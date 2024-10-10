+++
title = "标识符"
date = 2024-07-13T10:56:22+08:00
weight = 100
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

​	在Go的语言规范中[关于标识符的描述]({{< ref "/langSpec/LexicalElements#identifiers-标识符">}})如下：

> Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.
>
>  ​	标识符命名程序实体，如变量和类型。标识符是一个或多个字母和数字的序列。标识符中的第一个字符必须是字母。
>
> ```
> identifier = letter { letter | unicode_digit } .
> a
> _x9
> ThisVariableIsExported
> αβ
> ```
>
> Some identifiers are [predeclared](https://go.dev/ref/spec#Predeclared_identifiers).
>
> ​	一些标识符是预先声明的（[predeclared]({{< ref "/langSpec/DeclarationsAndScope#predeclared-identifiers--预先声明的标识符">}})）。

​	这里也将预先声明的标识符（合计44个）摘录如下：

```
types【共22个】:
	any【go1.18】 bool byte comparable【go1.20】
	complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants【共3个】:
	true false iota

Zero value【共1个】:
	nil

Functions【共18个】:
	append cap clear【go1.21】 close complex copy delete imag len
	make max【go1.21】 min【go1.21】 new panic print println real recover
```

​	标识符分布在：变量名、常量名、包名、函数名、方法名、类型名等的命名上，同时也分布在某些类型的进一步抽象上，如：`comparable`，以及一些常量值和零值上，如：`true`、`false`、`iota`、`nil`。
