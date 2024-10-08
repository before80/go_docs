+++
title = "objdump"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# objdump

> 原文：[https://pkg.go.dev/cmd/objdump@go1.19.3](https://pkg.go.dev/cmd/objdump@go1.19.3)

### Overview

Objdump disassembles executable files.

​	objdump 反汇编可执行文件。

Usage:

​	用法：

```
go tool objdump [-s symregexp] binary
```

Objdump prints a disassembly of all text symbols (code) in the binary. If the -s option is present, objdump only disassembles symbols with names matching the regular expression.

​	objdump 打印二进制文件中所有文本符号（代码）的反汇编。如果存在 -s 选项，objdump 只反汇编名称与正则表达式匹配的符号。

Alternate usage:

​	其他用法：

```
go tool objdump binary start end
```

In this mode, objdump disassembles the binary starting at the start address and stopping at the end address. The start and end addresses are program counters written in hexadecimal with optional leading 0x prefix. In this mode, objdump prints a sequence of stanzas of the form:

​	在此模式下，objdump 从起始地址开始反汇编二进制文件，并在结束地址停止。起始地址和结束地址是十六进制形式的程序计数器，前面可以加上可选的前缀 0x。在此模式下，objdump 打印一系列形式为：

```
file:line
 address: assembly
 address: assembly
 ...
```

Each stanza gives the disassembly for a contiguous range of addresses all mapped to the same original source file and line number. This mode is intended for use by pprof.

​	每个节都给出连续地址范围的反汇编，所有地址都映射到相同的原始源文件和行号。此模式旨在供 pprof 使用。

=== "main.go"

    ``` go 
    // Copyright 2012 The Go Authors. All rights reserved.
    // Use of this source code is governed by a BSD-style
    // license that can be found in the LICENSE file.
    
    // Objdump disassembles executable files.
    //
    // Usage:
    //
    //	go tool objdump [-s symregexp] binary
    //
    // Objdump prints a disassembly of all text symbols (code) in the binary.
    // If the -s option is present, objdump only disassembles
    // symbols with names matching the regular expression.
    //
    // Alternate usage:
    //
    //	go tool objdump binary start end
    //
    // In this mode, objdump disassembles the binary starting at the start address and
    // stopping at the end address. The start and end addresses are program
    // counters written in hexadecimal with optional leading 0x prefix.
    // In this mode, objdump prints a sequence of stanzas of the form:
    //
    //	file:line
    //	 address: assembly
    //	 address: assembly
    //	 ...
    //
    // Each stanza gives the disassembly for a contiguous range of addresses
    // all mapped to the same original source file and line number.
    // This mode is intended for use by pprof.
    package main
    
    import (
    	"flag"
    	"fmt"
    	"log"
    	"os"
    	"regexp"
    	"strconv"
    	"strings"
    
    	"cmd/internal/objfile"
    )
    
    var printCode = flag.Bool("S", false, "print Go code alongside assembly")
    var symregexp = flag.String("s", "", "only dump symbols matching this regexp")
    var gnuAsm = flag.Bool("gnu", false, "print GNU assembly next to Go assembly (where supported)")
    var symRE *regexp.Regexp
    
    func usage() {
    	fmt.Fprintf(os.Stderr, "usage: go tool objdump [-S] [-gnu] [-s symregexp] binary [start end]\n\n")
    	flag.PrintDefaults()
    	os.Exit(2)
    }
    
    func main() {
    	log.SetFlags(0)
    	log.SetPrefix("objdump: ")
    
    	flag.Usage = usage
    	flag.Parse()
    	if flag.NArg() != 1 && flag.NArg() != 3 {
    		usage()
    	}
    
    	if *symregexp != "" {
    		re, err := regexp.Compile(*symregexp)
    		if err != nil {
    			log.Fatalf("invalid -s regexp: %v", err)
    		}
    		symRE = re
    	}
    
    	f, err := objfile.Open(flag.Arg(0))
    	if err != nil {
    		log.Fatal(err)
    	}
    	defer f.Close()
    
    	dis, err := f.Disasm()
    	if err != nil {
    		log.Fatalf("disassemble %s: %v", flag.Arg(0), err)
    	}
    
    	switch flag.NArg() {
    	default:
    		usage()
    	case 1:
    		// disassembly of entire object
    		dis.Print(os.Stdout, symRE, 0, ^uint64(0), *printCode, *gnuAsm)
    
    	case 3:
    		// disassembly of PC range
    		start, err := strconv.ParseUint(strings.TrimPrefix(flag.Arg(1), "0x"), 16, 64)
    		if err != nil {
    			log.Fatalf("invalid start PC: %v", err)
    		}
    		end, err := strconv.ParseUint(strings.TrimPrefix(flag.Arg(2), "0x"), 16, 64)
    		if err != nil {
    			log.Fatalf("invalid end PC: %v", err)
    		}
    		dis.Print(os.Stdout, symRE, start, end, *printCode, *gnuAsm)
    	}
    }
    ```

=== "objdump_test.go"

    ```go 
    // Copyright 2014 The Go Authors. All rights reserved.
    // Use of this source code is governed by a BSD-style
    // license that can be found in the LICENSE file.
    
    package main
    
    import (
        "cmd/internal/notsha256"
        "flag"
        "fmt"
        "go/build"
        "internal/testenv"
        "os"
        "os/exec"
        "path/filepath"
        "runtime"
        "strings"
        "testing"
    )
    
    var tmp, exe string // populated by buildObjdump
    
    func TestMain(m *testing.M) {
        if !testenv.HasGoBuild() {
            return
        }
    
        var exitcode int
        if err := buildObjdump(); err == nil {
            exitcode = m.Run()
        } else {
            fmt.Println(err)
            exitcode = 1
        }
        os.RemoveAll(tmp)
        os.Exit(exitcode)
    }
    
    func buildObjdump() error {
        var err error
        tmp, err = os.MkdirTemp("", "TestObjDump")
        if err != nil {
            return fmt.Errorf("TempDir failed: %v", err)
        }
    
        exe = filepath.Join(tmp, "testobjdump.exe")
        gotool, err := testenv.GoTool()
        if err != nil {
            return err
        }
        out, err := exec.Command(gotool, "build", "-o", exe, "cmd/objdump").CombinedOutput()
        if err != nil {
            os.RemoveAll(tmp)
            return fmt.Errorf("go build -o %v cmd/objdump: %v\n%s", exe, err, string(out))
        }
    
        return nil
    }
    
    var x86Need = []string{ // for both 386 and AMD64
        "JMP main.main(SB)",
        "CALL main.Println(SB)",
        "RET",
    }
    
    var amd64GnuNeed = []string{
        "jmp",
        "callq",
        "cmpb",
    }
    
    var i386GnuNeed = []string{
        "jmp",
        "call",
        "cmp",
    }
    
    var armNeed = []string{
        "B main.main(SB)",
        "BL main.Println(SB)",
        "RET",
    }
    
    var arm64Need = []string{
        "JMP main.main(SB)",
        "CALL main.Println(SB)",
        "RET",
    }
    
    var armGnuNeed = []string{ // for both ARM and AMR64
        "ldr",
        "bl",
        "cmp",
    }
    
    var ppcNeed = []string{
        "BR main.main(SB)",
        "CALL main.Println(SB)",
        "RET",
    }
    
    var ppcGnuNeed = []string{
        "mflr",
        "lbz",
        "beq",
    }
    
    func mustHaveDisasm(t *testing.T) {
        switch runtime.GOARCH {
        case "loong64":
            t.Skipf("skipping on %s", runtime.GOARCH)
        case "mips", "mipsle", "mips64", "mips64le":
            t.Skipf("skipping on %s, issue 12559", runtime.GOARCH)
        case "riscv64":
            t.Skipf("skipping on %s, issue 36738", runtime.GOARCH)
        case "s390x":
            t.Skipf("skipping on %s, issue 15255", runtime.GOARCH)
        }
    }
    
    var target = flag.String("target", "", "test disassembly of `goos/goarch` binary")
    
    // objdump is fully cross platform: it can handle binaries
    // from any known operating system and architecture.
    // We could in principle add binaries to testdata and check
    // all the supported systems during this test. However, the
    // binaries would be about 1 MB each, and we don't want to
    // add that much junk to the hg repository. Instead, build a
    // binary for the current system (only) and test that objdump
    // can handle that one.
    
    func testDisasm(t *testing.T, srcfname string, printCode bool, printGnuAsm bool, flags ...string) {
        mustHaveDisasm(t)
        goarch := runtime.GOARCH
        if *target != "" {
            f := strings.Split(*target, "/")
            if len(f) != 2 {
                t.Fatalf("-target argument must be goos/goarch")
            }
            defer os.Setenv("GOOS", os.Getenv("GOOS"))
            defer os.Setenv("GOARCH", os.Getenv("GOARCH"))
            os.Setenv("GOOS", f[0])
            os.Setenv("GOARCH", f[1])
            goarch = f[1]
        }
    
        hash := notsha256.Sum256([]byte(fmt.Sprintf("%v-%v-%v-%v", srcfname, flags, printCode, printGnuAsm)))
        hello := filepath.Join(tmp, fmt.Sprintf("hello-%x.exe", hash))
        args := []string{"build", "-o", hello}
        args = append(args, flags...)
        args = append(args, srcfname)
        cmd := exec.Command(testenv.GoToolPath(t), args...)
        // "Bad line" bug #36683 is sensitive to being run in the source directory.
        cmd.Dir = "testdata"
        // Ensure that the source file location embedded in the binary matches our
        // actual current GOROOT, instead of GOROOT_FINAL if set.
        cmd.Env = append(os.Environ(), "GOROOT_FINAL=")
        t.Logf("Running %v", cmd.Args)
        out, err := cmd.CombinedOutput()
        if err != nil {
            t.Fatalf("go build %s: %v\n%s", srcfname, err, out)
        }
        need := []string{
            "TEXT main.main(SB)",
        }
    
        if printCode {
            need = append(need, `	Println("hello, world")`)
        } else {
            need = append(need, srcfname+":6")
        }
    
        switch goarch {
        case "amd64", "386":
            need = append(need, x86Need...)
        case "arm":
            need = append(need, armNeed...)
        case "arm64":
            need = append(need, arm64Need...)
        case "ppc64", "ppc64le":
            need = append(need, ppcNeed...)
        }
    
        if printGnuAsm {
            switch goarch {
            case "amd64":
                need = append(need, amd64GnuNeed...)
            case "386":
                need = append(need, i386GnuNeed...)
            case "arm", "arm64":
                need = append(need, armGnuNeed...)
            case "ppc64", "ppc64le":
                need = append(need, ppcGnuNeed...)
            }
        }
        args = []string{
            "-s", "main.main",
            hello,
        }
    
        if printCode {
            args = append([]string{"-S"}, args...)
        }
    
        if printGnuAsm {
            args = append([]string{"-gnu"}, args...)
        }
        cmd = exec.Command(exe, args...)
        cmd.Dir = "testdata" // "Bad line" bug #36683 is sensitive to being run in the source directory
        out, err = cmd.CombinedOutput()
        t.Logf("Running %v", cmd.Args)
    
        if err != nil {
            exename := srcfname[:len(srcfname)-len(filepath.Ext(srcfname))] + ".exe"
            t.Fatalf("objdump %q: %v\n%s", exename, err, out)
        }
    
        text := string(out)
        ok := true
        for _, s := range need {
            if !strings.Contains(text, s) {
                t.Errorf("disassembly missing '%s'", s)
                ok = false
            }
        }
        if goarch == "386" {
            if strings.Contains(text, "(IP)") {
                t.Errorf("disassembly contains PC-Relative addressing on 386")
                ok = false
            }
        }
    
        if !ok || testing.Verbose() {
            t.Logf("full disassembly:\n%s", text)
        }
    }
    
    func testGoAndCgoDisasm(t *testing.T, printCode bool, printGnuAsm bool) {
        t.Parallel()
        testDisasm(t, "fmthello.go", printCode, printGnuAsm)
        if build.Default.CgoEnabled {
            testDisasm(t, "fmthellocgo.go", printCode, printGnuAsm)
        }
    }
    
    func TestDisasm(t *testing.T) {
        testGoAndCgoDisasm(t, false, false)
    }
    
    func TestDisasmCode(t *testing.T) {
        testGoAndCgoDisasm(t, true, false)
    }
    
    func TestDisasmGnuAsm(t *testing.T) {
        testGoAndCgoDisasm(t, false, true)
    }
    
    func TestDisasmExtld(t *testing.T) {
        testenv.MustHaveCGO(t)
        switch runtime.GOOS {
        case "plan9", "windows":
            t.Skipf("skipping on %s", runtime.GOOS)
        }
        t.Parallel()
        testDisasm(t, "fmthello.go", false, false, "-ldflags=-linkmode=external")
    }
    
    func TestDisasmGoobj(t *testing.T) {
        mustHaveDisasm(t)
    
        hello := filepath.Join(tmp, "hello.o")
        args := []string{"tool", "compile", "-p=main", "-o", hello}
        args = append(args, "testdata/fmthello.go")
        out, err := exec.Command(testenv.GoToolPath(t), args...).CombinedOutput()
        if err != nil {
            t.Fatalf("go tool compile fmthello.go: %v\n%s", err, out)
        }
        need := []string{
            "main(SB)",
            "fmthello.go:6",
        }
    
        args = []string{
            "-s", "main",
            hello,
        }
    
        out, err = exec.Command(exe, args...).CombinedOutput()
        if err != nil {
            t.Fatalf("objdump fmthello.o: %v\n%s", err, out)
        }
    
        text := string(out)
        ok := true
        for _, s := range need {
            if !strings.Contains(text, s) {
                t.Errorf("disassembly missing '%s'", s)
                ok = false
            }
        }
        if runtime.GOARCH == "386" {
            if strings.Contains(text, "(IP)") {
                t.Errorf("disassembly contains PC-Relative addressing on 386")
                ok = false
            }
        }
        if !ok {
            t.Logf("full disassembly:\n%s", text)
        }
    }
    
    func TestGoobjFileNumber(t *testing.T) {
        // Test that file table in Go object file is parsed correctly.
        testenv.MustHaveGoBuild(t)
        mustHaveDisasm(t)
    
        t.Parallel()
    
        tmpdir, err := os.MkdirTemp("", "TestGoobjFileNumber")
        if err != nil {
            t.Fatal(err)
        }
        defer os.RemoveAll(tmpdir)
    
        obj := filepath.Join(tmpdir, "p.a")
        cmd := exec.Command(testenv.GoToolPath(t), "build", "-o", obj)
        cmd.Dir = filepath.Join("testdata/testfilenum")
        out, err := cmd.CombinedOutput()
        if err != nil {
            t.Fatalf("build failed: %v\n%s", err, out)
        }
    
        cmd = exec.Command(exe, obj)
        out, err = cmd.CombinedOutput()
        if err != nil {
            t.Fatalf("objdump failed: %v\n%s", err, out)
        }
    
        text := string(out)
        for _, s := range []string{"a.go", "b.go", "c.go"} {
            if !strings.Contains(text, s) {
                t.Errorf("output missing '%s'", s)
            }
        }
    
        if t.Failed() {
            t.Logf("output:\n%s", text)
        }
    }
    
    func TestGoObjOtherVersion(t *testing.T) {
        testenv.MustHaveExec(t)
        t.Parallel()
    
        obj := filepath.Join("testdata", "go116.o")
        cmd := exec.Command(exe, obj)
        out, err := cmd.CombinedOutput()
        if err == nil {
            t.Fatalf("objdump go116.o succeeded unexpectedly")
        }
        if !strings.Contains(string(out), "go object of a different version") {
            t.Errorf("unexpected error message:\n%s", out)
        }
    }
    ```

