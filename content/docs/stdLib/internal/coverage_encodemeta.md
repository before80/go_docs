# encodemeta

https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1







  




## 常量 [¶](https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1#pkg-constants)

This section is empty.

## 变量

This section is empty.

## 函数

#### func [HashFuncDesc](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/encodemeta/encode.go;l=190) [¶](https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1#HashFuncDesc)

```
func HashFuncDesc(f *coverage.FuncDesc) [16]byte
```

HashFuncDesc computes an md5 sum of a coverage.FuncDesc and returns a digest for it.

## 类型

### type [CoverageMetaDataBuilder](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/encodemeta/encode.go;l=24) [¶](https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1#CoverageMetaDataBuilder)

```
type CoverageMetaDataBuilder struct {
	// contains filtered or unexported fields
}
```

#### func [NewCoverageMetaDataBuilder](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/encodemeta/encode.go;l=36) [¶](https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1#NewCoverageMetaDataBuilder)

```
func NewCoverageMetaDataBuilder(pkgpath string, pkgname string, modulepath string) (*CoverageMetaDataBuilder, error)
```

#### (*CoverageMetaDataBuilder) [AddFunc](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/encodemeta/encode.go;l=67) [¶](https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1#CoverageMetaDataBuilder.AddFunc)

```
func (b *CoverageMetaDataBuilder) AddFunc(f coverage.FuncDesc) uint
```

AddFunc registers a new function with the meta data builder.

#### (*CoverageMetaDataBuilder) [Emit](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/encodemeta/encode.go;l=131) [¶](https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1#CoverageMetaDataBuilder.Emit)

```
func (b *CoverageMetaDataBuilder) Emit(w io.WriteSeeker) ([16]byte, error)
```

Emit writes the meta-data accumulated so far in this builder to 'w'. Returns a hash of the meta-data payload and an error.

### type [CoverageMetaFileWriter](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/encodemeta/encodefile.go;l=23) [¶](https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1#CoverageMetaFileWriter)

```
type CoverageMetaFileWriter struct {
	// contains filtered or unexported fields
}
```

#### func [NewCoverageMetaFileWriter](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/encodemeta/encodefile.go;l=31) [¶](https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1#NewCoverageMetaFileWriter)

```
func NewCoverageMetaFileWriter(mfname string, w io.Writer) *CoverageMetaFileWriter
```

#### (*CoverageMetaFileWriter) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/encodemeta/encodefile.go;l=42) [¶](https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1#CoverageMetaFileWriter.Write)

```
func (m *CoverageMetaFileWriter) Write(finalHash [16]byte, blobs [][]byte, mode coverage.CounterMode, granularity coverage.CounterGranularity) error
```