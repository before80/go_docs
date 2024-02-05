+++
title = "Monitor"
date = 2024-02-05T09:14:15+08:00
weight = 190
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/middleware/monitor]({{< ref "/fiber/API/Middleware/Monitor" >}})

# Monitor 监视器

Monitor middleware for [Fiber](https://github.com/gofiber/fiber) that reports server metrics, inspired by [express-status-monitor](https://github.com/RafalWilinski/express-status-monitor)

​	受 express-status-monitor 启发的 Fiber 服务器指标报告监视中间件

CAUTION
注意

Monitor is still in beta, API might change in the future!

​	Monitor 仍处于测试阶段，API 可能会在未来发生变化！

![img](https://i.imgur.com/nHAtBpJ.gif)

### Signatures 签名

```go
func New() fiber.Handler
```



### Examples 示例 

Import the middleware package that is part of the Fiber web framework

​	导入 Fiber Web 框架的一部分中间件包

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/monitor"
)
```



After you initiate your Fiber app, you can use the following possibilities:

​	在启动 Fiber 应用后，您可以使用以下可能性：

```go
// Initialize default config (Assign the middleware to /metrics)
app.Get("/metrics", monitor.New())

// Or extend your config for customization
// Assign the middleware to /metrics
// and change the Title to `MyService Metrics Page`
app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
```



You can also access the API endpoint with `curl -X GET -H "Accept: application/json" http://localhost:3000/metrics` which returns:

​	您还可以使用 `curl -X GET -H "Accept: application/json" http://localhost:3000/metrics` 访问 API 端点，它返回：

```json
{"pid":{ "cpu":0.4568381746582226, "ram":20516864,   "conns":3 },
 "os": { "cpu":8.759124087593099,  "ram":3997155328, "conns":44,
    "total_ram":8245489664, "load_avg":0.51 }}
```



## Config 配置

| Property 属性         | Type 输入               | Description 说明                                             | Default 默认                                                 |
| --------------------- | ----------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| Title 标题            | `string`                | Metrics page title 指标页面标题                              | "Fiber Monitor" “Fiber Monitor”                              |
| Refresh 刷新          | `time.Duration`         | Refresh period 刷新周期                                      | 3 seconds 3 秒                                               |
| APIOnly               | `bool`                  | Whether the service should expose only the monitoring API 服务是否应该仅公开监控 API | false                                                        |
| Next 下一步           | `func(*fiber.Ctx) bool` | Next defines a function to skip this middleware when returned true. 接下来定义一个函数，在返回 true 时跳过此中间件。 | `nil`                                                        |
| CustomHead 自定义头部 | `string`                | Custom HTML Code to Head Section(Before End) 自定义 HTML 代码到头部部分（结束前） | empty 空                                                     |
| FontURL 字体 URL      | `string`                | FontURL for specify font resource path or URL 字体 URL 用于指定字体资源路径或 URL | "https://fonts.googleapis.com/css2?family=Roboto:wght@400;900&display=swap" “https://fonts.googleapis.com/css2?family=Roboto:wght@400;900&display=swap” |
| ChartJsURL            | `string`                | ChartJsURL for specify ChartJS library path or URL ChartJsURL 用于指定 ChartJS 库路径或 URL | "https://cdn.jsdelivr.net/npm/chart.js@2.9/dist/Chart.bundle.min.js" |

## Default Config 默认配置 

```go
var ConfigDefault = Config{
    Title:      defaultTitle,
    Refresh:    defaultRefresh,
    FontURL:    defaultFontURL,
    ChartJsURL: defaultChartJSURL,
    CustomHead: defaultCustomHead,
    APIOnly:    false,
    Next:       nil,
    index: newIndex(viewBag{
        defaultTitle,
        defaultRefresh,
        defaultFontURL,
        defaultChartJSURL,
        defaultCustomHead,
    }),
}
```
