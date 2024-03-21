+++
title = "gcompress"
date = 2024-03-21T17:49:19+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gcompress

Package gcompress provides kinds of compression algorithms for binary/bytes data.

Package gcompress provides kinds of compression algorithms for binary/bytes data.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

##### func Gzip 

``` go
func Gzip(data []byte, level ...int) ([]byte, error)
```

Gzip compresses `data` using gzip algorithm. The optional parameter `level` specifies the compression level from 1 to 9 which means from none to the best compression.

Note that it returns error if given `level` is invalid.

##### func GzipFile 

``` go
func GzipFile(srcFilePath, dstFilePath string, level ...int) (err error)
```

GzipFile compresses the file `src` to `dst` using gzip algorithm.

##### func GzipPathWriter <-2.2.0

``` go
func GzipPathWriter(filePath string, writer io.Writer, level ...int) error
```

GzipPathWriter compresses `filePath` to `writer` using gzip compressing algorithm.

Note that the parameter `path` can be either a directory or a file.

##### func UnGzip 

``` go
func UnGzip(data []byte) ([]byte, error)
```

UnGzip decompresses `data` with gzip algorithm.

##### func UnGzipFile 

``` go
func UnGzipFile(srcFilePath, dstFilePath string) error
```

UnGzipFile decompresses srcFilePath `src` to `dst` using gzip algorithm.

##### func UnZipContent 

``` go
func UnZipContent(zippedContent []byte, dstFolderPath string, zippedPrefix ...string) error
```

UnZipContent decompresses `zippedContent` to `dstFolderPath` using zip compressing algorithm.

The parameter `dstFolderPath` should be a directory. The parameter `zippedPrefix` specifies the unzipped path of `zippedContent`, which can be used to specify part of the archive file to unzip.

##### func UnZipFile 

``` go
func UnZipFile(zippedFilePath, dstFolderPath string, zippedPrefix ...string) error
```

UnZipFile decompresses `archive` to `dstFolderPath` using zip compressing algorithm.

The parameter `dstFolderPath` should be a directory. The optional parameter `zippedPrefix` specifies the unzipped path of `zippedFilePath`, which can be used to specify part of the archive file to unzip.

##### func UnZlib 

``` go
func UnZlib(data []byte) ([]byte, error)
```

UnZlib decompresses `data` with zlib algorithm.

##### func ZipPath 

``` go
func ZipPath(fileOrFolderPaths, dstFilePath string, prefix ...string) error
```

ZipPath compresses `fileOrFolderPaths` to `dstFilePath` using zip compressing algorithm.

The parameter `paths` can be either a directory or a file, which supports multiple paths join with ','. The unnecessary parameter `prefix` indicates the path prefix for zip file.

##### func ZipPathContent <-2.2.0

``` go
func ZipPathContent(fileOrFolderPaths string, prefix ...string) ([]byte, error)
```

ZipPathContent compresses `fileOrFolderPaths` to []byte using zip compressing algorithm.

Note that the parameter `fileOrFolderPaths` can be either a directory or a file, which supports multiple paths join with ','. The unnecessary parameter `prefix` indicates the path prefix for zip file.

##### func ZipPathWriter 

``` go
func ZipPathWriter(fileOrFolderPaths string, writer io.Writer, prefix ...string) error
```

ZipPathWriter compresses `fileOrFolderPaths` to `writer` using zip compressing algorithm.

Note that the parameter `fileOrFolderPaths` can be either a directory or a file, which supports multiple paths join with ','. The unnecessary parameter `prefix` indicates the path prefix for zip file.

##### func Zlib 

``` go
func Zlib(data []byte) ([]byte, error)
```

Zlib compresses `data` with zlib algorithm.

### Types 

This section is empty.