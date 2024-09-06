+++
title = "fs"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/io/fs@go1.23.0](https://pkg.go.dev/io/fs@go1.23.0)

Package fs defines basic interfaces to a file system. A file system can be provided by the host operating system but also by other packages.

​	`fs`包定义了与文件系统交互的基本接口。文件系统可以由操作系统提供，也可以由其他包提供。

## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/io/fs/fs.go;l=136)

``` go 
var (
	ErrInvalid    = errInvalid()    // "参数无效" "invalid argument"
	ErrPermission = errPermission() // "权限被拒绝" "permission denied"
	ErrExist      = errExist()      // "文件已经存在" "file already exists"
	ErrNotExist   = errNotExist()   // "文件不存在" "file does not exist"
	ErrClosed     = errClosed()     // "文件已经关闭" "file already closed"
)
```

Generic file system errors. Errors returned by file systems can be tested against these errors using errors.Is.

​	通用的文件系统错误。可以使用errors.Is将文件系统返回的错误与这些错误进行比较。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/io/fs/walk.go;l=20)

``` go 
var SkipAll = errors.New("skip everything and stop the walk")
```

SkipAll is used as a return value from WalkDirFuncs to indicate that all remaining files and directories are to be skipped. It is not returned as an error by any function.

​	SkipAll是从WalkDirFuncs返回的值，用于指示所有剩余的文件和目录都应该被跳过。任何函数都不会将其作为错误返回。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/io/fs/walk.go;l=15)

``` go 
var SkipDir = errors.New("skip this directory")
```

SkipDir is used as a return value from WalkDirFuncs to indicate that the directory named in the call is to be skipped. It is not returned as an error by any function.

​	SkipDir是从WalkDirFuncs返回的值，用于指示应该跳过调用中指定的目录。任何函数都不会将其作为错误返回。

## 函数

### func FormatDirEntry <-go1.21.0

```go
func FormatDirEntry(dir DirEntry) string
```

FormatDirEntry returns a formatted version of dir for human readability. Implementations of DirEntry can call this from a String method. The outputs for a directory named subdir and a file named hello.go are:

​	FormatDirEntry 返回 dir 的格式化版本，以便于人类阅读。DirEntry 的实现可以从 String 方法中调用此函数。对于名为 `subdir` 的目录和名为 `hello.go` 的文件，输出结果如下：

```
d subdir/
- hello.go
```

### func FormatFileInfo <-go1.21.0

```go
func FormatFileInfo(info FileInfo) string
```

FormatFileInfo returns a formatted version of info for human readability. Implementations of FileInfo can call this from a String method. The output for a file named "hello.go", 100 bytes, mode 0o644, created January 1, 1970 at noon is

​	FormatFileInfo 返回 info 的格式化版本，以便于人类阅读。FileInfo 的实现可以从 String 方法中调用此函数。对于名为 "hello.go" 的文件，大小为 100 字节，权限模式为 0o644，创建于 1970 年 1 月 1 日中午，输出结果为

```
-rw-r--r-- 100 1970-01-01 12:00:00 hello.go
```

### func Glob 

``` go 
func Glob(fsys FS, pattern string) (matches []string, err error)
```

Glob returns the names of all files matching pattern or nil if there is no matching file. The syntax of patterns is the same as in path.Match. The pattern may describe hierarchical names such as usr/*/bin/ed.

​	Glob函数返回与pattern匹配的所有文件的名称，如果没有匹配的文件，则返回nil。模式的语法与path.Match中的语法相同。模式可以描述分层名称，例如`usr/*/bin/ed`。

Glob ignores file system errors such as I/O errors reading directories. The only possible returned error is path.ErrBadPattern, reporting that the pattern is malformed.

​	Glob函数忽略文件系统错误，例如读取目录的I/O错误。唯一可能返回的错误是path.ErrBadPattern，报告模式格式不正确。

If fs implements GlobFS, Glob calls fs.Glob. Otherwise, Glob uses ReadDir to traverse the directory tree and look for matches for the pattern.

​	如果fs实现了GlobFS，则Glob函数调用fs.Glob。否则，Glob函数使用ReadDir遍历目录树并查找模式匹配项。

#### Glob My Example

![image-20230823204039795](fs_img/image-20230823204039795.png)

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// 使用 Glob 进行文件路径匹配，返回匹配的文件路径列表
	matches1, err := fs.Glob(os.DirFS("."), "*.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Matched files1:")
	for _, match := range matches1 {
		fmt.Println(match)
	}

	matches2, err := fs.Glob(os.DirFS("./subdir"), "*.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Matched files2:")
	for _, match := range matches2 {
		fmt.Println(match)
	}

}

// Output:
//Matched files1:
//hello.txt
//world.txt
//Matched files2:
//hi.txt
//nice.txt
```



### func ReadFile 

``` go 
func ReadFile(fsys FS, name string) ([]byte, error) {
	if fsys, ok := fsys.(ReadFileFS); ok {
		return fsys.ReadFile(name)
	}

	file, err := fsys.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var size int
	if info, err := file.Stat(); err == nil {
		size64 := info.Size()
		if int64(int(size64)) == size64 {
			size = int(size64)
		}
	}

	data := make([]byte, 0, size+1)
	for {
		if len(data) >= cap(data) {
			d := append(data[:cap(data)], 0)
			data = d[:len(data)]
		}
		n, err := file.Read(data[len(data):cap(data)])
		data = data[:len(data)+n]
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return data, err
		}
	}
}
```

ReadFile reads the named file from the file system fs and returns its contents. A successful call returns a nil error, not io.EOF. (Because ReadFile reads the whole file, the expected EOF from the final Read is not treated as an error to be reported.)

​	ReadFile 函数从文件系统fs中读取指定的文件并返回其内容。成功调用返回nil错误，而不是io.EOF。(因为ReadFile读取整个文件，因此不会将最终的EOF视为要报告的错误。)

If fs implements ReadFileFS, ReadFile calls fs.ReadFile. Otherwise ReadFile calls fs.Open and uses Read and Close on the returned file.

​	如果fs实现了ReadFileFS，则ReadFile调用fs.ReadFile。否则，ReadFile调用fs.Open并在返回的文件上使用Read和Close。

#### ReadFile My Example

![image-20230823204943901](fs_img/image-20230823204943901.png)

```go
package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
)

type MyReadFileFS struct{}

func (m *MyReadFileFS) ReadFile(name string) ([]byte, error) {
	// 检查文件是否存在
	if _, err := os.Stat(name); err != nil {
		return nil, err
	}

	// 打开文件
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	// 关闭文件
	defer file.Close()

	// 创建一个带缓冲的读取器
	reader := bufio.NewReader(file)

	var data = make([]byte, 4096)
	_, err = reader.Read(data)
	data = append([]byte("这是该方法的自定义内容，之后才是文件中的内容！"), data...)
	return data, nil
}

func (m *MyReadFileFS) Open(name string) (file fs.File, err error) {
	// 检查文件是否存在
	if _, err := os.Stat(name); err != nil {
		return nil, err
	}

	// 打开文件
	file, err = os.Open(name)
	if err != nil {
		return nil, err
	}

	// 返回文件读取器
	return file, nil
}

func main() {
	content1, err := fs.ReadFile(os.DirFS("."), "hello.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Matched file content is:", string(content1))

	content2, err := fs.ReadFile(&MyReadFileFS{}, "hello.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Matched file content is:", string(content2))
}

// Output:
//Matched file content is: Hello! Nice to meet you! (notice: All in one line and no newline)
//Matched file content is: 这是该方法的自定义内容，之后才是文件中的内容！Hello! Nice to meet you! (notice: All in one line and no newline)
```

### func ValidPath 

``` go 
func ValidPath(name string) bool
```

ValidPath reports whether the given path name is valid for use in a call to Open.

​	ValidPath 函数报告给定的路径名在调用 Open 时是否有效。

Path names passed to open are UTF-8-encoded, unrooted, slash-separated sequences of path elements, like “x/y/z”. Path names must not contain an element that is “.” or “..” or the empty string, except for the special case that the root directory is named “.”. Paths must not start or end with a slash: “/x” and “x/” are invalid.

​	传递给 Open 的路径名是以 UTF-8 编码的、无根的、斜杠分隔的路径元素序列，例如 "`x/y/z`"。路径名不得包含 "`.`" 或 "`..`" 或空字符串，但根目录的特殊情况是命名为 "`.`"。路径不能以斜杠开头或结尾，即 "`/x`" 和 "`x/`" 是无效的。

Note that paths are slash-separated on all systems, even Windows. Paths containing other characters such as backslash and colon are accepted as valid, but those characters must never be interpreted by an FS implementation as path element separators.

​	请注意，路径在所有系统上都是以斜杠分隔的，即使在 Windows 上也是如此。包含反斜杠和冒号等其他字符的路径被接受为有效，但这些字符绝不能被 FS 实现解释为路径元素分隔符。

#### ValidPath My Example

```go
package main

import (
	"fmt"
	"io/fs"
)

func main() {
	// 待验证的路径
	paths := []string{
		`tmp/\example.txt`,
		`tmp/:example.txt`,
		`tmp/example.txt`,
		`/tmp/example.txt`,
		`./tmp/example.txt`,
		`../tmp/example.txt`,
		`../tmp/example.txt`,
		`../tmp/example.txt`,
		`tmp/\subdir`,
		`tmp/:subdir`,
		`tmp/subdir`,
		`/tmp/subdir`,
		`./tmp/subdir/`,
		`../tmp/subdir/`,
		`tmp/\subdir/example.txt`,
		`tmp/:subdir/example.txt`,
		`tmp/subdir/example.txt`,
		`/tmp/subdir/example.txt`,
		`./tmp/subdir/example.txt`,
		`../tmp/subdir/example.txt`,
	}

	var validPaths []string
	var invalidPaths []string

	for _, path := range paths {
		// 使用 ValidPath 函数检查路径是否有效
		if fs.ValidPath(path) {
			validPaths = append(validPaths, path)
		} else {
			invalidPaths = append(invalidPaths, path)
		}
	}

	fmt.Println("有效路径有：")
	for _, path := range validPaths {
		fmt.Println(path)
	}

	fmt.Println("无效路径有：")
	for _, path := range invalidPaths {
		fmt.Println(path)
	}
}

//Output:
//有效路径有：
//tmp/\example.txt
//tmp/:example.txt
//tmp/example.txt
//tmp/\subdir
//tmp/:subdir
//tmp/subdir
//tmp/\subdir/example.txt
//tmp/:subdir/example.txt
//tmp/subdir/example.txt
//无效路径有：
///tmp/example.txt
//./tmp/example.txt
//../tmp/example.txt
//../tmp/example.txt
//../tmp/example.txt
///tmp/subdir
//./tmp/subdir/
//../tmp/subdir/
///tmp/subdir/example.txt
//./tmp/subdir/example.txt
//../tmp/subdir/example.txt
```

### func WalkDir 

``` go 
func WalkDir(fsys FS, root string, fn WalkDirFunc) error {
	info, err := Stat(fsys, root)
	if err != nil {
		err = fn(root, nil, err)
	} else {
		err = walkDir(fsys, root, &statDirEntry{info}, fn)
	}
	if err == SkipDir || err == SkipAll {
		return nil
	}
	return err
}
```

WalkDir walks the file tree rooted at root, calling fn for each file or directory in the tree, including root.

​	WalkDir 函数遍历以 `root` 为根的文件树，在树中的每个文件或目录(包括 `root`)上调用 fn。

All errors that arise visiting files and directories are filtered by fn: see the fs.WalkDirFunc documentation for details.

​	`fn` 过滤了遍历文件和目录时出现的所有错误：详见 [fs.WalkDirFunc](#type-walkdirfunc) 文档。

The files are walked in lexical order, which makes the output deterministic but requires WalkDir to read an entire directory into memory before proceeding to walk that directory.

​	文件以字典序遍历，这使输出是确定性的，但需要在继续遍历该目录之前将整个目录读入内存。

WalkDir does not follow symbolic links found in directories, but if root itself is a symbolic link, its target will be walked.

​	WalkDir 函数不会跟踪目录中发现的符号链接，但如果 `root` 本身是符号链接，则会遍历其目标。

#### WalkDir My Example

![image-20230824144819156](fs_img/image-20230824144819156.png)

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// 定义一个目录
	dir := "dir"

	fmt.Println("----------------------------1-------------------------------")
	num := 0
	// 使用 WalkDir 函数遍历目录
	if err := fs.WalkDir(os.DirFS(dir), ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(num, "-----------------------")
		fmt.Printf("path=%v,", path)
		num++

		if err != nil {
			return err
		}

		if d.IsDir() {
			fmt.Println(path, "是一个目录")
		} else {
			fmt.Println(path, "是一个文件")
		}

		return nil
	}); err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------------------------2-------------------------------")
	num = 0
	// 使用 WalkDir 函数遍历目录
	if err := fs.WalkDir(os.DirFS(dir), "subdir1", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(num, "-----------------------")
		fmt.Printf("path=%v,", path)
		num++

		if err != nil {
			return err
		}

		if d.IsDir() {
			fmt.Println(path, "是一个目录")
			if path == "subdir1/subsubdir1" {
				fmt.Println("不遍历subdir1/subsubdir1目录")
				return fs.SkipDir
			}
		} else {
			fmt.Println(path, "是一个文件")
		}

		return nil
	}); err != nil {
		fmt.Println(err)
	}

	fmt.Println("----------------------------3-------------------------------")
	num = 0
	// 使用 WalkDir 函数遍历目录
	if err := fs.WalkDir(os.DirFS(dir), "subdir1/subsubdir1", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(num, "-----------------------")
		fmt.Printf("path=%v,", path)
		num++

		if err != nil {
			return err
		}

		if d.IsDir() {
			fmt.Println(path, "是一个目录")
		} else {
			fmt.Println(path, "是一个文件")
		}

		fmt.Println("d.Name()=", d.Name())
		fmt.Println("d.IsDir()=", d.IsDir())
		info, err := d.Info()
		fmt.Printf("d.Info()=%#v,err=%v\n", info, err)
		fmt.Println("d.Type()=", d.Type())

		return nil
	}); err != nil {
		fmt.Println(err)
	}
}

// Output:
//----------------------------1-------------------------------
//0 -----------------------
//path=.,. 是一个目录
//1 -----------------------
//path=0.html,0.html 是一个文件
//2 -----------------------
//path=0.txt,0.txt 是一个文件
//3 -----------------------
//path=subdir1,subdir1 是一个目录
//4 -----------------------
//path=subdir1/1.txt,subdir1/1.txt 是一个文件
//5 -----------------------
//path=subdir1/2.txt,subdir1/2.txt 是一个文件
//6 -----------------------
//path=subdir1/3.html,subdir1/3.html 是一个文件
//7 -----------------------
//path=subdir1/subsubdir1,subdir1/subsubdir1 是一个目录
//8 -----------------------
//path=subdir1/subsubdir1/1_1.txt,subdir1/subsubdir1/1_1.txt 是一个文件
//9 -----------------------
//path=subdir1/subsubdir1/1_2.html,subdir1/subsubdir1/1_2.html 是一个文件
//10 -----------------------
//path=subdir2,subdir2 是一个目录
//11 -----------------------
//path=subdir2/4.txt,subdir2/4.txt 是一个文件
//12 -----------------------
//path=subdir2/5.txt,subdir2/5.txt 是一个文件
//13 -----------------------
//path=subdir2/6.html,subdir2/6.html 是一个文件
//----------------------------2-------------------------------
//0 -----------------------
//path=subdir1,subdir1 是一个目录
//1 -----------------------
//path=subdir1/1.txt,subdir1/1.txt 是一个文件
//2 -----------------------
//path=subdir1/2.txt,subdir1/2.txt 是一个文件
//3 -----------------------
//path=subdir1/3.html,subdir1/3.html 是一个文件
//4 -----------------------
//path=subdir1/subsubdir1,subdir1/subsubdir1 是一个目录
//不遍历subdir1/subsubdir1目录
//----------------------------3-------------------------------
//0 -----------------------
//path=subdir1/subsubdir1,subdir1/subsubdir1 是一个目录
//d.Name()= subsubdir1
//d.IsDir()= true
//d.Info()=&os.fileStat{name:"subsubdir1", size:4096, mode:0x800001fd, modTime:time.Time{wall:0x2b836688, ext:63828778161, loc:(*time.Location)(0x544ec0)}, sys:syscall.Stat_t{Dev:0x820, Ino:0x2d74, Nlink:0x2, Mode:0x41fd, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:4096, Blksize:409
//6, Blocks:8, Atim:syscall.Timespec{Sec:1693181479, Nsec:10036347}, Mtim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, Ctim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, X__unused:[3]int64{0, 0, 0}}},err=<nil>
//d.Type()= d---------
//1 -----------------------
//path=subdir1/subsubdir1/1_1.txt,subdir1/subsubdir1/1_1.txt 是一个文件
//d.Name()= 1_1.txt
//d.IsDir()= false
//d.Info()=&os.fileStat{name:"1_1.txt", size:0, mode:0x1b4, modTime:time.Time{wall:0x2b836688, ext:63828778161, loc:(*time.Location)(0x544ec0)}, sys:syscall.Stat_t{Dev:0x820, Ino:0x2d75, Nlink:0x1, Mode:0x81b4, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:0, Blksize:4096, Blocks:0, A
//tim:syscall.Timespec{Sec:1693182481, Nsec:600036816}, Mtim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, Ctim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, X__unused:[3]int64{0, 0, 0}}},err=<nil>
//d.Type()= ----------
//2 -----------------------
//path=subdir1/subsubdir1/1_2.html,subdir1/subsubdir1/1_2.html 是一个文件
//d.Name()= 1_2.html
//d.IsDir()= false
//d.Info()=&os.fileStat{name:"1_2.html", size:123, mode:0x1b4, modTime:time.Time{wall:0x2b836688, ext:63828778161, loc:(*time.Location)(0x544ec0)}, sys:syscall.Stat_t{Dev:0x820, Ino:0x2d76, Nlink:0x1, Mode:0x81b4, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:123, Blksize:4096, Blocks
//:8, Atim:syscall.Timespec{Sec:1693181480, Nsec:430036484}, Mtim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, Ctim:syscall.Timespec{Sec:1693181361, Nsec:730031752}, X__unused:[3]int64{0, 0, 0}}},err=<nil>
//d.Type()= ----------
```

#### WalkDir Example

``` go 
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	root := "/usr/local/go/bin"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
}

```



## 类型

### type DirEntry 

``` go 
type DirEntry interface {
    // Name returns the name of the file (or subdirectory) described by the entry.
	// This name is only the final element of the path (the base name), not the entire path.
	// For example, Name would return "hello.go" not "home/gopher/hello.go".
    // Name 返回该条目描述的文件(或子目录)的名称。
    // 该名称仅为路径的最后一个元素(基本名称)，而不是整个路径。
    // 例如，Name 将返回"hello.go"而不是"home/gopher/hello.go"。
	Name() string

    // IsDir reports whether the entry describes a directory.
    // IsDir 报告该条目是否描述一个目录。
	IsDir() bool

    // Type returns the type bits for the entry.
	// The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
    // Type 返回该条目的类型位。
    // 类型位是 FileMode 常量的子集，
    // 是由 FileMode.Type 方法返回的常量之一。
	Type() FileMode

    // Info returns the FileInfo for the file or subdirectory described by the entry.
	// The returned FileInfo may be from the time of the original directory read
	// or from the time of the call to Info. If the file has been removed or renamed
	// since the directory read, Info may return an error satisfying errors.Is(err, ErrNotExist).
	// If the entry denotes a symbolic link, Info reports the information about the link itself,
	// not the link's target.
    // Info 返回该条目描述的文件或子目录的 FileInfo。
    // 返回的 FileInfo 可能来自原始目录读取的时间或 Info 调用的时间。
    // 如果文件已在目录读取后被删除或重命名，
    // Info 可能返回一个满足 errors.Is(err, ErrNotExist) 的错误。
    // 如果条目表示一个符号链接，
    // 则 Info 报告有关链接本身的信息，而不是链接的目标。
	Info() (FileInfo, error)
}
```

A DirEntry is an entry read from a directory (using the ReadDir function or a ReadDirFile's ReadDir method).

​	DirEntry 是从目录中读取的一个条目(使用 ReadDir 函数或 ReadDirFile 的 ReadDir 方法)。

#### func FileInfoToDirEntry  <- go1.17

``` go 
func FileInfoToDirEntry(info FileInfo) DirEntry
```

FileInfoToDirEntry returns a DirEntry that returns information from info. If info is nil, FileInfoToDirEntry returns nil.

​	FileInfoToDirEntry 函数返回一个从 `info` 中获取信息的 DirEntry。如果 info 为 nil，则 FileInfoToDirEntry 返回 nil。

##### FileInfoToDirEntry My Example

![image-20230824170411420](fs_img/image-20230824170411420.png)

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// 获取文件信息
	fileInfo1, err := os.Stat("dir")
	if err != nil {
		fmt.Println("无法获取文件信息:", err)
		return
	}

	// 将文件信息转换为目录条目
	dirEntry := fs.FileInfoToDirEntry(fileInfo1)

	// 打印目录条目的名称和类型
	fmt.Println("名称:", dirEntry.Name())
	fmt.Println("类型:", dirEntry.Type())
	fmt.Println("是目录？", dirEntry.IsDir())
	info, err := dirEntry.Info()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dirEntry.Info()=%#v\n", info)

	// 获取文件信息
	fileInfo2, err := os.Stat("dir/hello.txt")
	if err != nil {
		fmt.Println("无法获取文件信息:", err)
		return
	}

	// 将文件信息转换为目录条目
	dirEntry = fs.FileInfoToDirEntry(fileInfo2)

	// 打印目录条目的名称和类型
	fmt.Println("名称:", dirEntry.Name())
	fmt.Println("类型:", dirEntry.Type())
	fmt.Println("是目录？", dirEntry.IsDir())
	info, err = dirEntry.Info()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("dirEntry.Info()=%#v\n", info)
}
// Output:
//名称: hello.txt
//类型: ----------
//PS F:\Devs\MyCodes\go_std_examples\io\fs\f_FileInfoToDirEntry> go run .\FileInfoToDirEntry.go
//# command-line-arguments
//.\FileInfoToDirEntry.go:32:2: fileInfo2 declared and not used
//PS F:\Devs\MyCodes\go_std_examples\io\fs\f_FileInfoToDirEntry> go run .\FileInfoToDirEntry.go
//名称: dir
//类型: d---------
//是目录？ true
//dirEntry.Info()=&os.fileStat{name:"dir", FileAttributes:0x10, CreationTime:syscall.Filetime{LowDateTime:0xa023380f, HighDateTime:0x1d9d667}, LastAccessTime:sys
//call.Filetime{LowDateTime:0x612f7da5, HighDateTime:0x1d9d668}, LastWriteTime:syscall.Filetime{LowDateTime:0xd4831f40, HighDateTime:0x1d9d667}, FileSizeHigh:0x0
//, FileSizeLow:0x0, ReparseTag:0x0, filetype:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"F:\\Devs\\MyCodes\\go_std_examples\\io\\fs\\f_FileInfoToDirEntry\\d
//ir", vol:0x0, idxhi:0x0, idxlo:0x0}
//名称: hello.txt
//类型: ----------
//是目录？ false
//dirEntry.Info()=&os.fileStat{name:"hello.txt", FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xa6109178, HighDateTime:0x1d9d667}, LastAccessTi
//me:syscall.Filetime{LowDateTime:0x612f7da5, HighDateTime:0x1d9d668}, LastWriteTime:syscall.Filetime{LowDateTime:0xd4831f40, HighDateTime:0x1d9d667}, FileSizeHi
//gh:0x0, FileSizeLow:0x6, ReparseTag:0x0, filetype:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"F:\\Devs\\MyCodes\\go_std_examples\\io\\fs\\f_FileInfoToDirEn
//try\\dir\\hello.txt", vol:0x0, idxhi:0x0, idxlo:0x0}

```

#### func ReadDir 

``` go 
func ReadDir(fsys FS, name string) ([]DirEntry, error)
```

ReadDir reads the named directory and returns a list of directory entries sorted by filename.

​	ReadDir函数读取命名的目录并返回一个按文件名排序的目录条目列表。

If fs implements ReadDirFS, ReadDir calls fs.ReadDir. Otherwise ReadDir calls fs.Open and uses ReadDir and Close on the returned file.

​	如果 fs 实现了 ReadDirFS，则 ReadDir 调用 fs.ReadDir。否则，ReadDir调用fs.Open并对返回的文件使用ReadDir和Close。

##### ReadDir My Example

注意：ReadDir 函数并不会读取，命名的目录下的子目录中的文件！

另请参见[ReadDirFS My Example](#type-readdirfs)

![image-20230824171453802](fs_img/image-20230824171453802.png)

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	dirEntrys, err := fs.ReadDir(os.DirFS("dir"), ".")

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, dirEntry := range dirEntrys {
		fmt.Println(i, "--------------------------")
		fmt.Println("名称:", dirEntry.Name())
		fmt.Println("类型:", dirEntry.Type())
		fmt.Println("是目录？", dirEntry.IsDir())
		info, err := dirEntry.Info()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("dirEntry.Info()=%#v\n", info)
		fmt.Println("info.Name()=", info.Name())
		fmt.Println("info.Size()=", info.Size())
		fmt.Println("info.Mode()=", info.Mode())
		fmt.Println("info.ModTime()=", info.ModTime())
		fmt.Println("info.IsDir()=", info.IsDir())
		fmt.Printf("info.Sys()=%#v\n", info.Sys())
	}
}

// Output:
//0 --------------------------
//名称: 1.txt
//类型: ----------
//是目录？ false
//dirEntry.Info()=&os.fileStat{name:"1.txt", FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xfc0f5a33, HighDateTime:0x1d9d668}, LastAccessTime:syscall.Filetime{LowDat
//eTime:0xac576b6c, HighDateTime:0x1d9d669}, LastWriteTime:syscall.Filetime{LowDateTime:0xac576b6c, HighDateTime:0x1d9d669}, FileSizeHigh:0x0, FileSizeLow:0x8, ReparseTag:0x0, filetyp
//e:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"", vol:0x2c188df6, idxhi:0x20000, idxlo:0x9e959}
//info.Name()= 1.txt
//info.Size()= 8
//info.Mode()= -rw-rw-rw-
//info.ModTime()= 2023-08-24 17:02:13.5460716 +0800 CST
//info.IsDir()= false
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xfc0f5a33, HighDateTime:0x1d9d668}, LastAccessTime:syscall.Filetime{LowDat
//eTime:0xac576b6c, HighDateTime:0x1d9d669}, LastWriteTime:syscall.Filetime{LowDateTime:0xac576b6c, HighDateTime:0x1d9d669}, FileSizeHigh:0x0, FileSizeLow:0x8}
//1 --------------------------
//名称: 2.txt
//类型: ----------
//是目录？ false
//dirEntry.Info()=&os.fileStat{name:"2.txt", FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xcc72153, HighDateTime:0x1d9d669}, LastAccessTime:syscall.Filetime{LowDate
//Time:0xac4b7e88, HighDateTime:0x1d9d669}, LastWriteTime:syscall.Filetime{LowDateTime:0xac4b7e88, HighDateTime:0x1d9d669}, FileSizeHigh:0x0, FileSizeLow:0x8, ReparseTag:0x0, filetype
//:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"", vol:0x2c188df6, idxhi:0x50000, idxlo:0x9e956}
//info.Name()= 2.txt
//info.Size()= 8
//info.Mode()= -rw-rw-rw-
//info.ModTime()= 2023-08-24 17:02:13.4679176 +0800 CST
//info.IsDir()= false
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xcc72153, HighDateTime:0x1d9d669}, LastAccessTime:syscall.Filetime{LowDate
//Time:0xac4b7e88, HighDateTime:0x1d9d669}, LastWriteTime:syscall.Filetime{LowDateTime:0xac4b7e88, HighDateTime:0x1d9d669}, FileSizeHigh:0x0, FileSizeLow:0x8}
//2 --------------------------
//名称: subdir
//类型: d---------
//是目录？ true
//dirEntry.Info()=&os.fileStat{name:"subdir", FileAttributes:0x10, CreationTime:syscall.Filetime{LowDateTime:0x49b7a14a, HighDateTime:0x1d9d66a}, LastAccessTime:syscall.Filetime{LowDa
//teTime:0x54958ccd, HighDateTime:0x1d9d66a}, LastWriteTime:syscall.Filetime{LowDateTime:0x541df8f8, HighDateTime:0x1d9d66a}, FileSizeHigh:0x0, FileSizeLow:0x0, ReparseTag:0x0, filety
//pe:0x0, Mutex:sync.Mutex{state:0, sema:0x0}, path:"", vol:0x2c188df6, idxhi:0x50000, idxlo:0x9e95e}
//info.Name()= subdir
//info.Size()= 0
//info.Mode()= drwxrwxrwx
//info.ModTime()= 2023-08-24 17:06:55.0268152 +0800 CST
//info.IsDir()= true
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x10, CreationTime:syscall.Filetime{LowDateTime:0x49b7a14a, HighDateTime:0x1d9d66a}, LastAccessTime:syscall.Filetime{LowDat
//eTime:0x54958ccd, HighDateTime:0x1d9d66a}, LastWriteTime:syscall.Filetime{LowDateTime:0x541df8f8, HighDateTime:0x1d9d66a}, FileSizeHigh:0x0, FileSizeLow:0x0}

```

### type FS 

``` go 
type FS interface {
    // Open opens the named file.
	//
	// When Open returns an error, it should be of type *PathError
	// with the Op field set to "open", the Path field set to name,
	// and the Err field describing the problem.
	//
	// Open should reject attempts to open names that do not satisfy
	// ValidPath(name), returning a *PathError with Err set to
	// ErrInvalid or ErrNotExist.
    // Open打开命名的文件。
	//
	// 当Open返回一个错误时，它应该是*PathError类型，Op字段设置为 "open"，Path字段设置为name，Err字段描述了问题所在。
	//
	// Open应该拒绝打开不符合ValidPath(name)的名字的尝试，返回一个*PathError，Err设置为ErrInvalid或ErrNotExist。
	Open(name string) (File, error)
}
```

An FS provides access to a hierarchical file system.

​	FS 提供对一个分层文件系统的访问。

The FS interface is the minimum implementation required of the file system. A file system may implement additional interfaces, such as ReadFileFS, to provide additional or optimized functionality.

​	FS 接口是文件系统所需的最小实现。一个文件系统可能会实现其他接口(如 ReadFileFS)以提供附加或优化的功能。

#### func Sub 

``` go 
func Sub(fsys FS, dir string) (FS, error) {
	if !ValidPath(dir) {
		return nil, &PathError{Op: "sub", Path: dir, Err: errors.New("invalid name")}
	}
	if dir == "." {
		return fsys, nil
	}
	if fsys, ok := fsys.(SubFS); ok {
		return fsys.Sub(dir)
	}
	return &subFS{fsys, dir}, nil
}
```

Sub returns an FS corresponding to the subtree rooted at fsys's dir.

​	Sub 函数返回一个对应于以 `fsys` 的 `dir` 为根的子树的 FS。

If dir is ".", Sub returns fsys unchanged. Otherwise, if fs implements SubFS, Sub returns fsys.Sub(dir). Otherwise, Sub returns a new FS implementation sub that, in effect, implements sub.Open(name) as fsys.Open(path.Join(dir, name)). The implementation also translates calls to ReadDir, ReadFile, and Glob appropriately.

​	如果 dir 为"`.`"，则 Sub 返回未更改的 fsys。否则，如果 fs 实现了 SubFS，则 Sub 返回 fsys.Sub(dir)。否则，Sub 返回一个新的 FS 实现 sub，该实现实际上将 sub.Open(name) 实现为 fsys.Open(path.Join(dir, name))。该实现还适当地翻译了对 ReadDir、ReadFile 和 Glob 的调用。

Note that Sub(os.DirFS("/"), "prefix") is equivalent to os.DirFS("/prefix") and that neither of them guarantees to avoid operating system accesses outside "/prefix", because the implementation of os.DirFS does not check for symbolic links inside "/prefix" that point to other directories. That is, os.DirFS is not a general substitute for a chroot-style security mechanism, and Sub does not change that fact.

​	请注意，Sub(os.DirFS("/"), "prefix") 等同于 os.DirFS("/prefix")，并且它们都不能保证避免操作系统对"/prefix "之外的访问，因为os.DirFS的实现并不检查"/prefix "内指向其他目录的符号链接。也就是说，os.DirFS 并不是 chroot 式安全机制的通用替代品，Sub 并不能改变这个事实。

##### Sub My Example

![image-20230824175308249](fs_img/image-20230824175308249.png)

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println("----------------------------------------------")
	fsys1, err := fs.Sub(os.DirFS("dir"), "subdir")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%T\n", fsys1)

	file, err := fsys1.Open("1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileData := make([]byte, 4096)
	n, err := file.Read(fileData)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("共有%d个字节的内容,内容是：%s\n", n, string(fileData))

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("fileInfo.Name()=", fileInfo.Name())
	fmt.Println("fileInfo.Size()=", fileInfo.Size())
	fmt.Println("fileInfo.Mode()=", fileInfo.Mode())
	fmt.Println("fileInfo.ModTime()=", fileInfo.ModTime())
	fmt.Println("fileInfo.IsDir()=", fileInfo.IsDir())
	fmt.Printf("fileInfo.Sys()=%#v\n", fileInfo.Sys())

	fmt.Println("----------------------------------------------")
	fsys2, err := fs.Sub(os.DirFS("."), "dir")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%T\n", fsys2)

	fmt.Println("----------------------------------------------")
	fsys3, err := fs.Sub(os.DirFS("."), "dir/subdir")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%T\n", fsys3)

	file, err = fsys3.Open("1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileData = make([]byte, 4096)
	n, err = file.Read(fileData)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("共有%d个字节的内容,内容是：%s\n", n, string(fileData))

	fileInfo, err = file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("fileInfo.Name()=", fileInfo.Name())
	fmt.Println("fileInfo.Size()=", fileInfo.Size())
	fmt.Println("fileInfo.Mode()=", fileInfo.Mode())
	fmt.Println("fileInfo.ModTime()=", fileInfo.ModTime())
	fmt.Println("fileInfo.IsDir()=", fileInfo.IsDir())
	fmt.Printf("fileInfo.Sys()=%#v\n", fileInfo.Sys())
}

// Output:
//----------------------------------------------
//*fs.subFS
//共有8个字节的内容,内容是：content1
//fileInfo.Name()= 1.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-rw-
//fileInfo.ModTime()= 2023-08-24 17:23:58.9674647 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xfde7ea1, HighDateTime:0x1d9d66c}, LastAccessTime:syscall.Filetime{Low
//DateTime:0x7fcc99df, HighDateTime:0x1d9d670}, LastWriteTime:syscall.Filetime{LowDateTime:0xb66eea97, HighDateTime:0x1d9d66c}, FileSizeHigh:0x0, FileSizeLow:0x8}
//----------------------------------------------
//*fs.subFS
//----------------------------------------------
//*fs.subFS
//共有8个字节的内容,内容是：content1
//fileInfo.Name()= 1.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-rw-
//fileInfo.ModTime()= 2023-08-24 17:23:58.9674647 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0xfde7ea1, HighDateTime:0x1d9d66c}, LastAccessTime:syscall.Filetime{Low
//DateTime:0x7fcfd9df, HighDateTime:0x1d9d670}, LastWriteTime:syscall.Filetime{LowDateTime:0xb66eea97, HighDateTime:0x1d9d66c}, FileSizeHigh:0x0, FileSizeLow:0x8}

```

### type File 

``` go 
type File interface {
	Stat() (FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}
```

A File provides access to a single file. The File interface is the minimum implementation required of the file. Directory files should also implement ReadDirFile. A file may implement io.ReaderAt or io.Seeker as optimizations.

​	File接口提供对单个文件的访问。File接口是文件所需的最小实现。目录文件还应该实现ReadDirFile。文件可以实现io.ReaderAt或io.Seeker作为优化。

### type FileInfo 

``` go 
type FileInfo interface {
	Name() string       // base name of the file// 文件的基本名称
	Size() int64        // length in bytes for regular files; system-dependent for others// 普通文件的长度，以字节为单位；其他文件则取决于系统。
	Mode() FileMode     // file mode bits // 文件的模式位
	ModTime() time.Time // modification time// 修改时间
	IsDir() bool        // abbreviation for Mode().IsDir()// Mode().IsDir()的缩写。
	Sys() any           // underlying data source (can return nil)// 底层数据源(可以返回nil)。
}
```

A FileInfo describes a file and is returned by Stat.

​	FileInfo描述了一个文件，并由Stat函数返回。

#### func Stat 

``` go 
func Stat(fsys FS, name string) (FileInfo, error)
```

Stat returns a FileInfo describing the named file from the file system.

​	Stat函数从文件系统返回描述命名文件的FileInfo。

If fs implements StatFS, Stat calls fs.Stat. Otherwise, Stat opens the file to stat it.

​	如果fs实现了StatFS，则Stat函数调用fs.Stat。否则，Stat打开文件以获取其状态。

##### Stat My Example

![image-20230824180701916](fs_img/image-20230824180701916.png)

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	fmt.Println("1 ----------------------------------")
	fileInfo, err := fs.Stat(os.DirFS("dir"), "1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("2 ----------------------------------")
	fileInfo, err = fs.Stat(os.DirFS("dir/subdir1"), "1.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("fileInfo.Name()=", fileInfo.Name())
	fmt.Println("fileInfo.Size()=", fileInfo.Size())
	fmt.Println("fileInfo.Mode()=", fileInfo.Mode())
	fmt.Println("fileInfo.ModTime()=", fileInfo.ModTime())
	fmt.Println("fileInfo.IsDir()=", fileInfo.IsDir())
	fmt.Printf("fileInfo.Sys()=%#v\n", fileInfo.Sys())

	fmt.Println("3 ----------------------------------")
	fileInfo, err = fs.Stat(os.DirFS("dir/subdir2"), "2.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("fileInfo.Name()=", fileInfo.Name())
	fmt.Println("fileInfo.Size()=", fileInfo.Size())
	fmt.Println("fileInfo.Mode()=", fileInfo.Mode())
	fmt.Println("fileInfo.ModTime()=", fileInfo.ModTime())
	fmt.Println("fileInfo.IsDir()=", fileInfo.IsDir())
	fmt.Printf("fileInfo.Sys()=%#v\n", fileInfo.Sys())
}

// Output:
//1 ----------------------------------
//CreateFile 1.txt: The system cannot find the file specified.
//2 ----------------------------------
//fileInfo.Name()= 1.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-rw-
//fileInfo.ModTime()= 2023-08-24 17:57:10.8000609 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x52d85ed0, HighDateTime:0x1d9d671}, LastAccessTime:syscall.Filetime{Lo
//wDateTime:0x9596a893, HighDateTime:0x1d9d672}, LastWriteTime:syscall.Filetime{LowDateTime:0x59a87361, HighDateTime:0x1d9d671}, FileSizeHigh:0x0, FileSizeLow:0x8}
//3 ----------------------------------
//fileInfo.Name()= 2.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-rw-
//fileInfo.ModTime()= 2023-08-24 18:04:51.5907555 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x59a9fdde, HighDateTime:0x1d9d671}, LastAccessTime:syscall.Filetime{Lo
//wDateTime:0x9596a893, HighDateTime:0x1d9d672}, LastWriteTime:syscall.Filetime{LowDateTime:0x6c4f87e3, HighDateTime:0x1d9d672}, FileSizeHigh:0x0, FileSizeLow:0x8}

```

### type FileMode 

``` go 
type FileMode uint32
```

A FileMode represents a file's mode and permission bits. The bits have the same definition on all systems, so that information about files can be moved from one system to another portably. Not all bits apply to all systems. The only required bit is ModeDir for directories.

​	FileMode表示一个文件的模式和权限位。这些位在所有系统上具有相同的定义，以便可以在不同系统之间可移植地移动文件信息。并非所有位都适用于所有系统。唯一需要的位是目录的ModeDir。

``` go 
const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
    //单字母是String方法的格式化所使用的缩略语。
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory// d: 是一个目录
	ModeAppend                                     // a: append-only // a: 仅限附加使用
	ModeExclusive                                  // l: exclusive use // l：独家使用
	ModeTemporary                                  // T: temporary file; Plan 9 only// T：临时文件；仅适用于Plan 9
	ModeSymlink                                    // L: symbolic link// L：象征性链接
	ModeDevice                                     // D: device file// D: 设备文件
	ModeNamedPipe                                  // p: named pipe (FIFO) // p：命名的管道(FIFO)。
	ModeSocket                                     // S: Unix domain socket// S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set // c：Unix字符设备，当ModeDevice被设置时
	ModeSticky                                     // t: sticky // t：粘性
	ModeIrregular                                  // ?: non-regular file; nothing else is known about this file // ? :非规则文件；关于这个文件的其他信息一无所知。

	// Mask for the type bits. For regular files, none will be set.
    // 类型位的屏蔽。对于常规文件，将不设置。
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

	ModePerm FileMode = 0777 // Unix permission bits
)
```

The defined file mode bits are the most significant bits of the FileMode. The nine least-significant bits are the standard Unix rwxrwxrwx permissions. The values of these bits should be considered part of the public API and may be used in wire protocols or disk representations: they must not be changed, although new bits might be added.

​	定义的文件模式位是 FileMode 的最高位。最低的九位是标准 Unix rwxrwxrwx 权限位。这些位的值应该被视为公共 API 的一部分，可以在传输协议或磁盘表示中使用：它们不得被更改，但可以添加新的位。

#### (FileMode) IsDir 

``` go 
func (m FileMode) IsDir() bool
```

IsDir reports whether m describes a directory. That is, it tests for the ModeDir bit being set in m.

​	IsDir 方法报告 `m` 是否描述一个目录。也就是说，它测试 ModeDir 位是否在 `m` 中被设置。

##### IsDir My Example

![image-20230824190949850](fs_img/image-20230824190949850.png)

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取文件信息
	fileInfo, err := os.Stat("dir")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())

	// 检查文件是否为目录
	if fileInfo.Mode().IsDir() {
		fmt.Println(fileInfo.Name(), "是一个目录")
	} else {
		fmt.Println(fileInfo.Name(), "不是一个目录")
	}

	fileInfo, err = os.Stat("dir/1.txt")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())

	// 检查文件是否为目录
	if fileInfo.Mode().IsDir() {
		fmt.Println(fileInfo.Name(), "是一个目录")
	} else {
		fmt.Println(fileInfo.Name(), "不是一个目录")
	}
}

// Output:
//fileInfo.Mode()'s type is fs.FileMode
//dir 是一个目录
//fileInfo.Mode()'s type is fs.FileMode
//1.txt 不是一个目录

```

#### (FileMode) IsRegular 

``` go 
func (m FileMode) IsRegular() bool
```

IsRegular reports whether m describes a regular file. That is, it tests that no mode type bits are set.

​	IsRegular 方法报告 `m` 是否描述一个普通文件。也就是说，它测试是否没有设置任何模式类型位。

##### IsRegular My Example

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取文件信息
	fileInfo, err := os.Stat("dir")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())
	fmt.Println(fileInfo.Name(), "是一个普通文件？", fileInfo.Mode().IsRegular())

	fileInfo, err = os.Stat("dir/1.txt")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())
	fmt.Println(fileInfo.Name(), "是一个普通文件？", fileInfo.Mode().IsRegular())
}

// Output:
//fileInfo.Mode()'s type is fs.FileMode
//dir 是一个普通文件？ false
//fileInfo.Mode()'s type is fs.FileMode
//1.txt 是一个普通文件？ true
```

#### (FileMode) Perm 

``` go 
func (m FileMode) Perm() FileMode
```

Perm returns the Unix permission bits in m (m & ModePerm).

​	Perm 方法返回 `m` 中的 Unix 权限位(m＆ModePerm)。

##### Perm My Example

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取文件信息
	fileInfo, err := os.Stat("dir")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())
	fmt.Println(fileInfo.Name(), " Perm()=", fileInfo.Mode().Perm())
	fmt.Println(fileInfo.Name(), " String()=", fileInfo.Mode().String())
	fmt.Println(fileInfo.Name(), " Perm().String()=", fileInfo.Mode().Perm().String())

	fileInfo, err = os.Stat("dir/1.txt")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())
	fmt.Println(fileInfo.Name(), " Perm()=", fileInfo.Mode().Perm())
	fmt.Println(fileInfo.Name(), " String()=", fileInfo.Mode().String())
	fmt.Println(fileInfo.Name(), " Perm().String()=", fileInfo.Mode().Perm().String())
}

// Output:
//fileInfo.Mode()'s type is fs.FileMode
//dir  Perm()= -rwxrwxrwx
//dir  String()= drwxrwxrwx
//dir  Perm().String()= -rwxrwxrwx
//fileInfo.Mode()'s type is fs.FileMode
//1.txt  Perm()= -rw-rw-rw-
//1.txt  String()= -rw-rw-rw-
//1.txt  Perm().String()= -rw-rw-rw-
```

#### (FileMode) String 

``` go 
func (m FileMode) String() string
```

##### String My Example

参见 [Perm](#filemode-perm)。

#### (FileMode) Type 

``` go 
func (m FileMode) Type() FileMode
```

Type returns type bits in m (m & ModeType).

​	Type 方法返回 `m` 中的类型位(m＆ModeType)。

##### Type My Example

![image-20230824192734887](fs_img/image-20230824192734887.png)

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 获取文件信息
	fileInfo, err := os.Stat("dir")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())
	fmt.Printf("fileInfo.Mode().Type()'s type is %T\n", fileInfo.Mode().Type())
	fmt.Println(fileInfo.Name(), " Type()=", fileInfo.Mode().Type())
	fmt.Println(fileInfo.Name(), " Type().String()=", fileInfo.Mode().Type().String())

	fileInfo, err = os.Stat("dir/1.txt")
	if err != nil {
		log.Fatalf("无法获取文件信息: %v", err)
	}

	fmt.Printf("fileInfo.Mode()'s type is %T\n", fileInfo.Mode())
	fmt.Printf("fileInfo.Mode().Type()'s type is %T\n", fileInfo.Mode().Type())
	fmt.Println(fileInfo.Name(), " Type()=", fileInfo.Mode().Type())
	fmt.Println(fileInfo.Name(), " Type().String()=", fileInfo.Mode().Type().String())
}

// Output:
//fileInfo.Mode()'s type is fs.FileMode
//fileInfo.Mode().Type()'s type is fs.FileMode
//dir  Type()= d---------
//dir  Type().String()= d---------
//fileInfo.Mode()'s type is fs.FileMode       
//fileInfo.Mode().Type()'s type is fs.FileMode
//1.txt  Type()= ----------
//1.txt  Type().String()= ----------


```

### type GlobFS 

``` go 
type GlobFS interface {
	FS

    // Glob returns the names of all files matching pattern,
	// providing an implementation of the top-level
	// Glob function.
    // Glob 返回与 pattern 匹配的所有文件的名称，
    // 提供了顶层 Glob 函数的实现。
	Glob(pattern string) ([]string, error)
}
```

A GlobFS is a file system with a Glob method.

​	GlobFS 是具有 Glob 方法的文件系统。

#### GlobFS My Example

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"strings"
)

type MyFs struct {
	Ifs fs.FS
}

func (m MyFs) Open(name string) (fs.File, error) {
	return m.Ifs.Open(name)
}

func (m MyFs) Glob(pattern string) ([]string, error) {
	var filenames []string
	if strings.HasPrefix(pattern, ".") {
		pattern = strings.Replace(pattern, `.`, `\.`, -1)
		pattern = strings.TrimLeft(pattern, `\`)
	} else {
		pattern = strings.Replace(pattern, `.`, `\.`, -1)
	}

	if strings.HasPrefix(pattern, "*") {
		pattern = "." + pattern
	}

	//fmt.Println("pattern=", pattern)

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("pattern is invalid: %w", err)
	}
	if err := fs.WalkDir(m.Ifs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if re.MatchString(d.Name()) {
			filenames = append(filenames, d.Name())
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return filenames, nil
}

func main() {
	myFs := MyFs{os.DirFS("dir")}
	file, err := myFs.Open("1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileData := make([]byte, 4096)
	file.Read(fileData)
	fmt.Println("文件中的内容是：", string(fileData))

	matches, err := myFs.Glob(`*.txt`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("匹配到的文件有:")
	for _, match := range matches {
		fmt.Println(match)
	}
}

// Output:
//文件中的内容是： content1
//匹配到的文件有:
//1.txt
//2.txt
```

### type PathError 

``` go 
type PathError struct {
	Op   string
	Path string
	Err  error
}
```

PathError records an error and the operation and file path that caused it.

​	PathError 记录了一个错误以及导致该错误的操作和文件路径。

#### (*PathError) Error 

``` go 
func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

```

##### Error My Example

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// 尝试打开不存在的文件
	_, err := os.Open("nonexistent.txt")
	var pathErr error
	if err != nil {
		// 创建一个 PathError 错误
		pathErr = &fs.PathError{
			Op:   "open",
			Path: "nonexistent.txt",
			Err:  err,
		}
		// 打印错误信息
		fmt.Println("发生错误:", pathErr)
		fmt.Println("发生错误:", pathErr.Error())
	}
}
// Output:
//发生错误: open nonexistent.txt: open nonexistent.txt: The system cannot find the file specified.
//发生错误: open nonexistent.txt: open nonexistent.txt: The system cannot find the file specified.
```

#### (*PathError) Timeout 

``` go 
func (e *PathError) Timeout() bool {
	t, ok := e.Err.(interface{ Timeout() bool })
	return ok && t.Timeout()
}
```

Timeout reports whether this error represents a timeout.

​	Timeout方法报告此错误是否表示超时。

##### Timeout My Example

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// Create a timeout error manually
	pathErr := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrDeadlineExceeded,
	}

	if pathErr.Timeout() {
		fmt.Println("Timeout occurred")
	} else {
		fmt.Println("No timeout occurred")
	}

	// Simulate a non-timeout error
	pathErr2 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrNotExist,
	}

	if pathErr2.Timeout() {
		fmt.Println("Timeout occurred")
	} else {
		fmt.Println("No timeout occurred")
	}
}

// Output:
//Timeout occurred
//No timeout occurred

```



#### (*PathError) Unwrap 

``` go 
func (e *PathError) Unwrap() error { return e.Err }
```

##### Unwrap My Example

```go
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	pathErr1 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrInvalid,
	}
	fmt.Println("os.ErrInvalid after Unwrap ->", pathErr1.Unwrap())

	pathErr2 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrPermission,
	}
	fmt.Println("os.ErrPermission after Unwrap ->", pathErr2.Unwrap())

	pathErr3 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrExist,
	}
	fmt.Println("os.ErrExist after Unwrap ->", pathErr3.Unwrap())

	pathErr4 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrNotExist,
	}
	fmt.Println("os.ErrNotExist after Unwrap ->", pathErr4.Unwrap())

	pathErr5 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrClosed,
	}
	fmt.Println("os.ErrClosed after Unwrap ->", pathErr5.Unwrap())

	pathErr6 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrNoDeadline,
	}
	fmt.Println("os.ErrNoDeadline after Unwrap ->", pathErr6.Unwrap())

	pathErr7 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  os.ErrDeadlineExceeded,
	}
	fmt.Println("os.ErrDeadlineExceeded after Unwrap ->", pathErr7.Unwrap())

	pathErr8 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  fmt.Errorf("found error: %w", errors.New("未知错误")),
	}
	fmt.Println("fmt.Errorf after Unwrap ->", pathErr8.Unwrap())

	pathErr9 := &fs.PathError{
		Op:   "open",
		Path: "myfile.txt",
		Err:  fmt.Errorf("found error1: %w, found error2:: %w", errors.New("未知错误1"), errors.New("未知错误2")),
	}
	fmt.Println("fmt.Errorf after Unwrap ->", pathErr9.Unwrap())
}

// Output:
//os.ErrInvalid after Unwrap -> invalid argument
//os.ErrPermission after Unwrap -> permission denied
//os.ErrExist after Unwrap -> file already exists
//os.ErrNotExist after Unwrap -> file does not exist
//os.ErrClosed after Unwrap -> file already closed
//os.ErrNoDeadline after Unwrap -> file type does not support deadline
//os.ErrDeadlineExceeded after Unwrap -> i/o timeout
//os.ErrDeadlineExceeded after Unwrap -> i/o timeout
//fmt.Errorf after Unwrap -> found error: 未知错误
//fmt.Errorf after Unwrap -> found error1: 未知错误1, found error2:: 未知错误2
```

### type ReadDirFS 

``` go 
type ReadDirFS interface {
	FS

    // ReadDir reads the named directory
	// and returns a list of directory entries sorted by filename.
	// ReadDir读取指定的目录并返回按文件名排序的目录条目列表。
	ReadDir(name string) ([]DirEntry, error)
}
```

ReadDirFS is the interface implemented by a file system that provides an optimized implementation of ReadDir.

​	ReadDirFS是由文件系统实现的接口，它提供了ReadDir方法的优化实现。

#### ReadDirFS My Example

```go
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

type myFS struct {
	fs.FS
}

func (mfs myFS) ReadDir(name string) ([]fs.DirEntry, error) {
	fmt.Printf("Reading directory: %s\n", name)
	return fs.ReadDir(mfs.FS, name)
}

func main() {
	mfs := myFS{os.DirFS("dir")}

	fmt.Printf("mfs=%#v\n", mfs) // mfs=main.myFS{FS:"."}

	dirEntries, err := fs.ReadDir(mfs, ".")
	if err != nil {
		log.Fatal(err)
	}

	// 说是dirEntry，实际也包含 file
	for _, dirEntry := range dirEntries {
		fmt.Println(dirEntry.Name())
	}

	fmt.Println("-----------------------------")
	dirEntries, err = fs.ReadDir(mfs, "subdir1")
	if err != nil {
		log.Fatal(err)
	}

	// 说是dirEntry，实际也包含 file
	for _, dirEntry := range dirEntries {
		fmt.Println(dirEntry.Name())
	}

}
// Output:
//mfs=main.myFS{FS:"dir"}
//Reading directory: .
//0.txt
//subdir1
//subdir2
//-----------------------------
//Reading directory: subdir1
//1.txt
```

### type ReadDirFile 

``` go 
type ReadDirFile interface {
	File

    // ReadDir reads the contents of the directory and returns
	// a slice of up to n DirEntry values in directory order.
	// Subsequent calls on the same file will yield further DirEntry values.
	//
	// If n > 0, ReadDir returns at most n DirEntry structures.
	// In this case, if ReadDir returns an empty slice, it will return
	// a non-nil error explaining why.
	// At the end of a directory, the error is io.EOF.
	// (ReadDir must return io.EOF itself, not an error wrapping io.EOF.)
	//
	// If n <= 0, ReadDir returns all the DirEntry values from the directory
	// in a single slice. In this case, if ReadDir succeeds (reads all the way
	// to the end of the directory), it returns the slice and a nil error.
	// If it encounters an error before the end of the directory,
	// ReadDir returns the DirEntry list read until that point and a non-nil error.
	// ReadDir读取目录的内容并以目录顺序返回高达n个DirEntry值的切片。
    // 后续对同一文件的调用将产生更多的DirEntry值。
    //
    // 如果n> 0，则ReadDir最多返回n个DirEntry结构。
    // 在这种情况下，如果ReadDir返回一个空切片，
    // 则它将返回一个非nil错误解释原因。
    // 在目录的末尾，错误是io.EOF。
    // (ReadDir必须返回io.EOF本身，而不是包装io.EOF的错误。)
    //
    // 如果n <= 0，
    // 则ReadDir在单个切片中返回目录中的所有DirEntry值。
    // 在这种情况下，如果ReadDir成功(一直读到目录的末尾)，
    // 它将返回该切片和一个nil错误。
    // 如果在目录的末尾之前遇到错误，
    // 则ReadDir返回读取到该点的DirEntry列表和一个非nil错误。
    ReadDir(n int) ([]DirEntry, error)
}
```

A ReadDirFile is a directory file whose entries can be read with the ReadDir method. Every directory file should implement this interface. (It is permissible for any file to implement this interface, but if so ReadDir should return an error for non-directories.)

​	ReadDirFile是一个可以使用ReadDir方法读取其条目的目录文件。每个目录文件都应实现此接口。(任何文件都可以实现此接口，但如果这样做，对于非目录，ReadDir应返回一个错误。)

#### ReadDirFile My Example

![image-20230825082407827](fs_img/image-20230825082407827.png)

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
)

type customReadDirFile struct {
	file *os.File
}

func (f *customReadDirFile) ReadDir(n int) ([]fs.DirEntry, error) {
	fmt.Println("using custom ReadDir...")
	return f.file.ReadDir(n)
}

func (f *customReadDirFile) Stat() (fs.FileInfo, error) {
	return f.file.Stat()
}

func (f *customReadDirFile) Read(data []byte) (int, error) {
	return f.file.Read(data)
}

func (f *customReadDirFile) Close() error {
	return f.file.Close()
}

func main() {
	file, err := os.Open("dir")
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}

	customRdf := &customReadDirFile{file: file}

	var fsrdf fs.ReadDirFile
	fsrdf = customRdf
	_, ok := fsrdf.(fs.ReadDirFile)
	fmt.Println("customReadDirFile类型是否实现了fs.ReadDirFile？", ok)

	fmt.Println("1 ------------------------------------")
	entries, err := customRdf.ReadDir(-1)
	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}

	fmt.Println("读取到的列表：")
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	fmt.Println("2 ------------------------------------")
	entries, err = customRdf.ReadDir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("读取到的列表：")
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	fmt.Println("3 ------------------------------------")
	entries, err = customRdf.ReadDir(1)

	fmt.Println("若想每次调用 ReadDir 都有值，则需要重新打开对应目录。例如这里的 dir目录！")

	if err != nil {
		fmt.Println("发生错误：", err)
		return
	}

	fmt.Println("读取到的列表：")
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}
// Output:
//customReadDirFile类型是否实现了fs.ReadDirFile？ true
//1 ------------------------------------
//using custom ReadDir...
//读取到的列表：
//0.txt
//subdir1
//subdir2
//2 ------------------------------------
//using custom ReadDir...
//读取到的列表：
//3 ------------------------------------
//using custom ReadDir...
//若想每次调用 ReadDir 都有值，则需要重新打开对应目录。例如这里的 dir目录！
//发生错误： EOF

```

### type ReadFileFS 

``` go 
type ReadFileFS interface {
	FS

    // ReadFile reads the named file and returns its contents.
	// A successful call returns a nil error, not io.EOF.
	// (Because ReadFile reads the whole file, the expected EOF
	// from the final Read is not treated as an error to be reported.)
	//
	// The caller is permitted to modify the returned byte slice.
	// This method should return a copy of the underlying data.
    // ReadFile 读取指定的文件并返回其内容。
    // 一次成功的调用返回一个nil错误，而不是io.EOF。
    // (因为ReadFile读取整个文件，因此从最后一次读取的EOF不被视为要报告的错误。)
    //
    // 调用者可以修改返回的字节切片。
    // 该方法应该返回底层数据的副本。
	ReadFile(name string) ([]byte, error)
}
```

ReadFileFS is the interface implemented by a file system that provides an optimized implementation of ReadFile.

​	`ReadFileFS` 是一个由文件系统实现的接口，该接口提供了 `ReadFile` 的优化实现。

ReadFileFS My Example

![image-20230825085347673](fs_img/image-20230825085347673.png)

```go
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type myFS struct {
	ifs fs.FS
}

func (mfs myFS) Open(name string) (fs.File, error) {
	return mfs.ifs.Open(name)
}

func (mfs myFS) ReadFile(name string) ([]byte, error) {
	f, err := mfs.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, 4096)
	n, err := f.Read(buf)
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, errors.New("没法发现数据")
	}

	return buf, nil
}

func main() {
	mfs := myFS{os.DirFS("dir")}

	data, err := mfs.ReadFile("0.txt")
	if err != nil {
		fmt.Println("发生错误:", err)
	}
	fmt.Println("内容是：", string(data))
	fmt.Println("----------------------------")

	data, err = mfs.ReadFile("1.txt")
	if err != nil {
		fmt.Println("发生错误:", err)
	}

	fmt.Println("内容是：", string(data))
	fmt.Println("----------------------------")

	data, err = mfs.ReadFile("subdir1/1.txt")
	if err != nil {
		fmt.Println("发生错误:", err)
	}
	fmt.Println("内容是：", string(data))
}
// Output:
//内容是： content0
//----------------------------
//发生错误: open 1.txt: The system cannot find the file specified.
//内容是：
//----------------------------
//内容是： content1
```

### type StatFS 

``` go 
type StatFS interface {
	FS

    // Stat returns a FileInfo describing the file.
	// If there is an error, it should be of type *PathError.
    // Stat 返回描述文件的FileInfo。
    // 如果发生错误，它应该是类型为*PathError。
	Stat(name string) (FileInfo, error)
}
```

A StatFS is a file system with a Stat method.

​	StatFS是一个具有Stat方法的文件系统。

#### StatFS My Example

![image-20230825111439914](fs_img/image-20230825111439914.png)

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
)

type myFS struct {
	ifs fs.FS
}

func (mfs myFS) Open(name string) (fs.File, error) {
	return mfs.ifs.Open(name)
}

func (mfs myFS) Stat(name string) (fs.FileInfo, error) {
	f, err := mfs.Open(name)
	if err != nil {
		return nil, &fs.PathError{
			Op:   "open",
			Path: name,
			Err:  err,
		}
	}
	fmt.Printf("f type is %T\n", f)
	defer f.Close()

	return f.Stat()
}

func main() {
	mfs := myFS{os.DirFS("dir")}

	var i fs.StatFS
	i = mfs
	imfs, ok := i.(myFS)
	if ok {
		fmt.Println("myFs实现了fs.StatFS")
	}

	filepaths := []string{
		"1.txt",
		"2.txt",
		"subdir1/1.txt",
		"subdir2/2.txt",
	}
	for _, filePath := range filepaths {
		fmt.Println(filePath, "--------------------")
		fileInfo, err := imfs.Stat(filePath)
		if err != nil {
			fmt.Printf("发生错误：%v 错误类型：%T\n", err, err)
		} else {
			fmt.Println("fileInfo.Name()=", fileInfo.Name())
			fmt.Println("fileInfo.Size()=", fileInfo.Size())
			fmt.Println("fileInfo.Mode()=", fileInfo.Mode())
			fmt.Println("fileInfo.ModTime()=", fileInfo.ModTime())
			fmt.Println("fileInfo.IsDir()=", fileInfo.IsDir())
			fmt.Println("fileInfo.Sys()=", fileInfo.Sys())
		}
	}
}

// Output:
//myFs实现了fs.StatFS
//1.txt --------------------
//发生错误：open 1.txt: open 1.txt: no such file or directory 错误类型：*fs.PathError
//2.txt --------------------
//发生错误：open 2.txt: open 2.txt: no such file or directory 错误类型：*fs.PathError
//subdir1/1.txt --------------------
//f type is *os.File
//fileInfo.Name()= 1.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-r--
//fileInfo.ModTime()= 2023-08-28 08:09:21.730031752 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()= &{2080 11702 1 33204 1000 1000 0 0 8 4096 8 {1693181480 390036480} {1693181361 730031752} {1693181361 730031752} [0 0 0]}
//subdir2/2.txt --------------------
//f type is *os.File
//fileInfo.Name()= 2.txt
//fileInfo.Size()= 8
//fileInfo.Mode()= -rw-rw-r--
//fileInfo.ModTime()= 2023-08-28 08:09:21.730031752 +0800 CST
//fileInfo.IsDir()= false
//fileInfo.Sys()= &{2080 11704 1 33204 1000 1000 0 0 8 4096 8 {1693181480 390036480} {1693181361 730031752} {1693181361 730031752} [0 0 0]}
```

### type SubFS 

``` go 
type SubFS interface {
	FS

    // Sub returns an FS corresponding to the subtree rooted at dir.
	// Sub 返回与dir根目录对应的FS。
	Sub(dir string) (FS, error)
}
```

A SubFS is a file system with a Sub method.

​	SubFS是一个具有Sub方法的文件系统。



#### SubFS My Example

![image-20230825111414158](fs_img/image-20230825111414158.png)

```go
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

type myFS struct {
	// 嵌入一个*os.FS字段
	ifs  fs.FS
	Root string
}

func NewMyFS(dir string) myFS {
	return myFS{ifs: os.DirFS(dir), Root: dir}
}

func (mfs myFS) Open(name string) (fs.File, error) {
	return mfs.ifs.Open(name)
}

func (mfs myFS) Sub(dir string) (fs.FS, error) {
	return NewMyFS(path.Join(mfs.Root, dir)), nil
}

func main() {
	mfs := NewMyFS("dir")
	subMfs, err := mfs.Sub("subdir1")
	if err != nil {
		fmt.Println("发生错误：", err)
	}

	fs.WalkDir(subMfs, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println("path=", path, "------------------------")
		fmt.Println("d.Name()=", d.Name())
		fmt.Println("d.IsDir()=", d.IsDir())
		info, _ := d.Info()
		fmt.Println("info.Name()=", info.Name())
		fmt.Println("info.Size()=", info.Size())
		fmt.Println("info.Mode()=", info.Mode())
		fmt.Println("info.ModTime()=", info.ModTime())
		fmt.Println("info.IsDir()=", info.IsDir())
		fmt.Printf("info.Sys()=%#v\n", info.Sys())
		return nil
	})
}

// Output:
//path= . ------------------------
//d.Name()= .
//d.IsDir()= true
//info.Name()= .
//info.Size()= 0
//info.Mode()= drwxrwxrwx
//info.ModTime()= 2023-08-25 11:05:52.7061956 +0800 CST
//info.IsDir()= true
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x10, CreationTime:syscall.Filetime{LowDateTime:0x63d1b05, HighDateTime:0x1d9d6ff}, LastAccessTime:syscall.
//path= 2.txt ------------------------
//d.Name()= 2.txt
//d.IsDir()= false
//info.Name()= 2.txt
//info.Size()= 8
//info.Mode()= -rw-rw-rw-
//info.ModTime()= 2023-08-25 11:05:52.7046265 +0800 CST
//info.IsDir()= false
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x39bed17a, HighDateTime:0x1d9d700}, LastAccessTime:syscall
//.Filetime{LowDateTime:0xec7ac79, HighDateTime:0x1d9d701}, LastWriteTime:syscall.Filetime{LowDateTime:0xec7ac79, HighDateTime:0x1d9d701}, FileSizeHigh:0x0, FileSizeLo
//w:0x8}
//path= 3.html ------------------------
//d.Name()= 3.html
//d.IsDir()= false
//info.Name()= 3.html
//info.Size()= 132
//info.Mode()= -rw-rw-rw-
//info.ModTime()= 2023-08-25 11:05:52.6546262 +0800 CST
//info.IsDir()= false
//info.Sys()=&syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x44a0928e, HighDateTime:0x1d9d700}, LastAccessTime:syscall
//.Filetime{LowDateTime:0xec00b56, HighDateTime:0x1d9d701}, LastWriteTime:syscall.Filetime{LowDateTime:0xec00b56, HighDateTime:0x1d9d701}, FileSizeHigh:0x0, FileSizeLo
//w:0x84}
```



### type WalkDirFunc 

``` go 
type WalkDirFunc func(path string, d DirEntry, err error) error
```

WalkDirFunc is the type of the function called by WalkDir to visit each file or directory.

​	WalkDirFunc 是WalkDir函数用来访问每个文件或目录的函数类型。

The path argument contains the argument to WalkDir as a prefix. That is, if WalkDir is called with root argument "dir" and finds a file named "a" in that directory, the walk function will be called with argument "dir/a".

​	`path` 实参包含WalkDir函数的实参作为前缀。也就是说，如果使用根实参"dir"调用WalkDir函数，并在该目录中找到名为"a"的文件，则遍历函数将使用实参"dir/a"进行调用。

The d argument is the fs.DirEntry for the named path.

​	`d`实参是具有 fs.DirEntry 的命名路径。

The error result returned by the function controls how WalkDir continues. If the function returns the special value SkipDir, WalkDir skips the current directory (path if d.IsDir() is true, otherwise path's parent directory). If the function returns the special value SkipAll, WalkDir skips all remaining files and directories. Otherwise, if the function returns a non-nil error, WalkDir stops entirely and returns that error.

​	该函数返回的错误结果控制WalkDir函数的继续。如果函数返回特殊值`SkipDir`，则WalkDir函数跳过当前目录（如果 `d.IsDir()`为true，则为`path`，否则为`path`的父目录）。如果函数返回特殊值`SkipAll`，则WalkDir函数跳过所有剩余文件和目录。否则，如果该函数返回非nil错误，则WalkDir函数完全停止并返回该错误。

The err argument reports an error related to path, signaling that WalkDir will not walk into that directory. The function can decide how to handle that error; as described earlier, returning the error will cause WalkDir to stop walking the entire tree.

​	`err`参数报告与`path`相关的错误，表示WalkDir函数不会遍历该目录。该函数可以决定如何处理该错误；如前所述，返回错误将导致WalkDir函数停止遍历整个树。

WalkDir calls the function with a non-nil err argument in two cases.

​	WalkDir函数在两种情况下使用非nil `err`实参调用函数。

First, if the initial fs.Stat on the root directory fails, WalkDir calls the function with path set to root, d set to nil, and err set to the error from fs.Stat.

​	第一，如果根目录上的初始fs.Stat失败，WalkDir函数调用该函数时，`path`设置为`root`，`d`设置为nil，`err`设置为fs.Stat的错误。

Second, if a directory's ReadDir method fails, WalkDir calls the function with path set to the directory's path, d set to an fs.DirEntry describing the directory, and err set to the error from ReadDir. In this second case, the function is called twice with the path of the directory: the first call is before the directory read is attempted and has err set to nil, giving the function a chance to return SkipDir or SkipAll and avoid the ReadDir entirely. The second call is after a failed ReadDir and reports the error from ReadDir. (If ReadDir succeeds, there is no second call.)

​	第二，如果一个目录的ReadDir方法失败，WalkDir函数调用该函数，`path`设置为该目录的路径，`d`设置为描述该目录的fs.DirEntry，`err`设置为ReadDir的错误。 在这第二种情况下，该函数被调用两次，`path`为该目录：第一次调用是在试图读取目录之前，`err`设置为nil，给该函数一个机会返回`SkipDir`或`SkipAll`，并完全避免ReadDir。第二次调用是在ReadDir失败之后，并报告ReadDir的错误(如果ReadDir成功，则没有第二次调用)。

The differences between WalkDirFunc compared to filepath.WalkFunc are:

​	WalkDirFunc与filepath.WalkFunc的不同之处在于：

- The second argument has type fs.DirEntry instead of fs.FileInfo.

- 第二个实参的类型为fs.DirEntry，而不是fs.FileInfo。 
- The function is called before reading a directory, to allow SkipDir or SkipAll to bypass the directory read entirely or skip all remaining files and directories respectively.
- 该函数在读取目录之前被调用，以允许`SkipDir`或`SkipAll`完全跳过目录读取或跳过所有剩余的文件和目录。 
- If a directory read fails, the function is called a second time for that directory to report the error.
- 如果目录读取失败，该函数将被第二次被调用，以报告该目录的错误。



#### WalkDirFunc My Example

参见 [func WalkDir](#func-walkdir)