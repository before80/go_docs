+++
title = "base32"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# base32

https://pkg.go.dev/encoding/base32@go1.20.1



Package base32 implements base32 encoding as specified by [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html).

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

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/base32/base32.go;l=72)

``` go 
var StdEncoding = NewEncoding(encodeStd)
```

StdEncoding is the standard base32 encoding, as defined in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html).

## 函数

#### func NewDecoder 

``` go 
func NewDecoder(enc *Encoding, r io.Reader) io.Reader
```

NewDecoder constructs a new base32 stream decoder.

#### func NewEncoder 

``` go 
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser
```

NewEncoder returns a new base32 stream encoder. Data written to the returned writer will be encoded using enc and then written to w. Base32 encodings operate in 5-byte blocks; when finished writing, the caller must Close the returned encoder to flush any partially written blocks.

##### NewEncoder Example
``` go 
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

An Encoding is a radix 32 encoding/decoding scheme, defined by a 32-character alphabet. The most common is the "base32" encoding introduced for SASL GSSAPI and standardized in [RFC 4648](https://rfc-editor.org/rfc/rfc4648.html). The alternate "base32hex" encoding is used in DNSSEC.

#### func NewEncoding 

``` go 
func NewEncoding(encoder string) *Encoding
```

NewEncoding returns a new Encoding defined by the given alphabet, which must be a 32-byte string.

#### (*Encoding) Decode 

``` go 
func (enc *Encoding) Decode(dst, src []byte) (n int, err error)
```

Decode decodes src using the encoding enc. It writes at most DecodedLen(len(src)) bytes to dst and returns the number of bytes written. If src contains invalid base32 data, it will return the number of bytes successfully written and CorruptInputError. New line characters (\r and \n) are ignored.

##### Decode Example
``` go 
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

``` go 
func (enc *Encoding) DecodeString(s string) ([]byte, error)
```

DecodeString returns the bytes represented by the base32 string s.

##### DecodeString Example
``` go 
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

``` go 
func (enc *Encoding) DecodedLen(n int) int
```

DecodedLen returns the maximum length in bytes of the decoded data corresponding to n bytes of base32-encoded data.

#### (*Encoding) Encode 

``` go 
func (enc *Encoding) Encode(dst, src []byte)
```

Encode encodes src using the encoding enc, writing EncodedLen(len(src)) bytes to dst.

The encoding pads the output to a multiple of 8 bytes, so Encode is not appropriate for use on individual blocks of a large data stream. Use NewEncoder() instead.

##### Encode Example
``` go 
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

``` go 
func (enc *Encoding) EncodeToString(src []byte) string
```

EncodeToString returns the base32 encoding of src.

##### EncodeToString Example
``` go 
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

``` go 
func (enc *Encoding) EncodedLen(n int) int
```

EncodedLen returns the length in bytes of the base32 encoding of an input buffer of length n.

#### (Encoding) WithPadding  <- go1.9

``` go 
func (enc Encoding) WithPadding(padding rune) *Encoding
```

WithPadding creates a new encoding identical to enc except with a specified padding character, or NoPadding to disable padding. The padding character must not be '\r' or '\n', must not be contained in the encoding's alphabet and must be a rune equal or below '\xff'.