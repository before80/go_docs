+++
title = "gres"
date = 2024-03-21T17:56:51+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gres](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gres)

Package gres provides resource management and packing/unpacking feature between files and bytes.

​	Package gres 提供资源管理和文件和字节之间的打包/解包功能。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gres/gres_instance.go#L11)

```go
const (
	// DefaultName default group name for instance usage.
	DefaultName = "default"
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gres/gres.go#L10)

```go
const (
	// Separator for directories.
	Separator = "/"
)
```

## 变量

This section is empty.

## 函数

#### func Add

```go
func Add(content string, prefix ...string) error
```

Add unpacks and adds the `content` into the default resource object. The unnecessary parameter `prefix` indicates the prefix for each file storing into current resource object.

​	添加解压缩并将 添加到默认资源对象中 `content` 。不必要的参数 `prefix` 指示存储到当前资源对象中的每个文件的前缀。

#### func Contains

```go
func Contains(path string) bool
```

Contains checks whether the `path` exists in the default resource object.

​	包含检查默认资源对象中是否存在。 `path`

#### func Dump

```go
func Dump()
```

Dump prints the files of the default resource object.

​	转储打印默认资源对象的文件。

#### func Export

```go
func Export(src, dst string, option ...ExportOption) error
```

Export exports and saves specified path `src` and all its sub files to specified system path `dst` recursively.

​	导出 以递归方式导出并保存指定路径 `src` 及其所有子文件到指定的系统路径 `dst` 。

#### func GetContent

```go
func GetContent(path string) []byte
```

GetContent directly returns the content of `path` in default resource object.

​	GetContent 直接返回默认资源对象的内容 `path` 。

#### func IsEmpty

```go
func IsEmpty() bool
```

IsEmpty checks and returns whether the resource manager is empty.

​	IsEmpty 检查并返回资源管理器是否为空。

#### func Load

```go
func Load(path string, prefix ...string) error
```

Load loads, unpacks and adds the data from `path` into the default resource object. The unnecessary parameter `prefix` indicates the prefix for each file storing into current resource object.

​	加载、加载、解压缩数据并将其添加到默认资源对象 `path` 中。不必要的参数 `prefix` 指示存储到当前资源对象中的每个文件的前缀

func
功能[Pack
包](https://github.com/gogf/gf/blob/v2.6.4/os/gres/gres_func.go#L49)DEPRECATED

func
功能[PackToFile](https://github.com/gogf/gf/blob/v2.6.4/os/gres/gres_func.go#L77)DEPRECATE



#### func PackToFileWithOption <-2.2.1

```go
func PackToFileWithOption(srcPaths, dstPath string, option Option) error
```

PackToFileWithOption packs the path specified by `srcPaths` to target file `dstPath`.

​	PackToFileWithOption 将 `srcPaths` 指定的路径打包到 目标文件 `dstPath` 。

Note that parameter `srcPaths` supports multiple paths join with ‘,’.

​	请注意，参数 `srcPaths` 支持多个路径与“，”连接。

unc
功能[PackToGoFile
PackToGo文件](https://github.com/gogf/gf/blob/v2.6.4/os/gres/gres_func.go#L105)DEPRECATED

#### func PackToGoFileWithOption <-2.2.1

```go
func PackToGoFileWithOption(srcPath, goFilePath, pkgName string, option Option) error
```

PackToGoFileWithOption packs the path specified by `srcPaths` to target go file `goFilePath` with given package name `pkgName`.

​	PackToGoFileWithOption 将 `srcPaths` 指定的路径打包到目标 go 文件 `goFilePath` ，并具有给定的包名称 `pkgName` 。

Note that parameter `srcPaths` supports multiple paths join with ‘,’.

​	请注意，参数 `srcPaths` 支持多个路径与“，”连接。

#### func PackWithOption <-2.2.1

```go
func PackWithOption(srcPaths string, option Option) ([]byte, error)
```

PackWithOption packs the path specified by `srcPaths` into bytes.

​	PackWithOption 将指定的 `srcPaths` 路径打包到字节中。

Note that parameter `srcPaths` supports multiple paths join with ‘,’.

​	请注意，参数 `srcPaths` 支持多个路径与“，”连接。

## 类型

### type ExportOption

```go
type ExportOption struct {
	RemovePrefix string // Remove the prefix of file name from resource.
}
```

ExportOption is the option for function Export.

​	ExportOption 是函数 Export 的选项。

### type File

```go
type File struct {
	// contains filtered or unexported fields
}
```

#### func Get

```go
func Get(path string) *File
```

Get returns the file with given path.

​	get 返回具有给定路径的文件。

#### func GetWithIndex

```go
func GetWithIndex(path string, indexFiles []string) *File
```

GetWithIndex searches file with `path`, if the file is directory it then does index files searching under this directory.

​	GetWithIndex 搜索 的文件 `path` ，如果文件是目录，则在此目录下搜索索引文件。

GetWithIndex is usually used for http static file service.

​	GetWithIndex 通常用于 http 静态文件服务。

#### func ScanDir

```go
func ScanDir(path string, pattern string, recursive ...bool) []*File
```

ScanDir returns the files under the given path, the parameter `path` should be a folder type.

​	ScanDir 返回给定路径下的文件，参数 `path` 应为文件夹类型。

The pattern parameter `pattern` supports multiple file name patterns, using the ‘,’ symbol to separate multiple patterns.

​	pattern 参数 `pattern` 支持多个文件名模式，使用“，”符号分隔多个模式。

It scans directory recursively if given parameter `recursive` is true.

​	如果给定的参数 `recursive` 为 true，它会递归扫描目录。

#### func ScanDirFile

```go
func ScanDirFile(path string, pattern string, recursive ...bool) []*File
```

ScanDirFile returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

​	ScanDirFile 返回所有绝对路径为 given `path` 的子文件，如果给定参数 `recursive` 为 true，则递归扫描目录。

Note that it returns only files, exclusive of directories.

​	请注意，它仅返回文件，不包括目录。

#### func Unpack

```go
func Unpack(path string) ([]*File, error)
```

Unpack unpacks the content specified by `path` to []*File.

​	Unpack 解压缩 `path` to []*File 指定的内容。

#### func UnpackContent

```go
func UnpackContent(content string) ([]*File, error)
```

UnpackContent unpacks the content to []*File.

​	UnpackContent 将内容解压缩到 []*File。

#### (*File) Close

```go
func (f *File) Close() error
```

Close implements interface of http.File.

​	关闭实现http的接口。文件。

#### (*File) Content

```go
func (f *File) Content() []byte
```

Content returns the content of the file.

​	Content 返回文件的内容。

#### (*File) Export

```go
func (f *File) Export(dst string, option ...ExportOption) error
```

Export exports and saves all its sub files to specified system path `dst` recursively.

​	导出 以递归方式将其所有子文件导出并保存到指定的系统路径 `dst` 。

#### (*File) FileInfo

```go
func (f *File) FileInfo() os.FileInfo
```

FileInfo returns an os.FileInfo for the FileHeader.

​	FileInfo 返回一个 os。FileHeader 的 FileInfo。

#### (File) MarshalJSON

```go
func (f File) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。

#### (*File) Name

```go
func (f *File) Name() string
```

Name returns the name of the file.

​	Name 返回文件的名称。

#### (*File) Open

```go
func (f *File) Open() (io.ReadCloser, error)
```

Open returns a ReadCloser that provides access to the File’s contents. Multiple files may be read concurrently.

​	Open 返回一个 ReadCloser，用于提供对文件内容的访问。可以同时读取多个文件。

#### (*File) Read

```go
func (f *File) Read(b []byte) (n int, err error)
```

Read implements the io.Reader interface.

​	Read 实现 io。阅读器界面。

#### (*File) Readdir

```go
func (f *File) Readdir(count int) ([]os.FileInfo, error)
```

Readdir implements Readdir interface of http.File.

​	Readdir 实现了 http 的 Readdir 接口。文件。

#### (*File) Seek

```go
func (f *File) Seek(offset int64, whence int) (n int64, err error)
```

Seek implements the io.Seeker interface.

​	Seek 实现 io。导引头接口。

#### (*File) Stat

```go
func (f *File) Stat() (os.FileInfo, error)
```

Stat implements Stat interface of http.File.

​	Stat 实现了 http 的 Stat 接口。文件。

### type Option <-2.2.1

```go
type Option struct {
	Prefix   string // The file path prefix for each file item in resource manager.
	KeepPath bool   // Keep the passed path when packing, usually for relative path.
}
```

Option contains the extra options for Pack functions.

​	选项包含 Pack 函数的额外选项。

### type Resource

```go
type Resource struct {
	// contains filtered or unexported fields
}
```

#### func Instance

```go
func Instance(name ...string) *Resource
```

Instance returns an instance of Resource. The parameter `name` is the name for the instance.

​	Instance 返回 Resource 的实例。该参数 `name` 是实例的名称。

#### func New

```go
func New() *Resource
```

New creates and returns a new resource object.

​	New 创建并返回一个新的资源对象。

#### (*Resource) Add

```go
func (r *Resource) Add(content string, prefix ...string) error
```

Add unpacks and adds the `content` into current resource object. The unnecessary parameter `prefix` indicates the prefix for each file storing into current resource object.

​	添加解压缩并添加到当前资源对象中 `content` 。不必要的参数 `prefix` 指示存储到当前资源对象中的每个文件的前缀。

#### (*Resource) Contains

```go
func (r *Resource) Contains(path string) bool
```

Contains checks whether the `path` exists in current resource object.

​	包含检查当前资源对象中是否存在。 `path`

#### (*Resource) Dump

```go
func (r *Resource) Dump()
```

Dump prints the files of current resource object.

​	转储打印当前资源对象的文件。

#### (*Resource) Export

```go
func (r *Resource) Export(src, dst string, option ...ExportOption) error
```

Export exports and saves specified path `srcPath` and all its sub files to specified system path `dstPath` recursively.

​	导出 以递归方式导出并保存指定路径 `srcPath` 及其所有子文件到指定的系统路径 `dstPath` 。

#### (*Resource) Get

```go
func (r *Resource) Get(path string) *File
```

Get returns the file with given path.

​	get 返回具有给定路径的文件。

#### (*Resource) GetContent

```go
func (r *Resource) GetContent(path string) []byte
```

GetContent directly returns the content of `path`.

​	GetContent 直接返回 的内容 `path` 。

#### (*Resource) GetWithIndex

```go
func (r *Resource) GetWithIndex(path string, indexFiles []string) *File
```

GetWithIndex searches file with `path`, if the file is directory it then does index files searching under this directory.

​	GetWithIndex 搜索 的文件 `path` ，如果文件是目录，则在此目录下搜索索引文件。

GetWithIndex is usually used for http static file service.

​	GetWithIndex 通常用于 http 静态文件服务。

#### (*Resource) IsEmpty

```go
func (r *Resource) IsEmpty() bool
```

IsEmpty checks and returns whether the resource manager is empty.

​	IsEmpty 检查并返回资源管理器是否为空。

#### (*Resource) Load

```go
func (r *Resource) Load(path string, prefix ...string) error
```

Load loads, unpacks and adds the data from `path` into current resource object. The unnecessary parameter `prefix` indicates the prefix for each file storing into current resource object.

​	加载、加载、解压缩数据并将其添加到当前资源对象 `path` 中。不必要的参数 `prefix` 指示存储到当前资源对象中的每个文件的前缀。

#### (*Resource) ScanDir

```go
func (r *Resource) ScanDir(path string, pattern string, recursive ...bool) []*File
```

ScanDir returns the files under the given path, the parameter `path` should be a folder type.

​	ScanDir 返回给定路径下的文件，参数 `path` 应为文件夹类型。

The pattern parameter `pattern` supports multiple file name patterns, using the ‘,’ symbol to separate multiple patterns.

​	pattern 参数 `pattern` 支持多个文件名模式，使用“，”符号分隔多个模式。

It scans directory recursively if given parameter `recursive` is true.

​	如果给定的参数 `recursive` 为 true，它会递归扫描目录。

Note that the returned files does not contain given parameter `path`.

​	请注意，返回的文件不包含给定的参数 `path` 。

#### (*Resource) ScanDirFile

```go
func (r *Resource) ScanDirFile(path string, pattern string, recursive ...bool) []*File
```

ScanDirFile returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

​	ScanDirFile 返回所有绝对路径为 given `path` 的子文件，如果给定参数 `recursive` 为 true，则递归扫描目录。

Note that it returns only files, exclusive of directories.

​	请注意，它仅返回文件，不包括目录。