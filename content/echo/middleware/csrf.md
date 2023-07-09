+++
title = "csrf"
date = 2023-07-09T21:54:32+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# CSRF

https://echo.labstack.com/docs/middleware/csrf

Cross-site request forgery, also known as one-click attack or session riding and abbreviated as CSRF (sometimes pronounced sea-surf) or XSRF, is a type of malicious exploit of a website where unauthorized commands are transmitted from a user that the website trusts.

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



Example above uses `X-XSRF-TOKEN` request header to extract CSRF token.

*Example Configuration that reads token from Cookie*

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



## Accessing CSRF Token

### Server-side

CSRF token can be accessed from `Echo#Context` using `ContextKey` and passed to the client via template.

### Client-side

CSRF token can be accessed from CSRF cookie.

## Configuration

```go
CSRFConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // TokenLength is the length of the generated token.
  TokenLength uint8 `json:"token_length"`
  // Optional. Default value 32.

  // TokenLookup is a string in the form of "<source>:<key>" that is used
  // to extract token from the request.
  // Optional. Default value "header:X-CSRF-Token".
  // Possible values:
  // - "header:<name>"
  // - "form:<name>"
  // - "query:<name>"
  // - "cookie:<name>"
  TokenLookup string `json:"token_lookup"`

  // Context key to store generated CSRF token into context.
  // Optional. Default value "csrf".
  ContextKey string `json:"context_key"`

  // Name of the CSRF cookie. This cookie will store CSRF token.
  // Optional. Default value "_csrf".
  CookieName string `json:"cookie_name"`

  // Domain of the CSRF cookie.
  // Optional. Default value none.
  CookieDomain string `json:"cookie_domain"`

  // Path of the CSRF cookie.
  // Optional. Default value none.
  CookiePath string `json:"cookie_path"`

  // Max age (in seconds) of the CSRF cookie.
  // Optional. Default value 86400 (24hr).
  CookieMaxAge int `json:"cookie_max_age"`

  // Indicates if CSRF cookie is secure.
  // Optional. Default value false.
  CookieSecure bool `json:"cookie_secure"`

  // Indicates if CSRF cookie is HTTP only.
  // Optional. Default value false.
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