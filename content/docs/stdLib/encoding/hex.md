+++
title = "hex"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# hex

https://pkg.go.dev/encoding/hex@go1.20.1



Package hex implements hexadecimal encoding and decoding.

包hex实现了十六进制的编码和解码。












## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=57)

``` go linenums="1"
var ErrLength = errors.New("encoding/hex: odd length hex string")
```

ErrLength reports an attempt to decode an odd-length input using Decode or DecodeString. The stream-based Decoder returns io.ErrUnexpectedEOF instead of ErrLength.

ErrLength报告使用Decode或DecodeString解码一个奇长的输入的尝试。基于流的解码器返回io.ErrUnexpectedEOF而不是ErrLength。

## 函数

#### func [Decode](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=77) 

``` go linenums="1"
func Decode(dst, src []byte) (int, error)
```

Decode decodes src into DecodedLen(len(src)) bytes, returning the actual number of bytes written to dst.

Decode将src解码为DecodedLen(len(src))字节，返回写给dst的实际字节数。

Decode expects that src contains only hexadecimal characters and that src has even length. If the input is malformed, Decode returns the number of bytes decoded before the error.

解码期望src只包含十六进制的字符，并且src的长度是偶数。如果输入是畸形的，Decode会返回错误发生前的解码字节数。

##### Example
``` go linenums="1"
```

#### func [DecodeString](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=118) 

``` go linenums="1"
func DecodeString(s string) ([]byte, error)
```

DecodeString returns the bytes represented by the hexadecimal string s.

DecodeString返回十六进制字符串s所代表的字节数。

DecodeString expects that src contains only hexadecimal characters and that src has even length. If the input is malformed, DecodeString returns the bytes decoded before the error.

DecodeString期望src只包含十六进制的字符，并且src具有偶数长度。如果输入是畸形的，DecodeString将返回错误之前的解码字节。

##### Example
``` go linenums="1"
```

#### func [DecodedLen](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=68) 

``` go linenums="1"
func DecodedLen(x int) int
```

DecodedLen returns the length of a decoding of x source bytes. Specifically, it returns x / 2.

DecodedLen返回x个源字节的解码长度。具体来说，它返回x/2。

#### func [Dump](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=128) 

``` go linenums="1"
func Dump(data []byte) string
```

Dump returns a string that contains a hex dump of the given data. The format of the hex dump matches the output of `hexdump -C` on the command line.

Dump返回一个包含给定数据的十六进制转储的字符串。十六进制转储的格式与命令行中`hexdump -C`的输出相匹配。

##### Example
``` go linenums="1"
```

#### func [Dumper](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=224) 

``` go linenums="1"
func Dumper(w io.Writer) io.WriteCloser
```

Dumper returns a WriteCloser that writes a hex dump of all written data to w. The format of the dump matches the output of `hexdump -C` on the command line.

Dumper返回一个WriteCloser，将所有写入的数据的十六进制转储到w。

##### Example
``` go linenums="1"
```

#### func [Encode](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=44) 

``` go linenums="1"
func Encode(dst, src []byte) int
```

Encode encodes src into EncodedLen(len(src)) bytes of dst. As a convenience, it returns the number of bytes written to dst, but this value is always EncodedLen(len(src)). Encode implements hexadecimal encoding.

Encode将src编码为dst的EncodedLen(len(src))字节。为了方便起见，它返回写入dst的字节数，但这个值总是EncodedLen(len(src))。Encode实现了十六进制的编码。

##### Example
``` go linenums="1"
```

#### func [EncodeToString](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=106) 

``` go linenums="1"
func EncodeToString(src []byte) string
```

EncodeToString returns the hexadecimal encoding of src.

EncodeToString返回src的十六进制编码。

##### Example
``` go linenums="1"
```

#### func [EncodedLen](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=38) 

``` go linenums="1"
func EncodedLen(n int) int
```

EncodedLen returns the length of an encoding of n source bytes. Specifically, it returns n * 2.

EncodedLen返回n个源字节的编码的长度。具体来说，它返回n * 2。

#### func [NewDecoder](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=184)  <- go1.10

``` go linenums="1"
func NewDecoder(r io.Reader) io.Reader
```

NewDecoder returns an io.Reader that decodes hexadecimal characters from r. NewDecoder expects that r contain only an even number of hexadecimal characters.

NewDecoder返回一个io.Reader，对r中的十六进制字符进行解码。NewDecoder希望r中只包含偶数的十六进制字符。

#### func [NewEncoder](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=155)  <- go1.10

``` go linenums="1"
func NewEncoder(w io.Writer) io.Writer
```

NewEncoder returns an io.Writer that writes lowercase hexadecimal characters to w.

NewEncoder返回一个io.Writer，将小写的十六进制字符写入w中。

## 类型

### type [InvalidByteError](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=60) 

``` go linenums="1"
type InvalidByteError byte
```

InvalidByteError values describe errors resulting from an invalid byte in a hex string.

InvalidByteError值描述由十六进制字符串中的无效字节导致的错误。

#### (InvalidByteError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=62) 

``` go linenums="1"
func (e InvalidByteError) Error() string
```