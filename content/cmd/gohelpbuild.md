+++
title = "go help build"
date = 2024-02-24T11:12:32+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

usage: go build [-o output] [build flags] [packages]

​	用法：go build [-o output] [build flags] [packages]

Build compiles the packages named by the import paths, along with their dependencies, but it does not install the results.

​	build 编译由导入路径命名的包及其依赖项，但不会安装结果。

If the arguments to build are a list of .go files from a single directory, build treats them as a list of source files specifying a single package.

​	如果 build 的参数是来自单个目录的 .go 文件列表，则 build 将它们视为指定单个包的源文件列表。

When compiling packages, build ignores files that end in ‘_test.go’.

​	在编译包时，build 会忽略以“_test.go”结尾的文件。

When compiling a single main package, build writes the resulting executable to an output file named after the first source file (‘go build ed.go rx.go’ writes ’ed’ or ’ed.exe’) or the source code directory (‘go build unix/sam’ writes ‘sam’ or ‘sam.exe’). The ‘.exe’ suffix is added when writing a Windows executable.

​	在编译单个 main 包时，build 会将生成的执行文件写入以第一个源文件命名的输出文件（“go build ed.go rx.go”会写入“ed”或“ed.exe”）或源代码目录（“go build unix/sam”会写入“sam”或“sam.exe”）。在编写 Windows 可执行文件时，会添加“.exe”后缀。

When compiling multiple packages or a single non-main package, build compiles the packages but discards the resulting object, serving only as a check that the packages can be built.

​	在编译多个包或单个非 main 包时，build 会编译这些包，但会丢弃生成的对象，仅用作检查这些包是否可以构建。

The -o flag forces build to write the resulting executable or object to the named output file or directory, instead of the default behavior described in the last two paragraphs. If the named output is an existing directory or ends with a slash or backslash, then any resulting executables will be written to that directory.

​	-o 标志强制构建将生成的执行文件或对象写入命名的输出文件或目录，而不是最后两段中描述的默认行为。如果命名的输出是现有目录或以斜杠或反斜杠结尾，则任何生成的执行文件都将写入该目录。

The build flags are shared by the build, clean, get, install, list, run, and test commands:

​	构建标志由构建、清理、获取、安装、列出、运行和测试命令共享：

        -C dir
                Change to dir before running the command.
                Any files named on the command line are interpreted after
                changing directories.
                If used, this flag must be the first one in the command line.
                在运行命令之前切换到指定目录。
    			命令行中指定的任何文件将在更改目录后进行解释。
    			如果使用了此标志，则必须将其放在命令行中的第一个位置。
        -a
                force rebuilding of packages that are already up-to-date.
                强制重新构建已经是最新状态的包。
                
        -n
                print the commands but do not run them.
                打印命令但不运行它们。
        -p n
                the number of programs, such as build commands or
                test binaries, that can be run in parallel.
                The default is GOMAXPROCS, normally the number of CPUs available.
                可以并行运行的程序数量，例如构建命令或测试二进制文件。
    			默认值为 GOMAXPROCS，通常为可用的 CPU 数量。
        -race
                enable data race detection.
                Supported only on linux/amd64, freebsd/amd64, darwin/amd64, darwin/arm64, windows/amd64,
                linux/ppc64le and linux/arm64 (only for 48-bit VMA).
                启用数据竞争检测。
    			仅在 linux/amd64、freebsd/amd64、darwin/amd64、darwin/arm64、windows/amd64、
    			linux/ppc64le 和 linux/arm64（仅适用于 48 位 VMA）上支持。
        -msan
                enable interoperation with memory sanitizer.
                Supported only on linux/amd64, linux/arm64, freebsd/amd64
                and only with Clang/LLVM as the host C compiler.
                PIE build mode will be used on all platforms except linux/amd64.
                启用与内存污染检测器的互操作。
    			仅在 linux/amd64、linux/arm64、freebsd/amd64 上支持，
    			并且仅在使用 Clang/LLVM 作为主机 C 编译器时可用。
    			在所有平台上，除了 linux/amd64 外，都将使用 PIE 构建模式。
        -asan
                enable interoperation with address sanitizer.
                Supported only on linux/arm64, linux/amd64.
                Supported only on linux/amd64 or linux/arm64 and only with GCC 7 and higher
                or Clang/LLVM 9 and higher.
                启用与地址污染检测器的互操作。
    			仅在 linux/arm64、linux/amd64 上支持。
    			仅在 linux/amd64 或 linux/arm64 上，并且仅在 GCC 7 或更高版本、
    			或者 Clang/LLVM 9 或更高版本下支持时可用。
        -cover
                enable code coverage instrumentation.
                启用代码覆盖率检测。
                
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
                设置覆盖率分析的模式。
    			默认值为 "set"，除非启用了 -race，则为 "atomic"。
    			可选值为：
    				set: 布尔值：此语句是否执行？
    				count: 整数值：此语句执行了多少次？
    				atomic: 整数值：count，但在多线程测试中正确；成本昂贵。
    			设置 -cover。
        -coverpkg pattern1,pattern2,pattern3
                For a build that targets package 'main' (e.g. building a Go
                executable), apply coverage analysis to each package matching
                the patterns. The default is to apply coverage analysis to
                packages in the main Go module. See 'go help packages' for a
                description of package patterns.  Sets -cover.
                对于目标为 'main' 包（例如构建 Go 可执行文件）的构建，
    			将覆盖率分析应用于与模式匹配的每个包。
    			默认情况下，将覆盖率分析应用于主 Go 模块中的包。
    			有关包模式的描述，请参阅 'go help packages'。
    			设置 -cover。
        -v
                print the names of packages as they are compiled.
                在编译包时打印包的名称。
        -work
                print the name of the temporary work directory and
                do not delete it when exiting.
                打印临时工作目录的名称，并在退出时不删除它。
        -x
                print the commands.
                打印命令。
        -asmflags '[pattern=]arg list'
                arguments to pass on each go tool asm invocation.
                传递给每个 go tool asm 调用的参数。
        -buildmode mode
                build mode to use. See 'go help buildmode' for more.
                要使用的构建模式。有关更多信息，请参阅 'go help buildmode'。
        -buildvcs
                Whether to stamp binaries with version control information
                ("true", "false", or "auto"). By default ("auto"), version control
                information is stamped into a binary if the main package, the main module
                containing it, and the current directory are all in the same repository.
                Use -buildvcs=false to always omit version control information, or
                -buildvcs=true to error out if version control information is available but
                cannot be included due to a missing tool or ambiguous directory structure.
                是否在二进制文件中加入版本控制信息（"true"、"false" 或 "auto"）。
    			默认情况下（"auto"），如果主包、包含它的主模块和当前目录都在同一个仓库中，
    			则将版本控制信息加入二进制文件。
    			使用 -buildvcs=false 可以始终省略版本控制信息，或者使用 -buildvcs=true
    			如果版本控制信息可用但由于缺少工具或模糊的目录结构而无法包含，则报错。
        -compiler name
                name of compiler to use, as in runtime.Compiler (gccgo or gc
                要使用的编译器的名称，如 runtime.Compiler (gccgo 或 gc)。
        -gccgoflags '[pattern=]arg list'
                arguments to pass on each gccgo compiler/linker invocation.
                传递给每个 gccgo 编译器/链接器调用的参数。
        -gcflags '[pattern=]arg list'
                arguments to pass on each go tool compile invocation.
                传递给每个 go tool compile 调用的参数。            
        -installsuffix suffix
                a suffix to use in the name of the package installation directory,
                in order to keep output separate from default builds.
                If using the -race flag, the install suffix is automatically set to race
                or, if set explicitly, has _race appended to it. Likewise for the -msan
                and -asan flags. Using a -buildmode option that requires non-default compile
                flags has a similar effect.
                在包安装目录的名称中添加后缀，
    			以便将输出与默认构建分开。
    			如果使用 -race 标志，则自动设置安装后缀为 race，
    			或者如果显式设置了，则附加 _race。对于 -msan 和 -asan 标志也是如此。
    			使用需要非默认编译标志的 -buildmode 选项会产生类似的效果。
        -ldflags '[pattern=]arg list'
                arguments to pass on each go tool link invocation.
                传递给每个 go tool link 调用的参数。
        -linkshared
                build code that will be linked against shared libraries previously
                created with -buildmode=shared.
                构建将链接到先前使用 -buildmode=shared 创建的共享库的代码。
        -mod mode
                module download mode to use: readonly, vendor, or mod.
                By default, if a vendor directory is present and the go version in go.mod
                is 1.14 or higher, the go command acts as if -mod=vendor were set.
                Otherwise, the go command acts as if -mod=readonly were set.
                See https://golang.org/ref/mod#build-commands for details.
                要使用的模块下载模式：readonly、vendor 或 mod。
    			默认情况下，如果存在 vendor 目录并且 go.mod 中的 go 版本为 1.14 或更高，
    			则 go 命令将表现为设置了 -mod=vendor。
    			否则，go 命令将表现为设置了 -mod=readonly。
    			有关详情，请参阅 https://golang.org/ref/mod#build-commands。
        -modcacherw
                leave newly-created directories in the module cache read-write
                instead of making them read-only.
                将新创建的目录保留在模块缓存中可读写状态，而不是只读状态。
        -modfile file
                in module aware mode, read (and possibly write) an alternate go.mod
                file instead of the one in the module root directory. A file named
                "go.mod" must still be present in order to determine the module root
                directory, but it is not accessed. When -modfile is specified, an
                alternate go.sum file is also used: its path is derived from the
                -modfile flag by trimming the ".mod" extension and appending ".sum".
                在模块感知模式下，读取（可能写入）替代的 go.mod 文件，
    			而不是模块根目录中的文件。仍然必须存在名为 "go.mod" 的文件，
    			以确定模块根目录，但不会访问该文件。
    			当指定了 -modfile 时，还将使用替代的 go.sum 文件：其路径是从
    			-modfile 标志派生的，通过去除 ".mod" 扩展名并附加 ".sum"。
        -overlay file
                read a JSON config file that provides an overlay for build operations.
                The file is a JSON struct with a single field, named 'Replace', that
                maps each disk file path (a string) to its backing file path, so that
                a build will run as if the disk file path exists with the contents
                given by the backing file paths, or as if the disk file path does not
                exist if its backing file path is empty. Support for the -overlay flag
                has some limitations: importantly, cgo files included from outside the
                include path must be in the same directory as the Go package they are
                included from, and overlays will not appear when binaries and tests are
                run through go run and go test respectively.
                读取提供构建操作覆盖的 JSON 配置文件。
    			该文件是一个具有单个字段（名为 'Replace'）的 JSON 结构，
    			将每个磁盘文件路径（字符串）映射到其支持文件路径，
    			以便构建将以磁盘文件路径存在的内容运行，
    			给定由支持文件路径提供，或者如果支持文件路径为空，则表示磁盘文件路径不存在。
    			-overlay 标志的支持存在一些限制：
    			重要的是，从包含路径外部包含的 cgo 文件必须与它们所包含的 Go 包
    			相同目录中，覆盖将不会出现在通过 go run 和 go test 分别运行二进制文件和测试时。
        -pgo file
                specify the file path of a profile for profile-guided optimization (PGO).
                When the special name "auto" is specified, for each main package in the
                build, the go command selects a file named "default.pgo" in the package's
                directory if that file exists, and applies it to the (transitive)
                dependencies of the main package (other packages are not affected).
                Special name "off" turns off PGO. The default is "auto".
                指定用于基于配置文件的优化（PGO）的配置文件路径。
    			当指定特殊名称 "auto" 时，对于构建中的每个主包，
    			如果包的目录中存在名为 "default.pgo" 的文件，则 go 命令将选择该文件，
    			并将其应用于（传递）主包的依赖项（其他包不受影响）。
    			特殊名称 "off" 关闭 PGO。默认为 "auto"。
        -pkgdir dir
                install and load all packages from dir instead of the usual locations.
                For example, when building with a non-standard configuration,
                use -pkgdir to keep generated packages in a separate location.
                从 dir 而不是通常的位置安装和加载所有包。
    			例如，在使用非标准配置构建时，
    			使用 -pkgdir 将生成的包保留在单独的位置。
        -tags tag,list
                a comma-separated list of additional build tags to consider satisfied
                during the build. For more information about build tags, see
                'go help buildconstraint'. (Earlier versions of Go used a
                space-separated list, and that form is deprecated but still recognized.)
                一组用逗号分隔的额外构建标签，以在构建过程中考虑满足的。
    			有关构建标签的更多信息，请参见 'go help buildconstraint'。
    			（Go 的早期版本使用空格分隔的列表，这种形式已被弃用，但仍然被识别。）
        -trimpath
                remove all file system paths from the resulting executable.
                Instead of absolute file system paths, the recorded file names
                will begin either a module path@version (when using modules),
                or a plain import path (when using the standard library, or GOPATH).
                从结果可执行文件中删除所有文件系统路径。
    			记录的文件名将以模块路径@版本开头（当使用模块时），
    			或者以纯粹的导入路径开头（当使用标准库或 GOPATH 时）。
        -toolexec 'cmd args'
                a program to use to invoke toolchain programs like vet and asm.
                For example, instead of running asm, the go command will run
                'cmd args /path/to/asm <arguments for asm>'.
                The TOOLEXEC_IMPORTPATH environment variable will be set,
                matching 'go list -f {{.ImportPath}}' for the package being built.
                用于调用工具链程序（如 vet 和 asm）的程序。
    			例如，不是运行 asm，而是运行 'cmd args /path/to/asm <asm 的参数>'。
    			TOOLEXEC_IMPORTPATH 环境变量将被设置，
    			匹配正在构建的包的 'go list -f {{.ImportPath}}'。

The -asmflags, -gccgoflags, -gcflags, and -ldflags flags accept a space-separated list of arguments to pass to an underlying tool during the build. To embed spaces in an element in the list, surround it with either single or double quotes. The argument list may be preceded by a package pattern and an equal sign, which restricts the use of that argument list to the building of packages matching that pattern (see ‘go help packages’ for a description of package patterns). Without a pattern, the argument list applies only to the packages named on the command line. The flags may be repeated with different patterns in order to specify different arguments for different sets of packages. If a package matches patterns given in multiple flags, the latest match on the command line wins.

​	-asmflags、-gccgoflags、-gcflags 和 -ldflags 标志接受一个用空格分隔的列表，以便在构建期间将参数传递给底层工具。要在列表中的元素中嵌入空格，请用单引号或双引号将其括起来。参数列表前面可以加上包模式和等号，这会将该参数列表的使用限制为构建与该模式匹配的包（有关包模式的说明，请参阅“go help packages”）。如果没有模式，参数列表仅适用于命令行中命名的包。可以重复使用具有不同模式的标志，以便为不同的包集指定不同的参数。如果某个包与多个标志中给出的模式匹配，则命令行上的最新匹配获胜。

For example, ‘go build -gcflags=-S fmt’ prints the disassembly only for package fmt, while ‘go build -gcflags=all=-S fmt’ prints the disassembly for fmt and all its dependencies.

​	例如，“go build -gcflags=-S fmt”仅为包 fmt 打印反汇编，而“go build -gcflags=all=-S fmt”为 fmt 及其所有依赖项打印反汇编。

For more about specifying packages, see ‘go help packages’.

​	有关指定包的更多信息，请参阅“go help packages”。

For more about where packages and binaries are installed, run ‘go help gopath’.

​	有关包和二进制文件安装位置的更多信息，请运行“go help gopath”。

For more about calling between Go and C/C++, run ‘go help c’.

​	有关在 Go 和 C/C++ 之间进行调用的更多信息，请运行“go help c”。

Note: Build adheres to certain conventions such as those described by ‘go help gopath’. Not all projects can follow these conventions, however. Installations that have their own conventions or that use a separate software build system may choose to use lower-level invocations such as ‘go tool compile’ and ‘go tool link’ to avoid some of the overheads and design decisions of the build tool.

​	注意：构建遵循某些约定，例如“go help gopath”中描述的约定。但是，并非所有项目都能遵循这些约定。具有自己的约定或使用单独软件构建系统的安装可以选择使用更低级别的调用，例如“go tool compile”和“go tool link”，以避免构建工具的一些开销和设计决策。

See also: go install, go get, go clean.

​	另请参阅：go install、go get、go clean。