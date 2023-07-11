+++
title = "prometheus"
weight = 140
date = 2023-07-09T21:56:27+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Prometheus

https://echo.labstack.com/docs/middleware/prometheus

> 注意
>
> ​	Echo 社区贡献
>

​	Prometheus 中间件为 HTTP 请求生成度量指标（metrics ）。

​	Prometheus 中间件有两个版本：

- 最新版本（推荐）https://github.com/labstack/echo-contrib/blob/master/echoprometheus.go
- 旧版本（已弃用）https://github.com/labstack/echo-contrib/blob/master/prometheus/prometheus.go

​	从旧版本迁移到较新版本的迁移指南可以在[这里](https://github.com/labstack/echo-contrib/blob/master/echoprometheus/README.md)找到。

## Usage

- 添加所需的模块 `go get -u github.com/labstack/echo-contrib`

- 添加 Prometheus 中间件和度量指标（metrics）服务路由

  ```go
  e := echo.New()
  e.Use(echoprometheus.NewMiddleware("myapp")) // 添加中间件以收集（gather）度量指标（metrics）
  e.GET("/metrics", echoprometheus.NewHandler()) // 添加路由以提供已收集的度量指标（metrics）
  ```

  

## 示例

Serve metric from the same server as where metrics is gathered

​	从与收集指标的相同服务器提供指标

​	在与收集指标的服务器相同的服务器上提供指标。

```go
package main

import (
    "errors"
    "github.com/labstack/echo-contrib/echoprometheus"
    "github.com/labstack/echo/v4"
    "log"
    "net/http"
)

func main() {
    e := echo.New()
    e.Use(echoprometheus.NewMiddleware("myapp")) // 添加中间件以收集度量指标
    e.GET("/metrics", echoprometheus.NewHandler()) // 添加路由以提供已收集的度量指标
    
    e.GET("/hello", func(c echo.Context) error {
        return c.String(http.StatusOK, "hello")
    })

    if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
        log.Fatal(err)
    }
}
```



​	在单独的端口上提供度量指标

```go
func main() {
    app := echo.New() // 此 Echo 实例将在端口 8080 上提供路由
    app.Use(echoprometheus.NewMiddleware("myapp")) // 添加中间件以收集度量指标

    go func() {
        metrics := echo.New() // 此 Echo 将在单独的端口 8081 上运行
        metrics.GET("/metrics", echoprometheus.NewHandler()) // 添加路由以提供已收集的度量指标
        if err := metrics.Start(":8081"); err != nil && !errors.Is(err, http.ErrServerClosed) {
            log.Fatal(err)
        }
    }()

    app.GET("/hello", func(c echo.Context) error {
        return c.String(http.StatusOK, "hello")
    })

    if err := app.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
        log.Fatal(err)
    }
}
```



*示例输出（针对第一个示例）*

```bash
curl http://localhost:8080/metrics

# HELP echo_request_duration_seconds The HTTP request latencies in seconds.
# TYPE echo_request_duration_seconds summary
echo_request_duration_seconds_sum 0.41086482
echo_request_duration_seconds_count 1
# HELP echo_request_size_bytes The HTTP request sizes in bytes.
# TYPE echo_request_size_bytes summary
echo_request_size_bytes_sum 56
echo_request_size_bytes_count 1
# HELP echo_requests_total How many HTTP requests processed, partitioned by status code and HTTP method.
# TYPE echo_requests_total counter
echo_requests_total{code="200",host="localhost:8080",method="GET",url="/"} 1
# HELP echo_response_size_bytes The HTTP response sizes in bytes.
# TYPE echo_response_size_bytes summary
echo_response_size_bytes_sum 61
echo_response_size_bytes_count 1
...
```



## Custom Configuration

### 提供自定义的 Prometheus Metrics

*使用方法*

​	用 Prometheus 默认注册表来使用自定义指标：

```go
package main

import (
    "errors"
    "github.com/labstack/echo-contrib/echoprometheus"
    "github.com/labstack/echo/v4"
    "github.com/prometheus/client_golang/prometheus"
    "log"
    "net/http"
)

func main() {
    e := echo.New()

    customCounter := prometheus.NewCounter( // 创建新的计数器指标。这是`prometheus.Metric`结构体的替代品
        prometheus.CounterOpts{
            Name: "custom_requests_total",
            Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
        },
    )
    if err := prometheus.Register(customCounter); err != nil { // 使用默认的指标注册表注册自定义计数器指标
        log.Fatal(err)
    }

    e.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
        AfterNext: func(c echo.Context, err error) {
            customCounter.Inc() // 在中间件中使用自定义指标。每个请求之后递增计数器
        },
    }))
    e.GET("/metrics", echoprometheus.NewHandler()) // 注册以获取已收集的度量指标的路由

    if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
        log.Fatal(err)
    }
}
```



或者创建自己的注册表并使用其注册自定义指标：

```go
package main

import (
    "errors"
    "github.com/labstack/echo-contrib/echoprometheus"
    "github.com/labstack/echo/v4"
    "github.com/prometheus/client_golang/prometheus"
    "log"
    "net/http"
)

func main() {
    e := echo.New()

    customRegistry := prometheus.NewRegistry() // 创建自定义注册表以容纳自定义指标
    customCounter := prometheus.NewCounter(    // 创建新的计数器指标。这是`prometheus.Metric`结构体的替代品
        prometheus.CounterOpts{
            Name: "custom_requests_total",
            Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
        },
    )
    if err := customRegistry.Register(customCounter); err != nil { //  使用指标注册表注册自定义计数器指标
        log.Fatal(err)
    }

    e.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
        AfterNext: func(c echo.Context, err error) {
            customCounter.Inc() // 在中间件中使用自定义指标。每个请求之后递增计数器
        },
        Registerer: customRegistry, // 使用自定义注册表而不是默认的 Prometheus 注册表
    }))
    e.GET("/metrics", echoprometheus.NewHandlerWithConfig(echoprometheus.HandlerConfig{Gatherer: customRegistry})) // 注册获取来自自定义注册表的收集的度量指标数据的路由

    if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
        log.Fatal(err)
    }
}
```



### 跳过 URL

*使用方法*

​	可以传递一个中间件跳过器以避免为某些 URL 生成度量指标

```go
package main

import (
    "errors"
    "github.com/labstack/echo-contrib/echoprometheus"
    "github.com/labstack/echo/v4"
    "log"
    "net/http"
    "strings"
)

func main() {
    e := echo.New()

    mwConfig := echoprometheus.MiddlewareConfig{
        Skipper: func(c echo.Context) bool {
            return strings.HasPrefix(c.Path(), "/testurl")
        }, // 不会在以 `/testurl` 开头的路由上收集度量指标
    }
    e.Use(echoprometheus.NewMiddlewareWithConfig(mwConfig)) // 添加中间件以收集度量指标

    e.GET("/metrics", echoprometheus.NewHandler()) // 添加路由以提供收集的度量指标

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

    if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
        log.Fatal(err)
    }
}
```



## 复杂场景

​	示例：修改默认的 `echoprometheus` 指标定义

```go
package main

import (
    "errors"
    "github.com/labstack/echo-contrib/echoprometheus"
    "github.com/labstack/echo/v4"
    "github.com/prometheus/client_golang/prometheus"
    "log"
    "net/http"
)

func main() {
    e := echo.New()

    e.Use(echoprometheus.NewMiddlewareWithConfig(echoprometheus.MiddlewareConfig{
        // 默认指标的标签可以通过 `LabelFuncs` 函数进行修改或添加
        LabelFuncs: map[string]echoprometheus.LabelValueFunc{
            "scheme": func(c echo.Context, err error) string { // 附加自定义标签
                return c.Scheme()
            },
            "host": func(c echo.Context, err error) string { // 覆盖默认的 'host' 标签值
                return "y_" + c.Request().Host
            },
        },

        // `echoprometheus` 中间件默认注册了以下指标：
        // - 直方图（Histogram）：request_duration_seconds
        // - 直方图（Histogram）：response_size_bytes
        // - 直方图（Histogram）：request_size_bytes
        // - 计数器（Counter）：requests_total
        // 这些指标可以通过 `HistogramOptsFunc` 和 `CounterOptsFunc` 函数进行修改
        HistogramOptsFunc: func(opts prometheus.HistogramOpts) prometheus.HistogramOpts {
            if opts.Name == "request_duration_seconds" {
                opts.Buckets = []float64{1000.0, 10_000.0, 100_000.0, 1_000_000.0} // 1KB ,10KB, 100KB, 1MB
            }
            return opts
        },
        CounterOptsFunc: func(opts prometheus.CounterOpts) prometheus.CounterOpts {
            if opts.Name == "requests_total" {
                opts.ConstLabels = prometheus.Labels{"my_const": "123"}
            }
            return opts
        },
    })) // 添加中间件以收集度量指标

    e.GET("/metrics", echoprometheus.NewHandler()) // 添加路由以提供收集的度量指标

    e.GET("/hello", func(c echo.Context) error {
        return c.String(http.StatusOK, "hello")
    })

    if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
        log.Fatal(err)
    }
}
```



