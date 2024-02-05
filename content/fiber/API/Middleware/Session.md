+++
title = "Session"
date = 2024-02-05T09:14:15+08:00
weight = 260
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/session]({{< ref "/fiber/API/Middleware/Session" >}})

# Session

Session middleware for [Fiber](https://github.com/gofiber/fiber).

​	Fiber 的会话中间件。

NOTE
注意

This middleware uses our [Storage](https://github.com/gofiber/storage) package to support various databases through a single interface. The default configuration for this middleware saves data to memory, see the examples below for other databases.

​	此中间件使用我们的存储包通过单个接口支持各种数据库。此中间件的默认配置将数据保存到内存中，请参阅以下示例以了解其他数据库。

## Signatures 签名

```go
func New(config ...Config) *Store
func (s *Store) RegisterType(i interface{})
func (s *Store) Get(c *fiber.Ctx) (*Session, error)
func (s *Store) Delete(id string) error
func (s *Store) Reset() error

func (s *Session) Get(key string) interface{}
func (s *Session) Set(key string, val interface{})
func (s *Session) Delete(key string)
func (s *Session) Destroy() error
func (s *Session) Reset() error
func (s *Session) Regenerate() error
func (s *Session) Save() error
func (s *Session) Fresh() bool
func (s *Session) ID() string
func (s *Session) Keys() []string
func (s *Session) SetExpiry(exp time.Duration)
```



CAUTION
注意

Storing `interface{}` values are limited to built-ins Go types.

​	存储 `interface{}` 值仅限于内建的 Go 数据类型。

## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/session"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
// This stores all of your app's sessions
store := session.New()

app.Get("/", func(c *fiber.Ctx) error {
    // Get session from storage
    sess, err := store.Get(c)
    if err != nil {
        panic(err)
    }

    // Get value
    name := sess.Get("name")

    // Set key/value
    sess.Set("name", "john")

    // Get all Keys
    keys := sess.Keys()

    // Delete key
    sess.Delete("name")

    // Destroy session
    if err := sess.Destroy(); err != nil {
        panic(err)
    }

    // Sets a specific expiration for this session
    sess.SetExpiry(time.Second * 2)

    // Save session
    if err := sess.Save(); err != nil {
        panic(err)
    }

    return c.SendString(fmt.Sprintf("Welcome %v", name))
})
```



## Config 配置

| Property 属性                                                | Type 输入       | Description 说明                                             | Default 默认          |
| ------------------------------------------------------------ | --------------- | ------------------------------------------------------------ | --------------------- |
| Expiration 过期                                              | `time.Duration` | Allowed session duration. 允许的会话持续时间。               | `24 * time.Hour`      |
| Storage                                                      | `fiber.Storage` | Storage interface to store the session data. 存储会话数据的存储接口。 | `memory.New()`        |
| KeyLookup                                                    | `string`        | KeyLookup is a string in the form of "`<source>:<name>`" that is used to extract session id from the request. KeyLookup 是一个 " `<source>:<name>` " 形式的字符串，用于从请求中提取会话 ID。 | `"cookie:session_id"` |
| CookieDomain                                                 | `string`        | Domain of the cookie. Cookie 的域。                          | `""`                  |
| CookiePath                                                   | `string`        | Path of the cookie. Cookie 的路径。                          | `""`                  |
| CookieSecure                                                 | `bool`          | Indicates if cookie is secure. 指示 Cookie 是否安全。        | `false`               |
| CookieHTTPOnly 指明是否安全。                                | `bool`          | Indicates if cookie is HTTP only. 指示 Cookie 是否仅限 HTTP。 | `false`               |
| CookieSameSite Site 的值。                                   | `string`        | Value of SameSite cookie. “Lax”                              | `"Lax"`               |
| CookieSessionOnly 决定令牌是否仅持续一次浏览器的会话。如果设置为真，则忽略过期时间。 | `bool`          | Decides whether cookie should last for only the browser session. Ignores Expiration if set to true. 决定 Cookie 是否仅持续浏览器会话。如果设置为 true，则忽略过期时间。 | `false`               |
| KeyGenerator                                                 | `func() string` | KeyGenerator generates the session key. KeyGenerator 生成会话键。 | `utils.UUIDv4`        |
| CookieName (Deprecated) CookieName（已弃用）                 | `string`        | Deprecated: Please use KeyLookup. The session name. 已弃用：请使用 KeyLookup。会话名称。 | `""`                  |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Expiration:   24 * time.Hour,
    KeyLookup:    "cookie:session_id",
    KeyGenerator: utils.UUIDv4,
    source:       "cookie",
    sessionName:  "session_id",
}
```



## Constants 常量 

```go
const (
    SourceCookie   Source = "cookie"
    SourceHeader   Source = "header"
    SourceURLQuery Source = "query"
)
```



### Custom Storage/Database 自定义存储/数据库

You can use any storage from our [storage](https://github.com/gofiber/storage/) package.

​	您可以使用存储包中的任何存储。

```go
storage := sqlite3.New() // From github.com/gofiber/storage/sqlite3
store := session.New(session.Config{
    Storage: storage,
})
```



To use the store, see the [Examples](https://docs.gofiber.io/api/middleware/session/#examples).

​	要使用存储，请参阅示例。
