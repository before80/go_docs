+++
title = "o Protobuf：全新的 Opaque API "
date = 2025-03-31T11:34:43+08:00
weight = 840
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/protobuf-opaque](https://go.dev/blog/protobuf-opaque)

#  Go Protobuf: The new Opaque API - Go Protobuf：全新的 Opaque API 

Michael Stapelberg  
16 December 2024  

2024年12月16日  

[Protocol Buffers (Protobuf)](https://en.wikipedia.org/wiki/Protocol_Buffers) is Google’s language-neutral data interchange format. See [protobuf.dev](https://protobuf.dev/).  

​	[Protocol Buffers（Protobuf）](https://en.wikipedia.org/wiki/Protocol_Buffers) 是 Google 的语言中立数据交换格式。详见 [protobuf.dev](https://protobuf.dev/)。  

Back in March 2020, we released the `google.golang.org/protobuf` module, [a major overhaul of the Go Protobuf API](https://go.dev/blog/protobuf-apiv2). This package introduced first-class [support for reflection](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect), a [`dynamicpb`](https://pkg.go.dev/google.golang.org/protobuf/types/dynamicpb) implementation and the [`protocmp`](https://pkg.go.dev/google.golang.org/protobuf/testing/protocmp) package for easier testing.  

​	2020 年 3 月，我们发布了 `google.golang.org/protobuf` 模块（[Go Protobuf API 的重大重构](https://go.dev/blog/protobuf-apiv2)），该包引入了原生[反射支持](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect)、[`dynamicpb`](https://pkg.go.dev/google.golang.org/protobuf/types/dynamicpb) 实现以及便于测试的 [`protocmp`](https://pkg.go.dev/google.golang.org/protobuf/testing/protocmp) 包。  

That release introduced a new protobuf module with a new API. Today, we are releasing an additional API for generated code, meaning the Go code in the `.pb.go` files created by the protocol compiler (`protoc`). This blog post explains our motivation for creating a new API and shows you how to use it in your projects.  

​	该版本引入了全新的 Protobuf 模块及 API。今天，我们发布了一个针对生成代码（即由协议编译器 `protoc` 生成的 `.pb.go` 文件中的 Go 代码）的新增 API。本文阐述了我们创建新 API 的动机，并指导你如何在项目中使用它。  

To be clear: We are not removing anything. We will continue to support the existing API for generated code, just like we still support the older protobuf module (by wrapping the `google.golang.org/protobuf` implementation). Go is [committed to backwards compatibility](https://go.dev/blog/compat) and this applies to Go Protobuf, too!  

​	需要明确的是：我们不会移除任何现有功能。我们将继续支持生成代码的现有 API，正如我们仍支持旧版 Protobuf 模块（通过封装 `google.golang.org/protobuf` 实现）。Go 语言[坚持向后兼容性](https://go.dev/blog/compat)，Go Protobuf 亦遵循此原则！  

## 背景：现有的 Open Struct API - Background: the (existing) Open Struct API  

We now call the existing API the Open Struct API, because generated struct types are open to direct access. In the next section, we will see how it differs from the new Opaque API.  

​	我们将现有 API 称为 **Open Struct API**（开放结构体 API），因为生成的 struct 类型允许直接访问字段。下一节将展示它与新 Opaque API 的区别。  

To work with protocol buffers, you first create a `.proto` definition file like this one:  

​	使用 Protobuf 时，你需先创建如下的 `.proto` 定义文件：  

```protobuf
edition = "2023";  // successor to proto2 and proto3

package log;

message LogEntry {
  string backend_server = 1;
  uint32 request_size = 2;
  string ip_address = 3;
}
```

Then, you [run the protocol compiler (`protoc`)](https://protobuf.dev/getting-started/gotutorial/) to generate code like the following (in a `.pb.go` file):

​	然后，通过[运行协议编译器 `protoc`](https://protobuf.dev/getting-started/gotutorial/) 生成如下代码（位于 `.pb.go` 文件中）：

```go
package logpb

type LogEntry struct {
  BackendServer *string
  RequestSize   *uint32
  IPAddress     *string
  // …internal fields elided…
}

func (l *LogEntry) GetBackendServer() string { … }
func (l *LogEntry) GetRequestSize() uint32   { … }
func (l *LogEntry) GetIPAddress() string     { … }
```

Now you can import the generated `logpb` package from your Go code and call functions like [`proto.Marshal`](https://pkg.go.dev/google.golang.org/protobuf/proto#Marshal) to encode `logpb.LogEntry` messages into protobuf wire format.

​	现在你可在 Go 代码中导入生成的 `logpb` 包，并调用如 [`proto.Marshal`](https://pkg.go.dev/google.golang.org/protobuf/proto#Marshal) 的函数，将 `logpb.LogEntry` 消息编码为 Protobuf 二进制格式。

You can find more details in the [Generated Code API documentation](https://protobuf.dev/reference/go/go-generated/).

​	更多细节请参阅[生成代码 API 文档](https://protobuf.dev/reference/go/go-generated/)。

### 现有 Open Struct API：字段存在性  (Existing) Open Struct API: Field Presence

An important aspect of this generated code is how *field presence* (whether a field is set or not) is modeled. For instance, the above example models presence using pointers, so you could set the `BackendServer` field to:

​	生成代码的一个重要方面是如何建模**字段存在性**（即字段是否被设置）。例如，上述示例使用指针建模存在性，因此 `BackendServer` 字段可设置为：

1. `proto.String("zrh01.prod")`: the field is set and contains “zrh01.prod”
   设置为 `proto.String("zrh01.prod")`：字段已设置且值为 "zrh01.prod"
2. `proto.String("")`: the field is set (non-`nil` pointer) but contains an empty value 
   设置为 `proto.String("")`：字段已设置（非 `nil` 指针）但值为空
3. `nil` pointer: the field is not set 
   `nil` 指针：字段未设置

If you are used to generated code not having pointers, you are probably using `.proto` files that start with `syntax = "proto3"`. The field presence behavior changed over the years:

​	如果你习惯生成代码不含指针，可能使用了以 `syntax = "proto3"` 开头的 `.proto` 文件。字段存在性的行为随时间演变：

- `syntax = "proto2"` uses *explicit presence* by default
  `syntax = "proto2"` 默认使用**显式存在性**
- `syntax = "proto3"` used *implicit presence* by default (where cases 2 and 3 cannot be distinguished and are both represented by an empty string), but was later extended to allow [opting into explicit presence with the `optional` keyword](https://protobuf.dev/programming-guides/proto3/#field-labels)
  `syntax = "proto3"` 默认使用**隐式存在性**（无法区分情况 2 和 3，均以空字符串表示），但后续扩展支持通过 [`optional` 关键字](https://protobuf.dev/programming-guides/proto3/#field-labels)启用显式存在性
- `edition = "2023"`, the [successor to both proto2 and proto3](https://protobuf.dev/editions/overview/), uses [*explicit presence*](https://protobuf.dev/programming-guides/field_presence/) by default
  `edition = "2023"`（[proto2 和 proto3 的继任者](https://protobuf.dev/editions/overview/)）默认使用[显式存在性](https://protobuf.dev/programming-guides/field_presence/)

## 全新的 Opaque API - The new Opaque API

We created the new *Opaque API* to uncouple the [Generated Code API](https://protobuf.dev/reference/go/go-generated/) from the underlying in-memory representation. The (existing) Open Struct API has no such separation: it allows programs direct access to the protobuf message memory. For example, one could use the `flag` package to parse command-line flag values into protobuf message fields:

​	我们创建了 **Opaque API**（不透明 API）以解耦[生成代码 API](https://protobuf.dev/reference/go/go-generated/) 与底层内存表示。现有 Open Struct API 无此分离：它允许程序直接访问 Protobuf 消息内存。例如，可使用 `flag` 包将命令行标志值解析到 Protobuf 消息字段中：

```go
var req logpb.LogEntry
flag.StringVar(&req.BackendServer, "backend", os.Getenv("HOST"), "…")
flag.Parse() // fills the BackendServer field from -backend flag
```

The problem with such a tight coupling is that we can never change how we lay out protobuf messages in memory. Lifting this restriction enables many implementation improvements, which we’ll see below.

​	这种紧耦合的问题在于我们永远无法改变 Protobuf 消息的内存布局。解除此限制可实现多项改进，下文将详述。

What changes with the new Opaque API? Here is how the generated code from the above example would change:

​	Opaque API 带来了哪些变化？以下是上述示例生成代码的变更示例：

```go
package logpb

type LogEntry struct {
  xxx_hidden_BackendServer *string // no longer exported
  xxx_hidden_RequestSize   uint32  // no longer exported
  xxx_hidden_IPAddress     *string // no longer exported
  // …internal fields elided…
}

func (l *LogEntry) GetBackendServer() string { … }
func (l *LogEntry) HasBackendServer() bool   { … }
func (l *LogEntry) SetBackendServer(string)  { … }
func (l *LogEntry) ClearBackendServer()      { … }
// …
```

With the Opaque API, the struct fields are hidden and can no longer be directly accessed. Instead, the new accessor methods allow for getting, setting, or clearing a field.

​	通过 Opaque API，struct 字段被隐藏且无法直接访问，取而代之的是新的访问器方法用于获取、设置或清除字段。

### Opaque 结构体占用更少内存  Opaque structs use less memory

One change we made to the memory layout is to model field presence for elementary fields more efficiently:

​	我们对内存布局的改进之一是更高效地建模基础类型字段的存在性：

- The (existing) Open Struct API uses pointers, which adds a 64-bit word to the space cost of the field.
  现有 Open Struct API 使用指针，这会为字段增加 64 位的空间开销
- The Opaque API uses [bit fields](https://en.wikipedia.org/wiki/Bit_field), which require one bit per field (ignoring padding overhead).
  Opaque API 使用[位域](https://en.wikipedia.org/wiki/Bit_field)，每个字段仅需 1 位（忽略填充开销）

Using fewer variables and pointers also lowers load on the allocator and on the garbage collector.

​	减少变量和指针的使用还可降低分配器和垃圾回收器的负担。

The performance improvement depends heavily on the shapes of your protocol messages: The change only affects elementary fields like integers, bools, enums, and floats, but not strings, repeated fields, or submessages (because it is [less profitable](https://protobuf.dev/reference/go/opaque-faq/#memorylayout) for those types).

​	性能提升高度依赖协议消息的结构：此变更仅影响整数、布尔、枚举和浮点数等基础字段，不影响字符串、重复字段或子消息（因对后者[收益较低](https://protobuf.dev/reference/go/opaque-faq/#memorylayout)）。

Our benchmark results show that messages with few elementary fields exhibit performance that is as good as before, whereas messages with more elementary fields are decoded with significantly fewer allocations:

​	基准测试显示，基础字段少的消息性能与之前持平，而基础字段多的消息解码时分配次数显著减少：

```txt
             │ Open Struct API │             Opaque API             │
             │    allocs/op    │  allocs/op   vs base               │
Prod#1          360.3k ± 0%       360.3k ± 0%  +0.00% (p=0.002 n=6)
Search#1       1413.7k ± 0%       762.3k ± 0%  -46.08% (p=0.002 n=6)
Search#2        314.8k ± 0%       132.4k ± 0%  -57.95% (p=0.002 n=6)
```

Reducing allocations also makes decoding protobuf messages more efficient:

​	减少分配还提升了 Protobuf 消息解码效率：

```txt
             │ Open Struct API │             Opaque API            │
             │   user-sec/op   │ user-sec/op  vs base              │
Prod#1         55.55m ± 6%        55.28m ± 4%  ~ (p=0.180 n=6)
Search#1       324.3m ± 22%       292.0m ± 6%  -9.97% (p=0.015 n=6)
Search#2       67.53m ± 10%       45.04m ± 8%  -33.29% (p=0.002 n=6)
```

(All measurements done on an AMD Castle Peak Zen 2. Results on ARM and Intel CPUs are similar.)

（所有测试基于 AMD Castle Peak Zen 2，ARM 和 Intel CPU 结果类似）

Note: proto3 with implicit presence similarly does not use pointers, so you will not see a performance improvement if you are coming from proto3. If you were using implicit presence for performance reasons, forgoing the convenience of being able to distinguish empty fields from unset ones, then the Opaque API now makes it possible to use explicit presence without a performance penalty.

​	注：使用隐式存在性的 proto3 同样不使用指针，因此从 proto3 迁移不会看到性能提升。若你曾因性能原因使用隐式存在性（牺牲区分空字段与未设置字段的便利性），Opaque API 现可在不损失性能的情况下支持显式存在性。

### 动机：延迟解码 Motivation: Lazy Decoding

Lazy decoding is a performance optimization where the contents of a submessage are decoded when first accessed instead of during [`proto.Unmarshal`](https://pkg.go.dev/google.golang.org/protobuf/proto#Unmarshal). Lazy decoding can improve performance by avoiding unnecessarily decoding fields which are never accessed.

​	**延迟解码**是一种性能优化：子消息内容在首次访问时解码，而非在 [`proto.Unmarshal`](https://pkg.go.dev/google.golang.org/protobuf/proto#Unmarshal) 期间解码。通过避免解码从未访问的字段提升性能。

Lazy decoding can’t be supported safely by the (existing) Open Struct API. While the Open Struct API provides getters, leaving the (un-decoded) struct fields exposed would be extremely error-prone. To ensure that the decoding logic runs immediately before the field is first accessed, we must make the field private and mediate all accesses to it through getter and setter functions.

​	现有 Open Struct API 无法安全支持延迟解码。虽然 Open Struct API 提供 Getter 方法，但暴露未解码的 struct 字段极易引发错误。为确保解码逻辑在字段首次访问前运行，我们必须将字段设为私有，并通过 Getter/Setter 方法中介所有访问。

This approach made it possible to implement lazy decoding with the Opaque API. Of course, not every workload will benefit from this optimization, but for those that do benefit, the results can be spectacular: We have seen logs analysis pipelines that discard messages based on a top-level message condition (e.g. whether `backend_server` is one of the machines running a new Linux kernel version) and can skip decoding deeply nested subtrees of messages.

​	此方法使得 Opaque API 支持延迟解码成为可能。当然，并非所有场景都能受益于此优化，但对适用场景效果显著：我们观察到日志分析流水线可根据顶层消息条件（例如 `backend_server` 是否属于运行新 Linux 内核版本的机器）丢弃消息，从而跳过深度嵌套子消息的解码。

As an example, here are the results of the micro-benchmark we included, demonstrating how lazy decoding saves over 50% of the work and over 87% of allocations!

​	以下是微基准测试结果，展示延迟解码如何节省超 50% 的工作量和 87% 的分配次数：

As an example, here are the results of the micro-benchmark we included, demonstrating how lazy decoding saves over 50% of the work and over 87% of allocations!

​	例如，以下是我们包含的微基准测试的结果，展示了延迟解码如何节省超过50%的工作和超过87%的分配！

```txt
                  │   nolazy    │                lazy                │
                  │   sec/op    │   sec/op     vs base               │
Unmarshal/lazy-24   6.742µ ± 0%   2.816µ ± 0%  -58.23% (p=0.002 n=6)

                  │    nolazy    │                lazy                 │
                  │     B/op     │     B/op      vs base               │
Unmarshal/lazy-24   3.666Ki ± 0%   1.814Ki ± 0%  -50.51% (p=0.002 n=6)

                  │   nolazy    │               lazy                │
                  │  allocs/op  │ allocs/op   vs base               │
Unmarshal/lazy-24   64.000 ± 0%   8.000 ± 0%  -87.50% (p=0.002 n=6)
```



### 动机：减少指针比较错误 Motivation: reduce pointer comparison mistakes

Modeling field presence with pointers invites pointer-related bugs.

​	通过指针建模字段存在性容易引发指针相关的错误。 

Consider an enum, declared within the `LogEntry` message:

​	假设在 `LogEntry` 消息中声明了一个枚举类型：  

```protobuf
message LogEntry {
  enum DeviceType {
    DESKTOP = 0;
    MOBILE = 1;
    VR = 2;
  };
  DeviceType device_type = 1;
}
```

A simple mistake is to compare the `device_type` enum field like so:

​	一个常见错误是以下列方式比较枚举字段 `device_type`：

```go
if cv.DeviceType == logpb.LogEntry_DESKTOP.Enum() { // incorrect!
```

Did you spot the bug? The condition compares the memory address instead of the value. Because the `Enum()` accessor allocates a new variable on each call, the condition can never be true. The check should have read:

​	你是否发现了问题？该条件比较的是内存地址而非实际值。由于 `Enum()` 访问器每次调用都会分配新变量，此条件永远无法成立。正确写法应为：

```go
if cv.GetDeviceType() == logpb.LogEntry_DESKTOP {
```

The new Opaque API prevents this mistake: Because fields are hidden, all access must go through the getter.

​	Opaque API 避免了此错误：由于字段被隐藏，所有访问必须通过 Getter 方法。

### 动机：减少意外共享错误 Motivation: reduce accidental sharing mistakes

Let’s consider a slightly more involved pointer-related bug. Assume you are trying to stabilize an RPC service that fails under high load. The following part of the request middleware looks correct, but still the entire service goes down whenever just one customer sends a high volume of requests:

​	让我们看一个更复杂的指针相关错误案例。假设你正在修复一个高负载下崩溃的 RPC 服务，以下请求中间件代码看似正确，但当单个客户发送大量请求时服务仍会崩溃：

```go
logEntry.IPAddress = req.IPAddress
logEntry.BackendServer = proto.String(hostname)
// The redactIP() function redacts IPAddress to 127.0.0.1,
// unexpectedly not just in logEntry *but also* in req!
// redactIP() 函数会将 IPAddress 脱敏为 127.0.0.1，
// 但意外地同时修改了 logEntry 和 req 中的 IPAddress！
go auditlog(redactIP(logEntry))
if quotaExceeded(req) {
    // BUG: All requests end up here, regardless of their source.
    // BUG: 所有请求都会进入此分支，无论来源如何
    return fmt.Errorf("server overloaded")
}
```

Did you spot the bug? The first line accidentally copied the pointer (thereby sharing the pointed-to variable between the `logEntry` and `req` messages) instead of its value. It should have read:

​	是否发现了问题？第一行错误地复制了指针（导致 `logEntry` 和 `req` 消息共享同一指针指向的变量），而非复制值。正确写法应为：

```go
logEntry.IPAddress = proto.String(req.GetIPAddress())
```

The new Opaque API prevents this problem as the setter takes a value (`string`) instead of a pointer:

​	Opaque API 通过 Setter 方法接收值（`string`）而非指针，从而避免了此问题：

```go
logEntry.SetIPAddress(req.GetIPAddress())
```

### 动机：消除反射的尖锐问题 Motivation: Fix Sharp Edges: reflection

To write code that works not only with a specific message type (e.g. `logpb.LogEntry`), but with any message type, one needs some kind of reflection. The previous example used a function to redact IP addresses. To work with any type of message, it could have been defined as `func redactIP(proto.Message) proto.Message { … }`.

​	编写适用于任意消息类型（而不仅是特定类型如 `logpb.LogEntry`）的代码需要某种形式的反射。前例中的 IP 地址脱敏函数可定义为 `func redactIP(proto.Message) proto.Message { … }` 以支持所有消息类型。

Many years ago, your only option to implement a function like `redactIP` was to reach for [Go’s `reflect` package](https://go.dev/blog/laws-of-reflection), which resulted in very tight coupling: you had only the generator output and had to reverse-engineer what the input protobuf message definition might have looked like. The [`google.golang.org/protobuf` module release](https://go.dev/blog/protobuf-apiv2) (from March 2020) introduced [Protobuf reflection](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect), which should always be preferred: Go’s `reflect` package traverses the data structure’s representation, which should be an implementation detail. Protobuf reflection traverses the logical tree of protocol messages without regard to its representation.

​	多年前，实现类似 `redactIP` 的函数只能使用 [Go 的 `reflect` 包](https://go.dev/blog/laws-of-reflection)，这导致强耦合：开发者需根据生成代码逆向推断原始 Protobuf 消息定义。2020 年 3 月发布的 [`google.golang.org/protobuf` 模块](https://go.dev/blog/protobuf-apiv2) 引入了 [Protobuf 反射](https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect)，应始终优先使用：Go 的 `reflect` 包遍历数据结构的内存表示（属于实现细节），而 Protobuf 反射则遍历消息的逻辑树结构，与内存布局无关。

Unfortunately, merely *providing* protobuf reflection is not sufficient and still leaves some sharp edges exposed: In some cases, users might accidentally use Go reflection instead of protobuf reflection.

​	但仅提供 Protobuf 反射仍存在隐患：用户可能误用 Go 反射而非 Protobuf 反射。

For example, encoding a protobuf message with the `encoding/json` package (which uses Go reflection) was technically possible, but the result is not [canonical Protobuf JSON encoding](https://protobuf.dev/programming-guides/proto3/#json). Use the [`protojson`](https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson) package instead.

​	例如，使用 `encoding/json` 包（依赖 Go 反射）编码 Protobuf 消息虽技术上可行，但结果不符合 [Protobuf 标准 JSON 编码规范](https://protobuf.dev/programming-guides/proto3/#json)。应改用 [`protojson`](https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson) 包。

The new Opaque API prevents this problem because the message struct fields are hidden: accidental usage of Go reflection will see an empty message. This is clear enough to steer developers towards protobuf reflection.

​	Opaque API 解决了此问题：消息结构体字段被隐藏，误用 Go 反射时只能看到空消息。这种明显的异常能有效引导开发者使用 Protobuf 反射。

### 动机：实现理想内存布局 Motivation: Making the ideal memory layout possible

The benchmark results from the [More Efficient Memory Representation](https://go.dev/blog/protobuf-opaque#lessmemory) section have already shown that protobuf performance heavily depends on the specific usage: How are the messages defined? Which fields are set?

​	[更高效的内存表示](https://go.dev/blog/protobuf-opaque#lessmemory) 章节的基准测试表明，Protobuf 性能高度依赖具体使用场景：消息如何定义？哪些字段被设置？

To keep Go Protobuf as fast as possible for *everyone*, we cannot implement optimizations that help only one program, but hurt the performance of other programs.

​	为确保 Go Protobuf 对所有人保持高性能，我们不能实施仅优化特定程序而损害其他程序性能的改进。

The Go compiler used to be in a similar situation, up until [Go 1.20 introduced Profile-Guided Optimization (PGO)](https://go.dev/blog/go1.20). By recording the production behavior (through [profiling](https://go.dev/blog/pprof)) and feeding that profile back to the compiler, we allow the compiler to make better trade-offs *for a specific program or workload*.

​	Go 编译器曾面临类似困境，直到 [Go 1.20 引入 Profile-Guided Optimization (PGO)](https://go.dev/blog/go1.20)。通过记录生产环境行为（通过 [性能分析](https://go.dev/blog/pprof)）并将分析结果反馈给编译器，编译器可为特定程序或负载做出更优权衡。

We think using profiles to optimize for specific workloads is a promising approach for further Go Protobuf optimizations. The Opaque API makes those possible: Program code uses accessors and does not need to be updated when the memory representation changes, so we could, for example, move rarely set fields into an overflow struct.

​	我们认为基于性能分析针对特定负载优化是 Go Protobuf 未来的重要方向。Opaque API 为此铺平道路：程序代码使用访问器，内存布局变更时无需修改代码。例如，可将极少设置的字段移至溢出结构体中。

## 迁移 Migration

You can migrate on your own schedule, or even not at all—the (existing) Open Struct API will not be removed. But, if you’re not on the new Opaque API, you won’t benefit from its improved performance, or future optimizations that target it.

​	你可按自己的节奏迁移，甚至完全不迁移——现有 Open Struct API 不会被移除。但若不采用 Opaque API，将无法享受其性能提升及未来优化。

We recommend you select the Opaque API for new development. Protobuf Edition 2024 (see [Protobuf Editions Overview](https://protobuf.dev/editions/overview/) if you are not yet familiar) will make the Opaque API the default.

​	建议新项目选择 Opaque API。Protobuf 2024 版本（详见 [Protobuf 版本概述](https://protobuf.dev/editions/overview/)）将默认采用 Opaque API。

### 混合 API - The Hybrid API

Aside from the Open Struct API and Opaque API, there is also the Hybrid API, which keeps existing code working by keeping struct fields exported, but also enabling migration to the Opaque API by adding the new accessor methods.

​	除 Open Struct API 和 Opaque API 外，还有**混合 API**：保留导出的结构体字段以兼容旧代码，同时添加新访问器方法以支持迁移至 Opaque API。

With the Hybrid API, the protobuf compiler will generate code on two API levels: the `.pb.go` is on the Hybrid API, whereas the `_protoopaque.pb.go` version is on the Opaque API and can be selected by building with the `protoopaque` build tag.

​	混合 API 下，Protobuf 编译器生成两个 API 层级的代码：`.pb.go` 为混合 API，`_protoopaque.pb.go` 为 Opaque API，可通过 `protoopaque` 构建标签选择。

### 代码迁移至 Opaque API - Rewriting Code to the Opaque API

See the [migration guide](https://protobuf.dev/reference/go/opaque-migration/) for detailed instructions. The high-level steps are:

​	详见 [迁移指南](https://protobuf.dev/reference/go/opaque-migration/)，主要步骤包括：

1. Enable the Hybrid API. 启用混合 API
2. Update existing code using the `open2opaque` migration tool. 使用 `open2opaque` 迁移工具更新现有代码
3. Switch to the Opaque API. 切换至 Opaque API

### 对已发布生成代码的建议：使用混合 API - Advice for published generated code: Use Hybrid API

Small usages of protobuf can live entirely within the same repository, but usually, `.proto` files are shared between different projects that are owned by different teams. An obvious example is when different companies are involved: To call Google APIs (with protobuf), use the [Google Cloud Client Libraries for Go](https://github.com/googleapis/google-cloud-go) from your project. Switching the Cloud Client Libraries to the Opaque API is not an option, as that would be a breaking API change, but switching to the Hybrid API is safe.

​	小型 Protobuf 应用可完全在单一仓库内管理，但通常 `.proto` 文件会在不同团队的项目间共享。典型场景是跨公司协作：调用 Google API 时需使用项目中的 [Go 版 Google Cloud 客户端库](https://github.com/googleapis/google-cloud-go)。将客户端库切换至 Opaque API 会导致 API 破坏性变更，但切换至混合 API 是安全的。

Our advice for such packages that publish generated code (`.pb.go` files) is to switch to the Hybrid API please! Publish both the `.pb.go` and the `_protoopaque.pb.go` files, please. The `protoopaque` version allows your consumers to migrate on their own schedule.

​	对发布生成代码（`.pb.go` 文件）的包，我们建议切换至混合 API！请同时发布 `.pb.go` 和 `_protoopaque.pb.go` 文件，`protoopaque` 版本让用户可自主迁移。

### 启用延迟解码 Enabling Lazy Decoding

Lazy decoding is available (but not enabled) once you migrate to the Opaque API! 🎉

​	迁移至 Opaque API 后即可使用延迟解码功能（默认未启用）！🎉

To enable: in your `.proto` file, annotate your message-typed fields with the `[lazy = true]` annotation.

​	启用方法：在 `.proto` 文件中为消息类型字段添加 `[lazy = true]` 注解。

To opt out of lazy decoding (despite `.proto` annotations), the [`protolazy` package documentation](https://pkg.go.dev/google.golang.org/protobuf/runtime/protolazy) describes the available opt-outs, which affect either an individual Unmarshal operation or the entire program.

​	若需禁用延迟解码（即使有 `.proto` 注解），[`protolazy` 包文档](https://pkg.go.dev/google.golang.org/protobuf/runtime/protolazy) 描述了可用的退出机制，支持针对单个 Unmarshal 操作或整个程序禁用。

## 后续步骤 Next Steps

By using the open2opaque tool in an automated fashion over the last few years, we have converted the vast majority of Google’s `.proto` files and Go code to the Opaque API. We continuously improved the Opaque API implementation as we moved more and more production workloads to it.

​	过去几年，我们通过自动化使用 `open2opaque` 工具，已将 Google 绝大多数 `.proto` 文件和 Go 代码迁移至 Opaque API。随着更多生产负载迁移，我们持续优化了 Opaque API 实现。

Therefore, we expect you should not encounter problems when trying the Opaque API. In case you do encounter any issues after all, please [let us know on the Go Protobuf issue tracker](https://github.com/golang/protobuf/issues/).

​	因此，预计你在尝试 Opaque API 时不会遇到问题。若仍有问题，请通过 [Go Protobuf 问题追踪器](https://github.com/golang/protobuf/issues/) 反馈。

Reference documentation for Go Protobuf can be found on [protobuf.dev → Go Reference](https://protobuf.dev/reference/go/).

​	Go Protobuf 参考文档详见 [protobuf.dev → Go 参考](https://protobuf.dev/reference/go/)。