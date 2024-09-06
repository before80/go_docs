+++
title = "strings"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
> 原文：[https://pkg.go.dev/strings@go1.23.0](https://pkg.go.dev/strings@go1.23.0)

Package strings implements simple functions to manipulate UTF-8 encoded strings.

​	`strings`包实现了一些简单的函数来操作 UTF-8 编码的字符串。

For information about UTF-8 strings in Go, see https://blog.golang.org/strings.

​	有关 Go 中的 UTF-8 字符串的信息，请参阅 [https://blog.golang.org/strings](https://blog.golang.org/strings)。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Clone  <- go1.18

``` go 
func Clone(s string) string
```

Clone returns a fresh copy of s. It guarantees to make a copy of s into a new allocation, which can be important when retaining only a small substring of a much larger string. Using Clone can help such programs use less memory. Of course, since using Clone makes a copy, overuse of Clone can make programs use more memory. Clone should typically be used only rarely, and only when profiling indicates that it is needed. For strings of length zero the string "" will be returned and no allocation is made.

​	Clone函数返回 s 的一个全新副本。它保证将 s 复制到一个新的分配中，当仅保留一个更大字符串的小子串时，这可能很重要。使用 Clone函数可以帮助这些程序使用更少的内存。当然，由于使用 Clone函数会进行复制，过度使用 Clone函数可能会使程序使用更多的内存。通常应仅在分析表明需要时才使用 Clone函数。对于长度为零的字符串，将返回字符串 ""，并且不会进行任何分配。

### func Compare  <- go1.5

``` go 
func Compare(a, b string) int
```

Compare returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b.

​	Compare函数按字典顺序比较两个字符串。

如果 a == b，则结果为 0；

如果 a < b，则结果为 -1；

如果 a > b，则结果为 +1。

Compare is included only for symmetry with package bytes. It is usually clearer and always faster to use the built-in string comparison operators ==, <, >, and so on.

​	Compare 函数只是为了与 bytes 包保持对称性而被包含进来的。使用内置的字符串比较运算符 ==、<、> 等通常更加清晰明了，而且速度更快。

#### Compare Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
}
Output:

-1
0
1
```

### func Contains 

``` go 
func Contains(s, substr string) bool
```

Contains reports whether substr is within s.

​	Contains函数报告 substr 是否在 s 中。

#### Contains Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
}
Output:

true
false
true
true
```

### func ContainsAny 

``` go 
func ContainsAny(s, chars string) bool
```

ContainsAny reports whether any Unicode code points in chars are within s.

​	ContainsAny函数报告 s 中是否包含 chars 中的任何 Unicode 码点。

#### ContainsAny Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("fail", "ui"))
	fmt.Println(strings.ContainsAny("ure", "ui"))
	fmt.Println(strings.ContainsAny("failure", "ui"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))
}
Output:

false
true
true
true
false
false
```

### func ContainsFunc <-go1.21.0

```go
func ContainsFunc(s string, f func(rune) bool) bool
```

ContainsFunc reports whether any Unicode code points r within s satisfy f(r).

​	ContainsFunc 判断字符串 s 中是否有任意 Unicode 码点 r 满足函数 f(r) 的条件。

### func ContainsRune 

``` go 
func ContainsRune(s string, r rune) bool
```

ContainsRune reports whether the Unicode code point r is within s.

​	ContainsRune函数报告 Unicode 码点 r 是否在 s 中。

#### ContainsRune Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 查找字符串是否包含特定的 Unicode 码点。
	// 例如，小写字母"a"的码点为97。
	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
}
Output:

true
false
```

### func Count 

``` go 
func Count(s, substr string) int
```

Count counts the number of non-overlapping instances of substr in s. If substr is an empty string, Count returns 1 + the number of Unicode code points in s.

​	Count函数返回 s 中不重叠的 substr 实例数。如果 substr 是空字符串，则 Count 返回 s 中 Unicode 码点的数量加 1。

#### Count Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
}
Output:

3
5
```

### func Cut  <- go1.18

``` go 
func Cut(s, sep string) (before, after string, found bool)
```

Cut slices s around the first instance of sep, returning the text before and after sep. The found result reports whether sep appears in s. If sep does not appear in s, cut returns s, "", false.

​	Cut函数围绕第一个 sep 实例对 s 进行切片，返回 sep 之前和之后的文本。found 结果报告 sep 是否出现在 s 中。如果 sep 不出现在 s 中，则 Cut 返回 s、""、false。

#### Cut Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	show := func(s, sep string) {
		before, after, found := strings.Cut(s, sep)
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
	show("Gopher", "er")
	show("Gopher", "Badger")
}
Output:

Cut("Gopher", "Go") = "", "pher", true
Cut("Gopher", "ph") = "Go", "er", true
Cut("Gopher", "er") = "Goph", "", true
Cut("Gopher", "Badger") = "Gopher", "", false
```

### func CutPrefix  <- go1.20

``` go 
func CutPrefix(s, prefix string) (after string, found bool)
```

CutPrefix returns s without the provided leading prefix string and reports whether it found the prefix. If s doesn't start with prefix, CutPrefix returns s, false. If prefix is the empty string, CutPrefix returns s, true.

​	CutPrefix函数返回去掉前缀字符串 prefix 后的 s，并报告它是否找到了前缀。如果 s 不以 prefix 开头，则 CutPrefix函数返回 s、false。如果 prefix 是空字符串，则 CutPrefix 返回 s、true。

### func CutSuffix  <- go1.20

``` go 
func CutSuffix(s, suffix string) (before string, found bool)
```

CutSuffix returns s without the provided ending suffix string and reports whether it found the suffix. If s doesn't end with suffix, CutSuffix returns s, false. If suffix is the empty string, CutSuffix returns s, true.

​	CutSuffix函数返回s中不包含指定结尾后缀字符串的部分，并报告它是否找到了后缀。如果s不以后缀结尾，则CutSuffix函数返回s，false。如果后缀为空字符串，则CutSuffix函数返回s，true。

### func EqualFold 

``` go 
func EqualFold(s, t string) bool
```

EqualFold reports whether s and t, interpreted as UTF-8 strings, are equal under simple Unicode case-folding, which is a more general form of case-insensitivity.

​	EqualFold函数报告s和t在简单的Unicode折叠(一种更一般的不区分大小写的形式)下是否相等，这是一种更一般的不区分大小写的形式。

#### EqualFold Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.EqualFold("Go", "go"))
	fmt.Println(strings.EqualFold("AB", "ab")) // true because comparison uses simple case-folding
	fmt.Println(strings.EqualFold("ß", "ss"))  // false because comparison does not use full case-folding
}
Output:

true
true
false
```

### func Fields 

``` go 
func Fields(s string) []string
```

Fields splits the string s around each instance of one or more consecutive white space characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.

​	Fields函数将字符串s按一个或多个连续的空格字符(由unicode.IsSpace定义)分割，返回s的子字符串切片或一个空切片，如果s仅包含空格，则返回空切片。

#### Fields Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}
Output:

Fields are: ["foo" "bar" "baz"]
```

### func FieldsFunc 

``` go 
func FieldsFunc(s string, f func(rune) bool) []string
```

FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c) and returns an array of slices of s. If all code points in s satisfy f(c) or the string is empty, an empty slice is returned.

​	FieldsFunc函数将字符串s在每个运行满足f(c) 的Unicode码点c处分割，并返回s的切片数组。如果s中所有码点都满足f(c) 或字符串为空，则返回空切片。

FieldsFunc makes no guarantees about the order in which it calls f(c) and assumes that f always returns the same value for a given c.

​	FieldsFunc函数不保证调用f(c) 的顺序，并假定f始终为给定c返回相同的值。

#### FieldsFunc Example
``` go 
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
}
Output:

Fields are: ["foo1" "bar2" "baz3"]
```

### func HasPrefix 

``` go 
func HasPrefix(s, prefix string) bool
```

HasPrefix tests whether the string s begins with prefix.

​	HasPrefix函数测试字符串s是否以prefix开头。

#### HasPrefix Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("Amigo", "go"))
	fmt.Println(strings.HasSuffix("Amigo", "O"))
	fmt.Println(strings.HasSuffix("Amigo", "Ami"))
	fmt.Println(strings.HasSuffix("Amigo", ""))
}
Output:

true
false
false
true
```

### func HasSuffix 

``` go 
func HasSuffix(s, suffix string) bool
```

HasSuffix tests whether the string s ends with suffix.

​	HasSuffix函数测试字符串s是否以后缀suffix结尾。

#### HasSuffix Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
}

Output:

true
false
true
```

### func Index 

``` go 
func Index(s, substr string) int
```

Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.

​	Index函数返回substr在s中第一次出现的索引，如果substr不在s中，则返回-1。

#### Index Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
}
Output:

4
-1
```

### func IndexAny 

``` go 
func IndexAny(s, chars string) int
```

IndexAny returns the index of the first instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.

​	IndexAny函数返回s中任何Unicode码点中chars的第一个实例的索引，如果s中不存在来自chars的Unicode码点，则返回-1。

#### IndexAny Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
}
Output:

2
-1
```

### func IndexByte  <- go1.2

``` go 
func IndexByte(s string, c byte) int
```

IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.

​	IndexByte函数返回c在s中第一次出现的索引，如果c不在s中，则返回-1。

#### IndexByte Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("gophers", 'h'))
	fmt.Println(strings.IndexByte("golang", 'x'))
}
Output:

0
3
-1
```

### func IndexFunc 

``` go 
func IndexFunc(s string, f func(rune) bool) int
```

IndexFunc returns the index into s of the first Unicode code point satisfying f(c), or -1 if none do.

​	IndexFunc函数返回第一个满足f(c) 的Unicode码点c在s中的索引，如果没有则返回-1。

#### IndexFunc Example
``` go 
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
}
Output:

7
-1
```

### func IndexRune 

``` go 
func IndexRune(s string, r rune) int
```

IndexRune returns the index of the first instance of the Unicode code point r, or -1 if rune is not present in s. If r is utf8.RuneError, it returns the first instance of any invalid UTF-8 byte sequence.

​	IndexRune函数返回Unicode码点r在字符串s中第一次出现的索引，如果r不在s中，则返回-1。如果r是utf8.RuneError，则返回任何无效的UTF-8字节序列的第一个实例。

#### IndexRune Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexRune("chicken", 'k'))
	fmt.Println(strings.IndexRune("chicken", 'd'))
}
Output:

4
-1
```

### func Join 

``` go 
func Join(elems []string, sep string) string
```

Join concatenates the elements of its first argument to create a single string. The separator string sep is placed between elements in the resulting string.

​	Join函数将其第一个参数的元素连接起来以创建单个字符串。分隔符字符串sep放置在结果字符串中的元素之间。

#### Join Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
}
Output:

foo, bar, baz
```

### func LastIndex 

``` go 
func LastIndex(s, substr string) int
```

LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.

​	LastIndexByte函数返回c在s中最后一次出现的索引，如果c不在s中，则返回-1。

#### LastIndex Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))
}
Output:

0
3
-1
```

### func LastIndexAny 

``` go 
func LastIndexAny(s, chars string) int
```

LastIndexAny returns the index of the last instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.

​	LastIndexAny函数返回 chars 中任何 Unicode 码点在 s 中最后一次出现的索引，如果 s 中没有来自 chars 的 Unicode 码点，则返回 -1。

#### LastIndexAny Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.LastIndexAny("go gopher", "go"))
	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))
}
Output:

4
8
-1
```

### func LastIndexByte  <- go1.5

``` go 
func LastIndexByte(s string, c byte) int
```

LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.

​	LastIndexByte函数返回 c 在 s 中最后一次出现的索引，如果 c 不在 s 中，则返回 -1。

#### LastIndexByte Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.LastIndexByte("Hello, world", 'l'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'o'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'x'))
}
Output:

10
8
-1
```

### func LastIndexFunc 

``` go 
func LastIndexFunc(s string, f func(rune) bool) int
```

LastIndexFunc returns the index into s of the last Unicode code point satisfying f(c), or -1 if none do.

​	LastIndexFunc函数返回最后一个满足 f(c) 的 Unicode 码点在 s 中的索引，如果没有，则返回-1。

#### LastIndexFunc Example
``` go 
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))
}
Output:

5
2
-1
```

### func Map 

``` go 
func Map(mapping func(rune) rune, s string) string
```

Map returns a copy of the string s with all its characters modified according to the mapping function. If mapping returns a negative value, the character is dropped from the string with no replacement.

​	Map函数返回一个新的字符串，其中包含根据映射函数修改的 s 中的所有字符。如果 mapping 返回**负值**，则**删除该字符**而不替换。（例如返回`-1`）

#### Map Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
}
Output:

'Gjnf oevyyvt naq gur fyvgul tbcure...
```

### func Repeat 

``` go 
func Repeat(s string, count int) string
```

Repeat returns a new string consisting of count copies of the string s.

​	Repeat函数返回由 count 个字符串 s 组成的新字符串。

It panics if count is negative or if the result of (len(s) * count) overflows.

​	如果 count 为负数或 `(len(s) * count)` 的结果溢出，则会发生 panic。

#### Repeat Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("ba" + strings.Repeat("na", 2))
}
Output:

banana
```

### func Replace 

``` go 
func Replace(s, old, new string, n int) string
```

Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string. If n < 0, there is no limit on the number of replacements.

​	Replace函数将字符串s中前`n`个非重叠old实例替换为new，并返回新的字符串。如果old为空，则它匹配字符串开头和每个UTF-8序列之后，为k个rune字符串提供最多`k+1`个替换。如果n < 0，则不限制替换次数。

#### Replace Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}
Output:

oinky oinky oink
moo moo moo
```

### func ReplaceAll  <- go1.12

``` go 
func ReplaceAll(s, old, new string) string
```

ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.

​	ReplaceAll函数将字符串s中所有非重叠old实例替换为new，并返回新的字符串。如果old为空，则它匹配字符串开头和每个UTF-8序列之后，为k个rune字符串提供最多k+1个替换。

#### ReplaceAll Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
}
Output:

moo moo moo
```

### func Split 

``` go 
func Split(s, sep string) []string
```

Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.

​	Split函数将字符串s按sep分割成所有子字符串，并返回它们之间的子字符串切片。

If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.

​	如果s不包含sep且sep不为空，则Split返回长度为1的切片，其中唯一的元素为s。

If sep is empty, Split splits after each UTF-8 sequence. If both s and sep are empty, Split returns an empty slice.

​	如果sep为空，则Split在每个UTF-8序列之后拆分。如果s和sep都为空，则Split返回一个空切片。

It is equivalent to SplitN with a count of -1.

​	它等价于带有计数-1的SplitN函数。

To split around the first instance of a separator, see Cut.

​	要在第一个分隔符实例周围分割，请参见[Cut 函数](#func-cut-go118)。

#### Split Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}
Output:

["a" "b" "c"]
["" "man " "plan " "canal panama"]
[" " "x" "y" "z" " "]
[""]
```

### func SplitAfter 

``` go 
func SplitAfter(s, sep string) []string
```

SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.

​	SplitAfter函数将字符串s按sep分割成每个实例后的所有子字符串，并返回它们的子字符串切片。

If s does not contain sep and sep is not empty, SplitAfter returns a slice of length 1 whose only element is s.

​	如果s不包含sep且sep不为空，则SplitAfter返回长度为1的切片，其中唯一的元素为s。

If sep is empty, SplitAfter splits after each UTF-8 sequence. If both s and sep are empty, SplitAfter returns an empty slice.

​	如果sep为空，则SplitAfter在每个UTF-8序列之后拆分。如果s和sep都为空，则SplitAfter返回一个空切片。

It is equivalent to SplitAfterN with a count of -1.

​	它等价于带有计数`-1`的SplitAfterN函数。

#### SplitAfter Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
}
Output:

["a," "b," "c"]
```

### func SplitAfterN 

``` go 
func SplitAfterN(s, sep string, n int) []string
```

SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.

​	SplitAfterN函数将字符串s按sep拆分为每个实例之后的子字符串，并返回它们的子字符串切片。

The count determines the number of substrings to return:

​	计数参数n确定要返回的子字符串的数量：

n > 0：最多n个子字符串；最后一个子字符串将是未拆分的剩余部分。 

n == 0：结果为nil(零个子字符串) 

n < 0：所有子字符串

Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for SplitAfter.

​	s和sep的边缘情况(例如，空字符串)的处理方式如SplitAfter函数文档中所述。

#### SplitAfterN Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
}
Output:

["a," "b,c"]
```

### func SplitN 

``` go 
func SplitN(s, sep string, n int) []string
```

SplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators.

​	SplitN函数将字符串s按sep分割为每个分隔符之间的子字符串，并返回它们的子字符串切片。

The count determines the number of substrings to return:

​	计数参数n确定要返回的子字符串的数量：

n > 0：最多n个子字符串；最后一个子字符串将是未拆分的剩余部分。 

n == 0：结果为nil(零个子字符串) 

n < 0：所有子字符串

Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for Split.

​	对于 s 和 sep 的边界情况(例如空字符串)，将按照 Split函数文档中描述的方式处理。

To split around the first instance of a separator, see Cut.

​	如果要在第一个分隔符周围进行分割，请参阅 [Cut函数](#func-cut-go118)。

#### SplitN Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
}
Output:

["a" "b,c"]
[] (nil = true)
```

### func Title <- DEPRECATED

```go
func Title(s string) string
```

Title returns a copy of the string s with all Unicode letters that begin words mapped to their Unicode title case.

Deprecated: The rule Title uses for word boundaries does not handle Unicode punctuation properly. Use golang.org/x/text/cases instead.

#### Title Example

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Compare this example to the ToTitle example.
	fmt.Println(strings.Title("her royal highness"))
	fmt.Println(strings.Title("loud noises"))
	fmt.Println(strings.Title("хлеб"))
}
Output:

Her Royal Highness
Loud Noises
Хлеб
```



### func ToLower 

``` go 
func ToLower(s string) string
```

ToLower returns s with all Unicode letters mapped to their lower case.

​	ToLower函数返回s的所有Unicode字母均被映射为它们的小写形式的副本。

#### ToLower Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToLower("Gopher"))
}
Output:

gopher
```

### func ToLowerSpecial 

``` go 
func ToLowerSpecial(c unicode.SpecialCase, s string) string
```

ToLowerSpecial returns a copy of the string s with all Unicode letters mapped to their lower case using the case mapping specified by c.

​	ToLowerSpecial函数返回s的所有Unicode字母均被映射为它们的小写形式的副本，使用由c指定的大小写映射。

#### ToLowerSpecial Example
``` go 
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
}
Output:

önnek iş
```

### func ToTitle 

``` go 
func ToTitle(s string) string
```

ToTitle returns a copy of the string s with all Unicode letters mapped to their Unicode title case.

​	ToTitle函数返回s的所有Unicode字母均被映射为它们的Unicode标题形式的副本。

#### ToTitle Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Compare this example to the Title example.
	fmt.Println(strings.ToTitle("her royal highness"))
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))
}
Output:

HER ROYAL HIGHNESS
LOUD NOISES
ХЛЕБ
```

### func ToTitleSpecial 

``` go 
func ToTitleSpecial(c unicode.SpecialCase, s string) string
```

ToTitleSpecial returns a copy of the string s with all Unicode letters mapped to their Unicode title case, giving priority to the special casing rules.

​	ToTitleSpecial函数返回s的所有Unicode字母均被映射为它们的Unicode标题形式的副本，优先考虑特殊的大小写规则。

#### ToTitleSpecial Example
``` go 
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
}
Output:

DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR
```

### func ToUpper 

``` go 
func ToUpper(s string) string
```

ToUpper returns s with all Unicode letters mapped to their upper case.

​	ToUpper函数返回s的所有Unicode字母均被映射为它们的大写形式的副本。

#### ToUpper Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("Gopher"))
}
Output:

GOPHER
```

### func ToUpperSpecial 

``` go 
func ToUpperSpecial(c unicode.SpecialCase, s string) string
```

ToUpperSpecial returns a copy of the string s with all Unicode letters mapped to their upper case using the case mapping specified by c.

​	ToUpperSpecial函数返回s的所有Unicode字母均被映射为它们的大写形式的副本，使用由c指定的大小写映射。

#### ToUpperSpecial Example
``` go 
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}
Output:

ÖRNEK İŞ
```

### func ToValidUTF8  <- go1.13

``` go 
func ToValidUTF8(s, replacement string) string
```

ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences replaced by the replacement string, which may be empty.

​	ToValidUTF8函数返回s的每个无效UTF-8字节序列的运行都被替换为替换字符串的副本，替换字符串可以为空。

### func Trim 

``` go 
func Trim(s, cutset string) string
```

Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.

​	Trim函数返回字符串s的所有前导和尾随Unicode码点都包含在cutset中，这些码点被删除的切片。

#### Trim Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
}
Output:

Hello, Gophers
```

### func TrimFunc 

``` go 
func TrimFunc(s string, f func(rune) bool) string
```

TrimFunc returns a slice of the string s with all leading and trailing Unicode code points c satisfying f(c) removed.

​	TrimFunc函数返回字符串 s 中所有符合函数 f 的 Unicode 码点的前缀和后缀都被删除后的子字符串。

#### TrimFunc Example
``` go 
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}
Output:

Hello, Gophers
```

### func TrimLeft 

``` go 
func TrimLeft(s, cutset string) string
```

TrimLeft returns a slice of the string s with all leading Unicode code points contained in cutset removed.

​	TrimLeft函数返回字符串 s 去掉所有前导的包含在 cutset 中的 Unicode 码点后的子字符串。

To remove a prefix, use TrimPrefix instead.

​	如果要去掉一个前缀，请使用 TrimPrefix函数。

#### TrimLeft Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
}
Output:

Hello, Gophers!!!
```

### func TrimLeftFunc 

``` go 
func TrimLeftFunc(s string, f func(rune) bool) string
```

TrimLeftFunc returns a slice of the string s with all leading Unicode code points c satisfying f(c) removed.

​	TrimLeftFunc函数返回字符串 s 去掉所有前导符合函数 f 的 Unicode 码点后的子字符串。

#### TrimLeftFunc Example
``` go 
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}
Output:

Hello, Gophers!!!
```

### func TrimPrefix  <- go1.1

``` go 
func TrimPrefix(s, prefix string) string
```

TrimPrefix returns s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged.

​	TrimPrefix函数返回字符串 s 去掉提供的前缀字符串 prefix。如果 s 不以 prefix 开头，则返回 s 不变。

#### TrimPrefix Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimPrefix(s, "¡¡¡Hello, ")
	s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
	fmt.Print(s)
}
Output:

Gophers!!!
```

### xxxxxxxxxx1 1func (e *NumError) Unwrap() errorgo 

``` go 
func TrimRight(s, cutset string) string
```

TrimRight returns a slice of the string s, with all trailing Unicode code points contained in cutset removed.

​	TrimRight函数返回字符串 s 去掉所有后缀的包含在 cutset 中的 Unicode 码点后的子字符串。

To remove a suffix, use TrimSuffix instead.

​	如果要去掉一个后缀，请使用 TrimSuffix函数。

#### TrimRight Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
}
Output:

¡¡¡Hello, Gophers
```

### func TrimRightFunc 

``` go 
func TrimRightFunc(s string, f func(rune) bool) string
```

TrimRightFunc returns a slice of the string s with all trailing Unicode code points c satisfying f(c) removed.

​	TrimRightFunc函数返回字符串 s 去掉所有后缀符合函数 f 的 Unicode 码点后的子字符串。

#### TrimRightFunc Example
``` go 
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}
Output:

¡¡¡Hello, Gophers
```

### func TrimSpace 

``` go 
func TrimSpace(s string) string
```

TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.

​	TrimSpace函数返回字符串 s 去掉所有前导和后缀的空白字符，根据 Unicode 定义。

#### TrimSpace Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
}
Output:

Hello, Gophers
```

### func TrimSuffix  <- go1.1

``` go 
func TrimSuffix(s, suffix string) string
```

TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.

​	TrimSuffix函数返回字符串 s 去掉提供的后缀字符串 suffix。如果 s 不以 suffix 结尾，则返回 s 不变。

#### TrimSuffix Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimSuffix(s, ", Gophers!!!")
	s = strings.TrimSuffix(s, ", Marmots!!!")
	fmt.Print(s)
}
Output:

¡¡¡Hello
```

## 类型

### type Builder  <- go1.10

``` go 
type Builder struct {
    addr *Builder // of receiver, to detect copies by value
	buf  []byte
}
```

A Builder is used to efficiently build a string using Write methods. It minimizes memory copying. The zero value is ready to use. Do not copy a non-zero Builder.

​	Builder结构体用于使用 Write 方法高效地构建字符串。它最小化了内存复制。零值可直接使用。不要复制非零值 Builder。

#### Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())

}
Output:

3...2...1...ignition
```

#### (*Builder) Cap  <- go1.12

``` go 
func (b *Builder) Cap() int
```

Cap returns the capacity of the builder's underlying byte slice. It is the total space allocated for the string being built and includes any bytes already written.

​	Cap方法返回 builder 底层字节切片的容量。它是为正在构建的字符串分配的总空间，包括已经写入的任何字节。

#### (*Builder) Grow  <- go1.10

``` go 
func (b *Builder) Grow(n int)
```

Grow grows b's capacity, if necessary, to guarantee space for another n bytes. After Grow(n), at least n bytes can be written to b without another allocation. If n is negative, Grow panics.

​	Grow方法按需增加 b 的容量，以保证另外 n 个字节的空间。调用 Grow(n) 后，至少可以将 n 个字节写入 b，而不必另行分配。如果 n 为负数，则 Grow 会出现 panic。

#### (*Builder) Len  <- go1.10

``` go 
func (b *Builder) Len() int
```

Len returns the number of accumulated bytes; b.Len() == len(b.String()).

​	Len方法返回已经累积的字节数；b.Len() == len(b.String())。

#### (*Builder) Reset  <- go1.10

``` go 
func (b *Builder) Reset()
```

Reset resets the Builder to be empty.

​	Reset方法将 Builder 重置为空。

#### (*Builder) String  <- go1.10

``` go 
func (b *Builder) String() string
```

String returns the accumulated string.

​	String方法返回已经累积的字符串。

#### (*Builder) Write  <- go1.10

``` go 
func (b *Builder) Write(p []byte) (int, error)
```

Write appends the contents of p to b's buffer. Write always returns len(p), nil.

​	Write方法将 p 的内容追加到 b 的缓冲区中。Write 总是返回 len(p) 和 nil。

#### (*Builder) WriteByte  <- go1.10

``` go 
func (b *Builder) WriteByte(c byte) error
```

WriteByte appends the byte c to b's buffer. The returned error is always nil.

​	WriteByte方法将字节 c 追加到 b 的缓冲区中。返回的错误始终为 nil。

#### (*Builder) WriteRune  <- go1.10

``` go 
func (b *Builder) WriteRune(r rune) (int, error)
```

WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer. It returns the length of r and a nil error.

​	WriteRune方法将 Unicode 码点 r 的 UTF-8 编码附加到 b 的缓冲区中。它返回 r 的长度和 nil 错误。

#### (*Builder) WriteString  <- go1.10

``` go 
func (b *Builder) WriteString(s string) (int, error)
```

WriteString appends the contents of s to b's buffer. It returns the length of s and a nil error.

​	WriteString方法将字符串s的内容附加到b的缓冲区中。它返回s的长度和一个nil错误。

### type Reader 

``` go 
type Reader struct {
	// contains filtered or unexported fields
}
```

A Reader implements the io.Reader, io.ReaderAt, io.ByteReader, io.ByteScanner, io.RuneReader, io.RuneScanner, io.Seeker, and io.WriterTo interfaces by reading from a string. The zero value for Reader operates like a Reader of an empty string.

​	Reader结构体通过从字符串中读取数据来实现io.Reader、io.ReaderAt、io.ByteReader、io.ByteScanner、io.RuneReader、io.RuneScanner、io.Seeker和io.WriterTo接口。Reader的零值像一个空字符串的Reader。

#### func NewReader 

``` go 
func NewReader(s string) *Reader
```

NewReader returns a new Reader reading from s. It is similar to bytes.NewBufferString but more efficient and read-only.

​	NewReader函数返回一个从s读取的新Reader。它类似于bytes.NewBufferString，但更高效且只读。

#### (*Reader) Len 

``` go 
func (r *Reader) Len() int
```

Len returns the number of bytes of the unread portion of the string.

​	Len方法返回未读部分字符串的字节数。

#### (*Reader) Read 

``` go 
func (r *Reader) Read(b []byte) (n int, err error)
```

Read implements the io.Reader interface.

​	Read方法实现了io.Reader接口。

#### (*Reader) ReadAt 

``` go 
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
```

ReadAt implements the io.ReaderAt interface.

​	ReadAt方法实现了io.ReaderAt接口。

#### (*Reader) ReadByte 

``` go 
func (r *Reader) ReadByte() (byte, error)
```

ReadByte implements the io.ByteReader interface.

​	ReadByte方法实现了io.ByteReader接口。

#### (*Reader) ReadRune 

``` go 
func (r *Reader) ReadRune() (ch rune, size int, err error)
```

ReadRune implements the io.RuneReader interface.

​	ReadRune方法实现了io.RuneReader接口。

#### (*Reader) Reset  <- go1.7

``` go 
func (r *Reader) Reset(s string)
```

Reset resets the Reader to be reading from s.

​	Reset方法将Reader重置为从s中读取。

#### (*Reader) Seek 

``` go 
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```

Seek implements the io.Seeker interface.

​	Seek方法实现了io.Seeker接口。

#### (*Reader) Size  <- go1.5

``` go 
func (r *Reader) Size() int64
```

Size returns the original length of the underlying string. Size is the number of bytes available for reading via ReadAt. The returned value is always the same and is not affected by calls to any other method.

​	Size方法返回底层字符串的原始长度。Size方法是通过ReadAt方法可读取的字节数。返回值总是相同的，并不受任何其他方法的调用影响。

#### (*Reader) UnreadByte 

``` go 
func (r *Reader) UnreadByte() error
```

UnreadByte implements the io.ByteScanner interface.

​	UnreadByte方法实现了 io.ByteScanner 接口。

#### (*Reader) UnreadRune 

``` go 
func (r *Reader) UnreadRune() error
```

UnreadRune implements the io.RuneScanner interface.

​	UnreadRune方法实现了 io.RuneScanner 接口。

#### (*Reader) WriteTo  <- go1.1

``` go 
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```

WriteTo implements the io.WriterTo interface.

​	WriteTo方法实现了 io.WriterTo 接口。

### type Replacer 

``` go 
type Replacer struct {
	// contains filtered or unexported fields
}
```

Replacer replaces a list of strings with replacements. It is safe for concurrent use by multiple goroutines.

​	Replacer结构体可以用一组字符串替换另一组字符串。它可以被多个 goroutine 并发使用。

#### func NewReplacer 

``` go 
func NewReplacer(oldnew ...string) *Replacer
```

NewReplacer returns a new Replacer from a list of old, new string pairs. Replacements are performed in the order they appear in the target string, without overlapping matches. The old string comparisons are done in argument order.

​	NewReplacer函数通过一组 old,new 字符串对返回一个新的 Replacer。替换按目标字符串中它们出现的顺序执行，而不会重叠匹配。old 字符串的比较按参数顺序执行。

NewReplacer panics if given an odd number of arguments.

​	如果传入奇数个参数，NewReplacer函数将 panic。

##### NewReplacer Example
``` go 
package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))
}
Output:

This is &lt;b&gt;HTML&lt;/b&gt;!
```

#### (*Replacer) Replace 

``` go 
func (r *Replacer) Replace(s string) string
```

Replace returns a copy of s with all replacements performed.

​	Replace方法返回执行所有替换后的 s 的副本。

#### (*Replacer) WriteString 

``` go 
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
```

WriteString writes s to w with all replacements performed.

​	WriteString方法将 s 写入 w，执行所有替换。