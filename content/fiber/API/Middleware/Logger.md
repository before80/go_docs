+++
title = "Logger"
date = 2024-02-05T09:14:15+08:00
weight = 180
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/logger]({{< ref "/fiber/API/Middleware/Logger" >}})

# Logger 日志记录器

Logger middleware for [Fiber](https://github.com/gofiber/fiber) that logs HTTP request/response details.

​	Fiber 的日志记录器中间件记录 HTTP 请求/响应详细信息。

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
    "github.com/gofiber/fiber/v2/middleware/logger"
)
```



TIP

The order of registration plays a role. Only all routes that are registered after this one will be logged. The middleware should therefore be one of the first to be registered.

​	注册顺序起作用。只有在此之后注册的所有路由才会被记录。因此，中间件应该是最早注册的中间件之一。

After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config
app.Use(logger.New())

// Or extend your config for customization
// Logging remote IP and Port
app.Use(logger.New(logger.Config{
    Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
}))

// Logging Request ID
app.Use(requestid.New())
app.Use(logger.New(logger.Config{
    // For more options, see the Config section
    Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
}))

// Changing TimeZone & TimeFormat
app.Use(logger.New(logger.Config{
    Format:     "${pid} ${status} - ${method} ${path}\n",
    TimeFormat: "02-Jan-2006",
    TimeZone:   "America/New_York",
}))

// Custom File Writer
file, err := os.OpenFile("./123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
if err != nil {
    log.Fatalf("error opening file: %v", err)
}
defer file.Close()
app.Use(logger.New(logger.Config{
    Output: file,
}))

// Add Custom Tags
app.Use(logger.New(logger.Config{
    CustomTags: map[string]logger.LogFunc{
        "custom_tag": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
            return output.WriteString("it is a custom tag")
        },
    },
}))

// Callback after log is written
app.Use(logger.New(logger.Config{
    TimeFormat: time.RFC3339Nano,
    TimeZone:   "Asia/Shanghai",
    Done: func(c *fiber.Ctx, logString []byte) {
        if c.Response().StatusCode() != fiber.StatusOK {
            reporter.SendToSlack(logString) 
        }
    },
}))

// Disable colors when outputting to default format
app.Use(logger.New(logger.Config{
    DisableColors: true,
}))
```



## Config 配置

### Config 配置

| Property 属性             | Type 输入                  | Description 说明                                             | Default 默认             |
| ------------------------- | -------------------------- | ------------------------------------------------------------ | ------------------------ |
| Next 下一步               | `func(*fiber.Ctx) bool`    | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                    |
| Done                      | `func(*fiber.Ctx, []byte)` | Done is a function that is called after the log string for a request is written to Output, and pass the log string as parameter. Done 是一个函数，它在请求的日志字符串被写入 Output 后被调用，并将日志字符串作为参数传递。 | `nil`                    |
| CustomTags                | `map[string]LogFunc`       | tagFunctions defines the custom tag action. tagFunctions 定义自定义标签操作。 | `map[string]LogFunc`     |
| Format                    | `string`                   | Format defines the logging tags. Format 定义日志记录标签。   | `${time}                 |
| TimeFormat                | `string`                   | TimeFormat defines the time format for log timestamps. TimeFormat 定义日志时间戳的时间格式。 | `15:04:05`               |
| TimeZone 时区             | `string`                   | TimeZone can be specified, such as "UTC" and "America/New_York" and "Asia/Chongqing", etc 时区可以指定，例如“UTC”、“America/New_York”和“Asia/Chongqing”等 | `"Local"`                |
| TimeInterval 时间间隔     | `time.Duration`            | TimeInterval is the delay before the timestamp is updated. 时间间隔是在时间戳更新之前延迟的时间。 | `500 * time.Millisecond` |
| Output 输出               | `io.Writer`                | Output is a writer where logs are written. 输出是一个写入日志的编写器。 | `os.Stdout`              |
| DisableColors 禁用颜色    | `bool`                     | DisableColors defines if the logs output should be colorized. 禁用颜色定义是否应给日志输出着色。 | `false`                  |
| enableColors 启用颜色     | `bool`                     | Internal field for enabling colors in the log output. (This is not a user-configurable field) 日志输出中启用颜色的内部字段。（这不是用户可配置字段） | -                        |
| enableLatency 启用延迟    | `bool`                     | Internal field for enabling latency measurement in logs. (This is not a user-configurable field) 用于在日志中启用时差测量的内部字段。（这不是用户可配置字段） | -                        |
| timeZoneLocation 时区位置 | `*time.Location`           | Internal field for the time zone location. (This is not a user-configurable field) 时区位置的内部字段。（这不是用户可配置字段） | -                        |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Next:          nil,
    Done:          nil,
    Format:        "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
    TimeFormat:    "15:04:05",
    TimeZone:      "Local",
    TimeInterval:  500 * time.Millisecond,
    Output:        os.Stdout,
    DisableColors: false,
}
```



## Constants 常量 

```go
// Logger variables
const (
    TagPid               = "pid"
    TagTime              = "time"
    TagReferer           = "referer"
    TagProtocol          = "protocol"
    TagPort              = "port"
    TagIP                = "ip"
    TagIPs               = "ips"
    TagHost              = "host"
    TagMethod            = "method"
    TagPath              = "path"
    TagURL               = "url"
    TagUA                = "ua"
    TagLatency           = "latency"
    TagStatus            = "status"         // response status
    TagResBody           = "resBody"        // response body
    TagReqHeaders        = "reqHeaders"
    TagQueryStringParams = "queryParams"    // request query parameters
    TagBody              = "body"           // request body
    TagBytesSent         = "bytesSent"
    TagBytesReceived     = "bytesReceived"
    TagRoute             = "route"
    TagError             = "error"
    // DEPRECATED: Use TagReqHeader instead
    TagHeader            = "header:"        // request header
    TagReqHeader         = "reqHeader:"     // request header
    TagRespHeader        = "respHeader:"    // response header
    TagQuery             = "query:"         // request query
    TagForm              = "form:"          // request form
    TagCookie            = "cookie:"        // request cookie
    TagLocals            = "locals:"
    // colors
    TagBlack             = "black"
    TagRed               = "red"
    TagGreen             = "green"
    TagYellow            = "yellow"
    TagBlue              = "blue"
    TagMagenta           = "magenta"
    TagCyan              = "cyan"
    TagWhite             = "white"
    TagReset             = "reset"
)
```
