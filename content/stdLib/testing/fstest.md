+++
title = "fstest"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/testing/fstest@go1.20.1

fstest包实现了用于测试文件系统实现和用户的支持。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func TestFS 

``` go 
func TestFS(fsys fs.FS, expected ...string) error
```

​	TestFS函数测试文件系统实现。它遍历fsys中的所有文件树，打开并检查每个文件是否正确。它还检查文件系统至少包含所需的文件。特殊情况是，如果没有列出所需的文件，则fsys必须为空。否则，fsys必须至少包含所列出的文件；它也可以包含其他文件。fsys的内容不能与TestFS函数并发更改。

​	如果TestFS函数发现任何不良行为，它将返回一个报告所有不良行为的错误。错误文本跨多行，每行报告一个检测到的不良行为。

​	在测试内部的典型用法是：

```
if err := fstest.TestFS(myFS, "file/that/should/be/present"); err != nil {
	t.Fatal(err)
}
```

## 类型

### type MapFS 

``` go 
type MapFS map[string]*MapFile
```

​	MapFS是一个简单的内存文件系统，用于测试，表示为从路径名(Open的参数)到有关它们所表示的文件或目录的信息的映射。

​	该映射不必包括包含在映射中的文件的父目录；如果需要，这些父目录将被合成。但是，可以通过设置MapFile.Mode的ModeDir位来包括目录；这可能是对目录的FileInfo进行详细控制或创建空目录所必需的。

​	文件系统操作直接从映射中读取，因此可以通过编辑映射来更改文件系统。一个含义是文件系统操作不能与对映射的更改同时运行，这将导致竞争。另一个含义是打开或读取目录需要迭代整个映射，因此MapFS通常应仅用于少于几百个条目或目录读取的情况。

#### (MapFS) Glob 

``` go 
func (fsys MapFS) Glob(pattern string) ([]string, error)
```

#### (MapFS) Open 

``` go 
func (fsys MapFS) Open(name string) (fs.File, error)
```

Open opens the named file.

#### (MapFS) ReadDir 

``` go 
func (fsys MapFS) ReadDir(name string) ([]fs.DirEntry, error)
```

#### (MapFS) ReadFile 

``` go 
func (fsys MapFS) ReadFile(name string) ([]byte, error)
```

#### (MapFS) Stat 

``` go 
func (fsys MapFS) Stat(name string) (fs.FileInfo, error)
```

#### (MapFS) Sub 

``` go 
func (fsys MapFS) Sub(dir string) (fs.FS, error)
```

### type MapFile 

``` go 
type MapFile struct {
	Data    []byte      // file content
	Mode    fs.FileMode // FileInfo.Mode
	ModTime time.Time   // FileInfo.ModTime
	Sys     any         // FileInfo.Sys
}
```

​	MapFile描述MapFS中的单个文件。