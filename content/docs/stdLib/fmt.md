+++
title = "fmt"
linkTitle = "fmt"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# fmt

[https://pkg.go.dev/fmt@go1.20.1](https://pkg.go.dev/fmt@go1.20.1)

​	fmt包实现了类似于C的printf和scanf的格式化I/O输入Input 和输出 Ouput)功能。格式化的"verbs(动词)"来自于C，但更简单。

#### 打印 Printing  

##### 动词 Verbs 

###### 通用

```
%v	以默认格式打印值
	当打印结构体时，加号(%+v)会增加字段名。    
%#v	值的Go语法表示	
%T	值的类型的Go语法表示	
%%	一个字面百分号；不消耗任何值	
```

###### 布尔

```
%t	单词true或false
```

###### 整数

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

###### 浮点数和复数成分

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

###### 字符串和字节切片(使用这些动词等效)

```
%s	字符串或切片的未解释字节	
%q	使用Go语法安全转义的双引号字符串	
%x	十六进制，小写字母，每个字节两个字符	
%X	十六进制，大写字母，每个字节两个字符	
```

###### 切片

```
%p	以基数16表示的第0个元素的地址，带有前导0x
```

###### 指针

```
%p	以基数16表示，带有前导0x
	
%b、%d、%o、%x和%X动词也适用于指针，将该值格式化为整数
```

###### %v的默认格式

```
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d，使用%#v打印时为%#x
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

对于复合对象，元素使用这些规则递归地打印，以此方式展开：

```
struct:             {field0 field1 ...}
array, slice:       [elem0 elem1 ...]
maps:               map[key1:value1 key2:value2 ...]
pointer to above:   &{}, &[], &map[]
```

​	**宽度**(Width)是在动词之前的可选十进制数指定的。如果没有指定，则宽度为表示该值所需的任何宽度。**精度**(Precision )是在(可选的)宽度之后由一个句点后跟一个十进制数指定的。如果没有句号，则使用默认精度。没有后面的数字的句号指定精度为零。例如：

```
%f     默认宽度，默认精度
	   
%9f    宽度9，默认精度
	   
%.2f   默认宽度，精度为2
	   
%9.2f  宽度9，精度为2
	   
%9.f   宽度9，精度为0
```

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

​	对于大多数值，宽度是要输出的符文数的最小值，必要时用空格填充格式化形式。

​	对于字符串，字节片和字节数组，精度限制要格式化的输入的长度(而不是输出的大小)，必要时截断。通常，它以符文为单位衡量，但对于这些类型，如果使用`%x`或`%X`格式进行格式化，则以字节为单位衡量。

​		对于浮点值，宽度设置字段的最小宽度，精度设置小数点后的位数(如果适用)，但是对于`%g` 或 `%G`，精度设置最大有效数字的数量(尾随零被删除)。例如，给定12.345，格式`%6.3f`打印12.345，而`%.3g`打印12.3。`%e`，`%f`和`%＃g`的默认精度为6；对于`%g`，它是唯一确定该值的最小数字数。(不好理解，见以下例子)

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

​	对于复数，宽度和精度分别应用于两个组成部分，结果用圆括号括起来的，所以`%f`应用于`1.2+3.4i`将产生`(1.200000+3.400000i)`。

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





###### 其他标志

```
'+'	总是为数字值打印一个标记。
	保证在 %q 中输出 ASCII-only(使用 %+q)
	
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
	
```

​	对于不带标志的动词，标志将被忽略。例如，没有替代的十进制格式，因此 %#d 和 %d 的行为相同。

​	对于每个类似 Printf 的函数，也有一个不带格式的 Print 函数，它相当于为每个操作数都使用 %v。另一种变体 Println 在操作数之间插入空格并追加一个换行符。

​	无论动词如何，如果操作数是接口值，则使用内部具体值，而不是接口本身。因此：

``` go 
var i interface{} = 23
fmt.Printf("%v\n", i)
```

将打印23。

​	除非使用%T和%p这两个格式，否则针对实现特定接口的操作数，会有特殊的格式化考虑。按应用顺序分别为：

1. 如果操作数是reflect.Value，则将其替换为其持有的具体值，并继续打印下一个规则。
2. 如果操作数实现了Formatter接口，则将调用该接口。在这种情况下，动词和标志的解释由该实现控制。
3. 如果使用带有#标志(%#v)的%v动词并且操作数实现了GoStringer接口，则将调用它。

​	如果格式(在Println等中隐式为%v)对于字符串(%s %q %v %x %X)有效，则应用以下两个规则：

1. 如果操作数实现了错误接口，则将调用Error方法将对象转换为字符串，然后按照动词(如果有)的要求进行格式化。
2. 如果操作数实现了方法String() string，则将调用该方法将对象转换为字符串，然后按照动词(如果有)的要求进行格式化。

​	对于像切片和结构体这样的复合操作数，格式递归地应用于每个操作数的元素，而不是整个操作数。因此，%q将引用字符串切片的每个元素，%6.2f将控制浮点数数组的每个元素的格式。

​	但是，当使用类似字符串的动词(%s %q %x %X)打印字节切片时，它与字符串一样被视为单个项目。

​	为避免出现递归，例如

``` go 
type X string
func (x X) String() string { return Sprintf("<%s>", x) }
```

​	在递归之前将值转换：

``` go 
func (x X) String() string { return Sprintf("<%s>", string(x)) }
```

​	如果这种类型有String方法，则还可以触发自引用数据结构(例如包含自身作为元素的切片)的无限递归。然而，这种病理情况很少见，而且该包不会保护它们。

​	在打印结构体时，fmt不能且因此不会调用未导出字段的格式化方法，例如Error或String。

#### 显式参数索引

​	在Printf、Sprintf和Fprintf中，默认行为是为每个格式化动词格式化在调用中传递的连续参数。但是，在动词之前的[n]表示将格式化第n个一索引参数。在'`*`'前面的相同符号表示选择包含该值的参数索引。在处理括号表达式[n]后，后续的动词将使用n + 1、n + 2等参数，除非另有指示。

例如，

```
fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
```

将生成 "22 11"，而

```
fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
```

等价于

```
fmt.Sprintf("%6.2f", 12.0)
```

将产生" 12.00"。因为显式索引会影响后续动词，所以可以使用此符号将第一个要重复的参数的索引重置为多次打印相同的值：

```
fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
```

将生成 "16 17 0x10 0x11"。

#### 格式错误 

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

​	所有错误都以字符串"`%!`"开头，有时后跟单个字符(动词)，并以括号括起来的描述结尾。

​	如果一个Error或String方法在被打印例程调用时触发了panic，fmt包将重新格式化来自panic的错误消息，用指示它来自fmt包的标识进行修饰。例如，如果String方法调用panic("bad")，则生成的格式化消息将如下所示

```
%!s(PANIC=bad)
```

​	`%!s`只是显示故障发生时使用的打印动词。然而，如果panic是由于Error或String方法的nil接收器引起的，则输出为未装饰的字符串`"<nil>"`。

#### 扫描

​	一组类似的函数扫描格式化文本以产生值。Scan、Scanf 和 Scanln 从 os.Stdin 读取；Fscan、Fscanf 和 Fscanln 从指定的 io.Reader 读取；Sscan、Sscanf 和 Sscanln 从参数字符串中读取。

​	Scan、Fscan、Sscan 将输入中的换行符视为空格。

​	Scanln、Fscanln 和 Sscanln 在遇到换行符时停止扫描，并要求项后面跟随一个换行符或 EOF。

​	Scanf、Fscanf 和 Sscanf 根据格式字符串解析参数，类似于 Printf。在接下来的文本中，"空格"表示除换行符以外的任何 Unicode 空白字符。

​	在格式字符串中，由 % 字符引入的转换说明符消耗并解析输入；这些说明符在下面有更详细的描述。格式中除 %、空格或换行符外的字符完全消耗该输入字符，该字符必须存在。在格式字符串中，在换行符之前有零个或多个空格的换行符消耗输入中的零个或多个空格，后跟一个单个换行符或输入结束。在格式字符串中，在换行符后面有一个空格消耗输入中的零个或多个空格。否则，在格式字符串中任何一个或多个空格的运行消耗尽可能多的输入中的空格。除非空格的运行在格式字符串中紧邻换行符，否则该运行必须从输入中消耗至少一个空格或找到输入的结尾。

​	空格和换行符的处理与 C 的 scanf 家族不同：在 C 中，换行符被视为其他空格一样，当格式字符串中的空格序列在输入中找不到要消耗的空格时，这不是一个错误。

​	这些转换说明符的行为类似于 Printf 的行为。例如，%x 将十六进制数作为整数扫描，而 %v 将扫描值的默认表示格式。Printf 的说明符 %p 和 %T 以及标志 # 和 + 未实现。对于浮点数和复数值，所有有效的格式说明符(%b %e %E %f %F %g %G %x %X 和 %v)都是等效的，并且接受十进制和十六进制表示法(例如："2.3e+7"、"0x4.5p-8")和数字分隔符下划线(例如："3.14159_26535_89793")。

​	转换说明符处理的输入隐式以空格为分隔符：每个说明符的实现除了 %c 开头会丢弃剩余输入中的前导空格，%s 说明符(以及 %v 读入字符串)在遇到第一个空格或换行符时停止消耗输入。

​	在没有格式或使用%v动词的情况下扫描整数时，接受熟悉的基数设置前缀0b(二进制)、0o和0(八进制)和0x(十六进制)，以及数字分隔下划线。

​	宽度在输入文本中解释，但没有用于扫描精度的语法(没有%5.2f，只有%5f)。如果提供了宽度，则在修剪前导空格之后应用它，并指定满足动词所需读取的符文的最大数量。例如，

```
Sscanf(" 1234567 ", "%5s%d", &s, &i)
```

将s设置为"12345"，将i设置为67，而

```
Sscanf(" 12 34 567 ", "%5s%d", &s, &i)
```

将s设置为"12"，将i设置为34。

​	在所有扫描函数中，回车紧随换行符会被视为普通的换行符(\r\n和\n是等效的)。

​	在所有扫描函数中，如果操作数实现了Scan方法(即实现了Scanner接口)，则将使用该方法来扫描该操作数的文本。此外，如果扫描到的参数数量少于提供的参数数量，则会返回错误。

​	所有要扫描的参数必须是基本类型的指针或Scanner接口的实现。

​	与Scanf和Fscanf一样，Sscanf不需要消耗其整个输入。没有办法恢复Sscanf使用了多少输入字符串。

注意：Fscan等函数可以读取它们返回的输入之后的一个字符(符文)，这意味着调用扫描程序的循环可能会跳过部分输入。这通常仅在输入值之间没有空格时才会出现问题。如果提供给Fscan的读取器实现了ReadRune方法，则将使用该方法读取字符。如果读取器还实现了UnreadRune方法，则将使用该方法保存字符，并且连续的调用不会丢失数据。要将ReadRune和UnreadRune方法附加到没有该功能的读取器，请使用bufio.NewReader。

##### Example (Formats)

​	这些示例演示了使用格式化字符串进行打印的基础知识。Printf、Sprintf和Fprintf都采用格式化字符串，指定如何格式化后续参数。例如，%d(我们称之为"动词")表示要打印相应的参数，该参数必须是十进制整数(或包含整数的内容，例如int切片)。动词%v("v"表示"value")始终以默认形式格式化参数，就像Print或Println显示它一样。特殊动词%T("T"表示"Type")打印参数的类型而不是其值。这些示例并不详尽，有关所有细节，请参见包的注释。

```go 
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// 这是一组基本示例，展示了%v是默认格式，
    // 对于整数来说是十进制的，可以使用%d明确请求；
	// 输出结果与Println生成的输出相同。
	integer := 23
	// 这些每个都会输出"23"(不带引号)。
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	// 特殊的%T谓词显示项目的类型，而不是其值。
	fmt.Printf("%T %T\n", integer, &integer)
	// 结果：int *int

	// Println(x)与Printf("％v \ n"，x)相同，
    // 因此在以下示例中只使用Printf。
	// 每个示例演示如何格式化特定类型的值，例如整数或字符串。
	// 我们以%v开头每个格式字符串以显示默认输出，
    // 然后跟随一个或多个自定义格式。

	// 布尔值使用%v或%t打印为"true"或"false"。
	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// 结果：true true

	// 整数使用%v和％d打印为十进制，
	// 使用%x以十六进制打印，使用%o以八进制打印，使用%b以二进制打印。
	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// 结果：42 42 2a 52 101010

	// 浮点数有多种格式：％v和％g以紧凑的形式打印，
	// 而％f打印带有小数点的十进制数，％e使用指数表示法。 
    // 这里使用的格式％6.2f显示如何设置宽度和精度以控制浮点值的外观。 
    // 在此实例中，6是所印文本的总宽度(请注意输出中的额外空格)，
    // 2是要显示的小数位数。
	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// 结果：3.141592653589793 3.141592653589793 3.14(3.14)

	// 复数格式化为带括号的浮点数对，虚部后跟'i'。
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// 结果: (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)

	// Rune是整数，但当使用%c格式打印时，
    // 会显示具有该Unicode值的字符。
    // %q动词将它们显示为带引号的字符，
    // %U显示为十六进制Unicode码点，
    // %#U则显示为码点和带引号的可打印形式，
    // 如果rune是可打印的，则同时显示两者。
	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// 结果: 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'

	// 字符串通过%v和%s原样格式化，
    // 通过%q格式化为带引号的字符串，通过%#q格式化为反引号字符串。
	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// 结果: foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`

	// 格式化为%v的映射将以其默认格式显示键和值。 
    // %#v形式(#在这种情况下称为"标志")以Go源格式显示映射。 
    // 映射以一致的顺序打印，按键的值排序。
	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// 结果: map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}

	// 格式化为%v的结构体将以其默认格式显示字段值。 
    // %+v形式按名称显示字段，而%#v格式化结构体的Go源格式。
	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// 结果: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}

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

	// 数组和切片会将格式应用到每个元素。
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// 结果: [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)
	// 结果: [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}

	// 字节切片是特殊的。
    // 像%d这样的整数格式化指示符以那种格式打印每个元素。 
    // %s和%q格式将切片视为字符串。
    // %x指示符有一个特殊的格式，其中的空格标志将一个空格放在字节之间。
	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
	// 结果: [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98

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

##### Example (Printers) 

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

	// Print函数在非字符串类型参数之间插入空格。
    // 它不会在输出中添加换行符，因此我们需要显式添加一个换行符。
	fmt.Print("The vector (", a, b, ") has length ", h, ".\n")

	// Println函数始终在其参数之间插入空格，
    // 因此它不能用于在此情况下生成与Print相同的输出；
    // 它的输出带有额外的空格。
	// 此外，Println函数始终在输出中添加换行符。
	fmt.Println("The vector (", a, b, ") has length", h, ".")

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

#### func [Append](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=287)  <- go1.19

``` go 
func Append(b []byte, a ...any) []byte
```

​	Append函数使用操作数的默认格式进行格式化，将结果附加到字节切片中，并返回更新后的切片。

#### func [Appendf](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=247)  <- go1.19

``` go 
func Appendf(b []byte, format string, a ...any) []byte
```

​	Appendf函数按照格式说明符进行格式化，将结果附加到字节切片中，并返回更新后的切片。

#### func [Appendln](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=330)  <- go1.19

``` go 
func Appendln(b []byte, a ...any) []byte
```

​	Appendln函数使用操作数的默认格式进行格式化，将结果附加到字节切片中，并返回更新后的切片。在操作数之间始终添加空格，并附加一个换行符。

#### func [Errorf](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/errors.go;l=22) 

``` go 
func Errorf(format string, a ...any) error
```

​	Errorf函数按照格式说明符进行格式化，并将字符串作为满足error接口的值返回。

​	如果格式说明符包含一个带有错误操作数的`%w`动词，则返回的错误将实现一个返回该操作数的Unwrap方法。如果有多个`%w`动词，则返回的错误将实现一个返回按出现在参数中的顺序包含所有`%w`操作数的[]error类型的Unwrap方法。为`%w`动词提供未实现[错误接口](../builtin#type-error)的操作数是无效的。否则，`%w`动词是`%v`的同义词。

##### Errorf Example
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

#### func [FormatString](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=81)  <- go1.20

``` go 
func FormatString(state State, verb rune) string
```

​	FormatString函数返回一个字符串，表示由State捕获的完全限定的格式化指令，后跟操作数verb。(State本身不包含操作数。)结果具有一个前导百分号，后跟任何标志、宽度和精度。缺少的标志、宽度和精度将被省略。此函数允许Formatter重建触发调用Format的原始指令。

#### func [Fprint](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=260) 

``` go 
func Fprint(w io.Writer, a ...any) (n int, err error)
```

​	Fprint函数使用操作数的默认格式进行格式化，并写入w中。当没有一个操作数是字符串时，它们之间添加空格。它返回写入的字节数和遇到的任何写入错误。

##### Fprint Example
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

#### func [Fprintf](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=222) 

``` go 
func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
```

​	Fprintf函数按照格式说明符对数据进行格式化，并将结果写入w。它返回写入的字节数和遇到的任何写入错误。

##### Fprintf Example
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

#### func [Fprintln](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=302) 

``` go 
func Fprintln(w io.Writer, a ...any) (n int, err error)
```

​	Fprintln函数按照默认格式对数据进行格式化，并将结果写入w。操作数之间总是添加空格，并追加一个换行符。它返回写入的字节数和遇到的任何写入错误。

##### Fprintln Example
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

#### func [Fscan](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=121) 

``` go 
func Fscan(r io.Reader, a ...any) (n int, err error)
```

​	Fscan函数从r中读取文本，将连续的以空格分隔的值存储到连续的参数中。换行符也被视为空格。它返回成功扫描的条目数。如果返回值小于参数个数，则err报告失败的原因。

#### func [Fscanf](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=141) 

``` go 
func Fscanf(r io.Reader, format string, a ...any) (n int, err error)
```

​	Fscanf函数从r中读取文本，按照格式说明符确定的方式，将连续的以空格分隔的值存储到连续的参数中。它返回成功解析的条目数。输入中的换行符必须与格式中的换行符匹配。

##### Fscanf Example
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

#### func [Fscanln](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=130) 

``` go 
func Fscanln(r io.Reader, a ...any) (n int, err error)
```

​	Fscanln函数类似于Fscan，但会在换行符处停止扫描，在最后一个项目之后必须有一个换行符或EOF。

##### Fscanln Example
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

#### func [Print](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=271) 

``` go 
func Print(a ...any) (n int, err error)
```

​	Print函数按照默认格式对数据进行格式化，并写入标准输出。当两个操作数都不是字符串时，它们之间添加空格。它返回写入的字节数和遇到的任何写入错误。

##### Print Example
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

#### func [Printf](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=232) 

``` go 
func Printf(format string, a ...any) (n int, err error)
```

​	Printf函数根据格式说明符格式化并写入标准输出。它返回写入的字节数和任何遇到的写入错误。

##### Printf Example
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

#### func [Println](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=313) 

``` go 
func Println(a ...any) (n int, err error)
```

​	Println函数根据操作数的默认格式进行格式化并写入标准输出。无论操作数是什么，都会添加空格，并追加一个换行符。它返回写入的字节数和任何遇到的写入错误。

##### Println Example
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

#### func [Scan](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=63) 

``` go 
func Scan(a ...any) (n int, err error)
```

​	Scan函数扫描从标准输入读取的文本，将连续的以空格分隔的值存储到连续的参数中。换行符会被视为空格。它返回成功扫描的项数。如果它小于参数数目，那么 err 将会报告原因。

#### func [Scanf](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=80) 

``` go 
func Scanf(format string, a ...any) (n int, err error)
```

​	Scanf函数扫描从标准输入读取的文本，根据格式将连续的以空格分隔的值存储到连续的参数中。它返回成功解析的项数。如果它小于参数数目，那么 err 将会报告原因。输入中的换行符必须与格式中的换行符相匹配。唯一的例外是，%c 动词总是扫描输入中的下一个符文，即使它是空格(或制表符等)或换行符。

#### func [Scanln](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=69) 

``` go 
func Scanln(a ...any) (n int, err error)
```

​	Scanln函数与 Scan函数类似，但会在换行符处停止扫描，在最后一项后必须有一个换行符或 EOF。

#### func [Sprint](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=277) 

``` go 
func Sprint(a ...any) string
```

​	Sprint函数根据操作数的默认格式进行格式化，并返回生成的字符串。当两个操作数都不是字符串时，将它们之间添加空格。

##### Sprint Example
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

#### func [Sprintf](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=237) 

``` go 
func Sprintf(format string, a ...any) string
```

​	Sprintf函数根据格式说明符格式化并返回一个字符串。

##### Sprintf Example
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

#### func [Sprintln](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=319) 

``` go 
func Sprintln(a ...any) string
```

​	Sprintln函数使用其操作数的默认格式进行格式化，并返回生成的字符串。在操作数之间始终添加空格，并附加换行符。

##### Sprintln Example
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

#### func [Sscan](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=99) 

``` go 
func Sscan(str string, a ...any) (n int, err error)
```

​	Sscan函数扫描参数字符串，将连续的以空格分隔的值存储到连续的参数中。换行符视为空格。它返回成功扫描的项目数。如果此数小于参数数，则err会报告原因。

#### func [Sscanf](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=113) 

``` go 
func Sscanf(str string, format string, a ...any) (n int, err error)
```

​	Sscanf函数扫描参数字符串，将根据格式将连续的以空格分隔的值存储到连续的参数中。它返回成功解析的项目数。输入中的换行符必须与格式中的换行符匹配。

##### Sscanf  Example
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

#### func [Sscanln](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=105) 

``` go 
func Sscanln(str string, a ...any) (n int, err error)
```

​	Sscanln函数类似于 Sscan函数，但在换行符处停止扫描，且最后一项后必须有一个换行符或 EOF。

## 类型

### type [Formatter](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=54) 

``` go 
type Formatter interface {
	Format(f State, verb rune)
}
```

​	Formatter 由任何具有 Format 方法的值实现。实现控制如何解释 State 和 rune，并可以调用 Sprint(f) 或 Fprint(f) 等来生成其输出。

### type [GoStringer](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=71) 

``` go 
type GoStringer interface {
	GoString() string
}
```

​	GoStringer 由任何具有 GoString 方法的值实现，该方法定义该值的 Go 语法。GoString 方法用于打印作为 %#v 格式的操作数传递的值。

##### Example
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

### type [ScanState](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=21) 

``` go 
type ScanState interface {
	// ReadRune从输入中读取下一个符文(Unicode码点)。
	// 如果在Scanln、Fscanln或Sscanln期间调用，
    // ReadRune()将在返回第一个'\n'或读取超出指定宽度后返回EOF。
	ReadRune() (r rune, size int, err error)
	// UnreadRune导致下一次调用ReadRune返回相同的符文。
	UnreadRune() error
	// SkipSpace跳过输入中的空格。
    // 换行符将根据执行的操作适当处理；有关更多信息，请参见包文档。
	SkipSpace()
	// Token跳过输入中的空格(如果skipSpace为真)，
    // 然后返回满足f(c)的Unicode码点c的运行。
    // 如果f为nil，则使用！unicode.IsSpace(c)；
    // 也就是说，标记将包含非空格字符。
    // 换行符将根据执行的操作适当处理；有关更多信息，请参见包文档。
    // 返回的切片指向可以被下一次对Token的调用、
    // 使用ScanState作为输入的扫描函数的调用或调用调用扫描方法时
    // 可能被覆盖的共享数据。
	Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
	// Width返回宽度选项的值以及它是否已设置。
    // 单位是Unicode码点。
	Width() (wid int, ok bool)
	// 因为ReadRune是由接口实现的，
    // 所以扫描程序永远不应调用Read，
    // 而ScanState的有效实现可能选择始终从Read中返回错误。
	Read(buf []byte) (n int, err error)
}
```

​	ScanState表示传递给自定义扫描器的扫描器状态。扫描器可以逐个rune扫描，也可以要求ScanState发现下一个以空格分隔的标记。

### type [Scanner](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/scan.go;l=55) 

``` go 
type Scanner interface {
	Scan(state ScanState, verb rune) error
}
```

​	Scanner由具有Scan方法的任何值实现，该方法扫描输入以查找值的表示，并将结果存储在接收器中，后者必须是指针才能有用。对于实现Scan方法的任何参数，都将调用Scan，Scanf或Scanln。

### type [State](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=39) 

``` go 
type State interface {
	// Write是调用以将格式化输出发射到打印机的函数。
	Write(b []byte) (n int, err error)
	// Width返回宽度选项的值以及是否已设置。
	Width() (wid int, ok bool)
	// Precision返回精度选项的值以及是否已设置。
	Precision() (prec int, ok bool)

	// Flag报告标志c，一个字符，是否已设置。
	Flag(c int) bool
}
```

​	State表示传递给自定义格式化程序的打印机状态。它提供了对io.Writer接口的访问以及有关操作数格式说明符的标志和选项的信息。

### type [Stringer](https://cs.opensource.google/go/go/+/go1.20.1:src/fmt/print.go;l=63) 

``` go 
type Stringer interface {
	String() string
}
```

​	Stringer由任何具有String方法的值实现，该方法定义该值的"原生"格式。String方法用于打印作为接受字符串的任何格式的操作数或打印机传递的未格式化操作数，例如Print。

##### Example
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

