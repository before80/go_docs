+++
title = "os"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++
> 原文：[https://pkg.go.dev/os@go1.24.2](https://pkg.go.dev/os@go1.24.2)

Package os provides a platform-independent interface to operating system functionality. The design is Unix-like, although the error handling is Go-like; failing calls return values of type error rather than error numbers. Often, more information is available within the error. For example, if a call that takes a file name fails, such as Open or Stat, the error will include the failing file name when printed and will be of type *PathError, which may be unpacked for more information.

​	`os` 包提供了一个独立于平台的操作系统功能接口。设计类似于 Unix，尽管错误处理类似于 Go；失败的调用会返回类型为 error 而不是错误号的值。通常，在错误中还有更多的信息。例如，如果一个以文件名为参数的调用失败了，比如 Open 或 Stat，错误消息将在打印时包括失败的文件名，并且类型为 `*PathError`，可以拆开以获取更多信息。

The os interface is intended to be uniform across all operating systems. Features not generally available appear in the system-specific package syscall.

​	os 接口旨在在所有操作系统中保持统一。一般不可用的功能会出现在特定于系统的 syscall 包中。

Here is a simple example, opening a file and reading some of it.

​	这里是一个简单的例子，打开一个文件并读取其中一部分。

```go
file, err := os.Open("file.go") // 用于读取。
if err != nil {
	log.Fatal(err)
}
```

If the open fails, the error string will be self-explanatory, like

​	如果打开失败，错误字符串将是不言自明的，例如

```bash
open file.go: no such file or directory
```

The file's data can then be read into a slice of bytes. Read and Write take their byte counts from the length of the argument slice.

​	然后可以将文件的数据读入到一个字节切片中。Read 和 Write 的字节计数从参数切片的长度中获取。

```go
data := make([]byte, 100)
count, err := file.Read(data)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])
```

Note: The maximum number of concurrent operations on a File may be limited by the OS or the system. The number should be high, but exceeding it may degrade performance or cause other issues.

注意：File 上的最大并发操作数可能受到操作系统或系统的限制。该数字应该很高，但超出它可能会降低性能或引起其他问题。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/file.go;l=72)

``` go 
const (
    // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	// 必须精确指定 O_RDONLY、O_WRONLY 或 O_RDWR 中的一个。
	O_RDONLY int = syscall.O_RDONLY // 只读打开文件。 open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // 只写打开文件。 open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // 读写打开文件。 open the file read-write.
    // The remaining values may be or'ed in to control behavior.
	// 其余的值可以进行或运算来控制行为。
	O_APPEND int = syscall.O_APPEND // 写入数据时追加到文件中。 append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // 如果不存在则创建一个新文件。 create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // 与 O_CREATE 一起使用，文件必须不存在。 used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // 为同步 I/O 打开文件。 open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // 打开普通可写文件时截断。 truncate regular writable file when opened.
)
```

Flags to OpenFile wrapping those of the underlying system. Not all flags may be implemented on a given system.

​	OpenFile 的标志，包装了底层系统的标志。并非所有的标志都可能在给定系统上实现。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/file.go;l=88)

``` go 
const (
	SEEK_SET int = 0 // 从文件的起始位置寻址
	SEEK_CUR int = 1 // 相对于当前位置寻址
	SEEK_END int = 2 // 相对于文件末尾寻址
)
```

Seek whence values.

​	寻址偏移值。

Deprecated: Use io.SeekStart, io.SeekCurrent, and io.SeekEnd.

​	已弃用：请使用 io.SeekStart、io.SeekCurrent 和 io.SeekEnd。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/path_unix.go;l=9)

``` go 
const (
	PathSeparator     = '/' // 特定于操作系统的路径分隔符
	PathListSeparator = ':' // 特定于操作系统的路径列表分隔符
)
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/types.go;l=35)

``` go 
const (
    // The single letters are the abbreviations
	// used by the String method's formatting.
	// 单个字母是 String 方法格式化时使用的缩写。
	ModeDir        = fs.ModeDir        // d：目录 d: is a directory
	ModeAppend     = fs.ModeAppend     // a：仅追加 a: append-only
	ModeExclusive  = fs.ModeExclusive  // l：独占使用 l: exclusive use
	ModeTemporary  = fs.ModeTemporary  // T：临时文件；仅 Plan 9 T: temporary file; Plan 9 only
	ModeSymlink    = fs.ModeSymlink    // L：符号链接 L: symbolic link
	ModeDevice     = fs.ModeDevice     // D：设备文件 D: device file
	ModeNamedPipe  = fs.ModeNamedPipe  // p：命名管道(FIFO) p: named pipe (FIFO)
	ModeSocket     = fs.ModeSocket     // S：Unix 域套接字 S: Unix domain socket
	ModeSetuid     = fs.ModeSetuid     // u: setuid u: setuid
	ModeSetgid     = fs.ModeSetgid     // g: setgid g: setgid
	ModeCharDevice = fs.ModeCharDevice // c：Unix 字符设备，在设置 ModeDevice 时 c: Unix character device, when ModeDevice is set
	ModeSticky     = fs.ModeSticky     // t：粘性位  t: sticky
	ModeIrregular  = fs.ModeIrregular  // ?: 非常规文件；关于此文件没有其他已知信息 ?: non-regular file; nothing else is known about this file

    // Mask for the type bits. For regular files, none will be set.
	// 类型比特的掩码。对于普通文件，不会设置任何位。
	ModeType = fs.ModeType

	ModePerm = fs.ModePerm //  Unix 权限位，0o777 Unix permission bits, 0o777
)
```

The defined file mode bits are the most significant bits of the FileMode. The nine least-significant bits are the standard Unix rwxrwxrwx permissions. The values of these bits should be considered part of the public API and may be used in wire protocols or disk representations: they must not be changed, although new bits might be added.

​	已定义的文件模式比特是 FileMode 的最高有效位。最低有效位是标准 Unix rwxrwxrwx 权限位。这些位的值应该被视为公共 API 的一部分，可以在传输协议或磁盘表示中使用：不能更改它们，尽管可以添加新的位。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/file_unix.go;l=211)

``` go 
const DevNull = "/dev/null"
```

DevNull is the name of the operating system's “null device.” On Unix-like systems, it is "/dev/null"; on Windows, "NUL".

​	DevNull 是操作系统的"null 设备"的名称。在类 Unix 系统上，它是"/dev/null"；在 Windows 上，它是"NUL"。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/error.go;l=16)

``` go 
var (
    // ErrInvalid indicates an invalid argument.
	// Methods on File will return this error when the receiver is nil.
	// ErrInvalid 表示无效的参数
	// 如果接收器为 nil，则 File 上的方法将返回此错误。
	ErrInvalid = fs.ErrInvalid // "invalid argument"

	ErrPermission = fs.ErrPermission // "permission denied"
	ErrExist      = fs.ErrExist      // "file already exists"
	ErrNotExist   = fs.ErrNotExist   // "file does not exist"
	ErrClosed     = fs.ErrClosed     // "file already closed"

	ErrNoDeadline       = errNoDeadline()       // "file type does not support deadline"// "文件类型不支持期限"
	ErrDeadlineExceeded = errDeadlineExceeded() // "i/o timeout"
)
```

Portable analogs of some common system call errors.

​	一些常见系统调用错误的便携式类比。

Errors returned from this package may be tested against these errors with errors.Is.

​	此包返回的错误可以使用 errors.Is 进行测试。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/file.go;l=64)

``` go 
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```

Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.

​	Stdin、Stdout 和 Stderr 是打开的文件，它们指向标准输入、标准输出和标准错误文件描述符。

Note that the Go runtime writes to standard error for panics and crashes; closing Stderr may cause those messages to go elsewhere, perhaps to a file opened later.

​	请注意，Go 运行时会将 panics 和 crashes 的消息写入标准错误；关闭 Stderr 可能会导致这些消息转到其他位置，例如稍后打开的文件。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/proc.go;l=16)

``` go 
var Args []string
```

Args hold the command-line arguments, starting with the program name.

​	Args 保存命令行参数，从程序名开始。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/exec.go;l=18)

``` go 
var ErrProcessDone = errors.New("os: process already finished")
```

ErrProcessDone indicates a Process has finished.

​	ErrProcessDone 表示进程已经结束。

## 函数

### func Chdir 

``` go 
func Chdir(dir string) error
```

Chdir changes the current working directory to the named directory. If there is an error, it will be of type *PathError.

​	Chdir函数将当前工作目录切换为指定的目录。如果出错，将返回 `*PathError` 类型的错误。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败：", err)
		return
	}
	fmt.Println("当前工作目录：", currentDir)

	// 切换到指定目录
	err = os.Chdir("dir/subdir")
	if err != nil {
		fmt.Println("切换目录失败：", err)
		return
	}

	// 再次获取当前工作目录
	currentDir, err = os.Getwd()
	if err != nil {
		fmt.Println("获取当前工作目录失败：", err)
		return
	}
	fmt.Println("切换后的工作目录：", currentDir)
}
// Output:
//当前工作目录： F:\Devs\MyCodes\go_std_examples\os\os_self\f_Chdir
//切换后的工作目录： F:\Devs\MyCodes\go_std_examples\os\os_self\f_Chdir\dir\subdir
```



### func Chmod 

``` go 
func Chmod(name string, mode FileMode) error
```

Chmod changes the mode of the named file to mode. If the file is a symbolic link, it changes the mode of the link's target. If there is an error, it will be of type *PathError.

​	Chmod函数将指定文件的模式更改为 mode。如果该文件是符号链接，则更改链接目标的模式。如果出错，将返回 `*PathError` 类型的错误。

A different subset of the mode bits are used, depending on the operating system.

​	根据操作系统使用不同的模式比特的子集。

On Unix, the mode's permission bits, ModeSetuid, ModeSetgid, and ModeSticky are used.

​	在 Unix 上，使用 mode 的权限位 ModeSetuid、ModeSetgid 和 ModeSticky。

On Windows, only the 0200 bit (owner writable) of mode is used; it controls whether the file's read-only attribute is set or cleared. The other bits are currently unused. For compatibility with Go 1.12 and earlier, use a non-zero mode. Use mode 0400 for a read-only file and 0600 for a readable+writable file.

​	在 Windows 上，仅使用 mode 的 0200 位(所有者可写)；它控制文件的只读属性是设置还是清除。其他位当前未使用。为了与 Go 1.12 及更早版本兼容，请使用非零模式。对于只读文件，请使用 mode 0400，对于可读写文件，请使用 mode 0600。

On Plan 9, the mode's permission bits, ModeAppend, ModeExclusive, and ModeTemporary are used.

​	在 Plan 9 上，使用 mode 的权限位 ModeAppend、ModeExclusive 和 ModeTemporary。

```go 
package main

import (
	"log"
	"os"
)

func main() {
	if err := os.Chmod("some-filename", 0644); err != nil {
		log.Fatal(err)
	}
}

// Output:

// 2009/11/10 23:00:00 chmod some-filename: no such file or directory

```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.OpenFile("data.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开文件失败：", err)
		return
	}
	defer file.Close()

	// 获取文件的当前权限
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败：", err)
		return
	}
	fmt.Println("当前文件权限：", fileInfo.Mode().String())

	// 修改文件权限
	err = os.Chmod("data.txt", 0755)
	if err != nil {
		fmt.Println("修改文件权限失败：", err)
		return
	}
	fmt.Println("文件权限修改成功！")

	// 再次获取文件的权限
	fileInfo, err = file.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败：", err)
		return
	}
	fmt.Println("修改后的文件权限：", fileInfo.Mode().String())
}

// Output:
//当前文件权限： -rw-r--r--
//文件权限修改成功！
//修改后的文件权限： -rwxr-xr-x
```



### func Chown 

``` go 
func Chown(name string, uid, gid int) error
```

Chown changes the numeric uid and gid of the named file. If the file is a symbolic link, it changes the uid and gid of the link's target. A uid or gid of -1 means to not change that value. If there is an error, it will be of type *PathError.

​	Chown函数更改指定文件的数值 UID 和 GID。如果该文件是符号链接，则更改链接目标的 UID 和 GID。UID 或 GID 的值为 -1 表示不更改该值。如果出错，将返回 *PathError 类型的错误。

On Windows or Plan 9, Chown always returns the syscall.EWINDOWS or EPLAN9 error, wrapped in *PathError.

​	在 Windows 或 Plan 9 上，Chown 始终返回 syscall.EWINDOWS 或 EPLAN9 错误，包装在 *PathError 中。

```go
package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	// 使用 root 用户进行创建：
	// root 用户的ID 是 0，用户的组ID 是 0
	// lx 用户的ID 是 1000，用户的组ID 是 1000 和 1001
	filePath := "data.txt"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return
	}
	PrintFileInfo(filePath)
	file.Close()

	// 指定新的所有者和组
	uid := 1001 // 这是我新增的test_01用户的ID
	gid := 1001 // 这是我新增的test_01用户的组ID

	// 使用os.Chown更改文件的所有者和组
	if err = os.Chown(filePath, uid, gid); err != nil {
		fmt.Println("无法更改文件所有者和组:", err)
		return
	}
	fmt.Println("文件：" + filePath + "的所有者和组已成功更改")
	PrintFileInfo(filePath)
}

func PrintFileInfo(filepath string) {
	// 获取文件信息
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("无法获取文件信息:", err)
		return
	}

	// 获取文件的用户ID和用户组ID
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatal("类型错误")
	}

	uid := stat.Uid
	gid := stat.Gid
	fmt.Printf("文件：%s，当前的所属用户ID：%d，所属用户组ID：%d\n", filepath, uid, gid)
}
// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Chown$ go build Chown.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Chown$ sudo ./Chown
//文件：data.txt，当前的所属用户ID：0，所属用户组ID：0
//文件：data.txt的所有者和组已成功更改
//文件：data.txt，当前的所属用户ID：1001，所属用户组ID：1001
```



### func Chtimes 

``` go 
func Chtimes(name string, atime time.Time, mtime time.Time) error
```

Chtimes changes the access and modification times of the named file, similar to the Unix utime() or utimes() functions.

​	Chtimes函数更改指定文件的访问和修改时间，类似于 Unix 的 utime() 或 utimes() 函数。

The underlying filesystem may truncate or round the values to a less precise time unit. If there is an error, it will be of type *PathError.

​	底层文件系统可能会将值截断或舍入为较不精确的时间单位。如果出错，将返回 `*PathError` 类型的错误。

```go 
package main

import (
	"log"
	"os"
	"time"
)

func main() {
	mtime := time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC)
	atime := time.Date(2007, time.March, 2, 4, 5, 6, 0, time.UTC)
	if err := os.Chtimes("some-filename", atime, mtime); err != nil {
		log.Fatal(err)
	}
}

// Output:

// 2009/11/10 23:00:00 chtimes some-filename: no such file or directory
```

```go
package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	filePath := "data.txt"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("创建文件失败：", err)
		return
	}
	PrintFileInfo(filePath)
	file.Close()

	mtime := time.Date(2006, 2, 1, 3, 4, 5, 0, time.Local)
	atime := time.Date(2007, 3, 2, 4, 5, 6, 0, time.Local)
	if err := os.Chtimes(filePath, atime, mtime); err != nil {
		log.Fatal(err)
	}
	fmt.Println("使用Chtimes()之后...")
	PrintFileInfo(filePath)
}

func PrintFileInfo(filepath string) {
	// 获取文件信息
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("无法获取文件信息:", err)
		return
	}

	// 获取文件的用户ID和用户组ID
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		log.Fatal("类型错误")
	}
	//fmt.Printf("%#v\n", stat)
	fmt.Printf("文件：%s，当前的修改时间：%s，访问时间：%s\n", filepath, TransformTimeFormat(stat.Mtim), TransformTimeFormat(stat.Atim))
}

func TransformTimeFormat(ts syscall.Timespec) string {
	t := time.Unix(ts.Sec, ts.Nsec)
	// 格式化时间为所需的字符串格式
	return t.Format("2006-01-02 15:04:05")
}

// Output:
//文件：data.txt，当前的修改时间：2023-08-28 11:28:52，访问时间：2023-08-28 11:28:52
//使用Chtimes()之后...
//文件：data.txt，当前的修改时间：2006-02-01 03:04:05，访问时间：2007-03-02 04:05:06
```



### func Clearenv 

``` go 
func Clearenv()
```

Clearenv deletes all environment variables.

​	Clearenv函数删除所有环境变量。

> 个人注释
>
> Clearenv 函数只会清除当前进程的环境变量，不会影响系统中的其他进程的环境变量。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 设置环境变量
	os.Setenv("FOO", "bar")
	os.Setenv("BAR", "baz")

	// 打印环境变量
	fmt.Println("环境变量：", os.Environ())

	// 清除环境变量
	os.Clearenv()

	// 再次打印环境变量
	fmt.Println("环境变量：", os.Environ())
}
// 以下是连续执行两次的结果：
// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Clearenv$ go run Clearenv.go
//环境变量： [SHELL=/bin/bash WSL2_GUI_APPS_ENABLED=1 WSL_DISTRO_NAME=Ubuntu-22.04 ...还有很多... GOPATH=/mnt/f/GoPath:/home/lx/go _=/usr/local/go/bin/go FOO=bar BAR=baz]
//环境变量： []
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Clearenv$ go run Clearenv.go
//环境变量： [SHELL=/bin/bash WSL2_GUI_APPS_ENABLED=1 WSL_DISTRO_NAME=Ubuntu-22.04 ...还有很多... GOPATH=/mnt/f/GoPath:/home/lx/go _=/usr/local/go/bin/go FOO=bar BAR=baz]
//环境变量： []
```

### func CopyFS <- go1.23.0

``` go
func CopyFS(dir string, fsys fs.FS) error
```

CopyFS copies the file system fsys into the directory dir, creating dir if necessary.

​	CopyFS 将文件系统 fsys 复制到目录 dir 中，如果需要会创建 dir。

Files are created with mode 0o666 plus any execute permissions from the source, and directories are created with mode 0o777 (before umask).

​	文件会以模式 0o666 加上源文件的任何执行权限来创建，目录会以模式 0o777 创建（在应用 umask 之前）。

CopyFS will not overwrite existing files, and returns an error if a file name in fsys already exists in the destination.

​	CopyFS 不会覆盖已有文件，如果 fsys 中的文件名在目标目录中已经存在，则返回错误。

Symbolic links in fsys are not supported. A *PathError with Err set to ErrInvalid is returned when copying from a symbolic link.

​	不支持复制 fsys 中的符号链接。如果复制符号链接，将返回一个带有 ErrInvalid 的 *PathError。

Symbolic links in dir are followed.

​	在 dir 中的符号链接会被跟随。

Copying stops at and returns the first error encountered.

​	复制过程中遇到的第一个错误将停止复制并返回该错误。

### func DirFS  <- go1.16

``` go 
func DirFS(dir string) fs.FS
```

DirFS returns a file system (an fs.FS) for the tree of files rooted at the directory dir.

​	DirFS 函数返回以目录 dir 为根的文件树的文件系统(即 fs.FS)。

Note that DirFS("/prefix") only guarantees that the Open calls it makes to the operating system will begin with "/prefix": DirFS("/prefix").Open("file") is the same as os.Open("/prefix/file"). So if /prefix/file is a symbolic link pointing outside the /prefix tree, then using DirFS does not stop the access any more than using os.Open does. Additionally, the root of the fs.FS returned for a relative path, DirFS("prefix"), will be affected by later calls to Chdir. DirFS is therefore not a general substitute for a chroot-style security mechanism when the directory tree contains arbitrary content.

​	请注意，DirFS("/prefix") 只保证它对操作系统所做的 Open 调用将以 "/prefix" 开始：DirFS("/prefix").Open("file") 与 os.Open("/prefix/file") 相同。因此，如果 `/prefix/file` 是指向 `/prefix` 以外的符号链接，则使用 DirFS 不会比使用 os.Open 更停止访问。此外，对于相对路径，DirFS 返回的 fs.FS 的根目录，即 DirFS("prefix")，将受后续 Chdir 调用的影响。因此，当目录树包含任意内容时，DirFS 不是 chroot 类型的安全机制的通用替代品。

The directory dir must not be "".

​	目录 dir 不得为 ""。

The result implements fs.StatFS.

​	该结果实现了 fs.StatFS。

![image-20230828121722548](os_img/image-20230828121722548.png)

```go
package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

func main() {
	fs0 := os.DirFS("dir")
	file0, err := fs0.Open("0.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file0.Close()
	fileData, err := io.ReadAll(file0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("文件内容：", string(fileData))

	dirEntrys, err := fs.ReadDir(os.DirFS("dir"), ".")

	if err != nil {
		fmt.Println(err)
		return
	}

	for i, dirEntry := range dirEntrys {
		fmt.Println(i, "--------------------------")
		fmt.Println("名称:", dirEntry.Name())
		fmt.Println("类型:", dirEntry.Type())
		fmt.Println("是目录？", dirEntry.IsDir())
		info, err := dirEntry.Info()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("dirEntry.Info()=%#v\n", info)

		fmt.Println("info.Name()=", info.Name())
		fmt.Println("info.Size()=", info.Size())
		fmt.Println("info.Mode()=", info.Mode())
		fmt.Println("info.ModTime()=", info.ModTime())
		fmt.Println("info.IsDir()=", info.IsDir())
		fmt.Printf("info.Sys()=%#v\n", info.Sys())
	}
}
// Output:
//文件内容： content0
//0 --------------------------
//名称: 0.txt
//类型: ----------
//是目录？ false
//dirEntry.Info()=&os.fileStat{name:"0.txt", size:8, mode:0x1a4, modTime:time.Time{wall:0x2b8366b8, ext:63828778161, loc:(*time.Location)(0x547f00)}, sys:syscall.Stat_t{Dev:0x820, Ino:0xb068, Nlink:0x1, Mode:0x81a4, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:8, Blksize:4096, Blocks
//:8, Atim:syscall.Timespec{Sec:1693194959, Nsec:290014206}, Mtim:syscall.Timespec{Sec:1693181361, Nsec:730031800}, Ctim:syscall.Timespec{Sec:1693194959, Nsec:220014206}, X__unused:[3]int64{0, 0, 0}}}
//info.Name()= 0.txt
//info.Size()= 8
//info.Mode()= -rw-r--r--
//info.ModTime()= 2023-08-28 08:09:21.7300318 +0800 CST
//info.IsDir()= false
//info.Sys()=&syscall.Stat_t{Dev:0x820, Ino:0xb068, Nlink:0x1, Mode:0x81a4, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:8, Blksize:4096, Blocks:8, Atim:syscall.Timespec{Sec:1693194959, Nsec:290014206}, Mtim:syscall.Timespec{Sec:1693181361, Nsec:730031800}, Ctim:syscall.Timespec{Sec:
//1693194959, Nsec:220014206}, X__unused:[3]int64{0, 0, 0}}
//1 --------------------------
//名称: subdir1
//类型: d---------
//是目录？ true
//dirEntry.Info()=&os.fileStat{name:"subdir1", size:4096, mode:0x800001ed, modTime:time.Time{wall:0xe4e537e, ext:63828791759, loc:(*time.Location)(0x547f00)}, sys:syscall.Stat_t{Dev:0x820, Ino:0xb06c, Nlink:0x2, Mode:0x41ed, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:4096, Blksize:
//4096, Blocks:8, Atim:syscall.Timespec{Sec:1693194959, Nsec:250014206}, Mtim:syscall.Timespec{Sec:1693194959, Nsec:240014206}, Ctim:syscall.Timespec{Sec:1693194959, Nsec:240014206}, X__unused:[3]int64{0, 0, 0}}}
//info.Name()= subdir1
//info.Size()= 4096
//info.Size()= 4096
//info.Mode()= drwxr-xr-x
//info.ModTime()= 2023-08-28 11:55:59.260014206 +0800 CST
//info.IsDir()= true
//info.Sys()=&syscall.Stat_t{Dev:0x820, Ino:0xb071, Nlink:0x2, Mode:0x41ed, Uid:0x3e8, Gid:0x3e8, X__pad0:0, Rdev:0x0, Size:4096, Blksize:4096, Blocks:8, Atim:syscall.Timespec{Sec:1693194959, Nsec:910014210}, Mtim:syscall.Timespec{Sec:1693194959, Nsec:260014206}, Ctim:syscall.Timespec{S
//ec:1693194959, Nsec:260014206}, X__unused:[3]int64{0, 0, 0}}
```

​	参见[WalkDir My Example]({{< ref "/stdLib/io/fs#walkdir-my-example">}})

### func Environ 

``` go 
func Environ() []string
```

Environ returns a copy of strings representing the environment, in the form "key=value".

​	Environ函数返回表示环境变量的字符串副本，格式为"key=value"。

```go
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// 设置环境变量
	os.Setenv("TEST1", "bar")
	os.Setenv("TEST2", "baz")

	re, err := regexp.Compile(`^TEST\d{1}`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("匹配 `^TEST\\d{1}` 模式的环境变量有：")
	matchNum := 0
	noMatchNum := 0
	for _, env := range os.Environ() {
		if re.MatchString(env) {
			fmt.Println(env)
			matchNum++
		} else {
			noMatchNum++
		}
	}
	fmt.Printf("匹配的环境变量个数：%d，不匹配的环境变量个数：%d\n", matchNum, noMatchNum)
}

// Output:
//匹配 `^TEST\d{1}` 模式的环境变量有：
//TEST1=bar
//TEST2=baz
//匹配的环境变量个数：2，不匹配的环境变量个数：29
```

​	参见[func Clearenv](#func-clearenv) 

### func Executable  <- go1.8

``` go 
func Executable() (string, error)
```

Executable returns the path name for the executable that started the current process. There is no guarantee that the path is still pointing to the correct executable. If a symlink was used to start the process, depending on the operating system, the result might be the symlink or the path it pointed to. If a stable result is needed, path/filepath.EvalSymlinks might help.

​	Executable函数返回启动当前进程的可执行文件的路径名。不能保证路径仍指向正确的可执行文件。如果使用符号链接启动了进程，则根据操作系统，结果可能是符号链接或它所指向的路径。如果需要稳定的结果，`path/filepath.EvalSymlinks`可能有所帮助。

Executable returns an absolute path unless an error occurred.

​	Executable返回绝对路径，除非发生错误。

The main use case is finding resources located relative to an executable.

​	主要用例是查找相对于可执行文件的资源。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取可执行文件的路径
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印可执行文件的路径
	fmt.Println("可执行文件路径：", path)
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Executable$ go run Executable.go
//可执行文件路径： /tmp/go-build4226880423/b001/exe/Executable
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Executable$ go build Executable.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Executable$ sudo ./Executable
//[sudo] password for lx:
//可执行文件路径： /home/lx/goprojects/go_std_examples/os/os_self/f_Executable/Executable
```



### func Exit 

``` go 
func Exit(code int)
```

Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately; deferred functions are not run.

​	Exit函数使当前程序以给定的状态码退出。传统上，代码零表示成功，非零表示错误。程序立即终止；延迟函数不会被运行。

For portability, the status code should be in the range [0, 125].

​	为了可移植性，状态码应在[0，125]范围内。

```go
package main

import (
	"fmt"
	"os"
)

func Exit0() {
	defer func() {
		fmt.Println("可能你认为os.Exit(0)之后会执行 defer，但你想错了！被defer的函数也是没有执行的机会，不信你试试！")
	}()
	os.Exit(0)
	fmt.Println("这句话肯定打印不出来，不信你试试！")
}

func main() {
	Exit0()
}

// Output:
// 没有任何输出
```

```go
package main

import (
	"fmt"
	"os"
)

func Exit1() {
	defer func() {
		fmt.Println("可能你认为os.Exit(1)之后会执行 defer，但你想错了！被defer的函数也是没有执行的机会，不信你试试！")
	}()
	os.Exit(1)
	fmt.Println("这句话肯定打印不出来，不信你试试！")
}

func main() {
	Exit1()
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Exit1$ go run Exit.go
//exit status 1
```

```go
package main

import (
	"fmt"
	"os"
)

func Exit2() {
	defer func() {
		fmt.Println("可能你认为os.Exit(2)之后会执行 defer，但你想错了！被defer的函数也是没有执行的机会，不信你试试！")
	}()
	os.Exit(2)
	fmt.Println("这句话肯定打印不出来，不信你试试！")
}
func main() {
	Exit2()
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Exit2$ go run Exit.go
//exit status 2
```



### func Expand 

``` go 
func Expand(s string, mapping func(string) string) string
```

Expand replaces ${var} or $var in the string based on the mapping function. For example, os.ExpandEnv(s) is equivalent to os.Expand(s, os.Getenv).

​	Expand函数根据映射函数替换字符串中的`${var}`或`$var`。例如，os.ExpandEnv(s)相当于os.Expand(s，os.Getenv)。

``` go 
package main

import (
	"fmt"
	"os"
)

func main() {
	mapper := func(placeholderName string) string {
		switch placeholderName {
		case "DAY_PART":
			return "morning"
		case "NAME":
			return "Gopher"
		}

		return ""
	}

	fmt.Println(os.Expand("Good ${DAY_PART}, $NAME!", mapper))

}
Output:

Good morning, Gopher!

```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Expand("Good ${DAY}, $NAME!", func(placeholderName string) string {
		switch placeholderName {
		case "DAY":
			return "morning"
		case "NAME":
			return "Gopher"
		}
		return ""
	}))
}
// Output:
//Good morning, Gopher!
```



### func ExpandEnv 

``` go 
func ExpandEnv(s string) string
```

ExpandEnv replaces ${var} or $var in the string according to the values of the current environment variables. References to undefined variables are replaced by the empty string.

​	ExpandEnv函数根据当前环境变量的值替换字符串中的`${var}`或`$var`。对未定义的变量的引用将替换为空字符串。

``` go 
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

}

Output:

gopher lives in /usr/gopher.
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")
	os.Setenv("LANGUAGE", "go")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}.`${language}` is empty!"))
}

// Output:
//gopher lives in /usr/gopher.`` is empty!
```



### func Getegid 

``` go 
func Getegid() int
```

Getegid returns the numeric effective group id of the caller.

​	Getegid函数返回调用方的有效组ID。

> 个人注释
>
> ​	`Getegid`中的第二个`e`是`effective`的意思。
>
> ​	参见[func Getgid ](#func-getgid )。

On Windows, it returns -1.

​	在Windows上，它返回-1。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 当前用户是 lx
	fmt.Println("当前进程（或调用者）的有效组ID：", os.Getegid())
	fmt.Println("当前进程（或调用者）的组ID：", os.Getgid())
	//fmt.Println("over")
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ go build Getegid.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ ls -l Getegid
//-rwxr-xr-x 1 lx lx 1802350 Aug 29 08:37 Getegid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ sudo chmod g+s Getegid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ ./Getegid
//当前进程（或调用者）的有效组ID： 1000
//当前进程（或调用者）的组ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ sudo ./Getegid
//当前进程（或调用者）的有效组ID： 1000
//当前进程（或调用者）的组ID： 0
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ sudo su - root
//root@DESKTOP-2OAUARV:~# cd /home/lx/goprojects/go_std_examples/os/os_self/f_Getegid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ./Getegid
//当前进程（或调用者）的有效组ID： 1000
//当前进程（或调用者）的组ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ls -l Getegid
//-rwxr-sr-x 1 lx lx 1802350 Aug 29 08:37 Getegid

// ---------------------------- 这里使用了 root 用户重新编辑了 Getegid.go --------------------------
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# go build Getegid.go
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ls -l Getegid
//-rwxr-xr-x 1 root root 1802462 Aug 29 08:39 Getegid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ./Getegid
//当前进程（或调用者）的有效组ID： 0
//当前进程（或调用者）的组ID： 0
//over
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# chmod g+s Getegid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ls -l Getegid
//-rwxr-sr-x 1 root root 1802462 Aug 29 08:39 Getegid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ./Getegid
//当前进程（或调用者）的有效组ID： 0
//当前进程（或调用者）的组ID： 0
//over
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# su - lx
//lx@DESKTOP-2OAUARV:~$ cd goprojects/go_std_examples/os/os_self/f_Getegid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ ./Getegid
//当前进程（或调用者）的有效组ID： 0
//当前进程（或调用者）的组ID： 1000
//over
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ sudo ./Getegid
//[sudo] password for lx:
//当前进程（或调用者）的有效组ID： 0
//当前进程（或调用者）的组ID： 0
//over
```



### func Getenv 

``` go 
func Getenv(key string) string
```

Getenv retrieves the value of the environment variable named by the key. It returns the value, which will be empty if the variable is not present. To distinguish between an empty value and an unset value, use LookupEnv.

​	Getenv函数检索由键指定的环境变量的值。它返回值，如果变量不存在，则为空。要区分空值和未设置的值，请使用 LookupEnv。

``` go 
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Printf("%s lives in %s.\n", os.Getenv("NAME"), os.Getenv("BURROW"))

}
Output:

gopher lives in /usr/gopher.
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Printf("%s lives in %s.`%s` is empty.\n", os.Getenv("NAME"), os.Getenv("BURROW"), os.Getenv("MISSING_KEY"))
}
// Output:
//gopher lives in /usr/gopher.`` is empty.
```



### func Geteuid 

``` go 
func Geteuid() int
```

Geteuid returns the numeric effective user id of the caller.

​	Geteuid函数返回调用方的有效用户 ID。

> 个人注释
>
> ​	`Geteuid`中的第二个`e`是`effective`的意思。

On Windows, it returns -1.

​	在Windows上，它返回-1。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 当前用户是 lx
	fmt.Println("当前进程（或调用者）的有效用户ID：", os.Geteuid())
	fmt.Println("当前进程（或调用者）的用户ID：", os.Getuid())
	//fmt.Println("over")
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ go build Geteuid.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ ls -l Geteuid
//-rwxr-xr-x 1 lx lx 1802342 Aug 29 08:52 Geteuid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo chmod u+s Geteuid
//[sudo] password for lx:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo chmod u+s Geteuid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ ls -l Geteuid
//-rwsr-xr-x 1 lx lx 1802342 Aug 29 08:52 Geteuid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ ./Geteuid
//当前进程（或调用者）的有效用户ID： 1000
//当前进程（或调用者）的用户ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo ./Geteuid
//当前进程（或调用者）的有效用户ID： 1000
//当前进程（或调用者）的用户ID： 0
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo su - root
//root@DESKTOP-2OAUARV:~# cd /home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# ./Geteuid
//当前进程（或调用者）的有效用户ID： 1000
//当前进程（或调用者）的用户ID： 0

// ---------------------------- 这里使用了 root 用户重新编辑了 Geteuid.go --------------------------
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# go build Geteuid.go
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# ls -l Geteuid
//-rwxr-xr-x 1 root root 1802462 Aug 29 08:55 Geteuid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# chmod u+s Geteuid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# ls -l Geteuid
//-rwsr-xr-x 1 root root 1802462 Aug 29 08:55 Geteuid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# ./Geteuid
//当前进程（或调用者）的有效用户ID： 0
//当前进程（或调用者）的用户ID： 0
//over
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# su - lx
//lx@DESKTOP-2OAUARV:~$ cd goprojects/go_std_examples/os/os_self/f_Geteuid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ ./Geteuid
//当前进程（或调用者）的有效用户ID： 0
//当前进程（或调用者）的用户ID： 1000
//over
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo ./Geteuid
//[sudo] password for lx:
//当前进程（或调用者）的有效用户ID： 0
//当前进程（或调用者）的用户ID： 0
//over
```



### func Getgid 

``` go 
func Getgid() int
```

Getgid returns the numeric group id of the caller.

​	Getgid函数返回调用方的组 ID。

> 个人注释
>
> ​	问了下ChatGPT，说是Getegid函数获取的是有效组ID，而Getgid函数获取的是实际组ID。还是不明白！待后续处理！
>
> ​	看了下Linux 的用户和访问权限，其中提到了SetUID 和 SetGID，以及粘滞位。我猜想，应该是和这些有关的！
>
> > 在 Linux 中，SetUID（设置用户 ID）和 SetGID（设置组 ID）是特殊的权限位，用于控制可执行文件的执行权限。   	1. SetUID 权限（Set User ID）：    
> >
> > - SetUID 权限允许一个可执行文件在执行时临时获得该文件所有者的权限。    - 这意味着，无论谁执行该文件，都将以文件所有者的身份执行，而不是以执行者自己的身份。    
> > - SetUID 权限通常用于需要特定权限或特定身份执行的程序。    
> > - 例如，passwd 程序用于更改用户密码。它需要访问 /etc/shadow 文件，该文件的权限通常只允许 root 用户访问。但是，passwd 程序具有 SetUID 权限，因此当普通用户执行 passwd 时，它可以以 root 用户的身份访问 /etc/shadow 文件并修改密码。   
> >   2. SetGID 权限（Set Group ID）：    
> > - SetGID 权限允许一个可执行文件在执行时临时获得该文件所属组的权限。    
> > - 这意味着，无论谁执行该文件，都将以文件所属组的身份执行，而不是以执行者自己的身份。    
> > - SetGID 权限通常用于需要特定组权限或特定组身份执行的程序。    
> > - 例如，一个共享文件夹的可执行文件具有 SetGID 权限，使得所有执行者都以共享文件夹所属组的身份执行。这样，无论是哪个用户执行该文件，都可以获得共享文件夹所属组的权限，从而访问共享文件夹。
>
> ​	好吧，开始验证下！

On Windows, it returns -1.

​	在Windows上，它返回-1。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 当前用户是 lx
	fmt.Println("当前进程（或调用者）的有效组ID：", os.Getegid())
	fmt.Println("当前进程（或调用者）的组ID：", os.Getgid())
	//fmt.Println("over")
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ go build Getegid.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ ls -l Getegid
//-rwxr-xr-x 1 lx lx 1802350 Aug 29 08:37 Getegid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ sudo chmod g+s Getegid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ ./Getegid
//当前进程（或调用者）的有效组ID： 1000
//当前进程（或调用者）的组ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ sudo ./Getegid
//当前进程（或调用者）的有效组ID： 1000
//当前进程（或调用者）的组ID： 0
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ sudo su - root
//root@DESKTOP-2OAUARV:~# cd /home/lx/goprojects/go_std_examples/os/os_self/f_Getegid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ./Getegid
//当前进程（或调用者）的有效组ID： 1000
//当前进程（或调用者）的组ID： 0
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ls -l Getegid
//-rwxr-sr-x 1 lx lx 1802350 Aug 29 08:37 Getegid

// ---------------------------- 这里使用了 root 用户重新编辑了 Getegid.go --------------------------
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# go build Getegid.go
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ls -l Getegid
//-rwxr-xr-x 1 root root 1802462 Aug 29 08:39 Getegid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ./Getegid
//当前进程（或调用者）的有效组ID： 0
//当前进程（或调用者）的组ID： 0
//over
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# chmod g+s Getegid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ls -l Getegid
//-rwxr-sr-x 1 root root 1802462 Aug 29 08:39 Getegid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# ./Getegid
//当前进程（或调用者）的有效组ID： 0
//当前进程（或调用者）的组ID： 0
//over
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getegid# su - lx
//lx@DESKTOP-2OAUARV:~$ cd goprojects/go_std_examples/os/os_self/f_Getegid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ ./Getegid
//当前进程（或调用者）的有效组ID： 0
//当前进程（或调用者）的组ID： 1000
//over
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getegid$ sudo ./Getegid
//[sudo] password for lx:
//当前进程（或调用者）的有效组ID： 0
//当前进程（或调用者）的组ID： 0
//over
```



### func Getgroups 

``` go 
func Getgroups() ([]int, error)
```

Getgroups returns a list of the numeric ids of groups that the caller belongs to.

​	Getgroups函数返回调用方所属的组的 ID 列表。

On Windows, it returns syscall.EWINDOWS. See the os/user package for a possible alternative.

​	在Windows上，它返回syscall.EWINDOWS。有关可能的替代方案，请参见os/user包。

Linux:

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 当前用户是 lx
	gids, err := os.Getgroups()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("当前进程（或调用者）的组ID列表：%#v\n", gids)
}
// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ go run Getgroups.go
//当前进程（或调用者）的组ID列表：[]int{4, 20, 24, 25, 27, 29, 30, 44, 46, 116, 1000, 1001}
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ whoami
//lx
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ go build Getgroups.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ ./Getgroups
//当前进程（或调用者）的组ID列表：[]int{4, 20, 24, 25, 27, 29, 30, 44, 46, 116, 1000, 1001}
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ sudo ./Getgroups
//[sudo] password for lx:
//当前进程（或调用者）的组ID列表：[]int{0}
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ sudo su - root
//root@DESKTOP-2OAUARV:~# cd /home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups# ./Getgroups
//当前进程（或调用者）的组ID列表：[]int{0}
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups# go run Getgroups.go
//当前进程（或调用者）的组ID列表：[]int{0}
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups# go build Getgroups.go
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups# ./Getgroups
//当前进程（或调用者）的组ID列表：[]int{0}
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Getgroups# sudo su - lx
//lx@DESKTOP-2OAUARV:~$ cd goprojects/go_std_examples/os/os_self/f_Getgroups
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ ./Getgroups
//当前进程（或调用者）的组ID列表：[]int{4, 20, 24, 25, 27, 29, 30, 44, 46, 116, 1000, 1001}
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getgroups$ sudo ./Getgroups
//[sudo] password for lx:
//当前进程（或调用者）的组ID列表：[]int{0}
```

Windows 10:

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {	
	gids, err := os.Getgroups()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("当前进程（或调用者）的组ID列表：%#v\n", gids)
}

// Output:
//PS F:\Devs\MyCodes\go_std_examples\os\os_self\f_Getgroups> go run .\Getgroups.go
//2023/08/28 17:01:21 getgroups: not supported by windows
//exit status 1
```



### func Getpagesize 

``` go 
func Getpagesize() int
```

Getpagesize returns the underlying system's memory page size.

​	Getpagesize函数返回底层系统的内存页面大小。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpagesize()) // 4096
}
// Output:
// 4096
```



### func Getpid 

``` go 
func Getpid() int
```

Getpid returns the process id of the caller.

​	Getpid函数返回调用方的进程 ID。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getpid())
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getpid$ go run Getpid.go
//1202651
```



### func Getppid 

``` go 
func Getppid() int
```

Getppid returns the process id of the caller's parent.

​	Getppid函数返回调用方的父进程 ID。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getppid())
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getppid$ go run Getppid.go
//1203113
```



### func Getuid 

``` go 
func Getuid() int
```

Getuid returns the numeric user id of the caller.

​	Getuid函数返回调用方的用户 ID。

On Windows, it returns -1.

​	在Windows上，它返回-1。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 当前用户是 lx
	fmt.Println("当前进程（或调用者）的有效用户ID：", os.Geteuid())
	fmt.Println("当前进程（或调用者）的用户ID：", os.Getuid())
	//fmt.Println("over")
}

// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ go build Geteuid.go
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ ls -l Geteuid
//-rwxr-xr-x 1 lx lx 1802342 Aug 29 08:52 Geteuid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo chmod u+s Geteuid
//[sudo] password for lx:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo chmod u+s Geteuid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ ls -l Geteuid
//-rwsr-xr-x 1 lx lx 1802342 Aug 29 08:52 Geteuid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ ./Geteuid
//当前进程（或调用者）的有效用户ID： 1000
//当前进程（或调用者）的用户ID： 1000
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo ./Geteuid
//当前进程（或调用者）的有效用户ID： 1000
//当前进程（或调用者）的用户ID： 0
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo su - root
//root@DESKTOP-2OAUARV:~# cd /home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# ./Geteuid
//当前进程（或调用者）的有效用户ID： 1000
//当前进程（或调用者）的用户ID： 0

// ---------------------------- 这里使用了 root 用户重新编辑了 Geteuid.go --------------------------
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# go build Geteuid.go
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# ls -l Geteuid
//-rwxr-xr-x 1 root root 1802462 Aug 29 08:55 Geteuid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# chmod u+s Geteuid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# ls -l Geteuid
//-rwsr-xr-x 1 root root 1802462 Aug 29 08:55 Geteuid
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# ./Geteuid
//当前进程（或调用者）的有效用户ID： 0
//当前进程（或调用者）的用户ID： 0
//over
//root@DESKTOP-2OAUARV:/home/lx/goprojects/go_std_examples/os/os_self/f_Geteuid# su - lx
//lx@DESKTOP-2OAUARV:~$ cd goprojects/go_std_examples/os/os_self/f_Geteuid
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ ./Geteuid
//当前进程（或调用者）的有效用户ID： 0
//当前进程（或调用者）的用户ID： 1000
//over
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Geteuid$ sudo ./Geteuid
//[sudo] password for lx:
//当前进程（或调用者）的有效用户ID： 0
//当前进程（或调用者）的用户ID： 0
//over
```



### func Getwd 

``` go 
func Getwd() (dir string, err error)
```

Getwd returns a rooted path name corresponding to the current directory. If the current directory can be reached via multiple paths (due to symbolic links), Getwd may return any one of them.

​	Getwd函数返回对应于当前目录的根路径名。如果可以通过多个路径到达当前目录(由于符号链接)，Getwd可能返回其中任何一个。

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	i := 0
	for i < 1000 {
		i++
		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		if dir != "/home/lx/goprojects/go_std_examples/os/os_self/f_Getwd" {
			fmt.Println(dir)
		}
	}
}

//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Getwd$ cd ../for_test/
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd1
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd2
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd3
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd4
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd5
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ ln -s /home/lx/goprojects/go_std_examples/os/os_self/f_Getwd dir_f_Getwd6
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/for_test$ cd ../f_Getwd/

// 也没有像 “如果可以通过多个路径到达当前目录(由于符号链接)，Getwd可能返回其中任何一个。” 这句话所属的。
```



### func Hostname 

``` go 
func Hostname() (name string, err error)
```

Hostname returns the host name reported by the kernel.

​	Hostname函数返回内核报告的主机名。

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("主机名是：", name)
}
// Output:
//lx@DESKTOP-2OAUARV:~/goprojects/go_std_examples/os/os_self/f_Hostname$ go run Hostname.go
//主机名是： DESKTOP-2OAUARV
```



### func IsExist 

``` go 
func IsExist(err error) bool
```

IsExist returns a boolean indicating whether the error is known to report that a file or directory already exists. It is satisfied by ErrExist as well as some syscall errors.

​	IsExist函数返回一个布尔值，指示错误是否已知报告文件或目录已经存在。它满足ErrExist以及一些syscall错误。

This function predates errors.Is. It only supports errors returned by the os package. New code should use errors.Is(err, fs.ErrExist).

​	此函数先于errors.Is。它仅支持由os包返回的错误。新代码应使用errors.Is(err，fs.ErrExist)。

### func IsNotExist 

``` go 
func IsNotExist(err error) bool
```

IsNotExist returns a boolean indicating whether the error is known to report that a file or directory does not exist. It is satisfied by ErrNotExist as well as some syscall errors.

​	IsNotExist函数返回一个布尔值，指示错误是否已知报告文件或目录不存在。它满足ErrNotExist以及一些syscall错误。

This function predates errors.Is. It only supports errors returned by the os package. New code should use errors.Is(err, fs.ErrNotExist).

​	此函数先于errors.Is。它仅支持由os包返回的错误。新代码应使用errors.Is(err，fs.ErrNotExist)。

### func IsPathSeparator 

``` go 
func IsPathSeparator(c uint8) bool
```

IsPathSeparator reports whether c is a directory separator character.

​	IsPathSeparator函数返回一个布尔值，指示c是否为目录分隔符字符。

### func IsPermission 

``` go 
func IsPermission(err error) bool
```

IsPermission returns a boolean indicating whether the error is known to report that permission is denied. It is satisfied by ErrPermission as well as some syscall errors.

​	IsPermission函数返回一个布尔值，指示错误是否已知报告权限被拒绝。它与某些系统调用错误以及ErrPermission相符。

This function predates errors.Is. It only supports errors returned by the os package. New code should use errors.Is(err, fs.ErrPermission).

​	此函数早于errors.Is。它仅支持由os包返回的错误。新代码应使用errors.Is(err，fs.ErrPermission)。

### func IsTimeout  <- go1.10

``` go 
func IsTimeout(err error) bool
```

IsTimeout returns a boolean indicating whether the error is known to report that a timeout occurred.

​	IsTimeout函数返回一个布尔值，指示错误是否已知报告超时发生。

This function predates errors.Is, and the notion of whether an error indicates a timeout can be ambiguous. For example, the Unix error EWOULDBLOCK sometimes indicates a timeout and sometimes does not. New code should use errors.Is with a value appropriate to the call returning the error, such as os.ErrDeadlineExceeded.

​	此函数早于errors.Is，并且错误是否表示超时的概念可能不明确。例如，Unix错误EWOULDBLOCK有时指示超时，有时不是。新代码应使用与返回错误的调用相适应的值，例如os.ErrDeadlineExceeded。

### func Lchown 

``` go 
func Lchown(name string, uid, gid int) error
```

Lchown changes the numeric uid and gid of the named file. If the file is a symbolic link, it changes the uid and gid of the link itself. If there is an error, it will be of type *PathError.

​	Lchown函数更改命名文件的数值UID和GID。如果文件是符号链接，则更改链接本身的UID和GID。如果有错误，它将是*PathError类型。

On Windows, it always returns the syscall.EWINDOWS error, wrapped in *PathError.

​	在Windows上，它总是返回syscall.EWINDOWS错误，包装在*PathError中。

### func Link 

``` go 
func Link(oldname, newname string) error
```

Link creates newname as a hard link to the oldname file. If there is an error, it will be of type *LinkError.

​	Link函数将newname创建为oldname文件的硬链接。如果有错误，它将是`*LinkError`类型。

### func LookupEnv  <- go1.5

``` go 
func LookupEnv(key string) (string, bool)
```

LookupEnv retrieves the value of the environment variable named by the key. If the variable is present in the environment the value (which may be empty) is returned and the boolean is true. Otherwise the returned value will be empty and the boolean will be false.

​	LookupEnv函数检索由key指定的环境变量的值。如果变量存在于环境中，则返回值(可能为空)为true。否则，返回的值将为空，布尔值将为false。

``` go 
package main

import (
	"fmt"
	"os"
)

func main() {
	show := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("%s not set\n", key)
		} else {
			fmt.Printf("%s=%s\n", key, val)
		}
	}

	os.Setenv("SOME_KEY", "value")
	os.Setenv("EMPTY_KEY", "")

	show("SOME_KEY")
	show("EMPTY_KEY")
	show("MISSING_KEY")

}
Output:

SOME_KEY=value
EMPTY_KEY=
MISSING_KEY not set
```



### func Mkdir 

``` go 
func Mkdir(name string, perm FileMode) error
```

Mkdir creates a new directory with the specified name and permission bits (before umask). If there is an error, it will be of type *PathError.

​	Mkdir函数使用指定的名称和权限位(umask之前)创建一个新目录。如果有错误，它将是`*PathError`类型。

``` go 
package main

import (
	"log"
	"os"
)

func main() {
	err := os.Mkdir("testdir", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("testdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}

```



### func MkdirAll 

``` go 
func MkdirAll(path string, perm FileMode) error
```

MkdirAll creates a directory named path, along with any necessary parents, and returns nil, or else returns an error. The permission bits perm (before umask) are used for all directories that MkdirAll creates. If path is already a directory, MkdirAll does nothing and returns nil.

​	MkdirAll函数创建一个名为path的目录，并创建所有必要的父级目录，返回nil，否则返回错误。所有由MkdirAll创建的目录都将使用perm(在umask之前)设置的权限位。如果path已经是一个目录，则MkdirAll不执行任何操作并返回nil。

``` go 
package main

import (
	"log"
	"os"
)

func main() {
	err := os.MkdirAll("test/subdir", 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
	err = os.WriteFile("test/subdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
	if err != nil {
		log.Fatal(err)
	}
}

```



### func MkdirTemp  <- go1.16

``` go 
func MkdirTemp(dir, pattern string) (string, error)
```

MkdirTemp creates a new temporary directory in the directory dir and returns the pathname of the new directory. The new directory's name is generated by adding a random string to the end of pattern. If pattern includes a "*", the random string replaces the last "*" instead. If dir is the empty string, MkdirTemp uses the default directory for temporary files, as returned by TempDir. Multiple programs or goroutines calling MkdirTemp simultaneously will not choose the same directory. It is the caller's responsibility to remove the directory when it is no longer needed.

​	MkdirTemp函数在目录dir中创建一个新的临时目录，并返回新目录的路径名。新目录的名称通过在模式的末尾添加一个随机字符串来生成。如果模式包含一个" `*`"，则随机字符串替换最后一个"`*`"。如果dir为空字符串，则MkdirTemp使用临时文件的默认目录，由TempDir返回。同时调用MkdirTemp的多个程序或goroutine不会选择相同的目录。当不再需要目录时，由调用者负责删除该目录。

#### MkdirTemp Example

``` go 
package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.MkdirTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir) // clean up

	file := filepath.Join(dir, "tmpfile")
	if err := os.WriteFile(file, []byte("content"), 0666); err != nil {
		log.Fatal(err)
	}
}

```

#### MkdirTemp Example(Suffix)

``` go 
package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	logsDir, err := os.MkdirTemp("", "*-logs")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(logsDir) // clean up

	// Logs can be cleaned out earlier if needed by searching
	// for all directories whose suffix ends in *-logs.
	globPattern := filepath.Join(os.TempDir(), "*-logs")
	matches, err := filepath.Glob(globPattern)
	if err != nil {
		log.Fatalf("Failed to match %q: %v", globPattern, err)
	}

	for _, match := range matches {
		if err := os.RemoveAll(match); err != nil {
			log.Printf("Failed to remove %q: %v", match, err)
		}
	}
}

```



### func NewSyscallError 

``` go 
func NewSyscallError(syscall string, err error) error
```

NewSyscallError returns, as an error, a new SyscallError with the given system call name and error details. As a convenience, if err is nil, NewSyscallError returns nil.

​	NewSyscallError函数返回一个新的SyscallError，其中包含给定的系统调用名称和错误详细信息作为错误。为方便起见，如果err为nil，则NewSyscallError返回nil。

### func Pipe 

``` go 
func Pipe() (r *File, w *File, err error)
```

Pipe returns a connected pair of Files; reads from r return bytes written to w. It returns the files and an error, if any.

​	Pipe函数返回一对连接的文件。从r读取的字节返回到w中。如果有错误，则返回文件和错误。

### func ReadFile  <- go1.16

``` go 
func ReadFile(name string) ([]byte, error)
```

ReadFile reads the named file and returns the contents. A successful call returns err == nil, not err == EOF. Because ReadFile reads the whole file, it does not treat an EOF from Read as an error to be reported.

​	ReadFile函数读取指定的文件并返回文件内容。成功调用返回err == nil，而不是err == EOF。由于ReadFile读取整个文件，因此它不将从Read返回的EOF视为要报告的错误。

#### ReadFile Example

``` go 
package main

import (
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("testdata/hello")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write(data)

}

```



### func Readlink 

``` go 
func Readlink(name string) (string, error)
```

Readlink returns the destination of the named symbolic link. If there is an error, it will be of type *PathError.

​	Readlink函数返回指定符号链接的目标。如果有错误，则错误类型为* PathError。

### func Remove 

``` go 
func Remove(name string) error
```

Remove removes the named file or (empty) directory. If there is an error, it will be of type *PathError.

​	Remove函数删除指定的文件或(空)目录。如果有错误，则错误类型为`* PathError`。

### func RemoveAll 

``` go 
func RemoveAll(path string) error
```

RemoveAll removes path and any children it contains. It removes everything it can but returns the first error it encounters. If the path does not exist, RemoveAll returns nil (no error). If there is an error, it will be of type *PathError.

​	RemoveAll函数删除路径及其包含的所有子项。它会尽可能删除所有内容，但返回遇到的第一个错误。如果路径不存在，则RemoveAll返回nil(无错误)。如果有错误，则错误类型为`* PathError`。

### func Rename 

``` go 
func Rename(oldpath, newpath string) error
```

Rename renames (moves) oldpath to newpath. If newpath already exists and is not a directory, Rename replaces it. OS-specific restrictions may apply when oldpath and newpath are in different directories. Even within the same directory, on non-Unix platforms Rename is not an atomic operation. If there is an error, it will be of type *LinkError.

​	Rename函数将oldpath重命名(移动)为newpath。如果newpath已经存在且不是目录，则会替换它。当oldpath和newpath在不同目录中时，可能会受到操作系统特定的限制。即使在同一个目录中，对于非Unix平台，Rename操作也不是原子性的。如果出错，错误类型为`*LinkError`。

### func SameFile 

``` go 
func SameFile(fi1, fi2 FileInfo) bool
```

SameFile reports whether fi1 and fi2 describe the same file. For example, on Unix this means that the device and inode fields of the two underlying structures are identical; on other systems the decision may be based on the path names. SameFile only applies to results returned by this package's Stat. It returns false in other cases.

​	SameFile函数报告fi1和fi2描述的是否是同一文件。例如，在Unix上，这意味着两个底层结构的设备和inode字段是相同的；在其他系统上，决策可能基于路径名。SameFile仅适用于此包的Stat返回的结果。在其他情况下返回false。

### func Setenv 

``` go 
func Setenv(key, value string) error
```

Setenv sets the value of the environment variable named by the key. It returns an error, if any.

​	Setenv函数设置由键名key命名的环境变量的值。如果有任何错误，它将返回该错误。

### func Symlink 

``` go 
func Symlink(oldname, newname string) error
```

Symlink creates newname as a symbolic link to oldname. On Windows, a symlink to a non-existent oldname creates a file symlink; if oldname is later created as a directory the symlink will not work. If there is an error, it will be of type *LinkError.

​	Symlink函数将newname创建为指向oldname的符号链接。在Windows上，指向不存在的oldname的符号链接会创建一个文件符号链接；如果后来将oldname创建为目录，则符号链接将无法正常工作。如果出错，错误类型为`*LinkError`。

### func TempDir 

``` go 
func TempDir() string
```

TempDir returns the default directory to use for temporary files.

​	TempDir函数返回用于临时文件的默认目录。

On Unix systems, it returns $TMPDIR if non-empty, else /tmp. On Windows, it uses GetTempPath, returning the first non-empty value from %TMP%, %TEMP%, %USERPROFILE%, or the Windows directory. On Plan 9, it returns /tmp.

​	在Unix系统上，如果$TMPDIR不为空，则返回它，否则返回/tmp。在Windows上，它使用GetTempPath，从%TMP%，%TEMP%，%USERPROFILE%或Windows目录中返回第一个非空值。在Plan 9上，它返回/tmp。

The directory is neither guaranteed to exist nor have accessible permissions.

​	该目录既不保证存在，也不保证可访问权限。

### func Truncate 

``` go 
func Truncate(name string, size int64) error
```

Truncate changes the size of the named file. If the file is a symbolic link, it changes the size of the link's target. If there is an error, it will be of type *PathError.

​	Truncate函数更改命名文件的大小。如果文件是符号链接，则更改链接目标的大小。如果出错，错误类型为`*PathError`。

### func Unsetenv  <- go1.4

``` go 
func Unsetenv(key string) error
```

Unsetenv unsets a single environment variable.

​	Unsetenv函数取消设置单个环境变量。

#### Unsetenv Example 

``` go 
package main

import (
	"os"
)

func main() {
	os.Setenv("TMPDIR", "/my/tmp")
	defer os.Unsetenv("TMPDIR")
}

```



### func UserCacheDir  <- go1.11

``` go 
func UserCacheDir() (string, error)
```

UserCacheDir returns the default root directory to use for user-specific cached data. Users should create their own application-specific subdirectory within this one and use that.

​	UserCacheDir函数返回用于用户特定缓存数据的默认根目录。用户应在其中创建自己的应用程序特定子目录并使用该目录。

On Unix systems, it returns `$XDG_CACHE_HOME` as specified by https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if non-empty, else `$HOME/.cache`. On Darwin, it returns `$HOME/Library/Caches`. On Windows, it returns `%LocalAppData%`. On Plan 9, it returns `$home/lib/cache`.

​	在Unix系统上，如果`$XDG_CACHE_HOME`不为空，则根据https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html指定的内容返回该值，否则返回`$HOME/.cache`。在Darwin上，它返回`$HOME/Library/Caches`。在Windows上，它返回`%LocalAppData%`。在Plan 9上，它返回`$home/lib/cache`。

If the location cannot be determined (for example, $HOME is not defined), then it will return an error.

​	如果无法确定位置(例如，$HOME未定义)，则它将返回一个错误。

### func UserConfigDir  <- go1.13

``` go 
func UserConfigDir() (string, error)
```

UserConfigDir returns the default root directory to use for user-specific configuration data. Users should create their own application-specific subdirectory within this one and use that.

​	UserConfigDir函数返回用于用户特定配置数据的默认根目录。用户应在其中创建自己的应用程序特定子目录并使用该子目录。

On Unix systems, it returns `$XDG_CONFIG_HOME` as specified by https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if non-empty, else `$HOME/.config`. On Darwin, it returns `$HOME/Library/Application` Support. On Windows, it returns `%AppData%`. On Plan 9, it returns `$home/lib`.

​	在 Unix 系统上，如果非空，则返回 `$XDG_CONFIG_HOME`，否则返回 `$HOME/.config`。在 Darwin 上，它返回 `$HOME/Library/Application Support`。在 Windows 上，它返回 `%AppData%`。在 Plan 9 上，它返回 `$home/lib`。

If the location cannot be determined (for example, $HOME is not defined), then it will return an error.

​	如果无法确定位置(例如未定义 `$HOME`)，则返回错误。

### func UserHomeDir  <- go1.12

``` go 
func UserHomeDir() (string, error)
```

UserHomeDir returns the current user's home directory.

​	UserHomeDir函数返回当前用户的主目录。

On Unix, including macOS, it returns the $HOME environment variable. On Windows, it returns %USERPROFILE%. On Plan 9, it returns the $home environment variable.

​	在 Unix 系统(包括 macOS)上，它返回 `$HOME` 环境变量。在 Windows 上，它返回`%USERPROFILE%`。在 Plan 9 上，它返回 `$home` 环境变量。

### func WriteFile  <- go1.16

``` go 
func WriteFile(name string, data []byte, perm FileMode) error
```

WriteFile writes data to the named file, creating it if necessary. If the file does not exist, WriteFile creates it with permissions perm (before umask); otherwise WriteFile truncates it before writing, without changing permissions. Since Writefile requires multiple system calls to complete, a failure mid-operation can leave the file in a partially written state.

​	`WriteFile`函数将数据写入命名文件，如有必要创建它。如果文件不存在，则 WriteFile 使用权限 perm(在 umask 前)创建它；否则 WriteFile 在写入前将其截断，而不更改权限。由于 Writefile 需要多个系统调用才能完成，因此操作中的失败可能会使文件处于部分写入状态。

#### WriteFile Example

``` go 
package main

import (
	"log"
	"os"
)

func main() {
	err := os.WriteFile("testdata/hello", []byte("Hello, Gophers!"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

```



## 类型

### type DirEntry  <- go1.16

``` go 
type DirEntry = fs.DirEntry
```

A DirEntry is an entry read from a directory (using the ReadDir function or a File's ReadDir method).

​	DirEntry 是从目录中读取的条目(使用 ReadDir 函数或 File 的 ReadDir 方法)。

#### func ReadDir  <- go1.16

``` go 
func ReadDir(name string) ([]DirEntry, error)
```

ReadDir reads the named directory, returning all its directory entries sorted by filename. If an error occurs reading the directory, ReadDir returns the entries it was able to read before the error, along with the error.

​	ReadDir函数读取命名目录，返回其所有目录条目，按文件名排序。如果读取目录时出现错误，ReadDir 返回在错误之前能够读取的条目以及错误。

##### ReadDir Example

``` go 
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

```



### type File 

``` go 
type File struct {
	// contains filtered or unexported fields
	// 包含已过滤或未导出的字段
}
```

File represents an open file descriptor.

​	File结构体表示一个打开的文件描述符。

#### func Create 

``` go 
func Create(name string) (*File, error)
```

Create creates or truncates the named file. If the file already exists, it is truncated. If the file does not exist, it is created with mode 0666 (before umask). If successful, methods on the returned File can be used for I/O; the associated file descriptor has mode O_RDWR. If there is an error, it will be of type *PathError.

​	Create函数创建或截断命名文件。如果文件已经存在，则将其截断。如果文件不存在，则创建它以使用 mode 0666(在 umask 前)。如果成功，则返回的 File 上的方法可用于 I/O；相关的文件描述符具有 mode O_RDWR。如果有错误，它将是 `*PathError` 类型。

#### func CreateTemp  <- go1.16

``` go 
func CreateTemp(dir, pattern string) (*File, error)
```

CreateTemp creates a new temporary file in the directory dir, opens the file for reading and writing, and returns the resulting file. The filename is generated by taking pattern and adding a random string to the end. If pattern includes a "`*`", the random string replaces the last "`*`". If dir is the empty string, CreateTemp uses the default directory for temporary files, as returned by TempDir. Multiple programs or goroutines calling CreateTemp simultaneously will not choose the same file. The caller can use the file's Name method to find the pathname of the file. It is the caller's responsibility to remove the file when it is no longer needed.

​	CreateTemp函数在目录 dir 中创建一个新的临时文件，并打开它用于读写操作，并返回该文件。文件名由 pattern 生成，并在末尾添加一个随机字符串。如果 pattern 包含一个 "`*`", 随机字符串将替换最后一个 "`*`"。如果 dir 是空字符串，则 CreateTemp 使用返回的 TempDir 作为临时文件的默认目录。同时调用 CreateTemp 的多个程序或 goroutine 不会选择同一个文件。调用者可以使用文件的 Name 方法查找文件的路径名。当不再需要文件时，由调用者负责删除它。

CreateTemp Example

``` go 
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up

	if _, err := f.Write([]byte("content")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

```

##### CreateTemp Example (Suffix)

``` go 
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up

	if _, err := f.Write([]byte("content")); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

```



#### func NewFile 

``` go 
func NewFile(fd uintptr, name string) *File
```

NewFile returns a new File with the given file descriptor and name. The returned value will be nil if fd is not a valid file descriptor. On Unix systems, if the file descriptor is in non-blocking mode, NewFile will attempt to return a pollable File (one for which the SetDeadline methods work).

​	NewFile函数返回具有给定文件描述符和名称的新文件。如果 fd 不是有效的文件描述符，则返回值将为 nil。在 Unix 系统上，如果文件描述符处于非阻塞模式，则 NewFile 将尝试返回可轮询的文件(其 SetDeadline 方法有效)。

After passing it to NewFile, fd may become invalid under the same conditions described in the comments of the Fd method, and the same constraints apply.

​	将其传递给 NewFile 后，fd 可能会在与 Fd 方法的注释相同的条件下无效，同样的限制也适用。

#### func Open 

``` go 
func Open(name string) (*File, error)
```

Open opens the named file for reading. If successful, methods on the returned file can be used for reading; the associated file descriptor has mode O_RDONLY. If there is an error, it will be of type *PathError.

​	Open函数打开具有指定名称的文件以供读取。如果成功，可以使用返回的文件上的方法进行读取。关联的文件描述符具有 O_RDONLY 模式。如果出现错误，类型为 `*PathError`。

#### func OpenFile 

``` go 
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

OpenFile is the generalized open call; most users will use Open or Create instead. It opens the named file with specified flag (O_RDONLY etc.). If the file does not exist, and the O_CREATE flag is passed, it is created with mode perm (before umask). If successful, methods on the returned File can be used for I/O. If there is an error, it will be of type *PathError.

​	OpenFile函数是通用的打开调用，大多数用户将使用 Open 或 Create 来代替。使用指定标志(O_RDONLY 等)打开指定名称的文件。如果该文件不存在并且传递了 O_CREATE 标志，则使用模式 perm(位于 umask 之前)创建它。如果成功，则可以使用返回的文件上的方法进行 I/O。如果出现错误，类型为 `*PathError`。

##### OpenFile Example

``` go 
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

```

##### OpenFile Example(Append) 

``` go 
package main

import (
	"log"
	"os"
)

func main() {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("appended some data\n")); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

```

#### func OpenInRoot <- 1.24.0

```
func OpenInRoot(dir, name string) (*File, error)
```

OpenInRoot opens the file name in the directory dir. It is equivalent to OpenRoot(dir) followed by opening the file in the root.

OpenInRoot returns an error if any component of the name references a location outside of dir.

See [Root](https://pkg.go.dev/os@go1.24.2#Root) for details and limitations.

#### (*File) Chdir 

``` go 
func (f *File) Chdir() error
```

Chdir changes the current working directory to the file, which must be a directory. If there is an error, it will be of type *PathError.

​	Chdir方法将当前工作目录更改为文件。它必须是目录。如果出现错误，类型为 `*PathError`。

#### (*File) Chmod 

``` go 
func (f *File) Chmod(mode FileMode) error
```

Chmod changes the mode of the file to mode. If there is an error, it will be of type *PathError.

​	Chmod方法将文件的模式更改为 mode。如果出现错误，类型为 `*PathError`。

#### (*File) Chown 

``` go 
func (f *File) Chown(uid, gid int) error
```

Chown changes the numeric uid and gid of the named file. If there is an error, it will be of type *PathError.

​	Chown方法更改具有指定名称的文件的数字 uid 和 gid。如果出现错误，类型为 `*PathError`。

On Windows, it always returns the syscall.EWINDOWS error, wrapped in *PathError.

​	在 Windows 上，它始终返回 syscall.EWINDOWS 错误，包装在 `*PathError` 中。

#### (*File) Close 

``` go 
func (f *File) Close() error
```

Close closes the File, rendering it unusable for I/O. On files that support SetDeadline, any pending I/O operations will be canceled and return immediately with an ErrClosed error. Close will return an error if it has already been called.

​	Close方法关闭文件，使其不能用于I/O。对于支持SetDeadline的文件，任何待处理的I/O操作都将被取消，并立即返回ErrClosed错误。如果已经调用过Close方法，则Close将返回一个错误。

#### (*File) Fd 

``` go 
func (f *File) Fd() uintptr
```

Fd returns the integer Unix file descriptor referencing the open file. If f is closed, the file descriptor becomes invalid. If f is garbage collected, a finalizer may close the file descriptor, making it invalid; see runtime.SetFinalizer for more information on when a finalizer might be run. On Unix systems this will cause the SetDeadline methods to stop working. Because file descriptors can be reused, the returned file descriptor may only be closed through the Close method of f, or by its finalizer during garbage collection. Otherwise, during garbage collection the finalizer may close an unrelated file descriptor with the same (reused) number.

​	`Fd`方法返回整数Unix文件描述符，引用打开的文件。如果f已关闭，则文件描述符将变为无效。如果f被垃圾回收，终结器可能会关闭文件描述符，使其无效；有关何时运行终结器的更多信息，请参见runtime.SetFinalizer。在Unix系统中，这将导致SetDeadline方法停止工作。因为文件描述符可以被重用，所以返回的文件描述符只能通过f的Close方法或在垃圾回收期间的终结器关闭。否则，在垃圾回收期间，终结器可能会关闭具有相同(重用)编号的不相关文件描述符。

As an alternative, see the f.SyscallConn method.

​	作为替代方案，请参见f.SyscallConn方法。

#### (*File) Name 

``` go 
func (f *File) Name() string
```

Name returns the name of the file as presented to Open.

​	Name方法返回打开文件的名称。

#### (*File) Read 

``` go 
func (f *File) Read(b []byte) (n int, err error)
```

Read reads up to len(b) bytes from the File and stores them in b. It returns the number of bytes read and any error encountered. At end of file, Read returns 0, io.EOF.

​	Read方法从文件中读取最多len(b)字节，并将其存储在b中。它返回读取的字节数和任何遇到的错误。在文件末尾，Read返回0，io.EOF。

#### (*File) ReadAt 

``` go 
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
```

ReadAt reads len(b) bytes from the File starting at byte offset off. It returns the number of bytes read and the error, if any. ReadAt always returns a non-nil error when n < len(b). At end of file, that error is io.EOF.

​	ReadAt方法从文件中读取len(b)字节，从偏移量off处开始。它返回读取的字节数和错误(如果有)。当n < len(b)时，ReadAt总是返回非nil错误。在文件末尾，该错误为io.EOF。

#### (*File) ReadDir  <- go1.16

``` go 
func (f *File) ReadDir(n int) ([]DirEntry, error)
```

ReadDir reads the contents of the directory associated with the file f and returns a slice of DirEntry values in directory order. Subsequent calls on the same file will yield later DirEntry records in the directory.

​	ReadDir方法读取与文件f相关联的目录的内容，并按文件名返回一系列DirEntry值。对同一文件的后续调用将按目录顺序返回后续的DirEntry记录。

If n > 0, ReadDir returns at most n DirEntry records. In this case, if ReadDir returns an empty slice, it will return an error explaining why. At the end of a directory, the error is io.EOF.

​	如果n > 0，则ReadDir最多返回n个DirEntry记录。在这种情况下，如果ReadDir返回一个空切片，则它将返回一个解释原因的错误。在目录末尾，错误为io.EOF。

If n <= 0, ReadDir returns all the DirEntry records remaining in the directory. When it succeeds, it returns a nil error (not io.EOF).

​	如果n <= 0，则ReadDir将返回剩余目录中的所有DirEntry记录。当它成功时，它返回一个nil错误(而不是io.EOF)。

#### (*File) ReadFrom  <- go1.15

``` go 
func (f *File) ReadFrom(r io.Reader) (n int64, err error)
```

ReadFrom implements io.ReaderFrom.

​	ReadFrom方法实现了io.ReaderFrom。

#### (*File) Readdir 

``` go 
func (f *File) Readdir(n int) ([]FileInfo, error)
```

Readdir reads the contents of the directory associated with file and returns a slice of up to n FileInfo values, as would be returned by Lstat, in directory order. Subsequent calls on the same file will yield further FileInfos.

​	Readdir方法读取与文件 f 关联的目录并以目录顺序返回最多 n 个 FileInfo 值的切片，与 Lstat 返回的一样。在同一文件上的后续调用将返回更多的 FileInfos。

If n > 0, Readdir returns at most n FileInfo structures. In this case, if Readdir returns an empty slice, it will return a non-nil error explaining why. At the end of a directory, the error is io.EOF.

​	如果 n > 0，则 Readdir 返回最多 n 个 FileInfo 结构。在这种情况下，如果 Readdir 返回一个空切片，则它将返回一个非 nil 的错误来解释原因。在目录末尾，错误为 io.EOF。

If n <= 0, Readdir returns all the FileInfo from the directory in a single slice. In this case, if Readdir succeeds (reads all the way to the end of the directory), it returns the slice and a nil error. If it encounters an error before the end of the directory, Readdir returns the FileInfo read until that point and a non-nil error.

​	如果 n <= 0，则 Readdir 以单个切片返回目录中的所有 FileInfo。在这种情况下，如果 Readdir 成功(一直读到目录结尾)，它将返回切片和 nil 错误。如果在目录结尾之前遇到错误，则 Readdir 返回读取到该点的 FileInfo 和非 nil 错误。

Most clients are better served by the more efficient ReadDir method.

​	大多数客户端最好使用更高效的 ReadDir 方法。

#### (*File) Readdirnames 

``` go 
func (f *File) Readdirnames(n int) (names []string, err error)
```

Readdirnames reads the contents of the directory associated with file and returns a slice of up to n names of files in the directory, in directory order. Subsequent calls on the same file will yield further names.

​	Readdirnames方法读取与文件 f 关联的目录并按目录顺序返回最多 n 个文件名的切片。在同一文件上的后续调用将返回更多的文件名。

If n > 0, Readdirnames returns at most n names. In this case, if Readdirnames returns an empty slice, it will return a non-nil error explaining why. At the end of a directory, the error is io.EOF.

​	如果 n > 0，则 Readdirnames 返回最多 n 个文件名。在这种情况下，如果 Readdirnames 返回一个空切片，则它将返回一个非 nil 的错误来解释原因。在目录末尾，错误为 io.EOF。

If n <= 0, Readdirnames returns all the names from the directory in a single slice. In this case, if Readdirnames succeeds (reads all the way to the end of the directory), it returns the slice and a nil error. If it encounters an error before the end of the directory, Readdirnames returns the names read until that point and a non-nil error.

​	如果 n <= 0，则 Readdirnames 以单个切片返回目录中的所有名称。在这种情况下，如果 Readdirnames 成功(一直读到目录结尾)，它将返回切片和 nil 错误。如果在目录结尾之前遇到错误，则 Readdirnames 返回读取到该点的名称和非 nil 错误。

#### (*File) Seek 

``` go 
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
```

Seek sets the offset for the next Read or Write on file to offset, interpreted according to whence: 0 means relative to the origin of the file, 1 means relative to the current offset, and 2 means relative to the end. It returns the new offset and an error, if any. The behavior of Seek on a file opened with O_APPEND is not specified.

​	Seek方法将下一个文件上的读取或写入的偏移量设置为 offset，根据 whence 进行解释：0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾。它返回新的偏移量和错误(如果有)。对于使用 O_APPEND 打开的文件，Seek 的行为未指定。

#### (*File) SetDeadline  <- go1.10

``` go 
func (f *File) SetDeadline(t time.Time) error
```

SetDeadline sets the read and write deadlines for a File. It is equivalent to calling both SetReadDeadline and SetWriteDeadline.

​	SetDeadline方法为文件设置读取和写入的截止日期。它等价于同时调用 SetReadDeadline 和 SetWriteDeadline。

Only some kinds of files support setting a deadline. Calls to SetDeadline for files that do not support deadlines will return ErrNoDeadline. On most systems ordinary files do not support deadlines, but pipes do.

​	只有某些类型的文件支持设置截止日期。调用不支持截止日期的文件的 SetDeadline方法将返回 ErrNoDeadline。在大多数系统上，普通文件不支持截止日期，但管道支持。

A deadline is an absolute time after which I/O operations fail with an error instead of blocking. The deadline applies to all future and pending I/O, not just the immediately following call to Read or Write. After a deadline has been exceeded, the connection can be refreshed by setting a deadline in the future.

​	一个截止时间是一个绝对时间，在此之后，I/O 操作会失败并返回一个错误，而不是阻塞。截止时间适用于所有未来和待处理的 I/O，而不仅仅是对 Read 或 Write 的立即调用。超过截止时间后，可以通过设置将来的截止时间来刷新连接。

If the deadline is exceeded a call to Read or Write or to other I/O methods will return an error that wraps ErrDeadlineExceeded. This can be tested using errors.Is(err, os.ErrDeadlineExceeded). That error implements the Timeout method, and calling the Timeout method will return true, but there are other possible errors for which the Timeout will return true even if the deadline has not been exceeded.

​	如果超过截止时间，对 Read 或 Write 或其他 I/O 方法的调用将返回一个包装 ErrDeadlineExceeded 的错误。可以使用 errors.Is(err，os.ErrDeadlineExceeded) 进行测试。该错误实现了 Timeout 方法，调用 Timeout 方法将返回 true，但有其他可能的错误，即使超时时间尚未超过，Timeout 也会返回 true。

An idle timeout can be implemented by repeatedly extending the deadline after successful Read or Write calls.

​	可以通过在成功的 Read 或 Write 调用后反复延长截止时间来实现空闲超时。

A zero value for t means I/O operations will not time out.

​	t 的零值表示 I/O 操作不会超时。

#### (*File) SetReadDeadline  <- go1.10

``` go 
func (f *File) SetReadDeadline(t time.Time) error
```

SetReadDeadline sets the deadline for future Read calls and any currently-blocked Read call. A zero value for t means Read will not time out. Not all files support setting deadlines; see SetDeadline.

​	SetReadDeadline方法设置将来的 Read 调用和任何当前阻塞的 Read 调用的截止时间。t 的零值表示 Read 不会超时。不是所有文件都支持设置截止时间；请参见 SetDeadline。

#### (*File) SetWriteDeadline  <- go1.10

``` go 
func (f *File) SetWriteDeadline(t time.Time) error
```

SetWriteDeadline sets the deadline for any future Write calls and any currently-blocked Write call. Even if Write times out, it may return n > 0, indicating that some of the data was successfully written. A zero value for t means Write will not time out. Not all files support setting deadlines; see SetDeadline.

​	SetWriteDeadline方法设置任何将来的 Write 调用和任何当前阻塞的 Write 调用的截止时间。即使 Write 超时，它也可能返回 n > 0，表示某些数据已成功写入。t 的零值表示 Write 不会超时。不是所有文件都支持设置截止时间；请参见 SetDeadline。

#### (*File) Stat 

``` go 
func (f *File) Stat() (FileInfo, error)
```

Stat returns the FileInfo structure describing file. If there is an error, it will be of type *PathError.

​	Stat方法返回描述文件的 FileInfo 结构。如果有错误，它将是 `*PathError` 类型。

#### (*File) Sync 

``` go 
func (f *File) Sync() error
```

Sync commits the current contents of the file to stable storage. Typically, this means flushing the file system's in-memory copy of recently written data to disk.

​	Sync方法将文件的当前内容提交到稳定存储。通常，这意味着将文件系统的最近写入数据的内存副本刷新到磁盘。

#### (*File) SyscallConn  <- go1.12

``` go 
func (f *File) SyscallConn() (syscall.RawConn, error)
```

SyscallConn returns a raw file. This implements the syscall.Conn interface.

​	SyscallConn方法返回一个原始文件。它实现了 syscall.Conn 接口。

#### (*File) Truncate 

``` go 
func (f *File) Truncate(size int64) error
```

Truncate changes the size of the file. It does not change the I/O offset. If there is an error, it will be of type *PathError.

​	Truncate方法改变文件的大小。它不会改变 I/O 偏移量。如果有错误，它将是 `*PathError` 类型。

#### (*File) Write 

``` go 
func (f *File) Write(b []byte) (n int, err error)
```

Write writes len(b) bytes from b to the File. It returns the number of bytes written and an error, if any. Write returns a non-nil error when n != len(b).

​	Write方法将 len(b) 个字节从 b 写入 File 中。它返回写入的字节数和错误(如果有)。当 n != len(b) 时，Write 返回非 nil 错误。

#### (*File) WriteAt 

``` go 
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
```

WriteAt writes len(b) bytes to the File starting at byte offset off. It returns the number of bytes written and an error, if any. WriteAt returns a non-nil error when n != len(b).

​	WriteAt方法从字节偏移量 off 处开始，将 len(b) 个字节写入 File。它返回写入的字节数和错误(如果有)。当 n != len(b) 时，WriteAt 返回非 nil 错误。

If file was opened with the O_APPEND flag, WriteAt returns an error.

​	如果打开 file 时使用了 O_APPEND 标志，则 WriteAt 返回一个错误。

#### (*File) WriteString 

``` go 
func (f *File) WriteString(s string) (n int, err error)
```

WriteString is like Write, but writes the contents of string s rather than a slice of bytes.

​	WriteString方法类似于 Write，但它写入字符串 s 的内容而不是字节切片。

####  (*File) WriteTo <- go1.22.0

``` go
func (f *File) WriteTo(w io.Writer) (n int64, err error)
```

WriteTo implements io.WriterTo.

​	WriteTo 实现了 io.WriterTo。

### type FileInfo 

``` go 
type FileInfo = fs.FileInfo
```

A FileInfo describes a file and is returned by Stat and Lstat.

​	FileInfo 描述一个文件，并由 Stat 和 Lstat 返回。

#### func Lstat 

``` go 
func Lstat(name string) (FileInfo, error)
```

Lstat returns a FileInfo describing the named file. If the file is a symbolic link, the returned FileInfo describes the symbolic link. Lstat makes no attempt to follow the link. If there is an error, it will be of type *PathError.

​	Lstat函数返回描述命名文件的 FileInfo。如果文件是符号链接，则返回的 FileInfo 描述符号链接。Lstat 不会尝试跟随链接。如果有错误，它将是 `*PathError` 类型的。

#### func Stat 

``` go 
func Stat(name string) (FileInfo, error)
```

Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError.

​	Stat函数返回描述命名文件的 FileInfo。如果有错误，它将是 `*PathError` 类型的。

### type FileMode 

``` go 
type FileMode = fs.FileMode
```

A FileMode represents a file's mode and permission bits. The bits have the same definition on all systems, so that information about files can be moved from one system to another portably. Not all bits apply to all systems. The only required bit is ModeDir for directories.

​	FileMode表示文件的模式和权限位。这些位在所有系统上具有相同的定义，因此可以将关于文件的信息可移植地从一个系统移动到另一个系统。并非所有位都适用于所有系统。ModeDir 用于目录是唯一需要的位。

#### FileMode Example

``` go 
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	fi, err := os.Lstat("some-filename")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("permissions: %#o\n", fi.Mode().Perm()) // 0400, 0777, etc.
	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&fs.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&fs.ModeNamedPipe != 0:
		fmt.Println("named pipe")
	}
}

```



### type LinkError 

``` go 
type LinkError struct {
	Op  string
	Old string
	New string
	Err error
}
```

LinkError records an error during a link or symlink or rename system call and the paths that caused it.

​	LinkError结构体记录了link、symlink或rename系统调用期间发生的错误和引起错误的路径。

#### (*LinkError) Error 

``` go 
func (e *LinkError) Error() string
```

#### (*LinkError) Unwrap  <- go1.13

``` go 
func (e *LinkError) Unwrap() error
```

### type PathError 

``` go 
type PathError = fs.PathError
```

PathError records an error and the operation and file path that caused it.



​	PathError类型记录了错误和引起错误的操作和文件路径。

### type ProcAttr 

``` go 
type ProcAttr struct {
    // If Dir is non-empty, the child changes into the directory before
	// creating the process.
	// 如果Dir不为空，则子进程在创建进程之前进入该目录。
	Dir string
    // If Env is non-nil, it gives the environment variables for the
	// new process in the form returned by Environ.
	// If it is nil, the result of Environ will be used.
	// 如果Env不为nil，则以Environ返回的形式为新进程提供环境变量。
	// 如果为nil，则使用Environ的结果。
	Env []string
    // Files specifies the open files inherited by the new process. The
	// first three entries correspond to standard input, standard output, and
	// standard error. An implementation may support additional entries,
	// depending on the underlying operating system. A nil entry corresponds
	// to that file being closed when the process starts.
	// On Unix systems, StartProcess will change these File values
	// to blocking mode, which means that SetDeadline will stop working
	// and calling Close will not interrupt a Read or Write.
	// Files指定新进程继承的打开文件。
    // 前三个条目对应标准输入、标准输出和标准错误输出。
	// 根据底层操作系统的不同，实现可能会支持更多的条目。
    // 空条目表示该文件在进程启动时关闭。
	// 在Unix系统上，StartProcess将把这些File值更改为阻塞模式，
    // 这意味着SetDeadline将停止工作，并且调用Close不会中断读取或写入。
	Files []*File

    // Operating system-specific process creation attributes.
	// Note that setting this field means that your program
	// may not execute properly or even compile on some
	// operating systems.
	// 操作系统特定的进程创建属性。
	// 请注意，设置此字段意味着您的程序可能
    // 无法在某些操作系统上正确执行甚至无法编译。
	Sys *syscall.SysProcAttr
}
```

ProcAttr holds the attributes that will be applied to a new process started by StartProcess.

​	ProcAttr结构体保存将应用于由StartProcess启动的新进程的属性。

### type Process 

``` go 
type Process struct {
	Pid int
	// contains filtered or unexported fields
	// 包含已过滤或未导出的字段
}
```

Process stores the information about a process created by StartProcess.

​	Process结构体存储由StartProcess创建的进程的信息。

#### func FindProcess 

``` go 
func FindProcess(pid int) (*Process, error)
```

FindProcess looks for a running process by its pid.

​	FindProcess函数按其pid查找正在运行的进程。

The Process it returns can be used to obtain information about the underlying operating system process.

​	它返回的Process可用于获取有关底层操作系统进程的信息。

On Unix systems, FindProcess always succeeds and returns a Process for the given pid, regardless of whether the process exists.

​	在Unix系统上，FindProcess始终成功并返回给定pid的Process，无论进程是否存在。

#### func StartProcess 

``` go 
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
```

StartProcess starts a new process with the program, arguments and attributes specified by name, argv and attr. The argv slice will become os.Args in the new process, so it normally starts with the program name.

​	StartProcess函数使用由name、argv和attr指定的程序、参数和属性启动一个新的进程。argv切片将成为新进程的os.Args，因此它通常以程序名称开头。

If the calling goroutine has locked the operating system thread with runtime.LockOSThread and modified any inheritable OS-level thread state (for example, Linux or Plan 9 name spaces), the new process will inherit the caller's thread state.

​	如果调用的goroutine已使用runtime.LockOSThread锁定操作系统线程并修改了任何可继承的操作系统级线程状态(例如Linux或Plan 9名称空间)，新进程将继承调用者的线程状态。

StartProcess is a low-level interface. The os/exec package provides higher-level interfaces.

​	StartProcess是一个低级接口。os/exec包提供了更高级的接口。

If there is an error, it will be of type *PathError.

​	如果有错误，则为`*PathError`类型。

#### (*Process) Kill 

``` go 
func (p *Process) Kill() error
```

Kill causes the Process to exit immediately. Kill does not wait until the Process has actually exited. This only kills the Process itself, not any other processes it may have started.

​	Kill方法使进程立即退出。Kill不会等待进程实际退出。这只会杀死进程本身，而不是它可能启动的任何其他进程。

#### (*Process) Release 

``` go 
func (p *Process) Release() error
```

Release releases any resources associated with the Process p, rendering it unusable in the future. Release only needs to be called if Wait is not.

​	Release方法释放与进程p关联的任何资源，使其在将来无法使用。如果不使用Wait，则只需要调用Release。

#### (*Process) Signal 

``` go 
func (p *Process) Signal(sig Signal) error
```

Signal sends a signal to the Process. Sending Interrupt on Windows is not implemented.

​	Signal方法向进程发送信号。在Windows上发送中断信号未实现。

#### (*Process) Wait 

``` go 
func (p *Process) Wait() (*ProcessState, error)
```

Wait waits for the Process to exit, and then returns a ProcessState describing its status and an error, if any. Wait releases any resources associated with the Process. On most operating systems, the Process must be a child of the current process or an error will be returned.

​	Wait方法等待进程退出，然后返回ProcessState描述其状态和任何错误。Wait释放与进程关联的任何资源。在大多数操作系统上，进程必须是当前进程的子进程，否则将返回一个错误。

### type ProcessState 

``` go 
type ProcessState struct {
	// contains filtered or unexported fields
}
```

ProcessState stores information about a process, as reported by Wait.

​	ProcessState结构体存储关于进程的信息，由Wait方法报告。

#### (*ProcessState) ExitCode  <- go1.12

``` go 
func (p *ProcessState) ExitCode() int
```

ExitCode returns the exit code of the exited process, or -1 if the process hasn't exited or was terminated by a signal.

​	ExitCode方法返回退出的进程的退出代码，如果进程尚未退出或被信号终止，则返回-1。

#### (*ProcessState) Exited 

``` go 
func (p *ProcessState) Exited() bool
```

Exited reports whether the program has exited. On Unix systems this reports true if the program exited due to calling exit, but false if the program terminated due to a signal.

​	Exited方法报告程序是否已退出。在Unix系统上，如果程序由于调用exit而退出，则此项报告为true，但如果程序由于信号终止而终止，则此项报告为false。

#### (*ProcessState) Pid 

``` go 
func (p *ProcessState) Pid() int
```

Pid returns the process id of the exited process.

​	Pid方法返回已退出进程的进程ID。

#### (*ProcessState) String 

``` go 
func (p *ProcessState) String() string
```

#### (*ProcessState) Success 

``` go 
func (p *ProcessState) Success() bool
```

Success reports whether the program exited successfully, such as with exit status 0 on Unix.

​	Success方法报告程序是否成功退出，例如在Unix上以退出状态0退出。

#### (*ProcessState) Sys 

``` go 
func (p *ProcessState) Sys() any
```

Sys returns system-dependent exit information about the process. Convert it to the appropriate underlying type, such as syscall.WaitStatus on Unix, to access its contents.

​	Sys方法返回有关进程的系统相关退出信息。将其转换为适当的底层类型，例如在Unix上的syscall.WaitStatus，以访问其内容。

#### (*ProcessState) SysUsage 

``` go 
func (p *ProcessState) SysUsage() any
```

SysUsage returns system-dependent resource usage information about the exited process. Convert it to the appropriate underlying type, such as *syscall.Rusage on Unix, to access its contents. (On Unix, *syscall.Rusage matches struct rusage as defined in the getrusage(2) manual page.)

​	SysUsage方法返回有关已退出进程的系统相关资源使用情况信息。将其转换为适当的底层类型，例如在Unix上的`*syscall.Rusage`，以访问其内容。(在Unix上，`*syscall.Rusage`与getrusage(2)手册页中定义的struct rusage匹配。)

#### (*ProcessState) SystemTime 

``` go 
func (p *ProcessState) SystemTime() time.Duration
```

SystemTime returns the system CPU time of the exited process and its children.

​	SystemTime方法返回已退出进程及其子进程的系统CPU时间。

#### (*ProcessState) UserTime 

``` go 
func (p *ProcessState) UserTime() time.Duration
```

UserTime returns the user CPU time of the exited process and its children.

​	UserTime方法返回已退出进程及其子进程的用户CPU时间。

### type Root <- 1.24.0

```go
type Root struct {
	// contains filtered or unexported fields
}
```

Root may be used to only access files within a single directory tree.

Methods on Root can only access files and directories beneath a root directory. If any component of a file name passed to a method of Root references a location outside the root, the method returns an error. File names may reference the directory itself (.).

Methods on Root will follow symbolic links, but symbolic links may not reference a location outside the root. Symbolic links must not be absolute.

Methods on Root do not prohibit traversal of filesystem boundaries, Linux bind mounts, /proc special files, or access to Unix device files.

Methods on Root are safe to be used from multiple goroutines simultaneously.

On most platforms, creating a Root opens a file descriptor or handle referencing the directory. If the directory is moved, methods on Root reference the original directory in its new location.

Root's behavior differs on some platforms:

- When GOOS=windows, file names may not reference Windows reserved device names such as NUL and COM1.
- When GOOS=js, Root is vulnerable to TOCTOU (time-of-check-time-of-use) attacks in symlink validation, and cannot ensure that operations will not escape the root.
- When GOOS=plan9 or GOOS=js, Root does not track directories across renames. On these platforms, a Root references a directory name, not a file descriptor.

#### func OpenRoot <- 1.24.0

```go
func OpenRoot(name string) (*Root, error)
```

OpenRoot opens the named directory. If there is an error, it will be of type *PathError.

#### (*Root) Close <- 1.24.0

```go
func (r *Root) Close() error
```

Close closes the Root. After Close is called, methods on Root return errors.

#### (*Root) Create <- 1.24.0

```go
func (r *Root) Create(name string) (*File, error)
```

Create creates or truncates the named file in the root. See [Create](https://pkg.go.dev/os@go1.24.2#Create) for more details.

#### (*Root) FS <- 1.24.0

```go
func (r *Root) FS() fs.FS
```

FS returns a file system (an fs.FS) for the tree of files in the root.

The result implements [io/fs.StatFS](https://pkg.go.dev/io/fs#StatFS), [io/fs.ReadFileFS](https://pkg.go.dev/io/fs#ReadFileFS) and [io/fs.ReadDirFS](https://pkg.go.dev/io/fs#ReadDirFS).

#### (*Root) Lstat <- 1.24.0

```go
func (r *Root) Lstat(name string) (FileInfo, error)
```

Lstat returns a [FileInfo](https://pkg.go.dev/os@go1.24.2#FileInfo) describing the named file in the root. If the file is a symbolic link, the returned FileInfo describes the symbolic link. See [Lstat](https://pkg.go.dev/os@go1.24.2#Lstat) for more details.

#### (*Root) Mkdir <- 1.24.0

```go
func (r *Root) Mkdir(name string, perm FileMode) error
```

Mkdir creates a new directory in the root with the specified name and permission bits (before umask). See [Mkdir](https://pkg.go.dev/os@go1.24.2#Mkdir) for more details.

If perm contains bits other than the nine least-significant bits (0o777), OpenFile returns an error.

#### (*Root) Name <- 1.24.0

```go
func (r *Root) Name() string
```

Name returns the name of the directory presented to OpenRoot.

It is safe to call Name after [Close].

#### (*Root) Open <- 1.24.0

```go
func (r *Root) Open(name string) (*File, error)
```

Open opens the named file in the root for reading. See [Open](https://pkg.go.dev/os@go1.24.2#Open) for more details.

#### (*Root) OpenFile <- 1.24.0

```go
func (r *Root) OpenFile(name string, flag int, perm FileMode) (*File, error)
```

OpenFile opens the named file in the root. See [OpenFile](https://pkg.go.dev/os@go1.24.2#OpenFile) for more details.

If perm contains bits other than the nine least-significant bits (0o777), OpenFile returns an error.

#### (*Root) OpenRoot <- 1.24.0

```go
func (r *Root) OpenRoot(name string) (*Root, error)
```

OpenRoot opens the named directory in the root. If there is an error, it will be of type *PathError.

#### (*Root) Remove <- 1.24.0

```go
func (r *Root) Remove(name string) error
```

Remove removes the named file or (empty) directory in the root. See [Remove](https://pkg.go.dev/os@go1.24.2#Remove) for more details.

#### (*Root) Stat <- 1.24.0

```go
func (r *Root) Stat(name string) (FileInfo, error)
```

Stat returns a [FileInfo](https://pkg.go.dev/os@go1.24.2#FileInfo) describing the named file in the root. See [Stat](https://pkg.go.dev/os@go1.24.2#Stat) for more details.

### type Signal 

``` go 
type Signal interface {
	String() string
	Signal() // to distinguish from other Stringers
}
```

A Signal represents an operating system signal. The usual underlying implementation is operating system-dependent: on Unix it is syscall.Signal.

​	Signal接口表示操作系统信号。通常底层实现是操作系统相关的：在Unix上，它是syscall.Signal。

``` go 
var (
	Interrupt Signal = syscall.SIGINT
	Kill      Signal = syscall.SIGKILL
)
```

The only signal values guaranteed to be present in the os package on all systems are os.Interrupt (send the process an interrupt) and os.Kill (force the process to exit). On Windows, sending os.Interrupt to a process with os.Process.Signal is not implemented; it will return an error instead of sending a signal.

​	在所有系统上os包中保证存在的唯一信号值是os.Interrupt(发送中断信号给进程)和os.Kill(强制进程退出)。在Windows上，使用os.Process.Signal将os.Interrupt发送到进程不起作用；它会返回一个错误，而不是发送信号。

### type SyscallError 

``` go 
type SyscallError struct {
	Syscall string
	Err     error
}
```

SyscallError records an error from a specific system call.

​	SyscallError结构体记录特定系统调用的错误。

#### (*SyscallError) Error 

``` go 
func (e *SyscallError) Error() string
```

#### (*SyscallError) Timeout  <- go1.10

``` go 
func (e *SyscallError) Timeout() bool
```

Timeout reports whether this error represents a timeout.

​	Timeout方法报告此错误是否表示超时。

#### (*SyscallError) Unwrap  <- go1.13

``` go 
func (e *SyscallError) Unwrap() error
```