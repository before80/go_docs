+++
title = "位置说明符"
date = 2024-12-09T07:59:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/cli/locspec.md](https://github.com/go-delve/delve/blob/master/Documentation/cli/locspec.md)
>
> 收录该文档时间： `2024-12-09T07:59:34+08:00`

# Location Specifiers - 位置说明符



Several delve commands take a program location as an argument, the syntax accepted by this commands is:

​	多个 delve 命令接受程序位置作为参数，以下是这些命令接受的语法：

- `*<address>` Specifies the location of memory address *address*. *address* can be specified as a decimal, hexadecimal or octal number
  - `*<address>` 指定内存地址 *address* 的位置。*address* 可以是十进制、十六进制或八进制数字。
- `<filename>:<line>` Specifies the line *line* in *filename*. *filename* can be the partial path to a file or even just the base name as long as the expression remains unambiguous.
  - `<filename>:<line>` 指定 *filename* 中的 *line* 行。*filename* 可以是文件的部分路径，甚至仅仅是文件的基本名称，只要表达式没有歧义。
- `<line>` Specifies the line *line* in the current file
  - `<line>` 指定当前文件中的 *line* 行。
- `+<offset>` Specifies the line *offset* lines after the current one
  - `+<offset>` 指定当前行后 *offset* 行。
- `-<offset>` Specifies the line *offset* lines before the current one
  - `-<offset>` 指定当前行前 *offset* 行。
- `<function>[:<line>]` Specifies the line *line* inside *function*. The full syntax for *function* is `<package>.(*<receiver type>).<function name>` however the only required element is the function name, everything else can be omitted as long as the expression remains unambiguous. For setting a breakpoint on an init function (ex: main.init), the `<filename>:<line>` syntax should be used to break in the correct init function at the correct location.
  - `<function>[:<line>]` 指定 *function* 内的 *line* 行。*function* 的完整语法是 `<package>.(*<receiver type>).<function name>`，但仅需提供函数名，其他部分可以省略，只要表达式没有歧义。要在 init 函数（例如：main.init）上设置断点，应使用 `<filename>:<line>` 语法来在正确的位置断开正确的 init 函数。
- `/<regex>/` Specifies the location of all the functions matching *regex*
  - `/<regex>/` 指定所有匹配 *regex* 的函数的位置。
