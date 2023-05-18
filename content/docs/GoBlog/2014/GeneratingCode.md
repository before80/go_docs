+++
title = "生成代码"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Generating code - 生成代码

https://go.dev/blog/generate

Rob Pike
22 December 2014

A property of universal computation—Turing completeness—is that a computer program can write a computer program. This is a powerful idea that is not appreciated as often as it might be, even though it happens frequently. It’s a big part of the definition of a compiler, for instance. It’s also how the `go` `test` command works: it scans the packages to be tested, writes out a Go program containing a test harness customized for the package, and then compiles and runs it. Modern computers are so fast this expensive-sounding sequence can complete in a fraction of a second.

通用计算的一个属性--图灵完备性--是指一个计算机程序可以编写一个计算机程序。这是一个强大的想法，尽管它经常发生，但却没有得到重视。例如，它是编译器定义的一个重要部分。这也是go测试命令的工作方式：它扫描要测试的软件包，写出一个包含为软件包定制的测试线束的Go程序，然后编译并运行它。现代计算机的速度非常快，这个听起来很昂贵的序列可以在几分之一秒内完成。

There are lots of other examples of programs that write programs. [Yacc](https://godoc.org/golang.org/x/tools/cmd/goyacc), for instance, reads in a description of a grammar and writes out a program to parse that grammar. The protocol buffer “compiler” reads an interface description and emits structure definitions, methods, and other support code. Configuration tools of all sorts work like this too, examining metadata or the environment and emitting scaffolding customized to the local state.

还有很多其他写程序的例子。例如，Yacc读入一个语法描述，并写出一个程序来解析该语法。协议缓冲区 "编译器 "读取一个接口描述，并发出结构定义、方法和其他支持代码。各种各样的配置工具也是这样工作的，检查元数据或环境，并发出根据本地状态定制的脚手架。

Programs that write programs are therefore important elements in software engineering, but programs like Yacc that produce source code need to be integrated into the build process so their output can be compiled. When an external build tool like Make is being used, this is usually easy to do. But in Go, whose go tool gets all necessary build information from the Go source, there is a problem. There is simply no mechanism to run Yacc from the go tool alone.

因此，编写程序的程序是软件工程中的重要元素，但像Yacc这样产生源代码的程序需要被集成到构建过程中，以便它们的输出可以被编译。当使用Make这样的外部构建工具时，这通常很容易做到。但在Go中，其go工具从Go源码中获取所有必要的构建信息，这就存在一个问题。根本没有任何机制可以单独从go工具中运行Yacc。

Until now, that is.

直到现在，就是这样。

The [latest Go release](https://blog.golang.org/go1.4), 1.4, includes a new command that makes it easier to run such tools. It’s called `go` `generate`, and it works by scanning for special comments in Go source code that identify general commands to run. It’s important to understand that `go` `generate` is not part of `go` `build`. It contains no dependency analysis and must be run explicitly before running `go` `build`. It is intended to be used by the author of the Go package, not its clients.

最新的Go版本，1.4，包含了一个新的命令，使得运行这样的工具变得更加容易。它被称为go generate，它通过扫描Go源代码中的特殊注释来识别要运行的一般命令。重要的是要明白，go generate不是go build的一部分。它不包含依赖性分析，必须在运行go build之前明确运行。它的目的是供Go包的作者使用，而不是其客户。

The `go` `generate` command is easy to use. As a warmup, here’s how to use it to generate a Yacc grammar.

go generate 命令很容易使用。作为热身，下面介绍如何使用它来生成Yacc语法。

First, install Go’s Yacc tool:

首先，安装Go的Yacc工具：

```
go get golang.org/x/tools/cmd/goyacc
```

Say you have a Yacc input file called `gopher.y` that defines a grammar for your new language. To produce the Go source file implementing the grammar, you would normally invoke the command like this:

假设你有一个名为gopher.y的Yacc输入文件，为你的新语言定义了一个语法。为了生成实现该语法的Go源代码文件，你通常会调用这样的命令：

```
goyacc -o gopher.go -p parser gopher.y
```

The `-o` option names the output file while `-p` specifies the package name.

-o选项命名了输出文件，而-p则指定了软件包的名称。

To have `go` `generate` drive the process, in any one of the regular (non-generated) `.go` files in the same directory, add this comment anywhere in the file:

要让go生成驱动这个过程，在同一目录下的任何一个常规（非生成）的.go文件中，在文件的任何地方添加这个注释：

```
//go:generate goyacc -o gopher.go -p parser gopher.y
```

This text is just the command above prefixed by a special comment recognized by `go` `generate`. The comment must start at the beginning of the line and have no spaces between the `//` and the `go:generate`. After that marker, the rest of the line specifies a command for `go` `generate` to run.

这段文字只是上面的命令，前面有一个go generate所识别的特殊注释。该注释必须从该行的开头开始，并且在//和go:generate之间没有空格。在这个标记之后，该行的其余部分指定了一个让go generate运行的命令。

Now run it. Change to the source directory and run `go` `generate`, then `go` `build` and so on:

现在运行它。换到源码目录，运行go generate，然后运行go build，依次进行：

```shell linenums="1"
$ cd $GOPATH/myrepo/gopher
$ go generate
$ go build
$ go test
```

That’s it. Assuming there are no errors, the `go` `generate` command will invoke `yacc` to create `gopher.go`, at which point the directory holds the full set of Go source files, so we can build, test, and work normally. Every time `gopher.y` is modified, just rerun `go` `generate` to regenerate the parser.

这就是了。假设没有错误，go generate命令会调用yacc来创建gopher.go，这时目录里有全套的Go源文件，所以我们可以构建、测试，并正常工作。每次修改gopher.y，只需重新运行go generate来重新生成解析器。

For more details about how `go` `generate` works, including options, environment variables, and so on, see the [design document](https://go.dev/s/go1.4-generate).

关于 go generate 如何工作的更多细节，包括选项、环境变量等等，请参阅设计文档。

Go generate does nothing that couldn’t be done with Make or some other build mechanism, but it comes with the `go` tool—no extra installation required—and fits nicely into the Go ecosystem. Just keep in mind that it is for package authors, not clients, if only for the reason that the program it invokes might not be available on the target machine. Also, if the containing package is intended for import by `go` `get`, once the file is generated (and tested!) it must be checked into the source code repository to be available to clients.

Go generate并没有做任何不能用Make或其他构建机制完成的事情，但它是Go工具自带的，不需要额外的安装，很适合Go的生态系统。请记住，它是为包的作者而不是客户准备的，原因是它调用的程序可能在目标机器上无法使用。另外，如果包含的包是打算由go get导入的，一旦文件生成（并经过测试！），它必须被检查到源代码库中，以便对客户可用。

Now that we have it, let’s use it for something new. As a very different example of how `go` `generate` can help, there is a new program available in the `golang.org/x/tools` repository called `stringer`. It automatically writes string methods for sets of integer constants. It’s not part of the released distribution, but it’s easy to install:

现在我们有了它，让我们把它用于新的东西。作为go generate的一个非常不同的例子，golang.org/x/tools仓库里有一个新的程序，叫做stringer。它为整数常量的集合自动编写字符串方法。它不是已发布版本的一部分，但它很容易安装：

```shell linenums="1"
$ go get golang.org/x/tools/cmd/stringer
```

Here’s an example from the documentation for [`stringer`](https://godoc.org/golang.org/x/tools/cmd/stringer). Imagine we have some code that contains a set of integer constants defining different types of pills:

下面是stringer文档中的一个例子。想象一下，我们有一些代码，其中包含一组定义不同类型药片的整数常量：

```go linenums="1"
package painkiller

type Pill int

const (
    Placebo Pill = iota
    Aspirin
    Ibuprofen
    Paracetamol
    Acetaminophen = Paracetamol
)
```

For debugging, we’d like these constants to pretty-print themselves, which means we want a method with signature,

为了调试，我们希望这些常量能漂亮地打印出自己，这意味着我们要有一个带签名的方法，

```go linenums="1"
func (p Pill) String() string
```

It’s easy to write one by hand, perhaps like this:

用手写一个方法是很容易的，也许像这样：

```go linenums="1"
func (p Pill) String() string {
    switch p {
    case Placebo:
        return "Placebo"
    case Aspirin:
        return "Aspirin"
    case Ibuprofen:
        return "Ibuprofen"
    case Paracetamol: // == Acetaminophen
        return "Paracetamol"
    }
    return fmt.Sprintf("Pill(%d)", p)
}
```

There are other ways to write this function, of course. We could use a slice of strings indexed by Pill, or a map, or some other technique. Whatever we do, we need to maintain it if we change the set of pills, and we need to make sure it’s correct. (The two names for paracetamol make this trickier than it might otherwise be.) Plus the very question of which approach to take depends on the types and values: signed or unsigned, dense or sparse, zero-based or not, and so on.

当然，还有其他方法来写这个函数。我们可以使用一个以Pill为索引的字符串切片，或一个地图，或其他一些技术。不管我们怎么做，如果我们改变了药片的集合，我们需要维护它，而且我们需要确保它是正确的。(扑热息痛的两个名字使得这个问题比它可能的更棘手)。另外，采取哪种方法的问题本身取决于类型和数值：有符号或无符号，密集或稀疏，是否基于零，等等。

The `stringer` program takes care of all these details. Although it can be run in isolation, it is intended to be driven by `go` `generate`. To use it, add a generate comment to the source, perhaps near the type definition:

stringer程序处理了所有这些细节。尽管它可以单独运行，但它是由go generate驱动的。要使用它，请在源代码中添加一个生成注释，也许在类型定义附近：

```
//go:generate stringer -type=Pill
```

This rule specifies that `go` `generate` should run the `stringer` tool to generate a `String` method for type `Pill`. The output is automatically written to `pill_string.go` (a default we could override with the `-output` flag).

这条规则指定了go generate应该运行stringer工具来为Pill类型生成一个String方法。输出被自动写入 pill_string.go（我们可以用 -output 标志来覆盖这个默认值）。

Let’s run it:

让我们来运行它：

```shell linenums="1"
$ go generate
$ cat pill_string.go
// Code generated by stringer -type Pill pill.go; DO NOT EDIT.

package painkiller

import "fmt"

const _Pill_name = "PlaceboAspirinIbuprofenParacetamol"

var _Pill_index = [...]uint8{0, 7, 14, 23, 34}

func (i Pill) String() string {
    if i < 0 || i+1 >= Pill(len(_Pill_index)) {
        return fmt.Sprintf("Pill(%d)", i)
    }
    return _Pill_name[_Pill_index[i]:_Pill_index[i+1]]
}
$
```

Every time we change the definition of `Pill` or the constants, all we need to do is run

每次我们改变Pill的定义或常量时，我们需要做的就是运行

```
$ go generate
```

to update the `String` method. And of course if we’ve got multiple types set up this way in the same package, that single command will update all their `String` methods with a single command.

来更新String方法。当然，如果我们在同一个包中以这种方式设置了多个类型，那么这条命令将以单一的命令更新所有的String方法。

There’s no question the generated method is ugly. That’s OK, though, because humans don’t need to work on it; machine-generated code is often ugly. It’s working hard to be efficient. All the names are smashed together into a single string, which saves memory (only one string header for all the names, even if there are zillions of them). Then an array, `_Pill_index`, maps from value to name by a simple, efficient technique. Note too that `_Pill_index` is an array (not a slice; one more header eliminated) of `uint8`, the smallest integer sufficient to span the space of values. If there were more values, or there were negatives ones, the generated type of `_Pill_index` might change to `uint16` or `int8`: whatever works best.

毫无疑问，生成的方法是丑陋的。不过这没关系，因为人类不需要在上面工作；机器生成的代码往往是丑陋的。它正在努力工作以提高效率。所有的名字都被打碎成一个字符串，这样可以节省内存（所有的名字只有一个字符串头，即使有几十亿个名字）。然后，一个数组，_Pill_index，通过一个简单而有效的技术从值映射到名字。还要注意的是，_Pill_index是一个uint8的数组（不是一个片断；多了一个标题），这个最小的整数足以横跨整个数值空间。如果有更多的值，或者有负数，_Pill_index的生成类型可能会变成uint16或者int8：不管是什么，都是最好的。

The approach used by the methods printed by `stringer` varies according to the properties of the constant set. For instance, if the constants are sparse, it might use a map. Here’s a trivial example based on a constant set representing powers of two:

stringer打印的方法所使用的方法根据常量集的属性而不同。例如，如果常数是稀疏的，它可能会使用一个映射。下面是一个基于代表2的幂的常数集的微不足道的例子：

```go linenums="1"
const _Power_name = "p0p1p2p3p4p5..."

var _Power_map = map[Power]string{
    1:    _Power_name[0:2],
    2:    _Power_name[2:4],
    4:    _Power_name[4:6],
    8:    _Power_name[6:8],
    16:   _Power_name[8:10],
    32:   _Power_name[10:12],
    ...,
}

func (i Power) String() string {
    if str, ok := _Power_map[i]; ok {
        return str
    }
    return fmt.Sprintf("Power(%d)", i)
}
```

In short, generating the method automatically allows us to do a better job than we would expect a human to do.

简而言之，自动生成方法使我们能够做得比我们期望人类做的更好。

There are lots of other uses of `go` `generate` already installed in the Go tree. Examples include generating Unicode tables in the `unicode` package, creating efficient methods for encoding and decoding arrays in `encoding/gob`, producing time zone data in the `time` package, and so on.

在Go树上已经安装了go generate，还有很多其他用途。例如，在unicode包中生成Unicode表，在encoding/gob中创建高效的数组编码和解码方法，在time包中生成时区数据，等等。

Please use `go` `generate` creatively. It’s there to encourage experimentation.

请创造性地使用go generate。它的存在是为了鼓励实验。

And even if you don’t, use the new `stringer` tool to write your `String` methods for your integer constants. Let the machine do the work.

即使你不这样做，也请使用新的stringer工具为你的整数常量编写字符串方法。让机器来做这些工作。
