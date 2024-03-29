+++
title = "gsel"
date = 2024-03-21T17:53:24+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gsel](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gsel)

Package gsel provides selector definition and implements.

​	软件包 gsel 提供选择器定义和实现。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func SetBuilder

```go
func SetBuilder(builder Builder)
```

SetBuilder sets the default builder for globally used purpose.

​	SetBuilder 为全局使用目的设置默认构建器。

## 类型

### type Builder

```go
type Builder interface {
	Name() string
	Build() Selector
}
```

Builder creates and returns selector in runtime.

​	Builder 在运行时创建并返回选择器。

#### func GetBuilder

```go
func GetBuilder() Builder
```

GetBuilder returns the default builder for globally used purpose.

​	GetBuilder 返回用于全局使用的默认构建器。

#### func NewBuilderLeastConnection

```go
func NewBuilderLeastConnection() Builder
```

#### func NewBuilderRandom

```go
func NewBuilderRandom() Builder
```

#### func NewBuilderRoundRobin

```go
func NewBuilderRoundRobin() Builder
```

#### func NewBuilderWeight

```go
func NewBuilderWeight() Builder
```

### type DoneFunc

```go
type DoneFunc func(ctx context.Context, di DoneInfo)
```

DoneFunc is callback function when RPC invoke done.

​	DoneFunc 是 RPC 调用完成时的回调函数。

### type DoneInfo

```go
type DoneInfo struct {
	// Err is the rpc error the RPC finished with. It could be nil.
	Err error

	// Trailer contains the metadata from the RPC's trailer, if present.
	Trailer DoneInfoMD

	// BytesSent indicates if any bytes have been sent to the server.
	BytesSent bool

	// BytesReceived indicates if any byte has been received from the server.
	BytesReceived bool

	// ServerLoad is the load received from server. It's usually sent as part of
	// trailing metadata.
	//
	// The only supported type now is *orca_v1.LoadReport.
	ServerLoad interface{}
}
```

DoneInfo contains additional information for done.

​	DoneInfo 包含有关完成的其他信息。

### type DoneInfoMD

```go
type DoneInfoMD interface {
	// Len returns the number of items in md.
	Len() int

	// Get obtains the values for a given key.
	//
	// k is converted to lowercase before searching in md.
	Get(k string) []string

	// Set sets the value of a given key with a slice of values.
	//
	// k is converted to lowercase before storing in md.
	Set(key string, values ...string)

	// Append adds the values to key k, not overwriting what was already stored at
	// that key.
	//
	// k is converted to lowercase before storing in md.
	Append(k string, values ...string)

	// Delete removes the values for a given key k which is converted to lowercase
	// before removing it from md.
	Delete(k string)
}
```

DoneInfoMD is a mapping from metadata keys to value array. Users should use the following two convenience functions New and Pairs to generate MD.

​	DoneInfoMD 是从元数据键到值数组的映射。用户应使用以下两个便捷函数 New 和 Pairs 来生成 MD。

### type Node

```go
type Node interface {
	Service() gsvc.Service
	Address() string
}
```

Node is node interface.

​	节点是节点接口。

### type Nodes <-2.1.0

```go
type Nodes []Node
```

Nodes contains multiple Node.

​	节点包含多个节点。

#### (Nodes) String

```go
func (ns Nodes) String() string
```

String formats and returns Nodes as string.

​	String 格式化并返回 Nodes 作为字符串。

### type Selector

```go
type Selector interface {
	// Pick selects and returns service.
	Pick(ctx context.Context) (node Node, done DoneFunc, err error)

	// Update updates services into Selector.
	Update(ctx context.Context, nodes Nodes) error
}
```

Selector for service balancer.

​	服务平衡器的选择器。

#### func NewSelectorLeastConnection

```go
func NewSelectorLeastConnection() Selector
```

#### func NewSelectorRandom

```go
func NewSelectorRandom() Selector
```

#### func NewSelectorRoundRobin

```go
func NewSelectorRoundRobin() Selector
```

#### func NewSelectorWeight

```go
func NewSelectorWeight() Selector
```