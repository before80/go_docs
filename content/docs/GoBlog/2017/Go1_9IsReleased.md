+++
title = "go 1.9发布了"
weight = 4
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go 1.9 is released - go 1.9发布了

https://go.dev/blog/go1.9

Francesc Campoy
24 August 2017

Today the Go team is happy to announce the release of Go 1.9. You can get it from the [download page](https://go.dev/dl/). There are many changes to the language, standard library, runtime, and tooling. This post covers the most significant visible ones. Most of the engineering effort put into this release went to improvements of the runtime and tooling, which makes for a less exciting announcement, but nonetheless a great release.

今天Go团队很高兴地宣布Go 1.9的发布。您可以从下载页面获得它。语言、标准库、运行时和工具都有很多变化。这篇文章涵盖了最重要的可见变化。这个版本的大部分工程努力都用于改进运行时和工具，这使得公告不那么激动人心，但仍是一个伟大的版本。

The most important change to the language is the introduction of type aliases: a feature created to support gradual code repair. A type alias declaration has the form:

该语言最重要的变化是引入了类型别名：这是一个为支持逐步修复代码而创建的功能。一个类型别名声明的形式是：

```go linenums="1"
type T1 = T2
```

This declaration introduces an alias name `T1` for the type `T2`, in the same way that `byte` has always been an alias for `uint8`. The [type alias design document](https://go.dev/design/18130-type-alias) and [an article on refactoring](https://go.dev/talks/2016/refactor.article) cover this addition in more detail.

这个声明为类型T2引入了一个别名T1，就像byte一直是uint8的别名一样。类型别名设计文档和一篇关于重构的文章更详细地介绍了这一补充。

The new [math/bits](https://go.dev/pkg/math/bits) package provides bit counting and manipulation functions for unsigned integers, implemented by special CPU instructions when possible. For example, on x86-64 systems, `bits.TrailingZeros(x)` uses the [BSF](https://pdos.csail.mit.edu/6.828/2010/readings/i386/BSF.htm) instruction.

新的math/bits包为无符号整数提供了位数计算和操作功能，在可能的情况下由特殊的CPU指令实现。例如，在x86-64系统上，bits.TrailingZeros(x)使用BSF指令。

The `sync` package has added a new [Map](https://go.dev/pkg/sync#Map) type, safe for concurrent access. You can read more about it from its documentation and learn more about why it was created from this [GopherCon 2017 lightning talk](https://www.youtube.com/watch?v=C1EtfDnsdDs) ([slides](https://github.com/gophercon/2017-talks/blob/master/lightningtalks/BryanCMills-AnOverviewOfSyncMap/An Overview of sync.Map.pdf)). It is not a general replacement for Go’s map type; please see the documentation to learn when it should be used.

Sync包增加了一个新的Map类型，对并发访问是安全的。您可以从它的文档中读到更多关于它的信息，并从这个GopherCon 2017的闪电演讲（幻灯片）中了解更多关于它创建的原因。它不是Go的地图类型的一般替代物；请看文档以了解何时应该使用它。

The `testing` package also has an addition. The new `Helper` method, added to both [testing.T](https://go.dev/pkg/testing#T.Helper) and [testing.B](https://go.dev/pkg/testing#B.Helper), marks the calling function as a test helper function. When the testing package prints file and line information, it shows the location of the call to a helper function instead of a line in the helper function itself.

测试包也有一个补充。新的Helper方法，添加到testing.T和testing.B中，将调用的函数标记为测试辅助函数。当测试包打印文件和行信息时，它显示调用辅助函数的位置，而不是辅助函数本身的行。

For example, consider this test:

例如，考虑这个测试：

```go linenums="1"
package p

import "testing"

func failure(t *testing.T) {
    t.Helper() // This call silences this function in error reports.
    t.Fatal("failure")
}

func Test(t *testing.T) {
    failure(t)
}
```

Because `failure` identifies itself as a test helper, the error message printed during `Test` will indicate line 11, where `failure` is called, instead of line 7, where `failure` calls `t.Fatal`.

因为failure将自己识别为一个测试助手，所以在Test过程中打印的错误信息将显示第11行，即failure被调用的地方，而不是第7行，即failure调用t.Fatal。

The `time` package now transparently tracks monotonic time in each `Time` value, making computing durations between two `Time` values a safe operation in the presence of wall clock adjustments. For example, this code now computes the right elapsed time even across a leap second clock reset:

时间包现在可以透明地跟踪每个时间值中的单调时间，使计算两个时间值之间的持续时间在挂钟调整的情况下成为安全操作。例如，这段代码现在可以计算出正确的经过时间，即使在闰秒时钟重置的情况下：

```go linenums="1"
start := time.Now()
f()
elapsed := time.Since(start)
```

See the [package docs](http://beta.golang.org/pkg/time/#hdr-Monotonic_Clocks) and [design document](https://github.com/golang/proposal/blob/master/design/12914-monotonic.md) for details.

详情请见包的文档和设计文件。

Finally, as part of the efforts to make the Go compiler faster, Go 1.9 compiles functions in a package concurrently.

最后，作为使Go编译器更快的努力的一部分，Go 1.9对包中的函数进行并发编译。

Go 1.9 includes many more additions, improvements, and fixes. Find the complete set of changes, and more information about the improvements listed above, in the [Go 1.9 release notes](https://go.dev/doc/go1.9).

Go 1.9 包括了更多的新增内容、改进和修正。在Go 1.9发布说明中可以找到完整的变化，以及有关上述改进的更多信息。

To celebrate the release, Go User Groups around the world are holding [release parties](https://github.com/golang/cowg/blob/master/events/2017-08-go1.9-release-party.md).

为了庆祝该版本的发布，世界各地的Go用户组正在举行发布派对。
