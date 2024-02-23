+++
title = "聚焦外部 Go 库"
weight = 18
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Spotlight on external Go libraries - 聚焦外部 Go 库

> 原文：[https://go.dev/blog/external-libraries](https://go.dev/blog/external-libraries)

Andrew Gerrand
3 June 2011

While the Go authors have been working hard at improving Go’s standard library, the greater community has created a growing ecosystem of external libraries. In this post we look at some popular Go libraries and how they can be used.

当Go作者一直在努力改进Go的标准库时，更大的社区已经创造了一个不断增长的外部库的生态系统。在这篇文章中，我们将介绍一些流行的 Go 库以及如何使用它们。

[Mgo](http://labix.org/mgo) (pronounced "mango") is a MongoDB database driver. [MongoDB](http://www.mongodb.org/) is a [document-oriented database](http://en.wikipedia.org/wiki/Document-oriented_database) with a long list of features suitable for [a broad range of uses](http://www.mongodb.org/display/DOCS/Use%2BCases). The mgo package provides a rich, idiomatic Go API for working with MongoDB, from basic operations such as inserting and updating records to the more advanced [MapReduce](http://www.mongodb.org/display/DOCS/MapReduce) and [GridFS](http://www.mongodb.org/display/DOCS/GridFS) features. Mgo has a bunch of cool features including automated cluster discovery and result pre-fetching - see the [mgo homepage](http://labix.org/mgo) for details and example code. For working with large data sets Go, MongoDB, and mgo are a powerful combination.

Mgo（发音为 "芒果"）是一个MongoDB数据库驱动。MongoDB是一个面向文档的数据库，有一长串适合广泛使用的功能。mgo包为与MongoDB的合作提供了丰富的、习惯性的Go API，从插入和更新记录等基本操作到更高级的MapReduce和GridFS功能。Mgo有很多很酷的功能，包括自动集群发现和结果预取--详情和示例代码见mgo主页。对于处理大型数据集，Go、MongoDB和mgo是一个强大的组合。

[Authcookie](https://github.com/dchest/authcookie) is a web library for generating and verifying user authentication cookies. It allows web servers to hand out cryptographically secure tokens tied to a specific user that will expire after a specified time period. It has a simple API that makes it straightforward to add authentication to existing web applications. See the [README file](https://github.com/dchest/authcookie/blob/master/README.md) for details and example code.

Authcookie 是一个用于生成和验证用户认证 cookie 的网络库。它允许网络服务器发放与特定用户绑定的加密安全令牌，这些令牌将在指定时间段后失效。它有一个简单的API，可以直接将认证添加到现有的Web应用程序中。详情和示例代码见README文件。

[Go-charset](http://code.google.com/p/go-charset) provides support for converting between Go’s standard UTF-8 encoding and a variety of character sets. The go-charset package implements a translating io.Reader and io.Writer so you can wrap existing Readers and Writers (such as network connections or file descriptors), making it easy to communicate with systems that use other character encodings.

Go-charset 为 Go 的标准 UTF-8 编码和各种字符集之间的转换提供支持。Go-charset包实现了一个翻译的io.Reader和io.Writer，因此您可以包装现有的Readers和Writer（如网络连接或文件描述符），使其易于与使用其他字符编码的系统进行通信。

[Go-socket.io](https://github.com/madari/go-socket.io) is a Go implementation of [Socket.IO](http://socket.io/), a client/server API that allows web servers to push messages to web browsers. Depending on the capabilities of the user’s browser, Socket.IO uses the best transport for the connection, be it modern websockets, AJAX long polling, or some [other mechanism](http://socket.io/#transports). Go-socket.io bridges the gap between Go servers and rich JavaScript clients for a wide range of browsers. To get a feel for go-socket.io see the [chat server example](https://github.com/madari/go-socket.io/blob/master/example/example.go).

Go-socket.io是Socket.IO的Go实现，Socket.IO是一个客户端/服务器API，允许网络服务器向网络浏览器推送消息。根据用户浏览器的能力，Socket.IO为连接使用最佳的传输方式，无论是现代websockets、AJAX长轮询，还是其他机制。Go-socket.io为各种浏览器架起了Go服务器和丰富的JavaScript客户端之间的桥梁。要了解go-socket.io的情况，请看聊天服务器的例子。

It’s worth mentioning that these packages are [goinstallable](https://go.dev/cmd/goinstall/). With an up-to-date Go [installation](https://go.dev/doc/install.html) you can install them all with a single command:

值得一提的是，这些软件包都是可以被goinstall的。只要安装了最新的 Go，就可以用一个命令将它们全部安装：

```
goinstall launchpad.net/mgo \
    github.com/dchest/authcookie \
    go-charset.googlecode.com/hg/charset \
    github.com/madari/go-socket.io
```

Once goinstalled, the packages can be imported using those same paths:

一旦goinstalled，可以使用这些相同的路径导入软件包：

```shell linenums="1"
import (
    "launchpad.net/mgo"
    "github.com/dchest/authcookie"
    "go-charset.googlecode.com/hg/charset"
    "github.com/madari/go-socket.io"
)
```

Also, as they are now a part of the local Go system, we can inspect their documentation with [godoc](https://go.dev/cmd/godoc/):

另外，由于它们现在是本地Go系统的一部分，我们可以用godoc检查它们的文档：

```
godoc launchpad.net/mgo Database # see docs for Database type
```

Of course, this is just the tip of the iceberg; there are more great Go libraries listed on the [package dashboard](http://godashboard.appspot.com/package) and many more to come.

当然，这只是冰山一角；还有更多优秀的Go库被列在软件包仪表板上，还有更多的库将会出现。
