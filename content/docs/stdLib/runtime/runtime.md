+++
title = "runtime"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# runtime

https://pkg.go.dev/runtime@go1.20.1

​	runtime包包含与Go运行时系统交互的操作，例如控制goroutine的函数。它还包括reflect包使用的低级类型信息；有关可编程接口的运行时类型系统，请参见reflect的文档。

## 环境变量

​	以下环境变量(`$name`或`%name%`，具体取决于主机操作系统)控制Go程序的运行时行为。它们的含义和用途可能会随版本发布而发生变化。

​	GOGC变量设置初始垃圾回收目标百分比。当新分配数据与上一次回收后剩余的活跃数据之比达到该百分比时，将触发垃圾回收。默认值为GOGC = 100。设置GOGC = off可以完全禁用垃圾收集器。[runtime/debug.SetGCPercent](https://pkg.go.dev/runtime/debug#SetGCPercent)允许在运行时更改此百分比。

​	GOMEMLIMIT变量为运行时设置软内存限制。此内存限制包括Go堆和运行时管理的所有其他内存，并排除外部内存源，例如二进制映射本身、其他语言中管理的内存以及代表Go程序的操作系统持有的内存。 GOMEMLIMIT是一个以字节为单位的数字值，具有可选的单位后缀。支持的后缀包括B、KiB、MiB、GiB和TiB。这些后缀表示IEC 80000-13标准定义的字节数量。也就是说，它们基于二的幂：KiB表示$2^{10}$字节，MiB表示$2^{20}$字节，依此类推。默认设置为math.MaxInt64，这实际上禁用了内存限制。[runtime/debug.SetMemoryLimit](https://pkg.go.dev/runtime/debug#SetMemoryLimit)允许在运行时更改此限制。

​	GODEBUG变量控制运行时内部的调试变量。它是一个逗号分隔的name=val 对列表，设置这些命名变量：

```
allocfreetrace：设置allocfreetrace = 1会导致对每个分配进行分析，并在每个对象的分配和释放时打印栈跟踪。

clobberfree：设置clobberfree = 1会导致垃圾收集器在释放对象时使用错误内容覆盖对象的内存内容。

cpu.*：cpu.all = off禁用所有可选指令集扩展的使用。

cpu.extension = off禁用来自指定指令集扩展的指令。
扩展名是指指令集扩展的小写名称，例如sse41或avx，
如在internal / cpu软件包中列出的那样。例如，cpu.avx = off禁用运行时检测并因此禁用AVX指令的使用。

cgocheck：将cgocheck = 0设置为禁用所有检查，以使使用cgo的程序包不正确地将Go指针传递给非Go代码。
将cgocheck = 1(默认值)设置为启用相对便宜的检查，可能会漏检一些错误。
将cgocheck = 2设置为启用昂贵的检查，不应该漏检任何错误，但会导致程序运行较慢。

efence：设置efence = 1会导致分配器以每个对象分配一个唯一页面并且地址永远不会回收的模式运行。

gccheckmark：将gccheckmark = 1设置为启用验证垃圾收集器的并发标记阶段，通过在停止世界时执行第二个标记传递来完成。
如果第二个传递发现无法在并发标记中找到的可达对象，则垃圾收集器将引发紧急情况。

gcpacertrace：设置gcpacertrace = 1会导致垃圾收集器打印有关并发调节器的内部状态的信息。

gcshrinkstackoff：设置gcshrinkstackoff = 1会禁止将goroutine移动到较小的栈上。
在此模式下，goroutine的栈只能增长。

gcstoptheworld：将gcstoptheworld = 1设置为禁用并发垃圾收集，使每个垃圾收集成为停止世界事件。设置gcstoptheworld = 2还会在垃圾收集完成后禁用并发扫描。

gctrace: 将 gctrace 设置为 1，会导致垃圾回收器在每次回收时，向标准错误输出一行信息，概述收集的内存量和暂停的时间长度。此行的格式可能会改变。
目前，格式如下：
gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # MB stacks, #MB globals, # P
其中各字段的含义如下：
	gc #：垃圾收集的次数，每次垃圾收集都会递增。
	@#s：程序启动以来的时间(秒)。
	#%：程序启动以来在垃圾收集上花费的时间的百分比。
	#+...+#：垃圾收集阶段的墙钟/CPU时间。
	#->#-># MB：垃圾收集开始时、结束时和活动堆的堆大小。
	# MB goal：堆大小目标。
	# MB stacks：估计可扫描的栈大小。
	# MB globals：可扫描的全局大小。
	# P：使用的处理器数。
	
各阶段为暂停全局暂停(STW)扫描终止、并发标记和扫描，以及 STW 标记终止。标记/扫描的 CPU 时间会被分解为辅助时间(在分配时执行 GC)、后台 GC 时间和空闲 GC 时间。
如果该行以"(forced)"结尾，则此垃圾收集是由运行时的 runtime.GC() 调用强制触发的。

harddecommit: 将 harddecommit 设置为 1，会导致返回给操作系统的内存也被移除保护。这是 Windows 上唯一的操作模式，但在其他平台上调试垃圾回收器相关问题时也有用。目前仅在 Linux 上支持。

inittrace: 将 inittrace 设置为 1，会导致运行时在每个具有 init 工作的包上向标准错误输出一行信息，概述执行时间和内存分配。对于作为插件加载的 inits 以及没有用户定义和编译器生成 init 工作的包，不会打印任何信息。此行的格式可能会改变。目前，格式如下：
init # @#ms, # ms clock, # bytes, # allocs
其中各字段的含义如下：
	init #：包名称。
    @#ms：init 开始时的时间(毫秒)。
    # clock：包初始化工作的墙钟时间。
    # bytes：在堆上分配的内存量。
    # allocs：堆分配次数。

madvdontneed: 将madvdontneed设置为0会在将内存返回给内核时，在Linux上使用MADV_FREE而不是MADV_DONTNEED。这样更有效率，但意味着RSS数值只有在操作系统处于内存压力下时才会下降。在BSD和Illumos/Solaris上，设置madvdontneed=1会使用MADV_DONTNEED而不是MADV_FREE。这样不太有效率，但会导致RSS数值更快下降。

memprofilerate: 设置memprofilerate=X会更新runtime.MemProfileRate的值。当设置为0时，内存分析被禁用。有关默认值，请参阅MemProfileRate的描述。

pagetrace: 将pagetrace=/path/to/file设置为一个文件路径，将会写出一个页面事件的追踪，可以使用x/debug/cmd/pagetrace工具进行查看、分析和可视化。使用GOEXPERIMENT=pagetrace构建您的程序以启用此功能。如果您的程序是一个setuid二进制文件，则不要启用此功能，因为在这种情况下它会引入安全风险。目前不支持Windows、plan9或js/wasm。为某些应用程序设置此选项可能会产生大量的追踪信息，因此请谨慎使用。

invalidptr: invalidptr=1(默认值)会导致垃圾收集器和堆栈复制器在指针类型位置发现无效指针值(例如1)时崩溃程序。将invalidptr设置为0会禁用此检查。这应仅用作暂时的诊断有错误的代码的解决方法。真正的解决方法是不要在指针类型位置存储整数。

sbrk: 将sbrk设置为1会将内存分配器和垃圾收集器替换为一个简单的分配器，它从操作系统获取内存，并且永远不会回收任何内存。

scavtrace: 设置scavtrace=1会导致运行时大致每个GC周期在标准错误流中发出一行摘要，总结了清扫器所做的工作量、返回给操作系统的总内存量以及物理内存利用率的估计值。此行的格式可能会发生变化，但目前的格式是：
scav # KiB work, # KiB total, #% util
其中字段如下：
	# KiB work 上次行以来返回给操作系统的内存量
	# KiB total 返回给操作系统的总内存量
	#% util 未清理的所有内存中正在使用的部分的比例
	如果该行以"(forced)"结尾，则是通过调用debug.FreeOSMemory()强制进行的清理。

scheddetail: 设置 schedtrace=X 和 scheddetail=1 会导致调度器每 X 毫秒发出详细的多行信息，描述调度器、处理器、线程和 goroutine 的状态。

schedtrace: 设置 schedtrace=X 会导致调度器每 X 毫秒向标准错误发出单行信息，概述调度器状态。

tracebackancestors: 设置 tracebackancestors=N 将追溯信息扩展到创建 goroutine 的栈，其中 N 限制要报告的祖先 goroutine 数量。这还扩展了由 runtime.Stack 返回的信息。祖先 goroutine 的 ID 将引用创建时 goroutine 的 ID；此 ID 可能会重用于另一个 goroutine。将 N 设置为 0 将不报告祖先信息。

asyncpreemptoff: asyncpreemptoff=1 禁用基于信号的异步 goroutine 抢占。这使某些循环在长时间内不可抢占，可能会延迟 GC 和 goroutine 调度。这对于调试 GC 问题非常有用，因为它还禁用了用于异步抢占的保守栈扫描。

```

​	net和net/http包也引用了GODEBUG中的调试变量。有关详细信息，请参阅这些包的文档。

​	GOMAXPROCS 变量限制了可以同时执行用户级 Go 代码的操作系统线程数量。在代表 Go 代码阻塞的系统调用中，线程数量没有限制；它们不计入 GOMAXPROCS 限制。本包的GOMAXPROCS函数用于查询和更改此限制。

​	GORACE 变量配置了竞争检测器，用于使用 -race 构建的程序。有关详细信息，请参阅 [https://golang.org/doc/articles/race_detector.html](https://golang.org/doc/articles/race_detector.html)。

​	GOTRACEBACK 变量控制在 Go 程序由于未恢复的 panic 或意外的运行时条件而失败时生成的输出量。默认情况下，失败会为当前 goroutine 打印栈跟踪，省略运行时系统内部的函数，然后以退出码 2 退出。如果没有当前 goroutine 或失败是运行时内部的，则失败会打印所有 goroutine 的栈跟踪。GOTRACEBACK=none 完全省略 goroutine 栈跟踪。GOTRACEBACK=single(默认值)的行为如上所述。GOTRACEBACK=all 为所有用户创建的 goroutine 添加栈跟踪。GOTRACEBACK=system 类似于 "all"，但为运行时函数添加堆栈帧，并显示由运行时内部创建的 goroutine。GOTRACEBACK=crash 类似于 "system"，但以特定于操作系统的方式崩溃而不是退出。例如，在 Unix 系统上，崩溃会引发 SIGABRT 来触发核心转储。出于历史原因，GOTRACEBACK 设置 0、1 和 2 分别是 none、all 和 system 的同义词。runtime/debug 包的 SetTraceback 函数允许在运行时增加输出量，但不能将输出量减少到低于环境变量指定的水平。请参见 [https://golang.org/pkg/runtime/debug/#SetTraceback](https://golang.org/pkg/runtime/debug/#SetTraceback)。

​	GOARCH、GOOS、GOPATH 和 GOROOT 环境变量完成了 Go 环境变量的设置。它们影响构建 Go 程序(请参见 [https://golang.org/cmd/go](https://golang.org/cmd/go) 和 [https://golang.org/pkg/go/build](https://golang.org/pkg/go/build))。GOARCH、GOOS 和 GOROOT 在编译时记录并通过常量或该包中的函数提供，但它们不影响运行时系统的执行。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/compiler.go;l=12)

``` go 
const Compiler = "gc"
```

​	Compiler 是编译生成运行二进制文件的编译器工具链的名称。已知的工具链包括：

```
gc 		也称 cmd/compile。
gccgo   gccgo 前端，是 GCC 编译器套件的一部分。
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/extern.go;l=303)

``` go 
const GOARCH string = goarch.GOARCH
```

​	GOARCH 是运行程序的体系结构目标，例如 386、amd64、arm、s390x 等。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/extern.go;l=299)

``` go 
const GOOS string = goos.GOOS
```

​	GOOS 是运行程序的操作系统目标，例如 darwin、freebsd、linux 等。要查看 GOOS 和 GOARCH 的可能组合，请运行"go tool dist list"。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/mprof.go;l=595)

``` go 
var MemProfileRate int = 512 * 1024
```

​	MemProfileRate 控制记录和报告内存分析中的内存分配的部分。分析器旨在对每个 MemProfileRate 分配的平均样本进行采样。

​	要在分析文件(profile)中包含每个已分配的块，请将 MemProfileRate 设置为 1。要完全关闭分析，请将 MemProfileRate 设置为 0。

​	处理内存分析的工具假设分析速率在程序的整个生命周期中是恒定的，并且等于当前值。更改内存分析速率的程序应该只在程序执行的尽早时期(例如在 main 的开始处)执行一次。

## 函数

#### func BlockProfile  <- go1.1

``` go 
func BlockProfile(p []BlockProfileRecord) (n int, ok bool)
```

​	BlockProfile函数返回当前阻塞分析中的记录数 n。如果 len(p) >= n，则 BlockProfile 将分析副本复制到 p 并返回 n、true。如果 len(p) < n，则 BlockProfile 不会更改 p 并返回 n、false。

​	大多数客户端应该使用 runtime/pprof 包或 testing 包的 -test.blockprofile 标志，而不是直接调用 BlockProfile函数。

#### func Breakpoint 

``` go 
func Breakpoint()
```

​	Breakpoint函数执行断点陷阱。

#### func Caller 

``` go 
func Caller(skip int) (pc uintptr, file string, line int, ok bool)
```

​	Caller函数报告关于调用 goroutine 栈上函数调用的文件和行号信息。skip 是要上升的栈帧数，其中 0 表示 Caller函数的调用者(由于历史原因，skip 在 Caller函数和 Callers函数之间的含义不同)。返回值报告相应调用的程序计数器、文件名和文件中的行号。如果无法恢复信息，则布尔值 ok 为 false。

#### func Callers 

``` go 
func Callers(skip int, pc []uintptr) int
```

​	Callers函数将调用当前goroutine的栈上函数调用的返回程序计数器填充到切片pc中。参数skip表示在记录pc之前要跳过的栈帧数，其中0标识Callers本身的帧，1标识Callers的调用者。它返回写入到pc的条目数。

​	要将这些程序计数器转换为符号信息，例如函数名称和行号，请使用CallersFrames函数。 CallersFrames函数考虑了内联函数并将返回程序计数器调整为调用程序计数器。不建议直接迭代返回的PCs切片，也不建议在任何返回的PC上使用FuncForPC函数，因为这些都无法考虑到内联或返回程序计数器的调整。

#### func GC 

``` go 
func GC()
```

​	GC函数运行垃圾回收并阻塞调用者，直到垃圾回收完成。它也可能阻止整个程序。

#### func GOMAXPROCS 

``` go 
func GOMAXPROCS(n int) int
```

​	GOMAXPROCS函数设置可以同时执行的最大CPU数量并返回先前的设置。默认值为runtime.NumCPU的值。如果n < 1，则不更改当前设置。当调度程序改进时，此调用将被取消(This call will go away when the scheduler improves.)。

#### func GOROOT 

``` go 
func GOROOT() string
```

​	GOROOT函数返回Go树的根。如果在进程启动时设置了GOROOT环境变量，则使用它，否则使用Go构建期间使用的根目录。

#### func Goexit 

``` go 
func Goexit()
```

​	Goexit函数终止调用它的goroutine。不会影响其他goroutine。Goexit函数在终止goroutine之前运行所有延迟调用。因为Goexit函数不是一个panic，所以这些延迟函数中的任何recover函数调用都将返回nil。

​	从主goroutine调用Goexit将终止该goroutine，而不是返回func main。由于func main没有返回，程序将继续执行其他goroutine。如果所有其他goroutine退出，则程序崩溃。

#### func GoroutineProfile 

``` go 
func GoroutineProfile(p []StackRecord) (n int, ok bool)
```

​	GoroutineProfile函数返回n，活动goroutine栈分析中记录的数量。如果len(p) >= n，则GoroutineProfile函数将分析复制到p中并返回n，true。如果len(p) < n，则GoroutineProfile不更改p并返回n，false。

​	大多数客户端应该使用runtime/pprof包而不是直接调用GoroutineProfile函数。

#### func Gosched 

``` go 
func Gosched()
```

​	Gosched函数让出处理器，允许其他goroutine运行。它不挂起当前的goroutine，因此执行将自动恢复。

#### func KeepAlive  <- go1.7

``` go 
func KeepAlive(x any)
```

​	KeepAlive函数将其参数标记为当前可访问。这确保在调用KeepAlive函数的程序点之前不会释放对象，也不会运行其finalizer(终结器)。

​	一个非常简化的例子展示了 KeepAlive 的使用情况：

``` go 
type File struct { d int }
d, err := syscall.Open("/file/path", syscall.O_RDONLY, 0)
// ... do something if err != nil ...
p := &File{d}
runtime.SetFinalizer(p, func(p *File) { syscall.Close(p.d) })
var buf [10]byte
n, err := syscall.Read(p.d, buf[:])
// 确保在 Read 返回之前，p 不被释放。
runtime.KeepAlive(p)
// 在此之后，p 不再被使用。
```

​	如果没有 KeepAlive函数的调用，finalizer(终结器)可能会在 syscall.Read 开始时运行，(在 syscall.Read 实际进行系统调用之前)关闭文件描述符。

注意：KeepAlive函数应该仅用于防止终结器过早运行。特别地，当与 unsafe.Pointer 一起使用时，仍然适用于 unsafe.Pointer 的有效使用规则。

#### func LockOSThread 

``` go 
func LockOSThread()
```

​	LockOSThread函数将调用它的 goroutine 绑定到其当前的操作系统线程。调用 goroutine 将始终在该线程中执行，并且没有其他 goroutine 将在其中执行，直到调用 goroutine 调用 UnlockOSThread 的次数与 LockOSThread 的次数相同。如果调用 goroutine 在不解锁线程的情况下退出，线程将被终止。

​	所有的 init 函数都在启动线程上运行。从 init 函数中调用 LockOSThread函数将导致在该线程上调用主函数。

​	在调用 OS 服务或依赖于每个线程状态的非 Go 库函数之前，goroutine 应该调用 LockOSThread函数。

#### func MemProfile 

``` go 
func MemProfile(p []MemProfileRecord, inuseZero bool) (n int, ok bool)
```

​	MemProfile函数返回每个分配点分配和释放的内存的分析。

​	MemProfile函数返回 n，当前内存分析中的记录数。

​	如果 len(p) >= n，则 MemProfile函数将分析复制到 p 并返回 n，true。

​	如果 len(p) < n，则 MemProfile函数不会改变 p，并返回 n，false。

​	如果 inuseZero参数 为 true，则分析包括 r.AllocBytes > 0 但 r.AllocBytes == r.FreeBytes 的分配记录。这些是分配了内存但已经全部释放回运行时的站点。

​	返回的分析结果可能是最多两个垃圾收集周期之前的。这是为了避免向分析结果倾斜分配的情况；由于分配是实时发生的，但释放需要等到垃圾收集器进行扫描，因此该分析结果仅记录那些已经有机会被垃圾收集器释放的分配情况。

​	大多数客户端应该使用 runtime/pprof 包或testing 包的 -test.memprofile 标志，而不是直接调用 MemProfile函数。

#### func MutexProfile  <- go1.8

``` go 
func MutexProfile(p []BlockProfileRecord) (n int, ok bool)
```

​	MutexProfile函数返回当前mutex profile中记录的数量n。如果len(p) >= n，则MutexProfile函数将profile复制到p中并返回n和true。否则，MutexProfile函数不会更改p，并返回n和false。

​	大多数客户端应该使用runtime/pprof包而不是直接调用MutexProfile函数。

#### func NumCPU 

``` go 
func NumCPU() int
```

​	NumCPU函数返回当前进程可用的逻辑 CPU 数量。

The set of available CPUs is checked by querying the operating system at process startup. Changes to operating system CPU allocation after process startup are not reflected.

​	可用的 CPU 集合通过在进程启动时查询操作系统进行检查。进程启动后对操作系统 CPU 分配的更改不会反映出来。

#### func NumCgoCall 

``` go 
func NumCgoCall() int64
```

NumCgoCall returns the number of cgo calls made by the current process.

#### func NumGoroutine 

``` go 
func NumGoroutine() int
```

NumGoroutine returns the number of goroutines that currently exist.

#### func ReadMemStats 

``` go 
func ReadMemStats(m *MemStats)
```

ReadMemStats populates m with memory allocator statistics.

The returned memory allocator statistics are up to date as of the call to ReadMemStats. This is in contrast with a heap profile, which is a snapshot as of the most recently completed garbage collection cycle.

#### func ReadTrace  <- go1.5

``` go 
func ReadTrace() []byte
```

ReadTrace returns the next chunk of binary tracing data, blocking until data is available. If tracing is turned off and all the data accumulated while it was on has been returned, ReadTrace returns nil. The caller must copy the returned data before calling ReadTrace again. ReadTrace must be called from one goroutine at a time.

#### func SetBlockProfileRate  <- go1.1

``` go 
func SetBlockProfileRate(rate int)
```

SetBlockProfileRate controls the fraction of goroutine blocking events that are reported in the blocking profile. The profiler aims to sample an average of one blocking event per rate nanoseconds spent blocked.

To include every blocking event in the profile, pass rate = 1. To turn off profiling entirely, pass rate <= 0.

#### func SetCPUProfileRate 

``` go 
func SetCPUProfileRate(hz int)
```

SetCPUProfileRate sets the CPU profiling rate to hz samples per second. If hz <= 0, SetCPUProfileRate turns off profiling. If the profiler is on, the rate cannot be changed without first turning it off.

Most clients should use the runtime/pprof package or the testing package's -test.cpuprofile flag instead of calling SetCPUProfileRate directly.

#### func SetCgoTraceback  <- go1.7

``` go 
func SetCgoTraceback(version int, traceback, context, symbolizer unsafe.Pointer)
```

SetCgoTraceback records three C functions to use to gather traceback information from C code and to convert that traceback information into symbolic information. These are used when printing stack traces for a program that uses cgo.

The traceback and context functions may be called from a signal handler, and must therefore use only async-signal safe functions. The symbolizer function may be called while the program is crashing, and so must be cautious about using memory. None of the functions may call back into Go.

The context function will be called with a single argument, a pointer to a struct:

```
struct {
	Context uintptr
}
```

In C syntax, this struct will be

```
struct {
	uintptr_t Context;
};
```

If the Context field is 0, the context function is being called to record the current traceback context. It should record in the Context field whatever information is needed about the current point of execution to later produce a stack trace, probably the stack pointer and PC. In this case the context function will be called from C code.

If the Context field is not 0, then it is a value returned by a previous call to the context function. This case is called when the context is no longer needed; that is, when the Go code is returning to its C code caller. This permits the context function to release any associated resources.

While it would be correct for the context function to record a complete a stack trace whenever it is called, and simply copy that out in the traceback function, in a typical program the context function will be called many times without ever recording a traceback for that context. Recording a complete stack trace in a call to the context function is likely to be inefficient.

The traceback function will be called with a single argument, a pointer to a struct:

```
struct {
	Context    uintptr
	SigContext uintptr
	Buf        *uintptr
	Max        uintptr
}
```

In C syntax, this struct will be

```
struct {
	uintptr_t  Context;
	uintptr_t  SigContext;
	uintptr_t* Buf;
	uintptr_t  Max;
};
```

The Context field will be zero to gather a traceback from the current program execution point. In this case, the traceback function will be called from C code.

Otherwise Context will be a value previously returned by a call to the context function. The traceback function should gather a stack trace from that saved point in the program execution. The traceback function may be called from an execution thread other than the one that recorded the context, but only when the context is known to be valid and unchanging. The traceback function may also be called deeper in the call stack on the same thread that recorded the context. The traceback function may be called multiple times with the same Context value; it will usually be appropriate to cache the result, if possible, the first time this is called for a specific context value.

If the traceback function is called from a signal handler on a Unix system, SigContext will be the signal context argument passed to the signal handler (a C ucontext_t* cast to uintptr_t). This may be used to start tracing at the point where the signal occurred. If the traceback function is not called from a signal handler, SigContext will be zero.

Buf is where the traceback information should be stored. It should be PC values, such that Buf[0] is the PC of the caller, Buf[1] is the PC of that function's caller, and so on. Max is the maximum number of entries to store. The function should store a zero to indicate the top of the stack, or that the caller is on a different stack, presumably a Go stack.

Unlike runtime.Callers, the PC values returned should, when passed to the symbolizer function, return the file/line of the call instruction. No additional subtraction is required or appropriate.

On all platforms, the traceback function is invoked when a call from Go to C to Go requests a stack trace. On linux/amd64, linux/ppc64le, linux/arm64, and freebsd/amd64, the traceback function is also invoked when a signal is received by a thread that is executing a cgo call. The traceback function should not make assumptions about when it is called, as future versions of Go may make additional calls.

The symbolizer function will be called with a single argument, a pointer to a struct:

```
struct {
	PC      uintptr // program counter to fetch information for
	File    *byte   // file name (NUL terminated)
	Lineno  uintptr // line number
	Func    *byte   // function name (NUL terminated)
	Entry   uintptr // function entry point
	More    uintptr // set non-zero if more info for this PC
	Data    uintptr // unused by runtime, available for function
}
```

In C syntax, this struct will be

```
struct {
	uintptr_t PC;
	char*     File;
	uintptr_t Lineno;
	char*     Func;
	uintptr_t Entry;
	uintptr_t More;
	uintptr_t Data;
};
```

The PC field will be a value returned by a call to the traceback function.

The first time the function is called for a particular traceback, all the fields except PC will be 0. The function should fill in the other fields if possible, setting them to 0/nil if the information is not available. The Data field may be used to store any useful information across calls. The More field should be set to non-zero if there is more information for this PC, zero otherwise. If More is set non-zero, the function will be called again with the same PC, and may return different information (this is intended for use with inlined functions). If More is zero, the function will be called with the next PC value in the traceback. When the traceback is complete, the function will be called once more with PC set to zero; this may be used to free any information. Each call will leave the fields of the struct set to the same values they had upon return, except for the PC field when the More field is zero. The function must not keep a copy of the struct pointer between calls.

When calling SetCgoTraceback, the version argument is the version number of the structs that the functions expect to receive. Currently this must be zero.

The symbolizer function may be nil, in which case the results of the traceback function will be displayed as numbers. If the traceback function is nil, the symbolizer function will never be called. The context function may be nil, in which case the traceback function will only be called with the context field set to zero. If the context function is nil, then calls from Go to C to Go will not show a traceback for the C portion of the call stack.

SetCgoTraceback should be called only once, ideally from an init function.

#### func SetFinalizer 

``` go 
func SetFinalizer(obj any, finalizer any)
```

SetFinalizer sets the finalizer associated with obj to the provided finalizer function. When the garbage collector finds an unreachable block with an associated finalizer, it clears the association and runs finalizer(obj) in a separate goroutine. This makes obj reachable again, but now without an associated finalizer. Assuming that SetFinalizer is not called again, the next time the garbage collector sees that obj is unreachable, it will free obj.

SetFinalizer(obj, nil) clears any finalizer associated with obj.

The argument obj must be a pointer to an object allocated by calling new, by taking the address of a composite literal, or by taking the address of a local variable. The argument finalizer must be a function that takes a single argument to which obj's type can be assigned, and can have arbitrary ignored return values. If either of these is not true, SetFinalizer may abort the program.

Finalizers are run in dependency order: if A points at B, both have finalizers, and they are otherwise unreachable, only the finalizer for A runs; once A is freed, the finalizer for B can run. If a cyclic structure includes a block with a finalizer, that cycle is not guaranteed to be garbage collected and the finalizer is not guaranteed to run, because there is no ordering that respects the dependencies.

The finalizer is scheduled to run at some arbitrary time after the program can no longer reach the object to which obj points. There is no guarantee that finalizers will run before a program exits, so typically they are useful only for releasing non-memory resources associated with an object during a long-running program. For example, an os.File object could use a finalizer to close the associated operating system file descriptor when a program discards an os.File without calling Close, but it would be a mistake to depend on a finalizer to flush an in-memory I/O buffer such as a bufio.Writer, because the buffer would not be flushed at program exit.

It is not guaranteed that a finalizer will run if the size of *obj is zero bytes, because it may share same address with other zero-size objects in memory. See https://go.dev/ref/spec#Size_and_alignment_guarantees.

It is not guaranteed that a finalizer will run for objects allocated in initializers for package-level variables. Such objects may be linker-allocated, not heap-allocated.

Note that because finalizers may execute arbitrarily far into the future after an object is no longer referenced, the runtime is allowed to perform a space-saving optimization that batches objects together in a single allocation slot. The finalizer for an unreferenced object in such an allocation may never run if it always exists in the same batch as a referenced object. Typically, this batching only happens for tiny (on the order of 16 bytes or less) and pointer-free objects.

A finalizer may run as soon as an object becomes unreachable. In order to use finalizers correctly, the program must ensure that the object is reachable until it is no longer required. Objects stored in global variables, or that can be found by tracing pointers from a global variable, are reachable. For other objects, pass the object to a call of the KeepAlive function to mark the last point in the function where the object must be reachable.

For example, if p points to a struct, such as os.File, that contains a file descriptor d, and p has a finalizer that closes that file descriptor, and if the last use of p in a function is a call to syscall.Write(p.d, buf, size), then p may be unreachable as soon as the program enters syscall.Write. The finalizer may run at that moment, closing p.d, causing syscall.Write to fail because it is writing to a closed file descriptor (or, worse, to an entirely different file descriptor opened by a different goroutine). To avoid this problem, call KeepAlive(p) after the call to syscall.Write.

A single goroutine runs all finalizers for a program, sequentially. If a finalizer must run for a long time, it should do so by starting a new goroutine.

In the terminology of the Go memory model, a call SetFinalizer(x, f) "synchronizes before" the finalization call f(x). However, there is no guarantee that KeepAlive(x) or any other use of x "synchronizes before" f(x), so in general a finalizer should use a mutex or other synchronization mechanism if it needs to access mutable state in x. For example, consider a finalizer that inspects a mutable field in x that is modified from time to time in the main program before x becomes unreachable and the finalizer is invoked. The modifications in the main program and the inspection in the finalizer need to use appropriate synchronization, such as mutexes or atomic updates, to avoid read-write races.

#### func SetMutexProfileFraction  <- go1.8

``` go 
func SetMutexProfileFraction(rate int) int
```

SetMutexProfileFraction controls the fraction of mutex contention events that are reported in the mutex profile. On average 1/rate events are reported. The previous rate is returned.

To turn off profiling entirely, pass rate 0. To just read the current rate, pass rate < 0. (For n>1 the details of sampling may change.)

#### func Stack 

``` go 
func Stack(buf []byte, all bool) int
```

Stack formats a stack trace of the calling goroutine into buf and returns the number of bytes written to buf. If all is true, Stack formats stack traces of all other goroutines into buf after the trace for the current goroutine.

#### func StartTrace  <- go1.5

``` go 
func StartTrace() error
```

StartTrace enables tracing for the current process. While tracing, the data will be buffered and available via ReadTrace. StartTrace returns an error if tracing is already enabled. Most clients should use the runtime/trace package or the testing package's -test.trace flag instead of calling StartTrace directly.

#### func StopTrace  <- go1.5

``` go 
func StopTrace()
```

StopTrace stops tracing, if it was previously enabled. StopTrace only returns after all the reads for the trace have completed.

#### func ThreadCreateProfile 

``` go 
func ThreadCreateProfile(p []StackRecord) (n int, ok bool)
```

ThreadCreateProfile returns n, the number of records in the thread creation profile. If len(p) >= n, ThreadCreateProfile copies the profile into p and returns n, true. If len(p) < n, ThreadCreateProfile does not change p and returns n, false.

Most clients should use the runtime/pprof package instead of calling ThreadCreateProfile directly.

#### func UnlockOSThread 

``` go 
func UnlockOSThread()
```

UnlockOSThread undoes an earlier call to LockOSThread. If this drops the number of active LockOSThread calls on the calling goroutine to zero, it unwires the calling goroutine from its fixed operating system thread. If there are no active LockOSThread calls, this is a no-op.

Before calling UnlockOSThread, the caller must ensure that the OS thread is suitable for running other goroutines. If the caller made any permanent changes to the state of the thread that would affect other goroutines, it should not call this function and thus leave the goroutine locked to the OS thread until the goroutine (and hence the thread) exits.

#### func Version 

``` go 
func Version() string
```

Version returns the Go tree's version string. It is either the commit hash and date at the time of the build or, when possible, a release tag like "go1.3".

## 类型

### type BlockProfileRecord  <- go1.1

``` go 
type BlockProfileRecord struct {
	Count  int64
	Cycles int64
	StackRecord
}
```

BlockProfileRecord describes blocking events originated at a particular call sequence (stack trace).

### type Error 

``` go 
type Error interface {
	error

	// RuntimeError is a no-op function but
	// serves to distinguish types that are run time
	// errors from ordinary errors: a type is a
	// run time error if it has a RuntimeError method.
	RuntimeError()
}
```

The Error interface identifies a run time error.

### type Frame  <- go1.7

``` go 
type Frame struct {
	// PC is the program counter for the location in this frame.
	// For a frame that calls another frame, this will be the
	// program counter of a call instruction. Because of inlining,
	// multiple frames may have the same PC value, but different
	// symbolic information.
	PC uintptr

	// Func is the Func value of this call frame. This may be nil
	// for non-Go code or fully inlined functions.
	Func *Func

	// Function is the package path-qualified function name of
	// this call frame. If non-empty, this string uniquely
	// identifies a single function in the program.
	// This may be the empty string if not known.
	// If Func is not nil then Function == Func.Name().
	Function string

	// File and Line are the file name and line number of the
	// location in this frame. For non-leaf frames, this will be
	// the location of a call. These may be the empty string and
	// zero, respectively, if not known.
	File string
	Line int

	// Entry point program counter for the function; may be zero
	// if not known. If Func is not nil then Entry ==
	// Func.Entry().
	Entry uintptr
	// contains filtered or unexported fields
}
```

Frame is the information returned by Frames for each call frame.

### type Frames  <- go1.7

``` go 
type Frames struct {
	// contains filtered or unexported fields
}
```

Frames may be used to get function/file/line information for a slice of PC values returned by Callers.

##### Example
``` go 
package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	c := func() {
		// Ask runtime.Callers for up to 10 PCs, including runtime.Callers itself.
		pc := make([]uintptr, 10)
		n := runtime.Callers(0, pc)
		if n == 0 {
			// No PCs available. This can happen if the first argument to
			// runtime.Callers is large.
			//
			// Return now to avoid processing the zero Frame that would
			// otherwise be returned by frames.Next below.
			return
		}

		pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
		frames := runtime.CallersFrames(pc)

		// Loop to get frames.
		// A fixed number of PCs can expand to an indefinite number of Frames.
		for {
			frame, more := frames.Next()

			// Process this frame.
			//
			// To keep this example's output stable
			// even if there are changes in the testing package,
			// stop unwinding when we leave package runtime.
			if !strings.Contains(frame.File, "runtime/") {
				break
			}
			fmt.Printf("- more:%v | %s\n", more, frame.Function)

			// Check whether there are more frames to process after this one.
			if !more {
				break
			}
		}
	}

	b := func() { c() }
	a := func() { b() }

	a()
}
Output:

- more:true | runtime.Callers
- more:true | runtime_test.ExampleFrames.func1
- more:true | runtime_test.ExampleFrames.func2
- more:true | runtime_test.ExampleFrames.func3
- more:true | runtime_test.ExampleFrames
```

#### func CallersFrames  <- go1.7

``` go 
func CallersFrames(callers []uintptr) *Frames
```

CallersFrames takes a slice of PC values returned by Callers and prepares to return function/file/line information. Do not change the slice until you are done with the Frames.

#### (*Frames) Next  <- go1.7

``` go 
func (ci *Frames) Next() (frame Frame, more bool)
```

Next returns a Frame representing the next call frame in the slice of PC values. If it has already returned all call frames, Next returns a zero Frame.

The more result indicates whether the next call to Next will return a valid Frame. It does not necessarily indicate whether this call returned one.

See the Frames example for idiomatic usage.

### type Func 

``` go 
type Func struct {
	// contains filtered or unexported fields
}
```

A Func represents a Go function in the running binary.

#### func FuncForPC 

``` go 
func FuncForPC(pc uintptr) *Func
```

FuncForPC returns a *Func describing the function that contains the given program counter address, or else nil.

If pc represents multiple functions because of inlining, it returns the *Func describing the innermost function, but with an entry of the outermost function.

#### (*Func) Entry 

``` go 
func (f *Func) Entry() uintptr
```

Entry returns the entry address of the function.

#### (*Func) FileLine 

``` go 
func (f *Func) FileLine(pc uintptr) (file string, line int)
```

FileLine returns the file name and line number of the source code corresponding to the program counter pc. The result will not be accurate if pc is not a program counter within f.

#### (*Func) Name 

``` go 
func (f *Func) Name() string
```

Name returns the name of the function.

### type MemProfileRecord 

``` go 
type MemProfileRecord struct {
	AllocBytes, FreeBytes     int64       // number of bytes allocated, freed
	AllocObjects, FreeObjects int64       // number of objects allocated, freed
	Stack0                    [32]uintptr // stack trace for this record; ends at first 0 entry
}
```

A MemProfileRecord describes the live objects allocated by a particular call sequence (stack trace).

#### (*MemProfileRecord) InUseBytes 

``` go 
func (r *MemProfileRecord) InUseBytes() int64
```

InUseBytes returns the number of bytes in use (AllocBytes - FreeBytes).

#### (*MemProfileRecord) InUseObjects 

``` go 
func (r *MemProfileRecord) InUseObjects() int64
```

InUseObjects returns the number of objects in use (AllocObjects - FreeObjects).

#### (*MemProfileRecord) Stack 

``` go 
func (r *MemProfileRecord) Stack() []uintptr
```

Stack returns the stack trace associated with the record, a prefix of r.Stack0.

### type MemStats 

``` go 
type MemStats struct {

	// Alloc is bytes of allocated heap objects.
	//
	// This is the same as HeapAlloc (see below).
	Alloc uint64

	// TotalAlloc is cumulative bytes allocated for heap objects.
	//
	// TotalAlloc increases as heap objects are allocated, but
	// unlike Alloc and HeapAlloc, it does not decrease when
	// objects are freed.
	TotalAlloc uint64

	// Sys is the total bytes of memory obtained from the OS.
	//
	// Sys is the sum of the XSys fields below. Sys measures the
	// virtual address space reserved by the Go runtime for the
	// heap, stacks, and other internal data structures. It's
	// likely that not all of the virtual address space is backed
	// by physical memory at any given moment, though in general
	// it all was at some point.
	Sys uint64

	// Lookups is the number of pointer lookups performed by the
	// runtime.
	//
	// This is primarily useful for debugging runtime internals.
	Lookups uint64

	// Mallocs is the cumulative count of heap objects allocated.
	// The number of live objects is Mallocs - Frees.
	Mallocs uint64

	// Frees is the cumulative count of heap objects freed.
	Frees uint64

	// HeapAlloc is bytes of allocated heap objects.
	//
	// "Allocated" heap objects include all reachable objects, as
	// well as unreachable objects that the garbage collector has
	// not yet freed. Specifically, HeapAlloc increases as heap
	// objects are allocated and decreases as the heap is swept
	// and unreachable objects are freed. Sweeping occurs
	// incrementally between GC cycles, so these two processes
	// occur simultaneously, and as a result HeapAlloc tends to
	// change smoothly (in contrast with the sawtooth that is
	// typical of stop-the-world garbage collectors).
	HeapAlloc uint64

	// HeapSys is bytes of heap memory obtained from the OS.
	//
	// HeapSys measures the amount of virtual address space
	// reserved for the heap. This includes virtual address space
	// that has been reserved but not yet used, which consumes no
	// physical memory, but tends to be small, as well as virtual
	// address space for which the physical memory has been
	// returned to the OS after it became unused (see HeapReleased
	// for a measure of the latter).
	//
	// HeapSys estimates the largest size the heap has had.
	HeapSys uint64

	// HeapIdle is bytes in idle (unused) spans.
	//
	// Idle spans have no objects in them. These spans could be
	// (and may already have been) returned to the OS, or they can
	// be reused for heap allocations, or they can be reused as
	// stack memory.
	//
	// HeapIdle minus HeapReleased estimates the amount of memory
	// that could be returned to the OS, but is being retained by
	// the runtime so it can grow the heap without requesting more
	// memory from the OS. If this difference is significantly
	// larger than the heap size, it indicates there was a recent
	// transient spike in live heap size.
	HeapIdle uint64

	// HeapInuse is bytes in in-use spans.
	//
	// In-use spans have at least one object in them. These spans
	// can only be used for other objects of roughly the same
	// size.
	//
	// HeapInuse minus HeapAlloc estimates the amount of memory
	// that has been dedicated to particular size classes, but is
	// not currently being used. This is an upper bound on
	// fragmentation, but in general this memory can be reused
	// efficiently.
	HeapInuse uint64

	// HeapReleased is bytes of physical memory returned to the OS.
	//
	// This counts heap memory from idle spans that was returned
	// to the OS and has not yet been reacquired for the heap.
	HeapReleased uint64

	// HeapObjects is the number of allocated heap objects.
	//
	// Like HeapAlloc, this increases as objects are allocated and
	// decreases as the heap is swept and unreachable objects are
	// freed.
	HeapObjects uint64

	// StackInuse is bytes in stack spans.
	//
	// In-use stack spans have at least one stack in them. These
	// spans can only be used for other stacks of the same size.
	//
	// There is no StackIdle because unused stack spans are
	// returned to the heap (and hence counted toward HeapIdle).
	StackInuse uint64

	// StackSys is bytes of stack memory obtained from the OS.
	//
	// StackSys is StackInuse, plus any memory obtained directly
	// from the OS for OS thread stacks (which should be minimal).
	StackSys uint64

	// MSpanInuse is bytes of allocated mspan structures.
	MSpanInuse uint64

	// MSpanSys is bytes of memory obtained from the OS for mspan
	// structures.
	MSpanSys uint64

	// MCacheInuse is bytes of allocated mcache structures.
	MCacheInuse uint64

	// MCacheSys is bytes of memory obtained from the OS for
	// mcache structures.
	MCacheSys uint64

	// BuckHashSys is bytes of memory in profiling bucket hash tables.
	BuckHashSys uint64

	// GCSys is bytes of memory in garbage collection metadata.
	GCSys uint64

	// OtherSys is bytes of memory in miscellaneous off-heap
	// runtime allocations.
	OtherSys uint64

	// NextGC is the target heap size of the next GC cycle.
	//
	// The garbage collector's goal is to keep HeapAlloc ≤ NextGC.
	// At the end of each GC cycle, the target for the next cycle
	// is computed based on the amount of reachable data and the
	// value of GOGC.
	NextGC uint64

	// LastGC is the time the last garbage collection finished, as
	// nanoseconds since 1970 (the UNIX epoch).
	LastGC uint64

	// PauseTotalNs is the cumulative nanoseconds in GC
	// stop-the-world pauses since the program started.
	//
	// During a stop-the-world pause, all goroutines are paused
	// and only the garbage collector can run.
	PauseTotalNs uint64

	// PauseNs is a circular buffer of recent GC stop-the-world
	// pause times in nanoseconds.
	//
	// The most recent pause is at PauseNs[(NumGC+255)%256]. In
	// general, PauseNs[N%256] records the time paused in the most
	// recent N%256th GC cycle. There may be multiple pauses per
	// GC cycle; this is the sum of all pauses during a cycle.
	PauseNs [256]uint64

	// PauseEnd is a circular buffer of recent GC pause end times,
	// as nanoseconds since 1970 (the UNIX epoch).
	//
	// This buffer is filled the same way as PauseNs. There may be
	// multiple pauses per GC cycle; this records the end of the
	// last pause in a cycle.
	PauseEnd [256]uint64

	// NumGC is the number of completed GC cycles.
	NumGC uint32

	// NumForcedGC is the number of GC cycles that were forced by
	// the application calling the GC function.
	NumForcedGC uint32

	// GCCPUFraction is the fraction of this program's available
	// CPU time used by the GC since the program started.
	//
	// GCCPUFraction is expressed as a number between 0 and 1,
	// where 0 means GC has consumed none of this program's CPU. A
	// program's available CPU time is defined as the integral of
	// GOMAXPROCS since the program started. That is, if
	// GOMAXPROCS is 2 and a program has been running for 10
	// seconds, its "available CPU" is 20 seconds. GCCPUFraction
	// does not include CPU time used for write barrier activity.
	//
	// This is the same as the fraction of CPU reported by
	// GODEBUG=gctrace=1.
	GCCPUFraction float64

	// EnableGC indicates that GC is enabled. It is always true,
	// even if GOGC=off.
	EnableGC bool

	// DebugGC is currently unused.
	DebugGC bool

	// BySize reports per-size class allocation statistics.
	//
	// BySize[N] gives statistics for allocations of size S where
	// BySize[N-1].Size < S ≤ BySize[N].Size.
	//
	// This does not report allocations larger than BySize[60].Size.
	BySize [61]struct {
		// Size is the maximum byte size of an object in this
		// size class.
		Size uint32

		// Mallocs is the cumulative count of heap objects
		// allocated in this size class. The cumulative bytes
		// of allocation is Size*Mallocs. The number of live
		// objects in this size class is Mallocs - Frees.
		Mallocs uint64

		// Frees is the cumulative count of heap objects freed
		// in this size class.
		Frees uint64
	}
}
```

A MemStats records statistics about the memory allocator.

### type StackRecord 

``` go 
type StackRecord struct {
	Stack0 [32]uintptr // stack trace for this record; ends at first 0 entry
}
```

A StackRecord describes a single execution stack.

#### (*StackRecord) Stack 

``` go 
func (r *StackRecord) Stack() []uintptr
```

Stack returns the stack trace associated with the record, a prefix of r.Stack0.

### type TypeAssertionError 

``` go 
type TypeAssertionError struct {
	// contains filtered or unexported fields
}
```

A TypeAssertionError explains a failed type assertion.

#### (*TypeAssertionError) Error 

``` go 
func (e *TypeAssertionError) Error() string
```

#### (*TypeAssertionError) RuntimeError 

``` go 
func (*TypeAssertionError) RuntimeError()
```