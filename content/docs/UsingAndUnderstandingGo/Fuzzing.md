+++
title = "Go Fuzzing"
linkTitle = "Go Fuzzing"
weight = 20
date = 2023-05-17T15:03:14+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Go Fuzzing

> 原文：[https://go.dev/security/fuzz/](https://go.dev/security/fuzz/)
>

​	从Go 1.18开始，Go在其标准工具链中支持模糊测试。[OSS-Fuzz支持](https://google.github.io/oss-fuzz/getting-started/new-project-guide/go-lang/#native-go-fuzzing-support)本地Go模糊测试。

​	请尝试使用[Go进行模糊测试的教程](../../GettingStarted/TutorialGettingStartedWithFuzzing)。

## 概述

​	模糊测试是一种自动化测试，它持续操作程序的输入，以查找错误。Go模糊测试使用覆盖率指导，智能地遍历被模糊测试的代码，查找并向用户报告故障。由于它可以触及人类经常错过的边缘情况，模糊测试对于寻找安全漏洞和漏洞特别有价值。

​	下面是一个模糊测试的示例，强调了其主要组成部分。

![Example code showing the overall fuzz test, with a fuzz target within it. Before the fuzz target is a corpus addition with f.Add, and the parameters of the fuzz target are highlighted as the fuzzing arguments.](Fuzzing_img/example.png)

## 编写模糊测试

### 要求

以下是模糊测试必须遵循的规则。

- 模糊测试必须是一个名为`FuzzXxx`的函数，它只接受`*testing.F`，并且没有返回值。
- 模糊测试必须在`*_test.go`文件中才能运行。
- 一个[模糊目标](https://go.dev/security/fuzz/#glos-fuzz-target)必须是对`(*testing.F).Fuzz`的方法调用，它接受一个`*testing.T`作为第一个参数，后跟模糊参数。没有返回值。
- 每个模糊测试必须恰好有一个模糊目标。
- 所有种子语料库条目的类型必须与模糊测试参数的类型相同，顺序也相同。这对于调用`(*testing.F).Add`和模糊测试的testdata/fuzz目录中的任何语料库文件都是如此。
- 模糊参数只能是以下类型：
  - `string`, `[]byte`
  - `int`, `int8`, `int16`, `int32`/`rune`, `int64`
  - `uint`, `uint8`/`byte`, `uint16`, `uint32`, `uint64`
  - `float32`, `float64`
  - `bool`

### 建议 Suggestions

​	下面是一些建议，它们将帮助您充分利用模糊测试。

- 模糊目标应该是快速和确定的，这样模糊测试引擎就能有效地工作，新的故障和代码覆盖率就能轻易地重现。
- 由于模糊目标是以非确定性的顺序在多个工作进程中并行调用的，因此模糊目标的状态不应持续到每次调用结束之后，并且模糊目标的行为也不应依赖于全局状态。

## 运行模糊测试

​	有两种运行模糊测试的模式：作为单元测试（默认的 `go test`），或使用模糊测试（`go test -fuzz=FuzzTestName`）。

​	默认情况下，模糊测试的运行方式与单元测试非常相似。每个[种子语料库条目（seed corpus entry）](#seed-corpus)都会针对模糊测试目标进行测试，在退出前报告任何失败的情况。

​	要启用模糊测试，请在运行 `go test` 时使用 `-fuzz` 标志，并提供一个与单个模糊测试相匹配的正则表达式。默认情况下，该包中的所有其他测试将在模糊测试开始前运行。这是为了确保模糊测试不会报告任何已经被现有测试发现的问题。

​	请注意，您需要决定运行模糊测试的时间长度。如果没有发现任何错误，模糊测试的执行很可能会无限期地运行。未来将支持使用OSS-Fuzz等工具连续运行这些模糊测试，见[Issue #50192](https://github.com/golang/go/issues/50192)。

> 注意
>
> ​	应在支持覆盖率仪器（目前为AMD64和ARM64）的平台上运行模糊测试，这样语料库才能在运行时有意义地增长，并在模糊测试期间覆盖更多代码。

### 命令行输出

​	当模糊测试进行时，[模糊测试引擎（fuzzing engine）](#fuzzing-engine)会生成新的输入并将其运行到提供的模糊目标中。默认情况下，它继续运行，直到发现一个[失败的输入（failing input）](#failing-input)，或者用户取消这个过程（例如用 Ctrl^C）。

​	输出的格式类似于以下内容：

```sh
~ go test -fuzz FuzzFoo
fuzz: elapsed: 0s, gathering baseline coverage: 0/192 completed
fuzz: elapsed: 0s, gathering baseline coverage: 192/192 completed, now fuzzing with 8 workers
fuzz: elapsed: 3s, execs: 325017 (108336/sec), new interesting: 11 (total: 202)
fuzz: elapsed: 6s, execs: 680218 (118402/sec), new interesting: 12 (total: 203)
fuzz: elapsed: 9s, execs: 1039901 (119895/sec), new interesting: 19 (total: 210)
fuzz: elapsed: 12s, execs: 1386684 (115594/sec), new interesting: 21 (total: 212)
PASS
ok      foo 12.692s
```

​	开头几行表示在模糊测试开始前收集 "基准覆盖率（baseline coverage）"。

​	为了收集基准覆盖率，模糊测试引擎会执行[种子语料库（seed corpus）](#seed-corpus)和[生成语料库（generated corpus）](#generated-corpus)，以确保没有发生错误，并了解现有语料库已经提供的代码覆盖率。

​	随后的几行提供了有关当前模糊测试执行的见解：

- elapsed：从进程开始到现在已经过去了多少时间
- execs：针对模糊目标运行的输入总数（自最后一行日志以来的平均execs/sec)
- new interesting：在这次模糊测试执行过程中，被添加到生成的语料库中的 "有趣（interesting）"输入的总数（与整个语料库的总大小有关）。

​	要使输入 "有趣（interesting）"，它必须将代码覆盖范围扩大到现有生成的语料库所能达到的范围。典型的情况是，新的有趣输入的数量在开始时快速增长，最终放缓，随着新的分支被发现，偶尔会有爆发。

​	随着语料库中的输入开始覆盖更多的代码行，您应该期望看到"新的有趣的（new interesting）"数量随着时间的推移逐渐减少，如果模糊测试引擎发现新的代码路径，则偶尔会出现爆发。

### 失败的输入

​	模糊测试可能由于以下几个原因而失败：

- 代码或测试中出现了panic 。
- 模糊目标直接或通过`t.Error`或`t.Fatal`等方法调用了`t.Fail`。
- 发生了一个不可恢复的错误，如`os.Exit`或栈溢出。
- 模糊目标执行时间过长。目前，模糊目标的执行超时时间为 1 秒。这可能由于死锁或无限循环而失败，或由于代码中的预期行为而失败。这就是为什么[建议您的模糊目标要快速的一个原因](#suggestions)。

​	如果出现错误，模糊测试引擎将尝试将输入最小化为仍能产生错误的最小可读值。有关配置此功能的信息，请参阅[自定义设置（custom-settings）](#custom-settings)部分。

​	一旦最小化完成，错误信息将被记录下来，并且输出将以以下形式结束：

```
    Failing input written to testdata/fuzz/FuzzFoo/a878c3134fe0404d44eb1e662e5d8d4a24beb05c3d68354903670ff65513ff49
    To re-run:
    go test -run=FuzzFoo/a878c3134fe0404d44eb1e662e5d8d4a24beb05c3d68354903670ff65513ff49
FAIL
exit status 1
FAIL    foo 0.839s
```

​	模糊测试引擎将这个[失败的输入（failing input）](#failing-input)写入了该模糊测试的种子语料库，现在它将被默认为与`go test`一起运行，一旦该bug被修复，它将作为回归测试。

​	下一步是诊断问题，修复错误，通过重新运行`go test`来验证修复，并提交包含新testdata文件的补丁，以充当回归测试。

### 自定义设置 Custom settings 

​	默认的go命令设置应该适用于大多数模糊测试用例。因此，通常情况下，在命令行上执行的模糊测试应如下所示：

```sh
$ go test -fuzz={FuzzTestName}
```

​	然而，`go`命令在运行模糊测试时确实提供了一些设置。这些设置在 [cmd/go 包](../../References/CommandDocumentation/go)的文档中都有记载。

​	在此强调几个：

- `-fuzztime`：在退出前执行模糊目标的总时间或迭代次数，默认为无限期。
- `-fuzzminimizetime`：在每次最小化尝试中，模糊目标将被执行的时间或迭代次数，默认为60秒。您可以通过在进行模糊测试时设置`-fuzzminimizetime 0`来完全禁用最小化。
- `-parallel`：同时运行的模糊测试进程的数量，默认为`$GOMAXPROCS`。目前，在进行模糊测试时设置`-cpu`无效。

## 语料库文件格式 Corpus file format 

​	语料库文件采用一种特殊的格式进行编码。这种格式用于[种子语料库（seed corpus）](#seed-corpus)和[生成语料库（generated corpus）](#generated-corpus)。

​	以下是一个语料库文件的示例：

```
go test fuzz v1
[]byte("hello\\xbd\\xb2=\\xbc ⌘")
int64(572293)
```

​	第一行是用来告知模糊测试引擎文件的编码版本。虽然目前没有计划未来的编码格式版本，但设计必须支持这种可能性。

​	接下来的每一行都是构成语料库条目的值，如果需要，可以直接复制到Go代码中。

​	在上面的示例中，我们有一个`[]byte`，后跟一个`int64`。这些类型必须与模糊测试的参数完全匹配，按照这个顺序。这些类型的模糊目标应该是这样的：

```
f.Fuzz(func(*testing.T, []byte, int64) {})
```

​	指定您自己的种子语料库值的最简单方法是使用`(*testing.F).Add`方法。在上面的示例中，看起来会是这样的：

```
f.Add([]byte("hello\\xbd\\xb2=\\xbc ⌘"), int64(572293))
```

​	然而，您可能有一些大型二进制文件，不希望将其作为代码复制到测试中，而是作为单独的种子语料库条目保留在`testdata/fuzz/{FuzzTestName}`目录下。golang.org/x/tools/cmd/file2fuzz上的[file2fuzz](https://pkg.go.dev/golang.org/x/tools/cmd/file2fuzz)工具可以用来将这些二进制文件转换成`[]byte`编码的语料库文件。

​	要使用这个工具：

```
$ go install golang.org/x/tools/cmd/file2fuzz@latest
$ file2fuzz
```

## 资源 Resources

（a）教程：

- 试试[Go的模糊测试教程](../../GettingStarted/TutorialGettingStartedWithFuzzing)，深入了解新概念。
- 如果想了解更简短的Go模糊测试入门教程，请看[博文](../../GoBlog/2021/FuzzingIsBetaReady)。

（b）文档：

- [testing包](../../StdLib/testing/index)文档描述了用于编写模糊测试时使用的 `testing.F` 类型。
- [cmd/go 包](../../References/CommandDocumentation/go)文档描述了与模糊测试相关的标志。

（c）技术细节：

- [Design draft 设计草案](https://golang.org/s/draft-fuzzing-design)

- [Proposal 提议](https://golang.org/issue/44551)

## 术语表 Glossary 

### corpus entry 语料库条目

​	语料库中的一个输入，可以在模糊测试时使用。这可以是一个特殊格式的文件，也可以是对`（*testing.F）.Add`的调用。

### coverage guidance 覆盖率指导

​	一种模糊测试方法，它使用代码覆盖率的扩展来确定哪些语料库条目值得保留以供将来使用。

### failing input 失败的输入

​	失败的输入是一个语料库条目，它在与模糊目标运行时将导致错误或恐慌。

### fuzz target 模糊目标

​	对于语料库条目和生成的值执行的模糊测试功能。它是通过向`(*testing.F).Fuzz`传递函数来提供给模糊测试的。

### fuzz test 模糊测试

​	测试文件中的一个函数，其形式为`func FuzzXxx(*testing.F)`，可用于进行模糊测试。

### fuzzing 模糊测试

​	一种自动测试，它不断地操纵程序的输入，以发现问题，如代码可能易受影响的缺陷或[漏洞（vulnerability）](#vulnerability)。

### fuzzing arguments 模糊测试参数

​	将传递给模糊目标的类型，并由[突变器（mutator）](#mutator)进行突变。

### fuzzing engine 模糊测试引擎

 	用于管理模糊测试的工具，包括维护语料库、调用突变器、识别新的覆盖范围和报告失败。

### generated corpus 生成的语料库

​	一个由模糊测试引擎在模糊测试过程中长期维护的语料库，以跟踪进展。它被存储在`$GOCACHE/fuzz`中。这些条目只在模糊测试时使用。

### mutator 突变器

​	一个在模糊测试时使用的工具，它在将语料库条目传递给模糊测试目标之前随机地处理这些条目。

### package 包

​	在同一目录下的源文件的集合，它们被编译在一起。参见Go语言规范中的[包部分](https://go.dev/ref/spec#Packages)。

### seed corpus 种子语料库

​	用户提供的用于模糊测试的语料库，可用于指导模糊引擎。它由模糊测试中`f.Add`调用提供的语料库条目，以及包中testdata/fuzz/{FuzzTestName}目录下的文件组成。无论是否进行模糊测试，这些条目都会在`go test`中默认运行。

### test file 测试文件

​	格式为`xxx_test.go`的文件，可能包含测试、基准测试、示例和模糊测试。

### vulnerability 漏洞

​	代码中对安全敏感的弱点，可被攻击者利用。

## 反馈 Feedback 

​	如果您遇到任何问题或对功能有想法，请提出[问题](https://github.com/golang/go/issues/new?&labels=fuzz)。

​	关于该功能的讨论和一般反馈，您也可以参与Gophers Slack的[#fuzzing频道](https://gophers.slack.com/archives/CH5KV1AKE)。