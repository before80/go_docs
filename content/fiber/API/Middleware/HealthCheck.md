+++
title = "Health Check"
date = 2024-02-05T09:14:15+08:00
weight = 130
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/healthcheck]({{< ref "/fiber/API/Middleware/HealthCheck" >}})

# Health Check 运行状况检查

Liveness and readiness probes middleware for [Fiber](https://github.com/gofiber/fiber) that provides two endpoints for checking the liveness and readiness state of HTTP applications.

​	Liveness 和 readiness 探测中间件，用于 Fiber，它提供两个端点来检查 HTTP 应用程序的 liveness 和 readiness 状态。

## Overview 概述

- **Liveness Probe**: Checks if the server is up and running.

  ​	运行状况探测：检查服务器是否启动并运行。

  - **Default Endpoint**: `/livez`
    默认端点： `/livez`
  - **Behavior**: By default returns `true` immediately when the server is operational.
    行为：默认情况下，当服务器运行时立即返回 `true` 。

- **Readiness Probe**: Assesses if the application is ready to handle requests.

  ​	就绪探测：评估应用程序是否准备好处理请求。

  - **Default Endpoint**: `/readyz`
    默认端点： `/readyz`
  - **Behavior**: By default returns `true` immediately when the server is operational.
    行为：默认情况下，当服务器运行时立即返回 `true` 。

- **HTTP Status Codes**:

  ​	HTTP 状态代码：

  - `200 OK`: Returned when the checker function evaluates to `true`.
    `200 OK` ：当检查器函数评估为 `true` 时返回。
  - `503 Service Unavailable`: Returned when the checker function evaluates to `false`.
    `503 Service Unavailable` ：当检查器函数评估为 `false` 时返回。

## Signatures 签名

```go
func New(config Config) fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the [Fiber](https://github.com/gofiber/fiber) web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/healthcheck"
)
```



After you initiate your [Fiber](https://github.com/gofiber/fiber) app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Provide a minimal config
app.Use(healthcheck.New())

// Or extend your config for customization
app.Use(healthcheck.New(healthcheck.Config{
    LivenessProbe: func(c *fiber.Ctx) bool {
        return true
    },
    LivenessEndpoint: "/live",
    ReadinessProbe: func(c *fiber.Ctx) bool {
        return serviceA.Ready() && serviceB.Ready() && ...
    },
    ReadinessEndpoint: "/ready",
}))
```



## Config 配置

```go
type Config struct {
    // Next defines a function to skip this middleware when returned true.
    //
    // Optional. Default: nil
    Next func(c *fiber.Ctx) bool

    // Function used for checking the liveness of the application. Returns true if the application
    // is running and false if it is not. The liveness probe is typically used to indicate if 
    // the application is in a state where it can handle requests (e.g., the server is up and running).
    //
    // Optional. Default: func(c *fiber.Ctx) bool { return true }
    LivenessProbe HealthChecker

    // HTTP endpoint at which the liveness probe will be available.
    //
    // Optional. Default: "/livez"
    LivenessEndpoint string

    // Function used for checking the readiness of the application. Returns true if the application
    // is ready to process requests and false otherwise. The readiness probe typically checks if all necessary
    // services, databases, and other dependencies are available for the application to function correctly.
    //
    // Optional. Default: func(c *fiber.Ctx) bool { return true }
    ReadinessProbe HealthChecker

    // HTTP endpoint at which the readiness probe will be available.
    // Optional. Default: "/readyz"
    ReadinessEndpoint string
}
```



## Default Config 默认配置 

The default configuration used by this middleware is defined as follows:

​	此中间件使用的默认配置定义如下：

```go
func defaultLivenessProbe(*fiber.Ctx) bool { return true }

func defaultReadinessProbe(*fiber.Ctx) bool { return true }

var ConfigDefault = Config{
    LivenessProbe:     defaultLivenessProbe,
    ReadinessProbe:    defaultReadinessProbe,
    LivenessEndpoint:  "/livez",
    ReadinessEndpoint: "/readyz",
}
```
