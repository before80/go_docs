+++
title = "示例"
date = 2024-01-31T19:10:31+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-chi/chi/blob/master/_examples/README.md](https://github.com/go-chi/chi/blob/master/_examples/README.md)

# chi examples chi 示例

- [custom-handler](https://github.com/go-chi/chi/blob/master/_examples/custom-handler/main.go) - Use a custom handler function signature
  custom-handler - 使用自定义处理程序函数签名
- [custom-method](https://github.com/go-chi/chi/blob/master/_examples/custom-method/main.go) - Add a custom HTTP method
  custom-method - 添加自定义 HTTP 方法
- [fileserver](https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go) - Easily serve static files
  fileserver - 轻松提供静态文件
- [graceful](https://github.com/go-chi/chi/blob/master/_examples/graceful/main.go) - Graceful context signaling and server shutdown
  graceful - 优雅的上下文信令和服务器关闭
- [hello-world](https://github.com/go-chi/chi/blob/master/_examples/hello-world/main.go) - Hello World!
- [limits](https://github.com/go-chi/chi/blob/master/_examples/limits/main.go) - Timeouts and Throttling
  限制 - 超时和节流
- [logging](https://github.com/go-chi/chi/blob/master/_examples/logging/main.go) - Easy structured logging for any backend
  日志记录 - 适用于任何后端的简单结构化日志记录
- [rest](https://github.com/go-chi/chi/blob/master/_examples/rest/main.go) - REST APIs made easy, productive and maintainable
  rest - REST API 简便、高效且易于维护
- [router-walk](https://github.com/go-chi/chi/blob/master/_examples/router-walk/main.go) - Print to stdout a router's routes
  router-walk - 将路由器的路由打印到 stdout
- [todos-resource](https://github.com/go-chi/chi/blob/master/_examples/todos-resource/main.go) - Struct routers/handlers, an example of another code layout style
  todos-resource - 结构路由器/处理程序，另一个代码布局样式的示例
- [versions](https://github.com/go-chi/chi/blob/master/_examples/versions/main.go) - Demo of `chi/render` subpkg
  版本 - `chi/render` 子包的演示

## Usage 用法

1. `go get -v -d -u ./...` - fetch example deps 获取示例依赖项
2. `cd <example>/` ie. `cd rest/`
3. `go run *.go` - note, example services run on port 3333  注意，示例服务在端口 3333 上运行
4. Open another terminal and use curl to send some requests to your example service, `curl -v http://localhost:3333/`  打开另一个终端并使用 curl 向示例服务发送一些请求， `curl -v http://localhost:3333/`
5. Read /main.go source to learn how service works and read comments for usage 阅读 /main.go 源代码以了解服务的工作原理并阅读用法注释