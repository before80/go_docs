+++
title = "path"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/path@go1.20.1

​	path包实现了用于操作斜杠分隔路径的实用程序函数。

​	path包仅用于由正斜杠分隔的路径，例如URL中的路径。该包不处理带有驱动器字母或反斜杠的Windows路径；要操作操作系统路径，请使用`path/filepath`包。


## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/path/match.go;l=14)

``` go 
var ErrBadPattern = errors.New("syntax error in pattern")
```

​	ErrBadPattern表示模式格式不正确。

## 函数

#### func Base 

``` go 
func Base(path string) string
```

​	Base函数返回path的最后一个元素。在提取最后一个元素之前，尾部的斜杠将被删除。如果路径为空，则Base返回"."。如果路径完全由斜杠组成，则Base返回"/"。

##### Base Example
``` go 
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
}
Output:

b
/
.
```

#### func Clean 

``` go 
func Clean(path string) string
```

​	Clean函数通过纯词法处理返回与path等效的最短路径名。它迭代地应用以下规则，直到不能再进行进一步的处理：

1. 将多个斜杠替换为单个斜杠。 
2. 消除每个". "路径名元素(当前目录)。 
3. 消除每个".."路径名元素(上一级目录)以及其前面的非".."元素。 
4. 消除以根路径开头的".."元素：即，在路径开头用"/"替换"/.."。

​	返回的路径仅在根"/"的情况下以斜杠结尾。

​	如果此过程的结果是空字符串，则Clean返回字符串"."。

​	请参见Rob Pike，"Plan 9中的词法文件名或正确处理点-点"，https://9p.io/sys/doc/lexnames.html

##### Clean Example
``` go 
package main

import (
	"fmt"
	"path"
)

func main() {
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
		"",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}

}
Output:

Clean("a/c") = "a/c"
Clean("a//c") = "a/c"
Clean("a/c/.") = "a/c"
Clean("a/c/b/..") = "a/c"
Clean("/../a/c") = "/a/c"
Clean("/../a/b/../././/c") = "/a/c"
Clean("") = "."
```

#### func Dir 

``` go 
func Dir(path string) string
```

​	Dir函数返回路径path中除去最后一个元素后的路径(即路径的目录部分)。使用Split去掉最后一个元素，然后将路径清理并删除尾部的斜杠。如果路径为空，则Dir返回"."。如果路径由一系列斜杠后跟非斜杠字符组成，则Dir返回单个斜杠。在任何其他情况下，返回的路径不以斜杠结尾。

##### Dir Example
``` go 
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("/a/b/c"))
	fmt.Println(path.Dir("a/b/c"))
	fmt.Println(path.Dir("/a/"))
	fmt.Println(path.Dir("a/"))
	fmt.Println(path.Dir("/"))
	fmt.Println(path.Dir(""))
}
Output:

/a/b
a/b
/a
a
/
.
```

#### func Ext 

``` go 
func Ext(path string) string
```

​	Ext函数返回路径path的文件扩展名。扩展名是路径的最终斜杠分隔元素中最后一个点号开始的后缀；如果没有点，则为空。

##### Ext Example
``` go 
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Ext("/a/b/c/bar.css"))
	fmt.Println(path.Ext("/"))
	fmt.Println(path.Ext(""))
}
Output:

.css
```

#### func IsAbs 

``` go 
func IsAbs(path string) bool
```

​	IsAbs函数报告路径path是否是绝对路径。

##### IsAbs Example
``` go 
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.IsAbs("/dev/null"))
}
Output:

true
```

#### func Join 

``` go 
func Join(elem ...string) string
```

​	Join函数将任意数量的路径元素连接成单个路径，使用斜杠进行分隔。忽略空元素。结果会被清理。但是，如果参数列表为空或其所有元素都为空，则Join返回一个空字符串。

##### Join Example
``` go 
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))

	fmt.Println(path.Join("a/b", "../../../xyz"))

	fmt.Println(path.Join("", ""))
	fmt.Println(path.Join("a", ""))
	fmt.Println(path.Join("", "a"))

}
Output:

a/b/c
a/b/c
a/b/c
../xyz

a
a
```

#### func Match 

``` go 
func Match(pattern, name string) (matched bool, err error)
```

​	Match函数报告name是否与shell模式匹配。模式语法是：

```
pattern:
	{ term }
term:
	'*'         匹配任何非/字符的序列
	'?'         匹配任何单个非/字符
	'[' [ '^' ] { character-range } ']'
	            字符类(必须非空)
	c           匹配字符c (c != '*', '?', '\\', '[')
	'\\' c      匹配字符c

character-range:
	c           匹配字符c (c != '\\', '-', ']')
	'\\' c      匹配字符c
	lo '-' hi   匹配区间[lo,hi]内的字符c(lo <= c <= hi)
```

​	Match函数要求pattern与name完全匹配，而不仅仅是子字符串。可能返回的唯一错误是ErrBadPattern，当pattern格式不正确时。

##### Match Example
``` go 
package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Match("abc", "abc"))
	fmt.Println(path.Match("a*", "abc"))
	fmt.Println(path.Match("a*/b", "a/c/b"))
}
Output:

true <nil>
true <nil>
false <nil>
```

#### func Split 

``` go 
func Split(path string) (dir, file string)
```

​	Split函数在最后一个斜杠后立即拆分path，将其分成目录和文件名两个组成部分。如果path中没有斜杠，则Split返回一个空目录和一个文件名设置为path的文件名。返回值的属性为path = dir + file。

##### Split Example
``` go 
package main

import (
	"fmt"
	"path"
)

func main() {
	split := func(s string) {
		dir, file := path.Split(s)
		fmt.Printf("path.Split(%q) = dir: %q, file: %q\n", s, dir, file)
	}
	split("static/myfile.css")
	split("myfile.css")
	split("")
}
Output:

path.Split("static/myfile.css") = dir: "static/", file: "myfile.css"
path.Split("myfile.css") = dir: "", file: "myfile.css"
path.Split("") = dir: "", file: ""
```

## 类型

This section is empty.