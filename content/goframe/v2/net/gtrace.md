+++
title = "gtrace"
date = 2024-03-21T17:53:48+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/net/gtrace

Package gtrace provides convenience wrapping functionality for tracing feature using OpenTelemetry.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func CheckSetDefaultTextMapPropagator 

``` go
func CheckSetDefaultTextMapPropagator()
```

CheckSetDefaultTextMapPropagator sets the default TextMapPropagator if it is not set previously.

##### func CommonLabels 

``` go
func CommonLabels() []attribute.KeyValue
```

CommonLabels returns common used attribute labels: ip.intranet, hostname.

##### func GetBaggageMap 

``` go
func GetBaggageMap(ctx context.Context) *gmap.StrAnyMap
```

GetBaggageMap retrieves and returns the baggage values as map.

##### func GetBaggageVar 

``` go
func GetBaggageVar(ctx context.Context, key string) *gvar.Var
```

GetBaggageVar retrieves value and returns a *gvar.Var for specified key from baggage.

##### func GetDefaultTextMapPropagator 

``` go
func GetDefaultTextMapPropagator() propagation.TextMapPropagator
```

GetDefaultTextMapPropagator returns the default propagator for context propagation between peers.

##### func GetSpanID 

``` go
func GetSpanID(ctx context.Context) string
```

GetSpanID retrieves and returns SpanId from context. It returns an empty string is tracing feature is not activated.

##### func GetTraceID 

``` go
func GetTraceID(ctx context.Context) string
```

GetTraceID retrieves and returns TraceId from context. It returns an empty string is tracing feature is not activated.

##### func IsTracingInternal 

``` go
func IsTracingInternal() bool
```

IsTracingInternal returns whether tracing spans of internal components.

##### func IsUsingDefaultProvider 

``` go
func IsUsingDefaultProvider() bool
```

IsUsingDefaultProvider checks and return if currently using default trace provider.

##### func MaxContentLogSize 

``` go
func MaxContentLogSize() int
```

MaxContentLogSize returns the max log size for request and response body, especially for HTTP/RPC request.

##### func SetBaggageMap 

``` go
func SetBaggageMap(ctx context.Context, data map[string]interface{}) context.Context
```

SetBaggageMap is a convenient function for adding map key-value pairs to baggage. Note that it uses attribute.Any to set the key-value pair.

##### func SetBaggageValue 

``` go
func SetBaggageValue(ctx context.Context, key string, value interface{}) context.Context
```

SetBaggageValue is a convenient function for adding one key-value pair to baggage. Note that it uses attribute.Any to set the key-value pair.

##### func WithTraceID 

``` go
func WithTraceID(ctx context.Context, traceID string) (context.Context, error)
```

WithTraceID injects custom trace id into context to propagate.

##### func WithUUID <-2.2.0

``` go
func WithUUID(ctx context.Context, uuid string) (context.Context, error)
```

WithUUID injects custom trace id with UUID into context to propagate.

### Types 

#### type Baggage 

``` go
type Baggage struct {
	// contains filtered or unexported fields
}
```

Baggage holds the data through all tracing spans.

##### func NewBaggage 

``` go
func NewBaggage(ctx context.Context) *Baggage
```

NewBaggage creates and returns a new Baggage object from given tracing context.

##### (*Baggage) Ctx 

``` go
func (b *Baggage) Ctx() context.Context
```

Ctx returns the context that Baggage holds.

##### (*Baggage) GetMap 

``` go
func (b *Baggage) GetMap() *gmap.StrAnyMap
```

GetMap retrieves and returns the baggage values as map.

##### (*Baggage) GetVar 

``` go
func (b *Baggage) GetVar(key string) *gvar.Var
```

GetVar retrieves value and returns a *gvar.Var for specified key from baggage.

##### (*Baggage) SetMap 

``` go
func (b *Baggage) SetMap(data map[string]interface{}) context.Context
```

SetMap is a convenient function for adding map key-value pairs to baggage. Note that it uses attribute.Any to set the key-value pair.

##### (*Baggage) SetValue 

``` go
func (b *Baggage) SetValue(key string, value interface{}) context.Context
```

SetValue is a convenient function for adding one key-value pair to baggage. Note that it uses attribute.Any to set the key-value pair.

#### type Carrier 

``` go
type Carrier map[string]interface{}
```

Carrier is the storage medium used by a TextMapPropagator.

##### func NewCarrier 

``` go
func NewCarrier(data ...map[string]interface{}) Carrier
```

NewCarrier creates and returns a Carrier.

##### (Carrier) Get 

``` go
func (c Carrier) Get(k string) string
```

Get returns the value associated with the passed key.

##### (Carrier) Keys 

``` go
func (c Carrier) Keys() []string
```

Keys lists the keys stored in this carrier.

##### (Carrier) MustMarshal 

``` go
func (c Carrier) MustMarshal() []byte
```

MustMarshal .returns the JSON encoding of c

##### (Carrier) Set 

``` go
func (c Carrier) Set(k, v string)
```

Set stores the key-value pair.

##### (Carrier) String 

``` go
func (c Carrier) String() string
```

String converts and returns current Carrier as string.

##### (Carrier) UnmarshalJSON 

``` go
func (c Carrier) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements interface UnmarshalJSON for package json.

#### type Span 

``` go
type Span struct {
	trace.Span
}
```

Span warps trace.Span for compatibility and extension.

##### func NewSpan 

``` go
func NewSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, *Span)
```

NewSpan creates a span using default tracer.

#### type Tracer 

``` go
type Tracer struct {
	trace.Tracer
}
```

Tracer warps trace.Tracer for compatibility and extension.

##### func NewTracer 

``` go
func NewTracer(name ...string) *Tracer
```

NewTracer Tracer is a short function for retrieving Tracer.