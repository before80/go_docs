+++
title = "go help buildconstraint"
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 410
description = ""
isCJKLanguage = true
draft = false

+++

A build constraint, also known as a build tag, is a condition under which a file should be included in the package. Build constraints are given by a line comment that begins

​	构建约束，也称为构建标签，是在包中应包含文件的条件。构建约束由以

    //go:build

Constraints may appear in any kind of source file (not just Go), but they must appear near the top of the file, preceded only by blank lines and other line comments. These rules mean that in Go files a build constraint must appear before the package clause.

​	开头的行注释给出。约束可以出现在任何类型的源文件中（不仅仅是 Go 文件），但它们必须出现在文件的顶部附近，只能由空行和其他行注释预先。这些规则意味着在 Go 文件中，构建约束必须出现在包子句之前。

To distinguish build constraints from package documentation, a build constraint should be followed by a blank line.

​	为了区分构建约束和包文档，构建约束应该跟随一个空行。

A build constraint comment is evaluated as an expression containing build tags combined by ||, &&, and ! operators and parentheses.

​	构建约束注释被解释为一个包含由 ||、&& 和 ! 运算符和括号组合的构建标签的表达式。

Operators have the same meaning as in Go.

​	运算符在 Go 中具有相同的含义。

For example, the following build constraint constrains a file to build when the "linux" and "386" constraints are satisfied, or when "darwin" is satisfied and "cgo" is not:

​	例如，以下构建约束约束文件仅在满足 "linux" 和 "386" 约束时构建，或在满足 "darwin" 且 "cgo" 不满足时构建：

```go
//go:build (linux && 386) || (darwin && !cgo)
```

It is an error for a file to have more than one //go:build line.

​	文件具有多个 //go:build 行是错误的。

During a particular build, the following build tags are satisfied:

​	在特定构建期间，以下构建标签得到满足：

- the target operating system, as spelled by runtime.GOOS, set with the
  GOOS environment variable.
- 目标操作系统，由 runtime.GOOS 拼写，使用 GOOS 环境变量设置。
- the target architecture, as spelled by runtime.GOARCH, set with the
  GOARCH environment variable.
- 目标架构，由 runtime.GOARCH 拼写，使用 GOARCH 环境变量设置。
- any architecture features, in the form GOARCH.feature
  (for example, "amd64.v2"), as detailed below.
- 任何架构特性，形式为 GOARCH.feature（例如，"amd64.v2"），如下所述。
- "unix", if GOOS is a Unix or Unix-like system.
- "unix"，如果 GOOS 是 Unix 或类 Unix 系统。
- the compiler being used, either "gc" or "gccgo"
- 正在使用的编译器，要么是 "gc" 要么是 "gccgo"。
- "cgo", if the cgo command is supported (see CGO_ENABLED in
  'go help environment').
- "cgo"，如果支持 cgo 命令（请参见 'go help environment' 中的 CGO_ENABLED）。
- a term for each Go major release, through the current version:
  "go1.1" from Go version 1.1 onward, "go1.12" from Go 1.12, and so on.
- 每个 Go 主要版本的术语，直至当前版本："go1.1"从 Go 版本 1.1 开始，"go1.12"从 Go 1.12 开始，以此类推。
- any additional tags given by the -tags flag (see 'go help build').
- 由 -tags 标志（请参见 'go help build'）给出的任何附加标签。

There are no separate build tags for beta or minor releases.

​	对于 beta 或小版本，没有单独的构建标签。

If a file's name, after stripping the extension and a possible `_test` suffix, matches any of the following patterns:

​	如果文件名（在剥离扩展名和可能的 `_test` 后缀之后）匹配以下任一模式：

​        `*_GOOS`

​        `*_GOARCH`

​        `*_GOOS_GOARCH`

(example: source_windows_amd64.go) where GOOS and GOARCH represent any known operating system and architecture values respectively, then the file is considered to have an implicit build constraint requiring those terms (in addition to any explicit constraints in the file).

(例如：source_windows_amd64.go)，其中 GOOS 和 GOARCH 分别表示任何已知的操作系统和体系结构值，则认为该文件具有隐式构建约束，需要这些术语（除了文件中的任何显式约束）。

Using GOOS=android matches build tags and files as for GOOS=linux in addition to android tags and files.

​	使用 GOOS=android 将匹配构建标签和文件，就像使用 GOOS=linux 一样，除了 android 标签和文件。

Using GOOS=illumos matches build tags and files as for GOOS=solaris in addition to illumos tags and files.

​	使用 GOOS=illumos 将匹配构建标签和文件，就像使用 GOOS=solaris 一样，除了 illumos 标签和文件。

Using GOOS=ios matches build tags and files as for GOOS=darwin in addition to ios tags and files.

​	使用 GOOS=ios 将匹配构建标签和文件，就像使用 GOOS=darwin 一样，除了 ios 标签和文件。

The defined architecture feature build tags are:

​	已定义的架构特性构建标签包括：

- For GOARCH=386, GO386=387 and GO386=sse2 set the 386.387 and 386.sse2 build tags, respectively.
- 对于 GOARCH=386，GO386=387 和 GO386=sse2 分别设置 386.387 和 386.sse2 构建标签。
- For GOARCH=amd64, GOAMD64=v1, v2, and v3 correspond to the amd64.v1, amd64.v2, and amd64.v3 feature build tags.
- 对于 GOARCH=amd64，GOAMD64=v1、v2 和 v3 对应于 amd64.v1、amd64.v2 和 amd64.v3 特性构建标签。
- For GOARCH=arm, GOARM=5, 6, and 7 correspond to the arm.5, arm.6, and arm.7 feature build tags.
- 对于 GOARCH=arm，GOARM=5、6 和 7 对应于 arm.5、arm.6 和 arm.7 特性构建标签。
- For GOARCH=mips or mipsle,GOMIPS=hardfloat and softfloat correspond to the mips.hardfloat and mips.softfloat
  (or mipsle.hardfloat and mipsle.softfloat) feature build tags.
- 对于 GOARCH=mips 或 mipsle，GOMIPS=hardfloat 和 softfloat 对应于 mips.hardfloat 和 mips.softfloat（或 mipsle.hardfloat 和 mipsle.softfloat）特性构建标签。
- For GOARCH=mips64 or mips64le,GOMIPS64=hardfloat and softfloat correspond to the mips64.hardfloat and mips64.softfloat
  (or mips64le.hardfloat and mips64le.softfloat) feature build tags.
- 对于 GOARCH=mips64 或 mips64le，GOMIPS64=hardfloat 和 softfloat 对应于 mips64.hardfloat 和 mips64.softfloat（或 mips64le.hardfloat 和 mips64le.softfloat）特性构建标签。
- For GOARCH=ppc64 or ppc64le, GOPPC64=power8, power9, and power10 correspond to the ppc64.power8, ppc64.power9, and ppc64.power10 (or ppc64le.power8, ppc64le.power9, and ppc64le.power10) feature build tags.
- 对于 GOARCH=ppc64 或 ppc64le，GOPPC64=power8、power9 和 power10 对应于 ppc64.power8、ppc64.power9 和 ppc64.power10（或 ppc64le.power8、ppc64le.power9 和 ppc64le.power10）特性构建标签。
- For GOARCH=wasm, GOWASM=satconv and signext correspond to the wasm.satconv and wasm.signext feature build tags.
- 对于 GOARCH=wasm，GOWASM=satconv 和 signext 对应于 wasm.satconv 和 wasm.signext 特性构建标签。

For GOARCH=amd64, arm, ppc64, and ppc64le, a particular feature level sets the feature build tags for all previous levels as well.

​	对于 GOARCH=amd64、arm、ppc64 和 ppc64le，特定的特性级别会设置先前所有级别的特性构建标签。

For example, GOAMD64=v2 sets the amd64.v1 and amd64.v2 feature flags.

​	例如，GOAMD64=v2 设置 amd64.v1 和 amd64.v2 特性标志。

This ensures that code making use of v2 features continues to compile when, say, GOAMD64=v4 is introduced.

​	这确保了使用 v2 特性的代码在引入例如 GOAMD64=v4 时仍然能够编译。

Code handling the absence of a particular feature level should use a negation:

​	处理缺少特定特性级别的代码应使用否定：

```go
//go:build !amd64.v2
```

To keep a file from being considered for any build:

​	要防止文件被考虑用于任何构建：

```go
//go:build ignore
```

(Any other unsatisfied word will work as well, but "ignore" is conventional.)

（任何其他未满足的单词也可以，但 "ignore" 是传统的。）

To build a file only when using cgo, and only on Linux and OS X:

​	要仅在使用 cgo 且仅在 Linux 和 OS X 上构建文件：

```go
//go:build cgo && (linux || darwin)
```

Such a file is usually paired with another file implementing thedefault functionality for other systems, which in this case wouldcarry the constraint:

​	这样的文件通常与另一个文件配对，该文件实现其他系统的默认功能，本例中将携带约束：

```go
//go:build !(cgo && (linux || darwin))
```

Naming a file dns_windows.go will cause it to be included only whenbuilding the package for Windows; similarly, math_386.s will be includedonly when building the package for 32-bit x86.

​	命名文件为 dns_windows.go 将导致仅在为 Windows 构建包时包含它；类似地，math_386.s 将仅在为 32 位 x86 构建包时包含。

Go versions 1.16 and earlier used a different syntax for build constraints,with a "// +build" prefix. The gofmt command will add an equivalent //go:build constraint when encountering the older syntax.

​	Go 版本 1.16 及更早版本使用不同的语法进行构建约束，使用 "// +build" 前缀。gofmt 命令会在遇到旧语法时添加相应的 //go:build 约束。
