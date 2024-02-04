+++
title = "后台作业工作者"
date = 2024-02-04T21:16:50+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/workers/](https://gobuffalo.io/documentation/guides/workers/)

# Background Job Workers 后台作业工作者 

When building complex applications it is often nice to be able to run things in the “background”. While Go provides beautiful concurrency features, like the famed Goroutine, often one wants to run these on different machines, persist them using Redis, or any number of potential reasons why a simple Goroutine isn’t sufficient.

​	在构建复杂应用程序时，通常希望能够在“后台”运行某些内容。虽然 Go 提供了漂亮的并发特性，例如著名的 Goroutine，但通常希望在不同的机器上运行这些特性，使用 Redis 持久化这些特性，或者出于多种潜在原因，而简单的 Goroutine 不够用。

Workers shouldn’t be confused with [tasks](https://gobuffalo.io/documentation/guides/tasks): tasks are synchronous tools, whereas workers are intended to run asynchronously.
工作者不应与任务混淆：任务是同步工具，而工作者旨在异步运行。

## The Worker Interface 工作者接口 

In order to use background jobs, the `worker.Worker` interface must first be satisfied.

​	为了使用后台作业，必须首先满足 `worker.Worker` 接口。

```go
type Worker interface {
  // Start the worker with the given context
  Start(context.Context) error
  // Stop the worker
  Stop() error
  // Perform a job as soon as possibly
  Perform(Job) error
  // PerformAt performs a job at a particular time
  PerformAt(Job, time.Time) error
  // PerformIn performs a job after waiting for a specified amount of time
  PerformIn(Job, time.Duration) error
  // Register a Handler
  Register(string, Handler) error
}
```

Currently there are two official implementations of this interface:

​	目前，此接口有两个官方实现：

- The first is `worker.Simple`; It uses Goroutines to implement the interface. This is great for simple applications, but since the queues are not persisted, any jobs enqueued will be lost if the server was to be shut down. **This implementation is turned on by default**.
  第一个是 `worker.Simple` ；它使用 Goroutine 来实现该接口。这对于简单的应用程序来说非常棒，但由于队列不是持久的，因此如果服务器关闭，任何排队的作业都将丢失。此实现默认启用。
- The other implementation is the [`github.com/gobuffalo/gocraft-work-adapter`](https://github.com/gobuffalo/gocraft-work-adapter) package, which implements the [`github.com/gocraft/work`](https://github.com/gocraft/work) package using Redis as the backing store.
  另一个实现是 `github.com/gobuffalo/gocraft-work-adapter` 包，它使用 Redis 作为后备存储来实现 `github.com/gocraft/work` 包。

### Community implementations 社区实现 

The following Worker implementations are provided by Buffalo users (no official support):

​	Buffalo 用户提供了以下 Worker 实现（无官方支持）：

| Name 名称                                                    | Author 作者                                    | Description 说明                                             |
| :----------------------------------------------------------- | :--------------------------------------------- | :----------------------------------------------------------- |
| [AMQP worker adapter AMQP 工作器适配器](https://github.com/stanislas-m/amqp-work-adapter) | [@stanislas-m](https://github.com/stanislas-m) | A Worker implementation to use with AMQP-compatible brokers (such as [RabbitMQ](https://www.rabbitmq.com/)). 一个与 AMQP 兼容的代理（例如 RabbitMQ）配合使用的 Worker 实现。 |
| [Faktory worker adapter Faktory 工作器适配器](https://github.com/frankywahl/fwa) | [@frankywahl](https://github.com/frankywahl)   | A Worker implementation to use with [Faktory](https://contribsys.com/faktory/). 一个与 Faktory 配合使用的 Worker 实现。 |

## The Job type 作业类型 

A Job is a unit of work for a given Worker implementation.

​	作业是给定 Worker 实现的工作单元。

```go
// Args are the arguments passed into a job
type Args map[string]interface{}

// Job to be processed by a Worker
type Job struct {
  // Queue the job should be placed into
  Queue string
  // Args that will be passed to the Handler when run
  Args Args
  // Handler that will be run by the worker
  Handler string
}
```

## How to Use Background Tasks 如何使用后台任务 

To be able to use background tasks, you’ll need to setup a worker adapter, register job handlers and trigger jobs.

​	要能够使用后台任务，您需要设置一个工作器适配器，注册作业处理程序并触发作业。

### Setting Up a Worker Adapter 设置工作程序适配器 

When setting up your application you *can* assign a worker implementation to the `Worker` option.

​	在设置应用程序时，您可以将工作程序实现分配给 `Worker` 选项。

**app.go
app.go 请注意，如果您想使用基于 goroutine 的运行程序，此步骤是可选的。**

```go
import "github.com/gobuffalo/gocraft-work-adapter"
import "github.com/gomodule/redigo/redis"

// ...

app = buffalo.New(buffalo.Options{
  // ...
  Worker: gwa.New(gwa.Options{
    Pool: &redis.Pool{
      MaxActive: 5,
      MaxIdle:   5,
      Wait:      true,
      Dial: func() (redis.Conn, error) {
        return redis.Dial("tcp", ":6379")
      },
    },
    Name:           "myapp",
    MaxConcurrency: 25,
  }),
  // ...
})
```

Please note this step is optional, if you want to use the goroutines-based runner.

​	注册工作程序处理程序 #

### Registering a Worker Handler 处理程序是一个函数，它将运行以处理队列中给定类型的作业。这些处理程序必须首先向将运行它们的 worker 注册。

Handler is a function that will be run to process jobs for a given type in the queue. These handlers have to be first registered with the worker that will be running them.

​	每个处理程序都必须实现以下接口：

Each handler has to implement the following interface:

​	要将给定函数附加到作业类型，请使用 函数将其绑定到您的运行程序：

```go
// Handler function that will be run by the worker and given
// a slice of arguments
type Handler func(worker.Args) error
```

To attach a given function to a job type, bind it to your runner using the `Register` function:

​	排队作业 #

```go
import "github.com/gobuffalo/buffalo/worker"

var w worker.Worker

func init() {
  w = App().Worker // Get a ref to the previously defined Worker
  w.Register("send_email", func(args worker.Args) error {
    // do work to send an email
    return nil
  })
}
```

### Enqueueing a Job 现在工作程序处理程序已绑定，您需要将作业发送到队列。建议在排队作业时仅使用基本类型。例如，使用模型的 ID，而不是整个模型本身。

Now that the worker handlers are bound, you’ll need to send jobs to the queue. It is recommended to only use basic types when enqueueing a job. For example, use the ID of a model, and not the whole model itself.

You can choose to trigger jobs right now, or wait for a given time or duration.

​	您可以选择立即触发作业，或等待给定时间或持续时间。

#### `worker.Perform`

The `Perform` method enqueues the job, so the worker should try and run the job as soon as possible, based on the implementation of the worker itself.

​	 `Perform` 方法使作业入队，因此，根据工作器本身的实现，工作器应尝试尽快运行作业。

```go
func doWork() {
  // Send the send_email job to the queue, and process it as soon as possible.
  w.Perform(worker.Job{
    Queue: "default",
    Handler: "send_email",
    Args: worker.Args{
      "user_id": 123,
    },
  })
}
```

#### `worker.PerformIn`

The `PerformIn` method enqueues the job, so the worker should try and run the job after the duration has passed, based on the implementation of the worker itself.

​	 `PerformIn` 方法使作业入队，因此，根据工作器本身的实现，工作器应尝试在持续时间过去后运行作业。

```go
func doWork() {
  // Send the send_email job to the queue, and process it in 5 seconds.
  // Please note if no working unit is free at this time, it will wait for a free slot.
  w.PerformIn(worker.Job{
    Queue: "default",
    Handler: "send_email",
    Args: worker.Args{
      "user_id": 123,
    },
  }, 5 * time.Second)
}
```

#### `worker.PerformAt`

The `PerformAt` method enqueues the job, so the worker should try and run the job at (or near) the time specified, based on the implementation of the worker itself.

​	 `PerformAt` 方法使作业入队，因此，根据工作器本身的实现，工作器应尝试在指定时间（或接近指定时间）运行作业。

```go
func doWork() {
  // Send the send_email job to the queue, and process it at now + 5 seconds.
  // Please note if no working unit is free at this time, it will wait for a free slot.
  w.PerformAt(worker.Job{
    Queue: "default",
    Handler: "send_email",
    Args: worker.Args{
      "user_id": 123,
    },
  }, time.Now().Add(5 * time.Second))
}
```

### Starting and Stopping Workers 启动和停止工作器 

By default all Buffalo applications created will have a `main.go` that looks something like this:

​	默认情况下，创建的所有 Buffalo 应用程序都将具有一个 `main.go` ，如下所示：

```go
// cmd/app/main.go

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
```

The [`buffalo.App#Serve`](https://godoc.org/github.com/gobuffalo/buffalo#App.Serve) method will, by default, call the [`worker.Worker#Start`](https://godoc.org/github.com/gobuffalo/buffalo/worker#Worker) method for the registered worker. This will also call the [`worker.Worker#Stop`](https://godoc.org/github.com/gobuffalo/buffalo/worker#Worker) method when the application is shut down. This is the **recommended** approach for applications.

​	 `buffalo.App#Serve` 方法默认情况下将调用已注册工作器的 `worker.Worker#Start` 方法。这还将在应用程序关闭时调用 `worker.Worker#Stop` 方法。这是应用程序的推荐方法。

If you don’t want your workers to start automatically, you can set the option [`buffalo.Options#WorkerOff`](https://godoc.org/github.com/gobuffalo/buffalo#Options) to `true` when setting up your application.

​	如果您不希望工作器自动启动，则可以在设置应用程序时将选项 `buffalo.Options#WorkerOff` 设置为 `true` 。