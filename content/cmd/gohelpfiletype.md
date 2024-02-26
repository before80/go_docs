+++
title = "go help filetype"
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 540
description = ""
isCJKLanguage = true
draft = false

+++

### extensions 

The go command examines the contents of a restricted set of files in each directory. It identifies which files to examine based on the extension of the file name. These extensions are:

​	go 命令检查每个目录中一组受限制的文件的内容。它根据文件名的扩展名来识别要检查的文件。这些扩展名包括：

### .go

​        Go source files.

​	 Go 源文件。

### .c, .h

​        C source files. If the package uses cgo or SWIG, these will be compiled with the OS-native compiler (typically gcc); otherwise they will trigger an error.

​	 C 源文件。如果包使用 cgo 或 SWIG，则将使用 OS 本地编译器（通常是 gcc）进行编译；否则，它们将触发错误。

### .cc, .cpp, .cxx, .hh, .hpp, .hxx

​        C++ source files. Only useful with cgo or SWIG, and always compiled with the OS-native compiler.

​	 C++ 源文件。仅在使用 cgo 或 SWIG 时有用，始终使用 OS 本地编译器进行编译。

### .m

​        Objective-C source files. Only useful with cgo, and always compiled with the OS-native compiler.

​	Objective-C 源文件。仅在使用 cgo 时有用，始终使用 OS 本地编译器进行编译。

### .s, .S, .sx

​        Assembler source files. If the package uses cgo or SWIG, these will be assembled with the  OS-native assembler (typically gcc (sic)); otherwise they  will be assembled with the Go assembler.

​	汇编源文件。如果包使用 cgo 或 SWIG，则将使用 OS 本地汇编器（通常是 gcc）进行汇编；否则，它们将使用 Go 汇编器进行汇编。

### .swig, .swigcxx

​        SWIG definition files.

​	 SWIG 定义文件。

### .syso

​        System object files.

​	 系统对象文件。

Files of each of these types except .syso may contain build constraints, but the go command stops scanning for build constraints
at the first item in the file that is not a blank line or //-style line comment. See the go/build package documentation for more details.

​	除了 .syso 类型的文件外，每种类型的文件都可以包含构建约束，但是 go 命令在文件中找到的第一个不是空行或 //-style 行注释的项目处停止扫描构建约束。有关更多详细信息，请参阅 go/build 包文档。
