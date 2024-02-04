+++
title = "upgradeBeego"
date = 2024-02-04T09:36:06+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/intro/upgrade/]({{< ref "/beego/intro/upgradeBeego" >}})

# Upgrade Beego 升级 Beego



## Upgrade beego 2.0.0 升级 beego 2.0.0

Install the latest bee tool `go get -u github.com/beego/bee/v2` Upgrading Beego `go get -u github.com/beego/beego/v2`

​	安装最新的 bee 工具 `go get -u github.com/beego/bee/v2` 升级 Beego `go get -u github.com/beego/beego/v2`

Running: `bee fix -t 2`

​	正在运行： `bee fix -t 2`

If you are working on Windows platform, please run this command in WSL.

​	如果您在 Windows 平台上工作，请在 WSL 中运行此命令。

## Upgrade to beego 1.6.0 升级到 beego 1.6.0

Install the latest bee tool `go get -u github.com/beego/bee` Upgrade Beego `go get -u github.com/astaxie/beego@v1.6.0`

​	安装最新的 bee 工具 `go get -u github.com/beego/bee` 升级 Beego `go get -u github.com/astaxie/beego@v1.6.0`

Running : `bee fix`

​	正在运行： `bee fix`

## Upgrade to beego 1.4.2 升级到 beego 1.4.2

Change GetInt to GetInt64

​	将 GetInt 更改为 GetInt64

## Upgrade to beego 1.3 升级到 beego 1.3

1. `AddFilter` method was removed, and you could update your code from：

   ​	 `AddFilter` 方法已删除，您可以将代码从以下内容更新：

   ```
    beego.AddFilter("/user/*", "BeforeRouter", cpt.Handler)
   ```

   To

   ```
    beego.InsertFilter("/user/*", beego.BeforeRouter, cpt.Handler)
   ```

2. beego.AfterStatic was removed，please using beego.BeforeRouter

   ​	beego.AfterStatic 已删除，请使用 beego.BeforeRouter
