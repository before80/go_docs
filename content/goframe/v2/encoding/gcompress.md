+++
title = "gcompress"
date = 2024-03-21T17:49:19+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gcompress](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/encoding/gcompress)

Package gcompress provides kinds of compression algorithms for binary/bytes data.

​	软件包 gcompress 为二进制/字节数据提供了多种压缩算法。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

#### func Gzip

```go
func Gzip(data []byte, level ...int) ([]byte, error)
```

Gzip compresses `data` using gzip algorithm. The optional parameter `level` specifies the compression level from 1 to 9 which means from none to the best compression.

​	Gzip `data` 使用 gzip 算法进行压缩。可选参数 `level` 指定从 1 到 9 的压缩级别，这意味着从无到最佳压缩。

Note that it returns error if given `level` is invalid.

​	请注意，如果给定 `level` 无效，则返回错误。

#### func GzipFile

```go
func GzipFile(srcFilePath, dstFilePath string, level ...int) (err error)
```

GzipFile compresses the file `src` to `dst` using gzip algorithm.

​	GzipFile 使用 gzip 算法将文件 `src` 压缩。 `dst`

#### func GzipPathWriter <-2.2.0

```go
func GzipPathWriter(filePath string, writer io.Writer, level ...int) error
```

GzipPathWriter compresses `filePath` to `writer` using gzip compressing algorithm.

​	GzipPathWriter 压缩 `filePath` 为 `writer` 使用 gzip 压缩算法。

Note that the parameter `path` can be either a directory or a file.

​	请注意，该参数 `path` 可以是目录或文件。

#### func UnGzip

```go
func UnGzip(data []byte) ([]byte, error)
```

UnGzip decompresses `data` with gzip algorithm.

​	UnGzip `data` 使用 gzip 算法解压缩。

#### func UnGzipFile

```go
func UnGzipFile(srcFilePath, dstFilePath string) error
```

UnGzipFile decompresses srcFilePath `src` to `dst` using gzip algorithm.

​	UnGzipFile 将 srcFilePath `src` 解压缩为 `dst` 使用 gzip 算法。

#### func UnZipContent

```go
func UnZipContent(zippedContent []byte, dstFolderPath string, zippedPrefix ...string) error
```

UnZipContent decompresses `zippedContent` to `dstFolderPath` using zip compressing algorithm.

​	UnZipContent 解压缩 `zippedContent` 为 `dstFolderPath` 使用 zip 压缩算法。

The parameter `dstFolderPath` should be a directory. The parameter `zippedPrefix` specifies the unzipped path of `zippedContent`, which can be used to specify part of the archive file to unzip.

​	该参数 `dstFolderPath` 应为目录。该参数 `zippedPrefix` 指定 `zippedContent` 的解压缩路径，该路径可用于指定要解压缩的归档文件的一部分。

#### func UnZipFile

```go
func UnZipFile(zippedFilePath, dstFolderPath string, zippedPrefix ...string) error
```

UnZipFile decompresses `archive` to `dstFolderPath` using zip compressing algorithm.

​	UnZipFile 解压缩 `archive` 为 `dstFolderPath` 使用 zip 压缩算法。

The parameter `dstFolderPath` should be a directory. The optional parameter `zippedPrefix` specifies the unzipped path of `zippedFilePath`, which can be used to specify part of the archive file to unzip.

​	该参数 `dstFolderPath` 应为目录。可选参数 `zippedPrefix` 指定 `zippedFilePath` 的解压缩路径，该路径可用于指定要解压缩的归档文件的一部分。

#### func UnZlib

```go
func UnZlib(data []byte) ([]byte, error)
```

UnZlib decompresses `data` with zlib algorithm.

​	UnZlib `data` 使用 zlib 算法解压缩。

#### func ZipPath

```go
func ZipPath(fileOrFolderPaths, dstFilePath string, prefix ...string) error
```

ZipPath compresses `fileOrFolderPaths` to `dstFilePath` using zip compressing algorithm.

​	ZipPath 压缩 `fileOrFolderPaths` 为 `dstFilePath` 使用 zip 压缩算法。

The parameter `paths` can be either a directory or a file, which supports multiple paths join with ‘,’. The unnecessary parameter `prefix` indicates the path prefix for zip file.

​	该参数 `paths` 可以是目录或文件，支持多个路径与“，”连接。不必要的参数 `prefix` 指示 zip 文件的路径前缀。

#### func ZipPathContent <-2.2.0

```go
func ZipPathContent(fileOrFolderPaths string, prefix ...string) ([]byte, error)
```

ZipPathContent compresses `fileOrFolderPaths` to []byte using zip compressing algorithm.

​	ZipPathContent 使用 zip 压缩算法 `fileOrFolderPaths` 压缩到 []byte。

Note that the parameter `fileOrFolderPaths` can be either a directory or a file, which supports multiple paths join with ‘,’. The unnecessary parameter `prefix` indicates the path prefix for zip file.

​	请注意，该参数 `fileOrFolderPaths` 可以是目录或文件，它支持多个路径与“，”连接。不必要的参数 `prefix` 指示 zip 文件的路径前缀。

#### func ZipPathWriter

```go
func ZipPathWriter(fileOrFolderPaths string, writer io.Writer, prefix ...string) error
```

ZipPathWriter compresses `fileOrFolderPaths` to `writer` using zip compressing algorithm.

​	ZipPathWriter 压缩 `fileOrFolderPaths` 为 `writer` 使用 zip 压缩算法。

Note that the parameter `fileOrFolderPaths` can be either a directory or a file, which supports multiple paths join with ‘,’. The unnecessary parameter `prefix` indicates the path prefix for zip file.

​	请注意，该参数 `fileOrFolderPaths` 可以是目录或文件，它支持多个路径与“，”连接。不必要的参数 `prefix` 指示 zip 文件的路径前缀。

#### func Zlib

```go
func Zlib(data []byte) ([]byte, error)
```

Zlib compresses `data` with zlib algorithm.

​	Zlib `data` 使用 zlib 算法进行压缩。

## 类型

This section is empty.