+++
title = "token"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/token@go1.23.0](https://pkg.go.dev/go/token@go1.23.0)

Package token defines constants representing the lexical tokens of the Go programming language and basic operations on tokens (printing, predicates).

​	token 包定义表示 Go 编程语言的词法标记的常量以及对标记的基本操作（打印、谓词）。

## Example(RetrievePositionInfo)

```go
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()

	const src = `package main

import "fmt"

import "go/token"

//line :1:5
type p = token.Pos

const bad = token.NoPos

//line fake.go:42:11
func ok(pos p) bool {
	return pos != bad
}

/*line :7:9*/func main() {
	fmt.Println(ok(bad) == bad.IsValid())
}
`

	f, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the location and kind of each declaration in f.
	for _, decl := range f.Decls {
		// Get the filename, line, and column back via the file set.
		// We get both the relative and absolute position.
		// The relative position is relative to the last line directive.
		// The absolute position is the exact position in the source.
		pos := decl.Pos()
		relPosition := fset.Position(pos)
		absPosition := fset.PositionFor(pos, false)

		// Either a FuncDecl or GenDecl, since we exit on error.
		kind := "func"
		if gen, ok := decl.(*ast.GenDecl); ok {
			kind = gen.Tok.String()
		}

		// If the relative and absolute positions differ, show both.
		fmtPosition := relPosition.String()
		if relPosition != absPosition {
			fmtPosition += "[" + absPosition.String() + "]"
		}

		fmt.Printf("%s: %s\n", fmtPosition, kind)
	}

}
Output:


main.go:3:1: import
main.go:5:1: import
main.go:1:5[main.go:8:1]: type
main.go:3:1[main.go:10:1]: const
fake.go:42:11[main.go:13:1]: func
fake.go:7:9[main.go:17:14]: func
```

## 常量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=257)

```go
const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)
```

A set of constants for precedence-based expression parsing. Non-operators have lowest precedence, followed by operators starting with precedence 1 up to unary operators. The highest precedence serves as “catch-all” precedence for selector, indexing, and other operator and delimiter tokens.

​	一组用于基于优先级的表达式解析的常量。非运算符具有最低优先级，其次是优先级从 1 开始的运算符，直到一元运算符。最高优先级用作选择器、索引和其他运算符和分隔符标记的“万能”优先级。

## 变量

This section is empty.

## 函数

### func IsExported <- go1.13

```go
func IsExported(name string) bool
```

IsExported reports whether name starts with an upper-case letter.

​	IsExported 报告名称是否以大写字母开头。

### func IsIdentifier <- go1.13

```go
func IsIdentifier(name string) bool
```

IsIdentifier reports whether name is a Go identifier, that is, a non-empty string made up of letters, digits, and underscores, where the first character is not a digit. Keywords are not identifiers.

​	IsIdentifier 报告 name 是否是 Go 标识符，即由字母、数字和下划线组成的非空字符串，其中第一个字符不是数字。关键字不是标识符。

### func IsKeyword <- go1.13

```go
func IsKeyword(name string) bool
```

IsKeyword reports whether name is a Go keyword, such as “func” or “return”.

​	IsKeyword 报告 name 是否是 Go 关键字，例如“func”或“return”。

## 类型

### type File

```go
type File struct {
	// contains filtered or unexported fields
}
```

A File is a handle for a file belonging to a FileSet. A File has a name, size, and line offset table.

​	File 是属于 FileSet 的文件的句柄。File 具有名称、大小和行偏移表。

#### (*File) AddLine

```go
func (f *File) AddLine(offset int)
```

AddLine adds the line offset for a new line. The line offset must be larger than the offset for the previous line and smaller than the file size; otherwise the line offset is ignored.

​	AddLine 添加新行的行偏移。行偏移必须大于前一行的偏移，并且小于文件大小；否则，将忽略行偏移。

#### (*File) AddLineColumnInfo <- go1.11

```go
func (f *File) AddLineColumnInfo(offset int, filename string, line, column int)
```

AddLineColumnInfo adds alternative file, line, and column number information for a given file offset. The offset must be larger than the offset for the previously added alternative line info and smaller than the file size; otherwise the information is ignored.

​	AddLineColumnInfo 为给定的文件偏移量添加备用文件、行和列号信息。偏移量必须大于先前添加的备用行信息偏移量，并且小于文件大小；否则，将忽略该信息。

AddLineColumnInfo is typically used to register alternative position information for line directives such as //line filename:line:column.

​	AddLineColumnInfo 通常用于为行指令（例如 //line filename:line:column）注册备用位置信息。

#### (*File) AddLineInfo

```go
func (f *File) AddLineInfo(offset int, filename string, line int)
```

AddLineInfo is like AddLineColumnInfo with a column = 1 argument. It is here for backward-compatibility for code prior to Go 1.11.

​	AddLineInfo 与 AddLineColumnInfo 类似，列 = 1 参数。它在此处用于 Go 1.11 之前的代码的向后兼容性。

#### (*File) Base

```go
func (f *File) Base() int
```

Base returns the base offset of file f as registered with AddFile.

​	Base 返回使用 AddFile 注册的文件 f 的基本偏移量。

#### (*File) Line

```go
func (f *File) Line(p Pos) int
```

Line returns the line number for the given file position p; p must be a Pos value in that file or NoPos.

​	Line 返回给定文件位置 p 的行号；p 必须是该文件中的 Pos 值或 NoPos。

#### (*File) LineCount

```go
func (f *File) LineCount() int
```

LineCount returns the number of lines in file f.

​	LineCount 返回文件 f 中的行数。

#### (*File) LineStart <- go1.12

```go
func (f *File) LineStart(line int) Pos
```

LineStart returns the Pos value of the start of the specified line. It ignores any alternative positions set using AddLineColumnInfo. LineStart panics if the 1-based line number is invalid.

​	LineStart 返回指定行的开始处的 Pos 值。它忽略使用 AddLineColumnInfo 设置的任何备用位置。如果基于 1 的行号无效，则 LineStart 会引发 panic。

#### (*File) Lines <-go1.21.0

```go
func (f *File) Lines() []int
```

Lines returns the effective line offset table of the form described by SetLines. Callers must not mutate the result.

​	Lines 返回 SetLines 描述的有效行偏移表。调用者不得改变结果。

#### (*File) MergeLine <- go1.2

```go
func (f *File) MergeLine(line int)
```

MergeLine merges a line with the following line. It is akin to replacing the newline character at the end of the line with a space (to not change the remaining offsets). To obtain the line number, consult e.g. Position.Line. MergeLine will panic if given an invalid line number.

​	MergeLine 将一行与下一行合并。它类似于用空格替换行尾的新行字符（不更改剩余的偏移量）。要获取行号，请咨询例如 Position.Line。如果给定无效的行号，MergeLine 将引发 panic。

#### (*File) Name

```go
func (f *File) Name() string
```

Name returns the file name of file f as registered with AddFile.

​	Name 返回使用 AddFile 注册的文件 f 的文件名。

#### (*File) Offset

```go
func (f *File) Offset(p Pos) int
```

Offset returns the offset for the given file position p; p must be a valid Pos value in that file. f.Offset(f.Pos(offset)) == offset.

​	Offset 返回给定文件位置 p 的偏移量；p 必须是该文件中的有效 Pos 值。f.Offset(f.Pos(offset)) == offset。

#### (*File) Pos

```go
func (f *File) Pos(offset int) Pos
```

Pos returns the Pos value for the given file offset; the offset must be <= f.Size(). f.Pos(f.Offset(p)) == p.

​	Pos 返回给定文件偏移量的 Pos 值；偏移量必须 <= f.Size()。f.Pos(f.Offset(p)) == p。

#### (*File) Position

```go
func (f *File) Position(p Pos) (pos Position)
```

Position returns the Position value for the given file position p. Calling f.Position(p) is equivalent to calling f.PositionFor(p, true).

​	Position 返回给定文件位置 p 的 Position 值。调用 f.Position(p) 等同于调用 f.PositionFor(p, true)。

#### (*File) PositionFor <- go1.4 

```go
func (f *File) PositionFor(p Pos, adjusted bool) (pos Position)
```

PositionFor returns the Position value for the given file position p. If adjusted is set, the position may be adjusted by position-altering //line comments; otherwise those comments are ignored. p must be a Pos value in f or NoPos.

​	PositionFor 返回给定文件位置 p 的 Position 值。如果设置了 adjusted，则位置可能会因改变位置的 //line 注释而调整；否则，将忽略这些注释。p 必须是 f 中的 Pos 值或 NoPos。

#### (*File) SetLines 

```go
func (f *File) SetLines(lines []int) bool
```

SetLines sets the line offsets for a file and reports whether it succeeded. The line offsets are the offsets of the first character of each line; for instance for the content “ab\nc\n” the line offsets are {0, 3}. An empty file has an empty line offset table. Each line offset must be larger than the offset for the previous line and smaller than the file size; otherwise SetLines fails and returns false. Callers must not mutate the provided slice after SetLines returns.

​	SetLines 设置文件的行偏移量并报告是否成功。行偏移量是每行的第一个字符的偏移量；例如，对于内容“ab\nc\n”，行偏移量为 {0, 3}。空文件具有空行偏移量表。每个行偏移量必须大于前一行的偏移量且小于文件大小；否则，SetLines 将失败并返回 false。调用者在 SetLines 返回后不得改变提供的切片。

#### (*File) SetLinesForContent 

```go
func (f *File) SetLinesForContent(content []byte)
```

SetLinesForContent sets the line offsets for the given file content. It ignores position-altering //line comments.

​	SetLinesForContent 为给定的文件内容设置行偏移量。它忽略改变位置的 //line 注释。

#### (*File) Size 

```go
func (f *File) Size() int
```

Size returns the size of file f as registered with AddFile.

​	Size 返回使用 AddFile 注册的文件 f 的大小。

### type FileSet 

```go
type FileSet struct {
	// contains filtered or unexported fields
}
```

A FileSet represents a set of source files. Methods of file sets are synchronized; multiple goroutines may invoke them concurrently.

​	FileSet 表示一组源文件。文件集的方法是同步的；多个 goroutine 可以同时调用它们。

The byte offsets for each file in a file set are mapped into distinct (integer) intervals, one interval [base, base+size] per file. Base represents the first byte in the file, and size is the corresponding file size. A Pos value is a value in such an interval. By determining the interval a Pos value belongs to, the file, its file base, and thus the byte offset (position) the Pos value is representing can be computed.

​	文件集中每个文件的字节偏移量映射到不同的（整数）区间，每个文件一个区间 [base, base+size]。Base 表示文件中的第一个字节，size 是相应的文件大小。Pos 值是此类区间中的一个值。通过确定 Pos 值所属的区间，可以计算出文件、其文件基数，以及 Pos 值表示的字节偏移量（位置）。

When adding a new file, a file base must be provided. That can be any integer value that is past the end of any interval of any file already in the file set. For convenience, FileSet.Base provides such a value, which is simply the end of the Pos interval of the most recently added file, plus one. Unless there is a need to extend an interval later, using the FileSet.Base should be used as argument for FileSet.AddFile.

​	添加新文件时，必须提供文件基数。该基数可以是任何整数值，该值必须位于文件集中任何文件的任何区间的末尾之后。为了方便起见，FileSet.Base 提供了这样的值，该值只是最近添加的文件的 Pos 区间的末尾加一。除非以后需要扩展区间，否则应将 FileSet.Base 用作 FileSet.AddFile 的参数。

A File may be removed from a FileSet when it is no longer needed. This may reduce memory usage in a long-running application.

​	当不再需要文件时，可以从 FileSet 中将其删除。这可能会减少长时间运行的应用程序中的内存使用量。

#### func NewFileSet

```go
func NewFileSet() *FileSet
```

NewFileSet creates a new file set.

​	NewFileSet 创建新的文件集。

#### (*FileSet) AddFile

```go
func (s *FileSet) AddFile(filename string, base, size int) *File
```

AddFile adds a new file with a given filename, base offset, and file size to the file set s and returns the file. Multiple files may have the same name. The base offset must not be smaller than the FileSet’s Base(), and size must not be negative. As a special case, if a negative base is provided, the current value of the FileSet’s Base() is used instead.

​	AddFile 将一个具有给定文件名、基本偏移量和文件大小的新文件添加到文件集 s 并返回该文件。多个文件可能具有相同名称。基本偏移量不得小于 FileSet 的 Base()，并且大小不得为负数。作为特例，如果提供了负基本偏移量，则使用 FileSet 的 Base() 的当前值代替。

Adding the file will set the file set’s Base() value to base + size + 1 as the minimum base value for the next file. The following relationship exists between a Pos value p for a given file offset offs:

​	添加文件会将文件集的 Base() 值设置为 base + size + 1，作为下一个文件的最小基本值。给定文件偏移量 offs 的 Pos 值 p 与以下关系存在：

```
int(p) = base + offs
```

with offs in the range [0, size] and thus p in the range [base, base+size]. For convenience, File.Pos may be used to create file-specific position values from a file offset.

​	offs 在范围 [0, size] 内，因此 p 在范围 [base, base+size] 内。为了方便起见，可以使用 File.Pos 从文件偏移量创建特定于文件的定位值。

#### (*FileSet) Base

```go
func (s *FileSet) Base() int
```

Base returns the minimum base offset that must be provided to AddFile when adding the next file.

​	Base 返回在添加下一个文件时提供给 AddFile 的最小基本偏移量。

#### (*FileSet) File

```go
func (s *FileSet) File(p Pos) (f *File)
```

File returns the file that contains the position p. If no such file is found (for instance for p == NoPos), the result is nil.

​	File 返回包含定位 p 的文件。如果找不到此类文件（例如对于 p == NoPos），则结果为 nil。

#### (*FileSet) Iterate

```go
func (s *FileSet) Iterate(f func(*File) bool)
```

Iterate calls f for the files in the file set in the order they were added until f returns false.

​	Iterate 按添加顺序对文件集中的文件调用 f，直到 f 返回 false。

#### (*FileSet) Position

```go
func (s *FileSet) Position(p Pos) (pos Position)
```

Position converts a Pos p in the fileset into a Position value. Calling s.Position(p) is equivalent to calling s.PositionFor(p, true).

​	Position 将 fileset 中的 Pos p 转换为 Position 值。调用 s.Position(p) 等同于调用 s.PositionFor(p, true)。

#### (*FileSet) PositionFor <- go1.4

```go
func (s *FileSet) PositionFor(p Pos, adjusted bool) (pos Position)
```

PositionFor converts a Pos p in the fileset into a Position value. If adjusted is set, the position may be adjusted by position-altering //line comments; otherwise those comments are ignored. p must be a Pos value in s or NoPos.

​	PositionFor 将 fileset 中的 Pos p 转换为 Position 值。如果设置了 adjusted，则位置可能会因位置更改的 //line 注释而调整；否则，将忽略这些注释。p 必须是 s 中的 Pos 值或 NoPos。

#### (*FileSet) Read

```go
func (s *FileSet) Read(decode func(any) error) error
```

Read calls decode to deserialize a file set into s; s must not be nil.

​	Read 调用 decode 将文件集反序列化为 s；s 不能为 nil。

#### (*FileSet) RemoveFile <- go1.20

```go
func (s *FileSet) RemoveFile(file *File)
```

RemoveFile removes a file from the FileSet so that subsequent queries for its Pos interval yield a negative result. This reduces the memory usage of a long-lived FileSet that encounters an unbounded stream of files.

​	RemoveFile 从 FileSet 中删除一个文件，以便对它的 Pos 区间的后续查询产生负结果。这减少了遇到无限文件流的长寿命 FileSet 的内存使用量。

Removing a file that does not belong to the set has no effect.

​	删除不属于该集合的文件不会产生任何效果。

#### (*FileSet) Write

```go
func (s *FileSet) Write(encode func(any) error) error
```

Write calls encode to serialize the file set s.

​	Write 调用 encode 将文件集 s 序列化。

### type Pos

```go
type Pos int
```

Pos is a compact encoding of a source position within a file set. It can be converted into a Position for a more convenient, but much larger, representation.

​	Pos 是文件集中源位置的紧凑编码。它可以转换为 Position 以获得更方便但更大的表示形式。

The Pos value for a given file is a number in the range [base, base+size], where base and size are specified when a file is added to the file set. The difference between a Pos value and the corresponding file base corresponds to the byte offset of that position (represented by the Pos value) from the beginning of the file. Thus, the file base offset is the Pos value representing the first byte in the file.

​	给定文件的 Pos 值是范围 [base, base+size] 中的一个数字，其中 base 和 size 在将文件添加到文件集时指定。Pos 值与相应文件 base 之间的差值对应于该位置（由 Pos 值表示）与文件开头的字节偏移量。因此，文件 base 偏移量是表示文件第一个字节的 Pos 值。

To create the Pos value for a specific source offset (measured in bytes), first add the respective file to the current file set using FileSet.AddFile and then call File.Pos(offset) for that file. Given a Pos value p for a specific file set fset, the corresponding Position value is obtained by calling fset.Position(p).

​	要为特定源偏移量（以字节为单位测量）创建 Pos 值，首先使用 FileSet.AddFile 将相应的文件添加到当前文件集，然后为该文件调用 File.Pos(offset)。给定特定文件集 fset 的 Pos 值 p，可以通过调用 fset.Position(p) 获得相应的 Position 值。

Pos values can be compared directly with the usual comparison operators: If two Pos values p and q are in the same file, comparing p and q is equivalent to comparing the respective source file offsets. If p and q are in different files, p < q is true if the file implied by p was added to the respective file set before the file implied by q.

​	Pos 值可以直接与通常的比较运算符进行比较：如果两个 Pos 值 p 和 q 在同一个文件中，则比较 p 和 q 等同于比较各自的源文件偏移量。如果 p 和 q 在不同的文件中，则如果 p 暗示的文件在 q 暗示的文件之前添加到各自的文件集中，则 p < q 为真。

```go
const NoPos Pos = 0
```

The zero value for Pos is NoPos; there is no file and line information associated with it, and NoPos.IsValid() is false. NoPos is always smaller than any other Pos value. The corresponding Position value for NoPos is the zero value for Position.

​	Pos 的零值为 NoPos；没有与之关联的文件和行信息，并且 NoPos.IsValid() 为 false。NoPos 始终小于任何其他 Pos 值。NoPos 的相应 Position 值是 Position 的零值。

#### (Pos) IsValid

```go
func (p Pos) IsValid() bool
```

IsValid reports whether the position is valid.

​	IsValid 报告位置是否有效。

### type Position

```go
type Position struct {
	Filename string // filename, if any
	Offset   int    // offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (byte count)
}
```

Position describes an arbitrary source position including the file, line, and column location. A Position is valid if the line number is > 0.

​	Position 描述了一个任意源位置，包括文件、行和列位置。如果行号 > 0，则 Position 有效。

#### (*Position) IsValid

```go
func (pos *Position) IsValid() bool
```

IsValid reports whether the position is valid.

​	IsValid 报告位置是否有效。

#### (Position) String

```go
func (pos Position) String() string
```

String returns a string in one of several forms:

​	String 以以下几种形式之一返回字符串：

```
file:line:column    valid position with file name
file:line           valid position with file name but no column (column == 0)
line:column         valid position without file name
line                valid position without file name and no column (column == 0)
file                invalid position with file name
-                   invalid position without file name
```

### type Token

```go
type Token int
```

Token is the set of lexical tokens of the Go programming language.

​	Token 是 Go 编程语言的词法标记集。

```go
const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	COMMENT

	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT  // main
	INT    // 12345
	FLOAT  // 123.45
	IMAG   // 123.45i
	CHAR   // 'a'
	STRING // "abc"

	// Operators and delimiters
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	REM // %

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	REM_ASSIGN // %=

	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=

	LAND  // &&
	LOR   // ||
	ARROW // <-
	INC   // ++
	DEC   // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ      // !=
	LEQ      // <=
	GEQ      // >=
	DEFINE   // :=
	ELLIPSIS // ...

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :

	// Keywords
	BREAK
	CASE
	CHAN
	CONST
	CONTINUE

	DEFAULT
	DEFER
	ELSE
	FALLTHROUGH
	FOR

	FUNC
	GO
	GOTO
	IF
	IMPORT

	INTERFACE
	MAP
	PACKAGE
	RANGE
	RETURN

	SELECT
	STRUCT
	SWITCH
	TYPE
	VAR

	// additional tokens, handled in an ad-hoc manner
	TILDE
)
```

The list of tokens.

​	标记列表。

#### func Lookup

```go
func Lookup(ident string) Token
```

Lookup maps an identifier to its keyword token or IDENT (if not a keyword).

​	Lookup 将标识符映射到其关键字标记或 IDENT（如果不是关键字）。

#### (Token) IsKeyword

```go
func (tok Token) IsKeyword() bool
```

IsKeyword returns true for tokens corresponding to keywords; it returns false otherwise.

​	IsKeyword 对对应于关键字的标记返回 true；否则返回 false。

#### (Token) IsLiteral

```go
func (tok Token) IsLiteral() bool
```

IsLiteral returns true for tokens corresponding to identifiers and basic type literals; it returns false otherwise.

​	IsLiteral 对对应于标识符和基本类型文字的标记返回 true；否则返回 false。

#### (Token) IsOperator

```go
func (tok Token) IsOperator() bool
```

IsOperator returns true for tokens corresponding to operators and delimiters; it returns false otherwise.

​	IsOperator 对对应于运算符和分隔符的标记返回 true；否则返回 false。

#### (Token) Precedence

```go
func (op Token) Precedence() int
```

Precedence returns the operator precedence of the binary operator op. If op is not a binary operator, the result is LowestPrecedence.

​	优先级返回二元运算符 op 的运算符优先级。如果 op 不是二元运算符，则结果为 LowestPrecedence。

#### (Token) String 

```go
func (tok Token) String() string
```

String returns the string corresponding to the token tok. For operators, delimiters, and keywords the string is the actual token character sequence (e.g., for the token ADD, the string is “+”). For all other tokens the string corresponds to the token constant name (e.g. for the token IDENT, the string is “IDENT”).

​	String 返回与令牌 tok 相对应的字符串。对于运算符、分隔符和关键字，字符串是实际的令牌字符序列（例如，对于令牌 ADD，字符串是“+”）。对于所有其他令牌，字符串对应于令牌常量名称（例如，对于令牌 IDENT，字符串是“IDENT”）。