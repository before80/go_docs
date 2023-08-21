+++
title = "Go 工具链"
date = 2023-08-21T22:53:42+08:00
weight = 1000
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Go Toolchains - Go 工具链

https://go.dev/doc/toolchain

## 简介

Starting in Go 1.21, the Go distribution consists of a `go` command and a bundled Go toolchain, which is the standard library as well as the compiler, assembler, and other tools. The `go` command can use its bundled Go toolchain as well as other versions that it finds in the local `PATH` or downloads as needed.

​	从Go 1.21开始，Go发行版包含一个`go`命令和一个捆绑的Go工具链，其中包括标准库以及编译器、汇编器和其他工具。`go`命令可以使用其捆绑的Go工具链，以及它在本地`PATH`中找到的其他版本或根据需要下载的版本。

The choice of Go toolchain being used depends on the `GOTOOLCHAIN` environment setting and the `go` and `toolchain` lines in the main module’s `go.mod` file or the current workspace’s `go.work` file. As you move between different main modules and workspaces, the toolchain version being used can vary, just as module dependency versions do.

​	所使用的Go工具链的选择取决于`GOTOOLCHAIN`环境设置以及主模块的`go.mod`文件或当前工作区的`go.work`文件中的`go`和`toolchain`行。当您在不同的主模块和工作区之间移动时，所使用的工具链版本可能会有所不同，就像模块依赖版本一样。

In the standard configuration, the `go` command uses its own bundled toolchain when that toolchain is at least as new as the `go` or `toolchain` lines in the main module or workspace. For example, when using the `go` command bundled with Go 1.21.3 in a main module that says `go 1.21.0`, the `go` command uses Go 1.21.3. When the `go` or `toolchain` line is newer than the bundled toolchain, the `go` command runs the newer toolchain instead. For example, when using the `go` command bundled with Go 1.21.3 in a main module that says `go 1.21.9`, the `go` command finds and runs Go 1.21.9 instead. It first looks in the PATH for a program named `go1.21.9` and otherwise downloads and caches a copy of the Go 1.21.9 toolchain. This automatic toolchain switching can be disabled, but in that case, for more precise forwards compatibility, the `go` command will refuse to run in a main module or workspace in which the `go` line requires a newer version of Go. That is, the `go` line sets the minimum required Go version necessary to use a module or workspace.

​	在标准配置中，当捆绑的工具链至少与主模块或工作区中的`go`或`toolchain`行一样新时，`go`命令将使用其自己的捆绑工具链。例如，使用捆绑在Go 1.21.3中的`go`命令时，如果主模块中指定了`go 1.21.0`，则`go`命令将使用Go 1.21.3。当`go`或`toolchain`行比捆绑的工具链更新时，`go`命令将运行较新的工具链。例如，使用捆绑在Go 1.21.3中的`go`命令时，如果主模块中指定了`go 1.21.9`，`go`命令将找到并运行Go 1.21.9。它首先在PATH中查找名为`go1.21.9`的程序，否则会下载并缓存Go 1.21.9的工具链副本。这种自动工具链切换可以禁用，但在这种情况下，为了更精确的向前兼容性，`go`命令将拒绝在主模块或工作区中运行，其中`go`行需要更新版本的Go。也就是说，`go`行设置了使用模块或工作区所需的最低Go版本。

Modules that are dependencies of other modules may need to set a minimum Go version requirement lower than the preferred toolchain to use when working in that module directly. In this case, the `toolchain` line in `go.mod` or `go.work` sets a preferred toolchain that takes precedence over the `go` line when the `go` command is deciding which toolchain to use.

​	作为其他模块的依赖项的模块可能需要将最低Go版本要求设置为低于在直接工作在该模块时首选工具链的版本。在这种情况下，`go.mod`或`go.work`中的`toolchain`行设置了一个首选的工具链，该工具链在`go`命令决定使用哪个工具链时优先于`go`行。

The `go` and `toolchain` lines can be thought of as specifying the version requirements for the module’s dependency on the Go toolchain itself, just as the `require` lines in `go.mod` specify the version requirements for dependencies on other modules. The `go get` command manages the Go toolchain dependency just as it manages dependencies on other modules. For example, `go get go@latest` updates the module to require the latest released Go toolchain.

​	`go`和`toolchain`行可以被视为指定模块对Go工具链自身的依赖的版本要求，就像`go.mod`中的`require`行指定了对其他模块的依赖的版本要求一样。`go get`命令管理Go工具链依赖，就像它管理其他模块的依赖一样。例如，`go get go@latest`将更新模块以依赖于最新发布的Go工具链。

The `GOTOOLCHAIN` environment setting can force a specific Go version, overriding the `go` and `toolchain` lines. For example, to test a package with Go 1.21rc3:

​	`GOTOOLCHAIN`环境设置可以强制使用特定的Go版本，覆盖`go`和`toolchain`行。例如，要使用Go 1.21rc3测试一个包：

```
GOTOOLCHAIN=go1.21rc3 go test
```

The default `GOTOOLCHAIN` setting is `auto`, which enables the toolchain switching described earlier. The alternate form `<name>+auto` sets the default toolchain to use before deciding whether to switch further. For example `GOTOOLCHAIN=go1.21.3+auto` directs the `go` command to begin its decision with a default of using Go 1.21.3 but still use a newer toolchain if directed by `go` and `toolchain` lines. Because the default `GOTOOLCHAIN` setting can be changed with `go env -w`, if you have Go 1.21.0 or later installed, then

​	默认的`GOTOOLCHAIN`设置是`auto`，它启用了上述描述的工具链切换。替代形式`<name>+auto`在决定是否进一步切换之前设置了默认的工具链。例如，`GOTOOLCHAIN=go1.21.3+auto`将指导`go`命令从默认情况下开始使用Go 1.21.3，但如果由`go`和`toolchain`行指示，则仍然使用较新的工具链。由于默认的`GOTOOLCHAIN`设置可以通过`go env -w`进行更改，如果您安装了Go 1.21.0或更高版本，那么

```
go env -w GOTOOLCHAIN=go1.21.3+auto
```

is equivalent to replacing your Go 1.21.0 installation with Go 1.21.3.

等效于将您的Go 1.21.0安装替换为Go 1.21.3。

The rest of this document explains how Go toolchains are versioned, chosen, and managed in more detail.

​	本文档的其余部分将更详细地解释Go工具链的版本管理、选择和管理。

## Go版本

Released versions of Go use the version syntax ‘1.*N*.*P*’, denoting the *P*th release of Go 1.*N*. The initial release is 1.*N*.0, like in ‘1.21.0’. Later releases like 1.*N*.9 are often referred to as patch releases.

​	发布版本的Go使用版本语法‘1.*N*.*P*’，表示Go 1.*N*的第 *P* 次发布。初始版本为1.*N*.0，例如‘1.21.0’。后续版本如1.*N*.9通常称为补丁版本。

Go 1.*N* release candidates, which are issued before 1.*N*.0, use the version syntax ‘1.*N*rc*R*’. The first release candidate for Go 1.*N* has version 1.*N*rc1, like in `1.23rc1`.

​	Go 1.*N* 发布候选版本在1.*N*.0之前发布，使用版本语法‘1.*N*rc*R*’。Go 1.*N*的第一个发布候选版本的版本为1.*N*rc1，例如`1.23rc1`。

The syntax ‘1.*N*’ is called a “language version”. It denotes the overall family of Go releases implementing that version of the Go language and standard library.

​	语法‘1.*N*’ 被称为“语言版本”。它表示实现该Go版本的Go语言和标准库的整体系列。

The language version for a Go version is the result of truncating everything after the *N*: 1.21, 1.21rc2, and 1.21.3 all implement language version 1.21.

​	Go版本的语言版本是将*N*后面的所有内容截断的结果：1.21、1.21rc2和1.21.3都实现了语言版本1.21。

Released Go toolchains such as Go 1.21.0 and Go 1.21rc1 report that specific version (for example, `go1.21.0` or `go1.21rc1`) from `go version` and [`runtime.Version`](https://go.dev/pkg/runtime/#Version). Unreleased (still in development) Go toolchains built from the Go development repository instead report only the language version (for example, `go1.21`).

​	已发布的Go工具链，如Go 1.21.0和Go 1.21rc1，在`go version`和[`runtime.Version`](https://go.dev/pkg/runtime/#Version)中报告特定的版本（例如`go1.21.0`或`go1.21rc1`）。未发布（仍在开发中）的Go工具链从Go开发存储库构建时，只报告语言版本（例如`go1.21`）。

Any two Go versions can be compared to decide whether one is less than, greater than, or equal to the other. If the language versions are different, that decides the comparison: 1.21.9 < 1.22. Within a language version, the ordering from least to greatest is: the language version itself, then release candidates ordered by *R*, then releases ordered by *P*.

​	任何两个Go版本都可以进行比较，以决定一个是否小于、大于或等于另一个。如果语言版本不同，则决定比较：1.21.9 < 1.22。在语言版本内，从最小到最大的顺序是：语言版本本身，然后按*R*顺序排列的发布候选版本，然后按*P*顺序排列的发布版本。

For example, 1.21 < 1.21rc1 < 1.21rc2 < 1.21.0 < 1.21.1 < 1.21.2.

​	例如，1.21 < 1.21rc1 < 1.21rc2 < 1.21.0 < 1.21.1 < 1.21.2。

Before Go 1.21, the initial release of a Go toolchain was version 1.*N*, not 1.*N*.0, so for *N* < 21, the ordering is adjusted to place 1.*N* after the release candidates.

​	在Go 1.21之前，Go工具链的初始版本是版本1.*N*，而不是1.*N*.0，因此对于*N* < 21，顺序进行了调整，将1.*N*放在发布候选版本之后。

For example, 1.20rc1 < 1.20rc2 < 1.20rc3 < 1.20 < 1.20.1.

​	例如，1.20rc1 < 1.20rc2 < 1.20rc3 < 1.20 < 1.20.1。

Earlier versions of Go had beta releases, with versions like 1.18beta2. Beta releases are placed immediately before release candidates in the version ordering.

​	早期的Go版本有beta版本，版本号如1.18beta2。Beta版本排在版本顺序中的发布候选版本之前。

For example, 1.18beta1 < 1.18beta2 < 1.18rc1 < 1.18 < 1.18.1.

​	例如，1.18beta1 < 1.18beta2 < 1.18rc1 < 1.18 < 1.18.1。

## Go工具链名称 Go toolchain names

The standard Go toolchains are named `go*V*` where *V* is a Go version denoting a beta release, release candidate, or release. For example, `go1.21rc1` and `go1.21.0` are toolchain names; `go1.21` and `go1.22` are not (the initial releases are `go1.21.0` and `go1.22.0`), but `go1.20` and `go1.19` are.

​	标准Go工具链的名称为`go*V*`，其中*V*是表示beta版本、发布候选版本或发布的Go版本。例如，`go1.21rc1`和`go1.21.0`是工具链名称；`go1.21`和`go1.22`不是（初始版本是`go1.21.0`和`go1.22.0`），但`go1.20`和`go1.19`是。

Non-standard toolchains use names of the form `go*V*-*suffix*` for any suffix.

​	非标准工具链使用`go*V*-*suffix*`的形式命名，其中*suffix*可以是任何后缀。

Toolchains are compared by comparing the version `*V*` embedded in the name (dropping the initial `go` and discarding off any suffix beginning with `-`). For example, `go1.21.0` and `go1.21.0-custom` compare equal for ordering purposes.

​	工具链通过比较嵌入在名称中的版本`*V*`（去掉开头的`go`并且丢弃以`-`开始的任何后缀）进行比较。例如，对于排序目的，`go1.21.0`和`go1.21.0-custom`进行比较时相等。

## 模块和工作区配置 Module and workspace configuration

Go modules and workspaces specify version-related configuration in their `go.mod` or `go.work` files.

​	Go模块和工作区在其`go.mod`或`go.work`文件中指定与版本相关的配置。

The `go` line declares the minimum required Go version for using the module or workspace. For compatibility reasons, if the `go` line is omitted from a `go.mod` file, the module is considered to have an implicit `go 1.16` line, and if the `go` line is omitted from a `go.work` file, the workspace is considered to have an implicit `go 1.18` line.

​	`go`行声明使用模块或工作区所需的最低Go版本。出于兼容性原因，如果从`go.mod`文件中省略了`go`行，则认为该模块具有隐式的`go 1.16`行，如果从`go.work`文件中省略了`go`行，则认为该工作区具有隐式的`go 1.18`行。

The `toolchain` line declares a suggested toolchain to use with the module or workspace. As described in “[Go toolchain selection](https://go.dev/doc/toolchain#select)” below, the `go` command may run this specific toolchain when operating in that module or workspace if the default toolchain’s version is less than the suggested toolchain’s version. If the `toolchain` line is omitted, the module or workspace is considered to have an implicit `toolchain go*V*` line, where *V* is the Go version from the `go` line.

​	`toolchain`行声明要与模块或工作区一起使用的建议工具链。如下所述，在“[Go工具链选择](https://go.dev/doc/toolchain#select)”中，当`go`命令在该模块或工作区中操作时，可能会运行此特定工具链，如果默认工具链的版本小于建议的工具链的版本。如果省略了`toolchain`行，则认为模块或工作区具有隐式的`toolchain go*V*`行，其中*V*是`go`行中的Go版本。

For example, a `go.mod` that says `go 1.21.0` with no `toolchain` line is interpreted as if it had a `toolchain go1.21.0` line.

​	例如，`go.mod`中指定了`go 1.21.0`且没有`toolchain`行，将解释为具有`toolchain go1.21.0`行。

The Go toolchain refuses to load a module or workspace that declares a minimum required Go version greater than the toolchain’s own version.

​	Go工具链拒绝加载声明的最低Go版本大于工具链自身版本的模块或工作区。

For example, Go 1.21.2 will refuse to load a module or workspace with a `go 1.21.3` or `go 1.22` line.

​	例如，Go 1.21.2将拒绝加载具有`go 1.21.3`或`go 1.22`行的模块或工作区。

A module’s `go` line must declare a version greater than or equal to the `go` version declared by each of the modules listed in `require` statements. A workspace’s `go` line must declare a version greater than or equal to the `go` version declared by each of the modules listed in `use` statements.

​	一个模块的`go`行必须声明的版本大于或等于`require`语句中列出的每个模块声明的`go`版本。一个工作区的`go`行必须声明的版本大于或等于`use`语句中列出的每个模块声明的`go`版本。

For example, if module *M* requires a dependency *D* with a `go.mod` that declares `go 1.22.0`, then *M*’s `go.mod` cannot say `go 1.21.3`.

​	例如，如果模块*M*需要一个依赖项*D*，而*D*的`go.mod`声明了`go 1.22.0`，那么模块*M*的`go.mod`不能声明为`go 1.21.3`。

The `go` line for each module sets the language version the compiler enforces when compiling packages in that module. The language version can be changed on a per-file basis by using a [build constraint](https://go.dev/cmd/go#hdr-Build_constraints).

​	每个模块的`go`行设置了编译器在编译该模块中的包时执行的语言版本。可以通过使用[构建约束](https://go.dev/cmd/go#hdr-Build_constraints)在每个文件的基础上更改语言版本。

For example, a module containing code that uses the Go 1.21 language version should have a `go.mod` file with a `go` line such as `go 1.21` or `go 1.21.3`. If a specific source file should be compiled only when using a newer Go toolchain, adding `//go:build go1.22` to that source file both ensures that only Go 1.22 and newer toolchains will compile the file and also changes the language version in that file to Go 1.22.

​	例如，包含使用Go 1.21语言版本的代码的模块应该有一个`go.mod`文件，其中`go`行类似于`go 1.21`或`go 1.21.3`。如果特定源文件只应在使用较新的Go工具链时进行编译，可以在该源文件中添加`//go:build go1.22`，这既确保只有Go 1.22和更高版本的工具链会编译该文件，还会将该文件中的语言版本更改为Go 1.22。

The `go` and `toolchain` lines are most conveniently and safely modified by using `go get`; see the [section dedicated to `go get` below](https://go.dev/doc/toolchain#get).

​	`go`和`toolchain`行最方便且最安全地通过使用`go get`进行修改；请参见[下面专门介绍的`go get`部分](https://go.dev/doc/toolchain#get)。

Before Go 1.21, Go toolchains treated the `go` line as an advisory requirement: if builds succeeded the toolchain assumed everything worked, and if not it printed a note about the potential version mismatch. Go 1.21 changed the `go` line to be a mandatory requirement instead. This behavior is partly backported to earlier language versions: Go 1.19 releases starting at Go 1.19.13 and Go 1.20 releases starting at Go 1.20.8, refuse to load workspaces or modules declaring version Go 1.22 or later.

​	在Go 1.21之前，Go工具链将`go`行视为一个建议的要求：如果构建成功，工具链会认为一切正常，如果不成功，则会打印有关潜在版本不匹配的说明。Go 1.21将`go`行更改为强制性要求。这种行为在一定程度上回溯到更早的语言版本：Go 1.19版本从Go 1.19.13开始，Go 1.20版本从Go 1.20.8开始，拒绝加载声明版本为Go 1.22或更高版本的工作区或模块。

Before Go 1.21, toolchains did not require a module or workspace to have a `go` line greater than or equal to the `go` version required by each of its dependency modules.

​	在Go 1.21之前，工具链不要求模块或工作区的`go`行大于或等于其每个依赖模块所需的`go`版本。

## The `GOTOOLCHAIN` setting `GOTOOLCHAIN`设置

The `go` command selects the Go toolchain to use based on the `GOTOOLCHAIN` setting. To find the `GOTOOLCHAIN` setting, the `go` command uses the standard rules for any Go environment setting:

​	`go`命令基于`GOTOOLCHAIN`设置选择要使用的Go工具链。要查找`GOTOOLCHAIN`设置，`go`命令使用了任何Go环境设置的标准规则：

 

- If `GOTOOLCHAIN` is set to a non-empty value in the process environment (as queried by [`os.Getenv`](https://go.dev/pkg/os/#Getenv)), the `go` command uses that value.
- 如果进程环境中的`GOTOOLCHAIN`设置为非空值（由[`os.Getenv`](https://go.dev/pkg/os/#Getenv)查询），则`go`命令使用该值。
- Otherwise, if `GOTOOLCHAIN` is set in the user’s environment default file (managed with [`go env -w` and `go env -u`](https://go.dev/cmd/go/#hdr-Print_Go_environment_information)), the `go` command uses that value.
- 否则，如果在用户的环境默认文件中设置了`GOTOOLCHAIN`（通过[`go env -w`和`go env -u`](https://go.dev/cmd/go/#hdr-Print_Go_environment_information)进行管理），则`go`命令使用该值。
- Otherwise, if `GOTOOLCHAIN` is set in the bundled Go toolchain’s environment default file (`$GOROOT/go.env`), the `go` command uses that value.
- 否则，如果在捆绑的Go工具链的环境默认文件（`$GOROOT/go.env`）中设置了`GOTOOLCHAIN`，则`go`命令使用该值。

In standard Go toolchains, the `$GOROOT/go.env` file sets the default `GOTOOLCHAIN=auto`, but repackaged Go toolchains may change this value.

​	在标准的Go工具链中，`$GOROOT/go.env`文件设置了默认的`GOTOOLCHAIN=auto`，但重新打包的Go工具链可能会更改此值。

If the `$GOROOT/go.env` file is missing or does not set a default, the `go` command assumes `GOTOOLCHAIN=local`.

​	如果`$GOROOT/go.env`文件丢失或未设置默认值，则`go`命令会假设`GOTOOLCHAIN=local`。

Running `go env GOTOOLCHAIN` prints the `GOTOOLCHAIN` setting.

​	运行`go env GOTOOLCHAIN`会打印出`GOTOOLCHAIN`设置。

## Go工具链选择

At startup, the `go` command selects which Go toolchain to use. It consults the `GOTOOLCHAIN` setting, which takes the form `<name>`, `<name>+auto`, or `<name>+path`. `GOTOOLCHAIN=auto` is shorthand for `GOTOOLCHAIN=local+auto`; similarly, `GOTOOLCHAIN=path` is shorthand for `GOTOOLCHAIN=local+path`. The `<name>` sets the default Go toolchain: `local` indicates the bundled Go toolchain (the one that shipped with the `go` command being run), and otherwise `<name>` must be a specific Go toolchain name, such as `go1.21.0`. The `go` command prefers to run the default Go toolchain. As noted above, starting in Go 1.21, Go toolchains refuse to run in workspaces or modules that require newer Go versions. Instead, they report an error and exit.

​	在启动时，`go`命令会选择要使用的Go工具链。它会查询`GOTOOLCHAIN`设置，该设置采用`<name>`、`<name>+auto`或`<name>+path`的形式。`GOTOOLCHAIN=auto`是`GOTOOLCHAIN=local+auto`的缩写；同样，`GOTOOLCHAIN=path`是`GOTOOLCHAIN=local+path`的缩写。`<name>`设置默认的Go工具链：`local`表示捆绑的Go工具链（与运行的`go`命令一起提供的工具链），否则`<name>`必须是特定的Go工具链名称，例如`go1.21.0`。`go`命令倾向于运行默认的Go工具链。如上所述，从Go 1.21开始，Go工具链拒绝在要求更新的Go版本的工作区或模块中运行。相反，它们报告错误并退出。

When `GOTOOLCHAIN` is set to `local`, the `go` command always runs the bundled Go toolchain.

​	当`GOTOOLCHAIN`设置为`local`时，`go`命令始终会运行捆绑的Go工具链。

When `GOTOOLCHAIN` is set to `<name>` (for example, `GOTOOLCHAIN=go1.21.0`), the `go` command always runs that specific Go toolchain. If a binary with that name is found in the system PATH, the `go` command uses it. Otherwise the `go` command uses a Go toolchain it downloads and verifies.

​	当`GOTOOLCHAIN`设置为`<name>`（例如，`GOTOOLCHAIN=go1.21.0`）时，`go`命令始终会运行该特定的Go工具链。如果在系统PATH中找到了具有该名称的二进制文件（例如`go1.21.3`），`go`命令会使用它。否则，`go`命令会下载并验证Go工具链。

When `GOTOOLCHAIN` is set to `<name>+auto` or `<name>+path` (or the shorthands `auto` or `path`), the `go` command selects and runs a newer Go version as needed. Specifically, it consults the `toolchain` and `go` lines in the current workspace’s `go.work` file or, when there is no workspace, the main module’s `go.mod` file. If the `go.work` or `go.mod` file has a `toolchain <tname>` line and `<tname>` is newer than the default Go toolchain, then the `go` command runs `<tname>` instead. If the file has a `toolchain default` line, then the `go` command runs the default Go toolchain, disabling any attempt at updating beyond `<name>`. Otherwise, if the file has a `go <version>` line and `<version>` is newer than the default Go toolchain, then the `go` command runs `go<version>` instead.

​	当`GOTOOLCHAIN`设置为`<name>+auto`或`<name>+path`（或缩写的`auto`或`path`），`go`命令会根据需要选择并运行更新的Go版本。具体来说，它会查询当前工作区`go.work`文件中的`toolchain`和`go`行，或者当没有工作区时，查询主模块的`go.mod`文件中的`toolchain`和`go`行。如果`go.work`或`go.mod`文件有一个`toolchain <tname>`行，并且`<tname>`比默认的Go工具链更新，则`go`命令会运行`<tname>`。如果文件有一个`toolchain default`行，则`go`命令会运行默认的Go工具链，禁用对`<name>`以外的更新尝试。否则，如果文件有一个`go <version>`行，并且`<version>`比默认的Go工具链更新，则`go`命令会运行`go<version>`。

To run a toolchain other than the bundled Go toolchain, the `go` command searches the process’s executable path (`$PATH` on Unix and Plan 9, `%PATH%` on Windows) for a program with the given name (for example, `go1.21.3`) and runs that program. If no such program is found, the `go` command [downloads and runs the specified Go toolchain](https://go.dev/doc/toolchain#download). Using the `GOTOOLCHAIN` form `<name>+path` disables the download fallback, causing the `go` command to stop after searching the executable path.

​	要运行除捆绑的Go工具链以外的工具链，`go`命令会在进程的可执行路径（Unix和Plan 9上的`$PATH`，Windows上的`%PATH%`）中搜索具有给定名称的程序（例如`go1.21.3`），然后运行该程序。如果没有找到这样的程序，则`go`命令会[下载并运行指定的Go工具链](https://go.dev/doc/toolchain#download)。使用`GOTOOLCHAIN`形式`<name>+path`会禁用下载回退，导致`go`命令在搜索可执行路径后停止。

Running `go version` prints the selected Go toolchain’s version (by running the selected toolchain’s implementation of `go version`).

​	运行`go version`会打印出所选Go工具链的版本（通过运行所选工具链的`go version`实现）。

Running `GOTOOLCHAIN=local go version` prints the bundled Go toolchain’s version.

​	运行`GOTOOLCHAIN=local go version`会打印出捆绑的Go工具链的版本。

## Go工具链切换 Go toolchain switches

For most commands, the workspace’s `go.work` or the main module’s `go.mod` will have a `go` line that is at least as new as the `go` line in any module dependency, due to the version ordering [configuration requirements](https://go.dev/doc/toolchain#config). In this case, the startup toolchain selection runs a new enough Go toolchain to complete the command.

​	对于大多数命令，工作区的`go.work`或主模块的`go.mod`文件的`go`行至少与任何模块依赖项的`go`行一样新，这是由于版本排序的[配置要求](https://go.dev/doc/toolchain#config)。在这种情况下，启动时的工具链选择会运行足够新的Go工具链以完成命令。

Some commands incorporate new module versions as part of their operation: `go get` adds new module dependencies to the main module; `go work use` adds new local modules to the workspace; `go work sync` resynchronizes a workspace with local modules that may have been updated since the workspace was created; `go install package@version` and `go run package@version` effectively run in an empty main module and add `package@version` as a new dependency. All these commands may encounter a module with a `go.mod` `go` line requiring a newer Go version than the currently executed Go version.

​	某些命令在其操作中包括新的模块版本：`go get`将新的模块依赖项添加到主模块；`go work use`将新的本地模块添加到工作区；`go work sync`重新同步工作区和可能已自创建工作区以来已更新的本地模块；`go install package@version`和`go run package@version`在空的主模块中实际运行，并将`package@version`添加为新的依赖项。所有这些命令可能会遇到模块，其`go.mod`中的`go`行要求比当前执行的Go版本更新的版本。

When a command encounters a module requiring a newer Go version and `GOTOOLCHAIN` permits running different toolchains (it is one of the `auto` or `path` forms), the `go` command chooses and switches to an appropriate newer toolchain to continue executing the current command.

​	当命令遇到需要更新的Go版本的模块，并且`GOTOOLCHAIN`允许运行不同的工具链（它是`auto`或`path`形式之一），`go`命令会选择并切换到适当的更新工具链以继续执行当前命令。

Any time the `go` command switches toolchains after startup toolchain selection, it prints a message explaining why. For example:

​	每次`go`命令在启动工具链选择后切换工具链时，它都会打印一条解释原因的消息。例如：

```
go: module example.com/widget@v1.2.3 requires go >= 1.24rc1; switching to go 1.27.9
```

As shown in the example, the `go` command may switch to a toolchain newer than the discovered requirement. In general the `go` command aims to switch to a supported Go toolchain.

​	如示例所示，`go`命令可能会切换到比发现的要求更新的工具链。一般来说，`go`命令旨在切换到受支持的Go工具链。

To choose the toolchain, the `go` command first obtains a list of available toolchains. For the `auto` form, the `go` command downloads a list of available toolchains. For the `path` form, the `go` command scans the PATH for any executables named for valid toolchains and uses a list of all the toolchains it finds. Using that list of toolchains, the `go` command identifies up to three candidates:

​	为了选择工具链，`go`命令首先获取可用工具链的列表。对于`auto`形式，`go`命令会下载可用工具链的列表。对于`path`形式，`go`命令会在PATH中扫描任何以有效工具链命名的可执行文件，并使用它找到的所有工具链的列表。使用该工具链列表，`go`命令识别最多三个候选者：

 

- the latest release candidate of an unreleased Go language version (1.*N*₃rc*R*₃),
- the latest patch release of the most recently released Go language version (1.*N*₂.*P*₂), and
- the latest patch release of the previous Go language version (1.*N*₁.*P*₁).
- 未发布的Go语言版本的最新发布候选版（1.*N*₃rc*R*₃），
- 最近发布的Go语言版本的最新补丁发布版（1.*N*₂.*P*₂），以及
- 以前的Go语言版本的最新补丁发布版（1.*N*₁.*P*₁）。

These are the supported Go releases according to Go’s [release policy](https://go.dev/doc/devel/release#policy). Consistent with [minimal version selection](https://research.swtch.com/vgo-mvs), the `go` command then conservatively uses the candidate with the *minimum* (oldest) version that satisfies the new requirement.

​	这些是根据Go的[发布政策](https://go.dev/doc/devel/release#policy)支持的Go版本。符合[最小版本选择](https://research.swtch.com/vgo-mvs)的一致性，`go`命令然后保守地使用满足新要求的*最小*（最旧）版本的候选者。

For example, suppose `example.com/widget@v1.2.3` requires Go 1.24rc1 or later. The `go` command obtains the list of available toolchains and finds that the latest patch releases of the two most recent Go toolchains are Go 1.28.3 and Go 1.27.9, and the release candidate Go 1.29rc2 is also available. In this situation, the `go` command will choose Go 1.27.9. If `widget` had required Go 1.28 or later, the `go` command would choose Go 1.28.3, because Go 1.27.9 is too old. If `widget` had required Go 1.29 or later, the `go` command would choose Go 1.29rc2, because both Go 1.27.9 and Go 1.28.3 are too old.

​	例如，假设`example.com/widget@v1.2.3`需要Go 1.24rc1或更高版本。`go`命令获取可用工具链的列表，并发现最近的两个Go工具链的最新补丁版本是Go 1.28.3和Go 1.27.9，还有发布候选版Go 1.29rc2可用。在这种情况下，`go`命令将选择Go 1.27.9。如果`widget`需要Go 1.28或更高版本，`go`命令将选择Go 1.28.3，因为Go 1.27.9太旧。如果`widget`需要Go 1.29或更高版本，`go`命令将选择Go 1.29rc2，因为Go 1.27.9和Go 1.28.3都太旧。

Commands that incorporate new module versions that require new Go versions write the new minimum `go` version requirement to the current workspace’s `go.work` file or the main module’s `go.mod` file, updating the `go` line. For [repeatability](https://research.swtch.com/vgo-principles#repeatability), any command that updates the `go` line also updates the `toolchain` line to record its own toolchain name. The next time the `go` command runs in that workspace or module, it will use that updated `toolchain` line during [toolchain selection](https://go.dev/doc/toolchain#select).

​	需要包含新模块版本的命令会将新的最低`go`版本要求写入当前工作区的`go.work`文件或主模块的`go.mod`文件，更新`go`行。出于[可重复性](https://research.swtch.com/vgo-principles#repeatability)的考虑，更新`go`行的任何命令也会更新`toolchain`行，以记录其自身的工具链名称。下次在该工作区或模块中运行`go`命令时，它将在[工具链选择](https://go.dev/doc/toolchain#select)期间使用该更新的`toolchain`行。

For example, `go get example.com/widget@v1.2.3` may print a switching notice like above and switch to Go 1.27.9. Go 1.27.9 will complete the `go get` and update the `toolchain` line to say `toolchain go1.27.9`. The next `go` command run in that module or workspace will select `go1.27.9` during startup and will not print any switching message.

​	例如，`go get example.com/widget@v1.2.3`可能会像上面那样打印切换通知并切换到Go 1.27.9。Go 1.27.9将完成`go get`并将`toolchain`行更新为`toolchain go1.27.9`。在该模块或工作区中下一次运行`go`命令时，将在启动时选择`go1.27.9`，并且不会打印任何切换消息。

In general, if any `go` command is run twice, if the first prints a switching message, the second will not, because the first also updated `go.work` or `go.mod` to select the right toolchain at startup. The exception is the `go install package@version` and `go run package@version` forms, which run in no workspace or main module and cannot write a `toolchain` line. They print a switching message every time they need to switch to a newer toolchain.

​	一般来说，如果运行任何`go`命令两次，如果第一个命令打印了切换消息，则第二个命令将不会，因为第一个命令还会更新`go.work`或`go.mod`以在启动时选择正确的工具链。例外情况是`go install package@version`和`go run package@version`形式，它们在没有工作区或主模块的情况下运行，无法编写`toolchain`行。它们每次需要切换到更新的工具链时都会打印切换消息。

## 下载工具链 Downloading toolchains

When using `GOTOOLCHAIN=auto` or `GOTOOLCHAIN=<name>+auto`, the Go command downloads newer toolchains as needed. These toolchains are packaged as special modules with module path `golang.org/toolchain` and version `v0.0.1-go*VERSION*.*GOOS*-*GOARCH*`. Toolchains are downloaded like any other module, meaning that toolchain downloads can be proxied by setting `GOPROXY` and have their checksums checked by the Go checksum database. Because the specific toolchain used depends on the system’s own default toolchain as well as the local operating system and architecture (GOOS and GOARCH), it is not practical to write toolchain module checksums to `go.sum`. Instead, toolchain downloads fail for lack of verification if `GOSUMDB=off`. `GOPRIVATE` and `GONOSUMDB` patterns do not apply to the toolchain downloads.

​	当使用`GOTOOLCHAIN=auto`或`GOTOOLCHAIN=<name>+auto`时，Go命令会根据需要下载更新的工具链。这些工具链被打包为带有模块路径`golang.org/toolchain`和版本`v0.0.1-go*VERSION*.*GOOS*-*GOARCH*`的特殊模块。工具链的下载方式与任何其他模块相同，这意味着可以通过设置`GOPROXY`来代理工具链的下载，并且可以通过Go校验和数据库来检查其校验和。由于特定的工具链取决于系统自己的默认工具链以及本地操作系统和体系结构（GOOS和GOARCH），所以将工具链模块的校验和写入`go.sum`是不实际的。相反，在`GOSUMDB=off`的情况下，工具链下载因验证不足而失败。`GOPRIVATE`和`GONOSUMDB`模式不适用于工具链下载。

## 使用`go get`管理Go版本模块要求 Managing Go version module requirements with `go get`

In general the `go` command treats the `go` and `toolchain` lines as declaring versioned toolchain dependencies of the main module. The `go get` command can manage these lines just as it manages the `require` lines that specify versioned module dependencies.

​	一般来说，`go`命令将`go`和`toolchain`行视为主模块的版本化工具链依赖关系。`go get`命令可以像管理指定版本化模块依赖项的`require`行一样管理这些行。

For example, `go get go@1.22.1 toolchain@1.24rc1` changes the main module’s `go.mod` file to read `go 1.22.1` and `toolchain go1.24rc1`.

​	例如，`go get go@1.22.1 toolchain@1.24rc1`会将主模块的`go.mod`文件更改为`go 1.22.1`和`toolchain go1.24rc1`。

The `go` command understands that the `go` dependency requires a `toolchain` dependency with a greater or equal Go version.

​	`go`命令理解`go`依赖关系需要具有更高或等于Go版本的`toolchain`依赖关系。

Continuing the example, a later `go get go@1.25.0` will update the toolchain to `go1.25.0` as well. When the toolchain matches the `go` line exactly, it can be omitted and implied, so this `go get` will delete the `toolchain` line.

​	继续上面的示例，稍后运行`go get go@1.25.0`也会将工具链更新为`go1.25.0`。当工具链与`go`行完全匹配时，可以省略并隐含它，因此此`go get`会删除`toolchain`行。

The same requirement applies in reverse when downgrading: if the `go.mod` starts at `go 1.22.1` and `toolchain go1.24rc1`, then `go get toolchain@go1.22.9` will update only the `toolchain` line, but `go get toolchain@go1.21.3` will downgrade the `go` line to `go 1.21.3` as well. The effect will be to leave just `go 1.21.3` with no `toolchain` line.

​	在降级时，同样的要求适用于相反情况：如果`go.mod`从`go 1.22.1`开始，并且运行`go get toolchain@go1.22.9`，那么只会更新`toolchain`行，但是运行`go get toolchain@go1.21.3`会将`go`行降级为`go 1.21.3`。效果将是仅留下`go 1.21.3`，没有`toolchain`行。

The special form `toolchain@none` means to remove any `toolchain` line, as in `go get toolchain@none` or `go get go@1.25.0 toolchain@none`.

​	特殊形式`toolchain@none`表示删除任何`toolchain`行，例如`go get toolchain@none`或`go get go@1.25.0 toolchain@none`。

The `go` command understands the version syntax for `go` and `toolchain` dependencies as well as queries.

​	`go`命令同样理解`go`和`toolchain`依赖的版本语法以及查询。

For example, just as `go get example.com/widget@v1.2` uses the latest `v1.2` version of `example.com/widget` (perhaps `v1.2.3`), `go get go@1.22` uses the latest available release of the Go 1.22 language version (perhaps `1.22rc3`, or perhaps `1.22.3`). The same applies to `go get toolchain@go1.22`.

​	例如，正如`go get example.com/widget@v1.2`使用`example.com/widget`的最新`v1.2`版本（可能是`v1.2.3`）一样，`go get go@1.22`使用Go 1.22语言版本的最新可用版本（可能是`1.22rc3`，或者可能是`1.22.3`）。对于`go get toolchain@go1.22`也是同样适用的。

The `go get` and `go mod tidy` commands maintain the `go` line to be greater than or equal to the `go` line of any required dependency module.

​	`go get`和`go mod tidy`命令会保持`go`行大于或等于任何所需依赖模块的`go`行。

For example, if the main module has `go 1.22.1` and we run `go get example.com/widget@v1.2.3` which declares `go 1.24rc1`, then `go get` will update the main module’s `go` line to `go 1.24rc1`.

​	例如，如果主模块的`go 1.22.1`，并且运行`go get example.com/widget@v1.2.3`，该模块声明了`go 1.24rc1`，那么`go get`会将主模块的`go`行更新为`go 1.24rc1`。

Continuing the example, a later `go get go@1.22.1` will downgrade `example.com/widget` to a version compatible with Go 1.22.1 or else remove the requirement entirely, just as it would when downgrading any other dependency of `example.com/widget`.

​	继续上面的示例，稍后运行`go get go@1.22.1`将会将`example.com/widget`降级为与Go 1.22.1兼容的版本，否则会完全删除该要求，就像降级`example.com/widget`的任何其他依赖项一样。

Before Go 1.21, the suggested way to update a module to a new Go version (say, Go 1.22) was `go mod tidy -go=1.22`, to make sure that any adjustments specific to Go 1.22 were made to the `go.mod` at the same time that the `go` line is updated. That form is still valid, but the simpler `go get go@1.22` is now preferred.

​	在Go 1.21之前，更新模块到新的Go版本（例如Go 1.22）的建议方法是`go mod tidy -go=1.22`，以确保在更新`go`行的同时对`go.mod`进行适应Go 1.22的任何调整。这种形式仍然有效，但更简单的`go get go@1.22`现在更受欢迎。

When `go get` is run in a module in a directory contained in a workspace root, `go get` mostly ignores the workspace, but it does update the `go.work` file to upgrade the `go` line when the workspace would otherwise be left with too old a `go` line.

​	当在包含在工作区根目录中的目录中运行`go get`时，`go get`大多会忽略工作区，但是它会在工作区被留下太旧的`go`行的情况下，更新`go.work`文件以升级`go`行。

## 使用`go work`管理Go版本工作区要求 Managing Go version workspace requirements with `go work`

As noted in the previous section, `go get` run in a directory inside a workspace root will take care to update the `go.work` file’s `go` line as needed to be greater than or equal to any module inside that root. However, workspaces can also refer to modules outside the root directory; running `go get` in those directories may result in an invalid workspace configuration, one in which the `go` version declared in `go.work` is less than one or more of the modules in the `use` directives.

​	如前一节所述，运行`go get`在工作区根目录内的目录中会注意更新`go.work`文件的`go`行，以便大于或等于该根目录内任何模块的`go`行。但是，工作区也可以引用根目录外的模块；在这些目录中运行`go get`可能会导致无效的工作区配置，其中`go.work`文件中声明的`go`版本小于一个或多个`use`指令中的模块。

The command `go work use`, which adds new `use` directives, also checks that the `go` version in the `go.work` file is new enough for all the existing `use` directives. To update a workspace that has gotten its `go` version out of sync with its modules, run `go work use` with no arguments.

​	`go work use`命令会添加新的`use`指令，并且还会检查`go.work`文件中的`go`版本是否足够新，以适应所有现有的`use`指令。要更新`go`版本与模块不同步的工作区，可以不带参数地运行`go work use`。

The commands `go work init` and `go work sync` also update the `go` version as needed.

​	`go work init`和`go work sync`命令也会根据需要更新`go`版本。

To remove the `toolchain` line from a `go.work` file, use `go work edit -toolchain=none`.

​	要从`go.work`文件中删除`toolchain`行，可以使用`go work edit -toolchain=none`。