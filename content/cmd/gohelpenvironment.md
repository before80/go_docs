+++
title = "go help environment "
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 530
description = ""
isCJKLanguage = true
draft = false

+++

​	

The go command and the tools it invokes consult environment variables for configuration. If an environment variable is unset or empty, the go command uses a sensible default setting. To see the effective setting of the variable `<NAME>`, run '`go env <NAME>`'. To change the default setting, run '`go env -w <NAME>=<VALUE>`'. Defaults changed using '`go env -w`' are recorded in a Go environment configuration file stored in the per-user configuration directory, as reported by os.UserConfigDir. The location of the configuration file can be changed by setting the environment variable GOENV, and '`go env GOENV`' prints the effective location, but '`go env -w`' cannot change the default location.See '`go help env`' for details.

​	go 命令及其调用的工具会查询环境变量进行配置。如果环境变量未设置或为空，则 go 命令将使用合理的默认设置。要查看变量 `<NAME>` 的有效设置，请运行 '`go env <NAME>`'。要更改默认设置，请运行 '`go env -w <NAME>=<VALUE>`'。使用 '`go env -w`' 更改的默认值将记录在每个用户配置目录中存储的 Go 环境配置文件中，如 os.UserConfigDir 所报告的。配置文件的位置可以通过设置环境变量 GOENV 进行更改，'`go env GOENV`' 打印有效位置，但 '`go env -w`' 不能更改默认位置。有关详细信息，请参见 '`go help env`'。

## 通用环境变量 General-purpose environment variables

### GO111MODULE

​        Controls whether the go command runs in module-aware mode or GOPATH mode.May be "off", "on", or "auto". See https://golang.org/ref/mod#mod-commands.

​	 控制 go 命令是在模块感知模式还是 GOPATH 模式下运行。可以是 "off"、"on" 或 "auto"。请参见 https://golang.org/ref/mod#mod-commands。

### GCCGO

​        The gccgo command to run for 'go build -compiler=gccgo'.

​	 用于 'go build -compiler=gccgo' 的 gccgo 命令。

### GOARCH

​        The architecture, or processor, for which to compile code. Examples are amd64, 386, arm, ppc64.

​	 用于编译代码的架构或处理器。例如 amd64、386、arm、ppc64。

### GOBIN

​        The directory where 'go install' will install a command.

​	 'go install' 将安装命令的目录。

### GOCACHE

​        The directory where the go command will store cached information for reuse in future builds.

​	 go 命令将存储用于将来构建中重用的缓存信息的目录。

### GOMODCACHE

​        The directory where the go command will store downloaded modules.

​	 go 命令将存储下载的模块的目录。

### GODEBUG

​        Enable various debugging facilities. See https://go.dev/doc/godebug for details.

​	 启用各种调试设施。请参见 https://go.dev/doc/godebug 了解详细信息。

### GOENV

​        The location of the Go environment configuration file. Cannot be set using 'go env -w'. Setting GOENV=off in the environment disables the use of the default configuration file.

​	 Go 环境配置文件的位置。无法使用 'go env -w' 设置。在环境中设置 GOENV=off 将禁用默认配置文件的使用。

### GOFLAGS

​        A space-separated list of -flag=value settings to apply to go commands by default, when the given flag is known by the current command. Each entry must be a standalone flag. Because the entries are space-separated, flag values must not contain spaces. Flags listed on the command line are applied after this list and therefore override it.

​	 默认情况下，当当前命令已知给定标志时，要应用于 go 命令的 -flag=value 设置的空格分隔列表。每个条目都必须是独立的标志。由于条目是空格分隔的，因此标志值不能包含空格。命令行上列出的标志将在此列表之后应用，因此会覆盖它。

### GOINSECURE

​        Comma-separated list of glob patterns (in the syntax of Go's path.Match) of module path prefixes that should always be fetched in an insecure manner. Only applies to dependencies that are being fetched directly. GOINSECURE does not disable checksum database validation. GOPRIVATE or GONOSUMDB may be used to achieve that.

​	以逗号分隔的模块路径前缀的 glob 模式列表（以 Go 的 path.Match 语法为准），始终以不安全的方式获取。仅适用于直接获取的依赖项。GOINSECURE 不会禁用校验和数据库验证。可以使用 GOPRIVATE 或 GONOSUMDB 来实现。

### GOOS

​        The operating system for which to compile code. Examples are linux, darwin, windows, netbsd.

​	用于编译代码的操作系统。例如 linux、darwin、windows、netbsd。

### GOPATH

​        Controls where various files are stored. See: 'go help gopath'.

​	 控制各种文件的存储位置。请参见：'go help gopath'。

### GOPROXY

​        URL of Go module proxy. See https://golang.org/ref/mod#environment-variables and https://golang.org/ref/mod#module-proxy for details.

​	Go 模块代理的 URL。请参见 https://golang.org/ref/mod#environment-variables 和 https://golang.org/ref/mod#module-proxy 了解详细信息。

### GOPRIVATE, GONOPROXY, GONOSUMDB

​        Comma-separated list of glob patterns (in the syntax of Go's path.Match)  of module path prefixes that should always be fetched directly or that should not be compared against the checksum database.See https://golang.org/ref/mod#private-modules.

​	 模块路径前缀的 glob 模式列表（以 Go 的 path.Match 语法为准），始终直接获取或不与校验和数据库进行比较。请参见 https://golang.org/ref/mod#private-modules。

### GOROOT

​        The root of the go tree.

​	 go 树的根目录。

### GOSUMDB

​        The name of checksum database to use and optionally its public key and URL. See https://golang.org/ref/mod#authenticating.

​	 用于使用其公钥和 URL 的校验和数据库的名称。请参见 https://golang.org/ref/mod#authenticating。

### GOTOOLCHAIN

​        Controls which Go toolchain is used. See https://go.dev/doc/toolchain.

​	 控制使用哪个 Go 工具链。请参见 https://go.dev/doc/toolchain。

### GOTMPDIR

​        The directory where the go command will write temporary source files, packages, and binaries.

​	 go 命令将写入临时源文件、包和二进制文件的目录。

### GOVCS

​        Lists version control commands that may be used with matching servers. See 'go help vcs'.

​	列出可用于与匹配服务器一起使用的版本控制命令。请参见 'go help vcs'。

### GOWORK

​        In module aware mode, use the given go.work file as a workspace file. By default or when GOWORK is "auto", the go command searches for a file named go.work in the current directory and then containing directories until one is found. If a valid go.work file is found, the modules specified will collectively be used as the main modules. If GOWORK is "off", or a go.work file is not found in "auto" mode, workspace mode is disabled.

​	 在模块感知模式下，将给定的 go.work 文件用作工作区文件。默认情况下或当 GOWORK 为 "auto" 时，go 命令会在当前目录以及包含的目录中搜索名为 go.work 的文件，直到找到一个。如果找到有效的 go.work 文件，则指定的模块将作为主模块一起使用。如果 GOWORK 为 "off"，或在 "auto" 模式下未找到 go.work 文件，则禁用工作区模式。



## 用于与 cgo 一起使用的环境变量 Environment variables for use with cgo

### AR

​        The command to use to manipulate library archives when building with the gccgo compiler. The default is 'ar'.

​	 用于在使用 gccgo 编译器构建时操作库归档的命令。默认值为 'ar'。

### CC

​        The command to use to compile C code.

​	用于编译 C 代码的命令。

### CGO_ENABLED

​        Whether the cgo command is supported. Either 0 or 1.

​	是否支持 cgo 命令。可以是 0 或 1。

### CGO_CFLAGS

​        Flags that cgo will pass to the compiler when compiling C code.

​	 cgo 在编译 C 代码时将传递给编译器的标志。

### CGO_CFLAGS_ALLOW

​        A regular expression specifying additional flags to allow to appear in #cgo CFLAGS source code directives. Does not apply to the CGO_CFLAGS environment variable.

​	指定允许在 #cgo CFLAGS 源代码指令中出现的额外标志的正则表达式。不适用于 CGO_CFLAGS 环境变量。

### CGO_CFLAGS_DISALLOW

​        A regular expression specifying flags that must be disallowed from appearing in #cgo CFLAGS source code directives. Does not apply to the CGO_CFLAGS environment variable.

​	指定必须禁止在 #cgo CFLAGS 源代码指令中出现的标志的正则表达式。不适用于 CGO_CFLAGS 环境变量。

### CGO_CPPFLAGS, CGO_CPPFLAGS_ALLOW, CGO_CPPFLAGS_DISALLOW

​        Like CGO_CFLAGS, CGO_CFLAGS_ALLOW, and CGO_CFLAGS_DISALLOW, but for the C preprocessor.

​	类似于 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW，但适用于 C 预处理器。

### CGO_CXXFLAGS, CGO_CXXFLAGS_ALLOW, CGO_CXXFLAGS_DISALLOW

​        Like CGO_CFLAGS, CGO_CFLAGS_ALLOW, and CGO_CFLAGS_DISALLOW, but for the C++ compiler.

​	类似于 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW，但适用于 C++ 编译器。

### CGO_FFLAGS, CGO_FFLAGS_ALLOW, CGO_FFLAGS_DISALLOW

​        Like CGO_CFLAGS, CGO_CFLAGS_ALLOW, and CGO_CFLAGS_DISALLOW, but for the Fortran compiler.

​	类似于 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW，但适用于 Fortran 编译器。

### CGO_LDFLAGS, CGO_LDFLAGS_ALLOW, CGO_LDFLAGS_DISALLOW

​        Like CGO_CFLAGS, CGO_CFLAGS_ALLOW, and CGO_CFLAGS_DISALLOW, but for the linker.

​	类似于 CGO_CFLAGS、CGO_CFLAGS_ALLOW 和 CGO_CFLAGS_DISALLOW，但适用于链接器。

### CXX

​        The command to use to compile C++ code.

​	用于编译 C++ 代码的命令。

### FC

​        The command to use to compile Fortran code.

​	用于编译 Fortran 代码的命令。

### PKG_CONFIG

​        Path to pkg-config tool.

​	 pkg-config 工具的路径。



## 架构特定的环境变量 Architecture-specific environment variables

### GOARM

​        For GOARCH=arm, the ARM architecture for which to compile. Valid values are 5, 6, 7.

​	用于 GOARCH=arm 时，要编译的 ARM 架构。有效值为 5、6、7。

### GO386

​        For GOARCH=386, how to implement floating point instructions. Valid values are sse2 (default), softfloat.

​	用于 GOARCH=386 时，如何实现浮点指令。有效值为 sse2（默认）、softfloat。

### GOAMD64

​        For GOARCH=amd64, the microarchitecture level for which to compile. Valid values are v1 (default), v2, v3, v4. See https://golang.org/wiki/MinimumRequirements#amd64

​	用于 GOARCH=amd64 时，要编译的微架构级别。有效值为 v1（默认）、v2、v3、v4。请参见 https://golang.org/wiki/MinimumRequirements#amd64

### GOMIPS

​        For GOARCH=mips{,le}, whether to use floating point instructions. Valid values are hardfloat (default), softfloat.

​	用于 GOARCH=mips{,le} 时，是否使用浮点指令。有效值为 hardfloat（默认）、softfloat。

### GOMIPS64

​        For GOARCH=mips64{,le}, whether to use floating point instructions. Valid values are hardfloat (default), softfloat.

​	 用于 GOARCH=mips64{,le} 时，是否使用浮点指令。有效值为 hardfloat（默认）、softfloat。

### GOPPC64

​        For GOARCH=ppc64{,le}, the target ISA (Instruction Set Architecture). Valid values are power8 (default), power9, power10.

​	用于 GOARCH=ppc64{,le} 时，目标 ISA（指令集架构）。有效值为 power8（默认）、power9、power10。

### GOWASM

​        For GOARCH=wasm, comma-separated list of experimental WebAssembly features to use. Valid values are satconv, signext.

​	用于 GOARCH=wasm 时，要使用的实验性 WebAssembly 功能的逗号分隔列表。有效值为 satconv、signext。

## 用于代码覆盖率的环境变量 Environment variables for use with code coverage

### GOCOVERDIR

​        Directory into which to write code coverage data files generated by running a "go build -cover" binary. Requires that GOEXPERIMENT=coverageredesign is enabled.

​	 用于写入由运行 "go build -cover" 二进制文件生成的代码覆盖数据文件的目录。要求启用 GOEXPERIMENT=coverageredesign。

### 特殊用途的环境变量 Special-purpose environment variables

### GCCGOTOOLDIR

​        If set, where to find gccgo tools, such as cgo. The default is based on how gccgo was configured.

​	 如果设置，用于查找 gccgo 工具的位置，例如 cgo。默认值基于 gccgo 的配置方式。

### GOEXPERIMENT

​        Comma-separated list of toolchain experiments to enable or disable. The list of available experiments may change arbitrarily over time. See src/internal/goexperiment/flags.go for currently valid values. Warning: This variable is provided for the development and testing of the Go toolchain itself. Use beyond that purpose is unsupported.

​	 要启用或禁用的工具链实验的逗号分隔列表。可用实验列表可能会随时间任意更改。请参见 src/internal/goexperiment/flags.go 以获取当前有效的值。警告：此变量提供给 Go 工具链自身的开发和测试。超出该目的的使用不受支持。

### GOROOT_FINAL

​        The root of the installed Go tree, when it is installed in a location other than where it is built. File names in stack traces are rewritten from GOROOT to GOROOT_FINAL.

​	 已安装的 Go 树的根目录，当安装在与构建位置不同的位置时。堆栈跟踪中的文件名会从 GOROOT 重写为 GOROOT_FINAL。

### GO_EXTLINK_ENABLED

​        Whether the linker should use external linking mode when using -linkmode=auto with code that uses cgo. Set to 0 to disable external linking mode, 1 to enable it.

​	链接器在使用带有 cgo 的代码的 -linkmode=auto 时是否应使用外部链接模式。设置为 0 以禁用外部链接模式，设置为 1 以启用。

### GIT_ALLOW_PROTOCOL

​        Defined by Git. A colon-separated list of schemes that are allowed to be used with git fetch/clone. If set, any scheme not explicitly mentioned will be considered insecure by 'go get'.Because the variable is defined by Git, the default value cannot be set using 'go env -w'.

​	Git 定义。允许与 git fetch/clone 一起使用的方案的冒号分隔列表。如果设置，任何未明确提到的方案将被 'go get' 视为不安全，因为该变量是由 Git 定义的，无法使用 'go env -w' 设置默认值。



## Additional 

Additional information available from 'go env' but not read from the environment:

​	'go env' 提供的附加信息，但未从环境中读取：

### GOEXE

​        The executable file name suffix (".exe" on Windows, "" on other systems).

​	 可执行文件名后缀（在 Windows 上为 ".exe"，在其他系统上为空字符串）。

### GOGCCFLAGS

​        A space-separated list of arguments supplied to the CC command.

​	 由 CC 命令提供的一组空格分隔的参数列表。

### GOHOSTARCH

​        The architecture (GOARCH) of the Go toolchain binaries.

​	 Go 工具链二进制文件的架构（GOARCH）。

### GOHOSTOS

​        The operating system (GOOS) of the Go toolchain binaries.

​	 Go 工具链二进制文件的操作系统（GOOS）。

### GOMOD

​        The absolute path to the go.mod of the main module.If module-aware mode is enabled, but there is no go.mod, GOMOD will be os.DevNull ("/dev/null" on Unix-like systems, "NUL" on Windows). If module-aware mode is disabled, GOMOD will be the empty string.

​	主模块的 go.mod 的绝对路径。如果启用了模块感知模式，但没有 go.mod，则 GOMOD 将是 os.DevNull（在类 Unix 系统上为 "/dev/null"，在 Windows 上为 "NUL"）。如果禁用了模块感知模式，则 GOMOD 将为空字符串。

### GOTOOLDIR

​        The directory where the go tools (compile, cover, doc, etc...) are installed.

​	 安装 go 工具（compile、cover、doc 等）的目录。

### GOVERSION

​        The version of the installed Go tree, as reported by runtime.Version.

​	已安装的 Go 树的版本，由 runtime.Version 报告。
