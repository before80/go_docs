+++
title = "Go 1.21中的向后兼容性和工具链管理"
date = 2023-08-21T15:03:30+08:00
weight = 86
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Forward Compatibility and Toolchain Management in Go 1.21 - Go 1.21中的向后兼容性和工具链管理

https://go.dev/blog/toolchain

Russ Cox
14 August 2023

Russ Cox 2023年8月14日



Beyond Go 1.21’s [expanded commitment to backward compatibility]({{< ref "/goBlog/2023/BackwardCompatibilityGo1_21AndGo2">}}), Go 1.21 also introduces better forward compatibility for Go code, meaning that Go 1.21 and later will take better care not to miscompile code that requires an even newer version of Go. Specifically, the `go` line in `go.mod` now specifies a minimum required Go toolchain version, while in previous releases it was a mostly unenforced suggestion.

​	除了Go 1.21对[向后兼容性的扩展承诺]({{< ref "/goBlog/2023/BackwardCompatibilityGo1_21AndGo2">}})之外，Go 1.21还引入了更好的Go代码向后兼容性，这意味着Go 1.21及更高版本将更好地避免对需要更新版本的Go的代码进行错误编译。具体而言，`go.mod`文件中的`go`行现在指定了所需的最低Go工具链版本，而在以前的版本中，这只是一个大致没有强制执行的建议。

To make it easier to keep up with these requirements, Go 1.21 also introduces toolchain management, so that different modules can use different Go toolchains just as they can use different versions of a required module. After installing Go 1.21, you’ll never have to manually download and install a Go toolchain again. The `go` command can do it for you.

​	为了更容易满足这些要求，Go 1.21还引入了工具链管理，使不同的模块可以使用不同的Go工具链，就像它们可以使用不同版本的所需模块一样。安装了Go 1.21后，您将不再需要手动下载和安装Go工具链。`go`命令可以为您完成这项工作。

The rest of this post describes both of these Go 1.21 changes in more detail.

​	本文的其余部分将更详细地描述Go 1.21中的这两个变更。

## 向后兼容性 Forward Compatibility

Forward compatibility refers to what happens when a Go toolchain attempts to build Go code intended for a newer version of Go. If my program depends on a module M and needs a bug fix added in M v1.2.3, I can add `require M v1.2.3` to my `go.mod`, guaranteeing that my program won’t be compiled against older versions of M. But if my program requires a particular version of Go, there hasn’t been any way to express that: in particular, the `go.mod` `go` line did not express that.

​	向后兼容性指的是当Go工具链试图构建用于较新Go版本的Go代码时会发生什么。如果我的程序依赖于一个模块M并且需要在M v1.2.3中添加的错误修复，我可以在我的`go.mod`中添加`require M v1.2.3`，确保我的程序不会与旧版本的M编译。但是，如果我的程序需要特定版本的Go，以前没有任何方法来表达这一点：特别是`go.mod`中的`go`行没有表达出这一点。

For example, if I write code that uses the new generics added in Go 1.18, I can write `go 1.18` in my `go.mod` file, but that won’t stop earlier versions of Go from trying to compile the code, producing errors like:

​	例如，如果我编写的代码使用了Go 1.18中添加的新泛型，我可以在我的`go.mod`文件中写上`go 1.18`，但是这不会阻止早期版本的Go尝试编译代码，产生如下错误：

```bash
$ cat go.mod
go 1.18
module example

$ go version
go version go1.17

$ go build
# example
./x.go:2:6: missing function body
./x.go:2:7: syntax error: unexpected [, expecting (
note: module requires Go 1.18
$
```

The two compiler errors are misleading noise. The real problem is printed by the `go` command as a hint: the program failed to compile, so the `go` command points out the potential version mismatch.

​	这两个编译器错误是误导性的噪音。真正的问题由`go`命令打印出作为提示：程序无法编译，因此`go`命令指出了潜在的版本不匹配问题。

In this example, we’re lucky the build failed. If I write code that only runs correctly in Go 1.19 or later, because it depends on a bug fixed in that patch release, but I’m not using any Go 1.19-specific language features or packages in the code, earlier versions of Go will compile it and silently succeed.

​	在这个例子中，我们幸运地构建失败了。如果我编写的代码只在Go 1.19或更高版本中正确运行，因为它依赖于在该补丁版本中修复的错误，但是在代码中没有使用任何Go 1.19特定的语言特性或包，早期版本的Go将编译它并默默地成功。

Starting in Go 1.21, Go toolchains will treat the `go` line in `go.mod` not as a guideline but as a rule, and the line can list specific point releases or release candidates. That is, Go 1.21.0 understands that it cannot even build code that says `go 1.21.1` in its `go.mod` file, not to mention code that says much later versions like `go 1.22.0`.

​	从Go 1.21开始，Go工具链将不再将`go.mod`中的`go`行视为一个指导方针，而是视为一个规则，该行可以列出特定的点发布或发布候选版本。也就是说，Go 1.21.0知道它甚至不能构建代码，该代码在其`go.mod`文件中说`go 1.21.1`，更不用说说像`go 1.22.0`这样的版本了。

The main reason we allowed older versions of Go to try to compile newer code was to avoid unnecessary build failures. It’s very frustrating to be told that your version of Go is too old to build a program, especially if it might work anyway (maybe the requirement is unnecessarily conservative), and especially when updating to a newer Go version is a bit of a chore. To reduce the impact of enforcing the `go` line as a requirement, Go 1.21 adds toolchain management to the core distribution as well.

​	我们允许旧版本的Go尝试编译更新的代码的主要原因是为了避免不必要的构建失败。被告知您的Go版本太旧，无法构建程序非常令人沮丧，尤其是如果它可能仍然可以工作（也许要求过于保守），尤其是更新到较新的Go版本也是一项艰巨的任务。为了减少将`go`行强制执行为要求的影响，Go 1.21还将工具链管理添加到核心发行版中。

## 工具链管理 Toolchain Management

When you need a new version of a Go module, the `go` command downloads it for you. Starting in Go 1.21, when you need a newer Go toolchain, the `go` command downloads that for you too. This functionality is like Node’s `nvm` or Rust’s `rustup`, but built in to the core `go` command instead of being a separate tool.

​	当您需要新版本的Go模块时，`go`命令会为您下载它。从Go 1.21开始，当您需要更新的Go工具链时，`go`命令也会为您下载它。这个功能类似于Node的`nvm`或Rust的`rustup`，但内置到核心的`go`命令中，而不是作为一个独立的工具。

If you are running Go 1.21.0 and you run a `go` command, say, `go build`, in a module with a `go.mod` that says `go 1.21.1`, the Go 1.21.0 `go` command will notice that you need Go 1.21.1, download it, and re-invoke that version’s `go` command to finish the build. When the `go` command downloads and runs these other toolchains, it doesn’t install them in your PATH or overwrite the current installation. Instead, it downloads them as Go modules, inheriting all the [security and privacy benefits of modules]({{< ref "/goBlog/2019/ModuleMirrorAndChecksumDatabaseLaunched">}}), and then it runs them from the module cache.

​	如果您运行的是Go 1.21.0，并且在一个`go.mod`文件中运行`go`命令，例如`go build`，该文件中写着`go 1.21.1`，则Go 1.21.0 `go`命令将注意到您需要Go 1.21.1，下载它，并重新调用该版本的`go`命令以完成构建。当`go`命令下载并运行这些其他工具链时，它不会将它们安装在您的PATH中，也不会覆盖当前的安装。相反，它会将它们作为Go模块进行下载，继承所有[模块的安全性和隐私优势]({{< ref "/goBlog/2019/ModuleMirrorAndChecksumDatabaseLaunched">}})，然后从模块缓存中运行它们。

There is also a new `toolchain` line in `go.mod` that specifies the minimum Go toolchain to use when working in a particular module. In contrast to the `go` line, `toolchain` does not impose a requirement on other modules. For example, a `go.mod` might say:

​	`go.mod`中还有一个新的`toolchain`行，指定在特定模块中工作时要使用的最低Go工具链。与`go`行相比，`toolchain`不会对其他模块施加要求。例如，一个`go.mod`文件可能会这样写：

```
module m
go 1.21.0
toolchain go1.21.4
```

This says that other modules requiring `m` need to provide at least Go 1.21.0, but when we are working in `m` itself, we want an even newer toolchain, at least Go 1.21.4.

​	这意味着其他需要`m`的模块至少需要提供Go 1.21.0，但是当我们在`m`本身工作时，我们希望使用更新的工具链，至少是Go 1.21.4。

The `go` and `toolchain` requirements can be updated using `go get` like ordinary module requirements. For example, if you’re using one of the Go 1.21 release candidates, you can start using Go 1.21.0 in a particular module by running:

​	`go`和`toolchain`的要求可以像普通模块要求一样使用`go get`进行更新。例如，如果您正在使用Go 1.21的一个发行候选版本，您可以通过运行以下命令在特定模块中开始使用Go 1.21.0：

```
go get go@1.21.0
```

That will download and run Go 1.21.0 to update the `go` line, and future invocations of the `go` command will see the line `go 1.21.0` and automatically re-invoke that version.

​	这将下载并运行Go 1.21.0以更新`go`行，`go`命令的未来调用将看到行`go 1.21.0`并自动重新调用该版本。

Or if you want to start using Go 1.21.0 in a module but leave the `go` line set to an older version, to help maintain compatibility with users of earlier versions of Go, you can update the `toolchain` line:

​	或者，如果您想在模块中开始使用Go 1.21.0，但保留`go`行设置为旧版本，以便与早期版本的Go用户保持兼容性，您可以更新`toolchain`行：

```
go get toolchain@go1.21.0
```

If you’re ever wondering which Go version is running in a particular module, the answer is the same as before: run `go version`.

​	如果您想知道特定模块中正在运行的Go版本是哪个版本，答案与以前相同：运行`go version`。

You can force the use of a specific Go toolchain version using the GOTOOLCHAIN environment variable. For example, to test code with Go 1.20.4:

​	您可以使用GOTOOLCHAIN环境变量强制使用特定的Go工具链版本。例如，要使用Go 1.20.4测试代码：

```
GOTOOLCHAIN=go1.20.4 go test
```

Finally, a GOTOOLCHAIN setting of the form `version+auto` means to use `version` by default but allow upgrades to newer versions as well. If you have Go 1.21.0 installed, then when Go 1.21.1 is released, you can change your system default by setting a default GOTOOLCHAIN:

​	最后，形式为`version+auto`的GOTOOLCHAIN设置意味着默认情况下使用`version`，但也允许升级到更新版本。如果您已安装了Go 1.21.0，那么在发布Go 1.21.1时，您可以通过设置默认的GOTOOLCHAIN来更改系统默认设置：

```
go env -w GOTOOLCHAIN=go1.21.1+auto
```

You’ll never have to manually download and install a Go toolchain again. The `go` command will take care of it for you.

​	您将不再需要手动下载和安装Go工具链。`go`命令会为您处理这一切。

See “[Go Toolchains]({{< ref "/docs/GoToolchains">}})” for more details.

有关更多详细信息，请参阅“[Go工具链]({{< ref "/docs/GoToolchains">}})”。