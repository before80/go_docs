+++
title = "importer"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/go/importer@go1.23.0](https://pkg.go.dev/go/importer@go1.23.0)

Package importer provides access to export data importers.

​	importer 包提供对导出数据导入器的访问。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

### func Default

```go
func Default() types.Importer
```

Default returns an Importer for the compiler that built the running binary. If available, the result implements types.ImporterFrom.

​	Default 返回一个 Importer，用于构建正在运行的二进制文件的编译器。如果可用，结果将实现 types.ImporterFrom。

### func For <-DEPRECATED

```go
func For(compiler string, lookup Lookup) types.Importer
```

For calls ForCompiler with a new FileSet.

​	For 调用 ForCompiler 并使用新的 FileSet。

Deprecated: Use ForCompiler, which populates a FileSet with the positions of objects created by the importer.

​	已弃用：使用 ForCompiler，它使用导入器创建的对象的位置填充 FileSet。

### func ForCompiler <- go1.12

```go
func ForCompiler(fset *token.FileSet, compiler string, lookup Lookup) types.Importer
```

ForCompiler returns an Importer for importing from installed packages for the compilers “gc” and “gccgo”, or for importing directly from the source if the compiler argument is “source”. In this latter case, importing may fail under circumstances where the exported API is not entirely defined in pure Go source code (if the package API depends on cgo-defined entities, the type checker won’t have access to those).

​	ForCompiler 返回一个 Importer，用于从编译器“gc”和“gccgo”的已安装包中导入，或者在编译器参数为“source”时直接从源代码中导入。在后一种情况下，如果导出的 API 并非完全在纯 Go 源代码中定义，则导入可能会失败（如果包 API 依赖于 cgo 定义的实体，类型检查器将无法访问这些实体）。

The lookup function is called each time the resulting importer needs to resolve an import path. In this mode the importer can only be invoked with canonical import paths (not relative or absolute ones); it is assumed that the translation to canonical import paths is being done by the client of the importer.

​	每次生成的导入器需要解析导入路径时，都会调用查找函数。在此模式下，导入器只能使用规范导入路径（不是相对或绝对路径）进行调用；假定导入器的客户端正在执行到规范导入路径的转换。

A lookup function must be provided for correct module-aware operation. Deprecated: If lookup is nil, for backwards-compatibility, the importer will attempt to resolve imports in the $GOPATH workspace.

​	必须提供查找函数才能正确进行模块感知操作。已弃用：如果查找为 nil，则出于向后兼容性的考虑，导入器将尝试解析 $GOPATH 工作区中的导入。

## 类型

### type Lookup

```go
type Lookup func(path string) (io.ReadCloser, error)
```

A Lookup function returns a reader to access package data for a given import path, or an error if no matching package is found.

​	Lookup 函数返回一个读取器，用于访问给定导入路径的包数据，或者在找不到匹配包时返回一个错误。