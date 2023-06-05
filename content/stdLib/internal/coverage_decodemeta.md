# decodemeta

https://pkg.go.dev/internal/coverage/decodemeta@go1.20.1






  
  
  
  



  
  
  
  
  

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type CoverageMetaDataDecoder 

```
type CoverageMetaDataDecoder struct {
	// contains filtered or unexported fields
}
```

#### func NewCoverageMetaDataDecoder 

```
func NewCoverageMetaDataDecoder(b []byte, readonly bool) (*CoverageMetaDataDecoder, error)
```

#### (*CoverageMetaDataDecoder) ModulePath 

```
func (d *CoverageMetaDataDecoder) ModulePath() string
```

#### (*CoverageMetaDataDecoder) NumFuncs 

```
func (d *CoverageMetaDataDecoder) NumFuncs() uint32
```

#### (*CoverageMetaDataDecoder) PackageName 

```
func (d *CoverageMetaDataDecoder) PackageName() string
```

#### (*CoverageMetaDataDecoder) PackagePath 

```
func (d *CoverageMetaDataDecoder) PackagePath() string
```

#### (*CoverageMetaDataDecoder) ReadFunc 

```
func (d *CoverageMetaDataDecoder) ReadFunc(fidx uint32, f *coverage.FuncDesc) error
```

ReadFunc reads the coverage meta-data for the function with index 'findex', filling it into the FuncDesc pointed to by 'f'.

### type CoverageMetaFileReader 

```
type CoverageMetaFileReader struct {
	// contains filtered or unexported fields
}
```

CoverageMetaFileReader provides state and methods for reading a meta-data file from a code coverage run.

#### func NewCoverageMetaFileReader 

```
func NewCoverageMetaFileReader(f *os.File, fileView []byte) (*CoverageMetaFileReader, error)
```

NewCoverageMetaFileReader returns a new helper object for reading the coverage meta-data output file 'f'. The param 'fileView' is a read-only slice containing the contents of 'f' obtained by mmap'ing the file read-only; 'fileView' may be nil, in which case the helper will read the contents of the file using regular file Read operations.

#### (*CoverageMetaFileReader) CounterGranularity 

```
func (r *CoverageMetaFileReader) CounterGranularity() coverage.CounterGranularity
```

CounterMode returns the counter granularity (single counter per function, or counter per block) selected when building for coverage for the program that produce this meta-data file.

#### (*CoverageMetaFileReader) CounterMode 

```
func (r *CoverageMetaFileReader) CounterMode() coverage.CounterMode
```

CounterMode returns the counter mode (set, count, atomic) used when building for coverage for the program that produce this meta-data file.

#### (*CoverageMetaFileReader) FileHash 

```
func (r *CoverageMetaFileReader) FileHash() [16]byte
```

FileHash returns the hash computed for all of the package meta-data blobs. Coverage counter data files refer to this hash, and the hash will be encoded into the meta-data file name.

#### (*CoverageMetaFileReader) GetPackageDecoder 

```
func (r *CoverageMetaFileReader) GetPackageDecoder(pkIdx uint32, payloadbuf []byte) (*CoverageMetaDataDecoder, []byte, error)
```

GetPackageDecoder requests a decoder object for the package within the meta-data file whose index is 'pkIdx'. If the CoverageMetaFileReader was set up with a read-only file view, a pointer into that file view will be returned, otherwise the buffer 'payloadbuf' will be written to (or if it is not of sufficient size, a new buffer will be allocated). Return value is the decoder, a byte slice with the encoded meta-data, and an error.

#### (*CoverageMetaFileReader) GetPackagePayload 

```
func (r *CoverageMetaFileReader) GetPackagePayload(pkIdx uint32, payloadbuf []byte) ([]byte, error)
```

GetPackagePayload returns the raw (encoded) meta-data payload for the package with index 'pkIdx'. As with GetPackageDecoder, if the CoverageMetaFileReader was set up with a read-only file view, a pointer into that file view will be returned, otherwise the buffer 'payloadbuf' will be written to (or if it is not of sufficient size, a new buffer will be allocated). Return value is the decoder, a byte slice with the encoded meta-data, and an error.

#### (*CoverageMetaFileReader) NumPackages 

```
func (r *CoverageMetaFileReader) NumPackages() uint64
```

NumPackages returns the number of packages for which this file contains meta-data.