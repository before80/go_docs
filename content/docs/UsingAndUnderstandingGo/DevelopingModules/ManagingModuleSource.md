+++
title = "管理模块来源"
weight = 3
date = 2023-05-17T15:03:14+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Managing module source - 管理模块来源

> 原文：[https://go.dev/doc/modules/managing-source](https://go.dev/doc/modules/managing-source)

When you’re developing modules to publish for others to use, you can help ensure that your modules are easier for other developers to use by following the repository conventions described in this topic.

当你开发模块并发布给其他人使用时，你可以通过遵循本主题中描述的版本库惯例，帮助确保你的模块更容易被其他开发者使用。

This topic describes actions you might take when managing your module repository. For information about the sequence of workflow steps you’d take when revising from version to version, see [Module release and versioning workflow](https://go.dev/doc/modules/release-workflow).

本主题描述了你在管理你的模块库时可能采取的行动。关于你在修订版本时采取的工作流程步骤的顺序，请参阅模块发布和版本工作流程。

Some of the conventions described here are required in modules, while others are best practices. This content assumes you’re familiar with the basic module use practices described in [Managing dependencies](https://go.dev/doc/modules/managing-dependencies).

这里描述的一些惯例在模块中是必须的，而另一些则是最佳实践。本内容假定你已经熟悉了管理依赖关系中描述的基本模块使用惯例。

Go supports the following repositories for publishing modules: Git, Subversion, Mercurial, Bazaar, and Fossil.

Go支持以下发布模块的存储库。Git、Subversion、Mercurial、Bazaar和Fossil。

For an overview of module development, see [Developing and publishing modules](https://go.dev/doc/modules/developing).

关于模块开发的概述，请参阅开发和发布模块。

## How Go tools find your published module Go工具如何找到你发布的模块

In Go’s decentralized system for publishing modules and retrieving their code, you can publish your module while leaving the code in your repository. Go tools rely on naming rules that have repository paths and repository tags indicating a module’s name and version number. When your repository follows these requirements, your module code is downloadable from your repository by Go tools such as the [`go get` command](https://go.dev/ref/mod#go-get).

在 Go 的发布模块和检索其代码的分散系统中，你可以发布你的模块，同时将代码留在你的仓库中。Go工具依赖于命名规则，这些规则有版本库路径和版本库标签，表明模块的名称和版本号。当你的版本库遵循这些要求时，你的模块代码就可以被Go工具（如go get命令）从你的版本库下载。

When a developer uses the `go get` command to get source code for packages their code imports, the command does the following:

当开发者使用go get命令来获取他们的代码所导入的包的源代码时，该命令会做以下工作：

1. From `import` statements in Go source code, `go get` identifies the module path within the package path.从 Go 源代码中的导入语句，go get 识别包路径中的模块路径。
2. Using a URL derived from the module path, the command locates the module source on a module proxy server or at its repository directly.使用从模块路径导出的URL，该命令在模块代理服务器上或直接在其存储库中定位模块源。
3. Locates source for the module version to download by matching the module’s version number to a repository tag to discover the code in the repository. When a version number to use is not yet known, `go get` locates the latest release version.通过将模块的版本号与资源库的标签相匹配来发现资源库中的代码，从而找到要下载的模块版本的源。当使用的版本号还不知道时，go get会找到最新的发布版本。
4. Retrieves module source and downloads it to the developer’s local module cache.检索模块源并下载到开发者的本地模块缓存。

## Organizing code in the repository 在版本库中组织代码

You can keep maintenance simple and improve developers' experience with your module by following the conventions described here. Getting your module code into a repository is generally as simple as with other code.

你可以通过遵循这里描述的惯例来保持维护的简单性，并改善开发者对你的模块的体验。将你的模块代码放入版本库通常和其他代码一样简单。

The following diagram illustrates a source hierarchy for a simple module with two packages.

下图说明了一个有两个包的简单模块的源代码层次结构。

![Diagram illustrating a module source code hierarchy](ManagingModuleSource_img/source-hierarchy.png)

Diagram illustrating a module source code hierarchy

图示模块源代码的层次结构

Your initial commit should include files listed in the following table:

你的初始提交应该包括下表中所列的文件：

| File                                 | Description                                                  |
| ------------------------------------ | ------------------------------------------------------------ |
| LICENSE                              | The module's license. 模块的许可证。                         |
| go.mod                               | Describes the module, including its module path (in effect, its name) and its dependencies. For more, see the [go.mod reference](https://go.dev/doc/modules/gomod-ref).The module path will be given in a module directive, such as:`module example.com/mymodule`For more about choosing a module path, see [Managing dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module).Though you can edit the go.mod file, you'll find it more reliable to make changes through `go` commands.描述模块，包括它的模块路径（实际上是它的名字）和它的依赖关系。更多信息请参见 go.mod 参考。模块路径将在一个模块指令中给出，例如模块 example.com/mymodule。关于选择模块路径的更多信息，请参见管理依赖关系。尽管你可以编辑go.mod文件，但你会发现通过go命令进行修改更为可靠。 |
| go.sum                               | Contains cryptographic hashes that represent the module's dependencies. Go tools use these hashes to authenticate downloaded modules, attempting to confirm that the downloaded module is authentic. Where this confirmation fails, Go will display a security error.The file will be empty or not present when there are no dependencies. You shouldn't edit this file except by using the `go mod tidy` command, which removes unneeded entries.包含代表模块依赖关系的加密哈希值。Go工具使用这些哈希值来验证下载的模块，试图确认下载的模块是真实的。如果确认失败，Go 会显示一个安全错误。当没有依赖关系时，该文件将为空或不存在。你不应该编辑这个文件，除非使用 go mod tidy 命令，它可以删除不需要的条目。 |
| Package directories and .go sources. | Directories and .go files that comprise the Go packages and sources in the module.组成模块中Go包和源的目录和.go文件。 |

From the command-line, you can create an empty repository, add the files that will be part of your initial commit, and commit with a message. Here’s an example using git:

在命令行中，你可以创建一个空的仓库，添加将成为初始提交的一部分的文件，并提交一个消息。下面是一个使用git的例子：

```shell
$ git init
$ git add --all
$ git commit -m "mycode: initial commit"
$ git push
```

## Choosing repository scope 选择版本库范围

You publish code in a module when the code should be versioned independently from code in other modules.

当你在一个模块中发布代码时，该代码应该独立于其他模块的代码进行版本管理。

Designing your repository so that it hosts a single module at its root directory will help keep maintenance simpler, particularly over time as you publish new minor and patch versions, branch into new major versions, and so on. However, if your needs require it, you can instead maintain a collection of modules in a single repository.

设计你的版本库，使其在根目录下承载单一的模块，将有助于保持更简单的维护，特别是随着时间的推移，你发布新的次要和补丁版本，分支到新的主要版本，等等。然而，如果你的需求需要，你可以在一个版本库中维护一系列的模块。

### Sourcing one module per repository 每个版本库采购一个模块

You can maintain a repository that has a single module’s source in it. In this model, you place your go.mod file at the repository root, with package subdirectories containing Go source beneath.

你可以维护一个只有一个模块源代码的版本库。在这种模式下，你把go.mod文件放在版本库根目录下，下面是包含Go源代码的包子目录。

This is the simplest approach, making your module likely easier to manage over time. It helps you avoid the need to prefix a module version number with a directory path.

这是最简单的方法，使你的模块可能更容易长期管理。它可以帮助你避免在模块的版本号前加上目录路径。

![Diagram illustrating a single module's source in its repository](ManagingModuleSource_img/single-module.png)

Diagram illustrating a single module's source in its repository

图示单个模块在其资源库中的源代码

### Sourcing multiple modules in a single repository 在一个资源库中采购多个模块

You can publish multiple modules from a single repository. For example, you might have code in a single repository that constitutes multiple modules, but want to version those modules separately.

你可以从一个资源库发布多个模块。例如，你可能在一个版本库中拥有构成多个模块的代码，但想分别对这些模块进行版本管理。

Each subdirectory that is a module root directory must have its own go.mod file.

每个作为模块根目录的子目录必须有自己的 go.mod 文件。

Sourcing module code in subdirectories changes the form of the version tag you must use when publishing a module. You must prefix the version number part of the tag with the name of the subdirectory that is the module root. For more about version numbers, see [Module version numbering](https://go.dev/doc/modules/version-numbers).

在子目录中采购模块代码会改变你在发布模块时必须使用的版本标签的形式。你必须在标签的版本号部分前面加上作为模块根目录的子目录的名称。更多关于版本号的信息，请参见模块版本号。

For example, for module `example.com/mymodules/module1` below, you would have the following for version v1.2.3:

例如，对于下面的模块example.com/mymodules/module1，你会有以下版本v1.2.3：

- Module path: `example.com/mymodules/module1` 模块路径: example.com/mymodules/module1

- Version tag: `module1/v1.2.3`  版本标签： module1/v1.2.3

- Package path imported by a user: `example.com/mymodules/module1/package1`

  用户导入的软件包路径: example.com/mymodules/module1/package1

- Module path as given in a user’s require directive: `example.com/mymodules/module1 module1/v1.2.3`

  用户在require指令中给出的模块路径：example.com/mymodules/module1 module1/v1.2.3

![Diagram illustrating two modules in a single repository](ManagingModuleSource_img/multiple-modules.png)

Diagram illustrating two modules in a single repository

图中说明了单个资源库中的两个模块