+++
title = "指南"
date = 2024-01-22T10:01:15+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go Style Guide - Go 风格指南

> 原文：[https://google.github.io/styleguide/go/guide](https://google.github.io/styleguide/go/guide)

> **Note:** This is part of a series of documents that outline [Go Style](https://google.github.io/styleguide/go/index) at Google. This document is **[normative](https://google.github.io/styleguide/go/index#normative) and [canonical](https://google.github.io/styleguide/go/index#canonical)**. See [the overview](https://google.github.io/styleguide/go/index#about) for more information.
>
> ​	注意：这是概述 Google 中 Go 风格的一系列文档的一部分。此文档具有规范性和权威性。有关更多信息，请参阅概述。



## Style principles 风格原则

There are a few overarching principles that summarize how to think about writing readable Go code. The following are attributes of readable code, in order of importance:

​	有一些总体原则总结了如何考虑编写可读的 Go 代码。以下是可读代码的属性，按重要性排序：

1. **[Clarity](https://google.github.io/styleguide/go/guide#clarity)**: The code’s purpose and rationale is clear to the reader.
   清晰度：代码的目的和基本原理对读者来说很清楚。
2. **[Simplicity](https://google.github.io/styleguide/go/guide#simplicity)**: The code accomplishes its goal in the simplest way possible.
   简单性：代码以最简单的方式实现其目标。
3. **[Concision](https://google.github.io/styleguide/go/guide#concision)**: The code has a high signal-to-noise ratio.
   简洁性：代码具有很高的信噪比。
4. **[Maintainability](https://google.github.io/styleguide/go/guide#maintainability)**: The code is written such that it can be easily maintained.
   可维护性：代码的编写方式使其易于维护。
5. **[Consistency](https://google.github.io/styleguide/go/guide#consistency)**: The code is consistent with the broader Google codebase.
   一致性：代码与更广泛的 Google 代码库一致。



### Clarity 清晰度

The core goal of readability is to produce code that is clear to the reader.

​	可读性的核心目标是生成对读者来说清晰的代码。

Clarity is primarily achieved with effective naming, helpful commentary, and efficient code organization.

​	清晰性主要通过有效的命名、有用的注释和高效的代码组织来实现。

Clarity is to be viewed through the lens of the reader, not the author of the code. It is more important that code be easy to read than easy to write. Clarity in code has two distinct facets:

​	清晰性应从读者的角度来看，而不是代码的作者。代码易于阅读比易于编写更重要。代码的清晰性有两个不同的方面：

- [What is the code actually doing?
  代码实际上在做什么？](https://google.github.io/styleguide/go/guide#clarity-purpose)
- [Why is the code doing what it does?
  代码为什么要做它所做的事情？](https://google.github.io/styleguide/go/guide#clarity-rationale)



#### What is the code actually doing? 代码实际上在做什么？

Go is designed such that it should be relatively straightforward to see what the code is doing. In cases of uncertainty or where a reader may require prior knowledge in order to understand the code, it is worth investing time in order to make the code’s purpose clearer for future readers. For example, it may help to:

​	Go 的设计使得人们可以相对直接地看到代码在做什么。在不确定或读者可能需要事先了解才能理解代码的情况下，值得花时间让代码的目的对未来的读者来说更清晰。例如，可能会有所帮助：

- Use more descriptive variable names
  使用更具描述性的变量名
- Add additional commentary
  添加额外的注释
- Break up the code with whitespace and comments
  用空格和注释分解代码
- Refactor the code into separate functions/methods to make it more modular
  将代码重构为单独的功能/方法，使其更具模块化

There is no one-size-fits-all approach here, but it is important to prioritize clarity when developing Go code.

​	这里没有一刀切的方法，但在开发 Go 代码时，优先考虑清晰度非常重要。



#### Why is the code doing what it does? 代码为什么要做它所做的事情？

The code’s rationale is often sufficiently communicated by the names of variables, functions, methods, or packages. Where it is not, it is important to add commentary. The “Why?” is especially important when the code contains nuances that a reader may not be familiar with, such as:

​	代码的原理通常由变量、函数、方法或包的名称充分传达。如果没有，添加注释非常重要。“为什么？”当代码包含读者可能不熟悉的细微差别时，这一点尤其重要，例如：

- A nuance in the language, e.g., a closure will be capturing a loop variable, but the closure is many lines away
  语言中的细微差别，例如，闭包将捕获循环变量，但闭包距离多行
- A nuance of the business logic, e.g., an access control check that needs to distinguish between the actual user and someone impersonating a user
  业务逻辑的细微差别，例如，需要区分实际用户和冒充用户的访问控制检查

An API might require care to use correctly. For example, a piece of code may be intricate and difficult to follow for performance reasons, or a complex sequence of mathematical operations may use type conversions in an unexpected way. In these cases and many more, it is important that accompanying commentary and documentation explain these aspects so that future maintainers don’t make a mistake and so that readers can understand the code without needing to reverse-engineer it.

​	API 可能需要小心使用才能正确使用。例如，一段代码可能错综复杂且难以遵循，原因是出于性能考虑，或者复杂的数学运算序列可能以意外的方式使用类型转换。在这些情况下以及更多情况下，重要的是随附的注释和文档要解释这些方面，以便未来的维护人员不会犯错，并且读者无需对其进行逆向工程即可理解代码。

It is also important to be aware that some attempts to provide clarity (such as adding extra commentary) can actually obscure the code’s purpose by adding clutter, restating what the code already says, contradicting the code, or adding maintenance burden to keep the comments up-to-date. Allow the code to speak for itself (e.g., by making the symbol names themselves self-describing) rather than adding redundant comments. It is often better for comments to explain why something is done, not what the code is doing.

​	同样重要的是要注意，一些提供清晰性的尝试（例如添加额外的注释）实际上可能会通过添加混乱、重复代码已经说过的内容、与代码相矛盾或增加维护负担以使注释保持最新状态来掩盖代码的目的。让代码自行说明（例如，通过使符号名称本身具有自描述性）而不是添加冗余注释。注释通常最好解释为什么做某事，而不是代码在做什么。

The Google codebase is largely uniform and consistent. It is often the case that code that stands out (e.g., by using an unfamiliar pattern) is doing so for a good reason, typically for performance. Maintaining this property is important to make it clear to readers where they should focus their attention when reading a new piece of code.

​	Google 代码库在很大程度上是统一且一致的。通常情况下，突出的代码（例如，通过使用不熟悉的模式）这样做是有充分理由的，通常是为了性能。维护此属性非常重要，以便在阅读新代码时向读者明确说明他们应该将注意力集中在哪里。

The standard library contains many examples of this principle in action. Among them:

​	标准库包含许多此原则的实际示例。其中包括：

- Maintainer comments in [`package sort`](https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/sort/sort.go).
  `package sort` 中的维护者注释。
- Good [runnable examples in the same package](https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/sort/example_search_test.go), which benefit both users (they [show up in godoc](https://pkg.go.dev/sort#pkg-examples)) and maintainers (they [run as part of tests](https://google.github.io/styleguide/go/decisions#examples)).
  同一个包中的良好可运行示例，这些示例对用户（它们显示在 godoc 中）和维护者（它们作为测试的一部分运行）都有好处。
- [`strings.Cut`](https://pkg.go.dev/strings#Cut) is only four lines of code, but they improve the [clarity and correctness of callsites](https://github.com/golang/go/issues/46336).
  `strings.Cut` 只有四行代码，但它们提高了调用站点的清晰性和正确性。



### Simplicity 简单性

Your Go code should be simple for those using, reading, and maintaining it.

​	您的 Go 代码对于使用、阅读和维护它的人来说应该很简单。

Go code should be written in the simplest way that accomplishes its goals, both in terms of behavior and performance. Within the Google Go codebase, simple code:

​	Go 代码应该以最简单的方式编写，以实现其目标，无论是在行为还是性能方面。在 Google Go 代码库中，简单的代码：

- Is easy to read from top to bottom
  从上到下都很容易阅读
- Does not assume that you already know what it is doing
  不假设您已经知道它在做什么
- Does not assume that you can memorize all of the preceding code
  不假设您可以记住所有前面的代码
- Does not have unnecessary levels of abstraction
  没有不必要的抽象级别
- Does not have names that call attention to something mundane
  没有引起对平凡事物注意的名称
- Makes the propagation of values and decisions clear to the reader
  使值和决策的传播对读者来说很清楚
- Has comments that explain why, not what, the code is doing to avoid future deviation
  有解释为什么而不是解释代码在做什么的注释，以避免未来的偏差
- Has documentation that stands on its own
  有独立的文档
- Has useful errors and useful test failures
  有有用的错误和有用的测试失败
- May often be mutually exclusive with “clever” code
  通常可能与“聪明”代码互斥

Tradeoffs can arise between code simplicity and API usage simplicity. For example, it may be worthwhile to have the code be more complex so that the end user of the API may more easily call the API correctly. In contrast, it may also be worthwhile to leave a bit of extra work to the end user of the API so that the code remains simple and easy to understand.

​	代码简单性和 API 使用简单性之间可能会出现权衡。例如，让代码更复杂可能是值得的，以便 API 的最终用户可以更轻松地正确调用 API。相比之下，让 API 的最终用户多做一些工作也可能是值得的，以便代码保持简单易懂。

When code needs complexity, the complexity should be added deliberately. This is typically necessary if additional performance is required or where there are multiple disparate customers of a particular library or service. Complexity may be justified, but it should come with accompanying documentation so that clients and future maintainers are able to understand and navigate the complexity. This should be supplemented with tests and examples that demonstrate its correct usage, especially if there is both a “simple” and a “complex” way to use the code.

​	当代码需要复杂性时，应故意添加复杂性。如果需要额外的性能或特定库或服务的客户众多，通常需要这样做。复杂性可能是合理的，但它应该附带相应的文档，以便客户和未来的维护人员能够理解和驾驭复杂性。这应该辅以测试和示例来演示其正确用法，尤其是在有“简单”和“复杂”两种使用代码的方法时。

This principle does not imply that complex code cannot or should not be written in Go or that Go code is not allowed to be complex. We strive for a codebase that avoids unnecessary complexity so that when complexity does appear, it indicates that the code in question requires care to understand and maintain. Ideally, there should be accompanying commentary that explains the rationale and identifies the care that should be taken. This often arises when optimizing code for performance; doing so often requires a more complex approach, like preallocating a buffer and reusing it throughout a goroutine lifetime. When a maintainer sees this, it should be a clue that the code in question is performance-critical, and that should influence the care that is taken when making future changes. If employed unnecessarily, on the other hand, this complexity is a burden on those who need to read or change the code in the future.

​	此原则并不意味着不能或不应该用 Go 编写复杂的代码，也不意味着 Go 代码不允许复杂。我们努力实现一个避免不必要复杂性的代码库，以便在出现复杂性时，它表明相关代码需要小心理解和维护。理想情况下，应该有附带的注释来解释基本原理并确定应采取的谨慎措施。这通常在优化代码以提高性能时出现；这样做通常需要更复杂的方法，例如预先分配缓冲区并在整个 goroutine 生命周期中重用它。当维护人员看到这一点时，这应该是一个线索，表明相关代码对性能至关重要，并且应该在进行未来更改时影响所采取的谨慎措施。另一方面，如果使用不必要，这种复杂性会给那些需要在未来阅读或更改代码的人带来负担。

If code turns out to be very complex when its purpose should be simple, this is often a signal to revisit the implementation to see if there is a simpler way to accomplish the same thing.

​	如果代码在目的应该很简单时变得非常复杂，这通常是一个重新审视实现的信号，以查看是否有更简单的方法来完成相同的事情。



#### Least mechanism 最少机制

Where there are several ways to express the same idea, prefer the one that uses the most standard tools. Sophisticated machinery often exists, but should not be employed without reason. It is easy to add complexity to code as needed, whereas it is much harder to remove existing complexity after it has been found to be unnecessary.

​	如果有多种方式表达同一个想法，请优先选择使用最标准的工具。通常存在复杂的机制，但不要无缘无故地使用。根据需要，很容易向代码中添加复杂性，而发现不必要后，删除现有复杂性则要困难得多。

1. Aim to use a core language construct (for example a channel, slice, map, loop, or struct) when sufficient for your use case.
   当足以满足您的用例时，请尝试使用核心语言结构（例如通道、切片、映射、循环或结构）。
2. If there isn’t one, look for a tool within the standard library (like an HTTP client or a template engine).
   如果没有，请在标准库中查找工具（例如 HTTP 客户端或模板引擎）。
3. Finally, consider whether there is a core library in the Google codebase that is sufficient before introducing a new dependency or creating your own.
   最后，在引入新依赖项或创建自己的依赖项之前，请考虑 Google 代码库中是否有足够的核心库。

As an example, consider production code that contains a flag bound to a variable with a default value which must be overridden in tests. Unless intending to test the program’s command-line interface itself (say, with `os/exec`), it is simpler and therefore preferable to override the bound value directly rather than by using `flag.Set`.

​	例如，考虑包含一个标志的生产代码，该标志绑定到具有默认值且必须在测试中覆盖的变量。除非打算测试程序的命令行界面本身（例如，使用 `os/exec` ），否则直接覆盖绑定值比使用 `flag.Set` 更简单，因此更好。

Similarly, if a piece of code requires a set membership check, a boolean-valued map (e.g., `map[string]bool`) often suffices. Libraries that provide set-like types and functionality should only be used if more complicated operations are required that are impossible or overly complicated with a map.

​	同样，如果一段代码需要进行集合成员资格检查，布尔值映射（例如， `map[string]bool` ）通常就足够了。只有在需要更复杂的运算（使用映射无法实现或过于复杂）时，才应该使用提供类集合类型和功能的库。



### Concision 简洁

Concise Go code has a high signal-to-noise ratio. It is easy to discern the relevant details, and the naming and structure guide the reader through these details.

​	简洁的 Go 代码具有很高的信噪比。很容易辨别相关细节，命名和结构引导读者了解这些细节。

There are many things that can get in the way of surfacing the most salient details at any given time:

​	在任何给定时间，都有许多事情会妨碍突出最突出的细节：

- Repetitive code
  重复的代码
- Extraneous syntax
  无关的语法
- [Opaque names
  不透明的名称](https://google.github.io/styleguide/go/guide#naming)
- Unnecessary abstraction
  不必要的抽象
- Whitespace
  空格

Repetitive code especially obscures the differences between each nearly-identical section, and requires a reader to visually compare similar lines of code to find the changes. [Table-driven testing](https://github.com/golang/go/wiki/TableDrivenTests) is a good example of a mechanism that can concisely factor out the common code from the important details of each repetition, but the choice of which pieces to include in the table will have an impact on how easy the table is to understand.

​	重复的代码尤其会模糊每个几乎相同的节之间的差异，并且要求读者对类似的代码行进行视觉比较以查找更改。表驱动测试是一个很好的例子，它可以简洁地将公共代码从每次重复的重要细节中分解出来，但选择将哪些部分包含在表中会影响表的易理解程度。

When considering multiple ways to structure code, it is worth considering which way makes important details the most apparent.

​	在考虑多种构建代码的方式时，值得考虑哪种方式使重要细节最明显。

Understanding and using common code constructions and idioms are also important for maintaining a high signal-to-noise ratio. For example, the following code block is very common in [error handling](https://go.dev/blog/errors-are-values), and the reader can quickly understand the purpose of this block.

​	理解和使用常见的代码结构和习惯用法对于保持高信噪比也很重要。例如，以下代码块在错误处理中非常常见，读者可以快速理解此块的目的。

``` go
// Good:
if err := doSomething(); err != nil {
    // ...
}
```

If code looks very similar to this but is subtly different, a reader may not notice the change. In cases like this, it is worth intentionally [“boosting”](https://google.github.io/styleguide/go/best-practices#signal-boost) the signal of the error check by adding a comment to call attention to it.

​	如果代码看起来与此非常相似，但略有不同，读者可能不会注意到更改。在这种情况下，值得通过添加注释来有意“增强”错误检查的信号，以引起人们的注意。

``` go
// Good:
if err := doSomething(); err == nil { // if NO error
    // ...
}
```



### Maintainability 可维护性

Code is edited many more times than it is written. Readable code not only makes sense to a reader who is trying to understand how it works, but also to the programmer who needs to change it. Clarity is key.

​	代码的编辑次数远多于编写次数。可读的代码不仅对试图理解其工作原理的读者有意义，而且对需要更改它的程序员也有意义。清晰是关键。

Maintainable code:

​	可维护的代码：

- Is easy for a future programmer to modify correctly
  未来程序员可以轻松正确地修改
- Has APIs that are structured so that they can grow gracefully
  具有可以优雅增长的结构的 API
- Is clear about the assumptions that it makes and chooses abstractions that map to the structure of the problem, not to the structure of the code
  清楚地说明其做出的假设，并选择映射到问题结构（而不是代码结构）的抽象
- Avoids unnecessary coupling and doesn’t include features that are not used
  避免不必要的耦合，不包含未使用的功能
- Has a comprehensive test suite to ensure promised behaviors are maintained and important logic is correct, and the tests provide clear, actionable diagnostics in case of failure
  拥有一个全面的测试套件，以确保承诺的行为得到维护，重要的逻辑是正确的，并且测试在发生故障时提供清晰、可操作的诊断 在使用抽象（例如接口和类型）时，这些抽象根据定义从其使用环境中删除信息，因此务必确保它们提供足够的好处。编辑器和 IDE 可以直接连接到方法定义，并在使用具体类型时显示相应的文档，但只能引用接口定义。接口是一个强大的工具，但需要付出代价，因为维护人员可能需要了解底层实现的具体信息才能正确使用接口，而这些信息必须在接口文档或调用位置进行说明。 可维护的代码还避免在容易忽略的地方隐藏重要细节。例如，在以下每一行代码中，单个字符的存在或缺失对于理解该行至关重要：

When using abstractions like interfaces and types which by definition remove information from the context in which they are used, it is important to ensure that they provide sufficient benefit. Editors and IDEs can connect directly to a method definition and show the corresponding documentation when a concrete type is used, but can only refer to an interface definition otherwise. Interfaces are a powerful tool, but come with a cost, since the maintainer may need to understand the specifics of the underlying implementation in order to correctly use the interface, which must be explained within the interface documentation or at the call-site.

Maintainable code also avoids hiding important details in places that are easy to overlook. For example, in each of the following lines of code, the presence or lack of a single character is critical to understand the line:

``` go
// Bad:
// The use of = instead of := can change this line completely.
if user, err = db.UserByID(userID); err != nil {
    // ...
}
// Bad:
// The ! in the middle of this line is very easy to miss.
leap := (year%4 == 0) && (!(year%100 == 0) || (year%400 == 0))
```

Neither of these are incorrect, but both could be written in a more explicit fashion, or could have an accompanying comment that calls attention to the important behavior:

​	这些都不错误，但两者都可以写得更明确，或者可以附带一条注释来引起对重要行为的注意：

``` go
// Good:
u, err := db.UserByID(userID)
if err != nil {
    return fmt.Errorf("invalid origin user: %s", err)
}
user = u
// Good:
// Gregorian leap years aren't just year%4 == 0.
// See https://en.wikipedia.org/wiki/Leap_year#Algorithm.
var (
    leap4   = year%4 == 0
    leap100 = year%100 == 0
    leap400 = year%400 == 0
)
leap := leap4 && (!leap100 || leap400)
```

In the same way, a helper function that hides critical logic or an important edge-case could make it easy for a future change to fail to account for it properly.

​	同样，隐藏关键逻辑或重要边缘情况的辅助函数可能会使未来的更改无法正确解释它。

Predictable names are another feature of maintainable code. A user of a package or a maintainer of a piece of code should be able to predict the name of a variable, method, or function in a given context. Function parameters and receiver names for identical concepts should typically share the same name, both to keep documentation understandable and to facilitate refactoring code with minimal overhead.

​	可预测的名称是可维护代码的另一个特征。软件包的用户或代码维护者应该能够预测给定上下文中变量、方法或函数的名称。对于相同概念的功能参数和接收者名称通常应共享相同的名称，既可以保持文档的可理解性，又可以以最小的开销促进重构代码。

Maintainable code minimizes its dependencies (both implicit and explicit). Depending on fewer packages means fewer lines of code that can affect behavior. Avoiding dependencies on internal or undocumented behavior makes code less likely to impose a maintenance burden when those behaviors change in the future.

​	可维护的代码最大程度地减少其依赖项（隐式和显式）。依赖较少的软件包意味着可以影响行为的代码行数更少。避免依赖内部或未记录的行为可降低代码在将来这些行为发生更改时造成维护负担的可能性。

When considering how to structure or write code, it is worth taking the time to think through ways in which the code may evolve over time. If a given approach is more conducive to easier and safer future changes, that is often a good trade-off, even if it means a slightly more complicated design.

​	在考虑如何构建或编写代码时，值得花时间思考代码可能随着时间推移而演变的方式。如果给定方法更有利于更轻松、更安全的未来更改，那么这通常是一个不错的权衡，即使这意味着设计稍微复杂一些。



### Consistency 一致性

Consistent code is code that looks, feels, and behaves like similar code throughout the broader codebase, within the context of a team or package, and even within a single file.

​	一致的代码是指在团队或包的上下文中，甚至在单个文件中，看起来、感觉和行为都像整个代码库中类似代码的代码。

Consistency concerns do not override any of the principles above, but if a tie must be broken, it is often beneficial to break it in favor of consistency.

​	一致性问题不会凌驾于上述任何原则之上，但如果必须打破僵局，那么通常有利于打破僵局以支持一致性。

Consistency within a package is often the most immediately important level of consistency. It can be very jarring if the same problem is approached in multiple ways throughout a package, or if the same concept has many names within a file. However, even this should not override documented style principles or global consistency.

​	包内的一致性通常是最直接重要的级别的一致性。如果在整个包中以多种方式解决相同的问题，或者如果同一个概念在文件中有多个名称，那么可能会非常刺耳。但是，即使这样也不应凌驾于记录的样式原则或全局一致性之上。



## Core guidelines 核心准则

These guidelines collect the most important aspects of Go style that all Go code is expected to follow. We expect that these principles be learned and followed by the time readability is granted. These are not expected to change frequently, and new additions will have to clear a high bar.

​	这些准则收集了所有 Go 代码都应遵循的 Go 风格的最重要方面。我们希望在授予可读性时学习并遵循这些原则。这些原则预计不会经常更改，并且新添加的内容必须清除高标准。

The guidelines below expand on the recommendations in [Effective Go](https://go.dev/doc/effective_go), which provide a common baseline for Go code across the entire community.

​	以下准则扩展了 Effective Go 中的建议，这些建议为整个社区的 Go 代码提供了共同的基准。



### Formatting 格式化

All Go source files must conform to the format outputted by the `gofmt` tool. This format is enforced by a presubmit check in the Google codebase. [Generated code](https://docs.bazel.build/versions/main/be/general.html#genrule) should generally also be formatted (e.g., by using [`format.Source`](https://pkg.go.dev/go/format#Source)), as it is also browsable in Code Search.

​	所有 Go 源文件都必须符合 `gofmt` 工具输出的格式。此格式由 Google 代码库中的预提交检查强制执行。生成的代码通常也应该格式化（例如，通过使用 `format.Source` ），因为它们也可以在代码搜索中浏览。



### MixedCaps

Go source code uses `MixedCaps` or `mixedCaps` (camel case) rather than underscores (snake case) when writing multi-word names.

​	在编写多词名称时，Go 源代码使用 `MixedCaps` 或 `mixedCaps` （小驼峰式）而不是下划线（蛇形式）。

This applies even when it breaks conventions in other languages. For example, a constant is `MaxLength` (not `MAX_LENGTH`) if exported and `maxLength` (not `max_length`) if unexported.

​	即使它打破了其他语言中的惯例，也适用。例如，如果导出，则常量为 `MaxLength` （不是 `MAX_LENGTH` ），如果未导出，则为 `maxLength` （不是 `max_length` ）。

Local variables are considered [unexported](https://go.dev/ref/spec#Exported_identifiers) for the purpose of choosing the initial capitalization.

​	在选择初始大写字母时，将局部变量视为未导出。



### Line length 行长

There is no fixed line length for Go source code. If a line feels too long, it should be refactored instead of broken. If it is already as short as it is practical for it to be, the line should be allowed to remain long.

​	Go 源代码没有固定的行长。如果一行感觉太长，应该重构而不是拆分。如果它已经尽可能短了，那么应该允许该行保持较长。

Do not split a line:

​	不要拆分一行：

- Before an [indentation change](https://google.github.io/styleguide/go/decisions#indentation-confusion) (e.g., function declaration, conditional)
  在缩进更改之前（例如，函数声明、条件）
- To make a long string (e.g., a URL) fit into multiple shorter lines
  为了使长字符串（例如 URL）适合多行较短的行



### Naming 命名

Naming is more art than science. In Go, names tend to be somewhat shorter than in many other languages, but the same [general guidelines](https://testing.googleblog.com/2017/10/code-health-identifiernamingpostforworl.html) apply. Names should:

​	命名更多的是艺术而不是科学。在 Go 中，名称往往比许多其他语言中的名称短一些，但相同的通用准则适用。名称应该：

- Not feel [repetitive](https://google.github.io/styleguide/go/decisions#repetition) when they are used
  在使用时不会感到重复
- Take the context into consideration
  考虑上下文
- Not repeat concepts that are already clear
  不重复已经很清楚的概念

You can find more specific guidance on naming in [decisions](https://google.github.io/styleguide/go/decisions#naming).

​	您可以在决策中找到有关命名的更具体指导。



### Local consistency 局部一致性

Where the style guide has nothing to say about a particular point of style, authors are free to choose the style that they prefer, unless the code in close proximity (usually within the same file or package, but sometimes within a team or project directory) has taken a consistent stance on the issue.

​	如果风格指南对某个特定风格点没有说明，作者可以自由选择他们喜欢的风格，除非邻近的代码（通常在同一个文件或包中，但有时在团队或项目目录中）对该问题采取了一致的立场。

Examples of **valid** local style considerations:

​	有效局部风格考虑因素的示例：

- Use of `%s` or `%v` for formatted printing of errors
  使用 `%s` 或 `%v` 格式化打印错误
- Usage of buffered channels in lieu of mutexes
  使用缓冲通道代替互斥锁

Examples of **invalid** local style considerations:

​	无效局部风格考虑因素的示例：

- Line length restrictions for code
  代码的行长限制
- Use of assertion-based testing libraries
  使用基于断言的测试库

If the local style disagrees with the style guide but the readability impact is limited to one file, it will generally be surfaced in a code review for which a consistent fix would be outside the scope of the CL in question. At that point, it is appropriate to file a bug to track the fix.

​	如果局部风格与风格指南不一致，但可读性影响仅限于一个文件，通常会在代码审查中发现，而一致的修复将超出相关 CL 的范围。此时，提交一个 bug 来跟踪修复是合适的。

If a change would worsen an existing style deviation, expose it in more API surfaces, expand the number of files in which the deviation is present, or introduce an actual bug, then local consistency is no longer a valid justification for violating the style guide for new code. In these cases, it is appropriate for the author to clean up the existing codebase in the same CL, perform a refactor in advance of the current CL, or find an alternative that at least does not make the local problem worse.

​	如果更改会使现有的样式偏差变得更糟，在更多 API 表面中公开它，扩展偏差存在的文件的数量，或引入实际错误，那么局部一致性不再是违反新代码样式指南的有效理由。在这些情况下，作者可以在同一个 CL 中清理现有的代码库，在当前 CL 之前执行重构，或找到至少不会使局部问题变得更糟的替代方案，这是合适的。