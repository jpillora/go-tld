//go:generate sh generate.sh

//Package tld has the same API as net/url except
//tld.URL contains extra fields: Subdomain, Domain, TLD and Port.
package tld

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

//URL embeds net/url and adds extra fields ontop
type URL struct {
	Subdomain, Domain, TLD, Port string
	ICANN                        bool
	*url.URL
}

//Parse mirrors net/url.Parse except instead it returns
//a tld.URL, which contains extra fields.
func Parse(s string) (*URL, error) {
	url, err := url.Parse(s)
	if err != nil {
		return nil, err
	}
	if url.Host == "" {
		return &URL{URL: url}, nil
	}
	dom, port := domainPort(url.Host)
	//etld+1
	etld1, err := publicsuffix.EffectiveTLDPlusOne(dom)
	suffix, icann := publicsuffix.PublicSuffix(strings.ToLower(dom))
	// HACK: attempt to support valid domains which are not registered with ICAN
	if err != nil && !icann && suffix == dom {
		etld1 = dom
		err = nil
	}
	if err != nil {
		return nil, err
	}
	//convert to domain name, and tld
	i := strings.Index(etld1, ".")
	if i < 0 {
		return nil, fmt.Errorf("tld: failed parsing %q", s)
	}
	domName := etld1[0:i]
	tld := etld1[i+1:]
	//and subdomain
	sub := ""
	if rest := strings.TrimSuffix(dom, "."+etld1); rest != dom {
		sub = rest
	}
	return &URL{
		Subdomain: sub,
		Domain:    domName,
		TLD:       tld,
		Port:      port,
		ICANN:     icann,
		URL:       url,
	}, nil
}

func domainPort(host string) (string, string) {
	for i := len(host) - 1; i >= 0; i-- {
		if host[i] == ':' {
			return host[:i], host[i+1:]
		} else if host[i] < '0' || host[i] > '9' {
			return host, ""
		}
	}
	//will only land here if the string is all digits,
	//net/url should prevent that from happening
	return host, ""
}
