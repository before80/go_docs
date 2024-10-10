+++
title = "自定义"
weight = 20
date = 2023-07-09T21:49:53+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Customization - 自定义

> 原文：[https://echo.labstack.com/docs/customization](https://echo.labstack.com/docs/customization)

## 调试模式

​	`Echo#Debug` 可以用于启用/禁用调试模式。调试模式将日志级别设置为 `DEBUG`。

## 日志记录

​	日志记录的默认格式是 JSON，可以通过修改header进行更改。

### 日志 Header

​	`Echo#Logger.SetHeader(string)` 可以用于设置日志记录器的header。默认值为：

```js
{"time":"${time_rfc3339_nano}","level":"${level}","prefix":"${prefix}","file":"${short_file}","line":"${line}"}
```



*示例*

```go
import "github.com/labstack/gommon/log"

/* ... */

if l, ok := e.Logger.(*log.Logger); ok {
  l.SetHeader("${time_rfc3339} ${level}")
}
```



```sh
2018-05-08T20:30:06-07:00 INFO info
```



#### 可用标签

- `time_rfc3339`
- `time_rfc3339_nano`
- `level`
- `prefix`
- `long_file`
- `short_file`
- `line`

### 日志输出

​	`Echo#Logger.SetOutput(io.Writer)` 可以用于设置日志记录器的输出目标。默认值为 `os.Stdout`。

​	要完全禁用日志，请使用 `Echo#Logger.SetOutput(io.Discard)` 或 `Echo#Logger.SetLevel(log.OFF)`。

### 日志级别

​	`Echo#Logger.SetLevel(log.Lvl)` 可以用于设置日志记录器的日志级别。默认值为 `ERROR`。可能的取值有： 

- `DEBUG`
- `INFO`
- `WARN`
- `ERROR`
- `OFF`

### 自定义日志记录器

​	日志记录使用 `echo.Logger` 接口实现，允许您使用 `Echo#Logger` 注册自定义日志记录器。

## Startup Banner

​	`Echo#HideBanner` 可用于隐藏startup banner。

## 自定义监听器

​	`Echo#*Listener` 可用于运行自定义监听器。

*示例*

```go
l, err := net.Listen("tcp", ":1323")
if err != nil {
  e.Logger.Fatal(err)
}
e.Listener = l
e.Logger.Fatal(e.Start(""))
```



## 禁用HTTP/2

​	`Echo#DisableHTTP2` 可用于禁用 HTTP/2 协议。

## 读超时

​	`Echo#*Server#ReadTimeout` 可用于设置请求读取超时的最大持续时间。

## 写超时

​	`Echo#*Server#WriteTimeout` 可用于设置响应写入超时的最大持续时间。

## 验证器

​	`Echo#Validator` 可用于注册用于对请求有效载荷执行数据验证的验证器。

[了解更多](https://echo.labstack.com/docs/request#validate-data)

## 自定义绑定器

​	`Echo#Binder` 可用于注册自定义绑定器，用于绑定请求有效载荷。

[了解更多](https://echo.labstack.com/docs/request#custom-binder)

## 自定义 JSON 序列化器

​	`Echo#JSONSerializer` 可用于注册自定义 JSON 序列化器。

​	请查看 [json.go](https://github.com/labstack/echo/blob/master/json.go) 上的 `DefaultJSONSerializer`。

## 渲染器

​	`Echo#Renderer` 可用于注册模板渲染器。

[了解更多](https://echo.labstack.com/docs/templates)

## HTTP 错误处理器

​	`Echo#HTTPErrorHandler` 可用于注册自定义的 HTTP 错误处理器。

[了解更多](https://echo.labstack.com/docs/error-handling)