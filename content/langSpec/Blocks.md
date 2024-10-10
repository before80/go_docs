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

A *block* is a possibly empty sequence of declarations and statements within matching brace brackets. 

​	块是一对匹配的花括号内可能为空的声明和语句序列。

```
Block = "{" StatementList "}" .
StatementList = { Statement ";" } .
```

In addition to explicit blocks in the source code, there are implicit blocks:

​	源代码中除了显式块之外，还有隐式块：

1. The *universe block* encompasses all Go source text.
2. 包含所有的Go源码文本的 `universe block` 。
3. Each [package](https://go.dev/ref/spec#Packages) has a *package block* containing all Go source text for that package.
4. 每个[包](../Packages)都有一个`package block`，包含该包的所有 Go 源代码。
5. Each file has a *file block* containing all Go source text in that file.
6. 每个文件都有一个`file block`，包含该文件中的所有Go 源代码。
7. Each ["if"](https://go.dev/ref/spec#If_statements), ["for"](https://go.dev/ref/spec#For_statements), and ["switch"](https://go.dev/ref/spec#Switch_statements) statement is considered to be in its own implicit block.
8. 每个 "[if](../Statements#if-statements---if-语句)"、"[for](../Statements#for-statements----for-语句) "和 "[switch](../Statements#switch-statements----switch-语句) "语句都被认为是在自己的隐式块中。
9. Each clause in a ["switch"](https://go.dev/ref/spec#Switch_statements) or ["select"](https://go.dev/ref/spec#Select_statements) statement acts as an implicit block.
10. "[switch](../Statements#switch-statements----switch-语句)"或 "[select](../Statements#select-statements---select-语句) "语句中的每个子句都是一个隐式块。

Blocks nest and influence [scoping](https://go.dev/ref/spec#Declarations_and_scope).

​	块是可以嵌套并影响着[作用域](../DeclarationsAndScope)。