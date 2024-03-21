+++
title = "gfile"
date = 2024-03-21T17:55:29+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gfile

Package gfile provides easy-to-use operations for file system.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gfile/gfile.go#L23)

``` go
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

### Variables 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gfile/gfile_contents.go#L17)

``` go
var (
	// DefaultReadBuffer is the buffer size for reading file content.
	DefaultReadBuffer = 1024
)
```

### Functions 

##### func Abs 

``` go
func Abs(path string) string
```

Abs returns an absolute representation of path. If the path is not absolute it will be joined with the current working directory to turn it into an absolute path. The absolute path name for a given file is not guaranteed to be unique. Abs calls Clean on the result.

##### Example

``` go
```
##### func Basename 

``` go
func Basename(path string) string
```

Basename returns the last element of path, which contains file extension. Trailing path separators are removed before extracting the last element. If the path is empty, Base returns ".". If the path consists entirely of separators, Basename returns a single separator.

Example: Basename("/var/www/file.js") -> file.js Basename("file.js") -> file.js

##### Example

``` go
```
##### func Chdir 

``` go
func Chdir(dir string) (err error)
```

Chdir changes the current working directory to the named directory. If there is an error, it will be of type *PathError.

##### Example

``` go
```
##### func Chmod 

``` go
func Chmod(path string, mode os.FileMode) (err error)
```

Chmod is alias of os.Chmod. See os.Chmod.

##### Example

``` go
```
##### func Copy 

``` go
func Copy(src string, dst string, option ...CopyOption) error
```

Copy file/directory from `src` to `dst`.

If `src` is file, it calls CopyFile to implements copy feature, or else it calls CopyDir.

If `src` is file, but `dst` already exists and is a folder, it then creates a same name file of `src` in folder `dst`.

Eg: Copy("/tmp/file1", "/tmp/file2") => /tmp/file1 copied to /tmp/file2 Copy("/tmp/dir1", "/tmp/dir2") => /tmp/dir1 copied to /tmp/dir2 Copy("/tmp/file1", "/tmp/dir2") => /tmp/file1 copied to /tmp/dir2/file1 Copy("/tmp/dir1", "/tmp/file2") => error

##### Example

``` go
```
##### func CopyDir 

``` go
func CopyDir(src string, dst string, option ...CopyOption) (err error)
```

CopyDir recursively copies a directory tree, attempting to preserve permissions.

Note that, the Source directory must exist and symlinks are ignored and skipped.

##### func CopyFile 

``` go
func CopyFile(src, dst string, option ...CopyOption) (err error)
```

CopyFile copies the contents of the file named `src` to the file named by `dst`. The file will be created if it does not exist. If the destination file exists, all it's contents will be replaced by the contents of the source file. The file mode will be copied from the source and the copied data is synced/flushed to stable storage. Thanks: https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04

##### func Create 

``` go
func Create(path string) (*os.File, error)
```

Create creates a file with given `path` recursively. The parameter `path` is suggested to be absolute path.

##### Example

``` go
```
##### func Dir 

``` go
func Dir(path string) string
```

Dir returns all but the last element of path, typically the path's directory. After dropping the final element, Dir calls Clean on the path and trailing slashes are removed. If the `path` is empty, Dir returns ".". If the `path` is ".", Dir treats the path as current working directory. If the `path` consists entirely of separators, Dir returns a single separator. The returned path does not end in a separator unless it is the root directory.

Example: Dir("/var/www/file.js") -> "/var/www" Dir("file.js") -> "."

##### Example

``` go
```
##### func DirNames 

``` go
func DirNames(path string) ([]string, error)
```

DirNames returns sub-file names of given directory `path`. Note that the returned names are NOT absolute paths.

##### Example

``` go
```
##### func Exists 

``` go
func Exists(path string) bool
```

Exists checks whether given `path` exist.

##### Example

``` go
```
##### func Ext 

``` go
func Ext(path string) string
```

Ext returns the file name extension used by path. The extension is the suffix beginning at the final dot in the final element of path; it is empty if there is no dot. Note: the result contains symbol '.'.

Example: Ext("main.go") => .go Ext("api.json") => .json

##### Example

``` go
```
##### func ExtName 

``` go
func ExtName(path string) string
```

ExtName is like function Ext, which returns the file name extension used by path, but the result does not contain symbol '.'.

Example: ExtName("main.go") => go ExtName("api.json") => json

##### Example

``` go
```
##### func FormatSize 

``` go
func FormatSize(raw int64) string
```

FormatSize formats size `raw` for more manually readable.

##### Example

``` go
```
##### func GetBytes 

``` go
func GetBytes(path string) []byte
```

GetBytes returns the file content of `path` as []byte. It returns nil if it fails reading.

##### Example

``` go
```
##### func GetBytesByTwoOffsets 

``` go
func GetBytesByTwoOffsets(reader io.ReaderAt, start int64, end int64) []byte
```

GetBytesByTwoOffsets returns the binary content as []byte from `start` to `end`. Note: Returned value does not contain the character of the last position, which means it returns content range as [start, end).

##### func GetBytesByTwoOffsetsByPath 

``` go
func GetBytesByTwoOffsetsByPath(path string, start int64, end int64) []byte
```

GetBytesByTwoOffsetsByPath returns the binary content as []byte from `start` to `end`. Note: Returned value does not contain the character of the last position, which means it returns content range as [start, end). It opens file of `path` for reading with os.O_RDONLY flag and default perm.

##### Example

``` go
```
##### func GetBytesTilChar 

``` go
func GetBytesTilChar(reader io.ReaderAt, char byte, start int64) ([]byte, int64)
```

GetBytesTilChar returns the contents of the file as []byte until the next specified byte `char` position.

Note: Returned value contains the character of the last position.

##### func GetBytesTilCharByPath 

``` go
func GetBytesTilCharByPath(path string, char byte, start int64) ([]byte, int64)
```

GetBytesTilCharByPath returns the contents of the file given by `path` as []byte until the next specified byte `char` position. It opens file of `path` for reading with os.O_RDONLY flag and default perm.

Note: Returned value contains the character of the last position.

##### Example

``` go
```
##### func GetBytesWithCache 

``` go
func GetBytesWithCache(path string, duration ...time.Duration) []byte
```

GetBytesWithCache returns []byte content of given file by `path` from cache. If there's no content in the cache, it will read it from disk file specified by `path`. The parameter `expire` specifies the caching time for this file content in seconds.

##### func GetContents 

``` go
func GetContents(path string) string
```

GetContents returns the file content of `path` as string. It returns en empty string if it fails reading.

##### Example

``` go
```
##### func GetContentsWithCache 

``` go
func GetContentsWithCache(path string, duration ...time.Duration) string
```

GetContentsWithCache returns string content of given file by `path` from cache. If there's no content in the cache, it will read it from disk file specified by `path`. The parameter `expire` specifies the caching time for this file content in seconds.

##### Example

``` go
```
##### func GetNextCharOffset 

``` go
func GetNextCharOffset(reader io.ReaderAt, char byte, start int64) int64
```

GetNextCharOffset returns the file offset for given `char` starting from `start`.

##### func GetNextCharOffsetByPath 

``` go
func GetNextCharOffsetByPath(path string, char byte, start int64) int64
```

GetNextCharOffsetByPath returns the file offset for given `char` starting from `start`. It opens file of `path` for reading with os.O_RDONLY flag and default perm.

##### Example

``` go
```
##### func Glob 

``` go
func Glob(pattern string, onlyNames ...bool) ([]string, error)
```

Glob returns the names of all files matching pattern or nil if there is no matching file. The syntax of patterns is the same as in Match. The pattern may describe hierarchical names such as /usr/*/bin/ed (assuming the Separator is '/').

Glob ignores file system errors such as I/O errors reading directories. The only possible returned error is ErrBadPattern, when pattern is malformed.

##### Example

``` go
```
##### func Home 

``` go
func Home(names ...string) (string, error)
```

Home returns absolute path of current user's home directory. The optional parameter `names` specifies the sub-folders/sub-files, which will be joined with current system separator and returned with the path.

##### Example

``` go
```
##### func IsDir 

``` go
func IsDir(path string) bool
```

IsDir checks whether given `path` a directory. Note that it returns false if the `path` does not exist.

##### Example

``` go
```
##### func IsEmpty 

``` go
func IsEmpty(path string) bool
```

IsEmpty checks whether the given `path` is empty. If `path` is a folder, it checks if there's any file under it. If `path` is a file, it checks if the file size is zero.

Note that it returns true if `path` does not exist.

##### Example

``` go
```
##### func IsFile 

``` go
func IsFile(path string) bool
```

IsFile checks whether given `path` a file, which means it's not a directory. Note that it returns false if the `path` does not exist.

##### Example

``` go
```
##### func IsReadable 

``` go
func IsReadable(path string) bool
```

IsReadable checks whether given `path` is readable.

##### Example

``` go
```
##### func IsWritable 

``` go
func IsWritable(path string) bool
```

IsWritable checks whether given `path` is writable.

TODO improve performance; use golang.org/x/sys to cross-plat-form

##### Example

``` go
```
##### func Join 

``` go
func Join(paths ...string) string
```

Join joins string array paths with file separator of current system.

##### Example

``` go
```
##### func MTime 

``` go
func MTime(path string) time.Time
```

MTime returns the modification time of file given by `path` in second.

##### Example

``` go
```
##### func MTimestamp 

``` go
func MTimestamp(path string) int64
```

MTimestamp returns the modification time of file given by `path` in second.

##### Example

``` go
```
##### func MTimestampMilli 

``` go
func MTimestampMilli(path string) int64
```

MTimestampMilli returns the modification time of file given by `path` in millisecond.

##### Example

``` go
```
##### func MainPkgPath 

``` go
func MainPkgPath() string
```

MainPkgPath returns absolute file path of package main, which contains the entrance function main.

It's only available in develop environment.

Note1: Only valid for source development environments, IE only valid for systems that generate this executable.

Note2: When the method is called for the first time, if it is in an asynchronous goroutine, the method may not get the main package path.

##### func Mkdir 

``` go
func Mkdir(path string) (err error)
```

Mkdir creates directories recursively with given `path`. The parameter `path` is suggested to be an absolute path instead of relative one.

##### Example

``` go
```
##### func Move 

``` go
func Move(src string, dst string) (err error)
```

Move renames (moves) `src` to `dst` path. If `dst` already exists and is not a directory, it'll be replaced.

##### Example

``` go
```
##### func Name 

``` go
func Name(path string) string
```

Name returns the last element of path without file extension.

Example: Name("/var/www/file.js") -> file Name("file.js") -> file

##### Example

``` go
```
##### func Open 

``` go
func Open(path string) (*os.File, error)
```

Open opens file/directory READONLY.

##### Example

``` go
```
##### func OpenFile 

``` go
func OpenFile(path string, flag int, perm os.FileMode) (*os.File, error)
```

OpenFile opens file/directory with custom `flag` and `perm`. The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.

##### Example

``` go
```
##### func OpenWithFlag 

``` go
func OpenWithFlag(path string, flag int) (*os.File, error)
```

OpenWithFlag opens file/directory with default perm and custom `flag`. The default `perm` is 0666. The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc.

##### Example

``` go
```
##### func OpenWithFlagPerm 

``` go
func OpenWithFlagPerm(path string, flag int, perm os.FileMode) (*os.File, error)
```

OpenWithFlagPerm opens file/directory with custom `flag` and `perm`. The parameter `flag` is like: O_RDONLY, O_RDWR, O_RDWR|O_CREATE|O_TRUNC, etc. The parameter `perm` is like: 0600, 0666, 0777, etc.

##### func PutBytes 

``` go
func PutBytes(path string, content []byte) error
```

PutBytes puts binary `content` to file of `path`. It creates file of `path` recursively if it does not exist.

##### Example

``` go
```
##### func PutBytesAppend 

``` go
func PutBytesAppend(path string, content []byte) error
```

PutBytesAppend appends binary `content` to file of `path`. It creates file of `path` recursively if it does not exist.

##### Example

``` go
```
##### func PutContents 

``` go
func PutContents(path string, content string) error
```

PutContents puts string `content` to file of `path`. It creates file of `path` recursively if it does not exist.

##### Example

``` go
```
##### func PutContentsAppend 

``` go
func PutContentsAppend(path string, content string) error
```

PutContentsAppend appends string `content` to file of `path`. It creates file of `path` recursively if it does not exist.

##### Example

``` go
```
##### func Pwd 

``` go
func Pwd() string
```

Pwd returns absolute path of current working directory. Note that it returns an empty string if retrieving current working directory failed.

##### Example

``` go
```
##### func ReadLines 

``` go
func ReadLines(file string, callback func(line string) error) error
```

ReadLines reads file content line by line, which is passed to the callback function `callback` as string. It matches each line of text, separated by chars '\r' or '\n', stripped any trailing end-of-line marker.

Note that the parameter passed to callback function might be an empty value, and the last non-empty line will be passed to callback function `callback` even if it has no newline marker.

##### Example

``` go
```
##### func ReadLinesBytes 

``` go
func ReadLinesBytes(file string, callback func(bytes []byte) error) error
```

ReadLinesBytes reads file content line by line, which is passed to the callback function `callback` as []byte. It matches each line of text, separated by chars '\r' or '\n', stripped any trailing end-of-line marker.

Note that the parameter passed to callback function might be an empty value, and the last non-empty line will be passed to callback function `callback` even if it has no newline marker.

##### Example

``` go
```
##### func ReadableSize 

``` go
func ReadableSize(path string) string
```

ReadableSize formats size of file given by `path`, for more human readable.

##### Example

``` go
```
##### func RealPath 

``` go
func RealPath(path string) string
```

RealPath converts the given `path` to its absolute path and checks if the file path exists. If the file does not exist, return an empty string.

##### Example

``` go
```
##### func Remove 

``` go
func Remove(path string) (err error)
```

Remove deletes all file/directory with `path` parameter. If parameter `path` is directory, it deletes it recursively.

It does nothing if given `path` does not exist or is empty.

##### Example

``` go
```
##### func Rename 

``` go
func Rename(src string, dst string) error
```

Rename is alias of Move. See Move.

##### Example

``` go
```
##### func ReplaceDir 

``` go
func ReplaceDir(search, replace, path, pattern string, recursive ...bool) error
```

ReplaceDir replaces content for files under `path`. The parameter `pattern` specifies the file pattern which matches to be replaced. It does replacement recursively if given parameter `recursive` is true.

##### Example

``` go
```
##### func ReplaceDirFunc 

``` go
func ReplaceDirFunc(f func(path, content string) string, path, pattern string, recursive ...bool) error
```

ReplaceDirFunc replaces content for files under `path` with callback function `f`. The parameter `pattern` specifies the file pattern which matches to be replaced. It does replacement recursively if given parameter `recursive` is true.

##### Example

``` go
```
##### func ReplaceFile 

``` go
func ReplaceFile(search, replace, path string) error
```

ReplaceFile replaces content for file `path`.

##### Example

``` go
```
##### func ReplaceFileFunc 

``` go
func ReplaceFileFunc(f func(path, content string) string, path string) error
```

ReplaceFileFunc replaces content for file `path` with callback function `f`.

##### Example

``` go
```
##### func ScanDir 

``` go
func ScanDir(path string, pattern string, recursive ...bool) ([]string, error)
```

ScanDir returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

The pattern parameter `pattern` supports multiple file name patterns, using the ',' symbol to separate multiple patterns.

##### Example

``` go
```
##### func ScanDirFile 

``` go
func ScanDirFile(path string, pattern string, recursive ...bool) ([]string, error)
```

ScanDirFile returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

The pattern parameter `pattern` supports multiple file name patterns, using the ',' symbol to separate multiple patterns.

Note that it returns only files, exclusive of directories.

##### Example

``` go
```
##### func ScanDirFileFunc 

``` go
func ScanDirFileFunc(path string, pattern string, recursive bool, handler func(path string) string) ([]string, error)
```

ScanDirFileFunc returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

The pattern parameter `pattern` supports multiple file name patterns, using the ',' symbol to separate multiple patterns.

The parameter `recursive` specifies whether scanning the `path` recursively, which means it scans its sub-files and appends the file paths to result array if the sub-file is also a folder. It is false in default.

The parameter `handler` specifies the callback function handling each sub-file path of the `path` and its sub-folders. It ignores the sub-file path if `handler` returns an empty string, or else it appends the sub-file path to result slice.

Note that the parameter `path` for `handler` is not a directory but a file. It returns only files, exclusive of directories.

##### Example

``` go
```
##### func ScanDirFunc 

``` go
func ScanDirFunc(path string, pattern string, recursive bool, handler func(path string) string) ([]string, error)
```

ScanDirFunc returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

The pattern parameter `pattern` supports multiple file name patterns, using the ',' symbol to separate multiple patterns.

The parameter `recursive` specifies whether scanning the `path` recursively, which means it scans its sub-files and appends the files path to result array if the sub-file is also a folder. It is false in default.

The parameter `handler` specifies the callback function handling each sub-file path of the `path` and its sub-folders. It ignores the sub-file path if `handler` returns an empty string, or else it appends the sub-file path to result slice.

##### Example

``` go
```
##### func Search 

``` go
func Search(name string, prioritySearchPaths ...string) (realPath string, err error)
```

Search searches file by name `name` in following paths with priority: prioritySearchPaths, Pwd()、SelfDir()、MainPkgPath(). It returns the absolute file path of `name` if found, or en empty string if not found.

##### Example

``` go
```
##### func SelfDir 

``` go
func SelfDir() string
```

SelfDir returns absolute directory path of current running process(binary).

##### Example

``` go
```
##### func SelfName 

``` go
func SelfName() string
```

SelfName returns file name of current running process(binary).

##### Example

``` go
```
##### func SelfPath 

``` go
func SelfPath() string
```

SelfPath returns absolute file path of current running process(binary).

##### Example

``` go
```
##### func Size 

``` go
func Size(path string) int64
```

Size returns the size of file specified by `path` in byte.

##### Example

``` go
```
##### func SizeFormat 

``` go
func SizeFormat(path string) string
```

SizeFormat returns the size of file specified by `path` in format string.

##### Example

``` go
```
##### func SortFiles 

``` go
func SortFiles(files []string) []string
```

SortFiles sorts the `files` in order of: directory -> file. Note that the item of `files` should be absolute path.

##### Example

``` go
```
##### func Stat 

``` go
func Stat(path string) (os.FileInfo, error)
```

Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError.

##### Example

``` go
```
##### func StrToSize 

``` go
func StrToSize(sizeStr string) int64
```

StrToSize converts formatted size string to its size in bytes.

##### Example

``` go
```
##### func Temp 

``` go
func Temp(names ...string) string
```

Temp retrieves and returns the temporary directory of current system.

The optional parameter `names` specifies the sub-folders/sub-files, which will be joined with current system separator and returned with the path.

##### func Truncate 

``` go
func Truncate(path string, size int) (err error)
```

Truncate truncates file of `path` to given size by `size`.

### Types 

#### type CopyOption <-2.5.4

``` go
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