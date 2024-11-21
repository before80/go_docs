+++
title = "理解 Context"
date = 2024-11-21T08:12:38+08:00
weight = 60
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go-rod.github.io/i18n/zh-CN/#/understand-context](https://go-rod.github.io/i18n/zh-CN/#/understand-context)
>
> 收录该文档时间： `2024-11-21T08:12:38+08:00`

# 理解 Context

​	在此之前，先确保你已经学会了 [Goroutines](https://tour.golang.org/concurrency/1) 和 [Channels](https://tour.golang.org/concurrency/2)。 Context 主要用于在 Goroutines 之间传递上下文信息，包括：取消信号，超时，截止时间，键值对等等。

​	例如，我们现在有一个长时间运行的函数 `heartbeat`，它每秒会打印一次 `beat`：

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    heartbeat()
}

func heartbeat() {
    tick := time.Tick(time.Second)

    for {
        <-tick
        fmt.Println("beat")
    }
}
```

​	假如我们想在按下回车键的时候中断 heartbeat，我们可以这样修改代码：

```go
func main() {
    stop := make(chan struct{})
    go func() {
        fmt.Scanln()
        close(stop)
    }()

    heartbeat(stop)
}

func heartbeat(stop chan struct{}) {
    tick := time.Tick(time.Second)

    for {
        select {
        case <-tick:
        case <-stop:
            return
        }
        fmt.Println("beat")
    }
}
```

​	由于这种代码很常见，Golang 把它抽象为了一个 helper 包来处理这种情况，并称之为 [Context](https://golang.org/pkg/context/)。 现在用 Context 重新实现上面的代码：

```go
func main() {
    ctx, stop := context.WithCancel(context.Background())
    go func() {
        fmt.Scanln()
        stop()
    }()

    heartbeat(ctx)
}

func heartbeat(ctx context.Context) {
    tick := time.Tick(time.Second)

    for {
        select {
        case <-tick:
        case <-ctx.Done():
            return
        }
        fmt.Println("beat")
    }
}
```
