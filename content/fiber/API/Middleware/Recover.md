+++
title = "Recover"
date = 2024-02-05T09:14:15+08:00
weight = 220
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/recover]({{< ref "/fiber/API/Middleware/Recover" >}})

# Recover

Recover middleware for [Fiber](https://github.com/gofiber/fiber) that recovers from panics anywhere in the stack chain and handles the control to the centralized [ErrorHandler]({{< ref "/fiber/Guide/ErrorHandling" >}}).

​	Fiber 的恢复中间件，可从堆栈链中的任何位置恢复恐慌，并将控制权交给集中的 ErrorHandler。

## Signatures 签名

```go
func New(config ...Config) fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/recover"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(recover.New())

// This panic will be caught by the middleware
app.Get("/", func(c *fiber.Ctx) error {
    panic("I'm an error")
})
```



## Config 配置

| Property 属性     | Type 输入                       | Description 说明                                             | Default 默认             |
| ----------------- | ------------------------------- | ------------------------------------------------------------ | ------------------------ |
| Next 下一步       | `func(*fiber.Ctx) bool`         | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                    |
| EnableStackTrace  | `bool`                          | EnableStackTrace enables handling stack trace. EnableStackTrace 启用堆栈跟踪处理。 | `false`                  |
| StackTraceHandler | `func(*fiber.Ctx, interface{})` | StackTraceHandler defines a function to handle stack trace. StackTraceHandler 定义一个处理堆栈跟踪的函数。 | defaultStackTraceHandler |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next:              nil,
    EnableStackTrace:  false,
    StackTraceHandler: defaultStackTraceHandler,
}
```
