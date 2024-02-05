+++
title = "redirect"
weight = 180
date = 2023-07-09T21:57:33+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Redirect

> 原文：[https://echo.labstack.com/docs/middleware/redirect](https://echo.labstack.com/docs/middleware/redirect)

## HTTPS Redirect

HTTPS redirect middleware redirects http requests to https. For example, [http://labstack.com](http://labstack.com/) will be redirected to [https://labstack.com](https://labstack.com/).

### Usage

```go
e := echo.New()
e.Pre(middleware.HTTPSRedirect())
```



## HTTPS WWW Redirect

HTTPS WWW redirect redirects http requests to www https. For example, [http://labstack.com](http://labstack.com/) will be redirected to [https://www.labstack.com](https://www.labstack.com/).

## Usage

```go
e := echo.New()
e.Pre(middleware.HTTPSWWWRedirect())
```



## HTTPS NonWWW Redirect

HTTPS NonWWW redirect redirects http requests to https non www. For example, [http://www.labstack.com](http://www.labstack.com/) will be redirect to [https://labstack.com](https://labstack.com/).

### Usage

```go
e := echo.New()
e.Pre(middleware.HTTPSNonWWWRedirect())
```



## WWW Redirect

WWW redirect redirects non www requests to www.

For example, [http://labstack.com](http://labstack.com/) will be redirected to [http://www.labstack.com](http://www.labstack.com/).

### Usage

```go
e := echo.New()
e.Pre(middleware.WWWRedirect())
```



## NonWWW Redirect

NonWWW redirect redirects www requests to non www. For example, [http://www.labstack.com](http://www.labstack.com/) will be redirected to [http://labstack.com](http://labstack.com/).

### Usage

```go
e := echo.New()
e.Pre(middleware.NonWWWRedirect())
```



## Custom Configuration

### Usage

```go
e := echo.New()
e.Use(middleware.HTTPSRedirectWithConfig(middleware.RedirectConfig{
  Code: http.StatusTemporaryRedirect,
}))
```



Example above will redirect the request HTTP to HTTPS with status code `307 - StatusTemporaryRedirect`.

## Configuration

```go
RedirectConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Status code to be used when redirecting the request.
  // Optional. Default value http.StatusMovedPermanently.
  Code int `json:"code"`
}
```



### Default Configuration*

```go
DefaultRedirectConfig = RedirectConfig{
  Skipper: DefaultSkipper,
  Code:    http.StatusMovedPermanently,
}
```