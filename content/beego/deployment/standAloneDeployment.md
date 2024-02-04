+++
title = "独立部署"
date = 2024-02-04T09:12:43+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/deploy/beego/](https://beego.wiki/docs/deploy/beego/)

# Stand alone Deployment 独立部署



This will run application in the backend as a daemon.

&zeroWidthSpace;这将在后台以守护进程的形式运行应用程序。

## linux

In Linux we can use `nohup` command to run the application in the backend:

&zeroWidthSpace;在 Linux 中，我们可以使用 `nohup` 命令在后台运行应用程序：

```
nohup ./beepkg &
```

Your application is running in the keep process of Linux.

&zeroWidthSpace;您的应用程序正在 Linux 的保持进程中运行。

## Windows

In Windows, set to auto run in the backend on start. Two ways to do that:

&zeroWidthSpace;在 Windows 中，设置开机时在后台自动运行。有两种方法可以做到这一点：

1. Create a bat file and put it into “Run”
   创建一个 bat 文件并将其放入“运行”
2. Create a service 创建服务