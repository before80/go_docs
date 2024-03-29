+++
title = "gtrace"
date = 2024-03-21T17:53:48+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gtrace](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gtrace)

Package gtrace provides convenience wrapping functionality for tracing feature using OpenTelemetry.

​	包 gtrace 为使用 OpenTelemetry 的跟踪功能提供了便利的包装功能。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func CheckSetDefaultTextMapPropagator

```go
func CheckSetDefaultTextMapPropagator()
```

CheckSetDefaultTextMapPropagator sets the default TextMapPropagator if it is not set previously.

​	CheckSetDefaultTextMapPropagator 设置默认的 TextMapPropagator（如果之前未设置）。

#### func CommonLabels

```go
func CommonLabels() []attribute.KeyValue
```

CommonLabels returns common used attribute labels: ip.intranet, hostname.

​	CommonLabels 返回常用的属性标签：ip.intranet、hostname。

#### func GetBaggageMap

```go
func GetBaggageMap(ctx context.Context) *gmap.StrAnyMap
```

GetBaggageMap retrieves and returns the baggage values as map.

​	GetBaggageMap 检索行李值并将其作为 map 返回。

#### func GetBaggageVar

```go
func GetBaggageVar(ctx context.Context, key string) *gvar.Var
```

GetBaggageVar retrieves value and returns a *gvar.Var for specified key from baggage.

​	GetBaggageVar 检索值并返回 *gvar。Var 表示行李中的指定密钥。

#### func GetDefaultTextMapPropagator

```go
func GetDefaultTextMapPropagator() propagation.TextMapPropagator
```

GetDefaultTextMapPropagator returns the default propagator for context propagation between peers.

​	GetDefaultTextMapPropagator 返回用于在对等方之间进行上下文传播的默认传播器。

#### func GetSpanID

```go
func GetSpanID(ctx context.Context) string
```

GetSpanID retrieves and returns SpanId from context. It returns an empty string is tracing feature is not activated.

​	GetSpanID 从上下文中检索并返回 SpanId。它返回一个空字符串，即跟踪功能未激活。

#### func GetTraceID

```go
func GetTraceID(ctx context.Context) string
```

GetTraceID retrieves and returns TraceId from context. It returns an empty string is tracing feature is not activated.

​	GetTraceID 从上下文中检索并返回 TraceId。它返回一个空字符串，即跟踪功能未激活。

#### func IsTracingInternal

```go
func IsTracingInternal() bool
```

IsTracingInternal returns whether tracing spans of internal components.

​	IsTracingInternal 返回是否跟踪内部组件的跨度。

#### func IsUsingDefaultProvider

```go
func IsUsingDefaultProvider() bool
```

IsUsingDefaultProvider checks and return if currently using default trace provider.

​	IsUsingDefaultProvider 检查并返回当前是否使用默认跟踪提供程序。

#### func MaxContentLogSize

```go
func MaxContentLogSize() int
```

MaxContentLogSize returns the max log size for request and response body, especially for HTTP/RPC request.

​	MaxContentLogSize 返回请求和响应正文的最大日志大小，尤其是 HTTP/RPC 请求。

#### func SetBaggageMap

```go
func SetBaggageMap(ctx context.Context, data map[string]interface{}) context.Context
```

SetBaggageMap is a convenient function for adding map key-value pairs to baggage. Note that it uses attribute.Any to set the key-value pair.

​	SetBaggageMap 是一个方便的函数，用于将地图键值对添加到行李中。请注意，它使用属性。any 设置键值对。

#### func SetBaggageValue

```go
func SetBaggageValue(ctx context.Context, key string, value interface{}) context.Context
```

SetBaggageValue is a convenient function for adding one key-value pair to baggage. Note that it uses attribute.Any to set the key-value pair.

​	SetBaggageValue 是一个方便的函数，用于将一个键值对添加到行李中。请注意，它使用属性。any 设置键值对。

#### func WithTraceID

```go
func WithTraceID(ctx context.Context, traceID string) (context.Context, error)
```

WithTraceID injects custom trace id into context to propagate.

​	WithTraceID 将自定义跟踪 ID 注入到上下文中进行传播。

#### func WithUUID <-2.2.0

```go
func WithUUID(ctx context.Context, uuid string) (context.Context, error)
```

WithUUID injects custom trace id with UUID into context to propagate.

​	WithUUID 将带有 UUID 的自定义跟踪 ID 注入到上下文中以进行传播。

## 类型

### type Baggage

```go
type Baggage struct {
	// contains filtered or unexported fields
}
```

Baggage holds the data through all tracing spans.

​	行李在所有跟踪跨度中保存数据。

#### func NewBaggage

```go
func NewBaggage(ctx context.Context) *Baggage
```

NewBaggage creates and returns a new Baggage object from given tracing context.

​	NewBaggage 从给定的跟踪上下文创建并返回一个新的 Baggage 对象。

#### (*Baggage) Ctx

```go
func (b *Baggage) Ctx() context.Context
```

Ctx returns the context that Baggage holds.

​	Ctx 返回 Baggage 所包含的上下文。

#### (*Baggage) GetMap

```go
func (b *Baggage) GetMap() *gmap.StrAnyMap
```

GetMap retrieves and returns the baggage values as map.

​	GetMap 检索行李值并将其作为地图返回。

#### (*Baggage) GetVar

```go
func (b *Baggage) GetVar(key string) *gvar.Var
```

GetVar retrieves value and returns a *gvar.Var for specified key from baggage.

​	GetVar 检索值并返回 *gvar。Var 表示行李中的指定密钥。

#### (*Baggage) SetMap

```go
func (b *Baggage) SetMap(data map[string]interface{}) context.Context
```

SetMap is a convenient function for adding map key-value pairs to baggage. Note that it uses attribute.Any to set the key-value pair.

​	SetMap 是一个方便的功能，用于将地图键值对添加到行李中。请注意，它使用属性。any 设置键值对。

#### (*Baggage) SetValue

```go
func (b *Baggage) SetValue(key string, value interface{}) context.Context
```

SetValue is a convenient function for adding one key-value pair to baggage. Note that it uses attribute.Any to set the key-value pair.

​	SetValue 是一个方便的函数，用于将一个键值对添加到行李中。请注意，它使用属性。any 设置键值对。

### type Carrier

```go
type Carrier map[string]interface{}
```

Carrier is the storage medium used by a TextMapPropagator.

​	Carrier 是 TextMapPropagator 使用的存储介质。

#### func NewCarrier

```go
func NewCarrier(data ...map[string]interface{}) Carrier
```

NewCarrier creates and returns a Carrier.

​	NewCarrier 创建并返回一个 Carrier。

#### (Carrier) Get

```go
func (c Carrier) Get(k string) string
```

Get returns the value associated with the passed key.

​	Get 返回与传递的密钥关联的值。

#### (Carrier) Keys

```go
func (c Carrier) Keys() []string
```

Keys lists the keys stored in this carrier.

​	密钥列出了存储在此载体中的密钥。

#### (Carrier) MustMarshal

```go
func (c Carrier) MustMarshal() []byte
```

MustMarshal .returns the JSON encoding of c

​	MustMarshal .返回 c 的 JSON 编码

#### (Carrier) Set

```go
func (c Carrier) Set(k, v string)
```

Set stores the key-value pair.

​	Set 存储键值对。

#### (Carrier) String

```go
func (c Carrier) String() string
```

String converts and returns current Carrier as string.

​	String 转换并返回当前 Carrier 作为字符串。

#### (Carrier) UnmarshalJSON

```go
func (c Carrier) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements interface UnmarshalJSON for package json.

​	UnmarshalJSON 为包 json 实现接口 UnmarshalJSON。

### type Span

```go
type Span struct {
	trace.Span
}
```

Span warps trace.Span for compatibility and extension.

​	跨度翘曲轨迹。跨度用于兼容性和扩展。

#### func NewSpan

```go
func NewSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, *Span)
```

NewSpan creates a span using default tracer.

​	NewSpan 使用默认跟踪器创建跨度。

### type Tracer

```go
type Tracer struct {
	trace.Tracer
}
```

Tracer warps trace.Tracer for compatibility and extension.

​	示踪剂翘曲跟踪。用于兼容性和扩展的示踪剂。

#### func NewTracer

```go
func NewTracer(name ...string) *Tracer
```

NewTracer Tracer is a short function for retrieving Tracer.

​	NewTracer Tracer 是用于检索 Tracer 的简短函数。