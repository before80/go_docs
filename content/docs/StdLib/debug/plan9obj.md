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

Package plan9obj 实现对 Plan 9 a.out 对象文件的访问。

#### Security 

This package is not designed to be hardened against adversarial inputs, and is outside the scope of https://go.dev/security/policy. In particular, only basic validation is done when parsing object files. As such, care should be taken when parsing untrusted inputs, as parsing malformed files may consume significant resources, or cause panics.

该包没有设计用于抵御对抗性输入，并且超出了 https://go.dev/security/policy 的范围。特别是，在解析对象文件时只进行了基本验证。因此，在解析不受信任的输入时应该小心，因为解析格式错误的文件可能会消耗大量资源或导致崩溃。



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

如果文件中没有该节（section），则 File.Symbols 返回 ErrNoSymbols。

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

File 表示一个打开的 Plan 9 a.out 文件。

#### func NewFile 

``` go 
func NewFile(r io.ReaderAt) (*File, error)
```

NewFile creates a new File for accessing a Plan 9 binary in an underlying reader. The Plan 9 binary is expected to start at position 0 in the ReaderAt.

NewFile 创建一个新的 File，用于访问底层 reader 中的 Plan 9 二进制文件。Plan 9 二进制文件预期从 ReaderAt 的位置 0 开始。

#### func Open 

``` go 
func Open(name string) (*File, error)
```

Open opens the named file using os.Open and prepares it for use as a Plan 9 a.out binary.

Open 使用 os.Open 打开指定的文件，并准备将其用作 Plan 9 a.out 二进制文件。

#### (*File) Close 

``` go 
func (f *File) Close() error
```

Close closes the File. If the File was created using NewFile directly instead of Open, Close has no effect.

Close 关闭文件。如果 File 是通过 NewFile 直接创建而不是通过 Open 创建的，则 Close 不起作用。

#### (*File) Section 

``` go 
func (f *File) Section(name string) *Section
```

Section returns a section with the given name, or nil if no such section exists.

Section 返回具有给定名称的节（section），如果不存在则返回 nil。

#### (*File) Symbols 

``` go 
func (f *File) Symbols() ([]Sym, error)
```

Symbols returns the symbol table for f.

Symbols 返回 f 的符号表。

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

FileHeader 表示 Plan 9 a.out 文件头。

### type Section 

``` go 
type Section struct {
	SectionHeader

    // 通过嵌入 ReaderAt 实现 ReadAt 方法。
    // 不直接嵌入 SectionReader，以避免造成 Read 和 Seek 的冲突。
    // 如果客户端需要 Read 和 Seek，必须使用 Open() 方法以避免与其他客户端争用 Seek 偏移量。	
	io.ReaderAt
	// contains filtered or unexported fields
}
```

A Section represents a single section in a Plan 9 a.out file.

Section 表示 Plan 9 a.out 文件中的单个节（section）。

#### (*Section) Data 

``` go 
func (s *Section) Data() ([]byte, error)
```

Data reads and returns the contents of the Plan 9 a.out section.

Data 读取并返回 Plan 9 a.out 节的内容。

#### (*Section) Open 

``` go 
func (s *Section) Open() io.ReadSeeker
```

Open returns a new ReadSeeker reading the Plan 9 a.out section.

Open 返回一个新的 ReadSeeker，用于读取 Plan 9 a.out 节。

### type SectionHeader 

``` go 
type SectionHeader struct {
	Name   string
	Size   uint32
	Offset uint32
}
```

A SectionHeader represents a single Plan 9 a.out section header. This structure doesn't exist on-disk, but eases navigation through the object file.

SectionHeader 表示 Plan 9 a.out 单个节的头。该结构在磁盘上不存在，但可以方便地导航对象文件。

### type Sym 

``` go 
type Sym struct {
	Value uint64
	Type  rune
	Name  string
}
```

A Symbol represents an entry in a Plan 9 a.out symbol table section.

Sym 表示 Plan 9 a.out 符号表节中的条目。