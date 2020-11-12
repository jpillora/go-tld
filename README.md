# TLD Parser in Go

[![GoDoc](https://godoc.org/github.com/jpillora/go-tld?status.svg)](https://pkg.go.dev/github.com/jpillora/go-tld?tab=doc)  [![Tests](https://github.com/jpillora/go-tld/workflows/Tests/badge.svg)](https://github.com/jpillora/go-tld/actions?workflow=Tests)

The `tld` package has the same API ([see godoc](http://godoc.org/github.com/jpillora/go-tld)) as `net/url` except `tld.URL` contains extra fields: `Subdomain`, `Domain`, `TLD` and `Port`.

_Note:_ This is a wrapper around http://golang.org/x/net/publicsuffix.

### Install

```
go get github.com/jpillora/go-tld
```

### Usage

```go
package main

import (
	"fmt"

	tld "github.com/jpillora/go-tld"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://blog.google",
		"https://www.medi-cal.ca.gov/",
		"https://ato.gov.au",
		"http://a.very.complex-domain.co.uk:8080/foo/bar",
		"http://a.domain.that.is.unmanaged",
	}
	for _, url := range urls {
		u, _ := tld.Parse(url)
		fmt.Printf("%50s = [ %s ] [ %s ] [ %s ] [ %s ] [ %s ] [ %t ]\n",
			u, u.Subdomain, u.Domain, u.TLD, u.Port, u.Path, u.Icann)
	}
}
```

```
$ go run main.go
                                 http://google.com = [  ] [ google ] [ com ] [  ] [  ] [ true ]
                                http://blog.google = [  ] [ blog ] [ google ] [  ] [  ] [ true ]
                      https://www.medi-cal.ca.gov/ = [ www.medi-cal ] [ ca ] [ gov ] [  ] [ / ] [ true ]
                                https://ato.gov.au = [  ] [ ato ] [ gov.au ] [  ] [  ] [ true ]
   http://a.very.complex-domain.co.uk:8080/foo/bar = [ a.very ] [ complex-domain ] [ co.uk ] [ 8080 ] [ /foo/bar ] [ true ]
                 http://a.domain.that.is.unmanaged = [ a.domain.that ] [ is ] [ unmanaged ] [  ] [  ] [ false ]
```

#### MIT License

Copyright © 2019 Jaime Pillora &lt;dev@jpillora.com&gt;

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
