+++
title = "url"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# url

https://pkg.go.dev/net/url@go1.20.1



Package url parses URLs and implements query escaping.

































## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [JoinPath](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1261)  <- go1.19

``` go linenums="1"
func JoinPath(base string, elem ...string) (result string, err error)
```

JoinPath returns a URL string with the provided path elements joined to the existing path of base and the resulting path cleaned of any ./ or ../ elements.

#### func [PathEscape](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=281)  <- go1.8

``` go linenums="1"
func PathEscape(s string) string
```

PathEscape escapes the string so it can be safely placed inside a URL path segment, replacing special characters (including /) with %XX sequences as needed.

##### Example
``` go linenums="1"
```

#### func [PathUnescape](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=194)  <- go1.8

``` go linenums="1"
func PathUnescape(s string) (string, error)
```

PathUnescape does the inverse transformation of PathEscape, converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits.

PathUnescape is identical to QueryUnescape except that it does not unescape '+' to ' ' (space).

##### Example
``` go linenums="1"
```

#### func [QueryEscape](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=275) 

``` go linenums="1"
func QueryEscape(s string) string
```

QueryEscape escapes the string so it can be safely placed inside a URL query.

##### Example
``` go linenums="1"
```

#### func [QueryUnescape](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=183) 

``` go linenums="1"
func QueryUnescape(s string) (string, error)
```

QueryUnescape does the inverse transformation of QueryEscape, converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits.

##### Example
``` go linenums="1"
```

## 类型

### type [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=23) 

``` go linenums="1"
type Error struct {
	Op  string
	URL string
	Err error
}
```

Error reports an error and the operation and URL that caused it.

#### (*Error) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=30) 

``` go linenums="1"
func (e *Error) Error() string
```

#### (*Error) [Temporary](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=39)  <- go1.6

``` go linenums="1"
func (e *Error) Temporary() bool
```

#### (*Error) [Timeout](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=32)  <- go1.6

``` go linenums="1"
func (e *Error) Timeout() bool
```

#### (*Error) [Unwrap](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=29)  <- go1.13

``` go linenums="1"
func (e *Error) Unwrap() error
```

### type [EscapeError](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=84) 

``` go linenums="1"
type EscapeError string
```

#### (EscapeError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=86) 

``` go linenums="1"
func (e EscapeError) Error() string
```

### type [InvalidHostError](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=90)  <- go1.6

``` go linenums="1"
type InvalidHostError string
```

#### (InvalidHostError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=92)  <- go1.6

``` go linenums="1"
func (e InvalidHostError) Error() string
```

### type [URL](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=362) 

``` go linenums="1"
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

##### Example
``` go linenums="1"
```

##### Example
``` go linenums="1"
```

#### func [Parse](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=466) 

``` go linenums="1"
func Parse(rawURL string) (*URL, error)
```

Parse parses a raw url into a URL structure.

The url may be relative (a path, without a host) or absolute (starting with a scheme). Trying to parse a hostname and path without a scheme is invalid but may not necessarily return an error, due to parsing ambiguities.

#### func [ParseRequestURI](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=487) 

``` go linenums="1"
func ParseRequestURI(rawURL string) (*URL, error)
```

ParseRequestURI parses a raw url into a URL structure. It assumes that url was received in an HTTP request, so the url is interpreted only as an absolute URI or an absolute path. The string url is assumed not to have a #fragment suffix. (Web browsers strip #fragment before sending the URL to a web server.)

#### (*URL) [EscapedFragment](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=760)  <- go1.15

``` go linenums="1"
func (u *URL) EscapedFragment() string
```

EscapedFragment returns the escaped form of u.Fragment. In general there are multiple possible escaped forms of any fragment. EscapedFragment returns u.RawFragment when it is a valid escaping of u.Fragment. Otherwise EscapedFragment ignores u.RawFragment and computes an escaped form on its own. The String method uses EscapedFragment to construct its result. In general, code should call EscapedFragment instead of reading u.RawFragment directly.

##### Example
``` go linenums="1"
```

#### (*URL) [EscapedPath](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=697)  <- go1.5

``` go linenums="1"
func (u *URL) EscapedPath() string
```

EscapedPath returns the escaped form of u.Path. In general there are multiple possible escaped forms of any path. EscapedPath returns u.RawPath when it is a valid escaping of u.Path. Otherwise EscapedPath ignores u.RawPath and computes an escaped form on its own. The String and RequestURI methods use EscapedPath to construct their results. In general, code should call EscapedPath instead of reading u.RawPath directly.

##### Example
``` go linenums="1"
```

#### (*URL) [Hostname](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1146)  <- go1.8

``` go linenums="1"
func (u *URL) Hostname() string
```

Hostname returns u.Host, stripping any valid port number if present.

If the result is enclosed in square brackets, as literal IPv6 addresses are, the square brackets are removed from the result.

##### Example
``` go linenums="1"
```

#### (*URL) [IsAbs](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1061) 

``` go linenums="1"
func (u *URL) IsAbs() bool
```

IsAbs reports whether the URL is absolute. Absolute means that it has a non-empty scheme.

##### Example
``` go linenums="1"
```

#### (*URL) [JoinPath](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1196)  <- go1.19

``` go linenums="1"
func (u *URL) JoinPath(elem ...string) *URL
```

JoinPath returns a new URL with the provided path elements joined to any existing path and the resulting path cleaned of any ./ or ../ elements. Any sequences of multiple / characters will be reduced to a single /.

#### (*URL) [MarshalBinary](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1180)  <- go1.8

``` go linenums="1"
func (u *URL) MarshalBinary() (text []byte, err error)
```

##### Example
``` go linenums="1"
```

#### (*URL) [Parse](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1068) 

``` go linenums="1"
func (u *URL) Parse(ref string) (*URL, error)
```

Parse parses a URL in the context of the receiver. The provided URL may be relative or absolute. Parse returns nil, err on parse failure, otherwise its return value is the same as ResolveReference.

##### Example
``` go linenums="1"
```

#### (*URL) [Port](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1154)  <- go1.8

``` go linenums="1"
func (u *URL) Port() string
```

Port returns the port part of u.Host, without the leading colon.

If u.Host doesn't contain a valid numeric port, Port returns an empty string.

##### Example
``` go linenums="1"
```

#### (*URL) [Query](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1117) 

``` go linenums="1"
func (u *URL) Query() Values
```

Query parses RawQuery and returns the corresponding values. It silently discards malformed value pairs. To check errors use ParseQuery.

##### Example
``` go linenums="1"
```

#### (*URL) [Redacted](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=863)  <- go1.15

``` go linenums="1"
func (u *URL) Redacted() string
```

Redacted is like String but replaces any password with "xxxxx". Only the password in u.URL is redacted.

##### Example
``` go linenums="1"
```

#### (*URL) [RequestURI](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1124) 

``` go linenums="1"
func (u *URL) RequestURI() string
```

RequestURI returns the encoded path?query or opaque?query string that would be used in an HTTP request for u.

##### Example
``` go linenums="1"
```

#### (*URL) [ResolveReference](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1082) 

``` go linenums="1"
func (u *URL) ResolveReference(ref *URL) *URL
```

ResolveReference resolves a URI reference to an absolute URI from an absolute base URI u, per [RFC 3986 Section 5.2](https://rfc-editor.org/rfc/rfc3986.html#section-5.2). The URI reference may be relative or absolute. ResolveReference always returns a new URL instance, even if the returned URL is identical to either the base or reference. If ref is an absolute URL, then ResolveReference ignores base and returns a copy of ref.

##### Example
``` go linenums="1"
```

#### (*URL) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=808) 

``` go linenums="1"
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

##### Example
``` go linenums="1"
```

#### (*URL) [UnmarshalBinary](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=1184)  <- go1.8

``` go linenums="1"
func (u *URL) UnmarshalBinary(text []byte) error
```

##### Example
``` go linenums="1"
```

### type [Userinfo](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=398) 

``` go linenums="1"
type Userinfo struct {
	// contains filtered or unexported fields
}
```

The Userinfo type is an immutable encapsulation of username and password details for a URL. An existing Userinfo value is guaranteed to have a username set (potentially empty, as allowed by [RFC 2396](https://rfc-editor.org/rfc/rfc2396.html)), and optionally a password.

#### func [User](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=378) 

``` go linenums="1"
func User(username string) *Userinfo
```

User returns a Userinfo containing the provided username and no password set.

#### func [UserPassword](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=390) 

``` go linenums="1"
func UserPassword(username, password string) *Userinfo
```

UserPassword returns a Userinfo containing the provided username and password.

This functionality should only be used with legacy web sites. [RFC 2396](https://rfc-editor.org/rfc/rfc2396.html) warns that interpreting Userinfo this way "is NOT RECOMMENDED, because the passing of authentication information in clear text (such as URI) has proven to be a security risk in almost every case where it has been used."

#### (*Userinfo) [Password](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=413) 

``` go linenums="1"
func (u *Userinfo) Password() (string, bool)
```

Password returns the password in case it is set, and whether it is set.

#### (*Userinfo) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=422) 

``` go linenums="1"
func (u *Userinfo) String() string
```

String returns the encoded userinfo information in the standard form of "username[:password]".

#### (*Userinfo) [Username](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=405) 

``` go linenums="1"
func (u *Userinfo) Username() string
```

Username returns the username.

### type [Values](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=879) 

``` go linenums="1"
type Values map[string][]string
```

Values maps a string key to a list of values. It is typically used for query parameters and form values. Unlike in the http.Header map, the keys in a Values map are case-sensitive.

##### Example
``` go linenums="1"
```

#### func [ParseQuery](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=929) 

``` go linenums="1"
func ParseQuery(query string) (Values, error)
```

ParseQuery parses the URL-encoded query string and returns a map listing the values specified for each key. ParseQuery always returns a non-nil map containing all the valid query parameters found; err describes the first decoding error encountered, if any.

Query is expected to be a list of key=value settings separated by ampersands. A setting without an equals sign is interpreted as a key set to an empty value. Settings containing a non-URL-encoded semicolon are considered invalid.

##### Example
``` go linenums="1"
```

#### (Values) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=904) 

``` go linenums="1"
func (v Values) Add(key, value string)
```

Add adds the value to key. It appends to any existing values associated with key.

##### Example
``` go linenums="1"
```

#### (Values) [Del](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=909) 

``` go linenums="1"
func (v Values) Del(key string)
```

Del deletes the values associated with key.

##### Example
``` go linenums="1"
```

#### (Values) [Encode](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=968) 

``` go linenums="1"
func (v Values) Encode() string
```

Encode encodes the values into "URL encoded" form ("bar=baz&foo=quux") sorted by key.

##### Example
``` go linenums="1"
```

#### (Values) [Get](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=885) 

``` go linenums="1"
func (v Values) Get(key string) string
```

Get gets the first value associated with the given key. If there are no values associated with the key, Get returns the empty string. To access multiple values, use the map directly.

##### Example
``` go linenums="1"
```

#### (Values) [Has](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=914)  <- go1.17

``` go linenums="1"
func (v Values) Has(key string) bool
```

Has checks whether a given key is set.

##### Example
``` go linenums="1"
```

#### (Values) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/net/url/url.go;l=898) 

``` go linenums="1"
func (v Values) Set(key, value string)
```

Set sets the key to value. It replaces any existing values.

```go linenums="1"
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

