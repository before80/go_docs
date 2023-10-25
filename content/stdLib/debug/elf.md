+++
title = "elf"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/debug/elf@go1.20.1





Package elf implements access to ELF object files.

#### Security 

This package is not designed to be hardened against adversarial inputs, and is outside the scope of https://go.dev/security/policy. In particular, only basic validation is done when parsing object files. As such, care should be taken when parsing untrusted inputs, as parsing malformed files may consume significant resources, or cause panics.



## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/elf/elf.go;l=54)

``` go 
const (
	EI_CLASS      = 4  /* Class of machine. */
	EI_DATA       = 5  /* Data format. */
	EI_VERSION    = 6  /* ELF format version. */
	EI_OSABI      = 7  /* Operating system / ABI identification */
	EI_ABIVERSION = 8  /* ABI version */
	EI_PAD        = 9  /* Start of padding (per SVR4 ABI). */
	EI_NIDENT     = 16 /* Size of e_ident array. */
)
```

Indexes into the Header.Ident array.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/elf/elf.go;l=3232)

``` go 
const ARM_MAGIC_TRAMP_NUMBER = 0x5c000003
```

Magic number for the elf trampoline, chosen wisely to be an immediate value.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/elf/elf.go;l=65)

``` go 
const ELFMAG = "\177ELF"
```

Initial magic number for ELF files.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/elf/elf.go;l=3322)

``` go 
const Sym32Size = 16
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/elf/elf.go;l=3424)

``` go 
const Sym64Size = 24
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/elf/file.go;l=599)

``` go 
var ErrNoSymbols = errors.New("no symbol section")
```

ErrNoSymbols is returned by File.Symbols and File.DynamicSymbols if there is no such section in the File.

## 函数

#### func R_INFO 

``` go 
func R_INFO(sym, typ uint32) uint64
```

#### func R_INFO32 

``` go 
func R_INFO32(sym, typ uint32) uint32
```

#### func R_SYM32 

``` go 
func R_SYM32(info uint32) uint32
```

#### func R_SYM64 

``` go 
func R_SYM64(info uint64) uint32
```

#### func R_TYPE32 

``` go 
func R_TYPE32(info uint32) uint32
```

#### func R_TYPE64 

``` go 
func R_TYPE64(info uint64) uint32
```

#### func ST_INFO 

``` go 
func ST_INFO(bind SymBind, typ SymType) uint8
```

## 类型

### type Chdr32  <- go1.6

``` go 
type Chdr32 struct {
	Type      uint32
	Size      uint32
	Addralign uint32
}
```

ELF32 Compression header.

### type Chdr64  <- go1.6

``` go 
type Chdr64 struct {
	Type uint32

	Size      uint64
	Addralign uint64
	// contains filtered or unexported fields
}
```

ELF64 Compression header.

### type Class 

``` go 
type Class byte
```

Class is found in Header.Ident[EI_CLASS] and Header.Class.

``` go 
const (
	ELFCLASSNONE Class = 0 /* Unknown class. */
	ELFCLASS32   Class = 1 /* 32-bit architecture. */
	ELFCLASS64   Class = 2 /* 64-bit architecture. */
)
```

#### (Class) GoString 

``` go 
func (i Class) GoString() string
```

#### (Class) String 

``` go 
func (i Class) String() string
```

### type CompressionType  <- go1.6

``` go 
type CompressionType int
```

Section compression type.

``` go 
const (
	COMPRESS_ZLIB   CompressionType = 1          /* ZLIB compression. */
	COMPRESS_LOOS   CompressionType = 0x60000000 /* First OS-specific. */
	COMPRESS_HIOS   CompressionType = 0x6fffffff /* Last OS-specific. */
	COMPRESS_LOPROC CompressionType = 0x70000000 /* First processor-specific type. */
	COMPRESS_HIPROC CompressionType = 0x7fffffff /* Last processor-specific type. */
)
```

#### (CompressionType) GoString  <- go1.6

``` go 
func (i CompressionType) GoString() string
```

#### (CompressionType) String  <- go1.6

``` go 
func (i CompressionType) String() string
```

### type Data 

``` go 
type Data byte
```

Data is found in Header.Ident[EI_DATA] and Header.Data.

``` go 
const (
	ELFDATANONE Data = 0 /* Unknown data format. */
	ELFDATA2LSB Data = 1 /* 2's complement little-endian. */
	ELFDATA2MSB Data = 2 /* 2's complement big-endian. */
)
```

#### (Data) GoString 

``` go 
func (i Data) GoString() string
```

#### (Data) String 

``` go 
func (i Data) String() string
```

### type Dyn32 

``` go 
type Dyn32 struct {
	Tag int32  /* Entry type. */
	Val uint32 /* Integer/Address value. */
}
```

ELF32 Dynamic structure. The ".dynamic" section contains an array of them.

### type Dyn64 

``` go 
type Dyn64 struct {
	Tag int64  /* Entry type. */
	Val uint64 /* Integer/address value */
}
```

ELF64 Dynamic structure. The ".dynamic" section contains an array of them.

### type DynFlag 

``` go 
type DynFlag int
```

DT_FLAGS values.

``` go 
const (
	DF_ORIGIN DynFlag = 0x0001 /* Indicates that the object being loaded may
	   make reference to the
	   $ORIGIN substitution string */
	DF_SYMBOLIC DynFlag = 0x0002 /* Indicates "symbolic" linking. */
	DF_TEXTREL  DynFlag = 0x0004 /* Indicates there may be relocations in non-writable segments. */
	DF_BIND_NOW DynFlag = 0x0008 /* Indicates that the dynamic linker should
	   process all relocations for the object
	   containing this entry before transferring
	   control to the program. */
	DF_STATIC_TLS DynFlag = 0x0010 /* Indicates that the shared object or
	   executable contains code using a static
	   thread-local storage scheme. */
)
```

#### (DynFlag) GoString 

``` go 
func (i DynFlag) GoString() string
```

#### (DynFlag) String 

``` go 
func (i DynFlag) String() string
```

### type DynTag 

``` go 
type DynTag int
```

Dyn.Tag

``` go 
const (
	DT_NULL         DynTag = 0  /* Terminating entry. */
	DT_NEEDED       DynTag = 1  /* String table offset of a needed shared library. */
	DT_PLTRELSZ     DynTag = 2  /* Total size in bytes of PLT relocations. */
	DT_PLTGOT       DynTag = 3  /* Processor-dependent address. */
	DT_HASH         DynTag = 4  /* Address of symbol hash table. */
	DT_STRTAB       DynTag = 5  /* Address of string table. */
	DT_SYMTAB       DynTag = 6  /* Address of symbol table. */
	DT_RELA         DynTag = 7  /* Address of ElfNN_Rela relocations. */
	DT_RELASZ       DynTag = 8  /* Total size of ElfNN_Rela relocations. */
	DT_RELAENT      DynTag = 9  /* Size of each ElfNN_Rela relocation entry. */
	DT_STRSZ        DynTag = 10 /* Size of string table. */
	DT_SYMENT       DynTag = 11 /* Size of each symbol table entry. */
	DT_INIT         DynTag = 12 /* Address of initialization function. */
	DT_FINI         DynTag = 13 /* Address of finalization function. */
	DT_SONAME       DynTag = 14 /* String table offset of shared object name. */
	DT_RPATH        DynTag = 15 /* String table offset of library path. [sup] */
	DT_SYMBOLIC     DynTag = 16 /* Indicates "symbolic" linking. [sup] */
	DT_REL          DynTag = 17 /* Address of ElfNN_Rel relocations. */
	DT_RELSZ        DynTag = 18 /* Total size of ElfNN_Rel relocations. */
	DT_RELENT       DynTag = 19 /* Size of each ElfNN_Rel relocation. */
	DT_PLTREL       DynTag = 20 /* Type of relocation used for PLT. */
	DT_DEBUG        DynTag = 21 /* Reserved (not used). */
	DT_TEXTREL      DynTag = 22 /* Indicates there may be relocations in non-writable segments. [sup] */
	DT_JMPREL       DynTag = 23 /* Address of PLT relocations. */
	DT_BIND_NOW     DynTag = 24 /* [sup] */
	DT_INIT_ARRAY   DynTag = 25 /* Address of the array of pointers to initialization functions */
	DT_FINI_ARRAY   DynTag = 26 /* Address of the array of pointers to termination functions */
	DT_INIT_ARRAYSZ DynTag = 27 /* Size in bytes of the array of initialization functions. */
	DT_FINI_ARRAYSZ DynTag = 28 /* Size in bytes of the array of termination functions. */
	DT_RUNPATH      DynTag = 29 /* String table offset of a null-terminated library search path string. */
	DT_FLAGS        DynTag = 30 /* Object specific flag values. */
	DT_ENCODING     DynTag = 32 /* Values greater than or equal to DT_ENCODING
	   and less than DT_LOOS follow the rules for
	   the interpretation of the d_un union
	   as follows: even == 'd_ptr', even == 'd_val'
	   or none */
	DT_PREINIT_ARRAY   DynTag = 32 /* Address of the array of pointers to pre-initialization functions. */
	DT_PREINIT_ARRAYSZ DynTag = 33 /* Size in bytes of the array of pre-initialization functions. */
	DT_SYMTAB_SHNDX    DynTag = 34 /* Address of SHT_SYMTAB_SHNDX section. */

	DT_LOOS DynTag = 0x6000000d /* First OS-specific */
	DT_HIOS DynTag = 0x6ffff000 /* Last OS-specific */

	DT_VALRNGLO       DynTag = 0x6ffffd00
	DT_GNU_PRELINKED  DynTag = 0x6ffffdf5
	DT_GNU_CONFLICTSZ DynTag = 0x6ffffdf6
	DT_GNU_LIBLISTSZ  DynTag = 0x6ffffdf7
	DT_CHECKSUM       DynTag = 0x6ffffdf8
	DT_PLTPADSZ       DynTag = 0x6ffffdf9
	DT_MOVEENT        DynTag = 0x6ffffdfa
	DT_MOVESZ         DynTag = 0x6ffffdfb
	DT_FEATURE        DynTag = 0x6ffffdfc
	DT_POSFLAG_1      DynTag = 0x6ffffdfd
	DT_SYMINSZ        DynTag = 0x6ffffdfe
	DT_SYMINENT       DynTag = 0x6ffffdff
	DT_VALRNGHI       DynTag = 0x6ffffdff

	DT_ADDRRNGLO    DynTag = 0x6ffffe00
	DT_GNU_HASH     DynTag = 0x6ffffef5
	DT_TLSDESC_PLT  DynTag = 0x6ffffef6
	DT_TLSDESC_GOT  DynTag = 0x6ffffef7
	DT_GNU_CONFLICT DynTag = 0x6ffffef8
	DT_GNU_LIBLIST  DynTag = 0x6ffffef9
	DT_CONFIG       DynTag = 0x6ffffefa
	DT_DEPAUDIT     DynTag = 0x6ffffefb
	DT_AUDIT        DynTag = 0x6ffffefc
	DT_PLTPAD       DynTag = 0x6ffffefd
	DT_MOVETAB      DynTag = 0x6ffffefe
	DT_SYMINFO      DynTag = 0x6ffffeff
	DT_ADDRRNGHI    DynTag = 0x6ffffeff

	DT_VERSYM     DynTag = 0x6ffffff0
	DT_RELACOUNT  DynTag = 0x6ffffff9
	DT_RELCOUNT   DynTag = 0x6ffffffa
	DT_FLAGS_1    DynTag = 0x6ffffffb
	DT_VERDEF     DynTag = 0x6ffffffc
	DT_VERDEFNUM  DynTag = 0x6ffffffd
	DT_VERNEED    DynTag = 0x6ffffffe
	DT_VERNEEDNUM DynTag = 0x6fffffff

	DT_LOPROC DynTag = 0x70000000 /* First processor-specific type. */

	DT_MIPS_RLD_VERSION           DynTag = 0x70000001
	DT_MIPS_TIME_STAMP            DynTag = 0x70000002
	DT_MIPS_ICHECKSUM             DynTag = 0x70000003
	DT_MIPS_IVERSION              DynTag = 0x70000004
	DT_MIPS_FLAGS                 DynTag = 0x70000005
	DT_MIPS_BASE_ADDRESS          DynTag = 0x70000006
	DT_MIPS_MSYM                  DynTag = 0x70000007
	DT_MIPS_CONFLICT              DynTag = 0x70000008
	DT_MIPS_LIBLIST               DynTag = 0x70000009
	DT_MIPS_LOCAL_GOTNO           DynTag = 0x7000000a
	DT_MIPS_CONFLICTNO            DynTag = 0x7000000b
	DT_MIPS_LIBLISTNO             DynTag = 0x70000010
	DT_MIPS_SYMTABNO              DynTag = 0x70000011
	DT_MIPS_UNREFEXTNO            DynTag = 0x70000012
	DT_MIPS_GOTSYM                DynTag = 0x70000013
	DT_MIPS_HIPAGENO              DynTag = 0x70000014
	DT_MIPS_RLD_MAP               DynTag = 0x70000016
	DT_MIPS_DELTA_CLASS           DynTag = 0x70000017
	DT_MIPS_DELTA_CLASS_NO        DynTag = 0x70000018
	DT_MIPS_DELTA_INSTANCE        DynTag = 0x70000019
	DT_MIPS_DELTA_INSTANCE_NO     DynTag = 0x7000001a
	DT_MIPS_DELTA_RELOC           DynTag = 0x7000001b
	DT_MIPS_DELTA_RELOC_NO        DynTag = 0x7000001c
	DT_MIPS_DELTA_SYM             DynTag = 0x7000001d
	DT_MIPS_DELTA_SYM_NO          DynTag = 0x7000001e
	DT_MIPS_DELTA_CLASSSYM        DynTag = 0x70000020
	DT_MIPS_DELTA_CLASSSYM_NO     DynTag = 0x70000021
	DT_MIPS_CXX_FLAGS             DynTag = 0x70000022
	DT_MIPS_PIXIE_INIT            DynTag = 0x70000023
	DT_MIPS_SYMBOL_LIB            DynTag = 0x70000024
	DT_MIPS_LOCALPAGE_GOTIDX      DynTag = 0x70000025
	DT_MIPS_LOCAL_GOTIDX          DynTag = 0x70000026
	DT_MIPS_HIDDEN_GOTIDX         DynTag = 0x70000027
	DT_MIPS_PROTECTED_GOTIDX      DynTag = 0x70000028
	DT_MIPS_OPTIONS               DynTag = 0x70000029
	DT_MIPS_INTERFACE             DynTag = 0x7000002a
	DT_MIPS_DYNSTR_ALIGN          DynTag = 0x7000002b
	DT_MIPS_INTERFACE_SIZE        DynTag = 0x7000002c
	DT_MIPS_RLD_TEXT_RESOLVE_ADDR DynTag = 0x7000002d
	DT_MIPS_PERF_SUFFIX           DynTag = 0x7000002e
	DT_MIPS_COMPACT_SIZE          DynTag = 0x7000002f
	DT_MIPS_GP_VALUE              DynTag = 0x70000030
	DT_MIPS_AUX_DYNAMIC           DynTag = 0x70000031
	DT_MIPS_PLTGOT                DynTag = 0x70000032
	DT_MIPS_RWPLT                 DynTag = 0x70000034
	DT_MIPS_RLD_MAP_REL           DynTag = 0x70000035

	DT_PPC_GOT DynTag = 0x70000000
	DT_PPC_OPT DynTag = 0x70000001

	DT_PPC64_GLINK DynTag = 0x70000000
	DT_PPC64_OPD   DynTag = 0x70000001
	DT_PPC64_OPDSZ DynTag = 0x70000002
	DT_PPC64_OPT   DynTag = 0x70000003

	DT_SPARC_REGISTER DynTag = 0x70000001

	DT_AUXILIARY DynTag = 0x7ffffffd
	DT_USED      DynTag = 0x7ffffffe
	DT_FILTER    DynTag = 0x7fffffff

	DT_HIPROC DynTag = 0x7fffffff /* Last processor-specific type. */
)
```

#### (DynTag) GoString 

``` go 
func (i DynTag) GoString() string
```

#### (DynTag) String 

``` go 
func (i DynTag) String() string
```

### type File 

``` go 
type File struct {
	FileHeader
	Sections []*Section
	Progs    []*Prog
	// contains filtered or unexported fields
}
```

A File represents an open ELF file.

#### func NewFile 

``` go 
func NewFile(r io.ReaderAt) (*File, error)
```

NewFile creates a new File for accessing an ELF binary in an underlying reader. The ELF binary is expected to start at position 0 in the ReaderAt.

#### func Open 

``` go 
func Open(name string) (*File, error)
```

Open opens the named file using os.Open and prepares it for use as an ELF binary.

#### (*File) Close 

``` go 
func (f *File) Close() error
```

Close closes the File. If the File was created using NewFile directly instead of Open, Close has no effect.

#### (*File) DWARF 

``` go 
func (f *File) DWARF() (*dwarf.Data, error)
```

#### (*File) DynString  <- go1.1

``` go 
func (f *File) DynString(tag DynTag) ([]string, error)
```

DynString returns the strings listed for the given tag in the file's dynamic section.

The tag must be one that takes string values: DT_NEEDED, DT_SONAME, DT_RPATH, or DT_RUNPATH.

#### (*File) DynamicSymbols  <- go1.4

``` go 
func (f *File) DynamicSymbols() ([]Symbol, error)
```

DynamicSymbols returns the dynamic symbol table for f. The symbols will be listed in the order they appear in f.

If f has a symbol version table, the returned Symbols will have initialized Version and Library fields.

For compatibility with Symbols, DynamicSymbols omits the null symbol at index 0. After retrieving the symbols as symtab, an externally supplied index x corresponds to symtab[x-1], not symtab[x].

#### (*File) ImportedLibraries 

``` go 
func (f *File) ImportedLibraries() ([]string, error)
```

ImportedLibraries returns the names of all libraries referred to by the binary f that are expected to be linked with the binary at dynamic link time.

#### (*File) ImportedSymbols 

``` go 
func (f *File) ImportedSymbols() ([]ImportedSymbol, error)
```

ImportedSymbols returns the names of all symbols referred to by the binary f that are expected to be satisfied by other libraries at dynamic load time. It does not return weak symbols.

#### (*File) Section 

``` go 
func (f *File) Section(name string) *Section
```

Section returns a section with the given name, or nil if no such section exists.

#### (*File) SectionByType 

``` go 
func (f *File) SectionByType(typ SectionType) *Section
```

SectionByType returns the first section in f with the given type, or nil if there is no such section.

#### (*File) Symbols 

``` go 
func (f *File) Symbols() ([]Symbol, error)
```

Symbols returns the symbol table for f. The symbols will be listed in the order they appear in f.

For compatibility with Go 1.0, Symbols omits the null symbol at index 0. After retrieving the symbols as symtab, an externally supplied index x corresponds to symtab[x-1], not symtab[x].

### type FileHeader 

``` go 
type FileHeader struct {
	Class      Class
	Data       Data
	Version    Version
	OSABI      OSABI
	ABIVersion uint8
	ByteOrder  binary.ByteOrder
	Type       Type
	Machine    Machine
	Entry      uint64
}
```

A FileHeader represents an ELF file header.

### type FormatError 

``` go 
type FormatError struct {
	// contains filtered or unexported fields
}
```

#### (*FormatError) Error 

``` go 
func (e *FormatError) Error() string
```

### type Header32 

``` go 
type Header32 struct {
	Ident     [EI_NIDENT]byte /* File identification. */
	Type      uint16          /* File type. */
	Machine   uint16          /* Machine architecture. */
	Version   uint32          /* ELF format version. */
	Entry     uint32          /* Entry point. */
	Phoff     uint32          /* Program header file offset. */
	Shoff     uint32          /* Section header file offset. */
	Flags     uint32          /* Architecture-specific flags. */
	Ehsize    uint16          /* Size of ELF header in bytes. */
	Phentsize uint16          /* Size of program header entry. */
	Phnum     uint16          /* Number of program header entries. */
	Shentsize uint16          /* Size of section header entry. */
	Shnum     uint16          /* Number of section header entries. */
	Shstrndx  uint16          /* Section name strings section. */
}
```

ELF32 File header.

### type Header64 

``` go 
type Header64 struct {
	Ident     [EI_NIDENT]byte /* File identification. */
	Type      uint16          /* File type. */
	Machine   uint16          /* Machine architecture. */
	Version   uint32          /* ELF format version. */
	Entry     uint64          /* Entry point. */
	Phoff     uint64          /* Program header file offset. */
	Shoff     uint64          /* Section header file offset. */
	Flags     uint32          /* Architecture-specific flags. */
	Ehsize    uint16          /* Size of ELF header in bytes. */
	Phentsize uint16          /* Size of program header entry. */
	Phnum     uint16          /* Number of program header entries. */
	Shentsize uint16          /* Size of section header entry. */
	Shnum     uint16          /* Number of section header entries. */
	Shstrndx  uint16          /* Section name strings section. */
}
```

ELF64 file header.

### type ImportedSymbol 

``` go 
type ImportedSymbol struct {
	Name    string
	Version string
	Library string
}
```

### type Machine 

``` go 
type Machine uint16
```

Machine is found in Header.Machine.

``` go 
const (
	EM_NONE          Machine = 0   /* Unknown machine. */
	EM_M32           Machine = 1   /* AT&T WE32100. */
	EM_SPARC         Machine = 2   /* Sun SPARC. */
	EM_386           Machine = 3   /* Intel i386. */
	EM_68K           Machine = 4   /* Motorola 68000. */
	EM_88K           Machine = 5   /* Motorola 88000. */
	EM_860           Machine = 7   /* Intel i860. */
	EM_MIPS          Machine = 8   /* MIPS R3000 Big-Endian only. */
	EM_S370          Machine = 9   /* IBM System/370. */
	EM_MIPS_RS3_LE   Machine = 10  /* MIPS R3000 Little-Endian. */
	EM_PARISC        Machine = 15  /* HP PA-RISC. */
	EM_VPP500        Machine = 17  /* Fujitsu VPP500. */
	EM_SPARC32PLUS   Machine = 18  /* SPARC v8plus. */
	EM_960           Machine = 19  /* Intel 80960. */
	EM_PPC           Machine = 20  /* PowerPC 32-bit. */
	EM_PPC64         Machine = 21  /* PowerPC 64-bit. */
	EM_S390          Machine = 22  /* IBM System/390. */
	EM_V800          Machine = 36  /* NEC V800. */
	EM_FR20          Machine = 37  /* Fujitsu FR20. */
	EM_RH32          Machine = 38  /* TRW RH-32. */
	EM_RCE           Machine = 39  /* Motorola RCE. */
	EM_ARM           Machine = 40  /* ARM. */
	EM_SH            Machine = 42  /* Hitachi SH. */
	EM_SPARCV9       Machine = 43  /* SPARC v9 64-bit. */
	EM_TRICORE       Machine = 44  /* Siemens TriCore embedded processor. */
	EM_ARC           Machine = 45  /* Argonaut RISC Core. */
	EM_H8_300        Machine = 46  /* Hitachi H8/300. */
	EM_H8_300H       Machine = 47  /* Hitachi H8/300H. */
	EM_H8S           Machine = 48  /* Hitachi H8S. */
	EM_H8_500        Machine = 49  /* Hitachi H8/500. */
	EM_IA_64         Machine = 50  /* Intel IA-64 Processor. */
	EM_MIPS_X        Machine = 51  /* Stanford MIPS-X. */
	EM_COLDFIRE      Machine = 52  /* Motorola ColdFire. */
	EM_68HC12        Machine = 53  /* Motorola M68HC12. */
	EM_MMA           Machine = 54  /* Fujitsu MMA. */
	EM_PCP           Machine = 55  /* Siemens PCP. */
	EM_NCPU          Machine = 56  /* Sony nCPU. */
	EM_NDR1          Machine = 57  /* Denso NDR1 microprocessor. */
	EM_STARCORE      Machine = 58  /* Motorola Star*Core processor. */
	EM_ME16          Machine = 59  /* Toyota ME16 processor. */
	EM_ST100         Machine = 60  /* STMicroelectronics ST100 processor. */
	EM_TINYJ         Machine = 61  /* Advanced Logic Corp. TinyJ processor. */
	EM_X86_64        Machine = 62  /* Advanced Micro Devices x86-64 */
	EM_PDSP          Machine = 63  /* Sony DSP Processor */
	EM_PDP10         Machine = 64  /* Digital Equipment Corp. PDP-10 */
	EM_PDP11         Machine = 65  /* Digital Equipment Corp. PDP-11 */
	EM_FX66          Machine = 66  /* Siemens FX66 microcontroller */
	EM_ST9PLUS       Machine = 67  /* STMicroelectronics ST9+ 8/16 bit microcontroller */
	EM_ST7           Machine = 68  /* STMicroelectronics ST7 8-bit microcontroller */
	EM_68HC16        Machine = 69  /* Motorola MC68HC16 Microcontroller */
	EM_68HC11        Machine = 70  /* Motorola MC68HC11 Microcontroller */
	EM_68HC08        Machine = 71  /* Motorola MC68HC08 Microcontroller */
	EM_68HC05        Machine = 72  /* Motorola MC68HC05 Microcontroller */
	EM_SVX           Machine = 73  /* Silicon Graphics SVx */
	EM_ST19          Machine = 74  /* STMicroelectronics ST19 8-bit microcontroller */
	EM_VAX           Machine = 75  /* Digital VAX */
	EM_CRIS          Machine = 76  /* Axis Communications 32-bit embedded processor */
	EM_JAVELIN       Machine = 77  /* Infineon Technologies 32-bit embedded processor */
	EM_FIREPATH      Machine = 78  /* Element 14 64-bit DSP Processor */
	EM_ZSP           Machine = 79  /* LSI Logic 16-bit DSP Processor */
	EM_MMIX          Machine = 80  /* Donald Knuth's educational 64-bit processor */
	EM_HUANY         Machine = 81  /* Harvard University machine-independent object files */
	EM_PRISM         Machine = 82  /* SiTera Prism */
	EM_AVR           Machine = 83  /* Atmel AVR 8-bit microcontroller */
	EM_FR30          Machine = 84  /* Fujitsu FR30 */
	EM_D10V          Machine = 85  /* Mitsubishi D10V */
	EM_D30V          Machine = 86  /* Mitsubishi D30V */
	EM_V850          Machine = 87  /* NEC v850 */
	EM_M32R          Machine = 88  /* Mitsubishi M32R */
	EM_MN10300       Machine = 89  /* Matsushita MN10300 */
	EM_MN10200       Machine = 90  /* Matsushita MN10200 */
	EM_PJ            Machine = 91  /* picoJava */
	EM_OPENRISC      Machine = 92  /* OpenRISC 32-bit embedded processor */
	EM_ARC_COMPACT   Machine = 93  /* ARC International ARCompact processor (old spelling/synonym: EM_ARC_A5) */
	EM_XTENSA        Machine = 94  /* Tensilica Xtensa Architecture */
	EM_VIDEOCORE     Machine = 95  /* Alphamosaic VideoCore processor */
	EM_TMM_GPP       Machine = 96  /* Thompson Multimedia General Purpose Processor */
	EM_NS32K         Machine = 97  /* National Semiconductor 32000 series */
	EM_TPC           Machine = 98  /* Tenor Network TPC processor */
	EM_SNP1K         Machine = 99  /* Trebia SNP 1000 processor */
	EM_ST200         Machine = 100 /* STMicroelectronics (www.st.com) ST200 microcontroller */
	EM_IP2K          Machine = 101 /* Ubicom IP2xxx microcontroller family */
	EM_MAX           Machine = 102 /* MAX Processor */
	EM_CR            Machine = 103 /* National Semiconductor CompactRISC microprocessor */
	EM_F2MC16        Machine = 104 /* Fujitsu F2MC16 */
	EM_MSP430        Machine = 105 /* Texas Instruments embedded microcontroller msp430 */
	EM_BLACKFIN      Machine = 106 /* Analog Devices Blackfin (DSP) processor */
	EM_SE_C33        Machine = 107 /* S1C33 Family of Seiko Epson processors */
	EM_SEP           Machine = 108 /* Sharp embedded microprocessor */
	EM_ARCA          Machine = 109 /* Arca RISC Microprocessor */
	EM_UNICORE       Machine = 110 /* Microprocessor series from PKU-Unity Ltd. and MPRC of Peking University */
	EM_EXCESS        Machine = 111 /* eXcess: 16/32/64-bit configurable embedded CPU */
	EM_DXP           Machine = 112 /* Icera Semiconductor Inc. Deep Execution Processor */
	EM_ALTERA_NIOS2  Machine = 113 /* Altera Nios II soft-core processor */
	EM_CRX           Machine = 114 /* National Semiconductor CompactRISC CRX microprocessor */
	EM_XGATE         Machine = 115 /* Motorola XGATE embedded processor */
	EM_C166          Machine = 116 /* Infineon C16x/XC16x processor */
	EM_M16C          Machine = 117 /* Renesas M16C series microprocessors */
	EM_DSPIC30F      Machine = 118 /* Microchip Technology dsPIC30F Digital Signal Controller */
	EM_CE            Machine = 119 /* Freescale Communication Engine RISC core */
	EM_M32C          Machine = 120 /* Renesas M32C series microprocessors */
	EM_TSK3000       Machine = 131 /* Altium TSK3000 core */
	EM_RS08          Machine = 132 /* Freescale RS08 embedded processor */
	EM_SHARC         Machine = 133 /* Analog Devices SHARC family of 32-bit DSP processors */
	EM_ECOG2         Machine = 134 /* Cyan Technology eCOG2 microprocessor */
	EM_SCORE7        Machine = 135 /* Sunplus S+core7 RISC processor */
	EM_DSP24         Machine = 136 /* New Japan Radio (NJR) 24-bit DSP Processor */
	EM_VIDEOCORE3    Machine = 137 /* Broadcom VideoCore III processor */
	EM_LATTICEMICO32 Machine = 138 /* RISC processor for Lattice FPGA architecture */
	EM_SE_C17        Machine = 139 /* Seiko Epson C17 family */
	EM_TI_C6000      Machine = 140 /* The Texas Instruments TMS320C6000 DSP family */
	EM_TI_C2000      Machine = 141 /* The Texas Instruments TMS320C2000 DSP family */
	EM_TI_C5500      Machine = 142 /* The Texas Instruments TMS320C55x DSP family */
	EM_TI_ARP32      Machine = 143 /* Texas Instruments Application Specific RISC Processor, 32bit fetch */
	EM_TI_PRU        Machine = 144 /* Texas Instruments Programmable Realtime Unit */
	EM_MMDSP_PLUS    Machine = 160 /* STMicroelectronics 64bit VLIW Data Signal Processor */
	EM_CYPRESS_M8C   Machine = 161 /* Cypress M8C microprocessor */
	EM_R32C          Machine = 162 /* Renesas R32C series microprocessors */
	EM_TRIMEDIA      Machine = 163 /* NXP Semiconductors TriMedia architecture family */
	EM_QDSP6         Machine = 164 /* QUALCOMM DSP6 Processor */
	EM_8051          Machine = 165 /* Intel 8051 and variants */
	EM_STXP7X        Machine = 166 /* STMicroelectronics STxP7x family of configurable and extensible RISC processors */
	EM_NDS32         Machine = 167 /* Andes Technology compact code size embedded RISC processor family */
	EM_ECOG1         Machine = 168 /* Cyan Technology eCOG1X family */
	EM_ECOG1X        Machine = 168 /* Cyan Technology eCOG1X family */
	EM_MAXQ30        Machine = 169 /* Dallas Semiconductor MAXQ30 Core Micro-controllers */
	EM_XIMO16        Machine = 170 /* New Japan Radio (NJR) 16-bit DSP Processor */
	EM_MANIK         Machine = 171 /* M2000 Reconfigurable RISC Microprocessor */
	EM_CRAYNV2       Machine = 172 /* Cray Inc. NV2 vector architecture */
	EM_RX            Machine = 173 /* Renesas RX family */
	EM_METAG         Machine = 174 /* Imagination Technologies META processor architecture */
	EM_MCST_ELBRUS   Machine = 175 /* MCST Elbrus general purpose hardware architecture */
	EM_ECOG16        Machine = 176 /* Cyan Technology eCOG16 family */
	EM_CR16          Machine = 177 /* National Semiconductor CompactRISC CR16 16-bit microprocessor */
	EM_ETPU          Machine = 178 /* Freescale Extended Time Processing Unit */
	EM_SLE9X         Machine = 179 /* Infineon Technologies SLE9X core */
	EM_L10M          Machine = 180 /* Intel L10M */
	EM_K10M          Machine = 181 /* Intel K10M */
	EM_AARCH64       Machine = 183 /* ARM 64-bit Architecture (AArch64) */
	EM_AVR32         Machine = 185 /* Atmel Corporation 32-bit microprocessor family */
	EM_STM8          Machine = 186 /* STMicroeletronics STM8 8-bit microcontroller */
	EM_TILE64        Machine = 187 /* Tilera TILE64 multicore architecture family */
	EM_TILEPRO       Machine = 188 /* Tilera TILEPro multicore architecture family */
	EM_MICROBLAZE    Machine = 189 /* Xilinx MicroBlaze 32-bit RISC soft processor core */
	EM_CUDA          Machine = 190 /* NVIDIA CUDA architecture */
	EM_TILEGX        Machine = 191 /* Tilera TILE-Gx multicore architecture family */
	EM_CLOUDSHIELD   Machine = 192 /* CloudShield architecture family */
	EM_COREA_1ST     Machine = 193 /* KIPO-KAIST Core-A 1st generation processor family */
	EM_COREA_2ND     Machine = 194 /* KIPO-KAIST Core-A 2nd generation processor family */
	EM_ARC_COMPACT2  Machine = 195 /* Synopsys ARCompact V2 */
	EM_OPEN8         Machine = 196 /* Open8 8-bit RISC soft processor core */
	EM_RL78          Machine = 197 /* Renesas RL78 family */
	EM_VIDEOCORE5    Machine = 198 /* Broadcom VideoCore V processor */
	EM_78KOR         Machine = 199 /* Renesas 78KOR family */
	EM_56800EX       Machine = 200 /* Freescale 56800EX Digital Signal Controller (DSC) */
	EM_BA1           Machine = 201 /* Beyond BA1 CPU architecture */
	EM_BA2           Machine = 202 /* Beyond BA2 CPU architecture */
	EM_XCORE         Machine = 203 /* XMOS xCORE processor family */
	EM_MCHP_PIC      Machine = 204 /* Microchip 8-bit PIC(r) family */
	EM_INTEL205      Machine = 205 /* Reserved by Intel */
	EM_INTEL206      Machine = 206 /* Reserved by Intel */
	EM_INTEL207      Machine = 207 /* Reserved by Intel */
	EM_INTEL208      Machine = 208 /* Reserved by Intel */
	EM_INTEL209      Machine = 209 /* Reserved by Intel */
	EM_KM32          Machine = 210 /* KM211 KM32 32-bit processor */
	EM_KMX32         Machine = 211 /* KM211 KMX32 32-bit processor */
	EM_KMX16         Machine = 212 /* KM211 KMX16 16-bit processor */
	EM_KMX8          Machine = 213 /* KM211 KMX8 8-bit processor */
	EM_KVARC         Machine = 214 /* KM211 KVARC processor */
	EM_CDP           Machine = 215 /* Paneve CDP architecture family */
	EM_COGE          Machine = 216 /* Cognitive Smart Memory Processor */
	EM_COOL          Machine = 217 /* Bluechip Systems CoolEngine */
	EM_NORC          Machine = 218 /* Nanoradio Optimized RISC */
	EM_CSR_KALIMBA   Machine = 219 /* CSR Kalimba architecture family */
	EM_Z80           Machine = 220 /* Zilog Z80 */
	EM_VISIUM        Machine = 221 /* Controls and Data Services VISIUMcore processor */
	EM_FT32          Machine = 222 /* FTDI Chip FT32 high performance 32-bit RISC architecture */
	EM_MOXIE         Machine = 223 /* Moxie processor family */
	EM_AMDGPU        Machine = 224 /* AMD GPU architecture */
	EM_RISCV         Machine = 243 /* RISC-V */
	EM_LANAI         Machine = 244 /* Lanai 32-bit processor */
	EM_BPF           Machine = 247 /* Linux BPF – in-kernel virtual machine */
	EM_LOONGARCH     Machine = 258 /* LoongArch */

	/* Non-standard or deprecated. */
	EM_486         Machine = 6      /* Intel i486. */
	EM_MIPS_RS4_BE Machine = 10     /* MIPS R4000 Big-Endian */
	EM_ALPHA_STD   Machine = 41     /* Digital Alpha (standard value). */
	EM_ALPHA       Machine = 0x9026 /* Alpha (written in the absence of an ABI) */
)
```

#### (Machine) GoString 

``` go 
func (i Machine) GoString() string
```

#### (Machine) String 

``` go 
func (i Machine) String() string
```

### type NType 

``` go 
type NType int
```

NType values; used in core files.

``` go 
const (
	NT_PRSTATUS NType = 1 /* Process status. */
	NT_FPREGSET NType = 2 /* Floating point registers. */
	NT_PRPSINFO NType = 3 /* Process state info. */
)
```

#### (NType) GoString 

``` go 
func (i NType) GoString() string
```

#### (NType) String 

``` go 
func (i NType) String() string
```

### type OSABI 

``` go 
type OSABI byte
```

OSABI is found in Header.Ident[EI_OSABI] and Header.OSABI.

``` go 
const (
	ELFOSABI_NONE       OSABI = 0   /* UNIX System V ABI */
	ELFOSABI_HPUX       OSABI = 1   /* HP-UX operating system */
	ELFOSABI_NETBSD     OSABI = 2   /* NetBSD */
	ELFOSABI_LINUX      OSABI = 3   /* Linux */
	ELFOSABI_HURD       OSABI = 4   /* Hurd */
	ELFOSABI_86OPEN     OSABI = 5   /* 86Open common IA32 ABI */
	ELFOSABI_SOLARIS    OSABI = 6   /* Solaris */
	ELFOSABI_AIX        OSABI = 7   /* AIX */
	ELFOSABI_IRIX       OSABI = 8   /* IRIX */
	ELFOSABI_FREEBSD    OSABI = 9   /* FreeBSD */
	ELFOSABI_TRU64      OSABI = 10  /* TRU64 UNIX */
	ELFOSABI_MODESTO    OSABI = 11  /* Novell Modesto */
	ELFOSABI_OPENBSD    OSABI = 12  /* OpenBSD */
	ELFOSABI_OPENVMS    OSABI = 13  /* Open VMS */
	ELFOSABI_NSK        OSABI = 14  /* HP Non-Stop Kernel */
	ELFOSABI_AROS       OSABI = 15  /* Amiga Research OS */
	ELFOSABI_FENIXOS    OSABI = 16  /* The FenixOS highly scalable multi-core OS */
	ELFOSABI_CLOUDABI   OSABI = 17  /* Nuxi CloudABI */
	ELFOSABI_ARM        OSABI = 97  /* ARM */
	ELFOSABI_STANDALONE OSABI = 255 /* Standalone (embedded) application */
)
```

#### (OSABI) GoString 

``` go 
func (i OSABI) GoString() string
```

#### (OSABI) String 

``` go 
func (i OSABI) String() string
```

### type Prog 

``` go 
type Prog struct {
	ProgHeader

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

A Prog represents a single ELF program header in an ELF binary.

#### (*Prog) Open 

``` go 
func (p *Prog) Open() io.ReadSeeker
```

Open returns a new ReadSeeker reading the ELF program body.

### type Prog32 

``` go 
type Prog32 struct {
	Type   uint32 /* Entry type. */
	Off    uint32 /* File offset of contents. */
	Vaddr  uint32 /* Virtual address in memory image. */
	Paddr  uint32 /* Physical address (not used). */
	Filesz uint32 /* Size of contents in file. */
	Memsz  uint32 /* Size of contents in memory. */
	Flags  uint32 /* Access permission flags. */
	Align  uint32 /* Alignment in memory and file. */
}
```

ELF32 Program header.

### type Prog64 

``` go 
type Prog64 struct {
	Type   uint32 /* Entry type. */
	Flags  uint32 /* Access permission flags. */
	Off    uint64 /* File offset of contents. */
	Vaddr  uint64 /* Virtual address in memory image. */
	Paddr  uint64 /* Physical address (not used). */
	Filesz uint64 /* Size of contents in file. */
	Memsz  uint64 /* Size of contents in memory. */
	Align  uint64 /* Alignment in memory and file. */
}
```

ELF64 Program header.

### type ProgFlag 

``` go 
type ProgFlag uint32
```

Prog.Flag

``` go 
const (
	PF_X        ProgFlag = 0x1        /* Executable. */
	PF_W        ProgFlag = 0x2        /* Writable. */
	PF_R        ProgFlag = 0x4        /* Readable. */
	PF_MASKOS   ProgFlag = 0x0ff00000 /* Operating system-specific. */
	PF_MASKPROC ProgFlag = 0xf0000000 /* Processor-specific. */
)
```

#### (ProgFlag) GoString 

``` go 
func (i ProgFlag) GoString() string
```

#### (ProgFlag) String 

``` go 
func (i ProgFlag) String() string
```

### type ProgHeader 

``` go 
type ProgHeader struct {
	Type   ProgType
	Flags  ProgFlag
	Off    uint64
	Vaddr  uint64
	Paddr  uint64
	Filesz uint64
	Memsz  uint64
	Align  uint64
}
```

A ProgHeader represents a single ELF program header.

### type ProgType 

``` go 
type ProgType int
```

Prog.Type

``` go 
const (
	PT_NULL    ProgType = 0 /* Unused entry. */
	PT_LOAD    ProgType = 1 /* Loadable segment. */
	PT_DYNAMIC ProgType = 2 /* Dynamic linking information segment. */
	PT_INTERP  ProgType = 3 /* Pathname of interpreter. */
	PT_NOTE    ProgType = 4 /* Auxiliary information. */
	PT_SHLIB   ProgType = 5 /* Reserved (not used). */
	PT_PHDR    ProgType = 6 /* Location of program header itself. */
	PT_TLS     ProgType = 7 /* Thread local storage segment */

	PT_LOOS ProgType = 0x60000000 /* First OS-specific. */

	PT_GNU_EH_FRAME ProgType = 0x6474e550 /* Frame unwind information */
	PT_GNU_STACK    ProgType = 0x6474e551 /* Stack flags */
	PT_GNU_RELRO    ProgType = 0x6474e552 /* Read only after relocs */
	PT_GNU_PROPERTY ProgType = 0x6474e553 /* GNU property */
	PT_GNU_MBIND_LO ProgType = 0x6474e555 /* Mbind segments start */
	PT_GNU_MBIND_HI ProgType = 0x6474f554 /* Mbind segments finish */

	PT_PAX_FLAGS ProgType = 0x65041580 /* PAX flags */

	PT_OPENBSD_RANDOMIZE ProgType = 0x65a3dbe6 /* Random data */
	PT_OPENBSD_WXNEEDED  ProgType = 0x65a3dbe7 /* W^X violations */
	PT_OPENBSD_BOOTDATA  ProgType = 0x65a41be6 /* Boot arguments */

	PT_SUNW_EH_FRAME ProgType = 0x6474e550 /* Frame unwind information */
	PT_SUNWSTACK     ProgType = 0x6ffffffb /* Stack segment */

	PT_HIOS ProgType = 0x6fffffff /* Last OS-specific. */

	PT_LOPROC ProgType = 0x70000000 /* First processor-specific type. */

	PT_ARM_ARCHEXT ProgType = 0x70000000 /* Architecture compatibility */
	PT_ARM_EXIDX   ProgType = 0x70000001 /* Exception unwind tables */

	PT_AARCH64_ARCHEXT ProgType = 0x70000000 /* Architecture compatibility */
	PT_AARCH64_UNWIND  ProgType = 0x70000001 /* Exception unwind tables */

	PT_MIPS_REGINFO  ProgType = 0x70000000 /* Register usage */
	PT_MIPS_RTPROC   ProgType = 0x70000001 /* Runtime procedures */
	PT_MIPS_OPTIONS  ProgType = 0x70000002 /* Options */
	PT_MIPS_ABIFLAGS ProgType = 0x70000003 /* ABI flags */

	PT_S390_PGSTE ProgType = 0x70000000 /* 4k page table size */

	PT_HIPROC ProgType = 0x7fffffff /* Last processor-specific type. */
)
```

#### (ProgType) GoString 

``` go 
func (i ProgType) GoString() string
```

#### (ProgType) String 

``` go 
func (i ProgType) String() string
```

### type R_386 

``` go 
type R_386 int
```

Relocation types for 386.

``` go 
const (
	R_386_NONE          R_386 = 0  /* No relocation. */
	R_386_32            R_386 = 1  /* Add symbol value. */
	R_386_PC32          R_386 = 2  /* Add PC-relative symbol value. */
	R_386_GOT32         R_386 = 3  /* Add PC-relative GOT offset. */
	R_386_PLT32         R_386 = 4  /* Add PC-relative PLT offset. */
	R_386_COPY          R_386 = 5  /* Copy data from shared object. */
	R_386_GLOB_DAT      R_386 = 6  /* Set GOT entry to data address. */
	R_386_JMP_SLOT      R_386 = 7  /* Set GOT entry to code address. */
	R_386_RELATIVE      R_386 = 8  /* Add load address of shared object. */
	R_386_GOTOFF        R_386 = 9  /* Add GOT-relative symbol address. */
	R_386_GOTPC         R_386 = 10 /* Add PC-relative GOT table address. */
	R_386_32PLT         R_386 = 11
	R_386_TLS_TPOFF     R_386 = 14 /* Negative offset in static TLS block */
	R_386_TLS_IE        R_386 = 15 /* Absolute address of GOT for -ve static TLS */
	R_386_TLS_GOTIE     R_386 = 16 /* GOT entry for negative static TLS block */
	R_386_TLS_LE        R_386 = 17 /* Negative offset relative to static TLS */
	R_386_TLS_GD        R_386 = 18 /* 32 bit offset to GOT (index,off) pair */
	R_386_TLS_LDM       R_386 = 19 /* 32 bit offset to GOT (index,zero) pair */
	R_386_16            R_386 = 20
	R_386_PC16          R_386 = 21
	R_386_8             R_386 = 22
	R_386_PC8           R_386 = 23
	R_386_TLS_GD_32     R_386 = 24 /* 32 bit offset to GOT (index,off) pair */
	R_386_TLS_GD_PUSH   R_386 = 25 /* pushl instruction for Sun ABI GD sequence */
	R_386_TLS_GD_CALL   R_386 = 26 /* call instruction for Sun ABI GD sequence */
	R_386_TLS_GD_POP    R_386 = 27 /* popl instruction for Sun ABI GD sequence */
	R_386_TLS_LDM_32    R_386 = 28 /* 32 bit offset to GOT (index,zero) pair */
	R_386_TLS_LDM_PUSH  R_386 = 29 /* pushl instruction for Sun ABI LD sequence */
	R_386_TLS_LDM_CALL  R_386 = 30 /* call instruction for Sun ABI LD sequence */
	R_386_TLS_LDM_POP   R_386 = 31 /* popl instruction for Sun ABI LD sequence */
	R_386_TLS_LDO_32    R_386 = 32 /* 32 bit offset from start of TLS block */
	R_386_TLS_IE_32     R_386 = 33 /* 32 bit offset to GOT static TLS offset entry */
	R_386_TLS_LE_32     R_386 = 34 /* 32 bit offset within static TLS block */
	R_386_TLS_DTPMOD32  R_386 = 35 /* GOT entry containing TLS index */
	R_386_TLS_DTPOFF32  R_386 = 36 /* GOT entry containing TLS offset */
	R_386_TLS_TPOFF32   R_386 = 37 /* GOT entry of -ve static TLS offset */
	R_386_SIZE32        R_386 = 38
	R_386_TLS_GOTDESC   R_386 = 39
	R_386_TLS_DESC_CALL R_386 = 40
	R_386_TLS_DESC      R_386 = 41
	R_386_IRELATIVE     R_386 = 42
	R_386_GOT32X        R_386 = 43
)
```

#### (R_386) GoString 

``` go 
func (i R_386) GoString() string
```

#### (R_386) String 

``` go 
func (i R_386) String() string
```

### type R_390  <- go1.7

``` go 
type R_390 int
```

Relocation types for s390x processors.

``` go 
const (
	R_390_NONE        R_390 = 0
	R_390_8           R_390 = 1
	R_390_12          R_390 = 2
	R_390_16          R_390 = 3
	R_390_32          R_390 = 4
	R_390_PC32        R_390 = 5
	R_390_GOT12       R_390 = 6
	R_390_GOT32       R_390 = 7
	R_390_PLT32       R_390 = 8
	R_390_COPY        R_390 = 9
	R_390_GLOB_DAT    R_390 = 10
	R_390_JMP_SLOT    R_390 = 11
	R_390_RELATIVE    R_390 = 12
	R_390_GOTOFF      R_390 = 13
	R_390_GOTPC       R_390 = 14
	R_390_GOT16       R_390 = 15
	R_390_PC16        R_390 = 16
	R_390_PC16DBL     R_390 = 17
	R_390_PLT16DBL    R_390 = 18
	R_390_PC32DBL     R_390 = 19
	R_390_PLT32DBL    R_390 = 20
	R_390_GOTPCDBL    R_390 = 21
	R_390_64          R_390 = 22
	R_390_PC64        R_390 = 23
	R_390_GOT64       R_390 = 24
	R_390_PLT64       R_390 = 25
	R_390_GOTENT      R_390 = 26
	R_390_GOTOFF16    R_390 = 27
	R_390_GOTOFF64    R_390 = 28
	R_390_GOTPLT12    R_390 = 29
	R_390_GOTPLT16    R_390 = 30
	R_390_GOTPLT32    R_390 = 31
	R_390_GOTPLT64    R_390 = 32
	R_390_GOTPLTENT   R_390 = 33
	R_390_GOTPLTOFF16 R_390 = 34
	R_390_GOTPLTOFF32 R_390 = 35
	R_390_GOTPLTOFF64 R_390 = 36
	R_390_TLS_LOAD    R_390 = 37
	R_390_TLS_GDCALL  R_390 = 38
	R_390_TLS_LDCALL  R_390 = 39
	R_390_TLS_GD32    R_390 = 40
	R_390_TLS_GD64    R_390 = 41
	R_390_TLS_GOTIE12 R_390 = 42
	R_390_TLS_GOTIE32 R_390 = 43
	R_390_TLS_GOTIE64 R_390 = 44
	R_390_TLS_LDM32   R_390 = 45
	R_390_TLS_LDM64   R_390 = 46
	R_390_TLS_IE32    R_390 = 47
	R_390_TLS_IE64    R_390 = 48
	R_390_TLS_IEENT   R_390 = 49
	R_390_TLS_LE32    R_390 = 50
	R_390_TLS_LE64    R_390 = 51
	R_390_TLS_LDO32   R_390 = 52
	R_390_TLS_LDO64   R_390 = 53
	R_390_TLS_DTPMOD  R_390 = 54
	R_390_TLS_DTPOFF  R_390 = 55
	R_390_TLS_TPOFF   R_390 = 56
	R_390_20          R_390 = 57
	R_390_GOT20       R_390 = 58
	R_390_GOTPLT20    R_390 = 59
	R_390_TLS_GOTIE20 R_390 = 60
)
```

#### (R_390) GoString  <- go1.7

``` go 
func (i R_390) GoString() string
```

#### (R_390) String  <- go1.7

``` go 
func (i R_390) String() string
```

### type R_AARCH64  <- go1.4

``` go 
type R_AARCH64 int
```

Relocation types for AArch64 (aka arm64)

``` go 
const (
	R_AARCH64_NONE                            R_AARCH64 = 0
	R_AARCH64_P32_ABS32                       R_AARCH64 = 1
	R_AARCH64_P32_ABS16                       R_AARCH64 = 2
	R_AARCH64_P32_PREL32                      R_AARCH64 = 3
	R_AARCH64_P32_PREL16                      R_AARCH64 = 4
	R_AARCH64_P32_MOVW_UABS_G0                R_AARCH64 = 5
	R_AARCH64_P32_MOVW_UABS_G0_NC             R_AARCH64 = 6
	R_AARCH64_P32_MOVW_UABS_G1                R_AARCH64 = 7
	R_AARCH64_P32_MOVW_SABS_G0                R_AARCH64 = 8
	R_AARCH64_P32_LD_PREL_LO19                R_AARCH64 = 9
	R_AARCH64_P32_ADR_PREL_LO21               R_AARCH64 = 10
	R_AARCH64_P32_ADR_PREL_PG_HI21            R_AARCH64 = 11
	R_AARCH64_P32_ADD_ABS_LO12_NC             R_AARCH64 = 12
	R_AARCH64_P32_LDST8_ABS_LO12_NC           R_AARCH64 = 13
	R_AARCH64_P32_LDST16_ABS_LO12_NC          R_AARCH64 = 14
	R_AARCH64_P32_LDST32_ABS_LO12_NC          R_AARCH64 = 15
	R_AARCH64_P32_LDST64_ABS_LO12_NC          R_AARCH64 = 16
	R_AARCH64_P32_LDST128_ABS_LO12_NC         R_AARCH64 = 17
	R_AARCH64_P32_TSTBR14                     R_AARCH64 = 18
	R_AARCH64_P32_CONDBR19                    R_AARCH64 = 19
	R_AARCH64_P32_JUMP26                      R_AARCH64 = 20
	R_AARCH64_P32_CALL26                      R_AARCH64 = 21
	R_AARCH64_P32_GOT_LD_PREL19               R_AARCH64 = 25
	R_AARCH64_P32_ADR_GOT_PAGE                R_AARCH64 = 26
	R_AARCH64_P32_LD32_GOT_LO12_NC            R_AARCH64 = 27
	R_AARCH64_P32_TLSGD_ADR_PAGE21            R_AARCH64 = 81
	R_AARCH64_P32_TLSGD_ADD_LO12_NC           R_AARCH64 = 82
	R_AARCH64_P32_TLSIE_ADR_GOTTPREL_PAGE21   R_AARCH64 = 103
	R_AARCH64_P32_TLSIE_LD32_GOTTPREL_LO12_NC R_AARCH64 = 104
	R_AARCH64_P32_TLSIE_LD_GOTTPREL_PREL19    R_AARCH64 = 105
	R_AARCH64_P32_TLSLE_MOVW_TPREL_G1         R_AARCH64 = 106
	R_AARCH64_P32_TLSLE_MOVW_TPREL_G0         R_AARCH64 = 107
	R_AARCH64_P32_TLSLE_MOVW_TPREL_G0_NC      R_AARCH64 = 108
	R_AARCH64_P32_TLSLE_ADD_TPREL_HI12        R_AARCH64 = 109
	R_AARCH64_P32_TLSLE_ADD_TPREL_LO12        R_AARCH64 = 110
	R_AARCH64_P32_TLSLE_ADD_TPREL_LO12_NC     R_AARCH64 = 111
	R_AARCH64_P32_TLSDESC_LD_PREL19           R_AARCH64 = 122
	R_AARCH64_P32_TLSDESC_ADR_PREL21          R_AARCH64 = 123
	R_AARCH64_P32_TLSDESC_ADR_PAGE21          R_AARCH64 = 124
	R_AARCH64_P32_TLSDESC_LD32_LO12_NC        R_AARCH64 = 125
	R_AARCH64_P32_TLSDESC_ADD_LO12_NC         R_AARCH64 = 126
	R_AARCH64_P32_TLSDESC_CALL                R_AARCH64 = 127
	R_AARCH64_P32_COPY                        R_AARCH64 = 180
	R_AARCH64_P32_GLOB_DAT                    R_AARCH64 = 181
	R_AARCH64_P32_JUMP_SLOT                   R_AARCH64 = 182
	R_AARCH64_P32_RELATIVE                    R_AARCH64 = 183
	R_AARCH64_P32_TLS_DTPMOD                  R_AARCH64 = 184
	R_AARCH64_P32_TLS_DTPREL                  R_AARCH64 = 185
	R_AARCH64_P32_TLS_TPREL                   R_AARCH64 = 186
	R_AARCH64_P32_TLSDESC                     R_AARCH64 = 187
	R_AARCH64_P32_IRELATIVE                   R_AARCH64 = 188
	R_AARCH64_NULL                            R_AARCH64 = 256
	R_AARCH64_ABS64                           R_AARCH64 = 257
	R_AARCH64_ABS32                           R_AARCH64 = 258
	R_AARCH64_ABS16                           R_AARCH64 = 259
	R_AARCH64_PREL64                          R_AARCH64 = 260
	R_AARCH64_PREL32                          R_AARCH64 = 261
	R_AARCH64_PREL16                          R_AARCH64 = 262
	R_AARCH64_MOVW_UABS_G0                    R_AARCH64 = 263
	R_AARCH64_MOVW_UABS_G0_NC                 R_AARCH64 = 264
	R_AARCH64_MOVW_UABS_G1                    R_AARCH64 = 265
	R_AARCH64_MOVW_UABS_G1_NC                 R_AARCH64 = 266
	R_AARCH64_MOVW_UABS_G2                    R_AARCH64 = 267
	R_AARCH64_MOVW_UABS_G2_NC                 R_AARCH64 = 268
	R_AARCH64_MOVW_UABS_G3                    R_AARCH64 = 269
	R_AARCH64_MOVW_SABS_G0                    R_AARCH64 = 270
	R_AARCH64_MOVW_SABS_G1                    R_AARCH64 = 271
	R_AARCH64_MOVW_SABS_G2                    R_AARCH64 = 272
	R_AARCH64_LD_PREL_LO19                    R_AARCH64 = 273
	R_AARCH64_ADR_PREL_LO21                   R_AARCH64 = 274
	R_AARCH64_ADR_PREL_PG_HI21                R_AARCH64 = 275
	R_AARCH64_ADR_PREL_PG_HI21_NC             R_AARCH64 = 276
	R_AARCH64_ADD_ABS_LO12_NC                 R_AARCH64 = 277
	R_AARCH64_LDST8_ABS_LO12_NC               R_AARCH64 = 278
	R_AARCH64_TSTBR14                         R_AARCH64 = 279
	R_AARCH64_CONDBR19                        R_AARCH64 = 280
	R_AARCH64_JUMP26                          R_AARCH64 = 282
	R_AARCH64_CALL26                          R_AARCH64 = 283
	R_AARCH64_LDST16_ABS_LO12_NC              R_AARCH64 = 284
	R_AARCH64_LDST32_ABS_LO12_NC              R_AARCH64 = 285
	R_AARCH64_LDST64_ABS_LO12_NC              R_AARCH64 = 286
	R_AARCH64_LDST128_ABS_LO12_NC             R_AARCH64 = 299
	R_AARCH64_GOT_LD_PREL19                   R_AARCH64 = 309
	R_AARCH64_LD64_GOTOFF_LO15                R_AARCH64 = 310
	R_AARCH64_ADR_GOT_PAGE                    R_AARCH64 = 311
	R_AARCH64_LD64_GOT_LO12_NC                R_AARCH64 = 312
	R_AARCH64_LD64_GOTPAGE_LO15               R_AARCH64 = 313
	R_AARCH64_TLSGD_ADR_PREL21                R_AARCH64 = 512
	R_AARCH64_TLSGD_ADR_PAGE21                R_AARCH64 = 513
	R_AARCH64_TLSGD_ADD_LO12_NC               R_AARCH64 = 514
	R_AARCH64_TLSGD_MOVW_G1                   R_AARCH64 = 515
	R_AARCH64_TLSGD_MOVW_G0_NC                R_AARCH64 = 516
	R_AARCH64_TLSLD_ADR_PREL21                R_AARCH64 = 517
	R_AARCH64_TLSLD_ADR_PAGE21                R_AARCH64 = 518
	R_AARCH64_TLSIE_MOVW_GOTTPREL_G1          R_AARCH64 = 539
	R_AARCH64_TLSIE_MOVW_GOTTPREL_G0_NC       R_AARCH64 = 540
	R_AARCH64_TLSIE_ADR_GOTTPREL_PAGE21       R_AARCH64 = 541
	R_AARCH64_TLSIE_LD64_GOTTPREL_LO12_NC     R_AARCH64 = 542
	R_AARCH64_TLSIE_LD_GOTTPREL_PREL19        R_AARCH64 = 543
	R_AARCH64_TLSLE_MOVW_TPREL_G2             R_AARCH64 = 544
	R_AARCH64_TLSLE_MOVW_TPREL_G1             R_AARCH64 = 545
	R_AARCH64_TLSLE_MOVW_TPREL_G1_NC          R_AARCH64 = 546
	R_AARCH64_TLSLE_MOVW_TPREL_G0             R_AARCH64 = 547
	R_AARCH64_TLSLE_MOVW_TPREL_G0_NC          R_AARCH64 = 548
	R_AARCH64_TLSLE_ADD_TPREL_HI12            R_AARCH64 = 549
	R_AARCH64_TLSLE_ADD_TPREL_LO12            R_AARCH64 = 550
	R_AARCH64_TLSLE_ADD_TPREL_LO12_NC         R_AARCH64 = 551
	R_AARCH64_TLSDESC_LD_PREL19               R_AARCH64 = 560
	R_AARCH64_TLSDESC_ADR_PREL21              R_AARCH64 = 561
	R_AARCH64_TLSDESC_ADR_PAGE21              R_AARCH64 = 562
	R_AARCH64_TLSDESC_LD64_LO12_NC            R_AARCH64 = 563
	R_AARCH64_TLSDESC_ADD_LO12_NC             R_AARCH64 = 564
	R_AARCH64_TLSDESC_OFF_G1                  R_AARCH64 = 565
	R_AARCH64_TLSDESC_OFF_G0_NC               R_AARCH64 = 566
	R_AARCH64_TLSDESC_LDR                     R_AARCH64 = 567
	R_AARCH64_TLSDESC_ADD                     R_AARCH64 = 568
	R_AARCH64_TLSDESC_CALL                    R_AARCH64 = 569
	R_AARCH64_TLSLE_LDST128_TPREL_LO12        R_AARCH64 = 570
	R_AARCH64_TLSLE_LDST128_TPREL_LO12_NC     R_AARCH64 = 571
	R_AARCH64_TLSLD_LDST128_DTPREL_LO12       R_AARCH64 = 572
	R_AARCH64_TLSLD_LDST128_DTPREL_LO12_NC    R_AARCH64 = 573
	R_AARCH64_COPY                            R_AARCH64 = 1024
	R_AARCH64_GLOB_DAT                        R_AARCH64 = 1025
	R_AARCH64_JUMP_SLOT                       R_AARCH64 = 1026
	R_AARCH64_RELATIVE                        R_AARCH64 = 1027
	R_AARCH64_TLS_DTPMOD64                    R_AARCH64 = 1028
	R_AARCH64_TLS_DTPREL64                    R_AARCH64 = 1029
	R_AARCH64_TLS_TPREL64                     R_AARCH64 = 1030
	R_AARCH64_TLSDESC                         R_AARCH64 = 1031
	R_AARCH64_IRELATIVE                       R_AARCH64 = 1032
)
```

#### (R_AARCH64) GoString  <- go1.4

``` go 
func (i R_AARCH64) GoString() string
```

#### (R_AARCH64) String  <- go1.4

``` go 
func (i R_AARCH64) String() string
```

### type R_ALPHA 

``` go 
type R_ALPHA int
```

Relocation types for Alpha.

``` go 
const (
	R_ALPHA_NONE           R_ALPHA = 0  /* No reloc */
	R_ALPHA_REFLONG        R_ALPHA = 1  /* Direct 32 bit */
	R_ALPHA_REFQUAD        R_ALPHA = 2  /* Direct 64 bit */
	R_ALPHA_GPREL32        R_ALPHA = 3  /* GP relative 32 bit */
	R_ALPHA_LITERAL        R_ALPHA = 4  /* GP relative 16 bit w/optimization */
	R_ALPHA_LITUSE         R_ALPHA = 5  /* Optimization hint for LITERAL */
	R_ALPHA_GPDISP         R_ALPHA = 6  /* Add displacement to GP */
	R_ALPHA_BRADDR         R_ALPHA = 7  /* PC+4 relative 23 bit shifted */
	R_ALPHA_HINT           R_ALPHA = 8  /* PC+4 relative 16 bit shifted */
	R_ALPHA_SREL16         R_ALPHA = 9  /* PC relative 16 bit */
	R_ALPHA_SREL32         R_ALPHA = 10 /* PC relative 32 bit */
	R_ALPHA_SREL64         R_ALPHA = 11 /* PC relative 64 bit */
	R_ALPHA_OP_PUSH        R_ALPHA = 12 /* OP stack push */
	R_ALPHA_OP_STORE       R_ALPHA = 13 /* OP stack pop and store */
	R_ALPHA_OP_PSUB        R_ALPHA = 14 /* OP stack subtract */
	R_ALPHA_OP_PRSHIFT     R_ALPHA = 15 /* OP stack right shift */
	R_ALPHA_GPVALUE        R_ALPHA = 16
	R_ALPHA_GPRELHIGH      R_ALPHA = 17
	R_ALPHA_GPRELLOW       R_ALPHA = 18
	R_ALPHA_IMMED_GP_16    R_ALPHA = 19
	R_ALPHA_IMMED_GP_HI32  R_ALPHA = 20
	R_ALPHA_IMMED_SCN_HI32 R_ALPHA = 21
	R_ALPHA_IMMED_BR_HI32  R_ALPHA = 22
	R_ALPHA_IMMED_LO32     R_ALPHA = 23
	R_ALPHA_COPY           R_ALPHA = 24 /* Copy symbol at runtime */
	R_ALPHA_GLOB_DAT       R_ALPHA = 25 /* Create GOT entry */
	R_ALPHA_JMP_SLOT       R_ALPHA = 26 /* Create PLT entry */
	R_ALPHA_RELATIVE       R_ALPHA = 27 /* Adjust by program base */
)
```

#### (R_ALPHA) GoString 

``` go 
func (i R_ALPHA) GoString() string
```

#### (R_ALPHA) String 

``` go 
func (i R_ALPHA) String() string
```

### type R_ARM 

``` go 
type R_ARM int
```

Relocation types for ARM.

``` go 
const (
	R_ARM_NONE               R_ARM = 0 /* No relocation. */
	R_ARM_PC24               R_ARM = 1
	R_ARM_ABS32              R_ARM = 2
	R_ARM_REL32              R_ARM = 3
	R_ARM_PC13               R_ARM = 4
	R_ARM_ABS16              R_ARM = 5
	R_ARM_ABS12              R_ARM = 6
	R_ARM_THM_ABS5           R_ARM = 7
	R_ARM_ABS8               R_ARM = 8
	R_ARM_SBREL32            R_ARM = 9
	R_ARM_THM_PC22           R_ARM = 10
	R_ARM_THM_PC8            R_ARM = 11
	R_ARM_AMP_VCALL9         R_ARM = 12
	R_ARM_SWI24              R_ARM = 13
	R_ARM_THM_SWI8           R_ARM = 14
	R_ARM_XPC25              R_ARM = 15
	R_ARM_THM_XPC22          R_ARM = 16
	R_ARM_TLS_DTPMOD32       R_ARM = 17
	R_ARM_TLS_DTPOFF32       R_ARM = 18
	R_ARM_TLS_TPOFF32        R_ARM = 19
	R_ARM_COPY               R_ARM = 20 /* Copy data from shared object. */
	R_ARM_GLOB_DAT           R_ARM = 21 /* Set GOT entry to data address. */
	R_ARM_JUMP_SLOT          R_ARM = 22 /* Set GOT entry to code address. */
	R_ARM_RELATIVE           R_ARM = 23 /* Add load address of shared object. */
	R_ARM_GOTOFF             R_ARM = 24 /* Add GOT-relative symbol address. */
	R_ARM_GOTPC              R_ARM = 25 /* Add PC-relative GOT table address. */
	R_ARM_GOT32              R_ARM = 26 /* Add PC-relative GOT offset. */
	R_ARM_PLT32              R_ARM = 27 /* Add PC-relative PLT offset. */
	R_ARM_CALL               R_ARM = 28
	R_ARM_JUMP24             R_ARM = 29
	R_ARM_THM_JUMP24         R_ARM = 30
	R_ARM_BASE_ABS           R_ARM = 31
	R_ARM_ALU_PCREL_7_0      R_ARM = 32
	R_ARM_ALU_PCREL_15_8     R_ARM = 33
	R_ARM_ALU_PCREL_23_15    R_ARM = 34
	R_ARM_LDR_SBREL_11_10_NC R_ARM = 35
	R_ARM_ALU_SBREL_19_12_NC R_ARM = 36
	R_ARM_ALU_SBREL_27_20_CK R_ARM = 37
	R_ARM_TARGET1            R_ARM = 38
	R_ARM_SBREL31            R_ARM = 39
	R_ARM_V4BX               R_ARM = 40
	R_ARM_TARGET2            R_ARM = 41
	R_ARM_PREL31             R_ARM = 42
	R_ARM_MOVW_ABS_NC        R_ARM = 43
	R_ARM_MOVT_ABS           R_ARM = 44
	R_ARM_MOVW_PREL_NC       R_ARM = 45
	R_ARM_MOVT_PREL          R_ARM = 46
	R_ARM_THM_MOVW_ABS_NC    R_ARM = 47
	R_ARM_THM_MOVT_ABS       R_ARM = 48
	R_ARM_THM_MOVW_PREL_NC   R_ARM = 49
	R_ARM_THM_MOVT_PREL      R_ARM = 50
	R_ARM_THM_JUMP19         R_ARM = 51
	R_ARM_THM_JUMP6          R_ARM = 52
	R_ARM_THM_ALU_PREL_11_0  R_ARM = 53
	R_ARM_THM_PC12           R_ARM = 54
	R_ARM_ABS32_NOI          R_ARM = 55
	R_ARM_REL32_NOI          R_ARM = 56
	R_ARM_ALU_PC_G0_NC       R_ARM = 57
	R_ARM_ALU_PC_G0          R_ARM = 58
	R_ARM_ALU_PC_G1_NC       R_ARM = 59
	R_ARM_ALU_PC_G1          R_ARM = 60
	R_ARM_ALU_PC_G2          R_ARM = 61
	R_ARM_LDR_PC_G1          R_ARM = 62
	R_ARM_LDR_PC_G2          R_ARM = 63
	R_ARM_LDRS_PC_G0         R_ARM = 64
	R_ARM_LDRS_PC_G1         R_ARM = 65
	R_ARM_LDRS_PC_G2         R_ARM = 66
	R_ARM_LDC_PC_G0          R_ARM = 67
	R_ARM_LDC_PC_G1          R_ARM = 68
	R_ARM_LDC_PC_G2          R_ARM = 69
	R_ARM_ALU_SB_G0_NC       R_ARM = 70
	R_ARM_ALU_SB_G0          R_ARM = 71
	R_ARM_ALU_SB_G1_NC       R_ARM = 72
	R_ARM_ALU_SB_G1          R_ARM = 73
	R_ARM_ALU_SB_G2          R_ARM = 74
	R_ARM_LDR_SB_G0          R_ARM = 75
	R_ARM_LDR_SB_G1          R_ARM = 76
	R_ARM_LDR_SB_G2          R_ARM = 77
	R_ARM_LDRS_SB_G0         R_ARM = 78
	R_ARM_LDRS_SB_G1         R_ARM = 79
	R_ARM_LDRS_SB_G2         R_ARM = 80
	R_ARM_LDC_SB_G0          R_ARM = 81
	R_ARM_LDC_SB_G1          R_ARM = 82
	R_ARM_LDC_SB_G2          R_ARM = 83
	R_ARM_MOVW_BREL_NC       R_ARM = 84
	R_ARM_MOVT_BREL          R_ARM = 85
	R_ARM_MOVW_BREL          R_ARM = 86
	R_ARM_THM_MOVW_BREL_NC   R_ARM = 87
	R_ARM_THM_MOVT_BREL      R_ARM = 88
	R_ARM_THM_MOVW_BREL      R_ARM = 89
	R_ARM_TLS_GOTDESC        R_ARM = 90
	R_ARM_TLS_CALL           R_ARM = 91
	R_ARM_TLS_DESCSEQ        R_ARM = 92
	R_ARM_THM_TLS_CALL       R_ARM = 93
	R_ARM_PLT32_ABS          R_ARM = 94
	R_ARM_GOT_ABS            R_ARM = 95
	R_ARM_GOT_PREL           R_ARM = 96
	R_ARM_GOT_BREL12         R_ARM = 97
	R_ARM_GOTOFF12           R_ARM = 98
	R_ARM_GOTRELAX           R_ARM = 99
	R_ARM_GNU_VTENTRY        R_ARM = 100
	R_ARM_GNU_VTINHERIT      R_ARM = 101
	R_ARM_THM_JUMP11         R_ARM = 102
	R_ARM_THM_JUMP8          R_ARM = 103
	R_ARM_TLS_GD32           R_ARM = 104
	R_ARM_TLS_LDM32          R_ARM = 105
	R_ARM_TLS_LDO32          R_ARM = 106
	R_ARM_TLS_IE32           R_ARM = 107
	R_ARM_TLS_LE32           R_ARM = 108
	R_ARM_TLS_LDO12          R_ARM = 109
	R_ARM_TLS_LE12           R_ARM = 110
	R_ARM_TLS_IE12GP         R_ARM = 111
	R_ARM_PRIVATE_0          R_ARM = 112
	R_ARM_PRIVATE_1          R_ARM = 113
	R_ARM_PRIVATE_2          R_ARM = 114
	R_ARM_PRIVATE_3          R_ARM = 115
	R_ARM_PRIVATE_4          R_ARM = 116
	R_ARM_PRIVATE_5          R_ARM = 117
	R_ARM_PRIVATE_6          R_ARM = 118
	R_ARM_PRIVATE_7          R_ARM = 119
	R_ARM_PRIVATE_8          R_ARM = 120
	R_ARM_PRIVATE_9          R_ARM = 121
	R_ARM_PRIVATE_10         R_ARM = 122
	R_ARM_PRIVATE_11         R_ARM = 123
	R_ARM_PRIVATE_12         R_ARM = 124
	R_ARM_PRIVATE_13         R_ARM = 125
	R_ARM_PRIVATE_14         R_ARM = 126
	R_ARM_PRIVATE_15         R_ARM = 127
	R_ARM_ME_TOO             R_ARM = 128
	R_ARM_THM_TLS_DESCSEQ16  R_ARM = 129
	R_ARM_THM_TLS_DESCSEQ32  R_ARM = 130
	R_ARM_THM_GOT_BREL12     R_ARM = 131
	R_ARM_THM_ALU_ABS_G0_NC  R_ARM = 132
	R_ARM_THM_ALU_ABS_G1_NC  R_ARM = 133
	R_ARM_THM_ALU_ABS_G2_NC  R_ARM = 134
	R_ARM_THM_ALU_ABS_G3     R_ARM = 135
	R_ARM_IRELATIVE          R_ARM = 160
	R_ARM_RXPC25             R_ARM = 249
	R_ARM_RSBREL32           R_ARM = 250
	R_ARM_THM_RPC22          R_ARM = 251
	R_ARM_RREL32             R_ARM = 252
	R_ARM_RABS32             R_ARM = 253
	R_ARM_RPC24              R_ARM = 254
	R_ARM_RBASE              R_ARM = 255
)
```

#### (R_ARM) GoString 

``` go 
func (i R_ARM) GoString() string
```

#### (R_ARM) String 

``` go 
func (i R_ARM) String() string
```

### type R_LARCH  <- go1.19

``` go 
type R_LARCH int
```

Relocation types for LoongArch.

``` go 
const (
	R_LARCH_NONE                       R_LARCH = 0
	R_LARCH_32                         R_LARCH = 1
	R_LARCH_64                         R_LARCH = 2
	R_LARCH_RELATIVE                   R_LARCH = 3
	R_LARCH_COPY                       R_LARCH = 4
	R_LARCH_JUMP_SLOT                  R_LARCH = 5
	R_LARCH_TLS_DTPMOD32               R_LARCH = 6
	R_LARCH_TLS_DTPMOD64               R_LARCH = 7
	R_LARCH_TLS_DTPREL32               R_LARCH = 8
	R_LARCH_TLS_DTPREL64               R_LARCH = 9
	R_LARCH_TLS_TPREL32                R_LARCH = 10
	R_LARCH_TLS_TPREL64                R_LARCH = 11
	R_LARCH_IRELATIVE                  R_LARCH = 12
	R_LARCH_MARK_LA                    R_LARCH = 20
	R_LARCH_MARK_PCREL                 R_LARCH = 21
	R_LARCH_SOP_PUSH_PCREL             R_LARCH = 22
	R_LARCH_SOP_PUSH_ABSOLUTE          R_LARCH = 23
	R_LARCH_SOP_PUSH_DUP               R_LARCH = 24
	R_LARCH_SOP_PUSH_GPREL             R_LARCH = 25
	R_LARCH_SOP_PUSH_TLS_TPREL         R_LARCH = 26
	R_LARCH_SOP_PUSH_TLS_GOT           R_LARCH = 27
	R_LARCH_SOP_PUSH_TLS_GD            R_LARCH = 28
	R_LARCH_SOP_PUSH_PLT_PCREL         R_LARCH = 29
	R_LARCH_SOP_ASSERT                 R_LARCH = 30
	R_LARCH_SOP_NOT                    R_LARCH = 31
	R_LARCH_SOP_SUB                    R_LARCH = 32
	R_LARCH_SOP_SL                     R_LARCH = 33
	R_LARCH_SOP_SR                     R_LARCH = 34
	R_LARCH_SOP_ADD                    R_LARCH = 35
	R_LARCH_SOP_AND                    R_LARCH = 36
	R_LARCH_SOP_IF_ELSE                R_LARCH = 37
	R_LARCH_SOP_POP_32_S_10_5          R_LARCH = 38
	R_LARCH_SOP_POP_32_U_10_12         R_LARCH = 39
	R_LARCH_SOP_POP_32_S_10_12         R_LARCH = 40
	R_LARCH_SOP_POP_32_S_10_16         R_LARCH = 41
	R_LARCH_SOP_POP_32_S_10_16_S2      R_LARCH = 42
	R_LARCH_SOP_POP_32_S_5_20          R_LARCH = 43
	R_LARCH_SOP_POP_32_S_0_5_10_16_S2  R_LARCH = 44
	R_LARCH_SOP_POP_32_S_0_10_10_16_S2 R_LARCH = 45
	R_LARCH_SOP_POP_32_U               R_LARCH = 46
	R_LARCH_ADD8                       R_LARCH = 47
	R_LARCH_ADD16                      R_LARCH = 48
	R_LARCH_ADD24                      R_LARCH = 49
	R_LARCH_ADD32                      R_LARCH = 50
	R_LARCH_ADD64                      R_LARCH = 51
	R_LARCH_SUB8                       R_LARCH = 52
	R_LARCH_SUB16                      R_LARCH = 53
	R_LARCH_SUB24                      R_LARCH = 54
	R_LARCH_SUB32                      R_LARCH = 55
	R_LARCH_SUB64                      R_LARCH = 56
	R_LARCH_GNU_VTINHERIT              R_LARCH = 57
	R_LARCH_GNU_VTENTRY                R_LARCH = 58
	R_LARCH_B16                        R_LARCH = 64
	R_LARCH_B21                        R_LARCH = 65
	R_LARCH_B26                        R_LARCH = 66
	R_LARCH_ABS_HI20                   R_LARCH = 67
	R_LARCH_ABS_LO12                   R_LARCH = 68
	R_LARCH_ABS64_LO20                 R_LARCH = 69
	R_LARCH_ABS64_HI12                 R_LARCH = 70
	R_LARCH_PCALA_HI20                 R_LARCH = 71
	R_LARCH_PCALA_LO12                 R_LARCH = 72
	R_LARCH_PCALA64_LO20               R_LARCH = 73
	R_LARCH_PCALA64_HI12               R_LARCH = 74
	R_LARCH_GOT_PC_HI20                R_LARCH = 75
	R_LARCH_GOT_PC_LO12                R_LARCH = 76
	R_LARCH_GOT64_PC_LO20              R_LARCH = 77
	R_LARCH_GOT64_PC_HI12              R_LARCH = 78
	R_LARCH_GOT_HI20                   R_LARCH = 79
	R_LARCH_GOT_LO12                   R_LARCH = 80
	R_LARCH_GOT64_LO20                 R_LARCH = 81
	R_LARCH_GOT64_HI12                 R_LARCH = 82
	R_LARCH_TLS_LE_HI20                R_LARCH = 83
	R_LARCH_TLS_LE_LO12                R_LARCH = 84
	R_LARCH_TLS_LE64_LO20              R_LARCH = 85
	R_LARCH_TLS_LE64_HI12              R_LARCH = 86
	R_LARCH_TLS_IE_PC_HI20             R_LARCH = 87
	R_LARCH_TLS_IE_PC_LO12             R_LARCH = 88
	R_LARCH_TLS_IE64_PC_LO20           R_LARCH = 89
	R_LARCH_TLS_IE64_PC_HI12           R_LARCH = 90
	R_LARCH_TLS_IE_HI20                R_LARCH = 91
	R_LARCH_TLS_IE_LO12                R_LARCH = 92
	R_LARCH_TLS_IE64_LO20              R_LARCH = 93
	R_LARCH_TLS_IE64_HI12              R_LARCH = 94
	R_LARCH_TLS_LD_PC_HI20             R_LARCH = 95
	R_LARCH_TLS_LD_HI20                R_LARCH = 96
	R_LARCH_TLS_GD_PC_HI20             R_LARCH = 97
	R_LARCH_TLS_GD_HI20                R_LARCH = 98
	R_LARCH_32_PCREL                   R_LARCH = 99
	R_LARCH_RELAX                      R_LARCH = 100
)
```

#### (R_LARCH) GoString  <- go1.19

``` go 
func (i R_LARCH) GoString() string
```

#### (R_LARCH) String  <- go1.19

``` go 
func (i R_LARCH) String() string
```

### type R_MIPS  <- go1.6

``` go 
type R_MIPS int
```

Relocation types for MIPS.

``` go 
const (
	R_MIPS_NONE          R_MIPS = 0
	R_MIPS_16            R_MIPS = 1
	R_MIPS_32            R_MIPS = 2
	R_MIPS_REL32         R_MIPS = 3
	R_MIPS_26            R_MIPS = 4
	R_MIPS_HI16          R_MIPS = 5  /* high 16 bits of symbol value */
	R_MIPS_LO16          R_MIPS = 6  /* low 16 bits of symbol value */
	R_MIPS_GPREL16       R_MIPS = 7  /* GP-relative reference  */
	R_MIPS_LITERAL       R_MIPS = 8  /* Reference to literal section  */
	R_MIPS_GOT16         R_MIPS = 9  /* Reference to global offset table */
	R_MIPS_PC16          R_MIPS = 10 /* 16 bit PC relative reference */
	R_MIPS_CALL16        R_MIPS = 11 /* 16 bit call through glbl offset tbl */
	R_MIPS_GPREL32       R_MIPS = 12
	R_MIPS_SHIFT5        R_MIPS = 16
	R_MIPS_SHIFT6        R_MIPS = 17
	R_MIPS_64            R_MIPS = 18
	R_MIPS_GOT_DISP      R_MIPS = 19
	R_MIPS_GOT_PAGE      R_MIPS = 20
	R_MIPS_GOT_OFST      R_MIPS = 21
	R_MIPS_GOT_HI16      R_MIPS = 22
	R_MIPS_GOT_LO16      R_MIPS = 23
	R_MIPS_SUB           R_MIPS = 24
	R_MIPS_INSERT_A      R_MIPS = 25
	R_MIPS_INSERT_B      R_MIPS = 26
	R_MIPS_DELETE        R_MIPS = 27
	R_MIPS_HIGHER        R_MIPS = 28
	R_MIPS_HIGHEST       R_MIPS = 29
	R_MIPS_CALL_HI16     R_MIPS = 30
	R_MIPS_CALL_LO16     R_MIPS = 31
	R_MIPS_SCN_DISP      R_MIPS = 32
	R_MIPS_REL16         R_MIPS = 33
	R_MIPS_ADD_IMMEDIATE R_MIPS = 34
	R_MIPS_PJUMP         R_MIPS = 35
	R_MIPS_RELGOT        R_MIPS = 36
	R_MIPS_JALR          R_MIPS = 37

	R_MIPS_TLS_DTPMOD32    R_MIPS = 38 /* Module number 32 bit */
	R_MIPS_TLS_DTPREL32    R_MIPS = 39 /* Module-relative offset 32 bit */
	R_MIPS_TLS_DTPMOD64    R_MIPS = 40 /* Module number 64 bit */
	R_MIPS_TLS_DTPREL64    R_MIPS = 41 /* Module-relative offset 64 bit */
	R_MIPS_TLS_GD          R_MIPS = 42 /* 16 bit GOT offset for GD */
	R_MIPS_TLS_LDM         R_MIPS = 43 /* 16 bit GOT offset for LDM */
	R_MIPS_TLS_DTPREL_HI16 R_MIPS = 44 /* Module-relative offset, high 16 bits */
	R_MIPS_TLS_DTPREL_LO16 R_MIPS = 45 /* Module-relative offset, low 16 bits */
	R_MIPS_TLS_GOTTPREL    R_MIPS = 46 /* 16 bit GOT offset for IE */
	R_MIPS_TLS_TPREL32     R_MIPS = 47 /* TP-relative offset, 32 bit */
	R_MIPS_TLS_TPREL64     R_MIPS = 48 /* TP-relative offset, 64 bit */
	R_MIPS_TLS_TPREL_HI16  R_MIPS = 49 /* TP-relative offset, high 16 bits */
	R_MIPS_TLS_TPREL_LO16  R_MIPS = 50 /* TP-relative offset, low 16 bits */
)
```

#### (R_MIPS) GoString  <- go1.6

``` go 
func (i R_MIPS) GoString() string
```

#### (R_MIPS) String  <- go1.6

``` go 
func (i R_MIPS) String() string
```

### type R_PPC 

``` go 
type R_PPC int
```

Relocation types for PowerPC.

Values that are shared by both R_PPC and R_PPC64 are prefixed with R_POWERPC_ in the ELF standard. For the R_PPC type, the relevant shared relocations have been renamed with the prefix R_PPC_. The original name follows the value in a comment.

``` go 
const (
	R_PPC_NONE            R_PPC = 0  // R_POWERPC_NONE
	R_PPC_ADDR32          R_PPC = 1  // R_POWERPC_ADDR32
	R_PPC_ADDR24          R_PPC = 2  // R_POWERPC_ADDR24
	R_PPC_ADDR16          R_PPC = 3  // R_POWERPC_ADDR16
	R_PPC_ADDR16_LO       R_PPC = 4  // R_POWERPC_ADDR16_LO
	R_PPC_ADDR16_HI       R_PPC = 5  // R_POWERPC_ADDR16_HI
	R_PPC_ADDR16_HA       R_PPC = 6  // R_POWERPC_ADDR16_HA
	R_PPC_ADDR14          R_PPC = 7  // R_POWERPC_ADDR14
	R_PPC_ADDR14_BRTAKEN  R_PPC = 8  // R_POWERPC_ADDR14_BRTAKEN
	R_PPC_ADDR14_BRNTAKEN R_PPC = 9  // R_POWERPC_ADDR14_BRNTAKEN
	R_PPC_REL24           R_PPC = 10 // R_POWERPC_REL24
	R_PPC_REL14           R_PPC = 11 // R_POWERPC_REL14
	R_PPC_REL14_BRTAKEN   R_PPC = 12 // R_POWERPC_REL14_BRTAKEN
	R_PPC_REL14_BRNTAKEN  R_PPC = 13 // R_POWERPC_REL14_BRNTAKEN
	R_PPC_GOT16           R_PPC = 14 // R_POWERPC_GOT16
	R_PPC_GOT16_LO        R_PPC = 15 // R_POWERPC_GOT16_LO
	R_PPC_GOT16_HI        R_PPC = 16 // R_POWERPC_GOT16_HI
	R_PPC_GOT16_HA        R_PPC = 17 // R_POWERPC_GOT16_HA
	R_PPC_PLTREL24        R_PPC = 18
	R_PPC_COPY            R_PPC = 19 // R_POWERPC_COPY
	R_PPC_GLOB_DAT        R_PPC = 20 // R_POWERPC_GLOB_DAT
	R_PPC_JMP_SLOT        R_PPC = 21 // R_POWERPC_JMP_SLOT
	R_PPC_RELATIVE        R_PPC = 22 // R_POWERPC_RELATIVE
	R_PPC_LOCAL24PC       R_PPC = 23
	R_PPC_UADDR32         R_PPC = 24 // R_POWERPC_UADDR32
	R_PPC_UADDR16         R_PPC = 25 // R_POWERPC_UADDR16
	R_PPC_REL32           R_PPC = 26 // R_POWERPC_REL32
	R_PPC_PLT32           R_PPC = 27 // R_POWERPC_PLT32
	R_PPC_PLTREL32        R_PPC = 28 // R_POWERPC_PLTREL32
	R_PPC_PLT16_LO        R_PPC = 29 // R_POWERPC_PLT16_LO
	R_PPC_PLT16_HI        R_PPC = 30 // R_POWERPC_PLT16_HI
	R_PPC_PLT16_HA        R_PPC = 31 // R_POWERPC_PLT16_HA
	R_PPC_SDAREL16        R_PPC = 32
	R_PPC_SECTOFF         R_PPC = 33 // R_POWERPC_SECTOFF
	R_PPC_SECTOFF_LO      R_PPC = 34 // R_POWERPC_SECTOFF_LO
	R_PPC_SECTOFF_HI      R_PPC = 35 // R_POWERPC_SECTOFF_HI
	R_PPC_SECTOFF_HA      R_PPC = 36 // R_POWERPC_SECTOFF_HA
	R_PPC_TLS             R_PPC = 67 // R_POWERPC_TLS
	R_PPC_DTPMOD32        R_PPC = 68 // R_POWERPC_DTPMOD32
	R_PPC_TPREL16         R_PPC = 69 // R_POWERPC_TPREL16
	R_PPC_TPREL16_LO      R_PPC = 70 // R_POWERPC_TPREL16_LO
	R_PPC_TPREL16_HI      R_PPC = 71 // R_POWERPC_TPREL16_HI
	R_PPC_TPREL16_HA      R_PPC = 72 // R_POWERPC_TPREL16_HA
	R_PPC_TPREL32         R_PPC = 73 // R_POWERPC_TPREL32
	R_PPC_DTPREL16        R_PPC = 74 // R_POWERPC_DTPREL16
	R_PPC_DTPREL16_LO     R_PPC = 75 // R_POWERPC_DTPREL16_LO
	R_PPC_DTPREL16_HI     R_PPC = 76 // R_POWERPC_DTPREL16_HI
	R_PPC_DTPREL16_HA     R_PPC = 77 // R_POWERPC_DTPREL16_HA
	R_PPC_DTPREL32        R_PPC = 78 // R_POWERPC_DTPREL32
	R_PPC_GOT_TLSGD16     R_PPC = 79 // R_POWERPC_GOT_TLSGD16
	R_PPC_GOT_TLSGD16_LO  R_PPC = 80 // R_POWERPC_GOT_TLSGD16_LO
	R_PPC_GOT_TLSGD16_HI  R_PPC = 81 // R_POWERPC_GOT_TLSGD16_HI
	R_PPC_GOT_TLSGD16_HA  R_PPC = 82 // R_POWERPC_GOT_TLSGD16_HA
	R_PPC_GOT_TLSLD16     R_PPC = 83 // R_POWERPC_GOT_TLSLD16
	R_PPC_GOT_TLSLD16_LO  R_PPC = 84 // R_POWERPC_GOT_TLSLD16_LO
	R_PPC_GOT_TLSLD16_HI  R_PPC = 85 // R_POWERPC_GOT_TLSLD16_HI
	R_PPC_GOT_TLSLD16_HA  R_PPC = 86 // R_POWERPC_GOT_TLSLD16_HA
	R_PPC_GOT_TPREL16     R_PPC = 87 // R_POWERPC_GOT_TPREL16
	R_PPC_GOT_TPREL16_LO  R_PPC = 88 // R_POWERPC_GOT_TPREL16_LO
	R_PPC_GOT_TPREL16_HI  R_PPC = 89 // R_POWERPC_GOT_TPREL16_HI
	R_PPC_GOT_TPREL16_HA  R_PPC = 90 // R_POWERPC_GOT_TPREL16_HA
	R_PPC_EMB_NADDR32     R_PPC = 101
	R_PPC_EMB_NADDR16     R_PPC = 102
	R_PPC_EMB_NADDR16_LO  R_PPC = 103
	R_PPC_EMB_NADDR16_HI  R_PPC = 104
	R_PPC_EMB_NADDR16_HA  R_PPC = 105
	R_PPC_EMB_SDAI16      R_PPC = 106
	R_PPC_EMB_SDA2I16     R_PPC = 107
	R_PPC_EMB_SDA2REL     R_PPC = 108
	R_PPC_EMB_SDA21       R_PPC = 109
	R_PPC_EMB_MRKREF      R_PPC = 110
	R_PPC_EMB_RELSEC16    R_PPC = 111
	R_PPC_EMB_RELST_LO    R_PPC = 112
	R_PPC_EMB_RELST_HI    R_PPC = 113
	R_PPC_EMB_RELST_HA    R_PPC = 114
	R_PPC_EMB_BIT_FLD     R_PPC = 115
	R_PPC_EMB_RELSDA      R_PPC = 116
)
```

#### (R_PPC) GoString 

``` go 
func (i R_PPC) GoString() string
```

#### (R_PPC) String 

``` go 
func (i R_PPC) String() string
```

### type R_PPC64  <- go1.5

``` go 
type R_PPC64 int
```

Relocation types for 64-bit PowerPC or Power Architecture processors.

Values that are shared by both R_PPC and R_PPC64 are prefixed with R_POWERPC_ in the ELF standard. For the R_PPC64 type, the relevant shared relocations have been renamed with the prefix R_PPC64_. The original name follows the value in a comment.

``` go 
const (
	R_PPC64_NONE               R_PPC64 = 0  // R_POWERPC_NONE
	R_PPC64_ADDR32             R_PPC64 = 1  // R_POWERPC_ADDR32
	R_PPC64_ADDR24             R_PPC64 = 2  // R_POWERPC_ADDR24
	R_PPC64_ADDR16             R_PPC64 = 3  // R_POWERPC_ADDR16
	R_PPC64_ADDR16_LO          R_PPC64 = 4  // R_POWERPC_ADDR16_LO
	R_PPC64_ADDR16_HI          R_PPC64 = 5  // R_POWERPC_ADDR16_HI
	R_PPC64_ADDR16_HA          R_PPC64 = 6  // R_POWERPC_ADDR16_HA
	R_PPC64_ADDR14             R_PPC64 = 7  // R_POWERPC_ADDR14
	R_PPC64_ADDR14_BRTAKEN     R_PPC64 = 8  // R_POWERPC_ADDR14_BRTAKEN
	R_PPC64_ADDR14_BRNTAKEN    R_PPC64 = 9  // R_POWERPC_ADDR14_BRNTAKEN
	R_PPC64_REL24              R_PPC64 = 10 // R_POWERPC_REL24
	R_PPC64_REL14              R_PPC64 = 11 // R_POWERPC_REL14
	R_PPC64_REL14_BRTAKEN      R_PPC64 = 12 // R_POWERPC_REL14_BRTAKEN
	R_PPC64_REL14_BRNTAKEN     R_PPC64 = 13 // R_POWERPC_REL14_BRNTAKEN
	R_PPC64_GOT16              R_PPC64 = 14 // R_POWERPC_GOT16
	R_PPC64_GOT16_LO           R_PPC64 = 15 // R_POWERPC_GOT16_LO
	R_PPC64_GOT16_HI           R_PPC64 = 16 // R_POWERPC_GOT16_HI
	R_PPC64_GOT16_HA           R_PPC64 = 17 // R_POWERPC_GOT16_HA
	R_PPC64_COPY               R_PPC64 = 19 // R_POWERPC_COPY
	R_PPC64_GLOB_DAT           R_PPC64 = 20 // R_POWERPC_GLOB_DAT
	R_PPC64_JMP_SLOT           R_PPC64 = 21 // R_POWERPC_JMP_SLOT
	R_PPC64_RELATIVE           R_PPC64 = 22 // R_POWERPC_RELATIVE
	R_PPC64_UADDR32            R_PPC64 = 24 // R_POWERPC_UADDR32
	R_PPC64_UADDR16            R_PPC64 = 25 // R_POWERPC_UADDR16
	R_PPC64_REL32              R_PPC64 = 26 // R_POWERPC_REL32
	R_PPC64_PLT32              R_PPC64 = 27 // R_POWERPC_PLT32
	R_PPC64_PLTREL32           R_PPC64 = 28 // R_POWERPC_PLTREL32
	R_PPC64_PLT16_LO           R_PPC64 = 29 // R_POWERPC_PLT16_LO
	R_PPC64_PLT16_HI           R_PPC64 = 30 // R_POWERPC_PLT16_HI
	R_PPC64_PLT16_HA           R_PPC64 = 31 // R_POWERPC_PLT16_HA
	R_PPC64_SECTOFF            R_PPC64 = 33 // R_POWERPC_SECTOFF
	R_PPC64_SECTOFF_LO         R_PPC64 = 34 // R_POWERPC_SECTOFF_LO
	R_PPC64_SECTOFF_HI         R_PPC64 = 35 // R_POWERPC_SECTOFF_HI
	R_PPC64_SECTOFF_HA         R_PPC64 = 36 // R_POWERPC_SECTOFF_HA
	R_PPC64_REL30              R_PPC64 = 37 // R_POWERPC_ADDR30
	R_PPC64_ADDR64             R_PPC64 = 38
	R_PPC64_ADDR16_HIGHER      R_PPC64 = 39
	R_PPC64_ADDR16_HIGHERA     R_PPC64 = 40
	R_PPC64_ADDR16_HIGHEST     R_PPC64 = 41
	R_PPC64_ADDR16_HIGHESTA    R_PPC64 = 42
	R_PPC64_UADDR64            R_PPC64 = 43
	R_PPC64_REL64              R_PPC64 = 44
	R_PPC64_PLT64              R_PPC64 = 45
	R_PPC64_PLTREL64           R_PPC64 = 46
	R_PPC64_TOC16              R_PPC64 = 47
	R_PPC64_TOC16_LO           R_PPC64 = 48
	R_PPC64_TOC16_HI           R_PPC64 = 49
	R_PPC64_TOC16_HA           R_PPC64 = 50
	R_PPC64_TOC                R_PPC64 = 51
	R_PPC64_PLTGOT16           R_PPC64 = 52
	R_PPC64_PLTGOT16_LO        R_PPC64 = 53
	R_PPC64_PLTGOT16_HI        R_PPC64 = 54
	R_PPC64_PLTGOT16_HA        R_PPC64 = 55
	R_PPC64_ADDR16_DS          R_PPC64 = 56
	R_PPC64_ADDR16_LO_DS       R_PPC64 = 57
	R_PPC64_GOT16_DS           R_PPC64 = 58
	R_PPC64_GOT16_LO_DS        R_PPC64 = 59
	R_PPC64_PLT16_LO_DS        R_PPC64 = 60
	R_PPC64_SECTOFF_DS         R_PPC64 = 61
	R_PPC64_SECTOFF_LO_DS      R_PPC64 = 62
	R_PPC64_TOC16_DS           R_PPC64 = 63
	R_PPC64_TOC16_LO_DS        R_PPC64 = 64
	R_PPC64_PLTGOT16_DS        R_PPC64 = 65
	R_PPC64_PLTGOT_LO_DS       R_PPC64 = 66
	R_PPC64_TLS                R_PPC64 = 67 // R_POWERPC_TLS
	R_PPC64_DTPMOD64           R_PPC64 = 68 // R_POWERPC_DTPMOD64
	R_PPC64_TPREL16            R_PPC64 = 69 // R_POWERPC_TPREL16
	R_PPC64_TPREL16_LO         R_PPC64 = 70 // R_POWERPC_TPREL16_LO
	R_PPC64_TPREL16_HI         R_PPC64 = 71 // R_POWERPC_TPREL16_HI
	R_PPC64_TPREL16_HA         R_PPC64 = 72 // R_POWERPC_TPREL16_HA
	R_PPC64_TPREL64            R_PPC64 = 73 // R_POWERPC_TPREL64
	R_PPC64_DTPREL16           R_PPC64 = 74 // R_POWERPC_DTPREL16
	R_PPC64_DTPREL16_LO        R_PPC64 = 75 // R_POWERPC_DTPREL16_LO
	R_PPC64_DTPREL16_HI        R_PPC64 = 76 // R_POWERPC_DTPREL16_HI
	R_PPC64_DTPREL16_HA        R_PPC64 = 77 // R_POWERPC_DTPREL16_HA
	R_PPC64_DTPREL64           R_PPC64 = 78 // R_POWERPC_DTPREL64
	R_PPC64_GOT_TLSGD16        R_PPC64 = 79 // R_POWERPC_GOT_TLSGD16
	R_PPC64_GOT_TLSGD16_LO     R_PPC64 = 80 // R_POWERPC_GOT_TLSGD16_LO
	R_PPC64_GOT_TLSGD16_HI     R_PPC64 = 81 // R_POWERPC_GOT_TLSGD16_HI
	R_PPC64_GOT_TLSGD16_HA     R_PPC64 = 82 // R_POWERPC_GOT_TLSGD16_HA
	R_PPC64_GOT_TLSLD16        R_PPC64 = 83 // R_POWERPC_GOT_TLSLD16
	R_PPC64_GOT_TLSLD16_LO     R_PPC64 = 84 // R_POWERPC_GOT_TLSLD16_LO
	R_PPC64_GOT_TLSLD16_HI     R_PPC64 = 85 // R_POWERPC_GOT_TLSLD16_HI
	R_PPC64_GOT_TLSLD16_HA     R_PPC64 = 86 // R_POWERPC_GOT_TLSLD16_HA
	R_PPC64_GOT_TPREL16_DS     R_PPC64 = 87 // R_POWERPC_GOT_TPREL16_DS
	R_PPC64_GOT_TPREL16_LO_DS  R_PPC64 = 88 // R_POWERPC_GOT_TPREL16_LO_DS
	R_PPC64_GOT_TPREL16_HI     R_PPC64 = 89 // R_POWERPC_GOT_TPREL16_HI
	R_PPC64_GOT_TPREL16_HA     R_PPC64 = 90 // R_POWERPC_GOT_TPREL16_HA
	R_PPC64_GOT_DTPREL16_DS    R_PPC64 = 91 // R_POWERPC_GOT_DTPREL16_DS
	R_PPC64_GOT_DTPREL16_LO_DS R_PPC64 = 92 // R_POWERPC_GOT_DTPREL16_LO_DS
	R_PPC64_GOT_DTPREL16_HI    R_PPC64 = 93 // R_POWERPC_GOT_DTPREL16_HI
	R_PPC64_GOT_DTPREL16_HA    R_PPC64 = 94 // R_POWERPC_GOT_DTPREL16_HA
	R_PPC64_TPREL16_DS         R_PPC64 = 95
	R_PPC64_TPREL16_LO_DS      R_PPC64 = 96
	R_PPC64_TPREL16_HIGHER     R_PPC64 = 97
	R_PPC64_TPREL16_HIGHERA    R_PPC64 = 98
	R_PPC64_TPREL16_HIGHEST    R_PPC64 = 99
	R_PPC64_TPREL16_HIGHESTA   R_PPC64 = 100
	R_PPC64_DTPREL16_DS        R_PPC64 = 101
	R_PPC64_DTPREL16_LO_DS     R_PPC64 = 102
	R_PPC64_DTPREL16_HIGHER    R_PPC64 = 103
	R_PPC64_DTPREL16_HIGHERA   R_PPC64 = 104
	R_PPC64_DTPREL16_HIGHEST   R_PPC64 = 105
	R_PPC64_DTPREL16_HIGHESTA  R_PPC64 = 106
	R_PPC64_TLSGD              R_PPC64 = 107
	R_PPC64_TLSLD              R_PPC64 = 108
	R_PPC64_TOCSAVE            R_PPC64 = 109
	R_PPC64_ADDR16_HIGH        R_PPC64 = 110
	R_PPC64_ADDR16_HIGHA       R_PPC64 = 111
	R_PPC64_TPREL16_HIGH       R_PPC64 = 112
	R_PPC64_TPREL16_HIGHA      R_PPC64 = 113
	R_PPC64_DTPREL16_HIGH      R_PPC64 = 114
	R_PPC64_DTPREL16_HIGHA     R_PPC64 = 115
	R_PPC64_REL24_NOTOC        R_PPC64 = 116
	R_PPC64_ADDR64_LOCAL       R_PPC64 = 117
	R_PPC64_ENTRY              R_PPC64 = 118
	R_PPC64_PLTSEQ             R_PPC64 = 119
	R_PPC64_PLTCALL            R_PPC64 = 120
	R_PPC64_PLTSEQ_NOTOC       R_PPC64 = 121
	R_PPC64_PLTCALL_NOTOC      R_PPC64 = 122
	R_PPC64_PCREL_OPT          R_PPC64 = 123
	R_PPC64_D34                R_PPC64 = 128
	R_PPC64_D34_LO             R_PPC64 = 129
	R_PPC64_D34_HI30           R_PPC64 = 130
	R_PPC64_D34_HA30           R_PPC64 = 131
	R_PPC64_PCREL34            R_PPC64 = 132
	R_PPC64_GOT_PCREL34        R_PPC64 = 133
	R_PPC64_PLT_PCREL34        R_PPC64 = 134
	R_PPC64_PLT_PCREL34_NOTOC  R_PPC64 = 135
	R_PPC64_ADDR16_HIGHER34    R_PPC64 = 136
	R_PPC64_ADDR16_HIGHERA34   R_PPC64 = 137
	R_PPC64_ADDR16_HIGHEST34   R_PPC64 = 138
	R_PPC64_ADDR16_HIGHESTA34  R_PPC64 = 139
	R_PPC64_REL16_HIGHER34     R_PPC64 = 140
	R_PPC64_REL16_HIGHERA34    R_PPC64 = 141
	R_PPC64_REL16_HIGHEST34    R_PPC64 = 142
	R_PPC64_REL16_HIGHESTA34   R_PPC64 = 143
	R_PPC64_D28                R_PPC64 = 144
	R_PPC64_PCREL28            R_PPC64 = 145
	R_PPC64_TPREL34            R_PPC64 = 146
	R_PPC64_DTPREL34           R_PPC64 = 147
	R_PPC64_GOT_TLSGD_PCREL34  R_PPC64 = 148
	R_PPC64_GOT_TLSLD_PCREL34  R_PPC64 = 149
	R_PPC64_GOT_TPREL_PCREL34  R_PPC64 = 150
	R_PPC64_GOT_DTPREL_PCREL34 R_PPC64 = 151
	R_PPC64_REL16_HIGH         R_PPC64 = 240
	R_PPC64_REL16_HIGHA        R_PPC64 = 241
	R_PPC64_REL16_HIGHER       R_PPC64 = 242
	R_PPC64_REL16_HIGHERA      R_PPC64 = 243
	R_PPC64_REL16_HIGHEST      R_PPC64 = 244
	R_PPC64_REL16_HIGHESTA     R_PPC64 = 245
	R_PPC64_REL16DX_HA         R_PPC64 = 246 // R_POWERPC_REL16DX_HA
	R_PPC64_JMP_IREL           R_PPC64 = 247
	R_PPC64_IRELATIVE          R_PPC64 = 248 // R_POWERPC_IRELATIVE
	R_PPC64_REL16              R_PPC64 = 249 // R_POWERPC_REL16
	R_PPC64_REL16_LO           R_PPC64 = 250 // R_POWERPC_REL16_LO
	R_PPC64_REL16_HI           R_PPC64 = 251 // R_POWERPC_REL16_HI
	R_PPC64_REL16_HA           R_PPC64 = 252 // R_POWERPC_REL16_HA
	R_PPC64_GNU_VTINHERIT      R_PPC64 = 253
	R_PPC64_GNU_VTENTRY        R_PPC64 = 254
)
```

#### (R_PPC64) GoString  <- go1.5

``` go 
func (i R_PPC64) GoString() string
```

#### (R_PPC64) String  <- go1.5

``` go 
func (i R_PPC64) String() string
```

### type R_RISCV  <- go1.11

``` go 
type R_RISCV int
```

Relocation types for RISC-V processors.

``` go 
const (
	R_RISCV_NONE          R_RISCV = 0  /* No relocation. */
	R_RISCV_32            R_RISCV = 1  /* Add 32 bit zero extended symbol value */
	R_RISCV_64            R_RISCV = 2  /* Add 64 bit symbol value. */
	R_RISCV_RELATIVE      R_RISCV = 3  /* Add load address of shared object. */
	R_RISCV_COPY          R_RISCV = 4  /* Copy data from shared object. */
	R_RISCV_JUMP_SLOT     R_RISCV = 5  /* Set GOT entry to code address. */
	R_RISCV_TLS_DTPMOD32  R_RISCV = 6  /* 32 bit ID of module containing symbol */
	R_RISCV_TLS_DTPMOD64  R_RISCV = 7  /* ID of module containing symbol */
	R_RISCV_TLS_DTPREL32  R_RISCV = 8  /* 32 bit relative offset in TLS block */
	R_RISCV_TLS_DTPREL64  R_RISCV = 9  /* Relative offset in TLS block */
	R_RISCV_TLS_TPREL32   R_RISCV = 10 /* 32 bit relative offset in static TLS block */
	R_RISCV_TLS_TPREL64   R_RISCV = 11 /* Relative offset in static TLS block */
	R_RISCV_BRANCH        R_RISCV = 16 /* PC-relative branch */
	R_RISCV_JAL           R_RISCV = 17 /* PC-relative jump */
	R_RISCV_CALL          R_RISCV = 18 /* PC-relative call */
	R_RISCV_CALL_PLT      R_RISCV = 19 /* PC-relative call (PLT) */
	R_RISCV_GOT_HI20      R_RISCV = 20 /* PC-relative GOT reference */
	R_RISCV_TLS_GOT_HI20  R_RISCV = 21 /* PC-relative TLS IE GOT offset */
	R_RISCV_TLS_GD_HI20   R_RISCV = 22 /* PC-relative TLS GD reference */
	R_RISCV_PCREL_HI20    R_RISCV = 23 /* PC-relative reference */
	R_RISCV_PCREL_LO12_I  R_RISCV = 24 /* PC-relative reference */
	R_RISCV_PCREL_LO12_S  R_RISCV = 25 /* PC-relative reference */
	R_RISCV_HI20          R_RISCV = 26 /* Absolute address */
	R_RISCV_LO12_I        R_RISCV = 27 /* Absolute address */
	R_RISCV_LO12_S        R_RISCV = 28 /* Absolute address */
	R_RISCV_TPREL_HI20    R_RISCV = 29 /* TLS LE thread offset */
	R_RISCV_TPREL_LO12_I  R_RISCV = 30 /* TLS LE thread offset */
	R_RISCV_TPREL_LO12_S  R_RISCV = 31 /* TLS LE thread offset */
	R_RISCV_TPREL_ADD     R_RISCV = 32 /* TLS LE thread usage */
	R_RISCV_ADD8          R_RISCV = 33 /* 8-bit label addition */
	R_RISCV_ADD16         R_RISCV = 34 /* 16-bit label addition */
	R_RISCV_ADD32         R_RISCV = 35 /* 32-bit label addition */
	R_RISCV_ADD64         R_RISCV = 36 /* 64-bit label addition */
	R_RISCV_SUB8          R_RISCV = 37 /* 8-bit label subtraction */
	R_RISCV_SUB16         R_RISCV = 38 /* 16-bit label subtraction */
	R_RISCV_SUB32         R_RISCV = 39 /* 32-bit label subtraction */
	R_RISCV_SUB64         R_RISCV = 40 /* 64-bit label subtraction */
	R_RISCV_GNU_VTINHERIT R_RISCV = 41 /* GNU C++ vtable hierarchy */
	R_RISCV_GNU_VTENTRY   R_RISCV = 42 /* GNU C++ vtable member usage */
	R_RISCV_ALIGN         R_RISCV = 43 /* Alignment statement */
	R_RISCV_RVC_BRANCH    R_RISCV = 44 /* PC-relative branch offset */
	R_RISCV_RVC_JUMP      R_RISCV = 45 /* PC-relative jump offset */
	R_RISCV_RVC_LUI       R_RISCV = 46 /* Absolute address */
	R_RISCV_GPREL_I       R_RISCV = 47 /* GP-relative reference */
	R_RISCV_GPREL_S       R_RISCV = 48 /* GP-relative reference */
	R_RISCV_TPREL_I       R_RISCV = 49 /* TP-relative TLS LE load */
	R_RISCV_TPREL_S       R_RISCV = 50 /* TP-relative TLS LE store */
	R_RISCV_RELAX         R_RISCV = 51 /* Instruction pair can be relaxed */
	R_RISCV_SUB6          R_RISCV = 52 /* Local label subtraction */
	R_RISCV_SET6          R_RISCV = 53 /* Local label subtraction */
	R_RISCV_SET8          R_RISCV = 54 /* Local label subtraction */
	R_RISCV_SET16         R_RISCV = 55 /* Local label subtraction */
	R_RISCV_SET32         R_RISCV = 56 /* Local label subtraction */
	R_RISCV_32_PCREL      R_RISCV = 57 /* 32-bit PC relative */
)
```

#### (R_RISCV) GoString  <- go1.11

``` go 
func (i R_RISCV) GoString() string
```

#### (R_RISCV) String  <- go1.11

``` go 
func (i R_RISCV) String() string
```

### type R_SPARC 

``` go 
type R_SPARC int
```

Relocation types for SPARC.

``` go 
const (
	R_SPARC_NONE     R_SPARC = 0
	R_SPARC_8        R_SPARC = 1
	R_SPARC_16       R_SPARC = 2
	R_SPARC_32       R_SPARC = 3
	R_SPARC_DISP8    R_SPARC = 4
	R_SPARC_DISP16   R_SPARC = 5
	R_SPARC_DISP32   R_SPARC = 6
	R_SPARC_WDISP30  R_SPARC = 7
	R_SPARC_WDISP22  R_SPARC = 8
	R_SPARC_HI22     R_SPARC = 9
	R_SPARC_22       R_SPARC = 10
	R_SPARC_13       R_SPARC = 11
	R_SPARC_LO10     R_SPARC = 12
	R_SPARC_GOT10    R_SPARC = 13
	R_SPARC_GOT13    R_SPARC = 14
	R_SPARC_GOT22    R_SPARC = 15
	R_SPARC_PC10     R_SPARC = 16
	R_SPARC_PC22     R_SPARC = 17
	R_SPARC_WPLT30   R_SPARC = 18
	R_SPARC_COPY     R_SPARC = 19
	R_SPARC_GLOB_DAT R_SPARC = 20
	R_SPARC_JMP_SLOT R_SPARC = 21
	R_SPARC_RELATIVE R_SPARC = 22
	R_SPARC_UA32     R_SPARC = 23
	R_SPARC_PLT32    R_SPARC = 24
	R_SPARC_HIPLT22  R_SPARC = 25
	R_SPARC_LOPLT10  R_SPARC = 26
	R_SPARC_PCPLT32  R_SPARC = 27
	R_SPARC_PCPLT22  R_SPARC = 28
	R_SPARC_PCPLT10  R_SPARC = 29
	R_SPARC_10       R_SPARC = 30
	R_SPARC_11       R_SPARC = 31
	R_SPARC_64       R_SPARC = 32
	R_SPARC_OLO10    R_SPARC = 33
	R_SPARC_HH22     R_SPARC = 34
	R_SPARC_HM10     R_SPARC = 35
	R_SPARC_LM22     R_SPARC = 36
	R_SPARC_PC_HH22  R_SPARC = 37
	R_SPARC_PC_HM10  R_SPARC = 38
	R_SPARC_PC_LM22  R_SPARC = 39
	R_SPARC_WDISP16  R_SPARC = 40
	R_SPARC_WDISP19  R_SPARC = 41
	R_SPARC_GLOB_JMP R_SPARC = 42
	R_SPARC_7        R_SPARC = 43
	R_SPARC_5        R_SPARC = 44
	R_SPARC_6        R_SPARC = 45
	R_SPARC_DISP64   R_SPARC = 46
	R_SPARC_PLT64    R_SPARC = 47
	R_SPARC_HIX22    R_SPARC = 48
	R_SPARC_LOX10    R_SPARC = 49
	R_SPARC_H44      R_SPARC = 50
	R_SPARC_M44      R_SPARC = 51
	R_SPARC_L44      R_SPARC = 52
	R_SPARC_REGISTER R_SPARC = 53
	R_SPARC_UA64     R_SPARC = 54
	R_SPARC_UA16     R_SPARC = 55
)
```

#### (R_SPARC) GoString 

``` go 
func (i R_SPARC) GoString() string
```

#### (R_SPARC) String 

``` go 
func (i R_SPARC) String() string
```

### type R_X86_64 

``` go 
type R_X86_64 int
```

Relocation types for x86-64.

``` go 
const (
	R_X86_64_NONE            R_X86_64 = 0  /* No relocation. */
	R_X86_64_64              R_X86_64 = 1  /* Add 64 bit symbol value. */
	R_X86_64_PC32            R_X86_64 = 2  /* PC-relative 32 bit signed sym value. */
	R_X86_64_GOT32           R_X86_64 = 3  /* PC-relative 32 bit GOT offset. */
	R_X86_64_PLT32           R_X86_64 = 4  /* PC-relative 32 bit PLT offset. */
	R_X86_64_COPY            R_X86_64 = 5  /* Copy data from shared object. */
	R_X86_64_GLOB_DAT        R_X86_64 = 6  /* Set GOT entry to data address. */
	R_X86_64_JMP_SLOT        R_X86_64 = 7  /* Set GOT entry to code address. */
	R_X86_64_RELATIVE        R_X86_64 = 8  /* Add load address of shared object. */
	R_X86_64_GOTPCREL        R_X86_64 = 9  /* Add 32 bit signed pcrel offset to GOT. */
	R_X86_64_32              R_X86_64 = 10 /* Add 32 bit zero extended symbol value */
	R_X86_64_32S             R_X86_64 = 11 /* Add 32 bit sign extended symbol value */
	R_X86_64_16              R_X86_64 = 12 /* Add 16 bit zero extended symbol value */
	R_X86_64_PC16            R_X86_64 = 13 /* Add 16 bit signed extended pc relative symbol value */
	R_X86_64_8               R_X86_64 = 14 /* Add 8 bit zero extended symbol value */
	R_X86_64_PC8             R_X86_64 = 15 /* Add 8 bit signed extended pc relative symbol value */
	R_X86_64_DTPMOD64        R_X86_64 = 16 /* ID of module containing symbol */
	R_X86_64_DTPOFF64        R_X86_64 = 17 /* Offset in TLS block */
	R_X86_64_TPOFF64         R_X86_64 = 18 /* Offset in static TLS block */
	R_X86_64_TLSGD           R_X86_64 = 19 /* PC relative offset to GD GOT entry */
	R_X86_64_TLSLD           R_X86_64 = 20 /* PC relative offset to LD GOT entry */
	R_X86_64_DTPOFF32        R_X86_64 = 21 /* Offset in TLS block */
	R_X86_64_GOTTPOFF        R_X86_64 = 22 /* PC relative offset to IE GOT entry */
	R_X86_64_TPOFF32         R_X86_64 = 23 /* Offset in static TLS block */
	R_X86_64_PC64            R_X86_64 = 24 /* PC relative 64-bit sign extended symbol value. */
	R_X86_64_GOTOFF64        R_X86_64 = 25
	R_X86_64_GOTPC32         R_X86_64 = 26
	R_X86_64_GOT64           R_X86_64 = 27
	R_X86_64_GOTPCREL64      R_X86_64 = 28
	R_X86_64_GOTPC64         R_X86_64 = 29
	R_X86_64_GOTPLT64        R_X86_64 = 30
	R_X86_64_PLTOFF64        R_X86_64 = 31
	R_X86_64_SIZE32          R_X86_64 = 32
	R_X86_64_SIZE64          R_X86_64 = 33
	R_X86_64_GOTPC32_TLSDESC R_X86_64 = 34
	R_X86_64_TLSDESC_CALL    R_X86_64 = 35
	R_X86_64_TLSDESC         R_X86_64 = 36
	R_X86_64_IRELATIVE       R_X86_64 = 37
	R_X86_64_RELATIVE64      R_X86_64 = 38
	R_X86_64_PC32_BND        R_X86_64 = 39
	R_X86_64_PLT32_BND       R_X86_64 = 40
	R_X86_64_GOTPCRELX       R_X86_64 = 41
	R_X86_64_REX_GOTPCRELX   R_X86_64 = 42
)
```

#### (R_X86_64) GoString 

``` go 
func (i R_X86_64) GoString() string
```

#### (R_X86_64) String 

``` go 
func (i R_X86_64) String() string
```

### type Rel32 

``` go 
type Rel32 struct {
	Off  uint32 /* Location to be relocated. */
	Info uint32 /* Relocation type and symbol index. */
}
```

ELF32 Relocations that don't need an addend field.

### type Rel64 

``` go 
type Rel64 struct {
	Off  uint64 /* Location to be relocated. */
	Info uint64 /* Relocation type and symbol index. */
}
```

ELF64 relocations that don't need an addend field.

### type Rela32 

``` go 
type Rela32 struct {
	Off    uint32 /* Location to be relocated. */
	Info   uint32 /* Relocation type and symbol index. */
	Addend int32  /* Addend. */
}
```

ELF32 Relocations that need an addend field.

### type Rela64 

``` go 
type Rela64 struct {
	Off    uint64 /* Location to be relocated. */
	Info   uint64 /* Relocation type and symbol index. */
	Addend int64  /* Addend. */
}
```

ELF64 relocations that need an addend field.

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
	//
	// ReaderAt may be nil if the section is not easily available
	// in a random-access form. For example, a compressed section
	// may have a nil ReaderAt.
	io.ReaderAt
	// contains filtered or unexported fields
}
```

A Section represents a single section in an ELF file.

#### (*Section) Data 

``` go 
func (s *Section) Data() ([]byte, error)
```

Data reads and returns the contents of the ELF section. Even if the section is stored compressed in the ELF file, Data returns uncompressed data.

For an SHT_NOBITS section, Data always returns a non-nil error.

#### (*Section) Open 

``` go 
func (s *Section) Open() io.ReadSeeker
```

Open returns a new ReadSeeker reading the ELF section. Even if the section is stored compressed in the ELF file, the ReadSeeker reads uncompressed data.

For an SHT_NOBITS section, all calls to the opened reader will return a non-nil error.

### type Section32 

``` go 
type Section32 struct {
	Name      uint32 /* Section name (index into the section header string table). */
	Type      uint32 /* Section type. */
	Flags     uint32 /* Section flags. */
	Addr      uint32 /* Address in memory image. */
	Off       uint32 /* Offset in file. */
	Size      uint32 /* Size in bytes. */
	Link      uint32 /* Index of a related section. */
	Info      uint32 /* Depends on section type. */
	Addralign uint32 /* Alignment in bytes. */
	Entsize   uint32 /* Size of each entry in section. */
}
```

ELF32 Section header.

### type Section64 

``` go 
type Section64 struct {
	Name      uint32 /* Section name (index into the section header string table). */
	Type      uint32 /* Section type. */
	Flags     uint64 /* Section flags. */
	Addr      uint64 /* Address in memory image. */
	Off       uint64 /* Offset in file. */
	Size      uint64 /* Size in bytes. */
	Link      uint32 /* Index of a related section. */
	Info      uint32 /* Depends on section type. */
	Addralign uint64 /* Alignment in bytes. */
	Entsize   uint64 /* Size of each entry in section. */
}
```

ELF64 Section header.

### type SectionFlag 

``` go 
type SectionFlag uint32
```

Section flags.

``` go 
const (
	SHF_WRITE            SectionFlag = 0x1        /* Section contains writable data. */
	SHF_ALLOC            SectionFlag = 0x2        /* Section occupies memory. */
	SHF_EXECINSTR        SectionFlag = 0x4        /* Section contains instructions. */
	SHF_MERGE            SectionFlag = 0x10       /* Section may be merged. */
	SHF_STRINGS          SectionFlag = 0x20       /* Section contains strings. */
	SHF_INFO_LINK        SectionFlag = 0x40       /* sh_info holds section index. */
	SHF_LINK_ORDER       SectionFlag = 0x80       /* Special ordering requirements. */
	SHF_OS_NONCONFORMING SectionFlag = 0x100      /* OS-specific processing required. */
	SHF_GROUP            SectionFlag = 0x200      /* Member of section group. */
	SHF_TLS              SectionFlag = 0x400      /* Section contains TLS data. */
	SHF_COMPRESSED       SectionFlag = 0x800      /* Section is compressed. */
	SHF_MASKOS           SectionFlag = 0x0ff00000 /* OS-specific semantics. */
	SHF_MASKPROC         SectionFlag = 0xf0000000 /* Processor-specific semantics. */
)
```

#### (SectionFlag) GoString 

``` go 
func (i SectionFlag) GoString() string
```

#### (SectionFlag) String 

``` go 
func (i SectionFlag) String() string
```

### type SectionHeader 

``` go 
type SectionHeader struct {
	Name      string
	Type      SectionType
	Flags     SectionFlag
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	Addralign uint64
	Entsize   uint64

	// FileSize is the size of this section in the file in bytes.
	// If a section is compressed, FileSize is the size of the
	// compressed data, while Size (above) is the size of the
	// uncompressed data.
	FileSize uint64
}
```

A SectionHeader represents a single ELF section header.

### type SectionIndex 

``` go 
type SectionIndex int
```

Special section indices.

``` go 
const (
	SHN_UNDEF     SectionIndex = 0      /* Undefined, missing, irrelevant. */
	SHN_LORESERVE SectionIndex = 0xff00 /* First of reserved range. */
	SHN_LOPROC    SectionIndex = 0xff00 /* First processor-specific. */
	SHN_HIPROC    SectionIndex = 0xff1f /* Last processor-specific. */
	SHN_LOOS      SectionIndex = 0xff20 /* First operating system-specific. */
	SHN_HIOS      SectionIndex = 0xff3f /* Last operating system-specific. */
	SHN_ABS       SectionIndex = 0xfff1 /* Absolute values. */
	SHN_COMMON    SectionIndex = 0xfff2 /* Common data. */
	SHN_XINDEX    SectionIndex = 0xffff /* Escape; index stored elsewhere. */
	SHN_HIRESERVE SectionIndex = 0xffff /* Last of reserved range. */
)
```

#### (SectionIndex) GoString 

``` go 
func (i SectionIndex) GoString() string
```

#### (SectionIndex) String 

``` go 
func (i SectionIndex) String() string
```

### type SectionType 

``` go 
type SectionType uint32
```

Section type.

``` go 
const (
	SHT_NULL           SectionType = 0          /* inactive */
	SHT_PROGBITS       SectionType = 1          /* program defined information */
	SHT_SYMTAB         SectionType = 2          /* symbol table section */
	SHT_STRTAB         SectionType = 3          /* string table section */
	SHT_RELA           SectionType = 4          /* relocation section with addends */
	SHT_HASH           SectionType = 5          /* symbol hash table section */
	SHT_DYNAMIC        SectionType = 6          /* dynamic section */
	SHT_NOTE           SectionType = 7          /* note section */
	SHT_NOBITS         SectionType = 8          /* no space section */
	SHT_REL            SectionType = 9          /* relocation section - no addends */
	SHT_SHLIB          SectionType = 10         /* reserved - purpose unknown */
	SHT_DYNSYM         SectionType = 11         /* dynamic symbol table section */
	SHT_INIT_ARRAY     SectionType = 14         /* Initialization function pointers. */
	SHT_FINI_ARRAY     SectionType = 15         /* Termination function pointers. */
	SHT_PREINIT_ARRAY  SectionType = 16         /* Pre-initialization function ptrs. */
	SHT_GROUP          SectionType = 17         /* Section group. */
	SHT_SYMTAB_SHNDX   SectionType = 18         /* Section indexes (see SHN_XINDEX). */
	SHT_LOOS           SectionType = 0x60000000 /* First of OS specific semantics */
	SHT_GNU_ATTRIBUTES SectionType = 0x6ffffff5 /* GNU object attributes */
	SHT_GNU_HASH       SectionType = 0x6ffffff6 /* GNU hash table */
	SHT_GNU_LIBLIST    SectionType = 0x6ffffff7 /* GNU prelink library list */
	SHT_GNU_VERDEF     SectionType = 0x6ffffffd /* GNU version definition section */
	SHT_GNU_VERNEED    SectionType = 0x6ffffffe /* GNU version needs section */
	SHT_GNU_VERSYM     SectionType = 0x6fffffff /* GNU version symbol table */
	SHT_HIOS           SectionType = 0x6fffffff /* Last of OS specific semantics */
	SHT_LOPROC         SectionType = 0x70000000 /* reserved range for processor */
	SHT_MIPS_ABIFLAGS  SectionType = 0x7000002a /* .MIPS.abiflags */
	SHT_HIPROC         SectionType = 0x7fffffff /* specific section header types */
	SHT_LOUSER         SectionType = 0x80000000 /* reserved range for application */
	SHT_HIUSER         SectionType = 0xffffffff /* specific indexes */
)
```

#### (SectionType) GoString 

``` go 
func (i SectionType) GoString() string
```

#### (SectionType) String 

``` go 
func (i SectionType) String() string
```

### type Sym32 

``` go 
type Sym32 struct {
	Name  uint32
	Value uint32
	Size  uint32
	Info  uint8
	Other uint8
	Shndx uint16
}
```

ELF32 Symbol.

### type Sym64 

``` go 
type Sym64 struct {
	Name  uint32 /* String table index of name. */
	Info  uint8  /* Type and binding information. */
	Other uint8  /* Reserved (not used). */
	Shndx uint16 /* Section index of symbol. */
	Value uint64 /* Symbol value. */
	Size  uint64 /* Size of associated object. */
}
```

ELF64 symbol table entries.

### type SymBind 

``` go 
type SymBind int
```

Symbol Binding - ELFNN_ST_BIND - st_info

``` go 
const (
	STB_LOCAL  SymBind = 0  /* Local symbol */
	STB_GLOBAL SymBind = 1  /* Global symbol */
	STB_WEAK   SymBind = 2  /* like global - lower precedence */
	STB_LOOS   SymBind = 10 /* Reserved range for operating system */
	STB_HIOS   SymBind = 12 /*   specific semantics. */
	STB_LOPROC SymBind = 13 /* reserved range for processor */
	STB_HIPROC SymBind = 15 /*   specific semantics. */
)
```

#### func ST_BIND 

``` go 
func ST_BIND(info uint8) SymBind
```

#### (SymBind) GoString 

``` go 
func (i SymBind) GoString() string
```

#### (SymBind) String 

``` go 
func (i SymBind) String() string
```

### type SymType 

``` go 
type SymType int
```

Symbol type - ELFNN_ST_TYPE - st_info

``` go 
const (
	STT_NOTYPE  SymType = 0  /* Unspecified type. */
	STT_OBJECT  SymType = 1  /* Data object. */
	STT_FUNC    SymType = 2  /* Function. */
	STT_SECTION SymType = 3  /* Section. */
	STT_FILE    SymType = 4  /* Source file. */
	STT_COMMON  SymType = 5  /* Uninitialized common block. */
	STT_TLS     SymType = 6  /* TLS object. */
	STT_LOOS    SymType = 10 /* Reserved range for operating system */
	STT_HIOS    SymType = 12 /*   specific semantics. */
	STT_LOPROC  SymType = 13 /* reserved range for processor */
	STT_HIPROC  SymType = 15 /*   specific semantics. */
)
```

#### func ST_TYPE 

``` go 
func ST_TYPE(info uint8) SymType
```

#### (SymType) GoString 

``` go 
func (i SymType) GoString() string
```

#### (SymType) String 

``` go 
func (i SymType) String() string
```

### type SymVis 

``` go 
type SymVis int
```

Symbol visibility - ELFNN_ST_VISIBILITY - st_other

``` go 
const (
	STV_DEFAULT   SymVis = 0x0 /* Default visibility (see binding). */
	STV_INTERNAL  SymVis = 0x1 /* Special meaning in relocatable objects. */
	STV_HIDDEN    SymVis = 0x2 /* Not visible. */
	STV_PROTECTED SymVis = 0x3 /* Visible but not preemptible. */
)
```

#### func ST_VISIBILITY 

``` go 
func ST_VISIBILITY(other uint8) SymVis
```

#### (SymVis) GoString 

``` go 
func (i SymVis) GoString() string
```

#### (SymVis) String 

``` go 
func (i SymVis) String() string
```

### type Symbol 

``` go 
type Symbol struct {
	Name        string
	Info, Other byte
	Section     SectionIndex
	Value, Size uint64

	// Version and Library are present only for the dynamic symbol
	// table.
	Version string
	Library string
}
```

A Symbol represents an entry in an ELF symbol table section.

### type Type 

``` go 
type Type uint16
```

Type is found in Header.Type.

``` go 
const (
	ET_NONE   Type = 0      /* Unknown type. */
	ET_REL    Type = 1      /* Relocatable. */
	ET_EXEC   Type = 2      /* Executable. */
	ET_DYN    Type = 3      /* Shared object. */
	ET_CORE   Type = 4      /* Core file. */
	ET_LOOS   Type = 0xfe00 /* First operating system specific. */
	ET_HIOS   Type = 0xfeff /* Last operating system-specific. */
	ET_LOPROC Type = 0xff00 /* First processor-specific. */
	ET_HIPROC Type = 0xffff /* Last processor-specific. */
)
```

#### (Type) GoString 

``` go 
func (i Type) GoString() string
```

#### (Type) String 

``` go 
func (i Type) String() string
```

### type Version 

``` go 
type Version byte
```

Version is found in Header.Ident[EI_VERSION] and Header.Version.

``` go 
const (
	EV_NONE    Version = 0
	EV_CURRENT Version = 1
)
```

#### (Version) GoString 

``` go 
func (i Version) GoString() string
```

#### (Version) String 

``` go 
func (i Version) String() string
```