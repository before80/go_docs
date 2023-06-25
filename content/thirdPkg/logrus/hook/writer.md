+++
title = "Logrus 的 Writer 钩子"
date = 2023-06-25T09:29:23+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Writer Hooks for Logrus - Logrus 的 Writer 钩子

https://github.com/sirupsen/logrus/tree/v1.9.3/hooks/writer

​	将指定级别的日志发送到任何实现了 `io.Writer` 接口的对象。

## 用法

​	如果你想将高级别的日志发送到 `Stderr`，将正常执行的日志发送到 `Stdout`，可以按照以下方式实现：

```go
package main

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

func main() {
	log.SetOutput(ioutil.Discard) // 默认将所有日志发送到 nowhere

	log.AddHook(&writer.Hook{ // 将级别高于 warning 的日志发送到 stderr
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})
    
	log.AddHook(&writer.Hook{ // 将 info 和 debug 级别的日志发送到 stdout
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.InfoLevel,
			log.DebugLevel,
		},
	})
	log.Info("This will go to stdout")
	log.Warn("This will go to stderr")
}
```