+++
title = "go模块：V2版及以后"
weight = 4
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go Modules: v2 and Beyond - go模块：V2版及以后

https://go.dev/blog/v2-go-modules

Jean de Klerk and Tyler Bui-Palsulich
7 November 2019

## Introduction 简介

This post is part 4 in a series.

这篇文章是系列文章中的第四部分。

- Part 1 — [Using Go Modules 使用Go模块](https://go.dev/blog/using-go-modules)
- Part 2 — [Migrating To Go Modules 迁移到Go模块](https://go.dev/blog/migrating-to-go-modules)
- Part 3 — [Publishing Go Modules 发布Go模块](https://go.dev/blog/publishing-go-modules)
- **Part 4 — Go Modules: v2 and Beyond** (this post) Go模块：v2版及以后（本帖）
- Part 5 — [Keeping Your Modules Compatible 保持模块的兼容性](https://go.dev/blog/module-compatibility)

**Note:** For documentation on developing modules, see [Developing and publishing modules](https://go.dev/doc/modules/developing).

注意：关于开发模块的文档，请参见开发和发布模块。

As a successful project matures and new requirements are added, past features and design decisions might stop making sense. Developers may want to integrate lessons they’ve learned by removing deprecated functions, renaming types, or splitting complicated packages into manageable pieces. These kinds of changes require effort by downstream users to migrate their code to the new API, so they should not be made without careful consideration that the benefits outweigh the costs.

随着一个成功项目的成熟和新需求的增加，过去的功能和设计决定可能不再有意义。开发人员可能希望通过删除废弃的功能、重命名类型或将复杂的包拆分成可管理的部分来整合他们所学到的经验。这类改变需要下游用户努力将他们的代码迁移到新的API上，所以在没有仔细考虑收益大于成本的情况下，不应该做这些改变。

For projects that are still experimental — at major version `v0` — occasional breaking changes are expected by users. For projects which are declared stable — at major version `v1` or higher — breaking changes must be done in a new major version. This post explores major version semantics, how to create and publish a new major version, and how to maintain multiple major versions of a module.

对于仍处于实验阶段的项目--主要版本为v0--用户可以期待偶尔的破坏性改变。对于那些被宣布为稳定的项目--主要版本v1或更高--破坏性的改变必须在新的主要版本中完成。这篇文章探讨了主要版本的语义，如何创建和发布一个新的主要版本，以及如何维护一个模块的多个主要版本。

## Major versions and module paths 主版本和模块路径

Modules formalized an important principle in Go, the [**import compatibility rule**](https://research.swtch.com/vgo-import):

模块正式确立了Go中的一个重要原则，即导入兼容性规则：

```
If an old package and a new package have the same import path,
the new package must be backwards compatible with the old package.
```

By definition, a new major version of a package is not backwards compatible with the previous version. This means a new major version of a module must have a different module path than the previous version. Starting with `v2`, the major version must appear at the end of the module path (declared in the `module` statement in the `go.mod` file). For example, when the authors of the module `github.com/googleapis/gax-go` developed `v2`, they used the new module path `github.com/googleapis/gax-go/v2`. Users who wanted to use `v2` had to change their package imports and module requirements to `github.com/googleapis/gax-go/v2`.

根据定义，一个包的新的主要版本是不向后兼容以前的版本的。这意味着一个模块的新的主要版本必须有一个与前一个版本不同的模块路径。从v2开始，主要版本必须出现在模块路径的末尾（在go.mod文件的模块声明中声明）。例如，当模块github.com/googleapis/gax-go的作者开发v2时，他们使用新的模块路径github.com/googleapis/gax-go/v2。想使用v2的用户必须把他们的软件包导入和模块要求改为github.com/googleapis/gax-go/v2。

The need for major version suffixes is one of the ways Go modules differs from most other dependency management systems. Suffixes are needed to solve the [diamond dependency problem](https://research.swtch.com/vgo-import#dependency_story). Before Go modules, [gopkg.in](http://gopkg.in/) allowed package maintainers to follow what we now refer to as the import compatibility rule. With gopkg.in, if you depend on a package that imports `gopkg.in/yaml.v1` and another package that imports `gopkg.in/yaml.v2`, there is no conflict because the two `yaml` packages have different import paths — they use a version suffix, as with Go modules. Since gopkg.in shares the same version suffix methodology as Go modules, the Go command accepts the `.v2` in `gopkg.in/yaml.v2` as a valid major version suffix. This is a special case for compatibility with gopkg.in: modules hosted at other domains need a slash suffix like `/v2`.

对主要版本后缀的需求是Go模块与其他大多数依赖管理系统不同的地方之一。需要后缀来解决钻石依赖项问题。在Go模块之前，gopkg.in允许软件包维护者遵循我们现在所说的导入兼容性规则。在gopkg.in中，如果您依赖一个导入gopkg.in/yaml.v1的包，而另一个导入gopkg.in/yaml.v2的包，就不会有冲突，因为这两个yaml包的导入路径不同--它们使用的是版本后缀，与Go模块一样。由于gopkg.in与Go模块共享相同的版本后缀方法，Go命令接受gopkg.in/yaml.v2中的.v2作为一个有效的主要版本后缀。这是与gopkg.in兼容的特殊情况：托管在其他域的模块需要一个斜线后缀，如/v2。

## Major version strategies 主要版本策略

The recommended strategy is to develop `v2+` modules in a directory named after the major version suffix.

推荐的策略是在一个以主要版本后缀命名的目录中开发v2+模块。

```
github.com/googleapis/gax-go @ master branch
/go.mod    → module github.com/googleapis/gax-go
/v2/go.mod → module github.com/googleapis/gax-go/v2
```

This approach is compatible with tools that aren’t aware of modules: file paths within the repository match the paths expected by `go get` in `GOPATH` mode. This strategy also allows all major versions to be developed together in different directories.

这种方法与那些不了解模块的工具兼容：版本库内的文件路径与GOPATH模式下go get所期望的路径一致。这种策略也允许所有的主要版本在不同的目录下一起开发。

Other strategies may keep major versions on separate branches. However, if `v2+` source code is on the repository’s default branch (usually `master`), tools that are not version-aware — including the `go` command in `GOPATH` mode — may not distinguish between major versions.

其他策略可能会把主要版本放在不同的分支上。然而，如果 v2+ 源代码在版本库的默认分支（通常是主干）上，那些没有版本意识的工具--包括 GOPATH 模式下的 go 命令--可能无法区分主要版本。

The examples in this post will follow the major version subdirectory strategy, since it provides the most compatibility. We recommend that module authors follow this strategy as long as they have users developing in `GOPATH` mode.

这篇文章中的例子将遵循主要版本的子目录策略，因为它提供了最大的兼容性。我们建议模块作者只要有用户在GOPATH模式下开发，就应该遵循这个策略。

## Publishing v2 and beyond 发布v2及以后的版本

This post uses `github.com/googleapis/gax-go` as an example:

本帖以github.com/googleapis/gax-go为例：

```shell linenums="1"
$ pwd
/tmp/gax-go
$ ls
CODE_OF_CONDUCT.md  call_option.go  internal
CONTRIBUTING.md     gax.go          invoke.go
LICENSE             go.mod          tools.go
README.md           go.sum          RELEASING.md
header.go
$ cat go.mod
module github.com/googleapis/gax-go

go 1.9

require (
    github.com/golang/protobuf v1.3.1
    golang.org/x/exp v0.0.0-20190221220918-438050ddec5e
    golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3
    golang.org/x/tools v0.0.0-20190114222345-bf090417da8b
    google.golang.org/grpc v1.19.0
    honnef.co/go/tools v0.0.0-20190102054323-c2f93a96b099
)
$
```

To start development on `v2` of `github.com/googleapis/gax-go`, we’ll create a new `v2/` directory and copy our package into it.

为了在github.com/googleapis/gax-go的v2版上开始开发，我们将创建一个新的v2/目录并将我们的包复制到其中。

```shell linenums="1"
$ mkdir v2
$ cp *.go v2/
building file list ... done
call_option.go
gax.go
header.go
invoke.go
tools.go

sent 10588 bytes  received 130 bytes  21436.00 bytes/sec
total size is 10208  speedup is 0.95
$
```

Now, let’s create a v2 `go.mod` file by copying the current `go.mod` file and adding a `v2/` suffix to the module path:

现在，让我们通过复制当前的go.mod文件并在模块路径中添加v2/后缀来创建一个v2 go.mod文件：

```shell linenums="1"
$ cp go.mod v2/go.mod
$ go mod edit -module github.com/googleapis/gax-go/v2 v2/go.mod
$
```

Note that the `v2` version is treated as a separate module from the `v0 / v1` versions: both may coexist in the same build. So, if your `v2+` module has multiple packages, you should update them to use the new `/v2` import path: otherwise, your `v2+` module will depend on your `v0 / v1` module. For example, to update all `github.com/my/project` references to `github.com/my/project/v2`, you can use `find` and `sed`:

注意，v2版本被视为独立于v0/v1版本的模块：两者可以在同一个构建中共存。因此，如果您的v2+模块有多个软件包，您应该更新它们以使用新的/v2导入路径：否则，您的v2+模块将依赖于您的v0 / v1模块。例如，要将所有github.com/my/project的引用更新为github.com/my/project/v2，您可以使用find和sed：

```shell linenums="1"
$ find . -type f \
    -name '*.go' \
    -exec sed -i -e 's,github.com/my/project,github.com/my/project/v2,g' {} \;
$
```

Now we have a `v2` module, but we want to experiment and make changes before publishing a release. Until we release `v2.0.0` (or any version without a pre-release suffix), we can develop and make breaking changes as we decide on the new API. If we want users to be able to experiment with the new API before we officially make it stable, we can publish a `v2` pre-release version:

现在我们有了一个v2模块，但我们想在发布前进行实验和修改。在我们发布v2.0.0（或任何没有预发布后缀的版本）之前，我们可以在决定新的API时进行开发并进行突破性的修改。如果我们想让用户在我们正式将新的 API 稳定下来之前能够进行实验，我们可以发布一个 v2 预发布版本：

```shell linenums="1"
$ git tag v2.0.0-alpha.1
$ git push origin v2.0.0-alpha.1
$
```

Once we are happy with our `v2` API and are sure we don’t need any other breaking changes, we can tag `v2.0.0`:

一旦我们对我们的 v2 API 感到满意，并且确定我们不需要任何其他突破性的改动，我们就可以对 v2.0.0 进行标记：

```shell linenums="1"
$ git tag v2.0.0
$ git push origin v2.0.0
$
```

At that point, there are now two major versions to maintain. Backwards compatible changes and bug fixes will lead to new minor and patch releases (for example, `v1.1.0`, `v2.0.1`, etc.).

至此，现在有两个主要版本需要维护。向后兼容的变化和错误修复将导致新的次要和补丁版本（例如，v1.1.0，v2.0.1，等等）。

## Conclusion 结论

Major version changes result in development and maintenance overhead and require investment from downstream users to migrate. The larger the project, the larger these overheads tend to be. A major version change should only come after identifying a compelling reason. Once a compelling reason has been identified for a breaking change, we recommend developing multiple major versions in the master branch because it is compatible with a wider variety of existing tools.

主要的版本变更会导致开发和维护的开销，并且需要下游用户的投资来进行迁移。项目越大，这些开销往往就越大。只有在确定了一个令人信服的理由之后，才可以进行重大的版本变更。一旦确定了一个令人信服的理由，我们建议在主分支中开发多个主要版本，因为它可以与更多的现有工具兼容。

Breaking changes to a `v1+` module should always happen in a new, `vN+1` module. When a new module is released, it means additional work for the maintainers and for the users who need to migrate to the new package. Maintainers should therefore validate their APIs before making a stable release, and consider carefully whether breaking changes are really necessary beyond `v1`.

对v1+模块的破坏性修改应该总是发生在一个新的、vN+1模块中。当一个新模块发布时，对于维护者和需要迁移到新包的用户来说，这意味着额外的工作。因此，维护者应该在发布稳定版之前验证他们的API，并仔细考虑是否真的有必要在v1版之后进行突破性的修改。
