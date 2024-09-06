+++
title = "zip"
date = 2023-05-17T09:59:21+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/archive/zip@go1.23.0](https://pkg.go.dev/archive/zip@go1.23.0)

Package zip provides support for reading and writing ZIP archives.

​	zip 包提供对 ZIP 存档的读写支持。

See the [ZIP specification](https://www.pkware.com/appnote) for details.

​	有关详细信息，请参阅 [ZIP 规范](https://www.pkware.com/appnote)。

This package does not support disk spanning.

​	此包不支持跨磁盘。

A note about ZIP64:

​	有关 ZIP64 的说明：

To be backwards compatible the FileHeader has both 32 and 64 bit Size fields. The 64 bit fields will always contain the correct value and for normal archives both fields will be the same. For files requiring the ZIP64 format the 32 bit fields will be 0xffffffff and the 64 bit fields must be used instead.

​	为了向后兼容，FileHeader 同时具有 32 位和 64 位大小字段。64 位字段始终包含正确的值，对于普通存档，两个字段都相同。对于需要 ZIP64 格式的文件，32 位字段将为 0xffffffff，而必须使用 64 位字段。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/archive/zip/struct.go;l=31)

``` go 
const (
	Store   uint16 = 0 // no compression
	Deflate uint16 = 8 // DEFLATE compressed
)
```

Compression methods.

​	压缩方法。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/archive/zip/reader.go;l=27)

``` go 
var (
	ErrFormat       = errors.New("zip: not a valid zip file")
	ErrAlgorithm    = errors.New("zip: unsupported compression algorithm")
	ErrChecksum     = errors.New("zip: checksum error")
	ErrInsecurePath = errors.New("zip: insecure file path")
)
```

## 函数

### func RegisterCompressor  <- go1.2

``` go 
func RegisterCompressor(method uint16, comp Compressor)
```

RegisterCompressor registers custom compressors for a specified method ID. The common methods Store and Deflate are built in.

​	RegisterCompressor 为指定的方法 ID 注册自定义压缩器。内置了常见的方法 Store 和 Deflate。

### func RegisterDecompressor  <- go1.2

``` go 
func RegisterDecompressor(method uint16, dcomp Decompressor)
```

RegisterDecompressor allows custom decompressors for a specified method ID. The common methods Store and Deflate are built in.

​	RegisterDecompressor 允许为指定的方法 ID 使用自定义解压缩器。内置了常见方法 Store 和 Deflate。

## 类型

### type Compressor  <- go1.2

``` go 
type Compressor func(w io.Writer) (io.WriteCloser, error)
```

A Compressor returns a new compressing writer, writing to w. The WriteCloser's Close method must be used to flush pending data to w. The Compressor itself must be safe to invoke from multiple goroutines simultaneously, but each returned writer will be used only by one goroutine at a time.

​	Compressor 返回一个新的压缩写入器，写入 w。必须使用 WriteCloser 的 Close 方法将待处理数据刷新到 w。Compressor 本身必须安全，可以同时从多个 goroutine 调用，但每个返回的写入器一次只能由一个 goroutine 使用。

### type Decompressor  <- go1.2

``` go 
type Decompressor func(r io.Reader) io.ReadCloser
```

A Decompressor returns a new decompressing reader, reading from r. The ReadCloser's Close method must be used to release associated resources. The Decompressor itself must be safe to invoke from multiple goroutines simultaneously, but each returned reader will be used only by one goroutine at a time.

​	Decompressor 返回一个新的解压缩读取器，从 r 读取。必须使用 ReadCloser 的 Close 方法释放关联的资源。Decompressor 本身必须安全，可以同时从多个 goroutine 调用，但每个返回的读取器一次只能由一个 goroutine 使用。

### type File 

``` go 
type File struct {
	FileHeader
	// contains filtered or unexported fields
}
```

A File is a single file in a ZIP archive. The file information is in the embedded FileHeader. The file content can be accessed by calling Open.

​	File 是 ZIP 存档中的单个文件。文件信息位于嵌入式 FileHeader 中。可以通过调用 Open 访问文件内容。

#### (*File) DataOffset  <- go1.2

``` go 
func (f *File) DataOffset() (offset int64, err error)
```

DataOffset returns the offset of the file's possibly-compressed data, relative to the beginning of the zip file.

​	DataOffset 返回文件可能压缩的数据相对于 zip 文件开头的偏移量。

Most callers should instead use Open, which transparently decompresses data and verifies checksums.

​	大多数调用者应该改用 Open，它会透明地解压缩数据并验证校验和。

#### (*File) Open 

``` go 
func (f *File) Open() (io.ReadCloser, error)
```

Open returns a ReadCloser that provides access to the File's contents. Multiple files may be read concurrently.

​	打开返回一个 ReadCloser，它提供对文件内容的访问。可以同时读取多个文件。

#### (*File) OpenRaw  <- go1.17

``` go 
func (f *File) OpenRaw() (io.Reader, error)
```

OpenRaw returns a Reader that provides access to the File's contents without decompression.

​	OpenRaw 返回一个 Reader，它提供对文件内容的访问，无需解压缩。

### type FileHeader 

``` go 
type FileHeader struct {
	// Name is the name of the file.
	//
	// It must be a relative path, not start with a drive letter (such as "C:"),
	// and must use forward slashes instead of back slashes. A trailing slash
	// indicates that this file is a directory and should have no data.
	Name string

	// Comment is any arbitrary user-defined string shorter than 64KiB.
	Comment string

	// NonUTF8 indicates that Name and Comment are not encoded in UTF-8.
	//
	// By specification, the only other encoding permitted should be CP-437,
	// but historically many ZIP readers interpret Name and Comment as whatever
	// the system's local character encoding happens to be.
	//
	// This flag should only be set if the user intends to encode a non-portable
	// ZIP file for a specific localized region. Otherwise, the Writer
	// automatically sets the ZIP format's UTF-8 flag for valid UTF-8 strings.
	NonUTF8 bool

	CreatorVersion uint16
	ReaderVersion  uint16
	Flags          uint16

	// Method is the compression method. If zero, Store is used.
	Method uint16

	// Modified is the modified time of the file.
	//
	// When reading, an extended timestamp is preferred over the legacy MS-DOS
	// date field, and the offset between the times is used as the timezone.
	// If only the MS-DOS date is present, the timezone is assumed to be UTC.
	//
	// When writing, an extended timestamp (which is timezone-agnostic) is
	// always emitted. The legacy MS-DOS date field is encoded according to the
	// location of the Modified time.
	Modified time.Time

	// ModifiedTime is an MS-DOS-encoded time.
	//
	// Deprecated: Use Modified instead.
	ModifiedTime uint16

	// ModifiedDate is an MS-DOS-encoded date.
	//
	// Deprecated: Use Modified instead.
	ModifiedDate uint16

	// CRC32 is the CRC32 checksum of the file content.
	CRC32 uint32

	// CompressedSize is the compressed size of the file in bytes.
	// If either the uncompressed or compressed size of the file
	// does not fit in 32 bits, CompressedSize is set to ^uint32(0).
	//
	// Deprecated: Use CompressedSize64 instead.
	CompressedSize uint32

	// UncompressedSize is the compressed size of the file in bytes.
	// If either the uncompressed or compressed size of the file
	// does not fit in 32 bits, CompressedSize is set to ^uint32(0).
	//
	// Deprecated: Use UncompressedSize64 instead.
	UncompressedSize uint32

	// CompressedSize64 is the compressed size of the file in bytes.
	CompressedSize64 uint64

	// UncompressedSize64 is the uncompressed size of the file in bytes.
	UncompressedSize64 uint64

	Extra         []byte
	ExternalAttrs uint32 // Meaning depends on CreatorVersion
}
```

FileHeader describes a file within a ZIP file. See the [ZIP specification](https://www.pkware.com/appnote) for details.

​	FileHeader 描述 ZIP 文件中的一个文件。有关详细信息，请参阅 [ZIP 规范](https://www.pkware.com/appnote)。

#### func FileInfoHeader 

``` go 
func FileInfoHeader(fi fs.FileInfo) (*FileHeader, error)
```

FileInfoHeader creates a partially-populated FileHeader from an fs.FileInfo. Because fs.FileInfo's Name method returns only the base name of the file it describes, it may be necessary to modify the Name field of the returned header to provide the full path name of the file. If compression is desired, callers should set the FileHeader.Method field; it is unset by default.

​	FileInfoHeader 从 fs.FileInfo 创建一个部分填充的 FileHeader。由于 fs.FileInfo 的 Name 方法仅返回它所描述的文件的基本名称，因此可能需要修改返回的标头的 Name 字段以提供文件的完整路径名。如果需要压缩，调用者应设置 FileHeader.Method 字段；默认情况下它未设置。

#### (*FileHeader) FileInfo 

``` go 
func (h *FileHeader) FileInfo() fs.FileInfo
```

FileInfo returns an fs.FileInfo for the FileHeader.

​	FileInfo 为 FileHeader 返回一个 fs.FileInfo。

#### (*FileHeader) ModTime <-DEPRECATED

``` go
func (h *FileHeader) ModTime() time.Time
```

ModTime returns the modification time in UTC using the legacy ModifiedDate and ModifiedTime fields.

​	ModTime 使用旧 ModifiedDate 和 ModifiedTime 字段返回 UTC 中的修改时间。

Deprecated: Use Modified instead.

​	已弃用：改用 Modified。

#### (*FileHeader) Mode 

``` go 
func (h *FileHeader) Mode() (mode fs.FileMode)
```

Mode returns the permission and mode bits for the FileHeader.

​	Mode 返回 FileHeader 的权限和模式位。

#### (*FileHeader) SetModTime <-DEPRECATED

``` go
func (h *FileHeader) SetModTime(t time.Time)
```

SetModTime sets the Modified, ModifiedTime, and ModifiedDate fields to the given time in UTC.

​	SetModTime 将 Modified、ModifiedTime 和 ModifiedDate 字段设置为 UTC 中给定的时间。

Deprecated: Use Modified instead.

​	已弃用：改用 Modified。

#### (*FileHeader) SetMode 

``` go 
func (h *FileHeader) SetMode(mode fs.FileMode)
```

SetMode changes the permission and mode bits for the FileHeader.

​	SetMode 更改 FileHeader 的权限和模式位。

### type ReadCloser 

``` go 
type ReadCloser struct {
	Reader
	// contains filtered or unexported fields
}
```

A ReadCloser is a Reader that must be closed when no longer needed.

​	Readcloser 是一个 Reader，在不再需要时必须关闭。

#### func OpenReader 

``` go 
func OpenReader(name string) (*ReadCloser, error)
```

OpenReader will open the Zip file specified by name and return a ReadCloser.

​	OpenReader 将打开 name 指定的 Zip 文件并返回一个 Readcloser。

#### (*ReadCloser) Close 

``` go 
func (rc *ReadCloser) Close() error
```

Close closes the Zip file, rendering it unusable for I/O.

​	Close 关闭 Zip 文件，使其无法用于 I/O。

### type Reader 

``` go 
type Reader struct {
	File    []*File
	Comment string
	// contains filtered or unexported fields
}
```

A Reader serves content from a ZIP archive.

​	Reader 从 ZIP 存档中提供内容。

#### Example
``` go 
package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Open a zip archive for reading.
	r, err := zip.OpenReader("testdata/readme.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}
Output:

Contents of README:
This is the source code repository for the Go programming language.
```

#### func NewReader 

``` go 
func NewReader(r io.ReaderAt, size int64) (*Reader, error)
```

NewReader returns a new Reader reading from r, which is assumed to have the given size in bytes.

​	NewReader 返回一个新的 Reader，从 r 读取，假设 r 的大小为 bytes 字节。

If any file inside the archive uses a non-local name (as defined by [filepath.IsLocal](https://pkg.go.dev/path/filepath#IsLocal)) or a name containing backslashes and the GODEBUG environment variable contains `zipinsecurepath=0`, NewReader returns the reader with an ErrInsecurePath error. A future version of Go may introduce this behavior by default. Programs that want to accept non-local names can ignore the ErrInsecurePath error and use the returned reader.

​	如果存档中的任何文件使用非本地名称（由 [filepath.IsLocal]({{< ref "/stdLib/path/filepath#func-islocal----go120">}}) 定义）或包含反斜杠的名称，并且 GODEBUG 环境变量包含 `zipinsecurepath=0` ，NewReader 将返回带有 ErrInsecurePath 错误的读取器。未来版本的 Go 可能会默认引入此行为。想要接受非本地名称的程序可以忽略 ErrInsecurePath 错误并使用返回的读取器。

#### (*Reader) Open  <- go1.16

``` go 
func (r *Reader) Open(name string) (fs.File, error)
```

Open opens the named file in the ZIP archive, using the semantics of fs.FS.Open: paths are always slash separated, with no leading `/` or `../` elements.

​	Open 使用 fs.FS.Open 的语义打开 ZIP 存档中的指定文件：路径始终用斜杠分隔，没有前导 `/` 或 `../` 元素。

#### (*Reader) RegisterDecompressor  <- go1.6

``` go 
func (z *Reader) RegisterDecompressor(method uint16, dcomp Decompressor)
```

RegisterDecompressor registers or overrides a custom decompressor for a specific method ID. If a decompressor for a given method is not found, Reader will default to looking up the decompressor at the package level.

​	RegisterDecompressor 为特定方法 ID 注册或覆盖自定义解压缩器。如果找不到给定方法的解压缩器，Reader 将默认在包级别查找解压缩器。

### type Writer 

``` go 
type Writer struct {
	// contains filtered or unexported fields
}
```

Writer implements a zip file writer.

​	Writer 实现了一个 zip 文件写入器。

### Example
``` go 
package main

import (
	"archive/zip"
	"bytes"
	"log"
)

func main() {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

output:

```

#### func NewWriter 

``` go 
func NewWriter(w io.Writer) *Writer
```

NewWriter returns a new Writer writing a zip file to w.

​	NewWriter 函数返回一个新的 Writer，将 zip 文件写入 w。

#### (*Writer) AddFS <- go1.22.0

``` go
func (w *Writer) AddFS(fsys fs.FS) error
```

AddFS adds the files from fs.FS to the archive. It walks the directory tree starting at the root of the filesystem adding each file to the zip using deflate while maintaining the directory structure.

​	AddFS 将 fs.FS 中的文件添加到存档中。它从文件系统的根目录开始遍历目录树，使用 deflate 将每个文件添加到 zip 中，同时保持目录结构。

#### (*Writer) Close 

``` go 
func (w *Writer) Close() error
```

Close finishes writing the zip file by writing the central directory. It does not close the underlying writer.

​	Close 通过写入中心目录完成 zip 文件的写入。它不会关闭底层 writer。

#### (*Writer) Copy  <- go1.17

``` go 
func (w *Writer) Copy(f *File) error
```

Copy copies the file f (obtained from a Reader) into w. It copies the raw form directly bypassing decompression, compression, and validation.

​	Copy 将文件 f（从 Reader 获取）复制到 w 中。它直接复制原始格式，绕过解压缩、压缩和验证。

#### (*Writer) Create 

``` go 
func (w *Writer) Create(name string) (io.Writer, error)
```

Create adds a file to the zip file using the provided name. It returns a Writer to which the file contents should be written. The file contents will be compressed using the Deflate method. The name must be a relative path: it must not start with a drive letter (e.g. C:) or leading slash, and only forward slashes are allowed. To create a directory instead of a file, add a trailing slash to the name. The file's contents must be written to the io.Writer before the next call to Create, CreateHeader, or Close.

​	Create 使用提供的名称将文件添加到 zip 文件中。它返回一个 Writer，文件内容应写入其中。文件内容将使用 Deflate 方法进行压缩。名称必须是相对路径：它不能以驱动器号（例如 C:) 或前导斜杠开头，并且只允许使用正斜杠。要创建目录而不是文件，请在名称后添加尾随斜杠。文件的内容必须在下次调用 Create、CreateHeader 或 Close 之前写入 io.Writer。

#### (*Writer) CreateHeader 

``` go 
func (w *Writer) CreateHeader(fh *FileHeader) (io.Writer, error)
```

CreateHeader adds a file to the zip archive using the provided FileHeader for the file metadata. Writer takes ownership of fh and may mutate its fields. The caller must not modify fh after calling CreateHeader.

​	CreateHeader 使用提供的 FileHeader 作为文件元数据将文件添加到 zip 存档中。Writer 拥有 fh 并可能会改变其字段。调用 CreateHeader 后，调用者不得修改 fh。

This returns a Writer to which the file contents should be written. The file's contents must be written to the io.Writer before the next call to Create, CreateHeader, CreateRaw, or Close.

​	此方法返回一个 Writer，文件内容应写入其中。在下次调用 Create、CreateHeader、CreateRaw 或 Close 之前，必须将文件内容写入 io.Writer。

#### (*Writer) CreateRaw  <- go1.17

``` go 
func (w *Writer) CreateRaw(fh *FileHeader) (io.Writer, error)
```

CreateRaw adds a file to the zip archive using the provided FileHeader and returns a Writer to which the file contents should be written. The file's contents must be written to the io.Writer before the next call to Create, CreateHeader, CreateRaw, or Close.

​	CreateRaw 使用提供的 FileHeader 将文件添加到 zip 存档中，并返回一个 Writer，文件内容应写入其中。在下次调用 Create、CreateHeader、CreateRaw 或 Close 之前，必须将文件内容写入 io.Writer。

In contrast to CreateHeader, the bytes passed to Writer are not compressed.

​	与 CreateHeader 相反，传递给 Writer 的字节不会被压缩。

#### (*Writer) Flush  <- go1.4

``` go 
func (w *Writer) Flush() error
```

Flush flushes any buffered data to the underlying writer. Calling Flush is not normally necessary; calling Close is sufficient.

​	Flush 将任何缓冲数据刷新到底层 writer。通常不需要调用 Flush；调用 Close 就足够了。

#### (*Writer) RegisterCompressor  <- go1.6

``` go 
func (w *Writer) RegisterCompressor(method uint16, comp Compressor)
```

RegisterCompressor registers or overrides a custom compressor for a specific method ID. If a compressor for a given method is not found, Writer will default to looking up the compressor at the package level.

​	RegisterCompressor 为特定方法 ID 注册或覆盖自定义压缩器。如果找不到给定方法的压缩器，Writer 将默认在包级别查找压缩器。

##### RegisterCompressor  Example
``` go 
package main

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"io"
)

func main() {
	// Override the default Deflate compressor with a higher compression level.

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Register a custom Deflate compressor.
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// Proceed to add files to w.
}
Output:

```

#### (*Writer) SetComment  <- go1.10

``` go 
func (w *Writer) SetComment(comment string) error
```

SetComment sets the end-of-central-directory comment field. It can only be called before Close.

​	SetComment 设置中央目录末尾注释字段。它只能在 Close 之前调用。

#### (*Writer) SetOffset  <- go1.5

``` go 
func (w *Writer) SetOffset(n int64)
```

SetOffset sets the offset of the beginning of the zip data within the underlying writer. It should be used when the zip data is appended to an existing file, such as a binary executable. It must be called before any data is written.

​	SetOffset 设置基础 writer 中 zip 数据开始处的偏移量。当 zip 数据附加到现有文件（例如二进制可执行文件）时，应使用它。必须在写入任何数据之前调用它。