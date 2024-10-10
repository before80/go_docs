+++
title = "Setting up and using gccgo"
weight = 4
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Setting up and using gccgo 

> 原文：[https://go.dev/doc/install/gccgo](https://go.dev/doc/install/gccgo)

​	本文介绍了如何使用`gccgo`，这是Go语言的编译器。`gccgo`编译器是广泛使用的GNU编译器`GCC`的一个新前端。尽管前端本身采用BSD风格的许可证，但`gccgo`通常作为`GCC`的一部分使用，然后由[GNU通用公共许可证](https://www.gnu.org/licenses/gpl.html)涵盖（该许可证涵盖`gccgo`本身作为`GCC`的一部分；它不包括由`gccgo`生成的代码）。

> 请注意，gccgo 不是 `gc` 编译器；请参阅[安装 Go](../InstallingGo) 编译器的说明。

## 版本

​	安装`gccgo`的最简单方法是安装包含Go支持的`GCC`二进制版本。`GCC` 二进制版本可以从[不同的网站](https://gcc.gnu.org/install/binaries.html)上获得，通常作为 `GNU/Linux` 发行版的一部分。我们希望大多数构建这些二进制文件的人都会加入Go支持。



`GCC 4.7.1`版本和所有后来的4.7版本包含完整的[Go 1](https://go.dev/doc/go1.html)编译器和库。



由于时间关系，`GCC 4.8.0` 和 `4.8.1` 版本接近于 `Go 1.1`，但不完全相同。`GCC 4.8.2`版本包含一个完整的`Go 1.1.2`实现。



`GCC 4.9`版本包含一个完整的`Go 1.2`实现。



`GCC 5` 版本包含 `Go 1.4` 用户库的完整实现。`Go 1.4` 的运行时没有完全合并，但这对 Go 程序来说不应该是可见的。



`GCC 6` 版本包含 `Go 1.6.1` 用户库的完整实现。`Go 1.6` 运行时没有完全合并，但这对 Go 程序来说不应该是可见的。



`GCC 7` 版本包含 `Go 1.8.1` 用户库的完整实现。与早期版本一样，`Go 1.8` 的运行时没有完全合并，但这对 Go 程序来说不应该是可见的。



`GCC 8` 版本包含对 `Go 1.10.1` 版本的完整实现。`Go 1.10` 运行时现在已经完全合并到 `GCC` 开发源中，并且完全支持并发垃圾收集。



`GCC 9` 版本包含对 `Go 1.12.2` 版本的完整实现。



`GCC 10`版本包含对`Go 1.14.6`版本的完整实现。



`GCC 11`版本包含对`Go 1.16.3`版本的完整实现。

## 源代码

​	如果您不能使用某个发行版，或者喜欢自己构建`gccgo`，可以通过Git访问`gccgo`源代码。`GCC` 网站上有[获取 `GCC` 源代码的说明](https://gcc.gnu.org/git.html)。`gccgo`的源代码已经包括在内。为方便起见，在`GCC`主代码库的`devel/gccgo`分支中提供了Go支持的稳定版本：`git://gcc.gnu.org/git/gcc.git`。这个分支会定期更新稳定的 Go 编译器源代码。

> 请注意，尽管 `gcc.gnu.org` 是获取 Go 前端源代码的最便捷方式，但它并不是主源的所在。如果您想对 Go 前端编译器进行修改，请参见[对 gccgo 的贡献](https://go.dev/doc/gccgo_contribute.html)。

## 构建

​	构建`gccgo`就像构建`GCC`一样，只是多了一个或两个选项。见[gcc网站上的说明](https://gcc.gnu.org/install/)。当您运行`configure`时，添加选项`--enable-languages=c,c++,go`（以及其他您可能想要构建的语言）。如果您的目标是`32`位`x86`，若您希望所构建的`gccgo`默认支持锁定的比较和交换指令；那么可以通过使用`configure`选项`--with-arch=i586`（或更新的架构，取决于您需要您的程序在哪里运行）来实现。如果您的目标是`64`位`x86`，但有时想使用`-m32`选项，那么可以使用`configure`选项 `--with-arch-32=i586`。

### Gold

​	在x86 GNU/Linux系统上，gccgo编译器能够为`goroutines`使用一个小的不连续堆栈。这允许程序运行更多的`goroutine`，因为每个`goroutine`可以使用一个相对较小的堆栈。做到这一点需要使用2.22或更高版本的`gold linker`。您可以安装`GNU binutils 2.22`或更高版本，或者您可以自己构建`gold`。

​	要自己构建`gold`，可以在运行`configure`脚本时使用`--enable-gold=default`来构建`GNU binutils`。在构建之前，您必须安装`flex`和`bison`软件包。一个典型的顺序是这样的（您可以把`/opt/gold`替换成任何您有写权限的目录）：

```shell
git clone git://sourceware.org/git/binutils-gdb.git
mkdir binutils-objdir
cd binutils-objdir
../binutils-gdb/configure --enable-gold=default --prefix=/opt/gold
make
make install
```

​	无论您如何安装`gold`，当您配置`gccgo`时，请使用选项`--with-ld=GOLD_BINARY`。

### 先决条件

​	如[gcc网站](https://gcc.gnu.org/install/prerequisites.html)上所述，构建`GCC`需要许多先决条件。在运行gcc `configure`脚本之前，安装所有先决条件是很重要的。先决条件库可以通过`GCC`源代码中的脚本`contrib/download_prerequisites`方便地下载。

### 构建命令

​	一旦所有的先决条件都安装好了，那么一个典型的构建和安装顺序是这样的（如果您使用上述的`gold linker`，只需使用`--with-ld`选项）：

```
git clone --branch devel/gccgo git://gcc.gnu.org/git/gcc.git gccgo
mkdir objdir
cd objdir
../gccgo/configure --prefix=/opt/gccgo --enable-languages=c,c++,go --with-ld=/opt/gold/bin/ld
make
make install
```

## 使用 gccgo

​	`gccgo`编译器的工作方式与其他`gcc`前端一样。从`GCC 5`开始，`gccgo`的安装也包括一个`go`命令的版本，它可以用来编译Go程序，如https://go.dev/cmd/go 中所描述的。

不使用 go 命令编译文件：

```
gccgo -c file.go
```

这就产生了`file.o`。要把文件链接起来形成一个可执行文件：

```
gccgo -o file file.o
```

要运行生成的文件，您需要告诉程序在哪里可以找到编译后的Go软件包。有几种方法可以做到这一点：

- 设置`LD_LIBRARY_PATH`环境变量：

  ```
  LD_LIBRARY_PATH=${prefix}/lib/gcc/MACHINE/VERSION
  [or]
  LD_LIBRARY_PATH=${prefix}/lib64/gcc/MACHINE/VERSION
  export LD_LIBRARY_PATH
  ```

  Here `${prefix}` is the `--prefix` option used when building gccgo. For a binary install this is normally `/usr`. Whether to use `lib` or `lib64` depends on the target. Typically `lib64` is correct for x86_64 systems, and `lib` is correct for other systems. The idea is to name the directory where `libgo.so` is found.

  这里的`${prefix}`是构建`gccgo`时使用的`--prefix`选项。对于二进制安装来说，这通常是`/usr`。是否使用`lib`或`lib64`取决于目标。通常，`lib64`适用于`x86_64`系统，`lib`适用于其他系统。想法是找到 `libgo.so` 的所在目录作为其名称。=> 仍有疑问？？

- 在链接时传递一个 `-Wl,-R` 选项（如果适合您的系统，就用 `lib64` 替换 `lib`）：

  ```
  go build -gccgoflags -Wl,-R,${prefix}/lib/gcc/MACHINE/VERSION
  [or]
  gccgo -o file file.o -Wl,-R,${prefix}/lib/gcc/MACHINE/VERSION
  ```

- 使用`-static-libgo`选项来静态链接已编译的软件包。

- 使用`-static`选项可以进行完全静态链接（`gc`编译器的默认值）。


## 选项

​	`gccgo`编译器支持所有与语言无关的`GCC`选项，特别是`-O`和`-g`选项。

​	`-fgo-pkgpath=PKGPATH`选项可以用来为正在编译的软件包设置一个唯一的前缀。此选项会被`go`命令自动使用，但如果您直接调用`gccgo`，您可能需要使用它。此选项主要用于包含许多包的大型程序，以允许多个包使用相同的标识符作为包名。`PKGPATH`可以是任何字符串；字符串的一个好选择是用于导入软件包的路径。

​	`-I`和`-L`选项是编译器的同义词，可以用来设置寻找导入的搜索路径。如果您用`go`命令编译，则不需要这些选项。

## 导入

​	当您编译一个导出某些内容的文件时，导出信息将直接存储在对象文件中。如果您直接用`gccgo`编译，而不是用go命令，那么当您导入一个包时，您必须告诉`gccgo`如何找到该文件。

​	当您用`gccgo`导入`FILE`软件包时，它将在下列文件中寻找导入数据，并使用它找到的第一个文件。

- `FILE.gox`
- `libFILE.so`
- `libFILE.a`
- `FILE.o`

`FILE.gox`，在使用时，通常只包含导出数据。这可以通过以下方式从`FILE.o`生成：

```
objcopy -j .go_export FILE.o FILE.gox
```

​	`gccgo`编译器会在当前目录下寻找导入文件。在更复杂的情况下，您可以传递`-I`或`-L`选项给`gccgo`。这两个选项都需要设置为搜索的目录。`-L`选项也会传递给链接器。

​	`gccgo`编译器目前（2015-06-15）不在对象文件中记录导入包的文件名。您必须安排将导入的数据链接到程序中。同样，在用`go`命令构建时，这也是没有必要的。

```
gccgo -c mypackage.go              # Exports mypackage
gccgo -c main.go                   # Imports mypackage
gccgo -o main main.o mypackage.o   # Explicitly links with mypackage.o
```

## 调试

​	如果您在编译时使用`-g`选项，您可以在您的可执行文件上运行`gdb`。调试器对Go的了解是有限的。您可以设置断点、单步执行等。您可以打印变量，但它们会被打印成具有`C/C++`的类型。对于数字类型，这并不重要。Go 字符串和接口将显示为两个元素结构。Go映射和通道始终表示为指向运行时结构的C指针。

## C语言的互操作性

​	在使用 `gccgo` 时，与 C 或使用 `extern "C"` 编译的 C++ 代码的互操作性是有限的。

### 类型

​	基本类型直接映射关系是：Go中的`int32`是C中的`int32_t`，`int64`是`int64_t`，等等。Go中的`int`类型是一个整数，与指针的大小相同，因此对应于C中的`intptr_t`。Go的`byte`相当于C的无符号`char`。Go中的指针就是C中的指针。Go结构与C结构相同，具有相同的字段和类型。

​	Go`字符串`类型目前（在这里）被定义为一个双元素结构（这可能会被改变）：

```
struct __go_string {
  const unsigned char *__data;
  intptr_t __length;
};
```

​	您不能在`C`和`Go`之间传递数组。然而，Go中指向数组的指针等同于指向元素类型的C语言指针。例如，Go `*[10]int `相当于 C `int*`，前提是 C 指针确实指向 10 个元素。

Go中的切片（在这里）是一个结构。目前的定义是（这可能会改变）：

```
struct __go_slice {
  void *__values;
  intptr_t __count;
  intptr_t __capacity;
};
```

​	Go函数的类型（在这里）是一个指向结构体的指针（这可能会有变化）。结构中的第一个字段指向函数的代码，这将相当于一个指向C函数的指针，其参数类型是相等的，另外还有一个尾随参数。后面的参数是闭包，要传递的参数是指向Go函数结构的指针。当Go函数返回一个以上的值时，C函数返回一个结构。例如，这些函数大致上是等价的：

``` go
func GoFunction(int) (int, float64)
struct { int i; float64 f; } CFunction(int, void*)
```

​	Go的`interface`、`channel`和`map`类型没有对应的C类型（接口是一个双元素结构体，`channel`和`map`是指向C中的结构体的指针，但结构体是`故意不记录的（deliberately undocumented）`）。C语言的`enum`类型对应于某种整数类型，但在一般情况下是难以预测的整数类型；请使用强制转换（ cast）。C语言的`union`类型没有对应的Go类型。包含`位字段（bitfields）`的 C 结构类型没有对应的 Go 类型。C++的`class`类型没有对应的Go类型。

​	C和Go的`内存分配（memory allocation）`是完全不同的，因为Go使用垃圾收集。这方面的确切准则尚未确定，但很可能允许将分配的内存的指针从C语言传递到Go语言。最终释放指针的责任仍由C端承担，当然，如果C端释放了指针，而Go端仍有一个副本，程序就会失败。当把指针从Go传到C时，Go函数必须在某个Go变量中保留一个可见的副本。否则，Go的垃圾收集器可能会在C函数仍在使用该指针时删除它。

### 函数名称

​	Go代码可以使用在`gccgo`中实现的Go扩展来直接调用C函数：函数声明前面可以加上`//extern NAME`。例如，以下是C函数`open`在Go中的声明方式：

```
//extern open
func c_open(name *byte, mode int, perm int) int
```

​	C函数自然需要得到一个以`NUL`结尾的字符串，在Go中，这相当于一个指向数组（不是切片！）的指针，其中有一个结束的零字节。因此，Go中的一个示例调用看起来像（在导入`syscall`包之后）：

```
var name = [4]byte{'f', 'o', 'o', 0};
i := c_open(&name[0], syscall.O_RDONLY, 0);
```

(这只是一个例子，要在Go中打开一个文件，请使用Go的`os.Open`函数代替）。

> 注意
>
> ​	如果C函数可以阻塞，例如在调用`read` 时，调用C函数可能会阻塞Go程序。除非您清楚地了解您在做什么，否则所有C和Go之间的调用都应该通过`cgo`或`SWIG`来实现，就像`gc`编译器一样。

​	从C语言访问的Go函数的名称可能会改变。目前，没有接收器的 Go 函数的名称是 `prefix.package.Functionname`。前缀由编译包时使用的`-fgo-prefix`选项设置；如果没有使用该选项，默认为`go`。要从C语言中调用该函数，您必须使用`GCC`扩展来设置名称。

```
extern int go_function(int) __asm__ ("myprefix.mypackage.Function");
```

### 从C源代码中自动生成Go声明

​	`GCC`的Go版本支持从C代码中自动生成Go声明。这个功能比较笨拙，大多数用户应该使用带有`-gccgo`选项的`cgo`程序来代替。

​	像往常一样编译您的C代码，并添加选项`-fdump-go-spec=FILENAME`。这将创建 `FILENAME` 文件作为编译的副作用。这个文件将包含C代码中所声明的类型、变量和函数的Go声明。不能用Go表示的C类型将被记录为Go代码中的注释。生成的文件将没有`package`声明，但可以直接被`gccgo`编译。

​	这个程序充满了未说明的注意事项和限制，我们不保证它在将来不会改变。与其说它是一个常规程序，不如说它是真正Go代码的起点。