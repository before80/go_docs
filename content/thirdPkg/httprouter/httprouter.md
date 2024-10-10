+++
title = "httprouter文档"
date = 2023-06-05T09:12:08+08:00
type = "docs"
weight = 2
description = ""
isCJKLanguage = true
draft = false

+++

# HttpRouter

> 原文：[https://pkg.go.dev/github.com/julienschmidt/httprouter](https://pkg.go.dev/github.com/julienschmidt/httprouter)
>
> 版本：v1.3.0
>
> 发布日期：2019.9.29
>
> github网址：[https://github.com/julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)

​	HttpRouter 是一个轻量级高性能的 Go 语言 HTTP 请求路由器（也称为多路复用器*multiplexer*或简称 *mux*）。

​	与 Go 标准库中的 `net/http` 包中的[默认 mux](https://golang.org/pkg/net/http/#ServeMux) 不同，该路由器支持路由模式中的变量，并匹配请求方法。它也具有更好的可扩展性。

​	该路由器针对高性能和小内存占用进行了优化。即使在非常长的路径和大量路由的情况下，它也能良好地扩展。使用压缩动态 trie（基数树  **radix tree**）结构进行高效匹配。

## 特性

​	只有显式匹配：在其他路由器（如 [http.ServeMux](https://golang.org/pkg/net/http/#ServeMux)）中，请求的 URL 路径可以匹配多个模式。因此它们有一些尴尬的模式优先级规则，比如最长匹配或首先注册，首先匹配。通过这个路由器的设计，一个请求只能精确地匹配一个或零个路由。结果，也没有意外的匹配，这使得它非常适合 SEO 并改善了用户体验。

​	不再关心尾部斜杠：选择您喜欢的URL样式，如果缺少尾部斜杠或多了一个，路由器会自动重定向客户端。当然，只有当新路径有处理程序时，它才会这样做。如果您不喜欢，可以[关闭此行为](https://godoc.org/github.com/julienschmidt/httprouter#Router.RedirectTrailingSlash)。

​	路径自动纠正：除了在不增加额外成本的情况下检测缺少或附加的尾随斜杠外，路由器还可以修复错误的大小写和删除多余的路径元素（如 `../` 或 `//`）。你的用户中有 [CAPTAIN CAPS LOCK](http://www.urbandictionary.com/define.php?term=Captain+Caps+Lock) 吗？HttpRouter 可以通过进行大小写不敏感的查找并将其重定向到正确的 URL 来帮助他。

​	在路由模式中使用参数：停止解析请求的 URL 路径，只需给路径段一个名称，路由器就会将动态值传递给您。由于路由器的设计，路径参数非常便宜。

​	零垃圾：匹配和调度过程不会产生任何垃圾。仅进行堆分配以构建路径参数的键值对切片和构建新的上下文和请求对象（仅在标准的 `Handler/HandlerFunc` API 中）。在三参数 API 中，如果请求路径不包含参数，则不需要进行任何堆分配。

​	最佳性能：[基准测试说明一切](https://github.com/julienschmidt/go-http-routing-benchmark)。有关实现的技术细节，请参见下面。

​	不再有服务器崩溃的问题：您可以设置一个[Panic处理程序](https://godoc.org/github.com/julienschmidt/httprouter#Router.PanicHandler)来处理在处理HTTP请求时发生的panic。路由器会恢复并让`PanicHandler`记录发生的情况并提供一个漂亮的错误页面。

​	完美适用于API：该路由器设计鼓励构建合理的分层RESTful API。此外，它还具有内置的[OPTIONS请求](http://zacstewart.com/2012/04/14/http-options-method.html)和`405 Method Not Allowed`响应的本机支持。

​	当然，您也可以设置自定义的[NotFound](https://godoc.org/github.com/julienschmidt/httprouter#Router.NotFound)和[MethodNotAllowed](https://godoc.org/github.com/julienschmidt/httprouter#Router.MethodNotAllowed)处理程序并提供静态文件。

## 用法

​	这只是一个简单的介绍，请查看[文档](http://pkg.go.dev/github.com/julienschmidt/httprouter)以获取详细信息。

​	让我们从一个微不足道的例子开始：

```go linenums="1"
package main

import (
    "fmt"
    "net/http"
    "log"

    "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}
```

### 命名参数 

​	正如您所看到的，`:name`是一个命名参数。值可以通过`httprouter.Params`访问，它只是`httprouter.Params`的切片。您可以通过在切片中使用其索引或使用`ByName(name)`方法来获取参数的值：`:name`可以通过`ByName("name")`检索。

​	当使用`http.Handler`（使用`router.Handler`或`http.HandlerFunc`）而不是使用第三个函数参数的HttpRouter的处理API时，命名参数存储在`request.Context`中。在下面的"为什么它不能与`http.Handler`一起工作？"下面查看更多信息。

​	命名参数仅匹配单个路径段：

```
Pattern: /user/:user

 /user/gordon              match
 /user/you                 match
 /user/gordon/profile      no match
 /user/                    no match
```

注意：由于此路由器仅具有显式匹配，因此您无法在相同的路径段上同时为相同的请求方法注册静态路由和参数。例如，您不能同时为`/user/new`和`/user/:user`注册模式，以进行相同的请求方法。不同请求方法的路由与彼此独立。

### 捕获所有参数 

​	第二种类型是捕获所有参数，格式为`*name`。正如名称所示，它们匹配所有内容。因此，它们必须始终位于模式的末尾：

```
Pattern: /src/*filepath

 /src/                     match
 /src/somefile.go          match
 /src/subdir/somefile.go   match
```

## 它是如何工作的？ 

​	路由器依赖于树形结构，该结构大量使用公共前缀，基本上是一棵紧凑的前缀树（或者称作[Radix树](https://en.wikipedia.org/wiki/Radix_tree)）。具有公共前缀的节点也共享一个共同的父节点。以下是`GET`请求方法的路由树的简短示例：

```
Priority   Path             Handle
9          \                *<1>
3          ├s               nil
2          |├earch\         *<2>
1          |└upport\        *<3>
2          ├blog\           *<4>
1          |    └:post      nil
1          |         └\     *<5>
2          ├about-us\       *<6>
1          |        └team\  *<7>
1          └contact\        *<8>
```

​	每个 `*<num>` 表示一个处理程序函数（指针）的内存地址。如果您从根节点到叶子节点沿着树的路径走，您将获得完整的路由路径，例如 `\blog:post\`，其中 `:post` 只是一个占位符（参数），用于实际的帖子名称。与哈希映射不同，树形结构也允许我们使用 `:post` 参数等动态部分，因为我们实际上是针对路由模式进行匹配，而不仅仅是比较哈希值。正如[基准测试](https://github.com/julienschmidt/go-http-routing-benchmark)所显示的那样，这非常有效和高效。

​	由于 URL 路径具有层次结构并且只使用有限的字符集（字节值），很可能存在许多公共前缀。这使我们可以轻松地将路由缩小为越来越小的问题。此外，路由器为每个请求方法管理一个单独的树。一方面，它比在每个单独节点中保存方法->处理程序映射更节省空间，另一方面它还允许我们在甚至开始查找前缀树之前大大减少路由问题。

​	为了更好的可扩展性，每个树级别上的子节点按优先级排序，其中优先级仅为在子节点（孩子、孙子等）中注册的处理程序数量。这种方法有两个帮助：

1. 首先，先评估最多的路由路径节点，这有助于使尽可能多的路由尽快到达。 
2. 这是某种成本补偿。最长可达路径（最高成本）总是可以首先进行评估。下面的图表显示了树形结构。节点从上到下，从左到右进行评估。 

```
├------------
├---------
├-----
├----
├--
├--
└-
```

## 为什么它不能与`http.Handler`一起工作？ 

​	它可以！路由器本身实现了`http.Handler`接口。此外，路由器提供了方便的适配器，使得当注册路由时，可以将它们用作[httprouter.Handle](https://godoc.org/github.com/julienschmidt/httprouter#Router.Handle)来使用[http.Handlers](https://godoc.org/github.com/julienschmidt/httprouter#Router.Handler)和[http.HandlerFuncs](https://godoc.org/github.com/julienschmidt/httprouter#Router.HandlerFunc)。

​	命名参数可以通过`request.Context`访问：

```go linenums="1"
func Hello(w http.ResponseWriter, r *http.Request) {
    params := httprouter.ParamsFromContext(r.Context())

    fmt.Fprintf(w, "hello, %s!\n", params.ByName("name"))
}
```

​	或者，也可以使用`params := r.Context().Value(httprouter.ParamsKey)`而不是使用helper函数。

​	尝试自己使用它，使用HttpRouter非常简单。这个包是紧凑而简约的，但也可能是最容易设置的路由器之一。

## 自动OPTIONS响应和CORS 

​	有人可能希望修改自动响应OPTIONS请求的方式，例如支持[CORS预检请求](https://developer.mozilla.org/en-US/docs/Glossary/preflight_request)或设置其他标头。可以使用[Router.GlobalOPTIONS](https://godoc.org/github.com/julienschmidt/httprouter#Router.GlobalOPTIONS)处理程序来实现：

```go linenums="1"
router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Access-Control-Request-Method") != "" {
        // 设置CORS标头
        header := w.Header()
        header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
        header.Set("Access-Control-Allow-Origin", "*")
    }

    // 调整状态码为204
    w.WriteHeader(http.StatusNoContent)
})
```

## 我在哪里可以找到中间件X？ 

​	此包只提供一个非常高效的请求路由器和一些额外的功能。该路由器只是一个[http.Handler](https://golang.org/pkg/net/http/#Handler)，您可以在路由器之前链接任何http.Handler兼容的中间件，例如[Gorilla处理程序](http://www.gorillatoolkit.org/pkg/handlers)。或者你可以[自己编写](https://justinas.org/writing-http-middleware-in-go/)，这非常容易！

​	或者，您可以尝试基于HttpRouter的Web框架。

### 多域/子域 

​	以下是一个快速的例子：您的服务器是否提供多个域名/主机？您想使用子域吗？为每个主机定义一个路由器！

```go linenums="1"
// 我们需要一个实现http.Handler接口的对象。
// 因此，我们需要一种类型，在其中实现ServeHTTP方法。
// 我们只是在此处使用一个映射，其中我们将主机名（带端口）
// 映射到http.Handlers
type HostSwitch map[string]http.Handler

//在我们的新类型上实现ServeHTTP方法
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//检查是否为给定主机注册了http.Handler。
	//如果是，使用它来处理请求。
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		//处理未注册处理程序的主机名
		http.Error(w, "Forbidden", 403) //还是重定向？
	}
}

func main() {
	//像往常一样初始化路由器
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	//创建一个HostSwitch并插入路由器（我们的http处理程序）
	//例如.com和端口12345
	hs := make(HostSwitch)
	hs["example.com:12345"] = router

	//使用HostSwitch在端口12345上进行监听和服务
	log.Fatal(http.ListenAndServe(":12345", hs))
}
```

### 基本身份验证 

​	另一个快速的例子：处理基本身份验证（RFC 2617）：

```go linenums="1"
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func BasicAuth(h httprouter.Handle, requiredUser, requiredPassword string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
        // 获取基本身份验证凭据
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == requiredUser && password == requiredPassword {
			// 将请求委托给给定的处理程序
			h(w, r, ps)
		} else {
			// 否则请求基本身份验证
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Not protected!\n")
}

func Protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Protected!\n")
}

func main() {
	user := "gordon"
	pass := "secret!"

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/protected/", BasicAuth(Protected, user, pass))

	log.Fatal(http.ListenAndServe(":8080", router))
}
```

## 使用NotFound处理程序进行链接

注意：可能需要将[Router.HandleMethodNotAllowed](https://godoc.org/github.com/julienschmidt/httprouter#Router.HandleMethodNotAllowed)设置为`false`以避免出现问题。

​	您可以使用另一个[http.Handler](https://golang.org/pkg/net/http/#Handler)（例如另一个路由器）来处理无法由此路由器匹配的请求，方法是使用[Router.NotFound](https://godoc.org/github.com/julienschmidt/httprouter#Router.NotFound)处理程序进行链接。这样可以进行链接。

### 静态文件

​	`NotFound`处理程序可以例如用于从根路径`/`（如`index.html`文件以及其他资源）提供静态文件：

```
// Serve static files from the ./public directory
// 从./public目录提供静态文件
router.NotFound = http.FileServer(http.Dir("public"))
```

​	但是，此方法规避了此路由器的严格核心规则以避免路由问题。更干净的方法是使用一个不同的子路径来提供文件，例如`/static/*filepath`或`/files/*filepath`。

## 基于HttpRouter的Web框架 

​	如果HttpRouter对您来说有点过于简单，您可以尝试以下更高级的第三方Web框架，这些框架是基于HttpRouter包构建的： 

- [Ace](https://github.com/plimble/ace)：极速的 Go Web 框架
- [api2go](https://github.com/manyminds/api2go)：用于 Go 的 JSON API 实现
- [Gin](https://github.com/gin-gonic/gin)：具有类似 Martini 的 API 和更好的性能
- [Goat](https://github.com/bahlo/goat)：Go 中极简的 REST API 服务器
- [goMiddlewareChain](https://github.com/TobiEiss/goMiddlewareChain)：类似于 Express.js 的中间件链
- [Hikaru](https://github.com/najeira/hikaru)：支持独立和 Google AppEngine
- [Hitch](https://github.com/nbio/hitch)：将 httprouter、[httpcontext](https://github.com/nbio/httpcontext)和中间件捆绑在一起
- [httpway](https://github.com/corneldamian/httpway)：支持具有上下文的中间件扩展的 httprouter 和具有优雅关闭支持的服务器
- [kami](https://github.com/guregu/kami)：使用 x/net/context 的微型 Web 框架
- [Medeina](https://github.com/imdario/medeina)：受 Ruby 的 Roda 和 Cuba 启发
- [Neko](https://github.com/rocwong/neko)：用于 Golang 的轻量级 Web 应用框架
- [pbgo](https://github.com/chai2010/pbgo)：基于 Protobuf 的 mini RPC/REST 框架
- [River](https://github.com/abiosoft/river)：River 是一个简单轻量级的 REST 服务器
- [siesta](https://github.com/VividCortex/siesta)：具有上下文的可组合 HTTP 处理程序
- [xmux](https://github.com/rs/xmux)：xmux 是 httprouter 在 xhandler（net/context 意识）之上的 fork



## Documentation 

### 概述

Package httprouter is a trie based high performance HTTP request router.

​	httprouter 包是一个基于字典树的高性能 HTTP 请求路由器。

A trivial example is:

​	一个简单的例子是：

``` go
package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}
```

The router matches incoming requests by the request method and the path. If a handle is registered for this path and method, the router delegates the request to that function. For the methods GET, POST, PUT, PATCH and DELETE shortcut functions exist to register handles, for all other methods router.Handle can be used.

​	路由器通过请求方法和路径匹配传入的请求。如果为该路径和方法注册了一个处理程序，路由器会将请求委托给该函数。对于 `GET`、`POST`、`PUT`、`PATCH` 和 `DELETE` 方法，存在用于注册处理程序的快捷函数，对于所有其他方法，可以使用 `router.Handle`。

The registered path, against which the router matches incoming requests, can contain two types of parameters:

​	路由器用来匹配传入请求的已注册路径可以包含两种类型的参数：

```
Syntax    Type
:name     named parameter
*name     catch-all parameter
```

Named parameters are dynamic path segments. They match anything until the next '/' or the path end:

​	命名参数是动态路径段。它们匹配到下一个“/”或路径结束之前的所有内容：

```
Path: /blog/:category/:post

Requests:
 /blog/go/request-routers            match: category="go", post="request-routers"
 /blog/go/request-routers/           no match, but the router would redirect
 /blog/go/                           no match
 /blog/go/request-routers/comments   no match
```

Catch-all parameters match anything until the path end, including the directory index (the '/' before the catch-all). Since they match anything until the end, catch-all parameters must always be the final path element.

​	万能参数匹配到路径结束之前的所有内容，包括目录索引（万能参数之前的“/”）。由于它们匹配到结束之前的所有内容，因此万能参数必须始终是路径的最后一个元素。

```
Path: /files/*filepath

Requests:
 /files/                             match: filepath="/"
 /files/LICENSE                      match: filepath="/LICENSE"
 /files/templates/article.html       match: filepath="/templates/article.html"
 /files                              no match, but the router would redirect
```

The value of parameters is saved as a slice of the Param struct, consisting each of a key and a value. The slice is passed to the Handle func as a third parameter. There are two ways to retrieve the value of a parameter:

​	参数的值保存在 Param 结构的切片中，每个参数都由一个键和一个值组成。该切片作为第三个参数传递给 Handle 函数。有两种方法可以检索参数的值：

```
// by the name of the parameter
user := ps.ByName("user") // defined by :user or *user

// by the index of the parameter. This way you can also get the name (key)
thirdKey   := ps[2].Key   // the name of the 3rd parameter
thirdValue := ps[2].Value // the value of the 3rd parameter
```

### 常量

This section is empty.

### 变量

[View Source](https://github.com/julienschmidt/httprouter/blob/v1.3.0/router.go#L115)

```
var ParamsKey = paramsKey{}
```

ParamsKey is the request context key under which URL params are stored.

​	ParamsKey 是存储 URL 参数的请求上下文键。

### 函数

#### func CleanPath 

``` go
func CleanPath(p string) string
```

CleanPath is the URL version of path.Clean, it returns a canonical URL path for p, eliminating . and .. elements.

​	`CleanPath` 是 `path.Clean` 的 URL 版本，它返回 `p` 的规范 URL 路径，消除了 `.` 和 `..` 元素。

The following rules are applied iteratively until no further processing can be done:

​	以下规则会反复应用，直到无法进行进一步处理：

1. Replace multiple slashes with a single slash.
2. 将多个斜杠替换为单个斜杠。
3. Eliminate each . path name element (the current directory).
4. 消除每个 . 路径名元素（当前目录）。
5. Eliminate each inner .. path name element (the parent directory) along with the non-.. element that precedes it.
6. 消除每个内部 `..` 路径名元素（父目录）以及紧靠其前面的非 `..` 元素。
7. Eliminate .. elements that begin a rooted path: that is, replace "/.." by "/" at the beginning of a path.
8. 消除开始根路径的 `..` 元素`：`即，在路径开头将 “`/..`” 替换为 “`/`”。

If the result of this process is an empty string, "/" is returned

​	如果此过程的结果为空字符串，则返回“/”

### 类型

#### type Handle 

``` go
type Handle func(http.ResponseWriter, *http.Request, Params)
```

Handle is a function that can be registered to a route to handle HTTP requests. Like http.HandlerFunc, but has a third parameter for the values of wildcards (variables).

​	Handle 是可以注册到路由以处理 HTTP 请求的函数。与 `http.HandlerFunc` 类似，但具有第三个参数，用于通配符（变量）的值。

#### type Param  <- 1.1.0

``` go
type Param struct {
	Key   string
	Value string
}
```

Param is a single URL parameter, consisting of a key and a value.

​	`Param` 是单个 URL 参数，由键和值组成。

#### type Params  <- 1.1.0

``` go
type Params []Param
```

Params is a Param-slice, as returned by the router. The slice is ordered, the first URL parameter is also the first slice value. It is therefore safe to read values by the index.

​	`Params` 是 `Param` 切片，由路由器返回。该切片是有序的，第一个 URL 参数也是第一个切片值。因此，按索引读取值是安全的。

#### func ParamsFromContext  <- 1.2.0

``` go
func ParamsFromContext(ctx context.Context) Params
```

ParamsFromContext pulls the URL parameters from a request context, or returns nil if none are present.

​	`ParamsFromContext` 从请求上下文中提取 URL 参数，如果不存在，则返回 nil。

#### (Params) ByName  <- 1.1.0

``` go
func (ps Params) ByName(name string) string
```

ByName returns the value of the first Param which key matches the given name. If no matching Param is found, an empty string is returned.

​	`ByName` 返回键与给定名称匹配的第一个 Param 的值。如果找不到匹配的 Param，则返回一个空字符串。

#### type Router 

``` go
type Router struct {

	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	// For example if /foo/ is requested but a route only exists for /foo, the
	// client is redirected to /foo with http status code 301 for GET requests
	// and 307 for all other request methods.
	RedirectTrailingSlash bool

	// If enabled, the router tries to fix the current request path, if no
	// handle is registered for it.
	// First superfluous path elements like ../ or // are removed.
	// Afterwards the router does a case-insensitive lookup of the cleaned path.
	// If a handle can be found for this route, the router makes a redirection
	// to the corrected path with status code 301 for GET requests and 307 for
	// all other request methods.
	// For example /FOO and /..//Foo could be redirected to /foo.
	// RedirectTrailingSlash is independent of this option.
	RedirectFixedPath bool

	// If enabled, the router checks if another method is allowed for the
	// current route, if the current request can not be routed.
	// If this is the case, the request is answered with 'Method Not Allowed'
	// and HTTP status code 405.
	// If no other Method is allowed, the request is delegated to the NotFound
	// handler.
	HandleMethodNotAllowed bool

	// If enabled, the router automatically replies to OPTIONS requests.
	// Custom OPTIONS handlers take priority over automatic replies.
	HandleOPTIONS bool

	// An optional http.Handler that is called on automatic OPTIONS requests.
	// The handler is only called if HandleOPTIONS is true and no OPTIONS
	// handler for the specific path was set.
	// The "Allowed" header is set before calling the handler.
	GlobalOPTIONS http.Handler

	// Configurable http.Handler which is called when no matching route is
	// found. If it is not set, http.NotFound is used.
	NotFound http.Handler

	// Configurable http.Handler which is called when a request
	// cannot be routed and HandleMethodNotAllowed is true.
	// If it is not set, http.Error with http.StatusMethodNotAllowed is used.
	// The "Allow" header with allowed request methods is set before the handler
	// is called.
	MethodNotAllowed http.Handler

	// Function to handle panics recovered from http handlers.
	// It should be used to generate a error page and return the http error code
	// 500 (Internal Server Error).
	// The handler can be used to keep your server from crashing because of
	// unrecovered panics.
	PanicHandler func(http.ResponseWriter, *http.Request, interface{})
	// contains filtered or unexported fields
}
```

Router is a http.Handler which can be used to dispatch requests to different handler functions via configurable routes

​	Router 是一个 `http.Handler`，可用于通过可配置路由将请求分派到不同的处理程序函数

#### func New 

``` go
func New() *Router
```

New returns a new initialized Router. Path auto-correction, including trailing slashes, is enabled by default.

​	`New` 返回一个新的已初始化的 Router。默认情况下，启用路径自动更正，包括尾随斜杠。

#### (*Router) DELETE 

``` go
func (r *Router) DELETE(path string, handle Handle)
```

DELETE is a shortcut for router.Handle(http.MethodDelete, path, handle)

​	`DELETE` 是 `router.Handle(http.MethodDelete, path, handle)` 的快捷方式

#### (*Router) GET 

``` go
func (r *Router) GET(path string, handle Handle)
```

GET is a shortcut for router.Handle(http.MethodGet, path, handle)

​	`GET` 是 `router.Handle(http.MethodGet, path, handle)` 的快捷方式

#### (*Router) HEAD  <- 1.1.0

``` go
func (r *Router) HEAD(path string, handle Handle)
```

HEAD is a shortcut for router.Handle(http.MethodHead, path, handle)

​	`HEAD` 是 `router.Handle(http.MethodHead, path, handle)` 的快捷方式

#### (*Router) Handle 

``` go
func (r *Router) Handle(method, path string, handle Handle)
```

Handle registers a new request handle with the given path and method.

​	`Handle` 使用给定的路径和方法注册新的请求句柄。

For GET, POST, PUT, PATCH and DELETE requests the respective shortcut functions can be used.

​	对于 GET、POST、PUT、PATCH 和 DELETE 请求，可以使用相应的快捷方式函数。

This function is intended for bulk loading and to allow the usage of less frequently used, non-standardized or custom methods (e.g. for internal communication with a proxy).

​	此函数旨在进行批量加载，并允许使用不常用的、非标准化或自定义方法（例如，用于与代理进行内部通信）。

#### (*Router) Handler <- 1.1.0

``` go
func (r *Router) Handler(method, path string, handler http.Handler)
```

Handler is an adapter which allows the usage of an http.Handler as a request handle. The Params are available in the request context under ParamsKey.

​	`Handler` 是一个适配器，它允许将 `http.Handler` 用作请求句柄。`Params` 在 `ParamsKey` 下的请求上下文中可用。

#### (*Router) HandlerFunc 

``` go
func (r *Router) HandlerFunc(method, path string, handler http.HandlerFunc)
```

HandlerFunc is an adapter which allows the usage of an http.HandlerFunc as a request handle.

​	`HandlerFunc` 是一个适配器，它允许将 `http.HandlerFunc` 用作请求句柄。

#### (*Router) Lookup <- 1.1.0

``` go
func (r *Router) Lookup(method, path string) (Handle, Params, bool)
```

Lookup allows the manual lookup of a method + path combo. This is e.g. useful to build a framework around this router. If the path was found, it returns the handle function and the path parameter values. Otherwise the third return value indicates whether a redirection to the same path with an extra / without the trailing slash should be performed.

​	`Lookup` 允许手动查找方法 + 路径组合。例如，这对于围绕此路由器构建框架很有用。如果找到路径，它将返回处理函数和路径参数值。否则，第三个返回值指示是否应执行对具有额外 / 的相同路径的重定向，而无需尾随斜杠。

#### (*Router) OPTIONS <- 1.1.0

``` go
func (r *Router) OPTIONS(path string, handle Handle)
```

OPTIONS is a shortcut for router.Handle(http.MethodOptions, path, handle)

​	`OPTIONS` 是 `router.Handle(http.MethodOptions, path, handle)` 的快捷方式

#### (*Router) PATCH 

``` go
func (r *Router) PATCH(path string, handle Handle)
```

PATCH is a shortcut for router.Handle(http.MethodPatch, path, handle)

​	`PATCH` 是 `router.Handle(http.MethodPatch, path, handle)` 的快捷方式

#### (*Router) POST 

``` go
func (r *Router) POST(path string, handle Handle)
```

POST is a shortcut for router.Handle(http.MethodPost, path, handle)

​	`POST` 是 `router.Handle(http.MethodPost, path, handle)` 的快捷方式

#### (*Router) PUT 

``` go
func (r *Router) PUT(path string, handle Handle)
```

PUT is a shortcut for router.Handle(http.MethodPut, path, handle)

​	`PUT` 是 `router.Handle(http.MethodPut, path, handle)` 的快捷方式

#### (*Router) ServeFiles 

``` go
func (r *Router) ServeFiles(path string, root http.FileSystem)
```

ServeFiles serves files from the given file system root. The path must end with "`/*filepath`", files are then served from the local path `/defined/root/dir/*filepath`. For example if root is "`/etc`" and `*filepath` is "passwd", the local file "/etc/passwd" would be served. Internally a http.FileServer is used, therefore http.NotFound is used instead of the Router's NotFound handler. To use the operating system's file system implementation, use http.Dir:

​	`ServeFiles` 从给定的文件系统根目录提供文件。路径必须以“`/*filepath`”结尾，然后从本地路径 `/defined/root/dir/*filepath` 提供文件。例如，如果根目录是“`/etc`”，而 `*filepath` 是“`passwd`”，则会提供本地文件“`/etc/passwd`”。在内部使用 `http.FileServer`，因此使用 `http.NotFound` 而不是路由器的 NotFound 处理程序。要使用操作系统的文件系统实现，请使用 `http.Dir`：

```
router.ServeFiles("/src/*filepath", http.Dir("/var/www"))
```

#### (*Router) ServeHTTP 

``` go
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request)
```

ServeHTTP makes the router implement the http.Handler interface.

​	`ServeHTTP` 使路由器实现 `http.Handler` 接口。