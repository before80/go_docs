+++
title = "基础部分"
date = 2023-05-20T08:23:25+08:00
description = ""
isCJKLanguage = true
draft = false

+++



# 基础部分

## 环境搭建

## 关键字

[语言规范中的关键字]({{< ref  "/docs/References/LanguageSpecification/LexicalElements#keywords-关键字" >}})

按字母表顺序排序如下：

```go
break      default      func   interface  select
case       defer        go     map        struct
chan       else         goto   package    switch
const      fallthrough  if     range      type
continue   for          import return     var
```

按类型、用途分类如下：



## 保留字

```go

```



## 数据类型

[语言规范中的数据类型]({{<ref "/docs/References/LanguageSpecification/Types">}})

```go
// 布尔类型
bool // 预先声明 true 和 false 两个常量

// 数值型
// 【数值型】与体系结构无关的数值类型
int8      int16      int32   int64
uint8     uint16     uint32  uint64
float32   float64
complex64 complex128
byte // uint8 的别名
rune // int32 的别名
// 【数值型】与体系结构有关的数值类型
int // 32 或 64 位bit
uint // 32 或 64 位bit
uintptr

// 字符串类型
string

// 数组类型
[number]Type

// 切片类型
[]Type

// 结构体类型
struct

// 指针类型
*Type

// 函数类型
func (params) result

// 接口类型
interface

// 字典、映射类型
//【Python：字典 dict】
//【PHP：关联数组 array】
//【Rust：哈希表 HashMap】
//【Ruby：哈希 Hash】
map

// 通道类型
chan

// 可比较类型 Go 1.18引入的预声明类型
comparable
```

`comparable` 类型又是什么？

[官方博客： All your comparable types]({{< ref "/docs/GoBlog/2023/AllYourComparableTypes" >}})

## 控制结构





## 函数



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






