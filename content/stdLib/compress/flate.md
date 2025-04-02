+++
title = "flate"
date = 2023-05-17T09:59:21+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/compress/flate@go1.24.2](https://pkg.go.dev/compress/flate@go1.24.2)

Package flate implements the DEFLATE compressed data format, described in [RFC 1951](https://rfc-editor.org/rfc/rfc1951.html). The gzip and zlib packages implement access to DEFLATE-based file formats.

​	flate 包实现 DEFLATE 压缩数据格式，在 RFC 1951 中进行了描述。gzip 和 zlib 包实现对基于 DEFLATE 的文件格式的访问。

## Example (Dictionary)

A preset dictionary can be used to improve the compression ratio. The downside to using a dictionary is that the compressor and decompressor must agree in advance what dictionary to use.

​	可以使用预设字典来提高压缩率。使用字典的缺点是，压缩器和解压缩器必须预先协商好使用哪个字典。

``` go 
package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// The dictionary is a string of bytes. When compressing some input data,
	// the compressor will attempt to substitute substrings with matches found
	// in the dictionary. As such, the dictionary should only contain substrings
	// that are expected to be found in the actual data stream.
	const dict = `<?xml version="1.0"?>` + `<book>` + `<data>` + `<meta name="` + `" content="`

	// The data to compress should (but is not required to) contain frequent
	// substrings that match those in the dictionary.
	const data = `<?xml version="1.0"?>
<book>
	<meta name="title" content="The Go Programming Language"/>
	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	<meta name="published" content="2015-10-26"/>
	<meta name="isbn" content="978-0134190440"/>
	<data>...</data>
</book>
`

	var b bytes.Buffer

	// Compress the data using the specially crafted dictionary.
	zw, err := flate.NewWriterDict(&b, flate.DefaultCompression, []byte(dict))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(zw, strings.NewReader(data)); err != nil {
		log.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	// The decompressor must use the same dictionary as the compressor.
	// Otherwise, the input may appear as corrupted.
	fmt.Println("Decompressed output using the dictionary:")
	zr := flate.NewReaderDict(bytes.NewReader(b.Bytes()), []byte(dict))
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	// Substitute all of the bytes in the dictionary with a '#' to visually
	// demonstrate the approximate effectiveness of using a preset dictionary.
	fmt.Println("Substrings matched by the dictionary are marked with #:")
	hashDict := []byte(dict)
	for i := range hashDict {
		hashDict[i] = '#'
	}
	zr = flate.NewReaderDict(&b, hashDict)
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

}
Output:

Decompressed output using the dictionary:
<?xml version="1.0"?>
<book>
	<meta name="title" content="The Go Programming Language"/>
	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	<meta name="published" content="2015-10-26"/>
	<meta name="isbn" content="978-0134190440"/>
	<data>...</data>
</book>

Substrings matched by the dictionary are marked with #:
#####################
######
	############title###########The Go Programming Language"/#
	############authors###########Alan Donovan and Brian Kernighan"/#
	############published###########2015-10-26"/#
	############isbn###########978-0134190440"/#
	######...</#####
</#####
```

## Example  (Reset)

In performance critical applications, Reset can be used to discard the current compressor or decompressor state and reinitialize them quickly by taking advantage of previously allocated memory.

​	在对性能至关重要的应用程序中，可以使用 Reset 来丢弃当前压缩器或解压缩器状态，并通过利用先前分配的内存快速重新初始化它们。

``` go 
package main

import (
	"bytes"
	"compress/flate"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.\n",
		"Concurrency is not parallelism.\n",
		"The bigger the interface, the weaker the abstraction.\n",
		"Documentation is for users.\n",
	}

	var r strings.Reader
	var b bytes.Buffer
	buf := make([]byte, 32<<10)

	zw, err := flate.NewWriter(nil, flate.DefaultCompression)
	if err != nil {
		log.Fatal(err)
	}
	zr := flate.NewReader(nil)

	for _, s := range proverbs {
		r.Reset(s)
		b.Reset()

		// Reset the compressor and encode from some input stream.
		zw.Reset(&b)
		if _, err := io.CopyBuffer(zw, &r, buf); err != nil {
			log.Fatal(err)
		}
		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		// Reset the decompressor and decode to some output stream.
		if err := zr.(flate.Resetter).Reset(&b, nil); err != nil {
			log.Fatal(err)
		}
		if _, err := io.CopyBuffer(os.Stdout, zr, buf); err != nil {
			log.Fatal(err)
		}
		if err := zr.Close(); err != nil {
			log.Fatal(err)
		}
	}

}
Output:

Don't communicate by sharing memory, share memory by communicating.
Concurrency is not parallelism.
The bigger the interface, the weaker the abstraction.
Documentation is for users.
```

## Example  (Synchronization)

DEFLATE is suitable for transmitting compressed data across the network.

​	DEFLATE 适用于通过网络传输压缩数据。

``` go 
package main

import (
	"compress/flate"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	// Use io.Pipe to simulate a network connection.
	// A real network application should take care to properly close the
	// underlying connection.
	rp, wp := io.Pipe()

	// Start a goroutine to act as the transmitter.
	wg.Add(1)
	go func() {
		defer wg.Done()

		zw, err := flate.NewWriter(wp, flate.BestSpeed)
		if err != nil {
			log.Fatal(err)
		}

		b := make([]byte, 256)
		for _, m := range strings.Fields("A long time ago in a galaxy far, far away...") {
			// We use a simple framing format where the first byte is the
			// message length, followed the message itself.
			b[0] = uint8(copy(b[1:], m))

			if _, err := zw.Write(b[:1+len(m)]); err != nil {
				log.Fatal(err)
			}

			// Flush ensures that the receiver can read all data sent so far.
			if err := zw.Flush(); err != nil {
				log.Fatal(err)
			}
		}

		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Start a goroutine to act as the receiver.
	wg.Add(1)
	go func() {
		defer wg.Done()

		zr := flate.NewReader(rp)

		b := make([]byte, 256)
		for {
			// Read the message length.
			// This is guaranteed to return for every corresponding
			// Flush and Close on the transmitter side.
			if _, err := io.ReadFull(zr, b[:1]); err != nil {
				if err == io.EOF {
					break // The transmitter closed the stream
				}
				log.Fatal(err)
			}

			// Read the message content.
			n := int(b[0])
			if _, err := io.ReadFull(zr, b[:n]); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Received %d bytes: %s\n", n, b[:n])
		}
		fmt.Println()

		if err := zr.Close(); err != nil {
			log.Fatal(err)
		}
	}()

}
Output:

Received 1 bytes: A
Received 4 bytes: long
Received 4 bytes: time
Received 3 bytes: ago
Received 2 bytes: in
Received 1 bytes: a
Received 6 bytes: galaxy
Received 4 bytes: far,
Received 3 bytes: far
Received 7 bytes: away...
```

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/compress/flate/deflate.go;l=14)

``` go 
const (
	NoCompression      = 0
	BestSpeed          = 1
	BestCompression    = 9
	DefaultCompression = -1

	// HuffmanOnly disables Lempel-Ziv match searching and only performs Huffman
	// entropy encoding. This mode is useful in compressing data that has
	// already been compressed with an LZ style algorithm (e.g. Snappy or LZ4)
	// that lacks an entropy encoder. Compression gains are achieved when
	// certain bytes in the input stream occur more frequently than others.
	//
	// Note that HuffmanOnly produces a compressed output that is
	// RFC 1951 compliant. That is, any valid DEFLATE decompressor will
	// continue to be able to decompress this output.
	HuffmanOnly = -2
)
```

## 变量

This section is empty.

## 函数

### func NewReader 

``` go 
func NewReader(r io.Reader) io.ReadCloser
```

NewReader returns a new ReadCloser that can be used to read the uncompressed version of r. If r does not also implement io.ByteReader, the decompressor may read more data than necessary from r. The reader returns io.EOF after the final block in the DEFLATE stream has been encountered. Any trailing data after the final block is ignored.

​	NewReader 返回一个新的 ReadCloser，可用于读取 r 的未压缩版本。如果 r 也没有实现 io.ByteReader，则解压缩器可能会从 r 中读取比必要更多的 data。在遇到 DEFLATE 流中的最后一个块后，读取器返回 io.EOF。最后一个块之后的任何尾随数据都将被忽略。

The ReadCloser returned by NewReader also implements Resetter.

​	NewReader 返回的 ReadCloser 也实现了 Resetter。

### func NewReaderDict 

``` go 
func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
```

NewReaderDict is like NewReader but initializes the reader with a preset dictionary. The returned Reader behaves as if the uncompressed data stream started with the given dictionary, which has already been read. NewReaderDict is typically used to read data compressed by NewWriterDict.

​	NewReaderDict 类似于 NewReader，但使用预设词典初始化读取器。返回的读取器表现得就像未压缩数据流以给定词典开头，该词典已读入。NewReaderDict 通常用于读取由 NewWriterDict 压缩的数据。

The ReadCloser returned by NewReader also implements Resetter.

​	NewReader 返回的 ReadCloser 也实现了 Resetter。

## 类型

### type CorruptInputError 

``` go 
type CorruptInputError int64
```

A CorruptInputError reports the presence of corrupt input at a given offset.

​	CorruptInputError 报告给定偏移量处存在损坏的输入。

#### (CorruptInputError) Error 

``` go 
func (e CorruptInputError) Error() string
```

### type InternalError 

``` go 
type InternalError string
```

An InternalError reports an error in the flate code itself.

#### (InternalError) Error 

``` go 
func (e InternalError) Error() string
```

### type ReadError <- DEPRECATED

```go
type ReadError struct {
	Offset int64 // byte offset where error occurred
	Err    error // error returned by underlying Read
}
```

A ReadError reports an error encountered while reading input.

​	ReadError 报告在读取输入时遇到的错误。

Deprecated: No longer returned.

​	已弃用：不再返回。

#### func (*ReadError) Error

```go
func (e *ReadError) Error() string
```


### type Reader 

``` go 
type Reader interface {
	io.Reader
	io.ByteReader
}
```

The actual read interface needed by NewReader. If the passed in io.Reader does not also have ReadByte, the NewReader will introduce its own buffering.

​	NewReader 所需的实际读取接口。如果传入的 io.Reader 也没有 ReadByte，NewReader 将引入自己的缓冲。

### type Resetter  <- go1.4

``` go 
type Resetter interface {
	// Reset discards any buffered data and resets the Resetter as if it was
	// newly initialized with the given reader.
	Reset(r io.Reader, dict []byte) error
}
```

Resetter resets a ReadCloser returned by NewReader or NewReaderDict to switch to a new underlying Reader. This permits reusing a ReadCloser instead of allocating a new one.

​	Resetter 重置 NewReader 或 NewReaderDict 返回的 ReadCloser 以切换到新的底层 Reader。这允许重用 ReadCloser，而不是分配一个新的。

### type WriteError <- DEPRECATED

```go
type WriteError struct {
	Offset int64 // byte offset where error occurred
	Err    error // error returned by underlying Write
}
```

A WriteError reports an error encountered while writing output.

​	WriteError 报告在写入输出时遇到的错误。

Deprecated: No longer returned.

​	已弃用：不再返回。

####  (*WriteError) Error

```go
func (e *WriteError) Error() string
```

### type Writer 

``` go 
type Writer struct {
	// contains filtered or unexported fields
}
```

A Writer takes data written to it and writes the compressed form of that data to an underlying writer (see NewWriter).

​	Writer 接收写入其中的数据，并将该数据的压缩形式写入底层 writer（请参阅 NewWriter）。

#### func NewWriter 

``` go 
func NewWriter(w io.Writer, level int) (*Writer, error)
```

NewWriter returns a new Writer compressing data at the given level. Following zlib, levels range from 1 (BestSpeed) to 9 (BestCompression); higher levels typically run slower but compress more. Level 0 (NoCompression) does not attempt any compression; it only adds the necessary DEFLATE framing. Level -1 (DefaultCompression) uses the default compression level. Level -2 (HuffmanOnly) will use Huffman compression only, giving a very fast compression for all types of input, but sacrificing considerable compression efficiency.

​	NewWriter 返回一个在给定级别压缩数据的新的 Writer。按照 zlib，级别范围从 1（BestSpeed）到 9（BestCompression）；较高的级别通常运行得较慢，但压缩得更多。级别 0（NoCompression）不尝试任何压缩；它只添加必要的 DEFLATE 框架。级别 -1（DefaultCompression）使用默认压缩级别。级别 -2（HuffmanOnly）将仅使用霍夫曼压缩，为所有类型的输入提供非常快速的压缩，但牺牲了相当大的压缩效率。

If level is in the range [-2, 9] then the error returned will be nil. Otherwise the error returned will be non-nil.

​	如果级别在 `[-2, 9]` 范围内，则返回的错误将为 nil。否则，返回的错误将为非 nil。

#### func NewWriterDict 

``` go 
func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)
```

NewWriterDict is like NewWriter but initializes the new Writer with a preset dictionary. The returned Writer behaves as if the dictionary had been written to it without producing any compressed output. The compressed data written to w can only be decompressed by a Reader initialized with the same dictionary.

​	NewWriterDict 与 NewWriter 类似，但使用预设词典初始化新的 Writer。返回的 Writer 的行为就像词典已经写入其中而没有产生任何压缩输出一样。写入 w 的压缩数据只能由使用相同词典初始化的 Reader 解压缩。

#### (*Writer) Close 

``` go 
func (w *Writer) Close() error
```

Close flushes and closes the writer.

​	Close 刷新并关闭 writer。

#### (*Writer) Flush 

``` go 
func (w *Writer) Flush() error
```

Flush flushes any pending data to the underlying writer. It is useful mainly in compressed network protocols, to ensure that a remote reader has enough data to reconstruct a packet. Flush does not return until the data has been written. Calling Flush when there is no pending data still causes the Writer to emit a sync marker of at least 4 bytes. If the underlying writer returns an error, Flush returns that error.

​	Flush 将所有待处理数据刷新到底层写入器。它主要用于压缩网络协议，以确保远程读取器有足够的数据来重建数据包。在数据写入之前，Flush 不会返回。在没有待处理数据的情况下调用 Flush 仍会导致写入器发出至少 4 个字节的同步标记。如果底层写入器返回错误，Flush 将返回该错误。

In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.

​	在 zlib 库的术语中，Flush 等效于 Z_SYNC_FLUSH。

#### (*Writer) Reset  <- go1.2

``` go 
func (w *Writer) Reset(dst io.Writer)
```

Reset discards the writer's state and makes it equivalent to the result of NewWriter or NewWriterDict called with dst and w's level and dictionary.

​	Reset 丢弃写入器状态，使其等效于使用 dst 和 w 的级别和词典调用的 NewWriter 或 NewWriterDict 的结果。

#### (*Writer) Write 

``` go 
func (w *Writer) Write(data []byte) (n int, err error)
```

Write writes data to w, which will eventually write the compressed form of data to its underlying writer.

​	Write 将数据写入 w，最终将写入数据的压缩形式到其底层编写器。