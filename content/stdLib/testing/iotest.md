+++
title = "iotest"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/testing/iotest@go1.20.1

包 iotest 实现了一些用于测试的 Reader 和 Writer。

## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/testing/iotest/reader.go;l=74)

``` go 
var ErrTimeout = errors.New("timeout")
```

​	ErrTimeout 是一个虚假的超时错误。

## 函数

#### func DataErrReader 

``` go 
func DataErrReader(r io.Reader) io.Reader
```

​	DataErrReader函数更改 Reader 处理错误的方式。通常，Reader 在读取最后一块数据后的第一次读取调用中返回一个错误(通常是 EOF)。DataErrReader 包装一个 Reader 并更改其行为，以便最终的错误与最终数据一起返回，而不是在最终数据后的第一次调用中返回。

#### func ErrReader  <- go1.16

``` go 
func ErrReader(err error) io.Reader
```

​	ErrReader函数返回一个 io.Reader，该 Reader 的所有 Read 调用都会返回 0，err。

##### ErrReader Example
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

#### func HalfReader 

``` go 
func HalfReader(r io.Reader) io.Reader
```

​	HalfReader函数返回一个 Reader，它通过从 `r` 读取一半所需的字节数来实现 Read。

#### func NewReadLogger 

``` go 
func NewReadLogger(prefix string, r io.Reader) io.Reader
```

​	NewReadLogger函数返回一个 Reader，它的行为类似于 `r`，但记录(使用 log.Printf)每次读取到标准错误输出，打印前缀和十六进制读取的数据。

#### func NewWriteLogger 

``` go 
func NewWriteLogger(prefix string, w io.Writer) io.Writer
```

​	NewWriteLogger函数返回一个 Writer，它的行为类似于 `w`，但记录(使用 log.Printf)每次写入到标准错误输出，打印前缀和十六进制写入的数据。

#### func OneByteReader 

``` go 
func OneByteReader(r io.Reader) io.Reader
```

​	OneByteReader函数返回一个 Reader，它通过从 `r` 读取一个字节来实现每个非空的 Read。

#### func TestReader  <- go1.16

``` go 
func TestReader(r io.Reader, content []byte) error
```

​	TestReade函数测试从 `r` 读取数据是否返回预期的文件内容。它对不同大小的读取进行了测试，直到 EOF。如果 `r` 实现了 io.ReaderAt 方法或 io.Seeker方法，则 TestReader函数还检查这些操作的行为是否符合预期。

​	如果 TestReader函数发现任何行为不当，它将返回报告这些问题的错误。错误文本可能跨越多行。

#### func TimeoutReader 

``` go 
func TimeoutReader(r io.Reader) io.Reader
```

​	TimeoutReader函数在第二次无数据的读取上返回 ErrTimeout。随后的读取调用成功。

#### func TruncateWriter 

``` go 
func TruncateWriter(w io.Writer, n int64) io.Writer
```

​	TruncateWriter函数返回一个 Writer，它写入 `w`，但在写入 n 个字节后静默停止。

## 类型

This section is empty.