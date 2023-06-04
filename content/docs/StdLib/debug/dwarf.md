+++
title = "dwarf"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# dwarf

https://pkg.go.dev/debug/dwarf@go1.20.1

Package dwarf provides access to DWARF debugging information loaded from executable files, as defined in the DWARF 2.0 Standard at http://dwarfstd.org/doc/dwarf-2.0.0.pdf.

包 dwarf 提供对从可执行文件加载的 DWARF 调试信息的访问，其定义在 DWARF 2.0 标准中，网址为 http://dwarfstd.org/doc/dwarf-2.0.0.pdf。

#### Security  安全性

This package is not designed to be hardened against adversarial inputs, and is outside the scope of https://go.dev/security/policy. In particular, only basic validation is done when parsing object files. As such, care should be taken when parsing untrusted inputs, as parsing malformed files may consume significant resources, or cause panics.

该包不针对对抗性输入进行强化设计，超出了 https://go.dev/security/policy 的范围。特别是，当解析对象文件时，仅进行基本验证。因此，在解析不受信任的输入时应谨慎，因为解析格式错误的文件可能会消耗大量资源或导致恐慌。



## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/debug/dwarf/line.go;l=730)

``` go 
var ErrUnknownPC = errors.New("ErrUnknownPC")
```

ErrUnknownPC is the error returned by LineReader.ScanPC when the seek PC is not covered by any entry in the line table.

当寻找的 PC 不在行表中的任何条目覆盖范围时，LineReader.ScanPC 返回 ErrUnknownPC。

## 函数

This section is empty.

## 类型

### type AddrType 

``` go 
type AddrType struct {
	BasicType
}
```

An AddrType represents a machine address type.

AddrType 表示机器地址类型。

### type ArrayType 

``` go 
type ArrayType struct {
	CommonType
	Type          Type
	StrideBitSize int64 // if > 0, number of bits to hold each element 如果 > 0，则表示每个元素所需的位数
	Count         int64 // if == -1, an incomplete array, like char x[]. 如果等于 -1，则表示不完整的数组，例如 char x[]。
}
```

An ArrayType represents a fixed size array type.

ArrayType 表示固定大小的数组类型。

#### (*ArrayType) Size 

``` go 
func (t *ArrayType) Size() int64
```

#### (*ArrayType) String 

``` go 
func (t *ArrayType) String() string
```

### type Attr 

``` go 
type Attr uint32
```

An Attr identifies the attribute type in a DWARF Entry's Field.

Attr 标识 DWARF 条目字段中的属性类型。

``` go 
const (
	AttrSibling        Attr = 0x01
	AttrLocation       Attr = 0x02
	AttrName           Attr = 0x03
	AttrOrdering       Attr = 0x09
	AttrByteSize       Attr = 0x0B
	AttrBitOffset      Attr = 0x0C
	AttrBitSize        Attr = 0x0D
	AttrStmtList       Attr = 0x10
	AttrLowpc          Attr = 0x11
	AttrHighpc         Attr = 0x12
	AttrLanguage       Attr = 0x13
	AttrDiscr          Attr = 0x15
	AttrDiscrValue     Attr = 0x16
	AttrVisibility     Attr = 0x17
	AttrImport         Attr = 0x18
	AttrStringLength   Attr = 0x19
	AttrCommonRef      Attr = 0x1A
	AttrCompDir        Attr = 0x1B
	AttrConstValue     Attr = 0x1C
	AttrContainingType Attr = 0x1D
	AttrDefaultValue   Attr = 0x1E
	AttrInline         Attr = 0x20
	AttrIsOptional     Attr = 0x21
	AttrLowerBound     Attr = 0x22
	AttrProducer       Attr = 0x25
	AttrPrototyped     Attr = 0x27
	AttrReturnAddr     Attr = 0x2A
	AttrStartScope     Attr = 0x2C
	AttrStrideSize     Attr = 0x2E
	AttrUpperBound     Attr = 0x2F
	AttrAbstractOrigin Attr = 0x31
	AttrAccessibility  Attr = 0x32
	AttrAddrClass      Attr = 0x33
	AttrArtificial     Attr = 0x34
	AttrBaseTypes      Attr = 0x35
	AttrCalling        Attr = 0x36
	AttrCount          Attr = 0x37
	AttrDataMemberLoc  Attr = 0x38
	AttrDeclColumn     Attr = 0x39
	AttrDeclFile       Attr = 0x3A
	AttrDeclLine       Attr = 0x3B
	AttrDeclaration    Attr = 0x3C
	AttrDiscrList      Attr = 0x3D
	AttrEncoding       Attr = 0x3E
	AttrExternal       Attr = 0x3F
	AttrFrameBase      Attr = 0x40
	AttrFriend         Attr = 0x41
	AttrIdentifierCase Attr = 0x42
	AttrMacroInfo      Attr = 0x43
	AttrNamelistItem   Attr = 0x44
	AttrPriority       Attr = 0x45
	AttrSegment        Attr = 0x46
	AttrSpecification  Attr = 0x47
	AttrStaticLink     Attr = 0x48
	AttrType           Attr = 0x49
	AttrUseLocation    Attr = 0x4A
	AttrVarParam       Attr = 0x4B
	AttrVirtuality     Attr = 0x4C
	AttrVtableElemLoc  Attr = 0x4D
	// The following are new in DWARF 3.
    // 以下是 DWARF 3 中的新内容。
	AttrAllocated     Attr = 0x4E
	AttrAssociated    Attr = 0x4F
	AttrDataLocation  Attr = 0x50
	AttrStride        Attr = 0x51
	AttrEntrypc       Attr = 0x52
	AttrUseUTF8       Attr = 0x53
	AttrExtension     Attr = 0x54
	AttrRanges        Attr = 0x55
	AttrTrampoline    Attr = 0x56
	AttrCallColumn    Attr = 0x57
	AttrCallFile      Attr = 0x58
	AttrCallLine      Attr = 0x59
	AttrDescription   Attr = 0x5A
	AttrBinaryScale   Attr = 0x5B
	AttrDecimalScale  Attr = 0x5C
	AttrSmall         Attr = 0x5D
	AttrDecimalSign   Attr = 0x5E
	AttrDigitCount    Attr = 0x5F
	AttrPictureString Attr = 0x60
	AttrMutable       Attr = 0x61
	AttrThreadsScaled Attr = 0x62
	AttrExplicit      Attr = 0x63
	AttrObjectPointer Attr = 0x64
	AttrEndianity     Attr = 0x65
	AttrElemental     Attr = 0x66
	AttrPure          Attr = 0x67
	AttrRecursive     Attr = 0x68
	// The following are new in DWARF 4.
    // 以下是 DWARF 4 中的新内容。
	AttrSignature      Attr = 0x69
	AttrMainSubprogram Attr = 0x6A
	AttrDataBitOffset  Attr = 0x6B
	AttrConstExpr      Attr = 0x6C
	AttrEnumClass      Attr = 0x6D
	AttrLinkageName    Attr = 0x6E
	// The following are new in DWARF 5.
    // 以下是 DWARF 5 中的新内容。
	AttrStringLengthBitSize  Attr = 0x6F
	AttrStringLengthByteSize Attr = 0x70
	AttrRank                 Attr = 0x71
	AttrStrOffsetsBase       Attr = 0x72
	AttrAddrBase             Attr = 0x73
	AttrRnglistsBase         Attr = 0x74
	AttrDwoName              Attr = 0x76
	AttrReference            Attr = 0x77
	AttrRvalueReference      Attr = 0x78
	AttrMacros               Attr = 0x79
	AttrCallAllCalls         Attr = 0x7A
	AttrCallAllSourceCalls   Attr = 0x7B
	AttrCallAllTailCalls     Attr = 0x7C
	AttrCallReturnPC         Attr = 0x7D
	AttrCallValue            Attr = 0x7E
	AttrCallOrigin           Attr = 0x7F
	AttrCallParameter        Attr = 0x80
	AttrCallPC               Attr = 0x81
	AttrCallTailCall         Attr = 0x82
	AttrCallTarget           Attr = 0x83
	AttrCallTargetClobbered  Attr = 0x84
	AttrCallDataLocation     Attr = 0x85
	AttrCallDataValue        Attr = 0x86
	AttrNoreturn             Attr = 0x87
	AttrAlignment            Attr = 0x88
	AttrExportSymbols        Attr = 0x89
	AttrDeleted              Attr = 0x8A
	AttrDefaulted            Attr = 0x8B
	AttrLoclistsBase         Attr = 0x8C
)
```

#### (Attr) GoString 

``` go 
func (a Attr) GoString() string
```

#### (Attr) String 

``` go 
func (i Attr) String() string
```

### type BasicType 

``` go 
type BasicType struct {
	CommonType
	BitSize       int64
	BitOffset     int64
	DataBitOffset int64
}
```

A BasicType holds fields common to all basic types.

BasicType 包含所有基本类型的共同字段。

See the documentation for StructField for more info on the interpretation of the BitSize/BitOffset/DataBitOffset fields.

有关 BitSize/BitOffset/DataBitOffset 字段解释的更多信息，请参阅 StructField 的文档。

#### (*BasicType) Basic 

``` go 
func (b *BasicType) Basic() *BasicType
```

#### (*BasicType) String 

``` go 
func (t *BasicType) String() string
```

### type BoolType 

``` go 
type BoolType struct {
	BasicType
}
```

A BoolType represents a boolean type.

BoolType 表示布尔类型。

### type CharType 

``` go 
type CharType struct {
	BasicType
}
```

A CharType represents a signed character type.

CharType 表示有符号字符类型。

### type Class  <- go1.5

``` go 
type Class int
```

A Class is the DWARF 4 class of an attribute value.

Class 是属性值的 DWARF 4 类。

In general, a given attribute's value may take on one of several possible classes defined by DWARF, each of which leads to a slightly different interpretation of the attribute.

一般来说，给定属性的值可以采用 DWARF 定义的多个可能类别之一，每个类别对属性的解释略有不同。

DWARF version 4 distinguishes attribute value classes more finely than previous versions of DWARF. The reader will disambiguate coarser classes from earlier versions of DWARF into the appropriate DWARF 4 class. For example, DWARF 2 uses "constant" for constants as well as all types of section offsets, but the reader will canonicalize attributes in DWARF 2 files that refer to section offsets to one of the Class*Ptr classes, even though these classes were only defined in DWARF 3.

DWARF 4 版本比以前的 DWARF 版本更精细地区分属性值类别。阅读器会将早期版本 DWARF 中更粗略的类别区分为适当的 DWARF 4 类别。例如，DWARF 2 将 "constant" 用于常量以及所有类型的节偏移量，但阅读器会将引用节偏移量的 DWARF 2 文件中的属性规范化为 Class*Ptr 类之一，即使这些类别仅在 DWARF 3 中定义。

``` go 
const (
	// ClassUnknown represents values of unknown DWARF class.
	// ClassUnknown 表示未知的 DWARF 类别的值。
	ClassUnknown Class = iota

	// ClassAddress represents values of type uint64 that are
	// addresses on the target machine.
    // ClassAddress 表示目标机器上的 uint64 类型的地址值。

	ClassAddress

	// ClassBlock represents values of type []byte whose
	// interpretation depends on the attribute.
    // ClassBlock 表示 []byte 类型的值，其解释取决于属性。
	ClassBlock

	// ClassConstant represents values of type int64 that are
	// constants. The interpretation of this constant depends on
	// the attribute.
    // ClassConstant 表示 int64 类型的常量值。此常量的解释取决于属性。

	ClassConstant

	// ClassExprLoc represents values of type []byte that contain
	// an encoded DWARF expression or location description.
    // ClassExprLoc 表示包含编码的 DWARF 表达式或位置描述的 []byte 类型的值。

	ClassExprLoc

	// ClassFlag represents values of type bool.
    // ClassFlag 表示 bool 类型的值。  
	ClassFlag

	// ClassLinePtr represents values that are an int64 offset
	// into the "line" section.
    // ClassLinePtr 表示一个 int64 偏移量，指向 "line" 节段。
	ClassLinePtr

	// ClassLocListPtr represents values that are an int64 offset
	// into the "loclist" section.
    // ClassLocListPtr 表示一个 int64 偏移量，指向 "loclist" 节段。
	ClassLocListPtr

	// ClassMacPtr represents values that are an int64 offset into
	// the "mac" section.
    // ClassMacPtr 表示一个 int64 偏移量，指向 "mac" 节段。
	ClassMacPtr

	// ClassRangeListPtr represents values that are an int64 offset into
	// the "rangelist" section.
    // ClassRangeListPtr 表示一个 int64 偏移量，指向 "rangelist" 节段。
	ClassRangeListPtr

	// ClassReference represents values that are an Offset offset
	// of an Entry in the info section (for use with Reader.Seek).
	// The DWARF specification combines ClassReference and
	// ClassReferenceSig into class "reference".
    // ClassReference 表示一个 Entry 在 info 节段中的 Offset 偏移量（用于 Reader.Seek）。
	// DWARF 规范将 ClassReference 和 ClassReferenceSig 组合成 "reference" 类别。
	ClassReference

	// ClassReferenceSig represents values that are a uint64 type
	// signature referencing a type Entry.
    // ClassReferenceSig 表示一个引用类型 Entry 的 uint64 类型签名。
	ClassReferenceSig

	// ClassString represents values that are strings. If the
	// compilation unit specifies the AttrUseUTF8 flag (strongly
	// recommended), the string value will be encoded in UTF-8.
	// Otherwise, the encoding is unspecified.
    // ClassString 表示字符串值。如果编译单元指定 AttrUseUTF8 标志（强烈推荐），
    // 字符串值将使用 UTF-8 编码。
	// 否则，编码是未指定的。
	ClassString

	// ClassReferenceAlt represents values of type int64 that are
	// an offset into the DWARF "info" section of an alternate
	// object file.
    // ClassReferenceAlt 表示 int64 类型的值，是一个对于备用目标文件中 DWARF "info" 节段的偏移量。
	ClassReferenceAlt

	// ClassStringAlt represents values of type int64 that are an
	// offset into the DWARF string section of an alternate object
	// file.
    // ClassStringAlt 表示 int64 类型的值，是一个对于备用目标文件中 DWARF 字符串节段的偏移量。
	ClassStringAlt

	// ClassAddrPtr represents values that are an int64 offset
	// into the "addr" section.
    // ClassAddrPtr 表示一个 int64 偏移量，指向 "addr" 节段。
	ClassAddrPtr

	// ClassLocList represents values that are an int64 offset
	// into the "loclists" section.
    // ClassLocList 表示一个 int64 偏移量，指向 "loclists" 节段。
	ClassLocList

	// ClassRngList represents values that are a uint64 offset
	// from the base of the "rnglists" section.
    // ClassRngList 表示一个从 "rnglists" 节段的基地址开始的 uint64 偏移量。
	ClassRngList

	// ClassRngListsPtr represents values that are an int64 offset
	// into the "rnglists" section. These are used as the base for
	// ClassRngList values.
    // ClassRngListsPtr 表示一个 int64 偏移量，指向 "rnglists" 节段。
    // 这些偏移量用作 ClassRngList 值的基地址。
	ClassRngListsPtr

	// ClassStrOffsetsPtr represents values that are an int64
	// offset into the "str_offsets" section.
    // ClassStrOffsetsPtr 表示一个 int64 偏移量，指向 "str_offsets" 节段。
	ClassStrOffsetsPtr
)
```

#### (Class) GoString  <- go1.5

``` go 
func (i Class) GoString() string
```

#### (Class) String  <- go1.5

``` go 
func (i Class) String() string
```

### type CommonType 

``` go 
type CommonType struct {
	ByteSize int64  // size of value of this type, in bytes 此类型的值的大小，以字节为单位
	Name     string // name that can be used to refer to type 用于引用该类型的名称
}
```

A CommonType holds fields common to multiple types. If a field is not known or not applicable for a given type, the zero value is used.

CommonType 包含多个类型共有的字段。如果某个类型的字段未知或不适用，则使用零值。

#### (*CommonType) Common 

``` go 
func (c *CommonType) Common() *CommonType
```

#### (*CommonType) Size 

``` go 
func (c *CommonType) Size() int64
```

### type ComplexType 

``` go 
type ComplexType struct {
	BasicType
}
```

A ComplexType represents a complex floating point type.

ComplexType 表示复数浮点数类型。

### type Data 

``` go 
type Data struct {
	// contains filtered or unexported fields
}
```

Data represents the DWARF debugging information loaded from an executable file (for example, an ELF or Mach-O executable).

Data 表示从可执行文件（例如 ELF 或 Mach-O 可执行文件）加载的 DWARF 调试信息。

#### func New 

``` go 
func New(abbrev, aranges, frame, info, line, pubnames, ranges, str []byte) (*Data, error)
```

New returns a new Data object initialized from the given parameters. Rather than calling this function directly, clients should typically use the DWARF method of the File type of the appropriate package debug/elf, debug/macho, or debug/pe.

New 从给定的参数初始化并返回一个新的 Data 对象。通常情况下，客户端应该使用相应的 debug/elf、debug/macho 或 debug/pe 包的 File 类型的 DWARF 方法，而不是直接调用此函数。

The []byte arguments are the data from the corresponding debug section in the object file; for example, for an ELF object, abbrev is the contents of the ".debug_abbrev" section.

[]byte 参数是对象文件中相应调试节的数据；例如，对于 ELF 对象，abbrev 是 ".debug_abbrev" 节的内容。

#### (*Data) AddSection  <- go1.14

``` go 
func (d *Data) AddSection(name string, contents []byte) error
```

AddSection adds another DWARF section by name. The name should be a DWARF section name such as ".debug_addr", ".debug_str_offsets", and so forth. This approach is used for new DWARF sections added in DWARF 5 and later.

AddSection 按名称添加另一个 DWARF 节。名称应为 DWARF 节的名称，例如 ".debug_addr"、".debug_str_offsets" 等。这种方法用于在 DWARF 5 及更高版本中添加的新的 DWARF 节。

#### (*Data) AddTypes  <- go1.3

``` go 
func (d *Data) AddTypes(name string, types []byte) error
```

AddTypes will add one .debug_types section to the DWARF data. A typical object with DWARF version 4 debug info will have multiple .debug_types sections. The name is used for error reporting only, and serves to distinguish one .debug_types section from another.

AddTypes 将一个 .debug_types 节添加到 DWARF 数据中。具有 DWARF 版本 4 调试信息的典型对象将具有多个 .debug_types 节。名称仅用于错误报告，并用于区分一个 .debug_types 节和另一个 .debug_types 节。

#### (*Data) LineReader  <- go1.5

``` go 
func (d *Data) LineReader(cu *Entry) (*LineReader, error)
```

LineReader returns a new reader for the line table of compilation unit cu, which must be an Entry with tag TagCompileUnit.

LineReader 返回编译单元 cu 的行表的新读取器，cu 必须是具有标签 TagCompileUnit 的 Entry。

If this compilation unit has no line table, it returns nil, nil.

如果此编译单元没有行表，则返回 nil, nil。

#### (*Data) Ranges  <- go1.7

``` go 
func (d *Data) Ranges(e *Entry) ([][2]uint64, error)
```

Ranges returns the PC ranges covered by e, a slice of [low,high) pairs. Only some entry types, such as TagCompileUnit or TagSubprogram, have PC ranges; for others, this will return nil with no error.

Ranges 返回 e 覆盖的 PC 范围，即 [low,high) 对的切片。只有某些条目类型（例如 TagCompileUnit 或 TagSubprogram）具有 PC 范围；对于其他条目，将返回 nil，且无错误。

#### (*Data) Reader 

``` go 
func (d *Data) Reader() *Reader
```

Reader returns a new Reader for Data. The reader is positioned at byte offset 0 in the DWARF "info" section.

Reader 返回 Data 的新读取器。读取器位于 DWARF "info" 节的字节偏移量 0 处。

#### (*Data) Type 

``` go 
func (d *Data) Type(off Offset) (Type, error)
```

Type reads the type at off in the DWARF "info" section.

Type 在 DWARF "info" 节的 off 处读取类型。

### type DecodeError 

``` go 
type DecodeError struct {
	Name   string
	Offset Offset
	Err    string
}
```

#### (DecodeError) Error 

``` go 
func (e DecodeError) Error() string
```

### type DotDotDotType 

``` go 
type DotDotDotType struct {
	CommonType
}
```

A DotDotDotType represents the variadic ... function parameter.

一个 DotDotDotType 表示可变参数 ... 的函数参数。

#### (*DotDotDotType) String 

``` go 
func (t *DotDotDotType) String() string
```

### type Entry 

``` go 
type Entry struct {
	Offset   Offset // offset of Entry in DWARF info DWARF 信息中条目的偏移量
	Tag      Tag    // tag (kind of Entry) 标签（条目的类型）
	Children bool   // whether Entry is followed by children 条目后是否有子条目
	Field    []Field
}
```

An entry is a sequence of attribute/value pairs.

一个 entry 是属性/值对的序列。

#### (*Entry) AttrField  <- go1.5

``` go 
func (e *Entry) AttrField(a Attr) *Field
```

AttrField returns the Field associated with attribute Attr in Entry, or nil if there is no such attribute.

AttrField 返回与 Entry 中的属性 Attr 关联的 Field，如果没有该属性，则返回 nil。

#### (*Entry) Val 

``` go 
func (e *Entry) Val(a Attr) any
```

Val returns the value associated with attribute Attr in Entry, or nil if there is no such attribute.

Val 返回与 Entry 中的属性 Attr 关联的值，如果没有该属性，则返回 nil。

A common idiom is to merge the check for nil return with the check that the value has the expected dynamic type, as in:

一个常见的习惯用法是将对 nil 返回的检查与对值具有期望的动态类型的检查合并在一起，例如：

```
v, ok := e.Val(AttrSibling).(int64)
```

### type EnumType 

``` go 
type EnumType struct {
	CommonType
	EnumName string
	Val      []*EnumValue
}
```

An EnumType represents an enumerated type. The only indication of its native integer type is its ByteSize (inside CommonType).

一个 EnumType 表示一个枚举类型。其本机整数类型的唯一指示是其 ByteSize（在 CommonType 中）。

#### (*EnumType) String 

``` go 
func (t *EnumType) String() string
```

### type EnumValue 

``` go 
type EnumValue struct {
	Name string
	Val  int64
}
```

An EnumValue represents a single enumeration value.

一个 EnumValue 表示一个单独的枚举值。

### type Field 

``` go 
type Field struct {
	Attr  Attr
	Val   any
	Class Class
}
```

A Field is a single attribute/value pair in an Entry.

Field 是 Entry 中的单个属性/值对。

A value can be one of several "attribute classes" defined by DWARF. The Go types corresponding to each class are:

值可以是 DWARF 定义的几种 "属性类别" 之一。对应于每个类别的 Go 类型如下：

```
DWARF class       Go type        Class
-----------       -------        -----
address           uint64         ClassAddress
block             []byte         ClassBlock
constant          int64          ClassConstant
flag              bool           ClassFlag
reference
  to info         dwarf.Offset   ClassReference
  to type unit    uint64         ClassReferenceSig
string            string         ClassString
exprloc           []byte         ClassExprLoc
lineptr           int64          ClassLinePtr
loclistptr        int64          ClassLocListPtr
macptr            int64          ClassMacPtr
rangelistptr      int64          ClassRangeListPtr
```

For unrecognized or vendor-defined attributes, Class may be ClassUnknown.

对于未被识别或供应商定义的属性，Class 可能是 ClassUnknown。

### type FloatType 

``` go 
type FloatType struct {
	BasicType
}
```

A FloatType represents a floating point type.

FuncType 表示一个函数类型。

### type FuncType 

``` go 
type FuncType struct {
	CommonType
	ReturnType Type
	ParamType  []Type
}
```

A FuncType represents a function type.

IntType 表示一个有符号整数类型。

#### (*FuncType) String 

``` go 
func (t *FuncType) String() string
```

### type IntType 

``` go 
type IntType struct {
	BasicType
}
```

An IntType represents a signed integer type.

### type LineEntry  <- go1.5

``` go 
type LineEntry struct {
	// Address is the program-counter value of a machine
	// instruction generated by the compiler. This LineEntry
	// applies to each instruction from Address to just before the
	// Address of the next LineEntry.
    // Address 是编译器生成的机器指令的程序计数器值。
    // 这个 LineEntry 对应于从 Address 到下一个 LineEntry 的 Address 之前的每个指令。
	Address uint64

	// OpIndex is the index of an operation within a VLIW
	// instruction. The index of the first operation is 0. For
	// non-VLIW architectures, it will always be 0. Address and
	// OpIndex together form an operation pointer that can
	// reference any individual operation within the instruction
	// stream.
    // OpIndex 是 VLIW 指令中操作的索引。
    // 第一个操作的索引为 0。对于非 VLIW 架构，它将始终为 0。
    // Address 和 OpIndex 一起形成一个操作指针，可以引用指令流中的任何单个操作。
	OpIndex int

	// File is the source file corresponding to these
	// instructions.
    // File 是与这些指令对应的源文件。
	File *LineFile

	// Line is the source code line number corresponding to these
	// instructions. Lines are numbered beginning at 1. It may be
	// 0 if these instructions cannot be attributed to any source
	// line.
    // Line 是与这些指令对应的源代码行号。行号从 1 开始编号。
    // 如果这些指令无法归属于任何源代码行，则可能为 0。
	Line int

	// Column is the column number within the source line of these
	// instructions. Columns are numbered beginning at 1. It may
	// be 0 to indicate the "left edge" of the line.
    // Column 是这些指令在源代码行中的列号。列号从 1 开始编号。
    // 如果表示行的"左边缘"，则可能为 0。
	Column int

	// IsStmt indicates that Address is a recommended breakpoint
	// location, such as the beginning of a line, statement, or a
	// distinct subpart of a statement.
    // IsStmt 指示 Address 是否是推荐的断点位置，例如行的开头、语句的开头或语句的一个独立子部分。
	IsStmt bool

	// BasicBlock indicates that Address is the beginning of a
	// basic block.
    // BasicBlock 指示 Address 是否是基本块的开头。
	BasicBlock bool

	// PrologueEnd indicates that Address is one (of possibly
	// many) PCs where execution should be suspended for a
	// breakpoint on entry to the containing function.
	//
	// Added in DWARF 3.
    // PrologueEnd 指示 Address 是否是应在进入包含函数时中断执行的断点之一（可能是多个）。
	//
	// 在 DWARF 3 中添加。
	PrologueEnd bool

	// EpilogueBegin indicates that Address is one (of possibly
	// many) PCs where execution should be suspended for a
	// breakpoint on exit from this function.
	//
    // EpilogueBegin 指示 Address 是否是应在退出该函数时中断执行的断点之一（可能是多个）。
    //
    // 在 DWARF 3 中添加。
	EpilogueBegin bool

	// ISA is the instruction set architecture for these
	// instructions. Possible ISA values should be defined by the
	// applicable ABI specification.
	//
	// Added in DWARF 3.
    // ISA 是这些指令的指令集架构。可能的 ISA 值应由适用的 ABI 规范定义。
    //
    // 在 DWARF 3 中添加。
	ISA int

	// Discriminator is an arbitrary integer indicating the block
	// to which these instructions belong. It serves to
	// distinguish among multiple blocks that may all have with
	// the same source file, line, and column. Where only one
	// block exists for a given source position, it should be 0.
	//
    // Discriminator 是一个任意的整数，指示这些指令所属的块。
    // 它用于区分具有相同源文件、行号和列号的多个块。对于给定源位置只有一个块存在的情况下，它应为 0。
    //
    // 在 DWARF 3 中添加。
	Discriminator int    

	// EndSequence indicates that Address is the first byte after
	// the end of a sequence of target machine instructions. If it
	// is set, only this and the Address field are meaningful. A
	// line number table may contain information for multiple
	// potentially disjoint instruction sequences. The last entry
	// in a line table should always have EndSequence set.
    // EndSequence 指示 Address 是否是目标机器指令序列结束后的第一个字节。
    // 如果设置了它，那么只有这个字段和 Address 字段有意义。
    // 行号表可以包含多个可能不相交的指令序列的信息。
    // 行表中的最后一个条目应始终设置 EndSequence。
	EndSequence bool
}
```

A LineEntry is a row in a DWARF line table.

LineEntry 是 DWARF 行表中的一行。

### type LineFile  <- go1.5

``` go 
type LineFile struct {
	Name   string
	Mtime  uint64 // Implementation defined modification time, or 0 if unknown 实现定义的修改时间，如果未知则为 0
	Length int    // File length, or 0 if unknown 文件长度，如果未知则为 0
}
```

A LineFile is a source file referenced by a DWARF line table entry.

LineFile 是由 DWARF 行表条目引用的源文件。

### type LineReader  <- go1.5

``` go 
type LineReader struct {
	// contains filtered or unexported fields
}
```

A LineReader reads a sequence of LineEntry structures from a DWARF "line" section for a single compilation unit. LineEntries occur in order of increasing PC and each LineEntry gives metadata for the instructions from that LineEntry's PC to just before the next LineEntry's PC. The last entry will have its EndSequence field set.

LineReader 从单个编译单元的 DWARF "line" 节中读取一系列 LineEntry 结构。LineEntry 按照递增的 PC 顺序出现，每个 LineEntry 给出从该 LineEntry 的 PC 到下一个 LineEntry 的 PC 之前的指令的元数据。最后一个条目的 EndSequence 字段将被设置。

#### (*LineReader) Files  <- go1.14

``` go 
func (r *LineReader) Files() []*LineFile
```

Files returns the file name table of this compilation unit as of the current position in the line table. The file name table may be referenced from attributes in this compilation unit such as AttrDeclFile.

Files 返回行表中此编译单元的文件名表，根据当前位置。文件名表可以从该编译单元的属性（如 AttrDeclFile）中引用。

Entry 0 is always nil, since file index 0 represents "no file".

条目 0 始终为 nil，因为文件索引 0 表示"无文件"。

The file name table of a compilation unit is not fixed. Files returns the file table as of the current position in the line table. This may contain more entries than the file table at an earlier position in the line table, though existing entries never change.

编译单元的文件名表不是固定的。Files 返回当前位置的行表中的文件表。它可能包含比行表中较早位置的文件表更多的条目，尽管现有条目永远不会更改。

#### (*LineReader) Next  <- go1.5

``` go 
func (r *LineReader) Next(entry *LineEntry) error
```

Next sets *entry to the next row in this line table and moves to the next row. If there are no more entries and the line table is properly terminated, it returns io.EOF.

Next 将 *entry 设置为行表中的下一行，并移到下一行。如果没有更多的条目并且行表已正确终止，则返回 io.EOF。

Rows are always in order of increasing entry.Address, but entry.Line may go forward or backward.

行始终按照 entry.Address 递增的顺序排列，但 entry.Line 可能前进或后退。

#### (*LineReader) Reset  <- go1.5

``` go 
func (r *LineReader) Reset()
```

Reset repositions the line table reader at the beginning of the line table.

Reset 将行表读取器重新定位到行表的开头。

#### (*LineReader) Seek  <- go1.5

``` go 
func (r *LineReader) Seek(pos LineReaderPos)
```

Seek restores the line table reader to a position returned by Tell.

Seek 将行表读取器恢复到由 Tell 返回的位置。

The argument pos must have been returned by a call to Tell on this line table.

参数 pos 必须是对该行表调用 Tell 的返回值。

#### (*LineReader) SeekPC  <- go1.5

``` go 
func (r *LineReader) SeekPC(pc uint64, entry *LineEntry) error
```

SeekPC sets *entry to the LineEntry that includes pc and positions the reader on the next entry in the line table. If necessary, this will seek backwards to find pc.

SeekPC 将 *entry 设置为包含 pc 的 LineEntry，并将读取器定位到行表中的下一个条目。如果需要，它会向后搜索以找到 pc。

If pc is not covered by any entry in this line table, SeekPC returns ErrUnknownPC. In this case, *entry and the final seek position are unspecified.

如果 pc 在该行表中的任何条目中都没有覆盖，则 SeekPC 返回 ErrUnknownPC。在这种情况下，*entry 和最终搜索位置是未指定的。

Note that DWARF line tables only permit sequential, forward scans. Hence, in the worst case, this takes time linear in the size of the line table. If the caller wishes to do repeated fast PC lookups, it should build an appropriate index of the line table.

请注意，DWARF 行表仅允许顺序逐个扫描。因此，在最坏的情况下，这将花费与行表大小成线性关系的时间。如果调用者希望进行重复的快速 PC 查找，应构建适当的行表索引。

#### (*LineReader) Tell  <- go1.5

``` go 
func (r *LineReader) Tell() LineReaderPos
```

Tell returns the current position in the line table.

Tell 返回行表中的当前位置。

### type LineReaderPos  <- go1.5

``` go 
type LineReaderPos struct {
	// contains filtered or unexported fields
}
```

A LineReaderPos represents a position in a line table.

LineReaderPos 表示行表中的位置。

### type Offset 

``` go 
type Offset uint32
```

An Offset represents the location of an Entry within the DWARF info. (See Reader.Seek.)

Offset 表示 DWARF 信息中 Entry 的位置。（参见 Reader.Seek。）

### type PtrType 

``` go 
type PtrType struct {
	CommonType
	Type Type
}
```

A PtrType represents a pointer type.

PtrType 表示指针类型。

#### (*PtrType) String 

``` go 
func (t *PtrType) String() string
```

### type QualType 

``` go 
type QualType struct {
	CommonType
	Qual string
	Type Type
}
```

A QualType represents a type that has the C/C++ "const", "restrict", or "volatile" qualifier.

QualType 表示具有 C/C++ 的 "const"、"restrict" 或 "volatile" 修饰符的类型。

#### (*QualType) Size 

``` go 
func (t *QualType) Size() int64
```

#### (*QualType) String 

``` go 
func (t *QualType) String() string
```

### type Reader 

``` go 
type Reader struct {
	// contains filtered or unexported fields
}
```

A Reader allows reading Entry structures from a DWARF "info" section. The Entry structures are arranged in a tree. The Reader's Next function return successive entries from a pre-order traversal of the tree. If an entry has children, its Children field will be true, and the children follow, terminated by an Entry with Tag 0.

Reader 允许从 DWARF "info" 部分读取 Entry 结构。Entry 结构按树形排列。Reader 的 Next 函数按照树的前序遍历返回连续的条目。如果一个条目有子项，它的 Children 字段将为 true，后面跟着子项，以 Tag 为 0 的条目作为终止标志。

#### (*Reader) AddressSize  <- go1.5

``` go 
func (r *Reader) AddressSize() int
```

AddressSize returns the size in bytes of addresses in the current compilation unit.

AddressSize 返回当前编译单元中地址的字节大小。

#### (*Reader) ByteOrder  <- go1.14

``` go 
func (r *Reader) ByteOrder() binary.ByteOrder
```

ByteOrder returns the byte order in the current compilation unit.

ByteOrder 返回当前编译单元的字节顺序。

#### (*Reader) Next 

``` go 
func (r *Reader) Next() (*Entry, error)
```

Next reads the next entry from the encoded entry stream. It returns nil, nil when it reaches the end of the section. It returns an error if the current offset is invalid or the data at the offset cannot be decoded as a valid Entry.

Next 从编码的条目流中读取下一个条目。当到达该部分的结尾时，它返回 nil, nil。如果当前偏移量无效或偏移处的数据无法解码为有效的 Entry，则返回错误。

#### (*Reader) Seek 

``` go 
func (r *Reader) Seek(off Offset)
```

Seek positions the Reader at offset off in the encoded entry stream. Offset 0 can be used to denote the first entry.

Seek 将 Reader 定位到编码条目流中的偏移量 off 处。可以使用偏移量 0 表示第一个条目。

#### (*Reader) SeekPC  <- go1.7

``` go 
func (r *Reader) SeekPC(pc uint64) (*Entry, error)
```

SeekPC returns the Entry for the compilation unit that includes pc, and positions the reader to read the children of that unit. If pc is not covered by any unit, SeekPC returns ErrUnknownPC and the position of the reader is undefined.

SeekPC 返回包含 pc 的编译单元的 Entry，并将读取器定位到该单元的子项。如果 pc 不在任何单元范围内，则 SeekPC 返回 ErrUnknownPC，并且读取器的位置未定义。

Because compilation units can describe multiple regions of the executable, in the worst case SeekPC must search through all the ranges in all the compilation units. Each call to SeekPC starts the search at the compilation unit of the last call, so in general looking up a series of PCs will be faster if they are sorted. If the caller wishes to do repeated fast PC lookups, it should build an appropriate index using the Ranges method.

由于编译单元可以描述可执行文件的多个区域，在最坏的情况下，SeekPC 必须搜索所有编译单元中的所有范围。每次调用 SeekPC 都从上一次调用的编译单元开始搜索，因此通常情况下，如果排序了一系列的 pc，查找速度会更快。如果调用者希望进行重复的快速 PC 查找，应使用 Ranges 方法构建适当的索引。

#### (*Reader) SkipChildren 

``` go 
func (r *Reader) SkipChildren()
```

SkipChildren skips over the child entries associated with the last Entry returned by Next. If that Entry did not have children or Next has not been called, SkipChildren is a no-op.

SkipChildren 跳过与上次由 Next 返回的 Entry 相关联的子项。如果该 Entry 没有子项或者尚未调用 Next，则 SkipChildren 不执行任何操作。

### type StructField 

``` go 
type StructField struct {
	Name          string
	Type          Type
	ByteOffset    int64
	ByteSize      int64 // usually zero; use Type.Size() for normal fields 通常为零；对于普通字段，请使用 Type.Size()
	BitOffset     int64
	DataBitOffset int64
	BitSize       int64 // zero if not a bit field 如果不是位字段，则为零
}
```

A StructField represents a field in a struct, union, or C++ class type.

StructField 表示结构体、联合体或 C++ 类型中的字段。

#### Bit Fields 位字段

The BitSize, BitOffset, and DataBitOffset fields describe the bit size and offset of data members declared as bit fields in C/C++ struct/union/class types.

BitSize、BitOffset 和 DataBitOffset 字段描述在 C/C++ 结构体/联合体/类类型中声明为位字段的数据成员的位大小和偏移量。

BitSize is the number of bits in the bit field.

BitSize 是位字段的位数。

DataBitOffset, if non-zero, is the number of bits from the start of the enclosing entity (e.g. containing struct/class/union) to the start of the bit field. This corresponds to the DW_AT_data_bit_offset DWARF attribute that was introduced in DWARF 4.

如果 DataBitOffset 非零，则表示从包含实体（如包含的结构体/类/联合体）的开始到位字段的开始之间的位数。这对应于 DWARF 4 中引入的 DW_AT_data_bit_offset DWARF 属性。

BitOffset, if non-zero, is the number of bits between the most significant bit of the storage unit holding the bit field to the most significant bit of the bit field. Here "storage unit" is the type name before the bit field (for a field "unsigned x:17", the storage unit is "unsigned"). BitOffset values can vary depending on the endianness of the system. BitOffset corresponds to the DW_AT_bit_offset DWARF attribute that was deprecated in DWARF 4 and removed in DWARF 5.

如果 BitOffset 非零，则表示从存储位字段的存储单元的最高有效位到位字段的最高有效位之间的位数。这里的 "存储单元" 是位字段之前的类型名称（对于字段 "unsigned x:17"，存储单元是 "unsigned"）。BitOffset 的值可能因系统的字节顺序而异。BitOffset 对应于 DWARF 4 中已弃用并在 DWARF 5 中删除的 DW_AT_bit_offset DWARF 属性。

At most one of DataBitOffset and BitOffset will be non-zero; DataBitOffset/BitOffset will only be non-zero if BitSize is non-zero. Whether a C compiler uses one or the other will depend on compiler vintage and command line options.

DataBitOffset 和 BitOffset 中最多只有一个非零；只有在 BitSize 非零时，DataBitOffset/BitOffset 才会是非零值。C 编译器使用哪种偏移量取决于编译器的年代和命令行选项。

Here is an example of C/C++ bit field use, along with what to expect in terms of DWARF bit offset info. Consider this code:

下面是 C/C++ 位字段使用的示例，以及在 DWARF 位偏移信息方面的预期

```
struct S {
	int q;
	int j:5;
	int k:6;
	int m:5;
	int n:8;
} s;
```

For the code above, one would expect to see the following for DW_AT_bit_offset values (using GCC 8):

对于上面的代码，在使用 GCC 8 编译时，可以预期 DW_AT_bit_offset 值如下所示：

```
       Little   |     Big
       Endian   |    Endian
                |
"j":     27     |     0
"k":     21     |     5
"m":     16     |     11
"n":     8      |     16
```

Note that in the above the offsets are purely with respect to the containing storage unit for j/k/m/n -- these values won't vary based on the size of prior data members in the containing struct.

请注意，在上述示例中，j/k/m/n 的偏移量纯粹是相对于包含的存储单元的偏移量 - 这些值不会根据包含结构体中先前数据成员的大小而变化。

If the compiler emits DW_AT_data_bit_offset, the expected values would be:

如果编译器生成了 DW_AT_data_bit_offset，预期值将如下所示：

```
"j":     32
"k":     37
"m":     43
"n":     48
```

Here the value 32 for "j" reflects the fact that the bit field is preceded by other data members (recall that DW_AT_data_bit_offset values are relative to the start of the containing struct). Hence DW_AT_data_bit_offset values can be quite large for structs with many fields.

这里，"j" 的值为 32 反映了位字段之前存在其他数据成员的事实（请记住，DW_AT_data_bit_offset 值是相对于包含结构体的开头的）。因此，对于具有许多字段的结构体，DW_AT_data_bit_offset 值可能非常大。

DWARF also allow for the possibility of base types that have non-zero bit size and bit offset, so this information is also captured for base types, but it is worth noting that it is not possible to trigger this behavior using mainstream languages.

DWARF 还允许基本类型具有非零的位大小和位偏移量，因此也为基本类型捕获了这些信息，但值得注意的是，使用主流语言不可能触发此行为。

### type StructType 

``` go 
type StructType struct {
	CommonType
	StructName string
	Kind       string // "struct", "union", or "class".
	Field      []*StructField
	Incomplete bool // if true, struct, union, class is declared but not defined 如果为 true，则结构体、联合体、类被声明但未定义
}
```

A StructType represents a struct, union, or C++ class type.

StructType 表示结构体、联合体或 C++ 类型。

#### (*StructType) Defn 

``` go 
func (t *StructType) Defn() string
```

#### (*StructType) String 

``` go 
func (t *StructType) String() string
```

### type Tag 

``` go 
type Tag uint32
```

A Tag is the classification (the type) of an Entry.

Tag 是 Entry 的分类（类型）。

``` go 
const (
	TagArrayType              Tag = 0x01
	TagClassType              Tag = 0x02
	TagEntryPoint             Tag = 0x03
	TagEnumerationType        Tag = 0x04
	TagFormalParameter        Tag = 0x05
	TagImportedDeclaration    Tag = 0x08
	TagLabel                  Tag = 0x0A
	TagLexDwarfBlock          Tag = 0x0B
	TagMember                 Tag = 0x0D
	TagPointerType            Tag = 0x0F
	TagReferenceType          Tag = 0x10
	TagCompileUnit            Tag = 0x11
	TagStringType             Tag = 0x12
	TagStructType             Tag = 0x13
	TagSubroutineType         Tag = 0x15
	TagTypedef                Tag = 0x16
	TagUnionType              Tag = 0x17
	TagUnspecifiedParameters  Tag = 0x18
	TagVariant                Tag = 0x19
	TagCommonDwarfBlock       Tag = 0x1A
	TagCommonInclusion        Tag = 0x1B
	TagInheritance            Tag = 0x1C
	TagInlinedSubroutine      Tag = 0x1D
	TagModule                 Tag = 0x1E
	TagPtrToMemberType        Tag = 0x1F
	TagSetType                Tag = 0x20
	TagSubrangeType           Tag = 0x21
	TagWithStmt               Tag = 0x22
	TagAccessDeclaration      Tag = 0x23
	TagBaseType               Tag = 0x24
	TagCatchDwarfBlock        Tag = 0x25
	TagConstType              Tag = 0x26
	TagConstant               Tag = 0x27
	TagEnumerator             Tag = 0x28
	TagFileType               Tag = 0x29
	TagFriend                 Tag = 0x2A
	TagNamelist               Tag = 0x2B
	TagNamelistItem           Tag = 0x2C
	TagPackedType             Tag = 0x2D
	TagSubprogram             Tag = 0x2E
	TagTemplateTypeParameter  Tag = 0x2F
	TagTemplateValueParameter Tag = 0x30
	TagThrownType             Tag = 0x31
	TagTryDwarfBlock          Tag = 0x32
	TagVariantPart            Tag = 0x33
	TagVariable               Tag = 0x34
	TagVolatileType           Tag = 0x35
	// The following are new in DWARF 3.
    // 以下是 DWARF 3 中新增的。
	TagDwarfProcedure  Tag = 0x36
	TagRestrictType    Tag = 0x37
	TagInterfaceType   Tag = 0x38
	TagNamespace       Tag = 0x39
	TagImportedModule  Tag = 0x3A
	TagUnspecifiedType Tag = 0x3B
	TagPartialUnit     Tag = 0x3C
	TagImportedUnit    Tag = 0x3D
	TagMutableType     Tag = 0x3E // Later removed from DWARF. // 后来从 DWARF 中删除。
	TagCondition       Tag = 0x3F
	TagSharedType      Tag = 0x40
	// The following are new in DWARF 4.
    // 以下是 DWARF 4 中新增的。
	TagTypeUnit            Tag = 0x41
	TagRvalueReferenceType Tag = 0x42
	TagTemplateAlias       Tag = 0x43
	// The following are new in DWARF 5.
    // 以下是 DWARF 5 中新增的。
	TagCoarrayType       Tag = 0x44
	TagGenericSubrange   Tag = 0x45
	TagDynamicType       Tag = 0x46
	TagAtomicType        Tag = 0x47
	TagCallSite          Tag = 0x48
	TagCallSiteParameter Tag = 0x49
	TagSkeletonUnit      Tag = 0x4A
	TagImmutableType     Tag = 0x4B
)
```

#### (Tag) GoString 

``` go 
func (t Tag) GoString() string
```

#### (Tag) String 

``` go 
func (i Tag) String() string
```

### type Type 

``` go 
type Type interface {
	Common() *CommonType
	String() string
	Size() int64
}
```

A Type conventionally represents a pointer to any of the specific Type structures (CharType, StructType, etc.).

Type 通常表示指向特定 Type 结构（如 CharType、StructType 等）的指针。

### type TypedefType 

``` go 
type TypedefType struct {
	CommonType
	Type Type
}
```

A TypedefType represents a named type.

TypedefType 表示命名类型。

#### (*TypedefType) Size 

``` go 
func (t *TypedefType) Size() int64
```

#### (*TypedefType) String 

``` go 
func (t *TypedefType) String() string
```

### type UcharType 

``` go 
type UcharType struct {
	BasicType
}
```

A UcharType represents an unsigned character type.

UcharType 表示无符号字符类型。

### type UintType 

``` go 
type UintType struct {
	BasicType
}
```

A UintType represents an unsigned integer type.

UintType 表示无符号整数类型。

### type UnspecifiedType  <- go1.4

``` go 
type UnspecifiedType struct {
	BasicType
}
```

An UnspecifiedType represents an implicit, unknown, ambiguous or nonexistent type.

UnspecifiedType 表示隐式的、未知的、模糊的或不存在的类型。

### type UnsupportedType  <- go1.13

``` go 
type UnsupportedType struct {
	CommonType
	Tag Tag
}
```

An UnsupportedType is a placeholder returned in situations where we encounter a type that isn't supported.

UnsupportedType 是在遇到不支持的类型时返回的占位符。

#### (*UnsupportedType) String  <- go1.13

``` go 
func (t *UnsupportedType) String() string
```

### type VoidType 

``` go 
type VoidType struct {
	CommonType
}
```

A VoidType represents the C void type.

VoidType 表示 C 中的 void 类型。

#### (*VoidType) String 

``` go 
func (t *VoidType) String() string
```