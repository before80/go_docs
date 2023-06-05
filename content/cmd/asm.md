+++
title = "asm"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# asm

> 原文：https://pkg.go.dev/cmd/asm@go1.19.3

### Overview 概述

​	`asm`，通常以 "`go tool asm` "调用，将源文件组装成一个对象文件，该文件以参数源文件的基本名称命名，并使用`.o`后缀。之后，该对象文件可以与其他对象组合成一个包存档。

#### Command Line 命令行

使用方法：

```
go tool asm [flags] file
```

​	指定的 `file` 必须是一个Go汇编文件。所有的目标操作系统和体系结构都使用同一个汇编器。`GOOS`和`GOARCH`环境变量设置所需的目标。

标志：

```
-D name[=value]
	Predefine symbol name with an optional simple value.
	Can be repeated to define multiple symbols.
	=>	用一个可选的简单值预先定义符号名称。
	可以重复定义多个符号。
	
-I dir1 -I dir2
	Search for #include files in dir1, dir2, etc,
	after consulting $GOROOT/pkg/$GOOS_$GOARCH.
	=>	搜索dir1、dir2等地方的#include文件。
	在查阅了$GOROOT/pkg/$GOOS_$GOARCH之后。
	
-S
	Print assembly and machine code.
	=> 	打印汇编和机器代码。

		
-V
	Print assembler version and exit.
	=> 	打印汇编程序版本并退出。

	
-debug
	Dump instructions as they are parsed.
	=> 	在解析指令的过程中倾倒指令。

	
-dynlink
	Support references to Go symbols defined in other shared libraries.
	=> 	支持对其他共享库中定义的 Go 符号的引用。

	
-gensymabis
	Write symbol ABI information to output file. Don't assemble.
	=> 	将符号ABI信息写入输出文件。不要组装。

	
-o file
	Write output to file. The default is foo.o for /a/b/c/foo.s.
	=> 	将输出写入文件。默认是/a/b/c/foo.s的foo.o。

	
-shared
	Generate code that can be linked into a shared library.
	=> 	生成可以被链接到共享库中的代码。

	
-spectre list
	Enable spectre mitigations in list (all, ret).
	=> 	在列表中启用spectre缓解措施（全部，ret）。

	
-trimpath prefix
	Remove prefix from recorded source file paths.
	=> 	从记录的源文件路径中移除前缀。

	
```

Input language:

输入语言：

The assembler uses mostly the same syntax for all architectures, the main variation having to do with addressing modes. Input is run through a simplified C preprocessor that implements #include, #define, #ifdef/endif, but not #if or ##.

​	汇编器对所有的架构都使用相同的语法，主要的变化是与寻址模式有关。输入是通过一个简化的C预处理器运行的，它实现了#include、#define、#ifdef/endif，但没有#if或##。

​	更多信息请参见https://golang.org/doc/asm。



=== "doc.go"

```

```

=== "main.go"

```

```



