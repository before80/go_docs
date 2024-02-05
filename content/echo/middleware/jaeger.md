+++
title = "jaeger"
weight = 90
date = 2023-07-09T21:55:04+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Jaeger

> 原文：[https://echo.labstack.com/docs/middleware/jaeger](https://echo.labstack.com/docs/middleware/jaeger)

> 注意
>
> ​	Echo社区贡献

​	使用Jaeger Tracing中间件对Echo框架的请求进行追踪。

## Usage

```go
package main
import (
    "github.com/labstack/echo-contrib/jaegertracing"
    "github.com/labstack/echo/v4"
)
func main() {
    e := echo.New()
    // 启用追踪中间件
    c := jaegertracing.New(e, nil)
    defer c.Close()

    e.Logger.Fatal(e.Start(":1323"))
}
```



​	启用追踪中间件会为每个请求创建一个追踪器（tracer）和一个根追踪跨度（root tracing span）。

## Custom Configuration

​	默认情况下，追踪会发送到`localhost` Jaeger代理实例。要配置外部的Jaeger，请使用环境变量启动应用程序。

### Usage

```bash
$ JAEGER_AGENT_HOST=192.168.1.10 JAEGER_AGENT_PORT=6831 ./myserver
```



​	追踪器（tracer）可以使用来自环境变量的值进行初始化。所有环境变量都是可选的，并且可以通过直接在配置对象上设置属性来覆盖它们。

| 属性                             | 描述                                                         |
| -------------------------------- | ------------------------------------------------------------ |
| JAEGER_SERVICE_NAME              | 服务名称                                                     |
| JAEGER_AGENT_HOST                | 通过UDP与代理进行通信的主机名                                |
| JAEGER_AGENT_PORT                | 通过UDP与代理进行通信的端口                                  |
| JAEGER_ENDPOINT                  | 发送跨度到收集器的HTTP端点，例如http://jaeger-collector:14268/api/traces |
| JAEGER_USER                      | 在收集器端点作为"Basic"身份验证的一部分发送的用户名          |
| JAEGER_PASSWORD                  | 在收集器端点作为"Basic"身份验证的一部分发送的密码            |
| JAEGER_REPORTER_LOG_SPANS        | 是否还记录跨度的报告器（reporter ）                          |
| JAEGER_REPORTER_MAX_QUEUE_SIZE   | 报告器的最大队列大小                                         |
| JAEGER_REPORTER_FLUSH_INTERVAL   | 报告器的刷新间隔，使用单位，例如"500ms"或"2s"（[有效单位][timeunits]） |
| JAEGER_SAMPLER_TYPE              | 采样器类型                                                   |
| JAEGER_SAMPLER_PARAM             | 采样器参数（数字）                                           |
| JAEGER_SAMPLER_MANAGER_HOST_PORT | 使用远程采样器时的HTTP端点，例如http://jaeger-agent:5778/sampling |
| JAEGER_SAMPLER_MAX_OPERATIONS    | 采样器（sampler ）将跟踪的操作的最大数量                     |
| JAEGER_SAMPLER_REFRESH_INTERVAL  | 远程控制的采样器将定期轮询jaeger-agent以获取适当的采样策略，使用单位，例如"1m"或"30s"（[有效单位][timeunits]） |
| JAEGER_TAGS                      | 以逗号分隔的`name=value`追踪器级别标签列表，将添加到所有报告的跨度中。该值还可以使用`${envVarName:default}`的格式引用环境变量，其中`:default`是可选的，并且如果找不到环境变量，则指定要使用的值 |
| JAEGER_DISABLED                  | 是否禁用追踪器。如果为true，则使用默认的`opentracing.NoopTracer` |
| JAEGER_RPC_METRICS               | 是否存储RPC指标（metrics ）                                  |

​	默认情况下，客户端通过UDP将跟踪发送到`localhost:6831`的代理。使用`JAEGER_AGENT_HOST`和`JAEGER_AGENT_PORT`将UDP跟踪发送到不同的`host:port`。如果设置了`JAEGER_ENDPOINT`，客户端将通过`HTTP`将跟踪发送到该端点，`JAEGER_AGENT_HOST`和`JAEGER_AGENT_PORT`将不再使用。如果`JAEGER_ENDPOINT`是安全的，可以通过设置`JAEGER_USER`和`JAEGER_PASSWORD`环境变量来执行HTTP基本身份验证。

### Skipping URL(s)

​	可以传递一个中间件跳过器（skipper）来避免对特定URL进行跟踪的跨度。

*用法*

```go
package main
import (
    "strings"
    "github.com/labstack/echo-contrib/jaegertracing"
    "github.com/labstack/echo/v4"
)

// urlSkipper 忽略某些中间件的指标（metrics）路由
func urlSkipper(c echo.Context) bool {
    if strings.HasPrefix(c.Path(), "/testurl") {
        return true
    }
    return false
}

func main() {
    e := echo.New()
    // 启用追踪中间件
    c := jaegertracing.New(e, urlSkipper)
    defer c.Close()

    e.Logger.Fatal(e.Start(":1323"))
}
```



### TraceFunction

​	这是一个包装（wrapper ）函数，用于在调用函数的整个持续时间内无缝添加一个跨度（span ）。无需更改函数实参。

*用法*

```go
package main
import (
    "github.com/labstack/echo-contrib/jaegertracing"
    "github.com/labstack/echo/v4"
    "net/http"
    "time"
)
func main() {
    e := echo.New()
    // 启用追踪中间件
    c := jaegertracing.New(e, nil)
    defer c.Close()
    e.GET("/", func(c echo.Context) error {
        // Wrap slowFunc on a new span to trace it's execution passing the function arguments
        // 使用新的跨度包装slowFunc以跟踪其执行，传递函数实参
        jaegertracing.TraceFunction(c, slowFunc, "Test String")
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}

// 要包装的函数。由于跟踪，无需更改其实参
func slowFunc(s string) {
    time.Sleep(200 * time.Millisecond)
    return
}
```



### CreateChildSpan

​	通过调用`CreateChildSpan`函数可以更精确地控制Span，从而控制要附加到Span的数据，如日志消息、行李（baggages ）和标签。

*用法*

```go
package main
import (
    "github.com/labstack/echo-contrib/jaegertracing"
    "github.com/labstack/echo/v4"
)
func main() {
    e := echo.New()
    // 启用追踪中间件
    c := jaegertracing.New(e, nil)
    defer c.Close()
    e.GET("/", func(c echo.Context) error {
        // 在创建子跨度（child span）之前执行某些操作
        time.Sleep(40 * time.Millisecond)
        sp := jaegertracing.CreateChildSpan(c, "Child span for additional processing")
        defer sp.Finish()
        sp.LogEvent("Test log")
        sp.SetBaggageItem("Test baggage", "baggage")
        sp.SetTag("Test tag", "New Tag")
        time.Sleep(100 * time.Millisecond)
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}
```



## 参考

- Opentracing库: https://github.com/opentracing/opentracing-go
- Jaeger配置: https://github.com/jaegertracing/jaeger-client-go#environment-variables