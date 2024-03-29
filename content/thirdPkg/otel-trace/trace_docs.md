+++
title = "trace文档"
date = 2024-03-28T15:17:42+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/go.opentelemetry.io/otel/trace@v1.24.0](https://pkg.go.dev/go.opentelemetry.io/otel/trace@v1.24.0)

## Overview 概述

- [API Implementations - API 实现]({{< ref "/thirdPkg/otel-trace/trace_docs#api-实现">}})

Package trace provides an implementation of the tracing part of the OpenTelemetry API.

​	trace 包提供 OpenTelemetry API 的跟踪部分的实现。

To participate in distributed traces a Span needs to be created for the operation being performed as part of a traced workflow. In its simplest form:

​	要参与分布式跟踪，需要为作为跟踪工作流的一部分执行的操作创建一个 Span。以最简单的形式：

```go
var tracer trace.Tracer

func init() {
	tracer = otel.Tracer("instrumentation/package/name")
}

func operation(ctx context.Context) {
	var span trace.Span
	ctx, span = tracer.Start(ctx, "operation")
	defer span.End()
	// ...
}
```

A Tracer is unique to the instrumentation and is used to create Spans. Instrumentation should be designed to accept a TracerProvider from which it can create its own unique Tracer. Alternatively, the registered global TracerProvider from the go.opentelemetry.io/otel package can be used as a default.

​	Tracer 是检测的唯一值，用于创建 Span。检测应设计为接受 TracerProvider，从中可以创建自己唯一的 Tracer。或者，可以将 go.opentelemetry.io/otel 包中注册的全局 TracerProvider 用作默认值。

```go
const (
	name    = "instrumentation/package/name"
	version = "0.1.0"
)

type Instrumentation struct {
	tracer trace.Tracer
}

func NewInstrumentation(tp trace.TracerProvider) *Instrumentation {
	if tp == nil {
		tp = otel.TracerProvider()
	}
	return &Instrumentation{
		tracer: tp.Tracer(name, trace.WithInstrumentationVersion(version)),
	}
}

func operation(ctx context.Context, inst *Instrumentation) {
	var span trace.Span
	ctx, span = inst.tracer.Start(ctx, "operation")
	defer span.End()
	// ...
}
```

### API 实现

This package does not conform to the standard Go versioning policy; all of its interfaces may have methods added to them without a package major version bump. This non-standard API evolution could surprise an uninformed implementation author. They could unknowingly build their implementation in a way that would result in a runtime panic for their users that update to the new API.

​	此包不符合标准的 Go 版本控制策略;它的所有接口都可能添加了方法，而不会出现包主要版本颠簸。这种非标准的 API 演进可能会让不知情的实现作者感到惊讶。他们可能会在不知不觉中构建他们的实现，从而导致更新到新 API 的用户出现运行时恐慌。

The API is designed to help inform an instrumentation author about this non-standard API evolution. It requires them to choose a default behavior for unimplemented interface methods. There are three behavior choices they can make:

​	该 API 旨在帮助检测作者了解这种非标准 API 的演变。它要求他们为未实现的接口方法选择默认行为。他们可以做出三种行为选择：

- Compilation failure
  编译失败
- Panic
  恐慌
- Default to another implementation
  默认为其他实现

All interfaces in this API embed a corresponding interface from [go.opentelemetry.io/otel/trace/embedded](https://pkg.go.dev/go.opentelemetry.io/otel/trace@v1.24.0/embedded). If an author wants the default behavior of their implementations to be a compilation failure, signaling to their users they need to update to the latest version of that implementation, they need to embed the corresponding interface from [go.opentelemetry.io/otel/trace/embedded](https://pkg.go.dev/go.opentelemetry.io/otel/trace@v1.24.0/embedded) in their implementation. For example,

​	此 API 中的所有接口都嵌入了 go.opentelemetry.io/otel/trace/embedded 中的相应接口。如果作者希望其实现的默认行为是编译失败，向用户发出他们需要更新到该实现的最新版本的信号，则需要在其实现中嵌入 go.opentelemetry.io/otel/trace/embedded 相应的接口。例如

```go
import "go.opentelemetry.io/otel/trace/embedded"

type TracerProvider struct {
	embedded.TracerProvider
	// ...
}
```

If an author wants the default behavior of their implementations to panic, they can embed the API interface directly.

​	如果作者希望其实现的默认行为出现恐慌，他们可以直接嵌入 API 接口。

```go
import "go.opentelemetry.io/otel/trace"

type TracerProvider struct {
	trace.TracerProvider
	// ...
}
```

This option is not recommended. It will lead to publishing packages that contain runtime panics when users update to newer versions of [go.opentelemetry.io/otel/trace](https://pkg.go.dev/go.opentelemetry.io/otel/trace@v1.24.0), which may be done with a trasitive dependency.

​	不建议使用此选项。当用户更新到较新版本的 go.opentelemetry.io/otel/trace 时，这将导致发布包含运行时崩溃的包，这可以通过 trasive 依赖项来完成。

Finally, an author can embed another implementation in theirs. The embedded implementation will be used for methods not defined by the author. For example, an author who wants to default to silently dropping the call can use [go.opentelemetry.io/otel/trace/noop](https://pkg.go.dev/go.opentelemetry.io/otel/trace@v1.24.0/noop):

​	最后，作者可以在他们的实现中嵌入另一个实现。嵌入式实现将用于作者未定义的方法。例如，想要默认为静默挂断调用的作者可以使用 go.opentelemetry.io/otel/trace/noop：

```go
import "go.opentelemetry.io/otel/trace/noop"

type TracerProvider struct {
	noop.TracerProvider
	// ...
}
```

It is strongly recommended that authors only embed [go.opentelemetry.io/otel/trace/noop](https://pkg.go.dev/go.opentelemetry.io/otel/trace@v1.24.0/noop) if they choose this default behavior. That implementation is the only one OpenTelemetry authors can guarantee will fully implement all the API interfaces when a user updates their API.

​	强烈建议作者仅在选择此默认行为时才嵌入 go.opentelemetry.io/otel/trace/noop。该实现是 OpenTelemetry 作者唯一可以保证在用户更新其 API 时完全实现所有 API 接口的实现。

## 常量

[View Source](https://github.com/open-telemetry/opentelemetry-go/blob/trace/v1.24.0/trace/trace.go#L28)

```go
const (
	// FlagsSampled is a bitmask with the sampled bit set. A SpanContext
	// with the sampling bit set means the span is sampled.
	FlagsSampled = TraceFlags(0x01)
)
```

## 变量

This section is empty.

## 函数

#### func ContextWithRemoteSpanContext

```go
func ContextWithRemoteSpanContext(parent context.Context, rsc SpanContext) context.Context
```

ContextWithRemoteSpanContext returns a copy of parent with rsc set explicly as a remote SpanContext and as the current Span. The Span implementation that wraps rsc is non-recording and performs no operations other than to return rsc as the SpanContext from the SpanContext method.

​	ContextWithRemoteSpanContext 返回父级的副本，其中 rsc 显式设置为远程 SpanContext 和当前 Span。包装 rsc 的 Span 实现是非记录的，除了从 SpanContext 方法返回 rsc 作为 SpanContext 之外，不执行任何操作。

#### func ContextWithSpan

```go
func ContextWithSpan(parent context.Context, span Span) context.Context
```

ContextWithSpan returns a copy of parent with span set as the current Span.

​	ContextWithSpan 返回 span 设置为当前 Span 的父级副本。

#### func ContextWithSpanContext <- 0.20.0

```go
func ContextWithSpanContext(parent context.Context, sc SpanContext) context.Context
```

ContextWithSpanContext returns a copy of parent with sc as the current Span. The Span implementation that wraps sc is non-recording and performs no operations other than to return sc as the SpanContext from the SpanContext method.

​	ContextWithSpanContext 返回父级的副本，其中 sc 为当前 Span。包装 sc 的 Span 实现不记录，除了从 SpanContext 方法返回 sc 作为 SpanContext 之外，不执行任何操作。

## Types

### type EventConfig <- 1.0.0

```go
type EventConfig struct {
	// contains filtered or unexported fields
}
```

EventConfig is a group of options for an Event.

​	EventConfig 是 Event 的一组选项。

#### func NewEventConfig

```go
func NewEventConfig(options ...EventOption) EventConfig
```

NewEventConfig applies all the EventOptions to a returned EventConfig. If no timestamp option is passed, the returned EventConfig will have a Timestamp set to the call time, otherwise no validation is performed on the returned EventConfig.

​	NewEventConfig 将所有 EventOptions 应用于返回的 EventConfig。如果未传递时间戳选项，则返回的 EventConfig 将 Timestamp 设置为调用时间，否则不会对返回的 EventConfig 执行验证。

#### (*EventConfig) Attributes <- 1.0.0 

```go
func (cfg *EventConfig) Attributes() []attribute.KeyValue
```

Attributes describe the associated qualities of an Event.

​	属性描述事件的关联质量。

#### (*EventConfig) StackTrace <- 1.0.0 

```go
func (cfg *EventConfig) StackTrace() bool
```

StackTrace checks whether stack trace capturing is enabled.

​	StackTrace 检查是否启用了堆栈跟踪捕获。

#### (*EventConfig) Timestamp <- 1.0.0 

```go
func (cfg *EventConfig) Timestamp() time.Time
```

Timestamp is a time in an Event life-cycle.

​	时间戳是事件生命周期中的时间。

### type EventOption 

```go
type EventOption interface {
	// contains filtered or unexported methods
}
```

EventOption applies span event options to an EventConfig.

​	EventOption 将 span 事件选项应用于 EventConfig。

### type Link 

```go
type Link struct {
	// SpanContext of the linked Span.
	SpanContext SpanContext

	// Attributes describe the aspects of the link.
	Attributes []attribute.KeyValue
}
```

Link is the relationship between two Spans. The relationship can be within the same Trace or across different Traces.

​	链接是两个跨度之间的关系。关系可以位于同一跟踪中，也可以跨不同跟踪。

For example, a Link is used in the following situations:

​	例如，在以下情况下使用链接：

1. Batch Processing: A batch of operations may contain operations associated with one or more traces/spans. Since there can only be one parent SpanContext, a Link is used to keep reference to the SpanContext of all operations in the batch.
   批处理：一批操作可能包含与一个或多个跟踪/跨度关联的操作。由于只能有一个父 SpanContext，因此使用 Link 来保留对批处理中所有操作的 SpanContext 的引用。
2. Public Endpoint: A SpanContext for an in incoming client request on a public endpoint should be considered untrusted. In such a case, a new trace with its own identity and sampling decision needs to be created, but this new trace needs to be related to the original trace in some form. A Link is used to keep reference to the original SpanContext and track the relationship.
   公共终结点：公共终结点上传入客户端请求的 SpanContext 应被视为不受信任。在这种情况下，需要创建具有自己的标识和采样决策的新跟踪，但此新跟踪需要以某种形式与原始跟踪相关联。链接用于保留对原始 SpanContext 的引用并跟踪关系。

#### func LinkFromContext <- 1.0.0

```go
func LinkFromContext(ctx context.Context, attrs ...attribute.KeyValue) Link
```

LinkFromContext returns a link encapsulating the SpanContext in the provided ctx.

​	LinkFromContext 返回一个链接，该链接将 SpanContext 封装在提供的 ctx 中。

### type Span 

```go
type Span interface {
	// Users of the interface can ignore this. This embedded type is only used
	// by implementations of this interface. See the "API Implementations"
	// section of the package documentation for more information.
	embedded.Span

	// End completes the Span. The Span is considered complete and ready to be
	// delivered through the rest of the telemetry pipeline after this method
	// is called. Therefore, updates to the Span are not allowed after this
	// method has been called.
	End(options ...SpanEndOption)

	// AddEvent adds an event with the provided name and options.
	AddEvent(name string, options ...EventOption)

	// IsRecording returns the recording state of the Span. It will return
	// true if the Span is active and events can be recorded.
	IsRecording() bool

	// RecordError will record err as an exception span event for this span. An
	// additional call to SetStatus is required if the Status of the Span should
	// be set to Error, as this method does not change the Span status. If this
	// span is not being recorded or err is nil then this method does nothing.
	RecordError(err error, options ...EventOption)

	// SpanContext returns the SpanContext of the Span. The returned SpanContext
	// is usable even after the End method has been called for the Span.
	SpanContext() SpanContext

	// SetStatus sets the status of the Span in the form of a code and a
	// description, provided the status hasn't already been set to a higher
	// value before (OK > Error > Unset). The description is only included in a
	// status when the code is for an error.
	SetStatus(code codes.Code, description string)

	// SetName sets the Span name.
	SetName(name string)

	// SetAttributes sets kv as attributes of the Span. If a key from kv
	// already exists for an attribute of the Span it will be overwritten with
	// the value contained in kv.
	SetAttributes(kv ...attribute.KeyValue)

	// TracerProvider returns a TracerProvider that can be used to generate
	// additional Spans on the same telemetry pipeline as the current Span.
	TracerProvider() TracerProvider
}
```

Span is the individual component of a trace. It represents a single named and timed operation of a workflow that is traced. A Tracer is used to create a Span and it is then up to the operation the Span represents to properly end the Span when the operation itself ends.

​	Span 是跟踪的各个组件。它表示所跟踪的工作流的单个命名和定时操作。Tracer 用于创建 Span，然后由该 Span 表示的操作在操作本身结束时正确结束 Span。

Warning: Methods may be added to this interface in minor releases. See package documentation on API implementation for information on how to set default behavior for unimplemented methods.

​	警告：在次要版本中，可能会将方法添加到此接口。有关如何为未实现的方法设置默认行为的信息，请参阅有关 API 实现的包文档。

#### func SpanFromContext

```go
func SpanFromContext(ctx context.Context) Span
```

SpanFromContext returns the current Span from ctx.

​	SpanFromContext 从 ctx 返回当前 Span。

If no Span is currently set in ctx an implementation of a Span that performs no operations is returned.

​	如果当前未在 ctx 中设置 Span，则返回不执行任何操作的 Span 实现。

### type SpanConfig 

```go
type SpanConfig struct {
	// contains filtered or unexported fields
}
```

SpanConfig is a group of options for a Span.

​	SpanConfig 是 Span 的一组选项。

#### func NewSpanEndConfig <- 1.0.0

```go
func NewSpanEndConfig(options ...SpanEndOption) SpanConfig
```

NewSpanEndConfig applies all the options to a returned SpanConfig. No validation is performed on the returned SpanConfig (e.g. no uniqueness checking or bounding of data), it is left to the SDK to perform this action.

​	NewSpanEndConfig 将所有选项应用于返回的 SpanConfig。不会对返回的 SpanConfig 执行任何验证（例如，没有唯一性检查或数据边界），而是留给 SDK 执行此操作。

#### func NewSpanStartConfig <- 1.0.0

```go
func NewSpanStartConfig(options ...SpanStartOption) SpanConfig
```

NewSpanStartConfig applies all the options to a returned SpanConfig. No validation is performed on the returned SpanConfig (e.g. no uniqueness checking or bounding of data), it is left to the SDK to perform this action.

​	NewSpanStartConfig 将所有选项应用于返回的 SpanConfig。不会对返回的 SpanConfig 执行任何验证（例如，没有唯一性检查或数据边界），而是留给 SDK 执行此操作。

#### (*SpanConfig) Attributes 

```go
func (cfg *SpanConfig) Attributes() []attribute.KeyValue
```

Attributes describe the associated qualities of a Span.

​	属性描述 Span 的关联质量。

#### (*SpanConfig) Links 

```go
func (cfg *SpanConfig) Links() []Link
```

Links are the associations a Span has with other Spans.

​	链接是 Span 与其他 Span 的关联。

#### (*SpanConfig) NewRoot 

```go
func (cfg *SpanConfig) NewRoot() bool
```

NewRoot identifies a Span as the root Span for a new trace. This is commonly used when an existing trace crosses trust boundaries and the remote parent span context should be ignored for security.

​	NewRoot 将 Span 标识为新跟踪的根 Span。当现有跟踪跨越信任边界时，通常会使用此方法，并且为了安全起见，应忽略远程父跨域上下文。

#### (*SpanConfig) SpanKind 

```go
func (cfg *SpanConfig) SpanKind() SpanKind
```

SpanKind is the role a Span has in a trace.

​	SpanKind 是 Span 在跟踪中的角色。

#### (*SpanConfig) StackTrace <- 1.0.0

```go
func (cfg *SpanConfig) StackTrace() bool
```

StackTrace checks whether stack trace capturing is enabled.

​	StackTrace 检查是否启用了堆栈跟踪捕获。

#### (*SpanConfig) Timestamp

```go
func (cfg *SpanConfig) Timestamp() time.Time
```

Timestamp is a time in a Span life-cycle.

​	时间戳是 Span 生命周期中的时间。

### type SpanContext

```go
type SpanContext struct {
	// contains filtered or unexported fields
}
```

SpanContext contains identifying trace information about a Span.

​	SpanContext 包含有关 Span 的标识跟踪信息。

#### func NewSpanContext <- 0.19.0

```go
func NewSpanContext(config SpanContextConfig) SpanContext
```

NewSpanContext constructs a SpanContext using values from the provided SpanContextConfig.

​	NewSpanContext 使用提供的 SpanContextConfig 中的值构造 SpanContext。

#### func SpanContextFromContext

```go
func SpanContextFromContext(ctx context.Context) SpanContext
```

SpanContextFromContext returns the current Span’s SpanContext.

​	SpanContextFromContext 返回当前 Span 的 SpanContext。

#### (SpanContext) Equal <- 0.19.0 

```go
func (sc SpanContext) Equal(other SpanContext) bool
```

Equal is a predicate that determines whether two SpanContext values are equal.

​	Equal 是一个谓词，用于确定两个 SpanContext 值是否相等。

#### (SpanContext) HasSpanID 

```go
func (sc SpanContext) HasSpanID() bool
```

HasSpanID checks if the SpanContext has a valid SpanID.

​	HasSpanID 检查 SpanContext 是否具有有效的 SpanID。

#### (SpanContext) HasTraceID 

```go
func (sc SpanContext) HasTraceID() bool
```

HasTraceID checks if the SpanContext has a valid TraceID.

​	HasTraceID 检查 SpanContext 是否具有有效的 TraceID。

#### (SpanContext) IsRemote <- 0.19.0

```go
func (sc SpanContext) IsRemote() bool
```

IsRemote indicates whether the SpanContext represents a remotely-created Span.

​	IsRemote 指示 SpanContext 是否表示远程创建的 Span。

#### (SpanContext) IsSampled 

```go
func (sc SpanContext) IsSampled() bool
```

IsSampled returns if the sampling bit is set in the SpanContext’s TraceFlags.

​	如果在 SpanContext 的 TraceFlags 中设置了采样位，则返回 IsSampled。

#### (SpanContext) IsValid 

```go
func (sc SpanContext) IsValid() bool
```

IsValid returns if the SpanContext is valid. A valid span context has a valid TraceID and SpanID.

​	如果 SpanContext 有效，则 IsValid 返回。有效的 span 上下文具有有效的 TraceID 和 SpanID。

#### (SpanContext) MarshalJSON <- 0.19.0 

```go
func (sc SpanContext) MarshalJSON() ([]byte, error)
```

MarshalJSON implements a custom marshal function to encode a SpanContext.

​	MarshalJSON 实现自定义封送器函数来对 SpanContext 进行编码。

#### (SpanContext) SpanID

```go
func (sc SpanContext) SpanID() SpanID
```

SpanID returns the SpanID from the SpanContext.

​	SpanID 从 SpanContext 返回 SpanID。

#### (SpanContext) TraceFlags 

```go
func (sc SpanContext) TraceFlags() TraceFlags
```

TraceFlags returns the flags from the SpanContext.

​	TraceFlags 从 SpanContext 返回标志。

#### (SpanContext) TraceID 

```go
func (sc SpanContext) TraceID() TraceID
```

TraceID returns the TraceID from the SpanContext.

​	TraceID 从 SpanContext 返回 TraceID。

#### (SpanContext) TraceState 

```go
func (sc SpanContext) TraceState() TraceState
```

TraceState returns the TraceState from the SpanContext.

​	TraceState 从 SpanContext 返回 TraceState。

#### (SpanContext) WithRemote <- 0.19.0

```go
func (sc SpanContext) WithRemote(remote bool) SpanContext
```

WithRemote returns a copy of sc with the Remote property set to remote.

​	WithRemote 返回 sc 的副本，其中 Remote 属性设置为 remote。

#### (SpanContext) WithSpanID <- 0.19.0 

```go
func (sc SpanContext) WithSpanID(spanID SpanID) SpanContext
```

WithSpanID returns a new SpanContext with the SpanID replaced.

​	WithSpanID 返回替换了 SpanID 的新 SpanContext。

#### (SpanContext) WithTraceFlags <- 0.19.0 

```go
func (sc SpanContext) WithTraceFlags(flags TraceFlags) SpanContext
```

WithTraceFlags returns a new SpanContext with the TraceFlags replaced.

​	WithTraceFlags 返回一个新的 SpanContext，其中替换了 TraceFlags。

#### (SpanContext) WithTraceID <- 0.19.0 

```go
func (sc SpanContext) WithTraceID(traceID TraceID) SpanContext
```

WithTraceID returns a new SpanContext with the TraceID replaced.

​	WithTraceID 返回替换了 TraceID 的新 SpanContext。

#### (SpanContext) WithTraceState <- 0.19.0 

```go
func (sc SpanContext) WithTraceState(state TraceState) SpanContext
```

WithTraceState returns a new SpanContext with the TraceState replaced.

​	WithTraceState 返回替换了 TraceState 的新 SpanContext。

### type SpanContextConfig <- 0.19.0 

```go
type SpanContextConfig struct {
	TraceID    TraceID
	SpanID     SpanID
	TraceFlags TraceFlags
	TraceState TraceState
	Remote     bool
}
```

SpanContextConfig contains mutable fields usable for constructing an immutable SpanContext.

​	SpanContextConfig 包含可用于构造不可变 SpanContext 的可变字段。

### type SpanEndEventOption <- 1.0.0

```go
type SpanEndEventOption interface {
	SpanEndOption
	EventOption
}
```

SpanEndEventOption are options that can be used at the end of a span, or with an event.

​	SpanEndEventOption 是可在跨度末尾或与事件一起使用的选项。

#### func WithStackTrace <- 1.0.0

```go
func WithStackTrace(b bool) SpanEndEventOption
```

WithStackTrace sets the flag to capture the error with stack trace (e.g. true, false).

​	WithStackTrace 设置标志以捕获堆栈跟踪错误（例如 true、false）。

### type SpanEndOption <- 1.0.0

```go
type SpanEndOption interface {
	// contains filtered or unexported methods
}
```

SpanEndOption applies an option to a SpanConfig. These options are applicable only when the span is ended.

​	SpanEndOption 将选项应用于 SpanConfig。这些选项仅在跨度结束时适用。

### type SpanEventOption <- 1.0.0 

```go
type SpanEventOption interface {
	SpanOption
	EventOption
}
```

SpanEventOption are options that can be used with an event or a span.

​	SpanEventOption 是可用于事件或跨度的选项。

#### func WithTimestamp

```go
func WithTimestamp(t time.Time) SpanEventOption
```

WithTimestamp sets the time of a Span or Event life-cycle moment (e.g. started, stopped, errored).

​	WithTimestamp 设置跨度或事件生命周期时刻的时间（例如，已启动、已停止、出错）。

### type SpanID

```go
type SpanID [8]byte
```

SpanID is a unique identity of a span in a trace.

​	SpanID 是跟踪中跨度的唯一标识。

#### func SpanIDFromHex

```go
func SpanIDFromHex(h string) (SpanID, error)
```

SpanIDFromHex returns a SpanID from a hex string if it is compliant with the w3c trace-context specification. See more at https://www.w3.org/TR/trace-context/#parent-id

​	如果 SpanIDFromHex 符合 w3c 跟踪上下文规范，则从十六进制字符串返回 SpanID。在 https://www.w3.org/TR/trace-context/#parent-id 上查看更多信息

#### (SpanID) IsValid

```go
func (s SpanID) IsValid() bool
```

IsValid checks whether the SpanID is valid. A valid SpanID does not consist of zeros only.

​	IsValid 检查 SpanID 是否有效。有效的 SpanID 不仅由零组成。

#### (SpanID) MarshalJSON

```go
func (s SpanID) MarshalJSON() ([]byte, error)
```

MarshalJSON implements a custom marshal function to encode SpanID as a hex string.

​	MarshalJSON 实现自定义封送函数，将 SpanID 编码为十六进制字符串。

#### (SpanID) String

```go
func (s SpanID) String() string
```

String returns the hex string representation form of a SpanID.

​	String 返回 SpanID 的十六进制字符串表示形式。

### type SpanKind

```go
type SpanKind int
```

SpanKind is the role a Span plays in a Trace.

​	SpanKind 是 Span 在跟踪中扮演的角色。

```go
const (
	// SpanKindUnspecified is an unspecified SpanKind and is not a valid
	// SpanKind. SpanKindUnspecified should be replaced with SpanKindInternal
	// if it is received.
	SpanKindUnspecified SpanKind = 0
	// SpanKindInternal is a SpanKind for a Span that represents an internal
	// operation within an application.
	SpanKindInternal SpanKind = 1
	// SpanKindServer is a SpanKind for a Span that represents the operation
	// of handling a request from a client.
	SpanKindServer SpanKind = 2
	// SpanKindClient is a SpanKind for a Span that represents the operation
	// of client making a request to a server.
	SpanKindClient SpanKind = 3
	// SpanKindProducer is a SpanKind for a Span that represents the operation
	// of a producer sending a message to a message broker. Unlike
	// SpanKindClient and SpanKindServer, there is often no direct
	// relationship between this kind of Span and a SpanKindConsumer kind. A
	// SpanKindProducer Span will end once the message is accepted by the
	// message broker which might not overlap with the processing of that
	// message.
	SpanKindProducer SpanKind = 4
	// SpanKindConsumer is a SpanKind for a Span that represents the operation
	// of a consumer receiving a message from a message broker. Like
	// SpanKindProducer Spans, there is often no direct relationship between
	// this Span and the Span that produced the message.
	SpanKindConsumer SpanKind = 5
)
```

As a convenience, these match the proto definition, see https://github.com/open-telemetry/opentelemetry-proto/blob/30d237e1ff3ab7aa50e0922b5bebdd93505090af/opentelemetry/proto/trace/v1/trace.proto#L101-L129

​	为方便起见，这些与原始定义相匹配，请参阅 https://github.com/open-telemetry/opentelemetry-proto/blob/30d237e1ff3ab7aa50e0922b5bebdd93505090af/opentelemetry/proto/trace/v1/trace.proto#L101-L129

The unspecified value is not a valid `SpanKind`. Use `ValidateSpanKind()` to coerce a span kind to a valid value.

​	未指定的值不是有效的 `SpanKind` 。用于 `ValidateSpanKind()` 强制 span 类型为有效值。

#### func ValidateSpanKind

```go
func ValidateSpanKind(spanKind SpanKind) SpanKind
```

ValidateSpanKind returns a valid span kind value. This will coerce invalid values into the default value, SpanKindInternal.

​	ValidateSpanKind 返回有效的 span kind 值。这会将无效值强制转换为默认值 SpanKindInternal。

#### (SpanKind) String

```go
func (sk SpanKind) String() string
```

String returns the specified name of the SpanKind in lower-case.

​	String 以小写形式返回 SpanKind 的指定名称。

### type SpanOption

```go
type SpanOption interface {
	SpanStartOption
	SpanEndOption
}
```

SpanOption are options that can be used at both the beginning and end of a span.

​	SpanOption 是可以在跨度的开始和结束时使用的选项。

### type SpanStartEventOption <- 1.0.0

```go
type SpanStartEventOption interface {
	SpanStartOption
	EventOption
}
```

SpanStartEventOption are options that can be used at the start of a span, or with an event.

​	SpanStartEventOption 是可在跨度开始时使用的选项，也可以与事件一起使用。

#### func WithAttributes

```go
func WithAttributes(attributes ...attribute.KeyValue) SpanStartEventOption
```

WithAttributes adds the attributes related to a span life-cycle event. These attributes are used to describe the work a Span represents when this option is provided to a Span’s start or end events. Otherwise, these attributes provide additional information about the event being recorded (e.g. error, state change, processing progress, system event).

​	WithAttributes 添加与跨度生命周期事件相关的属性。这些属性用于描述在向 Span 的开始或结束事件提供此选项时 Span 所代表的工作。否则，这些属性将提供有关正在记录的事件的其他信息（例如错误、状态更改、处理进度、系统事件）。

If multiple of these options are passed the attributes of each successive option will extend the attributes instead of overwriting. There is no guarantee of uniqueness in the resulting attributes.

​	如果传递了其中多个选项，则每个连续选项的属性将扩展属性，而不是覆盖。无法保证生成的属性的唯一性。

### type SpanStartOption <- 1.0.0

```go
type SpanStartOption interface {
	// contains filtered or unexported methods
}
```

SpanStartOption applies an option to a SpanConfig. These options are applicable only when the span is created.

​	SpanStartOption 将选项应用于 SpanConfig。这些选项仅在创建跨度时适用。

#### func WithLinks

```go
func WithLinks(links ...Link) SpanStartOption
```

WithLinks adds links to a Span. The links are added to the existing Span links, i.e. this does not overwrite. Links with invalid span context are ignored.

​	WithLinks 将链接添加到 Span。这些链接被添加到现有的 Span 链接中，即这不会覆盖。将忽略具有无效 span 上下文的链接。

#### func WithNewRoot

```go
func WithNewRoot() SpanStartOption
```

WithNewRoot specifies that the Span should be treated as a root Span. Any existing parent span context will be ignored when defining the Span’s trace identifiers.

​	WithNewRoot 指定应将 Span 视为根 Span。定义 Span 的跟踪标识符时，将忽略任何现有的父 span 上下文。

#### func WithSpanKind

```go
func WithSpanKind(kind SpanKind) SpanStartOption
```

WithSpanKind sets the SpanKind of a Span.

​	WithSpanKind 将 SpanKind 设置为 Span。

### type TraceFlags <- 0.20.0

```go
type TraceFlags byte //nolint:revive // revive complains about stutter of `trace.TraceFlags`.
```

TraceFlags contains flags that can be set on a SpanContext.

​	TraceFlags 包含可在 SpanContext 上设置的标志。

#### (TraceFlags) IsSampled <- 0.20.0

```go
func (tf TraceFlags) IsSampled() bool
```

IsSampled returns if the sampling bit is set in the TraceFlags.

​	如果在 TraceFlags 中设置了采样位，则返回 IsSampled。

#### (TraceFlags) MarshalJSON <- 0.20.0

```go
func (tf TraceFlags) MarshalJSON() ([]byte, error)
```

MarshalJSON implements a custom marshal function to encode TraceFlags as a hex string.

​	MarshalJSON 实现自定义封送函数，将 TraceFlags 编码为十六进制字符串。

#### (TraceFlags) String <- 0.20.0

```go
func (tf TraceFlags) String() string
```

String returns the hex string representation form of TraceFlags.

​	String 返回 TraceFlags 的十六进制字符串表示形式。

#### (TraceFlags) WithSampled <- 0.20.0

```go
func (tf TraceFlags) WithSampled(sampled bool) TraceFlags
```

WithSampled sets the sampling bit in a new copy of the TraceFlags.

​	WithSampled 在 TraceFlags 的新副本中设置采样位。

### type TraceID

```go
type TraceID [16]byte
```

TraceID is a unique identity of a trace. nolint:revive // revive complains about stutter of `trace.TraceID`.

​	TraceID 是跟踪的唯一标识。nolint：revive // revive 抱怨 `trace.TraceID` .

#### func TraceIDFromHex

```go
func TraceIDFromHex(h string) (TraceID, error)
```

TraceIDFromHex returns a TraceID from a hex string if it is compliant with the W3C trace-context specification. See more at https://www.w3.org/TR/trace-context/#trace-id nolint:revive // revive complains about stutter of `trace.TraceIDFromHex`.

​	如果 TraceIDFromHex 符合 W3C 跟踪上下文规范，则从十六进制字符串返回 TraceID。有关详细信息 https://www.w3.org/TR/trace-context/#trace-id nolint：revive // revive 抱怨口吃。 `trace.TraceIDFromHex`

#### (TraceID) IsValid

```go
func (t TraceID) IsValid() bool
```

IsValid checks whether the trace TraceID is valid. A valid trace ID does not consist of zeros only.

​	IsValid 检查跟踪 TraceID 是否有效。有效的跟踪 ID 不仅由零组成。

#### (TraceID) MarshalJSON

```go
func (t TraceID) MarshalJSON() ([]byte, error)
```

MarshalJSON implements a custom marshal function to encode TraceID as a hex string.

​	MarshalJSON 实现自定义封送函数，将 TraceID 编码为十六进制字符串。

#### (TraceID) String

```go
func (t TraceID) String() string
```

String returns the hex string representation form of a TraceID.

​	String 返回 TraceID 的十六进制字符串表示形式。

### type TraceState

```go
type TraceState struct {
	// contains filtered or unexported fields
}
```

TraceState provides additional vendor-specific trace identification information across different distributed tracing systems. It represents an immutable list consisting of key/value pairs, each pair is referred to as a list-member.

​	TraceState 在不同的分布式跟踪系统中提供其他特定于供应商的跟踪标识信息。它表示一个由键/值对组成的不可变列表，每个对称为列表成员。

TraceState conforms to the W3C Trace Context specification (https://www.w3.org/TR/trace-context-1). All operations that create or copy a TraceState do so by validating all input and will only produce TraceState that conform to the specification. Specifically, this means that all list-member’s key/value pairs are valid, no duplicate list-members exist, and the maximum number of list-members (32) is not exceeded.

​	TraceState 符合 W3C 跟踪上下文规范 （https://www.w3.org/TR/trace-context-1）。创建或复制 TraceState 的所有操作都通过验证所有输入来执行此操作，并且只会生成符合规范的 TraceState。具体而言，这意味着所有列表成员的键/值对都有效，不存在重复的列表成员，并且不超过列表成员的最大数量 （32）。

#### func ParseTraceState <- 1.0.0

```go
func ParseTraceState(ts string) (TraceState, error)
```

ParseTraceState attempts to decode a TraceState from the passed string. It returns an error if the input is invalid according to the W3C Trace Context specification.

​	ParseTraceState 尝试从传递的字符串中解码 TraceState。如果根据 W3C 跟踪上下文规范，输入无效，则返回错误。

#### (TraceState) Delete

```go
func (ts TraceState) Delete(key string) TraceState
```

Delete returns a copy of the TraceState with the list-member identified by key removed.

​	Delete 返回 TraceState 的副本，其中 list-member 由已删除的键标识。

#### (TraceState) Get

```go
func (ts TraceState) Get(key string) string
```

Get returns the value paired with key from the corresponding TraceState list-member if it exists, otherwise an empty string is returned.

​	Get 返回与相应 TraceState list-member 中的 key 配对的值（如果存在），否则返回空字符串。

#### (TraceState) Insert

```go
func (ts TraceState) Insert(key, value string) (TraceState, error)
```

Insert adds a new list-member defined by the key/value pair to the TraceState. If a list-member already exists for the given key, that list-member’s value is updated. The new or updated list-member is always moved to the beginning of the TraceState as specified by the W3C Trace Context specification.

​	Insert 将键/值对定义的新列表成员添加到 TraceState。如果给定键已存在 list-member，则会更新该 list-member 的值。新的或更新的 list-member 始终移动到 W3C 跟踪上下文规范指定的 TraceState 的开头。

If key or value are invalid according to the W3C Trace Context specification an error is returned with the original TraceState.

​	如果键或值根据 W3C 跟踪上下文规范无效，则返回原始 TraceState 错误。

If adding a new list-member means the TraceState would have more members then is allowed, the new list-member will be inserted and the right-most list-member will be dropped in the returned TraceState.

​	如果添加新的列表成员意味着 TraceState 将具有更多成员，则将插入新的列表成员，并将最右边的列表成员删除到返回的 TraceState 中。

#### (TraceState) Len <- 1.0.0

```go
func (ts TraceState) Len() int
```

Len returns the number of list-members in the TraceState.

​	Len 返回 TraceState 中的列表成员数。

#### (TraceState) MarshalJSON

```go
func (ts TraceState) MarshalJSON() ([]byte, error)
```

MarshalJSON marshals the TraceState into JSON.

​	MarshalJSON 将 TraceState 封送到 JSON 中。

#### (TraceState) String

```go
func (ts TraceState) String() string
```

String encodes the TraceState into a string compliant with the W3C Trace Context specification. The returned string will be invalid if the TraceState contains any invalid members.

​	String 将 TraceState 编码为符合 W3C 跟踪上下文规范的字符串。如果 TraceState 包含任何无效成员，则返回的字符串将无效。

### type Tracer

```go
type Tracer interface {
	// Users of the interface can ignore this. This embedded type is only used
	// by implementations of this interface. See the "API Implementations"
	// section of the package documentation for more information.
	embedded.Tracer

	// Start creates a span and a context.Context containing the newly-created span.
	//
	// If the context.Context provided in `ctx` contains a Span then the newly-created
	// Span will be a child of that span, otherwise it will be a root span. This behavior
	// can be overridden by providing `WithNewRoot()` as a SpanOption, causing the
	// newly-created Span to be a root span even if `ctx` contains a Span.
	//
	// When creating a Span it is recommended to provide all known span attributes using
	// the `WithAttributes()` SpanOption as samplers will only have access to the
	// attributes provided when a Span is created.
	//
	// Any Span that is created MUST also be ended. This is the responsibility of the user.
	// Implementations of this API may leak memory or other resources if Spans are not ended.
	Start(ctx context.Context, spanName string, opts ...SpanStartOption) (context.Context, Span)
}
```

Tracer is the creator of Spans.

​	Tracer 是 Span 的创建者。

Warning: Methods may be added to this interface in minor releases. See package documentation on API implementation for information on how to set default behavior for unimplemented methods.

​	警告：在次要版本中，可能会将方法添加到此接口。有关如何为未实现的方法设置默认行为的信息，请参阅有关 API 实现的包文档。

### type TracerConfig

```go
type TracerConfig struct {
	// contains filtered or unexported fields
}
```

TracerConfig is a group of options for a Tracer.

​	TracerConfig 是 Tracer 的一组选项。

#### func NewTracerConfig

```go
func NewTracerConfig(options ...TracerOption) TracerConfig
```

NewTracerConfig applies all the options to a returned TracerConfig.

​	NewTracerConfig 将所有选项应用于返回的 TracerConfig。

#### (*TracerConfig) InstrumentationAttributes <- 1.14.0

```go
func (t *TracerConfig) InstrumentationAttributes() attribute.Set
```

InstrumentationAttributes returns the attributes associated with the library providing instrumentation.

​	InstrumentationAttributes 返回与提供检测的库关联的属性。

#### (*TracerConfig) InstrumentationVersion

```go
func (t *TracerConfig) InstrumentationVersion() string
```

InstrumentationVersion returns the version of the library providing instrumentation.

​	InstrumentationVersion 返回提供检测的库的版本。

#### (*TracerConfig) SchemaURL <- 1.0.0

```go
func (t *TracerConfig) SchemaURL() string
```

SchemaURL returns the Schema URL of the telemetry emitted by the Tracer.

​	SchemaURL 返回跟踪器发出的遥测数据的架构 URL。

### type TracerOption

```go
type TracerOption interface {
	// contains filtered or unexported methods
}
```

TracerOption applies an option to a TracerConfig.

​	TracerOption 将选项应用于 TracerConfig。

#### func WithInstrumentationAttributes <- 1.14.0

```go
func WithInstrumentationAttributes(attr ...attribute.KeyValue) TracerOption
```

WithInstrumentationAttributes sets the instrumentation attributes.

​	WithInstrumentationAttributes 设置检测属性。

The passed attributes will be de-duplicated.

​	传递的属性将被删除重复。

#### func WithInstrumentationVersion

```go
func WithInstrumentationVersion(version string) TracerOption
```

WithInstrumentationVersion sets the instrumentation version.

​	WithInstrumentationVersion 设置检测版本。

#### func WithSchemaURL <- 1.0.0

```go
func WithSchemaURL(schemaURL string) TracerOption
```

WithSchemaURL sets the schema URL for the Tracer.

​	WithSchemaURL 设置跟踪器的架构 URL。

### type TracerProvider

```go
type TracerProvider interface {
	// Users of the interface can ignore this. This embedded type is only used
	// by implementations of this interface. See the "API Implementations"
	// section of the package documentation for more information.
	embedded.TracerProvider

	// Tracer returns a unique Tracer scoped to be used by instrumentation code
	// to trace computational workflows. The scope and identity of that
	// instrumentation code is uniquely defined by the name and options passed.
	//
	// The passed name needs to uniquely identify instrumentation code.
	// Therefore, it is recommended that name is the Go package name of the
	// library providing instrumentation (note: not the code being
	// instrumented). Instrumentation libraries can have multiple versions,
	// therefore, the WithInstrumentationVersion option should be used to
	// distinguish these different codebases. Additionally, instrumentation
	// libraries may sometimes use traces to communicate different domains of
	// workflow data (i.e. using spans to communicate workflow events only). If
	// this is the case, the WithScopeAttributes option should be used to
	// uniquely identify Tracers that handle the different domains of workflow
	// data.
	//
	// If the same name and options are passed multiple times, the same Tracer
	// will be returned (it is up to the implementation if this will be the
	// same underlying instance of that Tracer or not). It is not necessary to
	// call this multiple times with the same name and options to get an
	// up-to-date Tracer. All implementations will ensure any TracerProvider
	// configuration changes are propagated to all provided Tracers.
	//
	// If name is empty, then an implementation defined default name will be
	// used instead.
	//
	// This method is safe to call concurrently.
	Tracer(name string, options ...TracerOption) Tracer
}
```

TracerProvider provides Tracers that are used by instrumentation code to trace computational workflows.

​	TracerProvider 提供检测代码用于跟踪计算工作流的跟踪器。

A TracerProvider is the collection destination of all Spans from Tracers it provides, it represents a unique telemetry collection pipeline. How that pipeline is defined, meaning how those Spans are collected, processed, and where they are exported, depends on its implementation. Instrumentation authors do not need to define this implementation, rather just use the provided Tracers to instrument code.

​	TracerProvider 是它提供的跟踪器中所有 Span 的收集目标，它表示唯一的遥测收集管道。如何定义该管道，即如何收集、处理这些 Span 以及导出它们的位置，取决于其实现。检测作者不需要定义此实现，而只需使用提供的跟踪器来检测代码。

Commonly, instrumentation code will accept a TracerProvider implementation at runtime from its users or it can simply use the globally registered one (see https://pkg.go.dev/go.opentelemetry.io/otel#GetTracerProvider).

​	通常，检测代码将在运行时接受其用户的 TracerProvider 实现，或者它可以简单地使用全局注册的实现（参见 https://pkg.go.dev/go.opentelemetry.io/otel#GetTracerProvider）。

Warning: Methods may be added to this interface in minor releases. See package documentation on API implementation for information on how to set default behavior for unimplemented methods.

​	警告：在次要版本中，可能会将方法添加到此接口。有关如何为未实现的方法设置默认行为的信息，请参阅有关 API 实现的包文档。