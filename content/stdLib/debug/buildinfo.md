+++
title = "buildinfo"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/debug/buildinfo@go1.23.0](https://pkg.go.dev/debug/buildinfo@go1.23.0)

Package buildinfo provides access to information embedded in a Go binary about how it was built. This includes the Go toolchain version, and the set of modules used (for binaries built in module mode).

​	`buildinfo`包提供了对嵌入Go二进制文件中的信息的访问，了解它是如何构建的。这包括Go工具链的版本，以及所使用的模块集(对于以模块模式构建的二进制文件)。

Build information is available for the currently running binary in runtime/debug.ReadBuildInfo.

​	在runtime/debug.ReadBuildInfo中，当前运行的二进制文件可以获得构建信息。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type BuildInfo 

``` go 
type BuildInfo = debug.BuildInfo
```

Type alias for build info. We cannot move the types here, since runtime/debug would need to import this package, which would make it a much larger dependency.

​	构建信息的类型别名。我们不能把类型移到这里，因为 runtime/debug 需要导入这个包，这将使它成为一个更大的依赖关系。

#### func Read 

``` go 
func Read(r io.ReaderAt) (*BuildInfo, error)
```

Read returns build information embedded in a Go binary file accessed through the given ReaderAt. Most information is only available for binaries built with module support.

​	`Read` 返回嵌入在通过给定的ReaderAt访问的Go二进制文件中的构建信息。大多数信息只适用于有模块支持的二进制文件。

#### func ReadFile 

``` go 
func ReadFile(name string) (info *BuildInfo, err error)
```

ReadFile returns build information embedded in a Go binary file at the given path. Most information is only available for binaries built with module support.

​	`ReadFile`返回嵌入在给定路径的Go二进制文件中的构建信息。大多数信息只适用于有模块支持的二进制文件。