+++
title = "事件"
date = 2024-02-04T21:18:02+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/events/](https://gobuffalo.io/documentation/guides/events/)

# Events 事件 

Since **0.13.0-beta.2**
自 0.13.0-beta.2 起



The [events](https://pkg.go.dev/github.com/gobuffalo/events) package allows for Go applications, including Buffalo applications, to listen, and emit, global event messages.

​	事件包允许 Go 应用程序（包括 Buffalo 应用程序）监听和发出全局事件消息。

## Listening for Events 监听事件 

To start listening for events a [events#Listener](https://pkg.go.dev/github.com/gobuffalo/events#Listener) must first be registered with the [events](https://pkg.go.dev/github.com/gobuffalo/events) package.

​	要开始监听事件，必须先将 events#Listener 注册到事件包中。

```go
func init() {
  _, err := events.Listen(func(e events.Event) {
    // do work
  })
}
```

Once registered this new listener function will be sent all events emitted through the [events](https://pkg.go.dev/github.com/gobuffalo/events) package.

​	注册后，此新侦听器函数将被发送通过事件包发出的所有事件。

## Emitting Events 发出事件 

When emitting events the `Kind` attribute should be a unique, but constant, string. It is this attribute that users will use to determine how to respond to events they receive.

​	发出事件时， `Kind` 属性应为唯一但不变的字符串。用户将使用此属性来确定如何响应他们收到的事件。

It is recommended to namespace this attribute like such, with error events being suffixed with `:err`.

​	建议像这样对该属性进行命名空间，错误事件以 `:err` 结尾。

```plain
"<package-name>:<additional-names>:<optional-error>"
"myapp:foo:start"
"myapp:foo:stop"
"mypkg:workers:bar:err"
```

This naming pattern makes it easier for users to filter events to only those that they care about. See [Filtering Events](https://gobuffalo.io/documentation/guides/events/#filtering-events) for more details.

​	这种命名模式使用户可以更轻松地将事件筛选为仅他们关心的事件。有关更多详细信息，请参阅筛选事件。

------

There are multiple ways to emit an [events#Event](https://pkg.go.dev/github.com/gobuffalo/events#Event) in your Go code. The [events#EmitError](https://pkg.go.dev/github.com/gobuffalo/events#EmitError) and [events#EmitPayload](https://pkg.go.dev/github.com/gobuffalo/events#EmitPayload) functions both accept a `payload interface{}` argument. It is recommended to use [events#Payload](https://pkg.go.dev/github.com/gobuffalo/events#Payload) for payloads; any other type passed in will get converted into a [events#Payload](https://pkg.go.dev/github.com/gobuffalo/events#Payload) with the argument set in the payload with the key, `data`.

​	在 Go 代码中有多种方式可以发出 events#Event。events#EmitError 和 events#EmitPayload 函数都接受 `payload interface{}` 参数。建议对有效负载使用 events#Payload；传入的任何其他类型都将转换为 events#Payload，其中参数在有效负载中设置，键为 `data` 。

- [events#Emit](https://pkg.go.dev/github.com/gobuffalo/events#Emit)

```go
func MyHandler(c buffalo.Context) error {
  e := events.Event{
    Kind:    "coke:myhandler:hello",
    Message: "hi!",
    Payload: events.Payload{"context": c},
  }
  if err := events.Emit(e); err != nil {
    return err
  }
  return c.Render(200, r.HTML("index.html"))
}
```

- [events#EmitError](https://pkg.go.dev/github.com/gobuffalo/events#EmitError)

```go
func MyHandler(c buffalo.Context) error {
  if err := events.EmitError("coke:myhandler:hello:err", errors.New("boom"), c); err != nil {
    return err
  }
  return c.Render(200, r.HTML("index.html"))
}
```

- [events#EmitPayload](https://pkg.go.dev/github.com/gobuffalo/events#EmitPayload)

```go
func MyHandler(c buffalo.Context) error {
  p := events.Payload{
    "message": "hi!",
  }
  if err := events.EmitPayload("coke:myhandler:hello", p); err != nil {
    return err
  }
  return c.Render(200, r.HTML("index.html"))
}
```

## Filtering Events 筛选事件 

In the [Emitting Events](https://gobuffalo.io/documentation/guides/events/#emitting-events) section the naming convention for [events#Event.Kind](https://pkg.go.dev/github.com/gobuffalo/events#Event.Kind) is described. By the checking the value of [events#Event.Kind](https://pkg.go.dev/github.com/gobuffalo/events#Event.Kind).

​	在发出事件部分中，描述了 events#Event.Kind 的命名约定。通过检查 events#Event.Kind 的值。

direct match
直接匹配

matching with a switch statement
使用 switch 语句进行匹配

matching error events
匹配错误事件

matching on prefix
匹配前缀

```go
// direct match
events.Listen(func(e events.Event) {
  if e.Kind != buffalo.EvtRouteStarted {
    // do nothing
    return
  }
  // do work on the route event
})
```

## Stop Listening for Events 停止监听事件 

When registering a new [events#Listener](https://pkg.go.dev/github.com/gobuffalo/events#Listener) a [events#DeleteFn](https://pkg.go.dev/github.com/gobuffalo/events#DeleteFn) is returned. This function should be held on to and used when you want to remove the added listener.

​	注册新的 events#Listener 时，会返回一个 events#DeleteFn。当您想要移除已添加的监听器时，应保留此函数并使用它。

```go
deleteFn, err := events.Listen(func(e events.Event) {
  // do work
})
if err != nil {
  return err
}
defer deleteFn()
```

## Listening with Plugins 使用插件进行监听 

To enable a plugin to a receive a JSON version of emitted events, the plugin can set the [events#Command.BuffaloCommand](https://pkg.go.dev/github.com/gobuffalo/buffalo-plugins/plugins#Command.BuffaloCommand) value to `events` when listing the `available` commands for the plugin.

​	为了使插件能够接收已发送事件的 JSON 版本，该插件可以在列出插件的 `available` 命令时将 events#Command.BuffaloCommand 值设置为 `events` 。

availableCmd

listenCmd

```go
// availableCmd
var availableCmd = &cobra.Command{
  Use:   "available",
  Short: "a list of available buffalo plugins",
  RunE: func(cmd *cobra.Command, args []string) error {
    plugs := plugins.Commands{
      {Name: "listen", UseCommand: "listen", BuffaloCommand: "events", Description: listenCmd.Short, Aliases: listenCmd.Aliases},
    }
    return json.NewEncoder(os.Stdout).Encode(plugs)
  },
}
```

## Integrating a Messaging Queue 集成消息队列 

It is often desirable to take events emitted and send them to a message queue, such as Kafka or Redis, to be processed externally. The [events](https://pkg.go.dev/github.com/gobuffalo/events) package does not have a directhook for this sort of functionality, the most direct way of enabling this behavior is to register a [events#Listener](https://pkg.go.dev/github.com/gobuffalo/events#Listener) that can then hand the event over to the appropriate message queue.

​	通常需要获取已发送的事件并将它们发送到消息队列（例如 Kafka 或 Redis）以供外部处理。events 包没有针对此类功能的直接挂钩，启用此行为的最直接方法是注册一个 events#Listener，然后该监听器可以将事件交给相应的消息队列。

```go
events.Listen(func(e events.Event) {
  myMessageQ.DoWork(e)
})
```

## Known Events 已知事件 

### Application Events 应用程序事件 

The following events are known to be emitted by Buffalo during the application lifecyle.

​	已知 Buffalo 在应用程序生命周期期间会发出以下事件。

| Constant 常量               | String 字符串                | Emitted When 发出时间                                        | Payload 有效负载                                             |
| --------------------------- | ---------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `buffalo.EvtAppStart`       | `"buffalo:app:start"`        | [buffalo#App.Serve](https://pkg.go.dev/github.com/gobuffalo/buffalo#App.Serve) is called 调用 buffalo#App.Serve | `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) |
| `buffalo.EvtAppStartErr`    | `"buffalo:app:start:err"`    | an error occurs calling [buffalo#App.Serve](https://pkg.go.dev/github.com/gobuffalo/buffalo#App.Serve) 调用 buffalo#App.Serve 时发生错误 | `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) |
| `buffalo.EvtAppStop`        | `"buffalo:app:stop"`         | [buffalo#App.Stop](https://pkg.go.dev/github.com/gobuffalo/buffalo#App.Stop) is called buffalo#App.Stop 被调用 | `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) |
| `buffalo.EvtAppStopErr`     | `"buffalo:app:stop:err"`     | an error occurs calling [buffalo#App.Stop](https://pkg.go.dev/github.com/gobuffalo/buffalo#App.Stop) 调用 buffalo#App.Stop 时发生错误 | `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) |
| `buffalo.EvtRouteStarted`   | `"buffalo:route:started"`    | a requested route is being processed 正在处理请求的路由      | `route`: [buffalo#RouteInfo](https://pkg.go.dev/github.com/gobuffalo/buffalo#RouteInfo) `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) `context`: [buffalo#Context](https://pkg.go.dev/github.com/gobuffalo/buffalo#Context) |
| `buffalo.EvtRouteFinished`  | `"buffalo:route:finished"`   | a requested route is completed 请求的路由已完成              | `route`: [buffalo#RouteInfo](https://pkg.go.dev/github.com/gobuffalo/buffalo#RouteInfo) `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) `context`: [buffalo#Context](https://pkg.go.dev/github.com/gobuffalo/buffalo#Context) |
| `buffalo.EvtRouteErr`       | `"buffalo:route:err"`        | there is a problem handling processing a route 处理处理路由时出现问题 | `route`: [buffalo#RouteInfo](https://pkg.go.dev/github.com/gobuffalo/buffalo#RouteInfo) `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) `context`: [buffalo#Context](https://pkg.go.dev/github.com/gobuffalo/buffalo#Context) |
| `buffalo.EvtWorkerStart`    | `"buffalo:worker:start"`     | [buffalo#App.Serve](https://pkg.go.dev/github.com/gobuffalo/buffalo#App.Serve) is called and workers are started buffalo#App.Serve 被调用，工作进程已启动 | `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) |
| `buffalo.EvtWorkerStartErr` | `"buffalo:worker:start:err"` | an error occurs when starting workers 启动工作进程时出错     | `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) |
| `buffalo.EvtWorkerStop`     | `"buffalo:worker:stop"`      | [buffalo#App.Stop](https://pkg.go.dev/github.com/gobuffalo/buffalo#App.Stop) is called and workers are stopped buffalo#App.Stop 被调用，工作进程已停止 | `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) |
| `buffalo.EvtWorkerStopErr`  | `"buffalo:worker:stop:err"`  | an error occurs when stopping workers 停止工作进程时出错     | `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) |
| `buffalo.EvtFailureErr`     | `"buffalo:failure:err"`      | something can’t be processed at all. it is a bad thing 根本无法处理某些内容。这是一件坏事 | `app`: [buffalo#App](https://pkg.go.dev/*github.com/gobuffalo/buffalo#App) `context`: [buffalo#Context](https://pkg.go.dev/github.com/gobuffalo/buffalo#Context) `app` ：buffalo#App `context` ：buffalo#Context |

### Buffalo Dev Events Buffalo Dev 事件 

The following events are known to be emitted by the `buffalo dev` during the development lifecyle.

​	已知以下事件在开发生命周期期间由 `buffalo dev` 发出。

| String 字符串                  | Emitted When 发出时间                           | Payload 有效负载                                             |
| ------------------------------ | ----------------------------------------------- | ------------------------------------------------------------ |
| `"buffalo:dev:raw"`            | an applicable file is modified 修改了适用的文件 | `event`: [fsnotify#Event](https://pkg.go.dev/github.com/fsnotify/fsnotify#Event) |
| `"buffalo:dev:build:started"`  | a build has started 构建已启动                  | `event`: [fsnotify#Event](https://pkg.go.dev/github.com/fsnotify/fsnotify#Event) `cmd`: string of the `go build` command (example: `"go build foo"`) `event` : fsnotify#Event `cmd` : `go build` 命令的字符串（示例： `"go build foo"` ） |
| `"buffalo:dev:build:finished"` | a build has completed 构建已完成                | `event`: [fsnotify#Event](https://pkg.go.dev/github.com/fsnotify/fsnotify#Event) `pid`: PID of the newly running binary `build_time`: the duration of the build `event` ：fsnotify#Event `pid` ：新运行的二进制文件的 PID `build_time` ：构建持续时间 |
| `"buffalo:dev:build:err"`      | a build error has occurred 发生构建错误         | `event`: [fsnotify#Event](https://pkg.go.dev/github.com/fsnotify/fsnotify#Event) `cmd`: string of the `go build` command (example: `"go build foo"`) `event` ：fsnotify#Event `cmd` ： `go build` 命令的字符串（示例： `"go build foo"` ） |