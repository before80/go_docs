+++
title = "gspath"
date = 2024-03-21T17:57:18+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gspath

Package gspath implements file index and search for folders.

It searches file internally with high performance in order by the directory adding sequence. Note that: If caching feature enabled, there would be a searching delay after adding/deleting files.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Search 

``` go
func Search(root string, name string, indexFiles ...string) (filePath string, isDir bool)
```

Search searches file `name` under path `root`. The parameter `root` should be an absolute path. It will not automatically convert `root` to absolute path for performance reason. The optional parameter `indexFiles` specifies the searching index files when the result is a directory. For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found, or else it returns `filePath`.

##### func SearchWithCache 

``` go
func SearchWithCache(root string, name string, indexFiles ...string) (filePath string, isDir bool)
```

SearchWithCache searches file `name` under path `root` with cache feature enabled. The parameter `root` should be an absolute path. It will not automatically convert `root` to absolute path for performance reason. The optional parameter `indexFiles` specifies the searching index files when the result is a directory. For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found, or else it returns `filePath`.

### Types 

#### type SPath 

``` go
type SPath struct {
	// contains filtered or unexported fields
}
```

SPath manages the path searching feature.

##### func Get 

``` go
func Get(root string, cache bool) *SPath
```

Get creates and returns an instance of searching manager for given path. The parameter `cache` specifies whether using cache feature for this manager. If cache feature is enabled, it asynchronously and recursively scans the path and updates all sub files/folders to the cache using package gfsnotify.

##### func New 

``` go
func New(path string, cache bool) *SPath
```

New creates and returns a new path searching manager.

##### (*SPath) Add 

``` go
func (sp *SPath) Add(path string) (realPath string, err error)
```

Add adds more searching directory to the manager. The manager will search file in added order.

##### (*SPath) AllPaths 

``` go
func (sp *SPath) AllPaths() []string
```

AllPaths returns all paths cached in the manager.

##### (*SPath) Paths 

``` go
func (sp *SPath) Paths() []string
```

Paths returns all searching directories.

##### (*SPath) Remove 

``` go
func (sp *SPath) Remove(path string)
```

Remove deletes the `path` from cache files of the manager. The parameter `path` can be either an absolute path or just a relative file name.

##### (*SPath) Search 

``` go
func (sp *SPath) Search(name string, indexFiles ...string) (filePath string, isDir bool)
```

Search searches file `name` in the manager. The optional parameter `indexFiles` specifies the searching index files when the result is a directory. For example, if the result `filePath` is a directory, and `indexFiles` is [index.html, main.html], it will also search [index.html, main.html] under `filePath`. It returns the absolute file path if any of them found, or else it returns `filePath`.

##### (*SPath) Set 

``` go
func (sp *SPath) Set(path string) (realPath string, err error)
```

Set deletes all other searching directories and sets the searching directory for this manager.

##### (*SPath) Size 

``` go
func (sp *SPath) Size() int
```

Size returns the count of the searching directories.

#### type SPathCacheItem 

``` go
type SPathCacheItem struct {
	// contains filtered or unexported fields
}
```

SPathCacheItem is a cache item for searching.