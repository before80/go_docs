# decodemeta

https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1






  
  
  
  



  
  
  
  
  

## 常量 [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#pkg-constants)

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type [CoverageMetaDataDecoder](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decode.go;l=22) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaDataDecoder)

```
type CoverageMetaDataDecoder struct {
	// contains filtered or unexported fields
}
```

#### func [NewCoverageMetaDataDecoder](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decode.go;l=30) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#NewCoverageMetaDataDecoder)

```
func NewCoverageMetaDataDecoder(b []byte, readonly bool) (*CoverageMetaDataDecoder, error)
```

#### (*CoverageMetaDataDecoder) [ModulePath](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decode.go;l=74) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaDataDecoder.ModulePath)

```
func (d *CoverageMetaDataDecoder) ModulePath() string
```

#### (*CoverageMetaDataDecoder) [NumFuncs](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decode.go;l=78) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaDataDecoder.NumFuncs)

```
func (d *CoverageMetaDataDecoder) NumFuncs() uint32
```

#### (*CoverageMetaDataDecoder) [PackageName](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decode.go;l=70) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaDataDecoder.PackageName)

```
func (d *CoverageMetaDataDecoder) PackageName() string
```

#### (*CoverageMetaDataDecoder) [PackagePath](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decode.go;l=66) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaDataDecoder.PackagePath)

```
func (d *CoverageMetaDataDecoder) PackagePath() string
```

#### (*CoverageMetaDataDecoder) [ReadFunc](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decode.go;l=84) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaDataDecoder.ReadFunc)

```
func (d *CoverageMetaDataDecoder) ReadFunc(fidx uint32, f *coverage.FuncDesc) error
```

ReadFunc reads the coverage meta-data for the function with index 'findex', filling it into the FuncDesc pointed to by 'f'.

### type [CoverageMetaFileReader](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decodefile.go;l=27) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaFileReader)

```
type CoverageMetaFileReader struct {
	// contains filtered or unexported fields
}
```

CoverageMetaFileReader provides state and methods for reading a meta-data file from a code coverage run.

#### func [NewCoverageMetaFileReader](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decodefile.go;l=45) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#NewCoverageMetaFileReader)

```
func NewCoverageMetaFileReader(f *os.File, fileView []byte) (*CoverageMetaFileReader, error)
```

NewCoverageMetaFileReader returns a new helper object for reading the coverage meta-data output file 'f'. The param 'fileView' is a read-only slice containing the contents of 'f' obtained by mmap'ing the file read-only; 'fileView' may be nil, in which case the helper will read the contents of the file using regular file Read operations.

#### (*CoverageMetaFileReader) [CounterGranularity](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decodefile.go;l=153) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaFileReader.CounterGranularity)

```
func (r *CoverageMetaFileReader) CounterGranularity() coverage.CounterGranularity
```

CounterMode returns the counter granularity (single counter per function, or counter per block) selected when building for coverage for the program that produce this meta-data file.

#### (*CoverageMetaFileReader) [CounterMode](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decodefile.go;l=146) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaFileReader.CounterMode)

```
func (r *CoverageMetaFileReader) CounterMode() coverage.CounterMode
```

CounterMode returns the counter mode (set, count, atomic) used when building for coverage for the program that produce this meta-data file.

#### (*CoverageMetaFileReader) [FileHash](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decodefile.go;l=160) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaFileReader.FileHash)

```
func (r *CoverageMetaFileReader) FileHash() [16]byte
```

FileHash returns the hash computed for all of the package meta-data blobs. Coverage counter data files refer to this hash, and the hash will be encoded into the meta-data file name.

#### (*CoverageMetaFileReader) [GetPackageDecoder](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decodefile.go;l=171) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaFileReader.GetPackageDecoder)

```
func (r *CoverageMetaFileReader) GetPackageDecoder(pkIdx uint32, payloadbuf []byte) (*CoverageMetaDataDecoder, []byte, error)
```

GetPackageDecoder requests a decoder object for the package within the meta-data file whose index is 'pkIdx'. If the CoverageMetaFileReader was set up with a read-only file view, a pointer into that file view will be returned, otherwise the buffer 'payloadbuf' will be written to (or if it is not of sufficient size, a new buffer will be allocated). Return value is the decoder, a byte slice with the encoded meta-data, and an error.

#### (*CoverageMetaFileReader) [GetPackagePayload](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decodefile.go;l=194) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaFileReader.GetPackagePayload)

```
func (r *CoverageMetaFileReader) GetPackagePayload(pkIdx uint32, payloadbuf []byte) ([]byte, error)
```

GetPackagePayload returns the raw (encoded) meta-data payload for the package with index 'pkIdx'. As with GetPackageDecoder, if the CoverageMetaFileReader was set up with a read-only file view, a pointer into that file view will be returned, otherwise the buffer 'payloadbuf' will be written to (or if it is not of sufficient size, a new buffer will be allocated). Return value is the decoder, a byte slice with the encoded meta-data, and an error.

#### (*CoverageMetaFileReader) [NumPackages](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodemeta/decodefile.go;l=139) [¶](https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1#CoverageMetaFileReader.NumPackages)

```
func (r *CoverageMetaFileReader) NumPackages() uint64
```

NumPackages returns the number of packages for which this file contains meta-data.