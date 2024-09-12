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

​	详见：[/comm/Go/basic/installs]({{< ref "/comm/Go/basic/installs" >}})

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

​	来源：[https://docs.python.org/zh-cn/3.12/reference/lexical_analysis.html#keywords](https://docs.python.org/zh-cn/3.12/reference/lexical_analysis.html#keywords)

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

​	详见 [comm/Go/basic/operators]({{< ref "/comm/Go/basic/operators">}})

​	Go语言比较特殊的操作符有：

- `&^`以及`&^=`，按位清除，按位清除赋值
- `++`和`--`是语句，且只有后置`++`、`--`，而没有前置`++`、`--`
- `:=`用于短变量声明
- `->`用在channel，目前有两种用法：用在函数参数、返回值类型；直接用在channel变量上，用于从接收channel变量的值，以及向channel发送指定值；
- `...`目前有两种用法：用在函数、方法的声明上的最后一个参数，表示可以接受多个参数，在函数和方法体重可以通过这个参数获得一个完整切片；用着函数或方法调用上，用于向函数、方法传递实参，目前可以在切片和字符串类型的变量或字面量上使用。

{{% /tab  %}}

{{% tab header="Python" %}}

​	 详见 [comm/Python/basic/operators]({{< ref "/comm/Python/basic/operators">}})

​	Python语言比较特殊的操作符有：

- 

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