+++
title = "gzip"
date = 2023-05-17T09:59:21+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/compress/gzip@go1.21.3](https://pkg.go.dev/compress/gzip@go1.21.3)

Package gzip implements reading and writing of gzip format compressed files, as specified in [RFC 1952](https://rfc-editor.org/rfc/rfc1952.html).

​	gzip 包实现了对 gzip 格式压缩文件（如 [RFC 1952](https://rfc-editor.org/rfc/rfc1952.html) 中所述）的读写。

## Example (CompressingReader)
``` go 
package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func main() {
	// This is an example of writing a compressing reader.
	// This can be useful for an HTTP client body, as shown.

	const testdata = "the data to be compressed"

	// This HTTP handler is just for testing purposes.
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		zr, err := gzip.NewReader(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Just output the data for the example.
		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// The remainder is the example code.

	// The data we want to compress, as an io.Reader
	dataReader := strings.NewReader(testdata)

	// bodyReader is the body of the HTTP request, as an io.Reader.
	// httpWriter is the body of the HTTP request, as an io.Writer.
	bodyReader, httpWriter := io.Pipe()

	// Make sure that bodyReader is always closed, so that the
	// goroutine below will always exit.
	defer bodyReader.Close()

	// gzipWriter compresses data to httpWriter.
	gzipWriter := gzip.NewWriter(httpWriter)

	// errch collects any errors from the writing goroutine.
	errch := make(chan error, 1)

	go func() {
		defer close(errch)
		sentErr := false
		sendErr := func(err error) {
			if !sentErr {
				errch <- err
				sentErr = true
			}
		}

		// Copy our data to gzipWriter, which compresses it to
		// gzipWriter, which feeds it to bodyReader.
		if _, err := io.Copy(gzipWriter, dataReader); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
		if err := gzipWriter.Close(); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
		if err := httpWriter.Close(); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
	}()

	// Send an HTTP request to the test server.
	req, err := http.NewRequest("PUT", ts.URL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	// Note that passing req to http.Client.Do promises that it
	// will close the body, in this case bodyReader.
	resp, err := ts.Client().Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// Check whether there was an error compressing the data.
	if err := <-errch; err != nil {
		log.Fatal(err)
	}

	// For this example we don't care about the response.
	resp.Body.Close()

}
Output:

the data to be compressed
```

## Example (WriterReader)
``` go 
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	// Setting the Header fields is optional.
	zw.Name = "a-new-hope.txt"
	zw.Comment = "an epic space opera by George Lucas"
	zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write([]byte("A long time ago in a galaxy far, far away..."))
	if err != nil {
		log.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

}
Output:

Name: a-new-hope.txt
Comment: an epic space opera by George Lucas
ModTime: 1977-05-25 00:00:00 +0000 UTC

A long time ago in a galaxy far, far away...
```

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.21.3:src/compress/gzip/gzip.go;l=18)

``` go 
const (
	NoCompression      = flate.NoCompression
	BestSpeed          = flate.BestSpeed
	BestCompression    = flate.BestCompression
	DefaultCompression = flate.DefaultCompression
	HuffmanOnly        = flate.HuffmanOnly
)
```

These constants are copied from the flate package, so that code that imports "compress/gzip" does not also have to import "compress/flate".

​	这些常量从 flate 包中复制，因此导入 "compress/gzip" 的代码不必再导入"compress/flate"。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/compress/gzip/gunzip.go;l=30)

``` go 
var (
	// ErrChecksum is returned when reading GZIP data that has an invalid checksum.
	ErrChecksum = errors.New("gzip: invalid checksum")
	// ErrHeader is returned when reading GZIP data that has an invalid header.
	ErrHeader = errors.New("gzip: invalid header")
)
```

## 函数

This section is empty.

## 类型

### type Header 

``` go 
type Header struct {
	Comment string    // comment
	Extra   []byte    // "extra data"
	ModTime time.Time // modification time
	Name    string    // file name
	OS      byte      // operating system type
}
```

The gzip file stores a header giving metadata about the compressed file. That header is exposed as the fields of the Writer and Reader structs.

​	gzip 文件存储一个标头，提供有关压缩文件元数据的信息。该标头作为 Writer 和 Reader 结构的字段公开。

Strings must be UTF-8 encoded and may only contain Unicode code points U+0001 through U+00FF, due to limitations of the GZIP file format.

​	由于 GZIP 文件格式的限制，字符串必须采用 UTF-8 编码，并且只能包含 Unicode 代码点 U+0001 到 U+00FF。

### type Reader 

``` go 
type Reader struct {
	Header // valid after NewReader or Reader.Reset
	// contains filtered or unexported fields
}
```

A Reader is an io.Reader that can be read to retrieve uncompressed data from a gzip-format compressed file.

​	Reader 是一个 io.Reader，可以读取它以从 gzip 格式压缩文件中检索未压缩的数据。

In general, a gzip file can be a concatenation of gzip files, each with its own header. Reads from the Reader return the concatenation of the uncompressed data of each. Only the first header is recorded in the Reader fields.

​	通常，gzip 文件可以是 gzip 文件的串联，每个文件都有自己的标头。从 Reader 读取的内容返回每个文件的未压缩数据的串联。只有第一个标头记录在 Reader 字段中。

Gzip files store a length and checksum of the uncompressed data. The Reader will return an ErrChecksum when Read reaches the end of the uncompressed data if it does not have the expected length or checksum. Clients should treat data returned by Read as tentative until they receive the io.EOF marking the end of the data.

​	Gzip 文件存储解压缩数据的长度和校验和。如果解压缩数据没有预期的长度或校验和，则当 Read 达到解压缩数据的末尾时，Reader 将返回 ErrCheckSum。应将 Read 返回的数据视为有效，直到接收到 io.EOF 作为数据结束的标志。

#### func NewReader 

``` go 
func NewReader(r io.Reader) (*Reader, error)
```

NewReader creates a new Reader reading the given reader. If r does not also implement io.ByteReader, the decompressor may read more data than necessary from r.

​	NewReader 创建一个新的 Reader 来读取给定的 reader。如果 r 也没有实现 io.ByteReader，则解压缩器可能会从 r 中读取比必要更多的 data。

It is the caller's responsibility to call Close on the Reader when done.

​	在完成时，由调用者负责对 Reader 调用 Close。

The Reader.Header fields will be valid in the Reader returned.

​	返回的 Reader 中的 Reader.Header 字段将是有效的。

#### (*Reader) Close 

``` go 
func (z *Reader) Close() error
```

Close closes the Reader. It does not close the underlying io.Reader. In order for the GZIP checksum to be verified, the reader must be fully consumed until the io.EOF.

​	Close 关闭 Reader。它不会关闭底层的 io.Reader。为了验证 Gzip 校验和，必须完全使用 reader 直到 io.EOF。

#### (*Reader) Multistream  <- go1.4

``` go 
func (z *Reader) Multistream(ok bool)
```

Multistream controls whether the reader supports multistream files.

​	Multistream 控件 reader 是否支持多流文件。

If enabled (the default), the Reader expects the input to be a sequence of individually gzipped data streams, each with its own header and trailer, ending at EOF. The effect is that the concatenation of a sequence of gzipped files is treated as equivalent to the gzip of the concatenation of the sequence. This is standard behavior for gzip readers.

​	如果启用（默认），Reader 期望输入为一系列单独的 gzip 数据流，每个数据流都有自己的头和尾，以 EOF 结尾。其效果是，将一系列 gzip 文件连接起来处理，等同于连接该序列的 gzip。这是 gzip 读取器的标准行为。

Calling Multistream(false) disables this behavior; disabling the behavior can be useful when reading file formats that distinguish individual gzip data streams or mix gzip data streams with other data streams. In this mode, when the Reader reaches the end of the data stream, Read returns io.EOF. The underlying reader must implement io.ByteReader in order to be left positioned just after the gzip stream. To start the next stream, call z.Reset(r) followed by z.Multistream(false). If there is no next stream, z.Reset(r) will return io.EOF.

​	调用 Multistream(false) 会禁用此行为；在读取区分各个 gzip 数据流或将 gzip 数据流与其他数据流混合的文件格式时，禁用此行为可能很有用。在此模式下，当 Reader 到达数据流的末尾时，Read 返回 io.EOF。底层读取器必须实现 io.ByteReader，以便将其定位在 gzip 流的正后方。要启动下一个流，请依次调用 z.Reset(r) 和 z.Multistream(false)。如果没有下一个流，z.Reset(r) 将返回 io.EOF。

##### Multistream Example
``` go 
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	var files = []struct {
		name    string
		comment string
		modTime time.Time
		data    string
	}{
		{"file-1.txt", "file-header-1", time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC), "Hello Gophers - 1"},
		{"file-2.txt", "file-header-2", time.Date(2007, time.March, 2, 4, 5, 6, 1, time.UTC), "Hello Gophers - 2"},
	}

	for _, file := range files {
		zw.Name = file.name
		zw.Comment = file.comment
		zw.ModTime = file.modTime

		if _, err := zw.Write([]byte(file.data)); err != nil {
			log.Fatal(err)
		}

		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		zw.Reset(&buf)
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	for {
		zr.Multistream(false)
		fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}

		fmt.Print("\n\n")

		err = zr.Reset(&buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

}

Output:

Name: file-1.txt
Comment: file-header-1
ModTime: 2006-02-01 03:04:05 +0000 UTC

Hello Gophers - 1

Name: file-2.txt
Comment: file-header-2
ModTime: 2007-03-02 04:05:06 +0000 UTC

Hello Gophers - 2
```

#### (*Reader) Read 

``` go 
func (z *Reader) Read(p []byte) (n int, err error)
```

Read implements io.Reader, reading uncompressed bytes from its underlying Reader.

​	Read 实现 io.Reader，从其底层 Reader 读取未压缩的字节。

#### (*Reader) Reset  <- go1.3

``` go 
func (z *Reader) Reset(r io.Reader) error
```

Reset discards the Reader z's state and makes it equivalent to the result of its original state from NewReader, but reading from r instead. This permits reusing a Reader rather than allocating a new one.

​	重置会丢弃 Reader z 的状态，并使其等同于 NewReader 的原始状态，但从 r 读取。这允许重用 Reader 而不是分配一个新的 Reader。

### type Writer 

``` go 
type Writer struct {
	Header // written at first call to Write, Flush, or Close
	// contains filtered or unexported fields
}
```

A Writer is an io.WriteCloser. Writes to a Writer are compressed and written to w.

​	Writer 是一个 io.WriteCloser。对 Writer 的写入会被压缩并写入 w。

#### func NewWriter 

``` go 
func NewWriter(w io.Writer) *Writer
```

NewWriter returns a new Writer. Writes to the returned writer are compressed and written to w.

​	NewWriter 返回一个新的 Writer。对返回的 writer 的写入会被压缩并写入 w。

It is the caller's responsibility to call Close on the Writer when done. Writes may be buffered and not flushed until Close.

​	在完成时，由调用者负责对 Writer 调用 Close。写入可能会被缓冲，直到 Close 才刷新。

Callers that wish to set the fields in Writer.Header must do so before the first call to Write, Flush, or Close.

​	希望设置 Writer.Header 中字段的调用者必须在第一次调用 Write、Flush 或 Close 之前这样做。

#### func NewWriterLevel 

``` go 
func NewWriterLevel(w io.Writer, level int) (*Writer, error)
```

NewWriterLevel is like NewWriter but specifies the compression level instead of assuming DefaultCompression.

​	NewWriterLevel 类似于 NewWriter，但指定压缩级别而不是假定 DefaultCompression。

The compression level can be DefaultCompression, NoCompression, HuffmanOnly or any integer value between BestSpeed and BestCompression inclusive. The error returned will be nil if the level is valid.

​	压缩级别可以是 DefaultCompression、NoCompression、HuffmanOnly 或介于 BestSpeed 和 BestCompression（包括）之间的任何整数值。如果级别有效，则返回的错误将为 nil。

#### (*Writer) Close 

``` go 
func (z *Writer) Close() error
```

Close closes the Writer by flushing any unwritten data to the underlying io.Writer and writing the GZIP footer. It does not close the underlying io.Writer.

​	Close 通过将任何未写入的数据刷新到底层 io.Writer 并写入 GZIP 页脚来关闭 Writer。它不会关闭底层 io.Writer。

#### (*Writer) Flush  <- go1.1

``` go 
func (z *Writer) Flush() error
```

Flush flushes any pending compressed data to the underlying writer.

​	Flush 刷新所有待处理的压缩数据到底层写入器。

It is useful mainly in compressed network protocols, to ensure that a remote reader has enough data to reconstruct a packet. Flush does not return until the data has been written. If the underlying writer returns an error, Flush returns that error.

​	它主要用于压缩网络协议，以确保远程读取器有足够的数据来重建数据包。在数据被写入之前，Flush 不会返回。如果底层写入器返回错误，Flush 将返回该错误。

In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.

​	在 zlib 库的术语中，Flush 等效于 Z_SYNC_FLUSH。

#### (*Writer) Reset  <- go1.2

``` go 
func (z *Writer) Reset(w io.Writer)
```

Reset discards the Writer z's state and makes it equivalent to the result of its original state from NewWriter or NewWriterLevel, but writing to w instead. This permits reusing a Writer rather than allocating a new one.

​	Reset  会丢弃 Writer z 的状态，并使其等同于 NewWriter 或 NewWriterLevel 的原始状态的结果，但会写入 w。这允许重用 Writer 而不是分配一个新的。

#### (*Writer) Write 

``` go 
func (z *Writer) Write(p []byte) (int, error)
```

Write writes a compressed form of p to the underlying io.Writer. The compressed bytes are not necessarily flushed until the Writer is closed.

​	Write 将 p 的一个已 compress 的形式写到 io.Writer。在 Writer 关闭之前，compress 的 bytes 不一定会被刷新。