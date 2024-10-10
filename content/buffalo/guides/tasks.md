+++
title = "任务"
date = 2024-02-04T22:22:23+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/tasks/]({{< ref "/buffalo/guides/tasks" >}})

# Tasks 任务 

Tasks are small scripts that are often needed when writing an application. These tasks might be along the lines of seeding a database, parsing a log file, or even a release script. Buffalo uses the [grift](https://github.com/gobuffalo/grift) package to make writing these tasks simple.

​	任务通常是在编写应用程序时经常需要的脚本。这些任务可能包括填充数据库、解析日志文件，甚至发布脚本。Buffalo 使用 grift 包来简化编写这些任务。

{{< vimeo "213096302">}}

## Writing Tasks 编写任务 

Tasks must all be in the `grifts` package. A simple task would look like following:

​	所有任务都必须在 `grifts` 包中。一个简单的任务如下所示：

```go
var _ = grift.Add("hello", func(c *grift.Context) error {
  fmt.Println("Hello!")
  return nil
})
```

## Tasks Generator 任务生成器 

```bash
$ buffalo g task foo:bar

--> grifts/bar.go
// grifts/bar.go
package grifts

import (
  . "github.com/gobuffalo/grift/grift"
)

var _ = Namespace("foo", func() {

  Desc("bar", "TODO")
  Add("bar", func(c *Context) error {
    return nil
  })

})
```

## Listing Available Tasks 列出可用任务 

```bash
$ buffalo task list

Available grifts
================
buffalo task db:seed       # Seeds a database
buffalo task middleware    # Prints out your middleware stack
buffalo task routes        # Print out all defined routes
buffalo task secret        # Generate a cryptographically secure secret key
```

## Running Tasks 运行任务 

### Development 开发 

Tasks can be run in development using the `buffalo task` command.

​	可以使用 `buffalo task` 命令在开发中运行任务。

```bash
$ buffalo task hello
```

### From a Built Binary 从已构建的二进制文件 

After a binary has been [built]({{< ref "/buffalo/deploy/packing" >}}), the tasks can be run with the `task` subcommand:

​	构建二进制文件后，可以使用 `task` 子命令运行任务：

```bash
$ myapp task hello
```
