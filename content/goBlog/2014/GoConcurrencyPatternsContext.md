+++
title = "go 并发模式：Context"
weight = 8
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go Concurrency Patterns: Context - go 并发模式：Context

> 原文：[https://go.dev/blog/context](https://go.dev/blog/context)

Sameer Ajmani
29 July 2014

2014年7月29日

## 简介 Introduction 

In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using.

​	在 Go 服务器中，每个传入的请求都在自己的 goroutine 中处理。请求处理程序经常会启动额外的 goroutine 来访问诸如数据库和 RPC 服务之类的后端。处理请求的一组 goroutine 通常需要访问请求特定的值，比如终端用户的身份、授权令牌和请求的截止时间。当请求被取消或超时时，所有处理该请求的 goroutine 应该迅速退出，以便系统可以回收它们正在使用的任何资源。

At Google, we developed a `context` package that makes it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request. The package is publicly available as [context](https://go.dev/pkg/context). This article describes how to use the package and provides a complete working example.

​	在 Google，我们开发了一个 `context` 包，用于轻松地在处理请求的所有涉及的 goroutine 之间传递请求范围的值、取消信号和截止时间。该包作为 [context](https://go.dev/pkg/context) 公开可用。本文描述了如何使用该包并提供了一个完整的工作示例。

## Context

The core of the `context` package is the `Context` type:

​	`context` 包的核心是 `Context` 类型：

```go
// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
// Context 携带截止时间、取消信号和跨 API 边界的请求范围的值。
// 它的方法对多个 goroutine 同时使用是安全的。
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    // Done 返回一个通道，在此 Context 被取消或超时时关闭。
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    // Err 在 Done 通道关闭后，指示为什么取消了此 Context。
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    // Deadline 返回此 Context 将被取消的时间（如果有的话）。
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    // Value 返回与键关联的值，如果没有则返回 nil。
    Value(key interface{}) interface{}
}
```

(This description is condensed; the [godoc](https://go.dev/pkg/context) is authoritative.)

（此描述已被缩减；[godoc](https://go.dev/pkg/context) 是权威的参考。）

The `Done` method returns a channel that acts as a cancellation signal to functions running on behalf of the `Context`: when the channel is closed, the functions should abandon their work and return. The `Err` method returns an error indicating why the `Context` was canceled. The [Pipelines and Cancellation](https://go.dev/blog/pipelines) article discusses the `Done` channel idiom in more detail.

​	`Done` 方法返回一个通道，作为函数在 `Context` 的代表上运行时的取消信号：当通道关闭时，函数应该放弃它们的工作并返回。`Err` 方法返回一个指示为什么取消了 `Context` 的错误。[管道和取消]({{< ref "/goBlog/2014/GoConcurrencyPatternsPipelinesAndCancellation">}}) 文章在更多细节中讨论了 `Done` 通道的惯用法。

A `Context` does *not* have a `Cancel` method for the same reason the `Done` channel is receive-only: the function receiving a cancellation signal is usually not the one that sends the signal. In particular, when a parent operation starts goroutines for sub-operations, those sub-operations should not be able to cancel the parent. Instead, the `WithCancel` function (described below) provides a way to cancel a new `Context` value.

​	`Context` 没有 `Cancel` 方法，原因与 `Done` 通道是只接收的原因相同：接收取消信号的函数通常不是发送信号的函数。特别地，当父操作为子操作启动 goroutine 时，这些子操作不应能够取消父操作。相反，`WithCancel` 函数（稍后会介绍）提供了一种取消新的 `Context` 值的方法。

A `Context` is safe for simultaneous use by multiple goroutines. Code can pass a single `Context` to any number of goroutines and cancel that `Context` to signal all of them.

​	`Context` 对于多个 goroutine 同时使用是安全的。代码可以将单个 `Context` 传递给任意数量的 goroutine，并取消该 `Context` 以向所有的goroutine发出信号。

The `Deadline` method allows functions to determine whether they should start work at all; if too little time is left, it may not be worthwhile. Code may also use a deadline to set timeouts for I/O operations.

​	`Deadline` 方法允许函数确定它们是否应该开始工作；如果剩余的时间太少，可能就不值得了。代码还可以使用截止时间设置 I/O 操作的超时。

`Value` allows a `Context` to carry request-scoped data. That data must be safe for simultaneous use by multiple goroutines.

​	`Value` 允许 `Context` 携带请求范围的数据。该数据必须对多个 goroutine 同时使用是安全的。

### 派生的上下文 - Derived contexts 

The `context` package provides functions to *derive* new `Context` values from existing ones. These values form a tree: when a `Context` is canceled, all `Contexts` derived from it are also canceled.

​	`context` 包提供了从现有上下文中 *派生* 新的 `Context` 值的函数。这些值形成一棵树：当一个 `Context` 被取消时，从它派生的所有 `Context` 也会被取消。

`Background` is the root of any `Context` tree; it is never canceled:

​	`Background` 是任何 `Context` 树的根；它永远不会被取消：

```go
// Background returns an empty Context. It is never canceled, has no deadline,
// and has no values. Background is typically used in main, init, and tests,
// and as the top-level Context for incoming requests.
// Background 返回一个空的 Context。
// 它永远不会被取消，没有截止时间，也没有值。
// 通常在 main、init 和测试中使用，以及作为传入请求的顶级 Context。
func Background() Context
```

`WithCancel` and `WithTimeout` return derived `Context` values that can be canceled sooner than the parent `Context`. The `Context` associated with an incoming request is typically canceled when the request handler returns. `WithCancel` is also useful for canceling redundant requests when using multiple replicas. `WithTimeout` is useful for setting a deadline on requests to backend servers:

​	`WithCancel` 和 `WithTimeout` 返回派生的 `Context` 值，可以比父 `Context` 更早地取消。与传入请求相关联的 `Context` 通常在请求处理程序返回时被取消。在使用多个副本时，`WithCancel` 也适用于取消冗余请求。`WithTimeout` 对于设置对后端服务器的请求的截止时间很有用：

```go
// WithCancel returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed or cancel is called.
// WithCancel 返回 parent 的副本，
// 其 Done 通道在 parent.Done 被关闭或调用 cancel 时立即关闭。
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// A CancelFunc cancels a Context.
// CancelFunc 取消一个 Context。
type CancelFunc func()

// WithTimeout returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed, cancel is called, or timeout elapses. The new
// Context's Deadline is the sooner of now+timeout and the parent's deadline, if
// any. If the timer is still running, the cancel function releases its
// resources.
// WithTimeout 返回 parent 的副本，
// 其 Done 通道在 parent.Done 被关闭、cancel 被调用或超时过去时立即关闭。
// 新 Context 的截止时间是 now+timeout 和父截止时间中较早的那个。
// 如果计时器仍在运行，取消函数会释放其资源。
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

`WithValue` provides a way to associate request-scoped values with a `Context`:

​	`WithValue` 提供了一种将请求范围的值与 `Context` 关联的方法：

```go
// WithValue returns a copy of parent whose Value method returns val for key.
// WithValue 返回 parent 的副本，其 Value 方法返回 key 的值 val。
func WithValue(parent Context, key interface{}, val interface{}) Context
```

The best way to see how to use the `context` package is through a worked example.

​	了解如何使用 `context` 包的最佳方法是通过一个实际的示例。

## 示例：Google Web 搜索 - Example: Google Web Search

Our example is an HTTP server that handles URLs like `/search?q=golang&timeout=1s` by forwarding the query "golang" to the [Google Web Search API](https://developers.google.com/web-search/docs/) and rendering the results. The `timeout` parameter tells the server to cancel the request after that duration elapses.

​	我们的示例是一个 HTTP 服务器，通过将查询“golang”转发到 [Google Web 搜索 API](https://developers.google.com/web-search/docs/) 并呈现结果，来处理类似 `/search?q=golang&timeout=1s` 的 URL。`timeout` 参数告诉服务器在持续时间到达后取消请求。

The code is split across three packages:

​	该代码分为三个包：

- [server](https://go.dev/blog/context/server/server.go) provides the `main` function and the handler for `/search`. 
- [server](https://go.dev/blog/context/server/server.go) 提供 `main` 函数和处理 `/search` 的处理程序。
- [userip](https://go.dev/blog/context/userip/userip.go) provides functions for extracting a user IP address from a request and associating it with a `Context`. 
- [userip](https://go.dev/blog/context/userip/userip.go) 提供从请求中提取用户 IP 地址并将其与 `Context` 关联的函数。
- [google](https://go.dev/blog/context/google/google.go) provides the `Search` function for sending a query to Google. 
- [google](https://go.dev/blog/context/google/google.go) 提供将查询发送到 Google 的 `Search` 函数。

### 服务器程序 - The server program

The [server](https://go.dev/blog/context/server/server.go) program handles requests like `/search?q=golang` by serving the first few Google search results for `golang`. It registers `handleSearch` to handle the `/search` endpoint. The handler creates an initial `Context` called `ctx` and arranges for it to be canceled when the handler returns. If the request includes the `timeout` URL parameter, the `Context` is canceled automatically when the timeout elapses:

​	[server](https://go.dev/blog/context/server/server.go) 程序处理类似 `/search?q=golang` 的请求，通过为 `golang` 提供 Google 搜索结果的前几个结果。它注册 `handleSearch` 来处理 `/search` 终端。处理程序创建一个名为 `ctx` 的初始 `Context`，并在处理程序返回时安排取消它。如果请求包括 `timeout` URL 参数，则在超时到达时 `Context` 会被自动取消：

```go
func handleSearch(w http.ResponseWriter, req *http.Request) {
    // ctx is the Context for this handler. Calling cancel closes the
    // ctx.Done channel, which is the cancellation signal for requests
    // started by this handler.
    // ctx 是此处理程序的 Context。
    // 调用 cancel 会关闭 ctx.Done 通道，
    // 这是由此处理程序启动的请求的取消信号。
    var (
        ctx    context.Context
        cancel context.CancelFunc
    )
    timeout, err := time.ParseDuration(req.FormValue("timeout"))
    if err == nil {
        // The request has a timeout, so create a context that is
        // canceled automatically when the timeout expires.
        // 请求具有超时，因此创建一个在超时到达时自动取消的 Context。
        ctx, cancel = context.WithTimeout(context.Background(), timeout)
    } else {
        ctx, cancel = context.WithCancel(context.Background())
    }
    defer cancel() // Cancel ctx as soon as handleSearch returns. 在 handleSearch 返回后立即取消 ctx。
```

The handler extracts the query from the request and extracts the client’s IP address by calling on the `userip` package. The client’s IP address is needed for backend requests, so `handleSearch` attaches it to `ctx`:

​	处理程序从请求中提取查询并通过调用 `userip` 包提取客户端的 IP 地址。客户端的 IP 地址在后端请求中是必需的，因此 `handleSearch` 将其附加到 `ctx`：

```go
    // Check the search query.
	// 检查搜索查询。
    query := req.FormValue("q")
    if query == "" {
        http.Error(w, "no query", http.StatusBadRequest)
        return
    }

    // Store the user IP in ctx for use by code in other packages.
 	// 将用户 IP 存储在 ctx 中，以供其他包中的代码使用。
    userIP, err := userip.FromRequest(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    ctx = userip.NewContext(ctx, userIP)
```

The handler calls `google.Search` with `ctx` and the `query`:

​	处理程序使用 `ctx` 和查询调用 `google.Search`：

```go
    // Run the Google search and print the results.
	// 运行 Google 搜索并打印结果。
    start := time.Now()
    results, err := google.Search(ctx, query)
    elapsed := time.Since(start)
```

If the search succeeds, the handler renders the results:

​	如果搜索成功，处理程序呈现结果：

```go
    if err := resultsTemplate.Execute(w, struct {
        Results          google.Results
        Timeout, Elapsed time.Duration
    }{
        Results: results,
        Timeout: timeout,
        Elapsed: elapsed,
    }); err != nil {
        log.Print(err)
        return
    }
```

### userip 包 - Package userip

The [userip](https://go.dev/blog/context/userip/userip.go) package provides functions for extracting a user IP address from a request and associating it with a `Context`. A `Context` provides a key-value mapping, where the keys and values are both of type `interface{}`. Key types must support equality, and values must be safe for simultaneous use by multiple goroutines. Packages like `userip` hide the details of this mapping and provide strongly-typed access to a specific `Context` value.

​	[userip](https://go.dev/blog/context/userip/userip.go) 包提供了从请求中提取用户 IP 地址并将其与 `Context` 关联的函数。`Context` 提供键值映射，其中键和值都是 `interface{}` 类型。键类型必须支持相等性，值必须对多个 goroutine 同时使用是安全的。像 `userip` 这样的包隐藏了此映射的细节，并为特定 `Context` 值提供了强类型的访问。

To avoid key collisions, `userip` defines an unexported type `key` and uses a value of this type as the context key:

​	为了避免键冲突，`userip` 定义了一个未导出的类型 `key`，并使用此类型的值作为上下文键：

```go
// The key type is unexported to prevent collisions with context keys defined in
// other packages.
// 未导出的 key 类型，以防止与其他包中定义的上下文键冲突。
type key int

// userIPkey is the context key for the user IP address.  Its value of zero is
// arbitrary.  If this package defined other context keys, they would have
// different integer values.
// userIPKey 是用户 IP 地址的上下文键。
// 它的值为零是任意的。如果此包定义了其他上下文键，它们将有不同的整数值。
const userIPKey key = 0
```

`FromRequest` extracts a `userIP` value from an `http.Request`:

​	`FromRequest` 从 `http.Request` 中提取 `userIP` 值：

```go
func FromRequest(req *http.Request) (net.IP, error) {
    ip, _, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
    }
```

`NewContext` returns a new `Context` that carries a provided `userIP` value:

​	`NewContext` 返回一个携带提供的 `userIP` 值的新 `Context`：

```go
func NewContext(ctx context.Context, userIP net.IP) context.Context {
    return context.WithValue(ctx, userIPKey, userIP)
}
```

`FromContext` extracts a `userIP` from a `Context`:

​	`FromContext` 从 `Context` 中提取一个 `userIP`：

```go
func FromContext(ctx context.Context) (net.IP, bool) {
    // ctx.Value returns nil if ctx has no value for the key;
    // the net.IP type assertion returns ok=false for nil.
    // 如果 ctx 没有值与键关联，则 ctx.Value 返回 nil；
    // net.IP 类型断言在 nil 时返回 ok=false。
    userIP, ok := ctx.Value(userIPKey).(net.IP)
    return userIP, ok
}
```

### google 包 - Package google

The [google.Search](https://go.dev/blog/context/google/google.go) function makes an HTTP request to the [Google Web Search API](https://developers.google.com/web-search/docs/) and parses the JSON-encoded result. It accepts a `Context` parameter `ctx` and returns immediately if `ctx.Done` is closed while the request is in flight.

​	[google.Search](https://go.dev/blog/context/google/google.go) 函数向 [Google Web 搜索 API](https://developers.google.com/web-search/docs/) 发出 HTTP 请求并解析 JSON 编码的结果。它接受一个 `Context` 参数 `ctx`，如果请求正在进行中时关闭 `ctx.Done`，则立即返回。

The Google Web Search API request includes the search query and the user IP as query parameters:

​	Google Web 搜索 API 请求包括查询和用户 IP 作为查询参数：

```go
func Search(ctx context.Context, query string) (Results, error) {
    // Prepare the Google Search API request.
    // 准备 Google Web 搜索 API 请求。
    req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
    if err != nil {
        return nil, err
    }
    q := req.URL.Query()
    q.Set("q", query)

    // If ctx is carrying the user IP address, forward it to the server.
    // Google APIs use the user IP to distinguish server-initiated requests
    // from end-user requests.
    // 如果 ctx 携带用户 IP 地址，请将其转发给服务器。
    // Google APIs 使用用户 IP 来区分服务器发起的请求和终端用户请求。
    if userIP, ok := userip.FromContext(ctx); ok {
        q.Set("userip", userIP.String())
    }
    req.URL.RawQuery = q.Encode()
```

`Search` uses a helper function, `httpDo`, to issue the HTTP request and cancel it if `ctx.Done` is closed while the request or response is being processed. `Search` passes a closure to `httpDo` handle the HTTP response:

​	`Search` 使用一个辅助函数 `httpDo` 来发出 HTTP 请求，并在请求或响应处理过程中关闭 `ctx.Done` 时取消它。`Search` 将闭包传递给 `httpDo` 处理 HTTP 响应：

```go
    var results Results
    err = httpDo(ctx, req, func(resp *http.Response, err error) error {
        if err != nil {
            return err
        }
        defer resp.Body.Close()

        // Parse the JSON search result.
        // https://developers.google.com/web-search/docs/#fonje
        // 解析 JSON 搜索结果。
        // https://developers.google.com/web-search/docs/#fonje
        var data struct {
            ResponseData struct {
                Results []struct {
                    TitleNoFormatting string
                    URL               string
                }
            }
        }
        if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
            return err
        }
        for _, res := range data.ResponseData.Results {
            results = append(results, Result{Title: res.TitleNoFormatting, URL: res.URL})
        }
        return nil
    })
    // httpDo waits for the closure we provided to return, so it's safe to
    // read results here.
	// httpDo 等待我们提供的闭包返回，因此在此处读取结果是安全的。
    return results, err
```

The `httpDo` function runs the HTTP request and processes its response in a new goroutine. It cancels the request if `ctx.Done` is closed before the goroutine exits:

​	`httpDo` 函数在新的 goroutine 中运行 HTTP 请求并处理其响应。如果 goroutine 退出之前关闭了 `ctx.Done`，它会取消请求：

```go
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
    // Run the HTTP request in a goroutine and pass the response to f.
    // 在 goroutine 中运行 HTTP 请求并将响应传递给 f。
    c := make(chan error, 1)
    req = req.WithContext(ctx)
    go func() { c <- f(http.DefaultClient.Do(req)) }()
    select {
    case <-ctx.Done():
        <-c // Wait for f to return. 等待 f 返回
        return ctx.Err()
    case err := <-c:
        return err
    }
}
```

## 为上下文调整代码 - Adapting code for Contexts

Many server frameworks provide packages and types for carrying request-scoped values. We can define new implementations of the `Context` interface to bridge between code using existing frameworks and code that expects a `Context` parameter.

​	许多服务器框架提供用于携带请求范围值的包和类型。我们可以为现有框架的代码和预期 `Context` 参数的代码之间的调用路径定义新的 `Context` 接口实现，以建立桥梁。

For example, Gorilla’s [github.com/gorilla/context](http://www.gorillatoolkit.org/pkg/context) package allows handlers to associate data with incoming requests by providing a mapping from HTTP requests to key-value pairs. In [gorilla.go](https://go.dev/blog/context/gorilla/gorilla.go), we provide a `Context` implementation whose `Value` method returns the values associated with a specific HTTP request in the Gorilla package.

​	例如，Gorilla 的 [github.com/gorilla/context](http://www.gorillatoolkit.org/pkg/context) 包允许处理程序通过提供从 HTTP 请求到键值对的映射，将数据与传入请求关联起来。在 [gorilla.go](https://go.dev/blog/context/gorilla/gorilla.go) 中，我们提供了一个 `Context` 实现，其 `Value` 方法返回 Gorilla 包中特定 HTTP 请求关联的值。

Other packages have provided cancellation support similar to `Context`. For example, [Tomb](https://godoc.org/gopkg.in/tomb.v2) provides a `Kill` method that signals cancellation by closing a `Dying` channel. `Tomb` also provides methods to wait for those goroutines to exit, similar to `sync.WaitGroup`. In [tomb.go](https://go.dev/blog/context/tomb/tomb.go), we provide a `Context` implementation that is canceled when either its parent `Context` is canceled or a provided `Tomb` is killed.

​	其他包也提供了类似于 `Context` 的取消支持。例如，[Tomb](https://godoc.org/gopkg.in/tomb.v2) 提供了一个 `Kill` 方法，通过关闭 `Dying` 通道来发出取消信号。`Tomb` 还提供了等待这些 goroutine 退出的方法，类似于 `sync.WaitGroup`。在 [tomb.go](https://go.dev/blog/context/tomb/tomb.go) 中，我们提供了一个 `Context` 实现，当其父 `Context` 被取消或提供的 `Tomb` 被 kill 时，它会被取消。

## 结论 - Conclusion

At Google, we require that Go programmers pass a `Context` parameter as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancellation and ensures that critical values like security credentials transit Go programs properly.

​	在 Google，我们要求 Go 程序员将一个 `Context` 参数作为传入和传出请求之间的调用路径上每个函数的第一个参数传递。这允许由许多不同团队开发的 Go 代码进行良好的互操作性。它提供了简单的超时和取消控制，并确保像安全凭据这样的关键值正确传输到 Go 程序中。

Server frameworks that want to build on `Context` should provide implementations of `Context` to bridge between their packages and those that expect a `Context` parameter. Their client libraries would then accept a `Context` from the calling code. By establishing a common interface for request-scoped data and cancellation, `Context` makes it easier for package developers to share code for creating scalable services.

​	希望在 `Context` 上构建的服务器框架应该提供 `Context` 实现，以建立它们的包和期望 `Context` 参数的包之间的桥梁。它们的客户端库随后将从调用代码接受一个 `Context`。通过为请求范围数据和取消建立一个共同的接口，`Context` 使得包开发者更容易共享用于创建可扩展服务的代码。
