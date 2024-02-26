+++
title = "go help testfunc"
date = 2023-12-12T14:13:21+08:00
type = "docs"
weight = 670
description = ""
isCJKLanguage = true
draft = false

+++

​	

The 'go test' command expects to find test, benchmark, and example functions in the "*_test.go" files corresponding to the package under test.

​	`go test` 命令期望在与待测试的包对应的 "*_test.go" 文件中找到测试、基准测试和示例函数。

A test function is one named TestXxx (where Xxx does not start with a lower case letter) and should have the signature,

​	测试函数以 `TestXxx` 命名（其中 `Xxx` 不以小写字母开头），其签名应为：

```go
    func TestXxx(t *testing.T) { ... }
```

A benchmark function is one named BenchmarkXxx and should have the signature,

​	基准测试函数以 `BenchmarkXxx` 命名，其签名应为：

```go
    func BenchmarkXxx(b *testing.B) { ... }
```

A fuzz test is one named FuzzXxx and should have the signature,

​	模糊测试函数以 `FuzzXxx` 命名，其签名应为：

```go
    func FuzzXxx(f *testing.F) { ... }
```

An example function is similar to a test function but, instead of using `*testing.T` to report success or failure, prints output to os.Stdout. If the last comment in the function starts with "Output:" then the output is compared exactly against the comment (see examples below). If the last comment begins with "Unordered output:" then the output is compared to the comment, however the order of the lines is ignored. An example with no such comment is compiled but not executed. An example with no text after "Output:" is compiled, executed, and expected to produce no output.

​	示例函数与测试函数类似，但不使用 `*testing.T` 来报告成功或失败，而是将输出打印到 `os.Stdout`。如果函数的最后一个注释以 "Output:" 开头，那么输出将与注释进行精确比较（请参阅下面的示例）。如果最后一个注释以 "Unordered output:" 开头，那么输出将与注释进行比较，但会忽略行的顺序。没有这种注释的示例会被编译但不会被执行。在 "Output:" 之后没有文本的示例会被编译、执行，并期望不产生输出。

Godoc displays the body of ExampleXxx to demonstrate the use of the function, constant, or variable Xxx. An example of a method M with receiver type T or *T is named ExampleT_M. There may be multiple examples for a given function, constant, or variable, distinguished by a trailing _xxx, where xxx is a suffix not beginning with an upper case letter.

​	Godoc 将显示 `ExampleXxx` 的主体，以演示函数、常量或变量 `Xxx` 的用法。带有接收器类型 `T` 或 `*T` 的方法 `M` 的示例被命名为 `ExampleT_M`。对于给定函数、常量或变量，可以有多个示例，它们通过后缀 `_xxx` 区分，其中 `xxx` 是不以大写字母开头的后缀。

Here is an example of an example:

​	以下是一个示例的示例：

```go
    func ExamplePrintln() {
            Println("The output of\nthis example.")
            // Output: The output of
            // this example.
    }
```

Here is another example where the ordering of the output is ignored:

​	以下是另一个示例，其中忽略了输出的顺序：

```go
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

The entire test file is presented as the example when it contains a single example function, at least one other function, type, variable, or constant declaration, and no tests, benchmarks, or fuzz tests.

​	当包含单个示例函数、至少一个其他函数、类型、变量或常量声明，并且没有测试、基准测试或模糊测试时，整个测试文件将作为示例呈现。

See the documentation of the testing package for more information.

​	有关更多信息，请参阅 testing 包的文档。
