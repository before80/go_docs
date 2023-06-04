+++
title = "保持您的模块兼容"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Keeping Your Modules Compatible - 保持您的模块兼容

https://go.dev/blog/module-compatibility

Jean de Klerk and Jonathan Amsterdam
7 July 2020

## Introduction 简介

This post is part 5 in a series.

这篇文章是系列文章的第五部分。

- Part 1 — [Using Go Modules 使用Go模块](https://go.dev/blog/using-go-modules)
- Part 2 — [Migrating To Go Modules 迁移到Go模块](https://go.dev/blog/migrating-to-go-modules)
- Part 3 — [Publishing Go Modules 发布Go模块](https://go.dev/blog/publishing-go-modules)
- Part 4 — [Go Modules: v2 and Beyond Go模块：V2版及以后](https://go.dev/blog/v2-go-modules)
- **Part 5 — Keeping Your Modules Compatible** (this post) 保持模块的兼容性（本帖）

**Note:** For documentation on developing modules, see [Developing and publishing modules](https://go.dev/doc/modules/developing).

注意：关于开发模块的文档，请参见开发和发布模块。

Your modules will evolve over time as you add new features, change behaviors, and reconsider parts of the module’s public surface. As discussed in [Go Modules: v2 and Beyond](https://go.dev/blog/v2-go-modules), breaking changes to a v1+ module must happen as part of a major version bump (or by adopting a new module path).

您的模块会随着时间的推移而发展，因为您会添加新的功能，改变行为，以及重新考虑模块的部分公共表面。正如《Go Modules: v2 and Beyond》中所讨论的，对 v1+ 模块的突破性改变必须作为主要版本升级的一部分（或通过采用新的模块路径）。

However, releasing a new major version is hard on your users. They have to find the new version, learn a new API, and change their code. And some users may never update, meaning you have to maintain two versions for your code forever. So it is usually better to change your existing package in a compatible way.

然而，发布一个新的主要版本对您的用户来说是困难的。他们必须找到新的版本，学习新的API，并改变他们的代码。而且有些用户可能永远不会更新，这意味着您必须永远为您的代码维护两个版本。所以，通常情况下，以一种兼容的方式改变您现有的包是更好的。

In this post, we’ll explore some techniques for introducing non-breaking changes. The common theme is: add, don’t change or remove. We’ll also talk about how to design your API for compatibility from the outset.

在这篇文章中，我们将探讨一些引入非破坏性变化的技巧。共同的主题是：添加，不要改变或删除。我们还将讨论如何从一开始就为兼容性设计您的API。

## Adding to a function 添加到一个函数

Often, breaking changes come in the form of new arguments to a function. We’ll describe some ways to deal with this sort of change, but first let’s look at a technique that doesn’t work.

通常情况下，破坏性变化是以函数的新参数形式出现的。我们将描述一些处理这种变化的方法，但首先让我们看看一种不起作用的技术。

When adding new arguments with sensible defaults, it’s tempting to add them as a variadic parameter. To extend the function

当添加具有合理默认值的新参数时，把它们作为一个变量参数来添加是很诱人的。为了扩展函数

```go linenums="1"
func Run(name string)
```

with an additional `size` argument which defaults to zero, one might propose

添加一个默认为零的额外的大小参数，我们可以提议

```go linenums="1"
func Run(name string, size ...int)
```

on the grounds that all existing call sites will continue to work. While that is true, other uses of `Run` could break, like this one:

的建议，理由是所有现有的调用站点将继续工作。虽然这是真的，但Run的其他用法可能会被破坏，比如这个：

```go linenums="1"
package mypkg
var runner func(string) = yourpkg.Run
```

The original `Run` function works here because its type is `func(string)`, but the new `Run` function’s type is `func(string, ...int)`, so the assignment fails at compile time.

原有的Run函数在这里工作，因为它的类型是func(string)，但新的Run函数的类型是func(string, ...int)，所以在编译时赋值失败。

This example illustrates that call compatibility is not enough for backward compatibility. There is, in fact, no backward-compatible change you can make to a function’s signature.

这个例子说明了调用兼容性对于向后兼容是不够的。事实上，您不能对函数的签名做任何向后兼容的改变。

Instead of changing a function’s signature, add a new function. As an example, after the `context` package was introduced, it became common practice to pass a `context.Context` as the first argument to a function. However, stable APIs could not change an exported function to accept a `context.Context` because it would break all uses of that function.

与其改变一个函数的签名，不如添加一个新的函数。举个例子，在引入context包后，将context.Context作为第一个参数传递给函数成为一种普遍做法。然而，稳定的API不能改变导出的函数以接受context.Context，因为这将破坏该函数的所有使用。

Instead, new functions were added. For example, the `database/sql` package’s `Query` method’s signature was (and still is)

相反，新的函数被添加。例如，数据库/sql包的Query方法的签名是（现在仍然是）

```go linenums="1"
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
```

When the `context` package was created, the Go team added a new method to `database/sql`:

在创建上下文包时，Go团队为数据库/sql添加了一个新方法：

```go linenums="1"
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)
```

To avoid copying code, the old method calls the new one:

为了避免复制代码，旧方法调用新方法：

```go linenums="1"
func (db *DB) Query(query string, args ...interface{}) (*Rows, error) {
    return db.QueryContext(context.Background(), query, args...)
}
```

Adding a method allows users to migrate to the new API at their own pace. Since the methods read similarly and sort together, and `Context` is in the name of the new method, this extension of the `database/sql` API did not degrade readability or comprehension of the package.

增加一个方法可以让用户按照自己的节奏迁移到新的API。由于这些方法的读取方式类似，并且在一起排序，而且Context在新方法的名称中，这种数据库/sql API的扩展并没有降低包的可读性或理解力。

If you anticipate that a function may need more arguments in the future, you can plan ahead by making optional arguments a part of the function’s signature. The simplest way to do that is to add a single struct argument, as the [crypto/tls.Dial](https://pkg.go.dev/crypto/tls?tab=doc#Dial) function does:

如果您预计一个函数在未来可能需要更多的参数，您可以通过将可选参数作为函数签名的一部分来提前计划。最简单的方法是添加一个结构参数，就像 crypto/tls.Dial 函数那样：

```go linenums="1"
func Dial(network, addr string, config *Config) (*Conn, error)
```

The TLS handshake conducted by `Dial` requires a network and address, but it has many other parameters with reasonable defaults. Passing a `nil` for `config` uses those defaults; passing a `Config` struct with some fields set will override the defaults for those fields. In the future, adding a new TLS configuration parameter only requires a new field on the `Config` struct, a change that is backward-compatible (almost always—see "Maintaining struct compatibility" below).

Dial进行的TLS握手需要一个网络和地址，但它有许多其他参数，有合理的默认值。传递一个nil的config会使用这些默认值；传递一个设置了某些字段的Config结构会覆盖这些字段的默认值。在未来，添加一个新的TLS配置参数只需要在Config结构上添加一个新的字段，这种变化是向后兼容的（几乎总是如此--见下面的 "保持结构兼容性"）。

Sometimes the techniques of adding a new function and adding options can be combined by making the options struct a method receiver. Consider the evolution of the `net` package’s ability to listen at a network address. Prior to Go 1.11, the `net` package provided only a `Listen` function with the signature

有时，添加新函数和添加选项的技术可以通过使选项结构成为方法接收器而结合起来。考虑一下net包监听网络地址的能力的演变。在 Go 1.11 之前，net 包只提供了一个 Listen 函数，其签名为

```go linenums="1"
func Listen(network, address string) (Listener, error)
```

For Go 1.11, two features were added to `net` listening: passing a context, and allowing the caller to provide a "control function" to adjust the raw connection after creation but before binding. The result could have been a new function that took a context, network, address and control function. Instead, the package authors added a [`ListenConfig`](https://pkg.go.dev/net@go1.11?tab=doc#ListenConfig) struct in anticipation that more options might be needed someday. And rather than define a new top-level function with a cumbersome name, they added a `Listen` method to `ListenConfig`:

在Go 1.11中，net监听增加了两个功能：传递一个上下文，以及允许调用者提供一个 "控制函数 "来在创建后但在绑定前调整原始连接。其结果可能是一个新的函数，需要一个上下文、网络、地址和控制函数。相反，包的作者添加了一个ListenConfig结构，因为他们预计有一天会需要更多的选项。他们没有定义一个名称繁琐的新顶层函数，而是给ListenConfig添加了一个Listen方法：

```go linenums="1"
type ListenConfig struct {
    Control func(network, address string, c syscall.RawConn) error
}

func (*ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)
```

Another way to provide new options in the future is the "Option types" pattern, where options are passed as variadic arguments, and each option is a function that changes the state of the value being constructed. They are described in more detail by Rob Pike’s post [Self-referential functions and the design of options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html). One widely used example is [google.golang.org/grpc](https://pkg.go.dev/google.golang.org/grpc?tab=doc)’s [DialOption](https://pkg.go.dev/google.golang.org/grpc?tab=doc#DialOption).

未来提供新选项的另一种方式是 "选项类型 "模式，其中选项作为变量参数被传递，每个选项是一个改变被构建值状态的函数。Rob Pike的帖子Self-referential functions and the design of options对它们进行了更详细的描述。一个广泛使用的例子是 google.golang.org/grpc 的 DialOption。

Option types fulfill the same role as struct options in function arguments: they are an extensible way to pass behavior-modifying configuration. Deciding which to choose is largely a matter of style. Consider this simple usage of gRPC’s `DialOption` option type:

选项类型实现了与函数参数中的结构选项相同的作用：它们是一种可扩展的方式来传递改变行为的配置。决定选择哪种类型在很大程度上是一个风格问题。考虑一下gRPC的DialOption选项类型的这个简单用法：

```go linenums="1"
grpc.Dial("some-target",
  grpc.WithAuthority("some-authority"),
  grpc.WithMaxDelay(time.Second),
  grpc.WithBlock())
```

This could also have been implemented as a struct option:

这也可以作为一个结构选项来实现：

```go linenums="1"
notgrpc.Dial("some-target", &notgrpc.Options{
  Authority: "some-authority",
  MaxDelay:  time.Second,
  Block:     true,
})
```

Functional options have some downsides: they require writing the package name before the option for each call; they increase the size of the package namespace; and it’s unclear what the behavior should be if the same option is provided twice. On the other hand, functions which take option structs require a parameter which might almost always be `nil`, which some find unattractive. And when a type’s zero value has a valid meaning, it is clumsy to specify that the option should have its default value, typically requiring a pointer or an additional boolean field.

功能性选项有一些缺点：它们需要在每次调用时在选项前写上包的名字；它们增加了包的名字空间的大小；而且不清楚如果同一个选项被提供两次应该有什么行为。另一方面，接受选项结构的函数需要一个几乎总是为零的参数，这让一些人觉得没有吸引力。而当一个类型的零值有一个有效的含义时，指定选项应该有其默认值是很笨拙的，通常需要一个指针或一个额外的布尔字段。

Either one is a reasonable choice for ensuring future extensibility of your module’s public API.

无论哪一种都是合理的选择，以确保您的模块的公共API未来的可扩展性。

## Working with interfaces 与接口一起工作

Sometimes, new features require changes to publicly-exposed interfaces: for example, an interface needs to be extended with new methods. Directly adding to an interface is a breaking change, though—how, then, can we support new methods on a publicly-exposed interface?

有时，新功能需要对公开暴露的接口进行修改：例如，一个接口需要用新的方法进行扩展。然而，直接添加到一个接口上是一种破坏性的改变--那么，我们如何才能支持公开暴露的接口上的新方法？

The basic idea is to define a new interface with the new method, and then wherever the old interface is used, dynamically check whether the provided type is the older type or the newer type.

基本的想法是定义一个带有新方法的新接口，然后在使用旧接口的地方，动态地检查所提供的类型是旧类型还是新类型。

Let’s illustrate this with an example from the [`archive/tar`](https://pkg.go.dev/archive/tar?tab=doc) package. [`tar.NewReader`](https://pkg.go.dev/archive/tar?tab=doc#NewReader) accepts an `io.Reader`, but over time the Go team realized that it would be more efficient to skip from one file header to the next if you could call [`Seek`](https://pkg.go.dev/io?tab=doc#Seeker). But, they could not add a `Seek` method to `io.Reader`: that would break all implementers of `io.Reader`.

让我们用archive/tar包中的一个例子来说明。tar.NewReader接受io.Reader，但随着时间的推移，Go团队意识到，如果您能调用Seek，从一个文件头跳到下一个文件头会更有效率。但是，他们不能给io.Reader增加一个Seek方法：这将破坏io.Reader的所有实现者。

Another ruled-out option was to change `tar.NewReader` to accept [`io.ReadSeeker`](https://pkg.go.dev/io?tab=doc#ReadSeeker) rather than `io.Reader`, since it supports both `io.Reader` methods and `Seek` (by way of `io.Seeker`). But, as we saw above, changing a function signature is also a breaking change.

另一个被排除的选项是改变tar.NewReader来接受io.ReadSeeker而不是io.Reader，因为它同时支持io.Reader方法和Seek（通过io.Seeker的方式）。但是，正如我们在上面看到的，改变一个函数签名也是一种破坏性的改变。

So, they decided to keep `tar.NewReader` signature unchanged, but type check for (and support) `io.Seeker` in `tar.Reader` methods:

所以，他们决定保持tar.NewReader签名不变，但在tar.Reader方法中对io.Seeker进行类型检查（并支持）：

```go linenums="1"
package tar

type Reader struct {
  r io.Reader
}

func NewReader(r io.Reader) *Reader {
  return &Reader{r: r}
}

func (r *Reader) Read(b []byte) (int, error) {
  if rs, ok := r.r.(io.Seeker); ok {
    // Use more efficient rs.Seek.
  }
  // Use less efficient r.r.Read.
}
```

(See [reader.go](https://github.com/golang/go/blob/60f78765022a59725121d3b800268adffe78bde3/src/archive/tar/reader.go#L837) for the actual code.)

(实际代码见reader.go)。

When you run into a case where you want to add a method to an existing interface, you may be able to follow this strategy. Start by creating a new interface with your new method, or identify an existing interface with the new method. Next, identify the relevant functions that need to support it, type check for the second interface, and add code that uses it.

当您遇到要给一个现有的接口添加方法的情况时，您也许可以遵循这个策略。首先，用您的新方法创建一个新的接口，或者确定一个现有的接口，用新方法。接下来，确定需要支持它的相关函数，对第二个接口进行类型检查，并添加使用它的代码。

This strategy only works when the old interface without the new method can still be supported, limiting the future extensibility of your module.

这种策略只有在没有新方法的旧接口仍能被支持时才有效，这限制了您的模块未来的可扩展性。

Where possible, it is better to avoid this class of problem entirely. When designing constructors, for example, prefer to return concrete types. Working with concrete types allows you to add methods in the future without breaking users, unlike interfaces. That property allows your module to be extended more easily in the future.

在可能的情况下，最好是完全避免这类问题。例如，在设计构造函数时，最好返回具体类型。与接口不同，使用具体类型的工作允许您在未来增加方法而不破坏用户，。这一特性使您的模块在未来更容易被扩展。

Tip: if you do need to use an interface but don’t intend for users to implement it, you can add an unexported method. This prevents types defined outside your package from satisfying your interface without embedding, freeing you to add methods later without breaking user implementations. For example, see [`testing.TB`’s `private()` function](https://github.com/golang/go/blob/83b181c68bf332ac7948f145f33d128377a09c42/src/testing/testing.go#L564-L567).

提示：如果您确实需要使用一个接口，但不打算让用户实现它，您可以添加一个未导出的方法。这可以防止在您的包之外定义的类型在没有嵌入的情况下满足您的接口，使您可以在以后添加方法而不破坏用户的实现。例如，请看 testing.TB 的 private() 函数。

```go linenums="1"
// TB is the interface common to T and B.
type TB interface {
    Error(args ...interface{})
    Errorf(format string, args ...interface{})
    // ...

    // A private method to prevent users implementing the
    // interface and so future additions to it will not
    // violate Go 1 compatibility.
    private()
}
```

This topic is also explored in more detail in Jonathan Amsterdam’s "Detecting Incompatible API Changes" talk ([video](https://www.youtube.com/watch?v=JhdL5AkH-AQ), [slides](https://github.com/gophercon/2019-talks/blob/master/JonathanAmsterdam-DetectingIncompatibleAPIChanges/slides.pdf)).

这个话题在Jonathan Amsterdam的 "检测不兼容的API变化 "讲座中也有更详细的探讨（视频，幻灯片）。

## Add configuration methods 添加配置方法

So far we’ve talked about overt breaking changes, where changing a type or a function would cause users' code to stop compiling. However, behavior changes can also break users, even if user code continues to compile. For example, many users expect [`json.Decoder`](https://pkg.go.dev/encoding/json?tab=doc#Decoder) to ignore fields in the JSON that are not in the argument struct. When the Go team wanted to return an error in that case, they had to be careful. Doing so without an opt-in mechanism would mean that the many users relying on those methods might start receiving errors where they hadn’t before.

到目前为止，我们已经谈到了公开的破坏性变化，即改变一个类型或一个函数会导致用户的代码停止编译。然而，行为改变也会破坏用户，即使用户代码继续编译。例如，许多用户希望json.Decoder能够忽略JSON中不在参数结构中的字段。当Go团队想在这种情况下返回一个错误时，他们必须要小心。在没有选择机制的情况下这样做，意味着许多依赖这些方法的用户可能会在以前没有的地方开始收到错误。

So, rather than changing the behavior for all users, they added a configuration method to the `Decoder` struct: [`Decoder.DisallowUnknownFields`](https://pkg.go.dev/encoding/json?tab=doc#Decoder.DisallowUnknownFields). Calling this method opts a user in to the new behavior, but not doing so preserves the old behavior for existing users.

因此，他们没有改变所有用户的行为，而是给解码器结构添加了一个配置方法。Decoder.DisallowUnknownFields。调用这个方法可以让用户选择新的行为，但不这样做的话，就可以为现有用户保留旧的行为。

## Maintaining struct compatibility 保持结构的兼容性

We saw above that any change to a function’s signature is a breaking change. The situation is much better with structs. If you have an exported struct type, you can almost always add a field or remove an unexported field without breaking compatibility. When adding a field, make sure that its zero value is meaningful and preserves the old behavior, so that existing code that doesn’t set the field continues to work.

我们在上面看到，对函数签名的任何改变都是一种破坏性改变。对于结构体来说，情况要好得多。如果您有一个导出的结构类型，您几乎总是可以添加一个字段或删除一个未导出的字段而不破坏兼容性。在添加字段时，要确保其零值是有意义的，并保留旧的行为，这样现有的不设置字段的代码就能继续工作。

Recall that the authors of the `net` package added `ListenConfig` in Go 1.11 because they thought more options might be forthcoming. Turns out they were right. In Go 1.13, the [`KeepAlive` field](https://pkg.go.dev/net@go1.13?tab=doc#ListenConfig) was added to allow for disabling keep-alive or changing its period. The default value of zero preserves the original behavior of enabling keep-alive with a default period.

记得net包的作者在Go 1.11中添加了ListenConfig，因为他们认为可能会有更多的选项。事实证明他们是对的。在Go 1.13中，KeepAlive字段被添加进来，允许禁用keep-alive或改变其周期。默认值为0，保留了原来的行为，即启用默认周期的keep-alive。

There is one subtle way a new field can break user code unexpectedly. If all the field types in a struct are comparable—meaning values of those types can be compared with `==` and `!=` and used as a map key—then the overall struct type is comparable too. In this case, adding a new field of uncomparable type will make the overall struct type non-comparable, breaking any code that compares values of that struct type.

有一种微妙的方式，新字段可能会意外地破坏用户代码。如果一个结构中的所有字段类型都是可比较的--意味着这些类型的值可以用==和!=来比较，并作为映射键使用--那么整个结构类型也是可比较的。在这种情况下，添加一个不可比较类型的新字段将使整个结构类型不可比较，从而破坏任何比较该结构类型值的代码。

To keep a struct comparable, don’t add non-comparable fields to it. You can write a test for that, or rely on the upcoming [gorelease](https://pkg.go.dev/golang.org/x/exp/cmd/gorelease?tab=doc) tool to catch it.

为了保持结构的可比性，不要向其添加不可比较的字段。您可以为此编写一个测试，或者依靠即将推出的gorelease工具来捕获它。

To prevent comparison in the first place, make sure the struct has a non-comparable field. It may have one already—no slice, map or function type is comparable—but if not, one can be added like so:

为了从一开始就防止比较，确保该结构有一个不可比较的字段。它可能已经有了--没有切片、地图或函数类型可供比较，但如果没有，可以像这样添加一个：

```go linenums="1"
type Point struct {
        _ [0]func()
        X int
        Y int
}
```

The `func()` type is not comparable, and the zero-length array takes up no space. We can define a type to clarify our intent:

func()类型没有可比性，零长度的数组不占用空间。我们可以定义一个类型来阐明我们的意图：

```go linenums="1"
type doNotCompare [0]func()

type Point struct {
        doNotCompare
        X int
        Y int
}
```

Should you use `doNotCompare` in your structs? If you’ve defined the struct to be used as a pointer—that is, it has pointer methods and perhaps a `NewXXX` constructor function that returns a pointer—then adding a `doNotCompare` field is probably overkill. Users of a pointer type understand that each value of the type is distinct: that if they want to compare two values, they should compare the pointers.

您应该在您的结构中使用doNotCompare吗？如果您定义的结构是作为指针使用的，也就是说，它有指针方法，也许还有一个返回指针的 NewXXX 构造函数，那么添加 doNotCompare 字段可能是多余的。指针类型的用户知道该类型的每个值都是不同的：如果他们想比较两个值，他们应该比较指针。

If you are defining a struct intended to be used as a value directly, like our `Point` example, then quite often you want it to be comparable. In the uncommon case that you have a value struct that you don’t want compared, then adding a `doNotCompare` field will give you the freedom to change the struct later without having to worry about breaking comparisons. On the downside, the type won’t be usable as a map key.

如果您定义的结构打算直接作为一个值使用，就像我们的Point例子，那么您经常希望它是可比较的。在不常见的情况下，您有一个不希望被比较的值结构，那么添加一个 doNotCompare 字段将给您自由，以后可以改变这个结构而不必担心破坏比较。缺点是，该类型不能作为地图键使用。

## Conclusion 结论

When planning an API from scratch, consider carefully how extensible the API will be to new changes in the future. And when you do need to add new features, remember the rule: add, don’t change or remove, keeping in mind the exceptions—interfaces, function arguments, and return values can’t be added in backwards-compatible ways.

在从头开始规划API时，要仔细考虑API对未来新变化的可扩展性如何。当您确实需要添加新的功能时，记住这个规则：添加，不要改变或删除，牢记例外情况--接口、函数参数和返回值不能以向后兼容的方式添加。

If you need to dramatically change an API, or if an API begins to lose its focus as more features are added, then it may be time for a new major version. But most of the time, making a backwards-compatible change is easy and avoids causing pain for your users.

如果您需要大幅度地改变一个API，或者一个API随着更多的功能的增加而开始失去重点，那么可能是时候进行新的主要版本了。但大多数时候，做一个向后兼容的改变是很容易的，而且可以避免给您的用户带来痛苦。
