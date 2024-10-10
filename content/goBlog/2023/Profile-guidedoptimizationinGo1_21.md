+++
title = "Go 1.21中的基于配置文件的优化"
date = 2023-10-04T14:31:31+08:00
type = "docs"
weight = 83
description = ""
isCJKLanguage = true
draft = false
+++

# Profile-guided optimization in Go 1.21 - Go 1.21中的基于配置文件的优化

> 原文：[https://go.dev/blog/pgo](https://go.dev/blog/pgo)

Michael Pratt
5 September 2023

迈克尔·普拉特
2023年9月5日

Earlier in 2023, Go 1.20 [shipped a preview of profile-guided optimization (PGO)](https://go.dev/blog/pgo-preview) for users to test. After addressing known limitations in the preview, and with additional refinements thanks to community feedback and contributions, PGO support in Go 1.21 is ready for general production use! See the [profile-guided optimization user guide](https://go.dev/doc/pgo) for complete documentation.

​	在2023年早些时候，Go 1.20发布了[基于配置文件的优化（PGO）的预览版本](https://go.dev/blog/pgo-preview)供用户测试。在解决了预览版本中已知的限制并通过社区反馈和贡献进行了进一步改进后，Go 1.21中的PGO支持已经准备好供生产使用！请参阅[基于配置文件的优化用户指南](https://go.dev/doc/pgo)获取完整的文档。

[Below](https://go.dev/blog/pgo#example) we will run through an example of using PGO to improve the performance of an application. Before we get to that, what exactly is “profile-guided optimization”?

​	[下面](https://go.dev/blog/pgo#example)我们将通过一个示例来演示如何使用PGO来提高应用程序的性能。在我们开始之前，什么是“基于配置文件的优化”？

When you build a Go binary, the Go compiler performs optimizations to try to generate the best performing binary it can. For example, constant propagation can evaluate constant expressions at compile time, avoiding runtime evaluation cost. Escape analysis avoids heap allocations for locally-scoped objects, avoiding GC overheads. Inlining copies the body of simple functions into callers, often enabling further optimization in the caller (such as additional constant propagation or better escape analysis). Devirtualization converts indirect calls on interface values whose type can be determined statically into direct calls to the concrete method (which often enables inlining of the call).

​	当构建Go二进制文件时，Go编译器会进行优化，以生成性能最佳的二进制文件。例如，常量传播可以在编译时计算常量表达式，避免运行时的计算成本。逃逸分析可以避免为局部作用域对象分配堆空间，减少垃圾回收开销。内联将简单函数的主体复制到调用者中，通常可以在调用者中进一步优化（如额外的常量传播或更好的逃逸分析）。去虚拟化将对接口值的间接调用转换为对具体方法的直接调用（这通常可以启用调用的内联）。

Go improves optimizations from release to release, but doing so is no easy task. Some optimizations are tunable, but the compiler can’t just “turn it up to 11” on every optimization because overly aggressive optimizations can actually hurt performance or cause excessive build times. Other optimizations require the compiler to make a judgment call about what the “common” and “uncommon” paths in a function are. The compiler must make a best guess based on static heuristics because it can’t know which cases will be common at run time.

​	Go在每个版本中都改进了优化，但这并不容易。一些优化是可调的，但编译器不能在每个优化上都“开到最大”，因为过于激进的优化实际上可能会降低性能或导致构建时间过长。其他优化需要编译器对函数中的“常见”和“不常见”路径进行判断。编译器必须根据静态启发式规则进行最佳猜测，因为它无法知道哪些情况在运行时是常见的。

Or can it?

或者，它可以吗？

With no definitive information about how the code is used in a production environment, the compiler can operate only on the source code of packages. But we do have a tool to evaluate production behavior: [profiling](https://go.dev/doc/diagnostics#profiling). If we provide a profile to the compiler, it can make more informed decisions: more aggressively optimizing the most frequently used functions, or more accurately selecting common cases.

​	在没有关于代码在生产环境中如何使用的确切信息的情况下，编译器只能根据包的源代码进行操作。但是，我们确实有一种工具可以评估生产行为：[性能分析](https://go.dev/doc/diagnostics#profiling)。如果我们为编译器提供一个配置文件，它可以做出更明智的决策：更积极地优化最常用的函数，或更准确地选择常见情况。

Using profiles of application behavior for compiler optimization is known as *Profile-Guided Optimization (PGO)* (also known as Feedback-Directed Optimization (FDO)).

​	使用应用程序行为的配置文件进行编译器优化称为*基于配置文件的优化（PGO）*（也称为反馈导向优化（FDO））。

## 示例 Example

Let’s build a service that converts Markdown to HTML: users upload Markdown source to `/render`, which returns the HTML conversion. We can use [`gitlab.com/golang-commonmark/markdown`](https://pkg.go.dev/gitlab.com/golang-commonmark/markdown) to implement this easily.

​	让我们构建一个将Markdown转换为HTML的服务：用户将Markdown源码上传到 `/render` ，然后返回HTML转换结果。我们可以使用[ `gitlab.com/golang-commonmark/markdown` ](https://pkg.go.dev/gitlab.com/golang-commonmark/markdown)来轻松实现这一点。

### 设置 Set up

```
$ go mod init example.com/markdown
$ go get gitlab.com/golang-commonmark/markdown@bf3e522c626a
```

In `main.go`:

​	在 `main.go` 中：

```
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

```
$ go build -o markdown.nopgo.exe
$ ./markdown.nopgo.exe
2023/08/23 03:55:51 Serving on port 8080...
```

Let’s try sending some Markdown from another terminal. We can use the `README.md` from the Go project as a sample document:

​	让我们尝试从另一个终端发送一些Markdown。我们可以使用Go项目的 `README.md` 作为示例文档：

```
$ curl -o README.md -L "https://raw.githubusercontent.com/golang/go/c16c2c49e2fa98ae551fc6335215fadd62d33542/README.md"
$ curl --data-binary @README.md http://localhost:8080/render
<h1>The Go Programming Language</h1>
<p>Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.</p>
...
```

### 性能分析 Profiling

Now that we have a working service, let’s collect a profile and rebuild with PGO to see if we get better performance.

​	现在我们有了一个工作的服务，让我们收集一个配置文件并使用PGO重新构建，看看是否能获得更好的性能。

In `main.go`, we imported [net/http/pprof](https://pkg.go.dev/net/http/pprof) which automatically adds a `/debug/pprof/profile` endpoint to the server for fetching a CPU profile.

​	在 `main.go` 中，我们导入了[net/http/pprof](https://pkg.go.dev/net/http/pprof)，它会自动在服务器上添加一个 `/debug/pprof/profile` 端点，用于获取CPU配置文件。

Normally you want to collect a profile from your production environment so that the compiler gets a representative view of behavior in production. Since this example doesn’t have a “production” environment, I have created a [simple program](https://github.com/prattmic/markdown-pgo/blob/main/load/main.go) to generate load while we collect a profile. Fetch and start the load generator (make sure the server is still running!):

​	通常情况下，您希望从生产环境中收集配置文件，以便编译器能够获得生产环境中行为的代表性视图。由于此示例没有“生产”环境，我创建了一个[简单的程序](https://github.com/prattmic/markdown-pgo/blob/main/load/main.go)，用于在我们收集配置文件时生成负载。下载并启动负载生成器（确保服务器仍在运行！）：

```
$ go run github.com/prattmic/markdown-pgo/load@latest
```

While that is running, download a profile from the server:

​	在运行时，从服务器下载一个配置文件：

```
$ curl -o cpu.pprof "http://localhost:8080/debug/pprof/profile?seconds=30"
```

Once this completes, kill the load generator and the server.

​	完成后，停止负载生成器和服务器。

### 使用配置文件 Using the profile

The Go toolchain will automatically enable PGO when it finds a profile named `default.pgo` in the main package directory. Alternatively, the `-pgo` flag to `go build` takes a path to a profile to use for PGO.

​	当Go工具链在主包目录中找到一个名为 `default.pgo` 的配置文件时，它将自动启用PGO。或者， `go build` 的 `-pgo` 标志可以接受一个配置文件的路径来用于PGO。

We recommend committing `default.pgo` files to your repository. Storing profiles alongside your source code ensures that users automatically have access to the profile simply by fetching the repository (either via the version control system, or via `go get`) and that builds remain reproducible.

​	我们建议将 `default.pgo` 文件提交到您的代码库中。将配置文件与源代码存储在一起可以确保用户仅通过获取代码库（无论是通过版本控制系统还是通过 `go get` ）就可以自动访问配置文件，并且构建仍然具有可重复性。

Let’s build:

​	让我们进行构建：

```
$ mv cpu.pprof default.pgo
$ go build -o markdown.withpgo.exe
```

We can check that PGO was enabled in the build with `go version`:

​	我们可以使用 `go version` 检查构建中是否启用了PGO：

```
$ go version -m markdown.withpgo.exe
./markdown.withpgo.exe: go1.21.0
...
        build   -pgo=/tmp/pgo121/default.pgo
```

### 评估 Evaluation

We will use a Go benchmark [version of the load generator](https://github.com/prattmic/markdown-pgo/blob/main/load/bench_test.go) to evaluate the effect of PGO on performance.

​	我们将使用一个Go基准测试的[负载生成器版本](https://github.com/prattmic/markdown-pgo/blob/main/load/bench_test.go)来评估PGO对性能的影响。

First, we will benchmark the server without PGO. Start that server:

​	首先，我们将对没有PGO的服务器进行基准测试。启动该服务器：

```
$ ./markdown.nopgo.exe
```

While that is running, run several benchmark iterations:

​	在服务器运行时，运行多次基准测试迭代：

```
$ go get github.com/prattmic/markdown-pgo@latest
$ go test github.com/prattmic/markdown-pgo/load -bench=. -count=40 -source $(pwd)/README.md > nopgo.txt
```

Once that completes, kill the original server and start the version with PGO:

​	完成后，停止原始服务器并启动带有PGO的版本：

```
$ ./markdown.withpgo.exe
```

While that is running, run several benchmark iterations:

​	在服务器运行时，运行多次基准测试迭代：

```
$ go test github.com/prattmic/markdown-pgo/load -bench=. -count=40 -source $(pwd)/README.md > withpgo.txt
```

Once that completes, let’s compare the results:

​	完成后，让我们比较结果：

```
$ go install golang.org/x/perf/cmd/benchstat@latest
$ benchstat nopgo.txt withpgo.txt
goos: linux
goarch: amd64
pkg: github.com/prattmic/markdown-pgo/load
cpu: Intel(R) Xeon(R) W-2135 CPU @ 3.70GHz
        │  nopgo.txt  │            withpgo.txt             │
        │   sec/op    │   sec/op     vs base               │
Load-12   374.5µ ± 1%   360.2µ ± 0%  -3.83% (p=0.000 n=40)
```

The new version is around 3.8% faster! In Go 1.21, workloads typically get between 2% and 7% CPU usage improvements from enabling PGO. Profiles contain a wealth of information about application behavior and Go 1.21 just begins to crack the surface by using this information for a limited set of optimizations. Future releases will continue improving performance as more parts of the compiler take advantage of PGO.

​	新版本快了约3.8％！在Go 1.21中，通过启用PGO，工作负载的CPU使用率通常会提高2％到7％。配置文件包含了关于应用程序行为的丰富信息，而Go 1.21仅仅开始利用这些信息进行一组有限的优化。随着编译器的更多部分利用PGO，未来的版本将继续改进性能。

## 下一步 Next steps

In this example, after collecting a profile, we rebuilt our server using the exact same source code used in the original build. In a real-world scenario, there is always ongoing development. So we may collect a profile from production, which is running last week’s code, and use it to build with today’s source code. That is perfectly fine! PGO in Go can handle minor changes to source code without issue. Of course, over time source code will drift more and more, so it is still important to update the profile occasionally.

​	在这个示例中，我们在收集配置文件之后，使用与原始构建中完全相同的源代码重新构建了服务器。在实际情况中，开发工作是持续进行的。因此，我们可能会从正在运行上周代码的生产环境中收集配置文件，并将其用于今天的源代码构建。这是完全可以的！Go中的PGO可以处理源代码的轻微更改而不会出现问题。当然，随着时间的推移，源代码会越来越偏离，因此定期更新配置文件仍然很重要。

For much more information on using PGO, best practices and caveats to be aware of, please see the [profile-guided optimization user guide](https://go.dev/doc/pgo). If you are curious about what is going on under the hood, keep reading!

​	关于使用PGO的更多信息，最佳实践和注意事项，请参阅[基于配置文件的优化用户指南](https://go.dev/doc/pgo)。如果您对内部工作原理感兴趣，请继续阅读！

## 内部工作原理 Under the hood

To get a better understanding of what made this application faster, let’s take a look under the hood to see how performance has changed. We are going to take a look at two different PGO-driven optimizations.

​	为了更好地理解这个应用程序为什么变得更快，让我们来看看内部是如何改变性能的。我们将观察两种不同的基于配置文件的优化。

### 内联 Inlining

To observe inlining improvements, let’s analyze this markdown application both with and without PGO.

​	为了观察内联改进，让我们同时分析使用PGO和不使用PGO的Markdown应用程序。

I will compare this using a technique called differential profiling, where we collect two profiles (one with PGO and one without) and compare them. For differential profiling, it’s important that both profiles represent the same amount of **work**, not the same amount of time, so I’ve adjusted the server to automatically collect profiles, and the load generator to send a fixed number of requests and then exit the server.

​	我将使用一种称为差分分析的技术来进行比较，其中我们收集两个配置文件（一个带有PGO，一个没有PGO）并进行比较。对于差分分析，重要的是两个配置文件代表相同数量的**工作**，而不是相同数量的时间，因此我调整了服务器以自动收集配置文件，并调整了负载生成器以发送固定数量的请求，然后退出服务器。

The changes I have made to the server as well as the profiles collected can be found at https://github.com/prattmic/markdown-pgo. The load generator was run with `-count=300000 -quit`.

​	我对服务器进行的更改以及收集的配置文件可以在https://github.com/prattmic/markdown-pgo找到。负载生成器使用 `-count=300000 -quit` 运行。

As a quick consistency check, let’s take a look at the total CPU time required to handle all 300k requests:

​	作为快速一致性检查，让我们看一下处理所有300k请求所需的总CPU时间：

```
$ go tool pprof -top cpu.nopgo.pprof | grep "Total samples"
Duration: 116.92s, Total samples = 118.73s (101.55%)
$ go tool pprof -top cpu.withpgo.pprof | grep "Total samples"
Duration: 113.91s, Total samples = 115.03s (100.99%)
```

CPU time dropped from ~118s to ~115s, or about 3%. This is in line with our benchmark results, which is a good sign that these profiles are representative.

​	CPU时间从约118秒降低到约115秒，约降低了3％。这与我们的基准测试结果一致，这是这些配置文件具有代表性的良好迹象。

Now we can open a differential profile to look for savings:

​	现在，我们可以打开一个差分配置文件来查找节省：

```
$ go tool pprof -diff_base cpu.nopgo.pprof cpu.withpgo.pprof
File: markdown.profile.withpgo.exe
Type: cpu
Time: Aug 28, 2023 at 10:26pm (EDT)
Duration: 230.82s, Total samples = 118.73s (51.44%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top -cum
Showing nodes accounting for -0.10s, 0.084% of 118.73s total
Dropped 268 nodes (cum <= 0.59s)
Showing top 10 nodes out of 668
      flat  flat%   sum%        cum   cum%
    -0.03s 0.025% 0.025%     -2.56s  2.16%  gitlab.com/golang-commonmark/markdown.ruleLinkify
     0.04s 0.034% 0.0084%     -2.19s  1.84%  net/http.(*conn).serve
     0.02s 0.017% 0.025%     -1.82s  1.53%  gitlab.com/golang-commonmark/markdown.(*Markdown).Render
     0.02s 0.017% 0.042%     -1.80s  1.52%  gitlab.com/golang-commonmark/markdown.(*Markdown).Parse
    -0.03s 0.025% 0.017%     -1.71s  1.44%  runtime.mallocgc
    -0.07s 0.059% 0.042%     -1.62s  1.36%  net/http.(*ServeMux).ServeHTTP
     0.04s 0.034% 0.0084%     -1.58s  1.33%  net/http.serverHandler.ServeHTTP
    -0.01s 0.0084% 0.017%     -1.57s  1.32%  main.render
     0.01s 0.0084% 0.0084%     -1.56s  1.31%  net/http.HandlerFunc.ServeHTTP
    -0.09s 0.076% 0.084%     -1.25s  1.05%  runtime.newobject
(pprof) top
Showing nodes accounting for -1.41s, 1.19% of 118.73s total
Dropped 268 nodes (cum <= 0.59s)
Showing top 10 nodes out of 668
      flat  flat%   sum%        cum   cum%
    -0.46s  0.39%  0.39%     -0.91s  0.77%  runtime.scanobject
    -0.40s  0.34%  0.72%     -0.40s  0.34%  runtime.nextFreeFast (inline)
     0.36s   0.3%  0.42%      0.36s   0.3%  gitlab.com/golang-commonmark/markdown.performReplacements
    -0.35s  0.29%  0.72%     -0.37s  0.31%  runtime.writeHeapBits.flush
     0.32s  0.27%  0.45%      0.67s  0.56%  gitlab.com/golang-commonmark/markdown.ruleReplacements
    -0.31s  0.26%  0.71%     -0.29s  0.24%  runtime.writeHeapBits.write
    -0.30s  0.25%  0.96%     -0.37s  0.31%  runtime.deductAssistCredit
     0.29s  0.24%  0.72%      0.10s 0.084%  gitlab.com/golang-commonmark/markdown.ruleText
    -0.29s  0.24%  0.96%     -0.29s  0.24%  runtime.(*mspan).base (inline)
    -0.27s  0.23%  1.19%     -0.42s  0.35%  bytes.(*Buffer).WriteRune
```

When specifying `pprof -diff_base`, the values in displayed in pprof are the *difference* between the two profiles. So, for instance, `runtime.scanobject` used 0.46s less CPU time with PGO than without. On the other hand, `gitlab.com/golang-commonmark/markdown.performReplacements` used 0.36s more CPU time. In a differential profile, we typically want to look at the absolute values (`flat` and `cum` columns), as the percentages aren’t meaningful.

​	在指定 `pprof -diff_base` 时，pprof中显示的值是两个配置文件之间的*差异*。因此，例如， `runtime.scanobject` 使用的CPU时间比没有PGO时少了0.46秒。另一方面， `gitlab.com/golang-commonmark/markdown.performReplacements` 使用的CPU时间增加了0.36秒。在差分配置文件中，我们通常希望查看绝对值（ `flat` 和 `cum` 列），因为百分比没有意义。

`top -cum` shows the top differences by cumulative change. That is, the difference in CPU of a function and all transitive callees from that function. This will generally show the outermost frames in our program’s call graph, such as `main` or another goroutine entry point. Here we can see most savings are coming from the `ruleLinkify` portion of handling HTTP requests.

​	 `top -cum` 按累积更改显示前几个差异。也就是说，函数及其所有传递的调用者的CPU差异。这通常显示我们程序调用图中的最外层框架，如 `main` 或另一个goroutine的入口点。在这里，我们可以看到大部分节省来自处理HTTP请求的 `ruleLinkify` 部分。

`top` shows the top differences limited only to changes in the function itself. This will generally show inner frames in our program’s call graph, where most of the actual work is happening. Here we can see that individual savings are coming mostly from `runtime` functions.

​	 `top` 仅显示函数本身的差异，限制在程序调用图中的内部框架，大部分实际工作发生在这里。在这里，我们可以看到个别节省主要来自 `runtime` 函数。

What are those? Let’s peek up the call stack to see where they come from:

​	这些是什么？让我们向上查看调用堆栈，看看它们来自哪里：

```
(pprof) peek scanobject$
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                            -0.86s 94.51% |   runtime.gcDrain
                                            -0.09s  9.89% |   runtime.gcDrainN
                                             0.04s  4.40% |   runtime.markrootSpans
    -0.46s  0.39%  0.39%     -0.91s  0.77%                | runtime.scanobject
                                            -0.19s 20.88% |   runtime.greyobject
                                            -0.13s 14.29% |   runtime.heapBits.nextFast (inline)
                                            -0.08s  8.79% |   runtime.heapBits.next
                                            -0.08s  8.79% |   runtime.spanOfUnchecked (inline)
                                             0.04s  4.40% |   runtime.heapBitsForAddr
                                            -0.01s  1.10% |   runtime.findObject
----------------------------------------------------------+-------------
(pprof) peek gcDrain$
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                               -1s   100% |   runtime.gcBgMarkWorker.func2
     0.15s  0.13%  0.13%        -1s  0.84%                | runtime.gcDrain
                                            -0.86s 86.00% |   runtime.scanobject
                                            -0.18s 18.00% |   runtime.(*gcWork).balance
                                            -0.11s 11.00% |   runtime.(*gcWork).tryGet
                                             0.09s  9.00% |   runtime.pollWork
                                            -0.03s  3.00% |   runtime.(*gcWork).tryGetFast (inline)
                                            -0.03s  3.00% |   runtime.markroot
                                            -0.02s  2.00% |   runtime.wbBufFlush
                                             0.01s  1.00% |   runtime/internal/atomic.(*Bool).Load (inline)
                                            -0.01s  1.00% |   runtime.gcFlushBgCredit
                                            -0.01s  1.00% |   runtime/internal/atomic.(*Int64).Add (inline)
----------------------------------------------------------+-------------
```

So `runtime.scanobject` is ultimately coming from `runtime.gcBgMarkWorker`. The [Go GC Guide](https://go.dev/doc/gc-guide#Identiying_costs) tells us that `runtime.gcBgMarkWorker` is part of the garbage collector, so `runtime.scanobject` savings must be GC savings. What about `nextFreeFast` and other `runtime` functions?

​	因此， `runtime.scanobject` 最终来自 `runtime.gcBgMarkWorker` 。[Go GC指南](https://go.dev/doc/gc-guide#Identiying_costs)告诉我们， `runtime.gcBgMarkWorker` 是垃圾回收器的一部分，因此 `runtime.scanobject` 的节省必须是GC的节省。 `nextFreeFast` 和其他 `runtime` 函数呢？

```
(pprof) peek nextFreeFast$
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                            -0.40s   100% |   runtime.mallocgc (inline)
    -0.40s  0.34%  0.34%     -0.40s  0.34%                | runtime.nextFreeFast
----------------------------------------------------------+-------------
(pprof) peek writeHeapBits
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                            -0.37s   100% |   runtime.heapBitsSetType
                                                 0     0% |   runtime.(*mspan).initHeapBits
    -0.35s  0.29%  0.29%     -0.37s  0.31%                | runtime.writeHeapBits.flush
                                            -0.02s  5.41% |   runtime.arenaIndex (inline)
----------------------------------------------------------+-------------
                                            -0.29s   100% |   runtime.heapBitsSetType
    -0.31s  0.26%  0.56%     -0.29s  0.24%                | runtime.writeHeapBits.write
                                             0.02s  6.90% |   runtime.arenaIndex (inline)
----------------------------------------------------------+-------------
(pprof) peek heapBitsSetType$
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                            -0.82s   100% |   runtime.mallocgc
    -0.12s   0.1%   0.1%     -0.82s  0.69%                | runtime.heapBitsSetType
                                            -0.37s 45.12% |   runtime.writeHeapBits.flush
                                            -0.29s 35.37% |   runtime.writeHeapBits.write
                                            -0.03s  3.66% |   runtime.readUintptr (inline)
                                            -0.01s  1.22% |   runtime.writeHeapBitsForAddr (inline)
----------------------------------------------------------+-------------
(pprof) peek deductAssistCredit$
Showing nodes accounting for -3.72s, 3.13% of 118.73s total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                            -0.37s   100% |   runtime.mallocgc
    -0.30s  0.25%  0.25%     -0.37s  0.31%                | runtime.deductAssistCredit
                                            -0.07s 18.92% |   runtime.gcAssistAlloc
----------------------------------------------------------+-------------
```

Looks like `nextFreeFast` and some of the others in the top 10 are ultimately coming from `runtime.mallocgc`, which the GC Guide tells us is the memory allocator.

​	在前10名中， `nextFreeFast` 和其他一些函数最终都是来自 `runtime.mallocgc` ，这是内存分配器。

Reduced costs in the GC and allocator imply that we are allocating less overall. Let’s take a look at the heap profiles for insight:

​	GC和分配器的成本降低意味着我们总体上分配的内存较少。让我们看一下堆剖析以获取更多信息：

```
$ go tool pprof -sample_index=alloc_objects -diff_base heap.nopgo.pprof heap.withpgo.pprof
File: markdown.profile.withpgo.exe
Type: alloc_objects
Time: Aug 28, 2023 at 10:28pm (EDT)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for -12044903, 8.29% of 145309950 total
Dropped 60 nodes (cum <= 726549)
Showing top 10 nodes out of 58
      flat  flat%   sum%        cum   cum%
  -4974135  3.42%  3.42%   -4974135  3.42%  gitlab.com/golang-commonmark/mdurl.Parse
  -4249044  2.92%  6.35%   -4249044  2.92%  gitlab.com/golang-commonmark/mdurl.(*URL).String
   -901135  0.62%  6.97%    -977596  0.67%  gitlab.com/golang-commonmark/puny.mapLabels
   -653998  0.45%  7.42%    -482491  0.33%  gitlab.com/golang-commonmark/markdown.(*StateInline).PushPending
   -557073  0.38%  7.80%    -557073  0.38%  gitlab.com/golang-commonmark/linkify.Links
   -557073  0.38%  8.18%    -557073  0.38%  strings.genSplit
   -436919   0.3%  8.48%    -232152  0.16%  gitlab.com/golang-commonmark/markdown.(*StateBlock).Lines
   -408617  0.28%  8.77%    -408617  0.28%  net/textproto.readMIMEHeader
    401432  0.28%  8.49%     499610  0.34%  bytes.(*Buffer).grow
    291659   0.2%  8.29%     291659   0.2%  bytes.(*Buffer).String (inline)
```

The `-sample_index=alloc_objects` option is showing us the count of allocations, regardless of size. This is useful since we are investigating a decrease in CPU usage, which tends to correlate more with allocation count rather than size. There are quite a few reductions here, but let’s focus on the biggest reduction, `mdurl.Parse`.

​	`-sample_index=alloc_objects` 选项显示了分配的数量，而不考虑大小。这很有用，因为我们正在研究CPU使用率的降低，而这往往与分配数量相关，而不是与大小相关。这里有很多减少，但让我们专注于最大的减少，即 `mdurl.Parse` 。

For reference, let’s look at the total allocation counts for this function without PGO:

​	为了参考，让我们看一下没有PGO的情况下该函数的总分配数量：

```
$ go tool pprof -sample_index=alloc_objects -top heap.nopgo.pprof | grep mdurl.Parse
   4974135  3.42% 68.60%    4974135  3.42%  gitlab.com/golang-commonmark/mdurl.Parse
```

The total count before was 4974135, meaning that `mdurl.Parse` has eliminated 100% of allocations!

​	之前的总数是4974135，这意味着 `mdurl.Parse` 消除了100%的分配！

Back in the differential profile, let’s gather a bit more context:

​	回到差异剖面，让我们收集一些更多的上下文：

```
(pprof) peek mdurl.Parse
Showing nodes accounting for -12257184, 8.44% of 145309950 total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context
----------------------------------------------------------+-------------
                                          -2956806 59.44% |   gitlab.com/golang-commonmark/markdown.normalizeLink
                                          -2017329 40.56% |   gitlab.com/golang-commonmark/markdown.normalizeLinkText
  -4974135  3.42%  3.42%   -4974135  3.42%                | gitlab.com/golang-commonmark/mdurl.Parse
----------------------------------------------------------+-------------
```

The calls to `mdurl.Parse` are coming from `markdown.normalizeLink` and `markdown.normalizeLinkText`.

​	对 `mdurl.Parse` 的调用来自 `markdown.normalizeLink` 和 `markdown.normalizeLinkText` 。

```
(pprof) list mdurl.Parse
Total: 145309950
ROUTINE ======================== gitlab.com/golang-commonmark/mdurl.Parse in /usr/local/google/home/mpratt/go/pkg/mod/gitlab.com/golang-commonmark/mdurl@v0.0.0-20191124015652-932350d1cb84/parse
.go
  -4974135   -4974135 (flat, cum)  3.42% of Total
         .          .     60:func Parse(rawurl string) (*URL, error) {
         .          .     61:   n, err := findScheme(rawurl)
         .          .     62:   if err != nil {
         .          .     63:           return nil, err
         .          .     64:   }
         .          .     65:
  -4974135   -4974135     66:   var url URL
         .          .     67:   rest := rawurl
         .          .     68:   hostless := false
         .          .     69:   if n > 0 {
         .          .     70:           url.RawScheme = rest[:n]
         .          .     71:           url.Scheme, rest = strings.ToLower(rest[:n]), rest[n+1:]
```

Full source for these functions and callers can be found at:

​	这些函数和调用者的完整源代码可以在以下位置找到：

- [`mdurl.Parse`](https://gitlab.com/golang-commonmark/mdurl/-/blob/bd573caec3d827ead19e40b1f141a3802d956710/parse.go#L60)
- [`markdown.normalizeLink`](https://gitlab.com/golang-commonmark/markdown/-/blob/fd7971701a0cab12e9347109a4c889f5c0a1a479/util.go#L53)
- [`markdown.normalizeLinkText`](https://gitlab.com/golang-commonmark/markdown/-/blob/fd7971701a0cab12e9347109a4c889f5c0a1a479/util.go#L68)

So what happened here? In a non-PGO build, `mdurl.Parse` is considered too large to be eligible for inlining. However, because our PGO profile indicated that the calls to this function were hot, the compiler did inline them. We can see this from the “(inline)” annotation in the profiles:

​	那么这里发生了什么？在非PGO构建中， `mdurl.Parse` 被认为太大而无法进行内联。然而，因为我们的PGO剖面表明对这个函数的调用是热点，编译器对它们进行了内联。我们可以从剖面中的“(inline)”注释中看到这一点：

```
$ go tool pprof -top cpu.nopgo.pprof | grep mdurl.Parse
     0.36s   0.3% 63.76%      2.75s  2.32%  gitlab.com/golang-commonmark/mdurl.Parse
$ go tool pprof -top cpu.withpgo.pprof | grep mdurl.Parse
     0.55s  0.48% 58.12%      2.03s  1.76%  gitlab.com/golang-commonmark/mdurl.Parse (inline)
```

`mdurl.Parse` creates a `URL` as a local variable on line 66 (`var url URL`), and then returns a pointer to that variable on line 145 (`return &url, nil`). Normally this requires the variable to be allocated on the heap, as a reference to it lives beyond function return. However, once `mdurl.Parse` is inlined into `markdown.normalizeLink`, the compiler can observe that the variable does not escape `normalizeLink`, which allows the compiler to allocate it on the stack. `markdown.normalizeLinkText` is similar to `markdown.normalizeLink`.

​	`mdurl.Parse` 在第66行创建了一个 `URL` 作为局部变量（ `var url URL` ），然后在第145行返回了对该变量的指针（ `return &url, nil` ）。通常情况下，这需要在堆上分配变量，因为对它的引用超出了函数返回。然而，一旦 `mdurl.Parse` 内联到 `markdown.normalizeLink` 中，编译器可以观察到该变量不会逃逸到 `normalizeLink` 之外，这允许编译器将其分配在堆栈上。 `markdown.normalizeLinkText` 类似于 `markdown.normalizeLink` 。

The second largest reduction shown in the profile, from `mdurl.(*URL).String` is a similar case of eliminating an escape after inlining.

​	在配置文件中显示的第二大减少，来自 `mdurl.(*URL).String` ，是类似的情况，在内联之后消除了逃逸。

In these cases, we got improved performance through fewer heap allocations. Part of the power of PGO and compiler optimizations in general is that effects on allocations are not part of the compiler’s PGO implementation at all. The only change that PGO made was to allow inlining of these hot function calls. All of the effects to escape analysis and heap allocation were standard optimizations that apply to any build. Improved escape behavior is a great downstream effect of inlining, but it is not the only effect. Many optimizations can take advantage of inlining. For example, constant propagation may be able to simplify the code in a function after inlining when some of the inputs are constants.

​	在这些情况下，通过减少堆分配来提高性能。PGO和编译器优化的部分优势在于对堆分配的影响根本不是编译器的PGO实现的一部分。PGO唯一的改变是允许内联这些热点函数调用。对逃逸分析和堆分配的所有影响都是适用于任何构建的标准优化。改进的逃逸行为是内联的一个很好的下游效果，但它不是唯一的效果。许多优化可以利用内联。例如，常量传播可以在内联后简化函数中的代码，当一些输入是常量时。

### 虚函数调用优化 Devirtualization

In addition to inling, which we saw in the example above, PGO can also drive conditional devirtualization of interface calls.

​	除了我们在上面的示例中看到的内联之外，PGO还可以驱动接口调用的条件虚函数调用优化。

Before getting to PGO-driven devirtualization, let’s step back and define “devirtualization” in general. Suppose you have code that looks like something like this:

​	在介绍PGO驱动的虚函数调用优化之前，让我们先回顾一下一般情况下的“虚函数调用优化”的定义。假设你有如下代码：

```
f, _ := os.Open("foo.txt")
var r io.Reader = f
r.Read(b)
```

Here we have a call to the `io.Reader` interface method `Read`. Since interfaces can have multiple implementations, the compiler generates an *indirect* function call, meaning it looks up the correct method to call at run time from the type in the interface value. Indirect calls have a small additional runtime cost compared to direct calls, but more importantly they preclude some compiler optimizations. For example, the compiler can’t perform escape analysis on an indirect call since it doesn’t know the concrete method implementation.

​	这里我们调用了 `io.Reader` 接口的方法 `Read` 。由于接口可以有多个实现，编译器会生成一个*间接*函数调用，这意味着它在运行时从接口值的类型中查找要调用的正确方法。间接调用与直接调用相比有一些额外的运行时开销，但更重要的是它排除了一些编译器优化。例如，编译器无法对间接调用执行逃逸分析，因为它不知道具体的方法实现。

But in the example above, we *do* know the concrete method implementation. It must be `os.(*File).Read`, since `*os.File` is the only type that could possibly be assigned to `r`. In this case, the compiler will perform *devirtualization*, where it replaces the indirect call to `io.Reader.Read` with a direct call to `os.(*File).Read`, thus allowing other optimizations.

​	但在上面的示例中，我们确实知道具体的方法实现。它必须是 `os.(*File).Read` ，因为 `*os.File` 是唯一可能被赋给 `r` 的类型。在这种情况下，编译器将执行*虚函数调用优化*，它将间接调用 `io.Reader.Read` 替换为直接调用 `os.(*File).Read` ，从而允许其他优化。

(You are probably thinking “that code is useless, why would anyone write it that way?” This is a good point, but note that code like above could be the result of inlining. Suppose `f` is passed into a function that takes an `io.Reader` argument. Once the function is inlined, now the `io.Reader` becomes concrete.)

（您可能会想，“这段代码是无用的，为什么有人会这样写？”这是一个很好的问题，但请注意，像上面的代码可能是内联的结果。假设 `f` 被传递给一个接受 `io.Reader` 参数的函数。一旦函数被内联，现在 `io.Reader` 变为具体类型。）

PGO-driven devirtualization extends this concept to situations where the concrete type is not statically known, but profiling can show that, for example, an `io.Reader.Read` call targets `os.(*File).Read` most of the time. In this case, PGO can replace `r.Read(b)` with something like:

​	PGO驱动的虚函数调用优化将这个概念扩展到静态未知具体类型的情况，但是分析显示，例如， `io.Reader.Read` 调用大部分时间都是针对 `os.(*File).Read` 的。在这种情况下，PGO可以将 `r.Read(b)` 替换为类似以下的代码：

```
if f, ok := r.(*os.File); ok {
    f.Read(b)
} else {
    r.Read(b)
}
```

That is, we add a runtime check for the concrete type that is most likely to appear, and if so use a concrete call, or otherwise fall back to the standard indirect call. The advantage here is that the common path (using `*os.File`) can be inlined and have additional optimizations applied, but we still maintain a fallback path because a profile is not a guarantee that this will always be the case.

​	也就是说，我们添加了一个运行时检查，以确定最有可能的具体类型，并在是的情况下使用具体调用，否则使用标准的间接调用。优势在于常见路径（使用 `*os.File` ）可以内联并应用其他优化，但我们仍然保留了备用路径，因为剖面不能保证这种情况总是发生。

In our analysis of the markdown server we didn’t see PGO-driven devirtualization, but we also only looked at the top impacted areas. PGO (and most compiler optimizations) generally yield their benefit in the aggregate of very small improvements in lots of different places, so there is likely more happening than just what we looked at.

​	在我们对markdown服务器的分析中，我们没有看到PGO驱动的虚函数调用优化，但我们只看了受影响最大的区域。PGO（和大多数编译器优化）通常通过在许多不同地方实现非常小的改进来产生效益，因此可能会发生更多的情况。

Inlining and devirtualization are the two PGO-driven optimizations available in Go 1.21, but as we’ve seen, these often unlock additional optimizations. In addition, future versions of Go will continue to improve PGO with additional optimizations.

​	内联和虚函数调用优化是Go 1.21中提供的两种PGO驱动的优化，但正如我们所看到的，这些优化通常会解锁其他优化。此外，未来的Go版本将继续改进PGO，并提供其他优化。