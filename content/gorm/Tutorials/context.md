+++
title = "Context"
date = 2023-10-28T14:29:31+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/context.html](https://gorm.io/docs/context.html)

GORM provides Context support, you can use it with method `WithContext`

## 单次会话模式 Single Session Mode

Single session mode usually used when you want to perform a single operation

​	单次会话模式通常用于执行单个操作，例如：

``` go
db.WithContext(ctx).Find(&users)
```

## 连续会话模式 Continuous session mode

Continuous session mode is usually used when you want to perform a group of operations, for example:

​	连续会话模式通常用于执行一组操作，例如：

``` go
tx := db.WithContext(ctx)
tx.First(&user, 1)
tx.Model(&user).Update("role", "admin")
```

## 上下文超时 Context timeout

You can pass in a context with a timeout to `db.WithContext` to set timeout for long running queries, for example:

​	您可以将带有超时的上下文传递给`db.WithContext`以设置长时间运行查询的超时时间，例如：

``` go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

db.WithContext(ctx).Find(&users)
```

## 钩子/回调中的上下文 Context in Hooks/Callbacks

You can access the `Context` object from the current `Statement`, for example:

​	您可以从当前`Statement`访问`Context`对象，例如：

``` go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  ctx := tx.Statement.Context
  // ...
  return
}
```

## Chi中间件示例 Chi Middleware Example

Continuous session mode which might be helpful when handling API requests, for example, you can set up `*gorm.DB` with Timeout Context in middlewares, and then use the `*gorm.DB` when processing all requests

​	在处理API请求时，连续会话模式可能很有用，例如，您可以在中间件中设置带有超时上下文的`*gorm.DB`，然后在处理所有请求时使用该`*gorm.DB`。

Following is a Chi middleware example:

​	以下是Chi中间件示例：

``` go
func SetDBMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    timeoutContext, _ := context.WithTimeout(context.Background(), time.Second)
    ctx := context.WithValue(r.Context(), "DB", db.WithContext(timeoutContext))
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

r := chi.NewRouter()
r.Use(SetDBMiddleware)

r.Get("/", func(w http.ResponseWriter, r *http.Request) {
  db, ok := ctx.Value("DB").(*gorm.DB)

  var users []User
  db.Find(&users)

  // lots of db operations
})

r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
  db, ok := ctx.Value("DB").(*gorm.DB)

  var user User
  db.First(&user)

  // lots of db operations
})
```

> **NOTE** Setting `Context` with `WithContext` is goroutine-safe, refer [Session](https://gorm.io/docs/session.html) for details
>
> **注意** 使用`WithContext`设置`Context`是线程安全的，请参阅[Session](https://gorm.io/docs/session.html)以获取详细信息

## 日志记录器 Logger

Logger accepts `Context` too, you can use it for log tracking, refer [Logger](https://gorm.io/docs/logger.html) for details

​	日志记录器也接受`Context`，您可以使用它进行日志跟踪，请参阅[Logger](https://gorm.io/docs/logger.html)以获取详细信息