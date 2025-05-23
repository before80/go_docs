# bytealg

> 原文：[https://pkg.go.dev/internal/bytealg@go1.24.2](https://pkg.go.dev/internal/bytealg@go1.24.2)





















## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/index_amd64.go;l=9)

``` go
const MaxBruteForce = 64
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/bytealg.go;l=32)

``` go
const PrimeRK = 16777619
```

PrimeRK is the prime base used in Rabin-Karp algorithm.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/bytealg.go;l=25)

``` go
var MaxLen int
```

MaxLen is the maximum length of the string to be searched for (argument b) in Index. If MaxLen is not 0, make sure MaxLen >= 4.

## 函数

#### func Compare 

``` go
func Compare(a, b []byte) int
```

#### func Count 

``` go
func Count(b []byte, c byte) int
```

#### func CountString 

``` go
func CountString(s string, c byte) int
```

#### func Cutover 

``` go
func Cutover(n int) int
```

Cutover reports the number of failures of IndexByte we should tolerate before switching over to Index. n is the number of bytes processed so far. See the bytes.Index implementation for details.

#### func Equal 

``` go
func Equal(a, b []byte) bool
```

Equal reports whether a and b are the same length and contain the same bytes. A nil argument is equivalent to an empty slice.

Equal is equivalent to bytes.Equal. It is provided here for convenience, because some packages cannot depend on bytes.

#### func HashStr  <- go1.15

``` go
func HashStr(sep string) (uint32, uint32)
```

HashStr returns the hash and the appropriate multiplicative factor for use in Rabin-Karp algorithm.

#### func HashStrBytes  <- go1.15

``` go
func HashStrBytes(sep []byte) (uint32, uint32)
```

HashStrBytes returns the hash and the appropriate multiplicative factor for use in Rabin-Karp algorithm.

#### func HashStrRev  <- go1.15

``` go
func HashStrRev(sep string) (uint32, uint32)
```

HashStrRev returns the hash of the reverse of sep and the appropriate multiplicative factor for use in Rabin-Karp algorithm.

#### func HashStrRevBytes  <- go1.15

``` go
func HashStrRevBytes(sep []byte) (uint32, uint32)
```

HashStrRevBytes returns the hash of the reverse of sep and the appropriate multiplicative factor for use in Rabin-Karp algorithm.

#### func Index 

``` go
func Index(a, b []byte) int
```

Index returns the index of the first instance of b in a, or -1 if b is not present in a. Requires 2 <= len(b) <= MaxLen.

#### func IndexByte 

``` go
func IndexByte(b []byte, c byte) int
```

#### func IndexByteString 

``` go
func IndexByteString(s string, c byte) int
```

#### func IndexRabinKarp  <- go1.15

``` go
func IndexRabinKarp(s, substr string) int
```

IndexRabinKarp uses the Rabin-Karp search algorithm to return the index of the first occurrence of substr in s, or -1 if not present.

#### func IndexRabinKarpBytes  <- go1.15

``` go
func IndexRabinKarpBytes(s, sep []byte) int
```

IndexRabinKarpBytes uses the Rabin-Karp search algorithm to return the index of the first occurrence of substr in s, or -1 if not present.

#### func IndexString 

``` go
func IndexString(a, b string) int
```

IndexString returns the index of the first instance of b in a, or -1 if b is not present in a. Requires 2 <= len(b) <= MaxLen.

## 类型

This section is empty.