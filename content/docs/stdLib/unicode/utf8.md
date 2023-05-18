+++
title = "utf8"
date = 2023-05-17T09:59:21+08:00
weight = 20
description = ""
isCJKLanguage = true
draft = false
+++

# utf8

[https://pkg.go.dev/unicode/utf8@go1.20.1](https://pkg.go.dev/unicode/utf8@go1.20.1)

​	utf8包实现了支持使用 UTF-8 编码的文本的函数和常量。它包括了在符文和 UTF-8 字节序列之间进行转换的函数。参见 https://en.wikipedia.org/wiki/UTF-8。


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=15)

``` go 
const (
	RuneError = '\uFFFD'     //  "错误" 符文或 "Unicode 替换字符"
	RuneSelf  = 0x80         // 低于 RuneSelf 的字符可以用单个字节表示。
	MaxRune   = '\U0010FFFF' // 最大有效 Unicode 代码点。
	UTFMax    = 4            // 一个 UTF-8 编码的 Unicode 字符的最大字节数。
)
```

这些数字是编码中的基本要素。

## 变量

This section is empty.

## 函数

#### func [AppendRune](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=375)  <- go1.18

``` go 
func AppendRune(p []byte, r rune) []byte
```

​	AppendRune函数将 r 的 UTF-8 编码附加到 p 的结尾并返回扩展后的缓冲区。如果符文超出范围，则附加 RuneError 的编码。

##### AppendRune Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf1 := utf8.AppendRune(nil, 0x10000)
	buf2 := utf8.AppendRune([]byte("init"), 0x10000)
	fmt.Println(string(buf1))
	fmt.Println(string(buf2))
}
Output:

𐀀
init𐀀
```

#### func [DecodeLastRune](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=247) 

``` go 
func DecodeLastRune(p []byte) (r rune, size int)
```

​	DecodeLastRune函数解码p中的最后一个UTF-8编码，并返回该符文及其占用的字节数。如果p为空，则返回(RuneError，0)。否则，如果编码无效，则返回(RuneError，1)。对于正确的非空UTF-8，这两种情况都是不可能的。

​	如果编码不正确，编码超出范围或不是该值的最短可能UTF-8编码，则编码无效。不执行其他验证。

##### DecodeLastRune Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("Hello, 世界")

	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		fmt.Printf("%c %v\n", r, size)

		b = b[:len(b)-size]
	}
}
Output:

界 3
世 3
  1
, 1
o 1
l 1
l 1
e 1
H 1
```

#### func [DecodeLastRuneInString](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=287) 

``` go 
func DecodeLastRuneInString(s string) (r rune, size int)
```

​	DecodeLastRuneInString函数类似于DecodeLastRune，但其输入为字符串。如果s为空，则返回(RuneError，0)。否则，如果编码无效，则返回(RuneError，1)。对于正确的非空UTF-8，这两种情况都是不可能的。

​	如果编码不正确，编码超出范围或不是该值的最短可能UTF-8编码，则编码无效。不执行其他验证。

##### DecodeLastRuneInString Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"

	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[:len(str)-size]
	}
}
Output:

界 3
世 3
  1
, 1
o 1
l 1
l 1
e 1
H 1
```

#### func [DecodeRune](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=151) 

``` go 
func DecodeRune(p []byte) (r rune, size int)
```

​	DecodeRune函数解码p中的第一个UTF-8编码，并返回该符文及其占用的字节数。如果p为空，则返回(RuneError，0)。否则，如果编码无效，则返回(RuneError，1)。对于正确的非空UTF-8，这两种情况都是不可能的。

​	如果编码不正确，编码超出范围或不是该值的最短可能UTF-8编码，则编码无效。不执行其他验证。

##### DecodeRune Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("Hello, 世界")

	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c %v\n", r, size)

		b = b[size:]
	}
}
Output:

H 1
e 1
l 1
l 1
o 1
, 1
  1
世 3
界 3
```

#### func [DecodeRuneInString](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=199) 

``` go 
func DecodeRuneInString(s string) (r rune, size int)
```

​	DecodeRuneInString函数类似于DecodeRune函数，但其输入为字符串。如果s为空，则返回(RuneError，0)。否则，如果编码无效，则返回(RuneError，1)。对于正确的非空UTF-8，这两种情况都是不可能的。

​	如果编码不正确，编码超出范围或不是该值的最短可能UTF-8编码，则编码无效。不执行其他验证。

##### DecodeRuneInString Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c %v\n", r, size)

		str = str[size:]
	}
}
Output:

H 1
e 1
l 1
l 1
o 1
, 1
  1
世 3
界 3
```

#### func [EncodeRune](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=342) 

``` go 
func EncodeRune(p []byte, r rune) int
```

EncodeRune writes into p (which must be large enough) the UTF-8 encoding of the rune. If the rune is out of range, it writes the encoding of RuneError. It returns the number of bytes written.

EncodeRune函数将符文的UTF-8编码写入p(p必须足够大)。如果符文超出范围，则写入RuneError的编码。返回写入的字节数。

##### EncodeRune Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	r := '世'
	buf := make([]byte, 3)

	n := utf8.EncodeRune(buf, r)

	fmt.Println(buf)
	fmt.Println(n)
}
Output:

[228 184 150]
3
```

##### EncodeRune Example(OutOfRange)
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	runes := []rune{
		// Less than 0, out of range.
		-1,
		// Greater than 0x10FFFF, out of range.
		0x110000,
		// The Unicode replacement character.
		utf8.RuneError,
	}
	for i, c := range runes {
		buf := make([]byte, 3)
		size := utf8.EncodeRune(buf, c)
		fmt.Printf("%d: %d %[2]s %d\n", i, buf, size)
	}
}
Output:

0: [239 191 189] � 3
1: [239 191 189] � 3
2: [239 191 189] � 3
```

#### func [FullRune](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=104) 

``` go 
func FullRune(p []byte) bool
```

​	FullRune函数报告p中的字节是否以完整的UTF-8符文编码开头。无效的编码被认为是完整的符文，因为它们将转换为宽度为1的错误符文。 

##### FullRune Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf := []byte{228, 184, 150} // 世
	fmt.Println(utf8.FullRune(buf))
	fmt.Println(utf8.FullRune(buf[:2]))
}
Output:

true
false
```

#### func [FullRuneInString](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=124) 

``` go 
func FullRuneInString(s string) bool
```

​	FullRuneInString函数类似于FullRune，但其输入是字符串。

##### FullRuneInString Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "世"
	fmt.Println(utf8.FullRuneInString(str))
	fmt.Println(utf8.FullRuneInString(str[:2]))
}
Output:

true
false
```

#### func [RuneCount](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=400) 

``` go 
func RuneCount(p []byte) int
```

​	RuneCount函数返回p中符文的数量。错误和短编码被视为宽度为1个字节的单个符文。

##### RuneCount Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf := []byte("Hello, 世界")
	fmt.Println("bytes =", len(buf))
	fmt.Println("runes =", utf8.RuneCount(buf))
}
Output:

bytes = 13
runes = 9
```

#### func [RuneCountInString](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=437) 

``` go 
func RuneCountInString(s string) (n int)
```

​	RuneCountInString函数类似于RuneCount，但其输入是字符串。

##### RuneCountInString Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"
	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", utf8.RuneCountInString(str))
}
Output:

bytes = 13
runes = 9
```

#### func [RuneLen](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=321) 

``` go 
func RuneLen(r rune) int
```

​	RuneLen函数返回编码符文所需的字节数。如果符文不是UTF-8的有效值，则返回-1。

##### RuneLen Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('界'))
}
Output:

1
3
```

#### func [RuneStart](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=474) 

``` go 
func RuneStart(b byte) bool
```

​	RuneStart函数报告字节是否可以是编码的第一个字节，可能无效。第二个及后续字节的前两位始终设置为10。

##### RuneStart Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf := []byte("a界")
	fmt.Println(utf8.RuneStart(buf[0]))
	fmt.Println(utf8.RuneStart(buf[1]))
	fmt.Println(utf8.RuneStart(buf[2]))
}
Output:

true
true
false
```

#### func [Valid](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=477) 

``` go 
func Valid(p []byte) bool
```

​	Valid函数报告p是否完全由有效的UTF-8编码符文组成。

##### Valid Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}

	fmt.Println(utf8.Valid(valid))
	fmt.Println(utf8.Valid(invalid))
}
Output:

true
false
```

#### func [ValidRune](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=575)  <- go1.1

``` go 
func ValidRune(r rune) bool
```

​	ValidRune函数报告r是否可以合法地编码为UTF-8。超出范围或替代字符的一半的代码点是非法的。

##### ValidRune Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := 'a'
	invalid := rune(0xfffffff)

	fmt.Println(utf8.ValidRune(valid))
	fmt.Println(utf8.ValidRune(invalid))
}
Output:

true
false
```

#### func [ValidString](https://cs.opensource.google/go/go/+/go1.20.1:src/unicode/utf8/utf8.go;l=528) 

``` go 
func ValidString(s string) bool
```

​	ValidString函数报告s是否完全由有效的UTF-8编码符文组成。

##### ValidString Example
``` go 
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	valid := "Hello, 世界"
	invalid := string([]byte{0xff, 0xfe, 0xfd})

	fmt.Println(utf8.ValidString(valid))
	fmt.Println(utf8.ValidString(invalid))
}
Output:

true
false
```

## 类型

This section is empty.