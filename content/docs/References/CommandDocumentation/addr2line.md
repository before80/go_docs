+++
title = "addr2line"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# addr2line

> 原文：[https://pkg.go.dev/cmd/addr2line@go1.19.3](https://pkg.go.dev/cmd/addr2line@go1.19.3)

### Overview 概述

​	`addr2line`是对`GNU addr2line`工具的最小模拟，足以支持`pprof`。

使用方法：

```
go tool addr2line binary
```

​	`addr2line`从标准输入中读取十六进制地址，每行一个，前缀为可选的`0x`。对于每个输入地址，`addr2line`打印两个输出行，首先是包含该地址的函数名称，其次是该地址对应的源代码的file:line。

​	这个工具仅用于`pprof`；它的接口可能会改变，或者在未来的版本中被完全删除。

=== "main.go"

```go 
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Addr2line is a minimal simulation of the GNU addr2line tool,
// just enough to support pprof.
//
// Usage:
//
//	go tool addr2line binary
//
// Addr2line reads hexadecimal addresses, one per line and with optional 0x prefix,
// from standard input. For each input address, addr2line prints two output lines,
// first the name of the function containing the address and second the file:line
// of the source code corresponding to that address.
//
// This tool is intended for use only by pprof; its interface may change or
// it may be deleted entirely in future releases.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"cmd/internal/objfile"
)

func printUsage(w *os.File) {
	fmt.Fprintf(w, "usage: addr2line binary\n")
	fmt.Fprintf(w, "reads addresses from standard input and writes two lines for each:\n")
	fmt.Fprintf(w, "\tfunction name\n")
	fmt.Fprintf(w, "\tfile:line\n")
}

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("addr2line: ")

	// pprof expects this behavior when checking for addr2line
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		printUsage(os.Stdout)
		os.Exit(0)
	}

	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}

	f, err := objfile.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	tab, err := f.PCLineTable()
	if err != nil {
		log.Fatalf("reading %s: %v", flag.Arg(0), err)
	}

	stdin := bufio.NewScanner(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	for stdin.Scan() {
		p := stdin.Text()
		if strings.Contains(p, ":") {
			// Reverse translate file:line to pc.
			// This was an extension in the old C version of 'go tool addr2line'
			// and is probably not used by anyone, but recognize the syntax.
			// We don't have an implementation.
			fmt.Fprintf(stdout, "!reverse translation not implemented\n")
			continue
		}
		pc, _ := strconv.ParseUint(strings.TrimPrefix(p, "0x"), 16, 64)
		file, line, fn := tab.PCToLine(pc)
		name := "?"
		if fn != nil {
			name = fn.Name
		} else {
			file = "?"
			line = 0
		}
		fmt.Fprintf(stdout, "%s\n%s:%d\n", name, file, line)
	}
	stdout.Flush()
}
```

=== "addr2line_test.go"

```go 
// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"internal/testenv"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func loadSyms(t *testing.T) map[string]string {
	cmd := exec.Command(testenv.GoToolPath(t), "tool", "nm", os.Args[0])
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go tool nm %v: %v\n%s", os.Args[0], err, string(out))
	}
	syms := make(map[string]string)
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		if len(f) < 3 {
			continue
		}
		syms[f[2]] = f[0]
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("error reading symbols: %v", err)
	}
	return syms
}

func runAddr2Line(t *testing.T, exepath, addr string) (funcname, path, lineno string) {
	cmd := exec.Command(exepath, os.Args[0])
	cmd.Stdin = strings.NewReader(addr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("go tool addr2line %v: %v\n%s", os.Args[0], err, string(out))
	}
	f := strings.Split(string(out), "\n")
	if len(f) < 3 && f[2] == "" {
		t.Fatal("addr2line output must have 2 lines")
	}
	funcname = f[0]
	pathAndLineNo := f[1]
	f = strings.Split(pathAndLineNo, ":")
	if runtime.GOOS == "windows" && len(f) == 3 {
		// Reattach drive letter.
		f = []string{f[0] + ":" + f[1], f[2]}
	}
	if len(f) != 2 {
		t.Fatalf("no line number found in %q", pathAndLineNo)
	}
	return funcname, f[0], f[1]
}

const symName = "cmd/addr2line.TestAddr2Line"

func testAddr2Line(t *testing.T, exepath, addr string) {
	funcName, srcPath, srcLineNo := runAddr2Line(t, exepath, addr)
	if symName != funcName {
		t.Fatalf("expected function name %v; got %v", symName, funcName)
	}
	fi1, err := os.Stat("addr2line_test.go")
	if err != nil {
		t.Fatalf("Stat failed: %v", err)
	}

	// Debug paths are stored slash-separated, so convert to system-native.
	srcPath = filepath.FromSlash(srcPath)
	fi2, err := os.Stat(srcPath)

	// If GOROOT_FINAL is set and srcPath is not the file we expect, perhaps
	// srcPath has had GOROOT_FINAL substituted for GOROOT and GOROOT hasn't been
	// moved to its final location yet. If so, try the original location instead.
	if gorootFinal := os.Getenv("GOROOT_FINAL"); gorootFinal != "" &&
		(os.IsNotExist(err) || (err == nil && !os.SameFile(fi1, fi2))) {
		// srcPath is clean, but GOROOT_FINAL itself might not be.
		// (See https://golang.org/issue/41447.)
		gorootFinal = filepath.Clean(gorootFinal)

		if strings.HasPrefix(srcPath, gorootFinal) {
			fi2, err = os.Stat(runtime.GOROOT() + strings.TrimPrefix(srcPath, gorootFinal))
		}
	}

	if err != nil {
		t.Fatalf("Stat failed: %v", err)
	}
	if !os.SameFile(fi1, fi2) {
		t.Fatalf("addr2line_test.go and %s are not same file", srcPath)
	}
	if srcLineNo != "106" {
		t.Fatalf("line number = %v; want 106", srcLineNo)
	}
}

// This is line 106. The test depends on that.
func TestAddr2Line(t *testing.T) {
	testenv.MustHaveGoBuild(t)

	tmpDir, err := os.MkdirTemp("", "TestAddr2Line")
	if err != nil {
		t.Fatal("TempDir failed: ", err)
	}
	defer os.RemoveAll(tmpDir)

	// Build copy of test binary with debug symbols,
	// since the one running now may not have them.
	exepath := filepath.Join(tmpDir, "testaddr2line_test.exe")
	out, err := exec.Command(testenv.GoToolPath(t), "test", "-c", "-o", exepath, "cmd/addr2line").CombinedOutput()
	if err != nil {
		t.Fatalf("go test -c -o %v cmd/addr2line: %v\n%s", exepath, err, string(out))
	}
	os.Args[0] = exepath

	syms := loadSyms(t)

	exepath = filepath.Join(tmpDir, "testaddr2line.exe")
	out, err = exec.Command(testenv.GoToolPath(t), "build", "-o", exepath, "cmd/addr2line").CombinedOutput()
	if err != nil {
		t.Fatalf("go build -o %v cmd/addr2line: %v\n%s", exepath, err, string(out))
	}

	testAddr2Line(t, exepath, syms[symName])
	testAddr2Line(t, exepath, "0x"+syms[symName])
}
```

