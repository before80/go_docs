+++
title = "tar"
date = 2023-05-17T09:59:21+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/archive/tar@go1.24.2](https://pkg.go.dev/archive/tar@go1.24.2)

Package tar implements access to tar archives.

​	tar包实现对tar存档的访问。

Tape archives (tar) are a file format for storing a sequence of files that can be read and written in a streaming manner. This package aims to cover most variations of the format, including those produced by GNU and BSD tar tools.

​	磁带存档（tar）是一种用于以流式方式读取和写入的存储文件序列的文件格式。该软件包旨在涵盖该格式的大多数变体，包括由GNU和BSD tar工具生成的那些。

## Example (Minimal)
``` go 
package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Create and add some files to the archive.
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatal(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	// Open and iterate through the files in the archive.
	tr := tar.NewReader(&buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}

}
Output:

Contents of readme.txt:
This archive contains some text files.
Contents of gopher.txt:
Gopher names:
George
Geoffrey
Gonzo
Contents of todo.txt:
Get animal handling license.
```







## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/archive/tar/common.go;l=60)

``` go 
const (
	// Type '0' indicates a regular file.
	TypeReg = '0'

	// Deprecated: Use TypeReg instead.
	TypeRegA = '\x00'

	// Type '1' to '6' are header-only flags and may not have a data body.
	TypeLink    = '1' // Hard link
	TypeSymlink = '2' // Symbolic link
	TypeChar    = '3' // Character device node
	TypeBlock   = '4' // Block device node
	TypeDir     = '5' // Directory
	TypeFifo    = '6' // FIFO node

	// Type '7' is reserved.
	TypeCont = '7'

	// Type 'x' is used by the PAX format to store key-value records that
	// are only relevant to the next file.
	// This package transparently handles these types.
	TypeXHeader = 'x'

	// Type 'g' is used by the PAX format to store key-value records that
	// are relevant to all subsequent files.
	// This package only supports parsing and composing such headers,
	// but does not currently support persisting the global state across files.
	TypeXGlobalHeader = 'g'

	// Type 'S' indicates a sparse file in the GNU format.
	TypeGNUSparse = 'S'

	// Types 'L' and 'K' are used by the GNU format for a meta file
	// used to store the path or link name for the next file.
	// This package transparently handles these types.
	TypeGNULongName = 'L'
	TypeGNULongLink = 'K'
)
```

Type flags for Header.Typeflag.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/archive/tar/common.go;l=32)

``` go 
var (
	ErrHeader          = errors.New("archive/tar: invalid tar header")
	ErrWriteTooLong    = errors.New("archive/tar: write too long")
	ErrFieldTooLong    = errors.New("archive/tar: header field too long")
	ErrWriteAfterClose = errors.New("archive/tar: write after close")
	ErrInsecurePath    = errors.New("archive/tar: insecure file path")
)
```

## 函数

This section is empty.

## 类型

### type FileInfoNames <- go1.23.0

``` go
type FileInfoNames interface {
	fs.FileInfo
	// Uname should give a user name.
	Uname() (string, error)
	// Gname should give a group name.
	Gname() (string, error)
}
```

FileInfoNames extends [fs.FileInfo](https://pkg.go.dev/io/fs#FileInfo). Passing an instance of this to [FileInfoHeader](https://pkg.go.dev/archive/tar@go1.23.0#FileInfoHeader) permits the caller to avoid a system-dependent name lookup by specifying the Uname and Gname directly.

​	FileInfoNames扩展[fs.FileInfo](https://pkg.go.dev/io/fs#FileInfo)。将此实例传递给[FileInfoHeader](https://pkg.go.dev/archive/tar@go1.23.0#FileInfoHeader)允许调用者通过直接指定Uname和Gname来避免与系统相关的名称查找。

### type Format  <- go1.10

``` go 
type Format int
```

Format represents the tar archive format.

​	Format 表示 tar 归档格式。

The original tar format was introduced in Unix V7. Since then, there have been multiple competing formats attempting to standardize or extend the V7 format to overcome its limitations. The most common formats are the USTAR, PAX, and GNU formats, each with their own advantages and limitations.

​	原始 tar 格式在 Unix V7 中引入。自那时起，出现了多种竞争格式，试图对 V7 格式进行标准化或扩展以克服其局限性。最常见的格式是 USTAR、PAX 和 GNU 格式，每个格式都有自己的优点和局限性。

The following table captures the capabilities of each format:

​	下表记录了每种格式的功能：

```
                  |  USTAR |       PAX |       GNU
------------------+--------+-----------+----------
Name              |   256B | unlimited | unlimited
Linkname          |   100B | unlimited | unlimited
Size              | uint33 | unlimited |    uint89
Mode              | uint21 |    uint21 |    uint57
Uid/Gid           | uint21 | unlimited |    uint57
Uname/Gname       |    32B | unlimited |       32B
ModTime           | uint33 | unlimited |     int89
AccessTime        |    n/a | unlimited |     int89
ChangeTime        |    n/a | unlimited |     int89
Devmajor/Devminor | uint21 |    uint21 |    uint57
------------------+--------+-----------+----------
string encoding   |  ASCII |     UTF-8 |    binary
sub-second times  |     no |       yes |        no
sparse files      |     no |       yes |       yes
```

The table's upper portion shows the Header fields, where each format reports the maximum number of bytes allowed for each string field and the integer type used to store each numeric field (where timestamps are stored as the number of seconds since the Unix epoch).

​	该表的顶部显示了 Header 字段，其中每种格式报告了每个字符串字段允许的最大字节数以及用于存储每个数字字段的整数类型（其中时间戳存储为自 Unix 纪元以来的秒数）。

The table's lower portion shows specialized features of each format, such as supported string encodings, support for sub-second timestamps, or support for sparse files.

​	该表的底部显示了每种格式的特殊功能，例如支持的字符串编码、对亚秒时间戳的支持或对稀疏文件支持。

The Writer currently provides no support for sparse files.

​	Writer 目前不支持稀疏文件。

``` go 
const (

	// FormatUnknown indicates that the format is unknown.
	FormatUnknown Format

	// FormatUSTAR represents the USTAR header format defined in POSIX.1-1988.
	//
	// While this format is compatible with most tar readers,
	// the format has several limitations making it unsuitable for some usages.
	// Most notably, it cannot support sparse files, files larger than 8GiB,
	// filenames larger than 256 characters, and non-ASCII filenames.
	//
	// Reference:
	//	http://pubs.opengroup.org/onlinepubs/9699919799/utilities/pax.html#tag_20_92_13_06
	FormatUSTAR

	// FormatPAX represents the PAX header format defined in POSIX.1-2001.
	//
	// PAX extends USTAR by writing a special file with Typeflag TypeXHeader
	// preceding the original header. This file contains a set of key-value
	// records, which are used to overcome USTAR's shortcomings, in addition to
	// providing the ability to have sub-second resolution for timestamps.
	//
	// Some newer formats add their own extensions to PAX by defining their
	// own keys and assigning certain semantic meaning to the associated values.
	// For example, sparse file support in PAX is implemented using keys
	// defined by the GNU manual (e.g., "GNU.sparse.map").
	//
	// Reference:
	//	http://pubs.opengroup.org/onlinepubs/009695399/utilities/pax.html
	FormatPAX

	// FormatGNU represents the GNU header format.
	//
	// The GNU header format is older than the USTAR and PAX standards and
	// is not compatible with them. The GNU format supports
	// arbitrary file sizes, filenames of arbitrary encoding and length,
	// sparse files, and other features.
	//
	// It is recommended that PAX be chosen over GNU unless the target
	// application can only parse GNU formatted archives.
	//
	// Reference:
	//	https://www.gnu.org/software/tar/manual/html_node/Standard.html
	FormatGNU
)
```

Constants to identify various tar formats.

​	用于标识各种 tar 格式的常量。

#### (Format) String  <- go1.10

``` go 
func (f Format) String() string
```

### type Header 

``` go 
type Header struct {
	// Typeflag is the type of header entry.
	// The zero value is automatically promoted to either TypeReg or TypeDir
	// depending on the presence of a trailing slash in Name.
	Typeflag byte

	Name     string // Name of file entry
	Linkname string // Target name of link (valid for TypeLink or TypeSymlink)

	Size  int64  // Logical file size in bytes
	Mode  int64  // Permission and mode bits
	Uid   int    // User ID of owner
	Gid   int    // Group ID of owner
	Uname string // User name of owner
	Gname string // Group name of owner

	// If the Format is unspecified, then Writer.WriteHeader rounds ModTime
	// to the nearest second and ignores the AccessTime and ChangeTime fields.
	//
	// To use AccessTime or ChangeTime, specify the Format as PAX or GNU.
	// To use sub-second resolution, specify the Format as PAX.
	ModTime    time.Time // Modification time
	AccessTime time.Time // Access time (requires either PAX or GNU support)
	ChangeTime time.Time // Change time (requires either PAX or GNU support)

	Devmajor int64 // Major device number (valid for TypeChar or TypeBlock)
	Devminor int64 // Minor device number (valid for TypeChar or TypeBlock)

	// Xattrs stores extended attributes as PAX records under the
	// "SCHILY.xattr." namespace.
	//
	// The following are semantically equivalent:
	//  h.Xattrs[key] = value
	//  h.PAXRecords["SCHILY.xattr."+key] = value
	//
	// When Writer.WriteHeader is called, the contents of Xattrs will take
	// precedence over those in PAXRecords.
	//
	// Deprecated: Use PAXRecords instead.
	Xattrs map[string]string

	// PAXRecords is a map of PAX extended header records.
	//
	// User-defined records should have keys of the following form:
	//	VENDOR.keyword
	// Where VENDOR is some namespace in all uppercase, and keyword may
	// not contain the '=' character (e.g., "GOLANG.pkg.version").
	// The key and value should be non-empty UTF-8 strings.
	//
	// When Writer.WriteHeader is called, PAX records derived from the
	// other fields in Header take precedence over PAXRecords.
	PAXRecords map[string]string

	// Format specifies the format of the tar header.
	//
	// This is set by Reader.Next as a best-effort guess at the format.
	// Since the Reader liberally reads some non-compliant files,
	// it is possible for this to be FormatUnknown.
	//
	// If the format is unspecified when Writer.WriteHeader is called,
	// then it uses the first format (in the order of USTAR, PAX, GNU)
	// capable of encoding this Header (see Format).
	Format Format
}
```

A Header represents a single header in a tar archive. Some fields may not be populated.

​	Header 表示 tar 存档中的单个标头。某些字段可能未填充。

For forward compatibility, users that retrieve a Header from Reader.Next, mutate it in some ways, and then pass it back to Writer.WriteHeader should do so by creating a new Header and copying the fields that they are interested in preserving.

​	为了向前兼容，从 Reader.Next 检索 Header、以某种方式对其进行变异，然后将其传回 Writer.WriteHeader 的用户应通过创建新的 Header 并复制他们感兴趣的要保留的字段来执行此操作。

#### func FileInfoHeader  <- go1.1

``` go 
func FileInfoHeader(fi fs.FileInfo, link string) (*Header, error)
```

FileInfoHeader creates a partially-populated Header from fi. If fi describes a symlink, FileInfoHeader records link as the link target. If fi describes a directory, a slash is appended to the name.

​	FileInfoHeader 从 fi 创建部分填充的 Header。如果 fi 描述了一个符号链接，FileInfoHeader 将记录 link 作为链接目标。如果 fi 描述了一个目录，则在名称后附加一个斜杠。

Since fs.FileInfo's Name method only returns the base name of the file it describes, it may be necessary to modify Header.Name to provide the full path name of the file.

​	由于 fs.FileInfo 的 Name 方法仅返回它描述的文件的基本名称，因此可能需要修改 Header.Name 以提供文件的完整路径名。

#### (*Header) FileInfo  <- go1.1

``` go 
func (h *Header) FileInfo() fs.FileInfo
```

FileInfo returns an fs.FileInfo for the Header.

​	FileInfo 为 Header 返回一个 fs.FileInfo。

### type Reader 

``` go 
type Reader struct {
	// contains filtered or unexported fields
}
```

Reader provides sequential access to the contents of a tar archive. Reader.Next advances to the next file in the archive (including the first), and then Reader can be treated as an io.Reader to access the file's data.

​	Reader 提供对 tar 存档内容的顺序访问。Reader.Next 进入存档中的下一个文件（包括第一个文件），然后 Reader 可被视为 io.Reader 以访问文件数据。

#### func NewReader 

``` go 
func NewReader(r io.Reader) *Reader
```

NewReader creates a new Reader reading from r.

​	NewReader 创建一个新的 Reader 从 r 读取。

#### (*Reader) Next 

``` go 
func (tr *Reader) Next() (*Header, error)
```

Next advances to the next entry in the tar archive. The Header.Size determines how many bytes can be read for the next file. Any remaining data in the current file is automatically discarded. At the end of the archive, Next returns the error io.EOF.

​	Next 进入 tar 存档中的下一个条目。Header.Size 确定可为下一个文件读取多少字节。当前文件中的任何剩余数据都会自动丢弃。在存档结束时，Next 返回错误 io.EOF。

If Next encounters a non-local name (as defined by [filepath.IsLocal](https://pkg.go.dev/path/filepath#IsLocal)) and the GODEBUG environment variable contains `tarinsecurepath=0`, Next returns the header with an ErrInsecurePath error. A future version of Go may introduce this behavior by default. Programs that want to accept non-local names can ignore the ErrInsecurePath error and use the returned header.

​	如果 Next 遇到非本地名称（由 filepath.IsLocal 定义），并且 GODEBUG 环境变量包含 `tarinsecurepath=0` ，则 Next 返回带有 ErrInsecurePath 错误的标头。未来版本的 Go 可能会默认引入此行为。想要接受非本地名称的程序可以忽略 ErrInsecurePath 错误并使用返回的标头。

#### (*Reader) Read 

``` go 
func (tr *Reader) Read(b []byte) (int, error)
```

Read reads from the current file in the tar archive. It returns (0, io.EOF) when it reaches the end of that file, until Next is called to advance to the next file.

​	Read 从 tar 存档中的当前文件读取。当它到达该文件的末尾时，它返回 (0, io.EOF)，直到调用 Next 进入下一个文件。

If the current file is sparse, then the regions marked as a hole are read back as NUL-bytes.

​	如果当前文件是稀疏的，则标记为孔的区域将被读回为 NUL 字节。

Calling Read on special types like TypeLink, TypeSymlink, TypeChar, TypeBlock, TypeDir, and TypeFifo returns (0, io.EOF) regardless of what the Header.Size claims.

​	对 TypeLink、TypeSymlink、TypeChar、TypeBlock、TypeDir 和 TypeFifo 等特殊类型的 Read 调用返回 (0, io.EOF)，而不管 Header.Size 声称是什么。

### type Writer 

``` go 
type Writer struct {
	// contains filtered or unexported fields
}
```

Writer provides sequential writing of a tar archive. Write.WriteHeader begins a new file with the provided Header, and then Writer can be treated as an io.Writer to supply that file's data.

​	Writer 提供 tar 存档的顺序写入。Write.WriteHeader 使用提供的 Header 开始一个新文件，然后 Writer 可以作为 io.Writer 来提供该文件的数据。

#### func NewWriter 

``` go 
func NewWriter(w io.Writer) *Writer
```

NewWriter creates a new Writer writing to w.

​	NewWriter 函数创建一个新的 Writer，写入 w。

#### (*Writer) AddFS <-go1.22.0

``` go
func (tw *Writer) AddFS(fsys fs.FS) error
```

AddFS adds the files from fs.FS to the archive. It walks the directory tree starting at the root of the filesystem adding each file to the tar archive while maintaining the directory structure.

​	AddFS 方法将 fs.FS 中的文件添加到存档中。它从文件系统的根目录开始遍历目录树，将每个文件添加到 tar 存档中，同时保持目录结构。

#### (*Writer) Close 

``` go 
func (tw *Writer) Close() error
```

Close closes the tar archive by flushing the padding, and writing the footer. If the current file (from a prior call to WriteHeader) is not fully written, then this returns an error.

​	Close 通过刷新填充并写入页脚来关闭 tar 存档。如果当前文件（来自 WriteHeader 的先前调用）未完全写入，则会返回错误。

#### (*Writer) Flush 

``` go 
func (tw *Writer) Flush() error
```

Flush finishes writing the current file's block padding. The current file must be fully written before Flush can be called.

​	Flush 完成写入当前文件的块填充。在调用 Flush 之前，必须完全写入当前文件。

This is unnecessary as the next call to WriteHeader or Close will implicitly flush out the file's padding.

​	这是不必要的，因为对 WriteHeader 或 Close 的下一次调用会隐式刷新文件的填充。

#### (*Writer) Write 

``` go 
func (tw *Writer) Write(b []byte) (int, error)
```

Write writes to the current file in the tar archive. Write returns the error ErrWriteTooLong if more than Header.Size bytes are written after WriteHeader.

​	Write 将内容写入 tar 存档中的当前文件。如果在 WriteHeader 之后写入的字节数超过 Header.Size，Write 将返回错误 ErrWriteTooLong。

Calling Write on special types like TypeLink, TypeSymlink, TypeChar, TypeBlock, TypeDir, and TypeFifo returns (0, ErrWriteTooLong) regardless of what the Header.Size claims.

​	对特殊类型（如 TypeLink、TypeSymlink、TypeChar、TypeBlock、TypeDir 和 TypeFifo）调用 Write 将返回 (0, ErrWriteTooLong)，而不管 Header.Size 声称是什么。

#### (*Writer) WriteHeader 

``` go 
func (tw *Writer) WriteHeader(hdr *Header) error
```

WriteHeader writes hdr and prepares to accept the file's contents. The Header.Size determines how many bytes can be written for the next file. If the current file is not fully written, then this returns an error. This implicitly flushes any padding necessary before writing the header.

​	WriteHeader 编写 hdr 并准备接受文件内容。Header.Size 决定了下一个文件可以写入多少字节。如果当前文件没有完全写入，则会返回一个错误。这会隐式刷新在写入头之前所需的任何填充。