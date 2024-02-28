+++
title = "标记法"
date = 2023-05-17T09:59:21+08:00
weight = 2
description = ""
isCJKLanguage = true
type = "docs"
draft = false
+++
## Notation 标记法

> 原文：[https://go.dev/ref/spec#Notation](https://go.dev/ref/spec#Notation)

The syntax is specified using a [variant](https://en.wikipedia.org/wiki/Wirth_syntax_notation) of Extended Backus-Naur Form (EBNF):

​	（标记法的）语法是用 Extended Backus-Naur Form (EBNF) 的一个[变体](https://en.wikipedia.org/wiki/Wirth_syntax_notation)来指定的。

```
Syntax      = { Production } .
Production  = production_name "=" [ Expression ] "." .
Expression  = Term { "|" Term } .
Term        = Factor { Factor } .
Factor      = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Productions are expressions constructed from terms and the following operators, in increasing precedence:

​	产生式是由术语和以下运算符构建的表达式，其优先级越来越高。

```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

Lowercase production names are used to identify lexical (terminal) tokens. Non-terminals are in CamelCase. Lexical tokens are enclosed in double quotes `""` or back quotes \`\`.

​	小写的产生式名称用于标识词法（终端）标记。非终端则使用`CamelCase`。词法标记用双引号`""`或反引号\`\`包裹起来。

The form `a … b` represents the set of characters from `a` through `b` as alternatives. The horizontal ellipsis `…` is also used elsewhere in the spec to informally denote various enumerations or code snippets that are not further specified. The character `…` (as opposed to the three characters `...`) is not a token of the Go language.

​	`a ... b`的形式表示从`a`到`b`的一组字符作为备选项。在规范的其他地方也使用水平省略号`...`来非正式地表示各种枚举或没有进一步指定的代码片段。字符`…`（相对于三个字符`...`而言）不是Go语言的标记。

A link of the form [[Go 1.xx](https://go.dev/ref/spec#Language_versions)] indicates that a described language feature (or some aspect of it) was changed or added with language version 1.xx and thus requires at minimum that language version to build. For details, see the [linked section](https://go.dev/ref/spec#Language_versions) in the [appendix](https://go.dev/ref/spec#Appendix).

​	表格的链接[[Go 1.xx]({{< ref "/langSpec/Appendix#language-versions-语言版本">}})]表示已描述语言功能(或其某些方面)已更改或添加了语言版本1.xx和 因此，至少需要构建该语言版本。 有关详细信息，请参见 [链接部分]({{< ref "/langSpec/Appendix#language-versions-语言版本">}}) 在 [附录]({{< ref "/langSpec/Appendix">}})。 

