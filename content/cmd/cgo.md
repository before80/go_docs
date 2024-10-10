+++
title = "cgo"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# cgo

> 原文：https://pkg.go.dev/cmd/cgo@go1.19.3

### Overview

Cgo enables the creation of Go packages that call C code.

​	cgo 允许创建调用 C 代码的 Go 包。

#### Using cgo with the go command

To use cgo write normal Go code that imports a pseudo-package "C". The Go code can then refer to types such as C.size_t, variables such as C.stdout, or functions such as C.putchar.

​	要使用 cgo，请编写导入伪包“C”的普通 Go 代码。然后，Go 代码可以引用类型（例如 C.size_t）、变量（例如 C.stdout）或函数（例如 C.putchar）。

If the import of "C" is immediately preceded by a comment, that comment, called the preamble, is used as a header when compiling the C parts of the package. For example:

​	如果“C”的导入紧接在注释之前，则该注释（称为前导注释）在编译包的 C 部分时用作头文件。例如：

```
// #include <stdio.h>
// #include <errno.h>
import "C"
```

The preamble may contain any C code, including function and variable declarations and definitions. These may then be referred to from Go code as though they were defined in the package "C". All names declared in the preamble may be used, even if they start with a lower-case letter. Exception: static variables in the preamble may not be referenced from Go code; static functions are permitted.

​	前导注释可以包含任何 C 代码，包括函数和变量声明及定义。然后，可以从 Go 代码中引用这些内容，就好像它们是在包“C”中定义的一样。可以使用的所有名称都声明在前导注释中，即使它们以小写字母开头。例外：前导注释中的静态变量不能从 Go 代码中引用；静态函数是允许的。

See `$GOROOT/misc/cgo/stdio` and `$GOROOT/misc/cgo/gmp` for examples. See "C? Go? Cgo!" for an introduction to using cgo: https://golang.org/doc/articles/c_go_cgo.html.

​	有关示例，请参阅 `$GOROOT/misc/cgo/stdio` 和 `$GOROOT/misc/cgo/gmp`。有关使用 cgo 的介绍，请参阅“C? Go? Cgo!”：https://golang.org/doc/articles/c_go_cgo.html。

CFLAGS, CPPFLAGS, CXXFLAGS, FFLAGS and LDFLAGS may be defined with pseudo #cgo directives within these comments to tweak the behavior of the C, C++ or Fortran compiler. Values defined in multiple directives are concatenated together. The directive can include a list of build constraints limiting its effect to systems satisfying one of the constraints (see https://golang.org/pkg/go/build/#hdr-Build_Constraints for details about the constraint syntax). For example:

​	可以在这些注释中使用伪 #cgo 指令定义 CFLAGS、CPPFLAGS、CXXFLAGS、FFLAGS 和 LDFLAGS，以调整 C、C++ 或 Fortran 编译器行为。在多个指令中定义的值会连接在一起。该指令可以包括一个构建约束列表，将其效果限制为满足其中一个约束的系统（有关约束语法的详细信息，请参阅 https://golang.org/pkg/go/build/#hdr-Build_Constraints）。例如：

```
// #cgo CFLAGS: -DPNG_DEBUG=1
// #cgo amd64 386 CFLAGS: -DX86=1
// #cgo LDFLAGS: -lpng
// #include <png.h>
import "C"
```

Alternatively, CPPFLAGS and LDFLAGS may be obtained via the pkg-config tool using a '#cgo pkg-config:' directive followed by the package names. For example:

​	或者，可以使用 '#cgo pkg-config:' 指令（后跟包名称）通过 pkg-config 工具获取 CPPFLAGS 和 LDFLAGS。例如：

```
// #cgo pkg-config: png cairo
// #include <png.h>
import "C"
```

The default pkg-config tool may be changed by setting the PKG_CONFIG environment variable.

​	可以通过设置 PKG_CONFIG 环境变量来更改默认的 pkg-config 工具。

For security reasons, only a limited set of flags are allowed, notably -D, -U, -I, and -l. To allow additional flags, set CGO_CFLAGS_ALLOW to a regular expression matching the new flags. To disallow flags that would otherwise be allowed, set CGO_CFLAGS_DISALLOW to a regular expression matching arguments that must be disallowed. In both cases the regular expression must match a full argument: to allow -mfoo=bar, use CGO_CFLAGS_ALLOW='-mfoo.*', not just CGO_CFLAGS_ALLOW='-mfoo'. Similarly named variables control the allowed CPPFLAGS, CXXFLAGS, FFLAGS, and LDFLAGS.

​	出于安全原因，只允许有限的一组标志，特别是 -D、-U、-I 和 -l。要允许其他标志，请将 CGO_CFLAGS_ALLOW 设置为与新标志匹配的正则表达式。要禁止本来允许的标志，请将 CGO_CFLAGS_DISALLOW 设置为与必须禁止的参数匹配的正则表达式。在这两种情况下，正则表达式必须与一个完整参数匹配：要允许 -mfoo=bar，请使用 CGO_CFLAGS_ALLOW='-mfoo.*'，而不仅仅是 CGO_CFLAGS_ALLOW='-mfoo'。类似命名的变量控制允许的 CPPFLAGS、CXXFLAGS、FFLAGS 和 LDFLAGS。

Also for security reasons, only a limited set of characters are permitted, notably alphanumeric characters and a few symbols, such as '.', that will not be interpreted in unexpected ways. Attempts to use forbidden characters will get a "malformed #cgo argument" error.

​	同样出于安全原因，只允许有限的一组字符，特别是字母数字字符和一些符号，例如不会以意外方式解释的“.”。尝试使用禁止的字符将收到“格式错误的 #cgo 参数”错误。

When building, the CGO_CFLAGS, CGO_CPPFLAGS, CGO_CXXFLAGS, CGO_FFLAGS and CGO_LDFLAGS environment variables are added to the flags derived from these directives. Package-specific flags should be set using the directives, not the environment variables, so that builds work in unmodified environments. Flags obtained from environment variables are not subject to the security limitations described above.

​	在构建时，CGO_CFLAGS、CGO_CPPFLAGS、CGO_CXXFLAGS、CGO_FFLAGS 和 CGO_LDFLAGS 环境变量会添加到从这些指令派生的标志中。应使用指令而不是环境变量设置特定于包的标志，以便构建在未修改的环境中工作。从环境变量获取的标志不受上述安全限制的约束。

All the cgo CPPFLAGS and CFLAGS directives in a package are concatenated and used to compile C files in that package. All the CPPFLAGS and CXXFLAGS directives in a package are concatenated and used to compile C++ files in that package. All the CPPFLAGS and FFLAGS directives in a package are concatenated and used to compile Fortran files in that package. All the LDFLAGS directives in any package in the program are concatenated and used at link time. All the pkg-config directives are concatenated and sent to pkg-config simultaneously to add to each appropriate set of command-line flags.

​	包中的所有 cgo CPPFLAGS 和 CFLAGS 指令会连接起来，并用于编译该包中的 C 文件。包中的所有 CPPFLAGS 和 CXXFLAGS 指令会连接起来，并用于编译该包中的 C++ 文件。包中的所有 CPPFLAGS 和 FFLAGS 指令会连接起来，并用于编译该包中的 Fortran 文件。程序中任何包中的所有 LDFLAGS 指令会连接起来，并在链接时使用。所有 pkg-config 指令会连接起来，并同时发送到 pkg-config，以添加到每组相应的命令行标志中。

When the cgo directives are parsed, any occurrence of the string `${SRCDIR}` will be replaced by the absolute path to the directory containing the source file. This allows pre-compiled static libraries to be included in the package directory and linked properly. For example if package foo is in the directory /go/src/foo:

​	当解析 cgo 指令时，字符串 `${SRCDIR}` 的任何出现都将被替换为包含源文件的目录的绝对路径。这允许将预编译的静态库包含在包目录中并正确链接。例如，如果包 foo 在目录 /go/src/foo 中：

```
// #cgo LDFLAGS: -L${SRCDIR}/libs -lfoo
```

Will be expanded to:

​	将展开为：

```
// #cgo LDFLAGS: -L/go/src/foo/libs -lfoo
```

When the Go tool sees that one or more Go files use the special import "C", it will look for other non-Go files in the directory and compile them as part of the Go package. Any .c, .s, .S or .sx files will be compiled with the C compiler. Any .cc, .cpp, or .cxx files will be compiled with the C++ compiler. Any .f, .F, .for or .f90 files will be compiled with the fortran compiler. Any .h, .hh, .hpp, or .hxx files will not be compiled separately, but, if these header files are changed, the package (including its non-Go source files) will be recompiled. Note that changes to files in other directories do not cause the package to be recompiled, so all non-Go source code for the package should be stored in the package directory, not in subdirectories. The default C and C++ compilers may be changed by the CC and CXX environment variables, respectively; those environment variables may include command line options.

​	当 Go 工具看到一个或多个 Go 文件使用特殊导入“C”时，它将在目录中查找其他非 Go 文件，并将它们作为 Go 包的一部分进行编译。任何 .c、.s、.S 或 .sx 文件都将使用 C 编译器进行编译。任何 .cc、.cpp 或 .cxx 文件都将使用 C++ 编译器进行编译。任何 .f、.F、.for 或 .f90 文件都将使用 fortran 编译器进行编译。任何 .h、.hh、.hpp 或 .hxx 文件都不会单独编译，但是，如果这些头文件发生更改，则会重新编译该包（包括其非 Go 源文件）。请注意，其他目录中的文件发生更改不会导致重新编译该包，因此该包的所有非 Go 源代码都应存储在包目录中，而不是子目录中。默认的 C 和 C++ 编译器可以分别通过 CC 和 CXX 环境变量进行更改；这些环境变量可能包括命令行选项。

The cgo tool will always invoke the C compiler with the source file's directory in the include path; i.e. -I${SRCDIR} is always implied. This means that if a header file foo/bar.h exists both in the source directory and also in the system include directory (or some other place specified by a -I flag), then "#include <foo/bar.h>" will always find the local version in preference to any other version.

​	cgo 工具将始终在包含路径中使用源文件的目录调用 C 编译器；即始终暗示 -I${SRCDIR}。这意味着如果头文件 foo/bar.h 同时存在于源目录和系统包含目录（或由 -I 标志指定的其他位置），则 "#include " 将始终优先于任何其他版本找到本地版本。

The cgo tool is enabled by default for native builds on systems where it is expected to work. It is disabled by default when cross-compiling. You can control this by setting the CGO_ENABLED environment variable when running the go tool: set it to 1 to enable the use of cgo, and to 0 to disable it. The go tool will set the build constraint "cgo" if cgo is enabled. The special import "C" implies the "cgo" build constraint, as though the file also said "// +build cgo". Therefore, if cgo is disabled, files that import "C" will not be built by the go tool. (For more about build constraints see https://golang.org/pkg/go/build/#hdr-Build_Constraints).

​	在预期 cgo 工具可以正常工作的系统上，该工具默认情况下处于启用状态。在交叉编译时，该工具默认情况下处于禁用状态。您可以在运行 go 工具时通过设置 CGO_ENABLED 环境变量来控制此行为：将其设置为 1 以启用 cgo 的使用，将其设置为 0 以禁用 cgo。如果启用了 cgo，go 工具将设置构建约束“cgo”。特殊导入“C”暗示“cgo”构建约束，就好像该文件还说“// +build cgo”。因此，如果禁用了 cgo，go 工具将不会构建导入“C”的文件。（有关构建约束的更多信息，请参阅 https://golang.org/pkg/go/build/#hdr-Build_Constraints）。

When cross-compiling, you must specify a C cross-compiler for cgo to use. You can do this by setting the generic CC_FOR_TARGET or the more specific `CC_FOR_${GOOS}_${GOARCH}` (for example, CC_FOR_linux_arm) environment variable when building the toolchain using make.bash, or you can set the CC environment variable any time you run the go tool.

​	在进行交叉编译时，您必须为 cgo 指定一个 C 交叉编译器。您可以在使用 make.bash 构建工具链时，通过设置通用的 CC_FOR_TARGET 或更具体的 `CC_FOR_${GOOS}_${GOARCH}`（例如，CC_FOR_linux_arm）环境变量来执行此操作，或者您可以在任何时候运行 go 工具时设置 CC 环境变量。

The CXX_FOR_TARGET, `CXX_FOR_${GOOS}_${GOARCH}`, and CXX environment variables work in a similar way for C++ code.

​	CXX_FOR_TARGET、`CXX_FOR_${GOOS}_${GOARCH}` 和 CXX 环境变量对 C++ 代码的作用方式类似。

#### Go references to C

Within the Go file, C's struct field names that are keywords in Go can be accessed by prefixing them with an underscore: if x points at a C struct with a field named "type", x._type accesses the field. C struct fields that cannot be expressed in Go, such as bit fields or misaligned data, are omitted in the Go struct, replaced by appropriate padding to reach the next field or the end of the struct.

​	在 Go 文件中，可以通过在 C 的结构字段名前加上下划线来访问 Go 中的关键字：如果 x 指向具有名为“type”的字段的 C 结构，则 x._type 访问该字段。无法在 Go 中表示的 C 结构字段（例如位字段或未对齐的数据）在 Go 结构中被省略，并用适当的填充来替换以到达下一个字段或结构的末尾。

The standard C numeric types are available under the names C.char, C.schar (signed char), C.uchar (unsigned char), C.short, C.ushort (unsigned short), C.int, C.uint (unsigned int), C.long, C.ulong (unsigned long), C.longlong (long long), C.ulonglong (unsigned long long), C.float, C.double, C.complexfloat (complex float), and C.complexdouble (complex double). The C type `void*` is represented by Go's unsafe.Pointer. The C types `__int128_t` and `__uint128_t` are represented by [16]byte.

​	标准 C 数值类型可通过以下名称使用：C.char、C.schar（有符号 char）、C.uchar（无符号 char）、C.short、C.ushort（无符号 short）、C.int、C.uint（无符号 int）、C.long、C.ulong（无符号 long）、C.longlong（long long）、C.ulonglong（无符号 long long）、C.float、C.double、C.complexfloat（复数 float）和 C.complexdouble（复数 double）。C 类型 `void*` 由 Go 的 unsafe.Pointer 表示。C 类型 `__int128_t` 和` __uint128_t` 由 [16]byte 表示。

A few special C types which would normally be represented by a pointer type in Go are instead represented by a uintptr. See the Special cases section below.

​	一些通常由 Go 中的指针类型表示的特殊 C 类型改由 uintptr 表示。请参阅下面的特殊情况部分。

To access a struct, union, or enum type directly, prefix it with struct_, union_, or enum_, as in C.struct_stat.

​	要直接访问结构体、联合或枚举类型，请在其前面加上 struct_、union_或 enum_，如 C.struct_stat。

The size of any C type T is available as C.sizeof_T, as in C.sizeof_struct_stat.

​	任何 C 类型 T 的大小都可作为 C.sizeof_T 使用，如 C.sizeof_struct_stat。

A C function may be declared in the Go file with a parameter type of the special name _GoString_. This function may be called with an ordinary Go string value. The string length, and a pointer to the string contents, may be accessed by calling the C functions

​	可以在 Go 文件中声明一个 C 函数，其参数类型为特殊名称 _GoString_。可以使用普通 Go 字符串值调用此函数。可以通过调用 C 函数访问字符串长度和指向字符串内容的指针

```
size_t _GoStringLen(_GoString_ s);
const char *_GoStringPtr(_GoString_ s);
```

These functions are only available in the preamble, not in other C files. The C code must not modify the contents of the pointer returned by _GoStringPtr. Note that the string contents may not have a trailing NUL byte.

​	这些函数仅在序言中可用，在其他 C 文件中不可用。C 代码不得修改 _GoStringPtr 返回的指针的内容。请注意，字符串内容可能没有尾随 NUL 字节。

As Go doesn't have support for C's union type in the general case, C's union types are represented as a Go byte array with the same length.

​	由于 Go 在一般情况下不支持 C 的联合类型，因此 C 的联合类型表示为具有相同长度的 Go 字节数组。

Go structs cannot embed fields with C types.

​	Go 结构不能嵌入具有 C 类型的字段。

Go code cannot refer to zero-sized fields that occur at the end of non-empty C structs. To get the address of such a field (which is the only operation you can do with a zero-sized field) you must take the address of the struct and add the size of the struct.

​	Go 代码不能引用出现在非空 C 结构末尾的零大小字段。要获取此类字段的地址（这是您可以对零大小字段执行的唯一操作），您必须获取结构的地址并添加结构的大小。

Cgo translates C types into equivalent unexported Go types. Because the translations are unexported, a Go package should not expose C types in its exported API: a C type used in one Go package is different from the same C type used in another.

​	Cgo 将 C 类型转换为等效的未导出的 Go 类型。由于这些转换是未导出的，因此 Go 包不应在其导出的 API 中公开 C 类型：在一个 Go 包中使用的 C 类型不同于在另一个 Go 包中使用的相同 C 类型。

Any C function (even void functions) may be called in a multiple assignment context to retrieve both the return value (if any) and the C errno variable as an error (use _ to skip the result value if the function returns void). For example:

​	可以在多重赋值上下文中调用任何 C 函数（甚至是 void 函数），以检索返回值（如果有）和 C errno 变量作为错误（如果函数返回 void，请使用 _ 跳过结果值）。例如：

```
n, err = C.sqrt(-1)
_, err := C.voidFunc()
var n, err = C.sqrt(1)
```

Calling C function pointers is currently not supported, however you can declare Go variables which hold C function pointers and pass them back and forth between Go and C. C code may call function pointers received from Go. For example:

​	目前不支持调用 C 函数指针，但是您可以声明保存 C 函数指针的 Go 变量，并在 Go 和 C 之间来回传递它们。C 代码可以调用从 Go 接收到的函数指针。例如：

```
package main

// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int fortytwo()
// {
//	    return 42;
// }
import "C"
import "fmt"

func main() {
	f := C.intFunc(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))
	// Output: 42
}
```

In C, a function argument written as a fixed size array actually requires a pointer to the first element of the array. C compilers are aware of this calling convention and adjust the call accordingly, but Go cannot. In Go, you must pass the pointer to the first element explicitly: C.f(&C.x[0]).

​	在 C 中，以固定大小数组形式编写的函数参数实际上需要指向数组第一个元素的指针。C 编译器知道此调用约定并相应地调整调用，但 Go 无法做到。在 Go 中，您必须显式传递指向第一个元素的指针：C.f(&C.x[0])。

Calling variadic C functions is not supported. It is possible to circumvent this by using a C function wrapper. For example:

​	不支持调用可变参数 C 函数。可以通过使用 C 函数包装器来规避此问题。例如：

```
package main

// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(char* s) {
//   printf("%s\n", s);
// }
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
```

A few special functions convert between Go and C types by making copies of the data. In pseudo-Go definitions:

​	少数特殊函数通过复制数据在 Go 和 C 类型之间进行转换。在伪 Go 定义中：

```
// Go string to C string
// The C string is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func C.CString(string) *C.char

// Go []byte slice to C array
// The C array is allocated in the C heap using malloc.
// It is the caller's responsibility to arrange for it to be
// freed, such as by calling C.free (be sure to include stdlib.h
// if C.free is needed).
func C.CBytes([]byte) unsafe.Pointer

// C string to Go string
func C.GoString(*C.char) string

// C data with explicit length to Go string
func C.GoStringN(*C.char, C.int) string

// C data with explicit length to Go []byte
func C.GoBytes(unsafe.Pointer, C.int) []byte
```

As a special case, C.malloc does not call the C library malloc directly but instead calls a Go helper function that wraps the C library malloc but guarantees never to return nil. If C's malloc indicates out of memory, the helper function crashes the program, like when Go itself runs out of memory. Because C.malloc cannot fail, it has no two-result form that returns errno.

​	作为特例，C.malloc 不会直接调用 C 库 malloc，而是调用一个 Go 帮助器函数，该函数包装 C 库 malloc，但保证绝不会返回 nil。如果 C 的 malloc 指示内存不足，则帮助器函数会使程序崩溃，就像 Go 本身内存不足时一样。由于 C.malloc 不会失败，因此它没有返回 errno 的双结果形式。

#### C references to Go

Go functions can be exported for use by C code in the following way:

​	Go 函数可以通过以下方式导出以供 C 代码使用：

```
//export MyFunction
func MyFunction(arg1, arg2 int, arg3 string) int64 {...}

//export MyFunction2
func MyFunction2(arg1, arg2 int, arg3 string) (int64, *C.char) {...}
```

They will be available in the C code as:

​	它们将在 C 代码中作为：

```
extern GoInt64 MyFunction(int arg1, int arg2, GoString arg3);
extern struct MyFunction2_return MyFunction2(int arg1, int arg2, GoString arg3);
```

found in the _cgo_export.h generated header, after any preambles copied from the cgo input files. Functions with multiple return values are mapped to functions returning a struct.

​	在从 cgo 输入文件中复制的任何前言之后，在生成的 _cgo_export.h 头文件中找到。具有多个返回值的函数映射到返回结构的函数。

Not all Go types can be mapped to C types in a useful way. Go struct types are not supported; use a C struct type. Go array types are not supported; use a C pointer.

​	并非所有 Go 类型都可以以有用的方式映射到 C 类型。不支持 Go 结构类型；使用 C 结构类型。不支持 Go 数组类型；使用 C 指针。

Go functions that take arguments of type string may be called with the C type _GoString_, described above. The _GoString_ type will be automatically defined in the preamble. Note that there is no way for C code to create a value of this type; this is only useful for passing string values from Go to C and back to Go.

​	可以带字符串类型参数的 Go 函数可以使用上面描述的 C 类型 _GoString_ 调用。_GoString_ 类型将在前言中自动定义。请注意，C 代码无法创建此类型的变量；这仅适用于将字符串值从 Go 传递到 C，再从 C 传递回 Go。

Using //export in a file places a restriction on the preamble: since it is copied into two different C output files, it must not contain any definitions, only declarations. If a file contains both definitions and declarations, then the two output files will produce duplicate symbols and the linker will fail. To avoid this, definitions must be placed in preambles in other files, or in C source files.

​	在文件中使用 //export 会对前言施加限制：由于它被复制到两个不同的 C 输出文件中，因此它不能包含任何定义，只能包含声明。如果一个文件同时包含定义和声明，那么两个输出文件将生成重复的符号，链接器将失败。为避免这种情况，必须将定义放在其他文件的前言中，或放在 C 源文件中。

#### Passing pointers

Go is a garbage collected language, and the garbage collector needs to know the location of every pointer to Go memory. Because of this, there are restrictions on passing pointers between Go and C.

​	Go 是一种垃圾回收语言，垃圾回收器需要知道每个指向 Go 内存的指针的位置。因此，在 Go 和 C 之间传递指针存在限制。

In this section the term Go pointer means a pointer to memory allocated by Go (such as by using the & operator or calling the predefined new function) and the term C pointer means a pointer to memory allocated by C (such as by a call to C.malloc). Whether a pointer is a Go pointer or a C pointer is a dynamic property determined by how the memory was allocated; it has nothing to do with the type of the pointer.

​	在本节中，术语 Go 指针是指向由 Go 分配的内存的指针（例如，通过使用 & 运算符或调用预定义的 new 函数），术语 C 指针是指向由 C 分配的内存的指针（例如，通过调用 C.malloc）。指针是 Go 指针还是 C 指针是一个动态属性，由内存的分配方式决定；它与指针的类型无关。

Note that values of some Go types, other than the type's zero value, always include Go pointers. This is true of string, slice, interface, channel, map, and function types. A pointer type may hold a Go pointer or a C pointer. Array and struct types may or may not include Go pointers, depending on the element types. All the discussion below about Go pointers applies not just to pointer types, but also to other types that include Go pointers.

​	请注意，某些 Go 类型的非零值总是包含 Go 指针。这适用于字符串、切片、接口、通道、映射和函数类型。指针类型可以保存 Go 指针或 C 指针。数组和结构类型可能包含或不包含 Go 指针，具体取决于元素类型。下面关于 Go 指针的所有讨论不仅适用于指针类型，也适用于包含 Go 指针的其他类型。

Go code may pass a Go pointer to C provided the Go memory to which it points does not contain any Go pointers. The C code must preserve this property: it must not store any Go pointers in Go memory, even temporarily. When passing a pointer to a field in a struct, the Go memory in question is the memory occupied by the field, not the entire struct. When passing a pointer to an element in an array or slice, the Go memory in question is the entire array or the entire backing array of the slice.

​	Go 代码可以传递一个 Go 指针给 C，前提是它所指向的 Go 内存不包含任何 Go 指针。C 代码必须保留此属性：它不能在 Go 内存中存储任何 Go 指针，即使是临时存储。当传递一个结构体中字段的指针时，有问题的 Go 内存是被该字段占用的内存，而不是整个结构体。当传递一个数组或切片中元素的指针时，有问题的 Go 内存是整个数组或切片的整个后备数组。

C code may not keep a copy of a Go pointer after the call returns. This includes the _GoString_ type, which, as noted above, includes a Go pointer; _GoString_ values may not be retained by C code.

​	C 代码在调用返回后不能保留 Go 指针的副本。这包括 _GoString_ 类型，如上所述，它包含一个 Go 指针；C 代码不能保留 _GoString_ 值。

A Go function called by C code may not return a Go pointer (which implies that it may not return a string, slice, channel, and so forth). A Go function called by C code may take C pointers as arguments, and it may store non-pointer or C pointer data through those pointers, but it may not store a Go pointer in memory pointed to by a C pointer. A Go function called by C code may take a Go pointer as an argument, but it must preserve the property that the Go memory to which it points does not contain any Go pointers.

​	由 C 代码调用的 Go 函数不能返回 Go 指针（这意味着它不能返回字符串、切片、通道等）。由 C 代码调用的 Go 函数可以将 C 指针作为参数，并且可以通过这些指针存储非指针或 C 指针数据，但它不能将 Go 指针存储在 C 指针指向的内存中。由 C 代码调用的 Go 函数可以将 Go 指针作为参数，但它必须保留其所指向的 Go 内存不包含任何 Go 指针的属性。

Go code may not store a Go pointer in C memory. C code may store Go pointers in C memory, subject to the rule above: it must stop storing the Go pointer when the C function returns.

​	Go 代码不能将 Go 指针存储在 C 内存中。C 代码可以将 Go 指针存储在 C 内存中，但须遵守上述规则：当 C 函数返回时，它必须停止存储 Go 指针。

These rules are checked dynamically at runtime. The checking is controlled by the cgocheck setting of the GODEBUG environment variable. The default setting is GODEBUG=cgocheck=1, which implements reasonably cheap dynamic checks. These checks may be disabled entirely using GODEBUG=cgocheck=0. Complete checking of pointer handling, at some cost in run time, is available via GODEBUG=cgocheck=2.

​	这些规则在运行时动态检查。检查由 GODEBUG 环境变量的 cgocheck 设置控制。默认设置是 GODEBUG=cgocheck=1，它实现了相当便宜的动态检查。可以使用 GODEBUG=cgocheck=0 完全禁用这些检查。通过 GODEBUG=cgocheck=2，可以以一定的运行时成本对指针处理进行完全检查。

It is possible to defeat this enforcement by using the unsafe package, and of course there is nothing stopping the C code from doing anything it likes. However, programs that break these rules are likely to fail in unexpected and unpredictable ways.

​	通过使用 unsafe 包可以击败此强制执行，当然，没有任何东西可以阻止 C 代码执行任何它喜欢的事情。但是，违反这些规则的程序可能会以意外和不可预测的方式失败。

The runtime/cgo.Handle type can be used to safely pass Go values between Go and C. See the runtime/cgo package documentation for details.

​	runtime/cgo.Handle 类型可用于在 Go 和 C 之间安全地传递 Go 值。有关详细信息，请参阅 runtime/cgo 包文档。

Note: the current implementation has a bug. While Go code is permitted to write nil or a C pointer (but not a Go pointer) to C memory, the current implementation may sometimes cause a runtime error if the contents of the C memory appear to be a Go pointer. Therefore, avoid passing uninitialized C memory to Go code if the Go code is going to store pointer values in it. Zero out the memory in C before passing it to Go.

​	注意：当前实现存在一个错误。虽然允许 Go 代码向 C 内存写入 nil 或 C 指针（但不允许写入 Go 指针），但如果 C 内存的内容看起来像 Go 指针，当前实现有时可能会导致运行时错误。因此，如果 Go 代码要在其中存储指针值，请避免将未初始化的 C 内存传递给 Go 代码。在将 C 内存传递给 Go 之前，先将其清零。

#### Special cases

A few special C types which would normally be represented by a pointer type in Go are instead represented by a uintptr. Those include:

​	一些通常在 Go 中由指针类型表示的特殊 C 类型改用 uintptr 表示。这些类型包括：

1. The *Ref types on Darwin, rooted at CoreFoundation's CFTypeRef type.


​	1. Darwin 上的 *Ref 类型，根植于 CoreFoundation 的 CFTypeRef 类型。

2. The object types from Java's JNI interface:


​	2. Java 的 JNI 接口中的对象类型：

```
jobject
jclass
jthrowable
jstring
jarray
jbooleanArray
jbyteArray
jcharArray
jshortArray
jintArray
jlongArray
jfloatArray
jdoubleArray
jobjectArray
jweak
```

3. The EGLDisplay and EGLConfig types from the EGL API.


​	3. EGL API 中的 EGLDisplay 和 EGLConfig 类型。

These types are uintptr on the Go side because they would otherwise confuse the Go garbage collector; they are sometimes not really pointers but data structures encoded in a pointer type. All operations on these types must happen in C. The proper constant to initialize an empty such reference is 0, not nil.

​	这些类型在 Go 侧是 uintptr，因为它们会混淆 Go 垃圾回收器；它们有时并不是真正的指针，而是以指针类型编码的数据结构。对这些类型的任何操作都必须在 C 中进行。初始化此类空引用的正确常量是 0，而不是 nil。

These special cases were introduced in Go 1.10. For auto-updating code from Go 1.9 and earlier, use the cftype or jni rewrites in the Go fix tool:

​	这些特殊情况在 Go 1.10 中引入。对于从 Go 1.9 及更早版本自动更新代码，请在 Go fix 工具中使用 cftype 或 jni 重写：

```
go tool fix -r cftype <pkg>
go tool fix -r jni <pkg>
```

It will replace nil with 0 in the appropriate places.

​	它将在适当的位置用 0 替换 nil。

The EGLDisplay case was introduced in Go 1.12. Use the egl rewrite to auto-update code from Go 1.11 and earlier:

​	EGLDisplay 情况在 Go 1.12 中引入。使用 egl 重写来自 Go 1.11 及更早版本的自动更新代码：

```
go tool fix -r egl <pkg>
```

The EGLConfig case was introduced in Go 1.15. Use the eglconf rewrite to auto-update code from Go 1.14 and earlier:

​	EGLConfig 情况在 Go 1.15 中引入。使用 eglconf 重写来自 Go 1.14 及更早版本的自动更新代码：

```
go tool fix -r eglconf <pkg>
```

#### Using cgo directly

Usage:

​	用法：

```
go tool cgo [cgo options] [-- compiler options] gofiles...
```

Cgo transforms the specified input Go source files into several output Go and C source files.

​	Cgo 将指定的输入 Go 源文件转换为多个输出 Go 和 C 源文件。

The compiler options are passed through uninterpreted when invoking the C compiler to compile the C parts of the package.

​	在调用 C 编译器编译包的 C 部分时，编译器选项会原样传递，不会进行解释。

The following options are available when running cgo directly:

​	直接运行 cgo 时，可以使用以下选项：

```
-V
	Print cgo version and exit.
-debug-define
	Debugging option. Print #defines.
-debug-gcc
	Debugging option. Trace C compiler execution and output.
-dynimport file
	Write list of symbols imported by file. Write to
	-dynout argument or to standard output. Used by go
	build when building a cgo package.
-dynlinker
	Write dynamic linker as part of -dynimport output.
-dynout file
	Write -dynimport output to file.
-dynpackage package
	Set Go package for -dynimport output.
-exportheader file
	If there are any exported functions, write the
	generated export declarations to file.
	C code can #include this to see the declarations.
-importpath string
	The import path for the Go package. Optional; used for
	nicer comments in the generated files.
-import_runtime_cgo
	If set (which it is by default) import runtime/cgo in
	generated output.
-import_syscall
	If set (which it is by default) import syscall in
	generated output.
-gccgo
	Generate output for the gccgo compiler rather than the
	gc compiler.
-gccgoprefix prefix
	The -fgo-prefix option to be used with gccgo.
-gccgopkgpath path
	The -fgo-pkgpath option to be used with gccgo.
-godefs
	Write out input file in Go syntax replacing C package
	names with real values. Used to generate files in the
	syscall package when bootstrapping a new target.
-objdir directory
	Put all generated files in directory.
-srcdir directory
```