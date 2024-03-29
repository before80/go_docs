+++
title = "gfpool"
date = 2024-03-21T17:55:35+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gfpool](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gfpool)

Package gfpool provides io-reusable pool for file pointer.

​	软件包 gfpool 为文件指针提供 io-reusable 池。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type File

```go
type File struct {
	*os.File // Underlying file pointer.
	// contains filtered or unexported fields
}
```

File is an item in the pool.

​	文件是池中的一个项目。

#### func Get <-2.1.1

```go
func Get(path string, flag int, perm os.FileMode, ttl ...time.Duration) (file *File)
```

Get returns a file item with given file path, flag and opening permission. It retrieves a file item from the file pointer pool after then.

​	Get 返回具有给定文件路径、标志和打开权限的文件项。之后，它从文件指针池中检索文件项。

#### func Open

```go
func Open(path string, flag int, perm os.FileMode, ttl ...time.Duration) (file *File, err error)
```

Open creates and returns a file item with given file path, flag and opening permission. It automatically creates an associated file pointer pool internally when it’s called first time. It retrieves a file item from the file pointer pool after then.

​	Open 创建并返回具有给定文件路径、标志和打开权限的文件项。首次调用时，它会自动在内部创建关联的文件指针池。之后，它从文件指针池中检索文件项。

#### (*File) Close

```go
func (f *File) Close(close ...bool) error
```

Close puts the file pointer back to the file pointer pool.

​	“关闭”将文件指针放回文件指针池。

#### (*File) Stat

```go
func (f *File) Stat() (os.FileInfo, error)
```

Stat returns the FileInfo structure describing file.

​	Stat 返回描述文件的 FileInfo 结构。

### type Pool

```go
type Pool struct {
	// contains filtered or unexported fields
}
```

Pool pointer pool.

​	池指针池。

#### func New

```go
func New(path string, flag int, perm os.FileMode, ttl ...time.Duration) *Pool
```

New creates and returns a file pointer pool with given file path, flag and opening permission.

​	New 创建并返回具有给定文件路径、标志和打开权限的文件指针池。

Note the expiration logic: ttl = 0 : not expired; ttl < 0 : immediate expired after use; ttl > 0 : timeout expired; It is not expired in default.

​	注意过期逻辑：ttl = 0 ： 未过期;ttl < 0：使用后立即过期; ttl > 0：超时过期;默认情况下，它不会过期。

#### (*Pool) Close

```go
func (p *Pool) Close()
```

Close closes current file pointer pool.

​	关闭关闭当前文件指针池。

#### (*Pool) File

```go
func (p *Pool) File() (*File, error)
```

File retrieves file item from the file pointer pool and returns it. It creates one if the file pointer pool is empty. Note that it should be closed when it will never be used. When it’s closed, it is not really closed the underlying file pointer but put back to the file pointer pool.

​	File 从文件指针池中检索文件项并返回它。如果文件指针池为空，它将创建一个。请注意，当它永远不会被使用时，它应该被关闭。当它关闭时，它并没有真正关闭基础文件指针，而是放回文件指针池。