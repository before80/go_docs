+++
title = "testing"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# testing

https://pkg.go.dev/testing@go1.20.1



​	testing包提供了对Go程序包进行自动化测试的支持。它旨在与"go test"命令一起使用，该命令自动执行以下形式的任何函数：

``` go 
func TestXxx(*testing.T)
```

其中Xxx不以小写字母开头。函数名称用于标识测试例程。

​	在这些函数内部，使用Error、Fail或相关方法来发出失败信号。

​	要编写新的测试套件，请创建一个文件，其中包含如上所述的TestXxx函数，并将该文件命名为以"_test.go"结尾的名称。该文件将在常规包构建时排除在外，但在运行"go test"命令时将被包括在内。

​	测试文件可以与被测试的包在同一个包中，也可以是具有后缀"_test"的对应包中。

(包内测试)如果测试文件在同一个包中，则可以引用包内未公开的标识符，如下面的示例：

```go  hl_lines="1 1"
package abs

import "testing"

func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

(包外测试)如果测试文件在一个单独的"_test"包中，被测试的包必须显式导入，并且只能使用其导出的标识符。这被称为"黑盒"测试

```go  hl_lines="1 1"
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

欲了解更多细节，请运行 "go help test"和 "go help testflag"。

## 基准测试 Benchmarks 

形如

``` go 
func BenchmarkXxx(*testing.B)
```

的函数被视为基准测试，并在 "go test" 命令提供其 -bench 标志时执行。基准测试按顺序运行。

​	有关测试标志的说明，，见[Testing flags](../../References/CommandDocumentation/go#testing-flags)。

​	基准测试函数的示例如下所示：

``` go 
func BenchmarkRandInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rand.Int()
    }
}
```

​	基准测试函数必须运行目标代码 `b.N` 次。在基准测试执行期间，`b.N` 会进行调整，直到基准测试函数持续时间足够长，以便可靠计时。输出

```
BenchmarkRandInt-8   	68453040	        17.8 ns/op
```

表示循环运行了 68453040 次，每次循环的速度为 17.8 纳秒。

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

​	基准测试结果格式的详细规范在 [Proposal: Go Benchmark Data Format](../../ProposalGoBenchmarkDataFormat)。

​	在[https://golang.org/x/perf/cmd](https://golang.org/x/perf/cmd)中有用于处理基准测试结果的标准工具。特别是，[https://golang.org/x/perf/cmd/benchstat](https://golang.org/x/perf/cmd/benchstat)可以进行统计学上健壮的 A/B 比较。

### 示例

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

​	没有输出注释的示例函数会被编译，但不会被执行。

​	用于声明包、函数F、类型T和类型T上的方法M的示例的命名规则是：

``` go 
func Example() { ... }
func ExampleF() { ... }
func ExampleT() { ... }
func ExampleT_M() { ... }
```

​	可以通过在名称后附加不同的后缀来为包/类型/函数/方法提供多个示例函数。后缀必须以小写字母开头。

``` go 
func Example_suffix() { ... }
func ExampleF_suffix() { ... }
func ExampleT_suffix() { ... }
func ExampleT_M_suffix() { ... }
```

​		当测试文件包含单个示例函数、至少一个其他函数、类型、变量或常量声明，并且没有测试或基准测试函数时，整个测试文件被展示为示例。

## 模糊测试 Fuzzing 

​	'go test' 和 testing 包支持模糊测试，这是一种使用随机生成的输入调用函数以发现单元测试未预期的错误的测试技术。

形如

``` go 
func FuzzXxx(*testing.F)
```

的函数被认为是模糊测试。

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

​		模糊测试维护一个种子语料库或默认情况下运行的一组输入，并可以生成输入。种子输入可以通过调用 `(*F).Add` 或将文件存储在包含模糊测试的包中的 `testdata/fuzz/<Name>` (其中`<Name>`是模糊测试的名称)目录中注册。种子输入是可选的，但是当提供一组具有良好代码覆盖率的小种子输入时，模糊测试引擎可能会更有效地发现错误。这些种子输入还可以作为模糊测试识别的漏洞的回归测试。	

​	在模糊测试中传递给 `(*F).Fuzz` 的函数被认为是模糊目标。一个模糊目标必须接受一个 `*T` 参数，后跟一个或多个随机输入的参数。传递给 `(*F).Add` 的参数类型必须与这些参数的类型相同。模糊目标可以通过调用 `T.Fail`(或调用它的任何方法，如 `T.Error` 或 `T.Fatal`)或引发 panic 的方式来指示它发现了问题，就像测试一样。

​	当启用模糊测试(通过将 `-fuzz` 标志设置为与特定模糊测试匹配的正则表达式)，模糊目标将使用对种子输入进行随机更改生成的参数进行调用。在受支持的平台上，"go test" 使用模糊覆盖率工具编译测试可执行文件。模糊测试引擎使用该工具来查找和缓存扩展覆盖范围的输入，从而增加发现错误的可能性。如果模糊目标对于给定的输入失败，则模糊测试引擎将导致引发失败的输入写入包目录中的 `testdata/fuzz/<Name>` 目录中的文件中。此文件随后用作种子输入。如果无法在该位置写入文件(例如，因为目录为只读)，则模糊测试引擎将文件写入构建缓存中的模糊缓存目录中。

​	当禁用模糊测试时，`F.Add`注册的种子输入和`testdata/fuzz/<Name>`中的种子输入将用于调用模糊目标。在此模式下，模糊测试的行为类似于常规测试，使用`F.Fuzz`启动子测试而不是`T.Run`。

​	有关模糊测试的文档，请参见[Go Fuzzing](../../UsingAndUnderstandingGo/fuzzing)。

## Skipping 

​	跳过测试或基准测试可以通过调用`*T`或`*B`的`Skip`方法来实现：

``` go 
func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    ...
}
```

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

​	每个子测试和子基准测试都有一个唯一的名称：由顶层测试的名称和传递给Run的名称序列组成，用斜杠分隔，可选地带有一个末尾的序列号以进行消除歧义。

​	`-run`、`-bench`和`-fuzz`命令行标志的参数是一个未锚定的正则表达式，用于匹配测试的名称。对于具有多个斜杠分隔元素(如子测试)的测试，参数本身是斜杠分隔的，其中表达式依次匹配每个名称元素。因为它是未锚定的，所以空表达式匹配任何字符串。例如，使用"matching"表示"whose name contains"：

> ​	未锚定的正则表达式是指没有明确指定正则表达式匹配字符串的起始位置和结束位置的表达式。在 Go 中，当使用 `-run` 命令行标志来运行测试用例时，如果正则表达式没有以 `^`(脱字符)开头和 `$`(美元符号)结尾，那么该表达式就是未锚定的。这意味着，表达式可以匹配任何位置出现的字符串。例如，如果运行 `go test -run TestFoo`，那么所有以 `TestFoo` 作为名字的测试用例都将被运行，包括 `TestFooBar`、`TestFooBaz` 等。

```
go test -run ''        # 运行所有测试。
go test -run Foo       # 运行与 "Foo"匹配的顶层测试，如 "TestFooBar"。
go test -run Foo/A=    # 对于匹配 "Foo"的顶级测试，运行匹配 "A="的子测试。
go test -run /A=1      # 对于所有顶级测试，运行匹配 "A=1"的子测试。
go test -fuzz FuzzFoo  # 对匹配 "FuzzFoo"的目标进行模糊处理。
```

​	-run参数也可以用于运行种子语料库中的特定值，以进行调试。例如：

```
go test -run=FuzzFoo/9ddb952d9814
```

​	可以同时设置`-fuzz`和`-run`标志，以模糊处理目标但跳过所有其他测试的执行。

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

​	有时，测试或基准测试程序需要在执行前后进行额外的设置或拆卸。有时还需要控制哪些代码在主线程上运行。为支持这些和其他情况，如果测试文件包含一个函数：

``` go 
func TestMain(m *testing.M)
```

则生成的测试将调用`TestMain(m)`而不是直接运行测试或基准测试。`TestMain`在主goroutine中运行，并可以在调用`m.Run`周围执行任何必要的设置和拆卸操作。`m.Run`将返回可能传递给`os.Exit`的退出代码。如果`TestMain`返回，则测试包装器将`m.Run`的结果传递给`os.Exit`。

​	在调用TestMain时，`flag.Parse`尚未运行。如果TestMain依赖于命令行标志，包括测试包的标志，它应该显式地调用`flag.Parse`。在测试或基准测试函数运行时，命令行标志总是被解析。

​	TestMain的一个简单实现是：

``` go 
func TestMain(m *testing.M) {
	// 如果TestMain使用标志，请在此处调用flag.Parse()
	os.Exit(m.Run())
}
```

​	TestMain 是一个底层的原语，对于普通测试需求，通常使用普通的测试函数就足够了。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func AllocsPerRun  <- go1.1

``` go 
func AllocsPerRun(runs int, f func()) (avg float64)
```

​	AllocsPerRun函数返回在调用 `f` 时的平均分配次数。虽然返回值的类型为 float64，但它始终为整数值。

​	为了计算分配次数，该函数将首先作为预热运行一次。然后将测量指定运行次数内的平均分配次数并返回。

​	AllocsPerRun函数在其测量期间将 `GOMAXPROCS` 设置为 1，并在返回前恢复它。

#### func CoverMode  <- go1.8

``` go 
func CoverMode() string
```

​	CoverMode函数报告测试覆盖模式设置为什么。可选值为 "set"、"count" 或 "atomic"。如果测试覆盖未启用，则返回值为空字符串。

#### func Coverage  <- go1.4

``` go 
func Coverage() float64
```

​	Coverage函数报告当前代码覆盖率，作为范围在 [0, 1] 内的分数。如果未启用覆盖，则 Coverage 返回 0。

​	在运行大型顺序测试集时，每次检查 Coverage函数可以有助于识别哪些测试用例涉及新的代码路径。它不是"go test -cover"和 "go tool cover"生成的报告的替代品。

#### func Init  <- go1.13

``` go 
func Init()
```

​	Init 函数注册测试标志。在运行测试函数之前，"go test"命令会自动注册这些标志，因此只有在调用诸如 Benchmark 等函数而不使用"go test"时才需要调用 Init 函数。

​	如果已经调用了 Init函数，则 Init 函数没有任何效果。

#### func Main 

``` go 
func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, benchmarks []InternalBenchmark, examples []InternalExample)
```

​	Main函数是 "go test" 命令实现的内部函数。它被导出是因为它是跨包的，并且早于 "internal" 包。它不再被 "go test" 使用，但为了其他系统可以模拟 "go test" 使用 Main，尽可能地保留了这个函数。但由于 testing 包添加了新功能，有时 Main 无法更新。模拟 "go test" 的系统应该更新为使用 MainStart。

#### func RegisterCover  <- go1.2

``` go 
func RegisterCover(c Cover)
```

​	RegisterCover函数记录测试的覆盖数据累加器。注意：这个函数是 testing 基础设施的内部函数，可能会更改。它 (目前) 不符合 Go 1 兼容性指南。

#### func RunBenchmarks 

``` go 
func RunBenchmarks(matchString func(pat, str string) (bool, error), benchmarks []InternalBenchmark)
```

​	RunBenchmarks函数是一个内部函数，但导出是因为它是跨包的；它是 "go test" 命令的实现的一部分。

#### func RunExamples 

``` go 
func RunExamples(matchString func(pat, str string) (bool, error), examples []InternalExample) (ok bool)
```

​	RunExamples函数是一个内部函数，但导出是因为它是跨包的；它是 "go test" 命令的实现的一部分。

#### func RunTests 

``` go 
func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)
```

​	RunTests函数是一个内部函数，但导出是因为它是跨包的；它是 "go test" 命令的实现的一部分。

#### func Short 

``` go 
func Short() bool
```

​	Short 函数报告是否设置了 `-test.short` 标志。

#### func Verbose  <- go1.1

``` go 
func Verbose() bool
```

​	Verbose函数报告是否设置了`-test.v`标志。

## 类型

### type B 

``` go 
type B struct {
	N int
	// contains filtered or unexported fields
}
```

​	B是传递给基准测试函数的类型，用于管理基准测试的时间和指定要运行的迭代次数。

​	当Benchmark函数返回或调用FailNow、Fatal、Fatalf、SkipNow、Skip或Skipf方法时，基准测试结束。这些方法只能从运行Benchmark函数的goroutine中调用。其他报告方法，例如Log和Error的变体，可以同时从多个goroutine调用。

​	与测试类似，基准测试日志在执行期间累积，并在完成时转储到标准输出。与测试不同，基准测试日志始终打印，以不隐藏可能影响基准测试结果的输出。

#### (*B) Cleanup  <- go1.14

``` go 
func (c *B) Cleanup(f func())
```

​	Cleanup方法注册一个函数，以在测试(或子测试)及其所有子测试完成时调用。清理函数将按最后添加的先调用的顺序调用。

#### (*B) Elapsed  <- go1.20

``` go 
func (b *B) Elapsed() time.Duration
```

​	Elapsed方法返回基准测试的测量经过时间。Elapsed报告的持续时间与StartTimer方法、StopTimer方法和ResetTimer方法测量的时间相匹配。

#### (*B) Error 

``` go 
func (c *B) Error(args ...any)
```

​	Error方法等效于Log方法后跟Fail方法。

#### (*B) Errorf 

``` go 
func (c *B) Errorf(format string, args ...any)
```

​	Errorf方法等效于Logf方法后跟Fail方法。

#### (*B) Fail 

``` go 
func (c *B) Fail()
```

​	Fail方法标记函数失败，但继续执行。

#### (*B) FailNow 

``` go 
func (c *B) FailNow()
```

​	FailNow方法标记函数失败并通过调用runtime.Goexit(然后在当前goroutine中运行所有延迟调用)停止其执行。执行将继续在下一个测试或基准测试中。FailNow方法必须从运行测试或基准测试函数的goroutine中调用，而不是从在测试期间创建的其他goroutine中调用。调用FailNow方法不会停止这些其他goroutine。

#### (*B) Failed 

``` go 
func (c *B) Failed() bool
```

​	Failed方法报告函数是否已失败。

#### (*B) Fatal 

``` go 
func (c *B) Fatal(args ...any)
```

​	Fatal方法等效于 Log方法后跟 FailNow方法。

#### (*B) Fatalf 

``` go 
func (c *B) Fatalf(format string, args ...any)
```

​	Fatalf方法等效于 Logf方法后跟 FailNow方法。

#### (*B) Helper  <- go1.9

``` go 
func (c *B) Helper()
```

​	Helper方法将调用函数标记为测试辅助函数。在打印文件和行信息时，该函数将被跳过。Helper方法可以同时从多个 goroutine 中调用。

#### (*B) Log 

``` go 
func (c *B) Log(args ...any)
```

​	Log方法使用默认格式对其参数进行格式化，类似于 Println，并将文本记录在错误日志中。对于测试，只有在测试失败或设置了 `-test.v` 标志时，才会打印该文本。对于基准测试，始终会打印文本，以避免性能依赖于 `-test.v` 标志的值。

#### (*B) Logf 

``` go 
func (c *B) Logf(format string, args ...any)
```

​	Logf方法根据格式对其参数进行格式化，类似于 Printf，并将文本记录在错误日志中。如果没有提供最终换行符，则会添加。对于测试，只有在测试失败或设置了 `-test.v` 标志时，才会打印该文本。对于基准测试，始终会打印文本，以避免性能依赖于 `-test.v` 标志的值。

#### (*B) Name  <- go1.8

``` go 
func (c *B) Name() string
```

​	Name方法返回正在运行的(子)测试或基准测试的名称。

​	名称将包括测试的名称以及任何嵌套的子测试的名称。如果两个同级的子测试具有相同的名称，则 Name 将附加后缀以保证返回的名称是唯一的。

#### (*B) ReportAllocs  <- go1.1

``` go 
func (b *B) ReportAllocs()
```

​	ReportAllocs方法为此基准测试启用 malloc 统计信息。它相当于设置 `-test.benchmem`，但仅影响调用 ReportAllocs方法的基准测试函数。

#### (*B) ReportMetric  <- go1.13

``` go 
func (b *B) ReportMetric(n float64, unit string)
```

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

​	ResetTimer方法将经过的基准测试时间和内存分配计数器归零，并删除用户报告的度量标准。它不会影响定时器是否在运行。

#### (*B) Run  <- go1.7

``` go 
func (b *B) Run(name string, f func(b *B)) bool
```

​	Run方法将f作为具有给定名称的子基准测试运行。它报告是否有任何失败。

​	子基准测试与任何其他基准测试相似。调用Run方法至少一次的基准测试本身不会被测量，并且将以N=1的方式调用一次。

#### (*B) RunParallel  <- go1.3

``` go 
func (b *B) RunParallel(body func(*PB))
```

​	RunParallel方法在并行中运行基准测试。它创建多个goroutine并将b.N次迭代分布在它们之间。 goroutine的数量默认为GOMAXPROCS。要增加非CPU绑定的基准测试的并行性，请在RunParallel之前调用SetParallelism。 RunParallel方法通常与go test -cpu标志一起使用。

​	body函数将在每个goroutine中运行。它应该设置任何goroutine-local状态，然后迭代直到pb.Next返回false。它不应该使用StartTimer方法、StopTimer方法或ResetTimer方法，因为它们具有全局效果。它也不应该调用Run方法。

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
	// 这是针对text/template.Template.Execute在单个对象上进行并行基准测试的代码。
	testing.Benchmark(func(b *testing.B) {
		templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
		// RunParallel将创建GOMAXPROCS个goroutine并在它们之间分配工作。
		b.RunParallel(func(pb *testing.PB) {
			// 每个goroutine都有自己的bytes.Buffer。
			var buf bytes.Buffer
			for pb.Next() {
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

​	SetBytes方法记录在单个操作中处理的字节数。如果调用了此函数，则基准测试将报告ns/op和MB/s。

#### (*B) SetParallelism  <- go1.3

``` go 
func (b *B) SetParallelism(p int)
```

​	SetParallelism方法设置RunParallel方法要使用的goroutine数为`p * GOMAXPROCS`。对于CPU绑定的基准测试，通常不需要调用SetParallelism方法。如果p小于1，则此调用将无效。

#### (*B) Setenv  <- go1.17

``` go 
func (c *B) Setenv(key, value string)
```

​	Setenv方法调用os.Setenv(key, value)，并使用Cleanup方法在测试后将环境变量恢复到其原始值。

​	因为Setenv方法会影响整个进程，所以它不能在并行测试或具有并行祖先的测试中使用。

#### (*B) Skip  <- go1.1

``` go 
func (c *B) Skip(args ...any)
```

​	Skip方法等同于Log方法后跟SkipNow方法。

#### (*B) SkipNow  <- go1.1

``` go 
func (c *B) SkipNow()
```

​	SkipNow方法将测试标记为已跳过，并通过调用runtime.Goexit停止其执行。如果测试失败(请参见Error、Errorf、Fail)，然后跳过，它仍然被认为是已失败的。执行将在下一个测试或基准测试上继续。另请参见FailNow。SkipNow必须从运行测试的goroutine而不是从测试期间创建的其他goroutine调用。调用SkipNow方法不会停止这些其他goroutine。

#### (*B) Skipf  <- go1.1

``` go 
func (c *B) Skipf(format string, args ...any)
```

​	Skipf方法等同于Logf方法后跟SkipNow方法。

#### (*B) Skipped  <- go1.1

``` go 
func (c *B) Skipped() bool
```

​	Skipped方法报告测试是否被跳过。

#### (*B) StartTimer 

``` go 
func (b *B) StartTimer()
```

​	StartTimer方法开始计时测试。该函数会在基准测试开始前自动调用，但也可以在需要测量但不想计入时间的复杂初始化操作后手动调用以恢复计时。

#### (*B) StopTimer 

``` go 
func (b *B) StopTimer()
```

​	StopTimer方法停止计时测试。可以在进行不需要测量时间的复杂初始化操作时使用它以暂停计时。

#### (*B) TempDir  <- go1.15

``` go 
func (c *B) TempDir() string
```

​	TempDir方法返回测试使用的临时目录。该目录会在测试及其所有子测试完成时由 Cleanup方法自动删除。每次调用 t.TempDir方法都会返回一个唯一的目录；如果目录创建失败，TempDir方法会通过调用 Fatal 方法终止测试。

### type BenchmarkResult 

``` go 
type BenchmarkResult struct {
	N         int           // 迭代次数。
	T         time.Duration // 所花费的总时间。
	Bytes     int64         // 单次迭代处理的字节数。
	MemAllocs uint64        // 内存分配总次数。
	MemBytes  uint64        // 分配的总字节数。

	// Extra 记录 ReportMetric 报告的其他度量标准。
	Extra map[string]float64
}
```

​	BenchmarkResult结构体包含基准测试运行的结果。

#### func Benchmark 

``` go 
func Benchmark(f func(b *B)) BenchmarkResult
```

​	Benchmark函数对单个函数进行基准测试。它适用于创建不使用 "go test" 命令的自定义基准测试。

​	如果 f 依赖于测试标志，则必须在调用 Benchmark 之前使用 Init方法注册这些标志，并在调用 flag.Parse 之前进行。

​	如果 f 调用 Run方法，则结果将是估计运行所有不调用 Run方法的子基准测试的结果。

#### (BenchmarkResult) AllocedBytesPerOp  <- go1.1

``` go 
func (r BenchmarkResult) AllocedBytesPerOp() int64
```

​	AllocedBytesPerOp方法返回 "B/op" 指标，它被计算为 r.MemBytes / r.N。

#### (BenchmarkResult) AllocsPerOp  <- go1.1

``` go 
func (r BenchmarkResult) AllocsPerOp() int64
```

​	AllocsPerOp方法返回 "allocs/op" 指标，它被计算为 r.MemAllocs / r.N。

#### (BenchmarkResult) MemString  <- go1.1

``` go 
func (r BenchmarkResult) MemString() string
```

​	MemString方法以与 'go test' 相同的格式返回 r.AllocedBytesPerOp 和 r.AllocsPerOp。

#### (BenchmarkResult) NsPerOp 

``` go 
func (r BenchmarkResult) NsPerOp() int64
```

​	NsPerOp方法返回"ns/op"指标。

#### (BenchmarkResult) String 

``` go 
func (r BenchmarkResult) String() string
```

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

​	CoverBlock结构体记录单个基本块的覆盖数据。这些字段是从编辑器中的 1 开始计数的：例如，文件的开头行号是 1。列以字节为单位测量。注意：此结构体对于测试基础设施是内部的，可能会更改。它尚未(但可能会)受到 Go 1 兼容性指南的影响。

### type F  <- go1.18

``` go 
type F struct {
	// contains filtered or unexported fields
}
```

​	F结构体是传递给模糊测试的类型。

​	模糊测试将生成的输入针对提供的模糊目标运行，可以找到并报告被测试代码中的潜在错误。

​	默认情况下，模糊测试运行种子语料库，其中包括 `(*F).Add` 中提供的条目和 `testdata/fuzz/<FuzzTestName>` 目录中的条目。在任何必要的设置和调用 `(*F).Add` 后，模糊测试必须调用 `(*F).Fuzz` 提供模糊目标。有关示例，请参见 testing 包文档，并参见 F.Fuzz 和 F.Add 方法文档获取详细信息。

​	`*F` 方法只能在 `(*F).Fuzz` 之前调用。一旦测试执行模糊目标，只能使用 `(*T)` 方法。`(*F).Failed` 和 `(*F).Name` 是 `(*F).Fuzz` 函数中允许的仅有的 `*F` 方法。

#### (*F) Add  <- go1.18

``` go 
func (f *F) Add(args ...any)
```

​	Add方法将参数添加到 fuzz 测试的种子语料库中。如果在模糊目标之后或其中调用，将不起作用，并且 `args` 必须与模糊目标的参数匹配。

#### (*F) Cleanup  <- go1.18

``` go 
func (c *F) Cleanup(f func())
```

​	Cleanup方法注册一个在测试(或子测试)及其所有子测试完成时调用的函数。Cleanup方法将按照最后添加的先调用的顺序调用。

#### (*F) Error  <- go1.18

``` go 
func (c *F) Error(args ...any)
```

​	Error 方法等同于 Log 方法后跟 Fail 方法。

#### (*F) Errorf  <- go1.18

``` go 
func (c *F) Errorf(format string, args ...any)
```

​	Errorf方法等同于 Logf 方法后跟 Fail 方法。

#### (*F) Fail  <- go1.18

``` go 
func (f *F) Fail()
```

​	Fail 方法标记函数失败但继续执行。

#### (*F) FailNow  <- go1.18

``` go 
func (c *F) FailNow()
```

​	FailNow方法标记函数失败并通过调用 runtime.Goexit 停止其执行(然后运行当前 goroutine 中的所有延迟调用)。执行将在下一个测试或基准中继续。FailNow方法必须从运行测试或基准函数的 goroutine 调用，而不是从测试期间创建的其他 goroutine 中调用。调用 FailNow方法不会停止这些其他 goroutine。

#### (*F) Failed  <- go1.18

``` go 
func (c *F) Failed() bool
```

​	Failed 方法报告函数是否已失败。

#### (*F) Fatal  <- go1.18

``` go 
func (c *F) Fatal(args ...any)
```

​	Fatal 方法等同于 Log 方法后跟 FailNow 方法。

#### (*F) Fatalf  <- go1.18

``` go 
func (c *F) Fatalf(format string, args ...any)
```

​	Fatalf 方法等同于 Logf 方法后跟 FailNow 方法。

#### (*F) Fuzz  <- go1.18

``` go 
func (f *F) Fuzz(ff any)
```

​	Fuzz 方法运行fuzz函数，ff，进行模糊测试。如果 ff 对一组参数失败，则这些参数将被添加到种子语料库中。

​	ff 必须是一个没有返回值的函数，其第一个参数为 `*T`，其余参数为要进行 fuzz 的类型。例如：

```
f.Fuzz(func(t *testing.T, b []byte, i int) { ... })
```

​	以下类型是允许的：[]byte、string、bool、byte、rune、float32、float64、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64。更多类型可能会在未来支持。

​	ff 不能调用任何 `*F` 方法，例如 `(*F).Log`、`(*F).Error`、`(*F).Skip`。应改用相应的 `*T` 方法。`(*F).Failed` 和 `(*F).Name` 是唯一允许在 `(*F).Fuzz` 函数中使用的 `*F` 方法

​	该函数应该快速且确定性，其行为不应该依赖于共享状态。不应在执行模糊函数之间保留可变输入参数或指向它们的指针，因为支持它们的内存可能在随后的调用期间被修改。ff不能修改由fuzz引擎提供的参数的底层数据。

​	在进行模糊测试时，F.Fuzz方法不会返回，直到发现问题，时间耗尽(使用`-fuzztime`设置)，或测试过程被信号中断。F.Fuzz应该仅被调用一次，除非在之前调用了F.Skip或F.Fail。

#### (*F) Helper  <- go1.18

``` go 
func (f *F) Helper()
```

​	Helper方法标记调用函数为测试辅助函数。在打印文件和行信息时，该函数将被跳过。Helper方法可以同时从多个goroutine调用。

#### (*F) Log  <- go1.18

``` go 
func (c *F) Log(args ...any)
```

​	Log方法使用默认格式对其参数进行格式化，类似于Println，并将文本记录在错误日志中。对于测试，仅当测试失败或设置了`-test.v`标志时才会打印文本。对于基准测试，为避免性能依赖于`-test.v`标志的值，始终打印文本。

#### (*F) Logf  <- go1.18

``` go 
func (c *F) Logf(format string, args ...any)
```

​	Logf方法根据格式对其参数进行格式化，类似于Printf，并将文本记录在错误日志中。如果没有提供最后一个换行符，则会添加一个。对于测试，仅当测试失败或设置了`-test.v`标志时才会打印文本。对于基准测试，为避免性能依赖于`-test.v`标志的值，始终打印文本。

#### (*F) Name  <- go1.18

``` go 
func (c *F) Name() string
```

​	Name方法返回正在运行的(子)测试或基准测试的名称。名称将包括测试的名称以及任何嵌套的子测试的名称。如果两个同级别的子测试名称相同，则Name方法将附加后缀以确保返回的名称是唯一的。

#### (*F) Setenv  <- go1.18

``` go 
func (c *F) Setenv(key, value string)
```

​	Setenv方法调用os.Setenv(key, value)，并使用Cleanup方法将环境变量还原为其原始值。由于Setenv方法影响整个进程，因此无法在并行测试或具有并行祖先的测试中使用。

#### (*F) Skip  <- go1.18

``` go 
func (c *F) Skip(args ...any)
```

​	Skip方法等同于Log方法后跟SkipNow方法。

#### (*F) SkipNow  <- go1.18

``` go 
func (c *F) SkipNow()
```

​	SkipNow方法将测试标记为已跳过，并通过调用runtime.Goexit停止其执行。如果测试失败(请参见Error、Errorf、Fail)，然后跳过，仍将被视为已失败。执行将在下一个测试或基准测试继续。请参阅FailNow方法。SkipNow方法必须从运行测试的goroutine中调用，而不是从测试期间创建的其他goroutine中调用。调用SkipNow方法不会停止这些其他goroutine。

#### (*F) Skipf  <- go1.18

``` go 
func (c *F) Skipf(format string, args ...any)
```

​	Skipf方法等同于Logf方法后跟SkipNow方法。

#### (*F) Skipped  <- go1.18

``` go 
func (f *F) Skipped() bool
```

​	Skipped方法报告测试是否被跳过。

#### (*F) TempDir  <- go1.18

``` go 
func (c *F) TempDir() string
```

​	TempDir方法返回一个临时目录供测试使用。当测试和它的所有子测试完成后，该目录会被Cleanup方法自动删除。每次对t.TempDir方法的后续调用都会返回一个唯一的目录；如果目录创建失败，TempDir方法会通过调用Fatal方法终止测试。

### type InternalBenchmark 

``` go 
type InternalBenchmark struct {
	Name string
	F    func(b *B)
}
```

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

​	InternalFuzzTarget结构体是一个内部类型，但导出因为它是"go test"命令实现的一部分。

### type InternalTest 

``` go 
type InternalTest struct {
	Name string
	F    func(*T)
}
```

​	InternalTest结构体是一个内部类型，但导出因为它是"go test"命令实现的一部分。

### type M  <- go1.4

``` go 
type M struct {
	// contains filtered or unexported fields
}
```

​	M结构体是一个传递给TestMain函数以运行实际测试的类型。

#### func MainStart  <- go1.4

``` go 
func MainStart(deps testDeps, tests []InternalTest, benchmarks []InternalBenchmark, fuzzTargets []InternalFuzzTarget, examples []InternalExample) *M
```

​	MainStart函数是供由"go test"生成的测试使用的。它不是直接调用的，并且不受Go 1兼容性文档的约束。它可能在不同的版本中改变签名。

#### (*M) Run  <- go1.4

``` go 
func (m *M) Run() (code int)
```

​	Run方法运行测试。它返回一个退出码以传递给os.Exit。

### type PB  <- go1.3

``` go 
type PB struct {
	// contains filtered or unexported fields
}
```

​	PB结构体被用于RunParallel方法以运行并行基准测试。

#### (*PB) Next  <- go1.3

``` go 
func (pb *PB) Next() bool
```

​	Next方法报告是否还有更多的迭代需要执行。

### type T 

``` go 
type T struct {
	// contains filtered or unexported fields
}
```

​	T结构体是传递给测试函数以管理测试状态并支持格式化测试日志的类型。

​	当Test函数返回或调用任何FailNow方法、Fatal方法、Fatalf方法、SkipNow方法、Skip方法或Skipf方法时，测试结束。这些方法以及Parallel方法只能从运行Test函数的goroutine中调用。

​	其他报告方法，如Log方法和Error方法的变体，可以同时从多个goroutine调用。

#### (*T) Cleanup  <- go1.14

``` go 
func (c *T) Cleanup(f func())
```

​	Cleanup方法注册一个函数，在测试(或子测试)及其所有子测试完成时调用。Cleanup方法将按照最后添加、最先调用的顺序调用。

#### (*T) Deadline  <- go1.15

``` go 
func (t *T) Deadline() (deadline time.Time, ok bool)
```

​	Deadline方法报告测试二进制文件将超过由`-timeout`标志指定的超时时间的时间。

​	如果`-timeout`标志指示"no timeout"(0)，则ok结果为false。

#### (*T) Error 

``` go 
func (c *T) Error(args ...any)
```

​	Error方法等同于Log方法后面跟着Fail方法。

#### (*T) Errorf 

``` go 
func (c *T) Errorf(format string, args ...any)
```

​	Errorf方法等同于Logf方法后跟Fail方法。

#### (*T) Fail 

``` go 
func (c *T) Fail()
```

​	Fail方法标记函数已经失败，但继续执行。

#### (*T) FailNow 

``` go 
func (c *T) FailNow()
```

​	FailNow方法标记函数已经失败并通过调用runtime.Goexit(然后在当前goroutine中运行所有延迟调用)停止其执行。执行将继续在下一个测试或基准测试上。FailNow方法必须从运行测试或基准测试函数的goroutine中调用，而不是从在测试期间创建的其他goroutine中调用。调用FailNow方法不会停止这些其他goroutine。

#### (*T) Failed 

``` go 
func (c *T) Failed() bool
```

​	Failed方法返回当前测试是否失败。

#### (*T) Fatal 

``` go 
func (c *T) Fatal(args ...any)
```

​	Fatal方法等同于 Log方法后跟FailNow方法。

#### (*T) Fatalf 

``` go 
func (c *T) Fatalf(format string, args ...any)
```

​	Fatalf方法等同于 Logf方法后跟 FailNow方法。

#### (*T) Helper  <- go1.9

``` go 
func (c *T) Helper()
```

​	Helper方法标记调用该函数的函数为测试帮助函数。在打印文件和行信息时，该函数将被跳过。多个 goroutine 可以同时调用 Helper方法。

#### (*T) Log 

``` go 
func (c *T) Log(args ...any)
```

​	Log方法使用默认格式化方式格式化其参数，类似于 Println，然后将文本记录在错误日志中。对于测试，只有在测试失败或 `-test.v` 标志设置时才会打印该文本。对于基准测试，总是会打印该文本，以避免性能受到 `-test.v` 标志值的影响。

#### (*T) Logf 

``` go 
func (c *T) Logf(format string, args ...any)
```

​	Logf方法根据格式进行参数格式化，类似于 Printf，然后将文本记录在错误日志中。如果未提供最后的换行符，则添加一个。对于测试，只有在测试失败或 `-test.v` 标志设置时才会打印该文本。对于基准测试，总是会打印该文本，以避免性能受到 `-test.v` 标志值的影响。

#### (*T) Name  <- go1.8

``` go 
func (c *T) Name() string
```

​	Name方法返回当前运行的(子)测试或基准测试的名称。

​	名称将包括测试的名称以及任何嵌套子测试的名称。如果两个同级别的子测试具有相同的名称，则 Name 方法将附加后缀以确保返回的名称是唯一的。

#### (*T) Parallel 

``` go 
func (t *T) Parallel()
```

​	Parallel方法表示该测试将与其他并行测试一起运行(且仅与其他并行测试一起运行)。当由于使用 `-test.count` 或 `-test.cpu` 而多次运行测试时，单个测试的多个实例永远不会彼此并行运行。

#### (*T) Run  <- go1.7

``` go 
func (t *T) Run(name string, f func(t *T)) bool
```

​	Run方法将`f`作为`t`的子测试运行，其名称为name。它在单独的goroutine中运行f并阻塞，直到`f`返回或调用`t.Parallel`成为并行测试。Run方法报告`f`是否成功(或至少在调用`t.Parallel`之前未失败)。

​	Run方法可以同时从多个goroutine调用，但是所有这些调用必须在`t`返回外部测试函数之前返回。

#### (*T) Setenv  <- go1.17

``` go 
func (t *T) Setenv(key, value string)
```

​	Setenv方法调用os.Setenv(key，value)并使用Cleanup将环境变量恢复为其原始值。测试完成后，Cleanup方法将在最后一个添加的函数优先调用的顺序下被调用。

​	由于Setenv方法会影响整个进程，因此它不能用于并行测试或具有并行祖先的测试。

#### (*T) Skip  <- go1.1

``` go 
func (c *T) Skip(args ...any)
```

​	Skip方法等效于Log方法后跟SkipNow方法。

#### (*T) SkipNow  <- go1.1

``` go 
func (c *T) SkipNow()
```

​	SkipNow方法将测试标记为已跳过，并通过调用runtime.Goexit停止其执行。如果测试失败(请参见Error，Errorf，Fail)，然后跳过它，仍将视为已失败。执行将继续在下一个测试或基准中进行。请参见FailNow。SkipNow方法必须从运行测试的goroutine中调用，而不是从测试期间创建的其他goroutine中调用。调用SkipNow不会停止这些其他goroutine。

#### (*T) Skipf  <- go1.1

``` go 
func (c *T) Skipf(format string, args ...any)
```

​	Skipf方法等效于Logf方法后跟SkipNow方法。

#### (*T) Skipped  <- go1.1

``` go 
func (c *T) Skipped() bool
```

​	Skipped方法报告测试是否已被跳过。

#### (*T) TempDir  <- go1.15

``` go 
func (c *T) TempDir() string
```

​	TempDir方法返回测试用于的临时目录。当测试及其所有子测试完成时，Cleanup方法会自动删除该目录。每个后续调用t.TempDir都会返回一个唯一的目录；如果目录创建失败，则TempDir方法通过调用Fatal方法终止测试。

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

TB 是 T、B 和 F 共有的接口。