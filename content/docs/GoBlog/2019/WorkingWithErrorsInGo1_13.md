+++
title = "在 go 1.13中处理错误"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Working with Errors in Go 1.13 - 在 go 1.13中处理错误

https://go.dev/blog/go1.13-errors

Damien Neil and Jonathan Amsterdam
17 October 2019

## Introduction 简介

Go’s treatment of [errors as values](https://blog.golang.org/errors-are-values) has served us well over the last decade. Although the standard library’s support for errors has been minimal—just the `errors.New` and `fmt.Errorf` functions, which produce errors that contain only a message—the built-in `error` interface allows Go programmers to add whatever information they desire. All it requires is a type that implements an `Error` method:

在过去的十年中，Go将错误作为数值处理，这让我们受益匪浅。尽管标准库对错误的支持微乎其微--只有 errors.New 和 fmt.Errorf 函数，它们产生的错误只包含一个消息--内置的错误接口允许 Go 程序员添加任何他们想要的信息。它所需要的只是一个实现了Error方法的类型：

```go linenums="1"
type QueryError struct {
    Query string
    Err   error
}

func (e *QueryError) Error() string { return e.Query + ": " + e.Err.Error() }
```

Error types like this one are ubiquitous, and the information they store varies widely, from timestamps to filenames to server addresses. Often, that information includes another, lower-level error to provide additional context.

像这样的错误类型无处不在，它们存储的信息差别很大，从时间戳到文件名到服务器地址。通常，这些信息包括另一个较低级别的错误，以提供额外的背景。

The pattern of one error containing another is so pervasive in Go code that, after [extensive discussion](https://go.dev/issue/29934), Go 1.13 added explicit support for it. This post describes the additions to the standard library that provide that support: three new functions in the `errors` package, and a new formatting verb for `fmt.Errorf`.

一个错误包含另一个错误的模式在Go代码中非常普遍，所以经过广泛的讨论，Go 1.13增加了对它的明确支持。这篇文章描述了为标准库增加的支持：错误包中的三个新函数，以及fmt.Errorf的一个新格式化动词。

Before describing the changes in detail, let’s review how errors are examined and constructed in previous versions of the language.

在详细描述这些变化之前，让我们回顾一下在以前版本的语言中是如何检查和构造错误的。

## Errors before Go 1.13 Go 1.13之前的错误

### Examining errors 检查错误

Go errors are values. Programs make decisions based on those values in a few ways. The most common is to compare an error to `nil` to see if an operation failed.

Go的错误是数值。程序以几种方式根据这些值进行决策。最常见的是将错误与nil进行比较，看是否操作失败。

```go linenums="1"
if err != nil {
    // something went wrong
}
```

Sometimes we compare an error to a known *sentinel* value, to see if a specific error has occurred.

有时，我们将一个错误与一个已知的哨兵值进行比较，以查看是否发生了特定的错误。

```go linenums="1"
var ErrNotFound = errors.New("not found")

if err == ErrNotFound {
    // something wasn't found
}
```

An error value may be of any type which satisfies the language-defined `error` interface. A program can use a type assertion or type switch to view an error value as a more specific type.

一个错误值可以是任何满足语言定义的错误接口的类型。程序可以使用类型断言或类型转换，将错误值视为一个更具体的类型。

```go linenums="1"
type NotFoundError struct {
    Name string
}

func (e *NotFoundError) Error() string { return e.Name + ": not found" }

if e, ok := err.(*NotFoundError); ok {
    // e.Name wasn't found
}
```

### Adding information 添加信息

Frequently a function passes an error up the call stack while adding information to it, like a brief description of what was happening when the error occurred. A simple way to do this is to construct a new error that includes the text of the previous one:

经常有一个函数在调用堆栈上传递一个错误，同时向其添加信息，比如错误发生时的简要描述。一个简单的方法是构造一个新的错误，包括前一个错误的文本：

```go linenums="1"
if err != nil {
    return fmt.Errorf("decompress %v: %v", name, err)
}
```

Creating a new error with `fmt.Errorf` discards everything from the original error except the text. As we saw above with `QueryError`, we may sometimes want to define a new error type that contains the underlying error, preserving it for inspection by code. Here is `QueryError` again:

用fmt.Errorf创建一个新的错误，除了文本外，抛弃了原始错误的所有内容。正如我们在上面看到的QueryError，我们有时可能想定义一个新的错误类型，包含基本的错误，保留它以便于代码检查。这里又是QueryError：

```go linenums="1"
type QueryError struct {
    Query string
    Err   error
}
```

Programs can look inside a `*QueryError` value to make decisions based on the underlying error. You’ll sometimes see this referred to as “unwrapping” the error.

程序可以查看*QueryError值的内部，以根据基础错误做出决定。有时你会看到这被称为 "解包 "错误。

```go linenums="1"
if e, ok := err.(*QueryError); ok && e.Err == ErrPermission {
    // query failed because of a permission problem
}
```

The `os.PathError` type in the standard library is another example of one error which contains another.

标准库中的os.PathError类型是另一个例子，即一个错误包含另一个错误。

## Errors in Go 1.13 Go 1.13中的错误

### The Unwrap method - Unwrap方法

Go 1.13 introduces new features to the `errors` and `fmt` standard library packages to simplify working with errors that contain other errors. The most significant of these is a convention rather than a change: an error which contains another may implement an `Unwrap` method returning the underlying error. If `e1.Unwrap()` returns `e2`, then we say that `e1` *wraps* `e2`, and that you can *unwrap* `e1` to get `e2`.

Go 1.13 为 errors 和 fmt 标准库包引入了新的功能，以简化对包含其他错误的错误的处理。其中最重要的是一个惯例，而不是一个变化：一个包含其他错误的错误可以实现一个返回底层错误的Unwrap方法。如果e1.Unwrap()返回e2，那么我们说e1包裹了e2，并且你可以解开e1来获得e2。

Following this convention, we can give the `QueryError` type above an `Unwrap` method that returns its contained error:

按照这个惯例，我们可以给上面的QueryError类型一个Unwrap方法，返回其包含的错误：

```go linenums="1"
func (e *QueryError) Unwrap() error { return e.Err }
```

The result of unwrapping an error may itself have an `Unwrap` method; we call the sequence of errors produced by repeated unwrapping the *error chain*.

解除一个错误的结果可能本身就有一个Unwrap方法；我们把重复解包产生的错误序列称为错误链。

### Examining errors with Is and As 用Is和As检查错误

The Go 1.13 `errors` package includes two new functions for examining errors: `Is` and `As`.

Go 1.13 错误包包括两个用于检查错误的新函数。Is和As。

The `errors.Is` function compares an error to a value.

errors.Is函数将一个错误与一个值进行比较。

```go linenums="1"
// Similar to:
//   if err == ErrNotFound { … }
if errors.Is(err, ErrNotFound) {
    // something wasn't found
}
```

The `As` function tests whether an error is a specific type.

As函数测试一个错误是否是一个特定的类型。

```go linenums="1"
// Similar to:
//   if e, ok := err.(*QueryError); ok { … }
var e *QueryError
// Note: *QueryError is the type of the error.
if errors.As(err, &e) {
    // err is a *QueryError, and e is set to the error's value
}
```

In the simplest case, the `errors.Is` function behaves like a comparison to a sentinel error, and the `errors.As` function behaves like a type assertion. When operating on wrapped errors, however, these functions consider all the errors in a chain. Let’s look again at the example from above of unwrapping a `QueryError` to examine the underlying error:

在最简单的情况下，errors.Is函数的行为就像与一个哨兵错误的比较，而errors.As函数的行为就像一个类型断言。然而，当对包裹的错误进行操作时，这些函数会考虑一个链中的所有错误。让我们再看一下上面的例子，解开QueryError的包装，检查底层错误：

```go linenums="1"
if e, ok := err.(*QueryError); ok && e.Err == ErrPermission {
    // query failed because of a permission problem
}
```

Using the `errors.Is` function, we can write this as:

使用errors.Is函数，我们可以这样写：

```go linenums="1"
if errors.Is(err, ErrPermission) {
    // err, or some error that it wraps, is a permission problem
}
```

The `errors` package also includes a new `Unwrap` function which returns the result of calling an error’s `Unwrap` method, or `nil` when the error has no `Unwrap` method. It is usually better to use `errors.Is` or `errors.As`, however, since these functions will examine the entire chain in a single call.

errors包还包括一个新的Unwrap函数，它返回调用一个错误的Unwrap方法的结果，或者当错误没有Unwrap方法时返回nil。然而，通常最好使用 errors.Is 或 errors.As，因为这些函数将在一次调用中检查整个链条。

Note: although it may feel odd to take a pointer to a pointer, in this case it is correct. Think of it instead as taking a pointer to a value of the error type; it so happens in this case that the returned error is a pointer type.

注意：尽管拿一个指针到一个指针可能感觉很奇怪，但在这种情况下是正确的。可以把它看作是取一个指向错误类型值的指针；在这种情况下，恰好返回的错误是一个指针类型。

### Wrapping errors with %w  用%w包住错误

As mentioned earlier, it is common to use the `fmt.Errorf` function to add additional information to an error.

如前所述，使用fmt.Errorf函数为错误添加额外信息是很常见的。

```go linenums="1"
if err != nil {
    return fmt.Errorf("decompress %v: %v", name, err)
}
```

In Go 1.13, the `fmt.Errorf` function supports a new `%w` verb. When this verb is present, the error returned by `fmt.Errorf` will have an `Unwrap` method returning the argument of `%w`, which must be an error. In all other ways, `%w` is identical to `%v`.

在Go 1.13中，fmt.Errorf函数支持一个新的%w动词。当这个动词出现时，由fmt.Errorf返回的错误将有一个Unwrap方法返回%w的参数，这个参数必须是一个错误。在所有其他方面，%w与%v是相同的。

```go linenums="1"
if err != nil {
    // Return an error which unwraps to err.
    return fmt.Errorf("decompress %v: %w", name, err)
}
```

Wrapping an error with `%w` makes it available to `errors.Is` and `errors.As`:

用%w包装一个错误，使其可用于 errors.Is 和 errors.As：

```go linenums="1"
err := fmt.Errorf("access denied: %w", ErrPermission)
...
if errors.Is(err, ErrPermission) ...
```

### Whether to Wrap 是否包裹

When adding additional context to an error, either with `fmt.Errorf` or by implementing a custom type, you need to decide whether the new error should wrap the original. There is no single answer to this question; it depends on the context in which the new error is created. Wrap an error to expose it to callers. Do not wrap an error when doing so would expose implementation details.

当给一个错误添加额外的上下文时，无论是用fmt.Errorf还是通过实现一个自定义类型，你需要决定新的错误是否应该包住原来的。这个问题没有唯一的答案；它取决于创建新错误的上下文。包裹一个错误是为了将其暴露给调用者。如果这样做会暴露出实现的细节，就不要包裹错误。

As one example, imagine a `Parse` function which reads a complex data structure from an `io.Reader`. If an error occurs, we wish to report the line and column number at which it occurred. If the error occurs while reading from the `io.Reader`, we will want to wrap that error to allow inspection of the underlying problem. Since the caller provided the `io.Reader` to the function, it makes sense to expose the error produced by it.

举个例子，设想一个Parse函数从io.Reader读取一个复杂的数据结构。如果发生错误，我们希望报告发生错误的行和列号。如果错误是在从io.Reader中读取时发生的，我们希望将该错误包裹起来，以便于检查根本问题。由于调用者向函数提供了io.Reader，暴露由它产生的错误是有意义的。

In contrast, a function which makes several calls to a database probably should not return an error which unwraps to the result of one of those calls. If the database used by the function is an implementation detail, then exposing these errors is a violation of abstraction. For example, if the `LookupUser` function of your package `pkg` uses Go’s `database/sql` package, then it may encounter a `sql.ErrNoRows` error. If you return that error with `fmt.Errorf("accessing DB: %v", err)` then a caller cannot look inside to find the `sql.ErrNoRows`. But if the function instead returns `fmt.Errorf("accessing DB: %w", err)`, then a caller could reasonably write

相反，一个对数据库进行多次调用的函数可能不应该返回一个错误，这个错误是对其中一次调用结果的解包。如果函数所使用的数据库是一个实现细节，那么暴露这些错误就违反了抽象原则。例如，如果你的包 pkg 的 LookupUser 函数使用 Go 的数据库/sql 包，那么它可能遇到 sql.ErrNoRows 错误。如果你用fmt.Errorf("accessing DB: %v", err)来返回这个错误，那么调用者就无法在里面找到sql.ErrNoRows。但是如果该函数返回fmt.Errorf("accessing DB: %w", err)，那么调用者可以合理地写道

```go linenums="1"
err := pkg.LookupUser(...)
if errors.Is(err, sql.ErrNoRows) …
```

At that point, the function must always return `sql.ErrNoRows` if you don’t want to break your clients, even if you switch to a different database package. In other words, wrapping an error makes that error part of your API. If you don’t want to commit to supporting that error as part of your API in the future, you shouldn’t wrap the error.

在这一点上，如果你不想破坏你的客户，即使你切换到一个不同的数据库包，该函数必须总是返回sql.ErrNoRows。换句话说，包装一个错误使得这个错误成为你的API的一部分。如果你不想在将来把这个错误作为你的API的一部分来支持，你就不应该包装这个错误。

It’s important to remember that whether you wrap or not, the error text will be the same. A *person* trying to understand the error will have the same information either way; the choice to wrap is about whether to give *programs* additional information so they can make more informed decisions, or to withhold that information to preserve an abstraction layer.

重要的是要记住，无论你包不包，错误文本都是一样的。试图理解错误的人无论如何都会有相同的信息；选择包装是为了给程序提供额外的信息，使他们能够做出更明智的决定，还是为了保留一个抽象层而隐瞒这些信息。

## Customizing error tests with Is and As methods 用Is和As方法定制错误测试

The `errors.Is` function examines each error in a chain for a match with a target value. By default, an error matches the target if the two are [equal](https://go.dev/ref/spec#Comparison_operators). In addition, an error in the chain may declare that it matches a target by implementing an `Is` *method*.

errors.Is函数检查链中的每个错误与目标值的匹配情况。默认情况下，如果两者相等，则错误与目标值匹配。此外，链中的一个错误可以通过实现Is方法来声明它与目标值相匹配。

As an example, consider this error inspired by the [Upspin error package](https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html) which compares an error against a template, considering only fields which are non-zero in the template:

作为一个例子，考虑这个由Upspin错误包启发的错误，它将一个错误与一个模板进行比较，只考虑模板中的非零字段：

```go linenums="1"
type Error struct {
    Path string
    User string
}

func (e *Error) Is(target error) bool {
    t, ok := target.(*Error)
    if !ok {
        return false
    }
    return (e.Path == t.Path || t.Path == "") &&
           (e.User == t.User || t.User == "")
}

if errors.Is(err, &Error{User: "someuser"}) {
    // err's User field is "someuser".
}
```

The `errors.As` function similarly consults an `As` method when present.

## Errors and package APIs 错误和包的API

A package which returns errors (and most do) should describe what properties of those errors programmers may rely on. A well-designed package will also avoid returning errors with properties that should not be relied upon.

一个返回错误的包（大多数都是如此）应该描述这些错误的哪些属性，程序员可以依赖这些属性。一个设计良好的包也会避免返回具有不应该被依赖的属性的错误。

The simplest specification is to say that operations either succeed or fail, returning a nil or non-nil error value respectively. In many cases, no further information is needed.

最简单的规范是说，操作要么成功，要么失败，分别返回一个nil或非nil的错误值。在许多情况下，不需要进一步的信息。

If we wish a function to return an identifiable error condition, such as “item not found,” we might return an error wrapping a sentinel.

如果我们希望一个函数返回一个可识别的错误条件，如 "未找到项目"，我们可以返回一个包裹着哨兵的错误。

```go linenums="1"
var ErrNotFound = errors.New("not found")

// FetchItem returns the named item.
//
// If no item with the name exists, FetchItem returns an error
// wrapping ErrNotFound.
func FetchItem(name string) (*Item, error) {
    if itemNotFound(name) {
        return nil, fmt.Errorf("%q: %w", name, ErrNotFound)
    }
    // ...
}
```

There are other existing patterns for providing errors which can be semantically examined by the caller, such as directly returning a sentinel value, a specific type, or a value which can be examined with a predicate function.

还有其他现有的模式，用于提供可由调用者进行语义检查的错误，例如直接返回一个哨兵值、一个特定的类型，或者一个可由谓语函数检查的值。

In all cases, care should be taken not to expose internal details to the user. As we touched on in “Whether to Wrap” above, when you return an error from another package you should convert the error to a form that does not expose the underlying error, unless you are willing to commit to returning that specific error in the future.

在所有情况下，都应该注意不要向用户暴露内部细节。正如我们在上面的 "是否包装 "中提到的，当你从另一个包中返回一个错误时，你应该将错误转换为不暴露基本错误的形式，除非你愿意承诺在将来返回那个特定的错误。

```go linenums="1"
f, err := os.Open(filename)
if err != nil {
    // The *os.PathError returned by os.Open is an internal detail.
    // To avoid exposing it to the caller, repackage it as a new
    // error with the same text. We use the %v formatting verb, since
    // %w would permit the caller to unwrap the original *os.PathError.
    return fmt.Errorf("%v", err)
}
```

If a function is defined as returning an error wrapping some sentinel or type, do not return the underlying error directly.

如果一个函数被定义为返回包裹某种哨兵或类型的错误，不要直接返回底层错误。

```go linenums="1"
var ErrPermission = errors.New("permission denied")

// DoSomething returns an error wrapping ErrPermission if the user
// does not have permission to do something.
func DoSomething() error {
    if !userHasPermission() {
        // If we return ErrPermission directly, callers might come
        // to depend on the exact error value, writing code like this:
        //
        //     if err := pkg.DoSomething(); err == pkg.ErrPermission { … }
        //
        // This will cause problems if we want to add additional
        // context to the error in the future. To avoid this, we
        // return an error wrapping the sentinel so that users must
        // always unwrap it:
        //
        //     if err := pkg.DoSomething(); errors.Is(err, pkg.ErrPermission) { ... }
        return fmt.Errorf("%w", ErrPermission)
    }
    // ...
}
```

## Conclusion 结论

Although the changes we’ve discussed amount to just three functions and a formatting verb, we hope they will go a long way toward improving how errors are handled in Go programs. We expect that wrapping to provide additional context will become commonplace, helping programs to make better decisions and helping programmers to find bugs more quickly.

虽然我们所讨论的变化只是三个函数和一个格式化动词，但我们希望它们能在很大程度上改善Go程序中的错误处理方式。我们希望提供额外上下文的包装将成为普遍现象，帮助程序做出更好的决定，并帮助程序员更快发现错误。

As Russ Cox said in his [GopherCon 2019 keynote](https://blog.golang.org/experiment), on the path to Go 2 we experiment, simplify and ship. Now that we’ve shipped these changes, we look forward to the experiments that will follow.



正如Russ Cox在他的GopherCon 2019主题演讲中所说，在通往Go 2的道路上，我们进行了实验、简化和发货。现在我们已经交付了这些变化，我们期待着接下来的实验。
