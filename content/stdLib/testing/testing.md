+++
title = "testing"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/testing@go1.24.2](https://pkg.go.dev/testing@go1.24.2)

Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the "go test" command, which automates execution of any function of the form

​	`testing`包提供了对Go程序包进行自动化测试的支持。它旨在与"go test"命令一起使用，该命令自动执行以下形式的任何函数：

``` go 
func TestXxx(*testing.T)
```

where Xxx does not start with a lowercase letter. The function name serves to identify the test routine.

其中Xxx不以小写字母开头。函数名称用于标识测试例程。

Within these functions, use the Error, Fail or related methods to signal failure.

​	在这些函数内部，使用Error、Fail或相关方法来发出失败信号。

To write a new test suite, create a file that contains the TestXxx functions as described here, and give that file a name ending in "_test.go". The file will be excluded from regular package builds but will be included when the "go test" command is run.

​	要编写新的测试套件，请创建一个文件，其中包含如上所述的TestXxx函数，并将该文件命名为以"_test.go"结尾的名称。该文件将在常规包构建时排除在外，但在运行"go test"命令时将被包括在内。

The test file can be in the same package as the one being tested, or in a corresponding package with the suffix "_test".

​	测试文件可以与被测试的包在同一个包中，也可以是具有后缀"_test"的对应包中。

If the test file is in the same package, it may refer to unexported identifiers within the package, as in this example:

(包内测试)如果测试文件在同一个包中，则可以引用包内未公开的标识符，如下面的示例：

```go
package abs

import "testing"

func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

If the file is in a separate "_test" package, the package being tested must be imported explicitly and only its exported identifiers may be used. This is known as "black box" testing.

(包外测试)如果测试文件在一个单独的"_test"包中，被测试的包必须显式导入，并且只能使用其导出的标识符。这被称为"黑盒"测试

```go
package abs_test

import (
	"testing"

	"path_to_pkg/abs"
)

func TestAbs(t *testing.T) {
    got := abs.Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

For more detail, run "go help test" and "go help testflag".

​	欲了解更多细节，请运行 "go help test"和 "go help testflag"。

## 基准测试 Benchmarks 

Functions of the form

形如

``` go 
func BenchmarkXxx(*testing.B)
```

are considered benchmarks, and are executed by the "go test" command when its -bench flag is provided. Benchmarks are run sequentially.

的函数被视为基准测试，并在 "go test" 命令提供其 -bench 标志时执行。基准测试按顺序运行。

For a description of the testing flags, see https://golang.org/cmd/go/#hdr-Testing_flags.

​	有关测试标志的说明，，见[Testing flags]({{< ref "/cmd/go#testing-flags">}})。

A sample benchmark function looks like this:

​	基准测试函数的示例如下所示：

``` go 
func BenchmarkRandInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rand.Int()
    }
}
```

The benchmark function must run the target code b.N times. During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably. The output

​	基准测试函数必须运行目标代码 `b.N` 次。在基准测试执行期间，`b.N` 会进行调整，直到基准测试函数持续时间足够长，以便可靠计时。输出

```
BenchmarkRandInt-8   	68453040	        17.8 ns/op
```

means that the loop ran 68453040 times at a speed of 17.8 ns per loop.

表示循环运行了 68453040 次，每次循环的速度为 17.8 纳秒。

If a benchmark needs some expensive setup before running, the timer may be reset:

​	如果基准测试需要在运行之前进行一些昂贵的设置，则可以重置计时器：

``` go  hl_lines="3 3"
func BenchmarkBigLen(b *testing.B) {
    big := NewBig()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        big.Len()
    }
}
```

If a benchmark needs to test performance in a parallel setting, it may use the RunParallel helper function; such benchmarks are intended to be used with the go test -cpu flag:

​	如果基准测试需要在并行设置中测试性能，则可以使用 `RunParallel` 帮助函数；这样的基准测试旨在与 `go test -cpu` 标志一起使用：

``` go  hl_lines="3 3"
func BenchmarkTemplateParallel(b *testing.B) {
    templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
    b.RunParallel(func(pb *testing.PB) {
        var buf bytes.Buffer
        for pb.Next() {
            buf.Reset()
            templ.Execute(&buf, "World")
        }
    })
}
```

A detailed specification of the benchmark results format is given in https://golang.org/design/14313-benchmark-format.

​	基准测试结果格式的详细规范在 [Proposal: Go Benchmark Data Format](../../ProposalGoBenchmarkDataFormat)。

There are standard tools for working with benchmark results at https://golang.org/x/perf/cmd. In particular, https://golang.org/x/perf/cmd/benchstat performs statistically robust A/B comparisons.

​	在[https://golang.org/x/perf/cmd](https://golang.org/x/perf/cmd)中有用于处理基准测试结果的标准工具。特别是，[https://golang.org/x/perf/cmd/benchstat](https://golang.org/x/perf/cmd/benchstat)可以进行统计学上健壮的 A/B 比较。

## Example

The package also runs and verifies example code. Example functions may include a concluding line comment that begins with "Output:" and is compared with the standard output of the function when the tests are run. (The comparison ignores leading and trailing space.) These are examples of an example:

​	该包还可以运行和验证示例代码。Example函数可以包括一个以 "`Output:`"开头的结尾行注释，并在测试运行时与函数的标准输出进行比较。(比较时忽略前导和尾部的空格。) 这些是众多示例中的一个示例：

​	该包还可以运行和验证示例代码。示例函数可以包含以"`Output:`"开头的结尾行注释，并在运行测试时与函数的标准输出进行比较。(比较忽略前导和尾随空格。)以下是众多示例中的一个示例：

``` go  hl_lines="3 3"
func ExampleHello() {
    fmt.Println("hello")
    // Output: hello
}

func ExampleSalutations() {
    fmt.Println("hello, and")
    fmt.Println("goodbye")
    // Output:
    // hello, and
    // goodbye
}
```

The comment prefix "Unordered output:" is like "Output:", but matches any line order:

​	前缀为"`Unordered output:`"的注释与"`Output:`"类似，但匹配任何行顺序：

``` go 
func ExamplePerm() {
    for _, value := range Perm(5) {
        fmt.Println(value)
    }
    // Unordered output: 4
    // 2
    // 1
    // 3
    // 0
}
```

Example functions without output comments are compiled but not executed.

​	没有输出注释的示例函数会被编译，但不会被执行。

The naming convention to declare examples for the package, a function F, a type T and method M on type T are:

​	用于声明包、函数F、类型T和类型T上的方法M的示例的命名规则是：

``` go 
func Example() { ... }
func ExampleF() { ... }
func ExampleT() { ... }
func ExampleT_M() { ... }
```

Multiple example functions for a package/type/function/method may be provided by appending a distinct suffix to the name. The suffix must start with a lower-case letter.

​	可以通过在名称后附加不同的后缀来为包/类型/函数/方法提供多个示例函数。后缀必须以小写字母开头。

``` go 
func Example_suffix() { ... }
func ExampleF_suffix() { ... }
func ExampleT_suffix() { ... }
func ExampleT_M_suffix() { ... }
```

The entire test file is presented as the example when it contains a single example function, at least one other function, type, variable, or constant declaration, and no test or benchmark functions.

​	当测试文件包含单个示例函数、至少一个其他函数、类型、变量或常量声明，并且没有测试或基准测试函数时，整个测试文件被展示为示例。

## 模糊测试 Fuzzing 

'go test' and the testing package support fuzzing, a testing technique where a function is called with randomly generated inputs to find bugs not anticipated by unit tests.

​	'go test' 和 testing 包支持模糊测试，这是一种使用随机生成的输入调用函数以发现单元测试未预期的错误的测试技术。

Functions of the form

形如

``` go 
func FuzzXxx(*testing.F)
```

are considered fuzz tests.

的函数被认为是模糊测试。

For example:

示例：

``` go 
func FuzzHex(f *testing.F) {
  for _, seed := range [][]byte{{}, {0}, {9}, {0xa}, {0xf}, {1, 2, 3, 4}} {
    f.Add(seed)
  }
  f.Fuzz(func(t *testing.T, in []byte) {
    enc := hex.EncodeToString(in)
    out, err := hex.DecodeString(enc)
    if err != nil {
      t.Fatalf("%v: decode: %v", in, err)
    }
    if !bytes.Equal(in, out) {
      t.Fatalf("%v: not equal after round trip: %v", in, out)
    }
  })
}
```

A fuzz test maintains a seed corpus, or a set of inputs which are run by default, and can seed input generation. Seed inputs may be registered by calling (*F).Add or by storing files in the directory testdata/fuzz/<Name> (where <Name> is the name of the fuzz test) within the package containing the fuzz test. Seed inputs are optional, but the fuzzing engine may find bugs more efficiently when provided with a set of small seed inputs with good code coverage. These seed inputs can also serve as regression tests for bugs identified through fuzzing.

​	模糊测试维护一个种子语料库或默认情况下运行的一组输入，并可以生成输入。种子输入可以通过调用 `(*F).Add` 或将文件存储在包含模糊测试的包中的 `testdata/fuzz/<Name>` (其中`<Name>`是模糊测试的名称)目录中注册。种子输入是可选的，但是当提供一组具有良好代码覆盖率的小种子输入时，模糊测试引擎可能会更有效地发现错误。这些种子输入还可以作为模糊测试识别的漏洞的回归测试。	

The function passed to (*F).Fuzz within the fuzz test is considered the fuzz target. A fuzz target must accept a *T parameter, followed by one or more parameters for random inputs. The types of arguments passed to (*F).Add must be identical to the types of these parameters. The fuzz target may signal that it's found a problem the same way tests do: by calling T.Fail (or any method that calls it like T.Error or T.Fatal) or by panicking.

​	在模糊测试中传递给 `(*F).Fuzz` 的函数被认为是模糊目标。一个模糊目标必须接受一个 `*T` 参数，后跟一个或多个随机输入的参数。传递给 `(*F).Add` 的参数类型必须与这些参数的类型相同。模糊目标可以通过调用 `T.Fail`(或调用它的任何方法，如 `T.Error` 或 `T.Fatal`)或引发 panic 的方式来指示它发现了问题，就像测试一样。

When fuzzing is enabled (by setting the -fuzz flag to a regular expression that matches a specific fuzz test), the fuzz target is called with arguments generated by repeatedly making random changes to the seed inputs. On supported platforms, 'go test' compiles the test executable with fuzzing coverage instrumentation. The fuzzing engine uses that instrumentation to find and cache inputs that expand coverage, increasing the likelihood of finding bugs. If the fuzz target fails for a given input, the fuzzing engine writes the inputs that caused the failure to a file in the directory testdata/fuzz/<Name> within the package directory. This file later serves as a seed input. If the file can't be written at that location (for example, because the directory is read-only), the fuzzing engine writes the file to the fuzz cache directory within the build cache instead.

​	当启用模糊测试(通过将 `-fuzz` 标志设置为与特定模糊测试匹配的正则表达式)，模糊目标将使用对种子输入进行随机更改生成的参数进行调用。在受支持的平台上，"go test" 使用模糊覆盖率工具编译测试可执行文件。模糊测试引擎使用该工具来查找和缓存扩展覆盖范围的输入，从而增加发现错误的可能性。如果模糊目标对于给定的输入失败，则模糊测试引擎将导致引发失败的输入写入包目录中的 `testdata/fuzz/<Name>` 目录中的文件中。此文件随后用作种子输入。如果无法在该位置写入文件(例如，因为目录为只读)，则模糊测试引擎将文件写入构建缓存中的模糊缓存目录中。

When fuzzing is disabled, the fuzz target is called with the seed inputs registered with F.Add and seed inputs from testdata/fuzz/<Name>. In this mode, the fuzz test acts much like a regular test, with subtests started with F.Fuzz instead of T.Run.

​	当禁用模糊测试时，`F.Add`注册的种子输入和`testdata/fuzz/<Name>`中的种子输入将用于调用模糊目标。在此模式下，模糊测试的行为类似于常规测试，使用`F.Fuzz`启动子测试而不是`T.Run`。

See https://go.dev/doc/fuzz for documentation about fuzzing.

​	有关模糊测试的文档，请参见[Go Fuzzing](../../UsingAndUnderstandingGo/fuzzing)。

## Skipping 

Tests or benchmarks may be skipped at run time with a call to the Skip method of *T or *B:

​	跳过测试或基准测试可以通过调用`*T`或`*B`的`Skip`方法来实现：

``` go 
func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    ...
}
```

The Skip method of *T can be used in a fuzz target if the input is invalid, but should not be considered a failing input. For example:

​	`*T`的`Skip`方法可以在模糊目标中使用，如果输入无效，但不应视为失败的输入。例如：

``` go 
func FuzzJSONMarshaling(f *testing.F) {
    f.Fuzz(func(t *testing.T, b []byte) {
        var v interface{}
        if err := json.Unmarshal(b, &v); err != nil {
            t.Skip()
        }
        if _, err := json.Marshal(v); err != nil {
            t.Errorf("Marshal: %v", err)
        }
    })
}
```

## 子测试和子基准测试 Subtests and Sub-benchmarks 

The Run methods of T and B allow defining subtests and sub-benchmarks, without having to define separate functions for each. This enables uses like table-driven benchmarks and creating hierarchical tests. It also provides a way to share common setup and tear-down code:

​	`*T`和`*B`的`Run`方法允许定义子测试和子基准测试，而无需为每个测试定义单独的函数。这使得可以使用表驱动的基准测试和创建分层测试。它还提供了一种共享通用设置和拆卸代码的方法：

``` go 
func TestFoo(t *testing.T) {
    // <setup code>
    t.Run("A=1", func(t *testing.T) { ... })
    t.Run("A=2", func(t *testing.T) { ... })
    t.Run("B=1", func(t *testing.T) { ... })
    // <tear-down code>
}
```

Each subtest and sub-benchmark has a unique name: the combination of the name of the top-level test and the sequence of names passed to Run, separated by slashes, with an optional trailing sequence number for disambiguation.

​	每个子测试和子基准测试都有一个唯一的名称：由顶层测试的名称和传递给Run的名称序列组成，用斜杠分隔，可选地带有一个末尾的序列号以进行消除歧义。

The argument to the -run, -bench, and -fuzz command-line flags is an unanchored regular expression that matches the test's name. For tests with multiple slash-separated elements, such as subtests, the argument is itself slash-separated, with expressions matching each name element in turn. Because it is unanchored, an empty expression matches any string. For example, using "matching" to mean "whose name contains":

​	`-run`、`-bench`和`-fuzz`命令行标志的参数是一个未锚定的正则表达式，用于匹配测试的名称。对于具有多个斜杠分隔元素(如子测试)的测试，参数本身是斜杠分隔的，其中表达式依次匹配每个名称元素。因为它是未锚定的，所以空表达式匹配任何字符串。例如，使用"matching"表示"whose name contains"：

> ​	未锚定的正则表达式是指没有明确指定正则表达式匹配字符串的起始位置和结束位置的表达式。在 Go 中，当使用 `-run` 命令行标志来运行测试用例时，如果正则表达式没有以 `^`(脱字符)开头和 `$`(美元符号)结尾，那么该表达式就是未锚定的。这意味着，表达式可以匹配任何位置出现的字符串。例如，如果运行 `go test -run TestFoo`，那么所有以 `TestFoo` 作为名字的测试用例都将被运行，包括 `TestFooBar`、`TestFooBaz` 等。

```
go test -run ''        # 运行所有测试。
go test -run Foo       # 运行与 "Foo"匹配的顶层测试，如 "TestFooBar"。
go test -run Foo/A=    # 对于匹配 "Foo"的顶级测试，运行匹配 "A="的子测试。
go test -run /A=1      # 对于所有顶级测试，运行匹配 "A=1"的子测试。
go test -fuzz FuzzFoo  # 对匹配 "FuzzFoo"的目标进行模糊处理。
```

The -run argument can also be used to run a specific value in the seed corpus, for debugging. For example:

​	-run参数也可以用于运行种子语料库中的特定值，以进行调试。例如：

```
go test -run=FuzzFoo/9ddb952d9814
```

The -fuzz and -run flags can both be set, in order to fuzz a target but skip the execution of all other tests.

​	可以同时设置`-fuzz`和`-run`标志，以模糊处理目标但跳过所有其他测试的执行。

Subtests can also be used to control parallelism. A parent test will only complete once all of its subtests complete. In this example, all tests are run in parallel with each other, and only with each other, regardless of other top-level tests that may be defined:

​	子测试还可以用于控制并行性。父测试只有在所有子测试完成后才会完成。在以下示例中，所有测试都会并行运行，仅与彼此并行运行，而不考虑其他可能已定义的顶级测试：

``` go 
func TestGroupedParallel(t *testing.T) {
    for _, tc := range tests {
        tc := tc // capture range variable
        t.Run(tc.Name, func(t *testing.T) {
            t.Parallel()
            ...
        })
    }
}
```

Run does not return until parallel subtests have completed, providing a way to clean up after a group of parallel tests:

​	`Run`不会返回，直到并行子测试完成，从而提供了一种在一组并行测试之后进行清理的方式：

``` go 
func TestTeardownParallel(t *testing.T) {
    // 在并行测试完成之前，此 Run 不会返回。
    t.Run("group", func(t *testing.T) {
        t.Run("Test1", parallelTest1)
        t.Run("Test2", parallelTest2)
        t.Run("Test3", parallelTest3)
    })
    // <tear-down code>
}
```

## Main 

It is sometimes necessary for a test or benchmark program to do extra setup or teardown before or after it executes. It is also sometimes necessary to control which code runs on the main thread. To support these and other cases, if a test file contains a function:

​	有时，测试或基准测试程序需要在执行前后进行额外的设置或拆卸。有时还需要控制哪些代码在主线程上运行。为支持这些和其他情况，如果测试文件包含一个函数：

``` go 
func TestMain(m *testing.M)
```

then the generated test will call TestMain(m) instead of running the tests or benchmarks directly. TestMain runs in the main goroutine and can do whatever setup and teardown is necessary around a call to m.Run. m.Run will return an exit code that may be passed to os.Exit. If TestMain returns, the test wrapper will pass the result of m.Run to os.Exit itself.

则生成的测试将调用`TestMain(m)`而不是直接运行测试或基准测试。`TestMain`在主goroutine中运行，并可以在调用`m.Run`周围执行任何必要的设置和拆卸操作。`m.Run`将返回可能传递给`os.Exit`的退出代码。如果`TestMain`返回，则测试包装器将`m.Run`的结果传递给`os.Exit`。

When TestMain is called, flag.Parse has not been run. If TestMain depends on command-line flags, including those of the testing package, it should call flag.Parse explicitly. Command line flags are always parsed by the time test or benchmark functions run.

​	在调用TestMain时，`flag.Parse`尚未运行。如果TestMain依赖于命令行标志，包括测试包的标志，它应该显式地调用`flag.Parse`。在测试或基准测试函数运行时，命令行标志总是被解析。

A simple implementation of TestMain is:

​	TestMain的一个简单实现是：

``` go 
func TestMain(m *testing.M) {
	// 如果TestMain使用标志，请在此处调用flag.Parse()
	os.Exit(m.Run())
}
```

TestMain is a low-level primitive and should not be necessary for casual testing needs, where ordinary test functions suffice.

​	TestMain 是一个底层的原语，对于普通测试需求，通常使用普通的测试函数就足够了。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func AllocsPerRun  <- go1.1

``` go 
func AllocsPerRun(runs int, f func()) (avg float64)
```

AllocsPerRun returns the average number of allocations during calls to f. Although the return value has type float64, it will always be an integral value.

​	AllocsPerRun函数返回在调用 `f` 时的平均分配次数。虽然返回值的类型为 float64，但它始终为整数值。

To compute the number of allocations, the function will first be run once as a warm-up. The average number of allocations over the specified number of runs will then be measured and returned.

​	为了计算分配次数，该函数将首先作为预热运行一次。然后将测量指定运行次数内的平均分配次数并返回。

AllocsPerRun sets GOMAXPROCS to 1 during its measurement and will restore it before returning.

​	AllocsPerRun函数在其测量期间将 `GOMAXPROCS` 设置为 1，并在返回前恢复它。

### func CoverMode  <- go1.8

``` go 
func CoverMode() string
```

CoverMode reports what the test coverage mode is set to. The values are "set", "count", or "atomic". The return value will be empty if test coverage is not enabled.

​	CoverMode函数报告测试覆盖模式设置为什么。可选值为 "set"、"count" 或 "atomic"。如果测试覆盖未启用，则返回值为空字符串。

### func Coverage  <- go1.4

``` go 
func Coverage() float64
```

Coverage reports the current code coverage as a fraction in the range [0, 1]. If coverage is not enabled, Coverage returns 0.

​	Coverage函数报告当前代码覆盖率，作为范围在 [0, 1] 内的分数。如果未启用覆盖，则 Coverage 返回 0。

When running a large set of sequential test cases, checking Coverage after each one can be useful for identifying which test cases exercise new code paths. It is not a replacement for the reports generated by 'go test -cover' and 'go tool cover'.

​	在运行大型顺序测试集时，每次检查 Coverage函数可以有助于识别哪些测试用例涉及新的代码路径。它不是"go test -cover"和 "go tool cover"生成的报告的替代品。

### func Init  <- go1.13

``` go 
func Init()
```

Init registers testing flags. These flags are automatically registered by the "go test" command before running test functions, so Init is only needed when calling functions such as Benchmark without using "go test".

​	Init 函数注册测试标志。在运行测试函数之前，"go test"命令会自动注册这些标志，因此只有在调用诸如 Benchmark 等函数而不使用"go test"时才需要调用 Init 函数。

Init has no effect if it was already called.

​	如果已经调用了 Init函数，则 Init 函数没有任何效果。

### func Main 

``` go 
func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, benchmarks []InternalBenchmark, examples []InternalExample)
```

Main is an internal function, part of the implementation of the "go test" command. It was exported because it is cross-package and predates "internal" packages. It is no longer used by "go test" but preserved, as much as possible, for other systems that simulate "go test" using Main, but Main sometimes cannot be updated as new functionality is added to the testing package. Systems simulating "go test" should be updated to use MainStart.

​	Main函数是 "go test" 命令实现的内部函数。它被导出是因为它是跨包的，并且早于 "internal" 包。它不再被 "go test" 使用，但为了其他系统可以模拟 "go test" 使用 Main，尽可能地保留了这个函数。但由于 testing 包添加了新功能，有时 Main 无法更新。模拟 "go test" 的系统应该更新为使用 MainStart。

### func RegisterCover  <- go1.2

``` go 
func RegisterCover(c Cover)
```

RegisterCover records the coverage data accumulators for the tests. NOTE: This function is internal to the testing infrastructure and may change. It is not covered (yet) by the Go 1 compatibility guidelines.

​	RegisterCover函数记录测试的覆盖数据累加器。注意：这个函数是 testing 基础设施的内部函数，可能会更改。它 (目前) 不符合 Go 1 兼容性指南。

### func RunBenchmarks 

``` go 
func RunBenchmarks(matchString func(pat, str string) (bool, error), benchmarks []InternalBenchmark)
```

RunBenchmarks is an internal function but exported because it is cross-package; it is part of the implementation of the "go test" command.

​	RunBenchmarks函数是一个内部函数，但导出是因为它是跨包的；它是 "go test" 命令的实现的一部分。

### func RunExamples 

``` go 
func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)
```

RunExamples is an internal function but exported because it is cross-package; it is part of the implementation of the "go test" command.

​	RunExamples函数是一个内部函数，但导出是因为它是跨包的；它是 "go test" 命令的实现的一部分。

### func RunTests 

``` go 
func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)
```

RunTests is an internal function but exported because it is cross-package; it is part of the implementation of the "go test" command.

​	RunTests函数是一个内部函数，但导出是因为它是跨包的；它是 "go test" 命令的实现的一部分。

### func Short 

``` go 
func Short() bool
```

Short reports whether the -test.short flag is set.

​	Short 函数报告是否设置了 `-test.short` 标志。

### func Testing <-go1.21.0

```go
func Testing() bool
```

Testing reports whether the current code is being run in a test. This will report true in programs created by "go test", false in programs created by "go build".

​	`Testing` 报告当前代码是否正在测试中运行。在由 `go test` 创建的程序中将返回 `true`，在由 `go build` 创建的程序中将返回 `false`。

### func Verbose  <- go1.1

``` go 
func Verbose() bool
```

Verbose reports whether the -test.v flag is set.

​	Verbose函数报告是否设置了`-test.v`标志。

## 类型

### type B 

``` go 
type B struct {
	N int
	// contains filtered or unexported fields
}
```

B is a type passed to Benchmark functions to manage benchmark timing and to specify the number of iterations to run.

​	B是传递给基准测试函数的类型，用于管理基准测试的时间和指定要运行的迭代次数。

A benchmark ends when its Benchmark function returns or calls any of the methods FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods must be called only from the goroutine running the Benchmark function. The other reporting methods, such as the variations of Log and Error, may be called simultaneously from multiple goroutines.

​	当Benchmark函数返回或调用FailNow、Fatal、Fatalf、SkipNow、Skip或Skipf方法时，基准测试结束。这些方法只能从运行Benchmark函数的goroutine中调用。其他报告方法，例如Log和Error的变体，可以同时从多个goroutine调用。

Like in tests, benchmark logs are accumulated during execution and dumped to standard output when done. Unlike in tests, benchmark logs are always printed, so as not to hide output whose existence may be affecting benchmark results.

​	与测试类似，基准测试日志在执行期间累积，并在完成时转储到标准输出。与测试不同，基准测试日志始终打印，以不隐藏可能影响基准测试结果的输出。

#### (*B) Chdir <- 1.24.0

```go
func (c *B) Chdir(dir string)
```

Chdir calls os.Chdir(dir) and uses Cleanup to restore the current working directory to its original value after the test. On Unix, it also sets PWD environment variable for the duration of the test.

Because Chdir affects the whole process, it cannot be used in parallel tests or tests with parallel ancestors.

#### (*B) Cleanup  <- go1.14

``` go 
func (c *B) Cleanup(f func())
```

Cleanup registers a function to be called when the test (or subtest) and all its subtests complete. Cleanup functions will be called in last added, first called order.

​	Cleanup方法注册一个函数，以在测试(或子测试)及其所有子测试完成时调用。清理函数将按最后添加的先调用的顺序调用。

#### (*B) Context <- 1.24.0

```go
func (c *B) Context() context.Context
```

Context returns a context that is canceled just before Cleanup-registered functions are called.

Cleanup functions can wait for any resources that shut down on Context.Done before the test or benchmark completes.

#### (*B) Elapsed  <- go1.20

``` go 
func (b *B) Elapsed() time.Duration
```

Elapsed returns the measured elapsed time of the benchmark. The duration reported by Elapsed matches the one measured by StartTimer, StopTimer, and ResetTimer.

​	Elapsed方法返回基准测试的测量经过时间。Elapsed报告的持续时间与StartTimer方法、StopTimer方法和ResetTimer方法测量的时间相匹配。

#### (*B) Error 

``` go 
func (c *B) Error(args ...any)
```

Error is equivalent to Log followed by Fail.

​	Error方法等效于Log方法后跟Fail方法。

#### (*B) Errorf 

``` go 
func (c *B) Errorf(format string, args ...any)
```

Errorf is equivalent to Logf followed by Fail.

​	Errorf方法等效于Logf方法后跟Fail方法。

#### (*B) Fail 

``` go 
func (c *B) Fail()
```

Fail marks the function as having failed but continues execution.

​	Fail方法标记函数失败，但继续执行。

#### (*B) FailNow 

``` go 
func (c *B) FailNow()
```

FailNow marks the function as having failed and stops its execution by calling runtime.Goexit (which then runs all deferred calls in the current goroutine). Execution will continue at the next test or benchmark. FailNow must be called from the goroutine running the test or benchmark function, not from other goroutines created during the test. Calling FailNow does not stop those other goroutines.

​	FailNow方法标记函数失败并通过调用runtime.Goexit(然后在当前goroutine中运行所有延迟调用)停止其执行。执行将继续在下一个测试或基准测试中。FailNow方法必须从运行测试或基准测试函数的goroutine中调用，而不是从在测试期间创建的其他goroutine中调用。调用FailNow方法不会停止这些其他goroutine。

#### (*B) Failed 

``` go 
func (c *B) Failed() bool
```

Failed reports whether the function has failed.

​	Failed方法报告函数是否已失败。

#### (*B) Fatal 

``` go 
func (c *B) Fatal(args ...any)
```

Fatal is equivalent to Log followed by FailNow.

​	Fatal方法等效于 Log方法后跟 FailNow方法。

#### (*B) Fatalf 

``` go 
func (c *B) Fatalf(format string, args ...any)
```

Fatalf is equivalent to Logf followed by FailNow.

​	Fatalf方法等效于 Logf方法后跟 FailNow方法。

#### (*B) Helper  <- go1.9

``` go 
func (c *B) Helper()
```

Helper marks the calling function as a test helper function. When printing file and line information, that function will be skipped. Helper may be called simultaneously from multiple goroutines.

​	Helper方法将调用函数标记为测试辅助函数。在打印文件和行信息时，该函数将被跳过。Helper方法可以同时从多个 goroutine 中调用。

#### (*B) Log 

``` go 
func (c *B) Log(args ...any)
```

Log formats its arguments using default formatting, analogous to Println, and records the text in the error log. For tests, the text will be printed only if the test fails or the -test.v flag is set. For benchmarks, the text is always printed to avoid having performance depend on the value of the -test.v flag.

​	Log方法使用默认格式对其参数进行格式化，类似于 Println，并将文本记录在错误日志中。对于测试，只有在测试失败或设置了 `-test.v` 标志时，才会打印该文本。对于基准测试，始终会打印文本，以避免性能依赖于 `-test.v` 标志的值。

#### (*B) Logf 

``` go 
func (c *B) Logf(format string, args ...any)
```

Logf formats its arguments according to the format, analogous to Printf, and records the text in the error log. A final newline is added if not provided. For tests, the text will be printed only if the test fails or the -test.v flag is set. For benchmarks, the text is always printed to avoid having performance depend on the value of the -test.v flag.

​	Logf方法根据格式对其参数进行格式化，类似于 Printf，并将文本记录在错误日志中。如果没有提供最终换行符，则会添加。对于测试，只有在测试失败或设置了 `-test.v` 标志时，才会打印该文本。对于基准测试，始终会打印文本，以避免性能依赖于 `-test.v` 标志的值。

#### (*B) Loop <- 1.24.0

```go
func (b *B) Loop() bool
```

Loop returns true as long as the benchmark should continue running.

A typical benchmark is structured like:

```go
func Benchmark(b *testing.B) {
	... setup ...
	for b.Loop() {
		... code to measure ...
	}
	... cleanup ...
}
```

Loop resets the benchmark timer the first time it is called in a benchmark, so any setup performed prior to starting the benchmark loop does not count toward the benchmark measurement. Likewise, when it returns false, it stops the timer so cleanup code is not measured.

The compiler never optimizes away calls to functions within the body of a "for b.Loop() { ... }" loop. This prevents surprises that can otherwise occur if the compiler determines that the result of a benchmarked function is unused. The loop must be written in exactly this form, and this only applies to calls syntactically between the curly braces of the loop. Optimizations are performed as usual in any functions called by the loop.

After Loop returns false, b.N contains the total number of iterations that ran, so the benchmark may use b.N to compute other average metrics.

Prior to the introduction of Loop, benchmarks were expected to contain an explicit loop from 0 to b.N. Benchmarks should either use Loop or contain a loop to b.N, but not both. Loop offers more automatic management of the benchmark timer, and runs each benchmark function only once per measurement, whereas b.N-based benchmarks must run the benchmark function (and any associated setup and cleanup) several times.

##### Loop Example

```go
package main

import (
	"math/rand/v2"
	"testing"
)

// ExBenchmark shows how to use b.Loop in a benchmark.
//
// (If this were a real benchmark, not an example, this would be named
// BenchmarkSomething.)
func ExBenchmark(b *testing.B) {
	// Generate a large random slice to use as an input.
	// Since this is done before the first call to b.Loop(),
	// it doesn't count toward the benchmark time.
	input := make([]int, 128<<10)
	for i := range input {
		input[i] = rand.Int()
	}

	// Perform the benchmark.
	for b.Loop() {
		// Normally, the compiler would be allowed to optimize away the call
		// to sum because it has no side effects and the result isn't used.
		// However, inside a b.Loop loop, the compiler ensures function calls
		// aren't optimized away.
		sum(input)
	}

	// Outside the loop, the timer is stopped, so we could perform
	// cleanup if necessary without affecting the result.
}

func sum(data []int) int {
	total := 0
	for _, value := range data {
		total += value
	}
	return total
}

func main() {
	testing.Benchmark(ExBenchmark)
}

```



#### (*B) Name  <- go1.8

``` go 
func (c *B) Name() string
```

Name returns the name of the running (sub-) test or benchmark.

​	Name方法返回正在运行的(子)测试或基准测试的名称。

The name will include the name of the test along with the names of any nested sub-tests. If two sibling sub-tests have the same name, Name will append a suffix to guarantee the returned name is unique.

​	名称将包括测试的名称以及任何嵌套的子测试的名称。如果两个同级的子测试具有相同的名称，则 Name 将附加后缀以保证返回的名称是唯一的。

#### (*B) ReportAllocs  <- go1.1

``` go 
func (b *B) ReportAllocs()
```

ReportAllocs enables malloc statistics for this benchmark. It is equivalent to setting -test.benchmem, but it only affects the benchmark function that calls ReportAllocs.

​	ReportAllocs方法为此基准测试启用 malloc 统计信息。它相当于设置 `-test.benchmem`，但仅影响调用 ReportAllocs方法的基准测试函数。

#### (*B) ReportMetric  <- go1.13

``` go 
func (b *B) ReportMetric(n float64, unit string)
```

ReportMetric adds "n unit" to the reported benchmark results. If the metric is per-iteration, the caller should divide by b.N, and by convention units should end in "/op". ReportMetric overrides any previously reported value for the same unit. ReportMetric panics if unit is the empty string or if unit contains any whitespace. If unit is a unit normally reported by the benchmark framework itself (such as "allocs/op"), ReportMetric will override that metric. Setting "ns/op" to 0 will suppress that built-in metric.

​	ReportMetric方法将"n unit"添加到报告的基准测试结果中。如果该指标是每次迭代的，则调用者应该除以 b.N，并且按照惯例，单位应该以"/op"结尾。ReportMetric 将覆盖先前报告的相同单位的任何值。如果单位为空字符串或单位包含任何空格，则 ReportMetric方法将 panic。如果单位通常由基准框架本身报告(如"allocs/op")，则 ReportMetric方法将覆盖该指标。将 "ns/op" 设置为 0 将禁止该内置度量。

##### ReportMetric Example

``` go 
package main

import (
	"sort"
	"testing"
)

func main() {
	// 这里报告了与特定算法相关的自定义基准测试度量标准(在本例中为排序)。
	testing.Benchmark(func(b *testing.B) {
		var compares int64
		for i := 0; i < b.N; i++ {
			s := []int{5, 4, 3, 2, 1}
			sort.Slice(s, func(i, j int) bool {
				compares++
				return s[i] < s[j]
			})
		}
		// 此度量标准是每个操作的，因此需要除以b.N，并将其报告为"/op"单位。
		b.ReportMetric(float64(compares)/float64(b.N), "compares/op")
		// 此度量标准是每个时间段的，因此需要除以b.Elapsed，并将其报告为"/ns"单位。
		b.ReportMetric(float64(compares)/float64(b.Elapsed().Nanoseconds()), "compares/ns")
	})
}

```

##### ReportMetric Example (Parallel) 

``` go 
package main

import (
	"sort"
	"sync/atomic"
	"testing"
)

func main() {
	// 这个函数报告了一个与特定算法相关的自定义基准度量(在本例中为排序)并且是并行执行的。
	testing.Benchmark(func(b *testing.B) {
		var compares atomic.Int64
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := []int{5, 4, 3, 2, 1}
				sort.Slice(s, func(i, j int) bool {
					// 因为RunParallel函数在并行中多次运行该函数，
                    // 我们必须以原子方式递增计数器以避免竞争写入。
					compares.Add(1)
					return s[i] < s[j]
				})
			}
		})

		// 注意：在所有并行调用完成后，每个度量值只需报告一次。

		// 此度量值是每个操作，所以需要除以 b.N，并将其报告为 "/op" 单位。
		b.ReportMetric(float64(compares.Load())/float64(b.N), "compares/op")
		// 此度量值是每个时间，所以需要除以 b.Elapsed 并将其报告为 "/ns" 单位。
		b.ReportMetric(float64(compares.Load())/float64(b.Elapsed().Nanoseconds()), "compares/ns")
	})
}

```



#### (*B) ResetTimer 

``` go 
func (b *B) ResetTimer()
```

ResetTimer zeroes the elapsed benchmark time and memory allocation counters and deletes user-reported metrics. It does not affect whether the timer is running.

​	ResetTimer方法将经过的基准测试时间和内存分配计数器归零，并删除用户报告的度量标准。它不会影响定时器是否在运行。

#### (*B) Run  <- go1.7

``` go 
func (b *B) Run(name string, f func(b *B)) bool
```

Run benchmarks f as a subbenchmark with the given name. It reports whether there were any failures.

​	Run方法将f作为具有给定名称的子基准测试运行。它报告是否有任何失败。

A subbenchmark is like any other benchmark. A benchmark that calls Run at least once will not be measured itself and will be called once with N=1.

​	子基准测试与任何其他基准测试相似。调用Run方法至少一次的基准测试本身不会被测量，并且将以N=1的方式调用一次。

#### (*B) RunParallel  <- go1.3

``` go 
func (b *B) RunParallel(body func(*PB))
```

RunParallel runs a benchmark in parallel. It creates multiple goroutines and distributes b.N iterations among them. The number of goroutines defaults to GOMAXPROCS. To increase parallelism for non-CPU-bound benchmarks, call SetParallelism before RunParallel. RunParallel is usually used with the go test -cpu flag.

​	RunParallel方法在并行中运行基准测试。它创建多个goroutine并将b.N次迭代分布在它们之间。 goroutine的数量默认为GOMAXPROCS。要增加非CPU绑定的基准测试的并行性，请在RunParallel之前调用SetParallelism。 RunParallel方法通常与go test -cpu标志一起使用。

The body function will be run in each goroutine. It should set up any goroutine-local state and then iterate until pb.Next returns false. It should not use the StartTimer, StopTimer, or ResetTimer functions, because they have global effect. It should also not call Run.

​	body函数将在每个goroutine中运行。它应该设置任何goroutine-local状态，然后迭代直到pb.Next返回false。它不应该使用StartTimer方法、StopTimer方法或ResetTimer方法，因为它们具有全局效果。它也不应该调用Run方法。

RunParallel reports ns/op values as wall time for the benchmark as a whole, not the sum of wall time or CPU time over each parallel goroutine.

​	RunParallel方法将ns/op值报告为整个基准测试的挂起时间，而不是每个并行goroutine的挂起时间或CPU时间总和。



##### RunParallel Example
``` go 
package main

import (
	"bytes"
	"testing"
	"text/template"
)

func main() {
    // Parallel benchmark for text/template.Template.Execute on a single object.
	// 这是针对text/template.Template.Execute在单个对象上进行并行基准测试的代码。
	testing.Benchmark(func(b *testing.B) {
		templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
        // RunParallel will create GOMAXPROCS goroutines
		// and distribute work among them.
		// RunParallel将创建GOMAXPROCS个goroutine并在它们之间分配工作。
		b.RunParallel(func(pb *testing.PB) {
            // Each goroutine has its own bytes.Buffer.
			// 每个goroutine都有自己的bytes.Buffer。
			var buf bytes.Buffer
			for pb.Next() {
                // The loop body is executed b.N times total across all goroutines.
				// 循环体在所有goroutine中总共执行b.N次。
				buf.Reset()
				templ.Execute(&buf, "World")
			}
		})
	})
}

```

#### (*B) SetBytes 

``` go 
func (b *B) SetBytes(n int64)
```

SetBytes records the number of bytes processed in a single operation. If this is called, the benchmark will report ns/op and MB/s.

​	SetBytes方法记录在单个操作中处理的字节数。如果调用了此函数，则基准测试将报告ns/op和MB/s。

#### (*B) SetParallelism  <- go1.3

``` go 
func (b *B) SetParallelism(p int)
```

SetParallelism sets the number of goroutines used by RunParallel to p*GOMAXPROCS. There is usually no need to call SetParallelism for CPU-bound benchmarks. If p is less than 1, this call will have no effect.

​	SetParallelism方法设置RunParallel方法要使用的goroutine数为`p * GOMAXPROCS`。对于CPU绑定的基准测试，通常不需要调用SetParallelism方法。如果p小于1，则此调用将无效。

#### (*B) Setenv  <- go1.17

``` go 
func (c *B) Setenv(key, value string)
```

Setenv calls os.Setenv(key, value) and uses Cleanup to restore the environment variable to its original value after the test.

​	Setenv方法调用os.Setenv(key, value)，并使用Cleanup方法在测试后将环境变量恢复到其原始值。

Because Setenv affects the whole process, it cannot be used in parallel tests or tests with parallel ancestors.

​	因为Setenv方法会影响整个进程，所以它不能在并行测试或具有并行祖先的测试中使用。

#### (*B) Skip  <- go1.1

``` go 
func (c *B) Skip(args ...any)
```

Skip is equivalent to Log followed by SkipNow.

​	Skip方法等同于Log方法后跟SkipNow方法。

#### (*B) SkipNow  <- go1.1

``` go 
func (c *B) SkipNow()
```

SkipNow marks the test as having been skipped and stops its execution by calling runtime.Goexit. If a test fails (see Error, Errorf, Fail) and is then skipped, it is still considered to have failed. Execution will continue at the next test or benchmark. See also FailNow. SkipNow must be called from the goroutine running the test, not from other goroutines created during the test. Calling SkipNow does not stop those other goroutines.

​	SkipNow方法将测试标记为已跳过，并通过调用runtime.Goexit停止其执行。如果测试失败(请参见Error、Errorf、Fail)，然后跳过，它仍然被认为是已失败的。执行将在下一个测试或基准测试上继续。另请参见FailNow。SkipNow必须从运行测试的goroutine而不是从测试期间创建的其他goroutine调用。调用SkipNow方法不会停止这些其他goroutine。

#### (*B) Skipf  <- go1.1

``` go 
func (c *B) Skipf(format string, args ...any)
```

Skipf is equivalent to Logf followed by SkipNow.

​	Skipf方法等同于Logf方法后跟SkipNow方法。

#### (*B) Skipped  <- go1.1

``` go 
func (c *B) Skipped() bool
```

Skipped reports whether the test was skipped.

​	Skipped方法报告测试是否被跳过。

#### (*B) StartTimer 

``` go 
func (b *B) StartTimer()
```

StartTimer starts timing a test. This function is called automatically before a benchmark starts, but it can also be used to resume timing after a call to StopTimer.

​	StartTimer方法开始计时测试。该函数会在基准测试开始前自动调用，但也可以在需要测量但不想计入时间的复杂初始化操作后手动调用以恢复计时。

#### (*B) StopTimer 

``` go 
func (b *B) StopTimer()
```

StopTimer stops timing a test. This can be used to pause the timer while performing complex initialization that you don't want to measure.

​	StopTimer方法停止计时测试。可以在进行不需要测量时间的复杂初始化操作时使用它以暂停计时。

#### (*B) TempDir  <- go1.15

``` go 
func (c *B) TempDir() string
```

TempDir returns a temporary directory for the test to use. The directory is automatically removed by Cleanup when the test and all its subtests complete. Each subsequent call to t.TempDir returns a unique directory; if the directory creation fails, TempDir terminates the test by calling Fatal.

​	TempDir方法返回测试使用的临时目录。该目录会在测试及其所有子测试完成时由 Cleanup方法自动删除。每次调用 t.TempDir方法都会返回一个唯一的目录；如果目录创建失败，TempDir方法会通过调用 Fatal 方法终止测试。

### type BenchmarkResult 

``` go 
type BenchmarkResult struct {
	N         int           // 迭代次数。 The number of iterations.
	T         time.Duration // 所花费的总时间。 The total time taken.
	Bytes     int64         // 单次迭代处理的字节数。 Bytes processed in one iteration.
	MemAllocs uint64        // 内存分配总次数。  The total number of memory allocations.
	MemBytes  uint64        // 分配的总字节数。 The total number of bytes allocated.

    // Extra records additional metrics reported by ReportMetric.
	// Extra 记录 ReportMetric 报告的其他度量标准。
	Extra map[string]float64
}
```

BenchmarkResult contains the results of a benchmark run.

​	BenchmarkResult结构体包含基准测试运行的结果。

#### func Benchmark 

``` go 
func Benchmark(f func(b *B)) BenchmarkResult
```

Benchmark benchmarks a single function. It is useful for creating custom benchmarks that do not use the "go test" command.

​	Benchmark函数对单个函数进行基准测试。它适用于创建不使用 "go test" 命令的自定义基准测试。

If f depends on testing flags, then Init must be used to register those flags before calling Benchmark and before calling flag.Parse.

​	如果 f 依赖于测试标志，则必须在调用 Benchmark 之前使用 Init方法注册这些标志，并在调用 flag.Parse 之前进行。

If f calls Run, the result will be an estimate of running all its subbenchmarks that don't call Run in sequence in a single benchmark.

​	如果 f 调用 Run方法，则结果将是估计运行所有不调用 Run方法的子基准测试的结果。

#### (BenchmarkResult) AllocedBytesPerOp  <- go1.1

``` go 
func (r BenchmarkResult) AllocedBytesPerOp() int64
```

AllocedBytesPerOp returns the "B/op" metric, which is calculated as r.MemBytes / r.N.

​	AllocedBytesPerOp方法返回 "B/op" 指标，它被计算为 r.MemBytes / r.N。

#### (BenchmarkResult) AllocsPerOp  <- go1.1

``` go 
func (r BenchmarkResult) AllocsPerOp() int64
```

AllocsPerOp returns the "allocs/op" metric, which is calculated as r.MemAllocs / r.N.

​	AllocsPerOp方法返回 "allocs/op" 指标，它被计算为 r.MemAllocs / r.N。

#### (BenchmarkResult) MemString  <- go1.1

``` go 
func (r BenchmarkResult) MemString() string
```

MemString returns r.AllocedBytesPerOp and r.AllocsPerOp in the same format as 'go test'.

​	MemString方法以与 'go test' 相同的格式返回 r.AllocedBytesPerOp 和 r.AllocsPerOp。

#### (BenchmarkResult) NsPerOp 

``` go 
func (r BenchmarkResult) NsPerOp() int64
```

NsPerOp returns the "ns/op" metric.

​	NsPerOp方法返回"ns/op"指标。

#### (BenchmarkResult) String 

``` go 
func (r BenchmarkResult) String() string
```

String returns a summary of the benchmark results. It follows the benchmark result line format from https://golang.org/design/14313-benchmark-format, not including the benchmark name. Extra metrics override built-in metrics of the same name. String does not include allocs/op or B/op, since those are reported by MemString.

​	String方法返回基准测试结果的摘要。它遵循 [https://golang.org/design/14313-benchmark-format](https://golang.org/design/14313-benchmark-format) 的基准测试结果行格式，但不包括基准测试名称。同名的额外指标会覆盖内置指标。String方法不包括 allocs/op 或 B/op，因为它们由 MemString方法报告。

### type Cover  <- go1.2

``` go 
type Cover struct {
	Mode            string
	Counters        map[string][]uint32
	Blocks          map[string][]CoverBlock
	CoveredPackages string
}
```

Cover records information about test coverage checking. NOTE: This struct is internal to the testing infrastructure and may change. It is not covered (yet) by the Go 1 compatibility guidelines.

​	Cover结构体记录有关测试覆盖检查的信息。注意：此结构体对于测试基础设施是内部的，可能会更改。它尚未(但可能会)受到 Go 1 兼容性指南的影响。

### type CoverBlock  <- go1.2

``` go 
type CoverBlock struct {
	Line0 uint32 // 块开头所在的行号
	Col0  uint16 // 块开头所在的列号
	Line1 uint32 // 块结尾所在的行号
	Col1  uint16 // 块结尾所在的列号
	Stmts uint16 // 此块中包含的语句数
}
```

CoverBlock records the coverage data for a single basic block. The fields are 1-indexed, as in an editor: The opening line of the file is number 1, for example. Columns are measured in bytes. NOTE: This struct is internal to the testing infrastructure and may change. It is not covered (yet) by the Go 1 compatibility guidelines.

​	CoverBlock结构体记录单个基本块的覆盖数据。这些字段是从编辑器中的 1 开始计数的：例如，文件的开头行号是 1。列以字节为单位测量。注意：此结构体对于测试基础设施是内部的，可能会更改。它尚未(但可能会)受到 Go 1 兼容性指南的影响。

### type F  <- go1.18

``` go 
type F struct {
	// contains filtered or unexported fields
}
```

F is a type passed to fuzz tests.

​	F结构体是传递给模糊测试的类型。

Fuzz tests run generated inputs against a provided fuzz target, which can find and report potential bugs in the code being tested.

​	模糊测试将生成的输入针对提供的模糊目标运行，可以找到并报告被测试代码中的潜在错误。

A fuzz test runs the seed corpus by default, which includes entries provided by `(*F).Add` and entries in the `testdata/fuzz/<FuzzTestName>` directory. After any necessary setup and calls to `(*F).Add,` the fuzz test must then call `(*F).Fuzz` to provide the fuzz target. See the testing package documentation for an example, and see the `F.Fuzz` and `F.Add` method documentation for details.

​	默认情况下，模糊测试运行种子语料库，其中包括 `(*F).Add` 中提供的条目和 `testdata/fuzz/<FuzzTestName>` 目录中的条目。在任何必要的设置和调用 `(*F).Add` 后，模糊测试必须调用 `(*F).Fuzz` 提供模糊目标。有关示例，请参见 testing 包文档，并参见 F.Fuzz 和 F.Add 方法文档获取详细信息。

`*F` methods can only be called before `(*F).Fuzz`. Once the test is executing the fuzz target, only `(*T)` methods can be used. The only *F methods that are allowed in the `(*F).Fuzz` function are `(*F).Failed` and `(*F).Name`.

​	`*F` 方法只能在 `(*F).Fuzz` 之前调用。一旦测试执行模糊目标，只能使用 `(*T)` 方法。`(*F).Failed` 和 `(*F).Name` 是 `(*F).Fuzz` 函数中允许的仅有的 `*F` 方法。

#### (*F) Add  <- go1.18

``` go 
func (f *F) Add(args ...any)
```

Add will add the arguments to the seed corpus for the fuzz test. This will be a no-op if called after or within the fuzz target, and args must match the arguments for the fuzz target.

​	`Add`方法将参数添加到 fuzz 测试的种子语料库中。如果在模糊目标之后或其中调用，将不起作用，并且 `args` 必须与模糊目标的参数匹配。

#### (*F) Chdir <- 1.24.0

```go
func (c *F) Chdir(dir string)
```

Chdir calls os.Chdir(dir) and uses Cleanup to restore the current working directory to its original value after the test. On Unix, it also sets PWD environment variable for the duration of the test.

Because Chdir affects the whole process, it cannot be used in parallel tests or tests with parallel ancestors.

#### (*F) Cleanup  <- go1.18

``` go 
func (c *F) Cleanup(f func())
```

Cleanup registers a function to be called when the test (or subtest) and all its subtests complete. Cleanup functions will be called in last added, first called order.

​	`Cleanup`方法注册一个在测试(或子测试)及其所有子测试完成时调用的函数。Cleanup方法将按照最后添加的先调用的顺序调用。

#### (*F) Context <- 1.24.0

```go
func (c *F) Context() context.Context
```

Context returns a context that is canceled just before Cleanup-registered functions are called.

Cleanup functions can wait for any resources that shut down on Context.Done before the test or benchmark completes.

#### (*F) Error  <- go1.18

``` go 
func (c *F) Error(args ...any)
```

Error is equivalent to Log followed by Fail.

​	Error 方法等同于 Log 方法后跟 Fail 方法。

#### (*F) Errorf  <- go1.18

``` go 
func (c *F) Errorf(format string, args ...any)
```

Errorf is equivalent to Logf followed by Fail.

​	Errorf方法等同于 Logf 方法后跟 Fail 方法。

#### (*F) Fail  <- go1.18

``` go 
func (f *F) Fail()
```

Fail marks the function as having failed but continues execution.

​	Fail 方法标记函数失败但继续执行。

#### (*F) FailNow  <- go1.18

``` go 
func (c *F) FailNow()
```

FailNow marks the function as having failed and stops its execution by calling runtime.Goexit (which then runs all deferred calls in the current goroutine). Execution will continue at the next test or benchmark. FailNow must be called from the goroutine running the test or benchmark function, not from other goroutines created during the test. Calling FailNow does not stop those other goroutines.

​	FailNow方法标记函数失败并通过调用 runtime.Goexit 停止其执行(然后运行当前 goroutine 中的所有延迟调用)。执行将在下一个测试或基准中继续。FailNow方法必须从运行测试或基准函数的 goroutine 调用，而不是从测试期间创建的其他 goroutine 中调用。调用 FailNow方法不会停止这些其他 goroutine。

#### (*F) Failed  <- go1.18

``` go 
func (c *F) Failed() bool
```

Failed reports whether the function has failed.

​	Failed 方法报告函数是否已失败。

#### (*F) Fatal  <- go1.18

``` go 
func (c *F) Fatal(args ...any)
```

Fatal is equivalent to Log followed by FailNow.

​	Fatal 方法等同于 Log 方法后跟 FailNow 方法。

#### (*F) Fatalf  <- go1.18

``` go 
func (c *F) Fatalf(format string, args ...any)
```

Fatalf is equivalent to Logf followed by FailNow.

​	Fatalf 方法等同于 Logf 方法后跟 FailNow 方法。

#### (*F) Fuzz  <- go1.18

``` go 
func (f *F) Fuzz(ff any)
```

Fuzz runs the fuzz function, ff, for fuzz testing. If ff fails for a set of arguments, those arguments will be added to the seed corpus.

​	Fuzz 方法运行fuzz函数，ff，进行模糊测试。如果 ff 对一组参数失败，则这些参数将被添加到种子语料库中。

ff must be a function with no return value whose first argument is *T and whose remaining arguments are the types to be fuzzed. For example:

​	ff 必须是一个没有返回值的函数，其第一个参数为 `*T`，其余参数为要进行 fuzz 的类型。例如：

```
f.Fuzz(func(t *testing.T, b []byte, i int) { ... })
```

The following types are allowed: []byte, string, bool, byte, rune, float32, float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64. More types may be supported in the future.

​	以下类型是允许的：[]byte、string、bool、byte、rune、float32、float64、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64。更多类型可能会在未来支持。

ff must not call any `*F` methods, e.g. `(*F).Log`, `(*F).Error`, `(*F).Skip`. Use the corresponding `*T` method instead. The only `*F` methods that are allowed in the `(*F).Fuzz` function are `(*F).Failed` and `(*F).Name`.

​	ff 不能调用任何 `*F` 方法，例如 `(*F).Log`、`(*F).Error`、`(*F).Skip`。应改用相应的 `*T` 方法。`(*F).Failed` 和 `(*F).Name` 是唯一允许在 `(*F).Fuzz` 函数中使用的 `*F` 方法

This function should be fast and deterministic, and its behavior should not depend on shared state. No mutatable input arguments, or pointers to them, should be retained between executions of the fuzz function, as the memory backing them may be mutated during a subsequent invocation. ff must not modify the underlying data of the arguments provided by the fuzzing engine.

​	该函数应该快速且确定性，其行为不应该依赖于共享状态。不应在执行模糊函数之间保留可变输入参数或指向它们的指针，因为支持它们的内存可能在随后的调用期间被修改。ff不能修改由fuzz引擎提供的参数的底层数据。

When fuzzing, F.Fuzz does not return until a problem is found, time runs out (set with -fuzztime), or the test process is interrupted by a signal. F.Fuzz should be called exactly once, unless F.Skip or F.Fail is called beforehand.

​	在进行模糊测试时，F.Fuzz方法不会返回，直到发现问题，时间耗尽(使用`-fuzztime`设置)，或测试过程被信号中断。F.Fuzz应该仅被调用一次，除非在之前调用了F.Skip或F.Fail。

#### (*F) Helper  <- go1.18

``` go 
func (f *F) Helper()
```

Helper marks the calling function as a test helper function. When printing file and line information, that function will be skipped. Helper may be called simultaneously from multiple goroutines.

​	Helper方法标记调用函数为测试辅助函数。在打印文件和行信息时，该函数将被跳过。Helper方法可以同时从多个goroutine调用。

#### (*F) Log  <- go1.18

``` go 
func (c *F) Log(args ...any)
```

Log formats its arguments using default formatting, analogous to Println, and records the text in the error log. For tests, the text will be printed only if the test fails or the -test.v flag is set. For benchmarks, the text is always printed to avoid having performance depend on the value of the -test.v flag.

​	`Log`方法使用默认格式对其参数进行格式化，类似于Println，并将文本记录在错误日志中。对于测试，仅当测试失败或设置了`-test.v`标志时才会打印文本。对于基准测试，为避免性能依赖于`-test.v`标志的值，始终打印文本。

#### (*F) Logf  <- go1.18

``` go 
func (c *F) Logf(format string, args ...any)
```

Logf formats its arguments according to the format, analogous to Printf, and records the text in the error log. A final newline is added if not provided. For tests, the text will be printed only if the test fails or the -test.v flag is set. For benchmarks, the text is always printed to avoid having performance depend on the value of the -test.v flag.

​	`Logf`方法根据格式对其参数进行格式化，类似于Printf，并将文本记录在错误日志中。如果没有提供最后一个换行符，则会添加一个。对于测试，仅当测试失败或设置了`-test.v`标志时才会打印文本。对于基准测试，为避免性能依赖于`-test.v`标志的值，始终打印文本。

#### (*F) Name  <- go1.18

``` go 
func (c *F) Name() string
```

Name returns the name of the running (sub-) test or benchmark.

​	`Name`方法返回正在运行的(子)测试或基准测试的名称。

The name will include the name of the test along with the names of any nested sub-tests. If two sibling sub-tests have the same name, Name will append a suffix to guarantee the returned name is unique.

​	名称将包括测试的名称以及任何嵌套的子测试的名称。如果两个同级别的子测试名称相同，则Name方法将附加后缀以确保返回的名称是唯一的。

#### (*F) Setenv  <- go1.18

``` go 
func (c *F) Setenv(key, value string)
```

Setenv calls os.Setenv(key, value) and uses Cleanup to restore the environment variable to its original value after the test.

​	Setenv方法调用os.Setenv(key, value)，并使用Cleanup方法将环境变量还原为其原始值。

Because Setenv affects the whole process, it cannot be used in parallel tests or tests with parallel ancestors.

​	由于Setenv方法影响整个进程，因此无法在并行测试或具有并行祖先的测试中使用。

#### (*F) Skip  <- go1.18

``` go 
func (c *F) Skip(args ...any)
```

Skip is equivalent to Log followed by SkipNow.

​	Skip方法等同于Log方法后跟SkipNow方法。

#### (*F) SkipNow  <- go1.18

``` go 
func (c *F) SkipNow()
```

SkipNow marks the test as having been skipped and stops its execution by calling runtime.Goexit. If a test fails (see Error, Errorf, Fail) and is then skipped, it is still considered to have failed. Execution will continue at the next test or benchmark. See also FailNow. SkipNow must be called from the goroutine running the test, not from other goroutines created during the test. Calling SkipNow does not stop those other goroutines.

​	SkipNow方法将测试标记为已跳过，并通过调用runtime.Goexit停止其执行。如果测试失败(请参见Error、Errorf、Fail)，然后跳过，仍将被视为已失败。执行将在下一个测试或基准测试继续。请参阅FailNow方法。SkipNow方法必须从运行测试的goroutine中调用，而不是从测试期间创建的其他goroutine中调用。调用SkipNow方法不会停止这些其他goroutine。

#### (*F) Skipf  <- go1.18

``` go 
func (c *F) Skipf(format string, args ...any)
```

Skipf is equivalent to Logf followed by SkipNow.

​	Skipf方法等同于Logf方法后跟SkipNow方法。

#### (*F) Skipped  <- go1.18

``` go 
func (f *F) Skipped() bool
```

Skipped reports whether the test was skipped.

​	Skipped方法报告测试是否被跳过。

#### (*F) TempDir  <- go1.18

``` go 
func (c *F) TempDir() string
```

TempDir returns a temporary directory for the test to use. The directory is automatically removed by Cleanup when the test and all its subtests complete. Each subsequent call to t.TempDir returns a unique directory; if the directory creation fails, TempDir terminates the test by calling Fatal.

​	TempDir方法返回一个临时目录供测试使用。当测试和它的所有子测试完成后，该目录会被Cleanup方法自动删除。每次对t.TempDir方法的后续调用都会返回一个唯一的目录；如果目录创建失败，TempDir方法会通过调用Fatal方法终止测试。

### type InternalBenchmark 

``` go 
type InternalBenchmark struct {
	Name string
	F    func(b *B)
}
```

InternalBenchmark is an internal type but exported because it is cross-package; it is part of the implementation of the "go test" command.

​	InternalBenchmark结构体是一个内部类型，但导出因为它是"go test"命令实现的一部分。

### type InternalExample 

``` go 
type InternalExample struct {
	Name      string
	F         func()
	Output    string
	Unordered bool
}
```

### type InternalFuzzTarget  <- go1.18

``` go 
type InternalFuzzTarget struct {
	Name string
	Fn   func(f *F)
}
```

InternalFuzzTarget is an internal type but exported because it is cross-package; it is part of the implementation of the "go test" command.

​	InternalFuzzTarget结构体是一个内部类型，但导出因为它是"go test"命令实现的一部分。

### type InternalTest 

``` go 
type InternalTest struct {
	Name string
	F    func(*T)
}
```

InternalTest is an internal type but exported because it is cross-package; it is part of the implementation of the "go test" command.

​	InternalTest结构体是一个内部类型，但导出因为它是"go test"命令实现的一部分。

### type M  <- go1.4

``` go 
type M struct {
	// contains filtered or unexported fields
}
```

M is a type passed to a TestMain function to run the actual tests.

​	M结构体是一个传递给TestMain函数以运行实际测试的类型。

#### func MainStart  <- go1.4

``` go 
func MainStart(deps testDeps, tests []InternalTest, benchmarks []InternalBenchmark, fuzzTargets []InternalFuzzTarget, examples []InternalExample) *M
```

MainStart is meant for use by tests generated by 'go test'. It is not meant to be called directly and is not subject to the Go 1 compatibility document. It may change signature from release to release.

​	MainStart函数是供由"go test"生成的测试使用的。它不是直接调用的，并且不受Go 1兼容性文档的约束。它可能在不同的版本中改变签名。

#### (*M) Run  <- go1.4

``` go 
func (m *M) Run() (code int)
```

Run runs the tests. It returns an exit code to pass to os.Exit.

​	Run方法运行测试。它返回一个退出码以传递给os.Exit。

### type PB  <- go1.3

``` go 
type PB struct {
	// contains filtered or unexported fields
}
```

A PB is used by RunParallel for running parallel benchmarks.

​	PB结构体被用于RunParallel方法以运行并行基准测试。

#### (*PB) Next  <- go1.3

``` go 
func (pb *PB) Next() bool
```

Next reports whether there are more iterations to execute.

​	Next方法报告是否还有更多的迭代需要执行。

### type T 

``` go 
type T struct {
	// contains filtered or unexported fields
}
```

T is a type passed to Test functions to manage test state and support formatted test logs.

​	`T`结构体是传递给测试函数以管理测试状态并支持格式化测试日志的类型。

A test ends when its Test function returns or calls any of the methods FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. Those methods, as well as the Parallel method, must be called only from the goroutine running the Test function.

​	当Test函数返回或调用任何`FailNow`方法、`Fatal`方法、`Fatalf`方法、`SkipNow`方法、`Skip`方法或`Skipf`方法时，测试结束。这些方法以及`Parallel`方法只能从运行Test函数的goroutine中调用。

The other reporting methods, such as the variations of Log and Error, may be called simultaneously from multiple goroutines.

​	其他报告方法，如`Log`方法和`Error`方法的变体，可以同时从多个goroutine调用。

#### (*T) Chdir <- 1.24.0

```go
func (t *T) Chdir(dir string)
```

Chdir calls os.Chdir(dir) and uses Cleanup to restore the current working directory to its original value after the test. On Unix, it also sets PWD environment variable for the duration of the test.

Because Chdir affects the whole process, it cannot be used in parallel tests or tests with parallel ancestors.

#### (*T) Cleanup  <- go1.14

``` go 
func (c *T) Cleanup(f func())
```

Cleanup registers a function to be called when the test (or subtest) and all its subtests complete. Cleanup functions will be called in last added, first called order.

​	`Cleanup`方法注册一个函数，在测试(或子测试)及其所有子测试完成时调用。`Cleanup`方法将按照最后添加、最先调用的顺序调用。

##### Cleanup  My Example 

```go
func TestExample(t *testing.T) {
    // 在测试开始时执行一些准备工作

    t.Cleanup(func() {
        // 在测试结束时执行清理工作
        // 这个函数将在测试函数返回或调用 t.FailNow、t.Fatal、t.Fatalf、t.SkipNow、t.Skip、t.Skipf 时执行
        // 通常用于释放资源、关闭连接等收尾工作
    })

    // 测试代码

    // 如果测试失败，t.Cleanup 中注册的清理函数将被执行
    // 如果测试通过，t.Cleanup 中注册的清理函数也会被执行
}
```

#### (*T) Context <- 1.24.0

```go
func (c *T) Context() context.Context
```

Context returns a context that is canceled just before Cleanup-registered functions are called.

Cleanup functions can wait for any resources that shut down on Context.Done before the test or benchmark completes.

#### (*T) Deadline  <- go1.15

``` go 
func (t *T) Deadline() (deadline time.Time, ok bool)
```

Deadline reports the time at which the test binary will have exceeded the timeout specified by the -timeout flag.

​	Deadline方法报告测试二进制文件将超过由`-timeout`标志指定的超时时间的时间。

The ok result is false if the -timeout flag indicates “no timeout” (0).

​	如果`-timeout`标志指示"no timeout"(0)，则`ok`结果为`false`。

##### Deadline My Example 	

​	超时设置可帮助确保测试在合理的时间内完成，避免无限执行时间导致的问题。

```go
func TestAbsDeadline(t *testing.T) {
	t.Logf("nowtime=%v", time.Now())
	time.Sleep(5 * time.Second)
	got := Abs(-1)
	// 获取测试函数的截止时间
	deadline, ok := t.Deadline()
	t.Logf("deadline=%v,ok=%v", deadline, ok)

	if !ok {
		t.Fatal("Test deadline not set") // 如果截止时间未设置，标记测试失败
	}

	// 计算距离截止时间的剩余时间
	remainingTime := time.Until(deadline)

	// 在测试代码中使用截止时间的信息
	if remainingTime < 0 {
		t.Fatal("Test has exceeded the deadline") // 如果超过了截止时间，标记测试失败
	}

	// 其他测试逻辑...
	if got != 1 {
		t.Fatalf("Abs(-1) = %f; want 1", got)
	}
	// 如果测试通过，截止时间方法也会被执行
	deadline, ok = t.Deadline()
	t.Logf("deadline=%v,ok=%v", deadline, ok)
}
```

​	请注意，超时时间的设置是可选的，如果不设置，默认情况下测试不会超时（但目前我在go1.21.4中发现，若没有设置`-timeout`，则默认为`-timeout=10m`）。

```cmd
PS F:\Devs\MyCodes\gin_learn\test\unit> go test -v
=== RUN   TestAbsDeadline
    unit_test.go:21: nowtime=2023-12-11 18:32:22.657096 +0800 CST m=+0.005378001
    unit_test.go:26: deadline=2023-12-11 18:42:22.6565738 +0800 CST m=+600.004855801,ok=true
    unit_test.go:46: deadline=2023-12-11 18:42:22.6565738 +0800 CST m=+600.004855801,ok=true
--- PASS: TestAbsDeadline (5.04s)
PASS
ok      github.com/before80/test/unit   5.106s


PS F:\Devs\MyCodes\gin_learn\test\unit> go test -v -timeout=10m
=== RUN   TestAbsDeadline
    unit_test.go:21: nowtime=2023-12-11 18:35:10.9793265 +0800 CST m=+0.005418601
    unit_test.go:26: deadline=2023-12-11 18:45:10.9788156 +0800 CST m=+600.004907701,ok=true
    unit_test.go:46: deadline=2023-12-11 18:45:10.9788156 +0800 CST m=+600.004907701,ok=true
--- PASS: TestAbsDeadline (5.03s)
PASS
ok      github.com/before80/test/unit   5.083s

```

#### (*T) Error 

``` go 
func (c *T) Error(args ...any)
```

Error is equivalent to Log followed by Fail.

​	`Error`方法等同于`Log`方法后面跟着`Fail`方法。

##### Error My Example

```go

func TestAbsError(t *testing.T) {
	tests := []struct {
		x     float64
		want  float64
		eWant float64
	}{
		{-0.1, 0.1, -0.1},
		{-2, 2, -2},
		{-0.3, 0.3, -0.3},
		{-4, 4, -4},
		{-0.5, 0.5, -0.5},
	}

	for _, tt := range tests {
		if got := Abs(tt.x); got != tt.eWant {
			t.Error("Abs(", tt.x, ") = ", got, "; want ", tt.want)
		}
	}
}
```

```cmd
PS F:\Devs\MyCodes\gin_learn\test\unit> go test -v
=== RUN   TestAbsError
    unit_test.go:63: Abs( -0.1 ) =  0.1 ; want  0.1
    unit_test.go:63: Abs( -2 ) =  2 ; want  2
    unit_test.go:63: Abs( -0.3 ) =  0.3 ; want  0.3
    unit_test.go:63: Abs( -4 ) =  4 ; want  4
    unit_test.go:63: Abs( -0.5 ) =  0.5 ; want  0.5
--- FAIL: TestAbsError (0.00s)
FAIL
exit status 1
FAIL    github.com/before80/test/unit   0.062s

```



#### (*T) Errorf 

``` go 
func (c *T) Errorf(format string, args ...any)
```

Errorf is equivalent to Logf followed by Fail.

​	`Errorf`方法等同于`Logf`方法后跟`Fail`方法。

##### Errorf My Example

```go
func TestAbsErrorf(t *testing.T) {
	tests := []struct {
		x     float64
		want  float64
		eWant float64
	}{
		{-0.1, 0.1, -0.1},
		{-2, 2, -2},
		{-0.3, 0.3, -0.3},
		{-4, 4, -4},
		{-0.5, 0.5, -0.5},
	}

	for _, tt := range tests {
		if got := Abs(tt.x); got != tt.eWant {
			t.Errorf("Abs(%v) = %v; want %v", tt.x, got, tt.want)
		}
	}
}
```

```cmd
PS F:\Devs\MyCodes\gin_learn\test\unit> go test -v
=== RUN   TestAbsError
    unit_test.go:64: Abs(-0.1) = 0.1; want 0.1
    unit_test.go:64: Abs(-2) = 2; want 2
    unit_test.go:64: Abs(-0.3) = 0.3; want 0.3
    unit_test.go:64: Abs(-4) = 4; want 4
    unit_test.go:64: Abs(-0.5) = 0.5; want 0.5
--- FAIL: TestAbsError (0.00s)
FAIL
exit status 1
FAIL    github.com/before80/test/unit   0.055s
```



#### (*T) Fail 

``` go 
func (c *T) Fail()
```

Fail marks the function as having failed but continues execution.

​	`Fail`方法标记函数已经失败，但继续执行。

##### Fail My Example

```go
func TestExample(t *testing.T) {
	// 在测试代码中执行一些逻辑...

	// 如果满足某些条件，标记测试失败
	if someCondition {
		t.Fail() // 标记测试失败
	}

	// 继续执行其他测试代码...

	// 如果测试通过，Fail 方法不会终止测试，测试会继续执行
}
```



#### (*T) FailNow <- 只能在Test函数的goroutine中调用

``` go 
func (c *T) FailNow()
```

FailNow marks the function as having failed and stops its execution by calling runtime.Goexit (which then runs all deferred calls in the current goroutine). Execution will continue at the next test or benchmark. FailNow must be called from the goroutine running the test or benchmark function, not from other goroutines created during the test. Calling FailNow does not stop those other goroutines.

​	`FailNow`方法标记函数已经失败并通过调用`runtime.Goexit`(然后在当前goroutine中运行所有延迟调用)停止其执行。执行将继续在下一个测试或基准测试上。`FailNow`方法必须从运行测试或基准测试函数的goroutine中调用，而不是从在测试期间创建的其他goroutine中调用。调用`FailNow`方法不会停止这些其他goroutine。

##### FailNow My Example 

```go
func TestExample(t *testing.T) {
	// 在测试代码中执行一些逻辑...

	// 如果满足某些条件，标记测试失败并立即终止本测试函数的测试
	if someCondition {
		t.FailNow() // 标记测试失败并立即终止本测试函数的测试
	}

	// 以下代码不会被执行，因为本测试函数已经被终止

	// 但会继续执行其他测试函数的代码...
}
```



#### (*T) Failed 

``` go 
func (c *T) Failed() bool
```

Failed reports whether the function has failed.

​	`Failed`方法返回当前测试是否失败。

##### Failed My Example

```go
func TestExample(t *testing.T) {
    // 在测试代码中执行一些逻辑...

    // 如果满足某些条件，标记测试失败
    if someCondition {
        t.Fail() // 标记测试失败
    }

    // 在后续的代码中检查测试是否失败
    if t.Failed() {
        // 在这里处理测试失败的逻辑
        t.Logf("Test failed: %v", someValue)
    }

    // 继续执行其他测试函数的代码...
}
```



#### (*T) Fatal  <- 只能在Test函数的goroutine中调用

``` go 
func (c *T) Fatal(args ...any)
```

Fatal is equivalent to Log followed by FailNow.

​	`Fatal`方法等同于 `Log`方法后跟`FailNow`方法。



##### Fatal My Example

```go
func TestExample(t *testing.T) {
    // 在测试代码中执行一些逻辑...

    // 如果满足某些条件，标记测试失败并立即终止本测试函数的测试
    if someCondition {
        t.Fatal("Test failed: some condition not met") // 标记测试失败并立即终止本测试函数的测试
    }

    // 以下代码不会被执行，因为本测试函数已经被终止

    // 但会继续执行其他测试函数的代码...
}
```



#### (*T) Fatalf <- 只能在Test函数的goroutine中调用

``` go 
func (c *T) Fatalf(format string, args ...any)
```

Fatalf is equivalent to Logf followed by FailNow.

​	`Fatalf`方法等同于 `Logf`方法后跟 `FailNow`方法。

##### Fatalf My Example

```go
func TestExample(t *testing.T) {
    // 在测试代码中执行一些逻辑...   

    // 如果满足某些条件，标记测试失败并立即终止本测试函数的测试
    if someCondition {
        t.Fatalf("Test failed: %s", "some condition not met") // 标记测试失败并立即终止本测试函数的测试
    }

    // 以下代码不会被执行，因为本测试函数已经被终止

    // 但会继续执行其他测试函数的代码...
}
```



#### (*T) Helper  <- go1.9

``` go 
func (c *T) Helper()
```

Helper marks the calling function as a test helper function. When printing file and line information, that function will be skipped. Helper may be called simultaneously from multiple goroutines.

​	`Helper`方法标记调用该函数的函数为测试帮助函数。在打印文件和行信息时，该函数将被跳过。多个 goroutine 可以同时调用 `Helper`方法。

```go
func failExample(t *testing.T) {
    // 在测试辅助函数中调用 Helper 方法
    t.Helper()

    // 在测试辅助函数中发现测试失败的情况
    // 这里使用 t.Error 方法模拟测试失败
    t.Error("This is a helper function example")
}

func TestHelperExample(t *testing.T) {
    // 在测试函数中调用测试辅助函数
    failExample(t)
}
```



#### (*T) Log 

``` go 
func (c *T) Log(args ...any)
```

Log formats its arguments using default formatting, analogous to Println, and records the text in the error log. For tests, the text will be printed only if the test fails or the -test.v flag is set. For benchmarks, the text is always printed to avoid having performance depend on the value of the -test.v flag.

​	`Log`方法使用默认格式化方式格式化其参数，类似于 `Println`，然后将文本记录在错误日志中。对于测试，只有在测试失败或 `-test.v` 标志设置时才会打印该文本。对于基准测试，总是会打印该文本，以避免性能受到 `-test.v` 标志值的影响。

> 个人注释
>
> ​	这里的`-test.v`即`-v`。

##### Log  My Example

```go
func TestExample(t *testing.T) {
    // 在测试代码中执行一些逻辑...   

    t.Log("This is a log message with a value: ", 42)

    // 以下代码会被执行

    // 继续执行其他测试函数的代码...
}
```



#### (*T) Logf 

``` go 
func (c *T) Logf(format string, args ...any)
```

Logf formats its arguments according to the format, analogous to Printf, and records the text in the error log. A final newline is added if not provided. For tests, the text will be printed only if the test fails or the -test.v flag is set. For benchmarks, the text is always printed to avoid having performance depend on the value of the -test.v flag.

​	`Logf`方法根据格式进行参数格式化，类似于 `Printf`，然后将文本记录在错误日志中。如果未提供最后的换行符，则添加一个。对于测试，只有在测试失败或 `-test.v` 标志设置时才会打印该文本。对于基准测试，总是会打印该文本，以避免性能受到 `-test.v` 标志值的影响。

> 个人注释
>
> ​	这里的`-test.v`即`-v`。

##### Logf  My Example

```go
func TestExample(t *testing.T) {
    // 在测试代码中执行一些逻辑...   

    t.Logf("This is a log message with a value: %d", 42)

    // 以下代码会被执行

    // 继续执行其他测试函数的代码...
}
```



#### (*T) Name  <- go1.8

``` go 
func (c *T) Name() string
```

Name returns the name of the running (sub-) test or benchmark.

​	`Name`方法返回当前运行的(子)测试或基准测试的名称。

The name will include the name of the test along with the names of any nested sub-tests. If two sibling sub-tests have the same name, Name will append a suffix to guarantee the returned name is unique.

​	名称将包括测试的名称以及任何嵌套子测试的名称。如果两个同级别的子测试具有相同的名称，则 `Name` 方法将附加后缀以确保返回的名称是唯一的。

##### Name My Example

```go
func TestAbsError(t *testing.T) {
	// 使用 Name 方法获取当前测试函数的名称
	testName := t.Name()
	t.Logf("Running test: %s", testName)

	tests := []struct {
		x     float64
		want  float64
		eWant float64
	}{
		{-0.1, 0.1, -0.1},
		{-2, 2, -2},
		{-0.3, 0.3, -0.3},
		{-4, 4, -4},
		{-0.5, 0.5, -0.5},
	}

	for _, tt := range tests {
		if got := Abs(tt.x); got != tt.want {
			t.Errorf("Abs(%v) = %v; want %v", tt.x, got, tt.want)
		}
	}
}
```

```cmd
PS F:\Devs\MyCodes\gin_learn\test\unit> go test -v
=== RUN   TestAbsError
    unit_test.go:51: Running test: TestAbsError
--- PASS: TestAbsError (0.00s)
PASS
ok      github.com/before80/test/unit   0.061s

```



#### (*T) Parallel <- 只能在Test函数的goroutine中调用

``` go 
func (t *T) Parallel()
```

Parallel signals that this test is to be run in parallel with (and only with) other parallel tests. When a test is run multiple times due to use of -test.count or -test.cpu, multiple instances of a single test never run in parallel with each other.

​	`Parallel`方法表示该测试将与其他并行测试一起运行(且仅与其他并行测试一起运行)。当由于使用 `-test.count` 或 `-test.cpu` 而多次运行测试时，单个测试的多个实例永远不会彼此并行运行。

##### Parallel My Example 

```

```



#### (*T) Run  <- go1.7

``` go 
func (t *T) Run(name string, f func(t *T)) bool
```

Run runs f as a subtest of t called name. It runs f in a separate goroutine and blocks until f returns or calls t.Parallel to become a parallel test. Run reports whether f succeeded (or at least did not fail before calling t.Parallel).

​	`Run`方法将`f`作为`t`的子测试运行，其名称为`name`。它在单独的goroutine中运行f并阻塞，直到`f`返回或调用`t.Parallel`成为并行测试。`Run`方法报告`f`是否成功(或至少在调用`t.Parallel`之前未失败)。

Run may be called simultaneously from multiple goroutines, but all such calls must return before the outer test function for t returns.

​	`Run`方法可以同时从多个goroutine调用，但是所有这些调用必须在`t`返回外部测试函数之前返回。

#### (*T) Setenv  <- go1.17

``` go 
func (t *T) Setenv(key, value string)
```

Setenv calls os.Setenv(key, value) and uses Cleanup to restore the environment variable to its original value after the test.

​	`Setenv`方法调用`os.Setenv(key，value)`并使用Cleanup将环境变量恢复为其原始值。测试完成后，`Cleanup`方法将在最后一个添加的函数优先调用的顺序下被调用。

Because Setenv affects the whole process, it cannot be used in parallel tests or tests with parallel ancestors.

​	由于`Setenv`方法会影响整个进程，因此它不能用于并行测试或具有并行祖先的测试。

#### (*T) Skip  <- go1.1 <- 只能在Test函数的goroutine中调用

``` go 
func (c *T) Skip(args ...any)
```

Skip is equivalent to Log followed by SkipNow.

​	`Skip`方法等效于`Log`方法后跟`SkipNow`方法。

#### (*T) SkipNow  <- go1.1 <- 只能在Test函数的goroutine中调用

``` go 
func (c *T) SkipNow()
```

SkipNow marks the test as having been skipped and stops its execution by calling runtime.Goexit. If a test fails (see Error, Errorf, Fail) and is then skipped, it is still considered to have failed. Execution will continue at the next test or benchmark. See also FailNow. SkipNow must be called from the goroutine running the test, not from other goroutines created during the test. Calling SkipNow does not stop those other goroutines.

​	`SkipNow`方法将测试标记为已跳过，并通过调用`runtime.Goexit`停止其执行。如果测试失败(请参见`Error`，`Errorf`，`Fail`)，然后跳过它，仍将视为已失败。执行将继续在下一个测试或基准中进行。请参见`FailNow`。`SkipNow`方法必须从运行测试的goroutine中调用，而不是从测试期间创建的其他goroutine中调用。调用`SkipNow`不会停止这些其他goroutine。

#### (*T) Skipf  <- go1.1 <- 只能在Test函数的goroutine中调用

``` go 
func (c *T) Skipf(format string, args ...any)
```

Skipf is equivalent to Logf followed by SkipNow.

​	`Skipf`方法等效于`Logf`方法后跟`SkipNow`方法。

#### (*T) Skipped  <- go1.1

``` go 
func (c *T) Skipped() bool
```

Skipped reports whether the test was skipped.

​	`Skipped`方法报告测试是否已被跳过。

#### (*T) TempDir  <- go1.15

``` go 
func (c *T) TempDir() string
```

TempDir returns a temporary directory for the test to use. The directory is automatically removed by Cleanup when the test and all its subtests complete. Each subsequent call to t.TempDir returns a unique directory; if the directory creation fails, TempDir terminates the test by calling Fatal.

​	`TempDir`方法返回测试用于的临时目录。当测试及其所有子测试完成时，`Cleanup`方法会自动删除该目录。每个后续调用`t.TempDir`都会返回一个唯一的目录；如果目录创建失败，则`TempDir`方法通过调用`Fatal`方法终止测试。

### type TB  <- go1.2

``` go 
type TB interface {
	Cleanup(func())
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	Setenv(key, value string)
	Skip(args ...any)
	SkipNow()
	Skipf(format string, args ...any)
	Skipped() bool
	TempDir() string
	// contains filtered or unexported methods
}
```

TB is the interface common to T, B, and F.

​	`TB` 是 `T`、`B` 和 `F` 共有的接口。