+++
title = "go help gopath"
date = 2023-12-12T14:13:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

​	

The Go path is used to resolve import statements. It is implemented by and documented in the go/build package.

​	Go路径用于解析导入语句。它由go/build包实现和记录。

The GOPATH environment variable lists places to look for Go code. On Unix, the value is a colon-separated string. On Windows, the value is a semicolon-separated string. On Plan 9, the value is a list.

​	GOPATH环境变量列出查找Go代码的位置。在Unix上，该值是一个以冒号分隔的字符串。在Windows上，该值是一个以分号分隔的字符串。在Plan 9上，该值是一个列表。

If the environment variable is unset, GOPATH defaults to a subdirectory named "go" in the user's home directory ($HOME/go on Unix, %USERPROFILE%\go on Windows), unless that directory holds a Go distribution. Run "go env GOPATH" to see the current GOPATH.

​	如果环境变量未设置，GOPATH将默认为用户主目录中名为"go"的子目录（在Unix上为$HOME/go，在Windows上为%USERPROFILE%\go），除非该目录包含Go分发版。运行"go env GOPATH"以查看当前的GOPATH。

See https://golang.org/wiki/SettingGOPATH to set a custom GOPATH.

​	请参阅https://golang.org/wiki/SettingGOPATH设置自定义GOPATH。

Each directory listed in GOPATH must have a prescribed structure:

​	GOPATH中列出的每个目录必须具有规定的结构：

The src directory holds source code. The path below src determines the import path or executable name.

​	src目录包含源代码。 src以下的路径确定导入路径或可执行文件名。

The pkg directory holds installed package objects. As in the Go tree, each target operating system and architecture pair has its own subdirectory of pkg (pkg/GOOS_GOARCH).

​	pkg目录包含已安装的包对象。与Go树一样，每个目标操作系统和体系结构对都有自己的pkg子目录（pkg/GOOS_GOARCH）。

If DIR is a directory listed in the GOPATH, a package with source in DIR/src/foo/bar can be imported as "foo/bar" and has its compiled form installed to "DIR/pkg/GOOS_GOARCH/foo/bar.a".

​	如果DIR是GOPATH中列出的目录，那么具有DIR/src/foo/bar中的源代码的包可以作为"foo/bar"导入，并且其编译形式安装到"DIR/pkg/GOOS_GOARCH/foo/bar.a"。

The bin directory holds compiled commands. Each command is named for its source directory, but only the final element, not the entire path. That is, the command with source in DIR/src/foo/quux is installed into DIR/bin/quux, not DIR/bin/foo/quux. The "foo/" prefix is stripped so that you can add DIR/bin to your PATH to get at the installed commands. If the GOBIN environment variable is set, commands are installed to the directory it names instead of DIR/bin. GOBIN must be an absolute path.

​	bin目录包含已编译的命令。每个命令以其源目录命名，但只有最终元素，而不是整个路径。也就是说，具有DIR/src/foo/quux中源代码的命令安装到DIR/bin/quux，而不是DIR/bin/foo/quux。删除"foo/"前缀，以便将DIR/bin添加到PATH以获取已安装的命令。如果设置了GOBIN环境变量，则命令将安装到其指定的目录，而不是DIR/bin。GOBIN必须是绝对路径。

Here's an example directory layout:

​	这是一个示例目录布局：

    GOPATH=/home/user/go
    
    /home/user/go/
        src/
            foo/
                bar/               (go code in package bar)
                    x.go
                quux/              (go code in package main)
                    y.go
        bin/
            quux                   (installed command)
        pkg/
            linux_amd64/
                foo/
                    bar.a          (installed package object)

Go searches each directory listed in GOPATH to find source code, but new packages are always downloaded into the first directory in the list.

​	Go搜索GOPATH中列出的每个目录以查找源代码，但新包始终下载到列表中的第一个目录中。

See https://golang.org/doc/code.html for an example.

​	有关示例，请参见https://golang.org/doc/code.html。

### GOPATH和模块 GOPATH and Modules

When using modules, GOPATH is no longer used for resolving imports. However, it is still used to store downloaded source code (in GOPATH/pkg/mod) and compiled commands (in GOPATH/bin).

在使用模块时，GOPATH不再用于解析导入。但是，它仍然用于存储已下载的源代码（在GOPATH/pkg/mod中）和已编译的命令（在GOPATH/bin中）。

### 内部目录 Internal Directories

Code in or below a directory named "internal" is importable only by code in the directory tree rooted at the parent of "internal". Here's an extended version of the directory layout above:

​	在名为"internal"的目录中或其下的代码仅可由树根为"internal"父级的目录树中的代码导入。以下是上面目录布局的扩展版本：

    /home/user/go/
        src/
            crash/
                bang/              (go code in package bang)
                    b.go
            foo/                   (go code in package foo)
                f.go
                bar/               (go code in package bar)
                    x.go
                internal/
                    baz/           (go code in package baz)
                        z.go
                quux/              (go code in package main)
                    y.go

The code in z.go is imported as "foo/internal/baz", but that import statement can only appear in source files in the subtree rooted at foo. The source files foo/f.go, foo/bar/x.go, and foo/quux/y.go can all import "foo/internal/baz", but the source file crash/bang/b.go cannot.

​	z.go 中的代码被导入为"foo/internal/baz"，但该导入语句只能出现在 foo 子树中的源文件中。源文件 foo/f.go、foo/bar/x.go 和 foo/quux/y.go 都可以导入"foo/internal/baz"，但源文件 crash/bang/b.go 不能。

See https://golang.org/s/go14internal for details.

​	有关详细信息，请参见https://golang.org/s/go14internal。

### vendor目录 Vendor Directories

Go 1.6 includes support for using local copies of external dependencies to satisfy imports of those dependencies, often referred to as vendoring.

​	Go 1.6 包括使用外部依赖项的本地副本来满足这些依赖项的导入，通常称为供应商。

Code below a directory named "vendor" is importable only by code in the directory tree rooted at the parent of "vendor", and only using an import path that omits the prefix up to and including the vendor element.

​	位于名为"vendor"的目录下的代码只能由树根为"vendor"父级的目录树中的代码导入，并且只能使用省略前缀直到并包括供应商元素的导入路径。

Here's the example from the previous section, but with the "internal" directory renamed to "vendor" and a new foo/vendor/crash/bang directory added:

​	下面是上一节的示例，但是将"internal"目录重命名为"vendor"，并添加了一个新的 foo/vendor/crash/bang 目录：

    /home/user/go/
        src/
            crash/
                bang/              (go code in package bang)
                    b.go
            foo/                   (go code in package foo)
                f.go
                bar/               (go code in package bar)
                    x.go
                vendor/
                    crash/
                        bang/      (go code in package bang)
                            b.go
                    baz/           (go code in package baz)
                        z.go
                quux/              (go code in package main)
                    y.go

The same visibility rules apply as for internal, but the code in z.go is imported as "baz", not as "foo/vendor/baz".

​	相同的可见性规则适用于内部，但是 z.go 中的代码作为"baz"导入，而不是作为"foo/vendor/baz"导入。

Code in vendor directories deeper in the source tree shadows code in higher directories. Within the subtree rooted at foo, an import of "crash/bang" resolves to "foo/vendor/crash/bang", not the top-level "crash/bang".

​	位于源代码树中的更深处的vendor 目录的代码会遮蔽更高目录中的代码。在 foo 子树中，"crash/bang" 的导入解析为 "foo/vendor/crash/bang"，而不是顶层的 "crash/bang"。

Code in vendor directories is not subject to import path checking (see 'go help importpath').

​	vendor 目录中的代码不受导入路径检查的影响（请参见'go help importpath'）。

When 'go get' checks out or updates a git repository, it now also updates submodules.

​	当'go get'检出或更新git存储库时，它现在还会更新子模块。

Vendor directories do not affect the placement of new repositories being checked out for the first time by 'go get': those are always placed in the main GOPATH, never in a vendor subtree.

​	vendor 目录不影响首次由'go get'检出的新存储库的放置：它们始终放置在主GOPATH中，而不是vendor子树中。

See https://golang.org/s/go15vendor for details.

​	有关详细信息，请参见https://golang.org/s/go15vendor。
