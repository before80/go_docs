# cformat

https://pkg.go.dev/internal/coverage/cformat@go1.20.1






  
  
  
  

## 常量 ¶

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Formatter [¶](https://pkg.go.dev/internal/coverage/cformat@go1.20.1#Formatter)

```
type Formatter struct {
	// contains filtered or unexported fields
}
```

#### func NewFormatter [¶](https://pkg.go.dev/internal/coverage/cformat@go1.20.1#NewFormatter)

```
func NewFormatter(cm coverage.CounterMode) *Formatter
```

#### (*Formatter) AddUnit [¶](https://pkg.go.dev/internal/coverage/cformat@go1.20.1#Formatter.AddUnit)

```
func (fm *Formatter) AddUnit(file string, fname string, isfnlit bool, unit coverage.CoverableUnit, count uint32)
```

AddUnit passes info on a single coverable unit (file, funcname, literal flag, range of lines, and counter value) to the formatter. Counter values will be accumulated where appropriate.

#### (*Formatter) EmitFuncs [¶](https://pkg.go.dev/internal/coverage/cformat@go1.20.1#Formatter.EmitFuncs)

```
func (fm *Formatter) EmitFuncs(w io.Writer) error
```

EmitFuncs writes out a function-level summary to the writer 'w'. A note on handling function literals: although we collect coverage data for unnamed literals, it probably does not make sense to include them in the function summary since there isn't any good way to name them (this is also consistent with the legacy cmd/cover implementation). We do want to include their counts in the overall summary however.

#### (*Formatter) EmitPercent [¶](https://pkg.go.dev/internal/coverage/cformat@go1.20.1#Formatter.EmitPercent)

```
func (fm *Formatter) EmitPercent(w io.Writer, covpkgs string, noteEmpty bool) error
```

EmitPercent writes out a "percentage covered" string to the writer 'w'.

#### (*Formatter) EmitTextual [¶](https://pkg.go.dev/internal/coverage/cformat@go1.20.1#Formatter.EmitTextual)

```
func (fm *Formatter) EmitTextual(w io.Writer) error
```

EmitTextual writes the accumulated coverage data in the legacy cmd/cover text format to the writer 'w'. We sort the data items by importpath, source file, and line number before emitting (this sorting is not explicitly mandated by the format, but seems like a good idea for repeatable/deterministic dumps).

#### (*Formatter) SetPackage [¶](https://pkg.go.dev/internal/coverage/cformat@go1.20.1#Formatter.SetPackage)

```
func (fm *Formatter) SetPackage(importpath string)
```

SetPackage tells the formatter that we're about to visit the coverage data for the package with the specified import path. Note that it's OK to call SetPackage more than once with the same import path; counter data values will be accumulated.