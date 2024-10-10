+++
title = "向后兼容性、Go 1.21和Go 2"
date = 2023-08-21T15:02:20+08:00
weight = 88
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Backward Compatibility, Go 1.21, and Go 2 - 向后兼容性、Go 1.21和Go 2

> 原文：[https://go.dev/blog/compat](https://go.dev/blog/compat)
>

Russ Cox
14 August 2023

Russ Cox 2023年8月14日



Go 1.21 includes new features to improve compatibility. Before you stop reading, I know that sounds boring. But boring can be good. Back in the early days of Go 1, Go was exciting and full of surprises. Each week we cut a new snapshot release and everyone got to roll the dice to see what we’d changed and how their programs would break. We released Go 1 and its compatibility promise to remove the excitement, so that new releases of Go would be boring.

​	Go 1.21版本包含了新功能，以提高兼容性。在您停下阅读之前，我知道这听起来很无聊。但无聊可能是好事。在Go 1的早期，Go是令人兴奋且充满惊喜的。每周我们都会发布一个新的快照版本，每个人都可以摇一摇骰子，看看我们做了哪些改变，以及他们的程序会如何崩溃。我们发布了Go 1和它的兼容性承诺，以消除兴奋感，使得新版本的Go变得乏味。

Boring is good. Boring is stable. Boring means being able to focus on your work, not on what’s different about Go. This post is about the important work we shipped in Go 1.21 to keep Go boring.

​	乏味是好事。乏味是稳定的。乏味意味着能够专注于工作，而不是关注Go的不同之处。这篇文章涉及了我们在Go 1.21中发布的重要工作，以保持Go的稳定性。

## Go 1的兼容性 Go 1 Compatibility

We’ve been focused on compatibility for over a decade. For Go 1, back in 2012, we published a document titled “[Go 1 and the Future of Go Programs](https://go.dev/doc/go1compat)” that lays out a very clear intention:

​	我们已经专注于兼容性十多年了。在2012年的Go 1时，我们发布了一份题为“[Go 1和Go程序的未来](https://go.dev/doc/go1compat)”的文档，其中明确表达了一个非常明确的意图：

> It is intended that programs written to the Go 1 specification will continue to compile and run correctly, unchanged, over the lifetime of that specification. … Go programs that work today should continue to work even as future releases of Go 1 arise.
>
> ​	Go 1规范的程序在其规范的生命周期内，将继续编译和运行，不变，而不会出现错误。...今天有效的Go程序，将继续在未来的Go 1版本中继续有效。

There are a few qualifications to that. First, compatibility means source compatibility. When you update to a new version of Go, you do have to recompile your code. Second, we can add new APIs, but not in a way that breaks existing code.

​	对此有一些限制。首先，兼容性指的是源代码兼容性。当您升级到新版本的Go时，确实需要重新编译您的代码。其次，我们可以添加新的API，但不能以一种破坏现有代码的方式添加。

The end of the document warns, “[It] is impossible to guarantee that no future change will break any program.” Then it lays out a number of reasons why programs might still break.

​	文档的结尾警告说：“[不可能保证任何未来的更改不会破坏任何程序。]”，然后它列出了一些程序可能仍然会出现问题的原因。

For example, it makes sense that if your program depends on a buggy behavior and we fix the bug, your program will break. But we try very hard to break as little as possible and keep Go boring. There are two main approaches we’ve used so far: API checking and testing.

​	例如，如果您的程序依赖于错误的行为并且我们修复了错误，那么您的程序将会崩溃。但我们非常努力地尽量减少破坏，保持Go的稳定性。迄今为止，我们主要采用了两种方法：API检查和测试。

## API检查 API Checking

Perhaps the clearest fact about compatibility is that we can’t take away API, or else programs using it will break.

​	兼容性最清楚的事实可能是，我们不能删除API，否则使用它的程序将会崩溃。

For example, here’s a program someone has written that we can’t break:

​	例如，这是一个某人编写的我们不能破坏的程序：

```go
package main

import "os"

func main() {
    os.Stdout.WriteString("hello, world\n")
}
```

We can’t remove the package `os`; we can’t remove the global variable `os.Stdout`, which is an `*os.File`; and we also can’t remove the `os.File` method `WriteString`. It should be clear that removing any of those would break this program.

​	我们不能删除包`os`；我们不能删除全局变量`os.Stdout`，它是一个`*os.File`；我们也不能删除`os.File`的方法`WriteString`。应该清楚，删除其中任何一个都会破坏这个程序。

It’s perhaps less clear that we can’t change the type of `os.Stdout` at all. Suppose we want to make it an interface with the same methods. The program we just saw wouldn’t break, but this one would:

​	或许不太清楚的是，我们根本不能更改`os.Stdout`的类型。假设我们想将其更改为具有相同方法的接口。刚刚看到的程序不会崩溃，但是这个会：

```go
package main

import "os"

func main() {
    greet(os.Stdout)
}

func greet(f *os.File) {
    f.WriteString(“hello, world\n”)
}
```

This program passes `os.Stdout` to a function named `greet` that requires an argument of type `*os.File`. So changing `os.Stdout` to an interface will break this program.

​	这个程序将`os.Stdout`传递给名为`greet`的函数，该函数需要一个类型为`*os.File`的参数。因此，将`os.Stdout`更改为接口将会破坏这个程序。

To help us as we develop Go, we use a tool that maintains a list of each package’s exported API in files separate from the actual packages:

​	为了在开发Go时帮助我们，我们使用一种工具，该工具将每个包的导出API列表维护在实际包之外的文件中：

```bash
% cat go/api/go1.21.txt
pkg bytes, func ContainsFunc([]uint8, func(int32) bool) bool #54386
pkg bytes, method (*Buffer) AvailableBuffer() []uint8 #53685
pkg bytes, method (*Buffer) Available() int #53685
pkg cmp, func Compare[$0 Ordered]($0, $0) int #59488
pkg cmp, func Less[$0 Ordered]($0, $0) bool #59488
pkg cmp, type Ordered interface {} #59488
pkg context, func AfterFunc(Context, func()) func() bool #57928
pkg context, func WithDeadlineCause(Context, time.Time, error) (Context, CancelFunc) #56661
pkg context, func WithoutCancel(Context) Context #40221
pkg context, func WithTimeoutCause(Context, time.Duration, error) (Context, CancelFunc) #56661
```

One of our standard tests checks that the actual package APIs match those files. If we add new API to a package, the test breaks unless we add it to the API files. And if we change or remove API, the test breaks too. This helps us avoid mistakes. However, a tool like this only finds a certain class of problems, namely API changes and removals. There are other ways to make incompatible changes to Go.

​	我们的标准测试之一是检查实际包API是否与这些文件匹配。如果我们向包中添加了新的API，则该测试会失败，除非我们将其添加到API文件中。如果更改或删除了API，测试也会失败。这有助于我们避免错误。然而，像这样的工具只能找到某种类别的问题，即API的更改和删除。还有其他方法可以对Go进行不兼容的更改。

That leads us to the second approach we use to keep Go boring: testing.

​	这使我们转向了第二种保持Go稳定的方法：测试。

## 测试 Testing

The most effective way to find unexpected incompatibilities is to run existing tests against the development version of the next Go release. We test the development version of Go against all of Google’s internal Go code on a rolling basis. When tests are passing, we install that commit as Google’s production Go toolchain.

​	发现意外的不兼容性最有效的方法是针对下一个Go版本的开发版本运行现有的测试。我们在滚动的基础上对Google内部的所有Go代码进行开发版本的测试。当测试通过时，我们将该提交安装为Google的生产Go工具链。

If a change breaks tests inside Google, we assume it will also break tests outside Google, and we look for ways to reduce the impact. Most of the time, we roll back the change entirely or find a way to rewrite it so that it doesn’t break any programs. Sometimes, however, we conclude that the change is important to make and “compatible” even though it does break some programs. In that case, we still work to reduce the impact as much as possible, and then we document the potential problem in the release notes.

​	如果更改在Google内部破坏了测试，我们会假设它也会在Google外部破坏了测试，然后我们会寻找减少影响的方法。大多数情况下，我们会完全回退更改，或者找到一种方法来重写它，以便它不会破坏任何程序。然而，有时我们会得出结论，该更改是重要的并且“兼容的”，即使它确实破坏了一些程序。在这种情况下，我们仍然会尽量减少影响，然后在发布说明中记录潜在的问题。

Here are two examples of that kind of subtle compatibility problems we found by testing Go inside Google but still included in Go 1.1.

​	以下是我们在Google内部测试Go时发现的这种类型的微妙兼容性问题的两个示例，尽管它们仍然包含在Go 1.1中。

## 结构体字面量和新字段 Struct Literals and New Fields

Here is some code that runs fine in Go 1:

​	以下是在Go 1中正常运行的代码：

```go
package main

import "net"

var myAddr = &net.TCPAddr{
    net.IPv4(18, 26, 4, 9),
    80,
}
```

Package `main` declares a global variable `myAddr`, which is a composite literal of type `net.TCPAddr`. In Go 1, package `net` defines the type `TCPAddr` as a struct with two fields, `IP` and `Port`. Those match the fields in the composite literal, so the program compiles.

​	`main`包声明了一个全局变量`myAddr`，它是类型为`net.TCPAddr`的组合字面量。在Go 1中，`net`包将类型`TCPAddr`定义为具有两个字段`IP`和`Port`的结构体。这些字段与组合字面量中的字段匹配，因此程序可以编译。

In Go 1.1, the program stopped compiling, with a compiler error that said “too few initializers in struct literal.” The problem is that we added a third field, `Zone`, to `net.TCPAddr`, and this program is missing the value for that third field. The fix is to rewrite the program using tagged literals, so that it builds in both versions of Go:

​	在Go 1.1中，该程序停止了编译，出现了编译器错误，错误消息为“struct字面量中的初始值太少”。问题在于，我们向`net.TCPAddr`添加了第三个字段`Zone`，而这个程序缺少了第三个字段的值。解决方法是使用标记字面量重新编写程序，以便在Go的两个版本中都能构建：

```go
var myAddr = &net.TCPAddr{
    IP:   net.IPv4(18, 26, 4, 9),
    Port: 80,
}
```

Since this literal doesn’t specify a value for `Zone`, it will use the zero value (an empty string in this case).

​	由于此字面量未为`Zone`指定值，因此它将使用零值（在这种情况下为空字符串）。

This requirement to use composite literals for standard library structs is explicitly called out in the [compatibility document](https://go.dev/doc/go1compat), and `go vet` reports literals that need tags to ensure compatibility with later versions of Go. This problem was new enough in Go 1.1 to merit a short comment in the release notes. Nowadays we just mention the new field.

​	需要使用组合字面量来为标准库结构体添加标签的要求明确写在了[兼容性文档](https://go.dev/doc/go1compat)中，而`go vet`会报告需要标签的字面量，以确保与Go的后续版本兼容。这个问题在Go 1.1中是足够新的，以至于在发布说明中需要一个简短的注释。现在我们只提到了新字段。

## 时间精度 Time Precision

The second problem we found while testing Go 1.1 had nothing to do with APIs at all. It had to do with time.

​	在测试Go 1.1时，我们发现的第二个问题与API无关。它与时间有关。

Shortly after Go 1 was released, someone pointed out that [`time.Now`](https://go.dev/pkg/time/#Now) returned times with microsecond precision, but with some extra code, it could return times with nanosecond precision instead. That sounds good, right? More precision is better. So we made that change.

​	Go 1发布不久后，有人指出[`time.Now`](https://go.dev/pkg/time/#Now)返回的时间具有微秒精度，但通过一些额外的代码，它可以返回具有纳秒精度的时间。听起来很好，对吧？更高的精度更好。因此，我们进行了这个更改。

That broke a handful of tests inside Google that were schematically like this one:

​	这破坏了一些谷歌内部的测试，这些测试的模式类似于这个：

```go
func TestSaveTime(t *testing.T) {
    t1 := time.Now()
    save(t1)
    if t2 := load(); t2 != t1 {
        t.Fatalf("load() = %v, want %v", t1, t2)
    }
}
```

This code calls `time.Now` and then round-trips the result through `save` and `load` and expects to get the same time back. If `save` and `load` use a representation that only stores microsecond precision, that will work fine in Go 1 but fail in Go 1.1.

​	这段代码调用`time.Now`，然后通过`save`和`load`进行往返，期望获得相同的时间。如果`save`和`load`使用只存储微秒精度的表示，这在Go 1中可以正常工作，但在Go 1.1中会失败。

To help fix tests like this, we added [`Round`](https://go.dev/pkg/time/#Time.Round) and [`Truncate`](https://go.dev/pkg/time/#Time.Truncate) methods to discard unwanted precision, and in the release notes, we documented the possible problem and the new methods to help fix it.

​	为了帮助修复这种情况下的测试，我们添加了[`Round`](https://go.dev/pkg/time/#Time.Round)和[`Truncate`](https://go.dev/pkg/time/#Time.Truncate)方法，以丢弃不需要的精度，并在发布说明中记录了可能的问题和新方法以帮助解决它。

These examples show how testing finds different kinds of incompatibility than the API checks do. Of course, testing is not a complete guarantee of compatibility either, but it’s more complete than just API checks. There are many examples of problems we’ve found while testing that we decided did break the compatibility rules and rolled back before the release. The time precision change is an interesting example of something that broke programs but that we released anyway. We made the change because the improved precision was better and was allowed within the documented behavior of the function.

​	这些示例显示了测试如何发现不同类型的不兼容性，而不仅仅是API检查。当然，测试也不是完全保证兼容性，但比单纯的API检查更加完整。我们发现了许多问题，经过测试后，我们决定在发布之前回退了这些问题。时间精度的更改是一个有趣的例子，它破坏了程序，但我们还是发布了它。我们进行了更改，因为改进的精度更好，并且在函数的文档行为范围内是允许的。

This example shows that sometimes, despite significant effort and attention, there are times when changing Go means breaking Go programs. The changes are, strictly speaking, “compatible” in the sense of the Go 1 document, but they still break programs. Most of these compatibility issues can be placed in one of three categories: output changes, input changes, and protocol changes.

​	这个例子表明，有时候，尽管付出了很大的努力和注意力，改变Go仍然意味着破坏Go程序。从Go 1的文件的角度来看，这些更改从严格意义上说是“兼容”的，但它们仍然会破坏程序。这些兼容性问题大多可以归类为三类：输出变化、输入变化和协议变化。

## 输出变化 Output Changes

An output change happens when a function gives a different output than it used to, but the new output is just as correct as, or even more correct than, the old output. If existing code is written to expect only the old output, it will break. We just saw an example of this, with `time.Now` adding nanosecond precision.

​	输出变化发生在函数提供的输出与以前不同，但新的输出与旧的输出一样正确，甚至更正确。如果现有代码只期望旧的输出，那么它将会破坏。我们刚刚看到了一个例子，`time.Now`添加了纳秒精度。

**Sort.** Another example happened in Go 1.6, when we changed the implementation of sort to run about 10% faster. Here’s an example program that sorts a list of colors by length of name:

​	**排序。** 另一个例子发生在Go 1.6中，我们改变了sort的实现，使其运行速度提高了约10%。以下是一个按名称长度对颜色列表进行排序的示例程序：

```go
colors := strings.Fields(
    `black white red orange yellow green blue indigo violet`)
sort.Sort(ByLen(colors))
fmt.Println(colors)

Go 1.5:  [red blue green white black yellow orange indigo violet]
Go 1.6:  [red blue white green black orange yellow indigo violet]
```

Changing sort algorithms often changes how equal elements are ordered, and that happened here. Go 1.5 returned green, white, black, in that order. Go 1.6 returned white, green, black.

​	更改排序算法通常会改变相等元素的排序方式，这在这里发生了。Go 1.5返回了green、white、black，按照这个顺序。Go 1.6返回了white、green、black。

Sort is clearly allowed to return equal results in any order it likes, and this change made it 10% faster, which is great. But programs that expect a specific output will break. This is a good example of why compatibility is so hard. We don’t want to break programs, but we also don’t want to be locked in to undocumented implementation details.

​	显然，排序可以以任何顺序返回相等的结果，而这个改变使得它的速度提高了10%，这很好。但是期望特定输出的程序会中断。这是兼容性非常难以处理的一个很好的例子。我们不想破坏程序，但我们也不想被锁定在未记录的实现细节中。

**Compress/flate.** As another example, in Go 1.8 we improved `compress/flate` to produce smaller outputs, with roughly the same CPU and memory overheads. That sounds like a win-win, but it broke a project inside Google that needed reproducible archive builds: now they couldn’t reproduce their old archives. They forked `compress/flate` and `compress/gzip` to keep a copy of the old algorithm.

​	**Compress/flate。** 另一个例子，在Go 1.8中，我们改进了`compress/flate`以产生更小的输出，同时大致保持了CPU和内存开销。听起来是一个双赢的局面，但它破坏了谷歌内部的一个项目，该项目需要可复现的存档构建：现在他们无法再复现他们的旧存档。他们复制了`compress/flate`和`compress/gzip`以保留旧算法的副本。

We do a similar thing with the Go compiler, using a fork of the `sort` package ([and others](https://go.googlesource.com/go/+/go1.21.0/src/cmd/dist/buildtool.go#22)) so that the compiler produces the same results even when it is built using earlier versions of Go.

​	我们在Go编译器中使用了类似的方法，使用`sort`包的分支（[和其他](https://go.googlesource.com/go/+/go1.21.0/src/cmd/dist/buildtool.go#22)），以便即使使用早期版本的Go构建编译器时，它产生的结果也是相同的。

For output change incompatibilities like these, the best answer is to write programs and tests that accept any valid output, and to use these kinds of breakages as an opportunity to change your testing strategy, not just update the expected answers. If you need truly reproducible outputs, the next best answer is to fork the code to insulate yourself from changes, but remember that you’re also insulating yourself from bug fixes.

​	对于这些输出变化不兼容性，最好的答案是编写接受任何有效输出的程序和测试，并将这些类型的破坏视为更改测试策略的机会，而不仅仅是更新预期的答案。如果需要真正可复现的输出，下一个最佳答案是分叉代码以隔离自己免受更改的影响，但请记住，这也将隔离自己免受错误修复的影响。

## 输入变化 Input Changes

An input change happens when a function changes which inputs it accepts or how it processes them.

​	输入变化发生在函数更改其接受的输入或处理它们的方式时。

**ParseInt.** For example, Go 1.13 added support for underscores in large numbers for readability. Along with the language change, we made `strconv.ParseInt` accept the new syntax. This change didn’t break anything inside Google, but much later we heard from an external user whose code did break. Their program used numbers separated by underscores as a data format. It tried `ParseInt` first and only fell back to checking for underscores if `ParseInt` failed. When `ParseInt` stopped failing, the underscore-handling code stopped running.

​	**ParseInt。** 例如，Go 1.13增加了对大数中使用下划线以提高可读性的支持。伴随着这个语言更改，我们使`strconv.ParseInt`接受新的语法。这个更改没有在谷歌内部破坏任何东西，但在很久之后，我们从一个外部用户那里听说他们的代码已经破坏了。他们的程序将数字用下划线分隔作为数据格式。他们首先尝试了`ParseInt`，只有在`ParseInt`失败时才检查下划线。当`ParseInt`停止失败时，处理下划线的代码也停止运行。

**ParseIP.** As another example, Go’s `net.ParseIP`, followed the examples in early IP RFCs, which often showed decimal IP addresses with leading zeros. It read the IP address 18.032.4.011 address as 18.32.4.11, just with a few extra zeros. We found out much later that BSD-derived C libraries interpret leading zeros in IP addresses as starting an octal number: in those libraries, 18.032.4.011 means 18.26.4.9!

​	**ParseIP。** 另一个例子是，Go的`net.ParseIP`遵循早期IP RFC的示例，这些示例经常显示具有前导零的十进制IP地址。它将IP地址18.032.4.011解释为18.32.4.11，只是多了一些额外的零。我们后来发现，基于BSD的C库将IP地址中的前导零解释为开始一个八进制数：在这些库中，18.032.4.011意味着18.26.4.9！

This was a serious mismatch between Go and the rest of the world, but changing the meaning of leading zeros from one Go release to the next would be a serious mismatch too. It would be a huge incompatibility. In the end, we decided to change `net.ParseIP` in Go 1.17 to reject leading zeros entirely. This stricter parsing ensures that when Go and C both parse an IP address successfully, or when old and new Go versions do, they all agree about what it means.

​	这是Go和其他语言之间严重不匹配的情况，但从一代Go版本到下一代改变前导零的含义将是一个严重的不匹配。这将是一个巨大的不兼容性。最终，我们决定在Go 1.17中更改`net.ParseIP`，以完全拒绝前导零。这种更严格的解析确保当Go和C同时成功解析IP地址时，或者当新旧Go版本都这样做时，它们都会对它的含义达成一致。

This change didn’t break anything inside Google, but the Kubernetes team was concerned about saved configurations that might have parsed before but would stop parsing with Go 1.17. Addresses with leading zeros probably should be removed from those configs, since Go interprets them differently from essentially every other language, but that should happen on Kubernetes’s timeline, not Go’s. To avoid the semantic change, Kubernetes started using its own forked copy of the original `net.ParseIP`.

​	这个更改没有在谷歌内部破坏任何东西，但Kubernetes团队担心可能已经解析的配置将在Go 1.17中停止解析。具有前导零的地址可能应该从这些配置中删除，因为Go与几乎所有其他语言解释它们的方式不同，但这应该在Kubernetes的时间表内完成，而不是Go的。为了避免这个语义上的改变，Kubernetes开始使用自己分支的原始`net.ParseIP`。

The best response to input changes is to process user input by first validating the syntax you want to accept before parsing the values, but sometimes you need to fork the code instead.

​	对于输入变化，最好的反应是在解析值之前首先验证所需的语法，但有时您需要分叉代码。

## 协议变化 Protocol Changes

The final common kind of incompatibility is protocol changes. A protocol change is a change made to a package that ends up externally visible in the protocols a program uses to communicate with the external world. Almost any change can become externally visible in certain programs, as we saw with `ParseInt` and `ParseIP`, but protocol changes are externally visible in essentially all programs.

​	最后一种常见的不兼容性类型是协议变化。协议变化是对程序用于与外部世界通信的协议的包进行的更改，这些更改在外部是可见的。几乎任何更改在某些程序中都可能成为外部可见的，就像我们在`ParseInt`和`ParseIP`中看到的那样，但协议变化在几乎所有程序中都是外部可见的。

**HTTP/2.** A clear example of a protocol change is when Go 1.6 added automatic support for HTTP/2. Suppose a Go 1.5 client is connecting to an HTTP/2-capable server over a network with middleboxes that happen to break HTTP/2. Since Go 1.5 only uses HTTP/1.1, the program works fine. But then updating to Go 1.6 breaks the program, because Go 1.6 starts using HTTP/2, and in this context, HTTP/2 doesn’t work.

​	**HTTP/2。** 协议变化的一个明显例子是在Go 1.6中自动添加对HTTP/2的支持。假设Go 1.5客户端通过支持HTTP/2的网络连接到支持HTTP/2的服务器，但网络中的中间盒破坏了HTTP/2。由于Go 1.5只使用HTTP/1.1，程序可以正常工作。但是，更新到Go 1.6会破坏程序，因为Go 1.6开始使用HTTP/2，在这种情况下，HTTP/2不起作用。

Go aims to support modern protocols by default, but this example shows that enabling HTTP/2 can break programs through no fault of their own (nor any fault of Go’s). Developers in this situation could go back to using Go 1.5, but that’s not very satisfying. Instead, Go 1.6 documented the change in the release notes and made it straightforward to disable HTTP/2.

​	Go的目标是默认支持现代协议，但这个例子表明，通过HTTP/2可以无缘无故地破坏程序（既不是程序的错误，也不是Go的错误）。在这种情况下，开发者可以回到使用Go 1.5，但那并不是很令人满意。相反，Go 1.6在发布说明中记录了这个更改，并且使得禁用HTTP/2变得简单。

In fact, [Go 1.6 documented two ways](https://go.dev/doc/go1.6#http2) to disable HTTP/2: configure the `TLSNextProto` field explicitly using the package API, or set the GODEBUG environment variable:

​	实际上，[Go 1.6文档中记录了两种禁用HTTP/2的方法](https://go.dev/doc/go1.6#http2)：使用包API显式配置`TLSNextProto`字段，或设置GODEBUG环境变量：

```
GODEBUG=http2client=0 ./myprog
GODEBUG=http2server=0 ./myprog
GODEBUG=http2client=0,http2server=0 ./myprog
```

As we’ll see later, Go 1.21 generalizes this GODEBUG mechanism to make it a standard for all potentially breaking changes.

​	正如我们将在后面看到的，Go 1.21在这个GODEBUG机制中进行了扩展，使它成为所有可能破坏性变化的标准。

**SHA1.** Here’s a subtler example of a protocol change. No one should be using SHA1-based certificates for HTTPS anymore. Certificate authorities stopped issuing them in 2015, and all the major browsers stopped accepting them in 2017. In early 2020, Go 1.18 disabled support for them by default, with a GODEBUG setting to override that change. We also announced our intent to remove the GODEBUG setting in Go 1.19.

​	**SHA1。** 下面是协议变化的一个更微妙的例子。不应再使用基于SHA1的HTTPS证书。证书颁发机构在2015年停止发放这种证书，所有主要的浏览器在2017年停止接受这种证书。在2020年初，Go 1.18默认禁用了对它们的支持，并通过GODEBUG设置覆盖了这个更改。我们还宣布了我们计划在Go 1.19中删除GODEBUG设置。

The Kubernetes team let us know that some installations still use private SHA1 certificates. Putting aside the security questions, it’s not Kubernetes’s place to force these enterprises to upgrade their certificate infrastructure, and it would be extremely painful to fork `crypto/tls` and `net/http` to keep SHA1 support. Instead, we agreed to keep the override in place longer than we had planned, to create more time for an orderly transition. After all, we want to break as few programs as possible.

​	Kubernetes团队告诉我们，一些安装仍然使用私有SHA1证书。不考虑安全问题，Kubernetes不应强制这些企业升级他们的证书基础设施，而且为了保持SHA1支持，分叉`crypto/tls`和`net/http`会非常痛苦。相反，我们同意保持这个覆盖的时间比计划的时间要长，以便为有序的过渡创造更多的时间。毕竟，我们希望尽量减少破坏的程序数量。

## 在Go 1.21中扩展的GODEBUG支持 Expanded GODEBUG Support in Go 1.21

To improve backwards compatibility even in these subtle cases we’ve been examining, Go 1.21 expands and formalizes the use of GODEBUG.

​	为了在我们刚刚讨论过的这些微妙情况下改进向后兼容性，Go 1.21扩展并正式使用了GODEBUG。

To begin with, for any change that is permitted by Go 1 compatibility but still might break existing programs, we do all the work we just saw to understand potential compatibility problems, and we engineer the change to keep as many existing programs working as possible. For the remaining programs, the new approach is:

​	首先，对于任何允许Go 1兼容性的更改，但仍可能破坏现有程序的更改，我们会做我们刚刚看到的所有工作，以了解潜在的兼容性问题，并设计更改，以使尽可能多的现有程序继续工作。对于剩余的程序，新的方法是：

1. We will define a new GODEBUG setting that allows individual programs to opt out of the new behavior. A GODEBUG setting may not be added if doing so is infeasible, but that should be extremely rare.
2. 我们将定义一个新的GODEBUG设置，允许单个程序选择退出新的行为。如果添加GODEBUG设置是不可行的，那么就几乎不可能。
3. GODEBUG settings added for compatibility will be maintained for a minimum of two years (four Go releases). Some, such as `http2client` and `http2server`, will be maintained much longer, even indefinitely.
4. 为了保持兼容性，为兼容性添加的GODEBUG设置将至少维护两年（四个Go版本）。一些设置，比如`http2client`和`http2server`，将维护更长时间，甚至是无限期的。
5. When possible, each GODEBUG setting has an associated [`runtime/metrics`](https://go.dev/pkg/runtime/metrics/) counter named `/godebug/non-default-behavior/<name>:events` that counts the number of times a particular program’s behavior has changed based on a non-default value for that setting. For example, when `GODEBUG=http2client=0` is set, `/godebug/non-default-behavior/http2client:events` counts the number of HTTP transports that the program has configured without HTTP/2 support.
6. 在可能的情况下，每个GODEBUG设置都有一个与之关联的[`runtime/metrics`](https://go.dev/pkg/runtime/metrics/)计数器，名称为`/godebug/non-default-behavior/<name>:events`，用于计算特定程序的行为已更改的次数，这是基于该设置的非默认值。例如，当设置了`GODEBUG=http2client=0`时，`/godebug/non-default-behavior/http2client:events`计算程序配置为不支持HTTP/2的HTTP传输的数量。
7. A program’s GODEBUG settings are configured to match the Go version listed in the main package’s `go.mod` file. If your program’s `go.mod` file says `go 1.20` and you update to a Go 1.21 toolchain, any GODEBUG-controlled behaviors changed in Go 1.21 will retain their old Go 1.20 behavior until you change the `go.mod` to say `go 1.21`.
8. 一个程序的GODEBUG设置被配置为与主包的`go.mod`文件中列出的Go版本相匹配。如果程序的`go.mod`文件说`go 1.20`，而你更新到Go 1.21工具链，那么在你将`go.mod`更改为`go 1.21`之前，Go 1.21中更改的基于GODEBUG的行为将保留其旧的Go 1.20行为。
9. A program can change individual GODEBUG settings by using `//go:debug` lines in package `main`.
10. 一个程序可以通过在包`main`中使用`//go:debug`行来更改单个GODEBUG设置。
11. All GODEBUG settings are documented in a [single, central list](https://go.dev/doc/godebug#history) for easy reference.
12. 所有GODEBUG设置都在一个[单一的、中央的列表](https://go.dev/doc/godebug#history)中进行了文档化，以便于参考。

This approach means that each new version of Go should be the best possible implementation of older versions of Go, even preserving behaviors that were changed in compatible-but-breaking ways in later releases when compiling old code.

​	这种方法意味着每个新版本的Go都应该是旧版本Go的最佳实现，甚至在编译旧代码时也会保留兼容性但具有破坏性的行为。

For example, in Go 1.21, `panic(nil)` now causes a (non-nil) runtime panic, so that the result of [`recover`](https://go.dev/ref/spec/#Handling_panics) now reliably reports whether the current goroutine is panicking. This new behavior is controlled by a GODEBUG setting and therefore dependent on the main package’s `go.mod`’s `go` line: if it says `go 1.20` or earlier, `panic(nil)` is still allowed. If it says `go 1.21` or later, `panic(nil)` turns into a panic with a `runtime.PanicNilError`. And the version-based default can be overridden explicitly by adding a line like this to package main:

​	例如，在Go 1.21中，`panic(nil)`现在会引发一个（非nil的）运行时panic，以便[`recover`](https://go.dev/ref/spec/#Handling_panics)的结果现在可靠地报告当前goroutine是否正在发生panic。这个新的行为受一个GODEBUG设置的控制，因此依赖于主包的`go.mod`的`go`行：如果它说`go 1.20`或更早，那么允许`panic(nil)`。如果它说`go 1.21`或更高，`panic(nil)`将变成一个带有`runtime.PanicNilError`的panic。而且版本基础的默认值可以通过在包main中添加一行来明确覆盖，例如：

```go
//go:debug panicnil=1
```

This combination of features means that programs can update to newer toolchains while preserving the behaviors of the earlier toolchains they used, can apply finer-grained control over specific settings as needed, and can use production monitoring to understand which jobs make use of these non-default behaviors in practice. Combined, those should make rolling out new toolchains even smoother than in the past.

​	这些特性的结合意味着程序可以在保持旧工具链行为的情况下更新到更新的工具链，可以根据需要对特定设置进行更精细的控制，并可以使用生产监视来了解哪些作业实际上使用了这些非默认行为。综合起来，这些特性应该使得新工具链的推出比过去更加平稳。

See “[Go, Backwards Compatibility, and GODEBUG](https://go.dev/doc/godebug)” for more details.

​	有关更多细节，请参见“[Go、向后兼容性和GODEBUG](https://go.dev/doc/godebug)”。

## 关于Go 2的最新消息 - An Update on Go 2

In the quoted text from “[Go 1 and the Future of Go Programs](https://go.dev/doc/go1compat)” at the top of this post, the ellipsis hid the following qualifier:

​	在本文顶部引用的来自“[Go 1和Go程序的未来](https://go.dev/doc/go1compat)”的文字中，省略号隐藏了以下限定语：

> At some indefinite point, a Go 2 specification may arise, but until that time, [… all the compatibility details …].
>
> 在某个不确定的时间点，可能会出现一个Go 2的规范，但在那之前，[...所有的兼容性细节...]。

That raises an obvious question: when should we expect the Go 2 specification that breaks old Go 1 programs?

​	这引发了一个明显的问题：我们应该何时期望会出现破坏旧Go 1程序的Go 2规范？

The answer is never. Go 2, in the sense of breaking with the past and no longer compiling old programs, is never going to happen. Go 2 in the sense of being the major revision of Go 1 we started toward in 2017 has already happened.

​	答案是永远不会。从过去中分离出来，不再编译旧程序的意义上来看，Go 2永远不会发生。从2017年开始朝着的Go 1的主要修订版本的意义上来看，Go 2已经发生了。

There will not be a Go 2 that breaks Go 1 programs. Instead, we are going to double down on compatibility, which is far more valuable than any possible break with the past. In fact, we believe that prioritizing compatibility was the most important design decision we made for Go 1.

​	不会有一个破坏Go 1程序的Go 2。相反，我们将加倍注重兼容性，这比与过去的任何可能的分离更有价值。实际上，我们认为将兼容性放在首位是我们为Go 1做出的最重要的设计决策。

So what you will see over the next few years is plenty of new, exciting work, but done in a careful, compatible way, so that we can keep your upgrades from one toolchain to the next as boring as possible.

​	因此，在未来的几年里，您将看到许多新的、令人兴奋的工作，但是以谨慎、兼容的方式完成，以便我们可以尽可能地让从一个工具链升级到下一个工具链的过程变得无聊。