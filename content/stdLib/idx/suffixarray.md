+++
title = "suffixarray"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/index/suffixarray@go1.23.0](https://pkg.go.dev/index/suffixarray@go1.23.0)

Package suffixarray implements substring search in logarithmic time using an in-memory suffix array.

​	suffixarray 包实现使用内存中后缀数组的对数时间子字符串搜索。

Example use:

​	示例用法：

```
// create index for some data
index := suffixarray.New(data)

// lookup byte slice s
offsets1 := index.Lookup(s, -1) // the list of all indices where s occurs in data
offsets2 := index.Lookup(s, 3)  // the list of at most 3 indices where s occurs in data
```

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Index 

``` go 
type Index struct {
	// contains filtered or unexported fields
}
```

Index implements a suffix array for fast substring search.

​	Index 实现后缀数组以进行快速子字符串搜索。

#### func New

```go
func New(data []byte) *Index
```

New creates a new Index for data. Index creation time is O(N) for N = len(data).

​	New 为数据创建一个新的 Index。对于 N = len(data)，Index 创建时间为 O(N)。

#### (*Index) Bytes

```go
func (x *Index) Bytes() []byte
```

Bytes returns the data over which the index was created. It must not be modified.

​	Bytes 返回创建索引的数据。不得修改它。

#### (*Index) FindAllIndex

```go
func (x *Index) FindAllIndex(r *regexp.Regexp, n int) (result [][]int)
```

FindAllIndex returns a sorted list of non-overlapping matches of the regular expression r, where a match is a pair of indices specifying the matched slice of x.Bytes(). If n < 0, all matches are returned in successive order. Otherwise, at most n matches are returned and they may not be successive. The result is nil if there are no matches, or if n == 0.

​	FindAllIndex 返回正则表达式 r 的非重叠匹配项的有序列表，其中匹配项是一对指定 x.Bytes() 的匹配切片的索引。如果 n < 0，则按连续顺序返回所有匹配项。否则，最多返回 n 个匹配项，并且它们可能不是连续的。如果没有匹配项或 n == 0，则结果为 nil。

#### (*Index) Lookup

```go
func (x *Index) Lookup(s []byte, n int) (result []int)
```

Lookup returns an unsorted list of at most n indices where the byte string s occurs in the indexed data. If n < 0, all occurrences are returned. The result is nil if s is empty, s is not found, or n == 0. Lookup time is O(log(N)*len(s) + len(result)) where N is the size of the indexed data.

​	Lookup 返回一个最多包含 n 个索引的未排序列表，其中字节字符串 s 出现在索引数据中。如果 n < 0，则返回所有出现位置。如果 s 为空、未找到 s 或 n == 0，则结果为 nil。Lookup 时间为 O(log(N)*len(s) + len(result))，其中 N 为索引数据的大小。

##### Lookup Example

```go
package main

import (
	"fmt"
	"index/suffixarray"
)

func main() {
	index := suffixarray.New([]byte("banana"))
	offsets := index.Lookup([]byte("ana"), -1)
	for _, off := range offsets {
		fmt.Println(off)
	}

}

Output:

1
3
```

#### (*Index) Read

```go
func (x *Index) Read(r io.Reader) error
```

Read reads the index from r into x; x must not be nil.

​	Read 从 r 中读取索引到 x；x 不能为 nil。

#### (*Index) Write

```go
func (x *Index) Write(w io.Writer) error
```

Write writes the index x to w.

​	Write 将索引 x 写入 w。