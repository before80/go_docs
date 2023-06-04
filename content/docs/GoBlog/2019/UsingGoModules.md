+++
title = "使用 go 模块"
weight = 17
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Using Go Modules - 使用 go 模块

https://go.dev/blog/using-go-modules

Tyler Bui-Palsulich and Eno Compton
19 March 2019

2019年3月19日

## Introduction 简介

这篇文章是系列的第一部分。

- Part 1 — 使用Go模块（本帖）
- Part 2 — [Migrating To Go Modules 迁移到Go模块](../MigratingToGoModules)
- Part 3 — [Publishing Go Modules 发布Go模块](../PublishingGoModules)
- Part 4 — [Go Modules: v2 and Beyond Go模块：V2版及以后](../GoModulesV2AndBeyond)
- Part 5 — [Keeping Your Modules Compatible 保持模块的兼容性](../../2020/KeepingYourModulesCompatible)

注意：关于管理模块的依赖项的文档，请参见[管理依赖项](../../../UsingAndUnderstandingGo/ManagingDependencies)。

​	Go 1.11 和 1.12 包括[对模块的初步支持](https://go.dev/doc/go1.11#modules)，Go的[新依赖项管理系统]()使依赖项版本信息明确且更易于管理。这篇博客文章介绍了开始使用模块所需的基本操作。

​	模块是存储在文件树中的Go包的集合，其根部带有`go.mod`文件。`go.mod`文件定义了模块的模块路径(也是用于根目录的导入路径)及其依赖项需求(这些模块是成功构建所需的其他模块)。每个依赖需求都被写成一个模块路径和一个特定的[semantic version（语义版本）](http://semver.org/)。

​	从Go 1.11开始，如果当前目录或任何父目录有`go.mod`时，只要该目录在`$GOPATH/src`之外，就可以用`go`命令使用模块。（在`$GOPATH/src`之内，为了兼容，`go`命令仍然以旧的`GOPATH`模式运行，即使发现有`go.mod`。详情见[go命令文档](../../../References/CommandDocumentation/go)）。从Go 1.13开始，模块模式将成为所有开发的默认模式。

​	本篇文章讲述了使用模块开发Go代码时出现的一系列常见操作：

- 创建一个新模块
- 添加一个依赖项。
- 升级依赖项。
- 在新的主版本上添加一个依赖项。
- 将一个依赖项升级到一个新的主版本。
- 移除未使用的依赖项。

## Creating a new module 创建一个新模块

​	让我们创建一个新的模块。

​	在`$GOPATH/src`之外的某个地方创建一个新的空目录，`cd`进入该目录，然后创建一个新的源代码文件：`hello.go`：

```go linenums="1"
package hello

func Hello() string {
    return "Hello, world."
}
```

​	我们也来写一个测试，放在`hello_test.go`中：

```go linenums="1"
package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
```

At this point, the directory contains a package, but not a module, because there is no `go.mod` file. If we were working in `/home/gopher/hello` and ran `go test` now, we’d see:

​	在这一点上，该目录包含一个包，但不是一个模块，因为没有go.mod文件。如果我们在/home/gopher/hello工作，现在运行go test，我们会看到：

```shell linenums="1"
$ go test
PASS
ok      _/home/gopher/hello 0.020s
$
```

The last line summarizes the overall package test. Because we are working outside `$GOPATH` and also outside any module, the `go` command knows no import path for the current directory and makes up a fake one based on the directory name: `_/home/gopher/hello`.

最后一行总结了整个软件包的测试情况。因为我们在$GOPATH之外工作，也在任何模块之外，go命令不知道当前目录的导入路径，而是根据目录名编了一个假的路径：_/home/gopher/hello。

Let’s make the current directory the root of a module by using `go mod init` and then try `go test` again:

让我们用go mod init使当前目录成为一个模块的根，然后再试试go test：

```shell linenums="1"
$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello
$ go test
PASS
ok      example.com/hello   0.020s
$
```

Congratulations! You’ve written and tested your first module.

恭喜您！您已经编写并测试了您的第一个模块。您已经编写并测试了您的第一个模块。

The `go mod init` command wrote a `go.mod` file:

go mod init命令写了一个go.mod文件：

```shell linenums="1"
$ cat go.mod
module example.com/hello

go 1.12
$
```

The `go.mod` file only appears in the root of the module. Packages in subdirectories have import paths consisting of the module path plus the path to the subdirectory. For example, if we created a subdirectory `world`, we would not need to (nor want to) run `go mod init` there. The package would automatically be recognized as part of the `example.com/hello` module, with import path `example.com/hello/world`.

go.mod文件只出现在模块的根目录中。子目录中的软件包的导入路径由模块路径加上子目录的路径组成。例如，如果我们创建了一个子目录world，我们就不需要（也不想）在那里运行go mod init。这个包会被自动识别为 example.com/hello 模块的一部分，导入路径为 example.com/hello/world。

## Adding a dependency 添加一个依赖项

The primary motivation for Go modules was to improve the experience of using (that is, adding a dependency on) code written by other developers.

Go模块的主要动机是为了改善使用（也就是添加依赖项）其他开发者编写的代码的体验。

Let’s update our `hello.go` to import `rsc.io/quote` and use it to implement `Hello`:

让我们更新我们的hello.go，导入rsc.io/quote并使用它来实现Hello：

```go linenums="1"
package hello

import "rsc.io/quote"

func Hello() string {
    return quote.Hello()
}
```

Now let’s run the test again:

```shell linenums="1"
$ go test
go: finding rsc.io/quote v1.5.2
go: downloading rsc.io/quote v1.5.2
go: extracting rsc.io/quote v1.5.2
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading rsc.io/sampler v1.3.0
go: extracting rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
PASS
ok      example.com/hello   0.023s
$
```

The `go` command resolves imports by using the specific dependency module versions listed in `go.mod`. When it encounters an `import` of a package not provided by any module in `go.mod`, the `go` command automatically looks up the module containing that package and adds it to `go.mod`, using the latest version. ("Latest" is defined as the latest tagged stable (non-[prerelease](https://semver.org/#spec-item-9)) version, or else the latest tagged prerelease version, or else the latest untagged version.) In our example, `go test` resolved the new import `rsc.io/quote` to the module `rsc.io/quote v1.5.2`. It also downloaded two dependencies used by `rsc.io/quote`, namely `rsc.io/sampler` and `golang.org/x/text`. Only direct dependencies are recorded in the `go.mod` file:

go命令通过使用go.mod中列出的特定依赖模块版本来解决导入问题。当它遇到go.mod中任何模块都没有提供的包的导入时，go命令会自动查找包含该包的模块并将其添加到go.mod中，使用最新的版本。("最新 "被定义为最新的有标签的稳定版（非发布版），或者最新的有标签的预发布版，或者最新的无标签版)。在我们的例子中，go test 将新的 import rsc.io/quote 解析为模块 rsc.io/quote v1.5.2。它还下载了rsc.io/quote使用的两个依赖项，即rsc.io/sampler和golang.org/x/text。在go.mod文件中只记录了直接依赖项：

```shell linenums="1"
$ cat go.mod
module example.com/hello

go 1.12

require rsc.io/quote v1.5.2
$
```

A second `go test` command will not repeat this work, since the `go.mod` is now up-to-date and the downloaded modules are cached locally (in `$GOPATH/pkg/mod`):

第二条go测试命令不会重复这项工作，因为go.mod现在是最新的，而且下载的模块都缓存在本地（在$GOPATH/pkg/mod）：

```shell linenums="1"
$ go test
PASS
ok      example.com/hello   0.020s
$
```

Note that while the `go` command makes adding a new dependency quick and easy, it is not without cost. Your module now literally *depends* on the new dependency in critical areas such as correctness, security, and proper licensing, just to name a few. For more considerations, see Russ Cox’s blog post, "[Our Software Dependency Problem](https://research.swtch.com/deps)."

请注意，虽然go命令使添加新的依赖项变得快速而简单，但它不是没有代价的。您的模块现在在关键领域依赖于新的依赖项，如正确性、安全性和适当的许可，仅举几例。关于更多的考虑，请参阅Russ Cox的博文 "我们的软件依赖问题"。

As we saw above, adding one direct dependency often brings in other indirect dependencies too. The command `go list -m all` lists the current module and all its dependencies:

正如我们在上面看到的那样，增加一个直接的依赖项往往也会带来其他间接的依赖项。命令go list -m all列出了当前模块及其所有依赖项：

```shell linenums="1"
$ go list -m all
example.com/hello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
$
```

In the `go list` output, the current module, also known as the *main module*, is always the first line, followed by dependencies sorted by module path.

在go list的输出中，当前模块，也就是主模块，总是第一行，后面是按模块路径排序的依赖项。

The `golang.org/x/text` version `v0.0.0-20170915032832-14c0d48ead0c` is an example of a [pseudo-version](https://go.dev/cmd/go/#hdr-Pseudo_versions), which is the `go` command’s version syntax for a specific untagged commit.

golang.org/x/text版本v0.0.0-20170915032832-14c0d48ead0c是一个伪版本的例子，它是go命令对特定未标记的提交的版本语法。

In addition to `go.mod`, the `go` command maintains a file named `go.sum` containing the expected [cryptographic hashes](https://go.dev/cmd/go/#hdr-Module_downloading_and_verification) of the content of specific module versions:

除了go.mod之外，go命令还维护一个名为go.sum的文件，其中包含特定模块版本内容的预期加密散列值：

```shell linenums="1"
$ cat go.sum
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZO...
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:Nq...
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3...
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPX...
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/Q...
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9...
$
```

The `go` command uses the `go.sum` file to ensure that future downloads of these modules retrieve the same bits as the first download, to ensure the modules your project depends on do not change unexpectedly, whether for malicious, accidental, or other reasons. Both `go.mod` and `go.sum` should be checked into version control.

go命令使用go.sum文件来确保这些模块的未来下载会检索到与第一次下载相同的比特，以确保您的项目所依赖的模块不会因为恶意的、意外的或其他原因而发生意外变化。go.mod和go.sum都应该被检查进版本控制。

## Upgrading dependencies 升级依赖项

With Go modules, versions are referenced with semantic version tags. A semantic version has three parts: major, minor, and patch. For example, for `v0.1.2`, the major version is 0, the minor version is 1, and the patch version is 2. Let’s walk through a couple minor version upgrades. In the next section, we’ll consider a major version upgrade.

对于Go模块，版本是用语义版本标签来引用的。一个语义版本有三个部分：主版本、次要版本和补丁。例如，对于v0.1.2，主版本是0，次版本是1，而补丁版本是2。在下一节，我们将考虑一个主版本的升级。

From the output of `go list -m all`, we can see we’re using an untagged version of `golang.org/x/text`. Let’s upgrade to the latest tagged version and test that everything still works:

从 go list -m all 的输出中，我们可以看到我们正在使用 golang.org/x/text 的一个未标记的版本。让我们升级到最新的有标签的版本，并测试所有东西是否还能工作：

```shell linenums="1"
$ go get golang.org/x/text
go: finding golang.org/x/text v0.3.0
go: downloading golang.org/x/text v0.3.0
go: extracting golang.org/x/text v0.3.0
$ go test
PASS
ok      example.com/hello   0.013s
$
```

Woohoo! Everything passes. Let’s take another look at `go list -m all` and the `go.mod` file:

呜呼! 一切都通过了。让我们再看一下go list -m all和go.mod文件：

```shell linenums="1"
$ go list -m all
example.com/hello
golang.org/x/text v0.3.0
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
$ cat go.mod
module example.com/hello

go 1.12

require (
    golang.org/x/text v0.3.0 // indirect
    rsc.io/quote v1.5.2
)
$
```

The `golang.org/x/text` package has been upgraded to the latest tagged version (`v0.3.0`). The `go.mod` file has been updated to specify `v0.3.0` too. The `indirect` comment indicates a dependency is not used directly by this module, only indirectly by other module dependencies. See `go help modules` for details.

golang.org/x/text 包已经升级到了最新的标记版本（v0.3.0）。go.mod 文件也已更新为指定的 v0.3.0。间接注释表示一个依赖没有被这个模块直接使用，只是间接地被其他模块依赖。详见go帮助模块。

Now let’s try upgrading the `rsc.io/sampler` minor version. Start the same way, by running `go get` and running tests:

现在让我们试试升级rsc.io/sampler的小版本。以同样的方式开始，运行go get并运行测试：

```shell linenums="1"
$ go get rsc.io/sampler
go: finding rsc.io/sampler v1.99.99
go: downloading rsc.io/sampler v1.99.99
go: extracting rsc.io/sampler v1.99.99
$ go test
--- FAIL: TestHello (0.00s)
    hello_test.go:8: Hello() = "99 bottles of beer on the wall, 99 bottles of beer, ...", want "Hello, world."
FAIL
exit status 1
FAIL    example.com/hello   0.014s
$
```

Uh, oh! The test failure shows that the latest version of `rsc.io/sampler` is incompatible with our usage. Let’s list the available tagged versions of that module:

呃，哦！测试失败表明rsc.io/sampler的最新版本与我们的用法不兼容。让我们列出该模块的可用标记版本：

```shell linenums="1"
$ go list -m -versions rsc.io/sampler
rsc.io/sampler v1.0.0 v1.2.0 v1.2.1 v1.3.0 v1.3.1 v1.99.99
$
```

We had been using v1.3.0; v1.99.99 is clearly no good. Maybe we can try using v1.3.1 instead:

我们一直在使用v1.3.0；v1.99.99显然是不行的。也许我们可以试试用v1.3.1代替：

```shell linenums="1"
$ go get rsc.io/sampler@v1.3.1
go: finding rsc.io/sampler v1.3.1
go: downloading rsc.io/sampler v1.3.1
go: extracting rsc.io/sampler v1.3.1
$ go test
PASS
ok      example.com/hello   0.022s
$
```

Note the explicit `@v1.3.1` in the `go get` argument. In general each argument passed to `go get` can take an explicit version; the default is `@latest`, which resolves to the latest version as defined earlier.

注意go get参数中明确的@v1.3.1。一般来说，传递给go get的每个参数都可以有一个明确的版本；默认的是@latest，它解析到前面定义的最新版本。

## Adding a dependency on a new major version 添加对一个新的主版本的依赖项

Let’s add a new function to our package: `func Proverb` returns a Go concurrency proverb, by calling `quote.Concurrency`, which is provided by the module `rsc.io/quote/v3`. First we update `hello.go` to add the new function:

让我们为我们的包添加一个新的函数：func Proverb通过调用rsc.io/quote/v3模块提供的quote.Concurrency，返回一个Go并发谚语。首先我们更新hello.go，添加新的函数：

```go linenums="1"
package hello

import (
    "rsc.io/quote"
    quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
    return quote.Hello()
}

func Proverb() string {
    return quoteV3.Concurrency()
}
```

Then we add a test to `hello_test.go`:

然后我们在hello_test.go中添加一个测试：

```go linenums="1"
func TestProverb(t *testing.T) {
    want := "Concurrency is not parallelism."
    if got := Proverb(); got != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}
```

Then we can test our code:

然后我们可以测试我们的代码：

```shell linenums="1"
$ go test
go: finding rsc.io/quote/v3 v3.1.0
go: downloading rsc.io/quote/v3 v3.1.0
go: extracting rsc.io/quote/v3 v3.1.0
PASS
ok      example.com/hello   0.024s
$
```

Note that our module now depends on both `rsc.io/quote` and `rsc.io/quote/v3`:

注意，我们的模块现在同时依赖于rsc.io/quote和rsc.io/quote/v3：

```shell linenums="1"
$ go list -m rsc.io/q...
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0
$
```

Each different major version (`v1`, `v2`, and so on) of a Go module uses a different module path: starting at `v2`, the path must end in the major version. In the example, `v3` of `rsc.io/quote` is no longer `rsc.io/quote`: instead, it is identified by the module path `rsc.io/quote/v3`. This convention is called [semantic import versioning](https://research.swtch.com/vgo-import), and it gives incompatible packages (those with different major versions) different names. In contrast, `v1.6.0` of `rsc.io/quote` should be backwards-compatible with `v1.5.2`, so it reuses the name `rsc.io/quote`. (In the previous section, `rsc.io/sampler` `v1.99.99` *should* have been backwards-compatible with `rsc.io/sampler` `v1.3.0`, but bugs or incorrect client assumptions about module behavior can both happen.)

Go模块的每个不同的主版本（v1、v2，以此类推）都使用不同的模块路径：从v2开始，路径必须以主版本结束。在这个例子中，rsc.io/quote的v3不再是rsc.io/quote：而是用模块路径rsc.io/quote/v3标识。这种惯例被称为语义导入版本划分，它给不兼容的包（那些主版本不同的包）起了不同的名字。相比之下，rsc.io/quote的v1.6.0版本应该是向后兼容v1.5.2版本的，所以它重用了rsc.io/quote这个名字。(在上一节中，rsc.io/sampler v1.99.99应该与rsc.io/sampler v1.3.0向后兼容，但错误或客户端对模块行为的不正确假设都可能发生)。

The `go` command allows a build to include at most one version of any particular module path, meaning at most one of each major version: one `rsc.io/quote`, one `rsc.io/quote/v2`, one `rsc.io/quote/v3`, and so on. This gives module authors a clear rule about possible duplication of a single module path: it is impossible for a program to build with both `rsc.io/quote v1.5.2` and `rsc.io/quote v1.6.0`. At the same time, allowing different major versions of a module (because they have different paths) gives module consumers the ability to upgrade to a new major version incrementally. In this example, we wanted to use `quote.Concurrency` from `rsc/quote/v3 v3.1.0` but are not yet ready to migrate our uses of `rsc.io/quote v1.5.2`. The ability to migrate incrementally is especially important in a large program or codebase.

go命令允许构建时最多包括任何特定模块路径的一个版本，意味着每个主版本最多一个：一个rsc.io/quote，一个rsc.io/quote/v2，一个rsc.io/quote/v3，等等。这为模块作者提供了一个关于单个模块路径可能重复的明确规则：一个程序不可能同时用rsc.io/quote v1.5.2和rsc.io/quote v1.6.0构建。同时，允许一个模块的不同主版本（因为它们有不同的路径）给模块消费者提供了逐步升级到新主版本的能力。在这个例子中，我们想使用来自rsc/quote/v3 v3.1.0的quote.Concurrency，但还没有准备好迁移我们对rsc.io/quote v1.5.2的使用。渐进式迁移的能力对于一个大型程序或代码库来说尤其重要。

## Upgrading a dependency to a new major version 将一个依赖项升级到一个新的主版本

Let’s complete our conversion from using `rsc.io/quote` to using only `rsc.io/quote/v3`. Because of the major version change, we should expect that some APIs may have been removed, renamed, or otherwise changed in incompatible ways. Reading the docs, we can see that `Hello` has become `HelloV3`:

让我们完成从使用rsc.io/quote到只使用rsc.io/quote/v3的转换。由于主版本的改变，我们应该预料到一些API可能已经被删除、重命名或以其他不兼容的方式改变。阅读文档，我们可以看到Hello已经变成了HelloV3：

```sehll linenums="1"
$ go doc rsc.io/quote/v3
package quote // import "rsc.io/quote/v3"

Package quote collects pithy sayings.

func Concurrency() string
func GlassV3() string
func GoV3() string
func HelloV3() string
func OptV3() string
$
```

We can update our use of `quote.Hello()` in `hello.go` to use `quoteV3.HelloV3()`:

我们可以在hello.go中更新对quote.Hello()的使用，使用quoteV3.HelloV3()：

```go linenums="1"
package hello

import quoteV3 "rsc.io/quote/v3"

func Hello() string {
    return quoteV3.HelloV3()
}

func Proverb() string {
    return quoteV3.Concurrency()
}
```

And then at this point, there’s no need for the renamed import anymore, so we can undo that:

然后在这一点上，已经不需要重命名的导入了，所以我们可以撤销这个：

```go linenums="1"
package hello

import "rsc.io/quote/v3"

func Hello() string {
    return quote.HelloV3()
}

func Proverb() string {
    return quote.Concurrency()
}
```

Let’s re-run the tests to make sure everything is working:

让我们重新运行测试以确保一切正常：

```shell linenums="1"
$ go test
PASS
ok      example.com/hello       0.014s
```

## Removing unused dependencies 移除未使用的依赖项

We’ve removed all our uses of `rsc.io/quote`, but it still shows up in `go list -m all` and in our `go.mod` file:

我们已经删除了对rsc.io/quote的所有使用，但它仍然显示在go list -m all和go.mod文件中：

```shell linenums="1"
$ go list -m all
example.com/hello
golang.org/x/text v0.3.0
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.1
$ cat go.mod
module example.com/hello

go 1.12

require (
    golang.org/x/text v0.3.0 // indirect
    rsc.io/quote v1.5.2
    rsc.io/quote/v3 v3.0.0
    rsc.io/sampler v1.3.1 // indirect
)
$
```

Why? Because building a single package, like with `go build` or `go test`, can easily tell when something is missing and needs to be added, but not when something can safely be removed. Removing a dependency can only be done after checking all packages in a module, and all possible build tag combinations for those packages. An ordinary build command does not load this information, and so it cannot safely remove dependencies.

为什么？因为构建单个软件包，比如用go build或go test，可以很容易地知道什么时候缺少什么东西需要添加，但不能知道什么时候可以安全地删除。只有在检查了一个模块中的所有包，以及这些包的所有可能的构建标签组合之后，才能删除一个依赖项。普通的构建命令不会加载这些信息，所以它不能安全地移除依赖项。

The `go mod tidy` command cleans up these unused dependencies:

go mod tidy 命令清理了这些未使用的依赖项：

```shell linenums="1"
$ go mod tidy
$ go list -m all
example.com/hello
golang.org/x/text v0.3.0
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.1
$ cat go.mod
module example.com/hello

go 1.12

require (
    golang.org/x/text v0.3.0 // indirect
    rsc.io/quote/v3 v3.1.0
    rsc.io/sampler v1.3.1 // indirect
)

$ go test
PASS
ok      example.com/hello   0.020s
$
```

## Conclusion 结论

Go modules are the future of dependency management in Go. Module functionality is now available in all supported Go versions (that is, in Go 1.11 and Go 1.12).

Go模块是Go中依赖项管理的未来。模块功能现在可以在所有支持的 Go 版本中使用（也就是在 Go 1.11 和 Go 1.12 中）。

This post introduced these workflows using Go modules:

本帖介绍了这些使用Go模块的工作流程：

- `go mod init` creates a new module, initializing the `go.mod` file that describes it. 创建一个新模块，初始化描述该模块的 go.mod 文件。
- `go build`, `go test`, and other package-building commands add new dependencies to `go.mod` as needed. go build、go test和其他建包命令根据需要向go.mod添加新的依赖项。
- `go list -m all` prints the current module’s dependencies.  go list -m all 打印出当前模块的依赖项。
- `go get` changes the required version of a dependency (or adds a new dependency). go get 更改某个依赖的必要版本（或添加一个新的依赖）。
- `go mod tidy` removes unused dependencies. go mod tidy 删除未使用的依赖项。

We encourage you to start using modules in your local development and to add `go.mod` and `go.sum` files to your projects. To provide feedback and help shape the future of dependency management in Go, please send us [bug reports](https://go.dev/issue/new) or [experience reports](https://go.dev/wiki/ExperienceReports).

我们鼓励您在本地开发中开始使用模块，并在您的项目中添加go.mod和go.sum文件。为了提供反馈并帮助塑造Go中依赖项管理的未来，请向我们发送错误报告或经验报告。

Thanks for all your feedback and help improving modules.

谢谢您的反馈和对改进模块的帮助。
