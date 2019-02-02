package main

import (
	"fmt"

	"github.com/jpillora/go-tld"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://blog.google",
		"https://www.medi-cal.ca.gov/",
		"https://ato.gov.au",
		"http://a.very.complex-domain.co.uk:8080/foo/bar",
	}
	for _, url := range urls {
		u, _ := tld.Parse(url)
		fmt.Printf("%50s = [ %s ] [ %s ] [ %s ] [ %s ] [ %s ]\n",
			u, u.Subdomain, u.Domain, u.TLD, u.Port, u.Path)
	}
}
