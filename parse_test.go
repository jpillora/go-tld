package tld

import "testing"

func run(input, sub, dom, tld string, icann bool, t *testing.T) {

	u, err := Parse(input)

	if err != nil {
		t.Errorf("errored '%s'", err)
	} else if u.TLD != tld {
		t.Errorf("should have TLD '%s', got '%s'", tld, u.TLD)
	} else if u.Domain != dom {
		t.Errorf("should have Domain '%s', got '%s'", dom, u.Domain)
	} else if u.Subdomain != sub {
		t.Errorf("should have Subdomain '%s', got '%s'", sub, u.Subdomain)
	} else if u.Icann != icann {
		t.Errorf("should have Icann '%t', got '%t'", icann, u.Icann)
	}
}

func Test0(t *testing.T) {
	run("http://foo.com", "", "foo", "com", true, t)
}

func Test1(t *testing.T) {
	run("http://zip.zop.foo.com", "zip.zop", "foo", "com", true, t)
}

func Test2(t *testing.T) {
	run("http://au.com.au", "", "au", "com.au", true, t)
}

func Test3(t *testing.T) {
	run("http://im.from.england.co.uk:1900", "im.from", "england", "co.uk", true, t)
}

func Test4(t *testing.T) {
	run("https://google.com", "", "google", "com", true, t)
}

func Test5(t *testing.T) {
	run("https://foo.notmanaged", "", "foo", "notmanaged", false, t)
}
