+++
title = "go"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# go

https://pkg.go.dev/cmd/go@go1.20.1

## Overview 

​	go 是管理 Go 源代码的工具。

用法：

```
go <command> [arguments]
```

命令包括：

```
bug         开始一个 bug 报告
build       编译包和依赖项
clean       删除对象文件和缓存文件
doc         显示包或符号的文档
env         打印 Go 环境信息
fix         更新包以使用新的 API
fmt         gofmt（重新格式化）包源代码
generate    通过处理源代码生成 Go 文件
get         添加依赖项到当前模块并安装它们
install     编译和安装包和依赖项
list        列出包或模块
mod         模块维护
work        工作区维护
run         编译并运行 Go 程序
test        测试包
tool        运行指定的 go 工具
version     打印 Go 版本
vet         报告包中可能存在的错误
```

使用"go help <command>"了解有关命令的更多信息。

其他帮助主题：

```
buildconstraint 构建限制条件
buildmode       构建模式
c               Go 和 C 之间的调用
cache 			构建和测试缓存
environment 	环境变量
filetype 		文件类型
go.mod 			go.mod 文件
gopath			GOPATH 环境变量
gopath-get      传统 GOPATH go get
goproxy 		模块代理协议
importpath 		导入路径语法
modules 		模块，模块版本等等
module-get 		支持模块的 go get
module-auth 	使用 go.sum 进行模块身份验证
packages	 	包列表和模式
private 		配置以下载非公共代码
testflag 		测试标志
testfunc 		测试函数
vcs 			使用 GOVCS 控制版本控制
```

使用"go help <topic>"了解该主题的更多信息。

#### go bug -> 开始一个 bug 报告

用法：

```
go bug
```

Bug 打开默认浏览器并启动一个新的 bug 报告。该报告包括有用的系统信息。

#### go build -> 编译包和依赖项  

用法：

```
go build [-o output] [build flags] [packages]
```

​	Build 编译由导入路径命名的包以及它们的依赖项，但不安装结果。

​	如果 build 的参数是来自单个目录的一组 .go 文件，则 build 将把它们视为指定单个包的源文件列表。

​	编译包时，build 忽略以"_test.go"结尾的文件。

​	当编译单个主包时，build 命令将生成的可执行文件写入以第一个源文件命名的输出文件（'go build ed.go rx.go' 会写入 'ed' 或 'ed.exe'），或者写入源代码目录（'go build unix/sam' 会写入 'sam' 或 'sam.exe'）。在 Windows 上编写可执行文件时，会添加 '.exe' 后缀。

​	当编译多个包或单个非主包时，build 命令会编译这些包，但会丢弃生成的对象，只用于检查这些包是否能够被构建。

​	-o 标志强制 build 命令将生成的可执行文件或对象写入命名的输出文件或目录，而不是采用最后两段所述的默认行为。如果命名的输出是一个现有目录或以斜线或反斜线结尾，则任何生成的可执行文件都将写入该目录。

​	构建标志由 build、clean、get、install、list、run 和 test 命令共享：

```
-C dir
	在运行命令之前切换到dir目录。
	任何在命令行上命名的文件在更改目录之后解释。
	
-a
	强制重新构建已经是最新版本的软件包。
	
-n
	打印命令但不运行它们。
	
-p n
	可以并行运行的程序数量，例如构建命令或测试二进制文件。
	默认值为GOMAXPROCS，通常为可用CPU的数量。
	
-race
	启用数据竞争检测。
	仅在linux/amd64、freebsd/amd64、darwin/amd64、darwin/arm64、	
    windows/amd64、linux/ppc64le和linux/arm64上支持（仅适用于48位VMA）。
    
-msan
	启用与内存污点检测器的互操作。
	仅在linux/amd64、linux/arm64、freebsd/amd64上支持，
	并且仅支持使用Clang/LLVM作为主机C编译器。
	PIE构建模式将在除linux/amd64外的所有平台上使用。
	
-asan
	启用与地址污点检测器的互操作。
	仅在linux/arm64、linux/amd64上支持。
	仅在linux/amd64或linux/arm64上支持，
	并且仅支持使用GCC 7及更高版本或Clang/LLVM 9及更高版本。
	
-cover
	启用代码覆盖率分析（需要设置GOEXPERIMENT=coverageredesign）。
	
-coverpkg pattern1，pattern2，pattern3
	针对目标为'main'的构建（例如构建Go可执行文件），
	将覆盖分析应用于与模式匹配的每个包。
	默认情况下，将覆盖分析应用于主Go模块中的包。
	有关包模式的说明，请参见"go help packages"。设置-cover。
-v
	编译软件包时打印软件包名称。
	
-work
	打印临时工作目录的名称，并在退出时不删除它。
	
-x
	打印命令。

-asmflags '[pattern=]arg list'
	在每个go工具asm调用中传递的参数。
	
-buildmode mode
	要使用的构建模式。有关更多信息，请参见"go help buildmode"。
	
-buildvcs
	是否在二进制文件中打印版本控制信息("true"、"false"或"auto")。
	默认情况下（"auto"），如果主包、包含它的主模块和当前目录都在同一个仓库中，
	则将版本控制信息打印到二进制文件中。
	使用-buildvcs=false始终省略版本控制信息，
	或者使用-buildvcs=true，
	如果版本控制信息可用但由于缺少工具或模糊的目录结构无法包含，则出错。
-compiler name
	指定要使用的编译器名称，如 runtime.Compiler 中的 gccgo 或 gc。
	
-gccgoflags '[pattern=]arg list'
	每个 gccgo 编译器/链接器调用传递的参数列表。
	
-gcflags '[pattern=]arg list'
	每个 go 工具编译调用传递的参数列表。
	
-installsuffix suffix
	用于包安装目录名称的后缀，以使输出与默认构建分开。
	如果使用 -race 标志，
	则自动将安装后缀设置为 race 或在显式设置后附加 _race。
	-msan 和 -asan 标志同理。
	使用需要非默认编译标志的 -buildmode 选项具有类似的效果。
	
-ldflags '[pattern=]arg list'
	每个 go 工具链接调用传递的参数列表。
	
-linkshared
	构建将链接到以 -buildmode=shared 创建的共享库的代码。
	
-mod mode
	要使用的模块下载模式：readonly、vendor 或 mod。
	默认情况下，如果存在 vendor 目录并且 go.mod 中的 go 版本为 1.14 或更高版本，
	则 go 命令会像设置了 -mod=vendor 一样操作。
	否则，go 命令会像设置了 -mod=readonly 一样操作。
	有关详细信息，请参见 https://golang.org/ref/mod#build-commands。
	
-modcacherw
	将新创建的目录保留在模块缓存中，以便进行读写，而不是只读。
	
-modfile file
	在模块感知模式下，读取（并可能写入）替代 go.mod 文件，
	而不是在模块根目录中的文件。
	仍然必须存在名为"go.mod"的文件，以确定模块根目录，
	但不会访问该文件。指定 -modfile 时，
	还会使用替代 go.sum 文件：
	其路径是通过从 -modfile 标志中删除".mod"扩展名并附加".sum"来派生的。
	
-overlay file
	读取 JSON 配置文件，为构建操作提供覆盖。
	文件是一个 JSON 结构，具有一个名为 'Replace' 的字段，
	该字段将每个磁盘文件路径（一个字符串）映射到其支持文件路径，
	以便在运行构建时，就像磁盘文件路径存在并具有由支持文件路径给定的内容一样，
	或者如果其支持文件路径为空，则磁盘文件路径将不存在。
	-overlay 标志的支持有一些限制：重要的是，
	从外部包含的 cgo 文件必须与它们所包含的 Go 包在同一个目录中，
	而覆盖在通过 go run 和 go test 运行二进制文件和测试时将不会出现。
	
-pgo file
	指定用于编译时的基于概要文件的优化（PGO）的文件路径。
	特殊名称"auto"将允许 go 命令在主包目录中选择名为"default.pgo"的文件
	（如果存在）。
	特殊名称"off"将关闭 PGO。
	
-pkgdir dir
	安装和从 dir 加载所有包，而不是使用通常的位置。
	例如，在使用非标准配置进行构建时，使用 -pkgdir 将生成的包保留在单独的位置。
	
-tags tag，list
	一个逗号分隔的构建标记列表，表示要在构建期间考虑的额外构建标记。
	有关构建标记的更多信息，请参见"go help buildconstraint"。
	（Go 的早期版本使用空格分隔的列表，虽然这种形式已被弃用但仍然可以识别。）
	
-trimpath
	从生成的可执行文件中删除所有文件系统路径。
	记录的文件名将以模块路径@版本（在使用模块时）或普通的导入路径
	（在使用标准库或 GOPATH 时）开头。
	
-toolexec 'cmd args'
	用于调用类似 vet 和 asm 的工具链程序的程序。
	例如，可以使用 -toolexec 来运行 asm 而不是直接运行，
	go 命令将运行"cmd args /path/to/asm <arguments for asm>"。
	TOOLEXEC_IMPORTPATH 环境变量将被设置，
	与正在构建的包的"go list -f {{.ImportPath}}"匹配。
```

​	`-asmflags`、`-gccgoflags`、`-gcflags` 和 `-ldflags` 标志接受一个以空格分隔的参数列表，用于在构建期间传递给底层工具。要在列表中的元素中嵌入空格，请用单引号或双引号括起来。参数列表可以用包模式和等号开头，限制该参数列表仅适用于构建与该模式匹配的包（有关包模式的描述，请参见"go help packages"）。没有模式时，参数列表仅适用于命令行上指定的包。这些标志可以重复使用不同的模式，以便为不同的包集指定不同的参数。如果一个包匹配在多个标志中给定的模式，命令行上最后匹配的标志将覆盖之前的所有标志。例如，"go build -gcflags=-S fmt"仅为包fmt打印反汇编，而"go build -gcflags=all=-S fmt"为fmt及其所有依赖项打印反汇编。

​	有关包的详细信息，请参见"go help packages"。有关安装包和二进制文件的位置，请运行"go help gopath"。有关在 Go 和 C/C++ 之间调用的更多信息，请运行"go help c"。

​	注意：构建遵循某些约定，如"go help gopath"所述。然而，并非所有项目都能遵循这些约定。具有自己的约定或使用单独的软件构建系统的安装可能选择使用较低级别的调用，如"go tool compile"和"go tool link"，以避免构建工具的一些开销和设计决策。

​	另请参阅：go install、go get、go clean。

#### go clean -> 删除对象文件和缓存文件

用法：

```
go clean [clean flags] [build flags] [packages]
```

​	clean从包源目录中删除对象文件。go命令在临时目录中构建大多数对象，因此go clean主要涉及其他工具或手动调用go build留下的对象文件。

​	如果给出了包参数或设置了-i或-r标志，则clean会从每个对应于导入路径的源目录中删除以下文件：

```
_obj/            旧的对象目录，由Makefile留下
_test/           旧的测试目录，由Makefile留下
_testmain.go     旧的 gotest 文件，由 Makefile 留下
test.out         旧的测试日志，由 Makefile 留下
build.out        旧的测试日志，由 Makefile 留下
*.[568ao]        由 Makefile 留下的对象文件

DIR(.exe)        通过 go build 生成的可执行文件
DIR.test(.exe)   通过 go test -c 生成的测试可执行文件
MAINFILE(.exe)   通过 go build MAINFILE.go 生成的可执行文件
*.so             由 SWIG 生成的文件
```

​	在列表中，DIR 表示目录的最终路径元素，而 MAINFILE 是目录中未包含在构建包中的任何 Go 源文件的基本名称。

​	`-i` 标志会使 clean 删除相应已安装的归档文件或二进制文件（相当于 'go install'）。

​	`-n` 标志会使 clean 打印出它将执行的删除命令，但不实际运行它们。

​	`-r` 标志会使 clean 递归应用于被导入路径命名的包的所有依赖项。

​	`-x` 标志会使 clean 打印出它执行删除命令的同时。

​	`-cache` 标志会使 clean 删除整个 go 构建缓存。

​	`-testcache` 标志会使 clean 过期 go 构建缓存中的所有测试结果。

​	`-modcache` 标志会使 clean 删除整个模块下载缓存，包括版本化依赖项的解压缩源代码。

​	`-fuzzcache` 标志会使 clean 删除用于模糊测试的存储在 Go 构建缓存中的文件。模糊引擎会缓存扩展代码覆盖率的文件，因此删除它们可能会使模糊测试变得不太有效，直到找到提供相同覆盖率的新输入。这些文件与存储在 testdata 目录中的文件不同；clean 不会删除这些文件。

​	有关构建标志的更多信息，请参阅 'go help build'。

​	有关指定包的更多信息，请参阅 'go help packages'。

#### go doc -> 显示包或符号的文档

用法：

```
go doc [doc flags] [package|[package.]symbol[.methodOrField]]
```

​	doc 打印与其参数所标识的项相关联的文档注释（一个包、常量、函数、类型、变量、方法或结构体字段），后跟每个该项"下面"一级项（一个包级别的声明、类型的方法等）的一行摘要。

​	doc 接受零个、一个或两个参数。

​	当不带参数运行时，即

```
go doc
```

它会打印当前目录中包的包文档。如果包是一个命令（package main），则除非提供了 -cmd 标志，否则该包的导出符号将从演示中省略。

​	当带有一个参数运行时，参数会被视为要文档化的项的 Go 语法样式表示形式。参数选择的内容取决于 GOROOT 和 GOPATH 中安装的内容，以及参数的形式，其概略如下：

```
go doc <pkg>
go doc <sym>[.<methodOrField>]
go doc [<pkg>.]<sym>[.<methodOrField>]
go doc [<pkg>.][<sym>.]<methodOrField>
```

​	列表中第一个匹配参数的项目是其文档将被打印的项目。（见下面的示例。）然而，如果参数以大写字母开头，则假定其为当前目录中标识符或方法。

​	对于包，扫描的顺序是按广度优先的词法顺序确定的。也就是说，呈现的包是与搜索匹配并且在其层次结构的根和词法上最先的包。在扫描GOPATH之前，始终完整扫描GOROOT树。

​	如果没有指定或匹配包，则选择当前目录中的包，因此"go doc Foo"会显示当前包中符号Foo的文档。

​	包路径必须是一个合格的路径或路径的后缀。go工具的常规包机制不适用于go doc。例如，包路径元素（如"."和"…"）没有被实现。

​	当使用两个参数运行时，第一个参数是包路径（完整路径或后缀），第二个是符号或具有方法或结构字段的符号：

```
go doc <pkg> <sym>[.<methodOrField>]
```

​	在所有形式中，当匹配符号时，参数中的小写字母匹配任何情况，但大写字母匹配确切的情况。这意味着，如果不同的符号具有不同的情况，则小写参数在包中可能有多个匹配项。如果出现这种情况，则打印所有匹配项的文档。

示例：

```sh
go doc
	显示当前包的文档。
	
go doc Foo
	显示当前包中Foo的文档。
	(Foo以大写字母开头，因此不能匹配包路径。)
	
go doc encoding/json
	显示encoding/json包的文档。
	
go doc json
	encoding/json的简写。
	
go doc json.Number (or go doc json.number)
	显示json.Number的文档和方法摘要。
	
go doc json.Number.Int64 (or go doc json.number.int64)
	显示json.Number的Int64方法的文档。
	
go doc cmd/doc
	显示doc命令的包文档。
	
go doc -cmd cmd/doc
	显示doc命令中的包文档和导出符号。
	
go doc template.new
	显示html/template的New函数的文档。
	(html/template在text/template之前按字典顺序排列)
	
go doc text/template.new # One argument
	显示text/template的New函数的文档。

go doc text/template new # Two arguments
	显示text/template的New函数的文档。

至少在当前树中，这些调用都打印json.Decoder的Decode方法的文档：

go doc json.Decoder.Decode
go doc json.decoder.decode
go doc json.decode
cd go/src/encoding/json; go doc decode
```

标志：

```
-all
	显示包中的所有文档。
	
-c
	在匹配符号时区分大小写。
	
-cmd
	将一个命令（package main）视为常规包。
	否则，显示包的顶层文档时将隐藏 package main 的导出符号。
	
-short
	每个符号显示一行的简要表示。
	
-src
	显示符号的完整源代码。
	这将显示其声明和定义的完整 Go 源代码，
	例如函数定义（包括主体）、类型声明或封闭 const 块。
	因此，输出可能包括未导出的详细信息。
	
-u
	显示未导出的符号、方法和字段的文档，以及导出的文档。
```

#### 打印Go环境信息

用法：

```
go env [-json] [-u] [-w] [var ...]
```

​	env打印Go环境信息。

​	默认情况下，env以shell脚本形式打印信息（在Windows上为批处理文件）。如果给出一个或多个变量名作为参数，env会在自己的行上打印每个命名变量的值。

​	`-json`标志以JSON格式而不是作为shell脚本打印环境。

​	`-u`标志需要一个或多个参数，并取消命名环境变量的默认设置，如果已使用"go env -w"设置。

​	`-w`标志需要一个或多个名称=值形式的参数，并将命名环境变量的默认设置更改为给定值。

​	有关环境变量的更多信息，请参见"go help environment"。

#### go fix -> 更新包以使用新API 

用法：

```
go fix [-fix list] [packages]
```

​	fix在导入路径命名的包上运行Go fix命令。

​	`-fix`标志设置要运行的逗号分隔的修复程序列表。默认为所有已知修复程序。 （其值传递给"go tool fix -r"。）

​	有关修复程序的更多信息，请参见"go doc cmd/fix"。有关指定包的更多信息，请参见"go help packages"。

​	要使用其他选项运行fix，请运行"go tool fix"。

​	另请参见：go fmt，go vet。

#### go fmt -> Gofmt（重新格式化）包源

用法：

```
go fmt [-n] [-x] [packages]
```

​	fmt在导入路径命名的包上运行命令'gofmt -l -w'。它打印已修改的文件的名称。

​	有关gofmt的更多信息，请参见"go doc cmd/gofmt"。有关指定包的更多信息，请参见"go help packages"。

​	`-n`标志打印将被执行的命令。-x标志按它们被执行的方式打印命令。

​	`-mod`标志的值设置要使用的模块下载模式：readonly或vendor。有关更多信息，请参见"go help modules"。

​	要使用特定选项运行gofmt，请运行gofmt本身。

​	另请参见：go fix，go vet。

#### go generate 通过处理源文件生成Go文件

用法：

```
go generate [-run regexp] [-n] [-v] [-x] [build flags] [file.go... | packages]
```

​	generate 运行由现有文件中的指令描述的命令。这些命令可以运行任何进程，但是其意图是创建或更新 Go 源文件。

​	go generate 不会被 go build、go test 等自动运行。必须显式运行它。

​	go generate 扫描指令文件，这些文件的指令是一行文本，格式如下：

```
//go:generate command argument...
```

(注意：没有前导空格，也没有 "//go" 中的空格)，其中 command 是要运行的生成器，对应于可以在本地运行的可执行文件。它必须在 shell 路径（gofmt）、完全限定路径（/usr/you/bin/mytool）或命令别名中。

​	请注意，go generate 不会解析文件，因此看起来像指令的行注释或多行字符串将被视为指令。

​	指令的参数是空格分隔的标记或双引号括起来的字符串，它们作为单独的参数传递给生成器在运行时。

​	引号括起来的字符串使用 Go 语法，并在执行之前进行评估；引号括起来的字符串在生成器中出现为单个参数。

​	为了让人类和机器工具知道代码是由生成器生成的，生成的源代码应该具有与以下正则表达式匹配的行（使用 Go 语法）：

```
^// Code generated .* DO NOT EDIT\.$
```

该行必须出现在文件中第一个非注释、非空白文本之前。

​	在运行生成器时，go generate 设置了几个变量：

```
$GOARCH
	执行的体系结构（arm、amd64 等）。
$GOOS
	执行的操作系统（linux、windows 等）。
$GOFILE
	文件的基本名称。
$GOLINE
	源文件中指令的行号。
$GOPACKAGE
	包含指令的文件的包的名称。
$GOROOT
	调用生成器的 'go' 命令的 GOROOT 目录，其中包含 Go 工具链和标准库。
$DOLLAR
	一个美元符号。
```

​	除了变量替换和引号括起来的字符串评估之外，命令行不执行任何特殊处理，例如 "globbing"。

​	作为执行命令前的最后一步，任何具有字母数字名称的环境变量调用（如`$GOFILE`或`$HOME`）都会在整个命令行中展开。变量扩展的语法在所有操作系统上均为`$NAME`。由于计算顺序，即使在引号内，变量也会被展开。如果未设置变量NAME，则`$NAME`会展开为空字符串。

​	以下是一个指令示例：

```
//go:generate -command xxx args...
```

该指令指定，在此源文件中，字符串xxx表示由参数标识的命令。这可用于创建别名或处理多个单词的生成器。例如：

```
//go:generate -command foo go tool foo
```

该指令指定命令"foo"表示生成器"go tool foo"。

​	在命令行中按照给定的顺序，逐一处理包，如果命令行列出了单个目录中的.go文件，则它们被视为单个包。在一个包内，按文件名顺序逐一处理源文件。在一个源文件内，生成器按它们出现在文件中的顺序逐一运行。go generate工具还设置了构建标记"generate"，因此文件可以通过go generate进行检查，但在构建过程中会被忽略。

​	对于包含无效代码的包，generate仅处理具有有效包从句的源文件。

​	如果任何生成器返回错误的退出状态，则"go generate"跳过该包的所有后续处理。

​	生成器在包的源目录中运行。

​	go generate接受两个特定的标志：

```
-run=""
	如果非空，则指定一个正则表达式，
	以选择原始源文本（不包括任何尾随空格和最后一个换行符）与表达式匹配的指令。

-skip=""
	如果非空，则指定一个正则表达式，
	以抑制原始源文本（不包括任何尾随空格和最后一个换行符）与表达式匹配的指令。如果一个指令同时与-run和-skip参数匹配，则它将被跳过。

```

​	它还接受标准构建标志，包括-v、-n和-x。-v标志会在处理过程中打印包和文件的名称。-n标志会打印将要执行的命令。-x标志会打印正在执行的命令。

​	有关构建标志的更多信息，请参见"go help build"。

​	有关指定包的更多信息，请参见"go help packages"。

#### go get -> 添加依赖项到当前模块并安装它们 

用法：

```
go get [-t] [-u] [-v] [build flags] [packages]
```

​	get 将命令行参数解析为特定模块版本的包，更新 go.mod 以要求这些版本，并将源代码下载到模块缓存中。

​	要为包添加依赖项或将其升级到最新版本：

```
go get example.com/pkg
```

​	要升级或降级特定版本的包：

```
go get example.com/pkg@v1.2.3
```

​	要删除对模块的依赖项并降级需要它的模块：

```
go get example.com/mod@none
```

​	有关详情，请参见 [Go模块参考中的go get命令](../../GoModulesReference/Module-awareCommands#go-get)。

​	在早期的 Go 版本中，"go get"用于构建和安装包。现在，"go get"专用于调整 go.mod 中的依赖项。"go install"可用于构建和安装命令。当指定版本时，"go install"以模块感知模式运行，并忽略当前目录中的 go.mod 文件。例如：

```
go install example.com/pkg@v1.2.3
go install example.com/pkg@latest
```

​	有关详情，请参见 'go help install' 或 [Go模块参考中的go install命令](../../GoModulesReference/Module-awareCommands#go-install)。

​	'go get' 接受以下标志：

- `-t` 标志指示 get 考虑构建命令行中指定的包的测试所需的模块。

- `-u` 标志指示 get 更新提供命令行中指定包的依赖项的模块以使用更高的次要或补丁版本。

- `-u=patch` 标志（不是 -u patch）也指示 get 更新依赖项，但将默认选择修补程序版本。

​			当 -t 和 -u 标志一起使用时，get 也会更新测试依赖项。

- `-x` 标志打印执行的命令。当直接从存储库下载模块时，这对于调试版本控制命令非常有用。

​	有关模块的更多信息，请参见 https://golang.org/ref/mod。

​	有关指定包的更多信息，请参见 'go help packages'。

​	本文描述了使用模块管理源代码和依赖项的 get 的行为。如果相反，go 命令在 GOPATH 模式下运行，则 get 的标志和效果的细节会改变，'go help get' 也会改变。请参阅 'go help gopath-get'。

​	另请参见：go build、go install、go clean、go mod。

#### go install -> 编译和安装包及其依赖项

用法：

```
go install [build flags] [packages]
```

​	install 编译并安装导入路径指定的包。

​	可执行文件将被安装到名为GOBIN的目录中。默认情况下，GOBIN环境变量为`$GOPATH/bin`或者`$HOME/go/bin`（如果GOPATH环境变量未设置）。`$GOROOT`中的可执行文件将被安装到`$GOROOT/bin`或`$GOTOOLDIR`中，而不是`$GOBIN`中。

​	如果参数有版本后缀（如`@latest`或`@v1.0.0`），则"go install"将在模块感知模式下进行构建，忽略当前目录或任何父目录中的go.mod文件。这对于安装可执行文件而不影响主模块的依赖项非常有用。为了消除构建中使用的模块版本的歧义，参数必须满足以下限制：

- 参数必须是包路径或包模式（带有"…"通配符）。它们不能是标准包（如fmt），元模式（std，cmd，all）或相对或绝对文件路径。
- 所有参数必须具有相同的版本后缀。不允许不同的查询，即使它们引用同一个版本。
- 所有参数必须引用同一模块中的相同版本的包。
- 包路径参数必须引用主包。模式参数只会匹配主包。
- 没有模块被视为"主"模块。如果命令行上的包所在的模块有go.mod文件，则该文件不得包含指令（replace和exclude），使其被解释为与主模块不同。该模块不得要求其自身的更高版本。
- 任何模块中均不使用供应商目录。 （供应商目录未包含在"go install"下载的模块zip文件中。）

​	如果参数没有版本后缀，则"go install"可以在模块感知模式或GOPATH模式下运行，这取决于GO111MODULE环境变量和是否存在go.mod文件。有关详细信息，请参见"go help modules"。如果启用了模块感知模式，则"go install"在主模块的上下文中运行。

​	在禁用模块感知模式时，非主要包将安装在目录`$GOPATH/pkg/$GOOS_$GOARCH`中。启用模块感知模式时，非主要包将被构建和缓存，但不会被安装。

​	在 Go 1.20 之前，标准库被安装到`$GOROOT/pkg/$GOOS_$GOARCH`中。从 Go 1.20 开始，标准库被构建和缓存，但不会被安装。设置 GODEBUG=installgoroot=all 可以恢复对`$GOROOT/pkg/$GOOS_$GOARCH`的使用。

​	有关构建标志的更多信息，请参见"go help build"。有关指定软件包的更多信息，请参见"go help packages"。

​	另请参阅：go build、go get、go clean。

#### go list -> 列出包或模块

用法：

```
go list [-f format] [-json] [-m] [list flags] [build flags] [packages]
```

​	list 命令将指定的包列出，每个包占一行。最常用的标志是 -f 和 -json，它们控制打印每个包时输出的格式。其他的 list 标志在下面有说明，它们控制更具体的细节。

​	默认输出显示包的导入路径：

```
bytes
encoding/json
github.com/gorilla/mux
golang.org/x/net/html
```

​	`-f` 标志指定了列表的替代格式，使用包模板语法。默认输出相当于 `-f '{{.ImportPath}}'`。传递给模板的结构体是：

``` go
type Package struct {
    Dir           string   // 包源文件所在的目录
    ImportPath    string   // 包在目录中的导入路径
    ImportComment string   // 包声明的导入注释中的路径
    Name          string   // 包名
    Doc           string   // 包的文档字符串
    Target        string   // 安装路径
    Shlib         string   // 包含该包的共享库（仅在使用 -linkshared 时设置）
    Goroot        bool     // 该包是否在 Go 根目录下？
    Standard      bool     // 该包是否是标准 Go 库的一部分？
    Stale         bool     // 对于该包，go install 是否会执行任何操作？
    StaleReason   string   // Stale==true 的原因说明
    Root          string   // 包所在的 Go 根目录或 Go path 目录
    ConflictDir   string   // 此目录遮盖了 $GOPATH 中的 Dir
    BinaryOnly    bool     // 仅限二进制包（不再支持）
    ForTest       string   // 该包仅供命名测试使用
    Export        string   // 包含导出数据的文件（使用 -export 时）
    BuildID       string   // 编译包的 build ID（使用 -export 时）
    Module        *Module  // 包所在模块的信息（如果有）（可能为 nil）
    Match         []string // 与此包匹配的命令行模式
    DepOnly       bool     // 该包仅作为依赖项，没有被显式列出

    // 源文件
    GoFiles         []string   // .go 源文件（不包括 CgoFiles、TestGoFiles 和 XTestGoFiles）
    CgoFiles        []string   // 导入了 "C" 的 .go 源文件

    CompiledGoFiles []string   // 向编译器展示的 .go 文件（使用 -compiled 时）
    IgnoredGoFiles  []string   //  因构建约束而被忽略的 .go 源文件
    IgnoredOtherFiles []string // 因构建约束而被忽略的非 .go 源文件
    CFiles          []string   // .c 源文件
    CXXFiles        []string   // .cc、.cxx 和 .cpp 源文件
    MFiles          []string   //  .m 源文件
    HFiles          []string   // .h、.hh、.hpp 和 .hxx 源文件
    FFiles          []string   // .f、.F、.for 和 .f90 Fortran 源文件
    SFiles          []string   //  .s 源文件
    SwigFiles       []string   //  .swig 文件
    SwigCXXFiles    []string   // .swigcxx 文件
    SysoFiles       []string   // 要添加到档案文件的 .syso 目标文件
    TestGoFiles     []string   // 包内的 _test.go 文件
    XTestGoFiles    []string   // 包外的 _test.go 文件

    // 嵌入式文件
    EmbedPatterns      []string // //go:embed 模式
    EmbedFiles         []string // 由 EmbedPatterns 匹配的文件
    TestEmbedPatterns  []string // TestGoFiles 中的 //go:embed 模式
    TestEmbedFiles     []string // 由 TestEmbedPatterns 匹配的文件
    XTestEmbedPatterns []string // XTestGoFiles 中的 //go:embed 模式
    XTestEmbedFiles    []string // 由 XTestEmbedPatterns 匹配的文件

    // Cgo 指令
    CgoCFLAGS    []string // cgo：C 编译器的标志
    CgoCPPFLAGS  []string // cgo：C 预处理器的标志
    CgoCXXFLAGS  []string // cgo：C++ 编译器的标志
    CgoFFLAGS    []string // cgo：Fortran 编译器的标志
    CgoLDFLAGS   []string // cgo：链接器的标志
    CgoPkgConfig []string // cgo：pkg-config 的名称

    // 依赖项信息
    Imports      []string          // 此包使用的导入路径
    ImportMap    map[string]string // 源导入到 ImportPath 的映射（省略标识条目）
    Deps         []string          // 所有（递归）导入的依赖项
    TestImports  []string          // TestGoFiles 中的导入项
    XTestImports []string          // XTestGoFiles 中的导入项

    // 错误信息
    Incomplete bool            // 此包或依赖项存在错误
    Error      *PackageError   // 加载包时出现的错误
    DepsErrors []*PackageError // 加载依赖项时出现的错误
}
```

​	存储在 vendor 目录中的包报告 ImportPath，该路径包括供应商目录的路径（例如，"d/vendor/p"而不是"p"），以便 ImportPath 唯一标识给定的包副本。Imports、Deps、TestImports 和 XTestImports 列表也包含这些扩展的导入路径。有关供应商的更多信息，请参见 golang.org/s/go15vendor。

​	如果有错误信息，则为：

``` go
type PackageError struct {
    ImportStack   []string // 从命令行命名的包到此包的最短路径
    Pos           string   // 错误的位置（如果存在，则为文件：行：列）
    Err           string   // 错误本身
}
```

​	模块信息是一个 Module 结构，定义在下面 list -m 的讨论中。

​	模板函数"join"调用 strings.Join。

​	模板函数"context"返回构建上下文，定义为：

``` go
type Context struct {
    GOARCH        string   // 目标架构
    GOOS          string   // 目标操作系统
    GOROOT        string   // Go 根目录
    GOPATH        string   // Go 路径
    CgoEnabled    bool     // 是否可以使用 cgo
    UseAllFiles   bool     // 使用文件，无论是否有 +build 行、文件名
    Compiler      string   // 在计算目标路径时要使用的编译器
    BuildTags     []string // 在 +build 行中匹配的构建约束
    ToolTags      []string // 工具链特定的构建约束
    ReleaseTags   []string // 当前版本兼容的版本
    InstallSuffix string   // 在安装目录的名称中使用的后缀
}
```

​	有关这些字段的含义的更多信息，请参阅 go/build 包的 Context 类型的文档。

​	`-json` 标志会导致以 JSON 格式而不是模板格式打印包数据。JSON 标志可以可选地与一组逗号分隔的所需字段名称一起提供，以输出这些所需字段。如果是这样，这些必需的字段将始终出现在 JSON 输出中，但其他字段可能被省略以节省计算 JSON 结构的工作。

​	`-compiled` 标志会导致 list 将 CompiledGoFiles 设置为提供给编译器的 Go 源文件。通常，这意味着它会重复列在 GoFiles 中的文件，然后还会添加通过处理 CgoFiles 和 SwigFiles 生成的 Go 代码。Imports 列表包含来自 GoFiles 和 CompiledGoFiles 的所有导入的并集。

​	`-deps`标志使列表不仅迭代命名包，还包括它们的所有依赖项。它以深度优先的后序遍历方式访问它们，这样一个包只有在所有依赖项之后才会列出。未在命令行中明确列出的包将具有DepOnly字段设置为true。

​	`-e`标志更改对错误包（无法找到或格式错误的包）的处理方式。默认情况下，list命令对于每个错误包在标准错误输出一个错误，并在通常的输出中省略这些包。使用-e标志时，list命令永远不会将错误打印到标准错误输出，并使用通常的输出处理错误包。错误的包将具有非空的ImportPath和非空的Error字段；其他信息可能存在或可能不存在（为零）。

​	`-export`标志使列表将导出字段设置为包含给定包的最新导出信息的文件的名称，并将BuildID字段设置为已编译包的构建ID。

​	`-find`标志使列表标识命名包但不解析它们的依赖项：Imports和Deps列表将为空。

​	`-test`标志不仅报告命名包，还报告它们的测试二进制文件（对于具有测试的包），以向源代码分析工具传达测试二进制文件的构造方式。测试二进制文件的报告导入路径是包的导入路径后跟一个".test"后缀，例如"math/rand.test"。在构建测试时，有时需要特别为该测试重新构建某些依赖项（最常见的是被测试的包本身）。为特定测试二进制文件重新编译的包的报告导入路径后跟一个空格和方括号中测试二进制文件的名称，例如"math/rand math/rand.test"或"regexp [sort.test]"。ForTest字段还设置为正在测试的包的名称（在前面的示例中为"math/rand"或"sort"）。

​	Dir，Target，Shlib，Root，ConflictDir和Export文件路径都是绝对路径。

​	默认情况下，GoFiles、CgoFiles等列表保存Dir中文件的名称（即相对于Dir的路径，而不是绝对路径）。当使用-compiled和-test标志时添加的生成文件是引用生成的Go源文件的缓存副本的绝对路径。虽然它们是Go源文件，但路径可能不以".go"结尾。

​	`-m`标志使列表列出模块而不是包。

​	当列出模块时，`-f`标志仍然指定应用于Go结构的格式模板，但现在是一个Module结构体：

``` go
type Module struct {
    Path       string        // 模块路径
    Query      string        // 对应于此版本的版本查询
    Version    string        // 模块版本
    Versions   []string      // 可用的模块版本
    Replace    *Module       // 被此模块替换
    Time       *time.Time    // 版本创建时间
    Update     *Module       // 可用更新（使用-u）
    Main       bool          // 是否为主模块？
    Indirect   bool          // 模块仅由主模块间接需要
    Dir        string        // 如果有的话，保存文件的本地副本的目录
    GoMod      string        // 描述模块的go.mod文件的路径（如果有）
    GoVersion  string        // 模块使用的Go版本
    Retracted  []string      // 撤回信息（使用-retracted或-u）
    Deprecated string        // 废弃消息（使用-u）
    Error      *ModuleError  // 加载模块时的错误
    Origin     any           // 模块来源
    Reuse      bool          // 旧模块信息的重用是安全的
}

type ModuleError struct {
    Err string // 错误本身
}
```

​	GoMod文件所指的文件可能在模块目录之外，如果模块在模块缓存中或使用了-modfile标志，则是如此。

​	默认输出是打印模块路径，然后是版本和替换信息（如果有）。例如，"go list -m all"可能会打印：

```
my/main/module
golang.org/x/text v0.3.0 => /tmp/text
rsc.io/pdf v0.1.1
```

​	Module结构具有String方法，用于格式化此行输出，因此默认格式等同于`-f'{{.String}}'`。

​	请注意，当模块已被替换时，其Replace字段描述替换模块，并且如果存在，则其Dir字段设置为替换的源代码。（也就是说，如果Replace不为nil，则Dir设置为Replace.Dir，没有访问被替换源代码的方式。）

​	`-u`标志会添加有关可用升级的信息。当给定模块的最新版本比当前版本更新时，list -u会将模块的Update字段设置为有关更高版本模块的信息。如果当前版本被撤回，则list -u还将设置模块的Retracted字段。如果可用升级，则模块的String方法通过在当前版本后用括号中的较新版本进行格式化来指示可用升级。如果版本被撤回，则字符串"（撤回）"将跟随它。例如，"go list -m -u all"可能会打印：

```
my/main/module
golang.org/x/text v0.3.0 [v0.4.0] => /tmp/text
rsc.io/pdf v0.1.1 (retracted) [v0.1.2]
```

(对于工具而言，'go list -m -u -json all' 更容易解析。)

​	`-versions` 标志会导致 list 将模块的 Versions 字段设置为该模块所有已知版本的列表，按照语义化版本控制的顺序从早到晚排序。该标志还更改默认输出格式，以模块路径为开头，后跟由空格分隔的版本列表。

​	`-retracted` 标志会导致 list 报告关于已撤销模块版本的信息。当 -retracted 与 -f 或 -json 结合使用时，Retracted 字段将被设置为一个字符串，解释为什么该版本被撤销。该字符串取自模块的 go.mod 文件中撤销指令的注释。当 -retracted 与 -versions 结合使用时，已撤销的版本会与未撤销的版本一起列出。-retracted 标志可以与或不带 -m 一起使用。

​	`list -m` 的参数被解释为模块列表，而不是包列表。主模块是包含当前目录的模块。活动模块是主模块和其依赖项。如果没有参数，list -m 会显示主模块。如果有参数，list -m 会显示由参数指定的模块。任何活动模块都可以用其模块路径来指定。特殊模式 "all" 指定所有活动模块，首先是主模块，然后是按模块路径排序的依赖项。包含 "..." 的模式指定其模块路径匹配该模式的活动模块。形式为 path@version 的查询指定该查询的结果，不限于活动模块。有关模块查询的详细信息，请参见 'go help modules'。

​	模板函数 "module" 接受一个字符串参数，必须是模块路径或查询，并将指定的模块返回为 Module 结构体。如果发生错误，结果将是一个带有非空错误字段的 Module 结构体。

​	使用 -m 时，-reuse=old.json 标志接受以前 'go list -m -json' 调用的 JSON 输出文件名，该调用具有相同的修改器标志集（例如 -u、-retracted 和 -versions）。go 命令可以使用此文件确定自上次调用以来模块未更改，并避免重新下载有关它的信息。未重新下载的模块将通过将 Reuse 字段设置为 true 在新输出中标记。通常，模块缓存会自动提供此类重用；-reuse 标志可用于不保留模块缓存的系统。

​	有关构建标志的更多信息，请参见 'go help build'。

​	有关指定包的更多信息，请参见 'go help packages'。

​	有关模块的更多信息，请参见 https://golang.org/ref/mod。

#### go mod -> 模块维护

​	go mod 提供了对模块操作的访问。

​	请注意，对模块的支持内置于所有 go 命令中，而不仅仅是 'go mod'。例如，日常添加、删除、升级和降级依赖关系应使用 'go get' 完成。有关模块功能的概述，请参阅 'go help modules'。

用法：

```
go mod <command> [arguments]
```

命令如下：

```
download 	将模块下载到本地缓存
edit 		从工具或脚本编辑 go.mod
graph 		打印模块依赖关系图
init 		在当前目录中初始化新模块
tidy 		添加缺少的模块并删除未使用的模块
vendor 		制作依赖关系的供应商副本
verify 		验证依赖项具有预期的内容
why 		解释需要哪些包或模块
```

使用 "go help mod <command>" 查看有关命令的更多信息。

#### go mod download -> 将模块下载到本地缓存

用法：

```
go mod download [-x] [-json] [-reuse=old.json] [modules]
```

​	download 命令会下载指定的模块，可以是选择主模块依赖项的模块模式，也可以是形式为 `path@version` 的模块查询。

​	如果没有指定参数，则 download 命令适用于构建和测试主模块中的包所需的模块：如果主模块处于 'go 1.17' 或更高版本，则为主模块明确要求的模块，否则为所有必需的传递模块（对于 'go 1.16' 或更低版本）。

​	在普通执行期间，go 命令将自动根据需要下载模块。"go mod download" 命令主要用于预先填充本地缓存或计算 Go 模块代理的答案。

​	默认情况下，download 不会向标准输出写入任何内容。它可能会将进度消息和错误打印到标准错误。

​	使用 -json 标志将导致 download 向标准输出打印一系列 JSON 对象，描述每个已下载的模块（或失败），对应于此 Go 结构体：

``` go
type Module struct {
    Path     string // 模块路径
    Query    string // 版本查询，对应于此版本
    Version  string // 模块版本
    Error    string // 加载模块时出现的错误
    Info     string // 缓存的 .info 文件的绝对路径
    GoMod    string // 缓存的 .mod 文件的绝对路径
    Zip      string // 缓存的 .zip 文件的绝对路径
    Dir      string // 缓存的源根目录的绝对路径
    Sum      string // 路径、版本的校验和（如 go.sum）
    GoModSum string // go.mod 的校验和（如 go.sum）
    Origin   any    // 模块来源的证明
    Reuse    bool   // 重用旧模块信息是安全的
}
```

​	`-reuse` 标记接受包含之前的 'go mod download -json' 调用的 JSON 输出的文件名。go 命令可以使用此文件确定模块是否自上次调用以来未更改并避免重新下载。不重新下载的模块将通过将 Reuse 字段设置为 true 来标记在新输出中。通常模块缓存会自动提供此类重用；-reuse 标记可以在不保留模块缓存的系统上很有用。

​	`-x` 标记会导致 download 打印出执行的命令。

​	有关"go mod download"的更多信息，请参见 https://golang.org/ref/mod#go-mod-download。

​	有关版本查询的更多信息，请参见 https://golang.org/ref/mod#version-queries。

#### go mod edit -> 从工具或脚本编辑 go.mod 

用法：

```
go mod edit [editing flags] [-fmt|-print|-json] [go.mod]
```

​	edit 提供了一个命令行界面，用于编辑 go.mod，主要供工具或脚本使用。它只读取 go.mod，不会查找有关所涉及的模块的信息。默认情况下，edit 读取和写入主模块的 go.mod 文件，但在编辑标记后可以指定不同的目标文件。

编辑标记指定一系列编辑操作。

​	`-fmt` 标记重新格式化 go.mod 文件而不进行其他更改。此重新格式化也隐含在使用或重写 go.mod 文件的任何其他修改中。只有在未指定其他标记时（例如 'go mod edit -fmt'）才需要此标记。

​	`-module` 标记更改模块的路径（go.mod 文件的模块行）。

​	`-require=path@version` 和 `-droprequire=path` 标志添加和删除给定模块路径和版本的要求。请注意，-require 覆盖路径上的所有现有要求。这些标志主要用于了解模块图的工具。用户应该优先使用 'go get path@version' 或 'go get path@none'，这些命令会根据其他模块的约束条件进行必要的 go.mod 调整。

​	`-exclude=path@version` 和 `-dropexclude=path@version` 标志添加和删除给定模块路径和版本的排除项。请注意，如果该排除项已存在，则 -exclude=path@version 不起作用。

​	`-replace=old[@v]=new[@v]` 标志添加给定模块路径和版本对的替换项。如果在 old@v 中省略 @v，则添加左侧没有版本的替换项，该替换项适用于该旧模块路径的所有版本。如果在 new@v 中省略 @v，则新路径应为本地模块根目录，而不是模块路径。请注意，-replace 覆盖了 old[@v] 的任何冗余替换项，因此省略 @v 将删除特定版本的现有替换项。

​	`-dropreplace=old[@v]` 标志删除给定模块路径和版本对的替换项。如果省略 @v，则删除左侧没有版本的替换项。

​	`-retract=version` 和 `-dropretract=version` 标志添加和删除给定版本的撤销。版本可以是单个版本，例如"v1.2.3"，也可以是封闭间隔，例如"[v1.1.0，v1.1.9]"。请注意，如果该撤回已经存在，则 -retract=version 不起作用。

​	`-require`、`-droprequire`、`-exclude`、`-dropexclude`、`-replace`、`-dropreplace`、`-retract` 和 `-dropretract` 编辑标志可以重复使用，更改按给定的顺序应用。

​	`-go=version` 标志设置预期的 Go 语言版本。

​	`-print` 标志以其文本格式打印最终的 go.mod，而不是将其写回 go.mod。

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

​	Retract 条目表示单个版本（而非区间），"Low" 和 "High" 字段将设置为相同的值。

​	请注意，这仅描述 go.mod 文件本身，而不是间接引用的其他模块。要查看可用于构建的完整模块集，请使用 'go list -m -json all'。

​	edit 还提供了 -C、-n 和 -x 构建标志。

​	有关 'go mod edit' 的更多信息，请参见 https://golang.org/ref/mod#go-mod-edit。

#### go mod graph -> 打印模块需求图 

用法：

```
go mod graph [-go=version] [-x]
```

​	graph 以文本形式打印模块需求图（已应用替换）。输出中的每一行都有两个空格分隔的字段：模块和其要求之一。除主模块外，每个模块都被标识为 path@version 格式的字符串，而主模块没有 @version 后缀。

​	`-go` 标志导致 graph 报告按照给定的 Go 版本加载的模块图，而不是 go.mod 文件中 'go' 指令指示的版本。

​	`-x` 标志导致 graph 打印 graph 执行的命令。

​	有关 'go mod graph' 的更多信息，请参见 https://golang.org/ref/mod#go-mod-graph。

#### go mod init -> 在当前目录中初始化新模块

用法：

```
go mod init [module-path]
```

​	init 在当前目录中初始化并写入新的 go.mod 文件，实际上创建一个以当前目录为根的新模块。go.mod 文件不能已经存在。

​	init 接受一个可选参数，新模块的模块路径。如果省略模块路径参数，则 init 将尝试使用 .go 文件中的导入注释、供应商工具配置文件（如 Gopkg.lock）和当前目录（如果在 GOPATH 中）来推断模块路径。

​	如果存在供应商工具的配置文件，则 init 将尝试从其导入模块要求。

​	有关 'go mod init' 的更多信息，请参见 https://golang.org/ref/mod#go-mod-init。

#### go mod tidy -> 添加缺失和删除未使用的模块 

用法：

```
go mod tidy [-e] [-v] [-x] [-go=version] [-compat=version]
```

​	tidy 确保 go.mod 与模块中的源代码匹配。它添加了构建当前模块的包和依赖项所需的任何缺失模块，并删除未提供任何相关包的未使用模块。它还添加任何缺失的条目到 go.sum，并删除任何不必要的条目。

​	`-v` 标志导致 tidy 打印有关已删除模块的信息到标准错误。

​	`-e` 标志导致 tidy 尝试在加载包时遇到错误时继续进行。

​	`-go` 标志会导致 tidy 更新 go.mod 文件中的 'go' 指令到给定的版本，这可能会改变哪些模块依赖项在 go.mod 文件中作为显式要求保留。（Go 版本 1.17 及更高版本会保留更多的要求以支持懒惰模块加载。）

​	`-compat` 标志保留所需的任何附加校验和，以便从指定的主 Go 发布版本成功加载模块图形，并且如果该 'go' 命令版本会从不同模块版本加载任何已导入的包，则会导致 tidy 出错。默认情况下，tidy 表现得好像将 -compat 标志设置为 go.mod 文件中指示的版本的前一个版本。

​	`-x` 标志会导致 tidy 打印下载命令执行的命令。

​	有关 'go mod tidy' 的更多信息，请参见 https://golang.org/ref/mod#go-mod-tidy。

#### go mod vendor -> 创建依赖项的副本以供vendor

用法：

```
go mod vendor [-e] [-v] [-o outdir]
```

​	vendor 重置主模块的供应商目录，以包括构建和测试所有主模块包所需的所有包。它不包括供应商软件包的测试代码。

​	`-v` 标志导致供应商将供应商的模块和包名称打印到标准错误输出。

​	`-e` 标志会导致供应商尝试在加载包时遇到错误后继续运行。

​	`-o` 标志会导致供应商在给定路径处创建供应商目录，而不是在 "vendor" 中。go 命令只能在模块根目录中命名为 "vendor" 的供应商目录，因此该标志主要对其他工具有用。

​	有关 'go mod vendor' 的更多信息，请参见 https://golang.org/ref/mod#go-mod-vendor。

#### go mod verify -> 验证依赖项是否具有预期内容

用法：

```
go mod verify
```

​	verify 检查当前模块的依赖项（存储在本地下载的源缓存中）是否在下载后已被修改。如果所有模块都未被修改，则 verify 打印 "all modules verified."。否则它会报告哪些模块已被更改，并导致 'go mod' 退出并显示一个非零状态码。

​	有关"go mod verify"的更多信息，请参见 https://golang.org/ref/mod#go-mod-verify。

#### go mod why -> 解释为什么需要包或模块

用法：

```
go mod why [-m] [-vendor] packages...
```

​	why 显示从主模块到每个列出的包的导入图中的最短路径。如果指定了 -m 标志，则 why 将参数视为模块列表，并为每个模块找到一个路径到任何包。

​	默认情况下，why 查询与 "go list all" 匹配的包图，其中包括可访问包的测试。-vendor 标志使 why 排除依赖项的测试。

​	输出是一系列段落，每个包或模块名称在命令行上都有一个段落，段落之间用空行分隔。每个段落以注释行 "# package" 或 "# module" 开头，给出目标包或模块。随后的行以一个包为一行，给出了导入图中的路径。如果从主模块中没有引用该包或模块，则段落将显示一个括号注释，指出这一点。

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

​	有关"go mod why"的更多信息，请参见https://golang.org/ref/mod#go-mod-why。

#### go work -> 工作区维护

​	work 提供了对工作区进行操作的访问。

​	请注意，对于工作区的支持内置于许多其他命令中，而不仅仅是 'go work'。

​	有关工作区的信息，请参见 'go help modules'。工作区是 Go 模块系统的一部分。

​	有关工作区的深入参考，请参见 https://go.dev/ref/mod#workspaces。

​	有关工作区的入门教程，请参见 https://go.dev/doc/tutorial/workspaces。

​	工作区由一个 go.work 文件指定，该文件使用 "use" 指令指定一组模块目录。这些模块由 go 命令用作构建和相关操作的根模块。未指定要使用的模块的工作区无法用于从本地模块进行构建。

​	go.work 文件是面向行的。每行包含一个指令，由关键字和参数组成。例如：

```
go 1.18

use ../foo/bar
use ./baz

replace example.com/foo v1.2.3 => example.com/bar v1.4.5
```

​	前导关键字可以从相邻行中分离出来以创建块，就像 Go 导入一样。

```
use (
  ../foo/bar
  ./baz
)
```

​	use 指令指定要包含在工作区主模块集中的模块。use 指令的参数是包含该模块的 go.mod 文件的目录。

​	go 指令指定了文件编写的 Go 版本。可能有将来的更改，这些更改可能由该版本控制工作区的语义，但是目前指定的版本没有影响。

​	replace 指令具有与 go.mod 文件中的 replace 指令相同的语法，并优先于 go.mod 文件中的替换。它主要用于覆盖不同工作区模块中的冲突替换。

​	要确定 go 命令是否在工作区模式下运行，请使用 "go env GOWORK" 命令。这将指定正在使用的工作区文件。

用法：

```
go work <command> [arguments]
```

命令为：

```
edit        从工具或脚本编辑 go.work
init        初始化工作区文件
sync        同步工作区构建列表到模块
use         向工作区文件中添加模块
```

​	有关命令的更多信息，请使用 "go help work <command>"。

#### go work edit -> 从工具或脚本编辑 go.work

用法：

```
go work edit [editing flags] [go.work]
```

​	edit 提供了一个命令行界面用于编辑 go.work 文件，主要供工具或脚本使用。它只读取 go.work 文件，不会查找有关模块的信息。如果未指定文件，则 Edit 会在当前目录和其父目录中查找 go.work 文件。

​	`editing`标志指定一系列编辑操作。

​	`-fmt` 标志重新格式化 go.work 文件，而不做其他更改。任何使用或重写 go.mod 文件的其他修改都会隐含此重新格式化。只有在没有指定其他标志的情况下才需要此标志，例如 go work edit -fmt。

​	`-use=path` 和 `-dropuse=path` 标志向 go.work 文件的模块目录集添加和删除 use 指令。

​	`-replace=old[@v]=new[@v]` 标志添加给定模块路径和版本对的替换。如果省略 old@v 中的 @v，则会添加左侧没有版本的替换，这适用于旧模块路径的所有版本。如果省略 new@v 中的 @v，则新路径应为本地模块根目录，而不是模块路径。请注意，-replace 会覆盖 old[@v] 的任何冗余替换，因此省略 @v 将删除特定版本的现有替换。

​	`-dropreplace=old[@v]` 标志删除给定模块路径和版本对的替换。如果省略 @v，则会删除左侧没有版本的替换。

​	`-use`、`-dropuse`、`-replace` 和 `-dropreplace` 编辑标志可以重复，更改按给定的顺序应用。

​	`-go=version` 标志设置期望的 Go 语言版本。

​	`-print` 标志以文本格式打印最终的 go.work，而不是将其写回 go.mod。

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

​	有关更多信息，请参见工作区参考：[Go模块参考中的工作区](../../GoModulesReference/Workspaces)。

#### go work init -> 初始化工作区文件

用法：

```
go work init [moddirs]
```

​	init 在当前目录中初始化并编写新的 go.work 文件，实际上创建了一个新的工作区。

​	go work init 可选地接受工作区模块的路径作为参数。如果省略参数，则将创建一个没有模块的空工作区。

​	每个路径参数都会添加到 go.work 文件的 use 指令中。当前的 Go 版本也将在 go.work 文件中列出。

​	有关更多信息，请参见工作区参考：[Go模块参考中的工作区](../../GoModulesReference/Workspaces)。

#### go work sync -> 同步工作区的构建清单到模块 

用法：

```
go work sync
```

​	sync命令将工作区的构建清单同步回工作区的模块

​	工作区的构建清单是所有（传递）依赖模块的版本集，这些依赖模块用于在工作区进行构建操作。 go work sync使用最小版本选择算法生成该构建清单，然后将这些版本与工作区指定的每个模块（使用指令）同步。

​	如果依赖模块的版本与构建清单的版本不同，则按顺序将工作区模块中指定的每个依赖模块升级到构建清单中的版本。请注意，最小版本选择保证构建清单中每个模块的版本始终相同或更高于每个工作区模块中的版本。

​	有关更多信息，请参见工作区参考：[Go模块参考中的工作区](../../GoModulesReference/Workspaces)。

#### go work use -> 将模块添加到工作区文件

用法：

```
go work use [-r] moddirs
```

​	use提供了一个命令行界面，用于将目录（可选地进行递归搜索）添加到go.work文件中。

​	如果go.work文件中存在，那么每个命令行上列出的参数目录都将在go.work文件中添加一个use指令，否则将从go.work文件中删除。

​	`-r`标志在参数目录中递归搜索模块，并且use命令操作方式与指定每个目录作为参数相同：即对于存在的目录将添加use指令，并且对于不存在的目录将删除use指令。

​	有关更多信息，请参见工作区参考：[Go模块参考中的工作区](../../GoModulesReference/Workspaces)。

#### go run -> 编译并运行Go程序 

用法：

```
go run [build flags] [-exec xprog] package [arguments...]
```

​	run编译并运行指定的Go 主包。通常，该程序被指定为来自单个目录的一系列.go源文件的列表，但也可以是导入路径、文件系统路径或与单个已知程序匹配的模式，例如"go run ."或"go run my/cmd"。

​	如果包参数有版本后缀（如`@latest`或`@v1.0.0`），"go run"将在模块感知模式下构建程序，忽略当前目录或任何父目录中的go.mod文件。这对于运行程序而不影响主模块的依赖项非常有用。

​	如果包参数没有版本后缀，"go run"可以在模块感知模式或GOPATH模式下运行，具体取决于GO111MODULE环境变量和go.mod文件的存在性。有关详细信息，请参见"go help modules"。如果启用了模块感知模式，"go run"将在主模块的上下文中运行。

​	默认情况下，"go run"直接运行已编译的二进制文件："a.out arguments..."。如果给出了"-exec"标志，"go run"将使用xprog调用二进制文件：

```
'xprog a.out arguments...'.
```

​	如果未给出"-exec"标志，且GOOS或GOARCH与系统默认值不同，并且在当前搜索路径中可以找到一个名为"`go_$GOOS_$GOARCH_exec`"的程序，例如"go_js_wasm_exec a.out arguments..."，则"go run"将使用该程序调用二进制文件。这使得可以使用模拟器或其他执行方法来执行跨编译程序。

​	默认情况下，"go run"编译二进制文件时不会生成调试器使用的信息，以减少构建时间。要在二进制文件中包含调试器信息，请使用"go build"。

​	run的退出状态不是已编译的二进制文件的退出状态。

​	有关构建标志的更多信息，请参见"go help build"。有关指定软件包的更多信息，请参见"go help packages"。

​	另请参见：go build。

#### go test -> 测试包 

用法：

```
go test [build/test flags] [packages] [build/test flags & test binary flags]
```

​	'go test' 自动测试由导入路径指定的包。它以以下格式打印测试结果的摘要：

```
ok   archive/tar   0.011s
FAIL archive/zip   0.022s
ok   compress/gzip 0.033s
...
```

之后是每个失败包的详细输出。

​	'go test' 重新编译每个包及其与文件名匹配的 "`*_test.go`" 文件。这些附加文件可以包含测试函数、基准函数、模糊测试和示例函数。有关更多信息，请参见 'go help testfunc'。每个列出的包都会导致执行一个单独的测试二进制文件。以"`_`"（包括"`_test.go`"）或"`.`"开头的文件将被忽略。

​	声明具有后缀 "_test" 的包的测试文件将被编译为单独的包，然后与主测试二进制文件链接和运行。

​	go 工具将忽略名为 "testdata" 的目录，使其可用于保存测试所需的辅助数据。

​	在构建测试二进制文件的过程中，go test 运行 go vet 命令，对包及其测试源文件进行静态分析，以识别重要的问题。如果 go vet 发现任何问题，go test 将报告这些问题并不运行测试二进制文件。仅使用默认的一组高置信度的 go vet 检查，包括 'atomic'、'bool'、'buildtags'、'errorsas'、'ifaceassert'、'nilfunc'、'printf' 和 'stringintconv'。您可以通过 "go doc cmd/vet" 查看这些和其他 vet 测试的文档。要禁用 go vet 的运行，请使用 -vet=off 标志。要运行所有检查，请使用 -vet=all 标志。

​	所有测试输出和摘要行都会打印到 go 命令的标准输出中，即使测试已将它们打印到自己的标准错误中。（go 命令的标准错误用于打印构建测试时的错误。）

go test 有两种不同的运行模式：

​	第一种被称为本地目录模式，当 go test 未使用包参数调用时（例如 'go test' 或 'go test -v'），它将编译当前目录中找到的包源和测试，然后运行生成的测试二进制文件。在这种模式下，禁用缓存（下面讨论）。包测试完成后，go test 打印一个总结行，显示测试状态（'ok' 或 'FAIL'）、包名和运行时间。

​	第二种情况称为包列表模式，当使用显式包参数调用go test时发生（例如'go test math'、'go test ./...'和'go test .'）。在此模式下，go test编译并测试命令行上列出的每个包。如果包测试通过，go test仅打印最终的"ok"摘要行。如果包测试失败，go test会打印完整的测试输出。如果使用-bench或-v标志调用，则为了显示请求的基准结果或详细日志记录，go test即使对于通过的包测试也会打印完整输出。在列出的所有包的包测试完成并输出后，如果任何包测试失败，go test将打印最终的"FAIL"状态。

​	仅在包列表模式下，go test会缓存成功的包测试结果，以避免重复运行测试。当测试结果可以从缓存中恢复时，go test会重新显示先前的输出，而不是再次运行测试二进制文件。当发生这种情况时，go test会在摘要行中的经过时间的位置打印"(cached)"。

​	缓存匹配的规则是运行涉及相同的测试二进制文件，并且命令行上的标志完全来自"可缓存"测试标志的受限集合，这些标志被定义为-benchtime、-cpu、-list、-parallel、-run、-short、-timeout、-failfast和-v。如果go test运行有任何测试或非测试标志在此集合之外，结果将不会被缓存。要禁用测试缓存，请使用除可缓存标志之外的任何测试标志或参数。显式禁用测试缓存的惯用方法是使用-count = 1。在包的源根目录（通常是$GOPATH）中打开文件或仅查看环境变量的测试将只匹配未来运行的情况，其中文件和环境变量未更改。缓存的测试结果被视为根本没有执行时间，因此成功的包测试结果将被缓存和重复使用，无论-timeout设置如何。

​	除了构建标志之外，'go test'本身处理的标志为：

```
-args
    将命令行剩余部分（-args后面的所有内容）传递给测试二进制文件，
    未经解释且不变。由于此标志消耗了命令行的剩余部分，
    因此包列表（如果存在）必须出现在此标志之前。

-c
    编译测试二进制文件到pkg.test，
    但不运行它（其中pkg是包导入路径的最后一个元素）。
    可以使用-o标志更改文件名。

-exec xprog
    使用xprog运行测试二进制文件。
    行为与"go run"相同。有关详细信息，请参见'go help run'。

-json
    将测试输出转换为适用于自动化处理的JSON格式。
有关编码细节，请参阅"go doc test2json"。

-o file
    将测试二进制文件编译到指定文件中。
	测试仍然运行（除非指定了-c或-i）。
```

​	测试二进制文件还接受控制测试执行的标志；这些标志也可以由"go test"访问。有关详细信息，请参阅"go help testflag"。

​	有关构建标志的更多信息，请参阅"go help build"。有关指定软件包的更多信息，请参阅"go help packages"。

​	另请参见：go build、go vet。

#### go tool -> 运行指定的go工具

用法：

```
go tool [-n] command [args...]
```

​	tool运行由参数标识的go工具命令。如果没有参数，则打印已知工具列表。

​	`-n`标志导致tool打印将要执行的命令，但不执行它。

​	有关每个工具命令的详细信息，请参见"go doc cmd/<command>"。

#### go version -> 打印Go版本 

用法：

```
go version [-m] [-v] [file ...]
```

​	version 打印Go二进制文件的构建信息。

​	go version 报告用于构建每个命名文件的Go版本。

​	如果命令行中没有命名文件，则go version打印自己的版本信息。

​	如果指定了目录，则go version递归地遍历该目录，寻找已知的Go二进制文件并报告它们的版本。默认情况下，go version不会报告目录扫描期间发现的无法识别的文件。-v标志会导致其报告无法识别的文件。

​	`-m`标志使go version在可能的情况下打印每个文件的嵌入式模块版本信息。在输出中，模块信息由版本行以下的多行组成，每行前面都有一个制表符。

​	另请参见：go doc runtime/debug.BuildInfo。

#### go vet -> 报告软件包中的可能错误 

用法：

```
go vet [-C dir] [-n] [-x] [-vettool prog] [build flags] [vet flags] [packages]
```

​	vet在由导入路径命名的软件包上运行go vet命令。

​	有关vet及其标志的更多信息，请参见"go doc cmd/vet"。有关指定软件包的更多信息，请参见"go help packages"。有关检查器及其标志的列表，请参见"go tool vet help"。有关特定检查器（例如'printf'）的详细信息，请参见"go tool vet help printf"。

​	`-C`标志在运行"go vet"命令之前更改到dir。-n标志打印将要执行的命令。-x标志在执行命令时打印命令。

​	`-vettool=prog`标志选择具有替代或附加检查的不同分析工具。例如，可以使用以下命令构建和运行"shadow"分析器：

```
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
go vet -vettool=$(which shadow)
```

​	go vet支持的构建标志是控制软件包解析和执行的标志，例如-n、-x、-v、-tags和-toolexec。有关这些标志的更多信息，请参见"go help build"。

​	另请参见：go fmt、go fix。

#### //go:build构建约束条件

​	构建约束条件（也称为构建标记）是决定一个文件是否应该被包含在包中的条件。构建约束条件是由以下行注释给出的：

```
//go:build
```

​	约束条件可以出现在任何类型的源文件（不仅仅是 Go 文件），但它们必须出现在文件的顶部附近，只有空行和其他行注释可以在它们之前。这些规则意味着在 Go 文件中，构建约束条件必须出现在包语句之前。

​	为了区分构建约束条件和包文档，构建约束条件应该在空行后跟随一个空行。

​	构建约束条件注释将被视为一个表达式，该表达式由 ||、&& 和 ! 运算符以及括号组合而成的构建标记组成。这些运算符的含义与 Go 中相同。

​	例如，以下构建约束条件约束了一个文件只有在满足"linux"和"386"约束条件，或者在满足"darwin"约束条件且不满足"cgo"约束条件时才能构建：

```
//go:build (linux && 386) || (darwin && !cgo)
```

​	一个文件如果有多于一个 //go:build 行注释，会引发错误。

​	在特定的构建过程中，以下构建标记将被满足：

- 目标操作系统，由 runtime.GOOS 拼写，使用 GOOS 环境变量设置。 

- 目标架构，由 runtime.GOARCH 拼写，使用 GOARCH 环境变量设置。

-  任何架构特性，采用 GOARCH.feature 的形式（例如，"amd64.v2"），如下所述。

-  如果 GOOS 是 Unix 或类 Unix 系统，则为"unix"。 

- 所使用的编译器，可以是"gc"或"gccgo"。 

- 如果支持 cgo 命令（请参阅"go help environment"中的 CGO_ENABLED），则为"cgo"。 

- 每个 Go 主要版本的术语，直到当前版本为止："go1.1"从 Go 1.1 开始，"go1.12"从 Go 1.12 开始，以此类推。 

- 通过 -tags 标记提供的任何其他标记（请参阅"go help build"）。

  

​	beta 版本或小版本没有单独的构建标记。

​	如果文件名在去掉扩展名和可能的 _test 后，匹配以下任一模式：

```
*_GOOS
*_GOARCH
*_GOOS_GOARCH
```

（例如：source_windows_amd64.go），其中 GOOS 和 GOARCH 分别表示任何已知的操作系统和架构值，则该文件被认为具有需要这些术语的隐式构建约束条件（除了文件中的任何显式约束条件）。

​	使用GOOS=android与GOOS=linux相匹配，以及android标记和文件。

​	使用GOOS=illumos与GOOS=solaris相匹配，并在illumos标记和文件上进行了补充。

​	使用GOOS=ios与GOOS=darwin相匹配，并在ios标记和文件上进行了补充。

​	已定义的架构特性构建标签为：

- 对于GOARCH=386，GO386=387和GO386=sse2，分别设置386.387和386.sse2构建标签。 
- 对于GOARCH=amd64，GOAMD64=v1、v2和v3对应于amd64.v1、amd64.v2和amd64.v3特性构建标签。
-  对于GOARCH=arm，GOARM=5、6和7对应于arm.5、arm.6和arm.7特性构建标签。 
- 对于GOARCH=mips或mipsle，GOMIPS=hardfloat和softfloat对应于mips.hardfloat和mips.softfloat（或mipsle.hardfloat和mipsle.softfloat）特性构建标签。 
- 对于GOARCH=mips64或mips64le，GOMIPS64=hardfloat和softfloat对应于mips64.hardfloat和mips64.softfloat（或mips64le.hardfloat和mips64le.softfloat）特性构建标签。 
- 对于GOARCH=ppc64或ppc64le，GOPPC64=power8、power9和power10对应于ppc64.power8、ppc64.power9和ppc64.power10（或ppc64le.power8、ppc64le.power9和ppc64le.power10）特性构建标签。
-  对于GOARCH=wasm，GOWASM=satconv和signext对应于wasm.satconv和wasm.signext特性构建标签。

​	对于GOARCH=amd64、arm、ppc64和ppc64le，特定特性级别也会设置先前所有级别的特性构建标签。例如，GOAMD64=v2设置了amd64.v1和amd64.v2特性标志。这确保使用v2特性的代码在引入GOAMD64=v4时继续编译。处理特定特性级别的缺失的代码应使用否定：

```
//go:build !amd64.v2
```

​	为了使文件不被考虑进行任何构建：

```
//go:build ignore
```

（任何其他未满足的单词都可以工作，但"ignore"是常规的。）

仅在使用cgo时并且仅在Linux和OS X上构建文件：

```
//go:build cgo && (linux || darwin)
```

这样的文件通常与实现其他系统的默认功能的另一个文件配对，该文件将带有以下限制：

```
//go:build !(cgo && (linux || darwin))
```

​	命名为dns_windows.go的文件将导致仅在为Windows构建包时包含它；同样，命名为math_386.s的文件将仅在为32位x86构建包时包含它。

​	Go版本1.16及更早版本使用不同的构建限制语法，使用"// +build"前缀。在遇到旧语法时，gofmt命令将添加等效的//go:build约束。

#### 构建模式 

​	'go build' 和 'go install' 命令可以使用 -buildmode 参数指定要构建哪种类型的目标文件。目前支持的值为：

```
-buildmode=archive
	将列出的非 main 包构建为 .a 文件。名为 main 的包将被忽略。

-buildmode=c-archive
	将列出的 main 包以及它导入的所有包构建为 C 归档文件。
	唯一可调用的符号将是那些使用 cgo 的 //export 注释导出的函数。
	需要列出一个 main 包。



-buildmode=c-shared
	将列出的 main 包以及它导入的所有包构建为 C 共享库。
	唯一可调用的符号将是那些使用 cgo 的 //export 注释导出的函数。
	需要列出一个 main 包。



-buildmode=default
	列出的 main 包构建为可执行文件，
	列出的非 main 包构建为 .a 文件（默认行为）。



-buildmode=shared
	将所有列出的非 main 包组合成一个共享库，
	将在使用 -linkshared 选项进行构建时使用。
	名为 main 的包将被忽略。



-buildmode=exe
	将列出的 main 包及其导入的所有内容构建为可执行文件。
	名为 main 以外的包将被忽略。



-buildmode=pie
	将列出的 main 包及其导入的所有内容构建为位置无关可执行文件（PIE）。
	名为 main 以外的包将被忽略。



-buildmode=plugin
	将列出的 main 包以及它们导入的所有包构建为 Go 插件。
	名为 main 以外的包将被忽略。

```

​	在 AIX 上，当链接使用 -buildmode=c-archive 构建的 Go 归档文件的 C 程序时，必须向 C 编译器传递 -Wl,-bnoobjreorder。

#### 调用 Go 和 C 之间的交互

​	调用 Go 和 C/C++ 代码有两种不同的方法。

​	第一种是 cgo 工具，它是 Go 发行版的一部分。有关如何使用它的信息，请参阅 cgo 文档(go doc cmd/cgo)。

​	第二种是 SWIG 程序，它是一种用于不同语言间接口的通用工具。关于 SWIG 的信息请参见 http://swig.org/。在运行 go build 时，任何具有 .swig 扩展名的文件都将被传递给 SWIG。任何具有 .swigcxx 扩展名的文件都将被传递给带有 -c++ 选项的 SWIG。

​	当使用 cgo 或 SWIG 时，go build 将任何 .c、.m、.s、.S 或 .sx 文件传递给 C 编译器，并将任何 .cc、.cpp、.cxx 文件传递给 C++ 编译器。CC 或 CXX 环境变量可以设置以确定要使用的 C 或 C++ 编译器。

#### 构建和测试缓存

​	go 命令缓存构建输出以便将来重复使用。缓存数据的默认位置是当前操作系统的标准用户缓存目录中名为 go-build 的子目录。设置 GOCACHE 环境变量将覆盖此默认设置，并且运行 'go env GOCACHE' 将打印当前缓存目录。

​	go 命令定期删除未经常用的缓存数据。运行 'go clean -cache' 将删除所有缓存的数据。

构建缓存正确计算 Go 源文件、编译器、编译器选项等更改：在典型用法中不需要显式清除缓存。但是，构建缓存不会检测使用 cgo 导入的 C 库的更改。如果您对系统上的 C 库进行了更改，则需要显式清除缓存，否则使用 -a 构建标志(请参阅 'go help build')强制重建依赖于更新的 C 库的包。

​	go 命令还会缓存成功的软件包测试结果。有关详细信息，请参见 'go help test'。运行 'go clean -testcache' 将删除所有缓存的测试结果(但不会删除缓存的构建结果)。

​	go 命令还会缓存用于 'go test -fuzz' 进行模糊测试的值，特别是将扩展代码覆盖率的值传递给模糊函数。这些值不用于常规构建和测试，但是它们存储在构建缓存的子目录中。运行 'go clean -fuzzcache' 将删除所有缓存的模糊值。这可能会使模糊测试在短时间内变得不那么有效。

​	GODEBUG环境变量可以启用有关缓存状态的调试信息：

​	GODEBUG=gocacheverify=1会导致go命令绕过使用任何缓存条目，而是重新构建所有内容并检查结果是否与现有的缓存条目匹配。

​	GODEBUG=gocachehash=1会导致go命令打印用于构建缓存查找键的所有内容哈希的输入。输出很多，但可用于调试缓存。

​	GODEBUG=gocachetest=1会导致go命令打印有关是否重用缓存测试结果的详细信息。

#### 环境变量 

​	go命令及其调用的工具会查询环境变量以进行配置。如果环境变量未设置，则go命令将使用合理的默认设置。要查看变量`<NAME>`的有效设置，请运行'go env `<NAME>`'。要更改默认设置，请运行'`go env -w <NAME>=<VALUE>`'。使用'go env -w'更改的默认值会记录在一个Go环境配置文件中，该文件存储在每个用户配置目录中，该目录由os.UserConfigDir报告。配置文件的位置可以通过设置环境变量GOENV来更改，'go env GOENV'打印有效位置，但'go env -w'无法更改默认位置。有关详细信息，请参见'go help env'。

​	通用环境变量：

```
GO111MODULE
	控制go命令运行在模块感知模式还是GOPATH模式下。
	可能的取值为"off"，"on"或"auto"。
	详见https://golang.org/ref/mod#mod-commands。
	
GCCGO
	运行'go build -compiler=gccgo'所用的gccgo命令。
	
GOARCH
	编译代码的架构或处理器。例如，amd64，386，arm，ppc64。
	
GOBIN
	'go install'命令安装命令的目录。
	
GOCACHE
	go命令存储缓存信息以便于在未来的构建中重复使用的目录。
	
GOMODCACHE
	go命令存储已下载模块的目录。
	
GODEBUG
	启用各种调试功能。详见'go doc runtime'。
	
GOENV
	Go环境配置文件的位置。
	不能使用'go env -w'设置。
	将环境变量GOENV设置为"off"禁用使用默认配置文件。
	
GOFLAGS
	要默认应用的一系列空格分隔的-flag=value标志设置，
	当当前命令已知该标志时。
	每个条目必须是一个独立的标志。
	因为条目是以空格分隔的，所以标志值不能包含空格。
	在命令行中列出的标志将在此列表之后应用，因此会覆盖它。
	
GOINSECURE
	用逗号分隔的模块路径前缀的通配符模式（使用Go的path.Match语法），
	它们始终以不安全的方式获取。
	仅适用于直接获取的依赖项。
	GOINSECURE不会禁用校验和数据库验证。
	GOPRIVATE或GONOSUMDB可用于实现该功能。

GOOS
	编译代码的操作系统。例如，linux，darwin，windows，netbsd。

GOPATH
	更多详见'go help gopath'。

GOPROXY
	Go 模块代理的 URL。
	有关详细信息，请参见 
	https://golang.org/ref/mod#environment-variables 
	和 https://golang.org/ref/mod#module-proxy。
	
GOPRIVATE,GONOPROXY,GONOSUMDB
	逗号分隔的模块路径前缀的通配符列表（使用 Go 的 path.Match 语法），
	它们应始终直接获取，或不应与校验和数据库进行比较。
	请参见 https://golang.org/ref/mod#private-modules。
	
GOROOT
	Go 树的根目录。
	
GOSUMDB
	要使用的校验和数据库的名称及其公钥和 URL（可选）。
	请参见 https://golang.org/ref/mod#authenticating。


GOTMPDIR
	Go 命令将写入临时源文件、包和二进制文件的目录。

GOVCS
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

​	用于与 cgo 一起使用的环境变量：

```
AR
	在使用 gccgo 编译器进行构建时，用于操作库档案的命令。默认值为 'ar'。

CC
	用于编译 C 代码的命令。

CGO_ENABLED
	是否支持 cgo 命令。为 0 或 1。
	
CGO_CFLAGS
	cgo 在编译 C 代码时将传递给编译器的标志。
	
CGO_CFLAGS_ALLOW
	指定允许在 #cgo CFLAGS 源代码指令中出现的其他标志的正则表达式。
	不适用于 CGO_CFLAGS 环境变量。
	
CGO_CFLAGS_DISALLOW
	指定必须禁止出现在 #cgo CFLAGS 源代码指令中的标志的正则表达式。
	不适用于 CGO_CFLAGS 环境变量。

CGO_CPPFLAGS, CGO_CPPFLAGS_ALLOW, CGO_CPPFLAGS_DISALLOW
	与 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW 类似，
	但用于 C 预处理器。
	
CGO_CXXFLAGS, CGO_CXXFLAGS_ALLOW, CGO_CXXFLAGS_DISALLOW
	Like CGO_CFLAGS, CGO_CFLAGS_ALLOW, and CGO_CFLAGS_DISALLOW,
	but for the C++ compiler.
	与 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW 类似，
	但是用于 C++ 编译器。
	
CGO_FFLAGS, CGO_FFLAGS_ALLOW, CGO_FFLAGS_DISALLOW
	与 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW 类似，
	但是用于 Fortran 编译器。
	
CGO_LDFLAGS, CGO_LDFLAGS_ALLOW, CGO_LDFLAGS_DISALLOW
	与 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW 类似，
	但是用于链接器。
	
CXX
	用于编译 C++ 代码的命令。
	
FC
	用于编译 Fortran 代码的命令。
	
PKG_CONFIG
	pkg-config 工具的路径。	
```

​	特定架构的环境变量：

```
GOARM
	用于 GOARCH=arm，表示编译的 ARM 架构。
	有效的值为 5、6、7。
	
GO386
	用于 GOARCH=386，指定如何实现浮点指令。
	有效的值为 sse2（默认）和 softfloat。
	
GOAMD64
	用于 GOARCH=amd64，表示编译的微架构级别。
	有效的值为 v1（默认）、v2、v3、v4。
	参见 https://golang.org/wiki/MinimumRequirements#amd64。
	
GOMIPS
	用于 GOARCH=mips{,le}，表示是否使用浮点指令。
	有效的值为 hardfloat（默认）和 softfloat。
	
GOMIPS64
	用于 GOARCH=mips64{,le}，表示是否使用浮点指令。
	有效的值为 hardfloat（默认）和 softfloat。
	
GOPPC64
	用于 GOARCH=ppc64{,le}，表示目标 ISA（指令集架构）。
	有效的值为 power8（默认）、power9 和 power10。
	
GOWASM
	用于 GOARCH=wasm，以逗号分隔的实验性 WebAssembly 功能列表。
	有效的值为 satconv 和 signext。
```

​	用于代码覆盖的环境变量：

```
GOCOVERDIR
	"go build -cover" 二进制文件生成的代码覆盖数据文件将写入其中的目录。
	需要启用 GOEXPERIMENT=coverageredesign。
```

​	特殊用途的环境变量：

```
GCCGOTOOLDIR
	如果设置了，gccgo 工具的所在位置，例如 cgo。
	默认值基于 gccgo 的配置。
	
GOEXPERIMENT
	启用或禁用的工具链实验的逗号分隔列表。
	可用实验的列表可能随时更改。
	有关目前有效值，请参见 src/internal/goexperiment/flags.go。
	警告：此变量仅为开发和测试 Go 工具链本身提供。
	超出此目的的使用不受支持。
	
GOROOT_FINAL
	安装 Go 树的根目录，当它安装在与构建位置不同的位置时。
	堆栈跟踪中的文件名将从 GOROOT 重写为 GOROOT_FINAL。
	
GO_EXTLINK_ENABLED
	使用 -linkmode=auto 与使用 cgo 的代码时，
	链接器是否应使用外部链接模式。
	将其设置为 0 以禁用外部链接模式，设置为 1 以启用。
	
GIT_ALLOW_PROTOCOL
	Git 定义的允许使用 git fetch/clone 的方案的冒号分隔列表。
	如果设置，未明确提及的任何方案都将被 'go get' 视为不安全。
	因为该变量由 Git 定义，所以不能使用 'go env -w' 设置默认值。
```

​	从 'go env' 中获取的其他信息，但不是从环境中读取的：

```
GOEXE
	可执行文件名后缀（Windows 上为".exe"，其他系统上为空字符串）。

GOGCCFLAGS
	CC 命令提供的参数的以空格分隔的列表。

GOHOSTARCH
	Go 工具链二进制文件的架构（GOARCH）。

GOHOSTOS
	Go 工具链二进制文件的操作系统（GOOS）。

GOMOD
	主模块的 go.mod 的绝对路径。
	如果启用了模块感知模式，
	但没有 go.mod，则 GOMOD 将是 os.DevNull
	（在类 Unix 系统上为"/dev/null"，在 Windows 上为"NUL"）。
	如果禁用了模块感知模式，则 GOMOD 将为空字符串。
	
GOTOOLDIR
	Go 工具（compile、cover、doc 等）所在的目录。
	
GOVERSION
	安装的 Go 树的版本，由 runtime.Version 报告。
```

#### 文件类型

​	go 命令检查每个目录中一组受限制的文件的内容。 它根据文件名的扩展名来确定要检查哪些文件。 这些扩展名是：

```
.go
	Go源文件。
.c, .h
	C源文件。
	如果包使用cgo或SWIG，
	则这些文件将使用本机编译器（通常是gcc）编译；
	否则它们会导致错误。
	
.cc, .cpp, .cxx, .hh, .hpp, .hxx
	C++源文件。仅在使用cgo或SWIG时有用，并且始终使用本机编译器编译。
	
.m
	Objective-C源文件。仅在使用cgo时有用，并始终使用本机编译器编译。
	
.s, .S, .sx
	汇编器源文件。
	如果包使用cgo或SWIG，则这些文件将使用本机汇编器（通常是gcc）汇编；
	否则它们将使用Go汇编器汇编。
	
.swig, .swigcxx
	SWIG定义文件。
	
.syso
	系统对象文件。	
```

​	除了.syso之外，每个这些类型的文件都可以包含构建约束条件，但是go命令会在第一个不是空白行或`//-`样式行注释的条目处停止扫描构建约束条件。有关更多详细信息，请参阅go/build包文档。

#### go.mod文件

​	模块版本由源文件树定义，在其根目录中有一个go.mod文件。运行go命令时，它会在当前目录中查找，然后在父目录中查找，以找到标记当前模块根的go.mod文件。

​	go.mod文件格式在https://golang.org/ref/mod#go-mod-file中详细描述。

​	要创建新的go.mod文件，请使用'go mod init'。有关详细信息，请参见'go help mod init'或https://golang.org/ref/mod#go-mod-init。

​	要添加缺少的模块要求或删除不需要的要求，请使用'go mod tidy'。有关详细信息，请参见'go help mod tidy'或https://golang.org/ref/mod#go-mod-tidy。

​	要添加、升级、降级或删除特定模块要求，请使用'go get'。有关详细信息，请参见'go help module-get'或https://golang.org/ref/mod#go-get。

​	要进行其他更改或将go.mod解析为JSON供其他工具使用，请使用'go mod edit'。请参阅'go help mod edit'或https://golang.org/ref/mod#go-mod-edit。

#### GOPATH环境变量

​	Go路径用于解析导入语句。它由go/build包实现并记录。

​	GOPATH环境变量列出了查找Go代码的位置。在Unix上，该值是以冒号分隔的字符串。在Windows上，该值是以分号分隔的字符串。在Plan 9上，该值是一个列表。

​	如果环境变量未设置，GOPATH默认为用户主目录中名为"go"的子目录（在Unix上为`$HOME/go`，在Windows上为`%USERPROFILE%\go`），除非该目录包含Go发行版。运行"go env GOPATH"可以查看当前GOPATH。

​	请参见https://golang.org/wiki/SettingGOPATH来设置自定义GOPATH。

​	GOPATH中列出的每个目录都必须具有预定的结构：

- src目录包含源代码。src下面的路径确定导入路径或可执行文件名。

- pkg目录保存已安装的包对象。与Go树中一样，每个目标操作系统和架构对都有自己的pkg子目录（pkg/GOOS_GOARCH）。

​			如果DIR是在GOPATH中列出的目录，则具有DIR/src/foo/bar中的源代码的包可以作为"foo/bar"导入，并将其编译形式安装到"DIR/pkg/GOOS_GOARCH/foo/bar.a"中。

- bin目录保存已编译的命令。每个命令以其源目录命名，但只使用最后一个元素，而不是整个路径。也就是说，具有DIR/src/foo/quux源代码的命令将安装到DIR/bin/quux，而不是DIR/bin/foo/quux。前缀"foo/"被删除，以便您可以将DIR/bin添加到PATH中以获得已安装命令。如果设置了GOBIN环境变量，则命令将安装到其命名的目录，而不是DIR/bin。GOBIN必须是绝对路径。

这是一个示例目录布局：

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

​	Go会在GOPATH列表中的每个目录中查找源代码，但新的包总是被下载到列表中第一个目录。

​	参见 https://golang.org/doc/code.html 获取示例。

#### GOPATH 和 Modules

​	使用 modules 时，GOPATH 不再用于解析导入。但是，它仍然用于存储已下载的源代码（在 GOPATH/pkg/mod 中）和编译的命令（在 GOPATH/bin 中）。

#### 内部目录

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

​	在 z.go 中导入为"foo/internal/baz"，但是该导入语句只能出现在以 foo 为根的子树中的源文件中。foo/f.go、foo/bar/x.go 和 foo/quux/y.go 中的源文件都可以导入"foo/internal/baz"，但是 crash/bang/b.go 中的源文件不能。

​	有关详细信息，请参见 https://golang.org/s/go14internal。

#### vendor目录

​	Go 1.6 包括使用本地副本来满足这些依赖项的导入的支持，通常称为供应商。

​	位于名为"vendor"的目录以下的代码只能由根目录为"vendor"的目录树中的代码导入，并且仅使用省略前缀直到包含 vendor 元素的路径。

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

​	同样的可见性规则适用于 internal，但是 z.go 中的代码导入时使用 "baz"，而不是 "foo/vendor/baz"。

​	位于更深层次的 vendor 目录中的代码会遮蔽较高层次的目录中的代码。在以 foo 为根的子树中，对 "crash/bang" 的导入会解析为 "foo/vendor/crash/bang"，而不是顶层的 "crash/bang"。

​	位于 vendor 目录中的代码不受导入路径检查的限制（参见 'go help importpath'）。

​	当 'go get' 检出或更新 git 子模块时，它现在也会更新子模块。

​	vendor 目录不影响由 'go get' 第一次检出的新代码库的位置：它们总是被放置在主 GOPATH 中，而不是在 vendor 子目录中。

​	详见 https://golang.org/s/go15vendor。

#### 遗留 GOPATH go get

​	'go get' 命令的行为取决于 go 命令是在模块感知模式还是遗留 GOPATH 模式下运行。即使在模块感知模式下，也可以通过 'go help gopath-get' 访问此帮助文本，它描述了 'go get' 在遗留 GOPATH 模式下的操作。

用法：

```
 go get [-d] [-f] [-t] [-u] [-v] [-fix] [build flags] [packages]
```

​	get 下载指定导入路径的包及其依赖项。它然后安装指定的包，就像 'go install' 一样。

​	`-d` 标志指示 get 在下载包后停止，即指示 get 不安装包。

​	`-f` 标志只在设置了 -u 标志时有效，强制 get -u 不验证每个包是否已从其导入路径所暗示的源代码库检出。如果源代码是原始代码的本地分支，则此功能可能很有用。

​	`-fix` 标志指示 get 在解析依赖项或构建代码之前对下载的包运行 fix 工具。

​	`-t` 标志指示 get 也下载构建指定包所需的测试包。

​	`-u` 标志指示 get 使用网络更新指定的包及其依赖项。默认情况下，get 使用网络检出缺少的包，但不使用网络查找现有包的更新。

​	`-v` 标志启用详细进度和调试输出。

​	get 命令还接受构建标志以控制安装。请参阅 'go help build'。

​	在检出新包时，get 会创建目标目录 GOPATH/src/<import-path>。如果 GOPATH 包含多个条目，则 get 会使用第一个条目。有关更多详细信息，请参阅 'go help gopath'。

​	在检出或更新包时，get 会查找与本地安装的 Go 版本匹配的分支或标记。最重要的规则是，如果本地安装正在运行版本 "go1"，则 get 会搜索名为 "go1" 的分支或标记。如果不存在此类版本，则检索包的默认分支。

​	当 go get 检出或更新 Git 存储库时，它也会更新由存储库引用的任何 git 子模块。

​	get 永远不会检出或更新存储在 vendor 目录中的代码。

​	有关指定包的更多信息，请参阅 'go help packages'。

​	有关 'go get' 如何查找要下载的源代码的更多信息，请参阅 'go help importpath'。

​	本文描述了使用 GOPATH 管理源代码和依赖项时的 get 行为。如果 go 命令是在模块感知模式下运行，则 get 的标志和效果的详细信息会发生更改，'go help get' 也会发生更改。请参阅 'go help modules' 和 'go help module-get'。

​	另请参阅：go build、go install、go clean。

#### 模块代理协议

​	Go 模块代理是任何可以响应特定格式 URL 的 GET 请求的 Web 服务器。请求没有查询参数，因此即使是从固定文件系统（包括 `file://URL`）服务的站点也可以是模块代理。

​	有关 GOPROXY 协议的详细信息，请参阅 https://golang.org/ref/mod#goproxy-protocol。

#### 导入路径语法

​	导入路径（请参阅 'go help packages'）表示存储在本地文件系统中的包。一般来说，导入路径表示标准包（如 "unicode/utf8"）或在工作区之一找到的包（有关详细信息，请参阅：'go help gopath'）。

#### 相对导入路径

​	以 `./` 或 `../` 开头的导入路径称为相对路径。工具链支持相对导入路径作为两种快捷方式。

​	首先，相对路径可以作为命令行上的快捷方式。如果您在包含代码作为 "unicode" 导入的目录中工作，并想要运行 "unicode/utf8" 的测试，则可以键入 "go test ./utf8" 而不需要指定完整路径。同样，在反向情况下，"go test .."将从 "unicode/utf8" 目录测试 "unicode"。相对模式也被允许，例如 "go test ./..." 可以测试所有子目录。有关模式语法的详细信息，请参见 'go help packages'。

​	其次，如果您在不在工作空间中编译 Go 程序，那么您可以在该程序的导入语句中使用相对路径，以引用附近的代码，也不在工作空间中。这使得在通常的工作空间之外实验小型多包程序非常容易，但这些程序无法使用 "go install" 进行安装（没有工作空间来安装它们），因此每次构建时都会重新构建。为避免歧义，Go 程序不能在工作空间内使用相对导入路径。

#### 远程导入路径 

​	某些导入路径也描述了如何使用版本控制系统获取软件包的源代码。

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

​	对于在其他服务器上托管的代码，导入路径可以使用版本控制类型进行限定，或者 go 工具可以通过 https/http 动态获取导入路径，并从 HTML 的 <meta> 标记中发现代码所在的位置。

​	要声明代码位置，形式为

```
repository.vcs/path
```

的导入路径指定了给定的仓库，带或不带 `.vcs` 后缀，使用命名的版本控制系统，然后是该仓库内的路径。支持的版本控制系统是：

```
Bazaar      .bzr
Fossil      .fossil
Git         .git
Mercurial   .hg
Subversion  .svn
```

例如，

```
import "example.org/user/foo.hg"
```

表示在 example.org/user/foo 或 foo.hg 上的 Mercurial 存储库的根目录，而

```
import "example.org/repo.git/foo/bar"
```

表示 example.org/repo 或 repo.git 上的 Git 存储库的 foo/bar 目录。

​	当版本控制系统支持多个协议时，每个协议都会依次尝试下载。例如，Git 下载尝试 https://，然后是 git+ssh://。

​	默认情况下，下载仅限于已知的安全协议（例如 https、ssh）。要覆盖此设置以进行 Git 下载，可以设置 GIT_ALLOW_PROTOCOL 环境变量（有关更多详细信息，请参见 'go help environment'）。

​	如果导入路径不是已知的代码托管站点，也缺乏版本控制限定符，则 go 工具尝试通过 https/http 获取导入，并在文档的 HTML `<head>` 中查找 `<meta>` 标签。

meta 标签的形式为：

```
<meta name="go-import" content="import-prefix vcs repo-root">
```

import-prefix 是与存储库根目录对应的导入路径。它必须是被 "go get" 获取的包的前缀或精确匹配。如果它不是精确匹配，则会在该前缀处进行另一个 http 请求以验证 <meta> 标签的匹配。

​	meta 标签应尽可能早地出现在文件中。特别是，它应该出现在任何原始 JavaScript 或 CSS 之前，以避免混淆 go 命令的受限解析器。

​	vcs 是 "bzr"、"fossil"、"git"、"hg"、"svn" 之一。

​	repo-root 是包含方案但不包含 .vcs 限定符的版本控制系统的根。

​	例如，

```
import "example.org/pkg/foo"
```

将导致以下请求：

```
https://example.org/pkg/foo?go-get=1（首选）
http://example.org/pkg/foo?go-get=1（回退，仅在正确设置了 GOINSECURE 的情况下）
```

如果该页面包含以下 meta 标签

```
<meta name="go-import" content="example.org git https://code.org/r/p/exproj">
```

go 工具将验证 https://example.org/?go-get=1 是否包含相同的 meta 标签，然后将 git 克隆 https://code.org/r/p/exproj 到 GOPATH/src/example.org。

​	在使用 GOPATH 时，下载的包会写入 GOPATH 环境变量中列出的第一个目录。（请参见 'go help gopath-get' 和 'go help gopath'。）

​	在使用模块时，下载的包存储在模块缓存中。请参阅 [Go模块参考中的模块缓存命令](../../GoModulesReference/ModuleCache)。

​	使用模块时，还有一种go-import meta标签的变体会被识别，并且优先于列出版本控制系统的标签。这个变体在content值中使用"mod"作为vcs，例如：

```
<meta name="go-import" content="example.org mod https://code.org/moduleproxy">
```

这个标签意味着要从可用于URL https://code.org/moduleproxy 的模块代理中获取以example.org开头的模块。有关代理协议的详细信息，请参见[Go模块参考中的模块代理](../../GoModulesReference/ModuleProxies#goproxy-protocol)。

#### 导入路径检查 

​	当上述自定义导入路径功能重定向到已知的代码托管站点时，每个结果包都有两个可能的导入路径，一个是使用自定义域名，另一个是使用已知的托管站点。

​	如果一个包语句紧接着（在下一个换行符之前）有一个这两种形式之一的注释，那么它被称为"导入注释"：

``` go
package math // import "path"
package math /* import "path" */
```

​	go命令将拒绝安装带有导入注释的包，除非它被引用了该导入路径。通过这种方式，导入注释让包的作者确保使用自定义导入路径，而不是底层的代码托管站点的直接路径。

​	在vendor树中发现的代码不会进行导入路径检查。这使得可以将代码复制到vendor树中的其他位置，而不需要更新导入注释。

​	当使用模块时，导入路径检查也被禁用了。导入路径注释被go.mod文件中的模块声明所取代。

​	有关详细信息，请参见https://golang.org/s/go14customimport。

#### 模块、模块版本等

​	模块是Go管理依赖项的方式。

​	一个模块是一组一起发布、版本化和分发的包。模块可以直接从版本控制存储库或模块代理服务器下载。

​	有关模块教程系列，请参见https://golang.org/doc/tutorial/create-module。

​	有关模块的详细参考，请参见[Go模块参考](../../GoModulesReference/Introduction)。

​	默认情况下，go命令可以从https://proxy.golang.org下载模块。它可以使用https://sum.golang.org上的checksum数据库对模块进行身份验证。这两个服务都由Google的Go团队运营。这些服务的隐私政策可在https://proxy.golang.org/privacy和https://sum.golang.org/privacy上获得。

​	可以使用GOPROXY、GOSUMDB、GOPRIVATE和其他环境变量来配置go命令的下载行为。有关更多信息，请参见"go help environment"和[Go模块参考中的私有模块中的隐私](../../GoModulesReference/PrivateModules#privacy)。

#### 使用 go.sum 进行模块认证 

​	当 go 命令将模块 zip 文件或 go.mod 文件下载到模块缓存中时，它会计算一个密码哈希值并将其与已知值进行比较，以验证该文件自从第一次下载以来没有发生更改。已知哈希值存储在模块根目录中名为 go.sum 的文件中。哈希值也可以根据 GOSUMDB、GOPRIVATE 和 GONOSUMDB 的值从校验和数据库中下载。

For details, see [Go模块参考中的验证模块](../../GoModulesReference/AuthenticatingModules).

#### 软件包列表和模式 

​	许多命令都适用于一组软件包：

```
go action [packages]
```

通常，[packages] 是一组导入路径。

​	一个根路径或以 `.` 或 `..` 元素开头的导入路径被解释为文件系统路径，并表示该目录中的软件包。

​	否则，导入路径 P 表示在 GOPATH 环境变量列出的某个 DIR/src/P 目录中找到的软件包（有关详细信息，请参阅 'go help gopath'）。

​	如果没有给出导入路径，则该操作适用于当前目录中的软件包。

​	有四个保留名称用于不应与 go 工具一起构建的软件包：

- "main"表示独立可执行文件中的顶级包。
- "all"展开为在所有 GOPATH 树中找到的所有软件包。例如，'go list all' 列出了本地系统上的所有软件包。当使用模块时，"all"会展开为主模块中的所有软件包及其依赖项，包括任何这些依赖项的测试所需的依赖项。
- "std"类似于"all"，但仅展开到标准 Go 库中的软件包。
- "cmd"展开为 Go 存储库的命令及其内部库。

​	以"cmd/"开头的导入路径仅匹配 Go 存储库中的源代码。

​	如果导入路径包含一个或多个"`...`"通配符，则导入路径是一个模式，每个通配符可以匹配任何字符串，包括空字符串和包含斜杠的字符串。这种模式将展开为 GOPATH 树中找到的所有软件包目录，其名称与模式匹配。

​	为了方便常见模式的匹配，有两个特殊情况。第一种情况是在模式末尾添加 `/...` 可以匹配空字符串，例如 `net/...` 同时匹配 net 和其子目录下的包，如 net/http。第二种情况是任何包含通配符的斜杠分隔模式元素都不会匹配 vendored 包路径中的 "vendor" 元素，这样 `./...` 就不会匹配 `./vendor` 或 `./mycode/vendor` 子目录下的包，但是 `./vendor/...` 和 `./mycode/vendor/...` 会。然而，一个包含代码的名为 vendor 的目录不是一个 vendored 包：cmd/vendor 将是一个名为 vendor 的命令，模式 `cmd/...` 会匹配它。有关 vendoring 的更多信息，请参见 golang.org/s/go15vendor。

​	导入路径也可以命名从远程仓库下载的包。有关详情，请运行"go help importpath"。

​	程序中的每个包都必须具有唯一的导入路径。按照惯例，这是通过将每个路径以属于您自己的唯一前缀开头来安排的。例如，在 Google 内部使用的路径都以 'google' 开头，表示远程仓库的路径以代码的路径为开头，例如 'github.com/user/repo'。

​	程序中的包名不需要唯一，但有两个保留的包名具有特殊含义。名称 main 表示命令，而不是库。命令被构建成二进制文件，不能被导入。名称 documentation 表示目录中非 Go 程序的文档。包文档中的文件将被 go 命令忽略。

​	作为特殊情况，如果包列表是来自单个目录的 .go 文件列表，则该命令将应用于一个合成包，该包由恰好这些文件组成，忽略这些文件中的任何构建约束，并忽略目录中的任何其他文件。

​	以 "." 或 "_" 开头的目录和文件名将被 go 工具忽略，同样以 "testdata" 命名的目录也会被忽略。

#### 用于下载非公共代码的配置

​	go 命令默认从公共 Go 模块镜像 proxy.golang.org 下载模块。无论源代码如何，它还默认针对公共 Go 校验和数据库 sum.golang.org 验证下载的模块。这些默认值适用于公开的源代码。

GOPRIVATE 环境变量控制了哪些模块被 go 命令视为私有模块（不公开可用），因此不应使用代理或校验和数据库。该变量是一个以逗号分隔的模块路径前缀的通配符模式列表（符合 Go 的 path.Match 语法）。例如，

```
GOPRIVATE=*.corp.example.com,rsc.io/private
```

会使 go 命令将任何具有与其中任意一种模式匹配的路径前缀的模块视为私有模块，包括 git.corp.example.com/xyzzy、rsc.io/private 和 rsc.io/private/quux 等。

​	为了对模块下载和校验有更精细的控制，GONOPROXY 和 GONOSUMDB 环境变量接受相同类型的通配符模式列表，并分别覆盖 GOPRIVATE，用于确定是否使用代理和校验和数据库。

​	例如，如果公司运行一个用于服务私有模块的模块代理，则用户可以使用以下方式配置 go：

```
GOPRIVATE=*.corp.example.com
GOPROXY=proxy.example.com
GONOPROXY=none
```

GOPRIVATE 变量还用于定义 GOVCS 变量的 "public" 和 "private" 模式；请参阅 'go help vcs'。对于这种用法，GOPRIVATE 即使在 GOPATH 模式下也适用。在这种情况下，它匹配导入路径而不是模块路径。

​	可以使用 'go env -w' 命令（参见 'go help env'）为将来的 go 命令调用设置这些变量。

​	有关更多详细信息，请参阅[Go模块参考中的私有模块](../../GoModulesReference/PrivateModules)。

#### 测试标志 

​	'go test' 命令接受适用于 'go test' 自身和适用于生成的测试二进制文件的标志。

​	其中一些标志控制分析，并编写适用于 "go tool pprof" 的执行分析文件；运行 "go tool pprof -h" 以获取更多信息。pprof 的 `--alloc_space`、`--alloc_objects` 和 `--show_bytes` 选项控制信息的显示方式。

​	以下标志由 'go test' 命令识别并控制任何测试的执行：

```
-bench regexp
    仅运行与正则表达式匹配的基准测试。
    默认情况下，不运行基准测试。
	要运行所有基准测试，请使用"-bench。"或"-bench =。"。
	正则表达式由未加括号的斜杠（/）字符拆分为一系列正则表达式，
	并且基准标识符的每个部分必须与序列中的相应元素匹配（如果有）。
	匹配项的可能父项以b.N = 1运行，以确定子基准测试。
	例如，给定-bench = X / Y，将使用b.N = 1运行与X匹配的顶级基准测试，
	以查找任何与Y匹配的子基准测试，然后以完整形式运行它们。

-benchtime t
    运行足够的迭代来花费t，t以time.Duration形式指定
	（例如，-benchtime 1h30s）。默认为1秒（1s）。
	特殊语法Nx表示运行基准测试N次（例如，-benchtime 100x）。

-count n
    对每个测试、基准测试和模糊种子运行n次（默认为1次）。
	如果设置了-cpu，则对每个GOMAXPROCS值运行n次。
	示例总是运行一次。-count不适用于由-fuzz匹配的模糊测试。

-cover
    启用覆盖分析。
	请注意，由于覆盖工作是通过编译前对源代码进行注释来完成的，
	因此启用覆盖并编译或测试失败可能会报告与原始源不对应的行号。

-covermode set,count,atomic
    为正在测试的包设置覆盖分析模式。默认为"set"，
	除非启用了-race，此时为"atomic"。
	值：
	set：bool：是否运行此语句？
	count：int：此语句运行多少次？
	atomic：int：count，但在多线程测试中正确；显著更昂贵。
	设置-cover。

-coverpkg pattern1,pattern2,pattern3
    在每个测试中将覆盖分析应用于与模式匹配的包。
	默认情况下，每个测试仅分析正在测试的包。
	有关包模式的描述，请参见"go help packages"。
	设置-cover。

-cpu 1,2,4
    指定测试、基准测试或模糊测试应执行的GOMAXPROCS值列表。
	默认值为GOMAXPROCS的当前值。-cpu不适用于由-fuzz匹配的模糊测试。

-failfast
    第一个测试失败后不要启动新测试。

-fuzz regexp
    运行匹配正则表达式的模糊测试。
    指定时，命令行参数必须恰好匹配主模块中的一个包，
    并且 regexp 必须恰好匹配该包中的一个模糊测试。
    模糊测试将在测试、基准测试、其他模糊测试的种子库和示例完成后进行。
    有关详细信息，请参见测试包文档中的"模糊测试"部分。

-fuzztime t
   运行足够的迭代次数来执行指定的时间 t 的模糊测试目标
   （例如，-fuzztime 1h30s）。
   默认情况下，将永久运行。
   特殊语法 Nx 表示运行模糊测试目标 N 次（例如，-fuzztime 1000x）。

-fuzzminimizetime t
    在每个最小化尝试期间运行足够的迭代次数以执行指定的时间 t
    （例如，-fuzzminimizetime 30s）。
    默认情况下为 60 秒。
    特殊语法 Nx 表示运行模糊测试目标 N 次
    （例如，-fuzzminimizetime 100x）。

-json
    以 JSON 形式记录详细输出和测试结果。
    这以机器可读的格式呈现与 -v 标志相同的信息。

-list regexp
    列出与正则表达式匹配的测试、基准测试、模糊测试或示例。
    不会运行任何测试、基准测试、模糊测试或示例。
    仅列出顶级测试。不会显示子测试或子基准测试。

-parallel n
    允许并行执行调用 t.Parallel 的测试函数，
    以及运行种子库时调用 t.Parallel 的模糊测试。
    此标志的值是同时运行的最大测试数。
    在进行模糊测试时，此标志的值是可以调用模糊函数的子进程的最大数量，
    无论是否调用了 T.Parallel。
    默认情况下，-parallel 设置为 GOMAXPROCS 的值。
    将 -parallel 设置为比 GOMAXPROCS 更高的值可能会导致因 CPU 争用
    而降低性能，特别是在进行模糊测试时。
    请注意，-parallel 仅适用于单个测试二进制文件。
    'go test' 命令也可以并行运行不同包的测试，
    根据 -p 标志的设置（请参阅 'go help build'）。

-run regexp
    仅运行与正则表达式匹配的测试、示例和模糊测试。
    对于测试来说，
    正则表达式被不带方括号的斜杠字符(/)分隔成一系列正则表达式，
    每个测试标识符的部分必须与序列中的相应元素匹配（如果有的话）。
    请注意，可能的匹配项的父级也会运行，
    因此-run=X/Y会匹配并运行和报告与X匹配的所有测试，
    即使没有子测试与Y匹配，也必须运行它们来查找这些子测试。
    另请参见-skip。

-short
    告诉长时间运行的测试缩短运行时间。
    默认情况下关闭，但在all.bash期间设置，
    以便安装Go树可以运行一次健全性检查，但不必花费时间运行详尽的测试。

-shuffle off,on,N
    随机执行测试和基准的顺序。默认情况下关闭。
    如果将-shuffle设置为on，则它将使用系统时钟作为种子生成随机数。
    如果将-shuffle设置为整数N，则N将用作种子值。
    在这两种情况下，种子将报告以实现可重现性。

-skip regexp
    仅运行与正则表达式不匹配的测试、示例、模糊测试和基准。
    就像对于-run和-bench一样，对于测试和基准，
    正则表达式被不带方括号的斜杠(/)字符分隔成一系列正则表达式，
    每个测试标识符的部分必须与序列中的相应元素匹配（如果有的话）。

-timeout d
    如果测试二进制文件运行时间超过d，
    将会发生panic。如果d为0，则禁用超时。默认为10分钟（10m）。

-v
    输出详细信息：记录所有运行的测试。即使测试成功，
    也打印来自Log和Logf调用的所有文本。

-vet list
    配置"go test"期间"go vet"的调用，
    使用逗号分隔的vet检查列表。
    如果列表为空，
    "go test"将使用一个由认为值得处理的检查组成的精选列表运行"go vet"。
    如果列表为"off"，"go test"根本不运行"go vet"。
```

以下标记也可以被'go test'识别并用于在执行过程中对测试进行分析：

```
-benchmem
    打印基准测试的内存分配统计信息。

-blockprofile block.out
    当所有测试完成时，将 goroutine 阻塞分析文件写入指定文件。
	写入测试二进制文件与 '-c' 相同。

-blockprofilerate n
    通过调用 runtime.SetBlockProfileRate(n) 
    控制 goroutine 阻塞分析中的详细信息。
	请参阅 'go doc runtime.SetBlockProfileRate'。
	分析器的目标是平均每 n 纳秒程序阻塞时采样一个阻塞事件。
	默认情况下，如果在未使用此标志的情况下设置了-test.blockprofile，
	则记录所有阻塞事件，相当于-test.blockprofilerate = 1。

-coverprofile cover.out
    在所有测试通过后将覆盖率分析结果写入文件。
	设置 -cover。

-cpuprofile cpu.out
    在退出之前将 CPU 分析结果写入指定的文件中。
	写入测试二进制文件，就像-c一样。

-memprofile mem.out
    在所有测试通过后，将分配分析结果写入指定的文件中。
	写入测试二进制文件，就像-c一样。

-memprofilerate n
    通过设置 runtime.MemProfileRate 
    来启用更精确（也更昂贵）的内存分配分析。
	请参阅"go doc runtime.MemProfileRate"。
	要分析所有内存分配，请使用-test.memprofilerate = 1。

-mutexprofile mutex.out
    当所有测试都完成后，将一个互斥锁争用分析文件写入指定的文件中。
	写入测试二进制文件，就像-c一样。

-mutexprofilefraction n
    对持有有争议的互斥锁的 goroutine 进行1:n的堆栈跟踪采样。

-outputdir directory
    将性能分析的输出文件放置在指定的目录中，
    默认情况下是"go test"正在运行的目录。

-trace trace.out
    在退出之前将执行跟踪结果写入指定的文件中。
```

​	每个标志也可以使用可选的"test."前缀，例如 -test.v。但是，当直接调用生成的测试二进制文件（即"go test -c"的结果）时，前缀是必需的。

​	"go test"命令会在可选的包列表之前和之后重写或删除已识别的标志，然后调用测试二进制文件。

​	例如，命令

```
go test -v -myflag testdata -cpuprofile=prof.out -x
```

将编译测试二进制文件，然后将其作为以下方式运行

```
pkg.test -test.v -myflag testdata -test.cpuprofile=prof.out
```

（-x 标志被删除，因为它仅适用于 go 命令的执行，而不适用于测试本身。）

生成剖析的测试标志（除了覆盖率之外）也会将测试二进制文件留在 pkg.test 中，以便在分析剖析时使用。

​	当"go test"运行测试二进制文件时，它是从相应的包的源代码目录中运行的。根据测试的不同，调用生成的测试二进制文件时可能需要执行相同的操作。因为该目录可能位于模块缓存中，而模块缓存可能是只读的并且由校验和进行验证，所以测试不能将其或模块中的任何其他目录写入，除非用户明确请求（例如使用 -fuzz 标志，将失败写入 testdata/fuzz）。

​	如果存在命令行包列表，则必须在任何未知于 go test 命令的标志之前出现。继续上面的例子，包列表必须出现在 -myflag 之前，但可以出现在 -v 的任何一侧。

​	当"go test"以包列表模式运行时，"go test"会缓存成功的包测试结果，以避免不必要的重复运行测试。要禁用测试缓存，请使用任何不可缓存的标志或参数。显式禁用测试缓存的惯用方法是使用 -count=1。

​	要使测试二进制文件的参数不被解释为已知标志或包名，请使用 -args（请参见"go help test"），该标志会将命令行的剩余部分传递给测试二进制文件，未经解释和未被改变。

​	例如，命令

```
go test -v -args -x -v
```

将编译测试二进制文件，然后将其作为以下方式运行

```
pkg.test -test.v -x -v
```

同样，

```
go test -args math
```

将编译测试二进制文件，然后将其作为以下方式运行

```
pkg.test math
```

​	在第一个例子中，-x和第二个-v参数被原封不动地传递给测试二进制文件，不会对go命令本身产生影响。在第二个例子中，参数math被传递给测试二进制文件，而不是被解释为包列表。

#### 测试函数 

​	'go test' 命令期望在对应被测试包的 "*_test.go" 文件中找到测试、基准测试和示例函数。

​	测试函数以 TestXxx 命名（其中 Xxx 不以小写字母开头），应该有如下签名：

``` go
func TestXxx(t *testing.T) { ... }
```

​	基准测试函数以 BenchmarkXxx 命名，应该有如下签名：

``` go
func BenchmarkXxx(b *testing.B) { ... }
```

​	模糊测试以 FuzzXxx 命名，应该有如下签名：

``` go
func FuzzXxx(f *testing.F) { ... }
```

​	示例函数与测试函数类似，但是输出到 os.Stdout 而不是使用 *testing.T 来报告成功或失败。如果函数的最后一个注释以 "Output:" 开头，则输出与注释进行完全匹配（参见下面的示例）。如果最后一个注释以 "Unordered output:" 开头，则输出与注释进行匹配，但是忽略行的顺序。没有此类注释的示例会被编译但不会被执行。在 "Output:" 后没有文本的示例会被编译、执行，并期望不产生输出。

​	Godoc 显示 ExampleXxx 的正文，以演示函数、常量或变量 Xxx 的使用。具有接收者类型 T 或 *T 的方法 M 的示例命名为 ExampleT_M。对于给定函数、常量或变量，可以有多个示例，它们通过一个尾部 _xxx 区分开，其中 xxx 是一个不以大写字母开头的后缀。

​	下面是一个示例：

``` go
func ExamplePrintln() {
	Println("The output of\nthis example.")
	// Output: The output of
	// this example.
}
```

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

​	当测试文件包含一个单独的示例函数、至少一个其他函数、类型、变量或常量声明以及没有测试、基准测试或模糊测试时，整个测试文件将被作为示例呈现。

​	有关更多信息，请参阅 testing 包的文档。

#### 用 GOVCS 控制版本控制

​	'go get' 命令可以运行版本控制命令（如 git）来下载导入的代码。这个功能对于 Go 分散式包生态系统至关重要，其中的代码可以从任何服务器导入，但如果恶意服务器找到了一种方法来让调用的版本控制命令运行未预期的代码，这也是一个潜在的安全问题。

​	为了平衡功能和安全问题，'go get' 命令默认只使用 git 和 hg 从公共服务器下载代码。但它将使用任何已知的版本控制系统（bzr、fossil、git、hg、svn）从私有服务器下载代码，这些服务器定义为托管匹配 GOPRIVATE 变量（请参见 'go help private'）的包。允许 Git 和 Mercurial 的原理在于这两个系统在作为不受信任服务器的客户端运行时已经得到了最多的问题关注。相比之下，Bazaar、Fossil 和 Subversion 主要用于受信任的身份验证环境中，并且没有像攻击面一样受到充分的审查。

​	版本控制命令限制仅适用于使用直接版本控制访问下载代码时。当从代理服务器下载模块时，'go get' 使用代理协议，该协议始终被允许。默认情况下，'go get' 命令使用 Go 模块镜像（proxy.golang.org）获取公共包，仅在模块镜像拒绝为公共包提供服务（通常是因为法律原因）时，才回退到版本控制以获取私有包。因此，客户端仍然可以访问由 Bazaar、Fossil 或 Subversion 存储库提供的公共代码，默认情况下，因为这些下载使用 Go 模块镜像，该镜像承担运行版本控制命令的安全风险，使用自定义沙箱。

​	GOVCS 变量可用于更改特定包（由模块或导入路径标识）的允许版本控制系统。GOVCS 变量在模块感知模式和 GOPATH 模式下构建包时都适用。当使用模块时，模式与模块路径匹配。使用 GOPATH 时，模式与对应于版本控制存储库根的导入路径匹配。

​	GOVCS 的一般形式为 pattern:vcslist 规则，由逗号分隔的一组规则构成。pattern 是必须匹配模块或导入路径的一个或多个前导元素的 glob 模式。vcslist 是允许的版本控制命令的管道分隔列表，或者是 "all" 表示允许使用任何已知命令，或者是 "off" 表示禁止使用所有命令。请注意，如果一个模块匹配一个 vcslist 为 "off" 的模式，则仍可以下载，如果源服务器使用 "mod" 方案，则指示 go 命令使用 GOPROXY 协议下载模块。列表中最早匹配的模式优先应用，即使后面的模式也匹配。

​	例如，请考虑：

```
GOVCS=github.com:git,evil.com:off,*:git|hg
```

使用此设置，具有以 `github.com/` 开头的模块或导入路径的代码只能使用 git。evil.com 上的路径不能使用任何版本控制命令，而其他所有路径（`*` 匹配所有）只能使用 git 或 hg。

​	特殊模式 "public" 和 "private" 匹配公共和私有模块或导入路径。如果路径与 GOPRIVATE 变量匹配，则为私有路径；否则为公共路径。

​	如果 GOVCS 变量中没有规则与特定模块或导入路径匹配，则 'go get' 命令应用其默认规则，该规则现在可以用 GOVCS 表示为 'public:git|hg,private:all'。

​	要允许任何包使用任何版本控制系统，可以使用：

```
GOVCS=*:all
```

​	要禁用所有版本控制的使用，请使用：

```
GOVCS=*:off
```

​	"go env -w" 命令（请参阅 "go help env"）可用于设置 GOVCS 变量以供将来的 go 命令调用。
