+++
title = "base32"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/encoding/base32@go1.21.3](https://pkg.go.dev/encoding/base32@go1.21.3)

Package base32 implements base32 encoding as specified by [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html).

​	base32 包实现了 [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html) 中指定的 base32 编码。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base32/base32.go;l=27)

``` go 
const (
	StdPadding rune = '=' // Standard padding character
	NoPadding  rune = -1  // No padding

)
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base32/base32.go;l=76)

``` go 
var HexEncoding = NewEncoding(encodeHex)
```

HexEncoding is the "Extended Hex Alphabet" defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html). It is typically used in DNS.

​	HexEncoding 是 [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html) 中定义的“扩展十六进制字母”。它通常用于 DNS。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base32/base32.go;l=72)

``` go 
var StdEncoding = NewEncoding(encodeStd)
```

StdEncoding is the standard base32 encoding, as defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html).

​	StdEncoding 是标准 base32 编码，如 [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html) 中定义的。

## 函数

### func NewDecoder 

``` go 
func NewDecoder(enc *Encoding, r io.Reader) io.Reader
```

NewDecoder constructs a new base32 stream decoder.

​	NewDecoder 构造一个新的 base32 流解码器。

### func NewEncoder

```go
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
```

NewEncoder returns a new base32 stream encoder. Data written to the returned writer will be encoded using enc and then written to w. Base32 encodings operate in 5-byte blocks; when finished writing, the caller must Close the returned encoder to flush any partially written blocks.

​	NewEncoder 返回一个新的 base32 流编码器。写入返回的编写器的数据将使用 enc 编码，然后写入 w。Base32 编码以 5 字节块操作；写完后，调用者必须关闭返回的编码器以刷新任何部分写入的块。

#### NewEncoder Example

```go
package main

import (
	"encoding/base32"
	"os"
)

func main() {
	input := []byte("foo\x00bar")
	encoder := base32.NewEncoder(base32.StdEncoding, os.Stdout)
	encoder.Write(input)
	// Must close the encoder when finished to flush any partial blocks.
	// If you comment out the following line, the last partial block "r"
	// won't be encoded.
	encoder.Close()
}

Output:

MZXW6ADCMFZA====
```

## 类型

### type CorruptInputError

```go
type CorruptInputError int64
```

#### (CorruptInputError) Error

```go
func (e CorruptInputError) Error() string
```

### type Encoding

```go
type Encoding struct {
	// contains filtered or unexported fields
}
```

An Encoding is a radix 32 encoding/decoding scheme, defined by a 32-character alphabet. The most common is the “base32” encoding introduced for SASL GSSAPI and standardized in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html). The alternate “base32hex” encoding is used in DNSSEC.

​	Encoding 是一个基数为 32 的编码/解码方案，由一个 32 个字符的字母表定义。最常见的是为 SASL GSSAPI 引入并已在 RFC 4648 中标准化的“base32”编码。备用的“base32hex”编码用于 DNSSEC。

#### func NewEncoding

```go
func NewEncoding(encoder string) *Encoding
```

NewEncoding returns a new Encoding defined by the given alphabet, which must be a 32-byte string.

​	NewEncoding 返回由给定字母表定义的新 Encoding，该字母表必须是一个 32 字节的字符串。

#### (*Encoding) AppendDecode <- go1.22.0

```
func (enc *Encoding) AppendDecode(dst, src []byte) ([]byte, error)
```

AppendDecode appends the base32 decoded src to dst and returns the extended buffer. If the input is malformed, it returns the partially decoded src and an error.

​	AppendDecode 将 base32 解码后的 src 追加到 dst，并返回扩展后的缓冲区。如果输入格式不正确，它会返回部分解码的 src 和一个错误。

#### (*Encoding) AppendEncode <- go1.22.0

```
func (enc *Encoding) AppendEncode(dst, src []byte) []byte
```

AppendEncode appends the base32 encoded src to dst and returns the extended buffer.

​	AppendEncode 将 base32 编码后的 src 追加到 dst，并返回扩展后的缓冲区。

#### (*Encoding) Decode

```go
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)
```

Decode decodes src using the encoding enc. It writes at most DecodedLen(len(src)) bytes to dst and returns the number of bytes written. If src contains invalid base32 data, it will return the number of bytes successfully written and CorruptInputError. New line characters (\r and \n) are ignored.

​	Decode 使用编码 enc 解码 src。它最多将 DecodedLen(len(src)) 字节写入 dst，并返回写入的字节数。如果 src 包含无效的 base32 数据，它将返回成功写入的字节数和 CorruptInputError。忽略换行符（\r 和 \n）。

##### Decode Example

```go
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	str := "JBSWY3DPFQQHO33SNRSCC==="
	dst := make([]byte, base32.StdEncoding.DecodedLen(len(str)))
	n, err := base32.StdEncoding.Decode(dst, []byte(str))
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	dst = dst[:n]
	fmt.Printf("%q\n", dst)
}

Output:

"Hello, world!"
```

#### (*Encoding) DecodeString

```go
func (enc *Encoding) DecodeString(s string) ([]byte, error)
```

DecodeString returns the bytes represented by the base32 string s.

​	DecodeString 返回由 base32 字符串 s 表示的字节。

##### DecodeString Example

```go
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	str := "ONXW2ZJAMRQXIYJAO5UXI2BAAAQGC3TEEDX3XPY="
	data, err := base32.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%q\n", data)
}

Output:

"some data with \x00 and \ufeff"
```

#### (*Encoding) DecodedLen

```go
func (enc *Encoding) DecodedLen(n int) int
```

DecodedLen returns the maximum length in bytes of the decoded data corresponding to n bytes of base32-encoded data.

​	DecodedLen 返回与 n 个字节的 base32 编码数据对应的解码数据的最大长度（以字节为单位）。

#### (*Encoding) Encode

```go
func (enc *Encoding) Encode(dst, src []byte)
```

Encode encodes src using the encoding enc, writing EncodedLen(len(src)) bytes to dst.

​	Encode 使用编码 enc 对 src 进行编码，将 EncodedLen(len(src)) 字节写入 dst。

The encoding pads the output to a multiple of 8 bytes, so Encode is not appropriate for use on individual blocks of a large data stream. Use NewEncoder() instead.

​	编码将输出填充为 8 个字节的倍数，因此 Encode 不适合用于大型数据流的各个块。请改用 NewEncoder()。

##### Encode Example

```go
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	data := []byte("Hello, world!")
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, data)
	fmt.Println(string(dst))
}
```

#### (*Encoding) EncodeToString

```go
func (enc *Encoding) EncodeToString(src []byte) string
```

EncodeToString returns the base32 encoding of src.

​	EncodeToString 返回 src 的 base32 编码。

##### EncodeToString Example

```go
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	data := []byte("any + old & data")
	str := base32.StdEncoding.EncodeToString(data)
	fmt.Println(str)
}

Output:

MFXHSIBLEBXWYZBAEYQGIYLUME======
```

#### (*Encoding) EncodedLen

```go
func (enc *Encoding) EncodedLen(n int) int
```

EncodedLen returns the length in bytes of the base32 encoding of an input buffer of length n.

​	EncodedLen 返回长度为 n 的输入缓冲区的 base32 编码的字节长度。

#### (Encoding) WithPadding <- go1.9

```go
func (enc Encoding) WithPadding(padding rune) *Encoding
```

WithPadding creates a new encoding identical to enc except with a specified padding character, or NoPadding to disable padding. The padding character must not be ‘\r’ or ‘\n’, must not be contained in the encoding’s alphabet and must be a rune equal or below ‘\xff’.

​	WithPadding 创建一个与 enc 相同的新编码，但使用指定的填充字符，或 NoPadding 来禁用填充。填充字符不能是 ‘\r’ 或 ‘\n’，不能包含在编码的字母表中，并且必须是小于或等于 ‘\xff’ 的符文。