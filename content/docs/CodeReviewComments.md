+++
title = "go 代码审查意见"
weight = 20
date = 2023-05-18T17:31:23+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Go Code Review Comments - Go代码审查意见

>原文：[https://github.com/golang/go/wiki/CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)

Ian Lance Taylor edited this page 12 days ago · [101 revisions](https://github.com/golang/go/wiki/CodeReviewComments/_history)

This page collects common comments made during reviews of Go code, so that a single detailed explanation can be referred to by shorthands. This is a laundry list of common mistakes, not a comprehensive style guide.

本页收集了审查围棋代码时的常见评论，以便通过速记法来参考单一的详细解释。这是一份常见错误的清单，而不是一份全面的风格指南。

You can view this as a supplement to [Effective Go](https://go.dev/doc/effective_go).

您可以把它看作是Effective Go的一个补充。

Additional comments related to testing can be found at [Go Test Comments](https://github.com/golang/go/wiki/TestComments)

与测试有关的其他评论可以在围棋测试评论中找到。

**Please [discuss changes](https://go.dev/issue/new?title=wiki%3A+CodeReviewComments+change&body=&labels=Documentation) before editing this page**, even *minor* ones. Many people have opinions and this is not the place for edit wars.https://github.com/golang/go/wiki/CodeReviewComments#variable-names)

请在编辑本页面之前讨论修改，即使是小的修改。很多人都有意见，这里不是编辑战争的地方。

## Gofmt

Run [gofmt](https://pkg.go.dev/cmd/gofmt/) on your code to automatically fix the majority of mechanical style issues. Almost all Go code in the wild uses `gofmt`. The rest of this document addresses non-mechanical style points.

在您的代码上运行gofmt，可以自动修复大部分的机械风格问题。几乎所有的 Go 代码都在使用 gofmt。本文档的其余部分将讨论非机械式风格问题。

An alternative is to use [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports), a superset of `gofmt` which additionally adds (and removes) import lines as necessary.

另一种方法是使用 goimports，它是 gofmt 的超集，可以在必要时增加（和删除）导入行。

## Comment Sentences 注释句子

See https://go.dev/doc/effective_go#commentary. Comments documenting declarations should be full sentences, even if that seems a little redundant. This approach makes them format well when extracted into godoc documentation. Comments should begin with the name of the thing being described and end in a period:

见 https://go.dev/doc/effective_go#commentary。记录声明的注释应该是完整的句子，即使这看起来有点多余。这种方法使它们在被提取到godoc文档中时格式化得很好。注释应该以被描述事物的名称开始，以句号结束：

```
// Request represents a request to run a command.
type Request struct { ...

// Encode writes the JSON encoding of req to w.
func Encode(w io.Writer, req *Request) { ...
```

and so on.

以此类推。

## Contexts

Values of the context.Context type carry security credentials, tracing information, deadlines, and cancellation signals across API and process boundaries. Go programs pass Contexts explicitly along the entire function call chain from incoming RPCs and HTTP requests to outgoing requests.

context.Context类型的值携带安全证书、跟踪信息、最后期限以及跨API和进程边界的取消信号。Go程序在整个函数调用链中明确地传递Context，从传入的RPC和HTTP请求到传出的请求。

Most functions that use a Context should accept it as their first parameter:

大多数使用 Context 的函数都应该接受它作为其第一个参数：

```
func F(ctx context.Context, /* other arguments */) {}
```

A function that is never request-specific may use context.Background(), but err on the side of passing a Context even if you think you don't need to. The default case is to pass a Context; only use context.Background() directly if you have a good reason why the alternative is a mistake.

一个从来没有特定请求的函数可以使用context.background()，但即使您认为不需要，也要在传递Context方面犯错误。默认情况下是传递一个Context；只有当您有充分的理由说明另一种选择是错误的时候才会直接使用context.background()。

Don't add a Context member to a struct type; instead add a ctx parameter to each method on that type that needs to pass it along. The one exception is for methods whose signature must match an interface in the standard library or in a third party library.

不要给结构类型添加Context成员；而是给该类型上需要传递Context的每个方法添加一个ctx参数。唯一的例外是那些签名必须符合标准库或第三方库中的接口的方法。

Don't create custom Context types or use interfaces other than Context in function signatures.

不要创建自定义的Context类型或在函数签名中使用Context以外的接口。

If you have application data to pass around, put it in a parameter, in the receiver, in globals, or, if it truly belongs there, in a Context value.

如果您有应用数据需要传递，把它放在一个参数中，放在接收器中，放在globals中，或者，如果它真的属于那里，放在一个Context值中。

Contexts are immutable, so it's fine to pass the same ctx to multiple calls that share the same deadline, cancellation signal, credentials, parent trace, etc.

Context是不可改变的，所以把相同的ctx传递给共享相同的截止日期、取消信号、证书、父级跟踪等的多个调用是可以的。

## Copying 复制

To avoid unexpected aliasing, be careful when copying a struct from another package. For example, the bytes.Buffer type contains a `[]byte` slice. If you copy a `Buffer`, the slice in the copy may alias the array in the original, causing subsequent method calls to have surprising effects.

为了避免意外的别名，在从其他包复制结构时要小心。例如，bytes.Buffer类型包含一个[]字节分片。如果您拷贝一个Buffer，拷贝中的片断可能会与原结构中的数组产生别名，导致后续的方法调用产生意外的效果。

In general, do not copy a value of type `T` if its methods are associated with the pointer type, `*T`.

一般来说，如果一个T类型的值的方法与指针类型*T有关，就不要复制它。

## Crypto Rand

Do not use package `math/rand` to generate keys, even throwaway ones. Unseeded, the generator is completely predictable. Seeded with `time.Nanoseconds()`, there are just a few bits of entropy. Instead, use `crypto/rand`'s Reader, and if you need text, print to hexadecimal or base64:

不要使用软件包math/rand来生成密钥，即使是抛弃式的。不加种子，生成器是完全可预测的。用time.Nanoseconds()作为种子，就只有几个比特的熵了。相反，使用crypto/rand的Reader，如果您需要文本，打印成十六进制或base64：

```go linenums="1"
import (
	"crypto/rand"
	// "encoding/base64"
	// "encoding/hex"
	"fmt"
)

func Key() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)  // out of randomness, should never happen
	}
	return fmt.Sprintf("%x", buf)
	// or hex.EncodeToString(buf)
	// or base64.StdEncoding.EncodeToString(buf)
}
```

## Declaring Empty Slices 声明空的切片

When declaring an empty slice, prefer

当声明一个空切片时，最好选择

```
var t []string
```

over

而不是

```
t := []string{}
```

The former declares a nil slice value, while the latter is non-nil but zero-length. They are functionally equivalent—their `len` and `cap` are both zero—but the nil slice is the preferred style.

前者声明的是一个空的片断值，而后者是非空的，但长度为零。它们在功能上是等价的--它们的长度和上限都是零--但是nil slice是首选样式。

Note that there are limited circumstances where a non-nil but zero-length slice is preferred, such as when encoding JSON objects (a `nil` slice encodes to `null`, while `[]string{}` encodes to the JSON array `[]`).

请注意，在有限的情况下，非零但零长度的分片是首选，例如在编码JSON对象时（nil分片编码为null，而[]string{}编码为JSON数组[]）。

When designing interfaces, avoid making a distinction between a nil slice and a non-nil, zero-length slice, as this can lead to subtle programming errors.

在设计接口时，避免区分nil片断和非nil的零长度片断，因为这可能会导致微妙的编程错误。

For more discussion about nil in Go see Francesc Campoy's talk [Understanding Nil](https://www.youtube.com/watch?v=ynoY2xz-F8s).

关于Go中nil的更多讨论，请参见Francesc Campoy的讲座Understanding Nil。

## Doc Comments 文件注释

All top-level, exported names should have doc comments, as should non-trivial unexported type or function declarations. See https://go.dev/doc/effective_go#commentary for more information about commentary conventions.

所有顶层的、导出的名字都应该有文档注释，正如非重要的未导出的类型或函数声明一样。请参阅 https://go.dev/doc/effective_go#commentary 了解更多关于注释惯例的信息。

## Don't Panic

See https://go.dev/doc/effective_go#errors. Don't use panic for normal error handling. Use error and multiple return values.

参见 https://go.dev/doc/effective_go#errors。不要在正常的错误处理中使用panic。使用错误和多个返回值。

## Error Strings 错误字符串

Error strings should not be capitalized (unless beginning with proper nouns or acronyms) or end with punctuation, since they are usually printed following other context. That is, use `fmt.Errorf("something bad")` not `fmt.Errorf("Something bad")`, so that `log.Printf("Reading %s: %v", filename, err)` formats without a spurious capital letter mid-message. This does not apply to logging, which is implicitly line-oriented and not combined inside other messages.

错误字符串不应大写（除非以专有名词或缩略语开始）或以标点符号结尾，因为它们通常是在其他上下文之后打印的。也就是说，使用fmt.Errorf("something bad")而不是fmt.Errorf("Something bad")，这样log.Printf("Reading %s: %v", filename, err)的格式就不会在信息中间出现虚假的大写字母。这不适用于日志，因为日志是隐含的以行为导向的，并且不在其他消息中组合。

## Examples 例子

When adding a new package, include examples of intended usage: a runnable Example, or a simple test demonstrating a complete call sequence.

当添加一个新的包时，包括预期使用的例子：一个可运行的例子，或者一个展示完整调用序列的简单测试。

Read more about [testable Example() functions](https://go.dev/blog/examples).

阅读更多关于可测试的 Example() 函数。

## Goroutine Lifetimes 协作程序寿命

When you spawn goroutines, make it clear when - or whether - they exit.

当您生成goroutine时，要明确说明它们何时或是否退出。

Goroutines can leak by blocking on channel sends or receives: the garbage collector will not terminate a goroutine even if the channels it is blocked on are unreachable.

goroutine可以通过阻塞通道的发送或接收来泄密：即使goroutine阻塞的通道无法到达，垃圾收集器也不会终止它。

Even when goroutines do not leak, leaving them in-flight when they are no longer needed can cause other subtle and hard-to-diagnose problems. Sends on closed channels panic. Modifying still-in-use inputs "after the result isn't needed" can still lead to data races. And leaving goroutines in-flight for arbitrarily long can lead to unpredictable memory usage.

即使goroutine没有泄漏，当它们不再需要时，让它们继续飞行也会引起其他微妙的、难以诊断的问题。在封闭的通道上发送消息会引起恐慌。在 "不需要结果之后 "修改仍在使用的输入，仍然会导致数据竞赛。让goroutines在飞行中停留任意长的时间会导致不可预知的内存使用。

Try to keep concurrent code simple enough that goroutine lifetimes are obvious. If that just isn't feasible, document when and why the goroutines exit.

尽量保持并发代码的简单性，使goroutine的寿命足够明显。如果这不可行，就记录下goroutines退出的时间和原因。

## Handle Errors 处理错误

See https://go.dev/doc/effective_go#errors. Do not discard errors using `_` variables. If a function returns an error, check it to make sure the function succeeded. Handle the error, return it, or, in truly exceptional situations, panic.

见https://go.dev/doc/effective_go#errors。不要使用_变量丢弃错误。如果一个函数返回一个错误，请检查它以确保该函数成功。处理错误，返回错误，或者，在真正特殊的情况下，惊慌失措。

## Imports 导入

Avoid renaming imports except to avoid a name collision; good package names should not require renaming. In the event of collision, prefer to rename the most local or project-specific import.

避免重命名进口，除非是为了避免名称冲突；好的包名不应该需要重命名。在发生冲突的情况下，最好重命名最本地的或项目特定的导入。

Imports are organized in groups, with blank lines between them. The standard library packages are always in the first group.

进口是按组组织的，它们之间有空行。标准库包总是在第一组。

```go linenums="1"
package main

import (
	"fmt"
	"hash/adler32"
	"os"

	"github.com/foo/bar"
	"rsc.io/goversion/version"
)
```

[goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) will do this for you.

goimports 会为您做这个。

## Import Blank 导入空白

Packages that are imported only for their side effects (using the syntax `import _ "pkg"`) should only be imported in the main package of a program, or in tests that require them.

只为其副作用而导入的包（使用 import _ "pkg" 语法）应该只在程序的主包中导入，或者在需要它们的测试中导入。

## Import Dot 导入点

The import . form can be useful in tests that, due to circular dependencies, cannot be made part of the package being tested:

import .形式在测试中很有用，由于循环依赖关系，不能成为被测试包的一部分。

```
package foo_test

import (
	"bar/testutil" // also imports "foo"
	. "foo"
)
```

In this case, the test file cannot be in package foo because it uses bar/testutil, which imports foo. So we use the 'import .' form to let the file pretend to be part of package foo even though it is not. Except for this one case, do not use import . in your programs. It makes the programs much harder to read because it is unclear whether a name like Quux is a top-level identifier in the current package or in an imported package.

在这种情况下，测试文件不可能在包foo中，因为它使用的是bar/testutil，它导入了foo。所以我们使用'import . '的形式，让文件假装是包foo的一部分，尽管它不是。除了这种情况，不要在您的程序中使用import .。它使程序更难阅读，因为不清楚像Quux这样的名字是当前包中的顶级标识符还是导入包中的顶级标识符。

## In-Band Errors 带内错误

In C and similar languages, it's common for functions to return values like -1 or null to signal errors or missing results:

在C语言和类似语言中，函数返回-1或null这样的值是很常见的，以示错误或丢失结果：

```
// Lookup returns the value for key or "" if there is no mapping for key.
func Lookup(key string) string

// Failing to check for an in-band error value can lead to bugs:
Parse(Lookup(key))  // returns "parse failure for value" instead of "no value for key"
```

Go's support for multiple return values provides a better solution. Instead of requiring clients to check for an in-band error value, a function should return an additional value to indicate whether its other return values are valid. This return value may be an error, or a boolean when no explanation is needed. It should be the final return value.

Go对多个返回值的支持提供了一个更好的解决方案。与其要求客户端检查一个带内错误值，一个函数应该返回一个额外的值来表明其其他返回值是否有效。这个返回值可以是一个错误，或者在不需要解释时是一个布尔值。它应该是最终的返回值。

```
// Lookup returns the value for key or ok=false if there is no mapping for key.
func Lookup(key string) (value string, ok bool)
```

This prevents the caller from using the result incorrectly:

这可以防止调用者错误地使用结果：

```
Parse(Lookup(key))  // compile-time error
```

And encourages more robust and readable code:

并鼓励更稳健和可读的代码：

```
value, ok := Lookup(key)
if !ok {
	return fmt.Errorf("no value for %q", key)
}
return Parse(value)
```

This rule applies to exported functions but is also useful for unexported functions.

这个规则适用于导出的函数，但对未导出的函数也很有用。

Return values like nil, "", 0, and -1 are fine when they are valid results for a function, that is, when the caller need not handle them differently from other values.

像nil、""、0和-1这样的返回值，当它们是一个函数的有效结果时，也就是说，当调用者不需要对它们进行与其他值不同的处理时，就可以了。

Some standard library functions, like those in package "strings", return in-band error values. This greatly simplifies string-manipulation code at the cost of requiring more diligence from the programmer. In general, Go code should return additional values for errors.

一些标准库函数，如包 "strings "中的函数，返回带内错误值。这大大简化了字符串处理的代码，但代价是要求程序员更加勤奋。一般来说，Go代码应该为错误返回额外的值。

## Indent Error Flow 缩进错误流程

Try to keep the normal code path at a minimal indentation, and indent the error handling, dealing with it first. This improves the readability of the code by permitting visually scanning the normal path quickly. For instance, don't write:

尽量使正常的代码路径保持最小的缩进，并缩进错误处理，首先处理错误。这可以提高代码的可读性，因为它允许在视觉上快速扫描正常路径。例如，不要写：

```
if err != nil {
	// error handling
} else {
	// normal code
}
```

Instead, write:

相反，要写：

```
if err != nil {
	// error handling
	return // or continue, etc.
}
// normal code
```

If the `if` statement has an initialization statement, such as:

如果if语句有一个初始化语句，比如说：

```
if x, err := f(); err != nil {
	// error handling
	return
} else {
	// use x
}
```

then this may require moving the short variable declaration to its own line:

那么这可能需要将短变量声明移到它自己的一行：

```
x, err := f()
if err != nil {
	// error handling
	return
}
// use x
```

## Initialisms 首字母缩写

Words in names that are initialisms or acronyms (e.g. "URL" or "NATO") have a consistent case. For example, "URL" should appear as "URL" or "url" (as in "urlPony", or "URLPony"), never as "Url". As an example: ServeHTTP not ServeHttp. For identifiers with multiple initialized "words", use for example "xmlHTTPRequest" or "XMLHTTPRequest".

名称中的单词如果是首字母缩写或缩略语（例如 "URL "或 "NATO"），则有一个统一的大小写。例如，"URL "应该显示为 "URL "或 "url"（如 "urlPony"，或 "URLPony"），而不是 "Url"。作为一个例子。ServeHTTP不是ServeHttp。对于有多个初始化 "词 "的标识符，使用例如 "xmlHTTPRequest "或 "XMLHTTPRequest"。

This rule also applies to "ID" when it is short for "identifier" (which is pretty much all cases when it's not the "id" as in "ego", "superego"), so write "appID" instead of "appId".

这个规则也适用于 "ID"，当它是 "标识符 "的简称时（这几乎是所有的情况，当它不是 "自我"、"超我 "中的 "ID "时），所以写 "appID "而不是 "appId"。

Code generated by the protocol buffer compiler is exempt from this rule. Human-written code is held to a higher standard than machine-written code.

由协议缓冲区编译器生成的代码不受这一规则影响。人写的代码要比机器写的代码有更高的标准。

## Interfaces 接口

Go interfaces generally belong in the package that uses values of the interface type, not the package that implements those values. The implementing package should return concrete (usually pointer or struct) types: that way, new methods can be added to implementations without requiring extensive refactoring.

Go接口通常属于使用接口类型值的包，而不是实现这些值的包。实现包应该返回具体的（通常是指针或结构）类型：这样一来，新的方法可以被添加到实现中，而不需要大量的重构。

Do not define interfaces on the implementor side of an API "for mocking"; instead, design the API so that it can be tested using the public API of the real implementation.

不要在API的实现者一方定义接口，"为了嘲弄"；相反，要设计API，使其可以使用真正实现的公共API来测试。

Do not define interfaces before they are used: without a realistic example of usage, it is too difficult to see whether an interface is even necessary, let alone what methods it ought to contain.

不要在使用之前就定义接口：如果没有一个真实的使用例子，就很难看出一个接口是否有必要，更不用说它应该包含哪些方法。

```
package consumer  // consumer.go

type Thinger interface { Thing() bool }

func Foo(t Thinger) string { … }
package consumer // consumer_test.go

type fakeThinger struct{ … }
func (t fakeThinger) Thing() bool { … }
…
if Foo(fakeThinger{…}) == "x" { … }
// DO NOT DO IT!!!
package producer

type Thinger interface { Thing() bool }

type defaultThinger struct{ … }
func (t defaultThinger) Thing() bool { … }

func NewThinger() Thinger { return defaultThinger{ … } }
```

Instead return a concrete type and let the consumer mock the producer implementation.

取而代之的是返回一个具体的类型，让消费者模拟生产者的实现。

```
package producer

type Thinger struct{ … }
func (t Thinger) Thing() bool { … }

func NewThinger() Thinger { return Thinger{ … } }
```

## Line Length 线条长度

There is no rigid line length limit in Go code, but avoid uncomfortably long lines. Similarly, don't add line breaks to keep lines short when they are more readable long--for example, if they are repetitive.

Go代码中没有严格的行长限制，但要避免不舒服的长行。同样地，如果行长更容易阅读，就不要添加换行符来保持行短--例如，如果行是重复的。

Most of the time when people wrap lines "unnaturally" (in the middle of function calls or function declarations, more or less, say, though some exceptions are around), the wrapping would be unnecessary if they had a reasonable number of parameters and reasonably short variable names. Long lines seem to go with long names, and getting rid of the long names helps a lot.

大多数时候，当人们 "不自然地 "换行时（在函数调用或函数声明的中间，或多或少，比如说，尽管周围有一些例外），如果他们有合理数量的参数和合理简短的变量名称，那么换行是不必要的。长线似乎与长名字相伴，摆脱长名字有很大的帮助。

In other words, break lines because of the semantics of what you're writing (as a general rule) and not because of the length of the line. If you find that this produces lines that are too long, then change the names or the semantics and you'll probably get a good result.

换句话说，断行是因为您所写的内容的语义（作为一般规则），而不是因为行的长度。如果您发现这样做产生的行太长，那么改变名称或语义，您可能会得到一个好结果。

This is, actually, exactly the same advice about how long a function should be. There's no rule "never have a function more than N lines long", but there is definitely such a thing as too long of a function, and of too repetitive tiny functions, and the solution is to change where the function boundaries are, not to start counting lines.

实际上，这与关于一个函数应该有多长的建议完全一样。没有 "永远不要让一个函数超过N行 "的规则，但绝对存在一个太长的函数，以及太多重复的小函数，解决方法是改变函数边界的位置，而不是开始计算行。

## Mixed Caps 混杂盖帽

See https://go.dev/doc/effective_go#mixed-caps. This applies even when it breaks conventions in other languages. For example an unexported constant is `maxLength` not `MaxLength` or `MAX_LENGTH`.

见https://go.dev/doc/effective_go#mixed-caps。这一点甚至在打破其他语言的惯例时也适用。例如，一个未导出的常数是maxLength而不是MaxLength或MAX_LENGTH。

Also see [Initialisms](https://github.com/golang/go/wiki/CodeReviewComments#initialisms).

也请参见首字母缩写。

## Named Result Parameters 命名的结果参数

Consider what it will look like in godoc. Named result parameters like:

考虑一下在godoc中会是什么样子。命名的结果参数如：

```
func (n *Node) Parent1() (node *Node) {}
func (n *Node) Parent2() (node *Node, err error) {}
```

will be repetitive in godoc; better to use:

在godoc中会有重复，最好使用：

```
func (n *Node) Parent1() *Node {}
func (n *Node) Parent2() (*Node, error) {}
```

On the other hand, if a function returns two or three parameters of the same type, or if the meaning of a result isn't clear from context, adding names may be useful in some contexts. Don't name result parameters just to avoid declaring a var inside the function; that trades off a minor implementation brevity at the cost of unnecessary API verbosity.

另一方面，如果一个函数返回两个或三个相同类型的参数，或者一个结果的含义在上下文中并不明确，那么在某些情况下添加名称可能是有用的。不要为了避免在函数中声明一个var而对结果参数进行命名；这样做是以不必要的API的冗长为代价来换取一个小的实现的简洁性。

```
func (f *Foo) Location() (float64, float64, error)
```

is less clear than:

不太清楚：

```
// Location returns f's latitude and longitude.
// Negative values mean south and west, respectively.
func (f *Foo) Location() (lat, long float64, err error)
```

Naked returns are okay if the function is a handful of lines. Once it's a medium sized function, be explicit with your return values. Corollary: it's not worth it to name result parameters just because it enables you to use naked returns. Clarity of docs is always more important than saving a line or two in your function.

如果函数只有寥寥几行，裸返是可以的。一旦它是一个中等规模的函数，就要明确您的返回值。推论：不值得为结果参数命名，因为这可以让您使用裸返回。文档的清晰性总是比在您的函数中节省一两行更重要。

Finally, in some cases you need to name a result parameter in order to change it in a deferred closure. That is always OK.

最后，在某些情况下，您需要命名一个结果参数，以便在一个延迟闭包中改变它。这总是可以的。

## Naked Returns 赤裸裸的返回

A `return` statement without arguments returns the named return values. This is known as a "naked" return.

一个没有参数的返回语句会返回指定的返回值。这就是所谓的 "裸 "返回。

```
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

See [Named Result Parameters](https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters).

参见命名的结果参数。

## Package Comments 包注释

Package comments, like all comments to be presented by godoc, must appear adjacent to the package clause, with no blank line.

包注释，像所有由godoc呈现的注释一样，必须出现在包子句的旁边，不能有空行。

```
// Package math provides basic constants and mathematical functions.
package math
/*
Package template implements data-driven templates for generating textual
output such as HTML.
....
*/
package template
```

For "package main" comments, other styles of comment are fine after the binary name (and it may be capitalized if it comes first), For example, for a `package main` in the directory `seedgen` you could write:

对于 "package main "的注释，其他样式的注释在二进制名称后面也可以（如果它在前面，可以大写），例如，对于目录seedgen中的package main，您可以写：

```
// Binary seedgen ...
package main
```

or 或

```
// Command seedgen ...
package main
```

or 或

```
// Program seedgen ...
package main
```

or 或

```
// The seedgen command ...
package main
```

or 或

```
// The seedgen program ...
package main
```

or 或

```
// Seedgen ..
package main
```

These are examples, and sensible variants of these are acceptable.

这些都是例子，其中合理的变体也是可以接受的。

Note that starting the sentence with a lower-case word is not among the acceptable options for package comments, as these are publicly-visible and should be written in proper English, including capitalizing the first word of the sentence. When the binary name is the first word, capitalizing it is required even though it does not strictly match the spelling of the command-line invocation.

请注意，以小写字母开始的句子不在软件包注释可接受的选项之列，因为这些注释是公开可见的，应该用正确的英语书写，包括将句子的第一个单词大写。当二进制名称是第一个词时，即使它与命令行调用的拼写不严格匹配，也需要大写。

See https://go.dev/doc/effective_go#commentary for more information about commentary conventions.

关于注释惯例的更多信息，见https://go.dev/doc/effective_go#commentary。

## Package Names  包名称

All references to names in your package will be done using the package name, so you can omit that name from the identifiers. For example, if you are in package chubby, you don't need type ChubbyFile, which clients will write as `chubby.ChubbyFile`. Instead, name the type `File`, which clients will write as `chubby.File`. Avoid meaningless package names like util, common, misc, api, types, and interfaces. See https://go.dev/doc/effective_go#package-names and https://go.dev/blog/package-names for more.

在您的包中所有对名字的引用都将使用包名，所以您可以在标识符中省略该名字。例如，如果您在包 chubby 中，您不需要类型 ChubbyFile，客户端会把它写成 chubby.ChubbyFile。取而代之的是，将类型File命名为客户将写成chubby.File。避免使用无意义的包名，如util、common、misc、api、types和interface。参见https://go.dev/doc/effective_go#package-names 和 https://go.dev/blog/package-names 了解更多。

## Pass Values 传递值

Don't pass pointers as function arguments just to save a few bytes. If a function refers to its argument `x` only as `*x` throughout, then the argument shouldn't be a pointer. Common instances of this include passing a pointer to a string (`*string`) or a pointer to an interface value (`*io.Reader`). In both cases the value itself is a fixed size and can be passed directly. This advice does not apply to large structs, or even small structs that might grow.

不要为了节省几个字节而传递指针作为函数参数。如果一个函数自始至终只把它的参数x称为*x，那么该参数就不应该是一个指针。常见的例子包括传递一个指向字符串的指针（*string）或一个指向接口值的指针（*io.Reader）。在这两种情况下，值本身是一个固定的大小，可以直接传递。这个建议不适用于大型结构，甚至是可能增长的小型结构。

## Receiver Names 接收器名称

The name of a method's receiver should be a reflection of its identity; often a one or two letter abbreviation of its type suffices (such as "c" or "cl" for "Client"). Don't use generic names such as "me", "this" or "self", identifiers typical of object-oriented languages that gives the method a special meaning. In Go, the receiver of a method is just another parameter and therefore, should be named accordingly. The name need not be as descriptive as that of a method argument, as its role is obvious and serves no documentary purpose. It can be very short as it will appear on almost every line of every method of the type; familiarity admits brevity. Be consistent, too: if you call the receiver "c" in one method, don't call it "cl" in another.

一个方法的接收者的名字应该反映它的身份；通常一个或两个字母的类型缩写就足够了（如 "c "或 "cl "代表 "Client"）。不要使用诸如 "me"、"this "或 "self "这样的通用名称，这些是面向对象语言的典型标识，它们赋予了方法以特殊的含义。在Go中，一个方法的接收者只是另一个参数，因此，应该相应地命名。这个名字不需要像方法参数那样具有描述性，因为它的作用是显而易见的，没有任何文件上的作用。它可以很短，因为它几乎会出现在该类型的每一个方法的每一行；熟悉的人都会接受简洁。也要保持一致：如果您在一个方法中称接收器为 "c"，不要在另一个方法中称它为 "cl"。

## Receiver Type 接收器类型

Choosing whether to use a value or pointer receiver on methods can be difficult, especially to new Go programmers. If in doubt, use a pointer, but there are times when a value receiver makes sense, usually for reasons of efficiency, such as for small unchanging structs or values of basic type. Some useful guidelines:

选择在方法上使用值或指针接收器可能是困难的，特别是对新的Go程序员来说。如果有疑问，请使用指针，但有时值接收器是有意义的，通常是出于效率的考虑，例如小型不变的结构或基本类型的值。一些有用的准则：

- If the receiver is a map, func or chan, don't use a pointer to them. If the receiver is a slice and the method doesn't reslice or reallocate the slice, don't use a pointer to it.如果接收器是一个map、func或chan，不要使用一个指针。如果接收方是一个片断，并且该方法不重新划分或重新分配片断，不要使用指向它的指针。
- If the method needs to mutate the receiver, the receiver must be a pointer.如果方法需要突变接收器，接收器必须是一个指针。
- If the receiver is a struct that contains a sync.Mutex or similar synchronizing field, the receiver must be a pointer to avoid copying.如果接收器是一个包含sync.Mutex或类似同步字段的结构，接收器必须是一个指针以避免复制。
- If the receiver is a large struct or array, a pointer receiver is more efficient. How large is large? Assume it's equivalent to passing all its elements as arguments to the method. If that feels too large, it's also too large for the receiver.如果接收器是一个大的结构体或数组，那么指针式的接收器会更有效。多大才算大？假设它相当于把所有的元素作为参数传递给方法。如果这感觉太大，那么对于接收器来说也是太大了。
- Can function or methods, either concurrently or when called from this method, be mutating the receiver? A value type creates a copy of the receiver when the method is invoked, so outside updates will not be applied to this receiver. If changes must be visible in the original receiver, the receiver must be a pointer.函数或方法，无论是并发的还是从这个方法中调用的，都会对接收器进行变异吗？当方法被调用时，一个值类型会创建一个接收器的副本，所以外部的更新不会被应用到这个接收器上。如果变化必须在原始接收器中可见，接收器必须是一个指针。
- If the receiver is a struct, array or slice and any of its elements is a pointer to something that might be mutating, prefer a pointer receiver, as it will make the intention clearer to the reader.如果接收器是一个结构体、数组或片断，并且它的任何元素都是一个指向可能发生变化的东西的指针，那么最好是一个指针接收器，因为它将使读者更清楚地了解其意图。
- If the receiver is a small array or struct that is naturally a value type (for instance, something like the time.Time type), with no mutable fields and no pointers, or is just a simple basic type such as int or string, a value receiver makes sense. A value receiver can reduce the amount of garbage that can be generated; if a value is passed to a value method, an on-stack copy can be used instead of allocating on the heap. (The compiler tries to be smart about avoiding this allocation, but it can't always succeed.) Don't choose a value receiver type for this reason without profiling first.如果接收器是一个小的数组或结构，自然是一个值类型（例如，像time.Time类型），没有可变的字段，也没有指针，或者只是一个简单的基本类型，如int或string，值接收器是有意义的。一个值接收器可以减少可能产生的垃圾量；如果一个值被传递给一个值方法，可以使用堆上拷贝而不是在堆上分配。(编译器试图聪明地避免这种分配，但它不可能总是成功）。在没有进行分析之前，不要因为这个原因选择一个值接收器类型。
- Don't mix receiver types. Choose either pointers or struct types for all available methods.不要混用接收器类型。为所有可用的方法选择指针或结构类型。
- Finally, when in doubt, use a pointer receiver.最后，当有疑问时，使用一个指针接收器。

## Synchronous Functions 同步函数

Prefer synchronous functions - functions which return their results directly or finish any callbacks or channel ops before returning - over asynchronous ones.

优先选择同步函数--直接返回结果或在返回前完成任何回调或通道操作的函数，而不是异步函数。

Synchronous functions keep goroutines localized within a call, making it easier to reason about their lifetimes and avoid leaks and data races. They're also easier to test: the caller can pass an input and check the output without the need for polling or synchronization.

同步函数使goroutines在调用中保持本地化，使其更容易推理其生命周期并避免泄漏和数据竞赛。它们也更容易测试：调用者可以传递一个输入并检查输出，而不需要轮询或同步。

If callers need more concurrency, they can add it easily by calling the function from a separate goroutine. But it is quite difficult - sometimes impossible - to remove unnecessary concurrency at the caller side.

如果调用者需要更多的并发性，他们可以通过从一个单独的goroutine中调用该函数来轻松地添加它。但是，在调用者一方删除不必要的并发是相当困难的，有时甚至是不可能的。

## Useful Test Failures 有用的测试失败

Tests should fail with helpful messages saying what was wrong, with what inputs, what was actually got, and what was expected. It may be tempting to write a bunch of assertFoo helpers, but be sure your helpers produce useful error messages. Assume that the person debugging your failing test is not you, and is not your team. A typical Go test fails like:

测试失败时应该有有用的信息，说明什么地方出了问题，用什么输入，实际得到了什么，以及预期得到了什么。写一堆assertFoo辅助工具可能很诱人，但要确保您的辅助工具产生有用的错误信息。假设调试您失败的测试的人不是您，也不是您的团队。一个典型的Go测试失败的情况如下：

```
if got != tt.want {
	t.Errorf("Foo(%q) = %d; want %d", tt.in, got, tt.want) // or Fatalf, if test can't test anything more past this point
}
```

Note that the order here is actual != expected, and the message uses that order too. Some test frameworks encourage writing these backwards: 0 != x, "expected 0, got x", and so on. Go does not.

注意，这里的顺序是actual != expected，而且消息也使用这个顺序。有些测试框架鼓励把这些东西倒过来写。0 != x，"预期0，得到x"，以此类推。Go则不然。

If that seems like a lot of typing, you may want to write a [table-driven test](https://github.com/golang/go/wiki/TableDrivenTests).

如果这看起来像大量的打字，您可能想写一个表驱动的测试。

Another common technique to disambiguate failing tests when using a test helper with different input is to wrap each caller with a different TestFoo function, so the test fails with that name:

在使用不同输入的测试助手时，另一个常见的技术是用不同的TestFoo函数来包装每个调用者，从而使测试以该名称失败：

```
func TestSingleValue(t *testing.T) { testHelper(t, []int{80}) }
func TestNoValues(t *testing.T)    { testHelper(t, []int{}) }
```

In any case, the onus is on you to fail with a helpful message to whoever's debugging your code in the future.

在任何情况下，您都有责任为将来调试您的代码的人提供一个有用的信息，使其失败。

## Variable Names 变量名

Variable names in Go should be short rather than long. This is especially true for local variables with limited scope. Prefer `c` to `lineCount`. Prefer `i` to `sliceIndex`.

Go中的变量名应该是短的而不是长的。这对范围有限的局部变量来说尤其如此。倾向于用c来表示lineCount。倾向于用i来表示sliceIndex。

The basic rule: the further from its declaration that a name is used, the more descriptive the name must be. For a method receiver, one or two letters is sufficient. Common variables such as loop indices and readers can be a single letter (`i`, `r`). More unusual things and global variables need more descriptive names.

基本规则：名字使用的时间离其声明越远，名字的描述性越强。对于一个方法接收器来说，一个或两个字母就足够了。常见的变量，如循环索引和读取器可以是一个字母（i，r）。更多不寻常的东西和全局变量需要更多的描述性名称。