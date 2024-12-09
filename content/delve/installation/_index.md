+++
title = "安装"
date = 2024-12-09T08:00:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/installation/README.md](https://github.com/go-delve/delve/blob/master/Documentation/installation/README.md)
>
> 收录该文档时间： `2024-12-09T08:00:34+08:00`

# Installation



The following instructions are known to work on Linux, macOS, Windows and FreeBSD.

​	以下说明适用于 Linux、macOS、Windows 和 FreeBSD。

With Go version 1.16 or later:

​	使用 Go 版本 1.16 或更高版本：

```
# Install the latest release:
$ go install github.com/go-delve/delve/cmd/dlv@latest

# Install from tree head:
$ go install github.com/go-delve/delve/cmd/dlv@master

# Install a specific version or pseudo-version:
$ go install github.com/go-delve/delve/cmd/dlv@v1.7.3
$ go install github.com/go-delve/delve/cmd/dlv@v1.7.4-0.20211208103735-2f13672765fe
```



See [Versions](https://go.dev/ref/mod#versions) and [Pseudo-versions](https://go.dev/ref/mod#pseudo-versions) for how to format the version suffixes.

​	请参阅 [版本](https://go.dev/ref/mod#versions) 和 [伪版本](https://go.dev/ref/mod#pseudo-versions) 了解如何格式化版本后缀。

Alternatively, clone the git repository and build:

​	或者，克隆 Git 仓库并构建：

```
$ git clone https://github.com/go-delve/delve
$ cd delve
$ go install github.com/go-delve/delve/cmd/dlv
```



See `go help install` for details on where the `dlv` executable is saved.

​	有关 `dlv` 可执行文件保存位置的详细信息，请参阅 `go help install`。

If during the install step you receive an error similar to this:

​	如果在安装步骤中收到如下错误：

```
found packages native (proc.go) and your_operating_system_and_architecture_combination_is_not_supported_by_delve (support_sentinel.go) in /home/pi/go/src/github.com/go-delve/delve/pkg/proc/native
```



It means that your combination of operating system and CPU architecture is not supported, check the output of `go version`.

​	这意味着您的操作系统和 CPU 架构组合不受支持，请检查 `go version` 的输出。

## macOS 注意事项 considerations



On macOS make sure you also install the command line developer tools:

​	在 macOS 上，请确保您还安装了命令行开发者工具：



```
$ xcode-select --install
```



If you didn't enable Developer Mode using Xcode you will be asked to authorize the debugger every time you use it. To enable Developer Mode and only have to authorize once per session use:

​	如果您没有通过 Xcode 启用开发者模式，每次使用调试器时都会要求授权。要启用开发者模式并在每次会话中只需授权一次，请使用：

```
sudo /usr/sbin/DevToolsSecurity -enable
```



You might also need to add your user to the developer group:

​	您可能还需要将您的用户添加到开发者组：

```
sudo dscl . append /Groups/_developer GroupMembership $(whoami)
```



## 编译 macOS 本地后端 Compiling macOS native backend



You do not need the macOS native backend and it [has known problems](https://github.com/go-delve/delve/issues/1112). If you still want to build it:

​	您不需要使用 macOS 本地后端，并且其[已知存在问题](https://github.com/go-delve/delve/issues/1112)。如果您仍希望构建它：

1. Run `xcode-select --install`
2. On macOS 10.14 manually install the legacy include headers by running  `/Library/Developer/CommandLineTools/Packages/macOS_SDK_headers_for_macOS_10.14.pkg`在 macOS 10.14 上，通过运行 `/Library/Developer/CommandLineTools/Packages/macOS_SDK_headers_for_macOS_10.14.pkg`手动安装旧版头文件
3. Clone the repo into `$GOPATH/src/github.com/go-delve/delve` 将代码仓库克隆到 `$GOPATH/src/github.com/go-delve/delve`
4. Run `make install` in that directory (on some versions of macOS this requires being root, the first time you run it, to install a new certificate) 在该目录中运行 `make install`（在某些版本的 macOS 上，这需要以 root 身份运行首次安装新证书）

The makefile will take care of creating and installing a self-signed certificate automatically.

​	Makefile 会自动处理创建和安装自签名证书的过程。
