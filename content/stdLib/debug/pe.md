+++
title = "pe"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/debug/pe@go1.23.0](https://pkg.go.dev/debug/pe@go1.23.0)

Package pe implements access to PE (Microsoft Windows Portable Executable) files.

​	`pe` 包实现对 PE（Microsoft Windows Portable Executable）文件的访问。

## Security 安全性

This package is not designed to be hardened against adversarial inputs, and is outside the scope of https://go.dev/security/policy. In particular, only basic validation is done when parsing object files. As such, care should be taken when parsing untrusted inputs, as parsing malformed files may consume significant resources, or cause panics.

​	该包没有设计用于抵御对抗性输入，并且超出了 https://go.dev/security/policy 的范围。特别地，在解析对象文件时仅进行基本验证。因此，在解析不受信任的输入时应当小心，因为解析格式错误的文件可能会消耗大量资源或导致崩溃。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/pe/pe.go;l=89)

``` go 
const (
	IMAGE_FILE_MACHINE_UNKNOWN     = 0x0
	IMAGE_FILE_MACHINE_AM33        = 0x1d3
	IMAGE_FILE_MACHINE_AMD64       = 0x8664
	IMAGE_FILE_MACHINE_ARM         = 0x1c0
	IMAGE_FILE_MACHINE_ARMNT       = 0x1c4
	IMAGE_FILE_MACHINE_ARM64       = 0xaa64
	IMAGE_FILE_MACHINE_EBC         = 0xebc
	IMAGE_FILE_MACHINE_I386        = 0x14c
	IMAGE_FILE_MACHINE_IA64        = 0x200
	IMAGE_FILE_MACHINE_LOONGARCH32 = 0x6232
	IMAGE_FILE_MACHINE_LOONGARCH64 = 0x6264
	IMAGE_FILE_MACHINE_M32R        = 0x9041
	IMAGE_FILE_MACHINE_MIPS16      = 0x266
	IMAGE_FILE_MACHINE_MIPSFPU     = 0x366
	IMAGE_FILE_MACHINE_MIPSFPU16   = 0x466
	IMAGE_FILE_MACHINE_POWERPC     = 0x1f0
	IMAGE_FILE_MACHINE_POWERPCFP   = 0x1f1
	IMAGE_FILE_MACHINE_R4000       = 0x166
	IMAGE_FILE_MACHINE_SH3         = 0x1a2
	IMAGE_FILE_MACHINE_SH3DSP      = 0x1a3
	IMAGE_FILE_MACHINE_SH4         = 0x1a6
	IMAGE_FILE_MACHINE_SH5         = 0x1a8
	IMAGE_FILE_MACHINE_THUMB       = 0x1c2
	IMAGE_FILE_MACHINE_WCEMIPSV2   = 0x169
	IMAGE_FILE_MACHINE_RISCV32     = 0x5032
	IMAGE_FILE_MACHINE_RISCV64     = 0x5064
	IMAGE_FILE_MACHINE_RISCV128    = 0x5128
)
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/pe/pe.go;l=120)

``` go 
const (
	IMAGE_DIRECTORY_ENTRY_EXPORT         = 0
	IMAGE_DIRECTORY_ENTRY_IMPORT         = 1
	IMAGE_DIRECTORY_ENTRY_RESOURCE       = 2
	IMAGE_DIRECTORY_ENTRY_EXCEPTION      = 3
	IMAGE_DIRECTORY_ENTRY_SECURITY       = 4
	IMAGE_DIRECTORY_ENTRY_BASERELOC      = 5
	IMAGE_DIRECTORY_ENTRY_DEBUG          = 6
	IMAGE_DIRECTORY_ENTRY_ARCHITECTURE   = 7
	IMAGE_DIRECTORY_ENTRY_GLOBALPTR      = 8
	IMAGE_DIRECTORY_ENTRY_TLS            = 9
	IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG    = 10
	IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT   = 11
	IMAGE_DIRECTORY_ENTRY_IAT            = 12
	IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT   = 13
	IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR = 14
)
```

IMAGE_DIRECTORY_ENTRY constants

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/pe/pe.go;l=139)

``` go 
const (
	IMAGE_FILE_RELOCS_STRIPPED         = 0x0001
	IMAGE_FILE_EXECUTABLE_IMAGE        = 0x0002
	IMAGE_FILE_LINE_NUMS_STRIPPED      = 0x0004
	IMAGE_FILE_LOCAL_SYMS_STRIPPED     = 0x0008
	IMAGE_FILE_AGGRESIVE_WS_TRIM       = 0x0010
	IMAGE_FILE_LARGE_ADDRESS_AWARE     = 0x0020
	IMAGE_FILE_BYTES_REVERSED_LO       = 0x0080
	IMAGE_FILE_32BIT_MACHINE           = 0x0100
	IMAGE_FILE_DEBUG_STRIPPED          = 0x0200
	IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP = 0x0400
	IMAGE_FILE_NET_RUN_FROM_SWAP       = 0x0800
	IMAGE_FILE_SYSTEM                  = 0x1000
	IMAGE_FILE_DLL                     = 0x2000
	IMAGE_FILE_UP_SYSTEM_ONLY          = 0x4000
	IMAGE_FILE_BYTES_REVERSED_HI       = 0x8000
)
```

Values of IMAGE_FILE_HEADER.Characteristics. These can be combined together.

IMAGE_FILE_HEADER.Characteristics 的取值。可以将它们组合在一起。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/pe/pe.go;l=158)

``` go 
const (
	IMAGE_SUBSYSTEM_UNKNOWN                  = 0
	IMAGE_SUBSYSTEM_NATIVE                   = 1
	IMAGE_SUBSYSTEM_WINDOWS_GUI              = 2
	IMAGE_SUBSYSTEM_WINDOWS_CUI              = 3
	IMAGE_SUBSYSTEM_OS2_CUI                  = 5
	IMAGE_SUBSYSTEM_POSIX_CUI                = 7
	IMAGE_SUBSYSTEM_NATIVE_WINDOWS           = 8
	IMAGE_SUBSYSTEM_WINDOWS_CE_GUI           = 9
	IMAGE_SUBSYSTEM_EFI_APPLICATION          = 10
	IMAGE_SUBSYSTEM_EFI_BOOT_SERVICE_DRIVER  = 11
	IMAGE_SUBSYSTEM_EFI_RUNTIME_DRIVER       = 12
	IMAGE_SUBSYSTEM_EFI_ROM                  = 13
	IMAGE_SUBSYSTEM_XBOX                     = 14
	IMAGE_SUBSYSTEM_WINDOWS_BOOT_APPLICATION = 16
)
```

OptionalHeader64.Subsystem and OptionalHeader32.Subsystem values.

​	OptionalHeader64.Subsystem 和 OptionalHeader32.Subsystem 的取值。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/pe/pe.go;l=177)

``` go 
const (
	IMAGE_DLLCHARACTERISTICS_HIGH_ENTROPY_VA       = 0x0020
	IMAGE_DLLCHARACTERISTICS_DYNAMIC_BASE          = 0x0040
	IMAGE_DLLCHARACTERISTICS_FORCE_INTEGRITY       = 0x0080
	IMAGE_DLLCHARACTERISTICS_NX_COMPAT             = 0x0100
	IMAGE_DLLCHARACTERISTICS_NO_ISOLATION          = 0x0200
	IMAGE_DLLCHARACTERISTICS_NO_SEH                = 0x0400
	IMAGE_DLLCHARACTERISTICS_NO_BIND               = 0x0800
	IMAGE_DLLCHARACTERISTICS_APPCONTAINER          = 0x1000
	IMAGE_DLLCHARACTERISTICS_WDM_DRIVER            = 0x2000
	IMAGE_DLLCHARACTERISTICS_GUARD_CF              = 0x4000
	IMAGE_DLLCHARACTERISTICS_TERMINAL_SERVER_AWARE = 0x8000
)
```

OptionalHeader64.DllCharacteristics and OptionalHeader32.DllCharacteristics values. These can be combined together.

​	OptionalHeader64.DllCharacteristics 和 OptionalHeader32.DllCharacteristics 的取值。可以将它们组合在一起。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/pe/section.go;l=110)

``` go 
const (
	IMAGE_SCN_CNT_CODE               = 0x00000020
	IMAGE_SCN_CNT_INITIALIZED_DATA   = 0x00000040
	IMAGE_SCN_CNT_UNINITIALIZED_DATA = 0x00000080
	IMAGE_SCN_LNK_COMDAT             = 0x00001000
	IMAGE_SCN_MEM_DISCARDABLE        = 0x02000000
	IMAGE_SCN_MEM_EXECUTE            = 0x20000000
	IMAGE_SCN_MEM_READ               = 0x40000000
	IMAGE_SCN_MEM_WRITE              = 0x80000000
)
```

Section characteristics flags.

​	节的特征标志。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/pe/symbol.go;l=174)

``` go 
const (
	IMAGE_COMDAT_SELECT_NODUPLICATES = 1
	IMAGE_COMDAT_SELECT_ANY          = 2
	IMAGE_COMDAT_SELECT_SAME_SIZE    = 3
	IMAGE_COMDAT_SELECT_EXACT_MATCH  = 4
	IMAGE_COMDAT_SELECT_ASSOCIATIVE  = 5
	IMAGE_COMDAT_SELECT_LARGEST      = 6
)
```

These constants make up the possible values for the 'Selection' field in an AuxFormat5.

​	这些常量构成 AuxFormat5 中 'Selection' 字段的可能取值。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/pe/symbol.go;l=16)

``` go 
const COFFSymbolSize = 18
```

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type COFFSymbol  <- go1.1

``` go 
type COFFSymbol struct {
	Name               [8]uint8
	Value              uint32
	SectionNumber      int16
	Type               uint16
	StorageClass       uint8
	NumberOfAuxSymbols uint8
}
```

COFFSymbol represents single COFF symbol table record.

​	COFFSymbol 表示单个 COFF 符号表记录。

#### (*COFFSymbol) FullName  <- go1.8

``` go 
func (sym *COFFSymbol) FullName(st StringTable) (string, error)
```

FullName finds real name of symbol sym. Normally name is stored in sym.Name, but if it is longer then 8 characters, it is stored in COFF string table st instead.

​	FullName 查找符号 sym 的实际名称。通常名称存储在 sym.Name 中，但如果超过 8 个字符，则存储在 COFF 字符串表 st 中。

### type COFFSymbolAuxFormat5  <- go1.19

``` go 
type COFFSymbolAuxFormat5 struct {
	Size           uint32
	NumRelocs      uint16
	NumLineNumbers uint16
	Checksum       uint32
	SecNum         uint16
	Selection      uint8
	// contains filtered or unexported fields
}
```

COFFSymbolAuxFormat5 describes the expected form of an aux symbol attached to a section definition symbol. The PE format defines a number of different aux symbol formats: format 1 for function definitions, format 2 for .be and .ef symbols, and so on. Format 5 holds extra info associated with a section definition, including number of relocations + line numbers, as well as COMDAT info. See https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-format-5-section-definitions for more on what's going on here.

​	COFFSymbolAuxFormat5 描述了附加到节定义符号的辅助符号的预期形式。PE 格式定义了多种不同的辅助符号格式：格式 1 用于函数定义，格式 2 用于 .be 和 .ef 符号，等等。格式 5 包含与节定义相关的额外信息，包括重定位数目+行号以及 COMDAT 信息。有关此处发生的更多信息，请参阅 https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-format-5-section-definitions。

### type DataDirectory  <- go1.3

``` go 
type DataDirectory struct {
	VirtualAddress uint32
	Size           uint32
}
```

### type File 

``` go 
type File struct {
	FileHeader
	OptionalHeader any // of type *OptionalHeader32 or *OptionalHeader64 类型为 *OptionalHeader32 或 *OptionalHeader64
	Sections       []*Section
	Symbols        []*Symbol    // COFF symbols with auxiliary symbol records removed 去除了附加符号记录的 COFF 符号
	COFFSymbols    []COFFSymbol // all COFF symbols (including auxiliary symbol records) 所有 COFF 符号（包括附加符号记录）
	StringTable    StringTable
	// contains filtered or unexported fields
}
```

A File represents an open PE file.

File 表示一个打开的 PE 文件。

#### func NewFile 

``` go 
func NewFile(r io.ReaderAt) (*File, error)
```

NewFile creates a new File for accessing a PE binary in an underlying reader.

​	NewFile 创建一个用于访问底层 Reader 中的 PE 二进制文件的新 File。

#### func Open 

``` go 
func Open(name string) (*File, error)
```

Open opens the named file using os.Open and prepares it for use as a PE binary.

​	Open 使用 os.Open 打开指定的文件，并准备将其用作 PE 二进制文件。

#### (*File) COFFSymbolReadSectionDefAux  <- go1.19

``` go 
func (f *File) COFFSymbolReadSectionDefAux(idx int) (*COFFSymbolAuxFormat5, error)
```

COFFSymbolReadSectionDefAux returns a blob of axiliary information (including COMDAT info) for a section definition symbol. Here 'idx' is the index of a section symbol in the main COFFSymbol array for the File. Return value is a pointer to the appropriate aux symbol struct. For more info, see:

​	COFFSymbolReadSectionDefAux 返回与节定义符号相关的辅助信息块（包括 COMDAT 信息）。这里的 'idx' 是 File 的主 COFFSymbol 数组中节符号的索引。返回值是适当的辅助符号结构体的指针。了解更多信息，请参阅：

auxiliary symbols: https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-symbol-records 

​	辅助符号：https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-symbol-records 

COMDAT sections: https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#comdat-sections-object-only 

​	COMDAT 节：https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#comdat-sections-object-only 

auxiliary info for section definitions: https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-format-5-section-definitions

​	节定义的辅助信息：https://docs.microsoft.com/en-us/windows/win32/debug/pe-format#auxiliary-format-5-section-definitions

#### (*File) Close 

``` go 
func (f *File) Close() error
```

Close closes the File. If the File was created using NewFile directly instead of Open, Close has no effect.

​	`Close` 关闭 File。如果 File 是直接使用 NewFile 创建而不是使用 Open，Close 不产生任何效果。

#### (*File) DWARF 

``` go 
func (f *File) DWARF() (*dwarf.Data, error)
```

​	`DWARF` 返回 PE 文件的 DWARF 调试信息。

#### (*File) ImportedLibraries 

``` go 
func (f *File) ImportedLibraries() ([]string, error)
```

ImportedLibraries returns the names of all libraries referred to by the binary f that are expected to be linked with the binary at dynamic link time.

​	`ImportedLibraries` 返回二进制文件 f 引用的所有库的名称，这些库预计在动态链接时与该二进制文件链接。



#### (*File) ImportedSymbols 

``` go 
func (f *File) ImportedSymbols() ([]string, error)
```

ImportedSymbols returns the names of all symbols referred to by the binary f that are expected to be satisfied by other libraries at dynamic load time. It does not return weak symbols.

​	`ImportedSymbols` 返回二进制文件 f 引用的所有符号的名称，这些符号预计在动态加载时由其他库满足。它不返回弱符号。

#### (*File) Section 

``` go 
func (f *File) Section(name string) *Section
```

Section returns the first section with the given name, or nil if no such section exists.

​	`Section` 返回具有给定名称的第一个节，如果不存在这样的节，则返回 nil。

### type FileHeader 

``` go 
type FileHeader struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}
```

### type FormatError 

``` go 
type FormatError struct {
}
```

FormatError is unused. The type is retained for compatibility.

​	`FormatError` 未被使用。该类型保留是为了保持兼容性。

#### (*FormatError) Error 

``` go 
func (e *FormatError) Error() string
```

### type ImportDirectory 

``` go 
type ImportDirectory struct {
	OriginalFirstThunk uint32
	TimeDateStamp      uint32
	ForwarderChain     uint32
	Name               uint32
	FirstThunk         uint32
	// contains filtered or unexported fields
}
```

### type OptionalHeader32  <- go1.3

``` go 
type OptionalHeader32 struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint32
	SizeOfStackCommit           uint32
	SizeOfHeapReserve           uint32
	SizeOfHeapCommit            uint32
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]DataDirectory
}
```

### type OptionalHeader64  <- go1.3

``` go 
type OptionalHeader64 struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]DataDirectory
}
```

### type Reloc  <- go1.8

``` go 
type Reloc struct {
	VirtualAddress   uint32
	SymbolTableIndex uint32
	Type             uint16
}
```

Reloc represents a PE COFF relocation. Each section contains its own relocation list.

​	`Reloc` 表示一个 PE COFF 重定位。每个节都包含自己的重定位列表。

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
    // 嵌入 ReaderAt 以便使用 ReadAt 方法。
    // 不直接嵌入 SectionReader，
    // 避免 Read 和 Seek 冲突。
    // 如果客户端需要 Read 和 Seek，必须使用 Open() 方法，
    // 以避免与其他客户端争夺 Seek 偏移量。
	io.ReaderAt
	// contains filtered or unexported fields
}
```

Section provides access to PE COFF section.

​	`Section` 提供对 PE COFF 节的访问。

#### (*Section) Data 

``` go 
func (s *Section) Data() ([]byte, error)
```

Data reads and returns the contents of the PE section s.

​	`Data` 读取并返回 PE 节 s 的内容。

#### (*Section) Open 

``` go 
func (s *Section) Open() io.ReadSeeker
```

Open returns a new ReadSeeker reading the PE section s.

​	`Open` 返回一个新的 `ReadSeeker`，用于读取 PE 节 s。

### type SectionHeader 

``` go 
type SectionHeader struct {
	Name                 string
	VirtualSize          uint32
	VirtualAddress       uint32
	Size                 uint32
	Offset               uint32
	PointerToRelocations uint32
	PointerToLineNumbers uint32
	NumberOfRelocations  uint16
	NumberOfLineNumbers  uint16
	Characteristics      uint32
}
```

SectionHeader is similar to SectionHeader32 with Name field replaced by Go string.

​	`SectionHeader` 类似于 `SectionHeader32`，但 `Name` 字段替换为 Go 字符串。

### type SectionHeader32 

``` go 
type SectionHeader32 struct {
	Name                 [8]uint8
	VirtualSize          uint32
	VirtualAddress       uint32
	SizeOfRawData        uint32
	PointerToRawData     uint32
	PointerToRelocations uint32
	PointerToLineNumbers uint32
	NumberOfRelocations  uint16
	NumberOfLineNumbers  uint16
	Characteristics      uint32
}
```

SectionHeader32 represents real PE COFF section header.

​	`SectionHeader32` 表示真实的 PE COFF 节头。

### type StringTable  <- go1.8

``` go 
type StringTable []byte
```

StringTable is a COFF string table.

​	`StringTable` 是 COFF 字符串表。

#### (StringTable) String  <- go1.8

``` go 
func (st StringTable) String(start uint32) (string, error)
```

String extracts string from COFF string table st at offset start.

​	`String` 从 COFF 字符串表 st 的偏移 start 处提取字符串。

### type Symbol  <- go1.1

``` go 
type Symbol struct {
	Name          string
	Value         uint32
	SectionNumber int16
	Type          uint16
	StorageClass  uint8
}
```

Symbol is similar to COFFSymbol with Name field replaced by Go string. Symbol also does not have NumberOfAuxSymbols.

​	`Symbol` 类似于 `COFFSymbol`，但 `Name` 字段替换为 `Go` 字符串。`Symbol` 也不包含 `NumberOfAuxSymbols`。