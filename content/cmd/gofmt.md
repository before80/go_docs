+++
title = "gofmt"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# gofmt

> 原文：[https://pkg.go.dev/cmd/gofmt@go1.19.3](https://pkg.go.dev/cmd/gofmt@go1.19.3)

### Overview 概述

Gofmt formats Go programs. It uses tabs for indentation and blanks for alignment. Alignment assumes that an editor is using a fixed-width font.

​	gofmt可以格式化Go程序。它使用制表符来缩进，使用空白来对齐。对齐的前提是编辑器使用固定宽度的字体。

Without an explicit path, it processes the standard input. Given a file, it operates on that file; given a directory, it operates on all `.go` files in that directory, recursively. (Files starting with a period are ignored.) By default, gofmt prints the reformatted sources to standard output.

​	没有明确的路径，它处理标准输入。给定一个文件，它就对该文件进行操作；给定一个目录，它就对该目录中的所有`.go`文件进行递归操作。(以句号开头的文件被忽略。)默认情况下，gofmt将重新格式化的源文件打印到标准输出。

Usage:

​	使用方法：

```
gofmt [flags] [path ...]
```

The flags are:

​	标志是：

```
-d
	Do not print reformatted sources to standard output.
	If a file's formatting is different than gofmt's, print diffs
	to standard output.
	=> 不将重新格式化的源文件打印到标准输出。
	如果一个文件的格式与gofmt的不同，则将差异打印到标准输出。到标准输出。
	
-e
	Print all (including spurious) errors.
	=> 打印所有（包括假的）错误。
	
-l
	Do not print reformatted sources to standard output.
	If a file's formatting is different from gofmt's, print its name
	to standard output.
	=> 不打印重新格式化的来源到标准输出。
	如果一个文件的格式与gofmt不同，则打印其名称
	到标准输出。
	
-r rule
	Apply the rewrite rule to the source before reformatting.
	=> 在重新格式化前对源文件应用重写规则。
	
-s
	Try to simplify code (after applying the rewrite rule, if any).
	=> 尝试简化代码（在应用重写规则后，如果有的话）。
	
-w
	Do not print reformatted sources to standard output.
	If a file's formatting is different from gofmt's, overwrite it
	with gofmt's version. If an error occurred during overwriting,
	the original file is restored from an automatic backup.
	=> 不将重新格式化的源文件打印到标准输出。
	如果一个文件的格式与gofmt的不同，就用gofmt的版本覆盖它。	用gofmt的版本覆盖。如果在覆盖的过程中发生错误。	原始文件将从自动备份中恢复。
```

Debugging support:

​	调试支持：

```
-cpuprofile filename
	Write cpu profile to the specified file.
```

The rewrite rule specified with the -r flag must be a string of the form:

​	用-r标志指定的重写规则必须是一个形式的字符串：

```
pattern -> replacement
```

Both pattern and replacement must be valid Go expressions. In the pattern, single-character lowercase identifiers serve as wildcards matching arbitrary sub-expressions; those expressions will be substituted for the same identifiers in the replacement.

​	pattern和replacement都必须是有效的Go表达式。在模式中，单字符小写标识符作为通配符匹配任意子表达式；这些表达式将被替换为替换中的相同标识符。

When gofmt reads from standard input, it accepts either a full Go program or a program fragment. A program fragment must be a syntactically valid declaration list, statement list, or expression. When formatting such a fragment, gofmt preserves leading indentation as well as leading and trailing spaces, so that individual sections of a Go program can be formatted by piping them through gofmt.

​	当gofmt从标准输入读取时，它接受一个完整的Go程序或程序片段。程序片段必须是一个语法上有效的声明列表、语句列表或表达式。当格式化这样的片段时，gofmt 会保留前导缩进以及前导和后导空格，因此 Go 程序的各个部分可以通过 gofmt 的管道来进行格式化。

#### Examples  例子

To check files for unnecessary parentheses:

​	检查文件是否有不必要的括号：

```
gofmt -r '(a) -> a' -l *.go
```

To remove the parentheses:

​	要删除括号：

```
gofmt -r '(a) -> a' -w *.go
```

To convert the package tree from explicit slice upper bounds to implicit ones:

​	将包树从明确的片断上界转换为隐含的上界：

```
gofmt -r 'α[β:len(α)] -> α[β:]' -w $GOROOT/src
```

#### The simplify command 简化命令

When invoked with -s gofmt will make the following source transformations where possible.

​	当与-s一起调用时，gofmt将尽可能地进行以下源转换。

```
An array, slice, or map composite literal of the form:
	[]T{T{}, T{}}
will be simplified to:
	[]T{{}, {}}

A slice expression of the form:
	s[a:len(s)]
will be simplified to:
	s[a:]

A range of the form:
	for x, _ = range v {...}
will be simplified to:
	for x = range v {...}

A range of the form:
	for _ = range v {...}
will be simplified to:
	for range v {...}
```

This may result in changes that are incompatible with earlier versions of Go.

​	这可能会导致与Go早期版本不兼容的变化。

### Notes  注意事项

### Bugs 

- The implementation of -r is a bit slow.
- -r的实现有点慢。
- If -w fails, the restored original file may not have some of the original file attributes.
- 如果-w失败，恢复的原始文件可能没有一些原始文件的属性。