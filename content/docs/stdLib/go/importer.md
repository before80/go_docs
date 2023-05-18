+++
title = "importer"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# importer

https://pkg.go.dev/go/importer@go1.20.1



Package importer provides access to export data importers.



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [Default](https://cs.opensource.google/go/go/+/go1.20.1:src/go/importer/importer.go;l=82) 

``` go 
func Default() types.Importer
```

Default returns an Importer for the compiler that built the running binary. If available, the result implements types.ImporterFrom.

##### Example
``` go 
```

#### func [ForCompiler](https://cs.opensource.google/go/go/+/go1.20.1:src/go/importer/importer.go;l=40)  <- go1.12

``` go 
func ForCompiler(fset *token.FileSet, compiler string, lookup Lookup) types.Importer
```

ForCompiler returns an Importer for importing from installed packages for the compilers "gc" and "gccgo", or for importing directly from the source if the compiler argument is "source". In this latter case, importing may fail under circumstances where the exported API is not entirely defined in pure Go source code (if the package API depends on cgo-defined entities, the type checker won't have access to those).

The lookup function is called each time the resulting importer needs to resolve an import path. In this mode the importer can only be invoked with canonical import paths (not relative or absolute ones); it is assumed that the translation to canonical import paths is being done by the client of the importer.

A lookup function must be provided for correct module-aware operation. Deprecated: If lookup is nil, for backwards-compatibility, the importer will attempt to resolve imports in the $GOPATH workspace.

## 类型

### type [Lookup](https://cs.opensource.google/go/go/+/go1.20.1:src/go/importer/importer.go;l=21) 

``` go 
type Lookup func(path string) (io.ReadCloser, error)
```

A Lookup function returns a reader to access package data for a given import path, or an error if no matching package is found.