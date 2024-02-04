+++
title = "闪存消息"
date = 2024-02-04T21:11:09+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/frontend-layer/flash-messages/](https://gobuffalo.io/documentation/frontend-layer/flash-messages/)

# Flash Messages 闪存消息 

## What are Flash Messages? 什么是闪存消息？ 

Flash messages are a means of communicating messages to the end user from inside of an application. These messages might be errors, warnings, or success types of messages.

​	闪存消息是一种从应用程序内部向最终用户传达消息的方式。这些消息可能是错误、警告或成功类型的消息。

Some examples of flash messages are:

​	闪存消息的一些示例包括：

- “You have been successfully logged out.”
  “您已成功注销。”
- “Your widget could not be updated.”
  “无法更新您的微件。”
- “There was a problem accessing your account.”
  “访问您的帐户时出现问题。”

Being able to set these messages in a Buffalo handler and then pass them down to views is incredibly helpful.

​	能够在 Buffalo 处理程序中设置这些消息，然后将它们传递给视图非常有用。

## Setting Flash Messages 设置闪存消息 

Creating flash messages can easily be done by using the `c.Flash()` function provided on the [`buffalo.Context`](https://gobuffalo.io/documentation/request_handling/context).

​	使用 `buffalo.Context` 上提供的 `c.Flash()` 函数可以轻松创建闪存消息。

```go
func WidgetsCreate(c buffalo.Context) error {
  // do some work
  c.Flash().Add("success", "Widget was successfully created!")
  // do more work and return
}
```

The names of the “keys”, in this example, “success”, are left up to your application to use as is appropriate. There are no “special” or “pre-defined” keys.

​	在这个示例中，“键”的名称（例如“success”）由您的应用程序酌情使用。没有“特殊”或“预定义”键。

## Accessing Flash Messages in Templates 在模板中访问闪存消息 

This document only applies when using https://github.com/gobuffalo/buffalo/tree/main/render.

​	仅当使用 https://github.com/go/tree/main/render 时，此文档才适用。

Please see [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) for more details on the underlying templating package.

​	有关基础模板包的详细信息，请参阅 github.com/go/plush。

### Looping Over all Flash Messages 循环访问所有闪存消息 

```html
<div class="row">
  <div class="col-md-12">
    <%= for (k, messages) in flash { %>
      <%= for (msg) in messages { %>
        <div class="alert alert-<%= k %>" role="alert">
          <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">×</span></button>
          <%= msg %>
        </div>
      <% } %>
    <% } %>
  </div>
</div>
```

### Looping Over a Specific Flash Message Key 循环访问特定闪存消息键 

```html
<div class="row">
  <div class="col-md-12">
    <%= for (message) in flash["success"] { %>
      <div class="alert alert-success" role="alert">
        <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <%= message %>
      </div>
    <% } %>
  </div>
</div>
```