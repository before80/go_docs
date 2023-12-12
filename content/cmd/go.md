+++
title = "go"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

https://pkg.go.dev/cmd/go@go1.21.5

## Overview 

### go 

Go is a tool for managing Go source code.

​	go 是管理go 源代码的工具。

Usage:

​	用法：

```
go <command> [arguments]
```

The commands are:

​	命令包括：

```
bug         开始一个 bug 报告 start a bug report
build       编译包和依赖项 compile packages and dependencies
clean       删除对象文件和缓存文件 remove object files and cached files
doc         显示包或符号的文档 show documentation for package or symbol
env         打印 Go 环境信息  print Go environment information
fix         更新包以使用新的 API  update packages to use new APIs
fmt         gofmt（重新格式化）包源代码 gofmt (reformat) package sources
generate    通过处理源代码生成 Go 文件 generate Go files by processing source
get         添加依赖项到当前模块并安装它们 add dependencies to current module and install them
install     编译和安装包和依赖项 compile and install packages and dependencies
list        列出包或模块 list packages or modules
mod         模块维护 module maintenance
work        工作区维护 workspace maintenance
run         编译并运行 Go 程序 compile and run Go program
test        测试包 test packages
tool        运行指定的 go 工具 run specified go tool
version     打印 Go 版本 print Go version
vet         报告包中可能存在的错误 report likely mistakes in packages
```

Use "go help <command>" for more information about a command.

​	使用"`go help <command>`"了解有关命令的更多信息。

Additional help topics:

​	其他帮助主题：

```
buildconstraint 构建限制条件 build constraints
buildmode       构建模式 build modes
c               Go 和 C 之间的调用 calling between Go and C
cache 			构建和测试缓存 build and test caching
environment 	环境变量 environment variables
filetype 		文件类型 file types
go.mod 			go.mod 文件 the go.mod file
gopath			GOPATH 环境变量 GOPATH environment variable
gopath-get      传统 GOPATH go get   legacy GOPATH go get
goproxy 		模块代理协议 module proxy protocol
importpath 		导入路径语法 import path syntax
modules 		模块，模块版本等等  modules, module versions, and more
module-get 		支持模块的 go get  module-aware go get
module-auth 	使用 go.sum 进行模块身份验证  module authentication using go.sum
packages	 	包列表和模式 package lists and patterns
private 		配置以下载非公共代码 configuration for downloading non-public code
testflag 		测试标志 testing flags
testfunc 		测试函数 testing functions
vcs 			使用 GOVCS 控制版本控制 controlling version control with GOVCS
```

Use "`go help <topic>`" for more information about that topic.

​	使用"`go help <topic>`"了解该主题的更多信息。

#### go bug -> 开始一个 bug 报告

Usage:

​	用法：

```
go bug
```

Bug opens the default browser and starts a new bug report. The report includes useful system information.

​	bug 打开默认浏览器并启动一个新的 bug 报告。该报告包括有用的系统信息。

#### go build -> 编译包和依赖项  

Usage:

​	用法：

```
go build [-o output] [build flags] [packages]
```

Build compiles the packages named by the import paths, along with their dependencies, but it does not install the results.

​	Build 编译由导入路径命名的包以及它们的依赖项，但不安装结果。

If the arguments to build are a list of .go files from a single directory, build treats them as a list of source files specifying a single package.

​	如果 build 的参数是来自单个目录的一组 .go 文件，则 build 将把它们视为指定单个包的源文件列表。

When compiling packages, build ignores files that end in '_test.go'.

​	编译包时，build 忽略以"_test.go"结尾的文件。

When compiling a single main package, build writes the resulting executable to an output file named after the first source file ('go build ed.go rx.go' writes 'ed' or 'ed.exe') or the source code directory ('go build unix/sam' writes 'sam' or 'sam.exe'). The '.exe' suffix is added when writing a Windows executable.

​	当编译单个主包时，build 命令将生成的可执行文件写入以第一个源文件命名的输出文件（'go build ed.go rx.go' 会写入 'ed' 或 'ed.exe'），或者写入源代码目录（'go build unix/sam' 会写入 'sam' 或 'sam.exe'）。在 Windows 上编写可执行文件时，会添加 '.exe' 后缀。

When compiling multiple packages or a single non-main package, build compiles the packages but discards the resulting object, serving only as a check that the packages can be built.

​	当编译多个包或单个非主包时，build 命令会编译这些包，但会丢弃生成的对象，只用于检查这些包是否能够被构建。

The -o flag forces build to write the resulting executable or object to the named output file or directory, instead of the default behavior described in the last two paragraphs. If the named output is an existing directory or ends with a slash or backslash, then any resulting executables will be written to that directory.

​	-o 标志强制 build 命令将生成的可执行文件或对象写入命名的输出文件或目录，而不是采用最后两段所述的默认行为。如果命名的输出是一个现有目录或以斜线或反斜线结尾，则任何生成的可执行文件都将写入该目录。

The build flags are shared by the build, clean, get, install, list, run, and test commands:

​	构建标志由 build、clean、get、install、list、run 和 test 命令共享：

```
-C dir
	Change to dir before running the command.
		在运行命令之前切换到dir目录。
	Any files named on the command line are interpreted after
	changing directories.	
		任何在命令行上命名的文件在更改目录之后解释。
	If used, this flag must be the first one in the command line.
	
-a
	force rebuilding of packages that are already up-to-date.
		强制重新构建已经是最新版本的包。
	
-n
	print the commands but do not run them.
		打印命令但不运行它们。
	
-p n
	the number of programs, such as build commands or
	test binaries, that can be run in parallel.
		可以并行运行的程序数量，例如构建命令或测试二进制文件。
	The default is GOMAXPROCS, normally the number of CPUs available.
		默认值为GOMAXPROCS，通常为可用CPU的数量。
	
-race
	enable data race detection.
		启用数据竞争检测。
	Supported only on linux/amd64, freebsd/amd64, darwin/amd64, darwin/arm64, windows/amd64,
	linux/ppc64le and linux/arm64 (only for 48-bit VMA).
		仅在linux/amd64、freebsd/amd64、darwin/amd64、darwin/arm64、	
    	windows/amd64、linux/ppc64le和linux/arm64上支持（仅适用于48位VMA）。
    
-msan
	enable interoperation with memory sanitizer.	
		启用与内存污点检测器的互操作。
	Supported only on linux/amd64, linux/arm64, freebsd/amd64
	and only with Clang/LLVM as the host C compiler.
	仅在linux/amd64、linux/arm64、freebsd/amd64上支持，
	并且仅支持使用Clang/LLVM作为主机C编译器。
	PIE build mode will be used on all platforms except linux/amd64.
	PIE构建模式将在除linux/amd64外的所有平台上使用。
	
-asan
	enable interoperation with address sanitizer.
		启用与地址污点检测器的互操作。
	Supported only on linux/arm64, linux/amd64.
		仅在linux/arm64、linux/amd64上支持。
	Supported only on linux/amd64 or linux/arm64 and only with GCC 7 and higher
	or Clang/LLVM 9 and higher.
		仅在linux/amd64或linux/arm64上支持，
		并且仅支持使用GCC 7及更高版本或Clang/LLVM 9及更高版本。
	
-cover
	enable code coverage instrumentation.
		启用代码覆盖率分析（需要设置GOEXPERIMENT=coverageredesign）。
		
-covermode set,count,atomic
	set the mode for coverage analysis.
	The default is "set" unless -race is enabled,
	in which case it is "atomic".
	The values:
	set: bool: does this statement run?
	count: int: how many times does this statement run?
	atomic: int: count, but correct in multithreaded tests;
		significantly more expensive.
	Sets -cover.
    
-coverpkg pattern1，pattern2，pattern3
	For a build that targets package 'main' (e.g. building a Go
	executable), apply coverage analysis to each package matching
	the patterns. 
		针对目标为'main'的构建（例如构建Go可执行文件），
		将覆盖分析应用于与模式匹配的每个包。
	The default is to apply coverage analysis to
	packages in the main Go module. See 'go help packages' for a
	description of package patterns.  Sets -cover.		
		默认情况下，将覆盖分析应用于主Go模块中的包。
		有关包模式的说明，请参见"go help packages"。设置-cover。
-v
	print the names of packages as they are compiled.
		编译包时打印包名称。
	
-work
	print the name of the temporary work directory and
	do not delete it when exiting.
		打印临时工作目录的名称，并在退出时不删除它。
	
-x
	print the commands.
	打印命令。

-asmflags '[pattern=]arg list'
	arguments to pass on each go tool asm invocation.
	在每个go工具asm调用中传递的参数。
	
-buildmode mode
	build mode to use. See 'go help buildmode' for more.
	要使用的构建模式。有关更多信息，请参见"go help buildmode"。
	
-buildvcs
	Whether to stamp binaries with version control information
	("true", "false", or "auto"). 
		是否在二进制文件中打印版本控制信息("true"、"false"或"auto")。
	By default ("auto"), version control
	information is stamped into a binary if the main package, the main module
	containing it, and the current directory are all in the same repository.
		默认情况下（"auto"），如果主包、包含它的主模块和当前目录都在同一个仓库中，
	则将版本控制信息打印到二进制文件中。
	Use -buildvcs=false to always omit version control information, or
	-buildvcs=true to error out if version control information is available but
	cannot be included due to a missing tool or ambiguous directory structure.
		使用-buildvcs=false始终省略版本控制信息，
		或者使用-buildvcs=true，
		如果版本控制信息可用但由于缺少工具或模糊的目录结构无法包含，则出错。
-compiler name
	name of compiler to use, as in runtime.Compiler (gccgo or gc).
		指定要使用的编译器名称，如 runtime.Compiler 中的 gccgo 或 gc。
	
-gccgoflags '[pattern=]arg list'
	arguments to pass on each gccgo compiler/linker invocation.
		每个 gccgo 编译器/链接器调用传递的参数列表。
	
-gcflags '[pattern=]arg list'
	arguments to pass on each go tool compile invocation.
		每个 go 工具编译调用传递的参数列表。
	
-installsuffix suffix
	a suffix to use in the name of the package installation directory,
	in order to keep output separate from default builds.
		用于包安装目录名称的后缀，以使输出与默认构建分开。
	If using the -race flag, the install suffix is automatically set to race
	or, if set explicitly, has _race appended to it. Likewise for the -msan
	and -asan flags. Using a -buildmode option that requires non-default compile
	flags has a similar effect.
		如果使用 -race 标志，
		则自动将安装后缀设置为 race 或在显式设置后附加 _race。
		-msan 和 -asan 标志同理。
		使用需要非默认编译标志的 -buildmode 选项具有类似的效果。
	
-ldflags '[pattern=]arg list'
	arguments to pass on each go tool link invocation.
		每个 go 工具链接调用传递的参数列表。
	
-linkshared
	build code that will be linked against shared libraries previously
	created with -buildmode=shared.
		构建将链接到以 -buildmode=shared 创建的共享库的代码。
	
-mod mode
	module download mode to use: readonly, vendor, or mod.
		要使用的模块下载模式：readonly、vendor 或 mod。
	By default, if a vendor directory is present and the go version in go.mod
	is 1.14 or higher, the go command acts as if -mod=vendor were set.
	Otherwise, the go command acts as if -mod=readonly were set.
	See https://golang.org/ref/mod#build-commands for details.
		默认情况下，如果存在 vendor 目录并且 go.mod 中的 go 版本为 1.14 或更高版本，
		则 go 命令会像设置了 -mod=vendor 一样操作。
		否则，go 命令会像设置了 -mod=readonly 一样操作。
		有关详细信息，请参见 https://golang.org/ref/mod#build-commands。
	
-modcacherw
	leave newly-created directories in the module cache read-write
	instead of making them read-only.
		将新创建的目录保留在模块缓存中，以便进行读写，而不是只读。
	
-modfile file
	in module aware mode, read (and possibly write) an alternate go.mod
	file instead of the one in the module root directory. 
		在模块感知模式下，读取（并可能写入）替代 go.mod 文件，
		而不是在模块根目录中的文件。
	A file named "go.mod" must still be present in order to determine the module root
	directory, but it is not accessed. 
		仍然必须存在名为"go.mod"的文件，以确定模块根目录，
		但不会访问该文件。
	When -modfile is specified, an
	alternate go.sum file is also used: its path is derived from the
	-modfile flag by trimming the ".mod" extension and appending ".sum".
		指定 -modfile 时，还会使用替代 go.sum 文件：
		其路径是通过从 -modfile 标志中删除".mod"扩展名并附加".sum"来派生的。
	
-overlay file
	read a JSON config file that provides an overlay for build operations.
		读取 JSON 配置文件，为构建操作提供覆盖。
	The file is a JSON struct with a single field, named 'Replace', that
	maps each disk file path (a string) to its backing file path, so that
	a build will run as if the disk file path exists with the contents
	given by the backing file paths, or as if the disk file path does not
	exist if its backing file path is empty. 
		文件是一个 JSON 结构，具有一个名为 'Replace' 的字段，
		该字段将每个磁盘文件路径（一个字符串）映射到其支持文件路径，
		以便在运行构建时，就像磁盘文件路径存在并具有由支持文件路径给定的内容一样，
		或者如果其支持文件路径为空，则磁盘文件路径将不存在。
	Support for the -overlay flag
	has some limitations: importantly, cgo files included from outside the
	include path must be in the same directory as the Go package they are
	included from, and overlays will not appear when binaries and tests are
	run through go run and go test respectively.
		-overlay 标志的支持有一些限制：重要的是，
		从外部包含的 cgo 文件必须与它们所包含的 Go 包在同一个目录中，
		而覆盖在通过 go run 和 go test 运行二进制文件和测试时将不会出现。
	
-pgo file
	specify the file path of a profile for profile-guided optimization (PGO).
		指定用于编译时的基于概要文件的优化（PGO）的文件路径。
	When the special name "auto" is specified, for each main package in the
	build, the go command selects a file named "default.pgo" in the package's
	directory if that file exists, and applies it to the (transitive)
	dependencies of the main package (other packages are not affected).
		特殊名称"auto"将允许 go 命令在主包目录中选择名为"default.pgo"的文件（如果存在）。
	Special name "off" turns off PGO. The default is "auto".
		特殊名称"off"将关闭 PGO。
	
-pkgdir dir
	install and load all packages from dir instead of the usual locations.
		安装和从 dir 加载所有包，而不是使用通常的位置。
	For example, when building with a non-standard configuration,
	use -pkgdir to keep generated packages in a separate location.
		例如，在使用非标准配置进行构建时，使用 -pkgdir 将生成的包保留在单独的位置。
	
-tags tag，list
	a comma-separated list of additional build tags to consider satisfied
	during the build. For more information about build tags, see
	'go help buildconstraint'. (Earlier versions of Go used a
	space-separated list, and that form is deprecated but still recognized.)
		一个逗号分隔的构建标记列表，表示要在构建期间考虑的额外构建标记。
		有关构建标记的更多信息，请参见"go help buildconstraint"。
		（Go 的早期版本使用空格分隔的列表，虽然这种形式已被弃用但仍然可以识别。）
	
-trimpath
	remove all file system paths from the resulting executable.
		从生成的可执行文件中删除所有文件系统路径。
	Instead of absolute file system paths, the recorded file names
	will begin either a module path@version (when using modules),
	or a plain import path (when using the standard library, or GOPATH).
		记录的文件名将以path@version（在使用模块时）或普通的导入路径
	（在使用标准库或 GOPATH 时）开头。
	
-toolexec 'cmd args'
	a program to use to invoke toolchain programs like vet and asm.
		用于调用类似 vet 和 asm 的工具链程序的程序。
	For example, instead of running asm, the go command will run
	'cmd args /path/to/asm <arguments for asm>'.
	The TOOLEXEC_IMPORTPATH environment variable will be set,
	matching 'go list -f {{.ImportPath}}' for the package being built.
		例如，可以使用 -toolexec 来运行 asm 而不是直接运行，
		go 命令将运行"cmd args /path/to/asm <arguments for asm>"。
		TOOLEXEC_IMPORTPATH 环境变量将被设置，
		与正在构建的包的"go list -f {{.ImportPath}}"匹配。
```

The -asmflags, -gccgoflags, -gcflags, and -ldflags flags accept a space-separated list of arguments to pass to an underlying tool during the build. To embed spaces in an element in the list, surround it with either single or double quotes. The argument list may be preceded by a package pattern and an equal sign, which restricts the use of that argument list to the building of packages matching that pattern (see 'go help packages' for a description of package patterns). Without a pattern, the argument list applies only to the packages named on the command line. The flags may be repeated with different patterns in order to specify different arguments for different sets of packages. If a package matches patterns given in multiple flags, the latest match on the command line wins. For example, 'go build -gcflags=-S fmt' prints the disassembly only for package fmt, while 'go build -gcflags=all=-S fmt' prints the disassembly for fmt and all its dependencies.

​	`-asmflags`、`-gccgoflags`、`-gcflags` 和 `-ldflags` 标志接受一个以空格分隔的参数列表，用于在构建期间传递给底层工具。要在列表中的元素中嵌入空格，请用单引号或双引号括起来。参数列表可以用包模式和等号开头，限制该参数列表仅适用于构建与该模式匹配的包（有关包模式的描述，请参见"go help packages"）。没有模式时，参数列表仅适用于命令行上指定的包。这些标志可以重复使用不同的模式，以便为不同的包集指定不同的参数。如果一个包匹配在多个标志中给定的模式，命令行上最后匹配的标志将覆盖之前的所有标志。例如，"go build -gcflags=-S fmt"仅为包fmt打印反汇编，而"go build -gcflags=all=-S fmt"为fmt及其所有依赖项打印反汇编。

For more about specifying packages, see 'go help packages'. For more about where packages and binaries are installed, run 'go help gopath'. For more about calling between Go and C/C++, run 'go help c'.

​	有关包的详细信息，请参见"go help packages"。有关安装包和二进制文件的位置，请运行"go help gopath"。有关在 Go 和 C/C++ 之间调用的更多信息，请运行"go help c"。

Note: Build adheres to certain conventions such as those described by 'go help gopath'. Not all projects can follow these conventions, however. Installations that have their own conventions or that use a separate software build system may choose to use lower-level invocations such as 'go tool compile' and 'go tool link' to avoid some of the overheads and design decisions of the build tool.

​	注意：构建遵循某些约定，如"go help gopath"所述。然而，并非所有项目都能遵循这些约定。具有自己的约定或使用单独的软件构建系统的安装可能选择使用较低级别的调用，如"go tool compile"和"go tool link"，以避免构建工具的一些开销和设计决策。

See also: go install, go get, go clean.

​	另请参阅：go install、go get、go clean。

#### go clean -> 删除对象文件和缓存文件

Usage:

​	用法：

```
go clean [clean flags] [build flags] [packages]
```

Clean removes object files from package source directories. The go command builds most objects in a temporary directory, so go clean is mainly concerned with object files left by other tools or by manual invocations of go build.

​	`clean`从包源目录中删除对象文件。go命令在临时目录中构建大多数对象，因此go clean主要涉及其他工具或手动调用go build留下的对象文件。

If a package argument is given or the -i or -r flag is set, clean removes the following files from each of the source directories corresponding to the import paths:

​	如果给出了包参数或设置了-i或-r标志，则clean会从每个对应于导入路径的源目录中删除以下文件：

```
_obj/            旧的对象目录，由Makefile留下 old object directory, left from Makefiles
_test/           旧的测试目录，由Makefile留下 old test directory, left from Makefiles
_testmain.go     旧的 gotest 文件，由 Makefile 留下 old gotest file, left from Makefiles
test.out         旧的测试日志，由 Makefile 留下 old test log, left from Makefiles
build.out        旧的测试日志，由 Makefile 留下 old test log, left from Makefiles
*.[568ao]        由 Makefile 留下的对象文件 object files, left from Makefiles

DIR(.exe)        通过 go build 生成的可执行文件  from go build
DIR.test(.exe)   通过 go test -c 生成的测试可执行文件 from go test -c
MAINFILE(.exe)   通过 go build MAINFILE.go 生成的可执行文件 from go build MAINFILE.go
*.so             由 SWIG 生成的文件 from SWIG
```

In the list, DIR represents the final path element of the directory, and MAINFILE is the base name of any Go source file in the directory that is not included when building the package.

​	在列表中，DIR 表示目录的最终路径元素，而 MAINFILE 是目录中未包含在构建包中的任何 Go 源文件的基本名称。

The -i flag causes clean to remove the corresponding installed archive or binary (what 'go install' would create).

​	`-i` 标志会使 clean 删除相应已安装的归档文件或二进制文件（相当于 'go install'）。

The -n flag causes clean to print the remove commands it would execute, but not run them.

​	`-n` 标志会使 clean 打印出它将执行的删除命令，但不实际运行它们。

The -r flag causes clean to be applied recursively to all the dependencies of the packages named by the import paths.

​	`-r` 标志会使 clean 递归应用于被导入路径命名的包的所有依赖项。

The -x flag causes clean to print remove commands as it executes them.

​	`-x` 标志会使 clean 打印出它执行删除命令的同时。

The -cache flag causes clean to remove the entire go build cache.

​	`-cache` 标志会使 clean 删除整个 go 构建缓存。

The -testcache flag causes clean to expire all test results in the go build cache.

​	`-testcache` 标志会使 clean 过期 go 构建缓存中的所有测试结果。

The -modcache flag causes clean to remove the entire module download cache, including unpacked source code of versioned dependencies.

​	`-modcache` 标志会使 clean 删除整个模块下载缓存，包括版本化依赖项的解压缩源代码。

The -fuzzcache flag causes clean to remove files stored in the Go build cache for fuzz testing. The fuzzing engine caches files that expand code coverage, so removing them may make fuzzing less effective until new inputs are found that provide the same coverage. These files are distinct from those stored in testdata directory; clean does not remove those files.

​	`-fuzzcache` 标志会使 clean 删除用于模糊测试的存储在 Go 构建缓存中的文件。模糊引擎会缓存扩展代码覆盖率的文件，因此删除它们可能会使模糊测试变得不太有效，直到找到提供相同覆盖率的新输入。这些文件与存储在 testdata 目录中的文件不同；clean 不会删除这些文件。

For more about build flags, see 'go help build'.

​	有关构建标志的更多信息，请参阅 'go help build'。

For more about specifying packages, see 'go help packages'.

​	有关指定包的更多信息，请参阅 'go help packages'。

#### go doc -> 显示包或符号的文档

Usage:

​	用法：

```
go doc [doc flags] [package|[package.]symbol[.methodOrField]]
```

Doc prints the documentation comments associated with the item identified by its arguments (a package, const, func, type, var, method, or struct field) followed by a one-line summary of each of the first-level items "under" that item (package-level declarations for a package, methods for a type, etc.).

​	`doc` 打印与其参数所标识的项相关联的文档注释（一个包、常量、函数、类型、变量、方法或结构体字段），后跟每个该项"下面"一级项（一个包级别的声明、类型的方法等）的一行摘要。

Doc accepts zero, one, or two arguments.

​	doc 接受零个、一个或两个参数。

Given no arguments, that is, when run as

​	当不带参数运行时，即

```
go doc
```

it prints the package documentation for the package in the current directory. If the package is a command (package main), the exported symbols of the package are elided from the presentation unless the -cmd flag is provided.

它会打印当前目录中包的包文档。如果包是一个命令（package main），则除非提供了 -cmd 标志，否则该包的导出符号将从演示中省略。

When run with one argument, the argument is treated as a Go-syntax-like representation of the item to be documented. What the argument selects depends on what is installed in GOROOT and GOPATH, as well as the form of the argument, which is schematically one of these:

​	当带有一个参数运行时，参数会被视为要文档化的项的 Go 语法样式表示形式。参数选择的内容取决于 GOROOT 和 GOPATH 中安装的内容，以及参数的形式，其概略如下：

```
go doc <pkg>
go doc <sym>[.<methodOrField>]
go doc [<pkg>.]<sym>[.<methodOrField>]
go doc [<pkg>.][<sym>.]<methodOrField>
```

The first item in this list matched by the argument is the one whose documentation is printed. (See the examples below.) However, if the argument starts with a capital letter it is assumed to identify a symbol or method in the current directory.

​	列表中第一个匹配参数的项目是其文档将被打印的项目。（见下面的示例。）然而，如果参数以大写字母开头，则假定其为当前目录中标识符或方法。

For packages, the order of scanning is determined lexically in breadth-first order. That is, the package presented is the one that matches the search and is nearest the root and lexically first at its level of the hierarchy. The GOROOT tree is always scanned in its entirety before GOPATH.

​	对于包，扫描的顺序是按广度优先的词法顺序确定的。也就是说，呈现的包是与搜索匹配并且在其层次结构的根和词法上最先的包。在扫描GOPATH之前，始终完整扫描GOROOT树。

If there is no package specified or matched, the package in the current directory is selected, so "go doc Foo" shows the documentation for symbol Foo in the current package.

​	如果没有指定或匹配包，则选择当前目录中的包，因此"go doc Foo"会显示当前包中符号Foo的文档。

The package path must be either a qualified path or a proper suffix of a path. The go tool's usual package mechanism does not apply: package path elements like . and ... are not implemented by go doc.

​	包路径必须是一个合格的路径或路径的后缀。go工具的常规包机制不适用于go doc。例如，包路径元素（如"."和"…"）没有被实现。

When run with two arguments, the first is a package path (full path or suffix), and the second is a symbol, or symbol with method or struct field:

​	当使用两个参数运行时，第一个参数是包路径（完整路径或后缀），第二个是符号或具有方法或结构字段的符号：

```
go doc <pkg> <sym>[.<methodOrField>]
```

In all forms, when matching symbols, lower-case letters in the argument match either case but upper-case letters match exactly. This means that there may be multiple matches of a lower-case argument in a package if different symbols have different cases. If this occurs, documentation for all matches is printed.

​	在所有形式中，当匹配符号时，参数中的小写字母匹配任何情况，但大写字母匹配确切的情况。这意味着，如果不同的符号具有不同的情况，则小写参数在包中可能有多个匹配项。如果出现这种情况，则打印所有匹配项的文档。

Examples:

​	示例：

```sh
go doc
	Show documentation for current package.
	显示当前包的文档。
	
go doc Foo
	Show documentation for Foo in the current package.
	(Foo starts with a capital letter so it cannot match
	a package path.)
		显示当前包中Foo的文档。
		(Foo以大写字母开头，因此不能匹配包路径。)
	
go doc encoding/json
	Show documentation for the encoding/json package.
		显示encoding/json包的文档。
	
go doc json
	Shorthand for encoding/json.
		encoding/json的简写。
	
go doc json.Number (or go doc json.number)
	Show documentation and method summary for json.Number.
		显示json.Number的文档和方法摘要。
	
go doc json.Number.Int64 (or go doc json.number.int64)
	Show documentation for json.Number's Int64 method.
		显示json.Number的Int64方法的文档。
	
go doc cmd/doc
	Show package docs for the doc command.
		显示doc命令的包文档。
	
go doc -cmd cmd/doc
	Show package docs and exported symbols within the doc command.
		显示doc命令中的包文档和导出符号。
	
go doc template.new
	Show documentation for html/template's New function.
	(html/template is lexically before text/template)
		显示html/template的New函数的文档。
		(html/template在text/template之前按字典顺序排列)
	
go doc text/template.new # One argument
	Show documentation for text/template's New function.
		显示text/template的New函数的文档。

go doc text/template new # Two arguments
	Show documentation for text/template's New function.
		显示text/template的New函数的文档。

At least in the current tree, these invocations all print the
documentation for json.Decoder's Decode method:
	至少在当前树中，这些调用都打印json.Decoder的Decode方法的文档：

go doc json.Decoder.Decode
go doc json.decoder.decode
go doc json.decode
cd go/src/encoding/json; go doc decode
```

Flags:

​	标志：

```
-all
	Show all the documentation for the package.
		显示包中的所有文档。
	
-c
	Respect case when matching symbols.
		在匹配符号时区分大小写。
	
-cmd
	Treat a command (package main) like a regular package.
	Otherwise package main's exported symbols are hidden
	when showing the package's top-level documentation.
		将一个命令（package main）视为常规包。
		否则，显示包的顶层文档时将隐藏 package main 的导出符号。
	
-short
	One-line representation for each symbol.
	每个符号显示一行的简要表示。
	
-src
	Show the full source code for the symbol. This will
	display the full Go source of its declaration and
	definition, such as a function definition (including
	the body), type declaration or enclosing const
	block. 
		显示符号的完整源代码。
		这将显示其声明和定义的完整 Go 源代码，
		例如函数定义（包括主体）、类型声明或封闭 const 块。
	The output may therefore include unexported
	details.
		因此，输出可能包括未导出的详细信息。
	
-u
	Show documentation for unexported as well as exported
	symbols, methods, and fields.
	显示未导出的符号、方法和字段的文档，以及导出的文档。
```

#### go env -> 打印go环境信息

Usage:

​	用法：

```
go env [-json] [-u] [-w] [var ...]
```

Env prints Go environment information.

​	`env`打印Go环境信息。

By default env prints information as a shell script (on Windows, a batch file). If one or more variable names is given as arguments, env prints the value of each named variable on its own line.

​	默认情况下，env以shell脚本形式打印信息（在Windows上为批处理文件）。如果给出一个或多个变量名作为参数，env会在自己的行上打印每个命名变量的值。

The -json flag prints the environment in JSON format instead of as a shell script.

​	`-json`标志以JSON格式而不是作为shell脚本打印环境。

The -u flag requires one or more arguments and unsets the default setting for the named environment variables, if one has been set with 'go env -w'.

​	`-u`标志需要一个或多个参数，并取消命名环境变量的默认设置，如果已使用"go env -w"设置。

The -w flag requires one or more arguments of the form NAME=VALUE and changes the default settings of the named environment variables to the given values.

​	`-w`标志需要一个或多个名称=值形式的参数，并将命名环境变量的默认设置更改为给定值。

For more about environment variables, see 'go help environment'.

​	有关环境变量的更多信息，请参见"go help environment"。

#### go fix -> 更新包以使用新API 

Usage:

​	用法：

```
go fix [-fix list] [packages]
```

Fix runs the Go fix command on the packages named by the import paths.

​	`fix`在导入路径命名的包上运行Go fix命令。

The -fix flag sets a comma-separated list of fixes to run. The default is all known fixes. (Its value is passed to 'go tool fix -r'.)

​	`-fix`标志设置要运行的逗号分隔的修复程序列表。默认为所有已知修复程序。 （其值传递给"go tool fix -r"。）

For more about fix, see 'go doc cmd/fix'. For more about specifying packages, see 'go help packages'.

​	有关修复程序的更多信息，请参见"go doc cmd/fix"。有关指定包的更多信息，请参见"go help packages"。

To run fix with other options, run 'go tool fix'.

​	要使用其他选项运行fix，请运行"go tool fix"。

See also: go fmt, go vet.

​	另请参见：go fmt，go vet。

#### go fmt -> Gofmt（重新格式化）包源

Usage:

​	用法：

```
go fmt [-n] [-x] [packages]
```

Fmt runs the command 'gofmt -l -w' on the packages named by the import paths. It prints the names of the files that are modified.

​	`fmt`在导入路径命名的包上运行命令'gofmt -l -w'。它打印已修改的文件的名称。

For more about gofmt, see 'go doc cmd/gofmt'. For more about specifying packages, see 'go help packages'.

​	有关`gofmt`的更多信息，请参见"go doc cmd/gofmt"。有关指定包的更多信息，请参见"go help packages"。

The -n flag prints commands that would be executed. The -x flag prints commands as they are executed.

​	`-n`标志打印将被执行的命令。-x标志按它们被执行的方式打印命令。

The -mod flag's value sets which module download mode to use: readonly or vendor. See 'go help modules' for more.

​	`-mod`标志的值设置要使用的模块下载模式：readonly或vendor。有关更多信息，请参见"go help modules"。

To run gofmt with specific options, run gofmt itself.

​	要使用特定选项运行gofmt，请运行gofmt本身。

See also: go fix, go vet.

​	另请参见：go fix，go vet。

#### go generate -> 通过处理源文件生成Go文件

Usage:

​	用法：

```
go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]
```

Generate runs commands described by directives within existing files. Those commands can run any process but the intent is to create or update Go source files.

​	`generate` 运行由现有文件中的指令描述的命令。这些命令可以运行任何进程，但是其意图是创建或更新 Go 源文件。

Go generate is never run automatically by go build, go test, and so on. It must be run explicitly.

​	`go generate` 不会被 `go build`、`go test` 等自动运行。必须显式运行它。

Go generate scans the file for directives, which are lines of the form,

​	`go generate` 扫描指令文件，这些文件的指令是一行文本，格式如下：

```
//go:generate command argument...
```

(note: no leading spaces and no space in "//go") where command is the generator to be run, corresponding to an executable file that can be run locally. It must either be in the shell path (gofmt), a fully qualified path (/usr/you/bin/mytool), or a command alias, described below.

(注意：没有前导空格，也没有 "//go" 中的空格)，其中 command 是要运行的生成器，对应于可以在本地运行的可执行文件。它必须在 shell 路径（gofmt）、完全限定路径（/usr/you/bin/mytool）或命令别名中。

Note that go generate does not parse the file, so lines that look like directives in comments or multiline strings will be treated as directives.

​	请注意，go generate 不会解析文件，因此看起来像指令的行注释或多行字符串将被视为指令。

The arguments to the directive are space-separated tokens or double-quoted strings passed to the generator as individual arguments when it is run.

​	指令的参数是空格分隔的标记或双引号括起来的字符串，它们作为单独的参数传递给生成器在运行时。

Quoted strings use Go syntax and are evaluated before execution; a quoted string appears as a single argument to the generator.

​	引号括起来的字符串使用 Go 语法，并在执行之前进行评估；引号括起来的字符串在生成器中出现为单个参数。

To convey to humans and machine tools that code is generated, generated source should have a line that matches the following regular expression (in Go syntax):

​	为了让人类和机器工具知道代码是由生成器生成的，生成的源代码应该具有与以下正则表达式匹配的行（使用 Go 语法）：

```
^// Code generated .* DO NOT EDIT\.$
```

This line must appear before the first non-comment, non-blank text in the file.

​	该行必须出现在文件中第一个非注释、非空白文本之前。

Go generate sets several variables when it runs the generator:

​	在运行生成器时，go generate 设置了几个变量：

```
$GOARCH
	The execution architecture (arm, amd64, etc.)
		执行的体系结构（arm、amd64 等）。
$GOOS
	The execution operating system (linux, windows, etc.)
		执行的操作系统（linux、windows 等）。
$GOFILE
	The base name of the file.
		文件的基本名称。
$GOLINE
	The line number of the directive in the source file.
		源文件中指令的行号。
$GOPACKAGE
	The name of the package of the file containing the directive.
		包含指令的文件的包的名称。
$GOROOT
	The GOROOT directory for the 'go' command that invoked the
	generator, containing the Go toolchain and standard library.
		调用生成器的 'go' 命令的 GOROOT 目录，其中包含 Go 工具链和标准库。
$DOLLAR
	A dollar sign.
	一个美元符号。
	
$PATH
	The $PATH of the parent process, with $GOROOT/bin
	placed at the beginning. This causes generators
	that execute 'go' commands to use the same 'go'
	as the parent 'go generate' command.
```

Other than variable substitution and quoted-string evaluation, no special processing such as "globbing" is performed on the command line.

​	除了变量替换和引号括起来的字符串评估之外，命令行不执行任何特殊处理，例如 "globbing"。

As a last step before running the command, any invocations of any environment variables with alphanumeric names, such as `$GOFILE` or `$HOME`, are expanded throughout the command line. The syntax for variable expansion is `$NAME` on all operating systems. Due to the order of evaluation, variables are expanded even inside quoted strings. If the variable NAME is not set, $NAME expands to the empty string.

​	作为执行命令前的最后一步，任何具有字母数字名称的环境变量调用（如`$GOFILE`或`$HOME`）都会在整个命令行中展开。变量扩展的语法在所有操作系统上均为`$NAME`。由于计算顺序，即使在引号内，变量也会被展开。如果未设置变量NAME，则`$NAME`会展开为空字符串。

A directive of the form,

​	以下是一个指令示例：

```
//go:generate -command xxx args...
```

specifies, for the remainder of this source file only, that the string xxx represents the command identified by the arguments. This can be used to create aliases or to handle multiword generators. For example,

该指令指定，在此源文件中，字符串xxx表示由参数标识的命令。这可用于创建别名或处理多个单词的生成器。例如：

```
//go:generate -command foo go tool foo
```

specifies that the command "foo" represents the generator "go tool foo".

该指令指定命令"foo"表示生成器"go tool foo"。

Generate processes packages in the order given on the command line, one at a time. If the command line lists .go files from a single directory, they are treated as a single package. Within a package, generate processes the source files in a package in file name order, one at a time. Within a source file, generate runs generators in the order they appear in the file, one at a time. The go generate tool also sets the build tag "generate" so that files may be examined by go generate but ignored during build.

​	在命令行中按照给定的顺序，逐一处理包，如果命令行列出了单个目录中的.go文件，则它们被视为单个包。在一个包内，按文件名顺序逐一处理源文件。在一个源文件内，生成器按它们出现在文件中的顺序逐一运行。go generate工具还设置了构建标记"generate"，因此文件可以通过go generate进行检查，但在构建过程中会被忽略。

For packages with invalid code, generate processes only source files with a valid package clause.

​	对于包含无效代码的包，generate仅处理具有有效包从句的源文件。

If any generator returns an error exit status, "go generate" skips all further processing for that package.

​	如果任何生成器返回错误的退出状态，则"go generate"跳过该包的所有后续处理。

The generator is run in the package's source directory.

​	生成器在包的源目录中运行。

Go generate accepts two specific flags:

​	go generate接受两个特定的标志：

```
-run=""
	if non-empty, specifies a regular expression to select
	directives whose full original source text (excluding
	any trailing spaces and final newline) matches the
	expression.
		如果非空，则指定一个正则表达式，
		以选择原始源文本（不包括任何尾随空格和最后一个换行符）与表达式匹配的指令。

-skip=""
	if non-empty, specifies a regular expression to suppress
	directives whose full original source text (excluding
	any trailing spaces and final newline) matches the
	expression. If a directive matches both the -run and
	the -skip arguments, it is skipped.
		如果非空，则指定一个正则表达式，
		以抑制原始源文本（不包括任何尾随空格和最后一个换行符）与表达式匹配的指令。
		如果一个指令同时与-run和-skip参数匹配，则它将被跳过。

```

It also accepts the standard build flags including -v, -n, and -x. The -v flag prints the names of packages and files as they are processed. The -n flag prints commands that would be executed. The -x flag prints commands as they are executed.

​	它还接受标准构建标志，包括-v、-n和-x。-v标志会在处理过程中打印包和文件的名称。-n标志会打印将要执行的命令。-x标志会打印正在执行的命令。

For more about build flags, see 'go help build'.

​	有关构建标志的更多信息，请参见"go help build"。

For more about specifying packages, see 'go help packages'.

​	有关指定包的更多信息，请参见"go help packages"。

#### go get -> 添加依赖项到当前模块并安装它们 

Usage:

​	用法：

```
go get [-t] [-u] [-v] [build flags] [packages]
```

Get resolves its command-line arguments to packages at specific module versions, updates go.mod to require those versions, and downloads source code into the module cache.

​	`get` 将命令行参数解析为特定模块版本的包，更新 go.mod 以要求这些版本，并将源代码下载到模块缓存中。

To add a dependency for a package or upgrade it to its latest version:

​	要为包添加依赖项或将其升级到最新版本：

```
go get example.com/pkg
```

To upgrade or downgrade a package to a specific version:

​	要升级或降级特定版本的包：

```
go get example.com/pkg@v1.2.3
```

To remove a dependency on a module and downgrade modules that require it:

​	要删除对模块的依赖项并降级需要它的模块：

```
go get example.com/mod@none
```

To upgrade the minimum required Go version to the latest released Go version:

See https://golang.org/ref/mod#go-get for details.

​	有关详情，请参见 [Go模块参考中的go get命令]({{< ref "/docs/References/GoModulesReference/Module-awareCommands#go-get">}})。

In earlier versions of Go, 'go get' was used to build and install packages. Now, 'go get' is dedicated to adjusting dependencies in go.mod. 'go install' may be used to build and install commands instead. When a version is specified, 'go install' runs in module-aware mode and ignores the go.mod file in the current directory. For example:

​	在早期的 Go 版本中，"go get"用于构建和安装包。现在，"go get"专用于调整 go.mod 中的依赖项。"go install"可用于构建和安装命令。当指定版本时，"go install"以模块感知模式运行，并忽略当前目录中的 go.mod 文件。例如：

```
go install example.com/pkg@v1.2.3
go install example.com/pkg@latest
```

See 'go help install' or https://golang.org/ref/mod#go-install for details.

​	有关详情，请参见 'go help install' 或 [Go模块参考中的go install命令]({{< ref "/docs/References/GoModulesReference/Module-awareCommands#go-install" >}})。

'go get' accepts the following flags.

​	'go get' 接受以下标志：

- The -t flag instructs get to consider modules needed to build tests of packages specified on the command line.

- `-t` 标志指示 get 考虑构建命令行中指定的包的测试所需的模块。

- The -u flag instructs get to update modules providing dependencies of packages named on the command line to use newer minor or patch releases when available.

- `-u` 标志指示 get 更新提供命令行中指定包的依赖项的模块以使用更高的次要或补丁版本。

- The -u=patch flag (not -u patch) also instructs get to update dependencies, but changes the default to select patch releases.

- `-u=patch` 标志（不是 -u patch）也指示 get 更新依赖项，但将默认选择修补程序版本。

  When the -t and -u flags are used together, get will update test dependencies as well.

​		当 -t 和 -u 标志一起使用时，get 也会更新测试依赖项。

- The -x flag prints commands as they are executed. This is useful for debugging version control commands when a module is downloaded directly from a repository.
- `-x` 标志打印执行的命令。当直接从存储库下载模块时，这对于调试版本控制命令非常有用。

For more about modules, see https://golang.org/ref/mod.

​	有关模块的更多信息，请参见 https://golang.org/ref/mod。

For more about using 'go get' to update the minimum Go version and suggested Go toolchain, see https://go.dev/doc/toolchain.

For more about specifying packages, see 'go help packages'.

​	有关指定包的更多信息，请参见 'go help packages'。

This text describes the behavior of get using modules to manage source code and dependencies. If instead the go command is running in GOPATH mode, the details of get's flags and effects change, as does 'go help get'. See 'go help gopath-get'.

​	本文描述了使用模块管理源代码和依赖项的 get 的行为。如果相反，go 命令在 GOPATH 模式下运行，则 get 的标志和效果的细节会改变，'go help get' 也会改变。请参阅 'go help gopath-get'。

See also: go build, go install, go clean, go mod.

​	另请参见：go build、go install、go clean、go mod。

#### go install -> 编译和安装包及其依赖项

Usage:

​	用法：

```
go install [build flags] [packages]
```

Install compiles and installs the packages named by the import paths.

​	`install` 编译并安装导入路径指定的包。

Executables are installed in the directory named by the GOBIN environment variable, which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH environment variable is not set. Executables in $GOROOT are installed in $GOROOT/bin or $GOTOOLDIR instead of $GOBIN.

​	可执行文件将被安装到名为GOBIN的目录中。默认情况下，GOBIN环境变量为`$GOPATH/bin`或者`$HOME/go/bin`（如果GOPATH环境变量未设置）。`$GOROOT`中的可执行文件将被安装到`$GOROOT/bin`或`$GOTOOLDIR`中，而不是`$GOBIN`中。

If the arguments have version suffixes (like @latest or @v1.0.0), "go install" builds packages in module-aware mode, ignoring the go.mod file in the current directory or any parent directory, if there is one. This is useful for installing executables without affecting the dependencies of the main module. To eliminate ambiguity about which module versions are used in the build, the arguments must satisfy the following constraints:

​	如果参数有版本后缀（如`@latest`或`@v1.0.0`），则"go install"将在模块感知模式下进行构建，忽略当前目录或任何父目录中的go.mod文件。这对于安装可执行文件而不影响主模块的依赖项非常有用。为了消除构建中使用的模块版本的歧义，参数必须满足以下限制：

- Arguments must be package paths or package patterns (with "..." wildcards). They must not be standard packages (like fmt), meta-patterns (std, cmd, all), or relative or absolute file paths.
- 参数必须是包路径或包模式（带有"…"通配符）。它们不能是标准包（如fmt），元模式（std，cmd，all）或相对或绝对文件路径。
- All arguments must have the same version suffix. Different queries are not allowed, even if they refer to the same version.
- 所有参数必须具有相同的版本后缀。不允许不同的查询，即使它们引用同一个版本。
- All arguments must refer to packages in the same module at the same version.
- 所有参数必须引用同一模块中的相同版本的包。
- Package path arguments must refer to main packages. Pattern arguments will only match main packages.
- 包路径参数必须引用主包。模式参数只会匹配主包。
- No module is considered the "main" module. If the module containing packages named on the command line has a go.mod file, it must not contain directives (replace and exclude) that would cause it to be interpreted differently than if it were the main module. The module must not require a higher version of itself.
- 没有模块被视为"主"模块。如果命令行上的包所在的模块有go.mod文件，则该文件不得包含指令（replace和exclude），使其被解释为与主模块不同。该模块不得要求其自身的更高版本。
- Vendor directories are not used in any module. (Vendor directories are not included in the module zip files downloaded by 'go install'.)
- 任何模块中均不使用供应商目录。 （供应商目录未包含在"go install"下载的模块zip文件中。）

If the arguments don't have version suffixes, "go install" may run in module-aware mode or GOPATH mode, depending on the GO111MODULE environment variable and the presence of a go.mod file. See 'go help modules' for details. If module-aware mode is enabled, "go install" runs in the context of the main module.

​	如果参数没有版本后缀，则"go install"可以在模块感知模式或GOPATH模式下运行，这取决于GO111MODULE环境变量和是否存在go.mod文件。有关详细信息，请参见"go help modules"。如果启用了模块感知模式，则"go install"在主模块的上下文中运行。

When module-aware mode is disabled, non-main packages are installed in the directory $GOPATH/pkg/$GOOS_$GOARCH. When module-aware mode is enabled, non-main packages are built and cached but not installed.

​	在禁用模块感知模式时，非主要包将安装在目录`$GOPATH/pkg/$GOOS_$GOARCH`中。启用模块感知模式时，非主要包将被构建和缓存，但不会被安装。

Before Go 1.20, the standard library was installed to `$GOROOT/pkg/$GOOS_$GOARCH`. Starting in Go 1.20, the standard library is built and cached but not installed. Setting GODEBUG=installgoroot=all restores the use of `$GOROOT/pkg/$GOOS_$GOARCH`.

​	在 Go 1.20 之前，标准库被安装到`$GOROOT/pkg/$GOOS_$GOARCH`中。从 Go 1.20 开始，标准库被构建和缓存，但不会被安装。设置 GODEBUG=installgoroot=all 可以恢复对`$GOROOT/pkg/$GOOS_$GOARCH`的使用。

For more about build flags, see 'go help build'.

​	有关构建标志的更多信息，请参见"go help build"。

For more about specifying packages, see 'go help packages'.

​	有关指定包的更多信息，请参见"go help packages"。

See also: go build, go get, go clean.

​	另请参阅：go build、go get、go clean。

#### go list -> 列出包或模块

Usage:

​	用法：

```
go list [-f format] [-json] [-m] [list flags] [build flags] [packages]
```

List lists the named packages, one per line. The most commonly-used flags are -f and -json, which control the form of the output printed for each package. Other list flags, documented below, control more specific details.

​	list 命令将指定的包列出，每个包占一行。最常用的标志是 -f 和 -json，它们控制打印每个包时输出的格式。其他的 list 标志在下面有说明，它们控制更具体的细节。

The default output shows the package import path:

​	默认输出显示包的导入路径：

```
bytes
encoding/json
github.com/gorilla/mux
golang.org/x/net/html
```

The -f flag specifies an alternate format for the list, using the syntax of package template. The default output is equivalent to `-f '{{.ImportPath}}'`. The struct being passed to the template is:

​	`-f` 标志指定了列表的替代格式，使用包模板语法。默认输出相当于 `-f '{{.ImportPath}}'`。传递给模板的结构体是：

``` go
type Package struct {
    Dir           string   // 包源文件所在的目录 directory containing package sources
    ImportPath    string   // 包在目录中的导入路径 import path of package in dir
    ImportComment string   // 包声明的导入注释中的路径 path in import comment on package statement
    Name          string   // 包名 package name
    Doc           string   // 包的文档字符串 package documentation string
    Target        string   // 安装路径 install path
    Shlib         string   // 包含该包的共享库（仅在使用 -linkshared 时设置） the shared library that contains this package (only set when -linkshared)
    Goroot        bool     // 该包是否在 Go 根目录下？ is this package in the Go root?
    Standard      bool     // 该包是否是标准 Go 库的一部分？ is this package part of the standard Go library?
    Stale         bool     // 对于该包，go install 是否会执行任何操作？ would 'go install' do anything for this package?
    StaleReason   string   // Stale==true 的原因说明 explanation for Stale==true
    Root          string   // 包所在的 Go 根目录或 Go path 目录 Go root or Go path dir containing this package
    ConflictDir   string   // 此目录遮盖了 $GOPATH 中的 Dir this directory shadows Dir in $GOPATH
    BinaryOnly    bool     // 仅限二进制包（不再支持） binary-only package (no longer supported)
    ForTest       string   // 该包仅供命名测试使用 package is only for use in named test
    Export        string   // 包含导出数据的文件（使用 -export 时） file containing export data (when using -export)
    BuildID       string   // 编译包的 build ID（使用 -export 时） build ID of the compiled package (when using -export)
    Module        *Module  // 包所在模块的信息（如果有）（可能为 nil） info about package's containing module, if any (can be nil)
    Match         []string // 与此包匹配的命令行模式 command-line patterns matching this package
    DepOnly       bool     // 该包仅作为依赖项，没有被显式列出 package is only a dependency, not explicitly listed
	DefaultGODEBUG string  // default GODEBUG setting, for main packages
    
    // 源文件 Source files
    GoFiles         []string   // .go 源文件（不包括 CgoFiles、TestGoFiles 和 XTestGoFiles） .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)
    CgoFiles        []string   // 导入了 "C" 的 .go 源文件 .go source files that import "C"
    CompiledGoFiles []string   // 向编译器展示的 .go 文件（使用 -compiled 时） .go files presented to compiler (when using -compiled)
    IgnoredGoFiles  []string   //  因构建约束而被忽略的 .go 源文件 .go source files ignored due to build constraints
    IgnoredOtherFiles []string // 因构建约束而被忽略的非 .go 源文件 non-.go source files ignored due to build constraints
    CFiles          []string   // .c 源文件 .c source files
    CXXFiles        []string   // .cc、.cxx 和 .cpp 源文件 .cc, .cxx and .cpp source files
    MFiles          []string   //  .m 源文件 .m source files
    HFiles          []string   // .h、.hh、.hpp 和 .hxx 源文件 .h, .hh, .hpp and .hxx source files
    FFiles          []string   // .f、.F、.for 和 .f90 Fortran 源文件 .f, .F, .for and .f90 Fortran source files
    SFiles          []string   //  .s 源文件 .s source files
    SwigFiles       []string   //  .swig 文件 .swig files
    SwigCXXFiles    []string   // .swigcxx 文件 .swigcxx files
    SysoFiles       []string   // 要添加到档案文件的 .syso 目标文件  .syso object files to add to archive
    TestGoFiles     []string   // 包内的 _test.go 文件 _test.go files in package
    XTestGoFiles    []string   // 包外的 _test.go 文件 _test.go files outside package

    // 嵌入式文件 Embedded files
    EmbedPatterns      []string // //go:embed 模式  //go:embed patterns
    EmbedFiles         []string // 由 EmbedPatterns 匹配的文件 files matched by EmbedPatterns
    TestEmbedPatterns  []string // TestGoFiles 中的 //go:embed 模式  //go:embed patterns in TestGoFiles
    TestEmbedFiles     []string // 由 TestEmbedPatterns 匹配的文件 files matched by TestEmbedPatterns
    XTestEmbedPatterns []string // XTestGoFiles 中的 //go:embed 模式  //go:embed patterns in XTestGoFiles
    XTestEmbedFiles    []string // 由 XTestEmbedPatterns 匹配的文件 files matched by XTestEmbedPatterns

    // Cgo 指令 Cgo directives
    CgoCFLAGS    []string // cgo：C 编译器的标志  cgo: flags for C compiler
    CgoCPPFLAGS  []string // cgo：C 预处理器的标志 cgo: flags for C preprocessor
    CgoCXXFLAGS  []string // cgo：C++ 编译器的标志 cgo: flags for C++ compiler
    CgoFFLAGS    []string // cgo：Fortran 编译器的标志 cgo: flags for Fortran compiler
    CgoLDFLAGS   []string // cgo：链接器的标志 cgo: flags for linker
    CgoPkgConfig []string // cgo：pkg-config 的名称 cgo: pkg-config names

    // 依赖项信息 Dependency information
    Imports      []string          // 此包使用的导入路径 import paths used by this package
    ImportMap    map[string]string // 源导入到 ImportPath 的映射（省略标识条目） map from source import to ImportPath (identity entries omitted)
    Deps         []string          // 所有（递归）导入的依赖项 all (recursively) imported dependencies
    TestImports  []string          // TestGoFiles 中的导入项  imports from TestGoFiles
    XTestImports []string          // XTestGoFiles 中的导入项 imports from XTestGoFiles

    // 错误信息
    Incomplete bool            // 此包或依赖项存在错误 this package or a dependency has an error
    Error      *PackageError   // 加载包时出现的错误 error loading package
    DepsErrors []*PackageError // 加载依赖项时出现的错误 errors loading dependencies
}
```

Packages stored in vendor directories report an ImportPath that includes the path to the vendor directory (for example, "d/vendor/p" instead of "p"), so that the ImportPath uniquely identifies a given copy of a package. The Imports, Deps, TestImports, and XTestImports lists also contain these expanded import paths. See golang.org/s/go15vendor for more about vendoring.

​	存储在 vendor 目录中的包报告 ImportPath，该路径包括供应商目录的路径（例如，"d/vendor/p"而不是"p"），以便 ImportPath 唯一标识给定的包副本。Imports、Deps、TestImports 和 XTestImports 列表也包含这些扩展的导入路径。有关供应商的更多信息，请参见 golang.org/s/go15vendor。

The error information, if any, is

​	如果有错误信息，则为：

``` go
type PackageError struct {
    ImportStack   []string // 从命令行命名的包到此包的最短路径 shortest path from package named on command line to this one
    Pos           string   // 错误的位置（如果存在，则为文件：行：列） position of error (if present, file:line:col)
    Err           string   // 错误本身 the error itself
}
```

The module information is a Module struct, defined in the discussion of list -m below.

​	模块信息是一个 Module 结构，定义在下面 list -m 的讨论中。

The template function "join" calls strings.Join.

​	模板函数"join"调用 strings.Join。

The template function "context" returns the build context, defined as:

​	模板函数"context"返回构建上下文，定义为：

``` go
type Context struct {
    GOARCH        string   // 目标架构 target architecture
    GOOS          string   // 目标操作系统 target operating system
    GOROOT        string   // Go 根目录 Go root
    GOPATH        string   // Go 路径 Go path
    CgoEnabled    bool     // 是否可以使用 cgo whether cgo can be used
    UseAllFiles   bool     // 使用文件，无论是否有 +build 行、文件名 use files regardless of //go:build lines, file names
    Compiler      string   // 在计算目标路径时要使用的编译器 compiler to assume when computing target paths
    BuildTags     []string // 在 +build 行中匹配的构建约束 build constraints to match in //go:build lines
    ToolTags      []string // 工具链特定的构建约束 toolchain-specific build constraints
    ReleaseTags   []string // 当前版本兼容的版本 releases the current release is compatible with
    InstallSuffix string   // 在安装目录的名称中使用的后缀 suffix to use in the name of the install dir
}
```

For more information about the meaning of these fields see the documentation for the go/build package's Context type.

​	有关这些字段的含义的更多信息，请参阅 go/build 包的 Context 类型的文档。

The -json flag causes the package data to be printed in JSON format instead of using the template format. The JSON flag can optionally be provided with a set of comma-separated required field names to be output. If so, those required fields will always appear in JSON output, but others may be omitted to save work in computing the JSON struct.

​	`-json` 标志会导致以 JSON 格式而不是模板格式打印包数据。JSON 标志可以可选地与一组逗号分隔的所需字段名称一起提供，以输出这些所需字段。如果是这样，这些必需的字段将始终出现在 JSON 输出中，但其他字段可能被省略以节省计算 JSON 结构的工作。

The -compiled flag causes list to set CompiledGoFiles to the Go source files presented to the compiler. Typically this means that it repeats the files listed in GoFiles and then also adds the Go code generated by processing CgoFiles and SwigFiles. The Imports list contains the union of all imports from both GoFiles and CompiledGoFiles.

​	`-compiled` 标志会导致 list 将 CompiledGoFiles 设置为提供给编译器的 Go 源文件。通常，这意味着它会重复列在 GoFiles 中的文件，然后还会添加通过处理 CgoFiles 和 SwigFiles 生成的 Go 代码。Imports 列表包含来自 GoFiles 和 CompiledGoFiles 的所有导入的并集。

The -deps flag causes list to iterate over not just the named packages but also all their dependencies. It visits them in a depth-first post-order traversal, so that a package is listed only after all its dependencies. Packages not explicitly listed on the command line will have the DepOnly field set to true.

​	`-deps`标志使列表不仅迭代命名包，还包括它们的所有依赖项。它以深度优先的后序遍历方式访问它们，这样一个包只有在所有依赖项之后才会列出。未在命令行中明确列出的包将具有DepOnly字段设置为true。

The -e flag changes the handling of erroneous packages, those that cannot be found or are malformed. By default, the list command prints an error to standard error for each erroneous package and omits the packages from consideration during the usual printing. With the -e flag, the list command never prints errors to standard error and instead processes the erroneous packages with the usual printing. Erroneous packages will have a non-empty ImportPath and a non-nil Error field; other information may or may not be missing (zeroed).

​	`-e`标志更改对错误包（无法找到或格式错误的包）的处理方式。默认情况下，list命令对于每个错误包在标准错误输出一个错误，并在通常的输出中省略这些包。使用-e标志时，list命令永远不会将错误打印到标准错误输出，并使用通常的输出处理错误包。错误的包将具有非空的ImportPath和非空的Error字段；其他信息可能存在或可能不存在（为零）。

The -export flag causes list to set the Export field to the name of a file containing up-to-date export information for the given package, and the BuildID field to the build ID of the compiled package.

​	`-export`标志使列表将导出字段设置为包含给定包的最新导出信息的文件的名称，并将BuildID字段设置为已编译包的构建ID。

The -find flag causes list to identify the named packages but not resolve their dependencies: the Imports and Deps lists will be empty. With the -find flag, the -deps, -test and -export commands cannot be used.

​	`-find`标志使列表标识命名包但不解析它们的依赖项：Imports和Deps列表将为空。

The -test flag causes list to report not only the named packages but also their test binaries (for packages with tests), to convey to source code analysis tools exactly how test binaries are constructed. The reported import path for a test binary is the import path of the package followed by a ".test" suffix, as in "math/rand.test". When building a test, it is sometimes necessary to rebuild certain dependencies specially for that test (most commonly the tested package itself). The reported import path of a package recompiled for a particular test binary is followed by a space and the name of the test binary in brackets, as in "math/rand [math/rand.test](https://pkg.go.dev/math/rand.test)" or "regexp [sort.test]". The ForTest field is also set to the name of the package being tested ("math/rand" or "sort" in the previous examples).

​	`-test`标志不仅报告命名包，还报告它们的测试二进制文件（对于具有测试的包），以向源代码分析工具传达测试二进制文件的构造方式。测试二进制文件的报告导入路径是包的导入路径后跟一个".test"后缀，例如"math/rand.test"。在构建测试时，有时需要特别为该测试重新构建某些依赖项（最常见的是被测试的包本身）。为特定测试二进制文件重新编译的包的报告导入路径后跟一个空格和方括号中测试二进制文件的名称，例如"math/rand math/rand.test"或"regexp [sort.test]"。ForTest字段还设置为正在测试的包的名称（在前面的示例中为"math/rand"或"sort"）。

The Dir, Target, Shlib, Root, ConflictDir, and Export file paths are all absolute paths.

​	Dir，Target，Shlib，Root，ConflictDir和Export文件路径都是绝对路径。

By default, the lists GoFiles, CgoFiles, and so on hold names of files in Dir (that is, paths relative to Dir, not absolute paths). The generated files added when using the -compiled and -test flags are absolute paths referring to cached copies of generated Go source files. Although they are Go source files, the paths may not end in ".go".

​	默认情况下，GoFiles、CgoFiles等列表保存Dir中文件的名称（即相对于Dir的路径，而不是绝对路径）。当使用-compiled和-test标志时添加的生成文件是引用生成的Go源文件的缓存副本的绝对路径。虽然它们是Go源文件，但路径可能不以".go"结尾。

The -m flag causes list to list modules instead of packages.

​	`-m`标志使列表列出模块而不是包。

When listing modules, the -f flag still specifies a format template applied to a Go struct, but now a Module struct:

​	当列出模块时，`-f`标志仍然指定应用于Go结构的格式模板，但现在是一个Module结构体：

``` go
type Module struct {
    Path       string        // 模块路径 module path
    Query      string        // 对应于此版本的版本查询 version query corresponding to this version
    Version    string        // 模块版本 module version
    Versions   []string      // 可用的模块版本 available module versions
    Replace    *Module       // 被此模块替换 replaced by this module
    Time       *time.Time    // 版本创建时间 time version was created
    Update     *Module       // 可用更新（使用-u） available update (with -u)
    Main       bool          // 是否为主模块？ is this the main module?
    Indirect   bool          // 模块仅由主模块间接需要  module is only indirectly needed by main module
    Dir        string        // 如果有的话，保存文件的本地副本的目录 directory holding local copy of files, if any
    GoMod      string        // 描述模块的go.mod文件的路径（如果有） path to go.mod file describing module, if any
    GoVersion  string        // 模块使用的Go版本 go version used in module
    Retracted  []string      // 撤回信息（使用-retracted或-u） retraction information, if any (with -retracted or -u)
    Deprecated string        // 废弃消息（使用-u） deprecation message, if any (with -u)
    Error      *ModuleError  // 加载模块时的错误  error loading module
    Origin     any           // 模块来源  provenance of module
    Reuse      bool          // 旧模块信息的重用是安全的 reuse of old module info is safe
}

type ModuleError struct {
    Err string // 错误本身 the error itself
}
```

The file GoMod refers to may be outside the module directory if the module is in the module cache or if the -modfile flag is used.

​	GoMod文件所指的文件可能在模块目录之外，如果模块在模块缓存中或使用了-modfile标志，则是如此。

The default output is to print the module path and then information about the version and replacement if any. For example, 'go list -m all' might print:

​	默认输出是打印模块路径，然后是版本和替换信息（如果有）。例如，"go list -m all"可能会打印：

```
my/main/module
golang.org/x/text v0.3.0 => /tmp/text
rsc.io/pdf v0.1.1
```

The Module struct has a String method that formats this line of output, so that the default format is equivalent to -f '{{.String}}'.

​	Module结构具有String方法，用于格式化此行输出，因此默认格式等同于`-f'{{.String}}'`。

Note that when a module has been replaced, its Replace field describes the replacement module, and its Dir field is set to the replacement's source code, if present. (That is, if Replace is non-nil, then Dir is set to Replace.Dir, with no access to the replaced source code.)

​	请注意，当模块已被替换时，其Replace字段描述替换模块，并且如果存在，则其Dir字段设置为替换的源代码。（也就是说，如果Replace不为nil，则Dir设置为Replace.Dir，没有访问被替换源代码的方式。）

The -u flag adds information about available upgrades. When the latest version of a given module is newer than the current one, list -u sets the Module's Update field to information about the newer module. list -u will also set the module's Retracted field if the current version is retracted. The Module's String method indicates an available upgrade by formatting the newer version in brackets after the current version. If a version is retracted, the string "(retracted)" will follow it. For example, 'go list -m -u all' might print:

​	`-u`标志会添加有关可用升级的信息。当给定模块的最新版本比当前版本更新时，list -u会将模块的Update字段设置为有关更高版本模块的信息。如果当前版本被撤回，则list -u还将设置模块的Retracted字段。如果可用升级，则模块的String方法通过在当前版本后用括号中的较新版本进行格式化来指示可用升级。如果版本被撤回，则字符串"（撤回）"将跟随它。例如，"go list -m -u all"可能会打印：

```
my/main/module
golang.org/x/text v0.3.0 [v0.4.0] => /tmp/text
rsc.io/pdf v0.1.1 (retracted) [v0.1.2]
```

(For tools, 'go list -m -u -json all' may be more convenient to parse.)

(对于工具而言，'go list -m -u -json all' 更容易解析。)

The -versions flag causes list to set the Module's Versions field to a list of all known versions of that module, ordered according to semantic versioning, earliest to latest. The flag also changes the default output format to display the module path followed by the space-separated version list.

​	`-versions` 标志会导致 list 将模块的 Versions 字段设置为该模块所有已知版本的列表，按照语义化版本控制的顺序从早到晚排序。该标志还更改默认输出格式，以模块路径为开头，后跟由空格分隔的版本列表。

The -retracted flag causes list to report information about retracted module versions. When -retracted is used with -f or -json, the Retracted field will be set to a string explaining why the version was retracted. The string is taken from comments on the retract directive in the module's go.mod file. When -retracted is used with -versions, retracted versions are listed together with unretracted versions. The -retracted flag may be used with or without -m.

​	`-retracted` 标志会导致 list 报告关于已撤销模块版本的信息。当 -retracted 与 -f 或 -json 结合使用时，Retracted 字段将被设置为一个字符串，解释为什么该版本被撤销。该字符串取自模块的 go.mod 文件中撤销指令的注释。当 -retracted 与 -versions 结合使用时，已撤销的版本会与未撤销的版本一起列出。-retracted 标志可以与或不带 -m 一起使用。

The arguments to list -m are interpreted as a list of modules, not packages. The main module is the module containing the current directory. The active modules are the main module and its dependencies. With no arguments, list -m shows the main module. With arguments, list -m shows the modules specified by the arguments. Any of the active modules can be specified by its module path. The special pattern "all" specifies all the active modules, first the main module and then dependencies sorted by module path. A pattern containing "..." specifies the active modules whose module paths match the pattern. A query of the form path@version specifies the result of that query, which is not limited to active modules. See 'go help modules' for more about module queries.

​	`list -m` 的参数被解释为模块列表，而不是包列表。主模块是包含当前目录的模块。活动模块是主模块和其依赖项。如果没有参数，list -m 会显示主模块。如果有参数，list -m 会显示由参数指定的模块。任何活动模块都可以用其模块路径来指定。特殊模式 "all" 指定所有活动模块，首先是主模块，然后是按模块路径排序的依赖项。包含 "..." 的模式指定其模块路径匹配该模式的活动模块。形式为 path@version 的查询指定该查询的结果，不限于活动模块。有关模块查询的详细信息，请参见 'go help modules'。

The template function "module" takes a single string argument that must be a module path or query and returns the specified module as a Module struct. If an error occurs, the result will be a Module struct with a non-nil Error field.

​	模板函数 "module" 接受一个字符串参数，必须是模块路径或查询，并将指定的模块返回为 Module 结构体。如果发生错误，结果将是一个带有非空错误字段的 Module 结构体。

When using -m, the -reuse=old.json flag accepts the name of file containing the JSON output of a previous 'go list -m -json' invocation with the same set of modifier flags (such as -u, -retracted, and -versions). The go command may use this file to determine that a module is unchanged since the previous invocation and avoid redownloading information about it. Modules that are not redownloaded will be marked in the new output by setting the Reuse field to true. Normally the module cache provides this kind of reuse automatically; the -reuse flag can be useful on systems that do not preserve the module cache.

​	使用 -m 时，-reuse=old.json 标志接受以前 'go list -m -json' 调用的 JSON 输出文件名，该调用具有相同的修改器标志集（例如 -u、-retracted 和 -versions）。go 命令可以使用此文件确定自上次调用以来模块未更改，并避免重新下载有关它的信息。未重新下载的模块将通过将 Reuse 字段设置为 true 在新输出中标记。通常，模块缓存会自动提供此类重用；-reuse 标志可用于不保留模块缓存的系统。

For more about build flags, see 'go help build'.

​	有关构建标志的更多信息，请参见 'go help build'。

For more about specifying packages, see 'go help packages'.

​	有关指定包的更多信息，请参见 'go help packages'。

For more about modules, see https://golang.org/ref/mod.

​	有关模块的更多信息，请参见 https://golang.org/ref/mod。

#### go mod -> 模块维护

Go mod provides access to operations on modules.

​	go mod 提供了对模块操作的访问。

Note that support for modules is built into all the go commands, not just 'go mod'. For example, day-to-day adding, removing, upgrading, and downgrading of dependencies should be done using 'go get'. See 'go help modules' for an overview of module functionality.

​	请注意，对模块的支持内置于所有 go 命令中，而不仅仅是 'go mod'。例如，日常添加、删除、升级和降级依赖关系应使用 'go get' 完成。有关模块功能的概述，请参阅 'go help modules'。

Usage:

​	用法：

```
go mod <command> [arguments]
```

The commands are:

​	命令如下：

```
download 	将模块下载到本地缓存 download modules to local cache
edit 		从工具或脚本编辑 go.mod edit go.mod from tools or scripts
graph 		打印模块依赖关系图 print module requirement graph
init 		在当前目录中初始化新模块  initialize new module in current directory
tidy 		添加缺少的模块并删除未使用的模块 add missing and remove unused modules
vendor 		制作依赖关系的供应商副本 make vendored copy of dependencies
verify 		验证依赖项具有预期的内容 verify dependencies have expected content
why 		解释需要哪些包或模块 explain why packages or modules are needed
```

Use "`go help mod <command>`" for more information about a command.

​	使用 "`go help mod <command>`" 查看有关命令的更多信息。

#### go mod download -> 将模块下载到本地缓存

Usage:

​	用法：

```
go mod download [-x] [-json] [-reuse=old.json] [modules]
```

Download downloads the named modules, which can be module patterns selecting dependencies of the main module or module queries of the form path@version.

​	download 命令会下载指定的模块，可以是选择主模块依赖项的模块模式，也可以是形式为 `path@version` 的模块查询。

With no arguments, download applies to the modules needed to build and test the packages in the main module: the modules explicitly required by the main module if it is at 'go 1.17' or higher, or all transitively-required modules if at 'go 1.16' or lower.

​	如果没有指定参数，则 download 命令适用于构建和测试主模块中的包所需的模块：如果主模块处于 'go 1.17' 或更高版本，则为主模块明确要求的模块，否则为所有必需的传递模块（对于 'go 1.16' 或更低版本）。

The go command will automatically download modules as needed during ordinary execution. The "go mod download" command is useful mainly for pre-filling the local cache or to compute the answers for a Go module proxy.

​	在普通执行期间，go 命令将自动根据需要下载模块。"go mod download" 命令主要用于预先填充本地缓存或计算 Go 模块代理的答案。

By default, download writes nothing to standard output. It may print progress messages and errors to standard error.

​	默认情况下，download 不会向标准输出写入任何内容。它可能会将进度消息和错误打印到标准错误。

The -json flag causes download to print a sequence of JSON objects to standard output, describing each downloaded module (or failure), corresponding to this Go struct:

​	使用 -json 标志将导致 download 向标准输出打印一系列 JSON 对象，描述每个已下载的模块（或失败），对应于此 Go 结构体：

``` go
type Module struct {
    Path     string // 模块路径 module path
    Query    string // 版本查询，对应于此版本 version query corresponding to this version
    Version  string // 模块版本 module version
    Error    string // 加载模块时出现的错误 error loading module
    Info     string // 缓存的 .info 文件的绝对路径 absolute path to cached .info file
    GoMod    string // 缓存的 .mod 文件的绝对路径 absolute path to cached .mod file
    Zip      string // 缓存的 .zip 文件的绝对路径 absolute path to cached .zip file
    Dir      string // 缓存的源根目录的绝对路径 absolute path to cached source root directory
    Sum      string // 路径、版本的校验和（如 go.sum） checksum for path, version (as in go.sum)
    GoModSum string // go.mod 的校验和（如 go.sum） checksum for go.mod (as in go.sum)
    Origin   any    // 模块来源的证明 provenance of module
    Reuse    bool   // 重用旧模块信息是安全的 reuse of old module info is safe
}
```

The -reuse flag accepts the name of file containing the JSON output of a previous 'go mod download -json' invocation. The go command may use this file to determine that a module is unchanged since the previous invocation and avoid redownloading it. Modules that are not redownloaded will be marked in the new output by setting the Reuse field to true. Normally the module cache provides this kind of reuse automatically; the -reuse flag can be useful on systems that do not preserve the module cache.

​	`-reuse` 标记接受包含之前的 'go mod download -json' 调用的 JSON 输出的文件名。go 命令可以使用此文件确定模块是否自上次调用以来未更改并避免重新下载。不重新下载的模块将通过将 Reuse 字段设置为 true 来标记在新输出中。通常模块缓存会自动提供此类重用；-reuse 标记可以在不保留模块缓存的系统上很有用。

The -x flag causes download to print the commands download executes.

​	`-x` 标记会导致 download 打印出执行的命令。

See https://golang.org/ref/mod#go-mod-download for more about 'go mod download'.

​	有关"go mod download"的更多信息，请参见 https://golang.org/ref/mod#go-mod-download。

See https://golang.org/ref/mod#version-queries for more about version queries.

​	有关版本查询的更多信息，请参见 https://golang.org/ref/mod#version-queries。

#### go mod edit -> 从工具或脚本编辑 go.mod 

Usage:

​	用法：

```
go mod edit [editing flags] [-fmt|-print|-json] [go.mod]
```

Edit provides a command-line interface for editing go.mod, for use primarily by tools or scripts. It reads only go.mod; it does not look up information about the modules involved. By default, edit reads and writes the go.mod file of the main module, but a different target file can be specified after the editing flags.

​	`edit` 提供了一个命令行界面，用于编辑 go.mod，主要供工具或脚本使用。它只读取 go.mod，不会查找有关所涉及的模块的信息。默认情况下，edit 读取和写入主模块的 go.mod 文件，但在编辑标记后可以指定不同的目标文件。

The editing flags specify a sequence of editing operations.

​	编辑标记指定一系列编辑操作。

The -fmt flag reformats the go.mod file without making other changes. This reformatting is also implied by any other modifications that use or rewrite the go.mod file. The only time this flag is needed is if no other flags are specified, as in 'go mod edit -fmt'.

​	`-fmt` 标记重新格式化 go.mod 文件而不进行其他更改。此重新格式化也隐含在使用或重写 go.mod 文件的任何其他修改中。只有在未指定其他标记时（例如 'go mod edit -fmt'）才需要此标记。

The -module flag changes the module's path (the go.mod file's module line).

​	`-module` 标记更改模块的路径（go.mod 文件的模块行）。

The -require=path@version and -droprequire=path flags add and drop a requirement on the given module path and version. Note that -require overrides any existing requirements on path. These flags are mainly for tools that understand the module graph. Users should prefer 'go get path@version' or 'go get path@none', which make other go.mod adjustments as needed to satisfy constraints imposed by other modules.

​	`-require=path@version` 和 `-droprequire=path` 标志添加和删除给定模块路径和版本的要求。请注意，-require 覆盖路径上的所有现有要求。这些标志主要用于了解模块图的工具。用户应该优先使用 'go get path@version' 或 'go get path@none'，这些命令会根据其他模块的约束条件进行必要的 go.mod 调整。

The -exclude=path@version and -dropexclude=path@version flags add and drop an exclusion for the given module path and version. Note that -exclude=path@version is a no-op if that exclusion already exists.

​	`-exclude=path@version` 和 `-dropexclude=path@version` 标志添加和删除给定模块路径和版本的排除项。请注意，如果该排除项已存在，则 -exclude=path@version 不起作用。

The -replace=old[@v]=new[@v] flag adds a replacement of the given module path and version pair. If the @v in old@v is omitted, a replacement without a version on the left side is added, which applies to all versions of the old module path. If the @v in new@v is omitted, the new path should be a local module root directory, not a module path. Note that -replace overrides any redundant replacements for old[@v], so omitting @v will drop existing replacements for specific versions.

​	`-replace=old[@v]=new[@v]` 标志添加给定模块路径和版本对的替换项。如果在 old@v 中省略 @v，则添加左侧没有版本的替换项，该替换项适用于该旧模块路径的所有版本。如果在 new@v 中省略 @v，则新路径应为本地模块根目录，而不是模块路径。请注意，-replace 覆盖了 old[@v] 的任何冗余替换项，因此省略 @v 将删除特定版本的现有替换项。

The -dropreplace=old[@v] flag drops a replacement of the given module path and version pair. If the @v is omitted, a replacement without a version on the left side is dropped.

​	`-dropreplace=old[@v]` 标志删除给定模块路径和版本对的替换项。如果省略 @v，则删除左侧没有版本的替换项。

The -retract=version and -dropretract=version flags add and drop a retraction on the given version. The version may be a single version like "v1.2.3" or a closed interval like "[v1.1.0,v1.1.9]". Note that -retract=version is a no-op if that retraction already exists.

​	`-retract=version` 和 `-dropretract=version` 标志添加和删除给定版本的撤销。版本可以是单个版本，例如"v1.2.3"，也可以是封闭间隔，例如"[v1.1.0，v1.1.9]"。请注意，如果该撤回已经存在，则 -retract=version 不起作用。

The -require, -droprequire, -exclude, -dropexclude, -replace, -dropreplace, -retract, and -dropretract editing flags may be repeated, and the changes are applied in the order given.

​	`-require`、`-droprequire`、`-exclude`、`-dropexclude`、`-replace`、`-dropreplace`、`-retract` 和 `-dropretract` 编辑标志可以重复使用，更改按给定的顺序应用。

The -go=version flag sets the expected Go language version.

​	`-go=version` 标志设置预期的 Go 语言版本。

The -toolchain=name flag sets the Go toolchain to use.

The -print flag prints the final go.mod in its text format instead of writing it back to go.mod.

​	`-print` 标志以其文本格式打印最终的 go.mod，而不是将其写回 go.mod。

The -json flag prints the final go.mod file in JSON format instead of writing it back to go.mod. The JSON output corresponds to these Go types:

​	`-json` 标志以 JSON 格式打印最终的 go.mod 文件，而不是将其写回 go.mod。JSON 输出对应于这些 Go 类型：

``` go
type Module struct {
	Path    string
	Version string
}

type GoMod struct {
	Module  ModPath
	Go      string
	Require []Require
	Exclude []Module
	Replace []Replace
	Retract []Retract
}

type ModPath struct {
	Path       string
	Deprecated string
}

type Require struct {
	Path string
	Version string
	Indirect bool
}

type Replace struct {
	Old Module
	New Module
}

type Retract struct {
	Low       string
	High      string
	Rationale string
}
```

Retract entries representing a single version (not an interval) will have the "Low" and "High" fields set to the same value.

​	retract 条目表示单个版本（而非区间），"Low" 和 "High" 字段将设置为相同的值。

Note that this only describes the go.mod file itself, not other modules referred to indirectly. For the full set of modules available to a build, use 'go list -m -json all'.

​	请注意，这仅描述 go.mod 文件本身，而不是间接引用的其他模块。要查看可用于构建的完整模块集，请使用 'go list -m -json all'。

Edit also provides the -C, -n, and -x build flags.

​	edit 还提供了 -C、-n 和 -x 构建标志。

See https://golang.org/ref/mod#go-mod-edit for more about 'go mod edit'.

​	有关 'go mod edit' 的更多信息，请参见 https://golang.org/ref/mod#go-mod-edit。

#### go mod graph -> 打印模块需求图 

Usage:

​	用法：

```
go mod graph [-go=version] [-x]
```

Graph prints the module requirement graph (with replacements applied) in text form. Each line in the output has two space-separated fields: a module and one of its requirements. Each module is identified as a string of the form path@version, except for the main module, which has no @version suffix.

​	graph 以文本形式打印模块需求图（已应用替换）。输出中的每一行都有两个空格分隔的字段：模块和其要求之一。除主模块外，每个模块都被标识为 path@version 格式的字符串，而主模块没有 @version 后缀。

The -go flag causes graph to report the module graph as loaded by the given Go version, instead of the version indicated by the 'go' directive in the go.mod file.

​	`-go` 标志导致 graph 报告按照给定的 Go 版本加载的模块图，而不是 go.mod 文件中 'go' 指令指示的版本。

The -x flag causes graph to print the commands graph executes.

​	`-x` 标志导致 graph 打印 graph 执行的命令。

See https://golang.org/ref/mod#go-mod-graph for more about 'go mod graph'.

​	有关 'go mod graph' 的更多信息，请参见 https://golang.org/ref/mod#go-mod-graph。

#### go mod init -> 在当前目录中初始化新模块

Usage:

​	用法：

```
go mod init [module-path]
```

Init initializes and writes a new go.mod file in the current directory, in effect creating a new module rooted at the current directory. The go.mod file must not already exist.

​	`init` 在当前目录中初始化并写入新的 go.mod 文件，实际上创建一个以当前目录为根的新模块。go.mod 文件不能已经存在。

Init accepts one optional argument, the module path for the new module. If the module path argument is omitted, init will attempt to infer the module path using import comments in .go files, vendoring tool configuration files (like Gopkg.lock), and the current directory (if in GOPATH).

​	`init` 接受一个可选参数，新模块的模块路径。如果省略模块路径参数，则 init 将尝试使用 .go 文件中的导入注释、供应商工具配置文件（如 Gopkg.lock）和当前目录（如果在 GOPATH 中）来推断模块路径。

If a configuration file for a vendoring tool is present, init will attempt to import module requirements from it.

​	如果存在供应商工具的配置文件，则 init 将尝试从其导入模块要求。

See https://golang.org/ref/mod#go-mod-init for more about 'go mod init'.

​	有关 'go mod init' 的更多信息，请参见 https://golang.org/ref/mod#go-mod-init。

#### go mod tidy -> 添加缺失和删除未使用的模块 

Usage:

​	用法：

```
go mod tidy [-e] [-v] [-x] [-go=version] [-compat=version]
```

Tidy makes sure go.mod matches the source code in the module. It adds any missing modules necessary to build the current module's packages and dependencies, and it removes unused modules that don't provide any relevant packages. It also adds any missing entries to go.sum and removes any unnecessary ones.

​	`tidy` 确保 go.mod 与模块中的源代码匹配。它添加了构建当前模块的包和依赖项所需的任何缺失模块，并删除未提供任何相关包的未使用模块。它还添加任何缺失的条目到 go.sum，并删除任何不必要的条目。

The -v flag causes tidy to print information about removed modules to standard error.

​	`-v` 标志导致 tidy 打印有关已删除模块的信息到标准错误。

The -e flag causes tidy to attempt to proceed despite errors encountered while loading packages.

​	`-e` 标志导致 tidy 尝试在加载包时遇到错误时继续进行。

The -go flag causes tidy to update the 'go' directive in the go.mod file to the given version, which may change which module dependencies are retained as explicit requirements in the go.mod file. (Go versions 1.17 and higher retain more requirements in order to support lazy module loading.)

​	`-go` 标志会导致 tidy 更新 go.mod 文件中的 'go' 指令到给定的版本，这可能会改变哪些模块依赖项在 go.mod 文件中作为显式要求保留。（Go 版本 1.17 及更高版本会保留更多的要求以支持懒惰模块加载。）

The -compat flag preserves any additional checksums needed for the 'go' command from the indicated major Go release to successfully load the module graph, and causes tidy to error out if that version of the 'go' command would load any imported package from a different module version. By default, tidy acts as if the -compat flag were set to the version prior to the one indicated by the 'go' directive in the go.mod file.

​	`-compat` 标志保留所需的任何附加校验和，以便从指定的主 Go 发布版本成功加载模块图形，并且如果该 'go' 命令版本会从不同模块版本加载任何已导入的包，则会导致 tidy 出错。默认情况下，tidy 表现得好像将 -compat 标志设置为 go.mod 文件中指示的版本的前一个版本。

The -x flag causes tidy to print the commands download executes.

​	`-x` 标志会导致 tidy 打印下载命令执行的命令。

See https://golang.org/ref/mod#go-mod-tidy for more about 'go mod tidy'.

​	有关 'go mod tidy' 的更多信息，请参见 https://golang.org/ref/mod#go-mod-tidy。

#### go mod vendor -> 创建依赖项的副本以供vendor    Make vendored copy of dependencies

Usage:

​	用法：

```
go mod vendor [-e] [-v] [-o outdir]
```

Vendor resets the main module's vendor directory to include all packages needed to build and test all the main module's packages. It does not include test code for vendored packages.

​	vendor 重置主模块的供应商目录，以包括构建和测试所有主模块包所需的所有包。它不包括供应商包的测试代码。

The -v flag causes vendor to print the names of vendored modules and packages to standard error.

​	`-v` 标志导致供应商将供应商的模块和包名称打印到标准错误输出。

The -e flag causes vendor to attempt to proceed despite errors encountered while loading packages.

​	`-e` 标志会导致供应商尝试在加载包时遇到错误后继续运行。

The -o flag causes vendor to create the vendor directory at the given path instead of "vendor". The go command can only use a vendor directory named "vendor" within the module root directory, so this flag is primarily useful for other tools.

​	`-o` 标志会导致供应商在给定路径处创建供应商目录，而不是在 "vendor" 中。go 命令只能在模块根目录中命名为 "vendor" 的供应商目录，因此该标志主要对其他工具有用。

See https://golang.org/ref/mod#go-mod-vendor for more about 'go mod vendor'.

​	有关 'go mod vendor' 的更多信息，请参见 https://golang.org/ref/mod#go-mod-vendor。

#### go mod verify -> 验证依赖项是否具有预期内容

Usage:

​	用法：

```
go mod verify
```

Verify checks that the dependencies of the current module, which are stored in a local downloaded source cache, have not been modified since being downloaded. If all the modules are unmodified, verify prints "all modules verified." Otherwise it reports which modules have been changed and causes 'go mod' to exit with a non-zero status.

​	verify 检查当前模块的依赖项（存储在本地下载的源缓存中）是否在下载后已被修改。如果所有模块都未被修改，则 verify 打印 "all modules verified."。否则它会报告哪些模块已被更改，并导致 'go mod' 退出并显示一个非零状态码。

See https://golang.org/ref/mod#go-mod-verify for more about 'go mod verify'.

​	有关"go mod verify"的更多信息，请参见 https://golang.org/ref/mod#go-mod-verify。

#### go mod why -> 解释为什么需要包或模块

Usage:

​	用法：

```
go mod why [-m] [-vendor] packages...
```

Why shows a shortest path in the import graph from the main module to each of the listed packages. If the -m flag is given, why treats the arguments as a list of modules and finds a path to any package in each of the modules.

​	`why` 显示从主模块到每个列出的包的导入图中的最短路径。如果指定了 -m 标志，则 why 将参数视为模块列表，并为每个模块找到一个路径到任何包。

By default, why queries the graph of packages matched by "go list all", which includes tests for reachable packages. The -vendor flag causes why to exclude tests of dependencies.

​	默认情况下，`why` 查询与 "go list all" 匹配的包图，其中包括可访问包的测试。-vendor 标志使 why 排除依赖项的测试。

The output is a sequence of stanzas, one for each package or module name on the command line, separated by blank lines. Each stanza begins with a comment line "# package" or "# module" giving the target package or module. Subsequent lines give a path through the import graph, one package per line. If the package or module is not referenced from the main module, the stanza will display a single parenthesized note indicating that fact.

​	输出是一系列段落，每个包或模块名称在命令行上都有一个段落，段落之间用空行分隔。每个段落以注释行 "# package" 或 "# module" 开头，给出目标包或模块。随后的行以一个包为一行，给出了导入图中的路径。如果从主模块中没有引用该包或模块，则段落将显示一个括号注释，指出这一点。

For example:

​	例如：

```
$ go mod why golang.org/x/text/language golang.org/x/text/encoding
# golang.org/x/text/language
rsc.io/quote
rsc.io/sampler
golang.org/x/text/language

# golang.org/x/text/encoding
(main module does not need package golang.org/x/text/encoding)
$
```

See https://golang.org/ref/mod#go-mod-why for more about 'go mod why'.

​	有关"`go mod why`"的更多信息，请参见https://golang.org/ref/mod#go-mod-why。

#### go work -> 工作区维护

Work provides access to operations on workspaces.

​	work 提供了对工作区进行操作的访问。

Note that support for workspaces is built into many other commands, not just 'go work'.

​	请注意，对于工作区的支持内置于许多其他命令中，而不仅仅是 'go work'。

See 'go help modules' for information about Go's module system of which workspaces are a part.

​	有关工作区的信息，请参见 'go help modules'。工作区是 Go 模块系统的一部分。

See https://go.dev/ref/mod#workspaces for an in-depth reference on workspaces.

​	有关工作区的深入参考，请参见 https://go.dev/ref/mod#workspaces。

See https://go.dev/doc/tutorial/workspaces for an introductory tutorial on workspaces.

​	有关工作区的入门教程，请参见 https://go.dev/doc/tutorial/workspaces。

A workspace is specified by a go.work file that specifies a set of module directories with the "use" directive. These modules are used as root modules by the go command for builds and related operations. A workspace that does not specify modules to be used cannot be used to do builds from local modules.

​	工作区由一个 go.work 文件指定，该文件使用 "use" 指令指定一组模块目录。这些模块由 go 命令用作构建和相关操作的根模块。未指定要使用的模块的工作区无法用于从本地模块进行构建。

go.work files are line-oriented. Each line holds a single directive, made up of a keyword followed by arguments. For example:

​	go.work 文件是面向行的。每行包含一个指令，由关键字和参数组成。例如：

```
go 1.18

use ../foo/bar
use ./baz

replace example.com/foo v1.2.3 => example.com/bar v1.4.5
```

The leading keyword can be factored out of adjacent lines to create a block, like in Go imports.

​	前导关键字可以从相邻行中分离出来以创建块，就像 Go 导入一样。

```
use (
  ../foo/bar
  ./baz
)
```

The use directive specifies a module to be included in the workspace's set of main modules. The argument to the use directive is the directory containing the module's go.mod file.

​	use 指令指定要包含在工作区主模块集中的模块。use 指令的参数是包含该模块的 go.mod 文件的目录。

The go directive specifies the version of Go the file was written at. It is possible there may be future changes in the semantics of workspaces that could be controlled by this version, but for now the version specified has no effect.

​	go 指令指定了文件编写的 Go 版本。可能有将来的更改，这些更改可能由该版本控制工作区的语义，但是目前指定的版本没有影响。

The replace directive has the same syntax as the replace directive in a go.mod file and takes precedence over replaces in go.mod files. It is primarily intended to override conflicting replaces in different workspace modules.

​	replace 指令具有与 go.mod 文件中的 replace 指令相同的语法，并优先于 go.mod 文件中的替换。它主要用于覆盖不同工作区模块中的冲突替换。

To determine whether the go command is operating in workspace mode, use the "go env GOWORK" command. This will specify the workspace file being used.

​	要确定 go 命令是否在工作区模式下运行，请使用 "go env GOWORK" 命令。这将指定正在使用的工作区文件。

Usage:

​	用法：

```
go work <command> [arguments]
```

The commands are:

​	命令为：

```
edit        从工具或脚本编辑 go.work
init        初始化工作区文件
sync        同步工作区构建列表到模块
use         向工作区文件中添加模块
```

Use "`go help work <command>`" for more information about a command.

​	有关命令的更多信息，请使用 "`go help work <command>`"。

#### go work edit -> 从工具或脚本编辑 go.work

Usage:

​	用法：

```
go work edit [editing flags] [go.work]
```

Edit provides a command-line interface for editing go.work, for use primarily by tools or scripts. It only reads go.work; it does not look up information about the modules involved. If no file is specified, Edit looks for a go.work file in the current directory and its parent directories

​	`edit` 提供了一个命令行界面用于编辑 go.work 文件，主要供工具或脚本使用。它只读取 go.work 文件，不会查找有关模块的信息。如果未指定文件，则 Edit 会在当前目录和其父目录中查找 go.work 文件。

The editing flags specify a sequence of editing operations.

​	`editing`标志指定一系列编辑操作。

The -fmt flag reformats the go.work file without making other changes. This reformatting is also implied by any other modifications that use or rewrite the go.mod file. The only time this flag is needed is if no other flags are specified, as in 'go work edit -fmt'.

​	`-fmt` 标志重新格式化 go.work 文件，而不做其他更改。任何使用或重写 go.mod 文件的其他修改都会隐含此重新格式化。只有在没有指定其他标志的情况下才需要此标志，例如 go work edit -fmt。

The -use=path and -dropuse=path flags add and drop a use directive from the go.work file's set of module directories.

​	`-use=path` 和 `-dropuse=path` 标志向 go.work 文件的模块目录集添加和删除 use 指令。

The -replace=old[@v]=new[@v] flag adds a replacement of the given module path and version pair. If the @v in old@v is omitted, a replacement without a version on the left side is added, which applies to all versions of the old module path. If the @v in new@v is omitted, the new path should be a local module root directory, not a module path. Note that -replace overrides any redundant replacements for old[@v], so omitting @v will drop existing replacements for specific versions.

​	`-replace=old[@v]=new[@v]` 标志添加给定模块路径和版本对的替换。如果省略 old@v 中的 @v，则会添加左侧没有版本的替换，这适用于旧模块路径的所有版本。如果省略 new@v 中的 @v，则新路径应为本地模块根目录，而不是模块路径。请注意，-replace 会覆盖 old[@v] 的任何冗余替换，因此省略 @v 将删除特定版本的现有替换。

The -dropreplace=old[@v] flag drops a replacement of the given module path and version pair. If the @v is omitted, a replacement without a version on the left side is dropped.

​	`-dropreplace=old[@v]` 标志删除给定模块路径和版本对的替换。如果省略 @v，则会删除左侧没有版本的替换。

The -use, -dropuse, -replace, and -dropreplace, editing flags may be repeated, and the changes are applied in the order given.

​	`-use`、`-dropuse`、`-replace` 和 `-dropreplace` 编辑标志可以重复，更改按给定的顺序应用。

The -go=version flag sets the expected Go language version.

​	`-go=version` 标志设置期望的 Go 语言版本。

The `-toolchain=name` flag sets the Go toolchain to use.

The -print flag prints the final go.work in its text format instead of writing it back to go.mod.

​	`-print` 标志以文本格式打印最终的 go.work，而不是将其写回 go.mod。

The -json flag prints the final go.work file in JSON format instead of writing it back to go.mod. The JSON output corresponds to these Go types:

​	`-json` 标志以 JSON 格式打印最终的 go.work 文件，而不是将其写回 go.mod。JSON 输出对应于这些 Go 类型：

``` go
type GoWork struct {
	Go      string
	Use     []Use
	Replace []Replace
}

type Use struct {
	DiskPath   string
	ModulePath string
}

type Replace struct {
	Old Module
	New Module
}

type Module struct {
	Path    string
	Version string
}
```

See the workspaces reference at https://go.dev/ref/mod#workspaces for more information.

​	有关更多信息，请参见工作区参考：[Go模块参考中的工作区](../../GoModulesReference/Workspaces)。

#### go work init -> 初始化工作区文件

Usage:

​	用法：

```
go work init [moddirs]
```

Init initializes and writes a new go.work file in the current directory, in effect creating a new workspace at the current directory.

​	init 在当前目录中初始化并编写新的 go.work 文件，实际上创建了一个新的工作区。

go work init optionally accepts paths to the workspace modules as arguments. If the argument is omitted, an empty workspace with no modules will be created.

​	go work init 可选地接受工作区模块的路径作为参数。如果省略参数，则将创建一个没有模块的空工作区。

Each argument path is added to a use directive in the go.work file. The current go version will also be listed in the go.work file.

​	每个路径参数都会添加到 go.work 文件的 use 指令中。当前的 Go 版本也将在 go.work 文件中列出。

See the workspaces reference at https://go.dev/ref/mod#workspaces for more information.

​	有关更多信息，请参见工作区参考：[Go模块参考中的工作区](../../GoModulesReference/Workspaces)。

#### go work sync -> 同步工作区的构建清单到模块 

Usage:

​	用法：

```
go work sync
```

Sync syncs the workspace's build list back to the workspace's modules

​	sync命令将工作区的构建清单同步回工作区的模块

The workspace's build list is the set of versions of all the (transitive) dependency modules used to do builds in the workspace. go work sync generates that build list using the Minimal Version Selection algorithm, and then syncs those versions back to each of modules specified in the workspace (with use directives).

​	工作区的构建清单是所有（传递）依赖模块的版本集，这些依赖模块用于在工作区进行构建操作。 go work sync使用最小版本选择算法生成该构建清单，然后将这些版本与工作区指定的每个模块（使用指令）同步。

The syncing is done by sequentially upgrading each of the dependency modules specified in a workspace module to the version in the build list if the dependency module's version is not already the same as the build list's version. Note that Minimal Version Selection guarantees that the build list's version of each module is always the same or higher than that in each workspace module.

​	如果依赖模块的版本与构建清单的版本不同，则按顺序将工作区模块中指定的每个依赖模块升级到构建清单中的版本。请注意，最小版本选择保证构建清单中每个模块的版本始终相同或更高于每个工作区模块中的版本。

See the workspaces reference at https://go.dev/ref/mod#workspaces for more information.

​	有关更多信息，请参见工作区参考：[Go模块参考中的工作区](../../GoModulesReference/Workspaces)。

#### go work use -> 将模块添加到工作区文件

Usage:

​	用法：

```
go work use [-r] moddirs
```

Use provides a command-line interface for adding directories, optionally recursively, to a go.work file.

​	`use`提供了一个命令行界面，用于将目录（可选地进行递归搜索）添加到go.work文件中。

A use directive will be added to the go.work file for each argument directory listed on the command line go.work file, if it exists, or removed from the go.work file if it does not exist. Use fails if any remaining use directives refer to modules that do not exist.

​	如果go.work文件中存在，那么每个命令行上列出的参数目录都将在go.work文件中添加一个use指令，否则将从go.work文件中删除。

Use updates the go line in go.work to specify a version at least as new as all the go lines in the used modules, both preexisting ones and newly added ones. With no arguments, this update is the only thing that go work use does.

​	`use` 会更新 `go.work` 文件中的 `go` 行，以指定一个版本，至少与所有已使用模块（包括预先存在的和新添加的）中的 `go` 行一样新。如果没有参数，那么 `go work use` 命令只会执行这个更新操作。

The -r flag searches recursively for modules in the argument directories, and the use command operates as if each of the directories were specified as arguments: namely, use directives will be added for directories that exist, and removed for directories that do not exist.

​	`-r`标志在参数目录中递归搜索模块，并且use命令操作方式与指定每个目录作为参数相同：即对于存在的目录将添加use指令，并且对于不存在的目录将删除use指令。

See the workspaces reference at https://go.dev/ref/mod#workspaces for more information.

​	有关更多信息，请参见工作区参考：[Go模块参考中的工作区](../../GoModulesReference/Workspaces)。

#### go run -> 编译并运行Go程序 

Usage:

​	用法：

```
go run [build flags] [-exec xprog] package [arguments...]
```

Run compiles and runs the named main Go package. Typically the package is specified as a list of .go source files from a single directory, but it may also be an import path, file system path, or pattern matching a single known package, as in 'go run .' or 'go run my/cmd'.

​	run编译并运行指定的Go 主包。通常，该程序被指定为来自单个目录的一系列.go源文件的列表，但也可以是导入路径、文件系统路径或与单个已知程序匹配的模式，例如"go run ."或"go run my/cmd"。

If the package argument has a version suffix (like @latest or @v1.0.0), "go run" builds the program in module-aware mode, ignoring the go.mod file in the current directory or any parent directory, if there is one. This is useful for running programs without affecting the dependencies of the main module.

​	如果包参数有版本后缀（如`@latest`或`@v1.0.0`），"go run"将在模块感知模式下构建程序，忽略当前目录或任何父目录中的go.mod文件。这对于运行程序而不影响主模块的依赖项非常有用。

If the package argument doesn't have a version suffix, "go run" may run in module-aware mode or GOPATH mode, depending on the GO111MODULE environment variable and the presence of a go.mod file. See 'go help modules' for details. If module-aware mode is enabled, "go run" runs in the context of the main module.

​	如果包参数没有版本后缀，"go run"可以在模块感知模式或GOPATH模式下运行，具体取决于GO111MODULE环境变量和go.mod文件的存在性。有关详细信息，请参见"go help modules"。如果启用了模块感知模式，"go run"将在主模块的上下文中运行。

By default, 'go run' runs the compiled binary directly: 'a.out arguments...'. If the -exec flag is given, 'go run' invokes the binary using xprog:

​	默认情况下，"go run"直接运行已编译的二进制文件："a.out arguments..."。如果给出了"-exec"标志，"go run"将使用xprog调用二进制文件：

```
'xprog a.out arguments...'.
```

If the -exec flag is not given, GOOS or GOARCH is different from the system default, and a program named go_$GOOS_$GOARCH_exec can be found on the current search path, 'go run' invokes the binary using that program, for example 'go_js_wasm_exec a.out arguments...'. This allows execution of cross-compiled programs when a simulator or other execution method is available.

​	如果未给出"-exec"标志，且GOOS或GOARCH与系统默认值不同，并且在当前搜索路径中可以找到一个名为"`go_$GOOS_$GOARCH_exec`"的程序，例如"go_js_wasm_exec a.out arguments..."，则"go run"将使用该程序调用二进制文件。这使得可以使用模拟器或其他执行方法来执行跨编译程序。

By default, 'go run' compiles the binary without generating the information used by debuggers, to reduce build time. To include debugger information in the binary, use 'go build'.

​	默认情况下，"go run"编译二进制文件时不会生成调试器使用的信息，以减少构建时间。要在二进制文件中包含调试器信息，请使用"go build"。

The exit status of Run is not the exit status of the compiled binary.

​	run的退出状态不是已编译的二进制文件的退出状态。

For more about build flags, see 'go help build'. For more about specifying packages, see 'go help packages'.

​	有关构建标志的更多信息，请参见"go help build"。有关指定包的更多信息，请参见"go help packages"。

See also: go build.

​	另请参见：go build。

#### go test -> 测试包 

Usage:

​	用法：

```
go test [build/test flags] [packages] [build/test flags & test binary flags]
```

'Go test' automates testing the packages named by the import paths. It prints a summary of the test results in the format:

​	`go test` 自动化测试由导入路径命名的包。它按以下格式打印测试结果的摘要：

```
ok   archive/tar   0.011s
FAIL archive/zip   0.022s
ok   compress/gzip 0.033s
...
```

followed by detailed output for each failed package.

接着是每个失败包的详细输出。

'Go test' recompiles each package along with any files with names matching the file pattern "`*_test.go`". These additional files can contain test functions, benchmark functions, fuzz tests and example functions. See 'go help testfunc' for more. Each listed package causes the execution of a separate test binary. Files whose names begin with "`_`" (including "`_test.go`") or "`.`" are ignored.

​	`go test`会重新编译每个包，以及与文件模式 "`*_test.go`" 匹配的任何文件。这些附加文件可以包含测试函数、基准函数、模糊测试和示例函数。有关更多信息，请参阅 'go help testfunc'。每个列出的包都会导致执行一个单独的测试二进制文件。以 "`_`"（包括 "`_test.go`"）或 "`.`" 开头的文件会被忽略。

Test files that declare a package with the suffix "_test" will be compiled as a separate package, and then linked and run with the main test binary.

​	声明带有后缀 "`_test`" 的包的测试文件将被编译为一个单独的包，然后与主测试二进制链接并运行。

The go tool will ignore a directory named "testdata", making it available to hold ancillary data needed by the tests.

​	go 工具会忽略名为 "testdata" 的目录，以用于保存测试所需的辅助数据。

As part of building a test binary, go test runs go vet on the package and its test source files to identify significant problems. If go vet finds any problems, go test reports those and does not run the test binary. Only a high-confidence subset of the default go vet checks are used. That subset is: atomic, bool, buildtags, directive, errorsas, ifaceassert, nilfunc, printf, and stringintconv. You can see the documentation for these and other vet tests via "go doc cmd/vet". To disable the running of go vet, use the -vet=off flag. To run all checks, use the -vet=all flag.

​	作为构建测试二进制的一部分，`go test` 还会在包及其测试源文件上运行 `go vet`，以识别重大的问题。如果 `go vet` 发现任何问题，`go test` 会报告这些问题，并且不会运行测试二进制。仅使用默认 go vet 检查的高置信度子集。该子集为：'atomic'、'bool'、'buildtags'、'errorsas'、'ifaceassert'、'nilfunc'、'printf' 和 'stringintconv'。您可以通过 "`go doc cmd/vet`" 查看这些和其他 vet 测试的文档。要禁用运行 `go vet`，使用 `-vet=off` 标志。要运行所有检查，使用 `-vet=all` 标志。

All test output and summary lines are printed to the go command's standard output, even if the test printed them to its own standard error. (The go command's standard error is reserved for printing errors building the tests.)

​	所有测试输出和摘要行都会打印到 go 命令的标准输出，即使测试将它们打印到自己的标准错误。 （go 命令的标准错误专用于打印构建测试时的错误。）

The go command places $GOROOT/bin at the beginning of $PATH in the test's environment, so that tests that execute 'go' commands use the same 'go' as the parent 'go test' command.

​	go 命令将 `$GOROOT/bin` 放在测试环境的 `$PATH` 开头，以便执行 `go` 命令的测试使用与父 `go test` 命令相同的 `go`。

Go test runs in two different modes:

​	`go test` 以两种不同的模式运行：

The first, called local directory mode, occurs when go test is invoked with no package arguments (for example, 'go test' or 'go test -v'). In this mode, go test compiles the package sources and tests found in the current directory and then runs the resulting test binary. In this mode, caching (discussed below) is disabled. After the package test finishes, go test prints a summary line showing the test status ('ok' or 'FAIL'), package name, and elapsed time.

​	第一种是本地目录模式，当无包参数调用 `go test` 时发生（例如 `go test` 或 `go test -v`）。在此模式下，`go test` 编译当前目录中找到的包源代码和测试，然后运行生成的测试二进制文件。在此模式下，缓存（下面讨论）被禁用。在包测试完成后，`go test` 打印一个摘要行，显示测试状态（'ok' 或 'FAIL'）、包名称和经过的时间。

The second, called package list mode, occurs when go test is invoked with explicit package arguments (for example 'go test math', 'go test ./...', and even 'go test .'). In this mode, go test compiles and tests each of the packages listed on the command line. If a package test passes, go test prints only the final 'ok' summary line. If a package test fails, go test prints the full test output. If invoked with the -bench or -v flag, go test prints the full output even for passing package tests, in order to display the requested benchmark results or verbose logging. After the package tests for all of the listed packages finish, and their output is printed, go test prints a final 'FAIL' status if any package test has failed.

​	第二种是包列表模式，在具有显式包参数的情况下调用 `go test` 时发生（例如 `go test math`、`go test ./...`，甚至 `go test .`）。在此模式下，go test 编译并测试命令行上列出的每个包。如果包测试通过，`go test` 只打印最终的 'ok' 摘要行。如果包测试失败，`go test` 打印完整的测试输出。如果使用了 `-bench` 或 `-v` 标志调用，`go test` 会为通过的包测试打印完整输出，以显示请求的基准结果或详细日志记录。在所有列出的包的包测试完成并打印出其输出后，如果任何包测试失败，`go test` 会打印最终的 'FAIL' 状态。

In package list mode only, go test caches successful package test results to avoid unnecessary repeated running of tests. When the result of a test can be recovered from the cache, go test will redisplay the previous output instead of running the test binary again. When this happens, go test prints '(cached)' in place of the elapsed time in the summary line.

​	仅在包列表模式下，`go test` 会将成功的包测试结果缓存起来，以避免不必要地重复运行测试。当测试结果可以从缓存中恢复时，`go test` 会重新显示先前的输出，而不是再次运行测试二进制文件。当这种情况发生时，`go test` 在摘要行的经过时间位置打印 '(cached)'。

The rule for a match in the cache is that the run involves the same test binary and the flags on the command line come entirely from a restricted set of 'cacheable' test flags, defined as -benchtime, -cpu, -list, -parallel, -run, -short, -timeout, -failfast, and -v. If a run of go test has any test or non-test flags outside this set, the result is not cached. To disable test caching, use any test flag or argument other than the cacheable flags. The idiomatic way to disable test caching explicitly is to use -count=1. Tests that open files within the package's source root (usually $GOPATH) or that consult environment variables only match future runs in which the files and environment variables are unchanged. A cached test result is treated as executing in no time at all, so a successful package test result will be cached and reused regardless of -timeout setting.

​	在缓存中的匹配规则是，运行涉及相同的测试二进制和命令行上的标志完全来自 'cacheable' 测试标志的受限集，这些标志定义为 `-benchtime`、`-cpu`、`-list`、`-parallel`、`-run`、`-short`、`-timeout`、`-failfast` 和 `-v`。如果 go test 的运行具有任何测试或非测试标志超出此集合，结果不会被缓存。显式禁用测试缓存的惯用方法是使用 `-count=1`。在包的源根目录（通常为 `$GOPATH`）中打开文件或仅查询环境变量的测试与未来的运行中文件和环境变量不发生变化的运行匹配。缓存的测试结果被视为根本没有执行时间，因此成功的包测试结果将被缓存并在不管 `-timeout` 设置如何的情况下重用。

In addition to the build flags, the flags handled by 'go test' itself are:

​	除了构建标志外，`go test` 本身处理的标志还有：

```
-args
	Pass the remainder of the command line (everything after -args)
    to the test binary, uninterpreted and unchanged.
    Because this flag consumes the remainder of the command line,
    the package list (if present) must appear before this flag.
    将命令行的其余部分（-args 之后的所有内容）
    传递给测试二进制文件，未解释和未更改。
    因为此标志使用命令行的其余部分，
    所以包列表（如果存在）必须出现在此标志之前。

-c
	Compile the test binary to pkg.test in the current directory but do not run it
    (where pkg is the last element of the package's import path).
    The file name or target directory can be changed with the -o flag.
    将测试二进制文件编译为 pkg.test，但不运行它
    （其中 pkg 是包的导入路径的最后一个元素）。
    文件名可以使用 -o 标志更改。

-exec xprog
	Run the test binary using xprog. The behavior is the same as
    in 'go run'. See 'go help run' for details.
    使用 xprog 运行测试二进制文件。行为与 'go run' 中相同。
    有关详情，请参阅 'go help run'。

-json
	Convert test output to JSON suitable for automated processing.
    See 'go doc test2json' for the encoding details.
    将测试输出转换为适用于自动化处理的 JSON。
    有关编码细节，请参阅 'go doc test2json'。

-o file
	Compile the test binary to the named file.
    The test still runs (unless -c or -i is specified).
    If file ends in a slash or names an existing directory,
    the test is written to pkg.test in that directory.
    将测试二进制文件编译为指定的文件。
    测试仍然会运行（除非指定了 -c 或 -i）。
```

The test binary also accepts flags that control execution of the test; these flags are also accessible by 'go test'. See 'go help testflag' for details.

​	测试二进制文件还接受控制测试执行的标志；这些标志也可由 `go test` 访问。有关详情，请参阅 'go help testflag'。

For more about build flags, see 'go help build'. For more about specifying packages, see 'go help packages'.

​	有关构建标志的更多信息，请参阅 'go help build'。有关指定包的更多信息，请参阅 'go help packages'。

See also: go build, go vet.

​	另请参阅：go build、go vet。

#### go tool -> 运行指定的go工具

Usage:

​	用法：

```
go tool [-n] command [args...]
```

Tool runs the go tool command identified by the arguments. With no arguments it prints the list of known tools.

​	`tool`运行由参数标识的go工具命令。如果没有参数，则打印已知工具列表。

The -n flag causes tool to print the command that would be executed but not execute it.

​	`-n`标志导致tool打印将要执行的命令，但不执行它。

For more about each tool command, see '`go doc cmd/<command>`'.

​	有关每个工具命令的详细信息，请参见"`go doc cmd/<command>`"。

#### go version -> 打印Go版本 

Usage:

​	用法：

```
go version [-m] [-v] [file ...]
```

Version prints the build information for Go binary files.

​	version 打印Go二进制文件的构建信息。

Go version reports the Go version used to build each of the named files.

​	go version 报告用于构建每个命名文件的Go版本。

If no files are named on the command line, go version prints its own version information.

​	如果命令行中没有命名文件，则go version打印自己的版本信息。

If a directory is named, go version walks that directory, recursively, looking for recognized Go binaries and reporting their versions. By default, go version does not report unrecognized files found during a directory scan. The -v flag causes it to report unrecognized files.

​	如果指定了目录，则go version递归地遍历该目录，寻找已知的Go二进制文件并报告它们的版本。默认情况下，go version不会报告目录扫描期间发现的无法识别的文件。-v标志会导致其报告无法识别的文件。

The -m flag causes go version to print each file's embedded module version information, when available. In the output, the module information consists of multiple lines following the version line, each indented by a leading tab character.

​	`-m`标志使go version在可能的情况下打印每个文件的嵌入式模块版本信息。在输出中，模块信息由版本行以下的多行组成，每行前面都有一个制表符。

See also: go doc runtime/debug.BuildInfo.

​	另请参见：go doc runtime/debug.BuildInfo。

#### go vet -> 报告包中的可能错误 

Usage:

​	用法：

```
go vet [-C dir] [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]
```

Vet runs the Go vet command on the packages named by the import paths.

​	vet在由导入路径命名的包上运行go vet命令。

For more about vet and its flags, see 'go doc cmd/vet'. For more about specifying packages, see 'go help packages'. For a list of checkers and their flags, see 'go tool vet help'. For details of a specific checker such as 'printf', see 'go tool vet help printf'.

​	有关vet及其标志的更多信息，请参见"go doc cmd/vet"。有关指定包的更多信息，请参见"go help packages"。有关检查器及其标志的列表，请参见"go tool vet help"。有关特定检查器（例如'printf'）的详细信息，请参见"go tool vet help printf"。

The -vettool=prog flag selects a different analysis tool with alternative or additional checks. For example, the 'shadow' analyzer can be built and run using these commands:

​	`-vettool=prog`标志选择具有替代或附加检查的不同分析工具。例如，可以使用以下命令构建和运行"shadow"分析器：

```
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
go vet -vettool=$(which shadow)
```

The build flags supported by go vet are those that control package resolution and execution, such as -C, -n, -x, -v, -tags, and -toolexec. For more about these flags, see 'go help build'.

​	go vet支持的构建标志是控制包解析和执行的标志，例如-n、-x、-v、-tags和-toolexec。有关这些标志的更多信息，请参见"go help build"。

See also: go fmt, go fix.

​	另请参见：go fmt、go fix。

#### //go:build -> 构建约束条件

A build constraint, also known as a build tag, is a condition under which a file should be included in the package. Build constraints are given by a line comment that begins

​	构建约束条件（也称为构建标记）是决定一个文件是否应该被包含在包中的条件。构建约束条件是由以下行注释给出的：

```
//go:build
```

Constraints may appear in any kind of source file (not just Go), but they must appear near the top of the file, preceded only by blank lines and other line comments. These rules mean that in Go files a build constraint must appear before the package clause.

​	约束条件可以出现在任何类型的源文件（不仅仅是 Go 文件），但它们必须出现在文件的顶部附近，只有空行和其他行注释可以在它们之前。这些规则意味着在 Go 文件中，构建约束条件必须出现在包语句之前。

To distinguish build constraints from package documentation, a build constraint should be followed by a blank line.

​	为了区分构建约束条件和包文档，构建约束条件应该在空行后跟随一个空行。

A build constraint comment is evaluated as an expression containing build tags combined by ||, &&, and ! operators and parentheses. Operators have the same meaning as in Go.

​	构建约束条件注释将被视为一个表达式，该表达式由 ||、&& 和 ! 运算符以及括号组合而成的构建标记组成。这些运算符的含义与 Go 中相同。

For example, the following build constraint constrains a file to build when the "linux" and "386" constraints are satisfied, or when "darwin" is satisfied and "cgo" is not:

​	例如，以下构建约束条件约束了一个文件只有在满足"linux"和"386"约束条件，或者在满足"darwin"约束条件且不满足"cgo"约束条件时才能构建：

```
//go:build (linux && 386) || (darwin && !cgo)
```

It is an error for a file to have more than one //go:build line.

​	一个文件如果有多于一个 //go:build 行注释，会引发错误。

During a particular build, the following build tags are satisfied:

​	在特定的构建过程中，以下构建标记将被满足：

- the target operating system, as spelled by runtime.GOOS, set with the GOOS environment variable.

- 目标操作系统，由 runtime.GOOS 拼写，使用 GOOS 环境变量设置。 

- the target architecture, as spelled by runtime.GOARCH, set with the GOARCH environment variable.

- 目标架构，由 runtime.GOARCH 拼写，使用 GOARCH 环境变量设置。

- any architecture features, in the form GOARCH.feature (for example, "amd64.v2"), as detailed below.

-  任何架构特性，采用 GOARCH.feature 的形式（例如，"amd64.v2"），如下所述。

-  "unix", if GOOS is a Unix or Unix-like system.

-  如果 GOOS 是 Unix 或类 Unix 系统，则为"unix"。 

-  the compiler being used, either "gc" or "gccgo"

- 所使用的编译器，可以是"gc"或"gccgo"。 

- "cgo", if the cgo command is supported (see CGO_ENABLED in 'go help environment').

- 如果支持 cgo 命令（请参阅"go help environment"中的 CGO_ENABLED），则为"cgo"。 

- a term for each Go major release, through the current version: "go1.1" from Go version 1.1 onward, "go1.12" from Go 1.12, and so on.

- 每个 Go 主要版本的术语，直到当前版本为止："go1.1"从 Go 1.1 开始，"go1.12"从 Go 1.12 开始，以此类推。 

- any additional tags given by the -tags flag (see 'go help build').

- 通过 -tags 标记提供的任何其他标记（请参阅"go help build"）。

There are no separate build tags for beta or minor releases.

​	beta 版本或小版本没有单独的构建标记。

If a file's name, after stripping the extension and a possible _test suffix, matches any of the following patterns:

​	如果文件名在去掉扩展名和可能的 _test 后，匹配以下任一模式：

```
*_GOOS
*_GOARCH
*_GOOS_GOARCH
```

(example: source_windows_amd64.go) where GOOS and GOARCH represent any known operating system and architecture values respectively, then the file is considered to have an implicit build constraint requiring those terms (in addition to any explicit constraints in the file).

（例如：source_windows_amd64.go），其中 GOOS 和 GOARCH 分别表示任何已知的操作系统和架构值，则该文件被认为具有需要这些术语的隐式构建约束条件（除了文件中的任何显式约束条件）。

Using GOOS=android matches build tags and files as for GOOS=linux in addition to android tags and files.

​	使用GOOS=android与GOOS=linux相匹配，以及android标记和文件。

Using GOOS=illumos matches build tags and files as for GOOS=solaris in addition to illumos tags and files.

​	使用GOOS=illumos与GOOS=solaris相匹配，并在illumos标记和文件上进行了补充。

Using GOOS=ios matches build tags and files as for GOOS=darwin in addition to ios tags and files.

​	使用GOOS=ios与GOOS=darwin相匹配，并在ios标记和文件上进行了补充。

The defined architecture feature build tags are:

​	已定义的架构特性构建标签为：

- For GOARCH=386, GO386=387 and GO386=sse2 set the 386.387 and 386.sse2 build tags, respectively.
- 对于GOARCH=386，GO386=387和GO386=sse2，分别设置386.387和386.sse2构建标签。 
- For GOARCH=amd64, GOAMD64=v1, v2, and v3 correspond to the amd64.v1, amd64.v2, and amd64.v3 feature build tags.
- 对于GOARCH=amd64，GOAMD64=v1、v2和v3对应于amd64.v1、amd64.v2和amd64.v3特性构建标签。
- For GOARCH=arm, GOARM=5, 6, and 7 correspond to the arm.5, arm.6, and arm.7 feature build tags.
-  对于GOARCH=arm，GOARM=5、6和7对应于arm.5、arm.6和arm.7特性构建标签。 
-  For GOARCH=mips or mipsle, GOMIPS=hardfloat and softfloat correspond to the mips.hardfloat and mips.softfloat (or mipsle.hardfloat and mipsle.softfloat) feature build tags.
- 对于GOARCH=mips或mipsle，GOMIPS=hardfloat和softfloat对应于mips.hardfloat和mips.softfloat（或mipsle.hardfloat和mipsle.softfloat）特性构建标签。 
- For GOARCH=mips64 or mips64le, GOMIPS64=hardfloat and softfloat correspond to the mips64.hardfloat and mips64.softfloat (or mips64le.hardfloat and mips64le.softfloat) feature build tags.
- 对于GOARCH=mips64或mips64le，GOMIPS64=hardfloat和softfloat对应于mips64.hardfloat和mips64.softfloat（或mips64le.hardfloat和mips64le.softfloat）特性构建标签。 
- For GOARCH=ppc64 or ppc64le, GOPPC64=power8, power9, and power10 correspond to the ppc64.power8, ppc64.power9, and ppc64.power10 (or ppc64le.power8, ppc64le.power9, and ppc64le.power10) feature build tags.
- 对于GOARCH=ppc64或ppc64le，GOPPC64=power8、power9和power10对应于ppc64.power8、ppc64.power9和ppc64.power10（或ppc64le.power8、ppc64le.power9和ppc64le.power10）特性构建标签。
- For GOARCH=wasm, GOWASM=satconv and signext correspond to the wasm.satconv and wasm.signext feature build tags.
-  对于GOARCH=wasm，GOWASM=satconv和signext对应于wasm.satconv和wasm.signext特性构建标签。

For GOARCH=amd64, arm, ppc64, and ppc64le, a particular feature level sets the feature build tags for all previous levels as well. For example, GOAMD64=v2 sets the amd64.v1 and amd64.v2 feature flags. This ensures that code making use of v2 features continues to compile when, say, GOAMD64=v4 is introduced. Code handling the absence of a particular feature level should use a negation:

​	对于GOARCH=amd64、arm、ppc64和ppc64le，特定特性级别也会设置先前所有级别的特性构建标签。例如，GOAMD64=v2设置了amd64.v1和amd64.v2特性标志。这确保使用v2特性的代码在引入GOAMD64=v4时继续编译。处理特定特性级别的缺失的代码应使用否定：

```
//go:build !amd64.v2
```

To keep a file from being considered for any build:

​	为了使文件不被考虑进行任何构建：

```
//go:build ignore
```

(Any other unsatisfied word will work as well, but "ignore" is conventional.)

​	（任何其他未满足的单词都可以工作，但"ignore"是常规的。）

To build a file only when using cgo, and only on Linux and OS X:

​	仅在使用cgo时并且仅在Linux和OS X上构建文件：

```
//go:build cgo && (linux || darwin)
```

Such a file is usually paired with another file implementing the default functionality for other systems, which in this case would carry the constraint:

​	这样的文件通常与实现其他系统的默认功能的另一个文件配对，该文件将带有以下限制：

```
//go:build !(cgo && (linux || darwin))
```

Naming a file dns_windows.go will cause it to be included only when building the package for Windows; similarly, math_386.s will be included only when building the package for 32-bit x86.

​	命名为dns_windows.go的文件将导致仅在为Windows构建包时包含它；同样，命名为math_386.s的文件将仅在为32位x86构建包时包含它。

Go versions 1.16 and earlier used a different syntax for build constraints, with a "// +build" prefix. The gofmt command will add an equivalent //go:build constraint when encountering the older syntax.

​	Go版本1.16及更早版本使用不同的构建限制语法，使用"// +build"前缀。在遇到旧语法时，gofmt命令将添加等效的//go:build约束。

#### 构建模式 

The 'go build' and 'go install' commands take a -buildmode argument which indicates which kind of object file is to be built. Currently supported values are:

​	'go build' 和 'go install' 命令可以使用 -buildmode 参数指定要构建哪种类型的目标文件。目前支持的值为：

```
-buildmode=archive
	Build the listed non-main packages into .a files. Packages named
	main are ignored.
	将列出的非 main 包构建为 .a 文件。名为 main 的包将被忽略。

-buildmode=c-archive
	Build the listed main package, plus all packages it imports,
	into a C archive file. The only callable symbols will be those
	functions exported using a cgo //export comment. Requires
	exactly one main package to be listed.
	将列出的 main 包以及它导入的所有包构建为 C 归档文件。
	唯一可调用的符号将是那些使用 cgo 的 //export 注释导出的函数。
	需要列出一个 main 包。



-buildmode=c-shared
	Build the listed main package, plus all packages it imports,
	into a C shared library. The only callable symbols will
	be those functions exported using a cgo //export comment.
	Requires exactly one main package to be listed.
	将列出的 main 包以及它导入的所有包构建为 C 共享库。
	唯一可调用的符号将是那些使用 cgo 的 //export 注释导出的函数。
	需要列出一个 main 包。



-buildmode=default
	Listed main packages are built into executables and listed
	non-main packages are built into .a files (the default
	behavior).
	列出的 main 包构建为可执行文件，
	列出的非 main 包构建为 .a 文件（默认行为）。



-buildmode=shared
	Combine all the listed non-main packages into a single shared
	library that will be used when building with the -linkshared
	option. Packages named main are ignored.
	将所有列出的非 main 包组合成一个共享库，
	将在使用 -linkshared 选项进行构建时使用。
	名为 main 的包将被忽略。



-buildmode=exe
	Build the listed main packages and everything they import into
	executables. Packages not named main are ignored.
	将列出的 main 包及其导入的所有内容构建为可执行文件。
	名为 main 以外的包将被忽略。



-buildmode=pie
	Build the listed main packages and everything they import into
	position independent executables (PIE). Packages not named
	main are ignored.
	将列出的 main 包及其导入的所有内容构建为位置无关可执行文件（PIE）。
	名为 main 以外的包将被忽略。



-buildmode=plugin
	Build the listed main packages, plus all packages that they
	import, into a Go plugin. Packages not named main are ignored.
	将列出的 main 包以及它们导入的所有包构建为 Go 插件。
	名为 main 以外的包将被忽略。

```

On AIX, when linking a C program that uses a Go archive built with -buildmode=c-archive, you must pass -Wl,-bnoobjreorder to the C compiler.

​	在 AIX 上，当链接使用 -buildmode=c-archive 构建的 Go 归档文件的 C 程序时，必须向 C 编译器传递 -Wl,-bnoobjreorder。

#### 调用 Go 和 C 之间的交互

There are two different ways to call between Go and C/C++ code.

​	调用 Go 和 C/C++ 代码有两种不同的方法。

The first is the cgo tool, which is part of the Go distribution. For information on how to use it see the cgo documentation (go doc cmd/cgo).

​	第一种是 cgo 工具，它是 Go 发行版的一部分。有关如何使用它的信息，请参阅 cgo 文档(go doc cmd/cgo)。

The second is the SWIG program, which is a general tool for interfacing between languages. For information on SWIG see http://swig.org/. When running go build, any file with a .swig extension will be passed to SWIG. Any file with a .swigcxx extension will be passed to SWIG with the -c++ option.

​	第二种是 SWIG 程序，它是一种用于不同语言间接口的通用工具。关于 SWIG 的信息请参见 http://swig.org/。在运行 go build 时，任何具有 .swig 扩展名的文件都将被传递给 SWIG。任何具有 .swigcxx 扩展名的文件都将被传递给带有 -c++ 选项的 SWIG。

When either cgo or SWIG is used, go build will pass any .c, .m, .s, .S or .sx files to the C compiler, and any .cc, .cpp, .cxx files to the C++ compiler. The CC or CXX environment variables may be set to determine the C or C++ compiler, respectively, to use.

​	当使用 cgo 或 SWIG 时，go build 将任何 .c、.m、.s、.S 或 .sx 文件传递给 C 编译器，并将任何 .cc、.cpp、.cxx 文件传递给 C++ 编译器。CC 或 CXX 环境变量可以设置以确定要使用的 C 或 C++ 编译器。

#### 构建和测试缓存

The go command caches build outputs for reuse in future builds. The default location for cache data is a subdirectory named go-build in the standard user cache directory for the current operating system. Setting the GOCACHE environment variable overrides this default, and running 'go env GOCACHE' prints the current cache directory.

​	go 命令缓存构建输出以便将来重复使用。缓存数据的默认位置是当前操作系统的标准用户缓存目录中名为 go-build 的子目录。设置 GOCACHE 环境变量将覆盖此默认设置，并且运行 'go env GOCACHE' 将打印当前缓存目录。

The go command periodically deletes cached data that has not been used recently. Running 'go clean -cache' deletes all cached data.

​	go 命令定期删除未经常用的缓存数据。运行 'go clean -cache' 将删除所有缓存的数据。

The build cache correctly accounts for changes to Go source files, compilers, compiler options, and so on: cleaning the cache explicitly should not be necessary in typical use. However, the build cache does not detect changes to C libraries imported with cgo. If you have made changes to the C libraries on your system, you will need to clean the cache explicitly or else use the -a build flag (see 'go help build') to force rebuilding of packages that depend on the updated C libraries.

​	构建缓存正确计算 Go 源文件、编译器、编译器选项等更改：在典型用法中不需要显式清除缓存。但是，构建缓存不会检测使用 cgo 导入的 C 库的更改。如果您对系统上的 C 库进行了更改，则需要显式清除缓存，否则使用 -a 构建标志(请参阅 'go help build')强制重建依赖于更新的 C 库的包。

The go command also caches successful package test results. See 'go help test' for details. Running 'go clean -testcache' removes all cached test results (but not cached build results).

​	go 命令还会缓存成功的包测试结果。有关详细信息，请参见 'go help test'。运行 'go clean -testcache' 将删除所有缓存的测试结果(但不会删除缓存的构建结果)。

The go command also caches values used in fuzzing with 'go test -fuzz', specifically, values that expanded code coverage when passed to a fuzz function. These values are not used for regular building and testing, but they're stored in a subdirectory of the build cache. Running 'go clean -fuzzcache' removes all cached fuzzing values. This may make fuzzing less effective, temporarily.

​	go 命令还会缓存用于 'go test -fuzz' 进行模糊测试的值，特别是将扩展代码覆盖率的值传递给模糊函数。这些值不用于常规构建和测试，但是它们存储在构建缓存的子目录中。运行 'go clean -fuzzcache' 将删除所有缓存的模糊值。这可能会使模糊测试在短时间内变得不那么有效。

The GODEBUG environment variable can enable printing of debugging information about the state of the cache:

​	GODEBUG环境变量可以启用有关缓存状态的调试信息：

GODEBUG=gocacheverify=1 causes the go command to bypass the use of any cache entries and instead rebuild everything and check that the results match existing cache entries.

​	GODEBUG=gocacheverify=1会导致go命令绕过使用任何缓存条目，而是重新构建所有内容并检查结果是否与现有的缓存条目匹配。

GODEBUG=gocachehash=1 causes the go command to print the inputs for all of the content hashes it uses to construct cache lookup keys. The output is voluminous but can be useful for debugging the cache.

​	GODEBUG=gocachehash=1会导致go命令打印用于构建缓存查找键的所有内容哈希的输入。输出很多，但可用于调试缓存。

GODEBUG=gocachetest=1 causes the go command to print details of its decisions about whether to reuse a cached test result.

​	GODEBUG=gocachetest=1会导致go命令打印有关是否重用缓存测试结果的详细信息。

#### 环境变量 

The go command and the tools it invokes consult environment variables for configuration. If an environment variable is unset or empty, the go command uses a sensible default setting. To see the effective setting of the variable <NAME>, run 'go env <NAME>'. To change the default setting, run 'go env -w <NAME>=<VALUE>'. Defaults changed using 'go env -w' are recorded in a Go environment configuration file stored in the per-user configuration directory, as reported by os.UserConfigDir. The location of the configuration file can be changed by setting the environment variable GOENV, and 'go env GOENV' prints the effective location, but 'go env -w' cannot change the default location. See 'go help env' for details.

​	go命令及其调用的工具会查询环境变量以进行配置。如果环境变量未设置，则go命令将使用合理的默认设置。要查看变量`<NAME>`的有效设置，请运行'go env `<NAME>`'。要更改默认设置，请运行'`go env -w <NAME>=<VALUE>`'。使用'go env -w'更改的默认值会记录在一个Go环境配置文件中，该文件存储在每个用户配置目录中，该目录由os.UserConfigDir报告。配置文件的位置可以通过设置环境变量GOENV来更改，'go env GOENV'打印有效位置，但'go env -w'无法更改默认位置。有关详细信息，请参见'go help env'。

General-purpose environment variables:

​	通用环境变量：

```
GO111MODULE
	Controls whether the go command runs in module-aware mode or GOPATH mode.
	May be "off", "on", or "auto".
	See https://golang.org/ref/mod#mod-commands.
	控制go命令运行在模块感知模式还是GOPATH模式下。
	可能的取值为"off"，"on"或"auto"。
	详见https://golang.org/ref/mod#mod-commands。
	
GCCGO
	The gccgo command to run for 'go build -compiler=gccgo'.
	运行'go build -compiler=gccgo'所用的gccgo命令。
	
GOARCH
	The architecture, or processor, for which to compile code.
	Examples are amd64, 386, arm, ppc64.
	编译代码的架构或处理器。例如，amd64，386，arm，ppc64。
	
GOBIN
	The directory where 'go install' will install a command.
	'go install'命令安装命令的目录。
	
GOCACHE
	The directory where the go command will store cached
	information for reuse in future builds.
	go命令存储缓存信息以便于在未来的构建中重复使用的目录。
	
GOMODCACHE
	The directory where the go command will store downloaded modules.
	go命令存储已下载模块的目录。
	
GODEBUG
	Enable various debugging facilities. See https://go.dev/doc/godebug
	for details.
	启用各种调试功能。详见'go doc runtime'。
	
GOENV
	The location of the Go environment configuration file.
	Cannot be set using 'go env -w'.
	Setting GOENV=off in the environment disables the use of the
	default configuration file.
	Go环境配置文件的位置。
	不能使用'go env -w'设置。
	将环境变量GOENV设置为"off"禁用使用默认配置文件。
	
GOFLAGS
	A space-separated list of -flag=value settings to apply
	to go commands by default, when the given flag is known by
	the current command. Each entry must be a standalone flag.
	Because the entries are space-separated, flag values must
	not contain spaces. Flags listed on the command line
	are applied after this list and therefore override it.
	要默认应用的一系列空格分隔的-flag=value标志设置，
	当当前命令已知该标志时。
	每个条目必须是一个独立的标志。
	因为条目是以空格分隔的，所以标志值不能包含空格。
	在命令行中列出的标志将在此列表之后应用，因此会覆盖它。
	
GOINSECURE
	Comma-separated list of glob patterns (in the syntax of Go's path.Match)
	of module path prefixes that should always be fetched in an insecure
	manner. Only applies to dependencies that are being fetched directly.
	GOINSECURE does not disable checksum database validation. GOPRIVATE or
	GONOSUMDB may be used to achieve that.
	用逗号分隔的模块路径前缀的通配符模式（使用Go的path.Match语法），
	它们始终以不安全的方式获取。
	仅适用于直接获取的依赖项。
	GOINSECURE不会禁用校验和数据库验证。
	GOPRIVATE或GONOSUMDB可用于实现该功能。

GOOS
	The operating system for which to compile code.
	Examples are linux, darwin, windows, netbsd.
	编译代码的操作系统。例如，linux，darwin，windows，netbsd。

GOPATH
	Controls where various files are stored. See: 'go help gopath'.
	更多详见'go help gopath'。

GOPROXY
	URL of Go module proxy. See https://golang.org/ref/mod#environment-variables
	and https://golang.org/ref/mod#module-proxy for details.
	Go 模块代理的 URL。
	有关详细信息，请参见 
	https://golang.org/ref/mod#environment-variables 
	和 https://golang.org/ref/mod#module-proxy。
	
GOPRIVATE,GONOPROXY,GONOSUMDB
	Comma-separated list of glob patterns (in the syntax of Go's path.Match)
	of module path prefixes that should always be fetched directly
	or that should not be compared against the checksum database.
	See https://golang.org/ref/mod#private-modules.
	逗号分隔的模块路径前缀的通配符列表（使用 Go 的 path.Match 语法），
	它们应始终直接获取，或不应与校验和数据库进行比较。
	请参见 https://golang.org/ref/mod#private-modules。
	
GOROOT
	The root of the go tree.
	Go 树的根目录。
	
GOSUMDB
	The name of checksum database to use and optionally its public key and
	URL. See https://golang.org/ref/mod#authenticating.
	要使用的校验和数据库的名称及其公钥和 URL（可选）。
	请参见 https://golang.org/ref/mod#authenticating。
	
GOTOOLCHAIN
	Controls which Go toolchain is used. See https://go.dev/doc/toolchain.
	控制使用哪个Go工具链。请参阅 https://go.dev/doc/toolchain。
	
GOTMPDIR
	The directory where the go command will write
	temporary source files, packages, and binaries.
	Go 命令将写入临时源文件、包和二进制文件的目录。

GOVCS
	Lists version control commands that may be used with matching servers.
	See 'go help vcs'.
	列出可能与匹配服务器一起使用的版本控制命令。请参见 'go help vcs'。

GOWORK

	在模块感知模式下，使用给定的 go.work 文件作为工作空间文件。
	默认情况下或当 GOWORK 为 "auto" 时，
	go 命令会在当前目录及包含目录中搜索名为 go.work 的文件，
	直到找到为止。
	如果找到有效的 go.work 文件，则所指定的模块将共同用作主要模块。
	如果 GOWORK 为 "off"，或者在 "auto" 模式下未找到 go.work 文件，
	则禁用工作空间模式。
```

Environment variables for use with cgo:

​	用于与 cgo 一起使用的环境变量：

```
AR
	The command to use to manipulate library archives when
	building with the gccgo compiler.
	The default is 'ar'.
	在使用 gccgo 编译器进行构建时，用于操作库档案的命令。默认值为 'ar'。

CC
	The command to use to compile C code.
	用于编译 C 代码的命令。

CGO_ENABLED
	Whether the cgo command is supported. Either 0 or 1.
	是否支持 cgo 命令。为 0 或 1。
	
CGO_CFLAGS
	Flags that cgo will pass to the compiler when compiling
	C code.
	cgo 在编译 C 代码时将传递给编译器的标志。
	
CGO_CFLAGS_ALLOW
	A regular expression specifying additional flags to allow
	to appear in #cgo CFLAGS source code directives.
	Does not apply to the CGO_CFLAGS environment variable.
	指定允许在 #cgo CFLAGS 源代码指令中出现的其他标志的正则表达式。
	不适用于 CGO_CFLAGS 环境变量。
	
CGO_CFLAGS_DISALLOW
	A regular expression specifying flags that must be disallowed
	from appearing in #cgo CFLAGS source code directives.
	Does not apply to the CGO_CFLAGS environment variable.
	指定必须禁止出现在 #cgo CFLAGS 源代码指令中的标志的正则表达式。
	不适用于 CGO_CFLAGS 环境变量。

CGO_CPPFLAGS, CGO_CPPFLAGS_ALLOW, CGO_CPPFLAGS_DISALLOW
	Like CGO_CFLAGS, CGO_CFLAGS_ALLOW, and CGO_CFLAGS_DISALLOW,
	but for the C preprocessor.
	与 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW 类似，
	但用于 C 预处理器。
	
CGO_CXXFLAGS, CGO_CXXFLAGS_ALLOW, CGO_CXXFLAGS_DISALLOW
	Like CGO_CFLAGS, CGO_CFLAGS_ALLOW, and CGO_CFLAGS_DISALLOW,
	but for the C++ compiler.
	Like CGO_CFLAGS, CGO_CFLAGS_ALLOW, and CGO_CFLAGS_DISALLOW,
	but for the C++ compiler.
	与 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW 类似，
	但是用于 C++ 编译器。
	
CGO_FFLAGS, CGO_FFLAGS_ALLOW, CGO_FFLAGS_DISALLOW
	Like CGO_CFLAGS, CGO_CFLAGS_ALLOW, and CGO_CFLAGS_DISALLOW,
	but for the Fortran compiler.
	与 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW 类似，
	但是用于 Fortran 编译器。
	
CGO_LDFLAGS, CGO_LDFLAGS_ALLOW, CGO_LDFLAGS_DISALLOW
	Like CGO_CFLAGS, CGO_CFLAGS_ALLOW, and CGO_CFLAGS_DISALLOW,
	but for the linker.
	与 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW 类似，
	但是用于链接器。
	
CXX
	The command to use to compile C++ code.
	用于编译 C++ 代码的命令。
	
FC
	The command to use to compile Fortran code.
	用于编译 Fortran 代码的命令。
	
PKG_CONFIG
	Path to pkg-config tool.
	pkg-config 工具的路径。	
```

Architecture-specific environment variables:

​	特定架构的环境变量：

```
GOARM
	For GOARCH=arm, the ARM architecture for which to compile.
	Valid values are 5, 6, 7.
	用于 GOARCH=arm，表示编译的 ARM 架构。
	有效的值为 5、6、7。
	
GO386
	For GOARCH=386, how to implement floating point instructions.
	Valid values are sse2 (default), softfloat.
	用于 GOARCH=386，指定如何实现浮点指令。
	有效的值为 sse2（默认）和 softfloat。
	
GOAMD64
	For GOARCH=amd64, the microarchitecture level for which to compile.
	Valid values are v1 (default), v2, v3, v4.
	See https://golang.org/wiki/MinimumRequirements#amd64
	用于 GOARCH=amd64，表示编译的微架构级别。
	有效的值为 v1（默认）、v2、v3、v4。
	参见 https://golang.org/wiki/MinimumRequirements#amd64。
	
GOMIPS
	For GOARCH=mips{,le}, whether to use floating point instructions.
	Valid values are hardfloat (default), softfloat.
	用于 GOARCH=mips{,le}，表示是否使用浮点指令。
	有效的值为 hardfloat（默认）和 softfloat。
	
GOMIPS64
	For GOARCH=mips64{,le}, whether to use floating point instructions.
	Valid values are hardfloat (default), softfloat.
	用于 GOARCH=mips64{,le}，表示是否使用浮点指令。
	有效的值为 hardfloat（默认）和 softfloat。
	
GOPPC64
	For GOARCH=ppc64{,le}, the target ISA (Instruction Set Architecture).
	Valid values are power8 (default), power9, power10.
	用于 GOARCH=ppc64{,le}，表示目标 ISA（指令集架构）。
	有效的值为 power8（默认）、power9 和 power10。
	
GOWASM
	For GOARCH=wasm, comma-separated list of experimental WebAssembly features to use.
	Valid values are satconv, signext.
	用于 GOARCH=wasm，以逗号分隔的实验性 WebAssembly 功能列表。
	有效的值为 satconv 和 signext。
```

Environment variables for use with code coverage:

​	用于代码覆盖的环境变量：

```
GOCOVERDIR
	Directory into which to write code coverage data files
	generated by running a "go build -cover" binary.
	Requires that GOEXPERIMENT=coverageredesign is enabled.
	"go build -cover" 二进制文件生成的代码覆盖数据文件将写入其中的目录。
	需要启用 GOEXPERIMENT=coverageredesign。
```

Special-purpose environment variables:

​	特殊用途的环境变量：

```
GCCGOTOOLDIR
	If set, where to find gccgo tools, such as cgo.
	The default is based on how gccgo was configured.
	如果设置了，gccgo 工具的所在位置，例如 cgo。
	默认值基于 gccgo 的配置。
	
GOEXPERIMENT
	Comma-separated list of toolchain experiments to enable or disable.
	The list of available experiments may change arbitrarily over time.
	See src/internal/goexperiment/flags.go for currently valid values.
	Warning: This variable is provided for the development and testing
	of the Go toolchain itself. Use beyond that purpose is unsupported.
	启用或禁用的工具链实验的逗号分隔列表。
	可用实验的列表可能随时更改。
	有关目前有效值，请参见 src/internal/goexperiment/flags.go。
	警告：此变量仅为开发和测试 Go 工具链本身提供。
	超出此目的的使用不受支持。
	
GOROOT_FINAL
	The root of the installed Go tree, when it is
	installed in a location other than where it is built.
	File names in stack traces are rewritten from GOROOT to
	GOROOT_FINAL.
	安装 Go 树的根目录，当它安装在与构建位置不同的位置时。
	堆栈跟踪中的文件名将从 GOROOT 重写为 GOROOT_FINAL。
	
GO_EXTLINK_ENABLED
	Whether the linker should use external linking mode
	when using -linkmode=auto with code that uses cgo.
	Set to 0 to disable external linking mode, 1 to enable it.
	使用 -linkmode=auto 与使用 cgo 的代码时，
	链接器是否应使用外部链接模式。
	将其设置为 0 以禁用外部链接模式，设置为 1 以启用。
	
GIT_ALLOW_PROTOCOL
	Defined by Git. A colon-separated list of schemes that are allowed
	to be used with git fetch/clone. If set, any scheme not explicitly
	mentioned will be considered insecure by 'go get'.
	Because the variable is defined by Git, the default value cannot
	be set using 'go env -w'.
	Git 定义的允许使用 git fetch/clone 的方案的冒号分隔列表。
	如果设置，未明确提及的任何方案都将被 'go get' 视为不安全。
	因为该变量由 Git 定义，所以不能使用 'go env -w' 设置默认值。
```

Additional information available from 'go env' but not read from the environment:

​	从 'go env' 中获取的其他信息，但不是从环境中读取的：

```
GOEXE
	The executable file name suffix (".exe" on Windows, "" on other systems).
	可执行文件名后缀（Windows 上为".exe"，其他系统上为空字符串）。

GOGCCFLAGS
	A space-separated list of arguments supplied to the CC command.
	CC 命令提供的参数的以空格分隔的列表。

GOHOSTARCH
	The architecture (GOARCH) of the Go toolchain binaries.
	Go 工具链二进制文件的架构（GOARCH）。

GOHOSTOS
	The operating system (GOOS) of the Go toolchain binaries.
	Go 工具链二进制文件的操作系统（GOOS）。

GOMOD
	The absolute path to the go.mod of the main module.
	If module-aware mode is enabled, but there is no go.mod, GOMOD will be
	os.DevNull ("/dev/null" on Unix-like systems, "NUL" on Windows).
	If module-aware mode is disabled, GOMOD will be the empty string.
	主模块的 go.mod 的绝对路径。
	如果启用了模块感知模式，
	但没有 go.mod，则 GOMOD 将是 os.DevNull
	（在类 Unix 系统上为"/dev/null"，在 Windows 上为"NUL"）。
	如果禁用了模块感知模式，则 GOMOD 将为空字符串。
	
GOTOOLDIR
	The directory where the go tools (compile, cover, doc, etc...) are installed.
	Go 工具（compile、cover、doc 等）所在的目录。
	
GOVERSION
	The version of the installed Go tree, as reported by runtime.Version.
	安装的 Go 树的版本，由 runtime.Version 报告。
```

#### 文件类型

The go command examines the contents of a restricted set of files in each directory. It identifies which files to examine based on the extension of the file name. These extensions are:

​	go 命令检查每个目录中一组受限制的文件的内容。 它根据文件名的扩展名来确定要检查哪些文件。 这些扩展名是：

```
.go
	Go source files.
	Go源文件。
	
.c, .h
	C source files.
	If the package uses cgo or SWIG, these will be compiled with the
	OS-native compiler (typically gcc); otherwise they will
	trigger an error.
	C源文件。
	如果包使用cgo或SWIG，
	则这些文件将使用本机编译器（通常是gcc）编译；
	否则它们会导致错误。
	
.cc, .cpp, .cxx, .hh, .hpp, .hxx
	C++ source files. Only useful with cgo or SWIG, and always
	compiled with the OS-native compiler.
	C++源文件。仅在使用cgo或SWIG时有用，并且始终使用本机编译器编译。
	
.m
	Objective-C source files. Only useful with cgo, and always
	compiled with the OS-native compiler.
	Objective-C源文件。仅在使用cgo时有用，并始终使用本机编译器编译。
	
.s, .S, .sx
	Assembler source files.
	If the package uses cgo or SWIG, these will be assembled with the
	OS-native assembler (typically gcc (sic)); otherwise they
	will be assembled with the Go assembler.
	汇编器源文件。
	如果包使用cgo或SWIG，则这些文件将使用本机汇编器（通常是gcc）汇编；
	否则它们将使用Go汇编器汇编。
	
.swig, .swigcxx
	SWIG definition files.
	SWIG定义文件。
	
.syso
	System object files.
	系统对象文件。	
```

Files of each of these types except .syso may contain build constraints, but the go command stops scanning for build constraints at the first item in the file that is not a blank line or //-style line comment. See the go/build package documentation for more details.

​	除了.syso之外，每个这些类型的文件都可以包含构建约束条件，但是go命令会在第一个不是空白行或`//-`样式行注释的条目处停止扫描构建约束条件。有关更多详细信息，请参阅go/build包文档。

#### go.mod文件

A module version is defined by a tree of source files, with a go.mod file in its root. When the go command is run, it looks in the current directory and then successive parent directories to find the go.mod marking the root of the main (current) module.

​	模块版本由源文件树定义，在其根目录中有一个go.mod文件。运行go命令时，它会在当前目录中查找，然后在父目录中查找，以找到标记当前模块根的go.mod文件。

The go.mod file format is described in detail at https://golang.org/ref/mod#go-mod-file.

​	go.mod文件格式在https://golang.org/ref/mod#go-mod-file中详细描述。

To create a new go.mod file, use 'go mod init'. For details see 'go help mod init' or https://golang.org/ref/mod#go-mod-init.

​	要创建新的go.mod文件，请使用'go mod init'。有关详细信息，请参见'go help mod init'或https://golang.org/ref/mod#go-mod-init。

To add missing module requirements or remove unneeded requirements, use 'go mod tidy'. For details, see 'go help mod tidy' or https://golang.org/ref/mod#go-mod-tidy.

​	要添加缺少的模块要求或删除不需要的要求，请使用'go mod tidy'。有关详细信息，请参见'go help mod tidy'或https://golang.org/ref/mod#go-mod-tidy。

To add, upgrade, downgrade, or remove a specific module requirement, use 'go get'. For details, see 'go help module-get' or https://golang.org/ref/mod#go-get.

​	要添加、升级、降级或删除特定模块要求，请使用'go get'。有关详细信息，请参见'go help module-get'或https://golang.org/ref/mod#go-get。

To make other changes or to parse go.mod as JSON for use by other tools, use 'go mod edit'. See 'go help mod edit' or https://golang.org/ref/mod#go-mod-edit.

​	要进行其他更改或将go.mod解析为JSON供其他工具使用，请使用'go mod edit'。请参阅'go help mod edit'或https://golang.org/ref/mod#go-mod-edit。

#### GOPATH环境变量

The Go path is used to resolve import statements. It is implemented by and documented in the go/build package.

​	Go路径用于解析导入语句。它由go/build包实现并记录。

The GOPATH environment variable lists places to look for Go code. On Unix, the value is a colon-separated string. On Windows, the value is a semicolon-separated string. On Plan 9, the value is a list.

​	GOPATH环境变量列出了查找Go代码的位置。在Unix上，该值是以冒号分隔的字符串。在Windows上，该值是以分号分隔的字符串。在Plan 9上，该值是一个列表。

If the environment variable is unset, GOPATH defaults to a subdirectory named "go" in the user's home directory ($HOME/go on Unix, %USERPROFILE%\go on Windows), unless that directory holds a Go distribution. Run "go env GOPATH" to see the current GOPATH.

​	如果环境变量未设置，GOPATH默认为用户主目录中名为"go"的子目录（在Unix上为`$HOME/go`，在Windows上为`%USERPROFILE%\go`），除非该目录包含Go发行版。运行"go env GOPATH"可以查看当前GOPATH。

See https://golang.org/wiki/SettingGOPATH to set a custom GOPATH.

​	请参见https://golang.org/wiki/SettingGOPATH来设置自定义GOPATH。

Each directory listed in GOPATH must have a prescribed structure:

​	GOPATH中列出的每个目录都必须具有预定的结构：

- The src directory holds source code. The path below src determines the import path or executable name.

- src目录包含源代码。src下面的路径确定导入路径或可执行文件名。

- The pkg directory holds installed package objects. As in the Go tree, each target operating system and architecture pair has its own subdirectory of pkg (pkg/GOOS_GOARCH).

- pkg目录保存已安装的包对象。与Go树中一样，每个目标操作系统和架构对都有自己的pkg子目录（pkg/GOOS_GOARCH）。

​	If DIR is a directory listed in the GOPATH, a package with source in DIR/src/foo/bar can be imported as "foo/bar" and has its compiled form installed to "DIR/pkg/GOOS_GOARCH/foo/bar.a".	

​	如果DIR是在GOPATH中列出的目录，则具有DIR/src/foo/bar中的源代码的包可以作为"foo/bar"导入，并将其编译形式安装到"DIR/pkg/GOOS_GOARCH/foo/bar.a"中。

- The bin directory holds compiled commands. Each command is named for its source directory, but only the final element, not the entire path. That is, the command with source in DIR/src/foo/quux is installed into DIR/bin/quux, not DIR/bin/foo/quux. The "foo/" prefix is stripped so that you can add DIR/bin to your PATH to get at the installed commands. If the GOBIN environment variable is set, commands are installed to the directory it names instead of DIR/bin. GOBIN must be an absolute path.
- bin目录保存已编译的命令。每个命令以其源目录命名，但只使用最后一个元素，而不是整个路径。也就是说，具有DIR/src/foo/quux源代码的命令将安装到DIR/bin/quux，而不是DIR/bin/foo/quux。前缀"foo/"被删除，以便您可以将DIR/bin添加到PATH中以获得已安装命令。如果设置了GOBIN环境变量，则命令将安装到其命名的目录，而不是DIR/bin。GOBIN必须是绝对路径。

Here's an example directory layout:

​	这是一个示例目录布局：

```
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
```

Go searches each directory listed in GOPATH to find source code, but new packages are always downloaded into the first directory in the list.

​	Go会在GOPATH列表中的每个目录中查找源代码，但新的包总是被下载到列表中第一个目录。

See https://golang.org/doc/code.html for an example.

​	参见 https://golang.org/doc/code.html 获取示例。

#### GOPATH 和 Modules

When using modules, GOPATH is no longer used for resolving imports. However, it is still used to store downloaded source code (in GOPATH/pkg/mod) and compiled commands (in GOPATH/bin).

​	使用 modules 时，GOPATH 不再用于解析导入。但是，它仍然用于存储已下载的源代码（在 GOPATH/pkg/mod 中）和编译的命令（在 GOPATH/bin 中）。

#### 内部目录

Code in or below a directory named "internal" is importable only by code in the directory tree rooted at the parent of "internal". Here's an extended version of the directory layout above:

​	位于名为"internal"的目录中或以下的代码只能被根目录为"internal"的目录树中的代码导入。下面是上述目录布局的扩展版本：

```
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
```

The code in z.go is imported as "foo/internal/baz", but that import statement can only appear in source files in the subtree rooted at foo. The source files foo/f.go, foo/bar/x.go, and foo/quux/y.go can all import "foo/internal/baz", but the source file crash/bang/b.go cannot.

​	在 z.go 中导入为"foo/internal/baz"，但是该导入语句只能出现在以 foo 为根的子树中的源文件中。foo/f.go、foo/bar/x.go 和 foo/quux/y.go 中的源文件都可以导入"foo/internal/baz"，但是 crash/bang/b.go 中的源文件不能。

See https://golang.org/s/go14internal for details.

​	有关详细信息，请参见 https://golang.org/s/go14internal。

#### vendor目录

Go 1.6 includes support for using local copies of external dependencies to satisfy imports of those dependencies, often referred to as vendoring.

​	Go 1.6 包括使用本地副本来满足这些依赖项的导入的支持，通常称为供应商。

Code below a directory named "vendor" is importable only by code in the directory tree rooted at the parent of "vendor", and only using an import path that omits the prefix up to and including the vendor element.

​	位于名为"vendor"的目录以下的代码只能由根目录为"vendor"的目录树中的代码导入，并且仅使用省略前缀直到包含 vendor 元素的路径。

Here's the example from the previous section, but with the "internal" directory renamed to "vendor" and a new foo/vendor/crash/bang directory added:

​	下面是前一节的示例，但将"internal"目录重命名为"vendor"并添加了一个新的 foo/vendor/crash/bang 目录：

```
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
```

The same visibility rules apply as for internal, but the code in z.go is imported as "baz", not as "foo/vendor/baz".

​	同样的可见性规则适用于 internal，但是 z.go 中的代码导入时使用 "baz"，而不是 "foo/vendor/baz"。

Code in vendor directories deeper in the source tree shadows code in higher directories. Within the subtree rooted at foo, an import of "crash/bang" resolves to "foo/vendor/crash/bang", not the top-level "crash/bang".

​	位于更深层次的 vendor 目录中的代码会遮蔽较高层次的目录中的代码。在以 foo 为根的子树中，对 "crash/bang" 的导入会解析为 "foo/vendor/crash/bang"，而不是顶层的 "crash/bang"。

Code in vendor directories is not subject to import path checking (see 'go help importpath').

​	位于 vendor 目录中的代码不受导入路径检查的限制（参见 'go help importpath'）。

When 'go get' checks out or updates a git repository, it now also updates submodules.

​	当 'go get' 检出或更新 git 子模块时，它现在也会更新子模块。

Vendor directories do not affect the placement of new repositories being checked out for the first time by 'go get': those are always placed in the main GOPATH, never in a vendor subtree.

​	vendor 目录不影响由 'go get' 第一次检出的新代码库的位置：它们总是被放置在主 GOPATH 中，而不是在 vendor 子目录中。

See https://golang.org/s/go15vendor for details.

​	详见 https://golang.org/s/go15vendor。

#### 遗留 GOPATH go get  Legacy GOPATH go get

The 'go get' command changes behavior depending on whether the go command is running in module-aware mode or legacy GOPATH mode. This help text, accessible as 'go help gopath-get' even in module-aware mode, describes 'go get' as it operates in legacy GOPATH mode.

​	'go get' 命令的行为取决于 go 命令是在模块感知模式还是遗留 GOPATH 模式下运行。即使在模块感知模式下，也可以通过 'go help gopath-get' 访问此帮助文本，它描述了 'go get' 在遗留 GOPATH 模式下的操作。

Usage:

​	用法：

```
 go get [-d] [-f] [-t] [-u] [-v] [-fix] [build flags] [packages]
```

Get downloads the packages named by the import paths, along with their dependencies. It then installs the named packages, like 'go install'.

​	get 下载指定导入路径的包及其依赖项。它然后安装指定的包，就像 'go install' 一样。

The -d flag instructs get to stop after downloading the packages; that is, it instructs get not to install the packages.

​	`-d` 标志指示 get 在下载包后停止，即指示 get 不安装包。

The -f flag, valid only when -u is set, forces get -u not to verify that each package has been checked out from the source control repository implied by its import path. This can be useful if the source is a local fork of the original.

​	`-f` 标志只在设置了 -u 标志时有效，强制 get -u 不验证每个包是否已从其导入路径所暗示的源代码库检出。如果源代码是原始代码的本地分支，则此功能可能很有用。

The -fix flag instructs get to run the fix tool on the downloaded packages before resolving dependencies or building the code.

​	`-fix` 标志指示 get 在解析依赖项或构建代码之前对下载的包运行 fix 工具。

The -t flag instructs get to also download the packages required to build the tests for the specified packages.

​	`-t` 标志指示 get 也下载构建指定包所需的测试包。

The -u flag instructs get to use the network to update the named packages and their dependencies. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages.

​	`-u` 标志指示 get 使用网络更新指定的包及其依赖项。默认情况下，get 使用网络检出缺少的包，但不使用网络查找现有包的更新。

The -v flag enables verbose progress and debug output.

​	`-v` 标志启用详细进度和调试输出。

Get also accepts build flags to control the installation. See 'go help build'.

​	get 命令还接受构建标志以控制安装。请参阅 'go help build'。

When checking out a new package, get creates the target directory `GOPATH/src/<import-path>`. If the GOPATH contains multiple entries, get uses the first one. For more details see: 'go help gopath'.

​	在检出新包时，get 会创建目标目录 `GOPATH/src/<import-path>`。如果 GOPATH 包含多个条目，则 get 会使用第一个条目。有关更多详细信息，请参阅 'go help gopath'。

When checking out or updating a package, get looks for a branch or tag that matches the locally installed version of Go. The most important rule is that if the local installation is running version "go1", get searches for a branch or tag named "go1". If no such version exists it retrieves the default branch of the package.

​	在检出或更新包时，get 会查找与本地安装的 Go 版本匹配的分支或标记。最重要的规则是，如果本地安装正在运行版本 "go1"，则 get 会搜索名为 "go1" 的分支或标记。如果不存在此类版本，则检索包的默认分支。

When go get checks out or updates a Git repository, it also updates any git submodules referenced by the repository.

​	当 go get 检出或更新 Git 存储库时，它也会更新由存储库引用的任何 git 子模块。

Get never checks out or updates code stored in vendor directories.

​	get 永远不会检出或更新存储在 vendor 目录中的代码。

For more about build flags, see 'go help build'.

​	有关构建标志的更多信息，请参阅 'go help build'。

For more about specifying packages, see 'go help packages'.

​	有关指定包的更多信息，请参阅 'go help packages'。

For more about how 'go get' finds source code to download, see 'go help importpath'.

​	有关 'go get' 如何查找要下载的源代码的更多信息，请参阅 'go help importpath'。

This text describes the behavior of get when using GOPATH to manage source code and dependencies. If instead the go command is running in module-aware mode, the details of get's flags and effects change, as does 'go help get'. See 'go help modules' and 'go help module-get'.

​	本文描述了使用 GOPATH 管理源代码和依赖项时的 get 行为。如果 go 命令是在模块感知模式下运行，则 get 的标志和效果的详细信息会发生更改，'go help get' 也会发生更改。请参阅 'go help modules' 和 'go help module-get'。

See also: go build, go install, go clean.

​	另请参阅：go build、go install、go clean。

#### 模块代理协议

A Go module proxy is any web server that can respond to GET requests for URLs of a specified form. The requests have no query parameters, so even a site serving from a fixed file system (including a file:/// URL) can be a module proxy.

​	Go 模块代理是任何可以响应特定格式 URL 的 GET 请求的 Web 服务器。请求没有查询参数，因此即使是从固定文件系统（包括 `file://URL`）服务的站点也可以是模块代理。

For details on the GOPROXY protocol, see https://golang.org/ref/mod#goproxy-protocol.

​	有关 GOPROXY 协议的详细信息，请参阅 https://golang.org/ref/mod#goproxy-protocol。

#### 导入路径语法

An import path (see 'go help packages') denotes a package stored in the local file system. In general, an import path denotes either a standard package (such as "unicode/utf8") or a package found in one of the work spaces (For more details see: 'go help gopath').

​	导入路径（请参阅 'go help packages'）表示存储在本地文件系统中的包。一般来说，导入路径表示标准包（如 "unicode/utf8"）或在工作区之一找到的包（有关详细信息，请参阅：'go help gopath'）。

#### 相对导入路径

An import path beginning with ./ or ../ is called a relative path. The toolchain supports relative import paths as a shortcut in two ways.

​	以 `./` 或 `../` 开头的导入路径称为相对路径。工具链支持相对导入路径作为两种快捷方式。

First, a relative path can be used as a shorthand on the command line. If you are working in the directory containing the code imported as "unicode" and want to run the tests for "unicode/utf8", you can type "go test ./utf8" instead of needing to specify the full path. Similarly, in the reverse situation, "go test .." will test "unicode" from the "unicode/utf8" directory. Relative patterns are also allowed, like "go test ./..." to test all subdirectories. See 'go help packages' for details on the pattern syntax.

​	首先，相对路径可以作为命令行上的快捷方式。如果您在包含代码作为 "unicode" 导入的目录中工作，并想要运行 "unicode/utf8" 的测试，则可以键入 "go test ./utf8" 而不需要指定完整路径。同样，在反向情况下，"go test .."将从 "unicode/utf8" 目录测试 "unicode"。相对模式也被允许，例如 "go test ./..." 可以测试所有子目录。有关模式语法的详细信息，请参见 'go help packages'。

Second, if you are compiling a Go program not in a work space, you can use a relative path in an import statement in that program to refer to nearby code also not in a work space. This makes it easy to experiment with small multipackage programs outside of the usual work spaces, but such programs cannot be installed with "go install" (there is no work space in which to install them), so they are rebuilt from scratch each time they are built. To avoid ambiguity, Go programs cannot use relative import paths within a work space.

​	其次，如果您在不在工作空间中编译 Go 程序，那么您可以在该程序的导入语句中使用相对路径，以引用附近的代码，也不在工作空间中。这使得在通常的工作空间之外实验小型多包程序非常容易，但这些程序无法使用 "go install" 进行安装（没有工作空间来安装它们），因此每次构建时都会重新构建。为避免歧义，Go 程序不能在工作空间内使用相对导入路径。

#### 远程导入路径 

Certain import paths also describe how to obtain the source code for the package using a revision control system.

​	某些导入路径也描述了如何使用版本控制系统获取包的源代码。

A few common code hosting sites have special syntax:

​	一些常见的代码托管站点有特殊的语法：

```
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
```

For code hosted on other servers, import paths may either be qualified with the version control type, or the go tool can dynamically fetch the import path over https/http and discover where the code resides from a <meta> tag in the HTML.

​	对于在其他服务器上托管的代码，导入路径可以使用版本控制类型进行限定，或者 go 工具可以通过 https/http 动态获取导入路径，并从 HTML 的 <meta> 标记中发现代码所在的位置。

To declare the code location, an import path of the form

​	要声明代码位置，形式为

```
repository.vcs/path
```

specifies the given repository, with or without the .vcs suffix, using the named version control system, and then the path inside that repository. The supported version control systems are:

的导入路径指定了给定的仓库，带或不带 `.vcs` 后缀，使用命名的版本控制系统，然后是该仓库内的路径。支持的版本控制系统是：

```
Bazaar      .bzr
Fossil      .fossil
Git         .git
Mercurial   .hg
Subversion  .svn
```

For example,

​	例如，

```
import "example.org/user/foo.hg"
```

denotes the root directory of the Mercurial repository at example.org/user/foo or foo.hg, and

表示在 example.org/user/foo 或 foo.hg 上的 Mercurial 存储库的根目录，而

```
import "example.org/repo.git/foo/bar"
```

denotes the foo/bar directory of the Git repository at example.org/repo or repo.git.

表示 example.org/repo 或 repo.git 上的 Git 存储库的 foo/bar 目录。

When a version control system supports multiple protocols, each is tried in turn when downloading. For example, a Git download tries https://, then git+ssh://.

​	当版本控制系统支持多个协议时，每个协议都会依次尝试下载。例如，Git 下载尝试 https://，然后是 git+ssh://。

By default, downloads are restricted to known secure protocols (e.g. https, ssh). To override this setting for Git downloads, the GIT_ALLOW_PROTOCOL environment variable can be set (For more details see: 'go help environment').

​	默认情况下，下载仅限于已知的安全协议（例如 https、ssh）。要覆盖此设置以进行 Git 下载，可以设置 GIT_ALLOW_PROTOCOL 环境变量（有关更多详细信息，请参见 'go help environment'）。

If the import path is not a known code hosting site and also lacks a version control qualifier, the go tool attempts to fetch the import over https/http and looks for a <meta> tag in the document's HTML <head>.

​	如果导入路径不是已知的代码托管站点，也缺乏版本控制限定符，则 go 工具尝试通过 https/http 获取导入，并在文档的 HTML `<head>` 中查找 `<meta>` 标签。

The meta tag has the form:

​	meta 标签的形式为：

```
<meta name="go-import" content="import-prefix vcs repo-root">
```

The import-prefix is the import path corresponding to the repository root. It must be a prefix or an exact match of the package being fetched with "go get". If it's not an exact match, another http request is made at the prefix to verify the <meta> tags match.

​	import-prefix 是与存储库根目录对应的导入路径。它必须是被 "go get" 获取的包的前缀或精确匹配。如果它不是精确匹配，则会在该前缀处进行另一个 http 请求以验证 <meta> 标签的匹配。

The meta tag should appear as early in the file as possible. In particular, it should appear before any raw JavaScript or CSS, to avoid confusing the go command's restricted parser.

​	meta 标签应尽可能早地出现在文件中。特别是，它应该出现在任何原始 JavaScript 或 CSS 之前，以避免混淆 go 命令的受限解析器。

The vcs is one of "bzr", "fossil", "git", "hg", "svn".

​	vcs 是 "bzr"、"fossil"、"git"、"hg"、"svn" 之一。

The repo-root is the root of the version control system containing a scheme and not containing a .vcs qualifier.

​	repo-root 是包含方案但不包含 .vcs 限定符的版本控制系统的根。

For example,

​	例如，

```
import "example.org/pkg/foo"
```

will result in the following requests:

将导致以下请求：

```
https://example.org/pkg/foo?go-get=1（首选）
http://example.org/pkg/foo?go-get=1（回退，仅在正确设置了 GOINSECURE 的情况下）
```

If that page contains the meta tag

​	如果该页面包含以下 meta 标签

```
<meta name="go-import" content="example.org git https://code.org/r/p/exproj">
```

the go tool will verify that https://example.org/?go-get=1 contains the same meta tag and then git clone https://code.org/r/p/exproj into GOPATH/src/example.org.

go 工具将验证 https://example.org/?go-get=1 是否包含相同的 meta 标签，然后将 git 克隆 https://code.org/r/p/exproj 到 GOPATH/src/example.org。

When using GOPATH, downloaded packages are written to the first directory listed in the GOPATH environment variable. (See 'go help gopath-get' and 'go help gopath'.)

​	在使用 GOPATH 时，下载的包会写入 GOPATH 环境变量中列出的第一个目录。（请参见 'go help gopath-get' 和 'go help gopath'。）

When using modules, downloaded packages are stored in the module cache. See https://golang.org/ref/mod#module-cache.

​	在使用模块时，下载的包存储在模块缓存中。请参阅 [Go模块参考中的模块缓存命令](../../GoModulesReference/ModuleCache)。

When using modules, an additional variant of the go-import meta tag is recognized and is preferred over those listing version control systems. That variant uses "mod" as the vcs in the content value, as in:

​	使用模块时，还有一种go-import meta标签的变体会被识别，并且优先于列出版本控制系统的标签。这个变体在content值中使用"mod"作为vcs，例如：

```
<meta name="go-import" content="example.org mod https://code.org/moduleproxy">
```

This tag means to fetch modules with paths beginning with example.org from the module proxy available at the URL https://code.org/moduleproxy. See https://golang.org/ref/mod#goproxy-protocol for details about the proxy protocol.

​	这个标签意味着要从可用于URL https://code.org/moduleproxy 的模块代理中获取以example.org开头的模块。有关代理协议的详细信息，请参见[Go模块参考中的模块代理](../../GoModulesReference/ModuleProxies#goproxy-protocol)。

#### 导入路径检查 

When the custom import path feature described above redirects to a known code hosting site, each of the resulting packages has two possible import paths, using the custom domain or the known hosting site.

​	当上述自定义导入路径功能重定向到已知的代码托管站点时，每个结果包都有两个可能的导入路径，一个是使用自定义域名，另一个是使用已知的托管站点。

A package statement is said to have an "import comment" if it is immediately followed (before the next newline) by a comment of one of these two forms:

​	如果一个包语句紧接着（在下一个换行符之前）有一个这两种形式之一的注释，那么它被称为"导入注释"：

``` go
package math // import "path"
package math /* import "path" */
```

The go command will refuse to install a package with an import comment unless it is being referred to by that import path. In this way, import comments let package authors make sure the custom import path is used and not a direct path to the underlying code hosting site.

​	go命令将拒绝安装带有导入注释的包，除非它被引用了该导入路径。通过这种方式，导入注释让包的作者确保使用自定义导入路径，而不是底层的代码托管站点的直接路径。

Import path checking is disabled for code found within vendor trees. This makes it possible to copy code into alternate locations in vendor trees without needing to update import comments.

​	在vendor树中发现的代码不会进行导入路径检查。这使得可以将代码复制到vendor树中的其他位置，而不需要更新导入注释。

Import path checking is also disabled when using modules. Import path comments are obsoleted by the go.mod file's module statement.

​	当使用模块时，导入路径检查也被禁用了。导入路径注释被go.mod文件中的模块声明所取代。

See https://golang.org/s/go14customimport for details.

​	有关详细信息，请参见https://golang.org/s/go14customimport。

#### 模块、模块版本等

Modules are how Go manages dependencies.

​	模块是Go管理依赖项的方式。

A module is a collection of packages that are released, versioned, and distributed together. Modules may be downloaded directly from version control repositories or from module proxy servers.

​	一个模块是一组一起发布、版本化和分发的包。模块可以直接从版本控制存储库或模块代理服务器下载。

For a series of tutorials on modules, see https://golang.org/doc/tutorial/create-module.

​	有关模块教程系列，请参见https://golang.org/doc/tutorial/create-module。

For a detailed reference on modules, see https://golang.org/ref/mod.

​	有关模块的详细参考，请参见[Go模块参考](../../GoModulesReference/Introduction)。

By default, the go command may download modules from [https://proxy.golang.org](https://proxy.golang.org/). It may authenticate modules using the checksum database at [https://sum.golang.org](https://sum.golang.org/). Both services are operated by the Go team at Google. The privacy policies for these services are available at https://proxy.golang.org/privacy and https://sum.golang.org/privacy, respectively.

​	默认情况下，go命令可以从https://proxy.golang.org下载模块。它可以使用https://sum.golang.org上的checksum数据库对模块进行身份验证。这两个服务都由Google的Go团队运营。这些服务的隐私政策可在https://proxy.golang.org/privacy和https://sum.golang.org/privacy上获得。

The go command's download behavior may be configured using GOPROXY, GOSUMDB, GOPRIVATE, and other environment variables. See 'go help environment' and https://golang.org/ref/mod#private-module-privacy for more information.

​	可以使用GOPROXY、GOSUMDB、GOPRIVATE和其他环境变量来配置go命令的下载行为。有关更多信息，请参见"go help environment"和[Go模块参考中的私有模块中的隐私](../../GoModulesReference/PrivateModules#privacy)。

#### 使用 go.sum 进行模块认证 

When the go command downloads a module zip file or go.mod file into the module cache, it computes a cryptographic hash and compares it with a known value to verify the file hasn't changed since it was first downloaded. Known hashes are stored in a file in the module root directory named go.sum. Hashes may also be downloaded from the checksum database depending on the values of GOSUMDB, GOPRIVATE, and GONOSUMDB.

​	当 go 命令将模块 zip 文件或 go.mod 文件下载到模块缓存中时，它会计算一个密码哈希值并将其与已知值进行比较，以验证该文件自从第一次下载以来没有发生更改。已知哈希值存储在模块根目录中名为 go.sum 的文件中。哈希值也可以根据 GOSUMDB、GOPRIVATE 和 GONOSUMDB 的值从校验和数据库中下载。

For details, see https://golang.org/ref/mod#authenticating.

​	For details, see [Go模块参考中的验证模块](../../GoModulesReference/AuthenticatingModules).

#### 包列表和模式 

Many commands apply to a set of packages:

​	许多命令都适用于一组包：

```
go action [packages]
```

Usually, [packages] is a list of import paths.

​	通常，[packages] 是一组导入路径。

An import path that is a rooted path or that begins with a . or .. element is interpreted as a file system path and denotes the package in that directory.

​	一个根路径或以 `.` 或 `..` 元素开头的导入路径被解释为文件系统路径，并表示该目录中的包。

Otherwise, the import path P denotes the package found in the directory DIR/src/P for some DIR listed in the GOPATH environment variable (For more details see: 'go help gopath').

​	否则，导入路径 P 表示在 GOPATH 环境变量列出的某个 DIR/src/P 目录中找到的包（有关详细信息，请参阅 'go help gopath'）。

If no import paths are given, the action applies to the package in the current directory.

​	如果没有给出导入路径，则该操作适用于当前目录中的包。

There are four reserved names for paths that should not be used for packages to be built with the go tool:

​	有四个保留名称用于不应与 go 工具一起构建的包：

- "main" denotes the top-level package in a stand-alone executable.
- "main"表示独立可执行文件中的顶级包。
- "all" expands to all packages found in all the GOPATH trees. For example, 'go list all' lists all the packages on the local system. When using modules, "all" expands to all packages in the main module and their dependencies, including dependencies needed by tests of any of those.
- "all"展开为在所有 GOPATH 树中找到的所有包。例如，'go list all' 列出了本地系统上的所有包。当使用模块时，"all"会展开为主模块中的所有包及其依赖项，包括任何这些依赖项的测试所需的依赖项。
-  "std" is like all but expands to just the packages in the standard Go library.
- "std"类似于"all"，但仅展开到标准 Go 库中的包。
-  "cmd" expands to the Go repository's commands and their internal libraries.
- "cmd"展开为 Go 存储库的命令及其内部库。

​	

Import paths beginning with "cmd/" only match source code in the Go repository.

​	以"cmd/"开头的导入路径仅匹配 Go 存储库中的源代码。

An import path is a pattern if it includes one or more "..." wildcards, each of which can match any string, including the empty string and strings containing slashes. Such a pattern expands to all package directories found in the GOPATH trees with names matching the patterns.

​	如果导入路径包含一个或多个"`...`"通配符，则导入路径是一个模式，每个通配符可以匹配任何字符串，包括空字符串和包含斜杠的字符串。这种模式将展开为 GOPATH 树中找到的所有包目录，其名称与模式匹配。

To make common patterns more convenient, there are two special cases. First, /... at the end of the pattern can match an empty string, so that net/... matches both net and packages in its subdirectories, like net/http. Second, any slash-separated pattern element containing a wildcard never participates in a match of the "vendor" element in the path of a vendored package, so that ./... does not match packages in subdirectories of ./vendor or ./mycode/vendor, but ./vendor/... and ./mycode/vendor/... do. Note, however, that a directory named vendor that itself contains code is not a vendored package: cmd/vendor would be a command named vendor, and the pattern cmd/... matches it. See golang.org/s/go15vendor for more about vendoring.

​	为了方便常见模式的匹配，有两个特殊情况。第一种情况是在模式末尾添加 `/...` 可以匹配空字符串，例如 `net/...` 同时匹配 net 和其子目录下的包，如 net/http。第二种情况是任何包含通配符的斜杠分隔模式元素都不会匹配 vendored 包路径中的 "vendor" 元素，这样 `./...` 就不会匹配 `./vendor` 或 `./mycode/vendor` 子目录下的包，但是 `./vendor/...` 和 `./mycode/vendor/...` 会。然而，一个包含代码的名为 vendor 的目录不是一个 vendored 包：cmd/vendor 将是一个名为 vendor 的命令，模式 `cmd/...` 会匹配它。有关 vendoring 的更多信息，请参见 golang.org/s/go15vendor。

An import path can also name a package to be downloaded from a remote repository. Run 'go help importpath' for details.

​	导入路径也可以命名从远程仓库下载的包。有关详情，请运行"go help importpath"。

Every package in a program must have a unique import path. By convention, this is arranged by starting each path with a unique prefix that belongs to you. For example, paths used internally at Google all begin with 'google', and paths denoting remote repositories begin with the path to the code, such as 'github.com/user/repo'.

​	程序中的每个包都必须具有唯一的导入路径。按照惯例，这是通过将每个路径以属于您自己的唯一前缀开头来安排的。例如，在 Google 内部使用的路径都以 'google' 开头，表示远程仓库的路径以代码的路径为开头，例如 'github.com/user/repo'。

Packages in a program need not have unique package names, but there are two reserved package names with special meaning. The name main indicates a command, not a library. Commands are built into binaries and cannot be imported. The name documentation indicates documentation for a non-Go program in the directory. Files in package documentation are ignored by the go command.

​	程序中的包名不需要唯一，但有两个保留的包名具有特殊含义。名称 main 表示命令，而不是库。命令被构建成二进制文件，不能被导入。名称 documentation 表示目录中非 Go 程序的文档。包文档中的文件将被 go 命令忽略。

As a special case, if the package list is a list of .go files from a single directory, the command is applied to a single synthesized package made up of exactly those files, ignoring any build constraints in those files and ignoring any other files in the directory.

​	作为特殊情况，如果包列表是来自单个目录的 .go 文件列表，则该命令将应用于一个合成包，该包由恰好这些文件组成，忽略这些文件中的任何构建约束，并忽略目录中的任何其他文件。

Directory and file names that begin with "." or "_" are ignored by the go tool, as are directories named "testdata".

​	以 "." 或 "_" 开头的目录和文件名将被 go 工具忽略，同样以 "testdata" 命名的目录也会被忽略。

#### 用于下载非公共代码的配置

The go command defaults to downloading modules from the public Go module mirror at proxy.golang.org. It also defaults to validating downloaded modules, regardless of source, against the public Go checksum database at sum.golang.org. These defaults work well for publicly available source code.

​	go 命令默认从公共 Go 模块镜像 proxy.golang.org 下载模块。无论源代码如何，它还默认针对公共 Go 校验和数据库 sum.golang.org 验证下载的模块。这些默认值适用于公开的源代码。

The GOPRIVATE environment variable controls which modules the go command considers to be private (not available publicly) and should therefore not use the proxy or checksum database. The variable is a comma-separated list of glob patterns (in the syntax of Go's path.Match) of module path prefixes. For example,

​	GOPRIVATE 环境变量控制了哪些模块被 go 命令视为私有模块（不公开可用），因此不应使用代理或校验和数据库。该变量是一个以逗号分隔的模块路径前缀的通配符模式列表（符合 Go 的 path.Match 语法）。例如，

```
GOPRIVATE=*.corp.example.com,rsc.io/private
```

causes the go command to treat as private any module with a path prefix matching either pattern, including git.corp.example.com/xyzzy, rsc.io/private, and rsc.io/private/quux.

会使 go 命令将任何具有与其中任意一种模式匹配的路径前缀的模块视为私有模块，包括 git.corp.example.com/xyzzy、rsc.io/private 和 rsc.io/private/quux 等。

For fine-grained control over module download and validation, the GONOPROXY and GONOSUMDB environment variables accept the same kind of glob list and override GOPRIVATE for the specific decision of whether to use the proxy and checksum database, respectively.

​	为了对模块下载和校验有更精细的控制，GONOPROXY 和 GONOSUMDB 环境变量接受相同类型的通配符模式列表，并分别覆盖 GOPRIVATE，用于确定是否使用代理和校验和数据库。

For example, if a company ran a module proxy serving private modules, users would configure go using:

​	例如，如果公司运行一个用于服务私有模块的模块代理，则用户可以使用以下方式配置 go：

```
GOPRIVATE=*.corp.example.com
GOPROXY=proxy.example.com
GONOPROXY=none
```

The GOPRIVATE variable is also used to define the "public" and "private" patterns for the GOVCS variable; see 'go help vcs'. For that usage, GOPRIVATE applies even in GOPATH mode. In that case, it matches import paths instead of module paths.

​	GOPRIVATE 变量还用于定义 GOVCS 变量的 "public" 和 "private" 模式；请参阅 'go help vcs'。对于这种用法，GOPRIVATE 即使在 GOPATH 模式下也适用。在这种情况下，它匹配导入路径而不是模块路径。

The 'go env -w' command (see 'go help env') can be used to set these variables for future go command invocations.

​	可以使用 'go env -w' 命令（参见 'go help env'）为将来的 go 命令调用设置这些变量。

For more details, see https://golang.org/ref/mod#private-modules.

​	有关更多详细信息，请参阅[Go模块参考中的私有模块](../../GoModulesReference/PrivateModules)。

#### 测试标志 

The 'go test' command takes both flags that apply to 'go test' itself and flags that apply to the resulting test binary.

​	`go test`命令既可以接受适用于 `go test`本身的标志，也可以接受适用于生成的测试二进制文件的标志。

Several of the flags control profiling and write an execution profile suitable for "go tool pprof"; run "go tool pprof -h" for more information. The --alloc_space, --alloc_objects, and --show_bytes options of pprof control how the information is presented.

​	其中一些标志用于控制性能分析，并生成适用于 "go tool pprof" 的执行分析文件；运行 "go tool pprof -h" 以获取更多信息。pprof 的 `--alloc_space`、`--alloc_objects` 和 `--show_bytes` 选项用于控制信息的显示方式。

The following flags are recognized by the 'go test' command and control the execution of any test:

​	以下是 `go test`命令可以识别并用于控制测试执行的标志：

```            
-bench regexp    
	Run only those benchmarks matching a regular expression.
    By default, no benchmarks are run.
    To run all benchmarks, use '-bench .' or '-bench=.'.
    The regular expression is split by unbracketed slash (/)
    characters into a sequence of regular expressions, and each
    part of a benchmark's identifier must match the corresponding
    element in the sequence, if any. Possible parents of matches
    are run with b.N=1 to identify sub-benchmarks. For example,
    given -bench=X/Y, top-level benchmarks matching X are run
    with b.N=1 to find any sub-benchmarks matching Y, which are
    then run in full.
    仅运行与正则表达式匹配的基准测试。
    默认情况下，不运行任何基准测试。
    要运行所有基准测试，使用 '-bench .' 或 '-bench=.'。
    正则表达式会根据未包含在括号中的斜杠 (/) 字符进行拆分，
    形成一系列正则表达式，基准测试标识的每个部分必须与序列中的
    相应元素匹配（如果有）。可能的父匹配会以 b.N=1 运行，
    以识别子基准测试。例如，给定 -bench=X/Y，
    将以 b.N=1 运行与 X 匹配的顶级基准测试，
    以查找与 Y 匹配的任何子基准测试，然后对其进行完整运行。

-benchtime t   
	Run enough iterations of each benchmark to take t, specified
    as a time.Duration (for example, -benchtime 1h30s).
    The default is 1 second (1s).
    The special syntax Nx means to run the benchmark N times
    (for example, -benchtime 100x).
    运行足够的基准测试迭代次数，以达到指定的时长 t，
    t以time.Duration形式指定（例如，-benchtime 1h30s）。
    默认情况下，时长为 1 秒（1s）。
    特殊语法 Nx 表示运行基准测试 N 次（例如，-benchtime 100x）。


-count n     
	Run each test, benchmark, and fuzz seed n times (default 1).
    If -cpu is set, run n times for each GOMAXPROCS value.
    Examples are always run once. -count does not apply to
    fuzz tests matched by -fuzz.
    运行每个测试、基准测试和模糊种子 n 次（默认为 1）。
    如果设置了 -cpu，则每个 GOMAXPROCS 值运行 n 次。
    示例总是运行一次。-count 不适用于由 -fuzz 匹配的模糊测试。


-cover       
	Enable coverage analysis.
    Note that because coverage works by annotating the source
    code before compilation, compilation and test failures with
    coverage enabled may report line numbers that don't correspond
    to the original sources.
    启用覆盖率(coverage)分析。
    请注意，由于覆盖率分析通过在编译之前对源代码进行注释来工作，
    因此启用覆盖率分析的编译和测试失败可能会报告不对应于原始源代码的行号。


-covermode set,count,atomic       
	Set the mode for coverage analysis for the package[s]
    being tested. The default is "set" unless -race is enabled,
    in which case it is "atomic".
    The values:
	set: bool: does this statement run?
	count: int: how many times does this statement run?
	atomic: int: count, but correct in multithreaded tests;
		significantly more expensive.
    Sets -cover.
    为正在测试的 package[s] 设置覆盖率(coverage)分析模式。
    默认为 "set"，除非启用了 -race，此时为 "atomic"。
    这些值分别是：
        set: bool：是否执行此语句？
        count: int：此语句运行多少次？
        atomic: int：计数，但在多线程测试中正确；代价较高。
    设置 -cover。


-coverpkg pattern1,pattern2,pattern3
	Apply coverage analysis in each test to packages matching the patterns.
    The default is for each test to analyze only the package being tested.
    See 'go help packages' for a description of package patterns.
    Sets -cover.
    在每个测试中将覆盖率(coverage)分析应用于与模式匹配的 package。
    默认情况下，每个测试仅分析正在测试的 package。
    有关 package 模式的描述，请参阅 'go help packages'。
    设置 -cover。
    
-cpu 1,2,4
	Specify a list of GOMAXPROCS values for which the tests, benchmarks or
    fuzz tests should be executed. The default is the current value
    of GOMAXPROCS. -cpu does not apply to fuzz tests matched by -fuzz.
    为测试、基准测试或模糊测试指定一组 GOMAXPROCS 值。
    默认值为当前的 GOMAXPROCS 值。
    -cpu 不适用于通过 -fuzz 匹配的模糊测试。


-failfast
	Do not start new tests after the first test failure.
    在第一个测试失败后，不再启动新的测试。
    
-fullpath
    Show full file names in the error messages.
    
-fuzz regexp
	Run the fuzz test matching the regular expression. When specified,
    the command line argument must match exactly one package within the
    main module, and regexp must match exactly one fuzz test within
    that package. Fuzzing will occur after tests, benchmarks, seed corpora
    of other fuzz tests, and examples have completed. See the Fuzzing
    section of the testing package documentation for details.
    运行与正则表达式匹配的模糊测试。
    在指定此标志时，命令行参数必须完全匹配主模块中的一个包，
    并且 regexp 必须完全匹配该包内的一个模糊测试。
    模糊测试将在测试、基准测试、其他模糊测试的种子语料库
    以及示例完成后进行。
    有关详细信息，请参阅testing包文档中的模糊测试部分。


-fuzztime t
	Run enough iterations of the fuzz target during fuzzing to take t,
    specified as a time.Duration (for example, -fuzztime 1h30s).
	The default is to run forever.
    The special syntax Nx means to run the fuzz target N times
    (for example, -fuzztime 1000x).
    在模糊测试期间运行足够的模糊目标迭代，以达到指定的时长 t，
    t以time.Duration形式指定（例如，-fuzztime 1h30s）。
    默认值为永远运行。
    特殊语法 Nx 表示运行模糊目标 N 次（例如，-fuzztime 1000x）。
    
-fuzzminimizetime t
	Run enough iterations of the fuzz target during each minimization
    attempt to take t, as specified as a time.Duration (for example,
    -fuzzminimizetime 30s).
	The default is 60s.
    The special syntax Nx means to run the fuzz target N times
    (for example, -fuzzminimizetime 100x).
    在每次最小化尝试期间运行足够的模糊目标迭代，以达到指定的时长 t，
    t以time.Duration形式指定（例如，-fuzzminimizetime 30s）。
    默认为 60 秒。
    特殊语法 Nx 表示运行模糊目标 N 次（例如，-fuzzminimizetime 100x）。


-json
	Log verbose output and test results in JSON. This presents the
    same information as the -v flag in a machine-readable format.
    以 JSON 格式记录详细输出和测试结果。
    这以机器可读的格式呈现与 -v 标志相同的信息。


-list regexp
	List tests, benchmarks, fuzz tests, or examples matching the regular
    expression. No tests, benchmarks, fuzz tests, or examples will be run.
    This will only list top-level tests. No subtest or subbenchmarks will be
    shown.
    列出与正则表达式匹配的测试、基准测试、模糊测试或示例。
    不运行任何测试、基准测试、模糊测试或示例。
    仅列出顶级测试。不显示子测试或子基准测试。
    
-parallel n    
	Allow parallel execution of test functions that call t.Parallel, and
    fuzz targets that call t.Parallel when running the seed corpus.
    The value of this flag is the maximum number of tests to run
    simultaneously.
    While fuzzing, the value of this flag is the maximum number of
    subprocesses that may call the fuzz function simultaneously, regardless of
    whether T.Parallel is called.
    By default, -parallel is set to the value of GOMAXPROCS.
    Setting -parallel to values higher than GOMAXPROCS may cause degraded
    performance due to CPU contention, especially when fuzzing.
    Note that -parallel only applies within a single test binary.
    The 'go test' command may run tests for different packages
    in parallel as well, according to the setting of the -p flag
    (see 'go help build').
    允许并行执行调用了 t.Parallel 的测试函数，
    以及运行种子语料库时调用了 t.Parallel 的模糊目标。
    该标志的值是要同时运行的最大测试数。
    在模糊测试时，该标志的值是可能同时调用模糊函数的最大子进程数，
    无论是否调用了 T.Parallel。
    默认情况下，-parallel 设置为 GOMAXPROCS 的值。
    将 -parallel 设置为高于 GOMAXPROCS 的值可能会导致性能降低，
    因为 CPU 冲突，特别是在模糊测试时。
    请注意，-parallel 仅在单个测试二进制文件中适用。
    根据 -p 标志的设置（请参阅 'go help build'），
    'go test' 命令也可以并行运行不同包的测试。


-run regexp
	Run only those tests, examples, and fuzz tests matching the regular
    expression. For tests, the regular expression is split by unbracketed
    slash (/) characters into a sequence of regular expressions, and each
    part of a test's identifier must match the corresponding element in
    the sequence, if any. Note that possible parents of matches are
    run too, so that -run=X/Y matches and runs and reports the result
    of all tests matching X, even those without sub-tests matching Y,
    because it must run them to look for those sub-tests.
    See also -skip.
    仅运行与正则表达式匹配的测试、示例和模糊测试。
    对于测试，正则表达式会通过未包含在括号中的斜杠（/）
    字符拆分为一系列正则表达式，而测试的每个部分必须与序列中的
    相应元素匹配（如果有的话）。请注意，可能的匹配父级也会运行，
    因此 -run=X/Y 会匹配并运行 X 匹配的所有测试的结果，
    即使没有子测试与 Y 匹配，因为必须运行它们以查找这些子测试。
    另请参阅 -skip。
    
-short
	Tell long-running tests to shorten their run time.
    It is off by default but set during all.bash so that installing
    the Go tree can run a sanity check but not spend time running
    exhaustive tests.
    告诉长时间运行的测试缩短其运行时间。
    默认情况下未启用，但在 all.bash 中设置为启用，
    以便在运行全面测试时可以运行一次健全性检查，
    而不会花费时间运行详尽测试。


-shuffle off,on,N
	Randomize the execution order of tests and benchmarks.
    It is off by default. If -shuffle is set to on, then it will seed
    the randomizer using the system clock. If -shuffle is set to an
    integer N, then N will be used as the seed value. In both cases,
    the seed will be reported for reproducibility.
    随机化测试和基准测试的执行顺序。
    默认情况下为off。如果将 -shuffle 设置为 on，
    则它将使用系统时钟对随机生成器（randomizer）进行种子化。
    如果将 -shuffle 设置为整数 N，
    则 N 将用作种子值。在这两种情况下，种子将被报告以便进行重现。


-skip regexp
	Run only those tests, examples, fuzz tests, and benchmarks that
    do not match the regular expression. Like for -run and -bench,
    for tests and benchmarks, the regular expression is split by unbracketed
    slash (/) characters into a sequence of regular expressions, and each
    part of a test's identifier must match the corresponding element in
    the sequence, if any.
    仅运行与正则表达式不匹配的测试、示例、模糊测试和基准测试。
    与 -run 和 -bench 一样，对于测试和基准测试，
    正则表达式会根据未括号化的斜杠 (/) 字符分成一系列正则表达式，
    每个测试的标识符部分必须与序列中的相应元素匹配（如果有的话）。


-timeout d
	If a test binary runs longer than duration d, panic.
    If d is 0, the timeout is disabled.
    The default is 10 minutes (10m).
    如果测试二进制运行时长超过 d，将触发 panic。
    如果 d 为 0，则禁用超时。
    默认值为 10 分钟（10m）。
    
-v
	Verbose output: log all tests as they are run. Also print all
    text from Log and Logf calls even if the test succeeds.
    详细输出：记录所有运行的测试。
    即使测试成功，也会打印来自 Log 和 Logf 调用的所有文本。


-vet list     
	Configure the invocation of "go vet" during "go test"
    to use the comma-separated list of vet checks.
    If list is empty, "go test" runs "go vet" with a curated list of
    checks believed to be always worth addressing.
    If list is "off", "go test" does not run "go vet" at all.
    配置在 "go test" 期间调用 "go vet" 
    使用逗号分隔的 vet 检查列表。
    如果列表为空，
    则 "go test" 使用相信始终值得解决的检查的精选列表运行 "go vet"。
    如果列表为 "off"，则 "go test" 根本不运行 "go vet"。
```

The following flags are also recognized by 'go test' and can be used to profile the tests during execution:

​	`go test` 命令还识别以下标志，用于在执行过程中对测试进行分析：

```
-benchmem
	Print memory allocation statistics for benchmarks.
    打印基准测试的内存分配统计信息。

-blockprofile block.out
    Write a goroutine blocking profile to the specified file
    when all tests are complete.
    Writes test binary as -c would.
    在所有测试完成后，将 goroutine 阻塞分析写入指定的文件。
    将测试二进制文件写入 -c 的目录。

-blockprofilerate n       
    通过调用 runtime.SetBlockProfileRate 
    设置 Goroutine 阻塞概要中提供的详细信息。
    请参阅 'go doc runtime.SetBlockProfileRate'。
    该分析器（profiler）的目标是平均每个程序被
    阻塞的 n 纳秒时间内采样一次阻塞事件。
    默认情况下，如果设置了 -test.blockprofile 但没有设置此标志，
    则会记录所有阻塞事件，相当于 -test.blockprofilerate=1。
        
-coverprofile cover.out
	Write a coverage profile to the file after all tests have passed.
    Sets -cover.
    在所有测试通过后，将覆盖率（coverage）分析写入文件。
    设置 -cover。
            
-cpuprofile cpu.out
	Write a CPU profile to the specified file before exiting.
    Writes test binary as -c would.
    在退出之前，将 CPU 分析（profile）写入指定的文件。
    将测试二进制文件写入 -c 的目录。

-memprofile mem.out
	Write an allocation profile to the file after all tests have passed.
    Writes test binary as -c would.
    在所有测试通过后，将分配（allocation）分析写入文件。
    将测试二进制文件写入 -c 的目录。

-memprofilerate n
	Enable more precise (and expensive) memory allocation profiles by
    setting runtime.MemProfileRate. See 'go doc runtime.MemProfileRate'.
    To profile all memory allocations, use -test.memprofilerate=1.
    通过设置 runtime.MemProfileRate 来启用更精确（但更昂贵）
    的内存分配分析。有关详细信息，
    请参阅 'go doc runtime.MemProfileRate'。
    若要分析所有内存分配，请使用 -test.memprofilerate=1。

-mutexprofile mutex.out
	Write a mutex contention profile to the specified file
    when all tests are complete.
    Writes test binary as -c would.
    在所有测试完成后，将互斥锁争用（mutex contention）分析写入指定的文件。
    将测试二进制文件写入 -c 的目录。

-mutexprofilefraction n
	Sample 1 in n stack traces of goroutines holding a
    contended mutex.
    对持有有争议的互斥锁的 goroutine 进行1:n的栈跟踪采样。

-outputdir directory
	Place output files from profiling in the specified directory,
    by default the directory in which "go test" is running.
    将性能分析的输出文件放置在指定的目录中，
    默认情况下是"go test"正在运行的目录。
        
-trace trace.out
 	Write an execution trace to the specified file before exiting.
    在退出之前，将执行跟踪写入指定的文件。
```

Each of these flags is also recognized with an optional 'test.' prefix, as in -test.v. When invoking the generated test binary (the result of 'go test -c') directly, however, the prefix is mandatory.


​	这些标志也可以使用可选的 'test.' 前缀来识别，例如 `-test.v`。然而，当直接调用生成的测试二进制文件（由 'go test -c' 生成）时，前缀是必需的。

The `go test`command rewrites or removes recognized flags, as appropriate, both before and after the optional package list, before invoking the test binary.

​	`go test` 命令会在调用测试二进制文件之前和之后，根据需要在可选的包列表之前和之后重写或删除已识别的标志。

For instance, the command

​	例如，以下命令：

        go test -v -myflag testdata -cpuprofile=prof.out -x

will compile the test binary and then run it as

将会编译测试二进制文件，然后作为以下方式运行：

        pkg.test -test.v -myflag testdata -test.cpuprofile=prof.out

(The -x flag is removed because it applies only to the go command's execution, not to the test itself.)

（`-x` 标志已被移除，因为它仅适用于 go 命令的执行，不适用于测试本身。）

The test flags that generate profiles (other than for coverage) also leave the test binary in pkg.test for use when analyzing the profiles.

​	生成性能分析的测试标志（除了覆盖率分析）还会保留测试二进制文件 pkg.test，以供分析性能分析时使用。

When 'go test' runs a test binary, it does so from within the corresponding package's source code directory. Depending on the test, it may be necessary to do the same when invoking a generated test binary directly. Because that directory may be located within the module cache, which may be read-only and is verified by checksums, the test must not write to it or any other directory within the module unless explicitly requested by the user (such as with the -fuzz flag, which writes failures to testdata/fuzz).

​	当 `go test` 运行测试二进制文件时，它是从相应包的源代码目录内运行的。根据测试的不同，当直接调用生成的测试二进制文件时，可能需要做相同的操作。因为该目录可能位于模块缓存中，模块缓存可能是只读的，并且通过校验和进行验证，所以测试不能将其写入或写入模块内的任何其他目录，除非用户明确要求（例如，使用 `-fuzz` 标志，将失败写入 `testdata/fuzz`目录）。

The command-line package list, if present, must appear before any flag not known to the go test command. Continuing the example above, the package list would have to appear before -myflag, but could appear on either side of -v.

​	命令行 package 列表（如果存在）必须出现在 go test 命令不识别的任何标志之前。延续上面的示例，package 列表必须出现在 `-myflag` 之前，但可以出现在 `-v` 的任一侧。

When 'go test' runs in package list mode, 'go test' caches successful package test results to avoid unnecessary repeated running of tests. To disable test caching, use any test flag or argument other than the cacheable flags. The idiomatic way to disable test caching explicitly is to use -count=1.

​	当 'go test' 在 package 列表模式下运行时，它会缓存成功的包测试结果，以避免不必要地重复运行测试。要禁用测试缓存，请使用除了可缓存标志之外的任何测试标志或参数。显式禁用测试缓存的惯用方法是使用 `-count=1`。

To keep an argument for a test binary from being interpreted as a known flag or a package name, use -args (see 'go help test') which passes the remainder of the command line through to the test binary uninterpreted and unaltered.

​	为了防止测试二进制文件的参数被解释为已知标志或包名，请使用 `-args`（参见 'go help test'），该标志将命令行的其余部分不加解释地传递给测试二进制文件。

For instance, the command

​	例如，以下命令：

        go test -v -args -x -v

will compile the test binary and then run it as

将编译测试二进制文件，然后作为以下方式运行：

        pkg.test -test.v -x -v

Similarly,

类似地，

        go test -args math

will compile the test binary and then run it as

将编译测试二进制文件，然后作为以下方式运行：

        pkg.test math

In the first example, the -x and the second -v are passed through to the test binary unchanged and with no effect on the go command itself. In the second example, the argument math is passed through to the test binary, instead of being interpreted as the package list.

​	在第一个示例中，`-x` 和第二个 `-v` 被不加解释地传递给测试二进制文件，不会对 go 命令本身产生影响。在第二个示例中，参数 `math` 被不加解释地传递给测试二进制文件，而不是被解释为 package 列表。

#### 测试函数 

The 'go test' command expects to find test, benchmark, and example functions in the "*_test.go" files corresponding to the package under test.

​	'go test' 命令期望在对应被测试包的 "*_test.go" 文件中找到测试、基准测试和示例函数。

A test function is one named TestXxx (where Xxx does not start with a lower case letter) and should have the signature,

​	测试函数以 TestXxx 命名（其中 Xxx 不以小写字母开头），应该有如下签名：

``` go
func TestXxx(t *testing.T) { ... }
```

A benchmark function is one named BenchmarkXxx and should have the signature,

​	基准测试函数以 BenchmarkXxx 命名，应该有如下签名：

``` go
func BenchmarkXxx(b *testing.B) { ... }
```

A fuzz test is one named FuzzXxx and should have the signature,

​	模糊测试以 FuzzXxx 命名，应该有如下签名：

``` go
func FuzzXxx(f *testing.F) { ... }
```

An example function is similar to a test function but, instead of using *testing.T to report success or failure, prints output to os.Stdout. If the last comment in the function starts with "Output:" then the output is compared exactly against the comment (see examples below). If the last comment begins with "Unordered output:" then the output is compared to the comment, however the order of the lines is ignored. An example with no such comment is compiled but not executed. An example with no text after "Output:" is compiled, executed, and expected to produce no output.

​	示例函数与测试函数类似，但是输出到 os.Stdout 而不是使用 *testing.T 来报告成功或失败。如果函数的最后一个注释以 "Output:" 开头，则输出与注释进行完全匹配（参见下面的示例）。如果最后一个注释以 "Unordered output:" 开头，则输出与注释进行匹配，但是忽略行的顺序。没有此类注释的示例会被编译但不会被执行。在 "Output:" 后没有文本的示例会被编译、执行，并期望不产生输出。

Godoc displays the body of ExampleXxx to demonstrate the use of the function, constant, or variable Xxx. An example of a method M with receiver type T or *T is named ExampleT_M. There may be multiple examples for a given function, constant, or variable, distinguished by a trailing _xxx, where xxx is a suffix not beginning with an upper case letter.

​	Godoc 显示 ExampleXxx 的正文，以演示函数、常量或变量 Xxx 的使用。具有接收者类型 T 或 *T 的方法 M 的示例命名为 ExampleT_M。对于给定函数、常量或变量，可以有多个示例，它们通过一个尾部 _xxx 区分开，其中 xxx 是一个不以大写字母开头的后缀。

Here is an example of an example:

​	下面是一个示例：

``` go
func ExamplePrintln() {
	Println("The output of\nthis example.")
	// Output: The output of
	// this example.
}
```

Here is another example where the ordering of the output is ignored:

​	下面是另一个示例，其中输出的顺序被忽略：

``` go
func ExamplePerm() {
	for _, value := range Perm(4) {
		fmt.Println(value)
	}

	// Unordered output: 4
	// 2
	// 1
	// 3
	// 0
}
```

The entire test file is presented as the example when it contains a single example function, at least one other function, type, variable, or constant declaration, and no tests, benchmarks, or fuzz tests.

​	当测试文件包含一个单独的示例函数、至少一个其他函数、类型、变量或常量声明以及没有测试、基准测试或模糊测试时，整个测试文件将被作为示例呈现。

See the documentation of the testing package for more information.

​	有关更多信息，请参阅 testing 包的文档。

#### 用 GOVCS 控制版本控制

The 'go get' command can run version control commands like git to download imported code. This functionality is critical to the decentralized Go package ecosystem, in which code can be imported from any server, but it is also a potential security problem, if a malicious server finds a way to cause the invoked version control command to run unintended code.

​	'go get' 命令可以运行版本控制命令（如 git）来下载导入的代码。这个功能对于 Go 分散式包生态系统至关重要，其中的代码可以从任何服务器导入，但如果恶意服务器找到了一种方法来让调用的版本控制命令运行未预期的代码，这也是一个潜在的安全问题。

To balance the functionality and security concerns, the 'go get' command by default will only use git and hg to download code from public servers. But it will use any known version control system (bzr, fossil, git, hg, svn) to download code from private servers, defined as those hosting packages matching the GOPRIVATE variable (see 'go help private'). The rationale behind allowing only Git and Mercurial is that these two systems have had the most attention to issues of being run as clients of untrusted servers. In contrast, Bazaar, Fossil, and Subversion have primarily been used in trusted, authenticated environments and are not as well scrutinized as attack surfaces.

​	为了平衡功能和安全问题，'go get' 命令默认只使用 git 和 hg 从公共服务器下载代码。但它将使用任何已知的版本控制系统（bzr、fossil、git、hg、svn）从私有服务器下载代码，这些服务器定义为托管匹配 GOPRIVATE 变量（请参见 'go help private'）的包。允许 Git 和 Mercurial 的原理在于这两个系统在作为不受信任服务器的客户端运行时已经得到了最多的问题关注。相比之下，Bazaar、Fossil 和 Subversion 主要用于受信任的身份验证环境中，并且没有像攻击面一样受到充分的审查。

The version control command restrictions only apply when using direct version control access to download code. When downloading modules from a proxy, 'go get' uses the proxy protocol instead, which is always permitted. By default, the 'go get' command uses the Go module mirror (proxy.golang.org) for public packages and only falls back to version control for private packages or when the mirror refuses to serve a public package (typically for legal reasons). Therefore, clients can still access public code served from Bazaar, Fossil, or Subversion repositories by default, because those downloads use the Go module mirror, which takes on the security risk of running the version control commands using a custom sandbox.

​	版本控制命令限制仅适用于使用直接版本控制访问下载代码时。当从代理服务器下载模块时，'go get' 使用代理协议，该协议始终被允许。默认情况下，'go get' 命令使用 Go 模块镜像（proxy.golang.org）获取公共包，仅在模块镜像拒绝为公共包提供服务（通常是因为法律原因）时，才回退到版本控制以获取私有包。因此，客户端仍然可以访问由 Bazaar、Fossil 或 Subversion 存储库提供的公共代码，默认情况下，因为这些下载使用 Go 模块镜像，该镜像承担运行版本控制命令的安全风险，使用自定义沙箱。

The GOVCS variable can be used to change the allowed version control systems for specific packages (identified by a module or import path). The GOVCS variable applies when building package in both module-aware mode and GOPATH mode. When using modules, the patterns match against the module path. When using GOPATH, the patterns match against the import path corresponding to the root of the version control repository.

​	GOVCS 变量可用于更改特定包（由模块或导入路径标识）的允许版本控制系统。GOVCS 变量在模块感知模式和 GOPATH 模式下构建包时都适用。当使用模块时，模式与模块路径匹配。使用 GOPATH 时，模式与对应于版本控制存储库根的导入路径匹配。

The general form of the GOVCS setting is a comma-separated list of pattern:vcslist rules. The pattern is a glob pattern that must match one or more leading elements of the module or import path. The vcslist is a pipe-separated list of allowed version control commands, or "all" to allow use of any known command, or "off" to disallow all commands. Note that if a module matches a pattern with vcslist "off", it may still be downloaded if the origin server uses the "mod" scheme, which instructs the go command to download the module using the GOPROXY protocol. The earliest matching pattern in the list applies, even if later patterns might also match.

​	GOVCS 的一般形式为 pattern:vcslist 规则，由逗号分隔的一组规则构成。pattern 是必须匹配模块或导入路径的一个或多个前导元素的 glob 模式。vcslist 是允许的版本控制命令的管道分隔列表，或者是 "all" 表示允许使用任何已知命令，或者是 "off" 表示禁止使用所有命令。请注意，如果一个模块匹配一个 vcslist 为 "off" 的模式，则仍可以下载，如果源服务器使用 "mod" 方案，则指示 go 命令使用 GOPROXY 协议下载模块。列表中最早匹配的模式优先应用，即使后面的模式也匹配。

For example, consider:

​	例如，请考虑：

```
GOVCS=github.com:git,evil.com:off,*:git|hg
```

​	With this setting, code with a module or import path beginning with github.com/ can only use git; paths on evil.com cannot use any version control command, and all other paths (* matches everything) can use only git or hg.

使用此设置，具有以 `github.com/` 开头的模块或导入路径的代码只能使用 git。evil.com 上的路径不能使用任何版本控制命令，而其他所有路径（`*` 匹配所有）只能使用 git 或 hg。

The special patterns "public" and "private" match public and private module or import paths. A path is private if it matches the GOPRIVATE variable; otherwise it is public.

​	特殊模式 "public" 和 "private" 匹配公共和私有模块或导入路径。如果路径与 GOPRIVATE 变量匹配，则为私有路径；否则为公共路径。

If no rules in the GOVCS variable match a particular module or import path, the 'go get' command applies its default rule, which can now be summarized in GOVCS notation as 'public:git|hg,private:all'.

​	如果 GOVCS 变量中没有规则与特定模块或导入路径匹配，则 'go get' 命令应用其默认规则，该规则现在可以用 GOVCS 表示为 'public:git|hg,private:all'。

To allow unfettered use of any version control system for any package, use:

​	要允许任何包使用任何版本控制系统，可以使用：

```
GOVCS=*:all
```

To disable all use of version control, use:

​	要禁用所有版本控制的使用，请使用：

```
GOVCS=*:off
```

The 'go env -w' command (see 'go help env') can be used to set the GOVCS variable for future go command invocations.

​	"go env -w" 命令（请参阅 "go help env"）可用于设置 GOVCS 变量以供将来的 go 命令调用。
