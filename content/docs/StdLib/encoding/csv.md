+++
title = "csv"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# csv

https://pkg.go.dev/encoding/csv@go1.20.1



Package csv reads and writes comma-separated values (CSV) files. There are many kinds of CSV files; this package supports the format described in [RFC 4180](https://rfc-editor.org/rfc/rfc4180.html).

csv包读取和写入逗号分隔的值(CSV)文件。有许多种CSV文件；这个包支持RFC 4180中描述的格式。

A csv file contains zero or more records of one or more fields per record. Each record is separated by the newline character. The final record may optionally be followed by a newline character.

一个csv文件包含零个或多个记录，每个记录有一个或多个字段。每条记录由换行符分隔。最后一条记录后面可以选择换行符。

```
field1,field2,field3
```

White space is considered part of a field.

白色空间被认为是字段的一部分。

Carriage returns before newline characters are silently removed.

换行符之前的回车符会被默默地删除。

Blank lines are ignored. A line with only whitespace characters (excluding the ending newline character) is not considered a blank line.

空白行被忽略。只有空白字符的行(不包括结尾的换行字符)不被视为空行。

Fields which start and stop with the quote character " are called quoted-fields. The beginning and ending quote are not part of the field.

以引号字符""开始和结束的字段被称为引号字段。开始和结束的引号不是字段的一部分。

The source:

来源：

```
normal string,"quoted-field"
```

results in the fields

的结果是字段

```
{`normal string`, `quoted-field`}
```

Within a quoted-field a quote character followed by a second quote character is considered a single quote.

在一个带引号的字段内，一个引号字符后面的第二个引号字符被认为是一个单引号。

```
"the ""word"" is true","a ""quoted-field"""
```

results in

结果是

```
{`the "word" is true`, `a "quoted-field"`}
```

Newlines and commas may be included in a quoted-field

换行符和逗号可以包含在一个引号字段中。

```
"Multi-line
field","comma is ,"
```

results in

结果是

```
{`Multi-line
field`, `comma is ,`}
```











## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/csv/reader.go;l=86)

``` go 
var (
	ErrBareQuote  = errors.New("bare \" in non-quoted-field")
	ErrQuote      = errors.New("extraneous or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")

	// Deprecated: ErrTrailingComma is no longer used.
	// 已废弃：ErrTrailingComma不再使用。
	ErrTrailingComma = errors.New("extra delimiter at end of line")
)
```

These are the errors that can be returned in ParseError.Err.

这些是可以在ParseError.Err中返回的错误。

## 函数

This section is empty.

## 类型

### type ParseError 

``` go 
type ParseError struct {
	StartLine int   // Line where the record starts// 记录开始的那一行
	Line      int   // Line where the error occurred// 发生错误的那一行
	Column    int   // Column (1-based byte index) where the error occurred // 发生错误的那一列(基于1的字节索引)。
	Err       error // The actual error // 实际的错误
}
```

A ParseError is returned for parsing errors. Line numbers are 1-indexed and columns are 0-indexed.

对于解析错误会返回一个ParseError。行号是1-索引的，列是0-索引的。

#### (*ParseError) Error 

``` go 
func (e *ParseError) Error() string
```

#### (*ParseError) Unwrap  <- go1.13

``` go 
func (e *ParseError) Unwrap() error
```

### type Reader 

``` go 
type Reader struct {
	// Comma is the field delimiter.
	// It is set to comma (',') by NewReader.
	// Comma must be a valid rune and must not be \r, \n,
	// or the Unicode replacement character (0xFFFD).
	// 逗号是字段分隔符。
	// 它被NewReader设置为逗号(',')。
	// 逗号必须是一个有效的符文，不能是 \r, \n,
	// 或Unicode替换字符(0xFFFD)。
	Comma rune

	// Comment, if not 0, is the comment character. Lines beginning with the
	// Comment character without preceding whitespace are ignored.
	// With leading whitespace the Comment character becomes part of the
	// field, even if TrimLeadingSpace is true.
	// Comment must be a valid rune and must not be \r, \n,
	// or the Unicode replacement character (0xFFFD).
	// It must also not be equal to Comma.
	//注释，如果不是0，就是注释字符。以Comment字符开始的行，如果没有前面的空白，将被忽略。
	// 如果有前导空白，Comment字符将成为字段的一部分，即使TrimLeadingSpace为真。
	// Comment必须是一个有效的符文，不能是 \r, \n, 或Unicode替换字符(0xFFFD)。
	// 它也不能等于逗号。
	Comment rune

	// FieldsPerRecord is the number of expected fields per record.
	// If FieldsPerRecord is positive, Read requires each record to
	// have the given number of fields. If FieldsPerRecord is 0, Read sets it to
	// the number of fields in the first record, so that future records must
	// have the same field count. If FieldsPerRecord is negative, no check is
	// made and records may have a variable number of fields.
	// FieldsPerRecord是每条记录的预期字段数。
	// 如果FieldsPerRecord是正数，Read要求每条记录都有给定的字段数。如果FieldsPerRecord为0，Read将其设置为第一条记录的字段数，这样以后的记录必须有相同的字段数。如果FieldsPerRecord为负数，则不做检查，记录可能有可变的字段数。
	FieldsPerRecord int

	// If LazyQuotes is true, a quote may appear in an unquoted field and a
	// non-doubled quote may appear in a quoted field.
	// 如果LazyQuotes为真，一个引号可能出现在一个无引号的字段中，一个非双引号可能出现在一个有引号的字段中。
	LazyQuotes bool

	// If TrimLeadingSpace is true, leading white space in a field is ignored.
	// This is done even if the field delimiter, Comma, is white space.
	// 如果TrimLeadingSpace为真，字段中的前导空白将被忽略。
	// 即使字段的分隔符Comma是空白的，也会这样做。
	TrimLeadingSpace bool

	// ReuseRecord controls whether calls to Read may return a slice sharing
	// the backing array of the previous call's returned slice for performance.
	// By default, each call to Read returns newly allocated memory owned by the caller.
	// ReuseRecord控制对Read的调用是否可以返回共享前一次调用返回的片断的支持数组的片断，以提高性能。
	// 默认情况下，每次对Read的调用都会返回由调用者拥有的新分配的内存。
	ReuseRecord bool

	// Deprecated: TrailingComma is no longer used.
	// 已废弃：不再使用TrailingComma。
	TrailingComma bool
	// contains filtered or unexported fields
}
```

A Reader reads records from a CSV-encoded file.

Reader 从一个CSV编码的文件中读取记录。

As returned by NewReader, a Reader expects input conforming to [RFC 4180](https://rfc-editor.org/rfc/rfc4180.html). The exported fields can be changed to customize the details before the first call to Read or ReadAll.

正如NewReader所返回的那样，Reader期望输入的内容符合RFC 4180的规定。在第一次调用Read或ReadAll之前，导出的字段可以被改变以定制细节。

The Reader converts all \r\n sequences in its input to plain \n, including in multiline field values, so that the returned data does not depend on which line-ending convention an input file uses.

Reader 将其输入中的所有\r\n序列转换为普通的\n，包括在多行字段值中，因此返回的数据不依赖于输入文件使用的行结束惯例。

##### Example
``` go 
```

##### Example
``` go 
```

#### func NewReader 

``` go 
func NewReader(r io.Reader) *Reader
```

NewReader returns a new Reader that reads from r.

NewReader返回一个新的阅读器，从r中读取数据。

#### (*Reader) FieldPos  <- go1.17

``` go 
func (r *Reader) FieldPos(field int) (line, column int)
```

FieldPos returns the line and column corresponding to the start of the field with the given index in the slice most recently returned by Read. Numbering of lines and columns starts at 1; columns are counted in bytes, not runes.

FieldPos返回对应于最近由Read返回的片断中具有给定索引的字段开始的行和列。行和列的编号从1开始；列的计数单位是字节，而不是符码。

If this is called with an out-of-bounds index, it panics.

如果在调用这个函数时，索引超出了范围，它就会惊慌失措。

#### (*Reader) InputOffset  <- go1.19

``` go 
func (r *Reader) InputOffset() int64
```

InputOffset returns the input stream byte offset of the current reader position. The offset gives the location of the end of the most recently read row and the beginning of the next row.

InputOffset返回当前阅读器位置的输入流字节偏移。这个偏移量给出了最近读取的行的结束和下一行的开始的位置。

#### (*Reader) Read 

``` go 
func (r *Reader) Read() (record []string, err error)
```

Read reads one record (a slice of fields) from r. If the record has an unexpected number of fields, Read returns the record along with the error ErrFieldCount. Except for that case, Read always returns either a non-nil record or a non-nil error, but not both. If there is no data left to be read, Read returns nil, io.EOF. If ReuseRecord is true, the returned slice may be shared between multiple calls to Read.

如果记录有一个意外的字段数，Read会返回记录和错误ErrFieldCount。除了这种情况，Read总是返回一个非空的记录或一个非空的错误，但不会同时返回。如果没有数据可读，Read返回nil，即io.EOF。如果ReuseRecord为真，返回的片断可以在多次调用Read时共享。

#### (*Reader) ReadAll 

``` go 
func (r *Reader) ReadAll() (records [][]string, err error)
```

ReadAll reads all the remaining records from r. Each record is a slice of fields. A successful call returns err == nil, not err == io.EOF. Because ReadAll is defined to read until EOF, it does not treat end of file as an error to be reported.

ReadAll从r读取所有剩余的记录。一个成功的调用返回err == nil，而不是err == io.EOF。因为ReadAll被定义为读到EOF为止，它不把文件结束作为一个错误来报告。

##### Example
``` go 
```

### type Writer 

``` go 
type Writer struct {
	Comma   rune // Field delimiter (set to ',' by NewWriter) // 字段分隔符(NewWriter设置为',' )。
	UseCRLF bool // True to use \r\n as the line terminator // True，使用 \r\n 作为行结束符。
	// contains filtered or unexported fields
}
```

A Writer writes records using CSV encoding.

Writer使用CSV编码来写记录。

As returned by NewWriter, a Writer writes records terminated by a newline and uses ',' as the field delimiter. The exported fields can be changed to customize the details before the first call to Write or WriteAll.

正如NewWriter所返回的那样，Writer写入的记录以换行方式结束，并使用','作为字段分隔符。在第一次调用Write或WriteAll之前，可以改变导出的字段以定制细节。

Comma is the field delimiter.

逗号是字段分隔符。

If UseCRLF is true, the Writer ends each output line with \r\n instead of \n.

如果UseCRLF为真，Writer会以\r\n而不是\n结束每个输出行。

The writes of individual records are buffered. After all data has been written, the client should call the Flush method to guarantee all data has been forwarded to the underlying io.Writer. Any errors that occurred should be checked by calling the Error method.

单个记录的写入是缓冲的。在所有数据被写入后，客户端应该调用Flush方法以保证所有数据都被转发到底层的io.Writer。任何发生的错误都应该通过调用Error方法来检查。

##### Example
``` go 
```

#### func NewWriter 

``` go 
func NewWriter(w io.Writer) *Writer
```

NewWriter returns a new Writer that writes to w.

NewWriter返回一个新的写入w的Writer。

#### (*Writer) Error  <- go1.1

``` go 
func (w *Writer) Error() error
```

Error reports any error that has occurred during a previous Write or Flush.

Error报告在之前的写或刷新过程中发生的任何错误。

#### (*Writer) Flush 

``` go 
func (w *Writer) Flush()
```

Flush writes any buffered data to the underlying io.Writer. To check if an error occurred during the Flush, call Error.

Flush将任何缓冲的数据写入底层的io.Writer。要检查在Flush过程中是否有错误发生，请调用Error。

#### (*Writer) Write 

``` go 
func (w *Writer) Write(record []string) error
```

Write writes a single CSV record to w along with any necessary quoting. A record is a slice of strings with each string being one field. Writes are buffered, so Flush must eventually be called to ensure that the record is written to the underlying io.Writer.

Write将一条CSV记录和任何必要的引号一起写到w中。一个记录是一个字符串的切片，每个字符串是一个字段。写入是缓冲的，所以最终必须调用Flush以确保记录被写入底层的io.Writer。

#### (*Writer) WriteAll 

``` go 
func (w *Writer) WriteAll(records [][]string) error
```

WriteAll writes multiple CSV records to w using Write and then calls Flush, returning any error from the Flush.

WriteAll使用Write将多个CSV记录写入w，然后调用Flush，返回Flush的任何错误。

``` go 
package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

```

