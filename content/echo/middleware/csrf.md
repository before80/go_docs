+++
title = "CSRF"
weight = 60
date = 2023-07-09T21:54:32+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# CSRF - 跨站请求伪造

https://echo.labstack.com/docs/middleware/csrf

​	跨站请求伪造（Cross-site request forgery），也称为一次点击攻击（one-click attack）、会话劫持（session riding），缩写为 CSRF（有时发音为 sea-surf）或 XSRF，是一种对网站的恶意利用（malicious exploit）方式，通过该方式，未经授权的命令可以从被网站信任的用户那里传输。

## Usage

```go
e.Use(middleware.CSRF())
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
  TokenLookup: "header:X-XSRF-TOKEN",
}))
```



​	以上示例使用 `X-XSRF-TOKEN` 请求头来提取 CSRF 令牌。

​	*从 Cookie 中读取令牌的示例配置*

```go
middleware.CSRFWithConfig(middleware.CSRFConfig{
    TokenLookup:    "cookie:_csrf",
    CookiePath:     "/",
    CookieDomain:   "example.com",
    CookieSecure:   true,
    CookieHTTPOnly: true,
    CookieSameSite: http.SameSiteStrictMode,
})
```



## 访问 CSRF 令牌

### 服务端

​	可以通过 `Echo#Context` 使用 `ContextKey` 来访问 CSRF 令牌，并通过模板传递给客户端。

### 客户端

​	可以通过 CSRF cookie 来访问 CSRF 令牌。

## Configuration

```go
CSRFConfig struct {
  // Skipper 定义一个用于跳过中间件的函数。
  Skipper Skipper

  // TokenLength 是生成的令牌的长度。
  TokenLength uint8 `json:"token_length"`
  // 可选。默认值 32

  // TokenLookup 是一个字符串，格式为 "<source>:<key>"，
  // 用于从请求中提取令牌。
  // 可选。默认值 "header:X-CSRF-Token"。
  // 可能的取值：
  // - "header:<name>"
  // - "form:<name>"
  // - "query:<name>"
  // - "cookie:<name>"
  TokenLookup string `json:"token_lookup"`

  // Context key to store generated CSRF token into context.
  // Optional. Default value "csrf".
  // 用于将生成的 CSRF 令牌存储到context中的Context键。
  // 可选。默认值 "csrf"。 
  ContextKey string `json:"context_key"`

  // CSRF cookie 的名称。该 cookie 用于存储 CSRF 令牌。
  // 可选。默认值 "_csrf"。
  CookieName string `json:"cookie_name"`

  // CSRF Cookie 的域。
  // 可选。默认值 none。
  CookieDomain string `json:"cookie_domain"`

  // CSRF cookie 的路径。
  // 可选。默认值 none。
  CookiePath string `json:"cookie_path"`

  // CSRF cookie 的最大有效期（以秒为单位）。
  // 可选。默认值 86400（24小时）。
  CookieMaxAge int `json:"cookie_max_age"`

  // 指示 CSRF cookie 是否为安全的。
  // 可选。默认值 false。
  CookieSecure bool `json:"cookie_secure"`

  // 指示 CSRF cookie 是否为 HTTP Only。
  // 可选。默认值 false。
  CookieHTTPOnly bool `json:"cookie_http_only"`
}
```



### Default Configuration

```go
DefaultCSRFConfig = CSRFConfig{
  Skipper:      DefaultSkipper,
  TokenLength:  32,
  TokenLookup:  "header:" + echo.HeaderXCSRFToken,
  ContextKey:   "csrf",
  CookieName:   "_csrf",
  CookieMaxAge: 86400,
}
```