+++
title = "EnvVar"
date = 2024-02-05T09:14:15+08:00
weight = 80
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/envvar]({{< ref "/fiber/API/Middleware/EnvVar" >}})

# EnvVar

EnvVar middleware for [Fiber](https://github.com/gofiber/fiber) that can be used to expose environment variables with various options.

​	Fiber 的 EnvVar 中间件，可用于公开具有各种选项的环境变量。

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
  "github.com/gofiber/fiber/v2/middleware/envvar"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use("/expose/envvars", envvar.New())

// Or extend your config for customization
app.Use("/expose/envvars", envvar.New(
    envvar.Config{
        ExportVars:  map[string]string{"testKey": "", "testDefaultKey": "testDefaultVal"},
        ExcludeVars: map[string]string{"excludeKey": ""},
    }),
)
```



NOTE
注意

You will need to provide a path to use the envvar middleware.

​	您需要提供一个路径来使用 envvar 中间件。

## Response

Http response contract:

​	Http 响应契约：

```text
{
  "vars": {
    "someEnvVariable": "someValue",
    "anotherEnvVariable": "anotherValue",
  }
}
```



## Config 配置

| Property 属性 | Type 输入           | Description 说明                                             | Default 默认 |
| ------------- | ------------------- | ------------------------------------------------------------ | ------------ |
| ExportVars    | `map[string]string` | ExportVars specifies the environment variables that should be exported. ExportVars 指定应导出的环境变量。 | `nil`        |
| ExcludeVars   | `map[string]string` | ExcludeVars specifies the environment variables that should not be exported. ExcludeVars 指定不应导出的环境变量。 | `nil`        |

## Default Config 默认配置 

```go
Config{}
```
