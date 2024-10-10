+++
title = "从源码安装 Go"
weight = 3
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Installing Go from source - 从源码安装 Go

> 原文：[https://go.dev/doc/install/source](https://go.dev/doc/install/source)

本主题描述了如何从源代码构建和运行 Go。要使用安装程序进行安装，请参阅[下载和安装](../InstallingGo)。

## 简介

​	Go 是一个开源项目，采用 [BSD 风格的许可证](https://go.dev/LICENSE)下发布。本文档介绍了如何检出源代码，在自己的机器上构建并运行它们。

​	大多数用户不需要这样做，而是按照[下载和安装](../InstallingGo)中的描述，从预编译的二进制包中安装，这个过程要简单得多。但是，如果您想帮助开发那些预编译包中的内容，请继续阅读。

​	有两个官方`Go编译器工具链`。本文档主要介绍 `gc` Go 编译器和工具。关于如何在`gccgo`上工作的信息，这是一个使用`GCC`后端的更传统的编译器，请参阅[设置和使用gccgo](../SettingUpAndUsingGccgo)。

`Go编译器`支持以下指令集：

- `amd64`, `386`

  `x86`指令集，64位（`amd64`）和32位（`386`）。

- `arm64`, `arm`

  `ARM`指令集，64位（`AArch64`）和32位。

- `loong64`

  64位的`LoongArch`指令集。

- `mips64`, `mips64le`, `mips`, `mipsle`

  `MIPS`指令集，大端、小端，64位和32位。

- `ppc64`, `ppc64le`

  64位`PowerPC`指令集，大端、小端。

- `riscv64`

  64位`RISC-V`指令集。

- `s390x`

  The IBM z/Architecture.

- `wasm`

  [WebAssembly](https://webassembly.org/).



​	编译器可以针对AIX, Android, DragonFly BSD, FreeBSD, Illumos, Linux, macOS/iOS (Darwin), NetBSD, OpenBSD, Plan 9, Solaris, 和Windows操作系统（尽管不是都支持所有操作系统的架构）。

​	被认为是 "第一类 "的端口列表可以在[第一类端口](https://go.dev/wiki/PortingPolicy#first-class-ports)wiki页面上找到。

​	支持的全部组合在下面关于[环境变量](#optional-environment-variables)的讨论中列出。

​	关于[整个系统的要求](../InstallingGo)，请参见主安装页面。以下附加限制适用于仅从源代码构建的系统：

- 对于 PowerPC 64 位的 Linux，最小支持的内核版本是 `2.6.37`，这意味着 Go 在这些系统上不支持 CentOS 6。

## 为引导程序安装 Go 编译器二进制文件

​	Go工具链是用Go编写的。要构建它，您需要安装一个 Go 编译器。执行工具初始构建的脚本会在 `$PATH` 中寻找 "go "命令，所以只要您在系统中安装了 Go 并在 `$PATH` 中配置了 Go，您就可以从源代码上构建 Go。或者如果您愿意，您可以将`$GOROOT_BOOTSTRAP`设置为Go安装的根目录，用来构建新的Go工具链；`$GOROOT_BOOTSTRAP/bin/go`应该是使用的go命令。

​	获得引导工具链有四种可用的方法：

- 下载Go的最新二进制版本。
- 使用已安装Go的系统`交叉编译`一个工具链。
- 使用 `gccgo`。
- 从 Go 1.4 编译一个工具链，这是`最后一个用 C 语言编写`的 Go 编译器。

下面将详细介绍这些方法。

### 从二进制版本引导工具链

​	要使用二进制发布版作为引导工具链，请参阅[下载页面](https://go.dev/dl/)或使用任何其他打包的 Go 发行版。

### 从交叉编译的源代码中引导工具链

​	要从源码交叉编译引导工具链，这在 Go 1.4 没有针对的系统（例如 `linux/ppc64le`）上是必要的，请在不同的系统上安装 Go 并运行 [bootstrap.bash](https://go.dev/src/bootstrap.bash)。

当以下列方式运行时（例如）：

```shell
$ GOOS=linux GOARCH=ppc64 ./bootstrap.bash
```

​	`bootstrap.bash`为`GOOS/GOARCH`的组合交叉编译一个工具链，将生成的树放在`././go-${GOOS}-${GOARCH}-bootstrap`。该树可以被复制到指定目标类型的机器上，并将其用作`GOROOT_BOOTSTRAP`来引导本地构建。

### 使用gccgo引导工具链

To use gccgo as the bootstrap toolchain, you need to arrange for `$GOROOT_BOOTSTRAP/bin/go` to be the go tool that comes as part of gccgo 5. For example on Ubuntu Vivid:

​	要使用`gccgo`作为引导工具链，您需要设置`$GOROOT_BOOTSTRAP/bin/go`为`gccgo 5`中go工具的一部分。例如，在`Ubuntu Vivid`上：

```shell
$ sudo apt-get install gccgo-5
$ sudo update-alternatives --set go /usr/bin/go-5
$ GOROOT_BOOTSTRAP=/usr ./make.bash
```

### 从C 源码引导工具链

​	要从C源代码构建一个引导工具链，可以使用git分支`release-branch.go1.4`或[go1.4-bootstrap-20171003.tar.gz](https://dl.google.com/go/go1.4-bootstrap-20171003.tar.gz)，其中包含Go 1.4的源代码以及保持工具在较新的操作系统上运行的累积补丁。(Go 1.4是最后一个用C语言编写工具链的发行版。)解压Go 1.4源代码后，`cd`到`src`子目录，在环境中设置`CGO_ENABLED=0`，然后运行`make.bash`（或者，在Windows中上运行`make.bat`）。

​	一旦 Go 1.4 源代码被解压到您的 `GOROOT_BOOTSTRAP` 目录中，您必须将这个 git 克隆实例检出到分支 `releas-branch.go1.4`。特别是，不要试图在后面的 "Fetch the repository. （获取版本库） "步骤中重复使用这个 git 克隆。go1.4 引导工具链`必须能够`正确地遍历go1.4的源代码，它假定这些源代码存在于这个版本库根目录下。

> 请注意，Go 1.4并不像后来的Go版本那样可以在所有系统上运行。特别是，`Go 1.4 不支持当前版本的 macOS`。在这样的系统上，必须使用其他方法获得引导工具链。

## 如果需要，安装Git

​	要执行下一个步骤，您必须安装Git。(在继续进行之前，请检查您是否有`git`命令)。

​	如果您没有安装好Git，请按照[Git下载](https://git-scm.com/downloads)页面上的说明进行操作。

## (可选）安装一个C语言编译器

​	要构建一个支持 `cgo` 的 Go 安装程序，`即允许 Go 程序导入 C 库`，必须先安装一个 C 编译器，如 `gcc` 或 `clang`。使用系统上的任何标准安装方法进行安装。

To build without `cgo`, set the environment variable `CGO_ENABLED=0` before running `all.bash` or `make.bash`.

​	要在没有 `cgo` 的情况下进行构建，请在运行 `all.bash` 或 `make.bash` 之前设置环境变量 `CGO_ENABLED=0`。

## 获取存储库

​	切换到您打算安装 Go 的目录，并确保 `goroot` 目录不存在。然后克隆存储库并查看最新的发布标签（例如`go1.12`）：

```shell
$ git clone https://go.googlesource.com/go goroot
$ cd goroot
$ git checkout <tag>
```

其中`<tag>`是发行版的版本字符串。

​	Go将被安装在它被检出的目录中。例如，如果Go被检出在`$HOME/goroot`，可执行文件将被安装在`$HOME/goroot/bin`中。该目录可以有任何名称，但请注意，如果 Go 被检出在 `$HOME/go`，它将与 `$GOPATH` 的默认位置冲突。参见下面的 [GOPATH](#set-up-your-work-environment)。

> 提醒一下
>
> ​	如果您选择同时从源代码编译bootstrap二进制文件（在前面的章节中），此时您仍然需要再次`git clone`（检出最新的`<tag>`），因为您必须保持您的go1.4仓库的独立性。


## (可选）切换到主分支

​	如果您打算修改 go 源代码，并[贡献您的修改](https://go.dev/doc/contribute.html)给项目，那么就把您的仓库从 release tag 上移到 master（开发）分支上。否则，跳过这一步。

```shell
$ git checkout master
```

## 安装 Go

要构建 Go 发行版，请运行

```shell
$ cd src
$ ./all.bash
```

(要在Windows下构建，请使用`all.bat`。)

如果一切顺利，它将完成打印输出，如：

```
ALL TESTS PASSED

---
Installed Go for linux/amd64 in /home/you/go.
Installed commands in /home/you/go/bin.
*** You need to add /home/you/go/bin to your $PATH. ***
```

​	其中最后几行的细节反映了安装过程中使用的操作系统、架构和根目录。

​	关于控制构建的方法的更多信息，请参见下面关于[环境变量](#optional-environment-variables)的讨论。 `all.bash` (或 `all.bat`) 为 Go 运行重要的测试，这可能比简单地构建 Go 需要花费更多时间。如果您不想运行测试套件，请使用 `make.bash` (或 `make.bat`) 来代替。

## 测试您的安装

通过构建一个简单的程序来检查Go的安装是否正确。

创建一个名为 `hello.go` 的文件，并将以下程序放入其中：

```go title="hello.go"
package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
}
```

然后用`go`工具运行它：

```shell
$ go run hello.go
hello, world
```

如果您看到 "hello, world "的信息，那么Go已经正确安装。

## 设置您的工作环境

您就快完成了。您只需要再做一点设置。

[**How to Write Go Code**  --  Learn how to set up and use the Go tools](../HowToWriteGoCode){ .md-button }

[How to Write Go Code](../HowToWriteGoCode) 文档提供了使用 Go 工具的`基本设置说明`。

## 安装其他工具

​	一些 Go 工具（包括 `gopls`）的源代码保存在 `golang.org/x/tools` 存储库中。要安装其中一个工具（本例中是`gopls`）。

```shell
$ go install golang.org/x/tools/gopls@latest
```

## 社区资源

​	在[帮助页面](https://go.dev/help/)上列出的常见社区资源都有活跃的开发者，可以帮助您解决安装或开发工作中的问题。对于那些想了解最新情况的人来说，还有一个邮件列表：[golang-checkins](https://groups.google.com/group/golang-checkins)，它可以收到一条汇总了每次对 Go 仓库签入的消息。

Bug 可以通过[Go问题跟踪器](https://go.dev/issue/new)报告。

## 跟进发布

​	新的版本会在 [golang-announce](https://groups.google.com/group/golang-announce) 邮件列表中公布。每个公告都会提到最新的发布标签，例如`go1.9`。

​	若要将现有的树更新到最新的版本，您可以运行：

```shell
$ cd go/src
$ git fetch
$ git checkout <tag>
$ ./all.bash
```

其中`<tag>`是版本的字符串。

## 可选的环境变量

​	Go 编译环境可以通过`环境变量`来定制。构建时不需要任何环境变量，但您可能希望设置一些来覆盖默认值。

- ```
  $GOROOT
  ```

  Go 树的 root，通常是 `$HOME/go1.X`。它的值在编译时被内置到树中，并默认为运行 `all.bash` 的目录的父目录。除非您想在版本库的多个本地副本之间切换，否则没有必要设置这个。

- ```
  $GOROOT_FINAL
  ```

  当`$GOROOT`没有明确设置时，安装的二进制文件和脚本所假定的值。它默认为 `$GOROOT` 的值。如果您想在一个位置构建 Go 树，但在构建后将其转移到其他地方，请将 `$GOROOT_FINAL` 设置为最终的位置。

- ```
  $GOPATH
  ```

  这是Go 发行版以外的 **Go 项目**通常被检出的目录。例如，`golang.org/x/tools` 可能检出到 `$GOPATH/src/golang.org/x/tools`。Go 发行版以外的可执行文件被安装在 `$GOPATH/bin`（或 `$GOBIN`，如果已设置）。模块被下载并缓存在`$GOPATH/pkg/mod`中。

  `$GOPATH`的默认位置是`$HOME/go`，通常没有必要明确设置`GOPATH`。然而，如果您已经将 Go 发行版检出到 `$HOME/go`，您必须将 `GOPATH` 设置到其他位置以避免冲突。

- ```
  $GOBIN
  ```

  这是 [go 命令](https://go.dev/cmd/go)安装 Go 发行版以外的**可执行文件**的安装目录。例如，`go install golang.org/x/tools/cmd/godoc@latest`：下载、构建和安装到`$GOBIN/godoc`。默认情况下，`$GOBIN`是`$GOPATH/bin`（如果没有设置`GOPATH`，则是`$HOME/go/bin`）。安装后，您要把这个目录添加到您的`$PATH`中，这样您就可以使用已安装的工具。

  **注意**Go发行版的可执行文件安装在`$GOROOT/bin`（用于由人调用的可执行文件）或`$GOTOOLDIR`（用于由go命令调用的可执行文件；默认为`$GOROOT/pkg/$GOOS_$GOARCH`）而不是`$GOBIN`。

- ```
  $GOOS
  ```

  目标操作系统和编译架构的名称。这些默认值分别为`$GOHOSTOS`和`$GOHOSTARCH`的值（如下所述）。

  `$GOOS` 的选择有：`android`，`darwin`， `dragonfly`，`freebsd`，`illumos`，`ios`，`js`，`linux`，`netbsd`，`openbsd`， `plan9`，`solaris` 和`windows`。

  `$GOARCH`的选择有：`amd64` (64-bit x86, the most mature port)， `386` (32-bit x86)， `arm` (32-bit ARM)， `arm64` (64-bit ARM)， `ppc64le` (PowerPC 64-bit, little-endian)，`ppc64` (PowerPC 64-bit, big-endian)， `mips64le` (MIPS 64-bit, little-endian)， `mips64` (MIPS 64-bit, big-endian)， `mipsle` (MIPS 32-bit, little-endian)，`mips` (MIPS 32-bit, big-endian)，`s390x` (IBM System z 64-bit, big-endian)， 和`wasm` (WebAssembly 32-bit).

  `$GOOS`和`$GOARCH`的有效组合有：

  | `$GOOS`     | `$GOARCH`  |
  | ----------- | ---------- |
  | `aix`       | `ppc64`    |
  | `android`   | `386`      |
  | `android`   | `amd64`    |
  | `android`   | `arm`      |
  | `android`   | `arm64`    |
  | `darwin`    | `amd64`    |
  | `darwin`    | `arm64`    |
  | `dragonfly` | `amd64`    |
  | `freebsd`   | `386`      |
  | `freebsd`   | `amd64`    |
  | `freebsd`   | `arm`      |
  | `illumos`   | `amd64`    |
  | `ios`       | `arm64`    |
  | `js`        | `wasm`     |
  | `linux`     | `386`      |
  | `linux`     | `amd64`    |
  | `linux`     | `arm`      |
  | `linux`     | `arm64`    |
  | `linux`     | `loong64`  |
  | `linux`     | `mips`     |
  | `linux`     | `mipsle`   |
  | `linux`     | `mips64`   |
  | `linux`     | `mips64le` |
  | `linux`     | `ppc64`    |
  | `linux`     | `ppc64le`  |
  | `linux`     | `riscv64`  |
  | `linux`     | `s390x`    |
  | `netbsd`    | `386`      |
  | `netbsd`    | `amd64`    |
  | `netbsd`    | `arm`      |
  | `openbsd`   | `386`      |
  | `openbsd`   | `amd64`    |
  | `openbsd`   | `arm`      |
  | `openbsd`   | `arm64`    |
  | `plan9`     | `386`      |
  | `plan9`     | `amd64`    |
  | `plan9`     | `arm`      |
  | `solaris`   | `amd64`    |
  | `windows`   | `386`      |
  | `windows`   | `amd64`    |
  | `windows`   | `arm`      |
  | `windows`   | `arm64`    |
  
- `$GOHOSTOS` and `$GOHOSTARCH`

  主机操作系统和编译架构的名称。这些默认为本地系统的操作系统和架构。

  有效的选择与上面列出的`$GOOS`和`$GOARCH`相同。指定的值必须与本地系统兼容。例如，在X86系统上，您不应该把`$GOHOSTARCH`设置为`arm`。

- `$GO386` (仅适用于`386`，默认值为`sse2`)

  这个变量控制`gc`如何实现浮点计算。

  - `GO386=softfloat`：使用软件浮点运算；应支持所有x86芯片（Pentium MMX或更高版本）。
  - `GO386=sse2`：使用SSE2进行浮点运算；有更好的性能，但只适用于Pentium 4/Opteron/Athlon 64或更高版本。
  
- `$GOARM` (仅适用于arm；如果在目标处理器上构建，默认为自动检测，如果不是，则为6)

  这设置了运行时应该针对的ARM浮点协处理器架构版本。如果您在目标系统上编译，它的值将被自动检测。

  - `GOARM=5`：使用软件浮点；当CPU没有VFP协处理器时
  - `GOARM=6`：只使用VFPv1；交叉编译时，则为默认值；通常是ARM11或更好的内核（也支持VFPv2或更好的）。
  - `GOARM=7`： 使用VFPv3；通常是Cortex-A内核

  如果有疑问，可以不设置这个变量，如果需要，可以在第一次运行Go可执行文件时调整它。[Go社区维基](https://go.dev/wiki)上的[GoARM](https://go.dev/wiki/GoArm)页面包含了关于Go的ARM支持的进一步细节。

- ```
  $GOAMD64 (仅适用于 amd64；默认为 v1)
  ```

  这设置了要编译的微架构级别。有效值有`v1`(默认)、`v2`、`v3`、`v4`。请参见 [the Go wiki MinimumRequirements page](https://go.dev/wiki/MinimumRequirements#amd64)。 

- ```
  $GOMIPS (仅适用于mips64和mips64le)
  ```

  这些变量设置是否使用浮点指令。设置为 "`hardfloat` "可以使用浮点指令；这是默认的。设置为 "`softfloat` "则使用软浮点（soft floating point）。

- ```
  $GOPPC64 (仅适用于ppc64和ppc64le)
  ```

  这个变量设置编译器将针对的处理器级别（即指令集架构版本）。默认是`power8`。

  - `GOPPC64=power8`：生成ISA v2.07指令
  - `GOPPC64=power9`：生成ISA v3.00指令。
  
- ```
  $GOWASM (仅适用于wasm)
  ```

  这个变量是一个逗号分隔的[实验性WebAssembly特性](https://github.com/WebAssembly/proposals)列表，允许编译后的WebAssembly二进制文件使用。默认情况是不使用实验性特征。

  - `GOWASM=satconv`：生成[饱和（非捕获）浮点数到int的转换](https://github.com/WebAssembly/nontrapping-float-to-int-conversions/blob/master/proposals/nontrapping-float-to-int-conversion/Overview.md)。
  - `GOWASM=signext`：生成[符号扩展运算符](https://github.com/WebAssembly/sign-extension-ops/blob/master/proposals/sign-extension-ops/Overview.md)

注意，`$GOARCH`和`$GOOS`标识的是`目标环境`，而不是您正在运行的环境。实际上，您总是在进行交叉编译。所谓架构，我们指的是目标环境可以运行的二进制文件的种类：一个运行仅有32位操作系统的x86-64系统必须将`GOARCH`设置为`386`，而不是`amd64`。

​	如果您选择覆盖默认值，请在您的shell配置文件（`$HOME/.bashrc`，`$HOME/.profile`，或类似的）中设置这些变量。设置的内容可能是这样的

```
export GOARCH=amd64
export GOOS=linux
```

不过，要重申的是，构建、安装和开发 Go 树不需要设置这些变量（这些变量指的是`$GOARCH` 和 `$GOOS` ？？）。