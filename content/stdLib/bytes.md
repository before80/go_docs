+++
title = "bytes"
linkTitle = "bytes"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/bytes@go1.20.1

​	bytes包实现了操作字节切片的功能。它类似于strings包的功能。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/bytes/buffer.go;l=191)

``` go 
const MinRead = 512
```

​	MinRead是Buffer.ReadFrom()方法在调用Buffer.Read()时传递的最小切片大小。只要Buffer至少有MinRead字节超出了`r`内容的所需的空间，Buffer.ReadFrom()就不会增长底层的缓冲区。

## 变量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/bytes/buffer.go;l=44)

``` go 
var ErrTooLarge = errors.New("bytes.Buffer: too large")
```

​	如果无法分配内存以存储缓冲区中的数据，ErrTooLarge将被传递给panic。

## 函数 

### func Clone  <- go1.20

``` go 
func Clone(b []byte) []byte
```

​	Clone函数返回b[:len(b)]的副本。其结果可能有额外的未使用的容量。	Clone(nil)返回nil。

### func Compare 

``` go 
func Compare(a, b []byte) int
```

​	Compare函数返回一个整数，按字典顺序比较两个字节切片。如果a == b，结果是0，如果a < b，结果是-1，如果a > b，结果是+1。nil参数等效于空切片。

#### Compare Example
``` go 
package main

import (
	"bytes"
)

func main() {
	// 将Compare的结果与零进行比较来解释。
	var a, b []byte
	if bytes.Compare(a, b) < 0 {
		// a less b
	}
	if bytes.Compare(a, b) <= 0 {
		// a less or equal b
	}
	if bytes.Compare(a, b) > 0 {
		// a greater b
	}
	if bytes.Compare(a, b) >= 0 {
		// a greater or equal b
	}

	// 倾向于用Equal来进行相等比较。
	if bytes.Equal(a, b) {
		// a equal b
	}
	if !bytes.Equal(a, b) {
		// a not equal b
	}
}

```

#### Compare Example (Search)
``` go 
package main

import (
	"bytes"
	"sort"
)

func main() {
	// 通过二进制搜索找到匹配的字节切片。
	var needle []byte
	var haystack [][]byte // 假设已排序
	i := sort.Search(len(haystack), func(i int) bool {
		// Return haystack[i] >= needle.
		return bytes.Compare(haystack[i], needle) >= 0
	})
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		// Found it!
	}
}

```

### func Contains 

``` go 
func Contains(b, subslice []byte) bool
```

​	Contains函数报告子字节切片subslice是否在字节切片b中。

#### Contains Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("foo")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("bar")))
	fmt.Println(bytes.Contains([]byte("seafood"), []byte("")))
	fmt.Println(bytes.Contains([]byte(""), []byte("")))
}
Output:

true
false
true
true
```

### func ContainsAny  <- go1.7

``` go 
func ContainsAny(b []byte, chars string) bool
```

​	ContainsAny函数报告chars中是否存在任何一个使用UTF-8编码的码点在字节切片b中。

#### ContainsAny Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "fÄo!"))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), "去是伟大的."))
	fmt.Println(bytes.ContainsAny([]byte("I like seafood."), ""))
	fmt.Println(bytes.ContainsAny([]byte(""), ""))
}
Output:

true
true
false
false
```

### func ContainsRune  <- go1.7

``` go 
func ContainsRune(b []byte, r rune) bool
```

​	ContainsRune函数报告符文`r`是否包含在UTF-8编码的字节切片b中。

#### ContainsRune Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'f'))
	fmt.Println(bytes.ContainsRune([]byte("I like seafood."), 'ö'))
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '大'))
	fmt.Println(bytes.ContainsRune([]byte("去是伟大的!"), '!'))
	fmt.Println(bytes.ContainsRune([]byte(""), '@'))
}
Output:

true
false
true
true
false
```

### func Count 

``` go 
func Count(s, sep []byte) int
```

​	Count函数计算s中sep的非重叠实例数。如果 sep 是一个空切片，则 Count 返回 s 中 UTF-8 编码的码点数加 1。

#### Count Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Count([]byte("cheese"), []byte("e")))
	fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each rune
}
Output:

3
5
```

### func Cut  <- go1.18

``` go 
func Cut(s, sep []byte) (before, after []byte, found bool)
```

​	Cut 函数将 s 切片在第一个 sep 的实例处分割，返回 sep 前面和后面的文本。 found的结果报告sep是否出现在s中。如果 sep 不出现在 s 中，则 Cut函数返回 s，nil，false。

​	Cut函数返回原始切片s的切片，而不是副本。

#### Cut Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	show := func(s, sep string) {
		before, after, found := bytes.Cut([]byte(s), []byte(sep))
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
func CutPrefix(s, prefix []byte) (after []byte, found bool)
```

​	CutPrefix函数返回没有提供前缀 prefix 的 s，并报告它是否找到了前缀。如果 s 不以 prefix 开头，则 CutPrefix函数返回 s、false。如果 prefix 是空字节切片，则 CutPrefix函数返回 s、true。

​	CutPrefix函数返回原始切片s的切片，而不是副本。

#### CutPrefix Example

```go 
s := []byte("foo bar foo bar foo")
prefix := []byte("foo ")
after, found := bytes.CutPrefix(s, prefix) 
// after = []byte("bar foo bar foo"), found = true
```



### func CutSuffix  <- go1.20

``` go 
func CutSuffix(s, suffix []byte) (before []byte, found bool)
```

​	CutSuffix函数返回没有提供后缀字节切片的s，并报告它是否找到了后缀。如果s没有以后缀结束，CutSuffix返回s，false。如果后缀是空字节切片，CutSuffix返回s，true。

​	CutSuffix 函数返回没有提供后缀 suffix 的 s，并报告它是否找到了后缀。如果 s 不以 suffix 结尾，则 CutSuffix 返回 s、false。如果 suffix 是空字节切片，则 CutSuffix函数返回 s、true。

​	CutSuffix函数返回原始切片s的切片，而不是副本。

#### CutSuffix Example

```go 
s := []byte("foo bar foo bar foo")
suffix := []byte(" foo")
before, found := bytes.CutSuffix(s, suffix) 
// before = []byte("foo bar foo bar"), found = true
```



### func Equal 

``` go 
func Equal(a, b []byte) bool
```

​	Equal 函数报告 a 和 b 是否具有相同的长度并包含相同的字节。nil 参数等价于空切片。

#### Equal Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Equal([]byte("Go"), []byte("Go")))
	fmt.Println(bytes.Equal([]byte("Go"), []byte("C++")))
}
Output:

true
false
```

### func EqualFold 

``` go 
func EqualFold(s, t []byte) bool
```

​	EqualFold 函数报告将 s 和 t(解释为 UTF-8 字符串)在简单的 Unicode 大小写折叠下是否相等，这是一种更通用的不区分大小写的形式。

#### EqualFold Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go")))
}
Output:

true
```

### func Fields 

``` go 
func Fields(s []byte) [][]byte
```

​	Fields函数将s解释为UTF-8编码的码点序列，并根据unicode.IsSpace定义的一个或多个连续的空白字符实例分隔切片s，返回s的子切片的切片，如果s仅包含空白，则返回一个空切片。

#### Fields Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
    // %q：输出带有双引号的字符串，同时可以将特殊字符以转义形式输出。
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
	fmt.Printf("Fields are: %#v\n", bytes.Fields([]byte("  foo bar  baz   ")))
}
Output:

Fields are: ["foo" "bar" "baz"]
Fields are: [][]uint8{[]uint8{0x66, 0x6f, 0x6f}, []uint8{0x62, 0x61, 0x72}, []uint8{0x62, 0x61, 0x7a}}
```

### func FieldsFunc 

``` go 
func FieldsFunc(s []byte, f func(rune) bool) [][]byte
```

​	FieldsFunc函数将s解释为UTF-8编码的码点序列，并根据满足f(c)的码点c的每个运行位置分隔切片s，并返回s的子切片。如果s中所有码点都满足f(c)，或者len(s) == 0，则返回一个空切片。

​	FieldsFunc函数对于调用f(c)的顺序没有保证，并且假定对于给定的c，f始终返回相同的值。

> 需要注意的是，bytes.FieldsFunc函数对于调用f(c)的顺序没有保证，这是因为FieldsFunc函数是并行调用f(c)的，以便更快地处理输入。

#### FieldsFunc Example
``` go 
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))
    
    s := []byte("a b c")
    fields := bytes.FieldsFunc(s, func(c rune) bool {
        return true
    })
    fmt.Printf("%q\n", fields)
    
    fields := bytes.FieldsFunc(s, func(c rune) bool {
    	return (int(c)-int('a'))%2 == 1
	})
    fmt.Printf("%q\n", fields)
}
Output:

Fields are: ["foo1" "bar2" "baz3"]
[]
["a " " c"]
```

### func HasPrefix 

``` go 
func HasPrefix(s, prefix []byte) bool
```

​	HasPrefix测试字节切片 s 是否以前缀 prefix 开头。

#### HasPrefix Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("")))
}
Output:

true
false
true
```

### func HasSuffix 

``` go 
func HasSuffix(s, suffix []byte) bool
```

​	HasSuffix函数测试字节切片s是否以后缀 suffix 结尾。

#### HasSuffix Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("go")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("O")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("Ami")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("")))
}
Output:

true
false
false
true
```

### func Index 

``` go 
func Index(s, sep []byte) int
```

​	Index函数返回sep在s中第一次出现的索引，如果sep不在s中，则返回-1。

#### Index Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Index([]byte("chicken"), []byte("ken")))
	fmt.Println(bytes.Index([]byte("chicken"), []byte("dmr")))
}
Output:

4
-1
```

### func IndexAny 

``` go 
func IndexAny(s []byte, chars string) int
```

​	IndexAny函数将s解释为UTF-8编码的Unicode码点序列。它返回chars中任何一个Unicode码点在s中第一次出现的字节索引。如果chars为空或在s中没有公共的码点，则返回-1。

#### IndexAny Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.IndexAny([]byte("chicken"), "aeiouy"))
	fmt.Println(bytes.IndexAny([]byte("crwth"), "aeiouy"))
}
Output:

2
-1
```

### func IndexByte 

``` go 
func IndexByte(b []byte, c byte) int
```

​	IndexByte函数返回c在b中第一次出现的索引，如果c不在b中，则返回-1。

#### IndexByte Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('k')))
	fmt.Println(bytes.IndexByte([]byte("chicken"), byte('g')))
}
Output:

4
-1
```

### func IndexFunc 

``` go 
func IndexFunc(s []byte, f func(r rune) bool) int
```

​	IndexFunc函数将s解释为UTF-8编码的码点序列。它返回s中第一个满足f(c)的Unicode码点的字节索引，如果没有则返回-1。

#### IndexFunc Example
``` go 
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(bytes.IndexFunc([]byte("Hello, 世界"), f))
	fmt.Println(bytes.IndexFunc([]byte("Hello, world"), f))
}
Output:

7
-1
```

### func IndexRune 

``` go 
func IndexRune(s []byte, r rune) int
```

​	IndexRune函数将s解释为UTF-8编码的码点序列。它返回给定符文在s中第一次出现的字节索引。如果rune不存在于s中，则返回-1。如果r是utf8.RuneError，则返回任何无效UTF-8字节序列的第一个实例。

#### IndexRune Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'k'))
	fmt.Println(bytes.IndexRune([]byte("chicken"), 'd'))
}
Output:

4
-1
```

### func Join 

``` go 
func Join(s [][]byte, sep []byte) []byte
```

Join concatenates the elements of s to create a new byte slice. The separator sep is placed between elements in the resulting slice.

Join函数将s的元素连接起来，创建一个新的字节切片。分隔符 sep 被放置在结果片断的元素之间。

​	Join函数将s中的元素连接起来以创建一个新的字节切片。分隔符sep放置在结果切片中的元素之间。

#### Join Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s", bytes.Join(s, []byte(", ")))
}
Output:

foo, bar, baz
```

### func LastIndex 

``` go 
func LastIndex(s, sep []byte) int
```

​	LastIndex函数返回sep在s中最后一次出现的索引，如果sep不在s中，则返回-1。

#### LastIndex Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.Index([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("rodent")))
}
Output:

0
3
-1
```

### func LastIndexAny 

``` go 
func LastIndexAny(s []byte, chars string) int
```

​	LastIndexAny函数将s解释为UTF-8编码的Unicode码点序列。它返回s中任何一个Unicode代码点(chars中的)的最后一次出现的字节索引。如果chars为空或没有共同码点，则返回-1。

#### LastIndexAny Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "MüQp"))
	fmt.Println(bytes.LastIndexAny([]byte("go 地鼠"), "地大"))
	fmt.Println(bytes.LastIndexAny([]byte("go gopher"), "z,!."))
}
Output:

5
3
-1
```

### func LastIndexByte  <- go1.5

``` go 
func LastIndexByte(s []byte, c byte) int
```

​	LastIndexByte函数返回c在s中最后一次出现的实例索引，如果c不在s中，则返回-1。

#### LastIndexByte Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('g')))
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('r')))
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('z')))
}
Output:

3
8
-1
```

### func LastIndexFunc 

``` go 
func LastIndexFunc(s []byte, f func(r rune) bool) int
```

​	LastIndexFunc将s解释为UTF-8编码的码点序列。它返回s中最后一个满足f(c)的Unicode码点(rune)的字节索引，如果没有符合条件的，则返回-1。

#### LastIndexFunc Example
``` go 
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsLetter))
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsPunct))
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsNumber))
}
Output:

8
9
-1
```

### func Map 

``` go 
func Map(mapping func(r rune) rune, s []byte) []byte
```

​	Map函数返回一个字节切片s的副本，其中所有字符根据映射函数进行修改。如果映射函数返回负值，则从字节切片中删除该字符而不作替换。s和输出中的字符被解释为UTF-8编码的码点。

#### Map Example

```go 
package main

import (
    "bytes"
    "fmt"
)

func main() {
    s := []byte("Hello, World!")
    mapping := func(r rune) rune {
        if r == 'o' {
            return '0'
        }
        return r
    }
    newS := bytes.Map(mapping, s)
    fmt.Printf("%s\n", newS)
}
output:

Hell0, W0rld!
```



### func Repeat 

``` go 
func Repeat(b []byte, count int) []byte
```

​	Repeat函数返回一个新的字节切片，其中包含count个b的副本。

​	如果count为负数或者(len(b) * count)的结果溢出，它就会发生panic。

#### Repeat Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("ba%s", bytes.Repeat([]byte("na"), 2))
}
Output:

banana
```

### func Replace 

``` go 
func Replace(s, old, new []byte, n int) []byte
```

Replace returns a copy of the slice s with the first n non-overlapping instances of old replaced by new. If old is empty, it matches at the beginning of the slice and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune slice. If n < 0, there is no limit on the number of replacements.

Replace函数返回一个片断s的副本，用new替换old的前n个不重叠的实例。如果old是空的，它在片断的开头和每个UTF-8序列之后进行匹配，对于一个k-rune片断，最多产生k+1个替换。如果n<0，则对替换的数量没有限制。

#### Replace Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("oink"), []byte("moo"), -1))
}
Output:

oinky oinky oink
moo moo moo
```

### func ReplaceAll  <- go1.12

``` go 
func ReplaceAll(s, old, new []byte) []byte
```

​	ReplaceAll函数返回将 s 中所有非重叠的 old 替换为 new 后得到的新 byte 切片。如果 old 是空的，则匹配从切片开头和每个 UTF-8 序列之后开始，最多得到 k+1 次替换，其中 k 是切片中的 Unicode 码点数。

#### ReplaceAll Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", bytes.ReplaceAll([]byte("oink oink oink"), []byte("oink"), []byte("moo")))
    fmt.Printf("%s\n", bytes.ReplaceAll([]byte("LOVE"), []byte(""), []byte("-123-")))
}
Output:

moo moo moo
-123-L-123-O-123-V-123-E-123-
```

### func Runes 

``` go 
func Runes(s []byte) []rune
```

​	Runes函数将 s 解释为一系列 UTF-8 编码的码点。它返回一个等价于 s 的码点切片(Unicode 码点)。

#### Runes Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	rs := bytes.Runes([]byte("go gopher"))
	for _, r := range rs {
        // %U 格式化操作符将整数值格式化为 Unicode 字符。
		// %#U 格式化操作符将整数值格式化为 Unicode 字符，
        // 并且在输出结果中附加 Unicode 字符的码点(code point)
		fmt.Printf("%#U\n", r)
	}
}
Output:

U+0067 'g'
U+006F 'o'
U+0020 ' '
U+0067 'g'
U+006F 'o'
U+0070 'p'
U+0068 'h'
U+0065 'e'
U+0072 'r'
```

### func Split 

``` go 
func Split(s, sep []byte) [][]byte
```

​	Split函数将s切成由sep分隔的所有子片，并返回这些分隔符之间的子片的一个片断。如果sep为空，Split会在每个UTF-8序列之后进行分割。它等同于SplitN，计数为-1。

​	要围绕第一个分隔符进行分割，请参见 [Cut](#func-cut-go118)。

#### Split Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
}
Output:

["a" "b" "c"]
["" "man " "plan " "canal panama"]
[" " "x" "y" "z" " "]
[""]
```

### func SplitAfter 

``` go 
func SplitAfter(s, sep []byte) [][]byte
```

​	SplitAfter函数将 s 切分成在每个 sep 实例之后的所有子切片，并返回这些子切片的切片。如果 sep 是空的，则 SplitAfter 在每个 UTF-8 序列之后分隔。它等同于SplitAfterN，计数为-1。

#### SplitAfter Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(",")))
}
Output:

["a," "b," "c"]
```

### func SplitAfterN 

``` go 
func SplitAfterN(s, sep []byte, n int) [][]byte
```

​	SplitAfterN 将 s 切分成每个 sep 实例之后的子切片，并返回这些子切片的切片。如果 sep 是空的，则 SplitAfterN 在每个 UTF-8 序列之后分隔。n 决定要返回的子切片数：



n > 0：最多 n 个子切片；最后一个子切片将是未拆分的剩余部分。

 n == 0：结果为 nil(零个子切片)

 n < 0：所有子切片

#### SplitAfterN Example

``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%q\n", bytes.SplitAfterN([]byte("a,b,c"), []byte(","), 2))
}
Output:

["a," "b,c"]
```

### func SplitN 

``` go 
func SplitN(s, sep []byte, n int) [][]byte
```

​	SplitN函数将 s 切分成由 sep 分隔的子切片，并返回这些分隔符之间的子切片的切片。如果 sep 是空的，则 SplitN 在每个 UTF-8 序列之后分隔。n 决定要返回的子切片数：



n > 0：最多 n 个子切片；最后一个子切片将是未拆分的剩余部分。 

n == 0：结果为 nil(零个子切片)

n < 0：所有子切片

​	要围绕第一个分隔符进行分割，请参见 [Cut](#func-cut-go118)。

#### SplitN Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 2))
	z := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
}
Output:

["a" "b,c"]
[] (nil = true)
```



### func ToLower 

``` go 
func ToLower(s []byte) []byte
```

​	ToLower函数返回一个将 s 中所有 Unicode 字母都转为其小写形式的新字节切片。

#### ToLower Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.ToLower([]byte("Gopher")))
}
Output:

gopher
```

### func ToLowerSpecial 

``` go 
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
```

​	ToLowerSpecial函数将 s 视为 UTF-8 编码的字节切片，返回一个将其中所有 Unicode 字母都转为其小写形式的新字节切片，其规则按照特定的 casing 规则进行。

#### ToLowerSpecial Example
``` go 
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	str := []byte("AHOJ VÝVOJÁRİ GOLANG")
	totitle := bytes.ToLowerSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))
	fmt.Println("ToLower : " + string(totitle))
}
Output:

Original : AHOJ VÝVOJÁRİ GOLANG
ToLower : ahoj vývojári golang
```

### func ToTitle 

``` go 
func ToTitle(s []byte) []byte
```

​	ToTitle函数将 s 视为 UTF-8 编码的字节切片，返回一个将其中所有 Unicode 字母都转为其 title case 形式的新字节切片。

#### ToTitle Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s\n", bytes.ToTitle([]byte("loud noises")))
	fmt.Printf("%s\n", bytes.ToTitle([]byte("хлеб")))
}
Output:

LOUD NOISES
ХЛЕБ
```

### func ToTitleSpecial 

``` go 
func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte
```

​	ToTitleSpecia函数将 s 视为 UTF-8 编码的字节切片，返回一个将其中所有 Unicode 字母都转为其 title case 形式的新字节切片，其规则按照特定的 casing 规则进行。

#### ToTitleSpecial Example
``` go 
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	str := []byte("ahoj vývojári golang")
	totitle := bytes.ToTitleSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))
	fmt.Println("ToTitle : " + string(totitle))
}
Output:

Original : ahoj vývojári golang
ToTitle : AHOJ VÝVOJÁRİ GOLANG
```

### func ToUpper 

``` go 
func ToUpper(s []byte) []byte
```

​	ToUppe函数返回一个将 s 中所有 Unicode 字母都转为其大写形式的新字节切片。

> ​	`bytes.ToUpper` 和 `bytes.ToTitle` 都是 bytes 包提供的字符串转换方法，其主要区别在于转换的方式不同。
>
> ​	`bytes.ToUpper` 方法将字符串中的所有字母字符转换成大写形式，并返回转换后的新字符串。其转换方式是基于简单的 ASCII 字符编码表进行的，不考虑各个字符在 Unicode 编码中的位置或语言特定的大小写规则。
>
> ​	`bytes.ToTitle` 方法将字符串中的所有字母字符转换成标题形式，并返回转换后的新字符串。它按照 Unicode 标准中的特定规则进行大小写转换，具体来说，它将每个字母转换成其 Unicode 标题形式。这个标题形式与普通大写形式有时会有所不同，比如德语中的 "ß" 字符，它的标题形式是 "SS"。
>
> ​	因此，`bytes.ToTitle` 方法在进行字符大小写转换时比`bytes.ToUpper` 更加智能和准确。但是，由于其对 Unicode 的完全支持，可能会导致一些性能问题，特别是在处理非常大的字符串时。

#### ToUpper Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.ToUpper([]byte("Gopher")))
}
Output:

GOPHER
```

### func ToUpperSpecial 

``` go 
func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte
```

​		ToUpperSpecial函数将 s 视为 UTF-8 编码的字节切片，返回一个将其中所有 Unicode 字母都转为其大写形式的新字节切片，其规则按照特定的 casing 规则进行。

#### ToUpperSpecial Example
``` go 
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	str := []byte("ahoj vývojári golang")
	totitle := bytes.ToUpperSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))
	fmt.Println("ToUpper : " + string(totitle))
}
Output:

Original : ahoj vývojári golang
ToUpper : AHOJ VÝVOJÁRİ GOLANG
```

### func ToValidUTF8  <- go1.13

``` go 
func ToValidUTF8(s, replacement []byte) []byte
```

​	ToValidUTF8 函数将 s 视为 UTF-8 编码的字节并返回一份副本，其中每个表示无效 UTF-8 的字节序列被替换为 replacement 中的字节，replacement 可以为空。

### func Trim 

``` go 
func Trim(s []byte, cutset string) []byte
```

​	Trim函数返回一个将 s 去掉开头和结尾的包含在 cutset 中的所有 UTF-8 编码的码点后的新字节切片。

#### Trim Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("[%q]", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! "))
}
Output:

["Achtung! Achtung"]
```

### func TrimFunc 

``` go 
func TrimFunc(s []byte, f func(r rune) bool) []byte
```

​	TrimFunc函数返回一个将 s 去掉开头和结尾的所有满足 f(c) 的 UTF-8 编码的码点后的新字节切片。

#### TrimFunc Example
``` go 
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimFunc([]byte("\"go-gopher!\""), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}
Output:

-gopher!
"go-gopher!"
go-gopher
go-gopher!
```

### func TrimLeft 

``` go 
func TrimLeft(s []byte, cutset string) []byte
```

​	TrimLeft函数返回一个去除了s中前部包含在cutset中的Unicode编码字符的子字节切片。

#### TrimLeft Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Print(string(bytes.TrimLeft([]byte("453gopher8257"), "0123456789")))
}
Output:

gopher8257
```

### func TrimLeftFunc 

``` go 
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte
```

​	TrimLeftFunc函数返回一个去除了s中前部符合函数f(c) 的Unicode编码字符的子字节切片。

#### TrimLeftFunc Example
``` go 
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(string(bytes.TrimLeftFunc([]byte("go-gopher"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimLeftFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimLeftFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}
Output:

-gopher
go-gopher!
go-gopher!567
```

### func TrimPrefix  <- go1.1

``` go 
func TrimPrefix(s, prefix []byte) []byte
```

​	TrimPrefix函数返回一个去除了s中前缀prefix的字节切片。如果s不是以prefix开头，则返回s本身。

#### TrimPrefix Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b = []byte("Goodbye,, world!")
	b = bytes.TrimPrefix(b, []byte("Goodbye,"))
	b = bytes.TrimPrefix(b, []byte("See ya,"))
	fmt.Printf("Hello%s", b)
}
Output:

Hello, world!
```

### func TrimRight 

``` go 
func TrimRight(s []byte, cutset string) []byte
```

​	TrimRight函数返回一个去除了s中尾部包含在cutset中的Unicode编码字符的子字节切片。

#### TrimRight Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Print(string(bytes.TrimRight([]byte("453gopher8257"), "0123456789")))
}
Output:

453gopher
```

### func TrimRightFunc 

``` go 
func TrimRightFunc(s []byte, f func(r rune) bool) []byte
```

​	TrimRightFunc 函数返回一个 s 的子切片，该子切片去除了满足 f(c) 的所有后缀 UTF-8 编码的码点。

#### TrimRightFunc Example
``` go 
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(string(bytes.TrimRightFunc([]byte("go-gopher"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimRightFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimRightFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}
Output:

go-
go-gopher
1234go-gopher!
```

### func TrimSpace 

``` go 
func TrimSpace(s []byte) []byte
```

​	TrimSpace 函数返回 s 的子切片，该子切片去除了所有的前导和尾随 Unicode 定义的空格。

#### TrimSpace Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("%s", bytes.TrimSpace([]byte(" \t\n a lone gopher \n\t\r\n")))
}
Output:

a lone gopher
```

### func TrimSuffix  <- go1.1

``` go 
func TrimSuffix(s, suffix []byte) []byte
```

​	TrimSuffix 函数返回不包含指定后缀的字节切片 s。如果 s 不以后缀结尾，则返回未更改的 s。

#### TrimSuffix Example
``` go 
package main

import (
	"bytes"
	"os"
)

func main() {
	var b = []byte("Hello, goodbye, etc!")
	b = bytes.TrimSuffix(b, []byte("goodbye, etc!"))
	b = bytes.TrimSuffix(b, []byte("gopher"))
	b = append(b, bytes.TrimSuffix([]byte("world!"), []byte("x!"))...)
	os.Stdout.Write(b)
}
Output:

Hello, world
```

## 类型 

### type Buffer 

``` go 
type Buffer struct {
	buf      []byte // buf是一个字节切片，
    				//它的内容是buf[off:len(buf)]，即从off到buf末尾的部分。
	off      int    // off是一个int类型，
    				// 表示下一次读取操作应该在&buf[off]处开始，
    				// 下一次写入操作应该在&buf[len(buf)]处开始。
	lastRead readOp // lastRead是一个readOp类型的变量，
    				// 表示最后一次读取操作的类型，以便Unread*方法可以正确地工作。
}
```

​	Buffer是一个可变大小的字节缓冲区，具有Read和Write方法。Buffer的零值是一个空缓冲区，可以直接使用。

> Reader实现了io.Reader、io.ReaderAt、io.WriterTo、io.Seeker、io.ByteScanner和io.RuneScanner接口，通过从一个字节切片中读取数据。与Buffer不同，Reader是只读的并支持寻址。Reader的零值类似于一个空切片的Reader。

#### Buffer Example
``` go 
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)
}
Output:

Hello world!
```

#### Buffer Example(Reader)
``` go 
package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
)

func main() {
	// A Buffer can turn a string or a []byte into an io.Reader.
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}
Output:

Gophers rule!
```

#### func NewBuffer 

``` go 
func NewBuffer(buf []byte) *Buffer
```

​	NewBuffer函数使用buf作为其初始内容创建并初始化一个新的Buffer。新的Buffer接管了buf，并且调用方不应在此调用之后使用buf。NewBuffer用于准备从现有数据中读取数据的缓冲区。它还可用于设置写入的内部缓冲区的初始大小。为此，buf应该具有所需的容量，但长度为零。 

​	在大多数情况下，new(Buffer)(或仅声明一个Buffer变量)足以初始化一个Buffer。

#### func NewBufferString 

``` go 
func NewBufferString(s string) *Buffer
```

​	NewBufferString函数使用字符串s作为其初始内容创建并初始化一个新的Buffer。它用于准备从现有字符串中读取数据。

​	在大多数情况下，new(Buffer)(或仅声明一个Buffer变量)足以初始化一个Buffer。

#### (*Buffer) Bytes 

``` go 
func (b *Buffer) Bytes() []byte
```

​	Bytes方法返回一个长度为b.Len()的切片，其中包含缓冲区未读部分。切片仅在下一次缓冲区修改之前有效(即只在下一次像Read、Write、Reset或Truncate这样的方法调用之前有效)。切片至少与缓冲区内容同步，直到下一次缓冲区修改，因此立即更改切片将影响将来读取的结果。

##### Example
``` go 
package main

import (
	"bytes"
	"os"
)

func main() {
	buf := bytes.Buffer{}
	buf.Write([]byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'})
	os.Stdout.Write(buf.Bytes())
}
Output:

hello world
```

#### (*Buffer) Cap  <- go1.5

``` go 
func (b *Buffer) Cap() int
```

​	Cap方法返回缓冲区底层字节切片的容量，即分配给缓冲区数据的总空间。

##### Cap Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf1 := bytes.NewBuffer(make([]byte, 10))
	buf2 := bytes.NewBuffer(make([]byte, 0, 10))
	fmt.Println(buf1.Cap())
	fmt.Println(buf2.Cap())
}
Output:

10
10
```

#### (*Buffer) Grow  <- go1.1

``` go 
func (b *Buffer) Grow(n int)
```

​	Grow方法增加缓冲区的容量(必要时)，以保证另外n个字节的空间。调用Grow(n)之后，可以将至少n个字节写入缓冲区，而不需要另一个分配。如果n为负数，Grow方法将发生panic。如果缓冲区无法增长，它将发生ErrTooLarge的panic。

##### Grow Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)
	bb := b.Bytes()
	b.Write([]byte("64 bytes or fewer"))
	fmt.Printf("%q", bb[:b.Len()])
}
Output:

"64 bytes or fewer"
```

#### (*Buffer) Len 

``` go 
func (b *Buffer) Len() int
```

​	Len方法返回缓冲区未读部分的字节数；`b.Len() == len(b.Bytes())`。

##### Len Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	fmt.Printf("%d", b.Len())
}
Output:

5
```

#### (*Buffer) Next 

``` go 
func (b *Buffer) Next(n int) []byte
```

​	Next方法返回一个包含从缓冲区中取出的下一个 n 个字节的切片，并将缓冲区推进，就像这些字节已经被 Read 返回一样。如果缓冲区中的字节数少于 n，则 Next 返回整个缓冲区。该切片只在下一次读或写方法调用之前有效。

##### Next Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	fmt.Printf("%s\n", string(b.Next(2)))
	fmt.Printf("%s\n", string(b.Next(2)))
	fmt.Printf("%s", string(b.Next(2)))
}
Output:

ab
cd
e
```

#### (*Buffer) Read 

``` go 
func (b *Buffer) Read(p []byte) (n int, err error)
```

​	Read方法从缓冲区读取 len(p) 个字节或直到缓冲区被耗尽。返回值 n 是读取的字节数。如果缓冲区没有数据可返回，则 err 为 io.EOF， len(p) == 0否则为 nil。

##### Read Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	rdbuf := make([]byte, 1)
	n, err := b.Read(rdbuf)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(b.String())
	fmt.Println(string(rdbuf))
	// Output
	// 1
	// bcde
	// a
}

```

#### (*Buffer) ReadByte 

``` go 
func (b *Buffer) ReadByte() (byte, error)
```

​	ReadByte方法从缓冲区读取并返回下一个字节。如果没有字节可用，则返回错误 io.EOF。

##### ReadByte Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	b.Grow(64)
	b.Write([]byte("abcde"))
	c, err := b.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
	fmt.Println(b.String())
	// Output
	// 97
	// bcde
}

```

#### (*Buffer) ReadBytes 

``` go 
func (b *Buffer) ReadBytes(delim byte) (line []byte, err error)
```

​	ReadBytes方法从输入中读取，直到遇到第一个分隔符 delim，返回包含数据和分隔符的切片。如果在找到分隔符之前遇到错误，则返回在错误之前读取的数据和错误本身(通常是 io.EOF)。如果返回的数据不以 delim 结尾，则 ReadBytes 返回 err != nil。

#### (*Buffer) ReadFrom 

``` go 
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
```

​	ReadFrom方法从 `r` 中读取数据，直到遇到 EOF，然后将其附加到缓冲区中，`并根据需要增加缓冲区的大小`。返回值 n 是读取的字节数。读取期间遇到的除 io.EOF 以外的任何错误也将被返回。如果缓冲区变得太大，ReadFrom 方法将使用 ErrTooLarge 引发 panic。

#### (*Buffer) ReadRune 

``` go 
func (b *Buffer) ReadRune() (r rune, size int, err error)
```

​	ReadRune方法从缓冲区读取并返回下一个 UTF-8 编码的 Unicode 码点。如果没有字节可用，则返回的错误是 io.EOF。如果字节是错误的 UTF-8 编码，则它会消耗一个字节并返回 U+FFFD，1。

#### (*Buffer) ReadString 

``` go 
func (b *Buffer) ReadString(delim byte) (line string, err error)
```

​	ReadString方法从输入中读取，直到遇到第一个分隔符 delim，返回包含数据和分隔符的字符串。如果在找到分隔符之前遇到错误，则返回在错误之前读取的数据和错误本身(通常是 io.EOF)。如果返回的数据不以 delim 结尾，则 ReadString 返回 err != nil。

#### (*Buffer) Reset 

``` go 
func (b *Buffer) Reset()
```

​	Reset方法将缓冲区重置为空，但保留底层存储空间供未来写入使用。Reset方法相当于Truncate(0)。

#### (*Buffer) String 

``` go 
func (b *Buffer) String() string
```

​	String方法返回缓冲区未读部分的内容作为字符串。如果Buffer是nil指针，则返回"`<nil>`"。

​	为了更有效地构建字符串，可以使用strings.Builder类型。

#### (*Buffer) Truncate 

``` go 
func (b *Buffer) Truncate(n int)
```

​	Truncate方法将缓冲区保留未读的前n个字节，但继续使用相同的分配存储空间。如果n为负数或大于缓冲区的长度，则会引发panic。

#### (*Buffer) UnreadByte 

``` go 
func (b *Buffer) UnreadByte() error
```

​	UnreadByte方法撤消最近一次成功读取至少一个字节的读操作返回的最后一个字节。如果自上次读取以来已经发生写入，或者上次读取返回错误，或者读取了零个字节，则UnreadByte会返回错误。

#### (*Buffer) UnreadRune 

``` go 
func (b *Buffer) UnreadRune() error
```

​	UnreadRune方法撤消ReadRune方法返回的最后一个符文。如果缓冲区上最近一次读取或写入操作不是成功的ReadRune，则UnreadRune会返回错误。(在这方面，它比UnreadByte更严格，后者会从任何读操作中撤消最后一个字节。)

#### (*Buffer) Write 

``` go 
func (b *Buffer) Write(p []byte) (n int, err error)
```

​	Write方法将`p`的内容附加到缓冲区，必要时扩展缓冲区。返回值`n`是p的长度；err始终为nil。如果缓冲区变得太大，Write方法会发生ErrTooLarge的panic。

#### (*Buffer) WriteByte 

``` go 
func (b *Buffer) WriteByte(c byte) error
```

​	WriteByte方法将字节`c`附加到缓冲区，必要时扩展缓冲区。返回的错误始终为nil，但包含在其中是为了与bufio.Writer的WriteByte方法匹配。如果缓冲区变得太大，WriteByte方法会发生ErrTooLarge的panic。

#### (*Buffer) WriteRune 

``` go 
func (b *Buffer) WriteRune(r rune) (n int, err error)
```

​	WriteRune方法将Unicode码点`r`的UTF-8编码附加到缓冲区，返回其长度和错误(始终为nil，但包含在其中是为了与bufio.Writer的WriteRune方法匹配)。必要时扩展缓冲区；如果缓冲区变得太大，则WriteRune方法会发生ErrTooLarge的panic。

#### (*Buffer) WriteString 

``` go 
func (b *Buffer) WriteString(s string) (n int, err error)
```

​	WriteString方法将`s`的内容附加到缓冲区，必要时扩展缓冲区。返回值`n`是s的长度；err总是nil。如果缓冲区变得太大，WriteString方法会发生ErrTooLarge的panic。

#### (*Buffer) WriteTo 

``` go 
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)
```

​	WriteTo方法将缓冲区中的数据写入`w`中，直到缓冲区被耗尽或出现错误。返回值`n`为写入的字节数；它总是能够适合int，但是为了匹配io.WriterTo接口，它是int64类型。任何在写入过程中遇到的错误也会被返回。

### type Reader 

``` go 
type Reader struct {
	s        []byte
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}
```

​	Reader通过从一个字节切片读取实现io.Reader、io.ReaderAt、io.WriterTo、io.Seeker、io.ByteScanner和io.RuneScanner接口。与Buffer不同，Reader是只读的，并支持寻找。Reader的零值操作方式类似于空切片的Reader。

#### func NewReader 

``` go 
func NewReader(b []byte) *Reader
```

​	NewReader函数返回一个从`b`读取的新Reader。

#### (*Reader) Len 

``` go 
func (r *Reader) Len() int
```

​	Len方法返回未读部分的字节数。

##### Len Example
``` go 
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(bytes.NewReader([]byte("Hi!")).Len())
	fmt.Println(bytes.NewReader([]byte("こんにちは!")).Len())
}
Output:

3
16
```

#### (*Reader) Read 

``` go 
func (r *Reader) Read(b []byte) (n int, err error)
```

​	Read方法实现io.Reader接口。

#### (*Reader) ReadAt 

``` go 
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
```

​	ReadAt方法实现io.ReaderAt接口。

#### (*Reader) ReadByte 

``` go 
func (r *Reader) ReadByte() (byte, error)
```

​	ReadByte方法实现io.ByteReader接口。

#### (*Reader) ReadRune 

``` go 
func (r *Reader) ReadRune() (ch rune, size int, err error)
```

​	ReadRune方法实现io.RuneReader接口。

#### (*Reader) Reset  <- go1.7

``` go 
func (r *Reader) Reset(b []byte)
```

​	Reset方法将Reader重置为从b读取。

#### (*Reader) Seek 

``` go 
func (r *Reader) Seek(offset int64, whence int) (int64, error)
```

​	Seek方法实现io.Seeker接口。

#### (*Reader) Size  <- go1.5

``` go 
func (r *Reader) Size() int64
```

​	Size方法返回底层字节切片的原始长度。Size方法是可通过ReadAt方法读取的字节数。该结果不受任何方法调用的影响，除了Reset方法。

#### (*Reader) UnreadByte 

``` go 
func (r *Reader) UnreadByte() error
```

​	UnreadByte 方法实现了 io.ByteScanner 接口，用于将上一次读取的一个字节退回到 Reader 中。

#### (*Reader) UnreadRune 

``` go 
func (r *Reader) UnreadRune() error
```

​	UnreadRune 方法实现了 io.RuneScanner 接口，用于将上一次读取的一个符文退回到 Reader 中。

#### (*Reader) WriteTo  <- go1.1

``` go 
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
```

​	WriteTo 方法实现了 io.WriterTo 接口，用于将 Reader 中未读取的数据写入 w，直到读取完毕或出现错误。返回值 n 是写入的字节数，它总是可以用 int 表示，但为了匹配 io.WriterTo 接口，返回值类型为 int64。在写入过程中遇到的任何错误也会一并返回。