+++
title = "routing"
date = 2024-01-31T19:05:12+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-chi.io/#/pages/routing](https://go-chi.io/#/pages/routing)

# 🔌 Routing 🔌 路由

## Introduction 简介

> Routing refers to how an application's endpoints (URIs) respond to client requests.
>
> ​	路由是指应用程序的端点 (URI) 如何响应客户端请求。

`Chi` allows you to route/handle any HTTP request method, such as all the usual suspects: GET, POST, HEAD, PUT, PATCH, DELETE, OPTIONS, TRACE, CONNECT

​	`Chi` 允许您路由/处理任何 HTTP 请求方法，例如所有常见的嫌疑犯：GET、POST、HEAD、PUT、PATCH、DELETE、OPTIONS、TRACE、CONNECT

## Handling HTTP Request Methods 处理 HTTP 请求方法

These methods are defined on the `chi.Router` as:

​	这些方法在 `chi.Router` 上定义为：

```go
// HTTP-method routing along `pattern`
Connect(pattern string, h http.HandlerFunc)
Delete(pattern string, h http.HandlerFunc)
Get(pattern string, h http.HandlerFunc)
Head(pattern string, h http.HandlerFunc)
Options(pattern string, h http.HandlerFunc)
Patch(pattern string, h http.HandlerFunc)
Post(pattern string, h http.HandlerFunc)
Put(pattern string, h http.HandlerFunc)
Trace(pattern string, h http.HandlerFunc)Copy to clipboardErrorCopied
```

and may set a route by calling ie. `r.Put("/path", myHandler)`.

​	并可以通过调用即 `r.Put("/path", myHandler)` 设置路由。

You may also register your own custom method names, by calling `chi.RegisterMethod("JELLO")` and then setting the routing handler via `r.Method("JELLO", "/path", myJelloMethodHandler)`

​	您还可以通过调用 `chi.RegisterMethod("JELLO")` 并通过 `r.Method("JELLO", "/path", myJelloMethodHandler)` 设置路由处理程序来注册您自己的自定义方法名称

## Routing patterns & url parameters 路由模式和 URL 参数

Each routing method accepts a URL `pattern` and chain of `handlers`.

​	每个路由方法接受一个 URL `pattern` 和 `handlers` 链。

The URL pattern supports named params (ie. `/users/{userID}`) and wildcards (ie. `/admin/*`).

​	URL 模式支持命名参数（即 `/users/{userID}` ）和通配符（即 `/admin/*` ）。

URL parameters can be fetched at runtime by calling `chi.URLParam(r, "userID")` for named parameters and `chi.URLParam(r, "*")` for a wildcard parameter.

​	可以通过调用 `chi.URLParam(r, "userID")` 获取命名参数的 URL 参数，或调用 `chi.URLParam(r, "*")` 获取通配符参数的 URL 参数。

**Routing a slug:
路由 slug：**

```go
r := chi.NewRouter()

r.Get("/articles/{date}-{slug}", getArticle)

func getArticle(w http.ResponseWriter, r *http.Request) {
  dateParam := chi.URLParam(r, "date")
  slugParam := chi.URLParam(r, "slug")
  article, err := database.GetArticle(date, slug)

  if err != nil {
    w.WriteHeader(422)
    w.Write([]byte(fmt.Sprintf("error fetching article %s-%s: %v", dateParam, slugParam, err)))
    return
  }
  
  if article == nil {
    w.WriteHeader(404)
    w.Write([]byte("article not found"))
    return
  }
  w.Write([]byte(article.Text()))
})Copy to clipboardErrorCopied
```

as you can see above, the url parameters are defined using the curly brackets `{}` with the parameter name in between, as `{date}` and `{slug}`.

​	如您在上面看到的，URL 参数使用带有参数名称的卷曲括号 `{}` 定义，如 `{date}` 和 `{slug}` 。

When a HTTP request is sent to the server and handled by the chi router, if the URL path matches the format of `/articles/{date}-{slug}`, then the `getArticle` function will be called to send a response to the client.

​	当 HTTP 请求发送到服务器并由 chi 路由器处理时，如果 URL 路径与 `/articles/{date}-{slug}` 的格式匹配，则将调用 `getArticle` 函数向客户端发送响应。

For instance, URL paths like `/articles/20200109-this-is-so-cool` will match the route, however, `/articles/1` will not.

​	例如，URL 路径（如 `/articles/20200109-this-is-so-cool` ）将匹配路由，但 `/articles/1` 不会。

We can also use regex in url patterns

​	我们还可以在 URL 模式中使用正则表达式

For Example:

​	例如：

```go
r := chi.NewRouter()
r.Get("/articles/{rid:^[0-9]{5,6}}", getArticle)Copy to clipboardErrorCopied
```

## Making Custom 404 and 405 Handlers 制作自定义 404 和 405 处理程序

You can create Custom `http.StatusNotFound` and `http.StatusMethodNotAllowed` handlers in `chi`

​	您可以在 `chi` 中创建自定义 `http.StatusNotFound` 和 `http.StatusMethodNotAllowed` 处理程序

```go
r.NotFound(func(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(404)
  w.Write([]byte("route does not exist"))
})
r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(405)
  w.Write([]byte("method is not valid"))
})Copy to clipboardErrorCopied
```

## Sub Routers 子路由

You can create New Routers and Mount them on the Main Router to act as Sub Routers.

​	您可以创建新路由并将其安装在主路由上，使其充当子路由。

For Example:

​	例如：

```go
func main(){
    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    })

    // Creating a New Router
    apiRouter := chi.NewRouter()
    apiRouter.Get("/articles/{date}-{slug}", getArticle)

    // Mounting the new Sub Router on the main router
    r.Mount("/api", apiRouter)
}Copy to clipboardErrorCopied
```

Another Way of Implementing Sub Routers would be:

​	实现子路由的另一种方法是：

```go
r.Route("/articles", func(r chi.Router) {
    r.With(paginate).Get("/", listArticles)                           // GET /articles
    r.With(paginate).Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017

    r.Post("/", createArticle)                                        // POST /articles
    r.Get("/search", searchArticles)                                  // GET /articles/search

    // Regexp url parameters:
    r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug)                // GET /articles/home-is-toronto

    // Subrouters:
    r.Route("/{articleID}", func(r chi.Router) {
      r.Use(ArticleCtx)
      r.Get("/", getArticle)                                          // GET /articles/123
      r.Put("/", updateArticle)                                       // PUT /articles/123
      r.Delete("/", deleteArticle)                                    // DELETE /articles/123
    })
  })Copy to clipboardErrorCopied
```

## Routing Groups 路由组

You can create Groups in Routers to segregate routes using a middleware and some not using a middleware

​	您可以在路由器中创建组，以使用中间件隔离路由，而有些则不使用中间件

for example:

​	例如：

```go
func main(){
    r := chi.NewRouter()
    
    // Public Routes
    r.Group(func(r chi.Router) {
        r.Get("/", HelloWorld)
        r.Get("/{AssetUrl}", GetAsset)
        r.Get("/manage/url/{path}", FetchAssetDetailsByURL)
        r.Get("/manage/id/{path}", FetchAssetDetailsByID)
    })

    // Private Routes
    // Require Authentication
    r.Group(func(r chi.Router) {
        r.Use(AuthMiddleware)
        r.Post("/manage", CreateAsset)
    })

}
```