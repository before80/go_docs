+++
title = "url"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/url@go1.23.0](https://pkg.go.dev/net/url@go1.23.0)

Package url parses URLs and implements query escaping.

​	 url 包解析 URL 并实现查询转义。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func JoinPath <- go1.19

```go
func JoinPath(base string, elem ...string) (result string, err error)
```

JoinPath returns a URL string with the provided path elements joined to the existing path of base and the resulting path cleaned of any ./ or ../ elements.

​	JoinPath 返回一个 URL 字符串，其中提供的路径元素与 base 的现有路径连接，并且结果路径已清除所有 ./ 或 ../ 元素。

### func PathEscape <- go1.8

```go
func PathEscape(s string) string
```

PathEscape escapes the string so it can be safely placed inside a URL path segment, replacing special characters (including /) with %XX sequences as needed.

​	PathEscape 转义字符串，以便可以安全地将其置于 URL 路径段中，根据需要将特殊字符（包括 /）替换为 %XX 序列。

#### PathEscape Example

```go
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

### func PathUnescape <- go1.8

```go
func PathUnescape(s string) (string, error)
```

PathUnescape does the inverse transformation of PathEscape, converting each 3-byte encoded substring of the form “%AB” into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits.

​	PathUnescape 执行 PathEscape 的逆转换，将每个 3 字节编码的“%AB”形式的子字符串转换为十六进制解码的字节 0xAB。如果任何 % 后面没有两个十六进制数字，它将返回一个错误。

PathUnescape is identical to QueryUnescape except that it does not unescape ‘+’ to ’ ’ (space).

​	PathUnescape 与 QueryUnescape 相同，只是它不会将“+”转义为“'”（空格）。

#### PathUnescape Example

```go
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

```go
func QueryEscape(s string) string
```

QueryEscape escapes the string so it can be safely placed inside a URL query.

​	QueryEscape 转义字符串，以便可以安全地将其置于 URL 查询中。

#### QueryEscape Example

```go
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

```go
func QueryUnescape(s string) (string, error)
```

QueryUnescape does the inverse transformation of QueryEscape, converting each 3-byte encoded substring of the form “%AB” into the hex-decoded byte 0xAB. It returns an error if any % is not followed by two hexadecimal digits.

​	QueryUnescape 执行 QueryEscape 的逆转换，将每个形式为“%AB”的 3 字节编码子字符串转换为十六进制解码的字节 0xAB。如果任何 % 后面没有两个十六进制数字，它将返回一个错误。

#### QueryUnescape Example

```go
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

```go
type Error struct {
	Op  string
	URL string
	Err error
}
```

Error reports an error and the operation and URL that caused it.

​	Error 报告错误以及导致错误的操作和 URL。

#### (*Error) Error 

```go
func (e *Error) Error() string
```

#### (*Error) Temporary <- go1.6

```go
func (e *Error) Temporary() bool
```

#### (*Error) Timeout <- go1.6

```go
func (e *Error) Timeout() bool
```

#### (*Error) Unwrap <- go1.13

```go
func (e *Error) Unwrap() error
```

### type EscapeError

```go
type EscapeError string
```

#### (EscapeError) Error

```go
func (e EscapeError) Error() string
```

### type InvalidHostError <- go1.6

```go
type InvalidHostError string
```

#### (InvalidHostError) Error <- go1.6

```go
func (e InvalidHostError) Error() string
```

### type URL

```go
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

​	URL 表示已解析的 URL（从技术上讲，是 URI 引用）。

The general form represented is:

​	表示的一般形式为：

```
[scheme:][//[userinfo@]host][/]path[?query][#fragment]
```

URLs that do not start with a slash after the scheme are interpreted as:

​	方案后不以斜杠开头的 URL 被解释为：

```
scheme:opaque[?query][#fragment]
```

Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/. A consequence is that it is impossible to tell which slashes in the Path were slashes in the raw URL and which were %2f. This distinction is rarely important, but when it is, the code should use the EscapedPath method, which preserves the original encoding of Path.

​	请注意，路径字段以解码形式存储：/%47%6f%2f 变为 /Go/。因此，无法分辨路径中的哪些斜杠是原始 URL 中的斜杠，哪些是 %2f。这种区别很少重要，但当它很重要时，代码应使用 EscapedPath 方法，该方法保留路径的原始编码。

The RawPath field is an optional field which is only set when the default encoding of Path is different from the escaped path. See the EscapedPath method for more details.

​	RawPath 字段是一个可选字段，仅在路径的默认编码与转义路径不同时设置。有关更多详细信息，请参阅 EscapedPath 方法。

URL’s String method uses the EscapedPath method to obtain the path.

​	URL 的 String 方法使用 EscapedPath 方法获取路径。

#### Example 

```go
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

```go
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

```go
func Parse(rawURL string) (*URL, error)
```

Parse parses a raw url into a URL structure.

​	Parse 将原始 URL 解析为 URL 结构。

The url may be relative (a path, without a host) or absolute (starting with a scheme). Trying to parse a hostname and path without a scheme is invalid but may not necessarily return an error, due to parsing ambiguities.

​	URL 可以是相对的（路径，没有主机）或绝对的（以方案开头）。尝试解析没有方案的主机名和路径是无效的，但由于解析歧义，可能不一定返回错误。

#### func ParseRequestURI

```go
func ParseRequestURI(rawURL string) (*URL, error)
```

ParseRequestURI parses a raw url into a URL structure. It assumes that url was received in an HTTP request, so the url is interpreted only as an absolute URI or an absolute path. The string url is assumed not to have a #fragment suffix. (Web browsers strip #fragment before sending the URL to a web server.)

​	ParseRequestURI 将原始 URL 解析为 URL 结构。它假定在 HTTP 请求中收到了 URL，因此 URL 仅解释为绝对 URI 或绝对路径。假定字符串 URL 没有 #fragment 后缀。（Web 浏览器在将 URL 发送到 Web 服务器之前会删除 #fragment。）

#### (*URL) EscapedFragment <- go1.15

```go
func (u *URL) EscapedFragment() string
```

EscapedFragment returns the escaped form of u.Fragment. In general there are multiple possible escaped forms of any fragment. EscapedFragment returns u.RawFragment when it is a valid escaping of u.Fragment. Otherwise EscapedFragment ignores u.RawFragment and computes an escaped form on its own. The String method uses EscapedFragment to construct its result. In general, code should call EscapedFragment instead of reading u.RawFragment directly.

​	EscapedFragment 返回 u.Fragment 的转义形式。通常，任何片段都有多种可能的转义形式。当 u.RawFragment 是 u.Fragment 的有效转义时，EscapedFragment 返回 u.RawFragment。否则，EscapedFragment 会忽略 u.RawFragment 并自行计算转义形式。String 方法使用 EscapedFragment 来构造其结果。通常，代码应调用 EscapedFragment，而不是直接读取 u.RawFragment。

##### EscapedFragment Example

```go
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

#### (*URL) EscapedPath <- go1.5

```go
func (u *URL) EscapedPath() string
```

EscapedPath returns the escaped form of u.Path. In general there are multiple possible escaped forms of any path. EscapedPath returns u.RawPath when it is a valid escaping of u.Path. Otherwise EscapedPath ignores u.RawPath and computes an escaped form on its own. The String and RequestURI methods use EscapedPath to construct their results. In general, code should call EscapedPath instead of reading u.RawPath directly.

​	EscapedPath 返回 u.Path 的转义形式。通常，任何路径都有多种可能的转义形式。当 u.RawPath 是 u.Path 的有效转义时，EscapedPath 返回 u.RawPath。否则，EscapedPath 会忽略 u.RawPath 并自行计算转义形式。String 和 RequestURI 方法使用 EscapedPath 来构造其结果。通常，代码应调用 EscapedPath，而不是直接读取 u.RawPath。

##### EscapedPath Example

```go
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

#### (*URL) Hostname <- go1.8

```go
func (u *URL) Hostname() string
```

Hostname returns u.Host, stripping any valid port number if present.

​	Hostname 返回 u.Host，如果存在，则剥离任何有效的端口号。

If the result is enclosed in square brackets, as literal IPv6 addresses are, the square brackets are removed from the result.

​	如果结果用方括号括起来，就像字面 IPv6 地址一样，则从结果中删除方括号。

##### Hostname Example 

```go
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

```go
func (u *URL) IsAbs() bool
```

IsAbs reports whether the URL is absolute. Absolute means that it has a non-empty scheme.

​	IsAbs 报告 URL 是否是绝对的。绝对意味着它具有非空方案。

##### IsAbs Example

```go
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

#### (*URL) JoinPath <- go1.19

```go
func (u *URL) JoinPath(elem ...string) *URL
```

JoinPath returns a new URL with the provided path elements joined to any existing path and the resulting path cleaned of any ./ or ../ elements. Any sequences of multiple / characters will be reduced to a single /.

​	JoinPath 返回一个新的 URL，其中提供的路径元素连接到任何现有路径，并且结果路径已清除任何 ./ 或 ../ 元素。任何多个 / 字符的序列都将减少为单个 /。

#### (*URL) MarshalBinary <- go1.8

```go
func (u *URL) MarshalBinary() (text []byte, err error)
```

##### MarshalBinary Example

```go
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

```go
func (u *URL) Parse(ref string) (*URL, error)
```

Parse parses a URL in the context of the receiver. The provided URL may be relative or absolute. Parse returns nil, err on parse failure, otherwise its return value is the same as ResolveReference.

​	Parse 在接收者的上下文中解析 URL。提供的 URL 可能为相对或绝对 URL。Parse 在解析失败时返回 nil 和 err，否则其返回值与 ResolveReference 相同。

##### Parse Example

```go
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

#### (*URL) Port <- go1.8

```go
func (u *URL) Port() string
```

Port returns the port part of u.Host, without the leading colon.

​	Port 返回 u.Host 的端口部分，不带前导冒号。

If u.Host doesn’t contain a valid numeric port, Port returns an empty string.

​	如果 u.Host 不包含有效的数字端口，Port 返回空字符串。

##### Port Example 

```go
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

```go
func (u *URL) Query() Values
```

Query parses RawQuery and returns the corresponding values. It silently discards malformed value pairs. To check errors use ParseQuery.

​	Query 解析 RawQuery 并返回相应的值。它会静默地丢弃格式错误的值对。要检查错误，请使用 ParseQuery。

##### Query Example

```go
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

#### (*URL) Redacted <- go1.15

```go
func (u *URL) Redacted() string
```

Redacted is like String but replaces any password with “xxxxx”. Only the password in u.URL is redacted.

​	Redacted 类似于 String，但会将任何密码替换为“xxxxx”。仅对 u.URL 中的密码进行编辑。

##### Redacted Example

```go
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

```go
func (u *URL) RequestURI() string
```

RequestURI returns the encoded path?query or opaque?query string that would be used in an HTTP request for u.

​	RequestURI 返回编码的 path?query 或 opaque?query 字符串，该字符串将用于 u 的 HTTP 请求中。

##### RequestURI Example

```go
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

```go
func (u *URL) ResolveReference(ref *URL) *URL
```

ResolveReference resolves a URI reference to an absolute URI from an absolute base URI u, per [RFC 3986 Section 5.2](https://rfc-editor.org/rfc/rfc3986.html#section-5.2). The URI reference may be relative or absolute. ResolveReference always returns a new URL instance, even if the returned URL is identical to either the base or reference. If ref is an absolute URL, then ResolveReference ignores base and returns a copy of ref.

​	ResolveReference 根据 RFC 3986 第 5.2 节，将 URI 引用解析为绝对基本 URI u 的绝对 URI。URI 引用可以是相对的或绝对的。ResolveReference 始终返回一个新的 URL 实例，即使返回的 URL 与基本 URL 或引用 URL 相同。如果 ref 是绝对 URL，则 ResolveReference 会忽略基本 URL 并返回 ref 的副本。

##### ResolveReference Example

```go
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

```go
func (u *URL) String() string
```

String reassembles the URL into a valid URL string. The general form of the result is one of:

​	String 将 URL 重新组合成有效的 URL 字符串。结果的一般形式之一为：

```
scheme:opaque?query#fragment
scheme://userinfo@host/path?query#fragment
```

If u.Opaque is non-empty, String uses the first form; otherwise it uses the second form. Any non-ASCII characters in host are escaped. To obtain the path, String uses u.EscapedPath().

​	如果 u.Opaque 为非空，String 使用第一种形式；否则，它使用第二种形式。host 中的任何非 ASCII 字符都将被转义。要获取路径，String 使用 u.EscapedPath()。

In the second form, the following rules apply:

​	在第二种形式中，应用以下规则：

- if u.Scheme is empty, scheme: is omitted.
  如果 u.Scheme 为空，则省略 scheme:。
- if u.User is nil, userinfo@ is omitted.
  如果 u.User 为 nil，则省略 userinfo@。
- if u.Host is empty, host/ is omitted.
  如果 u.Host 为空，则省略 host/。
- if u.Scheme and u.Host are empty and u.User is nil, the entire scheme://userinfo@host/ is omitted.
  如果 u.Scheme 和 u.Host 为空，并且 u.User 为 nil，则省略整个 scheme://userinfo@host/。
- if u.Host is non-empty and u.Path begins with a /, the form host/path does not add its own /.
  如果 u.Host 为非空，并且 u.Path 以 / 开头，则形式 host/path 不会添加其自己的 /。
- if u.RawQuery is empty, ?query is omitted.
  如果 u.RawQuery 为空，则省略 ?query。
- if u.Fragment is empty, #fragment is omitted.
  如果 u.Fragment 为空，则省略 #fragment。

##### String Example 

```go
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

#### (*URL) UnmarshalBinary <- go1.8

```go
func (u *URL) UnmarshalBinary(text []byte) error
```

##### UnmarshalBinary Example

```go
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

```go
type Userinfo struct {
	// contains filtered or unexported fields
}
```

The Userinfo type is an immutable encapsulation of username and password details for a URL. An existing Userinfo value is guaranteed to have a username set (potentially empty, as allowed by [RFC 2396](https://rfc-editor.org/rfc/rfc2396.html)), and optionally a password.

​	Userinfo 类型是对 URL 的用户名和密码详细信息的不可变封装。现有 Userinfo 值保证设置了用户名（可能为空，RFC 2396 允许），并且可以选择密码。

#### func User

```go
func User(username string) *Userinfo
```

User returns a Userinfo containing the provided username and no password set.

​	User 返回包含提供的用户名且未设置密码的 Userinfo。

#### func UserPassword

```go
func UserPassword(username, password string) *Userinfo
```

UserPassword returns a Userinfo containing the provided username and password.

​	UserPassword 返回包含提供的用户名和密码的 Userinfo。

This functionality should only be used with legacy web sites. [RFC 2396](https://rfc-editor.org/rfc/rfc2396.html) warns that interpreting Userinfo this way “is NOT RECOMMENDED, because the passing of authentication information in clear text (such as URI) has proven to be a security risk in almost every case where it has been used.”

​	此功能仅应与旧版网站配合使用。RFC 2396 警告以这种方式解释 Userinfo “不推荐，因为以明文（例如 URI）传递身份验证信息已被证明在几乎所有使用它的情况下都是安全风险。”

#### (*Userinfo) Password

```go
func (u *Userinfo) Password() (string, bool)
```

Password returns the password in case it is set, and whether it is set.

​	Password 返回设置的密码（如果已设置）以及是否已设置。

#### (*Userinfo) String

```go
func (u *Userinfo) String() string
```

String returns the encoded userinfo information in the standard form of “username[:password]”.

​	String 返回标准形式“username[:password]”中编码的用户信息。

#### (*Userinfo) Username

```go
func (u *Userinfo) Username() string
```

Username returns the username.

​	Username 返回用户名。

### type Values

```go
type Values map[string][]string
```

Values maps a string key to a list of values. It is typically used for query parameters and form values. Unlike in the http.Header map, the keys in a Values map are case-sensitive.

​	Values 将字符串键映射到值列表。它通常用于查询参数和表单值。与 http.Header 映射不同，Values 映射中的键区分大小写。

#### Example

```go
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

```go
func ParseQuery(query string) (Values, error)
```

ParseQuery parses the URL-encoded query string and returns a map listing the values specified for each key. ParseQuery always returns a non-nil map containing all the valid query parameters found; err describes the first decoding error encountered, if any.

​	ParseQuery 解析 URL 编码的查询字符串，并返回一个映射，其中列出了为每个键指定的值。ParseQuery 始终返回一个非 nil 映射，其中包含找到的所有有效查询参数；err 描述遇到的第一个解码错误（如果有）。

Query is expected to be a list of key=value settings separated by ampersands. A setting without an equals sign is interpreted as a key set to an empty value. Settings containing a non-URL-encoded semicolon are considered invalid.

​	Query 预计是通过与号分隔的一系列 key=value 设置。没有等号的设置将被解释为一个键，该键设置为一个空值。包含未经 URL 编码的分号的设置被认为无效。

##### ParseQuery Example

```go
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

```go
func (v Values) Add(key, value string)
```

Add adds the value to key. It appends to any existing values associated with key.

​	Add 将值添加到键。它追加到与键关联的任何现有值。

##### Add Example

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

}
Output:

[meow mew mau]
```

#### (Values) Del

```go
func (v Values) Del(key string)
```

Del deletes the values associated with key.

​	Del 删除与键关联的值。

##### Del Example

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

	v.Del("cat sounds")
	fmt.Println(v["cat sounds"])

}
Output:

[meow mew mau]
[]
```

#### (Values) Encode

```go
func (v Values) Encode() string
```

Encode encodes the values into “URL encoded” form (“bar=baz&foo=quux”) sorted by key.

​	Encode 将值编码为“URL 编码”形式（“bar=baz&foo=quux”），按键排序。

##### Encode Example

```go
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

```go
func (v Values) Get(key string) string
```

Get gets the first value associated with the given key. If there are no values associated with the key, Get returns the empty string. To access multiple values, use the map directly.

​	Get 获取与给定键关联的第一个值。如果与键关联的值不存在，则 Get 返回空字符串。要访问多个值，请直接使用映射。

##### Get Example

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
	fmt.Printf("%q\n", v.Get("cat sounds"))
	fmt.Printf("%q\n", v.Get("dog sounds"))

}
Output:

"meow"
""
```

#### (Values) Has <- go1.17

```go
func (v Values) Has(key string) bool
```

Has checks whether a given key is set.

​	Has 检查给定的键是否已设置。

##### Has Example

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
	fmt.Println(v.Has("cat sounds"))
	fmt.Println(v.Has("dog sounds"))

}
Output:

true
false
```

#### (Values) Set

```go
func (v Values) Set(key, value string)
```

Set sets the key to value. It replaces any existing values.

​	Set 将键设置为值。它将替换任何现有值。

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

