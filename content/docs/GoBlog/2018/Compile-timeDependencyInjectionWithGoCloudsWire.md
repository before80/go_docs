+++
title = "使用 go Cloud的Wire进行编译时依赖注入"
weight = 6
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Compile-time Dependency Injection With Go Cloud's Wire - 使用 go Cloud的Wire进行编译时依赖注入

https://go.dev/blog/wire

Robert van Gent
9 October 2018

## Overview 概述

The Go team recently [announced](https://blog.golang.org/go-cloud) the open source project [Go Cloud](https://github.com/google/go-cloud), with portable Cloud APIs and tools for [open cloud](https://cloud.google.com/open-cloud/) development. This post goes into more detail about Wire, a dependency injection tool used in Go Cloud.

Go团队最近宣布了开源项目Go Cloud，拥有可移植的云计算API和用于开放云开发的工具。这篇文章更详细地介绍了Wire，一个在Go Cloud中使用的依赖注入工具。

## What problem does Wire solve? - Wire能解决什么问题？

[Dependency injection](https://en.wikipedia.org/wiki/Dependency_injection) is a standard technique for producing flexible and loosely coupled code, by explicitly providing components with all of the dependencies they need to work. In Go, this often takes the form of passing dependencies to constructors:

依赖注入是一种标准的技术，通过明确提供组件工作所需的所有依赖，来产生灵活和松散耦合的代码。在Go中，这通常采取的形式是将依赖关系传递给构造函数。

```go linenums="1"
// NewUserStore returns a UserStore that uses cfg and db as dependencies.
func NewUserStore(cfg *Config, db *mysql.DB) (*UserStore, error) {...}
```

This technique works great at small scale, but larger applications can have a complex graph of dependencies, resulting in a big block of initialization code that’s order-dependent but otherwise not very interesting. It’s often hard to break up this code cleanly, especially because some dependencies are used multiple times. Replacing one implementation of a service with another can be painful because it involves modifying the dependency graph by adding a whole new set of dependencies (and their dependencies…), and removing unused old ones. In practice, making changes to initialization code in applications with large dependency graphs is tedious and slow.

这种技术在小范围内非常有效，但更大的应用程序可能会有一个复杂的依赖关系图，导致一大块初始化代码，这些代码的顺序是依赖性的，但其他方面并不有趣。通常很难将这些代码分解得很干净，特别是因为有些依赖关系被多次使用。用一个服务的一个实现替换另一个实现是很痛苦的，因为这涉及到修改依赖关系图，增加一组新的依赖关系（以及它们的依赖关系......），并删除未使用的旧依赖关系。在实践中，在具有大型依赖图的应用程序中对初始化代码进行修改是很乏味和缓慢的。

Dependency injection tools like Wire aim to simplify the management of initialization code. You describe your services and their dependencies, either as code or as configuration, then Wire processes the resulting graph to figure out ordering and how to pass each service what it needs. Make changes to an application’s dependencies by changing a function signature or adding or removing an initializer, and then let Wire do the tedious work of generating initialization code for the entire dependency graph.

像Wire这样的依赖注入工具，旨在简化初始化代码的管理。您以代码或配置的形式描述您的服务和它们的依赖关系，然后Wire处理产生的图，以找出排序和如何传递给每个服务它所需要的东西。通过改变函数签名或添加或删除初始化器来改变应用程序的依赖关系，然后让Wire完成为整个依赖关系图生成初始化代码的繁琐工作。

## Why is this part of Go Cloud? 为什么这是Go Cloud的一部分？

Go Cloud’s goal is to make it easier to write portable Cloud applications by providing idiomatic Go APIs for useful Cloud services. For example, [blob.Bucket](https://godoc.org/github.com/google/go-cloud/blob) provides a storage API with implementations for Amazon’s S3 and Google Cloud Storage (GCS); applications written using `blob.Bucket` can swap implementations without changing their application logic. However, the initialization code is inherently provider-specific, and each provider has a different set of dependencies.

Go Cloud的目标是通过为有用的云服务提供习惯性的Go API，使编写可移植的云应用程序更加容易。例如，blob.Bucket提供了一个存储API，为亚马逊的S3和谷歌云存储（GCS）提供了实现；使用blob.Bucket编写的应用程序可以在不改变其应用逻辑的情况下交换实现方式。然而，初始化代码在本质上是针对提供商的，每个提供商都有一套不同的依赖关系。

For example, [constructing a GCS `blob.Bucket`](https://godoc.org/github.com/google/go-cloud/blob/gcsblob#OpenBucket) requires a `gcp.HTTPClient`, which eventually requires `google.Credentials`, while [constructing one for S3](https://godoc.org/github.com/google/go-cloud/blob/s3blob) requires an `aws.Config`, which eventually requires AWS credentials. Thus, updating an application to use a different `blob.Bucket` implementation involves exactly the kind of tedious update to the dependency graph that we described above. The driving use case for Wire is to make it easy to swap implementations of Go Cloud portable APIs, but it’s also a general-purpose tool for dependency injection.

例如，构建GCS blob.Bucket需要gcp.HTTPClient，最终需要google.Credentials，而构建S3需要aws.Config，最终需要AWS credentials。因此，更新应用程序以使用不同的blob.Bucket实现，就需要对我们上面描述的依赖关系图进行繁琐的更新。Wire的主要用途是使Go Cloud可移植API的实现易于交换，但它也是一个通用的依赖注入工具。

## Hasn’t this been done already? 这不是已经完成了吗？

There are a number of dependency injection frameworks out there. For Go, [Uber’s dig](https://github.com/uber-go/dig) and [Facebook’s inject](https://github.com/facebookgo/inject) both use reflection to do runtime dependency injection. Wire was primarily inspired by Java’s [Dagger 2](https://google.github.io/dagger/), and uses code generation rather than reflection or [service locators](https://en.wikipedia.org/wiki/Service_locator_pattern).

现在有很多的依赖注入框架。对于Go来说，Uber的dig和Facebook的inject都使用反射来进行运行时依赖注入。Wire主要受Java的Dagger 2启发，使用代码生成而不是反射或服务定位器。

We think this approach has several advantages:

我们认为这种方法有几个优点：

- Runtime dependency injection can be hard to follow and debug when the dependency graph gets complex. Using code generation means that the initialization code that’s executed at runtime is regular, idiomatic Go code that’s easy to understand and debug. Nothing is obfuscated by an intervening framework doing "magic". In particular, problems like forgetting a dependency become compile-time errors, not run-time errors.当依赖关系图变得复杂时，运行时依赖关系注入会很难跟踪和调试。使用代码生成意味着在运行时执行的初始化代码是常规的、习惯性的Go代码，易于理解和调试。没有任何东西会被干预框架的 "魔法 "所迷惑。特别是，像忘记依赖关系这样的问题会成为编译时错误，而不是运行时错误。
- Unlike [service locators](https://en.wikipedia.org/wiki/Service_locator_pattern), there’s no need to make up arbitrary names or keys to register services. Wire uses Go types to connect components with their dependencies.与服务定位器不同，不需要编造任意的名字或键来注册服务。Wire使用Go类型来连接组件和它们的依赖关系。
- It’s easier to avoid dependency bloat. Wire’s generated code will only import the dependencies you need, so your binary won’t have unused imports. Runtime dependency injectors can’t identify unused dependencies until runtime.这更容易避免依赖性的膨胀。Wire生成的代码只导入您需要的依赖，所以您的二进制文件不会有未使用的导入。运行时的依赖性注入器在运行时才能识别未使用的依赖性。
- Wire’s dependency graph is knowable statically, which provides opportunities for tooling and visualization.Wire的依赖图是静态可知的，这为工具化和可视化提供了机会。

## How does it work? 它是如何工作的？

Wire has two basic concepts: providers and injectors.

Wire有两个基本概念：提供者和注入者。

*Providers* are ordinary Go functions that "provide" values given their dependencies, which are described simply as parameters to the function. Here’s some sample code that defines three providers:

提供者是普通的Go函数，它 "提供 "给它们的依赖值，这些依赖值被简单描述为函数的参数。下面是一些定义了三个提供者的示例代码：

```go linenums="1"
// NewUserStore is the same function we saw above; it is a provider for UserStore,
// with dependencies on *Config and *mysql.DB.
func NewUserStore(cfg *Config, db *mysql.DB) (*UserStore, error) {...}

// NewDefaultConfig is a provider for *Config, with no dependencies.
func NewDefaultConfig() *Config {...}

// NewDB is a provider for *mysql.DB based on some connection info.
func NewDB(info *ConnectionInfo) (*mysql.DB, error) {...}
```

Providers that are commonly used together can be grouped into `ProviderSets`. For example, it’s common to use a default `*Config` when creating a `*UserStore`, so we can group `NewUserStore` and `NewDefaultConfig` in a `ProviderSet`:

通常一起使用的提供者可以被分组为ProviderSets。例如，在创建*UserStore时，使用默认的*Config是很常见的，所以我们可以把NewUserStore和NewDefaultConfig归入一个ProviderSet：

```go linenums="1"
var UserStoreSet = wire.ProviderSet(NewUserStore, NewDefaultConfig)
```

*Injectors* are generated functions that call providers in dependency order. You write the injector’s signature, including any needed inputs as arguments, and insert a call to `wire.Build` with the list of providers or provider sets that are needed to construct the end result:

注入器是生成的函数，它按依赖关系的顺序调用提供者。您编写注入器的签名，包括任何需要的输入作为参数，并插入对 wire.Build 的调用，其中包括构建最终结果所需的提供者或提供者集的列表：

```go linenums="1"
func initUserStore() (*UserStore, error) {
    // We're going to get an error, because NewDB requires a *ConnectionInfo
    // and we didn't provide one.
    wire.Build(UserStoreSet, NewDB)
    return nil, nil  // These return values are ignored.
}
```

Now we run go generate to execute wire:

现在我们运行go generate来执行wire：

```shell linenums="1"
$ go generate
wire.go:2:10: inject initUserStore: no provider found for ConnectionInfo (required by provider of *mysql.DB)
wire: generate failed
```

Oops! We didn’t include a `ConnectionInfo` or tell Wire how to build one. Wire helpfully tells us the line number and types involved. We can either add a provider for it to `wire.Build`, or add it as an argument:

哎呀! 我们没有包括ConnectionInfo，也没有告诉Wire如何建立一个。Wire 很有帮助地告诉我们所涉及的行号和类型。我们可以在 wire.Build 中为它添加一个提供者，或者将它作为一个参数添加：

```go linenums="1"
func initUserStore(info ConnectionInfo) (*UserStore, error) {
    wire.Build(UserStoreSet, NewDB)
    return nil, nil  // These return values are ignored.
}
```

Now `go generate` will create a new file with the generated code:

现在go generate将用生成的代码创建一个新文件：

```go linenums="1"
// File: wire_gen.go
// Code generated by Wire. DO NOT EDIT.
//go:generate wire
//+build !wireinject

func initUserStore(info ConnectionInfo) (*UserStore, error) {
    defaultConfig := NewDefaultConfig()
    db, err := NewDB(info)
    if err != nil {
        return nil, err
    }
    userStore, err := NewUserStore(defaultConfig, db)
    if err != nil {
        return nil, err
    }
    return userStore, nil
}
```

Any non-injector declarations are copied into the generated file. There is no dependency on Wire at runtime: all of the written code is just normal Go code.

任何非注入器的声明都被复制到生成的文件中。在运行时没有对Wire的依赖性：所有编写的代码都是正常的Go代码。

As you can see, the output is very close to what a developer would write themselves. This was a trivial example with just three components, so writing the initializer by hand wouldn’t be too painful, but Wire saves a lot of manual toil for components and applications with more complex dependency graphs.

正如您所看到的，输出结果非常接近于开发者自己写的东西。这是一个只有三个组件的微不足道的例子，所以用手写初始化器不会太痛苦，但对于具有更复杂依赖关系图的组件和应用程序来说，Wire可以节省大量的手工劳作。

## How can I get involved and learn more? 我怎样才能参与并了解更多？

The [Wire README](https://github.com/google/wire/blob/master/README.md) goes into more detail about how to use Wire and its more advanced features. There’s also a [tutorial](https://github.com/google/wire/tree/master/_tutorial) that walks through using Wire in a simple application.

Wire的README更详细地介绍了如何使用Wire及其更高级的功能。还有一个教程，介绍了在一个简单的应用程序中使用Wire。

We appreciate any input you have about your experience with Wire! [Wire’s](https://github.com/google/wire) development is conducted on GitHub, so you can [file an issue](https://github.com/google/wire/issues/new/choose) to tell us what could be better. For updates and discussion about the project, join [the Go Cloud mailing list](https://groups.google.com/forum/#!forum/go-cloud).

我们感谢您对使用Wire的经验提出的任何意见。Wire的开发是在GitHub上进行的，所以您可以提交一个问题来告诉我们什么地方可以做得更好。关于项目的更新和讨论，请加入Go Cloud邮件列表。

Thank you for taking the time to learn about Go Cloud’s Wire. We’re excited to work with you to make Go the language of choice for developers building portable cloud applications.

感谢您花时间了解Go Cloud的Wire。我们很高兴能与您合作，使 Go 成为开发人员构建可移植云应用程序的首选语言。
