+++
title = "Logrus 的 Syslog 钩子"
date = 2023-06-25T09:29:33+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Syslog Hooks for Logrus - Logrus 的 Syslog 钩子

> 原文：[https://github.com/sirupsen/logrus/tree/v1.9.3/hooks/writer](https://github.com/sirupsen/logrus/tree/v1.9.3/hooks/writer)

## 用法

```go
import (
  "log/syslog"
  "github.com/sirupsen/logrus"
  lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

func main() {
  log       := logrus.New()
  hook, err := lSyslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")

  if err == nil {
    log.Hooks.Add(hook)
  }
}
```



If you want to connect to local syslog (Ex. "/dev/log" or "/var/run/syslog" or "/var/run/log"). Just assign empty string to the first two parameters of `NewSyslogHook`. It should look like the following.

​	如果你想连接到本地的 syslog（例如 "/dev/log" 或 "/var/run/syslog" 或 "/var/run/log"），只需将第一个和第二个参数传递为空字符串给 `NewSyslogHook` 函数。示例如下：

```go
import (
  "log/syslog"
  "github.com/sirupsen/logrus"
  lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

func main() {
  log       := logrus.New()
  hook, err := lSyslog.NewSyslogHook("", "", syslog.LOG_INFO, "")

  if err == nil {
    log.Hooks.Add(hook)
  }
}
```



### 本地和远程日志记录的不同日志级别

​	默认情况下，`NewSyslogHook()` 函数会将所有日志级别的日志通过钩子发送。如果你想在本地日志记录和 syslog 日志记录之间使用不同的日志级别（即要遵循传递给 `NewSyslogHook()` 的 `priority` 参数），你需要实现 `logrus_syslog.SyslogHook` 接口，并重写 `Levels()` 方法，只返回你感兴趣的日志级别。

​	下面的示例展示了如何在本地日志记录时使用 **DEBUG** 级别，在 syslog 日志记录时使用 **WARN** 级别：

```go
package main

import (
	"log/syslog"

	log "github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

type customHook struct {
	*logrus_syslog.SyslogHook
}

func (h *customHook) Levels() []log.Level {
	return []log.Level{log.WarnLevel}
}

func main() {
	log.SetLevel(log.DebugLevel)

	hook, err := logrus_syslog.NewSyslogHook("tcp", "localhost:5140", syslog.LOG_WARNING, "myTag")
	if err != nil {
		panic(err)
	}

	log.AddHook(&customHook{hook})

	//...
}
```