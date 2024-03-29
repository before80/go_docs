+++
title = "gspath"
date = 2024-03-21T17:57:18+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gspath](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gspath)

Package gspath implements file index and search for folders.

​	软件包 gspath 实现文件索引和文件夹搜索。

It searches file internally with high performance in order by the directory adding sequence. Note that: If caching feature enabled, there would be a searching delay after adding/deleting files.

​	它按目录添加顺序在内部以高性能搜索文件。请注意：如果启用了缓存功能，则添加/删除文件后会出现搜索延迟。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Search

```go
func Search(root string, name string, indexFiles ...string) (filePath string, isDir bool)
```

Search searches file `name` under path `root`. The parameter `root` should be an absolute path. It will not automatically convert `root` to absolute path for performance reason. The optional parameter `indexFiles` specifies the searching index files when the result is a directory. For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found, or else it returns `filePath`.

​	搜索搜索路径 `root` 下的文件 `name` 。该参数 `root` 应为绝对路径。出于性能原因，它不会自动转换为 `root` 绝对路径。可选参数 `indexFiles` 指定当结果为目录时搜索索引文件。例如，如果结果 `filePath` 是一个目录，并且是 `indexFiles` [index.html， main.html]，它还将在 `filePath` 下搜索 [index.html， main.html]。如果找到其中任何一个，它将返回绝对文件路径，否则返回 `filePath` 。

#### func SearchWithCache

```go
func SearchWithCache(root string, name string, indexFiles ...string) (filePath string, isDir bool)
```

SearchWithCache searches file `name` under path `root` with cache feature enabled. The parameter `root` should be an absolute path. It will not automatically convert `root` to absolute path for performance reason. The optional parameter `indexFiles` specifies the searching index files when the result is a directory. For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found, or else it returns `filePath`.

​	SearchWithCache 在启用了缓存功能的路径 `root` 下搜索文件 `name` 。该参数 `root` 应为绝对路径。出于性能原因，它不会自动转换为 `root` 绝对路径。可选参数 `indexFiles` 指定当结果为目录时搜索索引文件。例如，如果结果 `filePath` 是一个目录，并且是 `indexFiles` [index.html， main.html]，它还将在 `filePath` 下搜索 [index.html， main.html]。如果找到其中任何一个，它将返回绝对文件路径，否则返回 `filePath` 。

## 类型

### type SPath

```go
type SPath struct {
	// contains filtered or unexported fields
}
```

SPath manages the path searching feature.

​	SPath 管理路径搜索功能。

#### func Get

```go
func Get(root string, cache bool) *SPath
```

Get creates and returns an instance of searching manager for given path. The parameter `cache` specifies whether using cache feature for this manager. If cache feature is enabled, it asynchronously and recursively scans the path and updates all sub files/folders to the cache using package gfsnotify.

​	Get 创建并返回给定路径的搜索管理器的实例。该参数 `cache` 指定是否为此管理器使用缓存功能。如果启用了缓存功能，它会异步和递归扫描路径，并使用包 gfsnotify 将所有子文件/文件夹更新到缓存中。

#### func New

```go
func New(path string, cache bool) *SPath
```

New creates and returns a new path searching manager.

​	new 创建并返回新的路径搜索管理器。

#### (*SPath) Add

```go
func (sp *SPath) Add(path string) (realPath string, err error)
```

Add adds more searching directory to the manager. The manager will search file in added order.

​	Add 将更多搜索目录添加到管理器中。管理器将按添加的顺序搜索文件。

#### (*SPath) AllPaths

```go
func (sp *SPath) AllPaths() []string
```

AllPaths returns all paths cached in the manager.

​	AllPaths 返回管理器中缓存的所有路径。

#### (*SPath) Paths

```go
func (sp *SPath) Paths() []string
```

Paths returns all searching directories.

​	Paths 返回所有搜索目录。

#### (*SPath) Remove

```go
func (sp *SPath) Remove(path string)
```

Remove deletes the `path` from cache files of the manager. The parameter `path` can be either an absolute path or just a relative file name.

​	删除 `path` 从管理器的缓存文件中删除。该参数 `path` 可以是绝对路径，也可以只是相对文件名。

#### (*SPath) Search

```go
func (sp *SPath) Search(name string, indexFiles ...string) (filePath string, isDir bool)
```

Search searches file `name` in the manager. The optional parameter `indexFiles` specifies the searching index files when the result is a directory. For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found, or else it returns `filePath`.

​	搜索 在管理器中搜索文件 `name` 。可选参数 `indexFiles` 指定当结果为目录时搜索索引文件。例如，如果结果 `filePath` 是一个目录，并且是 `indexFiles` [index.html， main.html]，它还会在 下 `filePath` 搜索 [index.html， main.html]。如果找到其中任何一个，它将返回绝对文件路径，否则返回 `filePath` 。

#### (*SPath) Set

```go
func (sp *SPath) Set(path string) (realPath string, err error)
```

Set deletes all other searching directories and sets the searching directory for this manager.

​	Set 将删除所有其他搜索目录，并设置此管理器的搜索目录。

#### (*SPath) Size

```go
func (sp *SPath) Size() int
```

Size returns the count of the searching directories.

​	Size 返回搜索目录的计数。

### type SPathCacheItem

```go
type SPathCacheItem struct {
	// contains filtered or unexported fields
}
```

SPathCacheItem is a cache item for searching.

​	SPathCacheItem 是用于搜索的缓存项。