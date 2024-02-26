+++
title = "管理 Go 安装"
weight = 2
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Managing Go installations - 管理 Go 安装

> 原文：[https://go.dev/doc/manage-install](https://go.dev/doc/manage-install)

​	本主题介绍如何在同一台计算机上`安装多个版本的 Go`，以及`如何卸载 Go`。

对于其他关于安装的内容，您可能感兴趣：

- [下载并安装](../InstallingGo) —— 安装和运行的最简单方法。
- [从源码安装Go](../InstallingGoFromSource) —— 如何检出源码，在自己的机器上构建它们，并运行它们。

## 安装多个 Go 版本

​	您可以在同一台机器上安装多个 Go 版本。例如，您可能想在多个 Go 版本上测试您的代码。有关可以通过这种方式安装的版本的列表，请参见[下载页面](https://go.dev/dl/)。

> 注意：要使用这里描述的方法进行安装，您需要安装 [git](https://git-scm.com/)。

​	要安装额外的 Go 版本，请运行 `go install` 命令，指定您要安装的版本的下载位置。下面的例子以`1.10.7`版本为例说明：

```shell
$ go install golang.org/dl/go1.10.7@latest
$ go1.10.7 download
```

​	要用新下载的版本运行`go`命令，请在`go`命令后面加上版本号，如下所示。

```shell
$ go1.10.7 version
go version go1.10.7 linux/amd64
```

​	当您安装了多个版本时，您可以发现每个版本的安装位置，查看版本的`GOROOT`值。例如，运行下面这样的命令。

```
$ go1.10.7 env GOROOT
```

​	要卸载下载的版本，只需删除其 `GOROOT` 环境变量所指定的目录和 `goX.Y.Z` 二进制文件。

## 卸载 Go

您可以使用本主题中描述的步骤从您的系统中删除 Go。

### Linux / macOS / FreeBSD

1. 删除 go 目录

   这通常是 `/usr/local/go`。

2. 从您的 `PATH` 环境变量中删除 Go bin 目录。

   在 `Linux` 和 `FreeBSD` 下，编辑 `/etc/profile` 或 `$HOME/.profile`。如果您用的是 `macOS` 安装 Go，请删除 `/etc/paths.d/go` 文件。

### Windows

删除 Go 的最简单方法是通过 Windows `控制面板`的`添加/删除程序`。

1. 在`控制面板`中，双击`添加/删除程序`。
2. 在`添加/删除程序`中，选择Go编程语言，点击卸载，然后按照提示操作。

您也可以使用命令行工具来删除Go：

- 通过运行以下命令，使用命令行进行卸载。

  ```
  msiexec /x go{{version}}.windows-{{cpu-arch}}.msi /q
  ```

  注意：对Windows使用这个卸载过程将自动删除原始安装时创建的Windows环境变量。