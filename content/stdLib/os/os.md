+++
title = "os"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# os

https://pkg.go.dev/os@go1.20.1

​	os包提供了一个独立于平台的操作系统功能接口。设计类 Unix，但错误处理类 Go；失败的调用会返回类型为 error 而不是错误号的值。通常，在错误中还有更多的信息。例如，如果一个以文件名为参数的调用失败了，比如 Open 或 Stat，错误消息将在打印时包括失败的文件名，并且类型为 `*PathError`，可以拆开以获取更多信息。

​	os 接口旨在在所有操作系统中保持统一。一般不可用的功能会出现在特定于系统的 syscall 包中。

​	这里是一个简单的例子，打开一个文件并读取其中一部分。

```
file, err := os.Open("file.go") // 用于读取。
if err != nil {
	log.Fatal(err)
}
```

如果打开失败，错误字符串将是不言自明的，例如

```
open file.go: no such file or directory
```

然后可以将文件的数据读入到一个字节片中。Read 和 Write 的字节计数从参数切片的长度中获取。

```
data := make([]byte, 100)
count, err := file.Read(data)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])
```

注意：File 上的最大并发操作数可能受到操作系统或系统的限制。该数字应该很高，但超出它可能会降低性能或引起其他问题。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/file.go;l=72)

``` go 
const (
	// 必须精确指定 O_RDONLY、O_WRONLY 或 O_RDWR 中的一个。
	O_RDONLY int = syscall.O_RDONLY // 只读打开文件。
	O_WRONLY int = syscall.O_WRONLY // 只写打开文件。
	O_RDWR   int = syscall.O_RDWR   // 读写打开文件。
	// 其余的值可以进行或运算来控制行为。
	O_APPEND int = syscall.O_APPEND // 写入数据时追加到文件中。
	O_CREATE int = syscall.O_CREAT  // 如果不存在则创建一个新文件。
	O_EXCL   int = syscall.O_EXCL   // 与 O_CREATE 一起使用，文件必须不存在。
	O_SYNC   int = syscall.O_SYNC   // 为同步 I/O 打开文件。
	O_TRUNC  int = syscall.O_TRUNC  // 打开普通可写文件时截断。
)
```

​	OpenFile 的标志，包装了底层系统的标志。并非所有的标志都可能在给定系统上实现。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/file.go;l=88)

``` go 
const (
	SEEK_SET int = 0 // 从文件的起始位置寻址
	SEEK_CUR int = 1 // 相对于当前位置寻址
	SEEK_END int = 2 // 相对于文件末尾寻址
)
```

寻址偏移值。

已弃用：请使用 io.SeekStart、io.SeekCurrent 和 io.SeekEnd。

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
	// 单个字母是 String 方法格式化时使用的缩写。
	ModeDir        = fs.ModeDir        // d：目录
	ModeAppend     = fs.ModeAppend     // a：仅追加
	ModeExclusive  = fs.ModeExclusive  // l：独占使用
	ModeTemporary  = fs.ModeTemporary  // T：临时文件；仅 Plan 9
	ModeSymlink    = fs.ModeSymlink    // L：符号链接
	ModeDevice     = fs.ModeDevice     // D：设备文件
	ModeNamedPipe  = fs.ModeNamedPipe  // p：命名管道(FIFO)
	ModeSocket     = fs.ModeSocket     // S：Unix 域套接字
	ModeSetuid     = fs.ModeSetuid     // u: setuid
	ModeSetgid     = fs.ModeSetgid     // g: setgid
	ModeCharDevice = fs.ModeCharDevice // c：Unix 字符设备，在设置 ModeDevice 时
	ModeSticky     = fs.ModeSticky     // t：粘性位
	ModeIrregular  = fs.ModeIrregular  // ？：非常规文件；关于此文件没有其他已知信息

	// 类型比特的掩码。对于普通文件，不会设置任何位。
	ModeType = fs.ModeType

	ModePerm = fs.ModePerm //  Unix 权限位，0o777
)
```

​	已定义的文件模式比特是 FileMode 的最高有效位。最低有效位是标准 Unix rwxrwxrwx 权限位。这些位的值应该被视为公共 API 的一部分，可以在传输协议或磁盘表示中使用：不能更改它们，尽管可以添加新的位。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/file_unix.go;l=211)

``` go 
const DevNull = "/dev/null"
```

​	DevNull 是操作系统的"null 设备"的名称。在类 Unix 系统上，它是"/dev/null"；在 Windows 上，它是"NUL"。

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/error.go;l=16)

``` go 
var (
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

​	一些常见系统调用错误的便携式类比。

​	此包返回的错误可以使用 errors.Is 进行测试。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/file.go;l=64)

``` go 
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
```

​	Stdin、Stdout 和 Stderr 是打开的文件，它们指向标准输入、标准输出和标准错误文件描述符。

​	请注意，Go 运行时会将 panics 和 crashes 的消息写入标准错误；关闭 Stderr 可能会导致这些消息转到其他位置，例如稍后打开的文件。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/proc.go;l=16)

``` go 
var Args []string
```

​	Args 保存命令行参数，从程序名开始。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/os/exec.go;l=18)

``` go 
var ErrProcessDone = errors.New("os: process already finished")
```

​	ErrProcessDone 表示进程已经结束。

## 函数

#### func Chdir 

``` go 
func Chdir(dir string) error
```

​	Chdir函数将当前工作目录更改为指定的目录。如果出错，将返回 `*PathError` 类型的错误。

#### func Chmod 

``` go 
func Chmod(name string, mode FileMode) error
```

​	Chmod函数将指定文件的模式更改为 mode。如果该文件是符号链接，则更改链接目标的模式。如果出错，将返回 `*PathError` 类型的错误。

​	根据操作系统使用不同的模式比特的子集。

​	在 Unix 上，使用 mode 的权限位 ModeSetuid、ModeSetgid 和 ModeSticky。

​	在 Windows 上，仅使用 mode 的 0200 位(所有者可写)；它控制文件的只读属性是设置还是清除。其他位当前未使用。为了与 Go 1.12 及更早版本兼容，请使用非零模式。对于只读文件，请使用 mode 0400，对于可读写文件，请使用 mode 0600。

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



#### func Chown 

``` go 
func Chown(name string, uid, gid int) error
```

​	Chown函数更改指定文件的数值 UID 和 GID。如果该文件是符号链接，则更改链接目标的 UID 和 GID。UID 或 GID 的值为 -1 表示不更改该值。如果出错，将返回 *PathError 类型的错误。

​	在 Windows 或 Plan 9 上，Chown 始终返回 syscall.EWINDOWS 或 EPLAN9 错误，包装在 *PathError 中。

#### func Chtimes 

``` go 
func Chtimes(name string, atime time.Time, mtime time.Time) error
```

​	Chtimes函数更改指定文件的访问和修改时间，类似于 Unix 的 utime() 或 utimes() 函数。

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



#### func Clearenv 

``` go 
func Clearenv()
```

​	Clearenv函数删除所有环境变量。

#### func DirFS  <- go1.16

``` go 
func DirFS(dir string) fs.FS
```

​	DirFS函数返回以目录 dir 为根的文件树的文件系统(即 fs.FS)。

​	请注意，DirFS("/prefix") 只保证它对操作系统所做的 Open 调用将以 "/prefix" 开始：DirFS("/prefix").Open("file") 与 os.Open("/prefix/file") 相同。因此，如果 `/prefix/file` 是指向 `/prefix` 以外的符号链接，则使用 DirFS 不会比使用 os.Open 更停止访问。此外，对于相对路径，DirFS 返回的 fs.FS 的根目录，即 DirFS("prefix")，将受后续 Chdir 调用的影响。因此，当目录树包含任意内容时，DirFS 不是 chroot 类型的安全机制的通用替代品。

​	目录 dir 不得为 ""。

​	该结果实现 fs.StatFS。

#### func Environ 

``` go 
func Environ() []string
```

​	Environ函数返回表示环境变量的字符串副本，格式为"key=value"。

#### func Executable  <- go1.8

``` go 
func Executable() (string, error)
```

​	Executable函数返回启动当前进程的可执行文件的路径名。不能保证路径仍指向正确的可执行文件。如果使用符号链接启动了进程，则根据操作系统，结果可能是符号链接或它所指向的路径。如果需要稳定的结果，`path/filepath.EvalSymlinks`可能有所帮助。

​	Executable返回绝对路径，除非发生错误。

​	主要用例是查找相对于可执行文件的资源。

#### func Exit 

``` go 
func Exit(code int)
```

​	Exit函数使当前程序以给定的状态码退出。传统上，代码零表示成功，非零表示错误。程序立即终止；延迟函数不会运行。

​	为了可移植性，状态码应在[0，125]范围内。

#### func Expand 

``` go 
func Expand(s string, mapping func(string) string) string
```

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



#### func ExpandEnv 

``` go 
func ExpandEnv(s string) string
```

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



#### func Getegid 

``` go 
func Getegid() int
```

​	Getegid函数返回调用方的有效组ID。

​	在Windows上，它返回-1。

#### func Getenv 

``` go 
func Getenv(key string) string
```

​	Getenv函数检索由键指定的环境变量的值。它返回值，如果变量不存在，则为空。要区分空值和未设置的值，请使用LookupEnv。

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



#### func Geteuid 

``` go 
func Geteuid() int
```

​	Geteuid函数返回调用方的数字有效用户 ID。

​	在Windows上，它返回-1。

#### func Getgid 

``` go 
func Getgid() int
```

​	Getgid函数返回调用方的数字组 ID。

​	在Windows上，它返回-1。

#### func Getgroups 

``` go 
func Getgroups() ([]int, error)
```

​	Getgroups函数返回调用方所属的组的数字 ID 列表。

​	在Windows上，它返回syscall.EWINDOWS。有关可能的替代方案，请参见os/user包。

#### func Getpagesize 

``` go 
func Getpagesize() int
```

​	Getpagesize函数返回底层系统的内存页面大小。

#### func Getpid 

``` go 
func Getpid() int
```

​	Getpid函数返回调用方的进程 ID。

#### func Getppid 

``` go 
func Getppid() int
```

​	Getppid函数返回调用方的父进程 ID。

#### func Getuid 

``` go 
func Getuid() int
```

​	Getuid函数返回调用方的数字用户 ID。

​	在Windows上，它返回-1。

#### func Getwd 

``` go 
func Getwd() (dir string, err error)
```

​	Getwd函数返回对应于当前目录的根路径名。如果可以通过多个路径到达当前目录(由于符号链接)，Getwd可能返回其中任何一个。

#### func Hostname 

``` go 
func Hostname() (name string, err error)
```

​	Hostname函数返回内核报告的主机名。

#### func IsExist 

``` go 
func IsExist(err error) bool
```

​	IsExist函数返回一个布尔值，指示错误是否已知报告文件或目录已经存在。它满足ErrExist以及一些syscall错误。

​	此函数先于errors.Is。它仅支持由os包返回的错误。新代码应使用errors.Is(err，fs.ErrExist)。

#### func IsNotExist 

``` go 
func IsNotExist(err error) bool
```

​	IsNotExist函数返回一个布尔值，指示错误是否已知报告文件或目录不存在。它满足ErrNotExist以及一些syscall错误。

​	此函数先于errors.Is。它仅支持由os包返回的错误。新代码应使用errors.Is(err，fs.ErrNotExist)。

#### func IsPathSeparator 

``` go 
func IsPathSeparator(c uint8) bool
```

IsPathSeparator函数返回一个布尔值，指示c是否为目录分隔符字符。

#### func IsPermission 

``` go 
func IsPermission(err error) bool
```

​	IsPermission函数返回一个布尔值，指示错误是否已知报告权限被拒绝。它与某些系统调用错误以及ErrPermission相符。

​	此函数早于errors.Is。它仅支持由os包返回的错误。新代码应使用errors.Is(err，fs.ErrPermission)。

#### func IsTimeout  <- go1.10

``` go 
func IsTimeout(err error) bool
```

​	IsTimeout函数返回一个布尔值，指示错误是否已知报告超时发生。

​	此函数早于errors.Is，并且错误是否表示超时的概念可能不明确。例如，Unix错误EWOULDBLOCK有时指示超时，有时不是。新代码应使用与返回错误的调用相适应的值，例如os.ErrDeadlineExceeded。

#### func Lchown 

``` go 
func Lchown(name string, uid, gid int) error
```

​	Lchown函数更改命名文件的数值UID和GID。如果文件是符号链接，则更改链接本身的UID和GID。如果有错误，它将是*PathError类型。

​	在Windows上，它总是返回syscall.EWINDOWS错误，包装在*PathError中。

#### func Link 

``` go 
func Link(oldname, newname string) error
```

​	Link函数将newname创建为oldname文件的硬链接。如果有错误，它将是`*LinkError`类型。

#### func LookupEnv  <- go1.5

``` go 
func LookupEnv(key string) (string, bool)
```

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



#### func Mkdir 

``` go 
func Mkdir(name string, perm FileMode) error
```

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



#### func MkdirAll 

``` go 
func MkdirAll(path string, perm FileMode) error
```

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



#### func MkdirTemp  <- go1.16

``` go 
func MkdirTemp(dir, pattern string) (string, error)
```

​	MkdirTemp函数在目录dir中创建一个新的临时目录，并返回新目录的路径名。新目录的名称通过在模式的末尾添加一个随机字符串来生成。如果模式包含一个" `*`"，则随机字符串替换最后一个"`*`"。如果dir为空字符串，则MkdirTemp使用临时文件的默认目录，由TempDir返回。同时调用MkdirTemp的多个程序或goroutine不会选择相同的目录。当不再需要目录时，由调用者负责删除该目录。

##### MkdirTemp Example

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

##### MkdirTemp Example(Suffix)

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



#### func NewSyscallError 

``` go 
func NewSyscallError(syscall string, err error) error
```

​	NewSyscallError函数返回一个新的SyscallError，其中包含给定的系统调用名称和错误详细信息作为错误。为方便起见，如果err为nil，则NewSyscallError返回nil。

#### func Pipe 

``` go 
func Pipe() (r *File, w *File, err error)
```

​	Pipe函数返回一对连接的文件。从r读取的字节返回到w中。如果有错误，则返回文件和错误。

#### func ReadFile  <- go1.16

``` go 
func ReadFile(name string) ([]byte, error)
```

​	ReadFile函数读取指定的文件并返回文件内容。成功调用返回err == nil，而不是err == EOF。由于ReadFile读取整个文件，因此它不将从Read返回的EOF视为要报告的错误。

##### ReadFile Example

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



#### func Readlink 

``` go 
func Readlink(name string) (string, error)
```

​	Readlink函数返回指定符号链接的目标。如果有错误，则错误类型为* PathError。

#### func Remove 

``` go 
func Remove(name string) error
```

​	Remove函数删除指定的文件或(空)目录。如果有错误，则错误类型为`* PathError`。

#### func RemoveAll 

``` go 
func RemoveAll(path string) error
```

​	RemoveAll函数删除路径及其包含的所有子项。它会尽可能删除所有内容，但返回遇到的第一个错误。如果路径不存在，则RemoveAll返回nil(无错误)。如果有错误，则错误类型为`* PathError`。

#### func Rename 

``` go 
func Rename(oldpath, newpath string) error
```

​	Rename函数将oldpath重命名(移动)为newpath。如果newpath已经存在且不是目录，则会替换它。当oldpath和newpath在不同目录中时，可能会受到操作系统特定的限制。即使在同一个目录中，对于非Unix平台，Rename操作也不是原子性的。如果出错，错误类型为`*LinkError`。

#### func SameFile 

``` go 
func SameFile(fi1, fi2 FileInfo) bool
```

​	SameFile函数报告fi1和fi2描述的是否是同一文件。例如，在Unix上，这意味着两个底层结构的设备和inode字段是相同的；在其他系统上，决策可能基于路径名。SameFile仅适用于此包的Stat返回的结果。在其他情况下返回false。

#### func Setenv 

``` go 
func Setenv(key, value string) error
```

​	Setenv函数设置由键名key命名的环境变量的值。如果有任何错误，它将返回该错误。

#### func Symlink 

``` go 
func Symlink(oldname, newname string) error
```

​	Symlink函数将newname创建为指向oldname的符号链接。在Windows上，指向不存在的oldname的符号链接会创建一个文件符号链接；如果后来将oldname创建为目录，则符号链接将无法正常工作。如果出错，错误类型为`*LinkError`。

#### func TempDir 

``` go 
func TempDir() string
```

​	TempDir函数返回用于临时文件的默认目录。

​	在Unix系统上，如果$TMPDIR不为空，则返回它，否则返回/tmp。在Windows上，它使用GetTempPath，从%TMP%，%TEMP%，%USERPROFILE%或Windows目录中返回第一个非空值。在Plan 9上，它返回/tmp。

​	该目录既不保证存在，也不保证可访问权限。

#### func Truncate 

``` go 
func Truncate(name string, size int64) error
```

​	Truncate函数更改命名文件的大小。如果文件是符号链接，则更改链接目标的大小。如果出错，错误类型为`*PathError`。

#### func Unsetenv  <- go1.4

``` go 
func Unsetenv(key string) error
```

​	Unsetenv函数取消设置单个环境变量。

##### Unsetenv Example 

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



#### func UserCacheDir  <- go1.11

``` go 
func UserCacheDir() (string, error)
```

​	UserCacheDir函数返回用于用户特定缓存数据的默认根目录。用户应在其中创建自己的应用程序特定子目录并使用该目录。

​	在Unix系统上，如果`$XDG_CACHE_HOME`不为空，则根据https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html指定的内容返回该值，否则返回`$HOME/.cache`。在Darwin上，它返回`$HOME/Library/Caches`。在Windows上，它返回`%LocalAppData%`。在Plan 9上，它返回`$home/lib/cache`。

​	如果无法确定位置(例如，$HOME未定义)，则它将返回一个错误。

#### func UserConfigDir  <- go1.13

``` go 
func UserConfigDir() (string, error)
```

​	UserConfigDir函数返回用于用户特定配置数据的默认根目录。用户应在其中创建自己的应用程序特定子目录并使用该子目录。

​	在 Unix 系统上，如果非空，则返回 `$XDG_CONFIG_HOME`，否则返回 `$HOME/.config`。在 Darwin 上，它返回 `$HOME/Library/Application Support`。在 Windows 上，它返回 `%AppData%`。在 Plan 9 上，它返回 `$home/lib`。

​	如果无法确定位置(例如未定义 `$HOME`)，则返回错误。

#### func UserHomeDir  <- go1.12

``` go 
func UserHomeDir() (string, error)
```

​	UserHomeDir函数返回当前用户的主目录。

​	在 Unix 系统(包括 macOS)上，它返回 `$HOME` 环境变量。在 Windows 上，它返回`%USERPROFILE%`。在 Plan 9 上，它返回 `$home` 环境变量。

#### func WriteFile  <- go1.16

``` go 
func WriteFile(name string, data []byte, perm FileMode) error
```

`WriteFile函数将数据写入命名文件，如有必要创建它。如果文件不存在，则 WriteFile 使用权限 perm(在 umask 前)创建它；否则 WriteFile 在写入前将其截断，而不更改权限。由于 Writefile 需要多个系统调用才能完成，因此操作中的失败可能会使文件处于部分写入状态。

##### WriteFile Example

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

​	DirEntry 是从目录中读取的条目(使用 ReadDir 函数或 File 的 ReadDir 方法)。

#### func ReadDir  <- go1.16

``` go 
func ReadDir(name string) ([]DirEntry, error)
```

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

​	File结构体表示一个打开的文件描述符。

#### func Create 

``` go 
func Create(name string) (*File, error)
```

​	Create函数创建或截断命名文件。如果文件已经存在，则将其截断。如果文件不存在，则创建它以使用 mode 0666(在 umask 前)。如果成功，则返回的 File 上的方法可用于 I/O；相关的文件描述符具有 mode O_RDWR。如果有错误，它将是 `*PathError` 类型。

#### func CreateTemp  <- go1.16

``` go 
func CreateTemp(dir, pattern string) (*File, error)
```

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

​	NewFile函数返回具有给定文件描述符和名称的新文件。如果 fd 不是有效的文件描述符，则返回值将为 nil。在 Unix 系统上，如果文件描述符处于非阻塞模式，则 NewFile 将尝试返回可轮询的文件(其 SetDeadline 方法有效)。

​	将其传递给 NewFile 后，fd 可能会在与 Fd 方法的注释相同的条件下无效，同样的限制也适用。

#### func Open 

``` go 
func Open(name string) (*File, error)
```

​	Open函数打开具有指定名称的文件以供读取。如果成功，可以使用返回的文件上的方法进行读取。关联的文件描述符具有 O_RDONLY 模式。如果出现错误，类型为 `*PathError`。

#### func OpenFile 

``` go 
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

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



#### (*File) Chdir 

``` go 
func (f *File) Chdir() error
```

​	Chdir方法将当前工作目录更改为文件。它必须是目录。如果出现错误，类型为 `*PathError`。

#### (*File) Chmod 

``` go 
func (f *File) Chmod(mode FileMode) error
```

​	Chmod方法将文件的模式更改为 mode。如果出现错误，类型为 `*PathError`。

#### (*File) Chown 

``` go 
func (f *File) Chown(uid, gid int) error
```

​	Chown方法更改具有指定名称的文件的数字 uid 和 gid。如果出现错误，类型为 `*PathError`。

​	在 Windows 上，它始终返回 syscall.EWINDOWS 错误，包装在 `*PathError` 中。

#### (*File) Close 

``` go 
func (f *File) Close() error
```

​	Close方法关闭文件，使其不能用于I/O。对于支持SetDeadline的文件，任何待处理的I/O操作都将被取消，并立即返回ErrClosed错误。如果已经调用过Close方法，则Close将返回一个错误。

#### (*File) Fd 

``` go 
func (f *File) Fd() uintptr
```

​	Fd方法返回整数Unix文件描述符，引用打开的文件。如果f已关闭，则文件描述符将变为无效。如果f被垃圾回收，终结器可能会关闭文件描述符，使其无效；有关何时运行终结器的更多信息，请参见runtime.SetFinalizer。在Unix系统中，这将导致SetDeadline方法停止工作。因为文件描述符可以被重用，所以返回的文件描述符只能通过f的Close方法或在垃圾回收期间的终结器关闭。否则，在垃圾回收期间，终结器可能会关闭具有相同(重用)编号的不相关文件描述符。

​	作为替代方案，请参见f.SyscallConn方法。

#### (*File) Name 

``` go 
func (f *File) Name() string
```

​	Name方法返回打开文件的名称。

#### (*File) Read 

``` go 
func (f *File) Read(b []byte) (n int, err error)
```

​	Read方法从文件中读取最多len(b)字节，并将其存储在b中。它返回读取的字节数和任何遇到的错误。在文件末尾，Read返回0，io.EOF。

#### (*File) ReadAt 

``` go 
func (f *File) ReadAt(b []byte, off int64) (n int, err error)
```

​	ReadAt方法从文件中读取len(b)字节，从偏移量off处开始。它返回读取的字节数和错误(如果有)。当n < len(b)时，ReadAt总是返回非nil错误。在文件末尾，该错误为io.EOF。

#### (*File) ReadDir  <- go1.16

``` go 
func (f *File) ReadDir(n int) ([]DirEntry, error)
```

​	ReadDir方法读取与文件f相关联的目录的内容，并按文件名返回一系列DirEntry值。对同一文件的后续调用将按目录顺序返回后续的DirEntry记录。

​	如果n > 0，则ReadDir最多返回n个DirEntry记录。在这种情况下，如果ReadDir返回一个空切片，则它将返回一个解释原因的错误。在目录末尾，错误为io.EOF。

​	如果n <= 0，则ReadDir将返回剩余目录中的所有DirEntry记录。当它成功时，它返回一个nil错误(而不是io.EOF)。

#### (*File) ReadFrom  <- go1.15

``` go 
func (f *File) ReadFrom(r io.Reader) (n int64, err error)
```

​	ReadFrom方法实现了io.ReaderFrom。

#### (*File) Readdir 

``` go 
func (f *File) Readdir(n int) ([]FileInfo, error)
```

​	Readdir方法读取与文件 f 关联的目录并以目录顺序返回最多 n 个 FileInfo 值的切片，与 Lstat 返回的一样。在同一文件上的后续调用将返回更多的 FileInfos。

​	如果 n > 0，则 Readdir 返回最多 n 个 FileInfo 结构。在这种情况下，如果 Readdir 返回一个空切片，则它将返回一个非 nil 的错误来解释原因。在目录末尾，错误为 io.EOF。

​	如果 n <= 0，则 Readdir 以单个切片返回目录中的所有 FileInfo。在这种情况下，如果 Readdir 成功(一直读到目录结尾)，它将返回切片和 nil 错误。如果在目录结尾之前遇到错误，则 Readdir 返回读取到该点的 FileInfo 和非 nil 错误。

​	大多数客户端最好使用更高效的 ReadDir 方法。

#### (*File) Readdirnames 

``` go 
func (f *File) Readdirnames(n int) (names []string, err error)
```

​	Readdirnames方法读取与文件 f 关联的目录并按目录顺序返回最多 n 个文件名的切片。在同一文件上的后续调用将返回更多的文件名。

​	如果 n > 0，则 Readdirnames 返回最多 n 个文件名。在这种情况下，如果 Readdirnames 返回一个空切片，则它将返回一个非 nil 的错误来解释原因。在目录末尾，错误为 io.EOF。

​	如果 n <= 0，则 Readdirnames 以单个切片返回目录中的所有名称。在这种情况下，如果 Readdirnames 成功(一直读到目录结尾)，它将返回切片和 nil 错误。如果在目录结尾之前遇到错误，则 Readdirnames 返回读取到该点的名称和非 nil 错误。

#### (*File) Seek 

``` go 
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
```

​	Seek方法将下一个文件上的读取或写入的偏移量设置为 offset，根据 whence 进行解释：0 表示相对于文件起始处，1 表示相对于当前偏移量，2 表示相对于结尾。它返回新的偏移量和错误(如果有)。对于使用 O_APPEND 打开的文件，Seek 的行为未指定。

#### (*File) SetDeadline  <- go1.10

``` go 
func (f *File) SetDeadline(t time.Time) error
```

​	SetDeadline方法为文件设置读取和写入的截止日期。它等价于同时调用 SetReadDeadline 和 SetWriteDeadline。

​	只有某些类型的文件支持设置截止日期。调用不支持截止日期的文件的 SetDeadline方法将返回 ErrNoDeadline。在大多数系统上，普通文件不支持截止日期，但管道支持。

​	一个截止时间是一个绝对时间，在此之后，I/O 操作会失败并返回一个错误，而不是阻塞。截止时间适用于所有未来和待处理的 I/O，而不仅仅是对 Read 或 Write 的立即调用。超过截止时间后，可以通过设置将来的截止时间来刷新连接。

​	如果超过截止时间，对 Read 或 Write 或其他 I/O 方法的调用将返回一个包装 ErrDeadlineExceeded 的错误。可以使用 errors.Is(err，os.ErrDeadlineExceeded) 进行测试。该错误实现了 Timeout 方法，调用 Timeout 方法将返回 true，但有其他可能的错误，即使超时时间尚未超过，Timeout 也会返回 true。

​	可以通过在成功的 Read 或 Write 调用后反复延长截止时间来实现空闲超时。

​	t 的零值表示 I/O 操作不会超时。

#### (*File) SetReadDeadline  <- go1.10

``` go 
func (f *File) SetReadDeadline(t time.Time) error
```

​	SetReadDeadline方法设置将来的 Read 调用和任何当前阻塞的 Read 调用的截止时间。t 的零值表示 Read 不会超时。不是所有文件都支持设置截止时间；请参见 SetDeadline。

#### (*File) SetWriteDeadline  <- go1.10

``` go 
func (f *File) SetWriteDeadline(t time.Time) error
```

​	SetWriteDeadline方法设置任何将来的 Write 调用和任何当前阻塞的 Write 调用的截止时间。即使 Write 超时，它也可能返回 n > 0，表示某些数据已成功写入。t 的零值表示 Write 不会超时。不是所有文件都支持设置截止时间；请参见 SetDeadline。

#### (*File) Stat 

``` go 
func (f *File) Stat() (FileInfo, error)
```

​	Stat方法返回描述文件的 FileInfo 结构。如果有错误，它将是 `*PathError` 类型。

#### (*File) Sync 

``` go 
func (f *File) Sync() error
```

​	Sync方法将文件的当前内容提交到稳定存储。通常，这意味着将文件系统的最近写入数据的内存副本刷新到磁盘。

#### (*File) SyscallConn  <- go1.12

``` go 
func (f *File) SyscallConn() (syscall.RawConn, error)
```

​	SyscallConn方法返回一个原始文件。它实现了 syscall.Conn 接口。

#### (*File) Truncate 

``` go 
func (f *File) Truncate(size int64) error
```

​	Truncate方法改变文件的大小。它不会改变 I/O 偏移量。如果有错误，它将是 `*PathError` 类型。

#### (*File) Write 

``` go 
func (f *File) Write(b []byte) (n int, err error)
```

​	Write方法将 len(b) 个字节从 b 写入 File 中。它返回写入的字节数和错误(如果有)。当 n != len(b) 时，Write 返回非 nil 错误。

#### (*File) WriteAt 

``` go 
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
```

​	WriteAt方法从字节偏移量 off 处开始，将 len(b) 个字节写入 File。它返回写入的字节数和错误(如果有)。当 n != len(b) 时，WriteAt 返回非 nil 错误。

​	如果打开 file 时使用了 O_APPEND 标志，则 WriteAt 返回一个错误。

#### (*File) WriteString 

``` go 
func (f *File) WriteString(s string) (n int, err error)
```

​	WriteString方法类似于 Write，但它写入字符串 s 的内容而不是字节切片。

### type FileInfo 

``` go 
type FileInfo = fs.FileInfo
```

​	FileInfo 描述一个文件，并由 Stat 和 Lstat 返回。

#### func Lstat 

``` go 
func Lstat(name string) (FileInfo, error)
```

​	Lstat函数返回描述命名文件的 FileInfo。如果文件是符号链接，则返回的 FileInfo 描述符号链接。Lstat 不会尝试跟随链接。如果有错误，它将是 `*PathError` 类型的。

#### func Stat 

``` go 
func Stat(name string) (FileInfo, error)
```

​	Stat函数返回描述命名文件的 FileInfo。如果有错误，它将是 `*PathError` 类型的。

### type FileMode 

``` go 
type FileMode = fs.FileMode
```

​	FileMode表示文件的模式和权限位。这些位在所有系统上具有相同的定义，因此可以将关于文件的信息可移植地从一个系统移动到另一个系统。并非所有位都适用于所有系统。ModeDir 用于目录是唯一需要的位。

##### FileMode Example

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

​	PathError类型记录了错误和引起错误的操作和文件路径。

### type ProcAttr 

``` go 
type ProcAttr struct {
	// 如果Dir不为空，则子进程在创建进程之前进入该目录。
	Dir string
	// 如果Env不为nil，则以Environ返回的形式为新进程提供环境变量。
	// 如果为nil，则使用Environ的结果。
	Env []string
	// Files指定新进程继承的打开文件。
    // 前三个条目对应标准输入、标准输出和标准错误输出。
	// 根据底层操作系统的不同，实现可能会支持更多的条目。
    // 空条目表示该文件在进程启动时关闭。
	// 在Unix系统上，StartProcess将把这些File值更改为阻塞模式，
    // 这意味着SetDeadline将停止工作，并且调用Close不会中断读取或写入。
	Files []*File

	// 操作系统特定的进程创建属性。
	// 请注意，设置此字段意味着您的程序可能
    // 无法在某些操作系统上正确执行甚至无法编译。
	Sys *syscall.SysProcAttr
}
```

​	ProcAttr结构体保存将应用于由StartProcess启动的新进程的属性。

### type Process 

``` go 
type Process struct {
	Pid int
	// contains filtered or unexported fields
	// 包含已过滤或未导出的字段
}
```

​	Process结构体存储由StartProcess创建的进程的信息。

#### func FindProcess 

``` go 
func FindProcess(pid int) (*Process, error)
```

​	FindProcess函数按其pid查找正在运行的进程。

​	它返回的Process可用于获取有关底层操作系统进程的信息。

​	在Unix系统上，FindProcess始终成功并返回给定pid的Process，无论进程是否存在。

#### func StartProcess 

``` go 
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
```

​	StartProcess函数使用由name、argv和attr指定的程序、参数和属性启动一个新的进程。argv切片将成为新进程的os.Args，因此它通常以程序名称开头。

​	如果调用的goroutine已使用runtime.LockOSThread锁定操作系统线程并修改了任何可继承的操作系统级线程状态(例如Linux或Plan 9名称空间)，新进程将继承调用者的线程状态。

​	StartProcess是一个低级接口。os/exec包提供了更高级的接口。

​	如果有错误，则为`*PathError`类型。

#### (*Process) Kill 

``` go 
func (p *Process) Kill() error
```

​	Kill方法使进程立即退出。Kill不会等待进程实际退出。这只会杀死进程本身，而不是它可能启动的任何其他进程。

#### (*Process) Release 

``` go 
func (p *Process) Release() error
```

​	Release方法释放与进程p关联的任何资源，使其在将来无法使用。如果不使用Wait，则只需要调用Release。

#### (*Process) Signal 

``` go 
func (p *Process) Signal(sig Signal) error
```

​	Signal方法向进程发送信号。在Windows上发送中断信号未实现。

#### (*Process) Wait 

``` go 
func (p *Process) Wait() (*ProcessState, error)
```

​	Wait方法等待进程退出，然后返回ProcessState描述其状态和任何错误。Wait释放与进程关联的任何资源。在大多数操作系统上，进程必须是当前进程的子进程，否则将返回一个错误。

### type ProcessState 

``` go 
type ProcessState struct {
	// contains filtered or unexported fields
}
```

​	ProcessState结构体存储关于进程的信息，由Wait方法报告。

#### (*ProcessState) ExitCode  <- go1.12

``` go 
func (p *ProcessState) ExitCode() int
```

​	ExitCode方法返回退出的进程的退出代码，如果进程尚未退出或被信号终止，则返回-1。

#### (*ProcessState) Exited 

``` go 
func (p *ProcessState) Exited() bool
```

​	Exited方法报告程序是否已退出。在Unix系统上，如果程序由于调用exit而退出，则此项报告为true，但如果程序由于信号终止而终止，则此项报告为false。

#### (*ProcessState) Pid 

``` go 
func (p *ProcessState) Pid() int
```

​	Pid方法返回已退出进程的进程ID。

#### (*ProcessState) String 

``` go 
func (p *ProcessState) String() string
```

#### (*ProcessState) Success 

``` go 
func (p *ProcessState) Success() bool
```

​	Success方法报告程序是否成功退出，例如在Unix上以退出状态0退出。

#### (*ProcessState) Sys 

``` go 
func (p *ProcessState) Sys() any
```

​	Sys方法返回有关进程的系统相关退出信息。将其转换为适当的底层类型，例如在Unix上的syscall.WaitStatus，以访问其内容。

#### (*ProcessState) SysUsage 

``` go 
func (p *ProcessState) SysUsage() any
```

​	SysUsage方法返回有关已退出进程的系统相关资源使用情况信息。将其转换为适当的底层类型，例如在Unix上的`*syscall.Rusage`，以访问其内容。(在Unix上，`*syscall.Rusage`与getrusage(2)手册页中定义的struct rusage匹配。)

#### (*ProcessState) SystemTime 

``` go 
func (p *ProcessState) SystemTime() time.Duration
```

​	SystemTime方法返回已退出进程及其子进程的系统CPU时间。

#### (*ProcessState) UserTime 

``` go 
func (p *ProcessState) UserTime() time.Duration
```

​	UserTime方法返回已退出进程及其子进程的用户CPU时间。

### type Signal 

``` go 
type Signal interface {
	String() string
	Signal() // to distinguish from other Stringers
}
```

​	Signal接口表示操作系统信号。通常底层实现是操作系统相关的：在Unix上，它是syscall.Signal。

``` go 
var (
	Interrupt Signal = syscall.SIGINT
	Kill      Signal = syscall.SIGKILL
)
```

​	在所有系统上os包中保证存在的唯一信号值是os.Interrupt(发送中断信号给进程)和os.Kill(强制进程退出)。在Windows上，使用os.Process.Signal将os.Interrupt发送到进程不起作用；它会返回一个错误，而不是发送信号。

### type SyscallError 

``` go 
type SyscallError struct {
	Syscall string
	Err     error
}
```

​	SyscallError结构体记录特定系统调用的错误。

#### (*SyscallError) Error 

``` go 
func (e *SyscallError) Error() string
```

#### (*SyscallError) Timeout  <- go1.10

``` go 
func (e *SyscallError) Timeout() bool
```

​	Timeout方法报告此错误是否表示超时。

#### (*SyscallError) Unwrap  <- go1.13

``` go 
func (e *SyscallError) Unwrap() error
```