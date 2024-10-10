+++
title = "Timeout"
date = 2024-02-05T09:14:15+08:00
weight = 280
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/timeout]({{< ref "/fiber/API/Middleware/Timeout" >}})

# Timeout 超时

There exist two distinct implementations of timeout middleware [Fiber](https://github.com/gofiber/fiber).

​	存在两种不同的 Fiber 超时中间件实现。

**New**

Wraps a `fiber.Handler` with a timeout. If the handler takes longer than the given duration to return, the timeout error is set and forwarded to the centralized [ErrorHandler](https://docs.gofiber.io/error-handling).

​	用超时包装 `fiber.Handler` 。如果处理器花费的时间超过给定的持续时间才返回值，则会设定超时并将其转发到集中式 ErrorHandler。

CAUTION
注意

This has been deprecated since it raises race conditions.

​	此方法已弃用，因为它会产生竞争状况。

**NewWithContext**

As a `fiber.Handler` wrapper, it creates a context with `context.WithTimeout` and pass it in `UserContext`.

​	作为一个 `fiber.Handler` 包装器，它会创建一个包含 `context.WithTimeout` 的上下文，并将其传递给 `UserContext` 。

If the context passed executions (eg. DB ops, Http calls) takes longer than the given duration to return, the timeout error is set and forwarded to the centralized `ErrorHandler`.

​	如果传递的上下文（例如，数据库运算、Http 调用）花费的时间超过给定的持续时间才返回值，则会设定超时并将其转发到集中式 `ErrorHandler` 。

It does not cancel long running executions. Underlying executions must handle timeout by using `context.Context` parameter.

​	它不取消长时间运行的任务。基础层级必需使用 `context.Context` 参数来进行超时管理。

## Signatures 签名

```go
func New(handler fiber.Handler, timeout time.Duration, timeoutErrors ...error) fiber.Handler
func NewWithContext(handler fiber.Handler, timeout time.Duration, timeoutErrors ...error) fiber.Handler
```



## Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/timeout"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
func main() {
    app := fiber.New()

    h := func(c *fiber.Ctx) error {
        sleepTime, _ := time.ParseDuration(c.Params("sleepTime") + "ms")
        if err := sleepWithContext(c.UserContext(), sleepTime); err != nil {
            return fmt.Errorf("%w: execution error", err)
        }
        return nil
    }

    app.Get("/foo/:sleepTime", timeout.New(h, 2*time.Second))
    log.Fatal(app.Listen(":3000"))
}

func sleepWithContext(ctx context.Context, d time.Duration) error {
    timer := time.NewTimer(d)

    select {
    case <-ctx.Done():
        if !timer.Stop() {
            <-timer.C
        }
        return context.DeadlineExceeded
    case <-timer.C:
    }
    return nil
}
```



Test http 200 with curl:

​	使用 curl 测试 http 200：

```bash
curl --location -I --request GET 'http://localhost:3000/foo/1000' 
```



Test http 408 with curl:

​	使用 curl 测试 http 408：

```bash
curl --location -I --request GET 'http://localhost:3000/foo/3000' 
```



Use with custom error:

​	使用自定义的异常：

```go
var ErrFooTimeOut = errors.New("foo context canceled")

func main() {
    app := fiber.New()
    h := func(c *fiber.Ctx) error {
        sleepTime, _ := time.ParseDuration(c.Params("sleepTime") + "ms")
        if err := sleepWithContextWithCustomError(c.UserContext(), sleepTime); err != nil {
            return fmt.Errorf("%w: execution error", err)
        }
        return nil
    }

    app.Get("/foo/:sleepTime", timeout.NewWithContext(h, 2*time.Second, ErrFooTimeOut))
    log.Fatal(app.Listen(":3000"))
}

func sleepWithContextWithCustomError(ctx context.Context, d time.Duration) error {
    timer := time.NewTimer(d)
    select {
    case <-ctx.Done():
        if !timer.Stop() {
            <-timer.C
        }
        return ErrFooTimeOut
    case <-timer.C:
    }
    return nil
}
```



Sample usage with a DB call:

​	使用数据库调用的示例：

```go
func main() {
    app := fiber.New()
    db, _ := gorm.Open(postgres.Open("postgres://localhost/foodb"), &gorm.Config{})

    handler := func(ctx *fiber.Ctx) error {
        tran := db.WithContext(ctx.UserContext()).Begin()
        
        if tran = tran.Exec("SELECT pg_sleep(50)"); tran.Error != nil {
            return tran.Error
        }
        
        if tran = tran.Commit(); tran.Error != nil {
            return tran.Error
        }

        return nil
    }

    app.Get("/foo", timeout.NewWithContext(handler, 10*time.Second))
    log.Fatal(app.Listen(":3000"))
}
```
