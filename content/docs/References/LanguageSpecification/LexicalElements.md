+++
title = "词汇元素"
date = 2023-05-17T09:59:21+08:00
weight = 4
description = ""
isCJKLanguage = true
draft = false
+++
## Lexical elements 词汇元素

> 原文：[https://go.dev/ref/spec#Lexical_elements](https://go.dev/ref/spec#Lexical_elements)

### Comments 注释

​	注释服务于程序文档，有两种形式：

1. 行注释以字符序列`//`开始，并在行尾结束。
2. 通用注释以字符序列`/*`开始，并以随后的第一个字符序列`*/`结束。

​	注释不能从[rune](#rune-literals)或[string literal](#string-literals)开始，也不能从注释内部开始。一个不包含换行符的通用注释就像一个空格。任何其他的注释就像一个换行符。

### Tokens

​	tokens 构成了Go语言的词汇表。有四个类别：`标识符`、`关键字`、`操作符和标点符号`以及`字面量（literals）`。由空格（U+0020）、水平制表符（U+0009）、回车符（U+000D）和换行符（U+000A）组成的空白空间被忽略，除非它分隔本来会合并成单个标记的标记。此外，换行或文件结束可能会触发插入分号[semicolon](#Semicolons) 。当把输入分解为 tokens 时，下一个 token 是形成有效 token 的最长的字符序列。

### Semicolons 分号

​	正式语法（formal syntax）在许多结果（productions）中使用分号"`;`"作为终止符。Go程序可以通过以下两条规则省略大部分的分号：

a. 当输入被分解成 tokens 时，分号会自动插入标记流后，如果某行的最后一个 token 是：



   - 一个标识符（[identifier](#identifiers)）

   - 一个[整数字面量](#integer-literals)、[浮点数字面量](#floating-point-literals)、[虚数字面量](#imaginary-literals)、[符文字面量](#rune-literals)或[字符串字面量](#string-literals )

   - `break`、 `continue`、`fallthrough`、 `return`中的任意一个[关键字](#keywords)

   - `++`、`--`、`)`、`]`、 `}`中的任意一个[操作符或标点符号](#operators-and-punctuation)

        

b. 为了允许复杂的语句占用一行，在结尾的"`)`"或"`}`"之前可以省略分号。

​	为了响应惯用法（idiomatic use），本文档中的代码示例使用这些规则省略分号。

### Identifiers 标识符

​	标识符命名程序实体，如变量和类型。标识符是一个或多个字母和数字的序列。标识符中的第一个字符必须是字母。

```
identifier = letter { letter | unicode_digit } .
a
_x9
ThisVariableIsExported
αβ
```

一些标识符是预先声明的（[predeclared](../DeclarationsAndScope#predeclared-identifiers)）。

### Keywords 关键字

​	以下关键词被作为保留，不能作为标识符使用：

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Operators and punctuation 操作符和标点符号

​	以下字符序列代表运算符（[operators](../Expressions#Operators)）（包括赋值运算符（[assignment operators](../Statements#assignment-statements)））和标点符号：

```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

### Integer literals 整数字面量

​	整数字面量是代表一个整数常量的数字序列。可选的前缀用于设置非十进制的基数。二进制为`0b`或`0B`，八进制为`0`、`0o`或`0O`，十六进制为`0x`或`0X`。`单一的0被认为是十进制的0`。在十六进制字面量中，字母`a`到`f`和`A`到`F`代表数值10到15。

​	为了便于阅读，下划线字符`_`可以出现在基数前缀之后或连续的数字之间。这种下划线不会改变字面的值。`（只有从 go 1.几的版本及以上版本才可以使用）`

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

​	浮点数字面量是[浮点常量](../Constants)的十进制或十六进制表示。

​	十进制浮点数字面量由整数部分（integer part）（十进制数字）、小数点（ a radix point）、小数部分（fractional part ）（十进制数字）和指数部分（exponent part ）（`e`或`E`后面有可选的符号和十进制数字）组成。整数部分或小数部分中的一个可以省略；小数点或指数部分中的一个可以省略。指数值 exp 将尾数（mantissa ）（整数和小数部分）按$10^{exp}$ 进行缩放。

​	十六进制浮点数字面量由`0x`或`0X`前缀（prefix）、整数部分（integer part）（十六进制数字）、小数点（ a radix point）、小数部分（fractional part ）（十六进制数字）和指数部分（exponent part ）（`p`或`P`后面有可选的符号和`十进制数字`）组成。整数部分或小数部分中的一个可以省略；小数点也可以省略，但指数部分是必须要存在的。(这个语法与`IEEE 754-2008 §5.12.3`中给出的语法一致。) 指数值exp将尾数（mantissa ）(整数和小数部分)按$2^{exp}$ 进行缩放。

​	为了便于阅读，一个下划线字符`_`可以出现在基数前缀之后或连续的数字之间；这样的下划线不会改变字面量的值。`（只有从 go 1.几的版本及以上版本才可以使用）`

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
	     //=> 无效的：十六进制尾数需要 p 指数
1._5         // invalid: _ must separate successive digits
	      //=> 无效的：十六进制尾数需要 p 指数
1.5_e1       // invalid: _ must separate successive digits
	         //=> 无效的：十六进制尾数需要 p 指数
1.5e_1       // invalid: _ must separate successive digits
	         //=> 无效的：十六进制尾数需要 p 指数
1.5e1_       // invalid: _ must separate successive digits
	         //=> 无效的：十六进制尾数需要 p 指数
```

### Imaginary literals 虚数字面量

​	虚数字面量表示一个[复数常量](../Constants)的虚数部分。它由一个[整数字面量](#integer-literals)或[浮点数的字面量](#floating-point-literals)和小写字母`i`组成，虚数字面量的值是各个整数或浮点数字面量的值乘以虚数单位`i`。

```
imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
```

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

​	符文字面量表示一个[符文常量](../Constants)，一个识别Unicode码点的整数值。一个符文字面量表示为一个或多个字符，用单引号包裹起来，如`'x'`或`'\n'`。在引号内，除了换行符（newline）和未转义的单引号（unescaped single quote），任何字符都可以出现。一个单引号字符代表该字符本身的Unicode值，而以`反斜线`开始的多字符序列则以各种格式编码。

​	最简单的形式代表引号内的单个字符；由于Go源文本是以UTF-8编码的Unicode字符，多个UTF-8编码的字节可能代表一个整数值。例如，字面意义上的`'a'`持有一个字节，代表字面意义上的`a`，Unicode U+0061，数值为`0x61`，而`'ä'`持有两个字节（`0xc3 0xa4`），代表字面意义上的a-dieresis，U+00E4，数值为`0xe4`。

​	一些反斜线转义允许任意的值被编码为ASCII文本。有四种方法可以将整数值表示为数字常量：

1. `\x`后面正好有`两个`十六进制数字；
2. `\u`后面正好有`四个`十六进制数字；
3. `\U`后面正好有`八个`十六进制数字，以及
4. 一个普通的反斜线`\`后面正好有`三个`**八进制**数字。

在每种情况下，字面的值都是由相应基数的数字代表的值。

​	虽然这些表示方法都是一个整数，但它们的有效范围不同。八进制转义必须代表0到255之间的值。十六进制转义在结构上满足这一条件。转义`\u`和`\U`代表Unicode码点，所以在它们里面有些值是非法的，特别是那些高于`0x10FFFF`的值和 surrogate halves 的值。

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

在符文字面量中的反斜线后面无法识别的字符是非法的。

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

​	字符串字面量表示通过连接字符序列获得的[字符串常量](../Constants)。有两种形式：`原始字符串字面量`和`解释字符串字面量`。

​	原始字符串字面量是`反引号之间`的字符序列，如\`foo\`。在引号内，任何字符都可以出现，除了反引号。原始字符串字面量的值是由引号之间未解释的（隐含的UTF-8编码）字符组成的字符串；特别是，反斜线没有特殊含义，字符串可以包含新行。原始字符串字面量内的回车字符（'\r'）会从原始字符串值中丢弃。

​	解释字符串字面量是`双引号之间`的字符序列，如 "bar"。在引号内，除了换行和未转义的双引号，任何字符都可以出现。引号之间的文字构成了字面量的值，反斜线转义的解释与符文字面量中的解释相同（除了`\'`是非法的，`\"`是合法的），限制也相同。三位数的八进制(`\nnn`)和两位数的十六进制(`\xnn`)转义代表结果字符串的单个字节；所有其他转义表示单个字符的（可能是多字节的）UTF-8编码。因此，在一个字符串字面中，`\377`和`\xFF`代表值为`0xFF`=255的单个字节，而`ÿ`、`\u00FF`、`\U000000FF`和 `\xc3\xbf`代表字符`U+00FF`的UTF-8编码的两个字节`0xc3 0xbf`。

```
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
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

这些示例都表示相同的字符串：

```
"日本語"                                 // UTF-8 input text
`日本語`                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```

​	如果源代码将一个字符表示为两个码点，例如涉及重音和字母的组合形式，如果放在符文字面量中，结果将是一个错误（它不是一个单一码点），如果放在字符串字面量中，将显示为两个码点。