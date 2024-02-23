+++
title = "go help packages"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# go help packages

version go 1.20.6



Many commands apply to a set of packages:

​	许多命令适用于一组包：

```
    go action [packages]
```

Usually, [packages] is a list of import paths.

​	通常，[packages] 是导入路径的列表。

An import path that is a rooted path or that begins with a . or .. element is interpreted as a file system path and denotes the package in that directory.

​	作为根路径或以 . 或 .. 元素开头的导入路径被解释为文件系统路径，并表示该目录中的包。

Otherwise, the import path P denotes the package found in the directory DIR/src/P for some DIR listed in the GOPATH environment variable (For more details see: ‘go help gopath’).

​	否则，导入路径 P 表示在 GOPATH 环境变量中列出的某个 DIR 的目录 DIR/src/P 中找到的包（有关更多详细信息，请参阅：“go help gopath”）。

If no import paths are given, the action applies to the package in the current directory.

​	如果没有给出导入路径，则该操作适用于当前目录中的包。

There are four reserved names for paths that should not be used for packages to be built with the go tool:

​	对于不应使用 go 工具构建的包，有四个保留名称的路径：

- “main” denotes the top-level package in a stand-alone executable.

- “main”表示独立可执行文件中的顶级包。

- “all” expands to all packages found in all the GOPATH trees. For example, ‘go list all’ lists all the packages on the local system. When using modules, “all” expands to all packages in the main module and their dependencies, including dependencies needed by tests of any of those.

- “all”扩展到在所有 GOPATH 树中找到的所有包。例如，“go list all”列出本地系统上的所有包。使用模块时，“all”扩展到主模块中的所有包及其依赖项，包括任何这些包的测试所需的依赖项。

- “std” is like all but expands to just the packages in the standard Go library.

- “std”类似于 all，但仅扩展到标准 Go 库中的包。

- “cmd” expands to the Go repository’s commands and their internal libraries.

- “cmd”扩展到 Go 存储库的命令及其内部库。

Import paths beginning with “`cmd/`” only match source code in the Go repository.

​	以“ `cmd/` ”开头的导入路径仅匹配 Go 存储库中的源代码。

An import path is a pattern if it includes one or more “`...`” wildcards, each of which can match any string, including the empty string and strings containing slashes. Such a pattern expands to all package directories found in the GOPATH trees with names matching the patterns.

​	如果导入路径包含一个或多个“ `...` ”通配符，则该路径是一个模式，每个通配符可以匹配任何字符串，包括空字符串和包含斜杠的字符串。此类模式会扩展到在 GOPATH 树中找到的所有名称与模式匹配的包目录。

To make common patterns more convenient, there are two special cases.

​	为了使常用模式更方便，这里有两个特殊情况。

First, `/...` at the end of the pattern can match an empty string, so that net/… matches both net and packages in its subdirectories, like net/http.

​	首先，模式末尾的 `/...` 可以匹配空字符串，因此 net/… 既匹配 net，也匹配其子目录中的包，例如 net/http。

Second, any slash-separated pattern element containing a wildcard never participates in a match of the “vendor” element in the path of a vendored package, so that `./...` does not match packages in subdirectories of `./vendor` or `./mycode/vendor`, but `./vendor/...` and `./mycode/vendor/...` do.

​	其次，任何包含通配符的以斜杠分隔的模式元素绝不会参与对已售包路径中的“vendor”元素的匹配，因此 `./...` 不匹配 `./vendor` 或 `./mycode/vendor` 的子目录中的包，但 `./vendor/...` 和 `./mycode/vendor/...` 匹配。

Note, however, that a directory named vendor that itself contains code is not a vendored package: cmd/vendor would be a command named vendor, and the pattern `cmd/...` matches it.

​	但是，请注意，名为 vendor 且本身包含代码的目录不是已售包：cmd/vendor 将是名为 vendor 的命令，模式 `cmd/...` 与之匹配。

See [golang.org/s/go15vendor](https://go.googlesource.com/proposal/+/master/design/25719-go15vendor.md) for more about vendoring.

​	有关售卖的更多信息，请参阅 golang.org/s/go15vendor。

An import path can also name a package to be downloaded from a remote repository. Run ‘go help importpath’ for details.

​	导入路径还可以命名要从远程存储库下载的包。有关详细信息，请运行“go help importpath”。

Every package in a program must have a unique import path.

​	程序中的每个包都必须具有唯一的导入路径。

By convention, this is arranged by starting each path with a unique prefix that belongs to you. For example, paths used internally at Google all begin with ‘google’, and paths denoting remote repositories begin with the path to the code, such as ‘github.com/user/repo’.

​	根据惯例，这是通过以属于您的唯一前缀开始每个路径来安排的。例如，在 Google 内部使用的路径都以“google”开头，表示远程存储库的路径以代码路径开头，例如“github.com/user/repo”。

Packages in a program need not have unique package names, but there are two reserved package names with special meaning.

​	程序中的包不必具有唯一的包名称，但有两个具有特殊含义的保留包名称。

The name main indicates a command, not a library. Commands are built into binaries and cannot be imported.

​	名称 main 表示命令，而不是库。命令内置于二进制文件中，无法导入。

The name documentation indicates documentation for a non-Go program in the directory. Files in package documentation are ignored by the go command.

​	名称 documentation 表示目录中非 Go 程序的文档。go 命令会忽略包 documentation 中的文件。

As a special case, if the package list is a list of .go files from a single directory, the command is applied to a single synthesized package made up of exactly those files, ignoring any build constraints in those files and ignoring any other files in the directory.

​	作为特例，如果包列表是来自单个目录的 .go 文件列表，则该命令将应用于由这些文件组成的单个综合包，忽略这些文件中的任何构建约束并忽略目录中的任何其他文件。

Directory and file names that begin with “.” or “_” are ignored by the go tool, as are directories named “testdata”.

​	以“.”或“_”开头的目录和文件名以及名为“testdata”的目录会被 go 工具忽略。
