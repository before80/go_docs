+++
title = "Sessions"
date = 2024-02-04T21:08:45+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/request_handling/sessions/]({{< ref "/buffalo/requestHandling/sessions" >}})

# Sessions 会话 

An HTTP session is a non-persistent data storage, which is destroyed on browser shutdown (in the default browser configuration). It can be used to store flash messages, or any temporary user-specific data. Use [cookies]({{< ref "/buffalo/requestHandling/cookies" >}}) instead if you need a more persistent client side storage.

​	HTTP 会话是一种非持久性数据存储，在浏览器关闭时销毁（在默认浏览器配置中）。它可用于存储闪存消息或任何临时用户特定数据。如果您需要更持久性客户端存储，请改用 Cookie。

The session is available directly from the `buffalo.Context` inside of a handler.

​	会话可直接从处理程序内部的 `buffalo.Context` 获得。

```go
func MyHandler(c buffalo.Context) error {
  s := c.Session()
}
```

## The Session Type 会话类型 

The `buffalo.Session` type has everything needed to work with a session during a request. Under the covers Buffalo uses the [github.com/gorilla/sessions](http://www.gorillatoolkit.org/pkg/sessions) package for managing the session.

​	 `buffalo.Session` 类型具有在请求期间使用会话所需的一切。在后台，Buffalo 使用 github.com/gorilla/sessions 包来管理会话。

```go
type Session
  // Clear a session of all values
  func (s *Session) Clear()
  // Delete a specific value from the session
  func (s *Session) Delete(name interface{})
  // Get a value from the session
  func (s *Session) Get(name interface{}) interface{}
  // GetOnce gets a value from the current session and then deletes it.
  func (s *Session) GetOnce(name interface{}) interface{}
  // Save a session
  func (s *Session) Save() error
  // Set a value on the session
  func (s *Session) Set(name, value interface{})
```

## Session Store 会话存储 

By default Buffalo will setup a session store using [`sessions.CookieStore`](http://www.gorillatoolkit.org/pkg/sessions#CookieStore).

​	默认情况下，Buffalo 将使用 `sessions.CookieStore` 设置会话存储。

This can be changed when setting up a new Buffalo application using the `SessionStore` option:

​	这可以在使用 `SessionStore` 选项设置新 Buffalo 应用程序时更改：

```go
app = buffalo.New(buffalo.Options{
  Env:         ENV,
  SessionName: "_coke_session",
  SessionStore: sessions.NewCookieStore([]byte("some session secret")),
})
```

The ENV variable `SESSION_SECRET` should be set before running the application. If this is not set, you will see a warning in your logs that your session is not secured.

​	在运行应用程序之前，应设置 ENV 变量 `SESSION_SECRET` 。如果未设置此变量，您将在日志中看到警告，指出您的会话不安全。

For more information on this see the docs for [`buffalo.Options`](https://godoc.org/github.com/gobuffalo/buffalo#Options).

​	有关此内容的更多信息，请参阅 `buffalo.Options` 的文档。

## Storing Complex Types 存储复杂类型 

It is generally considered **not** good practice to store complex types in a session. There are lots of reasons for this, but it is recommended to store the ID of a type, instead of the “whole” value.

​	通常认为在会话中存储复杂类型不是一个好习惯。原因有很多，但建议存储类型的 ID，而不是“整个”值。

Should you need to store a complex type, like a `struct` you will first need to register the type with the [`encoding/gob`](https://golang.org/pkg/encoding/gob/) package.

​	如果您需要存储复杂类型，例如 `struct` ，则首先需要使用 `encoding/gob` 包注册该类型。

```go
import "encoding/gob"

func init() {
  gob.Register(&models.Person{})
}
```

## Saving a Session 保存会话 

Buffalo automatically saves your session for you, so you don’t have to. If there is an error when saving the session, Buffalo will return an error through the normal [error handling ]({{< ref "/buffalo/requestHandling/errorHanding" >}})process.

​	Buffalo 会自动为您保存会话，因此您不必这样做。如果在保存会话时出错，Buffalo 将通过正常的错误处理过程返回错误。

## Null Sessions for APIs API 的空会话 

When building API servers the default cookie session store is undesirable. The `sessions.Null` type is the recommended replacement for the default session store.

​	在构建 API 服务器时，默认的 cookie 会话存储是不可取的。 `sessions.Null` 类型是默认会话存储的推荐替代类型。

```go
app = buffalo.New(buffalo.Options{
  Env:          ENV,
  SessionStore: sessions.Null{},
  SessionName: "_coke_session",
})
```

When running `buffalo new` with the `--api` flag the default session will be set to `sessions.Null`.

​	使用 `--api` 标志运行 `buffalo new` 时，默认会话将设置为 `sessions.Null` 。
