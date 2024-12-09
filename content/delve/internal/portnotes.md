+++
title = "portnotes"
date = 2024-12-09T08:03:53+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md)
>
> 收录该文档时间： `2024-12-09T08:03:53+08:00`



# Notes on porting Delve to other architectures - 移植Delve到其他架构的笔记



## 持续集成要求 Continuous Integration requirements



Code that isn't tested doesn't work, we like to run CI on all supported platforms. Currently our CI is done on an [instance of TeamCity cloud provided by JetBrains](https://delve.teamcity.com/), with the exception of the FreeBSD port, which is tested by Cirrus-CI.

​	未经测试的代码无法正常工作，我们希望在所有支持的平台上运行持续集成（CI）。目前，我们的CI是在由JetBrains提供的[TeamCity云实例](https://delve.teamcity.com/)上完成的，FreeBSD移植版本由Cirrus-CI测试。

TeamCity settings are in `.teamcity/settings.kts` which in turn runs one of `_scripts/test_linux.sh`, `_scripts/test_mac.sh` or `_scripts/test_windows.ps1`. All test scripts eventually end up calling into our main test program `_scripts/make.go`, which makes the appropriate `go test` calls.

​	TeamCity的设置位于`.teamcity/settings.kts`，该文件会调用其中之一的`_scripts/test_linux.sh`、`_scripts/test_mac.sh`或`_scripts/test_windows.ps1`。所有的测试脚本最终都会调用我们主要的测试程序`_scripts/make.go`，该程序会执行相应的`go test`调用。

If you plan to port Delve to a new platform you should first figure out how we are going to add your port to our CI, there are three possible solutions:

​	如果你计划将Delve移植到新平台，你应首先弄清楚如何将你的移植版本添加到我们的CI中，目前有三种可能的解决方案：

1. the platform can be run on existing agents we have on TeamCity (linux/amd64, linux/arm64, windows/amd64, darwin/amd64, darwin/arm64) through Docker or similar solutions. 该平台可以通过Docker或类似解决方案在我们现有的TeamCity代理上运行（linux/amd64, linux/arm64, windows/amd64, darwin/amd64, darwin/arm64）。
2. there is a free CI service that integrates with GitHub that we can use 有一个可以与GitHub集成的免费的CI服务供我们使用。
3. you provide the hardware to be added to TeamCity to test it 你提供硬件并将其添加到TeamCity中进行测试。

Exception to this requirement can be discussed in special cases.

​	对于这个要求的例外情况，可以根据特殊情况讨论。

## 一般代码组织 General code organization



An introduction to the architecture of Delve can be found in the 2018 Gophercon Iceland talk: [slides](https://speakerdeck.com/aarzilli/internal-architecture-of-delve), [video](https://www.youtube.com/watch?v=IKnTr7Zms1k).

​	Delve的架构介绍可以在2018年Gophercon Iceland演讲中找到：[幻灯片](https://speakerdeck.com/aarzilli/internal-architecture-of-delve)、[视频](https://www.youtube.com/watch?v=IKnTr7Zms1k)。

### 你不需要担心的包 Packages you shouldn't worry about



- `cmd/dlv/...` implements the command line program
  - `cmd/dlv/...` 实现了命令行程序
- `pkg/terminal/...` implements the command line user interface
  - `pkg/terminal/...` 实现了命令行用户界面
- `service/...` with the exception of `service/test`, implements our API as well as DAP
  - `service/...`（`service/test`除外）实现了我们的API以及DAP
- `pkg/dwarf/...` with the exception of `pkg/dwarf/regnum`, implements DWARF features not covered by the standard library
  - `pkg/dwarf/...`（`pkg/dwarf/regnum`除外）实现了DWARF特性，标准库未涵盖的部分
- anything else in `pkg` that isn't inside `pkg/proc`
  - `pkg`中其他所有不在`pkg/proc`内的内容

### pkg/proc



`pkg/proc` is the symbolic layer of Delve, its job is to bridge the distance between the API and low level Operating System and CPU features. Almost all features of Delve are implemented here, **except for the interaction with the OS/CPU**, which is provided by one of three backends: `pkg/proc/native`, `pkg/proc/core` or `pkg/proc/gdbserial`.

​	`pkg/proc`是Delve的符号层，它的作用是桥接API和底层操作系统及CPU特性之间的距离。几乎所有Delve的功能都在这里实现，**除了与操作系统/CPU的交互**，这些交互由以下三个后端之一提供：`pkg/proc/native`、`pkg/proc/core`或`pkg/proc/gdbserial`。

This package also contains the main test suite for Delve's backends in `pkg/proc/proc_test.go`. The tests for reading variables, however, are inside `service/test` for historical reasons.

​	此包还包含Delve后端的主要测试套件，位于`pkg/proc/proc_test.go`。然而，读取变量的测试位于`service/test`，这出于历史原因。

When porting Delve to a new CPU a new instance of the `proc.Arch` structure should be filled, see `pkg/proc/arch.go` and `pkg/proc/amd64_arch.go` as an example. To do this you will have to:

​	当将Delve移植到新CPU时，应填写`proc.Arch`结构的一个新实例，参考`pkg/proc/arch.go`和`pkg/proc/amd64_arch.go`作为示例。为了做到这一点，你需要：

- provide a disassembler for the port architecture
  - 为目标架构提供反汇编器

- provide a mapping between DWARF register numbers and hardware registers in `pkg/dwarf/regnum` (see `pkg/dwarf/regnum/amd64.go` as an example). This mapping *is not arbitrary* it needs to be described in some standard document which should be linked to in the documentation of `pkg/dwarf/regnum`, for example the mapping for amd64 is described by the System V ABI AMD64 Architecture Processor Supplement v. 1.0 on page 61 figure 3.36.
  - 在`pkg/dwarf/regnum`中提供DWARF寄存器号与硬件寄存器之间的映射（参考`pkg/dwarf/regnum/amd64.go`作为示例）。此映射*不是任意的*，它需要在某个标准文档中进行描述，并应在`pkg/dwarf/regnum`的文档中提供链接。例如，amd64的映射描述在《System V ABI AMD64 Architecture Processor Supplement v. 1.0》第61页的图3.36中。
- if you don't know what `proc.Arch.fixFrameUnwindContext` or `porc.Arch.switchStack` should do *leave them empty*
  - 如果你不知道`proc.Arch.fixFrameUnwindContext`或`proc.Arch.switchStack`应该做什么，*请保持空白*。
- the `proc.Arch.prologues` field needs to be filled by looking at the relevant parts of the Go linker (`cmd/link`), the code is somewhere inside `$GOROOT/src/cmd/internal/obj`, usually the function is called `stacksplit`.
  - `proc.Arch.prologues`字段需要通过查看Go链接器（`cmd/link`）的相关部分来填写，代码位于`$GOROOT/src/cmd/internal/obj`中，通常该函数被称为`stacksplit`。

If your target OS uses an executable file format other than ELF, Mach-O or PE you will also have to change `pkg/proc/bininfo.go` .

​	如果你的目标操作系统使用ELF、Mach-O或PE之外的可执行文件格式，你还需要修改`pkg/proc/bininfo.go`。

**See also** the note on [build tags](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#buildtags).

​	**另见**：[构建标签](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#buildtags)相关说明。

### pkg/proc/gdbserial



This implements GDB remote serial protocol, it is used as the main backend on macOS as well as connecting to [rr](https://rr-project.org/).

​	该部分实现了GDB远程串行协议，作为macOS的主要后端，并与[rr](https://rr-project.org/)连接。

Unless you are making a macOS port you shouldn't worry about this.

​	除非你正在进行macOS的移植，否则不需要担心此部分。

### pkg/proc/core



This implements code for reading core files. You don't need to support this for the port platform, see the note on [skippable features](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#skippable). If you decide to do it anyway see the note on [build tags](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#buildtags).

​	该部分实现了读取core文件的功能。对于移植平台，你无需支持此功能，参见[可跳过的特性](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#skippable)。如果你决定实现此功能，参见[构建标签](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#buildtags)相关说明。

### pkg/proc/native



This is the interface between Delve and the OS/CPU, you will definitely want to work on this and it will be the bulk of the port job. The tests for this code however are not in this directory, they are in `pkg/proc` and `service/test`.

​	这是Delve与操作系统/CPU之间的接口，你肯定需要处理这一部分，它将是移植工作的主要部分。然而，针对这部分代码的测试并不在此目录中，而是在`pkg/proc`和`service/test`中。

## 一般移植过程 General port process



1. Edit `pkg/proc/native/support_sentinel.go` to disable it in the port platform 编辑`pkg/proc/native/support_sentinel.go`以在移植平台上禁用它
2. Fill `proc.Arch` struct for target architecture if it isn't supported already 填写目标架构的`proc.Arch`结构（如果该架构尚未被支持）
3. Go in `pkg/proc`
   - run `go test -v`
   - fix compiler errors 修复编译错误
   - repeat until compilation succeeds 重复直到编译成功
   - fix test failures 修复测试失败
   - repeat until almost all tests pass (see note on [skippable tests and features](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#skippable)) 重复直到几乎所有测试通过（参见[可跳过的测试和特性](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#skippable)）
4. Go in `service/test`
   - run `go test -v`
   - fix compiler errors 修复编译错误
   - repeat until compilation succeeds 重复直到编译成功
   - fix test failures 修复测试失败
   - repeat until almost all tests pass (see note on [skippable tests and features](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#skippable)) 重复直到几乎所有测试通过（参见[可跳过的测试和特性](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#skippable)）
5. Go to the root directory of the project 进入项目根目录
   - run `go run _scripts/make.go test`
   - fix compiler errors
   - repeat until compilation succeeds
   - fix test failures
   - repeat until almost all tests pass (see note on [skippable tests and features](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#skippable)) 重复直到几乎所有测试通过（参见[可跳过的测试和特性](https://github.com/go-delve/delve/blob/master/Documentation/internal/portnotes.md#skippable)）

## 杂项 Miscellaneous



### 构建标签、runtime.GOOS和runtime.GOARCH的使用 Uses of build tags, runtime.GOOS and runtime.GOARCH



Delve has the ability to read cross-platform core files: you can read a core file of any supported platform with Delve running on any other supported platform. For example, a core file produced by linux/arm64 can be read using Delve running under windows/amd64. This feature has far reaching consequences, for example the stack unwinding code in `pkg/proc/stack.go` could be asked to unwind a stack for an architecture different from the one its running under.

​	Delve具有读取跨平台core文件的能力：你可以在任何支持的平台上运行Delve并读取来自其他支持平台的core文件。例如，运行在linux/arm64上的core文件可以通过运行在windows/amd64上的Delve读取。这个特性有着深远的影响，例如，`pkg/proc/stack.go`中的栈展开代码可能需要展开一个与其运行平台不同架构的栈。

What this means in practice is that, in general, using build tags (like `_amd64.go`) or checking `runtime.GOOS` and `runtime.GOARCH` is forbidden throughout Delve's source tree, with two important exceptions:

​	实际应用中，这意味着通常情况下，整个Delve源码树中不允许使用构建标签（如`_amd64.go`）或检查`runtime.GOOS`和`runtime.GOARCH`，有两个重要例外：

- `pkg/proc/native` is allowed to check runtime.GOOS/runtime.GOARCH as well as using build tags
  - `pkg/proc/native`允许检查`runtime.GOOS/runtime.GOARCH`并使用构建标签
- test files are allowed to check runtime.GOOS/runtime.GOARCH as well as using build tags
  - 测试文件允许检查`runtime.GOOS/runtime.GOARCH`并使用构建标签

Other exceptions can be considered, but in general code outside of `pkg/proc/native` should:

​	其他例外情况可以考虑，但一般来说，`pkg/proc/native`之外的代码应：

- use `proc.BinaryInfo.GOOS` instead of `runtime.GOOS`
  - 使用`proc.BinaryInfo.GOOS`代替`runtime.GOOS`
- use `proc.BinaryInfo.Arch.Name` instead of `runtime.GOARCH`
  - 使用`proc.BinaryInfo.Arch.Name`代替`runtime.GOARCH`
- use `proc.BinaryInfo.Arch.PtrSize()` instead of determining the pointer size with `unsafe.Sizeof`
  - 使用`proc.BinaryInfo.Arch.PtrSize()`代替使用`unsafe.Sizeof`确定指针大小
- use `uint64` wherever an address-sized integer is needed, instead of `uintptr`
  - 在需要地址大小整数时使用`uint64`，而不是`uintptr`
- use `amd64_filename.go` instead of the build tag version, `filename_amd64.go`
  - 使用`amd64_filename.go`代替构建标签版本`filename_amd64.go`

### 可以被移植跳过的特性和测试 Features and tests that can be skipped by a port



Delve offers many features, however not all of them are necessary for a useful port of Delve. The following features are optional to implement for a port:

​	Delve提供了许多功能，但并非所有功能对于一个有用的移植版都是必要的。以下功能是可选的，移植时可以选择实现：

- Reading core files (i.e. `pkg/proc/core`)
  - 读取core文件（即`pkg/proc/core`）
- Writing core files (i.e. `pkg/proc/native/dump_*.go`)
  - 写入core文件（即`pkg/proc/native/dump_*.go`）
- Watchpoints (`(*nativeThread).writeHardwareBreakpoint` etc)
  - 监视点（`(*nativeThread).writeHardwareBreakpoint`等）
- Supporting CGO calls (`proc.Arch.switchStack`)
  - 支持CGO调用（`proc.Arch.switchStack`）
- eBPF (`pkg/proc/internal/ebpf`)
  - eBPF（`pkg/proc/internal/ebpf`）
- Working with Position Independent Executables (PIE), unless the default buildmode for the port platform is PIE
  - 处理位置无关可执行文件（PIE），除非移植平台的默认构建模式为PIE
- Function call injection (`pkg/proc/fncall.go` -- it is probably not supported on the port architecture anyway)
  - 函数调用注入（`pkg/proc/fncall.go` — 可能在移植架构上不被支持）

For all these features it is acceptable (and possibly advisable) to either leave the implementation empty or to write a stub that always returns an "not implemented" error. Tests relative to these features can be skipped, `proc_test.go` has a `skipOn` utility function that can be called to skip a specific test on some architectures.

​	对于这些特性，接受（并可能建议）将实现留空或编写一个总是返回“未实现”错误的存根。与这些特性相关的测试可以跳过，`proc_test.go`有一个`skipOn`工具函数，可以调用它跳过特定架构上的测试。

Other tests should pass reliably, it is acceptable to skip some of them as long as most of them will pass. The following tests should not be skipped even if you will be tempted to:

​	其他测试应该可靠地通过，尽管可以跳过其中一些，只要大多数测试能通过。即使你会被诱惑跳过，也不应跳过以下测试：

- `proc.TestNextConcurrent`
- `proc.TestNextConcurrentVariant2`
- `proc.TestBreakpointCounts` (enable `proc.TestBreakpointCountsWithDetection` if you have problems with this)
  - `proc.TestBreakpointCounts`（如果遇到问题，可以启用`proc.TestBreakpointCountsWithDetection`）
- `proc.TestStepConcurrentDirect`
- `proc.TestStepConcurrentPtr`

### 移植到大端架构 Porting to Big Endian architectures



Delve was initially written for amd64 and assumed 64bit and little endianness everywhere. The assumption on pointer size has been removed throughout the codebase, the assumption about endianness hasn't. Both `pkg/dwarf/loclist` and `pkg/dwarf/frame` incorrectly assume little endian encoding. Other parts of the code might do the same.

​	Delve最初是为amd64编写的，假设在所有地方都使用64位和小端字节序。指针大小的假设已经从代码中删除，但关于字节序的假设仍然存在。`pkg/dwarf/loclist`和`pkg/dwarf/frame`都错误地假设了小端编码。代码的其他部分也可能存在类似的假设。

### 移植到ARM32、MIPS和软件单步调试 Porting to ARM32, MIPS and Software Single Stepping



When resuming a thread stopped at a breakpoint Delve will:

​	当恢复线程时，Delve将：

1. remove the breakpoint temporarily 临时移除断点
2. make the thread execute a single instruction 执行单个指令
3. write the breakpoint back 重新设置断点

Step 2 is implemented using the hardware single step feature that many CPUs have and that is exposed via PTRACE_SINGLESTEP or similar features. ARM32 and MIPS do not have a hardware single step implemented, this means that it has to be implemented in software. The linux kernel used to have a software implementation of PTRACE_SINGLESTEP but those have been removed because they were too burdensome to maintain and delegated the feature to userspace debuggers entirely.

​	步骤2通过许多CPU具有的硬件单步调试功能来实现，该功能通过PTRACE_SINGLESTEP或类似的功能暴露出来。ARM32和MIPS没有实现硬件单步调试功能，这意味着必须通过软件实现。Linux内核曾经有一个PTRACE_SINGLESTEP的软实现，但由于维护困难，它们已被移除，改由用户空间调试器完全处理。

A software singlestep implementation would work like this:

​	软件单步调试的实现方式如下：

1. decode the current instruction 解码当前指令
2. figure out all the possible places where the PC registers could be after executing the current instruction 找出执行当前指令后PC寄存器可能的所有位置
3. set a breakpoint on all of them 在这些位置上设置断点
4. resume the thread normally 正常恢复线程
5. clear all breakpoints created on step 3 清除步骤3中设置的所有断点

Delve does not currently have any infrastructure to help implement this, which means that porting to architectures without hardware singlestep is even more complicated.

​	Delve当前没有帮助实现此功能的基础设施，这意味着在没有硬件单步调试功能的架构上进行移植会更加复杂。
