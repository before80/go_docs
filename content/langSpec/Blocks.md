+++
title = "块"
date = 2023-05-17T09:59:21+08:00
weight = 9
description = ""
isCJKLanguage = true
type = "docs"
draft = false
+++
## Blocks 块

> 原文：[https://go.dev/ref/spec#Blocks](https://go.dev/ref/spec#Blocks)

​	块是一对匹配的花括号内可能为空的声明和语句序列。

```
Block = "{" StatementList "}" .
StatementList = { Statement ";" } .
```

​	源代码中除了显式块之外，还有隐式块：

1. 包含所有的Go源码文本的 `universe block` 。
2. 每个[包](../Packages)都有一个`package block`，包含该包的所有 Go 源代码。
3. 每个文件都有一个`file block`，包含该文件中的所有Go 源代码。
4. 每个 "[if](../Statements#if-statements---if-语句)"、"[for](../Statements#for-statements----for-语句) "和 "[switch](../Statements#switch-statements----switch-语句) "语句都被认为是在自己的隐式块中。
5. "[switch](../Statements#switch-statements----switch-语句)"或 "[select](../Statements#select-statements---select-语句) "语句中的每个子句都是一个隐式块。

​	块是可以嵌套并影响着[作用域](../DeclarationsAndScope)。