+++
title = "macho"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# macho

https://pkg.go.dev/debug/macho@go1.20.1





Package macho implements access to Mach-O object files.

#### Security 

This package is not designed to be hardened against adversarial inputs, and is outside the scope of https://go.dev/security/policy. In particular, only basic validation is done when parsing object files. As such, care should be taken when parsing untrusted inputs, as parsing malformed files may consume significant resources, or cause panics.



## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/macho/macho.go;l=33)

``` go 
const (
	Magic32  uint32 = 0xfeedface
	Magic64  uint32 = 0xfeedfacf
	MagicFat uint32 = 0xcafebabe
)
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/macho/macho.go;l=203)

``` go 
const (
	FlagNoUndefs              uint32 = 0x1
	FlagIncrLink              uint32 = 0x2
	FlagDyldLink              uint32 = 0x4
	FlagBindAtLoad            uint32 = 0x8
	FlagPrebound              uint32 = 0x10
	FlagSplitSegs             uint32 = 0x20
	FlagLazyInit              uint32 = 0x40
	FlagTwoLevel              uint32 = 0x80
	FlagForceFlat             uint32 = 0x100
	FlagNoMultiDefs           uint32 = 0x200
	FlagNoFixPrebinding       uint32 = 0x400
	FlagPrebindable           uint32 = 0x800
	FlagAllModsBound          uint32 = 0x1000
	FlagSubsectionsViaSymbols uint32 = 0x2000
	FlagCanonical             uint32 = 0x4000
	FlagWeakDefines           uint32 = 0x8000
	FlagBindsToWeak           uint32 = 0x10000
	FlagAllowStackExecution   uint32 = 0x20000
	FlagRootSafe              uint32 = 0x40000
	FlagSetuidSafe            uint32 = 0x80000
	FlagNoReexportedDylibs    uint32 = 0x100000
	FlagPIE                   uint32 = 0x200000
	FlagDeadStrippableDylib   uint32 = 0x400000
	FlagHasTLVDescriptors     uint32 = 0x800000
	FlagNoHeapExecution       uint32 = 0x1000000
	FlagAppExtensionSafe      uint32 = 0x2000000
)
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/macho/fat.go;l=41)

``` go 
var ErrNotFat = &FormatError{0, "not a fat Mach-O file", nil}
```

ErrNotFat is returned from NewFatFile or OpenFat when the file is not a universal binary but may be a thin binary, based on its magic number.

## 函数

This section is empty.

## 类型

### type Cpu 

``` go 
type Cpu uint32
```

A Cpu is a Mach-O cpu type.

``` go 
const (
	Cpu386   Cpu = 7
	CpuAmd64 Cpu = Cpu386 | cpuArch64
	CpuArm   Cpu = 12
	CpuArm64 Cpu = CpuArm | cpuArch64
	CpuPpc   Cpu = 18
	CpuPpc64 Cpu = CpuPpc | cpuArch64
)
```

#### (Cpu) GoString 

``` go 
func (i Cpu) GoString() string
```

#### (Cpu) String 

``` go 
func (i Cpu) String() string
```

### type Dylib 

``` go 
type Dylib struct {
	LoadBytes
	Name           string
	Time           uint32
	CurrentVersion uint32
	CompatVersion  uint32
}
```

A Dylib represents a Mach-O load dynamic library command.

### type DylibCmd 

``` go 
type DylibCmd struct {
	Cmd            LoadCmd
	Len            uint32
	Name           uint32
	Time           uint32
	CurrentVersion uint32
	CompatVersion  uint32
}
```

A DylibCmd is a Mach-O load dynamic library command.

### type Dysymtab 

``` go 
type Dysymtab struct {
	LoadBytes
	DysymtabCmd
	IndirectSyms []uint32 // indices into Symtab.Syms
}
```

A Dysymtab represents a Mach-O dynamic symbol table command.

### type DysymtabCmd 

``` go 
type DysymtabCmd struct {
	Cmd            LoadCmd
	Len            uint32
	Ilocalsym      uint32
	Nlocalsym      uint32
	Iextdefsym     uint32
	Nextdefsym     uint32
	Iundefsym      uint32
	Nundefsym      uint32
	Tocoffset      uint32
	Ntoc           uint32
	Modtaboff      uint32
	Nmodtab        uint32
	Extrefsymoff   uint32
	Nextrefsyms    uint32
	Indirectsymoff uint32
	Nindirectsyms  uint32
	Extreloff      uint32
	Nextrel        uint32
	Locreloff      uint32
	Nlocrel        uint32
}
```

A DysymtabCmd is a Mach-O dynamic symbol table command.

### type FatArch  <- go1.3

``` go 
type FatArch struct {
	FatArchHeader
	*File
}
```

A FatArch is a Mach-O File inside a FatFile.

### type FatArchHeader  <- go1.3

``` go 
type FatArchHeader struct {
	Cpu    Cpu
	SubCpu uint32
	Offset uint32
	Size   uint32
	Align  uint32
}
```

A FatArchHeader represents a fat header for a specific image architecture.

### type FatFile  <- go1.3

``` go 
type FatFile struct {
	Magic  uint32
	Arches []FatArch
	// contains filtered or unexported fields
}
```

A FatFile is a Mach-O universal binary that contains at least one architecture.

#### func NewFatFile  <- go1.3

``` go 
func NewFatFile(r io.ReaderAt) (*FatFile, error)
```

NewFatFile creates a new FatFile for accessing all the Mach-O images in a universal binary. The Mach-O binary is expected to start at position 0 in the ReaderAt.

#### func OpenFat  <- go1.3

``` go 
func OpenFat(name string) (*FatFile, error)
```

OpenFat opens the named file using os.Open and prepares it for use as a Mach-O universal binary.

#### (*FatFile) Close  <- go1.3

``` go 
func (ff *FatFile) Close() error
```

### type File 

``` go 
type File struct {
	FileHeader
	ByteOrder binary.ByteOrder
	Loads     []Load
	Sections  []*Section

	Symtab   *Symtab
	Dysymtab *Dysymtab
	// contains filtered or unexported fields
}
```

A File represents an open Mach-O file.

#### func NewFile 

``` go 
func NewFile(r io.ReaderAt) (*File, error)
```

NewFile creates a new File for accessing a Mach-O binary in an underlying reader. The Mach-O binary is expected to start at position 0 in the ReaderAt.

#### func Open 

``` go 
func Open(name string) (*File, error)
```

Open opens the named file using os.Open and prepares it for use as a Mach-O binary.

#### (*File) Close 

``` go 
func (f *File) Close() error
```

Close closes the File. If the File was created using NewFile directly instead of Open, Close has no effect.

#### (*File) DWARF 

``` go 
func (f *File) DWARF() (*dwarf.Data, error)
```

DWARF returns the DWARF debug information for the Mach-O file.

#### (*File) ImportedLibraries 

``` go 
func (f *File) ImportedLibraries() ([]string, error)
```

ImportedLibraries returns the paths of all libraries referred to by the binary f that are expected to be linked with the binary at dynamic link time.

#### (*File) ImportedSymbols 

``` go 
func (f *File) ImportedSymbols() ([]string, error)
```

ImportedSymbols returns the names of all symbols referred to by the binary f that are expected to be satisfied by other libraries at dynamic load time.

#### (*File) Section 

``` go 
func (f *File) Section(name string) *Section
```

Section returns the first section with the given name, or nil if no such section exists.

#### (*File) Segment 

``` go 
func (f *File) Segment(name string) *Segment
```

Segment returns the first Segment with the given name, or nil if no such segment exists.

### type FileHeader 

``` go 
type FileHeader struct {
	Magic  uint32
	Cpu    Cpu
	SubCpu uint32
	Type   Type
	Ncmd   uint32
	Cmdsz  uint32
	Flags  uint32
}
```

A FileHeader represents a Mach-O file header.

### type FormatError 

``` go 
type FormatError struct {
	// contains filtered or unexported fields
}
```

FormatError is returned by some operations if the data does not have the correct format for an object file.

#### (*FormatError) Error 

``` go 
func (e *FormatError) Error() string
```

### type Load 

``` go 
type Load interface {
	Raw() []byte
}
```

A Load represents any Mach-O load command.

### type LoadBytes 

``` go 
type LoadBytes []byte
```

A LoadBytes is the uninterpreted bytes of a Mach-O load command.

#### (LoadBytes) Raw 

``` go 
func (b LoadBytes) Raw() []byte
```

### type LoadCmd 

``` go 
type LoadCmd uint32
```

A LoadCmd is a Mach-O load command.

``` go 
const (
	LoadCmdSegment    LoadCmd = 0x1
	LoadCmdSymtab     LoadCmd = 0x2
	LoadCmdThread     LoadCmd = 0x4
	LoadCmdUnixThread LoadCmd = 0x5 // thread+stack
	LoadCmdDysymtab   LoadCmd = 0xb
	LoadCmdDylib      LoadCmd = 0xc // load dylib command
	LoadCmdDylinker   LoadCmd = 0xf // id dylinker command (not load dylinker command)
	LoadCmdSegment64  LoadCmd = 0x19
	LoadCmdRpath      LoadCmd = 0x8000001c
)
```

#### (LoadCmd) GoString 

``` go 
func (i LoadCmd) GoString() string
```

#### (LoadCmd) String 

``` go 
func (i LoadCmd) String() string
```

### type Nlist32 

``` go 
type Nlist32 struct {
	Name  uint32
	Type  uint8
	Sect  uint8
	Desc  uint16
	Value uint32
}
```

An Nlist32 is a Mach-O 32-bit symbol table entry.

### type Nlist64 

``` go 
type Nlist64 struct {
	Name  uint32
	Type  uint8
	Sect  uint8
	Desc  uint16
	Value uint64
}
```

An Nlist64 is a Mach-O 64-bit symbol table entry.

### type Regs386 

``` go 
type Regs386 struct {
	AX    uint32
	BX    uint32
	CX    uint32
	DX    uint32
	DI    uint32
	SI    uint32
	BP    uint32
	SP    uint32
	SS    uint32
	FLAGS uint32
	IP    uint32
	CS    uint32
	DS    uint32
	ES    uint32
	FS    uint32
	GS    uint32
}
```

Regs386 is the Mach-O 386 register structure.

### type RegsAMD64 

``` go 
type RegsAMD64 struct {
	AX    uint64
	BX    uint64
	CX    uint64
	DX    uint64
	DI    uint64
	SI    uint64
	BP    uint64
	SP    uint64
	R8    uint64
	R9    uint64
	R10   uint64
	R11   uint64
	R12   uint64
	R13   uint64
	R14   uint64
	R15   uint64
	IP    uint64
	FLAGS uint64
	CS    uint64
	FS    uint64
	GS    uint64
}
```

RegsAMD64 is the Mach-O AMD64 register structure.

### type Reloc  <- go1.10

``` go 
type Reloc struct {
	Addr  uint32
	Value uint32
	// when Scattered == false && Extern == true, Value is the symbol number.
	// when Scattered == false && Extern == false, Value is the section number.
	// when Scattered == true, Value is the value that this reloc refers to.
	Type      uint8
	Len       uint8 // 0=byte, 1=word, 2=long, 3=quad
	Pcrel     bool
	Extern    bool // valid if Scattered == false
	Scattered bool
}
```

A Reloc represents a Mach-O relocation.

### type RelocTypeARM  <- go1.10

``` go 
type RelocTypeARM int
const (
	ARM_RELOC_VANILLA        RelocTypeARM = 0
	ARM_RELOC_PAIR           RelocTypeARM = 1
	ARM_RELOC_SECTDIFF       RelocTypeARM = 2
	ARM_RELOC_LOCAL_SECTDIFF RelocTypeARM = 3
	ARM_RELOC_PB_LA_PTR      RelocTypeARM = 4
	ARM_RELOC_BR24           RelocTypeARM = 5
	ARM_THUMB_RELOC_BR22     RelocTypeARM = 6
	ARM_THUMB_32BIT_BRANCH   RelocTypeARM = 7
	ARM_RELOC_HALF           RelocTypeARM = 8
	ARM_RELOC_HALF_SECTDIFF  RelocTypeARM = 9
)
```

#### (RelocTypeARM) GoString  <- go1.10

``` go 
func (r RelocTypeARM) GoString() string
```

#### (RelocTypeARM) String  <- go1.10

``` go 
func (i RelocTypeARM) String() string
```

### type RelocTypeARM64  <- go1.10

``` go 
type RelocTypeARM64 int
const (
	ARM64_RELOC_UNSIGNED            RelocTypeARM64 = 0
	ARM64_RELOC_SUBTRACTOR          RelocTypeARM64 = 1
	ARM64_RELOC_BRANCH26            RelocTypeARM64 = 2
	ARM64_RELOC_PAGE21              RelocTypeARM64 = 3
	ARM64_RELOC_PAGEOFF12           RelocTypeARM64 = 4
	ARM64_RELOC_GOT_LOAD_PAGE21     RelocTypeARM64 = 5
	ARM64_RELOC_GOT_LOAD_PAGEOFF12  RelocTypeARM64 = 6
	ARM64_RELOC_POINTER_TO_GOT      RelocTypeARM64 = 7
	ARM64_RELOC_TLVP_LOAD_PAGE21    RelocTypeARM64 = 8
	ARM64_RELOC_TLVP_LOAD_PAGEOFF12 RelocTypeARM64 = 9
	ARM64_RELOC_ADDEND              RelocTypeARM64 = 10
)
```

#### (RelocTypeARM64) GoString  <- go1.10

``` go 
func (r RelocTypeARM64) GoString() string
```

#### (RelocTypeARM64) String  <- go1.10

``` go 
func (i RelocTypeARM64) String() string
```

### type RelocTypeGeneric  <- go1.10

``` go 
type RelocTypeGeneric int
const (
	GENERIC_RELOC_VANILLA        RelocTypeGeneric = 0
	GENERIC_RELOC_PAIR           RelocTypeGeneric = 1
	GENERIC_RELOC_SECTDIFF       RelocTypeGeneric = 2
	GENERIC_RELOC_PB_LA_PTR      RelocTypeGeneric = 3
	GENERIC_RELOC_LOCAL_SECTDIFF RelocTypeGeneric = 4
	GENERIC_RELOC_TLV            RelocTypeGeneric = 5
)
```

#### (RelocTypeGeneric) GoString  <- go1.10

``` go 
func (r RelocTypeGeneric) GoString() string
```

#### (RelocTypeGeneric) String  <- go1.10

``` go 
func (i RelocTypeGeneric) String() string
```

### type RelocTypeX86_64  <- go1.10

``` go 
type RelocTypeX86_64 int
const (
	X86_64_RELOC_UNSIGNED   RelocTypeX86_64 = 0
	X86_64_RELOC_SIGNED     RelocTypeX86_64 = 1
	X86_64_RELOC_BRANCH     RelocTypeX86_64 = 2
	X86_64_RELOC_GOT_LOAD   RelocTypeX86_64 = 3
	X86_64_RELOC_GOT        RelocTypeX86_64 = 4
	X86_64_RELOC_SUBTRACTOR RelocTypeX86_64 = 5
	X86_64_RELOC_SIGNED_1   RelocTypeX86_64 = 6
	X86_64_RELOC_SIGNED_2   RelocTypeX86_64 = 7
	X86_64_RELOC_SIGNED_4   RelocTypeX86_64 = 8
	X86_64_RELOC_TLV        RelocTypeX86_64 = 9
)
```

#### (RelocTypeX86_64) GoString  <- go1.10

``` go 
func (r RelocTypeX86_64) GoString() string
```

#### (RelocTypeX86_64) String  <- go1.10

``` go 
func (i RelocTypeX86_64) String() string
```

### type Rpath  <- go1.10

``` go 
type Rpath struct {
	LoadBytes
	Path string
}
```

A Rpath represents a Mach-O rpath command.

### type RpathCmd  <- go1.10

``` go 
type RpathCmd struct {
	Cmd  LoadCmd
	Len  uint32
	Path uint32
}
```

A RpathCmd is a Mach-O rpath command.

### type Section 

``` go 
type Section struct {
	SectionHeader
	Relocs []Reloc

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

#### (*Section) Data 

``` go 
func (s *Section) Data() ([]byte, error)
```

Data reads and returns the contents of the Mach-O section.

#### (*Section) Open 

``` go 
func (s *Section) Open() io.ReadSeeker
```

Open returns a new ReadSeeker reading the Mach-O section.

### type Section32 

``` go 
type Section32 struct {
	Name     [16]byte
	Seg      [16]byte
	Addr     uint32
	Size     uint32
	Offset   uint32
	Align    uint32
	Reloff   uint32
	Nreloc   uint32
	Flags    uint32
	Reserve1 uint32
	Reserve2 uint32
}
```

A Section32 is a 32-bit Mach-O section header.

### type Section64 

``` go 
type Section64 struct {
	Name     [16]byte
	Seg      [16]byte
	Addr     uint64
	Size     uint64
	Offset   uint32
	Align    uint32
	Reloff   uint32
	Nreloc   uint32
	Flags    uint32
	Reserve1 uint32
	Reserve2 uint32
	Reserve3 uint32
}
```

A Section64 is a 64-bit Mach-O section header.

### type SectionHeader 

``` go 
type SectionHeader struct {
	Name   string
	Seg    string
	Addr   uint64
	Size   uint64
	Offset uint32
	Align  uint32
	Reloff uint32
	Nreloc uint32
	Flags  uint32
}
```

### type Segment 

``` go 
type Segment struct {
	LoadBytes
	SegmentHeader

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

A Segment represents a Mach-O 32-bit or 64-bit load segment command.

#### (*Segment) Data 

``` go 
func (s *Segment) Data() ([]byte, error)
```

Data reads and returns the contents of the segment.

#### (*Segment) Open 

``` go 
func (s *Segment) Open() io.ReadSeeker
```

Open returns a new ReadSeeker reading the segment.

### type Segment32 

``` go 
type Segment32 struct {
	Cmd     LoadCmd
	Len     uint32
	Name    [16]byte
	Addr    uint32
	Memsz   uint32
	Offset  uint32
	Filesz  uint32
	Maxprot uint32
	Prot    uint32
	Nsect   uint32
	Flag    uint32
}
```

A Segment32 is a 32-bit Mach-O segment load command.

### type Segment64 

``` go 
type Segment64 struct {
	Cmd     LoadCmd
	Len     uint32
	Name    [16]byte
	Addr    uint64
	Memsz   uint64
	Offset  uint64
	Filesz  uint64
	Maxprot uint32
	Prot    uint32
	Nsect   uint32
	Flag    uint32
}
```

A Segment64 is a 64-bit Mach-O segment load command.

### type SegmentHeader 

``` go 
type SegmentHeader struct {
	Cmd     LoadCmd
	Len     uint32
	Name    string
	Addr    uint64
	Memsz   uint64
	Offset  uint64
	Filesz  uint64
	Maxprot uint32
	Prot    uint32
	Nsect   uint32
	Flag    uint32
}
```

A SegmentHeader is the header for a Mach-O 32-bit or 64-bit load segment command.

### type Symbol 

``` go 
type Symbol struct {
	Name  string
	Type  uint8
	Sect  uint8
	Desc  uint16
	Value uint64
}
```

A Symbol is a Mach-O 32-bit or 64-bit symbol table entry.

### type Symtab 

``` go 
type Symtab struct {
	LoadBytes
	SymtabCmd
	Syms []Symbol
}
```

A Symtab represents a Mach-O symbol table command.

### type SymtabCmd 

``` go 
type SymtabCmd struct {
	Cmd     LoadCmd
	Len     uint32
	Symoff  uint32
	Nsyms   uint32
	Stroff  uint32
	Strsize uint32
}
```

A SymtabCmd is a Mach-O symbol table command.

### type Thread 

``` go 
type Thread struct {
	Cmd  LoadCmd
	Len  uint32
	Type uint32
	Data []uint32
}
```

A Thread is a Mach-O thread state command.

### type Type 

``` go 
type Type uint32
```

A Type is the Mach-O file type, e.g. an object file, executable, or dynamic library.

``` go 
const (
	TypeObj    Type = 1
	TypeExec   Type = 2
	TypeDylib  Type = 6
	TypeBundle Type = 8
)
```

#### (Type) GoString  <- go1.10

``` go 
func (t Type) GoString() string
```

#### (Type) String  <- go1.10

``` go 
func (t Type) String() string
```