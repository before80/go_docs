+++
title = "flate"
date = 2023-05-17T09:59:21+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# flate

https://pkg.go.dev/compress/flate@go1.20.1



Package flate implements the DEFLATE compressed data format, described in [RFC 1951](https://rfc-editor.org/rfc/rfc1951.html). The gzip and zlib packages implement access to DEFLATE-based file formats.

##### Example (Dictionary)

A preset dictionary can be used to improve the compression ratio. The downside to using a dictionary is that the compressor and decompressor must agree in advance what dictionary to use.

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

##### Example  (Reset)

In performance critical applications, Reset can be used to discard the current compressor or decompressor state and reinitialize them quickly by taking advantage of previously allocated memory.

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

##### Example  (Synchronization)

DEFLATE is suitable for transmitting compressed data across the network.

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

#### func NewReader 

``` go 
func NewReader(r io.Reader) io.ReadCloser
```

NewReader returns a new ReadCloser that can be used to read the uncompressed version of r. If r does not also implement io.ByteReader, the decompressor may read more data than necessary from r. The reader returns io.EOF after the final block in the DEFLATE stream has been encountered. Any trailing data after the final block is ignored.

The ReadCloser returned by NewReader also implements Resetter.

#### func NewReaderDict 

``` go 
func NewReaderDict(r io.Reader, dict []byte) io.ReadCloser
```

NewReaderDict is like NewReader but initializes the reader with a preset dictionary. The returned Reader behaves as if the uncompressed data stream started with the given dictionary, which has already been read. NewReaderDict is typically used to read data compressed by NewWriterDict.

The ReadCloser returned by NewReader also implements Resetter.

## 类型

### type CorruptInputError 

``` go 
type CorruptInputError int64
```

A CorruptInputError reports the presence of corrupt input at a given offset.

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

#### type ReadError <- DEPRECATED


### type Reader 

``` go 
type Reader interface {
	io.Reader
	io.ByteReader
}
```

The actual read interface needed by NewReader. If the passed in io.Reader does not also have ReadByte, the NewReader will introduce its own buffering.

### type Resetter  <- go1.4

``` go 
type Resetter interface {
	// Reset discards any buffered data and resets the Resetter as if it was
	// newly initialized with the given reader.
	Reset(r io.Reader, dict []byte) error
}
```

Resetter resets a ReadCloser returned by NewReader or NewReaderDict to switch to a new underlying Reader. This permits reusing a ReadCloser instead of allocating a new one.

#### type WriteError <- DEPRECATED
### type Writer 

``` go 
type Writer struct {
	// contains filtered or unexported fields
}
```

A Writer takes data written to it and writes the compressed form of that data to an underlying writer (see NewWriter).

#### func NewWriter 

``` go 
func NewWriter(w io.Writer, level int) (*Writer, error)
```

NewWriter returns a new Writer compressing data at the given level. Following zlib, levels range from 1 (BestSpeed) to 9 (BestCompression); higher levels typically run slower but compress more. Level 0 (NoCompression) does not attempt any compression; it only adds the necessary DEFLATE framing. Level -1 (DefaultCompression) uses the default compression level. Level -2 (HuffmanOnly) will use Huffman compression only, giving a very fast compression for all types of input, but sacrificing considerable compression efficiency.

If level is in the range [-2, 9] then the error returned will be nil. Otherwise the error returned will be non-nil.

#### func NewWriterDict 

``` go 
func NewWriterDict(w io.Writer, level int, dict []byte) (*Writer, error)
```

NewWriterDict is like NewWriter but initializes the new Writer with a preset dictionary. The returned Writer behaves as if the dictionary had been written to it without producing any compressed output. The compressed data written to w can only be decompressed by a Reader initialized with the same dictionary.

#### (*Writer) Close 

``` go 
func (w *Writer) Close() error
```

Close flushes and closes the writer.

#### (*Writer) Flush 

``` go 
func (w *Writer) Flush() error
```

Flush flushes any pending data to the underlying writer. It is useful mainly in compressed network protocols, to ensure that a remote reader has enough data to reconstruct a packet. Flush does not return until the data has been written. Calling Flush when there is no pending data still causes the Writer to emit a sync marker of at least 4 bytes. If the underlying writer returns an error, Flush returns that error.

In the terminology of the zlib library, Flush is equivalent to Z_SYNC_FLUSH.

#### (*Writer) Reset  <- go1.2

``` go 
func (w *Writer) Reset(dst io.Writer)
```

Reset discards the writer's state and makes it equivalent to the result of NewWriter or NewWriterDict called with dst and w's level and dictionary.

#### (*Writer) Write 

``` go 
func (w *Writer) Write(data []byte) (n int, err error)
```

Write writes data to w, which will eventually write the compressed form of data to its underlying writer.