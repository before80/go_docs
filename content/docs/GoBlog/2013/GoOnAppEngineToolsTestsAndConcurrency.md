+++
title = "App Engine上的go：工具、测试和并发"
weight = 22
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go on App Engine: tools, tests, and concurrency - App Engine上的go：工具、测试和并发

https://go.dev/blog/appengine-dec2013

Andrew Gerrand and Johan Euphrosine
13 December 2013

## Background 背景介绍

When we [launched Go for App Engine](https://blog.golang.org/go-and-google-app-engine) in May 2011 the SDK was just a modified version of the Python SDK. At the time, there was no canonical way to build or organize Go programs, so it made sense to take the Python approach. Since then Go 1.0 was released, including the [go tool](https://go.dev/cmd/go/) and a [convention](https://go.dev/doc/code.html) for organizing Go programs.

当我们在2011年5月为App Engine推出Go时，SDK只是Python SDK的一个修改版本。当时，还没有建立或组织Go程序的典型方法，所以采用Python方法是有意义的。此后，Go 1.0发布，包括go工具和组织Go程序的惯例。

In January 2013 we announced [better integration](https://blog.golang.org/the-app-engine-sdk-and-workspaces-gopath) between the Go App Engine SDK and the go tool, promoting the use of conventional import paths in App Engine apps and making it possible to use "go get" to fetch app dependencies.

2013年1月，我们宣布了Go应用引擎SDK和go工具之间更好的整合，促进了App Engine应用程序中传统导入路径的使用，并使使用 "go get "来获取应用程序的依赖项成为可能。

With the recent release of App Engine 1.8.8 we are pleased to announce more improvements to the developer experience for Go on App Engine.

随着最近App Engine 1.8.8的发布，我们很高兴地宣布在App Engine上对Go的开发者体验进行了更多改进。

## The goapp tool - Goapp工具

The Go App Engine SDK now includes the "goapp" tool, an App Engine-specific version of the "go" tool. The new name permits users to keep both the regular "go" tool and the "goapp" tool in their system PATH.

Go App Engine SDK现在包括 "goapp "工具，这是 "go "工具的特定App Engine版本。新名称允许用户在其系统PATH中同时保留常规的 "go "工具和 "goapp "工具。

In addition to the existing "go" tool [commands](https://go.dev/cmd/go/), the "goapp" tool provides new commands for working with App Engine apps. The "[goapp serve](https://developers.google.com/appengine/docs/go/tools/devserver)" command starts the local development server and the "[goapp deploy](https://developers.google.com/appengine/docs/go/tools/uploadinganapp)" command uploads an app to App Engine.

除了现有的 "go "工具命令之外，"goapp "工具还提供了与App Engine应用程序一起工作的新命令。goapp serve "命令启动本地开发服务器，"goapp deploy "命令将应用程序上传到App Engine。

The main advantages offered by the "goapp serve" and "goapp deploy" commands are a simplified user interface and consistency with existing commands like "go get" and "go fmt". For example, to run a local instance of the app in the current directory, run:

goapp serve "和 "goapp deploy "命令提供的主要优势是简化了用户界面，并与现有的命令如 "go get "和 "go fmt "保持一致。例如，要在当前目录下运行一个应用程序的本地实例，请运行：

```shell linenums="1"
$ goapp serve
```

To upload it to App Engine:

要把它上传到App Engine：

```shell linenums="1"
$ goapp deploy
```

You can also specify the Go import path to serve or deploy:

您也可以指定Go导入路径来服务或部署：

```shell linenums="1"
$ goapp serve github.com/user/myapp
```

You can even specify a YAML file to serve or deploy a specific [module](https://developers.google.com/appengine/docs/go/modules/):

您甚至可以指定一个YAML文件来服务或部署一个特定的模块。

```shell linenums="1"
$ goapp deploy mymodule.yaml
```

These commands can replace most uses of `dev_appserver.py` and `appcfg.py`, although the Python tools are still available for their less common uses.

这些命令可以取代dev_appserver.py和appcfg.py的大部分用途，尽管Python工具仍然可以用于它们不太常见的用途。

## Local unit testing 本地单元测试

The Go App Engine SDK now supports local unit testing, using Go’s native [testing package](https://developers.google.com/appengine/docs/go/tools/localunittesting) and the "[go test](https://go.dev/cmd/go/#hdr-Test_packages)" command (provided as "goapp test" by the SDK).

Go App Engine SDK现在支持本地单元测试，使用Go的本地测试包和 "go test "命令（SDK提供的是 "goapp test"）。

Furthermore, you can now write tests that use App Engine services. The [aetest package](https://developers.google.com/appengine/docs/go/tools/localunittesting#Go_Introducing_the_aetest_package) provides an appengine.Context value that delegates requests to a temporary instance of the development server.

此外，您现在可以编写使用App Engine服务的测试。aetest包提供了一个appengine.Context值，将请求委托给开发服务器的一个临时实例。

For more information about using "goapp test" and the aetest package, see the [Local Unit Testing for Go documentation](https://developers.google.com/appengine/docs/go/tools/localunittesting). Note that the aetest package is still in its early days; we hope to add more features over time.

关于使用 "goapp test "和aetest包的更多信息，请参阅Go的本地单元测试文档。请注意，aetest包仍处于早期阶段；我们希望随着时间的推移增加更多的功能。

## Better concurrency support 更好的并发性支持

It is now possible to configure the number of concurrent requests served by each of your app’s dynamic instances by setting the [`max_concurrent_requests`](https://developers.google.com/appengine/docs/go/modules/#max_concurrent_requests) option (available to [Automatic Scaling modules](https://developers.google.com/appengine/docs/go/modules/#automatic_scaling) only).

现在可以通过设置max_concurrent_requests选项（仅适用于自动扩展模块）来配置每个应用程序的动态实例所提供的并发请求的数量。

Here’s an example `app.yaml` file:

下面是一个app.yaml文件的例子：

```
application: maxigopher
version: 1
runtime: go
api_version: go1
automatic_scaling:
  max_concurrent_requests: 100
```

This configures each instance of the app to serve up to 100 requests concurrently (up from the default of 10). You can configure Go instances to serve up to a maximum of 500 concurrent requests.

这将配置应用程序的每个实例最多可以并发提供100个请求（高于默认的10个）。您可以将Go实例配置为最多提供500个并发请求。

This setting allows your instances to handle more simultaneous requests by taking advantage of Go’s efficient handling of concurrency, which should yield better instance utilization and ultimately fewer billable instance hours.

这个设置允许您的实例利用Go对并发的有效处理来处理更多的并发请求，这应该会产生更好的实例利用率，最终减少实例的计费时间。

## Conclusion 总结

With these changes Go on App Engine is more convenient and efficient than ever, and we hope you enjoy the improvements. Please join the [google-appengine-go group](http://groups.google.com/group/google-appengine-go/) to raise questions or discuss these changes with the engineering team and the rest of the community.

有了这些变化，App Engine上的Go比以前更方便、更高效了，我们希望您能享受这些改进。请加入google-appengine-go小组，提出问题或与工程团队和社区其他成员讨论这些变化。
