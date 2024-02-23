+++
title = "compile"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# compile

> 原文：[https://pkg.go.dev/cmd/compile@go1.19.3](https://pkg.go.dev/cmd/compile@go1.19.3)

### Overview 

Compile, typically invoked as “go tool compile,” compiles a single Go package comprising the files named on the command line. It then writes a single object file named for the basename of the first source file with a .o suffix. The object file can then be combined with other objects into a package archive or passed directly to the linker (“go tool link”). If invoked with -pack, the compiler writes an archive directly, bypassing the intermediate object file.

​	编译，通常称为“go tool compile”，编译一个由命令行上命名的文件组成的单个 Go 包。然后，它会为第一个源文件的基础名称编写一个名为 .o 后缀的对象文件。然后，可以将对象文件与其他对象组合成一个包存档，或者直接传递给链接器（“go tool link”）。如果使用 -pack 调用，编译器会直接编写一个存档，绕过中间对象文件。

The generated files contain type information about the symbols exported by the package and about types used by symbols imported by the package from other packages. It is therefore not necessary when compiling client C of package P to read the files of P's dependencies, only the compiled output of P.

​	生成的这些文件包含有关包导出的符号以及包从其他包导入的符号所使用的类型的类型信息。因此，在编译包 P 的客户端 C 时，无需读取 P 的依赖项的文件，只需读取 P 的已编译输出即可。

#### Command Line

Usage:

​	用法：

```
go tool compile [flags] file...
```

The specified files must be Go source files and all part of the same package. The same compiler is used for all target operating systems and architectures. The GOOS and GOARCH environment variables set the desired target.

​	指定的文件必须是 Go 源文件，并且都属于同一个包。所有目标操作系统和体系结构都使用相同的编译器。GOOS 和 GOARCH 环境变量设置所需的 target。

Flags:

​	标志：

```
-D path
	Set relative path for local imports.
-I dir1 -I dir2
	Search for imported packages in dir1, dir2, etc,
	after consulting $GOROOT/pkg/$GOOS_$GOARCH.
-L
	Show complete file path in error messages.
-N
	Disable optimizations.
-S
	Print assembly listing to standard output (code only).
-S -S
	Print assembly listing to standard output (code and data).
-V
	Print compiler version and exit.
-asmhdr file
	Write assembly header to file.
-asan
	Insert calls to C/C++ address sanitizer.
-buildid id
	Record id as the build id in the export metadata.
-blockprofile file
	Write block profile for the compilation to file.
-c int
	Concurrency during compilation. Set 1 for no concurrency (default is 1).
-complete
	Assume package has no non-Go components.
-cpuprofile file
	Write a CPU profile for the compilation to file.
-dynlink
	Allow references to Go symbols in shared libraries (experimental).
-e
	Remove the limit on the number of errors reported (default limit is 10).
-goversion string
	Specify required go tool version of the runtime.
	Exits when the runtime go version does not match goversion.
-h
	Halt with a stack trace at the first error detected.
-importcfg file
	Read import configuration from file.
	In the file, set importmap, packagefile to specify import resolution.
-installsuffix suffix
	Look for packages in $GOROOT/pkg/$GOOS_$GOARCH_suffix
	instead of $GOROOT/pkg/$GOOS_$GOARCH.
-l
	Disable inlining.
-lang version
	Set language version to compile, as in -lang=go1.12.
	Default is current version.
-linkobj file
	Write linker-specific object to file and compiler-specific
	object to usual output file (as specified by -o).
	Without this flag, the -o output is a combination of both
	linker and compiler input.
-m
	Print optimization decisions. Higher values or repetition
	produce more detail.
-memprofile file
	Write memory profile for the compilation to file.
-memprofilerate rate
	Set runtime.MemProfileRate for the compilation to rate.
-msan
	Insert calls to C/C++ memory sanitizer.
-mutexprofile file
	Write mutex profile for the compilation to file.
-nolocalimports
	Disallow local (relative) imports.
-o file
	Write object to file (default file.o or, with -pack, file.a).
-p path
	Set expected package import path for the code being compiled,
	and diagnose imports that would cause a circular dependency.
-pack
	Write a package (archive) file rather than an object file
-race
	Compile with race detector enabled.
-s
	Warn about composite literals that can be simplified.
-shared
	Generate code that can be linked into a shared library.
-spectre list
	Enable spectre mitigations in list (all, index, ret).
-traceprofile file
	Write an execution trace to file.
-trimpath prefix
	Remove prefix from recorded source file paths.
```

Flags related to debugging information:

​	与调试信息相关的标志：

```
-dwarf
	Generate DWARF symbols.
-dwarflocationlists
	Add location lists to DWARF in optimized mode.
-gendwarfinl int
	Generate DWARF inline info records (default 2).
```

Flags to debug the compiler itself:

​	用于调试编译器本身的标志：

```
-E
	Debug symbol export.
-K
	Debug missing line numbers.
-d list
	Print debug information about items in list. Try -d help for further information.
-live
	Debug liveness analysis.
-v
	Increase debug verbosity.
-%
	Debug non-static initializers.
-W
	Debug parse tree after type checking.
-f
	Debug stack frames.
-i
	Debug line number stack.
-j
	Debug runtime-initialized variables.
-r
	Debug generated wrappers.
-w
	Debug type checking.
```

#### Compiler Directives [¶](https://pkg.go.dev/cmd/compile@go1.19.3#hdr-Compiler_Directives)

The compiler accepts directives in the form of comments. To distinguish them from non-directive comments, directives require no space between the comment opening and the name of the directive. However, since they are comments, tools unaware of the directive convention or of a particular directive can skip over a directive like any other comment.

​	编译器接受以注释形式给出的指令。为了将它们与非指令注释区分开来，指令要求注释开头与指令名称之间没有空格。但是，由于它们是注释，因此不了解指令约定或特定指令的工具可以像跳过任何其他注释一样跳过指令。

Line directives come in several forms:

​	行指令有几种形式：

```
//line :line
//line :line:col
//line filename:line
//line filename:line:col
/*line :line*/
/*line :line:col*/
/*line filename:line*/
/*line filename:line:col*/
```

In order to be recognized as a line directive, the comment must start with //line or /*line followed by a space, and must contain at least one colon. The //line form must start at the beginning of a line. A line directive specifies the source position for the character immediately following the comment as having come from the specified file, line and column: For a //line comment, this is the first character of the next line, and for a /*line comment this is the character position immediately following the closing */. If no filename is given, the recorded filename is empty if there is also no column number; otherwise it is the most recently recorded filename (actual filename or filename specified by previous line directive). If a line directive doesn't specify a column number, the column is "unknown" until the next directive and the compiler does not report column numbers for that range. The line directive text is interpreted from the back: First the trailing :ddd is peeled off from the directive text if ddd is a valid number > 0. Then the second :ddd is peeled off the same way if it is valid. Anything before that is considered the filename (possibly including blanks and colons). Invalid line or column values are reported as errors.

​	 为了被识别为行指令，注释必须以 //line 或 /*line 开头，后跟一个空格，并且必须至少包含一个冒号。//line 形式必须从行的开头开始。行指令指定紧跟注释之后的字符的源位置来自指定的文件、行和列：对于 //line 注释，这是下一行的第一个字符，对于 /*line 注释，这是紧跟关闭的 */ 之后的字符位置。如果没有给出文件名，则在没有列号的情况下，记录的文件名为空；否则，它是最近记录的文件名（实际文件名或由前一行指令指定的文件名）。如果行指令未指定列号，则列为“未知”，直到下一条指令，并且编译器不会报告该范围的列号。行指令文本从后往前解释：首先，如果 ddd 是一个大于 0 的有效数字，则从指令文本中剥离掉尾随的 :ddd。然后，如果有效，则以相同的方式剥离第二个 :ddd。在此之前的任何内容都被视为文件名（可能包括空格和冒号）。 无效的行或列值报告为错误。

Examples:

​	示例：

```
//line foo.go:10      the filename is foo.go, and the line number is 10 for the next line
//line C:foo.go:10    colons are permitted in filenames, here the filename is C:foo.go, and the line is 10
//line  a:100 :10     blanks are permitted in filenames, here the filename is " a:100 " (excluding quotes)
/*line :10:20*/x      the position of x is in the current file with line number 10 and column number 20
/*line foo: 10 */     this comment is recognized as invalid line directive (extra blanks around line number)
```

Line directives typically appear in machine-generated code, so that compilers and debuggers will report positions in the original input to the generator.

​	行指令通常出现在机器生成的代码中，以便编译器和调试器将报告位置报告给生成器的原始输入。

The line directive is a historical special case; all other directives are of the form //go:name, indicating that they are defined by the Go toolchain. Each directive must be placed its own line, with only leading spaces and tabs allowed before the comment. Each directive applies to the Go code that immediately follows it, which typically must be a declaration.

​	行指令是一个历史特殊情况；所有其他指令都采用 //go:name 形式，表示它们由 Go 工具链定义。每个指令都必须放在自己的行中，注释前只能有前导空格和制表符。每个指令都适用于紧随其后的 Go 代码，通常必须是声明。

```
//go:noescape
```

The //go:noescape directive must be followed by a function declaration without a body (meaning that the function has an implementation not written in Go). It specifies that the function does not allow any of the pointers passed as arguments to escape into the heap or into the values returned from the function. This information can be used during the compiler's escape analysis of Go code calling the function.

​	//go:noescape 指令后面必须紧跟一个没有函数体的函数声明（这意味着该函数具有未用 Go 编写的实现）。它指定该函数不允许任何作为参数传递的指针逃逸到堆中或逃逸到从该函数返回的值中。此信息可用于编译器对调用该函数的 Go 代码进行逃逸分析时使用。

```
//go:uintptrescapes
```

The //go:uintptrescapes directive must be followed by a function declaration. It specifies that the function's uintptr arguments may be pointer values that have been converted to uintptr and must be on the heap and kept alive for the duration of the call, even though from the types alone it would appear that the object is no longer needed during the call. The conversion from pointer to uintptr must appear in the argument list of any call to this function. This directive is necessary for some low-level system call implementations and should be avoided otherwise.

​	//go:uintptrescapes 指令后面必须紧跟一个函数声明。它指定函数的 uintptr 参数可以是已转换为 uintptr 的指针值，并且必须位于堆上并在调用期间保持活动状态，即使从类型本身来看，该对象在调用期间似乎不再需要。从指针到 uintptr 的转换必须出现在对该函数的任何调用的参数列表中。此指令对于某些低级系统调用实现是必需的，否则应避免使用。

```
//go:noinline
```

The //go:noinline directive must be followed by a function declaration. It specifies that calls to the function should not be inlined, overriding the compiler's usual optimization rules. This is typically only needed for special runtime functions or when debugging the compiler.

​	//go:noinline 指令后面必须紧跟一个函数声明。它指定对函数的调用不应内联，从而覆盖编译器的常规优化规则。通常仅在特殊运行时函数或调试编译器时才需要这样做。

```
//go:norace
```

The //go:norace directive must be followed by a function declaration. It specifies that the function's memory accesses must be ignored by the race detector. This is most commonly used in low-level code invoked at times when it is unsafe to call into the race detector runtime.

​	//go:norace 指令后面必须紧跟一个函数声明。它指定函数的内存访问必须被竞态检测器忽略。这最常用于在调用竞态检测器运行时不安全时调用的低级代码中。

```
//go:nosplit
```

The //go:nosplit directive must be followed by a function declaration. It specifies that the function must omit its usual stack overflow check. This is most commonly used by low-level runtime code invoked at times when it is unsafe for the calling goroutine to be preempted.

​	//go:nosplit 指令必须紧跟函数声明。它指定函数必须省略其通常的堆栈溢出检查。这通常由在调用协程被抢占时不安全的时间调用的低级运行时代码使用。

```
//go:linkname localname [importpath.name]
```

This special directive does not apply to the Go code that follows it. Instead, the //go:linkname directive instructs the compiler to use “importpath.name” as the object file symbol name for the variable or function declared as “localname” in the source code. If the “importpath.name” argument is omitted, the directive uses the symbol's default object file symbol name and only has the effect of making the symbol accessible to other packages. Because this directive can subvert the type system and package modularity, it is only enabled in files that have imported "unsafe".

​	此特殊指令不适用于其后的 Go 代码。相反，//go:linkname 指令指示编译器将“importpath.name”用作源代码中声明为“localname”的变量或函数的可执行文件符号名称。如果省略“importpath.name”参数，该指令将使用符号的默认可执行文件符号名称，并且仅使符号可供其他包访问。由于此指令可以破坏类型系统和包模块化，因此仅在已导入“unsafe”的文件中启用它。



=== "doc.go"

```

```

=== "main.go"

```

```

