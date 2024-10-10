+++
title = "scanner"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/scanner@go1.23.0](https://pkg.go.dev/go/scanner@go1.23.0)

Package scanner implements a scanner for Go source text. It takes a []byte as source which can then be tokenized through repeated calls to the Scan method.

​	scanner 包实现了一个用于 Go 源代码的扫描器。它将 []byte 作为源代码，然后可以通过对 Scan 方法的重复调用对其进行标记化。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func PrintError

```go
func PrintError(w io.Writer, err error)
```

PrintError is a utility function that prints a list of errors to w, one error per line, if the err parameter is an ErrorList. Otherwise it prints the err string.

​	PrintError 是一个实用函数，如果 err 参数是 ErrorList，则将错误列表打印到 w，每行一个错误。否则，它将打印 err 字符串。

## 类型

### type Error

```go
type Error struct {
	Pos token.Position
	Msg string
}
```

In an ErrorList, an error is represented by an `*Error`. The position Pos, if valid, points to the beginning of the offending token, and the error condition is described by Msg.

​	在 ErrorList 中，错误由 `*Error` 表示。如果有效，位置 Pos 指向违规标记的开头，错误条件由 Msg 描述。

#### (Error) Error

```go
func (e Error) Error() string
```

Error implements the error interface.

​	Error 实现错误接口。

### type ErrorHandler

```go
type ErrorHandler func(pos token.Position, msg string)
```

An ErrorHandler may be provided to Scanner.Init. If a syntax error is encountered and a handler was installed, the handler is called with a position and an error message. The position points to the beginning of the offending token.

​	ErrorHandler 可以提供给 Scanner.Init。如果遇到语法错误并且安装了处理程序，则使用位置和错误消息调用处理程序。该位置指向违规标记的开头。

### type ErrorList

```go
type ErrorList []*Error
```

ErrorList is a list of `*Errors`. The zero value for an ErrorList is an empty ErrorList ready to use.

​	ErrorList 是 `*Errors` 的列表。ErrorList 的零值是一个空的 ErrorList，可以使用。

#### (*ErrorList) Add

```go
func (p *ErrorList) Add(pos token.Position, msg string)
```

Add adds an Error with given position and error message to an ErrorList.

​	Add 将具有给定位置和错误消息的 Error 添加到 ErrorList。

#### (ErrorList) Err

```go
func (p ErrorList) Err() error
```

Err returns an error equivalent to this error list. If the list is empty, Err returns nil.

​	Err 返回等效于此错误列表的错误。如果列表为空，Err 返回 nil。

#### (ErrorList) Error

```go
func (p ErrorList) Error() string
```

An ErrorList implements the error interface.

​	ErrorList 实现错误接口。

#### (ErrorList) Len

```go
func (p ErrorList) Len() int
```

ErrorList implements the sort Interface.

​	ErrorList 实现排序接口。

#### (ErrorList) Less

```go
func (p ErrorList) Less(i, j int) bool
```

#### (*ErrorList) RemoveMultiples

```go
func (p *ErrorList) RemoveMultiples()
```

RemoveMultiples sorts an ErrorList and removes all but the first error per line.

​	RemoveMultiples 对 ErrorList 进行排序，并删除每行除第一个错误之外的所有错误。

#### (*ErrorList) Reset

```go
func (p *ErrorList) Reset()
```

Reset resets an ErrorList to no errors.

​	Reset 将 ErrorList 重置为无错误。

#### (ErrorList) Sort

```go
func (p ErrorList) Sort()
```

Sort sorts an ErrorList. `*Error` entries are sorted by position, other errors are sorted by error message, and before any *Error entry.

​	Sort 对 ErrorList 进行排序。`*Error` 条目按位置排序，其他错误按错误消息排序，且位于任何 *Error 条目之前。

#### (ErrorList) Swap

```go
func (p ErrorList) Swap(i, j int)
```

### type Mode

```go
type Mode uint
```

A mode value is a set of flags (or 0). They control scanner behavior.

​	模式值是一组标志（或 0）。它们控制扫描器行为。

```go
const (
	ScanComments Mode = 1 << iota // return comments as COMMENT tokens

)
```

### type Scanner

```go
type Scanner struct {

	// public state - ok to modify
	ErrorCount int // number of errors encountered
	// contains filtered or unexported fields
}
```

A Scanner holds the scanner’s internal state while processing a given text. It can be allocated as part of another data structure but must be initialized via Init before use.

​	Scanner 在处理给定文本时保存扫描器的内部状态。它可以作为另一个数据结构的一部分进行分配，但必须在使用前通过 Init 进行初始化。

#### (*Scanner) Init

```go
func (s *Scanner) Init(file *token.File, src []byte, err ErrorHandler, mode Mode)
```

Init prepares the scanner s to tokenize the text src by setting the scanner at the beginning of src. The scanner uses the file set file for position information and it adds line information for each line. It is ok to re-use the same file when re-scanning the same file as line information which is already present is ignored. Init causes a panic if the file size does not match the src size.

​	Init 通过将扫描器设置在 src 的开头来准备扫描器 s 对文本 src 进行标记化。扫描器使用文件集 file 获取位置信息，并为每行添加行信息。在重新扫描同一文件时，可以重复使用同一文件，因为已经存在且被忽略的行信息。如果文件大小与 src 大小不匹配，Init 会导致恐慌。

Calls to Scan will invoke the error handler err if they encounter a syntax error and err is not nil. Also, for each error encountered, the Scanner field ErrorCount is incremented by one. The mode parameter determines how comments are handled.

​	如果 Scan 调用遇到语法错误且 err 不为 nil，则会调用错误处理程序 err。此外，对于遇到的每个错误，扫描器字段 ErrorCount 会增加一。mode 参数确定如何处理注释。

Note that Init may call err if there is an error in the first character of the file.

​	请注意，如果文件第一个字符有错误，Init 可能会调用 err。

#### (*Scanner) Scan 

```go
func (s *Scanner) Scan() (pos token.Pos, tok token.Token, lit string)
```

Scan scans the next token and returns the token position, the token, and its literal string if applicable. The source end is indicated by token.EOF.

​	扫描扫描下一个标记并返回标记位置、标记及其文字字符串（如果适用）。源代码的结束由标记.EOF 指示。

If the returned token is a literal (token.IDENT, token.INT, token.FLOAT, token.IMAG, token.CHAR, token.STRING) or token.COMMENT, the literal string has the corresponding value.

​	如果返回的标记是文字（标记.IDENT、标记.INT、标记.FLOAT、标记.IMAG、标记.CHAR、标记.STRING）或标记.COMMENT，则文字字符串具有相应的值。

If the returned token is a keyword, the literal string is the keyword.

​	如果返回的标记是关键字，则文字字符串是关键字。

If the returned token is token.SEMICOLON, the corresponding literal string is “;” if the semicolon was present in the source, and “\n” if the semicolon was inserted because of a newline or at EOF.

​	如果返回的标记是标记.SEMICOLON，则相应的文字字符串是“；”（如果源代码中存在分号）或“\n”（如果分号是由于换行或在 EOF 处插入的）。

If the returned token is token.ILLEGAL, the literal string is the offending character.

​	如果返回的标记是标记.ILLEGAL，则文字字符串是违规字符。

In all other cases, Scan returns an empty literal string.

​	在所有其他情况下，扫描返回一个空文字字符串。

For more tolerant parsing, Scan will return a valid token if possible even if a syntax error was encountered. Thus, even if the resulting token sequence contains no illegal tokens, a client may not assume that no error occurred. Instead it must check the scanner’s ErrorCount or the number of calls of the error handler, if there was one installed.

​	为了实现更宽容的解析，即使遇到语法错误，扫描也会尽可能返回一个有效标记。因此，即使结果标记序列不包含任何非法标记，客户端也不能假设没有发生错误。相反，它必须检查扫描器的 ErrorCount 或错误处理程序的调用次数（如果有安装）。

Scan adds line information to the file added to the file set with Init. Token positions are relative to that file and thus relative to the file set.

​	Scan 向使用 Init 添加到文件集的文件添加行信息。令牌位置相对于该文件，因此相对于文件集。

##### Example

```go
package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	// src is the input that we want to tokenize.
	src := []byte("cos(x) + 1i*sin(x) // Euler")

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

}
Output:

1:1	IDENT	"cos"
1:4	(	""
1:5	IDENT	"x"
1:6	)	""
1:8	+	""
1:10	IMAG	"1i"
1:12	*	""
1:13	IDENT	"sin"
1:16	(	""
1:17	IDENT	"x"
1:18	)	""
1:20	COMMENT	"// Euler"
1:28	;	"\n"
```

