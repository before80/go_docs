+++
title = "trace"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/runtime/trace@go1.24.2](https://pkg.go.dev/runtime/trace@go1.24.2)

Package trace contains facilities for programs to generate traces for the Go execution tracer.

​	trace 包包含程序为 Go 执行跟踪器生成跟踪的工具。

## Tracing runtime activities 跟踪运行时活动

The execution trace captures a wide range of execution events such as goroutine creation/blocking/unblocking, syscall enter/exit/block, GC-related events, changes of heap size, processor start/stop, etc. When CPU profiling is active, the execution tracer makes an effort to include those samples as well. A precise nanosecond-precision timestamp and a stack trace is captured for most events. The generated trace can be interpreted using `go tool trace`.

​	执行跟踪捕获各种执行事件，例如 goroutine 创建/阻塞/解除阻塞、syscall 进入/退出/阻塞、与 GC 相关的事件、堆大小更改、处理器启动/停止等。当 CPU 分析处于活动状态时，执行跟踪器也会努力包含这些样本。对于大多数事件，都会捕获精确的纳秒级时间戳和堆栈跟踪。可以使用 `go tool trace` 来解释生成的跟踪。

Support for tracing tests and benchmarks built with the standard testing package is built into `go test`. For example, the following command runs the test in the current directory and writes the trace file (trace.out).

​	对使用标准测试包构建的跟踪测试和基准测试的支持已内置于 `go test` 中。例如，以下命令运行当前目录中的测试并将跟踪文件 (trace.out) 写入。 此 runtime/trace 包提供 API，以便向独立程序添加等效的跟踪支持。请参阅演示如何使用此 API 来启用跟踪的示例。

```
go test -trace=trace.out
```

This runtime/trace package provides APIs to add equivalent tracing support to a standalone program. See the Example that demonstrates how to use this API to enable tracing.

​	还有一个用于跟踪数据的标准 HTTP 接口。添加以下行将在 /debug/pprof/trace URL 下安装一个处理程序，以下载实时跟踪：

There is also a standard HTTP interface to trace data. Adding the following line will install a handler under the /debug/pprof/trace URL to download a live trace:

```
import _ "net/http/pprof"
```

See the net/http/pprof package for more details about all of the debug endpoints installed by this import.

​	有关此导入安装的所有调试端点的更多详细信息，请参阅 net/http/pprof 包。

## User annotation 用户注释

Package trace provides user annotation APIs that can be used to log interesting events during execution.

​	Package trace 提供用户注释 API，可用于在执行期间记录有趣事件。

There are three types of user annotations: log messages, regions, and tasks.

​	有三种类型用户注释：日志消息、区域和任务。

Log emits a timestamped message to the execution trace along with additional information such as the category of the message and which goroutine called Log. The execution tracer provides UIs to filter and group goroutines using the log category and the message supplied in Log.

​	Log 会将带时间戳的消息连同其他信息（例如消息的类别和调用 Log 的 goroutine）一起发送到执行跟踪。执行跟踪器提供 UI，以便使用日志类别和 Log 中提供的消息来过滤和分组 goroutine。

A region is for logging a time interval during a goroutine’s execution. By definition, a region starts and ends in the same goroutine. Regions can be nested to represent subintervals. For example, the following code records four regions in the execution trace to trace the durations of sequential steps in a cappuccino making operation.

​	区域用于记录 goroutine 执行期间的时间间隔。根据定义，区域在同一个 goroutine 中开始和结束。区域可以嵌套以表示子间隔。例如，以下代码在执行跟踪中记录四个区域，以跟踪制作卡布奇诺咖啡的顺序步骤的持续时间。

```
trace.WithRegion(ctx, "makeCappuccino", func() {

   // orderID allows to identify a specific order
   // among many cappuccino order region records.
   trace.Log(ctx, "orderID", orderID)

   trace.WithRegion(ctx, "steamMilk", steamMilk)
   trace.WithRegion(ctx, "extractCoffee", extractCoffee)
   trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
})
```

A task is a higher-level component that aids tracing of logical operations such as an RPC request, an HTTP request, or an interesting local operation which may require multiple goroutines working together. Since tasks can involve multiple goroutines, they are tracked via a context.Context object. NewTask creates a new task and embeds it in the returned context.Context object. Log messages and regions are attached to the task, if any, in the Context passed to Log and WithRegion.

​	任务是一个高级组件，有助于跟踪逻辑操作，例如 RPC 请求、HTTP 请求或可能需要多个协程协同工作的有趣本地操作。由于任务可能涉及多个协程，因此它们通过 context.Context 对象进行跟踪。NewTask 创建一个新任务并将其嵌入返回的 context.Context 对象中。如果在传递给 Log 和 WithRegion 的 Context 中存在任务，则日志消息和区域将附加到该任务。

For example, assume that we decided to froth milk, extract coffee, and mix milk and coffee in separate goroutines. With a task, the trace tool can identify the goroutines involved in a specific cappuccino order.

​	例如，假设我们决定在单独的协程中打奶泡、提取咖啡并混合牛奶和咖啡。使用任务，跟踪工具可以识别参与特定卡布奇诺订单的协程。

```
ctx, task := trace.NewTask(ctx, "makeCappuccino")
trace.Log(ctx, "orderID", orderID)

milk := make(chan bool)
espresso := make(chan bool)

go func() {
        trace.WithRegion(ctx, "steamMilk", steamMilk)
        milk <- true
}()
go func() {
        trace.WithRegion(ctx, "extractCoffee", extractCoffee)
        espresso <- true
}()
go func() {
        defer task.End() // When assemble is done, the order is complete.
        <-espresso
        <-milk
        trace.WithRegion(ctx, "mixMilkCoffee", mixMilkCoffee)
}()
```

The trace tool computes the latency of a task by measuring the time between the task creation and the task end and provides latency distributions for each task type found in the trace.

​	跟踪工具通过测量任务创建和任务结束之间的时间来计算任务的延迟，并为跟踪中找到的每种任务类型提供延迟分布。

## Example

Example demonstrates the use of the trace package to trace the execution of a Go program. The trace output will be written to the file trace.out

​	示例演示了使用跟踪包来跟踪 Go 程序的执行。跟踪输出将被写入文件 trace.out

```go
package main

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
)

// Example demonstrates the use of the trace package to trace
// the execution of a Go program. The trace output will be
// written to the file trace.out
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	// your program here
	RunMyProgram()
}

func RunMyProgram() {
	fmt.Printf("this function will be traced")
}
```

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func IsEnabled <- go1.11

```go
func IsEnabled() bool
```

IsEnabled reports whether tracing is enabled. The information is advisory only. The tracing status may have changed by the time this function returns.

​	IsEnabled 报告是否启用了跟踪。该信息仅供参考。此函数返回时，跟踪状态可能已发生更改。

### func Log <- go1.11

```go
func Log(ctx context.Context, category, message string)
```

Log emits a one-off event with the given category and message. Category can be empty and the API assumes there are only a handful of unique categories in the system.

​	Log 发出一个一次性事件，其中包含给定的类别和消息。类别可以为空，API 假设系统中只有少数几个唯一类别。

### func Logf <- go1.11

```go
func Logf(ctx context.Context, category, format string, args ...any)
```

Logf is like Log, but the value is formatted using the specified format spec.

​	Logf 与 Log 类似，但使用指定格式规范格式化值。

### func Start

```go
func Start(w io.Writer) error
```

Start enables tracing for the current program. While tracing, the trace will be buffered and written to w. Start returns an error if tracing is already enabled.

​	Start 为当前程序启用跟踪。在跟踪期间，跟踪将被缓冲并写入 w。如果已启用跟踪，Start 将返回错误。

### func Stop

```go
func Stop()
```

Stop stops the current tracing, if any. Stop only returns after all the writes for the trace have completed.

​	Stop 停止当前跟踪（如果有）。Stop 仅在跟踪的所有写入完成后才返回。

### func WithRegion <- go1.11

```go
func WithRegion(ctx context.Context, regionType string, fn func())
```

WithRegion starts a region associated with its calling goroutine, runs fn, and then ends the region. If the context carries a task, the region is associated with the task. Otherwise, the region is attached to the background task.

​	WithRegion 启动与其调用协程相关联的区域，运行 fn，然后结束该区域。如果上下文包含任务，则该区域与该任务相关联。否则，该区域将附加到后台任务。

The regionType is used to classify regions, so there should be only a handful of unique region types.

​	regionType 用于对区域进行分类，因此应该只有少数几个唯一的区域类型。

## 类型

### type Region <- go1.11

```go
type Region struct {
	// contains filtered or unexported fields
}
```

Region is a region of code whose execution time interval is traced.

​	Region 是其执行时间间隔被跟踪的代码区域。

#### func StartRegion <- go1.11

```go
func StartRegion(ctx context.Context, regionType string) *Region
```

StartRegion starts a region and returns a function for marking the end of the region. The returned Region’s End function must be called from the same goroutine where the region was started. Within each goroutine, regions must nest. That is, regions started after this region must be ended before this region can be ended. Recommended usage is

​	StartRegion 启动一个区域并返回一个用于标记区域结束的函数。返回的 Region 的 End 函数必须在启动该区域的同一个协程中调用。在每个协程中，区域必须嵌套。也就是说，在此区域之后启动的区域必须在此区域结束之前结束。建议用法是

```
defer trace.StartRegion(ctx, "myTracedRegion").End()
```

#### (*Region) End <- go1.11

```go
func (r *Region) End()
```

End marks the end of the traced code region.

​	End 标记被跟踪的代码区域的结束。

### type Task <- go1.11

```go
type Task struct {
	// contains filtered or unexported fields
}
```

Task is a data type for tracing a user-defined, logical operation.

​	Task 是用于跟踪用户定义的逻辑操作的数据类型。

#### func NewTask <- go1.11

```go
func NewTask(pctx context.Context, taskType string) (ctx context.Context, task *Task)
```

NewTask creates a task instance with the type taskType and returns it along with a Context that carries the task. If the input context contains a task, the new task is its subtask.

​	NewTask 创建一个类型为 taskType 的任务实例，并返回该实例以及一个携带该任务的 Context。如果输入的 context 包含一个任务，则新任务是其子任务。

The taskType is used to classify task instances. Analysis tools like the Go execution tracer may assume there are only a bounded number of unique task types in the system.

​	taskType 用于对任务实例进行分类。Go 执行跟踪器等分析工具可能会假设系统中只有有限数量的唯一任务类型。

The returned end function is used to mark the task’s end. The trace tool measures task latency as the time between task creation and when the end function is called, and provides the latency distribution per task type. If the end function is called multiple times, only the first call is used in the latency measurement.

​	返回的 end 函数用于标记任务的结束。跟踪工具将任务延迟测量为任务创建与调用 end 函数之间的时间，并提供每个任务类型的延迟分布。如果多次调用 end 函数，则只有第一次调用用于延迟测量。

```
ctx, task := trace.NewTask(ctx, "awesomeTask")
trace.WithRegion(ctx, "preparation", prepWork)
// preparation of the task
go func() {  // continue processing the task in a separate goroutine.
    defer task.End()
    trace.WithRegion(ctx, "remainingWork", remainingWork)
}()
```

#### (*Task) End <- go1.11

```go
func (t *Task) End()
```

End marks the end of the operation represented by the Task.

​	End 标记 Task 表示的操作的结束。