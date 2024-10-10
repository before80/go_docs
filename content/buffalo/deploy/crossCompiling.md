+++
title = "交叉编译"
date = 2024-02-04T21:19:43+08:00
weight =2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/deploy/cross-compiling/]({{< ref "/buffalo/deploy/crossCompiling" >}})

# Cross-compiling a Buffalo application 交叉编译 Buffalo 应用程序 

Just like another Go application, you can cross-compile a Buffalo application. This means it’s possible to develop your app on a Mac, and compile it for a Linux target on your Mac.

​	就像其他 Go 应用程序一样，您可以交叉编译 Buffalo 应用程序。这意味着您可以在 Mac 上开发应用程序，并在 Mac 上为 Linux 目标编译它。

## GOOS and GOARCH GOOS 和 GOARCH 

The Go toolchain supports cross-compilation out of the box. You just need to provide the `GOOS` and `GOARCH` env variables.

​	Go 工具链开箱即用地支持交叉编译。您只需提供 `GOOS` 和 `GOARCH` 环境变量。

- `GOOS` sets the target OS (e.g. linux, windows, etc.)
  `GOOS` 设置目标操作系统（例如 linux、windows 等）
- `GOARCH` sets the target CPU architecture (e.g. amd64, 386, etc.)
  `GOARCH` 设置目标 CPU 架构（例如 amd64、386 等）

You can find the list of supported targets here: https://golang.org/doc/install/source#environment

​	您可以在此处找到受支持目标的列表：https://golang.org/doc/install/source#environment

## Examples 示例 

### Build for AMD64 Linux 针对 AMD64 Linux 构建 

```bash
$ GOOS=linux GOARCH=amd64 buffalo build
```

### Build for ARM64 Linux 针对 ARM64 Linux 构建 

```bash
$ GOOS=linux GOARCH=arm64 buffalo build
```

### Build for i386 Windows 适用于 i386 Windows 的构建 

```bash
$ GOOS=windows GOARCH=386 buffalo build
```
