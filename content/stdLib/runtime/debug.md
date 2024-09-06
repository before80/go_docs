+++
title = "debug"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/runtime/debug@go1.23.0](https://pkg.go.dev/runtime/debug@go1.23.0)

Package debug contains facilities for programs to debug themselves while they are running.

​	debug 包含程序在运行时调试自身的工具。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func FreeOSMemory <- go1.1

```go
func FreeOSMemory()
```

FreeOSMemory forces a garbage collection followed by an attempt to return as much memory to the operating system as possible. (Even if this is not called, the runtime gradually returns memory to the operating system in a background task.)

​	FreeOSMemory 强制进行垃圾回收，然后尝试尽可能多地将内存返回给操作系统。（即使不调用此函数，运行时也会在后台任务中逐渐将内存返回给操作系统。）

### func PrintStack

```go
func PrintStack()
```

PrintStack prints to standard error the stack trace returned by runtime.Stack.

​	PrintStack 将 runtime.Stack 返回的堆栈跟踪打印到标准错误输出。

### func ReadGCStats <- go1.1

```go
func ReadGCStats(stats *GCStats)
```

ReadGCStats reads statistics about garbage collection into stats. The number of entries in the pause history is system-dependent; stats.Pause slice will be reused if large enough, reallocated otherwise. ReadGCStats may use the full capacity of the stats.Pause slice. If stats.PauseQuantiles is non-empty, ReadGCStats fills it with quantiles summarizing the distribution of pause time. For example, if len(stats.PauseQuantiles) is 5, it will be filled with the minimum, 25%, 50%, 75%, and maximum pause times.

​	ReadGCStats 将有关垃圾回收的统计信息读入 stats。暂停历史记录中的条目数取决于系统；如果 stats.Pause 切片足够大，则会重复使用，否则会重新分配。ReadGCStats 可能会使用 stats.Pause 切片的全部容量。如果 stats.PauseQuantiles 不为空，ReadGCStats 会用概括暂停时间分布的分位数填充它。例如，如果 len(stats.PauseQuantiles) 为 5，它将填充最小值、25%、50%、75% 和最大暂停时间。

### func SetCrashOutput <- go1.23.0

``` go
func SetCrashOutput(f *os.File, opts CrashOptions) error
```

SetCrashOutput configures a single additional file where unhandled panics and other fatal errors are printed, in addition to standard error. There is only one additional file: calling SetCrashOutput again overrides any earlier call. SetCrashOutput duplicates f's file descriptor, so the caller may safely close f as soon as SetCrashOutput returns. To disable this additional crash output, call SetCrashOutput(nil). If called concurrently with a crash, some in-progress output may be written to the old file even after an overriding SetCrashOutput returns.

​	`SetCrashOutput` 配置一个额外的文件，用于输出未处理的 `panic` 和其他致命错误，除了标准错误输出之外。只有一个额外的文件：再次调用 `SetCrashOutput` 将覆盖之前的调用。`SetCrashOutput` 会复制文件描述符 `f`，因此调用方可以在 `SetCrashOutput` 返回后立即安全地关闭 `f`。要禁用此额外的崩溃输出，可以调用 `SetCrashOutput(nil)`。如果在崩溃时并发调用此函数，一些进行中的输出可能会写入旧文件，即使覆盖的 `SetCrashOutput` 已返回。

#### SetCrashOutput  Example (Monitor)

ExampleSetCrashOutput_monitor shows an example of using [debug.SetCrashOutput](https://pkg.go.dev/runtime/debug@go1.23.0#SetCrashOutput) to direct crashes to a "monitor" process, for automated crash reporting. The monitor is the same executable, invoked in a special mode indicated by an environment variable.

​	`ExampleSetCrashOutput_monitor` 展示了一个使用 [debug.SetCrashOutput](https://pkg.go.dev/runtime/debug@go1.23.0#SetCrashOutput) 的示例，将崩溃定向到一个“监视器”进程，用于自动化崩溃报告。监视器是同一个可执行文件，以环境变量指示的特殊模式运行。

```go
package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)

// ExampleSetCrashOutput_monitor shows an example of using
// [debug.SetCrashOutput] to direct crashes to a "monitor" process,
// for automated crash reporting. The monitor is the same executable,
// invoked in a special mode indicated by an environment variable.
// ExampleSetCrashOutput_monitor 展示了一个使用
// [debug.SetCrashOutput] 的示例，将崩溃定向到一个“监视器”进程，
// 用于自动化崩溃报告。监视器是同一个可执行文件，
// 以环境变量指示的特殊模式运行。
func main() {
	appmain()

	// This Example doesn't actually run as a test because its
	// purpose is to crash, so it has no "Output:" comment
	// within the function body.
	//
	// To observe the monitor in action, replace the entire text
	// of this comment with "Output:" and run this command:
	//
	//    $ go test -run=ExampleSetCrashOutput_monitor runtime/debug
	//    panic: oops
	//    ...stack...
	//    monitor: saved crash report at /tmp/10804884239807998216.crash
    // 此示例实际上不会作为测试运行，因为它的目的是崩溃，
	// 因此函数体中没有 "Output:" 注释。
	//
	// 要观察监视器的运行效果，可以将此注释的整个内容替换为 "Output:"，
	// 然后运行以下命令：
	//
	//    $ go test -run=ExampleSetCrashOutput_monitor runtime/debug
	//    panic: oops
	//    ...stack...
	//    monitor: saved crash report at /tmp/10804884239807998216.crash
}

// appmain represents the 'main' function of your application.
// appmain 表示应用程序的 'main' 函数。
func appmain() {
	monitor()

	// Run the application.
    // 运行应用程序。
	println("hello")
	panic("oops")
}

// monitor starts the monitor process, which performs automated
// crash reporting. Call this function immediately within main.
//
// This function re-executes the same executable as a child process,
// in a special mode. In that mode, the call to monitor will never
// return.
// monitor 启动监视器进程，执行自动化崩溃报告。
// 应在 main 函数中立即调用此函数。
//
// 此函数将以特殊模式重新执行同一个可执行文件作为子进程，
// 在这种模式下，调用 monitor 不会返回。
func monitor() {
	const monitorVar = "RUNTIME_DEBUG_MONITOR"
	if os.Getenv(monitorVar) != "" {
		// This is the monitor (child) process.
        // 这是监视器（子）进程。
		log.SetFlags(0)
		log.SetPrefix("monitor: ")

		crash, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("failed to read from input pipe: %v", err)
		}
		if len(crash) == 0 {
			// Parent process terminated without reporting a crash.
            // 父进程在没有报告崩溃的情况下终止。
			os.Exit(0)
		}

		// Save the crash report securely in the file system.
        // 将崩溃报告安全地保存到文件系统中。
		f, err := os.CreateTemp("", "*.crash")
		if err != nil {
			log.Fatal(err)
		}
		if _, err := f.Write(crash); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
		log.Fatalf("saved crash report at %s", f.Name())
	}

	// This is the application process.
	// Fork+exec the same executable in monitor mode.
    // 这是应用程序进程。
	// 以监视器模式 Fork+exec 同一个可执行文件。
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(exe, "-test.run=ExampleSetCrashOutput_monitor")
	cmd.Env = append(os.Environ(), monitorVar+"=1")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stderr
	pipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatalf("StdinPipe: %v", err)
	}
	debug.SetCrashOutput(pipe.(*os.File), debug.CrashOptions{}) // (this conversion is safe) （此转换是安全的）
	if err := cmd.Start(); err != nil {
		log.Fatalf("can't start monitor: %v", err)
	}
	// Now return and start the application proper...
    // 现在返回并开始实际的应用程序...
}
Output:
```



### func SetGCPercent <- go1.1

```go
func SetGCPercent(percent int) int
```

SetGCPercent sets the garbage collection target percentage: a collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. SetGCPercent returns the previous setting. The initial setting is the value of the GOGC environment variable at startup, or 100 if the variable is not set. This setting may be effectively reduced in order to maintain a memory limit. A negative percentage effectively disables garbage collection, unless the memory limit is reached. See SetMemoryLimit for more details.

​	SetGCPercent 设置垃圾回收目标百分比：当新分配的数据与上次回收后剩余的活动数据之比达到此百分比时，将触发回收。SetGCPercent 返回上一个设置。初始设置是启动时 GOGC 环境变量的值，如果未设置该变量，则为 100。可以有效地降低此设置以维持内存限制。负百分比有效地禁用垃圾回收，除非达到内存限制。有关更多详细信息，请参阅 SetMemoryLimit。

### func SetMaxStack <- go1.2

```go
func SetMaxStack(bytes int) int
```

SetMaxStack sets the maximum amount of memory that can be used by a single goroutine stack. If any goroutine exceeds this limit while growing its stack, the program crashes. SetMaxStack returns the previous setting. The initial setting is 1 GB on 64-bit systems, 250 MB on 32-bit systems. There may be a system-imposed maximum stack limit regardless of the value provided to SetMaxStack.

​	SetMaxStack 设置单个 goroutine 堆栈可使用的最大内存量。如果任何 goroutine 在增长其堆栈时超过此限制，则程序崩溃。SetMaxStack 返回上一个设置。初始设置在 64 位系统上为 1 GB，在 32 位系统上为 250 MB。无论提供给 SetMaxStack 的值如何，都可能存在系统强加的最大堆栈限制。

SetMaxStack is useful mainly for limiting the damage done by goroutines that enter an infinite recursion. It only limits future stack growth.

​	SetMaxStack 主要用于限制进入无限递归的 goroutine 所造成的损害。它仅限制将来的堆栈增长。

### func SetMaxThreads <- go1.2

```go
func SetMaxThreads(threads int) int
```

SetMaxThreads sets the maximum number of operating system threads that the Go program can use. If it attempts to use more than this many, the program crashes. SetMaxThreads returns the previous setting. The initial setting is 10,000 threads.

​	SetMaxThreads 设置 Go 程序可使用的操作系统线程的最大数量。如果尝试使用超过此数量的线程，程序将崩溃。SetMaxThreads 返回之前的设置。初始设置是 10,000 个线程。

The limit controls the number of operating system threads, not the number of goroutines. A Go program creates a new thread only when a goroutine is ready to run but all the existing threads are blocked in system calls, cgo calls, or are locked to other goroutines due to use of runtime.LockOSThread.

​	此限制控制操作系统线程的数量，而不是 goroutine 的数量。Go 程序仅在 goroutine 准备运行但所有现有线程都因系统调用、cgo 调用而被阻塞，或因使用 runtime.LockOSThread 而被锁定到其他 goroutine 时才创建新线程。

SetMaxThreads is useful mainly for limiting the damage done by programs that create an unbounded number of threads. The idea is to take down the program before it takes down the operating system.

​	SetMaxThreads 主要用于限制创建无限数量线程的程序造成的损害。其目的是在程序导致操作系统崩溃之前将其关闭。

### func SetMemoryLimit <- go1.19

```go
func SetMemoryLimit(limit int64) int64
```

SetMemoryLimit provides the runtime with a soft memory limit.

​	SetMemoryLimit 为运行时提供软内存限制。

The runtime undertakes several processes to try to respect this memory limit, including adjustments to the frequency of garbage collections and returning memory to the underlying system more aggressively. This limit will be respected even if GOGC=off (or, if SetGCPercent(-1) is executed).

​	运行时会执行多个进程来尝试遵守此内存限制，包括调整垃圾回收的频率以及更积极地将内存返回给底层系统。即使 GOGC=off（或执行 SetGCPercent(-1)），也会遵守此限制。

The input limit is provided as bytes, and includes all memory mapped, managed, and not released by the Go runtime. Notably, it does not account for space used by the Go binary and memory external to Go, such as memory managed by the underlying system on behalf of the process, or memory managed by non-Go code inside the same process. Examples of excluded memory sources include: OS kernel memory held on behalf of the process, memory allocated by C code, and memory mapped by syscall.Mmap (because it is not managed by the Go runtime).

​	输入限制以字节为单位提供，包括所有内存映射、管理和未由 Go 运行时释放的内存。值得注意的是，它不考虑 Go 二进制文件和 Go 外部内存使用的空间，例如由进程代表基础系统管理的内存，或由同一进程中的非 Go 代码管理的内存。排除的内存源示例包括：代表进程持有的操作系统内核内存、由 C 代码分配的内存以及由 syscall.Mmap 映射的内存（因为它不受 Go 运行时管理）。

More specifically, the following expression accurately reflects the value the runtime attempts to maintain as the limit:

​	更具体地说，以下表达式准确地反映了运行时尝试维持的限制值：

```
runtime.MemStats.Sys - runtime.MemStats.HeapReleased
```

or in terms of the runtime/metrics package:

​	或根据 runtime/metrics 包：

```
/memory/classes/total:bytes - /memory/classes/heap/released:bytes
```

A zero limit or a limit that’s lower than the amount of memory used by the Go runtime may cause the garbage collector to run nearly continuously. However, the application may still make progress.

​	零限制或低于 Go 运行时使用的内存量的限制可能会导致垃圾回收器几乎连续运行。但是，应用程序可能仍然会取得进展。

The memory limit is always respected by the Go runtime, so to effectively disable this behavior, set the limit very high. math.MaxInt64 is the canonical value for disabling the limit, but values much greater than the available memory on the underlying system work just as well.

​	Go 运行时始终会遵守内存限制，因此要有效地禁用此行为，请将限制设置得非常高。math.MaxInt64 是禁用限制的规范值，但远大于基础系统上可用内存的值也同样有效。

See https://go.dev/doc/gc-guide for a detailed guide explaining the soft memory limit in more detail, as well as a variety of common use-cases and scenarios.

​	有关软内存限制的详细说明，以及各种常见用例和场景，请参阅 https://go.dev/doc/gc-guide。

The initial setting is math.MaxInt64 unless the GOMEMLIMIT environment variable is set, in which case it provides the initial setting. GOMEMLIMIT is a numeric value in bytes with an optional unit suffix. The supported suffixes include B, KiB, MiB, GiB, and TiB. These suffixes represent quantities of bytes as defined by the IEC 80000-13 standard. That is, they are based on powers of two: KiB means 2^10 bytes, MiB means 2^20 bytes, and so on.

​	初始设置是 math.MaxInt64，除非设置了 GOMEMLIMIT 环境变量，在这种情况下，它提供初始设置。GOMEMLIMIT 是一个以字节为单位的数字值，带有可选的单位后缀。支持的后缀包括 B、KiB、MiB、GiB 和 TiB。这些后缀表示由 IEC 80000-13 标准定义的字节数量。也就是说，它们基于 2 的幂：KiB 表示 2^10 字节，MiB 表示 2^20 字节，依此类推。

SetMemoryLimit returns the previously set memory limit. A negative input does not adjust the limit, and allows for retrieval of the currently set memory limit.

​	SetMemoryLimit 返回先前设置的内存限制。负输入不会调整限制，并允许检索当前设置的内存限制。

### func SetPanicOnFault <- go1.3

```go
func SetPanicOnFault(enabled bool) bool
```

SetPanicOnFault controls the runtime’s behavior when a program faults at an unexpected (non-nil) address. Such faults are typically caused by bugs such as runtime memory corruption, so the default response is to crash the program. Programs working with memory-mapped files or unsafe manipulation of memory may cause faults at non-nil addresses in less dramatic situations; SetPanicOnFault allows such programs to request that the runtime trigger only a panic, not a crash. The runtime.Error that the runtime panics with may have an additional method:

​	SetPanicOnFault 控制程序在意外（非 nil）地址发生故障时的运行时行为。此类故障通常由错误引起，例如运行时内存损坏，因此默认响应是使程序崩溃。使用内存映射文件或不安全内存操作的程序可能会在不太严重的情况下导致非 nil 地址发生故障；SetPanicOnFault 允许此类程序请求运行时仅触发 panic，而不是崩溃。运行时引发 panic 的 runtime.Error 可能具有其他方法：

```
Addr() uintptr
```

If that method exists, it returns the memory address which triggered the fault. The results of Addr are best-effort and the veracity of the result may depend on the platform. SetPanicOnFault applies only to the current goroutine. It returns the previous setting.

​	如果该方法存在，它将返回触发故障的内存地址。Addr 的结果是尽力而为，结果的真实性可能取决于平台。SetPanicOnFault 仅适用于当前 goroutine。它返回先前的设置。

### func SetTraceback <- go1.6

```go
func SetTraceback(level string)
```

SetTraceback sets the amount of detail printed by the runtime in the traceback it prints before exiting due to an unrecovered panic or an internal runtime error. The level argument takes the same values as the GOTRACEBACK environment variable. For example, SetTraceback(“all”) ensure that the program prints all goroutines when it crashes. See the package runtime documentation for details. If SetTraceback is called with a level lower than that of the environment variable, the call is ignored.

​	SetTraceback 设置运行时在因无法恢复的 panic 或内部运行时错误而退出之前打印的回溯的详细信息量。level 参数采用与 GOTRACEBACK 环境变量相同的值。例如，SetTraceback(“all”) 确保程序在崩溃时打印所有 goroutine。有关详细信息，请参阅 package runtime 文档。如果 SetTraceback 被调用时使用的级别低于环境变量的级别，则忽略该调用。

### func Stack

```go
func Stack() []byte
```

Stack returns a formatted stack trace of the goroutine that calls it. It calls runtime.Stack with a large enough buffer to capture the entire trace.

​	Stack 返回调用它的 goroutine 的格式化堆栈跟踪。它使用足够大的缓冲区调用 runtime.Stack 来捕获整个跟踪。

### func WriteHeapDump <- go1.3

```go
func WriteHeapDump(fd uintptr)
```

WriteHeapDump writes a description of the heap and the objects in it to the given file descriptor.

​	WriteHeapDump 将堆及其中的对象描述写入给定的文件描述符。

WriteHeapDump suspends the execution of all goroutines until the heap dump is completely written. Thus, the file descriptor must not be connected to a pipe or socket whose other end is in the same Go process; instead, use a temporary file or network socket.

​	WriteHeapDump 暂停所有 goroutine 的执行，直到堆转储完全写入。因此，文件描述符不得连接到另一端在同一 Go 进程中的管道或套接字；而应使用临时文件或网络套接字。

The heap dump format is defined at https://golang.org/s/go15heapdump.

​	堆转储格式定义在 https://golang.org/s/go15heapdump。

## 类型

### type BuildInfo <- go1.12

```go
type BuildInfo struct {
	// GoVersion is the version of the Go toolchain that built the binary
	// (for example, "go1.19.2").
	GoVersion string

	// Path is the package path of the main package for the binary
	// (for example, "golang.org/x/tools/cmd/stringer").
	Path string

	// Main describes the module that contains the main package for the binary.
	Main Module

	// Deps describes all the dependency modules, both direct and indirect,
	// that contributed packages to the build of this binary.
	Deps []*Module

	// Settings describes the build settings used to build the binary.
	Settings []BuildSetting
}
```

BuildInfo represents the build information read from a Go binary.

​	BuildInfo 表示从 Go 二进制文件中读取的构建信息。

#### func ParseBuildInfo <- go1.18

```go
func ParseBuildInfo(data string) (bi *BuildInfo, err error)
```

#### func ReadBuildInfo <- go1.12

```go
func ReadBuildInfo() (info *BuildInfo, ok bool)
```

ReadBuildInfo returns the build information embedded in the running binary. The information is available only in binaries built with module support.

​	ReadBuildInfo 返回嵌入在正在运行的二进制文件中的构建信息。该信息仅在使用模块支持构建的二进制文件中可用。

#### (*BuildInfo) String <- go1.18

```go
func (bi *BuildInfo) String() string
```

### type BuildSetting <- go1.18

```go
type BuildSetting struct {
	// Key and Value describe the build setting.
	// Key must not contain an equals sign, space, tab, or newline.
	// Value must not contain newlines ('\n').
	Key, Value string
}
```

A BuildSetting is a key-value pair describing one setting that influenced a build.

​	BuildSetting 是一个键值对，描述影响构建的一个设置。

Defined keys include:

​	定义的键包括：

- -buildmode: the buildmode flag used (typically “exe”)
  -buildmode：使用的构建模式标志（通常为“exe”）
- -compiler: the compiler toolchain flag used (typically “gc”)
  -compiler：使用的编译器工具链标志（通常为“gc”）
- CGO_ENABLED: the effective CGO_ENABLED environment variable
  CGO_ENABLED：有效的 CGO_ENABLED 环境变量
- CGO_CFLAGS: the effective CGO_CFLAGS environment variable
  CGO_CFLAGS：有效的 CGO_CFLAGS 环境变量
- CGO_CPPFLAGS: the effective CGO_CPPFLAGS environment variable
  CGO_CPPFLAGS：有效的 CGO_CPPFLAGS 环境变量
- CGO_CXXFLAGS: the effective CGO_CPPFLAGS environment variable
  CGO_CXXFLAGS：有效的 CGO_CPPFLAGS 环境变量
- CGO_LDFLAGS: the effective CGO_CPPFLAGS environment variable
  CGO_LDFLAGS：有效的 CGO_CPPFLAGS 环境变量
- GOARCH: the architecture target
  GOARCH：体系结构目标
- GOAMD64/GOARM64/GO386/etc: the architecture feature level for GOARCH
  GOAMD64/GOARM64/GO386/等：GOARCH 的体系结构功能级别
- GOOS: the operating system target
  GOOS：操作系统目标
- vcs: the version control system for the source tree where the build ran
  vcs：运行构建的源树的版本控制系统
- vcs.revision: the revision identifier for the current commit or checkout
  vcs.revision: 当前提交或签出的修订标识符
- vcs.time: the modification time associated with vcs.revision, in RFC3339 format
  vcs.time: 与 vcs.revision 关联的修改时间，采用 RFC3339 格式
- vcs.modified: true or false indicating whether the source tree had local modifications
  vcs.modified: 指示源树是否有本地修改的 true 或 false

### type CrashOptions <- go1.23.0

``` go
type CrashOptions struct {
}
```

CrashOptions provides options that control the formatting of the fatal crash message.

​	`CrashOptions` 提供控制致命崩溃消息格式的选项。

### type GCStats <- go1.1

```go
type GCStats struct {
	LastGC         time.Time       // time of last collection
	NumGC          int64           // number of garbage collections
	PauseTotal     time.Duration   // total pause for all collections
	Pause          []time.Duration // pause history, most recent first
	PauseEnd       []time.Time     // pause end times history, most recent first
	PauseQuantiles []time.Duration
}
```

GCStats collect information about recent garbage collections.

​	GCStats 收集有关最近垃圾回收的信息。

### type Module <- go1.12

```go
type Module struct {
	Path    string  // module path
	Version string  // module version
	Sum     string  // checksum
	Replace *Module // replaced by this module
}
```

A Module describes a single module included in a build.

​	Module 描述了构建中包含的单个模块。