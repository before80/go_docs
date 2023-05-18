+++
title = "suffixarray"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# suffixarray

https://pkg.go.dev/index/suffixarray@go1.20.1



Package suffixarray implements substring search in logarithmic time using an in-memory suffix array.

Example use:

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

### type [Index](https://cs.opensource.google/go/go/+/go1.20.1:src/index/suffixarray/suffixarray.go;l=34) 

``` go 
type Index struct {
	// contains filtered or unexported fields
}
```

Index implements a suffix array for fast substring search.

#### func [New](https://cs.opensource.google/go/go/+/go1.20.1:src/index/suffixarray/suffixarray.go;l=75) 

``` go 
func New(data []byte) *Index
```

New creates a new Index for data. Index creation time is O(N) for N = len(data).

#### (*Index) [Bytes](https://cs.opensource.google/go/go/+/go1.20.1:src/index/suffixarray/suffixarray.go;l=232) 

``` go 
func (x *Index) Bytes() []byte
```

Bytes returns the data over which the index was created. It must not be modified.

#### (*Index) [FindAllIndex](https://cs.opensource.google/go/go/+/go1.20.1:src/index/suffixarray/suffixarray.go;l=286) 

``` go 
func (x *Index) FindAllIndex(r *regexp.Regexp, n int) (result [][]int)
```

FindAllIndex returns a sorted list of non-overlapping matches of the regular expression r, where a match is a pair of indices specifying the matched slice of x.Bytes(). If n < 0, all matches are returned in successive order. Otherwise, at most n matches are returned and they may not be successive. The result is nil if there are no matches, or if n == 0.

#### (*Index) [Lookup](https://cs.opensource.google/go/go/+/go1.20.1:src/index/suffixarray/suffixarray.go;l=256) 

``` go 
func (x *Index) Lookup(s []byte, n int) (result []int)
```

Lookup returns an unsorted list of at most n indices where the byte string s occurs in the indexed data. If n < 0, all occurrences are returned. The result is nil if s is empty, s is not found, or n == 0. Lookup time is O(log(N)*len(s) + len(result)) where N is the size of the indexed data.

##### Example
``` go 
```

#### (*Index) [Read](https://cs.opensource.google/go/go/+/go1.20.1:src/index/suffixarray/suffixarray.go;l=154) 

``` go 
func (x *Index) Read(r io.Reader) error
```

Read reads the index from r into x; x must not be nil.

#### (*Index) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/index/suffixarray/suffixarray.go;l=204) 

``` go 
func (x *Index) Write(w io.Writer) error
```

Write writes the index x to w.