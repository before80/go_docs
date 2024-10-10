+++
title = "jwt"
weight = 100
date = 2023-07-09T21:55:16+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# JWT

> 原文：[https://echo.labstack.com/docs/middleware/jwt](https://echo.labstack.com/docs/middleware/jwt)



​	JWT提供了JSON Web Token（JWT）身份验证中间件。Echo JWT中间件位于[https://github.com/labstack/echo-jwt](https://github.com/labstack/echo-jwt)



​	基本的中间件行为如下：

- 对于有效的令牌，它在上下文中设置用户并调用下一个处理程序。
- 对于无效的令牌，它发送"401 - Unauthorized"响应。
- 对于缺失或无效的`Authorization`头，它发送 "400 - Bad Request"。

## 依赖

```go
import "github.com/labstack/echo-jwt/v4"
```



## Usage

```go
e.Use(echojwt.JWT([]byte("secret")))
```



## Custom Configuration

### Usage

```go
e.Use(echojwt.WithConfig(echojwt.Config{
  // ...
  SigningKey:             []byte("secret"),
  // ...
}))
```



## Configuration

```go
type Config struct {
    // Skipper定义一个用于跳过中间件的函数。
    Skipper middleware.Skipper

    // BeforeFunc定义一个在中间件之前执行的函数。
    BeforeFunc middleware.BeforeFunc

    // SuccessHandler定义一个在令牌有效时执行的函数。
    SuccessHandler func(c echo.Context)

    // ErrorHandler定义当所有查找都完成并且没有一个通过验证器函数时执行的函数。
    // ErrorHandler在执行时会传入最后一个缺失的（ErrExtractionValueMissing）或一个无效的键。
    // 它可用于定义自定义的JWT错误。
    //
    // 注意：当错误处理程序忽略错误（返回`nil`）时，
    // 中间件会继续执行处理程序链以继续处理。
    // 这在您的站点/ API的某些部分可公开访问并且为授权用户提供额外功能的情况下非常有用。
    // 在这种情况下，
    // 您可以使用ErrorHandler将默认的公共JWT令牌值设置到请求中并继续处理程序链。
    ErrorHandler func(c echo.Context, err error) error

    // ContinueOnIgnoredError允许在ErrorHandler决定忽略错误时
    // （通过返回`nil`）调用下一个中间件/处理程序。
    // 这在您的站点/ API的某些部分允许公共访问并且某些授权路由提供额外功能的情况下非常有用。
    // 在这种情况下，
    // 您可以使用ErrorHandler在请求上下文中设置默认的公共JWT令牌值并继续执行。
    // 然后，剩余的执行链中的一些逻辑需要检查（公共）令牌值。
    ContinueOnIgnoredError bool

    // ContextKey用于将令牌中的用户信息存储到上下文中。
    // 可选。默认值为"user"。
    ContextKey string

    // SigningKey用于验证令牌的签名密钥。
    // 这是提供令牌验证密钥的三个选项之一。
    // 优先级顺序为：用户定义的KeyFunc、SigningKeys和SigningKey。
    // 如果未提供用户定义的KeyFunc和SigningKeys，则为必需。
    SigningKey interface{}

    // SigningKeys是一组用于根据kid字段使用的验证令牌的签名密钥。
    // 这是提供令牌验证密钥的三个选项之一。
    // 优先级顺序为：用户定义的KeyFunc、SigningKeys和SigningKey。
    // 如果未提供用户定义的KeyFunc和SigningKey，则为必需。
    SigningKeys map[string]interface{}

    // SigningMethod用于检查令牌的签名算法。
    // 可选。默认值为HS256。
    SigningMethod string

    // KeyFunc定义一个用户定义的函数，用于提供令牌验证的公钥。
    // 该函数应负责验证签名算法并选择正确的密钥。
    // 如果令牌由外部方发行，则用户定义的KeyFunc可能很有用。
    // 默认情况下使用ParseTokenFunc的实现。
    //
    // 当提供了用户定义的KeyFunc时，SigningKey、SigningKeys和SigningMethod将被忽略。
    // 这是提供令牌验证密钥的三个选项之一。
    // 优先级顺序为：用户定义的KeyFunc、SigningKeys和SigningKey。
    // 如果未提供SigningKeys和SigningKey，则为必需。
    // 如果设置了自定义的ParseTokenFunc，则不使用KeyFunc。
    // 默认情况下，KeyFunc采用内部实现，用于验证签名算法并选择适当的密钥。
    KeyFunc jwt.Keyfunc 
    
    // TokenLookup 是一个字符串，格式为 "<source>:<name>" 或 "<source>:<name>,<source>:<name>"，用于从请求中提取令牌。
	// 可选。默认值为 "header:Authorization"。
	// 可能的值：
	// - "header:<name>" 或 "header:<name>:<cut-prefix>"
	//          `<cut-prefix>` 是要剪切/修整提取值前缀的参数值。
	//          如果标头值具有固定前缀，例如 `Authorization: <auth-scheme> <authorisation-parameters>`，
    //          我们要剪切的部分是 `<auth-scheme> `，请注意末尾的空格。
	//          对于 JWT 令牌的情况，我们要剪切的前缀是 `Bearer `。
	//          如果前缀留空，则返回整个值。
    // - "query:<name>"
    // - "param:<name>"
    // - "cookie:<name>"
    // - "form:<name>"
    // 多个来源的示例：
    // - "header:Authorization:Bearer ,cookie:myowncookie"
    TokenLookup string

    // TokenLookupFuncs 定义了一组用户定义的函数，用于从给定的上下文中提取 JWT 令牌。
    // 这是提供令牌提取器的两个选项之一。
	// 优先级顺序为用户定义的 TokenLookupFuncs 和 TokenLookup。
	// 如果需要，您也可以同时提供两者。
    TokenLookupFuncs []middleware.ValuesExtractor
    
    // ParseTokenFunc 定义了一个用户定义的函数，该函数从给定的 auth 中解析令牌。
    // 在令牌解析失败或解析的令牌无效时返回错误。
	// 默认情况下，使用 `github.com/golang-jwt/jwt` 作为 JWT 实现库。
    ParseTokenFunc func(c echo.Context, auth string) (interface{}, error)

    // Claims 是定义令牌内容的可扩展声明数据。被默认的 ParseTokenFunc 实现使用。
	// 如果设置了自定义的 ParseTokenFunc，则不使用 Claims。
	// 可选。默认为返回 jwt.MapClaims 的函数。
    NewClaimsFunc func(c echo.Context) jwt.Claims
}
```



