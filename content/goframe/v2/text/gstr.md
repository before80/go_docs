+++
title = "gstr"
date = 2024-03-21T17:58:55+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/text/gstr](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/text/gstr)

Package gstr provides functions for string handling.

​	软件包 gstr 提供字符串处理函数。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/text/gstr/gstr.go#L10)

```go
const (
	// NotFoundIndex is the position index for string not found in searching functions.
	NotFoundIndex = -1
)
```

## 变量

This section is empty.

## 函数

#### func AddSlashes

```go
func AddSlashes(str string) string
```

AddSlashes quotes with slashes `\` for chars: ‘".

​	AddSlashes 引号与斜杠 `\` 表示字符：'“。

##### Example

``` go
```

#### func CaseCamel

```go
func CaseCamel(s string) string
```

CaseCamel converts a string to CamelCase.

​	CaseCamel 将字符串转换为 CamelCase。

Example: CaseCamel(“any_kind_of_string”) -> AnyKindOfString

​	示例：CaseCamel（“any_kind_of_string”） -> AnyKindOfString

##### Example

``` go
```

#### func CaseCamelLower

```go
func CaseCamelLower(s string) string
```

CaseCamelLower converts a string to lowerCamelCase.

​	CaseCamelLower 将字符串转换为 lowerCamelCase。

Example: CaseCamelLower(“any_kind_of_string”) -> anyKindOfString

​	示例：CaseCamelLower（“any_kind_of_string”） -> anyKindOfString

##### Example

``` go
```

#### func CaseConvert <-2.5.7

```go
func CaseConvert(s string, caseType CaseType) string
```

CaseConvert converts a string to the specified naming convention. Use CaseTypeMatch to match the case type from string.

​	CaseConvert 将字符串转换为指定的命名约定。使用 CaseTypeMatch 匹配字符串中的案例类型。

#### func CaseDelimited

```go
func CaseDelimited(s string, del byte) string
```

CaseDelimited converts a string to snake.case.delimited.

​	CaseDelimited 将字符串转换为 snake.case.delimited。

Example: CaseDelimited(“AnyKindOfString”, ‘.’) -> any.kind.of.string

​	示例：CaseDelimited（“AnyKindOfString”， '.'） -> any.kind.of.string

##### Example

``` go
```

#### func CaseDelimitedScreaming

```go
func CaseDelimitedScreaming(s string, del uint8, screaming bool) string
```

CaseDelimitedScreaming converts a string to DELIMITED.SCREAMING.CASE or delimited.screaming.case.

​	CaseDelimitedScreaming 将字符串转换为 DELIMITED。尖叫。CASE 或 delimited.screaming.case。

Example: CaseDelimitedScreaming(“AnyKindOfString”, ‘.’) -> ANY.KIND.OF.STRING

​	示例：CaseDelimitedScreaming（“AnyKindOfString”， '.'） -> ANY.类。之。字符串

##### Example

``` go
```

#### func CaseKebab

```go
func CaseKebab(s string) string
```

CaseKebab converts a string to kebab-case.

​	CaseKebab 将字符串转换为 kebab-case。

Example: CaseKebab(“AnyKindOfString”) -> any-kind-of-string

​	示例：CaseKebab（“AnyKindOfString”） -> any-kind-of-string

##### Example

``` go
```

#### func CaseKebabScreaming

```go
func CaseKebabScreaming(s string) string
```

CaseKebabScreaming converts a string to KEBAB-CASE-SCREAMING.

​	CaseKebabScreaming 将字符串转换为 KEBAB-CASE-SCREAMING。

Example: CaseKebab(“AnyKindOfString”) -> ANY-KIND-OF-STRING

​	示例：CaseKebab（“AnyKindOfString”） -> ANY-KIND-OF-STRING

##### Example

``` go
```

#### func CaseSnake

```go
func CaseSnake(s string) string
```

CaseSnake converts a string to snake_case.

​	CaseSnake 将字符串转换为snake_case。

Example: CaseSnake(“AnyKindOfString”) -> any_kind_of_string

​	示例：CaseSnake（“AnyKindOfString”） -> any_kind_of_string

##### Example

``` go
```

#### func CaseSnakeFirstUpper

```go
func CaseSnakeFirstUpper(word string, underscore ...string) string
```

CaseSnakeFirstUpper converts a string like “RGBCodeMd5” to “rgb_code_md5”. TODO for efficiency should change regexp to traversing string in future.

​	CaseSnakeFirstUpper 将类似“RGBCodeMd5”的字符串转换为“rgb_code_md5”。为了提高效率，TODO 将来应将正则表达式更改为遍历字符串。

Example: CaseSnakeFirstUpper(“RGBCodeMd5”) -> rgb_code_md5

​	示例：CaseSnakeFirstUpper（“RGBCodeMd5”） -> rgb_code_md5

##### Example

``` go
```

#### func CaseSnakeScreaming

```go
func CaseSnakeScreaming(s string) string
```

CaseSnakeScreaming converts a string to SNAKE_CASE_SCREAMING.

​	CaseSnakeScreaming 将字符串转换为SNAKE_CASE_SCREAMING。

Example: CaseSnakeScreaming(“AnyKindOfString”) -> ANY_KIND_OF_STRING

​	示例：CaseSnakeScreaming（“AnyKindOfString”） -> ANY_KIND_OF_STRING

##### Example

``` go
```

#### func Chr

```go
func Chr(ascii int) string
```

Chr return the ascii string of a number(0-255).

​	Chr 返回数字 （0-255） 的 ascii 字符串。

Example: Chr(65) -> “A”

​	示例：chr（65） -> “A”

##### Example

``` go
```

#### func ChunkSplit

```go
func ChunkSplit(body string, chunkLen int, end string) string
```

ChunkSplit splits a string into smaller chunks. Can be used to split a string into smaller chunks which is useful for e.g. converting BASE64 string output to match [RFC 2045](https://rfc-editor.org/rfc/rfc2045.html) semantics. It inserts end every chunkLen characters. It considers parameter `body` and `end` as unicode string.

​	ChunkSplit 将字符串拆分为更小的块。可用于将字符串拆分为更小的块，例如，转换 BASE64 字符串输出以匹配 RFC 2045 语义。它插入每个 chunkLen 字符的结尾。它考虑参数 `body` 和 `end` unicode 字符串。

##### Example

``` go
```

#### func Compare

```go
func Compare(a, b string) int
```

Compare returns an integer comparing two strings lexicographically. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.

​	Compare 返回一个按字典比较两个字符串的整数。如果 a==b，则结果为 0，如果 a < b，则结果为 -1，如果 a > b，则为 +1。

##### Example

``` go
```

#### func CompareVersion

```go
func CompareVersion(a, b string) int
```

CompareVersion compares `a` and `b` as standard GNU version.

​	CompareVersion 比较 `a` 和 `b` 作为标准 GNU 版本。

It returns 1 if `a` > `b`.

​	如果 `a` > `b` ，则返回 1。

It returns -1 if `a` < `b`.

​	如果 `a` < `b` ，则返回 -1。

It returns 0 if `a` = `b`.

​	如果 `a` = `b` ，则返回 0。

GNU standard version is like: v1.0 1 1.0.0 v1.0.1 v2.10.8 10.2.0 etc.

​	GNU标准版本如下：v1.0、1、1.0.0、v1.0.1、v2.10.8、10.2.0等。

##### Example

``` go
```

#### func CompareVersionGo

```go
func CompareVersionGo(a, b string) int
```

CompareVersionGo compares `a` and `b` as standard Golang version.

​	CompareVersionGo 将 `a` 和 `b` 作为标准 Golang 版本进行比较。

It returns 1 if `a` > `b`.

​	如果 `a` > `b` ，则返回 1。

It returns -1 if `a` < `b`.

​	如果 `a` < `b` ，则返回 -1。

It returns 0 if `a` = `b`.

​	如果 `a` = `b` ，则返回 0。

Golang standard version is like: 1.0.0 v1.0.1 v2.10.8 10.2.0 v0.0.0-20190626092158-b2ccc519800e v1.12.2-0.20200413154443-b17e3a6804fa v4.20.0+incompatible etc.

​	Golang标准版是这样的：1.0.0 v1.0.1 v2.10.8 10.2.0 v0.0.0-20190626092158-b2ccc519800e v1.12.2-0.20200413154443-b17e3a6804fa v4.20.0+不兼容等。

Docs: https://go.dev/doc/modules/version-numbers

​	文档：https://go.dev/doc/modules/version-numbers

##### Example

``` go
```

#### func Contains

```go
func Contains(str, substr string) bool
```

Contains reports whether `substr` is within `str`, case-sensitively.

​	包含是否 `substr` 在 、 `str` 区分大小写的报表。

##### Example

``` go
```

#### func ContainsAny

```go
func ContainsAny(s, chars string) bool
```

ContainsAny reports whether any Unicode code points in `chars` are within `s`.

​	ContainsAny 报告中 `chars` 是否有任何 Unicode 码位在 `s` 内。

##### Example

``` go
```

#### func ContainsI

```go
func ContainsI(str, substr string) bool
```

ContainsI reports whether substr is within str, case-insensitively.

​	ContainsI 报告 substr 是否在 str 内，不区分大小写。

##### Example

``` go
```

#### func Count

```go
func Count(s, substr string) int
```

Count counts the number of `substr` appears in `s`. It returns 0 if no `substr` found in `s`.

​	计数 计算 中 `s` 出现的次数 `substr` 。如果在 中 `s` 找不到 `substr` ，则返回 0。

##### Example

``` go
```

#### func CountChars

```go
func CountChars(str string, noSpace ...bool) map[string]int
```

CountChars returns information about chars’ count used in a string. It considers parameter `str` as unicode string.

​	CountChars 返回有关字符串中使用的字符计数的信息。它将参数 `str` 视为 unicode 字符串。

##### Example

``` go
```

#### func CountI

```go
func CountI(s, substr string) int
```

CountI counts the number of `substr` appears in `s`, case-insensitively. It returns 0 if no `substr` found in `s`.

​	CountI 不区分大小写地计算 中 `s` 出现的 `substr` 次数。如果在 中 `s` 找不到 `substr` ，则返回 0。

##### Example

``` go
```

#### func CountWords

```go
func CountWords(str string) map[string]int
```

CountWords returns information about words’ count used in a string. It considers parameter `str` as unicode string.

​	CountWords 返回有关字符串中使用的单词计数的信息。它将参数 `str` 视为 unicode 字符串。

##### Example

``` go
```

#### func Equal

```go
func Equal(a, b string) bool
```

Equal reports whether `a` and `b`, interpreted as UTF-8 strings, are equal under Unicode case-folding, case-insensitively.

​	相等报告 `a` 和 `b` （解释为 UTF-8 字符串）在 Unicode 大小写折叠下是否相等，不区分大小写。

##### Example

``` go
```

#### func Explode

```go
func Explode(delimiter, str string) []string
```

Explode splits string `str` by a string `delimiter`, to an array. See http://php.net/manual/en/function.explode.php.

​	Explode 将字符串 `str` 拆分为一个字符串 `delimiter` ，以转换为数组。请参见 http://php.net/manual/en/function.explode.php。

##### Example

``` go
```

#### func Fields

```go
func Fields(str string) []string
```

Fields returns the words used in a string as slice.

​	Fields 将字符串中使用的单词作为切片返回。

##### Example

``` go
```

#### func HasPrefix

```go
func HasPrefix(s, prefix string) bool
```

HasPrefix tests whether the string s begins with prefix.

​	HasPrefix 测试字符串 s 是否以前缀开头。

##### Example

``` go
```

#### func HasSuffix

```go
func HasSuffix(s, suffix string) bool
```

HasSuffix tests whether the string s ends with suffix.

​	HasSuffix 测试字符串 s 是否以后缀结尾。

##### Example

``` go
```

#### func HideStr

```go
func HideStr(str string, percent int, hide string) string
```

HideStr replaces part of the string `str` to `hide` by `percentage` from the `middle`. It considers parameter `str` as unicode string.

​	HideStr 将部分字符串 `str` 替换 `percentage` 为 `hide` by `middle` from the .它将参数 `str` 视为 unicode 字符串。

##### Example

``` go
```

#### func Implode

```go
func Implode(glue string, pieces []string) string
```

Implode joins array elements `pieces` with a string `glue`. http://php.net/manual/en/function.implode.php

​	Implode 使用字符串 `glue` 连接数组元素 `pieces` 。http://php.net/manual/en/function.implode.php

##### Example

``` go
```

#### func InArray

```go
func InArray(a []string, s string) bool
```

InArray checks whether string `s` in slice `a`.

​	InArray 检查 slice `a` 中的字符串 `s` 是否为 。

##### Example

``` go
```

#### func IsGNUVersion <-2.1.2

```go
func IsGNUVersion(version string) bool
```

IsGNUVersion checks and returns whether given `version` is valid GNU version string.

​	IsGNUVersion 检查并返回给定 `version` 的 GNU 版本字符串是否有效。

#### func IsLetterLower

```go
func IsLetterLower(b byte) bool
```

IsLetterLower tests whether the given byte b is in lower case.

​	IsLetterLower 测试给定的字节 b 是否为小写。

##### Example

``` go
```

#### func IsLetterUpper

```go
func IsLetterUpper(b byte) bool
```

IsLetterUpper tests whether the given byte b is in upper case.

​	IsLetterUpper 测试给定的字节 b 是否为大写。

##### Example

``` go
```

#### func IsNumeric

```go
func IsNumeric(s string) bool
```

IsNumeric tests whether the given string s is numeric.

​	IsNumeric 测试给定字符串 s 是否为数值。

##### Example

``` go
```

#### func IsSubDomain

```go
func IsSubDomain(subDomain string, mainDomain string) bool
```

IsSubDomain checks whether `subDomain` is sub-domain of mainDomain. It supports ‘*’ in `mainDomain`.

​	IsSubDomain 检查是否 `subDomain` 是 mainDomain 的子域。它支持 中的 `mainDomain` “*”。

##### Example

``` go
```

#### func Join

```go
func Join(array []string, sep string) string
```

Join concatenates the elements of `array` to create a single string. The separator string `sep` is placed between elements in the resulting string.

​	Join 将 的 `array` 元素连接起来以创建单个字符串。分隔符字符串 `sep` 放置在生成的字符串中的元素之间。

##### Example

``` go
```

#### func JoinAny

```go
func JoinAny(array interface{}, sep string) string
```

JoinAny concatenates the elements of `array` to create a single string. The separator string `sep` is placed between elements in the resulting string.

​	JoinAny 将 的 `array` 元素连接起来以创建单个字符串。分隔符字符串 `sep` 放置在生成的字符串中的元素之间。

The parameter `array` can be any type of slice, which be converted to string array.

​	该参数 `array` 可以是任何类型的切片，这些切片可以转换为字符串数组。

##### Example

``` go
```

#### func LcFirst

```go
func LcFirst(s string) string
```

LcFirst returns a copy of the string s with the first letter mapped to its lower case.

​	LcFirst 返回字符串 s 的副本，其中第一个字母映射到其小写字母。

##### Example

``` go
```

#### func LenRune

```go
func LenRune(str string) int
```

LenRune returns string length of unicode.

​	LenRune 返回 unicode 的字符串长度。

##### Example

``` go
```

#### func Levenshtein

```go
func Levenshtein(str1, str2 string, costIns, costRep, costDel int) int
```

Levenshtein calculates Levenshtein distance between two strings. costIns: Defines the cost of insertion. costRep: Defines the cost of replacement. costDel: Defines the cost of deletion. See http://php.net/manual/en/function.levenshtein.php.

​	Levenshtein 计算两根弦之间的 Levenshtein 距离。costIns：定义插入成本。costRep：定义更换成本。costDel：定义删除成本。请参见 http://php.net/manual/en/function.levenshtein.php。

##### Example

``` go
```

#### func List2 <-2.5.5

```go
func List2(str, delimiter string) (part1, part2 string)
```

List2 Split the `str` with `delimiter` and returns the result as two parts string.

​	List2 拆分 `str` with `delimiter` 并将结果作为两部分字符串返回。

#### func List3 <-2.5.5

```go
func List3(str, delimiter string) (part1, part2, part3 string)
```

List3 Split the `str` with `delimiter` and returns the result as three parts string.

​	List3 拆分 `str` with `delimiter` 并将结果作为三部分字符串返回。

#### func ListAndTrim2 <-2.5.5

```go
func ListAndTrim2(str, delimiter string) (part1, part2 string)
```

ListAndTrim2 SplitAndTrim the `str` with `delimiter` and returns the result as two parts string.

​	ListAndTrim2 SplitAndTrim with `str` `delimiter` 并将结果作为两部分字符串返回。

#### func ListAndTrim3 <-2.5.5

```go
func ListAndTrim3(str, delimiter string) (part1, part2, part3 string)
```

ListAndTrim3 SplitAndTrim the `str` with `delimiter` and returns the result as three parts string.

​	ListAndTrim3 SplitAndTrim with `str` `delimiter` 并将结果作为三部分字符串返回。

#### func Nl2Br

```go
func Nl2Br(str string, isXhtml ...bool) string
```

Nl2Br inserts HTML line breaks(`br`|

​	Nl2Br 插入 HTML 换行符（ `br` |
) before all newlines in a string: \n\r, \r\n, \r, \n. It considers parameter `str` as unicode string.

​	） 在字符串中的所有换行符之前：\n\r、\r\n、\r、\n。它将参数 `str` 视为 unicode 字符串。

##### Example

``` go
```

#### func NumberFormat

```go
func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string
```

NumberFormat formats a number with grouped thousands. Parameter `decimals`: Sets the number of decimal points. Parameter `decPoint`: Sets the separator for the decimal point. Parameter `thousandsSep`: Sets the thousands’ separator. See http://php.net/manual/en/function.number-format.php.

​	NumberFormat 使用分组的千位设置数字的格式。参数 `decimals` ：设置小数点数。参数 `decPoint` ：设置小数点的分隔符。参数 `thousandsSep` ：设置千位分隔符。请参见 http://php.net/manual/en/function.number-format.php。

Example: NumberFormat(1234.56, 2, “.”, “”) -> 1234,56 NumberFormat(1234.56, 2, “,”, " “) -> 1 234,56

​	示例：NumberFormat（1234.56， 2， “.”， “”） -> 1234,56 NumberFormat（1234.56， 2， “，”， “ ”） -> 1 234,56

##### Example

``` go
```

#### func OctStr

```go
func OctStr(str string) string
```

OctStr converts string container octal string to its original string, for example, to Chinese string.

​	OctStr 将字符串容器八进制字符串转换为其原始字符串，例如，转换为中文字符串。

Example: OctStr("\346\200\241”) -> 怡

​	示例：OctStr（“\346\200\241”） -> 怡

##### Example

``` go
```

#### func Ord

```go
func Ord(char string) int
```

Ord converts the first byte of a string to a value between 0 and 255.

​	Ord 将字符串的第一个字节转换为介于 0 和 255 之间的值。

Example: Chr(“A”) -> 65

​	示例：Chr（“A”） -> 65

##### Example

``` go
```

#### func Parse

```go
func Parse(s string) (result map[string]interface{}, err error)
```

Parse parses the string into map[string]interface{}.

​	Parse 将字符串解析为 map[string]interface{}。

v1=m&v2=n -> map[v1:m v2:n] v[a]=m&v[b]=n -> map[v:map[a:m b:n]] v[a][a]=m&v[a][b]=n -> map[v:map[a:map[a:m b:n]]] v[]=m&v[]=n -> map[v:[m n]] v[a][]=m&v[a][]=n -> map[v:map[a:[m n]]] v[][]=m&v[][]=n -> map[v:[map[]]] // Currently does not support nested slice. v=m&v[a]=n -> error a .[[b=c -> map[a___[b:c]

​	v1=m&v2=n -> map[v1：m v2：n] v[a]=m&v[b]=n -> map[v：map[a：m b：n]] v[a][a]=m&v[a][b]=n -> map[v：map[a：map[a：m b：n]]] v[]=m&v[]=n -> map[v：[m n]] v[a][]=m&v[a][]=n -> map[v：map[a：[m n]]]v[][]=m&v[][]=n -> map[v：[map[]]] // 目前不支持嵌套切片。v=m&v[a]=n -> 错误 a .[[b=c -> 地图[a___[b：c]

##### Example

``` go
```

#### func Pos

```go
func Pos(haystack, needle string, startOffset ...int) int
```

Pos returns the position of the first occurrence of `needle` in `haystack` from `startOffset`, case-sensitively. It returns -1, if not found.

​	Pos 区分大小写地返回 `needle` in `haystack` from `startOffset` 第一次出现的位置。如果未找到，则返回 -1。

##### Example

``` go
```

#### func PosI

```go
func PosI(haystack, needle string, startOffset ...int) int
```

PosI returns the position of the first occurrence of `needle` in `haystack` from `startOffset`, case-insensitively. It returns -1, if not found.

​	PosI 不区分大小写地返回 `needle` in `haystack` from `startOffset` 的第一个出现位置。如果未找到，则返回 -1。

##### Example

``` go
```

#### func PosIRune

```go
func PosIRune(haystack, needle string, startOffset ...int) int
```

PosIRune acts like function PosI but considers `haystack` and `needle` as unicode string.

​	PosIRune 的行为类似于函数 PosI，但将 和 `needle` 视为 `haystack` unicode 字符串。

##### Example

``` go
```

#### func PosR

```go
func PosR(haystack, needle string, startOffset ...int) int
```

PosR returns the position of the last occurrence of `needle` in `haystack` from `startOffset`, case-sensitively. It returns -1, if not found.

​	PosR 区分大小写地返回 `needle` in `haystack` from `startOffset` 的最后出现位置。如果未找到，则返回 -1。

##### Example

``` go
```

#### func PosRI

```go
func PosRI(haystack, needle string, startOffset ...int) int
```

PosRI returns the position of the last occurrence of `needle` in `haystack` from `startOffset`, case-insensitively. It returns -1, if not found.

​	PosRI 不区分大小写地返回 `needle` in `haystack` from `startOffset` 的最后一个位置。如果未找到，则返回 -1。

##### Example

``` go
```

#### func PosRIRune

```go
func PosRIRune(haystack, needle string, startOffset ...int) int
```

PosRIRune acts like function PosRI but considers `haystack` and `needle` as unicode string.

​	PosRIRune 的行为类似于函数 PosRI，但将 和 `needle` 视为 `haystack` unicode 字符串。

##### Example

``` go
```

#### func PosRRune

```go
func PosRRune(haystack, needle string, startOffset ...int) int
```

PosRRune acts like function PosR but considers `haystack` and `needle` as unicode string.

​	PosRRune 的行为类似于函数 PosR，但将 和 `needle` 视为 `haystack` unicode 字符串。

##### Example

``` go
```

#### func PosRune

```go
func PosRune(haystack, needle string, startOffset ...int) int
```

PosRune acts like function Pos but considers `haystack` and `needle` as unicode string.

​	PosRune 的作用类似于函数 Pos，但将 和 `needle` 视为 `haystack` unicode 字符串。

##### Example

``` go
```

#### func PrefixArray

```go
func PrefixArray(array []string, prefix string)
```

PrefixArray adds `prefix` string for each item of `array`.

​	PrefixArray 为 的每个 `array` 项目添加 `prefix` 字符串。

Example: PrefixArray([“a”,“b”], “gf_”) -> [“gf_a”, “gf_b”]

​	示例：PrefixArray（[“a”，“b”]， “gf_”） -> [“gf_a”， “gf_b”]

##### Example

``` go
```

#### func QuoteMeta

```go
func QuoteMeta(str string, chars ...string) string
```

QuoteMeta returns a version of `str` with a backslash character (`\`). If custom chars `chars` not given, it uses default chars: .+*?[^](http://ngd.cn/goframe/v2/text/gstr/$)

​	QuoteMeta 返回带有反斜杠字符 （ `\` ） 的 `str` 版本。如果 `chars` 未给出自定义字符，则使用默认字符：.+*？^

##### Example

``` go
```

#### func Repeat

```go
func Repeat(input string, multiplier int) string
```

Repeat returns a new string consisting of multiplier copies of the string input.

​	Repeat 返回一个新字符串，该字符串由字符串输入的乘法器副本组成。

Example: Repeat(“a”, 3) -> “aaa”

​	示例：Repeat（“a”， 3） -> “aaa”

##### Example

``` go
```

#### func Replace

```go
func Replace(origin, search, replace string, count ...int) string
```

Replace returns a copy of the string `origin` in which string `search` replaced by `replace` case-sensitively.

​	Replace 返回字符串 `origin` 的副本，其中字符串 `search` 区分大小写地替换。 `replace`

##### Example

``` go
```

#### func ReplaceByArray

```go
func ReplaceByArray(origin string, array []string) string
```

ReplaceByArray returns a copy of `origin`, which is replaced by a slice in order, case-sensitively.

​	ReplaceByArray 返回 `origin` 的副本，该副本按大小写区分顺序替换为切片。

##### Example

``` go
```

#### func ReplaceByMap

```go
func ReplaceByMap(origin string, replaces map[string]string) string
```

ReplaceByMap returns a copy of `origin`, which is replaced by a map in unordered way, case-sensitively.

​	ReplaceByMap 返回 `origin` 的副本，该副本以无序方式替换为映射，区分大小写。

##### Example

``` go
```

#### func ReplaceI

```go
func ReplaceI(origin, search, replace string, count ...int) string
```

ReplaceI returns a copy of the string `origin` in which string `search` replaced by `replace` case-insensitively.

​	ReplaceI 返回字符串 `origin` 的副本，其中字符串 `search` 替换为 `replace` 不区分大小写的字符串。

##### Example

``` go
```

#### func ReplaceIByArray

```go
func ReplaceIByArray(origin string, array []string) string
```

ReplaceIByArray returns a copy of `origin`, which is replaced by a slice in order, case-insensitively.

​	ReplaceIByArray 返回 的 `origin` 副本，该副本按大小写不区分大小写。

##### Example

``` go
```

#### func ReplaceIByMap

```go
func ReplaceIByMap(origin string, replaces map[string]string) string
```

ReplaceIByMap returns a copy of `origin`, which is replaced by a map in unordered way, case-insensitively.

​	ReplaceIByMap 返回 的 `origin` 副本，该副本以无序方式替换为映射，不区分大小写。

##### Example

``` go
```

#### func Reverse

```go
func Reverse(str string) string
```

Reverse returns a string which is the reverse of `str`.

​	Reverse 返回一个字符串，该字符串与 `str` 相反。

Example: Reverse(“123456”) -> “654321”

​	示例：reverse（“123456”） -> “654321”

##### Example

``` go
```

#### func SearchArray

```go
func SearchArray(a []string, s string) int
```

SearchArray searches string `s` in string slice `a` case-sensitively, returns its index in `a`. If `s` is not found in `a`, it returns -1.

​	SearchArray 区分 `a` 大小写地搜索字符串切片中的字符串 `s` ，返回其在 中的 `a` 索引。如果 `s` 在 `a` 中找不到，则返回 -1。

##### Example

``` go
```

#### func Shuffle

```go
func Shuffle(str string) string
```

Shuffle randomly shuffles a string. It considers parameter `str` as unicode string.

​	随机随机洗牌字符串。它将参数 `str` 视为 unicode 字符串。

Example: Shuffle(“123456”) -> “325164” Shuffle(“123456”) -> “231546” …

​	示例： Shuffle（“123456”） -> “325164” Shuffle（“123456”） -> “231546” ...

##### Example

``` go
```

#### func SimilarText

```go
func SimilarText(first, second string, percent *float64) int
```

SimilarText calculates the similarity between two strings. See http://php.net/manual/en/function.similar-text.php.

​	SimilarText 计算两个字符串之间的相似度。请参见 http://php.net/manual/en/function.similar-text.php。

##### Example

``` go
```

#### func Soundex

```go
func Soundex(str string) string
```

Soundex calculates the soundex key of a string. See http://php.net/manual/en/function.soundex.php.

​	Soundex 计算字符串的 soundex 键。请参见 http://php.net/manual/en/function.soundex.php。

##### Example

``` go
```

#### func Split

```go
func Split(str, delimiter string) []string
```

Split splits string `str` by a string `delimiter`, to an array.

​	将字符串 `str` 拆分为一个字符串 `delimiter` ，以拆分为一个数组。

##### Example

``` go
```

#### func SplitAndTrim

```go
func SplitAndTrim(str, delimiter string, characterMask ...string) []string
```

SplitAndTrim splits string `str` by a string `delimiter` to an array, and calls Trim to every element of this array. It ignores the elements which are empty after Trim.

​	SplitAndTrim 将字符串 `str` 逐个字符串 `delimiter` 拆分为数组，并为此数组的每个元素调用 Trim。它忽略 Trim 之后为空的元素。

##### Example

``` go
```

#### func Str

```go
func Str(haystack string, needle string) string
```

Str returns part of `haystack` string starting from and including the first occurrence of `needle` to the end of `haystack`.

​	Str 返回从 开始（包括 第一次 `needle` 出现）到 结尾 `haystack` 的 `haystack` 字符串的一部分。

This function performs exactly as function SubStr, but to implement the same function as PHP: http://php.net/manual/en/function.strstr.php.

​	此函数的执行方式与函数 SubStr 完全相同，但要实现与 PHP 相同的函数：http://php.net/manual/en/function.strstr.php。

Example: Str(“av.mp4”, “.”) -> “.mp4”

​	示例： str（“av.mp4”， “.”） -> “.mp4”

##### Example

``` go
```

#### func StrEx

```go
func StrEx(haystack string, needle string) string
```

StrEx returns part of `haystack` string starting from and excluding the first occurrence of `needle` to the end of `haystack`.

​	StrEx 返回从 开始并排除 到 结尾 `haystack` 的第一次出现 `needle` 的 `haystack` 字符串的一部分。

This function performs exactly as function SubStrEx, but to implement the same function as PHP: http://php.net/manual/en/function.strstr.php.

​	此函数的执行方式与函数 SubStrEx 完全相同，但要实现与 PHP 相同的函数：http://php.net/manual/en/function.strstr.php。

Example: StrEx(“av.mp4”, “.”) -> “mp4”

​	示例： StrEx（“av.mp4”， “.”） -> “mp4”

##### Example

``` go
```

#### func StrLimit

```go
func StrLimit(str string, length int, suffix ...string) string
```

StrLimit returns a portion of string `str` specified by `length` parameters, if the length of `str` is greater than `length`, then the `suffix` will be appended to the result string.

​	StrLimit 返回由参数指定的字符串 `str` 的一部分，如果 的 `str` 长度大于 `length` ，则 `suffix` 将追加到结果字符串中。 `length`

Example: StrLimit(“123456”, 3) -> “123…” StrLimit(“123456”, 3, “~”) -> “123~”

​	示例：StrLimit（“123456”， 3） -> “123...”StrLimit（“123456”， 3， “~”） -> “123~”

##### Example

``` go
```

#### func StrLimitRune

```go
func StrLimitRune(str string, length int, suffix ...string) string
```

StrLimitRune returns a portion of string `str` specified by `length` parameters, if the length of `str` is greater than `length`, then the `suffix` will be appended to the result string. StrLimitRune considers parameter `str` as unicode string.

​	StrLimitRune 返回参数指定的字符串 `str` 的一部分，如果 的 `str` 长度大于 `length` ，则 `suffix` 将追加到结果字符串 `length` 中。StrLimitRune 将参数 `str` 视为 unicode 字符串。

Example: StrLimitRune(“一起学习吧！”, 2) -> “一起…” StrLimitRune(“一起学习吧！”, 2, “~”) -> “一起~”

##### Example

``` go
```

#### func StrTill

```go
func StrTill(haystack string, needle string) string
```

StrTill returns part of `haystack` string ending to and including the first occurrence of `needle` from the start of `haystack`.

​	StrTill 返回以 开头结尾的 `haystack` 字符串的一部分，包括从 开头 `haystack` 出现的 `needle` 字符串。

Example: StrTill(“av.mp4”, “.”) -> “av.”

​	示例：StrTill（“av.mp4”， “.”） -> “off”。

##### Example

``` go
```

#### func StrTillEx

```go
func StrTillEx(haystack string, needle string) string
```

StrTillEx returns part of `haystack` string ending to and excluding the first occurrence of `needle` from the start of `haystack`.

​	StrTillEx 返回以 结尾的 `haystack` 字符串的一部分，并排除从 开头 `haystack` 出现的 `needle` 字符串。

Example: StrTillEx(“av.mp4”, “.”) -> “av”

​	示例：StrTillEx（“av.mp4”， “.”） -> “off”

##### Example

``` go
```

#### func StripSlashes

```go
func StripSlashes(str string) string
```

StripSlashes un-quotes a quoted string by AddSlashes.

​	StripSlashes 取消引用 AddSlashes 的带引号的字符串。

##### Example

``` go
```

#### func SubStr

```go
func SubStr(str string, start int, length ...int) (substr string)
```

SubStr returns a portion of string `str` specified by the `start` and `length` parameters. The parameter `length` is optional, it uses the length of `str` in default.

​	SubStr 返回由 `start` and `length` 参数指定的字符串 `str` 的一部分。该参数 `length` 是可选的，它使用默认的长度 `str` 。

Example: SubStr(“123456”, 1, 2) -> “23”

​	示例：SubStr（“123456”， 1， 2） -> “23”

##### Example

``` go
```

#### func SubStrFrom

```go
func SubStrFrom(str string, need string) (substr string)
```

SubStrFrom returns a portion of string `str` starting from first occurrence of and including `need` to the end of `str`.

​	SubStrFrom 返回 `str` 从第一次出现的字符串的一部分，包括 `need` 到 的 `str` 末尾。

Example: SubStrFrom(“av.mp4”, “.”) -> “.mp4”

​	示例：SubStrFrom（“av.mp4”， “.”） -> “.mp4”

##### Example

``` go
```

#### func SubStrFromEx

```go
func SubStrFromEx(str string, need string) (substr string)
```

SubStrFromEx returns a portion of string `str` starting from first occurrence of and excluding `need` to the end of `str`.

​	SubStrFromEx 返回 `str` 从第一次出现的字符串开始并排除 `need` 到结尾 `str` 的部分字符串。

Example: SubStrFromEx(“av.mp4”, “.”) -> “mp4”

​	示例：SubStrFromEx（“av.mp4”， “.”） -> “mp4”

##### Example

``` go
```

#### func SubStrFromR

```go
func SubStrFromR(str string, need string) (substr string)
```

SubStrFromR returns a portion of string `str` starting from last occurrence of and including `need` to the end of `str`.

​	SubStrFromR 返回 `str` 从上次出现的字符串开始的一部分，包括 `need` 到 的 `str` 末尾。

Example: SubStrFromR("/dev/vda", “/”) -> “/vda”

​	示例：SubStrFromR（“/dev/vda”， “/”） -> “/vda”

##### Example

``` go
```

#### func SubStrFromREx

```go
func SubStrFromREx(str string, need string) (substr string)
```

SubStrFromREx returns a portion of string `str` starting from last occurrence of and excluding `need` to the end of `str`.

​	SubStrFromREx 返回字符串的一部分 `str` ，从上次出现的字符串开始， `need` 排除到的 `str` 末尾。

Example: SubStrFromREx("/dev/vda", “/”) -> “vda”

​	示例：SubStrFromREx（“/dev/vda”， “/”） -> “vda”

##### Example

``` go
```

#### func SubStrRune

```go
func SubStrRune(str string, start int, length ...int) (substr string)
```

SubStrRune returns a portion of string `str` specified by the `start` and `length` parameters. SubStrRune considers parameter `str` as unicode string. The parameter `length` is optional, it uses the length of `str` in default.

​	SubStrRune 返回由 `start` and `length` 参数指定的字符串 `str` 的一部分。SubStrRune 将参数 `str` 视为 unicode 字符串。该参数 `length` 是可选的，它使用默认的长度 `str` 。

Example: SubStrRune(“一起学习吧！”, 2, 2) -> “学习”

##### Example

``` go
```

#### func ToLower

```go
func ToLower(s string) string
```

ToLower returns a copy of the string s with all Unicode letters mapped to their lower case.

​	ToLower 返回字符串 s 的副本，其中所有 Unicode 字母都映射到其小写字母。

##### Example

``` go
```

#### func ToUpper

```go
func ToUpper(s string) string
```

ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.

​	ToUpper 返回字符串 s 的副本，其中所有 Unicode 字母都映射到其大写字母。

##### Example

``` go
```

#### func Trim

```go
func Trim(str string, characterMask ...string) string
```

Trim strips whitespace (or other characters) from the beginning and end of a string. The optional parameter `characterMask` specifies the additional stripped characters.

​	修剪从字符串的开头和结尾去除空格（或其他字符）。可选参数 `characterMask` 指定其他剥离字符。

##### Example

``` go
```

#### func TrimAll

```go
func TrimAll(str string, characterMask ...string) string
```

TrimAll trims all characters in string `str`.

​	TrimAll 修剪字符串 `str` 中的所有字符。

##### Example

``` go
```

#### func TrimLeft

```go
func TrimLeft(str string, characterMask ...string) string
```

TrimLeft strips whitespace (or other characters) from the beginning of a string.

​	TrimLeft 从字符串的开头去除空格（或其他字符）。

##### Example

``` go
```

#### func TrimLeftStr

```go
func TrimLeftStr(str string, cut string, count ...int) string
```

TrimLeftStr strips all the given `cut` string from the beginning of a string. Note that it does not strip the whitespaces of its beginning.

​	TrimLeftStr 从字符串的开头剥离所有给定 `cut` 的字符串。请注意，它不会去除其开头的空格。

##### Example

``` go
```

#### func TrimRight

```go
func TrimRight(str string, characterMask ...string) string
```

TrimRight strips whitespace (or other characters) from the end of a string.

​	TrimRight 从字符串末尾去除空格（或其他字符）。

##### Example

``` go
```

#### func TrimRightStr

```go
func TrimRightStr(str string, cut string, count ...int) string
```

TrimRightStr strips all the given `cut` string from the end of a string. Note that it does not strip the whitespaces of its end.

​	TrimRightStr 从字符串的末尾剥离所有给定 `cut` 的字符串。请注意，它不会去除其末端的空格。

##### Example

``` go
```

#### func TrimStr

```go
func TrimStr(str string, cut string, count ...int) string
```

TrimStr strips all the given `cut` string from the beginning and end of a string. Note that it does not strip the whitespaces of its beginning or end.

​	TrimStr 从字符串的开头和结尾剥离所有给定 `cut` 的字符串。请注意，它不会剥离其开头或结尾的空格。

##### Example

``` go
```

#### func UcFirst

```go
func UcFirst(s string) string
```

UcFirst returns a copy of the string s with the first letter mapped to its upper case.

​	UcFirst 返回字符串 s 的副本，其中第一个字母映射到其大写字母。

##### Example

``` go
```

#### func UcWords

```go
func UcWords(str string) string
```

UcWords uppercase the first character of each word in a string.

​	UcWords 将字符串中每个单词的第一个字符大写。

##### Example

``` go
```

#### func WordWrap

```go
func WordWrap(str string, width int, br string) string
```

WordWrap wraps a string to a given number of characters. This function supports cut parameters of both english and chinese punctuations. TODO: Enable custom cut parameter, see http://php.net/manual/en/function.wordwrap.php.

​	WordWrap 将字符串包装为给定数量的字符。此功能支持英文和中文标点符号的剪切参数。TODO：启用自定义切割参数，请参阅 http://php.net/manual/en/function.wordwrap.php。

##### Example

``` go
```

## 类型

### type CaseType <-2.5.7

```go
type CaseType string
```

CaseType is the type for Case.

​	CaseType 是 Case 的类型。

```go
const (
	Camel           CaseType = "Camel"
	CamelLower      CaseType = "CamelLower"
	Snake           CaseType = "Snake"
	SnakeFirstUpper CaseType = "SnakeFirstUpper"
	SnakeScreaming  CaseType = "SnakeScreaming"
	Kebab           CaseType = "Kebab"
	KebabScreaming  CaseType = "KebabScreaming"
	Lower           CaseType = "Lower"
)
```

The case type constants.

​	大小写类型常量。

#### func CaseTypeMatch <-2.5.7

```go
func CaseTypeMatch(caseStr string) CaseType
```

CaseTypeMatch matches the case type from string.

​	CaseTypeMatch 与字符串中的案例类型匹配。