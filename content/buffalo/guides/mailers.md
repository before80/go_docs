+++
title = "邮件发送器"
date = 2024-02-04T21:17:02+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/mailers/]({{< ref "/buffalo/guides/mailers" >}})

# Mailers 邮件发送器 

E-mails are part of most (web) systems. It’s a standard way to inform users about platform news, confirmations, marketing stuff, and so on.

​	电子邮件是大多数（网络）系统的一部分。这是通知用户有关平台新闻、确认、营销内容等的一种标准方式。

Buffalo provides, out of the box, a mailer extension with a standard SMTP sender. A generator is included, to allow you to work with emails as fast as possible.

​	Buffalo 开箱即用，提供了一个带有标准 SMTP 发送器的邮件发送器扩展。其中包含一个生成器，以便您能够尽可能快地处理电子邮件。

## Generator 生成器 

When the generator is run for the first time it will bootstrap a new `mailers` package and a new `templates/mail` directory.

​	首次运行生成器时，它将引导一个新的 `mailers` 包和一个新的 `templates/mail` 目录。

```bash
$ buffalo generate mailer welcome_email
```

## Example Usage 示例用法 

```go
// mailers/mail.go
package x

import (
  "log"

  "github.com/gobuffalo/buffalo/render"
  "github.com/gobuffalo/envy"
  "github.com/gobuffalo/packr"
  "github.com/gobuffalo/plush"
  "github.com/gobuffalo/buffalo/mail"
  "github.com/pkg/errors"
  "gitlab.com/wawandco/app/models"
)

var smtp mail.Sender
var r *render.Engine

func init() {

  // Pulling config from the env.
  port := envy.Get("SMTP_PORT", "1025")
  host := envy.Get("SMTP_HOST", "localhost")
  user := envy.Get("SMTP_USER", "")
  password := envy.Get("SMTP_PASSWORD", "")

  var err error
  smtp, err = mail.NewSMTPSender(host, port, user, password)

  if err != nil {
    log.Fatal(err)
  }

  // The rendering engine, this is usually generated inside actions/render.go in your buffalo app.
  r = render.New(render.Options{
    TemplatesBox:   packr.NewBox("../templates"),
  })
}

// SendContactMessage Sends contact message to contact@myapp.com
func SendContactMessage(c *models.Contact) error {

  // Creates a new message
  m := mail.NewMessage()
  m.From = "sender@myapp.com"
  m.Subject = "New Contact"
  m.To = []string{"contact@myapp.com"}

  // Data that will be used inside the templates when rendering.
  data := map[string]interface{}{
    "contact": c,
  }

  // You can add multiple bodies to the message you're creating to have content-types alternatives.
  err := m.AddBodies(data, r.HTML("mail/contact.html"), r.Plain("mail/contact.txt"))

  if err != nil {
    return errors.WithStack(err)
  }

  err = smtp.Send(m)
  if err != nil {
    return errors.WithStack(err)
  }

  return nil
}
```

This `SendContactMessage` could be called by one of your actions, i.e. the action that handles your contact form submission.

​	此 `SendContactMessage` 可以由您的某个操作调用，即处理您的联系表单提交的操作。

```go
// actions/contact.go
...

func ContactFormHandler(c buffalo.Context) error {
  contact := &models.Contact{}
  c.Bind(contact)

  // Calling to send the message
  SendContactMessage(contact)
  return c.Redirect(302, "contact/thanks")
}
...
```

You can add your own custom plush functions by binding them in as data.

​	您可以通过将它们绑定为数据来添加您自己的自定义 plush 函数。

```go
func UUIDToString(u uuid.UUID) string {
  return fmt.Sprintf("%s", u)
}

  m := mail.NewMessage()
  ...
  
  // Data that will be used inside the templates when rendering.
  data := map[string]interface{}{
    "contact": c,
    "UUIDToString": UUIDToStringHelper,
  }
```

## Using Context Variables 使用上下文变量 

Since **0.13.0-rc1**
自 0.13.0-rc1



To use context variables such as [RouteHelpers]({{< ref "/buffalo/requestHandling/routing#using-route-helpers-in-templates" >}}) or those set with `c.Set(...)`, `mail.New` accepts a `buffalo.Context`.

​	要使用诸如 RouteHelpers 或使用 `c.Set(...)` 设置的上下文变量， `mail.New` 接受 `buffalo.Context` 。

```go
func SendMail(c buffalo.Context) error {
  m := mail.New(c)
  ...

  m.AddBody(r.HTML("mail.html"))
  return SMTP.Send(m)
}
<a href="\<%= awesomePath() %>">Click here</a>
```

## Additional Configuration 其他配置 

If you’re using Gmail or need to configure your SMTP connection, you can use the `Dialer` property on the SMTPSender, p.e: (for Gmail)

​	如果您使用 Gmail 或需要配置 SMTP 连接，则可以在 SMTPSender 上使用 `Dialer` 属性，例如：（对于 Gmail）

```go
// mailers/mail.go
...
var smtp mail.Sender

func init() {
  port := envy.Get("SMTP_PORT", "465")
  // or 587 with TLS

  host := envy.Get("SMTP_HOST", "smtp.gmail.com")
  user := envy.Get("SMTP_USER", "your@email.com")
  password := envy.Get("SMTP_PASSWORD", "yourp4ssw0rd")

  // Assigning to smtp later to preserve type
  var err error
  sender, err := mail.NewSMTPSender(host, port, user, password)
  sender.Dialer.SSL = true

  //or if TLS
  sender.Dialer.TLSConfig = &tls.Config{...}

  smtp = sender
}
...
```

## Sender Implementations 发送器实现 

Some alternate [`Sender`](https://godoc.org/github.com/gobuffalo/buffalo/mail#Sender) implementations are provided by the Buffalo community:

​	Buffalo 社区提供了一些备用 `Sender` 实现：

| Package 软件包 postmark-sender                               | Description 说明                                             | Author 作者                                    |
| :----------------------------------------------------------- | :----------------------------------------------------------- | :--------------------------------------------- |
| [postmark-sender 与 Postmark 配合使用的发送器 mocksmtp](https://github.com/paganotoni/postmark-sender) | A sender to work with [Postmark](https://postmarkapp.com/) 可用于测试的模拟实现 | [@paganotoni](https://github.com/paganotoni)   |
| [mocksmtp](https://github.com/stanislas-m/mocksmtp)          | A mock implementation that can be used for testing           | [@stanislas-m](https://github.com/stanislas-m) |
| [sendgrid-sender](https://github.com/paganotoni/sendgrid-sender) | A sender to work with [Sendgrid](https://sendgrid.com/) 一个与 Sendgrid 协同工作的发件人 | [@paganotoni](https://github.com/paganotoni)   |
| [mailopen](https://github.com/paganotoni/mailopen)           | A sender that opens emails in browser 一个在浏览器中打开电子邮件的发件人 | [@paganotoni](https://github.com/paganotoni)   |
