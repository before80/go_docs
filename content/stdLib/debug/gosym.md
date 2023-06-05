+++
title = "gosym"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# gosym

https://pkg.go.dev/debug/gosym@go1.20.1

Package gosym implements access to the Go symbol and line number tables embedded in Go binaries generated by the gc compilers.

包 gosym 实现了对由 gc 编译器生成的 Go 二进制文件中嵌入的符号和行号表的访问。

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

DecodingError 表示在解码符号表过程中发生的错误。

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
	Params    []*Sym // nil for Go 1.3 and later binaries   Go 1.3 及以后的二进制文件为 nil
	Locals    []*Sym // nil for Go 1.3 and later binaries   Go 1.3 及以后的二进制文件为 nil
	FrameSize int
	LineTable *LineTable
	Obj       *Obj
}
```

A Func collects information about a single function.

Func 收集关于单个函数的信息。

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

LineTable 是一个将程序计数器映射到行号的数据结构。

In Go 1.1 and earlier, each function (represented by a Func) had its own LineTable, and the line number corresponded to a numbering of all source lines in the program, across all files. That absolute line number would then have to be converted separately to a file name and line number within the file.

在 Go 1.1 及之前的版本中，每个函数（由 Func 表示）都有自己的 LineTable，而行号对应于程序中所有文件的所有源代码行的编号。然后，该绝对行号必须单独转换为文件名和文件内的行号。

In Go 1.2, the format of the data changed so that there is a single LineTable for the entire program, shared by all Funcs, and there are no absolute line numbers, just line numbers within specific files.

在 Go 1.2 中，数据的格式发生了变化，整个程序共享一个 LineTable，由所有的 Func 共享，并且没有绝对行号，只有特定文件内的行号。

For the most part, LineTable's methods should be treated as an internal detail of the package; callers should use the methods on Table instead.

在大多数情况下，LineTable 的方法应被视为包的内部细节；调用者应使用 Table 上的方法。

#### func NewLineTable 

``` go 
func NewLineTable(data []byte, text uint64) *LineTable
```

NewLineTable returns a new PC/line table corresponding to the encoded data. Text must be the start address of the corresponding text segment.

NewLineTable 返回对应于编码数据的新的 PC/行号表。Text 必须是相应文本段的起始地址。

### type Obj 

``` go 
type Obj struct {
	// Funcs is a list of functions in the Obj.
    // Funcs 是 Obj 中的函数列表。
	Funcs []Func

	// In Go 1.1 and earlier, Paths is a list of symbols corresponding
	// to the source file names that produced the Obj.
	// In Go 1.2, Paths is nil.
	// Use the keys of Table.Files to obtain a list of source files.
    // 在 Go 1.1 及之前的版本中，Paths 是与生成 Obj 的源文件名相对应的符号列表。
	// 在 Go 1.2 中，Paths 为 nil。
	// 使用 Table.Files 的键获取源文件的列表。
	Paths []Sym // meta
}
```

An Obj represents a collection of functions in a symbol table.

Obj 表示符号表中的一组函数。

The exact method of division of a binary into separate Objs is an internal detail of the symbol table format.

将二进制文件划分为单独的 Obj 的确切方法是符号表格式的内部细节。

In early versions of Go each source file became a different Obj.

在早期的 Go 版本中，每个源文件成为不同的 Obj。

In Go 1 and Go 1.1, each package produced one Obj for all Go sources and one Obj per C source file.

在 Go 1 和 Go 1.1 中，每个包为所有的 Go 源文件产生一个 Obj，并为每个 C 源文件产生一个 Obj。

In Go 1.2, there is a single Obj for the entire program.

在 Go 1.2 中，整个程序只有一个 Obj。

### type Sym 

``` go 
type Sym struct {
	Value  uint64
	Type   byte
	Name   string
	GoType uint64
	// If this symbol is a function symbol, the corresponding Func
    // 如果这个符号是一个函数符号，那么对应的是 Func。
	Func *Func
	// contains filtered or unexported fields
}
```

A Sym represents a single symbol table entry.

Sym 表示单个符号表条目。

#### (*Sym) BaseName 

``` go 
func (s *Sym) BaseName() string
```

BaseName returns the symbol name without the package or receiver name.

BaseName 返回不包含包名或接收器名称的符号名。

#### (*Sym) PackageName 

``` go 
func (s *Sym) PackageName() string
```

PackageName returns the package part of the symbol name, or the empty string if there is none.

PackageName 返回符号名的包部分，如果没有则返回空字符串。

#### (*Sym) ReceiverName 

``` go 
func (s *Sym) ReceiverName() string
```

ReceiverName returns the receiver type name of this symbol, or the empty string if there is none. A receiver name is only detected in the case that s.Name is fully-specified with a package name.

ReceiverName 返回此符号的接收器类型名称，如果没有则返回空字符串。仅在 s.Name 完全指定了包名的情况下才检测到接收器名称。

#### (*Sym) Static 

``` go 
func (s *Sym) Static() bool
```

Static reports whether this symbol is static (not visible outside its file).

Static 报告此符号是否为静态符号（在其文件外不可见）。

### type Table 

``` go 
type Table struct {
	Syms  []Sym // nil for Go 1.3 and later binaries   Go 1.3及以后的版本二进制文件为nil
	Funcs []Func
	Files map[string]*Obj // for Go 1.2 and later all files map to one Obj 对于 Go 1.2 及以后的版本，所有文件都映射到一个 Obj
	Objs  []Obj           // for Go 1.2 and later only one Obj in slice 对于 Go 1.2 及以后的版本，只有一个 Obj 在切片中
	// contains filtered or unexported fields
}
```

Table represents a Go symbol table. It stores all of the symbols decoded from the program and provides methods to translate between symbols, names, and addresses.

Table 表示 Go 符号表。它存储从程序中解码的所有符号，并提供在符号、名称和地址之间进行转换的方法。

#### func NewTable 

``` go 
func NewTable(symtab []byte, pcln *LineTable) (*Table, error)
```

NewTable decodes the Go symbol table (the ".gosymtab" section in ELF), returning an in-memory representation. Starting with Go 1.3, the Go symbol table no longer includes symbol data.

NewTable 解码 Go 符号表（ELF 中的 ".gosymtab" 部分），返回一个内存中的表示。从 Go 1.3 开始，Go 符号表不再包含符号数据。

#### (*Table) LineToPC 

``` go 
func (t *Table) LineToPC(file string, line int) (pc uint64, fn *Func, err error)
```

LineToPC looks up the first program counter on the given line in the named file. It returns UnknownPathError or UnknownLineError if there is an error looking up this line.

LineToPC 在指定文件中查找给定行号上的第一个程序计数器。如果查找该行时出现错误，将返回 UnknownPathError 或 UnknownLineError。

#### (*Table) LookupFunc 

``` go 
func (t *Table) LookupFunc(name string) *Func
```

LookupFunc returns the text, data, or bss symbol with the given name, or nil if no such symbol is found.

LookupFunc 返回具有给定名称的文本、数据或 bss 符号，如果找不到此类符号，则返回 nil。

#### (*Table) LookupSym 

``` go 
func (t *Table) LookupSym(name string) *Sym
```

LookupSym returns the text, data, or bss symbol with the given name, or nil if no such symbol is found.

LookupSym 返回具有给定名称的文本、数据或 bss 符号，如果找不到此类符号，则返回 nil。

#### (*Table) PCToFunc 

``` go 
func (t *Table) PCToFunc(pc uint64) *Func
```

PCToFunc returns the function containing the program counter pc, or nil if there is no such function.

PCToFunc 返回包含程序计数器 pc 的函数，如果没有这样的函数，则返回 nil。

#### (*Table) PCToLine 

``` go 
func (t *Table) PCToLine(pc uint64) (file string, line int, fn *Func)
```

PCToLine looks up line number information for a program counter. If there is no information, it returns fn == nil.

PCToLine 查找程序计数器的行号信息。如果没有信息，则返回 fn == nil。

#### (*Table) SymByAddr 

``` go 
func (t *Table) SymByAddr(addr uint64) *Sym
```

SymByAddr returns the text, data, or bss symbol starting at the given address.

SymByAddr 返回以给定地址开始的文本、数据或 bss 符号。

### type UnknownFileError 

``` go 
type UnknownFileError string
```

UnknownFileError represents a failure to find the specific file in the symbol table.

UnknownLineError 表示无法将行号映射到程序计数器，原因是行号超出文件范围或给定行上没有代码。

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

UnknownLineError 表示将行号映射到程序计数器失败，可能是因为行号超出文件范围或给定行上没有代码。

#### (*UnknownLineError) Error 

``` go 
func (e *UnknownLineError) Error() string
```