+++
title = "go help private"
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 640
description = ""
isCJKLanguage = true
draft = false

+++

​	

The go command defaults to downloading modules from the public Go module mirror at proxy.golang.org. It also defaults to validating downloaded modules, regardless of source, against the public Go checksum database at sum.golang.org. These defaults work well for publicly available source code.

​	go 命令默认从公共 Go 模块镜像 proxy.golang.org 下载模块。它还默认使用公共 Go 校验和数据库 sum.golang.org 验证下载的模块，无论其来源如何。这些默认设置对于公开可用的源代码非常有效。

The GOPRIVATE environment variable controls which modules the go command considers to be private (not available publicly) and should therefore not use the proxy or checksum database. The variable is a comma-separated list of glob patterns (in the syntax of Go's path.Match) of module path prefixes. For example,

​	GOPRIVATE 环境变量控制 go 命令认为哪些模块是私有的（不公开可用），因此不应使用代理或校验和数据库。该变量是一个以逗号分隔的 glob 模式列表（采用 Go 的 path.Match 语法），用于指定模块路径前缀。例如，

        GOPRIVATE=*.corp.example.com,rsc.io/private

causes the go command to treat as private any module with a path prefix matching either pattern, including git.corp.example.com/xyzzy, rsc.io/private, and rsc.io/private/quux.

会使 go 命令将匹配任一模式的路径前缀的模块视为私有，包括 git.corp.example.com/xyzzy、rsc.io/private 和 rsc.io/private/quux。

For fine-grained control over module download and validation, the GONOPROXY and GONOSUMDB environment variables accept the same kind of glob list and override GOPRIVATE for the specific decision of whether to use the proxy and checksum database, respectively.

​	为了对模块的下载和验证进行更精细的控制，GONOPROXY 和 GONOSUMDB 环境变量接受相同类型的 glob 列表，并分别覆盖 GOPRIVATE，用于特定的决策，即是否使用代理和校验和数据库。

For example, if a company ran a module proxy serving private modules, users would configure go using:

​	例如，如果公司运行了一个提供私有模块的模块代理，用户可以配置 go 如下：

        GOPRIVATE=*.corp.example.com
        GOPROXY=proxy.example.com
        GONOPROXY=none

The GOPRIVATE variable is also used to define the "public" and "private" patterns for the GOVCS variable; see 'go help vcs'. For that usage, GOPRIVATE applies even in GOPATH mode. In that case, it matches import paths instead of module paths.

​	GOPRIVATE 变量还用于定义 GOVCS 变量的 "public" 和 "private" 模式；参见 'go help vcs'。对于此用法，即使在 GOPATH 模式下，GOPRIVATE 也适用。在这种情况下，它匹配导入路径而不是模块路径。

The 'go env -w' command (see 'go help env') can be used to set these variables for future go command invocations.

​	'go env -w' 命令（参见 'go help env'）可用于设置这些变量，以供将来的 go 命令调用使用。

For more details, see https://golang.org/ref/mod#private-modules.

​	有关更多详细信息，请参阅 https://golang.org/ref/mod#private-modules。
