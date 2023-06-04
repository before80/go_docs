+++
title = "go 中的字符串、字节、符文和字符"
weight = 6
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Strings, bytes, runes and characters in Go - go 中的字符串、字节、符文和字符

https://go.dev/blog/strings

Rob Pike
23 October 2013

## Introduction 简介

The [previous blog post](https://blog.golang.org/slices) explained how slices work in Go, using a number of examples to illustrate the mechanism behind their implementation. Building on that background, this post discusses strings in Go. At first, strings might seem too simple a topic for a blog post, but to use them well requires understanding not only how they work, but also the difference between a byte, a character, and a rune, the difference between Unicode and UTF-8, the difference between a string and a string literal, and other even more subtle distinctions.

上一篇博文解释了Go中的分片是如何工作的，用一些例子来说明其实现背后的机制。在此背景下，本篇文章将讨论Go中的字符串。起初，对于一篇博文来说，字符串似乎过于简单，但要很好地使用它们，不仅需要了解它们的工作原理，还需要了解字节、字符和符文之间的区别，Unicode和UTF-8之间的区别，字符串和字符串字面的区别，以及其他更微妙的区别。

One way to approach this topic is to think of it as an answer to the frequently asked question, "When I index a Go string at position *n*, why don’t I get the *nth* character?" As you’ll see, this question leads us to many details about how text works in the modern world.

处理这个问题的一种方法是把它看作是对一个经常被问到的问题的回答："当我在位置n索引一个Go字符串时，为什么我没有得到第n个字符？" 正如您将看到的，这个问题将我们引向许多关于现代世界中文本如何工作的细节。

An excellent introduction to some of these issues, independent of Go, is Joel Spolsky’s famous blog post, [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](http://www.joelonsoftware.com/articles/Unicode.html). Many of the points he raises will be echoed here.

乔尔-斯波尔斯基（Joel Spolsky）的著名博文《每个软件开发者绝对必须了解的Unicode和字符集（没有借口！）》是对这些问题的一个很好的介绍，它与Go无关。他提出的许多观点将在这里得到回应。

## What is a string? 什么是字符串？

Let’s start with some basics.

让我们从一些基础知识开始。

In Go, a string is in effect a read-only slice of bytes. If you’re at all uncertain about what a slice of bytes is or how it works, please read the [previous blog post](https://blog.golang.org/slices); we’ll assume here that you have.

在Go中，字符串实际上是一个只读的字节片。如果您不知道什么是字节片，或者它是如何工作的，请阅读之前的博文；我们在此假定您已经阅读了。

It’s important to state right up front that a string holds *arbitrary* bytes. It is not required to hold Unicode text, UTF-8 text, or any other predefined format. As far as the content of a string is concerned, it is exactly equivalent to a slice of bytes.

重要的是要在前面说明，一个字符串可以容纳任意的字节。它不需要持有Unicode文本、UTF-8文本或任何其他预定义的格式。就字符串的内容而言，它完全等同于一个字节的片断。

Here is a string literal (more about those soon) that uses the `\xNN` notation to define a string constant holding some peculiar byte values. (Of course, bytes range from hexadecimal values 00 through FF, inclusive.)

下面是一个字符串字头（很快会有更多的介绍），它使用了\xNN符号来定义一个字符串常量，持有一些奇特的字节值。(当然，字节的范围是从十六进制的00到FF，包括在内）。

```go linenums="1"
    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
```

## Printing strings 打印字符串

Because some of the bytes in our sample string are not valid ASCII, not even valid UTF-8, printing the string directly will produce ugly output. The simple print statement

由于我们的样本字符串中的一些字节不是有效的ASCII码，甚至不是有效的UTF-8码，直接打印字符串将产生丑陋的输出。简单的打印语句

```go linenums="1"
    fmt.Println(sample)
```

produces this mess (whose exact appearance varies with the environment):

就会产生这种混乱的结果（其具体外观随环境变化而变化）:

```go linenums="1"
��=� ⌘
```

To find out what that string really holds, we need to take it apart and examine the pieces. There are several ways to do this. The most obvious is to loop over its contents and pull out the bytes individually, as in this `for` loop:

为了弄清这个字符串的真正含义，我们需要把它拆开来检查一下各个部分。有几种方法可以做到这一点。最明显的是在其内容上进行循环，并逐个拉出字节，如这个for循环：

```go linenums="1"
    for i := 0; i < len(sample); i++ {
        fmt.Printf("%x ", sample[i])
    }
```

As implied up front, indexing a string accesses individual bytes, not characters. We’ll return to that topic in detail below. For now, let’s stick with just the bytes. This is the output from the byte-by-byte loop:

正如前面所暗示的，索引字符串是访问单个字节，而不是字符。我们将在下面详细讨论这个问题。现在，让我们只关注字节。这就是逐个字节循环的输出：

```
bd b2 3d bc 20 e2 8c 98
```

Notice how the individual bytes match the hexadecimal escapes that defined the string.

注意各个字节是如何与定义字符串的十六进制转义相匹配的。

A shorter way to generate presentable output for a messy string is to use the `%x` (hexadecimal) format verb of `fmt.Printf`. It just dumps out the sequential bytes of the string as hexadecimal digits, two per byte.

为一个混乱的字符串生成可展示的输出的一个更短的方法是使用fmt.Printf的%x（十六进制）格式动词。它只是将字符串的连续字节以十六进制数字的形式跳出，每个字节两个。

```go linenums="1"
    fmt.Printf("%x\n", sample)
```

Compare its output to that above:

将其输出与上面的输出进行比较：

```
bdb23dbc20e28c98
```

A nice trick is to use the "space" flag in that format, putting a space between the `%` and the `x`. Compare the format string used here to the one above,

一个很好的技巧是在该格式中使用 "空格 "标志，在%和x之间放一个空格，将这里使用的格式字符串与上面的比较，

```go linenums="1"
    fmt.Printf("% x\n", sample)
```

and notice how the bytes come out with spaces between, making the result a little less imposing:

并注意到字节之间是如何产生空格的，从而使结果不那么令人讨厌：

```
bd b2 3d bc 20 e2 8c 98
```

There’s more. The `%q` (quoted) verb will escape any non-printable byte sequences in a string so the output is unambiguous.

还有更多。%q（引号）动词将转义字符串中任何不可打印的字节序列，因此输出是明确的。

```go linenums="1"
    fmt.Printf("%q\n", sample)
```

This technique is handy when much of the string is intelligible as text but there are peculiarities to root out; it produces:

当大部分字符串可以理解为文本，但有一些特殊性需要根除时，这种技术就很方便；它产生的结果是：

```
"\xbd\xb2=\xbc ⌘"
```

If we squint at that, we can see that buried in the noise is one ASCII equals sign, along with a regular space, and at the end appears the well-known Swedish "Place of Interest" symbol. That symbol has Unicode value U+2318, encoded as UTF-8 by the bytes after the space (hex value `20`): `e2` `8c` `98`.

如果我们眯着眼睛看，我们可以看到埋藏在噪音中的是一个ASCII等号，以及一个普通的空格，最后出现的是著名的瑞典语 "兴趣点 "符号。该符号的Unicode值为U+2318，通过空格后的字节（十六进制值为20）编码为UTF-8：e2 8c 98。

If we are unfamiliar or confused by strange values in the string, we can use the "plus" flag to the `%q` verb. This flag causes the output to escape not only non-printable sequences, but also any non-ASCII bytes, all while interpreting UTF-8. The result is that it exposes the Unicode values of properly formatted UTF-8 that represents non-ASCII data in the string:

如果我们对字符串中的奇怪数值不熟悉或感到困惑，我们可以使用%q动词的 "加 "标志。这个标志使输出不仅转义非打印序列，而且还转义任何非ASCII字节，同时解释UTF-8。其结果是，它暴露了正确格式化的UTF-8的Unicode值，代表字符串中的非ASCII数据：

```go linenums="1"
    fmt.Printf("%+q\n", sample)
```

With that format, the Unicode value of the Swedish symbol shows up as a `\u` escape:

有了这种格式，瑞典语符号的Unicode值就显示为\u转义。

```
"\xbd\xb2=\xbc \u2318"
```

These printing techniques are good to know when debugging the contents of strings, and will be handy in the discussion that follows. It’s worth pointing out as well that all these methods behave exactly the same for byte slices as they do for strings.

在调试字符串的内容时，这些打印技术是很好的知识，在接下来的讨论中也会很方便。值得指出的是，所有这些方法对字节片的行为和对字符串的行为是完全一样的。

Here’s the full set of printing options we’ve listed, presented as a complete program you can run (and edit) right in the browser:

下面是我们列出的全部打印选项，以一个完整的程序形式呈现，您可以在浏览器中直接运行（和编辑）：

```go linenums="1"
package main

import "fmt"

func main() {
    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    fmt.Println("Println:")
    fmt.Println(sample)

    fmt.Println("Byte loop:")
    for i := 0; i < len(sample); i++ {
        fmt.Printf("%x ", sample[i])
    }
    fmt.Printf("\n")

    fmt.Println("Printf with %x:")
    fmt.Printf("%x\n", sample)

    fmt.Println("Printf with % x:")
    fmt.Printf("% x\n", sample)

    fmt.Println("Printf with %q:")
    fmt.Printf("%q\n", sample)

    fmt.Println("Printf with %+q:")
    fmt.Printf("%+q\n", sample)
}
```

Run 运行

[Exercise: Modify the examples above to use a slice of bytes instead of a string. Hint: Use a conversion to create the slice.]

[练习。修改上面的例子，用一个字节的片断代替字符串。提示：使用一个转换来创建片断。］

[Exercise: Loop over the string using the `%q` format on each byte. What does the output tell you?]

[练习。在每个字节上使用%q格式在字符串上循环。输出结果告诉您什么？］

## UTF-8 and string literals - UTF-8和字符串字面意义

As we saw, indexing a string yields its bytes, not its characters: a string is just a bunch of bytes. That means that when we store a character value in a string, we store its byte-at-a-time representation. Let’s look at a more controlled example to see how that happens.

正如我们所看到的，对一个字符串进行索引可以得到它的字节，而不是它的字符：一个字符串只是一堆字节。这意味着，当我们在一个字符串中存储一个字符值时，我们存储的是它的字节表示。让我们看一个更有控制力的例子，看看这是如何发生的。

Here’s a simple program that prints a string constant with a single character three different ways, once as a plain string, once as an ASCII-only quoted string, and once as individual bytes in hexadecimal. To avoid any confusion, we create a "raw string", enclosed by back quotes, so it can contain only literal text. (Regular strings, enclosed by double quotes, can contain escape sequences as we showed above.)

这里有一个简单的程序，它以三种不同的方式打印一个带有单个字符的字符串常数，一次是普通字符串，一次是仅有ASCII引号的字符串，还有一次是十六进制的单个字节。为了避免混淆，我们创建了一个 "原始字符串"，用反引号括起来，所以它只能包含字面文本。(正规的字符串，用双引号括起来，可以包含转义序列，正如我们上面所展示的那样）。

```go linenums="1"
func main() {
    const placeOfInterest = `⌘`

    fmt.Printf("plain string: ")
    fmt.Printf("%s", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("quoted string: ")
    fmt.Printf("%+q", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("hex bytes: ")
    for i := 0; i < len(placeOfInterest); i++ {
        fmt.Printf("%x ", placeOfInterest[i])
    }
    fmt.Printf("\n")
}
```

Run 运行

The output is:

输出结果是：

```
plain string: ⌘
quoted string: "\u2318"
hex bytes: e2 8c 98
```

which reminds us that the Unicode character value U+2318, the "Place of Interest" symbol ⌘, is represented by the bytes `e2` `8c` `98`, and that those bytes are the UTF-8 encoding of the hexadecimal value 2318.

这提醒我们，Unicode字符值U+2318，即 "感兴趣的地方 "符号⌘，由字节e2 8c 98表示，而这些字节是十六进制值2318的UTF-8编码。

It may be obvious or it may be subtle, depending on your familiarity with UTF-8, but it’s worth taking a moment to explain how the UTF-8 representation of the string was created. The simple fact is: it was created when the source code was written.

这可能很明显，也可能很微妙，取决于您对UTF-8的熟悉程度，但值得花点时间解释一下这个字符串的UTF-8表示法是如何创建的。简单的事实是：它是在编写源代码的时候创建的。

Source code in Go is *defined* to be UTF-8 text; no other representation is allowed. That implies that when, in the source code, we write the text

Go中的源代码被定义为UTF-8文本，不允许有其他表示方式。这意味着，当我们在源代码中写下文本时

```
`⌘`
```

the text editor used to create the program places the UTF-8 encoding of the symbol ⌘ into the source text. When we print out the hexadecimal bytes, we’re just dumping the data the editor placed in the file.

用于创建程序的文本编辑器会将符号⌘的UTF-8编码放入源文本中。当我们打印出十六进制的字节时，我们只是在转储编辑器放在文件中的数据。

In short, Go source code is UTF-8, so *the source code for the string literal is UTF-8 text*. If that string literal contains no escape sequences, which a raw string cannot, the constructed string will hold exactly the source text between the quotes. Thus by definition and by construction the raw string will always contain a valid UTF-8 representation of its contents. Similarly, unless it contains UTF-8-breaking escapes like those from the previous section, a regular string literal will also always contain valid UTF-8.

简而言之，Go的源代码是UTF-8的，所以字符串字头的源代码是UTF-8文本。如果该字符串字面不包含转义序列，而原始字符串是不能包含转义序列的，那么构建的字符串将在引号之间准确地包含源文本。因此，根据定义和结构，原始字符串将始终包含其内容的有效UTF-8表示。同样地，除非它包含UTF-8中断转义，如上一节所述，否则正则字符串字面也总是包含有效的UTF-8。

Some people think Go strings are always UTF-8, but they are not: only string literals are UTF-8. As we showed in the previous section, string *values* can contain arbitrary bytes; as we showed in this one, string *literals* always contain UTF-8 text as long as they have no byte-level escapes.

有些人认为Go中的字符串总是UTF-8的，但事实并非如此：只有字符串字头是UTF-8的。正如我们在上一节中所展示的，字符串值可以包含任意的字节；正如我们在这一节中所展示的，只要没有字节级转义，字符串字面总是包含UTF-8文本。

To summarize, strings can contain arbitrary bytes, but when constructed from string literals, those bytes are (almost always) UTF-8.

总而言之，字符串可以包含任意的字节，但是当由字符串字头构建时，这些字节（几乎总是）是UTF-8。

## Code points, characters, and runes 代码点、字符和符文

We’ve been very careful so far in how we use the words "byte" and "character". That’s partly because strings hold bytes, and partly because the idea of "character" is a little hard to define. The Unicode standard uses the term "code point" to refer to the item represented by a single value. The code point U+2318, with hexadecimal value 2318, represents the symbol ⌘. (For lots more information about that code point, see [its Unicode page](http://unicode.org/cldr/utility/character.jsp?a=2318).)

到目前为止，我们在使用 "字节 "和 "字符 "这两个词时一直非常谨慎。这一方面是因为字符串持有字节，另一方面是因为 "字符 "的概念有点难以定义。Unicode标准使用术语 "码位 "来指代由单一数值代表的项目。代码点U+2318，十六进制值2318，代表符号⌘。(关于该码位的更多信息，请参见Unicode页面）。

To pick a more prosaic example, the Unicode code point U+0061 is the lower case Latin letter ‘A’: a.

举个更普通的例子，Unicode码位U+0061是小写拉丁字母 "A"：a。

But what about the lower case grave-accented letter ‘A’, à? That’s a character, and it’s also a code point (U+00E0), but it has other representations. For example we can use the "combining" grave accent code point, U+0300, and attach it to the lower case letter a, U+0061, to create the same character à. In general, a character may be represented by a number of different sequences of code points, and therefore different sequences of UTF-8 bytes.

但是，小写的重音字母 "A"，à，又是怎么回事呢？这是一个字符，它也是一个代码点（U+00E0），但它有其他的表示方法。例如，我们可以使用 "组合式 "重音码位，U+0300，并将其附加到小写字母a，U+0061，以创建相同的字符à。

The concept of character in computing is therefore ambiguous, or at least confusing, so we use it with care. To make things dependable, there are *normalization* techniques that guarantee that a given character is always represented by the same code points, but that subject takes us too far off the topic for now. A later blog post will explain how the Go libraries address normalization.

因此，计算中的字符概念是模糊的，或者至少是混乱的，所以我们使用它时要小心。为了使事情变得可靠，有一些规范化的技术可以保证一个给定的字符总是由相同的码位来表示，但是这个主题让我们现在离这个话题太远。稍后的博文将解释Go库如何解决规范化问题。

"Code point" is a bit of a mouthful, so Go introduces a shorter term for the concept: *rune*. The term appears in the libraries and source code, and means exactly the same as "code point", with one interesting addition.

"代码点 "有点拗口，所以Go为这个概念引入了一个更简短的术语：符文。该术语出现在库和源代码中，其含义与 "代码点 "完全相同，但有一个有趣的补充。

The Go language defines the word `rune` as an alias for the type `int32`, so programs can be clear when an integer value represents a code point. Moreover, what you might think of as a character constant is called a *rune constant* in Go. The type and value of the expression

Go语言将rune这个词定义为int32类型的别名，因此程序可以清楚地知道一个整数值代表一个码位。此外，您可能认为是字符常量的东西在Go中被称为符文常量。表达式的类型和值

```
'⌘'
```

is `rune` with integer value `0x2318`.

的类型和值是符文，整数值为0x2318。

To summarize, here are the salient points:

总结一下，以下是突出的几点：

- Go source code is always UTF-8. Go的源代码始终是UTF-8。
- A string holds arbitrary bytes. 字符串持有任意的字节。
- A string literal, absent byte-level escapes, always holds valid UTF-8 sequences. 在没有字节级转义的情况下，一个字符串字面总是持有有效的UTF-8序列。
- Those sequences represent Unicode code points, called runes. 这些序列代表Unicode代码点，称为符码。
- No guarantee is made in Go that characters in strings are normalized. Go中不保证字符串中的字符是规范化的。

## Range loops - range 循环

Besides the axiomatic detail that Go source code is UTF-8, there’s really only one way that Go treats UTF-8 specially, and that is when using a `for` `range` loop on a string.

除了 Go 源代码是 UTF-8 这一不言而喻的细节外，Go 只有一种方式会特别处理 UTF-8，那就是在字符串上使用 for 范围循环。

We’ve seen what happens with a regular `for` loop. A `for` `range` loop, by contrast, decodes one UTF-8-encoded rune on each iteration. Each time around the loop, the index of the loop is the starting position of the current rune, measured in bytes, and the code point is its value. Here’s an example using yet another handy `Printf` format, `%#U`, which shows the code point’s Unicode value and its printed representation:

我们已经看到了普通for循环的情况。相比之下，for range 循环在每次迭代时都会解码一个UTF-8编码的符文。循环的每一次，循环的索引是当前符文的起始位置，以字节为单位，代码点是其值。下面是一个使用另一种方便的Printf格式的例子，%#U，它显示了代码点的Unicode值和它的打印表示：

```go linenums="1"
    const nihongo = "日本語"
    for index, runeValue := range nihongo {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
    }
```

Run 运行

The output shows how each code point occupies multiple bytes:

输出显示了每个码位是如何占据多个字节的：

```
U+65E5 '日' starts at byte position 0
U+672C '本' starts at byte position 3
U+8A9E '語' starts at byte position 6
```

[Exercise: Put an invalid UTF-8 byte sequence into the string. (How?) What happens to the iterations of the loop?]

[练习。将一个无效的UTF-8字节序列放入字符串。(怎么做?) 循环的迭代会发生什么?］

## Libraries 库

Go’s standard library provides strong support for interpreting UTF-8 text. If a `for` `range` loop isn’t sufficient for your purposes, chances are the facility you need is provided by a package in the library.

Go的标准库对解释UTF-8文本提供了强大的支持。如果for range循环不能满足您的需求，您所需要的功能很可能是由库中的某个包提供的。

The most important such package is [`unicode/utf8`](https://go.dev/pkg/unicode/utf8/), which contains helper routines to validate, disassemble, and reassemble UTF-8 strings. Here is a program equivalent to the `for` `range` example above, but using the `DecodeRuneInString` function from that package to do the work. The return values from the function are the rune and its width in UTF-8-encoded bytes.

最重要的包是unicode/utf8，它包含了验证、反汇编和重新汇编UTF-8字符串的辅助程序。下面是一个等同于上面for range例子的程序，但使用该包中的DecodeRuneInString函数来完成这项工作。该函数的返回值是符文和它的宽度，以UTF-8编码的字节为单位。

```go linenums="1"
    const nihongo = "日本語"
    for i, w := 0, 0; i < len(nihongo); i += w {
        runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
    }
```

Run 运行

Run it to see that it performs the same. The `for` `range` loop and `DecodeRuneInString` are defined to produce exactly the same iteration sequence.

运行它，看它的执行情况是否相同。for range循环和DecodeRuneInString被定义为产生完全相同的迭代序列。

Look at the [documentation](https://go.dev/pkg/unicode/utf8/) for the `unicode/utf8` package to see what other facilities it provides.

看看unicode/utf8包的文档，看看它还提供了哪些设施。

## Conclusion 总结

To answer the question posed at the beginning: Strings are built from bytes so indexing them yields bytes, not characters. A string might not even hold characters. In fact, the definition of "character" is ambiguous and it would be a mistake to try to resolve the ambiguity by defining that strings are made of characters.

为了回答开头提出的问题。字符串是由字节构成的，所以对它们进行索引会产生字节，而不是字符。一个字符串甚至可能不包含字符。事实上，"字符 "的定义是模糊的，如果试图通过定义字符串是由字符构成的来解决这种模糊性，那就是一个错误。

There’s much more to say about Unicode, UTF-8, and the world of multilingual text processing, but it can wait for another post. For now, we hope you have a better understanding of how Go strings behave and that, although they may contain arbitrary bytes, UTF-8 is a central part of their design.

关于Unicode、UTF-8和多语言文本处理的世界，还有很多东西要讲，但可以等到下一篇文章。现在，我们希望您能更好地理解Go字符串的行为，虽然它们可能包含任意的字节，但UTF-8是它们设计的核心部分。
