+++
title = "安装 Beego"
date = 2024-02-04T09:08:56+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/install/install/]({{< ref "/beego/install/installUpgrade" >}})

# Installing Beego 安装 Beego



You can use the classic Go way to install Beego:

​	您可以使用经典的 Go 方式来安装 Beego：

```
go get github.com/beego/beego/v2
```

Frequently asked questions:

​	常问问题：

- git is not installed. Please install git for your system.

  ​	未安装 git。请为您的系统安装 git。

- git https is not accessible. Please config local git and close https validation:

  ​	无法访问 git https。请配置本地区 git 并关闭 https 验证：

  ```
    git config --global http.sslVerify false
  ```

- How can I install Beego offline? There is no good solution for now. We will create packages for downloading and installing for future releases.

  ​	我该如何离线安装 Beego？目前没有好的解决方案。我们将在未来版本中创建用于下载和安装的软件包。

# Upgrading Beego 升级 Beego

You can upgrade Beego through Go command or download and upgrade from source code.

​	您可以通过 Go 命令升级 Beego，或从源代码下载并升级。

- Through Go command (Recommended):

  ​	通过 Go 命令（推荐）：

  ```
    go get -u github.com/beego/beego/v2
  ```

- Through source code: visit `https://github.com/beego/beego/v2` and download the source code. Copy and overwrite to path `$GOPATH/src/github.com/beego/beego/v2`. Then run `go install` to upgrade Beego:

  ​	通过源代码：访问 `https://github.com/beego/beego/v2` 并下载源代码。复制并覆盖到路径 `$GOPATH/src/github.com/beego/beego/v2` 。然后运行 `go install` 来升级 Beego：

  ```
    go install 	github.com/beego/beego/v2
  ```

**Upgrading Prior to 1.0:** The API of Beego is stable after 1.0 and compatible with every upgrade. If you are still using a version lower than 1.0 you may need to configure your parameters based on the latest API.

​	1.0 之前的升级：Beego 的 API 在 1.0 之后是稳定的，并且与每次升级兼容。如果您仍在使用低于 1.0 的版本，则可能需要根据最新的 API 配置您的参数。
