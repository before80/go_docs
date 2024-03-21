+++
title = "gstr"
date = 2024-03-21T17:58:55+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/text/gstr

Package gstr provides functions for string handling.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/text/gstr/gstr.go#L10)

``` go
const (
	// NotFoundIndex is the position index for string not found in searching functions.
	NotFoundIndex = -1
)
```

### Variables 

This section is empty.

### Functions 

##### func AddSlashes 

``` go
func AddSlashes(str string) string
```

AddSlashes quotes with slashes `\` for chars: '"\.

##### Example

``` go
```
##### func CaseCamel 

``` go
func CaseCamel(s string) string
```

CaseCamel converts a string to CamelCase.

Example: CaseCamel("any_kind_of_string") -> AnyKindOfString

##### Example

``` go
```
##### func CaseCamelLower 

``` go
func CaseCamelLower(s string) string
```

CaseCamelLower converts a string to lowerCamelCase.

Example: CaseCamelLower("any_kind_of_string") -> anyKindOfString

##### Example

``` go
```
##### func CaseConvert <-2.5.7

``` go
func CaseConvert(s string, caseType CaseType) string
```

CaseConvert converts a string to the specified naming convention. Use CaseTypeMatch to match the case type from string.

##### func CaseDelimited 

``` go
func CaseDelimited(s string, del byte) string
```

CaseDelimited converts a string to snake.case.delimited.

Example: CaseDelimited("AnyKindOfString", '.') -> any.kind.of.string

##### Example

``` go
```
##### func CaseDelimitedScreaming 

``` go
func CaseDelimitedScreaming(s string, del uint8, screaming bool) string
```

CaseDelimitedScreaming converts a string to DELIMITED.SCREAMING.CASE or delimited.screaming.case.

Example: CaseDelimitedScreaming("AnyKindOfString", '.') -> ANY.KIND.OF.STRING

##### Example

``` go
```
##### func CaseKebab 

``` go
func CaseKebab(s string) string
```

CaseKebab converts a string to kebab-case.

Example: CaseKebab("AnyKindOfString") -> any-kind-of-string

##### Example

``` go
```
##### func CaseKebabScreaming 

``` go
func CaseKebabScreaming(s string) string
```

CaseKebabScreaming converts a string to KEBAB-CASE-SCREAMING.

Example: CaseKebab("AnyKindOfString") -> ANY-KIND-OF-STRING

##### Example

``` go
```
##### func CaseSnake 

``` go
func CaseSnake(s string) string
```

CaseSnake converts a string to snake_case.

Example: CaseSnake("AnyKindOfString") -> any_kind_of_string

##### Example

``` go
```
##### func CaseSnakeFirstUpper 

``` go
func CaseSnakeFirstUpper(word string, underscore ...string) string
```

CaseSnakeFirstUpper converts a string like "RGBCodeMd5" to "rgb_code_md5". TODO for efficiency should change regexp to traversing string in future.

Example: CaseSnakeFirstUpper("RGBCodeMd5") -> rgb_code_md5

##### Example

``` go
```
##### func CaseSnakeScreaming 

``` go
func CaseSnakeScreaming(s string) string
```

CaseSnakeScreaming converts a string to SNAKE_CASE_SCREAMING.

Example: CaseSnakeScreaming("AnyKindOfString") -> ANY_KIND_OF_STRING

##### Example

``` go
```
##### func Chr 

``` go
func Chr(ascii int) string
```

Chr return the ascii string of a number(0-255).

Example: Chr(65) -> "A"

##### Example

``` go
```
##### func ChunkSplit 

``` go
func ChunkSplit(body string, chunkLen int, end string) string
```

ChunkSplit splits a string into smaller chunks. Can be used to split a string into smaller chunks which is useful for e.g. converting BASE64 string output to match [RFC 2045](https://rfc-editor.org/rfc/rfc2045.html) semantics. It inserts end every chunkLen characters. It considers parameter `body` and `end` as unicode string.

##### Example

``` go
```
##### func Compare 

``` go
func Compare(a, b string) int
```

Compare returns an integer comparing two strings lexicographically. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.

##### Example

``` go
```
##### func CompareVersion 

``` go
func CompareVersion(a, b string) int
```

CompareVersion compares `a` and `b` as standard GNU version.

It returns 1 if `a` > `b`.

It returns -1 if `a` < `b`.

It returns 0 if `a` = `b`.

GNU standard version is like: v1.0 1 1.0.0 v1.0.1 v2.10.8 10.2.0 etc.

##### Example

``` go
```
##### func CompareVersionGo 

``` go
func CompareVersionGo(a, b string) int
```

CompareVersionGo compares `a` and `b` as standard Golang version.

It returns 1 if `a` > `b`.

It returns -1 if `a` < `b`.

It returns 0 if `a` = `b`.

Golang standard version is like: 1.0.0 v1.0.1 v2.10.8 10.2.0 v0.0.0-20190626092158-b2ccc519800e v1.12.2-0.20200413154443-b17e3a6804fa v4.20.0+incompatible etc.

Docs: https://go.dev/doc/modules/version-numbers

##### Example

``` go
```
##### func Contains 

``` go
func Contains(str, substr string) bool
```

Contains reports whether `substr` is within `str`, case-sensitively.

##### Example

``` go
```
##### func ContainsAny 

``` go
func ContainsAny(s, chars string) bool
```

ContainsAny reports whether any Unicode code points in `chars` are within `s`.

##### Example

``` go
```
##### func ContainsI 

``` go
func ContainsI(str, substr string) bool
```

ContainsI reports whether substr is within str, case-insensitively.

##### Example

``` go
```
##### func Count 

``` go
func Count(s, substr string) int
```

Count counts the number of `substr` appears in `s`. It returns 0 if no `substr` found in `s`.

##### Example

``` go
```
##### func CountChars 

``` go
func CountChars(str string, noSpace ...bool) map[string]int
```

CountChars returns information about chars' count used in a string. It considers parameter `str` as unicode string.

##### Example

``` go
```
##### func CountI 

``` go
func CountI(s, substr string) int
```

CountI counts the number of `substr` appears in `s`, case-insensitively. It returns 0 if no `substr` found in `s`.

##### Example

``` go
```
##### func CountWords 

``` go
func CountWords(str string) map[string]int
```

CountWords returns information about words' count used in a string. It considers parameter `str` as unicode string.

##### Example

``` go
```
##### func Equal 

``` go
func Equal(a, b string) bool
```

Equal reports whether `a` and `b`, interpreted as UTF-8 strings, are equal under Unicode case-folding, case-insensitively.

##### Example

``` go
```
##### func Explode 

``` go
func Explode(delimiter, str string) []string
```

Explode splits string `str` by a string `delimiter`, to an array. See http://php.net/manual/en/function.explode.php.

##### Example

``` go
```
##### func Fields 

``` go
func Fields(str string) []string
```

Fields returns the words used in a string as slice.

##### Example

``` go
```
##### func HasPrefix 

``` go
func HasPrefix(s, prefix string) bool
```

HasPrefix tests whether the string s begins with prefix.

##### Example

``` go
```
##### func HasSuffix 

``` go
func HasSuffix(s, suffix string) bool
```

HasSuffix tests whether the string s ends with suffix.

##### Example

``` go
```
##### func HideStr 

``` go
func HideStr(str string, percent int, hide string) string
```

HideStr replaces part of the string `str` to `hide` by `percentage` from the `middle`. It considers parameter `str` as unicode string.

##### Example

``` go
```
##### func Implode 

``` go
func Implode(glue string, pieces []string) string
```

Implode joins array elements `pieces` with a string `glue`. http://php.net/manual/en/function.implode.php

##### Example

``` go
```
##### func InArray 

``` go
func InArray(a []string, s string) bool
```

InArray checks whether string `s` in slice `a`.

##### Example

``` go
```
##### func IsGNUVersion <-2.1.2

``` go
func IsGNUVersion(version string) bool
```

IsGNUVersion checks and returns whether given `version` is valid GNU version string.

##### func IsLetterLower 

``` go
func IsLetterLower(b byte) bool
```

IsLetterLower tests whether the given byte b is in lower case.

##### Example

``` go
```
##### func IsLetterUpper 

``` go
func IsLetterUpper(b byte) bool
```

IsLetterUpper tests whether the given byte b is in upper case.

##### Example

``` go
```
##### func IsNumeric 

``` go
func IsNumeric(s string) bool
```

IsNumeric tests whether the given string s is numeric.

##### Example

``` go
```
##### func IsSubDomain 

``` go
func IsSubDomain(subDomain string, mainDomain string) bool
```

IsSubDomain checks whether `subDomain` is sub-domain of mainDomain. It supports '*' in `mainDomain`.

##### Example

``` go
```
##### func Join 

``` go
func Join(array []string, sep string) string
```

Join concatenates the elements of `array` to create a single string. The separator string `sep` is placed between elements in the resulting string.

##### Example

``` go
```
##### func JoinAny 

``` go
func JoinAny(array interface{}, sep string) string
```

JoinAny concatenates the elements of `array` to create a single string. The separator string `sep` is placed between elements in the resulting string.

The parameter `array` can be any type of slice, which be converted to string array.

##### Example

``` go
```
##### func LcFirst 

``` go
func LcFirst(s string) string
```

LcFirst returns a copy of the string s with the first letter mapped to its lower case.

##### Example

``` go
```
##### func LenRune 

``` go
func LenRune(str string) int
```

LenRune returns string length of unicode.

##### Example

``` go
```
##### func Levenshtein 

``` go
func Levenshtein(str1, str2 string, costIns, costRep, costDel int) int
```

Levenshtein calculates Levenshtein distance between two strings. costIns: Defines the cost of insertion. costRep: Defines the cost of replacement. costDel: Defines the cost of deletion. See http://php.net/manual/en/function.levenshtein.php.

##### Example

``` go
```
##### func List2 <-2.5.5

``` go
func List2(str, delimiter string) (part1, part2 string)
```

List2 Split the `str` with `delimiter` and returns the result as two parts string.

##### func List3 <-2.5.5

``` go
func List3(str, delimiter string) (part1, part2, part3 string)
```

List3 Split the `str` with `delimiter` and returns the result as three parts string.

##### func ListAndTrim2 <-2.5.5

``` go
func ListAndTrim2(str, delimiter string) (part1, part2 string)
```

ListAndTrim2 SplitAndTrim the `str` with `delimiter` and returns the result as two parts string.

##### func ListAndTrim3 <-2.5.5

``` go
func ListAndTrim3(str, delimiter string) (part1, part2, part3 string)
```

ListAndTrim3 SplitAndTrim the `str` with `delimiter` and returns the result as three parts string.

##### func Nl2Br 

``` go
func Nl2Br(str string, isXhtml ...bool) string
```

Nl2Br inserts HTML line breaks(`br`|<br />) before all newlines in a string: \n\r, \r\n, \r, \n. It considers parameter `str` as unicode string.

##### Example

``` go
```
##### func NumberFormat 

``` go
func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string
```

NumberFormat formats a number with grouped thousands. Parameter `decimals`: Sets the number of decimal points. Parameter `decPoint`: Sets the separator for the decimal point. Parameter `thousandsSep`: Sets the thousands' separator. See http://php.net/manual/en/function.number-format.php.

Example: NumberFormat(1234.56, 2, ".", "") -> 1234,56 NumberFormat(1234.56, 2, ",", " ") -> 1 234,56

##### Example

``` go
```
##### func OctStr 

``` go
func OctStr(str string) string
```

OctStr converts string container octal string to its original string, for example, to Chinese string.

Example: OctStr("\346\200\241") -> 怡

##### Example

``` go
```
##### func Ord 

``` go
func Ord(char string) int
```

Ord converts the first byte of a string to a value between 0 and 255.

Example: Chr("A") -> 65

##### Example

``` go
```
##### func Parse 

``` go
func Parse(s string) (result map[string]interface{}, err error)
```

Parse parses the string into map[string]interface{}.

v1=m&v2=n -> map[v1:m v2:n] v[a]=m&v[b]=n -> map[v:map[a:m b:n]] v[a][a]=m&v[a][b]=n -> map[v:map[a:map[a:m b:n]]] v[]=m&v[]=n -> map[v:[m n]] v[a][]=m&v[a][]=n -> map[v:map[a:[m n]]] v[][]=m&v[][]=n -> map[v:[map[]]] // Currently does not support nested slice. v=m&v[a]=n -> error a .[[b=c -> map[a___[b:c]

##### Example

``` go
```
##### func Pos 

``` go
func Pos(haystack, needle string, startOffset ...int) int
```

Pos returns the position of the first occurrence of `needle` in `haystack` from `startOffset`, case-sensitively. It returns -1, if not found.

##### Example

``` go
```
##### func PosI 

``` go
func PosI(haystack, needle string, startOffset ...int) int
```

PosI returns the position of the first occurrence of `needle` in `haystack` from `startOffset`, case-insensitively. It returns -1, if not found.

##### Example

``` go
```
##### func PosIRune 

``` go
func PosIRune(haystack, needle string, startOffset ...int) int
```

PosIRune acts like function PosI but considers `haystack` and `needle` as unicode string.

##### Example

``` go
```
##### func PosR 

``` go
func PosR(haystack, needle string, startOffset ...int) int
```

PosR returns the position of the last occurrence of `needle` in `haystack` from `startOffset`, case-sensitively. It returns -1, if not found.

##### Example

``` go
```
##### func PosRI 

``` go
func PosRI(haystack, needle string, startOffset ...int) int
```

PosRI returns the position of the last occurrence of `needle` in `haystack` from `startOffset`, case-insensitively. It returns -1, if not found.

##### Example

``` go
```
##### func PosRIRune 

``` go
func PosRIRune(haystack, needle string, startOffset ...int) int
```

PosRIRune acts like function PosRI but considers `haystack` and `needle` as unicode string.

##### Example

``` go
```
##### func PosRRune 

``` go
func PosRRune(haystack, needle string, startOffset ...int) int
```

PosRRune acts like function PosR but considers `haystack` and `needle` as unicode string.

##### Example

``` go
```
##### func PosRune 

``` go
func PosRune(haystack, needle string, startOffset ...int) int
```

PosRune acts like function Pos but considers `haystack` and `needle` as unicode string.

##### Example

``` go
```
##### func PrefixArray 

``` go
func PrefixArray(array []string, prefix string)
```

PrefixArray adds `prefix` string for each item of `array`.

Example: PrefixArray(["a","b"], "gf_") -> ["gf_a", "gf_b"]

##### Example

``` go
```
##### func QuoteMeta 

``` go
func QuoteMeta(str string, chars ...string) string
```

QuoteMeta returns a version of `str` with a backslash character (`\`). If custom chars `chars` not given, it uses default chars: .\+*?[^]($)

##### Example

``` go
```
##### func Repeat 

``` go
func Repeat(input string, multiplier int) string
```

Repeat returns a new string consisting of multiplier copies of the string input.

Example: Repeat("a", 3) -> "aaa"

##### Example

``` go
```
##### func Replace 

``` go
func Replace(origin, search, replace string, count ...int) string
```

Replace returns a copy of the string `origin` in which string `search` replaced by `replace` case-sensitively.

##### Example

``` go
```
##### func ReplaceByArray 

``` go
func ReplaceByArray(origin string, array []string) string
```

ReplaceByArray returns a copy of `origin`, which is replaced by a slice in order, case-sensitively.

##### Example

``` go
```
##### func ReplaceByMap 

``` go
func ReplaceByMap(origin string, replaces map[string]string) string
```

ReplaceByMap returns a copy of `origin`, which is replaced by a map in unordered way, case-sensitively.

##### Example

``` go
```
##### func ReplaceI 

``` go
func ReplaceI(origin, search, replace string, count ...int) string
```

ReplaceI returns a copy of the string `origin` in which string `search` replaced by `replace` case-insensitively.

##### Example

``` go
```
##### func ReplaceIByArray 

``` go
func ReplaceIByArray(origin string, array []string) string
```

ReplaceIByArray returns a copy of `origin`, which is replaced by a slice in order, case-insensitively.

##### Example

``` go
```
##### func ReplaceIByMap 

``` go
func ReplaceIByMap(origin string, replaces map[string]string) string
```

ReplaceIByMap returns a copy of `origin`, which is replaced by a map in unordered way, case-insensitively.

##### Example

``` go
```
##### func Reverse 

``` go
func Reverse(str string) string
```

Reverse returns a string which is the reverse of `str`.

Example: Reverse("123456") -> "654321"

##### Example

``` go
```
##### func SearchArray 

``` go
func SearchArray(a []string, s string) int
```

SearchArray searches string `s` in string slice `a` case-sensitively, returns its index in `a`. If `s` is not found in `a`, it returns -1.

##### Example

``` go
```
##### func Shuffle 

``` go
func Shuffle(str string) string
```

Shuffle randomly shuffles a string. It considers parameter `str` as unicode string.

Example: Shuffle("123456") -> "325164" Shuffle("123456") -> "231546" ...

##### Example

``` go
```
##### func SimilarText 

``` go
func SimilarText(first, second string, percent *float64) int
```

SimilarText calculates the similarity between two strings. See http://php.net/manual/en/function.similar-text.php.

##### Example

``` go
```
##### func Soundex 

``` go
func Soundex(str string) string
```

Soundex calculates the soundex key of a string. See http://php.net/manual/en/function.soundex.php.

##### Example

``` go
```
##### func Split 

``` go
func Split(str, delimiter string) []string
```

Split splits string `str` by a string `delimiter`, to an array.

##### Example

``` go
```
##### func SplitAndTrim 

``` go
func SplitAndTrim(str, delimiter string, characterMask ...string) []string
```

SplitAndTrim splits string `str` by a string `delimiter` to an array, and calls Trim to every element of this array. It ignores the elements which are empty after Trim.

##### Example

``` go
```
##### func Str 

``` go
func Str(haystack string, needle string) string
```

Str returns part of `haystack` string starting from and including the first occurrence of `needle` to the end of `haystack`.

This function performs exactly as function SubStr, but to implement the same function as PHP: http://php.net/manual/en/function.strstr.php.

Example: Str("av.mp4", ".") -> ".mp4"

##### Example

``` go
```
##### func StrEx 

``` go
func StrEx(haystack string, needle string) string
```

StrEx returns part of `haystack` string starting from and excluding the first occurrence of `needle` to the end of `haystack`.

This function performs exactly as function SubStrEx, but to implement the same function as PHP: http://php.net/manual/en/function.strstr.php.

Example: StrEx("av.mp4", ".") -> "mp4"

##### Example

``` go
```
##### func StrLimit 

``` go
func StrLimit(str string, length int, suffix ...string) string
```

StrLimit returns a portion of string `str` specified by `length` parameters, if the length of `str` is greater than `length`, then the `suffix` will be appended to the result string.

Example: StrLimit("123456", 3) -> "123..." StrLimit("123456", 3, "~") -> "123~"

##### Example

``` go
```
##### func StrLimitRune 

``` go
func StrLimitRune(str string, length int, suffix ...string) string
```

StrLimitRune returns a portion of string `str` specified by `length` parameters, if the length of `str` is greater than `length`, then the `suffix` will be appended to the result string. StrLimitRune considers parameter `str` as unicode string.

Example: StrLimitRune("一起学习吧！", 2) -> "一起..." StrLimitRune("一起学习吧！", 2, "~") -> "一起~"

##### Example

``` go
```
##### func StrTill 

``` go
func StrTill(haystack string, needle string) string
```

StrTill returns part of `haystack` string ending to and including the first occurrence of `needle` from the start of `haystack`.

Example: StrTill("av.mp4", ".") -> "av."

##### Example

``` go
```
##### func StrTillEx 

``` go
func StrTillEx(haystack string, needle string) string
```

StrTillEx returns part of `haystack` string ending to and excluding the first occurrence of `needle` from the start of `haystack`.

Example: StrTillEx("av.mp4", ".") -> "av"

##### Example

``` go
```
##### func StripSlashes 

``` go
func StripSlashes(str string) string
```

StripSlashes un-quotes a quoted string by AddSlashes.

##### Example

``` go
```
##### func SubStr 

``` go
func SubStr(str string, start int, length ...int) (substr string)
```

SubStr returns a portion of string `str` specified by the `start` and `length` parameters. The parameter `length` is optional, it uses the length of `str` in default.

Example: SubStr("123456", 1, 2) -> "23"

##### Example

``` go
```
##### func SubStrFrom 

``` go
func SubStrFrom(str string, need string) (substr string)
```

SubStrFrom returns a portion of string `str` starting from first occurrence of and including `need` to the end of `str`.

Example: SubStrFrom("av.mp4", ".") -> ".mp4"

##### Example

``` go
```
##### func SubStrFromEx 

``` go
func SubStrFromEx(str string, need string) (substr string)
```

SubStrFromEx returns a portion of string `str` starting from first occurrence of and excluding `need` to the end of `str`.

Example: SubStrFromEx("av.mp4", ".") -> "mp4"

##### Example

``` go
```
##### func SubStrFromR 

``` go
func SubStrFromR(str string, need string) (substr string)
```

SubStrFromR returns a portion of string `str` starting from last occurrence of and including `need` to the end of `str`.

Example: SubStrFromR("/dev/vda", "/") -> "/vda"

##### Example

``` go
```
##### func SubStrFromREx 

``` go
func SubStrFromREx(str string, need string) (substr string)
```

SubStrFromREx returns a portion of string `str` starting from last occurrence of and excluding `need` to the end of `str`.

Example: SubStrFromREx("/dev/vda", "/") -> "vda"

##### Example

``` go
```
##### func SubStrRune 

``` go
func SubStrRune(str string, start int, length ...int) (substr string)
```

SubStrRune returns a portion of string `str` specified by the `start` and `length` parameters. SubStrRune considers parameter `str` as unicode string. The parameter `length` is optional, it uses the length of `str` in default.

Example: SubStrRune("一起学习吧！", 2, 2) -> "学习"

##### Example

``` go
```
##### func ToLower 

``` go
func ToLower(s string) string
```

ToLower returns a copy of the string s with all Unicode letters mapped to their lower case.

##### Example

``` go
```
##### func ToUpper 

``` go
func ToUpper(s string) string
```

ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.

##### Example

``` go
```
##### func Trim 

``` go
func Trim(str string, characterMask ...string) string
```

Trim strips whitespace (or other characters) from the beginning and end of a string. The optional parameter `characterMask` specifies the additional stripped characters.

##### Example

``` go
```
##### func TrimAll 

``` go
func TrimAll(str string, characterMask ...string) string
```

TrimAll trims all characters in string `str`.

##### Example

``` go
```
##### func TrimLeft 

``` go
func TrimLeft(str string, characterMask ...string) string
```

TrimLeft strips whitespace (or other characters) from the beginning of a string.

##### Example

``` go
```
##### func TrimLeftStr 

``` go
func TrimLeftStr(str string, cut string, count ...int) string
```

TrimLeftStr strips all the given `cut` string from the beginning of a string. Note that it does not strip the whitespaces of its beginning.

##### Example

``` go
```
##### func TrimRight 

``` go
func TrimRight(str string, characterMask ...string) string
```

TrimRight strips whitespace (or other characters) from the end of a string.

##### Example

``` go
```
##### func TrimRightStr 

``` go
func TrimRightStr(str string, cut string, count ...int) string
```

TrimRightStr strips all the given `cut` string from the end of a string. Note that it does not strip the whitespaces of its end.

##### Example

``` go
```
##### func TrimStr 

``` go
func TrimStr(str string, cut string, count ...int) string
```

TrimStr strips all the given `cut` string from the beginning and end of a string. Note that it does not strip the whitespaces of its beginning or end.

##### Example

``` go
```
##### func UcFirst 

``` go
func UcFirst(s string) string
```

UcFirst returns a copy of the string s with the first letter mapped to its upper case.

##### Example

``` go
```
##### func UcWords 

``` go
func UcWords(str string) string
```

UcWords uppercase the first character of each word in a string.

##### Example

``` go
```
##### func WordWrap 

``` go
func WordWrap(str string, width int, br string) string
```

WordWrap wraps a string to a given number of characters. This function supports cut parameters of both english and chinese punctuations. TODO: Enable custom cut parameter, see http://php.net/manual/en/function.wordwrap.php.

##### Example

``` go
```
### Types 

#### type CaseType <-2.5.7

``` go
type CaseType string
```

CaseType is the type for Case.

``` go
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

##### func CaseTypeMatch <-2.5.7

``` go
func CaseTypeMatch(caseStr string) CaseType
```

CaseTypeMatch matches the case type from string.