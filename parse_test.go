package tld

import (
	"testing"
)

func run(input, sub, dom, tld string, icann, errorExpected bool, t *testing.T) {

	u, err := Parse(input)

	if err != nil && errorExpected {
		return
	} else if err != nil {
		t.Errorf("errored '%s'", err)
	} else if u.TLD != tld {
		t.Errorf("should have TLD '%s', got '%s'", tld, u.TLD)
	} else if u.Domain != dom {
		t.Errorf("should have Domain '%s', got '%s'", dom, u.Domain)
	} else if u.Subdomain != sub {
		t.Errorf("should have Subdomain '%s', got '%s'", sub, u.Subdomain)
	} else if u.ICANN != icann {
		t.Errorf("should have Icann '%t', got '%t'", icann, u.ICANN)
	}
}

func Test0(t *testing.T) {
	run("http://foo.com", "", "foo", "com", true, false, t)
}

func Test1(t *testing.T) {
	run("http://zip.zop.foo.com", "zip.zop", "foo", "com", true, false, t)
}

func Test2(t *testing.T) {
	run("http://au.com.au", "", "au", "com.au", true, false, t)
}

func Test3(t *testing.T) {
	run("http://im.from.england.co.uk:1900", "im.from", "england", "co.uk", true, false, t)
}

func Test4(t *testing.T) {
	run("https://google.com", "", "google", "com", true, false, t)
}

func Test5(t *testing.T) {
	run("https://foo.notmanaged", "", "foo", "notmanaged", false, false, t)
}

func Test6(t *testing.T) {
	run("https://google.Com", "", "google", "Com", true, false, t)
}

func Test7(t *testing.T) {
	run("https://github.io", "", "github", "io", false, false, t)
}

func Test8(t *testing.T) {
	// test expects error
	run("https://no_dot_should_not_panic", "", "", "", false, true, t)
}

func Test9(t *testing.T) {
	// test expects error
	run("https://.start_with_dot_should_fail", "", "", "", false, true, t)
}

func Test10(t *testing.T) {
	// test expects error
	run("https://ends_with_dot_should_fail.", "", "", "", false, true, t)
}
