+++
title = "Go 1.3+ Compiler Overhaul"
weight = 5
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# **Go 1.3+ Compiler Overhaul**

> 原文：[https://docs.google.com/document/d/1P3BLR31VA8cvLJLfMibSuTdwTuF7WWLux71CYD0eeD8/edit](https://docs.google.com/document/d/1P3BLR31VA8cvLJLfMibSuTdwTuF7WWLux71CYD0eeD8/edit)

golang.org/s/go13compiler

Russ Cox

December 2013

## Abstract 摘要

The Go compiler today is written in C. It is time to move to Go.

今天的Go编译器是用C语言编写的，现在是时候转向Go了。

[**Update**: This work was completed and presented at GopherCon. See “[Go from C to Go](https://www.google.com/url?q=https://www.youtube.com/watch?v%3DQIE5nV5fDwA&sa=D&source=editors&ust=1669903232902574&usg=AOvVaw10OaW5d1pEZdqpj20LZ1Dr)”.]

[更新：这项工作已经完成，并在GopherCon上发表。参见 "从C到Go"。］

## Background 背景

The “gc” Go toolchain is derived from the Plan 9 compiler toolchain. The assemblers, C compilers, and linkers are adopted essentially unchanged, and the Go compilers (in cmd/gc, cmd/5g, cmd/6g, and cmd/8g) are new C programs that fit into the toolchain.

gc "Go工具链是由Plan 9编译器工具链衍生而来。汇编器、C 编译器和链接器基本上没有改变，Go 编译器（在 cmd/gc、cmd/5g、cmd/6g 和 cmd/8g 中）是适合该工具链的新 C 程序。

Writing the compiler in C had some important advantages over using Go at the start of the project, most prominent among them the fact that, at first, Go did not exist and so could not be used to write a compiler, and the fact that, once Go did exist, it often changed in significant, backwards-incompatible ways. Using C instead of Go avoided both the initial and ongoing bootstrapping problems. Today, however, Go does exist, and its definition is stable as of Go 1, so the problems of bootstrapping are greatly reduced.

在项目开始时，用C语言编写编译器比使用Go有一些重要的优势，其中最突出的是，一开始，Go并不存在，所以不能用来编写编译器，而一旦Go存在，它往往会发生重大的、向后不兼容的变化。使用C语言而不是Go，可以避免最初和持续的引导问题。然而今天，Go确实存在，而且从Go 1开始它的定义就很稳定，所以引导的问题大大减少。

As the bootstrapping problems have receded, other engineering concerns have arisen that make Go much more attractive than C for the compiler implementation. The concerns include:

随着引导问题的消退，其他工程方面的问题也出现了，这使得Go在编译器实现方面比C更有吸引力。这些问题包括：

- It is easier to write correct Go code than to write correct C code.
- It is easier to debug incorrect Go code than to debug incorrect C code.
- Work on a Go compiler necessarily requires a good understanding of Go. Implementing the compiler in C adds an unnecessary second requirement.
- Go makes parallel execution trivial compared to C.
- Go has better standard support than C for modularity, for automated rewriting, for unit testing, and for profiling.
- Go is much more fun to use than C.
- 编写正确的Go代码比编写正确的C代码更容易。
  调试错误的Go代码比调试错误的C代码更容易。
  从事Go编译器的工作必然要求对Go有很好的理解。用C语言实现编译器会增加第二个不必要的要求。
  与C相比，Go使并行执行变得微不足道。
  在模块化、自动重写、单元测试和剖析方面，Go比C有更好的标准支持。
  Go在使用上比C语言更有趣。

For all these reasons, we believe it is time to switch to Go compilers written in Go.

基于所有这些原因，我们认为现在是转换到用Go语言编写的编译器的时候了。

## Proposed Plan 建议的计划

We plan to translate the existing compilers from C to Go by writing and then applying an automatic translator. The conversion will proceed in phases, starting in Go 1.3 but continuing into future releases.

我们计划通过编写和应用自动翻译器将现有的编译器从C语言翻译成Go语言。这一转换将分阶段进行，从Go 1.3开始，但会持续到未来的版本。

*Phase 1*. Develop and debug the translator. This can be done in parallel with ordinary development. In particular, it is fine for people to continue making changes to the C version of the compiler during this phase. The translator is a fair amount of work, but we are confident that we can build one that works for the specific case of translating the compilers. There are many corners of C that have no direct translation into Go; macros, unions, and bit fields are probably highest on the list. Fortunately (but not coincidentally), those features are rarely used, if at all, in the code being translated. Pointer arithmetic and arrays are also some work to translate, but even those are rare in the compiler, which primarily operates on trees and linked lists. The translator will preserve the comments and structure of the original C code, so the translation should be as readable as the current compiler.

第一阶段。开发和调试翻译器。这可以与普通开发同时进行。特别是，在这个阶段，人们可以继续对编译器的C版本进行修改。翻译器是一个相当大的工作量，但我们有信心，我们可以建立一个适用于翻译编译器这一特定情况的翻译器。C语言中的许多角落都没有直接翻译成Go语言；宏、联合体和位域可能是列表中最高的。幸运的是（但并非巧合），这些功能在被翻译的代码中很少使用，如果有的话。指针算术和数组也需要翻译，但即使是这些在编译器中也很少，因为编译器主要是在树和链接列表上操作。译者将保留原始C代码的注释和结构，所以译文应该和当前的编译器一样可读。

*Phase* *2*. Use the translator to convert the compilers from C to Go and delete the C copies. At this point we have transitioned to Go and still have a working compiler, but the compiler is still very much a C program. This *may* happen for Go 1.3, but that’s pretty aggressive. It is more likely to happen for Go 1.4.

第二阶段。使用翻译器将编译器从C语言转换为Go语言，并删除C语言副本。在这一点上，我们已经过渡到Go，并且仍然有一个可以工作的编译器，但编译器在很大程度上仍然是一个C程序。这可能会发生在Go 1.3上，但那是相当激进的。它更有可能发生在Go 1.4上。

*Phase 3.* Use some tools, perhaps derived from gofix and the Go oracle to split the compiler into packages, cleaning up and documenting the code, and adding unit tests as appropriate. This phase turns the compiler into an idiomatic Go program. This is targeted for Go 1.4.

第三阶段。使用一些工具，也许是从gofix和Go oracle派生出来的，把编译器拆成包，清理和记录代码，并适当地增加单元测试。这个阶段将编译器变成一个成语的Go程序。这个阶段的目标是Go 1.4。

*Phase 4a*. Apply standard profiling and measurement techniques to understand and optimize the memory and CPU usage of the compiler. This may include introducing parallelization; if so, the race detector is likely to be a significant help. This is targeted for Go 1.4, but parts may slip to Go 1.5. Some basic profiling and optimization may be done earlier, in Phase 3.

第4a阶段。应用标准的剖析和测量技术，了解并优化编译器的内存和CPU使用情况。这可能包括引入并行化；如果是这样，竞赛检测器可能会有很大帮助。这是为Go 1.4设计的，但部分内容可能滑向Go 1.5。一些基本的剖析和优化可以在第三阶段提前完成。

Phase 4b. (Concurrent with Phase 4a.) With the compiler split into packages with clearly defined boundaries, it should be straightforward to introduce a new middle representation between the architecture-independent unordered tree (Node*s) and the architecture-dependent ordered list (Prog*s) used today. That representation, which should be architecture-independent but contain information about precise order of execution, can be used to introduce order-dependent but architecture-independent optimizations like elimination of redundant nil checks and bounds checks. It may be based on SSA and if so would certainly take advantage of the lessons learned from Alan Donovan’s go.tools/ssa package.

第4b阶段。(与第4a阶段同时进行。)随着编译器被分割成具有明确边界的包，在与架构无关的无序树（Node*s）和目前使用的与架构有关的有序列表（Prog*s）之间引入一个新的中间表示法应该是很简单的。这种表示法应该是与架构无关的，但包含精确的执行顺序的信息，可以用来引入与顺序相关但与架构无关的优化，比如消除多余的nil检查和边界检查。它可能是基于SSA的，如果是这样的话，肯定会利用从Alan Donovan的go.tools/ssa包中获得的经验。

*Phase 5*. Replace the front end with the latest (perhaps new) versions of go/parser and go/types. Robert Griesemer has discussed the possibility of designing new go/parser and go/types APIs at some point, based on experience with the current ones (and under new names, to preserve Go 1 compatibility). The work of connecting them to a compiler back end may help guide design of new APIs.

第五阶段。用go/parser和go/types的最新（也许是新）版本替换前端。Robert Griesemer已经讨论了在某个时候设计新的go/parser和go/types API的可能性，这些API是基于目前的经验（并且用新的名字，以保持Go 1的兼容性）。将它们与编译器后端连接的工作可能有助于指导新的API的设计。

## Bootstrapping 引导

With a Go compiler written in Go, there must be a plan for bootstrapping from scratch. The rule we plan to adopt is that the Go 1.3 compiler must compile using Go 1.2, Go 1.4 must compile using Go 1.3, and so on. Then there is a clear path to generating current binaries: build the Go 1.2 toolchain (written in C), use it to build the Go 1.3 toolchain, and so on. There will be a shell script to do this; it will take CPU time but not human time. The bootstrapping only needs to be done once per machine; the Go 1.x binaries can be kept in a known location and reused each time all.bash is run during the development of Go 1.(x+1).

有了用Go编写的Go编译器，就必须有一个从头开始的引导计划。我们计划采用的规则是，Go 1.3编译器必须使用Go 1.2进行编译，Go 1.4必须使用Go 1.3进行编译，以此类推。那么就有一条明确的路径来生成当前的二进制文件：构建Go 1.2工具链（用C语言编写），用它来构建Go 1.3工具链，以此类推。会有一个shell脚本来做这件事；它将花费CPU时间，但不需要人的时间。引导工作只需在每台机器上进行一次；Go 1.x的二进制文件可以保存在一个已知的位置，并在Go 1.(x+1)的开发过程中每次运行all.bash时重复使用。

Obviously, this bootstrapping path scales poorly over time. Before too many releases have gone by, it may make sense to write a back end for the compiler that generates C code. The code need not be efficient or readable, just correct. That C version would be checked in, just as today we check in the y.tab.c file generated by yacc. The bootstrap sequence would invoke gcc on that C code to build a bootstrap compiler, and the bootstrap compiler would be used to build the real compiler. Like in the other scheme, the bootstrap compiler binary can be kept in a known location and reused (not rebuilt) each time all.bash is run.

很明显，这种引导路径随着时间的推移，其扩展性很差。在太多的版本过去之前，为编译器编写一个生成C代码的后端可能是有意义的。这些代码不需要高效或可读，只需要正确。这个C版本将被检查，就像今天我们检查yacc生成的y.tab.c文件一样。引导序列会在C代码上调用gcc来构建引导编译器，引导编译器会被用来构建真正的编译器。和其他方案一样，引导编译器的二进制文件可以保存在一个已知的位置，并且在每次运行all.bash时都可以重复使用（而不是重建）。

## Alternatives 替代方案

There are a few alternatives that would be obvious approaches to consider, and so it is worth explaining why we have decided against them.

有一些替代方案是可以考虑的，所以值得解释一下为什么我们决定不采用这些方案。

*Write new compilers from scratch*. The current compilers do have one very important property: they compile Go correctly (or at least correctly enough for nearly all current users). Despite Go’s simplicity, there are many subtle cases in the optimizations and other rewrites performed by the compilers, and it would be foolish to throw away the 10 or so man-years of effort that have gone into them.

从头开始编写新的编译器。目前的编译器确实有一个非常重要的特性：它们能正确地编译Go（或者至少对目前几乎所有的用户来说足够正确）。尽管Go很简单，但在编译器进行的优化和其他重写中，有许多微妙的情况，如果丢弃在这些编译器上花费了10来年的努力，那就太愚蠢了。

*Translate the compiler manually*. We have translated other, smaller C and C++ programs to Go manually. The process is tedious and therefore error-prone, and the mistakes can be very subtle and difficult to find. A mechanical translator will instead generate translations with consistent classes of errors, which should be easier to find, and it will not zone out during the tedious parts. The Go compilers are also significantly larger than anything we’ve converted: over 60,000 lines of C. Mechanical help will make the job much easier. As Dick Sites wrote in 1974, “I would rather write programs to help me write programs than write programs.” Translating the compiler mechanically also makes it easier for development on the C originals to proceed unhindered until we are ready for the switch.

手动翻译编译器。我们已经将其他较小的C和C++程序手动翻译成Go。这个过程很繁琐，因此很容易出错，而且错误可能非常细微，很难发现。机械翻译器反而会产生具有一致错误类别的译文，这应该更容易找到，而且在繁琐的部分也不会出现区隔。Go编译器也比我们转换过的任何东西都大得多：超过60,000行的C语言，机械帮助将使工作更容易。正如Dick Sites在1974年写道："我宁愿写程序来帮助我写程序，而不是写程序"。机械地翻译编译器也使得在C语言原件上的开发更容易不受阻碍地进行，直到我们准备好进行转换。

*Translate just the back ends and connect to go/parser and go/types immediately*. The data structures in the compiler that convey information from the front end to the back ends look nothing like the APIs presented by go/parser and go/types. Replacing the front end by those libraries would require writing code to convert from the go/parser and go/types data structures into the ones expected by the back ends, a very broad and error-prone undertaking. We do believe that it makes sense to use these packages, but it also makes sense to wait until the compiler is structured more like a Go program, into documented sub-packages of its own with defined boundaries and unit tests.

只翻译后端，立即连接到go/parser和go/types。编译器中从前端向后端传递信息的数据结构与go/parser和go/types所提供的API完全不同。用这些库来替代前端需要编写代码，将go/parser和go/types的数据结构转换为后端所期望的数据结构，这是一项非常广泛和容易出错的工作。我们相信使用这些包是有意义的，但是等到编译器的结构更像一个Go程序，变成自己的文档化的子包，并有明确的边界和单元测试，也是有意义的。

*Discard the current compilers and use gccgo (or go/parser + go/types + LLVM, or …)*. The current compilers are a large part of Go’s flexibility. Tying development of Go to a comparatively larger code base like GCC or LLVM seems likely to hurt that flexibility. Also, GCC is a large C (now partly C++) program and LLVM a large C++ program. All the reasons listed above justifying a move away from the current compiler code apply as much or more to these code bases.

抛弃目前的编译器，使用gccgo（或go/parser + go/types + LLVM，或...）。目前的编译器是 Go 灵活性的一个重要组成部分。将Go的开发与GCC或LLVM这样一个相对较大的代码库捆绑在一起，似乎会损害这种灵活性。另外，GCC是一个大型的C语言（现在部分是C++）程序，LLVM是一个大型的C++程序。上面列举的所有理由都证明了放弃目前的编译器代码对这些代码库同样适用，甚至更加适用。

## Long Term Use of C C语言的长期使用

Carried to completion, this plan still leaves the rest of the Plan 9 toolchain written in C. In the long term it would be nice to eliminate all C from the tree. This section speculates on how that might happen. It is not guaranteed to happen in this way or at all.

完成这个计划后，Plan 9工具链的其余部分仍然是用C语言编写的。从长远来看，最好是能从树上消除所有的C语言。本节将推测如何实现这一目标。并不保证会以这种方式发生，或者根本不会发生。

*Package runtime*. Most of the runtime is written in C, for many of the same reasons that the Go compiler is written in C. However, the runtime is much smaller than the compilers and it is already written in a mix of Go and C. It is plausible to convert the C to Go one piece at a time. The major pieces are the scheduler, the garbage collector, the hash map implementation, and the channel implementation. (The fine mixing of Go and C is possible here because the C is compiled with 6c, not gcc.)

包运行时。大部分运行时是用C语言编写的，这与Go编译器用C语言编写的原因相同。然而，运行时比编译器小得多，而且它已经是用Go和C语言混合编写的。主要的部分是调度器，垃圾收集器，哈希图的实现，以及通道的实现。(在这里，Go和C的精细混合是可能的，因为C是用6c而不是gcc编译的）。

*C compilers*. The Plan 9 C compilers are themselves written in C. If we remove all the C from Go package implementations (in particular, package runtime), we can remove these compilers: “go tool 6c” and so on would be no more, and .c files in Go package directory sources would no longer be supported. We would need to announce these plans early, so that external packages written partly in C have time to remove their uses. (Cgo, which uses gcc instead of 6c, would remain as a way to write parts of a package in C.) The Go 1 compatibility document excludes changes to the toolchain; deleting the C compilers is permitted.

C语言编译器。Plan 9的C语言编译器本身就是用C语言编写的。如果我们从Go包的实现中删除所有的C语言（特别是包的运行时间），我们可以删除这些编译器。"go tool 6c "等将不复存在，Go包目录源中的.c文件将不再被支持。我们需要尽早宣布这些计划，以便部分用C语言编写的外部软件包有时间删除它们的用途。(使用gcc而不是6c的Cgo将继续作为用C语言编写部分软件包的一种方式。) Go 1的兼容性文件不包括对工具链的改变；删除C编译器是允许的。

*Assemblers*. The Plan 9 assemblers are also written in C. However, the assembler is little more than a simple parser coupled with a serialization of the parse tree. That could easily be translated to Go, either automatically or by hand.

汇编器。Plan 9的汇编程序也是用C语言编写的。然而，汇编程序只不过是一个简单的解析器，加上解析树的序列化。这很容易被翻译成Go，无论是自动还是手工。

*Linkers*. The Plan 9 linkers are also written in C. Recent work has moved most of the linker in into the compilers, and there is already a plan to rewrite what is left as a new, much simpler Go program. The part of the linker that has moved into the Go compiler will now need to be translated along with the rest of the compiler.

链接器。Plan 9的链接器也是用C语言编写的。最近的工作将大部分链接器移到了编译器中，并且已经有计划将剩下的部分重写成一个新的、更简单的Go程序。链接器中已经移到Go编译器中的部分现在需要和编译器的其他部分一起被翻译。

*Libmach-based tools: nm, pack, addr2line, and objdump*. Nm has already been rewritten in Go. Pack and addr2line can be rewritten any day. Objdump currently depends on libmach’s disassemblers, but those should be straightforward to convert to go, whether mechanically or manually, and at that point libmach itself can be deleted.

基于 Libmach 的工具：nm、pack、addr2line 和 objdump。Nm 已经用 Go 重写了。Pack和addr2line随时可能被重写。Objdump目前依赖于libmach的反汇编器，但这些应该可以直接转换为Go，无论是机械的还是手动的，到那时libmach本身就可以被删除了。