+++
title = "url"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/net/url@go1.21.3

Package url parses URLs and implements query escaping.

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func JoinPath  <- go1.19

``` go 
func JoinPath(base string, elem ...string) (result string, err error)
```

JoinPath returns a URL string with the provided path elements joined to the existing path of base and the resulting path cleaned of any ./ or ../ elements.

### func PathEscape  <- go1.8

``` go 
func PathEscape(s string) string
```

PathEscape escapes the string so it can be safely placed inside a URL path segment, replacing special characters (including /) with %XX sequences as needed.

#### PathEscape  Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	path := url.PathEscape("my/cool+blog&about,stuff")
	fmt.Println(path)

}
Output:

my%2Fcool+blog&about%2Cstuff
```

### func PathUnescape  <- go1.8

``` go 
func PathUnescape(s string) (string, error)
```

PathUnescape does the inverse transformation of PathEscape, converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits.

PathUnescape is identical to QueryUnescape except that it does not unescape '+' to ' ' (space).

#### PathUnescape  Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	escapedPath := "my%2Fcool+blog&about%2Cstuff"
	path, err := url.PathUnescape(escapedPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)

}
Output:

my/cool+blog&about,stuff
```

### func QueryEscape 

``` go 
func QueryEscape(s string) string
```

QueryEscape escapes the string so it can be safely placed inside a URL query.

#### QueryEscape Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	query := url.QueryEscape("my/cool+blog&about,stuff")
	fmt.Println(query)

}
Output:

my%2Fcool%2Bblog%26about%2Cstuff
```

### func QueryUnescape 

``` go 
func QueryUnescape(s string) (string, error)
```

QueryUnescape does the inverse transformation of QueryEscape, converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits.

#### QueryUnescape Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	escapedQuery := "my%2Fcool%2Bblog%26about%2Cstuff"
	query, err := url.QueryUnescape(escapedQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(query)

}
Output:

my/cool+blog&about,stuff
```

## 类型

### type Error 

``` go 
type Error struct {
	Op  string
	URL string
	Err error
}
```

Error reports an error and the operation and URL that caused it.

#### (*Error) Error 

``` go 
func (e *Error) Error() string
```

#### (*Error) Temporary  <- go1.6

``` go 
func (e *Error) Temporary() bool
```

#### (*Error) Timeout  <- go1.6

``` go 
func (e *Error) Timeout() bool
```

#### (*Error) Unwrap  <- go1.13

``` go 
func (e *Error) Unwrap() error
```

### type EscapeError 

``` go 
type EscapeError string
```

#### (EscapeError) Error 

``` go 
func (e EscapeError) Error() string
```

### type InvalidHostError  <- go1.6

``` go 
type InvalidHostError string
```

#### (InvalidHostError) Error  <- go1.6

``` go 
func (e InvalidHostError) Error() string
```

### type URL 

``` go 
type URL struct {
	Scheme      string
	Opaque      string    // encoded opaque data
	User        *Userinfo // username and password information
	Host        string    // host or host:port
	Path        string    // path (relative paths may omit leading slash)
	RawPath     string    // encoded path hint (see EscapedPath method)
	OmitHost    bool      // do not emit empty host (authority)
	ForceQuery  bool      // append a query ('?') even if RawQuery is empty
	RawQuery    string    // encoded query values, without '?'
	Fragment    string    // fragment for references, without '#'
	RawFragment string    // encoded fragment hint (see EscapedFragment method)
}
```

A URL represents a parsed URL (technically, a URI reference).

The general form represented is:

```
[scheme:][//[userinfo@]host][/]path[?query][#fragment]
```

URLs that do not start with a slash after the scheme are interpreted as:

```
scheme:opaque[?query][#fragment]
```

Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/. A consequence is that it is impossible to tell which slashes in the Path were slashes in the raw URL and which were %2f. This distinction is rarely important, but when it is, the code should use the EscapedPath method, which preserves the original encoding of Path.

The RawPath field is an optional field which is only set when the default encoding of Path is different from the escaped path. See the EscapedPath method for more details.

URL's String method uses the EscapedPath method to obtain the path.

#### Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u)
}
Output:

https://google.com/search?q=golang
```

#### Example(Roundtrip)
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Parse + String preserve the original encoding.
	u, err := url.Parse("https://example.com/foo%2fbar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u.String())
}
Output:

/foo/bar
/foo%2fbar
https://example.com/foo%2fbar
```

#### func Parse 

``` go 
func Parse(rawURL string) (*URL, error)
```

Parse parses a raw url into a URL structure.

The url may be relative (a path, without a host) or absolute (starting with a scheme). Trying to parse a hostname and path without a scheme is invalid but may not necessarily return an error, due to parsing ambiguities.

#### func ParseRequestURI 

``` go 
func ParseRequestURI(rawURL string) (*URL, error)
```

ParseRequestURI parses a raw url into a URL structure. It assumes that url was received in an HTTP request, so the url is interpreted only as an absolute URI or an absolute path. The string url is assumed not to have a #fragment suffix. (Web browsers strip #fragment before sending the URL to a web server.)

#### (*URL) EscapedFragment  <- go1.15

``` go 
func (u *URL) EscapedFragment() string
```

EscapedFragment returns the escaped form of u.Fragment. In general there are multiple possible escaped forms of any fragment. EscapedFragment returns u.RawFragment when it is a valid escaping of u.Fragment. Otherwise EscapedFragment ignores u.RawFragment and computes an escaped form on its own. The String method uses EscapedFragment to construct its result. In general, code should call EscapedFragment instead of reading u.RawFragment directly.

##### EscapedFragment  Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://example.com/#x/y%2Fz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fragment:", u.Fragment)
	fmt.Println("RawFragment:", u.RawFragment)
	fmt.Println("EscapedFragment:", u.EscapedFragment())
}
Output:

Fragment: x/y/z
RawFragment: x/y%2Fz
EscapedFragment: x/y%2Fz
```

#### (*URL) EscapedPath  <- go1.5

``` go 
func (u *URL) EscapedPath() string
```

EscapedPath returns the escaped form of u.Path. In general there are multiple possible escaped forms of any path. EscapedPath returns u.RawPath when it is a valid escaping of u.Path. Otherwise EscapedPath ignores u.RawPath and computes an escaped form on its own. The String and RequestURI methods use EscapedPath to construct their results. In general, code should call EscapedPath instead of reading u.RawPath directly.

##### EscapedPath  Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("http://example.com/x/y%2Fz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Path:", u.Path)
	fmt.Println("RawPath:", u.RawPath)
	fmt.Println("EscapedPath:", u.EscapedPath())
}
Output:

Path: /x/y/z
RawPath: /x/y%2Fz
EscapedPath: /x/y%2Fz
```

#### (*URL) Hostname  <- go1.8

``` go 
func (u *URL) Hostname() string
```

Hostname returns u.Host, stripping any valid port number if present.

If the result is enclosed in square brackets, as literal IPv6 addresses are, the square brackets are removed from the result.

##### Hostname  Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.org:8000/path")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Hostname())
	u, err = url.Parse("https://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:17000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Hostname())
}
Output:

example.org
2001:0db8:85a3:0000:0000:8a2e:0370:7334
```

#### (*URL) IsAbs 

``` go 
func (u *URL) IsAbs() bool
```

IsAbs reports whether the URL is absolute. Absolute means that it has a non-empty scheme.

##### IsAbs Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := url.URL{Host: "example.com", Path: "foo"}
	fmt.Println(u.IsAbs())
	u.Scheme = "http"
	fmt.Println(u.IsAbs())
}
Output:

false
true
```

#### (*URL) JoinPath  <- go1.19

``` go 
func (u *URL) JoinPath(elem ...string) *URL
```

JoinPath returns a new URL with the provided path elements joined to any existing path and the resulting path cleaned of any ./ or ../ elements. Any sequences of multiple / characters will be reduced to a single /.

#### (*URL) MarshalBinary  <- go1.8

``` go 
func (u *URL) MarshalBinary() (text []byte, err error)
```

##### MarshalBinary  Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://example.org")
	b, err := u.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}
Output:

https://example.org
```

#### (*URL) Parse 

``` go 
func (u *URL) Parse(ref string) (*URL, error)
```

Parse parses a URL in the context of the receiver. The provided URL may be relative or absolute. Parse returns nil, err on parse failure, otherwise its return value is the same as ResolveReference.

##### Parse Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.org")
	if err != nil {
		log.Fatal(err)
	}
	rel, err := u.Parse("/foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rel)
	_, err = u.Parse(":foo")
	if _, ok := err.(*url.Error); !ok {
		log.Fatal(err)
	}
}
Output:

https://example.org/foo
```

#### (*URL) Port  <- go1.8

``` go 
func (u *URL) Port() string
```

Port returns the port part of u.Host, without the leading colon.

If u.Host doesn't contain a valid numeric port, Port returns an empty string.

##### Port  Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Port())
	u, err = url.Parse("https://example.org:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Port())
}
Output:


8080
```

#### (*URL) Query 

``` go 
func (u *URL) Query() Values
```

Query parses RawQuery and returns the corresponding values. It silently discards malformed value pairs. To check errors use ParseQuery.

##### Query Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.org/?a=1&a=2&b=&=3&&&&")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	fmt.Println(q["a"])
	fmt.Println(q.Get("b"))
	fmt.Println(q.Get(""))
}
Output:

[1 2]

3
```

#### (*URL) Redacted  <- go1.15

``` go 
func (u *URL) Redacted() string
```

Redacted is like String but replaces any password with "xxxxx". Only the password in u.URL is redacted.

##### Redacted  Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := &url.URL{
		Scheme: "https",
		User:   url.UserPassword("user", "password"),
		Host:   "example.com",
		Path:   "foo/bar",
	}
	fmt.Println(u.Redacted())
	u.User = url.UserPassword("me", "newerPassword")
	fmt.Println(u.Redacted())
}
Output:

https://user:xxxxx@example.com/foo/bar
https://me:xxxxx@example.com/foo/bar
```

#### (*URL) RequestURI 

``` go 
func (u *URL) RequestURI() string
```

RequestURI returns the encoded path?query or opaque?query string that would be used in an HTTP request for u.

##### RequestURI Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://example.org/path?foo=bar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.RequestURI())
}
Output:

/path?foo=bar
```

#### (*URL) ResolveReference 

``` go 
func (u *URL) ResolveReference(ref *URL) *URL
```

ResolveReference resolves a URI reference to an absolute URI from an absolute base URI u, per [RFC 3986 Section 5.2](https://rfc-editor.org/rfc/rfc3986.html#section-5.2). The URI reference may be relative or absolute. ResolveReference always returns a new URL instance, even if the returned URL is identical to either the base or reference. If ref is an absolute URL, then ResolveReference ignores base and returns a copy of ref.

##### ResolveReference Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("../../..//search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	base, err := url.Parse("http://example.com/directory/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base.ResolveReference(u))
}
Output:

http://example.com/search?q=dotnet
```

#### (*URL) String 

``` go 
func (u *URL) String() string
```

String reassembles the URL into a valid URL string. The general form of the result is one of:

```
scheme:opaque?query#fragment
scheme://userinfo@host/path?query#fragment
```

If u.Opaque is non-empty, String uses the first form; otherwise it uses the second form. Any non-ASCII characters in host are escaped. To obtain the path, String uses u.EscapedPath().

In the second form, the following rules apply:

- if u.Scheme is empty, scheme: is omitted.
- if u.User is nil, userinfo@ is omitted.
- if u.Host is empty, host/ is omitted.
- if u.Scheme and u.Host are empty and u.User is nil, the entire scheme://userinfo@host/ is omitted.
- if u.Host is non-empty and u.Path begins with a /, the form host/path does not add its own /.
- if u.RawQuery is empty, ?query is omitted.
- if u.Fragment is empty, #fragment is omitted.

##### String Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := &url.URL{
		Scheme:   "https",
		User:     url.UserPassword("me", "pass"),
		Host:     "example.com",
		Path:     "foo/bar",
		RawQuery: "x=1&y=2",
		Fragment: "anchor",
	}
	fmt.Println(u.String())
	u.Opaque = "opaque"
	fmt.Println(u.String())
}

```

#### (*URL) UnmarshalBinary  <- go1.8

``` go 
func (u *URL) UnmarshalBinary(text []byte) error
```

##### UnmarshalBinary  Example
``` go 
package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u := &url.URL{}
	err := u.UnmarshalBinary([]byte("https://example.org/foo"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", u)
}
Output:

https://example.org/foo
```

### type Userinfo 

``` go 
type Userinfo struct {
	// contains filtered or unexported fields
}
```

The Userinfo type is an immutable encapsulation of username and password details for a URL. An existing Userinfo value is guaranteed to have a username set (potentially empty, as allowed by [RFC 2396](https://rfc-editor.org/rfc/rfc2396.html)), and optionally a password.

#### func User 

``` go 
func User(username string) *Userinfo
```

User returns a Userinfo containing the provided username and no password set.

#### func UserPassword 

``` go 
func UserPassword(username, password string) *Userinfo
```

UserPassword returns a Userinfo containing the provided username and password.

This functionality should only be used with legacy web sites. [RFC 2396](https://rfc-editor.org/rfc/rfc2396.html) warns that interpreting Userinfo this way "is NOT RECOMMENDED, because the passing of authentication information in clear text (such as URI) has proven to be a security risk in almost every case where it has been used."

#### (*Userinfo) Password 

``` go 
func (u *Userinfo) Password() (string, bool)
```

Password returns the password in case it is set, and whether it is set.

#### (*Userinfo) String 

``` go 
func (u *Userinfo) String() string
```

String returns the encoded userinfo information in the standard form of "username[:password]".

#### (*Userinfo) Username 

``` go 
func (u *Userinfo) Username() string
```

Username returns the username.

### type Values 

``` go 
type Values map[string][]string
```

Values maps a string key to a list of values. It is typically used for query parameters and form values. Unlike in the http.Header map, the keys in a Values map are case-sensitive.

#### Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
}
Output:

Ava
Jess
[Jess Sarah Zoe]
```

#### func ParseQuery 

``` go 
func ParseQuery(query string) (Values, error)
```

ParseQuery parses the URL-encoded query string and returns a map listing the values specified for each key. ParseQuery always returns a non-nil map containing all the valid query parameters found; err describes the first decoding error encountered, if any.

Query is expected to be a list of key=value settings separated by ampersands. A setting without an equals sign is interpreted as a key set to an empty value. Settings containing a non-URL-encoded semicolon are considered invalid.

##### ParseQuery Example
``` go 
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	m, err := url.ParseQuery(`x=1&y=2&y=3`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(toJSON(m))
}

func toJSON(m any) string {
	js, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(string(js), ",", ", ")
}
Output:

{"x":["1"], "y":["2", "3"]}
```

#### (Values) Add 

``` go 
func (v Values) Add(key, value string)
```

Add adds the value to key. It appends to any existing values associated with key.

##### Add Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("cat sounds", "meow")
	v.Add("cat sounds", "mew")
	v.Add("cat sounds", "mau")
	fmt.Println(v["cat sounds"])

}
Output:

[meow mew mau]
```

#### (Values) Del 

``` go 
func (v Values) Del(key string)
```

Del deletes the values associated with key.

##### Del Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("cat sounds", "meow")
	v.Add("cat sounds", "mew")
	v.Add("cat sounds", "mau")
	fmt.Println(v["cat sounds"])

	v.Del("cat sounds")
	fmt.Println(v["cat sounds"])

}
Output:

[meow mew mau]
[]
```

#### (Values) Encode 

``` go 
func (v Values) Encode() string
```

Encode encodes the values into "URL encoded" form ("bar=baz&foo=quux") sorted by key.

##### Encode Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("cat sounds", "meow")
	v.Add("cat sounds", "mew/")
	v.Add("cat sounds", "mau$")
	fmt.Println(v.Encode())

}
Output:

cat+sounds=meow&cat+sounds=mew%2F&cat+sounds=mau%24
```

#### (Values) Get 

``` go 
func (v Values) Get(key string) string
```

Get gets the first value associated with the given key. If there are no values associated with the key, Get returns the empty string. To access multiple values, use the map directly.

##### Get Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("cat sounds", "meow")
	v.Add("cat sounds", "mew")
	v.Add("cat sounds", "mau")
	fmt.Printf("%q\n", v.Get("cat sounds"))
	fmt.Printf("%q\n", v.Get("dog sounds"))

}
Output:

"meow"
""
```

#### (Values) Has  <- go1.17

``` go 
func (v Values) Has(key string) bool
```

Has checks whether a given key is set.

##### Has  Example
``` go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("cat sounds", "meow")
	v.Add("cat sounds", "mew")
	v.Add("cat sounds", "mau")
	fmt.Println(v.Has("cat sounds"))
	fmt.Println(v.Has("dog sounds"))

}
Output:

true
false
```

#### (Values) Set 

``` go 
func (v Values) Set(key, value string)
```

Set sets the key to value. It replaces any existing values.

##### Set  Example

```go 
package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("cat sounds", "meow")
	v.Add("cat sounds", "mew")
	v.Add("cat sounds", "mau")
	fmt.Println(v["cat sounds"])

	v.Set("cat sounds", "meow")
	fmt.Println(v["cat sounds"])

}
Output:

[meow mew mau]
[meow]
```

