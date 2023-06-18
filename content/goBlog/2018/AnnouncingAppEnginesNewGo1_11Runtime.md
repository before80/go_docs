+++
title = "宣布App Engine的新 go 1.11运行时"
weight = 5
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Announcing App Engine’s New Go 1.11 Runtime - 宣布App Engine的新 go 1.11运行时

https://go.dev/blog/appengine-go111

Eno Compton and Tyler Bui-Palsulich
16 October 2018

[App Engine](https://cloud.google.com/appengine/) launched [experimental support for Go](https://blog.golang.org/go-and-google-app-engine) in 2011. In the subsequent years, the Go community has grown significantly and has settled on idiomatic patterns for cloud-based applications. Today, Google Cloud is [announcing a new Go 1.11 runtime](https://cloud.google.com/blog/products/application-development/go-1-11-is-now-available-on-app-engine) for the App Engine standard environment that provides all the power of App Engine—things like paying only for what you use, automatic scaling, and managed infrastructure—while supporting idiomatic Go.

App Engine在2011年推出了对Go的实验性支持。在随后的几年里，Go社区有了很大的发展，并为基于云的应用程序确定了成语模式。今天，谷歌云宣布为App Engine标准环境提供新的Go 1.11运行时，该运行时提供App Engine的所有功能--例如只为您使用的东西付费、自动扩展和管理基础设施--同时支持成语Go。

Starting with Go 1.11, Go on App Engine has no limits on application structure, supported packages, `context.Context` values, or HTTP clients. Write your Go application however you prefer, add an `app.yaml` file, and your app is ready to deploy on App Engine. [Specifying Dependencies](https://cloud.google.com/appengine/docs/standard/go111/specifying-dependencies) describes how the new runtime supports [vendoring](https://go.dev/cmd/go/#hdr-Vendor_Directories) and [modules](https://go.dev/doc/go1.11#modules) (experimental) for dependency management.

从Go 1.11开始，App Engine上的Go在应用结构、支持的包、context.Context值或HTTP客户端方面都没有限制。以您喜欢的方式编写您的Go应用程序，添加一个app.yaml文件，您的应用程序就可以在App Engine上部署了。指定依赖项描述了新的运行时如何支持vendoring和模块（实验性）的依赖项管理。

Along with [Cloud Functions support for Go](https://twitter.com/kelseyhightower/status/1035278586754813952) (more on that in a future post), App Engine provides a compelling way to run Go code on Google Cloud Platform (GCP) with no concern for the underlying infrastructure.

随着Cloud Functions对Go的支持（在未来的文章中会有更多介绍），App Engine为在Google云平台（GCP）上运行Go代码提供了一种引人注目的方式，而不必担心底层基础设施。

Let’s take a look at creating a small application for App Engine. For the example here, we assume a `GOPATH`-based workflow, although Go modules have [experimental support](https://cloud.google.com/appengine/docs/standard/go111/specifying-dependencies) as well.

让我们来看看如何为App Engine创建一个小的应用程序。在这里的例子中，我们假设是基于GOPATH的工作流程，尽管Go模块也有实验性的支持。

First, you create the application in your `GOPATH`:

首先，您在您的GOPATH中创建应用程序：

```go
// This server can run on App Engine.
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", hello)

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, 世界"))
}
```

The code contains an idiomatic setup for a small HTTP server that responds with "Hello, 世界." If you have previous App Engine experience, you’ll notice the absence of any call to `appengine.Main()`, which is now entirely optional. Furthermore, the application code is completely portable—there are no ties to the infrastructure that your application is deployed on.

这段代码包含了一个小型HTTP服务器的习惯性设置，它响应的是 "Hello, 世界"。如果您以前有App Engine的经验，您会注意到没有对appengine.Main()的任何调用，现在它完全是可选的。此外，应用程序的代码是完全可移植的--与您的应用程序所部署的基础设施没有任何联系。

If you need to use external dependencies, you can add those dependencies to a `vendor` directory or to a `go.mod` file, both of which the new runtime supports.

如果您需要使用外部依赖，您可以将这些依赖添加到供应商目录或go.mod文件中，新的运行时支持这两种方式。

With the application code complete, create an `app.yaml` file to specify the runtime:

应用程序代码完成后，创建一个app.yaml文件来指定运行时：

```
runtime: go111
```

Finally, set your machine up with a Google Cloud Platform account:

最后，在您的机器上设置一个谷歌云平台账户：

- Create an account with [GCP](https://cloud.google.com/).
- [Create a project](https://cloud.google.com/resource-manager/docs/creating-managing-projects).
- Install the [Cloud SDK](https://cloud.google.com/sdk/) on your system.
- 在GCP创建一个账户。
  创建一个项目。
  在您的系统上安装Cloud SDK。

With all the setup complete, you can deploy using one command:

所有设置完成后，您可以使用一个命令进行部署：

```
gcloud app deploy
```

We think Go developers will find the new Go 1.11 runtime for App Engine an exciting addition to the available options to run Go applications. There is a [free tier](https://cloud.google.com/free/). Check out the [getting started guide](https://cloud.google.com/appengine/docs/standard/go111/building-app/) or the [migration guide](https://cloud.google.com/appengine/docs/standard/go111/go-differences) and deploy an app to the new runtime today!

我们认为Go开发者会发现App Engine的新Go 1.11运行时是对运行Go应用程序可用选项的一个令人兴奋的补充。有一个免费层。请查看入门指南或迁移指南，并在今天将应用程序部署到新的运行时中。
