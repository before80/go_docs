+++
title = "http/cookiejar"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# cookiejar

https://pkg.go.dev/net/http/cookiejar@go1.20.1



Package cookiejar implements an in-memory [RFC 6265](https://rfc-editor.org/rfc/rfc6265.html)-compliant http.CookieJar.







## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type [Jar](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookiejar/jar.go;l=61) 

``` go 
type Jar struct {
	// contains filtered or unexported fields
}
```

Jar implements the http.CookieJar interface from the net/http package.

#### func [New](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookiejar/jar.go;l=78) 

``` go 
func New(o *Options) (*Jar, error)
```

New returns a new cookie jar. A nil *Options is equivalent to a zero Options.

##### Example
``` go 
```

#### (*Jar) [Cookies](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookiejar/jar.go;l=157) 

``` go 
func (j *Jar) Cookies(u *url.URL) (cookies []*http.Cookie)
```

Cookies implements the Cookies method of the http.CookieJar interface.

It returns an empty slice if the URL's scheme is not HTTP or HTTPS.

#### (*Jar) [SetCookies](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookiejar/jar.go;l=232) 

``` go 
func (j *Jar) SetCookies(u *url.URL, cookies []*http.Cookie)
```

SetCookies implements the SetCookies method of the http.CookieJar interface.

It does nothing if the URL's scheme is not HTTP or HTTPS.

### type [Options](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookiejar/jar.go;l=50) 

``` go 
type Options struct {
	// PublicSuffixList is the public suffix list that determines whether
	// an HTTP server can set a cookie for a domain.
	//
	// A nil value is valid and may be useful for testing but it is not
	// secure: it means that the HTTP server for foo.co.uk can set a cookie
	// for bar.co.uk.
	PublicSuffixList PublicSuffixList
}
```

Options are the options for creating a new Jar.

### type [PublicSuffixList](https://cs.opensource.google/go/go/+/go1.20.1:src/net/http/cookiejar/jar.go;l=35) 

``` go 
type PublicSuffixList interface {
	// PublicSuffix returns the public suffix of domain.
	//
	// TODO: specify which of the caller and callee is responsible for IP
	// addresses, for leading and trailing dots, for case sensitivity, and
	// for IDN/Punycode.
	PublicSuffix(domain string) string

	// String returns a description of the source of this public suffix
	// list. The description will typically contain something like a time
	// stamp or version number.
	String() string
}
```

PublicSuffixList provides the public suffix of a domain. For example:

- the public suffix of "example.com" is "com",
- the public suffix of "foo1.foo2.foo3.co.uk" is "co.uk", and
- the public suffix of "bar.pvt.k12.ma.us" is "pvt.k12.ma.us".

Implementations of PublicSuffixList must be safe for concurrent use by multiple goroutines.

An implementation that always returns "" is valid and may be useful for testing but it is not secure: it means that the HTTP server for foo.com can set a cookie for bar.com.

A public suffix list implementation is in the package golang.org/x/net/publicsuffix.