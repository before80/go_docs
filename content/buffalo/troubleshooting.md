+++
title = "故障排除"
date = 2024-02-04T21:04:53+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/troubleshooting/]({{< ref "/buffalo/troubleshooting" >}})

# Troubleshooting 故障排除 

## App Crashes with `securecookie: hash key is not set` 应用在 `securecookie: hash key is not set` 处崩溃 

After a recent change in the [github.com/gorilla/sessions](http://www.gorillatoolkit.org/pkg/sessions) Buffalo applications will fail to start with the error `securecookie: hash key is not set`.

​	在 github.com/gorilla/sessions 中最近的一次更改后，Buffalo 应用将无法启动，并出现错误 `securecookie: hash key is not set` 。

To fix this you must set an environment variable named `SESSION_SECRET`.

​	要修复此问题，您必须设置一个名为 `SESSION_SECRET` 的环境变量。

For information see [github.com/gobuffalo/buffalo/issues/1067](https://github.com/gobuffalo/buffalo/issues/1067)

​	有关信息，请参阅 github.com/gobuffalo/buffalo/issues/1067

## Command line is slow 命令行速度慢 

If executing `buffalo --help` or any other command from the terminal takes longer than expected, set `export BUFFALO_PLUGIN_PATH=$GOPATH/bin` in your shell config (e.g. .bash_profile).
如果从终端执行 `buffalo --help` 或任何其他命令所花费的时间比预期长，请在您的 shell 配置（例如 .bash_profile）中设置 `export BUFFALO_PLUGIN_PATH=$GOPATH/bin` 。

## Can’t find `buffalo` binary. 找不到 `buffalo` 二进制文件。 

If you can’t find the `buffalo` binary after a successful installation, this is problably caused because `$GOPATH/bin`, or `%GOPATH\bin` (Windows), isn’t in your `$PATH` variable. When a Go binary is installed it is placed in `$GOPATH/bin`. Adding this to your global `$PATH` will allow you to find **any** Go binary everywhere in your system. See [golang.org/doc/code.html#GOPATH](https://golang.org/doc/code.html#GOPATH) for more details.
如果在成功安装后找不到 `buffalo` 二进制文件，这可能是因为 `$GOPATH/bin` 或 `%GOPATH\bin` （Windows）不在您的 `$PATH` 变量中。安装 Go 二进制文件时，它会被放置在 `$GOPATH/bin` 中。将此添加到您的全局 `$PATH` 中将允许您在系统中的任何位置找到任何 Go 二进制文件。有关更多详细信息，请参阅 golang.org/doc/code.html#GOPATH。

## `buffalo new` fails to generate a complete project. `buffalo new` 无法生成完整的项目。 

This happens because the `buffalo new` command cannot find the templates it needs to generate a new application.

​	出现这种情况是因为 `buffalo new` 命令找不到生成新应用程序所需的模板。

There are a couple of things that could cause this issue.

​	可能导致此问题的原因有几个。

- Using multiple `$GOPATH`s. This can happen when you install Buffalo to one `$GOPATH` and then create a new, temporary, `$GOPATH` and try to create a new application there. Because the templates are in the first, original `$GOPATH`, the installer does not find them, and subsequently generates an incomplete application. To fix this, use just one `$GOPATH`.

  ​	使用多个 `$GOPATH` 。当您将 Buffalo 安装到一个 `$GOPATH` 中，然后创建一个新的临时 `$GOPATH` 并尝试在那里创建一个新应用程序时，可能会发生这种情况。由于模板位于第一个原始 `$GOPATH` 中，因此安装程序找不到它们，随后生成不完整的应用程序。要解决此问题，只需使用一个 `$GOPATH` 。

- Using a single `$GOPATH`. If you aren’t using multiple `$GOPATH`s and are still seeing this issue, it is most likely caused by a bad installation. Run `$ go get -u -v github.com/gobuffalo/buffalo/buffalo` again, and it should, hopefully, repair the installation for you.

  ​	使用单个 `$GOPATH` 。如果您没有使用多个 `$GOPATH` 并且仍然看到此问题，则很可能是由错误的安装引起的。再次运行 `$ go get -u -v github.com/gobuffalo/buffalo/buffalo` ，它应该可以为您修复安装。

The original ticket for this issue can be found at [github.com/gobuffalo/buffalo/issues/629](https://github.com/gobuffalo/buffalo/issues/629).

​	可以在 github.com/gobuffalo/buffalo/issues/629 中找到此问题的原始工单。

## `buffalo new` fails with NPM permissions issues. `buffalo new` 因 NPM 权限问题而失败。

This is caused by incorrectly setup Node/NPM installation.

​	这是由 Node/NPM 安装设置不正确引起的。

See [docs.npmjs.com/resolving-eacces-permissions-errors-when-installing-packages-globally](https://docs.npmjs.com/resolving-eacces-permissions-errors-when-installing-packages-globally) for information on how to fix this issue.

​	请参阅 docs.npmjs.com/resolving-eacces-permissions-errors-when-installing-packages-globally，了解有关如何解决此问题的信息。

## `buffalo dev` auto rebuild doesn’t work with NFS. `buffalo dev` 自动重新构建不适用于 NFS。

This is caused by the `fsnotify` package not supporting NFS.

​	这是因为 `fsnotify` 包不支持 NFS。

See [github.com/gobuffalo/buffalo/issues/620](https://github.com/gobuffalo/buffalo/issues/620) for more details and a workaround.

​	有关更多详细信息和解决方法，请参阅 github.com/gobuffalo/buffalo/issues/620。

## `buffalo new` fails with `import path does not begin with hostname` `buffalo new` 失败，并显示 `import path does not begin with hostname` 

This is caused by a mismatched `$GOPATH` and file system.

​	这是由于 `$GOPATH` 和文件系统不匹配而导致的。

```text
GOPATH: /Users/foobar/Documents/Programming/Go
ACTUAL: /Users/foobar/Documents/programming/go
```

Those are not the same and cause problems with a lot of Go tools. Correct the `$GOPATH` to match the file system and retry.

​	它们不同，并且会给许多 Go 工具带来问题。更正 `$GOPATH` 以匹配文件系统，然后重试。

## `buffalo new` fails looking for `golang.org/x/tools/go/gcimporter` `buffalo new` 失败，正在查找 `golang.org/x/tools/go/gcimporter` 

This is caused by an outdated copy of the `github.com/motemen/gore` package. To fix simply update `gore`:

​	这是由于 `github.com/motemen/gore` 包的副本已过时而导致的。要修复，只需更新 `gore` ：

```text
$ go get -u github.com/motemen/gore
```

For information see https://github.com/gobuffalo/buffalo/issues/108 and https://github.com/motemen/gore/issues/63.

​	有关信息，请参阅 https://github.com/gobuffalo/buffalo/issues/108 和 https://github.com/motemen/gore/issues/63。

## `buffalo dev` fails to start with `Unknown` `buffalo dev` 无法使用 `Unknown` 启动 

When starting `$ buffalo dev`, and you encounter this error:

​	在启动 `$ buffalo dev` 时，您遇到此错误：

```text
There was a problem starting the dev server: Unknown, Please review the troubleshooting docs.
```

This may be due to your system missing NodeJS/NPM, Ensure that Node/NPM is installed and is in your `$PATH`. If Node/NPM are indeed in your `$PATH`, try renaming webpack.config.js.

​	这可能是由于您的系统缺少 NodeJS/NPM，请确保已安装 Node/NPM 且位于您的 `$PATH` 中。如果您的 `$PATH` 中确实有 Node/NPM，请尝试重命名 webpack.config.js.

If you are still having issues after attempting the steps above, please reach out to the community in the #buffalo channel on Gophers Slack.

​	如果您在尝试上述步骤后仍然遇到问题，请在 Gophers Slack 的 #buffalo 频道中联系社区。

## `package context: unrecognized import path "context" (import path does not begin with hostname)`

When trying to install Buffalo `go get` returns this error:

​	在尝试安装 Buffalo 时 `go get` 返回此错误：

```text
package context: unrecognized import path "context" (import path does not begin with hostname)
```

This is due to an outdated version of Go. Buffalo requires Go 1.7 or higher. Please check your installation of Go and ensure you running the latest version.

​	这是由于 Go 版本过旧。Buffalo 需要 Go 1.7 或更高版本。请检查您的 Go 安装并确保您运行的是最新版本。

## Error: `unexpected directory layout:` during `go get` 错误： `unexpected directory layout:` 在 `go get` # 时

Occasionally when running `go get` on Buffalo you will get the following error:

​	偶尔在 Buffalo 上运行 `go get` 时，您会收到以下错误：

```text
unexpected directory layout:
import path: github.com/mattn/go-colorable
dir: /go/src/github.com/fatih/color/vendor/github.com/mattn/go-colorable
expand dir: /go/src/github.com/fatih/color/vendor/github.com/mattn/go-colorable
separator: /
```

This issue has been reported previously the Go team, [github.com/golang/go/issues/17597](https://github.com/golang/go/issues/17597).

​	此问题已向 Go 团队报告过，github.com/golang/go/issues/17597。

The best way to solve this problem is to run `go get` again, and it seems to fix itself.

​	解决此问题的最佳方法是再次运行 `go get` ，它似乎可以自行修复。

## Error: in `application.js` from UglifyJs 错误：在 UglifyJs 中的 `application.js` 中 

If you get this error when running `buffalo build` you need to update your `webpack.config.js` to work with https://github.com/gobuffalo/buffalo/pull/350/files.
如果您在运行 `buffalo build` 时收到此错误，则需要更新您的 `webpack.config.js` 以配合 https://github.com/gobuffalo/buffalo/pull/350/files 工作。

## Error: `Killed 9` when running `buffalo` on Mac OS X with Go 1.8.0 错误： `Killed 9` 在 Mac OS X 上运行 `buffalo` 时出现，Go 版本为 1.8.0 

This is a known issue with Go, github.com/golang/go/issues/19734, not with Buffalo.

​	这是一个 Go 的已知问题，github.com/golang/go/issues/19734，与 Buffalo 无关。

The best solution is to upgrade to Go 1.8.1, or greater, and rebuild your Go binaries.

​	最佳解决方案是升级到 Go 1.8.1 或更高版本，并重新构建您的 Go 二进制文件。

## Mac OS X: `Too many open files in system` error Mac OS X： `Too many open files in system` 错误 

If you get this error when running `buffalo dev` that means you are “watching” too many files, either `.go` files or asset files. To correct this you can [change](http://blog.mact.me/2014/10/22/yosemite-upgrade-changes-open-file-limit) the maximum number of open files on your system.
如果您在运行 `buffalo dev` 时收到此错误，则表示您“监视”的文件过多，可能是 `.go` 文件或资产文件。要更正此问题，您可以更改系统上的最大打开文件数。

## `buffalo new` fails trying to run `goimports` `buffalo new` 尝试运行 `goimports` 时失败 

The full error may appear something like the following, and seems to be the result of outdated go tools. To resolve run `rm -r $GOPATH/src/golang.org/`, then run `go get` again.

​	完整错误可能类似于以下内容，并且似乎是过时 go 工具的结果。要解决此问题，请运行 `rm -r $GOPATH/src/golang.org/` ，然后再次运行 `go get` 。

```bash
$ buffalo new myapp

--> go get -u golang.org/x/tools/cmd/goimports
package golang.org/x/tools/cmd/goimports: golang.org/x/tools is a custom import path for https://go.googlesource.com/tools, but /Users/foo/go/src/golang.org/x/tools is checked out from https://code.google.com/p/go.tools

Error: exit status 1
```

## `buffalo g goth` fails to generate `auth.go` `buffalo g goth` 无法生成 `auth.go` 

You might see errors similar to this when you build:

​	在构建时，您可能会看到类似这样的错误：

```bash
buffalo: 2018/01/19 20:58:47 === Error! ===
buffalo: 2018/01/19 20:58:47 === exit status 2
path/path/models
path/path/actions
# path/path/actions
actions/app.go:17:2: gothic redeclared as imported package name
    previous declaration at actions/app.go:15:2
actions/app.go:66:36: undefined: AuthCallback
actions/app.go:67:11: undefined: SetCurrentUser
actions/app.go:68:11: undefined: Authorize
actions/app.go:69:23: undefined: Authorize
```

This could be because the `goth` plugin isn’t able to find the templates for the different providers. This can happen when the `goth` plugin is available in `$PATH`, but the project isn’t in your current `$GOPATH`.

​	这可能是因为 `goth` 插件无法找到不同提供程序的模板。当 `goth` 插件在 `$PATH` 中可用，但项目不在您当前的 `$GOPATH` 中时，可能会发生这种情况。

To fix it, you can either `go get -u github.com/gobuffalo/buffalo-goth` in your project’s `$GOPATH` or use `dep`.

​	要修复它，您可以在项目的 `$GOPATH` 中 `go get -u github.com/gobuffalo/buffalo-goth` 或使用 `dep` 。
