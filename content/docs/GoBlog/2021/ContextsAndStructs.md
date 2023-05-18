+++
title = "上下文和结构"
weight = 94
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Contexts and structs - 上下文和结构

https://go.dev/blog/context-and-structs

Jean de Klerk, Matt T. Proud
24 February 2021

## Introduction 简介

In many Go APIs, especially modern ones, the first argument to functions and methods is often [`context.Context`](https://go.dev/pkg/context/). Context provides a means of transmitting deadlines, caller cancellations, and other request-scoped values across API boundaries and between processes. It is often used when a library interacts — directly or transitively — with remote servers, such as databases, APIs, and the like.

在许多Go API中，尤其是现代的API，函数和方法的第一个参数往往是context.Context。Context提供了一种跨API边界和进程之间传输截止日期、调用者取消和其他请求范围值的方法。当一个库与远程服务器（如数据库、API等）直接或转接地进行交互时，它经常被使用。

The [documentation for context](https://go.dev/pkg/context/) states:

上下文的文档指出：

> Contexts should not be stored inside a struct type, but instead passed to each function that needs it.上下文不应该存储在一个结构类型内，而应该传递给需要它的每个函数。

This article expands on that advice with reasons and examples describing why it’s important to pass Context rather than store it in another type. It also highlights a rare case where storing Context in a struct type may make sense, and how to do so safely.

本文对这一建议进行了阐述，并举例说明了为什么传递Context而不是将其存储在其他类型中很重要。它还强调了一种罕见的情况，即在结构类型中存储 Context 可能是合理的，以及如何安全地这样做。

## Prefer contexts passed as arguments 倾向于将上下文作为参数传递

To understand the advice to not store context in structs, let’s consider the preferred context-as-argument approach:

为了理解不要将上下文存储在结构中的建议，让我们考虑首选上下文作为参数的方法：

```go linenums="1"
// Worker fetches and adds works to a remote work orchestration server.
type Worker struct { /* … */ }

type Work struct { /* … */ }

func New() *Worker {
  return &Worker{}
}

func (w *Worker) Fetch(ctx context.Context) (*Work, error) {
  _ = ctx // A per-call ctx is used for cancellation, deadlines, and metadata.
}

func (w *Worker) Process(ctx context.Context, work *Work) error {
  _ = ctx // A per-call ctx is used for cancellation, deadlines, and metadata.
}
```

Here, the `(*Worker).Fetch` and `(*Worker).Process` methods both accept a context directly. With this pass-as-argument design, users can set per-call deadlines, cancellation, and metadata. And, it’s clear how the `context.Context` passed to each method will be used: there’s no expectation that a `context.Context` passed to one method will be used by any other method. This is because the context is scoped to as small an operation as it needs to be, which greatly increases the utility and clarity of `context` in this package.

这里，(*Worker).Fetch和(*Worker).Process方法都直接接受一个上下文。通过这种传递即参数的设计，用户可以设置每个调用的最后期限、取消和元数据。而且，传递给每个方法的context.Context将被如何使用是很清楚的：不存在传递给一个方法的context.Context会被其他方法使用的期望。这是因为上下文的范围是根据需要的小操作，这大大增加了这个包中上下文的效用和清晰度。

## Storing context in structs leads to confusion 将上下文存储在结构中会导致混乱

Let’s inspect again the `Worker` example above with the disfavored context-in-struct approach. The problem with it is that when you store the context in a struct, you obscure lifetime to the callers, or worse intermingle two scopes together in unpredictable ways:

让我们再次检查上面的Worker例子，用不受欢迎的context-in-struct方法。它的问题在于，当你将上下文存储在一个结构中时，你会对调用者的一生造成模糊，或者更糟糕的是以不可预测的方式将两个作用域混合在一起：

```go linenums="1"
type Worker struct {
  ctx context.Context
}

func New(ctx context.Context) *Worker {
  return &Worker{ctx: ctx}
}

func (w *Worker) Fetch() (*Work, error) {
  _ = w.ctx // A shared w.ctx is used for cancellation, deadlines, and metadata.
}

func (w *Worker) Process(work *Work) error {
  _ = w.ctx // A shared w.ctx is used for cancellation, deadlines, and metadata.
}
```

The `(*Worker).Fetch` and `(*Worker).Process` method both use a context stored in Worker. This prevents the callers of Fetch and Process (which may themselves have different contexts) from specifying a deadline, requesting cancellation, and attaching metadata on a per-call basis. For example: the user is unable to provide a deadline just for `(*Worker).Fetch`, or cancel just the `(*Worker).Process` call. The caller’s lifetime is intermingled with a shared context, and the context is scoped to the lifetime where the `Worker` is created.

(*Worker).Fetch和(*Worker).Process方法都使用存储在Worker中的上下文。这防止了Fetch和Process的调用者（它们本身可能有不同的上下文）在每个调用的基础上指定最后期限、请求取消和附加元数据。例如：用户无法仅仅为(*Worker).Fetch提供一个截止日期，或者仅仅取消(*Worker).Process的调用。调用者的生命周期与共享的上下文交织在一起，而上下文的范围是创建Worker的那个生命周期。

The API is also much more confusing to users compared to the pass-as-argument approach. Users might ask themselves:

与传递为参数的方法相比，该API对用户来说也更容易混淆。用户可能会问自己。

- Since `New` takes a `context.Context`, is the constructor doing work that needs cancellation or deadlines?既然New需要一个context.Context，那么构造函数是在做需要取消的工作还是死期？
- Does the `context.Context` passed in to `New` apply to work in `(*Worker).Fetch` and `(*Worker).Process`? Neither? One but not the other?传入New的context.Context是否适用于（*Worker）.Fetch和（*Worker）.Process中的工作？都不适用？一个但不是另一个？

The API would need a good deal of documentation to explicitly tell the user exactly what the `context.Context` is used for. The user might also have to read code rather than being able to rely on the structure of the API conveys.

API将需要大量的文档来明确地告诉用户context.Context到底是用来做什么的。用户可能还需要阅读代码，而不是依靠API所传达的结构。

And, finally, it can be quite dangerous to design a production-grade server whose requests don’t each have a context and thus can’t adequately honor cancellation. Without the ability to set per-call deadlines, [your process could backlog](https://sre.google/sre-book/handling-overload/) and exhaust its resources (like memory)!

而且，最后，设计一个生产级的服务器可能是相当危险的，因为它的请求并不是每个都有一个上下文，因此不能充分地履行取消。如果没有设置每个请求的最后期限的能力，你的进程可能会积压并耗尽它的资源（如内存）！因此，在设计生产级服务器时，要考虑到这一点。

## Exception to the rule: preserving backwards compatibility 规则的例外：保持向后的兼容性

When Go 1.7 — which [introduced context.Context](https://go.dev/doc/go1.7) — was released, a large number of APIs had to add context support in backwards compatible ways. For example, [`net/http`’s `Client` methods](https://go.dev/pkg/net/http/), like `Get` and `Do`, were excellent candidates for context. Each external request sent with these methods would benefit from having the deadline, cancellation, and metadata support that came with `context.Context`.

当Go 1.7--引入context.Context--发布时，大量的API不得不以向后兼容的方式增加对context的支持。例如，net/http的客户端方法，如Get和Do，是上下文的最佳候选者。用这些方法发送的每个外部请求都会从 context.Context 所提供的截止日期、取消和元数据支持中受益。

There are two approaches for adding support for `context.Context` in backwards compatible ways: including a context in a struct, as we’ll see in a moment, and duplicating functions, with duplicates accepting `context.Context` and having `Context` as their function name suffix. The duplicate approach should be preferred over the context-in-struct, and is further discussed in [Keeping your modules compatible](https://blog.golang.org/module-compatibility). However, in some cases it’s impractical: for example, if your API exposes a large number of functions, then duplicating them all might be infeasible.

有两种方法可以以向后兼容的方式添加对context.Context的支持：在一个结构中包含一个context，正如我们稍后看到的，以及复制函数，复制的函数接受context.Context并将Context作为其函数名的后缀。复制的方法应该比结构中的上下文更受欢迎，在《保持你的模块兼容》中会进一步讨论。然而，在某些情况下这是不切实际的：例如，如果你的 API 暴露了大量的函数，那么将它们全部重复可能是不可行的。

The `net/http` package chose the context-in-struct approach, which provides a useful case study. Let’s look at `net/http`’s `Do`. Prior to the introduction of `context.Context`, `Do` was defined as follows:

net/http 包选择了 context-in-struct 的方法，它提供了一个有用的案例研究。让我们来看看 net/http 的 Do。在引入 context.Context 之前，Do 的定义如下：

```go linenums="1"
// Do sends an HTTP request and returns an HTTP response [...]
func (c *Client) Do(req *Request) (*Response, error)
```

After Go 1.7, `Do` might have looked like the following, if not for the fact that it would break backwards compatibility:

在Go 1.7之后，如果不是因为Do会破坏向后的兼容性，它可能看起来像下面这样：

```go linenums="1"
// Do sends an HTTP request and returns an HTTP response [...]
func (c *Client) Do(ctx context.Context, req *Request) (*Response, error)
```

But, preserving the backwards compatibility and adhering to the [Go 1 promise of compatibility](https://go.dev/doc/go1compat) is crucial for the standard library. So, instead, the maintainers chose to add a `context.Context` on the `http.Request` struct in order to allow support `context.Context` without breaking backwards compatibility:

但是，保持向后的兼容性和遵守Go 1的兼容性承诺对标准库来说是至关重要的。因此，维护者选择在http.Request结构上添加context.Context，以允许支持context.Context而不破坏向后的兼容性：

```go linenums="1"
// A Request represents an HTTP request received by a server or to be sent by a client.
// ...
type Request struct {
  ctx context.Context

  // ...
}

// NewRequestWithContext returns a new Request given a method, URL, and optional
// body.
// [...]
// The given ctx is used for the lifetime of the Request.
func NewRequestWithContext(ctx context.Context, method, url string, body io.Reader) (*Request, error) {
  // Simplified for brevity of this article.
  return &Request{
    ctx: ctx,
    // ...
  }
}

// Do sends an HTTP request and returns an HTTP response [...]
func (c *Client) Do(req *Request) (*Response, error)
```

When retrofitting your API to support context, it may make sense to add a `context.Context` to a struct, as above. However, remember to first consider duplicating your functions, which allows retrofitting `context.Context` in a backwards compatibility without sacrificing utility and comprehension. For example:

当改造你的API以支持上下文时，将context.Context添加到一个结构中可能是有意义的，如上所述。然而，记得首先考虑重复你的函数，这样可以在不牺牲实用性和理解力的情况下，以向后兼容的方式改造context.Context。比如说：

```go linenums="1"
// Call uses context.Background internally; to specify the context, use
// CallContext.
func (c *Client) Call() error {
  return c.CallContext(context.Background())
}

func (c *Client) CallContext(ctx context.Context) error {
  // ...
}
```

## Conclusion 结论

Context makes it easy to propagate important cross-library and cross-API information down a calling stack. But, it must be used consistently and clearly in order to remain comprehensible, easy to debug, and effective.

Context使得重要的跨库和跨API信息可以很容易地在调用栈中传播。但是，为了保持可理解性、易于调试和有效，它的使用必须一致和明确。

When passed as the first argument in a method rather than stored in a struct type, users can take full advantage of its extensibility in order to build a powerful tree of cancellation, deadline, and metadata information through the call stack. And, best of all, its scope is clearly understood when it’s passed in as an argument, leading to clear comprehension and debuggability up and down the stack.

当作为方法的第一个参数传递而不是存储在结构类型中时，用户可以充分利用它的可扩展性，以便通过调用栈建立一个强大的取消、截止日期和元数据信息树。而且，最重要的是，当它被作为参数传入时，它的范围被清楚地理解，从而导致堆栈上下的清晰理解和可调试性。

When designing an API with context, remember the advice: pass `context.Context` in as an argument; don’t store it in structs.

在设计带有上下文的API时，请记住以下建议：将context.Context作为一个参数传入；不要将其存储在结构中。
