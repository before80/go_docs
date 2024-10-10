+++
title = "FAQ"
weight = 100
date = 2023-05-18T16:56:23+08:00
description = ""
isCJKLanguage = true
draft = false

+++
# Frequently Asked Questions (FAQ)

> 原文：[https://go.dev/doc/faq](https://go.dev/doc/faq)

## Origins 起源

### What is the purpose of the project? 这个项目的目的是什么？

​	在Go诞生的时候，也就是十年前，编程世界与今天不同。生产软件通常是用C++或Java编写的，GitHub还不存在，大多数计算机还不是多处理器，除了Visual Studio和Eclipse之外，几乎没有什么IDE或其他高级工具可用，更不用说在互联网上免费使用了。

​	与此同时，我们对使用我们所使用的语言来开发服务器软件所需要的过度复杂性感到沮丧。自从C、C++和Java等语言被首次开发以来，计算机已经变得非常快了，但是编程行为本身并没有得到同样的发展。另外，很明显，多处理器正在普及，但大多数语言对它们的有效和安全编程没有提供什么帮助。

​	我们决定退一步思考，随着技术的发展，哪些主要问题将在未来几年主导软件工程，以及一种新的语言如何帮助解决这些问题。例如，多核CPU的兴起表明，一种语言应该为某种并发性或并行性提供一流的支持。为了使大型并发程序中的资源管理变得可行，需要有垃圾收集，或者至少是某种安全的自动内存管理。

​	这些考虑导致了[一系列的讨论](https://commandcenter.blogspot.com/2017/09/go-ten-years-and-climbing.html)，Go就是在这些讨论中产生的，首先是一系列的想法和愿望，然后是作为一种语言。一个首要的目标是，Go要更多地帮助工作中的程序员，使工具化，使代码格式化等日常任务自动化，并消除在大型代码库中工作的障碍。

​	关于Go的目标以及如何实现或至少如何接近这些目标的更详尽描述，可以在[Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)中找到。

### What is the history of the project? 这个项目的历史是怎样的？

​	2007年9月21日，Robert Griesemer、Rob Pike和Ken Thompson开始在白板上勾画新语言的目标。在几天之内，这些目标就变成了一个做事的计划和一个关于它将是什么的合理想法。设计工作继续与无关的工作并行。到2008年1月，Ken已经开始了编译器的工作，用它来探索想法；它的输出是C代码。到了年中，该语言已经成为一个全职项目，并且已经稳定下来，可以尝试生产一个编译器。2008年5月，Ian Taylor利用规范草案独立开始了Go的GCC前端工作。Russ Cox在2008年底加入，并帮助将语言和库从原型变为现实。

​	Go在2009年11月10日成为一个公开的开源项目。社区中无数人贡献了想法、讨论和代码。

​	现在全世界有数以百万计的Go程序员——gophers，而且每天都有更多。Go的成功已经远远超过了我们的预期。

### What's the origin of the gopher mascot? 囊地鼠吉祥物的起源是什么？

​	吉祥物和标志是由[Renée French](https://reneefrench.blogspot.com/)设计的，她还设计了Plan 9的兔子[Glenda](https://9p.io/plan9/glenda.html)。一篇[关于囊地鼠的博客]({{< ref "/goBlog/2014/TheGoGopher" >}})文章解释了它是如何从她几年前用于[WFMU](https://wfmu.org/)T恤设计的一个囊地鼠衍生出来的。该标志和吉祥物属于[Creative Commons Attribution 4.0](https://creativecommons.org/licenses/by/4.0/)许可范围。

​	囊地鼠有一个[模型表](https://go.dev/doc/gopher/modelsheet.jpg)，说明他的特点和如何正确地表现它们。该模型表在Renée于2016年在Gophercon的[演讲](https://www.youtube.com/watch?v=4rw_B4yY69k)中首次展示。他有独特的特征；他是Go的囊地鼠，不是普通的囊地鼠。

### Is the language called Go or Golang? 这门语言叫Go还是Golang？

​	这门语言叫Go。出现 "golang "这个名字是因为网站最初是golang.org。(那时还没有.dev域名。)不过很多人都使用golang这个名字，而且它作为一个标签很方便。例如，该语言的Twitter标签是 "#golang"。但无论如何，该语言的名字只是普通的Go。

​	题外话：虽然[官方标志]({{< ref "/goBlog/2018/GosNewBrand" >}})有两个大写字母，但语言的名字是Go，而不是GO。

### Why did you create a new language? 您为什么要创造一种新的语言？

​	Go的诞生是出于对现有语言和环境的失望，我们在谷歌所做的工作。编程已经变得太困难了，语言的选择也是部分原因。人们不得不选择高效的编译、高效的执行，或者方便的编程；这三者在同一种主流语言中都不存在。有能力的程序员选择了轻松而不是安全和效率，他们转向动态类型的语言，如Python和JavaScript，而不是C++，或者在较小程度上使用Java。

​	我们并不孤单。在经历了多年的编程语言领域的平静之后，Go是几种新语言中的第一批——Rust、Elixir、Swift等等——它们使编程语言的开发再次成为一个活跃的、几乎是主流的领域。

​	Go解决了这些问题，它试图将解释型、动态类型语言的编程便利性与静态类型、编译型语言的效率和安全性相结合。它还旨在成为现代语言，支持网络和多核计算。最后，使用Go的目的是快速：在一台计算机上构建一个大的可执行文件最多只需要几秒钟。为了实现这些目标，需要解决一些语言问题：一个有表现力但轻量级的类型系统；并发性和垃圾收集；严格的依赖项规范等等。这些问题不能由库或工具来解决；需要一种新的语言。

​	[Go at Google](https://go.dev/talks/2012/splash.article)的文章讨论了Go语言设计的背景和动机，并对本FAQ中的许多答案提供了更多细节。

### What are Go's ancestors? Go 的祖先是什么？

​	Go主要属于C语言家族（基本语法），有来自Pascal/Modula/Oberon家族（声明、包）的重要输入，还有来自Tony Hoare的CSP语言的一些想法，如Newsqueak和Limbo（并发）。然而，它是一种全面的新语言。在每一个方面，这门语言的设计都是通过思考程序员的工作以及如何使编程，至少是我们所做的那种编程，更加有效，也就是更加有趣。

### What are the guiding principles in the design? 设计中的指导原则是什么？

​	在设计Go的时候，Java和C++是编写服务器最常用的语言，至少在Google是这样。我们觉得这些语言需要太多的记账和重复。一些程序员的反应是转向像Python这样更动态、更流畅的语言，但却牺牲了效率和类型安全。我们认为应该可以在一种语言中实现效率、安全和流畅性。

​	Go试图在这两种意义上减少输入的数量。在整个设计过程中，我们试图减少混乱和复杂性。没有前向声明，也没有头文件；所有东西都只声明一次。初始化是明确的、自动的，并且易于使用。语法简洁，关键字少。重复（`foo.Foo* myFoo = new(foo.Foo)`）通过使用 `:=` declare-and-initialize 结构进行简单的类型推导来减少。也许最根本的是，没有类型的层次结构：类型就是这样，它们不需要宣布它们的关系。这些简化使得Go在不牺牲复杂度的前提下实现了表现力和可理解性。

​	另一个重要原则是保持概念的正交性。`方法可以为任何类型实现；结构代表数据而接口代表抽象；`等等。正交性使我们更容易理解事物结合时的情况。

## Usage 使用方法

### Is Google using Go internally? Google在内部使用Go吗？

​	是的。Go 在 Google 内部的生产中被广泛使用。一个简单的例子是[golang.org](https://golang.org/)背后的服务器。它只是在 [Google App Engine](https://developers.google.com/appengine/) 上以生产配置运行的 [godoc](https://go.dev/cmd/godoc) 文档服务器。

​	一个更重要的例子是Google的下载服务器`dl.google.com`，它提供Chrome二进制文件和其他大型安装程序，如`apt-get`包。

​	Go不是谷歌唯一使用的语言，远非如此，但它是包括[网站可靠性工程（SRE）](https://go.dev/talks/2013/go-sreops.slide)和大规模数据处理在内的许多领域的关键语言。

### What other companies use Go? 还有哪些公司使用Go？

​	Go的使用在全球范围内不断增长，尤其是在云计算领域，但绝非仅限于此。`Docker` 和 `Kubernetes` 是用 Go 编写的几个主要的云计算基础设施项目，但还有很多。

​	不过，这不仅仅是云计算。Go Wiki有一个[页面](https://github.com/golang/go/wiki/GoUsers)，定期更新，其中列出了一些使用Go的公司。

​	Wiki还有一个页面，链接到使用该语言的公司和项目的[成功案例](https://github.com/golang/go/wiki/SuccessStories)。

### Do Go programs link with C/C++ programs? Go程序可以与C/C++程序链接吗？

​	可以在同一地址空间中同时使用 C 和 Go，但这并不是一种自然的配合，可能需要特殊的接口软件。另外，将 C 与 Go 代码连接起来会放弃 Go 提供的内存安全和堆栈管理特性。有时绝对有必要使用C语言库来解决问题，但这样做总是会引入纯Go代码所不具备的风险因素，所以要谨慎行事。

​	如果您确实需要在Go中使用C语言，如何进行取决于Go编译器的实现。Go 团队支持三种 Go 编译器实现。它们是 `gc`（默认的编译器）、`gccgo`（使用 GCC 后端）以及不太成熟的 `gollvm`（使用 LLVM 基础架构）。

​	`Gc`使用与C不同的调用约定和链接器，因此不能直接从C程序中调用，反之亦然。[cgo](https://go.dev/cmd/cgo/)程序提供了一个 "外来函数接口 "的机制，允许从Go代码中安全地调用C库。SWIG将这种能力扩展到C++库。

​	您也可以在`Gccgo`和`gollvm`中使用`cgo`和SWIG。由于它们使用的是传统的API，所以在非常小心的情况下，也可以将这些编译器的代码直接与GCC/LLVM编译的C或C++程序链接。然而，要安全地做到这一点，需要了解所有相关语言的调用惯例，以及从 Go 中调用 C 或 C++ 时对堆栈限制的关注。

### What IDEs does Go support? Go 支持哪些集成开发环境？

​	Go 项目不包括定制的集成开发环境，但语言和库的设计使其易于分析源代码。因此，大多数知名的编辑器和集成开发环境都很好地支持 Go，可以直接使用或通过插件使用。

​	对Go有良好支持的知名IDE和编辑器包括Emacs、Vim、VSCode、Atom、Eclipse、Sublime、IntelliJ（通过一个名为Goland的自定义变体），以及更多。您最喜欢的环境有可能是用Go编程的有效环境。

### Does Go support Google's protocol buffers? Go 是否支持 Google 的协议缓冲区？

​	一个单独的开源项目提供了必要的编译器插件和库。它可以在 [github.com/golang/protobuf/](https://github.com/golang/protobuf)上找到。

### Can I translate the Go home page into another language?  我可以将 Go 主页翻译成其他语言吗？

​	当然可以。我们鼓励开发者用自己的语言制作 Go 语言网站。然而，如果您选择在您的网站上添加谷歌的标志或品牌（它不会出现在golang.org上），您需要遵守www.google.com/permissions/guidelines.html 的准则



## Design 设计

### Does Go have a runtime? Go有运行时吗？

​	Go确实有一个广泛的库，称为 `runtime`，它是每个Go程序的一部分。`runtime`库实现了垃圾收集、并发、堆栈管理以及 Go 语言的其他关键功能。虽然它是语言的核心，但Go的`runtime`类似于`libc`，即C语言库。

​	然而，重要的是要理解Go的`runtime`不包括虚拟机，如Java`runtime`所提供的虚拟机。Go程序被提前编译为本地机器代码（或JavaScript或WebAssembly，用于某些变体实现）。因此，虽然这个词经常被用来描述程序运行的虚拟环境，但在 Go 中，"`runtime` "这个词`只是提供关键语言服务的库的名称`。

### What's up with Unicode identifiers? Unicode标识符是怎么回事？

​	在设计Go时，我们希望确保它不会过度以ASCII为中心，这意味着将标识符的空间从7位ASCII的范围内扩展出来。Go的规则——标识符必须是Unicode所定义的字母或数字——很容易理解和实现，`但也有限制`。例如，组合字符在设计上被排除在外，而这也排除了一些语言，如Devanagari。

​	这个规则还有一个不幸的后果。由于导出的标识符必须以大写字母开始，根据定义，由某些语言的字符创建的标识符不能被导出。目前，唯一的解决办法是使用类似`X日本語`的东西，这显然是不能令人满意的。

​	自从该语言的最早版本以来，人们一直在考虑如何最好地扩展标识符空间以适应使用其他母语的程序员。具体怎么做仍然是一个活跃的讨论话题，未来版本的语言可能会在标识符的定义上更加自由。例如，它可能会采用Unicode组织关于标识符的[建议](http://unicode.org/reports/tr31/)中的一些想法。无论发生什么，都必须在保留（或扩大）字母大小写决定标识符的可见性的同时进行兼容，这仍然是我们最喜欢的Go的特点之一。

​	目前，我们有一个简单的规则，以后可以在不破坏程序的情况下进行扩展，这个规则可以避免因允许模棱两可的标识符的规则而产生的错误。

### Why does Go not have feature X? 为什么Go没有X的特性？

​	每种语言都包含新奇的功能，并忽略了某些人最喜欢的功能。Go的设计着眼于编程的便利性、编译的速度、概念的正交性，以及支持并发和垃圾回收等功能的需要。您最喜欢的功能可能因为不合适而缺失，因为它影响了编译速度或设计的清晰度，或者因为它使基本的系统模型过于困难。

​	如果Go缺失了X功能让您感到困扰，请原谅我们，转而研究Go确实拥有的功能。您可能会发现它们以有趣的方式弥补了X的缺失。

### When did Go get generic types? Go 什么时候有了泛型？

​	Go 1.18版本在语言中加入了类型参数。这允许一种多态或泛型编程的形式。详情请参见[语言规范]({{< ref "/langSpec/DeclarationsAndScope#type-definitions">}})和[提案](https://go.dev/design/43651-type-parameters)。

### Why was Go initially released without generic types? 为什么 Go 最初发布时没有泛型？

​	Go 的目的是作为一种编写服务器程序的语言，以便于长期维护。(更多背景见[这篇文章](https://go.dev/talks/2012/splash.article)。)设计集中在可扩展性、可读性和并发性等方面。多态编程在当时看来对该语言的目标并不重要，所以最初为了简化而被排除在外。

​	泛型是很方便的，但它们的代价是类型系统和运行时间的复杂性。我们花了一些时间来开发一个我们认为能够提供与复杂性相称的价值的设计。

### Why does Go not have exceptions?  为什么Go没有异常？

​	我们认为将异常与控制结构相耦合，如`try-catch-finally`习语，会造成代码的复杂化。它还倾向于鼓励程序员将太多的普通错误，如无法打开文件，标记为异常。

​	Go采取了一种不同的方法。对于普通的错误处理，Go的`多值返回`使得报告错误很容易，而不需要重载返回值。[一个典型的错误类型，加上Go的其他特性](https://go.dev/doc/articles/error_handling.html)，使错误处理变得令人愉快，但与其他语言中的错误处理完全不同。

​	Go也有几个内置的函数，用来发出信号并从真正的特殊情况下恢复。恢复机制仅作为错误发生后函数状态被拆解的一部分来执行，这足以处理灾难，但不需要额外的控制结构，如果使用得好，可以产生干净的错误处理代码。

​	详见[Defer, Panic, and Recover]({{< ref "/goBlog/2010/DeferPanicAandRecover" >}})一文。另外，[Errors are values]({{< ref "/goBlog/2015/ErrorsAreValues" >}})的博文描述了一种在 Go 中干净地处理错误的方法，它表明，由于错误只是值，因此可以在错误处理中部署 Go 的全部力量。

### Why does Go not have assertions? 为什么 Go 没有断言？

​	Go 并没有提供断言。不可否认的是，它们很方便，但我们的经验是，程序员把它们当作拐杖，避免考虑正确的错误处理和报告。正确的错误处理意味着服务器可以继续运行，而不是在出现非致命错误后崩溃。正确的错误报告意味着错误是直接的，是有针对性的，使程序员不必解释一个大的崩溃跟踪。当看到错误的程序员不熟悉代码的时候，精确的错误就显得尤为重要。

​	我们理解这是一个争论点。Go语言和库中有许多与现代实践不同的东西，只是因为我们觉得有时值得尝试不同的方法。

### Why build concurrency on the ideas of CSP? 为什么要在CSP的思想上建立并发性？

​	随着时间的推移，并发和多线程编程以其难度而闻名。我们相信这部分是由于复杂的设计，比如[pthreads](https://en.wikipedia.org/wiki/POSIX_Threads)，部分是由于过度强调低级别的细节，比如mutexes，条件变量和内存障碍。更高级别的接口可以使代码简单得多，即使下面仍然有互斥等问题。

​	为并发提供高级语言支持的最成功的模型之一来自Hoare的Communicating Sequential Processes，即CSP。Occam和Erlang是源于CSP的两种著名语言。Go的并发原语来自家族树的另一部分，其主要贡献是将通道作为第一类对象的强大概念。早期几种语言的经验表明，CSP模型很适合程序语言框架。

### Why goroutines instead of threads? 为什么是goroutines而不是线程？

​	goroutines是使并发性易于使用的一部分。这个想法已经存在了一段时间，它是将独立执行的函数——coroutines——复用到一组线程中。当一个 coroutine 阻塞时，例如通过调用一个阻塞的系统调用，run-time 会自动将同一操作系统线程上的其他程序转移到不同的可运行线程，这样它们就不会被阻塞。程序员看不到这些，这就是问题所在。这样的结果，我们称之为goroutines，可以非常便宜：除了堆栈的内存，它们几乎没有开销，而堆栈的内存只有几千字节。

​	为了使堆栈变小，Go的run-time使用可调整大小的有界堆栈。一个新诞生的goroutine被赋予几千字节的内存，这几乎总是足够的。当它不够时，运行时间会自动增加（和缩小）用于存储堆栈的内存，从而使许多goroutine能够在适量的内存中生存。每个函数调用的CPU开销平均约为三条廉价指令。在同一地址空间中创建成百上千个goroutines是很实际的。如果goroutines只是线程，系统资源会在更少的数量上耗尽。

### Why are map operations not defined to be atomic? 为什么映射操作没有被定义为原子操作？

​	经过长时间的讨论，我们决定，映射的典型使用并不要求从多个goroutine中进行安全访问，而在那些需要访问的情况下，映射可能是一些更大的数据结构或计算的一部分，而这些数据结构或计算已经被同步了。因此，要求所有的映射操作都要抓取一个mutex，这将降低大多数程序的速度，并为少数程序增加安全性。然而，这并不是一个简单的决定，因为这意味着不受控制的映射访问会使程序崩溃。

​	该语言并不排除原子映射更新。在需要的时候，比如在托管一个不受信任的程序时，实现可以对映射访问进行联锁。

​	只有在更新发生时，映射访问才是不安全的。只要所有的goroutines只是阅读查找映射中的元素，包括使用for range循环遍历它——而不是通过向元素赋值或进行删除来改变映射，那么它们在没有同步的情况下同时访问映射是安全的。

​	作为对正确使用映射的帮助，一些语言的实现包含一个特殊的检查，当映射被并发执行不安全地修改时，会在运行时自动报告。

### Will you accept my language change? 您会接受我的语言修改吗？

​	人们经常建议对语言进行改进——[邮件列表](https://groups.google.com/group/golang-nuts)中包含了丰富的此类讨论历史——但这些修改很少被接受。

​	虽然 Go 是一个开源项目，但语言和库受到[兼容性承诺](https://go.dev/doc/go1compat.html)的保护，至少在源代码层面上，不会出现破坏现有程序的变化（程序可能需要偶尔重新编译以保持最新）。如果您的建议违反了Go 1的规范，我们甚至不能接受这个想法，不管它的优点是什么。未来的Go的主要版本可能会与Go 1不兼容，但关于这个话题的讨论才刚刚开始，有一点是肯定的：在这个过程中会很少引入这样的不兼容问题。此外，兼容性的承诺鼓励我们为老程序提供一条自动改进的道路，以便在出现这种情况时进行调整。

​	即使您的方案与Go 1规范兼容，它也可能不符合Go的设计目标的精神。文章[Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)一文解释了 Go 的起源及其设计背后的动机。

## Types 类型

### Is Go an object-oriented language? Go 是一种面向对象的语言吗？

​	是，也不是。虽然 Go 有类型和方法，并且允许面向对象的编程风格，但没有类型层次。Go 中的 "interface "概念提供了一种不同的方法，我们认为这种方法易于使用，而且在某些方面更加通用。还有一些方法可以将类型嵌入到其他类型中，以提供类似的东西，但不等同于子类。此外，Go中的方法比C++或Java中的方法更通用：它们可以为任何类型的数据定义，甚至是内置类型，如普通的、"未装箱的 "整数。它们并不局限于结构体（类）。

​	另外，由于缺乏类型层次，Go中的 "对象 "比C++或Java等语言更轻巧。

### How do I get dynamic dispatch of methods? 我如何获得方法的动态调度？

​	拥有动态分配方法的唯一方法是通过接口。结构体或任何其他具体类型的方法总是以静态方式解决。

### Why is there no type inheritance? 为什么没有类型继承？

​	面向对象的编程，至少在最著名的语言中，涉及到太多关于类型之间关系的讨论，这些关系往往可以自动衍生。Go采取了一种不同的方法。

​	在Go中，一个类型不是要求程序员提前声明两个类型的关系，而是自动满足任何指定其方法子集的接口。除了减少记账，这种方法还有真正的优势。类型可以同时满足许多接口，而没有传统的多重继承的复杂问题。接口可以是非常轻量级的——一个只有一个甚至是零个方法的接口可以表达一个有用的概念。如果有新的想法出现，或者为了测试，可以在事后添加接口，而不需要对原始类型进行注释。因为类型和接口之间没有明确的关系，所以不存在需要管理或讨论的类型层次结构。

​	我们可以用这些想法来构建类似于类型安全的 Unix 管道的东西。例如，看看`fmt.Fprintf`是如何实现对任何输出的格式化打印，而不仅仅是文件，或者`bufio`包是如何与文件I/O完全分离的，或者`image`包是如何生成压缩图像文件的。所有这些想法都源于一个接口（`io.Writer`），表示一个方法（`Write`）。而这仅仅是表面现象。Go的接口对程序的结构有深刻的影响。

​	这需要一些时间来适应，但这种隐式的类型依赖风格是Go最富有成效的地方之一。

### Why is `len` a function and not a method? 为什么`len`是一个函数而不是一个方法？

​	我们对这个问题进行了辩论，但最终决定将 len 和朋友们作为函数来实现，在实践中是没有问题的，也不会使基本类型的接口（在 Go 类型意义上）问题复杂化。

### Why does Go not support overloading of methods and operators? 为什么 Go 不支持方法和运算符的重载？

​	如果不需要同时进行类型匹配，方法调度就会被简化。其他语言的经验告诉我们，拥有各种同名但不同签名的方法偶尔是有用的，但在实践中也会让人感到困惑和脆弱。在Go的类型系统中，仅通过名称进行匹配并要求类型的一致性是一个重要的简化决定。

​	关于操作符重载，它似乎更像是一种便利，而不是绝对的要求。同样，没有它，事情会更简单。

### Why doesn't Go have "implements" declarations? 为什么 Go 没有 "implements"声明？

​	Go类型通过实现一个接口的方法来满足该接口，仅此而已。这个属性允许定义和使用接口，而不需要修改现有的代码。它实现了一种结构化的类型，促进了关注点的分离，提高了代码的重复使用，并使之更容易建立在代码发展过程中出现的模式上。接口的语义是Go的灵活、轻量级感觉的主要原因之一。

​	更多细节请参见关于[类型继承的问题](#why-is-there-no-type-inheritance-为什么没有类型继承)。

### How can I guarantee my type satisfies an interface? 如何保证我的类型满足接口的要求？

​	您可以要求编译器通过尝试使用 `T` 的零值或 `T` 的指针进行赋值，来检查类型 `T` 是否实现了接口 `I`，如果合适的话：

```go linenums="1"
type T struct{}
var _ I = T{}       // Verify that T implements I. => 验证 T 实现了 I
var _ I = (*T)(nil) // Verify that *T implements I. => 验证 *T 实现了 I
```

​	如果`T`（或`*T`，相应地）没有实现`I`，这个错误将在编译时被捕获。

​	如果您希望一个接口的用户明确声明他们实现了这个接口，您可以在接口的方法集中添加一个带有描述性名称的方法。比如说：

```go linenums="1"
type Fooer interface {
    Foo()
    ImplementsFooer()
}
```

​	然后一个类型必须实现`ImplementsFooer`方法才能成为`Fooer`，清楚地记录这一事实并在[go doc](https://go.dev/cmd/go/#hdr-Show_documentation_for_package_or_symbol)的输出中公布。

```go linenums="1"
type Bar struct{}
func (b Bar) ImplementsFooer() {}
func (b Bar) Foo() {}
```

​	大多数代码都不使用这种约束，因为它们限制了接口思想的实用性。但有时，它们对于解决类似接口之间的歧义是必要的。

### Why doesn't type T satisfy the Equal interface? 为什么类型T不满足 Equal 接口？

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

​	与某些多态类型系统中的类似情况不同，`T`并没有实现`Equaler`。`T.Equal`的实参类型是`T`，而不是字面上所要求的`Equaler`类型。

​	在Go中，类型系统不提升`Equal`的实参；这是程序员的责任，如类型`T2`所示，它确实实现了`Equaler`：

```go linenums="1"
type T2 int
func (t T2) Equal(u Equaler) bool { return t == u.(T2) }  // satisfies Equaler
```

不过，即使这样也和其他类型系统不同，因为在Go中，任何满足`Equaler`的类型都可以作为实参传给T2.Equal，在运行时我们必须检查参数是否为`T2`类型。有些语言在编译时就安排了这种保证。

​	一个相关的例子则恰恰相反：

```go linenums="1"
type Opener interface {
   Open() Reader
}

func (t T3) Open() *os.File
```

在Go中，`T3`并不满足`Opener`，尽管它在其他语言中可能满足。

​	虽然在这种情况下，Go的类型系统确实对程序员的帮助较小，但由于缺乏子类型，关于接口满足的规则非常容易说明：`函数的名称和签名是否正是接口的名称和签名？`Go的规则也很容易有效实现。我们觉得这些好处抵消了自动类型推导的不足。如果有一天Go采用了某种形式的多态类型，我们希望有一种方法来表达这些例子的想法，并让它们得到静态检查。

### Can I convert a []T to an []interface{}? 我可以将一个[]T转换为一个[]interface{}吗？

​	不能直接转换。语言规范不允许这样做，因为这两种类型在内存中没有相同的表示。有必要将元素单独复制到目标切片中。这个例子将一个`int`的切片转换为`interface{}`的切片：

```go linenums="1"
t := []int{1, 2, 3, 4}
s := make([]interface{}, len(t))
for i, v := range t {
    s[i] = v
}
```

### Can I convert []T1 to []T2 if T1 and T2 have the same underlying type? 如果T1和T2有相同的底层类型，我可以将[]T1转换成[]T2吗？

​	这段代码样本的最后一行不能编译。

```go linenums="1"
type T1 int
type T2 int
var t1 T1
var x = T2(t1) // OK
var st1 []T1
var sx = ([]T2)(st1) // NOT OK
```

​	`在Go中，类型与方法紧密相连，每个命名的类型都有一个（可能是空的）方法集。`一般的规则是，您可以改变被转换的类型的名称（从而可能改变其方法集），`但您不能改变复合类型的元素的名称（和方法集）`。Go要求您对类型转换进行明确说明。

### Why is my nil error value not equal to nil? 为什么我的nil错误值不等于nil？

​	在底层，接口被实现为两个元素，一个类型`T`和一个值`V`。`V`是一个具体的值，如`int`、`struct`或指针，绝不是接口本身，并且具有类型`T`。例如，如果我们在一个接口中存储`int`值3，那么产生的接口值就有，示意为（`T=int`，V=3）。值`V`也被称为接口的动态值，因为在程序执行过程中，一个给定的接口变量可能持有不同的值`V`（以及相应的类型`T`）。

​	只有当`V`和`T`都未设置时，接口值才是`nil`，（`T=nil`，`V未设置`），特别是，一个`nil`接口将永远持有一个`nil`类型。如果我们在一个接口值里面存储一个`*int`类型的`nil`指针，那么不管指针的值是什么，内部类型都是`*int`：（`T=*int`，`V=nil`）。因此，即使里面的指针值`V`是`nil`，这样的接口值也是非`nil`的。

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

​	如果一切顺利，该函数返回一个`nil`  `p`，所以返回值是一个`error`接口值持有（`T=*MyError`, `V=nil`）。这意味着，如果调用者将返回的错误与`nil`进行比较，即使没有发生什么坏事，也会一直看起来像有一个错误。为了向调用者返回一个适当的`nil`错误，该函数必须返回一个明确的`nil`：

```go linenums="1"
func returnsError() error {
    if bad() {
        return ErrBad
    }
    return nil
}
```

​	对于返回错误的函数来说，最好在其签名中使用`error`类型（就像我们上面做的那样），而不是一个具体的类型，如`*MyError`，以帮助保证错误被正确创建。作为一个例子，[os.Open](https://go.dev/pkg/os/#Open)返回一个`error`，尽管如果不是`nil`，它总是具体类型`*os.PathError`。

​	只要使用接口，就会出现与这里描述的类似情况。只要记住，`如果有任何具体的值被存储在接口中，接口就`不会是`nil`。更多信息，请参见[反射的法则]({{< ref "/goBlog/2011/TheLawsOfReflection" >}})。

### Why are there no untagged unions, as in C? 为什么没有像C语言那样的无标记的联合体？

​	没有标记的联合体会违反 Go 的内存安全保证。

### Why does Go not have variant types? 为什么 Go 没有变体类型？

​	变体类型，也被称为代数类型，提供了一种方法来指定一个值可能采取一系列其他类型中的一个，但只采取这些类型。在系统编程中，一个常见的例子是指定一个错误是，例如，一个网络错误、一个安全错误或一个应用程序错误，并允许调用者通过检查错误的类型来分辨问题的来源。另一个例子是语法树，其中每个节点可以是不同的类型：声明、语句、赋值等等。

​	我们考虑过在Go中加入变体类型，但经过讨论后决定不加入，因为它们与接口的重叠方式令人困惑。如果变体类型的元素本身就是接口，会发生什么？

​	另外，变体类型所涉及的一些内容已经被语言所涵盖。错误的例子很容易表达，用一个接口值来保持错误，用一个类型转换来区分情况。语法树的例子也是可以做到的，虽然没有那么优雅。

### Why does Go not have covariant result types? 为什么Go没有协变结果类型？

​	协变结果类型意味着一个接口，如

```go linenums="1"
type Copyable interface {
    Copy() interface{}
}
```

的方法会被满足

```go linenums="1"
func (v Value) Copy() Value
```

​	因为`Value`实现了这个空接口。`在Go中，方法类型必须完全匹配，所以Value并没有实现Copyable`。Go将一个类型所做的事情的概念——它的方法——与类型的实现分开。`如果两个方法返回不同的类型，它们就不是在做同一件事。`希望得到协变结果类型的程序员往往试图通过接口来表达类型的层次结构。在Go中，接口和实现之间的分离是比较自然的。

## Values 值

### Why does Go not provide implicit numeric conversions? 为什么 Go 不提供隐式数值转换？

​	在C语言中，数字类型之间的自动转换所带来的便利被其造成的混乱所抵消。一个表达式什么时候是无符号的？数值有多大？它是否会溢出？结果是否可移植，与执行它的机器无关？这也使编译器变得复杂；"通常的算术转换 "不容易实现，而且在不同的架构上也不一致。出于可移植性的考虑，我们决定以代码中一些明确的转换为代价，使事情变得清晰明了。在Go中对常量的定义——任意精度的值，没有符号和大小注释——大大改善了问题。

​	一个相关的细节是，与C不同，`int`和`int64`是不同的类型，即使`int`是一个64位的类型。`int` 类型是通用的；如果您关心一个整数拥有多少位，Go 鼓励您明确地表达出来。

### How do constants work in Go? 常量在 Go 中是如何工作的？

​	虽然 Go 对不同数字类型的变量之间的转换有严格的规定，但语言中的常量要灵活得多。字面常量如`23`、`3.14159`和[math.Pi](https://go.dev/pkg/math/#pkg-constants)占据了一种理想的数字空间，有任意的精度，没有溢出或下溢。例如，`math.Pi`的值在源代码中被指定有63处，涉及该值的常量表达式所保持的精度超过了`float64`所能容纳的。只有当常量或常量表达式被分配到一个变量——程序中的内存位置时，它才成为一个具有通常浮点属性和精度的 "computer"数字。

​	另外，`由于常量只是数字，而不是类型化的数值，Go中的常量可以比变量更自由地使用`，从而缓解了严格的转换规则带来的一些尴尬局面。我们可以写出这样的表达式

```go linenums="1"
sqrt2 := math.Sqrt(2)
```

这样的表达式，而不会被编译器抱怨，因为理想数字`2`可以被安全、准确地转换为`float64`来调用`math.Sqrt`。

​	一篇题为 "[Constants]({{< ref "/goBlog/2014/Constants" >}}) "的博文更详细地探讨了这个话题。

### Why are maps built in? 为什么内建映射？

​	和字符串的原因一样：它们是如此强大和重要的数据结构，提供一个具有语法支持的优秀实现可以使编程更加愉快。我们相信Go的映射实现足够强大，可以满足绝大多数的使用。如果一个特定的应用可以从自定义的实现中受益，那么就可以写一个，但在语法上就不那么方便了；这似乎是一个合理的权衡。

### Why don't maps allow slices as keys? 为什么映射不允许将切片作为键？

​	映射查询需要一个相等运算符，而切片并没有实现这个运算符。他们没有实现相等性，因为相等性在这种类型上没有得到很好的定义；有多种考虑，涉及到浅层与深层比较、指针与值比较、如何处理递归类型等等。我们可能会重新审视这个问题——为切片实现相等并不会使任何现有的程序失效，但是在没有明确切片相等的含义的情况下，暂时不考虑这个问题比较简单。

​	在Go 1中，与之前的版本不同，为结构体和数组定义了相等性，所以这类类型可以作为映射键使用。不过，切片仍然没有相等性的定义。

### Why are maps, slices, and channels references while arrays are values? 为什么映射、切片和通道是引用而数组是值？

​	关于这个话题有很多历史。早期，map和channel在语法上是指针，不可能声明或使用一个非指针实例。此外，我们还为数组应该如何工作而挣扎。最终我们决定，指针和值的严格分离将会使语言更难使用。将这些类型改变为对相关共享数据结构的引用，解决了这些问题。这一改变给语言增加了一些令人遗憾的复杂性，但对可用性产生了很大的影响。Go在推出后成为了一种更有生产力、更舒适的语言。

## Writing Code 编写代码

### How are libraries documented? 库是如何被文档化的？

​	有一个用Go编写的程序`godoc`，可以从源代码中提取包的文档，并将其作为网页提供给声明、文件等的链接。一个实例正在[go.dev/pkg/](https://go.dev/pkg/)上运行。事实上，`godoc`实现了[golang.org/](https://go.dev/)的完整网站。

​	一个`godoc`实例可以被配置为对它所显示的程序中的符号提供丰富的、交互式的静态分析；细节在[这里](../godoc/analysis/help)列出。

​	为了从命令行访问文档，[go](https://go.dev/pkg/cmd/go/)工具有一个[doc](https://go.dev/pkg/cmd/go/#hdr-Show_documentation_for_package_or_symbol)子命令，为相同的信息提供一个文本接口。

### Is there a Go programming style guide? 是否有Go编程风格指南？

​	虽然没有明确的风格指南，但肯定有一个可识别的 "Go 风格"。

​	Go 已经建立了惯例来指导命名、布局和文件组织方面的决策。文件 [Effective Go](../UsingAndUnderstandingGo/EffectiveGo) 包含了一些关于这些主题的建议。更直接地说，程序`gofmt`是一个漂亮的打印机，其目的是强制布局规则；它取代了允许解释的通常的注意事项的概要。存储库中的所有Go代码，以及开源世界中的绝大部分，都已经通过`gofmt`运行。

​	题为[Go Code Review Comments](../CodeReviewComments)的文件是一个非常简短的文章集，涉及到程序员经常忽略的Go习惯用语的细节。对于为Go项目做代码审查的人来说，它是一个方便的参考。

### How do I submit patches to the Go libraries? 我如何提交Go库的补丁？

​	库的源代码在版本库的 `src` 目录中。如果您想做一个重大的改变，请在开始之前在邮件列表中讨论。

​	有关如何进行的更多信息，请参见文档[Contributing to the Go project](../References/ContributionGuidel)。

### Why does "go get" use HTTPS when cloning a repository? 为什么 "go get" 在克隆版本库时使用 HTTPS？

​	公司通常只允许标准的TCP端口80（HTTP）和443（HTTPS）的出站流量，阻止其他端口的出站流量，包括TCP端口9418（git）和TCP端口22（SSH）。当使用HTTPS而不是HTTP时，`git`默认执行证书验证，提供保护以防止中间人、窃听和篡改攻击。因此，`go get`命令使用HTTPS来保证安全。

​	`Git` 可以被配置为通过 HTTPS 进行认证，或者使用 SSH 来代替 HTTPS。要通过HTTPS认证，可以在git查阅的`$HOME/.netrc`文件中添加一行。

```
machine github.com login USERNAME password APIKEY
```

​	对于GitHub账户，密码可以是[个人访问令牌](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/)。

​	`Git` 也可以被配置为使用 SSH 来代替 HTTPS 来处理与给定前缀匹配的 URL。例如，要对所有的 GitHub 访问使用 SSH，请在 `~/.gitconfig` 中添加这些行：

```
[url "ssh://git@github.com/"]
	insteadOf = https://github.com/
```

### How should I manage package versions using "go get"? 我应该如何使用 "go get "来管理软件包的版本？

​	Go工具链有一个内置系统，用于管理相关软件包的版本集，称为模块。模块是在 [Go 1.11](https://go.dev/doc/go1.11#modules) 中引入的，并从 [1.14](https://go.dev/doc/go1.14#introduction) 开始可用于生产。

​	要创建一个使用模块的项目，运行[go mod init](https://go.dev/ref/mod#go-mod-init)。这个命令会创建一个`go.mod`文件来跟踪依赖版本。

```shell
go mod init example/project
```

​	要添加、升级或降级一个依赖项，请运行 `go get`：

```shell
go get golang.org/x/text@v0.3.5
```

​	参见[Tutorial: Create a module](../GettingStarted/TutorialCreateAGoModule)，了解更多关于开始的信息。

​	请参阅 [Developing modules](../GoUserManual#developing-modules)，了解用模块管理依赖项的指南。

​	模块内的包在发展过程中应保持向后的兼容性，遵循[导入兼容性规则](https://research.swtch.com/vgo-import)：

> 如果一个旧包和一个新包有相同的导入路径，新的包必须向后兼容旧的包。

​	[Go 1的兼容性准则](../Go1AndTheFutureOfGoPrograms)在这里是一个很好的参考：不要删除导出的名字，鼓励标记的复合字面量，等等。如果需要不同的功能，添加一个新的名字，而不是改变一个旧的名字。

​	模块通过[语义版本化](https://semver.org/)和语义导入版本化来编纂这一点。如果需要打破兼容性，就在新的主版本上发布一个模块。主版本2和更高版本的模块需要一个[主版本后缀](../References/GoModulesReference/ModulesPackagesAndVersions#major-version-suffixes)作为其路径的一部分（如/v2）。这保留了导入兼容性规则：一个模块的不同主版本的包有不同的路径。

## Pointers and Allocation 指针和分配

### When are function parameters passed by value? 什么时候函数参数是按值传递的？

​	如同C系列的所有语言一样，`Go中的所有东西都是以值传递的`。也就是说，`一个函数总是得到一个被传递的东西的副本`，就像有一个赋值语句将值分配给参数一样。例如，向一个函数传递一个`int`值，就会得到`int`的副本，而`传递一个指针值就会得到指针的副本`，但不会得到它所指向的数据。(参见[后面一节](#should-I-define-methods-on-values-or-pointers)，讨论这对方法接收者的影响）。

​	`映射和切片值的行为类似于指针`：它们是包含指向底层映射或切片数据的指针的描述符。复制一个映射或切片值并不复制它所指向的数据。`复制一个接口值会复制存储在该接口值中的东西`。如果接口值持有一个结构，复制接口值就会复制该结构。`如果接口值持有一个指针，复制接口值会复制该指针，但同样不会复制它所指向的数据。`

​	请注意，这个讨论是关于操作的语义的。`实际的实现可以应用优化来避免复制`，只要这些优化不改变语义。

### When should I use a pointer to an interface? 什么时候应该使用一个指向接口的指针？

​	几乎不需要。指向接口值的指针只出现在罕见的、棘手的情况下，涉及到掩饰接口值的类型以延迟求值。

​	将一个接口值的指针传递给一个需要接口参数的函数，是一个常见的错误。编译器会抱怨这个错误，但这种情况仍然会让人困惑，因为有时需要[一个指针来满足一个接口](#why-do-t-and-t-have-different-method-sets)。我们的见解是，尽管一个指向具体类型的指针可以满足一个接口，但有一个例外，一个指向接口的指针永远不能满足一个接口。

​	考虑一下这个变量声明，

```go linenums="1"
var w io.Writer
```

​	打印函数`fmt.Fprintf`把一个满足`io.Writer`的值作为它的第一个参数——实现了典型的`Write`方法的东西。因此，我们可以写

```go linenums="1"
fmt.Fprintf(w, "hello, world\n")
```

然而，如果我们传递`w`的地址，程序将无法编译。

```go linenums="1"
fmt.Fprintf(&w, "hello, world\n") // Compile-time error. => 编译时错误。
```

​	有一个例外是，任何值，甚至是一个接口的指针，都可以被分配给一个空接口类型的变量（`interface{}`）。即便如此，如果值是一个指向接口的指针，几乎可以肯定是个错误；其结果可能是混乱的。

### Should I define methods on values or pointers? 我应该在值或指针上定义方法？

``` go
func (s *MyStruct) pointerMethod() { } // method on pointer
func (s MyStruct)  valueMethod()   { } // method on value
```

​	对于不习惯指针的程序员来说，这两个例子之间的区别可能会让人感到困惑，但实际上情况非常简单。当在一个类型上定义一个方法时，接收器（上述例子中的`s`）`的行为就像它是该方法的一个参数一样`。那么，将接收器定义为一个值还是一个指针，`与函数实参应该是一个值还是一个指针的问题是一样的`。这里有几个考虑因素。

​	首先，也是最重要的，方法是否需要修改接收器？如果需要，那么接收器必须是一个指针。(切片和映射作为引用，所以它们的故事更微妙一些，`但例如在一个方法中改变切片的长度，接收器仍然必须是一个指针`。) 在上面的例子中，如果`pointerMethod`修改了`s`的字段，那么调用者会看到这些变化，但是`valueMethod`是用调用者实参的副本来调用的（这就是传递值的定义），所以它所做的变化对调用者来说是不可见的。

​	顺便说一下，`在Java中，方法接收者总是指针`，尽管它们的指针性质在某种程度上被掩盖了（有一个建议是在语言中增加值接收者）。Go中的值接收器是不寻常的。

​	其次是对效率的考虑。如果接收器很大，例如一个大的结构体，那么使用指针接收器就会廉价很多。

​	其次是一致性。如果该类型的一些方法必须有指针接收器，那么其他的也应该有，所以无论该类型如何使用，方法集都是一致的。详见关于[方法集](#why-do-t-and-t-have-different-method-sets)的章节。

​	对于基本类型、切片和小结构体等类型，值接收器是非常廉价的，所以除非方法的语义需要指针，否则值接收器是高效且清晰的。

### What's the difference between new and make? new和make之间有什么区别？

​	简而言之：`new`是分配内存，而`make`是初始化slice、map和channel类型。

​	更多细节请参见[Effective Go](../UsingAndUnderstandingGo/EffectiveGo#allocation-with-new-new)的相关章节。

### What is the size of an `int` on a 64 bit machine? 在64位机器上int的大小是多少？

​	`int`和`uint`的大小是由具体实现决定的，但在特定的平台上彼此是一样的。为了可移植性，依赖特定大小的值的代码应该使用明确大小的类型，如`int64`。在32位机器上，编译器默认使用32位整数，而在64位机器上，整数有64位。(历史上，这并不总是正确的）。

​	另一方面，浮点标量和复数类型总是有大小的（没有`float`或`complex`的基本类型），因为程序员在使用浮点数字时应该注意精度。用于（无类型的）浮点常量的默认类型是`float64`。因此`foo :=3.0`声明了一个`float64`类型的变量`foo`。对于一个由（无类型）常量初始化的`float32`变量，必须在变量声明中明确指定变量类型：

```
var foo float32 = 3.0
```

​	或者，常量必须被赋予一个转换的类型，如`foo := float32(3.0)`。

### How do I know whether a variable is allocated on the heap or the stack? 我如何知道一个变量是在堆上还是在栈上分配的？

​	从正确性的角度来看，您不需要知道。只要有对它的引用，Go中的每个变量就存在。实现所选择的存储位置与语言的语义无关。

​	存储位置确实对编写高效程序有影响。在可能的情况下，Go编译器会在一个函数的栈框架中分配属于该函数的局部变量。但是，如果编译器不能证明该变量在函数返回后没有被引用，那么编译器必须在垃圾收集的堆上分配该变量，以避免`悬空指针错误`。另外，如果一个局部变量非常大，把它存储在`堆（heap ）`上而不是`栈（stack）`上可能更有意义。

​	在目前的编译器中，如果一个变量的地址被占用，那么这个变量就是在堆上分配的候选变量。然而，基本的转义分析认识到有些情况下，这种变量不会活过函数的返回，可以驻留在栈上。

### Why does my Go process use so much virtual memory? 为什么我的Go进程会使用这么多的虚拟内存？

​	Go内存分配器保留了一个大的虚拟内存区域作为分配的场所。这个虚拟内存是特定Go进程的本地内存；保留的内存不会剥夺其他进程的内存。

​	要找到分配给Go进程的实际内存量，请使用Unix的`top`命令，并查阅`RES`（Linux）或`RSIZE`（macOS）列。

## Concurrency 并发性

### What operations are atomic? What about mutexes? 哪些操作是原子性的？互斥是什么？

​	关于Go中操作的原子性的描述可以在[Go Memory Model](../References/TheGoMemoryModel)文档中找到。

​	低级别的同步和原子原语可以在[sync](https://go.dev/pkg/sync)和[sync/atomic](https://go.dev/pkg/sync/atomic)包中找到。这些包适合于简单的任务，如增加引用计数或保证小规模的互斥。

​	对于更高层次的操作，如并发服务器之间的协调，更高层次的技术可以带来更好的程序，Go通过其goroutines和通道支持这种方法。例如，您可以构造您的程序，使每次只有一个goroutine负责一个特定的数据。这种方法被[Go 谚语](https://www.youtube.com/watch?v=PAAkCSZUG1c)所概括，

Do not communicate by sharing memory. Instead, share memory by communicating.

`不要通过共享内存进行通信。相反，通过通信来共享内存。`

​	关于这个概念的详细讨论，请参见 [通过通信共享内存](https://go.dev/doc/codewalk/sharemem/) 的代码练习及其[相关文章]({{< ref "/goBlog/2010/ShareMemoryByCommunicating" >}})。

​	大型并发程序可能会借用这两个工具包。

### Why doesn't my program run faster with more CPUs? 为什么我的程序在更多的CPU下运行得不快？

​	一个程序是否在更多的CPU下运行得更快，取决于它所解决的问题。Go语言提供了并发原语，如goroutines和channel，但只有当底层问题本质上是并行的时候，并发才能实现并行化。本质上是顺序的问题不能通过添加更多的CPU来加速，而那些可以被分解成可以并行执行的片段的问题则可以加速，有时甚至可以大大加速。

​	有时，增加更多的CPU会使程序变慢。在实际应用中，当使用多个操作系统线程时，那些花在同步或通信上的时间多于进行有用的计算的程序可能会出现性能下降。这是因为在线程之间传递数据需要切换上下文，这需要很大的成本，而这种成本会随着CPU的增加而增加。例如，Go规范中的[素数筛例子]({{< ref "/langSpec/Packages#an-example-package">}})没有明显的并行性，尽管它启动了许多goroutine；增加线程（CPU）的数量更有可能使它变慢而不是变快。

​	关于这个话题的更多细节，请看题为 [并发性不是并行性]({{< ref "/goBlog/2013/ConcurrencyIsNotParallelism" >}}) 的讲座。

### How can I control the number of CPUs? 我怎样才能控制CPU的数量？

​	可用于同时执行goroutines的CPU数量由`GOMAXPROCS` shell环境变量控制，其默认值是可用的CPU核心数量。因此，具有并行执行潜力的程序应该在多CPU机器上默认实现。要改变使用的并行CPU数量，可以设置环境变量或使用 runtime 包的类似[功能](https://go.dev/pkg/runtime/#GOMAXPROCS)来配置运行时支持，以利用不同数量的线程。`将其设置为1会消除真正并行的可能性，迫使独立的goroutine轮流执行`。

​	runtime 可以分配比`GOMAXPROCS`值更多的线程，以服务于多个未处理的I/O请求。`GOMAXPROCS`只影响到有多少个goroutine可以同时实际执行；任意多的goroutine可能会在系统调用中被阻塞。

​	Go的`goroutine调度器`并不像它需要的那样好，尽管它随着时间的推移已经有所改进。在未来，它可能会更好地优化其对操作系统线程的使用。目前，如果有性能问题，在每个应用的基础上设置`GOMAXPROCS`可能会有帮助。

### Why is there no goroutine ID? 为什么没有goroutine的ID？

​	goroutines没有名字；它们只是匿名的工作程序。它们没有向程序员暴露出唯一的标识符、名称或数据结构。有些人对此感到惊讶，以为`go`语句会返回一些可以用来访问和控制goroutine的项目。

​	goroutines是匿名的，其根本原因是为了在编程并发代码时可以使用完整的Go语言。相比之下，当线程和goroutines被命名时，形成的使用模式会限制使用它们的库的能力。

​	下面是一个困难的例子。一旦人们命名了一个goroutine，并围绕它构建了一个模型，它就变得很特别，人们就会倾向于将所有的计算与该goroutines联系起来，而忽略了使用多个可能共享的goroutine进行处理的可能性。如果 `net/http` 包将每个请求的状态与一个 goroutine 关联起来，那么在提供请求时，客户端将无法使用更多的 goroutine。

​	此外，像那些要求所有处理都发生在 "主线程 "上的图形系统库的经验表明，当部署在一种并发语言中时，这种方法是多么的笨拙和局限。特殊线程或goroutine的存在，迫使程序员扭曲程序，以避免崩溃和其他因无意中在错误线程上操作而引起的问题。

​	对于那些特定的goroutine确实很特殊的情况，语言提供了诸如通道等功能，可以用灵活的方式与之交互。

## Functions and Methods 函数和方法

### Why do T and *T have different method sets? 为什么T和`*T`有不同的方法集？

​	正如[Go规范]({{< ref "/langSpec/Types">}})所说，一个类型`T`的方法集包括所有接收者类型为`T`的方法，而相应的指针类型`*T`的方法集包括所有接收者为`*T`或`T`的方法，这意味着`*T`的方法集包括`T`的方法集，但不是相反。

​	这种区别的产生是因为如果一个接口值包含一个指针`*T`，方法调用可以通过取消引用指针来获得一个值，但是如果一个接口值包含一个值`T`，方法调用就没有安全的方法来获得一个指针。(这样做将允许一个方法修改接口内的值的内容，这是语言规范所不允许的）。

​	即使在编译器可以获取传递给方法的值的地址的情况下，如果方法修改了这个值，则更改将在调用方中丢失。举个例子，如果[bytes.Buffer](https://go.dev/pkg/bytes/#Buffer)的`Write`方法使用了一个值接收器而不是一个指针，那么这段代码：

```go linenums="1"
var buf bytes.Buffer
io.Copy(buf, os.Stdin)
```

将会把标准输入复制到`buf`的一个副本中，而不是复制到`buf`本身。这几乎不是我们想要的行为。

### What happens with closures running as goroutines? 作为goroutines运行的闭包会发生什么？

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

​	人们可能误以为会看到`a, b, c`作为输出。您可能看到的是 `c, c, c`。这是因为循环的每一次迭代都使用同一个变量`v`的实例，所以每个闭包都共享这个单一的变量。当闭包运行时，它打印的是执行 `fmt.Println` 时的 `v` 值，但 `v` 可能在 goroutine 启动后被修改过。为了帮助在发生这种情况和其他问题之前发现它们，运行[go vet](https://go.dev/cmd/go/#hdr-Run_go_tool_vet_on_packages)。

​	为了在每个闭包启动时将`v`的当前值与之绑定，必须修改内循环以在每个迭代中创建一个新变量。一种方法是将该变量作为参数传递给闭包：

```go linenums="1" hl_lines="2 2"
    for _, v := range values {
        go func(u string) {
            fmt.Println(u)
            done <- true
        }(v)
    }
```

​	在这个例子中，`v`的值被作为一个参数传递给匿名函数。然后，该值可以在函数中作为变量`u`被访问。

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

​	语言的这种行为，即不为每个迭代定义一个新的变量，现在回想起来可能是一个错误。`它可能会在以后的版本中得到解决`，但`为了兼容，在Go 1版本中不能改变`。

## Control flow 控制流

### Why does Go not have the `?:` operator? 为什么Go没有?: 操作符？

​	Go中没有三元测试操作。您可以使用下面的方法来实现同样的结果：

```go linenums="1"
if expr {
    n = trueVal
} else {
    n = falseVal
}
```

​	Go中没有`?:`的原因是该语言的设计者看到该操作经常被用来创建令人难以理解的复杂表达式。if-else形式虽然较长，但无疑是更清晰的。一门语言只需要一个条件控制流结构。

## Type Parameters 类型参数

### Why does Go have type parameters? 为什么 Go 有类型参数？

​	类型参数允许所谓的泛型编程，其中函数和数据结构是以类型来定义的，这些类型在以后使用这些函数和数据结构时被指定。例如，它们使我们有可能编写一个返回任何有序类型的两个值的最小值的函数，而不必为每个可能的类型编写一个单独的版本。关于更深入的解释和例子，请看博文 [Why Generics?]({{< ref "/goBlog/2019/WhyGenerics" >}})。

### How are generics implemented in Go? 泛型是如何在Go中实现的？

​	编译器可以选择是单独编译每个实例，还是将合理相似的实例作为一个单一的实现来编译。单一实现的方法类似于带有接口参数的函数。不同的编译器会针对不同的情况做出不同的选择。标准的Go 1.18编译器通常为每个具有相同形状的类型实参发出一个单一的实例，其中的形状是由类型的属性决定的，例如它包含的指针的大小和位置。未来的版本将对编译时、运行时效率和代码大小之间的权衡进行试验。

### How do generics in Go compare to generics in other languages? Go中的泛型与其他语言的泛型相比如何？

​	所有语言中的基本功能都是相似的：可以使用以后指定的类型来编写类型和函数。尽管如此，还是有一些区别。

- Java

  在Java中，编译器在编译时检查泛型类型，但在运行时删除这些类型。这被称为[类型擦除](https://en.wikipedia.org/wiki/Generics_in_Java#Problems_with_type_erasure)。例如，一个在编译时被称为`List<Integer>`的Java类型在运行时将变成非泛型类型`List`。这意味着，例如，当使用Java形式的类型反射时，不可能区分`List<Integer>`类型的值和`List<Float>`类型的值。在Go中，泛型类型的反射信息包括完整的编译时类型信息。

  Java使用类型通配符，如`List<? extends Number>`或`List<? super Number>`来实现泛型的协变和反变。Go没有这些概念，这使得Go中的泛型变得更加简单。

- C++

  传统上C++模板不对类型实参执行任何约束，尽管C++20通过[概念](https://en.wikipedia.org/wiki/Concepts_(C%2B%2B))支持可选的约束。在Go中，对所有类型参数的约束是强制性的。C++20的概念被表达为小的代码片段，必须与类型实参一起编译。Go的约束是定义所有允许的类型实参集合的接口类型。C++支持模板元编程；Go则不支持。在实践中，所有的C++编译器都是在每个模板被实例化的地方进行编译；如上所述，Go可以也确实对不同的实例化使用不同的方法。

- Rust

  Rust版本的约束被称为`特性绑定`。在Rust中，特性绑定和类型之间的关联必须明确定义，要么在定义特性绑定的板块中，要么在定义类型的板块中。在Go中，类型实参隐式地满足约束，就像Go类型隐式地实现接口类型一样。Rust标准库为比较或加法等操作定义了标准特性；Go标准库则没有，因为这些可以通过接口类型在用户代码中表达。

- Python

  Python 不是静态类型语言，因此可以合理地说，所有 Python 函数在默认情况下总是泛型的：它们总是可以用任何类型的值来调用，并且在运行时检测到任何类型的错误。

### Why does Go use square brackets for type parameter lists? 为什么Go对类型参数列表使用方括号？

​	Java 和 C++ 对类型参数列表使用`角括号`，如 Java `List<Integer>` 和 C++ `std::vector<int>`。然而，Go没有这个选项，`因为它导致了一个语法问题`：当解析一个函数内的代码时，例如`v := F<T>`，在看到`<`的时候，我们看到的是一个实例化还是一个使用`<`操作符的表达式，这一点是模糊的。如果没有类型信息，这是很难解决的。

例如，考虑一个语句，如

```go linenums="1"
    a, b = w < x, y > (z)
```

​	如果没有类型信息，就不可能决定赋值的右边是一对表达式（`w < x and y > z`），还是一个返回两个结果值的泛型函数实例化和调用（`(w<x, y>)(z)`）。

​	`Go的一个关键设计决定`是`在没有类型信息的情况下进行解析`，这在使用角括号表示泛型时似乎是不可能的。

​	Go在使用方括号方面并不是唯一的或原创的；还有其他语言，如`Scala`，也使用方括号来表示泛型代码。

### Why does Go not support methods with type parameters? 为什么 Go 不支持带有类型参数的方法？

​	Go 允许泛型类型拥有方法，`但除了接收者之外，这些方法的参数不能使用参数化类型`。一个类型的方法决定了该类型所实现的接口，但不清楚这对泛型的方法的参数化实参如何操作。这将需要在运行时实例化函数，或者为每个可能的类型实参实例化每个泛型函数。这两种方法似乎都不可行。更多细节，包括一个例子，请看[提案](https://go.dev/design/43651-type-parameters#no-parameterized-methods)。不使用带类型参数的方法，而是使用带类型参数的顶层函数，或者将类型参数添加到接收器类型中。

### Why can't I use a more specific type for the receiver of a parameterized type? 为什么我不能为参数化类型的接收器使用一个更具体的类型？

​	泛型类型的方法声明是用一个包括类型参数名称的接收器来写的。有些人认为可以使用一个特定的类型，产生一个只对某些类型实参有效的方法：

```go linenums="1"
type S[T any] struct { f T }

func (s S[string]) Add(t string) string {
    return s.f + t
}
```

​	这个操作失败了，出现了编译器错误，如操作符`+`没有定义在`s.f上`（由any限制的字符串类型的变量），尽管`+`操作符当然对预先声明的字符串类型有效。

​	这是因为在方法`Add`的声明中使用`string`只是为类型参数引入一个名字，而这个名字就是`string`。这是一件有效的，尽管很奇怪的事情。字段`s.f`的类型是`string`，不是通常预先声明的`string`类型，而是`S`的类型参数，在这个方法中被命名为`string`。由于类型参数的约束是`any`，所以不允许使用`+`运算符。=>仍有疑问？？

### Why can't the compiler infer the type argument in my program? 为什么编译器不能推断出我程序中的类型实参？

​	有很多情况下，程序员可以很容易地看到一个泛型或函数的类型实参必须是什么，但`语言不允许编译器推断它`。`类型推断是有意限制的，以确保在推断哪种类型方面不会有任何混淆。`其他语言的经验表明，在阅读和调试程序时，意外的类型推断会导致相当大的混乱。我们总是可以指定在调用中使用的显式类型实参。在未来，新的推理形式可能会被支持，只要规则保持简单和清晰。

## Packages and Testing 包和测试

### How do I create a multifile package? 我如何创建一个多文件包？

​	把包的所有源文件单独放在一个目录中。源文件可以随意引用不同文件中的项；不需要正向声明或头文件。

​	除了被分割成多个文件外，该包的编译和测试就像一个单文件包。

### How do I write a unit test? 如何写一个单元测试？

​	在与您的包源文件相同的目录下，创建一个以`_test.go`结尾的新文件。在该文件中，`import "testing"`并编写以下形式的函数

```go linenums="1"
func TestFoo(t *testing.T) {
    ...
}
```

​	在该目录下运行`go test`。该脚本找到`Test`的函数，建立一个测试二进制文件，并运行它。

​	更多细节请参见[How to Write Go Code](../GettingStarted/HowToWriteGoCode)文档，[testing](https://go.dev/pkg/testing/)包和[go test](https://go.dev/cmd/go/#hdr-Test_packages)子命令。

### Where is my favorite helper function for testing? 我最喜欢的测试辅助函数在哪里？

​	Go 的标准[testing](https://go.dev/pkg/testing/)包使编写单元测试变得容易，但它缺乏其他语言的测试框架所提供的功能，如断言函数。本文的[前一部分](#why-does-go-not-have-assertions)解释了为什么Go没有断言，同样的论点也适用于测试中使用断言。正确的错误处理意味着在一个测试失败后让其他测试运行，这样调试失败的人就可以得到一个完整的错误信息。对于测试来说，报告 `isPrime` 对2、3、5和7(或对2、4、8和16)给出错误答案比报告 `isPrime` 对2给出错误答案更有用，因此没有（什么比）运行更多的测试更有用。触发测试失败的程序员可能不熟悉失败的代码。现在花时间写一个好的错误信息，在以后测试失败时就会得到回报。

​	与此相关的一点是，测试框架往往会发展成自己的迷您语言，带有条件、控件和打印机制，但Go已经具备所有这些功能，为什么还要重新创建它们呢？我们宁愿用Go来写测试；这样就少了一种需要学习的语言，而且这种方法可以使测试简单明了，易于理解。

​	如果编写好的错误所需的额外代码量看起来是重复的和压倒性的，且测试是由表驱动的，则在数据结构中定义的输入和输出的列表上迭代（Go对数据结构体字面量有很好的支持），可能效果更好。编写一个好的测试和好的错误信息的工作将被分摊到许多测试案例中。标准 Go 库中有很多说明性的例子，例如 [fmt 包的格式化测试](https://go.dev/src/fmt/fmt_test.go)。

### Why isn't *X* in the standard library? 为什么X不在标准库中？

​	标准库的目的是支持运行时，连接到操作系统，并提供许多 Go 程序需要的关键功能，如格式化 I/O 和网络。它还包含对网络编程很重要的元素，包括密码学和对HTTP、JSON和XML等标准的支持。

​	没有明确的标准来定义所包含的内容，因为在很长一段时间里，这是唯一的Go库。然而，有一些标准定义了今天被添加的内容。

​	新加入标准库的情况很少，而且加入标准库的标准很高。包含在标准库中的代码要承担大量的持续维护费用（通常由原作者以外的人承担），要遵守[Go 1的兼容性承诺](../Go1AndTheFutureOfGoPrograms)（阻止对API中任何缺陷的修复），并且要遵守Go的[发布时间表](../GoReleaseCycle)，阻止用户快速获得 bug 修复。

​	大多数新的代码应该生活在标准库之外，可以通过[go  tool](https://go.dev/cmd/go/)的`go get`命令进行访问。这样的代码可以有自己的维护者、发布周期和兼容性保证。用户可以在[https://pkg.go.dev/](https://pkg.go.dev/)上找到软件包并阅读其文档。

​	尽管标准库中有些部分并不真正属于x，比如`log/syslog`，但由于Go 1的兼容性承诺，我们继续维护库中的一切。但我们鼓励大多数新代码在其他地方使用。

## Implementation 实现

### What compiler technology is used to build the compilers? 使用什么编译器技术来构建编译器？

​	有几个Go的生产型编译器，还有一些正在为不同平台开发的编译器。

​	默认的编译器 `gc` 包含在 Go 发行版中，是对 `go` 命令支持的一部分。`Gc`最初是用C语言编写的，因为启动困难——您需要一个Go编译器来建立一个Go环境。但现在事情有了进展，`从Go 1.5版本开始，编译器就是一个Go程序`。编译器是使用自动翻译工具从C语言转换为Go语言的，如这个[设计文档](../Other/Go1_3PlusCompilerOverhaul)和[讲座](https://go.dev/talks/2015/gogo.slide#1)中所述。因此，编译器现在是 "`自举（self-hosting）` "的，这意味着我们需要面对启动（bootstrapping ）的问题。解决的办法是已经有了一个正常工作的Go的安装，就像通常有了一个正常工作的C的安装一样。关于如何从源码建立一个新的Go环境的故事在[这里](../Other/Go1_5BootstrapPlan)和[这里](../GettingStarted/InstallingGoFromSource)有描述。

​	`Gc`是用Go写的，有一个递归解析器，并使用一个定制的加载器，也是用Go写的，但基于Plan 9加载器，用来生成ELF/Mach-O/PE二进制文件。

​	在项目开始时，我们考虑过使用LLVM的`gc`，但认为它太大太慢，无法满足我们的性能目标。更重要的是，现在回想起来，使用LLVM会使我们更难引入一些ABI和相关的变化，例如堆栈管理，Go需要这些变化，但这些变化并不是标准C设置的一部分。然而，一个新的[LLVM实现](../Other/gollvm)现在已经开始了。

​	`Gccgo`编译器是一个用C++编写的前端，带有一个与标准GCC后端耦合的递归解析器。

​	Go被证明是实现Go编译器的一种很好的语言，尽管这并不是它最初的目标。从一开始就不是`自举（self-hosting）`的，这使得Go的设计能够集中在它最初的使用情况上，也就是网络服务器。如果我们一开始就决定Go应该自我编译，那么我们最终可能会得到一个更针对编译器构建的语言，这是一个值得追求的目标，但不是我们最初的目标。

​	虽然`gc`没有使用它们（还没有？），但在[go](https://go.dev/pkg/go/)包中有一个本地的 lexer 和解析器，还有一个本地的[类型检查器](https://go.dev/pkg/go/types)。

### How is the run-time support implemented? 运行时支持是如何实现的？

​	同样由于引导（bootstrapping ）问题，运行时代码最初主要是用C语言编写的（有一小部分汇编程序），但后来被翻译成Go语言（除了一些汇编程序部分）。`Gccgo`的运行时支持使用`glibc`。`gccgo`编译器使用一种叫做`分段堆栈（segmented stacks）`的技术来实现goroutines，这种技术得到了最近对 gold 链接器的修改的支持。`Gollvm`同样是建立在相应的LLVM基础设施上。

### Why is my trivial program such a large binary? 为什么我的微不足道的程序会有这么大的二进制文件？

​	`gc`工具链中的链接器默认会创建静态链接的二进制文件。因此，所有 Go 二进制文件都包括 Go 运行时，以及支持动态类型检查、反射、甚至恐慌时栈跟踪所需的运行时类型信息。

​	一个简单的C语言 "hello, world "程序在Linux上使用gcc静态编译和链接，大约是750 kB，包括一个`printf`的实现。一个使用`fmt.Printf`的同等Go程序重达几兆字节，但这包括更强大的运行时支持以及类型和调试信息。

​	用 `gc` 编译的 Go 程序可以用 `-ldflags=-w` 标志链接，以禁止生成 `DWARF`，从二进制文件中删除调试信息，但没有其他功能损失。这可以大大减少二进制文件的大小。

### Can I stop these complaints about my unused variable/import? 我可以停止这些关于我的未使用变量/导入的抱怨吗？

​	未使用的变量的存在可能表明了一个错误，而未使用的导入只是减缓了编译速度，随着时间的推移，程序的代码和程序员的积累，这种影响会变得很大。由于这些原因，Go拒绝编译带有未使用的变量或导入的程序，以短期的方便换取长期的编译速度和程序的清晰。

​	尽管如此，在开发代码时，临时创建这些情况是很常见的，在程序编译前必须将其编辑出来，这可能是很烦人的。

​	有些人要求提供一个编译器选项，以关闭这些检查，或至少将其减少为警告。不过这样的选项还没有被添加，因为编译器选项不应该影响语言的语义，而且Go编译器不报告警告，只报告妨碍编译的错误。

​	没有警告的原因有两个。首先，如果它值得抱怨，就值得在代码中加以修正。（如果不值得修正，也就不值得一提了。）第二，让编译器产生警告会鼓励实现者对那些会使编译变得混乱的弱情况发出警告，从而掩盖应该被修复的实际错误。

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

​	现在，大多数Go程序员都使用[goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)工具，它会自动重写 Go 源文件以获得正确的导入，从而在实践中消除了未使用的导入问题。这个程序很容易连接到大多数编辑器，以便在编写 Go 源文件时自动运行。

### Why does my virus-scanning software think my Go distribution or compiled binary is infected? 为什么我的病毒扫描软件认为我的 Go 发行版或编译的二进制文件被感染了？

​	这种情况很常见，尤其是在Windows机器上，而且几乎都是假阳性。商业病毒扫描程序常常被Go二进制文件的结构所迷惑，它们不象其他语言编译的文件那样经常看到这种结构。

​	如果您刚刚安装了Go发行版，而系统报告说它被感染了，这肯定是个错误。为了真正彻底，您可以通过将校验和与[下载页面](https://go.dev/dl/)上的校验和进行比较来验证下载。

​	在任何情况下，如果您认为报告是错误的，请向您的病毒扫描器的供应商报告一个错误。也许随着时间的推移，病毒扫描器可以学会理解Go程序。

## Performance 性能表现

### Why does Go perform badly on benchmark X? 为什么Go在基准X上的表现很差？

​	Go 的设计目标之一是使同类程序的性能接近 C，但在一些基准测试中，Go 的表现相当差，包括 [golang.org/x/exp/shootout](https://go.googlesource.com/exp/+/master/shootout/) 中的几个基准。最慢的是依赖于Go中没有可比性能版本的库。例如，[pidigits.go](https://go.googlesource.com/exp/+/master/shootout/pidigits.go)依赖于一个多精度数学包，而C版本与Go不同，（C版本它）使用[GMP](https://gmplib.org/)（用优化的汇编程序编写）。依赖于正则表达式的基准测试（例如[regex-dna.go](https://go.googlesource.com/exp/+/master/shootout/regex-dna.go)）基本上是将Go的本地[regexp package](https://go.dev/pkg/regexp)与成熟的、高度优化的正则表达式库（如PCRE）进行比较。

​	基准测试游戏是通过广泛的调整来赢得的，大多数基准测试的Go版本需要注意。如果您测量可比较的C和Go程序（[reverse-complement.go](https://go.googlesource.com/exp/+/master/shootout/reverse-complement.go)就是一个例子），您会发现这两种语言的原始性能比这个套件所显示的要接近得多。

​	尽管如此，仍有改进的余地。编译器很好，但可以做得更好，许多库需要进行大量的性能工作，而且垃圾收集器还不够快。(即使它够快，小心避免产生不必要的垃圾也会产生巨大的影响。)

​	在任何情况下，Go通常是非常有竞争力的。随着语言和工具的发展，许多程序的性能都有了明显的改善。请参阅有关 [profiling Go programs]({{< ref "/goBlog/2011/ProfilingGoPrograms" >}}) 的博文，了解一个信息丰富的例子。

## Changes from C - 与 C 的变化

### Why is the syntax so different from C? 为什么语法与C语言如此不同？

​	除了声明语法之外，其他的差别并不大，而且源于两个愿望。首先，语法应该给人以轻松的感觉，没有太多的强制性关键字、重复或奥秘。第二，该语言被设计成易于分析，不需要符号表就可以进行解析。这使得它更容易建立工具，如调试器、依赖分析器、自动文档提取器、IDE插件等等。C及其后代在这方面是出了名的困难。

### Why are declarations backwards? 为什么声明是相反的？

​	如果您习惯于C语言，它们就是相反的。在C语言中，变量的声明就像表示其类型的表达式一样，这是一个很好的想法，但是类型语法和表达式语法不能很好地混合，其结果可能是混乱的；考虑到函数指针。Go大多将表达式和类型语法分开，这就简化了事情（对指针使用前缀`*`是证明这一规则的一个例外）。在C语言中，声明

```go linenums="1"
    int* a, b;
```

声明了`a`是一个指针，但没有声明`b`；在Go中

```go linenums="1"
    var a, b *int
```

声明两者都是指针。这样更清晰，更有规律。另外，`:=`的短声明形式认为，完整的变量声明应该呈现与`:=`相同的顺序，所以

```go linenums="1"
    var a uint64 = 1
```

和

```go linenums="1"
    a := uint64(1)
```

有同样的效果。

​	通过为类型制定独特的语法，而不仅仅是表达式语法，解析工作也得到了简化；`func`和`chan`等关键字可以让事情变得清晰。

​	更多细节请参见[Go's Declaration Syntax]({{< ref "/goBlog/2010/GosDeclarationSyntax" >}})一文。

### Why is there no pointer arithmetic? 为什么没有指针算术？

Safety. Without pointer arithmetic it's possible to create a language that can never derive an illegal address that succeeds incorrectly. Compiler and hardware technology have advanced to the point where a loop using array indices can be as efficient as a loop using pointer arithmetic. Also, the lack of pointer arithmetic can simplify the implementation of the garbage collector.

​	安全性。如果没有指针算术，就有可能创造出一种语言，永远无法推导出一个非法的地址，从而错误地成功。编译器和硬件技术已经发展到这样的程度：使用数组索引的循环可以和使用指针算术的循环一样有效。而且，`没有指针算术可以简化垃圾收集器的实现`。

### Why are `++` and `--` statements and not expressions? And why postfix, not prefix? 为什么是++和--语句而不是表达式？以及为什么是后缀而不是前缀？

​	没有指针运算，前缀和后缀增量运算符的方便性就会下降。通过将它们从表达式层次中完全移除，表达式的语法被简化了，围绕`++`和`--`的求值顺序的混乱问题（考虑`f(i++)`和`p[i] = q[++i]`）也被消除了。这种简化是非常重要的。至于后缀和前缀，两者都可以使用，但后缀版本更为传统；坚持使用前缀是由STL引起的，STL是一种语言的库，讽刺的是，其名称包含后缀增量。

### Why are there braces but no semicolons? And why can't I put the opening brace on the next line? 为什么有大括号而没有分号？为什么我不能把开头的大括号放在下一行？

​	Go使用花括号对语句进行分组，这是使用过C语言的程序员都熟悉的语法。然而，分号是给解析器用的，不是给人用的，我们想尽可能地消除它们。为了实现这一目标，Go从`BCPL`中借鉴了一个技巧：分隔语句的分号在形式语法中存在，但在任何可能是语句结尾的行的末尾，由词法分析器自动注入，without lookahead。这在实践中效果很好，但有一个影响，那就是它强制采用了花括号样式。例如，一个函数的开头花括号不能单独出现在一行中。

​	有些人认为，lexer 应该进行查找，以允许括号出现在下一行。我们不同意这个观点。由于Go代码是由[gofmt]({{< ref "/cmd/gofmt">}})自动格式化的，所以必须选择一些风格。这种风格可能与您在C或Java中使用的不同，但Go是一种不同的语言，`gofmt`的风格与其他语言一样好。更重要的是——要重要得多的是——为所有Go程序提供单一的、程序化的强制格式的优势大大超过了特定风格的任何感知的缺点。还要注意的是，Go的风格意味着Go的交互式实现可以一行一行地使用标准语法，而不需要特别的规则。

### Why do garbage collection? Won't it be too expensive? 为什么要做垃圾收集？不会代价太大吗？

​	系统程序中最大的记账来源之一是管理分配对象的生命周期。在诸如C语言中，它是手动完成的，它可能会消耗程序员大量的时间，而且常常是致命的错误的原因。即使在像C++或Rust这样提供辅助机制的语言中，这些机制也会对软件的设计产生重大影响，往往会增加自身的编程开销。我们觉得消除这种程序员的开销是至关重要的，而过去几年垃圾收集技术的进步给了我们信心，它可以用足够低的成本地实现，并具有足够低的延迟，它可以成为网络系统的一个可行的方法。

​	并发编程的大部分困难都源于对象的生存期问题：当对象在线程之间传递时，保证它们安全释放就变得很麻烦。自动垃圾收集使并发代码更容易编写。当然，在并发环境中实现垃圾收集本身就是一个挑战，但一次就实现胜过在每个程序中都去（手动）实现，对每个人都有帮助。

​	最后，撇开并发性不谈，垃圾收集使接口更简单，因为它们不需要指定如何跨接口管理内存。

​	这并不是说最近在Rust等语言中为管理资源的问题带来新思路的工作是错误的；我们鼓励这项工作，并很高兴看到它的发展情况。但是Go采取了一种更传统的方法，即`通过垃圾收集来解决对象的生命周期问题，而且仅仅只是垃圾收集`。

​	目前的实现是一个标记-清除（mark-and-sweep）收集器。如果机器是多处理器，收集器会在一个单独的CPU核心上与主程序并行运行。近年来关于收集器的主要工作已经将停顿时间减少到了亚毫秒范围，甚至对于大堆也是如此，这几乎消除了网络服务器中垃圾收集的主要障碍之一。完善算法的工作仍在继续，进一步减少开销和延迟，并探索新的方法。Go团队的`Rick Hudson`在2018年[ISMM keynote]({{< ref "/goBlog/2018/GettingToGoTheJourneyOfGosGarbageCollector" >}})的主题演讲中描述了迄今为止的进展，并提出了一些未来的方法。

​	关于性能的话题，请记住，Go给了程序员对内存布局和分配相当大的控制权，比典型的垃圾收集型语言要多得多。细心的程序员可以通过很好地使用该语言来大幅减少垃圾收集的开销；请参阅为一个工作示例[剖析Go程序]({{< ref "/goBlog/2011/ProfilingGoPrograms" >}})的文章，其中包括Go的剖析工具的演示。