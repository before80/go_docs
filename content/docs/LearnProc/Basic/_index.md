+++
title = "基础部分"
date = 2023-05-20T08:23:25+08:00
description = ""
isCJKLanguage = true
draft = false

+++



# 基础部分

## 环境搭建

## 【25个】关键字

[语言规范中的关键字]({{< ref  "/docs/References/LanguageSpecification/LexicalElements#keywords-关键字" >}})

【25个】按字母表顺序排序如下：

```go
break      default      func   interface  select
case       defer        go     map        struct
chan       else         goto   package    switch
const      fallthrough  if     range      type
continue   for          import return     var
```

按类型、用途分类如下（来自《Go语言核心编程》李文塔/著）：

【8个】引导程序整体结构的关键字：

```go
package
import
const
var
func
defer
go
return
```

【4个】声明复合数据结构的关键字：

```go
struct
interface
map
chan
```

【13个】控制程序结构的关键字：

```go
if      else
for     range     break    continue
switch  select    case     fallthrough  default  type
goto
```



## 【41个】预先声明标识符

[语言规范中的预先声明标识符]({{< ref "/docs/References/LanguageSpecification/DeclarationsAndScope#predeclared-identifiers--预先声明的标识符" >}})

【41个】按类型分类如下：

```go
// 【22个】类型:
	any       bool       byte   comparable
	complex64 complex128 error  float32    float64
	int       int8       int16  int32      int64    rune    string
	uint      uint8      uint16 uint32     uint64   uintptr

// 【3个】常量:
	true false iota

// 【1个】零值:
	nil

// 【15个】函数:
	append  cap  close  complex copy     delete  imag     len
	make    new  panic  print   println  real    recover
```

> any：interface{} 的别名。
>
> comparable ：可比较类型， Go 1.18引入的预声明类型。



## 【28种】数据类型

[语言规范中的数据类型]({{<ref "/docs/References/LanguageSpecification/Types">}})

```go
// 【1种】布尔类型
bool // 预先声明 true 和 false 两个常量

// 数值型
// 【14种】【数值型】与体系结构无关的数值类型
int8      int16      int32   int64
uint8     uint16     uint32  uint64
float32   float64
complex64 complex128
byte // uint8 的别名
rune // int32 的别名
// 【3种】【数值型】与体系结构有关的数值类型
int // 32 或 64 位bit
uint // 32 或 64 位bit
uintptr

// 【1种】字符串类型
string

// 【1种】数组类型
[number]Type

// 【1种】切片类型
[]Type

// 【1种】结构体类型
struct

// 【1种】指针类型
*Type

// 【1种】函数类型
func (params) result

// 【1种】接口类型
interface

// 【1种】error 接口类型
error

// 【1种】字典、映射类型
//【Python：字典 dict】
//【PHP：关联数组 array】
//【Rust：哈希表 HashMap】
//【Ruby：哈希 Hash】
map

// 【1种】通道类型
chan

```

`comparable` 类型又是什么？在预先声明标识符中。

[官方博客： All your comparable types]({{< ref "/docs/GoBlog/2023/AllYourComparableTypes" >}})

## 【种】控制结构

```go
if
if...else...
if...else if...else

for ;; {}
for condition {}
for {}
for k,v := range {}

switch...case...default...


```



## 函数



### 【15个】内置函数

[语言规范中的内置函数]({{< ref  "/docs/References/LanguageSpecification/Built-inFunctions" >}})

【15个】按字母表顺序排序如下：

```go
append   delete   panic    
cap      imag     print
close    len      println
complex  make     real
copy     new      recover
```



## 方法



## 接口



## 并发



## 错误



## 异常



## 标准库



## go 命令



## 日志的处理



## GORM



## JSON的处理



## 授权的处理



## 跨域的处理



## Viper的使用



## Cobra的使用



## 测试



### 单元测试



### 基准测试



### 模糊测试



### 示例测试



## //go:build



## //go:embed





