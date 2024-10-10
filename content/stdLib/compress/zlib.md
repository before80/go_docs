+++
title = "zlib"
date = 2023-05-17T09:59:21+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/compress/zlib@go1.23.0](https://pkg.go.dev/compress/zlib@go1.23.0)

Package zlib implements reading and writing of zlib format compressed data, as specified in [RFC 1950](https://rfc-editor.org/rfc/rfc1950.html).

​	zlib 包实现了对 [RFC 1950](https://rfc-editor.org/rfc/rfc1950.html) 中指定的 zlib 格式压缩数据的读写。

The implementation provides filters that uncompress during reading and compress during writing. For example, to write compressed data to a buffer:

​	该实现提供了在读取期间取消压缩并在写入期间压缩的过滤器。例如，要将压缩数据写入缓冲区：

``` go 
var b bytes.Buffer
w := zlib.NewWriter(&b)
w.Write([]byte("hello, world\n"))
w.Close()
```

and to read that data back:

并读取该数据：

```go
r, err := zlib.NewReader(&b)
io.Copy(os.Stdout, r)
r.Close()
```


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/compress/zlib/writer.go;l=18)

``` go 
const (
	NoCompression      = flate.NoCompression
	BestSpeed          = flate.BestSpeed
	BestCompression    = flate.BestCompression
	DefaultCompression = flate.DefaultCompression
	HuffmanOnly        = flate.HuffmanOnly
)
```

These constants are copied from the flate package, so that code that imports "compress/zlib" does not also have to import "compress/flate".

​	这些常量从 flate 包中复制，以便导入“compress/zlib”的代码不必导入“compress/flate”。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/compress/zlib/reader.go;l=40)

``` go 
var (
	// ErrChecksum is returned when reading ZLIB data that has an invalid checksum.
	ErrChecksum = errors.New("zlib: invalid checksum")
	// ErrDictionary is returned when reading ZLIB data that has an invalid dictionary.
	ErrDictionary = errors.New("zlib: invalid dictionary")
	// ErrHeader is returned when reading ZLIB data that has an invalid header.
	ErrHeader = errors.New("zlib: invalid header")
)
```

## 函数

### func NewReader 

``` go 
func NewReader(r io.Reader) (io.ReadCloser, error)
```

NewReader creates a new ReadCloser. Reads from the returned ReadCloser read and decompress data from r. If r does not implement io.ByteReader, the decompressor may read more data than necessary from r. It is the caller's responsibility to call Close on the ReadCloser when done.

​	NewReader 创建一个新的 ReadCloser。从返回的 ReadCloser 读取并从 r 解压缩数据。如果 r 未实现 io.ByteReader，则解压缩器可能会从 r 读取多余的数据。在完成后，由调用者负责对 ReadCloser 调用 Close。

The ReadCloser returned by NewReader also implements Resetter.

​	NewReader 返回的 ReadCloser 也实现了 Resetter。

#### NewReader Example
``` go 
package main

import (
	"bytes"
	"compress/zlib"
	"io"
	"os"
)

func main() {
	buff := []byte{120, 156, 202, 72, 205, 201, 201, 215, 81, 40, 207,
		47, 202, 73, 225, 2, 4, 0, 0, 255, 255, 33, 231, 4, 147}
	b := bytes.NewReader(buff)

	r, err := zlib.NewReader(b)
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, r)

	r.Close()
}
Output:

hello, world
```

### func NewReaderDict 

``` go 
func NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)
```

NewReaderDict is like NewReader but uses a preset dictionary. NewReaderDict ignores the dictionary if the compressed data does not refer to it. If the compressed data refers to a different dictionary, NewReaderDict returns ErrDictionary.

​	NewReaderDict 与 NewReader 类似，但使用预设词典。如果压缩数据未引用词典，NewReaderDict 将忽略该词典。如果压缩数据引用了不同的词典，NewReaderDict 将返回 ErrDictionary。

The ReadCloser returned by NewReaderDict also implements Resetter.

​	NewReaderDict 返回的 ReadCloser 也实现了 Resetter。

## 类型

### type Resetter  <- go1.4

``` go 
type Resetter interface {
	// Reset discards any buffered data and resets the Resetter as if it was
	// newly initialized with the given reader.
	Reset(r io.Reader, dict []byte) error
}
```

Resetter resets a ReadCloser returned by NewReader or NewReaderDict to switch to a new underlying Reader. This permits reusing a ReadCloser instead of allocating a new one.

​	Resetter 重置 NewReader 或 NewReaderDict 返回的 ReadCloser 以切换到新的底层 Reader。这允许重用 ReadCloser，而不是分配一个新的。

### type Writer 

``` go 
type Writer struct {
	// contains filtered or unexported fields
}
```

A Writer takes data written to it and writes the compressed form of that data to an underlying writer (see NewWriter).

​	Writer 接收写入其中的数据，并将该数据的压缩形式写入底层 writer（请参阅 NewWriter）。

#### func NewWriter 

``` go 
func NewWriter(w io.Writer) *Writer
```

NewWriter creates a new Writer. Writes to the returned Writer are compressed and written to w.

​	NewWriter 创建一个新的 Writer。对返回的 Writer 的写入将被压缩并写入 w。

It is the caller's responsibility to call Close on the Writer when done. Writes may be buffered and not flushed until Close.

​	在完成时，由调用者负责对 Writer 调用 Close。写入可能会被缓冲，直到 Close 才刷新。

##### NewWriter Example
``` go 
package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
)

func main() {
	var b bytes.Buffer

	w := zlib.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()
	fmt.Println(b.Bytes())
}
Output:

[120 156 202 72 205 201 201 215 81 40 207 47 202 73 225 2 4 0 0 255 255 33 231 4 147]
```

#### func NewWriterLevel 

``` go 
func NewWriterLevel(w io.Writer, level int) (*Writer, error)
```

NewWriterLevel is like NewWriter but specifies the compression level instead of assuming DefaultCompression.

​	NewWriterLevel 类似于 NewWriter，但指定压缩级别而不是假定 DefaultCompression。

The compression level can be DefaultCompression, NoCompression, HuffmanOnly or any integer value between BestSpeed and BestCompression inclusive. The error returned will be nil if the level is valid.

​	压缩级别可以是 DefaultCompression、NoCompression、HuffmanOnly 或介于 BestSpeed 和 BestCompression（包括）之间的任何整数值。如果级别有效，则返回的错误将为 nil。

#### func NewWriterLevelDict 

``` go 
func NewWriterLevelDict(w io.Writer, level int, dict []byte) (*Writer, error)
```

NewWriterLevelDict is like NewWriterLevel but specifies a dictionary to compress with.

​	NewWriterLevelDict 与 NewWriterLevel 类似，但指定了一个词典来进行压缩。

The dictionary may be nil. If not, its contents should not be modified until the Writer is closed.

​	词典可以为 nil。如果不是，则在 Writer 关闭之前不应修改其内容。

#### (*Writer) Close 

``` go 
func (z *Writer) Close() error
```

Close closes the Writer, flushing any unwritten data to the underlying io.Writer, but does not close the underlying io.Writer.

​	Close 关闭 Writer，将所有未写入的数据刷新到底层 io.Writer，但不关闭底层 io.Writer。

#### (*Writer) Flush 

``` go 
func (z *Writer) Flush() error
```

Flush flushes the Writer to its underlying io.Writer.

​	Flush 将 Writer 刷新到其底层的 io.Writer。

#### (*Writer) Reset  <- go1.2

``` go 
func (z *Writer) Reset(w io.Writer)
```

Reset clears the state of the Writer z such that it is equivalent to its initial state from NewWriterLevel or NewWriterLevelDict, but instead writing to w.

​	Reset 清除 Writer z 的状态，使其等同于 NewWriterLevel 或 NewWriterLevelDict 的初始状态，但改为写入 w。

#### (*Writer) Write 

``` go 
func (z *Writer) Write(p []byte) (n int, err error)
```

Write writes a compressed form of p to the underlying io.Writer. The compressed bytes are not necessarily flushed until the Writer is closed or explicitly flushed.

​	Write 将 p 的压缩形式写入底层 io.Writer。在 Writer 关闭或显式刷新之前，压缩字节不一定被刷新。