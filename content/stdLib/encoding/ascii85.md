+++
title = "ascii85"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/encoding/ascii85@go1.23.0](https://pkg.go.dev/encoding/ascii85@go1.23.0)

Package ascii85 implements the ascii85 data encoding as used in the btoa tool and Adobe's PostScript and PDF document formats.

​	ascii85 包实现了 btoa 工具和 Adobe 的 PostScript 和 PDF 文档格式中使用的 ascii85 数据编码。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Decode 

``` go 
func Decode(dst, src []byte, flush bool) (ndst, nsrc int, err error)
```

Decode decodes src into dst, returning both the number of bytes written to dst and the number consumed from src. If src contains invalid ascii85 data, Decode will return the number of bytes successfully written and a CorruptInputError. Decode ignores space and control characters in src. Often, ascii85-encoded data is wrapped in <~ and ~> symbols. Decode expects these to have been stripped by the caller.


​	Decode 将 src 解码到 dst 中，同时返回写入 dst 的字节数和从 src 中消耗的字节数。如果 src 包含无效的 ascii85 数据，Decode 将返回成功写入的字节数和 CorruptInputError。Decode 会忽略 src 中的空间和控制字符。通常，ascii85 编码的数据会用 <~ 和 ~> 符号包装。Decode 预期调用者已将其剥离。

If flush is true, Decode assumes that src represents the end of the input stream and processes it completely rather than wait for the completion of another 32-bit block.

​	如果 flush 为 true，则 Decode 假设 src 表示输入流的末尾，并完全处理它，而不是等待另一个 32 位块完成。

NewDecoder wraps an io.Reader interface around Decode.

​	NewDecoder 将 io.Reader 接口包装在 Decode 周围。

### func Encode

```go
func Encode(dst, src []byte) int
```

Encode encodes src into at most MaxEncodedLen(len(src)) bytes of dst, returning the actual number of bytes written.

​	Encode 将 src 编码为最多 MaxEncodedLen(len(src)) 个 dst 字节，并返回实际写入的字节数。

The encoding handles 4-byte chunks, using a special encoding for the last fragment, so Encode is not appropriate for use on individual blocks of a large data stream. Use NewEncoder() instead.

​	编码处理 4 字节块，对最后一个片段使用特殊编码，因此 Encode 不适合用于大型数据流的各个块。请改用 NewEncoder()。

Often, ascii85-encoded data is wrapped in <~ and ~> symbols. Encode does not add these.

​	通常，ascii85 编码的数据用 <~ 和 ~> 符号包装。Encode 不会添加这些符号。

### func MaxEncodedLen

```go
func MaxEncodedLen(n int) int
```

MaxEncodedLen returns the maximum length of an encoding of n source bytes.

​	MaxEncodedLen 返回 n 个源字节的编码的最大长度。

### func NewDecoder

```go
func NewDecoder(r io.Reader) io.Reader
```

NewDecoder constructs a new ascii85 stream decoder.

​	NewDecoder 构造一个新的 ascii85 流解码器。

### func NewEncoder

```go
func NewEncoder(w io.Writer) io.WriteCloser
```

NewEncoder returns a new ascii85 stream encoder. Data written to the returned writer will be encoded and then written to w. Ascii85 encodings operate in 32-bit blocks; when finished writing, the caller must Close the returned encoder to flush any trailing partial block.

​	NewEncoder 返回一个新的 ascii85 流编码器。写入返回的编写器的数据将被编码，然后写入 w。Ascii85 编码以 32 位块操作；写完后，调用者必须关闭返回的编码器以刷新任何尾随的部分块。

## 类型

### type CorruptInputError 

``` go 
type CorruptInputError int64
```

#### (CorruptInputError) Error 

``` go 
func (e CorruptInputError) Error() string
```