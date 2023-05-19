+++
title = "plan9obj"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# plan9obj

https://pkg.go.dev/debug/plan9obj@go1.20.1





Package plan9obj implements access to Plan 9 a.out object files.

#### Security 

This package is not designed to be hardened against adversarial inputs, and is outside the scope of https://go.dev/security/policy. In particular, only basic validation is done when parsing object files. As such, care should be taken when parsing untrusted inputs, as parsing malformed files may consume significant resources, or cause panics.



## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/plan9obj/plan9obj.go;l=30)

``` go 
const (
	Magic64 = 0x8000 // 64-bit expanded header

	Magic386   = (4*11+0)*11 + 7
	MagicAMD64 = (4*26+0)*26 + 7 + Magic64
	MagicARM   = (4*20+0)*20 + 7
)
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/plan9obj/file.go;l=314)

``` go 
var ErrNoSymbols = errors.New("no symbol section")
```

ErrNoSymbols is returned by File.Symbols if there is no such section in the File.

## 函数

This section is empty.

## 类型

### type File 

``` go 
type File struct {
	FileHeader
	Sections []*Section
	// contains filtered or unexported fields
}
```

A File represents an open Plan 9 a.out file.

#### func NewFile 

``` go 
func NewFile(r io.ReaderAt) (*File, error)
```

NewFile creates a new File for accessing a Plan 9 binary in an underlying reader. The Plan 9 binary is expected to start at position 0 in the ReaderAt.

#### func Open 

``` go 
func Open(name string) (*File, error)
```

Open opens the named file using os.Open and prepares it for use as a Plan 9 a.out binary.

#### (*File) Close 

``` go 
func (f *File) Close() error
```

Close closes the File. If the File was created using NewFile directly instead of Open, Close has no effect.

#### (*File) Section 

``` go 
func (f *File) Section(name string) *Section
```

Section returns a section with the given name, or nil if no such section exists.

#### (*File) Symbols 

``` go 
func (f *File) Symbols() ([]Sym, error)
```

Symbols returns the symbol table for f.

### type FileHeader 

``` go 
type FileHeader struct {
	Magic       uint32
	Bss         uint32
	Entry       uint64
	PtrSize     int
	LoadAddress uint64
	HdrSize     uint64
}
```

A FileHeader represents a Plan 9 a.out file header.

### type Section 

``` go 
type Section struct {
	SectionHeader

	// Embed ReaderAt for ReadAt method.
	// Do not embed SectionReader directly
	// to avoid having Read and Seek.
	// If a client wants Read and Seek it must use
	// Open() to avoid fighting over the seek offset
	// with other clients.
	io.ReaderAt
	// contains filtered or unexported fields
}
```

A Section represents a single section in a Plan 9 a.out file.

#### (*Section) Data 

``` go 
func (s *Section) Data() ([]byte, error)
```

Data reads and returns the contents of the Plan 9 a.out section.

#### (*Section) Open 

``` go 
func (s *Section) Open() io.ReadSeeker
```

Open returns a new ReadSeeker reading the Plan 9 a.out section.

### type SectionHeader 

``` go 
type SectionHeader struct {
	Name   string
	Size   uint32
	Offset uint32
}
```

A SectionHeader represents a single Plan 9 a.out section header. This structure doesn't exist on-disk, but eases navigation through the object file.

### type Sym 

``` go 
type Sym struct {
	Value uint64
	Type  rune
	Name  string
}
```

A Symbol represents an entry in a Plan 9 a.out symbol table section.