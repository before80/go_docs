+++
title = "debug"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# debug

https://pkg.go.dev/runtime/debug@go1.20.1



Package debug contains facilities for programs to debug themselves while they are running.



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func FreeOSMemory  <- go1.1

``` go 
func FreeOSMemory()
```

FreeOSMemory forces a garbage collection followed by an attempt to return as much memory to the operating system as possible. (Even if this is not called, the runtime gradually returns memory to the operating system in a background task.)

#### func PrintStack 

``` go 
func PrintStack()
```

PrintStack prints to standard error the stack trace returned by runtime.Stack.

#### func ReadGCStats  <- go1.1

``` go 
func ReadGCStats(stats *GCStats)
```

ReadGCStats reads statistics about garbage collection into stats. The number of entries in the pause history is system-dependent; stats.Pause slice will be reused if large enough, reallocated otherwise. ReadGCStats may use the full capacity of the stats.Pause slice. If stats.PauseQuantiles is non-empty, ReadGCStats fills it with quantiles summarizing the distribution of pause time. For example, if len(stats.PauseQuantiles) is 5, it will be filled with the minimum, 25%, 50%, 75%, and maximum pause times.

#### func SetGCPercent  <- go1.1

``` go 
func SetGCPercent(percent int) int
```

SetGCPercent sets the garbage collection target percentage: a collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. SetGCPercent returns the previous setting. The initial setting is the value of the GOGC environment variable at startup, or 100 if the variable is not set. This setting may be effectively reduced in order to maintain a memory limit. A negative percentage effectively disables garbage collection, unless the memory limit is reached. See SetMemoryLimit for more details.

#### func SetMaxStack  <- go1.2

``` go 
func SetMaxStack(bytes int) int
```

SetMaxStack sets the maximum amount of memory that can be used by a single goroutine stack. If any goroutine exceeds this limit while growing its stack, the program crashes. SetMaxStack returns the previous setting. The initial setting is 1 GB on 64-bit systems, 250 MB on 32-bit systems. There may be a system-imposed maximum stack limit regardless of the value provided to SetMaxStack.

SetMaxStack is useful mainly for limiting the damage done by goroutines that enter an infinite recursion. It only limits future stack growth.

#### func SetMaxThreads  <- go1.2

``` go 
func SetMaxThreads(threads int) int
```

SetMaxThreads sets the maximum number of operating system threads that the Go program can use. If it attempts to use more than this many, the program crashes. SetMaxThreads returns the previous setting. The initial setting is 10,000 threads.

The limit controls the number of operating system threads, not the number of goroutines. A Go program creates a new thread only when a goroutine is ready to run but all the existing threads are blocked in system calls, cgo calls, or are locked to other goroutines due to use of runtime.LockOSThread.

SetMaxThreads is useful mainly for limiting the damage done by programs that create an unbounded number of threads. The idea is to take down the program before it takes down the operating system.

#### func SetMemoryLimit  <- go1.19

``` go 
func SetMemoryLimit(limit int64) int64
```

SetMemoryLimit provides the runtime with a soft memory limit.

The runtime undertakes several processes to try to respect this memory limit, including adjustments to the frequency of garbage collections and returning memory to the underlying system more aggressively. This limit will be respected even if GOGC=off (or, if SetGCPercent(-1) is executed).

The input limit is provided as bytes, and includes all memory mapped, managed, and not released by the Go runtime. Notably, it does not account for space used by the Go binary and memory external to Go, such as memory managed by the underlying system on behalf of the process, or memory managed by non-Go code inside the same process. Examples of excluded memory sources include: OS kernel memory held on behalf of the process, memory allocated by C code, and memory mapped by syscall.Mmap (because it is not managed by the Go runtime).

More specifically, the following expression accurately reflects the value the runtime attempts to maintain as the limit:

```
runtime.MemStats.Sys - runtime.MemStats.HeapReleased
```

or in terms of the runtime/metrics package:

```
/memory/classes/total:bytes - /memory/classes/heap/released:bytes
```

A zero limit or a limit that's lower than the amount of memory used by the Go runtime may cause the garbage collector to run nearly continuously. However, the application may still make progress.

The memory limit is always respected by the Go runtime, so to effectively disable this behavior, set the limit very high. math.MaxInt64 is the canonical value for disabling the limit, but values much greater than the available memory on the underlying system work just as well.

See https://go.dev/doc/gc-guide for a detailed guide explaining the soft memory limit in more detail, as well as a variety of common use-cases and scenarios.

The initial setting is math.MaxInt64 unless the GOMEMLIMIT environment variable is set, in which case it provides the initial setting. GOMEMLIMIT is a numeric value in bytes with an optional unit suffix. The supported suffixes include B, KiB, MiB, GiB, and TiB. These suffixes represent quantities of bytes as defined by the IEC 80000-13 standard. That is, they are based on powers of two: KiB means 2^10 bytes, MiB means 2^20 bytes, and so on.

SetMemoryLimit returns the previously set memory limit. A negative input does not adjust the limit, and allows for retrieval of the currently set memory limit.

#### func SetPanicOnFault  <- go1.3

``` go 
func SetPanicOnFault(enabled bool) bool
```

SetPanicOnFault controls the runtime's behavior when a program faults at an unexpected (non-nil) address. Such faults are typically caused by bugs such as runtime memory corruption, so the default response is to crash the program. Programs working with memory-mapped files or unsafe manipulation of memory may cause faults at non-nil addresses in less dramatic situations; SetPanicOnFault allows such programs to request that the runtime trigger only a panic, not a crash. The runtime.Error that the runtime panics with may have an additional method:

```
Addr() uintptr
```

If that method exists, it returns the memory address which triggered the fault. The results of Addr are best-effort and the veracity of the result may depend on the platform. SetPanicOnFault applies only to the current goroutine. It returns the previous setting.

#### func SetTraceback  <- go1.6

``` go 
func SetTraceback(level string)
```

SetTraceback sets the amount of detail printed by the runtime in the traceback it prints before exiting due to an unrecovered panic or an internal runtime error. The level argument takes the same values as the GOTRACEBACK environment variable. For example, SetTraceback("all") ensure that the program prints all goroutines when it crashes. See the package runtime documentation for details. If SetTraceback is called with a level lower than that of the environment variable, the call is ignored.

#### func Stack 

``` go 
func Stack() []byte
```

Stack returns a formatted stack trace of the goroutine that calls it. It calls runtime.Stack with a large enough buffer to capture the entire trace.

#### func WriteHeapDump  <- go1.3

``` go 
func WriteHeapDump(fd uintptr)
```

WriteHeapDump writes a description of the heap and the objects in it to the given file descriptor.

WriteHeapDump suspends the execution of all goroutines until the heap dump is completely written. Thus, the file descriptor must not be connected to a pipe or socket whose other end is in the same Go process; instead, use a temporary file or network socket.

The heap dump format is defined at https://golang.org/s/go15heapdump.

## 类型

### type BuildInfo  <- go1.12

``` go 
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

#### func ParseBuildInfo  <- go1.18

``` go 
func ParseBuildInfo(data string) (bi *BuildInfo, err error)
```

#### func ReadBuildInfo  <- go1.12

``` go 
func ReadBuildInfo() (info *BuildInfo, ok bool)
```

ReadBuildInfo returns the build information embedded in the running binary. The information is available only in binaries built with module support.

#### (*BuildInfo) String  <- go1.18

``` go 
func (bi *BuildInfo) String() string
```

### type BuildSetting  <- go1.18

``` go 
type BuildSetting struct {
	// Key and Value describe the build setting.
	// Key must not contain an equals sign, space, tab, or newline.
	// Value must not contain newlines ('\n').
	Key, Value string
}
```

A BuildSetting is a key-value pair describing one setting that influenced a build.

Defined keys include:

- -buildmode: the buildmode flag used (typically "exe")
- -compiler: the compiler toolchain flag used (typically "gc")
- CGO_ENABLED: the effective CGO_ENABLED environment variable
- CGO_CFLAGS: the effective CGO_CFLAGS environment variable
- CGO_CPPFLAGS: the effective CGO_CPPFLAGS environment variable
- CGO_CXXFLAGS: the effective CGO_CPPFLAGS environment variable
- CGO_LDFLAGS: the effective CGO_CPPFLAGS environment variable
- GOARCH: the architecture target
- GOAMD64/GOARM64/GO386/etc: the architecture feature level for GOARCH
- GOOS: the operating system target
- vcs: the version control system for the source tree where the build ran
- vcs.revision: the revision identifier for the current commit or checkout
- vcs.time: the modification time associated with vcs.revision, in RFC3339 format
- vcs.modified: true or false indicating whether the source tree had local modifications

### type GCStats  <- go1.1

``` go 
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

### type Module  <- go1.12

``` go 
type Module struct {
	Path    string  // module path
	Version string  // module version
	Sum     string  // checksum
	Replace *Module // replaced by this module
}
```

A Module describes a single module included in a build.