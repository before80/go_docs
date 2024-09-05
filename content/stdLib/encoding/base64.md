+++
title = "base64"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/encoding/base64@go1.21.3](https://pkg.go.dev/encoding/base64@go1.21.3)

Package base64 implements base64 encoding as specified by [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html).

​	base64包实现了[RFC 4648](https://rfc-editor.org/rfc/rfc4648.html)中规定的base64编码。

## Example
``` go 
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := "Hello, 世界"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))
}

Output:

SGVsbG8sIOS4lueVjA==
Hello, 世界
```


## 常量 

### StdPadding

### NoPadding

``` go 
const (
	StdPadding rune = '=' // Standard padding character  标准填充字符
	NoPadding  rune = -1  // No padding 无填充字符

)
```

## 变量

### RawStdEncoding

``` go 
var RawStdEncoding = StdEncoding.WithPadding(NoPadding)
```

RawStdEncoding is the standard raw, unpadded base64 encoding, as defined in [RFC 4648 section 3.2](https://rfc-editor.org/rfc/rfc4648.html#section-3.2). This is the same as StdEncoding but omits padding characters.

​	RawStdEncoding变量是标准的原始、无填充的base64编码，如[RFC 4648第3.2节](https://rfc-editor.org/rfc/rfc4648.html#section-3.2)中所定义。它与StdEncoding相同，但省略了填充字符。

### RawURLEncoding

``` go 
var RawURLEncoding = URLEncoding.WithPadding(NoPadding)
```

RawURLEncoding is the unpadded alternate base64 encoding defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html). It is typically used in URLs and file names. This is the same as URLEncoding but omits padding characters.

​	RawURLEncoding变量是[RFC 4648](https://rfc-editor.org/rfc/rfc4648.html)中定义的无填充的另一种base64编码。它通常用于URL和文件名。它与URLEncoding相同，但省略了填充字符。

### StdEncoding

``` go 
var StdEncoding = NewEncoding(encodeStd)
```

StdEncoding is the standard base64 encoding, as defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html).

​	StdEncoding变量是标准的base64编码，如[RFC 4648](https://rfc-editor.org/rfc/rfc4648.html)中所定义。

### URLEncoding

``` go 
var URLEncoding = NewEncoding(encodeURL)
```

URLEncoding is the alternate base64 encoding defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html). It is typically used in URLs and file names.

​	URLEncoding变量是[RFC 4648](https://rfc-editor.org/rfc/rfc4648.html)中定义的另一种base64编码。它通常用于URL和文件名。

## 函数

### func NewDecoder 

``` go 
func NewDecoder(enc *Encoding, r io.Reader) io.Reader
```

NewDecoder constructs a new base64 stream decoder.

​	NewDecoder函数构造一个新的base64流解码器。

### func NewEncoder 

``` go 
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
```

NewEncoder returns a new base64 stream encoder. Data written to the returned writer will be encoded using enc and then written to w. Base64 encodings operate in 4-byte blocks; when finished writing, the caller must Close the returned encoder to flush any partially written blocks.

​	NewEncoder返回一个新的base64流编码器。写入返回的写入器的数据将使用enc进行编码，然后写入w。Base64编码以4字节的块进行操作；当写完后，调用者必须关闭返回的编码器以冲刷任何部分写入的块。

​	NewEncoder函数返回一个新的base64流编码器。写入返回的写入器的数据将使用enc进行编码，然后写入w。Base64编码操作以4字节块为单位；在写入完成后，调用方必须关闭返回的编码器以刷新任何部分写入的块。

#### NewEncoder Example
``` go 
package main

import (
	"encoding/base64"
	"os"
)

func main() {
	input := []byte("foo\x00bar")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	// Must close the encoder when finished to flush any partial blocks.
	// If you comment out the following line, the last partial block "r"
	// won't be encoded.
	encoder.Close()
}

Output:

Zm9vAGJhcg==
```

## 类型

### type CorruptInputError 

``` go 
type CorruptInputError int64
```

#### (CorruptInputError) Error 

``` go 
func (e CorruptInputError) Error() string
```

### type Encoding 

``` go 
type Encoding struct {
	// contains filtered or unexported fields
}
```

An Encoding is a radix 64 encoding/decoding scheme, defined by a 64-character alphabet. The most common encoding is the "base64" encoding defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html) and used in MIME ([RFC 2045](https://rfc-editor.org/rfc/rfc2045.html)) and PEM ([RFC 1421](https://rfc-editor.org/rfc/rfc1421.html)). [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html) also defines an alternate encoding, which is the standard encoding with - and _ substituted for + and /.

​	Encoding结构体是一个基于64个字符的基数64编码/解码方案（a radix 64 encoding/decoding scheme）。最常见的编码是"base64"编码，如[RFC 4648](https://rfc-editor.org/rfc/rfc4648.html)中所定义，并在MIME ([RFC 2045](https://rfc-editor.org/rfc/rfc2045.html))和PEM ([RFC 1421](https://rfc-editor.org/rfc/rfc1421.html))中使用。[RFC 4648](https://rfc-editor.org/rfc/rfc4648.html)还定义了一种备用编码，它是将标准编码中的`+`和`/`分别替换为`-`和`_`。

#### func NewEncoding 

``` go 
func NewEncoding(encoder string) *Encoding
```

NewEncoding returns a new padded Encoding defined by the given alphabet, which must be a 64-byte string that does not contain the padding character or CR / LF ('\r', '\n'). The resulting Encoding uses the default padding character ('='), which may be changed or disabled via WithPadding.

​	NewEncoding函数返回一个由给定字母定义的新的填充编码，它必须是一个64字节的字符串，不包含填充字符或CR / LF('\r', '\n')。产生的Encoding使用默认的填充字符('=')，可以通过WithPadding进行更改或禁用。

#### (*Encoding) AppendDecode <- go1.22.0

```
func (enc *Encoding) AppendDecode(dst, src []byte) ([]byte, error)
```

AppendDecode appends the base64 decoded src to dst and returns the extended buffer. If the input is malformed, it returns the partially decoded src and an error.

​	AppendDecode 将 base64 解码后的 src 追加到 dst，并返回扩展后的缓冲区。如果输入格式不正确，它会返回部分解码的 src 和一个错误。

#### (*Encoding) AppendEncode <- go1.22.0

```
func (enc *Encoding) AppendEncode(dst, src []byte) []byte
```

AppendEncode appends the base64 encoded src to dst and returns the extended buffer.

​	AppendEncode 将 base64 编码后的 src 追加到 dst，并返回扩展后的缓冲区。

#### (*Encoding) Decode 

``` go 
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)
```

Decode decodes src using the encoding enc. It writes at most DecodedLen(len(src)) bytes to dst and returns the number of bytes written. If src contains invalid base64 data, it will return the number of bytes successfully written and CorruptInputError. New line characters (\r and \n) are ignored.

​	Decode方法使用 the encoding enc 解码`src`。它将最多`DecodedLen(len(src))`个字节写入`dst`，并返回写入的字节数。如果`src`包含无效的base64数据，它将返回成功写入的字节数和CorruptInputError。换行字符（`\r`和`\n`）将被忽略。

##### Decode Example
``` go 
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "SGVsbG8sIHdvcmxkIQ=="
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(str)))
	n, err := base64.StdEncoding.Decode(dst, []byte(str))
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

``` go 
func (enc *Encoding) DecodeString(s string) ([]byte, error)
```

DecodeString returns the bytes represented by the base64 string s.

​	DecodeString方法返回由base64字符串`s`表示的字节。

##### DecodeString Example
``` go 
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "c29tZSBkYXRhIHdpdGggACBhbmQg77u/"
	data, err := base64.StdEncoding.DecodeString(str)
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

``` go 
func (enc *Encoding) DecodedLen(n int) int
```

DecodedLen returns the maximum length in bytes of the decoded data corresponding to n bytes of base64-encoded data.

​	DecodedLen方法返回对应于`n`个字节的base64编码数据的解码数据的最大长度（以字节为单位）。

#### (*Encoding) Encode 

``` go 
func (enc *Encoding) Encode(dst, src []byte)
```

Encode encodes src using the encoding enc, writing EncodedLen(len(src)) bytes to dst.

​	Encode方法使用the encoding enc对`src`进行编码，将`EncodedLen(len(src))`个字节写入`dst`。

The encoding pads the output to a multiple of 4 bytes, so Encode is not appropriate for use on individual blocks of a large data stream. Use NewEncoder() instead.

​	该编码将输出填充为4字节的倍数，因此该Encode方法不适用于对大型数据流的单个块进行编码。请改用NewEncoder()函数。

##### Encode Example
``` go 
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("Hello, world!")
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	fmt.Println(string(dst))
}

Output:

SGVsbG8sIHdvcmxkIQ==
```

#### (*Encoding) EncodeToString 

``` go 
func (enc *Encoding) EncodeToString(src []byte) string
```

EncodeToString returns the base64 encoding of src.

​	EncodeToString方法返回`src`的base64编码字符串。

##### EncodeToString Example
``` go 
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("any + old & data")
	str := base64.StdEncoding.EncodeToString(data)
	fmt.Println(str)
}

Output:

YW55ICsgb2xkICYgZGF0YQ==
```

#### (*Encoding) EncodedLen 

``` go 
func (enc *Encoding) EncodedLen(n int) int
```

EncodedLen returns the length in bytes of the base64 encoding of an input buffer of length n.

​	EncodedLen方法返回长度为`n`的输入缓冲区的base64编码的字节数。

#### (Encoding) Strict  <- go1.8

``` go 
func (enc Encoding) Strict() *Encoding
```

Strict creates a new encoding identical to enc except with strict decoding enabled. In this mode, the decoder requires that trailing padding bits are zero, as described in [RFC 4648 section 3.5](https://rfc-editor.org/rfc/rfc4648.html#section-3.5).

​	Strict方法创建一个与enc相同但启用严格解码的新编码。在此模式下，解码器要求尾部填充位为零，如[RFC 4648第3.5节](https://rfc-editor.org/rfc/rfc4648.html#section-3.5)中所述。

Note that the input is still malleable, as new line characters (CR and LF) are still ignored.

请注意，输入仍然是可塑的（malleable），因为换行字符（CR和LF）仍然被忽略。

#### (Encoding) WithPadding  <- go1.5

``` go 
func (enc Encoding) WithPadding(padding rune) *Encoding
```

WithPadding creates a new encoding identical to enc except with a specified padding character, or NoPadding to disable padding. The padding character must not be '\r' or '\n', must not be contained in the encoding's alphabet and must be a rune equal or below '\xff'.

​	WithPadding方法创建一个与enc相同但具有指定填充字符的新编码，或者使用NoPadding禁用填充。填充字符不能是'\r'或'\n'，不能包含在编码的字母表中，并且必须是小于或等于'`\xff`'的符文。