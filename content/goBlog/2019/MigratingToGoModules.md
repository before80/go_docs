+++
title = "迁移到 go 模块"
weight = 9
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Migrating to Go Modules - 迁移到 go 模块

https://go.dev/blog/migrating-to-go-modules

Jean de Klerk
21 August 2019

## Introduction 简介

This post is part 2 in a series.

这篇文章是系列文章的第二部分。

- Part 1 — [Using Go Modules 使用Go模块](https://go.dev/blog/using-go-modules)
- **Part 2 — Migrating To Go Modules** (this post) 迁移到Go模块（本帖）
- Part 3 — [Publishing Go Modules 发布Go模块](https://go.dev/blog/publishing-go-modules)
- Part 4 — [Go Modules: v2 and Beyond Go模块：V2版及以后](https://go.dev/blog/v2-go-modules)
- Part 5 — [Keeping Your Modules Compatible 保持模块的兼容性](https://go.dev/blog/module-compatibility)

**Note:** For documentation, see [Managing dependencies](https://go.dev/doc/modules/managing-dependencies) and [Developing and publishing modules](https://go.dev/doc/modules/developing).

注意：关于文档，请参见管理依赖项和开发与发布模块。

Go projects use a wide variety of dependency management strategies. [Vendoring](https://go.dev/cmd/go/#hdr-Vendor_Directories) tools such as [dep](https://github.com/golang/dep) and [glide](https://github.com/Masterminds/glide) are popular, but they have wide differences in behavior and don’t always work well together. Some projects store their entire GOPATH directory in a single Git repository. Others simply rely on `go get` and expect fairly recent versions of dependencies to be installed in GOPATH.

Go项目使用了多种多样的依赖项管理策略。诸如dep和glide这样的销售工具很受欢迎，但它们在行为上有很大的差异，而且并不总是能很好地协同工作。有些项目将整个GOPATH目录存储在一个Git仓库中。其他项目则简单地依赖go get，并期望相当新的依赖版本被安装在GOPATH中。

Go’s module system, introduced in Go 1.11, provides an official dependency management solution built into the `go` command. This article describes tools and techniques for converting a project to modules.

Go的模块系统在Go 1.11中引入，提供了一个内置于go命令的官方依赖项管理解决方案。本文介绍了将项目转换为模块的工具和技术。

Please note: if your project is already tagged at v2.0.0 or higher, you will need to update your module path when you add a `go.mod` file. We’ll explain how to do that without breaking your users in a future article focused on v2 and beyond.

请注意：如果您的项目已经被标记为v2.0.0或更高版本，您需要在添加go.mod文件时更新您的模块路径。我们会在以后的文章中解释如何在不破坏用户的情况下做到这一点，主要是针对v2及以上版本。

## Migrating to Go modules in your project 在您的项目中迁移到Go模块

A project might be in one of three states when beginning the transition to Go modules:

在开始向 Go 模块过渡时，项目可能处于以下三种状态之一：

- A brand new Go project. 一个全新的 Go 项目。
- An established Go project with a non-modules dependency manager. 一个已经建立的 Go 项目，有一个非模块的依赖管理器。
- An established Go project without any dependency manager. 已建立的 Go 项目，没有任何依赖关系管理器。

The first case is covered in [Using Go Modules](https://blog.golang.org/using-go-modules); we’ll address the latter two in this post.

第一种情况在《使用 Go 模块》中有所介绍；我们将在这篇文章中讨论后两种情况。

## With a dependency manager 使用依赖管理器

To convert a project that already uses a dependency management tool, run the following commands:

要转换一个已经使用依赖管理工具的项目，请运行以下命令：

```shell linenums="1"
$ git clone https://github.com/my/project
[...]
$ cd project
$ cat Godeps/Godeps.json
{
    "ImportPath": "github.com/my/project",
    "GoVersion": "go1.12",
    "GodepVersion": "v80",
    "Deps": [
        {
            "ImportPath": "rsc.io/binaryregexp",
            "Comment": "v0.2.0-1-g545cabd",
            "Rev": "545cabda89ca36b48b8e681a30d9d769a30b3074"
        },
        {
            "ImportPath": "rsc.io/binaryregexp/syntax",
            "Comment": "v0.2.0-1-g545cabd",
            "Rev": "545cabda89ca36b48b8e681a30d9d769a30b3074"
        }
    ]
}
$ go mod init github.com/my/project
go: creating new go.mod: module github.com/my/project
go: copying requirements from Godeps/Godeps.json
$ cat go.mod
module github.com/my/project

go 1.12

require rsc.io/binaryregexp v0.2.1-0.20190524193500-545cabda89ca
$
```

`go mod init` creates a new go.mod file and automatically imports dependencies from `Godeps.json`, `Gopkg.lock`, or a number of [other supported formats](https://go.googlesource.com/go/+/362625209b6cd2bc059b6b0a67712ddebab312d9/src/cmd/go/internal/modconv/modconv.go#9). The argument to `go mod init` is the module path, the location where the module may be found.

go mod init会创建一个新的go.mod文件，并自动从Godeps.json、Gopkg.lock或其他一些支持的格式中导入依赖项。go mod init的参数是模块路径，即可以找到模块的位置。

This is a good time to pause and run `go build ./...` and `go test ./...` before continuing. Later steps may modify your `go.mod` file, so if you prefer to take an iterative approach, this is the closest your `go.mod` file will be to your pre-modules dependency specification.

这是一个暂停的好时机，在继续之前运行go build ./...和go test ./...。以后的步骤可能会修改您的go.mod文件，所以如果您喜欢采取迭代的方法，这是您的go.mod文件最接近于您的模块前依赖项规范。

```shell linenums="1"
$ go mod tidy
go: downloading rsc.io/binaryregexp v0.2.1-0.20190524193500-545cabda89ca
go: extracting rsc.io/binaryregexp v0.2.1-0.20190524193500-545cabda89ca
$ cat go.sum
rsc.io/binaryregexp v0.2.1-0.20190524193500-545cabda89ca h1:FKXXXJ6G2bFoVe7hX3kEX6Izxw5ZKRH57DFBJmHCbkU=
rsc.io/binaryregexp v0.2.1-0.20190524193500-545cabda89ca/go.mod h1:qTv7/COck+e2FymRvadv62gMdZztPaShugOCi3I+8D8=
$
```

`go mod tidy` finds all the packages transitively imported by packages in your module. It adds new module requirements for packages not provided by any known module, and it removes requirements on modules that don’t provide any imported packages. If a module provides packages that are only imported by projects that haven’t migrated to modules yet, the module requirement will be marked with an `// indirect` comment. It is always good practice to run `go mod tidy` before committing a `go.mod` file to version control.

go mod tidy 找到所有被您的模块中的软件包过渡性导入的软件包。它为任何已知模块没有提供的包添加新的模块需求，并删除没有提供任何导入包的模块的需求。如果一个模块提供的包只被那些还没有迁移到模块的项目所导入，那么这个模块的需求将被标记为//间接注释。在将go.mod文件提交到版本控制之前，运行go mod tidy总是一个好的做法。

Let’s finish by making sure the code builds and tests pass:

最后，让我们确保代码的构建和测试通过：

```shell linenums="1"
$ go build ./...
$ go test ./...
[...]
$
```

Note that other dependency managers may specify dependencies at the level of individual packages or entire repositories (not modules), and generally do not recognize the requirements specified in the `go.mod` files of dependencies. Consequently, you may not get exactly the same version of every package as before, and there’s some risk of upgrading past breaking changes. Therefore, it’s important to follow the above commands with an audit of the resulting dependencies. To do so, run

请注意，其他的依赖管理器可能会在单个包或整个软件库（而不是模块）的层面上指定依赖关系，并且通常不会识别依赖关系的go.mod文件中指定的要求。因此，您可能不会得到与以前完全相同的每个软件包的版本，而且有一些升级过去破坏变化的风险。因此，在执行上述命令后，对所产生的依赖关系进行审计是很重要的。要做到这一点，请运行

```shell linenums="1"
$ go list -m all
go: finding rsc.io/binaryregexp v0.2.1-0.20190524193500-545cabda89ca
github.com/my/project
rsc.io/binaryregexp v0.2.1-0.20190524193500-545cabda89ca
$
```

and compare the resulting versions with your old dependency management file to ensure that the selected versions are appropriate. If you find a version that wasn’t what you wanted, you can find out why using `go mod why -m` and/or `go mod graph`, and upgrade or downgrade to the correct version using `go get`. (If the version you request is older than the version that was previously selected, `go get` will downgrade other dependencies as needed to maintain compatibility.) For example,

并将得到的版本与您的旧的依赖项管理文件进行比较，以确保所选的版本是合适的。如果您发现一个版本不是您想要的，您可以用go mod why -m和/或go mod graph找出原因，然后用go get升级或降级到正确的版本。(如果您要求的版本比之前选择的版本要老，go get会根据需要降级其他依赖关系以保持兼容性）。比如说，

```shell linenums="1"
$ go mod why -m rsc.io/binaryregexp
[...]
$ go mod graph | grep rsc.io/binaryregexp
[...]
$ go get rsc.io/binaryregexp@v0.2.0
$
```

## Without a dependency manager 没有依赖项管理系统

For a Go project without a dependency management system, start by creating a `go.mod` file:

对于没有依赖项管理系统的Go项目，首先要创建一个go.mod文件：

```shell linenums="1"
$ git clone https://go.googlesource.com/blog
[...]
$ cd blog
$ go mod init golang.org/x/blog
go: creating new go.mod: module golang.org/x/blog
$ cat go.mod
module golang.org/x/blog

go 1.12
$
```

Without a configuration file from a previous dependency manager, `go mod init` will create a `go.mod` file with only the `module` and `go` directives. In this example, we set the module path to `golang.org/x/blog` because that is its [custom import path](https://go.dev/cmd/go/#hdr-Remote_import_paths). Users may import packages with this path, and we must be careful not to change it.

如果没有之前依赖管理器的配置文件，go mod init将创建一个只有模块和go指令的go.mod文件。在这个例子中，我们将模块路径设置为golang.org/x/blog，因为这是它的自定义导入路径。用户可以用这个路径导入包，我们必须注意不要改变它。

The `module` directive declares the module path, and the `go` directive declares the expected version of the Go language used to compile the code within the module.

模块指令声明了模块路径，而go指令声明了用于编译模块内代码的Go语言的预期版本。

Next, run `go mod tidy` to add the module’s dependencies:

接下来，运行 go mod tidy 来添加模块的依赖项：

```shell linenums="1"
$ go mod tidy
go: finding golang.org/x/website latest
go: finding gopkg.in/tomb.v2 latest
go: finding golang.org/x/net latest
go: finding golang.org/x/tools latest
go: downloading github.com/gorilla/context v1.1.1
go: downloading golang.org/x/tools v0.0.0-20190813214729-9dba7caff850
go: downloading golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7
go: extracting github.com/gorilla/context v1.1.1
go: extracting golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7
go: downloading gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637
go: extracting gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637
go: extracting golang.org/x/tools v0.0.0-20190813214729-9dba7caff850
go: downloading golang.org/x/website v0.0.0-20190809153340-86a7442ada7c
go: extracting golang.org/x/website v0.0.0-20190809153340-86a7442ada7c
$ cat go.mod
module golang.org/x/blog

go 1.12

require (
    github.com/gorilla/context v1.1.1
    golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7
    golang.org/x/text v0.3.2
    golang.org/x/tools v0.0.0-20190813214729-9dba7caff850
    golang.org/x/website v0.0.0-20190809153340-86a7442ada7c
    gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637
)
$ cat go.sum
cloud.google.com/go v0.26.0/go.mod h1:aQUYkXzVsufM+DwF1aE+0xfcU+56JwCaLick0ClmMTw=
cloud.google.com/go v0.34.0/go.mod h1:aQUYkXzVsufM+DwF1aE+0xfcU+56JwCaLick0ClmMTw=
git.apache.org/thrift.git v0.0.0-20180902110319-2566ecd5d999/go.mod h1:fPE2ZNJGynbRyZ4dJvy6G277gSllfV2HJqblrnkyeyg=
git.apache.org/thrift.git v0.0.0-20181218151757-9b75e4fe745a/go.mod h1:fPE2ZNJGynbRyZ4dJvy6G277gSllfV2HJqblrnkyeyg=
github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973/go.mod h1:Dwedo/Wpr24TaqPxmxbtue+5NUziq4I4S80YR8gNf3Q=
[...]
$
```

`go mod tidy` added module requirements for all the packages transitively imported by packages in your module and built a `go.sum` with checksums for each library at a specific version. Let’s finish by making sure the code still builds and tests still pass:

go mod tidy为所有被您的模块中的包转接导入的包添加了模块需求，并建立了一个go.sum，其中包括每个库的特定版本的校验和。最后，让我们确保代码仍然可以构建，测试仍然可以通过：

```go linenums="1"
$ go build ./...
$ go test ./...
ok      golang.org/x/blog   0.335s
?       golang.org/x/blog/content/appengine [no test files]
ok      golang.org/x/blog/content/cover 0.040s
?       golang.org/x/blog/content/h2push/server [no test files]
?       golang.org/x/blog/content/survey2016    [no test files]
?       golang.org/x/blog/content/survey2017    [no test files]
?       golang.org/x/blog/support/racy  [no test files]
$
```

Note that when `go mod tidy` adds a requirement, it adds the latest version of the module. If your `GOPATH` included an older version of a dependency that subsequently published a breaking change, you may see errors in `go mod tidy`, `go build`, or `go test`. If this happens, try downgrading to an older version with `go get` (for example, `go get github.com/broken/module@v1.1.0`), or take the time to make your module compatible with the latest version of each dependency.

注意，当go mod tidy添加一个需求时，它添加的是该模块的最新版本。如果您的GOPATH包含了一个旧版本的依赖，而这个依赖后来发布了一个突破性的变化，您可能会在go mod tidy、go build或go test中看到错误。如果发生这种情况，试着用go get降级到旧版本（例如，go get github.com/broken/module@v1.1.0），或者花时间使您的模块与每个依赖的最新版本兼容。

### Tests in module mode 模块模式下的测试

Some tests may need tweaks after migrating to Go modules.

有些测试在迁移到Go模块后可能需要进行调整。

If a test needs to write files in the package directory, it may fail when the package directory is in the module cache, which is read-only. In particular, this may cause `go test all` to fail. The test should copy files it needs to write to a temporary directory instead.

如果一个测试需要写入包目录中的文件，当包目录在模块缓存中时，它可能会失败，因为模块缓存是只读的。特别是，这可能导致go测试全部失败。测试应该把它需要写的文件复制到一个临时目录中。

If a test relies on relative paths (`../package-in-another-module`) to locate and read files in another package, it will fail if the package is in another module, which will be located in a versioned subdirectory of the module cache or a path specified in a `replace` directive. If this is the case, you may need to copy the test inputs into your module, or convert the test inputs from raw files to data embedded in `.go` source files.

如果一个测试依靠相对路径（.../package-in-another-module）来定位和读取另一个包中的文件，如果该包在另一个模块中，它将失败，该模块将位于模块缓存的一个版本子目录或替换指令中指定的路径。如果是这种情况，您可能需要将测试输入复制到您的模块中，或者将测试输入从原始文件转换成嵌入.go源文件的数据。

If a test expects `go` commands within the test to run in GOPATH mode, it may fail. If this is the case, you may need to add a `go.mod` file to the source tree to be tested, or set `GO111MODULE=off` explicitly.

如果一个测试期望测试中的go命令在GOPATH模式下运行，它可能会失败。如果是这种情况，您可能需要在要测试的源代码树上添加一个go.mod文件，或者明确设置GO111MODULE=off。

## Publishing a release 发布一个版本

Finally, you should tag and publish a release version for your new module. This is optional if you haven’t released any versions yet, but without an official release, downstream users will depend on specific commits using [pseudo-versions](https://go.dev/cmd/go/#hdr-Pseudo_versions), which may be more difficult to support.

最后，您应该为您的新模块标记并发布一个发布版本。如果您还没有发布任何版本，这是可选的，但如果没有正式的发布版本，下游用户将依赖于使用伪版本的特定提交，这可能更难支持。

```
$ git tag v1.2.0
$ git push origin v1.2.0
```

Your new `go.mod` file defines a canonical import path for your module and adds new minimum version requirements. If your users are already using the correct import path, and your dependencies haven’t made breaking changes, then adding the `go.mod` file is backwards-compatible — but it’s a significant change, and may expose existing problems. If you have existing version tags, you should increment the [minor version](https://semver.org/#spec-item-7). See [Publishing Go Modules](https://go.dev/blog/publishing-go-modules) to learn how to increment and publish versions.

您的新go.mod文件为您的模块定义了一个规范的导入路径，并增加了新的最低版本要求。如果您的用户已经在使用正确的导入路径，而且您的依赖关系也没有发生破坏性的变化，那么添加go.mod文件是向后兼容的--但这是一个重大变化，可能会暴露现有的问题。如果您有现有的版本标签，您应该增加次要版本。请参阅发布Go模块以了解如何增加和发布版本。

## Imports and canonical module paths 进口和规范的模块路径

Each module declares its module path in its `go.mod` file. Each `import` statement that refers to a package within the module must have the module path as a prefix of the package path. However, the `go` command may encounter a repository containing the module through many different [remote import paths](https://go.dev/cmd/go/#hdr-Remote_import_paths). For example, both `golang.org/x/lint` and `github.com/golang/lint` resolve to repositories containing the code hosted at [go.googlesource.com/lint](https://go.googlesource.com/lint). The [`go.mod` file](https://go.googlesource.com/lint/+/refs/heads/master/go.mod) contained in that repository declares its path to be `golang.org/x/lint`, so only that path corresponds to a valid module.

每个模块在其go.mod文件中声明其模块路径。每个引用模块中的包的导入语句必须将模块路径作为包路径的前缀。然而，go命令可能会通过许多不同的远程导入路径遇到包含该模块的仓库。例如，golang.org/x/lint和github.com/golang/lint都解析为包含代码的仓库，托管在go.googlesource.com/lint。该仓库中包含的go.mod文件声明其路径为golang.org/x/lint，所以只有该路径对应的模块才是有效的。

Go 1.4 provided a mechanism for declaring canonical import paths using [`// import` comments](https://go.dev/cmd/go/#hdr-Import_path_checking), but package authors did not always provide them. As a result, code written prior to modules may have used a non-canonical import path for a module without surfacing an error for the mismatch. When using modules, the import path must match the canonical module path, so you may need to update `import` statements: for example, you may need to change `import "github.com/golang/lint"` to `import "golang.org/x/lint"`.

Go 1.4提供了一个使用//导入注释来声明规范导入路径的机制，但包的作者并不总是提供这些注释。因此，在模块之前编写的代码可能使用了一个非标准的模块导入路径，而没有出现不匹配的错误。当使用模块时，导入路径必须与规范的模块路径相匹配，所以您可能需要更新导入语句：例如，您可能需要将导入 "github.com/golang/lint "改为导入 "golang.org/x/lint"。

Another scenario in which a module’s canonical path may differ from its repository path occurs for Go modules at major version 2 or higher. A Go module with a major version above 1 must include a major-version suffix in its module path: for example, version `v2.0.0` must have the suffix `/v2`. However, `import` statements may have referred to the packages within the module *without* that suffix. For example, non-module users of `github.com/russross/blackfriday/v2` at `v2.0.1` may have imported it as `github.com/russross/blackfriday` instead, and will need to update the import path to include the `/v2` suffix.

另一种情况是，一个模块的标准路径可能与它的仓库路径不同，这发生在主版本为2或更高的Go模块。主版本在1以上的Go模块必须在其模块路径中包含一个主版本后缀：例如，版本v2.0.0必须有后缀/v2.然而，导入语句可能已经提到了模块中的包，而没有这个后缀。例如，github.com/russross/blackfriday/v2在v2.0.1的非模块用户可能将其导入为github.com/russross/blackfriday，因此需要更新导入路径以包括/v2后缀。

## Conclusion 结论

Converting to Go modules should be a straightforward process for most users. Occasional issues may arise due to non-canonical import paths or breaking changes within a dependency. Future posts will explore [publishing new versions](https://go.dev/blog/publishing-go-modules), v2 and beyond, and ways to debug strange situations.

对大多数用户来说，转换到Go模块应该是一个简单的过程。偶尔会出现一些问题，由于非正统的导入路径或依赖关系中的破坏性变化。未来的文章将探讨发布新的版本，v2及以上版本，以及调试奇怪情况的方法。

To provide feedback and help shape the future of dependency management in Go, please send us [bug reports](https://go.dev/issue/new) or [experience reports](https://go.dev/wiki/ExperienceReports).

为了提供反馈并帮助塑造Go中依赖项管理的未来，请向我们发送错误报告或经验报告。

Thanks for all your feedback and help improving modules.

谢谢您的反馈和对改进模块的帮助。
