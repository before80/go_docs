+++
title = "CSRF"
date = 2024-02-05T09:14:15+08:00
weight = 50
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/csrf]({{< ref "/fiber/API/Middleware/CSRF" >}})

# CSRF

The CSRF middleware for [Fiber](https://github.com/gofiber/fiber) provides protection against [Cross-Site Request Forgery](https://en.wikipedia.org/wiki/Cross-site_request_forgery) (CSRF) attacks. Requests made using methods other than those defined as 'safe' by [RFC9110#section-9.2.1](https://datatracker.ietf.org/doc/html/rfc9110.html#section-9.2.1) (GET, HEAD, OPTIONS, and TRACE) are validated using tokens. If a potential attack is detected, the middleware will return a default 403 Forbidden error.

​	Fiber 的 CSRF 中间件提供针对跨站点请求伪造 (CSRF) 攻击的保护。使用 RFC9110#section-9.2.1（GET、HEAD、OPTIONS 和 TRACE）定义为“安全”以外的方法发出的请求将使用令牌进行验证。如果检测到潜在攻击，中间件将返回默认的 403 禁止错误。

This middleware offers two [Token Validation Patterns](https://docs.gofiber.io/api/middleware/csrf/#token-validation-patterns): the [Double Submit Cookie Pattern (default)](https://docs.gofiber.io/api/middleware/csrf/#double-submit-cookie-pattern-default), and the [Synchronizer Token Pattern (with Session)](https://docs.gofiber.io/api/middleware/csrf/#synchronizer-token-pattern-session).

​	此中间件提供两种令牌验证模式：双重提交 Cookie 模式（默认）和同步令牌模式（带会话）。

As a [Defense In Depth](https://docs.gofiber.io/api/middleware/csrf/#defense-in-depth) measure, this middleware performs [Referer Checking](https://docs.gofiber.io/api/middleware/csrf/#referer-checking) for HTTPS requests.

​	作为纵深防御措施，此中间件对 HTTPS 请求执行引用者检查。

## Token Generation 令牌生成

CSRF tokens are generated on 'safe' requests and when the existing token has expired or hasn't been set yet. If `SingleUseToken` is `true`, a new token is generated after each use. Retrieve the CSRF token using `c.Locals(contextKey)`, where `contextKey` is defined within the configuration.

​	CSRF 令牌在“安全”请求中生成，并且在现有令牌已过期或尚未设置时生成。如果 `SingleUseToken` 为 `true` ，则每次使用后都会生成一个新令牌。使用 `c.Locals(contextKey)` 检索 CSRF 令牌，其中 `contextKey` 在配置中定义。

## Security Considerations 安全注意事项

This middleware is designed to protect against CSRF attacks but does not protect against other attack vectors, such as XSS. It should be used in combination with other security measures.

​	此中间件旨在防御 CSRF 攻击，但不防御其他攻击媒介，例如 XSS。应将其与其他安全措施结合使用。

DANGER
危险

Never use 'safe' methods to mutate data, for example, never use a GET request to modify a resource. This middleware will not protect against CSRF attacks on 'safe' methods.

​	切勿使用“安全”方法来改变数据，例如，切勿使用 GET 请求来修改资源。此中间件不会防御针对“安全”方法的 CSRF 攻击。

### Token Validation Patterns 令牌验证模式

#### Double Submit Cookie Pattern (Default) 双重提交 Cookie 模式（默认）

By default, the middleware generates and stores tokens using the `fiber.Storage` interface. These tokens are not linked to any particular user session, and they are validated using the Double Submit Cookie pattern. The token is stored in a cookie, and then sent as a header on requests. The middleware compares the cookie value with the header value to validate the token. This is a secure pattern that does not require a user session.

​	默认情况下，中间件使用 `fiber.Storage` 接口生成并存储令牌。这些令牌不链接到任何特定用户会话，并且使用双重提交 Cookie 模式进行验证。令牌存储在 Cookie 中，然后作为请求的标头发送。中间件将 Cookie 值与标头值进行比较以验证令牌。这是一个不需要用户会话的安全模式。

When the authorization status changes, the previously issued token MUST be deleted, and a new one generated. See [Token Lifecycle](https://docs.gofiber.io/api/middleware/csrf/#token-lifecycle) [Deleting Tokens](https://docs.gofiber.io/api/middleware/csrf/#deleting-tokens) for more information.

​	当授权状态发生更改时，必须删除先前颁发的令牌，并生成新的令牌。有关更多信息，请参阅令牌生命周期删除令牌。

CAUTION
注意

When using this pattern, it's important to set the `CookieSameSite` option to `Lax` or `Strict` and ensure that the Extractor is not `CsrfFromCookie`, and KeyLookup is not `cookie:<name>`.

​	使用此模式时，重要的是将 `CookieSameSite` 选项设置为 `Lax` 或 `Strict` ，并确保 Extractor 不是 `CsrfFromCookie` ，KeyLookup 不是 `cookie:<name>` 。

NOTE
注意

When using this pattern, this middleware uses our [Storage](https://github.com/gofiber/storage) package to support various databases through a single interface. The default configuration for Storage saves data to memory. See [Custom Storage/Database](https://docs.gofiber.io/api/middleware/csrf/#custom-storagedatabase) for customizing the storage.

​	使用此模式时，此中间件使用我们的 Storage 包通过单个接口支持各种数据库。Storage 的默认配置将数据保存到内存中。有关自定义存储，请参阅自定义存储/数据库。 同步令牌模式（带会话）

#### Synchronizer Token Pattern (with Session)

When using this middleware with a user session, the middleware can be configured to store the token within the session. This method is recommended when using a user session, as it is generally more secure than the Double Submit Cookie Pattern.

​	当将此中间件与用户会话一起使用时，可以将中间件配置为将令牌存储在会话中。当使用用户会话时，建议使用此方法，因为它通常比双重提交 Cookie 模式更安全。

When using this pattern it's important to regenerate the session when the authorization status changes, this will also delete the token. See: [Token Lifecycle](https://docs.gofiber.io/api/middleware/csrf/#token-lifecycle) for more information.

​	当使用此模式时，在授权状态更改时重新生成会话非常重要，这也将删除令牌。有关更多信息，请参阅：令牌生命周期。

CAUTION
注意

Pre-sessions are required and will be created automatically if not present. Use a session value to indicate authentication instead of relying on presence of a session.

​	需要预会话，如果不存在，将自动创建。使用会话值来指示身份验证，而不是依赖会话的存在。

### Defense In Depth 纵深防御 

When using this middleware, it's recommended to serve your pages over HTTPS, set the `CookieSecure` option to `true`, and set the `CookieSameSite` option to `Lax` or `Strict`. This ensures that the cookie is only sent over HTTPS and not on requests from external sites.

​	当使用此中间件时，建议通过 HTTPS 提供页面，将 `CookieSecure` 选项设置为 `true` ，并将 `CookieSameSite` 选项设置为 `Lax` 或 `Strict` 。这可确保仅通过 HTTPS 发送 Cookie，而不通过外部网站的请求发送。

NOTE
注意

Cookie prefixes `__Host-` and `__Secure-` can be used to further secure the cookie. Note that these prefixes are not supported by all browsers and there are other limitations. See [MDN#Set-Cookie#cookie_prefixes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie#cookie_prefixes) for more information.

​	Cookie 前缀 `__Host-` 和 `__Secure-` 可用于进一步保护 Cookie。请注意，并非所有浏览器都支持这些前缀，还存在其他限制。有关更多信息，请参阅 MDN#Set-Cookie#cookie_prefixes。

To use these prefixes, set the `CookieName` option to `__Host-csrf_` or `__Secure-csrf_`.

​	要使用这些前缀，请将 `CookieName` 选项设置为 `__Host-csrf_` 或 `__Secure-csrf_` 。

### Referer Checking 引用者检查 

For HTTPS requests, this middleware performs strict referer checking. Even if a subdomain can set or modify cookies on your domain, it can't force a user to post to your application since that request won't come from your own exact domain.

​	对于 HTTPS 请求，此中间件执行严格的引用检查。即使子域可以在您的域上设置或修改 cookie，它也无法强制用户发布到您的应用程序，因为该请求不会来自您自己的确切域。

CAUTION
注意

When HTTPS requests are protected by CSRF, referer checking is always carried out.

​	当 HTTPS 请求受 CSRF 保护时，始终执行引用检查。

The Referer header is automatically included in requests by all modern browsers, including those made using the JS Fetch API. However, if you're making use of this middleware with a custom client, it's important to ensure that the client sends a valid Referer header.

​	所有现代浏览器都会自动在请求中包含 Referer 头，包括使用 JS Fetch API 发出的请求。但是，如果您将此中间件与自定义客户端一起使用，则务必确保客户端发送有效的 Referer 头。

### Token Lifecycle 令牌生命周期

Tokens are valid until they expire or until they are deleted. By default, tokens are valid for 1 hour, and each subsequent request extends the expiration by 1 hour. The token only expires if the user doesn't make a request for the duration of the expiration time.

​	令牌在过期或被删除之前有效。默认情况下，令牌有效期为 1 小时，并且每个后续请求将延长 1 小时。仅当用户在过期时间内未发出请求时，令牌才会过期。

#### Token Reuse 令牌重用

By default, tokens may be used multiple times. If you want to delete the token after it has been used, you can set the `SingleUseToken` option to `true`. This will delete the token after it has been used, and a new token will be generated on the next request.

​	默认情况下，令牌可以多次使用。如果您想在使用令牌后将其删除，可以将 `SingleUseToken` 选项设置为 `true` 。这将在使用令牌后将其删除，并在下一次请求中生成新的令牌。

INFO
信息

Using `SingleUseToken` comes with usability trade-offs and is not enabled by default. For example, it can interfere with the user experience if the user has multiple tabs open or uses the back button.

​	使用 `SingleUseToken` 会带来可用性权衡，并且默认情况下未启用。例如，如果用户打开了多个标签页或使用了后退按钮，它可能会干扰用户体验。

#### Deleting Tokens 删除令牌 

When the authorization status changes, the CSRF token MUST be deleted, and a new one generated. This can be done by calling `handler.DeleteToken(c)`.

​	当授权状态更改时，必须删除 CSRF 令牌并生成一个新的令牌。这可以通过调用 `handler.DeleteToken(c)` 来完成。

```go
if handler, ok := app.AcquireCtx(ctx).Locals(csrf.ConfigDefault.HandlerContextKey).(*CSRFHandler); ok {
    if err := handler.DeleteToken(app.AcquireCtx(ctx)); err != nil {
        // handle error
    }
}
```



TIP

If you are using this middleware with the fiber session middleware, then you can simply call `session.Destroy()`, `session.Regenerate()`, or `session.Reset()` to delete session and the token stored therein.

​	如果您将此中间件与 fiber 会话中间件一起使用，那么您可以简单地调用 `session.Destroy()` 、 `session.Regenerate()` 或 `session.Reset()` 来删除会话和其中存储的令牌。

### BREACH

It's important to note that the token is sent as a header on every request. If you include the token in a page that is vulnerable to [BREACH](https://en.wikipedia.org/wiki/BREACH), an attacker may be able to extract the token. To mitigate this, ensure your pages are served over HTTPS, disable HTTP compression, and implement rate limiting for requests.

​	请务必注意，令牌作为每个请求的标头发送。如果您将令牌包含在容易受到 BREACH 攻击的页面中，攻击者可能会提取该令牌。为了减轻这种情况，请确保您的页面通过 HTTPS 提供，禁用 HTTP 压缩，并对请求实施速率限制。

## Signatures 签名

```go
func New(config ...Config) fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework:

​	导入 Fiber Web 框架中包含的中间件包：

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/csrf"
)
```



After initializing your Fiber app, you can use the following code to initialize the middleware:

​	初始化 Fiber 应用后，您可以使用以下代码初始化中间件：

```go
// Initialize default config
app.Use(csrf.New())

// Or extend your config for customization
app.Use(csrf.New(csrf.Config{
    KeyLookup:      "header:X-Csrf-Token",
    CookieName:     "csrf_",
    CookieSameSite: "Lax",
    Expiration:     1 * time.Hour,
    KeyGenerator:   utils.UUIDv4,
}))
```



## Config 配置

| Property 属性                                                | Type 输入                                               | Description 说明                                             | Default 默认                                             |
| ------------------------------------------------------------ | ------------------------------------------------------- | ------------------------------------------------------------ | -------------------------------------------------------- |
| Next 下一步                                                  | `func(*fiber.Ctx) bool`                                 | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                                                    |
| KeyLookup                                                    | `string`                                                | KeyLookup is a string in the form of "`<source>:<key>`" that is used to create an Extractor that extracts the token from the request. Possible values: "`header:<name>`", "`query:<name>`", "`param:<name>`", "`form:<name>`", "`cookie:<name>`". Ignored if an Extractor is explicitly set. KeyLookup 是一个字符串，格式为 " `<source>:<key>` "，用于创建从请求中提取令牌的 Extractor。可能的值：" `header:<name>` "、" `query:<name>` "、" `param:<name>` "、" `form:<name>` "、" `cookie:<name>` "。如果明确设置了 Extractor，则忽略。 | "header:X-Csrf-Token"                                    |
| CookieName                                                   | `string`                                                | Name of the csrf cookie. This cookie will store the csrf key. csrf cookie 的名称。此 cookie 将存储 csrf 密钥。 | "csrf_"                                                  |
| CookieDomain                                                 | `string`                                                | Domain of the CSRF cookie. CSRF cookie 的域。                | ""                                                       |
| CookiePath                                                   | `string`                                                | Path of the CSRF cookie. CSRF cookie 的路径。                | ""                                                       |
| CookieSecure                                                 | `bool`                                                  | Indicates if the CSRF cookie is secure.                      | false                                                    |
| CookieHTTPOnly 指明是否安全。                                | `bool`                                                  | Indicates if the CSRF cookie is HTTP-only. 指明是否仅使用。  | false                                                    |
| CookieSameSite Site 的值。                                   | `string`                                                | Value of SameSite cookie. “Lax”                              | "Lax" 仅会话                                             |
| CookieSessionOnly 决定令牌是否仅持续一次浏览器的会话。如果设置为真，则忽略过期时间。 | `bool`                                                  | Decides whether the cookie should last for only the browser session. Ignores Expiration if set to true. 过期时间是令牌失效前的持续时间。 | false                                                    |
| Expiration 过期                                              | `time.Duration`                                         | Expiration is the duration before the CSRF token will expire. 1 * time.秒 | 1 * time.Hour                                            |
| SingleUseToken                                               | `bool`                                                  | SingleUseToken indicates if the CSRF token be destroyed and a new one generated on each use. (See TokenLifecycle) SingleUseToken 指示是否销毁 CSRF 令牌并在每次使用时生成一个新令牌。（请参阅 TokenLifecycle） | false                                                    |
| Storage                                                      | `fiber.Storage`                                         | Store is used to store the state of the middleware. Store 用于存储中间件的状态。 | `nil`                                                    |
| Session                                                      | `*session.Store`                                        | Session is used to store the state of the middleware. Overrides Storage if set. Session 用于存储中间件的状态。如果设置，则覆盖 Storage。 | `nil`                                                    |
| SessionKey                                                   | `string`                                                | SessionKey is the key used to store the token within the session. SessionKey 是用于在会话中存储令牌的密钥。 | "fiber.csrf.token"                                       |
| ContextKey                                                   | `inteface{}`                                            | Context key to store the generated CSRF token into the context. If left empty, the token will not be stored within the context. 上下文密钥，用于将生成的 CSRF 令牌存储到上下文中。如果留空，则不会将令牌存储在上下文中。 | ""                                                       |
| KeyGenerator                                                 | `func() string`                                         | KeyGenerator creates a new CSRF token. KeyGenerator 创建一个新的 CSRF 令牌。 | utils.UUID                                               |
| CookieExpires                                                | `time.Duration` (Deprecated) `time.Duration` （已弃用） | Deprecated: Please use Expiration. 已弃用：请使用 Expiration。 | 0                                                        |
| Cookie                                                       | `*fiber.Cookie` (Deprecated) `*fiber.Cookie` （已弃用） | Deprecated: Please use Cookie* related fields. 已弃用：请使用 Cookie* 相关字段。 | `nil`                                                    |
| TokenLookup                                                  | `string` (Deprecated) `string` （已弃用）               | Deprecated: Please use KeyLookup. 已弃用：请使用 KeyLookup。 | ""                                                       |
| ErrorHandler                                                 | `fiber.ErrorHandler`                                    | ErrorHandler is executed when an error is returned from fiber.Handler. 当从 fiber.Handler 返回错误时，将执行 ErrorHandler。 | DefaultErrorHandler                                      |
| Extractor                                                    | `func(*fiber.Ctx) (string, error)`                      | Extractor returns the CSRF token. If set, this will be used in place of an Extractor based on KeyLookup. Extractor 返回 CSRF 令牌。如果已设置，它将用于代替基于 KeyLookup 的 Extractor。 | Extractor based on KeyLookup 基于 KeyLookup 的 Extractor |
| HandlerContextKey                                            | `interface{}`                                           | HandlerContextKey is used to store the CSRF Handler into context. HandlerContextKey 用于将 CSRF Handler 存储到上下文中。 | "fiber.csrf.handler"                                     |

### Default Config 默认配置 

```go
var ConfigDefault = Config{
    KeyLookup:         "header:" + HeaderName,
    CookieName:        "csrf_",
    CookieSameSite:    "Lax",
    Expiration:        1 * time.Hour,
    KeyGenerator:      utils.UUIDv4,
    ErrorHandler:      defaultErrorHandler,
    Extractor:         CsrfFromHeader(HeaderName),
    SessionKey:        "fiber.csrf.token",
    HandlerContextKey: "fiber.csrf.handler",
}
```



### Recommended Config (with session) 推荐的配置（带会话）

It's recommended to use this middleware with [fiber/middleware/session]({{< ref "/fiber/API/Middleware/Session" >}}) to store the CSRF token within the session. This is generally more secure than the default configuration.

​	建议将此中间件与 fiber/middleware/session 一起使用，以便将 CSRF 令牌存储在会话中。这通常比默认配置更安全。

```go
var ConfigDefault = Config{
    KeyLookup:         "header:" + HeaderName,
    CookieName:        "__Host-csrf_",
    CookieSameSite:    "Lax",
    CookieSecure:      true,
    CookieSessionOnly: true,
    CookieHTTPOnly:    true,
    Expiration:        1 * time.Hour,
    KeyGenerator:      utils.UUIDv4,
    ErrorHandler:      defaultErrorHandler,
    Extractor:         CsrfFromHeader(HeaderName),
    Session:           session.Store,
    SessionKey:        "fiber.csrf.token",
    HandlerContextKey: "fiber.csrf.handler",
}
```



## Constants 常量 

```go
const (
    HeaderName = "X-Csrf-Token"
)
```



## Sentinel Errors

The CSRF middleware utilizes a set of sentinel errors to handle various scenarios and communicate errors effectively. These can be used within a [custom error handler](https://docs.gofiber.io/api/middleware/csrf/#custom-error-handler) to handle errors returned by the middleware.

​	CSRF 中间件利用一组哨兵错误来处理各种场景并有效地传达错误。这些错误可以在自定义错误处理程序中使用，以处理中间件返回的错误。

### Errors Returned to Error Handler 返回到错误处理程序的错误 

- `ErrTokenNotFound`: Indicates that the CSRF token was not found.
  `ErrTokenNotFound` ：指示未找到 CSRF 令牌。
- `ErrTokenInvalid`: Indicates that the CSRF token is invalid.
  `ErrTokenInvalid` ：指示 CSRF 令牌无效。
- `ErrNoReferer`: Indicates that the referer was not supplied.
  `ErrNoReferer` ：指示未提供引用者。
- `ErrBadReferer`: Indicates that the referer is invalid.
  `ErrBadReferer` ：指示引用者无效。

If you use the default error handler, the client will receive a 403 Forbidden error without any additional information.

​	如果您使用默认错误处理程序，则客户端将收到 403 Forbidden 错误，而没有任何其他信息。

## Custom Error Handler 自定义错误处理程序 

You can use a custom error handler to handle errors returned by the CSRF middleware. The error handler is executed when an error is returned from the middleware. The error handler is passed the error returned from the middleware and the fiber.Ctx.

​	您可以使用自定义错误处理程序来处理 CSRF 中间件返回的错误。当从中间件返回错误时，将执行错误处理程序。错误处理程序将传递从中间件返回的错误和 fiber.Ctx。

Example, returning a JSON response for API requests and rendering an error page for other requests:

​	示例，为 API 请求返回 JSON 响应并为其他请求呈现错误页面：

```go
app.Use(csrf.New(csrf.Config{
    ErrorHandler: func(c *fiber.Ctx, err error) error {
        accepts := c.Accepts("html", "json")
        path := c.Path()
        if accepts == "json" || strings.HasPrefix(path, "/api/") {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                "error": "Forbidden",
            })
        }
        return c.Status(fiber.StatusForbidden).Render("error", fiber.Map{
            "Title": "Forbidden",
            "Status": fiber.StatusForbidden,
        }, "layouts/main")
    },
}))
```



## Custom Storage/Database 自定义存储/数据库

You can use any storage from our [storage](https://github.com/gofiber/storage/) package.

​	您可以使用存储包中的任何存储。

```go
storage := sqlite3.New() // From github.com/gofiber/storage/sqlite3
app.Use(csrf.New(csrf.Config{
    Storage: storage,
}))
```
