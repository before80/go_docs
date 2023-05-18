+++
title = "expvar"
linkTitle = "expvar"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# expvar

https://pkg.go.dev/expvar@go1.20.1



Package expvar provides a standardized interface to public variables, such as operation counters in servers. It exposes these variables via HTTP at /debug/vars in JSON format.

包expvar为公共变量提供了一个标准化的接口，例如服务器中的操作计数器。它通过HTTP在/debug/vars上以JSON格式公开这些变量。

Operations to set or modify these public variables are atomic.

设置或修改这些公共变量的操作是原子性的。

In addition to adding the HTTP handler, this package registers the following variables:

除了添加HTTP处理程序外，这个包还注册了以下变量：

```
cmdline   os.Args
memstats  runtime.Memstats
```

The package is sometimes only imported for the side effect of registering its HTTP handler and the above variables. To use it this way, link this package into your program:

这个包有时被导入只是为了注册其HTTP处理程序和上述变量的副作用。要以这种方式使用它，请将这个包链接到你的程序中：

```
import _ "expvar"
```



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [Do](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=329) 

``` go 
func Do(f func(KeyValue))
```

Do calls f for each exported variable. The global variable map is locked during the iteration, but existing entries may be concurrently updated.

Do为每个导出的变量调用f。在迭代过程中，全局变量图被锁定，但现有的条目可以被同时更新。

#### func [Handler](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=355)  <- go1.8

``` go 
func Handler() http.Handler
```

Handler returns the expvar HTTP Handler.

Handler返回expvar HTTP Handler。

This is only needed to install the handler in a non-standard location.

只有在将处理程序安装在非标准位置时才需要这样做。

#### func [Publish](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=282) 

``` go 
func Publish(name string, v Var)
```

Publish declares a named exported variable. This should be called from a package's init function when it creates its Vars. If the name is already registered then this will log.Panic.

Publish 声明了一个命名的导出变量。这应该从包的init函数中调用，当它创建其Vars时。如果名字已经被注册了，那么这将导致log.Panic。

## 类型

### type [Float](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=69) 

``` go 
type Float struct {
	// contains filtered or unexported fields
}
```

Float is a 64-bit float variable that satisfies the Var interface.

Float是一个满足Var接口的64位浮点数变量。

#### func [NewFloat](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=308) 

``` go 
func NewFloat(name string) *Float
```

#### (*Float) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=83) 

``` go 
func (v *Float) Add(delta float64)
```

Add adds delta to v.

Add将delta添加到v中。

#### (*Float) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=96) 

``` go 
func (v *Float) Set(value float64)
```

Set sets v to value.

Set 将v设置为值。

#### (*Float) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=77) 

``` go 
func (v *Float) String() string
```

#### (*Float) [Value](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=73)  <- go1.8

``` go 
func (v *Float) Value() float64
```

### type [Func](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=261) 

``` go 
type Func func() any
```

Func implements Var by calling the function and formatting the returned value using JSON.

Func通过调用函数和使用JSON格式化返回值来实现Var。

#### (Func) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=267) 

``` go 
func (f Func) String() string
```

#### (Func) [Value](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=263)  <- go1.8

``` go 
func (f Func) Value() any
```

### type [Int](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=48) 

``` go 
type Int struct {
	// contains filtered or unexported fields
}
```

Int is a 64-bit integer variable that satisfies the Var interface.

Int是一个满足Var接口的64位整数变量。

#### func [NewInt](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=302) 

``` go 
func NewInt(name string) *Int
```

#### (*Int) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=60) 

``` go 
func (v *Int) Add(delta int64)
```

#### (*Int) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=64) 

``` go 
func (v *Int) Set(value int64)
```

#### (*Int) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=56) 

``` go 
func (v *Int) String() string
```

#### (*Int) [Value](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=52)  <- go1.8

``` go 
func (v *Int) Value() int64
```

### type [KeyValue](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=108) 

``` go 
type KeyValue struct {
	Key   string
	Value Var
}
```

KeyValue represents a single entry in a Map.

KeyValue代表了一个Map中的一个条目。

### type [Map](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=101) 

``` go 
type Map struct {
	// contains filtered or unexported fields
}
```

Map is a string-to-Var map variable that satisfies the Var interface.

Map是一个满足Var接口的字符串到Var的map变量。

#### func [NewMap](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=314) 

``` go 
func NewMap(name string) *Map
```

#### (*Map) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=180) 

``` go 
func (v *Map) Add(key string, delta int64)
```

Add adds delta to the *Int value stored under the given map key.

Add将delta添加到存储在给定map键下的*Int值中。

#### (*Map) [AddFloat](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=197) 

``` go 
func (v *Map) AddFloat(key string, delta float64)
```

AddFloat adds delta to the *Float value stored under the given map key.

AddFloat将delta添加到存储在给定map键下的*Float值。

#### (*Map) [Delete](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=214)  <- go1.12

``` go 
func (v *Map) Delete(key string)
```

Delete deletes the given key from the map.

Delete 将给定的键从map上删除。

#### (*Map) [Do](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=227) 

``` go 
func (v *Map) Do(f func(KeyValue))
```

Do calls f for each entry in the map. The map is locked during the iteration, but existing entries may be concurrently updated.

Do为map中的每个条目调用f。在迭代过程中，map被锁定，但现有的条目可能被同时更新。

#### (*Map) [Get](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=159) 

``` go 
func (v *Map) Get(key string) Var
```

#### (*Map) [Init](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=134) 

``` go 
func (v *Map) Init() *Map
```

Init removes all keys from the map.

Init会从map上删除所有的键。

#### (*Map) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=165) 

``` go 
func (v *Map) Set(key string, av Var)
```

#### (*Map) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=113) 

``` go 
func (v *Map) String() string
```

### type [String](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=238) 

``` go 
type String struct {
	// contains filtered or unexported fields
}
```

String is a string variable, and satisfies the Var interface.

String是一个字符串变量，并且满足Var接口。

#### func [NewString](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=320) 

``` go 
func NewString(name string) *String
```

#### (*String) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=255) 

``` go 
func (v *String) Set(value string)
```

#### (*String) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=249) 

``` go 
func (v *String) String() string
```

String implements the Var interface. To get the unquoted string use Value.

String实现了Var接口。要获得未引用的字符串请使用Value。

#### (*String) [Value](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=242)  <- go1.8

``` go 
func (v *String) Value() string
```

### type [Var](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=40) 

``` go 
type Var interface {
	// String returns a valid JSON value for the variable.
	// Types with String methods that do not return valid JSON
	// (such as time.Time) must not be used as a Var.
    // String为变量返回一个有效的JSON值。
	// 具有String方法的类型如果不返回有效的JSON值(如time.Time)，则不能作为Var使用。
	String() string
}
```

Var is an abstract type for all exported variables.

Var是一个抽象类型，用于所有导出的变量。

#### func [Get](https://cs.opensource.google/go/go/+/go1.20.1:src/expvar/expvar.go;l=294) 

``` go 
func Get(name string) Var
```

Get retrieves a named exported variable. It returns nil if the name has not been registered.

Get 检索一个命名的导出变量。如果这个名字没有被注册，它将返回nil。