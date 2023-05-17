+++
title = "常量"
date = 2023-05-17T09:59:21+08:00
weight = 5
description = ""
isCJKLanguage = true
draft = false
+++
## Constants 常量

> 原文：[https://go.dev/ref/spec#Constants](https://go.dev/ref/spec#Constants)

​	有布尔常量、符文常量、整数常量、浮点常量、复数常量和字符串常量。符文、整数、浮点和复数常量统称为`数值常量`。

​	常量值由一个[符文字面量](../LexicalElements#rune-literals)、[整数字面量](../LexicalElements#integer-literals)、[浮点数字面量](../LexicalElements#float-point-literals)、[虚数字面量](../LexicalElements#imaginary-literals)或[字符串字面量](../LexicalElements#string-literals)，表示常量的标识符，[常量表达式](../Expressions#constant-expressions)，结果为常量的[转换](../Expressions#conversions)，或一些内置函数的结果值表示，如`unsafe.Sizeof`应用于某些值，`cap`或`len`应用于一些表达式，`real`和`imag`应用于复数常量，`complex`应用于数值常量。布尔真值由预先声明的常数`true`和`false`表示。预先声明的标识符`iota`表示一个整数常量。

​	通常，复数常量是[常量表达式](../Expressions#constant-expressions)的一种形式，将在该节中讨论。

​	数值常量表示任意精度的精确值，不会溢出。因此，不存在表示IEEE-754负零、无穷大和非数字值的常量。

​	常量可以是[有类型的](../Types)的或无类型的。`字面常量`、`true`、`false`、`iota`，以及某些只包含无类型的常量操作数的[常量表达式](../Expressions#constant-expressions)是无类型的。

​	常量可以通过[常量声明](../DeclarationsAndScope#constant-declaration)或[转换](../Expressions#conversions)显式地给出类型，也可以在[变量声明](../DeclarationsAndScope#variable-declarations)、[赋值语句](../Statements#assignment-statements) 、作为[表达式](../Expressions)的操作数时，隐式赋予类型。如果常量值不能被[表示](../Types#representability)为相应类型的值，那就是一个错误。如果类型是一个[类型参数](../DeclarationsAndScope#type-parameter-declarations)，常量将被转换为类型参数的一个非常量值。

​	一个无类型常量有一个默认的类型，该类型是在需要类型化值的上下文中隐式转换为的类型，例如，在一个[短变量声明](../DeclarationsAndScope#short-variable-declarations)中，如`i := 0`，没有明确的类型。无类型常量的默认类型分别是`bool`, `rune`, `int`, `float64`, `complex128`或`string`，具体取决于它是一个布尔型常量、rune型常量、整数型常量、浮点型常量、复数型常量还是字符串型常量。

​	实现限制：尽管数值常量在语言中具有任意的精度，但编译器可以使用有限精度的内部表示法来实现它们。也就是说，每个实现都必须：

- 用至少256位来表示整数常量。
- 用至少256位的尾数和至少16位的有符号二进制指数来表示浮点常量，包括复数常量的对应部分。
- 如果不能精确表示一个整数常量，则给出一个错误。
- 如果由于溢出而无法表示一个浮点常量或复数常量，则给出一个错误。
- 如果由于精度的限制，无法表示一个浮点常量或复数常量，则四舍五入到最接近的可表示常量。

​	这些要求既适用于字面常量，也适用于[常量表达式](../Expressions#constant-expressions)的计算结果。