+++
title = "go help filetype"
date = 2023-12-12T14:13:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

### extensions 

The go command examines the contents of a restricted set of files in each directory. It identifies which files to examine based on the extension of the file name. These extensions are:

### .go

​        Go source files.

### .c, .h

​        C source files. If the package uses cgo or SWIG, these will be compiled with the OS-native compiler (typically gcc); otherwise they will trigger an error.

### .cc, .cpp, .cxx, .hh, .hpp, .hxx

​        C++ source files. Only useful with cgo or SWIG, and always compiled with the OS-native compiler.

### .m

​        Objective-C source files. Only useful with cgo, and always compiled with the OS-native compiler.

### .s, .S, .sx

​        Assembler source files. If the package uses cgo or SWIG, these will be assembled with the  OS-native assembler (typically gcc (sic)); otherwise they  will be assembled with the Go assembler.

### .swig, .swigcxx

​        SWIG definition files.

### .syso

​        System object files.

Files of each of these types except .syso may contain build constraints, but the go command stops scanning for build constraints
at the first item in the file that is not a blank line or //-style line comment. See the go/build package documentation for more details.
