+++
title = "Rewrite"
date = 2024-02-05T09:14:15+08:00
weight = 250
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/rewrite]({{< ref "/fiber/API/Middleware/Rewrite" >}})

# Rewrite 重写

Rewrite middleware rewrites the URL path based on provided rules. It can be helpful for backward compatibility or just creating cleaner and more descriptive links.

​	重写中间件根据提供的规则重写 URL 路径。它有助于向后兼容或只是创建更简洁、更具描述性的链接。

## Signatures 签名

```go
func New(config ...Config) fiber.Handler
```



## Config 配置

| Property 属性 | Type 输入               | Description 说明                                             | Default 默认        |
| ------------- | ----------------------- | ------------------------------------------------------------ | ------------------- |
| Next 下一步   | `func(*fiber.Ctx) bool` | Next defines a function to skip middleware. 接下来定义一个跳过中间件的函数。 | `nil`               |
| Rules 规则    | `map[string]string`     | Rules defines the URL path rewrite rules. The values captured in asterisk can be retrieved by index. 规则定义 URL 路径重写规则。星号中捕获的值可以通过索引检索。 | (Required) （必需） |

### Examples 示例 

```go
package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/rewrite"
)

func main() {
  app := fiber.New()
  
  app.Use(rewrite.New(rewrite.Config{
    Rules: map[string]string{
      "/old":   "/new",
      "/old/*": "/new/$1",
    },
  }))
  
  app.Get("/new", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })
  app.Get("/new/*", func(c *fiber.Ctx) error {
    return c.SendString("Wildcard: " + c.Params("*"))
  })
  
  app.Listen(":3000")
}
```



**Test:
测试：**

```curl
curl http://localhost:3000/old
curl http://localhost:3000/old/hello
```
