+++
title = "link"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# link

> 原文：[https://pkg.go.dev/cmd/link@go1.19.3](https://pkg.go.dev/cmd/link@go1.19.3)

### Overview 概述

​	`link`，通常以 "`go tool link` "的形式调用，读取包 `main` 的 Go 存档或对象及其依赖项，并将它们组合成一个可执行的二进制文件。

#### Command Line 命令行

使用方法：

```
go tool link [flags] main.a
```

标志：

```
-B note
	Add an ELF_NT_GNU_BUILD_ID note when using ELF.
	The value should start with 0x and be an even number of hex digits.
	=>在使用ELF时，添加一个ELF_NT_GNU_BUILD_ID注释。
	这个值应该以0x开始，并且是一个偶数的十六进制数字。
	
-D address
	Set data segment address.
	=> 设置数据段地址。
	
-E entry
	Set entry symbol name.
	=> 设置条目符号名称。
	
-H type
	Set executable format type.
	The default format is inferred from GOOS and GOARCH.
	On Windows, -H windowsgui writes a "GUI binary" instead of a "console binary."
	=> 设置可执行格式类型。
	默认格式是由GOOS和GOARCH推断出来的。
	在Windows上，-H windowsgui 编写 "GUI binary"，而不是 "console binary"。
	
-I interpreter
	Set the ELF dynamic linker to use.
	=> 设置要使用的 ELF 动态链接器。
	
-L dir1 -L dir2
	Search for imported packages in dir1, dir2, etc,
	after consulting $GOROOT/pkg/$GOOS_$GOARCH.
	=> 在dir1、dir2等处搜索导入的软件包。
	在查询$GOROOT/pkg/$GOOS_$GOARCH之后。
	
-R quantum
	Set address rounding quantum.
	=> 设置地址舍入量子。
	
-T address
	Set text segment address.
	=> 设置文本段地址。
	
-V
	Print linker version and exit.
	=> 打印链接器版本并退出。
	
-X importpath.name=value
	Set the value of the string variable in importpath named name to value.
	This is only effective if the variable is declared in the source code either uninitialized
	or initialized to a constant string expression. -X will not work if the initializer makes
	a function call or refers to other variables.
	Note that before Go 1.5 this option took two separate arguments.
	=> 将 importpath 中名为 name 的字符串变量的值设置为值。
	只有当该变量在源代码中未被初始化或被初始化为常量字符串表达式时，这才有效。
	或初始化为一个常量的字符串表达式。如果初始化器进行了函数调用或引用了其他变量，-X将不起作用。
	函数调用或引用其他变量时，-X不起作用。
	注意，在Go 1.5之前，这个选项需要两个单独的参数。
	
-a
	Disassemble output.
	=> 反汇编输出。
	
-asan
	Link with C/C++ address sanitizer support.
	=> 支持C/C++地址净化器的链接。
	
-buildid id
	Record id as Go toolchain build id.
	=> 记录ID作为Go工具链的构建ID。
	
-buildmode mode
	Set build mode (default exe).
	=> 设置构建模式（默认为 exe）。
	
-c
	Dump call graphs.
	=> 转储调用图。
	
-compressdwarf
	Compress DWARF if possible (default true).
	=> 如果可能的话，压缩DWARF（默认为true）。
	
-cpuprofile file
	Write CPU profile to file.
	=> 将CPU配置文件写到文件中。
	
-d
	Disable generation of dynamic executables.
	The emitted code is the same in either case; the option
	controls only whether a dynamic header is included.
	The dynamic header is on by default, even without any
	references to dynamic libraries, because many common
	system tools now assume the presence of the header.
	=> 禁用动态可执行文件的生成。
	在这两种情况下发出的代码都是一样的；
	该选项只控制是否包含动态头。
	动态头默认是打开的，即使没有任何对动态库的引用，
	因为现在许多常见的系统工具都假定存在头。
	
-debugtramp int
	Debug trampolines.
	=> 
	
-dumpdep
	Dump symbol dependency graph.
	=> 转储符号依赖图。
	
-extar ar
	Set the external archive program (default "ar").
	Used only for -buildmode=c-archive.
	=> 设置外部存档程序（默认为 "ar"）。
	仅用于 -buildmode=c-archive。
	
-extld linker
	Set the external linker (default "clang" or "gcc").
	=> 设置外部链接器（默认为 "clang "或 "gcc"）。
	
-extldflags flags
	Set space-separated flags to pass to the external linker.
	=> 设置以空格分隔的标志，传递给外部链接器。
	
-f
	Ignore version mismatch in the linked archives.
	=> 忽略链接存档中的版本不匹配。
	
-g
	Disable Go package data checks.
	=> 停用 Go 包的数据检查。
	
-importcfg file
	Read import configuration from file.
	In the file, set packagefile, packageshlib to specify import resolution.
	=> 从 file 中读取导入配置。
	在 file 中，设置packagefile、packageshlib来指定导入解析。
	
-installsuffix suffix
	Look for packages in $GOROOT/pkg/$GOOS_$GOARCH_suffix
	instead of $GOROOT/pkg/$GOOS_$GOARCH.
	=> 在$GOROOT/pkg/$GOOS_$GOARCH_suffix中寻找包
	而不是在 $GOROOT/pkg/$GOOS_$GOARCH。
	
-k symbol
	Set field tracking symbol. Use this flag when GOEXPERIMENT=fieldtrack is set.
	=> 设置字段跟踪符号。当设置GOEXPERIMENT=fieldtrack时，使用此标志。
	
-libgcc file
	Set name of compiler support library.
	This is only used in internal link mode.
	If not set, default value comes from running the compiler,
	which may be set by the -extld option.
	Set to "none" to use no support library.
	=> 设置编译器支持库的名称。
	这只在内部链接模式下使用。
	如果不设置，默认值来自运行的编译器，
	可以通过-extld选项设置。
	设置为 "none"表示不使用支持库。
	
-linkmode mode
	Set link mode (internal, external, auto).
	This sets the linking mode as described in cmd/cgo/doc.go.
	=> 设置链接模式（内部、外部、自动）。
	这设置了cmd/cgo/doc.go中描述的链接模式。
	
-linkshared
	Link against installed Go shared libraries (experimental).
	=> 针对已安装的 Go 共享库进行链接（实验性）。
	
-memprofile file
	Write memory profile to file.
	=> 将内存配置文件写入 file。
	
-memprofilerate rate
	Set runtime.MemProfileRate to rate.
	=> 将 runtime.MemProfileRate 设置为 rate。
	
-msan
	Link with C/C++ memory sanitizer support.
	=> 支持用C/C++内存净化器进行链接。
	
-n
	Dump symbol table.
	=> 转储符号表。
	
-o file
	Write output to file (default a.out, or a.out.exe on Windows).
	=> 将输出写入 file（默认为a.out，或Windows下的a.out.exe）。
	
-pluginpath path
	The path name used to prefix exported plugin symbols.
	=> 用来给导出的插件符号加前缀的 path 名称。
	
-r dir1:dir2:...
	Set the ELF dynamic linker search path.
	=> 设置ELF动态链接器的搜索路径。
	
-race
	Link with race detection libraries.
	=> 用竞争检测库链接。
	
-s
	Omit the symbol table and debug information.
	=> 省略符号表和调试信息。
	
-shared
	Generated shared object (implies -linkmode external; experimental).
	=> 生成的共享对象（意味着-linkmode external; 实验性）。
	
-tmpdir dir
	Write temporary files to dir.
	Temporary files are only used in external linking mode.
	=> 将临时文件写到 dir。
	临时文件只在外部链接模式下使用。
	
-u
	Reject unsafe packages.
	=> 拒绝不安全的件包。
	
-v
	Print trace of linker operations.
	=> 打印链接器操作的痕迹。
	
-w
	Omit the DWARF symbol table.
	=> 省略DWARF符号表。
	
```