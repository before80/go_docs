+++
title = "聊天室"
date = 2024-02-04T09:34:42+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/examples/chat/]({{< ref "/beego/demos/chatRoom" >}})

# Chat Room 聊天室

This demo shows two ways of implementing a Web Instant Messaging application:

​	此演示展示了实现 Web 即时消息应用程序的两种方式：

Using long polling. Using WebSocket.

​	使用长轮询。使用 WebSocket。

Both of them save data in memory by default so everything will be lost every time the application restarts, but you can change this setting in `conf/app.conf` to enable a database adapter for data persistence.

​	默认情况下，它们都将数据保存在内存中，因此每次应用程序重新启动时，所有内容都会丢失，但您可以在 `conf/app.conf` 中更改此设置以启用用于数据持久性的数据库适配器。

Here is the project structure:

​	以下是项目结构：

```bash
WebIM/
    WebIM.go            # File of main package
    conf
        app.conf        # Configuration file
    controllers
        app.go          # The welcome screen that allows the user to pick a technology and username
        chatroom.go     # Functions for data management
        longpolling.go  # Controller and methods for long polling chat demo
        websocket.go    # Controller and methods for WebSocket chat demo
    models
        archive.go      # Functions of chat data operations for both demos.
    views
        ...             # Template files
    static
        ...             # JavaScript and CSS files
```

[Browse the code on GitHub
在 GitHub 上浏览代码](https://github.com/beego/samples/tree/master/WebIM)
