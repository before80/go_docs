+++
title = "fmt"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
> 原文：[https://pkg.go.dev/fmt@go1.21.3](https://pkg.go.dev/fmt@go1.21.3)

Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.

​	fmt包实现了类似于C的printf和scanf的格式化I/O输入Input 和输出 Ouput)功能。格式化的"verbs(动词)"来自于C，但更简单。

## 打印 Printing  

### 动词 Verbs 

#### 通用 General

```
%v	以默认格式打印值
	当打印结构体时，加号(%+v)会增加字段名。    
%#v	值的Go语法表示	
%T	值的类型的Go语法表示	
%%	一个字面百分号；不消耗任何值	
```

#### 布尔 Boolean

```
%t	单词true或false
```

#### 整数 Integer

```
%b	二进制	
%c	对应的Unicode码点所表示的字符	
%d	十进制	
%o	八进制	
%O 	带有0o前缀的八进制	
%q	使用Go语法安全转义的单引号字符字面量。	
%x	十六进制，小写字母a-f	
%X	十六进制，大写字母A-F	
%U	Unicode格式：U+1234；与"U+%04X"相同
```

#### 浮点数和复数成分 Floating-point and complex constituents

```
%b	指数为二的幂的无小数科学计数法，
	类似于strconv.FormatFloat的'b'格式，例如-123456p-78	
%e	科学计数法，例如-1.234456e+78	
%E	科学计数法，例如-1.234456E+78	
%f	有小数点但没有指数，例如123.456	
%F	%f的同义词	
%g	大指数的%e，否则是%f。精度将在下面讨论。	
%G	大指数的%E，否则是%F	
%x	十六进制(带有两个十进制指数幂)，例如-0x1.23abcp+20	
%X	大写版%x，例如：-0X1.23ABCP+20	
```

#### 字符串和字节切片(使用这些动词等效) String and slice of bytes (treated equivalently with these verbs)

```
%s	字符串或切片的未解释字节	
%q	使用Go语法安全转义的双引号字符串	
	（
	
	%q（引用字符串）格式指令用于以Go语言的双引号形式输出字符串，其中会直接将可打印字符的可打印字面量输出，而其他不可打印字符则使用转义的形式输出。
	如果使用了+号修饰符，那么只有ASCII字符（从、U+0020到U+007E）会直接输出，而其他字符则以转义字符形式输出。
	如果使用了#修饰符，那么只要在可能的情况下就会输出Go原始字符串，否则输出以双引号引用的字符串。
		摘自《Go语言程序设计 - Mark Summerfield(英)著》
	）
	
    s := "End Ó ré ttlæti♥中国，世界"

	fmt.Printf("%s\n", s) // End Ó ré ttlæti♥中国，世界
	fmt.Printf("%q\n", s) // "End Ó ré ttlæti♥中国，世界"
	fmt.Printf("%+q\n", s) // "End \u00d3 r\u00e9 ttl\u00e6ti\u2665\u4e2d\u56fd\uff0c\u4e16\u754c"
	fmt.Printf("%#q\n", s) // `End Ó ré ttlæti♥中国，世界`
	
	
%x	十六进制，小写字母，每个字节两个字符	
%X	十六进制，大写字母，每个字节两个字符	
```

#### 切片 Slice

```
%p	以基数16表示的第0个元素的地址，带有前导0x
```

#### 指针 Pointer

```
%p	以基数16表示，带有前导0x
	
%b、%d、%o、%x和%X动词也适用于指针，将该值格式化为整数
```

#### %v的默认格式 The default format for %v is

```
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d，使用%#v打印时为%#x
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

For compound objects, the elements are printed using these rules, recursively, laid out like this:

​	对于复合对象，元素使用这些规则递归地打印，以此方式展开：

```
struct:             {field0 field1 ...}
array, slice:       [elem0 elem1 ...]
maps:               map[key1:value1 key2:value2 ...]
pointer to above:   &{}, &[], &map[]
```

Width is specified by an optional decimal number immediately preceding the verb. If absent, the width is whatever is necessary to represent the value. Precision is specified after the (optional) width by a period followed by a decimal number. If no period is present, a default precision is used. A period with no following number specifies a precision of zero. Examples:

​	**宽度**(Width)是在动词之前的可选十进制数指定的。如果没有指定，则宽度为表示该值所需的任何宽度。**精度**(Precision )是在(可选的)宽度之后由一个句点后跟一个十进制数指定的。如果没有句号，则使用默认精度。没有后面的数字的句号指定精度为零。例如：

```
%f     默认宽度，默认精度
	   
%9f    宽度9，默认精度
	   
%.2f   默认宽度，精度为2
	   
%9.2f  宽度9，精度为2
	   
%9.f   宽度9，精度为0
```

Width and precision are measured in units of Unicode code points, that is, runes. (This differs from C's printf where the units are always measured in bytes.) Either or both of the flags may be replaced with the character '*', causing their values to be obtained from the next operand (preceding the one to format), which must be of type int.

​	宽度和精度以Unicode码点为单位衡量，即符文。 (这与C的printf不同，其中单位始终以字节为单位衡量。)其中一个或两个标志可以用字符'`*`'替换，从而导致它们的值从要格式化的下一个操作数(必须是`int`类型的)获取。

```go 
	fmt.Printf("%0*.*f \n", 7, 3, 3.454489123456)  //003.454 
	fmt.Printf("%0*.*f \n", 8, 3, 3.454489123456)  //0003.454
	fmt.Printf("%0*.*f \n", 9, 3, 3.454489123456)  //00003.454  
	fmt.Printf("%0*.*f \n", 10, 3, 3.454489123456) //000003.454  
	fmt.Printf("%0*.*f \n", 11, 3, 3.454489123456) //00000003.454
	fmt.Printf("%0*.*f \n", 12, 3, 3.454489123456) //000000003.454
	fmt.Printf("%0*.*f \n", 13, 3, 3.454489123456) //0000000003.454
	fmt.Printf("%0*.*f \n", 14, 3, 3.454489123456) //00000000003.454
	fmt.Printf("%0*.*f \n", 15, 3, 3.454489123456) //000000000003.454
	fmt.Printf("%0*.*f \n", 16, 3, 3.454489123456) //0000000000003.454

	fmt.Printf("%0*.*f \n", 7, 3, 3.454589123456)  //003.455
	fmt.Printf("%0*.*f \n", 8, 3, 3.454589123456)  //0003.455
	fmt.Printf("%0*.*f \n", 9, 3, 3.454589123456)  //00003.455
	fmt.Printf("%0*.*f \n", 10, 3, 3.454589123456) //000003.455
	fmt.Printf("%0*.*f \n", 11, 3, 3.454589123456) //0000003.455
	fmt.Printf("%0*.*f \n", 12, 3, 3.454589123456) //00000003.455
	fmt.Printf("%0*.*f \n", 13, 3, 3.454589123456) //000000003.455
	fmt.Printf("%0*.*f \n", 14, 3, 3.454589123456) //0000000003.455
	fmt.Printf("%0*.*f \n", 15, 3, 3.454589123456) //00000000003.455
	fmt.Printf("%0*.*f \n", 16, 3, 3.454589123456) //000000000003.455

	fmt.Printf("%0*.*f \n", 7, 3, 3.454689123456)  //003.455
	fmt.Printf("%0*.*f \n", 8, 3, 3.454689123456)  //0003.455
	fmt.Printf("%0*.*f \n", 9, 3, 3.454689123456)  //00003.455
	fmt.Printf("%0*.*f \n", 10, 3, 3.456489123456) //000003.456
	fmt.Printf("%0*.*f \n", 11, 3, 3.456489123456) //0000003.456
	fmt.Printf("%0*.*f \n", 12, 3, 3.456489123456) //00000003.456
	fmt.Printf("%0*.*f \n", 13, 3, 3.456489123456) //000000003.456
	fmt.Printf("%0*.*f \n", 14, 3, 3.456489123456) //0000000003.456
	fmt.Printf("%0*.*f \n", 15, 3, 3.456489123456) //00000000003.456
	fmt.Printf("%0*.*f \n", 16, 3, 3.456489123456) //000000000003.456
```

For most values, width is the minimum number of runes to output, padding the formatted form with spaces if necessary.

​	对于大多数值，宽度是要输出的符文数的最小值，必要时用空格填充格式化形式。

For strings, byte slices and byte arrays, however, precision limits the length of the input to be formatted (not the size of the output), truncating if necessary. Normally it is measured in runes, but for these types when formatted with the %x or %X format it is measured in bytes.

​	对于字符串，字节切片和字节数组，精度限制要格式化的输入的长度(而不是输出的大小)，必要时截断。通常，它以符文为单位衡量，但对于这些类型，如果使用`%x`或`%X`格式进行格式化，则以字节为单位衡量。

For floating-point values, width sets the minimum width of the field and precision sets the number of places after the decimal, if appropriate, except that for %g/%G precision sets the maximum number of significant digits (trailing zeros are removed). For example, given 12.345 the format %6.3f prints 12.345 while %.3g prints 12.3. The default precision for %e, %f and %#g is 6; for %g it is the smallest number of digits necessary to identify the value uniquely.

​	对于浮点值，宽度设置字段的最小宽度，精度设置小数点后的位数(如果适用)，但是对于`%g` 或 `%G`，精度设置最大有效数字的数量(尾随零被删除)。例如，给定12.345，格式`%6.3f`打印12.345，而`%.3g`打印12.3。`%e`，`%f`和`%＃g`的默认精度为6；对于`%g`，它是唯一确定该值的最小数字数。(不好理解，见以下例子)

```go 
	fmt.Printf("%08.1f \n", 123.444) // 000123.4 
	fmt.Printf("%08.1f \n", 123.456) // 000123.5

	fmt.Printf("%0.1f \n", 123.444)  // 123.4
	fmt.Printf("%0.1f \n", 123.456)  // 123.5

	fmt.Printf("%08.2f \n", 123.444) // 00123.44
	fmt.Printf("%08.2f \n", 123.456) // 00123.46

	fmt.Printf("%0.2f \n", 123.444)  // 123.44
	fmt.Printf("%0.2f \n", 123.456)  // 123.46

	fmt.Printf("%08.1g \n", 123.444) // 0001e+02
	fmt.Printf("%08.1g \n", 123.456) // 0001e+02

	fmt.Printf("%0.1g \n", 123.444)  // 1e+02
	fmt.Printf("%0.1g \n", 123.456)  // 1e+02

	fmt.Printf("%08.2g \n", 123.444) // 01.2e+02
	fmt.Printf("%08.2g \n", 123.456) // 01.2e+02

	fmt.Printf("%08.3g \n", 123.444) // 00000123
	fmt.Printf("%08.3g \n", 123.456) // 00000123

	fmt.Printf("%08.4g \n", 123.444) // 000123.4
	fmt.Printf("%08.4g \n", 123.456) // 000123.5

	fmt.Printf("%08.5g \n", 123.444) // 00123.44
	fmt.Printf("%08.5g \n", 123.456) // 00123.46

	fmt.Printf("%09.5g \n", 123.444) // 000123.44
	fmt.Printf("%09.5g \n", 123.456) // 000123.46

	fmt.Printf("%010.6g \n", 123.444) // 000123.444 
	fmt.Printf("%010.6g \n", 123.456) // 000123.456 

	fmt.Printf("%011.6g \n", 123.444) // 0000123.444
	fmt.Printf("%011.6g \n", 123.456) // 0000123.456 
```

For complex numbers, the width and precision apply to the two components independently and the result is parenthesized, so %f applied to 1.2+3.4i produces (1.200000+3.400000i).

​	对于复数，宽度和精度分别应用于两个组成部分，结果用圆括号括起来的，所以`%f`应用于`1.2+3.4i`将产生`(1.200000+3.400000i)`。

When formatting a single integer code point or a rune string (type []rune) with %q, invalid Unicode code points are changed to the Unicode replacement character, U+FFFD, as in strconv.QuoteRune.

​	使用`%q`格式化单个整数码点或符文字符串(类型`[] rune`)时，无效的Unicode码点将更改为Unicode替换字符`U+FFFD`，如strconv.QuoteRune中所述。(不好理解，见以下例子)

```go 
package main

import (
	"fmt"
)

func main() {
	r := []rune("中国\x80")
	for _, v := range r {
		fmt.Printf("%q - %U\n", v, v)
	}
}
```

输出：

```
'中' - U+4E2D
'国' - U+56FD
'�' - U+FFFD
```



#### 其他标志 Other flags

```
'+'	总是为数字值打印一个标记。
	保证在 %q 中输出 ASCII-only(使用 %+q)
	（
	让格式指令在数值前面输出+号或者-号，为字符串输出ASCII字符（别的字符会被转义），为结构体输出其字段名字
	摘自《Go语言程序设计 - Mark Summerfield(英)著》
	）
	
'-'	在右侧填充空格而不是左侧(左对齐字段)
	
'#'	替代格式：
	在二进制(%#b)前添加前导的 0b，
	在八进制(%#o)前添加 0，
	在十六进制(%#x 或 %#X)前添加 0x 或 0X；
	对于 %q，如果 strconv.CanBackquote 返回 true，
	则打印原始的(反引号括起来的)字符串；
	对于 %e、%E、%f、%F、%g 和 %G，总是打印小数点；
	对于 %g 和 %G，不要移除尾随的零；
	对于 %U，如果字符可打印，则打印类似 U+0078 'x' 的形式(%#U)。
	
	（
	%#o输出以0打头的八进制数据；
	%#p输出不含0x打头的指针；
	%#q尽可能以原始字符串的形式输出一个字符串或者[]byte切片（使用反引号），否则输出以双引号引起来的字符串。
	
	摘自《Go语言程序设计 - Mark Summerfield(英)著》
	）
	
	
' '	(space) leave a space for elided sign in numbers (% d);
	put spaces between bytes printing strings or slices in hex (% x, % X)
	=> (空格)为数字中的省略标记留一个空格(% d)。
	在字节之间留出空格，打印字符串或十六进制的切片(% x, % X)
	(空格)在数字中留出一个空格用于省略的符号(% d)；
	在以十六进制形式(% x、% X)打印字符串或字节片时，将字节之间放置空格。
	
	=> fmt.Printf("数字以%% d格式打印后为% d值\n", 123456789) 
	//数字以% d格式打印后为 123456789值
	=> fmt.Printf("数字以%%d格式打印后为%d值\n", 123456789)
	//数字以%d格式打印后为123456789值
	
	=> fmt.Printf("% x\n", "hello world")
	//68 65 6c 6c 6f 20 77 6f 72 6c 64
	=> fmt.Printf("%x\n", "hello world")
	//68656c6c6f20776f726c64
	
	=> fmt.Printf("% x\n", []int{0x12, 0x13, 0x14})
	//[ 12  13  14]
	=> fmt.Printf("%x\n", []int{0x12, 0x13, 0x14})
	//[12 13 14]
	
'0'	在前导位置填充零而不是空格；
	对于数字，这将在符号后面移动填充；
	对于字符串、字节片和字节数组，将被忽略。
	
    f1 := -1.235
	f2 := 1.235

	fmt.Printf("%+09.3f\n", f1) // -0001.235
	fmt.Printf("%+09.3f\n", f2) // +0001.235	
```

Flags are ignored by verbs that do not expect them. For example there is no alternate decimal format, so %#d and %d behave identically.

​	对于不带标志的动词，标志将被忽略。例如，没有替代的十进制格式，因此 %#d 和 %d 的行为相同。

For each Printf-like function, there is also a Print function that takes no format and is equivalent to saying %v for every operand. Another variant Println inserts blanks between operands and appends a newline.

​	对于每个类似 Printf 的函数，也有一个不带格式的 Print 函数，它相当于为每个操作数都使用 %v。另一种变体 Println 在操作数之间插入空格并追加一个换行符。

Regardless of the verb, if an operand is an interface value, the internal concrete value is used, not the interface itself. Thus:

​	无论动词如何，如果操作数是接口值，则使用内部具体值，而不是接口本身。因此：

``` go 
var i interface{} = 23
fmt.Printf("%v\n", i)
```

will print 23.

将打印23。

Except when printed using the verbs %T and %p, special formatting considerations apply for operands that implement certain interfaces. In order of application:

​	除非使用%T和%p这两个格式，否则针对实现特定接口的操作数，会有特殊的格式化考虑。按应用顺序分别为：

1. If the operand is a reflect.Value, the operand is replaced by the concrete value that it holds, and printing continues with the next rule.

2. 如果操作数是reflect.Value，则将其替换为其持有的具体值，并继续打印下一个规则。

3. If an operand implements the Formatter interface, it will be invoked. In this case the interpretation of verbs and flags is controlled by that implementation.

4. 如果操作数实现了Formatter接口，则将调用该接口。在这种情况下，动词和标志的解释由该实现控制。

5. If the %v verb is used with the # flag (%#v) and the operand implements the GoStringer interface, that will be invoked.

6. 如果使用带有#标志(%#v)的%v动词并且操作数实现了GoStringer接口，则将调用它。

   If the format (which is implicitly %v for Println etc.) is valid for a string (%s %q %v %x %X), the following two rules apply:

   如果格式(在Println等中隐式为%v)对于字符串(%s %q %v %x %X)有效，则应用以下两个规则：

7. If an operand implements the error interface, the Error method will be invoked to convert the object to a string, which will then be formatted as required by the verb (if any).
8. 如果操作数实现了错误接口，则将调用Error方法将对象转换为字符串，然后按照动词(如果有)的要求进行格式化。
9.  If an operand implements method String() string, that method will be invoked to convert the object to a string, which will then be formatted as required by the verb (if any).
10. 如果操作数实现了方法String() string，则将调用该方法将对象转换为字符串，然后按照动词(如果有)的要求进行格式化。

For compound operands such as slices and structs, the format applies to the elements of each operand, recursively, not to the operand as a whole. Thus %q will quote each element of a slice of strings, and %6.2f will control formatting for each element of a floating-point array.

​	对于像切片和结构体这样的复合操作数，格式递归地应用于每个操作数的元素，而不是整个操作数。因此，%q将引用字符串切片的每个元素，%6.2f将控制浮点数数组的每个元素的格式。

However, when printing a byte slice with a string-like verb (%s %q %x %X), it is treated identically to a string, as a single item.

​	但是，当使用类似字符串的动词(%s %q %x %X)打印字节切片时，它与字符串一样被视为单个项。

To avoid recursion in cases such as

​	为避免出现递归，例如

``` go 
type X string
func (x X) String() string { return Sprintf("<%s>", x) }
```

convert the value before recurring:

在递归之前将值转换：

``` go 
func (x X) String() string { return Sprintf("<%s>", string(x)) }
```

Infinite recursion can also be triggered by self-referential data structures, such as a slice that contains itself as an element, if that type has a String method. Such pathologies are rare, however, and the package does not protect against them.

​	如果这种类型有String方法，则还可以触发自引用数据结构(例如包含自身作为元素的切片)的无限递归。然而，这种病理情况很少见，而且该包不会保护它们。

When printing a struct, fmt cannot and therefore does not invoke formatting methods such as Error or String on unexported fields.

​	在打印结构体时，fmt不能且因此不会调用未导出字段的格式化方法，例如Error或String。

### 显式参数索引 Explicit argument indexes

In Printf, Sprintf, and Fprintf, the default behavior is for each formatting verb to format successive arguments passed in the call. However, the notation [n] immediately before the verb indicates that the nth one-indexed argument is to be formatted instead. The same notation before a '*' for a width or precision selects the argument index holding the value. After processing a bracketed expression [n], subsequent verbs will use arguments n+1, n+2, etc. unless otherwise directed.

​	在Printf、Sprintf和Fprintf中，默认行为是为每个格式化动词格式化在调用中传递的连续参数。但是，在动词之前的[n]表示将格式化第n个一索引参数。在'`*`'前面的相同符号表示选择包含该值的参数索引。在处理括号表达式[n]后，后续的动词将使用n + 1、n + 2等参数，除非另有指示。

For example,

例如，

```
fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
```

will yield "22 11", while

将生成 "22 11"，而

```
fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
```

equivalent to

等价于

```
fmt.Sprintf("%6.2f", 12.0)
```

will yield " 12.00". Because an explicit index affects subsequent verbs, this notation can be used to print the same values multiple times by resetting the index for the first argument to be repeated:

将产生" 12.00"。因为显式索引会影响后续动词，所以可以使用此符号将第一个要重复的参数的索引重置为多次打印相同的值：

```
fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
```

will yield "16 17 0x10 0x11".

将生成 "16 17 0x10 0x11"。

### 格式错误 Format errors  

If an invalid argument is given for a verb, such as providing a string to %d, the generated string will contain a description of the problem, as in these examples:

​	如果对动词提供了无效的参数，例如在`%d`中提供了字符串，则生成的字符串将包含问题的描述，例如：

```
Wrong type or unknown verb: %!verb(type=value)
	Printf("%d", "hi"):        %!d(string=hi)
Too many arguments: %!(EXTRA type=value)
	Printf("hi", "guys"):      hi%!(EXTRA string=guys)
Too few arguments: %!verb(MISSING)
	Printf("hi%d"):            hi%!d(MISSING)
Non-int for width or precision: %!(BADWIDTH) or %!(BADPREC)
	Printf("%*s", 4.5, "hi"):  %!(BADWIDTH)hi
	Printf("%.*s", 4.5, "hi"): %!(BADPREC)hi
Invalid or invalid use of argument index: %!(BADINDEX)
	Printf("%*[2]d", 7):       %!d(BADINDEX)
	Printf("%.[2]d", 7):       %!d(BADINDEX)
```

All errors begin with the string "%!" followed sometimes by a single character (the verb) and end with a parenthesized description.

​	所有错误都以字符串"`%!`"开头，有时后跟单个字符(动词)，并以括号括起来的描述结尾。

If an Error or String method triggers a panic when called by a print routine, the fmt package reformats the error message from the panic, decorating it with an indication that it came through the fmt package. For example, if a String method calls panic("bad"), the resulting formatted message will look like

​	如果一个Error或String方法在被打印例程调用时触发了panic，fmt包将重新格式化来自panic的错误消息，用指示它来自fmt包的标识进行修饰。例如，如果String方法调用panic("bad")，则生成的格式化消息将如下所示

```
%!s(PANIC=bad)
```

The %!s just shows the print verb in use when the failure occurred. If the panic is caused by a nil receiver to an Error or String method, however, the output is the undecorated string, "<nil>".

​	`%!s`只是显示故障发生时使用的打印动词。然而，如果panic是由于Error或String方法的nil接收器引起的，则输出为未装饰的字符串`"<nil>"`。

### 扫描 Scanning 

An analogous set of functions scans formatted text to yield values. Scan, Scanf and Scanln read from os.Stdin; Fscan, Fscanf and Fscanln read from a specified io.Reader; Sscan, Sscanf and Sscanln read from an argument string.

​	一组类似的函数扫描格式化文本以产生值。Scan、Scanf 和 Scanln 从 os.Stdin 读取；Fscan、Fscanf 和 Fscanln 从指定的 io.Reader 读取；Sscan、Sscanf 和 Sscanln 从参数字符串中读取。

Scan, Fscan, Sscan treat newlines in the input as spaces.

​	Scan、Fscan、Sscan 将输入中的换行符视为空格。

Scanln, Fscanln and Sscanln stop scanning at a newline and require that the items be followed by a newline or EOF.

​	Scanln、Fscanln 和 Sscanln 在遇到换行符时停止扫描，并要求项后面跟随一个换行符或 EOF。

Scanf, Fscanf, and Sscanf parse the arguments according to a format string, analogous to that of Printf. In the text that follows, 'space' means any Unicode whitespace character except newline.

​	Scanf、Fscanf 和 Sscanf 根据格式字符串解析参数，类似于 Printf。在接下来的文本中，"空格"表示除换行符以外的任何 Unicode 空白字符。

In the format string, a verb introduced by the % character consumes and parses input; these verbs are described in more detail below. A character other than %, space, or newline in the format consumes exactly that input character, which must be present. A newline with zero or more spaces before it in the format string consumes zero or more spaces in the input followed by a single newline or the end of the input. A space following a newline in the format string consumes zero or more spaces in the input. Otherwise, any run of one or more spaces in the format string consumes as many spaces as possible in the input. Unless the run of spaces in the format string appears adjacent to a newline, the run must consume at least one space from the input or find the end of the input.

​	在格式字符串中，由 % 字符引入的转换说明符消耗并解析输入；这些说明符在下面有更详细的描述。格式中除 %、空格或换行符外的字符完全消耗该输入字符，该字符必须存在。在格式字符串中，在换行符之前有零个或多个空格的换行符消耗输入中的零个或多个空格，后跟一个单个换行符或输入结束。在格式字符串中，在换行符后面有一个空格消耗输入中的零个或多个空格。否则，在格式字符串中任何一个或多个空格的运行消耗尽可能多的输入中的空格。除非空格的运行在格式字符串中紧邻换行符，否则该运行必须从输入中消耗至少一个空格或找到输入的结尾。

The handling of spaces and newlines differs from that of C's scanf family: in C, newlines are treated as any other space, and it is never an error when a run of spaces in the format string finds no spaces to consume in the input.

​	空格和换行符的处理与 C 的 scanf 家族不同：在 C 中，换行符被视为其他空格一样，当格式字符串中的空格序列在输入中找不到要消耗的空格时，这不是一个错误。

The verbs behave analogously to those of Printf. For example, %x will scan an integer as a hexadecimal number, and %v will scan the default representation format for the value. The Printf verbs %p and %T and the flags # and + are not implemented. For floating-point and complex values, all valid formatting verbs (%b %e %E %f %F %g %G %x %X and %v) are equivalent and accept both decimal and hexadecimal notation (for example: "2.3e+7", "0x4.5p-8") and digit-separating underscores (for example: "3.14159_26535_89793").

​	这些转换说明符的行为类似于 Printf 的行为。例如，%x 将十六进制数作为整数扫描，而 %v 将扫描值的默认表示格式。Printf 的说明符 %p 和 %T 以及标志 # 和 + 未实现。对于浮点数和复数值，所有有效的格式说明符(%b %e %E %f %F %g %G %x %X 和 %v)都是等效的，并且接受十进制和十六进制表示法(例如："2.3e+7"、"0x4.5p-8")和数字分隔符下划线(例如："3.14159_26535_89793")。

Input processed by verbs is implicitly space-delimited: the implementation of every verb except %c starts by discarding leading spaces from the remaining input, and the %s verb (and %v reading into a string) stops consuming input at the first space or newline character.

​	转换说明符处理的输入隐式以空格为分隔符：每个说明符的实现除了 %c 开头会丢弃剩余输入中的前导空格，%s 说明符(以及 %v 读入字符串)在遇到第一个空格或换行符时停止消耗输入。

The familiar base-setting prefixes 0b (binary), 0o and 0 (octal), and 0x (hexadecimal) are accepted when scanning integers without a format or with the %v verb, as are digit-separating underscores.

​	在没有格式或使用%v动词的情况下扫描整数时，接受熟悉的基数设置前缀0b(二进制)、0o和0(八进制)和0x(十六进制)，以及数字分隔下划线。

Width is interpreted in the input text but there is no syntax for scanning with a precision (no %5.2f, just %5f). If width is provided, it applies after leading spaces are trimmed and specifies the maximum number of runes to read to satisfy the verb. For example,

​	宽度在输入文本中解释，但没有用于扫描精度的语法(没有%5.2f，只有%5f)。如果提供了宽度，则在修剪前导空格之后应用它，并指定满足动词所需读取的符文的最大数量。例如，

```
Sscanf(" 1234567 ", "%5s%d", &s, &i)
```

will set s to "12345" and i to 67 while

将s设置为"12345"，将i设置为67，而

```
Sscanf(" 12 34 567 ", "%5s%d", &s, &i)
```

will set s to "12" and i to 34.

将s设置为"12"，将i设置为34。

In all the scanning functions, a carriage return followed immediately by a newline is treated as a plain newline (\r\n means the same as \n).

​	在所有扫描函数中，回车紧随换行符会被视为普通的换行符(\r\n和\n是等效的)。

In all the scanning functions, if an operand implements method Scan (that is, it implements the Scanner interface) that method will be used to scan the text for that operand. Also, if the number of arguments scanned is less than the number of arguments provided, an error is returned.

​	在所有扫描函数中，如果操作数实现了Scan方法(即实现了Scanner接口)，则将使用该方法来扫描该操作数的文本。此外，如果扫描到的参数数量少于提供的参数数量，则会返回错误。

All arguments to be scanned must be either pointers to basic types or implementations of the Scanner interface.

​	所有要扫描的参数必须是基本类型的指针或Scanner接口的实现。

Like Scanf and Fscanf, Sscanf need not consume its entire input. There is no way to recover how much of the input string Sscanf used.

​	与Scanf和Fscanf一样，Sscanf不需要消耗其整个输入。没有办法恢复Sscanf使用了多少输入字符串。

Note: Fscan etc. can read one character (rune) past the input they return, which means that a loop calling a scan routine may skip some of the input. This is usually a problem only when there is no space between input values. If the reader provided to Fscan implements ReadRune, that method will be used to read characters. If the reader also implements UnreadRune, that method will be used to save the character and successive calls will not lose data. To attach ReadRune and UnreadRune methods to a reader without that capability, use bufio.NewReader.

注意：Fscan等函数可以读取它们返回的输入之后的一个字符(符文)，这意味着调用扫描程序的循环可能会跳过部分输入。这通常仅在输入值之间没有空格时才会出现问题。如果提供给Fscan的读取器实现了ReadRune方法，则将使用该方法读取字符。如果读取器还实现了UnreadRune方法，则将使用该方法保存字符，并且连续的调用不会丢失数据。要将ReadRune和UnreadRune方法附加到没有该功能的读取器，请使用bufio.NewReader。

#### Example (Formats)

These examples demonstrate the basics of printing using a format string. Printf, Sprintf, and Fprintf all take a format string that specifies how to format the subsequent arguments. For example, %d (we call that a 'verb') says to print the corresponding argument, which must be an integer (or something containing an integer, such as a slice of ints) in decimal. The verb %v ('v' for 'value') always formats the argument in its default form, just how Print or Println would show it. The special verb %T ('T' for 'Type') prints the type of the argument rather than its value. The examples are not exhaustive; see the package comment for all the details.

​	这些示例演示了使用格式化字符串进行打印的基础知识。Printf、Sprintf和Fprintf都采用格式化字符串，指定如何格式化后续参数。例如，%d(我们称之为"动词")表示要打印相应的参数，该参数必须是十进制整数(或包含整数的内容，例如int切片)。动词%v("v"表示"value")始终以默认形式格式化参数，就像Print或Println显示它一样。特殊动词%T("T"表示"Type")打印参数的类型而不是其值。这些示例并不详尽，有关所有细节，请参见包的注释。

```go 
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
    // A basic set of examples showing that %v is the default format, in this
	// case decimal for integers, which can be explicitly requested with %d;
	// the output is just what Println generates.
	// 这是一组基本示例，展示了%v是默认格式，
    // 对于整数来说是十进制的，可以使用%d明确请求；
	// 输出结果与Println生成的输出相同。
	integer := 23
    // Each of these prints "23" (without the quotes).
	// 这些每个都会输出"23"(不带引号)。
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

    // The special verb %T shows the type of an item rather than its value.
	// 特殊的%T谓词显示项的类型，而不是其值。
	fmt.Printf("%T %T\n", integer, &integer)
	// 结果：int *int

    // Println(x) is the same as Printf("%v\n", x) so we will use only Printf
	// in the following examples. Each one demonstrates how to format values of
	// a particular type, such as integers or strings. We start each format
	// string with %v to show the default output and follow that with one or
	// more custom formats.
	// Println(x)与Printf("％v \ n"，x)相同，
    // 因此在以下示例中只使用Printf。
	// 每个示例演示如何格式化特定类型的值，例如整数或字符串。
	// 我们以%v开头每个格式字符串以显示默认输出，
    // 然后跟随一个或多个自定义格式。

    // Booleans print as "true" or "false" with %v or %t.
	// 布尔值使用%v或%t打印为"true"或"false"。
	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// 结果：true true

    // Integers print as decimals with %v and %d,
	// or in hex with %x, octal with %o, or binary with %b.
	// 整数使用%v和％d打印为十进制，
	// 使用%x以十六进制打印，使用%o以八进制打印，使用%b以二进制打印。
	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// 结果：42 42 2a 52 101010

    // Floats have multiple formats: %v and %g print a compact representation,
	// while %f prints a decimal point and %e uses exponential notation. The
	// format %6.2f used here shows how to set the width and precision to
	// control the appearance of a floating-point value. In this instance, 6 is
	// the total width of the printed text for the value (note the extra spaces
	// in the output) and 2 is the number of decimal places to show.
	// 浮点数有多种格式：％v和％g以紧凑的形式打印，
	// 而％f打印带有小数点的十进制数，％e使用指数表示法。 
    // 这里使用的格式％6.2f显示如何设置宽度和精度以控制浮点值的外观。 
    // 在此实例中，6是所印文本的总宽度(请注意输出中的额外空格)，
    // 2是要显示的小数位数。
	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// 结果：3.141592653589793 3.141592653589793 3.14(3.14)

    // Complex numbers format as parenthesized pairs of floats, with an 'i'
	// after the imaginary part.
	// 复数格式化为带括号的浮点数对，虚部后跟'i'。
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// 结果: (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)

    // Runes are integers but when printed with %c show the character with that
	// Unicode value. The %q verb shows them as quoted characters, %U as a
	// hex Unicode code point, and %#U as both a code point and a quoted
	// printable form if the rune is printable.
	// Rune是整数，但当使用%c格式打印时，
    // 会显示具有该Unicode值的字符。
    // %q动词将它们显示为带引号的字符，
    // %U显示为十六进制Unicode码点，
    // %#U则显示为码点和带引号的可打印形式，
    // 如果rune是可打印的，则同时显示两者。
	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// 结果: 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'

    // Strings are formatted with %v and %s as-is, with %q as quoted strings,
	// and %#q as backquoted strings.
	// 字符串通过%v和%s原样格式化，
    // 通过%q格式化为带引号的字符串，通过%#q格式化为反引号字符串。
	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// 结果: foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`

    // Maps formatted with %v show keys and values in their default formats.
	// The %#v form (the # is called a "flag" in this context) shows the map in
	// the Go source format. Maps are printed in a consistent order, sorted
	// by the values of the keys.
	// 格式化为%v的映射将以其默认格式显示键和值。 
    // %#v形式(#在这种情况下称为"标志")以Go源格式显示映射。 
    // 映射以一致的顺序打印，按键的值排序。
	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// 结果: map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}

    // Structs formatted with %v show field values in their default formats.
	// The %+v form shows the fields by name, while %#v formats the struct in
	// Go source format.
	// 格式化为%v的结构体将以其默认格式显示字段值。 
    // %+v形式按名称显示字段，而%#v格式化结构体的Go源格式。
	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// 结果: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}

    // The default format for a pointer shows the underlying value preceded by
	// an ampersand. The %p verb prints the pointer value in hex. We use a
	// typed nil for the argument to %p here because the value of any non-nil
	// pointer would change from run to run; run the commented-out Printf
	// call yourself to see.
	// 指针的默认格式显示前面加上取地址符号"&"。 
    // %p动词以十六进制打印指针值。
    // 这里我们使用了一个类型化的nil作为%p的参数，
    // 因为任何非nil指针的值都会从一次运行到另一次运行更改; 
    // 请自己运行被注释的Printf调用以查看。
	pointer := &person
	fmt.Printf("%v %p\n", pointer, (*int)(nil))
	// 结果: &{Kim 22} 0x0
	// fmt.Printf("%v %p\n", pointer, pointer)
	// 结果: &{Kim 22} 0x010203 // See comment above.

    // Arrays and slices are formatted by applying the format to each element.
	// 数组和切片会将格式应用到每个元素。
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// 结果: [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)
	// 结果: [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}

    // Byte slices are special. Integer verbs like %d print the elements in
	// that format. The %s and %q forms treat the slice like a string. The %x
	// verb has a special form with the space flag that puts a space between
	// the bytes.
	// 字节切片是特殊的。
    // 像%d这样的整数格式化指示符以那种格式打印每个元素。 
    // %s和%q格式将切片视为字符串。
    // %x指示符有一个特殊的格式，其中的空格标志将一个空格放在字节之间。
	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
	// 结果: [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98

    // Types that implement Stringer are printed the same as strings. Because
	// Stringers return a string, we can print them using a string-specific
	// verb such as %q.
	// 实现了Stringer接口的类型与字符串一样打印。
    // 由于Stringers返回一个字符串，我们可以使用字符串特定的格式指示符(如%q)打印它们。
	now := time.Unix(123456789, 0).UTC() // time.Time实现了fmt.Stringer。
	fmt.Printf("%v %q\n", now, now)
	// 结果: 1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"

}
Output:

23
23
23
int *int
true true
42 42 2a 52 101010
3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00
(110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
128512 128512 😀 '😀' U+1F600 U+1F600 '😀'
foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`
map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}
{Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}
&{Kim 22} 0x0
[Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]
[Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}
[97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98
1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"
```

#### Example (Printers) 

Print, Println, and Printf lay out their arguments differently. In this example we can compare their behaviors. Println always adds blanks between the items it prints, while Print adds blanks only between non-string arguments and Printf does exactly what it is told. Sprint, Sprintln, Sprintf, Fprint, Fprintln, and Fprintf behave the same as their corresponding Print, Println, and Printf functions shown here.

​	Print、Println和Printf在它们的参数布局方面有所不同。在这个示例中，我们可以比较它们的行为。Println始终在它打印的项之间添加空格，而Print仅在非字符串参数之间添加空格，并且Printf完全按照指令执行。Sprint、Sprintln、Sprintf、Fprint、Fprintln和Fprintf的行为与它们对应的Print、Println和Printf函数相同。

```go 
package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := 3.0, 4.0
	h := math.Hypot(a, b)

    // Print inserts blanks between arguments when neither is a string.
    // It does not add a newline to the output, so we add one explicitly.
	// Print函数在非字符串类型参数之间插入空格。
    // 它不会在输出中添加换行符，因此我们需要显式添加一个换行符。
	fmt.Print("The vector (", a, b, ") has length ", h, ".\n")

    // Println always inserts spaces between its arguments,
	// so it cannot be used to produce the same output as Print in this case;
	// its output has extra spaces.
	// Also, Println always adds a newline to the output.
	// Println函数始终在其参数之间插入空格，
    // 因此它不能用于在此情况下生成与Print相同的输出；
    // 它的输出带有额外的空格。
	// 此外，Println函数始终在输出中添加换行符。
	fmt.Println("The vector (", a, b, ") has length", h, ".")

    // Printf provides complete control but is more complex to use.
	// It does not add a newline to the output, so we add one explicitly
	// at the end of the format specifier string.
	// Printf函数提供完全的控制，但使用起来更加复杂。
	// 它不会在输出中添加换行符，
    // 因此我们需要在格式说明符字符串的末尾显式添加一个换行符。
	fmt.Printf("The vector (%g %g) has length %g.\n", a, b, h)

}
Output:

The vector (3 4) has length 5.
The vector ( 3 4 ) has length 5 .
The vector (3 4) has length 5.
```



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Append  <- go1.19

``` go 
func Append(b []byte, a ...any) []byte
```

Append formats using the default formats for its operands, appends the result to the byte slice, and returns the updated slice.

​	Append函数使用操作数的默认格式进行格式化，将结果附加到字节切片中，并返回更新后的切片。

#### Example My Append

```go
package main

import "fmt"

func main() {
	var b []byte
	b = fmt.Append(b, "a", "b", "中国", 1, 0xff, 012, map[string]int{"age": 18})
	fmt.Println(b)
	fmt.Println(string(b))
}

Output:
[97 98 228 184 173 229 155 189 49 32 50 53 53 32 49 48 32 109 97 112 91 97 103 101 58 49 56 93]
ab中国1 255 10 map[age:18]
```



### func Appendf  <- go1.19

``` go 
func Appendf(b []byte, format string, a ...any) []byte
```

Appendf formats according to a format specifier, appends the result to the byte slice, and returns the updated slice.

​	Appendf函数按照格式说明符进行格式化，将结果附加到字节切片中，并返回更新后的切片。

#### Example My Appendf

```go
package main

import "fmt"

func main() {
	var b []byte
	b1 := fmt.Appendf(b, "%q %s %+q %d %x %o %+v", "a", "b", "中国", 1, 0xff, 012, map[string]int{"age": 18})
	b2 := fmt.Appendf(b, "%s %q %s %d %#x %#o %v", "a", "b", "中国", 1, 0xff, 012, map[string]int{"age": 18})
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(string(b1))
	fmt.Println(string(b2))
}
Output:
[34 97 34 32 98 32 34 92 117 52 101 50 100 92 117 53 54 102 100 34 32 49 32 102
102 32 49 50 32 109 97 112 91 97 103 101 58 49 56 93]
[97 32 34 98 34 32 228 184 173 229 155 189 32 49 32 48 120 102 102 32 48 49 50 3
2 109 97 112 91 97 103 101 58 49 56 93]
"a" b "\u4e2d\u56fd" 1 ff 12 map[age:18]
a "b" 中国 1 0xff 012 map[age:18]
```



### func Appendln  <- go1.19

``` go 
func Appendln(b []byte, a ...any) []byte
```

Appendln formats using the default formats for its operands, appends the result to the byte slice, and returns the updated slice. Spaces are always added between operands and a newline is appended.

​	Appendln函数使用操作数的默认格式进行格式化，将结果附加到字节切片中，并返回更新后的切片。在操作数之间始终添加空格，并附加一个换行符。

#### Example My Appendln 

```go
package main

import "fmt"

func main() {
	var b1, b2 []byte
	b1 = fmt.Appendln(b1, "a", "b", "中国", 1, 0xff, 012, map[string]int{"age": 18})
	b2 = fmt.Append(b2, "a", "b", "中国", 1, 0xff, 012, map[string]int{"age": 18})
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(string(b1))
	fmt.Println(string(b1))
	fmt.Println(string(b2))
	fmt.Println(string(b2))
}


Output:
[97 32 98 32 228 184 173 229 155 189 32 49 32 50 53 53 32 49 48 32 109 97 112 91
 97 103 101 58 49 56 93 10]
[97 98 228 184 173 229 155 189 49 32 50 53 53 32 49 48 32 109 97 112 91 97 103 1
01 58 49 56 93]
a b 中国 1 255 10 map[age:18]

a b 中国 1 255 10 map[age:18]

ab中国1 255 10 map[age:18]
ab中国1 255 10 map[age:18]

```



### func Errorf 

``` go 
func Errorf(format string, a ...any) error
```

Errorf formats according to a format specifier and returns the string as a value that satisfies error.

​	Errorf函数按照格式说明符进行格式化，并将字符串作为满足error接口的值返回。

If the format specifier includes a %w verb with an error operand, the returned error will implement an Unwrap method returning the operand. If there is more than one %w verb, the returned error will implement an Unwrap method returning a []error containing all the %w operands in the order they appear in the arguments. It is invalid to supply the %w verb with an operand that does not implement the error interface. The %w verb is otherwise a synonym for %v.

​	如果格式说明符包含一个带有错误操作数的`%w`动词，则返回的错误将实现一个返回该操作数的Unwrap方法。如果有多个`%w`动词，则返回的错误将实现一个返回按出现在参数中的顺序包含所有`%w`操作数的[]error类型的Unwrap方法。为`%w`动词提供未实现[错误接口](../builtin#type-error)的操作数是无效的。否则，`%w`动词是`%v`的同义词。

#### Errorf Example

The Errorf function lets us use formatting features to create descriptive error messages.

​	`Errorf` 函数允许我们使用格式化功能来创建描述性错误消息。

``` go 
package main

import (
	"fmt"
)

func main() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())

}
Output:

user "bueller" (id 17) not found
```

### func FormatString  <- go1.20

``` go 
func FormatString(state State, verb rune) string
```

FormatString returns a string representing the fully qualified formatting directive captured by the State, followed by the argument verb. (State does not itself contain the verb.) The result has a leading percent sign followed by any flags, the width, and the precision. Missing flags, width, and precision are omitted. This function allows a Formatter to reconstruct the original directive triggering the call to Format.

​	FormatString函数返回一个字符串，表示由State捕获的完全限定的格式化指令，后跟操作数verb。(State本身不包含操作数。)结果具有一个前导百分号，后跟任何标志、宽度和精度。缺少的标志、宽度和精度将被省略。此函数允许Formatter重建触发调用Format的原始指令。

### func Fprint 

``` go 
func Fprint(w io.Writer, a ...any) (n int, err error)
```

Fprint formats using the default formats for its operands and writes to w. Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.

​	Fprint函数使用操作数的默认格式进行格式化，并写入w中。当没有一个操作数是字符串时，它们之间添加空格。它返回写入的字节数和遇到的任何写入错误。

#### Fprint Example
``` go 
package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")

}
Output:

Kim is 22 years old.
21 bytes written.
```

### func Fprintf 

``` go 
func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
```

Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.

​	Fprintf函数按照格式说明符对数据进行格式化，并将结果写入w。它返回写入的字节数和遇到的任何写入错误。

#### Fprintf Example
``` go 
package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// The n and err return values from Fprintf are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

}
Output:

Kim is 22 years old.
21 bytes written.
```

### func Fprintln 

``` go 
func Fprintln(w io.Writer, a ...any) (n int, err error)
```

Fprintln formats using the default formats for its operands and writes to w. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.

​	Fprintln函数按照默认格式对数据进行格式化，并将结果写入w。操作数之间总是添加空格，并追加一个换行符。它返回写入的字节数和遇到的任何写入错误。

#### Fprintln Example
``` go 
package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// The n and err return values from Fprintln are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")

}
Output:

Kim is 22 years old.
21 bytes written.
```

### func Fscan 

``` go 
func Fscan(r io.Reader, a ...any) (n int, err error)
```

Fscan scans text read from r, storing successive space-separated values into successive arguments. Newlines count as space. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.

​	Fscan函数从r中读取文本，将连续的以空格分隔的值存储到连续的参数中。换行符也被视为空格。它返回成功扫描的条目数。如果返回值小于参数个数，则err报告失败的原因。

### func Fscanf 

``` go 
func Fscanf(r io.Reader, format string, a ...any) (n int, err error)
```

Fscanf scans text read from r, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully parsed. Newlines in the input must match newlines in the format.

​	Fscanf函数从r中读取文本，按照格式说明符确定的方式，将连续的以空格分隔的值存储到连续的参数中。它返回成功解析的条目数。输入中的换行符必须与格式中的换行符匹配。

#### Fscanf Example
``` go 
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)
}
Output:

5 true gophers
3
```

### func Fscanln 

``` go 
func Fscanln(r io.Reader, a ...any) (n int, err error)
```

Fscanln is similar to Fscan, but stops scanning at a newline and after the final item there must be a newline or EOF.

​	Fscanln函数类似于Fscan，但会在换行符处停止扫描，在最后一个项之后必须有一个换行符或EOF。

#### Fscanln Example
``` go 
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	s := `dmr 1771 1.61803398875
	ken 271828 3.14159`
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	for {
		n, err := fmt.Fscanln(r, &a, &b, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
	}
}
Output:

3: dmr, 1771, 1.618034
3: ken, 271828, 3.141590
```

### func Print 

``` go 
func Print(a ...any) (n int, err error)
```

Print formats using the default formats for its operands and writes to standard output. Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.

​	Print函数按照默认格式对数据进行格式化，并写入标准输出。当两个操作数都不是字符串时，它们之间添加空格。它返回写入的字节数和遇到的任何写入错误。

#### Print Example
``` go 
package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	fmt.Print(name, " is ", age, " years old.\n")

	// It is conventional not to worry about any
	// error returned by Print.

}
Output:

Kim is 22 years old.
```

### func Printf 

``` go 
func Printf(format string, a ...any) (n int, err error)
```

Printf formats according to a format specifier and writes to standard output. It returns the number of bytes written and any write error encountered.

​	Printf函数根据格式说明符格式化并写入标准输出。它返回写入的字节数和任何遇到的写入错误。

#### Printf Example
``` go 
package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.

}
Output:

Kim is 22 years old.
```

### func Println 

``` go 
func Println(a ...any) (n int, err error)
```

Println formats using the default formats for its operands and writes to standard output. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.

​	Println函数根据操作数的默认格式进行格式化并写入标准输出。无论操作数是什么，都会添加空格，并追加一个换行符。它返回写入的字节数和任何遇到的写入错误。

#### Println Example
``` go 
package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	fmt.Println(name, "is", age, "years old.")

	// It is conventional not to worry about any
	// error returned by Println.

}
Output:

Kim is 22 years old.
```

### func Scan 

``` go 
func Scan(a ...any) (n int, err error)
```

Scan scans text read from standard input, storing successive space-separated values into successive arguments. Newlines count as space. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.

​	Scan函数扫描从`标准输入`读取的文本，将连续的以空格分隔的值存储到连续的参数中。换行符会被视为空格。它返回成功扫描的项数。如果它小于参数个数，那么 err 将会报告原因。

### func Scanf 

``` go 
func Scanf(format string, a ...any) (n int, err error)
```

Scanf scans text read from standard input, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why. Newlines in the input must match newlines in the format. The one exception: the verb %c always scans the next rune in the input, even if it is a space (or tab etc.) or newline.

​	Scanf函数扫描从`标准输入`读取的文本，根据`format`将连续的以空格分隔的值存储到连续的参数中。它返回成功解析的项数。如果它小于参数个数，那么 err 将会报告原因。输入中的换行符必须与格式中的换行符相匹配。唯一的例外是，%c 动词总是扫描输入中的下一个符文，即使它是空格(或制表符等)或换行符。

### func Scanln 

``` go 
func Scanln(a ...any) (n int, err error)
```

Scanln is similar to Scan, but stops scanning at a newline and after the final item there must be a newline or EOF.

​	Scanln函数与 Scan函数类似，但会在换行符处停止扫描，在最后一项后必须有一个换行符或 EOF。

### func Sprint 

``` go 
func Sprint(a ...any) string
```

Sprint formats using the default formats for its operands and returns the resulting string. Spaces are added between operands when neither is a string.

​	Sprint函数根据操作数的默认格式进行格式化，并返回生成的字符串。当两个操作数都不是字符串时，将在它们之间添加空格。

#### Sprint Example
``` go 
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, age = "Kim", 22
	s := fmt.Sprint(name, " is ", age, " years old.\n")

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}
Output:

Kim is 22 years old.
```

### func Sprintf 

``` go 
func Sprintf(format string, a ...any) string
```

Sprintf formats according to a format specifier and returns the resulting string.

​	Sprintf函数根据格式说明符格式化并返回一个字符串。

#### Sprintf Example
``` go 
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}
Output:

Kim is 22 years old.
```

### func Sprintln 

``` go 
func Sprintln(a ...any) string
```

Sprintln formats using the default formats for its operands and returns the resulting string. Spaces are always added between operands and a newline is appended.

​	Sprintln函数使用其操作数的默认格式进行格式化，并返回生成的字符串。在操作数之间始终添加空格，并附加换行符。

#### Sprintln Example
``` go 
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, age = "Kim", 22
	s := fmt.Sprintln(name, "is", age, "years old.")

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}
Output:

Kim is 22 years old.
```

### func Sscan 

``` go 
func Sscan(str string, a ...any) (n int, err error)
```

Sscan scans the argument string, storing successive space-separated values into successive arguments. Newlines count as space. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why.

​	Sscan函数扫描参数`字符串`，将连续的以空格分隔的值存储到连续的参数中。换行符视为空格。它返回成功扫描的项数。如果此数小于参数个数，则err会报告原因。

### func Sscanf 

``` go 
func Sscanf(str string, format string, a ...any) (n int, err error)
```

Sscanf scans the argument string, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully parsed. Newlines in the input must match newlines in the format.

​	Sscanf函数扫描参数`字符串`，将根据格式将连续的以空格分隔的值存储到连续的参数中。它返回成功解析的项数。输入中的换行符必须与格式中的换行符匹配。

#### Sscanf  Example
``` go 
package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)

}
Output:

2: Kim, 22
```

### func Sscanln 

``` go 
func Sscanln(str string, a ...any) (n int, err error)
```

Sscanln is similar to Sscan, but stops scanning at a newline and after the final item there must be a newline or EOF.

​	Sscanln函数类似于 Sscan函数，但在换行符处停止扫描，且最后一项后必须有一个换行符或 EOF。

## 类型

### type Formatter 

``` go 
type Formatter interface {
	Format(f State, verb rune)
}
```

Formatter is implemented by any value that has a Format method. The implementation controls how State and rune are interpreted, and may call Sprint(f) or Fprint(f) etc. to generate its output.

​	Formatter 由任何具有 Format 方法的值实现。实现控制如何解释 State 和 rune，并可以调用 Sprint(f) 或 Fprint(f) 等来生成其输出。

### type GoStringer 

``` go 
type GoStringer interface {
	GoString() string
}
```

GoStringer is implemented by any value that has a GoString method, which defines the Go syntax for that value. The GoString method is used to print values passed as an operand to a %#v format.

​	GoStringer 由任何具有 GoString 方法的值实现，该方法定义该值的 Go 语法。GoString 方法用于打印作为 %#v 格式的操作数传递的值。

#### Example
``` go 
package main

import (
	"fmt"
)

// Address has a City, State and a Country.
type Address struct {
	City    string
	State   string
	Country string
}

// Person has a Name, Age and Address.
type Person struct {
	Name string
	Age  uint
	Addr *Address
}

// GoString makes Person satisfy the GoStringer interface.
// The return value is valid Go code that can be used to reproduce the Person struct.
func (p Person) GoString() string {
	if p.Addr != nil {
		return fmt.Sprintf("Person{Name: %q, Age: %d, Addr: &Address{City: %q, State: %q, Country: %q}}", p.Name, int(p.Age), p.Addr.City, p.Addr.State, p.Addr.Country)
	}
	return fmt.Sprintf("Person{Name: %q, Age: %d}", p.Name, int(p.Age))
}

func main() {
	p1 := Person{
		Name: "Warren",
		Age:  31,
		Addr: &Address{
			City:    "Denver",
			State:   "CO",
			Country: "U.S.A.",
		},
	}
	// If GoString() wasn't implemented, the output of `fmt.Printf("%#v", p1)` would be similar to
	// Person{Name:"Warren", Age:0x1f, Addr:(*main.Address)(0x10448240)}
	fmt.Printf("%#v\n", p1)

	p2 := Person{
		Name: "Theia",
		Age:  4,
	}
	// If GoString() wasn't implemented, the output of `fmt.Printf("%#v", p2)` would be similar to
	// Person{Name:"Theia", Age:0x4, Addr:(*main.Address)(nil)}
	fmt.Printf("%#v\n", p2)

}
Output:

Person{Name: "Warren", Age: 31, Addr: &Address{City: "Denver", State: "CO", Country: "U.S.A."}}
Person{Name: "Theia", Age: 4}
```

### type ScanState 

``` go 
type ScanState interface {
    // ReadRune reads the next rune (Unicode code point) from the input.
	// If invoked during Scanln, Fscanln, or Sscanln, ReadRune() will
	// return EOF after returning the first '\n' or when reading beyond
	// the specified width.
	// ReadRune从输入中读取下一个符文(Unicode码点)。
	// 如果在Scanln、Fscanln或Sscanln期间调用，
    // ReadRune()将在返回第一个'\n'或读取超出指定宽度后返回EOF。
	ReadRune() (r rune, size int, err error)
    // UnreadRune causes the next call to ReadRune to return the same rune.
	// UnreadRune导致下一次调用ReadRune返回相同的符文。
	UnreadRune() error
    // SkipSpace skips space in the input. Newlines are treated appropriately
	// for the operation being performed; see the package documentation
	// for more information.
	// SkipSpace跳过输入中的空格。
    // 换行符将根据执行的操作适当处理；有关更多信息，请参见包文档。
	SkipSpace()
    // Token skips space in the input if skipSpace is true, then returns the
	// run of Unicode code points c satisfying f(c).  If f is nil,
	// !unicode.IsSpace(c) is used; that is, the token will hold non-space
	// characters. Newlines are treated appropriately for the operation being
	// performed; see the package documentation for more information.
	// The returned slice points to shared data that may be overwritten
	// by the next call to Token, a call to a Scan function using the ScanState
	// as input, or when the calling Scan method returns.
	// Token跳过输入中的空格(如果skipSpace为真)，
    // 然后返回满足f(c)的Unicode码点c的运行。
    // 如果f为nil，则使用！unicode.IsSpace(c)；
    // 也就是说，标记将包含非空格字符。
    // 换行符将根据执行的操作适当处理；有关更多信息，请参见包文档。
    // 返回的切片指向可以被下一次对Token的调用、
    // 使用ScanState作为输入的扫描函数的调用或调用调用扫描方法时
    // 可能被覆盖的共享数据。
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
    // Width returns the value of the width option and whether it has been set.
	// The unit is Unicode code points.
	// Width返回宽度选项的值以及它是否已设置。
    // 单位是Unicode码点。
	Width() (wid int, ok bool)
    // Because ReadRune is implemented by the interface, Read should never be
	// called by the scanning routines and a valid implementation of
	// ScanState may choose always to return an error from Read.
	// 因为ReadRune是由接口实现的，
    // 所以扫描程序永远不应调用Read，
    // 而ScanState的有效实现可能选择始终从Read中返回错误。
	Read(buf []byte) (n int, err error)
}
```

ScanState represents the scanner state passed to custom scanners. Scanners may do rune-at-a-time scanning or ask the ScanState to discover the next space-delimited token.

​	ScanState表示传递给自定义扫描器的扫描器状态。扫描器可以逐个rune扫描，也可以要求ScanState发现下一个以空格分隔的标记。

### type Scanner 

``` go 
type Scanner interface {
	Scan(state ScanState, verb rune) error
}
```

Scanner is implemented by any value that has a Scan method, which scans the input for the representation of a value and stores the result in the receiver, which must be a pointer to be useful. The Scan method is called for any argument to Scan, Scanf, or Scanln that implements it.

​	Scanner由具有Scan方法的任何值实现，该方法扫描输入以查找值的表示，并将结果存储在接收器中，后者必须是指针才能有用。对于实现Scan方法的任何参数，都将调用Scan，Scanf或Scanln。

### type State 

``` go 
type State interface {
    // Write is the function to call to emit formatted output to be printed.
	// Write是调用以将格式化输出发射到打印机的函数。
	Write(b []byte) (n int, err error)
    // Width returns the value of the width option and whether it has been set.
	// Width返回宽度选项的值以及是否已设置。
	Width() (wid int, ok bool)
    // Precision returns the value of the precision option and whether it has been set.
	// Precision返回精度选项的值以及是否已设置。
	Precision() (prec int, ok bool)

    // Flag reports whether the flag c, a character, has been set.
	// Flag报告标志c，一个字符，是否已设置。
	Flag(c int) bool
}
```

State represents the printer state passed to custom formatters. It provides access to the io.Writer interface plus information about the flags and options for the operand's format specifier.

​	State表示传递给自定义格式化程序的打印机状态。它提供了对io.Writer接口的访问以及有关操作数格式说明符的标志和选项的信息。

### type Stringer 

``` go 
type Stringer interface {
	String() string
}
```

Stringer is implemented by any value that has a String method, which defines the “native” format for that value. The String method is used to print values passed as an operand to any format that accepts a string or to an unformatted printer such as Print.

​	Stringer由任何具有String方法的值实现，该方法定义该值的"原生"格式。String方法用于打印作为接受字符串的任何格式的操作数或打印机传递的未格式化操作数，例如Print。

#### Example
``` go 
package main

import (
	"fmt"
)

// Animal has a Name and an Age to represent an animal.
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
}
Output:

Gopher (2)
```

