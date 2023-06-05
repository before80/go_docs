+++
title = "FAQ（中英对照）"
weight = 100
date = 2023-05-18T16:56:22+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Frequently Asked Questions (FAQ)

> 原文：[https://go.dev/doc/faq](https://go.dev/doc/faq)

## Origins 起源

### What is the purpose of the project? 这个项目的目的是什么？

At the time of Go's inception, only a decade ago, the programming world was different from today. Production software was usually written in C++ or Java, GitHub did not exist, most computers were not yet multiprocessors, and other than Visual Studio and Eclipse there were few IDEs or other high-level tools available at all, let alone for free on the Internet.

​	在Go诞生的时候，也就是十年前，编程世界与今天不同。生产软件通常是用C++或Java编写的，GitHub还不存在，大多数计算机还不是多处理器，除了Visual Studio和Eclipse之外，几乎没有什么IDE或其他高级工具可用，更不用说在互联网上免费使用了。

Meanwhile, we had become frustrated by the undue complexity required to use the languages we worked with to develop server software. Computers had become enormously quicker since languages such as C, C++ and Java were first developed but the act of programming had not itself advanced nearly as much. Also, it was clear that multiprocessors were becoming universal but most languages offered little help to program them efficiently and safely.

​	与此同时，我们对使用我们所使用的语言来开发服务器软件所需要的过度复杂性感到沮丧。自从C、C++和Java等语言被首次开发以来，计算机已经变得非常快了，但是编程行为本身并没有得到同样的发展。另外，很明显，多处理器正在普及，但大多数语言对它们的有效和安全编程没有提供什么帮助。

We decided to take a step back and think about what major issues were going to dominate software engineering in the years ahead as technology developed, and how a new language might help address them. For instance, the rise of multicore CPUs argued that a language should provide first-class support for some sort of concurrency or parallelism. And to make resource management tractable in a large concurrent program, garbage collection, or at least some sort of safe automatic memory management was required.

​	我们决定退一步思考，随着技术的发展，哪些主要问题将在未来几年主导软件工程，以及一种新的语言如何帮助解决这些问题。例如，多核CPU的兴起表明，一种语言应该为某种并发性或并行性提供一流的支持。为了使大型并发程序中的资源管理变得可行，需要有垃圾收集，或者至少是某种安全的自动内存管理。

These considerations led to [a series of discussions](https://commandcenter.blogspot.com/2017/09/go-ten-years-and-climbing.html) from which Go arose, first as a set of ideas and desiderata, then as a language. An overarching goal was that Go do more to help the working programmer by enabling tooling, automating mundane tasks such as code formatting, and removing obstacles to working on large code bases.

​	这些考虑导致了[一系列的讨论](https://commandcenter.blogspot.com/2017/09/go-ten-years-and-climbing.html)，Go就是在这些讨论中产生的，首先是一系列的想法和愿望，然后是作为一种语言。一个首要的目标是，Go要更多地帮助工作中的程序员，使工具化，使代码格式化等日常任务自动化，并消除在大型代码库中工作的障碍。

A much more expansive description of the goals of Go and how they are met, or at least approached, is available in the article, [Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article).

​	关于Go的目标以及如何实现或至少如何接近这些目标的更详尽描述，可以在[Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)中找到。

### What is the history of the project? 这个项目的历史是怎样的？

Robert Griesemer, Rob Pike and Ken Thompson started sketching the goals for a new language on the white board on September 21, 2007. Within a few days the goals had settled into a plan to do something and a fair idea of what it would be. Design continued part-time in parallel with unrelated work. By January 2008, Ken had started work on a compiler with which to explore ideas; it generated C code as its output. By mid-year the language had become a full-time project and had settled enough to attempt a production compiler. In May 2008, Ian Taylor independently started on a GCC front end for Go using the draft specification. Russ Cox joined in late 2008 and helped move the language and libraries from prototype to reality.

​	2007年9月21日，Robert Griesemer、Rob Pike和Ken Thompson开始在白板上勾画新语言的目标。在几天之内，这些目标就变成了一个做事的计划和一个关于它将是什么的合理想法。设计工作继续与无关的工作并行。到2008年1月，Ken已经开始了编译器的工作，用它来探索想法；它的输出是C代码。到了年中，该语言已经成为一个全职项目，并且已经稳定下来，可以尝试生产一个编译器。2008年5月，Ian Taylor利用规范草案独立开始了Go的GCC前端工作。Russ Cox在2008年底加入，并帮助将语言和库从原型变为现实。

Go became a public open source project on November 10, 2009. Countless people from the community have contributed ideas, discussions, and code.

​	Go在2009年11月10日成为一个公开的开源项目。社区中无数人贡献了想法、讨论和代码。

There are now millions of Go programmers—gophers—around the world, and there are more every day. Go's success has far exceeded our expectations.

​	现在全世界有数以百万计的Go程序员——gophers，而且每天都有更多。Go的成功已经远远超过了我们的预期。

### What's the origin of the gopher mascot? 囊地鼠吉祥物的起源是什么？

The mascot and logo were designed by [Renée French](https://reneefrench.blogspot.com/), who also designed [Glenda](https://9p.io/plan9/glenda.html), the Plan 9 bunny. A [blog post](https://blog.golang.org/gopher) about the gopher explains how it was derived from one she used for a [WFMU](https://wfmu.org/) T-shirt design some years ago. The logo and mascot are covered by the [Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0/) license.

​	吉祥物和标志是由[Renée French](https://reneefrench.blogspot.com/)设计的，她还设计了Plan 9的兔子[Glenda](https://9p.io/plan9/glenda.html)。一篇[关于囊地鼠的博客]({{< ref "/goBlog/2014/TheGoGopher" >}})文章解释了它是如何从她几年前用于[WFMU](https://wfmu.org/)T恤设计的一个囊地鼠衍生出来的。该标志和吉祥物属于[Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0/)许可范围。

The gopher has a [model sheet](https://go.dev/doc/gopher/modelsheet.jpg) illustrating his characteristics and how to represent them correctly. The model sheet was first shown in a [talk](https://www.youtube.com/watch?v=4rw_B4yY69k) by Renée at Gophercon in 2016. He has unique features; he's the *Go gopher*, not just any old gopher.

​	囊地鼠有一个[模型表](https://go.dev/doc/gopher/modelsheet.jpg)，说明他的特点和如何正确地表现它们。该模型表在Renée于2016年在Gophercon的[演讲](https://www.youtube.com/watch?v=4rw_B4yY69k)中首次展示。他有独特的特征；他是Go的囊地鼠，不是普通的囊地鼠。

### Is the language called Go or Golang? 这门语言叫Go还是Golang？

The language is called Go. The "golang" moniker arose because the web site was originally *golang.org*. (There was no *.dev* domain then.) Many use the golang name, though, and it is handy as a label. For instance, the Twitter tag for the language is "#golang". The language's name is just plain Go, regardless.

​	这门语言叫Go。出现 "golang "这个名字是因为网站最初是golang.org。(那时还没有.dev域名。)不过很多人都使用golang这个名字，而且它作为一个标签很方便。例如，该语言的Twitter标签是 "#golang"。但无论如何，该语言的名字只是普通的Go。

A side note: Although the [official logo](https://blog.golang.org/go-brand) has two capital letters, the language name is written Go, not GO.

​	题外话：虽然[官方标志]({{< ref "/goBlog/2018/GosNewBrand" >}})有两个大写字母，但语言的名字是Go，而不是GO。

### Why did you create a new language? 您为什么要创造一种新的语言？

Go was born out of frustration with existing languages and environments for the work we were doing at Google. Programming had become too difficult and the choice of languages was partly to blame. One had to choose either efficient compilation, efficient execution, or ease of programming; all three were not available in the same mainstream language. Programmers who could were choosing ease over safety and efficiency by moving to dynamically typed languages such as Python and JavaScript rather than C++ or, to a lesser extent, Java.

​	Go的诞生是出于对现有语言和环境的失望，我们在谷歌所做的工作。编程已经变得太困难了，语言的选择也是部分原因。人们不得不选择高效的编译、高效的执行，或者方便的编程；这三者在同一种主流语言中都不存在。有能力的程序员选择了轻松而不是安全和效率，他们转向动态类型的语言，如Python和JavaScript，而不是C++，或者在较小程度上使用Java。

We were not alone in our concerns. After many years with a pretty quiet landscape for programming languages, Go was among the first of several new languages—Rust, Elixir, Swift, and more—that have made programming language development an active, almost mainstream field again.

​	我们并不孤单。在经历了多年的编程语言领域的平静之后，Go是几种新语言中的第一批——Rust、Elixir、Swift等等——它们使编程语言的开发再次成为一个活跃的、几乎是主流的领域。

Go addressed these issues by attempting to combine the ease of programming of an interpreted, dynamically typed language with the efficiency and safety of a statically typed, compiled language. It also aimed to be modern, with support for networked and multicore computing. Finally, working with Go is intended to be *fast*: it should take at most a few seconds to build a large executable on a single computer. To meet these goals required addressing a number of linguistic issues: an expressive but lightweight type system; concurrency and garbage collection; rigid dependency specification; and so on. These cannot be addressed well by libraries or tools; a new language was called for.

​	Go解决了这些问题，它试图将解释型、动态类型语言的编程便利性与静态类型、编译型语言的效率和安全性相结合。它还旨在成为现代语言，支持网络和多核计算。最后，使用Go的目的是快速：在一台计算机上构建一个大的可执行文件最多只需要几秒钟。为了实现这些目标，需要解决一些语言问题：一个有表现力但轻量级的类型系统；并发性和垃圾收集；严格的依赖项规范等等。这些问题不能由库或工具来解决；需要一种新的语言。

The article [Go at Google](https://go.dev/talks/2012/splash.article) discusses the background and motivation behind the design of the Go language, as well as providing more detail about many of the answers presented in this FAQ.

​	[Go at Google](https://go.dev/talks/2012/splash.article)的文章讨论了Go语言设计的背景和动机，并对本FAQ中的许多答案提供了更多细节。

### What are Go's ancestors? Go 的祖先是什么？

Go is mostly in the C family (basic syntax), with significant input from the Pascal/Modula/Oberon family (declarations, packages), plus some ideas from languages inspired by Tony Hoare's CSP, such as Newsqueak and Limbo (concurrency). However, it is a new language across the board. In every respect the language was designed by thinking about what programmers do and how to make programming, at least the kind of programming we do, more effective, which means more fun.

​	Go主要属于C语言家族（基本语法），有来自Pascal/Modula/Oberon家族（声明、包）的重要输入，还有来自Tony Hoare的CSP语言的一些想法，如Newsqueak和Limbo（并发）。然而，它是一种全面的新语言。在每一个方面，这门语言的设计都是通过思考程序员的工作以及如何使编程，至少是我们所做的那种编程，更加有效，也就是更加有趣。

### What are the guiding principles in the design? 设计中的指导原则是什么？

When Go was designed, Java and C++ were the most commonly used languages for writing servers, at least at Google. We felt that these languages required too much bookkeeping and repetition. Some programmers reacted by moving towards more dynamic, fluid languages like Python, at the cost of efficiency and type safety. We felt it should be possible to have the efficiency, the safety, and the fluidity in a single language.

​	在设计Go的时候，Java和C++是编写服务器最常用的语言，至少在Google是这样。我们觉得这些语言需要太多的记账和重复。一些程序员的反应是转向像Python这样更动态、更流畅的语言，但却牺牲了效率和类型安全。我们认为应该可以在一种语言中实现效率、安全和流畅性。

Go attempts to reduce the amount of typing in both senses of the word. Throughout its design, we have tried to reduce clutter and complexity. There are no forward declarations and no header files; everything is declared exactly once. Initialization is expressive, automatic, and easy to use. Syntax is clean and light on keywords. Repetition (`foo.Foo* myFoo = new(foo.Foo)`) is reduced by simple type derivation using the `:=` declare-and-initialize construct. And perhaps most radically, there is no type hierarchy: types just *are*, they don't have to announce their relationships. These simplifications allow Go to be expressive yet comprehensible without sacrificing, well, sophistication.

​	Go试图在这两种意义上减少输入的数量。在整个设计过程中，我们试图减少混乱和复杂性。没有前向声明，也没有头文件；所有东西都只声明一次。初始化是明确的、自动的，并且易于使用。语法简洁，关键字少。重复（`foo.Foo* myFoo = new(foo.Foo)`）通过使用 `:=` declare-and-initialize 结构进行简单的类型推导来减少。也许最根本的是，没有类型的层次结构：类型就是这样，它们不需要宣布它们的关系。这些简化使得Go在不牺牲复杂度的前提下实现了表现力和可理解性。

Another important principle is to keep the concepts orthogonal. Methods can be implemented for any type; structures represent data while interfaces represent abstraction; and so on. Orthogonality makes it easier to understand what happens when things combine.

​	另一个重要原则是保持概念的正交性。`方法可以为任何类型实现；结构代表数据而接口代表抽象；`等等。正交性使我们更容易理解事物结合时的情况。

## Usage 使用方法

### Is Google using Go internally? Google在内部使用Go吗？

Yes. Go is used widely in production inside Google. One easy example is the server behind [golang.org](https://golang.org/). It's just the [`godoc`](https://go.dev/cmd/godoc) document server running in a production configuration on [Google App Engine](https://developers.google.com/appengine/).

​	是的。Go 在 Google 内部的生产中被广泛使用。一个简单的例子是[golang.org](https://golang.org/)背后的服务器。它只是在 [Google App Engine](https://developers.google.com/appengine/) 上以生产配置运行的 [godoc](https://go.dev/cmd/godoc) 文档服务器。

A more significant instance is Google's download server, `dl.google.com`, which delivers Chrome binaries and other large installables such as `apt-get` packages.

​	一个更重要的例子是Google的下载服务器`dl.google.com`，它提供Chrome二进制文件和其他大型安装程序，如`apt-get`包。

Go is not the only language used at Google, far from it, but it is a key language for a number of areas including [site reliability engineering (SRE)](https://go.dev/talks/2013/go-sreops.slide) and large-scale data processing.

​	Go不是谷歌唯一使用的语言，远非如此，但它是包括[网站可靠性工程（SRE）](https://go.dev/talks/2013/go-sreops.slide)和大规模数据处理在内的许多领域的关键语言。

### What other companies use Go? 还有哪些公司使用Go？

Go usage is growing worldwide, especially but by no means exclusively in the cloud computing space. A couple of major cloud infrastructure projects written in Go are Docker and Kubernetes, but there are many more.

​	Go的使用在全球范围内不断增长，尤其是在云计算领域，但绝非仅限于此。`Docker` 和 `Kubernetes` 是用 Go 编写的几个主要的云计算基础设施项目，但还有很多。

It's not just cloud, though. The Go Wiki includes a [page](https://github.com/golang/go/wiki/GoUsers), updated regularly, that lists some of the many companies using Go.

​	不过，这不仅仅是云计算。Go Wiki有一个[页面](https://github.com/golang/go/wiki/GoUsers)，定期更新，其中列出了一些使用Go的公司。

The Wiki also has a page with links to [success stories](https://github.com/golang/go/wiki/SuccessStories) about companies and projects that are using the language.

​	Wiki还有一个页面，链接到使用该语言的公司和项目的[成功案例](https://github.com/golang/go/wiki/SuccessStories)。

### Do Go programs link with C/C++ programs? Go程序可以与C/C++程序链接吗？

It is possible to use C and Go together in the same address space, but it is not a natural fit and can require special interface software. Also, linking C with Go code gives up the memory safety and stack management properties that Go provides. Sometimes it's absolutely necessary to use C libraries to solve a problem, but doing so always introduces an element of risk not present with pure Go code, so do so with care.

​	可以在同一地址空间中同时使用 C 和 Go，但这并不是一种自然的配合，可能需要特殊的接口软件。另外，将 C 与 Go 代码连接起来会放弃 Go 提供的内存安全和堆栈管理特性。有时绝对有必要使用C语言库来解决问题，但这样做总是会引入纯Go代码所不具备的风险因素，所以要谨慎行事。

If you do need to use C with Go, how to proceed depends on the Go compiler implementation. There are three Go compiler implementations supported by the Go team. These are `gc`, the default compiler, `gccgo`, which uses the GCC back end, and a somewhat less mature `gollvm`, which uses the LLVM infrastructure.

​	如果您确实需要在Go中使用C语言，如何进行取决于Go编译器的实现。Go 团队支持三种 Go 编译器实现。它们是 `gc`（默认的编译器）、`gccgo`（使用 GCC 后端）以及不太成熟的 `gollvm`（使用 LLVM 基础架构）。

`Gc` uses a different calling convention and linker from C and therefore cannot be called directly from C programs, or vice versa. The [`cgo`](https://go.dev/cmd/cgo/) program provides the mechanism for a "foreign function interface" to allow safe calling of C libraries from Go code. SWIG extends this capability to C++ libraries.

​	`Gc`使用与C不同的调用约定和链接器，因此不能直接从C程序中调用，反之亦然。[cgo](https://go.dev/cmd/cgo/)程序提供了一个 "外来函数接口 "的机制，允许从Go代码中安全地调用C库。SWIG将这种能力扩展到C++库。

You can also use `cgo` and SWIG with `Gccgo` and `gollvm`. Since they use a traditional API, it's also possible, with great care, to link code from these compilers directly with GCC/LLVM-compiled C or C++ programs. However, doing so safely requires an understanding of the calling conventions for all languages concerned, as well as concern for stack limits when calling C or C++ from Go.

​	您也可以在`Gccgo`和`gollvm`中使用`cgo`和SWIG。由于它们使用的是传统的API，所以在非常小心的情况下，也可以将这些编译器的代码直接与GCC/LLVM编译的C或C++程序链接。然而，要安全地做到这一点，需要了解所有相关语言的调用惯例，以及从 Go 中调用 C 或 C++ 时对堆栈限制的关注。

### What IDEs does Go support? Go 支持哪些集成开发环境？

The Go project does not include a custom IDE, but the language and libraries have been designed to make it easy to analyze source code. As a consequence, most well-known editors and IDEs support Go well, either directly or through a plugin.

​	Go 项目不包括定制的集成开发环境，但语言和库的设计使其易于分析源代码。因此，大多数知名的编辑器和集成开发环境都很好地支持 Go，可以直接使用或通过插件使用。

The list of well-known IDEs and editors that have good Go support available includes Emacs, Vim, VSCode, Atom, Eclipse, Sublime, IntelliJ (through a custom variant called Goland), and many more. Chances are your favorite environment is a productive one for programming in Go.

​	对Go有良好支持的知名IDE和编辑器包括Emacs、Vim、VSCode、Atom、Eclipse、Sublime、IntelliJ（通过一个名为Goland的自定义变体），以及更多。您最喜欢的环境有可能是用Go编程的有效环境。

### Does Go support Google's protocol buffers? Go 是否支持 Google 的协议缓冲区？

A separate open source project provides the necessary compiler plugin and library. It is available at [github.com/golang/protobuf/](https://github.com/golang/protobuf).

​	一个单独的开源项目提供了必要的编译器插件和库。它可以在 [github.com/golang/protobuf/](https://github.com/golang/protobuf)上找到。

### Can I translate the Go home page into another language?  我可以将 Go 主页翻译成其他语言吗？

Absolutely. We encourage developers to make Go Language sites in their own languages. However, if you choose to add the Google logo or branding to your site (it does not appear on [golang.org](https://go.dev/)), you will need to abide by the guidelines at [www.google.com/permissions/guidelines.html](https://www.google.com/permissions/guidelines.html)

​	当然可以。我们鼓励开发者用自己的语言制作 Go 语言网站。然而，如果您选择在您的网站上添加谷歌的标志或品牌（它不会出现在golang.org上），您需要遵守www.google.com/permissions/guidelines.html 的准则



## Design 设计

### Does Go have a runtime? Go有运行时吗？

Go does have an extensive library, called the *runtime*, that is part of every Go program. The runtime library implements garbage collection, concurrency, stack management, and other critical features of the Go language. Although it is more central to the language, Go's runtime is analogous to `libc`, the C library.

​	Go确实有一个广泛的库，称为 `runtime`，它是每个Go程序的一部分。`runtime`库实现了垃圾收集、并发、堆栈管理以及 Go 语言的其他关键功能。虽然它是语言的核心，但Go的`runtime`类似于`libc`，即C语言库。

It is important to understand, however, that Go's runtime does not include a virtual machine, such as is provided by the Java runtime. Go programs are compiled ahead of time to native machine code (or JavaScript or WebAssembly, for some variant implementations). Thus, although the term is often used to describe the virtual environment in which a program runs, in Go the word "runtime" is just the name given to the library providing critical language services.

​	然而，重要的是要理解Go的`runtime`不包括虚拟机，如Java`runtime`所提供的虚拟机。Go程序被提前编译为本地机器代码（或JavaScript或WebAssembly，用于某些变体实现）。因此，虽然这个词经常被用来描述程序运行的虚拟环境，但在 Go 中，"`runtime` "这个词`只是提供关键语言服务的库的名称`。

### What's up with Unicode identifiers? Unicode标识符是怎么回事？

When designing Go, we wanted to make sure that it was not overly ASCII-centric, which meant extending the space of identifiers from the confines of 7-bit ASCII. Go's rule—identifier characters must be letters or digits as defined by Unicode—is simple to understand and to implement but has restrictions. Combining characters are excluded by design, for instance, and that excludes some languages such as Devanagari.

​	在设计Go时，我们希望确保它不会过度以ASCII为中心，这意味着将标识符的空间从7位ASCII的范围内扩展出来。Go的规则——标识符必须是Unicode所定义的字母或数字——很容易理解和实现，`但也有限制`。例如，组合字符在设计上被排除在外，而这也排除了一些语言，如Devanagari。

This rule has one other unfortunate consequence. Since an exported identifier must begin with an upper-case letter, identifiers created from characters in some languages can, by definition, not be exported. For now the only solution is to use something like `X日本語`, which is clearly unsatisfactory.

​	这个规则还有一个不幸的后果。由于导出的标识符必须以大写字母开始，根据定义，由某些语言的字符创建的标识符不能被导出。目前，唯一的解决办法是使用类似`X日本語`的东西，这显然是不能令人满意的。

Since the earliest version of the language, there has been considerable thought into how best to expand the identifier space to accommodate programmers using other native languages. Exactly what to do remains an active topic of discussion, and a future version of the language may be more liberal in its definition of an identifier. For instance, it might adopt some of the ideas from the Unicode organization's [recommendations](http://unicode.org/reports/tr31/) for identifiers. Whatever happens, it must be done compatibly while preserving (or perhaps expanding) the way letter case determines visibility of identifiers, which remains one of our favorite features of Go.

​	自从该语言的最早版本以来，人们一直在考虑如何最好地扩展标识符空间以适应使用其他母语的程序员。具体怎么做仍然是一个活跃的讨论话题，未来版本的语言可能会在标识符的定义上更加自由。例如，它可能会采用Unicode组织关于标识符的[建议](http://unicode.org/reports/tr31/)中的一些想法。无论发生什么，都必须在保留（或扩大）字母大小写决定标识符的可见性的同时进行兼容，这仍然是我们最喜欢的Go的特点之一。

For the time being, we have a simple rule that can be expanded later without breaking programs, one that avoids bugs that would surely arise from a rule that admits ambiguous identifiers.

​	目前，我们有一个简单的规则，以后可以在不破坏程序的情况下进行扩展，这个规则可以避免因允许模棱两可的标识符的规则而产生的错误。

### Why does Go not have feature X? 为什么Go没有X的特性？

Every language contains novel features and omits someone's favorite feature. Go was designed with an eye on felicity of programming, speed of compilation, orthogonality of concepts, and the need to support features such as concurrency and garbage collection. Your favorite feature may be missing because it doesn't fit, because it affects compilation speed or clarity of design, or because it would make the fundamental system model too difficult.

​	每种语言都包含新奇的功能，并忽略了某些人最喜欢的功能。Go的设计着眼于编程的便利性、编译的速度、概念的正交性，以及支持并发和垃圾回收等功能的需要。您最喜欢的功能可能因为不合适而缺失，因为它影响了编译速度或设计的清晰度，或者因为它使基本的系统模型过于困难。

If it bothers you that Go is missing feature X, please forgive us and investigate the features that Go does have. You might find that they compensate in interesting ways for the lack of X.

​	如果Go缺失了X功能让您感到困扰，请原谅我们，转而研究Go确实拥有的功能。您可能会发现它们以有趣的方式弥补了X的缺失。

### When did Go get generic types? Go 什么时候有了泛型？

The Go 1.18 release added type parameters to the language. This permits a form of polymorphic or generic programming. See the [language spec](https://go.dev/ref/spec) and the [proposal](https://go.dev/design/43651-type-parameters) for details.

Go 1.18版本在语言中加入了类型参数。这允许一种多态或泛型编程的形式。详情请参见[语言规范](../References/LanguageSpecification/DeclarationsAndScope#type-definitions )和[提案](https://go.dev/design/43651-type-parameters)。

### Why was Go initially released without generic types? 为什么 Go 最初发布时没有泛型？

Go was intended as a language for writing server programs that would be easy to maintain over time. (See [this article](https://go.dev/talks/2012/splash.article) for more background.) The design concentrated on things like scalability, readability, and concurrency. Polymorphic programming did not seem essential to the language's goals at the time, and so was initially left out for simplicity.

​	Go 的目的是作为一种编写服务器程序的语言，以便于长期维护。(更多背景见[这篇文章](https://go.dev/talks/2012/splash.article)。)设计集中在可扩展性、可读性和并发性等方面。多态编程在当时看来对该语言的目标并不重要，所以最初为了简化而被排除在外。

Generics are convenient but they come at a cost in complexity in the type system and run-time. It took a while to develop a design that we believe gives value proportionate to the complexity.

​	泛型是很方便的，但它们的代价是类型系统和运行时间的复杂性。我们花了一些时间来开发一个我们认为能够提供与复杂性相称的价值的设计。

### Why does Go not have exceptions?  为什么Go没有异常？

We believe that coupling exceptions to a control structure, as in the `try-catch-finally` idiom, results in convoluted code. It also tends to encourage programmers to label too many ordinary errors, such as failing to open a file, as exceptional.

​	我们认为将异常与控制结构相耦合，如`try-catch-finally`习语，会造成代码的复杂化。它还倾向于鼓励程序员将太多的普通错误，如无法打开文件，标记为异常。

Go takes a different approach. For plain error handling, Go's multi-value returns make it easy to report an error without overloading the return value. [A canonical error type, coupled with Go's other features](https://go.dev/doc/articles/error_handling.html), makes error handling pleasant but quite different from that in other languages.

​	Go采取了一种不同的方法。对于普通的错误处理，Go的`多值返回`使得报告错误很容易，而不需要重载返回值。[一个典型的错误类型，加上Go的其他特性](https://go.dev/doc/articles/error_handling.html)，使错误处理变得令人愉快，但与其他语言中的错误处理完全不同。

Go also has a couple of built-in functions to signal and recover from truly exceptional conditions. The recovery mechanism is executed only as part of a function's state being torn down after an error, which is sufficient to handle catastrophe but requires no extra control structures and, when used well, can result in clean error-handling code.

​	Go也有几个内置的函数，用来发出信号并从真正的特殊情况下恢复。恢复机制仅作为错误发生后函数状态被拆解的一部分来执行，这足以处理灾难，但不需要额外的控制结构，如果使用得好，可以产生干净的错误处理代码。

See the [Defer, Panic, and Recover](https://go.dev/doc/articles/defer_panic_recover.html) article for details. Also, the [Errors are values](https://blog.golang.org/errors-are-values) blog post describes one approach to handling errors cleanly in Go by demonstrating that, since errors are just values, the full power of Go can be deployed in error handling.

​	详见[Defer, Panic, and Recover]({{< ref "/goBlog/2010/DeferPanicAandRecover" >}})一文。另外，[Errors are values]({{< ref "/goBlog/2015/ErrorsAreValues" >}})的博文描述了一种在 Go 中干净地处理错误的方法，它表明，由于错误只是值，因此可以在错误处理中部署 Go 的全部力量。

### Why does Go not have assertions? 为什么 Go 没有断言？

Go doesn't provide assertions. They are undeniably convenient, but our experience has been that programmers use them as a crutch to avoid thinking about proper error handling and reporting. Proper error handling means that servers continue to operate instead of crashing after a non-fatal error. Proper error reporting means that errors are direct and to the point, saving the programmer from interpreting a large crash trace. Precise errors are particularly important when the programmer seeing the errors is not familiar with the code.

​	Go 并没有提供断言。不可否认的是，它们很方便，但我们的经验是，程序员把它们当作拐杖，避免考虑正确的错误处理和报告。正确的错误处理意味着服务器可以继续运行，而不是在出现非致命错误后崩溃。正确的错误报告意味着错误是直接的，是有针对性的，使程序员不必解释一个大的崩溃跟踪。当看到错误的程序员不熟悉代码的时候，精确的错误就显得尤为重要。

We understand that this is a point of contention. There are many things in the Go language and libraries that differ from modern practices, simply because we feel it's sometimes worth trying a different approach.

​	我们理解这是一个争论点。Go语言和库中有许多与现代实践不同的东西，只是因为我们觉得有时值得尝试不同的方法。

### Why build concurrency on the ideas of CSP? 为什么要在CSP的思想上建立并发性？

Concurrency and multi-threaded programming have over time developed a reputation for difficulty. We believe this is due partly to complex designs such as [pthreads](https://en.wikipedia.org/wiki/POSIX_Threads) and partly to overemphasis on low-level details such as mutexes, condition variables, and memory barriers. Higher-level interfaces enable much simpler code, even if there are still mutexes and such under the covers.

​	随着时间的推移，并发和多线程编程以其难度而闻名。我们相信这部分是由于复杂的设计，比如[pthreads](https://en.wikipedia.org/wiki/POSIX_Threads)，部分是由于过度强调低级别的细节，比如mutexes，条件变量和内存障碍。更高级别的接口可以使代码简单得多，即使下面仍然有互斥等问题。

One of the most successful models for providing high-level linguistic support for concurrency comes from Hoare's Communicating Sequential Processes, or CSP. Occam and Erlang are two well known languages that stem from CSP. Go's concurrency primitives derive from a different part of the family tree whose main contribution is the powerful notion of channels as first class objects. Experience with several earlier languages has shown that the CSP model fits well into a procedural language framework.

​	为并发提供高级语言支持的最成功的模型之一来自Hoare的Communicating Sequential Processes，即CSP。Occam和Erlang是源于CSP的两种著名语言。Go的并发原语来自家族树的另一部分，其主要贡献是将通道作为第一类对象的强大概念。早期几种语言的经验表明，CSP模型很适合程序语言框架。

### Why goroutines instead of threads? 为什么是goroutines而不是线程？

Goroutines are part of making concurrency easy to use. The idea, which has been around for a while, is to multiplex independently executing functions—coroutines—onto a set of threads. When a coroutine blocks, such as by calling a blocking system call, the run-time automatically moves other coroutines on the same operating system thread to a different, runnable thread so they won't be blocked. The programmer sees none of this, which is the point. The result, which we call goroutines, can be very cheap: they have little overhead beyond the memory for the stack, which is just a few kilobytes.

​	goroutines是使并发性易于使用的一部分。这个想法已经存在了一段时间，它是将独立执行的函数——coroutines——复用到一组线程中。当一个 coroutine 阻塞时，例如通过调用一个阻塞的系统调用，run-time 会自动将同一操作系统线程上的其他程序转移到不同的可运行线程，这样它们就不会被阻塞。程序员看不到这些，这就是问题所在。这样的结果，我们称之为goroutines，可以非常便宜：除了堆栈的内存，它们几乎没有开销，而堆栈的内存只有几千字节。

To make the stacks small, Go's run-time uses resizable, bounded stacks. A newly minted goroutine is given a few kilobytes, which is almost always enough. When it isn't, the run-time grows (and shrinks) the memory for storing the stack automatically, allowing many goroutines to live in a modest amount of memory. The CPU overhead averages about three cheap instructions per function call. It is practical to create hundreds of thousands of goroutines in the same address space. If goroutines were just threads, system resources would run out at a much smaller number.

​	为了使堆栈变小，Go的run-time使用可调整大小的有界堆栈。一个新诞生的goroutine被赋予几千字节的内存，这几乎总是足够的。当它不够时，运行时间会自动增加（和缩小）用于存储堆栈的内存，从而使许多goroutine能够在适量的内存中生存。每个函数调用的CPU开销平均约为三条廉价指令。在同一地址空间中创建成百上千个goroutines是很实际的。如果goroutines只是线程，系统资源会在更少的数量上耗尽。

### Why are map operations not defined to be atomic? 为什么映射操作没有被定义为原子操作？

After long discussion it was decided that the typical use of maps did not require safe access from multiple goroutines, and in those cases where it did, the map was probably part of some larger data structure or computation that was already synchronized. Therefore requiring that all map operations grab a mutex would slow down most programs and add safety to few. This was not an easy decision, however, since it means uncontrolled map access can crash the program.

​	经过长时间的讨论，我们决定，映射的典型使用并不要求从多个goroutine中进行安全访问，而在那些需要访问的情况下，映射可能是一些更大的数据结构或计算的一部分，而这些数据结构或计算已经被同步了。因此，要求所有的映射操作都要抓取一个mutex，这将降低大多数程序的速度，并为少数程序增加安全性。然而，这并不是一个简单的决定，因为这意味着不受控制的映射访问会使程序崩溃。

The language does not preclude atomic map updates. When required, such as when hosting an untrusted program, the implementation could interlock map access.

​	该语言并不排除原子映射更新。在需要的时候，比如在托管一个不受信任的程序时，实现可以对映射访问进行联锁。

Map access is unsafe only when updates are occurring. As long as all goroutines are only reading—looking up elements in the map, including iterating through it using a `for` `range` loop—and not changing the map by assigning to elements or doing deletions, it is safe for them to access the map concurrently without synchronization.

​	只有在更新发生时，映射访问才是不安全的。只要所有的goroutines只是阅读查找映射中的元素，包括使用for range循环遍历它——而不是通过向元素赋值或进行删除来改变映射，那么它们在没有同步的情况下同时访问映射是安全的。

As an aid to correct map use, some implementations of the language contain a special check that automatically reports at run time when a map is modified unsafely by concurrent execution.

​	作为对正确使用映射的帮助，一些语言的实现包含一个特殊的检查，当映射被并发执行不安全地修改时，会在运行时自动报告。

### Will you accept my language change? 您会接受我的语言修改吗？

People often suggest improvements to the language—the [mailing list](https://groups.google.com/group/golang-nuts) contains a rich history of such discussions—but very few of these changes have been accepted.

​	人们经常建议对语言进行改进——[邮件列表](https://groups.google.com/group/golang-nuts)中包含了丰富的此类讨论历史——但这些修改很少被接受。

Although Go is an open source project, the language and libraries are protected by a [compatibility promise](https://go.dev/doc/go1compat.html) that prevents changes that break existing programs, at least at the source code level (programs may need to be recompiled occasionally to stay current). If your proposal violates the Go 1 specification we cannot even entertain the idea, regardless of its merit. A future major release of Go may be incompatible with Go 1, but discussions on that topic have only just begun and one thing is certain: there will be very few such incompatibilities introduced in the process. Moreover, the compatibility promise encourages us to provide an automatic path forward for old programs to adapt should that situation arise.

​	虽然 Go 是一个开源项目，但语言和库受到[兼容性承诺](https://go.dev/doc/go1compat.html)的保护，至少在源代码层面上，不会出现破坏现有程序的变化（程序可能需要偶尔重新编译以保持最新）。如果您的建议违反了Go 1的规范，我们甚至不能接受这个想法，不管它的优点是什么。未来的Go的主要版本可能会与Go 1不兼容，但关于这个话题的讨论才刚刚开始，有一点是肯定的：在这个过程中会很少引入这样的不兼容问题。此外，兼容性的承诺鼓励我们为老程序提供一条自动改进的道路，以便在出现这种情况时进行调整。

Even if your proposal is compatible with the Go 1 spec, it might not be in the spirit of Go's design goals. The article *[Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)* explains Go's origins and the motivation behind its design.

​	即使您的方案与Go 1规范兼容，它也可能不符合Go的设计目标的精神。文章[Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)一文解释了 Go 的起源及其设计背后的动机。

## Types 类型

### Is Go an object-oriented language? Go 是一种面向对象的语言吗？

Yes and no. Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy. The concept of "interface" in Go provides a different approach that we believe is easy to use and in some ways more general. There are also ways to embed types in other types to provide something analogous—but not identical—to subclassing. Moreover, methods in Go are more general than in C++ or Java: they can be defined for any sort of data, even built-in types such as plain, "unboxed" integers. They are not restricted to structs (classes).

​	是，也不是。虽然 Go 有类型和方法，并且允许面向对象的编程风格，但没有类型层次。Go 中的 "interface "概念提供了一种不同的方法，我们认为这种方法易于使用，而且在某些方面更加通用。还有一些方法可以将类型嵌入到其他类型中，以提供类似的东西，但不等同于子类。此外，Go中的方法比C++或Java中的方法更通用：它们可以为任何类型的数据定义，甚至是内置类型，如普通的、"未装箱的 "整数。它们并不局限于结构体（类）。

Also, the lack of a type hierarchy makes "objects" in Go feel much more lightweight than in languages such as C++ or Java.

​	另外，由于缺乏类型层次，Go中的 "对象 "比C++或Java等语言更轻巧。

### How do I get dynamic dispatch of methods? 我如何获得方法的动态调度？

The only way to have dynamically dispatched methods is through an interface. Methods on a struct or any other concrete type are always resolved statically.

​	拥有动态分配方法的唯一方法是通过接口。结构体或任何其他具体类型的方法总是以静态方式解决。

### Why is there no type inheritance? 为什么没有类型继承？

Object-oriented programming, at least in the best-known languages, involves too much discussion of the relationships between types, relationships that often could be derived automatically. Go takes a different approach.

​	面向对象的编程，至少在最著名的语言中，涉及到太多关于类型之间关系的讨论，这些关系往往可以自动衍生。Go采取了一种不同的方法。

Rather than requiring the programmer to declare ahead of time that two types are related, in Go a type automatically satisfies any interface that specifies a subset of its methods. Besides reducing the bookkeeping, this approach has real advantages. Types can satisfy many interfaces at once, without the complexities of traditional multiple inheritance. Interfaces can be very lightweight—an interface with one or even zero methods can express a useful concept. Interfaces can be added after the fact if a new idea comes along or for testing—without annotating the original types. Because there are no explicit relationships between types and interfaces, there is no type hierarchy to manage or discuss.

​	在Go中，一个类型不是要求程序员提前声明两个类型的关系，而是自动满足任何指定其方法子集的接口。除了减少记账，这种方法还有真正的优势。类型可以同时满足许多接口，而没有传统的多重继承的复杂问题。接口可以是非常轻量级的——一个只有一个甚至是零个方法的接口可以表达一个有用的概念。如果有新的想法出现，或者为了测试，可以在事后添加接口，而不需要对原始类型进行注释。因为类型和接口之间没有明确的关系，所以不存在需要管理或讨论的类型层次结构。

It's possible to use these ideas to construct something analogous to type-safe Unix pipes. For instance, see how `fmt.Fprintf` enables formatted printing to any output, not just a file, or how the `bufio` package can be completely separate from file I/O, or how the `image` packages generate compressed image files. All these ideas stem from a single interface (`io.Writer`) representing a single method (`Write`). And that's only scratching the surface. Go's interfaces have a profound influence on how programs are structured.

​	我们可以用这些想法来构建类似于类型安全的 Unix 管道的东西。例如，看看`fmt.Fprintf`是如何实现对任何输出的格式化打印，而不仅仅是文件，或者`bufio`包是如何与文件I/O完全分离的，或者`image`包是如何生成压缩图像文件的。所有这些想法都源于一个接口（`io.Writer`），表示一个方法（`Write`）。而这仅仅是表面现象。Go的接口对程序的结构有深刻的影响。

It takes some getting used to but this implicit style of type dependency is one of the most productive things about Go.

​	这需要一些时间来适应，但这种隐式的类型依赖风格是Go最富有成效的地方之一。

### Why is `len` a function and not a method? 为什么`len`是一个函数而不是一个方法？

We debated this issue but decided implementing `len` and friends as functions was fine in practice and didn't complicate questions about the interface (in the Go type sense) of basic types.

​	我们对这个问题进行了辩论，但最终决定将 len 和朋友们作为函数来实现，在实践中是没有问题的，也不会使基本类型的接口（在 Go 类型意义上）问题复杂化。

### Why does Go not support overloading of methods and operators? 为什么 Go 不支持方法和运算符的重载？

Method dispatch is simplified if it doesn't need to do type matching as well. Experience with other languages told us that having a variety of methods with the same name but different signatures was occasionally useful but that it could also be confusing and fragile in practice. Matching only by name and requiring consistency in the types was a major simplifying decision in Go's type system.

​	如果不需要同时进行类型匹配，方法调度就会被简化。其他语言的经验告诉我们，拥有各种同名但不同签名的方法偶尔是有用的，但在实践中也会让人感到困惑和脆弱。在Go的类型系统中，仅通过名称进行匹配并要求类型的一致性是一个重要的简化决定。

Regarding operator overloading, it seems more a convenience than an absolute requirement. Again, things are simpler without it.

​	关于操作符重载，它似乎更像是一种便利，而不是绝对的要求。同样，没有它，事情会更简单。

### Why doesn't Go have "implements" declarations? 为什么 Go 没有 "implements"声明？

A Go type satisfies an interface by implementing the methods of that interface, nothing more. This property allows interfaces to be defined and used without needing to modify existing code. It enables a kind of [structural typing](https://en.wikipedia.org/wiki/Structural_type_system) that promotes separation of concerns and improves code re-use, and makes it easier to build on patterns that emerge as the code develops. The semantics of interfaces is one of the main reasons for Go's nimble, lightweight feel.

​	Go类型通过实现一个接口的方法来满足该接口，仅此而已。这个属性允许定义和使用接口，而不需要修改现有的代码。它实现了一种结构化的类型，促进了关注点的分离，提高了代码的重复使用，并使之更容易建立在代码发展过程中出现的模式上。接口的语义是Go的灵活、轻量级感觉的主要原因之一。

See the [question on type inheritance](https://go.dev/doc/faq#inheritance) for more detail.

​	更多细节请参见关于[类型继承的问题](#why-is-there-no-type-inheritance)。

### How can I guarantee my type satisfies an interface? 如何保证我的类型满足接口的要求？

You can ask the compiler to check that the type `T` implements the interface `I` by attempting an assignment using the zero value for `T` or pointer to `T`, as appropriate:

​	您可以要求编译器通过尝试使用 `T` 的零值或 `T` 的指针进行赋值，来检查类型 `T` 是否实现了接口 `I`，如果合适的话：

```go linenums="1"
type T struct{}
var _ I = T{}       // Verify that T implements I. => 验证 T 实现了 I
var _ I = (*T)(nil) // Verify that *T implements I. => 验证 *T 实现了 I
```

If `T` (or `*T`, accordingly) doesn't implement `I`, the mistake will be caught at compile time.

​	如果`T`（或`*T`，相应地）没有实现`I`，这个错误将在编译时被捕获。

If you wish the users of an interface to explicitly declare that they implement it, you can add a method with a descriptive name to the interface's method set. For example:

​	如果您希望一个接口的用户明确声明他们实现了这个接口，您可以在接口的方法集中添加一个带有描述性名称的方法。比如说：

```go linenums="1"
type Fooer interface {
    Foo()
    ImplementsFooer()
}
```

A type must then implement the `ImplementsFooer` method to be a `Fooer`, clearly documenting the fact and announcing it in [go doc](https://go.dev/cmd/go/#hdr-Show_documentation_for_package_or_symbol)'s output.

​	然后一个类型必须实现`ImplementsFooer`方法才能成为`Fooer`，清楚地记录这一事实并在[go doc](https://go.dev/cmd/go/#hdr-Show_documentation_for_package_or_symbol)的输出中公布。

```go linenums="1"
type Bar struct{}
func (b Bar) ImplementsFooer() {}
func (b Bar) Foo() {}
```

Most code doesn't make use of such constraints, since they limit the utility of the interface idea. Sometimes, though, they're necessary to resolve ambiguities among similar interfaces.

​	大多数代码都不使用这种约束，因为它们限制了接口思想的实用性。但有时，它们对于解决类似接口之间的歧义是必要的。

### Why doesn't type T satisfy the Equal interface? 为什么类型T不满足 Equal 接口？

Consider this simple interface to represent an object that can compare itself with another value:

​	考虑用这个简单的接口来表示一个可以将自己与另一个值进行比较的对象：

```go linenums="1"
type Equaler interface {
    Equal(Equaler) bool
}
```

和这个类型，`T`：

```go linenums="1"
type T int
func (t T) Equal(u T) bool { return t == u } // does not satisfy Equaler => 不满足 Equaler
```

Unlike the analogous situation in some polymorphic type systems, `T` does not implement `Equaler`. The argument type of `T.Equal` is `T`, not literally the required type `Equaler`.

​	与某些多态类型系统中的类似情况不同，`T`并没有实现`Equaler`。`T.Equal`的实参类型是`T`，而不是字面上所要求的`Equaler`类型。

In Go, the type system does not promote the argument of `Equal`; that is the programmer's responsibility, as illustrated by the type `T2`, which does implement `Equaler`:

​	在Go中，类型系统不提升`Equal`的实参；这是程序员的责任，如类型`T2`所示，它确实实现了`Equaler`：

```go linenums="1"
type T2 int
func (t T2) Equal(u Equaler) bool { return t == u.(T2) }  // satisfies Equaler
```

Even this isn't like other type systems, though, because in Go *any* type that satisfies `Equaler` could be passed as the argument to `T2.Equal`, and at run time we must check that the argument is of type `T2`. Some languages arrange to make that guarantee at compile time.

不过，即使这样也和其他类型系统不同，因为在Go中，任何满足`Equaler`的类型都可以作为实参传给T2.Equal，在运行时我们必须检查参数是否为`T2`类型。有些语言在编译时就安排了这种保证。

A related example goes the other way:

​	一个相关的例子则恰恰相反：

```go linenums="1"
type Opener interface {
   Open() Reader
}

func (t T3) Open() *os.File
```

In Go, `T3` does not satisfy `Opener`, although it might in another language.

在Go中，`T3`并不满足`Opener`，尽管它在其他语言中可能满足。

While it is true that Go's type system does less for the programmer in such cases, the lack of subtyping makes the rules about interface satisfaction very easy to state: are the function's names and signatures exactly those of the interface? Go's rule is also easy to implement efficiently. We feel these benefits offset the lack of automatic type promotion. Should Go one day adopt some form of polymorphic typing, we expect there would be a way to express the idea of these examples and also have them be statically checked.

​	虽然在这种情况下，Go的类型系统确实对程序员的帮助较小，但由于缺乏子类型，关于接口满足的规则非常容易说明：`函数的名称和签名是否正是接口的名称和签名？`Go的规则也很容易有效实现。我们觉得这些好处抵消了自动类型推导的不足。如果有一天Go采用了某种形式的多态类型，我们希望有一种方法来表达这些例子的想法，并让它们得到静态检查。

### Can I convert a []T to an []interface{}? 我可以将一个[]T转换为一个[]interface{}吗？

Not directly. It is disallowed by the language specification because the two types do not have the same representation in memory. It is necessary to copy the elements individually to the destination slice. This example converts a slice of `int` to a slice of `interface{}`:

​	不能直接转换。语言规范不允许这样做，因为这两种类型在内存中没有相同的表示。有必要将元素单独复制到目标切片中。这个例子将一个`int`的切片转换为`interface{}`的切片：

```go linenums="1"
t := []int{1, 2, 3, 4}
s := make([]interface{}, len(t))
for i, v := range t {
    s[i] = v
}
```

### Can I convert []T1 to []T2 if T1 and T2 have the same underlying type? 如果T1和T2有相同的底层类型，我可以将[]T1转换成[]T2吗？

This last line of this code sample does not compile.

​	这段代码样本的最后一行不能编译。

```go linenums="1"
type T1 int
type T2 int
var t1 T1
var x = T2(t1) // OK
var st1 []T1
var sx = ([]T2)(st1) // NOT OK
```

In Go, types are closely tied to methods, in that every named type has a (possibly empty) method set. The general rule is that you can change the name of the type being converted (and thus possibly change its method set) but you can't change the name (and method set) of elements of a composite type. Go requires you to be explicit about type conversions.

​	`在Go中，类型与方法紧密相连，每个命名的类型都有一个（可能是空的）方法集。`一般的规则是，您可以改变被转换的类型的名称（从而可能改变其方法集），`但您不能改变复合类型的元素的名称（和方法集）`。Go要求您对类型转换进行明确说明。

### Why is my nil error value not equal to nil? 为什么我的nil错误值不等于nil？

Under the covers, interfaces are implemented as two elements, a type `T` and a value `V`. `V` is a concrete value such as an `int`, `struct` or pointer, never an interface itself, and has type `T`. For instance, if we store the `int` value 3 in an interface, the resulting interface value has, schematically, (`T=int`, `V=3`). The value `V` is also known as the interface's *dynamic* value, since a given interface variable might hold different values `V` (and corresponding types `T`) during the execution of the program.

​	在底层，接口被实现为两个元素，一个类型`T`和一个值`V`。`V`是一个具体的值，如`int`、`struct`或指针，绝不是接口本身，并且具有类型`T`。例如，如果我们在一个接口中存储`int`值3，那么产生的接口值就有，示意为（`T=int`，V=3）。值`V`也被称为接口的动态值，因为在程序执行过程中，一个给定的接口变量可能持有不同的值`V`（以及相应的类型`T`）。

An interface value is `nil` only if the `V` and `T` are both unset, (`T=nil`, `V` is not set), In particular, a `nil` interface will always hold a `nil` type. If we store a `nil` pointer of type `*int` inside an interface value, the inner type will be `*int` regardless of the value of the pointer: (`T=*int`, `V=nil`). Such an interface value will therefore be non-`nil` *even when the pointer value `V` inside is* `nil`.

​	只有当`V`和`T`都未设置时，接口值才是`nil`，（`T=nil`，`V未设置`），特别是，一个`nil`接口将永远持有一个`nil`类型。如果我们在一个接口值里面存储一个`*int`类型的`nil`指针，那么不管指针的值是什么，内部类型都是`*int`：（`T=*int`，`V=nil`）。因此，即使里面的指针值`V`是`nil`，这样的接口值也是非`nil`的。

This situation can be confusing, and arises when a `nil` value is stored inside an interface value such as an `error` return:

​	这种情况可能会引起混淆，当一个`nil`值被存储在一个接口值里面时就会出现，比如一个`error`返回：

```go linenums="1"
func returnsError() error {
    var p *MyError = nil
    if bad() {
        p = ErrBad
    }
    return p // Will always return a non-nil error. =>  将始终返回非 nil error
}
```

If all goes well, the function returns a `nil` `p`, so the return value is an `error` interface value holding (`T=*MyError`, `V=nil`). This means that if the caller compares the returned error to `nil`, it will always look as if there was an error even if nothing bad happened. To return a proper `nil` `error` to the caller, the function must return an explicit `nil`:

​	如果一切顺利，该函数返回一个`nil`  `p`，所以返回值是一个`error`接口值持有（`T=*MyError`, `V=nil`）。这意味着，如果调用者将返回的错误与`nil`进行比较，即使没有发生什么坏事，也会一直看起来像有一个错误。为了向调用者返回一个适当的`nil`错误，该函数必须返回一个明确的`nil`：

```go linenums="1"
func returnsError() error {
    if bad() {
        return ErrBad
    }
    return nil
}
```

It's a good idea for functions that return errors always to use the `error` type in their signature (as we did above) rather than a concrete type such as `*MyError`, to help guarantee the error is created correctly. As an example, [`os.Open`](https://go.dev/pkg/os/#Open) returns an `error` even though, if not `nil`, it's always of concrete type [`*os.PathError`](https://go.dev/pkg/os/#PathError).

​	对于返回错误的函数来说，最好在其签名中使用`error`类型（就像我们上面做的那样），而不是一个具体的类型，如`*MyError`，以帮助保证错误被正确创建。作为一个例子，[os.Open](https://go.dev/pkg/os/#Open)返回一个`error`，尽管如果不是`nil`，它总是具体类型`*os.PathError`。

Similar situations to those described here can arise whenever interfaces are used. Just keep in mind that if any concrete value has been stored in the interface, the interface will not be `nil`. For more information, see [The Laws of Reflection](https://go.dev/doc/articles/laws_of_reflection.html).

​	只要使用接口，就会出现与这里描述的类似情况。只要记住，`如果有任何具体的值被存储在接口中，接口就`不会是`nil`。更多信息，请参见[反射的法则]({{< ref "/goBlog/2011/TheLawsOfReflection" >}})。

### Why are there no untagged unions, as in C? 为什么没有像C语言那样的无标记的联合体？

Untagged unions would violate Go's memory safety guarantees.

​	没有标记的联合体会违反 Go 的内存安全保证。

### Why does Go not have variant types? 为什么 Go 没有变体类型？

Variant types, also known as algebraic types, provide a way to specify that a value might take one of a set of other types, but only those types. A common example in systems programming would specify that an error is, say, a network error, a security error or an application error and allow the caller to discriminate the source of the problem by examining the type of the error. Another example is a syntax tree in which each node can be a different type: declaration, statement, assignment and so on.

​	变体类型，也被称为代数类型，提供了一种方法来指定一个值可能采取一系列其他类型中的一个，但只采取这些类型。在系统编程中，一个常见的例子是指定一个错误是，例如，一个网络错误、一个安全错误或一个应用程序错误，并允许调用者通过检查错误的类型来分辨问题的来源。另一个例子是语法树，其中每个节点可以是不同的类型：声明、语句、赋值等等。

We considered adding variant types to Go, but after discussion decided to leave them out because they overlap in confusing ways with interfaces. What would happen if the elements of a variant type were themselves interfaces?

​	我们考虑过在Go中加入变体类型，但经过讨论后决定不加入，因为它们与接口的重叠方式令人困惑。如果变体类型的元素本身就是接口，会发生什么？

Also, some of what variant types address is already covered by the language. The error example is easy to express using an interface value to hold the error and a type switch to discriminate cases. The syntax tree example is also doable, although not as elegantly.

​	另外，变体类型所涉及的一些内容已经被语言所涵盖。错误的例子很容易表达，用一个接口值来保持错误，用一个类型转换来区分情况。语法树的例子也是可以做到的，虽然没有那么优雅。

### Why does Go not have covariant result types? 为什么Go没有协变结果类型？

Covariant result types would mean that an interface like

​	协变结果类型意味着一个接口，如

```go linenums="1"
type Copyable interface {
    Copy() interface{}
}
```

would be satisfied by the method

的方法会被满足

```go linenums="1"
func (v Value) Copy() Value
```

because `Value` implements the empty interface. In Go method types must match exactly, so `Value` does not implement `Copyable`. Go separates the notion of what a type does—its methods—from the type's implementation. If two methods return different types, they are not doing the same thing. Programmers who want covariant result types are often trying to express a type hierarchy through interfaces. In Go it's more natural to have a clean separation between interface and implementation.

​	因为`Value`实现了这个空接口。`在Go中，方法类型必须完全匹配，所以Value并没有实现Copyable`。Go将一个类型所做的事情的概念——它的方法——与类型的实现分开。`如果两个方法返回不同的类型，它们就不是在做同一件事。`希望得到协变结果类型的程序员往往试图通过接口来表达类型的层次结构。在Go中，接口和实现之间的分离是比较自然的。

## Values 值

### Why does Go not provide implicit numeric conversions? 为什么 Go 不提供隐式数值转换？

The convenience of automatic conversion between numeric types in C is outweighed by the confusion it causes. When is an expression unsigned? How big is the value? Does it overflow? Is the result portable, independent of the machine on which it executes? It also complicates the compiler; "the usual arithmetic conversions" are not easy to implement and inconsistent across architectures. For reasons of portability, we decided to make things clear and straightforward at the cost of some explicit conversions in the code. The definition of constants in Go—arbitrary precision values free of signedness and size annotations—ameliorates matters considerably, though.

​	在C语言中，数字类型之间的自动转换所带来的便利被其造成的混乱所抵消。一个表达式什么时候是无符号的？数值有多大？它是否会溢出？结果是否可移植，与执行它的机器无关？这也使编译器变得复杂；"通常的算术转换 "不容易实现，而且在不同的架构上也不一致。出于可移植性的考虑，我们决定以代码中一些明确的转换为代价，使事情变得清晰明了。在Go中对常量的定义——任意精度的值，没有符号和大小注释——大大改善了问题。

A related detail is that, unlike in C, `int` and `int64` are distinct types even if `int` is a 64-bit type. The `int` type is generic; if you care about how many bits an integer holds, Go encourages you to be explicit.

​	一个相关的细节是，与C不同，`int`和`int64`是不同的类型，即使`int`是一个64位的类型。`int` 类型是通用的；如果您关心一个整数拥有多少位，Go 鼓励您明确地表达出来。

### How do constants work in Go? 常量在 Go 中是如何工作的？

Although Go is strict about conversion between variables of different numeric types, constants in the language are much more flexible. Literal constants such as `23`, `3.14159` and [`math.Pi`](https://go.dev/pkg/math/#pkg-constants) occupy a sort of ideal number space, with arbitrary precision and no overflow or underflow. For instance, the value of `math.Pi` is specified to 63 places in the source code, and constant expressions involving the value keep precision beyond what a `float64` could hold. Only when the constant or constant expression is assigned to a variable—a memory location in the program—does it become a "computer" number with the usual floating-point properties and precision.

​	虽然 Go 对不同数字类型的变量之间的转换有严格的规定，但语言中的常量要灵活得多。字面常量如`23`、`3.14159`和[math.Pi](https://go.dev/pkg/math/#pkg-constants)占据了一种理想的数字空间，有任意的精度，没有溢出或下溢。例如，`math.Pi`的值在源代码中被指定有63处，涉及该值的常量表达式所保持的精度超过了`float64`所能容纳的。只有当常量或常量表达式被分配到一个变量——程序中的内存位置时，它才成为一个具有通常浮点属性和精度的 "computer"数字。

Also, because they are just numbers, not typed values, constants in Go can be used more freely than variables, thereby softening some of the awkwardness around the strict conversion rules. One can write expressions such as

​	另外，`由于常量只是数字，而不是类型化的数值，Go中的常量可以比变量更自由地使用`，从而缓解了严格的转换规则带来的一些尴尬局面。我们可以写出这样的表达式

```go linenums="1"
sqrt2 := math.Sqrt(2)
```

without complaint from the compiler because the ideal number `2` can be converted safely and accurately to a `float64` for the call to `math.Sqrt`.

这样的表达式，而不会被编译器抱怨，因为理想数字`2`可以被安全、准确地转换为`float64`来调用`math.Sqrt`。

A blog post titled [Constants](https://blog.golang.org/constants) explores this topic in more detail.

​	一篇题为 "[Constants]({{< ref "/goBlog/2014/Constants" >}}) "的博文更详细地探讨了这个话题。

### Why are maps built in? 为什么内建映射？

The same reason strings are: they are such a powerful and important data structure that providing one excellent implementation with syntactic support makes programming more pleasant. We believe that Go's implementation of maps is strong enough that it will serve for the vast majority of uses. If a specific application can benefit from a custom implementation, it's possible to write one but it will not be as convenient syntactically; this seems a reasonable tradeoff.

​	和字符串的原因一样：它们是如此强大和重要的数据结构，提供一个具有语法支持的优秀实现可以使编程更加愉快。我们相信Go的映射实现足够强大，可以满足绝大多数的使用。如果一个特定的应用可以从自定义的实现中受益，那么就可以写一个，但在语法上就不那么方便了；这似乎是一个合理的权衡。

### Why don't maps allow slices as keys? 为什么映射不允许将切片作为键？

Map lookup requires an equality operator, which slices do not implement. They don't implement equality because equality is not well defined on such types; there are multiple considerations involving shallow vs. deep comparison, pointer vs. value comparison, how to deal with recursive types, and so on. We may revisit this issue—and implementing equality for slices will not invalidate any existing programs—but without a clear idea of what equality of slices should mean, it was simpler to leave it out for now.

​	映射查询需要一个相等运算符，而切片并没有实现这个运算符。他们没有实现相等性，因为相等性在这种类型上没有得到很好的定义；有多种考虑，涉及到浅层与深层比较、指针与值比较、如何处理递归类型等等。我们可能会重新审视这个问题——为切片实现相等并不会使任何现有的程序失效，但是在没有明确切片相等的含义的情况下，暂时不考虑这个问题比较简单。

In Go 1, unlike prior releases, equality is defined for structs and arrays, so such types can be used as map keys. Slices still do not have a definition of equality, though.

​	在Go 1中，与之前的版本不同，为结构体和数组定义了相等性，所以这类类型可以作为映射键使用。不过，切片仍然没有相等性的定义。

### Why are maps, slices, and channels references while arrays are values? 为什么映射、切片和通道是引用而数组是值？

There's a lot of history on that topic. Early on, maps and channels were syntactically pointers and it was impossible to declare or use a non-pointer instance. Also, we struggled with how arrays should work. Eventually we decided that the strict separation of pointers and values made the language harder to use. Changing these types to act as references to the associated, shared data structures resolved these issues. This change added some regrettable complexity to the language but had a large effect on usability: Go became a more productive, comfortable language when it was introduced.

​	关于这个话题有很多历史。早期，map和channel在语法上是指针，不可能声明或使用一个非指针实例。此外，我们还为数组应该如何工作而挣扎。最终我们决定，指针和值的严格分离将会使语言更难使用。将这些类型改变为对相关共享数据结构的引用，解决了这些问题。这一改变给语言增加了一些令人遗憾的复杂性，但对可用性产生了很大的影响。Go在推出后成为了一种更有生产力、更舒适的语言。

## Writing Code 编写代码

### How are libraries documented? 库是如何被文档化的？

There is a program, `godoc`, written in Go, that extracts package documentation from the source code and serves it as a web page with links to declarations, files, and so on. An instance is running at [golang.org/pkg/](https://go.dev/pkg/). In fact, `godoc` implements the full site at [golang.org/](https://go.dev/).

​	有一个用Go编写的程序`godoc`，可以从源代码中提取包的文档，并将其作为网页提供给声明、文件等的链接。一个实例正在[go.dev/pkg/](https://go.dev/pkg/)上运行。事实上，`godoc`实现了[golang.org/](https://go.dev/)的完整网站。

A `godoc` instance may be configured to provide rich, interactive static analyses of symbols in the programs it displays; details are listed [here](https://go.dev/lib/godoc/analysis/help.html).

​	一个`godoc`实例可以被配置为对它所显示的程序中的符号提供丰富的、交互式的静态分析；细节在[这里](../godoc/analysis/help)列出。

For access to documentation from the command line, the [go](https://go.dev/pkg/cmd/go/) tool has a [doc](https://go.dev/pkg/cmd/go/#hdr-Show_documentation_for_package_or_symbol) subcommand that provides a textual interface to the same information.

​	为了从命令行访问文档，[go](https://go.dev/pkg/cmd/go/)工具有一个[doc](https://go.dev/pkg/cmd/go/#hdr-Show_documentation_for_package_or_symbol)子命令，为相同的信息提供一个文本接口。

### Is there a Go programming style guide? 是否有Go编程风格指南？

There is no explicit style guide, although there is certainly a recognizable "Go style".

​	虽然没有明确的风格指南，但肯定有一个可识别的 "Go 风格"。

Go has established conventions to guide decisions around naming, layout, and file organization. The document [Effective Go](https://go.dev/doc/effective_go.html) contains some advice on these topics. More directly, the program `gofmt` is a pretty-printer whose purpose is to enforce layout rules; it replaces the usual compendium of dos and don'ts that allows interpretation. All the Go code in the repository, and the vast majority in the open source world, has been run through `gofmt`.

​	Go 已经建立了惯例来指导命名、布局和文件组织方面的决策。文件 [Effective Go](../UsingAndUnderstandingGo/EffectiveGo) 包含了一些关于这些主题的建议。更直接地说，程序`gofmt`是一个漂亮的打印机，其目的是强制布局规则；它取代了允许解释的通常的注意事项的概要。存储库中的所有Go代码，以及开源世界中的绝大部分，都已经通过`gofmt`运行。

The document titled [Go Code Review Comments](https://go.dev/s/comments) is a collection of very short essays about details of Go idiom that are often missed by programmers. It is a handy reference for people doing code reviews for Go projects.

​	题为[Go Code Review Comments](../CodeReviewComments)的文件是一个非常简短的文章集，涉及到程序员经常忽略的Go习惯用语的细节。对于为Go项目做代码审查的人来说，它是一个方便的参考。

### How do I submit patches to the Go libraries? 我如何提交Go库的补丁？

The library sources are in the `src` directory of the repository. If you want to make a significant change, please discuss on the mailing list before embarking.

​	库的源代码在版本库的 `src` 目录中。如果您想做一个重大的改变，请在开始之前在邮件列表中讨论。

See the document [Contributing to the Go project](https://go.dev/doc/contribute.html) for more information about how to proceed.

​	有关如何进行的更多信息，请参见文档[Contributing to the Go project](../References/ContributionGuide)。

### Why does "go get" use HTTPS when cloning a repository? 为什么 "go get" 在克隆版本库时使用 HTTPS？

Companies often permit outgoing traffic only on the standard TCP ports 80 (HTTP) and 443 (HTTPS), blocking outgoing traffic on other ports, including TCP port 9418 (git) and TCP port 22 (SSH). When using HTTPS instead of HTTP, `git` enforces certificate validation by default, providing protection against man-in-the-middle, eavesdropping and tampering attacks. The `go get` command therefore uses HTTPS for safety.

​	公司通常只允许标准的TCP端口80（HTTP）和443（HTTPS）的出站流量，阻止其他端口的出站流量，包括TCP端口9418（git）和TCP端口22（SSH）。当使用HTTPS而不是HTTP时，`git`默认执行证书验证，提供保护以防止中间人、窃听和篡改攻击。因此，`go get`命令使用HTTPS来保证安全。

`Git` can be configured to authenticate over HTTPS or to use SSH in place of HTTPS. To authenticate over HTTPS, you can add a line to the `$HOME/.netrc` file that git consults:

​	`Git` 可以被配置为通过 HTTPS 进行认证，或者使用 SSH 来代替 HTTPS。要通过HTTPS认证，可以在git查阅的`$HOME/.netrc`文件中添加一行。

```
machine github.com login USERNAME password APIKEY
```

For GitHub accounts, the password can be a [personal access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/).

​	对于GitHub账户，密码可以是[个人访问令牌](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/)。

`Git` can also be configured to use SSH in place of HTTPS for URLs matching a given prefix. For example, to use SSH for all GitHub access, add these lines to your `~/.gitconfig`:

​	`Git` 也可以被配置为使用 SSH 来代替 HTTPS 来处理与给定前缀匹配的 URL。例如，要对所有的 GitHub 访问使用 SSH，请在 `~/.gitconfig` 中添加这些行：

```
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

### How should I manage package versions using "go get"? 我应该如何使用 "go get "来管理软件包的版本？

The Go toolchain has a built-in system for managing versioned sets of related packages, known as modules. Modules were introduced in [Go 1.11](https://go.dev/doc/go1.11#modules) and have been ready for production use since [1.14](https://go.dev/doc/go1.14#introduction).

​	Go工具链有一个内置系统，用于管理相关软件包的版本集，称为模块。模块是在 [Go 1.11](https://go.dev/doc/go1.11#modules) 中引入的，并从 [1.14](https://go.dev/doc/go1.14#introduction) 开始可用于生产。

To create a project using modules, run [`go mod init`](https://go.dev/ref/mod#go-mod-init). This command creates a `go.mod` file that tracks dependency versions.

​	要创建一个使用模块的项目，运行[go mod init](https://go.dev/ref/mod#go-mod-init)。这个命令会创建一个`go.mod`文件来跟踪依赖版本。

```shell
go mod init example/project
```

To add, upgrade, or downgrade a dependency, run [`go get`](https://go.dev/ref/mod#go-get):

​	要添加、升级或降级一个依赖项，请运行 `go get`：

```shell
go get golang.org/x/text@v0.3.5
```

See [Tutorial: Create a module](https://go.dev/doc/tutorial/create-module.html) for more information on getting started.

​	参见[Tutorial: Create a module](../GettingStarted/TutorialCreateAGoModule)，了解更多关于开始的信息。

See [Developing modules](https://go.dev/doc/#developing-modules) for guides on managing dependencies with modules.

​	请参阅 [Developing modules](../GoUserManual#developing-modules)，了解用模块管理依赖项的指南。

Packages within modules should maintain backward compatibility as they evolve, following the [import compatibility rule](https://research.swtch.com/vgo-import):

​	模块内的包在发展过程中应保持向后的兼容性，遵循[导入兼容性规则](https://research.swtch.com/vgo-import)：

> If an old package and a new package have the same import path, the new package must be backwards compatible with the old package.
> 
>如果一个旧包和一个新包有相同的导入路径，新的包必须向后兼容旧的包。

The [Go 1 compatibility guidelines](https://go.dev/doc/go1compat.html) are a good reference here: don't remove exported names, encourage tagged composite literals, and so on. If different functionality is required, add a new name instead of changing an old one.

​	[Go 1的兼容性准则](../Go1AndTheFutureOfGoPrograms)在这里是一个很好的参考：不要删除导出的名字，鼓励标记的复合字面量，等等。如果需要不同的功能，添加一个新的名字，而不是改变一个旧的名字。

Modules codify this with [semantic versioning](https://semver.org/) and semantic import versioning. If a break in compatibility is required, release a module at a new major version. Modules at major version 2 and higher require a [major version suffix](https://go.dev/ref/mod#major-version-suffixes) as part of their path (like `/v2`). This preserves the import compatibility rule: packages in different major versions of a module have distinct paths.

​	模块通过[语义版本化](https://semver.org/)和语义导入版本化来编纂这一点。如果需要打破兼容性，就在新的主版本上发布一个模块。主版本2和更高版本的模块需要一个[主版本后缀](../References/GoModulesReference/ModulesPackagesAndVersions#major-version-suffixes)作为其路径的一部分（如/v2）。这保留了导入兼容性规则：一个模块的不同主版本的包有不同的路径。

## Pointers and Allocation 指针和分配

### When are function parameters passed by value? 什么时候函数参数是按值传递的？

As in all languages in the C family, everything in Go is passed by value. That is, a function always gets a copy of the thing being passed, as if there were an assignment statement assigning the value to the parameter. For instance, passing an `int` value to a function makes a copy of the `int`, and passing a pointer value makes a copy of the pointer, but not the data it points to. (See a [later section](https://go.dev/doc/faq#methods_on_values_or_pointers) for a discussion of how this affects method receivers.)

​	如同C系列的所有语言一样，`Go中的所有东西都是以值传递的`。也就是说，`一个函数总是得到一个被传递的东西的副本`，就像有一个赋值语句将值分配给参数一样。例如，向一个函数传递一个`int`值，就会得到`int`的副本，而`传递一个指针值就会得到指针的副本`，但不会得到它所指向的数据。(参见[后面一节](#should-I-define-methods-on-values-or-pointers)，讨论这对方法接收者的影响）。

Map and slice values behave like pointers: they are descriptors that contain pointers to the underlying map or slice data. Copying a map or slice value doesn't copy the data it points to. Copying an interface value makes a copy of the thing stored in the interface value. If the interface value holds a struct, copying the interface value makes a copy of the struct. If the interface value holds a pointer, copying the interface value makes a copy of the pointer, but again not the data it points to.

​	`映射和切片值的行为类似于指针`：它们是包含指向底层映射或切片数据的指针的描述符。复制一个映射或切片值并不复制它所指向的数据。`复制一个接口值会复制存储在该接口值中的东西`。如果接口值持有一个结构，复制接口值就会复制该结构。`如果接口值持有一个指针，复制接口值会复制该指针，但同样不会复制它所指向的数据。`

Note that this discussion is about the semantics of the operations. Actual implementations may apply optimizations to avoid copying as long as the optimizations do not change the semantics.

​	请注意，这个讨论是关于操作的语义的。`实际的实现可以应用优化来避免复制`，只要这些优化不改变语义。

### When should I use a pointer to an interface? 什么时候应该使用一个指向接口的指针？

Almost never. Pointers to interface values arise only in rare, tricky situations involving disguising an interface value's type for delayed evaluation.

​	几乎不需要。指向接口值的指针只出现在罕见的、棘手的情况下，涉及到掩饰接口值的类型以延迟求值。

It is a common mistake to pass a pointer to an interface value to a function expecting an interface. The compiler will complain about this error but the situation can still be confusing, because sometimes a [pointer is necessary to satisfy an interface](https://go.dev/doc/faq#different_method_sets). The insight is that although a pointer to a concrete type can satisfy an interface, with one exception *a pointer to an interface can never satisfy an interface*.

​	将一个接口值的指针传递给一个需要接口参数的函数，是一个常见的错误。编译器会抱怨这个错误，但这种情况仍然会让人困惑，因为有时需要[一个指针来满足一个接口](#why-do-t-and-t-have-different-method-sets)。我们的见解是，尽管一个指向具体类型的指针可以满足一个接口，但有一个例外，一个指向接口的指针永远不能满足一个接口。

Consider the variable declaration,

​	考虑一下这个变量声明，

```go linenums="1"
var w io.Writer
```

The printing function `fmt.Fprintf` takes as its first argument a value that satisfies `io.Writer`—something that implements the canonical `Write` method. Thus we can write

​	打印函数`fmt.Fprintf`把一个满足`io.Writer`的值作为它的第一个参数——实现了典型的`Write`方法的东西。因此，我们可以写

```go linenums="1"
fmt.Fprintf(w, "hello, world\n")
```

If however we pass the address of `w`, the program will not compile.

然而，如果我们传递`w`的地址，程序将无法编译。

```go linenums="1"
fmt.Fprintf(&w, "hello, world\n") // Compile-time error. => 编译时错误。
```

The one exception is that any value, even a pointer to an interface, can be assigned to a variable of empty interface type (`interface{}`). Even so, it's almost certainly a mistake if the value is a pointer to an interface; the result can be confusing.

​	有一个例外是，任何值，甚至是一个接口的指针，都可以被分配给一个空接口类型的变量（`interface{}`）。即便如此，如果值是一个指向接口的指针，几乎可以肯定是个错误；其结果可能是混乱的。

### Should I define methods on values or pointers? 我应该在值或指针上定义方法？

``` go
func (s *MyStruct) pointerMethod() { } // method on pointer
func (s MyStruct)  valueMethod()   { } // method on value
```

For programmers unaccustomed to pointers, the distinction between these two examples can be confusing, but the situation is actually very simple. When defining a method on a type, the receiver (`s` in the above examples) behaves exactly as if it were an argument to the method. Whether to define the receiver as a value or as a pointer is the same question, then, as whether a function argument should be a value or a pointer. There are several considerations.

​	对于不习惯指针的程序员来说，这两个例子之间的区别可能会让人感到困惑，但实际上情况非常简单。当在一个类型上定义一个方法时，接收器（上述例子中的`s`）`的行为就像它是该方法的一个参数一样`。那么，将接收器定义为一个值还是一个指针，`与函数实参应该是一个值还是一个指针的问题是一样的`。这里有几个考虑因素。

First, and most important, does the method need to modify the receiver? If it does, the receiver *must* be a pointer. (Slices and maps act as references, so their story is a little more subtle, but for instance to change the length of a slice in a method the receiver must still be a pointer.) In the examples above, if `pointerMethod` modifies the fields of `s`, the caller will see those changes, but `valueMethod` is called with a copy of the caller's argument (that's the definition of passing a value), so changes it makes will be invisible to the caller.

​	首先，也是最重要的，方法是否需要修改接收器？如果需要，那么接收器必须是一个指针。(切片和映射作为引用，所以它们的故事更微妙一些，`但例如在一个方法中改变切片的长度，接收器仍然必须是一个指针`。) 在上面的例子中，如果`pointerMethod`修改了`s`的字段，那么调用者会看到这些变化，但是`valueMethod`是用调用者实参的副本来调用的（这就是传递值的定义），所以它所做的变化对调用者来说是不可见的。

By the way, in Java method receivers are always pointers, although their pointer nature is somewhat disguised (and there is a proposal to add value receivers to the language). It is the value receivers in Go that are unusual.

​	顺便说一下，`在Java中，方法接收者总是指针`，尽管它们的指针性质在某种程度上被掩盖了（有一个建议是在语言中增加值接收者）。Go中的值接收器是不寻常的。

Second is the consideration of efficiency. If the receiver is large, a big `struct` for instance, it will be much cheaper to use a pointer receiver.

​	其次是对效率的考虑。如果接收器很大，例如一个大的结构体，那么使用指针接收器就会廉价很多。

Next is consistency. If some of the methods of the type must have pointer receivers, the rest should too, so the method set is consistent regardless of how the type is used. See the section on [method sets](https://go.dev/doc/faq#different_method_sets) for details.

​	其次是一致性。如果该类型的一些方法必须有指针接收器，那么其他的也应该有，所以无论该类型如何使用，方法集都是一致的。详见关于[方法集](#why-do-t-and-t-have-different-method-sets)的章节。

For types such as basic types, slices, and small `structs`, a value receiver is very cheap so unless the semantics of the method requires a pointer, a value receiver is efficient and clear.

​	对于基本类型、切片和小结构体等类型，值接收器是非常廉价的，所以除非方法的语义需要指针，否则值接收器是高效且清晰的。

### What's the difference between new and make? new和make之间有什么区别？

In short: `new` allocates memory, while `make` initializes the slice, map, and channel types.

​	简而言之：`new`是分配内存，而`make`是初始化slice、map和channel类型。

See the [relevant section of Effective Go](https://go.dev/doc/effective_go.html#allocation_new) for more details.

更多细节请参见[Effective Go](../UsingAndUnderstandingGo/EffectiveGo#allocation-with-new-new)的相关章节。

### What is the size of an `int` on a 64 bit machine? 在64位机器上int的大小是多少？

The sizes of `int` and `uint` are implementation-specific but the same as each other on a given platform. For portability, code that relies on a particular size of value should use an explicitly sized type, like `int64`. On 32-bit machines the compilers use 32-bit integers by default, while on 64-bit machines integers have 64 bits. (Historically, this was not always true.)

​	`int`和`uint`的大小是由具体实现决定的，但在特定的平台上彼此是一样的。为了可移植性，依赖特定大小的值的代码应该使用明确大小的类型，如`int64`。在32位机器上，编译器默认使用32位整数，而在64位机器上，整数有64位。(历史上，这并不总是正确的）。

On the other hand, floating-point scalars and complex types are always sized (there are no `float` or `complex` basic types), because programmers should be aware of precision when using floating-point numbers. The default type used for an (untyped) floating-point constant is `float64`. Thus `foo` `:=` `3.0` declares a variable `foo` of type `float64`. For a `float32` variable initialized by an (untyped) constant, the variable type must be specified explicitly in the variable declaration:

​	另一方面，浮点标量和复数类型总是有大小的（没有`float`或`complex`的基本类型），因为程序员在使用浮点数字时应该注意精度。用于（无类型的）浮点常量的默认类型是`float64`。因此`foo :=3.0`声明了一个`float64`类型的变量`foo`。对于一个由（无类型）常量初始化的`float32`变量，必须在变量声明中明确指定变量类型：

```
var foo float32 = 3.0
```

Alternatively, the constant must be given a type with a conversion as in `foo := float32(3.0)`.

​	或者，常量必须被赋予一个转换的类型，如`foo := float32(3.0)`。

### How do I know whether a variable is allocated on the heap or the stack? 我如何知道一个变量是在堆上还是在栈上分配的？

From a correctness standpoint, you don't need to know. Each variable in Go exists as long as there are references to it. The storage location chosen by the implementation is irrelevant to the semantics of the language.

​	从正确性的角度来看，您不需要知道。只要有对它的引用，Go中的每个变量就存在。实现所选择的存储位置与语言的语义无关。

The storage location does have an effect on writing efficient programs. When possible, the Go compilers will allocate variables that are local to a function in that function's stack frame. However, if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors. Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.

​	存储位置确实对编写高效程序有影响。在可能的情况下，Go编译器会在一个函数的栈框架中分配属于该函数的局部变量。但是，如果编译器不能证明该变量在函数返回后没有被引用，那么编译器必须在垃圾收集的堆上分配该变量，以避免`悬空指针错误`。另外，如果一个局部变量非常大，把它存储在`堆（heap ）`上而不是`栈（stack）`上可能更有意义。

In the current compilers, if a variable has its address taken, that variable is a candidate for allocation on the heap. However, a basic *escape analysis* recognizes some cases when such variables will not live past the return from the function and can reside on the stack.

​	在目前的编译器中，如果一个变量的地址被占用，那么这个变量就是在堆上分配的候选变量。然而，基本的转义分析认识到有些情况下，这种变量不会活过函数的返回，可以驻留在栈上。

### Why does my Go process use so much virtual memory? 为什么我的Go进程会使用这么多的虚拟内存？

The Go memory allocator reserves a large region of virtual memory as an arena for allocations. This virtual memory is local to the specific Go process; the reservation does not deprive other processes of memory.

​	Go内存分配器保留了一个大的虚拟内存区域作为分配的场所。这个虚拟内存是特定Go进程的本地内存；保留的内存不会剥夺其他进程的内存。

To find the amount of actual memory allocated to a Go process, use the Unix `top` command and consult the `RES` (Linux) or `RSIZE` (macOS) columns.

​	要找到分配给Go进程的实际内存量，请使用Unix的`top`命令，并查阅`RES`（Linux）或`RSIZE`（macOS）列。

## Concurrency 并发性

### What operations are atomic? What about mutexes? 哪些操作是原子性的？互斥是什么？

A description of the atomicity of operations in Go can be found in the [Go Memory Model](https://go.dev/ref/mem) document.

​	关于Go中操作的原子性的描述可以在[Go Memory Model](../References/TheGoMemoryModel)文档中找到。

Low-level synchronization and atomic primitives are available in the [sync](https://go.dev/pkg/sync) and [sync/atomic](https://go.dev/pkg/sync/atomic) packages. These packages are good for simple tasks such as incrementing reference counts or guaranteeing small-scale mutual exclusion.

​	低级别的同步和原子原语可以在[sync](https://go.dev/pkg/sync)和[sync/atomic](https://go.dev/pkg/sync/atomic)包中找到。这些包适合于简单的任务，如增加引用计数或保证小规模的互斥。

For higher-level operations, such as coordination among concurrent servers, higher-level techniques can lead to nicer programs, and Go supports this approach through its goroutines and channels. For instance, you can structure your program so that only one goroutine at a time is ever responsible for a particular piece of data. That approach is summarized by the original [Go proverb](https://www.youtube.com/watch?v=PAAkCSZUG1c),

​	对于更高层次的操作，如并发服务器之间的协调，更高层次的技术可以带来更好的程序，Go通过其goroutines和通道支持这种方法。例如，您可以构造您的程序，使每次只有一个goroutine负责一个特定的数据。这种方法被[Go 谚语](https://www.youtube.com/watch?v=PAAkCSZUG1c)所概括，

Do not communicate by sharing memory. Instead, share memory by communicating.

`不要通过共享内存进行通信。相反，通过通信来共享内存。`

See the [Share Memory By Communicating](https://go.dev/doc/codewalk/sharemem/) code walk and its [associated article](https://blog.golang.org/2010/07/share-memory-by-communicating.html) for a detailed discussion of this concept.

​	关于这个概念的详细讨论，请参见 [通过通信共享内存](https://go.dev/doc/codewalk/sharemem/) 的代码练习及其[相关文章]({{< ref "/goBlog/2010/ShareMemoryByCommunicating" >}})。

Large concurrent programs are likely to borrow from both these toolkits.

​	大型并发程序可能会借用这两个工具包。

### Why doesn't my program run faster with more CPUs? 为什么我的程序在更多的CPU下运行得不快？

Whether a program runs faster with more CPUs depends on the problem it is solving. The Go language provides concurrency primitives, such as goroutines and channels, but concurrency only enables parallelism when the underlying problem is intrinsically parallel. Problems that are intrinsically sequential cannot be sped up by adding more CPUs, while those that can be broken into pieces that can execute in parallel can be sped up, sometimes dramatically.

​	一个程序是否在更多的CPU下运行得更快，取决于它所解决的问题。Go语言提供了并发原语，如goroutines和channel，但只有当底层问题本质上是并行的时候，并发才能实现并行化。本质上是顺序的问题不能通过添加更多的CPU来加速，而那些可以被分解成可以并行执行的片段的问题则可以加速，有时甚至可以大大加速。

Sometimes adding more CPUs can slow a program down. In practical terms, programs that spend more time synchronizing or communicating than doing useful computation may experience performance degradation when using multiple OS threads. This is because passing data between threads involves switching contexts, which has significant cost, and that cost can increase with more CPUs. For instance, the [prime sieve example](https://go.dev/ref/spec#An_example_package) from the Go specification has no significant parallelism although it launches many goroutines; increasing the number of threads (CPUs) is more likely to slow it down than to speed it up.

​	有时，增加更多的CPU会使程序变慢。在实际应用中，当使用多个操作系统线程时，那些花在同步或通信上的时间多于进行有用的计算的程序可能会出现性能下降。这是因为在线程之间传递数据需要切换上下文，这需要很大的成本，而这种成本会随着CPU的增加而增加。例如，Go规范中的[素数筛例子](../References/LanguageSpecification/Packages#an-example-package)没有明显的并行性，尽管它启动了许多goroutine；增加线程（CPU）的数量更有可能使它变慢而不是变快。

For more detail on this topic see the talk entitled [Concurrency is not Parallelism](https://blog.golang.org/2013/01/concurrency-is-not-parallelism.html).

​	关于这个话题的更多细节，请看题为 [并发性不是并行性]({{< ref "/goBlog/2013/ConcurrencyIsNotParallelism" >}}) 的讲座。

### How can I control the number of CPUs? 我怎样才能控制CPU的数量？

The number of CPUs available simultaneously to executing goroutines is controlled by the `GOMAXPROCS` shell environment variable, whose default value is the number of CPU cores available. Programs with the potential for parallel execution should therefore achieve it by default on a multiple-CPU machine. To change the number of parallel CPUs to use, set the environment variable or use the similarly-named [function](https://go.dev/pkg/runtime/#GOMAXPROCS) of the runtime package to configure the run-time support to utilize a different number of threads. Setting it to 1 eliminates the possibility of true parallelism, forcing independent goroutines to take turns executing.

​	可用于同时执行goroutines的CPU数量由`GOMAXPROCS` shell环境变量控制，其默认值是可用的CPU核心数量。因此，具有并行执行潜力的程序应该在多CPU机器上默认实现。要改变使用的并行CPU数量，可以设置环境变量或使用 runtime 包的类似[功能](https://go.dev/pkg/runtime/#GOMAXPROCS)来配置运行时支持，以利用不同数量的线程。`将其设置为1会消除真正并行的可能性，迫使独立的goroutine轮流执行`。

The runtime can allocate more threads than the value of `GOMAXPROCS` to service multiple outstanding I/O requests. `GOMAXPROCS` only affects how many goroutines can actually execute at once; arbitrarily more may be blocked in system calls.

​	runtime 可以分配比`GOMAXPROCS`值更多的线程，以服务于多个未处理的I/O请求。`GOMAXPROCS`只影响到有多少个goroutine可以同时实际执行；任意多的goroutine可能会在系统调用中被阻塞。

Go's goroutine scheduler is not as good as it needs to be, although it has improved over time. In the future, it may better optimize its use of OS threads. For now, if there are performance issues, setting `GOMAXPROCS` on a per-application basis may help.

​	Go的`goroutine调度器`并不像它需要的那样好，尽管它随着时间的推移已经有所改进。在未来，它可能会更好地优化其对操作系统线程的使用。目前，如果有性能问题，在每个应用的基础上设置`GOMAXPROCS`可能会有帮助。

### Why is there no goroutine ID? 为什么没有goroutine的ID？

Goroutines do not have names; they are just anonymous workers. They expose no unique identifier, name, or data structure to the programmer. Some people are surprised by this, expecting the `go` statement to return some item that can be used to access and control the goroutine later.

​	goroutines没有名字；它们只是匿名的工作程序。它们没有向程序员暴露出唯一的标识符、名称或数据结构。有些人对此感到惊讶，以为`go`语句会返回一些可以用来访问和控制goroutine的项目。

The fundamental reason goroutines are anonymous is so that the full Go language is available when programming concurrent code. By contrast, the usage patterns that develop when threads and goroutines are named can restrict what a library using them can do.

​	goroutines是匿名的，其根本原因是为了在编程并发代码时可以使用完整的Go语言。相比之下，当线程和goroutines被命名时，形成的使用模式会限制使用它们的库的能力。

Here is an illustration of the difficulties. Once one names a goroutine and constructs a model around it, it becomes special, and one is tempted to associate all computation with that goroutine, ignoring the possibility of using multiple, possibly shared goroutines for the processing. If the `net/http` package associated per-request state with a goroutine, clients would be unable to use more goroutines when serving a request.

​	下面是一个困难的例子。一旦人们命名了一个goroutine，并围绕它构建了一个模型，它就变得很特别，人们就会倾向于将所有的计算与该goroutines联系起来，而忽略了使用多个可能共享的goroutine进行处理的可能性。如果 `net/http` 包将每个请求的状态与一个 goroutine 关联起来，那么在提供请求时，客户端将无法使用更多的 goroutine。

Moreover, experience with libraries such as those for graphics systems that require all processing to occur on the "main thread" has shown how awkward and limiting the approach can be when deployed in a concurrent language. The very existence of a special thread or goroutine forces the programmer to distort the program to avoid crashes and other problems caused by inadvertently operating on the wrong thread.

​	此外，像那些要求所有处理都发生在 "主线程 "上的图形系统库的经验表明，当部署在一种并发语言中时，这种方法是多么的笨拙和局限。特殊线程或goroutine的存在，迫使程序员扭曲程序，以避免崩溃和其他因无意中在错误线程上操作而引起的问题。

For those cases where a particular goroutine is truly special, the language provides features such as channels that can be used in flexible ways to interact with it.

​	对于那些特定的goroutine确实很特殊的情况，语言提供了诸如通道等功能，可以用灵活的方式与之交互。

## Functions and Methods 函数和方法

### Why do T and *T have different method sets? 为什么T和`*T`有不同的方法集？

As the [Go specification](https://go.dev/ref/spec#Types) says, the method set of a type `T` consists of all methods with receiver type `T`, while that of the corresponding pointer type `*T` consists of all methods with receiver `*T` or `T`. That means the method set of `*T` includes that of `T`, but not the reverse.

​	正如[Go规范](../References/LanguageSpecification/Types)所说，一个类型`T`的方法集包括所有接收者类型为`T`的方法，而相应的指针类型`*T`的方法集包括所有接收者为`*T`或`T`的方法，这意味着`*T`的方法集包括`T`的方法集，但不是相反。

This distinction arises because if an interface value contains a pointer `*T`, a method call can obtain a value by dereferencing the pointer, but if an interface value contains a value `T`, there is no safe way for a method call to obtain a pointer. (Doing so would allow a method to modify the contents of the value inside the interface, which is not permitted by the language specification.)

​	这种区别的产生是因为如果一个接口值包含一个指针`*T`，方法调用可以通过取消引用指针来获得一个值，但是如果一个接口值包含一个值`T`，方法调用就没有安全的方法来获得一个指针。(这样做将允许一个方法修改接口内的值的内容，这是语言规范所不允许的）。

Even in cases where the compiler could take the address of a value to pass to the method, if the method modifies the value the changes will be lost in the caller. As an example, if the `Write` method of [`bytes.Buffer`](https://go.dev/pkg/bytes/#Buffer) used a value receiver rather than a pointer, this code:

​	即使在编译器可以获取传递给方法的值的地址的情况下，如果方法修改了这个值，则更改将在调用方中丢失。举个例子，如果[bytes.Buffer](https://go.dev/pkg/bytes/#Buffer)的`Write`方法使用了一个值接收器而不是一个指针，那么这段代码：

```go linenums="1"
var buf bytes.Buffer
io.Copy(buf, os.Stdin)
```

would copy standard input into a *copy* of `buf`, not into `buf` itself. This is almost never the desired behavior.

将会把标准输入复制到`buf`的一个副本中，而不是复制到`buf`本身。这几乎不是我们想要的行为。

### What happens with closures running as goroutines? 作为goroutines运行的闭包会发生什么？

Some confusion may arise when using closures with concurrency. Consider the following program:

​	在使用具有并发性的闭包时可能会出现一些混乱。考虑一下下面的程序：

```go linenums="1"
func main() {
    done := make(chan bool)

    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }

    // wait for all goroutines to complete before exiting => 等待所有 goroutine 完成，然后退出
    for _ = range values {
        <-done
    }
}
```

One might mistakenly expect to see `a, b, c` as the output. What you'll probably see instead is `c, c, c`. This is because each iteration of the loop uses the same instance of the variable `v`, so each closure shares that single variable. When the closure runs, it prints the value of `v` at the time `fmt.Println` is executed, but `v` may have been modified since the goroutine was launched. To help detect this and other problems before they happen, run [`go vet`](https://go.dev/cmd/go/#hdr-Run_go_tool_vet_on_packages).

​	人们可能误以为会看到`a, b, c`作为输出。您可能看到的是 `c, c, c`。这是因为循环的每一次迭代都使用同一个变量`v`的实例，所以每个闭包都共享这个单一的变量。当闭包运行时，它打印的是执行 `fmt.Println` 时的 `v` 值，但 `v` 可能在 goroutine 启动后被修改过。为了帮助在发生这种情况和其他问题之前发现它们，运行[go vet](https://go.dev/cmd/go/#hdr-Run_go_tool_vet_on_packages)。

To bind the current value of `v` to each closure as it is launched, one must modify the inner loop to create a new variable each iteration. One way is to pass the variable as an argument to the closure:

​	为了在每个闭包启动时将`v`的当前值与之绑定，必须修改内循环以在每个迭代中创建一个新变量。一种方法是将该变量作为参数传递给闭包：

```go linenums="1" hl_lines="2 2"
    for _, v := range values {
        go func(u string) {
            fmt.Println(u)
            done <- true
        }(v)
    }
```

In this example, the value of `v` is passed as an argument to the anonymous function. That value is then accessible inside the function as the variable `u`.

​	在这个例子中，`v`的值被作为一个参数传递给匿名函数。然后，该值可以在函数中作为变量`u`被访问。

Even easier is just to create a new variable, using a declaration style that may seem odd but works fine in Go:

​	更简单的方法是`直接创建一个新的变量`，使用的声明方式可能看起来很奇怪，但在Go中很好用：

```go linenums="1" hl_lines="2 2"
    for _, v := range values {
        v := v // create a new 'v'.
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }
```

This behavior of the language, not defining a new variable for each iteration, may have been a mistake in retrospect. It may be addressed in a later version but, for compatibility, cannot change in Go version 1.

​	语言的这种行为，即不为每个迭代定义一个新的变量，现在回想起来可能是一个错误。`它可能会在以后的版本中得到解决`，但`为了兼容，在Go 1版本中不能改变`。

## Control flow 控制流

### Why does Go not have the `?:` operator? 为什么Go没有?: 操作符？

There is no ternary testing operation in Go. You may use the following to achieve the same result:

​	Go中没有三元测试操作。您可以使用下面的方法来实现同样的结果：

```go linenums="1"
if expr {
    n = trueVal
} else {
    n = falseVal
}
```

The reason `?:` is absent from Go is that the language's designers had seen the operation used too often to create impenetrably complex expressions. The `if-else` form, although longer, is unquestionably clearer. A language needs only one conditional control flow construct.

​	Go中没有`?:`的原因是该语言的设计者看到该操作经常被用来创建令人难以理解的复杂表达式。if-else形式虽然较长，但无疑是更清晰的。一门语言只需要一个条件控制流结构。

## Type Parameters 类型参数

### Why does Go have type parameters? 为什么 Go 有类型参数？

Type parameters permit what is known as generic programming, in which functions and data structures are defined in terms of types that are specified later, when those functions and data structures are used. For example, they make it possible to write a function that returns the minimum of two values of any ordered type, without having to write a separate version for each possible type. For a more in-depth explanation with examples see the blog post [Why Generics?](https://go.dev/blog/why-generics).

​	类型参数允许所谓的泛型编程，其中函数和数据结构是以类型来定义的，这些类型在以后使用这些函数和数据结构时被指定。例如，它们使我们有可能编写一个返回任何有序类型的两个值的最小值的函数，而不必为每个可能的类型编写一个单独的版本。关于更深入的解释和例子，请看博文 [Why Generics?]({{< ref "/goBlog/2019/WhyGenerics" >}})。

### How are generics implemented in Go? 泛型是如何在Go中实现的？

The compiler can choose whether to compile each instantiation separately or whether to compile reasonably similar instantiations as a single implementation. The single implementation approach is similar to a function with an interface parameter. Different compilers will make different choices for different cases. The standard Go 1.18 compiler ordinarily emits a single instantiation for every type argument with the same shape, where the shape is determined by properties of the type such as the size and the location of pointers that it contains. Future releases will experiment with the tradeoff between compile time, run-time efficiency, and code size.

​	编译器可以选择是单独编译每个实例，还是将合理相似的实例作为一个单一的实现来编译。单一实现的方法类似于带有接口参数的函数。不同的编译器会针对不同的情况做出不同的选择。标准的Go 1.18编译器通常为每个具有相同形状的类型实参发出一个单一的实例，其中的形状是由类型的属性决定的，例如它包含的指针的大小和位置。未来的版本将对编译时、运行时效率和代码大小之间的权衡进行试验。

### How do generics in Go compare to generics in other languages? Go中的泛型与其他语言的泛型相比如何？

The basic functionality in all languages is similar: it is possible to write types and functions using types that are specified later. That said, there are some differences.

​	所有语言中的基本功能都是相似的：可以使用以后指定的类型来编写类型和函数。尽管如此，还是有一些区别。

- Java

  In Java, the compiler checks generic types at compile time but removes the types at run time. This is known as [type erasure](https://en.wikipedia.org/wiki/Generics_in_Java#Problems_with_type_erasure). For example, a Java type known as `List<Integer>` at compile time will become the non-generic type `List` at run time. This means, for example, that when using the Java form of type reflection it is impossible to distinguish a value of type `List<Integer>` from a value of type `List<Float>`. In Go the reflection information for a generic type includes the full compile-time type information.

  在Java中，编译器在编译时检查泛型类型，但在运行时删除这些类型。这被称为[类型擦除](https://en.wikipedia.org/wiki/Generics_in_Java#Problems_with_type_erasure)。例如，一个在编译时被称为`List<Integer>`的Java类型在运行时将变成非泛型类型`List`。这意味着，例如，当使用Java形式的类型反射时，不可能区分`List<Integer>`类型的值和`List<Float>`类型的值。在Go中，泛型类型的反射信息包括完整的编译时类型信息。

  Java uses type wildcards such as `List<? extends Number>` or `List<? super Number>` to implement generic covariance and contravariance. Go does not have these concepts, which makes generic types in Go much simpler. 

  Java使用类型通配符，如`List<? extends Number>`或`List<? super Number>`来实现泛型的协变和反变。Go没有这些概念，这使得Go中的泛型变得更加简单。

- C++

  Traditionally C++ templates do not enforce any constraints on type arguments, although C++20 supports optional constraints via [concepts](https://en.wikipedia.org/wiki/Concepts_(C%2B%2B)). In Go constraints are mandatory for all type parameters. C++20 concepts are expressed as small code fragments that must compile with the type arguments. Go constraints are interface types that define the set of all permitted type arguments.C++ supports template metaprogramming; Go does not. In practice, all C++ compilers compile each template at the point where it is instantiated; as noted above, Go can and does use different approaches for different instantiations.

  传统上C++模板不对类型实参执行任何约束，尽管C++20通过[概念](https://en.wikipedia.org/wiki/Concepts_(C%2B%2B))支持可选的约束。在Go中，对所有类型参数的约束是强制性的。C++20的概念被表达为小的代码片段，必须与类型实参一起编译。Go的约束是定义所有允许的类型实参集合的接口类型。C++支持模板元编程；Go则不支持。在实践中，所有的C++编译器都是在每个模板被实例化的地方进行编译；如上所述，Go可以也确实对不同的实例化使用不同的方法。

- Rust

  The Rust version of constraints is known as trait bounds. In Rust the association between a trait bound and a type must be defined explicitly, either in the crate that defines the trait bound or the crate that defines the type. In Go type arguments implicitly satisfy constraints, just as Go types implicitly implement interface types. The Rust standard library defines standard traits for operations such as comparison or addition; the Go standard library does not, as these can be expressed in user code via interface types.

  Rust版本的约束被称为`特性绑定`。在Rust中，特性绑定和类型之间的关联必须明确定义，要么在定义特性绑定的板块中，要么在定义类型的板块中。在Go中，类型实参隐式地满足约束，就像Go类型隐式地实现接口类型一样。Rust标准库为比较或加法等操作定义了标准特性；Go标准库则没有，因为这些可以通过接口类型在用户代码中表达。

- Python

  Python is not a statically typed language, so one can reasonably say that all Python functions are always generic by default: they can always be called with values of any type, and any type errors are detected at run time.
  
  Python 不是静态类型语言，因此可以合理地说，所有 Python 函数在默认情况下总是泛型的：它们总是可以用任何类型的值来调用，并且在运行时检测到任何类型的错误。

### Why does Go use square brackets for type parameter lists? 为什么Go对类型参数列表使用方括号？

Java and C++ use angle brackets for type parameter lists, as in Java `List<Integer>` and C++ `std::vector<int>`. However, that option was not available for Go, because it leads to a syntactic problem: when parsing code within a function, such as `v := F<T>`, at the point of seeing the `<` it's ambiguous whether we are seeing an instantiation or an expression using the `<` operator. This is very difficult to resolve without type information.

​	Java 和 C++ 对类型参数列表使用`角括号`，如 Java `List<Integer>` 和 C++ `std::vector<int>`。然而，Go没有这个选项，`因为它导致了一个语法问题`：当解析一个函数内的代码时，例如`v := F<T>`，在看到`<`的时候，我们看到的是一个实例化还是一个使用`<`操作符的表达式，这一点是模糊的。如果没有类型信息，这是很难解决的。

For example, consider a statement like

例如，考虑一个语句，如

```go linenums="1"
    a, b = w < x, y > (z)
```

Without type information, it is impossible to decide whether the right hand side of the assignment is a pair of expressions (`w < x` and `y > z`), or whether it is a generic function instantiation and call that returns two result values (`(w<x, y>)(z)`).

​	如果没有类型信息，就不可能决定赋值的右边是一对表达式（`w < x and y > z`），还是一个返回两个结果值的泛型函数实例化和调用（`(w<x, y>)(z)`）。

It is a key design decision of Go that parsing be possible without type information, which seems impossible when using angle brackets for generics.

​	`Go的一个关键设计决定`是`在没有类型信息的情况下进行解析`，这在使用角括号表示泛型时似乎是不可能的。

Go is not unique or original in using square brackets; there are other languages such as Scala that also use square brackets for generic code.

​	Go在使用方括号方面并不是唯一的或原创的；还有其他语言，如`Scala`，也使用方括号来表示泛型代码。

### Why does Go not support methods with type parameters? 为什么 Go 不支持带有类型参数的方法？

Go permits a generic type to have methods, but, other than the receiver, the arguments to those methods cannot use parameterized types. The methods of a type determines the interfaces that the type implements, but it is not clear how this would work with parameterized arguments for methods of generic types. It would require either instantiating functions at run time or instantiating every generic function for every possible type argument. Neither approach seems feasible. For more details, including an example, see the [proposal](https://go.dev/design/43651-type-parameters#no-parameterized-methods). Instead of methods with type parameters, use top-level functions with type parameters, or add the type parameters to the receiver type.

​	Go 允许泛型类型拥有方法，`但除了接收者之外，这些方法的参数不能使用参数化类型`。一个类型的方法决定了该类型所实现的接口，但不清楚这对泛型的方法的参数化实参如何操作。这将需要在运行时实例化函数，或者为每个可能的类型实参实例化每个泛型函数。这两种方法似乎都不可行。更多细节，包括一个例子，请看[提案](https://go.dev/design/43651-type-parameters#no-parameterized-methods)。不使用带类型参数的方法，而是使用带类型参数的顶层函数，或者将类型参数添加到接收器类型中。

### Why can't I use a more specific type for the receiver of a parameterized type? 为什么我不能为参数化类型的接收器使用一个更具体的类型？

The method declarations of a generic type are written with a receiver that includes the type parameter names. Some people think that a specific type can be used, producing a method that only works for certain type arguments:

​	泛型类型的方法声明是用一个包括类型参数名称的接收器来写的。有些人认为可以使用一个特定的类型，产生一个只对某些类型实参有效的方法：

```go linenums="1"
type S[T any] struct { f T }

func (s S[string]) Add(t string) string {
    return s.f + t
}
```

This fails with a compiler error like `operator + not defined on s.f (variable of type string constrained by any)`, even though the `+` operator does of course work on the predeclared type `string`.

​	这个操作失败了，出现了编译器错误，如操作符`+`没有定义在`s.f上`（由any限制的字符串类型的变量），尽管`+`操作符当然对预先声明的字符串类型有效。

This is because the use of `string` in the declaration of the method `Add` is simply introducing a name for the type parameter, and the name is `string`. This is a valid, if strange, thing to do. The field `s.f` has type `string`, not the usual predeclared type `string`, but rather the type parameter of `S`, which in this method is named `string`. Since the constraint of the type parameter is `any`, the `+` operator is not permitted.

​	这是因为在方法`Add`的声明中使用`string`只是为类型参数引入一个名字，而这个名字就是`string`。这是一件有效的，尽管很奇怪的事情。字段`s.f`的类型是`string`，不是通常预先声明的`string`类型，而是`S`的类型参数，在这个方法中被命名为`string`。由于类型参数的约束是`any`，所以不允许使用`+`运算符。=>仍有疑问？？

### Why can't the compiler infer the type argument in my program? 为什么编译器不能推断出我程序中的类型实参？

There are many cases where a programmer can easily see what the type argument for a generic type or function must be, but the language does not permit the compiler to infer it. Type inference is intentionally limited to ensure that there is never any confusion as to which type is inferred. Experience with other languages suggests that unexpected type inference can lead to considerable confusion when reading and debugging a program. It is always possible to specify the explicit type argument to be used in the call. In the future new forms of inference may be supported, as long as the rules remain simple and clear.

​	有很多情况下，程序员可以很容易地看到一个泛型或函数的类型实参必须是什么，但`语言不允许编译器推断它`。`类型推断是有意限制的，以确保在推断哪种类型方面不会有任何混淆。`其他语言的经验表明，在阅读和调试程序时，意外的类型推断会导致相当大的混乱。我们总是可以指定在调用中使用的显式类型实参。在未来，新的推理形式可能会被支持，只要规则保持简单和清晰。

## Packages and Testing 包和测试

### How do I create a multifile package? 我如何创建一个多文件包？

Put all the source files for the package in a directory by themselves. Source files can refer to items from different files at will; there is no need for forward declarations or a header file.

​	把包的所有源文件单独放在一个目录中。源文件可以随意引用不同文件中的项；不需要正向声明或头文件。

Other than being split into multiple files, the package will compile and test just like a single-file package.

​	除了被分割成多个文件外，该包的编译和测试就像一个单文件包。

### How do I write a unit test? 如何写一个单元测试？

Create a new file ending in `_test.go` in the same directory as your package sources. Inside that file, `import "testing"` and write functions of the form

​	在与您的包源文件相同的目录下，创建一个以`_test.go`结尾的新文件。在该文件中，`import "testing"`并编写以下形式的函数

```go linenums="1"
func TestFoo(t *testing.T) {
    ...
}
```

Run `go test` in that directory. That script finds the `Test` functions, builds a test binary, and runs it.

​	在该目录下运行`go test`。该脚本找到`Test`的函数，建立一个测试二进制文件，并运行它。

See the [How to Write Go Code](https://go.dev/doc/code.html) document, the [`testing`](https://go.dev/pkg/testing/) package and the [`go test`](https://go.dev/cmd/go/#hdr-Test_packages) subcommand for more details.

​	更多细节请参见[How to Write Go Code](../GettingStarted/HowToWriteGoCode)文档，[testing](https://go.dev/pkg/testing/)包和[go test](https://go.dev/cmd/go/#hdr-Test_packages)子命令。

### Where is my favorite helper function for testing? 我最喜欢的测试辅助函数在哪里？

Go's standard [`testing`](https://go.dev/pkg/testing/) package makes it easy to write unit tests, but it lacks features provided in other language's testing frameworks such as assertion functions. An [earlier section](https://go.dev/doc/faq#assertions) of this document explained why Go doesn't have assertions, and the same arguments apply to the use of `assert` in tests. Proper error handling means letting other tests run after one has failed, so that the person debugging the failure gets a complete picture of what is wrong. It is more useful for a test to report that `isPrime` gives the wrong answer for 2, 3, 5, and 7 (or for 2, 4, 8, and 16) than to report that `isPrime` gives the wrong answer for 2 and therefore no more tests were run. The programmer who triggers the test failure may not be familiar with the code that fails. Time invested writing a good error message now pays off later when the test breaks.

​	Go 的标准[testing](https://go.dev/pkg/testing/)包使编写单元测试变得容易，但它缺乏其他语言的测试框架所提供的功能，如断言函数。本文的[前一部分](#why-does-go-not-have-assertions)解释了为什么Go没有断言，同样的论点也适用于测试中使用断言。正确的错误处理意味着在一个测试失败后让其他测试运行，这样调试失败的人就可以得到一个完整的错误信息。对于测试来说，报告 `isPrime` 对2、3、5和7(或对2、4、8和16)给出错误答案比报告 `isPrime` 对2给出错误答案更有用，因此没有（什么比）运行更多的测试更有用。触发测试失败的程序员可能不熟悉失败的代码。现在花时间写一个好的错误信息，在以后测试失败时就会得到回报。

A related point is that testing frameworks tend to develop into mini-languages of their own, with conditionals and controls and printing mechanisms, but Go already has all those capabilities; why recreate them? We'd rather write tests in Go; it's one fewer language to learn and the approach keeps the tests straightforward and easy to understand.

​	与此相关的一点是，测试框架往往会发展成自己的迷您语言，带有条件、控件和打印机制，但Go已经具备所有这些功能，为什么还要重新创建它们呢？我们宁愿用Go来写测试；这样就少了一种需要学习的语言，而且这种方法可以使测试简单明了，易于理解。

If the amount of extra code required to write good errors seems repetitive and overwhelming, the test might work better if table-driven, iterating over a list of inputs and outputs defined in a data structure (Go has excellent support for data structure literals). The work to write a good test and good error messages will then be amortized over many test cases. The standard Go library is full of illustrative examples, such as in [the formatting tests for the `fmt` package](https://go.dev/src/fmt/fmt_test.go).

​	如果编写好的错误所需的额外代码量看起来是重复的和压倒性的，且测试是由表驱动的，则在数据结构中定义的输入和输出的列表上迭代（Go对数据结构体字面量有很好的支持），可能效果更好。编写一个好的测试和好的错误信息的工作将被分摊到许多测试案例中。标准 Go 库中有很多说明性的例子，例如 [fmt 包的格式化测试](https://go.dev/src/fmt/fmt_test.go)。

### Why isn't *X* in the standard library? 为什么X不在标准库中？

The standard library's purpose is to support the runtime, connect to the operating system, and provide key functionality that many Go programs require, such as formatted I/O and networking. It also contains elements important for web programming, including cryptography and support for standards like HTTP, JSON, and XML.

​	标准库的目的是支持运行时，连接到操作系统，并提供许多 Go 程序需要的关键功能，如格式化 I/O 和网络。它还包含对网络编程很重要的元素，包括密码学和对HTTP、JSON和XML等标准的支持。

There is no clear criterion that defines what is included because for a long time, this was the *only* Go library. There are criteria that define what gets added today, however.

​	没有明确的标准来定义所包含的内容，因为在很长一段时间里，这是唯一的Go库。然而，有一些标准定义了今天被添加的内容。

New additions to the standard library are rare and the bar for inclusion is high. Code included in the standard library bears a large ongoing maintenance cost (often borne by those other than the original author), is subject to the [Go 1 compatibility promise](https://go.dev/doc/go1compat.html) (blocking fixes to any flaws in the API), and is subject to the Go [release schedule](https://go.dev/s/releasesched), preventing bug fixes from being available to users quickly.

​	新加入标准库的情况很少，而且加入标准库的标准很高。包含在标准库中的代码要承担大量的持续维护费用（通常由原作者以外的人承担），要遵守[Go 1的兼容性承诺](../Go1AndTheFutureOfGoPrograms)（阻止对API中任何缺陷的修复），并且要遵守Go的[发布时间表](../GoReleaseCycle)，阻止用户快速获得 bug 修复。

Most new code should live outside of the standard library and be accessible via the [`go` tool](https://go.dev/cmd/go/)'s `go get` command. Such code can have its own maintainers, release cycle, and compatibility guarantees. Users can find packages and read their documentation at [https://pkg.go.dev/](https://pkg.go.dev/).

​	大多数新的代码应该生活在标准库之外，可以通过[go  tool](https://go.dev/cmd/go/)的`go get`命令进行访问。这样的代码可以有自己的维护者、发布周期和兼容性保证。用户可以在[https://pkg.go.dev/](https://pkg.go.dev/)上找到软件包并阅读其文档。

Although there are pieces in the standard library that don't really belong, such as `log/syslog`, we continue to maintain everything in the library because of the Go 1 compatibility promise. But we encourage most new code to live elsewhere.

​	尽管标准库中有些部分并不真正属于x，比如`log/syslog`，但由于Go 1的兼容性承诺，我们继续维护库中的一切。但我们鼓励大多数新代码在其他地方使用。

## Implementation 实现

### What compiler technology is used to build the compilers? 使用什么编译器技术来构建编译器？

There are several production compilers for Go, and a number of others in development for various platforms.

​	有几个Go的生产型编译器，还有一些正在为不同平台开发的编译器。

The default compiler, `gc`, is included with the Go distribution as part of the support for the `go` command. `Gc` was originally written in C because of the difficulties of bootstrapping—you'd need a Go compiler to set up a Go environment. But things have advanced and since the Go 1.5 release the compiler has been a Go program. The compiler was converted from C to Go using automatic translation tools, as described in this [design document](https://go.dev/s/go13compiler) and [talk](https://go.dev/talks/2015/gogo.slide#1). Thus the compiler is now "self-hosting", which means we needed to face the bootstrapping problem. The solution is to have a working Go installation already in place, just as one normally has with a working C installation. The story of how to bring up a new Go environment from source is described [here](https://go.dev/s/go15bootstrap) and [here](https://go.dev/doc/install/source).

​	默认的编译器 `gc` 包含在 Go 发行版中，是对 `go` 命令支持的一部分。`Gc`最初是用C语言编写的，因为启动困难——您需要一个Go编译器来建立一个Go环境。但现在事情有了进展，`从Go 1.5版本开始，编译器就是一个Go程序`。编译器是使用自动翻译工具从C语言转换为Go语言的，如这个[设计文档](../Other/Go1_3PlusCompilerOverhaul)和[讲座](https://go.dev/talks/2015/gogo.slide#1)中所述。因此，编译器现在是 "`自举（self-hosting）` "的，这意味着我们需要面对启动（bootstrapping ）的问题。解决的办法是已经有了一个正常工作的Go的安装，就像通常有了一个正常工作的C的安装一样。关于如何从源码建立一个新的Go环境的故事在[这里](../Other/Go1_5BootstrapPlan)和[这里](../GettingStarted/InstallingGoFromSource)有描述。

`Gc` is written in Go with a recursive descent parser and uses a custom loader, also written in Go but based on the Plan 9 loader, to generate ELF/Mach-O/PE binaries.

​	`Gc`是用Go写的，有一个递归解析器，并使用一个定制的加载器，也是用Go写的，但基于Plan 9加载器，用来生成ELF/Mach-O/PE二进制文件。

At the beginning of the project we considered using LLVM for `gc` but decided it was too large and slow to meet our performance goals. More important in retrospect, starting with LLVM would have made it harder to introduce some of the ABI and related changes, such as stack management, that Go requires but are not part of the standard C setup. A new [LLVM implementation](https://go.googlesource.com/gollvm/) is starting to come together now, however.

​	在项目开始时，我们考虑过使用LLVM的`gc`，但认为它太大太慢，无法满足我们的性能目标。更重要的是，现在回想起来，使用LLVM会使我们更难引入一些ABI和相关的变化，例如堆栈管理，Go需要这些变化，但这些变化并不是标准C设置的一部分。然而，一个新的[LLVM实现](../Other/gollvm)现在已经开始了。

The `Gccgo` compiler is a front end written in C++ with a recursive descent parser coupled to the standard GCC back end.

​	`Gccgo`编译器是一个用C++编写的前端，带有一个与标准GCC后端耦合的递归解析器。

Go turned out to be a fine language in which to implement a Go compiler, although that was not its original goal. Not being self-hosting from the beginning allowed Go's design to concentrate on its original use case, which was networked servers. Had we decided Go should compile itself early on, we might have ended up with a language targeted more for compiler construction, which is a worthy goal but not the one we had initially.

​	Go被证明是实现Go编译器的一种很好的语言，尽管这并不是它最初的目标。从一开始就不是`自举（self-hosting）`的，这使得Go的设计能够集中在它最初的使用情况上，也就是网络服务器。如果我们一开始就决定Go应该自我编译，那么我们最终可能会得到一个更针对编译器构建的语言，这是一个值得追求的目标，但不是我们最初的目标。

Although `gc` does not use them (yet?), a native lexer and parser are available in the [`go`](https://go.dev/pkg/go/) package and there is also a native [type checker](https://go.dev/pkg/go/types).

​	虽然`gc`没有使用它们（还没有？），但在[go](https://go.dev/pkg/go/)包中有一个本地的 lexer 和解析器，还有一个本地的[类型检查器](https://go.dev/pkg/go/types)。

### How is the run-time support implemented? 运行时支持是如何实现的？

Again due to bootstrapping issues, the run-time code was originally written mostly in C (with a tiny bit of assembler) but it has since been translated to Go (except for some assembler bits). `Gccgo`'s run-time support uses `glibc`. The `gccgo` compiler implements goroutines using a technique called segmented stacks, supported by recent modifications to the gold linker. `Gollvm` similarly is built on the corresponding LLVM infrastructure.

​	同样由于引导（bootstrapping ）问题，运行时代码最初主要是用C语言编写的（有一小部分汇编程序），但后来被翻译成Go语言（除了一些汇编程序部分）。`Gccgo`的运行时支持使用`glibc`。`gccgo`编译器使用一种叫做`分段堆栈（segmented stacks）`的技术来实现goroutines，这种技术得到了最近对 gold 链接器的修改的支持。`Gollvm`同样是建立在相应的LLVM基础设施上。

### Why is my trivial program such a large binary? 为什么我的微不足道的程序会有这么大的二进制文件？

The linker in the `gc` toolchain creates statically-linked binaries by default. All Go binaries therefore include the Go runtime, along with the run-time type information necessary to support dynamic type checks, reflection, and even panic-time stack traces.

​	`gc`工具链中的链接器默认会创建静态链接的二进制文件。因此，所有 Go 二进制文件都包括 Go 运行时，以及支持动态类型检查、反射、甚至恐慌时栈跟踪所需的运行时类型信息。

A simple C "hello, world" program compiled and linked statically using gcc on Linux is around 750 kB, including an implementation of `printf`. An equivalent Go program using `fmt.Printf` weighs a couple of megabytes, but that includes more powerful run-time support and type and debugging information.

​	一个简单的C语言 "hello, world "程序在Linux上使用gcc静态编译和链接，大约是750 kB，包括一个`printf`的实现。一个使用`fmt.Printf`的同等Go程序重达几兆字节，但这包括更强大的运行时支持以及类型和调试信息。

A Go program compiled with `gc` can be linked with the `-ldflags=-w` flag to disable DWARF generation, removing debugging information from the binary but with no other loss of functionality. This can reduce the binary size substantially.

​	用 `gc` 编译的 Go 程序可以用 `-ldflags=-w` 标志链接，以禁止生成 `DWARF`，从二进制文件中删除调试信息，但没有其他功能损失。这可以大大减少二进制文件的大小。

### Can I stop these complaints about my unused variable/import? 我可以停止这些关于我的未使用变量/导入的抱怨吗？

The presence of an unused variable may indicate a bug, while unused imports just slow down compilation, an effect that can become substantial as a program accumulates code and programmers over time. For these reasons, Go refuses to compile programs with unused variables or imports, trading short-term convenience for long-term build speed and program clarity.

​	未使用的变量的存在可能表明了一个错误，而未使用的导入只是减缓了编译速度，随着时间的推移，程序的代码和程序员的积累，这种影响会变得很大。由于这些原因，Go拒绝编译带有未使用的变量或导入的程序，以短期的方便换取长期的编译速度和程序的清晰。

Still, when developing code, it's common to create these situations temporarily and it can be annoying to have to edit them out before the program will compile.

​	尽管如此，在开发代码时，临时创建这些情况是很常见的，在程序编译前必须将其编辑出来，这可能是很烦人的。

Some have asked for a compiler option to turn those checks off or at least reduce them to warnings. Such an option has not been added, though, because compiler options should not affect the semantics of the language and because the Go compiler does not report warnings, only errors that prevent compilation.

​	有些人要求提供一个编译器选项，以关闭这些检查，或至少将其减少为警告。不过这样的选项还没有被添加，因为编译器选项不应该影响语言的语义，而且Go编译器不报告警告，只报告妨碍编译的错误。

There are two reasons for having no warnings. First, if it's worth complaining about, it's worth fixing in the code. (And if it's not worth fixing, it's not worth mentioning.) Second, having the compiler generate warnings encourages the implementation to warn about weak cases that can make compilation noisy, masking real errors that *should* be fixed.

​	没有警告的原因有两个。首先，如果它值得抱怨，就值得在代码中加以修正。（如果不值得修正，也就不值得一提了。）第二，让编译器产生警告会鼓励实现者对那些会使编译变得混乱的弱情况发出警告，从而掩盖应该被修复的实际错误。

It's easy to address the situation, though. Use the blank identifier to let unused things persist while you're developing.

​	不过要解决这种情况很容易。在开发过程中，使用空白标识符可以保留未使用的内容。

```go linenums="1"
import "unused"

// This declaration marks the import as used by referencing an
// item from the package.
var _ = unused.Item  // TODO: Delete before committing!

func main() {
    debugData := debug.Profile()
    _ = debugData // Used only during debugging.
    ....
}
```

Nowadays, most Go programmers use a tool, [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports), which automatically rewrites a Go source file to have the correct imports, eliminating the unused imports issue in practice. This program is easily connected to most editors to run automatically when a Go source file is written.

​	现在，大多数Go程序员都使用[goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)工具，它会自动重写 Go 源文件以获得正确的导入，从而在实践中消除了未使用的导入问题。这个程序很容易连接到大多数编辑器，以便在编写 Go 源文件时自动运行。。

### Why does my virus-scanning software think my Go distribution or compiled binary is infected? 为什么我的病毒扫描软件认为我的 Go 发行版或编译的二进制文件被感染了？

This is a common occurrence, especially on Windows machines, and is almost always a false positive. Commercial virus scanning programs are often confused by the structure of Go binaries, which they don't see as often as those compiled from other languages.

​	这种情况很常见，尤其是在Windows机器上，而且几乎都是假阳性。商业病毒扫描程序常常被Go二进制文件的结构所迷惑，它们不象其他语言编译的文件那样经常看到这种结构。

If you've just installed the Go distribution and the system reports it is infected, that's certainly a mistake. To be really thorough, you can verify the download by comparing the checksum with those on the [downloads page](https://go.dev/dl/).

​	如果您刚刚安装了Go发行版，而系统报告说它被感染了，这肯定是个错误。为了真正彻底，您可以通过将校验和与[下载页面](https://go.dev/dl/)上的校验和进行比较来验证下载。

In any case, if you believe the report is in error, please report a bug to the supplier of your virus scanner. Maybe in time virus scanners can learn to understand Go programs.

​	在任何情况下，如果您认为报告是错误的，请向您的病毒扫描器的供应商报告一个错误。也许随着时间的推移，病毒扫描器可以学会理解Go程序。

## Performance 性能表现

### Why does Go perform badly on benchmark X? 为什么Go在基准X上的表现很差？

One of Go's design goals is to approach the performance of C for comparable programs, yet on some benchmarks it does quite poorly, including several in [golang.org/x/exp/shootout](https://go.googlesource.com/exp/+/master/shootout/). The slowest depend on libraries for which versions of comparable performance are not available in Go. For instance, [pidigits.go](https://go.googlesource.com/exp/+/master/shootout/pidigits.go) depends on a multi-precision math package, and the C versions, unlike Go's, use [GMP](https://gmplib.org/) (which is written in optimized assembler). Benchmarks that depend on regular expressions ([regex-dna.go](https://go.googlesource.com/exp/+/master/shootout/regex-dna.go), for instance) are essentially comparing Go's native [regexp package](https://go.dev/pkg/regexp) to mature, highly optimized regular expression libraries like PCRE.

​	Go 的设计目标之一是使同类程序的性能接近 C，但在一些基准测试中，Go 的表现相当差，包括 [golang.org/x/exp/shootout](https://go.googlesource.com/exp/+/master/shootout/) 中的几个基准。最慢的是依赖于Go中没有可比性能版本的库。例如，[pidigits.go](https://go.googlesource.com/exp/+/master/shootout/pidigits.go)依赖于一个多精度数学包，而C版本与Go不同，（C版本它）使用[GMP](https://gmplib.org/)（用优化的汇编程序编写）。依赖于正则表达式的基准测试（例如[regex-dna.go](https://go.googlesource.com/exp/+/master/shootout/regex-dna.go)）基本上是将Go的本地[regexp package](https://go.dev/pkg/regexp)与成熟的、高度优化的正则表达式库（如PCRE）进行比较。

Benchmark games are won by extensive tuning and the Go versions of most of the benchmarks need attention. If you measure comparable C and Go programs ([reverse-complement.go](https://go.googlesource.com/exp/+/master/shootout/reverse-complement.go) is one example), you'll see the two languages are much closer in raw performance than this suite would indicate.

​	基准测试游戏是通过广泛的调整来赢得的，大多数基准测试的Go版本需要注意。如果您测量可比较的C和Go程序（[reverse-complement.go](https://go.googlesource.com/exp/+/master/shootout/reverse-complement.go)就是一个例子），您会发现这两种语言的原始性能比这个套件所显示的要接近得多。

Still, there is room for improvement. The compilers are good but could be better, many libraries need major performance work, and the garbage collector isn't fast enough yet. (Even if it were, taking care not to generate unnecessary garbage can have a huge effect.)

​	尽管如此，仍有改进的余地。编译器很好，但可以做得更好，许多库需要进行大量的性能工作，而且垃圾收集器还不够快。(即使它够快，小心避免产生不必要的垃圾也会产生巨大的影响。)

In any case, Go can often be very competitive. There has been significant improvement in the performance of many programs as the language and tools have developed. See the blog post about [profiling Go programs](https://blog.golang.org/2011/06/profiling-go-programs.html) for an informative example.

​	在任何情况下，Go通常是非常有竞争力的。随着语言和工具的发展，许多程序的性能都有了明显的改善。请参阅有关 [profiling Go programs]({{< ref "/goBlog/2011/ProfilingGoPrograms" >}}) 的博文，了解一个信息丰富的例子。

## Changes from C - 与 C 的变化

### Why is the syntax so different from C? 为什么语法与C语言如此不同？

Other than declaration syntax, the differences are not major and stem from two desires. First, the syntax should feel light, without too many mandatory keywords, repetition, or arcana. Second, the language has been designed to be easy to analyze and can be parsed without a symbol table. This makes it much easier to build tools such as debuggers, dependency analyzers, automated documentation extractors, IDE plug-ins, and so on. C and its descendants are notoriously difficult in this regard.

​	除了声明语法之外，其他的差别并不大，而且源于两个愿望。首先，语法应该给人以轻松的感觉，没有太多的强制性关键字、重复或奥秘。第二，该语言被设计成易于分析，不需要符号表就可以进行解析。这使得它更容易建立工具，如调试器、依赖分析器、自动文档提取器、IDE插件等等。C及其后代在这方面是出了名的困难。

### Why are declarations backwards? 为什么声明是倒退的？

They're only backwards if you're used to C. In C, the notion is that a variable is declared like an expression denoting its type, which is a nice idea, but the type and expression grammars don't mix very well and the results can be confusing; consider function pointers. Go mostly separates expression and type syntax and that simplifies things (using prefix `*` for pointers is an exception that proves the rule). In C, the declaration

​	如果您习惯于C语言，它们就是相反的。在C语言中，变量的声明就像表示其类型的表达式一样，这是一个很好的想法，但是类型语法和表达式语法不能很好地混合，其结果可能是混乱的；考虑到函数指针。Go大多将表达式和类型语法分开，这就简化了事情（对指针使用前缀`*`是证明这一规则的一个例外）。在C语言中，声明

```go linenums="1"
    int* a, b;
```

declares `a` to be a pointer but not `b`; in Go

声明了`a`是一个指针，但没有声明`b`；在Go中

```go linenums="1"
    var a, b *int
```

declares both to be pointers. This is clearer and more regular. Also, the `:=` short declaration form argues that a full variable declaration should present the same order as `:=` so

声明两者都是指针。这样更清晰，更有规律。另外，`:=`的短声明形式认为，完整的变量声明应该呈现与`:=`相同的顺序，所以

```go linenums="1"
    var a uint64 = 1
```

has the same effect as

和

```go linenums="1"
    a := uint64(1)
```

有同样的效果。

Parsing is also simplified by having a distinct grammar for types that is not just the expression grammar; keywords such as `func` and `chan` keep things clear.

​	通过为类型制定独特的语法，而不仅仅是表达式语法，解析工作也得到了简化；`func`和`chan`等关键字可以让事情变得清晰。

See the article about [Go's Declaration Syntax](https://go.dev/doc/articles/gos_declaration_syntax.html) for more details.

​	更多细节请参见[Go's Declaration Syntax]({{< ref "/goBlog/2010/GosDeclarationSyntax" >}})一文。

### Why is there no pointer arithmetic? 为什么没有指针算术？

Safety. Without pointer arithmetic it's possible to create a language that can never derive an illegal address that succeeds incorrectly. Compiler and hardware technology have advanced to the point where a loop using array indices can be as efficient as a loop using pointer arithmetic. Also, the lack of pointer arithmetic can simplify the implementation of the garbage collector.

​	安全性。如果没有指针算术，就有可能创造出一种语言，永远无法推导出一个非法的地址，从而错误地成功。编译器和硬件技术已经发展到这样的程度：使用数组索引的循环可以和使用指针算术的循环一样有效。而且，`没有指针算术可以简化垃圾收集器的实现`。

### Why are `++` and `--` statements and not expressions? And why postfix, not prefix? 为什么是++和--语句而不是表达式？以及为什么是后缀而不是前缀？

Without pointer arithmetic, the convenience value of pre- and postfix increment operators drops. By removing them from the expression hierarchy altogether, expression syntax is simplified and the messy issues around order of evaluation of `++` and `--` (consider `f(i++)` and `p[i] = q[++i]`) are eliminated as well. The simplification is significant. As for postfix vs. prefix, either would work fine but the postfix version is more traditional; insistence on prefix arose with the STL, a library for a language whose name contains, ironically, a postfix increment.

​	没有指针运算，前缀和后缀增量运算符的方便性就会下降。通过将它们从表达式层次中完全移除，表达式的语法被简化了，围绕`++`和`--`的求值顺序的混乱问题（考虑`f(i++)`和`p[i] = q[++i]`）也被消除了。这种简化是非常重要的。至于后缀和前缀，两者都可以使用，但后缀版本更为传统；坚持使用前缀是由STL引起的，STL是一种语言的库，讽刺的是，其名称包含后缀增量。

### Why are there braces but no semicolons? And why can't I put the opening brace on the next line? 为什么有大括号而没有分号？为什么我不能把开头的大括号放在下一行？

Go uses brace brackets for statement grouping, a syntax familiar to programmers who have worked with any language in the C family. Semicolons, however, are for parsers, not for people, and we wanted to eliminate them as much as possible. To achieve this goal, Go borrows a trick from BCPL: the semicolons that separate statements are in the formal grammar but are injected automatically, without lookahead, by the lexer at the end of any line that could be the end of a statement. This works very well in practice but has the effect that it forces a brace style. For instance, the opening brace of a function cannot appear on a line by itself.

​	Go使用花括号对语句进行分组，这是使用过C语言的程序员都熟悉的语法。然而，分号是给解析器用的，不是给人用的，我们想尽可能地消除它们。为了实现这一目标，Go从`BCPL`中借鉴了一个技巧：分隔语句的分号在形式语法中存在，但在任何可能是语句结尾的行的末尾，由词法分析器自动注入，without lookahead。这在实践中效果很好，但有一个影响，那就是它强制采用了花括号样式。例如，一个函数的开头花括号不能单独出现在一行中。

Some have argued that the lexer should do lookahead to permit the brace to live on the next line. We disagree. Since Go code is meant to be formatted automatically by [`gofmt`](https://go.dev/cmd/gofmt/), *some* style must be chosen. That style may differ from what you've used in C or Java, but Go is a different language and `gofmt`'s style is as good as any other. More important—much more important—the advantages of a single, programmatically mandated format for all Go programs greatly outweigh any perceived disadvantages of the particular style. Note too that Go's style means that an interactive implementation of Go can use the standard syntax one line at a time without special rules.

​	有些人认为，lexer 应该进行查找，以允许括号出现在下一行。我们不同意这个观点。由于Go代码是由[gofmt]({{< ref "/cmd/gofmt">}})自动格式化的，所以必须选择一些风格。这种风格可能与您在C或Java中使用的不同，但Go是一种不同的语言，`gofmt`的风格与其他语言一样好。更重要的是——要重要得多的是——为所有Go程序提供单一的、程序化的强制格式的优势大大超过了特定风格的任何感知的缺点。还要注意的是，Go的风格意味着Go的交互式实现可以一行一行地使用标准语法，而不需要特别的规则。

### Why do garbage collection? Won't it be too expensive? 为什么要做垃圾收集？不会代价太大吗？

One of the biggest sources of bookkeeping in systems programs is managing the lifetimes of allocated objects. In languages such as C in which it is done manually, it can consume a significant amount of programmer time and is often the cause of pernicious bugs. Even in languages like C++ or Rust that provide mechanisms to assist, those mechanisms can have a significant effect on the design of the software, often adding programming overhead of its own. We felt it was critical to eliminate such programmer overheads, and advances in garbage collection technology in the last few years gave us confidence that it could be implemented cheaply enough, and with low enough latency, that it could be a viable approach for networked systems.

​	系统程序中最大的记账来源之一是管理分配对象的生命周期。在诸如C语言中，它是手动完成的，它可能会消耗程序员大量的时间，而且常常是致命的错误的原因。即使在像C++或Rust这样提供辅助机制的语言中，这些机制也会对软件的设计产生重大影响，往往会增加自身的编程开销。我们觉得消除这种程序员的开销是至关重要的，而过去几年垃圾收集技术的进步给了我们信心，它可以用足够低的成本地实现，并具有足够低的延迟，它可以成为网络系统的一个可行的方法。

Much of the difficulty of concurrent programming has its roots in the object lifetime problem: as objects get passed among threads it becomes cumbersome to guarantee they become freed safely. Automatic garbage collection makes concurrent code far easier to write. Of course, implementing garbage collection in a concurrent environment is itself a challenge, but meeting it once rather than in every program helps everyone.

​	并发编程的大部分困难都源于对象的生存期问题：当对象在线程之间传递时，保证它们安全释放就变得很麻烦。自动垃圾收集使并发代码更容易编写。当然，在并发环境中实现垃圾收集本身就是一个挑战，但一次就实现胜过在每个程序中都去（手动）实现，对每个人都有帮助。

Finally, concurrency aside, garbage collection makes interfaces simpler because they don't need to specify how memory is managed across them.

​	最后，撇开并发性不谈，垃圾收集使接口更简单，因为它们不需要指定如何跨接口管理内存。

This is not to say that the recent work in languages like Rust that bring new ideas to the problem of managing resources is misguided; we encourage this work and are excited to see how it evolves. But Go takes a more traditional approach by addressing object lifetimes through garbage collection, and garbage collection alone.

​	这并不是说最近在Rust等语言中为管理资源的问题带来新思路的工作是错误的；我们鼓励这项工作，并很高兴看到它的发展情况。但是Go采取了一种更传统的方法，即`通过垃圾收集来解决对象的生命周期问题，而且仅仅只是垃圾收集`。

The current implementation is a mark-and-sweep collector. If the machine is a multiprocessor, the collector runs on a separate CPU core in parallel with the main program. Major work on the collector in recent years has reduced pause times often to the sub-millisecond range, even for large heaps, all but eliminating one of the major objections to garbage collection in networked servers. Work continues to refine the algorithm, reduce overhead and latency further, and to explore new approaches. The 2018 [ISMM keynote](https://blog.golang.org/ismmkeynote) by Rick Hudson of the Go team describes the progress so far and suggests some future approaches.

​	目前的实现是一个标记-清除（mark-and-sweep）收集器。如果机器是多处理器，收集器会在一个单独的CPU核心上与主程序并行运行。近年来关于收集器的主要工作已经将停顿时间减少到了亚毫秒范围，甚至对于大堆也是如此，这几乎消除了网络服务器中垃圾收集的主要障碍之一。完善算法的工作仍在继续，进一步减少开销和延迟，并探索新的方法。Go团队的`Rick Hudson`在2018年[ISMM keynote]({{< ref "/goBlog/2018/GettingToGoTheJourneyOfGosGarbageCollector" >}})的主题演讲中描述了迄今为止的进展，并提出了一些未来的方法。

On the topic of performance, keep in mind that Go gives the programmer considerable control over memory layout and allocation, much more than is typical in garbage-collected languages. A careful programmer can reduce the garbage collection overhead dramatically by using the language well; see the article about [profiling Go programs](https://blog.golang.org/2011/06/profiling-go-programs.html) for a worked example, including a demonstration of Go's profiling tools.

​	关于性能的话题，请记住，Go给了程序员对内存布局和分配相当大的控制权，比典型的垃圾收集型语言要多得多。细心的程序员可以通过很好地使用该语言来大幅减少垃圾收集的开销；请参阅为一个工作示例[剖析Go程序]({{< ref "/goBlog/2011/ProfilingGoPrograms" >}})的文章，其中包括Go的剖析工具的演示。