+++
title = "教程：fuzzing 入门"
weight = 12
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Tutorial: Getting started with fuzzing - 教程：fuzzing 入门

> 原文：[https://go.dev/doc/tutorial/fuzz](https://go.dev/doc/tutorial/fuzz)

​	本教程介绍了Go中的模糊测试的基本知识。通过模糊测试，针对您的测试运行随机数据，试图找到漏洞或导致崩溃的输入。通过fuzzing可以发现的一些漏洞的例子有：`SQL注入`、`缓冲区溢出`、`拒绝服务`和`跨站脚本攻击`。

​	在本教程中，您将为一个简单的函数写一个模糊测试，运行go命令，并调试和修复代码中的问题。

​	关于本教程中的术语帮助，请参见 [Go Fuzzing glossary](../../UsingAndUnderstandingGo/Fuzzing#glossary)。

您将通过以下几个部分取得进展：

1. [为您的代码创建一个文件夹](#为您的代码创建一个文件夹)。
3. [添加测试代码](#添加测试代码)。
4. [添加一个单元测试](#添加一个单元测试)。
5. [添加一个模糊测试](#添加一个模糊测试)。
6. [修复两个bug](#修复无效字符串的错误)。
7. [探索其他资源](#总结)。

注意：关于其他教程，请参见 [Tutorials](../Tutorials)。

> 注意：Go模糊测试目前支持[Go模糊测试文档](../../UsingAndUnderstandingGo/Fuzzing#要求)中列出的内置类型子集，未来将增加对更多内置类型的支持。

## 前提条件

- 安装 Go 1.18 或更高版本。关于安装说明，请参阅 [Installing Go](../InstallingGo)。
- 编辑代码的工具。任何您拥有的文本编辑器都可以使用。
- 命令终端。在 Linux 和 Mac 上使用任何终端，以及在 Windows 上使用 `PowerShell` 或 `cmd`，Go 都能很好地工作。
- 支持模糊处理的环境。目前只有在`AMD64`和`ARM64`架构上可以使用覆盖检测技术进行模糊化。

## 为您的代码创建一个文件夹

首先，为您要写的代码创建一个文件夹。

a. 打开一个命令提示符，切换到您的主目录。

在Linux或Mac上：

```shell
$ cd
```

在Windows上：

```shell
C:\> cd %HOMEPATH%
```

​	本教程的其余部分将显示一个`$`作为提示符。您所使用的命令在Windows上也会起作用。

b. 在命令提示符下，为您的代码创建一个名为`fuzz`的目录。

```shell
$ mkdir fuzz
$ cd fuzz
```

c. 创建一个模块来存放您的代码。

​	运行 `go mod init` 命令，给它您的新代码的模块路径。

```shell
$ go mod init example/fuzz
go: creating new go.mod: module example/fuzz
```

注意：对于生产代码，您可以根据自己的需要指定一个更具体的模块路径。更多信息，请务必参阅[管理依赖项](../../UsingAndUnderstandingGo/ManagingDependencies#命名一个模块)。

​	接下来，您将添加一些简单的代码来反转一个字符串，我们稍后将对其进行模糊处理。

## 添加测试代码

在这一步，您将添加一个函数来反转一个字符串。

### 编写代码

a. 使用您的文本编辑器，在`fuzz`目录下创建一个名为`main.go`的文件。

b. 在`main.go`文件的顶部，粘贴以下包声明。

```go
package main
```

A standalone program (as opposed to a library) is always in package `main`.

一个独立的程序（相对于一个库）总是在package main中。

c. 在包声明的下面，粘贴以下函数声明。

```go linenums="1" hl_lines="2 2"
func Reverse(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}
```

​	这个函数将接受一个`string`，每次循环一个`byte`，并在最后返回反转的字符串。

注意：这段代码是基于`golang.org/x/example`中的`stringutil.Reverse`函数。

d. 在`main.go`的顶部，在包声明的下面，粘贴以下`main`函数，以初始化一个字符串，反转它，打印输出，然后重复。

```go linenums="1"
func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev := Reverse(input)
    doubleRev := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q\n", rev)
    fmt.Printf("reversed again: %q\n", doubleRev)
}
```

​	这个函数将运行一些`Reverse`（反转）操作，然后将输出打印到命令行。这有助于查看正在运行的代码，而且有可能用于调试。

e. `main`函数使用`fmt`包，所以您需要导入它。

第一行代码应该是这样的：

```go linenums="1"
package main

import "fmt"
```

### 运行该代码

在包含`main.go`的目录下的命令行中，运行该代码。

```shell
$ go run .
original: "The quick brown fox jumped over the lazy dog"
reversed: "god yzal eht revo depmuj xof nworb kciuq ehT"
reversed again: "The quick brown fox jumped over the lazy dog"
```

​	您可以看到原始字符串，反转后的结果，然后是再次反转的结果，它等价于原始字符串。

​	现在代码正在运行，是时候测试它了。

## 添加一个单元测试

在这一步，您将为`Reverse`函数写一个基本的单元测试。

### 编写代码

a. 使用您的文本编辑器，在`fuzz`目录下创建一个名为`reverse_test.go`的文件。

b. 将以下代码粘贴到`reverse_test.go`中。

```go title="reverse_test.go" linenums="1"
package main

import (
    "testing"
)

func TestReverse(t *testing.T) {
    testcases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
    }
    for _, tc := range testcases {
        rev := Reverse(tc.in)
        if rev != tc.want {
                t.Errorf("Reverse: %q, want %q", rev, tc.want)
        }
    }
}
```

这个简单的测试将断言列出的输入字符串将被正确反转的字符串。

### 运行代码

使用`go test`运行单元测试

```shell
$ go test
PASS
ok      example/fuzz  0.013s
```

接下来，您将把单元测试改为模糊测试。

## 添加一个模糊测试

​	`单元测试有其局限性`，即每个输入必须由开发人员添加到测试中。模糊测试的一个好处是，它为您的代码提供输入，并可以识别出您提供的测试用例没有达到的边缘情况。

​	在本节中，您将把单元测试转换为模糊测试，这样您就能以更少的工作量产生更多的输入了！

注意，您可以把单元测试（unit tests）、基准测试（benchmarks）和模糊测试（fuzz tests）放在同一个`*_test.go`文件中，但在这个例子中，您将把单元测试转换为模糊测试。

### 编写代码

在您的文本编辑器中，用以下模糊测试替换 `reverse_test.go` 中的单元测试。

```go linenums="1"
func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev := Reverse(orig)
        doubleRev := Reverse(rev)
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

​	`Fuzzing` 也有一些限制。在您的单元测试中，您可以预测`Reverse`函数的预期输出，并验证实际输出是否符合这些预期。

​	例如，在测试用例`Reverse("Hello, world")`中，单元测试指定返回为 `"dlrow ,olleH"`。

​	当模糊测试时，您无法预测预期的输出，因为您无法控制输入。

​	然而，您可以在模糊测试中验证`Reverse`函数的一些属性。在这个模糊测试中被检查的两个属性是：

1. 对字符串进行两次反转将保留原始值
4. 反向字符串将其状态保留为有效的 UTF-8

注意单元测试和模糊测试之间的语法差异：

- 该函数以`FuzzXxx`开头，而不是`TestXxx`，并且使用*`testing.F`而不是*`testing.T`
- 在您期望看到`t.Run`执行的地方，您看到的是`f.Fuzz`，它接收一个模糊目标函数，其参数是`*testing.T`和要模糊的类型。您的单元测试的输入被作为种子语料库的输入使用`f.Add`提供。

确保新的包，`unicode/utf8`已经被导入。

```go linenums="1" hl_lines="5 5"
package main

import (
    "testing"
    "unicode/utf8"
)
```

将单元测试转换为模糊测试后，现在是再次运行测试的时候了。

### 运行代码

a. 运行模糊测试而不 fuzzing 它，以确保种子（seed）输入通过。

```shell
$ go test
PASS
ok      example/fuzz  0.013s
```

​	如果您在该文件中有其他测试，而您只想运行模糊测试，您也可以运行 `go test -run=FuzzReverse`。

b. 使用 fuzzing 运行 `FuzzReverse` ，看看任何随机生成的字符串输入是否会导致失败。这是用 `go test` 和一个新的标志 `-fuzz` 来执行的。

```shell
$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/3 completed
fuzz: elapsed: 0s, gathering baseline coverage: 3/3 completed, now fuzzing with 8 workers
fuzz: minimizing 38-byte failing input file...
--- FAIL: FuzzReverse (0.01s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:20: Reverse produced invalid UTF-8 string "\x9c\xdd"

    Failing input written to testdata/fuzz/FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
    To re-run:
    go test -run=FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a
FAIL
exit status 1
FAIL    example/fuzz  0.030s
```

​	在 fuzzing 时发生了故障，导致问题的输入被写入种子语料库文件，在下次调用`go test` 时将被运行，即使没有`-fuzz`标志。要查看导致失败的输入，可以用文本编辑器打开写在`testdata/fuzz/FuzzReverse`目录下的语料库文件。您的种子语料库文件`可能包含不同的字符串`，但其格式是相同的。

```shell
go test fuzz v1
string("泃")
```

​	语料库文件的第一行表示编码版本。接下来的每一行表示构成语料库条目的每个类型的值。由于fuzz目标只接受1个输入，所以版本后面只有1个值。

c. 再次运行`go test`，不使用`-fuzz`标志；将使用新的失败的种子语料库条目：

```shell
$ go test
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/af69258a12129d6cbba438df5d5f25ba0ec050461c116f777e77ea7c9a0d217a (0.00s)
        reverse_test.go:20: Reverse produced invalid string
FAIL
exit status 1
FAIL    example/fuzz  0.016s
```

既然我们的测试已经失败了，现在是时候进行调试了。

## 修复无效字符串的错误

​	在这一节中，您将对失败进行调试，并修复这个错误。

​	在继续之前，请随意花一些时间来思考，并尝试自己修复这个问题。

### 诊断错误

​	有几种不同的方法可以调试这个错误。如果您使用`VS Code`作为您的文本编辑器，您可以[设置您的调试器](https://github.com/golang/vscode-go/blob/master/docs/debugging.md)来调查。

​	在本教程中，我们将把有用的调试信息记录到您的终端。

首先，考虑[utf8.ValidString]({{< ref "/stdLib/unicode/utf8#func-validstring">}})的文档。

```
ValidString reports whether s consists entirely of valid UTF-8-encoded runes. `ValidString`报告s是否完全由有效的utf -8编码的符文组成。
```

​	目前的`Reverse`函数是`逐个字节地 （byte-by-byte）`反转字符串，这就是我们的问题所在。为了保留原始字符串的UTF-8编码的符文，我们必须`逐个符文地（rune-by-rune）`反转字符串。

​	为了检查为什么输入（在本例中是中文字符 "`泃`"）会导致`Reverse`在反转时产生一个无效的字符串，您可以检查反转字符串中的符文数量。

#### 编写代码

​	在您的文本编辑器中，将`FuzzReverse`中的 fuzz 目标替换为以下内容。

```go linenums="1"
f.Fuzz(func(t *testing.T, orig string) {
    rev := Reverse(orig)
    doubleRev := Reverse(rev)
    t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
    if orig != doubleRev {
        t.Errorf("Before: %q, after: %q", orig, doubleRev)
    }
    if utf8.ValidString(orig) && !utf8.ValidString(rev) {
        t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
    }
})
```

​	如果发生错误，或者用`-v`执行测试，这个`t.Logf`行将打印到命令行，这可以帮助您调试这个特殊问题。

#### 运行代码

使用`go test`运行测试

```shell
$ go test
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0 (0.00s)
        reverse_test.go:16: Number of runes: orig=1, rev=3, doubleRev=1
        reverse_test.go:21: Reverse produced invalid UTF-8 string "\x83\xb3\xe6"
FAIL
exit status 1
FAIL    example/fuzz    0.598s
```

​	整个种子语料库使用的字符串中，每个字符都是一个字节。但是，像 "`泃` "这样的字符可能需要几个字节。因此，逐个字节地反转字符串将使`多字节的字符`失效。

> 注意：如果您对Go如何处理字符串感到好奇，请阅读博文《[Strings, bytes, runes and characters in Go]({{< ref "/goBlog/2013/StringsBytesRunesAndCharactersInGo">}})》以加深理解。

​	在对这个错误有了更深入的了解后，在`Reverse`函数中纠正这个错误。

### 修正错误

​	为了纠正`Reverse`函数，让我们通过符文而不是通过字节来遍历字符串。

#### 编写代码

​	在您的文本编辑器中，将现有的`Reverse()`函数替换为以下内容。

```go linenums="1" hl_lines="2 2"
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

​	关键的区别是`Reverse`现在是对字符串中的每个`rune`（符文）进行迭代，而不是每个`byte`（字节）。

#### 运行代码

a. 使用`go test`运行测试

```shell
$ go test
PASS
ok      example/fuzz  0.016s
```

现在测试通过了!

b. 再次用`go test -fuzz`测试一下，看看是否有新的bug。

```shell
$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/37 completed
fuzz: minimizing 506-byte failing input file...
fuzz: elapsed: 0s, gathering baseline coverage: 5/37 completed
--- FAIL: FuzzReverse (0.02s)
    --- FAIL: FuzzReverse (0.00s)
        reverse_test.go:33: Before: "\x91", after: "�"

    Failing input written to testdata/fuzz/FuzzReverse/1ffc28f7538e29d79fce69fef20ce5ea72648529a9ca10bea392bcff28cd015c
    To re-run:
    go test -run=FuzzReverse/1ffc28f7538e29d79fce69fef20ce5ea72648529a9ca10bea392bcff28cd015c
FAIL
exit status 1
FAIL    example/fuzz  0.032s
```

​	我们可以看到，`经过两次反转后的字符串与原来的不同`。这一次的输入本身是无效的unicode。如果我们是用字符串进行模糊处理，这怎么可能呢？

我们再来调试一下。

## 修复两次反转错误

​	在这一节中，您将调试两次反转失败，并修复这个错误。

​	在继续之前，请随意花一些时间来思考这个问题，并尝试自己修复这个问题。

### 诊断错误

​	像以前一样，您有几种方法可以调试这个故障。在这种情况下，使用[debugger （调试器）](https://github.com/golang/vscode-go/blob/master/docs/debugging.md)将是一个很好的方法。

​	在本教程中，我们将在`Reverse`函数中记录有用的调试信息。

​	仔细观察反转的字符串来发现错误。在Go中，[字符串是一个只读的字节切片]({{< ref "/goBlog/2013/StringsBytesRunesAndCharactersInGo">}})，可以包含`无效UTF-8的字节`。原始字符串是一个字节切片，其中有一个字节`'\x91'`。当输入字符串被设置为`[]rune`时，Go将字节切片编码为UTF-8，并将该字节替换为UTF-8字符�。当我们将替换的UTF-8字符与输入的字节切片进行比较时，它们显然是不相等的。

#### 编写代码

a. 在您的文本编辑器中，将`Reverse`函数替换为以下内容。

```go linenums="1" hl_lines="2 4"
func Reverse(s string) string {
    fmt.Printf("input: %q\n", s)
    r := []rune(s)
    fmt.Printf("runes: %q\n", r)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

这将帮助我们了解在将字符串转换为符文片时出了什么问题。

#### 编写代码

​	这一次，我们只想运行失败的测试，以便检查日志。要做到这一点，我们将使用`go test -run`。

```shell
$ go test -run=FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0
input: "\x91"
runes: ['�']
input: "�"
runes: ['�']
--- FAIL: FuzzReverse (0.00s)
    --- FAIL: FuzzReverse/28f36ef487f23e6c7a81ebdaa9feffe2f2b02b4cddaa6252e87f69863046a5e0 (0.00s)
        reverse_test.go:16: Number of runes: orig=1, rev=1, doubleRev=1
        reverse_test.go:18: Before: "\x91", after: "�"
FAIL
exit status 1
FAIL    example/fuzz    0.145s
```

​	要在`FuzzXxx/testdata`中运行一个特定的语料库条目，您可以向`-run`提供`{FuzzTestName}/{filename}`。这在调试时可能会有帮助。

​	知道了输入是无效的 unicode，让我们在`Reverse`函数中修复这个错误。

### 修复错误

​	为了解决这个问题，如果`Reverse`的输入不是有效的UTF-8，让我们返回一个错误。

#### 编写代码

a. 在您的文本编辑器中，将现有的`Reverse`函数替换为以下内容。

```go linenums="1"
func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}
```

​	如果输入的字符串包含无效的UTF-8字符，这一改变将返回一个错误。

b. 由于`Reverse`函数现在返回一个错误，请修改`main`函数以丢弃额外的错误值。将现有的`main`函数改为以下内容。

```go linenums="1"
func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev, revErr := Reverse(input)
    doubleRev, doubleRevErr := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}
```

​	这些对`Reverse`的调用应该返回一个`nil`错误，当输入的字符串是有效的UTF-8。

c. 您将需要导入错误和`unicode/utf8`包。`main.go`中的`import`语句应该如下所示。

```go linenums="1"
import (
    "errors"
    "fmt"
    "unicode/utf8"
)
```

d. 修改 `reverse_test.go` 文件以检查错误，如果返回时产生了错误，则跳过测试。

```go linenums="1" hl_lines="8 10"
func FuzzReverse(f *testing.F) {
    testcases := []string {"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc)  // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
             return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

​	您也可以调用`t.Skip()`而不是返回，以停止执行该模糊输入。

#### 运行代码

a. 使用`go test`运行测试

```shell
$ go test
PASS
ok      example/fuzz  0.019s
```

b. Fuzz it with `go test -fuzz=Fuzz`, then after a few seconds has passed, stop fuzzing with `ctrl-C`. 用`go test -fuzz=Fuzz`进行模糊测试，然后在几秒钟后，用ctrl-C停止模糊测试。

```shell
$ go test -fuzz=Fuzz
fuzz: elapsed: 0s, gathering baseline coverage: 0/38 completed
fuzz: elapsed: 0s, gathering baseline coverage: 38/38 completed, now fuzzing with 4 workers
fuzz: elapsed: 3s, execs: 86342 (28778/sec), new interesting: 2 (total: 35)
fuzz: elapsed: 6s, execs: 193490 (35714/sec), new interesting: 4 (total: 37)
fuzz: elapsed: 9s, execs: 304390 (36961/sec), new interesting: 4 (total: 37)
...
fuzz: elapsed: 3m45s, execs: 7246222 (32357/sec), new interesting: 8 (total: 41)
^Cfuzz: elapsed: 3m48s, execs: 7335316 (31648/sec), new interesting: 8 (total: 41)
PASS
ok      example/fuzz  228.000s
```

​	除非您通过`-fuzztime`标志，否则`模糊测试会一直运行到遇到失败的输入`。默认情况下，`如果没有失败发生，就会永远运行下去`，而且可以用ctrl-C中断这个过程。

c. 用`go test -fuzz=Fuzz -fuzztime 30s`对其进行模糊处理，如果没有发现故障，将模糊处理30秒后退出。

```shell
$ go test -fuzz=Fuzz -fuzztime 30s
fuzz: elapsed: 0s, gathering baseline coverage: 0/5 completed
fuzz: elapsed: 0s, gathering baseline coverage: 5/5 completed, now fuzzing with 4 workers
fuzz: elapsed: 3s, execs: 80290 (26763/sec), new interesting: 12 (total: 12)
fuzz: elapsed: 6s, execs: 210803 (43501/sec), new interesting: 14 (total: 14)
fuzz: elapsed: 9s, execs: 292882 (27360/sec), new interesting: 14 (total: 14)
fuzz: elapsed: 12s, execs: 371872 (26329/sec), new interesting: 14 (total: 14)
fuzz: elapsed: 15s, execs: 517169 (48433/sec), new interesting: 15 (total: 15)
fuzz: elapsed: 18s, execs: 663276 (48699/sec), new interesting: 15 (total: 15)
fuzz: elapsed: 21s, execs: 771698 (36143/sec), new interesting: 15 (total: 15)
fuzz: elapsed: 24s, execs: 924768 (50990/sec), new interesting: 16 (total: 16)
fuzz: elapsed: 27s, execs: 1082025 (52427/sec), new interesting: 17 (total: 17)
fuzz: elapsed: 30s, execs: 1172817 (30281/sec), new interesting: 17 (total: 17)
fuzz: elapsed: 31s, execs: 1172817 (0/sec), new interesting: 17 (total: 17)
PASS
ok      example/fuzz  31.025s
```

Fuzzing通过了!

除了`-fuzz`标志外，还有几个新的标志被添加到`go test`中，可以在[文档](../../UsingAndUnderstandingGo/Fuzzing#自定义设置-custom-settings)中查看。

## 总结

​	做得很好! 您刚刚向自己介绍了Go中的fuzzing。

​	下一步是在您的代码中选择一个您想模糊处理的函数，并尝试使用它! 如果 fuzzing 在您的代码中发现了一个bug，可以考虑把它加入[trophy case （战利品箱、奖杯箱）](https://github.com/golang/go/wiki/Fuzzing-trophy-case)。

​	如果您遇到了任何问题或有关于特性的想法，[file an issue](https://github.com/golang/go/issues/new/?&labels=fuzz)。

​	对于有关该特性的讨论和一般反馈，您也可以参与Gophers Slack的[#fuzzing channel](https://gophers.slack.com/archives/CH5KV1AKE)。

​	请查看 [go.dev/security/fuzz](../../UsingAndUnderstandingGo/Fuzzing) 的文档，以进一步阅读。

## 完整的代码

{{< tabpane text=true >}}
{{< tab header="main.go" >}}

```go title="main.go" linenums="1"
package main

import (
    "errors"
    "fmt"
    "unicode/utf8"
)

func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev, revErr := Reverse(input)
    doubleRev, doubleRevErr := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}
```

{{< /tab >}}

{{< tab header="reverse_test.go" >}}

```go title="reverse_test.go" linenums="1"
package main

import (
    "testing"
    "unicode/utf8"
)

func FuzzReverse(f *testing.F) {
    testcases := []string{"Hello, world", " ", "!12345"}
    for _, tc := range testcases {
        f.Add(tc) // Use f.Add to provide a seed corpus
    }
    f.Fuzz(func(t *testing.T, orig string) {
        rev, err1 := Reverse(orig)
        if err1 != nil {
            return
        }
        doubleRev, err2 := Reverse(rev)
        if err2 != nil {
            return
        }
        if orig != doubleRev {
            t.Errorf("Before: %q, after: %q", orig, doubleRev)
        }
        if utf8.ValidString(orig) && !utf8.ValidString(rev) {
            t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
        }
    })
}
```

{{< /tab >}}

{{< /tabpane >}}

