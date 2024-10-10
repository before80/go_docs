+++
title = "cors"
weight = 50
date = 2023-07-09T21:54:25+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# CORS - 跨域资源共享 Cross Origin Resources Sharing

> 原文：[https://echo.labstack.com/docs/middleware/cors](https://echo.labstack.com/docs/middleware/cors)

​	CORS 中间件实现了 [CORS](http://www.w3.org/TR/cors) 规范。CORS 可以给 web 服务器提供跨域访问控制，从而实现安全的跨域数据传输。

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
  // Skipper 定义一个用于跳过中间件的函数。
  Skipper Skipper

  // AllowOrigin 定义可以访问资源的源的列表。
  // 可选。默认值 []string{"*"}。
  AllowOrigins []string `yaml:"allow_origins"`

  // AllowOriginFunc 是一个自定义函数，用于验证来源（origin）。
  // 它以来源作为实参，如果允许则返回 true，否则返回 false。
  // 如果返回错误，该错误将被处理程序返回。
  // 如果设置了此选项，则AllowOrigins将被忽略。
  // 可选。
  AllowOriginFunc func(origin string) (bool, error) `yaml:"allow_origin_func"`

  // AllowMethods 定义在访问资源时允许的方法列表。
  // 这用于响应预检（preflight）请求。
  // 可选。默认值 DefaultCORSConfig.AllowMethods。
  AllowMethods []string `yaml:"allow_methods"`

  // AllowHeaders 定义可以在实际请求中使用的请求头列表。
  // 这是对预检（preflight）请求的响应。
  // 可选。默认值 []string{}。
  AllowHeaders []string `yaml:"allow_headers"`

  // AllowCredentials 指示响应是否可以在凭证标志为 true 时公开（be exposed）。
  // 当作为预检（preflight）请求的一部分使用时，它指示是否可以使用凭证进行实际请求。
  // 可选。默认值 false。
  AllowCredentials bool `yaml:"allow_credentials"`

  // ExposeHeaders 定义客户端可以访问的白名单（whitelist）标头。
  // 可选。默认值 []string{}。
  ExposeHeaders []string `yaml:"expose_headers"`

  // MaxAge 指示预检（preflight）请求的结果可以被缓存多长时间（以秒为单位）。
  // 可选。默认值 0。
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