+++
title = "pack"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# pack

> 原文：[https://pkg.go.dev/cmd/pack@go1.19.3](https://pkg.go.dev/cmd/pack@go1.19.3)

### Overview 概述

​	`pack`是传统`Unix ar`工具的一个简单版本。它只实现Go所需的操作。

使用方法：

```
go tool pack op file.a [name...]
```

Pack applies the operation to the archive, using the names as arguments to the operation.

​	pack将操作应用于存档，使用名称（names ）作为操作的参数。

The operation op is given by one of these letters:

​	操作`op`由下列字母中之一给出：

```
c	append files (from the file system) to a new archive
	=> 将文件（来自文件系统）追加到一个新的存档中
	
p	print files from the archive
	=> 打印存档中的文件
	
r	append files (from the file system) to the archive
	=> 将文件（从文件系统）追加到存档中
	
t	list files from the archive
	=> 列出存档中的文件
	
x	extract files from the archive
	=> 从存档中提取文件
```

​	`c`命令的存档参数必须是不存在的，或者是一个有效的（在添加新条目之前会被清除的）存档文件。如果文件存在但不是归档文件，则会是一个错误。

​	对于`p`、`t`和`x`命令，在命令行中不列出名称将导致操作应用于存档中的所有文件。

​	与`Unix ar`不同的是，`r`操作总是追加到归档文件中，即使归档文件中已经存在一个给定名称的文件。在这种情况下，`pack`的`r`操作更像`Unix ar`的`rq`操作。

​	在操作中添加字母`v`，如`pv`或`rv`，可以启用详细操作：

- 对于`c`和`r`命令，在添加文件的时候会打印文件名。

- 对于`p`命令，每个文件的前缀都是一行中单独的名称。

- 对于`t`命令，列表中包括额外的文件元数据。
- 对于`x`命令，在提取文件时打印名称。



=== "doc.go"

    ``` go 
    // Copyright 2014 The Go Authors. All rights reserved.
    // Use of this source code is governed by a BSD-style
    // license that can be found in the LICENSE file.
    
    /*
    Pack is a simple version of the traditional Unix ar tool.
    It implements only the operations needed by Go.
    
    Usage:
    
    	go tool pack op file.a [name...]
    
    Pack applies the operation to the archive, using the names as arguments to the operation.
    
    The operation op is given by one of these letters:
    
    	c	append files (from the file system) to a new archive
    	p	print files from the archive
    	r	append files (from the file system) to the archive
    	t	list files from the archive
    	x	extract files from the archive
    
    The archive argument to the c command must be non-existent or a
    valid archive file, which will be cleared before adding new entries. It
    is an error if the file exists but is not an archive.
    
    For the p, t, and x commands, listing no names on the command line
    causes the operation to apply to all files in the archive.
    
    In contrast to Unix ar, the r operation always appends to the archive,
    even if a file with the given name already exists in the archive. In this way
    pack's r operation is more like Unix ar's rq operation.
    
    Adding the letter v to an operation, as in pv or rv, enables verbose operation:
    For the c and r commands, names are printed as files are added.
    For the p command, each file is prefixed by the name on a line by itself.
    For the t command, the listing includes additional file metadata.
    For the x command, names are printed as files are extracted.
    */
    
    ```



=== "pack.go"

~~~go 
```go 
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "cmd/internal/archive"
    "fmt"
    "io"
    "io/fs"
    "log"
    "os"
    "path/filepath"
)

const usageMessage = `Usage: pack op file.a [name....]
Where op is one of cprtx optionally followed by v for verbose output.
For compatibility with old Go build environments the op string grc is
accepted as a synonym for c.

For more information, run
    go doc cmd/pack`

func usage() {
    fmt.Fprintln(os.Stderr, usageMessage)
    os.Exit(2)
}

func main() {
    log.SetFlags(0)
    log.SetPrefix("pack: ")
    // need "pack op archive" at least.
    if len(os.Args) < 3 {
        log.Print("not enough arguments")
        fmt.Fprintln(os.Stderr)
        usage()
    }
    setOp(os.Args[1])
    var ar *Archive
    switch op {
    case 'p':
        ar = openArchive(os.Args[2], os.O_RDONLY, os.Args[3:])
        ar.scan(ar.printContents)
    case 'r':
        ar = openArchive(os.Args[2], os.O_RDWR|os.O_CREATE, os.Args[3:])
        ar.addFiles()
    case 'c':
        ar = openArchive(os.Args[2], os.O_RDWR|os.O_TRUNC|os.O_CREATE, os.Args[3:])
        ar.addPkgdef()
        ar.addFiles()
    case 't':
        ar = openArchive(os.Args[2], os.O_RDONLY, os.Args[3:])
        ar.scan(ar.tableOfContents)
    case 'x':
        ar = openArchive(os.Args[2], os.O_RDONLY, os.Args[3:])
        ar.scan(ar.extractContents)
    default:
        log.Printf("invalid operation %q", os.Args[1])
        fmt.Fprintln(os.Stderr)
        usage()
    }
    if len(ar.files) > 0 {
        log.Fatalf("file %q not in archive", ar.files[0])
    }
}

// The unusual ancestry means the arguments are not Go-standard.
// These variables hold the decoded operation specified by the first argument.
// op holds the operation we are doing (prtx).
// verbose tells whether the 'v' option was specified.
var (
    op      rune
    verbose bool
)

// setOp parses the operation string (first argument).
func setOp(arg string) {
    // Recognize 'go tool pack grc' because that was the
    // formerly canonical way to build a new archive
    // from a set of input files. Accepting it keeps old
    // build systems working with both Go 1.2 and Go 1.3.
    if arg == "grc" {
        arg = "c"
    }

    for _, r := range arg {
        switch r {
        case 'c', 'p', 'r', 't', 'x':
            if op != 0 {
                // At most one can be set.
                usage()
            }
            op = r
        case 'v':
            if verbose {
                // Can be set only once.
                usage()
            }
            verbose = true
        default:
            usage()
        }
    }
}

const (
    arHeader = "!<arch>\n"
)

// An Archive represents an open archive file. It is always scanned sequentially
// from start to end, without backing up.
type Archive struct {
    a        *archive.Archive
    files    []string // Explicit list of files to be processed.
    pad      int      // Padding bytes required at end of current archive file
    matchAll bool     // match all files in archive
}

// archive opens (and if necessary creates) the named archive.
func openArchive(name string, mode int, files []string) *Archive {
    f, err := os.OpenFile(name, mode, 0666)
    if err != nil {
        log.Fatal(err)
    }
    var a *archive.Archive
    if mode&os.O_TRUNC != 0 { // the c command
        a, err = archive.New(f)
    } else {
        a, err = archive.Parse(f, verbose)
        if err != nil && mode&os.O_CREATE != 0 { // the r command
            a, err = archive.New(f)
        }
    }
    if err != nil {
        log.Fatal(err)
    }
    return &Archive{
        a:        a,
        files:    files,
        matchAll: len(files) == 0,
    }
}

// scan scans the archive and executes the specified action on each entry.
func (ar *Archive) scan(action func(*archive.Entry)) {
    for i := range ar.a.Entries {
        e := &ar.a.Entries[i]
        action(e)
    }
}

// listEntry prints to standard output a line describing the entry.
func listEntry(e *archive.Entry, verbose bool) {
    if verbose {
        fmt.Fprintf(stdout, "%s\n", e.String())
    } else {
        fmt.Fprintf(stdout, "%s\n", e.Name)
    }
}

// output copies the entry to the specified writer.
func (ar *Archive) output(e *archive.Entry, w io.Writer) {
    r := io.NewSectionReader(ar.a.File(), e.Offset, e.Size)
    n, err := io.Copy(w, r)
    if err != nil {
        log.Fatal(err)
    }
    if n != e.Size {
        log.Fatal("short file")
    }
}

// match reports whether the entry matches the argument list.
// If it does, it also drops the file from the to-be-processed list.
func (ar *Archive) match(e *archive.Entry) bool {
    if ar.matchAll {
        return true
    }
    for i, name := range ar.files {
        if e.Name == name {
            copy(ar.files[i:], ar.files[i+1:])
            ar.files = ar.files[:len(ar.files)-1]
            return true
        }
    }
    return false
}

// addFiles adds files to the archive. The archive is known to be
// sane and we are positioned at the end. No attempt is made
// to check for existing files.
func (ar *Archive) addFiles() {
    if len(ar.files) == 0 {
        usage()
    }
    for _, file := range ar.files {
        if verbose {
            fmt.Printf("%s\n", file)
        }

        f, err := os.Open(file)
        if err != nil {
            log.Fatal(err)
        }
        aro, err := archive.Parse(f, false)
        if err != nil || !isGoCompilerObjFile(aro) {
            f.Seek(0, io.SeekStart)
            ar.addFile(f)
            goto close
        }

        for _, e := range aro.Entries {
            if e.Type != archive.EntryGoObj || e.Name != "_go_.o" {
                continue
            }
            ar.a.AddEntry(archive.EntryGoObj, filepath.Base(file), 0, 0, 0, 0644, e.Size, io.NewSectionReader(f, e.Offset, e.Size))
        }
    close:
        f.Close()
    }
    ar.files = nil
}

// FileLike abstracts the few methods we need, so we can test without needing real files.
type FileLike interface {
    Name() string
    Stat() (fs.FileInfo, error)
    Read([]byte) (int, error)
    Close() error
}

// addFile adds a single file to the archive
func (ar *Archive) addFile(fd FileLike) {
    // Format the entry.
    // First, get its info.
    info, err := fd.Stat()
    if err != nil {
        log.Fatal(err)
    }
    // mtime, uid, gid are all zero so repeated builds produce identical output.
    mtime := int64(0)
    uid := 0
    gid := 0
    ar.a.AddEntry(archive.EntryNativeObj, info.Name(), mtime, uid, gid, info.Mode(), info.Size(), fd)
}

// addPkgdef adds the __.PKGDEF file to the archive, copied
// from the first Go object file on the file list, if any.
// The archive is known to be empty.
func (ar *Archive) addPkgdef() {
    done := false
    for _, file := range ar.files {
        f, err := os.Open(file)
        if err != nil {
            log.Fatal(err)
        }
        aro, err := archive.Parse(f, false)
        if err != nil || !isGoCompilerObjFile(aro) {
            goto close
        }

        for _, e := range aro.Entries {
            if e.Type != archive.EntryPkgDef {
                continue
            }
            if verbose {
                fmt.Printf("__.PKGDEF # %s\n", file)
            }
            ar.a.AddEntry(archive.EntryPkgDef, "__.PKGDEF", 0, 0, 0, 0644, e.Size, io.NewSectionReader(f, e.Offset, e.Size))
            done = true
        }
    close:
        f.Close()
        if done {
            break
        }
    }
}

// Finally, the actual commands. Each is an action.

// can be modified for testing.
var stdout io.Writer = os.Stdout

// printContents implements the 'p' command.
func (ar *Archive) printContents(e *archive.Entry) {
    ar.extractContents1(e, stdout)
}

// tableOfContents implements the 't' command.
func (ar *Archive) tableOfContents(e *archive.Entry) {
    if ar.match(e) {
        listEntry(e, verbose)
    }
}

// extractContents implements the 'x' command.
func (ar *Archive) extractContents(e *archive.Entry) {
    ar.extractContents1(e, nil)
}

func (ar *Archive) extractContents1(e *archive.Entry, out io.Writer) {
    if ar.match(e) {
        if verbose {
            listEntry(e, false)
        }
        if out == nil {
            f, err := os.OpenFile(e.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0444 /*e.Mode*/)
            if err != nil {
                log.Fatal(err)
            }
            defer f.Close()
            out = f
        }
        ar.output(e, out)
    }
}

// isGoCompilerObjFile reports whether file is an object file created
// by the Go compiler, which is an archive file with exactly one entry
// of __.PKGDEF, or _go_.o, or both entries.
func isGoCompilerObjFile(a *archive.Archive) bool {
    switch len(a.Entries) {
    case 1:
        return (a.Entries[0].Type == archive.EntryGoObj && a.Entries[0].Name == "_go_.o") ||
            (a.Entries[0].Type == archive.EntryPkgDef && a.Entries[0].Name == "__.PKGDEF")
    case 2:
        var foundPkgDef, foundGo bool
        for _, e := range a.Entries {
            if e.Type == archive.EntryPkgDef && e.Name == "__.PKGDEF" {
                foundPkgDef = true
            }
            if e.Type == archive.EntryGoObj && e.Name == "_go_.o" {
                foundGo = true
            }
        }
        return foundPkgDef && foundGo
    default:
        return false
    }
}
```
~~~

=== "pack_test.go"

    ```go 
    // Copyright 2014 The Go Authors. All rights reserved.
    // Use of this source code is governed by a BSD-style
    // license that can be found in the LICENSE file.
    
    package main
    
    import (
    	"bufio"
    	"bytes"
    	"cmd/internal/archive"
    	"fmt"
    	"internal/testenv"
    	"io"
    	"io/fs"
    	"os"
    	"os/exec"
    	"path/filepath"
    	"testing"
    	"time"
    )
    
    // testCreate creates an archive in the specified directory.
    func testCreate(t *testing.T, dir string) {
    	name := filepath.Join(dir, "pack.a")
    	ar := openArchive(name, os.O_RDWR|os.O_CREATE, nil)
    	// Add an entry by hand.
    	ar.addFile(helloFile.Reset())
    	ar.a.File().Close()
    	// Now check it.
    	ar = openArchive(name, os.O_RDONLY, []string{helloFile.name})
    	var buf bytes.Buffer
    	stdout = &buf
    	verbose = true
    	defer func() {
    		stdout = os.Stdout
    		verbose = false
    	}()
    	ar.scan(ar.printContents)
    	ar.a.File().Close()
    	result := buf.String()
    	// Expect verbose output plus file contents.
    	expect := fmt.Sprintf("%s\n%s", helloFile.name, helloFile.contents)
    	if result != expect {
    		t.Fatalf("expected %q got %q", expect, result)
    	}
    }
    
    // Test that we can create an archive, write to it, and get the same contents back.
    // Tests the rv and then the pv command on a new archive.
    func TestCreate(t *testing.T) {
    	dir := t.TempDir()
    	testCreate(t, dir)
    }
    
    // Test that we can create an archive twice with the same name (Issue 8369).
    func TestCreateTwice(t *testing.T) {
    	dir := t.TempDir()
    	testCreate(t, dir)
    	testCreate(t, dir)
    }
    
    // Test that we can create an archive, put some files in it, and get back a correct listing.
    // Tests the tv command.
    func TestTableOfContents(t *testing.T) {
    	dir := t.TempDir()
    	name := filepath.Join(dir, "pack.a")
    	ar := openArchive(name, os.O_RDWR|os.O_CREATE, nil)
    
    	// Add some entries by hand.
    	ar.addFile(helloFile.Reset())
    	ar.addFile(goodbyeFile.Reset())
    	ar.a.File().Close()
    
    	// Now print it.
    	var buf bytes.Buffer
    	stdout = &buf
    	verbose = true
    	defer func() {
    		stdout = os.Stdout
    		verbose = false
    	}()
    	ar = openArchive(name, os.O_RDONLY, nil)
    	ar.scan(ar.tableOfContents)
    	ar.a.File().Close()
    	result := buf.String()
    	// Expect verbose listing.
    	expect := fmt.Sprintf("%s\n%s\n", helloFile.Entry(), goodbyeFile.Entry())
    	if result != expect {
    		t.Fatalf("expected %q got %q", expect, result)
    	}
    
    	// Do it again without verbose.
    	verbose = false
    	buf.Reset()
    	ar = openArchive(name, os.O_RDONLY, nil)
    	ar.scan(ar.tableOfContents)
    	ar.a.File().Close()
    	result = buf.String()
    	// Expect non-verbose listing.
    	expect = fmt.Sprintf("%s\n%s\n", helloFile.name, goodbyeFile.name)
    	if result != expect {
    		t.Fatalf("expected %q got %q", expect, result)
    	}
    
    	// Do it again with file list arguments.
    	verbose = false
    	buf.Reset()
    	ar = openArchive(name, os.O_RDONLY, []string{helloFile.name})
    	ar.scan(ar.tableOfContents)
    	ar.a.File().Close()
    	result = buf.String()
    	// Expect only helloFile.
    	expect = fmt.Sprintf("%s\n", helloFile.name)
    	if result != expect {
    		t.Fatalf("expected %q got %q", expect, result)
    	}
    }
    
    // Test that we can create an archive, put some files in it, and get back a file.
    // Tests the x command.
    func TestExtract(t *testing.T) {
    	dir := t.TempDir()
    	name := filepath.Join(dir, "pack.a")
    	ar := openArchive(name, os.O_RDWR|os.O_CREATE, nil)
    	// Add some entries by hand.
    	ar.addFile(helloFile.Reset())
    	ar.addFile(goodbyeFile.Reset())
    	ar.a.File().Close()
    	// Now extract one file. We chdir to the directory of the archive for simplicity.
    	pwd, err := os.Getwd()
    	if err != nil {
    		t.Fatal("os.Getwd: ", err)
    	}
    	err = os.Chdir(dir)
    	if err != nil {
    		t.Fatal("os.Chdir: ", err)
    	}
    	defer func() {
    		err := os.Chdir(pwd)
    		if err != nil {
    			t.Fatal("os.Chdir: ", err)
    		}
    	}()
    	ar = openArchive(name, os.O_RDONLY, []string{goodbyeFile.name})
    	ar.scan(ar.extractContents)
    	ar.a.File().Close()
    	data, err := os.ReadFile(goodbyeFile.name)
    	if err != nil {
    		t.Fatal(err)
    	}
    	// Expect contents of file.
    	result := string(data)
    	expect := goodbyeFile.contents
    	if result != expect {
    		t.Fatalf("expected %q got %q", expect, result)
    	}
    }
    
    // Test that pack-created archives can be understood by the tools.
    func TestHello(t *testing.T) {
    	testenv.MustHaveGoBuild(t)
    
    	dir := t.TempDir()
    	hello := filepath.Join(dir, "hello.go")
    	prog := `
    		package main
    		func main() {
    			println("hello world")
    		}
    	`
    	err := os.WriteFile(hello, []byte(prog), 0666)
    	if err != nil {
    		t.Fatal(err)
    	}
    
    	run := func(args ...string) string {
    		return doRun(t, dir, args...)
    	}
    
    	goBin := testenv.GoToolPath(t)
    	run(goBin, "build", "cmd/pack") // writes pack binary to dir
    	run(goBin, "tool", "compile", "-p=main", "hello.go")
    	run("./pack", "grc", "hello.a", "hello.o")
    	run(goBin, "tool", "link", "-o", "a.out", "hello.a")
    	out := run("./a.out")
    	if out != "hello world\n" {
    		t.Fatalf("incorrect output: %q, want %q", out, "hello world\n")
    	}
    }
    
    // Test that pack works with very long lines in PKGDEF.
    func TestLargeDefs(t *testing.T) {
    	if testing.Short() {
    		t.Skip("skipping in -short mode")
    	}
    	testenv.MustHaveGoBuild(t)
    
    	dir := t.TempDir()
    	large := filepath.Join(dir, "large.go")
    	f, err := os.Create(large)
    	if err != nil {
    		t.Fatal(err)
    	}
    	b := bufio.NewWriter(f)
    
    	printf := func(format string, args ...any) {
    		_, err := fmt.Fprintf(b, format, args...)
    		if err != nil {
    			t.Fatalf("Writing to %s: %v", large, err)
    		}
    	}
    
    	printf("package large\n\ntype T struct {\n")
    	for i := 0; i < 1000; i++ {
    		printf("f%d int `tag:\"", i)
    		for j := 0; j < 100; j++ {
    			printf("t%d=%d,", j, j)
    		}
    		printf("\"`\n")
    	}
    	printf("}\n")
    	if err = b.Flush(); err != nil {
    		t.Fatal(err)
    	}
    	if err = f.Close(); err != nil {
    		t.Fatal(err)
    	}
    
    	main := filepath.Join(dir, "main.go")
    	prog := `
    		package main
    		import "large"
    		var V large.T
    		func main() {
    			println("ok")
    		}
    	`
    	err = os.WriteFile(main, []byte(prog), 0666)
    	if err != nil {
    		t.Fatal(err)
    	}
    
    	run := func(args ...string) string {
    		return doRun(t, dir, args...)
    	}
    
    	goBin := testenv.GoToolPath(t)
    	run(goBin, "build", "cmd/pack") // writes pack binary to dir
    	run(goBin, "tool", "compile", "-p=large", "large.go")
    	run("./pack", "grc", "large.a", "large.o")
    	run(goBin, "tool", "compile", "-p=main", "-I", ".", "main.go")
    	run(goBin, "tool", "link", "-L", ".", "-o", "a.out", "main.o")
    	out := run("./a.out")
    	if out != "ok\n" {
    		t.Fatalf("incorrect output: %q, want %q", out, "ok\n")
    	}
    }
    
    // Test that "\n!\n" inside export data doesn't result in a truncated
    // package definition when creating a .a archive from a .o Go object.
    func TestIssue21703(t *testing.T) {
    	testenv.MustHaveGoBuild(t)
    
    	dir := t.TempDir()
    
    	const aSrc = `package a; const X = "\n!\n"`
    	err := os.WriteFile(filepath.Join(dir, "a.go"), []byte(aSrc), 0666)
    	if err != nil {
    		t.Fatal(err)
    	}
    
    	const bSrc = `package b; import _ "a"`
    	err = os.WriteFile(filepath.Join(dir, "b.go"), []byte(bSrc), 0666)
    	if err != nil {
    		t.Fatal(err)
    	}
    
    	run := func(args ...string) string {
    		return doRun(t, dir, args...)
    	}
    
    	goBin := testenv.GoToolPath(t)
    	run(goBin, "build", "cmd/pack") // writes pack binary to dir
    	run(goBin, "tool", "compile", "-p=a", "a.go")
    	run("./pack", "c", "a.a", "a.o")
    	run(goBin, "tool", "compile", "-p=b", "-I", ".", "b.go")
    }
    
    // Test the "c" command can "see through" the archive generated by the compiler.
    // This is peculiar. (See issue #43271)
    func TestCreateWithCompilerObj(t *testing.T) {
    	testenv.MustHaveGoBuild(t)
    
    	dir := t.TempDir()
    	src := filepath.Join(dir, "p.go")
    	prog := "package p; var X = 42\n"
    	err := os.WriteFile(src, []byte(prog), 0666)
    	if err != nil {
    		t.Fatal(err)
    	}
    
    	run := func(args ...string) string {
    		return doRun(t, dir, args...)
    	}
    
    	goBin := testenv.GoToolPath(t)
    	run(goBin, "build", "cmd/pack") // writes pack binary to dir
    	run(goBin, "tool", "compile", "-pack", "-p=p", "-o", "p.a", "p.go")
    	run("./pack", "c", "packed.a", "p.a")
    	fi, err := os.Stat(filepath.Join(dir, "p.a"))
    	if err != nil {
    		t.Fatalf("stat p.a failed: %v", err)
    	}
    	fi2, err := os.Stat(filepath.Join(dir, "packed.a"))
    	if err != nil {
    		t.Fatalf("stat packed.a failed: %v", err)
    	}
    	// For compiler-generated object file, the "c" command is
    	// expected to get (essentially) the same file back, instead
    	// of packing it into a new archive with a single entry.
    	if want, got := fi.Size(), fi2.Size(); want != got {
    		t.Errorf("packed file with different size: want %d, got %d", want, got)
    	}
    
    	// Test -linkobj flag as well.
    	run(goBin, "tool", "compile", "-p=p", "-linkobj", "p2.a", "-o", "p.x", "p.go")
    	run("./pack", "c", "packed2.a", "p2.a")
    	fi, err = os.Stat(filepath.Join(dir, "p2.a"))
    	if err != nil {
    		t.Fatalf("stat p2.a failed: %v", err)
    	}
    	fi2, err = os.Stat(filepath.Join(dir, "packed2.a"))
    	if err != nil {
    		t.Fatalf("stat packed2.a failed: %v", err)
    	}
    	if want, got := fi.Size(), fi2.Size(); want != got {
    		t.Errorf("packed file with different size: want %d, got %d", want, got)
    	}
    
    	run("./pack", "c", "packed3.a", "p.x")
    	fi, err = os.Stat(filepath.Join(dir, "p.x"))
    	if err != nil {
    		t.Fatalf("stat p.x failed: %v", err)
    	}
    	fi2, err = os.Stat(filepath.Join(dir, "packed3.a"))
    	if err != nil {
    		t.Fatalf("stat packed3.a failed: %v", err)
    	}
    	if want, got := fi.Size(), fi2.Size(); want != got {
    		t.Errorf("packed file with different size: want %d, got %d", want, got)
    	}
    }
    
    // Test the "r" command creates the output file if it does not exist.
    func TestRWithNonexistentFile(t *testing.T) {
    	testenv.MustHaveGoBuild(t)
    
    	dir := t.TempDir()
    	src := filepath.Join(dir, "p.go")
    	prog := "package p; var X = 42\n"
    	err := os.WriteFile(src, []byte(prog), 0666)
    	if err != nil {
    		t.Fatal(err)
    	}
    
    	run := func(args ...string) string {
    		return doRun(t, dir, args...)
    	}
    
    	goBin := testenv.GoToolPath(t)
    	run(goBin, "build", "cmd/pack") // writes pack binary to dir
    	run(goBin, "tool", "compile", "-p=p", "-o", "p.o", "p.go")
    	run("./pack", "r", "p.a", "p.o") // should succeed
    }
    
    // doRun runs a program in a directory and returns the output.
    func doRun(t *testing.T, dir string, args ...string) string {
    	cmd := exec.Command(args[0], args[1:]...)
    	cmd.Dir = dir
    	out, err := cmd.CombinedOutput()
    	if err != nil {
    		t.Fatalf("%v: %v\n%s", args, err, string(out))
    	}
    	return string(out)
    }
    
    // Fake implementation of files.
    
    var helloFile = &FakeFile{
    	name:     "hello",
    	contents: "hello world", // 11 bytes, an odd number.
    	mode:     0644,
    }
    
    var goodbyeFile = &FakeFile{
    	name:     "goodbye",
    	contents: "Sayonara, Jim", // 13 bytes, another odd number.
    	mode:     0644,
    }
    
    // FakeFile implements FileLike and also fs.FileInfo.
    type FakeFile struct {
    	name     string
    	contents string
    	mode     fs.FileMode
    	offset   int
    }
    
    // Reset prepares a FakeFile for reuse.
    func (f *FakeFile) Reset() *FakeFile {
    	f.offset = 0
    	return f
    }
    
    // FileLike methods.
    
    func (f *FakeFile) Name() string {
    	// A bit of a cheat: we only have a basename, so that's also ok for FileInfo.
    	return f.name
    }
    
    func (f *FakeFile) Stat() (fs.FileInfo, error) {
    	return f, nil
    }
    
    func (f *FakeFile) Read(p []byte) (int, error) {
    	if f.offset >= len(f.contents) {
    		return 0, io.EOF
    	}
    	n := copy(p, f.contents[f.offset:])
    	f.offset += n
    	return n, nil
    }
    
    func (f *FakeFile) Close() error {
    	return nil
    }
    
    // fs.FileInfo methods.
    
    func (f *FakeFile) Size() int64 {
    	return int64(len(f.contents))
    }
    
    func (f *FakeFile) Mode() fs.FileMode {
    	return f.mode
    }
    
    func (f *FakeFile) ModTime() time.Time {
    	return time.Time{}
    }
    
    func (f *FakeFile) IsDir() bool {
    	return false
    }
    
    func (f *FakeFile) Sys() any {
    	return nil
    }
    
    // Special helpers.
    
    func (f *FakeFile) Entry() *archive.Entry {
    	return &archive.Entry{
    		Name:  f.name,
    		Mtime: 0, // Defined to be zero.
    		Uid:   0, // Ditto.
    		Gid:   0, // Ditto.
    		Mode:  f.mode,
    		Data:  archive.Data{Size: int64(len(f.contents))},
    	}
    }
    ```

