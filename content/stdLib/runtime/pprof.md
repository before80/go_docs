+++
title = "pprof"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/runtime/pprof@go1.23.0](https://pkg.go.dev/runtime/pprof@go1.23.0)

Package pprof writes runtime profiling data in the format expected by the pprof visualization tool.

​	pprof 包以 pprof 可视化工具期望的格式编写运行时分析数据。

## Profiling a Go program 分析 Go 程序

The first step to profiling a Go program is to enable profiling. Support for profiling benchmarks built with the standard testing package is built into go test. For example, the following command runs benchmarks in the current directory and writes the CPU and memory profiles to cpu.prof and mem.prof:

​	分析 Go 程序的第一步是启用分析。对使用标准测试包构建的基准测试进行分析的支持已内置到 go test 中。例如，以下命令在当前目录中运行基准测试，并将 CPU 和内存分析结果分别写入 cpu.prof 和 mem.prof：

```
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
```

To add equivalent profiling support to a standalone program, add code like the following to your main function:

​	要向独立程序添加等效的分析支持，请将类似以下的代码添加到您的 main 函数中：

```go
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

    // ... rest of the program ...

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
    }
}
```

There is also a standard HTTP interface to profiling data. Adding the following line will install handlers under the /debug/pprof/ URL to download live profiles:

​	还提供了一个标准的 HTTP 接口来分析数据。添加以下行将在 /debug/pprof/ URL 下安装处理程序以下载实时分析结果：

```
import _ "net/http/pprof"
```

See the net/http/pprof package for more details.

​	有关更多详细信息，请参阅 net/http/pprof 包。

Profiles can then be visualized with the pprof tool:

​	然后可以使用 pprof 工具可视化分析结果：

```
go tool pprof cpu.prof
```

There are many commands available from the pprof command line. Commonly used commands include “top”, which prints a summary of the top program hot-spots, and “web”, which opens an interactive graph of hot-spots and their call graphs. Use “help” for information on all pprof commands.

​	pprof 命令行提供了许多命令。常用的命令包括“top”，它会打印程序热点摘要，以及“web”，它会打开热点及其调用图的交互式图表。使用“help”可了解所有 pprof 命令的信息。

For more information about pprof, see https://github.com/google/pprof/blob/master/doc/README.md.

​	有关 pprof 的更多信息，请参阅 https://github.com/google/pprof/blob/master/doc/README.md。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func Do <- go1.9

```go
func Do(ctx context.Context, labels LabelSet, f func(context.Context))
```

Do calls f with a copy of the parent context with the given labels added to the parent’s label map. Goroutines spawned while executing f will inherit the augmented label-set. Each key/value pair in labels is inserted into the label map in the order provided, overriding any previous value for the same key. The augmented label map will be set for the duration of the call to f and restored once f returns.

​	Do 使用带有添加到父标签映射的给定标签的父上下文副本调用 f。在执行 f 时产生的协程将继承扩充的标签集。标签中的每个键/值对按提供的顺序插入到标签映射中，覆盖相同键的任何先前值。扩充的标签映射将设置在对 f 的调用持续时间内，并在 f 返回后还原。

### func ForLabels <- go1.9

```go
func ForLabels(ctx context.Context, f func(key, value string) bool)
```

ForLabels invokes f with each label set on the context. The function f should return true to continue iteration or false to stop iteration early.

​	ForLabels 使用上下文中的每个标签集调用 f。函数 f 应返回 true 以继续迭代，或返回 false 以提前停止迭代。

### func Label <- go1.9

```go
func Label(ctx context.Context, key string) (string, bool)
```

Label returns the value of the label with the given key on ctx, and a boolean indicating whether that label exists.

​	Label 返回 ctx 上具有给定键的标签的值，以及一个布尔值，指示该标签是否存在。

### func SetGoroutineLabels <- go1.9

```go
func SetGoroutineLabels(ctx context.Context)
```

SetGoroutineLabels sets the current goroutine’s labels to match ctx. A new goroutine inherits the labels of the goroutine that created it. This is a lower-level API than Do, which should be used instead when possible.

​	SetGoroutineLabels 将当前协程的标签设置为与 ctx 匹配。新协程继承创建它的协程的标签。这是一个比 Do 更底层的 API，在可能的情况下应使用 Do。

### func StartCPUProfile

```go
func StartCPUProfile(w io.Writer) error
```

StartCPUProfile enables CPU profiling for the current process. While profiling, the profile will be buffered and written to w. StartCPUProfile returns an error if profiling is already enabled.

​	StartCPUProfile 为当前进程启用 CPU 分析。在分析时，分析结果将被缓冲并写入 w。如果已启用分析，StartCPUProfile 将返回一个错误。

On Unix-like systems, StartCPUProfile does not work by default for Go code built with -buildmode=c-archive or -buildmode=c-shared. StartCPUProfile relies on the SIGPROF signal, but that signal will be delivered to the main program’s SIGPROF signal handler (if any) not to the one used by Go. To make it work, call os/signal.Notify for syscall.SIGPROF, but note that doing so may break any profiling being done by the main program.

​	在类似 Unix 的系统上，对于使用 -buildmode=c-archive 或 -buildmode=c-shared 构建的 Go 代码，StartCPUProfile 默认情况下不起作用。StartCPUProfile 依赖于 SIGPROF 信号，但该信号将传递给主程序的 SIGPROF 信号处理程序（如果有），而不是 Go 使用的信号处理程序。要使其起作用，请为 syscall.SIGPROF 调用 os/signal.Notify，但请注意这样做可能会破坏主程序执行的任何分析。

### func StopCPUProfile

```go
func StopCPUProfile()
```

StopCPUProfile stops the current CPU profile, if any. StopCPUProfile only returns after all the writes for the profile have completed.

​	StopCPUProfile 停止当前的 CPU 分析（如果有）。StopCPUProfile 仅在分析的所有写入完成后才返回。

### func WithLabels <- go1.9

```go
func WithLabels(ctx context.Context, labels LabelSet) context.Context
```

WithLabels returns a new context.Context with the given labels added. A label overwrites a prior label with the same key.

​	WithLabels 返回一个新的 context.Context，其中添加了给定的标签。标签会覆盖具有相同键的先前标签。

### func WriteHeapProfile

```go
func WriteHeapProfile(w io.Writer) error
```

WriteHeapProfile is shorthand for Lookup(“heap”).WriteTo(w, 0). It is preserved for backwards compatibility.

​	WriteHeapProfile 是 Lookup(“heap”).WriteTo(w, 0) 的简写。保留它是为了向后兼容。

## 类型

### type LabelSet <- go1.9

```go
type LabelSet struct {
	// contains filtered or unexported fields
}
```

LabelSet is a set of labels.

​	LabelSet 是一组标签。

#### func Labels <- go1.9

```go
func Labels(args ...string) LabelSet
```

Labels takes an even number of strings representing key-value pairs and makes a LabelSet containing them. A label overwrites a prior label with the same key. Currently only the CPU and goroutine profiles utilize any labels information. See https://golang.org/issue/23458 for details.

​	标签采用偶数个字符串来表示键值对，并生成包含它们的 LabelSet。标签会覆盖具有相同键的先前标签。目前只有 CPU 和 goroutine 配置文件使用任何标签信息。有关详细信息，请参阅 https://golang.org/issue/23458。

### type Profile

```go
type Profile struct {
	// contains filtered or unexported fields
}
```

A Profile is a collection of stack traces showing the call sequences that led to instances of a particular event, such as allocation. Packages can create and maintain their own profiles; the most common use is for tracking resources that must be explicitly closed, such as files or network connections.

​	Profile 是一个堆栈跟踪集合，显示导致特定事件（例如分配）实例的调用序列。软件包可以创建和维护自己的配置文件；最常见的用途是跟踪必须显式关闭的资源，例如文件或网络连接。

A Profile’s methods can be called from multiple goroutines simultaneously.

​	可以从多个 goroutine 同时调用 Profile 的方法。

Each Profile has a unique name. A few profiles are predefined:

​	每个配置文件都有一个唯一名称。几个配置文件是预定义的：

```
goroutine    - stack traces of all current goroutines
heap         - a sampling of memory allocations of live objects
allocs       - a sampling of all past memory allocations
threadcreate - stack traces that led to the creation of new OS threads
block        - stack traces that led to blocking on synchronization primitives
mutex        - stack traces of holders of contended mutexes
```

These predefined profiles maintain themselves and panic on an explicit Add or Remove method call.

​	这些预定义的配置文件会维护自身，并在显式 Add 或 Remove 方法调用时引发恐慌。

The heap profile reports statistics as of the most recently completed garbage collection; it elides more recent allocation to avoid skewing the profile away from live data and toward garbage. If there has been no garbage collection at all, the heap profile reports all known allocations. This exception helps mainly in programs running without garbage collection enabled, usually for debugging purposes.

​	堆配置文件报告最近完成的垃圾回收的统计信息；它会忽略最近的分配，以避免将配置文件偏离活动数据而偏向垃圾。如果根本没有垃圾回收，则堆配置文件会报告所有已知分配。此异常主要有助于在未启用垃圾回收的情况下运行的程序，通常用于调试目的。

The heap profile tracks both the allocation sites for all live objects in the application memory and for all objects allocated since the program start. Pprof’s -inuse_space, -inuse_objects, -alloc_space, and -alloc_objects flags select which to display, defaulting to -inuse_space (live objects, scaled by size).

​	堆配置文件会跟踪应用程序内存中所有活动对象以及自程序启动以来分配的所有对象的分配站点。Pprof 的 -inuse_space、-inuse_objects、-alloc_space 和 -alloc_objects 标志选择要显示的内容，默认为 -inuse_space（活动对象，按大小缩放）。

The allocs profile is the same as the heap profile but changes the default pprof display to -alloc_space, the total number of bytes allocated since the program began (including garbage-collected bytes).

​	allocs 配置文件与堆配置文件相同，但将默认 pprof 显示更改为 -alloc_space，即自程序开始以来分配的总字节数（包括垃圾回收的字节）。

The CPU profile is not available as a Profile. It has a special API, the StartCPUProfile and StopCPUProfile functions, because it streams output to a writer during profiling.

​	CPU 配置文件不可用作配置文件。它有一个特殊的 API，即 StartCPUProfile 和 StopCPUProfile 函数，因为它在分析期间将输出流式传输到编写器。

#### func Lookup

```go
func Lookup(name string) *Profile
```

Lookup returns the profile with the given name, or nil if no such profile exists.

​	Lookup 返回具有给定名称的配置文件，如果不存在此类配置文件，则返回 nil。

#### func NewProfile

```go
func NewProfile(name string) *Profile
```

NewProfile creates a new profile with the given name. If a profile with that name already exists, NewProfile panics. The convention is to use a ‘import/path.’ prefix to create separate name spaces for each package. For compatibility with various tools that read pprof data, profile names should not contain spaces.

​	NewProfile 创建具有给定名称的新配置文件。如果已存在具有该名称的配置文件，则 NewProfile 会引发 panic。惯例是使用“import/path.”前缀为每个包创建单独的名称空间。为了与读取 pprof 数据的各种工具兼容，配置文件名称不应包含空格。

#### func Profiles

```go
func Profiles() []*Profile
```

Profiles returns a slice of all the known profiles, sorted by name.

​	Profiles 返回按名称排序的所有已知配置文件的切片。

#### (*Profile) Add

```go
func (p *Profile) Add(value any, skip int)
```

Add adds the current execution stack to the profile, associated with value. Add stores value in an internal map, so value must be suitable for use as a map key and will not be garbage collected until the corresponding call to Remove. Add panics if the profile already contains a stack for value.

​	Add 将当前执行堆栈添加到配置文件，并与值相关联。Add 将值存储在内部映射中，因此值必须适合用作映射键，并且在对 Remove 的相应调用之前不会被垃圾回收。如果配置文件已包含值的堆栈，则 Add 会引发 panic。

The skip parameter has the same meaning as runtime.Caller’s skip and controls where the stack trace begins. Passing skip=0 begins the trace in the function calling Add. For example, given this execution stack:

​	skip 参数与 runtime.Caller 的 skip 具有相同的含义，并控制堆栈跟踪的开始位置。传递 skip=0 会在调用 Add 的函数中开始跟踪。例如，给定此执行堆栈：

```
Add
called from rpc.NewClient
called from mypkg.Run
called from main.main
```

Passing skip=0 begins the stack trace at the call to Add inside rpc.NewClient. Passing skip=1 begins the stack trace at the call to NewClient inside mypkg.Run.

​	传递 skip=0 会在 rpc.NewClient 中对 Add 的调用处开始堆栈跟踪。传递 skip=1 会在 mypkg.Run 中对 NewClient 的调用处开始堆栈跟踪。

#### (*Profile) Count

```go
func (p *Profile) Count() int
```

Count returns the number of execution stacks currently in the profile.

​	Count 返回当前配置文件中的执行堆栈数。

#### (*Profile) Name

```go
func (p *Profile) Name() string
```

Name returns this profile’s name, which can be passed to Lookup to reobtain the profile.

​	Name 返回此配置文件的名称，该名称可以传递给 Lookup 以重新获取配置文件。

#### (*Profile) Remove

```go
func (p *Profile) Remove(value any)
```

Remove removes the execution stack associated with value from the profile. It is a no-op if the value is not in the profile.

​	Remove 从配置文件中删除与值关联的执行堆栈。如果值不在配置文件中，则它是一个空操作。

#### (*Profile) WriteTo

```go
func (p *Profile) WriteTo(w io.Writer, debug int) error
```

WriteTo writes a pprof-formatted snapshot of the profile to w. If a write to w returns an error, WriteTo returns that error. Otherwise, WriteTo returns nil.

​	WriteTo 将配置文件的 pprof 格式快照写入 w。如果对 w 的写入返回错误，WriteTo 返回该错误。否则，WriteTo 返回 nil。

The debug parameter enables additional output. Passing debug=0 writes the gzip-compressed protocol buffer described in https://github.com/google/pprof/tree/master/proto#overview. Passing debug=1 writes the legacy text format with comments translating addresses to function names and line numbers, so that a programmer can read the profile without tools.

​	debug 参数启用其他输出。传递 debug=0 将写入 https://github.com/google/pprof/tree/master/proto#overview 中描述的 gzip 压缩协议缓冲区。传递 debug=1 将写入带有注释的旧文本格式，将地址转换为函数名称和行号，以便程序员无需工具即可读取配置文件。

The predefined profiles may assign meaning to other debug values; for example, when printing the “goroutine” profile, debug=2 means to print the goroutine stacks in the same form that a Go program uses when dying due to an unrecovered panic.

​	预定义的配置文件可能会为其他 debug 值指定含义；例如，在打印“goroutine”配置文件时，debug=2 表示以 Go 程序在因未恢复的 panic 而终止时使用的相同形式打印 goroutine 堆栈。

## Notes

## Bugs

- Profiles are only as good as the kernel support used to generate them. See https://golang.org/issue/13841 for details about known problems.
- 配置文件的好坏取决于用于生成它们的内核支持。有关已知问题详情，请参阅 https://golang.org/issue/13841。