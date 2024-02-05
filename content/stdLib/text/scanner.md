+++
title = "scanner"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/text/scanner@go1.21.3](https://pkg.go.dev/text/scanner@go1.21.3)

Package scanner provides a scanner and tokenizer for UTF-8-encoded text. It takes an io.Reader providing the source, which then can be tokenized through repeated calls to the Scan function. For compatibility with existing tools, the NUL character is not allowed. If the first character in the source is a UTF-8 encoded byte order mark (BOM), it is discarded.

​	`scanner`包提供了用于UTF-8编码文本的扫描器和标记器。它接受一个提供源代码的io.Reader，然后可以通过重复调用Scan函数来对其进行标记化。为了与现有工具兼容，不允许出现NUL字符。如果源代码中的第一个字符是UTF-8编码的字节顺序标记（BOM），它将被丢弃。

By default, a Scanner skips white space and Go comments and recognizes all literals as defined by the Go language specification. It may be customized to recognize only a subset of those literals and to recognize different identifier and white space characters.

​	默认情况下，Scanner会跳过空白字符和Go注释，并识别符合Go语言规范定义的所有字面量。它可以定制为仅识别这些字面量的子集，并识别不同的标识符和空白字符。

## Example
``` go 
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	const src = `
// This is scanned code.
if a > 10 {
	someParsable = text
}`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

}
//Output:

//example:3:1: if
//example:3:4: a
//example:3:6: >
//example:3:8: 10
//example:3:11: {
//example:4:2: someParsable
//example:4:15: =
//example:4:17: text
//example:5:1: }
```

## Example(IsIdentRune)
``` go 
package main

import (
	"fmt"
	"strings"
	"text/scanner"
	"unicode"
)

func main() {
	const src = "%var1 var2%"

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "default"

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

	fmt.Println()
	s.Init(strings.NewReader(src))
	s.Filename = "percent"

	// treat leading '%' as part of an identifier
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '%' && i == 0 || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0
	}

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Printf("%s: %s\n", s.Position, s.TokenText())
	}

}
//Output:

//default:1:1: %
//default:1:2: var1
//default:1:7: var2
//default:1:11: %

//percent:1:1: %var1
//percent:1:7: var2
//percent:1:11: %
```

## Example(Mode)
``` go 
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	const src = `
    // Comment begins at column 5.

This line should not be included in the output.

/*
This multiline comment
should be extracted in
its entirety.
*/
`

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "comments"
	s.Mode ^= scanner.SkipComments // don't skip comments

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		if strings.HasPrefix(txt, "//") || strings.HasPrefix(txt, "/*") {
			fmt.Printf("%s: %s\n", s.Position, txt)
		}
	}

}
//Output:

//comments:2:5: // Comment begins at column 5.
//comments:6:1: /*
//This multiline comment
//should be extracted in
//its entirety.
//*/
```

## Example (Whitespace)
``` go 
package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	// tab-separated values
	const src = `aa	ab	ac	ad
ba	bb	bc	bd
ca	cb	cc	cd
da	db	dc	dd`

	var (
		col, row int
		s        scanner.Scanner
		tsv      [4][4]string // large enough for example above
	)
	s.Init(strings.NewReader(src))
	s.Whitespace ^= 1<<'\t' | 1<<'\n' // don't skip tabs and new lines

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case '\n':
			row++
			col = 0
		case '\t':
			col++
		default:
			tsv[row][col] = s.TokenText()
		}
	}

	fmt.Print(tsv)

}
//Output:

//[[aa ab ac ad] [ba bb bc bd] [ca cb cc cd] [da db dc dd]]
```




## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/text/scanner/scanner.go;l=63)

``` go 
const (
	ScanIdents     = 1 << -Ident
	ScanInts       = 1 << -Int
	ScanFloats     = 1 << -Float // includes Ints and hexadecimal floats 包括Ints和十六进制浮点数
	ScanChars      = 1 << -Char
	ScanStrings    = 1 << -String
	ScanRawStrings = 1 << -RawString
	ScanComments   = 1 << -Comment
	SkipComments   = 1 << -skipComment // if set with ScanComments, comments become white space 如果与ScanComments一起设置，则注释将变为空白字符
	GoTokens       = ScanIdents | ScanFloats | ScanChars | ScanStrings | ScanRawStrings | ScanComments | SkipComments
)
```

Predefined mode bits to control recognition of tokens. For instance, to configure a Scanner such that it only recognizes (Go) identifiers, integers, and skips comments, set the Scanner's Mode field to:

预定义的模式位，用于控制对标记的识别。例如，要配置一个Scanner，使其仅识别（Go）标识符、整数，并跳过注释，可以将Scanner的Mode字段设置为：

```
ScanIdents | ScanInts | SkipComments
```

With the exceptions of comments, which are skipped if SkipComments is set, unrecognized tokens are not ignored. Instead, the scanner simply returns the respective individual characters (or possibly sub-tokens). For instance, if the mode is ScanIdents (not ScanStrings), the string "foo" is scanned as the token sequence '"' Ident '"'.

除了如果设置了SkipComments，则跳过注释，否则不会忽略无法识别的标记。相反，扫描器只会返回相应的单个字符（或可能是子标记）。例如，如果模式是ScanIdents（而不是ScanStrings），则字符串"foo"将被扫描为标记序列'"' Ident '"。

Use GoTokens to configure the Scanner such that it accepts all Go literal tokens including Go identifiers. Comments will be skipped.

使用GoTokens来配置Scanner，使其接受包括Go标识符在内的所有Go字面量标记。注释将被跳过。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/text/scanner/scanner.go;l=76)

``` go 
const (
	EOF = -(iota + 1)
	Ident
	Int
	Float
	Char
	String
	RawString
	Comment
)
```

The result of Scan is one of these tokens or a Unicode character.

Scan的结果是这些标记之一或Unicode字符。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/text/scanner/scanner.go;l=111)

``` go 
const GoWhitespace = 1<<'\t' | 1<<'\n' | 1<<'\r' | 1<<' '
```

GoWhitespace is the default value for the Scanner's Whitespace field. Its value selects Go's white space characters.

GoWhitespace是Scanner的Whitespace字段的默认值。它的值选择了Go

## 变量

This section is empty.

## 函数

### func TokenString 

``` go 
func TokenString(tok rune) string
```

TokenString returns a printable string for a token or Unicode character.

TokenString返回标记或Unicode字符的可打印字符串。

## 类型

### type Position 

``` go 
type Position struct {
	Filename string // filename, if any 文件名（如果有）
	Offset   int    // byte offset, starting at 0 字节偏移量，从0开始
	Line     int    // line number, starting at 1 行号，从1开始
	Column   int    // column number, starting at 1 (character count per line) 列号，从1开始（每行的字符计数）
}
```

Position is a value that represents a source position. A position is valid if Line > 0.

Position是表示源代码位置的值。如果Line > 0，则位置是有效的。

#### (*Position) IsValid 

``` go 
func (pos *Position) IsValid() bool
```

IsValid reports whether the position is valid.

IsValid报告位置是否有效。

#### (Position) String 

``` go 
func (pos Position) String() string
```

### type Scanner 

``` go 
type Scanner struct {

	// Error is called for each error encountered. If no Error
	// function is set, the error is reported to os.Stderr.
    // Error用于每次遇到错误时调用。如果未设置Error函数，错误将报告到os.Stderr。
	Error func(s *Scanner, msg string)

	// ErrorCount is incremented by one for each error encountered.
    // ErrorCount用于记录遇到的错误数。每遇到一个错误，ErrorCount加1。
	ErrorCount int

	// The Mode field controls which tokens are recognized. For instance,
	// to recognize Ints, set the ScanInts bit in Mode. The field may be
	// changed at any time.
    // Mode字段控制识别哪些标记。例如，要识别Ints，请在Mode中设置ScanInts位。可以随时更改该字段。
	Mode uint

	// The Whitespace field controls which characters are recognized
	// as white space. To recognize a character ch <= ' ' as white space,
	// set the ch'th bit in Whitespace (the Scanner's behavior is undefined
	// for values ch > ' '). The field may be changed at any time.
    // Whitespace字段控制哪些字符被识别为空白字符。要将字符ch <= ' '识别为空白字符，请在Whitespace中设置ch位（对于ch > ' '的值，Scanner的行为未定义）。可以随时更改该字段。
	Whitespace uint64

	// IsIdentRune is a predicate controlling the characters accepted
	// as the ith rune in an identifier. The set of valid characters
	// must not intersect with the set of white space characters.
	// If no IsIdentRune function is set, regular Go identifiers are
	// accepted instead. The field may be changed at any time.
    // IsIdentRune是一个断言函数，用于控制在标识符的第i个rune中接受的字符。有效字符集不能与空白字符集相交。如果未设置IsIdentRune函数，则接受常规的Go标识符。可以随时更改该字段。
	IsIdentRune func(ch rune, i int) bool

	// Start position of most recently scanned token; set by Scan.
	// Calling Init or Next invalidates the position (Line == 0).
	// The Filename field is always left untouched by the Scanner.
	// If an error is reported (via Error) and Position is invalid,
	// the scanner is not inside a token. Call Pos to obtain an error
	// position in that case, or to obtain the position immediately
	// after the most recently scanned token.
    // 最近扫描的标记的起始位置；由Scan设置。调用Init或Next将使位置无效（Line == 0）。文件名字段始终不受Scanner的影响。如果报告了错误（通过Error）并且Position无效，则扫描器不在标记内。在这种情况下，调用Pos获取错误位置，或获取最近扫描的标记之后的位置。
	Position
	// contains filtered or unexported fields
}
```

A Scanner implements reading of Unicode characters and tokens from an io.Reader.

Scanner实现了从io.Reader中读取Unicode字符和标记。

#### (*Scanner) Init 

``` go 
func (s *Scanner) Init(src io.Reader) *Scanner
```

Init initializes a Scanner with a new source and returns s. Error is set to nil, ErrorCount is set to 0, Mode is set to GoTokens, and Whitespace is set to GoWhitespace.

Init使用新的源代码初始化Scanner并返回s。Error设置为nil，ErrorCount设置为0，Mode设置为GoTokens，Whitespace设置为GoWhitespace。

#### (*Scanner) Next 

``` go 
func (s *Scanner) Next() rune
```

Next reads and returns the next Unicode character. It returns EOF at the end of the source. It reports a read error by calling s.Error, if not nil; otherwise it prints an error message to os.Stderr. Next does not update the Scanner's Position field; use Pos() to get the current position.

Next读取并返回下一个Unicode字符。在源代码的末尾返回EOF。如果不为nil，则通过调用s.Error报告读取错误；否则，它将打印错误消息到os.Stderr。Next不会更新Scanner的Position字段；使用Pos()获取当前位置。

#### (*Scanner) Peek 

``` go 
func (s *Scanner) Peek() rune
```

Peek returns the next Unicode character in the source without advancing the scanner. It returns EOF if the scanner's position is at the last character of the source.

Peek返回源代码中的下一个Unicode字符，但不推进扫描器。如果扫描器的位置在源代码的最后一个字符处，则返回EOF。

#### (*Scanner) Pos 

``` go 
func (s *Scanner) Pos() (pos Position)
```

Pos returns the position of the character immediately after the character or token returned by the last call to Next or Scan. Use the Scanner's Position field for the start position of the most recently scanned token.

Pos返回在上次调用Next或Scan时返回的字符或标记之后的字符的位置。使用Scanner的Position字段获取最近扫描的标记的起始位置。

#### (*Scanner) Scan 

``` go 
func (s *Scanner) Scan() rune
```

Scan reads the next token or Unicode character from source and returns it. It only recognizes tokens t for which the respective Mode bit (1<<-t) is set. It returns EOF at the end of the source. It reports scanner errors (read and token errors) by calling s.Error, if not nil; otherwise it prints an error message to os.Stderr.

Scan从源代码中读取下一个标记或Unicode字符并返回。它仅识别设置了相应Mode位（1<<-t）的标记t。在源代码的末尾返回EOF。如果不为nil，则通过调用s.Error报告扫描器错误（读取和标记错误）；否则，它将打印错误消息到os.Stderr。

#### (*Scanner) TokenText 

``` go 
func (s *Scanner) TokenText() string
```

TokenText returns the string corresponding to the most recently scanned token. Valid after calling Scan and in calls of Scanner.Error.

TokenText返回与最近扫描的标记对应的字符串。在调用Scan后和调用Scanner.Error时有效。