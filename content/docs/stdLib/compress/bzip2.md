+++
title = "bzip2"
date = 2023-05-17T09:59:21+08:00
weight = 1
description = ""
isCJKLanguage = true
draft = false
+++
# bzip2

https://pkg.go.dev/compress/bzip2@go1.20.1



Package bzip2 implements bzip2 decompression.



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [NewReader](https://cs.opensource.google/go/go/+/go1.20.1:src/compress/bzip2/bzip2.go;l=46) 

``` go linenums="1"
func NewReader(r io.Reader) io.Reader
```

NewReader returns an io.Reader which decompresses bzip2 data from r. If r does not also implement io.ByteReader, the decompressor may read more data than necessary from r.

## 类型

### type [StructuralError](https://cs.opensource.google/go/go/+/go1.20.1:src/compress/bzip2/bzip2.go;l=17) 

``` go linenums="1"
type StructuralError string
```

A StructuralError is returned when the bzip2 data is found to be syntactically invalid.

#### (StructuralError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/compress/bzip2/bzip2.go;l=19) 

``` go linenums="1"
func (s StructuralError) Error() string
```