+++
title = "hex"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/encoding/hex@go1.21.3](https://pkg.go.dev/encoding/hex@go1.21.3)

Package hex implements hexadecimal encoding and decoding.

​	`hex`包实现了十六进制的编码和解码。


## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/hex/hex.go;l=57)

``` go 
var ErrLength = errors.New("encoding/hex: odd length hex string")
```

ErrLength reports an attempt to decode an odd-length input using Decode or DecodeString. The stream-based Decoder returns io.ErrUnexpectedEOF instead of ErrLength.

​	ErrLength报告使用Decode或DecodeString解码一个奇长的输入的尝试。基于流的解码器返回io.ErrUnexpectedEOF而不是ErrLength。

## 函数

### func AppendDecode <- go1.22.0

```
func AppendDecode(dst, src []byte) ([]byte, error)
```

AppendDecode appends the hexadecimally decoded src to dst and returns the extended buffer. If the input is malformed, it returns the partially decoded src and an error.

​	AppendDecode 将十六进制解码后的 src 追加到 dst，并返回扩展后的缓冲区。如果输入格式不正确，它将返回部分解码的 src 和一个错误。

### func AppendEncode <- go1.22.0

```
func AppendEncode(dst, src []byte) []byte
```

AppendEncode appends the hexadecimally encoded src to dst and returns the extended buffer.

​	AppendEncode 将十六进制编码的 src 追加到 dst，并返回扩展后的缓冲区。

### func Decode 

``` go 
func Decode(dst, src []byte) (int, error)
```

Decode decodes src into DecodedLen(len(src)) bytes, returning the actual number of bytes written to dst.

​	Decode将src解码为DecodedLen(len(src))字节，返回写给dst的实际字节数。

Decode expects that src contains only hexadecimal characters and that src has even length. If the input is malformed, Decode returns the number of bytes decoded before the error.

​	解码期望src只包含十六进制的字符，并且src的长度是偶数。如果输入是畸形的，Decode会返回错误发生前的解码字节数。

#### Decode Example
``` go 
package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	src := []byte("48656c6c6f20476f7068657221")

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])

}

Output:

Hello Gopher!
```

### func DecodeString 

``` go 
func DecodeString(s string) ([]byte, error)
```

DecodeString returns the bytes represented by the hexadecimal string s.

​	DecodeString返回十六进制字符串s所代表的字节数。

DecodeString expects that src contains only hexadecimal characters and that src has even length. If the input is malformed, DecodeString returns the bytes decoded before the error.

​	DecodeString期望src只包含十六进制的字符，并且src具有偶数长度。如果输入是畸形的，DecodeString将返回错误之前的解码字节。

#### DecodeString Example
``` go 
package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	const s = "48656c6c6f20476f7068657221"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)

}

Output:

Hello Gopher!
```

### func DecodedLen 

``` go 
func DecodedLen(x int) int
```

DecodedLen returns the length of a decoding of x source bytes. Specifically, it returns x / 2.

​	DecodedLen返回x个源字节的解码长度。具体来说，它返回x/2。

### func Dump 

``` go 
func Dump(data []byte) string
```

Dump returns a string that contains a hex dump of the given data. The format of the hex dump matches the output of `hexdump -C` on the command line.

​	Dump返回一个包含给定数据的十六进制转储的字符串。十六进制转储的格式与命令行中`hexdump -C`的输出相匹配。

#### Dump Example
``` go 
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	content := []byte("Go is an open source programming language.")

	fmt.Printf("%s", hex.Dump(content))

}

Output:

00000000  47 6f 20 69 73 20 61 6e  20 6f 70 65 6e 20 73 6f  |Go is an open so|
00000010  75 72 63 65 20 70 72 6f  67 72 61 6d 6d 69 6e 67  |urce programming|
00000020  20 6c 61 6e 67 75 61 67  65 2e                    | language.|
```

### func Dumper 

``` go 
func Dumper(w io.Writer) io.WriteCloser
```

Dumper returns a WriteCloser that writes a hex dump of all written data to w. The format of the dump matches the output of `hexdump -C` on the command line.

​	Dumper返回一个WriteCloser，将所有写入的数据的十六进制转储到w。

#### Dumper Example
``` go 
package main

import (
	"encoding/hex"
	"os"
)

func main() {
	lines := []string{
		"Go is an open source programming language.",
		"\n",
		"We encourage all Go users to subscribe to golang-announce.",
	}

	stdoutDumper := hex.Dumper(os.Stdout)

	defer stdoutDumper.Close()

	for _, line := range lines {
		stdoutDumper.Write([]byte(line))
	}

}

Output:

00000000  47 6f 20 69 73 20 61 6e  20 6f 70 65 6e 20 73 6f  |Go is an open so|
00000010  75 72 63 65 20 70 72 6f  67 72 61 6d 6d 69 6e 67  |urce programming|
00000020  20 6c 61 6e 67 75 61 67  65 2e 0a 57 65 20 65 6e  | language..We en|
00000030  63 6f 75 72 61 67 65 20  61 6c 6c 20 47 6f 20 75  |courage all Go u|
00000040  73 65 72 73 20 74 6f 20  73 75 62 73 63 72 69 62  |sers to subscrib|
00000050  65 20 74 6f 20 67 6f 6c  61 6e 67 2d 61 6e 6e 6f  |e to golang-anno|
00000060  75 6e 63 65 2e                                    |unce.|
```

### func Encode 

``` go 
func Encode(dst, src []byte) int
```

Encode encodes src into EncodedLen(len(src)) bytes of dst. As a convenience, it returns the number of bytes written to dst, but this value is always EncodedLen(len(src)). Encode implements hexadecimal encoding.

​	Encode将src编码为dst的EncodedLen(len(src))字节。为了方便起见，它返回写入dst的字节数，但这个值总是EncodedLen(len(src))。Encode实现了十六进制的编码。

#### Encode Example
``` go 
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("Hello Gopher!")

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	fmt.Printf("%s\n", dst)

}

Output:

48656c6c6f20476f7068657221
```

### func EncodeToString 

``` go 
func EncodeToString(src []byte) string
```

EncodeToString returns the hexadecimal encoding of src.

​	EncodeToString返回src的十六进制编码。

#### EncodeToString Example
``` go 
package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)

}

Output:

48656c6c6f
```

### func EncodedLen 

``` go 
func EncodedLen(n int) int
```

EncodedLen returns the length of an encoding of n source bytes. Specifically, it returns n * 2.

​	EncodedLen返回n个源字节的编码的长度。具体来说，它返回n * 2。

### func NewDecoder  <- go1.10

``` go 
func NewDecoder(r io.Reader) io.Reader
```

NewDecoder returns an io.Reader that decodes hexadecimal characters from r. NewDecoder expects that r contain only an even number of hexadecimal characters.

​	NewDecoder返回一个io.Reader，对r中的十六进制字符进行解码。NewDecoder希望r中只包含偶数的十六进制字符。

### func NewEncoder  <- go1.10

``` go 
func NewEncoder(w io.Writer) io.Writer
```

NewEncoder returns an io.Writer that writes lowercase hexadecimal characters to w.

​	NewEncoder返回一个io.Writer，将小写的十六进制字符写入w中。

## 类型

### type InvalidByteError 

``` go 
type InvalidByteError byte
```

InvalidByteError values describe errors resulting from an invalid byte in a hex string.

​	InvalidByteError值描述由十六进制字符串中的无效字节导致的错误。

#### (InvalidByteError) Error 

``` go 
func (e InvalidByteError) Error() string
```