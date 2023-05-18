+++
title = "base64"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# base64

https://pkg.go.dev/encoding/base64@go1.20.1



Package base64 implements base64 encoding as specified by [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html).

包 base64 实现了 RFC 4648 所规定的 base64 编码。

##### Example
``` go 
```












## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=30)

``` go 
const (
	StdPadding rune = '=' // Standard padding character // 标准填充字符
	NoPadding  rune = -1  // No padding  // 无填充字符

)
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=123)

``` go 
var RawStdEncoding = StdEncoding.WithPadding(NoPadding)
```

RawStdEncoding is the standard raw, unpadded base64 encoding, as defined in [RFC 4648 section 3.2](https://rfc-editor.org/rfc/rfc4648.html#section-3.2). This is the same as StdEncoding but omits padding characters.

RawStdEncoding是标准的原始、无填充的base64编码，定义于RFC 4648第3.2节。这与StdEncoding相同，但省略了填充字符。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=128)

``` go 
var RawURLEncoding = URLEncoding.WithPadding(NoPadding)
```

RawURLEncoding is the unpadded alternate base64 encoding defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html). It is typically used in URLs and file names. This is the same as URLEncoding but omits padding characters.

RawURLEncoding是RFC 4648中定义的无填充的另一种base64编码。它通常在URL和文件名中使用。这与URLEncoding相同，但省略了填充字符。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=114)

``` go 
var StdEncoding = NewEncoding(encodeStd)
```

StdEncoding is the standard base64 encoding, as defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html).

StdEncoding是标准的base64编码，如RFC 4648所定义。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=118)

``` go 
var URLEncoding = NewEncoding(encodeURL)
```

URLEncoding is the alternate base64 encoding defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html). It is typically used in URLs and file names.

URLEncoding是RFC 4648中定义的备用base64编码。它通常在URL和文件名中使用。

## 函数

#### func [NewDecoder](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=614) 

``` go 
func NewDecoder(enc *Encoding, r io.Reader) io.Reader
```

NewDecoder constructs a new base64 stream decoder.

NewDecoder构建一个新的base64流解码器。

#### func [NewEncoder](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=270) 

``` go 
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
```

NewEncoder returns a new base64 stream encoder. Data written to the returned writer will be encoded using enc and then written to w. Base64 encodings operate in 4-byte blocks; when finished writing, the caller must Close the returned encoder to flush any partially written blocks.

NewEncoder返回一个新的base64流编码器。写入返回的写入器的数据将使用enc进行编码，然后写入w。Base64编码以4字节的块进行操作；当写完后，调用者必须关闭返回的编码器以冲刷任何部分写入的块。

##### Example
``` go 
```

## 类型

### type [CorruptInputError](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=287) 

``` go 
type CorruptInputError int64
```

#### (CorruptInputError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=289) 

``` go 
func (e CorruptInputError) Error() string
```

### type [Encoding](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=23) 

``` go 
type Encoding struct {
	// contains filtered or unexported fields
}
```

An Encoding is a radix 64 encoding/decoding scheme, defined by a 64-character alphabet. The most common encoding is the "base64" encoding defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html) and used in MIME ([RFC 2045](https://rfc-editor.org/rfc/rfc2045.html)) and PEM ([RFC 1421](https://rfc-editor.org/rfc/rfc1421.html)). [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html) also defines an alternate encoding, which is the standard encoding with - and _ substituted for + and /.

编码是一个弧度为64的编码/解码方案，由一个64字符的字母表定义。最常见的编码是RFC 4648中定义的 "base64 "编码，在MIME(RFC 2045)和PEM(RFC 1421)中使用。RFC 4648还定义了一个备用编码，即用-和_代替+和/的标准编码。

#### func [NewEncoding](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=60) 

``` go 
func NewEncoding(encoder string) *Encoding
```

NewEncoding returns a new padded Encoding defined by the given alphabet, which must be a 64-byte string that does not contain the padding character or CR / LF ('\r', '\n'). The resulting Encoding uses the default padding character ('='), which may be changed or disabled via WithPadding.

NewEncoding返回一个由给定字母定义的新的填充编码，它必须是一个64字节的字符串，不包含填充字符或CR / LF('\r', '\n')。产生的Encoding使用默认的padding字符('=')，可以通过WithPadding改变或禁用。

#### (*Encoding) [Decode](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=484) 

``` go 
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)
```

Decode decodes src using the encoding enc. It writes at most DecodedLen(len(src)) bytes to dst and returns the number of bytes written. If src contains invalid base64 data, it will return the number of bytes successfully written and CorruptInputError. New line characters (\r and \n) are ignored.

Decode使用enc编码对src进行解码。它最多向dst写入DecodedLen(len(src))字节，并返回写入的字节数。如果src包含无效的base64数据，它将返回成功写入的字节数和CorruptInputError。新行字符(\r 和 \n)被忽略。

##### Example
``` go 
```

#### (*Encoding) [DecodeString](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=396) 

``` go 
func (enc *Encoding) DecodeString(s string) ([]byte, error)
```

DecodeString returns the bytes represented by the base64 string s.

DecodeString返回base64字符串s所代表的字节。

##### Example
``` go 
```

#### (*Encoding) [DecodedLen](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=620) 

``` go 
func (enc *Encoding) DecodedLen(n int) int
```

DecodedLen returns the maximum length in bytes of the decoded data corresponding to n bytes of base64-encoded data.

DecodedLen返回对应于base64编码的n个字节的解码数据的最大长度(字节)。

#### (*Encoding) [Encode](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=140) 

``` go 
func (enc *Encoding) Encode(dst, src []byte)
```

Encode encodes src using the encoding enc, writing EncodedLen(len(src)) bytes to dst.

Encode使用enc编码对src进行编码，向dst写入EncodedLen(len(src))字节。

The encoding pads the output to a multiple of 4 bytes, so Encode is not appropriate for use on individual blocks of a large data stream. Use NewEncoder() instead.

编码将输出填充为4字节的倍数，所以Encode不适合用于大数据流的单个块。请使用NewEncoder()代替。

##### Example
``` go 
```

#### (*Encoding) [EncodeToString](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=192) 

``` go 
func (enc *Encoding) EncodeToString(src []byte) string
```

EncodeToString returns the base64 encoding of src.

EncodeToString返回src的base64编码。

##### Example
``` go 
```

#### (*Encoding) [EncodedLen](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=276) 

``` go 
func (enc *Encoding) EncodedLen(n int) int
```

EncodedLen returns the length in bytes of the base64 encoding of an input buffer of length n.

EncodedLen返回一个长度为n的输入缓冲区的base64编码的字节长度。

#### (Encoding) [Strict](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=107)  <- go1.8

``` go 
func (enc Encoding) Strict() *Encoding
```

Strict creates a new encoding identical to enc except with strict decoding enabled. In this mode, the decoder requires that trailing padding bits are zero, as described in [RFC 4648 section 3.5](https://rfc-editor.org/rfc/rfc4648.html#section-3.5).

严格创建一个与enc相同的新编码，只是启用了严格解码。在这种模式下，解码器要求尾部填充位为零，如RFC 4648第3.5节所述。

Note that the input is still malleable, as new line characters (CR and LF) are still ignored.

注意，输入仍然是可塑的，因为新的行字符(CR和LF)仍然被忽略。

#### (Encoding) [WithPadding](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base64/base64.go;l=86)  <- go1.5

``` go 
func (enc Encoding) WithPadding(padding rune) *Encoding
```

WithPadding creates a new encoding identical to enc except with a specified padding character, or NoPadding to disable padding. The padding character must not be '\r' or '\n', must not be contained in the encoding's alphabet and must be a rune equal or below '\xff'.

WithPadding创建一个与enc相同的新编码，除了指定的padding字符，或者用NoPadding来禁用padding。填充字符不能是'\r'或'\n'，不能包含在编码的字母表中，必须是等于或低于'\xff'的符文。