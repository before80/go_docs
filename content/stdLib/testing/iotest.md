+++
title = "iotest"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/testing/iotest@go1.21.3](https://pkg.go.dev/testing/iotest@go1.21.3)

Package iotest implements Readers and Writers useful mainly for testing.

​	`iotest` 包实现了一些用于测试的 `Reader` 和 `Writer`。

## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/iotest/reader.go;l=74)

``` go 
var ErrTimeout = errors.New("timeout")
```

ErrTimeout is a fake timeout error.

​	`ErrTimeout` 是一个虚假的超时错误。

## 函数

### func DataErrReader 

``` go 
func DataErrReader(r io.Reader) io.Reader
```

DataErrReader changes the way errors are handled by a Reader. Normally, a Reader returns an error (typically EOF) from the first Read call after the last piece of data is read. DataErrReader wraps a Reader and changes its behavior so the final error is returned along with the final data, instead of in the first call after the final data.

​	`DataErrReader`函数更改 Reader 处理错误的方式。通常，Reader 在读取最后一块数据后的第一次读取调用中返回一个错误(通常是 EOF)。DataErrReader 包装一个 Reader 并更改其行为，以便最终的错误与最终数据一起返回，而不是在最终数据后的第一次调用中返回。

### func ErrReader  <- go1.16

``` go 
func ErrReader(err error) io.Reader
```

ErrReader returns an io.Reader that returns 0, err from all Read calls.

​	`ErrReader`函数返回一个 io.Reader，该 Reader 的所有 Read 调用都会返回 0，err。

#### ErrReader Example
``` go 
package main

import (
	"errors"
	"fmt"
	"testing/iotest"
)

func main() {
	// A reader that always returns a custom error.
	r := iotest.ErrReader(errors.New("custom error"))
	n, err := r.Read(nil)
	fmt.Printf("n:   %d\nerr: %q\n", n, err)

}
Output:

n:   0
err: "custom error"
```

### func HalfReader 

``` go 
func HalfReader(r io.Reader) io.Reader
```

HalfReader returns a Reader that implements Read by reading half as many requested bytes from r.

​	`HalfReader`函数返回一个 Reader，它通过从 `r` 读取一半所需的字节数来实现 Read。

### func NewReadLogger 

``` go 
func NewReadLogger(prefix string, r io.Reader) io.Reader
```

NewReadLogger returns a reader that behaves like r except that it logs (using log.Printf) each read to standard error, printing the prefix and the hexadecimal data read.

​	`NewReadLogger`函数返回一个 Reader，它的行为类似于 `r`，但记录(使用 log.Printf)每次读取到标准错误输出，打印前缀和十六进制读取的数据。

### func NewWriteLogger 

``` go 
func NewWriteLogger(prefix string, w io.Writer) io.Writer
```

NewWriteLogger returns a writer that behaves like w except that it logs (using log.Printf) each write to standard error, printing the prefix and the hexadecimal data written.

​	`NewWriteLogger`函数返回一个 Writer，它的行为类似于 `w`，但记录(使用 log.Printf)每次写入到标准错误输出，打印前缀和十六进制写入的数据。

### func OneByteReader 

``` go 
func OneByteReader(r io.Reader) io.Reader
```

OneByteReader returns a Reader that implements each non-empty Read by reading one byte from r.

​	`OneByteReader`函数返回一个 Reader，它通过从 `r` 读取一个字节来实现每个非空的 Read。

### func TestReader  <- go1.16

``` go 
func TestReader(r io.Reader, content []byte) error
```

TestReader tests that reading from r returns the expected file content. It does reads of different sizes, until EOF. If r implements io.ReaderAt or io.Seeker, TestReader also checks that those operations behave as they should.

​	`TestReade`函数测试从 `r` 读取数据是否返回预期的文件内容。它对不同大小的读取进行了测试，直到 EOF。如果 `r` 实现了 io.ReaderAt 方法或 io.Seeker方法，则 TestReader函数还检查这些操作的行为是否符合预期。

If TestReader finds any misbehaviors, it returns an error reporting them. The error text may span multiple lines.

​	如果 `TestReader`函数发现任何行为不当，它将返回报告这些问题的错误。错误文本可能跨越多行。

### func TimeoutReader 

``` go 
func TimeoutReader(r io.Reader) io.Reader
```

TimeoutReader returns ErrTimeout on the second read with no data. Subsequent calls to read succeed.

​	`TimeoutReader`函数在第二次无数据的读取上返回 ErrTimeout。随后的读取调用成功。

### func TruncateWriter 

``` go 
func TruncateWriter(w io.Writer, n int64) io.Writer
```

TruncateWriter returns a Writer that writes to w but stops silently after n bytes.

​	`TruncateWriter`函数返回一个 Writer，它写入 `w`，但在写入 n 个字节后静默停止。

## 类型

This section is empty.