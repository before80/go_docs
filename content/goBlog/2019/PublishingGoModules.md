+++
title = "发布 go 模块"
weight = 6
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Publishing Go Modules - 发布 go 模块

https://go.dev/blog/publishing-go-modules

Tyler Bui-Palsulich
26 September 2019

## Introduction 简介

This post is part 3 in a series.

这篇文章是系列文章的第三部分。

- Part 1 — [Using Go Modules 使用Go模块](https://go.dev/blog/using-go-modules)
- Part 2 — [Migrating To Go Modules 迁移到Go模块](https://go.dev/blog/migrating-to-go-modules)
- **Part 3 — Publishing Go Modules** (this post) 发布Go模块（本帖）
- Part 4 — [Go Modules: v2 and Beyond Go模块：V2版及以后](https://go.dev/blog/v2-go-modules)
- Part 5 — [Keeping Your Modules Compatible 保持模块的兼容性](https://go.dev/blog/module-compatibility)

**Note:** For documentation on developing modules, see [Developing and publishing modules](https://go.dev/doc/modules/developing).

注意：关于开发模块的文档，请参见开发和发布模块。

This post discusses how to write and publish modules so other modules can depend on them.

这篇文章讨论了如何编写和发布模块，以便其他模块可以依赖它们。

Please note: this post covers development up to and including `v1`. If you are interested in `v2`, please see [Go Modules: v2 and Beyond](https://go.dev/blog/v2-go-modules).

如果您对v2版感兴趣，请看Go Modules: v2 and Beyond。

This post uses [Git](https://git-scm.com/) in examples. [Mercurial](https://www.mercurial-scm.org/), [Bazaar](http://wiki.bazaar.canonical.com/), and others are supported as well.

本帖在例子中使用了Git。也支持Mercurial、Bazaar和其他。

## Project setup 项目设置

For this post, you’ll need an existing project to use as an example. So, start with the files from the end of the [Using Go Modules](https://blog.golang.org/using-go-modules) article:

在这篇文章中，您需要一个现有的项目来作为例子。因此，从《使用Go模块》一文末尾的文件开始：

```shell linenums="1"
$ cat go.mod
module example.com/hello

go 1.12

require rsc.io/quote/v3 v3.1.0

$ cat go.sum
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote/v3 v3.1.0 h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=
rsc.io/quote/v3 v3.1.0/go.mod h1:yEA65RcK8LyAZtP9Kv3t0HmxON59tX3rD+tICJqUlj0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=

$ cat hello.go
package hello

import "rsc.io/quote/v3"

func Hello() string {
    return quote.HelloV3()
}

func Proverb() string {
    return quote.Concurrency()
}

$ cat hello_test.go
package hello

import (
    "testing"
)

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}

func TestProverb(t *testing.T) {
    want := "Concurrency is not parallelism."
    if got := Proverb(); got != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}

$
```

Next, create a new `git` repository and add an initial commit. If you’re publishing your own project, be sure to include a `LICENSE` file. Change to the directory containing the `go.mod` then create the repo:

接下来，创建一个新的git仓库并添加一个初始提交。如果您要发布自己的项目，一定要包括一个LICENSE文件。改变到包含go.mod的目录，然后创建版本库：

```shell linenums="1"
$ git init
$ git add LICENSE go.mod go.sum hello.go hello_test.go
$ git commit -m "hello: initial commit"
$
```

## Semantic versions and modules 语义版本和模块

Every required module in a `go.mod` has a [semantic version](https://semver.org/), the minimum version of that dependency to use to build the module.

go.mod中的每个所需模块都有一个语义版本，即构建模块时使用的该依赖关系的最小版本。

A semantic version has the form `vMAJOR.MINOR.PATCH`.

语义版本的形式为vMAJOR.MINOR.PATCH。

- Increment the `MAJOR` version when you make a [backwards incompatible](https://go.dev/doc/go1compat) change to the public API of your module. This should only be done when absolutely necessary.当您对您的模块的公共API进行向后不兼容的修改时，要增加MAJOR版本。只有在绝对必要的时候才应该这样做。
- Increment the `MINOR` version when you make a backwards compatible change to the API, like changing dependencies or adding a new function, method, struct field, or type.当您对API进行向后兼容的改变时，增加MINOR版本，比如改变依赖关系或添加新的函数、方法、结构域或类型。
- Increment the `PATCH` version after making minor changes that don’t affect your module’s public API or dependencies, like fixing a bug.在做了不影响您的模块的公共API或依赖关系的小改动后，增加PATCH版本，如修复一个错误。

You can specify pre-release versions by appending a hyphen and dot separated identifiers (for example, `v1.0.1-alpha` or `v2.2.2-beta.2`). Normal releases are preferred by the `go` command over pre-release versions, so users must ask for pre-release versions explicitly (for example, `go get example.com/hello@v1.0.1-alpha`) if your module has any normal releases.

您可以通过附加一个连字符和点分隔的标识符来指定预发布版本（例如，v1.0.1-alpha 或 v2.2.2-beta.2）。普通版本是go命令的首选，而不是预发布版本，所以如果您的模块有任何普通版本，用户必须明确要求预发布版本（例如，go get example.com/hello@v1.0.1-alpha）。

`v0` major versions and pre-release versions do not guarantee backwards compatibility. They let you refine your API before making stability commitments to your users. However, `v1` major versions and beyond require backwards compatibility within that major version.

v0主要版本和预发布版本并不保证向后兼容。它们让您在向用户做出稳定性承诺之前完善您的API。然而，v1主要版本及以后的版本需要在该主要版本内向后兼容。

The version referenced in a `go.mod` may be an explicit release tagged in the repository (for example, `v1.5.2`), or it may be a [pseudo-version](https://go.dev/ref/mod#pseudo-versions) based on a specific commit (for example, `v0.0.0-20170915032832-14c0d48ead0c`). Pseudo-versions are a special type of pre-release version. Pseudo-versions are useful when a user needs to depend on a project that has not published any semantic version tags, or develop against a commit that hasn’t been tagged yet, but users should not assume that pseudo-versions provide a stable or well-tested API. Tagging your modules with explicit versions signals to your users that specific versions are fully tested and ready to use.

go.mod中引用的版本可能是版本库中明确标记的版本（例如，v1.5.2），也可能是基于特定提交的伪版本（例如，v0.0.0-20170915032832-14c0d48ead0c）。伪版本是一种特殊类型的预发布版本。当用户需要依赖一个尚未发布任何语义版本标签的项目，或者需要针对一个尚未被标记的提交进行开发时，伪版本是非常有用的，但是用户不应该认为伪版本提供了一个稳定的或者经过良好测试的API。用明确的版本来标记您的模块，向您的用户发出信号，表明特定的版本已经过充分的测试，可以使用。

Once you start tagging your repo with versions, it’s important to keep tagging new releases as you develop your module. When users request a new version of your module (with `go get -u` or `go get example.com/hello`), the `go` command will choose the greatest semantic release version available, even if that version is several years old and many changes behind the primary branch. Continuing to tag new releases will make your ongoing improvements available to your users.

一旦您开始用版本标记您的仓库，重要的是在您开发模块时不断标记新版本。当用户请求您的模块的新版本时（用go get -u或go get example.com/hello），go命令会选择可用的最大语义发布版本，即使该版本是几年前的，而且在主分支后面有许多变化。继续对新的版本进行标记，将使您正在进行的改进能够为您的用户所用。

Do not delete version tags from your repo. If you find a bug or a security issue with a version, release a new version. If people depend on a version that you have deleted, their builds may fail. Similarly, once you release a version, do not change or overwrite it. The [module mirror and checksum database](https://blog.golang.org/module-mirror-launch) store modules, their versions, and signed cryptographic hashes to ensure that the build of a given version remains reproducible over time.

请不要从您的版本库中删除版本标签。如果您发现某个版本有错误或安全问题，就发布一个新的版本。如果人们依赖一个被您删除的版本，他们的构建可能会失败。同样，一旦您发布了一个版本，就不要改变或覆盖它。模块镜像和校验数据库存储模块、它们的版本和签名的加密哈希值，以确保特定版本的构建在一段时间内保持可重复性。

## v0: the initial, unstable version v0：最初的、不稳定的版本

Let’s tag the module with a `v0` semantic version. A `v0` version does not make any stability guarantees, so nearly all projects should start with `v0` as they refine their public API.

让我们给模块贴上v0语义版本的标签。v0版本不做任何稳定性保证，因此几乎所有的项目在完善其公共API时都应该从v0开始。

Tagging a new version has a few steps:

给新版本打上标签有几个步骤：

1. Run `go mod tidy`, which removes any dependencies the module might have accumulated that are no longer necessary.运行go mod tidy，删除该模块可能积累的、不再需要的任何依赖关系。
2. Run `go test ./...` a final time to make sure everything is working. 最后一次运行go test ./...以确保一切正常。
3. Tag the project with a new version using [`git tag`](https://git-scm.com/docs/git-tag). 用git tag给项目打上新版本的标签。
4. Push the new tag to the origin repository. 将新的标签推送到源码库中。

```shell linenums="1"
$ go mod tidy
$ go test ./...
ok      example.com/hello       0.015s
$ git add go.mod go.sum hello.go hello_test.go
$ git commit -m "hello: changes for v0.1.0"
$ git tag v0.1.0
$ git push origin v0.1.0
$
```

Now other projects can depend on `v0.1.0` of `example.com/hello`. For your own module, you can run `go list -m example.com/hello@v0.1.0` to confirm the latest version is available (this example module does not exist, so no versions are available). If you don’t see the latest version immediately and you’re using the Go module proxy (the default since Go 1.13), try again in a few minutes to give the proxy time to load the new version.

现在其他项目可以依赖 example.com/hello 的 v0.1.0。对于您自己的模块，您可以运行 go list -m example.com/hello@v0.1.0 来确认最新的版本是否可用（这个例子的模块不存在，所以没有版本可用）。如果您没有立即看到最新的版本，而且您使用的是Go模块代理（从Go 1.13开始默认的），过几分钟再试试，让代理有时间加载新版本。

If you add to the public API, make a breaking change to a `v0` module, or upgrade the minor or version of one of your dependencies, increment the `MINOR` version for your next release. For example, the next release after `v0.1.0` would be `v0.2.0`.

如果您增加了公共 API，对 v0 模块做了突破性的改变，或者升级了您的一个依赖关系的次要版本，为您的下一个版本增加 MINOR 版本。例如，v0.1.0之后的下一个版本将是v0.2.0。

If you fix a bug in an existing version, increment the `PATCH` version. For example, the next release after `v0.1.0` would be `v0.1.1`.

如果您修复了一个现有版本中的错误，请增加 PATCH 版本。例如，v0.1.0之后的下一个版本将是v0.1.1。

## v1: the first stable version - v1：第一个稳定版本

Once you are absolutely sure your module’s API is stable, you can release `v1.0.0`. A `v1` major version communicates to users that no incompatible changes will be made to the module’s API. They can upgrade to new `v1` minor and patch releases, and their code should not break. Function and method signatures will not change, exported types will not be removed, and so on. If there are changes to the API, they will be backwards compatible (for example, adding a new field to a struct) and will be included in a new minor release. If there are bug fixes (for example, a security fix), they will be included in a patch release (or as part of a minor release).

一旦您完全确定您的模块的API是稳定的，您就可以发布v1.0.0。一个v1大版本向用户传达了将不会对模块的API进行不兼容的修改。他们可以升级到新的v1小版本和补丁版本，而且他们的代码应该不会被破坏。函数和方法的签名不会改变，导出的类型不会被删除，等等。如果API有变化，它们将是向后兼容的（例如，为结构体添加新的字段），并将包含在新的次要版本中。如果有错误修复（例如，安全修复），它们将被包含在一个补丁版本中（或作为一个次要版本的一部分）。

Sometimes, maintaining backwards compatibility can lead to awkward APIs. That’s OK. An imperfect API is better than breaking users' existing code.

有时，保持向后的兼容性会导致尴尬的API。这没关系。一个不完美的API总比破坏用户的现有代码要好。

The standard library’s `strings` package is a prime example of maintaining backwards compatibility at the cost of API consistency.

标准库的字符串包是以API的一致性为代价来维持向后兼容性的一个典型例子。

- [`Split`](https://godoc.org/strings#Split) slices a string into all substrings separated by a separator and returns a slice of the substrings between those separators. Split将一个字符串切成由分隔符分隔的所有子字符串，并返回这些分隔符之间的子字符串的一个片断。
- [`SplitN`](https://godoc.org/strings#SplitN) can be used to control the number of substrings to return. SplitN可以用来控制要返回的子串的数量。

However, [`Replace`](https://godoc.org/strings#Replace) took a count of how many instances of the string to replace from the beginning (unlike `Split`).

然而，Replace从头开始计算要替换多少个字符串实例（与Split不同）。

Given `Split` and `SplitN`, you would expect functions like `Replace` and `ReplaceN`. But, we couldn’t change the existing `Replace` without breaking callers, which we promised not to do. So, in Go 1.12, we added a new function, [`ReplaceAll`](https://godoc.org/strings#ReplaceAll). The resulting API is a little odd, since `Split` and `Replace` behave differently, but that inconsistency is better than a breaking change.

鉴于Split和SplitN，您会想到Replace和ReplaceN这样的函数。但是，我们不能在不破坏调用者的情况下改变现有的Replace，我们承诺不会这样做。因此，在Go 1.12中，我们添加了一个新的函数，ReplaceAll。由此产生的API有点奇怪，因为Split和Replace的行为是不同的，但这种不一致总比破坏性的改变好。

Let’s say you’re happy with the API of `example.com/hello` and you want to release `v1` as the first stable version.

假设您对example.com/hello的API很满意，您想发布v1作为第一个稳定版本。

Tagging `v1` uses the same process as tagging a `v0` version: run `go mod tidy` and `go test ./...`, tag the version, and push the tag to the origin repository:

标记v1版本的过程与标记v0版本的过程相同：运行go mod tidy和go test ./...，标记版本，并将标记推送到源码库：

```shell linenums="1"
$ go mod tidy
$ go test ./...
ok      example.com/hello       0.015s
$ git add go.mod go.sum hello.go hello_test.go
$ git commit -m "hello: changes for v1.0.0"
$ git tag v1.0.0
$ git push origin v1.0.0
$
```

At this point, the `v1` API of `example.com/hello` is solidified. This communicates to everyone that our API is stable and they should feel comfortable using it.

至此，example.com/hello的v1版API就稳固了。这向大家传达了我们的API是稳定的，他们应该放心地使用它。

## Conclusion 总结

This post walked through the process of tagging a module with semantic versions and when to release `v1`. A future post will cover how to maintain and publish modules at `v2` and beyond.

本篇文章介绍了用语义版本标记模块的过程，以及何时发布V1版。

To provide feedback and help shape the future of dependency management in Go, please send us [bug reports](https://go.dev/issue/new) or [experience reports](https://go.dev/wiki/ExperienceReports).

为了提供反馈并帮助塑造Go中依赖项管理的未来，请向我们发送错误报告或经验报告。

Thanks for all your feedback and help improving Go modules.

谢谢您的反馈和对改进Go模块的帮助。
