+++
title = "embed"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/embed@go1.21.3](https://pkg.go.dev/embed@go1.21.3)

Package embed provides access to files embedded in the running Go program.

​	`embed`包提供了访问嵌入在运行的 Go 程序中的文件的功能。

Go source files that import "embed" can use the //go:embed directive to initialize a variable of type string, []byte, or FS with the contents of files read from the package directory or subdirectories at compile time.

​	导入 "embed" 的 Go 源文件可以使用 `//go:embed` 指令，在编译时使用来自包目录或子目录中读取的文件的内容来初始化一个`string`类型、`[]byte` 类型或 `FS` 类型的变量。

For example, here are three ways to embed a file named hello.txt and then print its contents at run time.

​	例如，下面是三种嵌入名为 `hello.txt` 的文件并在运行时打印其内容的方式。

Embedding one file into a string:

​	将一个文件嵌入到一个`string`中：

```go
import _ "embed"

//go:embed hello.txt
var s string
print(s)
```

> 个人注释

{{< tabpane text=true >}}

{{< tab header="main.go" >}}

```go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var s string

func main() {
    fmt.Println(s) // 你好！	
}

```

{{< /tab >}}

{{< tab header="hello.txt" >}}

```txt
你好！
```

{{< /tab >}}

{{< /tabpane >}}

Embedding one file into a slice of bytes:

​	将一个文件嵌入到一个`[]byte`中：

```go
import _ "embed"

//go:embed hello.txt
var b []byte
print(string(b))
```

> 个人注释

{{< tabpane text=true >}}

{{< tab header="main.go" >}}

```go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var b []byte

func main() {
	fmt.Println(string(b)) // 你好！
	// go:embed hello.txt //报错： go:embed cannot apply to var inside func
	//var bb []byte
	//fmt.Println(string(bb))
}

```

{{< /tab >}}

{{< tab header="hello.txt" >}}

```txt
你好！
```

{{< /tab >}}

{{< /tabpane >}}

Embedded one or more files into a file system:

​	将一个或多个文件嵌入到一个文件系统（`FS`）中：

```go
import "embed"

//go:embed hello.txt
var f embed.FS
data, _ := f.ReadFile("hello.txt")
print(string(data))
```

> 个人注释

{{< tabpane text=true >}}

{{< tab header="main.go" >}}

```go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var f embed.FS

func main() {
    data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data)) // 你好！	
}

```

{{< /tab >}}

{{< tab header="hello.txt" >}}

```txt
你好！
```

{{< /tab >}}

{{< /tabpane >}}

## 指令 Directives

A //go:embed directive above a variable declaration specifies which files to embed, using one or more path.Match patterns.

​	变量声明上方的一个`//go:embed`指令使用一个或多个`path.Match`模式来指定要嵌入的文件。

The directive must immediately precede a line containing the declaration of a single variable. Only blank lines and ‘//’ line comments are permitted between the directive and the declaration.

​	该指令必须紧接在包含单个变量声明的行之前。在指令和声明之间只允许有空行和'`//`'行注释。

The type of the variable must be a string type, or a slice of a byte type, or FS (or an alias of FS).

​	变量的类型必须是`string` 类型，或`[]bype`类型，或`FS`(或FS的别名)。

For example:

​	例如：

``` go 
package server

import "embed"

// content holds our static web server content.
// 变量 content 存储我们的静态Web服务器内容。
//go:embed image/* template/*
//go:embed html/index.html
var content embed.FS
```

The Go build system will recognize the directives and arrange for the declared variable (in the example above, content) to be populated with the matching files from the file system.

​	Go构建系统（Go build system）将识别这些指令，并安排在文件系统中匹配的文件填充所声明的变量（在上面的例子中是 `content`）。

The //go:embed directive accepts multiple space-separated patterns for brevity, but it can also be repeated, to avoid very long lines when there are many patterns. The patterns are interpreted relative to the package directory containing the source file. The path separator is a forward slash, even on Windows systems. Patterns may not contain ‘.’ or ‘..’ or empty path elements, nor may they begin or end with a slash. To match everything in the current directory, use ‘*’ instead of ‘.’. To allow for naming files with spaces in their names, patterns can be written as Go double-quoted or back-quoted string literals.

​	`//go:embed`指令接受多个空格分隔的模式，以达到简洁表示，但也可以重复使用，以避免在有许多模式时出现很长的行。这些模式是相对于包含源文件的包目录解释的。路径分隔符是一个正斜杠，即使在Windows系统上也是如此。模式不能包含'`.`'或'`..`'或空的路径元素，也不能以斜杠开始或结束。要匹配当前目录中的所有内容，可以使用'`*`'代替'`.`'。为了允许命名文件时在其名称中加入空格，模式可以写成Go双引号（double-quoted）或反引号（back-quoted）的字符串字面量。

If a pattern names a directory, all files in the subtree rooted at that directory are embedded (recursively), except that files with names beginning with ‘.’ or ‘_’ are excluded. So the variable in the above example is almost equivalent to:

​	如果一个模式命名了一个目录，那么该目录下以及其子目录中的所有文件都会被嵌入（递归），但文件名以'`.`'或'`_`'开头的文件会被排除在外。因此，上述例子中的变量几乎等同于：

```go
// content 是我们的静态Web服务器内容。
//go:embed image template html/index.html
var content embed.FS
```

The difference is that ‘image/*’ embeds ‘image/.tempfile’ while ‘image’ does not. Neither embeds ‘image/dir/.tempfile’.

​	不同的是，'`image/*`'嵌入了'`image/.tempfile`'，而'`image`'却没有。两者都没有嵌入 '`image/dir/.tempfile`'。

If a pattern begins with the prefix ‘all:’, then the rule for walking directories is changed to include those files beginning with ‘.’ or ‘_’. For example, ‘all:image’ embeds both ‘image/.tempfile’ and ‘image/dir/.tempfile’.

​	如果一个模式以前缀 '`all:`' 开头，那么目录遍历的规则会改变，包括以 '`.`' 或 '`_`' 开头的文件。例如，'`all:image`' 嵌入了 '`image/.tempfile`' 和 '`image/dir/.tempfile`'。

The //go:embed directive can be used with both exported and unexported variables, depending on whether the package wants to make the data available to other packages. It can only be used with variables at package scope, not with local variables.

​	`//go:embed`指令可以用于导出的和未导出的变量，这取决于该包是否想让其他包获得这些数据。**它只能与包范围的变量一起使用，不能与局部变量一起使用**。

Patterns must not match files outside the package's module, such as ‘.git/*’ or symbolic links. Patterns must not match files whose names include the special punctuation characters " * < > ? ` ' | / \ and :. Matches for empty directories are ignored. After that, each pattern in a //go:embed line must match at least one file or non-empty directory.

​	模式不能匹配该包模块之外的文件，例如 '`.git/*`' 或符号链接。模式不能匹配文件名包含特殊标点符号字符 `"` `*` `<` `>` `?` \` `'` `|` `/` `\` 和 `:`。匹配空目录的模式会被忽略。除此之外，在 `//go:embed` 行中的每个模式必须至少匹配一个文件或非空目录。

If any patterns are invalid or have invalid matches, the build will fail.

​	如果有任何模式是无效的或具有无效匹配，那么构建将会失败。

## Strings 和 Bytes

The //go:embed line for a variable of type string or []byte can have only a single pattern, and that pattern can match only a single file. The string or []byte is initialized with the contents of that file.

​	`string` 或 `[]byte` 类型的变量的`//go:embed`行只能有一个模式，并且该模式只能匹配一个文件。`string` 或`[]byte`用该（匹配到的）文件的内容进行初始化。

The //go:embed directive requires importing "embed", even when using a string or []byte. In source files that don't refer to embed.FS, use a blank import (import _ "embed").

​	即使是使用`string`或 `[]byte`，`//go:embed` 指令也需要导入 "`embed`"。在不引用 `embed.FS` 的源文件中，使用一个空白导入（`import _ "embed"`）。

## File Systems

For embedding a single file, a variable of type string or []byte is often best. The FS type enables embedding a tree of files, such as a directory of static web server content, as in the example above.

​	对于嵌入单个文件，`string`  或 `[]byte`类型的变量通常是最好的的选择。FS 类型可以嵌入文件树，例如静态网络服务器内容的目录，就像上面的示例中那样。

FS implements the io/fs package's FS interface, so it can be used with any package that understands file systems, including net/http, text/template, and html/template.

​	FS 实现了 `io/fs` 包的 `FS` 接口，因此它可以与任何理解文件系统的包一起使用，包括 `net/http`、`text/template` 和 `html/template`。

For example, given the content variable in the example above, we can write:

​	例如，给定上面示例中的 `content` 变量，我们可以编写如下代码：

```go
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(content))))

template.ParseFS(content, "*.tmpl")
```

## Tools

To support tools that analyze Go packages, the patterns found in //go:embed lines are available in “go list” output. See the EmbedPatterns, TestEmbedPatterns, and XTestEmbedPatterns fields in the “go help list” output.

​	为了支持分析Go包的工具，在`//go:embed`行中发现的模式可以在 "`go list`"输出中使用。请参阅 "`go help list` "输出中的EmbedPatterns、TestEmbedPatterns和XTestEmbedPatterns字段。

## Example
``` go 
package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed internal/embedtest/testdata/*.txt
var content embed.FS

func main() {
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}
Output:

```

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type FS 

``` go 
type FS struct {
	// 编译器了解该结构体的布局。
    // 请参阅 cmd/compile/internal/staticdata 的 WriteEmbed。
    //
    // 文件列表按名称排序，但不是使用简单的字符串比较排序。
    // 相反，每个文件的名称采用 "dir/elem" 或 "dir/elem/" 的形式。
    // 可选的尾部斜杠表示该文件本身是一个目录。
    // 文件列表首先按照 dir 排序（如果 dir 不存在，则认为 dir 为 "."），
    // 然后按照 base 排序，因此以下文件列表：
	//
	//	p
	//	q/
	//	q/r
	//	q/s/
	//	q/s/t
	//	q/s/u
	//	q/v
	//	w
	//
	// 实际上按照以下顺序排序：
	//
	//	p       # dir=.    elem=p
	//	q/      # dir=.    elem=q
	//	w/      # dir=.    elem=w
	//	q/r     # dir=q    elem=r
	//	q/s/    # dir=q    elem=s
	//	q/v     # dir=q    elem=v
	//	q/s/t   # dir=q/s  elem=t
	//	q/s/u   # dir=q/s  elem=u
	//
	// This order brings directory contents together in contiguous sections
	// of the list, allowing a directory read to use binary search to find
	// the relevant sequence of entries.
    // 这种排序方式将目录内容连续地放置在列表的相邻部分，
	// 允许目录读取使用二分查找来查找相关的条目序列。
	files *[]file
}
```

An FS is a read-only collection of files, usually initialized with a //go:embed directive. When declared without a //go:embed directive, an FS is an empty file system.

​	`FS`结构体是一个只读的文件集合，通常用一个`//go:embed`指令来初始化。当声明时没有`//go:embed`指令，FS是一个空的文件系统。

An FS is a read-only value, so it is safe to use from multiple goroutines simultaneously and also safe to assign values of type FS to each other.

​	`FS` 是一个只读值，因此可以同时从多个 goroutine 中安全使用，也可以将类型为 FS 的值安全地赋值给其他（变量）。

FS implements fs.FS, so it can be used with any package that understands file system interfaces, including net/http, text/template, and html/template.

​	`FS` 实现了 `fs.FS` 接口，因此可以与任何理解文件系统接口的包一起使用，包括 `net/http`、`text/template` 和 `html/template`。

See the package documentation for more details about initializing an FS.

​	有关初始化 `FS` 的更多详细信息，请参阅该包文档。

#### (FS) Open 

``` go 
func (f FS) Open(name string) (fs.File, error)
```

Open opens the named file for reading and returns it as an fs.File.

​	`Open`方法打开指定的文件以进行读取，并将其作为 `fs.File` 返回。

The returned file implements io.Seeker when the file is not a directory.

​	当文件不是目录时，返回的文件实现了 `io.Seeker`。

#### (FS) ReadDir 

``` go 
func (f FS) ReadDir(name string) ([]fs.DirEntry, error)
```

ReadDir reads and returns the entire named directory.

​	`ReadDir`方法读取并返回整个命名的目录。??

​	`ReadDir`方法读取并返回指定目录的全部内容。??

#### (FS) ReadFile 

``` go 
func (f FS) ReadFile(name string) ([]byte, error)
```

ReadFile reads and returns the content of the named file.

​	`ReadFile`方法读取并返回指定文件的内容。