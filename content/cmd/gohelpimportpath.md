+++
title = "go help importpath"
date = 2023-12-12T14:13:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

​	

An import path (see 'go help packages') denotes a package stored in the local file system. In general, an import path denotes either a standard package (such as "unicode/utf8") or a package found in one of the work spaces (For more details see: 'go help gopath').

​	导入路径（参见'go help packages'）表示存储在本地文件系统中的包。一般来说，导入路径表示标准包（例如"unicode/utf8"）或在工作空间之一中找到的包（有关详细信息，请参见：'go help gopath'）。

## 相对导入路径 Relative import paths

An import path beginning with ./ or ../ is called a relative path. The toolchain supports relative import paths as a shortcut in two ways.

​	以./或../开头的导入路径称为相对路径。工具链支持相对导入路径，作为两种方式的快捷方式。

First, a relative path can be used as a shorthand on the command line. If you are working in the directory containing the code imported as "unicode" and want to run the tests for "unicode/utf8", you can type "go test ./utf8" instead of needing to specify the full path.

​	首先，相对路径可以用作命令行上的快捷方式。如果您正在包含代码导入为"unicode"的目录中工作，并希望运行"unicode/utf8"的测试，则可以输入"go test ./utf8"，而不需要指定完整路径。

Similarly, in the reverse situation, "go test .." will test "unicode" from the "unicode/utf8" directory. Relative patterns are also allowed, like "go test ./..." to test all subdirectories. See 'go help packages' for details on the pattern syntax.

​	同样，在相反的情况下，"go test .."将测试"unicode/utf8"目录中的"unicode"。相对模式也是允许的，例如"go test ./..."以测试所有子目录。有关模式语法的详细信息，请参见'go help packages'。

Second, if you are compiling a Go program not in a work space, you can use a relative path in an import statement in that program to refer to nearby code also not in a work space.

​	其次，如果您正在编译不在工作空间中的Go程序，可以在该程序的导入语句中使用相对路径来引用不在工作空间中的附近代码。

This makes it easy to experiment with small multipackage programs outside of the usual work spaces, but such programs cannot be installed with "go install" (there is no work space in which to install them),so they are rebuilt from scratch each time they are built.

​	这使得在通常的工作空间之外轻松实验小型多包程序，但是这些程序不能使用"go install"安装（因为没有工作空间可以安装它们），因此它们每次构建时都会从头开始构建。

To avoid ambiguity, Go programs cannot use relative import paths within a work space.

​	为避免歧义，Go程序不能在工作空间内使用相对导入路径。

### 远程导入路径 Remote import paths

Certain import paths also describe how to obtain the source code for the package using a revision control system.

​	某些导入路径还描述了如何使用修订控制系统获取包的源代码。

A few common code hosting sites have special syntax:

​	一些常见的代码托管站点具有特殊的语法：

        Bitbucket (Git, Mercurial)
    
                import "bitbucket.org/user/project"
                import "bitbucket.org/user/project/sub/directory"
    
        GitHub (Git)
    
                import "github.com/user/project"
                import "github.com/user/project/sub/directory"
    
        Launchpad (Bazaar)
    
                import "launchpad.net/project"
                import "launchpad.net/project/series"
                import "launchpad.net/project/series/sub/directory"
    
                import "launchpad.net/~user/project/branch"
                import "launchpad.net/~user/project/branch/sub/directory"
    
        IBM DevOps Services (Git)
    
                import "hub.jazz.net/git/user/project"
                import "hub.jazz.net/git/user/project/sub/directory"

For code hosted on other servers, import paths may either be qualified with the version control type, or the go tool can dynamically fetch the import path over https/http and discover where the code resides from a `<meta>` tag in the HTML.

​	对于托管在其他服务器上的代码，导入路径可以限定版本控制类型，或者go工具可以通过https/http动态获取导入路径，并从HTML的`<meta>`标签中发现代码所在的位置。

To declare the code location, an import path of the form

​	为了声明代码位置，形式为

        repository.vcs/path

specifies the given repository, with or without the .vcs suffix, using the named version control system, and then the path inside that repository. The supported version control systems are:

的导入路径指定了给定仓库，带有或不带有.vcs后缀，使用指定的版本控制系统，然后是该仓库内的路径。支持的版本控制系统有：

        Bazaar      .bzr
        Fossil      .fossil
        Git         .git
        Mercurial   .hg
        Subversion  .svn

For example,

​	例如，

```go
import "example.org/user/foo.hg"
```

denotes the root directory of the Mercurial repository at example.org/user/foo or foo.hg, and

表示位于example.org/user/foo或foo.hg的Mercurial存储库的根目录，以及

```go
import "example.org/repo.git/foo/bar"
```

denotes the foo/bar directory of the Git repository at example.org/repo or repo.git.

表示example.org/repo或repo.git的Git存储库中foo/bar目录。

When a version control system supports multiple protocols, each is tried in turn when downloading. For example, a Git download tries https://, then git+ssh://.

​	当版本控制系统支持多个协议时，下载时会依次尝试每个协议。例如，Git下载会尝试https://，然后是git+ssh://。

By default, downloads are restricted to known secure protocols (e.g. https, ssh). To override this setting for Git downloads, the GIT_ALLOW_PROTOCOL environment variable can be set (For more details see: 'go help environment').

​	默认情况下，下载仅限于已知的安全协议（例如https、ssh）。要覆盖此设置以进行Git下载，可以设置GIT_ALLOW_PROTOCOL环境变量（有关详细信息，请参见：'go help environment'）。

If the import path is not a known code hosting site and also lacks a version control qualifier, the go tool attempts to fetch the import over https/http and looks for a `<meta>` tag in the document's HTML `<head>`.

​	如果导入路径不是已知的代码托管站点，并且还缺少版本控制限定符，则go工具将尝试通过https/http获取导入，并在文档的HTML `<head>`中查找`<meta>`标签。

The meta tag has the form:

​	meta标签的形式为：

```html
<meta name="go-import" content="import-prefix vcs repo-root">
```

The import-prefix is the import path corresponding to the repository root. It must be a prefix or an exact match of the package being fetched with "go get". If it's not an exact match, another http request is made at the prefix to verify the <meta> tags match.

​	import-prefix是与存储库根对应的导入路径。它必须是包含在"go get"中的包的前缀或精确匹配。如果它不是精确匹配，则在前缀处进行另一个http请求，以验证<meta>标签是否匹配。

The meta tag should appear as early in the file as possible.

​	meta标签应尽早出现在文件中。

In particular, it should appear before any raw JavaScript or CSS, to avoid confusing the go command's restricted parser.

​	特别是，它应该出现在任何原始JavaScript或CSS之前，以避免混淆go命令的受限解析器。

The vcs is one of "bzr", "fossil", "git", "hg", "svn".

​	vcs是以下之一："bzr"、"fossil"、"git"、"hg"、"svn"。

The repo-root is the root of the version control system containing a scheme and not containing a .vcs qualifier.

​	repo-root是包含方案但不包含.vcs限定符的版本控制系统的根。

For example,

​	例如，

```go
import "example.org/pkg/foo"
```

will result in the following requests:

将导致以下请求：

        https://example.org/pkg/foo?go-get=1 (preferred)
        http://example.org/pkg/foo?go-get=1  (fallback, only with use of correctly set GOINSECURE)

If that page contains the meta tag

​	如果该页面包含以下meta标签

```html
<meta name="go-import" content="example.org git https://code.org/r/p/exproj">
```

the go tool will verify that https://example.org/?go-get=1 contains the same meta tag and then git clone https://code.org/r/p/exproj into GOPATH/src/example.org.

go工具将验证https://example.org/?go-get=1包含相同的meta标签，然后git克隆https://code.org/r/p/exproj到GOPATH/src/example.org。

When using GOPATH, downloaded packages are written to the first directory listed in the GOPATH environment variable. (See 'go help gopath-get' and 'go help gopath'.)

​	在使用GOPATH时，下载的包会写入GOPATH环境变量中列出的第一个目录。 （参见'go help gopath-get'和'go help gopath'。）

When using modules, downloaded packages are stored in the module cache. See https://golang.org/ref/mod#module-cache.

​	在使用模块时，下载的包存储在模块缓存中。 有关模块缓存的详细信息，请参阅https://golang.org/ref/mod#module-cache。

When using modules, an additional variant of the go-import meta tag is recognized and is preferred over those listing version control systems. That variant uses "mod" as the vcs in the content value, as in:

​	在使用模块时，还会识别并优先使用go-import meta标签的另一种变体，而不是那些列出版本控制系统的变体。 该变体在内容值中使用"mod"作为vcs，例如：

```html
<meta name="go-import" content="example.org mod https://code.org/moduleproxy">
```

This tag means to fetch modules with paths beginning with example.org from the module proxy available at the URL https://code.org/moduleproxy. See https://golang.org/ref/mod#goproxy-protocol for details about the proxy protocol.

​	此标签表示从URL https://code.org/moduleproxy获取以example.org开头的模块。 有关代理协议的详细信息，请参见https://golang.org/ref/mod#goproxy-protocol。

### 导入路径检查 Import path checking

When the custom import path feature described above redirects to a known code hosting site, each of the resulting packages has two possible import paths, using the custom domain or the known hosting site.

​	当上述自定义导入路径功能重定向到已知的代码托管站点时，生成的每个包都有两个可能的导入路径，即使用自定义域或已知的托管站点。

A package statement is said to have an "import comment" if it is immediately followed (before the next newline) by a comment of one of these two forms:

​	如果包语句后面（在下一个换行符之前）紧跟一个以下两种形式的注释，则说该包语句具有"导入注释"：

```go
    package math // import "path"
    package math /* import "path" */
```

The go command will refuse to install a package with an import comment unless it is being referred to by that import path. In this way, import comments let package authors make sure the custom import path is used and not a direct path to the underlying code hosting site.

​	go命令将拒绝安装带有导入注释的包，除非正在引用该导入路径。通过这种方式，导入注释让包的作者确保使用自定义导入路径而不是底层代码托管站点的直接路径。

Import path checking is disabled for code found within vendor trees. This makes it possible to copy code into alternate locations in vendor trees without needing to update import comments.

​	对于在vendor树中找到的代码，禁用导入路径检查。这样可以将代码复制到vendor树中的替代位置，而无需更新导入注释。

Import path checking is also disabled when using modules. Import path comments are obsoleted by the go.mod file's module statement.

​	在使用模块时，导入路径检查也被禁用。导入路径注释被go.mod文件的模块语句所淘汰。

See https://golang.org/s/go14customimport for details.

​	有关详细信息，请参阅https://golang.org/s/go14customimport。
