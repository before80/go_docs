+++
title = "basic-auth"
date = 2023-07-09T21:53:36+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Basic Auth

https://echo.labstack.com/docs/middleware/basic-auth

Basic auth middleware provides an HTTP basic authentication.

- For valid credentials it calls the next handler.
- For missing or invalid credentials, it sends "401 - Unauthorized" response.

## Usage

```go
e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
    // Be careful to use constant time comparison to prevent timing attacks
    if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
        subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
        return true, nil
    }
    return false, nil
}))
```



## Custom Configuration

### Usage

```go
e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}))
```



## Configuration

```go
BasicAuthConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Validator is a function to validate BasicAuth credentials.
  // Required.
  Validator BasicAuthValidator

  // Realm is a string to define realm attribute of BasicAuth.
  // Default value "Restricted".
  Realm string
}
```



### Default Configuration

```go
DefaultBasicAuthConfig = BasicAuthConfig{
    Skipper: DefaultSkipper,
}
```