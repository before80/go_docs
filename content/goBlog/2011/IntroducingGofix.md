+++
title = "介绍一下 gofix"
weight = 23
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Introducing Gofix - 介绍一下 gofix

https://go.dev/blog/introducing-gofix

Russ Cox
15 April 2011

The next Go release will include significant API changes in several fundamental Go packages. Code that [implements an HTTP server handler](http://codereview.appspot.com/4239076), [calls `net.Dial`](http://codereview.appspot.com/4244055), [calls `os.Open`](http://codereview.appspot.com/4357052), or [uses the reflect package](http://codereview.appspot.com/4281055) will not build unless it is updated to use the new APIs. Now that our releases are [more stable and less frequent](https://blog.golang.org/2011/03/go-becomes-more-stable.html), this will be a common situation. Each of these API changes happened in a different weekly snapshot and might have been manageable on its own; together, however, they represent a significant amount of manual effort to update existing code.

下一个Go版本将包括几个基本Go包的重大API变化。实现HTTP服务器处理程序、调用net.Dial、调用os.Open或使用reflect包的代码将无法构建，除非它被更新为使用新的API。现在我们的发布更加稳定，频率也更低，这将是一种常见的情况。这些API的变化都发生在不同的周快照中，可能本身是可以管理的；但是，它们加在一起，代表了大量的手动工作来更新现有的代码。

[Gofix](https://go.dev/cmd/fix/) is a new tool that reduces the amount of effort it takes to update existing code. It reads a program from a source file, looks for uses of old APIs, rewrites them to use the current API, and writes the program back to the file. Not all API changes preserve all the functionality of an old API, so gofix cannot always do a perfect job. When gofix cannot rewrite a use of an old API, it prints a warning giving the file name and line number of the use, so that a developer can examine and rewrite the code. Gofix takes care of the easy, repetitive, tedious changes, so that a developer can focus on the ones that truly merit attention.

Gofix是一个新的工具，它减少了更新现有代码的工作量。它从源文件中读取程序，寻找旧API的用途，将其改写为使用当前的API，并将程序写回文件中。并非所有的API变化都能保留旧API的所有功能，所以gofix不能总是做得很完美。当gofix不能重写一个旧API的使用时，它会打印出一个警告，给出使用的文件名和行号，以便开发者可以检查并重写代码。Gofix负责处理那些简单的、重复的、繁琐的修改，这样开发者就可以专注于那些真正值得关注的修改。

Each time we make a significant API change we’ll add code to gofix to take care of the conversion, as much as mechanically possible. When you update to a new Go release and your code no longer builds, just run gofix on your source directory.

每次我们做出重大的API改变时，我们都会在gofix中添加代码，以尽可能机械地处理转换问题。当您更新到一个新的Go版本，您的代码不再构建时，只需在您的源代码目录上运行gofix。

You can extend gofix to support changes to your own APIs. The gofix program is a simple driver around plugins called fixes that each handle a particular API change. Right now, writing a new fix requires doing some scanning and rewriting of the go/ast syntax tree, usually in proportion to how complex the API changes are. If you want to explore, the [`netdialFix`](https://go.googlesource.com/go/+/go1/src/cmd/fix/netdial.go), [`osopenFix`](https://go.googlesource.com/go/+/go1/src/cmd/fix/osopen.go), [`httpserverFix`](https://go.googlesource.com/go/+/go1/src/cmd/fix/httpserver.go), and [`reflectFix`](https://go.googlesource.com/go/+/go1/src/cmd/fix/reflect.go) are all illustrative examples, in increasing order of complexity.

您可以扩展gofix以支持您自己的API的变化。gofix程序是一个简单的驱动，围绕着称为fix的插件，每个插件都处理一个特定的API变化。现在，编写一个新的fix需要对go/ast语法树做一些扫描和重写，通常与API变化的复杂程度成正比。如果您想探索，netdialFix、osopenFix、httpserverFix和reflectFix都是说明性的例子，复杂程度依次递增。

We write Go code too, of course, and our code is just as affected by these API changes as yours. Typically, we write the gofix support at the same time as the API change and then use gofix to rewrite the uses in the main source tree. We use gofix to update other Go code bases and our personal projects. We even use gofix to update Google’s internal source tree when it is time to build against a new Go release.

当然，我们也写Go代码，我们的代码和您的代码一样会受到这些API变化的影响。通常情况下，我们在API变化的同时编写gofix支持，然后用gofix重写主源码树中的用途。我们用gofix来更新其他Go代码库和我们的个人项目。当需要针对新的Go版本进行构建时，我们甚至使用gofix来更新谷歌的内部源代码树。

As an example, gofix can rewrite code like [this snippet from `fmt/print.go`](http://codereview.appspot.com/4353043/diff/10001/src/pkg/fmt/print.go#newcode657):

举个例子，gofix可以重写像fmt/print.go中的这段代码：

```go
switch f := value.(type) {
case *reflect.BoolValue:
    p.fmtBool(f.Get(), verb, field)
case *reflect.IntValue:
    p.fmtInt64(f.Get(), verb, field)
// ...
case reflect.ArrayOrSliceValue:
    // Byte slices are special.
    if f.Type().(reflect.ArrayOrSliceType).Elem().Kind() == reflect.Uint8 {
        // ...
    }
// ...
}
```

to adapt it to the new reflect API:

以使其适应新的reflect API：

```go
switch f := value; f.Kind() {
case reflect.Bool:
    p.fmtBool(f.Bool(), verb, field)
case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
    p.fmtInt64(f.Int(), verb, field)
// ...
case reflect.Array, reflect.Slice:
    // Byte slices are special.
    if f.Type().Elem().Kind() == reflect.Uint8 {
        // ...
    }
// ...
}
```

Nearly every line above changed in some small way. The changes involved in the rewrite are extensive but nearly entirely mechanical, just the kind of thing that computers are great at doing.

上面的每一行几乎都有一些细微的变化。重写所涉及的变化非常广泛，但几乎完全是机械性的，这正是计算机所擅长的事情。

Gofix is possible because Go has support in its standard libraries for [parsing Go source files into syntax trees](https://go.dev/pkg/go/parser) and also for [printing those syntax trees back to Go source code](https://go.dev/pkg/go/printer). Importantly, the Go printing library prints a program in the official format (typically enforced via the gofmt tool), allowing gofix to make mechanical changes to Go programs without causing spurious formatting changes. In fact, one of the key motivations for creating gofmt—perhaps second only to avoiding debates about where a particular brace belongs—was to simplify the creation of tools that rewrite Go programs, as gofix does.

Gofix之所以能够实现，是因为Go的标准库支持将Go源文件解析为语法树，并将这些语法树打印为Go源代码。重要的是，Go打印库以官方格式打印程序（通常通过gofmt工具强制执行），允许gofix对Go程序进行机械性修改，而不会引起虚假的格式变化。事实上，创建gofmt的主要动机之一--也许仅次于避免关于特定括号归属的争论--是为了简化重写Go程序的工具的创建，正如gofix所做的那样。

Gofix has already made itself indispensable. In particular, the recent reflect changes would have been unpalatable without automated conversion, and the reflect API badly needed to be redone. Gofix gives us the ability to fix mistakes or completely rethink package APIs without worrying about the cost of converting existing code. We hope you find gofix as useful and convenient as we have.

Gofix已经让自己变得不可或缺了。特别是最近的reflect变化，如果没有自动转换，就会变得很难受，而reflect的API也亟需重做。Gofix让我们有能力修复错误或完全重新思考包的API，而不必担心转换现有代码的成本。我们希望您发现gofix和我们一样有用和方便。
