+++
title = "变量"
date = 2023-05-17T09:59:21+08:00
weight = 6
description = ""
isCJKLanguage = true
draft = false
+++
## Variables 变量

> 原文：[https://go.dev/ref/spec#Variables ](https://go.dev/ref/spec#Variables )

​	变量是用于保存值的存储位置。允许值的集合是由变量的[类型](../Types)决定的。

​	[变量声明](../DeclarationsAndScope#variable-declarations) 、函数参数和结果、[函数声明](../DeclarationsAndScope#function-declarations)或[函数字面量](../Expressions#function-literals)的签名为指定的变量保留存储空间。调用内置函数`new`或获取[复合字面量](../Expressions#composite-literals)的地址会在运行时为变量分配存储空间。这样的匿名变量是通过(可能是隐式的)[指针间接](../Expressions#address-operators)引用的。

​	[数组](../Types#array-types)、[切片](../Types#slice-types)和[结构体](../Types/struct-types)等类型的结构化变量具有可以被单独[寻址](../Expressions#address-operators)的元素和字段。每个这样的元素都像一个变量。

​	变量的静态类型（或仅仅是类型）是在其声明中给出的类型、在`new`调用或复合字面量中提供的类型、或是结构化变量的元素的类型。接口类型的变量也有一个独特的动态类型，它是运行时分配给变量的值的（非接口）类型（除非该值是预先声明的标识符`nil`，它没有类型）。在执行过程中，动态类型可能会发生变化，但是存储在接口变量中的值总是可以赋给变量的静态类型。

```go linenums="1"
var x interface{}  // x is nil and has static type interface{}  => x 为 nil，具有静态类型 interface{}
var v *T           // v has value nil, static type *T =>  v 具有值 nil，静态类型 * T
x = 42             // x has value 42 and dynamic type int =>  x 具有值42，动态类型 int
x = v              // x has value (*T)(nil) and dynamic type *T => x 具有值(* T)(nil)和动态类型 * T
```

​	变量的值是通过在[表达式](../Expressions)中引用该变量来检索的；它是最近[分配](../statements#assignment-statements)给该变量的值。如果一个变量还没有被赋值，它的值就是其类型的[零值](../ProgramInitializationAndExecution#the-zero-value)。