+++
title = "lzw"
date = 2023-05-17T09:59:21+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/compress/lzw@go1.23.0](https://pkg.go.dev/compress/lzw@go1.23.0)

Package lzw implements the Lempel-Ziv-Welch compressed data format, described in T. A. Welch, "A Technique for High-Performance Data Compression", Computer, 17(6) (June 1984), pp 8-19.

​	lzw 包实现 Lempel-Ziv-Welch 压缩数据格式，在 T. A. Welch，“一种高性能数据压缩技术”，计算机，17(6)（1984 年 6 月），第 8-19 页中进行了描述。

In particular, it implements LZW as used by the GIF and PDF file formats, which means variable-width codes up to 12 bits and the first two non-literal codes are a clear code and an EOF code.

​	具体来说，它实现了 GIF 和 PDF 文件格式所使用的 LZW，这意味着可变宽度的代码最多为 12 位，并且前两个非文本代码是清除代码和 EOF 代码。

The TIFF file format uses a similar but incompatible version of the LZW algorithm. See the golang.org/x/image/tiff/lzw package for an implementation.

​	TIFF 文件格式使用类似但又不兼容的 LZW 算法版本。有关实现，请参阅 golang.org/x/image/tiff/lzw 包。



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func NewReader 

``` go 
func NewReader(r io.Reader, order Order, litWidth int) io.ReadCloser
```

NewReader creates a new io.ReadCloser. Reads from the returned io.ReadCloser read and decompress data from r. If r does not also implement io.ByteReader, the decompressor may read more data than necessary from r. It is the caller's responsibility to call Close on the ReadCloser when finished reading. The number of bits to use for literal codes, litWidth, must be in the range [2,8] and is typically 8. It must equal the litWidth used during compression.

​	NewReader 创建一个新的 io.ReadCloser。从返回的 io.ReadCloser 读取并解压缩来自 r 的数据。如果 r 也没有实现 io.ByteReader，则解压缩器可能会从 r 读取比必要更多的 data。在完成读取后，由调用者负责对 ReadCloser 调用 Close。用于文本代码的位数 litWidth 必须在 [2,8] 范围内，通常为 8。它必须等于压缩期间使用的 litWidth。

It is guaranteed that the underlying type of the returned io.ReadCloser is a *Reader.

​	保证返回的 io.ReadCloser 的底层类型为 *Reader。

### func NewWriter 

``` go 
func NewWriter(w io.Writer, order Order, litWidth int) io.WriteCloser
```

NewWriter creates a new io.WriteCloser. Writes to the returned io.WriteCloser are compressed and written to w. It is the caller's responsibility to call Close on the WriteCloser when finished writing. The number of bits to use for literal codes, litWidth, must be in the range [2,8] and is typically 8. Input bytes must be less than 1<<litWidth.

​	NewWriter 创建一个新的 io.WriteCloser。写入返回的 io.WriteCloser 的内容会被压缩并写入 w。调用者负责在完成写入后对 WriteCloser 调用 Close。用于文字代码的位数 litWidth 必须在 [2,8] 范围内，通常为 8。输入字节必须小于 1«litWidth。

It is guaranteed that the underlying type of the returned io.WriteCloser is a *Writer.

​	保证返回的 io.WriteCloser 的底层类型为 *Writer。

## 类型

### type Order 

``` go 
type Order int
```

Order specifies the bit ordering in an LZW data stream.

​	Order 指定 LZW 数据流中的位顺序。

``` go 
const (
	// LSB means Least Significant Bits first, as used in the GIF file format.
	LSB Order = iota
	// MSB means Most Significant Bits first, as used in the TIFF and PDF
	// file formats.
	MSB
)
```

### type Reader  <- go1.17

``` go 
type Reader struct {
	// contains filtered or unexported fields
}
```

Reader is an io.Reader which can be used to read compressed data in the LZW format.

​	Reader 是一个 io.Reader，可用于读取 LZW 格式的压缩数据。

#### (*Reader) Close  <- go1.17

``` go 
func (r *Reader) Close() error
```

Close closes the Reader and returns an error for any future read operation. It does not close the underlying io.Reader.

​	Close 关闭 Reader 并返回未来任何读取操作的错误。它不会关闭底层的 io.Reader。

#### (*Reader) Read  <- go1.17

``` go 
func (r *Reader) Read(b []byte) (int, error)
```

Read implements io.Reader, reading uncompressed bytes from its underlying Reader.

​	Read 实现 io.Reader，从其底层 Reader 读取未压缩的字节。

#### (*Reader) Reset  <- go1.17

``` go 
func (r *Reader) Reset(src io.Reader, order Order, litWidth int)
```

Reset clears the Reader's state and allows it to be reused again as a new Reader.

​	Reset 清除 Reader 的状态，并允许将其再次用作新的 Reader。

### type Writer  <- go1.17

``` go 
type Writer struct {
	// contains filtered or unexported fields
}
```

Writer is an LZW compressor. It writes the compressed form of the data to an underlying writer (see NewWriter).

​	Writer 是一个 LZW 压缩器。它将数据的压缩形式写入基础 writer（请参阅 NewWriter）。

#### (*Writer) Close  <- go1.17

``` go 
func (w *Writer) Close() error
```

Close closes the Writer, flushing any pending output. It does not close w's underlying writer.

​	Close 关闭 Writer，刷新所有待处理的输出。它不会关闭 w 的基础 writer。

#### (*Writer) Reset  <- go1.17

``` go 
func (w *Writer) Reset(dst io.Writer, order Order, litWidth int)
```

Reset clears the Writer's state and allows it to be reused again as a new Writer.

​	Reset 清除 Writer 的状态，并允许将其再次用作新的 Writer。

#### (*Writer) Write  <- go1.17

``` go 
func (w *Writer) Write(p []byte) (n int, err error)
```

Write writes a compressed representation of p to w's underlying writer.

​	Write 将 p 的压缩表示写入 w 的基础 writer。