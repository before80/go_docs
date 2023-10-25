+++
title = "quotedprintable"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/mime/quotedprintable@go1.20.1



Package quotedprintable implements quoted-printable encoding as specified by [RFC 2045](https://rfc-editor.org/rfc/rfc2045.html).








## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Reader 

``` go 
type Reader struct {
	// contains filtered or unexported fields
}
```

Reader is a quoted-printable decoder.

#### func NewReader 

``` go 
func NewReader(r io.Reader) *Reader
```

NewReader returns a quoted-printable reader, decoding from r.

##### Example
``` go 
```

#### (*Reader) Read 

``` go 
func (r *Reader) Read(p []byte) (n int, err error)
```

Read reads and decodes quoted-printable data from the underlying reader.

### type Writer 

``` go 
type Writer struct {
	// Binary mode treats the writer's input as pure binary and processes end of
	// line bytes as binary data.
	Binary bool
	// contains filtered or unexported fields
}
```

A Writer is a quoted-printable writer that implements io.WriteCloser.

#### func NewWriter 

``` go 
func NewWriter(w io.Writer) *Writer
```

NewWriter returns a new Writer that writes to w.

##### Example
``` go 
```

#### (*Writer) Close 

``` go 
func (w *Writer) Close() error
```

Close closes the Writer, flushing any unwritten data to the underlying io.Writer, but does not close the underlying io.Writer.

#### (*Writer) Write 

``` go 
func (w *Writer) Write(p []byte) (n int, err error)
```

Write encodes p using quoted-printable encoding and writes it to the underlying io.Writer. It limits line length to 76 characters. The encoded bytes are not necessarily flushed until the Writer is closed.