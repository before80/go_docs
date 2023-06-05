# encodecounter

https://pkg.go.dev/internal/coverage/encodecounter@go1.20.1








  

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type CounterVisitor 

```
type CounterVisitor interface {
	NumFuncs() (int, error)
	VisitFuncs(f CounterVisitorFn) error
}
```

CounterVisitor describes a helper object used during counter file writing; when writing counter data files, clients pass a CounterVisitor to the write/emit routines. The writers will then first invoke the visitor's NumFuncs() method to find out how many function's worth of data to write, then it will invoke VisitFuncs. The expectation is that the VisitFuncs method will then invoke the callback "f" with data for each function to emit to the file.

### type CounterVisitorFn 

```
type CounterVisitorFn func(pkid uint32, funcid uint32, counters []uint32) error
```

CounterVisitorFn describes a callback function invoked when writing coverage counter data.

### type CoverageDataWriter 

```
type CoverageDataWriter struct {
	// contains filtered or unexported fields
}
```

#### func NewCoverageDataWriter 

```
func NewCoverageDataWriter(w io.Writer, flav coverage.CounterFlavor) *CoverageDataWriter
```

#### (*CoverageDataWriter) AppendSegment 

```
func (cfw *CoverageDataWriter) AppendSegment(args map[string]string, visitor CounterVisitor) error
```

AppendSegment appends a new segment to a counter data, with a new args section followed by a payload of counter data clauses.

#### (*CoverageDataWriter) Write 

```
func (cfw *CoverageDataWriter) Write(metaFileHash [16]byte, args map[string]string, visitor CounterVisitor) error
```

Write writes the contents of the count-data file to the writer previously supplied to NewCoverageDataWriter. Returns an error if something went wrong somewhere with the write.