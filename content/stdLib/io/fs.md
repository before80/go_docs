+++
title = "fs"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# fs

https://pkg.go.dev/io/fs@go1.20.1

​	fs包定义了与文件系统交互的基本接口。文件系统可以由操作系统提供，也可以由其他包提供。

## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/io/fs/fs.go;l=136)

``` go 
var (
	ErrInvalid    = errInvalid()    // "参数无效"
	ErrPermission = errPermission() // "权限被拒绝"
	ErrExist      = errExist()      // "文件已经存在"
	ErrNotExist   = errNotExist()   // "文件不存在"
	ErrClosed     = errClosed()     // "文件已经关闭"
)
```

​	通用的文件系统错误。可以使用errors.Is将文件系统返回的错误与这些错误进行比较。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/io/fs/walk.go;l=20)

``` go 
var SkipAll = errors.New("skip everything and stop the walk")
```

​	SkipAll是从WalkDirFuncs返回的值，用于指示所有剩余的文件和目录都应该被跳过。任何函数都不会将其作为错误返回。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/io/fs/walk.go;l=15)

``` go 
var SkipDir = errors.New("skip this directory")
```

​	SkipDir是从WalkDirFuncs返回的值，用于指示应该跳过调用中指定的目录。任何函数都不会将其作为错误返回。

## 函数

#### func Glob 

``` go 
func Glob(fsys FS, pattern string) (matches []string, err error)
```

​	Glob函数返回与pattern匹配的所有文件的名称，如果没有匹配的文件，则返回nil。模式的语法与path.Match中的语法相同。模式可以描述分层名称，例如`usr/*/bin/ed`。

​	Glob函数忽略文件系统错误，例如读取目录的I/O错误。唯一可能返回的错误是path.ErrBadPattern，报告模式格式不正确。

​	如果fs实现了GlobFS，则Glob函数调用fs.Glob。否则，Glob函数使用ReadDir遍历目录树并查找模式匹配项。

##### Glob My Example

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



#### func ReadFile 

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

​	ReadFile 函数从文件系统fs中读取指定的文件并返回其内容。成功调用返回nil错误，而不是io.EOF。(因为ReadFile读取整个文件，因此不会将最终的EOF视为要报告的错误。)

​	如果fs实现了ReadFileFS，则ReadFile调用fs.ReadFile。否则，ReadFile调用fs.Open并在返回的文件上使用Read和Close。

##### ReadFile My Example

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

#### func ValidPath 

``` go 
func ValidPath(name string) bool
```

ValidPath reports whether the given path name is valid for use in a call to Open.

ValidPath 报告给定的路径名称是否可以在调用 Open 时使用。

ValidPath 函数返回给定路径名在调用 Open 时是否有效。

Path names passed to open are UTF-8-encoded, unrooted, slash-separated sequences of path elements, like "x/y/z". Path names must not contain an element that is "." or ".." or the empty string, except for the special case that the root directory is named ".". Paths must not start or end with a slash: "/x" and "x/" are invalid.

传递给open的路径名是UTF-8编码的、无根的、斜线分隔的路径元素序列，如 "x/y/z"。路径名不能包含". "或". "或空字符串的元素，除了根目录被命名为". "的特殊情况。路径不能以斜线开始或结束："/x "和 "x/"是无效的。

传递给 Open 的路径名是以 UTF-8 编码的、未根化的、斜杠分隔的路径元素序列，例如 "x/y/z"。路径名不得包含 "." 或 ".." 或空字符串，但根目录的特殊情况是命名为 "."。路径不能以斜杠开头或结尾，即 "/x" 和 "x/" 是无效的。

Note that paths are slash-separated on all systems, even Windows. Paths containing other characters such as backslash and colon are accepted as valid, but those characters must never be interpreted by an FS implementation as path element separators.

请注意，在所有的系统上，甚至是Windows，路径都是以斜线分隔的。含有反斜杠和冒号等其他字符的路径可以被接受为有效，但这些字符决不能被FS实现解释为路径元素分隔符。

请注意，路径在所有系统上都是以斜杠分隔的，即使在 Windows 上也是如此。包含反斜杠和冒号等其他字符的路径被接受为有效，但这些字符绝不能被 FS 实现解释为路径元素分隔符。

##### My Example

```go

```

#### func WalkDir 

``` go 
func WalkDir(fsys FS, root string, fn WalkDirFunc) error
```

WalkDir walks the file tree rooted at root, calling fn for each file or directory in the tree, including root.

WalkDir行走以根为根的文件树，为树中的每个文件或目录调用fn，包括根。

WalkDir 遍历以 root 为根的文件树，在树中的每个文件或目录(包括 root)上调用 fn。

All errors that arise visiting files and directories are filtered by fn: see the fs.WalkDirFunc documentation for details.

所有访问文件和目录出现的错误都由fn过滤：详情请参见fs.WalkDirFunc文档。

fn 过滤了遍历文件和目录时出现的所有错误：详见 fs.WalkDirFunc 文档。

The files are walked in lexical order, which makes the output deterministic but requires WalkDir to read an entire directory into memory before proceeding to walk that directory.

文件是按词法顺序走的，这使得输出是确定的，但要求WalkDir在继续走该目录之前将整个目录读入内存。

文件以字典序遍历，这使输出是确定性的，但需要在继续遍历该目录之前将整个目录读入内存。

WalkDir does not follow symbolic links found in directories, but if root itself is a symbolic link, its target will be walked.

WalkDir不跟踪在目录中发现的符号链接，但是如果root本身是一个符号链接，它的目标将被步行。

WalkDir 不会遵循目录中发现的符号链接，但如果 root 本身是符号链接，则会遍历其目标。

##### My Example

```go

```

##### WalkDir Example

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
    // Name 返回该条目描述的文件(或子目录)的名称。
    // 该名称仅为路径的最后一个元素(基本名称)，而不是整个路径。
    // 例如，Name 将返回"hello.go"而不是"home/gopher/hello.go"。
	Name() string

    // IsDir 报告该条目是否描述一个目录。
	IsDir() bool

    // Type 返回该条目的类型位。
    // 类型位是 FileMode 常量的子集，
    // 是由 FileMode.Type 方法返回的常量之一。
	Type() FileMode

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

DirEntry 是一个从目录中读取的条目(使用 ReadDir 函数或 ReadDirFile 的 ReadDir 方法)。

DirEntry 是从目录中读取的一个条目(使用 ReadDir 函数或 ReadDirFile 的 ReadDir 方法)。

#### func FileInfoToDirEntry  <- go1.17

``` go 
func FileInfoToDirEntry(info FileInfo) DirEntry
```

FileInfoToDirEntry returns a DirEntry that returns information from info. If info is nil, FileInfoToDirEntry returns nil.

FileInfoToDirEntry 返回一个 DirEntry，它从 info 中返回信息。如果info是nil，FileInfoToDirEntry返回nil。

FileInfoToDirEntry 返回一个从 info 中获取信息的 DirEntry。如果 info 为 nil，则 FileInfoToDirEntry 返回 nil。

##### My Example

```go

```

#### func ReadDir 

``` go 
func ReadDir(fsys FS, name string) ([]DirEntry, error)
```

ReadDir reads the named directory and returns a list of directory entries sorted by filename.

ReadDir读取命名的目录并返回一个按文件名排序的目录条目列表。

If fs implements ReadDirFS, ReadDir calls fs.ReadDir. Otherwise ReadDir calls fs.Open and uses ReadDir and Close on the returned file.

如果fs实现了ReadDirFS，ReadDir调用fs.ReadDir。否则ReadDir调用fs.Open并对返回的文件使用ReadDir和Close。

如果 fs 实现了 ReadDirFS，则 ReadDir 调用 fs.ReadDir。否则，ReadDir 调用 fs.Open 并使用返回的文件上的 ReadDir 和 Close。

##### My Example

```go

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

一个FS提供了对一个分层文件系统的访问。

FS 提供对分层文件系统的访问。

The FS interface is the minimum implementation required of the file system. A file system may implement additional interfaces, such as ReadFileFS, to provide additional or optimized functionality.

FS接口是文件系统所需的最小实现。一个文件系统可以实现额外的接口，如ReadFileFS，以提供额外的或优化的功能。

FS 接口是文件系统所需的最小实现。文件系统可能会实现其他接口(如 ReadFileFS)以提供附加或优化的功能。

#### func Sub 

``` go 
func Sub(fsys FS, dir string) (FS, error)
```

Sub returns an FS corresponding to the subtree rooted at fsys's dir.

Sub返回一个对应于以fsys的dir为根的子树的FS。

Sub 返回一个对应于以 fsys 的 dir 为根的子树的 FS。

If dir is ".", Sub returns fsys unchanged. Otherwise, if fs implements SubFS, Sub returns fsys.Sub(dir). Otherwise, Sub returns a new FS implementation sub that, in effect, implements sub.Open(name) as fsys.Open(path.Join(dir, name)). The implementation also translates calls to ReadDir, ReadFile, and Glob appropriately.

如果dir是"."，Sub返回fsys而不改变。否则，如果fs实现了SubFS，Sub返回fsys.Sub(dir)。否则，Sub返回一个新的FS实现sub，实际上，它将sub.Open(name)实现为fsys.Open(path.Join(dir, name))。该实现还适当地翻译了对 ReadDir、ReadFile 和 Glob 的调用。

如果 dir 为"."，则 Sub 返回未更改的 fsys。否则，如果 fs 实现了 SubFS，则 Sub 返回 fsys.Sub(dir)。否则，Sub 返回一个新的 FS 实现 sub，该实现实际上将 sub.Open(name) 实现为 fsys.Open(path.Join(dir, name))。该实现还适当地翻译对 ReadDir、ReadFile 和 Glob 的调用。

Note that Sub(os.DirFS("/"), "prefix") is equivalent to os.DirFS("/prefix") and that neither of them guarantees to avoid operating system accesses outside "/prefix", because the implementation of os.DirFS does not check for symbolic links inside "/prefix" that point to other directories. That is, os.DirFS is not a general substitute for a chroot-style security mechanism, and Sub does not change that fact.

注意Sub(os.DirFS("/"), "prefix")等同于os.DirFS("/prefix")，它们都不能保证避免操作系统对"/prefix "之外的访问，因为os.DirFS的实现并不检查"/prefix "内指向其他目录的符号链接。也就是说，os.DirFS并不是chroot式安全机制的一般替代品，Sub并不能改变这一事实。

请注意，Sub(os.DirFS("/"), "prefix") 等同于 os.DirFS("/prefix")，并且它们都不能保证避免超出"/prefix"范围的操作系统访问，因为 os.DirFS 的实现不检查指向其他目录的"/prefix"内部符号链接。也就是说，os.DirFS 不是 chroot 样式安全机制的通用替代品，Sub 也不改变这个事实。

##### My Example

```go

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

一个文件提供对单个文件的访问。File接口是文件所需的最小实现。目录文件也应该实现ReadDirFile。一个文件可以实现io.ReaderAt或io.Seeker作为优化。

File接口提供对单个文件的访问。File接口是文件所需的最小实现。目录文件还应该实现ReadDirFile。文件可以实现io.ReaderAt或io.Seeker作为优化。

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

FileInfo描述了一个文件，并由Stat返回。

FileInfo接口描述文件并由Stat返回。

#### func Stat 

``` go 
func Stat(fsys FS, name string) (FileInfo, error)
```

Stat returns a FileInfo describing the named file from the file system.

Stat返回一个描述文件系统中的命名文件的FileInfo。

Stat从文件系统返回描述命名文件的FileInfo。

If fs implements StatFS, Stat calls fs.Stat. Otherwise, Stat opens the file to stat it.

如果fs实现了StatFS，Stat调用fs.Stat。否则，Stat打开文件以进行统计。

如果fs实现了StatFS，则Stat调用fs.Stat。否则，Stat打开文件以获取其状态。

##### My Example

```go

```

### type FileMode 

``` go 
type FileMode uint32
```

A FileMode represents a file's mode and permission bits. The bits have the same definition on all systems, so that information about files can be moved from one system to another portably. Not all bits apply to all systems. The only required bit is ModeDir for directories.

一个FileMode代表一个文件的模式和权限位。这些位在所有系统上都有相同的定义，因此关于文件的信息可以从一个系统移植到另一个系统。不是所有的位都适用于所有的系统。唯一需要的位是目录的ModeDir。

FileMode表示文件的模式和权限位。这些位在所有系统上具有相同的定义，以便可以在不同系统之间可移植地移动文件信息。并非所有位都适用于所有系统。唯一必需的位是ModeDir，适用于目录。

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

定义的文件模式位是FileMode中最重要的位。九个最不重要的位是标准的Unix rwxrwxrwx权限。这些位的值应该被认为是公共API的一部分，可以在线程协议或磁盘表示法中使用：它们不能被改变，尽管可能会添加新的位。

定义的文件模式位是 FileMode 的最高位。最低的九位是标准 Unix rwxrwxrwx 权限位。这些位的值应该被视为公共 API 的一部分，可以在传输协议或磁盘表示中使用：它们不得更改，但可以添加新的位。

#### (FileMode) IsDir 

``` go 
func (m FileMode) IsDir() bool
```

IsDir reports whether m describes a directory. That is, it tests for the ModeDir bit being set in m.

IsDir报告m是否描述了一个目录。也就是说，它测试ModeDir位是否被设置在m中。

IsDir 报告 m 是否描述一个目录。也就是说，它测试 ModeDir 位是否在 m 中被设置。

##### My Example

```go

```

#### (FileMode) IsRegular 

``` go 
func (m FileMode) IsRegular() bool
```

IsRegular reports whether m describes a regular file. That is, it tests that no mode type bits are set.

IsRegular报告m是否描述了一个常规文件。也就是说，它测试没有模式类型位被设置。

IsRegular 报告 m 是否描述一个普通文件。也就是说，它测试是否没有设置任何模式类型位。

##### My Example

```go

```

#### (FileMode) Perm 

``` go 
func (m FileMode) Perm() FileMode
```

Perm returns the Unix permission bits in m (m & ModePerm).

Perm返回m中的Unix权限位(m & ModePerm)。

Perm 返回 m 中的 Unix 权限位(m＆ModePerm)。

##### My Example

```go

```

#### (FileMode) String 

``` go 
func (m FileMode) String() string
```

##### My Example

```go

```

#### (FileMode) Type 

``` go 
func (m FileMode) Type() FileMode
```

Type returns type bits in m (m & ModeType).

Type 返回m中的类型位(m & ModeType)。

Type 返回 m 中的类型位(m＆ModeType)。

##### My Example

```go

```

### type GlobFS 

``` go 
type GlobFS interface {
	FS

    // Glob 返回与 pattern 匹配的所有文件的名称，
    // 提供了顶级 Glob 函数的实现。
	Glob(pattern string) ([]string, error)
}
```

A GlobFS is a file system with a Glob method.

GlobFS是一个具有Glob方法的文件系统。

GlobFS 是具有 Glob 方法的文件系统。

##### My Example

```go

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

PathError记录了一个错误以及导致该错误的操作和文件路径。

PathError 记录了一个错误以及导致该错误的操作和文件路径。

#### (*PathError) Error 

``` go 
func (e *PathError) Error() string
```

##### My Example

```go

```

#### (*PathError) Timeout 

``` go 
func (e *PathError) Timeout() bool
```

Timeout reports whether this error represents a timeout.

Timeout 报告这个错误是否代表超时。

Timeout报告此错误是否表示超时。

#### (*PathError) Unwrap 

``` go 
func (e *PathError) Unwrap() error
```

##### My Example

```go

```

### type ReadDirFS 

``` go 
type ReadDirFS interface {
	FS

	// ReadDir读取指定的目录并返回按文件名排序的目录条目列表。
	ReadDir(name string) ([]DirEntry, error)
}
```

ReadDirFS is the interface implemented by a file system that provides an optimized implementation of ReadDir.

ReadDirFS是由文件系统实现的接口，它提供了ReadDir的优化实现。

ReadDirFS是由提供了ReadDir的文件系统所实现的接口。

##### My Example

```go

```

### type ReadDirFile 

``` go 
type ReadDirFile interface {
	File

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

ReadDirFile是一个目录文件，其条目可以用ReadDir方法读取。每个目录文件都应该实现这个接口。(任何文件都可以实现这个接口，但如果这样的话，ReadDir应该对非目录文件返回一个错误。)

ReadDirFile是一个可以使用ReadDir方法读取其条目的目录文件。每个目录文件都应实现此接口。(任何文件都可以实现此接口，但如果这样做，对于非目录，ReadDir应返回一个错误。)

##### My Example

```go

```

### type ReadFileFS 

``` go 
type ReadFileFS interface {
	FS

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

ReadFileFS是由文件系统实现的接口，它提供了ReadFile的优化实现。

ReadFileFS是一个文件系统，它提供了ReadFile的优化实现。

##### My Example

```go

```

### type StatFS 

``` go 
type StatFS interface {
	FS

    // Stat 返回描述文件的FileInfo。
    // 如果发生错误，它应该是类型为*PathError。
	Stat(name string) (FileInfo, error)
}
```

A StatFS is a file system with a Stat method.

一个StatFS是一个具有Stat方法的文件系统。

​	StatFS是一个具有Stat方法的文件系统。

##### My Example

```go

```

### type SubFS 

``` go 
type SubFS interface {
	FS

	// Sub 返回与dir根目录对应的FS。
	Sub(dir string) (FS, error)
}
```

A SubFS is a file system with a Sub method.

一个SubFS是一个具有Sub方法的文件系统。

SubFS是一个具有Sub方法的文件系统。

### type WalkDirFunc 

``` go 
type WalkDirFunc func(path string, d DirEntry, err error) error
```

WalkDirFunc is the type of the function called by WalkDir to visit each file or directory.

WalkDirFunc是由WalkDir调用的访问每个文件或目录的函数的类型。

WalkDirFunc是WalkDir用来访问每个文件或目录的函数类型。

The path argument contains the argument to WalkDir as a prefix. That is, if WalkDir is called with root argument "dir" and finds a file named "a" in that directory, the walk function will be called with argument "dir/a".

path参数包含作为前缀的WalkDir的参数。也就是说，如果用根参数 "dir "调用WalkDir，并在该目录中找到一个名为 "a "的文件，将用参数 "dir/a "调用Walk函数。

path参数包含WalkDir的参数作为前缀。也就是说，如果使用根参数"dir"调用WalkDir，并在该目录中找到名为"a"的文件，则遍历函数将使用参数"dir/a"进行调用。

The d argument is the fs.DirEntry for the named path.

d参数是命名路径的fs.DirEntry。

d参数是具有fs.DirEntry的命名路径。

The error result returned by the function controls how WalkDir continues. If the function returns the special value SkipDir, WalkDir skips the current directory (path if d.IsDir() is true, otherwise path's parent directory). If the function returns the special value SkipAll, WalkDir skips all remaining files and directories. Otherwise, if the function returns a non-nil error, WalkDir stops entirely and returns that error.

该函数返回的错误结果控制WalkDir如何继续。如果函数返回特殊值SkipDir，WalkDir将跳过当前目录(如果d.IsDir()为真，则为path，否则为path的父目录)。如果函数返回特殊值SkipAll，WalkDir将跳过所有剩余的文件和目录。否则，如果函数返回一个非零的错误，WalkDir完全停止并返回该错误。

函数返回的错误结果控制WalkDir的继续。如果函数返回特殊值SkipDir，则WalkDir跳过当前目录(如果d.IsDir()为true，则为路径，否则为路径的父目录)。如果函数返回特殊值SkipAll，则WalkDir跳过所有剩余文件和目录。否则，如果函数返回非nil错误，则WalkDir完全停止并返回该错误。

The err argument reports an error related to path, signaling that WalkDir will not walk into that directory. The function can decide how to handle that error; as described earlier, returning the error will cause WalkDir to stop walking the entire tree.

err参数报告一个与路径有关的错误，表示WalkDir不会进入该目录。该函数可以决定如何处理该错误；如前所述，返回该错误将导致WalkDir停止行走整个树。

err参数报告与路径相关的错误，表示WalkDir不会遍历该目录。函数可以决定如何处理该错误；如前所述，返回错误将导致WalkDir停止遍历整个树。

WalkDir calls the function with a non-nil err argument in two cases.

在两种情况下，WalkDir用一个非零的err参数调用该函数。

WalkDir在两种情况下使用非nil err参数调用函数。

First, if the initial fs.Stat on the root directory fails, WalkDir calls the function with path set to root, d set to nil, and err set to the error from fs.Stat.

首先，如果根目录上的初始fs.Stat失败，WalkDir调用该函数时，路径设置为root，d设置为nil，err设置为fs.Stat的错误。

首先，如果根目录的fs.Stat失败，则WalkDir使用path设置为根，d设置为nil，并使用从fs.Stat返回的错误设置err调用函数。

Second, if a directory's ReadDir method fails, WalkDir calls the function with path set to the directory's path, d set to an fs.DirEntry describing the directory, and err set to the error from ReadDir. In this second case, the function is called twice with the path of the directory: the first call is before the directory read is attempted and has err set to nil, giving the function a chance to return SkipDir or SkipAll and avoid the ReadDir entirely. The second call is after a failed ReadDir and reports the error from ReadDir. (If ReadDir succeeds, there is no second call.)

第二，如果一个目录的ReadDir方法失败，WalkDir调用该函数，path设置为该目录的路径，d设置为描述该目录的fs.DirEntry，err设置为ReadDir的错误。 在这第二种情况下，该函数被调用两次，路径为该目录：第一次调用是在试图读取目录之前，err设置为nil，给该函数一个机会返回SkipDir或SkipAll，完全避免ReadDir。第二次调用是在ReadDir失败之后，并报告ReadDir的错误(如果ReadDir成功，则没有第二次调用)。

其次，如果目录的ReadDir方法失败，则WalkDir使用path设置为目录的路径，d设置为描述目录的fs.DirEntry，并使用从ReadDir返回的错误设置err调用函数。在第二种情况下，该函数使用目录的路径两次进行调用：第一次调用在尝试读取目录之前进行，并将err设置为nil，给函数一次机会返回SkipDir或SkipAll，并完全避免ReadDir。第二次调用是在ReadDir失败之后，报告ReadDir的错误。(如果ReadDir成功，则没有第二次调用。)

The differences between WalkDirFunc compared to filepath.WalkFunc are:

与filepath.WalkFunc相比，WalkDirFunc的不同之处在于。

WalkDirFunc相对于filepath.WalkFunc的区别是：

WalkDirFunc与filepath.WalkFunc的不同之处在于：

- The second argument has type fs.DirEntry instead of fs.FileInfo. 第二个参数的类型是fs.DirEntry而不是fs.FileInfo。 第二个参数的类型为fs.DirEntry，而不是fs.FileInfo。 
- The function is called before reading a directory, to allow SkipDir or SkipAll to bypass the directory read entirely or skip all remaining files and directories respectively.该函数在读取目录之前被调用，以允许SkipDir或SkipAll完全绕过目录读取或分别跳过所有剩余的文件和目录。 函数在读取目录之前调用，以允许SkipDir或SkipAll完全跳过目录读取或跳过所有剩余的文件和目录。 
- If a directory read fails, the function is called a second time for that directory to report the error.如果目录读取失败，该函数将被第二次调用，以报告该目录的错误。如果目录读取失败，则会为该目录再次调用该函数以报告错误。



##### My Example

```go

```