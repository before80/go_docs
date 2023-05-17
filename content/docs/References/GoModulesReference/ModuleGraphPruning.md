+++
title = "模块图的修剪"
date = 2023-05-17T09:59:21+08:00
weight = 5
description = ""
isCJKLanguage = true
draft = false
+++
## Module graph pruning 模块图的修剪

> 原文：[https://go.dev/ref/mod#graph-pruning](https://go.dev/ref/mod#graph-pruning)

​	如果主模块是`go 1.17`或更高版本，用于[minimal version selection（最小化版本选择）](../MVS)的[模块图](../Glossary#module-graph)仅包括每个模块依赖项的 immediate（直接）requirements ，这些依赖项在其自己的 `go.mod` 文件中指定到`go 1.17`或更高版本，除非该版本的模块也被`go 1.16`或更低版本的其他依赖（过渡性）所需要。(`go 1.17`版本的过渡性依赖项会从模块图中删去）。

​	由于 `go 1.17` 的 `go.mod` 文件包括构建任何包或测试所需的每个依赖项的 [require]() 指令，因此修剪后的模块图包括`go build`或`go test`[main module（主模块）](../Glossary#main-module)明确要求的任何依赖项中的包所需的所有依赖项。不需要构建任何包或测试的模块不能影响其包的运行时行为，因此从模块图中修剪出来的依赖项只会导致其他不相关的模块之间的干扰。

​	那些 requirements 被修剪掉的模块仍然出现在模块图中，并且仍然被`go list -m all`报告：它们[selected versions（所选择的版本）](../Glossary#selected-version)是已知的并且定义良好，并且包可以从这些模块中加载（例如，作为从其他模块加载的测试的过渡性依赖项）。然而，由于`go`命令不能轻易识别这些模块的哪些依赖项得到了满足，所以`go build`和`go test`的参数不能包括那些 requirements 已被修剪掉的模块的包。 [go get](../Module-awareCommands#go-get)将包含每个命名包的模块提升为显式的依赖项，允许在该包上调用`go build`或`go test`。

​	**因为Go 1.16和更早的版本不支持模块图的修剪**，所以对于每个指定Go 1.16或更低版本的模块来说，完整的依赖项的过渡性闭包 —— 包括过渡性的`go 1.17`依赖项 —— 仍然被包括在内。(在 `go 1.16` 及以下版本中，`go.mod` 文件只包括[direct dependencies（直接依赖项）](../Glossary#direct-dependency)，因此必须加载更大的图以确保包括所有间接依赖项）。

​	默认情况下，[go mod tidy](../Module-awareCommands#go-mod-tidy)为模块记录的 [go.sum 文件](../AuthenticatingModules#go.sum-files) 包括Go 版本在 [go 指令](../Module-awareCommands#go-directive)中指定的版本的前一个版本所需的校验和。因此，`go 1.17` 的模块包括 Go 1.16 所加载的完整模块图所需的校验和，但 `go 1.18` 的模块将只包括 Go 1.17 所加载的修剪模块图所需的校验和。可以使用`-compat`标志来覆盖默认版本（例如，在`go 1.17`模块中更积极地修剪`go.sum`文件）。

​	更多细节参见[设计文档](https://go.googlesource.com/proposal/+/master/design/36460-lazy-module-loading.md)。

### Lazy module loading 延迟模块加载

​	为模块图修剪增加的更全面的 requirements 也使得在模块内工作时可以进行另一种优化。如果主模块是在`go 1.17`或更高版本，`go`命令会避免加载完整的模块图，直到（或除非）需要它。相反，它只加载主模块的`go.mod`文件，然后尝试只使用这些 requirements 来加载要构建的包。如果在这些 requirements中没有找到要导入的包（例如，主模块之外的包的测试依赖项），那么模块图的其余部分将被加载。

​	如果在不加载模块图的情况下能找到所有导入的包，`go`命令将只加载包含这些包的模块的`go.mod`文件，并将它们的requirements 与主模块的requirements 进行核对，以确保它们是本地一致的。(不一致可能是由于版本控制的合并，手工编辑，以及使用本地文件系统路径[替换](../Module-awareCommands#replace-directive)的模块的更改造成的）。