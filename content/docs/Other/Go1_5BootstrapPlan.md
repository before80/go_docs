+++
title = "go 1.5 Bootstrap计划"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Go 1.5 Bootstrap Plan - go 1.5 Bootstrap计划

> 原文：[https://docs.google.com/document/d/1OaatvGhEAq7VseQ9kkavxKNAfepWy2yhPUBs96FGV28/edit](https://docs.google.com/document/d/1OaatvGhEAq7VseQ9kkavxKNAfepWy2yhPUBs96FGV28/edit)

golang.org/s/go15bootstrap

([comments on golang-dev](https://www.google.com/url?q=https://groups.google.com/d/msg/golang-dev/3bTIOleL8Ik/D8gICLOiUJEJ&sa=D&source=editors&ust=1669904018103563&usg=AOvVaw1DLJEz3FJWZF4mVZ-SdIIZ))

Russ Cox

January 2015

## Abstract 摘要

Go 1.5 will use a toolchain written in Go (at least in part).

Go 1.5将使用由Go编写的工具链（至少是部分）。

Question: how do you build Go if you need Go built already?

问题：如果你需要Go已经构建好了，那么你如何构建Go？

Answer: building Go 1.5 will require having Go 1.4 available.

回答：构建 Go 1.5 需要有 Go 1.4 可用。

## Background 背景介绍

We have been planning for a year now to eliminate all C programs from the Go source tree. The C compilers (5c, 6c, 8c, 9c) have already been removed. The remaining C programs will be converted to Go: they are the Go compilers ([golang.org/s/go13compiler](https://www.google.com/url?q=https://golang.org/s/go13compiler&sa=D&source=editors&ust=1669904018104772&usg=AOvVaw3K8jVYnoviw9iACYfkHSEw)), the assemblers, the linkers ([golang.org/s/go13linker](https://www.google.com/url?q=https://golang.org/s/go13linker&sa=D&source=editors&ust=1669904018104998&usg=AOvVaw2JHWfrTIVMS9bKCGqOBfvQ)), and cmd/dist. If these programs are written in Go, that introduces a bootstrapping problem when building completely from source code: you need a working Go toolchain in order to build a Go toolchain.

我们已经计划了一年，以消除 Go 源代码树中的所有 C 程序。C语言编译器（5c、6c、8c、9c）已经被删除。剩下的C程序将被转换为Go：它们是Go编译器（golang.org/s/go13compiler）、汇编器、链接器（golang.org/go13linker）以及cmd/dist。如果这些程序是用Go编写的，那么在完全从源代码构建时就会引入一个引导问题：你需要一个可以工作的Go工具链来构建一个Go工具链。

## Proposal 提议

To build Go 1.x, for x ≥ 5, it will be necessary to have Go 1.4 (or newer) installed already, in `$GOROOT_BOOTSTRAP`. The default value of `$GOROOT_BOOTSTRAP` is `$HOME/go1.4`. In general we'll keep using Go 1.4 as the bootstrap base version for as long as possible. The toolchain proper (compiler, assemblers, linkers) will need to be buildable with Go 1.4, whether by restricting their feature use to what is in Go 1.4 or by using build tags.

要构建 Go 1.x，对于 x ≥ 5，必须在 `$GOROOT_BOOTSTRAP` 中已经安装了 Go 1.4（或更新版本）。`$GOROOT_BOOTSTRAP`的默认值是`$HOME/go1.4`。一般来说，我们会尽可能长时间地使用Go 1.4作为引导的基础版本。工具链本身（编译器、汇编器、链接器）将需要在Go 1.4中构建，无论是通过将其功能使用限制在Go 1.4中的内容还是通过使用构建标签。

For comparison with what will follow, the old build process for Go 1.4 is:

为了与后面的内容相比较，Go 1.4的旧构建过程是：

1. Build cmd/dist with gcc (or clang).
2. Using dist, build compiler toolchain with gcc (or clang)
3. NOP
4. Using dist, build cmd/go (as go_bootstrap) with compiler toolchain.
5. Using go_bootstrap, build the remaining standard library and commands.
6. 用gcc（或clang）构建cmd/dist。
7. 使用dist，用gcc（或clang）构建编译器工具链。
8. NOP
9. 使用dist，用编译器工具链构建cmd/go（作为go_bootstrap）。
10. 使用go_bootstrap，构建其余的标准库和命令。

The new build process for Go 1.x (x ≥ 5) will be:

Go 1.x (x ≥ 5) 的新构建过程将是：

1. Build cmd/dist with Go 1.4.
2. Using dist, build Go 1.x compiler toolchain with Go 1.4.
3. Using dist, rebuild Go 1.x compiler toolchain with itself.
4. Using dist, build Go 1.x cmd/go (as go_bootstrap) with Go 1.x compiler toolchain.
5. Using go_bootstrap, build the remaining Go 1.x standard library and commands.
6. 用 Go 1.4 构建 cmd/dist。
7. 使用 dist，用 Go 1.4 构建 Go 1.x 编译器工具链。
8. 使用dist，用它来重建Go 1.x编译器工具链。
9. 使用 dist，用 Go 1.x 编译器工具链构建 Go 1.x cmd/go（作为 go_bootstrap）。
10. 使用 go_bootstrap 构建剩余的 Go 1.x 标准库和命令。

There are two changes.

有两个变化。

The first change is that we replace gcc (or clang) with Go 1.4.

第一个变化是我们用Go 1.4代替gcc（或clang）。

The second change is the introduction of step 3, which rebuilds the Go 1.x compiler toolchain with itself. The 6g built in Step 2 is a Go 1.x compiler built using Go 1.4 libraries and compilers. The 6g built in Step 3 is the same Go 1.x compiler, but built using Go 1.x libraries and compilers. If Go 1.x has changed the format of debug info or some other detail of the binaries, it may matter to tools whether 6g is a Go 1.4 binary or a Go 1.x binary. If Go 1.x has introduced any performance or stability improvements in the libraries, the compiler in Step 3 will be faster or more stable than the compiler in Step 2. Of course, if Go 1.x is buggier, the 6g built in Step 3 will also be buggier, so it will be possible to disable step 3 for debugging.

第二个变化是引入了步骤3，它用自己重建了Go 1.x编译器工具链。在步骤2中构建的6g是一个使用Go 1.4库和编译器构建的Go 1.x编译器。步骤3中构建的6g是同一个Go 1.x编译器，但使用Go 1.x库和编译器构建。如果 Go 1.x 改变了调试信息的格式或二进制文件的一些其他细节，那么 6g 是 Go 1.4 二进制文件还是 Go 1.x 二进制文件对工具来说可能很重要。如果Go 1.x在库中引入了任何性能或稳定性的改进，步骤3中的编译器将比步骤2中的编译器更快或更稳定。当然，如果Go 1.x的bug较多，步骤3中构建的6g也会有bug，所以可以禁用步骤3进行调试。

Step 3 could make make.bash take longer. As an upper bound on the slowdown, the current build process steps 1-4 take 20 seconds on my MacBook Pro, out of the total 40 seconds required for make.bash. In the new process, I can’t see step 3 adding more than 50% to the make.bash run time, and I expect it would be significantly less than that. On the other hand, the C compilations being replaced are very I/O heavy; two Go compilations might still be faster, especially on I/O-constrained ARM devices. In any event, if make.bash does slow down, I will speed up run.bash at least as much, so that all.bash time does not increase.

第3步可能会使make.bash花费更长的时间。作为减速的上限，在我的MacBook Pro上，当前构建过程的第1-4步需要20秒，而make.bash总共需要40秒。在新的过程中，我看不出步骤3会给make.bash的运行时间增加50%以上，而且我估计会比这少得多。另一方面，被替换的C语言编译的I/O量非常大；两个Go语言编译可能仍然更快，特别是在I/O受限的ARM设备上。无论如何，如果make.bash确实变慢了，我将至少加快run.bash的速度，这样所有.bash的时间就不会增加。

## New Ports 新端口

Bootstrapping makes new ports a little more complex. It was possible in the past to check out the Go tree on a new system and run all.bash to build the toolchain (and it would fail, and you’d make some edits, and try again). Now, it will not be possible to run all.bash until that system is fully supported by Go.

Bootstrapping 使得新的 port 更加复杂。过去有可能在一个新的系统上检查 Go 树， 然后运行 all.bash 来构建工具链 (但会失败， 你会做一些编辑， 然后再试一次)。现在，在该系统被Go完全支持之前，是不可能运行all.bash的。

For Go 1.x (x ≥ 5), new ports will have to be done by cross-compiling test binaries on a working system, copying the binaries over to the target, and running and debugging them there. This is already well-supported by all.bash via the go_$GOOS_$GOARCH_exec scripts (see ‘go help run’). Once all.bash can be run in that mode, the resulting compilers and libraries can be copied to the target system and used directly.

对于 Go 1.x (x ≥ 5)，新的移植必须通过在工作系统上交叉编译测试二进制文件，将二进制文件复制到目标系统，并在那里运行和调试。all.bash 已经通过 go_$GOOS_$GOARCH_exec 脚本很好地支持了这一点 (参见 'go help run')。一旦 all.bash 能够以这种模式运行， 所产生的编译器和库就可以被复制到目标系统中并直接使用。

Once a port works well enough that the compilers and linkers can run on the target machine, the script bootstrap.bash (run on an old system) will prepare a GOROOT_BOOTSTRAP directory for use on the new system.

一旦移植工作顺利，编译器和链接器可以在目标机器上运行，脚本bootstrap.bash（在旧系统上运行）将准备一个GOROOT_BOOTSTRAP目录，在新系统上使用。

## Deployment 部署

Today we are still using the Go 1.4 build process above.

今天我们仍然使用上面的Go 1.4构建过程。

The first step in the transition will be to convert cmd/dist itself to Go and change make.bash to use Go 1.4 to build cmd/dist. That replaces “gcc (or clang)” with “Go 1.4” in step 1 of the build and changes nothing else. This will mainly exercise the integration of Go 1.4 into the build.

过渡的第一步将是把 cmd/dist 本身转换为 Go，并修改 make.bash 以使用 Go 1.4 来构建 cmd/dist。这将在构建的第一步中用 "Go 1.4 "取代 "gcc (或clang)"，而不改变其他内容。这将主要锻炼 Go 1.4 在构建中的集成度。

After that first step, we can convert the remaining C programs in whatever order makes sense. Each conversion will require minor modifications to cmd/dist to build the Go version instead of the C version. I am not sure whether the new linker or the new assemblers will be converted first. I expect the Go compiler to be converted last.

在第一步之后，我们可以按照任何有意义的顺序转换其余的C程序。每一次转换都需要对cmd/dist进行小的修改，以构建Go版本而不是C版本。我不确定是新的链接器还是新的汇编器将首先被转换。我希望 Go 编译器能最后转换。

We will probably do the larger conversions on the dev.cc branch and merge into master at good checkpoints, so that multiple people can work on the conversion (coordinated via Git) but able to break certain builds for long amounts of time without affecting other developers. This is similar to what we did for dev.cc and dev.garbage in 2014.

我们可能会在 dev.cc 分支上进行较大的转换，并在良好的检查点上合并到 master 分支上，这样多人就可以在转换过程中工作（通过 Git 协调），但又能在不影响其他开发者的情况下长期中断某些构建。这与我们在2014年为dev.cc和dev.garbage所做的类似。

Go 1.5 will require Go 1.4 to build. The goal is to convert all the C programs—the Go compiler, the linker, the assemblers, and cmd/dist—for Go 1.5. We may not reach that goal, but certainly some of that list will be converted.

Go 1.5将需要Go 1.4来构建。我们的目标是将所有的C程序——Go编译器、链接器、汇编器和cmd/dist转换为Go 1.5。我们可能达不到这个目标，但肯定会转换其中一些程序。