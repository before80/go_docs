+++
title = "embed"
linkTitle = "embed"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# embed

https://pkg.go.dev/embed@go1.20.1








Package embed provides access to files embedded in the running Go program.

包embed提供了对运行中的Go程序中嵌入的文件的访问。

Go source files that import "embed" can use the //go:embed directive to initialize a variable of type string, []byte, or FS with the contents of files read from the package directory or subdirectories at compile time.

导入 "embed "的 Go 源文件可以使用 `//go:embed` 指令，用编译时从包目录或子目录中读取的文件内容初始化字符串、[]byte或 FS 类型的变量。

For example, here are three ways to embed a file named hello.txt and then print its contents at run time.

例如，这里有三种方法来嵌入一个名为`hello.txt`的文件，然后在运行时打印其内容。

Embedding one file into a string:

将一个文件嵌入到一个字符串中：

```
import _ "embed"

//go:embed hello.txt
var s string
print(s)
```

Embedding one file into a slice of bytes:

将一个文件嵌入到一个字节片中：

```
import _ "embed"

//go:embed hello.txt
var b []byte
print(string(b))
```

Embedded one or more files into a file system:

将一个或多个文件嵌入到一个文件系统中：

```
import "embed"

//go:embed hello.txt
var f embed.FS
data, _ := f.ReadFile("hello.txt")
print(string(data))
```

#### Directives  指令

A //go:embed directive above a variable declaration specifies which files to embed, using one or more path.Match patterns.

变量声明上面的一个`//go:embed`指令使用一个或多个`path.Match`模式来指定要嵌入的文件。

The directive must immediately precede a line containing the declaration of a single variable. Only blank lines and ‘//’ line comments are permitted between the directive and the declaration.

该指令必须紧接在包含单个变量声明的行之前。在指令和声明之间只允许有空行和'//'行注释。

The type of the variable must be a string type, or a slice of a byte type, or FS (or an alias of FS).

变量的类型必须是字符串类型，或字节类型的片断，或FS(或FS的别名)。

For example:

例如：

``` go 
package server

import "embed"

// content holds our static web server content.
// content保存我们的静态Web服务器内容。
//go:embed image/* template/*
//go:embed html/index.html
var content embed.FS
```

The Go build system will recognize the directives and arrange for the declared variable (in the example above, content) to be populated with the matching files from the file system.

Go构建系统将识别这些指令，并安排声明的变量(在上面的例子中是content)与文件系统中的匹配文件一起填充。

The //go:embed directive accepts multiple space-separated patterns for brevity, but it can also be repeated, to avoid very long lines when there are many patterns. The patterns are interpreted relative to the package directory containing the source file. The path separator is a forward slash, even on Windows systems. Patterns may not contain ‘.’ or ‘..’ or empty path elements, nor may they begin or end with a slash. To match everything in the current directory, use ‘*’ instead of ‘.’. To allow for naming files with spaces in their names, patterns can be written as Go double-quoted or back-quoted string literals.

`//go:embed`指令接受多个空格分隔的模式，以达到简洁的目的，但也可以重复使用，以避免在有许多模式时出现很长的行。这些模式是相对于包含源文件的包目录解释的。路径分隔符是一个正斜杠，即使在Windows系统上也是如此。模式不能包含'.'或'.'或空的路径元素，也不能以斜线开始或结束。要匹配当前目录中的所有内容，使用'`*`'而不是'`.`'。为了允许命名文件时在其名称中加入空格，模式可以写成Go双引号或反引号的字符串字面。

If a pattern names a directory, all files in the subtree rooted at that directory are embedded (recursively), except that files with names beginning with ‘.’ or ‘_’ are excluded. So the variable in the above example is almost equivalent to:

如果一个模式命名了一个目录，那么根植于该目录的子树中的所有文件都会被嵌入(递归)，但名字以'.'或'_'开头的文件会被排除。因此，上述例子中的变量几乎等同于：

```
// content is our static web server content.
// content是我们的静态Web服务器内容。
//go:embed image template html/index.html
var content embed.FS
```

The difference is that ‘image/*’ embeds ‘image/.tempfile’ while ‘image’ does not. Neither embeds ‘image/dir/.tempfile’.

不同的是，'image/*'嵌入了'image/.tempfile'，而'image'却没有。也不嵌入'image/dir/.tempfile'。

If a pattern begins with the prefix ‘all:’, then the rule for walking directories is changed to include those files beginning with ‘.’ or ‘_’. For example, ‘all:image’ embeds both ‘image/.tempfile’ and ‘image/dir/.tempfile’.

如果一个模式以前缀'all:'开始，那么行走目录的规则就会改变，包括那些以'.'或'_'开始的文件。例如，'all:image'同时嵌入'image/.tempfile'和'image/dir/.tempfile'。

The //go:embed directive can be used with both exported and unexported variables, depending on whether the package wants to make the data available to other packages. It can only be used with variables at package scope, not with local variables.

`//go:embed`指令可以用于导出的和未导出的变量，这取决于包是否想让其他包获得这些数据。它只能用于包范围内的变量，不能用于本地变量。

Patterns must not match files outside the package's module, such as ‘.git/*’ or symbolic links. Patterns must not match files whose names include the special punctuation characters " * < > ? ` ' | / \ and :. Matches for empty directories are ignored. After that, each pattern in a //go:embed line must match at least one file or non-empty directory.

模式不能与包的模块之外的文件相匹配，如'.git/*'或符号链接。模式不能与名称中包含特殊标点符号的文件相匹配 " * < > ? ` ' | / 和 :。空目录的匹配会被忽略。之后，//go:embed行中的每个模式必须至少匹配一个文件或非空目录。

If any patterns are invalid or have invalid matches, the build will fail.

如果有任何模式是无效的或有无效的匹配，构建将失败。

#### Strings and Bytes 字符串和字节

The //go:embed line for a variable of type string or []byte can have only a single pattern, and that pattern can match only a single file. The string or []byte is initialized with the contents of that file.

字符串或[]byte 类型的变量的`//go:embed`行只能有一个模式，并且该模式只能匹配一个文件。字符串或[]字节用该文件的内容进行初始化。

The //go:embed directive requires importing "embed", even when using a string or []byte. In source files that don't refer to embed.FS, use a blank import (import _ "embed").

`//go:embed`指令要求导入 "embed"，即使是在使用字符串或[]字节的时候。在没有引用embed.FS的源文件中，使用一个空白的导入(import _ "embed")。

#### File Systems 文件系统

For embedding a single file, a variable of type string or []byte is often best. The FS type enables embedding a tree of files, such as a directory of static web server content, as in the example above.

对于嵌入单个文件，字符串或[]字节类型的变量通常是最好的。FS类型可以嵌入一棵树状的文件，例如静态网络服务器内容的目录，如上面的例子。

FS implements the io/fs package's FS interface, so it can be used with any package that understands file systems, including net/http, text/template, and html/template.

FS 实现了 io/fs 包的 FS 接口，所以它可以与任何理解文件系统的包一起使用，包括 net/http、text/template 和 html/template。

For example, given the content variable in the example above, we can write:

例如，给定上面例子中的内容变量，我们可以这样写：

```
http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(content))))

template.ParseFS(content, "*.tmpl")
```

#### Tools 工具

To support tools that analyze Go packages, the patterns found in //go:embed lines are available in "go list" output. See the EmbedPatterns, TestEmbedPatterns, and XTestEmbedPatterns fields in the "go help list" output.

为了支持分析Go包的工具，在//go:embed行中发现的模式可以在 "go list "输出中使用。参见 "go help list "输出中的EmbedPatterns、TestEmbedPatterns和XTestEmbedPatterns字段。

##### Example
``` go 
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
	// contains filtered or unexported fields
}
```

An FS is a read-only collection of files, usually initialized with a //go:embed directive. When declared without a //go:embed directive, an FS is an empty file system.

FS是一个只读的文件集合，通常用一个//go:embed指令来初始化。当声明时没有//go:embed指令，FS是一个空的文件系统。

An FS is a read-only value, so it is safe to use from multiple goroutines simultaneously and also safe to assign values of type FS to each other.

一个FS是一个只读的值，所以它可以从多个goroutine同时使用，也可以安全地将FS类型的值分配给对方。

FS implements fs.FS, so it can be used with any package that understands file system interfaces, including net/http, text/template, and html/template.

FS实现了fs.FS，所以它可以和任何理解文件系统接口的包一起使用，包括net/http、text/template和html/template。

See the package documentation for more details about initializing an FS.

关于初始化一个FS的更多细节，请看包的文档。

#### (FS) Open 

``` go 
func (f FS) Open(name string) (fs.File, error)
```

Open opens the named file for reading and returns it as an fs.File.

打开指定的文件供阅读，并将其作为fs.File返回。

The returned file implements io.Seeker when the file is not a directory.

当文件不是一个目录时，返回的文件实现了io.Seeker。

#### (FS) ReadDir 

``` go 
func (f FS) ReadDir(name string) ([]fs.DirEntry, error)
```

ReadDir reads and returns the entire named directory.

ReadDir读取并返回整个命名的目录。

#### (FS) ReadFile 

``` go 
func (f FS) ReadFile(name string) ([]byte, error)
```

ReadFile reads and returns the content of the named file.

ReadFile读取并返回指定文件的内容。