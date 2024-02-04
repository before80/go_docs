+++
title = "FAQ"
date = 2024-02-04T09:41:30+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/faq/]({{< ref "/beego/faq" >}})

# FAQ

1. Can’t find the template files or configuration files or nil pointer error?

   ​	找不到模板文件或配置文件或空指针错误？

   It may be because you used `go run main.go` to run your application. `go run` will compile the file and put it into a tmp folder to run it. But Beego needs the static files, templates and config files. So you need to use `go build` and run the application by `./app`. Or you can use `bee run app` to run your application.

   ​	这可能是因为您使用 `go run main.go` 运行应用程序。 `go run` 将编译文件并将其放入 tmp 文件夹中以运行它。但 Beego 需要静态文件、模板和配置文件。因此，您需要使用 `go build` 并通过 `./app` 运行应用程序。或者，您可以使用 `bee run app` 运行应用程序。

2. Can Beego be used for production?

   ​	Beego 能用于生产吗？

   Yes. Beego has been used in production. E.g.: SNDA’s CDN system, 360 search API, Bmob mobile cloud API, weico backend API etc. They are all high concurrence and high performance applications.

   ​	可以。Beego 已用于生产。例如：SNDA 的 CDN 系统、360 搜索 API、Bmob 移动云 API、weico 后端 API 等。它们都是高并发和高性能的应用程序。

3. Will the future upgrades affect the API I am using right now?

   ​	未来的升级是否会影响我现在正在使用的 API？

   Beego is keeping the stable API since version 0.1. Many applications upgraded to the latest Beego easily. We will try to keep the API stable in the future.

   ​	自 0.1 版本以来，Beego 一直保持稳定的 API。许多应用程序轻松升级到最新的 Beego。我们将在未来努力保持 API 的稳定性。

4. Will Beego keep developing?

   ​	Beego 会继续开发吗？

   Many people are worried about open source projects that stop developing. We have four people who are contributing to the code. We can keep making Beego better and better.

   ​	许多人担心停止开发的开源项目。我们有四个人在为代码做出贡献。我们可以让 Beego 变得越来越好。

5. Why I got “github.com/beego/beego/v2” package not found error?

   ​	为什么我收到“找不到 github.com/beego/beego/v2”包错误？

   In BeegoV2, we are using go mod. So you must enable go module feature in your environment. In general, you should set `GO111MODULE=on`.

   ​	在 BeegoV2 中，我们使用 go mod。因此您必须在您的环境中启用 go module 功能。通常，您应该设置 `GO111MODULE=on` 。

6. Why I always got i/o timeout when I run `go get github.com/beego/beego/v2`?

   ​	为什么我在运行 `go get github.com/beego/beego/v2` 时总是超时？

   It means that your network has some problem. Sometimes it was caused by the firewall. If you are in China, this is a common case, and you could set `GOPROXY`, for example: `export GORPOXY=https://goproxy.cn"`

   ​	这意味着您的网络存在一些问题。有时这是由防火墙引起的。如果您在中国，这是一个常见的情况，您可以设置 `GOPROXY` ，例如： `export GORPOXY=https://goproxy.cn"`
