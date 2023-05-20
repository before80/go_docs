+++
title = "gosym"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# gosym

https://pkg.go.dev/debug/gosym@go1.20.1

Package gosym implements access to the Go symbol and line number tables embedded in Go binaries generated by the gc compilers.



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type DecodingError 

``` go 
type DecodingError struct {
	// contains filtered or unexported fields
}
```

DecodingError represents an error during the decoding of the symbol table.

#### (*DecodingError) Error 

``` go 
func (e *DecodingError) Error() string
```

### type Func 

``` go 
type Func struct {
	Entry uint64
	*Sym
	End       uint64
	Params    []*Sym // nil for Go 1.3 and later binaries
	Locals    []*Sym // nil for Go 1.3 and later binaries
	FrameSize int
	LineTable *LineTable
	Obj       *Obj
}
```

A Func collects information about a single function.

### type LineTable 

``` go 
type LineTable struct {
	Data []byte
	PC   uint64
	Line int
	// contains filtered or unexported fields
}
```

A LineTable is a data structure mapping program counters to line numbers.

In Go 1.1 and earlier, each function (represented by a Func) had its own LineTable, and the line number corresponded to a numbering of all source lines in the program, across all files. That absolute line number would then have to be converted separately to a file name and line number within the file.

In Go 1.2, the format of the data changed so that there is a single LineTable for the entire program, shared by all Funcs, and there are no absolute line numbers, just line numbers within specific files.

For the most part, LineTable's methods should be treated as an internal detail of the package; callers should use the methods on Table instead.

#### func NewLineTable 

``` go 
func NewLineTable(data []byte, text uint64) *LineTable
```

NewLineTable returns a new PC/line table corresponding to the encoded data. Text must be the start address of the corresponding text segment.

##### Example
``` go 
```

##### Example
``` go 
```

### type Obj 

``` go 
type Obj struct {
	// Funcs is a list of functions in the Obj.
	Funcs []Func

	// In Go 1.1 and earlier, Paths is a list of symbols corresponding
	// to the source file names that produced the Obj.
	// In Go 1.2, Paths is nil.
	// Use the keys of Table.Files to obtain a list of source files.
	Paths []Sym // meta
}
```

An Obj represents a collection of functions in a symbol table.

The exact method of division of a binary into separate Objs is an internal detail of the symbol table format.

In early versions of Go each source file became a different Obj.

In Go 1 and Go 1.1, each package produced one Obj for all Go sources and one Obj per C source file.

In Go 1.2, there is a single Obj for the entire program.

### type Sym 

``` go 
type Sym struct {
	Value  uint64
	Type   byte
	Name   string
	GoType uint64
	// If this symbol is a function symbol, the corresponding Func
	Func *Func
	// contains filtered or unexported fields
}
```

A Sym represents a single symbol table entry.

#### (*Sym) BaseName 

``` go 
func (s *Sym) BaseName() string
```

BaseName returns the symbol name without the package or receiver name.

#### (*Sym) PackageName 

``` go 
func (s *Sym) PackageName() string
```

PackageName returns the package part of the symbol name, or the empty string if there is none.

#### (*Sym) ReceiverName 

``` go 
func (s *Sym) ReceiverName() string
```

ReceiverName returns the receiver type name of this symbol, or the empty string if there is none. A receiver name is only detected in the case that s.Name is fully-specified with a package name.

#### (*Sym) Static 

``` go 
func (s *Sym) Static() bool
```

Static reports whether this symbol is static (not visible outside its file).

### type Table 

``` go 
type Table struct {
	Syms  []Sym // nil for Go 1.3 and later binaries
	Funcs []Func
	Files map[string]*Obj // for Go 1.2 and later all files map to one Obj
	Objs  []Obj           // for Go 1.2 and later only one Obj in slice
	// contains filtered or unexported fields
}
```

Table represents a Go symbol table. It stores all of the symbols decoded from the program and provides methods to translate between symbols, names, and addresses.

#### func NewTable 

``` go 
func NewTable(symtab []byte, pcln *LineTable) (*Table, error)
```

NewTable decodes the Go symbol table (the ".gosymtab" section in ELF), returning an in-memory representation. Starting with Go 1.3, the Go symbol table no longer includes symbol data.

#### (*Table) LineToPC 

``` go 
func (t *Table) LineToPC(file string, line int) (pc uint64, fn *Func, err error)
```

LineToPC looks up the first program counter on the given line in the named file. It returns UnknownPathError or UnknownLineError if there is an error looking up this line.

#### (*Table) LookupFunc 

``` go 
func (t *Table) LookupFunc(name string) *Func
```

LookupFunc returns the text, data, or bss symbol with the given name, or nil if no such symbol is found.

#### (*Table) LookupSym 

``` go 
func (t *Table) LookupSym(name string) *Sym
```

LookupSym returns the text, data, or bss symbol with the given name, or nil if no such symbol is found.

#### (*Table) PCToFunc 

``` go 
func (t *Table) PCToFunc(pc uint64) *Func
```

PCToFunc returns the function containing the program counter pc, or nil if there is no such function.

#### (*Table) PCToLine 

``` go 
func (t *Table) PCToLine(pc uint64) (file string, line int, fn *Func)
```

PCToLine looks up line number information for a program counter. If there is no information, it returns fn == nil.

#### (*Table) SymByAddr 

``` go 
func (t *Table) SymByAddr(addr uint64) *Sym
```

SymByAddr returns the text, data, or bss symbol starting at the given address.

### type UnknownFileError 

``` go 
type UnknownFileError string
```

UnknownFileError represents a failure to find the specific file in the symbol table.

#### (UnknownFileError) Error 

``` go 
func (e UnknownFileError) Error() string
```

### type UnknownLineError 

``` go 
type UnknownLineError struct {
	File string
	Line int
}
```

UnknownLineError represents a failure to map a line to a program counter, either because the line is beyond the bounds of the file or because there is no code on the given line.

#### (*UnknownLineError) Error 

``` go 
func (e *UnknownLineError) Error() string
```