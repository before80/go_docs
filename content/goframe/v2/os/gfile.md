+++
title = "gfile"
date = 2024-03-21T17:55:29+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gfile](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gfile)

Package gfile provides easy-to-use operations for file system.

​	软件包 gfile 为文件系统提供了简单易用的操作。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gfile/gfile.go#L23)

```go
const (
	// Separator for file system.
	// It here defines the separator as variable
	// to allow it modified by developer if necessary.
	Separator = string(filepath.Separator)

	// DefaultPermOpen is the default perm for file opening.
	DefaultPermOpen = os.FileMode(0666)

	// DefaultPermCopy is the default perm for file/folder copy.
	DefaultPermCopy = os.FileMode(0755)
)
```

## 变量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gfile/gfile_contents.go#L17)

```go
var (
	// DefaultReadBuffer is the buffer size for reading file content.
	DefaultReadBuffer = 1024
)
```

## 函数

#### func Abs

```go
func Abs(path string) string
```

Abs returns an absolute representation of path. If the path is not absolute it will be joined with the current working directory to turn it into an absolute path. The absolute path name for a given file is not guaranteed to be unique. Abs calls Clean on the result.

​	Abs 返回路径的绝对表示形式。如果路径不是绝对路径，它将与当前工作目录联接，以将其转换为绝对路径。不保证给定文件的绝对路径名是唯一的。Abs 在结果上调用 Clean。

##### Example

``` go
```

#### func Basename

```go
func Basename(path string) string
```

Basename returns the last element of path, which contains file extension. Trailing path separators are removed before extracting the last element. If the path is empty, Base returns “.”. If the path consists entirely of separators, Basename returns a single separator.

​	Basename 返回 path 的最后一个元素，其中包含文件扩展名。在提取最后一个元素之前，将删除尾随路径分隔符。如果路径为空，则 Base 返回“.”。如果路径完全由分隔符组成，则 Basename 将返回单个分隔符。

Example: Basename("/var/www/file.js") -> file.js Basename(“file.js”) -> file.js

​	示例：basename（“/var/www/file.js”） -> file.js basename（“file.js”） -> file.js

##### Example

``` go
```

#### func Chdir

```go
func Chdir(dir string) (err error)
```

Chdir changes the current working directory to the named directory. If there is an error, it will be of type *PathError.

​	Chdir 将当前工作目录更改为命名目录。如果存在错误，则其类型为 *PathError。

##### Example

``` go
```

#### func Chmod

```go
func Chmod(path string, mode os.FileMode) (err error)
```

Chmod is alias of os.Chmod. See os.Chmod.

​	Chmod 是 os 的别名。Chmod的。请参见 os。Chmod的。

##### Example

``` go
```

#### func Copy

```go
func Copy(src string, dst string, option ...CopyOption) error
```

Copy file/directory from `src` to `dst`.

​	将文件/目录从 `src` 复制到 `dst` 。

If `src` is file, it calls CopyFile to implements copy feature, or else it calls CopyDir.

​	如果 `src` 是 file，则调用 CopyFile 来实现复制功能，否则调用 CopyDir。

If `src` is file, but `dst` already exists and is a folder, it then creates a same name file of `src` in folder `dst`.

​	如果 `src` 是 file，但 `dst` 已经存在并且是一个文件夹，则它会创建一个与 `src` in 文件夹 同名的文件 `dst` 。

Eg: Copy("/tmp/file1", “/tmp/file2”) => /tmp/file1 copied to /tmp/file2 Copy("/tmp/dir1", “/tmp/dir2”) => /tmp/dir1 copied to /tmp/dir2 Copy("/tmp/file1", “/tmp/dir2”) => /tmp/file1 copied to /tmp/dir2/file1 Copy("/tmp/dir1", “/tmp/file2”) => error

​	例如： Copy（“/tmp/file1”， “/tmp/file2”） => /tmp/file1 复制到 /tmp/file2 复制（“/tmp/dir1”， “/tmp/dir2”） => /tmp/dir1 复制到 /tmp/dir2 复制（“/tmp/file1”， “/tmp/dir2”） => /tmp/file1 复制到 /tmp/dir2/file1 复制（“/tmp/dir1”， “/tmp/file2”） => error

##### Example

``` go
```

#### func CopyDir

```go
func CopyDir(src string, dst string, option ...CopyOption) (err error)
```

CopyDir recursively copies a directory tree, attempting to preserve permissions.

​	CopyDir 以递归方式复制目录树，尝试保留权限。

Note that, the Source directory must exist and symlinks are ignored and skipped.

​	请注意，源目录必须存在，符号链接将被忽略和跳过。

#### func CopyFile

```go
func CopyFile(src, dst string, option ...CopyOption) (err error)
```

CopyFile copies the contents of the file named `src` to the file named by `dst`. The file will be created if it does not exist. If the destination file exists, all it’s contents will be replaced by the contents of the source file. The file mode will be copied from the source and the copied data is synced/flushed to stable storage. Thanks: https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04

​	CopyFile 将命名 `src` 的文件的内容复制到以 `dst` 命名的文件。如果文件不存在，则将创建该文件。如果目标文件存在，则其所有内容都将替换为源文件的内容。文件模式将从源复制，复制的数据将同步/刷新到稳定存储。谢谢： https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04

#### func Create

```go
func Create(path string) (*os.File, error)
```

Create creates a file with given `path` recursively. The parameter `path` is suggested to be absolute path.

​	Create 以递归方式创建一个给定 `path` 的文件。建议该参数 `path` 为绝对路径。

##### Example

``` go
```

#### func Dir

```go
func Dir(path string) string
```

Dir returns all but the last element of path, typically the path’s directory. After dropping the final element, Dir calls Clean on the path and trailing slashes are removed. If the `path` is empty, Dir returns “.”. If the `path` is “.”, Dir treats the path as current working directory. If the `path` consists entirely of separators, Dir returns a single separator. The returned path does not end in a separator unless it is the root directory.

​	Dir 返回除 path 的最后一个元素之外的所有元素，通常是路径的目录。删除最后一个元素后，Dir 在路径上调用 Clean，并删除尾部斜杠。如果 为 `path` 空，则 Dir 返回 “.”。如果为 `path` “.”，则 Dir 将路径视为当前工作目录。如果 完全 `path` 由分隔符组成，则 Dir 返回单个分隔符。返回的路径不会以分隔符结尾，除非它是根目录。

Example: Dir("/var/www/file.js") -> “/var/www” Dir(“file.js”) -> “.”

​	示例： dir（“/var/www/file.js”） -> “/var/www” dir（“file.js”） -> “.”

##### Example

``` go
```

#### func DirNames

```go
func DirNames(path string) ([]string, error)
```

DirNames returns sub-file names of given directory `path`. Note that the returned names are NOT absolute paths.

​	DirNames 返回给定目录的子文件名 `path` 。请注意，返回的名称不是绝对路径。

##### Example

``` go
```

#### func Exists

```go
func Exists(path string) bool
```

Exists checks whether given `path` exist.

​	存在检查给定 `path` 是否存在。

##### Example

``` go
```

#### func Ext

```go
func Ext(path string) string
```

Ext returns the file name extension used by path. The extension is the suffix beginning at the final dot in the final element of path; it is empty if there is no dot. Note: the result contains symbol ‘.’.

​	Ext 返回 path 使用的文件扩展名。扩展名是从路径的最后一个元素中的最后一个点开始的后缀;如果没有点，则为空。注意：结果包含符号“.”。

Example: Ext(“main.go”) => .go Ext(“api.json”) => .json

​	示例：ext（“main.go”） => .go ext（“api.json”） => .json

##### Example

``` go
```

#### func ExtName

```go
func ExtName(path string) string
```

ExtName is like function Ext, which returns the file name extension used by path, but the result does not contain symbol ‘.’.

​	ExtName 类似于函数 Ext，它返回 path 使用的文件扩展名，但结果不包含符号“.”。

Example: ExtName(“main.go”) => go ExtName(“api.json”) => json

​	示例：ExtName（“main.go”） => go ExtName（“api.json”） => json

##### Example

``` go
```

#### func FormatSize

```go
func FormatSize(raw int64) string
```

FormatSize formats size `raw` for more manually readable.

​	FormatSize 格式大小 `raw` ，以便手动读取。

##### Example

``` go
```

#### func GetBytes

```go
func GetBytes(path string) []byte
```

GetBytes returns the file content of `path` as []byte. It returns nil if it fails reading.

​	GetBytes 以 []byte 的形式返回文件 `path` 内容。如果读取失败，则返回 nil。

##### Example

``` go
```

#### func GetBytesByTwoOffsets

```go
func GetBytesByTwoOffsets(reader io.ReaderAt, start int64, end int64) []byte
```

GetBytesByTwoOffsets returns the binary content as []byte from `start` to `end`. Note: Returned value does not contain the character of the last position, which means it returns content range as [start, end).

​	GetBytesByTwoOffsets 将二进制内容作为 []byte from `start` to `end` 返回。注意：返回值不包含最后一个位置的字符，这意味着它返回内容范围为 [start， end]。

#### func GetBytesByTwoOffsetsByPath

```go
func GetBytesByTwoOffsetsByPath(path string, start int64, end int64) []byte
```

GetBytesByTwoOffsetsByPath returns the binary content as []byte from `start` to `end`. Note: Returned value does not contain the character of the last position, which means it returns content range as [start, end). It opens file of `path` for reading with os.O_RDONLY flag and default perm.

​	GetBytesByTwoOffsetsByPath 将二进制内容作为 []byte 从 `start` 到 `end` 返回。注意：返回值不包含最后一个位置的字符，这意味着它返回内容范围为 [start， end]。它打开文件 `path` 以使用操作系统读取。O_RDONLY标志和默认烫发。

##### Example

``` go
```

#### func GetBytesTilChar

```go
func GetBytesTilChar(reader io.ReaderAt, char byte, start int64) ([]byte, int64)
```

GetBytesTilChar returns the contents of the file as []byte until the next specified byte `char` position.

​	GetBytesTilChar 以 []byte 的形式返回文件的内容，直到下一个指定的字节 `char` 位置。

Note: Returned value contains the character of the last position.

​	注意：返回值包含最后一个位置的字符。

#### func GetBytesTilCharByPath

```go
func GetBytesTilCharByPath(path string, char byte, start int64) ([]byte, int64)
```

GetBytesTilCharByPath returns the contents of the file given by `path` as []byte until the next specified byte `char` position. It opens file of `path` for reading with os.O_RDONLY flag and default perm.

​	GetBytesTilCharByPath 返回 `path` 作为 []byte 给出的文件的内容，直到下一个指定的字节 `char` 位置。它打开文件 `path` 以使用操作系统读取。O_RDONLY标志和默认烫发。

Note: Returned value contains the character of the last position.

​	注意：返回值包含最后一个位置的字符。

##### Example

``` go
```

#### func GetBytesWithCache

```go
func GetBytesWithCache(path string, duration ...time.Duration) []byte
```

GetBytesWithCache returns []byte content of given file by `path` from cache. If there’s no content in the cache, it will read it from disk file specified by `path`. The parameter `expire` specifies the caching time for this file content in seconds.

​	GetBytesWithCache 从缓存中返回给定文件的 `path` [] 字节内容。如果缓存中没有内容，它将从 指定的 `path` 磁盘文件中读取内容。该参数 `expire` 指定此文件内容的缓存时间（以秒为单位）。

#### func GetContents

```go
func GetContents(path string) string
```

GetContents returns the file content of `path` as string. It returns en empty string if it fails reading.

​	GetContents 以字符串形式返回文件 `path` 内容。如果读取失败，则返回 en 空字符串。

##### Example

``` go
```

#### func GetContentsWithCache

```go
func GetContentsWithCache(path string, duration ...time.Duration) string
```

GetContentsWithCache returns string content of given file by `path` from cache. If there’s no content in the cache, it will read it from disk file specified by `path`. The parameter `expire` specifies the caching time for this file content in seconds.

​	GetContentsWithCache 从缓存中返回给定文件的 `path` 字符串内容。如果缓存中没有内容，它将从 指定的 `path` 磁盘文件中读取内容。该参数 `expire` 指定此文件内容的缓存时间（以秒为单位）。

##### Example

``` go
```

#### func GetNextCharOffset

```go
func GetNextCharOffset(reader io.ReaderAt, char byte, start int64) int64
```

GetNextCharOffset returns the file offset for given `char` starting from `start`.

​	GetNextCharOffset 返回给定 `char` 的文件偏移量， `start` 从 开始。

#### func GetNextCharOffsetByPath

```go
func GetNextCharOffsetByPath(path string, char byte, start int64) int64
```

GetNextCharOffsetByPath returns the file offset for given `char` starting from `start`. It opens file of `path` for reading with os.O_RDONLY flag and default perm.

​	GetNextCharOffsetByPath 返回给定 `char` 的文件偏移量， `start` 从 开始。它打开文件 `path` 以使用操作系统读取。O_RDONLY标志和默认烫发。

##### Example

``` go
```

#### func Glob

```go
func Glob(pattern string, onlyNames ...bool) ([]string, error)
```

Glob returns the names of all files matching pattern or nil if there is no matching file. The syntax of patterns is the same as in Match. The pattern may describe hierarchical names such as /usr/*/bin/ed (assuming the Separator is ‘/’).

​	Glob 返回所有匹配 pattern 的文件的名称，如果没有匹配的文件，则返回 nil。模式的语法与 Match 中的语法相同。该模式可以描述分层名称，例如 /usr/*/bin/ed（假设分隔符为 '/'）。

Glob ignores file system errors such as I/O errors reading directories. The only possible returned error is ErrBadPattern, when pattern is malformed.

​	Glob 忽略文件系统错误，例如读取目录的 I/O 错误。当模式格式不正确时，唯一可能返回的错误是 ErrBadPattern。

##### Example

``` go
```

#### func Home

```go
func Home(names ...string) (string, error)
```

Home returns absolute path of current user’s home directory. The optional parameter `names` specifies the sub-folders/sub-files, which will be joined with current system separator and returned with the path.

​	Home 返回当前用户主目录的绝对路径。optional 参数 `names` 指定子文件夹/子文件，这些子文件夹/子文件将与当前系统分隔符联接并与路径一起返回。

##### Example

``` go
```

#### func IsDir

```go
func IsDir(path string) bool
```

IsDir checks whether given `path` a directory. Note that it returns false if the `path` does not exist.

​	IsDir 检查是否给定 `path` 了目录。请注意，如果 不存在， `path` 则返回 false。

##### Example

``` go
```

#### func IsEmpty

```go
func IsEmpty(path string) bool
```

IsEmpty checks whether the given `path` is empty. If `path` is a folder, it checks if there’s any file under it. If `path` is a file, it checks if the file size is zero.

​	IsEmpty 检查给定 `path` 的是否为空。如果 `path` 是一个文件夹，它会检查它下面是否有任何文件。如果 `path` 是文件，则检查文件大小是否为零。

Note that it returns true if `path` does not exist.

​	请注意，如果 `path` 不存在，则返回 true。

##### Example

``` go
```

#### func IsFile

```go
func IsFile(path string) bool
```

IsFile checks whether given `path` a file, which means it’s not a directory. Note that it returns false if the `path` does not exist.

​	IsFile 检查是否给定 `path` 文件，这意味着它不是目录。请注意，如果 不存在， `path` 则返回 false。

##### Example

``` go
```

#### func IsReadable

```go
func IsReadable(path string) bool
```

IsReadable checks whether given `path` is readable.

​	IsReadable 检查给定 `path` 的是否可读。

##### Example

``` go
```

#### func IsWritable

```go
func IsWritable(path string) bool
```

IsWritable checks whether given `path` is writable.

​	IsWritable 检查给定 `path` 是否可写。

TODO improve performance; use golang.org/x/sys to cross-plat-form

​	TODO提高性能;使用 golang.org/x/sys 交叉平台形成

##### Example

``` go
```

#### func Join

```go
func Join(paths ...string) string
```

Join joins string array paths with file separator of current system.

​	Join 将字符串数组路径与当前系统的文件分隔符连接起来。

##### Example

``` go
```

#### func MTime

```go
func MTime(path string) time.Time
```

MTime returns the modification time of file given by `path` in second.

​	MTime 返回 `path` 秒 给出的文件的修改时间。

##### Example

``` go
```

#### func MTimestamp

```go
func MTimestamp(path string) int64
```

MTimestamp returns the modification time of file given by `path` in second.

​	MTimestamp 返回以 `path` 秒为单位给出的文件的修改时间。

##### Example

``` go
```

#### func MTimestampMilli

```go
func MTimestampMilli(path string) int64
```

MTimestampMilli returns the modification time of file given by `path` in millisecond.

​	MTimestampMilli 返回以毫秒为 `path` 单位给出的文件的修改时间。

##### Example

``` go
```

#### func MainPkgPath

```go
func MainPkgPath() string
```

MainPkgPath returns absolute file path of package main, which contains the entrance function main.

​	MainPkgPath 返回包 main 的绝对文件路径，其中包含入口函数 main。

It’s only available in develop environment.

​	它仅在开发环境中可用。

Note1: Only valid for source development environments, IE only valid for systems that generate this executable.

​	注1：仅对源代码开发环境有效，IE仅对生成此可执行文件的系统有效。

Note2: When the method is called for the first time, if it is in an asynchronous goroutine, the method may not get the main package path.

​	注2：首次调用该方法时，如果该方法处于异步 goroutine 中，则该方法可能无法获取主包路径。

#### func Mkdir

```go
func Mkdir(path string) (err error)
```

Mkdir creates directories recursively with given `path`. The parameter `path` is suggested to be an absolute path instead of relative one.

​	Mkdir 使用给定 `path` 的 .建议将该参数 `path` 设置为绝对路径，而不是相对路径。

##### Example

``` go
```

#### func Move

```go
func Move(src string, dst string) (err error)
```

Move renames (moves) `src` to `dst` path. If `dst` already exists and is not a directory, it’ll be replaced.

​	将重命名（移动） `src` 移动到 `dst` 路径。如果 `dst` 已存在并且不是目录，则将替换它。

##### Example

``` go
```

#### func Name

```go
func Name(path string) string
```

Name returns the last element of path without file extension.

​	Name 返回不带文件扩展名的 path 的最后一个元素。

Example: Name("/var/www/file.js") -> file Name(“file.js”) -> file

​	示例：Name（“/var/www/file.js”） -> file Name（“file.js”） -> file

##### Example

``` go
```

#### func Open

```go
func Open(path string) (*os.File, error)
```

Open opens file/directory READONLY.

​	打开文件/目录 READONLY。

##### Example

``` go
```

#### func OpenFile

```go
func OpenFile(path string, flag int, perm os.FileMode) (*os.File, error)
```

OpenFile opens file/directory with custom `flag` and `perm`. The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.

​	OpenFile 使用 custom `flag` 和 `perm` 打开文件/目录。参数 `flag` 如下：O_RDONLY、O_RDWR、O_RDWR|O_CREATE|O_TRUNC等

##### Example

``` go
```

#### func OpenWithFlag

```go
func OpenWithFlag(path string, flag int) (*os.File, error)
```

OpenWithFlag opens file/directory with default perm and custom `flag`. The default `perm` is 0666. The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.

​	OpenWithFlag 使用默认 perm 和 custom `flag` 打开文件/目录。默认值 `perm` 为 0666。参数 `flag` 如下：O_RDONLY、O_RDWR、O_RDWR|O_CREATE|O_TRUNC等

##### Example

``` go
```

#### func OpenWithFlagPerm

```go
func OpenWithFlagPerm(path string, flag int, perm os.FileMode) (*os.File, error)
```

OpenWithFlagPerm opens file/directory with custom `flag` and `perm`. The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc. The parameter `perm` is like: 0600, 0666, 0777, etc.

​	OpenWithFlagPerm 使用 custom `flag` 和 `perm` 打开文件/目录。参数 `flag` 如下：O_RDONLY、O_RDWR、O_RDWR|O_CREATE|O_TRUNC等参数 `perm` 如下：0600、0666、0777 等。

#### func PutBytes

```go
func PutBytes(path string, content []byte) error
```

PutBytes puts binary `content` to file of `path`. It creates file of `path` recursively if it does not exist.

​	PutBytes 将二进制文件 `content` 放入 `path` 的文件中。如果它不存在，它会以递归方式创建文件 `path` 。

##### Example

``` go
```

#### func PutBytesAppend

```go
func PutBytesAppend(path string, content []byte) error
```

PutBytesAppend appends binary `content` to file of `path`. It creates file of `path` recursively if it does not exist.

​	PutBytesAppend 将二进制文件 `content` 追加到 的 `path` 文件。如果它不存在，它会以递归方式创建文件 `path` 。

##### Example

``` go
```

#### func PutContents

```go
func PutContents(path string, content string) error
```

PutContents puts string `content` to file of `path`. It creates file of `path` recursively if it does not exist.

​	PutContents 将字符串 `content` 放入 的 `path` 文件中。如果它不存在，它会以递归方式创建文件 `path` 。

##### Example

``` go
```

#### func PutContentsAppend

```go
func PutContentsAppend(path string, content string) error
```

PutContentsAppend appends string `content` to file of `path`. It creates file of `path` recursively if it does not exist.

​	PutContentsAppend 将字符串 `content` 追加到 的 `path` 文件。如果它不存在，它会以递归方式创建文件 `path` 。

##### Example

``` go
```

#### func Pwd

```go
func Pwd() string
```

Pwd returns absolute path of current working directory. Note that it returns an empty string if retrieving current working directory failed.

​	Pwd 返回当前工作目录的绝对路径。请注意，如果检索当前工作目录失败，它将返回一个空字符串。

##### Example

``` go
```

#### func ReadLines

```go
func ReadLines(file string, callback func(line string) error) error
```

ReadLines reads file content line by line, which is passed to the callback function `callback` as string. It matches each line of text, separated by chars ‘\r’ or ‘\n’, stripped any trailing end-of-line marker.

​	ReadLines 逐行读取文件内容，并将其作为字符串传递给回调函数 `callback` 。它匹配每行文本，用字符“\r”或“\n”分隔，去除任何尾随的行尾标记。

Note that the parameter passed to callback function might be an empty value, and the last non-empty line will be passed to callback function `callback` even if it has no newline marker.

​	请注意，传递给回调函数的参数可能是一个空值，最后一个非空行将传递给回调函数 `callback` ，即使它没有换行标记。

##### Example

``` go
```

#### func ReadLinesBytes

```go
func ReadLinesBytes(file string, callback func(bytes []byte) error) error
```

ReadLinesBytes reads file content line by line, which is passed to the callback function `callback` as []byte. It matches each line of text, separated by chars ‘\r’ or ‘\n’, stripped any trailing end-of-line marker.

​	ReadLinesBytes 逐行读取文件内容，并将其作为 []byte 传递给回调函数 `callback` 。它匹配每行文本，用字符“\r”或“\n”分隔，去除任何尾随的行尾标记。

Note that the parameter passed to callback function might be an empty value, and the last non-empty line will be passed to callback function `callback` even if it has no newline marker.

​	请注意，传递给回调函数的参数可能是一个空值，最后一个非空行将被传递给回调函数 `callback` ，即使它没有换行标记。

##### Example

``` go
```

#### func ReadableSize

```go
func ReadableSize(path string) string
```

ReadableSize formats size of file given by `path`, for more human readable.

​	ReadableSize 格式 给出 `path` 的文件大小为 ，以便更易于阅读。

##### Example

``` go
```

#### func RealPath

```go
func RealPath(path string) string
```

RealPath converts the given `path` to its absolute path and checks if the file path exists. If the file does not exist, return an empty string.

​	RealPath 将给定 `path` 路径转换为其绝对路径，并检查文件路径是否存在。如果文件不存在，则返回空字符串。

##### Example

``` go
```

#### func Remove

```go
func Remove(path string) (err error)
```

Remove deletes all file/directory with `path` parameter. If parameter `path` is directory, it deletes it recursively.

​	删除删除所有带有 `path` 参数的文件/目录。如果 parameter `path` 是 directory，它会递归删除它。

It does nothing if given `path` does not exist or is empty.

​	如果给定 `path` 的不存在或为空，则它什么都不做。

##### Example

``` go
```

#### func Rename

```go
func Rename(src string, dst string) error
```

Rename is alias of Move. See Move.

​	Rename 是 Move 的别名。请参阅移动。

##### Example

``` go
```

#### func ReplaceDir

```go
func ReplaceDir(search, replace, path, pattern string, recursive ...bool) error
```

ReplaceDir replaces content for files under `path`. The parameter `pattern` specifies the file pattern which matches to be replaced. It does replacement recursively if given parameter `recursive` is true.

​	ReplaceDir 替换 下 `path` 文件的内容。该参数 `pattern` 指定要替换的匹配文件模式。如果给定的参数 `recursive` 为 true，则以递归方式进行替换。

##### Example

``` go
```

#### func ReplaceDirFunc

```go
func ReplaceDirFunc(f func(path, content string) string, path, pattern string, recursive ...bool) error
```

ReplaceDirFunc replaces content for files under `path` with callback function `f`. The parameter `pattern` specifies the file pattern which matches to be replaced. It does replacement recursively if given parameter `recursive` is true.

​	ReplaceDirFunc 将 下 `path` 的文件的内容替换为 回调函数 `f` 。该参数 `pattern` 指定要替换的匹配文件模式。如果给定的参数 `recursive` 为 true，则以递归方式进行替换。

##### Example

``` go
```

#### func ReplaceFile

```go
func ReplaceFile(search, replace, path string) error
```

ReplaceFile replaces content for file `path`.

​	ReplaceFile 将内容替换为 file `path` 。

##### Example

``` go
```

#### func ReplaceFileFunc

```go
func ReplaceFileFunc(f func(path, content string) string, path string) error
```

ReplaceFileFunc replaces content for file `path` with callback function `f`.

​	ReplaceFileFunc 将文件 `path` 的内容替换为回调函数 `f` 。

##### Example

``` go
```

#### func ScanDir

```go
func ScanDir(path string, pattern string, recursive ...bool) ([]string, error)
```

ScanDir returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

​	ScanDir 返回所有绝对路径为 given `path` 的子文件，如果给定参数 `recursive` 为 true，则递归扫描目录。

The pattern parameter `pattern` supports multiple file name patterns, using the ‘,’ symbol to separate multiple patterns.

​	pattern 参数 `pattern` 支持多个文件名模式，使用“，”符号分隔多个模式。

##### Example

``` go
```

#### func ScanDirFile

```go
func ScanDirFile(path string, pattern string, recursive ...bool) ([]string, error)
```

ScanDirFile returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

​	ScanDirFile 返回所有绝对路径为 given `path` 的子文件，如果给定参数 `recursive` 为 true，则递归扫描目录。

The pattern parameter `pattern` supports multiple file name patterns, using the ‘,’ symbol to separate multiple patterns.

​	pattern 参数 `pattern` 支持多个文件名模式，使用“，”符号分隔多个模式。

Note that it returns only files, exclusive of directories.

​	请注意，它仅返回文件，不包括目录。

##### Example

``` go
```

#### func ScanDirFileFunc

```go
func ScanDirFileFunc(path string, pattern string, recursive bool, handler func(path string) string) ([]string, error)
```

ScanDirFileFunc returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

​	ScanDirFileFunc返回所有绝对路径为given `path` 的子文件，如果given参数 `recursive` 为true，则递归扫描目录。

The pattern parameter `pattern` supports multiple file name patterns, using the ‘,’ symbol to separate multiple patterns.

​	pattern 参数 `pattern` 支持多个文件名模式，使用“，”符号分隔多个模式。

The parameter `recursive` specifies whether scanning the `path` recursively, which means it scans its sub-files and appends the file paths to result array if the sub-file is also a folder. It is false in default.

​	该参数 `recursive` 指定是否以递归方式扫描， `path` 这意味着如果子文件也是一个文件夹，则它会扫描其子文件并将文件路径附加到结果数组。默认情况下为 false。

The parameter `handler` specifies the callback function handling each sub-file path of the `path` and its sub-folders. It ignores the sub-file path if `handler` returns an empty string, or else it appends the sub-file path to result slice.

​	该参数 `handler` 指定处理 及其 `path` 子文件夹的每个子文件路径的回调函数。如果 `handler` 返回空字符串，它将忽略子文件路径，否则会将子文件路径追加到结果切片。

Note that the parameter `path` for `handler` is not a directory but a file. It returns only files, exclusive of directories.

​	请注意， `path` 参数 `handler` 不是目录，而是文件。它只返回文件，不包括目录。

##### Example

``` go
```

#### func ScanDirFunc

```go
func ScanDirFunc(path string, pattern string, recursive bool, handler func(path string) string) ([]string, error)
```

ScanDirFunc returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

​	ScanDirFunc返回所有绝对路径为given `path` 的子文件，如果given参数 `recursive` 为true，则递归扫描目录。

The pattern parameter `pattern` supports multiple file name patterns, using the ‘,’ symbol to separate multiple patterns.

​	pattern 参数 `pattern` 支持多个文件名模式，使用“，”符号分隔多个模式。

The parameter `recursive` specifies whether scanning the `path` recursively, which means it scans its sub-files and appends the files path to result array if the sub-file is also a folder. It is false in default.

​	该参数 `recursive` 指定是否以递归方式扫描， `path` 这意味着如果子文件也是一个文件夹，则它会扫描其子文件并将文件路径附加到结果数组。默认情况下为 false。

The parameter `handler` specifies the callback function handling each sub-file path of the `path` and its sub-folders. It ignores the sub-file path if `handler` returns an empty string, or else it appends the sub-file path to result slice.

​	该参数 `handler` 指定处理 及其 `path` 子文件夹的每个子文件路径的回调函数。如果 `handler` 返回空字符串，它将忽略子文件路径，否则会将子文件路径追加到结果切片。

##### Example

``` go
```

#### func Search

```go
func Search(name string, prioritySearchPaths ...string) (realPath string, err error)
```

Search searches file by name `name` in following paths with priority: prioritySearchPaths, Pwd()、SelfDir()、MainPkgPath(). It returns the absolute file path of `name` if found, or en empty string if not found.

​	搜索在以下路径中按名称 `name` 搜索文件，优先级为：prioritySearchPaths、Pwd（）、SelfDir（）、MainPkgPath（）。它返回 if found 的 `name` 绝对文件路径，如果未找到，则返回 en 空字符串。

##### Example

``` go
```

#### func SelfDir

```go
func SelfDir() string
```

SelfDir returns absolute directory path of current running process(binary).

​	SelfDir 返回当前正在运行的进程（二进制）的绝对目录路径。

##### Example

``` go
```

#### func SelfName

```go
func SelfName() string
```

SelfName returns file name of current running process(binary).

​	SelfName 返回当前正在运行的进程（二进制）的文件名。

##### Example

``` go
```

#### func SelfPath

```go
func SelfPath() string
```

SelfPath returns absolute file path of current running process(binary).

​	SelfPath 返回当前正在运行的进程（二进制）的绝对文件路径。

##### Example

``` go
```

#### func Size

```go
func Size(path string) int64
```

Size returns the size of file specified by `path` in byte.

​	Size 返回以 `path` 字节为单位指定的文件大小。

##### Example

``` go
```

#### func SizeFormat

```go
func SizeFormat(path string) string
```

SizeFormat returns the size of file specified by `path` in format string.

​	SizeFormat 返回 `path` in format string 指定的文件大小。

##### Example

``` go
```

#### func SortFiles

```go
func SortFiles(files []string) []string
```

SortFiles sorts the `files` in order of: directory -> file. Note that the item of `files` should be absolute path.

​	SortFiles `files` 按以下顺序对文件进行排序：directory -> 文件。请注意，的 `files` 项应该是绝对路径。

##### Example

``` go
```

#### func Stat

```go
func Stat(path string) (os.FileInfo, error)
```

Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError.

​	Stat 返回描述命名文件的 FileInfo。如果存在错误，则其类型为 *PathError。

##### Example

``` go
```

#### func StrToSize

```go
func StrToSize(sizeStr string) int64
```

StrToSize converts formatted size string to its size in bytes.

​	StrToSize 将格式化的大小字符串转换为其大小（以字节为单位）。

##### Example

``` go
```

#### func Temp

```go
func Temp(names ...string) string
```

Temp retrieves and returns the temporary directory of current system.

​	Temp 检索并返回当前系统的临时目录。

The optional parameter `names` specifies the sub-folders/sub-files, which will be joined with current system separator and returned with the path.

​	optional 参数 `names` 指定子文件夹/子文件，这些子文件夹/子文件将与当前系统分隔符联接并与路径一起返回。

#### func Truncate

```go
func Truncate(path string, size int) (err error)
```

Truncate truncates file of `path` to given size by `size`.

​	截断 截断给定大小的文件 `path` `size` 。

## 类型

### type CopyOption <-2.5.4

```go
type CopyOption struct {
	// Auto call file sync after source file content copied to target file.
	Sync bool

	// Preserve the mode of the original file to the target file.
	// If true, the Mode attribute will make no sense.
	PreserveMode bool

	// Destination created file mode.
	// The default file mode is DefaultPermCopy if PreserveMode is false.
	Mode os.FileMode
}
```

CopyOption is the option for Copy* functions.

​	CopyOption 是 Copy* 函数的选项。