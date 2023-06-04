+++
title = "strconv"
linkTitle = "strconv"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# strconv

https://pkg.go.dev/strconv@go1.20.1

​	strconv包实现了基本数据类型的字符串表示与其相互转换的功能。

## 数字转换

​	最常见的数字转换是Atoi(字符串到整数)和Itoa(整数到字符串)。

> ​	函数名`Atoi`是将字符串转换为整数的缩写，其中`A`代表`ASCII`字符集，即该函数只能处理使用`ASCII`字符集表示的数字字符，而不能处理使用其他字符集表示的数字字符。`toi`则是`to int`的缩写，表示将字符串转换为整数。因此，`Atoi`函数名的含义为将使用`ASCII`字符集表示的字符串转换为整数。

```
i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)
```

​	这些假设十进制和Go int类型。

​	ParseBool函数、ParseFloat函数、ParseInt函数和ParseUint函数将字符串转换为相应的数值：

```
b, err := strconv.ParseBool("true")
f, err := strconv.ParseFloat("3.1415", 64)
i, err := strconv.ParseInt("-42", 10, 64)
u, err := strconv.ParseUint("42", 10, 64)
```

​	解析函数返回最宽的类型(float64、int64和uint64)，但如果size参数指定了更窄的宽度，则结果可以转换为该更窄的类型而不会丢失数据：

```
s := "2147483647" // biggest int32 // 最大的int32
i64, err := strconv.ParseInt(s, 10, 32)
...
i := int32(i64)
```

​	FormatBool函数、FormatFloat函数、FormatInt函数和FormatUint函数将值转换为字符串：

```
s := strconv.FormatBool(true)
s := strconv.FormatFloat(3.1415, 'E', -1, 64)
s := strconv.FormatInt(-42, 16)
s := strconv.FormatUint(42, 16)
```

​	AppendBool函数、AppendFloat函数、AppendInt函数和AppendUint函数类似，但将格式化的值附加到目标切片。

## 字符串转换

​	Quote函数和QuoteToASCII函数将字符串转换为带引号的Go字符串文字。后者通过使用`\u`转义任何非ASCII Unicode来保证结果是ASCII字符串：

```
q := strconv.Quote("Hello, 世界")
q := strconv.QuoteToASCII("Hello, 世界")
```

​	QuoteRune函数和QuoteRuneToASCII函数类似，但接受符文并返回带引号的Go符文文字。

​	Unquote函数和UnquoteChar函数取消引用Go字符串和符文文字。


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/strconv/atoi.go;l=71)

``` go 
const IntSize = intSize
```

​	IntSize 是 int 或 uint 值的位数大小。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/strconv/atoi.go;l=18)

``` go 
var ErrRange = errors.New("value out of range")
```

​	ErrRange 表示该值超出目标类型的范围。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/strconv/atoi.go;l=21)

``` go 
var ErrSyntax = errors.New("invalid syntax")
```

​	ErrSyntax 表示该值对于目标类型来说语法不正确。

## 函数

#### func AppendBool 

``` go 
func AppendBool(dst []byte, b bool) []byte
```

​	AppendBool函数根据 b 的值将 "true" 或 "false" 追加到 dst 中并返回扩展后的缓冲区。

##### AppendBool Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))

}
Output:

bool:true
```

#### func AppendFloat 

``` go 
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
```

​	AppendFloat函数将浮点数 f 的字符串形式(由 FormatFloat 生成)追加到 dst 中并返回扩展后的缓冲区。

##### AppendFloat Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(b32))

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64))

}
Output:

float32:3.1415927E+00
float64:3.1415926535E+00
```

#### func AppendInt 

``` go 
func AppendInt(dst []byte, i int64, base int) []byte
```

​	AppendInt函数将整数 i 的字符串形式(由 FormatInt函数 生成)追加到 dst 中并返回扩展后的缓冲区。

##### AppendInt Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))

}
Output:

int (base 10):-42
int (base 16):-2a
```

#### func AppendQuote 

``` go 
func AppendQuote(dst []byte, s string) []byte
```

​	AppendQuote函数将表示 s 的双引号 Go 字符串文字(由 Quote函数 生成)追加到 dst 中并返回扩展后的缓冲区。

##### AppendQuote Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("quote:")
	b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

}
Output:

quote:"\"Fran & Freddie's Diner\""
```

#### func AppendQuoteRune 

``` go 
func AppendQuoteRune(dst []byte, r rune) []byte
```

​	AppendQuoteRune函数将表示符文的单引号 Go 字符文字(由 QuoteRune函数 生成)追加到 dst 中并返回扩展后的缓冲区。

##### AppendQuoteRune Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("rune:")
	b = strconv.AppendQuoteRune(b, '☺')
	fmt.Println(string(b))

}
Output:

rune:'☺'
```

#### func AppendQuoteRuneToASCII 

``` go 
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
```

​	AppendQuoteRuneToASCII函数将表示rune的单引号Go字符文本(由QuoteRuneToASCII生成)附加到dst并返回扩展的缓冲区。

##### AppendQuoteRuneToASCII Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("rune (ascii):")
	b = strconv.AppendQuoteRuneToASCII(b, '☺')
	fmt.Println(string(b))

}
Output:

rune (ascii):'\u263a'
```

#### func AppendQuoteRuneToGraphic  <- go1.6

``` go 
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte
```

​	AppendQuoteRuneToGraphic函数将表示rune的单引号Go字符文本(由QuoteRuneToGraphic函数生成)附加到dst并返回扩展的缓冲区。

#### func AppendQuoteToASCII 

``` go 
func AppendQuoteToASCII(dst []byte, s string) []byte
```

​	AppendQuoteToASCII函数将表示s的双引号Go字符串文本(由QuoteToASCII函数生成)附加到dst并返回扩展的缓冲区。

##### AppendQuoteToASCII Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("quote (ascii):")
	b = strconv.AppendQuoteToASCII(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

}
Output:

quote (ascii):"\"Fran & Freddie's Diner\""
```

#### func AppendQuoteToGraphic  <- go1.6

``` go 
func AppendQuoteToGraphic(dst []byte, s string) []byte
```

​	AppendQuoteToGraphic函数将表示s的双引号Go字符串文本(由QuoteToGraphic函数生成)附加到dst并返回扩展的缓冲区。

#### func AppendUint 

``` go 
func AppendUint(dst []byte, i uint64, base int) []byte
```

​	AppendUint函数将生成的无符号整数i的字符串形式(由FormatUint函数生成)附加到dst并返回扩展的缓冲区。

##### AppendUint Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b10 := []byte("uint (base 10):")
	b10 = strconv.AppendUint(b10, 42, 10)
	fmt.Println(string(b10))

	b16 := []byte("uint (base 16):")
	b16 = strconv.AppendUint(b16, 42, 16)
	fmt.Println(string(b16))

}
Output:

uint (base 10):42
uint (base 16):2a
```

#### func Atoi 

``` go 
func Atoi(s string) (int, error)
```

​	Atoi函数等价于ParseInt(s，10，0)，转换为int类型。

##### Atoi Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "10"
	if s, err := strconv.Atoi(v); err == nil {
		fmt.Printf("%T, %v", s, s)
	}

}
Output:

int, 10
```

#### func CanBackquote 

``` go 
func CanBackquote(s string) bool
```

​	CanBackquote函数报告字符串s是否可以表示为单行反引号字符串而不带控制字符(除制表符之外)。

> `CanBackquote`函数会返回一个布尔值，指示是否可以使用Go语言中的反引号来包裹给定的字符串。如果可以使用反引号，则返回`true`，否则返回`false`。

##### CanBackquote Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.CanBackquote("Fran & Freddie's Diner ☺"))
	fmt.Println(strconv.CanBackquote("`can't backquote this`"))

}
Output:

true
false
```

#### func FormatBool 

``` go 
func FormatBool(b bool) string
```

​	FormatBool函数根据 b 的值返回 "true" 或 "false"。

##### FormatBool Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := true
	s := strconv.FormatBool(v)
	fmt.Printf("%T, %v\n", s, s)

}
Output:

string, true
```

#### func FormatComplex  <- go1.15

``` go 
func FormatComplex(c complex128, fmt byte, prec, bitSize int) string
```

​	FormatComplex函数将复数 c 格式化为 (a+bi) 形式的字符串，其中 a 和 b 是实部和虚部，根据格式fmt和精度prec进行格式化。

​	格式 fmt 和精度 prec 的含义与 FormatFloat函数 相同。它假定原始值是从 bitSize 位的复数值(complex64 的 bitSize 必须是 64，而 complex128 的 bitSize 必须是 128)获得的，四舍五入结果。

#### func FormatFloat 

``` go 
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
```

​	FormatFloat函数将浮点数 f 格式化为字符串，格式由 fmt 和精度 prec 指定。它假定原始值是从 bitSize 位(float32 为 32，float64 为 64)的浮点数值获得的。

​	格式 fmt 是 'b'(-ddddp±ddd，二进制指数)、'e'(-d.dddde±dd，十进制指数)、'E'(-d.ddddE±dd，十进制指数)、'f'(-ddd.dddd，无指数)、'g'(大指数时为 'e'，否则为 'f')、'G'(大指数时为 'E'，否则为 'f')、'x'(-0xd.ddddp±ddd，十六进制小数和二进制指数)或 'X'(-0Xd.ddddP±ddd，十六进制小数和二进制指数)之一。

​	精度 prec 控制由 'e'、'E'、'f'、'g'、'G'、'x' 和 'X' 格式打印的数字的位数(不包括指数)。对于 'e'、'E'、'f'、'x' 和 'X'，它是小数点后的数字位数。对于 'g' 和 'G'，它是最大的有效数字位数(尾随零被删除)。特殊精度 -1 使用最少数量的位数，使 ParseFloat 精确返回 f。

##### FormatFloat Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := 3.1415926535

	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)

}
Output:

string, 3.1415927E+00
string, 3.1415926535E+00
```

#### func FormatInt 

``` go 
func FormatInt(i int64, base int) string
```

​	FormatInt函数返回基于给定进制 base 中 i 的字符串表示形式，2 <= base <= 36。结果对于值 >= 10 的数字使用小写字母 'a' 到 'z'。

##### FormatInt Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := int64(-42)

	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatInt(v, 16)
	fmt.Printf("%T, %v\n", s16, s16)

}
Output:

string, -42
string, -2a
```

#### func FormatUint 

``` go 
func FormatUint(i uint64, base int) string
```

​	FormatUint函数返回基于给定进制 base 中 i 的字符串表示形式，2 <= base <= 36。结果对于值 >= 10 的数字使用小写字母 'a' 到 'z'。

##### FormatUint Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := uint64(42)

	s10 := strconv.FormatUint(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatUint(v, 16)
	fmt.Printf("%T, %v\n", s16, s16)

}
Output:

string, 42
string, 2a
```

#### func IsGraphic  <- go1.6

``` go 
func IsGraphic(r rune) bool
```

​	IsGraphic函数报告 r 是否被 Unicode 定义为图形字符。这些字符包括类别 L、M、N、P、S和Z。

##### IsGraphic Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	shamrock := strconv.IsGraphic('☘')
	fmt.Println(shamrock)

	a := strconv.IsGraphic('a')
	fmt.Println(a)

	bel := strconv.IsGraphic('\007')
	fmt.Println(bel)

}
Output:

true
true
false
```

#### func IsPrint 

``` go 
func IsPrint(r rune) bool
```

​	IsPrint 函数判断 rune 是否可打印，其定义与 unicode.IsPrint 相同：字母、数字、标点符号、符号和 ASCII 空格都是可打印的。

##### IsPrint Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := strconv.IsPrint('\u263a')
	fmt.Println(c)

	bel := strconv.IsPrint('\007')
	fmt.Println(bel)

}
Output:

true
false
```

#### func Itoa 

``` go 
func Itoa(i int) string
```

​	Itoa函数将 int 类型的 i 转换成对应的十进制字符串。

​	Itoa函数等同于FormatInt(int64(i), 10)。

##### Itoa Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", s, s)

}
Output:

string, 10
```

#### func ParseBool 

``` go 
func ParseBool(str string) (bool, error)
```

​	ParseBool 函数将字符串 str 解析为 bool 类型的值。它接受 1、t、T、TRUE、true、True、0、f、F、FALSE、false、False，其他任何值都会返回一个错误。

##### ParseBool Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "true"
	if s, err := strconv.ParseBool(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

}
Output:

bool, true
```

#### func ParseComplex  <- go1.15

``` go 
func ParseComplex(s string, bitSize int) (complex128, error)
```

​	ParseComplex函数将字符串s转换为复数，精度由bitSize指定：64表示complex64，128表示complex128。当bitSize=64时，结果仍为complex128类型，但可以转换为complex64而不改变其值。

​	s表示为N、Ni或N±Ni的形式，其中N表示由ParseFloat函数识别的浮点数，i是虚部。如果第二个N是无符号的，则需要+符号将两个组件连接起来，如±所示。如果第二个N是NaN，则只接受+符号。该形式可以括在括号中，不能包含任何空格。由ParseFloat函数转换的两个组件构成的结果复数。

​	ParseComplex函数返回的错误具有具体类型`*NumError`，并包括err.Num = s。

​	如果s的语法不正确，则ParseComplex函数返回err.Err = ErrSyntax。

​	如果s的语法正确，但任一组件距离给定组件大小的最大浮点数超过1/2 ULP，则ParseComplex函数返回err.Err = ErrRange和c =±Inf，分别对应于组件。

#### func ParseFloat 

``` go 
func ParseFloat(s string, bitSize int) (float64, error)
```

​	ParseFloat函数将字符串s转换为浮点数，精度由bitSize指定：32表示float32，64表示float64。当bitSize=32时，结果仍为float64类型，但可以转换为float32而不改变其值。

​	ParseFloat函数接受十进制和十六进制的浮点数，这是Go语法对浮点数字面的定义。如果s是格式良好且接近有效的浮点数，ParseFloat函数会返回最近的浮点数，并使用IEEE754无偏差四舍五入。(解析十六进制浮点数时，只有当十六进制表示的位数多于尾数时才会进行四舍五入)。

​	ParseFloat函数返回的错误具有具体类型`*NumError`，并包括err.Num = s。

​	如果s的语法不正确，则ParseFloat函数返回err.Err = ErrSyntax。

​	如果s的语法正确，但距离给定大小的最大浮点数超过1/2 ULP，则ParseFloat函数返回f =±Inf，err.Err = ErrRange。

​	ParseFloat函数将字符串"NaN"和(可能带符号的)字符串"Inf"和"Infinity"识别为它们各自的特殊浮点值。它在匹配时忽略大小写。

##### ParseFloat Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v32 := "-354634382"
	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	v64 := "-3546343826724305832"
	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

}

```

#### func ParseInt 

``` go 
func ParseInt(s string, base int, bitSize int) (i int64, err error)
```

​	ParseInt函数将给定进制(0、2到36)和位大小(0到64)的字符串s解释为对应的值i并返回。

​	字符串可能以符号"+"或"-"开头。

​	如果base参数为0，则真实基数是由字符串前缀随后的符号(如果存在)隐含指定的：对于"0b"，基数为2；对于"0"或"0o"，基数为8；对于"0x"，基数为16；否则基数为10。此外，仅针对base为0的情况，根据Go整数字面量的语法，下划线字符是允许的。

​	bitSize参数指定结果必须适合的整数类型。位大小0、8、16、32和64对应于int、int8、int16、int32和int64。如果bitSize小于0或大于64，则会返回错误。

​	ParseInt函数返回的错误具有具体类型`*NumError`，并包括err.Num = s。如果s为空或包含无效数字，则err.Err = ErrSyntax，返回值为0；如果s对应的值无法由给定大小的有符号整数表示，则err.Err = ErrRange，返回值为适当的bitSize和符号的最大幅度整数。

##### ParseInt Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "42"
	if s, err := strconv.ParseUint(v, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseUint(v, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

}
Output:

uint64, 42
uint64, 42
```

#### func ParseUint 

``` go 
func ParseUint(s string, base int, bitSize int) (uint64, error)
```

​	ParseUint函数类似于ParseInt函数，但用于无符号数字。

​	不允许符号前缀。

##### ParseUint Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// This string literal contains a tab character.
	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)

}
Output:

"\"Fran & Freddie's Diner\t☺\""
```

#### func Quote 

``` go 
func Quote(s string) string
```

​	Quote函数返回表示字符串s的双引号Go字符串文本。返回的字符串使用Go转义序列(\t，\n，\xFF，\u0100)表示控制字符和非可打印字符，如IsPrint定义的。

##### Quote Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// This string literal contains a tab character.
	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)

}
Output:

"\"Fran & Freddie's Diner\t☺\""
```

#### func QuoteRune 

``` go 
func QuoteRune(r rune) string
```

​	QuoteRune函数返回表示符文的单引号Go字符文本。返回的字符串使用Go转义序列(\t，\n，\xFF，\u0100)表示控制字符和非可打印字符，如IsPrint函数定义的。如果r不是有效的Unicode代码点，则将其解释为Unicode替换字符U+FFFD。

> `\u`和`\U`转义字符表示Unicode编码时，后面必须跟着4个(`\u`)或8个(`\U`)十六进制数字来表示该字符的Unicode编码，例如`\u4E2D`表示中文字符"中"的Unicode编码，`\U0001F600`表示笑脸表情的Unicode编码。需要注意的是，`\u`只能用来表示码位在BMP(基本多文种平面，即Unicode的第0平面)中的字符，而`\U`可以用来表示码位在任意平面的字符。
>
> ​	U+数字表示Unicode编码时，U后面紧跟着的数字表示该字符的Unicode编码值，例如U+4E2D表示中文字符"中"的Unicode编码，U+1F600表示笑脸表情的Unicode编码。需要注意的是，U+数字只能用来表示码位在BMP中的字符，而不能表示码位在其他平面中的字符，需要使用`\U`或者专门的非BMP编码方式来表示。
>
> ​	总之，`\u`和`\U`是表示Unicode编码的通用转义字符，适用于任何字符，而U+数字只适用于码位在BMP中的字符。

> BMP(Basic Multilingual Plane，基本多文种平面)是 Unicode 标准中的一个字符编码平面，其中包含了 0x0000 至 0xFFFF 这 65536 个字符的编码。这个编码范围内包含了大部分常用的字符，如 ASCII 码中的字符、拉丁字母、希腊字母、西里尔字母、汉字等等。因此，BMP 平面也被称为 Unicode 字符集的核心区域。
>
> 在 Go 语言中，使用 `\u`+四位十六进制数来表示 BMP 平面中的 Unicode 字符。例如，`\u4e2d`表示汉字"中"的 Unicode 编码 U+4E2D。而使用 `\U`+八位十六进制数来表示 Unicode 字符，这种表示方式可以用来表示 BMP 平面以外的字符，例如 `\U0001F600` 表示一个笑脸表情"😀"，它的 Unicode 编码为 U+1F600，超出了 BMP 平面的编码范围。

##### QuoteRune Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteRune('☺')
	fmt.Println(s)

}
Output:

'☺'
```

#### func QuoteRuneToASCII 

``` go 
func QuoteRuneToASCII(r rune) string
```

​	QuoteRuneToASCII函数返回表示符文的 Go 单引号字符字面量。返回的字符串使用 Go 转义序列 (\t, \n, \xFF, \u0100) 表示非 ASCII 字符和由 IsPrint 定义的不可打印字符。如果 r 不是一个有效的 Unicode 代码点，则将其解释为 Unicode 替换字符 U+FFFD。

##### QuoteRuneToASCII Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteRuneToASCII('☺')
	fmt.Println(s)

}
Output:

'\u263a'
```

#### func QuoteRuneToGraphic  <- go1.6

``` go 
func QuoteRuneToGraphic(r rune) string
```

​	QuoteRuneToGraphic函数返回表示符文的 Go 单引号字符字面量。如果符文不是一个 Unicode 图形字符，如由 IsGraphic函数定义，返回的字符串将使用 Go 转义序列 (\t, \n, \xFF, \u0100)。如果 r 不是一个有效的 Unicode 代码点，则将其解释为 Unicode 替换字符 U+FFFD。

##### QuoteRuneToGraphic Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteRuneToGraphic('☺')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('\u263a')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('\u000a')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('	') // tab character
	fmt.Println(s)

}
Output:

'☺'
'☺'
'\n'
'\t'
```

#### func QuoteToASCII 

``` go 
func QuoteToASCII(s string) string
```

​	QuoteToASCII函数返回表示字符串 s 的 Go 双引号字符串字面量。返回的字符串使用 Go 转义序列 (\t, \n, \xFF, \u0100) 表示非 ASCII 字符和由 IsPrint函数定义的不可打印字符。

##### QuoteToASCII Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 这个字符串字面量包含一个制表符。
	s := strconv.QuoteToASCII(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)

}
Output:

"\"Fran & Freddie's Diner\t\u263a\""
```

#### func QuoteToGraphic  <- go1.6

``` go 
func QuoteToGraphic(s string) string
```

​	QuoteToGraphic函数返回表示字符串 s 的 Go 双引号字符串字面量。返回的字符串保留 Unicode 图形字符(由 IsGraphic函数定义)，对于非图形字符使用 Go 转义序列 (\t, \n, \xFF, \u0100)。

##### QuoteToGraphic Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteToGraphic("☺")
	fmt.Println(s)

	// This string literal contains a tab character.
	s = strconv.QuoteToGraphic("This is a \u263a	\u000a")
	fmt.Println(s)

	s = strconv.QuoteToGraphic(`" This is a ☺ \n "`)
	fmt.Println(s)

}
Output:

"☺"
"This is a ☺\t\n"
"\" This is a ☺ \\n \""
```

#### func QuotedPrefix  <- go1.17

``` go 
func QuotedPrefix(s string) (string, error)
```

​	QuotedPrefix函数返回 s 的前缀处的带引号字符串(如 Unquote函数理解的那样)。如果 s 不以有效的带引号字符串开头，则 QuotedPrefix函数返回一个错误。

#### func Unquote 

``` go 
func Unquote(s string) (string, error)
```

​	Unquote函数将 s 解释为单引号、双引号或反引号包裹的 Go 字符串字面量，并返回 s 引用的字符串值。(如果 s 是单引号引用的，则它是一个 Go 字符字面量；Unquote函数返回相应的单个字符字符串。)

> ​	Unquote 函数接收一个字符串参数 s，并返回一个解析后的字符串值和一个错误。如果解析成功，则该错误为 nil；否则，错误包含一个具体的错误消息。
>
> ​	Unquote 函数支持三种引号类型：单引号、双引号和反引号。其中，单引号表示 Go 语言字符字面值，双引号表示 Go 语言字符串字面值，反引号表示 Go 语言原始字符串字面值。
>
> ​	在解析字符串字面值时，Unquote 函数会自动处理转义字符，例如 \t、\n、" 和 ' 等。同时，它还支持 Unicode 转义，例如 \uXXXX 和 \UXXXXXXXX。

##### Unquote Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
    s, err := strconv.Unquote(`"You can unquote a string with quotes"`)
	fmt.Printf("%q, %v\n", s, err)
	s, err := strconv.Unquote("You can't unquote a string without quotes")
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("\"The string must be either double-quoted\"")
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("`or backquoted.`")
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("'\u263a'") // single character only allowed in single quotes
	fmt.Printf("%q, %v\n", s, err)
	s, err = strconv.Unquote("'\u2639\u2639'")
	fmt.Printf("%q, %v\n", s, err)

}
Output:
"You can unquote a string with quotes",<nil>
"", invalid syntax
"The string must be either double-quoted", <nil>
"or backquoted.", <nil>
"☺", <nil>
"", invalid syntax
```

#### func UnquoteChar 

``` go 
func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
```

​	UnquoteChar函数解码字符串 s 中转义的字符串或字符字面值表示的第一个字符或字节。它返回四个值：

1. value，解码后的 Unicode 码点或字节值； 
2. multibyte，一个布尔值，指示解码后的字符是否需要多字节 UTF-8 表示形式； 
3. tail，字符之后剩余的字符串； 
4. 一个错误，如果字符在语法上有效，则为 nil。

​	第二个参数 quote 指定要解析的文本类型，因此允许哪个转义引号字符。如果设置为单引号，则允许序列 `'`，并禁止未转义的`'`。如果设置为双引号，则允许 `"` 并禁止未转义的`"`。如果设置为零值，则不允许任何转义，且允许两个引号字符未转义出现。

##### UnquoteChar Example
``` go 
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	v, mb, t, err := strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("value:", string(v))
	fmt.Println("multibyte:", mb)
	fmt.Println("tail:", t)
    // 解码转义字符
    s := `\"hello\"`
    value, multibyte, tail, err := strconv.UnquoteChar(s, '"')
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(value, multibyte, tail)
	// 解码Unicode字符
	s = `\u0048\u0065\u006c\u006c\u006f\u0020\u0057\u006f\u0072\u006c\u0064`
    value, multibyte, tail, err = strconv.UnquoteChar(s, '\'')
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%c %t %s", value, multibyte, tail)

}
Output:

value: "
multibyte: false
tail: Fran & Freddie's Diner\"
34 false hello\"
H true \u0065\u006c\u006c\u006f\u0020\u0057\u006f\u0072\u006c\u0064
```

## 类型

### type NumError 

``` go 
type NumError struct {
	Func string // 失败的函数(ParseBool, ParseInt, ParseUint, ParseFloat, ParseComplex)
	Num  string // 输入的数据
	Err  error  // 转换失败的原因(例如 ErrRange, ErrSyntax, 等)。
}
```

​	NumError函数记录了转换失败的情况。

##### Example
``` go 
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "Not a number"
	if _, err := strconv.ParseFloat(str, 64); err != nil {
		e := err.(*strconv.NumError)
		fmt.Println("Func:", e.Func)
		fmt.Println("Num:", e.Num)
		fmt.Println("Err:", e.Err)
		fmt.Println(err)
	}

}
Output:

Func: ParseFloat
Num: Not a number
Err: invalid syntax
strconv.ParseFloat: parsing "Not a number": invalid syntax
```

#### (*NumError) Error 

``` go 
func (e *NumError) Error() string
```

#### (*NumError) Unwrap  <- go1.14

``` go 
func (e *NumError) Unwrap() error
```