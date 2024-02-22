+++
title = "quotedprintable"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/mime/quotedprintable@go1.21.3](https://pkg.go.dev/mime/quotedprintable@go1.21.3)

Package quotedprintable implements quoted-printable encoding as specified by [RFC 2045](https://rfc-editor.org/rfc/rfc2045.html).

​	Package quotedprintable 实现 RFC 2045 中指定的 quoted-printable 编码。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Reader

```go
type Reader struct {
	// contains filtered or unexported fields
}
```

Reader is a quoted-printable decoder.

​	Reader 是一个 quoted-printable 解码器。

#### func NewReader

```go
func NewReader(r io.Reader) *Reader
```

NewReader returns a quoted-printable reader, decoding from r.

​	NewReader 返回一个 quoted-printable 读取器，从 r 解码。

##### NewReader Example 

```go
package main

import (
	"fmt"
	"io"
	"mime/quotedprintable"
	"strings"
)

func main() {
	for _, s := range []string{
		`=48=65=6C=6C=6F=2C=20=47=6F=70=68=65=72=73=21`,
		`invalid escape: <b style="font-size: 200%">hello</b>`,
		"Hello, Gophers! This symbol will be unescaped: =3D and this will be written in =\r\none line.",
	} {
		b, err := io.ReadAll(quotedprintable.NewReader(strings.NewReader(s)))
		fmt.Printf("%s %v\n", b, err)
	}
}
Output:

Hello, Gophers! <nil>
invalid escape: <b style="font-size: 200%">hello</b> <nil>
Hello, Gophers! This symbol will be unescaped: = and this will be written in one line. <nil>
```

#### (*Reader) Read

```go
func (r *Reader) Read(p []byte) (n int, err error)
```

Read reads and decodes quoted-printable data from the underlying reader.

​	Read 从底层读取器读取并解码 quoted-printable 数据。

### type Writer

```go
type Writer struct {
	// Binary mode treats the writer's input as pure binary and processes end of
	// line bytes as binary data.
	Binary bool
	// contains filtered or unexported fields
}
```

A Writer is a quoted-printable writer that implements io.WriteCloser.

​	Writer 是一个实现 io.WriteCloser 的 quoted-printable 写入器。

#### func NewWriter

```go
func NewWriter(w io.Writer) *Writer
```

NewWriter returns a new Writer that writes to w.

​	NewWriter 返回一个新的 Writer，写入 w。

##### NewWriter Example

```go
package main

import (
	"mime/quotedprintable"
	"os"
)

func main() {
	w := quotedprintable.NewWriter(os.Stdout)
	w.Write([]byte("These symbols will be escaped: = \t"))
	w.Close()

}
Output:

These symbols will be escaped: =3D =09
```

#### (*Writer) Close

```go
func (w *Writer) Close() error
```

Close closes the Writer, flushing any unwritten data to the underlying io.Writer, but does not close the underlying io.Writer.

​	Close 关闭 Writer，将任何未写入的数据刷新到底层 io.Writer，但不关闭底层 io.Writer。

#### (*Writer) Write

```go
func (w *Writer) Write(p []byte) (n int, err error)
```

Write encodes p using quoted-printable encoding and writes it to the underlying io.Writer. It limits line length to 76 characters. The encoded bytes are not necessarily flushed until the Writer is closed.

​	Write 使用 quoted-printable 编码对 p 进行编码，并将其写入底层 io.Writer。它将行长限制为 76 个字符。在关闭 Writer 之前，不会强制刷新已编码的字节。