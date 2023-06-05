# decodecounter

https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1






  
  
  
  
  
  


## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type CounterDataReader 

``` go
type CounterDataReader struct {
	// contains filtered or unexported fields
}
```

#### func NewCounterDataReader 

``` go
func NewCounterDataReader(fn string, rs io.ReadSeeker) (*CounterDataReader, error)
```

#### (*CounterDataReader) BeginNextSegment 

``` go
func (cdr *CounterDataReader) BeginNextSegment() (bool, error)
```

BeginNextSegment sets up the the reader to read the next segment, returning TRUE if we do have another segment to read, or FALSE if we're done with all the segments (also an error if something went wrong).

#### (*CounterDataReader) Goarch 

``` go
func (cdr *CounterDataReader) Goarch() string
```

Goarch returns the GOARCH setting in effect for the "-cover" binary that produced this counter data file. The GOARCH value may be empty in the case where the counter data file was produced from a merge in which more than one GOARCH value was present.

#### (*CounterDataReader) Goos 

``` go
func (cdr *CounterDataReader) Goos() string
```

Goos returns the GOOS setting in effect for the "-cover" binary that produced this counter data file. The GOOS value may be empty in the case where the counter data file was produced from a merge in which more than one GOOS value was present.

#### (*CounterDataReader) NextFunc 

``` go
func (cdr *CounterDataReader) NextFunc(p *FuncPayload) (bool, error)
```

NextFunc reads data for the next function in this current segment into "p", returning TRUE if the read was successful or FALSE if we've read all the functions already (also an error if something went wrong with the read or we hit a premature EOF).

#### (*CounterDataReader) NumFunctionsInSegment 

``` go
func (cdr *CounterDataReader) NumFunctionsInSegment() uint32
```

NumFunctionsInSegment returns the number of live functions in the currently selected segment.

#### (*CounterDataReader) NumSegments 

``` go
func (cdr *CounterDataReader) NumSegments() uint32
```

NumSegments returns the number of execution segments in the file.

#### (*CounterDataReader) OsArgs 

``` go
func (cdr *CounterDataReader) OsArgs() []string
```

OsArgs returns the program arguments (saved from os.Args during the run of the instrumented binary) read from the counter data file. Not all coverage data files will have os.Args values; for example, if a data file is produced by merging coverage data from two distinct runs, no os args will be available (an empty list is returned).

### type FuncPayload 

``` go
type FuncPayload struct {
	PkgIdx   uint32
	FuncIdx  uint32
	Counters []uint32
}
```

FuncPayload encapsulates the counter data payload for a single function as read from a counter data file.