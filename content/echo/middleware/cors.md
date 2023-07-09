+++
title = "cors"
date = 2023-07-09T21:54:25+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# CORS

https://echo.labstack.com/docs/middleware/cors

CORS middleware implements [CORS](http://www.w3.org/TR/cors) specification. CORS gives web servers cross-domain access controls, which enable secure cross-domain data transfers.

## Usage

```go
e.Use(middleware.CORS())
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
}))
```



## Configuration

```go
CORSConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // AllowOrigin defines a list of origins that may access the resource.
  // Optional. Default value []string{"*"}.
  AllowOrigins []string `yaml:"allow_origins"`

  // AllowOriginFunc is a custom function to validate the origin. It takes the
  // origin as an argument and returns true if allowed or false otherwise. If
  // an error is returned, it is returned by the handler. If this option is
  // set, AllowOrigins is ignored.
  // Optional.
  AllowOriginFunc func(origin string) (bool, error) `yaml:"allow_origin_func"`

  // AllowMethods defines a list methods allowed when accessing the resource.
  // This is used in response to a preflight request.
  // Optional. Default value DefaultCORSConfig.AllowMethods.
  AllowMethods []string `yaml:"allow_methods"`

  // AllowHeaders defines a list of request headers that can be used when
  // making the actual request. This is in response to a preflight request.
  // Optional. Default value []string{}.
  AllowHeaders []string `yaml:"allow_headers"`

  // AllowCredentials indicates whether or not the response to the request
  // can be exposed when the credentials flag is true. When used as part of
  // a response to a preflight request, this indicates whether or not the
  // actual request can be made using credentials.
  // Optional. Default value false.
  AllowCredentials bool `yaml:"allow_credentials"`

  // ExposeHeaders defines a whitelist headers that clients are allowed to
  // access.
  // Optional. Default value []string{}.
  ExposeHeaders []string `yaml:"expose_headers"`

  // MaxAge indicates how long (in seconds) the results of a preflight request
  // can be cached.
  // Optional. Default value 0.
  MaxAge int `yaml:"max_age"`
}
```



### Default Configuration

```go
DefaultCORSConfig = CORSConfig{
  Skipper:      DefaultSkipper,
  AllowOrigins: []string{"*"},
  AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
}
```