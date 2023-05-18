+++
title = "Third-第三方库：goprotobuf及其他"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Third-party libraries: goprotobuf and beyond - 第三方库：goprotobuf及其他

https://go.dev/blog/protobuf

Andrew Gerrand
20 April 2010

2010 年 4 月 20 日

​	在 3 月 24 日，Rob Pike 宣布了 [goprotobuf](http://code.google.com/p/goprotobuf/)，即 Google 数据交换格式 [Protocol Buffers](http://code.google.com/apis/protocolbuffers/docs/overview.html) 的 Go 语言绑定。与 C++、Java 和 Python 一样，Go 语言提供了官方的 protobuf 实现，这是实现现有系统与使用 Go 语言构建的系统互操作性的重要里程碑。

​	goprotobuf 项目包含两部分：一个“协议编译器插件”，用于生成 Go 源文件，一旦编译，就可以访问和管理协议缓冲区；以及一个 Go 包，用于实现编码（marshaling 编组）、解码（unmarshaling 解组）和访问协议缓冲区的运行时支持。

​	要使用 goprotobuf，首先需要安装 Go 和 [protobuf](http://code.google.com/p/protobuf/)。然后可以使用 [goinstall](https://go.dev/cmd/goinstall/)安装“proto”包：

```shell
goinstall goprotobuf.googlecode.com/hg/proto
```

然后安装 protobuf 编译器插件：

```shell
cd $GOROOT/src/pkg/goprotobuf.googlecode.com/hg/compiler
make install
```

更多详情请参见项目的 [README](http://code.google.com/p/goprotobuf/source/browse/README) 文件。

​	这是越来越多的第三方 [Go 项目](http://godashboard.appspot.com/package)之一。自 goprotobuf 的发布以来，X Go 绑定已从标准库分离到 [x-go-binding](http://code.google.com/p/x-go-binding/) 项目中，并且已经开始进行 [Freetype](http://www.freetype.org/)移植，即 [freetype-go](http://code.google.com/p/freetype-go/)。其他受欢迎的第三方项目包括轻量级的 Web 框架 [web.go](http://github.com/hoisie/web.go)，以及 Go GTK 绑定 [gtk-go](http://github.com/mattn/go-gtk)。

​	我们希望通过开源社区的开发来鼓励其他有用的软件包的开发。如果您正在开发某些东西，请不要独自承受，让我们通过我们的邮件列表 g[golang-nuts](http://groups.google.com/group/golang-nuts) 知道。
