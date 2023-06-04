+++
title = "go 集成测试的代码覆盖率"
weight = 96
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Code coverage for Go integration tests - Go集成测试的代码覆盖率

[https://go.dev/blog/integration-test-coverage](https://go.dev/blog/integration-test-coverage)

Than McIntosh
8 March 2023

2023年3月8日

​	代码覆盖率工具帮助开发者确定在执行给定测试套件时有多少源代码被执行（覆盖）。

​	Go自1.2版本开始就[提供了对包级别代码覆盖率](../../stdLib/2013/TheCoverStory)的支持，使用"go test"命令的"-cover"标志。

​	这个工具在大多数情况下工作良好，但对于较大的Go应用程序有一些弱点。对于这样的应用程序，开发者经常编写"integration（集成）"测试，以验证整个程序的行为（除了包级别的单元测试）。

​	这种类型的测试通常涉及构建一个完整的应用程序二进制文件，然后在一组代表性输入（或生产负载下，如果是服务器）上运行该二进制文件，以确保所有组件包都正确地一起工作，而不是在隔离中测试各个包。

​	由于集成测试二进制文件是使用"go build"而不是"go test"构建的，因此Go的工具没有提供任何方便的方法来收集这些测试的覆盖率分析文件，直到现在。

​	从Go 1.20开始，您现在可以使用"go build -cover"构建带有覆盖率工具化的的程序，然后将这些工具化的二进制文件提供给集成测试，以扩展覆盖率测试的范围。

​	在本博客文章中，我们将举例说明这些新功能的工作原理，并概述从集成测试收集覆盖率分析文件的用例和工作流程。

## 示例

​	我们将以一个非常小的示例程序为例，为其编写一个简单的集成测试，然后从集成测试中收集覆盖率分析文件。

​	为了进行这个练习，我们将使用来自`gitlab.com/golang-commonmark/mdtool`的"mdtool" Markdown处理工具。这是一个演示程序，旨在展示客户端如何使用`gitlab.com/golang-commonmark/markdown`包，这是一个Markdown到HTML的转换库。

## 为 mdtool 进行设置

​	首先让我们下载 "mdtool" 的一个副本（为了使这些步骤可重复，我们选择一个特定的版本）：

```sh
$ git clone https://gitlab.com/golang-commonmark/mdtool.git
...
$ cd mdtool
$ git tag example e210a4502a825ef7205691395804eefce536a02f
$ git checkout example
...
$
```

## 一个简单的集成测试

​	现在我们将为 "mdtool" 编写一个简单的集成测试；我们的测试将构建 "mdtool" 二进制文件，然后在一组输入的 markdown 文件上运行它。这个非常简单的脚本在测试数据目录中的每个文件上运行 "mdtool" 二进制文件，检查它是否产生了一些输出并且没有崩溃。

```sh
$ cat integration_test.sh
#!/bin/sh
BUILDARGS="$*"
#
# 如果以下任何命令未能成功完成，则终止测试。
#
set -e
#
# 下载一些测试输入（'website'仓库包含各种*.md文件）。
#
if [ ! -d testdata ]; then
  git clone https://go.googlesource.com/website testdata
  git -C testdata tag example 8bb4a56901ae3b427039d490207a99b48245de2c
  git -C testdata checkout example
fi
#
# 为测试目的构建mdtool二进制文件。
#
rm -f mdtool.exe
go build $BUILDARGS -o mdtool.exe .
#
# 在来自'testdata'的一组输入文件上运行该工具。
#
FILES=$(find testdata -name "*.md" -print)
N=$(echo $FILES | wc -w)
for F in $FILES
do
  ./mdtool.exe +x +a $F > /dev/null
done
echo "finished processing $N files, no crashes"
$
```

以下是我们的测试的一个示例运行：

```sh
$ /bin/sh integration_test.sh
...
finished processing 380 files, no crashes
$
```

成功：我们已经验证了 "mdtool" 二进制文件成功转化了一组输入文件……但是我们实际上运行了工具的多少源代码呢？在下一节中，我们将收集一个覆盖率分析文件来找出答案。

## 使用集成测试收集覆盖率数据

​	让我们编写另一个包装脚本来调用前面的脚本，但是为了收集覆盖率，它会构建该工具，然后后处理生成的分析文件：

```sh
$ cat wrap_test_for_coverage.sh
#!/bin/sh
set -e
PKGARGS="$*"
#
# 准备
#
rm -rf covdatafiles
mkdir covdatafiles
#
# 将"-cover"传递给脚本以进行覆盖率构建，然后
# 设置GOCOVERDIR并运行。
#
GOCOVERDIR=covdatafiles \
  /bin/sh integration_test.sh -cover $PKGARGS
#
# 后处理生成的分析文件。
#
go tool covdata percent -i=covdatafiles
$
```

​	关于该包装脚本，需要注意的一些重要事项：

- 当运行`integration_test.sh`时，它传入"-cover"标志，这给了我们一个覆盖率工具"mdtool.exe"二进制文件。 
- 它将GOCOVERDIR环境变量设置为一个目录，覆盖率数据文件将被写入其中。
- 当测试完成时，它运行"go tool covdata percent"来生成语句覆盖率报告。 

​	当我们运行这个新的包装脚本时，输出如下：

```sh
$ /bin/sh wrap_test_for_coverage.sh
...
    gitlab.com/golang-commonmark/mdtool coverage: 48.1% of statements
$
# Note: covdatafiles now contains 381 files.
```

​	太好了！（Voila!）我们现在对于集成测试如何使用"mdtool"应用程序的源代码有了一些了解。

​	如果我们更改测试套件以增强其功能，然后进行第二次覆盖率收集运行，我们将在覆盖率报告中看到更改的影响。例如，假设我们通过向`integration_test.sh`添加以下两行额外的内容来改进我们的测试：

```sh
./mdtool.exe +ty testdata/README.md  > /dev/null
./mdtool.exe +ta < testdata/README.md  > /dev/null
```

​	再次运行覆盖测试包装脚本：

```sh
$ /bin/sh wrap_test_for_coverage.sh
finished processing 380 files, no crashes
    gitlab.com/golang-commonmark/mdtool coverage: 54.6% of statements
$
```

​	我们可以看到我们的更改效果：语句覆盖率从48％增加到54％。

## 选择要覆盖的包 

​	默认情况下，"go build -cover"只会对正在构建的 Go 模块的包进行工具化处理，而在本例中，这个包是 `gitlab.com/golang-commonmark/mdtool`。但在某些情况下，将覆盖工具化扩展到其他包是有用的；可以通过将"-coverpkg"传递给"go build -cover"来实现这一点。

​	对于我们的示例程序，"mdtool"实际上大部分只是围绕 `gitlab.com/golang-commonmark/markdown` 包的包装器，因此将 `markdown` 包包括在被工具化的包集中是有意义的。

​	这是"mdtool"的 `go.mod` 文件：

```sh
$ head go.mod
module gitlab.com/golang-commonmark/mdtool

go 1.17

require (
    github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
    gitlab.com/golang-commonmark/markdown v0.0.0-20211110145824-bf3e522c626a
)
```

​	我们可以使用"-coverpkg"标志来控制哪些包被选择用于包含在覆盖分析中，以包含上述任何一个依赖项。下面是一个例子：

```sh
$ /bin/sh wrap_test_for_coverage.sh -coverpkg=gitlab.com/golang-commonmark/markdown,gitlab.com/golang-commonmark/mdtool
...
    gitlab.com/golang-commonmark/markdown   coverage: 70.6% of statements
    gitlab.com/golang-commonmark/mdtool coverage: 54.6% of statements
$
```

## 使用覆盖数据文件

​	当覆盖率集成测试完成并写出一组原始数据文件（在我们的示例中，是 `covdatafiles` 目录的内容）后，我们可以以各种方式后处理这些文件。

### 将profiles转换为"-coverprofile"文本格式 

​	在处理单元测试时，可以运行 `go test -coverprofile=abc.txt` 来为给定的覆盖测试运行编写文本格式的覆盖率profile。

​	对于使用 `go build -cover` 构建的二进制文件，可以在事后运行 `go tool covdata textfmt` 命令生成文本格式的profile，该命令会针对 GOCOVERDIR 目录中生成的文件进行处理。

​	完成此步骤后，您可以使用 `go tool cover -func=<file>` 或 `go tool cover -html=<file>` 来解释/可视化数据，就像使用 `go test -coverprofile` 一样。

示例：

```sh
$ /bin/sh wrap_test_for_coverage.sh
...
$ go tool covdata textfmt -i=covdatafiles -o=cov.txt
$ go tool cover -func=cov.txt
gitlab.com/golang-commonmark/mdtool/main.go:40:     readFromStdin   100.0%
gitlab.com/golang-commonmark/mdtool/main.go:44:     readFromFile    80.0%
gitlab.com/golang-commonmark/mdtool/main.go:54:     readFromWeb 0.0%
gitlab.com/golang-commonmark/mdtool/main.go:64:     readInput   80.0%
gitlab.com/golang-commonmark/mdtool/main.go:74:     extractText 100.0%
gitlab.com/golang-commonmark/mdtool/main.go:88:     writePreamble   100.0%
gitlab.com/golang-commonmark/mdtool/main.go:111:    writePostamble  100.0%
gitlab.com/golang-commonmark/mdtool/main.go:118:    handler     0.0%
gitlab.com/golang-commonmark/mdtool/main.go:139:    main        51.6%
total:                          (statements)    54.6%
$
```

### 使用"go tool covdata merge"合并原始profiles

​	每次执行"-cover"构建的应用程序都会将一个或多个数据文件写入到GOCOVERDIR环境变量指定的目录中。如果一个集成测试执行了N次程序执行，则输出目录中将有O(N)个文件。数据文件中通常有大量重复内容，因此可以使用`go tool covdata merge`命令来合并profiles，以压缩数据和/或组合来自不同集成测试运行的数据集。例如：

```sh
$ /bin/sh wrap_test_for_coverage.sh
finished processing 380 files, no crashes
    gitlab.com/golang-commonmark/mdtool coverage: 54.6% of statements
$ ls covdatafiles
covcounters.13326b42c2a107249da22f6e0d35b638.772307.1677775306041466651
covcounters.13326b42c2a107249da22f6e0d35b638.772314.1677775306053066987
...
covcounters.13326b42c2a107249da22f6e0d35b638.774973.1677775310032569308
covmeta.13326b42c2a107249da22f6e0d35b638
$ ls covdatafiles | wc
    381     381   27401
$ rm -rf merged ; mkdir merged ; go tool covdata merge -i=covdatafiles -o=merged
$ ls merged
covcounters.13326b42c2a107249da22f6e0d35b638.0.1677775331350024014
covmeta.13326b42c2a107249da22f6e0d35b638
$
```

​	`go tool covdata merge`操作还接受`-pkg`标志，该标志可用于选择特定的包或包集，如果需要的话。

​	此合并功能还可用于组合来自不同类型的测试运行的结果，包括由其他测试工具生成的运行结果。

## 总结 Wrap-up

​	到此为止：随着1.20版本的发布，Go的覆盖率工具不再仅限于包测试，而是支持从更大的集成测试中收集profiles。我们希望您能充分利用新功能，帮助了解您的更大更复杂的测试的工作情况，以及它们正在测试哪些部分的源代码。

​	请尝试这些新功能，如果遇到问题，请像往常一样在我们的[GitHub问题跟踪器](https://github.com/golang/go/issues)上提出问题。谢谢。
