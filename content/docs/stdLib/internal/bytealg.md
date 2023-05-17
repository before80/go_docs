# bytealg

https://pkg.go.dev/internal/bytealg@go1.20.1





















## 常量 [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#pkg-constants)

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/index_amd64.go;l=9)

```
const MaxBruteForce = 64
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/bytealg.go;l=32)

```
const PrimeRK = 16777619
```

PrimeRK is the prime base used in Rabin-Karp algorithm.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/bytealg.go;l=25)

```
var MaxLen int
```

MaxLen is the maximum length of the string to be searched for (argument b) in Index. If MaxLen is not 0, make sure MaxLen >= 4.

## 函数

#### func [Compare](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/compare_native.go;l=12) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#Compare)

```
func Compare(a, b []byte) int
```

#### func [Count](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/count_native.go;l=10) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#Count)

```
func Count(b []byte, c byte) int
```

#### func [CountString](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/count_native.go;l=13) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#CountString)

```
func CountString(s string, c byte) int
```

#### func [Cutover](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/index_amd64.go;l=23) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#Cutover)

```
func Cutover(n int) int
```

Cutover reports the number of failures of IndexByte we should tolerate before switching over to Index. n is the number of bytes processed so far. See the bytes.Index implementation for details.

#### func [Equal](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/equal_generic.go;l=14) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#Equal)

```
func Equal(a, b []byte) bool
```

Equal reports whether a and b are the same length and contain the same bytes. A nil argument is equivalent to an empty slice.

Equal is equivalent to bytes.Equal. It is provided here for convenience, because some packages cannot depend on bytes.

#### func [HashStr](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/bytealg.go;l=53) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#HashStr)added in go1.15

```
func HashStr(sep string) (uint32, uint32)
```

HashStr returns the hash and the appropriate multiplicative factor for use in Rabin-Karp algorithm.

#### func [HashStrBytes](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/bytealg.go;l=36) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#HashStrBytes)added in go1.15

```
func HashStrBytes(sep []byte) (uint32, uint32)
```

HashStrBytes returns the hash and the appropriate multiplicative factor for use in Rabin-Karp algorithm.

#### func [HashStrRev](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/bytealg.go;l=87) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#HashStrRev)added in go1.15

```
func HashStrRev(sep string) (uint32, uint32)
```

HashStrRev returns the hash of the reverse of sep and the appropriate multiplicative factor for use in Rabin-Karp algorithm.

#### func [HashStrRevBytes](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/bytealg.go;l=70) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#HashStrRevBytes)added in go1.15

```
func HashStrRevBytes(sep []byte) (uint32, uint32)
```

HashStrRevBytes returns the hash of the reverse of sep and the appropriate multiplicative factor for use in Rabin-Karp algorithm.

#### func [Index](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/index_native.go;l=13) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#Index)

```
func Index(a, b []byte) int
```

Index returns the index of the first instance of b in a, or -1 if b is not present in a. Requires 2 <= len(b) <= MaxLen.

#### func [IndexByte](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/indexbyte_native.go;l=10) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#IndexByte)

```
func IndexByte(b []byte, c byte) int
```

#### func [IndexByteString](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/indexbyte_native.go;l=13) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#IndexByteString)

```
func IndexByteString(s string, c byte) int
```

#### func [IndexRabinKarp](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/bytealg.go;l=129) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#IndexRabinKarp)added in go1.15

```
func IndexRabinKarp(s, substr string) int
```

IndexRabinKarp uses the Rabin-Karp search algorithm to return the index of the first occurrence of substr in s, or -1 if not present.

#### func [IndexRabinKarpBytes](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/bytealg.go;l=104) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#IndexRabinKarpBytes)added in go1.15

```
func IndexRabinKarpBytes(s, sep []byte) int
```

IndexRabinKarpBytes uses the Rabin-Karp search algorithm to return the index of the first occurrence of substr in s, or -1 if not present.

#### func [IndexString](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/bytealg/index_native.go;l=19) [¶](https://pkg.go.dev/internal/bytealg@go1.20.1#IndexString)

```
func IndexString(a, b string) int
```

IndexString returns the index of the first instance of b in a, or -1 if b is not present in a. Requires 2 <= len(b) <= MaxLen.

## 类型

This section is empty.