+++
title = "build"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/build@go1.23.0](https://pkg.go.dev/go/build@go1.23.0)

Package build gathers information about Go packages.

​	`build` 包收集了有关Go包的信息。

## Go Path 

The Go path is a list of directory trees containing Go source code. It is consulted to resolve imports that cannot be found in the standard Go tree. The default path is the value of the GOPATH environment variable, interpreted as a path list appropriate to the operating system (on Unix, the variable is a colon-separated string; on Windows, a semicolon-separated string; on Plan 9, a list).

​	Go路径是一个包含Go源代码的目录树列表。它被用来解决在标准Go树中找不到的导入问题。默认路径是 GOPATH 环境变量的值，被解释为适合操作系统的路径列表(在 Unix 上，该变量是一个以冒号分隔的字符串；在 Windows 上，是一个以分号分隔的字符串；在 Plan 9 上，是一个列表)。

Each directory listed in the Go path must have a prescribed structure:

​	Go路径中列出的每个目录必须有规定的结构：

The src/ directory holds source code. The path below 'src' determines the import path or executable name.

​	src/目录存放源代码。src "下面的路径决定了导入路径或可执行名称。

The pkg/ directory holds installed package objects. As in the Go tree, each target operating system and architecture pair has its own subdirectory of pkg (pkg/GOOS_GOARCH).

​	pkg/目录存放已安装的包对象。与Go树一样，每个目标操作系统和架构对都有自己的pkg子目录(pkg/GOOS_GOARCH)。

If DIR is a directory listed in the Go path, a package with source in DIR/src/foo/bar can be imported as "foo/bar" and has its compiled form installed to "DIR/pkg/GOOS_GOARCH/foo/bar.a" (or, for gccgo, "DIR/pkg/gccgo/foo/libbar.a").

​	如果DIR是Go路径中列出的目录，一个源代码在DIR/src/foo/bar的包可以被导入为 "foo/bar"，并将其编译后的形式安装到 "DIR/pkg/GOOS_GOARCH/foo/bar.a"(或者，对于gccgo，"DIR/pkg/gccgo/foo/libbar.a")。

The bin/ directory holds compiled commands. Each command is named for its source directory, but only using the final element, not the entire path. That is, the command with source in DIR/src/foo/quux is installed into DIR/bin/quux, not DIR/bin/foo/quux. The foo/ is stripped so that you can add DIR/bin to your PATH to get at the installed commands.

​	bin/目录存放已编译的命令。每个命令都以其源目录命名，但只使用最后一个元素，而不是整个路径。也就是说，源码在DIR/src/foo/quux的命令被安装到DIR/bin/quux，而不是DIR/bin/foo/quux。foo/被删除了，这样您就可以在PATH中加入DIR/bin来获取已安装的命令。

Here's an example directory layout:

​	下面是一个目录布局的例子：

```
GOPATH=/home/user/gocode

/home/user/gocode/
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
```

## Build Constraints  构建约束

A build constraint, also known as a build tag, is a condition under which a file should be included in the package. Build constraints are given by a line comment that begins

​	构建约束，也被称为构建标签，是一个文件应该被包含在包中的条件。构建约束是由一行注释给出的，注释的开头是

```
//go:build
```

Build constraints may also be part of a file's name (for example, source_windows.go will only be included if the target operating system is windows).

​	构建约束也可以是文件名称的一部分(例如，source_windows.go只有在目标操作系统是windows时才会被包含)。

See 'go help buildconstraint' (https://golang.org/cmd/go/#hdr-Build_constraints) for details.

## Binary-Only Packages 

In Go 1.12 and earlier, it was possible to distribute packages in binary form without including the source code used for compiling the package. The package was distributed with a source file not excluded by build constraints and containing a "//go:binary-only-package" comment. Like a build constraint, this comment appeared at the top of a file, preceded only by blank lines and other line comments and with a blank line following the comment, to separate it from the package documentation. Unlike build constraints, this comment is only recognized in non-test Go source files.

​	在Go 1.12和更早的版本中，可以以二进制形式发布包，而不包括用于编译包的源代码。包的发布包含一个没有被构建约束排除的源文件，并且包含一个"//go:binary-only-package "注释。和构建约束一样，这个注释出现在文件的顶部，前面只有空行和其他行的注释，并且在注释后面有一个空行，把它和包的文档分开。与构建约束不同，这个注释只在非测试Go源文件中被识别。

The minimal source code for a binary-only package was therefore:

​	因此，一个仅有二进制的包的最小源代码是：

```
//go:binary-only-package

package mypkg
```

The source code could include additional Go code. That code was never compiled but would be processed by tools like godoc and might be useful as end-user documentation.

​	源代码可以包括额外的Go代码。这些代码从未被编译过，但会被godoc等工具处理，并可能作为最终用户的文档而有用。

"go build" and other commands no longer support binary-only-packages. Import and ImportDir will still set the BinaryOnly flag in packages containing these comments for use in tools and error messages.

​	"go build "和其他命令不再支持只用二进制的包。Import和ImportDir仍然会在包含这些注释的包中设置BinaryOnly标志，以便在工具和错误信息中使用。



## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/go/build/build.go;l=1996)

``` go 
var ToolDir = getToolDir()
```

ToolDir is the directory containing build tools.

​	`ToolDir`是包含构建工具的目录。

## 函数

### func ArchChar 

``` go 
func ArchChar(goarch string) (string, error)
```

ArchChar returns "?" and an error. In earlier versions of Go, the returned string was used to derive the compiler and linker tool names, the default object file suffix, and the default linker output name. As of Go 1.5, those strings no longer vary by architecture; they are compile, link, .o, and a.out, respectively.

​	`ArchChar`返回"？"和一个错误。在Go的早期版本中，返回的字符串被用来推导出编译器和链接器的工具名称、默认对象文件后缀和默认链接器输出名称。从 Go 1.5 开始，这些字符串不再因架构而异；它们分别是编译、链接、.o 和 a.out。

### func IsLocalImport 

``` go 
func IsLocalImport(path string) bool
```

IsLocalImport reports whether the import path is a local import path, like ".", "..", "./foo", or "../foo".

​	`IsLocalImport`报告导入路径是否为本地导入路径，如"."、"..."、"./foo "或"./foo"。

## 类型

### type Context 

``` go 
type Context struct {
	GOARCH string // target architecture
	GOOS   string // target operating system  // 目标操作系统
	GOROOT string // Go root
	GOPATH string // Go paths

	// Dir is the caller's working directory, or the empty string to use
	// the current directory of the running process. In module mode, this is used
	// to locate the main module.
	//
	// If Dir is non-empty, directories passed to Import and ImportDir must
	// be absolute.
	// Dir是调用者的工作目录，或者用空字符串表示使用运行进程的当前目录。在模块模式下，这是用来定位主模块的。
	//
	// 如果Dir不是空的，传递给Import和ImportDir的目录必须是绝对的。
	Dir string

	CgoEnabled  bool   // whether cgo files are included // 是否包括cgo文件。
	UseAllFiles bool   // use files regardless of go:build lines, file names // 使用文件而不考虑go:build行和文件名。
	Compiler    string // compiler to assume when computing target paths // 计算目标路径时假定的编译器。

	// The build, tool, and release tags specify build constraints
	// that should be considered satisfied when processing go:build lines.
	// Clients creating a new context may customize BuildTags, which
	// defaults to empty, but it is usually an error to customize ToolTags or ReleaseTags.
	// ToolTags defaults to build tags appropriate to the current Go toolchain configuration.
	// ReleaseTags defaults to the list of Go releases the current release is compatible with.
	// BuildTags is not set for the Default build Context.
	// In addition to the BuildTags, ToolTags, and ReleaseTags, build constraints
	// consider the values of GOARCH and GOOS as satisfied tags.
	// The last element in ReleaseTags is assumed to be the current release.
	// build、tool和release标签指定了在处理go:build行时应考虑满足的构建约束。
	// 创建新上下文的客户可以自定义BuildTags，其默认值为空，但自定义ToolTags或ReleaseTags通常是错误的。
	// ToolTags 默认为适合当前 Go 工具链配置的构建标记。
	// ReleaseTags 默认为当前版本所兼容的 Go 版本的列表。
	// BuildTags 没有被设置为默认构建上下文。
	// 除了 BuildTags、ToolTags 和 ReleaseTags 之外，构建约束还考虑 GOARCH 和 GOOS 的值作为满意的标签。
	// ReleaseTags 中的最后一个元素被认为是当前版本。
	BuildTags   []string
	ToolTags    []string
	ReleaseTags []string

	// The install suffix specifies a suffix to use in the name of the installation
	// directory. By default it is empty, but custom builds that need to keep
	// their outputs separate can set InstallSuffix to do so. For example, when
	// using the race detector, the go command uses InstallSuffix = "race", so
	// that on a Linux/386 system, packages are written to a directory named
	// "linux_386_race" instead of the usual "linux_386".
	// 安装后缀指定了安装目录名称中使用的后缀。默认情况下，它是空的，但需要保持输出独立的自定义构建可以设置InstallSuffix来做到这一点。例如，当使用种族检测器时，go命令使用InstallSuffix = "race"，这样在Linux/386系统上，包被写入一个名为 "linux_386_race "的目录，而不是通常的 "linux_386"。
	InstallSuffix string

	// JoinPath joins the sequence of path fragments into a single path.
	// If JoinPath is nil, Import uses filepath.Join.
	// JoinPath将路径片段的序列连接成一个单一的路径。
	// 如果JoinPath为nil，则导入使用filepath.Join。
	JoinPath func(elem ...string) string

	// SplitPathList splits the path list into a slice of individual paths.
	// If SplitPathList is nil, Import uses filepath.SplitList.
	// SplitPathList将路径列表分割成单个路径的片段。
	// 如果SplitPathList为零，导入时使用filepath.SplitList。
	SplitPathList func(list string) []string

	// IsAbsPath reports whether path is an absolute path.
	// If IsAbsPath is nil, Import uses filepath.IsAbs.
	// IsAbsPath报告路径是否是一个绝对路径。
	// 如果IsAbsPath是nil，Import使用filepath.IsAbs。
	IsAbsPath func(path string) bool

	// IsDir reports whether the path names a directory.
	// If IsDir is nil, Import calls os.Stat and uses the result's IsDir method.
	// IsDir报告路径是否命名了一个目录。
	// 如果IsDir为零，Import调用os.Stat并使用结果的IsDir方法。
	IsDir func(path string) bool

	// HasSubdir reports whether dir is lexically a subdirectory of
	// root, perhaps multiple levels below. It does not try to check
	// whether dir exists.
	// If so, HasSubdir sets rel to a slash-separated path that
	// can be joined to root to produce a path equivalent to dir.
	// If HasSubdir is nil, Import uses an implementation built on
	// filepath.EvalSymlinks.
	// HasSubdir 报告 dir 是否是根目录的词法子目录，也许是下面的多级。它并不试图检查dir是否存在。
	// 如果是的话，HasSubdir将rel设置为一个斜线分隔的路径，该路径可以与root连接，产生一个相当于dir的路径。
	// 如果HasSubdir为零，Import使用建立在filepath.EvalSymlinks上的实现。
	HasSubdir func(root, dir string) (rel string, ok bool)

	// ReadDir returns a slice of fs.FileInfo, sorted by Name,
	// describing the content of the named directory.
	// If ReadDir is nil, Import uses os.ReadDir.
	// ReadDir返回fs.FileInfo的一个片断，按Name排序，描述命名目录的内容。
	// 如果ReadDir为nil，则导入使用os.ReadDir。
	ReadDir func(dir string) ([]fs.FileInfo, error)

	// OpenFile opens a file (not a directory) for reading.
	// If OpenFile is nil, Import uses os.Open.
	// OpenFile 打开一个文件(不是一个目录)供读取。
	// 如果OpenFile是nil，Import使用os.Open。
	OpenFile func(path string) (io.ReadCloser, error)
}
```

A Context specifies the supporting context for a build.

​	`Context`指定了支持构建的上下文。

``` go 
var Default Context = defaultContext()
```

Default is the default Context for builds. It uses the GOARCH, GOOS, GOROOT, and GOPATH environment variables if set, or else the compiled code's GOARCH, GOOS, and GOROOT.

​	`Default`是用于构建的默认上下文。它使用 GOARCH、GOOS、GOROOT 和 GOPATH 环境变量(如果设置了)，否则就使用编译后的代码的 GOARCH、GOOS 和 GOROOT。

#### (*Context) Import 

``` go 
func (ctxt *Context) Import(path string, srcDir string, mode ImportMode) (*Package, error)
```

Import returns details about the Go package named by the import path, interpreting local import paths relative to the srcDir directory. If the path is a local import path naming a package that can be imported using a standard import path, the returned package will set p.ImportPath to that path.

​	`Import`返回由导入路径命名的Go包的详细信息，解释相对于srcDir目录的本地导入路径。如果路径是命名一个可以使用标准导入路径导入的包的本地导入路径，返回的包将把p.ImportPath设置为该路径。

In the directory containing the package, .go, .c, .h, and .s files are considered part of the package except for:

​	在包含包的目录中，.go、.c、.h和.s文件被认为是包的一部分，除了：

- .go files in package documentation 包文件中的.go文件
- files starting with _ or . (likely editor temporary files) 以_或.开头的文件(可能是编辑器的临时文件)
- files with build constraints not satisfied by the context 上下文不满足构建约束的文件

If an error occurs, Import returns a non-nil error and a non-nil *Package containing partial information.

​	如果发生错误，Import会返回一个非零的错误和一个非零的*Package，包含部分信息。

#### (*Context) ImportDir 

``` go 
func (ctxt *Context) ImportDir(dir string, mode ImportMode) (*Package, error)
```

ImportDir is like Import but processes the Go package found in the named directory.

​	`ImportDir`与`Import`类似，但处理在指定目录中发现的Go包。

#### (*Context) MatchFile  <- go1.2

``` go 
func (ctxt *Context) MatchFile(dir, name string) (match bool, err error)
```

MatchFile reports whether the file with the given name in the given directory matches the context and would be included in a Package created by ImportDir of that directory.

​	`MatchFile`报告在给定目录中具有给定名称的文件是否与上下文相匹配，以及是否会被包含在由该目录的ImportDir创建的包中。

MatchFile considers the name of the file and may use ctxt.OpenFile to read some or all of the file's content.

​	`MatchFile`考虑文件的名称，并可能使用ctxt.OpenFile读取部分或全部文件的内容。

#### (*Context) SrcDirs 

``` go 
func (ctxt *Context) SrcDirs() []string
```

SrcDirs returns a list of package source root directories. It draws from the current Go root and Go path but omits directories that do not exist.

​	`SrcDirs` 返回包源码根目录的列表。它从当前的 Go 根目录和 Go 路径中提取，但会省略不存在的目录。

#### type Directive <-go1.21.0

```go
type Directive struct {
	Text string         // full line comment including leading slashes
	Pos  token.Position // position of comment
}
```

A Directive is a Go directive comment (//go:zzz...) found in a source file.

### type ImportMode 

``` go 
type ImportMode uint
```

An ImportMode controls the behavior of the Import method.

​	`ImportMode`控制导入方法的行为。

``` go 
const (
	// If FindOnly is set, Import stops after locating the directory
	// that should contain the sources for a package. It does not
	// read any files in the directory.
	// 如果设置了FindOnly，Import会在定位到应该包含包源的目录后停止。它不会读取该目录中的任何文件。
	FindOnly ImportMode = 1 << iota

	// If AllowBinary is set, Import can be satisfied by a compiled
	// package object without corresponding sources.
	//
	// Deprecated:
	// The supported way to create a compiled-only package is to
	// write source code containing a //go:binary-only-package comment at
	// the top of the file. Such a package will be recognized
	// regardless of this flag setting (because it has source code)
	// and will have BinaryOnly set to true in the returned Package.
	// 如果设置了 AllowBinary，Import 可以被一个没有相应源代码的编译包对象所满足。
	//
	// 已废弃：
	// 支持的创建只编译包的方法是编写源代码，在文件顶部包含一个//go:binary-only-package注释。这样的包将被识别，不管这个标志设置如何(因为它有源代码)，并在返回的Package中把BinaryOnly设置为true。
	AllowBinary

	// If ImportComment is set, parse import comments on package statements.
	// Import returns an error if it finds a comment it cannot understand
	// or finds conflicting comments in multiple source files.
	// See golang.org/s/go14customimport for more information.
	// 如果ImportComment被设置，则解析包语句的导入注释。
	// 如果导入发现无法理解的注释或在多个源文件中发现冲突的注释，则返回错误。
	// 参见 golang.org/s/go14customimport 获取更多信息。
	ImportComment

	// By default, Import searches vendor directories
	// that apply in the given source directory before searching
	// the GOROOT and GOPATH roots.
	// If an Import finds and returns a package using a vendor
	// directory, the resulting ImportPath is the complete path
	// to the package, including the path elements leading up
	// to and including "vendor".
	// For example, if Import("y", "x/subdir", 0) finds
	// "x/vendor/y", the returned package's ImportPath is "x/vendor/y",
	// not plain "y".
	// See golang.org/s/go15vendor for more information.
	//
	// Setting IgnoreVendor ignores vendor directories.
	//
	// In contrast to the package's ImportPath,
	// the returned package's Imports, TestImports, and XTestImports
	// are always the exact import paths from the source files:
	// Import makes no attempt to resolve or check those paths.
	// 默认情况下，Import 在搜索 GOROOT 和 GOPATH 根目录之前会搜索适用于给定源目录的供应商目录。
	// 如果一个 Import 找到并返回一个使用供应商目录的包，产生的 ImportPath 是通往该包的完整路径，包括通往 "vendor "的路径元素。
	// 例如，如果 Import("y", "x/subdir", 0) 找到 "x/vendor/y"，返回的包的 ImportPath 是 "x/vendor/y"，而不是简单的 "y"。
	// 参见 golang.org/s/go15vendor 获取更多信息。
	//
	// 设置 IgnoreVendor 会忽略供应商目录。
	//
	// 与包的ImportPath相反。
	// 返回的包的Imports、TestImports和XTestImports总是来自源文件的精确导入路径。
	// Import不会尝试去解决或检查这些路径。
	IgnoreVendor
)
```

### type MultiplePackageError  <- go1.4

``` go 
type MultiplePackageError struct {
	Dir      string   // directory containing files // 包含文件的目录
	Packages []string // package names found // 找到的包名称
	Files    []string // corresponding files: Files[i] declares package Packages[i] // 相应的文件。Files[i]声明了包Packages[i]。
}
```

MultiplePackageError describes a directory containing multiple buildable Go source files for multiple packages.

​	`MultiplePackageError` 描述了一个包含多个包的多个可构建Go源代码文件的目录。

#### (*MultiplePackageError) Error  <- go1.4

``` go 
func (e *MultiplePackageError) Error() string
```

### type NoGoError 

``` go 
type NoGoError struct {
	Dir string
}
```

NoGoError is the error used by Import to describe a directory containing no buildable Go source files. (It may still contain test files, files hidden by build tags, and so on.)

​	`NoGoError`是Import用来描述一个不包含可构建的Go源文件的目录的错误。(它可能仍然包含测试文件、被构建标签隐藏的文件，等等。)

#### (*NoGoError) Error 

``` go 
func (e *NoGoError) Error() string
```

### type Package 

``` go 
type Package struct {
	Dir           string   // directory containing package sources  // 包含包源的目录
	Name          string   // package name // 包的名称
	ImportComment string   // path in import comment on package statement // 包声明中导入注释的路径
	Doc           string   // documentation synopsis // 文档概要
	ImportPath    string   // import path of package ("" if unknown) // 包的导入路径(""如果未知)。
	Root          string   // root of Go tree where this package lives // 这个包所在的Go树的根部
	SrcRoot       string   // package source root directory ("" if unknown) // 包的来源根目录 ("" 如果未知)
	PkgRoot       string   // package install root directory ("" if unknown)  // 包的安装根目录 (" "如果未知)
	PkgTargetRoot string   // architecture dependent install root directory ("" if unknown)// 与架构相关的安装根目录 (" "如果未知)
	BinDir        string   // command install directory ("" if unknown)  // 命令安装目录 (" "如果未知)
	Goroot        bool     // package found in Go root // 在Go根目录下找到的包
	PkgObj        string   // installed .a file // 安装的.a文件
	AllTags       []string // tags that can influence file selection in this directory // 与架构相关的安装根目录 (" "如果未知)
	ConflictDir   string   // this directory shadows Dir in $GOPATH   // 此目录在$GOPATH中的Dir有阴影
	BinaryOnly    bool     // cannot be rebuilt from source (has //go:binary-only-package comment)// 不能从源文件重建(有//go:binary-only-package的注释)

	// Source files
	GoFiles           []string // .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
	CgoFiles          []string // .go source files that import "C"
	IgnoredGoFiles    []string // .go source files ignored for this build (including ignored _test.go files)
	InvalidGoFiles    []string // .go source files with detected problems (parse error, wrong package name, and so on)
	IgnoredOtherFiles []string // non-.go source files ignored for this build
	CFiles            []string // .c source files
	CXXFiles          []string // .cc, .cpp and .cxx source files
	MFiles            []string // .m (Objective-C) source files
	HFiles            []string // .h, .hh, .hpp and .hxx source files
	FFiles            []string // .f, .F, .for and .f90 Fortran source files
	SFiles            []string // .s source files
	SwigFiles         []string // .swig files
	SwigCXXFiles      []string // .swigcxx files
	SysoFiles         []string // .syso system object files to add to archive

	// Cgo directives
	CgoCFLAGS    []string // Cgo CFLAGS directives
	CgoCPPFLAGS  []string // Cgo CPPFLAGS directives
	CgoCXXFLAGS  []string // Cgo CXXFLAGS directives
	CgoFFLAGS    []string // Cgo FFLAGS directives
	CgoLDFLAGS   []string // Cgo LDFLAGS directives
	CgoPkgConfig []string // Cgo pkg-config directives

	// Test information
	TestGoFiles  []string // _test.go files in package
	XTestGoFiles []string // _test.go files outside package

	// Dependency information
	Imports        []string                    // import paths from GoFiles, CgoFiles
	ImportPos      map[string][]token.Position // line information for Imports
	TestImports    []string                    // import paths from TestGoFiles
	TestImportPos  map[string][]token.Position // line information for TestImports
	XTestImports   []string                    // import paths from XTestGoFiles
	XTestImportPos map[string][]token.Position // line information for XTestImports

	// //go:embed patterns found in Go source files
	// For example, if a source file says
	//	//go:embed a* b.c
	// then the list will contain those two strings as separate entries.
	// (See package embed for more details about //go:embed.)
	EmbedPatterns        []string                    // patterns from GoFiles, CgoFiles
	EmbedPatternPos      map[string][]token.Position // line information for EmbedPatterns
	TestEmbedPatterns    []string                    // patterns from TestGoFiles
	TestEmbedPatternPos  map[string][]token.Position // line information for TestEmbedPatterns
	XTestEmbedPatterns   []string                    // patterns from XTestGoFiles
	XTestEmbedPatternPos map[string][]token.Position // line information for XTestEmbedPatternPos
}
```

A Package describes the Go package found in a directory.

​	一个Package描述了在一个目录中发现的Go包。

#### func Import 

``` go 
func Import(path, srcDir string, mode ImportMode) (*Package, error)
```

Import is shorthand for Default.Import.

​	`Import`是Default.Import的简写。

#### func ImportDir 

``` go 
func ImportDir(dir string, mode ImportMode) (*Package, error)
```

ImportDir is shorthand for Default.ImportDir.

​	`ImportDir`是Default.ImportDir的简写。

#### (*Package) IsCommand 

``` go 
func (p *Package) IsCommand() bool
```

IsCommand reports whether the package is considered a command to be installed (not just a library). Packages named "main" are treated as commands.

​	`IsCommand` 报告包是否被认为是一个要安装的命令(而不仅仅是一个库)。命名为 "main "的包被当作命令来对待。