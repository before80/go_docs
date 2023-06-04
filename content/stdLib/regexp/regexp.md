+++
title = "regexp"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# regexp

https://pkg.go.dev/regexp@go1.20.1

​	regexp包实现了正则表达式搜索。

​	所接受的正则表达式语法与Perl、Python和其他语言使用的通用语法相同。更准确地说，它是RE2接受的语法，并在https://golang.org/s/re2syntax上描述，除了`\C`。要了解语法的概述，请运行

```
go doc regexp/syntax
```

​	此包提供的regexp实现保证以输入大小为线性时间运行。(这是大多数开源正则表达式实现不保证的属性。)有关此属性的更多信息，请参见

```
https://swtch.com/~rsc/regexp/regexp1.html
```

或有关自动机理论的任何书籍。

​	所有字符都是UTF-8编码的码点。在utf8.DecodeRune之后，无效的UTF-8序列的每个字节都被视为编码utf8.RuneError(U+FFFD)。

​	Regexp有16个方法可以匹配正则表达式并标识匹配的文本。它们的名称由此正则表达式匹配：

```
Find(All)?(String)?(Submatch)?(Index)?
```

​	如果存在'All'，则该例程将匹配整个表达式的连续不重叠的匹配项。忽略与前面匹配的空匹配项。返回值是一个包含相应非'All'例程的连续返回值的切片。这些例程需要一个额外的整数参数n。如果n >= 0，则函数返回最多n个匹配项/子匹配项；否则，它返回它们所有。

​	如果存在'String'，则参数是一个字符串；否则它是字节切片；返回值适当调整。

​	如果存在'Submatch'，则返回值是一个标识表达式的连续子匹配项的切片。子匹配项是正则表达式中括号表达式(也称为捕获组)的匹配项，按从左到右顺序编号。子匹配项0是整个表达式的匹配项，子匹配项1是第一个括号表达式的匹配项，依此类推。

​	如果存在'Index'，则匹配项和子匹配项通过输入字符串中的字节索引对进行标识：`result[2*n：2*n+2]`标识第n个子匹配项的索引。n==0的对标识整个表达式的匹配项。如果不存在'Index'，则匹配项由匹配/子匹配的文本标识。如果索引为负或文本为nil，则表示该子表达式未匹配输入中的任何字符串。对于'String'版本，空字符串表示没有匹配或空匹配。

​	还有一组可以应用于从RuneReader读取的文本的方法：

```
MatchReader, FindReaderIndex, FindReaderSubmatchIndex
```

​	这个集合可能会增长。注意，正则表达式匹配可能需要检查超过匹配返回的文本，因此从RuneReader匹配文本的方法在返回之前可能会读取任意数量的输入。

(还有一些不符合此模式的方法。)

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Compile the expression once, usually at init time.
	// Use raw strings to avoid having to quote the backslashes.
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
}
Output:

true
true
false
false
```

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func Match 

``` go 
func Match(pattern string, b []byte) (matched bool, err error)
```

​	Match函数报告字节切片b是否包含正则表达式模式的任何匹配项。更复杂的查询需要使用Compile和完整的Regexp接口。

##### Match Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`bar.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`a(b`, []byte(`seafood`))
	fmt.Println(matched, err)

}
Output:

true <nil>
false <nil>
false error parsing regexp: missing closing ): `a(b`
```



#### func MatchReader 

``` go 
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
```

​	MatchReader函数报告RuneReader返回的文本是否包含正则表达式模式的任何匹配项。更复杂的查询需要使用Compile和完整的Regexp接口。

#### func MatchString 

``` go 
func MatchString(pattern string, s string) (matched bool, err error)
```

​	MatchString函数报告字符串s是否包含正则表达式模式的任何匹配项。更复杂的查询需要使用Compile和完整的Regexp接口。

##### MatchString Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.MatchString(`foo.*`, "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(`bar.*`, "seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(`a(b`, "seafood")
	fmt.Println(matched, err)
}
Output:

true <nil>
false <nil>
false error parsing regexp: missing closing ): `a(b`
```



#### func QuoteMeta 

``` go 
func QuoteMeta(s string) string
```

​	QuoteMeta函数返回一个字符串，其中转义了参数文本中的所有正则表达式元字符；返回的字符串是与字面文本匹配的正则表达式。

##### QuoteMeta Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`))
}

```



## 类型

### type Regexp 

``` go 
type Regexp struct {
	// contains filtered or unexported fields
    // 包含已过滤或未导出的字段
}
```

​	Regexp 是已编译的正则表达式的表示。Regexp 可以被多个 goroutine 安全地并发使用，除了 Longest 等配置方法之外。

#### func Compile 

``` go 
func Compile(expr string) (*Regexp, error)
```

​	Compile函数解析一个正则表达式，如果成功则返回 Regexp 对象，该对象可以用于对文本进行匹配。

​	在匹配文本时，正则表达式返回一个尽可能早地开始的匹配(最左边的)，并从其中选择一个回溯搜索最先发现的匹配。这种所谓的最左优先匹配是 Perl、Python 和其他实现使用的语义，尽管此包实现它而不需要回溯的开销。有关 POSIX 最左最长匹配，请参见 CompilePOSIX。

#### func CompilePOSIX 

``` go 
func CompilePOSIX(expr string) (*Regexp, error)
```

​	CompilePOSIX函数类似于 Compile函数，但将正则表达式限制为 POSIX ERE(egrep)语法，并将匹配语义更改为最左最长匹配。

​	也就是说，在匹配文本时，正则表达式返回尽可能早地开始的匹配(最左边的)，并从其中选择尽可能长的匹配。这种所谓的最左最长匹配是早期正则表达式实现所使用的语义，也是 POSIX 规定的。

​	但是，可能会有多个最左最长匹配，其中具有不同子匹配选择，这里的此包与 POSIX 不同。在可能的最左最长匹配中，此包选择回溯搜索最先发现的匹配，而 POSIX 规定应选择使第一个子表达式的长度最大，然后是第二个表达式，依此类推，从左到右。POSIX 规则计算上过于繁琐，甚至无法定义。有关详细信息，请参见 https://swtch.com/~rsc/regexp/regexp2.html#posix。

#### func MustCompile 

``` go 
func MustCompile(str string) *Regexp
```

​	MustCompile函数类似于 Compile方法，但如果无法解析表达式，则会 panic。它简化了持有已编译的正则表达式的全局变量的安全初始化。

#### func MustCompilePOSIX 

``` go 
func MustCompilePOSIX(str string) *Regexp
```

​	MustCompilePOSIX函数类似于 CompilePOSIX，但如果无法解析表达式，则会 panic。它简化了持有已编译的正则表达式的全局变量的安全初始化。

#### (*Regexp) Expand 

``` go 
func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte
```

​	Expand方法将模板追加到dst并返回结果；在追加期间，它使用从src中获得的相应匹配项替换模板中的变量。match切片应该由FindSubmatchIndex返回。

​	在模板中，变量由形如`$name`或`${name}`的子字符串表示，其中name是字母、数字和下划线的非空序列。纯数字名称(例如`$1`)引用相应索引的子匹配项；其他名称引用使用`(?P<name>...)`语法命名的捕获括号。引用超出范围或未匹配索引或不存在于正则表达式中的名称将被替换为一个空切片。

​	在`$name`形式中，name被认为尽可能长：`$1x`等同于`${1x}`，而不是`${1}x`，`$10`等同于`${10}`，而不是`${1}0`。

​	要在输出中插入文字`$`，请在模板中使用`$$`。

##### Expand Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`)

	// 正则表达式模式从内容中捕获"key: value"对。
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// 将"key: value"转换为"key=value"的模板，
    // 通过引用正则表达式模式捕获的值。
	template := []byte("$key=$value\n")

	result := []byte{}

	// 对于内容中的每个正则表达式匹配。
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		// 将捕获的子匹配应用于模板，并将输出附加到结果中。
		result = pattern.Expand(result, template, content, submatches)
	}
	fmt.Println(string(result))
}
Output:

option1=value1
option2=value2
option3=value3
```





#### (*Regexp) ExpandString 

``` go 
func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte
```

​	ExpandString方法类似于Expand方法，但模板和源为字符串。它将追加到并返回字节片，以便调用代码控制分配。

##### ExpandString Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := `
	# comment line
	option1: value1
	option2: value2

	# another comment line
	option3: value3
`

	// 正则表达式模式从内容中捕获"key: value"对。
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// 使用模板将"key: value"转换为"key=value"，
    // 并参照正则表达式模式中捕获的值。
	template := "$key=$value\n"

	result := []byte{}

	// 对于内容中的每个正则表达式匹配。
	for _, submatches := range pattern.FindAllStringSubmatchIndex(content, -1) {
		// 将捕获的子匹配应用于模板，并将输出附加到结果中。
		result = pattern.ExpandString(result, template, content, submatches)
	}
	fmt.Println(string(result))
}
Output:

option1=value1
option2=value2
option3=value3
```



#### (*Regexp) Find 

``` go 
func (re *Regexp) Find(b []byte) []byte
```

​	Find方法返回一个切片，其中包含b中最左侧的正则表达式匹配项的文本。返回值为nil表示无匹配项。

##### Find Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.Find([]byte(`seafood fool`)))

}
Output:

"food"
```



#### (*Regexp) FindAll 

``` go 
func (re *Regexp) FindAll(b []byte, n int) [][]byte
```

​	FindAll方法是Find方法的"All"版本；它返回所有连续匹配表达式的切片，如包注释中的"All"描述所定义。返回值为nil表示无匹配项。

##### FindAll Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), -1))

}
Output:

["food" "fool"]
```



#### (*Regexp) FindAllIndex 

``` go 
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
```

​	FindAllIndex方法是FindIndex方法的"All"版本；它返回所有连续匹配表达式的索引切片，如包注释中的"All"描述所定义。返回值为nil表示无匹配项。

##### FindAllIndex Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := []byte("London")
	re := regexp.MustCompile(`o.`)
	fmt.Println(re.FindAllIndex(content, 1))
	fmt.Println(re.FindAllIndex(content, -1))
}
Output:

[[1 3]]
[[1 3] [4 6]]
```



#### (*Regexp) FindAllString 

``` go 
func (re *Regexp) FindAllString(s string, n int) []string
```

​	FindAllString方法是FindString方法的"All"版本；它返回一个切片，其中包含表达式的所有连续匹配项，如包注释中的"All"描述所定义。返回值nil表示没有匹配项。

##### FindAllString Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a.`)
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
}
Output:

[ar an al]
[ar an]
[aa]
[]
```



#### (*Regexp) FindAllStringIndex 

``` go 
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
```

​	FindAllStringIndex方法是FindStringIndex方法的"All"版本；它返回一个切片，其中包含表达式的所有连续匹配项，如包注释中的"All"描述所定义。返回值nil表示没有匹配项。

#### (*Regexp) FindAllStringSubmatch 

``` go 
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
```

​	FindAllStringSubmatch方法是FindStringSubmatch方法的"All"版本；它返回一个切片，其中包含表达式的所有连续匹配项，如包注释中的"All"描述所定义。返回值nil表示没有匹配项。

##### FindAllStringSubmatch Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))
}
Output:

[["ab" ""]]
[["axxb" "xx"]]
[["ab" ""] ["axb" "x"]]
[["axxb" "xx"] ["ab" ""]]
```



#### (*Regexp) FindAllStringSubmatchIndex 

``` go 
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
```

​	FindAllStringSubmatchIndex方法是FindStringSubmatchIndex方法的"All"版本；它返回一个切片，其中包含表达式的所有连续匹配项，如包注释中的"All"描述所定义。返回值nil表示没有匹配项。

##### FindAllStringSubmatchIndex Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(x*)b`)
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))
}
Output:

[[1 3 2 2]]
[[1 5 2 4]]
[[1 3 2 2] [4 7 5 6]]
[[1 5 2 4] [6 8 7 7]]
[]
```



#### (*Regexp) FindAllSubmatch 

``` go 
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
```

​	FindAllSubmatch方法是FindSubmatch方法的"All"版本；它返回一个切片，其中包含表达式的所有连续匹配项，如包注释中的"All"描述所定义。返回值nil表示没有匹配项。

##### FindAllSubmatch Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`), -1))

}
Output:

[["food" "d"] ["fool" "l"]]
```



#### (*Regexp) FindAllSubmatchIndex 

``` go 
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
```

​	FindAllSubmatchIndex方法是FindSubmatchIndex方法的"All"版本；它返回一个切片，其中包含表达式的所有连续匹配项，如包注释中的"All"描述所定义。返回值nil表示没有匹配项。

##### FindAllSubmatchIndex Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2
`)
	// 正则表达式模式从内容中捕获"key: value"对。
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	allIndexes := pattern.FindAllSubmatchIndex(content, -1)
	for _, loc := range allIndexes {
		fmt.Println(loc)
		fmt.Println(string(content[loc[0]:loc[1]]))
		fmt.Println(string(content[loc[2]:loc[3]]))
		fmt.Println(string(content[loc[4]:loc[5]]))
	}
}
Output:

[18 33 18 25 27 33]
option1: value1
option1
value1
[35 50 35 42 44 50]
option2: value2
option2
value2
```



#### (*Regexp) FindIndex 

``` go 
func (re *Regexp) FindIndex(b []byte) (loc []int)
```

​	FindIndex方法返回一个包含两个整数的切片，这个切片定义了b中与正则表达式最左匹配的位置。匹配本身位于`b[loc[0]:loc[1]]`。如果没有匹配，返回值为nil。

##### FindIndex Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	content := []byte(`
	# comment line
	option1: value1
	option2: value2
`)
	// 正则表达式模式从内容中捕获"key: value"对。
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	loc := pattern.FindIndex(content)
	fmt.Println(loc)
	fmt.Println(string(content[loc[0]:loc[1]]))
}
Output:

[18 33]
option1: value1
```



#### (*Regexp) FindReaderIndex 

``` go 
func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)
```

​	FindReaderIndex方法返回一个包含两个整数的切片，这个切片定义了从RuneReader读取的文本中与正则表达式最左匹配的位置。匹配文本在输入流中的字节偏移量为`loc[0]`至l`oc[1]-1`。如果没有匹配，返回值为nil。

#### (*Regexp) FindReaderSubmatchIndex 

``` go 
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int
```

​	FindReaderSubmatchIndex方法返回一个切片，该切片包含标识正则表达式左侧最匹配的文本的索引对，以及它的子表达式(如果有)的匹配，如包注释中的"Submatch"和"Index"描述所定义。如果没有匹配，返回值为nil。

#### (*Regexp) FindString 

``` go 
func (re *Regexp) FindString(s string) string
```

​	FindString方法返回一个字符串，该字符串包含了s中与正则表达式最左匹配的文本。如果没有匹配，则返回值为空字符串，但如果正则表达式成功地匹配了一个空字符串，返回值也将为空字符串。如果需要区分这些情况，请使用FindStringIndex或FindStringSubmatch。

##### FindString Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindString("seafood fool"))
	fmt.Printf("%q\n", re.FindString("meat"))
}
Output:

"food"
""
```



#### (*Regexp) FindStringIndex 

``` go 
func (re *Regexp) FindStringIndex(s string) (loc []int)
```

FindStringIndex方法返回一个包含两个整数的切片，这个切片定义了s中与正则表达式最左匹配的位置。匹配本身位于s[loc[0]:loc[1]]。如果没有匹配，返回值为nil。

##### FindStringIndex Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`ab?`)
	fmt.Println(re.FindStringIndex("tablett"))
	fmt.Println(re.FindStringIndex("foo") == nil)
}
Output:

[1 3]
true
```



#### (*Regexp) FindStringSubmatch 

``` go 
func (re *Regexp) FindStringSubmatch(s string) []string
```

FindStringSubmatch方法返回一个字符串切片，该切片包含s中正则表达式最左匹配的文本和其子表达式(如果有)的匹配，如包注释中的"Submatch"描述所定义。如果没有匹配，返回值为nil。

##### FindStringSubmatch Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
}
Output:

["axxxbyc" "xxx" "y"]
["abzc" "" "z"]
```



#### (*Regexp) FindStringSubmatchIndex 

``` go 
func (re *Regexp) FindStringSubmatchIndex(s string) []int
```

​	FindStringSubmatchIndex方法返回一个切片，其中包含识别字符串s中最左侧的正则表达式匹配项及其子表达式(如果有)的索引对，如"Submatch"和"Index"描述所定义。如果没有匹配项，则返回nil。

#### (*Regexp) FindSubmatch 

``` go 
func (re *Regexp) FindSubmatch(b []byte) [][]byte
```

​	FindSubmatch方法返回一个切片，其中包含识别字节切片b中最左侧的正则表达式匹配项及其子表达式(如果有)的文本，如包注释中"Submatch"描述所定义。如果没有匹配项，则返回nil。

##### FindSubmatch Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindSubmatch([]byte(`seafood fool`)))

}
Output:

["food" "d"]
```



#### (*Regexp) FindSubmatchIndex 

``` go 
func (re *Regexp) FindSubmatchIndex(b []byte) []int
```

​	FindSubmatchIndex方法返回一个切片，其中包含识别字节切片b中最左侧的正则表达式匹配项及其子表达式(如果有)的索引对，如"Submatch"和"Index"描述所定义。如果没有匹配项，则返回nil。

##### FindSubmatchIndex Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(x*)b`)
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
	fmt.Println(re.FindSubmatchIndex([]byte("-ab-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-axxb-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-ab-axb-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-axxb-ab-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-foo-")))
}
Output:

[1 3 2 2]
[1 5 2 4]
[1 3 2 2]
[1 5 2 4]
[]
```



#### (*Regexp) LiteralPrefix 

``` go 
func (re *Regexp) LiteralPrefix() (prefix string, complete bool)
```

​	LiteralPrefix方法返回一个字符串文本，该文本必须以正则表达式re的开头。如果文本字符串包含整个正则表达式，则返回布尔值true。

#### (*Regexp) Longest  <- go1.1

``` go 
func (re *Regexp) Longest()
```

​	Longest方法使未来的搜索更倾向于最左侧最长匹配。也就是说，在匹配文本时，正则表达式返回一个在输入中尽可能早的匹配项(最左侧)，并从中选择尽可能长的匹配项。此方法修改Regexp，并且可能不与任何其他方法同时调用。

##### Longest Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(|b)`)
	fmt.Println(re.FindString("ab"))
	re.Longest()
	fmt.Println(re.FindString("ab"))
}
Output:

a
ab
```



#### (*Regexp) Match 

``` go 
func (re *Regexp) Match(b []byte) bool
```

​	Match方法报告byte切片b是否包含正则表达式re的任何匹配项。

##### Match Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`foo.?`)
	fmt.Println(re.Match([]byte(`seafood fool`)))
	fmt.Println(re.Match([]byte(`something else`)))

}
Output:

true
false
```



#### (*Regexp) MatchReader 

``` go 
func (re *Regexp) MatchReader(r io.RuneReader) bool
```

​	MatchReader方法报告RuneReader返回的文本是否包含正则表达式re的任何匹配项。

#### (*Regexp) MatchString 

``` go 
func (re *Regexp) MatchString(s string) bool
```

​	MatchString方法报告字符串s是否包含正则表达式re的任何匹配项。

##### MatchString Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(gopher){2}`)
	fmt.Println(re.MatchString("gopher"))
	fmt.Println(re.MatchString("gophergopher"))
	fmt.Println(re.MatchString("gophergophergopher"))
}
Output:

false
true
true
```



#### (*Regexp) NumSubexp 

``` go 
func (re *Regexp) NumSubexp() int
```

​	NumSubexp方法返回此Regexp中圆括号子表达式的数量。

##### NumSubexp Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re0 := regexp.MustCompile(`a.`)
	fmt.Printf("%d\n", re0.NumSubexp())

	re := regexp.MustCompile(`(.*)((a)b)(.*)a`)
	fmt.Println(re.NumSubexp())
}
Output:

0
4
```



#### (*Regexp) ReplaceAll 

``` go 
func (re *Regexp) ReplaceAll(src, repl []byte) []byte
```

​	ReplaceAll方法返回src的副本，将Regexp的匹配项替换为替换文本repl。在repl内，`$`符号的解释与Expand方法中的解释相同，例如`$1`表示第一个子匹配的文本。

##### ReplaceAll Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("T")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))
}
Output:

-T-T-
--xx-
---
-W-xxW-
```



#### (*Regexp) ReplaceAllFunc 

``` go 
func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
```

​	ReplaceAllFunc方法返回src的副本，其中所有Regexp的匹配项都已被应用于匹配的字节切片的repl函数的返回值所替换。 repl返回的替换会直接被替换，而不使用Expand方法。

#### (*Regexp) ReplaceAllLiteral 

``` go 
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
```

​	ReplaceAllLiteral方法返回一个 src 的副本，将所有与 Regexp 匹配的内容替换为 replacement bytes repl。替换 repl 时，不使用 Expand。

#### (*Regexp) ReplaceAllLiteralString 

``` go 
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
```

​	ReplaceAllLiteralString方法返回一个 src 的副本，将所有与 Regexp 匹配的内容替换为 replacement string repl。替换 repl 时，不使用 Expand。

##### ReplaceAllLiteralString Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))
}
Output:

-T-T-
-$1-$1-
-${1}-${1}-
```





#### (*Regexp) ReplaceAllString 

``` go 
func (re *Regexp) ReplaceAllString(src, repl string) string
```

​	ReplaceAllString方法返回一个 src 的副本，将所有与 Regexp 匹配的内容替换为 替换字符串repl。替换 repl 时，`$` 符号的解释方式与 Expand 相同，因此例如 `$1` 表示第一个 submatch 的文本。

##### ReplaceAllString Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))
}
Output:

-T-T-
--xx-
---
-W-xxW-
```



#### (*Regexp) ReplaceAllStringFunc 

``` go 
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string
```

​	ReplaceAllStringFunc方法返回一个 src 的副本，将所有与 Regexp 匹配的内容替换为应用于匹配的子字符串的函数 repl 的返回值。替换 repl 时，不使用 Expand。

##### ReplaceAllStringFunc Example

``` go 
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	re := regexp.MustCompile(`[^aeiou]`)
	fmt.Println(re.ReplaceAllStringFunc("seafood fool", strings.ToUpper))
}
Output:

SeaFooD FooL
```



#### (*Regexp) Split  <- go1.1

``` go 
func (re *Regexp) Split(s string, n int) []string
```

​	Split方法将字符串 s 分割为由表达式分隔的子字符串，并返回这些表达式匹配之间的子字符串的切片。

​	此方法返回的切片由 s 的所有不包含在 FindAllString 返回的切片中的子字符串组成。如果调用不包含元字符的表达式，则等效于 strings.SplitN。

示例：

```
s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
// s: ["", "b", "b", "c", "cadaaae"]
```

计数决定要返回的子字符串数：

n > 0：最多 n 个子字符串；最后一个子字符串是未分割的剩余部分。

n == 0：结果为 nil(零个子字符串) 

n < 0：所有子字符串

##### Split Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	a := regexp.MustCompile(`a`)
	fmt.Println(a.Split("banana", -1))
	fmt.Println(a.Split("banana", 0))
	fmt.Println(a.Split("banana", 1))
	fmt.Println(a.Split("banana", 2))
	zp := regexp.MustCompile(`z+`)
	fmt.Println(zp.Split("pizza", -1))
	fmt.Println(zp.Split("pizza", 0))
	fmt.Println(zp.Split("pizza", 1))
	fmt.Println(zp.Split("pizza", 2))
}
Output:

[b n n ]
[]
[banana]
[b nana]
[pi a]
[]
[pizza]
[pi a]
```



#### (*Regexp) String 

``` go 
func (re *Regexp) String() string
```

​	String方法返回用于编译正则表达式的源文本。

#### (*Regexp) SubexpIndex  <- go1.15

``` go 
func (re *Regexp) SubexpIndex(name string) int
```

SubexpIndex方法返回第一个名称为 name 的子表达式的索引；如果没有具有该名称的子表达式，则返回 -1。

​	请注意，多个子表达式可以使用相同的名称编写，例如 `(?P<bob>a+)(?P<bob>b+)`，这声明了两个名为"bob"的子表达式。在这种情况下，SubexpIndex 返回正则表达式中最左边的子表达式的索引。

##### SubexpIndex  Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Println(re.MatchString("Alan Turing"))
	matches := re.FindStringSubmatch("Alan Turing")
	lastIndex := re.SubexpIndex("last")
	fmt.Printf("last => %d\n", lastIndex)
	fmt.Println(matches[lastIndex])
}
Output:

true
last => 2
Turing
```



#### (*Regexp) SubexpNames 

``` go 
func (re *Regexp) SubexpNames() []string
```

​	SubexpNames方法返回正则表达式re中有命名的捕获组的名称。第一个子表达式的名称为names[1]，因此如果m是一个匹配切片，m[i]的名称为SubexpNames()[i]。由于整个正则表达式不能被命名，因此names[0]始终为空字符串。切片不应被修改。

##### SubexpNames  Example

``` go 
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Println(re.MatchString("Alan Turing"))
	fmt.Printf("%q\n", re.SubexpNames())
	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed))
}
Output:

true
["" "first" "last"]
${last} ${first}
Turing Alan
```

