+++
title = "使用 slog 进行结构化日志记录"
weight = 85
date = 2023-08-25T21:19:14+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Structured Logging with slog - 使用 slog 进行结构化日志记录

https://go.dev/blog/slog

Jonathan Amsterdam
22 August 2023

乔纳森·阿姆斯特丹 2023年8月22日

The new `log/slog` package in Go 1.21 brings structured logging to the standard library. Structured logs use key-value pairs so they can be parsed, filtered, searched, and analyzed quickly and reliably. For servers, logging is an important way for developers to observe the detailed behavior of the system, and often the first place they go to debug it. Logs therefore tend to be voluminous, and the ability to search and filter them quickly is essential.

​	Go 1.21 中的新 `log/slog` 包为标准库引入了结构化日志记录。结构化日志使用键值对，因此可以快速可靠地解析、过滤、搜索和分析。对于服务器来说，日志记录是开发人员观察系统详细行为的重要方式，通常是他们调试的第一个地方。因此，日志往往是大量的，快速搜索和过滤它们的能力是至关重要的。

The standard library has had a logging package, `log`, since Go’s initial release over a decade ago. Over time, we’ve learned that structured logging is important to Go programmers. It has consistently ranked high in our annual survey, and many packages in the Go ecosystem provide it. Some of these are quite popular: one of the first structured logging packages for Go, [logrus](https://pkg.go.dev/github.com/sirupsen/logrus), is used in over 100,000 other packages.

​	自 Go 十多年前的初始发布以来，标准库就有一个日志记录包 `log`。随着时间的推移，我们了解到结构化日志对于 Go 程序员非常重要。在我们的年度调查中，它一直排名靠前，并且 Go 生态系统中的许多包都提供了它。其中一些非常受欢迎：Go 的第一个结构化日志包之一 [logrus](https://pkg.go.dev/github.com/sirupsen/logrus) 在其他超过 100,000 个包中使用。

With many structured logging packages to choose from, large programs will often end up including more than one through their dependencies. The main program might have to configure each of these logging packages so that the log output is consistent: it all goes to the same place, in the same format. By including structured logging in the standard library, we can provide a common framework that all the other structured logging packages can share.

​	由于有许多结构化日志包可供选择，大型程序通常会通过它们的依赖关系包含不止一个包。主程序可能需要配置每个这些日志包，以使日志输出保持一致：所有日志都输出到相同的位置，使用相同的格式。通过将结构化日志记录包含在标准库中，我们可以提供一个共同的框架，所有其他结构化日志记录包都可以共享。

## 对 `slog` 的介绍 A tour of `slog`

Here is the simplest program that uses `slog`:

​	以下是使用 `slog` 的最简单的程序示例：

```go
package main

import "log/slog"

func main() {
    slog.Info("hello, world")
}
```

As of this writing, it prints:

​	截至本文撰写时，它会输出：

```bash
2023/08/04 16:09:19 INFO hello, world
```

The `Info` function prints a message at the Info log level using the default logger, which in this case is the default logger from the `log` package—the same logger you get when you write `log.Printf`. That explains why the output looks so similar: only the “INFO” is new. Out of the box, `slog` and the original `log` package work together to make it easy to get started.

​	`Info` 函数使用默认的记录器在 Info 日志级别打印消息，在本例中即来自 `log` 包的默认记录器，就像你在使用 `log.Printf` 时获得的记录器一样。这就解释了为什么输出看起来如此相似：只有“INFO”是新增的。在默认情况下，`slog` 和原始的 `log` 包一起工作，使得入门变得容易。

Besides `Info`, there are functions for three other levels—`Debug`, `Warn`, and `Error`—as well as a more general `Log` function that takes the level as an argument. In `slog`, levels are just integers, so you aren’t limited to the four named levels. For example, `Info` is zero and `Warn` is 4, so if your logging system has a level in between those, you can use 2 for it.

​	除了 `Info`，还有三个其他级别的函数——`Debug`、`Warn` 和 `Error`，以及一个更通用的 `Log` 函数，该函数将级别作为参数。在 `slog` 中，级别只是整数，因此不限于这四个命名级别。例如，`Info` 为零，`Warn` 为 4，因此如果你的日志记录系统在这两者之间有一个级别，你可以使用 2 代表它。

Unlike with the `log` package, we can easily add key-value pairs to our output by writing them after the message:

​	与 `log` 包不同的是，我们可以通过在消息之后写入键值对来轻松地将键值对添加到输出中：

```go
slog.Info("hello, world", "user", os.Getenv("USER"))
```

The output now looks like this:

​	现在输出看起来像这样：

```bash
2023/08/04 16:27:19 INFO hello, world user=jba
```

As we mentioned, `slog`’s top-level functions use the default logger. We can get this logger explicitly, and call its methods:

​	正如我们所提到的，`slog` 的顶层函数使用默认的记录器。我们可以显式地获取这个记录器，并调用它的方法：

```go
logger := slog.Default()
logger.Info("hello, world", "user", os.Getenv("USER"))
```

Every top-level function corresponds to a method on a `slog.Logger`. The output is the same as before.

​	每个顶层函数对应于 `slog.Logger` 上的一个方法。输出与之前相同。

Initially, slog’s output goes through the default `log.Logger`, producing the output we’ve seen above. We can change the output by changing the *handler* used by the logger. `slog` comes with two built-in handlers. A `TextHandler` emits all log information in the form `key=value`. This program creates a new logger using a `TextHandler` and makes the same call to the `Info` method:

​	最初，slog 的输出通过默认的 `log.Logger` 进行处理，产生了我们上面看到的输出。我们可以通过改变记录器使用的 *处理程序* 来改变输出。`slog` 配备了两个内置的处理程序。`TextHandler` 以 `key=value` 的形式发出所有日志信息。以下程序使用 `TextHandler` 创建一个新的记录器，并进行与 `Info` 方法相同的调用：

```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
logger.Info("hello, world", "user", os.Getenv("USER"))
```

Now the output looks like this:

​	现在输出看起来像这样：

```bash
time=2023-08-04T16:56:03.786-04:00 level=INFO msg="hello, world" user=jba
```

Everything has been turned into a key-value pair, with strings quoted as needed to preserve structure.

​	所有内容都已转换为键值对，需要时字符串被引用以保留结构。

For JSON output, install the built-in `JSONHandler` instead:

​	要获得 JSON 输出，可以安装内置的 `JSONHandler`：

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("hello, world", "user", os.Getenv("USER"))
```

Now our output is a sequence of JSON objects, one per logging call:

​	现在我们的输出是一系列 JSON 对象，每个日志调用一个对象：

```json
{"time":"2023-08-04T16:58:02.939245411-04:00","level":"INFO","msg":"hello, world","user":"jba"}
```

You are not limited to the built-in handlers. Anyone can write a handler by implementing the `slog.Handler` interface. A handler can generate output in a particular format, or it can wrap another handler to add functionality. One of the [examples](https://pkg.go.dev/log/slog@master#example-Handler-LevelHandler) in the `slog` documentation shows how to write a wrapping handler that changes the minimum level at which log messages will be displayed.

​	你不受限于内置的处理程序。任何人都可以通过实现 `slog.Handler` 接口来编写处理程序。处理程序可以生成特定格式的输出，也可以包装另一个处理程序以添加功能。`slog` 文档中的一个 [示例](https://pkg.go.dev/log/slog@master#example-Handler-LevelHandler) 显示了如何编写一个包装处理程序，以更改将显示日志消息的最低级别。

The alternating key-value syntax for attributes that we’ve been using so far is convenient, but for frequently executed log statements it may be more efficient to use the `Attr` type and call the `LogAttrs` method. These work together to minimize memory allocations. There are functions for building `Attr`s out of strings, numbers, and other common types. This call to `LogAttrs` produces the same output as above, but does it faster:

​	到目前为止，我们一直使用的属性的交替键-值语法很方便，但对于频繁执行的日志语句，使用 `Attr` 类型并调用 `LogAttrs` 方法可能更高效。它们共同工作以最小化内存分配。有用于构建字符串、数字和其他常见类型的 `Attr` 的函数。以下对 `LogAttrs` 的调用会产生与上述相同的输出，但速度更快：

```go
slog.LogAttrs(context.Background(), slog.LevelInfo, "hello, world",
    slog.String("user", os.Getenv("USER")))
```

There is a lot more to `slog`:

​	`slog` 还有更多功能：

- As the call to `LogAttrs` shows, you can pass a `context.Context` to some log functions so a handler can extract context information like trace IDs. (Canceling the context does not prevent the log entry from being written.)
- 如 `LogAttrs` 调用所示，可以将 `context.Context` 传递给某些日志函数，以便处理程序可以提取上下文信息，例如跟踪 ID（取消上下文不会阻止写入日志条目）。
- You can call `Logger.With` to add attributes to a logger that will appear in all of its output, effectively factoring out the common parts of several log statements. This is not only convenient, but it can also help performance, as discussed below.
- 可以调用 `Logger.With` 为记录器添加属性，这些属性将出现在所有输出中，从而将多个日志语句的共同部分分解出来。这不仅方便，而且还可以提高性能，如下文所述。
- Attributes can be combined into groups. This can add more structure to your log output and can help to disambiguate keys that would otherwise be identical.
- 属性可以组合成组。这可以为您的日志输出添加更多结构，并有助于消除否则相同的键。
- You can control how a value appears in the logs by providing its type with a `LogValue` method. That can be used to [log the fields of a struct as a group](https://pkg.go.dev/log/slog@master#example-LogValuer-Group) or [redact sensitive data](https://pkg.go.dev/log/slog@master#example-LogValuer-Secret), among other things.
- 可以通过提供 `LogValue` 方法中的类型来控制值在日志中的显示方式。这可用于将结构的字段作为组[记录](https://pkg.go.dev/log/slog@master#example-LogValuer-Group)或[隐藏敏感数据](https://pkg.go.dev/log/slog@master#example-LogValuer-Secret)等其他操作。

The best place to learn about all of `slog` is the [package documentation](https://pkg.go.dev/log/slog).

​	学习有关 `slog` 的全部内容的最佳位置是 [包文档](https://pkg.go.dev/log/slog)。

## 性能 Performance

We wanted `slog` to be fast. For large-scale performance gains, we designed [the `Handler` interface](https://pkg.go.dev/log/slog#Handler) to provide optimization opportunities. The `Enabled` method is called at the beginning of every log event, giving the handler a chance to drop unwanted log events quickly. The `WithAttrs` and `WithGroup` methods let the handler format attributes added by `Logger.With` once, rather than at each logging call. This pre-formatting can provide a significant speedup when large attributes, like an `http.Request`, are added to a `Logger` and then used in many logging calls.

​	我们希望 `slog` 具有高性能。为了获得大规模性能提升，我们设计了 [Handler 接口](https://pkg.go.dev/log/slog#Handler) 来提供优化机会。`Enabled` 方法在每个日志事件开始时调用，使处理程序有机会快速丢弃不需要的日志事件。`WithAttrs` 和 `WithGroup` 方法允许处理程序一次格式化由 `Logger.With` 添加的属性，而不是在每个日志调用时进行格式化。当将大型属性（例如 `http.Request`）添加到 `Logger` 中并在许多日志调用中使用时，这种预格式化可以显著提速。

To inform our performance optimization work, we investigated typical patterns of logging in existing open-source projects. We found that over 95% of calls to logging methods pass five or fewer attributes. We also categorized the types of attributes, finding that a handful of common types accounted for the majority. We then wrote benchmarks that captured the common cases, and used them as a guide to see where the time went. The greatest gains came from paying careful attention to memory allocation.

​	为了指导我们的性能优化工作，我们调查了现有开源项目中的典型日志模式。我们发现超过 95% 的日志方法调用传递了五个或更少的属性。我们还对属性的类型进行了分类，发现少数常见类型占了大部分。然后，我们编写了捕捉常见情况的基准测试，并将其用作指南，以了解时间花在哪里。最大的收益来自于仔细关注内存分配。

## 设计过程 The design process

The `slog` package is one of the largest additions to the standard library since Go 1 was released in 2012. We wanted to take our time designing it, and we knew that community feedback would be essential.

​	`slog` 包是自 2012 年 Go 1 发布以来标准库中最大的增加之一。我们希望花时间来设计它，我们知道社区的反馈意见至关重要。

By April 2022, we had gathered enough data to demonstrate the importance of structured logging to the Go community. The Go team decided to explore adding it to the standard library.

​	到 2022 年 4 月，我们已经收集了足够的数据来向 Go 社区证明结构化日志对其重要性。Go 团队决定探索将其添加到标准库中。

We began by looking at how the existing structured logging packages were designed. We also took advantage of the large collection of open-source Go code stored on the Go module proxy to learn how these packages were actually used. Our first design was informed by this research as well as Go’s spirit of simplicity. We wanted an API that is light on the page and easy to understand, without sacrificing performance.

​	我们开始研究现有的结构化日志包是如何设计的。我们还利用存储在 Go 模块代理上的大量开源 Go 代码来了解这些包的实际使用情况。我们的第一个设计受到这项研究以及 Go 简洁精神的启发。我们希望 API 在页面上的表现轻盈且易于理解，而不会牺牲性能。

It was never a goal to replace existing third-party logging packages. They are all good at what they do, and replacing existing code that works well is rarely a good use of a developer’s time. We divided the API into a frontend, `Logger`, that calls a backend interface, `Handler`. That way, existing logging packages can talk to a common backend, so the packages that use them can interoperate without having to be rewritten. Handlers are written or in progress for many common logging packages, including [Zap](https://github.com/uber-go/zap/tree/master/exp/zapslog), [logr](https://github.com/go-logr/logr/pull/196) and [hclog](https://github.com/evanphx/go-hclog-slog).

​	从未将替换现有的第三方日志记录包作为目标。它们在它们所做的事情上都表现良好，而且替换运行良好的现有代码很少是开发人员时间的良好利用。我们将 API 划分为前端 `Logger` 和调用后端接口 `Handler`。这样，现有的日志包可以与共同的后端进行通信，因此使用它们的包可以相互操作，而无需重写。针对许多常见的日志包编写了处理程序，包括 [Zap](https://github.com/uber-go/zap/tree/master/exp/zapslog)、[logr](https://github.com/go-logr/logr/pull/196) 和 [hclog](https://github.com/evanphx/go-hclog-slog)。

We shared our initial design within the Go team and other developers who had extensive logging experience. We made alterations based on their feedback, and by August of 2022 we felt we had a workable design. On August 29, we made our [experimental implementation](https://github.com/golang/exp/tree/master/slog) public and began a [GitHub discussion](https://github.com/golang/go/discussions/54763) to hear what the community had to say. The response was enthusiastic and largely positive. Thanks to insightful comments from the designers and users of other structured logging packages, we made several changes and added a few features, like groups and the `LogValuer` interface. We changed the mapping from log levels to integers twice.

​	我们在 Go 团队和其他有丰富日志记录经验的开发人员中分享了我们的初始设计。我们根据他们的反馈进行了修改，到2022年8月，我们认为我们有了一个可行的设计。在2022年8月29日，我们使我们的[实验性实现](https://github.com/golang/exp/tree/master/slog)公开，并开始了一个[GitHub 讨论](https://github.com/golang/go/discussions/54763)来听取社区的意见。反应积极且大多数是正面的。得益于其他结构化日志包的设计师和用户提供的深刻评论，我们进行了几次更改并添加了一些功能，比如组和 `LogValuer` 接口。我们两次更改了从日志级别到整数的映射。

After two months and about 300 comments, we felt we were ready for an actual [proposal](https://go.dev/issue/56345) and accompanying [design doc](https://go.googlesource.com/proposal/+/03441cb358c7b27a8443bca839e5d7a314677ea6/design/56345-structured-logging.md). The proposal issue garnered over 800 comments and resulted in many improvements to the API and the implementation. Here are two examples of API changes, both concerning `context.Context`:

​	经过两个月和约300个评论后，我们觉得我们已经准备好了一个实际的[提案](https://go.dev/issue/56345)，以及伴随的[设计文档](https://go.googlesource.com/proposal/+/03441cb358c7b27a8443bca839e5d7a314677ea6/design/56345-structured-logging.md)。提案问题收到了800多个评论，并且对 API 和实现进行了许多改进。以下是两个关于 `context.Context` 的 API 更改示例： 

1. Originally the API supported adding loggers to a context. Many felt that this was a convenient way to plumb a logger easily through levels of code that didn’t care about it. But others felt it was smuggling in an implicit dependency, making the code harder to understand. Ultimately, we removed the feature as being too controversial.
2. 最初，API 支持将记录器添加到上下文中。许多人认为这是一种方便的方式，可以轻松地将记录器传递到不关心它的代码层次中。但其他人认为这是引入隐含依赖，使代码变得更难理解。最终，我们将该功能删除，因为它过于有争议。
3. We also wrestled with the related question of passing a context to logging methods, trying a number of designs. We initially resisted the standard pattern of passing the context as the first argument because we didn’t want every logging call to require a context, but ultimately created two sets of logging methods, one with a context and one without.
4. 我们还纠结于将上下文传递给日志方法的相关问题，尝试了许多设计。最初，我们抵制了将上下文作为第一个参数传递的标准模式，因为我们不想要每次日志调用都需要上下文，但最终创建了两组日志方法，一组带有上下文，一组不带。

One change we did not make concerned the alternating key-and-value syntax for expressing attributes:

​	我们没有进行的更改涉及用于表示属性的交替键和值语法：

```go
slog.Info("message", "k1", v1, "k2", v2)
```

Many felt strongly that this was a bad idea. They found it hard to read and easy to get wrong by omitting a key or value. They preferred explicit attributes for expressing structure:

​	许多人强烈认为这是一个不好的主意。他们发现很难阅读，易于由于遗漏键或值而出错。他们更喜欢使用显式属性来表示结构：

```go
slog.Info("message", slog.Int("k1", v1), slog.String("k2", v2))
```

But we felt that the lighter syntax was important to keeping Go easy and fun to use, especially for new Go programmers. We also knew that several Go logging packages, like `logr`, `go-kit/log` and `zap` (with its `SugaredLogger`) successfully used alternating keys and values. We added a [vet check](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/slog) to catch common mistakes, but did not change the design.

​	但我们认为轻便的语法对于保持 Go 易于使用和有趣对于新的 Go 程序员来说非常重要。我们还知道几个 Go 日志记录包，如 `logr`、`go-kit/log` 和 `zap`（其 `SugaredLogger`）成功使用交替键和值。我们添加了一个[vet 检查](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/slog)，以捕捉常见错误，但没有更改设计。

On March 15, 2023, the proposal was accepted, but there were still some minor unresolved issues. Over the next few weeks, ten additional changes were proposed and resolved. By early July, the `log/slog` package implementation was complete, along with the `testing/slogtest` package for verifying handlers and the vet check for correct usage of alternating keys and values.

​	2023年3月15日，提案被接受，但仍存在一些小的未解决问题。在接下来的几周里，提出并解决了十个附加的更改。到7月初，`log/slog` 包的实现完成了，以及用于验证处理程序的 `testing/slogtest` 包，以及用于正确使用交替键和值的 vet 检查。

And on August 8, Go 1.21 was released, and `slog` with it. We hope you find it useful, and as fun to use as it was to build.

​	在2023年8月8日，Go 1.21 发布，`slog` 随之发布。我们希望你发现它有用，并且使用起来和构建起来一样有趣。

And a big thanks to everyone who participated in the discussion and the proposal process. Your contributions improved `slog` immensely.

​	同时也要感谢所有参与讨论和提案过程的人。你们的贡献极大地改进了 `slog`。

## 资源 Resources

The [documentation](https://pkg.go.dev/log/slog) for the `log/slog` package explains how to use it and provides several examples.

​	`log/slog` 包的[文档](https://pkg.go.dev/log/slog)解释了如何使用它并提供了几个示例。

The [wiki page](https://github.com/golang/go/wiki/Resources-for-slog) has additional resources provided by the Go community, including a variety of handlers.

​	[维基页面](https://github.com/golang/go/wiki/Resources-for-slog)提供了由 Go 社区提供的额外资源，包括各种处理程序。

If you want to write a handler, consult the [handler writing guide](https://github.com/golang/example/blob/master/slog-handler-guide/README.md).

​	如果你想编写一个处理程序，请参考[处理程序编写指南](https://github.com/golang/example/blob/master/slog-handler-guide/README.md)。