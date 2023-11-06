+++
title = "filepath"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/path/filepath@go1.21.3

Package filepath implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined file paths.

​	`filepath` 包实现了操作文件名路径的实用程序，以与目标操作系统定义的文件路径兼容的方式。

The filepath package uses either forward slashes or backslashes, depending on the operating system. To process paths such as URLs that always use forward slashes regardless of the operating system, see the path package.

​	filepath 包根据操作系统使用正斜杠或反斜杠。要处理始终使用正斜杠(无论操作系统如何)的路径(如 URL)，请参见 path 包。


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/path/filepath/path.go;l=62)

``` go 
const (
	Separator     = os.PathSeparator
	ListSeparator = os.PathListSeparator
)
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/path/filepath/match.go;l=17)

``` go 
var ErrBadPattern = errors.New("syntax error in pattern")
```

ErrBadPattern indicates a pattern was malformed.

​	ErrBadPattern 指示模式格式错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/path/filepath/path.go;l=399)

``` go 
var SkipAll error = fs.SkipAll
```

SkipAll is used as a return value from WalkFuncs to indicate that all remaining files and directories are to be skipped. It is not returned as an error by any function.

​	SkipAll 用作 WalkFuncs 的返回值，以指示所有剩余的文件和目录都将被跳过。它不作为任何函数的错误返回。 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/path/filepath/path.go;l=394)

``` go 
var SkipDir error = fs.SkipDir
```

SkipDir is used as a return value from WalkFuncs to indicate that the directory named in the call is to be skipped. It is not returned as an error by any function.

​	SkipDir 用作 WalkFuncs 的返回值，以指示调用中命名的目录将被跳过。它不作为任何函数的错误返回。



## 函数

### func Abs 

``` go 
func Abs(path string) (string, error)
```

Abs returns an absolute representation of path. If the path is not absolute it will be joined with the current working directory to turn it into an absolute path. The absolute path name for a given file is not guaranteed to be unique. Abs calls Clean on the result.

​	Abs函数返回 path 的绝对路径表示形式。如果 path 不是绝对路径，它将与当前工作目录连接以将其转换为绝对路径。给定文件的绝对路径名不能保证唯一。Abs函数对结果调用 Clean函数。

### func Base 

``` go 
func Base(path string) string
```

Base returns the last element of path. Trailing path separators are removed before extracting the last element. If the path is empty, Base returns ".". If the path consists entirely of separators, Base returns a single separator.

​	Base函数返回路径的最后一个元素。 在提取最后一个元素之前，末尾的路径分隔符会被删除。如果路径为空，则Base函数返回"。"。如果路径完全由分隔符组成，则Base函数返回单个分隔符。

Base Example

``` go 
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Base("/foo/bar/baz.js"))
	fmt.Println(filepath.Base("/foo/bar/baz"))
	fmt.Println(filepath.Base("/foo/bar/baz/"))
	fmt.Println(filepath.Base("dev.txt"))
	fmt.Println(filepath.Base("../todo.txt"))
	fmt.Println(filepath.Base(".."))
	fmt.Println(filepath.Base("."))
	fmt.Println(filepath.Base("/"))
	fmt.Println(filepath.Base(""))

}
Output:

On Unix:
baz.js
baz
baz
dev.txt
todo.txt
..
.
/
.
```

### func Clean 

``` go 
func Clean(path string) string
```

Clean returns the shortest path name equivalent to path by purely lexical processing. It applies the following rules iteratively until no further processing can be done:

​	Clean函数通过纯词汇处理返回等效于路径的最短路径名。它迭代应用以下规则，直到不能再进行处理：

1. Replace multiple Separator elements with a single one.

2. 将多个Separator元素替换为一个。 

3. Eliminate each . path name element (the current directory).

4. 消除每个"."路径名称元素(当前目录)。 

5. Eliminate each inner .. path name element (the parent directory) along with the non-.. element that precedes it.

6. 消除每个内部的".."路径名称元素(父目录)以及其前面的非".."元素。 

7. Eliminate .. elements that begin a rooted path: that is, replace "/.." by "/" at the beginning of a path, assuming Separator is '/'.

8. 消除以根路径开头的".."元素：即在路径开头将"/.."替换为"/"，假设Separator为"/"。

   

The returned path ends in a slash only if it represents a root directory, such as "/" on Unix or `C:\` on Windows.

​	返回的路径仅以斜杠结尾，如果它代表根目录(例如在Unix上的"/"或Windows上的"C:\")。

Finally, any occurrences of slash are replaced by Separator.

​	最后，将任何出现的斜杠替换为Separator。

If the result of this process is an empty string, Clean returns the string ".".

​	如果此过程的结果为空字符串，则Clean返回字符串"."。

See also Rob Pike, “Lexical File Names in Plan 9 or Getting Dot-Dot Right,” https://9p.io/sys/doc/lexnames.html

​	另请参见Rob Pike，"Plan 9中的词汇文件名或获取点点正确"，https://9p.io/sys/doc/lexnames.html

### func Dir 

``` go 
func Dir(path string) string
```

Dir returns all but the last element of path, typically the path's directory. After dropping the final element, Dir calls Clean on the path and trailing slashes are removed. If the path is empty, Dir returns ".". If the path consists entirely of separators, Dir returns a single separator. The returned path does not end in a separator unless it is the root directory.

​	Dir函数返回除路径的最后一个元素(通常是路径的目录)之外的所有内容。删除最后一个元素后，Dir函数在路径上调用Clean函数，尾随斜杠被删除。如果路径为空，则Dir返回"。"。如果路径完全由分隔符组成，则Dir返回单个分隔符。返回的路径不以分隔符结尾，除非它是根目录。

#### Dir Example
``` go 
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Dir("/foo/bar/baz.js"))
	fmt.Println(filepath.Dir("/foo/bar/baz"))
	fmt.Println(filepath.Dir("/foo/bar/baz/"))
	fmt.Println(filepath.Dir("/dirty//path///"))
	fmt.Println(filepath.Dir("dev.txt"))
	fmt.Println(filepath.Dir("../todo.txt"))
	fmt.Println(filepath.Dir(".."))
	fmt.Println(filepath.Dir("."))
	fmt.Println(filepath.Dir("/"))
	fmt.Println(filepath.Dir(""))

}
Output:

On Unix:
/foo/bar
/foo/bar
/foo/bar/baz
/dirty/path
.
..
.
.
/
.
```

### func EvalSymlinks 

``` go 
func EvalSymlinks(path string) (string, error)
```

EvalSymlinks returns the path name after the evaluation of any symbolic links. If path is relative the result will be relative to the current directory, unless one of the components is an absolute symbolic link. EvalSymlinks calls Clean on the result.

​	EvalSymlinks函数返回任何符号链接评估后的路径名。如果路径是相对路径，则结果将相对于当前目录，除非其中一个组件是绝对符号链接。EvalSymlinks函数在结果上调用Clean函数。

### func Ext 

``` go 
func Ext(path string) string
```

Ext returns the file name extension used by path. The extension is the suffix beginning at the final dot in the final element of path; it is empty if there is no dot.

​	Ext函数返回 path 使用的文件名扩展名。扩展名是从 path 的最后一个元素的最后一个句点开始的后缀；如果没有句点，则为空。

#### Ext Example
``` go 
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Printf("No dots: %q\n", filepath.Ext("index"))
	fmt.Printf("One dot: %q\n", filepath.Ext("index.js"))
	fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js"))
}
Output:

No dots: ""
One dot: ".js"
Two dots: ".js"
```

### func FromSlash 

``` go 
func FromSlash(path string) string
```

FromSlash returns the result of replacing each slash ('/') character in path with a separator character. Multiple slashes are replaced by multiple separators.

​	FromSlash函数返回将 path 中的每个斜杠('/')字符替换为分隔符字符的结果。多个斜杠将被多个分隔符替换。

### func Glob 

``` go 
func Glob(pattern string) (matches []string, err error)
```

Glob returns the names of all files matching pattern or nil if there is no matching file. The syntax of patterns is the same as in Match. The pattern may describe hierarchical names such as /usr/*/bin/ed (assuming the Separator is '/').

​	Glob函数返回与 pattern 匹配的所有文件的名称，如果没有匹配的文件，则返回 nil。模式的语法与 Match 中相同。该模式可以描述分层名称，例如 `/usr/*/bin/ed`(假设 Separator 是'/')。

Glob ignores file system errors such as I/O errors reading directories. The only possible returned error is ErrBadPattern, when pattern is malformed.

​	Glob函数忽略文件系统错误，例如读取目录时的 I/O 错误。当模式格式错误时，仅可能返回 ErrBadPattern。

### func HasPrefix <-DEPRECATED

```
func HasPrefix(p, prefix string) bool
```

HasPrefix exists for historical compatibility and should not be used.

Deprecated: HasPrefix does not respect path boundaries and does not ignore case when required.

### func IsAbs 

``` go 
func IsAbs(path string) bool
```

IsAbs reports whether the path is absolute.

​	IsAbs函数报告路径是否为绝对路径。

#### IsAbs Example
``` go 
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.IsAbs("/home/gopher"))
	fmt.Println(filepath.IsAbs(".bashrc"))
	fmt.Println(filepath.IsAbs(".."))
	fmt.Println(filepath.IsAbs("."))
	fmt.Println(filepath.IsAbs("/"))
	fmt.Println(filepath.IsAbs(""))

}
Output:

On Unix:
true
false
false
false
true
false
```

### func IsLocal  <- go1.20

``` go 
func IsLocal(path string) bool
```

IsLocal reports whether path, using lexical analysis only, has all of these properties:

​	IsLocal函数仅使用词法分析，报告 path 是否具有以下所有属性：

- is within the subtree rooted at the directory in which path is evaluated
- 在 path 评估的目录为根的子树中 
- is not an absolute path
- 不是绝对路径 
- is not empty
- 不为空 
- on Windows, is not a reserved name such as "NUL"
- 在 Windows 上，不是保留名称，例如"NUL"

If IsLocal(path) returns true, then Join(base, path) will always produce a path contained within base and Clean(path) will always produce an unrooted path with no ".." path elements.

​	如果 IsLocal(path) 返回 true，则 Join(base, path) 总是生成包含在 base 中的路径，并且 Clean(path) 总是生成没有 ".." 路径元素的未根路径。

IsLocal is a purely lexical operation. In particular, it does not account for the effect of any symbolic links that may exist in the filesystem.

​	IsLocal函数是纯词法操作。特别地，它不考虑文件系统中可能存在的任何符号链接的影响。

### func Join 

``` go 
func Join(elem ...string) string
```

Join joins any number of path elements into a single path, separating them with an OS specific Separator. Empty elements are ignored. The result is Cleaned. However, if the argument list is empty or all its elements are empty, Join returns an empty string. On Windows, the result will only be a UNC path if the first non-empty element is a UNC path.

​	Join函数将任意数量的路径元素连接为单个路径，并使用特定于操作系统的分隔符进行分隔。空元素将被忽略。结果将经过清理。但是，如果参数列表为空或其所有元素为空，则 Join函数返回一个空字符串。在 Windows 上，如果第一个非空元素是 UNC 路径，则结果将只是一个 UNC 路径。

#### Join Example
``` go 
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))

	fmt.Println(filepath.Join("a/b", "../../../xyz"))

}
Output:

On Unix:
a/b/c
a/b/c
a/b/c
a/b/c
../xyz
```

### func Match 

``` go 
func Match(pattern, name string) (matched bool, err error)
```

Match reports whether name matches the shell file name pattern. The pattern syntax is:

​	Match 函数报告 name 是否与 shell 文件名模式匹配。模式语法如下：

```
pattern:
	{ term }
term:
	'*'         匹配任意非分隔符字符的序列
	'?'         匹配任意单个非分隔符字符
	'[' [ '^' ] { character-range } ']'
	            字符类(必须非空)
	c           匹配字符 c (c != '*', '?', '\\', '[')
	'\\' c      匹配字符 c

character-range:
	c           匹配字符 c (c != '\\', '-', ']')
	'\\' c      匹配字符 c
	lo '-' hi   匹配 lo <= c <= hi 的字符 c
```

Match requires pattern to match all of name, not just a substring. The only possible returned error is ErrBadPattern, when pattern is malformed.

​	Match函数要求模式与 name 完全匹配，而不仅仅是子字符串。可能的返回错误仅为 ErrBadPattern，当模式存在格式错误时。

On Windows, escaping is disabled. Instead, '\\' is treated as path separator.

​	在 Windows 平台上，转义字符被禁用。相反，'\' 被视为路径分隔符。

#### Match Example

``` go 
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo"))
	fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo/bar"))
	fmt.Println(filepath.Match("/home/?opher", "/home/gopher"))
	fmt.Println(filepath.Match("/home/\\*", "/home/*"))

}
Output:

On Unix:
true <nil>
false <nil>
true <nil>
true <nil>
```

### func Rel 

``` go 
func Rel(basepath, targpath string) (string, error)
```

Rel returns a relative path that is lexically equivalent to targpath when joined to basepath with an intervening separator. That is, Join(basepath, Rel(basepath, targpath)) is equivalent to targpath itself. On success, the returned path will always be relative to basepath, even if basepath and targpath share no elements. An error is returned if targpath can't be made relative to basepath or if knowing the current working directory would be necessary to compute it. Rel calls Clean on the result.

​	Rel函数返回一个相对路径，该路径在插入分隔符后连接到 basepath 上时与 targpath 在词法上等效。也就是说，Join(basepath, Rel(basepath, targpath)) 等价于 targpath。成功时，返回的路径将始终相对于 basepath，即使 basepath 和 targpath 没有共享元素。如果无法将 targpath 变成相对于 basepath 的路径，或者需要知道当前工作目录才能计算它，则会返回错误。Rel函数对结果调用 Clean函数。

#### Rel Example
``` go 
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	paths := []string{
		"/a/b/c",
		"/b/c",
		"./b/c",
	}
	base := "/a"

	fmt.Println("On Unix:")
	for _, p := range paths {
		rel, err := filepath.Rel(base, p)
		fmt.Printf("%q: %q %v\n", p, rel, err)
	}

}
Output:

On Unix:
"/a/b/c": "b/c" <nil>
"/b/c": "../b/c" <nil>
"./b/c": "" Rel: can't make ./b/c relative to /a
```

### func Split 

``` go 
func Split(path string) (dir, file string)
```

Split splits path immediately following the final Separator, separating it into a directory and file name component. If there is no Separator in path, Split returns an empty dir and file set to path. The returned values have the property that path = dir+file.

​	Split函数在最后一个分隔符之后立即分割路径，将其分割为目录和文件名组件。如果 path 中没有分隔符，则 Split 返回一个 dir 和 file 均为 path 的空字符串。返回的值具有 path = dir+file 的属性。

#### Split Example
``` go 
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	paths := []string{
		"/home/arnie/amelia.jpg",
		"/mnt/photos/",
		"rabbit.jpg",
		"/usr/local//go",
	}
	fmt.Println("On Unix:")
	for _, p := range paths {
		dir, file := filepath.Split(p)
		fmt.Printf("input: %q\n\tdir: %q\n\tfile: %q\n", p, dir, file)
	}
}
Output:

On Unix:
input: "/home/arnie/amelia.jpg"
	dir: "/home/arnie/"
	file: "amelia.jpg"
input: "/mnt/photos/"
	dir: "/mnt/photos/"
	file: ""
input: "rabbit.jpg"
	dir: ""
	file: "rabbit.jpg"
input: "/usr/local//go"
	dir: "/usr/local//"
	file: "go"
```

### func SplitList 

``` go 
func SplitList(path string) []string
```

SplitList splits a list of paths joined by the OS-specific ListSeparator, usually found in PATH or GOPATH environment variables. Unlike strings.Split, SplitList returns an empty slice when passed an empty string.

​	`SplitList`函数将由 OS 特定 ListSeparator 连接的路径列表拆分为多个路径，通常在 PATH 或 GOPATH 环境变量中找到。与 strings.Split 不同，当传入空字符串时，SplitList函数返回一个空切片。

#### SplitList Example
``` go 
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin"))
}
Output:

On Unix: [/a/b/c /usr/bin]
```

### func ToSlash 

``` go 
func ToSlash(path string) string
```

ToSlash returns the result of replacing each separator character in path with a slash ('/') character. Multiple separators are replaced by multiple slashes.

​	`ToSlash`函数返回将路径中的每个分隔符字符替换为斜杠('/')字符的结果。多个分隔符被多个斜杠替换。

### func VolumeName 

``` go 
func VolumeName(path string) string
```

VolumeName returns leading volume name. Given "C:\foo\bar" it returns "C:" on Windows. Given "\\host\share\foo" it returns "\\host\share". On other platforms it returns "".

​	VolumeName函数返回 Windows 系统下的卷名，例如 "C:\foo\bar" 将返回 "C:" ，"\host\share\foo" 将返回 "\host\share"。在其他平台上，返回空字符串。

### func Walk 

``` go 
func Walk(root string, fn WalkFunc) error
```

Walk walks the file tree rooted at root, calling fn for each file or directory in the tree, including root.

​	`Walk` 函数遍历以 root 为根的文件树，对树中的每个文件或目录都会调用 fn 函数，包括根目录。

All errors that arise visiting files and directories are filtered by fn: see the WalkFunc documentation for details.

​	所有访问文件和目录时产生的错误都会通过 fn 进行过滤：有关详细信息，请参见 WalkFunc 的文档。

The files are walked in lexical order, which makes the output deterministic but requires Walk to read an entire directory into memory before proceeding to walk that directory.

​	文件按词法顺序遍历，这使得输出是确定性的，但要求 Walk 读取整个目录到内存中，然后才能继续遍历该目录。

Walk does not follow symbolic links.

​	Walk 函数不会遵循符号链接。

Walk is less efficient than WalkDir, introduced in Go 1.16, which avoids calling os.Lstat on every visited file or directory.

​	Walk 函数不如在 Go 1.16 中引入的 WalkDir 函数高效，后者避免了在访问每个文件或目录时调用 os.Lstat 的问题。

#### Walk Example
``` go 
//go:build !windows && !plan9

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func prepareTestDirTree(tree string) (string, error) {
	tmpDir, err := os.MkdirTemp("", "")
	if err != nil {
		return "", fmt.Errorf("error creating temp directory: %v\n", err)
	}

	err = os.MkdirAll(filepath.Join(tmpDir, tree), 0755)
	if err != nil {
		os.RemoveAll(tmpDir)
		return "", err
	}

	return tmpDir, nil
}

func main() {
	tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	if err != nil {
		fmt.Printf("unable to create test dir tree: %v\n", err)
		return
	}
	defer os.RemoveAll(tmpDir)
	os.Chdir(tmpDir)

	subDirToSkip := "skip"

	fmt.Println("On Unix:")
	err = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		return
	}
}
Output:

On Unix:
visited file or dir: "."
visited file or dir: "dir"
visited file or dir: "dir/to"
visited file or dir: "dir/to/walk"
skipping a dir without errors: skip
```

### func WalkDir  <- go1.16

``` go 
func WalkDir(root string, fn fs.WalkDirFunc) error
```

WalkDir walks the file tree rooted at root, calling fn for each file or directory in the tree, including root.

​	`WalkDir`函数遍历以 root 为根的文件树，对树中的每个文件或目录都会调用 fn 函数，包括根目录。

All errors that arise visiting files and directories are filtered by fn: see the fs.WalkDirFunc documentation for details.

​	所有访问文件和目录时产生的错误都会通过 fn 进行过滤：有关详细信息，请参见 fs.WalkDirFunc 的文档。

The files are walked in lexical order, which makes the output deterministic but requires WalkDir to read an entire directory into memory before proceeding to walk that directory.

​	文件按词法顺序遍历，这使得输出是确定性的，但要求 WalkDir 读取整个目录到内存中，然后才能继续遍历该目录。

WalkDir does not follow symbolic links.

​	WalkDir函数不会遵循符号链接。

WalkDir calls fn with paths that use the separator character appropriate for the operating system. This is unlike [io/fs.WalkDir](https://pkg.go.dev/io/fs#WalkDir), which always uses slash separated paths.

​	WalkDir函数调用 fn 以使用适用于操作系统的分隔符字符的路径，这与 [io/fs.WalkDir](https://pkg.go.dev/io/fs#WalkDir) 不同，后者总是使用斜杠分隔的路径。

## 类型

### type WalkFunc 

``` go 
type WalkFunc func(path string, info fs.FileInfo, err error) error
```

WalkFunc is the type of the function called by Walk to visit each file or directory.

​	WalkFunc 是 Walk 函数调用访问每个文件或目录的函数的类型。

The path argument contains the argument to Walk as a prefix. That is, if Walk is called with root argument "dir" and finds a file named "a" in that directory, the walk function will be called with argument "dir/a".

​	path 参数包含 Walk 的参数作为前缀。也就是说，如果 Walk 的根参数为 "dir"，并且在该目录中找到名为 "a" 的文件，则将使用参数 "dir/a" 调用 walk 函数。

The directory and file are joined with Join, which may clean the directory name: if Walk is called with the root argument "x/../dir" and finds a file named "a" in that directory, the walk function will be called with argument "dir/a", not "x/../dir/a".

​	目录和文件使用 Join 进行连接，这可能会清理目录名称：如果 Walk 被使用根参数 "x/../dir" 调用，并且在该目录中找到名为 "a" 的文件，则 walk 函数将使用参数 "dir/a" 而不是 "x/../dir/a" 调用。

The info argument is the fs.FileInfo for the named path.

​	info 参数是命名路径的 fs.FileInfo。

The error result returned by the function controls how Walk continues. If the function returns the special value SkipDir, Walk skips the current directory (path if info.IsDir() is true, otherwise path's parent directory). If the function returns the special value SkipAll, Walk skips all remaining files and directories. Otherwise, if the function returns a non-nil error, Walk stops entirely and returns that error.

​	函数返回的错误结果控制了 Walk 的继续。如果函数返回特殊值 SkipDir，则 Walk 跳过当前目录(如果 info.IsDir() 为 true，则为 path，否则为 path 的父目录)。如果函数返回特殊值 SkipAll，则 Walk 跳过所有剩余的文件和目录。否则，如果函数返回非 nil 错误，Walk 将停止遍历整个树并返回该错误。

The err argument reports an error related to path, signaling that Walk will not walk into that directory. The function can decide how to handle that error; as described earlier, returning the error will cause Walk to stop walking the entire tree.

​	err 参数报告与 path 相关的错误，表示 Walk 不会遍历该目录。函数可以决定如何处理该错误；如前所述，返回错误将导致 Walk 停止遍历整个树。

Walk calls the function with a non-nil err argument in two cases.

​	Walk 在两种情况下使用非 nil 的 err 参数调用函数。

First, if an os.Lstat on the root directory or any directory or file in the tree fails, Walk calls the function with path set to that directory or file's path, info set to nil, and err set to the error from os.Lstat.

​	首先，如果根目录或树中的任何目录或文件的 os.Lstat 失败，则 Walk 使用路径设置为该目录或文件的路径、info 设置为 nil 和 err 设置为 os.Lstat 的错误调用该函数。

Second, if a directory's Readdirnames method fails, Walk calls the function with path set to the directory's path, info, set to an fs.FileInfo describing the directory, and err set to the error from Readdirnames.

​	其次，如果一个目录的 Readdirnames 方法失败，Walk 使用路径设置为该目录的路径、info 设置为一个描述目录的 fs.FileInfo 和 err 设置为来自 Readdirnames 的错误调用该函数。