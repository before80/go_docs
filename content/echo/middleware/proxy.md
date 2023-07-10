+++
title = "proxy"
weight = 150
date = 2023-07-09T21:56:40+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Proxy

https://echo.labstack.com/docs/middleware/proxy

Proxy provides an HTTP/WebSocket reverse proxy middleware. It forwards a request to upstream server using a configured load balancing technique.

### Usage

```go
url1, err := url.Parse("http://localhost:8081")
if err != nil {
  e.Logger.Fatal(err)
}
url2, err := url.Parse("http://localhost:8082")
if err != nil {
  e.Logger.Fatal(err)
}
e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
  {
    URL: url1,
  },
  {
    URL: url2,
  },
})))
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{}))
```



### Configuration

```go
// ProxyConfig defines the config for Proxy middleware.
  ProxyConfig struct {
    // Skipper defines a function to skip middleware.
    Skipper Skipper

    // Balancer defines a load balancing technique.
    // Required.
    Balancer ProxyBalancer

    // Rewrite defines URL path rewrite rules. The values captured in asterisk can be
    // retrieved by index e.g. $1, $2 and so on.
    Rewrite map[string]string

    // RegexRewrite defines rewrite rules using regexp.Rexexp with captures
    // Every capture group in the values can be retrieved by index e.g. $1, $2 and so on.
    RegexRewrite map[*regexp.Regexp]string

    // Context key to store selected ProxyTarget into context.
    // Optional. Default value "target".
    ContextKey string

    // To customize the transport to remote.
    // Examples: If custom TLS certificates are required.
    Transport http.RoundTripper

    // ModifyResponse defines function to modify response from ProxyTarget.
    ModifyResponse func(*http.Response) error
```



### Default Configuration

| Name       | Value          |
| ---------- | -------------- |
| Skipper    | DefaultSkipper |
| ContextKey | `target`       |

### Regex-based Rules

For advanced rewriting of proxy requests rules may also be defined using regular expression. Normal capture groups can be defined using `()` and referenced by index (`$1`, `$2`, ...) for the rewritten path.

`RegexRules` and normal `Rules` can be combined.

```go
  e.Use(ProxyWithConfig(ProxyConfig{
    Balancer: rrb,
    Rewrite: map[string]string{
      "^/v1/*":     "/v2/$1",
    },
    RegexRewrite: map[*regexp.Regexp]string{
      regexp.MustCompile("^/foo/([0-9].*)"):  "/num/$1",
      regexp.MustCompile("^/bar/(.+?)/(.*)"): "/baz/$2/$1",
    },
  }))
```



## [Example](https://echo.labstack.com/docs/cookbook/reverse-proxy)