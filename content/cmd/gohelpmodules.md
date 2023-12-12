+++
title = "go help modules"
date = 2023-12-12T14:13:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

​	

A module is a collection of packages that are released, versioned, and distributed together. Modules may be downloaded directly from version control repositories or from module proxy servers.

​	模块是一组一起发布、进行版本管理和分发的包。可以直接从版本控制存储库或模块代理服务器下载模块。

For a series of tutorials on modules, see https://golang.org/doc/tutorial/create-module.

​	有关模块的一系列教程，请参见 https://golang.org/doc/tutorial/create-module。

For a detailed reference on modules, see https://golang.org/ref/mod.

​	有关模块的详细参考，请参见 https://golang.org/ref/mod。

By default, the go command may download modules from https://proxy.golang.org. It may authenticate modules using the checksum database at https://sum.golang.org. Both services are operated by the Go team at Google.

​	默认情况下，go 命令可以从 [https://proxy.golang.org](https://proxy.golang.org/) 下载模块。它可以使用 [https://sum.golang.org](https://sum.golang.org/) 的校验和数据库对模块进行身份验证。这两项服务都由谷歌的 Go 团队运营。

The privacy policies for these services are available at https://proxy.golang.org/privacy and https://sum.golang.org/privacy, respectively.

​	这些服务的隐私政策可在 https://proxy.golang.org/privacy 和 https://sum.golang.org/privacy 找到。

The go command's download behavior may be configured using GOPROXY, GOSUMDB, GOPRIVATE, and other environment variables. See 'go help environment' and https://golang.org/ref/mod#private-module-privacy for more information.

​	可以使用 GOPROXY、GOSUMDB、GOPRIVATE 和其他环境变量配置 go 命令的下载行为。有关更多信息，请参阅 'go help environment' 和 https://golang.org/ref/mod#private-module-privacy。
