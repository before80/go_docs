+++
title = "responseFormats"
date = 2024-02-04T09:58:40+08:00
weight = 10
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/jsonxml/](https://beego.wiki/docs/mvc/controller/jsonxml/)

# Response formats 响应格式



## JSON, XML, JSONP and YAML JSON、XML、JSONP 和 YAML

Beego is also designed for the creation of API applications. When we build an API application, we often need to respond with JSON or XML. Beego provides a simple approach:

&zeroWidthSpace;Beego 还设计用于创建 API 应用程序。当我们构建 API 应用程序时，我们经常需要使用 JSON 或 XML 来响应。Beego 提供了一种简单的方法：

- Respond with JSON data:

  &zeroWidthSpace;使用 JSON 数据响应：

  ```go
  type mystruct struct {
    FieldOne string `json:"field_one"`
  }
  
  func (this *AddController) Get() {
  	mystruct := { ... }
  	this.Data["json"] = &mystruct
  	this.ServeJSON()
  }
  ```

  ServeJson will set `content-type` to `application/json` and JSONify the data.

  &zeroWidthSpace;ServeJson 将 `content-type` 设置为 `application/json` 并对数据进行 JSON 化。

- Respond with XML data:

  &zeroWidthSpace;使用 XML 数据响应：

  ```go
  func (this *AddController) Get() {
  	mystruct := { ... }
  	this.Data["xml"]=&mystruct
  	this.ServeXML()
  }
  ```

  ServeXml will set `content-type` to `application/xml` and convert the data into XML.

  &zeroWidthSpace;ServeXml 将 `content-type` 设置为 `application/xml` 并将数据转换为 XML。

- Respond with jsonp

  &zeroWidthSpace;使用 jsonp 响应

  ```go
  func (this *AddController) Get() {
  	mystruct := { ... }
  	this.Data["jsonp"] = &mystruct
  	this.ServeJSONP()
  }
  ```

  ServeJsonp will set `content-type` to `application/javascript` , JSONify the data and respond to jsonp based on the request parameter `callback`.

  &zeroWidthSpace;ServeJsonp 将 `content-type` 设置为 `application/javascript` ，对数据进行 JSON 化并根据请求参数 `callback` 响应 jsonp。

- Renspond based on Accept Header in request

  &zeroWidthSpace;根据请求中的 Accept 头部进行响应

  ```go
  func (this *AddController) Get() {
  	mystruct := { ... }
  	this.Resp(mystruct)
  }
  ```

  Based on the Accept Header value response will be either JSON, XML or YAML. If Accept header is none of the above by default response will be in JSON format

  &zeroWidthSpace;根据 Accept 头部值，响应将是 JSON、XML 或 YAML。如果 Accept 头部不是上述任何一种，则默认情况下响应将采用 JSON 格式

In version 1.6 names of methods were changed, it is ServeJSON(), ServeXML(), ServeJSONP() from now on.

&zeroWidthSpace;在 1.6 版本中，方法的名称已更改，从现在开始是 ServeJSON()、ServeXML()、ServeJSONP()。