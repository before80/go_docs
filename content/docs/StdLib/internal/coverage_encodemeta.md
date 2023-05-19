# encodemeta

https://pkg.go.dev/internal/coverage/encodemeta@go1.20.1







  




## 常量 ¶

This section is empty.

## 变量

This section is empty.

## 函数

#### func HashFuncDesc 

```
func HashFuncDesc(f *coverage.FuncDesc) [16]byte
```

HashFuncDesc computes an md5 sum of a coverage.FuncDesc and returns a digest for it.

## 类型

### type CoverageMetaDataBuilder 

```
type CoverageMetaDataBuilder struct {
	// contains filtered or unexported fields
}
```

#### func NewCoverageMetaDataBuilder 

```
func NewCoverageMetaDataBuilder(pkgpath string, pkgname string, modulepath string) (*CoverageMetaDataBuilder, error)
```

#### (*CoverageMetaDataBuilder) AddFunc 

```
func (b *CoverageMetaDataBuilder) AddFunc(f coverage.FuncDesc) uint
```

AddFunc registers a new function with the meta data builder.

#### (*CoverageMetaDataBuilder) Emit 

```
func (b *CoverageMetaDataBuilder) Emit(w io.WriteSeeker) ([16]byte, error)
```

Emit writes the meta-data accumulated so far in this builder to 'w'. Returns a hash of the meta-data payload and an error.

### type CoverageMetaFileWriter 

```
type CoverageMetaFileWriter struct {
	// contains filtered or unexported fields
}
```

#### func NewCoverageMetaFileWriter 

```
func NewCoverageMetaFileWriter(mfname string, w io.Writer) *CoverageMetaFileWriter
```

#### (*CoverageMetaFileWriter) Write 

```
func (m *CoverageMetaFileWriter) Write(finalHash [16]byte, blobs [][]byte, mode coverage.CounterMode, granularity coverage.CounterGranularity) error
```