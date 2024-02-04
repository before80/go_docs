+++
title = "路由设置"
date = 2024-02-04T09:10:13+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/quickstart/router/](https://beego.wiki/docs/quickstart/router/)

# Routing settings 路由设置



## Project routing settings 项目路由设置

The previous section covered creating and running a Beego project. This section will investigate the operation of the main file (main.go):

&zeroWidthSpace;上一节介绍了如何创建和运行 Beego 项目。本节将研究主文件 (main.go) 的操作：

```
package main

import (
        _ "quickstart/routers"
        "github.com/beego/beego/v2/server/web"
)

func main() {
        web.Run()
}
```

This code imports the package `quickstart/routers`. This file contains the following (routers/router.go):

&zeroWidthSpace;此代码导入包 `quickstart/routers` 。此文件包含以下内容 (routers/router.go)：

```
package routers

import (
        "quickstart/controllers"
        "github.com/beego/beego/v2/server/web"
)

func init() {
        web.Router("/", &controllers.MainController{})
}
```

There are two relevant lines here; `web.Router` and `web.Run`.

&zeroWidthSpace;这里有两行相关代码； `web.Router` 和 `web.Run` 。

1. `web.Router` is used to register a router address. This command accepts two arguments. The first argument specifes the request uri, which is `/` here to indicate that no uri is requested. The second argument is used to indicate the Controller that will handle requests for this uri.
   `web.Router` 用于注册路由器地址。此命令接受两个参数。第一个参数指定请求 uri，此处为 `/` ，表示未请求 uri。第二个参数用于指示将处理此 uri 的请求的控制器。

Alternately, a router can be registered in this format:

&zeroWidthSpace;或者，可以按此格式注册路由器：

```
	web.Router("/user", &controllers.UserController{})
```

The user can visit `/user` to invoke the logic in UserController. For further information on router usage please see [beego router settings](https://beego.wiki/docs/mvc/controller/router).

&zeroWidthSpace;用户可以访问 `/user` 以调用 UserController 中的逻辑。有关路由器使用情况的更多信息，请参阅 beego 路由器设置。

1. `web.Run` will actively listen on the specified port when executed. The following tasks are performed behind the scenes upon execution:
   执行时， `web.Run` 将主动侦听指定端口。执行时，后台将执行以下任务：

- Parse the [configuration file](https://beego.wiki/docs/mvc/controller/config) Beego will parse the configuration file `app.conf` in conf folder to change the port, enable session management and set the application’s name.

  &zeroWidthSpace;解析配置文件 Beego 将解析 conf 文件夹中的配置文件 `app.conf` 以更改端口、启用会话管理并设置应用程序的名称。

- Initialize the [user session](https://beego.wiki/docs/mvc/controller/session) Beego will initialize the user session, based on the setting in the configuration file.

  &zeroWidthSpace;初始化用户会话 Beego 将根据配置文件中的设置初始化用户会话。

- Compile the [views](https://beego.wiki/docs/quickstart/view) Beego will compile the views in the views folder. This is done on startup to avoid compiling multiple times and improve efficiency.

  &zeroWidthSpace;编译视图 Beego 将编译 views 文件夹中的视图。这在启动时完成，以避免多次编译并提高效率。

- Starting the [supervisor module](https://beego.wiki/docs/advantage/monitor) By visiting port `8088` the user can access information about QPS, cpu, memory, GC, goroutine and thread information.

  &zeroWidthSpace;启动监督程序模块 通过访问端口 `8088` ，用户可以访问有关 QPS、cpu、内存、GC、goroutine 和线程的信息。

- Listen on the service port Beego will listen http requests on port `8080`. It takes advantage of goroutines by calling `ListenAndServe`.

  &zeroWidthSpace;监听服务端口 Beego 将在端口 `8080` 上监听 http 请求。它通过调用 `ListenAndServe` 来利用 goroutine。

- When the application is running our server will serve incoming requests from port `8080` and supervising from port `8088`.

  &zeroWidthSpace;当应用程序正在运行时，我们的服务器将处理来自端口 `8080` 的传入请求，并从端口 `8088` 进行监督。

The next section will cover the operation of the controller [next section](https://beego.wiki/docs/quickstart/controller).

&zeroWidthSpace;下一节将介绍控制器的操作下一节。