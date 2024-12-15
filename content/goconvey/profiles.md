+++
title = "Profiles"
date = 2024-12-15T11:22:23+08:00
weight = 13
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Profiles](https://github.com/smartystreets/goconvey/wiki/Profiles)
>
> 收录该文档时间： `2024-12-15T11:22:23+08:00`

# Profiles



Manuel Mendez edited this page on Apr 21, 2016 · [5 revisions](https://github.com/smartystreets/goconvey/wiki/Profiles/_history)

​	Manuel Mendez 于 2016 年 4 月 21 日编辑 · [5 次修订](https://github.com/smartystreets/goconvey/wiki/Profiles/_history)

## GoConvey 测试包配置文件 GoConvey Test Package Profiles



When using the [web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) it would be nice to be able to customize the flags used when running `go test` in a specific package*. It would also be nice to somehow mark a package as persistently ignored or disabled so that when the UI starts sometime in the future, you don't have to click ignore on that package again. Well, there's a way to do both of those things by creating a text file in that package that satisfied the following regex:

​	使用 [Web UI](https://github.com/smartystreets/goconvey/wiki/Web-UI) 时，如果可以自定义运行特定包中的 `go test` 命令所用的标志，会非常方便。同样，如果可以将某个包标记为永久忽略或禁用，那么在未来启动 UI 时，就不必再手动点击忽略该包了。幸运的是，通过在包中创建一个满足以下正则表达式的文本文件，可以实现这两种功能：

```
.+\.goconvey
```



有效文件名示例：Examples of good names:

- `example.goconvey`
- `hi.goconvey`

无效文件名示例：Examples of bad names:

- `.goconvey`
- `hi.txt`

The contents of that text file may include blank lines, commented lines (start the line with a `#` or `//`), test flags (`-short`, etc...) or the word `ignore` (casing is unimportant). The only hard and fast rule is that if you include the word `ignore` it must come before any test flags.

​	该文本文件的内容可以包括空行、注释行（以 `#` 或 `//` 开头）、测试标志（如 `-short` 等）或单词 `ignore`（大小写不敏感）。唯一的硬性规则是，如果包含单词 `ignore`，它必须出现在所有测试标志之前。

The result of creating such a file is that if `ignore` is found as the first non-comment, non-blank line, the package will be ignored by the GoConvey server until that line is removed or commented. Otherwise, all non-blank, non-comment lines are treated as arguments to be used whenever calling `go test` on your package (like when you save a file). Most of these arguments will probably arguments defined by the golang testing command but you can include arbitrary arguments to your package. There are only a few caveats to be aware of:

​	创建这样的文件后，如果 `ignore` 是第一个非注释、非空行，GoConvey 服务器将在 UI 中忽略该包，直到移除或注释掉该行。否则，所有非注释、非空行将作为参数传递给调用 `go test` 的命令（例如保存文件时）。这些参数大部分可能是 Golang 测试命令中定义的参数，但您也可以为包指定任意自定义参数。需要注意以下几点：

1. `-v` is always used so there's no need to ever include this flag.`-v` 标志始终默认启用，因此无需显式添加。
2. `-cover`, `-covermode` and `-coverprofile` are specified by the GoConvey server so including these flags will have no effect. (Let me know if you really want to be able to specify `-covermode` for values other than `set` which is what GoConvey uses by default.) `-cover`、`-covermode` 和 `-coverprofile` 标志由 GoConvey 服务器指定，因此手动添加这些标志无效。（如果您确实需要指定 `-covermode` 的值为非默认的 `set`，请联系我们。）
3. `-tags` will be passed appropriately through to `go test`, `go list`, etc. calls, so if you need a special tag to build your package or your tests, it should get picked up correctly. `-tags` 标志会被正确传递到 `go test`、`go list` 等命令调用中，因此如果需要特定标签来构建包或测试，它将被正确处理。
4. Many of the profiling/benchmarking flags haven't really been used much by the core goconvey developers so we aren't sure if they will work as they normally do when applied to a goconvey package profile. 核心开发者对某些分析/基准标志的使用较少，因此不确定它们在应用于 GoConvey 配置文件时是否按预期工作。

Far from being an exhaustive list, here are some intended use cases:

​	虽然以上不是详尽的列表，但以下是一些常见的用途示例：

- Using the `-short` flag can allow you to toggle execution of long-running integration tests.
  - 使用 `-short` 标志可以切换长时间运行的集成测试的执行状态。

- Using the `-run` flag can allow you to focus on specific test functions in a package. Combined with `FocusConvey` you could limit test execution to a single test case within a single test function--handy if displaying lots of debug information.
  - 使用 `-run` 标志可以专注于包中的特定测试函数。结合 `FocusConvey`，您可以将测试限制在单个测试函数中的单个测试用例上，这在显示大量调试信息时非常有用。

- Using the `-timeout` flag may help prevent manual restarts of the goconvey server in the event of a mistaken infinite loop (come on, admit it--you know it's happened to you!).
  - 使用 `-timeout` 标志可以防止因意外的无限循环而需要手动重启 GoConvey 服务器。（承认吧，这种情况您肯定遇到过！）


Please see the `examples.goconvey` profile in the examples package ([`github.com/smartystreets/goconvey/examples`](https://github.com/smartystreets/goconvey/tree/master/examples)) for, well, an example. Happy testing!

​	请参考示例包中的 `examples.goconvey` 配置文件（[`github.com/smartystreets/goconvey/examples`](https://github.com/smartystreets/goconvey/tree/master/examples)）以了解具体示例。祝您测试愉快！

- A profile only applies to the containing package (same folder), not any nested packages (sub-folders).
  - 配置文件仅适用于所在的包（同一文件夹），不适用于嵌套包（子文件夹）。
