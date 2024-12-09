+++
title = "已知 Bug"
date = 2024-12-09T08:08:56+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/KnownBugs.md](https://github.com/go-delve/delve/blob/master/Documentation/KnownBugs.md)
>
> 收录该文档时间： `2024-12-09T08:08:56+08:00`

# Known Bugs - 已知 Bug



- When Delve is compiled with versions of go prior to 1.7.0 it is not possible to set a breakpoint on a function in a remote package using the `Receiver.MethodName` syntax. See [Issue #528](https://github.com/go-delve/delve/issues/528).
  - 当 Delve 使用 Go 1.7.0 之前的版本编译时，无法使用 `Receiver.MethodName` 语法在远程包中的函数上设置断点。请参见 [Issue #528](https://github.com/go-delve/delve/issues/528)。
- When running Delve on binaries compiled with a version of go prior to 1.9.0 `locals` will print all local variables, including ones that are out of scope, the shadowed flag will be applied arbitrarily. If there are multiple variables defined with the same name in the current function `print` will not be able to select the correct one for the current line.
  - 在使用 Go 1.9.0 之前版本编译的二进制文件上运行 Delve 时，`locals` 会打印所有局部变量，包括超出作用域的变量，`shadowed` 标志会被任意应用。如果当前函数中定义了多个相同名称的变量，`print` 将无法为当前行选择正确的变量。
- `reverse step` will not reverse step into functions called by deferred calls.
  - `reverse step` 无法逆向进入由延迟调用（deferred calls）调用的函数。
