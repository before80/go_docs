+++
title = "Go 1.22 中的安全随机性"
date = 2024-05-30T10:14:57+08:00
weight = 930
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/chacha8rand](https://go.dev/blog/chacha8rand)

# Secure Randomness in Go 1.22 - Go 1.22 中的安全随机性

Russ Cox and Filippo Valsorda

​	Russ Cox 和 Filippo Valsorda
2 May 2024

​	2024 年 5 月 2 日

Computers aren’t random. On the contrary, hardware designers work very hard to make sure computers run every program the same way every time. So when a program does need random numbers, that requires extra effort. Traditionally, computer scientists and programming languages have distinguished between two different kinds of random numbers: statistical and cryptographic randomness. In Go, those are provided by [`math/rand`](https://go.dev/pkg/math/rand/) and [`crypto/rand`](https://go.dev/pkg/crypto/rand), respectively. This post is about how Go 1.22 brings the two closer together, by using a cryptographic random number source in `math/rand` (as well as `math/rand/v2`, as mentioned in our [previous post](https://go.dev/blog/randv2)). The result is better randomness and far less damage when developers accidentally use `math/rand` instead of `crypto/rand`.

​	计算机并非随机的。相反，硬件设计人员非常努力地确保计算机每次都以相同的方式运行每个程序。因此，当程序确实需要随机数时，这就需要额外的努力。传统上，计算机科学家和编程语言区分了两种不同类型的随机数：统计随机性和密码随机性。在 Go 中，它们分别由 `math/rand` 和 `crypto/rand` 提供。这篇文章讲述了 Go 1.22 如何通过在 `math/rand` 中使用密码随机数源（以及 `math/rand/v2` ，如我们在上一篇文章中提到的）将两者更紧密地结合在一起。结果是更好的随机性，并且当开发人员意外使用 `math/rand` 而不是 `crypto/rand` 时，造成的损害要小得多。

Before we can explain what Go 1.22 did, let’s take a closer look at statistical randomness compared to cryptographic randomness.

​	在我们解释 Go 1.22 做了什么之前，让我们仔细看看统计随机性与加密随机性之间的比较。

## Statistical Randomness 统计随机性

Random numbers that pass basic statistical tests are usually appropriate for use cases like simulations, sampling, numerical analysis, non-cryptographic randomized algorithms, [random testing](https://go.dev/doc/security/fuzz/), [shuffling inputs](https://en.wikipedia.org/wiki/Fisher–Yates_shuffle), and [random exponential backoff](https://en.wikipedia.org/wiki/Exponential_backoff#Collision_avoidance). Very basic, easy to compute mathematical formulas turn out to work well enough for these use cases. Because the methods are so simple, however, an observer who knows what algorithm is being used can typically predict the rest of the sequence after seeing enough values.

​	通过基本统计检验的随机数通常适用于模拟、抽样、数值分析、非加密随机算法、随机测试、输入洗牌和随机指数退避等用例。非常基本、易于计算的数学公式足以满足这些用例。然而，由于这些方法非常简单，因此了解所用算法的观察者通常可以在看到足够多的值后预测序列的其余部分。

Essentially all programming environments provide a mechanism for generating statistical random numbers that traces back through C to Research Unix Third Edition (V3), which added a pair of functions: `srand` and `rand`. The manual page included a note that read:

​	本质上，所有编程环境都提供了一种生成统计随机数的机制，该机制可通过 C 追溯到 Research Unix 第三版 (V3)，其中添加了一对函数： `srand` 和 `rand` 。手册页中包含一条注释，内容如下：

> *WARNING  The author of this routine has been writing random-number generators for many years and has never been known to write one that worked.
> 警告 本例程的作者编写随机数生成器已有许多年，但从未听说过他编写过一个可行的生成器。*

This note was partly a joke but also an acknowledgement that such generators are [inherently not random](https://www.tuhs.org/pipermail/tuhs/2024-March/029587.html).

​	这个注释部分是玩笑，但也承认这样的生成器本质上不是随机的。

The source code of the generator makes clear how trivial it is. Translated from PDP-11 assembly to modern C, it was:

​	生成器的源代码清楚地表明了它是多么微不足道。从 PDP-11 汇编语言翻译成现代 C 语言，它就是：

```
uint16 ranx;

void
srand(uint16 seed)
{
    ranx = seed;
}

int16
rand(void)
{
    ranx = 13077*ranx + 6925;
    return ranx & ~0x8000;
}
```

Calling `srand` seeds the generator with a single integer seed, and `rand` returns the next number from the generator. The AND in the return statement clears the sign bit to make sure the result is positive.

​	调用 `srand` 用一个整数种子播种生成器，而 `rand` 返回生成器的下一个数字。返回语句中的 AND 清除符号位以确保结果为正。

This function is an instance of the general class of [linear congruential generators (LCGs)](https://en.wikipedia.org/wiki/Linear_congruential_generator), which Knuth analyzes in *The Art of Computer Programming*, Volume 2, section 3.2.1. The main benefit of LCGs is that constants can be chosen such that they emit every possible output value once before repeating, as the Unix implementation did for 15-bit outputs. A serious problem with LCGs, however, is that the high bits of the state do not affect the low bits at all, so every truncation of the sequence to *k* bits necessarily repeats with a smaller period. The low bit must toggle: 0, 1, 0, 1, 0, 1. The low two bits must count up or down: 0, 1, 2, 3, 0, 1, 2, 3, or else 0, 3, 2, 1, 0, 3, 2, 1. There are four possible three-bit sequences; the original Unix implementation repeats 0, 5, 6, 3, 4, 1, 2, 7. (These problems can be avoided by reducing the value modulo a prime, but that would have been quite expensive at the time. See S. K. Park and K. W. Miller’s 1988 CACM paper “[Random number generators: good ones are hard to find](https://dl.acm.org/doi/10.1145/63039.63042)” for a short analysis and the first chapter of Knuth Volume 2 for a longer one.)

​	此函数是线性同余生成器 (LCG) 的一般类的实例，Knuth 在《计算机编程的艺术》第 2 卷第 3.2.1 节中对此进行了分析。LCG 的主要好处是可以选择常数，以便在重复之前发出每个可能的输出值一次，就像 Unix 实现对 15 位输出所做的那样。然而，LCG 的一个严重问题是状态的高位根本不影响低位，因此序列对 k 位的任何截断必然以较小的周期重复。低位必须切换：0、1、0、1、0、1。低两位必须向上或向下计数：0、1、2、3、0、1、2、3，或者 0、3、2、1、0、3、2、1。有四种可能的三位序列；最初的 Unix 实现重复 0、5、6、3、4、1、2、7。（可以通过将值模减一个质数来避免这些问题，但当时这样做会非常昂贵。请参阅 S. K. Park 和 K. W. Miller 在 1988 年发表的 CACM 论文“随机数生成器：很难找到好的”，以进行简要分析，以及 Knuth 第 2 卷的第一章以进行更长的分析。）

Even with these known problems, the `srand` and `rand` functions were included in the first C standard, and equivalent functionality was included in essentially every language since then. LCGs were once the dominant implementation strategy, although they’ve fallen off in popularity due to some important drawbacks. One significant remaining use is [`java.util.Random`](https://github.com/openjdk/jdk8u-dev/blob/master/jdk/src/share/classes/java/util/Random.java), which powers [`java.lang.Math.random`](https://github.com/openjdk/jdk8u-dev/blob/master/jdk/src/share/classes/java/util/Random.java).

​	即使存在这些已知问题， `srand` 和 `rand` 函数也被包含在第一个 C 标准中，并且从那时起，几乎每种语言中都包含了等效的功能。LCG 曾经是主要的实现策略，尽管由于一些重要的缺点，它们的受欢迎程度有所下降。一个仍然重要的用途是 `java.util.Random` ，它为 `java.lang.Math.random` 提供支持。

Another thing you can see from the implementation above is that the internal state is completely exposed by the result of `rand`. An observer who knows the algorithm and sees a single result can easily compute all future results. If you are running a server that calculates some random values that become public and some random values that must stay secret, using this kind of generator would be disastrous: the secrets wouldn’t be secret.

​	从上面的实现中，你还可以看到，内部状态完全由 `rand` 的结果公开。一个了解算法并看到单个结果的观察者可以轻松计算出所有未来的结果。如果你正在运行一个服务器，它计算一些公开的随机值和一些必须保密的随机值，那么使用这种生成器将是灾难性的：秘密将不再是秘密。

More modern random generators aren’t as terrible as the original Unix one, but they’re still not completely unpredictable. To make that point, next we will look at the original `math/rand` generator from Go 1 and the PCG generator we added in `math/rand/v2`.

​	更现代的随机生成器不像最初的 Unix 那样糟糕，但它们仍然不是完全不可预测的。为了说明这一点，接下来我们将查看 Go 1 中的原始 `math/rand` 生成器和我们在 `math/rand/v2` 中添加的 PCG 生成器。

## The Go 1 Generator Go 1 生成器

The generator used in Go 1’s `math/rand` is an instance of what is called a [linear-feedback shift register](https://en.wikipedia.org/wiki/Linear-feedback_shift_register). The algorithm is based on an idea by George Marsaglia, tweaked by Don Mitchell and Jim Reeds, and further customized by Ken Thompson for Plan 9 and then Go. It has no official name, so this post calls it the Go 1 generator.

​	Go 1 中使用的生成器 `math/rand` 是一个线性反馈移位寄存器的实例。该算法基于 George Marsaglia 的一个想法，由 Don Mitchell 和 Jim Reeds 调整，然后由 Ken Thompson 为 Plan 9 和 Go 进一步定制。它没有官方名称，因此这篇文章称它为 Go 1 生成器。

The Go 1 generator’s internal state is a slice `vec` of 607 uint64s. In that slice, there are two distinguished elements: `vec[606]`, the last element, is called the “tap”, and `vec[334]` is called the “feed”. To generate the next random number, the generator adds the tap and the feed to produce a value `x`, stores `x` back into the feed, shifts the entire slice one position to the right (the tap moves to `vec[0]` and `vec[i]` moves to `vec[i+1]`), and returns `x`. The generator is called “linear feedback” because the tap is *added* to the feed; the entire state is a “shift register” because each step shifts the slice entries.

​	Go 1 生成器的内部状态是 607 个 uint64 的切片 `vec` 。在该切片中，有两个不同的元素： `vec[606]` ，最后一个元素，称为“抽头”， `vec[334]` 称为“馈送”。为了生成下一个随机数，生成器将抽头和馈送相加以生成一个值 `x` ，将 `x` 存储回馈送，将整个切片向右移动一个位置（抽头移动到 `vec[0]` ， `vec[i]` 移动到 `vec[i+1]` ），并返回 `x` 。生成器被称为“线性反馈”，因为抽头被添加到馈送中；整个状态是一个“移位寄存器”，因为每一步都会移位切片条目。

Of course, actually moving every slice entry forward would be prohibitively expensive, so instead the implementation leaves the slice data in place and moves the tap and feed positions backward on each step. The code looks like:

​	当然，实际上向前移动每个切片条目将非常昂贵，因此，实现将切片数据保留在原处，并在每一步上向后移动抽头和进给位置。代码如下所示：

```
func (r *rngSource) Uint64() uint64 {
    r.tap--
    if r.tap < 0 {
        r.tap += len(r.vec)
    }

    r.feed--
    if r.feed < 0 {
        r.feed += len(r.vec)
    }

    x := r.vec[r.feed] + r.vec[r.tap]
    r.vec[r.feed] = x
    return uint64(x)
}
```

Generating the next number is quite cheap: two subtractions, two conditional adds, two loads, one add, one store.

​	生成下一个数字非常便宜：两个减法，两个条件加法，两个加载，一个加法，一个存储。

Unfortunately, because the generator directly returns one slice element from its internal state vector, reading 607 values from the generator completely exposes all its state. With those values, you can predict all the future values, by filling in your own `vec` and then running the algorithm. You can also recover all the previous values, by running the algorithm backward (subtracting the tap from the feed and shifting the slice to the left).

​	不幸的是，由于生成器直接从其内部状态向量返回一个切片元素，从生成器读取 607 个值会完全公开其所有状态。有了这些值，您可以通过填写您自己的 `vec` 然后运行算法来预测所有未来值。您还可以通过反向运行算法（从进料中减去抽头并向左移动切片）来恢复所有先前值。

As a complete demonstration, here is an [insecure program](https://go.dev/play/p/v0QdGjUAtzC) generating pseudorandom authentication tokens along with code that predicts the next token given a sequence of earlier tokens. As you can see, the Go 1 generator provides no security at all (nor was it meant to). The quality of the generated numbers also depends on the initial setting of `vec`.

​	作为一个完整的演示，这里有一个不安全的程序，它生成伪随机身份验证令牌，以及给定一系列早期令牌预测下一个令牌的代码。正如你所看到的，Go 1 生成器根本不提供任何安全性（它也不是为了提供安全性而设计的）。生成数字的质量还取决于 `vec` 的初始设置。

## The PCG Generator PCG 生成器

For `math/rand/v2`, we wanted to provide a more modern statistical random generator and settled on Melissa O’Neill’s PCG algorithm, published in 2014 in her paper “[PCG: A Family of Simple Fast Space-Efficient Statistically Good Algorithms for Random Number Generation](https://www.pcg-random.org/pdf/hmc-cs-2014-0905.pdf)”. The exhaustive analysis in the paper can make it hard to notice at first glance how utterly trivial the generators are: PCG is a post-processed 128-bit LCG.

​	对于 `math/rand/v2` ，我们希望提供一个更现代的统计随机生成器，并确定了 Melissa O’Neill 在 2014 年发表的论文“PCG：用于随机数生成的简单、快速、空间高效且统计良好的算法系列”中的 PCG 算法。论文中详尽的分析可能让人一开始很难注意到生成器的简单性：PCG 是经过后处理的 128 位 LCG。

If the state `p.x` were a `uint128` (hypothetically), the code to compute the next value would be:

​	如果状态 `p.x` 是一个 `uint128` （假设），计算下一个值的代码将是：

```
const (
    pcgM = 0x2360ed051fc65da44385df649fccf645
    pcgA = 0x5851f42d4c957f2d14057b7ef767814f
)

type PCG struct {
    x uint128
}

func (p *PCG) Uint64() uint64 {
    p.x = p.x * pcgM + pcgA
    return scramble(p.x)
}
```

The entire state is a single 128-bit number, and the update is a 128-bit multiply and add. In the return statement, the `scramble` function reduces the 128-bit state down to a 64-bit state. The original PCG used (again using a hypothetical `uint128` type):

​	整个状态是一个 128 位数字，更新是一个 128 位乘法和加法。在 return 语句中， `scramble` 函数将 128 位状态缩减为 64 位状态。最初的 PCG 使用（再次使用一个假设的 `uint128` 类型）：

```
func scramble(x uint128) uint64 {
    return bits.RotateLeft(uint64(x>>64) ^ uint64(x), -int(x>>122))
}
```

This code XORs the two halves of the 128-bit state together and then rotates the result according to the top six bits of the state. This version is called PCG-XSL-RR, for “xor shift low, right rotate”.

​	此代码对 128 位状态的两半进行异或运算，然后根据状态的前六位对结果进行旋转。此版本称为 PCG-XSL-RR，表示“异或低位移，右旋转”。

Based on a [suggestion from O’Neill during proposal discussion](https://go.dev/issue/21835#issuecomment-739065688), Go’s PCG uses a new scramble function based on multiplication, which mixes the bits more aggressively:

​	根据 O'Neill 在提案讨论期间的建议，Go 的 PCG 使用基于乘法的新的扰动函数，它更激进地混合位：

```
func scramble(x uint128) uint64 {
    hi, lo := uint64(x>>64), uint64(x)
    hi ^= hi >> 32
    hi *= 0xda942042e4dd58b5
    hi ^= hi >> 48
    hi *= lo | 1
}
```

O’Neill calls PCG with this scrambler PCG-DXSM, for “double xorshift multiply.” Numpy uses this form of PCG as well.

​	O’Neill 使用这个名为 PCG-DXSM 的扰频器调用 PCG，表示“双 xorshift 乘法”。Numpy 也使用这种形式的 PCG。

Although PCG uses more computation to generate each value, it uses significantly less state: two uint64s instead of 607. It is also much less sensitive to the initial values of that state, and [it passes many statistical tests that other generators do not](https://www.pcg-random.org/statistical-tests.html). In many ways it is an ideal statistical generator.

​	尽管 PCG 使用更多计算来生成每个值，但它使用的状态明显更少：两个 uint64，而不是 607。它对该状态的初始值也不那么敏感，并且它通过了许多其他生成器无法通过的统计测试。在很多方面，它都是一个理想的统计生成器。

Even so, PCG is not unpredictable. While the scrambling of bits to prepare the result does not expose the state directly like in the LCG and Go 1 generators, [PCG-XSL-RR can still be be reversed](https://pdfs.semanticscholar.org/4c5e/4a263d92787850edd011d38521966751a179.pdf), and it would not be surprising if PCG-DXSM could too. For secrets, we need something different.

​	即便如此，PCG 并非不可预测。虽然为准备结果而对位元进行扰乱不会像 LCG 和 Go 1 生成器那样直接暴露状态，但 PCG-XSL-RR 仍然可以被逆转，而 PCG-DXSM 也可以被逆转也就不足为奇了。对于秘密，我们需要一些不同的东西。

## Cryptographic Randomness 密码随机性

*Cryptographic random numbers* need to be utterly unpredictable in practice, even to an observer who knows how they are generated and has observed any number of previously generated values. The safety of cryptographic protocols, secret keys, modern commerce, online privacy, and more all critically depend on access to cryptographic randomness.

​	密码随机数在实践中需要完全不可预测，即使对于知道它们是如何生成的并且已经观察到任何数量的先前生成值的人来说也是如此。密码协议、密钥、现代商业、在线隐私等的安全性都严重依赖于对密码随机性的访问。

Providing cryptographic randomness is ultimately the job of the operating system, which can gather true randomness from physical devices—timings of the mouse, keyboard, disks, and network, and more recently [electrical noise measured directly by the CPU itself](https://web.archive.org/web/20141230024150/http://www.cryptography.com/public/pdf/Intel_TRNG_Report_20120312.pdf). Once the operating system has gathered a meaningful amount of randomness—say, at least 256 bits—it can use cryptographic hashing or encryption algorithms to stretch that seed into an arbitrarily long sequence of random numbers. (In practice the operating system is also constantly gathering and adding new randomness to the sequence too.)

​	提供加密随机性最终是操作系统的工作，它可以从物理设备中收集真正的随机性——鼠标、键盘、磁盘和网络的时间，以及最近由 CPU 本身直接测量的电噪声。一旦操作系统收集到有意义的随机性（比如至少 256 位），它就可以使用加密哈希或加密算法将该种子扩展为任意长的随机数序列。（实际上，操作系统也在不断收集和向序列中添加新的随机性。）

The exact operating system interfaces have evolved over time. A decade ago, most systems provided a device file named `/dev/random` or something similar. Today, in recognition of how fundamental randomness has become, operating systems provide a direct system call instead. (This also allows programs to read randomness even when cut off from the file system.) In Go, the [`crypto/rand`](https://go.dev/pkg/crypto/rand/) package abstracts away those details, providing the same interface on every operating system: [`rand.Read`](https://go.dev/pkg/crypto/rand/#Read).

​	确切的操作系统接口随着时间推移而演变。十年前，大多数系统提供了一个名为 `/dev/random` 或类似名称的设备文件。如今，为了认识到随机性的基本性，操作系统直接提供系统调用。（这也允许程序在与文件系统断开连接时读取随机性。）在 Go 中， `crypto/rand` 包抽象了这些细节，在每个操作系统上提供相同的接口： `rand.Read` 。

It would not be practical for `math/rand` to ask the operating system for randomness each time it needs a `uint64`. But we can use cryptographic techniques to define an in-process random generator that improves on LCGs, the Go 1 generator, and even PCG.

​	对于 `math/rand` 来说，每次需要 `uint64` 时都向操作系统请求随机数是不切实际的。但是，我们可以使用加密技术来定义一个进程内随机生成器，它比 LCG、Go 1 生成器甚至 PCG 都有所改进。

## The ChaCha8Rand Generator ChaCha8Rand 生成器

Our new generator, which we unimaginatively named ChaCha8Rand for specification purposes and implemented as `math/rand/v2`’s [`rand.ChaCha8`](https://go.dev/pkg/math/rand/v2/#ChaCha8), is a lightly modified version of Daniel J. Bernstein’s [ChaCha stream cipher](https://cr.yp.to/chacha.html). ChaCha is widely used in a 20-round form called ChaCha20, including in TLS and SSH. Jean-Philippe Aumasson’s paper “[Too Much Crypto](https://eprint.iacr.org/2019/1492.pdf)” argues persuasively that the 8-round form ChaCha8 is secure too (and it’s roughly 2.5X faster). We used ChaCha8 as the core of ChaCha8Rand.

​	我们的新生成器，我们出于规范目的而毫不费力地将其命名为 ChaCha8Rand，并将其实现为 `math/rand/v2` 的 `rand.ChaCha8` ，是 Daniel J. Bernstein 的 ChaCha 流密码的轻微修改版本。ChaCha 广泛用于称为 ChaCha20 的 20 轮形式中，包括在 TLS 和 SSH 中。Jean-Philippe Aumasson 的论文“太多加密”有说服力地论证了 8 轮形式 ChaCha8 也安全（并且大约快 2.5 倍）。我们使用 ChaCha8 作为 ChaCha8Rand 的核心。

Most stream ciphers, including ChaCha8, work by defining a function that is given a key and a block number and produces a fixed-size block of apparently random data. The cryptographic standard these aim for (and usually meet) is for this output to be indistinguishable from actual random data in the absence of some kind of exponentially costly brute force search. A message is encrypted or decrypted by XOR’ing successive blocks of input data with successive randomly generated blocks. To use ChaCha8 as a `rand.Source`, we use the generated blocks directly instead of XOR’ing them with input data (this is equivalent to encrypting or decrypting all zeros).

​	大多数流密码（包括 ChaCha8）通过定义一个函数来工作，该函数给定一个密钥和一个块号，并生成一个固定大小的明显随机数据块。这些函数的目标（并且通常会达到）是加密标准，即在没有某种指数级代价的暴力搜索的情况下，此输出与实际随机数据无法区分。通过将连续的输入数据块与连续的随机生成块进行 XOR 操作，对消息进行加密或解密。若要将 ChaCha8 用作 `rand.Source` ，我们直接使用生成的块，而不是将它们与输入数据进行 XOR 操作（这相当于对所有零进行加密或解密）。

We changed a few details to make ChaCha8Rand more suitable for generating random numbers. Briefly:

​	我们更改了一些细节，以使 ChaCha8Rand 更适合生成随机数。简而言之：

- ChaCha8Rand takes a 32-byte seed, used as the ChaCha8 key.
  ChaCha8Rand 采用 32 字节种子，用作 ChaCha8 密钥。
- ChaCha8 generates 64-byte blocks, with calculations treating a block as 16 `uint32`s. A common implementation is to compute four blocks at a time using [SIMD instructions](https://en.wikipedia.org/wiki/Single_instruction,_multiple_data) on 16 vector registers of four `uint32`s each. This produces four interleaved blocks that must be unshuffled for XOR’ing with the input data. ChaCha8Rand defines that the interleaved blocks are the random data stream, removing the cost of the unshuffle. (For security purposes, this can be viewed as standard ChaCha8 followed by a reshuffle.)
  ChaCha8 生成 64 字节块，计算将块视为 16 `uint32` s。一种常见的实现是使用 SIMD 指令同时计算四个块，每个块有 16 个向量寄存器，每个向量寄存器有四个 `uint32` s。这会产生四个交错块，必须取消交错才能与输入数据进行 XOR。ChaCha8Rand 定义交错块是随机数据流，从而消除了取消交错的成本。（出于安全目的，这可以视为标准 ChaCha8，然后重新洗牌。）
- ChaCha8 finishes a block by adding certain values to each `uint32` in the block. Half the values are key material and the other half are known constants. ChaCha8Rand defines that the known constants are not re-added, removing half of the final adds. (For security purposes, this can be viewed as standard ChaCha8 followed by subtracting the known constants.)
  ChaCha8 通过向块中的每个 `uint32` 添加特定值来完成一个块。一半的值是密钥材料，另一半是已知的常量。ChaCha8Rand 定义已知的常量不会被重新添加，从而删除了一半的最终添加。（出于安全目的，这可以被视为标准 ChaCha8，然后减去已知的常量。）
- Every 16th generated block, ChaCha8Rand takes the final 32 bytes of the block for itself, making them the key for the next 16 blocks. This provides a kind of [forward secrecy](https://en.wikipedia.org/wiki/Forward_secrecy): if a system is compromised by an attack that recovers the entire memory state of the generator, only values generated since the last rekeying can be recovered. The past is inaccessible. ChaCha8Rand as defined so far must generate 4 blocks at a time, but we chose to do this key rotation every 16 blocks to leave open the possibility of faster implementations using 256-bit or 512-bit vectors, which could generate 8 or 16 blocks at a time.
  每生成 16 个区块，ChaCha8Rand 会为自己获取区块的最后 32 个字节，使其成为下一个 16 个区块的密钥。这提供了一种前向保密性：如果系统因攻击而遭到破坏，导致生成器的整个内存状态被恢复，则只能恢复自上次重新加锁以来生成的值。过去是不可访问的。到目前为止，定义的 ChaCha8Rand 必须一次生成 4 个区块，但我们选择每 16 个区块进行一次密钥轮换，以便为使用 256 位或 512 位向量的更快速实现留出可能性，这些实现一次可以生成 8 个或 16 个区块。

We wrote and published a [C2SP specification for ChaCha8Rand](https://c2sp.org/chacha8rand), along with test cases. This will enable other implementations to share repeatability with the Go implementation for a given seed.

​	我们为 ChaCha8Rand 编写并发布了 C2SP 规范，以及测试用例。这将使其他实现能够与 Go 实现共享给定种子的可重复性。

The Go runtime now maintains a per-core ChaCha8Rand state (300 bytes), seeded with operating system-supplied cryptographic randomness, so that random numbers can be generated quickly without any lock contention. Dedicating 300 bytes per core may sound expensive, but on a 16-core system, it is about the same as storing a single shared Go 1 generator state (4,872 bytes). The speed is worth the memory. This per-core ChaCha8Rand generator is now used in three different places in the Go standard library:

​	Go 运行时现在维护每个内核的 ChaCha8Rand 状态（300 字节），使用操作系统提供的加密随机数进行种子设置，以便在没有任何锁争用的情况下快速生成随机数。为每个内核分配 300 字节听起来可能很昂贵，但在 16 核系统上，它与存储单个共享 Go 1 生成器状态（4,872 字节）大致相同。速度值得内存。此每个内核的 ChaCha8Rand 生成器现在在 Go 标准库中的三个不同位置使用：

1. The `math/rand/v2` package functions, such as [`rand.Float64`](https://go.dev/pkg/math/rand/v2/#Float64) and [`rand.N`](https://go.dev/pkg/math/rand/v2/#N), always use ChaCha8Rand.

   ​	 `math/rand/v2` 包函数（如 `rand.Float64` 和 `rand.N` ）始终使用 ChaCha8Rand。

2. The `math/rand` package functions, such as [`rand.Float64`](https://go.dev/pkg/math/rand/#Float64) and [`rand.Intn`](https://go.dev/pkg/math/rand/#Intn), use ChaCha8Rand when [`rand.Seed`](https://go.dev/pkg/math/rand/#Seed) has not been called. Applying ChaCha8Rand in `math/rand` improves the security of programs even before they update to `math/rand/v2`, provided they are not calling `rand.Seed`. (If `rand.Seed` is called, the implementation is required to fall back to the Go 1 generator for compatibility.)

   ​	 `math/rand` 包函数（如 `rand.Float64` 和 `rand.Intn` ）在未调用 `rand.Seed` 时使用 ChaCha8Rand。在 `math/rand` 中应用 ChaCha8Rand 可提高程序的安全性，即使在程序更新到 `math/rand/v2` 之前也是如此，前提是它们不调用 `rand.Seed` 。（如果调用 `rand.Seed` ，则实现需要回退到 Go 1 生成器以实现兼容性。）

3. The runtime chooses the hash seed for each new map using ChaCha8Rand instead of a less secure [wyrand-based generator](https://github.com/wangyi-fudan/wyhash) it previously used. Random seeds are needed because if an attacker knows the specific hash function used by a map implementation, they can prepare input that drives the map into quadratic behavior (see Crosby and Wallach’s “[Denial of Service via Algorithmic Complexity Attacks](https://www.usenix.org/conference/12th-usenix-security-symposium/denial-service-algorithmic-complexity-attacks)”). Using a per-map seed, instead of one global seed for all maps, also avoids [other degenerate behaviors](https://accidentallyquadratic.tumblr.com/post/153545455987/rust-hash-iteration-reinsertion). It is not strictly clear that maps need a cryptographically random seed, but it’s also not clear that they don’t. It seemed prudent and was trivial to switch.

   ​	运行时使用 ChaCha8Rand 为每个新映射选择哈希种子，而不是之前使用的安全性较低的基于 wyrand 的生成器。需要随机种子，因为如果攻击者知道映射实现使用的特定哈希函数，他们可以准备将映射变为二次行为的输入（参见 Crosby 和 Wallach 的“通过算法复杂性攻击进行拒绝服务”）。使用每个映射的种子，而不是所有映射的一个全局种子，还可以避免其他退化行为。映射是否需要加密随机种子并不完全明确，但它们不需要也不明确。切换起来似乎很谨慎，而且也很简单。

Code that needs its own ChaCha8Rand instances can create its own [`rand.ChaCha8`](https://go.dev/pkg/math/rand/v2/#ChaCha8) directly.

​	需要其自己的 ChaCha8Rand 实例的代码可以直接创建其自己的 `rand.ChaCha8` 。

## Fixing Security Mistakes 修复安全错误

Go aims to help developers write code that is secure by default. When we observe a common mistake with security consequences, we look for ways to reduce the risk of that mistake or eliminate it entirely. In this case, `math/rand`’s global generator was far too predictable, leading to serious problems in a variety of contexts.

​	Go 的目标是帮助开发者编写默认情况下安全的代码。当我们观察到一个具有安全后果的常见错误时，我们会寻找方法来降低该错误的风险或彻底消除它。在这种情况下， `math/rand` 的全局生成器过于可预测，导致在各种上下文中出现严重问题。

For example, when Go 1.20 deprecated [`math/rand`’s `Read`](https://go.dev/pkg/math/rand/#Read), we heard from developers who discovered (thanks to tooling pointing out use of deprecated functionality) they had been using it in places where [`crypto/rand`’s `Read`](https://go.dev/pkg/crypto/rand/#Read) was definitely needed, like generating key material. Using Go 1.20, that mistake is a serious security problem that merits a detailed investigation to understand the damage. Where were the keys used? How were the keys exposed? Were other random outputs exposed that might allow an attacker to derive the keys? And so on. Using Go 1.22, that mistake is just a mistake. It’s still better to use `crypto/rand`, because the operating system kernel can do a better job keeping the random values secret from various kinds of prying eyes, the kernel is continually adding new entropy to its generator, and the kernel has had more scrutiny. But accidentally using `math/rand` is no longer a security catastrophe.

​	例如，当 Go 1.20 弃用 `math/rand` 的 `Read` 时，我们从开发人员那里得知（感谢工具指出了弃用功能的使用），他们一直在使用它，在 `crypto/rand` 的 `Read` 绝对需要的地方，例如生成密钥材料。使用 Go 1.20，该错误是一个严重的安全问题，需要进行详细调查以了解损害。密钥在哪里使用？密钥是如何暴露的？是否暴露了其他随机输出，可能允许攻击者推导出密钥？等等。使用 Go 1.22，该错误只是一个错误。最好还是使用 `crypto/rand` ，因为操作系统内核可以更好地保护随机值不被各种窥探者窃取，内核会不断为其生成器添加新的熵，并且内核已经接受了更多的审查。但意外使用 `math/rand` 不再是安全灾难。

There are also a variety of use cases that don’t seem like “crypto” but nonetheless need unpredictable randomness. These cases are made more robust by using ChaCha8Rand instead of the Go 1 generator.

​	还有一些看起来不像“加密”但仍然需要不可预测随机性的各种用例。使用 ChaCha8Rand 而不是 Go 1 生成器可以使这些用例更加健壮。

For example, consider generating a [random UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random)). Since UUIDs are not secret, using `math/rand` might seem fine. But if `math/rand` has been seeded with the current time, then running it at the same instant on different computers will produce the same value, making them not “universally unique”. This is especially likely on systems where the current time is only available with millisecond precision. Even with auto-seeding using OS-provided entropy, as introduced in Go 1.20, the Go 1 generator’s seed is only a 63-bit integer, so a program that generates a UUID at startup can only generate 2⁶³ possible UUIDs and is likely to see collisions after 2³¹ or so UUIDs. Using Go 1.22, the new ChaCha8Rand generator is seeded from 256 bits of entropy and can generate 2²⁵⁶ possible first UUIDs. It does not need to worry about collisions.

​	例如，考虑生成一个随机 UUID。由于 UUID 不是秘密，因此使用 `math/rand` 似乎很好。但是，如果 `math/rand` 已使用当前时间播种，那么在不同的计算机上同时运行它将产生相同的值，从而使其不再“通用唯一”。在当前时间只能以毫秒精度获得的系统上尤其如此。即使使用 Go 1.20 中引入的 OS 提供的熵进行自动播种，Go 1 生成器的种子也只是一个 63 位整数，因此在启动时生成 UUID 的程序只能生成 2⁶³ 个可能的 UUID，并且在 2³¹ 个左右的 UUID 之后可能会看到冲突。使用 Go 1.22，新的 ChaCha8Rand 生成器从 256 位熵中播种，并且可以生成 2²⁵⁶ 个可能的第一个 UUID。它不必担心冲突。

As another example, consider load balancing in a front-end server that randomly assigns incoming requests to back-end servers. If an attacker can observe the assignments and knows the predictable algorithm generating them, then the attacker could send a stream of mostly cheap requests but arrange for all the expensive requests to land on a single back-end server. This is an unlikely but plausible problem using the Go 1 generator. Using Go 1.22, it’s not a problem at all.

​	作为另一个示例，考虑前端服务器中的负载平衡，它将传入请求随机分配给后端服务器。如果攻击者可以观察分配并知道生成它们的算法是可预测的，那么攻击者可以发送一个流，其中大部分是廉价请求，但安排所有昂贵的请求都落在一个后端服务器上。这是一个不太可能但合理的问题，使用 Go 1 生成器。使用 Go 1.22，这根本不是问题。

In all these examples, Go 1.22 has eliminated or greatly reduced security problems.

​	在所有这些示例中，Go 1.22 已消除或极大地减少了安全问题。

## Performance 性能

The security benefits of ChaCha8Rand do have a small cost, but ChaCha8Rand is still in the same ballpark as both the Go 1 generator and PCG. The following graphs compare the performance of the three generators, across a variety of hardware, running two operations: the primitive operation “Uint64,” which returns the next `uint64` in the random stream, and the higher-level operation “N(1000),” which returns a random value in the range [0, 1000).

​	ChaCha8Rand 的安全优势确实有小代价，但 ChaCha8Rand 仍然与 Go 1 生成器和 PCG 处于同一水平。以下图表比较了三个生成器的性能，它们在各种硬件上运行两个操作：基本操作“Uint64”，它返回随机流中的下一个 `uint64` ，以及更高级别的操作“N(1000)”，它返回 [0, 1000) 范围内的随机值。

![img](./SecureRandomnessinGo1_22_img/amd.svg+xml) ![img](./SecureRandomnessinGo1_22_img/intel.svg+xml) ![img](./SecureRandomnessinGo1_22_img/amd32.svg+xml) ![img](./SecureRandomnessinGo1_22_img/intel32.svg+xml) ![img](./SecureRandomnessinGo1_22_img/m1.svg+xml) ![img](./SecureRandomnessinGo1_22_img/m3.svg+xml) ![img](./SecureRandomnessinGo1_22_img/taut2a.svg+xml)

The “running 32-bit code” graphs show modern 64-bit x86 chips executing code built with `GOARCH=386`, meaning they are running in 32-bit mode. In that case, the fact that PCG requires 128-bit multiplications makes it slower than ChaCha8Rand, which only uses 32-bit SIMD arithmetic. Actual 32-bit systems matter less every year, but it is still interesting that ChaCha8Rand is faster than PCG on those systems.

​	“运行 32 位代码”图表显示，现代 64 位 x86 芯片执行使用 `GOARCH=386` 构建的代码，这意味着它们在 32 位模式下运行。在这种情况下，PCG 需要 128 位乘法这一事实使其比 ChaCha8Rand 慢，后者仅使用 32 位 SIMD 算术。真正的 32 位系统每年都变得不那么重要，但 ChaCha8Rand 在这些系统上比 PCG 快仍然很有趣。

On some systems, “Go 1: Uint64” is faster than “PCG: Uint64”, but “Go 1: N(1000)” is slower than “PCG: N(1000)”. This happens because “Go 1: N(1000)” is using `math/rand`’s algorithm for reducing a random `int64` down to a value in the range [0, 1000), and that algorithm does two 64-bit integer divide operations. In contrast, “PCG: N(1000)” and “ChaCha8: N(1000)” use the [faster `math/rand/v2` algorithm](https://go.dev/blog/randv2#problem.rand), which almost always avoids the divisions. Removing the 64-bit divisions dominates the algorithm change for 32-bit execution and on the Ampere.

​	在某些系统上，“Go 1: Uint64”比“PCG: Uint64”快，但“Go 1: N(1000)”比“PCG: N(1000)”慢。这是因为“Go 1: N(1000)”使用 `math/rand` 的算法将随机 `int64` 缩小到[0, 1000)范围内的值，并且该算法执行两个 64 位整数除法运算。相比之下，“PCG: N(1000)”和“ChaCha8: N(1000)”使用更快的 `math/rand/v2` 算法，该算法几乎总是避免除法。对于 32 位执行和 Ampere，删除 64 位除法主导了算法更改。

Overall, ChaCha8Rand is slower than the Go 1 generator, but it is never more than twice as slow, and on typical servers the difference is never more than 3ns. Very few programs will be bottlenecked by this difference, and many programs will enjoy the improved security.

​	总体而言，ChaCha8Rand 比 Go 1 生成器慢，但它从不慢两倍以上，在典型服务器上，差异从不超过 3ns。很少有程序会因这种差异而成为瓶颈，而许多程序会享受改进后的安全性。

## Conclusion 结论

Go 1.22 makes your programs more secure without any code changes. We did this by identifying the common mistake of accidentally using `math/rand` instead of `crypto/rand` and then strengthening `math/rand`. This is one small step in Go’s ongoing journey to keep programs safe by default.

​	Go 1.22 让你的程序更安全，无需任何代码更改。我们通过识别意外使用 `math/rand` 而不是 `crypto/rand` 的常见错误，然后加强 `math/rand` 来做到这一点。这是 Go 持续进行的让程序默认保持安全的旅程中的一小步。

These kinds of mistakes are not unique to Go. For example, the npm `keypair` package tries to generate an RSA key pair using Web Crypto APIs, but if they’re not available, it falls back to JavaScript’s `Math.random`. This is hardly an isolated case, and the security of our systems cannot depend on developers not making mistakes. Instead, we hope that eventually all programming languages will move to cryptographically strong pseudorandom generators even for “mathematical” randomness, eliminating this kind of mistake, or at least greatly reducing its blast radius. Go 1.22’s [ChaCha8Rand](https://c2sp.org/chacha8rand) implementation proves that this approach is competitive with other generators.

​	此类错误并非 Go 独有。例如，npm `keypair` 包尝试使用 Web Crypto API 生成 RSA 密钥对，但如果这些 API 不可用的情况下，它会回退到 JavaScript 的 `Math.random` 。这绝非孤立案例，我们系统的安全性不能依赖于开发人员不犯错。相反，我们希望最终所有编程语言都转向密码学强伪随机生成器，即使是“数学”随机性，从而消除此类错误，或至少大幅减小其影响范围。Go 1.22 的 ChaCha8Rand 实现证明了此方法与其他生成器具有竞争力。