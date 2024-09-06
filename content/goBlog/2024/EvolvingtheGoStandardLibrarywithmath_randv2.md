+++
title = "使用 math/rand/v2 演变 Go 标准库"
date = 2024-05-30T10:14:27+08:00
weight = 940
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/randv2](https://go.dev/blog/randv2)

# Evolving the Go Standard Library with math/rand/v2 - 使用 math/rand/v2 演变 Go 标准库

Russ Cox
1 May 2024

​	2024 年 5 月 1 日

Since Go 1 was [released in March 2012](https://go.dev/blog/go1), changes to the standard library have been constrained by Go’s [compatibility promise](https://go.dev/doc/go1compat). Overall, compatibility has been a boon for Go users, providing a stable base for production systems, documentation, tutorials, books, and more. Over time, however, we’ve realized mistakes in the original APIs that cannot be fixed compatibly; in other cases, best practices and convention have changed. We need a plan for making important, breaking changes too.

​	自 2012 年 3 月发布 Go 1 以来，对标准库的更改受到 Go 兼容性承诺的约束。总体而言，兼容性一直是 Go 用户的福音，为生产系统、文档、教程、书籍等提供了稳定的基础。然而，随着时间的推移，我们意识到原始 API 中无法兼容修复的错误；在其他情况下，最佳实践和惯例已经改变。我们需要一个计划来进行重要的重大更改。

This blog post is about Go 1.22’s new [`math/rand/v2`](https://go.dev/pkg/math/rand/v2/) package, the first “v2” in the standard library. It brings needed improvements to the [`math/rand`](https://go.dev/pkg/math/rand/) API, but more importantly it sets an example for how we can revise other standard library packages as the need arises.

​	这篇博文是关于 Go 1.22 的新 `math/rand/v2` 包，这是标准库中的第一个“v2”。它对 `math/rand` API 进行了必要的改进，但更重要的是，它为我们如何根据需要修改其他标准库包树立了一个榜样。

(In Go, `math/rand` and `math/rand/v2` are two different packages with different import paths. Go 1 and every release after it have included `math/rand`; Go 1.22 added `math/rand/v2`. A Go program can import either package, or both.)

​	（在 Go 中， `math/rand` 和 `math/rand/v2` 是两个不同的包，具有不同的导入路径。Go 1 及其之后的每个版本都包含 `math/rand` ；Go 1.22 添加了 `math/rand/v2` 。Go 程序可以导入任一包，或同时导入两个包。）

This post discusses the specific rationale for the changes in `math/rand/v2` and then [reflects on the general principles](https://go.dev/blog/randv2#principles) that will guide new versions of other packages.

​	本文讨论了 `math/rand/v2` 中更改的具体原理，然后反映了将指导其他软件包新版本的通用原则。

## Pseudorandom Number Generators 伪随机数生成器

Before we look at `math/rand`, which is an API for a pseudorandom number generator, let’s take a moment to understand what that means.

​	在我们了解 `math/rand` 之前，这是一个伪随机数生成器的 API，让我们花点时间来理解一下这意味着什么。

A pseudorandom number generator is a deterministic program that generates a long sequence of seemingly random numbers from a small seed input, although the numbers are not in fact random at all. In the case of `math/rand`, the seed is a single int64, and the algorithm produces a sequence of int64s using a variant of a [linear-feedback shift register (LFSR)](https://en.wikipedia.org/wiki/Linear-feedback_shift_register). The algorithm is based on an idea by George Marsaglia, tweaked by Don Mitchell and Jim Reeds, and further customized by Ken Thompson for Plan 9 and then Go. It has no official name, so this post calls it the Go 1 generator.

​	伪随机数生成器是一个确定性程序，它从一个小的种子输入生成一个看似随机数的长序列，尽管这些数实际上根本不是随机的。在 `math/rand` 的情况下，种子是一个 int64，该算法使用线性反馈移位寄存器 (LFSR) 的一个变体生成一个 int64 序列。该算法基于 George Marsaglia 的一个想法，由 Don Mitchell 和 Jim Reeds 调整，然后由 Ken Thompson 为 Plan 9 和 Go 进一步定制。它没有官方名称，因此这篇文章称它为 Go 1 生成器。

The goal is for these generators to be fast, repeatable, and random enough to support simulations, shuffling, and other non-cryptographic use cases. Repeatability is particularly important for uses like numerical simulations or randomized testing. For example, a randomized tester might pick a seed (perhaps based on the current time), generate a large random test input, and repeat. When the tester finds a failure, it only needs to print the seed to allow repeating the test with that specific large input.

​	目标是让这些生成器快速、可重复且随机性足够高，以支持模拟、洗牌和其他非加密用例。可重复性对于数值模拟或随机测试等用途尤为重要。例如，随机测试器可能会选择一个种子（可能基于当前时间），生成一个大型随机测试输入，然后重复。当测试器发现故障时，它只需要打印种子即可重复使用该特定的大型输入进行测试。

Repeatability also matters over time: given a particular seed, a new version of Go needs to generate the same sequence of values that an older version did. We didn’t realize this when we released Go 1; instead, we discovered it the hard way, when we tried to make a change in Go 1.2 and got reports that we had broken certain tests and other use cases. At that point, we decided Go 1 compatibility included the specific random outputs for a given seed and [added a test](https://go.dev/change/5aca0514941ce7dd0f3cea8d8ffe627dbcd542ca).

​	可重复性在一段时间内也很重要：给定一个特定的种子，新版本的 Go 需要生成与旧版本相同的数值序列。我们在发布 Go 1 时没有意识到这一点；相反，我们以困难的方式发现了这一点，当我们尝试在 Go 1.2 中进行更改时，并收到报告称我们破坏了某些测试和其他用例。在这一点上，我们决定 Go 1 兼容性包括给定种子的特定随机输出，并添加了一个测试。

It is not a goal for these kinds of generators to produce random numbers suitable for deriving cryptographic keys or other important secrets. Because the seed is only 63 bits, any output drawn from the generator, no matter how long, will also only contain 63 bits of entropy. For example, using `math/rand` to generate a 128-bit or 256-bit AES key would be a serious mistake, since the key would be easier to brute force. For that kind of use, you need a cryptographically strong random number generator, as provided by [`crypto/rand`](https://go.dev/pkg/crypto/rand/).

​	对于这些类型的生成器来说，生成适合于派生加密密钥或其他重要秘密的随机数并不是目标。由于种子只有 63 位，因此从生成器中提取的任何输出，无论多长，都只会包含 63 位熵。例如，使用 `math/rand` 生成 128 位或 256 位 AES 密钥将是一个严重的错误，因为该密钥更容易被暴力破解。对于这种用途，您需要一个密码学上强随机数生成器，如 `crypto/rand` 所提供的。

That’s enough background that we can move on to what needed fixing in the `math/rand` package.

​	有了足够的背景知识，我们可以继续讨论 `math/rand` 包中需要修复的内容。

## Problems with `math/rand` `math/rand` 的问题

Over time, we noticed more and more problems with `math/rand`. The most serious were the following.

​	随着时间的推移，我们注意到 `math/rand` 出现了越来越多的问题。最严重的问题如下。

### Generator Algorithm 生成器算法

The generator itself needed replacement.

​	发电机本身需要更换。

The initial implementation of Go, while production ready, was in many ways a “pencil sketch” of the entire system, working well enough to serve as a base for future development: the compiler and runtime were written in C; the garbage collector was a conservative, single-threaded, stop-the-world collector; and the libraries used basic implementations throughout. From Go 1 through around Go 1.5, we went back and drew the “fully inked” version of each of these: we converted the compiler and runtime to Go; we wrote a new, precise, parallel, concurrent garbage collection with microsecond pause times; and we replaced standard library implementations with more sophisticated, optimized algorithms as needed.

​	Go 的初始实现虽然已准备好投入生产，但在很多方面是整个系统的“铅笔素描”，足以作为未来开发的基础：编译器和运行时是用 C 编写的；垃圾收集器是一个保守的、单线程的、停止世界的收集器；并且库在整个过程中使用了基本实现。从 Go 1 到 Go 1.5 左右，我们返回并绘制了每个实现的“完全着墨”版本：我们将编译器和运行时转换为 Go；我们编写了一个新的、精确的、并行的、并发垃圾收集，暂停时间为微秒；并且根据需要用更复杂、优化的算法替换了标准库实现。

Unfortunately, the repeatability requirement in `math/rand` meant that we couldn’t replace the generator there without breaking compatibility. We were stuck with the Go 1 generator, which is reasonably fast (about 1.8ns per number on my M3 Mac) but maintains an internal state of almost 5 kilobytes. In contrast, Melissa O’Neill’s [PCG family of generators](https://www.pcg-random.org/) generates better random numbers in about 2.1ns per number with only 16 bytes of internal state. We also wanted to explore using Daniel J. Bernstein’s [ChaCha stream cipher](https://cr.yp.to/chacha.html) as a generator. A [follow-up post](https://go.dev/blog/chacha8rand) discusses that generator specifically.

​	不幸的是， `math/rand` 中的可重复性要求意味着我们无法在不破坏兼容性的情况下替换那里的生成器。我们只能使用 Go 1 生成器，它相当快（在我的 M3 Mac 上每个数字约 1.8ns），但保持着近 5 千字节的内部状态。相比之下，Melissa O’Neill 的 PCG 生成器系列以每个数字约 2.1ns 的速度生成更好的随机数，内部状态只有 16 字节。我们还想探索使用 Daniel J. Bernstein 的 ChaCha 流密码作为生成器。后续文章专门讨论了该生成器。

### Source Interface 源接口

The [`rand.Source` interface](https://go.dev/pkg/math/rand/#Source) was wrong. That interface defines the concept of a low-level random number generator that generates non-negative `int64` values:

​	 `rand.Source` 接口是错误的。该接口定义了一个低级随机数生成器的概念，该生成器生成非负 `int64` 值：

```
% go doc -src math/rand.Source
package rand // import "math/rand"

// A Source represents a source of uniformly-distributed
// pseudo-random int64 values in the range [0, 1<<63).
//
// A Source is not safe for concurrent use by multiple goroutines.
type Source interface {
    Int63() int64
    Seed(seed int64)
}

func NewSource(seed int64) Source
%
```

(In the doc comment, “[0, N)” denotes a [half-open interval](https://en.wikipedia.org/wiki/Interval_(mathematics)#Definitions_and_terminology), meaning the range includes 0 but ends just before 2⁶³.)

​	（在文档注释中，“[0, N)”表示半开区间，这意味着该范围包括 0，但结束于 2⁶³ 之前。）

The [`rand.Rand` type](https://go.dev/pkg/math/rand/#Rand) wraps a `Source` to implement a richer set of operations, such as generating [an integer between 0 and N](https://go.dev/pkg/math/rand/#Rand.Intn), generating [floating-point numbers](https://go.dev/pkg/math/rand/#Rand.Float64), and so on.

​	 `rand.Rand` 类型包装 `Source` 以实现更丰富的操作集，例如生成 0 到 N 之间的整数、生成浮点数等等。

We defined the `Source` interface to return a shortened 63-bit value instead of a uint64 because that’s what the Go 1 generator and other widely-used generators produce, and it matches the convention set by the C standard library. But this was a mistake: more modern generators produce full-width uint64s, which is a more convenient interface.

​	我们定义了 `Source` 接口以返回一个缩短的 63 位值，而不是 uint64，因为这是 Go 1 生成器和其他广泛使用的生成器产生的，并且它与 C 标准库设置的约定相匹配。但这却是一个错误：更现代的生成器会生成全宽 uint64，这是一个更方便的接口。

Another problem is the `Seed` method hard-coding an `int64` seed: some generators are seeded by larger values, and the interface provides no way to handle that.

​	另一个问题是 `Seed` 方法硬编码了一个 `int64` 种子：一些生成器由更大的值播种，而该接口没有提供处理该值的方法。

### Seeding Responsibility 播种责任

A bigger problem with `Seed` is that responsibility for seeding the global generator was unclear. Most users don’t use `Source` and `Rand` directly. Instead, the `math/rand` package provides a global generator accessed by top-level functions like [`Intn`](https://go.dev/pkg/math/rand/#Intn). Following the C standard library, the global generator defaults to behaving as if `Seed(1)` is called at startup. This is good for repeatability but bad for programs that want their random outputs to be different from one run to the next. The package documentation suggests using `rand.Seed(time.Now().UnixNano())` in that case, to make the generator’s output time-dependent, but what code should do this?

​	 `Seed` 的一个更大问题是，播种全局生成器的责任不明确。大多数用户不会直接使用 `Source` 和 `Rand` 。相反， `math/rand` 包提供了一个全局生成器，可通过顶级函数（如 `Intn` ）访问。遵循 C 标准库，全局生成器默认为在启动时调用 `Seed(1)` 。这有利于可重复性，但对于希望其随机输出在每次运行中都不同的程序来说却不好。在这种情况下，包文档建议使用 `rand.Seed(time.Now().UnixNano())` ，以使生成器的输出依赖于时间，但应该使用什么代码来执行此操作？

Probably the main package should be in charge of how `math/rand` is seeded: it would be unfortunate for imported libraries to configure global state themselves, since their choices might conflict with other libraries or the main package. But what happens if a library needs some random data and wants to use `math/rand`? What if the main package doesn’t even know `math/rand` is being used? We found that in practice many libraries add init functions that seed the global generator with the current time, “just to be sure”.

​	可能主包应该负责 `math/rand` 的种子方式：如果导入的库自行配置全局状态，那将很不幸，因为它们的选择可能会与其他库或主包冲突。但是，如果某个库需要一些随机数据并希望使用 `math/rand` ，会发生什么情况？如果主包甚至不知道 `math/rand` 正在被使用，会发生什么情况？我们发现，在实践中，许多库都会添加 init 函数，使用当前时间为全局生成器设置种子，“以确保”。

Library packages seeding the global generator themselves causes a new problem. Suppose package main imports two packages that both use `math/rand`: package A assumes the global generator will be seeded by package main, but package B seeds it in an `init` func. And suppose that package main doesn’t seed the generator itself. Now package A’s correct operation depends on the coincidence that package B is also imported in the program. If package main stops importing package B, package A will stop getting random values. We observed this happening in practice in large codebases.

​	库包自己播种全局生成器会导致一个新问题。假设包 main 导入两个都使用 `math/rand` 的包：包 A 假设全局生成器将由包 main 播种，但包 B 在 `init` 函数中播种它。并且假设包 main 本身不播种生成器。现在，包 A 的正确操作取决于包 B 也在程序中导入的巧合。如果包 main 停止导入包 B，包 A 将停止获取随机值。我们观察到这种情况在大型代码库中实际发生。

In retrospect, it was clearly a mistake to follow the C standard library here: seeding the global generator automatically would remove the confusion about who seeds it, and users would stop being surprised by repeatable output when they didn’t want that.

​	回顾过去，很明显按照 C 标准库来做是一个错误：自动播种全局生成器将消除关于谁播种它的困惑，并且用户将不再对他们不想要时可重复的输出感到惊讶。

### Scalability 可扩展性

The global generator also did not scale well. Because top-level functions like [`rand.Intn`](https://go.dev/pkg/math/rand/#Intn) can be called simultaneously from multiple goroutines, the implementation needed a lock protecting the shared generator state. In parallel usage, acquiring and releasing this lock was more expensive than the actual generation. It would make sense instead to have a per-thread generator state, but doing so would break repeatability in programs without concurrent use of `math/rand`.

​	全局生成器也没有很好地扩展。因为像 `rand.Intn` 这样的顶级函数可以从多个 goroutine 同时调用，所以实现需要一个保护共享生成器状态的锁。在并行使用中，获取和释放此锁比实际生成更昂贵。相反，拥有每个线程生成器状态是有意义的，但这样做会破坏在不并发使用 `math/rand` 的程序中的可重复性。

### The `Rand` implementation was missing important optimizations `Rand` 实现缺少重要的优化

The [`rand.Rand` type](https://go.dev/pkg/math/rand/#Rand) wraps a `Source` to implement a richer set of operations For example, here is the Go 1 implementation of `Int63n`, which returns a random integer in the range [0, `n`).

​	 `rand.Rand` 类型包装 `Source` 以实现更丰富的操作集。例如，以下是 `Int63n` 的 Go 1 实现，它返回 [0, `n` ) 范围内的随机整数。

```go
func (r *Rand) Int63n(n int64) int64 {
    if n <= 0 {
        panic("invalid argument to Int63n")
    }
    max := int64((1<<63 - 1)  - (1<<63)%uint64(n))
    v := r.src.Int63()
    for v > max {
        v = r.Int63()
    }
    return v % n
}
```

The actual conversion is easy: `v % n`. However, no algorithm can convert 2⁶³ equally likely values into `n` equally likely values unless 2⁶³ is a multiple of `n`: otherwise some outputs will necessarily happen more often than others. (As a simpler example, try converting 4 equally likely values into 3.) The code computes `max` such that `max+1` is the largest multiple of `n` less than or equal to 2⁶³, and then the loop rejects random values greater than or equal to `max+1`. Rejecting this too-large values ensures that all `n` outputs are equally likely. For small `n`, needing to reject any value at all is rare; rejection becomes more common and more important for larger values. Even without the rejection loop, the two (slow) modulus operations can make the conversion more expensive than generating the random value `v` in the first place.

​	实际转换很简单： `v % n` 。但是，除非 2⁶³ 是 `n` 的倍数，否则没有算法可以将 2⁶³ 个等可能的数值转换为 `n` 个等可能的数值：否则，某些输出必然比其他输出更频繁。（作为一个更简单的示例，尝试将 4 个等可能的数值转换为 3 个。）该代码计算 `max` ，使得 `max+1` 是小于或等于 2⁶³ 的 `n` 的最大倍数，然后循环拒绝大于或等于 `max+1` 的随机值。拒绝这些过大的值可确保所有 `n` 输出都等可能。对于较小的 `n` ，根本不需要拒绝任何值的情况很少见；对于较大的值，拒绝变得更加常见和更加重要。即使没有拒绝循环，两个（慢）模运算也会使转换比最初生成随机值 `v` 更加昂贵。

In 2018, [Daniel Lemire found an algorithm](https://arxiv.org/abs/1805.10941) that avoids the divisions nearly all the time (see also his [2019 blog post](https://lemire.me/blog/2019/06/06/nearly-divisionless-random-integer-generation-on-various-systems/)). In `math/rand`, adopting Lemire’s algorithm would make `Intn(1000)` 20-30% faster, but we can’t: the faster algorithm generates different values than the standard conversion, breaking repeatability.

​	2018 年，Daniel Lemire 发现了一种几乎始终避免除法的算法（另请参阅他 2019 年的博客文章）。在 `math/rand` 中，采用 Lemire 的算法将使 `Intn(1000)` 快 20-30%，但我们不能这样做：更快的算法会生成与标准转换不同的值，从而破坏可重复性。

Other methods are also slower than they could be, constrained by repeatability. For example, the `Float64` method could easily be sped up by about 10% if we could change the generated value stream. (This was the change we tried to make in Go 1.2 and rolled back, mentioned earlier.)

​	其他方法也比它们可能的速度慢，受制于可重复性。例如，如果我们能够更改生成的值流， `Float64` 方法很容易就能加快大约 10%。 （这是我们尝试在 Go 1.2 中进行的更改，并回滚，前面已提到。）

### The `Read` Mistake `Read` 的错误

As mentioned earlier, `math/rand` is not intended for and not suitable for generating cryptographic secrets. The `crypto/rand` package does that, and its fundamental primitive is its [`Read` function](https://go.dev/pkg/crypto/rand/#Read) and [`Reader`](https://go.dev/pkg/crypto/rand/#Reader) variable.

​	如前所述， `math/rand` 不适用于生成加密密钥，也不适合用于此目的。 `crypto/rand` 包可用于此目的，其基本原语是 `Read` 函数和 `Reader` 变量。

In 2015, we accepted a proposal to make `rand.Rand` implement `io.Reader` as well, along with [adding a top-level `Read` function](https://go.dev/pkg/math/rand/#Read). This seemed reasonable at the time, but in retrospect we did not pay enough attention to the software engineering aspects of this change. Now if you want to read random data, now you have two choices: `math/rand.Read` and `crypto/rand.Read`. If the data is going to be used for key material, it is very important to use `crypto/rand`, but now it is possible to use `math/rand` instead, potentially with disastrous consequences.

​	2015 年，我们接受了一项提案，让 `rand.Rand` 也实现 `io.Reader` ，同时添加一个顶级 `Read` 函数。当时这似乎很合理，但回想起来，我们没有足够重视此更改的软件工程方面。现在，如果您想读取随机数据，您有两个选择： `math/rand.Read` 和 `crypto/rand.Read` 。如果数据将用于密钥材料，则使用 `crypto/rand` 非常重要，但现在可以使用 `math/rand` 代替，这可能会产生灾难性的后果。

Tools like `goimports` and `gopls` have a special case to make sure they prefer to use `rand.Read` from `crypto/rand` instead of `math/rand`, but that’s not a complete fix. It would be better to remove `Read` entirely.

​	像 `goimports` 和 `gopls` 这样的工具有一个特殊情况，以确保它们更喜欢使用 `crypto/rand` 中的 `rand.Read` 而不是 `math/rand` ，但这并不是一个完整的修复。最好完全删除 `Read` 。

## Fixing `math/rand` directly 直接修复 `math/rand`

Making a new, incompatible major version of a package is never our first choice: that new version only benefits programs that switch to it, leaving all existing usage of the old major version behind. In contrast, fixing a problem in the existing package has much more impact, since it fixes all the existing usage. We should never create a `v2` without doing as much as possible to fix `v1`. In the case of `math/rand`, we were able to partly address a few of the problems described above:

​	制作一个新的、不兼容的软件包主要版本永远不是我们的首选：该新版本仅使切换到它的程序受益，而将旧主要版本的所有现有用法都抛在脑后。相比之下，修复现有软件包中的问题影响更大，因为它修复了所有现有用法。我们永远不应该在不尽力修复 `v1` 的情况下创建一个 `v2` 。在 `math/rand` 的情况下，我们能够部分解决上面描述的几个问题：

- Go 1.8 introduced an optional [`Source64` interface](https://go.dev/pkg/math/rand/#Uint64) with a `Uint64` method. If a `Source` also implements `Source64`, then `Rand` uses that method when appropriate. This “extension interface” pattern provides a compatible (if slightly awkward) way to revise an interface after the fact.

  ​	Go 1.8 引入了一个可选的 `Source64` 接口，其中包含一个 `Uint64` 方法。如果 `Source` 也实现了 `Source64` ，那么 `Rand` 会在适当的时候使用该方法。这种“扩展接口”模式提供了一种兼容的（尽管有点别扭）方式来事后修改接口。

- Go 1.20 automatically seeded the top-level generator and deprecated [`rand.Seed`](https://go.dev/pkg/math/rand/#Seed). Although this may seem like an incompatible change given our focus on repeatability of the output stream, [we reasoned](https://go.dev/issue/56319) that any imported package that called [`rand.Int`](https://go.dev/pkg/math/rand/#Int) at init time or inside any computation would also visibly change the output stream, and surely adding or removing such a call cannot be considered a breaking change. And if that’s true, then auto-seeding is no worse, and it would eliminate this source of fragility for future programs. We also added a [GODEBUG setting](https://go.dev/doc/godebug) to opt back into the old behavior. Then we marked the top-level `rand.Seed` as [deprecated](https://go.dev/wiki/Deprecated). (Programs that need seeded repeatability can still use `rand.New(rand.NewSource(seed))` to obtain a local generator instead of using the global one.)

  ​	Go 1.20 自动播种顶级生成器并弃用 `rand.Seed` 。尽管鉴于我们专注于输出流的可重复性，这看起来像是不兼容的更改，但我们认为，任何在初始化时或在任何计算中调用 `rand.Int` 的导入包也会明显改变输出流，并且肯定添加或删除这样的调用不能被视为重大更改。如果这是真的，那么自动播种不会更糟，并且它将消除未来程序的这种脆弱性来源。我们还添加了一个 GODEBUG 设置以选择恢复旧行为。然后，我们将顶级 `rand.Seed` 标记为已弃用。（需要播种可重复性的程序仍然可以使用 `rand.New(rand.NewSource(seed))` 来获取本地生成器，而不是使用全局生成器。）

- Having eliminated repeatability of the global output stream, Go 1.20 was also able to make the global generator scale better in programs that don’t call `rand.Seed`, replacing the Go 1 generator with a very cheap per-thread [wyrand generator](https://github.com/wangyi-fudan/wyhash) already used inside the Go runtime. This removed the global mutex and made the top-level functions scale much better. Programs that do call `rand.Seed` fall back to the mutex-protected Go 1 generator.

  ​	通过消除全局输出流的可重复性，Go 1.20 还能够在不调用 `rand.Seed` 的程序中更好地扩展全局生成器，使用 Go 运行时中已使用的非常便宜的每线程 wyrand 生成器替换 Go 1 生成器。这移除了全局互斥锁，并使顶级函数扩展得更好。确实调用 `rand.Seed` 的程序会回退到受互斥锁保护的 Go 1 生成器。

- We were able to adopt Lemire’s optimization in the Go runtime, and we also used it inside [`rand.Shuffle`](https://go.dev/pkg/math/rand/#Shuffle), which was implemented after Lemire’s paper was published.

  ​	我们能够在 Go 运行时采用 Lemire 的优化，并且我们还在 `rand.Shuffle` 中使用了它，它是在 Lemire 的论文发表后实现的。

- Although we couldn’t remove [`rand.Read`](https://go.dev/pkg/math/rand/#Read) entirely, Go 1.20 marked it [deprecated](https://go.dev/wiki/Deprecated) in favor of `crypto/rand`. We have since heard from people who discovered that they were accidentally using `math/rand.Read` in cryptographic contexts when their editors flagged the use of the deprecated function.

  ​	尽管我们无法完全移除 `rand.Read` ，但 Go 1.20 已将其标记为弃用，转而支持 `crypto/rand` 。我们已经听取了那些发现自己意外地在加密上下文中使用 `math/rand.Read` 的人的意见，因为他们的编辑器标记了弃用函数的使用。

These fixes are imperfect and incomplete but also real improvements that helped all users of the existing `math/rand` package. For more complete fixes, we needed to turn our attention to `math/rand/v2`.

​	这些修复不完美且不完整，但也是帮助现有 `math/rand` 包的所有用户的真正改进。对于更完整的修复，我们需要将注意力转向 `math/rand/v2` 。

## Fixing the rest in `math/rand/v2` 修复 `math/rand/v2` 中的其余部分

Defining `math/rand/v2` took significant planning, then a [GitHub Discussion](https://go.dev/issue/60751) and then a [proposal discussion](https://go.dev/issue/61716). It is the same as `math/rand` with the following breaking changes addressing the problems outlined above:

​	定义 `math/rand/v2` 需要大量的规划，然后是 GitHub 讨论，然后是提案讨论。它与 `math/rand` 相同，其中包含以下针对上述问题进行解决的重大更改：

- We removed the Go 1 generator entirely, replacing it with two new generators, [PCG](https://go.dev/pkg/math/rand/v2/#PCG) and [ChaCha8](https://go.dev/pkg/math/rand/v2/#ChaCha8). The new types are named for their algorithms (avoiding the generic name `NewSource`) so that if another important algorithm needs to be added, it will fit well into the naming scheme.

  ​	我们完全移除了 Go 1 生成器，用两个新生成器 PCG 和 ChaCha8 取代了它。新类型以其算法命名（避免使用通用名称 `NewSource` ），这样如果需要添加另一个重要算法，它将很好地融入命名方案中。

  Adopting a suggestion from the proposal discussion, the new types implement the [`encoding.BinaryMarshaler`](https://go.dev/pkg/encoding/#BinaryMarshaler) and [`encoding.BinaryUnmarshaler`](https://go.dev/pkg/encoding/#BinaryUnmarshaler) interfaces.

  ​	采纳提案讨论中的建议，新类型实现了 `encoding.BinaryMarshaler` 和 `encoding.BinaryUnmarshaler` 接口。

- We changed the `Source` interface, replacing the `Int63` method with a `Uint64` method and deleting the `Seed` method. Implementations that support seeding can provide their own concrete methods, like [`PCG.Seed`](https://go.dev/pkg/math/rand/v2/#PCG.Seed) and [`ChaCha8.Seed`](https://go.dev/pkg/math/rand/v2/#ChaCha8.Seed). Note that the two take different seed types, and neither is a single `int64`.

  ​	我们更改了 `Source` 接口，用 `Uint64` 方法替换了 `Int63` 方法，并删除了 `Seed` 方法。支持播种的实现可以提供它们自己的具体方法，如 `PCG.Seed` 和 `ChaCha8.Seed` 。请注意，这两个方法采用不同的种子类型，并且都不是单个 `int64` 。

- We removed the top-level `Seed` function: the global functions like `Int` can only be used in auto-seeded form now.

  ​	我们移除了顶级 `Seed` 函数：现在只能以自动播种形式使用 `Int` 等全局函数。

- Removing the top-level `Seed` also let us hard-code the use of scalable, per-thread generators by the top-level methods, avoiding a GODEBUG check at each use.

  ​	移除顶级 `Seed` 也让我们能够硬编码顶级方法使用可扩展的、每个线程的生成器，避免在每次使用时进行 GODEBUG 检查。

- We implemented Lemire’s optimization for `Intn` and related functions. The concrete `rand.Rand` API is now locked in to that value stream, so we will not be able to take advantage of any optimizations yet to be discovered, but at least we are up to date once again. We also implemented the `Float32` and `Float64` optimizations we wanted to use back in Go 1.2.

  ​	我们为 `Intn` 和相关函数实现了 Lemire 的优化。具体的 `rand.Rand` API 现在已锁定到该值流中，因此我们无法利用尚未发现的任何优化，但至少我们再次更新了。我们还实现了我们希望在 Go 1.2 中使用的 `Float32` 和 `Float64` 优化。

- During the proposal discussion, a contributor pointed out detectable bias in the implementations of `ExpFloat64` and `NormFloat64`. We fixed that bias and locked in the new value streams.

  ​	在提案讨论期间，一位贡献者指出了 `ExpFloat64` 和 `NormFloat64` 实现中的可检测偏差。我们修复了该偏差并锁定了新的值流。

- `Perm` and `Shuffle` used different shuffling algorithms and produced different value streams, because `Shuffle` happened second and used a faster algorithm. Deleting `Perm` entirely would have made migration harder for users. Instead we implemented `Perm` in terms of `Shuffle`, which still lets us delete an implementation.

  ​	 `Perm` 和 `Shuffle` 使用了不同的洗牌算法并生成了不同的值流，因为 `Shuffle` 发生在第二位并使用了更快的算法。完全删除 `Perm` 会让用户更难迁移。相反，我们根据 `Shuffle` 来实现 `Perm` ，这仍然允许我们删除一个实现。

- We renamed `Int31`, `Int63`, `Intn`, `Int31n`, and `Int63n` to `Int32`, `Int64`, `IntN`, `Int32N`, and `Int64N`. The 31 and 63 in the names were unnecessarily pedantic and confusing, and the capitalized N is more idiomatic for a second “word” in the name in Go.

  ​	我们把 `Int31` 、 `Int63` 、 `Intn` 、 `Int31n` 和 `Int63n` 重命名为 `Int32` 、 `Int64` 、 `IntN` 、 `Int32N` 和 `Int64N` 。名称中的 31 和 63 不必要地繁琐且令人困惑，而大写的 N 对于 Go 中名称中的第二个“单词”来说更符合惯例。

- We added `Uint`, `Uint32`, `Uint64`, `UintN`, `Uint32N`, and `Uint64N` top-level functions and methods. We needed to add `Uint64` to provide direct access to the core `Source` functionality, and it seemed inconsistent not to add the others.

  ​	我们添加了 `Uint` 、 `Uint32` 、 `Uint64` 、 `UintN` 、 `Uint32N` 和 `Uint64N` 顶级函数和方法。我们需要添加 `Uint64` 以提供对核心 `Source` 功能的直接访问，并且不添加其他功能似乎不一致。

- Adopting another suggestion from the proposal discussion, we added a new top-level, generic function `N` that is like `Int64N` or `Uint64N` but works for any integer type. In the old API, to create a random duration of up to 5 seconds, it was necessary to write:

  ​	采纳提案讨论中的另一条建议，我们添加了一个新的顶级通用函数 `N` ，它类似于 `Int64N` 或 `Uint64N` ，但适用于任何整数类型。在旧 API 中，要创建一个最长为 5 秒的随机持续时间，必须编写：

  ```go
  d := time.Duration(rand.Int63n(int64(5*time.Second)))
  ```

  Using `N`, the equivalent code is:

  ​	使用 `N` ，等效代码为：

  ```go
  d := rand.N(5 * time.Second)
  ```

  `N` is only a top-level function; there is no `N` method on `rand.Rand` because there are no generic methods in Go. (Generic methods are not likely in the future, either; they conflict badly with interfaces, and a complete implementation would require either run-time code generation or slow execution.)

  ​	 `N` 只是一个顶级函数； `rand.Rand` 上没有 `N` 方法，因为 Go 中没有泛型方法。（将来也不太可能出现泛型方法；它们与接口严重冲突，并且完整的实现需要运行时代码生成或缓慢执行。）

- To ameliorate misuse of `math/rand` in cryptographic contexts, we made `ChaCha8` the default generator used in global functions, and we also changed the Go runtime to use it (replacing wyrand). Programs are still strongly encouraged to use `crypto/rand` to generate cryptographic secrets, but accidentally using `math/rand/v2` is not as catastrophic as using `math/rand` would be. Even in `math/rand`, the global functions now use the `ChaCha8` generator when not explicitly seeded.

  ​	为了改善在加密上下文中对 `math/rand` 的滥用，我们让 `ChaCha8` 成为全局函数中使用的默认生成器，并且我们还更改了 Go 运行时以使用它（替换 wyrand）。强烈建议程序仍然使用 `crypto/rand` 来生成加密密钥，但意外使用 `math/rand/v2` 并不像使用 `math/rand` 那样具有灾难性。即使在 `math/rand` 中，当没有明确设置种子时，全局函数现在也使用 `ChaCha8` 生成器。

## Principles for evolving the Go standard library Go 标准库演进原则

As mentioned at the start this post, one of the goals for this work was to establish principles and a pattern for how we approach all v2 packages in the standard library. There will not be a glut of v2 packages in the next few Go releases. Instead, we will handle one package at a time, making sure we set a quality bar that will last for another decade. Many packages will not need a v2 at all. But for those that do, our approach boils down to three principles.

​	正如本文开头所述，这项工作的目标之一是为我们如何处理标准库中的所有 v2 包制定原则和模式。在接下来的几个 Go 版本中不会出现过多的 v2 包。相反，我们将一次处理一个包，确保我们设定一个将在未来十年内持续存在的质量标准。许多包根本不需要 v2。但对于那些需要的包，我们的方法归结为三个原则。

First, a new, incompatible version of a package will use `that/package/v2` as its import path, following [semantic import versioning](https://research.swtch.com/vgo-import) just like a v2 module outside the standard library would. This allows uses of the original package and the v2 package to coexist in a single program, which is critical for a [gradual conversion](https://go.dev/talks/2016/refactor.article) to the new API.

​	首先，一个新版本的不兼容包将使用 `that/package/v2` 作为其导入路径，遵循语义导入版本控制，就像标准库之外的 v2 模块一样。这允许原始包和 v2 包在单个程序中共存，这对于逐步转换到新 API 至关重要。

Second, all changes must be rooted in respect for existing usage and users: we must not introduce needless churn, whether in the form of unnecessary changes to an existing package or an entirely new package that must be learned instead. In practice, that means we take the existing package as the starting point and only make changes that are well motivated and provide a value that justifies the cost to users of updating.

​	其次，所有更改都必须植根于对现有用法和用户的尊重：我们不能引入不必要的变动，无论是对现有软件包进行不必要的更改，还是必须学习的全新软件包。在实践中，这意味着我们将现有软件包作为起点，并且仅进行有充分动机且提供价值的更改，以证明更新对用户的成本是合理的。

Third, the v2 package must not leave v1 users behind. Ideally, the v2 package should be able to do everything the v1 package could do, and when v2 is released, the v1 package should be rewritten to be a thin wrapper around v2. This would ensure that existing uses of v1 continue to benefit from bug fixes and performance optimizations in v2. Of course, given that v2 is introducing breaking changes, this is not always possible, but it is always something to consider carefully. For `math/rand/v2`, we arranged for the auto-seeded v1 functions to call the v2 generator, but we were unable to share other code due to the repeatability violations. Ultimately `math/rand` is not a lot of code and does not require regular maintenance, so the duplication is manageable. In other contexts, more work to avoid duplication could be worthwhile. For example, in the [encoding/json/v2 design (still in progress)](https://go.dev/issue/63397), although the default semantics and the API are changed, the package provides configuration knobs that make it possible to implement the v1 API. When we eventually ship `encoding/json/v2`, `encoding/json` (v1) will become a thin wrapper around it, ensuring that users who don’t migrate from v1 still benefit from optimizations and security fixes in v2.

​	第三，v2 包不能抛弃 v1 用户。理想情况下，v2 包应该能够完成 v1 包可以完成的所有事情，并且当 v2 发布时，v1 包应该被重写为 v2 周围的一个薄包装器。这将确保 v1 的现有用途继续受益于 v2 中的错误修复和性能优化。当然，鉴于 v2 引入了重大更改，这并不总是可能的，但始终需要仔细考虑。对于 `math/rand/v2` ，我们安排自动播种的 v1 函数调用 v2 生成器，但由于可重复性违规，我们无法共享其他代码。最终 `math/rand` 不是很多代码，并且不需要定期维护，因此重复是可以管理的。在其他情况下，避免重复的更多工作可能是值得的。例如，在 encoding/json/v2 设计（仍在进行中）中，尽管默认语义和 API 已更改，但该包提供了配置旋钮，使实现 v1 API 成为可能。 当我们最终发布 `encoding/json/v2` 时， `encoding/json` (v1) 将成为其一个轻量级封装，确保未从 v1 迁移的用户仍能从 v2 中的优化和安全修复中受益。

A [follow-up blog post](https://go.dev/blog/chacha8rand) presents the `ChaCha8` generator in more detail.

​	后续博客文章更详细地介绍了 `ChaCha8` 生成器。