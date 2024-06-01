+++
title = "词汇元素"
date = 2023-05-17T09:59:21+08:00
weight = 4
description = ""
isCJKLanguage = true
type = "docs"
math = true
draft = false

+++
## Lexical elements 词汇元素

> 原文：[https://go.dev/ref/spec#Lexical_elements](https://go.dev/ref/spec#Lexical_elements)

### Comments 注释

Comments serve as program documentation. There are two forms:

​	注释服务于程序文档，有两种形式：

1. *Line comments* start with the character sequence `//` and stop at the end of the line.
2. 行注释以字符序列`//`开始，并在行尾结束。
3. *General comments* start with the character sequence `/*` and stop with the first subsequent character sequence `*/`.
4. 通用注释以字符序列`/*`开始，并以随后的第一个字符序列`*/`结束。

A comment cannot start inside a [rune](https://go.dev/ref/spec#Rune_literals) or [string literal](https://go.dev/ref/spec#String_literals), or inside a comment. A general comment containing no newlines acts like a space. Any other comment acts like a newline.

​	注释不能从[rune]({{< ref "/langSpec/LexicalElements#rune-literals-符文字面量">}})或[string literal]({{< ref "/langSpec/LexicalElements#string-literals-字符串字面量">}})开始，也不能从注释内部开始。一个不包含换行符的通用注释就像一个空格。任何其他的注释就像一个换行符。

### Tokens

Tokens form the vocabulary of the Go language. There are four classes: *identifiers*, *keywords*, *operators and punctuation*, and *literals*. *White space*, formed from spaces (U+0020), horizontal tabs (U+0009), carriage returns (U+000D), and newlines (U+000A), is ignored except as it separates tokens that would otherwise combine into a single token. Also, a newline or end of file may trigger the insertion of a [semicolon](https://go.dev/ref/spec#Semicolons). While breaking the input into tokens, the next token is the longest sequence of characters that form a valid token.

​	tokens 构成了Go语言的词汇表。有四个类别：`标识符`、`关键字`、`操作符和标点符号`以及`字面量（literals）`。由空格（U+0020）、水平制表符（U+0009）、回车符（U+000D）和换行符（U+000A）组成的空白空间被忽略，除非它分隔本来会合并成单个标记的标记。此外，换行或文件结束可能会触发插入分号[semicolon](#semicolons-分号) 。当把输入分解为 tokens 时，下一个 token 是形成有效 token 的最长的字符序列。

### Semicolons 分号

The formal syntax uses semicolons `";"` as terminators in a number of productions. Go programs may omit most of these semicolons using the following two rules:

​	正式语法（formal syntax）在许多结果（productions）中使用分号"`;`"作为终止符。Go程序可以通过以下两条规则省略大部分的分号：

a. When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is

​	当输入被分解成 tokens 时，分号会自动插入标记流后，如果某行的最后一个 token 是：

   - an [identifier](https://go.dev/ref/spec#Identifiers)  

   - 一个标识符（[identifier](#identifiers-标识符)）

   - an [integer](https://go.dev/ref/spec#Integer_literals), [floating-point](https://go.dev/ref/spec#Floating-point_literals), [imaginary](https://go.dev/ref/spec#Imaginary_literals), [rune](https://go.dev/ref/spec#Rune_literals), or [string](https://go.dev/ref/spec#String_literals) literal

   - 一个[整数字面量]({{< ref "/langSpec/LexicalElements#integer-literals-整数字面量">}})、[浮点数字面量]({{< ref "/langSpec/LexicalElements#floating-point-literals-浮点数字面量">}})、[虚数字面量]({{< ref "/langSpec/LexicalElements#imaginary-literals-虚数字面量">}})、[符文字面量]({{< ref "/langSpec/LexicalElements#rune-literals-符文字面量">}})或[字符串字面量]({{< ref "/langSpec/LexicalElements#string-literals-字符串字面量">}} )

   - one of the [keywords](https://go.dev/ref/spec#Keywords) `break`, `continue`, `fallthrough`, or `return`

   - `break`、 `continue`、`fallthrough`、 `return`中的任意一个[关键字](#keywords-关键字)

   - one of the [operators and punctuation](https://go.dev/ref/spec#Operators_and_punctuation) `++`, `--`, `)`, `]`, or `}`

   - `++`、`--`、`)`、`]`、 `}`中的任意一个[操作符或标点符号](#operators-and-punctuation-操作符和标点符号)

        

b. To allow complex statements to occupy a single line, a semicolon may be omitted before a closing `")"` or `"}"`.

​	为了允许复杂的语句占用一行，在结尾的"`)`"或"`}`"之前可以省略分号。

To reflect idiomatic use, code examples in this document elide semicolons using these rules.

​	为了响应惯用法（idiomatic use），本文档中的代码示例使用这些规则省略分号。

### Identifiers 标识符

Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

​	标识符命名程序实体，如变量和类型。标识符是一个或多个字母和数字的序列。标识符中的第一个字符必须是字母。

```
identifier = letter { letter | unicode_digit } .
a
_x9
ThisVariableIsExported
αβ
```

Some identifiers are [predeclared](https://go.dev/ref/spec#Predeclared_identifiers).

​	一些标识符是预先声明的（[predeclared]({{< ref "/langSpec/DeclarationsAndScope#predeclared-identifiers--预先声明的标识符">}})）。

### Keywords 关键字

The following keywords are reserved and may not be used as identifiers.

​	以下关键词被作为保留，不能作为标识符使用：

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Operators and punctuation 操作符和标点符号

The following character sequences represent [operators](https://go.dev/ref/spec#Operators) (including [assignment operators](https://go.dev/ref/spec#Assignment_statements)) and punctuation [[Go 1.18]({{< ref "/langSpec/Appendix#go-118">}})]:

​	以下字符序列代表运算符（[operators]({{< ref "/langSpec/Expressions#operators-操作符">}})）（包括赋值运算符（[assignment operators]({{< ref "/langSpec/Statements#assignment-statements-赋值语句">}})））和标点符号：

```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

### Integer literals 整数字面量

An integer literal is a sequence of digits representing an [integer constant](https://go.dev/ref/spec#Constants). An optional prefix sets a non-decimal base: `0b` or `0B` for binary, `0`, `0o`, or `0O` for octal, and `0x` or `0X` for hexadecimal [[Go 1.13]({{< ref "/langSpec/Appendix#go-113">}})]. A single `0` is considered a decimal zero. In hexadecimal literals, letters `a` through `f` and `A` through `F` represent values 10 through 15.

​	整数字面量是代表一个整数常量的数字序列。可选的前缀用于设置非十进制的基数。二进制为`0b`或`0B`，八进制为`0`、`0o`或`0O`，十六进制为`0x`或`0X` [[Go 1.13]({{< ref "/langSpec/Appendix#go-113">}})]。`单一的0被认为是十进制的0`。在十六进制字面量中，字母`a`到`f`和`A`到`F`代表数值10到15。

For readability, an underscore character `_` may appear after a base prefix or between successive digits; such underscores do not change the literal's value.

​	为了便于阅读，下划线字符`_`可以出现在基数前缀之后或连续的数字之间。这种下划线不会改变字面的值。`（个人注释：只有从 go 1.几的版本及以上版本才可以使用）`

```
int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .

decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
binary_digits  = binary_digit { [ "_" ] binary_digit } .
octal_digits   = octal_digit { [ "_" ] octal_digit } .
hex_digits     = hex_digit { [ "_" ] hex_digit } .
```

```go
42
4_2
0600
0_600
0o600
0O600       // second character is capital letter 'O'  => 第二个字符是字母 'O'
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727

_42         // an identifier, not an integer literal => 标识符,非整数字面量42
42_         // invalid: _ must separate successive digits 
	    //=> 无效的：_ 必须分隔连续的数字 （但0x_1FA或 0X_1FA 等字面量却是有效的）
4__2        // invalid: only one _ at a time 
	     //=> 无效的：每次只能使用一个 _（但0x_1FA或 0X_1FA 等字面量却是有效的）
0_xBadFace  // invalid: _ must separate successive digits 
                           //=>  无效的：_ 必须分隔连续的数字（但0x_1FA或 0X_1FA 等字面量却是有效的）
```

### Floating-point literals 浮点数字面量

A floating-point literal is a decimal or hexadecimal representation of a [floating-point constant](https://go.dev/ref/spec#Constants).

​	浮点数字面量是[浮点常量]({{< ref "/langSpec/Constants">}})的十进制或十六进制表示。

A decimal floating-point literal consists of an integer part (decimal digits), a decimal point, a fractional part (decimal digits), and an exponent part (`e` or `E` followed by an optional sign and decimal digits). One of the integer part or the fractional part may be elided; one of the decimal point or the exponent part may be elided. An exponent value exp scales the mantissa (integer and fractional part) by \\(10^{exp}\\) .

​	十进制浮点数字面量由整数部分（integer part）（十进制数字）、小数点（ a radix point）、小数部分（fractional part ）（十进制数字）和指数部分（exponent part ）（`e`或`E`后面有可选的符号和十进制数字）组成。整数部分或小数部分中的一个可以省略；小数点或指数部分中的一个可以省略。指数值 exp 将尾数（mantissa ）（整数和小数部分）按\\(10^{exp}\\) 进行缩放。

A hexadecimal floating-point literal consists of a `0x` or `0X` prefix, an integer part (hexadecimal digits), a radix point, a fractional part (hexadecimal digits), and an exponent part (`p` or `P` followed by an optional sign and decimal digits). One of the integer part or the fractional part may be elided; the radix point may be elided as well, but the exponent part is required. (This syntax matches the one given in IEEE 754-2008 §5.12.3.) An exponent value exp scales the mantissa (integer and fractional part) by \\(2^{exp}\\)  [[Go 1.13]({{< ref "/langSpec/Appendix#go-113">}})].

​	十六进制浮点数字面量由`0x`或`0X`前缀（prefix）、整数部分（integer part）（十六进制数字）、小数点（ a radix point）、小数部分（fractional part ）（十六进制数字）和指数部分（exponent part ）（`p`或`P`后面有可选的符号和`十进制数字`）组成。整数部分或小数部分中的一个可以省略；小数点也可以省略，但指数部分是必须要存在的。(这个语法与`IEEE 754-2008 §5.12.3`中给出的语法一致。) 指数值exp将尾数（mantissa ）(整数和小数部分)按\\(2^{exp}\\)​ 进行缩放。

For readability, an underscore character `_` may appear after a base prefix or between successive digits; such underscores do not change the literal value.

​	为了便于阅读，一个下划线字符`_`可以出现在基数前缀之后或连续的数字之间；这样的下划线不会改变字面量的值。`（个人注释：只有从 go 1.几的版本及以上版本才可以使用）`

```
float_lit         = decimal_float_lit | hex_float_lit .

decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
                    decimal_digits decimal_exponent |
                    "." decimal_digits [ decimal_exponent ] .
decimal_exponent  = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .

hex_float_lit     = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
hex_mantissa      = [ "_" ] hex_digits "." [ hex_digits ] |
                    [ "_" ] hex_digits |
                    "." hex_digits .
hex_exponent      = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .
```

```
0.
72.40
072.40       // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
1_5.         // == 15.0    <=  省略小数部分
0.15e+0_2    // == 15.0  <= 这里有一个可选符号 +   

0x1p-2       // == 0.25  <= 这里有一个可选符号 -   
0x2.p10      // == 2048.0        
0x1.Fp+0     // == 1.9375 <= 这里有一个可选符号 +  
0X.8p-0      // == 0.5 <= 这里有一个可选符号 -
0X_1FFFP-16  // == 0.1249847412109375
0x15e-2      // == 0x15e - 2 (integer subtraction) 整数减法  
	          //<= 相当于 ( 0x15e)  -  2  =  348。
	          //注意这里的 e 是 十六进制表示法中的 e （表示十进制 14），
	          //而不是十进制浮点数字面量中的指数部分前面的 e

0x.p1        // invalid: mantissa has no digits  
	       	 //=> 无效的： 尾数没有数字
1p-2         // invalid: p exponent requires hexadecimal mantissa  
	      	 //=> 无效的：p 指数需要十六进制尾数
0x1.5e-2     // invalid: hexadecimal mantissa requires p exponent 
	         //=> 无效的：十六进制尾数需要 p 指数
1_.5         // invalid: _ must separate successive digits
	     	 //=> 无效的：_ 必须分开连续的数字
1._5         // invalid: _ must separate successive digits
	      	 //=> 无效的：_ 必须分开连续的数字
1.5_e1       // invalid: _ must separate successive digits
	         //=> 无效的：_ 必须分开连续的数字
1.5e_1       // invalid: _ must separate successive digits
	         //=> 无效的：_ 必须分开连续的数字
1.5e1_       // invalid: _ must separate successive digits
	         //=> 无效的：_ 必须分开连续的数字
```

### Imaginary literals 虚数字面量

An imaginary literal represents the imaginary part of a [complex constant](https://go.dev/ref/spec#Constants). It consists of an [integer](https://go.dev/ref/spec#Integer_literals) or [floating-point](https://go.dev/ref/spec#Floating-point_literals) literal followed by the lowercase letter `i`. The value of an imaginary literal is the value of the respective integer or floating-point literal multiplied by the imaginary unit *i* [[Go 1.13]({{< ref "/langSpec/Appendix#go-113">}})]

​	虚数字面量表示一个[复数常量]({{< ref "/langSpec/Constants">}})的虚数部分。它由一个[整数字面量](#integer-literals-整数字面量)或[浮点数的字面量](#floating-point-literals-浮点数字面量)和小写字母`i`组成，虚数字面量的值是各个整数或浮点数字面量的值乘以虚数单位`i` [[Go 1.13]({{< ref "/langSpec/Appendix#go-113">}})]。

```
imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
```

For backward compatibility, an imaginary literal's integer part consisting entirely of decimal digits (and possibly underscores) is considered a decimal integer, even if it starts with a leading `0`. 

​	为了向后兼容，虚数字面量的整数部分完全由十进制数字（可能还有下划线）组成，被认为是一个十进制整数，即使它以前导`0`开始。

```
0i
0123i         // == 123i for backward-compatibility
0o123i        // == 0o123 * 1i == 83i
0xabci        // == 0xabc * 1i == 2748i
0.i
2.71828i
1.e+0i
6.67428e-11i
1E6i
.25i
.12345E+5i
0x1p-2i       // == 0x1p-2 * 1i == 0.25i
```

### Rune literals 符文字面量

A rune literal represents a [rune constant](https://go.dev/ref/spec#Constants), an integer value identifying a Unicode code point. A rune literal is expressed as one or more characters enclosed in single quotes, as in `'x'` or `'\n'`. Within the quotes, any character may appear except newline and unescaped single quote. A single quoted character represents the Unicode value of the character itself, while multi-character sequences beginning with a backslash encode values in various formats.

​	符文字面量表示一个[符文常量]({{< ref "/langSpec/Constants">}})，一个识别Unicode码点的整数值。一个符文字面量表示为一个或多个字符，用单引号包裹起来，如`'x'`或`'\n'`。在引号内，除了换行符（newline）和未转义的单引号（unescaped single quote），任何字符都可以出现。一个单引号字符代表该字符本身的Unicode值，而以`反斜线`开始的多字符序列则以各种格式编码。

The simplest form represents the single character within the quotes; since Go source text is Unicode characters encoded in UTF-8, multiple UTF-8-encoded bytes may represent a single integer value. For instance, the literal `'a'` holds a single byte representing a literal `a`, Unicode U+0061, value `0x61`, while `'ä'` holds two bytes (`0xc3` `0xa4`) representing a literal `a`-dieresis, U+00E4, value `0xe4`.

​	最简单的形式代表引号内的单个字符；由于Go源文本是以UTF-8编码的Unicode字符，多个UTF-8编码的字节可能代表一个整数值。例如，字面意义上的`'a'`持有一个字节，代表字面意义上的`a`，Unicode U+0061，数值为`0x61`，而`'ä'`持有两个字节（`0xc3 0xa4`），代表字面意义上的a-dieresis，U+00E4，数值为`0xe4`。

Several backslash escapes allow arbitrary values to be encoded as ASCII text. There are four ways to represent the integer value as a numeric constant: `\x` followed by exactly two hexadecimal digits; `\u` followed by exactly four hexadecimal digits; `\U` followed by exactly eight hexadecimal digits, and a plain backslash `\` followed by exactly three octal digits. In each case the value of the literal is the value represented by the digits in the corresponding base.

​	一些反斜线转义允许任意的值被编码为ASCII文本。有四种方法可以将整数值表示为数字常量：

1. `\x`后面正好有`两个`十六进制数字；
2. `\u`后面正好有`四个`十六进制数字；
3. `\U`后面正好有`八个`十六进制数字，以及
4. 一个普通的反斜线`\`后面正好有`三个`**八进制**数字。

​	在每种情况下，字面的值都是由相应基数的数字代表的值。

Although these representations all result in an integer, they have different valid ranges. Octal escapes must represent a value between 0 and 255 inclusive. Hexadecimal escapes satisfy this condition by construction. The escapes `\u` and `\U` represent Unicode code points so within them some values are illegal, in particular those above `0x10FFFF` and surrogate halves.

​	虽然这些表示方法都是一个整数，但它们的有效范围不同。八进制转义必须代表0到255之间的值。十六进制转义在结构上满足这一条件。转义`\u`和`\U`代表Unicode码点，所以在它们里面有些值是非法的，特别是那些高于`0x10FFFF`的值和 surrogate halves 的值。

> 个人注释
>
> ​	解释下什么是surrogate halves ？
>
> 以下是来自文心一言的解释：
>
> ​	在计算机科学中，特别是与Unicode字符编码相关时，"surrogate halves"或"surrogates"是一个特定的概念。Unicode标准设计为能够表示世界上几乎所有的书面语言的字符，包括一些非常特殊的字符，如象形文字、表情符号等。为了实现这一点，Unicode使用一个相对较大的编码空间，其中一些区域被指定为"surrogate"区域。
>
> ​	具体来说，Unicode中的"surrogate halves"或"surrogates"是指位于U+D800到U+DFFF之间的码点。这些码点不是用来直接表示字符的，而是被用作表示那些位于U+10000到U+10FFFF之间的字符的"代理"。
>
> ​	为什么需要这种代理机制呢？因为最初的Unicode设计是使用16位来表示每个字符的，这意味着它最多可以表示2^16 = 65536个字符。然而，随着Unicode的不断扩展，很快就发现这个空间不足以表示所有的字符。为了解决这个问题，Unicode引入了一个扩展机制，允许使用两个16位的代码单元（即一个32位的值）来表示一个字符。这就是所谓的"UTF-16"编码。
>
> ​	在UTF-16编码中，U+D800到U+DFFF之间的码点被用作"高代理"（high surrogates）和"低代理"（low surrogates）。一个高代理和一个低代理结合在一起，形成一个"代理对"（surrogate pair），可以表示一个位于U+10000到U+10FFFF之间的字符。
>
> ​	例如，字符"😀"（U+1F600，一个笑脸表情）在UTF-16编码中会被表示为一个代理对：U+D83D（高代理）和U+DE00（低代理）。
>
> ​	总结来说，"surrogate halves"或"surrogates"是Unicode中用于表示那些超出基本16位编码空间的字符的一种机制。它们不是直接表示字符的，而是作为表示这些字符的代理对的一部分。
>
> ```go
> package main
> 
> import (
> 	"fmt"
> 	"unicode/utf16"
> )
> 
> func PrintIt(r ...rune) {
> 	if len(r) > 0 {
> 		for _, ri := range r {
> 			fmt.Printf("%3d,%X -> %q\n", ri, ri, string(ri))
> 		}
> 	}
> }
> 
> func main() {
> 	r1 := '\x61'
> 	r2 := '\x69'
> 	r3 := '\xFF'
> 	fmt.Println("\\x---------------")
> 	PrintIt(r1, r2, r3)
> 
> 	r1 = '\u0061'
> 	r2 = '\u0069'
> 	r3 = '\uFFFF'
> 	fmt.Println("\\u---------------")
> 	PrintIt(r1, r2, r3)
> 
> 	r1 = '\U00000061'
> 	r2 = '\U00000069'
> 	r3 = '\U0010FFFF' // 不得大于 \U0010FFFF
> 	fmt.Println("\\U---------------")
> 	PrintIt(r1, r2, r3)
> 
> 	r1 = '\141'
> 	r2 = '\151'
> 	r3 = '\377' // 不得大于 \377
> 	PrintIt(r1, r2, r3)
> 
> 	fmt.Println("surrogate halves-----")
> 	fmt.Println(string('\U0001F600'))
> 	PrintIt(utf16.EncodeRune('\U0001F600'))
> 
> 	// 代理对的示例
> 	var runes = []rune{0xD7FF, 0xD800, 0xD801, 0xD802, 0xDFFE, 0xDFFF, 0xE000}
> 
> 	for _, r := range runes {
> 		if utf16.IsSurrogate(r) { // 判断 r的码点是否在位于 U+D800到U+DFFF 之间
> 			fmt.Printf("U+%04X is a surrogate pair\n", r)
> 		} else {
> 			fmt.Printf("U+%04X is not a surrogate pair\n", r)
> 		}
> 	}
> }
> 
> \x---------------
>  97,61 -> "a"                 
> 105,69 -> "i"                 
> 255,FF -> "ÿ"                 
> \u---------------             
>  97,61 -> "a"                 
> 105,69 -> "i"                 
> 65535,FFFF -> "\uffff"        
> \U---------------             
>  97,61 -> "a"                 
> 105,69 -> "i"                 
> 1114111,10FFFF -> "\U0010ffff"
>  97,61 -> "a"                 
> 105,69 -> "i"                 
> 255,FF -> "ÿ"                 
> surrogate halves-----         
> 😀                            
> 55357,D83D -> "�"             
> 56832,DE00 -> "�"             
> U+D7FF is not a surrogate pair
> U+D800 is a surrogate pair    
> U+D801 is a surrogate pair    
> U+D802 is a surrogate pair    
> U+DFFE is a surrogate pair    
> U+DFFF is a surrogate pair    
> U+E000 is not a surrogate pair 
> ```
>
> 

After a backslash, certain single-character escapes represent special values:

​	在反斜线之后，某些单字符转义表示特殊的值：

```
\a   U+0007 alert or bell
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed or newline
\r   U+000D carriage return  => 回车
\t   U+0009 horizontal tab => 水平制表符
\v   U+000B vertical tab => 垂直制表符
\\   U+005C backslash => 反斜线
\'   U+0027 single quote  (valid escape only within rune literals) => 单引号(仅在符文字面量内有效转义)
\"   U+0022 double quote  (valid escape only within string literals) => 双引号(仅在字符串字面值内有效转义)
```

An unrecognized character following a backslash in a rune literal is illegal.

​	在符文字面量中的反斜线后面无法识别的字符是非法的。

```
rune_lit         = "'" ( unicode_value | byte_value ) "'" .
unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value       = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value   = `\` "x" hex_digit hex_digit .
little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
                           hex_digit hex_digit hex_digit hex_digit .
escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
```

```
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'\''         // rune literal containing single quote character
'aa'         // illegal: too many characters
'\k'         // illegal: k is not recognized after a backslash
'\xa'        // illegal: too few hexadecimal digits
'\0'         // illegal: too few octal digits
'\400'       // illegal: octal value over 255
'\uDFFF'     // illegal: surrogate half
'\U00110000' // illegal: invalid Unicode code point
```

### String literals 字符串字面量

A string literal represents a [string constant](https://go.dev/ref/spec#Constants) obtained from concatenating a sequence of characters. There are two forms: raw string literals and interpreted string literals.

​	字符串字面量表示通过连接字符序列获得的[字符串常量]({{< ref "/langSpec/Constants">}})。有两种形式：`原始字符串字面量`和`解释字符串字面量`。

Raw string literals are character sequences between back quotes, as in ``foo``. Within the quotes, any character may appear except back quote. The value of a raw string literal is the string composed of the uninterpreted (implicitly UTF-8-encoded) characters between the quotes; in particular, backslashes have no special meaning and the string may contain newlines. Carriage return characters ('\r') inside raw string literals are discarded from the raw string value.

​	原始字符串字面量是`反引号之间`的字符序列，如\`foo\`。在引号内，任何字符都可以出现，除了反引号。原始字符串字面量的值是由引号之间未解释的（隐含的UTF-8编码）字符组成的字符串；特别是，反斜线没有特殊含义，字符串可以包含新行。原始字符串字面量内的回车字符（'\r'）会从原始字符串值中丢弃。

Interpreted string literals are character sequences between double quotes, as in `"bar"`. Within the quotes, any character may appear except newline and unescaped double quote. The text between the quotes forms the value of the literal, with backslash escapes interpreted as they are in [rune literals](https://go.dev/ref/spec#Rune_literals) (except that `\'` is illegal and `\"` is legal), with the same restrictions. The three-digit octal (`\`*nnn*) and two-digit hexadecimal (`\x`*nn*) escapes represent individual *bytes* of the resulting string; all other escapes represent the (possibly multi-byte) UTF-8 encoding of individual *characters*. Thus inside a string literal `\377` and `\xFF` represent a single byte of value `0xFF`=255, while `ÿ`, `\u00FF`, `\U000000FF` and `\xc3\xbf` represent the two bytes `0xc3` `0xbf` of the UTF-8 encoding of character U+00FF.

​	解释字符串字面量是`双引号之间`的字符序列，如 `"bar"`。在引号内，除了换行和未转义的双引号，任何字符都可以出现。引号之间的文字构成了字面量的值，反斜线转义的解释与符文字面量中的解释相同（除了`\'`是非法的，`\"`是合法的），限制也相同。三位数的八进制(`\nnn`)和两位数的十六进制(`\xnn`)转义代表结果字符串的单个字节；所有其他转义表示单个字符的（可能是多字节的）UTF-8编码。因此，在一个字符串字面中，`\377`和`\xFF`代表值为`0xFF`=255的单个字节，而`ÿ`、`\u00FF`、`\U000000FF`和 `\xc3\xbf`代表字符`U+00FF`的UTF-8编码的两个字节`0xc3` ` 0xbf`。

```
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
```

```
`abc`                // same as "abc"
`\n
\n`                  // same as "\\n\n\\n"
"\n"
"\""                 // same as `"`
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // illegal: surrogate half
"\U00110000"         // illegal: invalid Unicode code point
```

These examples all represent the same string:

​	这些示例都表示相同的字符串：

```
"日本語"                                 // UTF-8 input text
`日本語`                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```

If the source code represents a character as two code points, such as a combining form involving an accent and a letter, the result will be an error if placed in a rune literal (it is not a single code point), and will appear as two code points if placed in a string literal.

​	如果源代码将一个字符表示为两个码点，例如涉及重音和字母的组合形式，如果放在符文字面量中，结果将是一个错误（它不是一个单一码点），如果放在字符串字面量中，将显示为两个码点。