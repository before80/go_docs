+++
title = "高级 Beego"
date = 2024-02-04T09:33:12+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/advantage/]({{< ref "/beego/advancedBeego" >}})

# Advanced Beego 高级 Beego



We have demonstrated the basic usage of Beego. Now we will talk about more advanced topics.

​	我们已经演示了 Beego 的基本用法。现在我们将讨论更高级的话题。

- [In process monitor
  正在进行的监控]({{< ref "/beego/advancedBeego/liveMonitor" >}})

  Beego will serve as two ports by default. One is 8080 for application to serve users. Another is 8088, to monitor the process status, execute tasks and so on.

  ​	Beego 默认情况下将提供两个端口。一个是 8080，用于应用程序为用户提供服务。另一个是 8088，用于监控进程状态、执行任务等。

- [Filters
  过滤器]({{< ref "/beego/mvcIntroduction/controllers/filters" >}})

  Filters is a very convenient feature for you to extend your logic. You can easily implement user authentication, log visiting, compatibility switching and so on.

  ​	过滤器是一个非常方便的功能，可用于扩展您的逻辑。您可以轻松实现用户身份验证、记录访问、兼容性切换等。

- [Reload
  重新加载]({{< ref "/beego/modules/grace" >}})

  Reload is always mentioned in web development that allows deploying application without interrupt user requests.

  ​	重新加载始终在 Web 开发中提及，它允许在不中断用户请求的情况下部署应用程序。

> This feature is not well done yet. It only tested on Mac and Linux. It haven’t been tested on production environment yet. It’s still under testing, so take your own risk to use it. It’s recommended to use upstream of nginx.
>
> ​	此功能尚未完善。它仅在 Mac 和 Linux 上进行了测试。尚未在生产环境中进行测试。它仍在测试中，因此请自行承担使用风险。建议在 nginx 上游使用。
