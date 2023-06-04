+++
title = "plugin"
linkTitle = "plugin"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# plugin

https://pkg.go.dev/plugin@go1.20.1





Package plugin implements loading and symbol resolution of Go plugins.

plugin 包实现了Go插件的加载和符号解析。

A plugin is a Go main package with exported functions and variables that has been built with:

plugin 是一个带有导出函数和变量的 Go 主包，它是通过以下方式构建的：

```
go build -buildmode=plugin
```

When a plugin is first opened, the init functions of all packages not already part of the program are called. The main function is not run. A plugin is only initialized once, and cannot be closed.

当一个插件第一次被打开时，所有还没有成为程序一部分的包的初始函数被调用。主函数不会被运行。一个插件只被初始化一次，并且不能被关闭。

#### Warnings 

The ability to dynamically load parts of an application during execution, perhaps based on user-defined configuration, may be a useful building block in some designs. In particular, because applications and dynamically loaded functions can share data structures directly, plugins may enable very high-performance integration of separate parts.

在执行过程中动态加载应用程序的部分的能力，也许是基于用户定义的配置，在某些设计中可能是一个有用的构建块。特别是，由于应用程序和动态加载的函数可以直接共享数据结构，插件可以使独立的部分得到非常高性能的整合。

However, the plugin mechanism has many significant drawbacks that should be considered carefully during the design. For example:

然而，插件机制有许多明显的缺点，在设计过程中应该仔细考虑。比如说：

- Plugins are currently supported only on Linux, FreeBSD, and macOS, making them unsuitable for applications intended to be portable.插件目前只在Linux、FreeBSD和macOS上支持，这使得它们不适合用于打算移植的应用程序。
- Applications that use plugins may require careful configuration to ensure that the various parts of the program be made available in the correct location in the file system (or container image). By contrast, deploying an application consisting of a single static executable is straightforward. 使用插件的应用程序可能需要仔细配置，以确保程序的各个部分在文件系统(或容器镜像)中的正确位置可用。相比之下，部署一个由单个静态可执行文件组成的应用程序是很直接的。
- Reasoning about program initialization is more difficult when some packages may not be initialized until long after the application has started running. 当一些包可能在应用程序开始运行很久后才被初始化时，对程序初始化的推理就比较困难。
- Bugs in applications that load plugins could be exploited by an an attacker to load dangerous or untrusted libraries.
- 攻击者可以利用加载插件的应用程序中的漏洞来加载危险的或不受信任的库。
- Runtime crashes are likely to occur unless all parts of the program (the application and all its plugins) are compiled using exactly the same version of the toolchain, the same build tags, and the same values of certain flags and environment variables.
- 除非程序的所有部分(应用程序和它的所有插件)都是使用完全相同的工具链版本、相同的构建标签以及某些标志和环境变量的相同值来编译的，否则就有可能发生运行时崩溃。
- Similar crashing problems are likely to arise unless all common dependencies of the application and its plugins are built from exactly the same source code.
- 除非应用程序及其插件的所有公共依赖都是由完全相同的源代码构建的，否则很可能会出现类似的崩溃问题。
- Together, these restrictions mean that, in practice, the application and its plugins must all be built together by a single person or component of a system. In that case, it may be simpler for that person or component to generate Go source files that blank-import the desired set of plugins and then compile a static executable in the usual way.
- 这些限制加在一起意味着，在实践中，应用程序及其插件必须由一个人或一个系统的组成部分共同构建。在这种情况下，对该人或组件来说，生成Go源文件以空白方式导入所需的插件集，然后以通常的方式编译静态可执行文件可能更简单。

For these reasons, many users decide that traditional interprocess communication (IPC) mechanisms such as sockets, pipes, remote procedure call (RPC), shared memory mappings, or file system operations may be more suitable despite the performance overheads.

由于这些原因，许多用户决定采用传统的进程间通信(IPC)机制，如套接字、管道、远程过程调用(RPC)、共享内存映射或文件系统操作，尽管有性能开销，但可能更适合。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Plugin 

``` go 
type Plugin struct {
	// contains filtered or unexported fields
}
```

Plugin is a loaded Go plugin.

Plugin是一个加载的Go插件。

#### func Open 

``` go 
func Open(path string) (*Plugin, error)
```

Open opens a Go plugin. If a path has already been opened, then the existing *Plugin is returned. It is safe for concurrent use by multiple goroutines.

开启一个Go插件。如果一个路径已经被打开，那么就会返回现有的*Plugin。它对多个goroutine的并发使用是安全的。

#### (*Plugin) Lookup 

``` go 
func (p *Plugin) Lookup(symName string) (Symbol, error)
```

Lookup searches for a symbol named symName in plugin p. A symbol is any exported variable or function. It reports an error if the symbol is not found. It is safe for concurrent use by multiple goroutines.

Lookup在插件p中搜索一个名为symName的符号。如果没有找到该符号，它会报告一个错误。它对多个goroutine的并发使用是安全的。

### type Symbol 

``` go 
type Symbol any
```

A Symbol is a pointer to a variable or function.

一个符号是一个指向变量或函数的指针。

For example, a plugin defined as

例如，一个插件定义为

``` go 
package main

import "fmt"

var V int

func F() { fmt.Printf("Hello, number %d\n", V) }
```

may be loaded with the Open function and then the exported package symbols V and F can be accessed

可以用Open函数加载，然后可以访问导出的包符号V和F

```
p, err := plugin.Open("plugin_name.so")
if err != nil {
	panic(err)
}
v, err := p.Lookup("V")
if err != nil {
	panic(err)
}
f, err := p.Lookup("F")
if err != nil {
	panic(err)
}
*v.(*int) = 7
f.(func())() // prints "Hello, number 7"
```