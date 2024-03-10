+++
title = "go help module-get"
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 610
description = ""
isCJKLanguage = true
draft = false

+++

​	

The 'go get' command changes behavior depending on whether the go command is running in module-aware mode or legacy GOPATH mode. This help text, accessible as 'go help module-get' even in legacy GOPATH mode, describes 'go get' as it operates in module-aware mode.

​	'go get'命令的行为取决于go命令是在模块感知模式还是传统的GOPATH模式下运行。这篇帮助文档，即使在传统的GOPATH模式下也可以通过'go help module-get'访问，描述了在模块感知模式下'go get'的操作方式。

Usage:

​	用法:

```cmd
go get [-t] [-u] [-v] [build flags] [packages]
```

Get resolves its command-line arguments to packages at specific module versions, updates go.mod to require those versions, and downloads source code into the module cache.

​	Get将其命令行参数解析为特定模块版本的包，更新go.mod以要求这些版本，并将源代码下载到模块缓存中。

To add a dependency for a package or upgrade it to its latest version:

​	要为包添加依赖项或升级到其最新版本：

        go get example.com/pkg

To upgrade or downgrade a package to a specific version:

​	要升级或降级到特定版本的包：

        go get example.com/pkg@v1.2.3

To remove a dependency on a module and downgrade modules that require it:

​	要删除对模块的依赖关系并降级依赖它的模块：

        go get example.com/mod@none

To upgrade the minimum required Go version to the latest released Go version:

​	要升级最低要求的Go版本到最新发布的Go版本：

        go get go@latest

To upgrade the Go toolchain to the latest patch release of the current Go toolchain:

​	要升级Go工具链到当前Go工具链的最新补丁版本：

        go get toolchain@patch

See https://golang.org/ref/mod#go-get for details.

​	有关详细信息，请参见 https://golang.org/ref/mod#go-get。

In earlier versions of Go, 'go get' was used to build and install packages. Now, 'go get' is dedicated to adjusting dependencies in go.mod. 'go install' may be used to build and install commands instead. When a version is specified, 'go install' runs in module-aware mode and ignores the go.mod file in the current directory. For example:

​	在Go的早期版本中，'go get'用于构建和安装包。现在，'go get'专用于调整go.mod中的依赖关系。相反，可以使用'go install'来构建和安装命令。当指定版本时，'go install'在模块感知模式下运行，并忽略当前目录中的go.mod文件。例如：

        go install example.com/pkg@v1.2.3
        go install example.com/pkg@latest

See 'go help install' or https://golang.org/ref/mod#go-install for details.

​	有关详细信息，请参见'go help install'或 https://golang.org/ref/mod#go-install。

'go get' accepts the following flags.

​	'go get'接受以下标志。

The -t flag instructs get to consider modules needed to build tests of packages specified on the command line.

​	-t标志指示get考虑用于构建命令行上指定的包的测试的模块。

The -u flag instructs get to update modules providing dependencies of packages named on the command line to use newer minor or patch releases when available.

​	-u标志指示get将命令行上指定的包的依赖项提供的模块更新为使用较新的次要或补丁版本（如果可用）。

The -u=patch flag (not -u patch) also instructs get to update dependencies, but changes the default to select patch releases.

​	-u=patch标志（不是-u patch）也指示get更新依赖项，但将默认更改为选择补丁版本。

When the -t and -u flags are used together, get will update test dependencies as well.

​	当-t和-u标志一起使用时，get将同时更新测试依赖项。

The -x flag prints commands as they are executed. This is useful for debugging version control commands when a module is downloaded directly from a repository.

​	-x标志会打印执行的命令。这对于在直接从存储库下载模块时调试版本控制命令非常有用。

For more about modules, see https://golang.org/ref/mod.

​	有关模块的更多信息，请参见 https://golang.org/ref/mod。

For more about using 'go get' to update the minimum Go version and suggested Go toolchain, see https://go.dev/doc/toolchain.

​	有关使用'go get'更新最低Go版本和建议的Go工具链的详细信息，请参见 https://go.dev/doc/toolchain。

For more about specifying packages, see 'go help packages'.

​	有关指定包的详细信息，请参见'go help packages'。

This text describes the behavior of get using modules to manage source code and dependencies. If instead the go command is running in GOPATH mode, the details of get's flags and effects change, as does 'go help get'. See 'go help gopath-get'.

​	此文本描述了使用模块管理源代码和依赖项的get的行为。如果相反，go命令在GOPATH模式下运行，get的标志和效果的详细信息将发生变化，'go help get'也将发生变化。请参阅'go help gopath-get'。

See also: go build, go install, go clean, go mod.



​	另请参见：go build、go install、go clean、go mod。
