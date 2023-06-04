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

https://go.dev/blog/context

Sameer Ajmani
29 July 2014

## Introduction 简介

In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using.

在Go服务器中，每个传入的请求都在自己的goroutine中处理。请求处理程序经常启动额外的goroutine来访问后端，如数据库和RPC服务。处理一个请求的goroutine集通常需要访问请求的特定值，如最终用户的身份、授权令牌和请求的截止日期。当一个请求被取消或超时时，所有在该请求上工作的goroutines应该迅速退出，以便系统可以回收他们正在使用的任何资源。

At Google, we developed a `context` package that makes it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request. The package is publicly available as [context](https://go.dev/pkg/context). This article describes how to use the package and provides a complete working example.

在谷歌，我们开发了一个上下文包，它可以很容易地将请求范围的值、取消信号和最后期限跨越API边界传递给所有参与处理请求的goroutines。该包以context的形式公开提供。这篇文章描述了如何使用该包，并提供了一个完整的工作实例。

## Context 上下文

The core of the `context` package is the `Context` type:

Context包的核心是Context类型：

```go linenums="1"
// A Context carries a deadline, cancellation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```

(This description is condensed; the [godoc](https://go.dev/pkg/context) is authoritative.)

(这个描述是浓缩的，godoc是权威的。)

The `Done` method returns a channel that acts as a cancellation signal to functions running on behalf of the `Context`: when the channel is closed, the functions should abandon their work and return. The `Err` method returns an error indicating why the `Context` was canceled. The [Pipelines and Cancellation](https://go.dev/blog/pipelines) article discusses the `Done` channel idiom in more detail.

Done方法返回一个通道，作为代表Context运行的函数的取消信号：当通道关闭时，这些函数应该放弃它们的工作并返回。Err方法返回一个错误，表明Context被取消的原因。Pipelines and Cancellation一文更详细地讨论了Done通道的习性。

A `Context` does *not* have a `Cancel` method for the same reason the `Done` channel is receive-only: the function receiving a cancellation signal is usually not the one that sends the signal. In particular, when a parent operation starts goroutines for sub-operations, those sub-operations should not be able to cancel the parent. Instead, the `WithCancel` function (described below) provides a way to cancel a new `Context` value.

Context没有Cancel方法的原因与Done通道只接收信号的原因相同：接收取消信号的函数通常不是发送信号的那个。特别是，当一个父操作为子操作启动goroutines时，这些子操作不应该能够取消父操作。相反，WithCancel函数（如下所述）提供了一种取消新Context值的方法。

A `Context` is safe for simultaneous use by multiple goroutines. Code can pass a single `Context` to any number of goroutines and cancel that `Context` to signal all of them.

一个Context对于多个goroutine同时使用是安全的。代码可以将一个单一的Context传递给任意数量的goroutine，并取消该Context以向所有的goroutine发出信号。

The `Deadline` method allows functions to determine whether they should start work at all; if too little time is left, it may not be worthwhile. Code may also use a deadline to set timeouts for I/O operations.

Deadline方法允许函数决定它们是否应该开始工作；如果剩下的时间太少，可能就不值得了。代码也可以使用截止日期来设置I/O操作的超时。

`Value` allows a `Context` to carry request-scoped data. That data must be safe for simultaneous use by multiple goroutines.

Value允许一个Context携带请求范围的数据。该数据必须是安全的，可以被多个goroutine同时使用。

### Derived contexts 派生语境

The `context` package provides functions to *derive* new `Context` values from existing ones. These values form a tree: when a `Context` is canceled, all `Contexts` derived from it are also canceled.

context包提供了从现有的Context值派生新Context的函数。这些值形成一棵树：当一个Context被取消时，所有从它派生的Context也被取消。

`Background` is the root of any `Context` tree; it is never canceled:

Background是任何Context树的根；它从不被取消：

```go linenums="1"
// Background returns an empty Context. It is never canceled, has no deadline,
// and has no values. Background is typically used in main, init, and tests,
// and as the top-level Context for incoming requests.
func Background() Context
```

`WithCancel` and `WithTimeout` return derived `Context` values that can be canceled sooner than the parent `Context`. The `Context` associated with an incoming request is typically canceled when the request handler returns. `WithCancel` is also useful for canceling redundant requests when using multiple replicas. `WithTimeout` is useful for setting a deadline on requests to backend servers:

WithCancel和WithTimeout返回派生的Context值，这些值可以比父Context更早被取消。与传入请求相关的Context通常在请求处理程序返回时被取消。WithCancel对于使用多个副本时取消多余的请求也很有用。WithTimeout对于设置对后端服务器的请求的最后期限很有用：

```go linenums="1"
// WithCancel returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed or cancel is called.
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// A CancelFunc cancels a Context.
type CancelFunc func()

// WithTimeout returns a copy of parent whose Done channel is closed as soon as
// parent.Done is closed, cancel is called, or timeout elapses. The new
// Context's Deadline is the sooner of now+timeout and the parent's deadline, if
// any. If the timer is still running, the cancel function releases its
// resources.
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

`WithValue` provides a way to associate request-scoped values with a `Context`:

WithValue提供了一种将请求范围的值与一个上下文联系起来的方法：

```go linenums="1"
// WithValue returns a copy of parent whose Value method returns val for key.
func WithValue(parent Context, key interface{}, val interface{}) Context
```

The best way to see how to use the `context` package is through a worked example.

了解如何使用Context包的最好方法是通过一个工作实例。

## Example: Google Web Search 例子：谷歌网络搜索

Our example is an HTTP server that handles URLs like `/search?q=golang&timeout=1s` by forwarding the query "golang" to the [Google Web Search API](https://developers.google.com/web-search/docs/) and rendering the results. The `timeout` parameter tells the server to cancel the request after that duration elapses.

我们的例子是一个HTTP服务器，通过将查询 "golang "转发给Google Web Search API并呈现结果来处理/search?q=golang&timeout=1s等URL。超时参数告诉服务器在该时间段过后取消请求。

The code is split across three packages:

该代码被分成三个包：

- [server](https://go.dev/blog/context/server/server.go) provides the `main` function and the handler for `/search`. server提供主函数和/search的处理程序。
- [userip](https://go.dev/blog/context/userip/userip.go) provides functions for extracting a user IP address from a request and associating it with a `Context`. userip提供了从请求中提取用户IP地址并将其与Context关联的函数。
- [google](https://go.dev/blog/context/google/google.go) provides the `Search` function for sending a query to Google. google提供搜索功能，用于向Google发送查询。

### The server program 服务器程序

The [server](https://go.dev/blog/context/server/server.go) program handles requests like `/search?q=golang` by serving the first few Google search results for `golang`. It registers `handleSearch` to handle the `/search` endpoint. The handler creates an initial `Context` called `ctx` and arranges for it to be canceled when the handler returns. If the request includes the `timeout` URL parameter, the `Context` is canceled automatically when the timeout elapses:

服务器程序处理像 /search?q=golang 这样的请求，提供 golang 的前几个 Google 搜索结果。它注册了 handleSearch 来处理 /search 端点。该处理程序创建了一个名为 ctx 的初始 Context，并安排它在处理程序返回时被取消。如果请求包括超时的 URL 参数，当超时过后，Context 会自动取消。

```go linenums="1"
func handleSearch(w http.ResponseWriter, req *http.Request) {
    // ctx is the Context for this handler. Calling cancel closes the
    // ctx.Done channel, which is the cancellation signal for requests
    // started by this handler.
    var (
        ctx    context.Context
        cancel context.CancelFunc
    )
    timeout, err := time.ParseDuration(req.FormValue("timeout"))
    if err == nil {
        // The request has a timeout, so create a context that is
        // canceled automatically when the timeout expires.
        ctx, cancel = context.WithTimeout(context.Background(), timeout)
    } else {
        ctx, cancel = context.WithCancel(context.Background())
    }
    defer cancel() // Cancel ctx as soon as handleSearch returns.
```

The handler extracts the query from the request and extracts the client’s IP address by calling on the `userip` package. The client’s IP address is needed for backend requests, so `handleSearch` attaches it to `ctx`:

处理程序从请求中提取查询，并通过调用 userip 包来提取客户端的 IP 地址。后台请求需要客户端的 IP 地址，所以 handleSearch 将其附加到 ctx 上：

```go linenums="1"
    // Check the search query.
    query := req.FormValue("q")
    if query == "" {
        http.Error(w, "no query", http.StatusBadRequest)
        return
    }

    // Store the user IP in ctx for use by code in other packages.
    userIP, err := userip.FromRequest(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    ctx = userip.NewContext(ctx, userIP)
```

The handler calls `google.Search` with `ctx` and the `query`:

处理程序用ctx和查询调用google.Search：

```go linenums="1"
    // Run the Google search and print the results.
    start := time.Now()
    results, err := google.Search(ctx, query)
    elapsed := time.Since(start)
```

If the search succeeds, the handler renders the results:

如果搜索成功，处理程序将渲染结果：

```go linenums="1"
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

### Package userip - userip 包

The [userip](https://go.dev/blog/context/userip/userip.go) package provides functions for extracting a user IP address from a request and associating it with a `Context`. A `Context` provides a key-value mapping, where the keys and values are both of type `interface{}`. Key types must support equality, and values must be safe for simultaneous use by multiple goroutines. Packages like `userip` hide the details of this mapping and provide strongly-typed access to a specific `Context` value.

userip包提供了从请求中提取用户IP地址的函数，并将其与一个Context相关联。Context提供了一个键值映射，其中的键和值都是interface{}类型。键的类型必须支持平等，而值必须是安全的，可以被多个goroutine同时使用。像userip这样的包隐藏了这种映射的细节，并提供对特定Context值的强类型访问。

To avoid key collisions, `userip` defines an unexported type `key` and uses a value of this type as the context key:

为了避免键的碰撞，userip定义了一个未导出的键类型，并使用该类型的值作为上下文键：

```go linenums="1"
// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

// userIPkey is the context key for the user IP address.  Its value of zero is
// arbitrary.  If this package defined other context keys, they would have
// different integer values.
const userIPKey key = 0
```

`FromRequest` extracts a `userIP` value from an `http.Request`:

FromRequest从一个http.Request中提取一个userIP值：

```go linenums="1"
func FromRequest(req *http.Request) (net.IP, error) {
    ip, _, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
    }
```

`NewContext` returns a new `Context` that carries a provided `userIP` value:

NewContext返回一个携带所提供的userip值的新Context：

```go linenums="1"
func NewContext(ctx context.Context, userIP net.IP) context.Context {
    return context.WithValue(ctx, userIPKey, userIP)
}
```

`FromContext` extracts a `userIP` from a `Context`:

FromContext从一个Context中提取一个userIP：

```go linenums="1"
func FromContext(ctx context.Context) (net.IP, bool) {
    // ctx.Value returns nil if ctx has no value for the key;
    // the net.IP type assertion returns ok=false for nil.
    userIP, ok := ctx.Value(userIPKey).(net.IP)
    return userIP, ok
}
```

### Package google - google 包

The [google.Search](https://go.dev/blog/context/google/google.go) function makes an HTTP request to the [Google Web Search API](https://developers.google.com/web-search/docs/) and parses the JSON-encoded result. It accepts a `Context` parameter `ctx` and returns immediately if `ctx.Done` is closed while the request is in flight.

google.Search函数向谷歌网络搜索API发出HTTP请求，并解析JSON编码的结果。它接受一个Context参数ctx，如果ctx.Done在请求运行中被关闭，则立即返回。

The Google Web Search API request includes the search query and the user IP as query parameters:

谷歌网络搜索API请求包括搜索查询和用户IP作为查询参数：

```go linenums="1"
func Search(ctx context.Context, query string) (Results, error) {
    // Prepare the Google Search API request.
    req, err := http.NewRequest("GET", "https://ajax.googleapis.com/ajax/services/search/web?v=1.0", nil)
    if err != nil {
        return nil, err
    }
    q := req.URL.Query()
    q.Set("q", query)

    // If ctx is carrying the user IP address, forward it to the server.
    // Google APIs use the user IP to distinguish server-initiated requests
    // from end-user requests.
    if userIP, ok := userip.FromContext(ctx); ok {
        q.Set("userip", userIP.String())
    }
    req.URL.RawQuery = q.Encode()
```

`Search` uses a helper function, `httpDo`, to issue the HTTP request and cancel it if `ctx.Done` is closed while the request or response is being processed. `Search` passes a closure to `httpDo` handle the HTTP response:

搜索使用一个辅助函数httpDo来发出HTTP请求，如果ctx.Done在处理请求或响应时被关闭，则取消该请求。搜索将一个闭包传递给httpDo处理HTTP响应：

```go linenums="1"
    var results Results
    err = httpDo(ctx, req, func(resp *http.Response, err error) error {
        if err != nil {
            return err
        }
        defer resp.Body.Close()

        // Parse the JSON search result.
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
    return results, err
```

The `httpDo` function runs the HTTP request and processes its response in a new goroutine. It cancels the request if `ctx.Done` is closed before the goroutine exits:

httpDo函数在一个新的goroutine中运行HTTP请求并处理其响应。如果ctx.Done在goroutine退出前被关闭，它将取消请求：

```go linenums="1"
func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
    // Run the HTTP request in a goroutine and pass the response to f.
    c := make(chan error, 1)
    req = req.WithContext(ctx)
    go func() { c <- f(http.DefaultClient.Do(req)) }()
    select {
    case <-ctx.Done():
        <-c // Wait for f to return.
        return ctx.Err()
    case err := <-c:
        return err
    }
}
```

## Adapting code for Contexts 为上下文调整代码

Many server frameworks provide packages and types for carrying request-scoped values. We can define new implementations of the `Context` interface to bridge between code using existing frameworks and code that expects a `Context` parameter.

许多服务器框架为携带请求范围的值提供了包和类型。我们可以定义新的Context接口的实现，在使用现有框架的代码和期望有Context参数的代码之间架起桥梁。

For example, Gorilla’s [github.com/gorilla/context](http://www.gorillatoolkit.org/pkg/context) package allows handlers to associate data with incoming requests by providing a mapping from HTTP requests to key-value pairs. In [gorilla.go](https://go.dev/blog/context/gorilla/gorilla.go), we provide a `Context` implementation whose `Value` method returns the values associated with a specific HTTP request in the Gorilla package.

例如，Gorilla的github.com/gorilla/context包允许处理程序通过提供从HTTP请求到键值对的映射，将数据与传入的请求联系起来。在gorilla.go中，我们提供了一个Context实现，其Value方法返回与Gorilla包中特定HTTP请求相关的值。

Other packages have provided cancellation support similar to `Context`. For example, [Tomb](https://godoc.org/gopkg.in/tomb.v2) provides a `Kill` method that signals cancellation by closing a `Dying` channel. `Tomb` also provides methods to wait for those goroutines to exit, similar to `sync.WaitGroup`. In [tomb.go](https://go.dev/blog/context/tomb/tomb.go), we provide a `Context` implementation that is canceled when either its parent `Context` is canceled or a provided `Tomb` is killed.

其他包也提供了类似于Context的取消支持。例如，Tomb提供了一个Kill方法，通过关闭一个Dying通道来发出取消信号。Tomb还提供了等待那些goroutines退出的方法，类似于sync.WaitGroup。在tomb.go中，我们提供了一个Context的实现，当它的父Context被取消或提供的Tomb被杀死时，Context就会被取消。

## Conclusion 总结

At Google, we require that Go programmers pass a `Context` parameter as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancellation and ensures that critical values like security credentials transit Go programs properly.

在Google，我们要求Go程序员在传入和传出请求之间的调用路径上向每个函数传递一个Context参数作为第一个参数。这使得许多不同团队开发的 Go 代码能够很好地互操作。它提供了对超时和取消的简单控制，并确保像安全凭证这样的关键值能够正确地转运到Go程序中。

Server frameworks that want to build on `Context` should provide implementations of `Context` to bridge between their packages and those that expect a `Context` parameter. Their client libraries would then accept a `Context` from the calling code. By establishing a common interface for request-scoped data and cancellation, `Context` makes it easier for package developers to share code for creating scalable services.

希望建立在Context基础上的服务器框架应该提供Context的实现，在他们的包和那些期望有Context参数的包之间建立桥梁。他们的客户端库将接受来自调用代码的Context。通过为请求范围内的数据和取消建立一个通用的接口，Context使包的开发者更容易分享代码以创建可扩展的服务。
