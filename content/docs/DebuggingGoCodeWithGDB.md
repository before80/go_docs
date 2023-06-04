+++
title = "使用GDB调试 go 代码"
weight = 22
date = 2023-05-18T17:31:23+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Debugging Go Code with GDB - 使用GDB调试 go 代码

> 原文：[https://go.dev/doc/gdb](https://go.dev/doc/gdb)

​	下面的说明适用于标准工具链(gc Go编译器和工具)。Gccgo具有原生的gdb支持。

> ​	​请注意，在使用标准工具链构建的Go程序进行调试时，[Delve](https://github.com/go-delve/delve)是比GDB更好的选择。它比GDB更好地理解Go运行时、数据结构和表达式。Delve目前支持Linux、OSX和Windows的amd64平台。有关支持的最新平台列表，请参阅[Delve文档](https://github.com/go-delve/delve/tree/master/Documentation/installation)。
>

​	GDB不太理解Go程序。堆栈管理、线程和运行时包含与GDB预期执行模型有足够不同的方面，即使使用gccgo编译程序，它们也可能混淆调试器并导致不正确的结果。因此，尽管在某些情况下(例如调试Cgo代码或调试运行时本身)GDB可能有用，但对于Go程序，特别是高度并发的程序，它不是一个可靠的调试器。此外，解决这些困难的问题不是Go项目的重点。

​	简而言之，下面的说明只应被视为在GDB正常工作时使用GDB的指南，而不是成功的保证。除了这个概述，您可能还想参考[GDB手册](https://sourceware.org/gdb/current/onlinedocs/gdb/)。

## 简介

​	当您在Linux、macOS、FreeBSD或NetBSD上使用gc工具链编译和链接Go程序时，生成的二进制文件包含DWARFv4调试信息，GDB调试器的最新版本(≥7.5)可以使用它来检查实时进程或核心转储。

​	将"`-w`"标志传递给链接器以省略调试信息(例如，`go build -ldflags=-w prog.go`)。

​	gc编译器生成的代码包括函数调用的内联和变量的寄存器化。这些优化有时会使使用gdb进行调试更加困难。如果您发现需要禁用这些优化，请使用`go build -gcflags=all="-N -l"`构建您的程序。

​	如果您想使用gdb检查核心转储，请在环境中设置`GOTRACEBACK=crash`，以在系统允许的情况下在程序崩溃时触发转储(有关更多信息，请参阅[运行时包文档](https://go.dev/pkg/runtime/))。

### 常见操作

- 显示代码的文件和行号，设置断点和反汇编：

  ```
  (gdb) list
  (gdb) list line
  (gdb) list file.go:line
  (gdb) break line
  (gdb) break file.go:line
  (gdb) disas
  ```

- 显示回溯和展开栈帧：

  ```
  (gdb) bt
  (gdb) frame n
  ```

- 显示本地变量、参数和返回值在栈帧中的名称、类型和位置：

  ```
  (gdb) info locals
  (gdb) info args
  (gdb) p variable
  (gdb) whatis variable
  ```

- 显示全局变量的名称、类型和位置：

  ```
  (gdb) info variables regexp
  ```

### Go扩展

​	最近的扩展机制允许GDB加载给定二进制文件的扩展脚本。工具链使用这个机制通过一些命令来扩展GDB，以便检查运行时代码(如goroutine)的内部情况，并漂亮地打印内置的map、slice和channel类型。

- 漂亮打印字符串、slice、map、channel或interface：

  ```
  (gdb) p var
  ```

- 针对字符串、slice和map的`$len()`和`$cap()`函数：

  ```
  (gdb) p $len(var)
  ```

-  用于将接口转换为它们的动态类型的函数：

  ```
  (gdb) p $dtype(var)
  (gdb) iface var
  ```

  **已知问题**：如果接口值的长名称与其短名称不同，则GDB无法自动找到其动态类型(在打印栈跟踪时很烦人，漂亮打印器退回到打印短类型名称和指针)。

- 检查goroutines：

  ```
  (gdb) info goroutines
  (gdb) goroutine n cmd
  (gdb) help goroutine
  ```

  例如：

  ```
  (gdb) goroutine 12 bt
  ```

  您可以通过传递`all`而不是特定goroutine的ID来检查所有goroutine。例如： 

  ```
  (gdb) goroutine all bt
  ```

​	如果您想了解如何工作或扩展它，请查看Go源发行版中的[src/runtime/runtime-gdb.py](https://go.dev/src/runtime/runtime-gdb.py)。它依赖于一些特殊的魔法类型(`hash<T,U>`)和变量(runtime.m和runtime.g)，链接器([src/cmd/link/internal/ld/dwarf.go](https://go.dev/src/cmd/link/internal/ld/dwarf.go))确保在DWARF代码中描述它们。

​	如果您对调试信息的外观感兴趣，请运行`objdump -W a.out`，并浏览`.debug_*`部分。

### 已知问题

1. 字符串的漂亮打印仅对类型为string的字符串触发，对于派生类型则无效。
2. C部分的运行时库缺少类型信息。
3. GDB不理解Go的名称限定，将"fmt.Print"视为带有需要引用的"."的非结构化文字。对于形式为`pkg.(*MyType).Meth`的方法名称，它甚至更强烈地反对。
4. 截至Go 1.11，默认情况下调试信息已压缩。旧版本的gdb，例如MacOS上默认提供的版本，无法理解该压缩格式。您可以使用`go build -ldflags=-compressdwarf=false`来生成未压缩的调试信息。(为方便起见，可以将-ldflags选项放入[GOFLAGS环境变量](https://go.dev/cmd/go/#hdr-Environment_variables)中，以免每次都需要指定它。)

## 教程

​	在本教程中，我们将检查[regexp](https://go.dev/pkg/regexp/)包的单元测试的二进制文件。要构建二进制文件，请切换到`$GOROOT/src/regexp`并运行`go test -c`。这应该会生成一个名为`regexp.test`的可执行文件。

### 入门

​	启动 GDB，调试 regexp.test：

```
$ gdb regexp.test
GNU gdb (GDB) 7.2-gg8
Copyright (C) 2010 Free Software Foundation, Inc.
License GPLv  3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
Type "show copying" and "show warranty" for licensing/warranty details.
This GDB was configured as "x86_64-linux".

Reading symbols from  /home/user/go/src/regexp/regexp.test...
done.
Loading Go Runtime support.
(gdb)
```

​	"Loading Go Runtime support" 消息表示 GDB 从 `$GOROOT/src/runtime/runtime-gdb.py` 加载了扩展。

​	为了帮助 GDB 找到 Go 运行时源和相关的支持脚本，请使用 '-d' 标志传递 `$GOROOT`：

```
$ gdb regexp.test -d $GOROOT
```

​	如果由于某种原因GDB仍然无法找到该目录或该脚本，则可以手动加载它，告诉gdb(假设您的go源代码位于`~/go/`)：

```
(gdb) source ~/go/src/runtime/runtime-gdb.py
Loading Go Runtime support.
```

### 检查源代码

​	使用 "l" 或 "list" 命令检查源代码。

```
(gdb) l
```

​	列出特定部分的源代码，使用函数名参数化 "list"(它必须与其包名一起限定)。

```
(gdb) l main.main
```

​	列出特定文件和行号：

```
(gdb) l regexp.go:1
(gdb) # Hit enter to repeat last command. Here, this lists next 10 lines.
```

### 命名

​	变量和函数名称必须限定为它们所属的包的名称。regexp 包的 Compile 函数被 GDB 认为是 'regexp.Compile'。

​	方法必须带上其接收器类型的名称。例如，`*Regexp` 类型的 String 方法在 GDB 中被称为 `'regexp.(*Regexp).String'`。

​	变量遮盖其他变量时，在调试信息中会自动添加数字后缀。闭包引用的变量会在指针前自动添加'&'。

### 设置断点

​	在 TestFind 函数处设置断点：

```
(gdb) b 'regexp.TestFind'
Breakpoint 1 at 0x424908: file /home/user/go/src/regexp/find_test.go, line 148.
```

​	运行程序：

```
(gdb) run
Starting program: /home/user/go/src/regexp/regexp.test

Breakpoint 1, regexp.TestFind (t=0xf8404a89c0) at /home/user/go/src/regexp/find_test.go:148
148	func TestFind(t *testing.T) {
```

​	程序已在断点处暂停。查看正在运行的 goroutine，以及它们正在做什么：

```
(gdb) info goroutines
  1  waiting runtime.gosched
* 13  running runtime.goexit
```

标有`*`的是当前 goroutine。

### 检查栈

​	查看我们暂停程序时的栈跟踪：

```
(gdb) bt  # backtrace
#0  regexp.TestFind (t=0xf8404a89c0) at /home/user/go/src/regexp/find_test.go:148
#1  0x000000000042f60b in testing.tRunner (t=0xf8404a89c0, test=0x573720) at /home/user/go/src/testing/testing.go:156
#2  0x000000000040df64 in runtime.initdone () at /home/user/go/src/runtime/proc.c:242
#3  0x000000f8404a89c0 in ?? ()
#4  0x0000000000573720 in ?? ()
#5  0x0000000000000000 in ?? ()
```

​	其他 goroutine(编号1)被阻塞在 `runtime.gosched` 上的通道接收操作中：

```
(gdb) goroutine 1 bt
#0  0x000000000040facb in runtime.gosched () at /home/user/go/src/runtime/proc.c:873
#1  0x00000000004031c9 in runtime.chanrecv (c=void, ep=void, selected=void, received=void)
 at  /home/user/go/src/runtime/chan.c:342
#2  0x0000000000403299 in runtime.chanrecv1 (t=void, c=void) at/home/user/go/src/runtime/chan.c:423
#3  0x000000000043075b in testing.RunTests (matchString={void (struct string, struct string, bool *, error *)}
 0x7ffff7f9ef60, tests=  []testing.InternalTest = {...}) at /home/user/go/src/testing/testing.go:201
#4  0x00000000004302b1 in testing.Main (matchString={void (struct string, struct string, bool *, error *)}
 0x7ffff7f9ef80, tests= []testing.InternalTest = {...}, benchmarks= []testing.InternalBenchmark = {...})
at /home/user/go/src/testing/testing.go:168
#5  0x0000000000400dc1 in main.main () at /home/user/go/src/regexp/_testmain.go:98
#6  0x00000000004022e7 in runtime.mainstart () at /home/user/go/src/runtime/amd64/asm.s:78
#7  0x000000000040ea6f in runtime.initdone () at /home/user/go/src/runtime/proc.c:243
#8  0x0000000000000000 in ?? ()
```

​	该栈帧显示我们当前正在执行regexp.TestFind函数，正如我们所期望的那样。

```
(gdb) info frame
Stack level 0, frame at 0x7ffff7f9ff88:
 rip = 0x425530 in regexp.TestFind (/home/user/go/src/regexp/find_test.go:148);
    saved rip 0x430233
 called by frame at 0x7ffff7f9ffa8
 source language minimal.
 Arglist at 0x7ffff7f9ff78, args: t=0xf840688b60
 Locals at 0x7ffff7f9ff78, Previous frame's sp is 0x7ffff7f9ff88
 Saved registers:
  rip at 0x7ffff7f9ff80
```

​	命令`info locals`列出了函数本地变量及其值，但使用时有些危险，因为它还会尝试打印未初始化的变量。未初始化的切片可能会导致gdb尝试打印任意大的数组。

​	该函数的参数：

```
(gdb) info args
t = 0xf840688b60
```

​	在打印参数时，请注意它是指向`Regexp`值的指针。请注意，GDB已将*错误地放在类型名称的右侧，并使用传统C样式创造了一个'struct'关键字。

```
(gdb) p re
(gdb) p t
$1 = (struct testing.T *) 0xf840688b60
(gdb) p t
$1 = (struct testing.T *) 0xf840688b60
(gdb) p *t
$2 = {errors = "", failed = false, ch = 0xf8406f5690}
(gdb) p *t->ch
$3 = struct hchan<*testing.T>
```

​	该`struct hchan<*testing.T>`是通道的运行时内部表示形式。 它当前为空，否则gdb会美观地打印其内容。

​	向前迈进：

```
(gdb) n  # execute next line
149             for _, test := range findTests {
(gdb)    # enter is repeat
150                     re := MustCompile(test.pat)
(gdb) p test.pat
$4 = ""
(gdb) p re
$5 = (struct regexp.Regexp *) 0xf84068d070
(gdb) p *re
$6 = {expr = "", prog = 0xf840688b80, prefix = "", prefixBytes =  []uint8, prefixComplete = true,
  prefixRune = 0, cond = 0 '\000', numSubexp = 0, longest = false, mu = {state = 0, sema = 0},
  machine =  []*regexp.machine}
(gdb) p *re->prog
$7 = {Inst =  []regexp/syntax.Inst = {{Op = 5 '\005', Out = 0, Arg = 0, Rune =  []int}, {Op =
    6 '\006', Out = 2, Arg = 0, Rune =  []int}, {Op = 4 '\004', Out = 0, Arg = 0, Rune =  []int}},
  Start = 1, NumCap = 2}
```

​	我们可以使用"s"进入String函数调用：

```
(gdb) s
regexp.(*Regexp).String (re=0xf84068d070, noname=void) at /home/user/go/src/regexp/regexp.go:97
97      func (re *Regexp) String() string {
```

​	获取栈跟踪以查看我们所在的位置：

```
(gdb) bt
#0  regexp.(*Regexp).String (re=0xf84068d070, noname=void)
    at /home/user/go/src/regexp/regexp.go:97
#1  0x0000000000425615 in regexp.TestFind (t=0xf840688b60)
    at /home/user/go/src/regexp/find_test.go:151
#2  0x0000000000430233 in testing.tRunner (t=0xf840688b60, test=0x5747b8)
    at /home/user/go/src/testing/testing.go:156
#3  0x000000000040ea6f in runtime.initdone () at /home/user/go/src/runtime/proc.c:243
....
```

查看源代码：

```
(gdb) l
92              mu      sync.Mutex
93              machine []*machine
94      }
95
96      // String returns the source text used to compile the regular expression.
97      func (re *Regexp) String() string {
98              return re.expr
99      }
100
101     // Compile parses a regular expression and returns, if successful,
```

### 漂亮的打印

​	GDB的漂亮打印机制是通过正则表达式匹配类型名称触发的。一个slice的例子：

```
(gdb) p utf
$22 =  []uint8 = {0 '\000', 0 '\000', 0 '\000', 0 '\000'}
```

​	由于slice、数组和字符串不是C指针，GDB不能为您解释下标操作，但您可以查看运行时表示来完成这个操作(在这里使用tab键自动完成可以帮助您)：

```
(gdb) p slc
$11 =  []int = {0, 0}
(gdb) p slc-><TAB>
array  slc    len
(gdb) p slc->array
$12 = (int *) 0xf84057af00
(gdb) p slc->array[1]
$13 = 0
```

扩展函数`$len`和`$cap`适用于字符串、数组和slice：

```
(gdb) p $len(utf)
$23 = 4
(gdb) p $cap(utf)
$24 = 4
```

​	通道和映射是'reference'类型，gdb将它们显示为指向C++类似类型`hash<int,string>*`的指针。取消引用将触发漂亮的打印

​	接口在运行时表示为类型描述符指针和值指针的组合。Go GDB运行时扩展将对其进行解码并自动触发运行时类型的漂亮打印。扩展函数`$dtype`会对动态类型进行解码(示例取自`regexp.go`第293行处的断点)。

```
(gdb) p i
$4 = {str = "cbb"}
(gdb) whatis i
type = regexp.input
(gdb) p $dtype(i)
$26 = (struct regexp.inputBytes *) 0xf8400b4930
(gdb) iface i
regexp.input: struct regexp.inputBytes *
```