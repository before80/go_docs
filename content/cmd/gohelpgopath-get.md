+++
title = "go help gopath-get"
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 570
description = ""
isCJKLanguage = true
draft = false

+++

​	

The 'go get' command changes behavior depending on whether the go command is running in module-aware mode or legacy GOPATH mode.

​	'go get'命令的行为会根据go命令是在模块感知模式还是传统的GOPATH模式下运行而发生变化。

This help text, accessible as 'go help gopath-get' even in module-aware mode, describes 'go get' as it operates in legacy GOPATH mode.

​	即使在模块感知模式下，作为'go help gopath-get'可访问的此帮助文本也描述了在传统的GOPATH模式下'go get'的操作方式。

Usage:

​	用法：

```cmd
go get [-d] [-f] [-t] [-u] [-v] [-fix] [build flags] [packages]
```

Get downloads the packages named by the import paths, along with their dependencies. It then installs the named packages, like 'go install'.

​	Get下载由导入路径指定的包及其依赖项。然后，它安装了指定的包，就像'go install'一样。

The -d flag instructs get to stop after downloading the packages; that is, it instructs get not to install the packages.

​	-d标志指示get在下载包后停止；也就是说，它指示get不要安装包。

The -f flag, valid only when -u is set, forces get -u not to verify that each package has been checked out from the source control repository implied by its import path. This can be useful if the source is a local fork of the original.

​	-f标志仅在设置了-u时有效，它强制get -u不验证每个包是否已从其导入路径暗示的源代码库签出。如果源代码是原始代码的本地分支，则这可能很有用。

The -fix flag instructs get to run the fix tool on the downloaded packages before resolving dependencies or building the code.

​	-fix标志指示get在解析依赖项或构建代码之前在下载的包上运行fix工具。

The -t flag instructs get to also download the packages required to build the tests for the specified packages.

​	-t标志指示get还下载构建指定包的测试所需的包。

The -u flag instructs get to use the network to update the named packages and their dependencies. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages.

​	-u标志指示get使用网络更新指定的包及其依赖项。默认情况下，get使用网络检出缺失的包，但不使用网络查找现有包的更新。

The -v flag enables verbose progress and debug output.

​	-v标志启用详细的进度和调试输出。

Get also accepts build flags to control the installation. See 'go help build'.

​	Get还接受用于控制安装的构建标志。请参见'go help build'。

When checking out a new package, get creates the target directory GOPATH/src/<import-path>. If the GOPATH contains multiple entries, get uses the first one. For more details see: 'go help gopath'.

​	在检出新包时，get会创建目标目录GOPATH/src/<import-path>。如果GOPATH包含多个条目，则使用第一个。有关详细信息，请参阅：'go help gopath'。

When checking out or updating a package, get looks for a branch or tag that matches the locally installed version of Go. The most important rule is that if the local installation is running version "go1", get searches for a branch or tag named "go1". If no such version exists it retrieves the default branch of the package.

​	在检出或更新包时，get会查找与本地安装的Go版本匹配的分支或标记。最重要的规则是，如果本地安装的版本为"go1"，则get会搜索名为"go1"的分支或标记。如果没有这样的版本，则检索该包的默认分支。

When go get checks out or updates a Git repository, it also updates any git submodules referenced by the repository.

​	当go get检出或更新Git存储库时，它还会更新由存储库引用的任何git子模块。

Get never checks out or updates code stored in vendor directories.

​	Get永远不会检出或更新存储在vendor目录中的代码。

For more about build flags, see 'go help build'.

​	有关构建标志的更多信息，请参见'go help build'。

For more about specifying packages, see 'go help packages'.

​	有关指定包的更多信息，请参见'go help packages'。

For more about how 'go get' finds source code to download, see 'go help importpath'.

​	有关'go get'如何查找要下载的源代码的更多信息，请参见'go help importpath'。

This text describes the behavior of get when using GOPATH to manage source code and dependencies. If instead the go command is running in module-aware mode, the details of get's flags and effects change, as does 'go help get'. See 'go help modules' and 'go help module-get'.

​	此文本描述了使用GOPATH管理源代码和依赖项时get的行为。如果相反，go命令正在模块感知模式下运行，则get的标志和效果的详细信息会发生变化，'go help get'也会有所不同。请参见'go help modules'和'go help module-get'。

See also: go build, go install, go clean.

​	另请参见：go build，go install，go clean。
