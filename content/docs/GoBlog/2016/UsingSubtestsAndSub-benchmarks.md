+++
title = "使用子测试和子基准"
weight = 5
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Using Subtests and Sub-benchmarks - 使用子测试和子基准

https://go.dev/blog/subtests

Marcel van Lohuizen
3 October 2016

## Introduction 简介

In Go 1.7, the `testing` package introduces a Run method on the [`T`](https://go.dev/pkg/testing/#T.Run) and [`B`](https://go.dev/pkg/testing/#B.Run) types that allows for the creation of subtests and sub-benchmarks. The introduction of subtests and sub-benchmarks enables better handling of failures, fine-grained control of which tests to run from the command line, control of parallelism, and often results in simpler and more maintainable code.

在Go 1.7中，测试包在T和B类型上引入了一个运行方法，允许创建子测试和子基准。子测试和子基准的引入可以更好地处理故障，细粒度地控制从命令行中运行的测试，控制并行性，并经常导致更简单和更可维护的代码。

## Table-driven tests basics 表驱动测试的基础知识

Before digging into the details, let’s first discuss a common way of writing tests in Go. A series of related checks can be implemented by looping over a slice of test cases:

在深入研究细节之前，我们首先讨论一下Go中编写测试的一种常见方式。一系列相关的检查可以通过循环测试案例的片断来实现：

```go linenums="1"
func TestTime(t *testing.T) {
    testCases := []struct {
        gmt  string
        loc  string
        want string
    }{
        {"12:31", "Europe/Zuri", "13:31"},     // incorrect location name
        {"12:31", "America/New_York", "7:31"}, // should be 07:31
        {"08:08", "Australia/Sydney", "18:08"},
    }
    for _, tc := range testCases {
        loc, err := time.LoadLocation(tc.loc)
        if err != nil {
            t.Fatalf("could not load location %q", tc.loc)
        }
        gmt, _ := time.Parse("15:04", tc.gmt)
        if got := gmt.In(loc).Format("15:04"); got != tc.want {
            t.Errorf("In(%s, %s) = %s; want %s", tc.gmt, tc.loc, got, tc.want)
        }
    }
}
```

This approach, commonly referred to as table-driven tests, reduces the amount of repetitive code compared to repeating the same code for each test and makes it straightforward to add more test cases.

这种方法，通常被称为表驱动测试，与每个测试重复相同的代码相比，减少了重复的代码量，并且可以直接增加更多的测试案例。

## Table-driven benchmarks 表驱动的基准

Before Go 1.7 it was not possible to use the same table-driven approach for benchmarks. A benchmark tests the performance of an entire function, so iterating over benchmarks would just measure all of them as a single benchmark.

在Go 1.7之前，不可能将同样的表驱动方法用于基准测试。基准测试的是整个函数的性能，所以对基准进行迭代，只是把所有函数作为一个单一的基准来测量。

A common workaround was to define separate top-level benchmarks that each call a common function with different parameters. For instance, before 1.7 the `strconv` package’s benchmarks for `AppendFloat` looked something like this:

一个常见的解决方法是定义单独的顶层基准，每个基准都以不同的参数调用一个共同的函数。例如，在1.7之前，strconv包的AppendFloat的基准看起来是这样的：

```go linenums="1"
func benchmarkAppendFloat(b *testing.B, f float64, fmt byte, prec, bitSize int) {
    dst := make([]byte, 30)
    b.ResetTimer() // Overkill here, but for illustrative purposes.
    for i := 0; i < b.N; i++ {
        AppendFloat(dst[:0], f, fmt, prec, bitSize)
    }
}

func BenchmarkAppendFloatDecimal(b *testing.B) { benchmarkAppendFloat(b, 33909, 'g', -1, 64) }
func BenchmarkAppendFloat(b *testing.B)        { benchmarkAppendFloat(b, 339.7784, 'g', -1, 64) }
func BenchmarkAppendFloatExp(b *testing.B)     { benchmarkAppendFloat(b, -5.09e75, 'g', -1, 64) }
func BenchmarkAppendFloatNegExp(b *testing.B)  { benchmarkAppendFloat(b, -5.11e-95, 'g', -1, 64) }
func BenchmarkAppendFloatBig(b *testing.B)     { benchmarkAppendFloat(b, 123456789123456789123456789, 'g', -1, 64) }
...
```

Using the `Run` method available in Go 1.7, the same set of benchmarks is now expressed as a single top-level benchmark:

使用Go 1.7中的Run方法，同样的一组基准现在被表达为一个单一的顶层基准：

```go linenums="1"
func BenchmarkAppendFloat(b *testing.B) {
    benchmarks := []struct{
        name    string
        float   float64
        fmt     byte
        prec    int
        bitSize int
    }{
        {"Decimal", 33909, 'g', -1, 64},
        {"Float", 339.7784, 'g', -1, 64},
        {"Exp", -5.09e75, 'g', -1, 64},
        {"NegExp", -5.11e-95, 'g', -1, 64},
        {"Big", 123456789123456789123456789, 'g', -1, 64},
        ...
    }
    dst := make([]byte, 30)
    for _, bm := range benchmarks {
        b.Run(bm.name, func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                AppendFloat(dst[:0], bm.float, bm.fmt, bm.prec, bm.bitSize)
            }
        })
    }
}
```

Each invocation of the `Run` method creates a separate benchmark. An enclosing benchmark function that calls a `Run` method is only run once and is not measured.

运行方法的每次调用都会创建一个单独的基准。一个调用Run方法的封闭式基准函数只运行一次，不被测量。

The new code has more lines of code, but is more maintainable, more readable, and consistent with the table-driven approach commonly used for testing. Moreover, common setup code is now shared between runs while eliminating the need to reset the timer.

新的代码有更多的行数，但更容易维护，更容易阅读，并且与测试中常用的表驱动方法一致。此外，共同的设置代码现在在运行之间共享，同时消除了重置定时器的需要。

## Table-driven tests using subtests 使用子测试的表驱动测试

Go 1.7 also introduces a `Run` method for creating subtests. This test is a rewritten version of our earlier example using subtests:

Go 1.7还引入了创建子测试的运行方法。这个测试是我们先前使用子测试的例子的重写版本：

```go linenums="1"
func TestTime(t *testing.T) {
    testCases := []struct {
        gmt  string
        loc  string
        want string
    }{
        {"12:31", "Europe/Zuri", "13:31"},
        {"12:31", "America/New_York", "7:31"},
        {"08:08", "Australia/Sydney", "18:08"},
    }
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("%s in %s", tc.gmt, tc.loc), func(t *testing.T) {
            loc, err := time.LoadLocation(tc.loc)
            if err != nil {
                t.Fatal("could not load location")
            }
            gmt, _ := time.Parse("15:04", tc.gmt)
            if got := gmt.In(loc).Format("15:04"); got != tc.want {
                t.Errorf("got %s; want %s", got, tc.want)
            }
        })
    }
}
```

The first thing to note is the difference in output from the two implementations. The original implementation prints:

首先要注意的是两个实现的输出的不同。原始的实现会打印：

```shell linenums="1"
--- FAIL: TestTime (0.00s)
    time_test.go:62: could not load location "Europe/Zuri"
```

Even though there are two errors, execution of the test halts on the call to `Fatalf` and the second test never runs.

尽管有两个错误，但测试的执行在调用Fatalf时停止了，第二个测试从未运行。

The implementation using `Run` prints both:

使用Run的实现会打印出这两个：

```shell linenums="1"
--- FAIL: TestTime (0.00s)
    --- FAIL: TestTime/12:31_in_Europe/Zuri (0.00s)
        time_test.go:84: could not load location
    --- FAIL: TestTime/12:31_in_America/New_York (0.00s)
        time_test.go:88: got 07:31; want 7:31
```

`Fatal` and its siblings causes a subtest to be skipped but not its parent or subsequent subtests.

Fatal和它的兄弟姐妹会导致一个子测试被跳过，但不会跳过它的父测试或后续子测试。

Another thing to note is the shorter error messages in the new implementation. Since the subtest name uniquely identifies the subtest there is no need to identify the test again within the error messages.

另一件需要注意的事情是，在新的实现中，错误信息更短。由于子测试名称唯一地标识了子测试，因此没有必要在错误信息中再次标识测试。

There are several other benefits to using subtests or sub-benchmarks, as clarified by the following sections.

使用子测试或子基准还有其他一些好处，正如以下各节所阐明的。

## Running specific tests or benchmarks 运行特定的测试或基准

Both subtests and sub-benchmarks can be singled out on the command line using the [`-run` or `-bench` flag](https://go.dev/cmd/go/#hdr-Description_of_testing_flags). Both flags take a slash-separated list of regular expressions that match the corresponding parts of the full name of the subtest or sub-benchmark. 

子测试和子基准都可以使用-run或-bench标志在命令行中单独列出。这两个标志都接受一个斜线分隔的正则表达式列表，该列表与子测试或子基准的全名的相应部分相匹配。

The full name of a subtest or sub-benchmark is a slash-separated list of its name and the names of all of its parents, starting with the top-level. The name is the corresponding function name for top-level tests and benchmarks, and the first argument to `Run` otherwise. To avoid display and parsing issues, a name is sanitized by replacing spaces with underscores and escaping non-printable characters. The same sanitizing is applied to the regular expressions passed to the `-run` or `-bench` flags.

子测试或子基准的全名是一个斜线分隔的列表，包括它的名字和它所有父级的名字，从顶级开始。对于顶层测试和基准，该名称是相应的函数名称，否则就是Run的第一个参数。为了避免显示和解析问题，通过用下划线替换空格和转义不可打印的字符，对名称进行消毒处理。同样的处理方法也适用于传递给-run或-bench标志的正则表达式。

A few examples:

举几个例子：

Run tests that use a timezone in Europe:

运行使用欧洲时区的测试：

```shell linenums="1"
$ go test -run=TestTime/"in Europe"
--- FAIL: TestTime (0.00s)
    --- FAIL: TestTime/12:31_in_Europe/Zuri (0.00s)
        time_test.go:85: could not load location
```

Run only tests for times after noon:

只运行中午以后的时间的测试：

```shell linenums="1"
$ go test -run=Time/12:[0-9] -v
=== RUN   TestTime
=== RUN   TestTime/12:31_in_Europe/Zuri
=== RUN   TestTime/12:31_in_America/New_York
--- FAIL: TestTime (0.00s)
    --- FAIL: TestTime/12:31_in_Europe/Zuri (0.00s)
        time_test.go:85: could not load location
    --- FAIL: TestTime/12:31_in_America/New_York (0.00s)
        time_test.go:89: got 07:31; want 7:31
```

Perhaps a bit surprising, using `-run=TestTime/New_York` won’t match any tests. This is because the slash present in the location names is treated as a separator as well. Instead use:

也许有点令人惊讶，使用-run=TestTime/New_York不会匹配任何测试。这是因为位置名称中的斜线也被当作分隔符来处理。相反，使用：

```shell linenums="1"
$ go test -run=Time//New_York
--- FAIL: TestTime (0.00s)
    --- FAIL: TestTime/12:31_in_America/New_York (0.00s)
        time_test.go:88: got 07:31; want 7:31
```

Note the `//` in the string passed to `-run`. The `/` in time zone name `America/New_York` is handled as if it were a separator resulting from a subtest. The first regular expression of the pattern (`TestTime`) matches the top-level test. The second regular expression (the empty string) matches anything, in this case the time and the continent part of the location. The third regular expression (`New_York`) matches the city part of the location.

注意传递给-run的字符串中的//。时区名称America/New_York中的/被处理，就像它是一个子测试产生的分隔符。模式中的第一个正则表达式（TestTime）与顶层测试相匹配。第二个正则表达式（空字符串）匹配任何东西，在这种情况下是时间和地点的大陆部分。第三个正则表达式（New_York）匹配位置的城市部分。

Treating slashes in names as separators allows the user to refactor hierarchies of tests without the need to change the naming. It also simplifies the escaping rules. The user should escape slashes in names, for instance by replacing them with backslashes, if this poses a problem.

将名称中的斜线视为分隔符，允许用户重构测试的层次，而不需要改变命名。它还简化了转义规则。用户应该转义名字中的斜线，例如用反斜线代替，如果这构成一个问题。

A unique sequence number is appended to test names that are not unique. So one could just pass an empty string to `Run` if there is no obvious naming scheme for subtests and the subtests can easily be identified by their sequence number.

一个唯一的序列号被附加到非唯一的测试名称上。因此，如果没有明显的子测试的命名方案，并且子测试可以很容易地通过其序列号来识别，那么可以直接传递一个空字符串给Run。

## Setup and Tear-down 设置和拆分

Subtests and sub-benchmarks can be used to manage common setup and tear-down code:

子测试和子基准可以用来管理常见的设置和拆解代码：

```go linenums="1"
func TestFoo(t *testing.T) {
    // <setup code>
    t.Run("A=1", func(t *testing.T) { ... })
    t.Run("A=2", func(t *testing.T) { ... })
    t.Run("B=1", func(t *testing.T) {
        if !test(foo{B:1}) {
            t.Fail()
        }
    })
    // <tear-down code>
}
```

The setup and tear-down code will run if any of the enclosed subtests are run and will run at most once. This applies even if any of the subtests calls `Skip`, `Fail`, or `Fatal`.

如果任何一个封闭的子测试被运行，设置和拆分代码将被运行，并且最多运行一次。即使任何一个子测试调用Skip、Fail或Fatal，这也适用。

## Control of Parallelism 平行性的控制

Subtests allow fine-grained control over parallelism. To understand how to use subtests in the way it is important to understand the semantics of parallel tests.

子测试允许对并行性进行细粒度的控制。为了理解如何使用子测试的方式，理解并行测试的语义很重要。

Each test is associated with a test function. A test is called a parallel test if its test function calls the Parallel method on its instance of `testing.T`. A parallel test never runs concurrently with a sequential test and its execution is suspended until its calling test function, that of the parent test, has returned. The `-parallel` flag defines the maximum number of parallel tests that can run in parallel.

每个测试都与一个测试函数相关。如果一个测试的测试函数在其testing.T的实例上调用Parallel方法，则该测试被称为并行测试。并行测试从不与顺序测试同时运行，其执行被暂停，直到其调用的测试函数，即父测试的测试函数返回。-parallel标志定义了可以并行运行的最大数量的并行测试。

A test blocks until its test function returns and all of its subtests have completed. This means that the parallel tests that are run by a sequential test will complete before any other consecutive sequential test is run.

一个测试阻塞，直到它的测试函数返回，并且所有的子测试都完成。这意味着由一个顺序测试运行的并行测试将在任何其他连续的顺序测试运行之前完成。

This behavior is identical for tests created by `Run` and top-level tests. In fact, under the hood top-level tests are implemented as subtests of a hidden master test.

这种行为对于由Run和顶层测试创建的测试是相同的。事实上，在引擎盖下，顶级测试被实现为一个隐藏的主测试的子测试。

### Run a group of tests in parallel 并行运行一组测试

The above semantics allows for running a group of tests in parallel with each other but not with other parallel tests:

上述语义允许并行运行一组测试，但不能与其他并行测试并行：

```go linenums="1"
func TestGroupedParallel(t *testing.T) {
    for _, tc := range testCases {
        tc := tc // capture range variable
        t.Run(tc.Name, func(t *testing.T) {
            t.Parallel()
            if got := foo(tc.in); got != tc.out {
                t.Errorf("got %v; want %v", got, tc.out)
            }
            ...
        })
    }
}
```

The outer test will not complete until all parallel tests started by `Run` have completed. As a result, no other parallel tests can run in parallel to these parallel tests.

在所有由Run启动的并行测试完成之前，外部测试将不会完成。因此，没有其他并行测试可以与这些并行测试并行运行。

Note that we need to capture the range variable to ensure that `tc` gets bound to the correct instance.

注意，我们需要捕获范围变量以确保tc被绑定到正确的实例。

### Cleaning up after a group of parallel tests 在一组并行测试之后进行清理

In the previous example we used the semantics to wait on a group of parallel tests to complete before commencing other tests. The same technique can be used to clean up after a group of parallel tests that share common resources:

在前面的例子中，我们使用语义来等待一组并行测试的完成，然后再开始其他测试。同样的技术可以用来清理一组共享资源的并行测试：

```go linenums="1"
func TestTeardownParallel(t *testing.T) {
    // <setup code>
    // This Run will not return until its parallel subtests complete.
    t.Run("group", func(t *testing.T) {
        t.Run("Test1", parallelTest1)
        t.Run("Test2", parallelTest2)
        t.Run("Test3", parallelTest3)
    })
    // <tear-down code>
}
```

The behavior of waiting on a group of parallel tests is identical to that of the previous example.

在一组并行测试上的等待行为与前面的例子相同。

## Conclusion 总结

Go 1.7’s addition of subtests and sub-benchmarks allows you to write structured tests and benchmarks in a natural way that blends nicely into the existing tools. One way to think about this is that earlier versions of the testing package had a 1-level hierarchy: the package-level test was structured as a set of individual tests and benchmarks. Now that structure has been extended to those individual tests and benchmarks, recursively. In fact, in the implementation, the top-level tests and benchmarks are tracked as if they were subtests and sub-benchmarks of an implicit master test and benchmark: the treatment really is the same at all levels.

Go 1.7 增加了子测试和子基准，允许你以一种自然的方式编写结构化的测试和基准，并很好地融入现有的工具。一种思考方式是，早期版本的测试包有一个1级的层次结构：包级测试被结构化为一组单独的测试和基准。现在，这个结构已经扩展到那些单独的测试和基准，递归地。事实上，在实现中，顶级测试和基准被跟踪，就像它们是隐含的主测试和基准的子测试和子基准一样：所有级别的处理都是一样的。

The ability for tests to define this structure enables fine-grained execution of specific test cases, shared setup and teardown, and better control over test parallelism. We are excited to see what other uses people find. Enjoy.

测试定义这种结构的能力使特定测试用例的细粒度执行、共享设置和拆除以及对测试并行性的更好控制。我们很高兴看到人们发现的其他用途。请欣赏。
