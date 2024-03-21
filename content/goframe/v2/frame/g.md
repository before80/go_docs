+++
title = "g"
date = 2024-03-21T17:51:29+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/frame/g

Package g provides commonly used type/function defines and coupled calling for creating commonly used objects.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Cfg 

``` go
func Cfg(name ...string) *gcfg.Config
```

Cfg is alias of Config. See Config.

##### func Client 

``` go
func Client() *gclient.Client
```

Client is a convenience function, which creates and returns a new HTTP client.

##### func Config 

``` go
func Config(name ...string) *gcfg.Config
```

Config returns an instance of config object with specified name.

##### func DB 

``` go
func DB(name ...string) gdb.DB
```

DB returns an instance of database ORM object with specified configuration group name.

##### func Dump 

``` go
func Dump(values ...interface{})
```

Dump dumps a variable to stdout with more manually readable.

##### func DumpJson <-2.4.2

``` go
func DumpJson(jsonContent string)
```

DumpJson pretty dumps json content to stdout.

##### func DumpTo 

``` go
func DumpTo(writer io.Writer, value interface{}, option gutil.DumpOption)
```

DumpTo writes variables `values` as a string in to `writer` with more manually readable

##### func DumpWithOption 

``` go
func DumpWithOption(value interface{}, option gutil.DumpOption)
```

DumpWithOption returns variables `values` as a string with more manually readable.

##### func DumpWithType 

``` go
func DumpWithType(values ...interface{})
```

DumpWithType acts like Dump, but with type information. Also see Dump.

##### func Go <-2.5.3

``` go
func Go(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
	recoverFunc func(ctx context.Context, exception error),
)
```

Go creates a new asynchronous goroutine function with specified recover function.

The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`. If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will throw to parent goroutine.

But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.

##### func I18n 

``` go
func I18n(name ...string) *gi18n.Manager
```

I18n returns an instance of gi18n.Manager. The parameter `name` is the name for the instance.

##### func IsEmpty 

``` go
func IsEmpty(value interface{}, traceSource ...bool) bool
```

IsEmpty checks whether given `value` empty. It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0. Or else it returns true.

The parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer that also points to a pointer. It returns true if the source is empty when `traceSource` is true. Note that it might use reflect feature which affects performance a little.

##### func IsNil 

``` go
func IsNil(value interface{}, traceSource ...bool) bool
```

IsNil checks whether given `value` is nil. Parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer that also points to a pointer. It returns nil if the source is nil when `traceSource` is true. Note that it might use reflect feature which affects performance a little.

##### func Listen 

``` go
func Listen()
```

Listen is an alias of gproc.Listen, which handles the signals received and automatically calls registered signal handler functions. It blocks until shutdown signals received and all registered shutdown handlers done.

##### func Log 

``` go
func Log(name ...string) *glog.Logger
```

Log returns an instance of glog.Logger. The parameter `name` is the name for the instance.

##### func Model 

``` go
func Model(tableNameOrStruct ...interface{}) *gdb.Model
```

Model creates and returns a model based on configuration of default database group.

##### func ModelRaw 

``` go
func ModelRaw(rawSql string, args ...interface{}) *gdb.Model
```

ModelRaw creates and returns a model based on a raw sql not a table.

##### func Redis 

``` go
func Redis(name ...string) *gredis.Redis
```

Redis returns an instance of redis client with specified configuration group name.

##### func RequestFromCtx 

``` go
func RequestFromCtx(ctx context.Context) *ghttp.Request
```

RequestFromCtx retrieves and returns the Request object from context.

##### func Res 

``` go
func Res(name ...string) *gres.Resource
```

Res is alias of Resource. See Resource.

##### func Resource 

``` go
func Resource(name ...string) *gres.Resource
```

Resource returns an instance of Resource. The parameter `name` is the name for the instance.

##### func Server 

``` go
func Server(name ...interface{}) *ghttp.Server
```

Server returns an instance of http server with specified name.

##### Example

``` go
```
##### func SetDebug 

``` go
func SetDebug(enabled bool)
```

SetDebug enables/disables the GoFrame internal logging manually. Note that this function is not concurrent safe, be aware of the DATA RACE, which means you should call this function in your boot but not the runtime.

##### func TCPServer 

``` go
func TCPServer(name ...interface{}) *gtcp.Server
```

TCPServer returns an instance of tcp server with specified name.

##### func Throw 

``` go
func Throw(exception interface{})
```

Throw throws an exception, which can be caught by TryCatch function.

##### func Try 

``` go
func Try(ctx context.Context, try func(ctx context.Context)) (err error)
```

Try implements try... logistics using internal panic...recover. It returns error if any exception occurs, or else it returns nil.

##### func TryCatch 

``` go
func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error))
```

TryCatch implements try...catch... logistics using internal panic...recover. It automatically calls function `catch` if any exception occurs and passes the exception as an error.

But, note that, if function `catch` also throws panic, the current goroutine will panic.

##### func UDPServer 

``` go
func UDPServer(name ...interface{}) *gudp.Server
```

UDPServer returns an instance of udp server with specified name.

##### func Validator 

``` go
func Validator() *gvalid.Validator
```

Validator is a convenience function, which creates and returns a new validation manager object.

##### func View 

``` go
func View(name ...string) *gview.View
```

View returns an instance of template engine object with specified name.

##### func Wait 

``` go
func Wait()
```

Wait is an alias of ghttp.Wait, which blocks until all the web servers shutdown. It's commonly used in multiple servers' situation.

### Types 

#### type Array 

``` go
type Array = []interface{} // Array is alias of frequently-used slice type []interface{}.
```

#### type ArrayAny 

``` go
type ArrayAny = []interface{} // ArrayAny is alias of frequently-used slice type []interface{}.
```

#### type ArrayInt 

``` go
type ArrayInt = []int // ArrayInt is alias of frequently-used slice type []int.
```

#### type ArrayStr 

``` go
type ArrayStr = []string // ArrayStr is alias of frequently-used slice type []string.
```

#### type Ctx 

``` go
type Ctx = context.Context // Ctx is alias of frequently-used type context.Context.
```

#### type List 

``` go
type List = []Map // List is alias of frequently-used slice type []Map.
```

#### type ListAnyAny 

``` go
type ListAnyAny = []MapAnyAny // ListAnyAny is alias of frequently-used slice type []MapAnyAny.
```

#### type ListAnyBool 

``` go
type ListAnyBool = []MapAnyBool // ListAnyBool is alias of frequently-used slice type []MapAnyBool.
```

#### type ListAnyInt 

``` go
type ListAnyInt = []MapAnyInt // ListAnyInt is alias of frequently-used slice type []MapAnyInt.
```

#### type ListAnyStr 

``` go
type ListAnyStr = []MapAnyStr // ListAnyStr is alias of frequently-used slice type []MapAnyStr.
```

#### type ListIntAny 

``` go
type ListIntAny = []MapIntAny // ListIntAny is alias of frequently-used slice type []MapIntAny.
```

#### type ListIntBool 

``` go
type ListIntBool = []MapIntBool // ListIntBool is alias of frequently-used slice type []MapIntBool.
```

#### type ListIntInt 

``` go
type ListIntInt = []MapIntInt // ListIntInt is alias of frequently-used slice type []MapIntInt.
```

#### type ListIntStr 

``` go
type ListIntStr = []MapIntStr // ListIntStr is alias of frequently-used slice type []MapIntStr.
```

#### type ListStrAny 

``` go
type ListStrAny = []MapStrAny // ListStrAny is alias of frequently-used slice type []MapStrAny.
```

#### type ListStrBool 

``` go
type ListStrBool = []MapStrBool // ListStrBool is alias of frequently-used slice type []MapStrBool.
```

#### type ListStrInt 

``` go
type ListStrInt = []MapStrInt // ListStrInt is alias of frequently-used slice type []MapStrInt.
```

#### type ListStrStr 

``` go
type ListStrStr = []MapStrStr // ListStrStr is alias of frequently-used slice type []MapStrStr.
```

#### type Map 

``` go
type Map = map[string]interface{} // Map is alias of frequently-used map type map[string]interface{}.
```

#### type MapAnyAny 

``` go
type MapAnyAny = map[interface{}]interface{} // MapAnyAny is alias of frequently-used map type map[interface{}]interface{}.
```

#### type MapAnyBool 

``` go
type MapAnyBool = map[interface{}]bool // MapAnyBool is alias of frequently-used map type map[interface{}]bool.
```

#### type MapAnyInt 

``` go
type MapAnyInt = map[interface{}]int // MapAnyInt is alias of frequently-used map type map[interface{}]int.
```

#### type MapAnyStr 

``` go
type MapAnyStr = map[interface{}]string // MapAnyStr is alias of frequently-used map type map[interface{}]string.
```

#### type MapIntAny 

``` go
type MapIntAny = map[int]interface{} // MapIntAny is alias of frequently-used map type map[int]interface{}.
```

#### type MapIntBool 

``` go
type MapIntBool = map[int]bool // MapIntBool is alias of frequently-used map type map[int]bool.
```

#### type MapIntInt 

``` go
type MapIntInt = map[int]int // MapIntInt is alias of frequently-used map type map[int]int.
```

#### type MapIntStr 

``` go
type MapIntStr = map[int]string // MapIntStr is alias of frequently-used map type map[int]string.
```

#### type MapStrAny 

``` go
type MapStrAny = map[string]interface{} // MapStrAny is alias of frequently-used map type map[string]interface{}.
```

#### type MapStrBool 

``` go
type MapStrBool = map[string]bool // MapStrBool is alias of frequently-used map type map[string]bool.
```

#### type MapStrInt 

``` go
type MapStrInt = map[string]int // MapStrInt is alias of frequently-used map type map[string]int.
```

#### type MapStrStr 

``` go
type MapStrStr = map[string]string // MapStrStr is alias of frequently-used map type map[string]string.
```

#### type Meta 

``` go
type Meta = gmeta.Meta // Meta is alias of frequently-used type gmeta.Meta.
```

#### type Slice 

``` go
type Slice = []interface{} // Slice is alias of frequently-used slice type []interface{}.
```

#### type SliceAny 

``` go
type SliceAny = []interface{} // SliceAny is alias of frequently-used slice type []interface{}.
```

#### type SliceInt 

``` go
type SliceInt = []int // SliceInt is alias of frequently-used slice type []int.
```

#### type SliceStr 

``` go
type SliceStr = []string // SliceStr is alias of frequently-used slice type []string.
```

#### type Var 

``` go
type Var = gvar.Var // Var is a universal variable interface, like generics.
```

##### func NewVar 

``` go
func NewVar(i interface{}, safe ...bool) *Var
```

NewVar returns a gvar.Var.