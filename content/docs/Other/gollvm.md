+++
title = "Gollvm"
weight = 4
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Gollvm

> 原文：[https://go.googlesource.com/gollvm/](https://go.googlesource.com/gollvm/)

Gollvm is an LLVM-based Go compiler. It incorporates "gofrontend" (a Go language front end written in C++ and shared with GCCGO), a bridge component (which translates from gofrontend IR to LLVM IR), and a driver that sends the resulting IR through the LLVM back end.

Gollvm是一个基于LLVM的Go编译器。它包含 "gofrontend"（用C++编写的Go语言前端，与GCCGO共享），一个桥接组件（将gofrontend IR翻译成LLVM IR），以及一个通过LLVM后端发送IR结果的驱动。

Gollvm is set up to be a subproject within the LLVM tools directory, similar to how things work for "clang" or "compiler-rt": you check out a copy of the LLVM source tree, then within the LLVM tree you check out additional git repos.

Gollvm被设置为LLVM工具目录下的一个子项目，类似于 "clang "或 "compiler-rt "的工作方式：您检查出LLVM源代码树的副本，然后在LLVM树中检查出额外的git repos。



# Building gollvm - 构建gollvm

Gollvm is currently in development -- releases are not yet available for download. Instructions for building gollvm follow.

Gollvm目前正在开发中 -- 版本还不能下载。构建gollvm的说明如下。

## Setting up a gollvm work area 设置一个gollvm工作区

To set up a work area for Gollvm, check out a copy of LLVM, the overlay the gollvm repo (and other associated dependencies) within the LLVM tools subdir, as follows:

要为Gollvm建立一个工作区，请查看LLVM的副本，将gollvm repo（以及其他相关的依赖项）覆盖在LLVM工具子目录中，如下所示：

```
// Here 'workarea' will contain a copy of the LLVM source tree and one or more build areas
% mkdir workarea
% cd workarea

// Sources
% git clone https://github.com/llvm/llvm-project.git
...
% cd llvm-project/llvm/tools
% git clone https://go.googlesource.com/gollvm
...
% cd gollvm
% git clone https://go.googlesource.com/gofrontend
...
% cd libgo
% git clone https://github.com/libffi/libffi.git
...
% git clone https://github.com/ianlancetaylor/libbacktrace.git
...
%
```

## Building gollvm with cmake and ninja 用cmake和ninja构建gollvm

You'll need to have an up-to-date copy of cmake on your system (3.6 or later vintage) to build Gollvm, as well as a C/C++ compiler (V10.0 or later for Clang, or V6.0 or later of GCC), and a working copy of ‘m4’.

您需要在您的系统上有一个最新的cmake副本（3.6或更高版本）来构建Gollvm，以及一个C/C++编译器（Clang的V10.0或更高版本，或GCC的V6.0或更高版本），和一个'm4'的工作副本。

Create a build directory (separate from the source tree) and run ‘cmake’ within the build area to set up for the build. Assuming that ‘workarea’ is the directory created as above:

创建一个构建目录（与源代码树分开），并在构建区域内运行'cmake'为构建做准备。假设'workarea'是上面创建的目录：

```
% cd workarea
% mkdir build-debug
% cd build-debug
% cmake -DCMAKE_BUILD_TYPE=Debug -DLLVM_USE_LINKER=gold -G Ninja ../llvm-project/llvm
...
% ninja gollvm
...
%
```

This will build the various tools and libraries needed for Gollvm. To select a specific C/C++ compiler for the build, you can use the "-DCMAKE_C_COMPILER" and "-DCMAKE_CXX_COMPILER" options to select your desired C/C++ compiler when invoking cmake (details [here](https://gitlab.kitware.com/cmake/community/wikis/FAQ#how-do-i-use-a-different-compiler)). Use the "-DLLVM_USE_LINKER=" cmake variable to control which linker is selected to link the Gollvm compiler and tools (where variant is one of "bfd", "gold", "lld", etc).

这将构建Gollvm所需的各种工具和库。要选择特定的C/C++编译器进行构建，您可以使用"-DCMAKE_C_COMPILER "和"-DCMAKE_CXX_COMPILER "选项，在调用cmake时选择您想要的C/C++编译器（细节在这里）。使用"-DLLVM_USE_LINKER="cmake变量来控制选择哪个链接器来链接Gollvm编译器和工具（其中变量是 "bfd"、"gold"、"ld "等之一）。

The Gollvm compiler driver defaults to using the gold linker when linking Go programs. If some other linker is desired, this can be accomplished by passing "-DGOLLVM_DEFAULT_LINKER=" when running cmake. Note that this default can still be overridden on the command line using the "-fuse-ld" option.

Gollvm编译器驱动程序默认在链接Go程序时使用gold链接器。如果需要其他链接器，可以在运行cmake时通过"-DGOLLVM_DEFAULT_LINKER="来完成。请注意，这个默认值仍然可以在命令行中使用"-fuse-ld "选项进行覆盖。

Gollvm's cmake rules expect a valid value for the SHELL environment variable; if not set, a default shell of /bin/bash will be used.

Gollvm的cmake规则期望SHELL环境变量的有效值；如果没有设置，将使用默认的/bin/bash shell。

## Installing gollvm 安装gollvm

A gollvm installation will contain ‘llvm-goc’ (the compiler driver), the libgo standard Go libraries, and the standard Go tools ("go", "vet", "cgo", etc).

gollvm的安装将包含'llvm-goc'（编译器驱动程序）、libgo标准Go库和标准Go工具（"go"、"vet"、"cgo "等等）。

The installation directory for gollvm needs to be specified when invoking cmake prior to the build:

在编译前调用 cmake 时需要指定 gollvm 的安装目录：

```
% mkdir build.rel
% cd build.rel
% cmake -DCMAKE_INSTALL_PREFIX=/my/install/dir -DCMAKE_BUILD_TYPE=Release -DLLVM_USE_LINKER=gold -G Ninja ../llvm-project/llvm

// Build all of gollvm
% ninja gollvm
...

// Install gollvm to "/my/install/dir"
% ninja install-gollvm
```

## Using an installed copy of gollvm 使用已安装的gollvm副本

Programs build with the Gollvm Go compiler default to shared linkage, meaning that they need to pick up the Go runtime library via LD_LIBRARY_PATH:

用Gollvm Go编译器编译的程序默认为共享链接，这意味着它们需要通过LD_LIBRARY_PATH来获取Go运行库：

```
// Root of Gollvm install is /tmp/gollvm-install
% export LD_LIBRARY_PATH=/tmp/gollvm-install/lib64
% export PATH=/tmp/gollvm-install/bin:$PATH
% go run himom.go
hi mom!
%
```

# Information for gollvm developers 给gollvm开发者的信息

## Source code structure 源代码结构

Within <workarea>/llvm/tools/gollvm, the following directories are of interest:

在<workarea>/llvm/tools/gollvm中，以下目录值得关注：

.../llvm/tools/gollvm:

- contains rules to build third party libraries needed for gollvm, along with common definitions for subdirs.包含构建gollvm所需的第三方库的规则，以及子目录的通用定义。

.../llvm/tools/gollvm/driver, .../llvm/tools/gollvm/driver-main:

- contains build rules and source code for llvm-goc 包含llvm-goc的构建规则和源代码。

.../llvm/tools/gollvm/gofrontend:

- source code for gofrontend and libgo (note: no cmake files here) gofrontend和libgo的源代码（注意：这里没有cmake文件）。

.../llvm/tools/gollvm/bridge:

- contains build rules for the libLLVMCppGoFrontEnd.a, a library that contains both the gofrontend code and the LLVM-specific middle layer (for example, the definition of the class Llvm_backend, which inherits from Backend).包含libLLVMCppGoFrontEnd.a的构建规则，这个库同时包含gofrontend代码和LLVM特有的中间层（例如，Llvm_backend类的定义，它继承于Backend）。

.../llvm/tools/gollvm/libgo:

- build rules and supporting infrastructure to build Gollvm's copy of the Go runtime and standard packages.构建规则和支持基础设施，用于构建Gollvm的Go运行时副本和标准包。

.../llvm/tools/gollvm/unittests:

- source code for the unit tests 单元测试的源代码

## The llvm-goc program - llvm-goc程序

The executable llvm-goc is the main compiler driver for gollvm; it functions as a compiler (consuming source for a Go package and producing an object file), an assembler, and/or a linker. While it is possible to build and run llvm-goc directly from the command line, in practice there is little point in doing this (better to build using "go build", which will invoke llvm-goc on your behalf.

可执行的llvm-goc是gollvm的主要编译器驱动；它的功能是编译器（消耗Go包的源代码并产生一个目标文件）、汇编器和/或链接器。虽然可以直接从命令行中构建和运行llvm-goc，但实际上这样做没有什么意义（最好使用 "go build "来构建，它将代表您调用llvm-goc。

```
// From within <workarea>/build.opt:

% ninja llvm-goc
...
% cat micro.go
package foo
func Bar() int {
	return 1
}
% ./bin/llvm-goc -fgo-pkgpath=foo -O3 -S -o micro.s micro.go
%
```

## Building and running the unit tests 构建和运行单元测试

Here are instructions on building and running the unit tests for the middle layer:

以下是关于构建和运行中间层的单元测试的说明：

```
// From within <workarea>/build.opt:

// Build unit test
% ninja GoBackendCoreTests

// Run a unit test
% ./tools/gollvm/unittests/BackendCore/GoBackendCoreTests
[==========] Running 10 tests from 2 test cases.
[----------] Global test environment set-up.
[----------] 9 tests from BackendCoreTests
[ RUN      ] BackendCoreTests.MakeBackend
[       OK ] BackendCoreTests.MakeBackend (1 ms)
[ RUN      ] BackendCoreTests.ScalarTypes
[       OK ] BackendCoreTests.ScalarTypes (0 ms)
[ RUN      ] BackendCoreTests.StructTypes
[       OK ] BackendCoreTests.StructTypes (1 ms)
[ RUN      ] BackendCoreTests.ComplexTypes
[       OK ] BackendCoreTests.ComplexTypes (0 ms)
[ RUN      ] BackendCoreTests.FunctionTypes
[       OK ] BackendCoreTests.FunctionTypes (0 ms)
[ RUN      ] BackendCoreTests.PlaceholderTypes
[       OK ] BackendCoreTests.PlaceholderTypes (0 ms)
[ RUN      ] BackendCoreTests.ArrayTypes
[       OK ] BackendCoreTests.ArrayTypes (0 ms)
[ RUN      ] BackendCoreTests.NamedTypes
[       OK ] BackendCoreTests.NamedTypes (0 ms)
[ RUN      ] BackendCoreTests.TypeUtils

...

[  PASSED  ] 10 tests.
```

The unit tests currently work by instantiating an LLVM Backend instance and making backend method calls (to mimic what the frontend would do), then inspects the results to make sure they are as expected. Here is an example:

该单元测试目前的工作方式是实例化LLVM后端实例，并进行后端方法调用（模仿前端会做什么），然后检查结果以确保它们符合预期。下面是一个例子：

```
TEST(BackendCoreTests, ComplexTypes) {
  LLVMContext C;

  Type *ft = Type::getFloatTy(C);
  Type *dt = Type::getDoubleTy(C);

  std::unique_ptr<Backend> be(go_get_backend(C, llvm::CallingConv::X86_64_SysV));
  Btype *c32 = be->complex_type(64);
  ASSERT_TRUE(c32 != NULL);
  ASSERT_EQ(c32->type(), mkTwoFieldLLvmStruct(C, ft, ft));
  Btype *c64 = be->complex_type(128);
  ASSERT_TRUE(c64 != NULL);
  ASSERT_EQ(c64->type(), mkTwoFieldLLvmStruct(C, dt, dt));
}
```

The test above makes sure that the LLVM type we get as a result of calling Backend::complex_type() is kosher and matches up to expectations.

上面的测试确保了我们在调用Backend::complex_type()后得到的LLVM类型是正确的，并且与预期相符。

## Building libgo (Go runtime and standard libraries) 构建 libgo (Go 运行时和标准库)

To build the Go runtime and standard libraries, use the following:

要构建Go运行时和标准库，请使用以下方法：

```
// From within <workarea>/build.opt:

// Build Go runtime and standard libraries
% ninja libgo_all
```

This will compile static (*.a) and dynamic (*.so) versions of the library.

这将编译静态（*.a）和动态（*.so）版本的库。

# FAQ 常见问题

## Where should I post questions about gollvm? 我应该在哪里发表关于gollvm的问题？

Please send questions about gollvm to the [golang-nuts](https://groups.google.com/d/forum/golang-nuts) mailing list. Posting questions to the issue tracker is generally not the right way to start discussions or get information.

请将关于gollvm的问题发送到golang-nuts邮件列表。在问题追踪器上发布问题通常不是开始讨论或获取信息的正确方式。

## Where should I file gollvm bugs? 我应该在哪里提交gollvm的bug？

Please file an issue on the golang [issue tracker](https://github.com/golang/go/issues); please be sure to use "gollvm" somewhere in the headline.

请在golang问题追踪器上提交问题；请确保在标题中使用 "gollvm"。

## How can I go about contributing to gollvm? 我怎样才能为gollvm做贡献？

Please see the Go project guidelines at https://golang.org/doc/contribute.html. Changes to https://go.googlesource.com/gollvm can be made by any Go contributor; for changes to gofrontend see [the gccgo guidelines](https://golang.org/doc/gccgo_contribute.html).

请参阅Go项目指南：https://golang.org/doc/contribute.html。任何 Go 贡献者都可以对 https://go.googlesource.com/gollvm 进行修改；对 gofrontend 的修改请参见 gccgo 指南。

## Is gollvm a replacement for the main Go compiler? (gc) gollvm是否可以替代主Go编译器？ (gc)

Gollvm is not intended as a replacement for the main Go compiler -- the expectation is that the bulk of users will want to continue to use the main Go compiler due to its superior compilation speed, ease of use, broader functionality, and higher-performance runtime. Gollvm is intended to provide a Go compiler with a more powerful back end, enabling such benefits as better inlining, vectorization, register allocation, etc.

Gollvm 并不是要取代主 Go 编译器 -- 我们的期望是大部分用户会继续使用主 Go 编译器，因为它有卓越的编译速度、易用性、更广泛的功能和更高性能的运行时间。Gollvm旨在为Go编译器提供一个更强大的后端，实现更好的内联、矢量化、寄存器分配等优点。

## Which architectures and operating systems are supported for gollvm? - Gollvm支持哪些架构和操作系统？

Gollvm is currently supported only for x86_64 and aarch64 Linux.

Gollvm目前只支持x86_64和arch64 Linux。

## How does the gollvm runtime differ from the main Go runtime? - gollvm运行时与主Go运行时有何不同？

The main Go runtime supports generation of accurate stack maps, which allows the garbage collector to do precise stack scanning; gollvm does not yet support stack map generation (note that we're actively working on fixing this), hence for gollvm the garbage collector has to scan stacks conservatively (which can lead to longer scan times and increased memory usage). The main Go runtime compiles to a different calling convention, whereas Gollvm uses the standard C/C++ calling convention. There are many other smaller differences as well.

主Go运行时支持生成精确的堆栈图，这使得垃圾收集器可以进行精确的堆栈扫描；而gollvm还不支持堆栈图的生成（注意我们正在积极解决这个问题），因此对于gollvm，垃圾收集器必须保守地扫描堆栈（这可能导致更长的扫描时间和更多的内存使用）。Go的主运行时编译时采用了不同的调用约定，而Gollvm则使用标准的C/C++调用约定。还有许多其他较小的区别。

## Shared linkage is the default for gollvm. How do I build non-shared? 共享链接是gollvm的默认方式。我如何构建非共享的？

Linking with "-static-libgo" will yield a binary that incorporates a full copy of the Go runtime. Example:

使用"-static-libgo "链接将产生一个包含Go运行时完整副本的二进制文件。例子：

```
 % go build -gccgoflags -static-libgo myprogram.go
```

Note that this will increase binary size.

注意，这将增加二进制文件的大小。

## What command line options are supported for gollvm? - gollvm支持哪些命令行选项？

You can run ‘llvm-goc -help’ to see a full set of supported options. These can be passed to the compiler via ‘-gccgoflags’ option. Example:

您可以运行 "llvm-goc -help "来查看一整套支持的选项。这些可以通过'-gccgoflags'选项传递给编译器。例子：

```
% go build -gccgoflags -fno-inline mumble.go
```

## How do I see the LLVM IR generated by gollvm? 我怎样才能看到gollvm生成的LLVM IR？

The ‘llvm-goc’ command supports the -emit-llvm flag, however passing this option to a "go build" command is not practical, since the "go build" won't be expecting the compiler to emit LLVM bitcode or assembly.

llvm-goc "命令支持-emit-llvm标志，但是把这个选项传递给 "go build "命令并不实际，因为 "go build "不会期望编译器发出LLVM位码或汇编。

A better recipe is to run "go build" with "-x -work" to capture the commands being executed, then rerun the llvm-goc command shown adding "-S -emit-llvm". The resulting output will be an LLVM IR dump. Example:

一个更好的方法是用"-x -work "来运行 "go build"，以捕获正在执行的命令，然后重新运行llvm-goc命令，并添加"-S -emit-llvm"。结果输出将是一个LLVM IR dump。例子：

```
% go build -work -x mypackage.go 1> transcript.txt 2>&1
% egrep '(WORK=|llvm-goc -c)' transcript.txt
WORK=/tmp/go-build887931787
/t/bin/llvm-goc -c -g -m64 -fdebug-prefix-map=$WORK=/tmp/go-build \
  -gno-record-gcc-switches -fgo-pkgpath=command-line-arguments \
  -fgo-relative-import-path=/mygopath/src/tmp -o $WORK/b001/_go_.o \
  -I $WORK/b001/_importcfgroot_ ./mypackage.go
% /t/bin/llvm-goc -c -g -m64 -fdebug-prefix-map=$WORK=/tmp/go-build \
  -gno-record-gcc-switches -fgo-pkgpath=command-line-arguments \
  -fgo-relative-import-path=/mygopath/src/tmp \
  -I $WORK/b001/_importcfgroot_ -o mypackage.ll -S -emit-llvm \
  ./mypackage.go
% ls -l mypackage.ll
...
%
```

## What is the relationship between gollvm and gccgo? - gollvm和gccgo之间是什么关系？

Gollvm and gccgo share a common front end (gofrontend) and associated runtime (libgo), however each uses a separate back end. When using "go build", the Go command currently treats gollvm as an instance of gccgo (hence the need to pass compile flags via "-gccgoflags"). This is expected to be temporary.

Gollvm和gccgo共享一个共同的前端（gofrontend）和相关的运行时（libgo），但各自使用一个单独的后端。当使用 "go build "时，Go命令目前将gollvm视为gccgo的一个实例（因此需要通过"-gccgoflags "传递编译标志）。预计这将是暂时的。

## Can I use FDO or ThinLTO with gollvm? 我可以在gollvm中使用FDO或ThinLTO吗？

There are plans to support FDO, AutoFDO, and ThinLTO for gollvm, however these features have not yet been implemented.

有计划为gollvm支持FDO、AutoFDO和ThinLTO，但是这些功能还没有实现。

## Can I use the race detector? 我可以使用竞争检测器吗？

Gollvm does not support the Go race detector; please use the main Go compiler for this purpose.

Gollvm不支持Go竞争检测器；请使用主Go编译器来实现这一目的。

## I am seeing "undefined symbol: `__get_cpuid_count`" from my gollvm install  我看到 "未定义的符号：__get_cpuid_count" 从我的gollvm安装中看到

The Gollvm build procedure requires an up-to-date C/C++ compiler; there is code in the gollvm runtime (libgo) that refers to functions defined in `<cpuid.h>`, however some older versions of clang (prior to 5.0) don't provide definitions for all the needed functions. If you encounter this problem, rerun `cmake` to configure your build to use a more recent version of Clang (or use GCC), as described above.

Gollvm的构建过程需要一个最新的C/C++编译器；在gollvm运行时（libgo）中有一些代码引用了<cpuid.h>中定义的函数，然而一些旧版本的clang（5.0以前）并没有提供所有需要的函数的定义。如果您遇到这个问题，请重新运行cmake来配置您的构建，以使用较新版本的clang（或使用GCC），如上所述。