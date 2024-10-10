+++
title = "包"
date = 2023-05-17T09:59:21+08:00
weight = 14
description = ""
isCJKLanguage = true
type = "docs"
draft = false
+++
## Packages 包

> 原文：[https://go.dev/ref/spec#Packages ](https://go.dev/ref/spec#Packages )

Go programs are constructed by linking together *packages*. A package in turn is constructed from one or more source files that together declare constants, types, variables and functions belonging to the package and which are accessible in all files of the same package. Those elements may be [exported](https://go.dev/ref/spec#Exported_identifiers) and used in another package.

​	Go程序是通过链接包来构建的。一个包是由一个或多个源文件构成的，这些文件共同声明了属于该包的常量、类型、变量和函数，并且可以在同一包的所有文件中访问。这些元素可以[导出](../DeclarationsAndScope#exported-identifiers-可导出的标识符)并在另一个包中使用。

### Source file organization 源文件组织

Each source file consists of a package clause defining the package to which it belongs, followed by a possibly empty set of import declarations that declare packages whose contents it wishes to use, followed by a possibly empty set of declarations of functions, types, variables, and constants.

​	每个源文件都由一个 package 子句（package clause）组成，该子句定义了它所属的包，然后是一组可能为空的导入声明，这些导入声明声明了它希望使用的包的内容，然后是一组可能为空的函数、类型、变量和常量的声明。

```
SourceFile       = PackageClause ";" { ImportDecl ";" } { TopLevelDecl ";" } .
```

### Package clause 包子句

A package clause begins each source file and defines the package to which the file belongs.

​	每个源文件的开头都有一个包子句，定义了该文件所属的包。

``` go
packageClause  = "package" PackageName .
PackageName    = identifier .
```

The PackageName must not be the [blank identifier](https://go.dev/ref/spec#Blank_identifier).

​	`PackageName（包名）`不能是[空白标识符（即`_`）](../DeclarationsAndScope#blank-identifier-空白标识符)。

```go 
package math
```

A set of files sharing the same PackageName form the implementation of a package. An implementation may require that all source files for a package inhabit the same directory.

​	共享相同`PackageName（包名）`的一组文件构成了一个包的实现。实现可能会要求一个包的所有源文件存放在同一个目录中。

### Import declarations 导入声明

An import declaration states that the source file containing the declaration depends on functionality of the *imported* package ([§Program initialization and execution](https://go.dev/ref/spec#Program_initialization_and_execution)) and enables access to [exported](https://go.dev/ref/spec#Exported_identifiers) identifiers of that package. The import names an identifier (PackageName) to be used for access and an ImportPath that specifies the package to be imported.

​	导入声明指出，包含该声明的源文件依赖于导入包的功能（[§程序初始化和执行](../ProgramInitializationAndExecution)），并且能够访问这些被导入的包的[导出](../DeclarationsAndScope#exported-identifiers-可导出的标识符)标识符。该导入命名了一个用于访问的标识符（PackageName 包名）和一个指定要导入的包的ImportPath（导入路径）。

```
ImportDecl       = "import" ( ImportSpec | "(" { ImportSpec ";" } ")" ) .
ImportSpec       = [ "." | PackageName ] ImportPath .
ImportPath       = string_lit .
```

The PackageName is used in [qualified identifiers](https://go.dev/ref/spec#Qualified_identifiers) to access exported identifiers of the package within the importing source file. It is declared in the [file block](https://go.dev/ref/spec#Blocks). If the PackageName is omitted, it defaults to the identifier specified in the [package clause](https://go.dev/ref/spec#Package_clause) of the imported package. If an explicit period (`.`) appears instead of a name, all the package's exported identifiers declared in that package's [package block](https://go.dev/ref/spec#Blocks) will be declared in the importing source file's file block and must be accessed without a qualifier.

​	PackageName （包名）用于[限定标识符](../Expressions#qualified-identifiers-限定标识符)，以访问导入源文件中包的导出标识符。它被声明在[文件块](../Blocks)中。如果（导入声明中）PackageName（包名）被省略，则它默认为在导入包的[包子句](#package-clause-包子句)中指定的标识符。如果出现一个明确的句号（`.`）而不是名字，那么在该包的[包块](../Blocks)中声明的所有包的导出标识符将在导入源文件的文件块中声明，并且必须在没有限定符的情况下进行访问。

The interpretation of the ImportPath is implementation-dependent but it is typically a substring of the full file name of the compiled package and may be relative to a repository of installed packages.

​	对ImportPath（导入路径）的解释是依赖于实现的，但它通常是已编译包的完整文件名的子字符串，并且可能是相对于已安装包的存储库的。

Implementation restriction: A compiler may restrict ImportPaths to non-empty strings using only characters belonging to [Unicode's](https://www.unicode.org/versions/Unicode6.3.0/) L, M, N, P, and S general categories (the Graphic characters without spaces) and may also exclude the characters `!"#$%&'()*,:;<=>?[\]^‘{|}` and the Unicode replacement character U+FFFD.

​	实现限制：编译器可以将ImportPaths（导入路径）限制为非空字符串，只使用属于[Unicode的](https://www.unicode.org/versions/Unicode6.3.0/)L、M、N、P和S一般类别的字符（不含空格的图形字符），也可能排除字符`!"#$%&'()*,:;<=>?[\]^‘{|}`和Unicode替换字符U+FFFD。

Consider a compiled a package containing the package clause `package math`, which exports function `Sin`, and installed the compiled package in the file identified by `"lib/math"`. This table illustrates how `Sin` is accessed in files that import the package after the various types of import declaration.

​	考虑一个包含包子句`package math`的编译包，它导出了函数`Sin`，并将编译包安装在由 "`lib/math`"标识的文件中。下表说明了在各种类型的导入声明之后，在导入包的文件中如何访问Sin。

```
Import declaration          Local name of Sin
导入声明                       本地 Sin 名

import   "lib/math"         math.Sin
import m "lib/math"         m.Sin
import . "lib/math"         Sin
```

An import declaration declares a dependency relation between the importing and imported package. It is illegal for a package to import itself, directly or indirectly, or to directly import a package without referring to any of its exported identifiers. To import a package solely for its side-effects (initialization), use the [blank](https://go.dev/ref/spec#Blank_identifier) identifier as explicit package name:

​	导入声明声明了导入包和被导入包之间的依赖关系。如果一个包直接或间接地导入自身，或者直接导入一个包而不引用其任何导出的标识符，都是非法的。要导入一个包只为了它的副作用（初始化），可以使用[空白标识符（即`_`）](../DeclarationsAndScope#blank-identifier-空白标识符)作为显式的包名：

```
import _ "lib/math"
```

### An example package 一个包的例子

Here is a complete Go package that implements a concurrent prime sieve.

​	下面是一个完整的Go包，它实现了一个并发素数筛选。

```go 
package main

import "fmt"

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i  // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {  // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i  // Send 'i' to channel 'dst'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int)  // Create a new channel.
	go generate(ch)       // Start generate() as a subprocess.
	for {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve()
}
```

