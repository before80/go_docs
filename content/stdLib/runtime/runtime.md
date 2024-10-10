+++
title = "runtime"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/runtime@go1.23.0](https://pkg.go.dev/runtime@go1.23.0)

Package runtime contains operations that interact with Go's runtime system, such as functions to control goroutines. It also includes the low-level type information used by the reflect package; see reflect's documentation for the programmable interface to the run-time type system.

​	`runtime`包包含与Go运行时系统交互的操作，例如控制goroutine的函数。它还包括reflect包使用的低级类型信息；有关可编程接口的运行时类型系统，请参见reflect的文档。

## 环境变量 Environment Variables

The following environment variables (`$name` or `%name%`, depending on the host operating system) control the run-time behavior of Go programs. The meanings and use may change from release to release.

​	以下环境变量(`$name`或`%name%`，具体取决于主机操作系统)控制Go程序的运行时行为。它们的含义和用途可能会随版本发布而发生变化。

The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely. [runtime/debug.SetGCPercent](https://pkg.go.dev/runtime/debug#SetGCPercent) allows changing this percentage at run time.

​	GOGC变量设置初始垃圾回收目标百分比。当新分配数据与上一次回收后剩余的活跃数据之比达到该百分比时，将触发垃圾回收。默认值为GOGC = 100。设置GOGC = off可以完全禁用垃圾收集器。[runtime/debug.SetGCPercent](https://pkg.go.dev/runtime/debug#SetGCPercent)允许在运行时更改此百分比。

The GOMEMLIMIT variable sets a soft memory limit for the runtime. This memory limit includes the Go heap and all other memory managed by the runtime, and excludes external memory sources such as mappings of the binary itself, memory managed in other languages, and memory held by the operating system on behalf of the Go program. GOMEMLIMIT is a numeric value in bytes with an optional unit suffix. The supported suffixes include B, KiB, MiB, GiB, and TiB. These suffixes represent quantities of bytes as defined by the IEC 80000-13 standard. That is, they are based on powers of two: KiB means 2^10 bytes, MiB means 2^20 bytes, and so on. The default setting is math.MaxInt64, which effectively disables the memory limit. [runtime/debug.SetMemoryLimit](https://pkg.go.dev/runtime/debug#SetMemoryLimit) allows changing this limit at run time.

​	GOMEMLIMIT变量为运行时设置软内存限制。此内存限制包括Go堆和运行时管理的所有其他内存，并排除外部内存源，例如二进制映射本身、其他语言中管理的内存以及代表Go程序的操作系统持有的内存。 GOMEMLIMIT是一个以字节为单位的数字值，具有可选的单位后缀。支持的后缀包括B、KiB、MiB、GiB和TiB。这些后缀表示IEC 80000-13标准定义的字节数量。也就是说，它们基于二的幂：KiB表示$2^{10}$字节，MiB表示$2^{20}$字节，依此类推。默认设置为math.MaxInt64，这实际上禁用了内存限制。[runtime/debug.SetMemoryLimit](https://pkg.go.dev/runtime/debug#SetMemoryLimit)允许在运行时更改此限制。

The GODEBUG variable controls debugging variables within the runtime. It is a comma-separated list of name=val pairs setting these named variables:

​	GODEBUG变量控制运行时内部的调试变量。它是一个逗号分隔的name=val 对列表，设置这些命名变量：

```
clobberfree: setting clobberfree=1 causes the garbage collector to
clobber the memory content of an object with bad content when it frees
the object.
clobberfree：将 clobberfree=1 设置为 1 会导致垃圾收集器在释放对象时将对象的内存内容覆盖为无效内容。



cpu.*: cpu.all=off disables the use of all optional instruction set extensions.
cpu.extension=off disables use of instructions from the specified instruction set extension.
extension is the lower case name for the instruction set extension such as sse41 or avx
as listed in internal/cpu package. As an example cpu.avx=off disables runtime detection
and thereby use of AVX instructions.
cpu.*：设置 cpu.all=off 会禁用所有可选的指令集扩展。设置 cpu.extension=off 会禁用来自指定指令集扩展的指令。extension 是指令集扩展的名称，如 sse41 或 avx，详细列表请参考 internal/cpu 包。比如，设置 cpu.avx=off 会禁用运行时检测，从而禁用 AVX 指令的使用。



cgocheck: setting cgocheck=0 disables all checks for packages
using cgo to incorrectly pass Go pointers to non-Go code.
Setting cgocheck=1 (the default) enables relatively cheap
checks that may miss some errors. A more complete, but slow,
cgocheck mode can be enabled using GOEXPERIMENT (which
requires a rebuild), see https://pkg.go.dev/internal/goexperiment for details.
cgocheck：设置 cgocheck=0 会禁用所有针对使用 cgo 的包进行检查，以防止错误地将 Go 指针传递给非 Go 代码。设置 cgocheck=1（默认值）会启用相对简单的检查，但可能会遗漏一些错误。可以使用 GOEXPERIMENT 启用更全面但较慢的 cgocheck 模式（需要重新构建），详情请参阅 https://pkg.go.dev/internal/goexperiment。



disablethp: setting disablethp=1 on Linux disables transparent huge pages for the heap.
It has no effect on other platforms. disablethp is meant for compatibility with versions
of Go before 1.21, which stopped working around a Linux kernel default that can result
in significant memory overuse. See https://go.dev/issue/64332. This setting will be
removed in a future release, so operators should tweak their Linux configuration to suit
their needs before then. See https://go.dev/doc/gc-guide#Linux_transparent_huge_pages.
disablethp：在 Linux 上，将 disablethp=1 会禁用堆上的透明大页功能。此设置对其他平台无效。disablethp 用于与 1.21 之前的 Go 版本兼容，这些版本在处理 Linux 内核的默认设置时可能导致显著的内存过度使用。详情请参阅 https://go.dev/issue/64332。此设置将在未来的版本中被移除，因此用户应在此之前调整其 Linux 配置以满足需求。更多信息请参阅 https://go.dev/doc/gc-guide#Linux_transparent_huge_pages。



dontfreezetheworld: by default, the start of a fatal panic or throw
"freezes the world", preempting all threads to stop all running
goroutines, which makes it possible to traceback all goroutines, and
keeps their state close to the point of panic. Setting
dontfreezetheworld=1 disables this preemption, allowing goroutines to
continue executing during panic processing. Note that goroutines that
naturally enter the scheduler will still stop. This can be useful when
debugging the runtime scheduler, as freezetheworld perturbs scheduler
state and thus may hide problems.
dontfreezetheworld：默认情况下，致命的 panic 或 throw 开始时会“冻结世界”，抢占所有线程以停止所有正在运行的 goroutine，从而可以回溯所有 goroutine，并使其状态接近 panic 点。设置 dontfreezetheworld=1 会禁用这种抢占，允许 goroutine 在 panic 处理过程中继续执行。请注意，自然进入调度器的 goroutine 仍会停止。当调试运行时调度器时，此功能可能很有用，因为冻结世界会干扰调度器状态，从而可能隐藏问题。



efence: setting efence=1 causes the allocator to run in a mode
where each object is allocated on a unique page and addresses are
never recycled.
efence：设置 efence=1 会使分配器在每个对象上分配唯一页面，并且地址永不重复使用。



gccheckmark: setting gccheckmark=1 enables verification of the
garbage collector's concurrent mark phase by performing a
second mark pass while the world is stopped.  If the second
pass finds a reachable object that was not found by concurrent
mark, the garbage collector will panic.
gccheckmark：设置 gccheckmark=1 会启用垃圾收集器并发标记阶段的验证，在世界停止时执行第二次标记。如果第二次标记发现一个可达的对象未在并发标记中找到，垃圾收集器将触发 panic。



gcpacertrace: setting gcpacertrace=1 causes the garbage collector to
print information about the internal state of the concurrent pacer.
gcpacertrace：设置 gcpacertrace=1 会使垃圾收集器输出有关并发调节器内部状态的信息。



gcshrinkstackoff: setting gcshrinkstackoff=1 disables moving goroutines
onto smaller stacks. In this mode, a goroutine's stack can only grow.
gcshrinkstackoff：设置 gcshrinkstackoff=1 会禁用将 goroutine 移动到更小栈上的操作。在此模式下，goroutine 的栈只能增长。



gcstoptheworld: setting gcstoptheworld=1 disables concurrent garbage collection,
making every garbage collection a stop-the-world event. Setting gcstoptheworld=2
also disables concurrent sweeping after the garbage collection finishes.
gcstoptheworld：设置 gcstoptheworld=1 会禁用并发垃圾收集，使每次垃圾收集都成为一个停止世界事件。设置 gcstoptheworld=2 还会禁用垃圾收集完成后的并发清扫。



gctrace: setting gctrace=1 causes the garbage collector to emit a single line to standard
error at each collection, summarizing the amount of memory collected and the
length of the pause. The format of this line is subject to change. Included in
the explanation below is also the relevant runtime/metrics metric for each field.
Currently, it is:
	gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # MB stacks, #MB globals, # P
where the fields are as follows:
	gc #         the GC number, incremented at each GC
	@#s          time in seconds since program start
	#%           percentage of time spent in GC since program start
	#+...+#      wall-clock/CPU times for the phases of the GC
	#->#-># MB   heap size at GC start, at GC end, and live heap, or /gc/scan/heap:bytes
	# MB goal    goal heap size, or /gc/heap/goal:bytes
	# MB stacks  estimated scannable stack size, or /gc/scan/stack:bytes
	# MB globals scannable global size, or /gc/scan/globals:bytes
	# P          number of processors used, or /sched/gomaxprocs:threads
The phases are stop-the-world (STW) sweep termination, concurrent
mark and scan, and STW mark termination. The CPU times
for mark/scan are broken down in to assist time (GC performed in
line with allocation), background GC time, and idle GC time.
If the line ends with "(forced)", this GC was forced by a
runtime.GC() call.
gctrace：设置 gctrace=1 会使垃圾收集器在每次收集时向标准错误输出一行，概述收集到的内存量和暂停时长。该行的格式可能会有所变动。当前格式如下： 
gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # MB stacks, #MB globals, # P 其中字段含义如下： 
gc # 垃圾收集编号，每次垃圾收集递增 
@#s 程序启动后经过的时间（秒） 
#% 程序启动后在垃圾收集上花费的时间百分比 
#+...+# 垃圾收集各阶段的时钟时间/CPU 时间 
#->#-># MB 垃圾收集开始时的堆大小、结束时的堆大小和活动堆大小，或 /gc/scan/heap
# MB goal 目标堆大小，或 /gc/heap/goal
# MB stacks 可扫描栈大小的估计值，或 /gc/scan/stack
# MB globals 可扫描全局变量大小，或 /gc/scan/globals
# P 使用的处理器数量，或 /sched/gomaxprocs
阶段包括停止世界 (STW) 扫描终止、并发标记和扫描，以及 STW 标记终止。标记/扫描的 CPU 时间被分解为辅助时间（随分配执行的垃圾收集）、后台垃圾收集时间和空闲垃圾收集时间。如果行末为“(forced)”，则表示该垃圾收集是由 runtime.GC() 调用强制触发的。

harddecommit: setting harddecommit=1 causes memory that is returned to the OS to
also have protections removed on it. This is the only mode of operation on Windows,
but is helpful in debugging scavenger-related issues on other platforms. Currently,
only supported on Linux.
harddecommit：设置 harddecommit=1 会在将内存返回给操作系统时同时移除其保护。这是 Windows 上的唯一操作模式，但在其他平台上有助于调试与清道夫相关的问题。目前仅支持 Linux。



inittrace: setting inittrace=1 causes the runtime to emit a single line to standard
error for each package with init work, summarizing the execution time and memory
allocation. No information is printed for inits executed as part of plugin loading
and for packages without both user defined and compiler generated init work.
The format of this line is subject to change. Currently, it is:
	init # @#ms, # ms clock, # bytes, # allocs
where the fields are as follows:
	init #      the package name
	@# ms       time in milliseconds when the init started since program start
	# clock     wall-clock time for package initialization work
	# bytes     memory allocated on the heap
	# allocs    number of heap allocations
inittrace：设置 inittrace=1 会使运行时在每个有初始化工作的包执行时向标准错误输出一行，概述执行时间和内存分配情况。插件加载时执行的初始化和没有用户定义和编译器生成的初始化工作的包不会打印信息。该行的格式可能会有所变动。当前格式如下： 
init # @#ms, # ms clock, # bytes, # allocs 
字段含义如下： 
init # 包名 
@# ms 初始化开始时，程序启动后的时间（毫秒） 
# clock 包初始化工作的时钟时间 
# bytes 在堆上分配的内存 
# allocs 堆分配次数


madvdontneed: setting madvdontneed=0 will use MADV_FREE
instead of MADV_DONTNEED on Linux when returning memory to the
kernel. This is more efficient, but means RSS numbers will
drop only when the OS is under memory pressure. On the BSDs and
Illumos/Solaris, setting madvdontneed=1 will use MADV_DONTNEED instead
of MADV_FREE. This is less efficient, but causes RSS numbers to drop
more quickly.
madvdontneed：设置 madvdontneed=0 将在将内存返回内核时使用 MADV_FREE 而不是 MADV_DONTNEED。这更高效，但意味着 RSS 数字只有在操作系统处于内存压力下时才会下降。在 BSD 和 Illumos/Solaris 系统上，设置 madvdontneed=1 将使用 MADV_DONTNEED 而不是 MADV_FREE。虽然效率较低，但会使 RSS 数字更快下降。



memprofilerate: setting memprofilerate=X will update the value of runtime.MemProfileRate.
When set to 0 memory profiling is disabled.  Refer to the description of
MemProfileRate for the default value.
memprofilerate：设置 memprofilerate=X 会更新 runtime.MemProfileRate 的值。设置为 0 时将禁用内存分析。有关默认值的描述，请参阅 MemProfileRate。



profstackdepth: profstackdepth=128 (the default) will set the maximum stack
depth used by all pprof profilers except for the CPU profiler to 128 frames.
Stack traces that exceed this limit will be truncated to the limit starting
from the leaf frame. Setting profstackdepth to any value above 1024 will
silently default to 1024. Future versions of Go may remove this limitation
and extend profstackdepth to apply to the CPU profiler and execution tracer.
profstackdepth：设置 profstackdepth=128（默认值）会将所有 pprof 分析器（CPU 分析器除外）使用的最大栈深度设置为 128 帧。超过此限制的栈跟踪将从叶帧开始截断到限制值。将 profstackdepth 设置为 1024 以上的任何值将默认为 1024。Go 的未来版本可能会移除此限制，并将 profstackdepth 扩展到 CPU 分析器和执行跟踪器。



pagetrace: setting pagetrace=/path/to/file will write out a trace of page events
that can be viewed, analyzed, and visualized using the x/debug/cmd/pagetrace tool.
Build your program with GOEXPERIMENT=pagetrace to enable this functionality. Do not
enable this functionality if your program is a setuid binary as it introduces a security
risk in that scenario. Currently not supported on Windows, plan9 or js/wasm. Setting this
option for some applications can produce large traces, so use with care.
pagetrace：设置 pagetrace=/path/to/file 将写出页面事件跟踪记录，您可以使用 x/debug/cmd/pagetrace 工具查看、分析和可视化这些记录。使用 GOEXPERIMENT=pagetrace 构建程序以启用此功能。如果您的程序是 setuid 二进制文件，请勿启用此功能，因为在这种情况下会带来安全风险。目前不支持 Windows、plan9 或 js/wasm 平台。对某些应用程序设置此选项可能会生成大量跟踪记录，请谨慎使用。



panicnil: setting panicnil=1 disables the runtime error when calling panic with nil
interface value or an untyped nil.
panicnil：设置 panicnil=1 会在调用 panic 且传入的 interface 值为 nil 或无类型 nil 时禁用运行时错误。



runtimecontentionstacks: setting runtimecontentionstacks=1 enables inclusion of call stacks
related to contention on runtime-internal locks in the "mutex" profile, subject to the
MutexProfileFraction setting. When runtimecontentionstacks=0, contention on
runtime-internal locks will report as "runtime._LostContendedRuntimeLock". When
runtimecontentionstacks=1, the call stacks will correspond to the unlock call that released
the lock. But instead of the value corresponding to the amount of contention that call
stack caused, it corresponds to the amount of time the caller of unlock had to wait in its
original call to lock. A future release is expected to align those and remove this setting.
runtimecontentionstacks：设置 runtimecontentionstacks=1 会将与运行时内部锁争用相关的调用栈包含在“mutex”分析中，受 MutexProfileFraction 设置的影响。当 runtimecontentionstacks=0 时，运行时内部锁的争用会报告为“runtime._LostContendedRuntimeLock”。当 runtimecontentionstacks=1 时，调用栈将对应于释放锁的 unlock 调用。但相应值将反映 unlock 调用者在其最初调用 lock 时的等待时间，而不是该调用栈引发的争用量。预计未来的版本会对齐这些设置并移除此选项。



invalidptr: invalidptr=1 (the default) causes the garbage collector and stack
copier to crash the program if an invalid pointer value (for example, 1)
is found in a pointer-typed location. Setting invalidptr=0 disables this check.
This should only be used as a temporary workaround to diagnose buggy code.
The real fix is to not store integers in pointer-typed locations.
invalidptr：invalidptr=1（默认值）会使垃圾收集器和栈复制器在找到无效指针值（例如 1）位于指针类型的位置时崩溃程序。设置 invalidptr=0 将禁用此检查。此选项应仅用作诊断错误代码的临时解决方案。真正的修复方法是不在指针类型的位置存储整数。



sbrk: setting sbrk=1 replaces the memory allocator and garbage collector
with a trivial allocator that obtains memory from the operating system and
never reclaims any memory.
sbrk：设置 sbrk=1 会用一个从操作系统获取内存并且从不回收任何内存的简单分配器替换内存分配器和垃圾收集器。



scavtrace: setting scavtrace=1 causes the runtime to emit a single line to standard
error, roughly once per GC cycle, summarizing the amount of work done by the
scavenger as well as the total amount of memory returned to the operating system
and an estimate of physical memory utilization. The format of this line is subject
to change, but currently it is:
	scav # KiB work (bg), # KiB work (eager), # KiB total, #% util
where the fields are as follows:
	# KiB work (bg)    the amount of memory returned to the OS in the background since
	                   the last line
	# KiB work (eager) the amount of memory returned to the OS eagerly since the last line
	# KiB now          the amount of address space currently returned to the OS
	#% util            the fraction of all unscavenged heap memory which is in-use
If the line ends with "(forced)", then scavenging was forced by a
debug.FreeOSMemory() call.
scavtrace：设置 scavtrace=1 会使运行时大约每个 GC 周期向标准错误输出一行，概述清道夫完成的工作量以及返回给操作系统的内存总量和物理内存利用率的估计值。此行的格式可能会有所变动，但当前格式如下： 
scav # KiB work (bg), # KiB work (eager), # KiB total, #% util 
字段含义如下： 
# KiB work (bg) 自上一行以来在后台返回给操作系统的内存量 
# KiB work (eager) 自上一行以来积极返回给操作系统的内存量 
# KiB now 当前返回给操作系统的地址空间量 
#% util 所有未清理堆内存中正在使用的部分的比例 
如果行末为“(forced)”，则表示清理是由 debug.FreeOSMemory() 调用强制触发的。



scheddetail: setting schedtrace=X and scheddetail=1 causes the scheduler to emit
detailed multiline info every X milliseconds, describing state of the scheduler,
processors, threads and goroutines.
scheddetail：设置 schedtrace=X 和 scheddetail=1 会使调度器每 X 毫秒输出详细的多行信息，描述调度器、处理器、线程和 goroutine 的状态。



schedtrace: setting schedtrace=X causes the scheduler to emit a single line to standard
error every X milliseconds, summarizing the scheduler state.
schedtrace：设置 schedtrace=X 会使调度器每 X 毫秒向标准错误输出一行，概述调度器状态。



tracebackancestors: setting tracebackancestors=N extends tracebacks with the stacks at
which goroutines were created, where N limits the number of ancestor goroutines to
report. This also extends the information returned by runtime.Stack.
Setting N to 0 will report no ancestry information.
tracebackancestors：设置 tracebackancestors=N 会扩展跟踪记录，以包含 goroutine 创建时的栈，其中 N 限制报告的祖先 goroutine 数量。这也会扩展 runtime.Stack 返回的信息。设置 N 为 0 将不报告祖先信息。



tracefpunwindoff: setting tracefpunwindoff=1 forces the execution tracer to
use the runtime's default stack unwinder instead of frame pointer unwinding.
This increases tracer overhead, but could be helpful as a workaround or for
debugging unexpected regressions caused by frame pointer unwinding.
tracefpunwindoff：设置 tracefpunwindoff=1 会强制执行跟踪器使用运行时的默认栈展开器，而不是帧指针展开。这会增加跟踪器的开销，但在帧指针展开引发的意外回归时可能会有帮助，或者作为调试的临时解决方案。



traceadvanceperiod: the approximate period in nanoseconds between trace generations. Only
applies if a program is built with GOEXPERIMENT=exectracer2. Used primarily for testing
and debugging the execution tracer.
traceadvanceperiod：跟踪生成之间的大致周期（纳秒）。仅适用于使用 GOEXPERIMENT=exectracer2 构建的程序。主要用于测试和调试执行跟踪器。



tracecheckstackownership: setting tracecheckstackownership=1 enables a debug check in the
execution tracer to double-check stack ownership before taking a stack trace.
tracecheckstackownership：设置 tracecheckstackownership=1 会在执行跟踪器中启用调试检查，以在获取栈跟踪之前仔细检查栈的所有权。



asyncpreemptoff: asyncpreemptoff=1 disables signal-based
asynchronous goroutine preemption. This makes some loops
non-preemptible for long periods, which may delay GC and
goroutine scheduling. This is useful for debugging GC issues
because it also disables the conservative stack scanning used
for asynchronously preempted goroutines.
asyncpreemptoff：设置 asyncpreemptoff=1 会禁用基于信号的异步 goroutine 抢占。这会使某些循环长时间不可抢占，可能会延迟 GC 和 goroutine 调度。这对于调试 GC 问题非常有用，因为它还禁用了用于异步抢占的 goroutine 的保守栈扫描。

```

The net and net/http packages also refer to debugging variables in GODEBUG. See the documentation for those packages for details.

​	net和net/http包也引用了GODEBUG中的调试变量。有关详细信息，请参阅这些包的文档。

The GOMAXPROCS variable limits the number of operating system threads that can execute user-level Go code simultaneously. There is no limit to the number of threads that can be blocked in system calls on behalf of Go code; those do not count against the GOMAXPROCS limit. This package's GOMAXPROCS function queries and changes the limit.

​	GOMAXPROCS 变量限制了可以同时执行用户级 Go 代码的操作系统线程数量。在代表 Go 代码阻塞的系统调用中，线程数量没有限制；它们不计入 GOMAXPROCS 限制。本包的GOMAXPROCS函数用于查询和更改此限制。

The GORACE variable configures the race detector, for programs built using -race. See https://golang.org/doc/articles/race_detector.html for details.

​	GORACE 变量配置了竞争检测器，用于使用 -race 构建的程序。有关详细信息，请参阅 [https://golang.org/doc/articles/race_detector.html](https://golang.org/doc/articles/race_detector.html)。

The GOTRACEBACK variable controls the amount of output generated when a Go program fails due to an unrecovered panic or an unexpected runtime condition. By default, a failure prints a stack trace for the current goroutine, eliding functions internal to the run-time system, and then exits with exit code 2. The failure prints stack traces for all goroutines if there is no current goroutine or the failure is internal to the run-time. GOTRACEBACK=none omits the goroutine stack traces entirely. GOTRACEBACK=single (the default) behaves as described above. GOTRACEBACK=all adds stack traces for all user-created goroutines. GOTRACEBACK=system is like “all” but adds stack frames for run-time functions and shows goroutines created internally by the run-time. GOTRACEBACK=crash is like “system” but crashes in an operating system-specific manner instead of exiting. For example, on Unix systems, the crash raises SIGABRT to trigger a core dump. GOTRACEBACK=wer is like “crash” but doesn't disable Windows Error Reporting (WER). For historical reasons, the GOTRACEBACK settings 0, 1, and 2 are synonyms for none, all, and system, respectively. The runtime/debug package's SetTraceback function allows increasing the amount of output at run time, but it cannot reduce the amount below that specified by the environment variable. See https://golang.org/pkg/runtime/debug/#SetTraceback.

​	GOTRACEBACK 变量控制在 Go 程序由于未恢复的 panic 或意外的运行时条件而失败时生成的输出量。默认情况下，失败会为当前 goroutine 打印栈跟踪，省略运行时系统内部的函数，然后以退出码 2 退出。如果没有当前 goroutine 或失败是运行时内部的，则失败会打印所有 goroutine 的栈跟踪。GOTRACEBACK=none 完全省略 goroutine 栈跟踪。GOTRACEBACK=single(默认值)的行为如上所述。GOTRACEBACK=all 为所有用户创建的 goroutine 添加栈跟踪。GOTRACEBACK=system 类似于 "all"，但为运行时函数添加栈帧，并显示由运行时内部创建的 goroutine。GOTRACEBACK=crash 类似于 "system"，但以特定于操作系统的方式崩溃而不是退出。例如，在 Unix 系统上，崩溃会引发 SIGABRT 来触发核心转储。出于历史原因，GOTRACEBACK 设置 0、1 和 2 分别是 none、all 和 system 的同义词。runtime/debug 包的 SetTraceback 函数允许在运行时增加输出量，但不能将输出量减少到低于环境变量指定的水平。请参见 [https://golang.org/pkg/runtime/debug/#SetTraceback](https://golang.org/pkg/runtime/debug/#SetTraceback)。

The GOARCH, GOOS, GOPATH, and GOROOT environment variables complete the set of Go environment variables. They influence the building of Go programs (see https://golang.org/cmd/go and https://golang.org/pkg/go/build). GOARCH, GOOS, and GOROOT are recorded at compile time and made available by constants or functions in this package, but they do not influence the execution of the run-time system.

​	GOARCH、GOOS、GOPATH 和 GOROOT 环境变量完成了 Go 环境变量的设置。它们影响构建 Go 程序(请参见 [https://golang.org/cmd/go](https://golang.org/cmd/go) 和 [https://golang.org/pkg/go/build](https://golang.org/pkg/go/build))。GOARCH、GOOS 和 GOROOT 在编译时记录并通过常量或该包中的函数提供，但它们不影响运行时系统的执行。



## 安全性 Security

On Unix platforms, Go's runtime system behaves slightly differently when a binary is setuid/setgid or executed with setuid/setgid-like properties, in order to prevent dangerous behaviors. On Linux this is determined by checking for the AT_SECURE flag in the auxiliary vector, on the BSDs and Solaris/Illumos it is determined by checking the issetugid syscall, and on AIX it is determined by checking if the uid/gid match the effective uid/gid.

​	在 Unix 平台上，当二进制文件具有 setuid/setgid 或以类似 setuid/setgid 的方式执行时，Go 的运行时系统行为会有所不同，以防止危险行为发生。在 Linux 上，这是通过检查辅助向量中的 AT_SECURE 标志来确定的；在 BSD 和 Solaris/Illumos 上，则通过检查 issetugid 系统调用来确定；在 AIX 上，通过检查用户 ID (uid)/组 ID (gid) 是否与有效用户 ID (effective uid)/有效组 ID (effective gid) 匹配来确定。

When the runtime determines the binary is setuid/setgid-like, it does three main things:

​	当运行时系统判断二进制文件具有类似 setuid/setgid 的属性时，会进行以下三个主要操作：

- The standard input/output file descriptors (0, 1, 2) are checked to be open. If any of them are closed, they are opened pointing at /dev/null.
- 检查标准输入/输出文件描述符 (0, 1, 2) 是否处于打开状态。如果有任何一个文件描述符是关闭的，则会将其指向 /dev/null 并打开。

- The value of the GOTRACEBACK environment variable is set to 'none'.
- 设置环境变量 GOTRACEBACK 的值为 'none'。
- When a signal is received that terminates the program, or the program encounters an unrecoverable panic that would otherwise override the value of GOTRACEBACK, the goroutine stack, registers, and other memory related information are omitted.
- 当接收到终止程序的信号，或程序遇到不可恢复的 panic 并且该 panic 通常会覆盖 GOTRACEBACK 的值时，不会显示 goroutine 栈、寄存器和其他内存相关信息。



## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/compiler.go;l=12)

``` go 
const Compiler = "gc"
```

Compiler is the name of the compiler toolchain that built the running binary. Known toolchains are:

​	Compiler 是编译生成运行二进制文件的编译器工具链的名称。已知的工具链包括：

```
gc 		也称 cmd/compile。 Also known as cmd/compile.
gccgo   gccgo 前端，是 GCC 编译器套件的一部分。 The gccgo front end, part of the GCC compiler suite.
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/extern.go;l=303)

``` go 
const GOARCH string = goarch.GOARCH
```

GOARCH is the running program's architecture target: one of 386, amd64, arm, s390x, and so on.

​	GOARCH 是运行程序的体系结构目标，例如 386、amd64、arm、s390x 等。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/extern.go;l=299)

``` go 
const GOOS string = goos.GOOS
```

GOOS is the running program's operating system target: one of darwin, freebsd, linux, and so on. To view possible combinations of GOOS and GOARCH, run "go tool dist list".

​	GOOS 是运行程序的操作系统目标，例如 darwin、freebsd、linux 等。要查看 GOOS 和 GOARCH 的可能组合，请运行"go tool dist list"。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/runtime/mprof.go;l=595)

``` go 
var MemProfileRate int = 512 * 1024
```

MemProfileRate controls the fraction of memory allocations that are recorded and reported in the memory profile. The profiler aims to sample an average of one allocation per MemProfileRate bytes allocated.

​	MemProfileRate 控制记录和报告内存分析中的内存分配的部分。分析器旨在对每个 MemProfileRate 分配的平均样本进行采样。

To include every allocated block in the profile, set MemProfileRate to 1. To turn off profiling entirely, set MemProfileRate to 0.

​	要在分析文件(profile)中包含每个已分配的块，请将 MemProfileRate 设置为 1。要完全关闭分析，请将 MemProfileRate 设置为 0。

The tools that process the memory profiles assume that the profile rate is constant across the lifetime of the program and equal to the current value. Programs that change the memory profiling rate should do so just once, as early as possible in the execution of the program (for example, at the beginning of main).

​	处理内存分析的工具假设分析速率在程序的整个生命周期中是恒定的，并且等于当前值。更改内存分析速率的程序应该只在程序执行的尽早时期(例如在 main 的开始处)执行一次。

## 函数

### func BlockProfile  <- go1.1

``` go 
func BlockProfile(p []BlockProfileRecord) (n int, ok bool)
```

BlockProfile returns n, the number of records in the current blocking profile. If len(p) >= n, BlockProfile copies the profile into p and returns n, true. If len(p) < n, BlockProfile does not change p and returns n, false.

​	BlockProfile函数返回当前阻塞分析中的记录数 n。如果 len(p) >= n，则 BlockProfile 将分析副本复制到 p 并返回 n、true。如果 len(p) < n，则 BlockProfile 不会更改 p 并返回 n、false。

Most clients should use the runtime/pprof package or the testing package's -test.blockprofile flag instead of calling BlockProfile directly.

​	大多数客户端应该使用 runtime/pprof 包或 testing 包的 -test.blockprofile 标志，而不是直接调用 BlockProfile函数。

### func Breakpoint 

``` go 
func Breakpoint()
```

Breakpoint executes a breakpoint trap.

​	Breakpoint函数执行断点陷阱。

### func CPUProfile <-DEPRECATED

```go
func CPUProfile() []byte
```

CPUProfile panics. It formerly provided raw access to chunks of a pprof-format profile generated by the runtime. The details of generating that format have changed, so this functionality has been removed.

​	`CPUProfile` 会触发 `panic`。它之前提供对由运行时生成的 pprof 格式分析的原始访问。由于生成该格式的细节已更改，因此此功能已被移除。

Deprecated: Use the runtime/pprof package, or the handlers in the net/http/pprof package, or the testing package's -test.cpuprofile flag instead.

​	已弃用：请使用 `runtime/pprof` 包，或者使用 `net/http/pprof` 包中的处理程序，或者使用 `testing` 包的 `-test.cpuprofile` 标志。

### func Caller 

``` go 
func Caller(skip int) (pc uintptr, file string, line int, ok bool)
```

Caller reports file and line number information about function invocations on the calling goroutine's stack. The argument skip is the number of stack frames to ascend, with 0 identifying the caller of Caller. (For historical reasons the meaning of skip differs between Caller and Callers.) The return values report the program counter, file name, and line number within the file of the corresponding call. The boolean ok is false if it was not possible to recover the information.

​	Caller函数报告关于调用 goroutine 栈上函数调用的文件和行号信息。skip 是要上升的栈帧数，其中 0 表示 Caller函数的调用者(由于历史原因，skip 在 Caller函数和 Callers函数之间的含义不同)。返回值报告相应调用的程序计数器、文件名和文件中的行号。如果无法恢复信息，则布尔值 ok 为 false。

### func Callers 

``` go 
func Callers(skip int, pc []uintptr) int
```

Callers fills the slice pc with the return program counters of function invocations on the calling goroutine's stack. The argument skip is the number of stack frames to skip before recording in pc, with 0 identifying the frame for Callers itself and 1 identifying the caller of Callers. It returns the number of entries written to pc.

​	Callers函数将调用当前goroutine的栈上函数调用的返回程序计数器填充到切片pc中。参数skip表示在记录pc之前要跳过的栈帧数，其中0标识Callers本身的帧，1标识Callers的调用者。它返回写入到pc的条目数。

To translate these PCs into symbolic information such as function names and line numbers, use CallersFrames. CallersFrames accounts for inlined functions and adjusts the return program counters into call program counters. Iterating over the returned slice of PCs directly is discouraged, as is using FuncForPC on any of the returned PCs, since these cannot account for inlining or return program counter adjustment.

​	要将这些程序计数器转换为符号信息，例如函数名称和行号，请使用CallersFrames函数。 CallersFrames函数考虑了内联函数并将返回程序计数器调整为调用程序计数器。不建议直接迭代返回的PCs切片，也不建议在任何返回的PC上使用FuncForPC函数，因为这些都无法考虑到内联或返回程序计数器的调整。

### func GC 

``` go 
func GC()
```

GC runs a garbage collection and blocks the caller until the garbage collection is complete. It may also block the entire program.

​	GC函数运行垃圾回收并阻塞调用者，直到垃圾回收完成。它也可能阻止整个程序。

### func GOMAXPROCS 

``` go 
func GOMAXPROCS(n int) int
```

GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns the previous setting. It defaults to the value of runtime.NumCPU. If n < 1, it does not change the current setting. This call will go away when the scheduler improves.

​	GOMAXPROCS函数设置可以同时执行的最大CPU数量并返回先前的设置。默认值为runtime.NumCPU的值。如果n < 1，则不更改当前设置。当调度程序改进时，此调用将被取消(This call will go away when the scheduler improves.)。

### func GOROOT 

``` go 
func GOROOT() string
```

GOROOT returns the root of the Go tree. It uses the GOROOT environment variable, if set at process start, or else the root used during the Go build.

​	GOROOT函数返回Go树的根。如果在进程启动时设置了GOROOT环境变量，则使用它，否则使用Go构建期间使用的根目录。

### func Goexit 

``` go 
func Goexit()
```

Goexit terminates the goroutine that calls it. No other goroutine is affected. Goexit runs all deferred calls before terminating the goroutine. Because Goexit is not a panic, any recover calls in those deferred functions will return nil.

​	Goexit函数终止调用它的goroutine。不会影响其他goroutine。Goexit函数在终止goroutine之前运行所有延迟调用。因为Goexit函数不是一个panic，所以这些延迟函数中的任何recover函数调用都将返回nil。

Calling Goexit from the main goroutine terminates that goroutine without func main returning. Since func main has not returned, the program continues execution of other goroutines. If all other goroutines exit, the program crashes.

​	从主goroutine调用Goexit将终止该goroutine，而不是返回func main。由于func main没有返回，程序将继续执行其他goroutine。如果所有其他goroutine退出，则程序崩溃。

### func GoroutineProfile 

``` go 
func GoroutineProfile(p []StackRecord) (n int, ok bool)
```

GoroutineProfile returns n, the number of records in the active goroutine stack profile. If len(p) >= n, GoroutineProfile copies the profile into p and returns n, true. If len(p) < n, GoroutineProfile does not change p and returns n, false.

​	GoroutineProfile函数返回n，活动goroutine栈分析中记录的数量。如果len(p) >= n，则GoroutineProfile函数将分析复制到p中并返回n，true。如果len(p) < n，则GoroutineProfile不更改p并返回n，false。

Most clients should use the runtime/pprof package instead of calling GoroutineProfile directly.

​	大多数客户端应该使用runtime/pprof包而不是直接调用GoroutineProfile函数。

### func Gosched 

``` go 
func Gosched()
```

Gosched yields the processor, allowing other goroutines to run. It does not suspend the current goroutine, so execution resumes automatically.

​	Gosched函数让出处理器，允许其他goroutine运行。它不挂起当前的goroutine，因此执行将自动恢复。

### func KeepAlive  <- go1.7

``` go 
func KeepAlive(x any)
```

KeepAlive marks its argument as currently reachable. This ensures that the object is not freed, and its finalizer is not run, before the point in the program where KeepAlive is called.

​	KeepAlive函数将其参数标记为当前可访问。这确保在调用KeepAlive函数的程序点之前不会释放对象，也不会运行其finalizer(终结器)。

A very simplified example showing where KeepAlive is required:

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

Without the KeepAlive call, the finalizer could run at the start of syscall.Read, closing the file descriptor before syscall.Read makes the actual system call.

​	如果没有 KeepAlive函数的调用，finalizer(终结器)可能会在 syscall.Read 开始时运行，(在 syscall.Read 实际进行系统调用之前)关闭文件描述符。

Note: KeepAlive should only be used to prevent finalizers from running prematurely. In particular, when used with unsafe.Pointer, the rules for valid uses of unsafe.Pointer still apply.

注意：KeepAlive函数应该仅用于防止终结器过早运行。特别地，当与 unsafe.Pointer 一起使用时，仍然适用于 unsafe.Pointer 的有效使用规则。

### func LockOSThread 

``` go 
func LockOSThread()
```

LockOSThread wires the calling goroutine to its current operating system thread. The calling goroutine will always execute in that thread, and no other goroutine will execute in it, until the calling goroutine has made as many calls to UnlockOSThread as to LockOSThread. If the calling goroutine exits without unlocking the thread, the thread will be terminated.

​	LockOSThread函数将调用它的 goroutine 绑定到其当前的操作系统线程。调用 goroutine 将始终在该线程中执行，并且没有其他 goroutine 将在其中执行，直到调用 goroutine 调用 UnlockOSThread 的次数与 LockOSThread 的次数相同。如果调用 goroutine 在不解锁线程的情况下退出，线程将被终止。

All init functions are run on the startup thread. Calling LockOSThread from an init function will cause the main function to be invoked on that thread.

​	所有的 init 函数都在启动线程上运行。从 init 函数中调用 LockOSThread函数将导致在该线程上调用主函数。

A goroutine should call LockOSThread before calling OS services or non-Go library functions that depend on per-thread state.

​	在调用 OS 服务或依赖于每个线程状态的非 Go 库函数之前，goroutine 应该调用 LockOSThread函数。

### func MemProfile 

``` go 
func MemProfile(p []MemProfileRecord, inuseZero bool) (n int, ok bool)
```

MemProfile returns a profile of memory allocated and freed per allocation site.

​	MemProfile函数返回每个分配点分配和释放的内存的分析。

MemProfile returns n, the number of records in the current memory profile. If len(p) >= n, MemProfile copies the profile into p and returns n, true. If len(p) < n, MemProfile does not change p and returns n, false.

​	MemProfile函数返回 n，当前内存分析中的记录数。

​	如果 len(p) >= n，则 MemProfile函数将分析复制到 p 并返回 n，true。

​	如果 len(p) < n，则 MemProfile函数不会改变 p，并返回 n，false。

If inuseZero is true, the profile includes allocation records where r.AllocBytes > 0 but r.AllocBytes == r.FreeBytes. These are sites where memory was allocated, but it has all been released back to the runtime.

​	如果 inuseZero参数 为 true，则分析包括 r.AllocBytes > 0 但 r.AllocBytes == r.FreeBytes 的分配记录。这些是分配了内存但已经全部释放回运行时的站点。

The returned profile may be up to two garbage collection cycles old. This is to avoid skewing the profile toward allocations; because allocations happen in real time but frees are delayed until the garbage collector performs sweeping, the profile only accounts for allocations that have had a chance to be freed by the garbage collector.

​	返回的分析结果可能是最多两个垃圾收集周期之前的。这是为了避免向分析结果倾斜分配的情况；由于分配是实时发生的，但释放需要等到垃圾收集器进行扫描，因此该分析结果仅记录那些已经有机会被垃圾收集器释放的分配情况。

Most clients should use the runtime/pprof package or the testing package's -test.memprofile flag instead of calling MemProfile directly.

​	大多数客户端应该使用 runtime/pprof 包或testing 包的 -test.memprofile 标志，而不是直接调用 MemProfile函数。

### func MutexProfile  <- go1.8

``` go 
func MutexProfile(p []BlockProfileRecord) (n int, ok bool)
```

MutexProfile returns n, the number of records in the current mutex profile. If len(p) >= n, MutexProfile copies the profile into p and returns n, true. Otherwise, MutexProfile does not change p, and returns n, false.

​	MutexProfile函数返回当前mutex profile中记录的数量n。如果len(p) >= n，则MutexProfile函数将profile复制到p中并返回n和true。否则，MutexProfile函数不会更改p，并返回n和false。

Most clients should use the runtime/pprof package instead of calling MutexProfile directly.

​	大多数客户端应该使用runtime/pprof包而不是直接调用MutexProfile函数。

### func NumCPU 

``` go 
func NumCPU() int
```

NumCPU returns the number of logical CPUs usable by the current process.

​	NumCPU函数返回当前进程可用的逻辑 CPU 数量。

The set of available CPUs is checked by querying the operating system at process startup. Changes to operating system CPU allocation after process startup are not reflected.

​	可用的 CPU 集合通过在进程启动时查询操作系统进行检查。进程启动后对操作系统 CPU 分配的更改不会反映出来。

### func NumCgoCall 

``` go 
func NumCgoCall() int64
```

NumCgoCall returns the number of cgo calls made by the current process.

​	NumCgoCall 返回当前进程中发生的 cgo 调用次数。

### func NumGoroutine 

``` go 
func NumGoroutine() int
```

NumGoroutine returns the number of goroutines that currently exist.

​	NumGoroutine 返回当前存在的 goroutine 数量。

### func ReadMemStats 

``` go 
func ReadMemStats(m *MemStats)
```

ReadMemStats populates m with memory allocator statistics.

​	ReadMemStats 用内存分配器的统计信息填充 `m`。

The returned memory allocator statistics are up to date as of the call to ReadMemStats. This is in contrast with a heap profile, which is a snapshot as of the most recently completed garbage collection cycle.

​	返回的内存分配器统计信息是在调用 ReadMemStats 时最新的数据。与栈剖面不同，栈剖面是最近完成的垃圾回收周期的快照。

### func ReadTrace  <- go1.5

``` go 
func ReadTrace() []byte
```

ReadTrace returns the next chunk of binary tracing data, blocking until data is available. If tracing is turned off and all the data accumulated while it was on has been returned, ReadTrace returns nil. The caller must copy the returned data before calling ReadTrace again. ReadTrace must be called from one goroutine at a time.

​	ReadTrace 返回下一个二进制跟踪数据块，阻塞直到数据可用。如果跟踪已关闭且所有在打开时累积的数据都已返回，则 ReadTrace 返回 nil。调用者必须在再次调用 ReadTrace 之前复制返回的数据。ReadTrace 必须一次只从一个 goroutine 中调用。

### func SetBlockProfileRate  <- go1.1

``` go 
func SetBlockProfileRate(rate int)
```

SetBlockProfileRate controls the fraction of goroutine blocking events that are reported in the blocking profile. The profiler aims to sample an average of one blocking event per rate nanoseconds spent blocked.

​	SetBlockProfileRate 控制在阻塞剖面中报告的 goroutine 阻塞事件的比例。探查器旨在每隔 `rate` 纳秒的阻塞时间采样一个阻塞事件。

To include every blocking event in the profile, pass rate = 1. To turn off profiling entirely, pass rate <= 0.

​	要在剖面中包含每个阻塞事件，请传递 `rate = 1`。要完全关闭探查，请传递 `rate <= 0`。

### func SetCPUProfileRate 

``` go 
func SetCPUProfileRate(hz int)
```

SetCPUProfileRate sets the CPU profiling rate to hz samples per second. If hz <= 0, SetCPUProfileRate turns off profiling. If the profiler is on, the rate cannot be changed without first turning it off.

​	SetCPUProfileRate 设置 CPU 探查速率为每秒 `hz` 个样本。如果 `hz <= 0`，则 SetCPUProfileRate 关闭探查。如果探查器已开启，则必须先关闭它，才能更改速率。

Most clients should use the runtime/pprof package or the testing package's -test.cpuprofile flag instead of calling SetCPUProfileRate directly.

​	大多数客户端应使用 `runtime/pprof` 包或测试包的 `-test.cpuprofile` 标志，而不是直接调用 SetCPUProfileRate。

### func SetCgoTraceback  <- go1.7

``` go 
func SetCgoTraceback(version int, traceback, context, symbolizer unsafe.Pointer)
```

SetCgoTraceback records three C functions to use to gather traceback information from C code and to convert that traceback information into symbolic information. These are used when printing stack traces for a program that uses cgo.

​	SetCgoTraceback 记录三个 C 函数，用于从 C 代码中收集回溯信息，并将该回溯信息转换为符号信息。这些函数用于为使用 cgo 的程序打印栈跟踪。

The traceback and context functions may be called from a signal handler, and must therefore use only async-signal safe functions. The symbolizer function may be called while the program is crashing, and so must be cautious about using memory. None of the functions may call back into Go.

​	traceback 和 context 函数可能会从信号处理程序中调用，因此必须仅使用异步信号安全函数。symbolizer 函数可能会在程序崩溃时调用，因此必须小心使用内存。这些函数都不能回调到 Go 代码中。

The context function will be called with a single argument, a pointer to a struct:

​	上下文函数将会被调用并带有一个参数，即指向结构体的指针：

```
struct {
	Context uintptr
}
```

In C syntax, this struct will be

​	在 C 语法中，这个结构体将会是

```
struct {
	uintptr_t Context;
};
```

If the Context field is 0, the context function is being called to record the current traceback context. It should record in the Context field whatever information is needed about the current point of execution to later produce a stack trace, probably the stack pointer and PC. In this case the context function will be called from C code.

​	如果 `Context` 字段为 0，那么该上下文函数被调用的目的是记录当前的追踪上下文。它应该在 `Context` 字段中记录关于当前执行点所需的任何信息，以便稍后生成栈跟踪，可能需要记录栈指针和 PC（程序计数器）。在这种情况下，上下文函数将从 C 代码中调用。

If the Context field is not 0, then it is a value returned by a previous call to the context function. This case is called when the context is no longer needed; that is, when the Go code is returning to its C code caller. This permits the context function to release any associated resources.

​	如果 `Context` 字段不为 0，那么它是先前调用上下文函数时返回的一个值。这个情况是在不再需要上下文时调用的；也就是说，当 Go 代码返回到其 C 代码调用者时调用。这允许上下文函数释放与该上下文相关的资源。

While it would be correct for the context function to record a complete a stack trace whenever it is called, and simply copy that out in the traceback function, in a typical program the context function will be called many times without ever recording a traceback for that context. Recording a complete stack trace in a call to the context function is likely to be inefficient.

​	虽然上下文函数在每次调用时记录完整的栈跟踪是正确的，但在典型程序中，上下文函数会被多次调用，而并不会为该上下文记录栈跟踪。在调用上下文函数时记录完整的栈跟踪可能效率不高。

The traceback function will be called with a single argument, a pointer to a struct:

​	追踪函数将会被调用并带有一个参数，即指向结构体的指针：

```
struct {
	Context    uintptr
	SigContext uintptr
	Buf        *uintptr
	Max        uintptr
}
```

In C syntax, this struct will be

​	在 C 语法中，这个结构体将会是

```
struct {
	uintptr_t  Context;
	uintptr_t  SigContext;
	uintptr_t* Buf;
	uintptr_t  Max;
};
```

The Context field will be zero to gather a traceback from the current program execution point. In this case, the traceback function will be called from C code.

​	当 `Context` 字段为 0 时，将从当前程序执行点收集栈跟踪。在这种情况下，追踪函数将从 C 代码中调用。

Otherwise Context will be a value previously returned by a call to the context function. The traceback function should gather a stack trace from that saved point in the program execution. The traceback function may be called from an execution thread other than the one that recorded the context, but only when the context is known to be valid and unchanging. The traceback function may also be called deeper in the call stack on the same thread that recorded the context. The traceback function may be called multiple times with the same Context value; it will usually be appropriate to cache the result, if possible, the first time this is called for a specific context value.

​	否则，`Context` 将是先前调用上下文函数时返回的一个值。追踪函数应该从程序执行的那个保存点收集栈跟踪。追踪函数可能从与记录上下文的线程不同的执行线程中调用，但只有在上下文已知是有效且不变时才可以这样调用。追踪函数也可能在同一线程中比记录上下文更深的调用栈中被调用。追踪函数可能会多次使用相同的 `Context` 值调用；通常情况下，第一次为特定 `Context` 值调用时缓存结果是合适的。

If the traceback function is called from a signal handler on a Unix system, SigContext will be the signal context argument passed to the signal handler (a C ucontext_t* cast to uintptr_t). This may be used to start tracing at the point where the signal occurred. If the traceback function is not called from a signal handler, SigContext will be zero.

​	如果追踪函数在 Unix 系统上的信号处理程序中被调用，`SigContext` 将是传递给信号处理程序的信号上下文参数（一个转换为 `uintptr_t` 的 C `ucontext_t*`）。这可以用于从信号发生点开始追踪。如果追踪函数不是从信号处理程序中调用的，则 `SigContext` 为 0。

Buf is where the traceback information should be stored. It should be PC values, such that Buf[0] is the PC of the caller, Buf[1] is the PC of that function's caller, and so on. Max is the maximum number of entries to store. The function should store a zero to indicate the top of the stack, or that the caller is on a different stack, presumably a Go stack.

​	`Buf` 是存储栈跟踪信息的地方。它应该存储 PC 值，其中 `Buf[0]` 是调用者的 PC，`Buf[1]` 是该函数调用者的 PC，依此类推。`Max` 是要存储的最大条目数。函数应该存储一个 0，以指示栈的顶部，或者调用者在不同的栈上，可能是在 Go 栈上。

Unlike runtime.Callers, the PC values returned should, when passed to the symbolizer function, return the file/line of the call instruction. No additional subtraction is required or appropriate.

​	与 `runtime.Callers` 不同，返回的 PC 值在传递给符号解析器函数时应该返回调用指令的文件/行号。无需进行额外的减法操作。

On all platforms, the traceback function is invoked when a call from Go to C to Go requests a stack trace. On linux/amd64, linux/ppc64le, linux/arm64, and freebsd/amd64, the traceback function is also invoked when a signal is received by a thread that is executing a cgo call. The traceback function should not make assumptions about when it is called, as future versions of Go may make additional calls.

​	在所有平台上，当从 Go 到 C 再到 Go 的调用请求栈跟踪时，将调用追踪函数。在 `linux/amd64`、`linux/ppc64le`、`linux/arm64` 和 `freebsd/amd64` 上，当执行 cgo 调用的线程收到信号时，追踪函数也会被调用。追踪函数不应对其被调用的时机做出假设，因为未来版本的 Go 可能会增加更多的调用。

The symbolizer function will be called with a single argument, a pointer to a struct:

​	符号解析器函数将被调用并带有一个参数，即指向结构体的指针：

```
struct {
	PC      uintptr // program counter to fetch information for 要获取信息的程序计数器
	File    *byte   // file name (NUL terminated) 文件名（以 NUL 结尾）
	Lineno  uintptr // line number 行号
	Func    *byte   // function name (NUL terminated) 函数名（以 NUL 结尾）
	Entry   uintptr // function entry point 函数入口点
	More    uintptr // set non-zero if more info for this PC 如果该 PC 有更多信息则设为非零
	Data    uintptr // unused by runtime, available for function 运行时未使用，可供函数使用
}
```

In C syntax, this struct will be

​	在 C 语法中，这个结构体将会是

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

​	`PC` 字段将是从追踪函数调用返回的值。

The first time the function is called for a particular traceback, all the fields except PC will be 0. The function should fill in the other fields if possible, setting them to 0/nil if the information is not available. The Data field may be used to store any useful information across calls. The More field should be set to non-zero if there is more information for this PC, zero otherwise. If More is set non-zero, the function will be called again with the same PC, and may return different information (this is intended for use with inlined functions). If More is zero, the function will be called with the next PC value in the traceback. When the traceback is complete, the function will be called once more with PC set to zero; this may be used to free any information. Each call will leave the fields of the struct set to the same values they had upon return, except for the PC field when the More field is zero. The function must not keep a copy of the struct pointer between calls.

​	第一次为特定的追踪调用该函数时，除了 `PC` 之外，所有字段都将为 0。该函数应尽可能填充其他字段，如果信息不可用则将它们设置为 0/nil。`Data` 字段可用于在多次调用之间存储有用信息。如果 `More` 字段设置为非零，则该函数将再次被调用并使用相同的 `PC`，并且可能返回不同的信息（这是为内联函数使用的）。如果 `More` 为零，则该函数将使用追踪中下一个 `PC` 值被调用。当追踪完成时，该函数将再次被调用并将 `PC` 设置为零；这可以用来释放任何信息。每次调用将使结构体的字段保持与返回时相同的值，除了当 `More` 字段为零时的 `PC` 字段。函数不得在调用之间保留结构体指针的副本。

When calling SetCgoTraceback, the version argument is the version number of the structs that the functions expect to receive. Currently this must be zero.

​	调用 `SetCgoTraceback` 时，`version` 参数是函数期望接收的结构体版本号。目前必须为 0。

The symbolizer function may be nil, in which case the results of the traceback function will be displayed as numbers. If the traceback function is nil, the symbolizer function will never be called. The context function may be nil, in which case the traceback function will only be called with the context field set to zero. If the context function is nil, then calls from Go to C to Go will not show a traceback for the C portion of the call stack.

​	符号解析器函数可以为 nil，在这种情况下，追踪函数的结果将以数字形式显示。如果追踪函数为 nil，则符号解析器函数永远不会被调用。上下文函数可以为 nil，在这种情况下，追踪函数将仅在 `Context` 字段设置为 0 时被调用。如果上下文函数为 nil，那么从 Go 到 C 再到 Go 的调用将不会显示 C 部分调用栈的栈跟踪。

SetCgoTraceback should be called only once, ideally from an init function.

​	`SetCgoTraceback` 应仅被调用一次，最好是在初始化函数中调用。

### func SetFinalizer 

``` go 
func SetFinalizer(obj any, finalizer any)
```

SetFinalizer sets the finalizer associated with obj to the provided finalizer function. When the garbage collector finds an unreachable block with an associated finalizer, it clears the association and runs finalizer(obj) in a separate goroutine. This makes obj reachable again, but now without an associated finalizer. Assuming that SetFinalizer is not called again, the next time the garbage collector sees that obj is unreachable, it will free obj.

​	`SetFinalizer` 将与 `obj` 关联的最终器（finalizer）设置为提供的最终器函数。当垃圾收集器找到一个不可达的内存块且其关联了一个最终器时，它会清除该关联并在一个单独的 goroutine 中运行 `finalizer(obj)`。这会使 `obj` 再次可达，但此时不再有关联的最终器。假设没有再次调用 `SetFinalizer`，下次垃圾收集器发现 `obj` 不可达时，它将释放 `obj`。

SetFinalizer(obj, nil) clears any finalizer associated with obj.

​	`SetFinalizer(obj, nil)` 会清除与 `obj` 关联的任何最终器。

The argument obj must be a pointer to an object allocated by calling new, by taking the address of a composite literal, or by taking the address of a local variable. The argument finalizer must be a function that takes a single argument to which obj's type can be assigned, and can have arbitrary ignored return values. If either of these is not true, SetFinalizer may abort the program.

​	参数 `obj` 必须是通过调用 `new` 分配的对象的指针，或是通过获取复合文字的地址，或是获取局部变量的地址。参数 `finalizer` 必须是一个只接受一个参数的函数，该参数可以赋值给 `obj` 的类型，且函数可以有任意被忽略的返回值。如果不满足这些条件，`SetFinalizer` 可能会中止程序。

Finalizers are run in dependency order: if A points at B, both have finalizers, and they are otherwise unreachable, only the finalizer for A runs; once A is freed, the finalizer for B can run. If a cyclic structure includes a block with a finalizer, that cycle is not guaranteed to be garbage collected and the finalizer is not guaranteed to run, because there is no ordering that respects the dependencies.

​	最终器按依赖顺序运行：如果 A 指向 B，且两者都有最终器，且它们在其他方面不可达，则只有 A 的最终器运行；一旦 A 被释放，B 的最终器才能运行。如果循环结构中包含有最终器的块，可能无法保证该循环会被垃圾回收，最终器也可能不会运行，因为无法尊重依赖关系来确定执行顺序。

The finalizer is scheduled to run at some arbitrary time after the program can no longer reach the object to which obj points. There is no guarantee that finalizers will run before a program exits, so typically they are useful only for releasing non-memory resources associated with an object during a long-running program. For example, an os.File object could use a finalizer to close the associated operating system file descriptor when a program discards an os.File without calling Close, but it would be a mistake to depend on a finalizer to flush an in-memory I/O buffer such as a bufio.Writer, because the buffer would not be flushed at program exit.

​	最终器会在程序无法再访问 `obj` 指向的对象后某个任意时间被安排运行。不能保证最终器会在程序退出之前运行，因此通常只有在长时间运行的程序中释放与对象相关的非内存资源时才有用。例如，一个 `os.File` 对象可以使用最终器来关闭与操作系统文件描述符关联的文件，当程序丢弃 `os.File` 而没有调用 `Close` 时，最终器就会自动关闭文件描述符，但依赖最终器来刷新内存中的 I/O 缓冲区（如 `bufio.Writer`）则是错误的做法，因为在程序退出时缓冲区可能不会被刷新。

It is not guaranteed that a finalizer will run if the size of *obj is zero bytes, because it may share same address with other zero-size objects in memory. See https://go.dev/ref/spec#Size_and_alignment_guarantees.

​	如果 `*obj` 的大小为零字节，最终器不一定会运行，因为它可能与内存中其他零大小的对象共享相同的地址。请参阅 https://go.dev/ref/spec#Size_and_alignment_guarantees。

It is not guaranteed that a finalizer will run for objects allocated in initializers for package-level variables. Such objects may be linker-allocated, not heap-allocated.

​	对于包级变量初始化器分配的对象，不保证最终器会运行。这些对象可能是由链接器分配的，而不是堆分配的。

Note that because finalizers may execute arbitrarily far into the future after an object is no longer referenced, the runtime is allowed to perform a space-saving optimization that batches objects together in a single allocation slot. The finalizer for an unreferenced object in such an allocation may never run if it always exists in the same batch as a referenced object. Typically, this batching only happens for tiny (on the order of 16 bytes or less) and pointer-free objects.

​	请注意，由于最终器可能会在对象不再被引用后很长时间才执行，运行时可以进行一种节省空间的优化，将多个对象批量分配在一个内存槽中。如果一个未引用对象与一个仍被引用的对象始终位于同一批次分配中，则该未引用对象的最终器可能永远不会运行。通常，这种批量分配只会发生在微小的（大约16字节或更小）且无指针的对象上。

A finalizer may run as soon as an object becomes unreachable. In order to use finalizers correctly, the program must ensure that the object is reachable until it is no longer required. Objects stored in global variables, or that can be found by tracing pointers from a global variable, are reachable. For other objects, pass the object to a call of the KeepAlive function to mark the last point in the function where the object must be reachable.

​	最终器可能会在对象变为不可达后立即运行。为了正确使用最终器，程序必须确保对象在不再需要之前是可达的。存储在全局变量中的对象，或可以通过从全局变量跟踪指针找到的对象，是可达的。对于其他对象，请将对象传递给 `KeepAlive` 函数，以标记该对象在函数中的最后一个可达点。

For example, if p points to a struct, such as os.File, that contains a file descriptor d, and p has a finalizer that closes that file descriptor, and if the last use of p in a function is a call to syscall.Write(p.d, buf, size), then p may be unreachable as soon as the program enters syscall.Write. The finalizer may run at that moment, closing p.d, causing syscall.Write to fail because it is writing to a closed file descriptor (or, worse, to an entirely different file descriptor opened by a different goroutine). To avoid this problem, call KeepAlive(p) after the call to syscall.Write.

​	例如，如果 `p` 指向一个结构体，如 `os.File`，该结构体包含文件描述符 `d`，并且 `p` 有一个关闭该文件描述符的最终器，如果函数中的 `p` 的最后一次使用是对 `syscall.Write(p.d, buf, size)` 的调用，那么当程序进入 `syscall.Write` 时，`p` 可能会变得不可达。此时最终器可能会运行并关闭 `p.d`，导致 `syscall.Write` 失败，因为它正尝试写入一个已关闭的文件描述符（或更糟糕的是，写入一个不同 goroutine 打开的完全不同的文件描述符）。为了避免这个问题，请在调用 `syscall.Write` 之后调用 `KeepAlive(p)`。

A single goroutine runs all finalizers for a program, sequentially. If a finalizer must run for a long time, it should do so by starting a new goroutine.

​	一个单独的 goroutine 顺序地运行程序中的所有最终器。如果最终器必须长时间运行，它应该通过启动一个新的 goroutine 来执行。

In the terminology of the Go memory model, a call SetFinalizer(x, f) "synchronizes before" the finalization call f(x). However, there is no guarantee that KeepAlive(x) or any other use of x "synchronizes before" f(x), so in general a finalizer should use a mutex or other synchronization mechanism if it needs to access mutable state in x. For example, consider a finalizer that inspects a mutable field in x that is modified from time to time in the main program before x becomes unreachable and the finalizer is invoked. The modifications in the main program and the inspection in the finalizer need to use appropriate synchronization, such as mutexes or atomic updates, to avoid read-write races.

​	在 Go 内存模型的术语中，调用 `SetFinalizer(x, f)` 会“在 f(x) 调用之前同步”。然而，不能保证 `KeepAlive(x)` 或对 `x` 的任何其他使用会“在 f(x) 调用之前同步”，因此一般情况下，如果最终器需要访问 `x` 中的可变状态，应该使用互斥锁或其他同步机制。例如，考虑一个最终器，它检查 `x` 中一个在主程序中时不时修改的可变字段，而在 `x` 变得不可达并且最终器被调用之前。这种情况下，主程序中的修改和最终器中的检查需要使用适当的同步机制，如互斥锁或原子更新，以避免读写竞争。

### func SetMutexProfileFraction  <- go1.8

``` go 
func SetMutexProfileFraction(rate int) int
```

SetMutexProfileFraction controls the fraction of mutex contention events that are reported in the mutex profile. On average 1/rate events are reported. The previous rate is returned.

​	SetMutexProfileFraction 控制在互斥锁配置文件中报告的互斥锁争用事件的比例。平均来说，每 1/rate 个事件被报告。返回之前的 rate 值。

To turn off profiling entirely, pass rate 0. To just read the current rate, pass rate < 0. (For n>1 the details of sampling may change.)

​	要完全关闭配置文件，请传入 rate 为 0。仅读取当前 rate，请传入 rate < 0。（对于 n>1，采样的详细信息可能会有所变化。）

### func Stack 

``` go 
func Stack(buf []byte, all bool) int
```

Stack formats a stack trace of the calling goroutine into buf and returns the number of bytes written to buf. If all is true, Stack formats stack traces of all other goroutines into buf after the trace for the current goroutine.

​	Stack 将调用 goroutine 的栈跟踪格式化到 buf 中，并返回写入 buf 的字节数。如果 all 为 true，Stack 会在当前 goroutine 的跟踪之后，将其他 goroutine 的栈跟踪格式化到 buf 中。

### func StartTrace  <- go1.5

``` go 
func StartTrace() error
```

StartTrace enables tracing for the current process. While tracing, the data will be buffered and available via ReadTrace. StartTrace returns an error if tracing is already enabled. Most clients should use the runtime/trace package or the testing package's -test.trace flag instead of calling StartTrace directly.

​	StartTrace 启用当前进程的跟踪。跟踪时，数据将被缓冲，并可通过 ReadTrace 获取。如果跟踪已启用，StartTrace 返回错误。大多数客户端应使用 runtime/trace 包或 testing 包的 -test.trace 标志，而不是直接调用 StartTrace。

### func StopTrace  <- go1.5

``` go 
func StopTrace()
```

StopTrace stops tracing, if it was previously enabled. StopTrace only returns after all the reads for the trace have completed.

​	StopTrace 停止跟踪（如果先前已启用）。StopTrace 仅在跟踪的所有读取完成后返回。

### func ThreadCreateProfile 

``` go 
func ThreadCreateProfile(p []StackRecord) (n int, ok bool)
```

ThreadCreateProfile returns n, the number of records in the thread creation profile. If len(p) >= n, ThreadCreateProfile copies the profile into p and returns n, true. If len(p) < n, ThreadCreateProfile does not change p and returns n, false.

​	ThreadCreateProfile 返回 n，即线程创建配置文件中的记录数。如果 len(p) >= n，ThreadCreateProfile 将配置文件复制到 p 中并返回 n，true。如果 len(p) < n，ThreadCreateProfile 不会更改 p，并返回 n，false。

Most clients should use the runtime/pprof package instead of calling ThreadCreateProfile directly.

​	大多数客户端应使用 runtime/pprof 包，而不是直接调用 ThreadCreateProfile。

### func UnlockOSThread 

``` go 
func UnlockOSThread()
```

UnlockOSThread undoes an earlier call to LockOSThread. If this drops the number of active LockOSThread calls on the calling goroutine to zero, it unwires the calling goroutine from its fixed operating system thread. If there are no active LockOSThread calls, this is a no-op.

​	UnlockOSThread 撤销先前对 LockOSThread 的调用。如果这使调用 goroutine 上的活动 LockOSThread 调用次数降至零，则会解除调用 goroutine 与其固定操作系统线程的绑定。如果没有活动的 LockOSThread 调用，这是一个无操作（no-op）。

Before calling UnlockOSThread, the caller must ensure that the OS thread is suitable for running other goroutines. If the caller made any permanent changes to the state of the thread that would affect other goroutines, it should not call this function and thus leave the goroutine locked to the OS thread until the goroutine (and hence the thread) exits.

​	在调用 UnlockOSThread 之前，调用者必须确保操作系统线程适合运行其他 goroutine。如果调用者对线程状态进行了任何可能影响其他 goroutine 的永久性更改，则不应调用此函数，从而使 goroutine 锁定在操作系统线程上，直到 goroutine（因此线程）退出。

### func Version 

``` go 
func Version() string
```

Version returns the Go tree's version string. It is either the commit hash and date at the time of the build or, when possible, a release tag like "go1.3".

​	Version 返回 Go 树的版本字符串。它要么是构建时的提交哈希和日期，要么（在可能的情况下）是类似于 "go1.3" 的发布标签。

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

​	BlockProfileRecord 描述了在特定调用序列（栈跟踪）中源自的阻塞事件。

### type Error 

``` go 
type Error interface {
	error

	// RuntimeError is a no-op function but
	// serves to distinguish types that are run time
	// errors from ordinary errors: a type is a
	// run time error if it has a RuntimeError method.
    // RuntimeError 是一个无操作函数，
	// 但它用于区分运行时错误与普通错误：
	// 如果一个类型有 RuntimeError 方法，则它是一个运行时错误。
	RuntimeError()
}
```

The Error interface identifies a run time error.

​	Error 接口用于标识运行时错误。

### type Frame  <- go1.7

``` go 
type Frame struct {
	// PC is the program counter for the location in this frame.
	// For a frame that calls another frame, this will be the
	// program counter of a call instruction. Because of inlining,
	// multiple frames may have the same PC value, but different
	// symbolic information.
    // PC 是此帧中位置的程序计数器。
	// 对于调用另一个帧的帧，这将是调用指令的程序计数器。
	// 由于内联，多个帧可能具有相同的 PC 值，但符号信息不同。
	PC uintptr

	// Func is the Func value of this call frame. This may be nil
	// for non-Go code or fully inlined functions.
    // Func 是此调用帧的 Func 值。对于非 Go 代码或完全内联的函数，这可能为 nil。
	Func *Func

	// Function is the package path-qualified function name of
	// this call frame. If non-empty, this string uniquely
	// identifies a single function in the program.
	// This may be the empty string if not known.
	// If Func is not nil then Function == Func.Name().
    // Function 是此调用帧的包路径限定的函数名称。
	// 如果不为空，此字符串唯一标识程序中的单个函数。
	// 如果未知，此值可能为空字符串。
	// 如果 Func 不为 nil，则 Function == Func.Name()。
	Function string

	// File and Line are the file name and line number of the
	// location in this frame. For non-leaf frames, this will be
	// the location of a call. These may be the empty string and
	// zero, respectively, if not known.
    // File 和 Line 是此帧中位置的文件名和行号。
	// 对于非叶帧，这将是调用的位置。如果未知，它们可能为空字符串和零。
	File string
	Line int

	// Entry point program counter for the function; may be zero
	// if not known. If Func is not nil then Entry ==
	// Func.Entry().
    // 函数的入口点程序计数器；如果未知，可能为零。
	// 如果 Func 不为 nil，则 Entry == Func.Entry()。
	Entry uintptr
	// contains filtered or unexported fields
}
```

Frame is the information returned by Frames for each call frame.

​	Frame 是 Frames 为每个调用帧返回的信息。

### type Frames  <- go1.7

``` go 
type Frames struct {
	// contains filtered or unexported fields
}
```

Frames may be used to get function/file/line information for a slice of PC values returned by Callers.

​	Frames 可用于获取由 Callers 返回的 PC 值片段的函数/文件/行号信息。

#### Example
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
        // 请求 runtime.Callers 获取最多 10 个 PC 值，包括 runtime.Callers 本身。
		pc := make([]uintptr, 10)
		n := runtime.Callers(0, pc)
		if n == 0 {
			// No PCs available. This can happen if the first argument to
			// runtime.Callers is large.
			//
			// Return now to avoid processing the zero Frame that would
			// otherwise be returned by frames.Next below.
            // 没有可用的 PC 值。这可能发生在 runtime.Callers 的第一个参数较大时。
			//
			// 现在返回，以避免处理 frames.Next 下面可能返回的零 Frame。
			return
		}

		pc = pc[:n] // pass only valid pcs to runtime.CallersFrames 仅传递有效的 pc 给 runtime.CallersFrames
		frames := runtime.CallersFrames(pc)

		// Loop to get frames.
		// A fixed number of PCs can expand to an indefinite number of Frames.
        // 循环以获取帧。
		// 固定数量的 PC 可以扩展为无限数量的 Frame。
		for {
			frame, more := frames.Next()

			// Process this frame.
			//
			// To keep this example's output stable
			// even if there are changes in the testing package,
			// stop unwinding when we leave package runtime.
            // 处理此帧。
			//
			// 为了保持此示例的输出稳定，
			// 即使 testing 包发生了变化，
			// 在我们离开 runtime 包时停止展开。
			if !strings.Contains(frame.File, "runtime/") {
				break
			}
			fmt.Printf("- more:%v | %s\n", more, frame.Function)

			// Check whether there are more frames to process after this one.
            // 检查此帧后是否还有更多帧需要处理。
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

​	CallersFrames 接受由 Callers 返回的 PC 值片段，并准备返回函数/文件/行号信息。在完成 Frames 之前不要更改片段。

#### (*Frames) Next  <- go1.7

``` go 
func (ci *Frames) Next() (frame Frame, more bool)
```

Next returns a Frame representing the next call frame in the slice of PC values. If it has already returned all call frames, Next returns a zero Frame.

​	Next 返回表示 PC 值片段中下一个调用帧的 Frame。如果已经返回了所有调用帧，Next 返回零 Frame。

The more result indicates whether the next call to Next will return a valid Frame. It does not necessarily indicate whether this call returned one.

​	more 结果指示下次调用 Next 时是否会返回有效的 Frame。它不一定表示此次调用是否返回了一个。

See the Frames example for idiomatic usage.

​	请参阅 Frames 示例以了解惯用用法。

### type Func 

``` go 
type Func struct {
	// contains filtered or unexported fields
}
```

A Func represents a Go function in the running binary.

​	Func 表示运行二进制文件中的 Go 函数。

#### func FuncForPC 

``` go 
func FuncForPC(pc uintptr) *Func
```

FuncForPC returns a *Func describing the function that contains the given program counter address, or else nil.

​	FuncForPC 返回描述包含给定程序计数器地址的函数的 *Func，否则返回 nil。

If pc represents multiple functions because of inlining, it returns the *Func describing the innermost function, but with an entry of the outermost function.

​	如果 pc 代表多个函数（由于内联），它会返回描述最内层函数的 *Func，但入口点为最外层函数的入口。

#### (*Func) Entry 

``` go 
func (f *Func) Entry() uintptr
```

Entry returns the entry address of the function.

​	Entry 返回函数的入口地址。

#### (*Func) FileLine 

``` go 
func (f *Func) FileLine(pc uintptr) (file string, line int)
```

FileLine returns the file name and line number of the source code corresponding to the program counter pc. The result will not be accurate if pc is not a program counter within f.

​	FileLine 返回与程序计数器 pc 对应的源代码文件名和行号。如果 pc 不是 f 中的程序计数器，则结果将不准确。

#### (*Func) Name 

``` go 
func (f *Func) Name() string
```

Name returns the name of the function.

​	Name 返回函数的名称。

### type MemProfileRecord 

``` go 
type MemProfileRecord struct {
	AllocBytes, FreeBytes     int64       // number of bytes allocated, freed 分配的字节数，释放的字节数
	AllocObjects, FreeObjects int64       // number of objects allocated, freed 分配的对象数，释放的对象数
	Stack0                    [32]uintptr // stack trace for this record; ends at first 0 entry 此记录的栈跟踪；在第一个 0 条目处结束
}
```

A MemProfileRecord describes the live objects allocated by a particular call sequence (stack trace).

​	MemProfileRecord 描述了由特定调用序列（栈跟踪）分配的活动对象。

#### (*MemProfileRecord) InUseBytes 

``` go 
func (r *MemProfileRecord) InUseBytes() int64
```

InUseBytes returns the number of bytes in use (AllocBytes - FreeBytes).

​	InUseBytes 返回使用中的字节数（AllocBytes - FreeBytes）。

#### (*MemProfileRecord) InUseObjects 

``` go 
func (r *MemProfileRecord) InUseObjects() int64
```

InUseObjects returns the number of objects in use (AllocObjects - FreeObjects).

​	InUseObjects 返回使用中的对象数量（AllocObjects - FreeObjects）。

#### (*MemProfileRecord) Stack 

``` go 
func (r *MemProfileRecord) Stack() []uintptr
```

Stack returns the stack trace associated with the record, a prefix of r.Stack0.

​	Stack 返回与该记录相关联的栈跟踪，是 `r.Stack0` 的前缀。

### type MemStats 

``` go 
type MemStats struct {

	// Alloc is bytes of allocated heap objects.
	//
	// This is the same as HeapAlloc (see below).
    // Alloc 是已分配的堆对象的字节数。
    //
    // 这与 HeapAlloc 相同（见下文）。
	Alloc uint64

	// TotalAlloc is cumulative bytes allocated for heap objects.
	//
	// TotalAlloc increases as heap objects are allocated, but
	// unlike Alloc and HeapAlloc, it does not decrease when
	// objects are freed.
    // TotalAlloc 是累计分配的堆对象的字节数。
    //
    // TotalAlloc 随着堆对象的分配而增加，但与 Alloc 和 HeapAlloc 不同，
    // 在对象被释放时它不会减少。
	TotalAlloc uint64

	// Sys is the total bytes of memory obtained from the OS.
	//
	// Sys is the sum of the XSys fields below. Sys measures the
	// virtual address space reserved by the Go runtime for the
	// heap, stacks, and other internal data structures. It's
	// likely that not all of the virtual address space is backed
	// by physical memory at any given moment, though in general
	// it all was at some point.
    // Sys 是从操作系统获取的内存总字节数。
    //
    // Sys 是以下 XSys 字段的总和。Sys 测量 Go 运行时为堆、栈和其他内部数据结构
    // 保留的虚拟地址空间。虽然通常情况下，并非所有的虚拟地址空间在任何时候都
    // 由物理内存支持，但通常它们曾经都支持过。
	Sys uint64

	// Lookups is the number of pointer lookups performed by the
	// runtime.
	//
	// This is primarily useful for debugging runtime internals.
    // Lookups 是运行时执行的指针查找次数。
    //
    // 这主要用于调试运行时内部。
	Lookups uint64

	// Mallocs is the cumulative count of heap objects allocated.
	// The number of live objects is Mallocs - Frees.
    // Mallocs 是分配的堆对象的累计计数。
    // 活跃对象的数量为 Mallocs - Frees。
	Mallocs uint64

	// Frees is the cumulative count of heap objects freed.
    // Frees 是释放的堆对象的累计计数。
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
    // HeapAlloc 是已分配的堆对象的字节数。
    //
    // "已分配"的堆对象包括所有可达的对象，以及垃圾回收器尚未释放的不可达对象。
    // 具体来说，HeapAlloc 随着堆对象的分配而增加，随着堆的清扫和不可达对象的释放而减少。
    // 清扫是增量进行的，因此这些过程同时发生，因此 HeapAlloc 通常平稳变化
    // （与典型的停止世界垃圾回收器的锯齿波形相反）。
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
    // HeapSys 是从操作系统获取的堆内存的字节数。
    //
    // HeapSys 测量为堆保留的虚拟地址空间的大小。这包括已保留但尚未使用的虚拟地址空间，
    // 这些地址空间不会消耗物理内存，但通常很小，还包括物理内存已被返回给操作系统的虚拟地址空间
    // （请参见 HeapReleased 以测量后者）。
    //
    // HeapSys 估计堆的最大大小。
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
    // HeapIdle 是空闲（未使用）的跨度中的字节数。
    //
    // 空闲跨度中没有对象。这些跨度可以（或可能已经）返回给操作系统，
    // 或者可以重新用于堆分配，或者可以重新用于栈内存。
    //
    // HeapIdle 减去 HeapReleased 估计可以返回给操作系统的内存量，
    // 但仍由运行时保留，以便在不向操作系统请求更多内存的情况下扩展堆。
    // 如果此差值显著大于堆大小，表明最近存在堆大小的瞬时峰值。
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
     // HeapInuse 是正在使用的跨度中的字节数。
    //
    // 正在使用的跨度中至少有一个对象。这些跨度只能用于其他大小大致相同的对象。
    //
    // HeapInuse 减去 HeapAlloc 估计了专用于特定大小类别但当前未使用的内存量。
    // 这是碎片的上限，但通常这种内存可以有效地重新使用。
	HeapInuse uint64

	// HeapReleased is bytes of physical memory returned to the OS.
	//
	// This counts heap memory from idle spans that was returned
	// to the OS and has not yet been reacquired for the heap.
    // HeapReleased 是返回给操作系统的物理内存的字节数。
    //
    // 这是从空闲跨度返回给操作系统的堆内存，尚未重新获取用于堆。
	HeapReleased uint64

	// HeapObjects is the number of allocated heap objects.
	//
	// Like HeapAlloc, this increases as objects are allocated and
	// decreases as the heap is swept and unreachable objects are
	// freed.
    // HeapObjects 是已分配的堆对象的数量。
    //
    // 与 HeapAlloc 一样，随着对象的分配而增加，随着堆的清扫和不可达对象的释放而减少。
	HeapObjects uint64

	// StackInuse is bytes in stack spans.
	//
	// In-use stack spans have at least one stack in them. These
	// spans can only be used for other stacks of the same size.
	//
	// There is no StackIdle because unused stack spans are
	// returned to the heap (and hence counted toward HeapIdle).
    // StackInuse 是栈跨度中的字节数。
    //
    // 正在使用的栈跨度中至少有一个栈。这些跨度只能用于相同大小的其他栈。
    //
    // 没有 StackIdle，因为未使用的栈跨度会返回给堆（因此计入 HeapIdle）。
	StackInuse uint64

	// StackSys is bytes of stack memory obtained from the OS.
	//
	// StackSys is StackInuse, plus any memory obtained directly
	// from the OS for OS thread stacks (which should be minimal).
    // StackSys 是从操作系统获取的栈内存的字节数。
    //
    // StackSys 是 StackInuse，加上从操作系统直接获取的用于操作系统线程栈的任何内存（应该很少）。
	StackSys uint64

	// MSpanInuse is bytes of allocated mspan structures.
    // MSpanInuse 是已分配的 mspan 结构的字节数。
	MSpanInuse uint64

	// MSpanSys is bytes of memory obtained from the OS for mspan
	// structures.
    // MSpanSys 是从操作系统获取的用于 mspan 结构的内存字节数。
	MSpanSys uint64

	// MCacheInuse is bytes of allocated mcache structures.
    // MCacheInuse 是已分配的 mcache 结构的字节数。
	MCacheInuse uint64

	// MCacheSys is bytes of memory obtained from the OS for
	// mcache structures.
    // MCacheSys 是从操作系统获取的用于 mcache 结构的内存字节数。
	MCacheSys uint64

	// BuckHashSys is bytes of memory in profiling bucket hash tables.
    // BuckHashSys 是内存分析 bucket 哈希表中的字节数。
	BuckHashSys uint64

	// GCSys is bytes of memory in garbage collection metadata.
    // GCSys 是垃圾回收元数据中的字节数。
	GCSys uint64

	// OtherSys is bytes of memory in miscellaneous off-heap
	// runtime allocations.
    // OtherSys 是其他运行时分配的堆外内存的字节数。
	OtherSys uint64

	// NextGC is the target heap size of the next GC cycle.
	//
	// The garbage collector's goal is to keep HeapAlloc ≤ NextGC.
	// At the end of each GC cycle, the target for the next cycle
	// is computed based on the amount of reachable data and the
	// value of GOGC.
    // NextGC 是下一个 GC 周期的目标堆大小。
    //
    // 垃圾回收器的目标是保持 HeapAlloc ≤ NextGC。
    // 在每个 GC 周期结束时，下一个周期的目标是根据可达数据量和 GOGC 的值计算的。
	NextGC uint64

	// LastGC is the time the last garbage collection finished, as
	// nanoseconds since 1970 (the UNIX epoch).
    // LastGC 是上一次垃圾回收结束的时间，自 1970 年（UNIX 纪元）以来的纳秒数。
	LastGC uint64

	// PauseTotalNs is the cumulative nanoseconds in GC
	// stop-the-world pauses since the program started.
	//
	// During a stop-the-world pause, all goroutines are paused
	// and only the garbage collector can run.
    // PauseTotalNs 是自程序启动以来 GC 停止世界暂停的累计纳秒数。
    //
    // 在停止世界暂停期间，所有 goroutine 都会暂停，只有垃圾回收器可以运行。
	PauseTotalNs uint64

	// PauseNs is a circular buffer of recent GC stop-the-world
	// pause times in nanoseconds.
	//
	// The most recent pause is at PauseNs[(NumGC+255)%256]. In
	// general, PauseNs[N%256] records the time paused in the most
	// recent N%256th GC cycle. There may be multiple pauses per
	// GC cycle; this is the sum of all pauses during a cycle.
    // PauseNs 是最近 GC 停止世界暂停时间的循环缓冲区，单位为纳秒。
    //
    // 最近的暂停时间在 PauseNs[(NumGC+255)%256] 中。
    // 通常，PauseNs[N%256] 记录最近 N%256 次 GC 周期的暂停时间。
    // 每个 GC 周期可能有多个暂停；这是周期中所有暂停的总和。
	PauseNs [256]uint64

	// PauseEnd is a circular buffer of recent GC pause end times,
	// as nanoseconds since 1970 (the UNIX epoch).
	//
	// This buffer is filled the same way as PauseNs. There may be
	// multiple pauses per GC cycle; this records the end of the
	// last pause in a cycle.
    // PauseEnd 是最近 GC 暂停结束时间的循环缓冲区，自 1970 年（UNIX 纪元）以来的纳秒数。
    //
    // 这个缓冲区的填充方式与 PauseNs 相同。每个 GC 周期可能有多个暂停；
    // 这记录了周期中最后一次暂停的结束时间。
	PauseEnd [256]uint64

	// NumGC is the number of completed GC cycles.
    // NumGC 是已完成的 GC 周期数。
	NumGC uint32

	// NumForcedGC is the number of GC cycles that were forced by
	// the application calling the GC function.
    // NumForcedGC 是应用程序调用 GC 函数强制的 GC 周期数。
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
    // GCCPUFraction 是自程序启动以来由 GC 使用的 CPU 时间的比例。
    //
    // GCCPUFraction 表示为 0 到 1 之间的数字，其中 0 表示 GC 未消耗此程序的 CPU。
    // 程序的可用 CPU 时间定义为 GOMAXPROCS 自程序启动以来的积分。
    // 也就是说，如果 GOMAXPROCS 是 2，程序运行了 10 秒，那么它的“可用 CPU”是 20 秒。
    // GCCPUFraction 不包括用于写屏障活动的 CPU 时间。
    //
    // 这与 GODEBUG=gctrace=1 报告的 CPU 百分比相同。
	GCCPUFraction float64

	// EnableGC indicates that GC is enabled. It is always true,
	// even if GOGC=off.
    // EnableGC 表示启用了 GC。即使 GOGC=off，它也始终为 true。
	EnableGC bool

	// DebugGC is currently unused.
    // DebugGC 当前未使用。
	DebugGC bool

	// BySize reports per-size class allocation statistics.
	//
	// BySize[N] gives statistics for allocations of size S where
	// BySize[N-1].Size < S ≤ BySize[N].Size.
	//
	// This does not report allocations larger than BySize[60].Size.
    // BySize 报告按大小分类的分配统计信息。
    //
    // BySize[N] 提供了 S 大小的分配统计信息，其中 BySize[N-1].Size < S ≤ BySize[N].Size。
    //
    // 这不会报告大于 BySize[60].Size 的分配。
	BySize [61]struct {
		// Size is the maximum byte size of an object in this
		// size class.
        // Size 是此大小类别中对象的最大字节大小。
		Size uint32

		// Mallocs is the cumulative count of heap objects
		// allocated in this size class. The cumulative bytes
		// of allocation is Size*Mallocs. The number of live
		// objects in this size class is Mallocs - Frees.
        // Mallocs 是此大小类别中已分配堆对象的累计计数。
        // 分配的累计字节数为 Size*Mallocs。此大小类别中的活跃对象数为 Mallocs - Frees。
		Mallocs uint64

		// Frees is the cumulative count of heap objects freed
		// in this size class.
        // Frees 是此大小类别中已释放的堆对象的累计计数。
		Frees uint64
	}
}
```

A MemStats records statistics about the memory allocator.

​	MemStats 记录了内存分配器的统计信息。

#### type PanicNilError <-go1.21.0

```go
type PanicNilError struct {
	// contains filtered or unexported fields
}
```

A PanicNilError happens when code calls panic(nil).

​	PanicNilError 发生在代码调用 `panic(nil)` 时。

Before Go 1.21, programs that called panic(nil) observed recover returning nil. Starting in Go 1.21, programs that call panic(nil) observe recover returning a *PanicNilError. Programs can change back to the old behavior by setting GODEBUG=panicnil=1.

​	在 Go 1.21 之前，调用 `panic(nil)` 的程序会观察到 `recover` 返回 `nil`。从 Go 1.21 开始，调用 `panic(nil)` 的程序会观察到 `recover` 返回一个 `*PanicNilError`。程序可以通过设置 `GODEBUG=panicnil=1` 来恢复旧的行为。

#### (*PanicNilError) Error <-go1.21.0

```go
func (*PanicNilError) Error() string
```

#### (*PanicNilError) RuntimeError <-go1.21.0

```go
func (*PanicNilError) RuntimeError()
```

#### type Pinner <-go1.21.0

```go
type Pinner struct {
	// contains filtered or unexported fields
}
```

A Pinner is a set of pinned Go objects. An object can be pinned with the Pin method and all pinned objects of a Pinner can be unpinned with the Unpin method.

​	Pinner 是一组固定的 Go 对象。可以使用 `Pin` 方法固定一个对象，并且可以使用 `Unpin` 方法解锁所有固定的对象。

#### (*Pinner) Pin <-go1.21.0

```go
func (p *Pinner) Pin(pointer any)
```

Pin pins a Go object, preventing it from being moved or freed by the garbage collector until the Unpin method has been called.

​	`Pin` 固定一个 Go 对象，防止垃圾回收器在调用 `Unpin` 方法之前移动或释放该对象。

A pointer to a pinned object can be directly stored in C memory or can be contained in Go memory passed to C functions. If the pinned object itself contains pointers to Go objects, these objects must be pinned separately if they are going to be accessed from C code.

​	固定对象的指针可以直接存储在 C 内存中，也可以包含在传递给 C 函数的 Go 内存中。如果固定的对象本身包含指向其他 Go 对象的指针，这些对象在从 C 代码中访问时必须分别固定。

The argument must be a pointer of any type or an unsafe.Pointer. It must be the result of calling new, taking the address of a composite literal, or taking the address of a local variable. If one of these conditions is not met, Pin will panic.

​	参数必须是任意类型的指针或 `unsafe.Pointer`。它必须是调用 `new` 的结果、取复合字面量的地址或取局部变量的地址。如果不满足这些条件之一，`Pin` 将会触发 `panic`。

#### (*Pinner) Unpin <-go1.21.0

```go
func (p *Pinner) Unpin()
```

Unpin unpins all pinned objects of the Pinner.

​	`Unpin` 解锁 Pinner 中所有固定的对象。

### type StackRecord 

``` go 
type StackRecord struct {
	Stack0 [32]uintptr // stack trace for this record; ends at first 0 entry 此记录的栈跟踪；在第一个 0 条目处结束
}
```

A StackRecord describes a single execution stack.

​	StackRecord 描述单个执行栈。

#### (*StackRecord) Stack 

``` go 
func (r *StackRecord) Stack() []uintptr
```

Stack returns the stack trace associated with the record, a prefix of r.Stack0.

​	`Stack` 返回与该记录相关联的栈跟踪，是 `r.Stack0` 的前缀。

### type TypeAssertionError 

``` go 
type TypeAssertionError struct {
	// contains filtered or unexported fields
}
```

A TypeAssertionError explains a failed type assertion.

​	TypeAssertionError 解释了类型断言失败的原因。

#### (*TypeAssertionError) Error 

``` go 
func (e *TypeAssertionError) Error() string
```

#### (*TypeAssertionError) RuntimeError 

``` go 
func (*TypeAssertionError) RuntimeError()
```