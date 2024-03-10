+++
title = "vet"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# vet

> 原文：[https://pkg.go.dev/cmd/vet@go1.19.3](https://pkg.go.dev/cmd/vet@go1.19.3)

### Overview 概述

Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string. Vet uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers.

​	vet检查Go源代码并报告可疑的结构，例如参数与格式字符串不一致的Printf调用。vet 使用启发式方法，不能保证所有报告都是真正的问题，但它可以发现编译器没有发现的错误。

Vet is normally invoked through the go command. This command vets the package in the current directory:

​	vet通常是通过go命令调用的。该命令对当前目录下的软件包进行审查：

```
go vet
```

whereas this one vets the packages whose path is provided:

而这个命令是对提供了路径的软件包进行审查：

```
go vet my/project/...
```

Use "go help packages" to see other ways of specifying which packages to vet.

​	使用 "go help packages "来查看其他指定审查软件包的方法。

Vet's exit code is non-zero for erroneous invocation of the tool or if a problem was reported, and 0 otherwise. Note that the tool does not check every possible problem and depends on unreliable heuristics, so it should be used as guidance only, not as a firm indicator of program correctness.

​	vet的退出代码在错误地调用工具或报告问题时为非零，否则为0。请注意，该工具并没有检查每一个可能的问题，而是依赖于不可靠的启发式方法，所以它应该只作为指导，而不是作为程序正确性的一个确定指标。

To list the available checks, run "go tool vet help":

​	要列出可用的检查，请运行 "go tool vet help"：

```
asmdecl      report mismatches between assembly files and Go declarations
			 检查汇编文件与Go声明之间的不匹配情况			 
assign       check for useless assignments
			 检查无用的赋值操作	
atomic       check for common mistakes using the sync/atomic package
			 检查使用sync/atomic包时常见的错误
bools        check for common mistakes involving boolean operators
			 检查涉及布尔运算符的常见错误
buildtag     check that +build tags are well-formed and correctly located
			 检查+build标签是否格式良好且位置正确
cgocall      detect some violations of the cgo pointer passing rules
			 检测一些违反cgo指针传递规则的情况
composites   check for unkeyed composite literals
			 检查未键控的复合文字
copylocks    check for locks erroneously passed by value
			 检查错误地通过值传递锁的情况
httpresponse check for mistakes using HTTP responses
			 检查使用HTTP响应时的错误
loopclosure  check references to loop variables from within nested functions
			 检查从嵌套函数中引用循环变量的情况
lostcancel   check cancel func returned by context.WithCancel is called
			 检查是否调用了由context.WithCancel返回的cancel函数
nilfunc      check for useless comparisons between functions and nil
			 检查函数与nil之间无用的比较
printf       check consistency of Printf format strings and arguments
			 检查Printf格式字符串和实参的一致性
shift        check for shifts that equal or exceed the width of the integer
			 检查等于或超过整数宽度的位移操作
stdmethods   check signature of methods of well-known interfaces
			 检查已知接口方法的签名
structtag    check that struct field tags conform to reflect.StructTag.Get
			 检查结构字段标签是否符合reflect.StructTag.Get的规范
tests        check for common mistaken usages of tests and examples
			 检查测试和示例的常见错误用法
unmarshal    report passing non-pointer or non-interface values to unmarshal
			 报告将非指针或非接口值传递给unmarshal的情况
unreachable  check for unreachable code
			 检查不可到达的代码
unsafeptr    check for invalid conversions of uintptr to unsafe.Pointer
			 检查将uintptr无效转换为unsafe.Pointer的情况
unusedresult check for unused results of calls to some functions
			 检查对某些函数调用的未使用结果
```

For details and flags of a particular check, such as printf, run "go tool vet help printf".

​	要了解某个特定检查的细节和标志，例如printf，请运行 "go tool vet help printf"。

By default, all checks are performed. If any flags are explicitly set to true, only those tests are run. Conversely, if any flag is explicitly set to false, only those tests are disabled. Thus -printf=true runs the printf check, and -printf=false runs all checks except the printf check.

​	默认情况下，所有的检查都被执行。如果任何标志被明确地设置为 "true"，则只有这些测试被运行。反之，如果任何标志被明确设置为false，则只有这些测试被禁用。因此，-printf=true运行printf检查，而-printf=false运行除printf检查之外的所有检查。

For information on writing a new check, see golang.org/x/tools/go/analysis.

​	有关编写新检查的信息，请参见golang.org/x/tools/go/analysis。

Core flags:

​	核心标志：

```
-c=N
  	display offending line plus N lines of surrounding context
-json
  	emit analysis diagnostics (and errors) in JSON format
```



=== "main.go"

    ```go 
    // Copyright 2012 The Go Authors. All rights reserved.
    // Use of this source code is governed by a BSD-style
    // license that can be found in the LICENSE file.
    
    package main
    
    import (
    	"cmd/internal/objabi"
    
    	"golang.org/x/tools/go/analysis/unitchecker"
    
    	"golang.org/x/tools/go/analysis/passes/asmdecl"
    	"golang.org/x/tools/go/analysis/passes/assign"
    	"golang.org/x/tools/go/analysis/passes/atomic"
    	"golang.org/x/tools/go/analysis/passes/bools"
    	"golang.org/x/tools/go/analysis/passes/buildtag"
    	"golang.org/x/tools/go/analysis/passes/cgocall"
    	"golang.org/x/tools/go/analysis/passes/composite"
    	"golang.org/x/tools/go/analysis/passes/copylock"
    	"golang.org/x/tools/go/analysis/passes/errorsas"
    	"golang.org/x/tools/go/analysis/passes/framepointer"
    	"golang.org/x/tools/go/analysis/passes/httpresponse"
    	"golang.org/x/tools/go/analysis/passes/ifaceassert"
    	"golang.org/x/tools/go/analysis/passes/loopclosure"
    	"golang.org/x/tools/go/analysis/passes/lostcancel"
    	"golang.org/x/tools/go/analysis/passes/nilfunc"
    	"golang.org/x/tools/go/analysis/passes/printf"
    	"golang.org/x/tools/go/analysis/passes/shift"
    	"golang.org/x/tools/go/analysis/passes/sigchanyzer"
    	"golang.org/x/tools/go/analysis/passes/stdmethods"
    	"golang.org/x/tools/go/analysis/passes/stringintconv"
    	"golang.org/x/tools/go/analysis/passes/structtag"
    	"golang.org/x/tools/go/analysis/passes/testinggoroutine"
    	"golang.org/x/tools/go/analysis/passes/tests"
    	"golang.org/x/tools/go/analysis/passes/unmarshal"
    	"golang.org/x/tools/go/analysis/passes/unreachable"
    	"golang.org/x/tools/go/analysis/passes/unsafeptr"
    	"golang.org/x/tools/go/analysis/passes/unusedresult"
    )
    
    func main() {
    	objabi.AddVersionFlag()
    
    	unitchecker.Main(
    		asmdecl.Analyzer,
    		assign.Analyzer,
    		atomic.Analyzer,
    		bools.Analyzer,
    		buildtag.Analyzer,
    		cgocall.Analyzer,
    		composite.Analyzer,
    		copylock.Analyzer,
    		errorsas.Analyzer,
    		framepointer.Analyzer,
    		httpresponse.Analyzer,
    		ifaceassert.Analyzer,
    		loopclosure.Analyzer,
    		lostcancel.Analyzer,
    		nilfunc.Analyzer,
    		printf.Analyzer,
    		shift.Analyzer,
    		sigchanyzer.Analyzer,
    		stdmethods.Analyzer,
    		stringintconv.Analyzer,
    		structtag.Analyzer,
    		tests.Analyzer,
    		testinggoroutine.Analyzer,
    		unmarshal.Analyzer,
    		unreachable.Analyzer,
    		unsafeptr.Analyzer,
    		unusedresult.Analyzer,
    	)
    }
    ```

=== "vet_test.go"

    ```go 
    // Copyright 2013 The Go Authors. All rights reserved.
    // Use of this source code is governed by a BSD-style
    // license that can be found in the LICENSE file.
    
    package main_test
    
    import (
    	"bytes"
    	"errors"
    	"fmt"
    	"internal/testenv"
    	"log"
    	"os"
    	"os/exec"
    	"path"
    	"path/filepath"
    	"regexp"
    	"strconv"
    	"strings"
    	"sync"
    	"testing"
    )
    
    const dataDir = "testdata"
    
    var binary string
    
    // We implement TestMain so remove the test binary when all is done.
    func TestMain(m *testing.M) {
    	os.Exit(testMain(m))
    }
    
    func testMain(m *testing.M) int {
    	dir, err := os.MkdirTemp("", "vet_test")
    	if err != nil {
    		fmt.Fprintln(os.Stderr, err)
    		return 1
    	}
    	defer os.RemoveAll(dir)
    	binary = filepath.Join(dir, "testvet.exe")
    
    	return m.Run()
    }
    
    var (
    	buildMu sync.Mutex // guards following
    	built   = false    // We have built the binary.
    	failed  = false    // We have failed to build the binary, don't try again.
    )
    
    func Build(t *testing.T) {
    	buildMu.Lock()
    	defer buildMu.Unlock()
    	if built {
    		return
    	}
    	if failed {
    		t.Skip("cannot run on this environment")
    	}
    	testenv.MustHaveGoBuild(t)
    	cmd := exec.Command(testenv.GoToolPath(t), "build", "-o", binary)
    	output, err := cmd.CombinedOutput()
    	if err != nil {
    		failed = true
    		fmt.Fprintf(os.Stderr, "%s\n", output)
    		t.Fatal(err)
    	}
    	built = true
    }
    
    func vetCmd(t *testing.T, arg, pkg string) *exec.Cmd {
    	cmd := exec.Command(testenv.GoToolPath(t), "vet", "-vettool="+binary, arg, path.Join("cmd/vet/testdata", pkg))
    	cmd.Env = os.Environ()
    	return cmd
    }
    
    func TestVet(t *testing.T) {
    	t.Parallel()
    	Build(t)
    	for _, pkg := range []string{
    		"asm",
    		"assign",
    		"atomic",
    		"bool",
    		"buildtag",
    		"cgo",
    		"composite",
    		"copylock",
    		"deadcode",
    		"httpresponse",
    		"lostcancel",
    		"method",
    		"nilfunc",
    		"print",
    		"rangeloop",
    		"shift",
    		"structtag",
    		"testingpkg",
    		// "testtag" has its own test
    		"unmarshal",
    		"unsafeptr",
    		"unused",
    	} {
    		pkg := pkg
    		t.Run(pkg, func(t *testing.T) {
    			t.Parallel()
    
    			// Skip cgo test on platforms without cgo.
    			if pkg == "cgo" && !cgoEnabled(t) {
    				return
    			}
    
    			cmd := vetCmd(t, "-printfuncs=Warn,Warnf", pkg)
    
    			// The asm test assumes amd64.
    			if pkg == "asm" {
    				cmd.Env = append(cmd.Env, "GOOS=linux", "GOARCH=amd64")
    			}
    
    			dir := filepath.Join("testdata", pkg)
    			gos, err := filepath.Glob(filepath.Join(dir, "*.go"))
    			if err != nil {
    				t.Fatal(err)
    			}
    			asms, err := filepath.Glob(filepath.Join(dir, "*.s"))
    			if err != nil {
    				t.Fatal(err)
    			}
    			var files []string
    			files = append(files, gos...)
    			files = append(files, asms...)
    
    			errchk(cmd, files, t)
    		})
    	}
    }
    
    func cgoEnabled(t *testing.T) bool {
    	// Don't trust build.Default.CgoEnabled as it is false for
    	// cross-builds unless CGO_ENABLED is explicitly specified.
    	// That's fine for the builders, but causes commands like
    	// 'GOARCH=386 go test .' to fail.
    	// Instead, we ask the go command.
    	cmd := exec.Command(testenv.GoToolPath(t), "list", "-f", "{{context.CgoEnabled}}")
    	out, _ := cmd.CombinedOutput()
    	return string(out) == "true\n"
    }
    
    func errchk(c *exec.Cmd, files []string, t *testing.T) {
    	output, err := c.CombinedOutput()
    	if _, ok := err.(*exec.ExitError); !ok {
    		t.Logf("vet output:\n%s", output)
    		t.Fatal(err)
    	}
    	fullshort := make([]string, 0, len(files)*2)
    	for _, f := range files {
    		fullshort = append(fullshort, f, filepath.Base(f))
    	}
    	err = errorCheck(string(output), false, fullshort...)
    	if err != nil {
    		t.Errorf("error check failed: %s", err)
    	}
    }
    
    // TestTags verifies that the -tags argument controls which files to check.
    func TestTags(t *testing.T) {
    	t.Parallel()
    	Build(t)
    	for tag, wantFile := range map[string]int{
    		"testtag":     1, // file1
    		"x testtag y": 1,
    		"othertag":    2,
    	} {
    		tag, wantFile := tag, wantFile
    		t.Run(tag, func(t *testing.T) {
    			t.Parallel()
    			t.Logf("-tags=%s", tag)
    			cmd := vetCmd(t, "-tags="+tag, "tagtest")
    			output, err := cmd.CombinedOutput()
    
    			want := fmt.Sprintf("file%d.go", wantFile)
    			dontwant := fmt.Sprintf("file%d.go", 3-wantFile)
    
    			// file1 has testtag and file2 has !testtag.
    			if !bytes.Contains(output, []byte(filepath.Join("tagtest", want))) {
    				t.Errorf("%s: %s was excluded, should be included", tag, want)
    			}
    			if bytes.Contains(output, []byte(filepath.Join("tagtest", dontwant))) {
    				t.Errorf("%s: %s was included, should be excluded", tag, dontwant)
    			}
    			if t.Failed() {
    				t.Logf("err=%s, output=<<%s>>", err, output)
    			}
    		})
    	}
    }
    
    // All declarations below were adapted from test/run.go.
    
    // errorCheck matches errors in outStr against comments in source files.
    // For each line of the source files which should generate an error,
    // there should be a comment of the form // ERROR "regexp".
    // If outStr has an error for a line which has no such comment,
    // this function will report an error.
    // Likewise if outStr does not have an error for a line which has a comment,
    // or if the error message does not match the <regexp>.
    // The <regexp> syntax is Perl but it's best to stick to egrep.
    //
    // Sources files are supplied as fullshort slice.
    // It consists of pairs: full path to source file and its base name.
    func errorCheck(outStr string, wantAuto bool, fullshort ...string) (err error) {
    	var errs []error
    	out := splitOutput(outStr, wantAuto)
    	// Cut directory name.
    	for i := range out {
    		for j := 0; j < len(fullshort); j += 2 {
    			full, short := fullshort[j], fullshort[j+1]
    			out[i] = strings.ReplaceAll(out[i], full, short)
    		}
    	}
    
    	var want []wantedError
    	for j := 0; j < len(fullshort); j += 2 {
    		full, short := fullshort[j], fullshort[j+1]
    		want = append(want, wantedErrors(full, short)...)
    	}
    	for _, we := range want {
    		var errmsgs []string
    		if we.auto {
    			errmsgs, out = partitionStrings("<autogenerated>", out)
    		} else {
    			errmsgs, out = partitionStrings(we.prefix, out)
    		}
    		if len(errmsgs) == 0 {
    			errs = append(errs, fmt.Errorf("%s:%d: missing error %q", we.file, we.lineNum, we.reStr))
    			continue
    		}
    		matched := false
    		n := len(out)
    		for _, errmsg := range errmsgs {
    			// Assume errmsg says "file:line: foo".
    			// Cut leading "file:line: " to avoid accidental matching of file name instead of message.
    			text := errmsg
    			if _, suffix, ok := strings.Cut(text, " "); ok {
    				text = suffix
    			}
    			if we.re.MatchString(text) {
    				matched = true
    			} else {
    				out = append(out, errmsg)
    			}
    		}
    		if !matched {
    			errs = append(errs, fmt.Errorf("%s:%d: no match for %#q in:\n\t%s", we.file, we.lineNum, we.reStr, strings.Join(out[n:], "\n\t")))
    			continue
    		}
    	}
    
    	if len(out) > 0 {
    		errs = append(errs, fmt.Errorf("Unmatched Errors:"))
    		for _, errLine := range out {
    			errs = append(errs, fmt.Errorf("%s", errLine))
    		}
    	}
    
    	if len(errs) == 0 {
    		return nil
    	}
    	if len(errs) == 1 {
    		return errs[0]
    	}
    	var buf bytes.Buffer
    	fmt.Fprintf(&buf, "\n")
    	for _, err := range errs {
    		fmt.Fprintf(&buf, "%s\n", err.Error())
    	}
    	return errors.New(buf.String())
    }
    
    func splitOutput(out string, wantAuto bool) []string {
    	// gc error messages continue onto additional lines with leading tabs.
    	// Split the output at the beginning of each line that doesn't begin with a tab.
    	// <autogenerated> lines are impossible to match so those are filtered out.
    	var res []string
    	for _, line := range strings.Split(out, "\n") {
    		line = strings.TrimSuffix(line, "\r") // normalize Windows output
    		if strings.HasPrefix(line, "\t") {
    			res[len(res)-1] += "\n" + line
    		} else if strings.HasPrefix(line, "go tool") || strings.HasPrefix(line, "#") || !wantAuto && strings.HasPrefix(line, "<autogenerated>") {
    			continue
    		} else if strings.TrimSpace(line) != "" {
    			res = append(res, line)
    		}
    	}
    	return res
    }
    
    // matchPrefix reports whether s starts with file name prefix followed by a :,
    // and possibly preceded by a directory name.
    func matchPrefix(s, prefix string) bool {
    	i := strings.Index(s, ":")
    	if i < 0 {
    		return false
    	}
    	j := strings.LastIndex(s[:i], "/")
    	s = s[j+1:]
    	if len(s) <= len(prefix) || s[:len(prefix)] != prefix {
    		return false
    	}
    	if s[len(prefix)] == ':' {
    		return true
    	}
    	return false
    }
    
    func partitionStrings(prefix string, strs []string) (matched, unmatched []string) {
    	for _, s := range strs {
    		if matchPrefix(s, prefix) {
    			matched = append(matched, s)
    		} else {
    			unmatched = append(unmatched, s)
    		}
    	}
    	return
    }
    
    type wantedError struct {
    	reStr   string
    	re      *regexp.Regexp
    	lineNum int
    	auto    bool // match <autogenerated> line
    	file    string
    	prefix  string
    }
    
    var (
    	errRx       = regexp.MustCompile(`// (?:GC_)?ERROR(NEXT)? (.*)`)
    	errAutoRx   = regexp.MustCompile(`// (?:GC_)?ERRORAUTO(NEXT)? (.*)`)
    	errQuotesRx = regexp.MustCompile(`"([^"]*)"`)
    	lineRx      = regexp.MustCompile(`LINE(([+-])([0-9]+))?`)
    )
    
    // wantedErrors parses expected errors from comments in a file.
    func wantedErrors(file, short string) (errs []wantedError) {
    	cache := make(map[string]*regexp.Regexp)
    
    	src, err := os.ReadFile(file)
    	if err != nil {
    		log.Fatal(err)
    	}
    	for i, line := range strings.Split(string(src), "\n") {
    		lineNum := i + 1
    		if strings.Contains(line, "////") {
    			// double comment disables ERROR
    			continue
    		}
    		var auto bool
    		m := errAutoRx.FindStringSubmatch(line)
    		if m != nil {
    			auto = true
    		} else {
    			m = errRx.FindStringSubmatch(line)
    		}
    		if m == nil {
    			continue
    		}
    		if m[1] == "NEXT" {
    			lineNum++
    		}
    		all := m[2]
    		mm := errQuotesRx.FindAllStringSubmatch(all, -1)
    		if mm == nil {
    			log.Fatalf("%s:%d: invalid errchk line: %s", file, lineNum, line)
    		}
    		for _, m := range mm {
    			replacedOnce := false
    			rx := lineRx.ReplaceAllStringFunc(m[1], func(m string) string {
    				if replacedOnce {
    					return m
    				}
    				replacedOnce = true
    				n := lineNum
    				if strings.HasPrefix(m, "LINE+") {
    					delta, _ := strconv.Atoi(m[5:])
    					n += delta
    				} else if strings.HasPrefix(m, "LINE-") {
    					delta, _ := strconv.Atoi(m[5:])
    					n -= delta
    				}
    				return fmt.Sprintf("%s:%d", short, n)
    			})
    			re := cache[rx]
    			if re == nil {
    				var err error
    				re, err = regexp.Compile(rx)
    				if err != nil {
    					log.Fatalf("%s:%d: invalid regexp \"%#q\" in ERROR line: %v", file, lineNum, rx, err)
    				}
    				cache[rx] = re
    			}
    			prefix := fmt.Sprintf("%s:%d", short, lineNum)
    			errs = append(errs, wantedError{
    				reStr:   rx,
    				re:      re,
    				prefix:  prefix,
    				auto:    auto,
    				lineNum: lineNum,
    				file:    short,
    			})
    		}
    	}
    
    	return
    }
    ```