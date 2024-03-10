+++
title = "源代码表示"
date = 2023-05-17T09:59:21+08:00
weight = 3
description = ""
isCJKLanguage = true
type = "docs"
draft = false
+++
## Source code representation 源代码表示

> 原文：[https://go.dev/ref/spec#Source_code_representation](https://go.dev/ref/spec#Source_code_representation)

Source code is Unicode text encoded in [UTF-8](https://en.wikipedia.org/wiki/UTF-8). The text is not canonicalized, so a single accented code point is distinct from the same character constructed from combining an accent and a letter; those are treated as two code points. For simplicity, this document will use the unqualified term *character* to refer to a Unicode code point in the source text.

​	源代码是以 [UTF-8](https://en.wikipedia.org/wiki/UTF-8)编码的Unicode文本。该文本未被规范化，因此单个重音码点与由重音和字母组合而成的相同字符是不同的；这些被视为两个码点。为了简单起见，本文将使用`非限定术语字符`来指代源文本中的Unicode码点。

Each code point is distinct; for instance, uppercase and lowercase letters are different characters.

​	每个码点都是不同的；例如，大写字母和小写字母是不同的字符。

Implementation restriction: For compatibility with other tools, a compiler may disallow the NUL character (U+0000) in the source text.

​	实现限制：为了与其他工具兼容，编译器可能不允许源文本中出现`NUL`字符（U+0000）。

Implementation restriction: For compatibility with other tools, a compiler may ignore a UTF-8-encoded byte order mark (U+FEFF) if it is the first Unicode code point in the source text. A byte order mark may be disallowed anywhere else in the source.

​	实现限制：为了与其他工具兼容，编译器可以忽略UTF-8编码的字节顺序标记（U+FEFF），如果它是源文本中的第一个Unicode码点。字节顺序标记可能在源代码的其他任何地方被禁用。

### Characters 字符

The following terms are used to denote specific Unicode character categories:

​	以下术语用于表示特定的Unicode字符类别：

```
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point categorized as "Letter" */ .
unicode_digit  = /* a Unicode code point categorized as "Number, decimal digit" */ .
```

In [The Unicode Standard 8.0](https://www.unicode.org/versions/Unicode8.0.0/), Section 4.5 "General Category" defines a set of character categories. Go treats all characters in any of the Letter categories Lu, Ll, Lt, Lm, or Lo as Unicode letters, and those in the Number category Nd as Unicode digits.

​	在 [The Unicode Standard 8.0](https://www.unicode.org/versions/Unicode8.0.0/)中，第4.5节 "`General Category` "定义了一组字符类别。Go将字母类别`Lu`、`Ll`、`Lt`、`Lm`或`Lo`中的所有字符视为`Unicode字母`，将数字类别`Nd`中的字符视为`Unicode数字`。

### Letters and digits 字母和数字

The underscore character `_` (U+005F) is considered a lowercase letter.

​	下划线字符`_`（U+005F）被视为小写字母。

```
letter        = unicode_letter | "_" .
decimal_digit = "0" … "9" .
binary_digit  = "0" | "1" .
octal_digit   = "0" … "7" .
hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
```