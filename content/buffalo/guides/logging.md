+++
title = "日志记录"
date = 2024-02-04T21:18:32+08:00
weight = 12
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/logging/]({{< ref "/buffalo/guides/templateEngines" >}})

# Logging 日志记录 

Buffalo logs are managed using the [logrus](https://github.com/sirupsen/logrus) package.

​	Buffalo 日志使用 logrus 包进行管理。

## Defaults 默认值 

The default logger outputs logs in a human-readable format:

​	默认记录器以人类可读的格式输出日志：

```plain
INFO[2020-02-21T07:42:34+01:00] /en/ content_type=text/html duration=26.189949ms human_size="21 kB" method=GET params="{\"lang\":[\"en\"]}" path=/en/ render=22.730816ms request_id=9b8d9260225fe99609a2-7cc679f4ae458b9925e3 size=21182 status=200
```

## Customize the logger 自定义记录器 

```go
// JSONLogger wraps a logrus JSON logger into a buffalo Logger
func JSONLogger(lvl logger.Level) logger.FieldLogger {
    l := logrus.New()
    l.Level = lvl
    l.SetFormatter(&logrus.JSONFormatter{})
    l.SetOutput(os.Stdout)
    return logger.Logrus{FieldLogger: l}
}

//... 

app = buffalo.New(buffalo.Options{
// ...
    Logger:       JSONLogger(logger.DebugLevel),
}
```
