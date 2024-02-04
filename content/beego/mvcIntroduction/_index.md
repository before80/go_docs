+++
title = "MVC 简介"
date = 2024-02-04T09:11:29+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/]({{< ref "/beego/mvcIntroduction" >}})

# MVC Introduction - MVC 简介



## Introduction to Beego’s MVC - Beego 的 MVC 简介

Beego uses a typical Model-View-Controller (MVC) framework. This diagram illustrates how request handling logic is processed:

​	Beego 使用典型的模型-视图-控制器 (MVC) 框架。此图表说明了如何处理请求处理逻辑：

![img](./_index_img/detail.png)

The whole logic handling process is explained below:

​	整个逻辑处理过程如下所述：

1. Data is recieved from the listening port. The listening port is set to 8080 by default.
   从侦听端口接收数据。侦听端口默认设置为 8080。
2. Beego begins processing the requested data after the request reachs port 8080
   请求到达端口 8080 后，Beego 开始处理请求的数据
3. The Context object is initialized. WebSocket requests will be set as Input after the request method has been verified as a standard method (get, post, put, delete, patch, options, head) in order to protect from hostile attack.
   初始化 Context 对象。在将请求方法验证为标准方法（get、post、put、delete、patch、options、head）后，WebSocket 请求将被设置为 Input，以防止恶意攻击。
4. If the `BeforeRouter` filter has been set by the user it is executed. If `responseWriter` has output data while executing this filter the request will be finished and the supervise checking step (see item 21) will be performed next.
   如果用户已设置 `BeforeRouter` 过滤器，则执行该过滤器。如果 `responseWriter` 在执行此过滤器时有输出数据，则请求将完成，然后执行监督检查步骤（参见项目 21）。
5. Start handling of static files. If the requested url matches the prefix set by `StaticDir`, the `ServeFile` in the `http` package will be used to handle the static file requests.
   开始处理静态文件。如果请求的 url 与 `StaticDir` 设置的前缀匹配，则 `http` 包中的 `ServeFile` 将用于处理静态文件请求。
6. If the request is not for a static file, and the session module is enabled, the session module will initialize. An error will occur if session is being used in the `BeforeRouter` filter (see item 4). Use the `AfterStatic` filter (see item 7) instead to avoid this error.
   如果请求不是针对静态文件，并且启用了会话模块，则会话模块将初始化。如果在 `BeforeRouter` 过滤器中使用会话，则会发生错误（请参阅第 4 项）。改用 `AfterStatic` 过滤器（请参阅第 7 项）以避免此错误。
7. The `AfterStatic` filter is executed. If `responseWriter` already has output data while executing this filter the request will be finished and the supervise checking step (see item 21) will be performed next.
   执行 `AfterStatic` 过滤器。如果在执行此过滤器时 `responseWriter` 已有输出数据，则请求将完成，接下来将执行监督检查步骤（请参阅第 21 项）。
8. After all filters have been executed Beego will start to match any requested urls with fixed routing rules. These connections will only be made if the whole string matches. For example: `/hello` does not match the url of `/hello/world`. If any matching pairs are found the appropriate logic will execute.
   在执行完所有过滤器后，Beego 将开始将任何请求的 URL 与固定的路由规则进行匹配。仅当整个字符串匹配时才会建立这些连接。例如： `/hello` 与 `/hello/world` 的 URL 不匹配。如果找到任何匹配对，则将执行相应的逻辑。
9. Regex matching is executed based on the order that the user added. This means the order of the regex routing rules will affect Regex matching. If any matches are found the appropriate logic will execute.
   根据用户添加的顺序执行正则表达式匹配。这意味着正则表达式路由规则的顺序将影响正则表达式匹配。如果找到任何匹配项，则将执行相应的逻辑。
10. If the user registered `AutoRouter`, `controller/method` will be used to match the Controller and method. If any matches are found the appropriate logic will execute. Otherwise the supervise checking step (see item 21) will be performed next.
    如果用户注册了 `AutoRouter` ，则将使用 `controller/method` 来匹配控制器和方法。如果找到任何匹配项，则将执行相应的逻辑。否则，接下来将执行监督检查步骤（请参阅第 21 项）。
11. If a Controller is found Beego will start this logic. `BeforeExec` will execute first. If `responseWriter` already has output data while executing this filter the request will be finished and the supervise checking step (see item 21) will be performed next.
    如果找到控制器，Beego 将启动此逻辑。 `BeforeExec` 将首先执行。如果 `responseWriter` 在执行此过滤器时已具有输出数据，则请求将完成，接下来将执行监督检查步骤（参见第 21 项）。
12. Controller will start executing the `Init` function to initialize basic information. `bee.Controller` will usually be initialized as part of this item, so modifiying this function is not recommened while inheriting the Controller.
    控制器将开始执行 `Init` 函数以初始化基本信息。 `bee.Controller` 通常会作为此项的一部分进行初始化，因此在继承控制器时不建议修改此函数。
13. If XSRF is enabled it will call `XsrfToken` of `Controller`. If this is a POST request `CheckXsrfCookie` will be called.
    如果启用了 XSRF，它将调用 `Controller` 的 `XsrfToken` 。如果这是 POST 请求，则将调用 `CheckXsrfCookie` 。
14. The `Prepare` function of `Controller` will be executed. This function is normally used by the user to launch initialization. If `responseWriter` already has data while executing this filter it will go directly to the `Finish` function (see item 17).
    将执行 `Controller` 的 `Prepare` 函数。此函数通常由用户用于启动初始化。如果 `responseWriter` 在执行此过滤器时已具有数据，它将直接转到 `Finish` 函数（参见第 17 项）。
15. If there is no output the user registered function will be executed. If there is no function registered by the user the method in `http.Method` (GET/POST and so on) will be called. This will execute logic inluding reading data, assigning data, rendering templates, or outputing JSON or XML.
    如果没有输出，将执行用户注册的函数。如果用户未注册任何函数，则将调用 `http.Method` （GET/POST 等）中的方法。这将执行包括读取数据、分配数据、呈现模板或输出 JSON 或 XML 在内的逻辑。
16. If there is no output by `responseWrite` the `Render` function will be called to output template.
    如果 `responseWrite` 没有输出，则将调用 `Render` 函数以输出模板。
17. Execute the `Finish` function of `Controller`. This function works as an override to allow the user to release resources, such as data initialized in `Init`.
    执行 `Controller` 的 `Finish` 函数。此函数用作覆盖，允许用户释放资源，例如在 `Init` 中初始化的数据。
18. Execute the `AfterExec` filter. If there is output it will jump to the supervise checking step (see item 21).
    执行 `AfterExec` 过滤器。如果有输出，它将跳转到监督检查步骤（参见项目 21）。
19. Execute `Destructor` in `Controller` to release data allocated in `Init`.
    在 `Controller` 中执行 `Destructor` 以释放 `Init` 中分配的数据。
20. If there is no router has been found the 404 page will be shown.
    如果没有找到路由器，将显示 404 页面。
21. Eventually, all logic paths lead to supervise checking. If the supervisor module is enabled (default on port 8088), the request will be sent to supervisor module to log data such as QPS of the request, visiting time, and request url.
    最终，所有逻辑路径都会导致监督检查。如果启用了监督模块（默认端口 8088），则会将请求发送到监督模块以记录数据，例如请求的 QPS、访问时间和请求 URL。

The next sections will detail the first step of Beego’s MVC, routing:

​	下一节将详细介绍 Beego 的 MVC 的第一步，路由：

- [Routing 路由]({{< ref "/beego/mvcIntroduction/controllers/routing" >}})
- [Controller functions 控制器函数]({{< ref "/beego/mvcIntroduction/controllers/controllerFuncs" >}})
- [Cross-site request forgery (XSRF)
  跨站点请求伪造 (XSRF)]({{< ref "/beego/mvcIntroduction/controllers/xfrfFiltering" >}})
- [Session control 会话控制]({{< ref "/beego/mvcIntroduction/controllers/sessionControl" >}})
- [Message flashing 消息闪烁]({{< ref "/beego/mvcIntroduction/controllers/flashMessages" >}})
- [Accessing Request Data 访问请求数据]({{< ref "/beego/mvcIntroduction/controllers/requestParameters" >}})
- [Multiple Response Formats
  多种响应格式]({{< ref "/beego/mvcIntroduction/controllers/responseFormats" >}})
- [Form validation 表单验证]({{< ref "/beego/mvcIntroduction/controllers/formValidation" >}})
- [Rendering templates 渲染模板]({{< ref "/beego/mvcIntroduction/views/templateParsing" >}})
- [Template functions 模板函数]({{< ref "/beego/mvcIntroduction/views/templateFunctions" >}})
- [Error handling 错误处理]({{< ref "/beego/mvcIntroduction/controllers/errorHanding" >}})
- [Handling static files 处理静态文件]({{< ref "/beego/mvcIntroduction/views/staticFiles" >}})
- [Parameter configuration 参数配置]({{< ref "/beego/mvcIntroduction/controllers/configuration" >}})
