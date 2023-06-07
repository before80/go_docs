+++
title = "错误处理和 go"
weight = 15
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Error handling and Go - 错误处理和 go

[https://go.dev/blog/error-handling-and-go](https://go.dev/blog/error-handling-and-go)

Andrew Gerrand
12 July 2011

## Introduction 简介

If you have written any Go code you have probably encountered the built-in `error` type. Go code uses `error` values to indicate an abnormal state. For example, the `os.Open` function returns a non-nil `error` value when it fails to open a file.

如果您写过任何Go代码，您可能遇到过内置的错误类型。Go代码使用错误值来表示异常的状态。例如，os.Open 函数在打开一个文件失败时返回一个非零的错误值。

```go linenums="1"
func Open(name string) (file *File, err error)
```

The following code uses `os.Open` to open a file. If an error occurs it calls `log.Fatal` to print the error message and stop.

下面的代码使用os.Open来打开一个文件。如果发生错误，它会调用log.Fatal来打印错误信息并停止。

```go linenums="1"
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```

You can get a lot done in Go knowing just this about the `error` type, but in this article we’ll take a closer look at `error` and discuss some good practices for error handling in Go.

仅仅知道这些关于错误类型的信息，您就可以在Go中完成很多事情，但在这篇文章中，我们将仔细研究错误，并讨论Go中错误处理的一些良好做法。

## The error type 错误类型

The `error` type is an interface type. An `error` variable represents any value that can describe itself as a string. Here is the interface’s declaration:

错误类型是一种接口类型。一个错误变量代表任何可以用字符串描述自己的值。下面是接口的声明。

```go linenums="1"
type error interface {
    Error() string
}
```

The `error` type, as with all built in types, is [predeclared](https://go.dev/doc/go_spec.html#Predeclared_identifiers) in the [universe block](https://go.dev/doc/go_spec.html#Blocks).

错误类型和所有的内置类型一样，都是在宇宙块中预先声明的。

The most commonly-used `error` implementation is the [errors](https://go.dev/pkg/errors/) package’s unexported `errorString` type.

最常用的错误实现是错误包的未导出的errorString类型。

```go linenums="1"
// errorString is a trivial implementation of error.
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

You can construct one of these values with the `errors.New` function. It takes a string that it converts to an `errors.errorString` and returns as an `error` value.

您可以用error.New函数构建这些值中的一个。它接受一个字符串，将其转换为error.errorString并作为一个错误值返回。

```go linenums="1"
// New returns an error that formats as the given text.
func New(text string) error {
    return &errorString{text}
}
```

Here’s how you might use `errors.New`:

下面是您如何使用error.New。

```go linenums="1"
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // implementation
}
```

A caller passing a negative argument to `Sqrt` receives a non-nil `error` value (whose concrete representation is an `errors.errorString` value). The caller can access the error string ("math: square root of…") by calling the `error`’s `Error` method, or by just printing it:

传递负数参数给Sqrt的调用者会收到一个非零的错误值（其具体表示是error.errorString值）。调用者可以通过调用错误的Error方法访问错误字符串（"math: square root of..."），或者直接打印它。

```go linenums="1"
f, err := Sqrt(-1)
if err != nil {
    fmt.Println(err)
}
```

The [fmt](https://go.dev/pkg/fmt/) package formats an `error` value by calling its `Error() string` method.

fmt包通过调用其Error()字符串方法来格式化一个错误值。

It is the error implementation’s responsibility to summarize the context. The error returned by `os.Open` formats as "open /etc/passwd: permission denied," not just "permission denied." The error returned by our `Sqrt` is missing information about the invalid argument.

总结上下文是错误实现的责任。os.Open返回的错误格式为 "open /etc/passwd: permission denied"，而不只是 "permission denied"。我们的Sqrt返回的错误缺少关于无效参数的信息。

To add that information, a useful function is the `fmt` package’s `Errorf`. It formats a string according to `Printf`’s rules and returns it as an `error` created by `errors.New`.

为了增加这些信息，一个有用的函数是fmt包的Errorf。它根据Printf的规则格式化一个字符串，并将其作为error.New创建的错误返回。

```go linenums="1"
if f < 0 {
    return 0, fmt.Errorf("math: square root of negative number %g", f)
}
```

In many cases `fmt.Errorf` is good enough, but since `error` is an interface, you can use arbitrary data structures as error values, to allow callers to inspect the details of the error.

在许多情况下，fmt.Errorf已经足够好了，但由于error是一个接口，您可以使用任意的数据结构作为错误值，以允许调用者检查错误的细节。

For instance, our hypothetical callers might want to recover the invalid argument passed to `Sqrt`. We can enable that by defining a new error implementation instead of using `errors.errorString`:

例如，我们假设的调用者可能想恢复传递给Sqrt的无效参数。我们可以通过定义一个新的错误实现而不是使用error.errorString来实现。

```go linenums="1"
type NegativeSqrtError float64

func (f NegativeSqrtError) Error() string {
    return fmt.Sprintf("math: square root of negative number %g", float64(f))
}
```

A sophisticated caller can then use a [type assertion](https://go.dev/doc/go_spec.html#Type_assertions) to check for a `NegativeSqrtError` and handle it specially, while callers that just pass the error to `fmt.Println` or `log.Fatal` will see no change in behavior.

复杂的调用者可以使用类型断言来检查NegativeSqrtError并进行特别处理，而那些只是将错误传递给fmt.Println或log.Fatal的调用者将不会看到行为上的变化。

As another example, the [json](https://go.dev/pkg/encoding/json/) package specifies a `SyntaxError` type that the `json.Decode` function returns when it encounters a syntax error parsing a JSON blob.

作为另一个例子，json包指定了一个SyntaxError类型，当json.Decode函数在解析JSON blob时遇到语法错误时返回。

```go linenums="1"
type SyntaxError struct {
    msg    string // description of error
    Offset int64  // error occurred after reading Offset bytes
}

func (e *SyntaxError) Error() string { return e.msg }
```

The `Offset` field isn’t even shown in the default formatting of the error, but callers can use it to add file and line information to their error messages:

Offset字段甚至没有显示在错误的默认格式中，但是调用者可以用它来给错误信息添加文件和行的信息。

```go linenums="1"
if err := dec.Decode(&val); err != nil {
    if serr, ok := err.(*json.SyntaxError); ok {
        line, col := findLine(f, serr.Offset)
        return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
    }
    return err
}
```

(This is a slightly simplified version of some [actual code](https://github.com/camlistore/go4/blob/03efcb870d84809319ea509714dd6d19a1498483/jsonconfig/eval.go#L123-L135) from the [Camlistore](http://camlistore.org/) project.)

(这是对Camlistore项目中一些实际代码的略微简化版本）。

The `error` interface requires only a `Error` method; specific error implementations might have additional methods. For instance, the [net](https://go.dev/pkg/net/) package returns errors of type `error`, following the usual convention, but some of the error implementations have additional methods defined by the `net.Error` interface:

错误接口只需要一个Error方法；特定的错误实现可能有额外的方法。例如，net包按照通常的惯例返回error类型的错误，但是一些错误实现有net.Error接口定义的额外方法。

```go linenums="1"
package net

type Error interface {
    error
    Timeout() bool   // Is the error a timeout?
    Temporary() bool // Is the error temporary?
}
```

Client code can test for a `net.Error` with a type assertion and then distinguish transient network errors from permanent ones. For instance, a web crawler might sleep and retry when it encounters a temporary error and give up otherwise.

客户端代码可以用一个类型断言来测试net.Error，然后将暂时性的网络错误与永久性的错误区分开。例如，一个网络爬虫可能会在遇到临时错误时睡眠并重试，否则就放弃。

```go linenums="1"
if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
    time.Sleep(1e9)
    continue
}
if err != nil {
    log.Fatal(err)
}
```

## Simplifying repetitive error handling 简化重复性错误处理

In Go, error handling is important. The language’s design and conventions encourage you to explicitly check for errors where they occur (as distinct from the convention in other languages of throwing exceptions and sometimes catching them). In some cases this makes Go code verbose, but fortunately there are some techniques you can use to minimize repetitive error handling.

在Go中，错误处理很重要。该语言的设计和惯例鼓励您在错误发生时明确地检查错误（与其他语言中抛出异常和有时捕捉异常的惯例不同）。在某些情况下，这使得Go代码变得冗长，但幸运的是，您可以使用一些技术来减少重复的错误处理。

Consider an [App Engine](https://cloud.google.com/appengine/docs/go/) application with an HTTP handler that retrieves a record from the datastore and formats it with a template.

考虑一个带有HTTP处理程序的App Engine应用程序，该处理程序从数据存储中获取一条记录，并以模板格式化。

```go linenums="1"
func init() {
    http.HandleFunc("/view", viewRecord)
}

func viewRecord(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    if err := viewTemplate.Execute(w, record); err != nil {
        http.Error(w, err.Error(), 500)
    }
}
```

This function handles errors returned by the `datastore.Get` function and `viewTemplate`’s `Execute` method. In both cases, it presents a simple error message to the user with the HTTP status code 500 ("Internal Server Error"). This looks like a manageable amount of code, but add some more HTTP handlers and you quickly end up with many copies of identical error handling code.

这个函数处理由datastore.Get函数和viewTemplate的Execute方法返回的错误。在这两种情况下，它都会向用户呈现一个简单的错误信息，其HTTP状态码为500（"内部服务器错误"）。这看起来是一个可控的代码量，但如果增加一些HTTP处理程序，您很快就会出现许多相同的错误处理代码的副本。

To reduce the repetition we can define our own HTTP `appHandler` type that includes an `error` return value:

为了减少重复，我们可以定义我们自己的HTTP appHandler类型，包括一个错误返回值。

```go linenums="1"
type appHandler func(http.ResponseWriter, *http.Request) error
```

Then we can change our `viewRecord` function to return errors:

然后我们可以改变我们的viewRecord函数来返回错误。

```go linenums="1"
func viewRecord(w http.ResponseWriter, r *http.Request) error {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        return err
    }
    return viewTemplate.Execute(w, record)
}
```

This is simpler than the original version, but the [http](https://go.dev/pkg/net/http/) package doesn’t understand functions that return `error`. To fix this we can implement the `http.Handler` interface’s `ServeHTTP` method on `appHandler`:

这比原来的版本要简单，但http包不理解返回错误的函数。为了解决这个问题，我们可以在appHandler上实现http.Handler接口的ServeHTTP方法。

```go linenums="1"
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := fn(w, r); err != nil {
        http.Error(w, err.Error(), 500)
    }
}
```

The `ServeHTTP` method calls the `appHandler` function and displays the returned error (if any) to the user. Notice that the method’s receiver, `fn`, is a function. (Go can do that!) The method invokes the function by calling the receiver in the expression `fn(w, r)`.

ServeHTTP方法调用appHandler函数并向用户显示返回的错误（如果有的话）。请注意，该方法的接收器fn是一个函数。(Go可以做到这一点！）该方法通过调用表达式fn(w, r)中的接收器来调用该函数。

Now when registering `viewRecord` with the http package we use the `Handle` function (instead of `HandleFunc`) as `appHandler` is an `http.Handler` (not an `http.HandlerFunc`).

现在当用http包注册viewRecord时，我们使用Handle函数（而不是HandleFunc），因为appHandler是一个http.Handler（而不是http.HandlerFunc）。

```go linenums="1"
func init() {
    http.Handle("/view", appHandler(viewRecord))
}
```

With this basic error handling infrastructure in place, we can make it more user friendly. Rather than just displaying the error string, it would be better to give the user a simple error message with an appropriate HTTP status code, while logging the full error to the App Engine developer console for debugging purposes.

有了这个基本的错误处理基础设施，我们就可以让它变得更加友好。与其只显示错误字符串，不如给用户一个简单的错误信息，并附上适当的HTTP状态码，同时将完整的错误记录到App Engine开发者控制台，以便进行调试。

To do this we create an `appError` struct containing an `error` and some other fields:

```go linenums="1"
type appError struct {
    Error   error
    Message string
    Code    int
}
```

Next we modify the appHandler type to return `*appError` values:

接下来我们修改appHandler类型以返回*appError值。

```go linenums="1"
type appHandler func(http.ResponseWriter, *http.Request) *appError
```

(It’s usually a mistake to pass back the concrete type of an error rather than `error`, for reasons discussed in [the Go FAQ](https://go.dev/doc/go_faq.html#nil_error), but it’s the right thing to do here because `ServeHTTP` is the only place that sees the value and uses its contents.)

(由于Go FAQ中讨论的原因，通常传回错误的具体类型而不是错误是一个错误，但在这里做的是正确的，因为ServeHTTP是唯一看到该值并使用其内容的地方)。

And make `appHandler`’s `ServeHTTP` method display the `appError`’s `Message` to the user with the correct HTTP status `Code` and log the full `Error` to the developer console:

并使appHandler的ServeHTTP方法以正确的HTTP状态码向用户显示appError的信息，并将完整的Error记录到开发者控制台。

```go linenums="1"
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if e := fn(w, r); e != nil { // e is *appError, not os.Error.
        c := appengine.NewContext(r)
        c.Errorf("%v", e.Error)
        http.Error(w, e.Message, e.Code)
    }
}
```

Finally, we update `viewRecord` to the new function signature and have it return more context when it encounters an error:

最后，我们将viewRecord更新为新的函数签名，让它在遇到错误时返回更多的上下文。

```go linenums="1"
func viewRecord(w http.ResponseWriter, r *http.Request) *appError {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
    record := new(Record)
    if err := datastore.Get(c, key, record); err != nil {
        return &appError{err, "Record not found", 404}
    }
    if err := viewTemplate.Execute(w, record); err != nil {
        return &appError{err, "Can't display record", 500}
    }
    return nil
}
```

This version of `viewRecord` is the same length as the original, but now each of those lines has specific meaning and we are providing a friendlier user experience.

这个版本的viewRecord与原来的长度相同，但现在每一行都有特定的含义，我们提供了一个更友好的用户体验。

It doesn’t end there; we can further improve the error handling in our application. Some ideas:

这还没有结束，我们可以进一步改善我们应用程序中的错误处理。一些想法。

- give the error handler a pretty HTML template, 给错误处理程序一个漂亮的HTML模板。
- make debugging easier by writing the stack trace to the HTTP response when the user is an administrator, 当用户是管理员时，将堆栈跟踪写到HTTP响应中，使调试更容易。
- write a constructor function for `appError` that stores the stack trace for easier debugging, 为appError写一个构造函数，存储堆栈跟踪，以方便调试。
- recover from panics inside the `appHandler`, logging the error to the console as "Critical," while telling the user "a serious error has occurred." This is a nice touch to avoid exposing the user to inscrutable error messages caused by programming errors. See the [Defer, Panic, and Recover](https://go.dev/doc/articles/defer_panic_recover.html) article for more details.在appHandler中恢复恐慌，将错误记录到控制台，称为 "关键"，同时告诉用户 "发生了一个严重的错误"。这是一个很好的提示，可以避免用户暴露在由编程错误引起的难以捉摸的错误信息中。更多细节请参见《延缓、恐慌和恢复》一文。

## Conclusion 结论

Proper error handling is an essential requirement of good software. By employing the techniques described in this post you should be able to write more reliable and succinct Go code.

正确的错误处理是优秀软件的一个基本要求。通过采用本篇文章中描述的技术，您应该能够写出更可靠和简洁的Go代码。
