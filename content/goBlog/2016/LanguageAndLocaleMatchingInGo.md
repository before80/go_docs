+++
title = "go 中的语言和地域匹配"
weight = 9
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Language and Locale Matching in Go - go 中的语言和地域匹配

https://go.dev/blog/matchlang

Marcel van Lohuizen
9 February 2016

## Introduction 简介

Consider an application, such as a web site, with support for multiple languages in its user interface. When a user arrives with a list of preferred languages, the application must decide which language it should use in its presentation to the user. This requires finding the best match between the languages the application supports and those the user prefers. This post explains why this is a difficult decision and how Go can help.

考虑到一个应用程序，例如一个网站，在其用户界面上支持多种语言。当一个用户带着他喜欢的语言列表来到这里时，应用程序必须决定在向用户展示时应该使用哪种语言。这需要在应用程序支持的语言和用户喜欢的语言之间找到最佳匹配。这篇文章解释了为什么这是一个困难的决定，以及Go如何帮助。

## Language Tags 语言标签

Language tags, also known as locale identifiers, are machine-readable identifiers for the language and/or dialect being used. The most common reference for them is the IETF BCP 47 standard, and that is the standard the Go libraries follow. Here are some examples of BCP 47 language tags and the language or dialect they represent.

语言标签，也被称为地域标识符，是机器可读的语言和/或使用的方言的标识符。它们最常见的参考是 IETF BCP 47 标准，这也是 Go 库所遵循的标准。下面是一些 BCP 47 语言标签的例子以及它们所代表的语言或方言。

| Tag         | Description                              |
| ----------- | ---------------------------------------- |
| en          | English                                  |
| en-US       | American English 美式英语                |
| cmn         | Mandarin Chinese                         |
| zh          | Chinese, typically Mandarin              |
| nl          | Dutch                                    |
| nl-BE       | Flemish                                  |
| es-419      | Latin American Spanish                   |
| az, az-Latn | both Azerbaijani written in Latin script |
| az-Arab     | Azerbaijani written in Arabic            |

The general form of the language tag is a language code ("en", "cmn", "zh", "nl", "az" above) followed by an optional subtag for script ("-Arab"), region ("-US", "-BE", "-419"), variants ("-oxendict" for Oxford English Dictionary spelling), and extensions ("-u-co-phonebk" for phone-book sorting). The most common form is assumed if a subtag is omitted, for instance "az-Latn-AZ" for "az".

语言标签的一般形式是一个语言代码（"en"，"cmn"，"zh"，"nl"，"az "以上），后面是一个可选的副标签，用于文字（"-Arab"），地区（"-US"，"-BE"，"-419"），变体（"-oxendict "用于牛津英语词典拼写），和扩展（"-u-co-phonebk "用于电话簿排序）。如果省略了一个子标签，则假定最常见的形式，例如 "az-Latn-AZ "表示 "az"。

The most common use of language tags is to select from a set of system-supported languages according to a list of the user’s language preferences, for example deciding that a user who prefers Afrikaans would be best served (assuming Afrikaans is not available) by the system showing Dutch. Resolving such matches involves consulting data on mutual language comprehensibility.

语言标签最常见的用途是根据用户的语言偏好列表从一套系统支持的语言中进行选择，例如，决定一个喜欢南非荷兰语的用户最好由系统显示荷兰语（假设南非荷兰语不可用）。解决这样的匹配需要查阅相互之间的语言可理解性数据。

The tag resulting from this match is subsequently used to obtain language-specific resources such as translations, sorting order, and casing algorithms. This involves a different kind of matching. For example, as there is no specific sorting order for Portuguese, a collate package may fall back to the sorting order for the default, or "root", language.

这种匹配产生的标签随后被用来获得特定语言的资源，如翻译、排序和套管的算法。这涉及到一种不同的匹配。例如，由于葡萄牙语没有特定的排序顺序，整理包可能会退回到默认或 "根 "语言的排序顺序。

## The Messy Nature of Matching Languages 匹配语言的混乱性质

Handling language tags is tricky. This is partly because the boundaries of human languages are not well defined and partly because of the legacy of evolving language tag standards. In this section we will show some of the messy aspects of handling language tags.

处理语言标签是很棘手的。部分原因是人类语言的边界没有被很好地定义，部分原因是不断发展的语言标签标准的遗留问题。在本节中，我们将展示处理语言标签的一些混乱的方面。

*Tags with different language codes can indicate the same language*

具有不同语言代码的标签可以表示同一种语言

For historical and political reasons, many language codes have changed over time, leaving languages with an older legacy code as well as a new one. But even two current codes may refer to the same language. For example, the official language code for Mandarin is "cmn", but "zh" is by far the most commonly used designator for this language. The code "zh" is officially reserved for a so called macro language, identifying the group of Chinese languages. Tags for macro languages are often used interchangeably with the most-spoken language in the group.

由于历史和政治的原因，许多语言代码随着时间的推移而改变，使语言有一个旧的遗留代码以及一个新的代码。但是，即使是两个当前的代码也可能指的是同一种语言。例如，普通话的官方语言代码是 "cmn"，但 "zh "是迄今为止该语言最常用的代号。编码 "zh "是官方为所谓的宏观语言保留的，用于识别中国的语言组。宏观语言的标签通常与该组语言中使用频率最高的语言交替使用。

*Matching language code alone is not sufficient*

仅仅匹配语言代码是不够的

Azerbaijani ("az"), for example, is written in different scripts depending on the country in which it is spoken: "az-Latn" for Latin (the default script), "az-Arab" for Arabic, and "az-Cyrl" for Cyrillic. If you replace "az-Arab" with just "az", the result will be in Latin script and may not be understandable to a user who only knows the Arabic form.

例如，阿塞拜疆语（"az"），根据其使用国家的不同，用不同的文字书写："az-Latn "代表拉丁语（默认文字），"az-Arab "代表阿拉伯语，"az-Cyrl "代表西里尔语。如果您把 "az-Arab "替换成 "az"，结果将是拉丁文，对于只知道阿拉伯语形式的用户来说，可能无法理解。

Also different regions may imply different scripts. For example: "zh-TW" and "zh-SG" respectively imply the use of Traditional and Simplified Han. As another example, "sr" (Serbian) defaults to Cyrillic script, but "sr-RU" (Serbian as written in Russia) implies the Latin script! A similar thing can be said for Kyrgyz and other languages.

另外，不同地区可能意味着不同的脚本。比如说。"zh-TW "和 "zh-SG "分别意味着使用繁体和简体汉字。再比如，"sr"（塞尔维亚语）默认为西里尔字体，但 "sr-RU"（在俄罗斯书写的塞尔维亚语）则意味着拉丁文体 对于吉尔吉斯语和其他语言也有类似的情况。

If you ignore subtags, you might as well present Greek to the user.

如果您忽略了副标记，您还不如把希腊语呈现给用户。

*The best match might be a language not listed by the user*

最好的匹配可能是用户没有列出的语言

The most common written form of Norwegian ("nb") looks an awful lot like Danish. If Norwegian is not available, Danish may be a good second choice. Similarly, a user requesting Swiss German ("gsw") will likely be happy to be presented German ("de"), though the converse is far from true. A user requesting Uygur may be happier to fall back to Chinese than to English. Other examples abound. If a user-requested language is not supported, falling back to English is often not the best thing to do.

挪威语最常见的书面形式（"nb"）看起来非常像丹麦语。如果挪威语不可用，丹麦语可能是一个很好的第二选择。同样，一个要求使用瑞士德语（"gsw"）的用户可能会很高兴得到德语（"de"），尽管反过来也是如此。一个要求使用维吾尔语的用户可能会更乐意回到中文而不是英语。其他的例子比比皆是。如果一个用户要求的语言不被支持，退回到英语往往不是最好的办法。

*The choice of language decides more than translation*

语言的选择比翻译更有决定性

Suppose a user asks for Danish, with German as a second choice. If an application chooses German, it must not only use German translations but also use German (not Danish) collation. Otherwise, for example, a list of animals might sort "Bär" before "Äffin".

假设一个用户要求使用丹麦语，并将德语作为第二选择。如果一个应用程序选择了德语，它不仅要使用德语翻译，还要使用德语（而不是丹麦语）整理。否则，例如，一个动物列表可能会将 "Bär "排序在 "Äffin "之前。

Selecting a supported language given the user’s preferred languages is like a handshaking algorithm: first you determine which protocol to communicate in (the language) and then you stick with this protocol for all communication for the duration of a session.

考虑到用户的首选语言，选择支持的语言就像握手算法：首先确定用哪种协议进行通信（语言），然后在会话期间坚持用这种协议进行所有通信。

*Using a "parent" of a language as fallback is non-trivial*

使用一种语言的 "父 "作为后备语言是不容易的。

Suppose your application supports Angolan Portuguese ("pt-AO"). Packages in [golang.org/x/text](https://golang.org/x/text), like collation and display, may not have specific support for this dialect. The correct course of action in such cases is to match the closest parent dialect. Languages are arranged in a hierarchy, with each specific language having a more general parent. For example, the parent of "en-GB-oxendict" is "en-GB", whose parent is "en", whose parent is the undefined language "und", also known as the root language. In the case of collation, there is no specific collation order for Portugese, so the collate package will select the sorting order of the root language. The closest parent to Angolan Portuguese supported by the display package is European Portuguese ("pt-PT") and not the more obvious "pt", which implies Brazilian.

假设您的应用程序支持安哥拉葡萄牙语（"pt-AO"）。golang.org/x/text中的包，如整理和显示，可能没有对这种方言的具体支持。在这种情况下，正确的做法是匹配最接近的父方言。语言是按层次排列的，每一种特定的语言都有一个更普遍的父方言。例如，"en-GB-oxendict "的父级是 "en-GB"，其父级是 "en"，其父级是未定义语言 "und"，也被称为根语言。在整理的情况下，葡萄牙语没有特定的整理顺序，所以整理包会选择根语言的排序顺序。显示包支持的与安哥拉葡萄牙语最接近的父语是欧洲葡萄牙语（"pt-PT"），而不是更明显的 "pt"，后者意味着巴西语。

In general, parent relationships are non-trivial. To give a few more examples, the parent of "es-CL" is "es-419", the parent of "zh-TW" is "zh-Hant", and the parent of "zh-Hant" is "und". If you compute the parent by simply removing subtags, you may select a "dialect" that is incomprehensible to the user.

一般来说，父系关系是不简单的。再举几个例子，"es-CL "的父级是 "es-419"，"zh-TW "的父级是 "zh-Hant"，而 "zh-Hant "的父级是 "und"。如果您通过简单地去除子标记来计算父语言，您可能会选择一种用户无法理解的 "方言"。

## Language Matching in Go - Go中的语言匹配

The Go package [golang.org/x/text/language](https://golang.org/x/text/language) implements the BCP 47 standard for language tags and adds support for deciding which language to use based on data published in the Unicode Common Locale Data Repository (CLDR).

Go软件包golang.org/x/text/language实现了BCP 47标准的语言标签，并增加了对根据Unicode通用地域数据存储库（CLDR）中公布的数据决定使用哪种语言的支持。

Here is a sample program, explained below, matching a user’s language preferences against an application’s supported languages:

下面是一个示例程序，解释如下，将用户的语言偏好与应用程序的支持语言相匹配：

```go linenums="1"
// +build OMIT

package main

import (
    "fmt"

    "golang.org/x/text/language"
    "golang.org/x/text/language/display"
)

var userPrefs = []language.Tag{
    language.Make("gsw"), // Swiss German
    language.Make("fr"),  // French
}

var serverLangs = []language.Tag{
    language.AmericanEnglish, // en-US fallback
    language.German,          // de
}

var matcher = language.NewMatcher(serverLangs)

func main() {
    tag, index, confidence := matcher.Match(userPrefs...)

    fmt.Printf("best match: %s (%s) index=%d confidence=%v\n",
        display.English.Tags().Name(tag),
        display.Self.Name(tag),
        index, confidence)
    // best match: German (Deutsch) index=1 confidence=High
}
```

### Creating Language Tags 创建语言标签

The simplest way to create a language.Tag from a user-given language code string is with language.Make. It extracts meaningful information even from malformed input. For example, "en-USD" will result in "en" even though USD is not a valid subtag.

从用户给定的语言代码字符串中创建一个language.Tag的最简单方法是使用language.Make。它甚至可以从畸形的输入中提取有意义的信息。例如，"en-USD "将产生 "en"，尽管USD不是一个有效的子标签。

Make doesn’t return an error. It is common practice to use the default language if an error occurs anyway so this makes it more convenient. Use Parse to handle any error manually.

Make不会返回一个错误。如果发生错误，通常的做法是使用默认的语言，所以这使它更方便。使用Parse来手动处理任何错误。

The HTTP Accept-Language header is often used to pass a user’s desired languages. The ParseAcceptLanguage function parses it into a slice of language tags, ordered by preference.

HTTP Accept-Language标头经常被用来传递用户想要的语言。ParseAcceptLanguage函数将其解析为一个语言标签片，按偏好排序。

By default, the language package does not canonicalize tags. For example, it does not follow the BCP 47 recommendation of eliminating scripts if it is the common choice in the "overwhelming majority". It similarly ignores CLDR recommendations: "cmn" is not replaced by "zh" and "zh-Hant-HK" is not simplified to "zh-HK". Canonicalizing tags may throw away useful information about user intent. Canonicalization is handled in the Matcher instead. A full array of canonicalization options are available if the programmer still desires to do so.

默认情况下，语言包不对标签进行规范化处理。例如，它不遵循BCP 47的建议，即如果脚本是 "绝大多数人 "的共同选择，则应将其消除。同样，它也忽略了CLDR的建议。"cmn "没有被 "zh "取代，"zh-Hant-HK "没有被简化为 "zh-HK"。对标签进行规范化处理可能会丢掉关于用户意图的有用信息。冠词化在匹配器中处理。如果程序员仍然希望这样做的话，可以使用一系列的规范化选项。

### Matching User-Preferred Languages to Supported Languages 将用户偏好的语言与支持的语言相匹配

A Matcher matches user-preferred languages to supported languages. Users are strongly advised to use it if they don’t want to deal with all the intricacies of matching languages.

匹配器将用户偏好的语言与支持的语言相匹配。如果用户不想处理所有复杂的语言匹配问题，我们强烈建议他们使用它。

The Match method may pass through user settings (from BCP 47 extensions) from the preferred tags to the selected supported tag. It is therefore important that the tag returned by Match is used to obtain language-specific resources. For example, "de-u-co-phonebk" requests phone-book ordering for German. The extension is ignored for matching, but is used by the collate package to select the respective sorting order variant.

匹配方法可能会将用户设置（来自BCP 47扩展）从首选标记传递到所选的支持标记。因此，重要的是，Match返回的标签要用于获得特定语言的资源。例如，"de-u-co-phonebk "请求为德语订购电话簿。扩展名在匹配时被忽略，但被collate包用来选择相应的排序变量。

A Matcher is initialized with the languages supported by an application, which are usually the languages for which there are translations. This set is typically fixed, allowing a matcher to be created at startup. Matcher is optimized to improve the performance of Match at the expense of initialization cost.

匹配器被初始化为一个应用程序所支持的语言，这通常是有翻译的语言。这个集合通常是固定的，允许在启动时创建一个匹配器。匹配器经过优化，以初始化成本为代价提高匹配器的性能。

The language package provides a predefined set of the most commonly used language tags that can be used for defining the supported set. Users generally don’t have to worry about the exact tags to pick for supported languages. For example, AmericanEnglish ("en-US") may be used interchangeably with the more common English ("en"), which defaults to American. It is all the same for the Matcher. An application may even add both, allowing for more specific American slang for "en-US".

语言包提供了一个预定义的最常用的语言标签集，可用于定义支持的集。用户一般不需要担心为支持的语言挑选确切的标签。例如，AmericanEnglish（"en-US"）可以与更常见的English（"en"）互换使用，后者默认为美国人。这对匹配器来说都是一样的。一个应用程序甚至可以同时添加，允许 "en-US "使用更具体的美国俚语。

### Matching Example 匹配实例

Consider the following Matcher and lists of supported languages:

考虑一下下面的匹配器和支持的语言列表：

```go linenums="1"
var supported = []language.Tag{
    language.AmericanEnglish,    // en-US: first language is fallback
    language.German,             // de
    language.Dutch,              // nl
    language.Portuguese          // pt (defaults to Brazilian)
    language.EuropeanPortuguese, // pt-pT
    language.Romanian            // ro
    language.Serbian,            // sr (defaults to Cyrillic script)
    language.SerbianLatin,       // sr-Latn
    language.SimplifiedChinese,  // zh-Hans
    language.TraditionalChinese, // zh-Hant
}
var matcher = language.NewMatcher(supported)
```

Let’s look at the matches against this list of supported languages for various user preferences.

让我们看看针对不同用户偏好的支持语言列表的匹配情况。

For a user preference of "he" (Hebrew), the best match is "en-US" (American English). There is no good match, so the matcher uses the fallback language (the first in the supported list).

对于 "he"（希伯来语）的用户偏好，最佳匹配是 "en-US"（美国英语）。没有好的匹配，所以匹配器使用后备语言（支持列表中的第一个）。

For a user preference of "hr" (Croatian), the best match is "sr-Latn" (Serbian with Latin script), because, once they are written in the same script, Serbian and Croatian are mutually intelligible.

对于用户偏好的 "hr"（克罗地亚语），最好的匹配是 "sr-Latn"（带拉丁字母的塞尔维亚语），因为一旦它们用相同的字体书写，塞尔维亚语和克罗地亚语是可以相互理解的。

For a user preference of "ru, mo" (Russian, then Moldavian), the best match is "ro" (Romanian), because Moldavian is now canonically classified as "ro-MD" (Romanian in Moldova).

对于用户偏好的 "ru, mo"（俄语，然后是摩尔多瓦语），最佳匹配是 "ro"（罗马尼亚语），因为摩尔多瓦语现在被归类为 "ro-MD"（摩尔多瓦的罗马尼亚语）。

For a user preference of "zh-TW" (Mandarin in Taiwan), the best match is "zh-Hant" (Mandarin written in Traditional Chinese), not "zh-Hans" (Mandarin written in Simplified Chinese).

对于用户偏好的 "zh-TW"（台湾的普通话），最佳匹配是 "zh-Hant"（用繁体中文书写的普通话），而不是 "zh-Hans"（用简体中文书写的普通话）。

For a user preference of "af, ar" (Afrikaans, then Arabic), the best match is "nl" (Dutch). Neither preference is supported directly, but Dutch is a significantly closer match to Afrikaans than the fallback language English is to either.

对于 "af, ar"（南非荷兰语，然后是阿拉伯语）的用户偏好，最佳匹配是 "nl"（荷兰语）。这两种偏好都不被直接支持，但荷兰语与南非荷兰语的匹配度明显高于后备语言英语与两者的匹配度。

For a user preference of "pt-AO, id" (Angolan Portuguese, then Indonesian), the best match is "pt-PT" (European Portuguese), not "pt" (Brazilian Portuguese).

对于 "pt-AO, id"（安哥拉葡萄牙语，然后是印度尼西亚语）的用户偏好，最佳匹配是 "pt-PT"（欧洲葡萄牙语），而不是 "pt"（巴西葡萄牙语）。

For a user preference of "gsw-u-co-phonebk" (Swiss German with phone-book collation order), the best match is "de-u-co-phonebk" (German with phone-book collation order). German is the best match for Swiss German in the server’s language list, and the option for phone-book collation order has been carried over.

对于用户偏好的 "gsw-u-co-phonebk"（采用电话簿排序的瑞士德语），最佳匹配是 "de-u-co-phonebk"（采用电话簿排序的德语）。在服务器的语言列表中，德语是与瑞士德语最匹配的，电话簿整理顺序的选项被延续下来。

### Confidence Scores 信任度评分

Go uses coarse-grained confidence scoring with rule-based elimination. A match is classified as Exact, High (not exact, but no known ambiguity), Low (probably the correct match, but maybe not), or No. In case of multiple matches, there is a set of tie-breaking rules that are executed in order. The first match is returned in the case of multiple equal matches. These confidence scores may be useful, for example, to reject relatively weak matches. They are also used to score, for example, the most likely region or script from a language tag.

Go使用粗粒度的置信度评分和基于规则的消除法。匹配被分为精确、高（不精确，但没有已知的歧义）、低（可能是正确的匹配，但可能不是）或否。在有多个相同匹配的情况下，将返回第一个匹配。这些信心分数可能是有用的，例如，拒绝相对较弱的匹配。它们也被用来打分，例如，从一个语言标签中找出最可能的地区或文字。

Implementations in other languages often use more fine-grained, variable-scale scoring. We found that using coarse-grained scoring in the Go implementation ended up simpler to implement, more maintainable, and faster, meaning that we could handle more rules.

其他语言的实施通常使用更精细的、可变尺度的评分。我们发现，在Go的实现中使用粗粒度的评分，最终实现起来更简单，更容易维护，而且速度更快，这意味着我们可以处理更多的规则。

### Displaying Supported Languages 显示支持的语言

The [golang.org/x/text/language/display](https://golang.org/x/text/language/display) package allows naming language tags in many languages. It also contains a "Self" namer for displaying a tag in its own language.

golang.org/x/text/language/display包允许以许多语言命名语言标签。它还包含一个 "Self "命名器，用于在自己的语言中显示一个标签。

For example:

比如说：

```go linenums="1"
    var supported = []language.Tag{
        language.English,            // en
        language.French,             // fr
        language.Dutch,              // nl
        language.Make("nl-BE"),      // nl-BE
        language.SimplifiedChinese,  // zh-Hans
        language.TraditionalChinese, // zh-Hant
        language.Russian,            // ru
    }

    en := display.English.Tags()
    for _, t := range supported {
        fmt.Printf("%-20s (%s)\n", en.Name(t), display.Self.Name(t))
    }
```

prints 打印

```
English              (English)
French               (français)
Dutch                (Nederlands)
Flemish              (Vlaams)
Simplified Chinese   (简体中文)
Traditional Chinese  (繁體中文)
Russian              (русский)
```

In the second column, note the differences in capitalization, reflecting the rules of the respective language.

在第二栏中，注意大写字母的差异，反映了各自语言的规则。

## Conclusion 总结

At first glance, language tags look like nicely structured data, but because they describe human languages, the structure of relationships between language tags is actually quite complex. It is often tempting, especially for English-speaking programmers, to write ad-hoc language matching using nothing other than string manipulation of the language tags. As described above, this can produce awful results.

乍一看，语言标签看起来是很好的结构化数据，但由于它们描述的是人类语言，语言标签之间的关系结构实际上是相当复杂的。特别是对于讲英语的程序员来说，编写临时性的语言匹配往往是很诱人的，除了对语言标记进行字符串操作外，什么都不用。如上所述，这可能会产生糟糕的结果。

Go’s [golang.org/x/text/language](https://golang.org/x/text/language) package solves this complex problem while still presenting a simple, easy-to-use API. Enjoy.

Go的golang.org/x/text/language包解决了这个复杂的问题，同时还提供了一个简单、易于使用的API。请享用。
