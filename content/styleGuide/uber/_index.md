+++
title = "Uber Go 风格指南"
date = 2024-01-22T09:32:55+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Uber Go Style Guide - Uber Go 风格指南

> 原文： [https://github.com/uber-go/guide/blob/master/style.md](https://github.com/uber-go/guide/blob/master/style.md)

## Introduction 简介

Styles are the conventions that govern our code. The term style is a bit of a misnomer, since these conventions cover far more than just source file formatting—gofmt handles that for us.

​	风格是我们代码遵循的惯例。术语风格有点名不副实，因为这些惯例涵盖的范围远不止源文件格式——gofmt 会为我们处理这些。

The goal of this guide is to manage this complexity by describing in detail the Dos and Don’ts of writing Go code at Uber. These rules exist to keep the code base manageable while still allowing engineers to use Go language features productively.

​	本指南的目的是通过详细描述在 Uber 编写 Go 代码的注意事项来管理这种复杂性。这些规则的存在是为了保持代码库的可管理性，同时仍允许工程师高效地使用 Go 语言特性。

This guide was originally created by [Prashant Varanasi](https://github.com/prashantv) and [Simon Newton](https://github.com/nomis52) as a way to bring some colleagues up to speed with using Go. Over the years it has been amended based on feedback from others.

​	这份指南最初由 Prashant Varanasi 和 Simon Newton 创建，目的是让一些同事快速了解如何使用 Go。多年来，它根据他人的反馈进行了修改。

This documents idiomatic conventions in Go code that we follow at Uber. A lot of these are general guidelines for Go, while others extend upon external resources:

​	本文档记录了我们在 Uber 遵循的 Go 代码中的惯用约定。其中许多是 Go 的一般准则，而其他一些则扩展了外部资源：

1. [Effective Go
   有效的 Go](https://golang.org/doc/effective_go.html)
2. [Go Common Mistakes
   常见的 Go 错误](https://github.com/golang/go/wiki/CommonMistakes)
3. [Go Code Review Comments
   Go 代码审查评论](https://github.com/golang/go/wiki/CodeReviewComments)

We aim for the code samples to be accurate for the two most recent minor versions of Go [releases](https://go.dev/doc/devel/release).

​	我们的目标是让代码示例适用于最近两个次要版本的 Go 版本。

All code should be error-free when run through `golint` and `go vet`. We recommend setting up your editor to:

​	通过 `golint` 和 `go vet` 运行时，所有代码都应无错误。我们建议将编辑器设置为：

- Run `goimports` on save
  保存时运行 `goimports`
- Run `golint` and `go vet` to check for errors
  运行 `golint` 和 `go vet` 以检查错误

You can find information in editor support for Go tools here: https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

​	您可以在此处找到有关 Go 工具的编辑器支持信息：https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

## Guidelines 指南

### Pointers to Interfaces 指向接口的指针

You almost never need a pointer to an interface. You should be passing interfaces as values—the underlying data can still be a pointer.

​	您几乎不需要指向接口的指针。您应该将接口作为值传递——底层数据仍然可以是指针。

An interface is two fields:

​	接口是两个字段：

1. A pointer to some type-specific information. You can think of this as “type.”
   指向某些类型特定信息的指针。您可以将其视为“类型”。
2. Data pointer. If the data stored is a pointer, it’s stored directly. If the data stored is a value, then a pointer to the value is stored.
   数据指针。如果存储的数据是指针，则直接存储。如果存储的数据是值，则存储指向该值的指针。

If you want interface methods to modify the underlying data, you must use a pointer.

​	如果希望接口方法修改底层数据，则必须使用指针。

### Verify Interface Compliance 验证接口合规性

Verify interface compliance at compile time where appropriate. This includes:

​	在适当的时候在编译时验证接口合规性。这包括：

- Exported types that are required to implement specific interfaces as part of their API contract
  作为其 API 契约的一部分，需要实现特定接口的导出类型
- Exported or unexported types that are part of a collection of types implementing the same interface
  作为实现相同接口的类型集合的一部分的已导出或未导出的类型
- Other cases where violating an interface would break users
  违反接口会破坏用户的情况

```go
// Bad
type Handler struct {
  // ...
}



func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  ...
}
// Good
type Handler struct {
  // ...
}

var _ http.Handler = (*Handler)(nil)

func (h *Handler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}
```

The statement `var _ http.Handler = (*Handler)(nil)` will fail to compile if `*Handler` ever stops matching the `http.Handler` interface.

​	如果 `*Handler` 停止匹配 `http.Handler` 接口，则语句 `var _ http.Handler = (*Handler)(nil)` 将无法编译。

The right hand side of the assignment should be the zero value of the asserted type. This is `nil` for pointer types (like `*Handler`), slices, and maps, and an empty struct for struct types.

​	赋值的右侧应该是断言类型的零值。对于指针类型（如 `*Handler` ）、切片和映射，这是 `nil` ，对于结构类型，这是空结构。 接收者和接口

```go
type LogHandler struct {
  h   http.Handler
  log *zap.Logger
}

var _ http.Handler = LogHandler{}

func (h LogHandler) ServeHTTP(
  w http.ResponseWriter,
  r *http.Request,
) {
  // ...
}
```

### Receivers and Interfaces 具有值接收者的方法可以调用指针和值。具有指针接收者的方法只能调用指针或可寻址值。

Methods with value receivers can be called on pointers as well as values. Methods with pointer receivers can only be called on pointers or [addressable values](https://golang.org/ref/spec#Method_values).

​	例如，

For example,

​	同样，即使方法具有值接收者，接口也可以由指针满足。

```go
type S struct {
  data string
}

func (s S) Read() string {
  return s.data
}

func (s *S) Write(str string) {
  s.data = str
}

// We cannot get pointers to values stored in maps, because they are not
// addressable values.
sVals := map[int]S{1: {"A"}}

// We can call Read on values stored in the map because Read
// has a value receiver, which does not require the value to
// be addressable.
sVals[1].Read()

// We cannot call Write on values stored in the map because Write
// has a pointer receiver, and it's not possible to get a pointer
// to a value stored in a map.
//
//  sVals[1].Write("test")

sPtrs := map[int]*S{1: {"A"}}

// You can call both Read and Write if the map stores pointers,
// because pointers are intrinsically addressable.
sPtrs[1].Read()
sPtrs[1].Write("test")
```

Similarly, an interface can be satisfied by a pointer, even if the method has a value receiver.

​	Effective Go 对指针与值进行了很好的撰写。

```go
type F interface {
  f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

s1Val := S1{}
s1Ptr := &S1{}
s2Val := S2{}
s2Ptr := &S2{}

var i F
i = s1Val
i = s1Ptr
i = s2Ptr

// The following doesn't compile, since s2Val is a value, and there is no value receiver for f.
//   i = s2Val
```

Effective Go has a good write up on [Pointers vs. Values](https://golang.org/doc/effective_go.html#pointers_vs_values).

​	零值互斥锁有效

### Zero-value Mutexes are Valid

The zero-value of `sync.Mutex` and `sync.RWMutex` is valid, so you almost never need a pointer to a mutex.

​	 `sync.Mutex` 和 `sync.RWMutex` 的零值有效，因此您几乎不需要指向互斥体的指针。

```go
// Bad
mu := new(sync.Mutex)
mu.Lock()
// Good
var mu sync.Mutex
mu.Lock()
```

If you use a struct by pointer, then the mutex should be a non-pointer field on it. Do not embed the mutex on the struct, even if the struct is not exported.

​	如果您通过指针使用结构体，那么互斥体应该是非指针字段。即使结构体未导出，也不要将互斥体嵌入到结构体中。

```go
// Bad
type SMap struct {
  sync.Mutex

  data map[string]string
}

func NewSMap() *SMap {
  return &SMap{
    data: make(map[string]string),
  }
}

func (m *SMap) Get(k string) string {
  m.Lock()
  defer m.Unlock()

  return m.data[k]
}
// Good
type SMap struct {
  mu sync.Mutex

  data map[string]string
}

func NewSMap() *SMap {
  return &SMap{
    data: make(map[string]string),
  }
}

func (m *SMap) Get(k string) string {
  m.mu.Lock()
  defer m.mu.Unlock()

  return m.data[k]
}
```

### Copy Slices and Maps at Boundaries 在边界处复制切片和映射

Slices and maps contain pointers to the underlying data so be wary of scenarios when they need to be copied.

​	切片和映射包含指向基础数据的指针，因此在需要复制它们时要小心。

#### Receiving Slices and Maps 接收切片和映射

Keep in mind that users can modify a map or slice you received as an argument if you store a reference to it.

​	请记住，如果您存储对映射或切片的引用，则用户可以修改您作为参数接收的映射或切片。

```go
// Bad
func (d *Driver) SetTrips(trips []Trip) {
  d.trips = trips
}

trips := ...
d1.SetTrips(trips)

// Did you mean to modify d1.trips?
trips[0] = ...
// Good
func (d *Driver) SetTrips(trips []Trip) {
  d.trips = make([]Trip, len(trips))
  copy(d.trips, trips)
}

trips := ...
d1.SetTrips(trips)

// We can now modify trips[0] without affecting d1.trips.
trips[0] = ...
```

#### Returning Slices and Maps 返回切片和映射

Similarly, be wary of user modifications to maps or slices exposing internal state.

​	同样，要小心用户对映射或切片的修改，这些修改会公开内部状态。

```go
// Bad
type Stats struct {
  mu sync.Mutex
  counters map[string]int
}

// Snapshot returns the current stats.
func (s *Stats) Snapshot() map[string]int {
  s.mu.Lock()
  defer s.mu.Unlock()

  return s.counters
}

// snapshot is no longer protected by the mutex, so any
// access to the snapshot is subject to data races.
snapshot := stats.Snapshot()
// Good
type Stats struct {
  mu sync.Mutex
  counters map[string]int
}

func (s *Stats) Snapshot() map[string]int {
  s.mu.Lock()
  defer s.mu.Unlock()

  result := make(map[string]int, len(s.counters))
  for k, v := range s.counters {
    result[k] = v
  }
  return result
}

// Snapshot is now a copy.
snapshot := stats.Snapshot()
```

### Defer to Clean Up 推迟清理

Use defer to clean up resources such as files and locks.

​	使用 defer 清理资源，例如文件和锁。

```go
// Bad
p.Lock()
if p.count < 10 {
  p.Unlock()
  return p.count
}

p.count++
newCount := p.count
p.Unlock()

return newCount

// easy to miss unlocks due to multiple returns
// Good
p.Lock()
defer p.Unlock()

if p.count < 10 {
  return p.count
}

p.count++
return p.count

// more readable
```

Defer has an extremely small overhead and should be avoided only if you can prove that your function execution time is in the order of nanoseconds. The readability win of using defers is worth the miniscule cost of using them. This is especially true for larger methods that have more than simple memory accesses, where the other computations are more significant than the `defer`.

​	延迟的开销极小，只有当您可以证明函数执行时间以纳秒为单位时才应避免使用延迟。使用延迟的可读性优势值得付出极小的使用成本。对于具有不止简单内存访问的较大方法尤其如此，其中其他计算比 `defer` 更重要。

### Channel Size is One or None 通道大小为一或无

Channels should usually have a size of one or be unbuffered. By default, channels are unbuffered and have a size of zero. Any other size must be subject to a high level of scrutiny. Consider how the size is determined, what prevents the channel from filling up under load and blocking writers, and what happens when this occurs.

​	通道通常应大小为一或无缓冲。默认情况下，通道无缓冲且大小为零。任何其他大小都必须经过严格审查。考虑如何确定大小，防止通道在负载下填满并阻塞写入器，以及发生这种情况时会发生什么。

```go
// Bad
// Ought to be enough for anybody!
c := make(chan int, 64)
// Good
// Size of one
c := make(chan int, 1) // or
// Unbuffered channel, size of zero
c := make(chan int)
```

### Start Enums at One 从一处开始枚举

The standard way of introducing enumerations in Go is to declare a custom type and a `const` group with `iota`. Since variables have a 0 default value, you should usually start your enums on a non-zero value.

​	在 Go 中引入枚举的标准方法是声明一个自定义类型和一个带有 `iota` 的 `const` 组。由于变量的默认值为 0，因此您通常应从非零值开始枚举。

```go
// Bad
type Operation int

const (
  Add Operation = iota
  Subtract
  Multiply
)

// Add=0, Subtract=1, Multiply=2
// Good
type Operation int

const (
  Add Operation = iota + 1
  Subtract
  Multiply
)

// Add=1, Subtract=2, Multiply=3
```

There are cases where using the zero value makes sense, for example when the zero value case is the desirable default behavior.

​	在某些情况下，使用零值是有意义的，例如当零值情况是理想的默认行为时。

```go
type LogOutput int

const (
  LogToStdout LogOutput = iota
  LogToFile
  LogToRemote
)

// LogToStdout=0, LogToFile=1, LogToRemote=2
```

### Use `"time"` to handle time 使用 `"time"` 处理时间

Time is complicated. Incorrect assumptions often made about time include the following.

​	时间很复杂。关于时间经常做出的错误假设包括以下内容。

1. A day has 24 hours
   一天有 24 小时
2. An hour has 60 minutes
   一小时有 60 分钟
3. A week has 7 days
   一周有 7 天
4. A year has 365 days
   一年有 365 天
5. [And a lot more
   等等](https://infiniteundo.com/post/25326999628/falsehoods-programmers-believe-about-time)

For example, *1* means that adding 24 hours to a time instant will not always yield a new calendar day.

​	例如，1 表示将 24 小时添加到时间瞬间并不总是会产生新的日历日。

Therefore, always use the [`"time"`](https://golang.org/pkg/time/) package when dealing with time because it helps deal with these incorrect assumptions in a safer, more accurate manner.

​	因此，在处理时间时始终使用 `"time"` 包，因为它有助于以更安全、更准确的方式处理这些错误假设。

#### Use `time.Time` for instants of time 对时间瞬间使用 `time.Time`

Use [`time.Time`](https://golang.org/pkg/time/#Time) when dealing with instants of time, and the methods on `time.Time` when comparing, adding, or subtracting time.

​	在处理时间瞬间时使用 `time.Time` ，并在比较、添加或减去时间时使用 `time.Time` 上的方法。

```go
// Bad
func isActive(now, start, stop int) bool {
  return start <= now && now < stop
}
// Good
func isActive(now, start, stop time.Time) bool {
  return (start.Before(now) || start.Equal(now)) && now.Before(stop)
}
```

#### Use `time.Duration` for periods of time 使用 `time.Duration` 表示时间段

Use [`time.Duration`](https://golang.org/pkg/time/#Duration) when dealing with periods of time.

​	处理时间段时，请使用 `time.Duration` 。

```go
// Bad
func poll(delay int) {
  for {
    // ...
    time.Sleep(time.Duration(delay) * time.Millisecond)
  }
}

poll(10) // was it seconds or milliseconds?
// Good
func poll(delay time.Duration) {
  for {
    // ...
    time.Sleep(delay)
  }
}

poll(10*time.Second)
```

Going back to the example of adding 24 hours to a time instant, the method we use to add time depends on intent. If we want the same time of the day, but on the next calendar day, we should use [`Time.AddDate`](https://golang.org/pkg/time/#Time.AddDate). However, if we want an instant of time guaranteed to be 24 hours after the previous time, we should use [`Time.Add`](https://golang.org/pkg/time/#Time.Add).

​	回到将 24 小时添加到时间戳的示例，我们用于添加时间的方法取决于意图。如果我们想要同一天的相同时间，但位于下一个日历日，我们应该使用 `Time.AddDate` 。但是，如果我们想要一个时间戳，保证在之前的时间之后 24 小时，我们应该使用 `Time.Add` 。

```
newDay := t.AddDate(0 /* years */, 0 /* months */, 1 /* days */)
maybeNewDay := t.Add(24 * time.Hour)
```

#### Use `time.Time` and `time.Duration` with external systems 与外部系统一起使用 `time.Time` 和 `time.Duration`

Use `time.Duration` and `time.Time` in interactions with external systems when possible. For example:

​	尽可能在与外部系统的交互中使用 `time.Duration` 和 `time.Time` 。例如：

- Command-line flags: [`flag`](https://golang.org/pkg/flag/) supports `time.Duration` via [`time.ParseDuration`](https://golang.org/pkg/time/#ParseDuration)
  命令行标志： `flag` 通过 `time.ParseDuration` 支持 `time.Duration`
- JSON: [`encoding/json`](https://golang.org/pkg/encoding/json/) supports encoding `time.Time` as an [RFC 3339](https://tools.ietf.org/html/rfc3339) string via its [`UnmarshalJSON` method](https://golang.org/pkg/time/#Time.UnmarshalJSON)
  JSON: `encoding/json` 通过其 `UnmarshalJSON` 方法支持将 `time.Time` 编码为 RFC 3339 字符串
- SQL: [`database/sql`](https://golang.org/pkg/database/sql/) supports converting `DATETIME` or `TIMESTAMP` columns into `time.Time` and back if the underlying driver supports it
  SQL: `database/sql` 支持将 `DATETIME` 或 `TIMESTAMP` 列转换为 `time.Time` ，反之亦然（如果底层驱动程序支持）
- YAML: [`gopkg.in/yaml.v2`](https://godoc.org/gopkg.in/yaml.v2) supports `time.Time` as an [RFC 3339](https://tools.ietf.org/html/rfc3339) string, and `time.Duration` via [`time.ParseDuration`](https://golang.org/pkg/time/#ParseDuration).
  YAML: `gopkg.in/yaml.v2` 支持 `time.Time` 作为 RFC 3339 字符串，以及通过 `time.ParseDuration` 支持 `time.Duration` 。

When it is not possible to use `time.Duration` in these interactions, use `int` or `float64` and include the unit in the name of the field.

​	如果在这些交互中无法使用 `time.Duration` ，请使用 `int` 或 `float64` ，并将单位包含在字段名称中。

For example, since `encoding/json` does not support `time.Duration`, the unit is included in the name of the field.

​	例如，由于 `encoding/json` 不支持 `time.Duration` ，因此将单位包含在字段名称中。

```go
// Bad
// {"interval": 2}
type Config struct {
  Interval int `json:"interval"`
}
// Good
// {"intervalMillis": 2000}
type Config struct {
  IntervalMillis int `json:"intervalMillis"`
}
```

When it is not possible to use `time.Time` in these interactions, unless an alternative is agreed upon, use `string` and format timestamps as defined in [RFC 3339](https://tools.ietf.org/html/rfc3339). This format is used by default by [`Time.UnmarshalText`](https://golang.org/pkg/time/#Time.UnmarshalText) and is available for use in `Time.Format` and `time.Parse` via [`time.RFC3339`](https://golang.org/pkg/time/#RFC3339).

​	如果在这些交互中无法使用 `time.Time` ，除非另行达成一致，请使用 `string` ，并将时间戳格式化为 RFC 3339 中定义的格式。此格式默认由 `Time.UnmarshalText` 使用，并且可通过 `time.RFC3339` 在 `Time.Format` 和 `time.Parse` 中使用。

Although this tends to not be a problem in practice, keep in mind that the `"time"` package does not support parsing timestamps with leap seconds ([8728](https://github.com/golang/go/issues/8728)), nor does it account for leap seconds in calculations ([15190](https://github.com/golang/go/issues/15190)). If you compare two instants of time, the difference will not include the leap seconds that may have occurred between those two instants.

​	尽管在实践中这往往不是问题，但请记住， `"time"` 包不支持解析带闰秒的时间戳 ( 8728)，也不在计算中考虑闰秒 ( 15190)。如果您比较两个时间点，则差值不会包括在这两个时间点之间可能发生的闰秒。

### Errors 错误

#### Error Types 错误类型

There are few options for declaring errors. Consider the following before picking the option best suited for your use case.

​	声明错误的方法很少。在选择最适合您的用例的选项之前，请考虑以下几点。

- Does the caller need to match the error so that they can handle it? If yes, we must support the [`errors.Is`](https://golang.org/pkg/errors/#Is) or [`errors.As`](https://golang.org/pkg/errors/#As) functions by declaring a top-level error variable or a custom type.
  调用者是否需要匹配错误以便能够处理错误？如果是，我们必须通过声明顶级错误变量或自定义类型来支持 `errors.Is` 或 `errors.As` 函数。
- Is the error message a static string, or is it a dynamic string that requires contextual information? For the former, we can use [`errors.New`](https://golang.org/pkg/errors/#New), but for the latter we must use [`fmt.Errorf`](https://golang.org/pkg/fmt/#Errorf) or a custom error type.
  错误消息是静态字符串还是需要上下文信息的动态字符串？对于前者，我们可以使用 `errors.New` ，但对于后者，我们必须使用 `fmt.Errorf` 或自定义错误类型。
- Are we propagating a new error returned by a downstream function? If so, see the [section on error wrapping](https://github.com/uber-go/guide/blob/master/style.md#error-wrapping).
  我们是否正在传播下游函数返回的新错误？如果是，请参阅错误包装部分。

| Error matching? 错误匹配？ | Error Message 错误消息 | Guidance 指导                                                |
| -------------------------- | ---------------------- | ------------------------------------------------------------ |
| No                         | static 静态            | [`errors.New`](https://golang.org/pkg/errors/#New)           |
| No                         | dynamic 动态           | [`fmt.Errorf`](https://golang.org/pkg/fmt/#Errorf)           |
| Yes                        | static 静态            | top-level `var` with [`errors.New`](https://golang.org/pkg/errors/#New) 顶级 `var` 与 `errors.New` |
| Yes                        | dynamic 动态           | custom `error` type 自定义 `error` 类型                      |

For example, use [`errors.New`](https://golang.org/pkg/errors/#New) for an error with a static string. Export this error as a variable to support matching it with `errors.Is` if the caller needs to match and handle this error.

​	例如，对于具有静态字符串的错误，请使用 `errors.New` 。将此错误导出为变量，以支持与 `errors.Is` 匹配，如果调用方需要匹配并处理此错误。

```go
// No error matching
// package foo

func Open() error {
  return errors.New("could not open")
}

// package bar

if err := foo.Open(); err != nil {
  // Can't handle the error.
  panic("unknown error")
}
// Error matching
// package foo

var ErrCouldNotOpen = errors.New("could not open")

func Open() error {
  return ErrCouldNotOpen
}

// package bar

if err := foo.Open(); err != nil {
  if errors.Is(err, foo.ErrCouldNotOpen) {
    // handle the error
  } else {
    panic("unknown error")
  }
}
```

For an error with a dynamic string, use [`fmt.Errorf`](https://golang.org/pkg/fmt/#Errorf) if the caller does not need to match it, and a custom `error` if the caller does need to match it.

​	对于具有动态字符串的错误，如果调用方不需要匹配它，请使用 `fmt.Errorf` ；如果调用方确实需要匹配它，请使用自定义 `error` 。

```go
// No error matching
// package foo

func Open(file string) error {
  return fmt.Errorf("file %q not found", file)
}

// package bar

if err := foo.Open("testfile.txt"); err != nil {
  // Can't handle the error.
  panic("unknown error")
}
// Error matching
// package foo

type NotFoundError struct {
  File string
}

func (e *NotFoundError) Error() string {
  return fmt.Sprintf("file %q not found", e.File)
}

func Open(file string) error {
  return &NotFoundError{File: file}
}


// package bar

if err := foo.Open("testfile.txt"); err != nil {
  var notFound *NotFoundError
  if errors.As(err, &notFound) {
    // handle the error
  } else {
    panic("unknown error")
  }
}
```

Note that if you export error variables or types from a package, they will become part of the public API of the package.

​	请注意，如果您从包中导出错误变量或类型，它们将成为包的公共 API 的一部分。

#### Error Wrapping 错误包装

There are three main options for propagating errors if a call fails:

​	如果调用失败，有三种主要选项可用于传播错误：

- return the original error as-is
  按原样返回原始错误
- add context with `fmt.Errorf` and the `%w` verb
  使用 `fmt.Errorf` 和 `%w` 动词添加上下文
- add context with `fmt.Errorf` and the `%v` verb
  使用 `fmt.Errorf` 和 `%v` 动词添加上下文

Return the original error as-is if there is no additional context to add. This maintains the original error type and message. This is well suited for cases when the underlying error message has sufficient information to track down where it came from.

​	如果没有要添加的其他上下文，请按原样返回原始错误。这将保留原始错误类型和消息。当基础错误消息有足够的信息来跟踪其来源时，这非常适合。

Otherwise, add context to the error message where possible so that instead of a vague error such as “connection refused”, you get more useful errors such as “call service foo: connection refused”.

​	否则，尽可能为错误消息添加上下文，以便您获得“连接被拒绝”等更有用的错误，而不是“连接被拒绝”等模糊的错误。

Use `fmt.Errorf` to add context to your errors, picking between the `%w` or `%v` verbs based on whether the caller should be able to match and extract the underlying cause.

​	使用 `fmt.Errorf` 为错误添加上下文，根据调用者是否能够匹配并提取根本原因，在 `%w` 或 `%v` 动词之间进行选择。

- Use `%w` if the caller should have access to the underlying error. This is a good default for most wrapped errors, but be aware that callers may begin to rely on this behavior. So for cases where the wrapped error is a known `var` or type, document and test it as part of your function’s contract.
  如果调用者应该有权访问基础错误，请使用 `%w` 。对于大多数包装错误来说，这是一个很好的默认值，但请注意，调用者可能会开始依赖此行为。因此，对于包装错误是已知的 `var` 或类型的情况，请将其作为函数契约的一部分进行记录和测试。
- Use `%v` to obfuscate the underlying error. Callers will be unable to match it, but you can switch to `%w` in the future if needed.
  使用 `%v` 混淆基础错误。调用者将无法匹配它，但您可以在需要时在将来切换到 `%w` 。

When adding context to returned errors, keep the context succinct by avoiding phrases like “failed to”, which state the obvious and pile up as the error percolates up through the stack:

​	在为返回的错误添加上下文时，请避免使用“失败”等短语来保持上下文简洁，这些短语说明了显而易见的事情，并且随着错误在堆栈中渗透而堆积起来：

```go
// Bad
s, err := store.New()
if err != nil {
    return fmt.Errorf(
        "failed to create new store: %w", err)
}

// failed to x: failed to y: failed to create new store: the error
// Good
s, err := store.New()
if err != nil {
    return fmt.Errorf(
        "new store: %w", err)
}

// x: y: new store: the error
```

However once the error is sent to another system, it should be clear the message is an error (e.g. an `err` tag or “Failed” prefix in logs).

​	但是，一旦错误被发送到另一个系统，就应该清楚消息是一个错误（例如，日志中的 `err` 标记或“失败”前缀）。

See also [Don’t just check errors, handle them gracefully](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully).

​	另请参阅不要仅仅检查错误，还要妥善处理它们。

#### Error Naming 错误命名

For error values stored as global variables, use the prefix `Err` or `err` depending on whether they’re exported. This guidance supersedes the [Prefix Unexported Globals with _](https://github.com/uber-go/guide/blob/master/style.md#prefix-unexported-globals-with-_).

​	对于存储为全局变量的错误值，请使用前缀 `Err` 或 `err` ，具体取决于它们是否已导出。此指南取代了使用 _ 为未导出的全局变量添加前缀。

```go
var (
  // The following two errors are exported
  // so that users of this package can match them
  // with errors.Is.

  ErrBrokenLink = errors.New("link is broken")
  ErrCouldNotOpen = errors.New("could not open")

  // This error is not exported because
  // we don't want to make it part of our public API.
  // We may still use it inside the package
  // with errors.Is.

  errNotFound = errors.New("not found")
)
```

For custom error types, use the suffix `Error` instead.

​	对于自定义错误类型，请改用后缀 `Error` 。

```
// Similarly, this error is exported
// so that users of this package can match it
// with errors.As.

type NotFoundError struct {
  File string
}

func (e *NotFoundError) Error() string {
  return fmt.Sprintf("file %q not found", e.File)
}

// And this error is not exported because
// we don't want to make it part of the public API.
// We can still use it inside the package
// with errors.As.

type resolveError struct {
  Path string
}

func (e *resolveError) Error() string {
  return fmt.Sprintf("resolve %q", e.Path)
}
```

#### Handle Errors Once 一次处理错误

When a caller receives an error from a callee, it can handle it in a variety of different ways depending on what it knows about the error.

​	当调用方从被调用方收到错误时，它可以根据对错误的了解以各种不同的方式处理错误。

These include, but not are limited to:

​	这些包括但不限于：

- if the callee contract defines specific errors, matching the error with `errors.Is` or `errors.As` and handling the branches differently
  如果被调用方合约定义了特定错误，则使用 `errors.Is` 或 `errors.As` 匹配错误并以不同的方式处理分支
- if the error is recoverable, logging the error and degrading gracefully
  如果错误是可恢复的，则记录错误并优雅地降级
- if the error represents a domain-specific failure condition, returning a well-defined error
  如果错误表示特定于域的故障条件，则返回定义明确的错误
- returning the error, either [wrapped](https://github.com/uber-go/guide/blob/master/style.md#error-wrapping) or verbatim
  返回错误，无论是包装的还是逐字的

Regardless of how the caller handles the error, it should typically handle each error only once. The caller should not, for example, log the error and then return it, because *its* callers may handle the error as well.

​	无论调用者如何处理错误，通常都应该只处理一次每个错误。例如，调用者不应记录错误然后返回错误，因为其调用者也可能处理错误。

For example, consider the following cases:

​	例如，考虑以下情况：

**Bad**: Log the error and return it. Callers further up the stack will likely take a similar action with the error. Doing so causing a lot of noise in the application logs for little value.

​	错误：记录错误并返回错误。堆栈上方的调用者可能会对错误采取类似的操作。这样做会导致应用程序日志中产生大量噪音，而价值却很小。

```go
// Bad
u, err := getUser(id)
if err != nil {
  // BAD: See description
  log.Printf("Could not get user %q: %v", id, err)
  return err
}
```

**Good**: Wrap the error and return it. Callers further up the stack will handle the error. Use of `%w` ensures they can match the error with `errors.Is` or `errors.As` if relevant.

​	正确：包装错误并返回错误。堆栈上方的调用者将处理错误。使用 `%w` 确保它们可以将错误与 `errors.Is` 或 `errors.As` （如果相关）匹配。

```go
// Good
u, err := getUser(id)
if err != nil {
  return fmt.Errorf("get user %q: %w", id, err)
}
```

**Good**: Log the error and degrade gracefully. If the operation isn’t strictly necessary, we can provide a degraded but unbroken experience by recovering from it.

​	正确：记录错误并优雅地降级。如果操作不是严格必需的，我们可以通过从中恢复来提供降级但未中断的体验。

```go
// Good
if err := emitMetrics(); err != nil {
  // Failure to write metrics should not
  // break the application.
  log.Printf("Could not emit metrics: %v", err)
}
```

**Good**: Match the error and degrade gracefully. If the callee defines a specific error in its contract, and the failure is recoverable, match on that error case and degrade gracefully. For all other cases, wrap the error and return it. Callers further up the stack will handle other errors.

​	正确：匹配错误并优雅地降级。如果被调用者在其契约中定义了特定错误，并且该故障是可恢复的，则匹配该错误情况并优雅地降级。对于所有其他情况，请包装错误并返回错误。堆栈上方的调用者将处理其他错误。

```go
// Good
tz, err := getUserTimeZone(id)
if err != nil {
  if errors.Is(err, ErrUserNotFound) {
    // User doesn't exist. Use UTC.
    tz = time.UTC
  } else {
    return fmt.Errorf("get user %q: %w", id, err)
  }
}
```

### Handle Type Assertion Failures 处理类型断言失败

The single return value form of a [type assertion](https://golang.org/ref/spec#Type_assertions) will panic on an incorrect type. Therefore, always use the “comma ok” idiom.

​	类型断言的单一返回值形式会在类型不正确时引发恐慌。因此，始终使用“逗号 ok”习惯用法。

```go
// Bad
t := i.(string)
// Good
t, ok := i.(string)
if !ok {
  // handle the error gracefully
}
```

### Don’t Panic 不要恐慌

Code running in production must avoid panics. Panics are a major source of [cascading failures](https://en.wikipedia.org/wiki/Cascading_failure). If an error occurs, the function must return an error and allow the caller to decide how to handle it.

​	在生产环境中运行的代码必须避免恐慌。恐慌是级联故障的主要来源。如果发生错误，函数必须返回错误并允许调用者决定如何处理它。

```go
// Bad
func run(args []string) {
  if len(args) == 0 {
    panic("an argument is required")
  }
  // ...
}

func main() {
  run(os.Args[1:])
}
// Good
func run(args []string) error {
  if len(args) == 0 {
    return errors.New("an argument is required")
  }
  // ...
  return nil
}

func main() {
  if err := run(os.Args[1:]); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
```

Panic/recover is not an error handling strategy. A program must panic only when something irrecoverable happens such as a nil dereference. An exception to this is program initialization: bad things at program startup that should abort the program may cause panic.

​	恐慌/恢复不是错误处理策略。程序仅在发生无法恢复的情况（例如 nil 解引用）时才应恐慌。对此的一个例外是程序初始化：程序启动时应中止程序的错误可能会导致恐慌。

```go
var _statusTemplate = template.Must(template.New("name").Parse("_statusHTML"))
```

Even in tests, prefer `t.Fatal` or `t.FailNow` over panics to ensure that the test is marked as failed.

​	即使在测试中，也更喜欢 `t.Fatal` 或 `t.FailNow` ，而不是恐慌，以确保将测试标记为失败。

```go
// Bad
// func TestFoo(t *testing.T)

f, err := os.CreateTemp("", "test")
if err != nil {
  panic("failed to set up test")
}
// Good
// func TestFoo(t *testing.T)

f, err := os.CreateTemp("", "test")
if err != nil {
  t.Fatal("failed to set up test")
}
```

### Use go.uber.org/atomic 使用 go.uber.org/atomic

Atomic operations with the [sync/atomic](https://golang.org/pkg/sync/atomic/) package operate on the raw types (`int32`, `int64`, etc.) so it is easy to forget to use the atomic operation to read or modify the variables.

​	使用 sync/atomic 包进行原子操作时，操作的是原始类型（ `int32` 、 `int64` 等），因此很容易忘记使用原子操作来读取或修改变量。

[go.uber.org/atomic](https://godoc.org/go.uber.org/atomic) adds type safety to these operations by hiding the underlying type. Additionally, it includes a convenient `atomic.Bool` type.

​	go.uber.org/atomic 通过隐藏底层类型，为这些操作添加了类型安全性。此外，它还包括一个方便的 `atomic.Bool` 类型。

```go
// Bad
type foo struct {
  running int32  // atomic
}

func (f* foo) start() {
  if atomic.SwapInt32(&f.running, 1) == 1 {
     // already running…
     return
  }
  // start the Foo
}

func (f *foo) isRunning() bool {
  return f.running == 1  // race!
}
// Good
type foo struct {
  running atomic.Bool
}

func (f *foo) start() {
  if f.running.Swap(true) {
     // already running…
     return
  }
  // start the Foo
}

func (f *foo) isRunning() bool {
  return f.running.Load()
}
```

### Avoid Mutable Globals 避免可变全局变量

Avoid mutating global variables, instead opting for dependency injection. This applies to function pointers as well as other kinds of values.

​	避免改变全局变量，而选择依赖注入。这适用于函数指针以及其他类型的值。

```go
// Bad
// sign.go

var _timeNow = time.Now

func sign(msg string) string {
  now := _timeNow()
  return signWithTime(msg, now)
}

// sign_test.go

func TestSign(t *testing.T) {
  oldTimeNow := _timeNow
  _timeNow = func() time.Time {
    return someFixedTime
  }
  defer func() { _timeNow = oldTimeNow }()

  assert.Equal(t, want, sign(give))
}
// Good
// sign.go

type signer struct {
  now func() time.Time
}

func newSigner() *signer {
  return &signer{
    now: time.Now,
  }
}

func (s *signer) Sign(msg string) string {
  now := s.now()
  return signWithTime(msg, now)
}

// sign_test.go

func TestSigner(t *testing.T) {
  s := newSigner()
  s.now = func() time.Time {
    return someFixedTime
  }

  assert.Equal(t, want, s.Sign(give))
}
```

### Avoid Embedding Types in Public Structs 避免在公共结构中嵌入类型

These embedded types leak implementation details, inhibit type evolution, and obscure documentation.

​	这些嵌入的类型会泄露实现细节，抑制类型演化，并模糊文档。

Assuming you have implemented a variety of list types using a shared `AbstractList`, avoid embedding the `AbstractList` in your concrete list implementations. Instead, hand-write only the methods to your concrete list that will delegate to the abstract list.

​	假设您使用共享的 `AbstractList` 实现各种列表类型，请避免在具体列表实现中嵌入 `AbstractList` 。相反，只手写将委托给抽象列表的具体列表的方法。

```go
type AbstractList struct {}

// Add adds an entity to the list.
func (l *AbstractList) Add(e Entity) {
  // ...
}

// Remove removes an entity from the list.
func (l *AbstractList) Remove(e Entity) {
  // ...
}
// Bad
// ConcreteList is a list of entities.
type ConcreteList struct {
  *AbstractList
}
// Good
// ConcreteList is a list of entities.
type ConcreteList struct {
  list *AbstractList
}

// Add adds an entity to the list.
func (l *ConcreteList) Add(e Entity) {
  l.list.Add(e)
}

// Remove removes an entity from the list.
func (l *ConcreteList) Remove(e Entity) {
  l.list.Remove(e)
}
```

Go allows [type embedding](https://golang.org/doc/effective_go.html#embedding) as a compromise between inheritance and composition. The outer type gets implicit copies of the embedded type’s methods. These methods, by default, delegate to the same method of the embedded instance.

​	Go 允许类型嵌入作为继承和组合之间的折衷方案。外部类型会隐式获取嵌入类型方法的副本。默认情况下，这些方法委托给嵌入实例的相同方法。

The struct also gains a field by the same name as the type. So, if the embedded type is public, the field is public. To maintain backward compatibility, every future version of the outer type must keep the embedded type.

​	该结构还获得一个与类型同名的字段。因此，如果嵌入类型是公共的，则该字段是公共的。为了保持向后兼容性，外部类型的每个未来版本都必须保留嵌入类型。

An embedded type is rarely necessary. It is a convenience that helps you avoid writing tedious delegate methods.

​	嵌入类型很少有必要。这是一种帮助您避免编写繁琐的委托方法的便利。

Even embedding a compatible AbstractList *interface*, instead of the struct, would offer the developer more flexibility to change in the future, but still leak the detail that the concrete lists use an abstract implementation.

​	即使嵌入兼容的 AbstractList 接口（而不是结构体），也会为开发人员提供更大的灵活性以在将来进行更改，但仍然会泄露具体列表使用抽象实现的细节。

```go
// Bad
// AbstractList is a generalized implementation
// for various kinds of lists of entities.
type AbstractList interface {
  Add(Entity)
  Remove(Entity)
}

// ConcreteList is a list of entities.
type ConcreteList struct {
  AbstractList
}
// Good
// ConcreteList is a list of entities.
type ConcreteList struct {
  list AbstractList
}

// Add adds an entity to the list.
func (l *ConcreteList) Add(e Entity) {
  l.list.Add(e)
}

// Remove removes an entity from the list.
func (l *ConcreteList) Remove(e Entity) {
  l.list.Remove(e)
}
```

Either with an embedded struct or an embedded interface, the embedded type places limits on the evolution of the type.

​	无论使用嵌入式结构体还是嵌入式接口，嵌入式类型都会限制类型的演变。

- Adding methods to an embedded interface is a breaking change.
  向嵌入式接口添加方法是一种重大更改。
- Removing methods from an embedded struct is a breaking change.
  从嵌入式结构体中删除方法是一种重大更改。
- Removing the embedded type is a breaking change.
  删除嵌入式类型是一种重大更改。
- Replacing the embedded type, even with an alternative that satisfies the same interface, is a breaking change.
  替换嵌入式类型（即使使用满足相同接口的替代类型）也是一种重大更改。

Although writing these delegate methods is tedious, the additional effort hides an implementation detail, leaves more opportunities for change, and also eliminates indirection for discovering the full List interface in documentation.

​	尽管编写这些委托方法很繁琐，但额外的努力隐藏了实现细节，留出了更多更改机会，还消除了在文档中发现完整 List 接口的间接性。

### Avoid Using Built-In Names 避免使用内置名称

The Go [language specification](https://golang.org/ref/spec) outlines several built-in, [predeclared identifiers](https://golang.org/ref/spec#Predeclared_identifiers) that should not be used as names within Go programs.

​	Go 语言规范概述了几个内置的、预声明的标识符，这些标识符不应在 Go 程序中用作名称。

Depending on context, reusing these identifiers as names will either shadow the original within the current lexical scope (and any nested scopes) or make affected code confusing. In the best case, the compiler will complain; in the worst case, such code may introduce latent, hard-to-grep bugs.

​	根据上下文，将这些标识符重新用作名称，要么在当前词法作用域（和任何嵌套作用域）内隐藏原始标识符，要么使受影响的代码令人困惑。在最好的情况下，编译器会发出警告；在最坏的情况下，此类代码可能会引入潜在的、难以 grep 的错误。

```go
// Bad
var error string
// `error` shadows the builtin

// or

func handleErrorMessage(error string) {
    // `error` shadows the builtin
}



type Foo struct {
    // While these fields technically don't
    // constitute shadowing, grepping for
    // `error` or `string` strings is now
    // ambiguous.
    error  error
    string string
}

func (f Foo) Error() error {
    // `error` and `f.error` are
    // visually similar
    return f.error
}

func (f Foo) String() string {
    // `string` and `f.string` are
    // visually similar
    return f.string
}
// Good
var errorMessage string
// `error` refers to the builtin

// or

func handleErrorMessage(msg string) {
    // `error` refers to the builtin
}



type Foo struct {
    // `error` and `string` strings are
    // now unambiguous.
    err error
    str string
}

func (f Foo) Error() error {
    return f.err
}

func (f Foo) String() string {
    return f.str
}
```

Note that the compiler will not generate errors when using predeclared identifiers, but tools such as `go vet` should correctly point out these and other cases of shadowing.

​	请注意，在使用预声明标识符时，编译器不会生成错误，但诸如 `go vet` 之类的工具应正确指出这些和其他隐藏情况。

### Avoid `init()` 避免 `init()`

Avoid `init()` where possible. When `init()` is unavoidable or desirable, code should attempt to:

​	尽可能避免 `init()` 。当 `init()` 不可避免或合乎需要时，代码应尝试：

1. Be completely deterministic, regardless of program environment or invocation.
   完全确定性，无论程序环境或调用如何。
2. Avoid depending on the ordering or side-effects of other `init()` functions. While `init()` ordering is well-known, code can change, and thus relationships between `init()` functions can make code brittle and error-prone.
   避免依赖其他 `init()` 函数的顺序或副作用。虽然 `init()` 顺序是众所周知的，但代码可能会发生变化，因此 `init()` 函数之间的关系可能会使代码变得脆弱且容易出错。
3. Avoid accessing or manipulating global or environment state, such as machine information, environment variables, working directory, program arguments/inputs, etc.
   避免访问或操作全局或环境状态，例如机器信息、环境变量、工作目录、程序参数/输入等。
4. Avoid I/O, including both filesystem, network, and system calls.
   避免 I/O，包括文件系统、网络和系统调用。

Code that cannot satisfy these requirements likely belongs as a helper to be called as part of `main()` (or elsewhere in a program’s lifecycle), or be written as part of `main()` itself. In particular, libraries that are intended to be used by other programs should take special care to be completely deterministic and not perform “init magic”.

​	无法满足这些要求的代码可能属于作为 `main()` （或程序生命周期中的其他位置）的一部分被调用的帮助程序，或者作为 `main()` 本身的一部分编写。特别是，旨在供其他程序使用的库应特别注意完全确定性，并且不执行“init magic”。

```go
// Bad
type Foo struct {
    // ...
}

var _defaultFoo Foo

func init() {
    _defaultFoo = Foo{
        // ...
    }
}


type Config struct {
    // ...
}

var _config Config

func init() {
    // Bad: based on current directory
    cwd, _ := os.Getwd()

    // Bad: I/O
    raw, _ := os.ReadFile(
        path.Join(cwd, "config", "config.yaml"),
    )

    yaml.Unmarshal(raw, &_config)
}
// Good
var _defaultFoo = Foo{
    // ...
}

// or, better, for testability:

var _defaultFoo = defaultFoo()

func defaultFoo() Foo {
    return Foo{
        // ...
    }
}



type Config struct {
    // ...
}

func loadConfig() Config {
    cwd, err := os.Getwd()
    // handle err

    raw, err := os.ReadFile(
        path.Join(cwd, "config", "config.yaml"),
    )
    // handle err

    var config Config
    yaml.Unmarshal(raw, &config)

    return config
}
```

Considering the above, some situations in which `init()` may be preferable or necessary might include:

​	考虑到上述情况，可能更可取或必要使用 `init()` 的一些情况可能包括：

- Complex expressions that cannot be represented as single assignments.
  无法表示为单个赋值的复杂表达式。
- Pluggable hooks, such as `database/sql` dialects, encoding type registries, etc.
  可插入的挂钩，例如 `database/sql` 方言、编码类型注册表等。
- Optimizations to [Google Cloud Functions](https://cloud.google.com/functions/docs/bestpractices/tips#use_global_variables_to_reuse_objects_in_future_invocations) and other forms of deterministic precomputation.
  对 Google Cloud Functions 和其他形式的确定性预计算的优化。

### Exit in Main 在 Main 中退出

Go programs use [`os.Exit`](https://golang.org/pkg/os/#Exit) or [`log.Fatal*`](https://golang.org/pkg/log/#Fatal) to exit immediately. (Panicking is not a good way to exit programs, please [don’t panic](https://github.com/uber-go/guide/blob/master/style.md#dont-panic).)

​	Go 程序使用 `os.Exit` 或 `log.Fatal*` 立即退出。（引发恐慌不是退出程序的好方法，请不要惊慌。）

Call one of `os.Exit` or `log.Fatal*` **only in `main()`**. All other functions should return errors to signal failure.

​	仅在 `main()` 中调用 `os.Exit` 或 `log.Fatal*` 之一。所有其他函数都应返回错误以发出失败信号。

```go
// Bad
func main() {
  body := readFile(path)
  fmt.Println(body)
}

func readFile(path string) string {
  f, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }

  b, err := io.ReadAll(f)
  if err != nil {
    log.Fatal(err)
  }

  return string(b)
}
// Good
func main() {
  body, err := readFile(path)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(body)
}

func readFile(path string) (string, error) {
  f, err := os.Open(path)
  if err != nil {
    return "", err
  }

  b, err := io.ReadAll(f)
  if err != nil {
    return "", err
  }

  return string(b), nil
}
```

Rationale: Programs with multiple functions that exit present a few issues:

​	基本原理：具有多个退出函数的程序会带来一些问题：

- Non-obvious control flow: Any function can exit the program so it becomes difficult to reason about the control flow.
  不明显的控制流：任何函数都可以退出程序，因此很难推理控制流。
- Difficult to test: A function that exits the program will also exit the test calling it. This makes the function difficult to test and introduces risk of skipping other tests that have not yet been run by `go test`.
  难以测试：退出程序的函数也会退出调用它的测试。这使得函数难以测试，并引入跳过尚未由 `go test` 运行的其他测试的风险。
- Skipped cleanup: When a function exits the program, it skips function calls enqueued with `defer` statements. This adds risk of skipping important cleanup tasks.
  跳过的清理：当函数退出程序时，它会跳过使用 `defer` 语句排队的函数调用。这会增加跳过重要清理任务的风险。

#### Exit Once 退出一次

If possible, prefer to call `os.Exit` or `log.Fatal` **at most once** in your `main()`. If there are multiple error scenarios that halt program execution, put that logic under a separate function and return errors from it.

​	如果可能，最好在 `main()` 中最多调用 `os.Exit` 或 `log.Fatal` 一次。如果有多个导致程序执行停止的错误场景，请将该逻辑放在单独的函数中，并从中返回错误。

This has the effect of shortening your `main()` function and putting all key business logic into a separate, testable function.

​	这会缩短 `main()` 函数，并将所有关键业务逻辑放入单独的可测试函数中。

```go
// Bad
package main

func main() {
  args := os.Args[1:]
  if len(args) != 1 {
    log.Fatal("missing file")
  }
  name := args[0]

  f, err := os.Open(name)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  // If we call log.Fatal after this line,
  // f.Close will not be called.

  b, err := io.ReadAll(f)
  if err != nil {
    log.Fatal(err)
  }

  // ...
}
// Good
package main

func main() {
  if err := run(); err != nil {
    log.Fatal(err)
  }
}

func run() error {
  args := os.Args[1:]
  if len(args) != 1 {
    return errors.New("missing file")
  }
  name := args[0]

  f, err := os.Open(name)
  if err != nil {
    return err
  }
  defer f.Close()

  b, err := io.ReadAll(f)
  if err != nil {
    return err
  }

  // ...
}
```

The example above uses `log.Fatal`, but the guidance also applies to `os.Exit` or any library code that calls `os.Exit`.

​	上面的示例使用 `log.Fatal` ，但该指南也适用于 `os.Exit` 或任何调用 `os.Exit` 的库代码。

```go
func main() {
  if err := run(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
```

You may alter the signature of `run()` to fit your needs. For example, if your program must exit with specific exit codes for failures, `run()` may return the exit code instead of an error. This allows unit tests to verify this behavior directly as well.

​	您可以更改 `run()` 的签名以满足您的需求。例如，如果您的程序必须使用特定的退出代码退出以表示失败，则 `run()` 可以返回退出代码而不是错误。这也允许单元测试直接验证此行为。

```go
func main() {
  os.Exit(run(args))
}

func run() (exitCode int) {
  // ...
}
```

More generally, note that the `run()` function used in these examples is not intended to be prescriptive. There’s flexibility in the name, signature, and setup of the `run()` function. Among other things, you may:

​	更一般地说，请注意，这些示例中使用的 `run()` 函数并非旨在具有规定性。在 `run()` 函数的名称、签名和设置中存在灵活性。除其他事项外，您可能：

- accept unparsed command line arguments (e.g., `run(os.Args[1:])`)
  接受未解析的命令行参数（例如， `run(os.Args[1:])` ）
- parse command line arguments in `main()` and pass them onto `run`
  在 `main()` 中解析命令行参数并将其传递给 `run`
- use a custom error type to carry the exit code back to `main()`
  使用自定义错误类型将退出代码传回 `main()`
- put business logic in a different layer of abstraction from `package main`
  将业务逻辑放在与 `package main` 不同的抽象层中

This guidance only requires that there’s a single place in your `main()` responsible for actually exiting the process.

​	此指南仅要求在您的 `main()` 中有一个负责实际退出进程的位置。

### Use field tags in marshaled structs 在编组结构中使用字段标记

Any struct field that is marshaled into JSON, YAML, or other formats that support tag-based field naming should be annotated with the relevant tag.

​	任何编组到 JSON、YAML 或其他支持基于标记的字段命名的格式的结构字段都应使用相关标记进行注释。

```go
// Bad
type Stock struct {
  Price int
  Name  string
}

bytes, err := json.Marshal(Stock{
  Price: 137,
  Name:  "UBER",
})
// Good
type Stock struct {
  Price int    `json:"price"`
  Name  string `json:"name"`
  // Safe to rename Name to Symbol.
}

bytes, err := json.Marshal(Stock{
  Price: 137,
  Name:  "UBER",
})
```

Rationale: The serialized form of the structure is a contract between different systems. Changes to the structure of the serialized form–including field names–break this contract. Specifying field names inside tags makes the contract explicit, and it guards against accidentally breaking the contract by refactoring or renaming fields.

​	基本原理：结构的序列化形式是不同系统之间的契约。对序列化形式的结构（包括字段名称）进行更改会破坏此契约。在标记内指定字段名称使契约显式化，并且它可以防止通过重构或重命名字段来意外破坏契约。

### Don’t fire-and-forget goroutines 不要使用即用即弃的 goroutine

Goroutines are lightweight, but they’re not free: at minimum, they cost memory for their stack and CPU to be scheduled. While these costs are small for typical uses of goroutines, they can cause significant performance issues when spawned in large numbers without controlled lifetimes. Goroutines with unmanaged lifetimes can also cause other issues like preventing unused objects from being garbage collected and holding onto resources that are otherwise no longer used.

​	Goroutine 很轻量，但它们不是免费的：至少，它们需要内存来存储它们的堆栈和 CPU 来进行调度。虽然对于 Goroutine 的典型用途来说这些成本很小，但当大量生成 Goroutine 且没有受控的生命周期时，它们可能会导致严重的性能问题。具有不受管理的生命周期的 Goroutine 还会导致其他问题，例如阻止未使用的对象被垃圾回收，以及持有不再使用的资源。

Therefore, do not leak goroutines in production code. Use [go.uber.org/goleak](https://pkg.go.dev/go.uber.org/goleak) to test for goroutine leaks inside packages that may spawn goroutines.

​	因此，不要在生产代码中泄漏 Goroutine。使用 go.uber.org/goleak 来测试可能生成 Goroutine 的包中的 Goroutine 泄漏。

In general, every goroutine:

​	通常，每个 Goroutine：

- must have a predictable time at which it will stop running; or
  必须具有可预测的停止运行时间；或者
- there must be a way to signal to the goroutine that it should stop
  必须有一种方法向 Goroutine 发出信号，指示它应该停止

In both cases, there must be a way code to block and wait for the goroutine to finish.

​	在这两种情况下，都必须有一种方法来阻塞代码并等待 Goroutine 完成。

For example:

​	例如：

```go
// Bad
go func() {
  for {
    flush()
    time.Sleep(delay)
  }
}()
```

There’s no way to stop this goroutine. This will run until the application exits.

​	没有办法停止此 Goroutine。这将一直运行，直到应用程序退出。

```go
// Good
var (
  stop = make(chan struct{}) // tells the goroutine to stop
  done = make(chan struct{}) // tells us that the goroutine exited
)
go func() {
  defer close(done)

  ticker := time.NewTicker(delay)
  defer ticker.Stop()
  for {
    select {
    case <-ticker.C:
      flush()
    case <-stop:
      return
    }
  }
}()

// Elsewhere...
close(stop)  // signal the goroutine to stop
<-done       // and wait for it to exit
```

This goroutine can be stopped with close(stop), and we can wait for it to exit with <-done.

​	可以使用 close(stop) 停止此 Goroutine，我们可以使用 <-done 等待它退出。

#### Wait for goroutines to exit 等待 Goroutine 退出

Given a goroutine spawned by the system, there must be a way to wait for the goroutine to exit. There are two popular ways to do this:

​	如果系统生成一个协程，那么肯定有办法等待协程退出。有两种流行的方法可以做到这一点：

- Use a `sync.WaitGroup`. Do this if there are multiple goroutines that you want to wait for

  ​	使用 `sync.WaitGroup` 。如果要等待多个协程，请执行此操作

  ```
  var wg sync.WaitGroup
  for i := 0; i < N; i++ {
    wg.Add(1)
    go func() {
      defer wg.Done()
      // ...
    }()
  }
  
  // To wait for all to finish:
  wg.Wait()
  ```

- Add another `chan struct{}` that the goroutine closes when it’s done. Do this if there’s only one goroutine.

  ​	添加另一个 `chan struct{}` ，协程完成后将其关闭。如果只有一个协程，请执行此操作。

  ```
  done := make(chan struct{})
  go func() {
    defer close(done)
    // ...
  }()
  
  // To wait for the goroutine to finish:
  <-done
  ```

#### No goroutines in `init()` `init()` 中没有协程

`init()` functions should not spawn goroutines. See also [Avoid init()](https://github.com/uber-go/guide/blob/master/style.md#avoid-init).

​	 `init()` 函数不应生成协程。另请参阅避免使用 init()。

If a package has need of a background goroutine, it must expose an object that is responsible for managing a goroutine’s lifetime. The object must provide a method (`Close`, `Stop`, `Shutdown`, etc) that signals the background goroutine to stop, and waits for it to exit.

​	如果某个包需要后台协程，则它必须公开一个负责管理协程生命周期的对象。该对象必须提供一个方法（ `Close` 、 `Stop` 、 `Shutdown` 等）来指示后台协程停止，并等待其退出。

```go
// Bad
func init() {
  go doWork()
}

func doWork() {
  for {
    // ...
  }
}
```

Spawns a background goroutine unconditionally when the user exports this package. The user has no control over the goroutine or a means of stopping it.

​	当用户导出此包时，无条件生成后台协程。用户无法控制协程或停止协程的方法。

```go
// Good
type Worker struct{ /* ... */ }

func NewWorker(...) *Worker {
  w := &Worker{
    stop: make(chan struct{}),
    done: make(chan struct{}),
    // ...
  }
  go w.doWork()
}

func (w *Worker) doWork() {
  defer close(w.done)
  for {
    // ...
    case <-w.stop:
      return
  }
}

// Shutdown tells the worker to stop
// and waits until it has finished.
func (w *Worker) Shutdown() {
  close(w.stop)
  <-w.done
}
```

Spawns the worker only if the user requests it. Provides a means of shutting down the worker so that the user can free up resources used by the worker.

​	仅在用户请求时生成工作进程。提供关闭工作进程的方法，以便用户可以释放工作进程使用的资源。

Note that you should use `WaitGroup`s if the worker manages multiple goroutines. See [Wait for goroutines to exit](https://github.com/uber-go/guide/blob/master/style.md#wait-for-goroutines-to-exit).

​	请注意，如果工作进程管理多个协程，则应使用 `WaitGroup` 。请参阅等待协程退出。

## Performance 性能

Performance-specific guidelines apply only to the hot path.

​	性能特定准则仅适用于热路径。

### Prefer strconv over fmt 优先使用 strconv 而非 fmt

When converting primitives to/from strings, `strconv` is faster than `fmt`.

​	将基本类型转换为字符串或从字符串转换基本类型时， `strconv` 比 `fmt` 更快。

```go
// Bad
for i := 0; i < b.N; i++ {
  s := fmt.Sprint(rand.Int())
}

// BenchmarkFmtSprint-4    143 ns/op    2 allocs/op
// Good
for i := 0; i < b.N; i++ {
  s := strconv.Itoa(rand.Int())
}

// BenchmarkStrconv-4    64.2 ns/op    1 allocs/op
```

### Avoid repeated string-to-byte conversions 避免重复的字符串到字节的转换

Do not create byte slices from a fixed string repeatedly. Instead, perform the conversion once and capture the result.

​	不要反复从固定字符串创建字节切片。相反，执行一次转换并捕获结果。

```go
// Bad
for i := 0; i < b.N; i++ {
  w.Write([]byte("Hello world"))
}

// BenchmarkBad-4   50000000   22.2 ns/op
// Good
data := []byte("Hello world")
for i := 0; i < b.N; i++ {
  w.Write(data)
}

// BenchmarkGood-4  500000000   3.25 ns/op
```

### Prefer Specifying Container Capacity 优先指定容器容量

Specify container capacity where possible in order to allocate memory for the container up front. This minimizes subsequent allocations (by copying and resizing of the container) as elements are added.

​	尽可能指定容器容量，以便预先为容器分配内存。随着元素的添加，这可以最大程度地减少后续分配（通过复制和调整容器大小）。

#### Specifying Map Capacity Hints 指定映射容量提示

Where possible, provide capacity hints when initializing maps with `make()`.

​	尽可能在使用 `make()` 初始化映射时提供容量提示。

```
make(map[T1]T2, hint)
```

Providing a capacity hint to `make()` tries to right-size the map at initialization time, which reduces the need for growing the map and allocations as elements are added to the map.

​	向 `make()` 提供容量提示会尝试在初始化时调整映射大小，从而减少随着元素添加到映射中而需要增加映射和分配的情况。

Note that, unlike slices, map capacity hints do not guarantee complete, preemptive allocation, but are used to approximate the number of hashmap buckets required. Consequently, allocations may still occur when adding elements to the map, even up to the specified capacity.

​	请注意，与切片不同，映射容量提示不保证完全的抢占式分配，但用于估算所需的哈希映射存储桶数。因此，即使达到指定容量，在向映射中添加元素时仍可能发生分配。

```go
// Bad
m := make(map[string]os.FileInfo)

files, _ := os.ReadDir("./files")
for _, f := range files {
    m[f.Name()] = f
}
```

m is created without a size hint; there may be more allocations at assignment time.

​	m 在没有大小提示的情况下创建；在分配时可能会有更多分配。

```go
// Good

files, _ := os.ReadDir("./files")

m := make(map[string]os.DirEntry, len(files))
for _, f := range files {
    m[f.Name()] = f
}
```

`m` is created with a size hint; there may be fewer allocations at assignment time.

​	 `m` 在具有大小提示的情况下创建；在分配时可能会有更少的分配。

#### Specifying Slice Capacity 指定切片容量

Where possible, provide capacity hints when initializing slices with `make()`, particularly when appending.

​	在可能的情况下，在使用 `make()` 初始化切片时提供容量提示，尤其是在追加时。

```
make([]T, length, capacity)
```

Unlike maps, slice capacity is not a hint: the compiler will allocate enough memory for the capacity of the slice as provided to `make()`, which means that subsequent `append()` operations will incur zero allocations (until the length of the slice matches the capacity, after which any appends will require a resize to hold additional elements).

​	与映射不同，切片容量不是提示：编译器将为 `make()` 提供的切片容量分配足够的内存，这意味着后续 `append()` 操作将产生零个分配（直到切片长度与容量匹配，之后任何追加都需要调整大小以容纳其他元素）。

```go
// Bad
for n := 0; n < b.N; n++ {
  data := make([]int, 0)
  for k := 0; k < size; k++{
    data = append(data, k)
  }
}

// BenchmarkBad-4    100000000    2.48s
// Good
for n := 0; n < b.N; n++ {
  data := make([]int, 0, size)
  for k := 0; k < size; k++{
    data = append(data, k)
  }
}

// BenchmarkGood-4   100000000    0.21s
```

## Style 样式

### Avoid overly long lines 避免过长的行

Avoid lines of code that require readers to scroll horizontally or turn their heads too much.

​	避免需要读者水平滚动或过多转动头部才能阅读的代码行。

We recommend a soft line length limit of **99 characters**. Authors should aim to wrap lines before hitting this limit, but it is not a hard limit. Code is allowed to exceed this limit.

​	我们建议将软行长度限制为 99 个字符。作者应在达到此限制之前换行，但这不是硬性限制。代码可以超过此限制。

### Be Consistent 保持一致

Some of the guidelines outlined in this document can be evaluated objectively; others are situational, contextual, or subjective.

​	本文档中概述的一些准则可以客观评估；其他准则则具有情境性、上下文性或主观性。

Above all else, **be consistent**.

​	最重要的是，保持一致。

Consistent code is easier to maintain, is easier to rationalize, requires less cognitive overhead, and is easier to migrate or update as new conventions emerge or classes of bugs are fixed.

​	一致的代码更易于维护、更易于合理化、需要更少的认知开销，并且随着新约定的出现或错误类别的修复，更易于迁移或更新。

Conversely, having multiple disparate or conflicting styles within a single codebase causes maintenance overhead, uncertainty, and cognitive dissonance, all of which can directly contribute to lower velocity, painful code reviews, and bugs.

​	相反，在单个代码库中有多种不同或冲突的样式会导致维护开销、不确定性和认知失调，所有这些都会直接导致速度降低、痛苦的代码审查和错误。

When applying these guidelines to a codebase, it is recommended that changes are made at a package (or larger) level: application at a sub-package level violates the above concern by introducing multiple styles into the same code.

​	将这些准则应用于代码库时，建议在包（或更大）级别进行更改：在子包级别进行应用会违反上述关注点，因为会在同一代码中引入多种样式。

### Group Similar Declarations 对类似声明进行分组

Go supports grouping similar declarations.

​	Go 支持对类似声明进行分组。

```go
// Bad
import "a"
import "b"
// Good
import (
  "a"
  "b"
)
```

This also applies to constants, variables, and type declarations.

​	这也适用于常量、变量和类型声明。

```go
// Bad

const a = 1
const b = 2



var a = 1
var b = 2



type Area float64
type Volume float64
// Good
const (
  a = 1
  b = 2
)

var (
  a = 1
  b = 2
)

type (
  Area float64
  Volume float64
)
```

Only group related declarations. Do not group declarations that are unrelated.

​	仅对相关的声明进行分组。不要对不相关的声明进行分组。

```go
// Bad
type Operation int

const (
  Add Operation = iota + 1
  Subtract
  Multiply
  EnvVar = "MY_ENV"
)
// Good
type Operation int

const (
  Add Operation = iota + 1
  Subtract
  Multiply
)

const EnvVar = "MY_ENV"
```

Groups are not limited in where they can be used. For example, you can use them inside of functions.

​	组的使用位置不受限制。例如，您可以在函数内部使用它们。

```go
// Bad
func f() string {
  red := color.New(0xff0000)
  green := color.New(0x00ff00)
  blue := color.New(0x0000ff)

  // ...
}
// Good
func f() string {
  var (
    red   = color.New(0xff0000)
    green = color.New(0x00ff00)
    blue  = color.New(0x0000ff)
  )

  // ...
}
```

Exception: Variable declarations, particularly inside functions, should be grouped together if declared adjacent to other variables. Do this for variables declared together even if they are unrelated.

​	异常：变量声明，尤其是在函数内部，如果与其他变量相邻声明，则应将其分组在一起。即使变量不相关，也要对一起声明的变量执行此操作。

```go
// Bad
func (c *client) request() {
  caller := c.name
  format := "json"
  timeout := 5*time.Second
  var err error

  // ...
}
// Good
func (c *client) request() {
  var (
    caller  = c.name
    format  = "json"
    timeout = 5*time.Second
    err error
  )

  // ...
}
```

### Import Group Ordering 导入组排序

There should be two import groups:

​	应有两个导入组：

- Standard library
  标准库
- Everything else
  其他所有内容

This is the grouping applied by goimports by default.

​	这是 goimports 默认应用的分组。

```go
// Bad
import (
  "fmt"
  "os"
  "go.uber.org/atomic"
  "golang.org/x/sync/errgroup"
)
// Good
import (
  "fmt"
  "os"

  "go.uber.org/atomic"
  "golang.org/x/sync/errgroup"
)
```

### Package Names 包名称

When naming packages, choose a name that is:

​	在命名包时，请选择一个名称：

- All lower-case. No capitals or underscores.
  全部小写。没有大写字母或下划线。
- Does not need to be renamed using named imports at most call sites.
  在大多数调用站点无需使用命名导入重命名。
- Short and succinct. Remember that the name is identified in full at every call site.
  简短精炼。请记住，名称在每个调用站点都以完整形式标识。
- Not plural. For example, `net/url`, not `net/urls`.
  不是复数。例如， `net/url` ，而不是 `net/urls` 。
- Not “common”, “util”, “shared”, or “lib”. These are bad, uninformative names.
  不是“common”、“util”、“shared”或“lib”。这些都是糟糕的、没有信息量的名称。

See also [Package Names](https://blog.golang.org/package-names) and [Style guideline for Go packages](https://rakyll.org/style-packages/).

​	另请参阅 Go 包的包名称和样式指南。

### Function Names 函数名称

We follow the Go community’s convention of using [MixedCaps for function names](https://golang.org/doc/effective_go.html#mixed-caps). An exception is made for test functions, which may contain underscores for the purpose of grouping related test cases, e.g., `TestMyFunction_WhatIsBeingTested`.

​	我们遵循 Go 社区的惯例，对函数名称使用 MixedCaps。测试函数除外，测试函数可能包含下划线，以便对相关的测试用例进行分组，例如， `TestMyFunction_WhatIsBeingTested` 。

### Import Aliasing 导入别名

Import aliasing must be used if the package name does not match the last element of the import path.

​	如果包名称与导入路径的最后一个元素不匹配，则必须使用导入别名。

```
import (
  "net/http"

  client "example.com/client-go"
  trace "example.com/trace/v2"
)
```

In all other scenarios, import aliases should be avoided unless there is a direct conflict between imports.

​	在所有其他情况下，应避免导入别名，除非导入之间存在直接冲突。

```go
// Bad
import (
  "fmt"
  "os"


  nettrace "golang.net/x/trace"
)
// Good
import (
  "fmt"
  "os"
  "runtime/trace"

  nettrace "golang.net/x/trace"
)
```

### Function Grouping and Ordering 函数分组和排序

- Functions should be sorted in rough call order.
  应按大致调用顺序对函数进行排序。
- Functions in a file should be grouped by receiver.
  文件中应按接收者对函数进行分组。

Therefore, exported functions should appear first in a file, after `struct`, `const`, `var` definitions.

​	因此，导出的函数应首先出现在文件中，位于 `struct` 、 `const` 、 `var` 定义之后。

A `newXYZ()`/`NewXYZ()` may appear after the type is defined, but before the rest of the methods on the receiver.

​	 `newXYZ()` / `NewXYZ()` 可能会在定义类型之后但接收者上的其他方法之前出现。

Since functions are grouped by receiver, plain utility functions should appear towards the end of the file.

​	由于函数按接收者分组，因此普通实用函数应出现在文件的末尾。

```go
// Bad
func (s *something) Cost() {
  return calcCost(s.weights)
}

type something struct{ ... }

func calcCost(n []int) int {...}

func (s *something) Stop() {...}

func newSomething() *something {
    return &something{}
}
// Good
type something struct{ ... }

func newSomething() *something {
    return &something{}
}

func (s *something) Cost() {
  return calcCost(s.weights)
}

func (s *something) Stop() {...}

func calcCost(n []int) int {...}
```

### Reduce Nesting 减少嵌套

Code should reduce nesting where possible by handling error cases/special conditions first and returning early or continuing the loop. Reduce the amount of code that is nested multiple levels.

​	代码应尽可能通过首先处理错误情况/特殊情况并尽早返回或继续循环来减少嵌套。减少嵌套多层的代码量。

```go
// Bad
for _, v := range data {
  if v.F1 == 1 {
    v = process(v)
    if err := v.Call(); err == nil {
      v.Send()
    } else {
      return err
    }
  } else {
    log.Printf("Invalid v: %v", v)
  }
}
// Good
for _, v := range data {
  if v.F1 != 1 {
    log.Printf("Invalid v: %v", v)
    continue
  }

  v = process(v)
  if err := v.Call(); err != nil {
    return err
  }
  v.Send()
}
```

### Unnecessary Else 不必要的 Else

If a variable is set in both branches of an if, it can be replaced with a single if.

​	如果变量在 if 的两个分支中都设置了，则可以用一个 if 替换它。

```go
// Bad
var a int
if b {
  a = 100
} else {
  a = 10
}
// Good
a := 10
if b {
  a = 100
}
```

### Top-level Variable Declarations 顶级变量声明

At the top level, use the standard `var` keyword. Do not specify the type, unless it is not the same type as the expression.

​	在顶级，使用标准 `var` 关键字。不要指定类型，除非它与表达式的类型不同。

```go
// Bad
var _s string = F()

func F() string { return "A" }
// Good
var _s = F()
// Since F already states that it returns a string, we don't need to specify
// the type again.

func F() string { return "A" }
```

Specify the type if the type of the expression does not match the desired type exactly.

​	如果表达式的类型与所需类型不完全匹配，则指定类型。

```go
type myError struct{}

func (myError) Error() string { return "error" }

func F() myError { return myError{} }

var _e error = F()
// F returns an object of type myError but we want error.
```

### Prefix Unexported Globals with _ 前缀未导出的全局变量为 _

Prefix unexported top-level `var`s and `const`s with `_` to make it clear when they are used that they are global symbols.

​	前缀未导出的顶级 `var` 和 `const` 为 `_` ，以便在使用时清楚地表明它们是全局符号。

Rationale: Top-level variables and constants have a package scope. Using a generic name makes it easy to accidentally use the wrong value in a different file.

​	基本原理：顶级变量和常量具有包范围。使用通用名称很容易在不同的文件中意外使用错误的值。

```go
// Bad
// foo.go

const (
  defaultPort = 8080
  defaultUser = "user"
)

// bar.go

func Bar() {
  defaultPort := 9090
  ...
  fmt.Println("Default port", defaultPort)

  // We will not see a compile error if the first line of
  // Bar() is deleted.
}
// Good
// foo.go

const (
  _defaultPort = 8080
  _defaultUser = "user"
)
```

**Exception**: Unexported error values may use the prefix `err` without the underscore. See [Error Naming](https://github.com/uber-go/guide/blob/master/style.md#error-naming).

​	例外：未导出的错误值可以使用前缀 `err` 而没有下划线。请参阅错误命名。

### Embedding in Structs 嵌入到结构中

Embedded types should be at the top of the field list of a struct, and there must be an empty line separating embedded fields from regular fields.

​	嵌入式类型应位于结构的字段列表顶部，并且必须有空行将嵌入式字段与常规字段分隔开。

```go
// Bad
type Client struct {
  version int
  http.Client
}
// Good
type Client struct {
  http.Client

  version int
}
```

Embedding should provide tangible benefit, like adding or augmenting functionality in a semantically-appropriate way. It should do this with zero adverse user-facing effects (see also: [Avoid Embedding Types in Public Structs](https://github.com/uber-go/guide/blob/master/style.md#avoid-embedding-types-in-public-structs)).

​	嵌入应该提供切实的益处，例如以语义上合适的方式添加或增强功能。它应该做到这一点，而不会对用户产生任何不利影响（另请参阅：避免在公共结构中嵌入类型）。

Exception: Mutexes should not be embedded, even on unexported types. See also: [Zero-value Mutexes are Valid](https://github.com/uber-go/guide/blob/master/style.md#zero-value-mutexes-are-valid).

​	例外：即使在未导出的类型上，也不应嵌入互斥锁。另请参阅：零值互斥锁有效。

Embedding **should not**:

​	嵌入不应：

- Be purely cosmetic or convenience-oriented.
  纯粹是装饰性的或以方便为导向。
- Make outer types more difficult to construct or use.
  使外部类型更难构造或使用。
- Affect outer types’ zero values. If the outer type has a useful zero value, it should still have a useful zero value after embedding the inner type.
  影响外部类型的零值。如果外部类型具有有用的零值，那么在嵌入内部类型后，它仍应具有有用的零值。
- Expose unrelated functions or fields from the outer type as a side-effect of embedding the inner type.
  将外部类型中不相关的函数或字段作为嵌入内部类型的副作用暴露出来。
- Expose unexported types.
  暴露未导出的类型。
- Affect outer types’ copy semantics.
  影响外部类型的复制语义。
- Change the outer type’s API or type semantics.
  更改外部类型的 API 或类型语义。
- Embed a non-canonical form of the inner type.
  嵌入内部类型的非规范形式。
- Expose implementation details of the outer type.
  公开外部类型的实现细节。
- Allow users to observe or control type internals.
  允许用户观察或控制类型内部。
- Change the general behavior of inner functions through wrapping in a way that would reasonably surprise users.
  通过包装改变内部函数的一般行为，这种方式可能会合理地让用户感到惊讶。

Simply put, embed consciously and intentionally. A good litmus test is, “would all of these exported inner methods/fields be added directly to the outer type”; if the answer is “some” or “no”, don’t embed the inner type - use a field instead.

​	简单来说，有意识地和有意图地嵌入。一个很好的试金石是，“所有这些导出的内部方法/字段是否会直接添加到外部类型”；如果答案是“有些”或“没有”，则不要嵌入内部类型 - 而要使用字段。

```go
// Bad
type A struct {
    // Bad: A.Lock() and A.Unlock() are
    //      now available, provide no
    //      functional benefit, and allow
    //      users to control details about
    //      the internals of A.
    sync.Mutex
}

// ---------------------------------------

type Book struct {
    // Bad: pointer changes zero value usefulness
    io.ReadWriter

    // other fields
}

// later

var b Book
b.Read(...)  // panic: nil pointer
b.String()   // panic: nil pointer
b.Write(...) // panic: nil pointer

// ---------------------------------------

type Client struct {
    sync.Mutex
    sync.WaitGroup
    bytes.Buffer
    url.URL
}
// Good
type countingWriteCloser struct {
    // Good: Write() is provided at this
    //       outer layer for a specific
    //       purpose, and delegates work
    //       to the inner type's Write().
    io.WriteCloser

    count int
}

func (w *countingWriteCloser) Write(bs []byte) (int, error) {
    w.count += len(bs)
    return w.WriteCloser.Write(bs)
}

// ---------------------------------------

type Book struct {
    // Good: has useful zero value
    bytes.Buffer

    // other fields
}

// later

var b Book
b.Read(...)  // ok
b.String()   // ok
b.Write(...) // ok

// ---------------------------------------

type Client struct {
    mtx sync.Mutex
    wg  sync.WaitGroup
    buf bytes.Buffer
    url url.URL
}
```

### Local Variable Declarations 局部变量声明

Short variable declarations (`:=`) should be used if a variable is being set to some value explicitly.

​	如果变量被明确地设置为某个值，则应使用短变量声明 ( `:=` )。

```go
// Bad
var s = "foo"
// Good
s := "foo"
```

However, there are cases where the default value is clearer when the `var` keyword is used. [Declaring Empty Slices](https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices), for example.

​	但是，在使用 `var` 关键字时，默认值更清晰的情况也有。例如，声明空切片。

```go
// Bad
func f(list []int) {
  filtered := []int{}
  for _, v := range list {
    if v > 10 {
      filtered = append(filtered, v)
    }
  }
}
// Good
func f(list []int) {
  var filtered []int
  for _, v := range list {
    if v > 10 {
      filtered = append(filtered, v)
    }
  }
}
```

### nil is a valid slice nil 是一个有效的切片

`nil` is a valid slice of length 0. This means that,

​	 `nil` 是一个长度为 0 的有效切片。这意味着，

- You should not return a slice of length zero explicitly. Return `nil` instead.

  ​	您不应该显式地返回长度为零的切片。相反，返回 `nil` 。

  ```go
  // Bad
  if x == "" {
    return []int{}
  }
  ```

  ```go
  // Good
  if x == "" {
    return nil
  }
  ```

- To check if a slice is empty, always use `len(s) == 0`. Do not check for `nil`.

  ​	要检查切片是否为空，请始终使用 `len(s) == 0` 。不要检查 `nil` 。

  ```go
  // Bad
  func isEmpty(s []string) bool {
    return s == nil
  }
  ```

  ```go
  // Good
  func isEmpty(s []string) bool {
    return len(s) == 0
  }
  ```

- The zero value (a slice declared with `var`) is usable immediately without `make()`.

  ​	零值（使用 `var` 声明的切片）可以直接使用，而无需 `make()` 。

  ```go
  // Bad
  nums := []int{}
  // or, nums := make([]int)
  
  if add1 {
    nums = append(nums, 1)
  }
  
  if add2 {
    nums = append(nums, 2)
  }
  ```

  ```go
  // Good
  var nums []int
  
  if add1 {
    nums = append(nums, 1)
  }
  
  if add2 {
    nums = append(nums, 2)
  }
  ```

Remember that, while it is a valid slice, a nil slice is not equivalent to an allocated slice of length 0 - one is nil and the other is not - and the two may be treated differently in different situations (such as serialization).

​	请记住，虽然它是有效的切片，但 nil 切片不等于长度为 0 的已分配切片——一个是 nil，另一个不是——并且在不同情况下（例如序列化）可能会对它们进行不同的处理。

### Reduce Scope of Variables 减少变量的作用域

Where possible, reduce scope of variables. Do not reduce the scope if it conflicts with [Reduce Nesting](https://github.com/uber-go/guide/blob/master/style.md#reduce-nesting).

​	尽可能减少变量的作用域。如果与减少嵌套冲突，请不要减少作用域。

```go
// Bad
err := os.WriteFile(name, data, 0644)
if err != nil {
 return err
}
// Good
if err := os.WriteFile(name, data, 0644); err != nil {
 return err
}
```

If you need a result of a function call outside of the if, then you should not try to reduce the scope.

​	如果您需要在 if 之外使用函数调用的结果，那么您不应该尝试减少作用域。

```go
// Bad
if data, err := os.ReadFile(name); err == nil {
  err = cfg.Decode(data)
  if err != nil {
    return err
  }

  fmt.Println(cfg)
  return nil
} else {
  return err
}
// Good
data, err := os.ReadFile(name)
if err != nil {
   return err
}

if err := cfg.Decode(data); err != nil {
  return err
}

fmt.Println(cfg)
return nil
```

### Avoid Naked Parameters 避免裸参数

Naked parameters in function calls can hurt readability. Add C-style comments (`/* ... */`) for parameter names when their meaning is not obvious.

​	函数调用中的裸参数可能会损害可读性。当参数的含义不明显时，请为参数名称添加 C 样式注释（ `/* ... */` ）。

```go
// Bad
// func printInfo(name string, isLocal, done bool)

printInfo("foo", true, true)
// Good
// func printInfo(name string, isLocal, done bool)

printInfo("foo", true /* isLocal */, true /* done */)
```

Better yet, replace naked `bool` types with custom types for more readable and type-safe code. This allows more than just two states (true/false) for that parameter in the future.

​	更好的是，用自定义类型替换裸 `bool` 类型，以获得更具可读性和类型安全性的代码。这允许该参数在未来具有不止两种状态（真/假）。

```go
type Region int

const (
  UnknownRegion Region = iota
  Local
)

type Status int

const (
  StatusReady Status = iota + 1
  StatusDone
  // Maybe we will have a StatusInProgress in the future.
)

func printInfo(name string, region Region, status Status)
```

### Use Raw String Literals to Avoid Escaping 使用原始字符串字面量避免转义

Go supports [raw string literals](https://golang.org/ref/spec#raw_string_lit), which can span multiple lines and include quotes. Use these to avoid hand-escaped strings which are much harder to read.

​	Go 支持原始字符串字面量，它可以跨越多行并包含引号。使用这些来避免手动转义的字符串，这些字符串更难阅读。

```go
// Bad
wantError := "unknown name:\"test\""
// Good
wantError := `unknown error:"test"`
```

### Initializing Structs 初始化结构体

#### Use Field Names to Initialize Structs 使用字段名初始化结构体

You should almost always specify field names when initializing structs. This is now enforced by [`go vet`](https://golang.org/cmd/vet/).

​	初始化结构体时，您几乎应该始终指定字段名。现在由 `go vet` 强制执行。

```go
// Bad
k := User{"John", "Doe", true}
// Good
k := User{
    FirstName: "John",
    LastName: "Doe",
    Admin: true,
}
```

Exception: Field names *may* be omitted in test tables when there are 3 or fewer fields.

​	例外：当字段少于或等于 3 个时，可以在测试表中省略字段名。

```
tests := []struct{
  op Operation
  want string
}{
  {Add, "add"},
  {Subtract, "subtract"},
}
```

#### Omit Zero Value Fields in Structs 在结构体中省略零值字段

When initializing structs with field names, omit fields that have zero values unless they provide meaningful context. Otherwise, let Go set these to zero values automatically.

​	使用字段名初始化结构体时，省略具有零值的字段，除非它们提供有意义的上下文。否则，让 Go 自动将这些字段设置为零值。

```go
// Bad
user := User{
  FirstName: "John",
  LastName: "Doe",
  MiddleName: "",
  Admin: false,
}
// Good
user := User{
  FirstName: "John",
  LastName: "Doe",
}
```

This helps reduce noise for readers by omitting values that are default in that context. Only meaningful values are specified.

​	这有助于通过省略在该上下文中为默认值的字段来减少阅读者的干扰。只指定有意义的值。

Include zero values where field names provide meaningful context. For example, test cases in [Test Tables](https://github.com/uber-go/guide/blob/master/style.md#test-tables) can benefit from names of fields even when they are zero-valued.

​	在字段名提供有意义的上下文时包含零值。例如，测试表中的测试用例即使在为零值时也能从字段名中受益。

```
tests := []struct{
  give string
  want int
}{
  {give: "0", want: 0},
  // ...
}
```

#### Use `var` for Zero Value Structs 使用 `var` 表示零值结构

When all the fields of a struct are omitted in a declaration, use the `var` form to declare the struct.

​	如果声明中省略了结构的所有字段，请使用 `var` 形式来声明结构。

```go
// Bad
user := User{}
var user User
```

This differentiates zero valued structs from those with non-zero fields similar to the distinction created for [map initialization](https://github.com/uber-go/guide/blob/master/style.md#initializing-maps), and matches how we prefer to [declare empty slices](https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices).

​	这将零值结构与具有非零字段的结构区分开来，类似于为映射初始化创建的区分，并且与我们更喜欢声明空切片的方式相匹配。

#### Initializing Struct References 初始化结构引用

Use `&T{}` instead of `new(T)` when initializing struct references so that it is consistent with the struct initialization.

​	在初始化结构引用时使用 `&T{}` 而不是 `new(T)` ，以便与结构初始化保持一致。

```go
// Bad
sval := T{Name: "foo"}

// inconsistent
sptr := new(T)
sptr.Name = "bar"
// Good
sval := T{Name: "foo"}

sptr := &T{Name: "bar"}
```

### Initializing Maps 初始化映射

Prefer `make(..)` for empty maps, and maps populated programmatically. This makes map initialization visually distinct from declaration, and it makes it easy to add size hints later if available.

​	对于空映射和以编程方式填充的映射，更喜欢使用 `make(..)` 。这使得映射初始化在视觉上与声明不同，并且如果以后有可用的大小提示，则可以轻松添加。

- Declaration and initialization are visually similar.
  声明和初始化在视觉上相似。

```go
// Bad
var (
  // m1 is safe to read and write;
  // m2 will panic on writes.
  m1 = map[T1]T2{}
  m2 map[T1]T2
)
```

- Declaration and initialization are visually distinct.
  声明和初始化在视觉上不同。

```go
// Good
var (
  // m1 is safe to read and write;
  // m2 will panic on writes.
  m1 = make(map[T1]T2)
  m2 map[T1]T2
)
```

Where possible, provide capacity hints when initializing maps with `make()`. See [Specifying Map Capacity Hints](https://github.com/uber-go/guide/blob/master/style.md#specifying-map-capacity-hints) for more information.

​	如果可能，请在使用 `make()` 初始化映射时提供容量提示。有关更多信息，请参阅指定映射容量提示。

On the other hand, if the map holds a fixed list of elements, use map literals to initialize the map.

​	另一方面，如果映射包含固定元素列表，请使用映射字面量来初始化映射。

```go
// Bad
m := make(map[T1]T2, 3)
m[k1] = v1
m[k2] = v2
m[k3] = v3
// Good
m := map[T1]T2{
  k1: v1,
  k2: v2,
  k3: v3,
}
```

The basic rule of thumb is to use map literals when adding a fixed set of elements at initialization time, otherwise use `make` (and specify a size hint if available).

​	基本经验法则是，在初始化时添加固定元素集时使用映射字面量，否则使用 `make` （如果可用，请指定大小提示）。

### Format Strings outside Printf Printf 之外的格式字符串

If you declare format strings for `Printf`-style functions outside a string literal, make them `const` values.

​	如果您在字符串字面量之外为 `Printf` 样式函数声明格式字符串，请使其成为 `const` 值。

This helps `go vet` perform static analysis of the format string.

​	这有助于 `go vet` 对格式字符串执行静态分析。

```go
// Bad
msg := "unexpected values %v, %v\n"
fmt.Printf(msg, 1, 2)
// Good
const msg = "unexpected values %v, %v\n"
fmt.Printf(msg, 1, 2)
```

### Naming Printf-style Functions 命名 Printf 样式函数

When you declare a `Printf`-style function, make sure that `go vet` can detect it and check the format string.

​	当您声明 `Printf` 样式函数时，请确保 `go vet` 可以检测到它并检查格式字符串。

This means that you should use predefined `Printf`-style function names if possible. `go vet` will check these by default. See [Printf family](https://golang.org/cmd/vet/#hdr-Printf_family) for more information.

​	这意味着您应该尽可能使用预定义的 `Printf` 样式函数名称。 `go vet` 将默认检查这些名称。有关更多信息，请参阅 Printf 系列。

If using the predefined names is not an option, end the name you choose with f: `Wrapf`, not `Wrap`. `go vet` can be asked to check specific `Printf`-style names but they must end with f.

​	如果使用预定义名称不是一种选择，请以 f 结尾： `Wrapf` ，而不是 `Wrap` 。可以要求 `go vet` 检查特定的 `Printf` 样式名称，但它们必须以 f 结尾。

```
go vet -printfuncs=wrapf,statusf
```

See also [go vet: Printf family check](https://kuzminva.wordpress.com/2017/11/07/go-vet-printf-family-check/).

​	另请参阅 go vet：Printf 系列检查。

## Patterns 模式

### Test Tables 测试表

Table-driven tests with [subtests](https://blog.golang.org/subtests) can be a helpful pattern for writing tests to avoid duplicating code when the core test logic is repetitive.

​	具有子测试的表驱动测试可以成为编写测试的有用模式，以避免在核心测试逻辑重复时复制代码。

If a system under test needs to be tested against *multiple conditions* where certain parts of the the inputs and outputs change, a table-driven test should be used to reduce redundancy and improve readability.

​	如果需要针对某些输入和输出部分发生变化的多种条件对被测系统进行测试，则应使用表驱动测试来减少冗余并提高可读性。

```go
// Bad
// func TestSplitHostPort(t *testing.T)

host, port, err := net.SplitHostPort("192.0.2.0:8000")
require.NoError(t, err)
assert.Equal(t, "192.0.2.0", host)
assert.Equal(t, "8000", port)

host, port, err = net.SplitHostPort("192.0.2.0:http")
require.NoError(t, err)
assert.Equal(t, "192.0.2.0", host)
assert.Equal(t, "http", port)

host, port, err = net.SplitHostPort(":8000")
require.NoError(t, err)
assert.Equal(t, "", host)
assert.Equal(t, "8000", port)

host, port, err = net.SplitHostPort("1:8")
require.NoError(t, err)
assert.Equal(t, "1", host)
assert.Equal(t, "8", port)
// Good
// func TestSplitHostPort(t *testing.T)

tests := []struct{
  give     string
  wantHost string
  wantPort string
}{
  {
    give:     "192.0.2.0:8000",
    wantHost: "192.0.2.0",
    wantPort: "8000",
  },
  {
    give:     "192.0.2.0:http",
    wantHost: "192.0.2.0",
    wantPort: "http",
  },
  {
    give:     ":8000",
    wantHost: "",
    wantPort: "8000",
  },
  {
    give:     "1:8",
    wantHost: "1",
    wantPort: "8",
  },
}

for _, tt := range tests {
  t.Run(tt.give, func(t *testing.T) {
    host, port, err := net.SplitHostPort(tt.give)
    require.NoError(t, err)
    assert.Equal(t, tt.wantHost, host)
    assert.Equal(t, tt.wantPort, port)
  })
}
```

Test tables make it easier to add context to error messages, reduce duplicate logic, and add new test cases.

​	测试表使添加上下文到错误消息、减少重复逻辑和添加新测试用例变得更加容易。

We follow the convention that the slice of structs is referred to as `tests` and each test case `tt`. Further, we encourage explicating the input and output values for each test case with `give` and `want` prefixes.

​	我们遵循将结构切片称为 `tests` 和每个测试用例 `tt` 的约定。此外，我们鼓励使用 `give` 和 `want` 前缀明确说明每个测试用例的输入和输出值。

```
tests := []struct{
  give     string
  wantHost string
  wantPort string
}{
  // ...
}

for _, tt := range tests {
  // ...
}
```

#### Avoid Unnecessary Complexity in Table Tests 避免在表测试中出现不必要的复杂性

Table tests can be difficult to read and maintain if the subtests contain conditional assertions or other branching logic. Table tests should **NOT** be used whenever there needs to be complex or conditional logic inside subtests (i.e. complex logic inside the `for` loop).

​	如果子测试包含条件断言或其他分支逻辑，则表测试可能难以阅读和维护。每当子测试中需要复杂或条件逻辑（即 `for` 循环内的复杂逻辑）时，都不应使用表测试。

Large, complex table tests harm readability and maintainability because test readers may have difficulty debugging test failures that occur.

​	大型、复杂的表测试会损害可读性和可维护性，因为测试阅读器可能难以调试发生的测试失败。

Table tests like this should be split into either multiple test tables or multiple individual `Test...` functions.

​	像这样的表格测试应该拆分为多个测试表格或多个单独的 `Test...` 函数。

Some ideals to aim for are:

​	一些要达到的理想目标是：

- Focus on the narrowest unit of behavior
  关注最窄的行为单元
- Minimize “test depth”, and avoid conditional assertions (see below)
  最小化“测试深度”，并避免条件断言（见下文）
- Ensure that all table fields are used in all tests
  确保所有表格字段在所有测试中都使用
- Ensure that all test logic runs for all table cases
  确保所有测试逻辑对所有表格用例都运行

In this context, “test depth” means “within a given test, the number of successive assertions that require previous assertions to hold” (similar to cyclomatic complexity). Having “shallower” tests means that there are fewer relationships between assertions and, more importantly, that those assertions are less likely to be conditional by default.

​	在此上下文中，“测试深度”是指“在给定测试中，需要先前的断言成立的连续断言的数量”（类似于圈复杂度）。进行“较浅”的测试意味着断言之间的关系更少，更重要的是，这些断言默认情况下不太可能是有条件的。

Concretely, table tests can become confusing and difficult to read if they use multiple branching pathways (e.g. `shouldError`, `expectCall`, etc.), use many `if` statements for specific mock expectations (e.g. `shouldCallFoo`), or place functions inside the table (e.g. `setupMocks func(*FooMock)`).

​	具体来说，如果表格测试使用多个分支路径（例如 `shouldError` 、 `expectCall` 等），对特定模拟期望使用许多 `if` 语句（例如 `shouldCallFoo` ），或在表格中放置函数（例如 `setupMocks func(*FooMock)` ），则表格测试可能会变得混乱且难以阅读。

However, when testing behavior that only changes based on changed input, it may be preferable to group similar cases together in a table test to better illustrate how behavior changes across all inputs, rather than splitting otherwise comparable units into separate tests and making them harder to compare and contrast.

​	然而，在测试仅基于更改的输入而更改的行为时，最好将类似的情况组合在一个表测试中，以更好地说明行为如何跨所有输入发生变化，而不是将其他可比较的单元拆分为单独的测试，并使它们更难比较和对比。

If the test body is short and straightforward, it’s acceptable to have a single branching pathway for success versus failure cases with a table field like `shouldErr` to specify error expectations.

​	如果测试主体简短且直接，则可以接受为成功与失败案例提供一个分支路径，并使用 `shouldErr` 之类的表字段来指定错误预期。

```go
// Bad
func TestComplicatedTable(t *testing.T) {
  tests := []struct {
    give          string
    want          string
    wantErr       error
    shouldCallX   bool
    shouldCallY   bool
    giveXResponse string
    giveXErr      error
    giveYResponse string
    giveYErr      error
  }{
    // ...
  }

  for _, tt := range tests {
    t.Run(tt.give, func(t *testing.T) {
      // setup mocks
      ctrl := gomock.NewController(t)
      xMock := xmock.NewMockX(ctrl)
      if tt.shouldCallX {
        xMock.EXPECT().Call().Return(
          tt.giveXResponse, tt.giveXErr,
        )
      }
      yMock := ymock.NewMockY(ctrl)
      if tt.shouldCallY {
        yMock.EXPECT().Call().Return(
          tt.giveYResponse, tt.giveYErr,
        )
      }

      got, err := DoComplexThing(tt.give, xMock, yMock)

      // verify results
      if tt.wantErr != nil {
        require.EqualError(t, err, tt.wantErr)
        return
      }
      require.NoError(t, err)
      assert.Equal(t, want, got)
    })
  }
}
// Good
func TestShouldCallX(t *testing.T) {
  // setup mocks
  ctrl := gomock.NewController(t)
  xMock := xmock.NewMockX(ctrl)
  xMock.EXPECT().Call().Return("XResponse", nil)

  yMock := ymock.NewMockY(ctrl)

  got, err := DoComplexThing("inputX", xMock, yMock)

  require.NoError(t, err)
  assert.Equal(t, "want", got)
}

func TestShouldCallYAndFail(t *testing.T) {
  // setup mocks
  ctrl := gomock.NewController(t)
  xMock := xmock.NewMockX(ctrl)

  yMock := ymock.NewMockY(ctrl)
  yMock.EXPECT().Call().Return("YResponse", nil)

  _, err := DoComplexThing("inputY", xMock, yMock)
  assert.EqualError(t, err, "Y failed")
}
```

This complexity makes it more difficult to change, understand, and prove the correctness of the test.

​	这种复杂性使得更改、理解和证明测试的正确性变得更加困难。

While there are no strict guidelines, readability and maintainability should always be top-of-mind when deciding between Table Tests versus separate tests for multiple inputs/outputs to a system.

​	虽然没有严格的准则，但在决定针对系统中的多个输入/输出使用表测试还是单独的测试时，可读性和可维护性应始终是首要考虑因素。

#### Parallel Tests 并行测试

Parallel tests, like some specialized loops (for example, those that spawn goroutines or capture references as part of the loop body), must take care to explicitly assign loop variables within the loop’s scope to ensure that they hold the expected values.

​	并行测试（例如某些专门的循环，例如那些在循环体中生成 goroutine 或捕获引用的循环）必须注意在循环范围内显式分配循环变量，以确保它们具有预期值。

```
tests := []struct{
  give string
  // ...
}{
  // ...
}

for _, tt := range tests {
  tt := tt // for t.Parallel
  t.Run(tt.give, func(t *testing.T) {
    t.Parallel()
    // ...
  })
}
```

In the example above, we must declare a `tt` variable scoped to the loop iteration because of the use of `t.Parallel()` below. If we do not do that, most or all tests will receive an unexpected value for `tt`, or a value that changes as they’re running.

​	在上面的示例中，我们必须声明一个 `tt` 变量，其作用域为循环迭代，因为下面使用了 `t.Parallel()` 。如果不这样做，大多数或所有测试都会收到 `tt` 的意外值，或者在运行时会更改的值。

### Functional Options 功能选项

Functional options is a pattern in which you declare an opaque `Option` type that records information in some internal struct. You accept a variadic number of these options and act upon the full information recorded by the options on the internal struct.

​	功能选项是一种模式，您可以在其中声明一个不透明的 `Option` 类型，该类型在某些内部结构中记录信息。您可以接受这些选项的可变数量，并根据内部结构上选项记录的完整信息采取行动。

Use this pattern for optional arguments in constructors and other public APIs that you foresee needing to expand, especially if you already have three or more arguments on those functions.

​	在您预见到需要扩展的构造函数和其他公共 API 中使用此模式来获取可选参数，尤其是当这些函数上已经存在三个或更多参数时。

```go
// Bad
// package db

func Open(
  addr string,
  cache bool,
  logger *zap.Logger
) (*Connection, error) {
  // ...
}
```

The cache and logger parameters must always be provided, even if the user wants to use the default.

​	即使用户想要使用默认值，也必须始终提供缓存和记录器参数。

```go
db.Open(addr, db.DefaultCache, zap.NewNop())
db.Open(addr, db.DefaultCache, log)
db.Open(addr, false /* cache */, zap.NewNop())
db.Open(addr, false /* cache */, log)
// Good
// package db

type Option interface {
  // ...
}

func WithCache(c bool) Option {
  // ...
}

func WithLogger(log *zap.Logger) Option {
  // ...
}

// Open creates a connection.
func Open(
  addr string,
  opts ...Option,
) (*Connection, error) {
  // ...
}
```

Options are provided only if needed.

​	仅在需要时才提供选项。

```go
db.Open(addr)
db.Open(addr, db.WithLogger(log))
db.Open(addr, db.WithCache(false))
db.Open(
  addr,
  db.WithCache(false),
  db.WithLogger(log),
)
```

Our suggested way of implementing this pattern is with an `Option` interface that holds an unexported method, recording options on an unexported `options` struct.

​	我们建议使用包含未导出方法的 `Option` 接口来实现此模式，该方法在未导出的 `options` 结构上记录选项。

```go
type options struct {
  cache  bool
  logger *zap.Logger
}

type Option interface {
  apply(*options)
}

type cacheOption bool

func (c cacheOption) apply(opts *options) {
  opts.cache = bool(c)
}

func WithCache(c bool) Option {
  return cacheOption(c)
}

type loggerOption struct {
  Log *zap.Logger
}

func (l loggerOption) apply(opts *options) {
  opts.logger = l.Log
}

func WithLogger(log *zap.Logger) Option {
  return loggerOption{Log: log}
}

// Open creates a connection.
func Open(
  addr string,
  opts ...Option,
) (*Connection, error) {
  options := options{
    cache:  defaultCache,
    logger: zap.NewNop(),
  }

  for _, o := range opts {
    o.apply(&options)
  }

  // ...
}
```

Note that there’s a method of implementing this pattern with closures but we believe that the pattern above provides more flexibility for authors and is easier to debug and test for users. In particular, it allows options to be compared against each other in tests and mocks, versus closures where this is impossible. Further, it lets options implement other interfaces, including `fmt.Stringer` which allows for user-readable string representations of the options.

​	请注意，有一种使用闭包实现此模式的方法，但我们认为上述模式为作者提供了更大的灵活性，并且用户更容易调试和测试。特别是，它允许在测试和模拟中相互比较选项，而闭包则无法做到这一点。此外，它允许选项实现其他接口，包括 `fmt.Stringer` ，这允许用户可读的字符串表示选项。

See also,

​	另请参阅，

- [Self-referential functions and the design of options
  自引用函数和选项设计](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)
- [Functional options for friendly APIs
  友好 API 的函数选项](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

## Linting

More importantly than any “blessed” set of linters, lint consistently across a codebase.

​	比任何“推荐”的 linter 集合更重要的是，在整个代码库中一致地进行 lint。

We recommend using the following linters at a minimum, because we feel that they help to catch the most common issues and also establish a high bar for code quality without being unnecessarily prescriptive:

​	我们建议至少使用以下 linter，因为我们认为它们有助于发现最常见的问题，并且在不必要地规定性情况下为代码质量设定了很高的标准：

- [errcheck](https://github.com/kisielk/errcheck) to ensure that errors are handled
  errcheck 确保处理错误
- [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) to format code and manage imports
  goimports 格式化代码并管理导入
- [golint](https://github.com/golang/lint) to point out common style mistakes
  golint 指出常见的样式错误
- [govet](https://golang.org/cmd/vet/) to analyze code for common mistakes
  govet 分析代码以查找常见错误
- [staticcheck](https://staticcheck.io/) to do various static analysis checks
  staticcheck 执行各种静态分析检查

### Lint Runners Lint 运行程序

We recommend [golangci-lint](https://github.com/golangci/golangci-lint) as the go-to lint runner for Go code, largely due to its performance in larger codebases and ability to configure and use many canonical linters at once. This repo has an example [.golangci.yml](https://github.com/uber-go/guide/blob/master/.golangci.yml) config file with recommended linters and settings.

​	我们推荐 golangci-lint 作为 Go 代码的首选 lint 运行程序，这在很大程度上是因为它在较大的代码库中的性能以及一次配置和使用许多规范化 linter 的能力。此存储库有一个示例 .golangci.yml 配置文件，其中包含推荐的 linter 和设置。

golangci-lint has [various linters](https://golangci-lint.run/usage/linters/) available for use. The above linters are recommended as a base set, and we encourage teams to add any additional linters that make sense for their projects.

​	golangci-lint 有各种可供使用的 linter。上述 linter 建议作为基本集，我们鼓励团队添加任何对他们的项目有意义的其他 linter。

