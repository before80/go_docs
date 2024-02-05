+++
title = "Log"
date = 2024-02-05T09:14:15+08:00
weight = 50
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文： [https://docs.gofiber.io/api/log]({{< ref "/fiber/API/Log" >}})

# 📃 Log

We can use logs to observe program behavior, diagnose problems, or configure corresponding alarms. And defining a well structured log can improve search efficiency and facilitate handling of problems.

​	我们可以使用日志来观察程序行为、诊断问题或配置相应的告警。并且定义一个结构良好的日志可以提高搜索效率，方便问题的处理。

Fiber provides a default way to print logs in the standard output. It also provides several global functions, such as `log.Info`, `log.Errorf`, `log.Warnw`, etc.

​	Fiber 提供了一种在标准输出中打印日志的默认方式。它还提供了一些全局函数，例如 `log.Info` 、 `log.Errorf` 、 `log.Warnw` 等。

## Log levels 日志级别 

```go
const (
    LevelTrace Level = iota
    LevelDebug
    LevelInfo
    LevelWarn
    LevelError
    LevelFatal
    LevelPanic
)
```



## Custom log 自定义日志 

Fiber provides the `AllLogger` interface for adapting the various log libraries.

​	Fiber 提供了 `AllLogger` 接口，用于适配各种日志库。

```go
type CommonLogger interface {
    Logger
    FormatLogger
    WithLogger
}

type AllLogger interface {
    CommonLogger
    ControlLogger
    WithLogger
}
```



## Print log 打印日志 

Note: The method of calling the Fatal level will interrupt the program running after printing the log, please use it with caution. Directly print logs of different levels, which will be entered into messageKey, the default is msg.

​	注意：调用 Fatal 级别的方法会在打印日志后中断程序运行，请谨慎使用。直接打印不同级别的日志，会进入 messageKey，默认是 msg。

```go
log.Info("Hello, World!")
log.Debug("Are you OK?")
log.Info("42 is the answer to life, the universe, and everything")
log.Warn("We are under attack!")
log.Error("Houston, we have a problem.")
log.Fatal("So Long, and Thanks for All the Fislog.")
log.Panic("The system is down.")
```



Format and print logs of different levels, all methods end with f

​	格式化并打印不同级别的日志，所有方法都以 f 结尾

```go
log.Debugf("Hello %s", "boy")
log.Infof("%d is the answer to life, the universe, and everything", 233)
log.Warnf("We are under attack %s!", "boss")
log.Errorf("%s, we have a problem.", "Master Shifu")
log.Fatalf("So Long, and Thanks for All the %s.", "banana")
```



Print a message with the key and value, or `KEYVALS UNPAIRED` if the key and value are not a pair.

​	打印带有键和值的消息，或者如果键和值不是一对，则打印 `KEYVALS UNPAIRED` 。

```go
log.Debugw("", "Hello", "boy")
log.Infow("", "number", 233)
log.Warnw("", "job", "boss")
log.Errorw("", "name", "Master Shifu")
log.Fatalw("", "fruit", "banana")
```



## Global log 全局日志 

If you are in a project and just want to use a simple log function that can be printed at any time in the global, we provide a global log.

​	如果您在一个项目中，只想使用一个可以在全局随时打印的简单日志功能，我们提供了一个全局日志。

```go
import "github.com/gofiber/fiber/v2/log"

log.Info("info")
log.Warn("warn")
```



The above is using the default `log.DefaultLogger` standard output. You can also find an already implemented adaptation under contrib, or use your own implemented Logger and use `log.SetLogger` to set the global log logger.

​	以上使用默认 `log.DefaultLogger` 标准输出。您还可以在 contrib 下找到已实现的改编，或使用您自己实现的 Logger 并使用 `log.SetLogger` 设置全局日志记录器。

```go
import (
    "log"
    fiberlog "github.com/gofiber/fiber/v2/log"
)

var _ log.AllLogger = (*customLogger)(nil)

type customLogger struct {
    stdlog *log.Logger
}

// ...
// inject your custom logger
fiberlog.SetLogger(customLogger)
```



## Set Level 设置级别 

`log.SetLevel` sets the level of logs below which logs will not be output. The default logger is LevelTrace.

​	 `log.SetLevel` 设置日志级别，低于该级别的日志将不会输出。默认记录器是 LevelTrace。

Note that this method is not **concurrent-safe**.

​	请注意，此方法不是并发安全的。

```go
import "github.com/gofiber/fiber/v2/log"

log.SetLevel(log.LevelInfo)
```



## Set output 设置输出 

`log.SetOutput` sets the output destination of the logger. The default logger types the log in the console.

​	 `log.SetOutput` 设置记录器的输出目标。默认记录器在控制台中键入日志。

```go
var logger AllLogger = &defaultLogger{
    stdlog: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
    depth:  4,
}
```



Set the output destination to the file.

​	将输出目标设置为文件。

```go
// Output to ./test.log file
f, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
if err != nil {
    return
}
log.SetOutput(f)
```



Set the output destination to the console and file.

​	将输出目标设置为控制台和文件。

```go
// Output to ./test.log file
file, _ := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
iw := io.MultiWriter(os.Stdout, file)
log.SetOutput(iw)
```



## Bind context 绑定上下文 

Set the context, using the following method will return a `CommonLogger` instance bound to the specified context

​	设置上下文，使用以下方法将返回一个绑定到指定上下文的 `CommonLogger` 实例

```go
commonLogger := log.WithContext(ctx)
commonLogger.Info("info")
```
