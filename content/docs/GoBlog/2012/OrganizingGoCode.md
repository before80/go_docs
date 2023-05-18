+++
title = "组织 go 代码"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Organizing Go code - 组织 go 代码

https://go.dev/blog/organizing-go-code

Andrew Gerrand
16 August 2012

## Introduction 简介

Go code is organized differently to that of other languages. This post discusses how to name and package the elements of your Go program to best serve its users.

Go代码的组织方式与其他语言的代码不同。本篇文章讨论了如何命名和打包您的 Go 程序中的元素，以便为其用户提供最佳服务。

## Choose good names 选择好的名字

The names you choose affect how you think about your code, so take care when naming your package and its exported identifiers.

你选择的名字会影响你对代码的思考，所以在命名你的包及其导出的标识符时要注意。

A package’s name provides context for its contents. For instance, the [bytes package](https://go.dev/pkg/bytes/) from the standard library exports the `Buffer` type. On its own, the name `Buffer` isn’t very descriptive, but when combined with its package name its meaning becomes clear: `bytes.Buffer`. If the package had a less descriptive name, like `util`, the buffer would likely acquire the longer and clumsier name `util.BytesBuffer`.

一个包的名字为其内容提供了背景。例如，标准库中的字节包导出了Buffer类型。就其本身而言，Buffer这个名字并不具有很强的描述性，但当它与包的名字结合在一起时，其含义就很清楚了：bytes.Buffer。如果这个包有一个描述性不强的名字，比如util，那么缓冲区很可能会获得一个更长、更笨的名字util.BytesBuffer。

Don’t be shy about renaming things as you work. As you spend time with your program you will better understand how its pieces fit together and, therefore, what their names should be. There’s no need to lock yourself into early decisions. (The [gofmt command](https://go.dev/cmd/gofmt/) has a `-r` flag that provides a syntax-aware search and replace, making large-scale refactoring easier.)

在你工作的过程中，不要羞于重命名东西。当你花时间研究你的程序时，你会更好地理解它的各个部分是如何结合在一起的，因此，它们的名字应该是什么。没有必要把自己锁定在早期的决定上。(gofmt命令有一个-r标志，提供了一个语法感知的搜索和替换，使大规模的重构更容易。)

A good name is the most important part of a software interface: the name is the first thing every client of the code will see. A well-chosen name is therefore the starting point for good documentation. Many of the following practices result organically from good naming.

一个好的名字是软件界面中最重要的部分：名字是代码的每个客户首先看到的东西。因此，一个精心选择的名字是好的文档的起点。以下许多做法都是由良好的命名有机地产生的。

## Choose a good import path (make your package “go get”-able) 选择一个好的导入路径（让你的包可以 "去获取"）。

An import path is the string with which users import a package. It specifies the directory (relative to `$GOROOT/src/pkg` or `$GOPATH/src`) in which the package’s source code resides.

导入路径是用户导入软件包时使用的字符串。它指定了包的源代码所在的目录（相对于$GOROOT/src/pkg或$GOPATH/src）。

Import paths should be globally unique, so use the path of your source repository as its base. For instance, the `websocket` package from the `go.net` sub-repository has an import path of `"golang.org/x/net/websocket"`. The Go project owns the path `"github.com/golang"`, so that path cannot be used by another author for a different package. Because the repository URL and import path are one and the same, the `go get` command can fetch and install the package automatically.

导入路径应该是全局唯一的，所以使用你的源码库的路径作为其基础。例如，来自go.net子仓库的websocket包的导入路径为 "golang.org/x/net/websocket"。Go项目拥有 "github.com/golang "这一路径，因此该路径不能被其他作者用于不同的包。因为存储库的URL和导入路径是相同的，所以go get命令可以自动获取并安装该包。

If you don’t use a hosted source repository, choose some unique prefix such as a domain, company, or project name. As an example, the import path of all Google’s internal Go code starts with the string `"google"`.

如果你不使用托管的源码库，选择一些独特的前缀，如域名、公司或项目名称。举个例子，Google所有内部Go代码的导入路径都以 "google "这个字符串开始。

The last element of the import path is typically the same as the package name. For instance, the import path `"net/http"` contains package `http`. This is not a requirement - you can make them different if you like - but you should follow the convention for predictability’s sake: a user might be surprised that import `"foo/bar"` introduces the identifier `quux` into the package name space.

导入路径的最后一个元素通常与包名相同。例如，导入路径 "net/http "包含包 http。这不是一个要求--如果你愿意，你可以让它们不同--但是为了可预测性，你应该遵循这个惯例：用户可能会惊讶于导入 "foo/bar "会将标识符quux引入包名空间。

Sometimes people set `GOPATH` to the root of their source repository and put their packages in directories relative to the repository root, such as `"src/my/package"`. On one hand, this keeps the import paths short (`"my/package"` instead of `"github.com/me/project/my/package"`), but on the other it breaks `go get` and forces users to re-set their `GOPATH` to use the package. Don’t do this.

有时人们会将 GOPATH 设置为他们源码库的根，并将他们的包放在相对于源码库根的目录中，例如 "src/my/package"。一方面，这可以保持导入路径的简短（"my/package "而不是 "github.com/me/project/my/package"），但另一方面，它破坏了go get，迫使用户重新设置GOPATH以使用该包。不要这样做。

## Minimize the exported interface 最小化导出的接口

Your code is likely composed of many small pieces of useful code, and so it is tempting to expose much of that functionality in your package’s exported interface. Resist that urge!

你的代码很可能是由许多有用的小段代码组成的，因此，在你的包的导出接口中暴露出许多功能是很诱人的。请抵制这种冲动

The larger the interface you provide, the more you must support. Users will quickly come to depend on every type, function, variable, and constant you export, creating an implicit contract that you must honor in perpetuity or risk breaking your users' programs. In preparing Go 1 we carefully reviewed the standard library’s exported interfaces and removed the parts we weren’t ready to commit to. You should take similar care when distributing your own libraries.

你提供的接口越大，你必须支持的就越多。用户很快就会依赖你导出的每一个类型、函数、变量和常量，这就形成了一个隐含的契约，你必须永远遵守，否则就有可能破坏用户的程序。在准备Go 1的过程中，我们仔细审查了标准库的导出接口，并删除了我们不准备接受的部分。在发布你自己的库时，你也应该采取类似的谨慎态度。

If in doubt, leave it out!

如果有疑问的话，就不要去管它

## What to put into a package 把什么放进一个包里

It is easy to just throw everything into a “grab bag” package, but this dilutes the meaning of the package name (as it must encompass a lot of functionality) and forces the users of small parts of the package to compile and link a lot of unrelated code.

把所有的东西都扔进一个 "抓包 "的包里是很容易的，但这冲淡了包名的意义（因为它必须包含很多功能），并迫使包中小部分的用户去编译和链接很多不相关的代码。

On the other hand, it is also easy to go overboard in splitting your code into small packages, in which case you will likely become bogged down in interface design, rather than just getting the job done.

另一方面，把你的代码分割成小包也很容易过火，在这种情况下，你很可能会在界面设计上陷入困境，而不是仅仅完成工作。

Look to the Go standard libraries as a guide. Some of its packages are large and some are small. For instance, the [http package](https://go.dev/pkg/net/http/) comprises 17 go source files (excluding tests) and exports 109 identifiers, and the [hash package](https://go.dev/pkg/hash/) consists of one file that exports just three declarations. There is no hard and fast rule; both approaches are appropriate given their context.

请看Go标准库作为指导。它的一些包是大的，一些是小的。例如，http包由17个go源文件组成（不包括测试），输出109个标识符，而hash包由一个文件组成，只输出三个声明。这没有什么硬性规定；考虑到它们的背景，两种方法都是合适的。

With that said, package main is often larger than other packages. Complex commands contain a lot of code that is of little use outside the context of the executable, and often it’s simpler to just keep it all in the one place. For instance, the go tool is more than 12000 lines spread across [34 files](https://go.dev/src/cmd/go/).

说到这里，包main往往比其他包大。复杂的命令包含大量的代码，这些代码在可执行文件的上下文之外没有什么用处，通常把它们都放在一个地方会更简单。例如，go工具有12000多行，分布在34个文件中。

## Document your code 记录你的代码

Good documentation is an essential quality of usable and maintainable code. Read the [Godoc: documenting Go code](https://go.dev/doc/articles/godoc_documenting_go_code.html) article to learn how to write good doc comments.

好的文档是可用和可维护的代码的一个基本质量。阅读Godoc：记录Go代码的文章，了解如何写好文档注释。
