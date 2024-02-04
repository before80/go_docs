+++
title = "flashMessages"
date = 2024-02-04T09:58:15+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/flash/]({{< ref "/beego/mvcIntroduction/controllers/flashMessages" >}})

# Flash messages 闪存消息



## Flash Messages 闪存消息

Flash messages are not related to Adobe/Macromedia Flash. They are temporary messages between two logic blocks. All flash messages will be cleared after the very next logic block. They are normally used to send notes and error messages. Their use is suited for the [Post/Redirect/Get](http://en.wikipedia.org/wiki/Post/Redirect/Get) model. For example:

​	闪存消息与 Adobe/Macromedia Flash 无关。它们是两个逻辑块之间的临时消息。所有闪存消息将在下一个逻辑块之后被清除。它们通常用于发送注释和错误消息。它们的使用适用于 Post/Redirect/Get 模型。例如：

```go
// Display settings message
func (c *MainController) Get() {
    flash := web.ReadFromRequest(&c.Controller)
    if n, ok := flash.Data["notice"]; ok {
        // Display settings successful
        c.TplName = "set_success.html"
    } else if n, ok = flash.Data["error"]; ok {
        // Display error messages
        c.TplName = "set_error.html"
    } else {
        // Display default settings page
        this.Data["list"] = GetInfo()
        c.TplName = "setting_list.html"
    }
}

// Process settings messages
func (c *MainController) Post() {
    flash := web.NewFlash()
    setting := Settings{}
    valid := Validation{}
    c.ParseForm(&setting)
    if b, err := valid.Valid(setting); err != nil {
        flash.Error("Settings invalid!")
        flash.Store(&c.Controller)
        c.Redirect("/setting", 302)
        return
    } else if b != nil {
        flash.Error("validation err!")
        flash.Store(&c.Controller)
        c.Redirect("/setting", 302)
        return
    }
    saveSetting(setting)
    flash.Notice("Settings saved!")
    flash.Store(&c.Controller)
    c.Redirect("/setting", 302)
}
```

The logic of the code above is as follows:

​	以上代码的逻辑如下：

1. Execute GET method. There’s no flash data, so display settings page.
   执行 GET 方法。没有闪存数据，因此显示设置页面。
2. After submission, execute POST and initialize a flash. If checking failed, set error flash message. If checking passed, save settings and set flash message to successful.
   提交后，执行 POST 并初始化闪存。如果检查失败，则设置错误闪存消息。如果检查通过，则保存设置并将闪存消息设置为成功。
3. Redirect to GET request.
   重定向到 GET 请求。
4. GET request receives flash message and executes the related logic. Show error page or success page based on the type of message.
   GET 请求接收闪存消息并执行相关逻辑。根据消息类型显示错误页面或成功页面。

`ReadFromRequest` assigns messages to flash, so you can use it in your template:

​	 `ReadFromRequest` 将消息分配给闪存，以便您可以在模板中使用它：

```
{{.flash.error}}
{{.flash.warning}}
{{.flash.success}}
{{.flash.notice}}
```

There are 4 different levels of flash messages:

​	闪存消息有 4 个不同的级别：

- Notice: Notice message
  通知：通知消息
- Success: Success message
  成功：成功消息
- Warning: Warning message
  警告：警告消息
- Error: Error message
  错误：错误消息
