+++
title = "g"
date = 2024-03-21T17:51:29+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/frame/g](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/frame/g)

Package g provides commonly used type/function defines and coupled calling for creating commonly used objects.

​	软件包 g 提供了常用的类型/函数定义和耦合调用，用于创建常用对象。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Cfg

```go
func Cfg(name ...string) *gcfg.Config
```

Cfg is alias of Config. See Config.

​	Cfg 是 Config 的别名。请参阅配置。

#### func Client

```go
func Client() *gclient.Client
```

Client is a convenience function, which creates and returns a new HTTP client.

​	客户端是一个方便的函数，它创建并返回一个新的 HTTP 客户端。

#### func Config

```go
func Config(name ...string) *gcfg.Config
```

Config returns an instance of config object with specified name.

​	Config 返回具有指定名称的 config 对象的实例。

#### func DB

```go
func DB(name ...string) gdb.DB
```

DB returns an instance of database ORM object with specified configuration group name.

​	数据库返回具有指定配置组名称的数据库 ORM 对象的实例。

#### func Dump

```go
func Dump(values ...interface{})
```

Dump dumps a variable to stdout with more manually readable.

​	Dump 将变量转储到 stdout，具有更多的手动可读性。

#### func DumpJson <-2.4.2

```go
func DumpJson(jsonContent string)
```

DumpJson pretty dumps json content to stdout.

​	DumpJson 漂亮地将 json 内容转储到 stdout。

#### func DumpTo

```go
func DumpTo(writer io.Writer, value interface{}, option gutil.DumpOption)
```

DumpTo writes variables `values` as a string in to `writer` with more manually readable

​	DumpTo 将 `values` 变量作为字符串写入 to `writer` ，具有更多的手动可读性

#### func DumpWithOption

```go
func DumpWithOption(value interface{}, option gutil.DumpOption)
```

DumpWithOption returns variables `values` as a string with more manually readable.

​	DumpWithOption `values` 以字符串形式返回变量，具有更多的手动可读性。

#### func DumpWithType

```go
func DumpWithType(values ...interface{})
```

DumpWithType acts like Dump, but with type information. Also see Dump.

​	DumpWithType 的行为类似于 Dump，但具有类型信息。另请参阅转储。

#### func Go <-2.5.3

```go
func Go(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
	recoverFunc func(ctx context.Context, exception error),
)
```

Go creates a new asynchronous goroutine function with specified recover function.

​	Go 使用指定的 recover 函数创建一个新的异步 goroutine 函数。

The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`. If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will throw to parent goroutine.

​	当在执行 期间 `goroutineFunc` 出现任何恐慌时， `recoverFunc` 将调用该参数。如果 `recoverFunc` 给定 nil，它会忽略来自 `goroutineFunc` 的恐慌，并且不会向父 goroutine 抛出恐慌。

But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.

​	但是，请注意，如果 `recoverFunc` 也引发恐慌，这种恐慌将被抛给父 goroutine。

#### func I18n

```go
func I18n(name ...string) *gi18n.Manager
```

I18n returns an instance of gi18n.Manager. The parameter `name` is the name for the instance.

​	I18n 返回 gi18n 的实例。经理。该参数 `name` 是实例的名称。

#### func IsEmpty

```go
func IsEmpty(value interface{}, traceSource ...bool) bool
```

IsEmpty checks whether given `value` empty. It returns true if `value` is in: 0, nil, false, “”, len(slice/map/chan) == 0. Or else it returns true.

​	IsEmpty 检查是否给定 `value` 为空。如果 `value` in 为 0， nil， false， “”， len（slice/map/chan） == 0，则返回 true。否则，它将返回 true。

The parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer that also points to a pointer. It returns true if the source is empty when `traceSource` is true. Note that it might use reflect feature which affects performance a little.

​	如果给定 `value` 的指针类型也指向指针，则该参数 `traceSource` 用于跟踪源变量。如果源为 true， `traceSource` 则返回 true。请注意，它可能会使用稍微影响性能的反射功能。

#### func IsNil

```go
func IsNil(value interface{}, traceSource ...bool) bool
```

IsNil checks whether given `value` is nil. Parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer that also points to a pointer. It returns nil if the source is nil when `traceSource` is true. Note that it might use reflect feature which affects performance a little.

​	IsNil 检查给定 `value` 是否为 nil。参数 `traceSource` 用于跟踪源变量，如果给定 `value` 的是也指向指针的指针类型。如果源为 nil，则返回 nil，而 `traceSource` true 为 true。请注意，它可能会使用稍微影响性能的反射功能。

#### func Listen

```go
func Listen()
```

Listen is an alias of gproc.Listen, which handles the signals received and automatically calls registered signal handler functions. It blocks until shutdown signals received and all registered shutdown handlers done.

​	Listen 是 gproc 的别名。侦听，处理接收到的信号并自动调用已注册的信号处理程序函数。它会阻塞，直到接收到关机信号并完成所有注册的关机处理程序。

#### func Log

```go
func Log(name ...string) *glog.Logger
```

Log returns an instance of glog.Logger. The parameter `name` is the name for the instance.

​	log 返回 glog 的实例。记录。该参数 `name` 是实例的名称。

#### func Model

```go
func Model(tableNameOrStruct ...interface{}) *gdb.Model
```

Model creates and returns a model based on configuration of default database group.

​	模型根据默认数据库组的配置创建并返回模型。

#### func ModelRaw

```go
func ModelRaw(rawSql string, args ...interface{}) *gdb.Model
```

ModelRaw creates and returns a model based on a raw sql not a table.

​	ModelRaw 创建并返回基于原始 sql 而不是表的模型。

#### func Redis

```go
func Redis(name ...string) *gredis.Redis
```

Redis returns an instance of redis client with specified configuration group name.

​	Redis 返回具有指定配置组名称的 redis 客户端实例。

#### func RequestFromCtx

```go
func RequestFromCtx(ctx context.Context) *ghttp.Request
```

RequestFromCtx retrieves and returns the Request object from context.

​	RequestFromCtx 从上下文中检索并返回 Request 对象。

#### func Res

```go
func Res(name ...string) *gres.Resource
```

Res is alias of Resource. See Resource.

​	Res 是 Resource 的别名。请参阅资源。

#### func Resource

```go
func Resource(name ...string) *gres.Resource
```

Resource returns an instance of Resource. The parameter `name` is the name for the instance.

​	Resource 返回 Resource 的实例。该参数 `name` 是实例的名称。

#### func Server

```go
func Server(name ...interface{}) *ghttp.Server
```

Server returns an instance of http server with specified name.

​	Server 返回具有指定名称的 http 服务器实例。

##### Example

``` go
```

#### func SetDebug

```go
func SetDebug(enabled bool)
```

SetDebug enables/disables the GoFrame internal logging manually. Note that this function is not concurrent safe, be aware of the DATA RACE, which means you should call this function in your boot but not the runtime.

​	SetDebug 手动启用/禁用 GoFrame 内部日志记录。请注意，此函数不是并发安全的，请注意 DATA RACE，这意味着您应该在启动中调用此函数，而不是在运行时调用此函数。

#### func TCPServer

```go
func TCPServer(name ...interface{}) *gtcp.Server
```

TCPServer returns an instance of tcp server with specified name.

​	TCPServer 返回具有指定名称的 tcp 服务器实例。

#### func Throw

```go
func Throw(exception interface{})
```

Throw throws an exception, which can be caught by TryCatch function.

​	Throw 抛出异常，TryCatch 函数可以捕获该异常。

#### func Try

```go
func Try(ctx context.Context, try func(ctx context.Context)) (err error)
```

Try implements try… logistics using internal panic…recover. It returns error if any exception occurs, or else it returns nil.

​	尝试实现尝试...物流使用内部恐慌...恢复。如果发生任何异常，它将返回错误，否则返回 nil。

#### func TryCatch

```go
func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error))
```

TryCatch implements try…catch… logistics using internal panic…recover. It automatically calls function `catch` if any exception occurs and passes the exception as an error.

​	TryCatch 实现 try...抓住。。。物流使用内部恐慌...恢复。如果发生任何异常，它会自动调用函数 `catch` ，并将异常作为错误传递。

But, note that, if function `catch` also throws panic, the current goroutine will panic.

​	但是，请注意，如果函数 `catch` 也引发恐慌，则当前的 goroutine 将恐慌。

#### func UDPServer

```go
func UDPServer(name ...interface{}) *gudp.Server
```

UDPServer returns an instance of udp server with specified name.

​	UDPServer 返回具有指定名称的 udp 服务器实例。

#### func Validator

```go
func Validator() *gvalid.Validator
```

Validator is a convenience function, which creates and returns a new validation manager object.

​	Validator 是一个方便的函数，它创建并返回一个新的验证管理器对象。

#### func View

```go
func View(name ...string) *gview.View
```

View returns an instance of template engine object with specified name.

​	View 返回具有指定名称的模板引擎对象的实例。

#### func Wait

```go
func Wait()
```

Wait is an alias of ghttp.Wait, which blocks until all the web servers shutdown. It’s commonly used in multiple servers’ situation.

​	Wait 是 ghttp 的别名。等待，直到所有 Web 服务器关闭为止。它通常用于多个服务器的情况。

## 类型

### type Array

```go
type Array = []interface{} // Array is alias of frequently-used slice type []interface{}.
```

### type ArrayAny

```go
type ArrayAny = []interface{} // ArrayAny is alias of frequently-used slice type []interface{}.
```

### type ArrayInt

```go
type ArrayInt = []int // ArrayInt is alias of frequently-used slice type []int.
```

### type ArrayStr

```go
type ArrayStr = []string // ArrayStr is alias of frequently-used slice type []string.
```

### type Ctx

```go
type Ctx = context.Context // Ctx is alias of frequently-used type context.Context.
```

### type List

```go
type List = []Map // List is alias of frequently-used slice type []Map.
```

### type ListAnyAny

```go
type ListAnyAny = []MapAnyAny // ListAnyAny is alias of frequently-used slice type []MapAnyAny.
```

### type ListAnyBool

```go
type ListAnyBool = []MapAnyBool // ListAnyBool is alias of frequently-used slice type []MapAnyBool.
```

### type ListAnyInt

```go
type ListAnyInt = []MapAnyInt // ListAnyInt is alias of frequently-used slice type []MapAnyInt.
```

### type ListAnyStr

```go
type ListAnyStr = []MapAnyStr // ListAnyStr is alias of frequently-used slice type []MapAnyStr.
```

### type ListIntAny

```go
type ListIntAny = []MapIntAny // ListIntAny is alias of frequently-used slice type []MapIntAny.
```

### type ListIntBool

```go
type ListIntBool = []MapIntBool // ListIntBool is alias of frequently-used slice type []MapIntBool.
```

### type ListIntInt

```go
type ListIntInt = []MapIntInt // ListIntInt is alias of frequently-used slice type []MapIntInt.
```

### type ListIntStr

```go
type ListIntStr = []MapIntStr // ListIntStr is alias of frequently-used slice type []MapIntStr.
```

### type ListStrAny

```go
type ListStrAny = []MapStrAny // ListStrAny is alias of frequently-used slice type []MapStrAny.
```

### type ListStrBool

```go
type ListStrBool = []MapStrBool // ListStrBool is alias of frequently-used slice type []MapStrBool.
```

### type ListStrInt

```go
type ListStrInt = []MapStrInt // ListStrInt is alias of frequently-used slice type []MapStrInt.
```

### type ListStrStr

```go
type ListStrStr = []MapStrStr // ListStrStr is alias of frequently-used slice type []MapStrStr.
```

### type Map

```go
type Map = map[string]interface{} // Map is alias of frequently-used map type map[string]interface{}.
```

### type MapAnyAny

```go
type MapAnyAny = map[interface{}]interface{} // MapAnyAny is alias of frequently-used map type map[interface{}]interface{}.
```

### type MapAnyBool

```go
type MapAnyBool = map[interface{}]bool // MapAnyBool is alias of frequently-used map type map[interface{}]bool.
```

### type MapAnyInt

```go
type MapAnyInt = map[interface{}]int // MapAnyInt is alias of frequently-used map type map[interface{}]int.
```

### type MapAnyStr

```go
type MapAnyStr = map[interface{}]string // MapAnyStr is alias of frequently-used map type map[interface{}]string.
```

### type MapIntAny

```go
type MapIntAny = map[int]interface{} // MapIntAny is alias of frequently-used map type map[int]interface{}.
```

### type MapIntBool

```go
type MapIntBool = map[int]bool // MapIntBool is alias of frequently-used map type map[int]bool.
```

### type MapIntInt

```go
type MapIntInt = map[int]int // MapIntInt is alias of frequently-used map type map[int]int.
```

### type MapIntStr

```go
type MapIntStr = map[int]string // MapIntStr is alias of frequently-used map type map[int]string.
```

### type MapStrAny

```go
type MapStrAny = map[string]interface{} // MapStrAny is alias of frequently-used map type map[string]interface{}.
```

### type MapStrBool

```go
type MapStrBool = map[string]bool // MapStrBool is alias of frequently-used map type map[string]bool.
```

### type MapStrInt

```go
type MapStrInt = map[string]int // MapStrInt is alias of frequently-used map type map[string]int.
```

### type MapStrStr

```go
type MapStrStr = map[string]string // MapStrStr is alias of frequently-used map type map[string]string.
```

### type Meta

```go
type Meta = gmeta.Meta // Meta is alias of frequently-used type gmeta.Meta.
```

### type Slice

```go
type Slice = []interface{} // Slice is alias of frequently-used slice type []interface{}.
```

### type SliceAny

```go
type SliceAny = []interface{} // SliceAny is alias of frequently-used slice type []interface{}.
```

### type SliceInt

```go
type SliceInt = []int // SliceInt is alias of frequently-used slice type []int.
```

### type SliceStr

```go
type SliceStr = []string // SliceStr is alias of frequently-used slice type []string.
```

### type Var

```go
type Var = gvar.Var // Var is a universal variable interface, like generics.
```

#### func NewVar

```go
func NewVar(i interface{}, safe ...bool) *Var
```

NewVar returns a gvar.Var.

​	NewVar 返回 gvar.Var。