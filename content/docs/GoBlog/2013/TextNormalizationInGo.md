+++
title = "go 的文本规范化"
weight = 7
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Text normalization in Go - go 的文本规范化

https://go.dev/blog/normalization

Marcel van Lohuizen
26 November 2013

## Introduction 简介

An earlier [post](https://blog.golang.org/strings) talked about strings, bytes and characters in Go. I’ve been working on various packages for multilingual text processing for the go.text repository. Several of these packages deserve a separate blog post, but today I want to focus on [go.text/unicode/norm](https://pkg.go.dev/golang.org/x/text/unicode/norm), which handles normalization, a topic touched in the [strings article](https://blog.golang.org/strings) and the subject of this post. Normalization works at a higher level of abstraction than raw bytes.

之前的一篇文章谈到了Go中的字符串、字节和字符。我一直在为go.text资源库开发各种多语言文本处理包。其中有几个包值得单独写一篇博文，但今天我想重点介绍go.text/unicode/norm，它可以处理规范化，这是字符串文章中提到的一个话题，也是这篇文章的主题。归一化是在比原始字节更高的抽象层次上工作的。

To learn pretty much everything you ever wanted to know about normalization (and then some), [Annex 15 of the Unicode Standard](http://unicode.org/reports/tr15/) is a good read. A more approachable article is the corresponding [Wikipedia page](http://en.wikipedia.org/wiki/Unicode_equivalence). Here we focus on how normalization relates to Go.

要了解你想知道的关于规范化的几乎所有内容（还有一些），Unicode标准的附件15是一个很好的读物。相应的维基百科页面是一篇更容易理解的文章。在这里，我们主要讨论规范化与Go的关系。

## What is normalization? 什么是规范化？

There are often several ways to represent the same string. For example, an é (e-acute) can be represented in a string as a single rune ("\u00e9") or an ‘e’ followed by an acute accent (“e\u0301”). According to the Unicode standard, these two are “canonically equivalent” and should be treated as equal.

同一个字符串通常有几种表示方法。例如，一个é（e-acute）在字符串中可以表示为一个单独的符文（""u00e9"）或一个'e'后面跟着一个锐角重音（"e/u0301"）。根据Unicode标准，这两个是 "典型的等价物"，应该被视为相等。

Using a byte-to-byte comparison to determine equality would clearly not give the right result for these two strings. Unicode defines a set of normal forms such that if two strings are canonically equivalent and are normalized to the same normal form, their byte representations are the same.

对于这两个字符串，使用逐个字节的比较来确定相等，显然不会得到正确的结果。Unicode定义了一组正常形式，如果两个字符串在规范上是等价的，并且被规范化为相同的正常形式，那么它们的字节表示就是相同的。

Unicode also defines a “compatibility equivalence” to equate characters that represent the same characters, but may have a different visual appearance. For example, the superscript digit ‘⁹’ and the regular digit ‘9’ are equivalent in this form.

Unicode还定义了一个 "兼容性等价"，以等价表示相同的字符，但可能有不同的视觉外观。例如，上标数字"⁹"和普通数字 "9 "在这种形式下是等同的。

For each of these two equivalence forms, Unicode defines a composing and decomposing form. The former replaces runes that can combine into a single rune with this single rune. The latter breaks runes apart into their components. This table shows the names, all starting with NF, by which the Unicode Consortium identifies these forms:

对于这两种等价形式中的每一种，Unicode都定义了一种合成和分解形式。前者用这个单一的符文取代了可以结合成一个符文的符文。后者则将符文分解为其组成部分。该表显示了Unicode联盟用来识别这些形式的名称，全部以NF开头。

|                                      | Composing 组成 | Decomposing 分解 |
| ------------------------------------ | -------------- | ---------------- |
| Canonical equivalence 正则等价       | NFC            | NFD              |
| Compatibility equivalence 兼容性等价 | NFKC           | NFKD             |

## Go’s approach to normalization - Go的规范化方法

As mentioned in the strings blog post, Go does not guarantee that characters in a string are normalized. However, the go.text packages can compensate. For example, the [collate](https://pkg.go.dev/golang.org/x/text/collate) package, which can sort strings in a language-specific way, works correctly even with unnormalized strings. The packages in go.text do not always require normalized input, but in general normalization may be necessary for consistent results.

正如在字符串博文中提到的，Go 并不保证字符串中的字符是规范化的。然而，go.text 包可以进行补偿。例如，collate包可以以特定语言的方式对字符串进行排序，即使在未规范化的字符串中也能正确工作。go.text中的包并不总是需要规范化的输入，但一般来说，为了获得一致的结果，规范化可能是必要的。

Normalization isn’t free but it is fast, particularly for collation and searching or if a string is either in NFD or in NFC and can be converted to NFD by decomposing without reordering its bytes. In practice, [99.98%](http://www.macchiato.com/unicode/nfc-faq#TOC-How-much-text-is-already-NFC-) of the web’s HTML page content is in NFC form (not counting markup, in which case it would be more). By far most NFC can be decomposed to NFD without the need for reordering (which requires allocation). Also, it is efficient to detect when reordering is necessary, so we can save time by doing it only for the rare segments that need it.

归一化不是免费的，但它是快速的，特别是对于整理和搜索，或者如果一个字符串是NFD或NFC，并且可以通过分解转换为NFD而不重新排序其字节。在实践中，99.98%的网络HTML页面内容是以NFC形式存在的（不算标记，在这种情况下会更多）。到目前为止，大多数NFC可以被分解为NFD，而不需要重新排序（这需要分配）。另外，检测何时需要重新排序是很有效的，所以我们可以通过只对极少数需要重新排序的片段进行重新排序来节省时间。

To make things even better, the collation package typically does not use the norm package directly, but instead uses the norm package to interleave normalization information with its own tables. Interleaving the two problems allows for reordering and normalization on the fly with almost no impact on performance. The cost of on-the-fly normalization is compensated by not having to normalize text beforehand and ensuring that the normal form is maintained upon edits. The latter can be tricky. For instance, the result of concatenating two NFC-normalized strings is not guaranteed to be in NFC.

为了使事情变得更好，整理包通常不直接使用规范包，而是使用规范包将规范化信息与它自己的表进行交错。将这两个问题交织在一起，可以在飞行中进行重新排序和规范化，对性能几乎没有影响。即时规范化的成本被补偿了，因为不必事先对文本进行规范化，并确保在编辑时保持正常的形式。后者可能很棘手。例如，将两个NFC规范化的字符串连接起来的结果不能保证是NFC。

Of course, we can also avoid the overhead outright if we know in advance that a string is already normalized, which is often the case.

当然，如果我们事先知道一个字符串已经被规范化了，我们也可以直接避免开销，这种情况经常发生。

## Why bother? 为什么要这样做呢？

After all this discussion about avoiding normalization, you might ask why it’s worth worrying about at all. The reason is that there are cases where normalization is required and it is important to understand what those are, and in turn how to do it correctly.

在讨论了这么多关于避免规范化的问题之后，你可能会问为什么值得担心这个问题。原因是有一些情况是需要规范化的，重要的是要了解这些情况是什么，以及如何正确地做到这一点。

Before discussing those, we must first clarify the concept of ‘character’.

在讨论这些之前，我们必须首先澄清 "字符 "的概念。

## What is a character? 什么是字符？

As was mentioned in the strings blog post, characters can span multiple runes. For example, an ‘e’ and ‘◌́’ (acute “\u0301”) can combine to form ‘é’ (“e\u0301” in NFD).  Together these two runes are one character. The definition of a character may vary depending on the application. For normalization we will define it as a sequence of runes that starts with a starter, a rune that does not modify or combine backwards with any other rune, followed by possibly empty sequence of non-starters, that is, runes that do (typically accents). The normalization algorithm processes one character at a time.

正如串联博文中提到的，字符可以跨越多个符文。例如，"e "和"◌́"（锐利的"\u0301"）可以组合成 "é"（NFD中的 "e\u0301"）。 这两个符文合在一起就是一个字符。一个字符的定义可能因应用而异。对于规范化来说，我们将其定义为一个符文序列，以起始符文开始，即一个不修改或不与任何其他符文反向组合的符文，后面可能是空的非起始符文序列，即有的符文（通常是重音）。归一化算法一次处理一个字符。

Theoretically, there is no bound to the number of runes that can make up a Unicode character. In fact, there are no restrictions on the number of modifiers that can follow a character and a modifier may be repeated, or stacked. Ever seen an ‘e’ with three acutes? Here you go: ‘é́́’. That is a perfectly valid 4-rune character according to the standard.

理论上，组成Unicode字符的符文数量没有限制。事实上，对一个字符后面可以有多少个修饰符没有限制，一个修饰符可以重复，也可以叠加。见过一个有三个尖锐的'e'吗？给你。'é́́'. 根据该标准，这是一个完全有效的4韵母字符。

As a consequence, even at the lowest level, text needs to be processed in increments of unbounded chunk sizes. This is especially awkward with a streaming approach to text processing, as used by Go’s standard Reader and Writer interfaces, as that model potentially requires any intermediate buffers to have unbounded size as well. Also, a straightforward implementation of normalization will have a O(n²) running time.

因此，即使在最底层，文本也需要以无界块大小的增量进行处理。这在文本处理的流式方法中尤其尴尬，就像Go的标准Reader和Writer接口所使用的那样，因为该模型可能要求任何中间缓冲区也有无界的大小。另外，一个直接的规范化实现会有一个O(n²)的运行时间。

There are really no meaningful interpretations for such large sequences of modifiers for practical applications. Unicode defines a Stream-Safe Text format, which allows capping the number of modifiers (non-starters) to at most 30, more than enough for any practical purpose. Subsequent modifiers will be placed after a freshly inserted Combining Grapheme Joiner (CGJ or U+034F). Go adopts this approach for all normalization algorithms. This decision gives up a little conformance but gains a little safety.

在实际应用中，对这样大的修饰语序列确实没有有意义的解释。Unicode定义了一种流安全文本格式，它允许将修饰符（非起始符）的数量限制在最多30个，对于任何实际用途来说都绰绰有余。后续的修饰符将被放置在新插入的组合字素连接符（CGJ或U+034F）之后。Go对所有规范化算法都采用这种方法。这个决定放弃了一点一致性，但获得了一点安全性。

## Writing in normal form 以正常形式书写

Even if you don’t need to normalize text within your Go code, you might still want to do so when communicating to the outside world. For example, normalizing to NFC might compact your text, making it cheaper to send down a wire. For some languages, like Korean, the savings can be substantial. Also, some external APIs might expect text in a certain normal form. Or you might just want to fit in and output your text as NFC like the rest of the world.

即使您不需要在 Go 代码中对文本进行规范化处理，但在与外部世界交流时，您可能仍然希望这样做。例如，对NFC进行规范化处理可能会压缩你的文本，使其在发送时更便宜。对于某些语言，如韩语，可以节省大量的费用。此外，一些外部API可能希望文本采用某种正常形式。或者你可能只是想适应并像世界上其他地方一样将你的文本输出为NFC。

To write your text as NFC, use the [unicode/norm](https://pkg.go.dev/golang.org/x/text/unicode/norm) package to wrap your `io.Writer` of choice:

要把你的文本写成NFC，使用unicode/norm包来包装你选择的io.Writer：

```go linenums="1"
wc := norm.NFC.Writer(w)
defer wc.Close()
// write as before...
```

If you have a small string and want to do a quick conversion, you can use this simpler form:

如果你有一个小的字符串，想做快速转换，你可以使用这种更简单的形式：

```go linenums="1"
norm.NFC.Bytes(b)
```

Package norm provides various other methods for normalizing text. Pick the one that suits your needs best.

包norm提供了其他各种规范化文本的方法。挑选一个最适合你的需求。

## Catching look-alikes 捕捉外观相似的东西

Can you tell the difference between ‘K’ ("\u004B") and ‘K’ (Kelvin sign “\u212A”) or ‘Ω’ ("\u03a9") and ‘Ω’ (Ohm sign “\u2126”)? It is easy to overlook the sometimes minute differences between variants of the same underlying character. It is generally a good idea to disallow such variants in identifiers or anything where deceiving users with such look-alikes can pose a security hazard.

你能分辨出'K'（"\u004B"）和'K'（开尔文符号"\u212A"）或'Ω'（"\u03a9"）和'Ω'（欧姆符号"\u2126"）之间的区别吗？我们很容易忽视同一个基本字符的变体之间有时存在的微小差异。一般来说，不允许在标识符或任何用这种外观相似的东西欺骗用户的地方出现这种变体是一个好主意。

The compatibility normal forms, NFKC and NFKD, will map many visually nearly identical forms to a single value. Note that it will not do so when two symbols look alike, but are really from two different alphabets. For example the Latin ‘o’, Greek ‘ο’, and Cyrillic ‘о’ are still different characters as defined by these forms.

兼容性正常形式，NFKC和NFKD，将把许多视觉上几乎相同的形式映射到一个单一的值。请注意，当两个符号看起来很像，但实际上来自两个不同的字母时，它不会这样做。例如，拉丁文的 "o"、希腊文的 "ο "和西里尔文的 "о "仍然是这些形式所定义的不同字符。

## Correct text modifications 正确的文本修改

The norm package might also come to the rescue when one needs to modify text. Consider a case where you want to search and replace the word “cafe” with its plural form “cafes”.  A code snippet could look like this.

当人们需要修改文本时，规范包也可能会出手相助。考虑一下这样的情况：你想用复数形式 "cafes "来搜索和替换 "cafe "这个词。 一个代码片断可以是这样的。

```go linenums="1"
s := "We went to eat at multiple cafe"
cafe := "cafe"
if p := strings.Index(s, cafe); p != -1 {
    p += len(cafe)
    s = s[:p] + "s" + s[p:]
}
fmt.Println(s)
```

This prints “We went to eat at multiple cafes” as desired and expected. Now consider our text contains the French spelling “café” in NFD form:

这就打印出了 "我们去了多家咖啡馆吃饭"，这是所希望的，也是预期的。现在考虑我们的文本包含NFD形式的法语拼写 "café"：

```go linenums="1"
s := "We went to eat at multiple cafe\u0301"
```

Using the same code from above, the plural “s” would still be inserted after the ‘e’, but before the acute, resulting in  “We went to eat at multiple cafeś”.  This behavior is undesirable.

使用上面的相同代码，复数 "s "仍将插入'e'之后，但在急性之前，结果是 "We went to eat at multiple cafeś"。 这种行为是不可取的。

The problem is that the code does not respect the boundaries between multi-rune characters and inserts a rune in the middle of a character.  Using the norm package, we can rewrite this piece of code as follows:

问题在于，代码没有尊重多符文字符之间的界限，在字符中间插入了一个符文。 使用规范包，我们可以把这段代码改写成如下：

```go linenums="1"
s := "We went to eat at multiple cafe\u0301"
cafe := "cafe"
if p := strings.Index(s, cafe); p != -1 {
    p += len(cafe)
    if bp := norm.FirstBoundary(s[p:]); bp > 0 {
        p += bp
    }
    s = s[:p] + "s" + s[p:]
}
fmt.Println(s)
```

This may be a contrived example, but the gist should be clear. Be mindful of the fact that characters can span multiple runes. Generally these kinds of problems can be avoided by using search functionality that respects character boundaries (such as the planned go.text/search package.)

这可能是一个矫揉造作的例子，但要点应该很清楚。请注意，字符可以跨越多个符文的事实。一般来说，这些问题可以通过使用尊重字符边界的搜索功能来避免（比如计划中的go.text/search包。）

## Iteration 迭代

Another tool provided by the norm package that may help dealing with character boundaries is its iterator, [`norm.Iter`](https://pkg.go.dev/golang.org/x/text/unicode/norm#Iter). It iterates over characters one at a time in the normal form of choice.

norm包提供的另一个有助于处理字符边界的工具是它的迭代器，norm.Iter。它以选择的法线形式一个一个地迭代字符。

## Performing magic 施展魔法

As mentioned earlier, most text is in NFC form, where base characters and modifiers are combined into a single rune whenever possible.  For the purpose of analyzing characters, it is often easier to handle runes after decomposition into their smallest components. This is where the NFD form comes in handy. For example, the following piece of code creates a `transform.Transformer` that decomposes text into its smallest parts, removes all accents, and then recomposes the text into NFC:

如前所述，大多数文本都是NFC形式的，其中基本字符和修饰符尽可能合并成一个符文。 为了分析字符，将符文分解成最小的组成部分后，往往更容易处理。这就是NFD形式的用武之地。例如，下面这段代码创建了一个Transform.Transformer，将文本分解成最小的部分，删除所有重音，然后将文本重新组成NFC：

```go linenums="1"
import (
    "unicode"

    "golang.org/x/text/transform"
    "golang.org/x/text/unicode/norm"
)

isMn := func(r rune) bool {
    return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
```

The resulting `Transformer` can be used to remove accents from an `io.Reader` of choice as follows:

产生的Transformer可以用来从选择的io.Reader中移除重音，如下所示：

```go linenums="1"
r = transform.NewReader(r, t)
// read as before ...
```

This will, for example, convert any mention of “cafés” in the text to “cafes”, regardless of the normal form in which the original text was encoded.

例如，这将把文本中提到的任何 "cafés "转换为 "cafes"，而不管原始文本是以何种规范形式编码的。

## Normalization info 正常化信息

As mentioned earlier, some packages precompute normalizations into their tables to minimize the need for normalization at run time. The type `norm.Properties` provides access to the per-rune information needed by these packages, most notably the Canonical Combining Class and decomposition information. Read the [documentation](https://pkg.go.dev/golang.org/x/text/unicode/norm#Properties) for this type if you want to dig deeper.

如前所述，一些包在它们的表中预先计算了规范化，以尽量减少在运行时对规范化的需求。norm.Properties类型提供了对这些包所需要的每个规范化信息的访问，最明显的是Canonical Combining Class和分解信息。如果你想深入研究，请阅读这个类型的文档。

## Performance 性能

To give an idea of the performance of normalization, we compare it against the performance of strings.ToLower. The sample in the first row is both lowercase and NFC and can in every case be returned as is. The second sample is neither and requires writing a new version.

为了了解规范化的性能，我们将其与strings.ToLower的性能进行比较。第一行的样本既是小写字母又是NFC，在任何情况下都可以原样返回。第二个例子既不是，又需要写一个新的版本。

| Input               | ToLower | NFC Append | NFC Transform | NFC Iter        |
| ------------------- | ------- | ---------- | ------------- | --------------- |
| nörmalization       | 199 ns  | 137 ns     | 133 ns        | 251 ns (621 ns) |
| No\u0308rmalization | 427 ns  | 836 ns     | 845 ns        | 573 ns (948 ns) |

The column with the results for the iterator shows both the measurement with and without initialization of the iterator, which contain buffers that don’t need to be reinitialized upon reuse.

有迭代器结果的那一列显示了有迭代器初始化和无迭代器初始化的测量结果，其中包含的缓冲区在重复使用时不需要被重新初始化。

As you can see, detecting whether a string is normalized can be quite efficient. A lot of the cost of normalizing in the second row is for the initialization of buffers, the cost of which is amortized when one is processing larger strings. As it turns out, these buffers are rarely needed, so we may change the implementation at some point to speed up the common case for small strings even further.

正如你所看到的，检测一个字符串是否被规范化可以是相当有效的。第二行中规范化的很多成本是用于缓冲区的初始化，当人们处理较大的字符串时，其成本会被摊薄。事实证明，这些缓冲区很少需要，所以我们可能会在某个时候改变实现，以进一步加快小字符串的常见情况。

## Conclusion 结论

If you’re dealing with text inside Go, you generally do not have to use the unicode/norm package to normalize your text. The package may still be useful for things like ensuring that strings are normalized before sending them out or to do advanced text manipulation.

如果您在Go中处理文本，您一般不需要使用unicode/norm包来规范您的文本。该包对于确保字符串在发送前被规范化或进行高级文本操作等方面仍有帮助。

This article briefly mentioned the existence of other go.text packages as well as multilingual text processing and it may have raised more questions than it has given answers. The discussion of these topics, however, will have to wait until another day.

这篇文章简要地提到了其他go.text包的存在以及多语言文本处理，它所提出的问题可能比它所给出的答案要多。然而，对这些话题的讨论，将不得不等到下一天。
