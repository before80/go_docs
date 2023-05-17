+++
title = "token"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# token

https://pkg.go.dev/go/token@go1.20.1



Package token defines constants representing the lexical tokens of the Go programming language and basic operations on tokens (printing, predicates).

##### Example
``` go linenums="1"
```







## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=257)

``` go linenums="1"
const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)
```

A set of constants for precedence-based expression parsing. Non-operators have lowest precedence, followed by operators starting with precedence 1 up to unary operators. The highest precedence serves as "catch-all" precedence for selector, indexing, and other operator and delimiter tokens.

## 变量

This section is empty.

## 函数

#### func [IsExported](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=316)  <- go1.13

``` go linenums="1"
func IsExported(name string) bool
```

IsExported reports whether name starts with an upper-case letter.

#### func [IsIdentifier](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=331)  <- go1.13

``` go linenums="1"
func IsIdentifier(name string) bool
```

IsIdentifier reports whether name is a Go identifier, that is, a non-empty string made up of letters, digits, and underscores, where the first character is not a digit. Keywords are not identifiers.

#### func [IsKeyword](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=322)  <- go1.13

``` go linenums="1"
func IsKeyword(name string) bool
```

IsKeyword reports whether name is a Go keyword, such as "func" or "return".

## 类型

### type [File](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=95) 

``` go linenums="1"
type File struct {
	// contains filtered or unexported fields
}
```

A File is a handle for a file belonging to a FileSet. A File has a name, size, and line offset table.

#### (*File) [AddLine](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=132) 

``` go linenums="1"
func (f *File) AddLine(offset int)
```

AddLine adds the line offset for a new line. The line offset must be larger than the offset for the previous line and smaller than the file size; otherwise the line offset is ignored.

#### (*File) [AddLineColumnInfo](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=246)  <- go1.11

``` go linenums="1"
func (f *File) AddLineColumnInfo(offset int, filename string, line, column int)
```

AddLineColumnInfo adds alternative file, line, and column number information for a given file offset. The offset must be larger than the offset for the previously added alternative line info and smaller than the file size; otherwise the information is ignored.

AddLineColumnInfo is typically used to register alternative position information for line directives such as //line filename:line:column.

#### (*File) [AddLineInfo](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=234) 

``` go linenums="1"
func (f *File) AddLineInfo(offset int, filename string, line int)
```

AddLineInfo is like AddLineColumnInfo with a column = 1 argument. It is here for backward-compatibility for code prior to Go 1.11.

#### (*File) [Base](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=112) 

``` go linenums="1"
func (f *File) Base() int
```

Base returns the base offset of file f as registered with AddFile.

#### (*File) [Line](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=276) 

``` go linenums="1"
func (f *File) Line(p Pos) int
```

Line returns the line number for the given file position p; p must be a Pos value in that file or NoPos.

#### (*File) [LineCount](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=122) 

``` go linenums="1"
func (f *File) LineCount() int
```

LineCount returns the number of lines in file f.

#### (*File) [LineStart](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=210)  <- go1.12

``` go linenums="1"
func (f *File) LineStart(line int) Pos
```

LineStart returns the Pos value of the start of the specified line. It ignores any alternative positions set using AddLineColumnInfo. LineStart panics if the 1-based line number is invalid.

#### (*File) [MergeLine](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=144)  <- go1.2

``` go linenums="1"
func (f *File) MergeLine(line int)
```

MergeLine merges a line with the following line. It is akin to replacing the newline character at the end of the line with a space (to not change the remaining offsets). To obtain the line number, consult e.g. Position.Line. MergeLine will panic if given an invalid line number.

#### (*File) [Name](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=107) 

``` go linenums="1"
func (f *File) Name() string
```

Name returns the file name of file f as registered with AddFile.

#### (*File) [Offset](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=267) 

``` go linenums="1"
func (f *File) Offset(p Pos) int
```

Offset returns the offset for the given file position p; p must be a valid Pos value in that file. f.Offset(f.Pos(offset)) == offset.

#### (*File) [Pos](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=257) 

``` go linenums="1"
func (f *File) Pos(offset int) Pos
```

Pos returns the Pos value for the given file offset; the offset must be <= f.Size(). f.Pos(f.Offset(p)) == p.

#### (*File) [Position](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=345) 

``` go linenums="1"
func (f *File) Position(p Pos) (pos Position)
```

Position returns the Position value for the given file position p. Calling f.Position(p) is equivalent to calling f.PositionFor(p, true).

#### (*File) [PositionFor](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=333)  <- go1.4

``` go linenums="1"
func (f *File) PositionFor(p Pos, adjusted bool) (pos Position)
```

PositionFor returns the Position value for the given file position p. If adjusted is set, the position may be adjusted by position-altering //line comments; otherwise those comments are ignored. p must be a Pos value in f or NoPos.

#### (*File) [SetLines](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=170) 

``` go linenums="1"
func (f *File) SetLines(lines []int) bool
```

SetLines sets the line offsets for a file and reports whether it succeeded. The line offsets are the offsets of the first character of each line; for instance for the content "ab\nc\n" the line offsets are {0, 3}. An empty file has an empty line offset table. Each line offset must be larger than the offset for the previous line and smaller than the file size; otherwise SetLines fails and returns false. Callers must not mutate the provided slice after SetLines returns.

#### (*File) [SetLinesForContent](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=188) 

``` go linenums="1"
func (f *File) SetLinesForContent(content []byte)
```

SetLinesForContent sets the line offsets for the given file content. It ignores position-altering //line comments.

#### (*File) [Size](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=117) 

``` go linenums="1"
func (f *File) Size() int
```

Size returns the size of file f as registered with AddFile.

### type [FileSet](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=374) 

``` go linenums="1"
type FileSet struct {
	// contains filtered or unexported fields
}
```

A FileSet represents a set of source files. Methods of file sets are synchronized; multiple goroutines may invoke them concurrently.

The byte offsets for each file in a file set are mapped into distinct (integer) intervals, one interval [base, base+size] per file. Base represents the first byte in the file, and size is the corresponding file size. A Pos value is a value in such an interval. By determining the interval a Pos value belongs to, the file, its file base, and thus the byte offset (position) the Pos value is representing can be computed.

When adding a new file, a file base must be provided. That can be any integer value that is past the end of any interval of any file already in the file set. For convenience, FileSet.Base provides such a value, which is simply the end of the Pos interval of the most recently added file, plus one. Unless there is a need to extend an interval later, using the FileSet.Base should be used as argument for FileSet.AddFile.

A File may be removed from a FileSet when it is no longer needed. This may reduce memory usage in a long-running application.

#### func [NewFileSet](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=382) 

``` go linenums="1"
func NewFileSet() *FileSet
```

NewFileSet creates a new file set.

#### (*FileSet) [AddFile](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=413) 

``` go linenums="1"
func (s *FileSet) AddFile(filename string, base, size int) *File
```

AddFile adds a new file with a given filename, base offset, and file size to the file set s and returns the file. Multiple files may have the same name. The base offset must not be smaller than the FileSet's Base(), and size must not be negative. As a special case, if a negative base is provided, the current value of the FileSet's Base() is used instead.

Adding the file will set the file set's Base() value to base + size + 1 as the minimum base value for the next file. The following relationship exists between a Pos value p for a given file offset offs:

```
int(p) = base + offs
```

with offs in the range [0, size] and thus p in the range [base, base+size]. For convenience, File.Pos may be used to create file-specific position values from a file offset.

#### (*FileSet) [Base](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=390) 

``` go linenums="1"
func (s *FileSet) Base() int
```

Base returns the minimum base offset that must be provided to AddFile when adding the next file.

#### (*FileSet) [File](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=506) 

``` go linenums="1"
func (s *FileSet) File(p Pos) (f *File)
```

File returns the file that contains the position p. If no such file is found (for instance for p == NoPos), the result is nil.

#### (*FileSet) [Iterate](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=462) 

``` go linenums="1"
func (s *FileSet) Iterate(f func(*File) bool)
```

Iterate calls f for the files in the file set in the order they were added until f returns false.

#### (*FileSet) [Position](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=528) 

``` go linenums="1"
func (s *FileSet) Position(p Pos) (pos Position)
```

Position converts a Pos p in the fileset into a Position value. Calling s.Position(p) is equivalent to calling s.PositionFor(p, true).

#### (*FileSet) [PositionFor](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=517)  <- go1.4

``` go linenums="1"
func (s *FileSet) PositionFor(p Pos, adjusted bool) (pos Position)
```

PositionFor converts a Pos p in the fileset into a Position value. If adjusted is set, the position may be adjusted by position-altering //line comments; otherwise those comments are ignored. p must be a Pos value in s or NoPos.

#### (*FileSet) [Read](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/serialize.go;l=22) 

``` go linenums="1"
func (s *FileSet) Read(decode func(any) error) error
```

Read calls decode to deserialize a file set into s; s must not be nil.

#### (*FileSet) [RemoveFile](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=447)  <- go1.20

``` go linenums="1"
func (s *FileSet) RemoveFile(file *File)
```

RemoveFile removes a file from the FileSet so that subsequent queries for its Pos interval yield a negative result. This reduces the memory usage of a long-lived FileSet that encounters an unbounded stream of files.

Removing a file that does not belong to the set has no effect.

#### (*FileSet) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/serialize.go;l=49) 

``` go linenums="1"
func (s *FileSet) Write(encode func(any) error) error
```

Write calls encode to serialize the file set s.

### type [Pos](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=77) 

``` go linenums="1"
type Pos int
```

Pos is a compact encoding of a source position within a file set. It can be converted into a Position for a more convenient, but much larger, representation.

The Pos value for a given file is a number in the range [base, base+size], where base and size are specified when a file is added to the file set. The difference between a Pos value and the corresponding file base corresponds to the byte offset of that position (represented by the Pos value) from the beginning of the file. Thus, the file base offset is the Pos value representing the first byte in the file.

To create the Pos value for a specific source offset (measured in bytes), first add the respective file to the current file set using FileSet.AddFile and then call File.Pos(offset) for that file. Given a Pos value p for a specific file set fset, the corresponding Position value is obtained by calling fset.Position(p).

Pos values can be compared directly with the usual comparison operators: If two Pos values p and q are in the same file, comparing p and q is equivalent to comparing the respective source file offsets. If p and q are in different files, p < q is true if the file implied by p was added to the respective file set before the file implied by q.

``` go linenums="1"
const NoPos Pos = 0
```

The zero value for Pos is NoPos; there is no file and line information associated with it, and NoPos.IsValid() is false. NoPos is always smaller than any other Pos value. The corresponding Position value for NoPos is the zero value for Position.

#### (Pos) [IsValid](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=86) 

``` go linenums="1"
func (p Pos) IsValid() bool
```

IsValid reports whether the position is valid.

### type [Position](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=20) 

``` go linenums="1"
type Position struct {
	Filename string // filename, if any
	Offset   int    // offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (byte count)
}
```

Position describes an arbitrary source position including the file, line, and column location. A Position is valid if the line number is > 0.

#### (*Position) [IsValid](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=28) 

``` go linenums="1"
func (pos *Position) IsValid() bool
```

IsValid reports whether the position is valid.

#### (Position) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/position.go;l=38) 

``` go linenums="1"
func (pos Position) String() string
```

String returns a string in one of several forms:

```
file:line:column    valid position with file name
file:line           valid position with file name but no column (column == 0)
line:column         valid position without file name
line                valid position without file name and no column (column == 0)
file                invalid position with file name
-                   invalid position without file name
```

### type [Token](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=16) 

``` go linenums="1"
type Token int
```

Token is the set of lexical tokens of the Go programming language.

``` go linenums="1"
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

#### func [Lookup](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=292) 

``` go linenums="1"
func Lookup(ident string) Token
```

Lookup maps an identifier to its keyword token or IDENT (if not a keyword).

#### (Token) [IsKeyword](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=313) 

``` go linenums="1"
func (tok Token) IsKeyword() bool
```

IsKeyword returns true for tokens corresponding to keywords; it returns false otherwise.

#### (Token) [IsLiteral](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=303) 

``` go linenums="1"
func (tok Token) IsLiteral() bool
```

IsLiteral returns true for tokens corresponding to identifiers and basic type literals; it returns false otherwise.

#### (Token) [IsOperator](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=307) 

``` go linenums="1"
func (tok Token) IsOperator() bool
```

IsOperator returns true for tokens corresponding to operators and delimiters; it returns false otherwise.

#### (Token) [Precedence](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=266) 

``` go linenums="1"
func (op Token) Precedence() int
```

Precedence returns the operator precedence of the binary operator op. If op is not a binary operator, the result is LowestPrecedence.

#### (Token) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/go/token/token.go;l=241) 

``` go linenums="1"
func (tok Token) String() string
```

String returns the string corresponding to the token tok. For operators, delimiters, and keywords the string is the actual token character sequence (e.g., for the token ADD, the string is "+"). For all other tokens the string corresponds to the token constant name (e.g. for the token IDENT, the string is "IDENT").