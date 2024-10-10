+++
title = "Logger"
weight = 120
date = 2023-07-09T21:55:43+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Logger

> 原文：[https://echo.labstack.com/docs/middleware/logger](https://echo.labstack.com/docs/middleware/logger)

​	Logger 中间件用于记录每个 HTTP 请求的信息。

​	Echo 提供了两种不同的日志记录中间件：

- 较旧的基于字符串模板的日志记录器 [`Logger`](https://github.com/labstack/echo/blob/master/middleware/logger.go) —— 简单易用但功能有限
- 较新的可自定义的基于函数的日志记录器 [`RequestLogger`](https://github.com/labstack/echo/blob/master/middleware/request_logger.go) —— 允许开发人员完全自定义日志记录内容和方式，适用于与第三方日志记录库一起使用。

## 字符串模板

## Usage

```go
e.Use(middleware.Logger())
```



*示例输出*

```js
{"time":"2017-01-12T08:58:07.372015644-08:00","remote_ip":"::1","host":"localhost:1323","method":"GET","uri":"/","status":200,"error":"","latency":14743,"latency_human":"14.743µs","bytes_in":0,"bytes_out":2}
```



## 自定义配置

### Usage

```go
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
  Format: "method=${method}, uri=${uri}, status=${status}\n",
}))
```



​	上述示例使用了 `Format`，记录请求方法和请求 URI。

*示例输出*

```sh
method=GET, uri=/, status=200
```



## Configuration

```go
// LoggerConfig 定义了 Logger 中间件的配置。
LoggerConfig struct {
  // Skipper 定义了一个用于跳过中间件的函数。
  Skipper Skipper

  // Tags 用于构造日志记录格式。
  //
  // - time_unix
  // - time_unix_milli
  // - time_unix_micro
  // - time_unix_nano
  // - time_rfc3339
  // - time_rfc3339_nano
  // - time_custom
  // - id (Request ID)
  // - remote_ip
  // - uri
  // - host
  // - method
  // - path
  // - protocol
  // - referer
  // - user_agent
  // - status
  // - error
  // - latency (In nanoseconds)
  // - latency_human (Human readable)
  // - bytes_in (Bytes received)
  // - bytes_out (Bytes sent)
  // - header:<NAME>
  // - query:<NAME>
  // - form:<NAME>
  //
  // 示例 "${remote_ip} ${status}"
  //
  // 可选。默认值为 DefaultLoggerConfig.Format。
  Format string `yaml:"format"`

  // 可选。默认值为 DefaultLoggerConfig.CustomTimeFormat。
  CustomTimeFormat string `yaml:"custom_time_format"`

  // Output 是一个以 JSON 格式写入日志的writer。
  // 可选。默认值为 os.Stdout。
  Output io.Writer
}
```



### Default Configuration

```go
DefaultLoggerConfig = LoggerConfig{
  Skipper: DefaultSkipper,
  Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
    `"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
    `"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
    `,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
  CustomTimeFormat: "2006-01-02 15:04:05.00000",
}
```



## 可自定义的函数

​	RequestLogger 中间件允许开发人员完全自定义日志记录内容和方式，更适合与第三方（结构化日志记录）库一起使用。

​	请参阅 [RequestLoggerConfig](https://github.com/labstack/echo/blob/master/middleware/request_logger.go) 结构体的字段，了解日志记录器可提取的值。

### 示例

​	使用原生的 `fmt.Printf` 的示例：

```go
e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
    LogStatus: true,
    LogURI:    true,
    BeforeNextFunc: func(c echo.Context) {
        c.Set("customValueFromContext", 42)
    },
    LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
        value, _ := c.Get("customValueFromContext").(int)
        fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status, value)
        return nil
    },
}))
```



​	使用 Zerolog (https://github.com/rs/zerolog) 的示例：

```go
logger := zerolog.New(os.Stdout)
e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
    LogURI:    true,
    LogStatus: true,
    LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
        logger.Info().
            Str("URI", v.URI).
            Int("status", v.Status).
            Msg("request")

        return nil
    },
}))
```



​	使用 Zap (https://github.com/uber-go/zap) 的示例：

```go
logger, _ := zap.NewProduction()
e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
    LogURI:    true,
    LogStatus: true,
    LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
        logger.Info("request",
            zap.String("URI", v.URI),
            zap.Int("status", v.Status),
        )

        return nil
    },
}))
```



​	使用 Logrus (https://github.com/sirupsen/logrus) 的示例：

```go
log := logrus.New()
e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
    LogURI:    true,
    LogStatus: true,
    LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
        log.WithFields(logrus.Fields{
            "URI":   values.URI,
            "status": values.Status,
        }).Info("request")

        return nil
    },
}))
```