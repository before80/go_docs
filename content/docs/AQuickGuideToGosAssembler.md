+++
title = "Go 汇编器速成指南"
weight = 28
date = 2023-05-18T17:31:23+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# A Quick Guide to Go's Assembler - Go汇编器速成指南

> 原文：[https://go.dev/doc/asm](https://go.dev/doc/asm)

## Go汇编器速成指南

​	本文档是gc Go编译器使用的不寻常汇编语言的快速概述。本文档并不全面。

​	汇编器基于Plan 9汇编程序的输入样式，该样式在[其他地方](https://9p.io/sys/doc/asm.html)有详细记录。如果您打算编写汇编语言，您应该阅读该文档，尽管其中大部分是针对Plan 9的。本文档提供了语法概述和与该文档中解释的内容的区别，并描述了在编写用于与Go交互的汇编代码时适用的特殊规定。

​	关于Go汇编器最重要的事情是它不是底层机器的直接表示。其中一些细节精确映射到机器，但有些则不是。这是因为编译器套件(请参见[此描述](https://9p.io/sys/doc/compiler.html))在常规管道中不需要汇编器传递。相反，编译器使用一种半抽象指令集进行操作，指令选择部分在代码生成后进行。汇编器基于半抽象形式工作，因此当您看到像`MOV`这样的指令时，工具链实际生成的操作可能根本不是移动指令，而可能是清除或加载。或者它可能与具有该名称的机器指令完全对应。通常，特定于机器的操作往往以它们自己的形式出现，而诸如内存移动和子例程调用和返回之类的更一般的概念更抽象。细节因架构而异，我们为不精确而道歉；情况并未得到明确定义。

​	汇编程序是解析半抽象指令集描述并将其转换为输入到链接器的指令的一种方法。如果要查看给定体系结构(例如amd64)的指令在汇编中的样子，标准库的源代码中有许多示例，例如[runtime](https://go.dev/pkg/runtime/)和[math/big](math/big)包。您还可以查看编译器生成的汇编代码(实际输出可能与此处所见不同)：

```
$ cat x.go
package main

func main() {
	println(3)
}
$ GOOS=linux GOARCH=amd64 go tool compile -S x.go        # or: go build -gcflags -S x.go
"".main STEXT size=74 args=0x0 locals=0x10
	0x0000 00000 (x.go:3)	TEXT	"".main(SB), $16-0
	0x0000 00000 (x.go:3)	MOVQ	(TLS), CX
	0x0009 00009 (x.go:3)	CMPQ	SP, 16(CX)
	0x000d 00013 (x.go:3)	JLS	67
	0x000f 00015 (x.go:3)	SUBQ	$16, SP
	0x0013 00019 (x.go:3)	MOVQ	BP, 8(SP)
	0x0018 00024 (x.go:3)	LEAQ	8(SP), BP
	0x001d 00029 (x.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (x.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (x.go:3)	FUNCDATA	$2, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (x.go:4)	PCDATA	$0, $0
	0x001d 00029 (x.go:4)	PCDATA	$1, $0
	0x001d 00029 (x.go:4)	CALL	runtime.printlock(SB)
	0x0022 00034 (x.go:4)	MOVQ	$3, (SP)
	0x002a 00042 (x.go:4)	CALL	runtime.printint(SB)
	0x002f 00047 (x.go:4)	CALL	runtime.printnl(SB)
	0x0034 00052 (x.go:4)	CALL	runtime.printunlock(SB)
	0x0039 00057 (x.go:5)	MOVQ	8(SP), BP
	0x003e 00062 (x.go:5)	ADDQ	$16, SP
	0x0042 00066 (x.go:5)	RET
	0x0043 00067 (x.go:5)	NOP
	0x0043 00067 (x.go:3)	PCDATA	$1, $-1
	0x0043 00067 (x.go:3)	PCDATA	$0, $-1
	0x0043 00067 (x.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x0048 00072 (x.go:3)	JMP	0
...
```

​	`FUNCDATA` 和 `PCDATA` 指令包含垃圾回收器使用的信息，由编译器引入。

​	要查看链接后的二进制文件中放入了什么，请使用 `go tool objdump`：

```
$ go build -o x.exe x.go
$ go tool objdump -s main.main x.exe
TEXT main.main(SB) /tmp/x.go
  x.go:3		0x10501c0		65488b0c2530000000	MOVQ GS:0x30, CX
  x.go:3		0x10501c9		483b6110		CMPQ 0x10(CX), SP
  x.go:3		0x10501cd		7634			JBE 0x1050203
  x.go:3		0x10501cf		4883ec10		SUBQ $0x10, SP
  x.go:3		0x10501d3		48896c2408		MOVQ BP, 0x8(SP)
  x.go:3		0x10501d8		488d6c2408		LEAQ 0x8(SP), BP
  x.go:4		0x10501dd		e86e45fdff		CALL runtime.printlock(SB)
  x.go:4		0x10501e2		48c7042403000000	MOVQ $0x3, 0(SP)
  x.go:4		0x10501ea		e8e14cfdff		CALL runtime.printint(SB)
  x.go:4		0x10501ef		e8ec47fdff		CALL runtime.printnl(SB)
  x.go:4		0x10501f4		e8d745fdff		CALL runtime.printunlock(SB)
  x.go:5		0x10501f9		488b6c2408		MOVQ 0x8(SP), BP
  x.go:5		0x10501fe		4883c410		ADDQ $0x10, SP
  x.go:5		0x1050202		c3			RET
  x.go:3		0x1050203		e83882ffff		CALL runtime.morestack_noctxt(SB)
  x.go:3		0x1050208		ebb6			JMP main.main(SB)
```

### 常量

Although the assembler takes its guidance from the Plan 9 assemblers, it is a distinct program, so there are some differences. One is in constant evaluation. Constant expressions in the assembler are parsed using Go's operator precedence, not the C-like precedence of the original. Thus `3&1<<2` is 4, not 0—it parses as `(3&1)<<2` not `3&(1<<2)`. Also, constants are always evaluated as 64-bit unsigned integers. Thus `-2` is not the integer value minus two, but the unsigned 64-bit integer with the same bit pattern. The distinction rarely matters but to avoid ambiguity, division or right shift where the right operand's high bit is set is rejected.

​	尽管汇编程序遵循 Plan 9 汇编程序的指导方针，但它是一个独立的程序，因此有一些差异。一个差异在于常量的求值。汇编程序中的常量表达式使用 Go 的运算符优先级进行解析，而不是原始程序的类 C 优先级。因此，`3&1<<2` 是 4，而不是 0 —— 它解析为 `(3&1)<<2` 而不是 `3&(1<<2)`。此外，常量总是作为 64 位无符号整数计算。因此，-2不是整数值减二，而是具有相同位模式的64位无符号整数。这种区别很少有影响，但为避免歧义，当右操作数的高位设置时，除法或右移将被拒绝。

### 符号

​	一些符号，例如 `R1` 或 `LR`，是预定义的并且指向寄存器。确切的集取决于体系结构。

​	有四个预定义的符号，它们指向伪寄存器。这些不是实际的寄存器，而是由工具链维护的虚拟寄存器，例如帧指针。伪寄存器集在所有架构中都相同：

- `FP`: 帧指针：参数和局部变量。
- `PC`: 程序计数器：跳转和分支。
- `SB`: 静态基地址指针：全局符号。
- `SP`: 栈指针：本地栈帧中的最高地址。

​	所有用户定义的符号都写成相对于伪寄存器 `FP`(参数和局部变量)和 `SB`(全局变量)的偏移量。

​	可以将伪寄存器 `SB` 视为内存的起点，因此符号 `foo(SB)` 是 `foo` 的内存地址。这种形式用于命名全局函数和数据。在名称中添加 `<>`，例如 `foo<>(SB)`，使名称仅在当前源文件中可见，就像 C 文件中的顶级静态声明。在名称后添加偏移量将引用该偏移量距符号地址的距离，因此 `foo+4(SB)` 是从 `foo` 开始的四个字节之后。

​	FP 伪寄存器是用于引用函数参数的虚拟帧指针。编译器维护一个虚拟帧指针，并将堆栈上的参数作为该伪寄存器的偏移量引用。因此，`0(FP)`是函数的第一个参数，`8(FP)`是第二个参数(在 64 位机器上)，以此类推。但是，在这种方式引用函数参数时，必须在开头放置一个名称，例如 `first_arg+0(FP)`和 `second_arg+8(FP)`。这里的偏移量的含义——从帧指针的偏移量——与在 `SB` 中使用的含义不同，SB 中的偏移量是从符号的偏移量。汇编程序执行这个约定，拒绝使用纯 `0(FP)`和 `8(FP)`。实际名称在语义上是无关紧要的，但应该用于记录参数的名称。值得强调的是，`FP` 总是伪寄存器，即使在具有硬件帧指针的体系结构中也是如此。

​	对于带有 Go 原型的汇编函数，go vet 将检查参数名称和偏移量是否匹配。在 32 位系统上，64 位值的低 32 位和高 32 位通过向名称添加 `_lo` 或 `_hi` 后缀来区分，例如 `arg_lo+0(FP)`或 `arg_hi+4(FP)`。如果 Go 原型没有为其结果命名，则预期的汇编名称为 ret。

​	`SP` 伪寄存器是用于引用框架本地变量和准备用于函数调用的参数的虚拟堆栈指针。它指向局部堆栈帧内的最高地址，因此引用应使用范围为 [−framesize,0) 的负偏移量：`x-8(SP)`、`y-4(SP)`等。

​	在具有名为 `SP` 的硬件寄存器的体系结构中，名称前缀区分对虚拟堆栈指针和对架构 `SP` 寄存器的引用。也就是说，`x-8(SP)`和`-8(SP)`是不同的内存位置：第一个是对虚拟堆栈指针伪寄存器的引用，而第二个是对硬件 `SP` 寄存器的引用。

​	在 SP 和 PC 通常是物理编号寄存器的体系结构上，在 Go 汇编器中，SP 和 PC 名称仍然被特殊处理；例如，对 SP 的引用需要一个符号，就像 FP 一样。要访问实际的硬件寄存器，请使用真正的 R 名称。例如，在 ARM 体系结构上，硬件 SP 和 PC 可以访问为 R13 和 R15。

​	分支和直接跳转总是写成PC的偏移量或标签跳转的形式：

```
label:
	MOVW $0, R1
	JMP label
```

​	每个标签只在定义它的函数中可见。因此，一个文件中的多个函数可以定义和使用相同的标签名称。直接跳转和调用指令可以指向文本符号，例如`name(SB)`，但不能指向符号的偏移量，例如`name+4(SB)`。

​	指令、寄存器和汇编指令总是大写的，以提醒您汇编编程是一项危险的任务。(例外：ARM上的`g`寄存器重命名。)

​	在Go对象文件和二进制文件中，符号的完整名称是包路径后跟句点和符号名称：`fmt.Printf`或`math/rand.Int`。因为汇编程序的解析器将句点和斜杠视为标点符号，所以这些字符串不能直接用作标识符名称。相反，汇编程序允许在标识符中使用中间点字符`U+00B7`和除号`U+2215`，并将其重写为普通的句点和斜杠。在汇编源文件中，上述符号被写成`fmt·Printf`和`math∕rand·Int`。编译器使用`-S`标志生成的汇编列表直接显示句点和斜杠，而不是汇编程序需要的Unicode替换。

​	大多数手写汇编文件的符号名称中不包括完整的包路径，因为链接器会在以句点开头的任何名称之前插入当前对象文件的包路径：在math/rand包实现的汇编源文件中，可以将该包的Int函数称为`·Int`。这种约定避免了在自己的源代码中硬编码包的导入路径的需要，使得将代码从一个位置移动到另一个位置更容易。

### 指令

​	汇编程序使用各种指令将文本和数据绑定到符号名称。例如，这里是一个简单的完整函数定义。TEXT指令声明符号`runtime·profileloop`，接下来的指令形成函数体。`TEXT`块中的最后一个指令必须是某种跳转，通常是`RET`(伪)指令。(如果不是，链接器将附加一个跳转到自身的指令；`TEXT`中没有fallthrough。)在符号之后，参数是标志(见下文)和帧大小，一个常量(但见下文)：

```
TEXT runtime·profileloop(SB),NOSPLIT,$8
	MOVQ	$runtime·profileloop1(SB), CX
	MOVQ	CX, 0(SP)
	CALL	runtime·externalthreadhandler(SB)
	RET
```

​	一般情况下，帧大小后面跟着一个参数大小，用减号分隔。这不是减法，只是特有的语法。例如，帧大小 `$24-8` 表示该函数有一个 24 字节的帧，并使用 8 字节的参数调用，这些参数存在于调用者的帧中。如果未在 `TEXT` 中指定 `NOSPLIT`，则必须提供参数大小。对于具有 Go 原型的汇编函数，`go vet` 将检查参数大小是否正确。

​	请注意，符号名称使用中间点分隔组件，并指定为从静态基址伪寄存器 `SB` 的偏移量。在 Go 源文件中，这个函数会使用简单的名称 `profileloop` 被称为 `runtime` 包。

​	全局数据符号由一系列初始化的 `DATA` 指令后跟一个 `GLOBL` 指令定义。每个 `DATA` 指令都初始化相应内存的一个部分。未显式初始化的内存被清零。`DATA` 指令的一般形式是

```
DATA	symbol+offset(SB)/width, value
```

​	它使用给定的值初始化给定偏移量和宽度的符号内存。给定符号的 `DATA` 指令必须按偏移量递增的顺序编写。

​	`GLOBL` 指令声明一个符号为全局。参数是可选标志和声明为全局的数据大小，除非 `DATA` 指令已将其初始化，否则初始值全为零。`GLOBL` 指令必须跟随任何相应的 `DATA` 指令。

例如，

```
DATA divtab<>+0x00(SB)/4, $0xf4f8fcff
DATA divtab<>+0x04(SB)/4, $0xe6eaedf0
...
DATA divtab<>+0x3c(SB)/4, $0x81828384
GLOBL divtab<>(SB), RODATA, $64

GLOBL runtime·tlsoffset(SB), NOPTR, $4
```

声明并初始化 `divtab<>`，一个只读的 64 字节表，其中每个元素为 4 字节整数值，并声明 `runtime·tlsoffset`，一个隐式清零、不包含指针的 4 字节变量。

​	指令可能有一或两个参数。如果有两个参数，第一个参数是标志位掩码，可以写成数值表达式，加或或或在一起，或者可以使用符号设置，以便更容易被人吸收。它们的值，在标准 `#include` 文件 `textflag.h` 中定义，包括：

- `NOPROF` = 1
  (对于 `TEXT` 项)不要对标记的函数进行性能分析。该标志已被弃用。
  
- `DUPOK` = 2
  在单个二进制文件中可以有多个该符号的实例。链接器将选择其中一个副本进行使用。
  
- `NOSPLIT` = 4
  (对于`TEXT`项。)不要插入开头检查堆栈是否必须分裂的前导部分。例程的帧，以及它调用的任何东西，都必须适合当前栈段中剩余的备用空间。用于保护例程，例如栈分裂代码本身。
  
- `RODATA` = 8
  (对于`DATA`和`GLOBL`项。)将这些数据放入只读部分。

- `NOPTR` = 16
  (对于`DATA`和`GLOBL`项。)此数据不包含指针，因此不需要被垃圾收集器扫描。

- `WRAPPER` = 32

  (对于`TEXT`项。)这是一个包装函数，不应视为禁用`recover(恢复)`。

- `NEEDCTXT` = 64

  (对于`TEXT`项。)此函数是闭包，因此使用其传入的上下文寄存器。

- `LOCAL` = 128
  此符号局部于动态共享对象。
  
- `TLSBSS` = 256
  
  (对于`DATA`和`GLOBL`项。)将这些数据放入线程局部存储中。
  
- `NOFRAME` = 512

  (对于`TEXT`项。)不插入指令以分配栈帧并保存/恢复返回地址，即使这不是叶子函数也是如此。仅适用于声明帧大小为0的函数。

- `TOPFRAME` = 2048

  (对于`TEXT`项。)函数是调用栈的最外层帧。回溯应在此函数中停止。

### 与Go类型和常量交互

​	如果一个包有任何 `.s` 文件，那么 `go build` 将指导编译器发出一个名为 `go_asm.h` 的特殊头文件，`.s` 文件可以将其 `#include`。该文件包含符号 `#define` 常量，用于 Go 结构字段的偏移量、Go 结构类型的大小以及当前包中定义的大多数 Go `const` 声明。Go 汇编应避免对 Go 类型布局做出假设，而应使用这些常量。这提高了汇编代码的可读性，并使其能够适应 Go 类型定义或 Go 编译器使用的布局规则的数据布局变化。

​	常量的形式为 `const_*name*`。例如，给定 Go 声明 `const bufSize = 1024`，汇编代码可以将这个常量的值称为 `const_bufSize`。

​	字段偏移量的形式为 `*type*_*field*`。结构体大小的形式为 `*type*__size`。例如，考虑以下 Go 定义：

``` go
type reader struct {
	buf [bufSize]byte
	r   int
}
```

​	汇编可以将此结构体的大小称为 `reader__size`，两个字段的偏移量分别称为 `reader_buf` 和 `reader_r`。因此，如果寄存器 `R1` 包含一个指向 `reader` 的指针，汇编可以将 r 字段引用为 `reader_r(R1)`。

​	如果这些 `#define` 名称中有任何歧义(例如一个带有 `_size` 字段的结构体)，`#include "go_asm.h"` 将失败并显示"redefinition of macro(重定义宏)"的错误。

### 运行时协调

​	为了使垃圾回收正常运行，运行时必须知道所有全局数据和大多数堆栈帧中指针的位置。Go编译器在编译Go源文件时发出此信息，但汇编程序必须显式定义它。

​	用`NOPTR`标记的数据符号(见上文)被视为不包含指向运行时分配的数据的指针。用`RODATA`标记的数据符号被分配在只读内存中，因此被视为隐式标记为`NOPTR`。如果一个数据符号的总大小小于指针大小，它也被视为隐式标记为`NOPTR`。在汇编源文件中定义包含指针的符号是不可能的，这样的符号必须在Go源文件中定义。即使没有`DATA`和`GLOBL`指令，汇编源代码仍可以通过名称引用符号。一个好的通用经验法则是将所有非`RODATA`符号定义在Go中而不是汇编中。

​	每个函数还需要注释其参数、结果和本地栈帧中活动指针的位置。对于没有指针结果和没有本地栈帧或没有函数调用的汇编函数，唯一的要求是在同一包中的Go源文件中为该函数定义一个Go原型。汇编函数的名称不得包含包名称组件(例如，syscall包中的`Syscall`函数应该使用名称`·Syscall`而不是等效名称`syscall·Syscall`在其`TEXT`指令中)。对于更复杂的情况，需要显式注释。这些注释使用标准的`#include`文件`funcdata.h`中定义的伪指令。

​	如果一个函数没有参数和结果，指针信息可以省略。这可以通过`TEXT`指令上的参数大小注释`$*n*-0`来表示。否则，必须在Go源文件中为该函数提供指针信息，即使是从未直接从Go调用的汇编函数也是如此。 (原型也将使`go vet`检查参数引用。)在函数开始时，假定参数已初始化，但结果未初始化。如果结果在调用指令期间保存活动指针，则函数应从零开始并执行伪指令`GO_RESULTS_INITIALIZED`。此指令记录结果现在已初始化，应在堆栈移动和垃圾回收期间进行扫描。通常更容易安排汇编函数不返回指针或不包含调用指令；标准库中没有汇编函数使用`GO_RESULTS_INITIALIZED`。

​	如果一个函数没有本地栈帧，则可以省略指针信息。这可以通过在`TEXT`指令上使用局部栈帧大小注释`$0-n`来表示。如果函数不包含函数调用指令，则也可以省略指针信息。否则，本地栈框架不能包含指针，并且汇编必须通过执行伪指令`NO_LOCAL_POINTERS`来确认这一事实。由于堆栈调整是通过移动堆栈来实现的，在任何函数调用期间堆栈指针可能会发生变化：即使是对堆栈数据的指针也不应保留在局部变量中。

​	汇编函数应始终提供Go原型，既为了为参数和结果提供指针信息，也为了让`go vet`检查访问它们所使用的偏移量是否正确。

## Architecture-specific details 架构特定的细节

​	列出每个机器的所有指令和其他细节是不切实际的。要查看给定机器(例如ARM)定义了哪些指令，请查看该体系结构的`obj`支持库的源代码，位于目录`src/cmd/internal/obj/arm`中。在该目录中有一个名为`a.out.go`的文件；它包含一长串以`A`开头的常量，如下所示：

``` go
const (
	AAND = obj.ABaseARM + obj.A_ARCHSPECIFIC + iota
	AEOR
	ASUB
	ARSB
	AADD
	...
```

​	这是每种体系结构的汇编器和链接器所知道的指令及其拼写的列表。在此列表中，每个指令以大写字母`A`开头，因此`AAND`代表按位与指令，`AND`(没有前导的`A`)在汇编源中的书写方式为AND。枚举大多按字母顺序排列。(独立于体系结构的`AXXX`，定义在`cmd/internal/obj`包中，代表一个无效指令)。`A`名称的顺序与实际机器指令的编码无关。`cmd/internal/obj`包负责处理该细节。

​	386和AMD64架构的指令列在`cmd/internal/obj/x86/a.out.go`中。

​	这些架构共享常见寻址模式的语法，例如`(R1)`(寄存器间接)、`4(R1)`(带偏移的寄存器间接)和`$foo(SB)`(绝对地址)。汇编器还支持每个架构特定的一些(不一定是所有)寻址模式。下面的部分列出了这些模式。

​	从前面几节的示例中可以看到的一个细节是指令中的数据从左到右流动：`MOVQ $0`，`CX`清除了`CX`。即使在传统符号使用相反方向的体系结构上，也适用此规则。

​	以下是受支持架构的一些关键Go特定细节的描述。

### 32位Intel 386

​	运行时指向`g`结构的指针通过MMU中另外未使用的(就Go而言)寄存器的值来维护。在runtime包中，汇编代码可以包含`go_tls.h`，它定义了一个特定于操作系统和体系结构的宏`get_tls`，用于访问此寄存器。`get_tls`宏接受一个参数，即将`g`指针加载到其中的寄存器。

​	例如，使用`CX`加载`g`和`m`的序列如下：

```
#include "go_tls.h"
#include "go_asm.h"
...
get_tls(CX)
MOVL	g(CX), AX     // Move g into AX.
MOVL	g_m(AX), BX   // Move g.m into BX.
```

​	`get_tls`宏也被定义在[amd64](https://go.dev/doc/asm#amd64)上。

​	寻址模式：

- `(DI)(BX*2)`: 地址为`DI`加上`BX*2`处的位置。
- `64(DI)(BX*2)`: 地址为`DI`加上`BX*2`加上64处的位置。这些模式仅接受1、2、4和8作为比例因子。

​	当使用编译器和汇编器的`-dynlink`或`-shared`模式时，对于任何固定的内存位置，例如全局变量的任何读取或写入操作，都必须假定会覆盖`CX`。因此，为了在这些模式下安全使用，汇编源代码通常应避免使用`CX`，除非在内存引用之间。

### 64位Intel 386(又称amd64)

​	这两种架构在汇编器层面上基本相同。在64位版本上访问`m`和`g`指针的汇编代码与32位386上相同，只是它使用`MOVQ`而不是`MOVL`：

```
get_tls(CX)
MOVQ	g(CX), AX     // Move g into AX.
MOVQ	g_m(AX), BX   // Move g.m into BX.
```

​	寄存器`BP`是被调用者保存的。当帧大小大于零时，汇编器会自动插入`BP`保存/恢复。允许将`BP`用作通用寄存器，但它可能会干扰基于采样的分析。

### ARM

​	寄存器`R10`和`R11`由编译器和链接器保留。

​	`R10`指向`g`(goroutine)结构。在汇编源代码中，必须将此指针称为`g`；不识别`R10`名称。

​	为了让人们和编译器更容易编写汇编代码，ARM链接器允许使用一般寻址形式和伪操作，例如`DIV`或`MOD`，这些操作可能无法使用单个硬件指令表达。它将这些形式实现为多个指令，通常使用`R11`寄存器来保存临时值。手写汇编可以使用`R11`，但这要求确定链接器未使用它来实现函数中的任何其他指令。

​	当定义`TEXT`时，指定帧大小`$-4`告诉链接器，这是一个不需要在入口保存`LR`的叶函数。

​	名称`SP`始终指代前面描述的虚拟栈指针。对于硬件寄存器，请使用`R13`。

​	条件代码语法是将一个点和一个或两个字母代码附加到指令上，例如`MOVW.EQ`。可以附加多个代码修饰符：`MOVM.IA.W`。代码修饰符的顺序是不相关的。

​	寻址模式：

- `R0->16`
  `R0>>16`
  `R0<<16`
  `R0@>16`:  对于`<<`，将`R0`左移16位。其他代码是`->`(算术右移)，`>>`(逻辑右移)和`@`>(向右旋转)。
- `R0->R1`
  `R0>>R1`
  `R0<<R1`
  `R0@>R1`:  对于`<<`，将`R0`左移`R1`次数。其他代码是`->`(算术右移)，`>>`(逻辑右移)和`@>`(向右旋转)。
- `[R0,g,R12-R15]`:  对于多寄存器指令，这是由`R0`、`g`和`R12`到`R15`共同组成的集合。
- `(R5, R6)`:  目标寄存器对

### ARM64

​	`R18` 是"platform register(平台寄存器)"，在 Apple 平台上被保留。为了防止意外误用，该寄存器被命名为 `R18_PLATFORM`。`R27` 和 `R28` 被编译器和链接器保留。`R29` 是帧指针，`R30` 是链接寄存器。

​	指令修饰符附加在句点后面的指令中。唯一的修饰符是 `P`(后自增)和 `W`(前自增)：`MOVW.P`，`MOVW.W`

​	寻址模式：

- `R0->16`
  `R0>>16`
  `R0<<16`
  `R0@>16`: 这些与 32 位 ARM 上相同。
- `$(8<<12)`: 将立即数 8 左移 12 位。
- `8(R0)`: 将 R0 的值和 8 相加。
- `(R2)(R0)`: 位于 `R0` 和 `R2` 之和的位置。
- `R0.UXTB`
  `R0.UXTB<<imm`:  `UXTB`：从 `R0` 的低位中提取一个 8 位值，并将其零扩展到 `R0` 的大小。`R0.UXTB<<imm`：将 `R0.UXTB` 的结果左移 `imm` 位。`imm` 的值可以是 0、1、2、3 或 4。其他扩展包括 `UXTH`(16 位)、`UXTW`(32 位)和 `UXTX`(64 位)
- `R0.SXTB`
  `R0.SXTB<<imm`:  `SXTB`：从 `R0` 的低位中提取一个 8 位值，并将其符号扩展到 `R0` 的大小。`R0.SXTB<<imm`：将 `R0.SXTB` 的结果左移 imm 位。imm 的值可以是 0、1、2、3 或 4。其他扩展包括 `SXTH`(16 位)、`SXTW`(32 位)和 `SXTX`(64 位)。
- `(R5, R6)`:  LDAXP/LDP/LDXP/STLXP/STP/STP 的寄存器对。

​	参考：[Go ARM64 汇编指令参考手册](https://go.dev/pkg/cmd/internal/obj/arm64)

### PPC64

​	参考: [Go PPC64 汇编指令参考手册](https://go.dev/pkg/cmd/internal/obj/ppc64)

### IBM z/Architecture，也称为s390x

​	寄存器`R10`和`R11`是保留的。汇编器在汇编某些指令时会使用它们来保存临时值。

​	寄存器`R13`指向`g`(goroutine)结构。在汇编源代码中必须将此寄存器称为`g`；不认可使用寄存器名`R13`。

​	寄存器`R15`指向栈帧，通常只应使用虚拟寄存器`SP`和`FP`访问。

​	加载和存储多个指令对一组寄存器进行操作。寄存器范围由起始寄存器和结束寄存器指定。例如，`LMG(R9)`，`R5`，`R7`会将64位值`0(R9)`，`8(R9)`和`16(R9)`处的值分别加载到`R5`、`R6`和`R7`中。

​	储存和存储指令，如MVC和XC，是以长度作为第一个参数写入的。例如，XC $8, (R9), (R9)将清除R9中指定地址的8个字节。

​	存储和存储指令(如`MVC`和`XC`)的长度作为第一个参数写入。例如，`XC $8,(R9),(R9)`将清除`R9`指定的地址处的8个字节。

​	如果向量指令需要一个长度或索引作为参数，则它将是第一个参数。例如，`VLEIF $1,$16,V2`将值16加载到`V2`的索引1中。在使用向量指令时，应注意确保它们在运行时可用。要使用向量指令，机器必须具有向量功能(设施列表中的位129)和内核支持。如果没有内核支持，则向量指令将没有效果(它将等效于`NOP`指令)。

寻址模式：

- `(R5)(R6*1)`: `R5`加上`R6`的位置。与x86一样，这是一个缩放模式，但是唯一允许的缩放是`1`。

### MIPS, MIPS64

​	通用寄存器的名称为`R0`到`R31`，浮点寄存器的名称为`F0`到`F31`。

​	`R30`保留用于指向`g`，`R23`用作临时寄存器。

​	在`TEXT`指令中，MIPS的帧大小为`$-4`，MIPS64的帧大小为`$-8`，指示链接器不保存`LR`。

​	`SP`指向虚拟栈指针。对于硬件寄存器，请使用`R29`。

​	寻址模式：

- `16(R1)`: `R1`加上16的位置。
- `(R1)`:  等同于0(R1)。

​	通过预定义`GOMIPS_hardfloat`或`GOMIPS_softfloat`，`GOMIPS`环境变量的值(`hardfloat`或`softfloat`)可以在汇编代码中使用。

​	通过预定义`GOMIPS64_hardfloat`或`GOMIPS64_softfloat`，GOMIPS64环境变量的值(`hardfloat`或`softfloat`)可以在汇编代码中使用。

### 不支持的操作码

​	汇编器的设计是为了支持编译器，因此并非所有硬件指令都适用于所有架构：如果编译器不生成它，则可能不存在。如果您需要使用缺失的指令，有两种方法可供选择。一种方法是更新汇编器以支持该指令，这很简单，但只有在有可能再次使用该指令时才值得。相反，对于简单的一次性情况，可以在`TEXT`中使用`BYTE`和`WORD`指令将显式数据放置到指令流中。以下是386运行时如何定义64位原子加载函数。

```
// uint64 atomicload64(uint64 volatile* addr);
// 因此实际上是
// void atomicload64(uint64 *res, uint64 volatile *addr);
TEXT runtime·atomicload64(SB), NOSPLIT, $0-12
	MOVL	ptr+0(FP), AX
	TESTL	$7, AX
	JZ	2(PC)
	MOVL	0, AX // crash with nil ptr deref
	LEAL	ret_lo+4(FP), BX
	// MOVQ (%EAX), %MM0
	BYTE $0x0f; BYTE $0x6f; BYTE $0x00
	// MOVQ %MM0, 0(%EBX)
	BYTE $0x0f; BYTE $0x7f; BYTE $0x03
	// EMMS
	BYTE $0x0F; BYTE $0x77
	RET
```