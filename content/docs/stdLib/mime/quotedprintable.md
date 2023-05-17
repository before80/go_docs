+++
title = "quotedprintable"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# quotedprintable

https://pkg.go.dev/mime/quotedprintable@go1.20.1



Package quotedprintable implements quoted-printable encoding as specified by [RFC 2045](https://rfc-editor.org/rfc/rfc2045.html).








## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type [Reader](https://cs.opensource.google/go/go/+/go1.20.1:src/mime/quotedprintable/reader.go;l=17) 

``` go linenums="1"
type Reader struct {
	// contains filtered or unexported fields
}
```

Reader is a quoted-printable decoder.

#### func [NewReader](https://cs.opensource.google/go/go/+/go1.20.1:src/mime/quotedprintable/reader.go;l=24) 

``` go linenums="1"
func NewReader(r io.Reader) *Reader
```

NewReader returns a quoted-printable reader, decoding from r.

##### Example
``` go linenums="1"
```

#### (*Reader) [Read](https://cs.opensource.google/go/go/+/go1.20.1:src/mime/quotedprintable/reader.go;l=72) 

``` go linenums="1"
func (r *Reader) Read(p []byte) (n int, err error)
```

Read reads and decodes quoted-printable data from the underlying reader.

### type [Writer](https://cs.opensource.google/go/go/+/go1.20.1:src/mime/quotedprintable/writer.go;l=12) 

``` go linenums="1"
type Writer struct {
	// Binary mode treats the writer's input as pure binary and processes end of
	// line bytes as binary data.
	Binary bool
	// contains filtered or unexported fields
}
```

A Writer is a quoted-printable writer that implements io.WriteCloser.

#### func [NewWriter](https://cs.opensource.google/go/go/+/go1.20.1:src/mime/quotedprintable/writer.go;l=24) 

``` go linenums="1"
func NewWriter(w io.Writer) *Writer
```

NewWriter returns a new Writer that writes to w.

##### Example
``` go linenums="1"
```

#### (*Writer) [Close](https://cs.opensource.google/go/go/+/go1.20.1:src/mime/quotedprintable/writer.go;l=67) 

``` go linenums="1"
func (w *Writer) Close() error
```

Close closes the Writer, flushing any unwritten data to the underlying io.Writer, but does not close the underlying io.Writer.

#### (*Writer) [Write](https://cs.opensource.google/go/go/+/go1.20.1:src/mime/quotedprintable/writer.go;l=31) 

``` go linenums="1"
func (w *Writer) Write(p []byte) (n int, err error)
```

Write encodes p using quoted-printable encoding and writes it to the underlying io.Writer. It limits line length to 76 characters. The encoded bytes are not necessarily flushed until the Writer is closed.