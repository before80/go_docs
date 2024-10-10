+++
title = "log 模块"
date = 2024-02-04T09:31:23+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/module/logs/]({{< ref "/beego/modules/logs" >}})

# Logs Module 日志模块



## Logging 日志记录

The logging module is inspired by `database/sql`. It supports file, console, net and smtp as destination providers by default. It is installed like this:

​	日志记录模块的灵感来自 `database/sql` 。它默认支持文件、控制台、网络和 smtp 作为目标提供程序。它的安装方式如下：

```
go get github.com/beego/beego/v2/core/logs
```

## Basic Usage 基本用法

### General Usage 常规用法

Import package: 
​	导入包：

```
import (
	"github.com/beego/beego/v2/core/logs"
)
```

Initialize log variable (10000 is the cache size):

​	初始化日志变量（10000 是缓存大小）：

```
log := logs.NewLogger(10000)
```

Then add the output provider (it supports outputting to multiple providers at the same time). The first parameter is the provider name (`console`, `file`,`multifile`, `conn` , `smtp` or `es`).

​	然后添加输出提供程序（它支持同时向多个提供程序输出）。第一个参数是提供程序名称（ `console` 、 `file` 、 `multifile` 、 `conn` 、 `smtp` 或 `es` ）。

```
log.SetLogger("console")
```

The second parameter is a provider-specific configuration string (see below for details). logs.SetLogger(logs.AdapterFile,`{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)

​	第二个参数是特定于提供程序的配置字符串（有关详细信息，请参见下文）。 logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}` )

Then we can use it in our code:

​	然后我们可以在代码中使用它：

```
package main

import (
	"github.com/beego/beego/v2/core/logs"
)

func main() {    
	//an official log.Logger
	l := logs.GetLogger()
	l.Println("this is a message of http")
	//an official log.Logger with prefix ORM
	logs.GetLogger("ORM").Println("this is a message of orm")

    logs.Debug("my book is bought in the year of ", 2016)
 	logs.Info("this %s cat is %v years old", "yellow", 3)
 	logs.Warn("json is a type of kv like", map[string]int{"key": 2016})
   	logs.Error(1024, "is a very", "good game")
   	logs.Critical("oh,crash")
}
```

### Another Way 另一种方法

beego/logs supports to declare a single logger to use

​	beego/logs 支持声明一个要使用的单一记录器

```
    package main
    
    import (
    	"github.com/beego/beego/v2/core/logs"
    )
    
    func main() {
    	log := logs.NewLogger()
    	log.SetLogger(logs.AdapterConsole)
    	log.Debug("this is a debug message")
    }
```

## Logging caller information (file name & line number) 记录调用者信息（文件名和行号）

The module can be configured to include the file & line number of the log calls in the logging output. This functionality is disabled by default, but can be enabled using the following code:

​	该模块可以配置为在日志输出中包含日志调用的文件和行号。此功能默认情况下处于禁用状态，但可以使用以下代码启用：

```
logs.EnableFuncCallDepth(true)
```

Use `true` to turn file & line number logging on, and `false` to turn it off. Default is `false`.

​	使用 `true` 打开文件和行号日志记录，使用 `false` 关闭它。默认值为 `false` 。

If your application encapsulates the call to the log methods, you may need use `SetLogFuncCallDepth` to set the number of stack frames to be skipped before the caller information is retrieved. The default is 2.

​	如果您的应用程序封装了对日志方法的调用，您可能需要使用 `SetLogFuncCallDepth` 来设置在检索调用者信息之前要跳过的堆栈帧数。默认值为 2。

```
logs.SetLogFuncCallDepth(3)
```

## Logging asynchronously 异步记录

You can set logger to asynchronous logging to improve performance:

​	您可以将记录器设置为异步记录以提高性能：

```
logs.Async()
```

Add a parameter to set the length of buffer channel logs.Async(1e3)

​	添加一个参数来设置缓冲区通道的长度 logs.Async(1e3)

## Provider configuration 提供程序配置

Each provider supports a set of configuration options.

​	每个提供程序都支持一组配置选项。

- console 
  ​	控制台

  Can set output level or use default. Uses `os.Stdout` by default.

  ​	可以设置输出级别或使用默认值。默认情况下使用 `os.Stdout` 。

  ```
    logs.SetLogger(logs.AdapterConsole, `{"level":1}`)
  ```

- file

  E.g.: 
  ​	例如：

  ```
    logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)
  ```

  Parameters: 
  ​	参数：

  - filename: Save to filename.
    filename：保存到 filename。
  - maxlines: Maximum lines for each log file, 1000000 by default.
    maxlines：每个 日志文件的最大行数，默认值为 1000000。
  - maxsize: Maximum size of each log file, 1 « 28 or 256M by default.
    maxsize：每个 日志文件的最大大小，默认值为 1 « 28 或 256M。
  - daily: If log rotates by day, true by default.
    daily：如果按天轮转，则默认为 true。
  - maxdays: Maximum number of days log files will be kept, 7 by default.
    maxdays：将保留的 日志文件的最大天数，默认值为 7。
  - rotate: Enable logrotate or not, true by default.
    rotate：启用 logrotate 或不启用，默认值为 true。
  - level: Log level, Trace by default.
    level：默认值为 Trace 的 日志级别。
  - perm: Log file permission
    perm： 日志文件权限

- multifile

  E.g.: 
  ​	例如：

  ```
    logs.SetLogger(logs.AdapterMultiFile, ``{"filename":"test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}``)
  ```

  Parameters: 
  ​	参数：

  - filename: Save to filename.
    filename：保存到 filename。
  - maxlines: Maximum lines for each log file, 1000000 by default.
    maxlines：每个日志文件的最大行数，默认值为 1000000。
  - maxsize: Maximum size of each log file, 1 « 28 or 256M by default.
    maxsize：每个日志文件的最大大小，默认值为 1 « 28 或 256M。
  - daily: If log rotates by day, true by default.
    daily：如果日志按天轮换，则为 true，默认值为 true。
  - maxdays: Maximum number of days log files will be kept, 7 by default.
    maxdays：将保留日志文件的最大天数，默认值为 7。
  - rotate: Enable logrotate or not, true by default.
    rotate：启用或不启用 logrotate，默认值为 true。
  - level: Log level, Trace by default.
    level：日志级别，默认值为 Trace。
  - perm: Log file permission
    perm: 日志文件权限
  - separate: Log file will separate to test.error.log/test.debug.log as the log level set in the json array
    separate: 日志文件将按照 json 数组中日志级别所设，分别写入 test.error.log/test.debug.log

- conn

  Net output:

  ```
    logs.SetLogger(logs.AdapterConn, `{"net":"tcp","addr":":7020"}`)
  ```

  Parameters:

  - reconnectOnMsg: If true: reopen and close connection every time a message is sent. False by default.
    reconnectOnMsg: 如果为 true：每次发送信息时重新打开并关闭连接。默认值为 false。
  - reconnect: If true: auto connect. False by default.
    reconnect: 如果为 true：自动连接。默认值为 false。
  - net: connection type: tcp, unix or udp.
    net: 连接的种类：tcp、unix 或 udp。
  - addr: net connection address.
    addr: net 连接的通信端口。
  - level: Log level, Trace by default.
    level: 日志级别，默认值为 Trace。

- smtp

  Log by email: 
  ​	电子邮件日志：

  ```
    logs.SetLogger(logs.AdapterMail, `{"username":"beegotest@gmail.com","password":"xxxxxxxx","host":"smtp.gmail.com:587","sendTos":["xiemengjun@gmail.com"]}`)
  ```

  Parameters: 
  ​	参数：

  - username: smtp username.
    username：smtp 用户名。
  - password: smtp password.
    password：smtp 密码。
  - host: SMTP server host.
    host：SMTP 服务器主机。
  - sendTos: email addresses to which the logs will be sent.
    sendTos：将日志发送到的电子邮件地址。
  - subject: email subject, `Diagnostic message from server` by default.
    subject：电子邮件主题，默认情况下为 `Diagnostic message from server` 。
  - level: Log level, Trace by default.
    level：日志级别，默认情况下为跟踪。

- ElasticSearch

  Log to ElasticSearch: 
  ​	记录到 ElasticSearch：

  ```
    logs.SetLogger(logs.AdapterEs, `{"dsn":"http://localhost:9200/","level":1}`)
  ```

- JianLiao

  Log to JianLiao 
  ​	记录到 JianLiao

  ```
    logs.SetLogger(logs.AdapterJianLiao, `{"authorname":"xxx","title":"beego", "webhookurl":"https://jianliao.com/xxx", "redirecturl":"https://jianliao.com/xxx","imageurl":"https://jianliao.com/xxx","level":1}`)
  ```

- Slack

  Log to Slack 
  ​	记录到 Slack

  ```
    logs.SetLogger(logs.AdapterSlack, `{"webhookurl":"https://slack.com/xxx","level":1}`)
  ```

## Custom format logging 自定义格式记录

A new feature of the 2.0 release of beego is the ability to have custom formatting applied to your logs before being sent to your preferred adapter. Here is an example of it in use:

​	beego 2.0 版本的新功能是能够在将日志发送到首选适配器之前对日志应用自定义格式。以下是一个使用示例：

```go
package main

import (
	"fmt"

	beego "github.com/beego/beego/v2/pkg"
	"github.com/beego/beego/v2/pkg/logs"
)

type MainController struct {
	beego.Controller
}


func customFormatter(lm *logs.LogMsg) string {
	return fmt.Sprintf("[CUSTOM FILE LOGGING] %s", lm.Msg)
}

func GlobalFormatter(lm *logs.LogMsg) string {
	return fmt.Sprintf("[GLOBAL] %s", lm.Msg)
}

func main() {

	beego.BConfig.Log.AccessLogs = true

	// GlobalFormatter only overrides default log adapters. Hierarchy is like this:
	// adapter specific formatter > global formatter > default formatter
	logs.SetGlobalFormatter(GlobalFormatter)

	logs.SetLogger("console", "")

	logs.SetLoggerWithOpts("file", []string{`{"filename":"test.json"}`}, customFormatter)

	beego.Run()
}
```

## Global formatter 全局格式化程序

With the global formatter you can override and *default* logging formatters. This means that setting a global formatter will override any `logs.SetLogger()` adapters but will not override and `logs.SetLoggerWithOpts()` adapters. Default logging formatters are any adapters set using the following syntax:

​	使用全局格式化程序，您可以覆盖默认日志记录格式化程序。这意味着设置全局格式化程序将覆盖任何 `logs.SetLogger()` 适配器，但不会覆盖任何 `logs.SetLoggerWithOpts()` 适配器。默认日志记录格式化程序是使用以下语法设置的任何适配器：

```go
logs.SetLogger("adapterName", '{"key":"value"}')
```

## Adapter Specific formatters 适配器特定格式化程序

Apapter specific formatters can be set and will override any default or global formatter that has been set for a given adapter. Adapter specific logging formatters can be set using the following syntax:

​	可以设置特定于适配器的格式化程序，它将覆盖为给定适配器设置的任何默认或全局格式化程序。可以使用以下语法设置特定于适配器的日志记录格式化程序：

```go
logs.SetLoggerWithOpts("adapterName", []string{'{"key":"value"}'}, utils.KV{Key:"formatter", Value: formatterFunc})
```
