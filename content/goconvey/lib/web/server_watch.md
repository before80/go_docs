+++
title = "server_watch"
date = 2024-12-15T21:22:39+08:00
weight = 8
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/watch](https://pkg.go.dev/github.com/smartystreets/goconvey@v1.8.1/web/server/watch)
>
> 收录该文档时间： `2024-12-15T21:22:39+08:00`

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func ActiveFolders 

``` go
func ActiveFolders(folders messaging.Folders) messaging.Folders
```

### func AttachProfiles 

``` go
func AttachProfiles(folders messaging.Folders, items []*FileSystemItem)
```

### func CreateFolders 

``` go
func CreateFolders(items []*FileSystemItem) messaging.Folders
```

### func LimitDepth 

``` go
func LimitDepth(folders messaging.Folders, depth int)
```

### func MarkIgnored 

``` go
func MarkIgnored(folders messaging.Folders, ignored map[string]struct{})
```

### func ParseProfile 

``` go
func ParseProfile(profile string) (isDisabled bool, tags, arguments []string)
```

### func ReadContents 

``` go
func ReadContents(path string) string
```

ReadContents reads files wholesale. This function is only called on files that end in '.goconvey'. These files should be very small, probably not ever more than a few hundred bytes. The ignored errors are ok because in the event of an IO error all that need be returned is an empty string.

### func Sum 

``` go
func Sum(folders messaging.Folders, items []*FileSystemItem) int64
```

### func YieldFileSystemItems 

``` go
func YieldFileSystemItems(root string, excludedDirs []string) chan *FileSystemItem
```

## 类型

### type FileSystemItem 

``` go
type FileSystemItem struct {
	Root     string
	Path     string
	Name     string
	Size     int64
	Modified int64
	IsFolder bool

	ProfileDisabled  bool
	ProfileTags      []string
	ProfileArguments []string
}
```

#### func Categorize 

``` go
func Categorize(items chan *FileSystemItem, root string, watchSuffixes []string) (folders, profiles, goFiles []*FileSystemItem)
```

### type Watcher 

``` go
type Watcher struct {
	// contains filtered or unexported fields
}
```

#### func NewWatcher 

``` go
func NewWatcher(rootFolder string, folderDepth int, nap time.Duration,
	input chan messaging.WatcherCommand, output chan messaging.Folders, watchSuffixes string, excludedDirs []string) *Watcher
```

#### (*Watcher) Listen 

``` go
func (this *Watcher) Listen()
```
