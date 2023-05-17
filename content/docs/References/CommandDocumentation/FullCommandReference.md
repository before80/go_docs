+++
title = "全部命令"
date = 2023-05-17T09:59:21+08:00
weight = 1
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/cmd](https://pkg.go.dev/cmd)



| **Name**                                                     | **Synopsis**  简述                                           |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [addr2line](../addr2line)                                    | `addr2line`是对`GNU addr2line`工具的最小模拟，刚好足以支持pprof。 |
| [api](../api)                                                | `api` 计算一组 Go 包的导出 API。                             |
| [asm](../asm)                                                | `asm`，通常以 "`go tool asm` "进行调用，它将源文件组装成一个对象文件，该文件以参数源文件的基名命名，并带有`.o`后缀。 |
| [buildid](../buildid)                                        | `buildid` 显示或更新存储在 Go 包或二进制文件中的构建 ID。    |
| [cgo](../cgo)                                                | `cgo` 可以创建调用 C 代码的 Go 包。                          |
| [compile](../compile)                                        | `compile` ，通常以 "`go tool compile` "调用，编译一个由命令行上命名的文件组成的单一 Go 包。 |
| [cover](../cover)                                            | `cover` 是一个用于分析由 "`go test -coverprofile=cover.out` "生成的覆盖率配置文件的程序。 |
| [dist](../dist)                                              | `dist` 帮助引导、构建和测试 Go 发行版。                      |
| [doc](../doc)                                                | `doc`（通常以 `go doc` 的形式运行）接受零个、一个或两个参数。 |
| [fix](../fix)                                                | `fix` 找到使用旧 API 的 Go 程序，并将其改写为使用新的 API。  |
| [go](../go)                                                  | `go` 是一个管理 Go 源代码的工具。                            |
| [gofmt](../gofmt)                                            | `gofmt` 用于格式化 Go 程序。                                 |
| [link](../link)                                              | `link`，通常以 "`go tool link` "的方式调用，读取包main的Go存档或对象，以及它的依赖项，并将它们组合成可执行的二进制文件。 |
| [nm](../nm)                                                  | `nm` 列出对象文件、存档、可执行文件等所定义或使用的符号。    |
| [objdump](../objdump)                                        | `objdump` 反汇编可执行文件。                                 |
| [pack](../pack)                                              | `pack` 是传统 Unix `ar` 工具的一个简单版本。                 |
| [pprof](../pprof)                                            | `pprof` 解释并显示 Go 程序的配置文件。                       |
| [test2json](../test2json)                                    | `test2json` 将 Go 测试输出转换为机器可读的 JSON 流。         |
| [trace](../trace)                                            | `trace` 是一个用于查看跟踪文件的工具。                       |
| [vet](../vet)                                                | `vet` 检查 Go 源代码并报告可疑的结构，例如实参与格式字符串不一致的 `Printf` 调用。 |
| **internal**                                                 |                                                              |
| [archive](https://pkg.go.dev/cmd/internal/archive@go1.19.3)  | `archive` 包实现了对 Go 工具链生成的归档文件的读取。         |
| [bio](https://pkg.go.dev/cmd/internal/bio@go1.19.3)          | `bio` 包实现了 Go 工具链中使用的常见 I/O 抽象。              |
| [browser](https://pkg.go.dev/cmd/internal/browser@go1.19.3)  | `browser`包提供了与用户的浏览器进行交互的实用工具。          |
| [buildid](https://pkg.go.dev/cmd/internal/buildid@go1.19.3)  |                                                              |
| [codesign](https://pkg.go.dev/cmd/internal/codesign@go1.19.3) | `codesign`包为`Mach-O`文件的临时代码签名提供基本功能。       |
| [dwarf](https://pkg.go.dev/cmd/internal/dwarf@go1.19.3)      | `dwarf`包生成`DWARF`调试信息。                               |
| [edit](https://pkg.go.dev/cmd/internal/edit@go1.19.3)        | `edit`包实现了对字节切片的基于位置的缓冲编辑。               |
| [gcprog](https://pkg.go.dev/cmd/internal/gcprog@go1.19.3)    | `gcprog`包为打包的 GC 指针位图实现了一个编码器（被称为`GC`程序）。 |
| [goobj](https://pkg.go.dev/cmd/internal/goobj@go1.19.3)      |                                                              |
| [notsha256](https://pkg.go.dev/cmd/internal/notsha256@go1.19.3) | `notsha256`包实现了`NOTSHA256`算法，这是一个散列，定义为 `SHA256`的按位 NOT。 |
| [obj](https://pkg.go.dev/cmd/internal/obj@go1.19.3)          |                                                              |
| [obj/arm](https://pkg.go.dev/cmd/internal/obj/arm@go1.19.3)  |                                                              |
| [obj/arm64](https://pkg.go.dev/cmd/internal/obj/arm64@go1.19.3) | `arm64`包实现了一个`ARM64`汇编器。                           |
| [obj/loong64](https://pkg.go.dev/cmd/internal/obj/loong64@go1.19.3) |                                                              |
| [obj/mips](https://pkg.go.dev/cmd/internal/obj/mips@go1.19.3) |                                                              |
| [obj/ppc64](https://pkg.go.dev/cmd/internal/obj/ppc64@go1.19.3) | `ppc64` 包实现了一个 `PPC64` 汇编器，它将 `Go asm` 汇编成 `Power ISA 3.0B` 所定义的相应 `PPC64` 指令。 |
| [obj/riscv](https://pkg.go.dev/cmd/internal/obj/riscv@go1.19.3) |                                                              |
| [obj/s390x](https://pkg.go.dev/cmd/internal/obj/s390x@go1.19.3) |                                                              |
| [obj/wasm](https://pkg.go.dev/cmd/internal/obj/wasm@go1.19.3) |                                                              |
| [obj/x86](https://pkg.go.dev/cmd/internal/obj/x86@go1.19.3)  |                                                              |
| [objabi](https://pkg.go.dev/cmd/internal/objabi@go1.19.3)    |                                                              |
| [objfile](https://pkg.go.dev/cmd/internal/objfile@go1.19.3)  | `objfile`包实现了对操作系统特定可执行文件的可移植访问。      |
| [osinfo](https://pkg.go.dev/cmd/internal/osinfo@go1.19.3)    | `osinfo`包提供了操作系统元数据。                             |
| [pkgpath](https://pkg.go.dev/cmd/internal/pkgpath@go1.19.3)  | `pkgpath` 包决定了 `gccgo/GoLLVM` 符号所使用的包路径。       |
| [quoted](https://pkg.go.dev/cmd/internal/quoted@go1.19.3)    | `quoted`包提供了字符串操作的实用工具。                       |
| [src](https://pkg.go.dev/cmd/internal/src@go1.19.3)          |                                                              |
| [sys](https://pkg.go.dev/cmd/internal/sys@go1.19.3)          |                                                              |
| [test2json](https://pkg.go.dev/cmd/internal/test2json@go1.19.3) | `test2json`包实现了测试二进制输出到JSON的转换。              |
| [traceviewer](https://pkg.go.dev/cmd/internal/traceviewer@go1.19.3) | `traceviewer`包提供了`Chrome trace viewer`所使用的JSON数据结构的定义。 |