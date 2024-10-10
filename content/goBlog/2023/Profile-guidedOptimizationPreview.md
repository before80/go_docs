+++
title = "性能分析引导优化预览"
weight = 98
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Profile-guided optimization preview - 性能分析引导优化预览

> 原文：[https://go.dev/blog/pgo-preview](https://go.dev/blog/pgo-preview)
>

Michael Pratt
8 February 2023

2023年2月8日

When you build a Go binary, the Go compiler performs optimizations to try to generate the best performing binary it can. For example, constant propagation can evaluate constant expressions at compile time, avoiding runtime evaluation cost. Escape analysis avoids heap allocations for locally-scoped objects, avoiding GC overheads. Inlining copies the body of simple functions into callers, often enabling further optimization in the caller (such as additional constant propagation or better escape analysis).

​	当您构建一个 Go 二进制文件时，Go 编译器会进行优化，以尝试生成性能最佳的二进制文件。例如，常量传播可以在编译时评估常量表达式，避免运行时的评估成本。逃逸分析可以避免在局部作用域对象上进行堆分配，避免垃圾回收的开销。内联将简单函数的主体复制到调用方，通常可以在调用方进行进一步的优化（例如更多的常量传播或更好的逃逸分析）。

Go improves optimizations from release to release, but this is not always an easy task. Some optimizations are tunable, but the compiler can’t just "turn it up to 11" on every function because overly aggressive optimizations can actually hurt performance or cause excessive build times. Other optimizations require the compiler to make a judgment call about what the "common" and "uncommon" paths in a function are. The compiler must make a best guess based on static heuristics because it can’t know which cases will be common at run time.

​	Go 在每个版本中改进了优化，但这并不总是一项容易的任务。某些优化是可调整的，但编译器不能简单地对每个函数进行过度激进的优化，因为过度激进的优化实际上可能会降低性能或导致过长的构建时间。其他优化需要编译器根据函数中的“常见”和“不常见”路径做出判断。编译器必须根据静态启发式方法做出最佳猜测，因为它无法知道运行时哪些情况将是常见的。

Or can it?

​	那么，它能做到吗？

With no definitive information about how the code is used in a production environment, the compiler can operate only on the source code of packages. But we do have a tool to evaluate production behavior: [profiling](https://go.dev/doc/diagnostics#profiling). If we provide a profile to the compiler, it can make more informed decisions: more aggressively optimizing the most frequently used functions, or more accurately selecting common cases.

​	在没有关于代码在生产环境中如何使用的明确信息的情况下，编译器只能根据包的源代码进行操作。但是，我们确实有一种工具来评估生产行为：[性能分析](https://go.dev/doc/diagnostics#profiling)。如果我们向编译器提供一个性能分析文件，它可以做出更明智的决策：更积极地优化最常使用的函数，或更准确地选择常见的情况。

Using profiles of application behavior for compiler optimization is known as *Profile-Guided Optimization (PGO)* (also known as Feedback-Directed Optimization (FDO)).

​	使用应用程序行为的分析文件进行编译器优化称为*性能分析引导优化（PGO）*（也称为反馈导向优化（FDO））。

Go 1.20 includes initial support for PGO as a preview. See the [profile-guided optimization user guide](https://go.dev/doc/pgo) for complete documentation. There are still some rough edges that may prevent production use, but we would love for you to try it out and [send us any feedback or issues you encounter](https://go.dev/issue/new).

​	Go 1.20 版本包含了对 PGO 的初始支持，作为一项预览功能。完整的文档，请参阅[性能分析引导优化用户指南](https://go.dev/doc/pgo)。尽管还存在一些可能阻止在生产环境中使用的问题，但我们非常希望您尝试并[向我们发送您遇到的任何反馈或问题](https://go.dev/issue/new)。

## 示例 Example

Let’s build a service that converts Markdown to HTML: users upload Markdown source to `/render`, which returns the HTML conversion. We can use [`gitlab.com/golang-commonmark/markdown`](https://pkg.go.dev/gitlab.com/golang-commonmark/markdown) to implement this easily.

​	让我们构建一个将 Markdown 转换为 HTML 的服务：用户将 Markdown 源文件上传到 `/render`，该服务返回 HTML 转换结果。我们可以使用 [`gitlab.com/golang-commonmark/markdown`](https://pkg.go.dev/gitlab.com/golang-commonmark/markdown) 轻松实现这个功能。

### Set up

```
$ go mod init example.com/markdown
$ go get gitlab.com/golang-commonmark/markdown@bf3e522c626a
```

In `main.go`:

​	在 `main.go` 中：

```go
package main

import (
    "bytes"
    "io"
    "log"
    "net/http"
    _ "net/http/pprof"

    "gitlab.com/golang-commonmark/markdown"
)

func render(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return
    }

    src, err := io.ReadAll(r.Body)
    if err != nil {
        log.Printf("error reading body: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    md := markdown.New(
        markdown.XHTMLOutput(true),
        markdown.Typographer(true),
        markdown.Linkify(true),
        markdown.Tables(true),
    )

    var buf bytes.Buffer
    if err := md.Render(&buf, src); err != nil {
        log.Printf("error converting markdown: %v", err)
        http.Error(w, "Malformed markdown", http.StatusBadRequest)
        return
    }

    if _, err := io.Copy(w, &buf); err != nil {
        log.Printf("error writing response: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/render", render)
    log.Printf("Serving on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Build and run the server:

​	构建并运行服务器：

```sh
$ go build -o markdown.nopgo.exe
$ ./markdown.nopgo.exe
2023/01/19 14:26:24 Serving on port 8080...
```

Let’s try sending some Markdown from another terminal. We can use the README from the Go project as a sample document:

​	让我们尝试从另一个终端发送一些 Markdown 内容。我们可以使用 Go 项目的 README 作为示例文档：

```sh
$ curl -o README.md -L "https://raw.githubusercontent.com/golang/go/c16c2c49e2fa98ae551fc6335215fadd62d33542/README.md"
$ curl --data-binary @README.md http://localhost:8080/render
<h1>The Go Programming Language</h1>
<p>Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.</p>
...
```

### 性能分析 Profiling

Now that we have a working service, let’s collect a profile and rebuild with PGO to see if we get better performance.

​	现在我们有了一个工作中的服务，让我们收集一个性能分析文件，并使用 PGO 重新构建，看看是否可以获得更好的性能。

In `main.go`, we imported [net/http/pprof](https://pkg.go.dev/net/http/pprof) which automatically adds a `/debug/pprof/profile` endpoint to the server for fetching a CPU profile.

​	在 `main.go` 中，我们导入了 [net/http/pprof](https://pkg.go.dev/net/http/pprof)，它会自动将一个 `/debug/pprof/profile` 的端点添加到服务器，用于获取 CPU 分析。

Normally you want to collect a profile from your production environment so that the compiler gets a representative view of behavior in production. Since this example doesn’t have a "production" environment, we will create a simple program to generate load while we collect a profile. Copy the source of [this program](https://go.dev/play/p/yYH0kfsZcpL) to `load/main.go` and start the load generator (make sure the server is still running!).

​	通常，您希望从生产环境中收集性能分析，以便编译器获得在生产环境中行为的代表性视图。由于此示例没有“生产”环境，我们将创建一个简单的程序，在收集性能分析时生成负载。将[此程序的源代码](https://go.dev/play/p/yYH0kfsZcpL)复制到 `load/main.go` 并启动负载生成器（确保服务器仍在运行！）。

```sh
$ go run example.com/markdown/load
```

While that is running, download a profile from the server:

​	在运行时，从服务器下载一个性能分析文件：

```sh
$ curl -o cpu.pprof "http://localhost:8080/debug/pprof/profile?seconds=30"
```

Once this completes, kill the load generator and the server.

​	完成后，停止负载生成器和服务器。

### 使用性能分析 Using the profile

We can ask the Go toolchain to build with PGO using the `-pgo` flag to `go build`. `-pgo` takes either the path to the profile to use, or `auto`, which will use the `default.pgo` file in the main package directory.

​	我们可以使用 `-pgo` 标志来要求 Go 工具链使用 PGO 进行构建。`-pgo` 接受分析文件的路径，或者使用 `auto`，它将使用主包目录中的 `default.pgo` 文件。

We recommending commiting `default.pgo` profiles to your repository. Storing profiles alongside your source code ensures that users automatically have access to the profile simply by fetching the repository (either via the version control system, or via `go get`) and that builds remain reproducible. In Go 1.20, `-pgo=off` is the default, so users still need to add `-pgo=auto`, but a future version of Go is expected to change the default to `-pgo=auto`, automatically giving anyone that builds the binary the benefit of PGO.

​	我们建议将 `default.pgo` 分析文件提交到您的存储库中。将分析文件与源代码放在一起，可以确保用户仅需获取存储库（通过版本控制系统或 `go get`）即可自动访问分析文件，并且构建结果可复现。在 Go 1.20 中，`-pgo=off` 是默认值，因此用户仍然需要添加 `-pgo=auto`，但预计未来版本的 Go 将将默认值更改为 `-pgo=auto`，这样任何构建二进制文件的人都能够获得 PGO 的好处。

Let’s build:

​	让我们进行构建：

```sh
$ mv cpu.pprof default.pgo
$ go build -pgo=auto -o markdown.withpgo.exe
```

### 评估 Evaluation

We will use a Go benchmark version of the load generator to evaluate the effect of PGO on performance. Copy [this benchmark](https://go.dev/play/p/6FnQmHfRjbh) to `load/bench_test.go`.

​	我们将使用 Go 基准测试版本的负载生成器来评估 PGO 对性能的影响。将[此基准测试](https://go.dev/play/p/6FnQmHfRjbh)复制到 `load/bench_test.go`。

First, we will benchmark the server without PGO. Start that server:

​	首先，我们将在没有 PGO 的情况下对服务器进行基准测试。启动该服务器：

```sh
$ ./markdown.nopgo.exe
```

While that is running, run several benchmark iterations:

​	在运行时，运行多个基准测试迭代：

```sh
$ go test example.com/markdown/load -bench=. -count=20 -source ../README.md > nopgo.txt
```

Once that completes, kill the original server and start the version with PGO:

​	完成后，停止原始服务器并启动带有 PGO 的版本：

```sh
$ ./markdown.withpgo.exe
```

While that is running, run several benchmark iterations:

​	在运行时，运行多个基准测试迭代：

```sh
$ go test example.com/markdown/load -bench=. -count=20 -source ../README.md > withpgo.txt
```

Once that completes, let’s compare the results:

​	完成后，我们比较结果：

```sh
$ go install golang.org/x/perf/cmd/benchstat@latest
$ benchstat nopgo.txt withpgo.txt
goos: linux
goarch: amd64
pkg: example.com/markdown/load
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
        │  nopgo.txt  │            withpgo.txt             │
        │   sec/op    │   sec/op     vs base               │
Load-12   393.8µ ± 1%   383.6µ ± 1%  -2.59% (p=0.000 n=20)
```

The new version is around 2.6% faster! In Go 1.20, workloads typically get between 2% and 4% CPU usage improvements from enabling PGO. Profiles contain a wealth of information about application behavior and Go 1.20 just begins to crack the surface by using this information for inlining. Future releases will continue improving performance as more parts of the compiler take advantage of PGO.

​	新版本大约快了 2.6%！在 Go 1.20 中，启用 PGO 可以使工作负载的 CPU 使用率提高 2% 到 4%。分析文件包含了有关应用程序行为的丰富信息，Go 1.20 仅仅开始利用此信息进行内联优化。未来的版本将继续改进性能，让编译器的更多部分利用 PGO。

## 后续步骤 Next steps

In this example, after collecting a profile, we rebuilt our server using the exact same source code used in the original build. In a real-world scenario, there is always ongoing development. So we may collect a profile from production, which is running last week’s code, and use it to build with today’s source code. That is perfectly fine! PGO in Go can handle minor changes to source code without issue.

​	在这个示例中，我们在收集了性能分析文件之后，使用与原始构建中完全相同的源代码重新构建了服务器。在实际情况中，开发工作通常是持续进行的。因此，我们可能会从正在运行上周代码的生产环境中收集一个性能分析文件，并将其用于使用今天的源代码进行构建。这完全没有问题！Go 中的 PGO 可以处理源代码的小变化而无需处理问题。

For much more information on using PGO, best practices and caveats to be aware of, please see the [profile-guided optimization user guide](https://go.dev/doc/pgo).

​	有关使用 PGO 的更多信息、最佳实践和注意事项，请参阅[性能分析引导优化用户指南](https://go.dev/doc/pgo)。

Please send us your feedback! PGO is still in preview and we’d love to hear about anything that is difficult to use, doesn’t work correctly, etc. Please file issues at https://go.dev/issue/new.

​	请向我们发送您的反馈！PGO 仍处于预览阶段，我们很乐意听到任何难以使用、无法正常工作等问题。请在 https://go.dev/issue/new 提交问题。
