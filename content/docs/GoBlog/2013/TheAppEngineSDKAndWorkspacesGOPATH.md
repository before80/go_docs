+++
title = "The App Engine SDK和工作区(GOPATH)"
weight = 20
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# The App Engine SDK and workspaces (GOPATH) - The App Engine SDK和工作区(GOPATH)

https://go.dev/blog/appengine-gopath

Andrew Gerrand
9 January 2013

## Introduction 简介

When we released Go 1 we introduced the [go tool](https://go.dev/cmd/go/) and, with it, the concept of workspaces. Workspaces (specified by the GOPATH environment variable) are a convention for organizing code that simplifies fetching, building, and installing Go packages. If you’re not familiar with workspaces, please read [this article](https://go.dev/doc/code.html) or watch [this screencast](http://www.youtube.com/watch?v=XCsL89YtqCs) before reading on.

当我们发布Go 1的时候，我们引入了go工具，并随之引入了工作空间的概念。工作空间（由GOPATH环境变量指定）是一种组织代码的惯例，可以简化Go包的获取、构建和安装。如果您不熟悉工作空间，在继续阅读之前，请阅读这篇文章或观看这个截屏。

Until recently, the tools in the App Engine SDK were not aware of workspaces. Without workspaces the "[go get](https://go.dev/cmd/go/#hdr-Download_and_install_packages_and_dependencies)" command cannot function, and so app authors had to install and update their app dependencies manually. It was a pain.

直到最近，App Engine SDK中的工具还没有意识到工作空间的存在。没有工作空间，"go get "命令就不能发挥作用，因此应用程序作者不得不手动安装和更新他们的应用程序依赖。这是一种痛苦。

This has all changed with version 1.7.4 of the App Engine SDK. The [dev_appserver](https://developers.google.com/appengine/docs/go/tools/devserver) and [appcfg](https://developers.google.com/appengine/docs/go/tools/uploadinganapp) tools are now workspace-aware. When running locally or uploading an app, these tools now search for dependencies in the workspaces specified by the GOPATH environment variable. This means you can now use "go get" while building App Engine apps, and switch between normal Go programs and App Engine apps without changing your environment or habits.

在App Engine SDK的1.7.4版本中，这一切都改变了。dev_appserver和appcfg工具现在可以感知工作区。在本地运行或上传应用程序时，这些工具现在会在GOPATH环境变量指定的工作空间中搜索依赖关系。这意味着您现在可以在构建App Engine应用程序时使用 "go get"，并在正常的Go程序和App Engine应用程序之间切换，而无需改变您的环境或习惯。

For example, let’s say you want to build an app that uses OAuth 2.0 to authenticate with a remote service. A popular OAuth 2.0 library for Go is the [oauth2](https://godoc.org/golang.org/x/oauth2) package, which you can install to your workspace with this command:

例如，假设您想构建一个使用OAuth 2.0来验证远程服务的应用程序。一个流行的OAuth 2.0库是oauth2包，您可以用这个命令把它安装到您的工作区：

```shell linenums="1"
go get golang.org/x/oauth2
```

When writing your App Engine app, import the oauth package just as you would in a regular Go program:

在编写App Engine应用程序时，就像在普通Go程序中那样导入oauth包：

```go linenums="1"
import "golang.org/x/oauth2"
```

Now, whether running your app with the dev_appserver or deploying it with appcfg, the tools will find the oauth package in your workspace. It just works.

现在，无论是用dev_appserver运行您的应用，还是用appcfg部署它，工具都会在您的工作区找到oauth包。它就这样工作了。

## Hybrid stand-alone/App Engine apps 混合型独立应用/应用引擎应用

The Go App Engine SDK builds on Go’s standard [net/http](https://go.dev/pkg/net/http/) package to serve web requests and, as a result, many Go web servers can be run on App Engine with only a few changes. For example, [godoc](https://go.dev/cmd/godoc/) is included in the Go distribution as a stand-alone program, but it can also run as an App Engine app (godoc serves [golang.org](https://go.dev/) from App Engine).

Go App Engine SDK建立在Go的标准net/http包之上，用于服务网络请求，因此，许多Go网络服务器只需做一些改变就可以在App Engine上运行。例如，godoc作为一个独立的程序包含在Go发行版中，但它也可以作为一个App Engine应用程序运行（godoc从App Engine为golang.org提供服务）。

But wouldn’t it be nice if you could write a program that is both a stand-alone web server and an App Engine app? By using [build constraints](https://go.dev/pkg/go/build/#hdr-Build_Constraints), you can.

但是如果您能写一个既是独立的Web服务器又是App Engine应用的程序，那不是很好吗？通过使用构建约束，您可以做到。

Build constraints are line comments that determine whether a file should be included in a package. They are most often used in code that handles a variety of operating systems or processor architectures. For instance, the [path/filepath](https://go.dev/pkg/path/filepath/) package includes the file [symlink.go](https://go.dev/src/pkg/path/filepath/symlink.go), which specifies a build constraint to ensure that it is not built on Windows systems (which do not have symbolic links):

构建约束是决定一个文件是否应该包含在一个包中的行注释。它们最常被用于处理各种操作系统或处理器架构的代码中。例如，path/filepath包包括文件symlink.go，它指定了一个构建约束，以确保它不在Windows系统（没有符号链接）上构建：

```
// +build !windows
```

The App Engine SDK introduces a new build constraint term: "appengine". Files that specify

App Engine SDK引入了一个新的构建约束术语。"appengine"。文件如果指定

```
// +build appengine
```

will be built by the App Engine SDK and ignored by the go tool. Conversely, files that specify

的文件将由App Engine SDK构建，而被Go工具忽略。反之，如果文件指定

```
// +build !appengine
```

are ignored by the App Engine SDK, while the go tool will happily build them.

的文件会被App Engine SDK忽略，而go工具会很高兴地构建它们。

The [goprotobuf](http://code.google.com/p/goprotobuf/) library uses this mechanism to provide two implementations of a key part of its encode/decode machinery: [pointer_unsafe.go](http://code.google.com/p/goprotobuf/source/browse/proto/pointer_unsafe.go) is the faster version that cannot be used on App Engine because it uses the [unsafe package](https://go.dev/pkg/unsafe/), while [pointer_reflect.go](http://code.google.com/p/goprotobuf/source/browse/proto/pointer_reflect.go) is a slower version that avoids unsafe by using the [reflect package](https://go.dev/pkg/reflect/) instead.

goprotobuf库使用这种机制为其编码/解码机制的一个关键部分提供了两种实现：pointer_unsafe.go是较快的版本，因为它使用了不安全包而不能在App Engine上使用，而pointer_reflect.go是较慢的版本，通过使用reflect包而避免了不安全。

Let’s take a simple Go web server and turn it into a hybrid app. This is main.go:

让我们拿一个简单的Go网络服务器，把它变成一个混合应用程序。这就是main.go：

```go linenums="1"
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello!")
}
```

Build this with the go tool and you’ll get a stand-alone web server executable.

用go工具构建这个，您会得到一个独立的Web服务器可执行文件。

The App Engine infrastructure provides its own main function that runs its equivalent to ListenAndServe. To convert main.go to an App Engine app, drop the call to ListenAndServe and register the handler in an init function (which runs before main). This is app.go:

App Engine基础设施提供了它自己的main函数，运行它相当于ListenAndServe。为了将main.go转换为App Engine应用程序，放弃对ListenAndServe的调用，并在init函数（运行在main之前）中注册处理程序。这就是app.go：

```go linenums="1"
package main

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello!")
}
```

To make this a hybrid app, we need to split it into an App Engine-specific part, an stand-alone binary-specific part, and the parts common to both versions. In this case, there is no App Engine-specific part, so we split it into just two files:

为了使这成为一个混合应用程序，我们需要把它分成一个App Engine特定的部分，一个独立的二进制特定的部分，以及两个版本的共同部分。在本例中，没有针对App Engine的部分，所以我们把它分成两个文件。

app.go specifies and registers the handler function. It is identical to the code listing above, and requires no build constraints as it should be included in all versions of the program.

app.go指定并注册了处理函数。它与上面的代码列表相同，不需要构建约束，因为它应该包含在所有版本的程序中。

main.go runs the web server. It includes the "!appengine" build constraint, as it must only be included when building the stand-alone binary.

main.go运行网络服务器。它包括"!appengine "构建约束，因为只有在构建独立的二进制文件时才必须包括它。

```go linenums="1"
// +build !appengine

package main

import "net/http"

func main() {
    http.ListenAndServe("localhost:8080", nil)
}
```

To see a more complex hybrid app, take a look at the [present tool](https://godoc.org/golang.org/x/tools/present).

要看一个更复杂的混合应用程序，请看一下本工具。

## Conclusions 结论

We hope these changes will make it easier to work on apps with external dependencies, and to maintain code bases that contain both stand-alone programs and App Engine apps.

我们希望这些变化能够使我们更容易地处理具有外部依赖性的应用程序，并维护包含独立程序和App Engine应用程序的代码库。
