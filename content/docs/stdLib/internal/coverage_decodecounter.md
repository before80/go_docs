# decodecounter

https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1






  
  
  
  
  
  


## 常量 [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#pkg-constants)

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type [CounterDataReader](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodecounter/decodecounterfile.go;l=22) [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#CounterDataReader)

```
type CounterDataReader struct {
	// contains filtered or unexported fields
}
```

#### func [NewCounterDataReader](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodecounter/decodecounterfile.go;l=39) [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#NewCounterDataReader)

```
func NewCounterDataReader(fn string, rs io.ReadSeeker) (*CounterDataReader, error)
```

#### (*CounterDataReader) [BeginNextSegment](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodecounter/decodecounterfile.go;l=243) [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#CounterDataReader.BeginNextSegment)

```
func (cdr *CounterDataReader) BeginNextSegment() (bool, error)
```

BeginNextSegment sets up the the reader to read the next segment, returning TRUE if we do have another segment to read, or FALSE if we're done with all the segments (also an error if something went wrong).

#### (*CounterDataReader) [Goarch](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodecounter/decodecounterfile.go;l=222) [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#CounterDataReader.Goarch)

```
func (cdr *CounterDataReader) Goarch() string
```

Goarch returns the GOARCH setting in effect for the "-cover" binary that produced this counter data file. The GOARCH value may be empty in the case where the counter data file was produced from a merge in which more than one GOARCH value was present.

#### (*CounterDataReader) [Goos](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodecounter/decodecounterfile.go;l=214) [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#CounterDataReader.Goos)

```
func (cdr *CounterDataReader) Goos() string
```

Goos returns the GOOS setting in effect for the "-cover" binary that produced this counter data file. The GOOS value may be empty in the case where the counter data file was produced from a merge in which more than one GOOS value was present.

#### (*CounterDataReader) [NextFunc](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodecounter/decodecounterfile.go;l=274) [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#CounterDataReader.NextFunc)

```
func (cdr *CounterDataReader) NextFunc(p *FuncPayload) (bool, error)
```

NextFunc reads data for the next function in this current segment into "p", returning TRUE if the read was successful or FALSE if we've read all the functions already (also an error if something went wrong with the read or we hit a premature EOF).

#### (*CounterDataReader) [NumFunctionsInSegment](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodecounter/decodecounterfile.go;l=263) [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#CounterDataReader.NumFunctionsInSegment)

```
func (cdr *CounterDataReader) NumFunctionsInSegment() uint32
```

NumFunctionsInSegment returns the number of live functions in the currently selected segment.

#### (*CounterDataReader) [NumSegments](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodecounter/decodecounterfile.go;l=235) [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#CounterDataReader.NumSegments)

```
func (cdr *CounterDataReader) NumSegments() uint32
```

NumSegments returns the number of execution segments in the file.

#### (*CounterDataReader) [OsArgs](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodecounter/decodecounterfile.go;l=206) [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#CounterDataReader.OsArgs)

```
func (cdr *CounterDataReader) OsArgs() []string
```

OsArgs returns the program arguments (saved from os.Args during the run of the instrumented binary) read from the counter data file. Not all coverage data files will have os.Args values; for example, if a data file is produced by merging coverage data from two distinct runs, no os args will be available (an empty list is returned).

### type [FuncPayload](https://cs.opensource.google/go/go/+/go1.20.1:src/internal/coverage/decodecounter/decodecounterfile.go;l=228) [¶](https://pkg.go.dev/internal/coverage/decodecounter@go1.20.1#FuncPayload)

```
type FuncPayload struct {
	PkgIdx   uint32
	FuncIdx  uint32
	Counters []uint32
}
```

FuncPayload encapsulates the counter data payload for a single function as read from a counter data file.