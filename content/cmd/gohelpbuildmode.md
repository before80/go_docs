+++
title = "go help buildmode"
date = 2023-12-12T14:13:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

​	

The '`go build`' and '`go install`' commands take a `-buildmode` argument which indicates which kind of object file is to be built. Currently supported values are:

        -buildmode=archive
                Build the listed non-main packages into .a files. Packages named
                main are ignored.
    
        -buildmode=c-archive
                Build the listed main package, plus all packages it imports,
                into a C archive file. The only callable symbols will be those
                functions exported using a cgo //export comment. Requires
                exactly one main package to be listed.
    
        -buildmode=c-shared
                Build the listed main package, plus all packages it imports,
                into a C shared library. The only callable symbols will
                be those functions exported using a cgo //export comment.
                Requires exactly one main package to be listed.
    
        -buildmode=default
                Listed main packages are built into executables and listed
                non-main packages are built into .a files (the default
                behavior).
    
        -buildmode=shared
                Combine all the listed non-main packages into a single shared
                library that will be used when building with the -linkshared
                option. Packages named main are ignored.
    
        -buildmode=exe
                Build the listed main packages and everything they import into
                executables. Packages not named main are ignored.
    
        -buildmode=pie
                Build the listed main packages and everything they import into
                position independent executables (PIE). Packages not named
                main are ignored.
    
        -buildmode=plugin
                Build the listed main packages, plus all packages that they
                import, into a Go plugin. Packages not named main are ignored.

On AIX, when linking a C program that uses a Go archive built with `-buildmode=c-archive`, you must pass `-Wl`,`-bnoobjreorder` to the C compiler.
