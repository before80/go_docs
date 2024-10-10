+++
title = "context 模块"
date = 2024-02-04T09:31:39+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/module/context/]({{< ref "/beego/modules/context" >}})

# Context Module 上下文模块

The Context module is an encapsulation for http request and response. The Context module provides an Input object for user input which is the request and an Output object for output which is the response.

​	上下文模块是对 http 请求和响应的封装。上下文模块提供了一个用于用户输入的输入对象（即请求）和一个用于输出的输出对象（即响应）。

## Context Object 上下文对象

Here are the functions encapsulated for input and output in the context object.

​	以下是上下文对象中封装的用于输入和输出的函数。

- Redirect 重定向
- Abort 中止
- WriteString
- GetCookie
- SetCookie

Context object is the parameter of a Filter function so that you can use a filter to manipulate it or finish the process in advance.

​	Context 对象是 Filter 函数的参数，因此您可以使用过滤器对其进行操作或提前完成处理。

## Input Object 输入对象

The Input object is the encapsulation of request. Here are the implemented methods:

​	输入对象是对请求的封装。以下是实现的方法：

- Protocol 
  ​	协议

  Get request protocol. E.g.: `HTTP/1.0`

  ​	获取请求协议。例如： `HTTP/1.0`

- Uri

  The RequestURI of request. E.g.: `/hi`

  ​	请求的 RequestURI。例如： `/hi`

- Url

  The URL of request. E.g.: `http://beego.wiki/about?username=astaxie`

  ​	请求的 URL。例如： `http://beego.wiki/about?username=astaxie`

- Site 
  ​	站点

  The combination of scheme and domain. E.g.: `http://beego.wiki`

  ​	方案和域的组合。例如： `http://beego.wiki`

- Scheme 
  ​	方案

  The request scheme. E.g.: `http`, `https`

  ​	请求方案。例如： `http` 、 `https`

- Domain 
  ​	域

  The request domain. E.g.: `beego.wiki`

  ​	请求域。例如： `beego.wiki`

- Host 
  ​	主机

  The request domain. Same as Domain.

  ​	请求域。与域相同。

- Method 
  ​	方法

  The request method. It’s a standard http request method. E.g.: `GET`, `POST`,

  ​	请求方法。它是标准的 http 请求方法。例如： `GET` 、 `POST` 、

- Is

  Test if it’s a http method. E.g.: `Is("GET")` will return true or false

  ​	测试它是否是一个 http 方法。例如： `Is("GET")` 将返回真或假

- IsAjax

  Test if it’s a ajax request. Return true or false.

  ​	测试是否为 ajax 请求。返回 true 或 false。

- IsSecure

  Test if the request is an https request. Return true or false.

  ​	测试请求是否为 https 请求。返回 true 或 false。

- IsWebsocket

  Test if the request is a Websocket request. Return true or false.

  ​	测试请求是否为 Websocket 请求。返回 true 或 false。

- IsUpload

  Test if there a is file uploaded in the request. Return true or false.

  ​	测试请求中是否上传了文件。返回 true 或 false。

- IP

  Return the IP of the requesting user. If the user is using a proxy, it will get the real IP recursively.

  ​	返回请求用户的 IP。如果用户使用代理，它将递归获取真实 IP。

- Proxy

  Return all IP addresses of the proxy request.

  ​	返回代理请求的所有 IP 地址。

- Refer

  Return the refer of the request.

  ​	返回请求的来源。

- SubDomains

  Return the sub domains of the request. For example, request domain is `blog.beego.wiki`, then this function returns `blog`.

  ​	返回请求的子域。例如，请求域为 `blog.beego.wiki` ，则此函数返回 `blog` 。

- Port

  Return the port of request. E.g.: 8080

  ​	返回请求的端口。例如：8080

- UserAgent

  Return `UserAgent` of request. E.g.: `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36`

  ​	返回请求的 `UserAgent` 。例如： `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36`

- Param

  Can be set in router config. Used to get those params. E.g.: `Param(":id")` return 12

  ​	可以在路由器配置中设置。用于获取这些参数。例如： `Param(":id")` 返回 12

- Query 
  ​	查询

  Return all params in GET and POST requests. This is similar as `$_REQUEST` in PHP

  ​	返回 GET 和 POST 请求中的所有参数。这与 PHP 中的 `$_REQUEST` 类似

- Header 
  ​	标题

  Return request header. E.g.: `Header("Accept-Language")` will return the value in request header, E.g.: `zh-CN,zh;q=0.8,en;q=0.6`

  ​	返回请求标题。例如： `Header("Accept-Language")` 将返回请求标题中的值，例如： `zh-CN,zh;q=0.8,en;q=0.6`

- Cookie

  Return request Cookie. E.g.: `Cookie("username")` will return the value of username in cookies

  ​	返回请求 Cookie。例如： `Cookie("username")` 将返回 Cookie 中的 username 值

- Session 
  ​	会话

  Return initialized session. It is the Session object in the session module of Beego. Return the related data stored on the server.

  ​	返回已初始化的会话。它是 Beego 会话模块中的 Session 对象。返回存储在服务器上的相关数据。

- Body 
  ​	正文

  Return request body. E.g.: in API application request sends JSON data and it can’t be retrieved by Query. You need to use Body to get the JSON data.

  ​	返回请求正文。例如：在 API 应用程序请求中发送 JSON 数据，且无法通过查询检索该数据。您需要使用正文来获取 JSON 数据。

- GetData

  Get value of `Data` in `Input`

  ​	获取 `Input` 中 `Data` 的值

- SetData

  Set value of `Data` in `Input`. `GetData` and `SetData` is used to pass data from Filter to Controller.

  ​	在 `Input` 中设置 `Data` 的值。 `GetData` 和 `SetData` 用于将数据从过滤器传递到控制器。

## Output Object 输出对象

Output object is the encapsulation of response. Here are the implemented methods:

​	输出对象是对响应的封装。以下是实现的方法：

- Header

  Set response header. E.g.: `Header("Server","beego")`

  ​	设置响应头。例如： `Header("Server","beego")`

- Body

  Set response body. E.g.: `Body([]byte("astaxie"))`

  ​	设置响应主体。例如： `Body([]byte("astaxie"))`

- Cookie

  Set response cookie. E.g.: `Cookie("sessionID","beegoSessionID")`

  ​	设置响应 Cookie。例如： `Cookie("sessionID","beegoSessionID")`

- Json

  Parse Data into JSON and call `Body` to return it.

  ​	将数据解析为 JSON 并调用 `Body` 返回它。

- Jsonp

  Parse Data into JSONP and call `Body` to return it.

  ​	将数据解析为 JSONP 并调用 `Body` 返回它。

- Xml

  Parse Data into XML and call `Body` to return it.

  ​	将数据解析为 XML 并调用 `Body` 返回它。

- Download 
  ​	下载

  Pass in file path and output file.

  ​	传入文件路径和输出文件。

- ContentType

  Set response ContentType

  ​	设置响应 ContentType

- SetStatus

  Set response status 
  ​	设置响应状态

- Session

  Set the value which will be stored on the server. E.g.: `Session("username","astaxie")`. Then it can be read later.

  ​	设置将在服务器上存储的值。例如： `Session("username","astaxie")` 。然后稍后可以读取它。

- IsCachable

  Test if it’s a cacheable status based on status.

  ​	根据状态测试是否可缓存的状态。

- IsEmpty

  Test if output is empty based on status.

  ​	根据状态测试输出是否为空。

- IsOk

  Test if response is 200 based on status.

  ​	根据状态测试响应是否为 200。 IsSuccessful

- IsSuccessful 
  ​	根据状态测试响应是否成功。 IsRedirect

  Test if response is successful based on status.

  ​	根据状态测试响应是否被重定向。 IsForbidden

- IsRedirect 
  ​	根据状态测试响应是否被禁止。 IsNotFound

  Test if response is redirected based on status.

  ​	根据状态测试响应是否被禁止。

- IsForbidden

  Test if response is forbidden based on status.

- IsNotFound

  Test if response is forbidden based on status.

- IsClientError

  Test if response is client error based on status.

  ​	根据状态测试响应是否是客户端错误。

- IsServerError

  Test if response is server error based on status.

  ​	根据状态测试响应是否是服务器错误。
