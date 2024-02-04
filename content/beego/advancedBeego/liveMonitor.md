+++
title = "实时监控"
date = 2024-02-04T09:33:42+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/advantage/monitor/]({{< ref "/beego/advancedBeego/liveMonitor" >}})

# Live Monitor 实时监控



We discussed the toolbox module before. It will listen on `127.0.0.1:8088` by default when the application is running. It can’t be accessed from the internet but you can browse to it by other means such as by nginx proxy.

​	我们之前讨论过工具箱模块。在应用程序运行时，它默认监听 `127.0.0.1:8088` 。无法从互联网访问它，但您可以通过其他方式（例如通过 nginx 代理）浏览它。

> For security reason it is recommended that you block port 8088 in firewall.
>
> ​	出于安全原因，建议您在防火墙中阻止端口 8088。

Monitor is disabled by default. You can enable it by adding the following line in `conf/app.conf` file:

​	默认情况下，监视器处于禁用状态。您可以通过在 `conf/app.conf` 文件中添加以下行来启用它：

```
EnableAdmin = true
```

Also you can change the port it listens on:

​	您还可以更改它监听的端口：

```
AdminAddr = "localhost"
AdminPort = 8088
```

Open browser and visit `http://localhost:8088/` you will see `Welcome to Admin Dashboard`.

​	打开浏览器并访问 `http://localhost:8088/` ，您将看到 `Welcome to Admin Dashboard` 。

## Requests statistics 请求统计信息

Browse to `http://localhost:8088/qps` and you will see the following:

​	浏览到 `http://localhost:8088/qps` ，您将看到以下内容：

![img](https://beego.wiki/docs/images/monitoring.png)

## Performance profiling 性能分析

You can also see the information for `goroutine`, `heap`, `threadcreate`, `block`, `cpuprof`, `memoryprof`, `gc summary` and do profiling.

​	您还可以查看 `goroutine` 、 `heap` 、 `threadcreate` 、 `block` 、 `cpuprof` 、 `memoryprof` 、 `gc summary` 的信息并进行分析。

## Healthcheck

You need to manually register the healthcheck logic to see the status of the healthcheck from `http://localhost:8088/healthcheck`

​	您需要手动注册健康检查逻辑才能从 `http://localhost:8088/healthcheck` 查看健康检查的状态

## Tasks 任务

You can add task in your application and check the task status or trigger it manually.

​	您可以在应用程序中添加任务并检查任务状态或手动触发任务。

- Check task status: `http://localhost:8088/task`
  检查任务状态： `http://localhost:8088/task`
- Run task manually: `http://localhost:8088/runtask?taskname=task_name`
  手动运行任务： `http://localhost:8088/runtask?taskname=task_name`

## Config Status 配置状态

After the development of the application, we may also want to know the config when the application is running. Beego’s Monitor also provided this feature.

​	在应用程序开发完成后，我们可能还想了解应用程序运行时的配置。Beego 的 Monitor 也提供了此功能。

- Show all configurations: `http://localhost:8088/listconf?command=conf`
  显示所有配置： `http://localhost:8088/listconf?command=conf`
- Show all routers: `http://localhost:8088/listconf?command=router`
  显示所有路由： `http://localhost:8088/listconf?command=router`
- Show all filters: `http://localhost:8088/listconf?command=filter`
  显示所有过滤器： `http://localhost:8088/listconf?command=filter`
