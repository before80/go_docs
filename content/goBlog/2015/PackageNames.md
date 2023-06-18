+++
title = "包名"
weight = 10
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Package names - 包名

https://go.dev/blog/package-names

Sameer Ajmani
4 February 2015

## Introduction 简介

Go code is organized into packages. Within a package, code can refer to any identifier (name) defined within, while clients of the package may only reference the package’s exported types, functions, constants, and variables. Such references always include the package name as a prefix: `foo.Bar` refers to the exported name `Bar` in the imported package named `foo`.

Go代码被组织成包。在一个包内，代码可以引用其中定义的任何标识符（名称），而包的客户只能引用包中导出的类型、函数、常量和变量。这种引用总是包括作为前缀的包名：foo.Bar指的是名为foo的导入包中的导出名称Bar。

Good package names make code better. A package’s name provides context for its contents, making it easier for clients to understand what the package is for and how to use it. The name also helps package maintainers determine what does and does not belong in the package as it evolves. Well-named packages make it easier to find the code you need.

好的包名能使代码变得更好。包的名称为其内容提供了背景，使客户更容易理解包的用途和如何使用它。名字也有助于包的维护者在包的发展过程中确定什么是属于包的，什么是不属于包的。命名良好的包使您更容易找到您需要的代码。

Effective Go provides [guidelines](https://go.dev/doc/effective_go.html#names) for naming packages, types, functions, and variables. This article expands on that discussion and surveys names found in the standard library. It also discusses bad package names and how to fix them.

Effective Go提供了命名包、类型、函数和变量的指南。本文对该讨论进行了扩展，并对标准库中的名称进行了调查。它还讨论了不好的包名以及如何修复它们。

## Package names 包名

Good package names are short and clear. They are lower case, with no `under_scores` or `mixedCaps`. They are often simple nouns, such as:

好的包名是简短而清晰的。它们是小写的，没有下划线或混合大写。它们通常是简单的名词，例如：

- `time` (provides functionality for measuring and displaying time) time (提供测量和显示时间的功能)
- `list` (implements a doubly linked list) list（实现了一个双链表）
- `http` (provides HTTP client and server implementations) http (提供HTTP客户端和服务器的实现)

The style of names typical of another language might not be idiomatic in a Go program. Here are two examples of names that might be good style in other languages but do not fit well in Go:

另一种语言的典型名称风格在Go程序中可能并不习惯。这里有两个名字的例子，在其他语言中可能是很好的风格，但在Go中却不太适合：

- `computeServiceClient`
- `priority_queue`

A Go package may export several types and functions. For example, a `compute` package could export a `Client` type with methods for using the service as well as functions for partitioning a compute task across several clients.

一个Go包可能会输出几种类型和函数。例如，一个计算包可以导出一个客户端类型，其中包含使用服务的方法，以及将一个计算任务划分到几个客户端的函数。

**Abbreviate judiciously.** Package names may be abbreviated when the abbreviation is familiar to the programmer. Widely-used packages often have compressed names:

谨慎地进行缩写。当程序员熟悉缩写的时候，包的名字可以被缩写。广泛使用的包经常有压缩的名字：

- `strconv` (string conversion) (字符串转换)
- `syscall` (system call)  (系统调用)
- `fmt` (formatted I/O) (格式化I/O)

On the other hand, if abbreviating a package name makes it ambiguous or unclear, don’t do it.

另一方面，如果缩写一个包的名字会使它变得含糊不清或不明确，就不要这样做。

**Don’t steal good names from the user.** Avoid giving a package a name that is commonly used in client code. For example, the buffered I/O package is called `bufio`, not `buf`, since `buf` is a good variable name for a buffer.

不要从用户那里偷取好名字。避免给一个包起一个在客户端代码中常用的名字。例如，缓冲I/O包被称为bufio，而不是buf，因为buf是一个好的缓冲区的变量名。

## Naming package contents 命名包的内容

A package name and its contents' names are coupled, since client code uses them together. When designing a package, take the client’s point of view.

包的名称和其内容的名称是耦合的，因为客户端代码会同时使用它们。当设计一个包时，要从客户的角度出发。

**Avoid repetition.** Since client code uses the package name as a prefix when referring to the package contents, the names for those contents need not repeat the package name. The HTTP server provided by the `http` package is called `Server`, not `HTTPServer`. Client code refers to this type as `http.Server`, so there is no ambiguity.

避免重复。由于客户端代码在引用包的内容时使用包的名字作为前缀，这些内容的名字不需要重复包的名字。由http包提供的HTTP服务器被称为Server，而不是HTTPServer。客户端代码将这种类型称为http.Server，这样就不会有歧义了。

**Simplify function names.** When a function in package pkg returns a value of type `pkg.Pkg` (or `*pkg.Pkg`), the function name can often omit the type name without confusion:

简化函数名称。当包pkg中的一个函数返回一个类型为pkg.Pkg（或*pkg.Pkg）的值时，函数名通常可以省略类型名而不会引起混淆：

```go
start := time.Now()                                  // start is a time.Time
t, err := time.Parse(time.Kitchen, "6:06PM")         // t is a time.Time
ctx = context.WithTimeout(ctx, 10*time.Millisecond)  // ctx is a context.Context
ip, ok := userip.FromContext(ctx)                    // ip is a net.IP
```

A function named `New` in package `pkg` returns a value of type `pkg.Pkg`. This is a standard entry point for client code using that type:

包 pkg 中一个名为 New 的函数返回一个 pkg.Pkg 类型的值。这是使用该类型的客户端代码的一个标准入口点：

```go
 q := list.New()  // q is a *list.List
```

When a function returns a value of type `pkg.T`, where `T` is not `Pkg`, the function name may include `T` to make client code easier to understand. A common situation is a package with multiple New-like functions:

当一个函数返回一个pkg.T类型的值时，其中T不是Pkg，函数名可能包括T，以使客户端代码更容易理解。一个常见的情况是一个包有多个类似New的函数：

```go
d, err := time.ParseDuration("10s")  // d is a time.Duration
elapsed := time.Since(start)         // elapsed is a time.Duration
ticker := time.NewTicker(d)          // ticker is a *time.Ticker
timer := time.NewTimer(d)            // timer is a *time.Timer
```

Types in different packages can have the same name, because from the client’s point of view such names are discriminated by the package name. For example, the standard library includes several types named `Reader`, including `jpeg.Reader`, `bufio.Reader`, and `csv.Reader`. Each package name fits with `Reader` to yield a good type name.

不同包中的类型可以有相同的名字，因为从客户的角度来看，这种名字是由包名来区分的。例如，标准库包括几个名为Reader的类型，包括jpeg.Reader、bufio.Reader和csv.Reader。每个包的名字都与Reader相匹配，产生一个好的类型名称。

If you cannot come up with a package name that’s a meaningful prefix for the package’s contents, the package abstraction boundary may be wrong. Write code that uses your package as a client would, and restructure your packages if the result seems poor. This approach will yield packages that are easier for clients to understand and for the package developers to maintain.

如果您不能想出一个对包的内容有意义的前缀的包名，那么包的抽象边界可能是错误的。编写代码，像客户一样使用您的包，如果结果看起来不好，就重组您的包。这种方法将产生对客户来说更容易理解的包，对包的开发者来说更容易维护。

## Package paths 包的路径

A Go package has both a name and a path. The package name is specified in the package statement of its source files; client code uses it as the prefix for the package’s exported names. Client code uses the package path when importing the package. By convention, the last element of the package path is the package name:

Go软件包有一个名称和一个路径。包的名称是在其源文件的包声明中指定的；客户代码将其作为包的导出名称的前缀。客户端代码在导入包时使用包的路径。根据惯例，包路径的最后一个元素是包名：

```go
import (
    "context"                // package context
    "fmt"                    // package fmt
    "golang.org/x/time/rate" // package rate
    "os/exec"                // package exec
)
```

Build tools map package paths onto directories. The go tool uses the [GOPATH](https://go.dev/doc/code.html#GOPATH) environment variable to find the source files for path `"github.com/user/hello"` in directory `$GOPATH/src/github.com/user/hello`. (This situation should be familiar, of course, but it’s important to be clear about the terminology and structure of packages.)

构建工具将包的路径映射到目录上。go工具使用GOPATH环境变量在$GOPATH/src/github.com/user/hello目录下找到路径 "github.com/user/hello "的源文件。（当然，这种情况应该很熟悉，但必须清楚软件包的术语和结构）。

**Directories.** The standard library uses directories like `crypto`, `container`, `encoding`, and `image` to group packages for related protocols and algorithms. There is no actual relationship among the packages in one of these directories; a directory just provides a way to arrange the files. Any package can import any other package provided the import does not create a cycle.

目录。标准库使用像crypto、container、encoding和image这样的目录来分组相关协议和算法的包。这些目录中的包之间没有实际的关系；目录只是提供了一种安排文件的方式。任何包都可以导入任何其他的包，只要这个导入不产生循环。

Just as types in different packages can have the same name without ambiguity, packages in different directories can have the same name. For example, [runtime/pprof](https://go.dev/pkg/runtime/pprof) provides profiling data in the format expected by the [pprof](https://github.com/google/pprof) profiling tool, while [net/http/pprof](https://go.dev/pkg/net/http/pprof) provides HTTP endpoints to present profiling data in this format. Client code uses the package path to import the package, so there is no confusion. If a source file needs to import both `pprof` packages, it can [rename](https://go.dev/ref/spec#Import_declarations) one or both locally. When renaming an imported package, the local name should follow the same guidelines as package names (lower case, no `under_scores` or `mixedCaps`).

就像不同包中的类型可以有相同的名字而不会产生歧义一样，不同目录中的包也可以有相同的名字。例如，runtime/pprof 以 pprof 剖析工具所期望的格式提供剖析数据，而 net/http/pprof 提供 HTTP 端点，以这种格式呈现剖析数据。客户端代码使用包的路径来导入包，所以不会出现混淆。如果一个源文件需要导入两个pprof包，它可以在本地重命名一个或两个包。当重命名一个导入的包时，本地名称应该遵循与包名称相同的准则（小写，没有under_scores或mixedCaps）。

## Bad package names 不好的包名

Bad package names make code harder to navigate and maintain. Here are some guidelines for recognizing and fixing bad names.

不好的包名会使代码更难浏览和维护。这里有一些识别和修复坏名字的准则。

**Avoid meaningless package names.** Packages named `util`, `common`, or `misc` provide clients with no sense of what the package contains. This makes it harder for clients to use the package and makes it harder for maintainers to keep the package focused. Over time, they accumulate dependencies that can make compilation significantly and unnecessarily slower, especially in large programs. And since such package names are generic, they are more likely to collide with other packages imported by client code, forcing clients to invent names to distinguish them.

避免无意义的包名。命名为util、common或misc的包使客户对该包所包含的内容毫无感觉。这使得客户更难使用该包，也使得维护者更难保持该包的重点。随着时间的推移，它们积累的依赖关系会使编译速度大大降低，特别是在大型程序中，这是不必要的。而且，由于这种包的名字是通用的，它们更有可能与客户代码导入的其他包发生冲突，迫使客户发明名字来区分它们。

**Break up generic packages.** To fix such packages, look for types and functions with common name elements and pull them into their own package. For example, if you have

分解通用包。为了修复这样的包，寻找具有共同名称元素的类型和函数，并将它们拉到自己的包中。例如，如果您有

```go
package util
func NewStringSet(...string) map[string]bool {...}
func SortStringSet(map[string]bool) []string {...}
```

then client code looks like

那么客户端代码看起来像

```go
set := util.NewStringSet("c", "a", "b")
fmt.Println(util.SortStringSet(set))
```

Pull these functions out of `util` into a new package, choosing a name that fits the contents:

把这些函数从util中拉出来放到一个新的包中，选择一个符合内容的名字：

```go
package stringset
func New(...string) map[string]bool {...}
func Sort(map[string]bool) []string {...}
```

then the client code becomes

那么客户端的代码就变成了

```go
set := stringset.New("c", "a", "b")
fmt.Println(stringset.Sort(set))
```

Once you’ve made this change, it’s easier to see how to improve the new package:

一旦您做了这个改变，就更容易看到如何改进新包了：

```go
package stringset
type Set map[string]bool
func New(...string) Set {...}
func (s Set) Sort() []string {...}
```

which yields even simpler client code:

这产生了更简单的客户端代码：

```go
set := stringset.New("c", "a", "b")
fmt.Println(set.Sort())
```

The name of the package is a critical piece of its design. Work to eliminate meaningless package names from your projects.

包的名称是其设计的一个关键部分。努力从您的项目中消除无意义的包名。

**Don’t use a single package for all your APIs.** Many well-intentioned programmers put all the interfaces exposed by their program into a single package named `api`, `types`, or `interfaces`, thinking it makes it easier to find the entry points to their code base. This is a mistake. Such packages suffer from the same problems as those named `util` or `common`, growing without bound, providing no guidance to users, accumulating dependencies, and colliding with other imports. Break them up, perhaps using directories to separate public packages from implementation.

不要为您所有的API使用一个包。许多用心良苦的程序员把他们的程序所暴露的所有接口放到一个名为api、type或interface的包中，认为这样可以更容易地找到他们代码库的入口。这是个错误。这样的包和那些命名为util或common的包有同样的问题，它们无限制地增长，不给用户提供任何指导，积累依赖关系，并与其他导入物发生冲突。把它们分开，也许可以用目录把公共包和实现分开。

**Avoid unnecessary package name collisions.** While packages in different directories may have the same name, packages that are frequently used together should have distinct names. This reduces confusion and the need for local renaming in client code. For the same reason, avoid using the same name as popular standard packages like `io` or `http`.

避免不必要的包名冲突。虽然不同目录下的包可能有相同的名字，但经常一起使用的包应该有不同的名字。这可以减少混乱和在客户端代码中进行局部重命名的需要。出于同样的原因，避免使用与流行的标准包（如io或http）相同的名称。

## Conclusion 总结

Package names are central to good naming in Go programs. Take the time to choose good package names and organize your code well. This helps clients understand and use your packages and helps maintainers to grow them gracefully.

包名是Go程序中良好命名的核心。花点时间选择好的包名，并好好组织您的代码。这有助于客户理解和使用您的包，并帮助维护者优雅地增长它们。

## Further reading 进一步阅读

- [Effective Go 有效的Go](https://go.dev/doc/effective_go.html)
- [How to Write Go Code 如何编写Go代码](https://go.dev/doc/code.html)
- [Organizing Go Code (2012 blog post) 组织Go代码（2012年博文）](https://blog.golang.org/organizing-go-code)
- [Organizing Go Code (2014 Google I/O talk) 组织Go代码（2014年谷歌I/O演讲）](https://go.dev/talks/2014/organizeio.slide)
