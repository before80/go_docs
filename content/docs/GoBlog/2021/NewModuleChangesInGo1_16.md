+++
title = "go 1.16中的新模块变化"
weight = 96
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# New module changes in Go 1.16 - go 1.16中的新模块变化

[https://go.dev/blog/go116-module-changes](https://go.dev/blog/go116-module-changes)

Jay Conrod
18 February 2021

We hope you’re enjoying Go 1.16! This release has a lot of new features, especially for modules. The [release notes](https://go.dev/doc/go1.16) describe these changes briefly, but let’s explore a few of them in depth.

​	我们希望您喜欢Go 1.16! 这个版本有很多新功能，特别是对模块而言。[发行说明](https://go.dev/doc/go1.16)简要描述了这些变化，但让我们深入探讨一下其中的几个。

## Modules on by default 默认开启模块

The `go` command now builds packages in module-aware mode by default, even when no `go.mod` is present. This is a big step toward using modules in all projects.

`go`命令现在默认以模块感知模式构建包，即使没有`go.mod`存在。这是朝着在所有项目中使用模块迈出的一大步。

It’s still possible to build packages in GOPATH mode by setting the `GO111MODULE` environment variable to `off`. You can also set `GO111MODULE` to `auto` to enable module-aware mode only when a go.mod file is present in the current directory or any parent directory. This was previously the default. Note that you can set `GO111MODULE` and other variables permanently with `go env -w`:

​	通过将`GO111MODULE`环境变量设置为`off`，仍然可以在`GOPATH`模式下构建包。您也可以将`GO111MODULE`设置为`auto`，以便只有在当前目录或任何父目录中存在`go.mod`文件时才启用模块感知模式。这在以前是默认的。注意，您可以用`go env -w`永久地设置`GO111MODULE`和其他变量：

```shell
go env -w GO111MODULE=auto
```

We plan to drop support for GOPATH mode in Go 1.17. In other words, Go 1.17 will ignore `GO111MODULE`. If you have projects that do not build in module-aware mode, now is the time to migrate. If there is a problem preventing you from migrating, please consider filing an [issue](https://github.com/golang/go/issues/new) or an [experience report](https://github.com/golang/go/wiki/ExperienceReports).

​	我们计划在Go 1.17中放弃对`GOPATH`模式的支持。换句话说，Go 1.17将忽略`GO111MODULE`。如果您的项目没有在模块感知模式下构建，现在是迁移的时候了。如果有问题阻碍您迁移，请考虑提交[issue](https://github.com/golang/go/issues/new)或[体验报告](https://github.com/golang/go/wiki/ExperienceReports)。

## No automatic changes to go.mod and go.sum 不自动改变go.mod和go.sum

Previously, when the `go` command found a problem with `go.mod` or `go.sum` like a missing `require` directive or a missing sum, it would attempt to fix the problem automatically. We received a lot of feedback that this behavior was surprising, especially for commands like `go list` that don’t normally have side effects. The automatic fixes weren’t always desirable: if an imported package wasn’t provided by any required module, the `go` command would add a new dependency, possibly triggering upgrades of common dependencies. Even a misspelled import path would result in a (failed) network lookup.

​	以前，当`go`命令发现`go.mod`或`go.sum`有问题时，比如缺少`require`指令或缺少sum，它会试图自动修复这个问题。我们收到了很多反馈，认为这种行为令人惊讶，尤其是对于像`go list`这样通常不会有副作用的命令。自动修复并不总是可取的：如果一个导入的包没有被任何所需的模块提供，`go`命令会添加一个新的依赖项，可能会触发常见依赖项的升级。即使是拼错的导入路径也会导致网络查询（失败）。

In Go 1.16, module-aware commands report an error after discovering a problem in `go.mod` or `go.sum` instead of attempting to fix the problem automatically. In most cases, the error message recommends a command to fix the problem.

​	在Go 1.16中，模块感知命令在发现`go.mod`或`go.sum`中的问题后会报告一个错误，而不是试图自动修复问题。在大多数情况下，错误信息推荐了一条修复问题的命令。

```shell
$ go build
example.go:3:8: no required module provides package golang.org/x/net/html; to add it:
    go get golang.org/x/net/html
$ go get golang.org/x/net/html
$ go build
```

As before, the `go` command may use the `vendor` directory if it’s present (see [Vendoring](../../../References/GoModulesReference/Module-awareCommands#vendoring) for details). Commands like `go get` and `go mod tidy` still modify `go.mod` and `go.sum`, since their main purpose is to manage dependencies.

​	和以前一样，`go`命令可以使用`vendor`目录，如果它存在的话（详见[Vendoring]()）。像`go get`和`go mod tidy`这样的命令仍然会修改`go.mod`和`go.sum`，因为它们的主要目的是为了管理依赖项。

## Installing an executable at a specific version 安装一个特定版本的可执行文件

The `go install` command can now install an executable at a specific version by specifying an `@version` suffix.

​	`go install` 命令现在可以通过指定 `@version` 后缀来安装特定版本的可执行文件。

```shell
go install golang.org/x/tools/gopls@v0.6.5
```

When using this syntax, `go install` installs the command from that exact module version, ignoring any `go.mod` files in the current directory and parent directories. (Without the `@version` suffix, `go install` continues to operate as it always has, building the program using the version requirements and replacements listed in the current module’s `go.mod`.)

​	当使用这个语法时，`go install` 从该模块的确切版本进行安装，忽略当前目录和父目录下的任何 `go.mod` 文件。(如果没有 `@version` 后缀，`go install` 将继续像以前一样操作，使用当前模块的 `go.mod` 中列出的版本需求和替换来构建程序。)

We used to recommend `go get -u program` to install an executable, but this use caused too much confusion with the meaning of `go get` for adding or changing module version requirements in `go.mod`. And to avoid accidentally modifying `go.mod`, people started suggesting more complex commands like:

​	我们曾经建议用 `go get -u program` 来安装可执行程序，但这种用法与 `go get` 在 `go.mod` 中增加或改变模块版本需求的含义产生了太多混淆。而且为了避免意外地修改`go.mod`，人们开始建议使用更复杂的命令，比如：

```shell
cd $HOME; GO111MODULE=on go get program@latest
```

Now we can all use `go install program@latest` instead. See [`go install`](https://go.dev/ref/mod#go-install) for details.

​	现在我们都可以用`go install program@latest`来代替。详见[go install](../../../References/GoModulesReference/Module-awareCommands#go-install)。

In order to eliminate ambiguity about which versions are used, there are several restrictions on what directives may be present in the program’s `go.mod` file when using this install syntax. In particular, `replace` and `exclude` directives are not allowed, at least for now. In the long term, once the new `go install program@version` is working well for enough use cases, we plan to make `go get` stop installing command binaries. See [issue 43684](https://go.dev/issue/43684) for details.

​	为了消除关于使用哪个版本的歧义，在使用这种安装语法时，对程序的`go.mod`文件中可能出现的指令有一些限制。特别是`replace`和`exclude`指令是不允许的，至少目前是这样。从长远来看，一旦新的`go install program@version`在足够多的使用情况下运行良好，我们计划让`go get`停止安装命令二进制文件。详情见[issue 43684](https://go.dev/issue/43684)。

## Module retraction 模块撤回

Have you ever accidentally published a module version before it was ready? Or have you discovered a problem right after a version was published that needed to be fixed quickly? Mistakes in published versions are difficult to correct. To keep module builds deterministic, a version cannot be modified after it is published. Even if you delete or change a version tag, [`proxy.golang.org`](https://proxy.golang.org/) and other proxies probably already have the original cached.

​	您是否曾经不小心在模块准备好之前就发布了模块版本？或者您是否在一个版本发布后马上发现了一个问题，需要迅速修复？发布的版本中的错误是很难纠正的。为了保持模块构建的确定性，一个版本在发布后不能被修改。即使您删除或改变了一个版本标签，[proxy.golang.org](https://proxy.golang.org/)和其他代理可能已经有了原始缓存。

Module authors can now *retract* module versions using the `retract` directive in `go.mod`. A retracted version still exists and can be downloaded (so builds that depend on it won’t break), but the `go` command won’t select it automatically when resolving versions like `@latest`. `go get` and `go list -m -u` will print warnings about existing uses.

​	模块作者现在可以使用 `go.mod` 中的 `retract` 指令撤回模块版本。被撤回的版本仍然存在并可以被下载（因此依赖于它的构建不会被破坏），但 `go` 命令在解析版本时不会自动选择它，如 `@latest`。 `go get` 和 `go list -m -u` 会打印关于现有用途的警告。

For example, suppose the author of a popular library `example.com/lib` releases `v1.0.5`, then discovers a new security issue. They can add a directive to their `go.mod` file like the one below:

​	例如，假设一个流行的库`example.com/lib`的作者发布了`v1.0.5`，然后发现了一个新的安全问题。他们可以在他们的`go.mod`文件中添加一个类似下面的指令：

```
// Remote-triggered crash in package foo. See CVE-2021-01234.
retract v1.0.5
```

Next, the author can tag and push version `v1.0.6`, the new highest version. After this, users that already depend on `v1.0.5` will be notified of the retraction when they check for updates or when they upgrade a dependent package. The notification message may include text from the comment above the `retract` directive.

​	接下来，作者可以标记并推送版本`v1.0.6`，即新的最高版本。在此之后，已经依赖`v1.0.5`版本的用户在检查更新或升级所依赖的包时，就会收到撤回的通知。通知信息可能包括`retract`指令上面的注释文本。

```shell
$ go list -m -u all
example.com/lib v1.0.0 (retracted)
$ go get .
go: warning: example.com/lib@v1.0.5: retracted by module author:
    Remote-triggered crash in package foo. See CVE-2021-01234.
go: to switch to the latest unretracted version, run:
    go get example.com/lib@latest
```

For an interactive, browser-based guide, check out [Retract Module Versions](https://play-with-go.dev/retract-module-versions_go116_en/) on [play-with-go.dev](https://play-with-go.dev/). See the [`retract` directive docs](https://go.dev/ref/mod#go-mod-file-retract) for syntax details.

​	对于一个互动的、基于浏览器的指南，请查看 [play-with-go.dev](https://play-with-go.dev/) 上的 [Retract 模块版本](https://play-with-go.dev/retract-module-versions_go116_en/)。语法细节见[retract指令](../../../References/GoModulesReference/gomodFiles#retract-directive)文档。

## Controlling version control tools with GOVCS 用GOVCS控制版本控制工具

The `go` command can download module source code from a mirror like [proxy.golang.org](https://proxy.golang.org/) or directly from a version control repository using `git`, `hg`, `svn`, `bzr`, or `fossil`. Direct version control access is important, especially for private modules that aren’t available on proxies, but it’s also potentially a security problem: a bug in a version control tool may be exploited by a malicious server to run unintended code.

​	`go`指令可以从[proxy.golang.org](https://proxy.golang.org/)这样的镜像中下载模块源代码，或者直接从使用`git`、`hg`、`svn`、`bzr`或`fossil`的版本控制存储库中下载。直接的版本控制访问是很重要的，特别是对于那些在代理上不可用的私有模块，但这也是一个潜在的安全问题：版本控制工具中的错误可能被恶意的服务器利用来运行非预期的代码。

Go 1.16 introduces a new configuration variable, `GOVCS`, which lets the user specify which modules are allowed to use specific version control tools. `GOVCS` accepts a comma-separated list of `pattern:vcslist` rules. The `pattern` is a [`path.Match`](https://go.dev/pkg/path#Match) pattern matching one or more leading elements of a module path. The special patterns `public` and `private` match public and private modules (`private` is defined as modules matched by patterns in `GOPRIVATE`; `public` is everything else). The `vcslist` is a pipe-separated list of allowed version control commands or the keyword `all` or `off`.

​	Go 1.16 引入了一个新的配置变量，`GOVCS`，让用户可以指定哪些模块可以使用特定的版本控制工具。`GOVCS` 接受一个以逗号分隔的 `pattern:vcslist` 规则列表。模式是一个[path.Match](https://go.dev/pkg/path#Match)模式，匹配模块路径的一个或多个前导元素。特殊模式`public`和`private`匹配公共和私人模块（`private`被定义为由`GOPRIVATE`中的模式匹配的模块；`public`是其他一切）。`vcslist`是一个管道分隔的允许的版本控制命令的列表，或者关键字`all`或`off`。

For example:

​	例如：

```
GOVCS=github.com:git,evil.com:off,*:git|hg
```

With this setting, modules with paths on `github.com` can be downloaded using `git`; paths on `evil.com` cannot be downloaded using any version control command, and all other paths (`*` matches everything) can be downloaded using `git` or `hg`.

​	有了这个设置，路径在`github.com`上的模块可以用`git`下载；路径在`evil.com`上的模块不能用任何版本控制命令下载，而所有其他路径（`*`匹配所有）可以用`git`或`hg`下载。

If `GOVCS` is not set, or if a module does not match any pattern, the `go` command uses this default: `git` and `hg` are allowed for public modules, and all tools are allowed for private modules. The rationale behind allowing only Git and Mercurial is that these two systems have had the most attention to issues of being run as clients of untrusted servers. In contrast, Bazaar, Fossil, and Subversion have primarily been used in trusted, authenticated environments and are not as well scrutinized as attack surfaces. That is, the default setting is:

​	如果没有设置`GOVCS`，或者某个模块不匹配任何模式，`go`命令就会使用这个默认值：公共模块允许使用`git`和`hg`，私有模块允许使用所有工具。只允许`Git`和`Mercurial`的理由是，这两个系统作为不受信任的服务器的客户端运行的问题最受关注。相比之下，`Bazaar`、`Fossil`和`Subversion`主要用于可信的、经过验证的环境中，没有被作为攻击面进行仔细检查。也就是说，默认设置是：

```
GOVCS=public:git|hg,private:all
```

See [Controlling version control tools with `GOVCS`](https://go.dev/ref/mod#vcs-govcs) for more details.

参见[用GOVCS控制版本控制工具](../../../References/GoModulesReference/VersionControlSystems#controlling-version-control-tools-with-govcs-govcs)，了解更多细节。

## What’s next? 下一步是什么？

We hope you find these features useful. We’re already hard at work on the next set of module features for Go 1.17, particularly [lazy module loading](https://github.com/golang/go/issues/36460), which should make the module loading process faster and more stable. As always, if you run into new bugs, please let us know on the [issue tracker](https://github.com/golang/go/issues). Happy coding!

​	我们希望您觉得这些功能很有用。我们已经在为Go 1.17的下一组模块功能努力工作了，特别是懒模块加载，它应该使[延迟模块加载](https://github.com/golang/go/issues/36460)过程更快、更稳定。像往常一样，如果您遇到新的错误，请在[issue tracker](https://github.com/golang/go/issues)上告诉我们。编码愉快!
