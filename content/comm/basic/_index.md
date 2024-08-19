+++
title = "基础"
date = 2024-03-01T15:18:46+08:00
weight = -100
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

## 安装、配置、卸载、更新、运行代码

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}

​	

{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 关键字、保留字

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}

​	Go目前有25个关键字。

来源：[https://go.dev/ref/spec#Keywords](https://go.dev/ref/spec#Keywords)

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```



{{% /tab  %}}

{{% tab header="Python" %}}

​	Python目前有35个关键字。

来源：[https://docs.python.org/zh-cn/3.12/reference/lexical_analysis.html#keywords](https://docs.python.org/zh-cn/3.12/reference/lexical_analysis.html#keywords)

```python
False      await      else       import     pass
None       break      except     in         raise
True       class      finally    is         return
and        continue   for        lambda     try
as         def        from       nonlocal   while
assert     del        global     not        with
async      elif       if         or         yield
```

{{% /tab  %}}

{{% tab header="Java" %}}

来源：[https://docs.oracle.com/javase/specs/jls/se21/html/jls-3.html#jls-3.9](https://docs.oracle.com/javase/specs/jls/se21/html/jls-3.html#jls-3.9)

*ReservedKeyword:*

```java
abstract   continue   for          new         switch
assert     default    if           package     synchronized
boolean    do         goto         private     this
break      double     implements   protected   throw
byte       else       import       public      throws
case       enum       instanceof   return      transient
catch      extends    int          short       try
char       final      interface    static      void
class      finally    long         strictfp    volatile
const      float      native       super       while
_ (underscore)
```

*ContextualKeyword:*

```java
exports      opens      requires     uses   yield
module       permits    sealed       var         
non-sealed   provides   to           when        
open         record     transitive   with    
```

{{% /tab  %}}

{{% tab header="Rust" %}}

来源：[https://doc.rust-lang.org/reference/keywords.html](https://doc.rust-lang.org/reference/keywords.html)



{{% /tab  %}}

{{% tab header="C/C++" %}}

来源：[https://en.cppreference.com/w/cpp/keyword](https://en.cppreference.com/w/cpp/keyword)



{{% /tab  %}}

{{% tab header="JavaScript" %}}

来源：[https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar#keywords](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar#keywords)

{{% /tab  %}}

{{% tab header="TypeScript" %}}

来源：

{{% /tab  %}}

{{% tab header="C#" %}}

来源：

{{% /tab  %}}

{{% tab header="Erlang" %}}

来源：

{{% /tab  %}}

{{% tab header="PHP" %}}

来源：[https://www.php.net/manual/en/reserved.keywords.php](https://www.php.net/manual/en/reserved.keywords.php)

```php
__halt_compiler()	abstract	and	array()	as
break	callable	case	catch	class
clone	const	continue	declare	default
die()	do	echo	else	elseif
empty()	enddeclare	endfor	endforeach	endif
endswitch	endwhile	eval()	exit()	extends
final	finally	fn (as of PHP 7.4)	for	foreach
function	global	goto	if	implements
include	include_once	instanceof	insteadof	interface
isset()	list()	match (as of PHP 8.0)	namespace	new
or	print	private	protected	public
readonly (as of PHP 8.1.0) *	require	require_once	return	static
switch	throw	trait	try	unset()
use	var	while	xor	yield
yield from
```

{{% /tab  %}}

{{% tab header="Ruby" %}}

来源：

{{% /tab  %}}

{{< /tabpane >}}

## 操作符和标点符号



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}

操作符和标点符号（48个）：

```go
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

赋值操作符（11个）：

```go
+=
-=
*=
/=
%=
&=
|=
^=
<<=
>>=
&^= 
```

算术操作符（7个）：

```go
+
-
*
/ 除法，若两个操作数都是整数，结果是整数（向0取整）；
若两个操作数都是浮点数，则结果是浮点数；
若一整数一浮点数，则编译报错：invalid operation: 变量1 / 变量2 (mismatched types int and float64)
% 取余
++ Go语言中++是一个语句，且只有后置++，而没有前置++
-- Go语言中——是一个语句，且只有后置——，而没有前置——
```

逻辑操作符（3个）：

```go
&& 逻辑与
|| 逻辑或
!  逻辑非
```

关系操作符（7个）：

```go
==
!=
<
>
=
<=
>=
```

位操作符（5个）：

```go
& 按位与，以二进制数字为例，1100 & 1010 的结果为 1000
| 按位或，以二进制数字为例，1100 | 1010 的结果为 1111
^ 按位异或，以二进制数字为例，1100 ^ 1010 的结果为 0110
&^ 按位清除，以二进制数字为例，1100 &^ 1010 的结果为 0100
   对第一个操作数进行与运算，但仅在第二个操作数的相应位为 0 时保留第一个操作数的相应位。
>> 二进制左移位
<< 二进制右移位

```

指针操作符（2个）：

```go
* 声明指针变量、解引用指针
& 获取变量的地址
```

其他操作符（12个）：

```go
( 左括号，用于运算时分组、import分组、函数或方法的定义以及调用，以及变量定义、常量定义、类型定义时共用var、const、type关键字
) 右括号，用于运算时分组、import分组、函数或方法的定义以及调用，以及变量定义、常量定义、类型定义时共用var、const、type关键字
[ 左方括号，用于切片、数组的定义，以及切片、数组获取指定索引的值、map获取指定键的值时使用。
] 右方括号，用于切片、数组的定义，以及切片、数组获取指定索引的值、map获取指定键的值时使用。
{ 用于切片、数组、map等字面值，以及函数、方法、接口、结构体的定义
} 用于切片、数组、map等字面值，以及函数、方法、接口、结构体的定义
:= 用于短变量声明
. 调用实例的方法、指定包中的函数或方法
, 逗号分隔符
; 分号，位于语句末尾，则分号可省略，若一行中有多个语句，则中间语句后面的分号不可省略
... 1.放在函数或方法的形参列表的最后一项前，表示该函数或方法接收任意多个实参；
	2.放在append内置函数的形参列表的最后一项后（最后一项必须是字符串类型、切片类型，
	对于切片，则其元素的类型必须与append函数的第一个实参的元素类型一致，
	对于字符串类型，则append函数的第一个实参的元素类型必须是byte类型）
: 1.定义标签后跟的一个冒号（goto语句跳转指定的标签、break指定的标签等）；
  2.case、default语句后跟的一个冒号；
  3.结构体字段标签中的一个冒号；
  4.切片表达式中的一个冒号；

~ 指定基础类型
-> channel操作
| 用于类型联合
```

操作符优先级：

```go
优先级从高到低排列：
（1）一元操作符（右结合性）
+（正号）
-（负号）
!（逻辑非）
^（按位取反）
*（指针解引用）
&（取地址）
<-（接收运算符，用于从通道中接收数据）

（2）乘法、除法和取余（左结合性）
*（乘法）
/（除法）
%（取余）
<<（左移）
>>（右移）
&（按位与）
&^（按位清除）

（3）加法和减法（左结合性）
+（加法）
-（减法）
|（按位或）
^（按位异或）

（4）关系操作符
==（等于）
!=（不等于）
<（小于）
<=（小于等于）
>（大于）
>=（大于等于）

（5）逻辑运算符
&&（逻辑与）
||（逻辑或）


（6）赋值运算符（右结合性）
=（赋值）
+=（加后赋值）
-=（减后赋值）
*=（乘后赋值）
/=（除后赋值）
%=（取余后赋值）
<<=（左移后赋值）
>>=（右移后赋值）
&=（按位与后赋值）
&^=（按位清除后赋值）
|=（按位或后赋值）
^=（按位异或后赋值）

（7）其他运算符（左结合性）
,（逗号，用于多变量声明和函数参数列表）
```

操作符的结合行：

​	参见以上优先级部分。

{{% /tab  %}}

{{% tab header="Python" %}}

操作符（Operator）（48个）：

```python
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
//   >>    //=   >>=    --                .    :
%          %=           ~
**         **=  
@
```

赋值操作符（11个）：

```python
+=
-=
*=
/=
//=
%=
&=
|=
^=
<<=
>>=
&^=
```

算术操作符（7个）：

```python
+
-
*
/
%
++
--
```

逻辑操作符（3个）：

```python
&&
||
!
```

关系操作符（7个）：

```python
==
!=
<
>
=
<=
>=
```

位操作符（5个）：

```python
& 按位与
| 按位或
^ 按位异或
>> 二进制左移位
<< 二进制右移位
```

指针操作符（2个）：

```python
* 声明指针变量、解引用指针
& 获取变量的地址
```

其他操作符（10个）：

```python
( 左括号
) 右括号
[ 左方括号，用于切片、数组的定义，以及切片、数组获取指定索引的值、map获取指定键的值时使用。
] 右方括号，用于切片、数组的定义，以及切片、数组获取指定索引的值、map获取指定键的值时使用。
:= 用于短变量声明
, 逗号分隔符
; 分号，位于语句末尾，则分号可省略，若一行中有多个语句，则中间语句后面的分号不可省略
... 1.放在函数或方法的形参列表的最后一项前，表示该函数或方法接收任意多个实参；
	2.放在append内置函数的形参列表的最后一项后（最后一项必须是字符串类型、切片类型，
	对于切片，则其的元素类型必须与append函数的第一个实参的元素类型一致，
	对于字符串类型，则append函数的第一个实参的元素类型必须是byte类型）
: 1.定义标签后跟的一个冒号（goto语句跳转指定的标签、break指定的标签等）；
  2.case、default语句后跟的一个冒号；
  3.结构体字段标签中的一个冒号；
  4.切片表达式中的一个冒号；

~ 指定基础类型
```

操作符优先级：

```

```

{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 数据类型

### 数据类型

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 类型转换

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 声明和作用域



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 常量



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 变量



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 命名规范

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 控制语句

### 判断语句（选择语句）

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 循环语句

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 内置函数

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 字符串

### 字符串字面值（string literal）

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 格式化字符串

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 字符串常见操作

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}



## 指针



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 模块



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 测试

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}



## 异常和错误

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 继承

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 编码

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 正则表达式

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## I/O操作

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}