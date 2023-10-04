+++
title = "修复Go 1.22中的for循环问题"
date = 2023-10-04T14:33:00+08:00
type = "docs"
weight = 80
description = ""
isCJKLanguage = true
draft = false

+++

# Fixing For Loops in Go 1.22 - 修复Go 1.22中的for循环问题

> 原文：[https://go.dev/blog/loopvar-preview](https://go.dev/blog/loopvar-preview)

David Chase and Russ Cox
19 September 2023

大卫·查斯（David Chase）和Russ Cox
2023年9月19日

Go 1.21 includes a preview of a change to `for` loop scoping that we plan to ship in Go 1.22, removing one of the most common Go mistakes.

​	Go 1.21中包含了一个 `for` 循环作用域的预览，我们计划在Go 1.22中发布该变更，从而消除了最常见的Go错误之一。

## 问题 The Problem

If you’ve written any amount of Go code, you’ve probably made the mistake of keeping a reference to a loop variable past the end of its iteration, at which point it takes on a new value that you didn’t want. For example, consider this program:

​	如果您编写了任意数量的Go代码，您可能犯过一个错误，即在迭代结束后保留对循环变量的引用，此时它会取一个您不希望的新值。例如，考虑以下程序：

```
func main() {
    done := make(chan bool)

    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }

    // wait for all goroutines to complete before exiting
    // 等待所有goroutine完成后退出
    for _ = range values {
        <-done
    }
}
```

The three created goroutines are all printing the same variable `v`, so they usually print “c”, “c”, “c”, instead of printing “a”, “b”, and “c” in some order.

​	这三个创建的goroutine都打印相同的变量 `v` ，因此它们通常会打印“c”，“c”，“c”，而不是以某种顺序打印“a”，“b”和“c”。

The [Go FAQ entry “What happens with closures running as goroutines?”](https://go.dev/doc/faq#closures_and_goroutines), gives this example and remarks “Some confusion may arise when using closures with concurrency.”

​	[Go FAQ条目“使用goroutine运行的闭包会有什么问题？”](https://go.dev/doc/faq#closures_and_goroutines)提供了这个示例，并指出“在使用并发时可能会出现一些困惑”。

Although concurrency is often involved, it need not be. This example has the same problem but no goroutines:

​	尽管通常涉及并发，但并不一定如此。这个示例也存在相同的问题，但没有goroutine：

```
func main() {
    var prints []func()
    for i := 1; i <= 3; i++ {
        prints = append(prints, func() { fmt.Println(i) })
    }
    for _, print := range prints {
        print()
    }
}
```

This kind of mistake has caused production problems at many companies, including a [publicly documented issue at Lets Encrypt](https://bugzilla.mozilla.org/show_bug.cgi?id=1619047). In that instance, the accidental capture of the loop variable was spread across multiple functions and much more difficult to notice:

​	这种错误已经导致许多公司的生产问题，包括[Let's Encrypt的一个公开记录的问题](https://bugzilla.mozilla.org/show_bug.cgi?id=1619047)。在该实例中，循环变量的意外捕获分布在多个函数中，更难以注意到：

```
// authz2ModelMapToPB converts a mapping of domain name to authz2Models into a
// protobuf authorizations map
// authz2ModelMapToPB将域名到authz2Models的映射转换为protobuf授权映射
func authz2ModelMapToPB(m map[string]authz2Model) (*sapb.Authorizations, error) {
    resp := &sapb.Authorizations{}
    for k, v := range m {
        // Make a copy of k because it will be reassigned with each loop.
        // 复制k，因为它将在每次循环中重新赋值。
        kCopy := k
        authzPB, err := modelToAuthzPB(&v)
        if err != nil {
            return nil, err
        }
        resp.Authz = append(resp.Authz, &sapb.Authorizations_MapElement{
            Domain: &kCopy,
            Authz: authzPB,
        })
    }
    return resp, nil
}
```

The author of this code clearly understood the general problem, because they made a copy of `k`, but it turns out `modelToAuthzPB` used pointers to fields in `v` when constructing its result, so the loop also needed to make a copy of `v`.

​	这段代码的作者显然理解了这个普遍问题，因为他们复制了 `k` ，但事实证明， `modelToAuthzPB` 在构造其结果时使用了 `v` 中字段的指针，因此循环还需要复制 `v` 。

Tools have been written to identify these mistakes, but it is hard to analyze whether references to a variable outlive its iteration or not. These tools must choose between false negatives and false positives. The `loopclosure` analyzer used by `go vet` and `gopls` opts for false negatives, only reporting when it is sure there is a problem but missing others. Other checkers opt for false positives, accusing correct code of being incorrect. We ran an analysis of commits adding `x := x` lines in open-source Go code, expecting to find bug fixes. Instead we found many unnecessary lines being added, suggesting instead that popular checkers have significant false positive rates, but developers add the lines anyway to keep the checkers happy.

​	已经编写了工具来识别这些错误，但很难分析变量的引用是否超出了其迭代的范围。这些工具必须在误报和漏报之间进行选择。 `go vet` 和 `gopls` 使用的 `loopclosure` 分析器选择了漏报，只有在确定存在问题时才报告，并且会错过其他情况。其他检查器则选择误报，将正确的代码指责为错误。我们对开源Go代码中添加了 `x := x` 行的提交进行了分析，希望找到bug修复。结果我们发现，添加了许多不必要的行，这表明流行的检查器存在相当高的误报率，但开发人员仍然添加这些行以使检查器保持良好的状态。

One pair of examples we found was particularly illuminating:

​	我们发现了一对特别有启发性的示例：

This diff was in one program:

​	一个程序中的差异是这样的：

```
     for _, informer := range c.informerMap {
+        informer := informer
         go informer.Run(stopCh)
     }
```

And this diff was in another program:

​	另一个程序中的差异是这样的：

```
     for _, a := range alarms {
+        a := a
         go a.Monitor(b)
     }
```

One of these two diffs is a bug fix; the other is an unnecessary change. You can’t tell which is which unless you know more about the types and functions involved.

​	这两个差异中，一个是bug修复，另一个是不必要的更改。除非您了解涉及的类型和函数的更多信息，否则无法确定哪个是哪个。

## 修复方法 The Fix

For Go 1.22, we plan to change `for` loops to make these variables have per-iteration scope instead of per-loop scope. This change will fix the examples above, so that they are no longer buggy Go programs; it will end the production problems caused by such mistakes; and it will remove the need for imprecise tools that prompt users to make unnecessary changes to their code.

​	对于Go 1.22，我们计划更改 `for` 循环，使这些变量具有每次迭代的作用域，而不是每次循环的作用域。这个变更将修复上面的示例，使它们不再是有错误的Go程序；它将解决由此类错误引起的生产问题；它还将消除不精确的工具，这些工具提示用户对其代码进行不必要的更改。

To ensure backwards compatibility with existing code, the new semantics will only apply in packages contained in modules that declare `go 1.22` or later in their `go.mod` files. This per-module decision provides developer control of a gradual update to the new semantics throughout a codebase. It is also possible to use `//go:build` lines to control the decision on a per-file basis.

​	为了确保与现有代码的向后兼容性，新的语义仅适用于在其 `go.mod` 文件中声明了 `go 1.22` 或更高版本的模块中包含的包。这种逐模块的决策为开发人员提供了在整个代码库中逐渐更新到新语义的控制。还可以使用 `//go:build` 行来在每个文件的基础上控制决策。

Old code will continue to mean exactly what it means today: the fix only applies to new or updated code. This will give developers control over when the semantics change in a particular package. As a consequence of our [forward compatibility work](https://go.dev/blog/toolchain), Go 1.21 will not attempt to compile code that declares `go 1.22` or later. We included a special case with the same effect in the point releases Go 1.20.8 and Go 1.19.13, so when Go 1.22 is released, code written depending on the new semantics will never be compiled with the old semantics, unless people are using very old, [unsupported Go versions](https://go.dev/doc/devel/release#policy).

​	旧代码将继续完全按照当前的含义进行解释：修复仅适用于新代码或更新的代码。这将使开发人员能够控制特定包中的语义何时发生变化。由于我们的[向前兼容性工作](https://go.dev/blog/toolchain)，Go 1.21将不会尝试编译声明了 `go 1.22` 或更高版本的代码。我们在Go 1.20.8和Go 1.19.13的点发布版本中包含了相同效果的特殊情况，因此当发布Go 1.22时，依赖新语义的代码将永远不会使用旧语义进行编译，除非人们使用非常旧的、[不受支持的Go版本](https://go.dev/doc/devel/release#policy)。

## 预览修复 Previewing The Fix

Go 1.21 includes a preview of the scoping change. If you compile your code with `GOEXPERIMENT=loopvar` set in your environment, then the new semantics are applied to all loops (ignoring the `go.mod` `go` lines). For example, to check whether your tests still pass with the new loop semantics applied to your package and all your dependencies:

​	Go 1.21中包含了这个作用域变更的预览。如果您在环境中设置了 `GOEXPERIMENT=loopvar` ，则新的语义将应用于所有循环（忽略 `go.mod` 中的 `go` 行）。例如，要检查在应用新的循环语义到包和所有依赖项后，您的测试是否仍然通过：

```
GOEXPERIMENT=loopvar go test
```

We patched our internal Go toolchain at Google to force this mode during all builds at the start of May 2023, and in the past four months we have had zero reports of any problems in production code.

​	我们在Google的内部Go工具链中修补了此模式，以便在2023年5月初的所有构建中强制使用此模式，在过去的四个月中，我们没有收到任何生产代码的问题报告。

You can also try test programs to better understand the semantics on the Go playground by including a `// GOEXPERIMENT=loopvar` comment at the top of the program, like in [this program](https://go.dev/play/p/YchKkkA1ETH). (This comment only applies in the Go playground.)

​	您还可以在Go playground中尝试测试程序，以更好地理解循环语义，只需在程序顶部包含一个 `// GOEXPERIMENT=loopvar` 注释，就像在[这个程序](https://go.dev/play/p/YchKkkA1ETH)中一样（此注释仅适用于Go playground）。

## 修复有错误的测试 Fixing Buggy Tests

Although we’ve had no production problems, to prepare for that switch, we did have to correct many buggy tests that were not testing what they thought they were, like this:

​	尽管我们没有遇到生产问题，但为了准备进行此切换，我们确实必须纠正许多有错误的测试，这些测试并没有测试它们认为的内容，就像这样的测试：

```
func TestAllEvenBuggy(t *testing.T) {
    testCases := []int{1, 2, 4, 6}
    for _, v := range testCases {
        t.Run("sub", func(t *testing.T) {
            t.Parallel()
            if v&1 != 0 {
                t.Fatal("odd v", v)
            }
        })
    }
}
```

In Go 1.21, this test passes because `t.Parallel` blocks each subtest until the entire loop has finished and then runs all the subtests in parallel. When the loop has finished, `v` is always 6, so the subtests all check that 6 is even, so the test passes. Of course, this test really should fail, because 1 is not even. Fixing for loops exposes this kind of buggy test.

​	在Go 1.21中，这个测试通过，因为 `t.Parallel` 会阻塞每个子测试，直到整个循环完成，然后并行运行所有子测试。当循环完成时， `v` 始终是6，因此所有子测试都检查6是否为偶数，因此测试通过。当然，这个测试实际上应该失败，因为1不是偶数。修复for循环将暴露出这种错误的测试。

To help prepare for this kind of discovery, we improved the precision of the `loopclosure` analyzer in Go 1.21 so that it can identify and report this problem. You can see the report [in this program](https://go.dev/play/p/WkJkgXRXg0m) on the Go playground. If `go vet` is reporting this kind of problem in your own tests, fixing them will prepare you better for Go 1.22.

​	为了帮助准备这种发现，我们在Go 1.21中提高了 `loopclosure` 分析器的精度，以便在应用新语义时可以识别和报告此问题。您可以在Go playground中查看[此程序](https://go.dev/play/p/WkJkgXRXg0m)中的报告。如果 `go vet` 在您自己的测试中报告了此类问题，请修复它们，这将更好地为Go 1.22做准备。

If you run into other problems, [the FAQ](https://github.com/golang/go/wiki/LoopvarExperiment#my-test-fails-with-the-change-how-can-i-debug-it) has links to examples and details about using a tool we’ve written to identify which specific loop is causing a test failure when the new semantics are applied.

​	如果遇到其他问题，[FAQ](https://github.com/golang/go/wiki/LoopvarExperiment#my-test-fails-with-the-change-how-can-i-debug-it)中提供了链接和详细信息，介绍了我们编写的一个工具，用于识别在应用新语义时导致测试失败的特定循环。

## 更多信息 More Information

For more information about the change, see the [design document](https://go.googlesource.com/proposal/+/master/design/60078-loopvar.md) and the [FAQ](https://go.dev/wiki/LoopvarExperiment).

​	有关此更改的更多信息，请参阅[设计文档](https://go.googlesource.com/proposal/+/master/design/60078-loopvar.md)和[常见问题解答](https://go.dev/wiki/LoopvarExperiment)。