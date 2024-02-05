+++
title = "http/cookiejar"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/net/http/cookiejar@go1.21.3](https://pkg.go.dev/net/http/cookiejar@go1.21.3)

Package cookiejar implements an in-memory [RFC 6265](https://rfc-editor.org/rfc/rfc6265.html)-compliant http.CookieJar.

​	 `cookiejar` 包实现了一个内存中符合 [RFC 6265](https://rfc-editor.org/rfc/rfc6265.html) 标准的 http.CookieJar。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Jar 

``` go 
type Jar struct {
	// 包含已过滤或未公开的字段
}
```

Jar implements the http.CookieJar interface from the net/http package.

​	Jar 实现了 net/http 包中的 http.CookieJar 接口。

#### func New 

``` go 
func New(o *Options) (*Jar, error)
```

New returns a new cookie jar. A nil *Options is equivalent to a zero Options.

​	New 函数返回一个新的 cookie 存储器。空的 `*Options` 等同于零值的 Options。

##### New Example
``` go 
// Start a server to give us cookies.
// 启动一个服务器以提供给我们 cookies。
ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("Flavor"); err != nil {
		http.SetCookie(w, &http.Cookie{Name: "Flavor", Value: "Chocolate Chip"})
	} else {
		cookie.Value = "Oatmeal Raisin"
		http.SetCookie(w, cookie)
	}
}))
defer ts.Close()

u, err := url.Parse(ts.URL)
if err != nil {
	log.Fatal(err)
}

// All users of cookiejar should import "golang.org/x/net/publicsuffix"
// 所有使用 cookiejar 的用户应该导入 "golang.org/x/net/publicsuffix"。
jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
if err != nil {
	log.Fatal(err)
}

client := &http.Client{
	Jar: jar,
}

if _, err = client.Get(u.String()); err != nil {
	log.Fatal(err)
}

fmt.Println("After 1st request:")
for _, cookie := range jar.Cookies(u) {
	fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
}

if _, err = client.Get(u.String()); err != nil {
	log.Fatal(err)
}

fmt.Println("After 2nd request:")
for _, cookie := range jar.Cookies(u) {
	fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
}

// Output:

After 1st request:
  Flavor: Chocolate Chip
After 2nd request:
  Flavor: Oatmeal Raisin
```

#### (*Jar) Cookies 

``` go 
func (j *Jar) Cookies(u *url.URL) (cookies []*http.Cookie)
```

Cookies implements the Cookies method of the http.CookieJar interface.

​	Cookies 方法实现了 http.CookieJar 接口的 Cookies 方法。

It returns an empty slice if the URL's scheme is not HTTP or HTTPS.

​	如果 URL 的 scheme 不是 HTTP 或 HTTPS，则返回一个空切片。

#### (*Jar) SetCookies 

``` go 
func (j *Jar) SetCookies(u *url.URL, cookies []*http.Cookie)
```

SetCookies implements the SetCookies method of the http.CookieJar interface.

​	SetCookies 方法实现了 http.CookieJar 接口的 SetCookies 方法。

It does nothing if the URL's scheme is not HTTP or HTTPS.

​	如果 URL 的 scheme 不是 HTTP 或 HTTPS，则不执行任何操作。

### type Options 

``` go 
type Options struct {
    // PublicSuffixList is the public suffix list that determines whether
	// an HTTP server can set a cookie for a domain.
	//
	// A nil value is valid and may be useful for testing but it is not
	// secure: it means that the HTTP server for foo.co.uk can set a cookie
	// for bar.co.uk.
    // PublicSuffixList 是确定 HTTP 服务器是否能够为域设置 cookie 的公共后缀列表。
	//
	// nil 值是有效的，对于测试可能是有用的，但不安全：
	// 这意味着 foo.co.uk 的 HTTP 服务器可以为 bar.co.uk 设置 cookie。
	PublicSuffixList PublicSuffixList
}
```

Options are the options for creating a new Jar.

​	Options 是创建新的 Jar 的选项。

### type PublicSuffixList 

``` go 
type PublicSuffixList interface {
	// PublicSuffix returns the public suffix of domain.
	//
	// TODO: specify which of the caller and callee is responsible for IP
	// addresses, for leading and trailing dots, for case sensitivity, and
	// for IDN/Punycode.
    // PublicSuffix 返回域的公共后缀。
	//
	// TODO：指定调用方和被调用方谁负责 IP 地址、前导和尾随点、大小写敏感性以及 IDN/Punycode。
	PublicSuffix(domain string) string

	// String returns a description of the source of this public suffix
	// list. The description will typically contain something like a time
	// stamp or version number.
    // String 返回此公共后缀列表源的描述。
	// 描述通常会包含时间戳或版本号等内容。
	String() string
}
```

PublicSuffixList provides the public suffix of a domain. For example:

​	PublicSuffixList 提供了一个域的公共后缀。例如： 

- the public suffix of "example.com" is "com",
- "example.com" 的公共后缀是 "com"，
- the public suffix of "foo1.foo2.foo3.co.uk" is "co.uk", and
- "foo1.foo2.foo3.co.uk" 的公共后缀是 "co.uk"，以及
- the public suffix of "bar.pvt.k12.ma.us" is "pvt.k12.ma.us".
- "bar.pvt.k12.ma.us" 的公共后缀是 "pvt.k12.ma.us"。

Implementations of PublicSuffixList must be safe for concurrent use by multiple goroutines.

​	PublicSuffixList 的实现必须能够在多个 goroutine 中安全地进行并发使用。

An implementation that always returns "" is valid and may be useful for testing but it is not secure: it means that the HTTP server for foo.com can set a cookie for bar.com.

​	总是返回 "" 的实现是有效的，可能对于测试有用，但不安全： 这意味着 foo.com 的 HTTP 服务器可以为 bar.com 设置 cookie。

A public suffix list implementation is in the package golang.org/x/net/publicsuffix.

​	在 golang.org/x/net/publicsuffix 包中有一个公共后缀列表的实现。