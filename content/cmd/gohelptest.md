+++
title = "go help test"
date = 2024-02-24T10:57:59+08:00
type = "docs"
weight = 650
description = ""
isCJKLanguage = true
draft = false

+++

usage: go test [build/test flags] [packages] [build/test flags & test binary flags]

‘Go test’ automates testing the packages named by the import paths.

​	“go test” 自动化测试由导入路径命名的包。

It prints a summary of the test results in the format:

​	它以以下格式打印测试结果的摘要：

```
    ok   archive/tar   0.011s
    FAIL archive/zip   0.022s
    ok   compress/gzip 0.033s
    ...
```

followed by detailed output for each failed package.

​	然后是每个失败包的详细输出。

‘Go test’ recompiles each package along with any files with names matching the file pattern “`*_test.go"`.

​	“go test” 重新编译每个包以及任何文件名与文件模式“ `*_test.go"` ”匹配的文件。

These additional files can contain test functions, benchmark functions, fuzz tests and example functions. See ‘go help testfunc’ for more.

​	这些其他文件可以包含测试函数、基准函数、模糊测试和示例函数。有关更多信息，请参阅“go help testfunc”。

Each listed package causes the execution of a separate test binary. Files whose names begin with “`_`” (including “`_test.go`”) or “`.`” are ignored.

​	每个列出的包都会导致执行单独的测试二进制文件。以“`_`”（包括“`_test.go`”）或“`.`”开头的文件将被忽略。

Test files that declare a package with the suffix “`_test`” will be compiled as a separate package, and then linked and run with the main test binary._

​	声明后缀为“`_test`”的包的测试文件将作为单独的包进行编译，然后与主测试二进制文件链接并运行。

The go tool will ignore a directory named “testdata”, making it available to hold ancillary data needed by the tests.

​	go 工具将忽略名为“testdata”的目录，使其可用于保存测试所需的其他数据。

As part of building a test binary, go test runs go vet on the package and its test source files to identify significant problems. If go vet finds any problems, go test reports those and does not run the test binary. Only a high-confidence subset of the default go vet checks are used. That subset is: atomic, bool, buildtags, directive, errorsas, ifaceassert, nilfunc, printf, and stringintconv. You can see the documentation for these and other vet tests via “go doc cmd/vet”. To disable the running of go vet, use the -vet=off flag. To run all checks, use the -vet=all flag.

​	作为构建测试二进制文件的一部分，go test 在包及其测试源文件上运行 go vet 以识别重大问题。如果 go vet 发现任何问题，go test 将报告这些问题并不会运行测试二进制文件。仅使用默认 go vet 检查中高度可信的子集。该子集包括：atomic、bool、buildtags、directive、errorsas、ifaceassert、nilfunc、printf 和 stringintconv。您可以通过“go doc cmd/vet”查看这些和其他 vet 测试的文档。要禁用 go vet 的运行，请使用 -vet=off 标志。要运行所有检查，请使用 -vet=all 标志。

All test output and summary lines are printed to the go command’s standard output, even if the test printed them to its own standard error. (The go command’s standard error is reserved for printing errors building the tests.)

​	即使测试将它们打印到其自己的标准错误中，所有测试输出和摘要行也会打印到 go 命令的标准输出。（go 命令的标准错误保留用于打印构建测试的错误。）

The go command places `$GOROOT/bin` at the beginning of `$PATH` in the test’s environment, so that tests that execute ‘go’ commands use the same ‘go’ as the parent ‘go test’ command.

​	go 命令将 `$GOROOT/bin` 置于 `$PATH` 环境的开头，以便执行“go”命令的测试使用与父“go test”命令相同的“go”。

Go test runs in two different modes:

​	Go test 以两种不同的模式运行：

The first, called local directory mode, occurs when go test is invoked with no package arguments (for example, ‘go test’ or ‘go test -v’). In this mode, go test compiles the package sources and tests found in the current directory and then runs the resulting test binary. In this mode, caching (discussed below) is disabled. After the package test finishes, go test prints a summary line showing the test status (‘ok’ or ‘FAIL’), package name, and elapsed time.

​	第一种称为本地目录模式，当使用无包参数调用 go test 时发生（例如，“go test”或“go test -v”）。在此模式中，go test 编译在当前目录中找到的包源和测试，然后运行生成的测试二进制文件。在此模式中，禁用缓存（如下所述）。包测试完成后，go test 打印显示测试状态（“ok”或“FAIL”）、包名称和经过时间的摘要行。

The second, called package list mode, occurs when go test is invoked with explicit package arguments (for example ‘go test math’, ‘go test ./…’, and even ‘go test .’). In this mode, go test compiles and tests each of the packages listed on the command line. If a package test passes, go test prints only the final ‘ok’ summary line. If a package test fails, go test prints the full test output.

​	第二种称为包列表模式，当使用显式包参数调用 go test 时发生（例如“go test math”、“go test ./…”，甚至“go test .”）。在此模式中，go test 编译并测试命令行上列出的每个包。如果包测试通过，go test 仅打印最终的“ok”摘要行。如果包测试失败，go test 打印完整的测试输出。

If invoked with the -bench or -v flag, go test prints the full output even for passing package tests, in order to display the requested benchmark results or verbose logging. After the package tests for all of the listed packages finish, and their output is printed, go test prints a final ‘FAIL’ status if any package test has failed.

​	如果使用 -bench 或 -v 标志调用，go test 即使对于通过的包测试也会打印完整输出，以便显示请求的基准结果或详细日志记录。在所有列出的包的包测试完成后并打印其输出后，如果任何包测试失败，go test 会打印最终的“FAIL”状态。

In package list mode only, go test caches successful package test results to avoid unnecessary repeated running of tests. When the result of a test can be recovered from the cache, go test will redisplay the previous output instead of running the test binary again. When this happens, go test prints ‘(cached)’ in place of the elapsed time in the summary line.

​	仅在包列表模式下，go test 会缓存成功的包测试结果，以避免不必要地重复运行测试。当可以从缓存中恢复测试结果时，go test 将重新显示以前的输出，而不是再次运行测试二进制文件。发生这种情况时，go test 会在摘要行中用“（已缓存）”代替经过的时间。

The rule for a match in the cache is that the run involves the same test binary and the flags on the command line come entirely from a restricted set of ‘cacheable’ test flags, defined as -benchtime, -cpu, -list, -parallel, -run, -short, -timeout, -failfast, and -v.

​	缓存中匹配的规则是，运行涉及相同的测试二进制文件，并且命令行上的标志完全来自一组受限的“可缓存”测试标志，定义为 -benchtime、-cpu、-list、-parallel、-run、-short、-timeout、-failfast 和 -v。

If a run of go test has any test or non-test flags outside this set, the result is not cached. To disable test caching, use any test flag or argument other than the cacheable flags. The idiomatic way to disable test caching explicitly is to use -count=1. Tests that open files within the package’s source root (usually `$GOPATH`) or that consult environment variables only match future runs in which the files and environment variables are unchanged. A cached test result is treated as executing in no time at all, so a successful package test result will be cached and reused regardless of -timeout setting.

​	如果 go test 的运行有任何不在此集合中的测试或非测试标志，则结果不会被缓存。要禁用测试缓存，请使用除可缓存标志之外的任何测试标志或参数。显式禁用测试缓存的惯用方式是使用 -count=1。在包的源根目录（通常是 `$GOPATH` ）内打开文件或仅查询环境变量的测试仅匹配文件和环境变量保持不变的未来运行。缓存的测试结果被视为立即执行，因此无论 -timeout 设置如何，成功的包测试结果都将被缓存并重新使用。

In addition to the build flags, the flags handled by ‘go test’ itself are:

​	除了构建标志外，“go test” 本身处理的标志有：

```
    -args
        Pass the remainder of the command line (everything after -args)
        to the test binary, uninterpreted and unchanged.
        Because this flag consumes the remainder of the command line,
        the package list (if present) must appear before this flag.

    -c
        Compile the test binary to pkg.test in the current directory but do not run it
        (where pkg is the last element of the package's import path).
        The file name or target directory can be changed with the -o flag.

    -exec xprog
        Run the test binary using xprog. The behavior is the same as
        in 'go run'. See 'go help run' for details.

    -json
        Convert test output to JSON suitable for automated processing.
        See 'go doc test2json' for the encoding details.

    -o file
        Compile the test binary to the named file.
        The test still runs (unless -c or -i is specified).
        If file ends in a slash or names an existing directory,
        the test is written to pkg.test in that directory.
```

The test binary also accepts flags that control execution of the test; these flags are also accessible by ‘go test’. See ‘go help testflag’ for details.

​	测试二进制文件还接受控制测试执行的标志；这些标志也可以通过“go test” 访问。有关详细信息，请参阅“go help testflag”。

For more about build flags, see ‘go help build’.

​	有关构建标志的更多信息，请参阅“go help build”。

For more about specifying packages, see ‘go help packages’.

​	有关指定包的更多信息，请参阅“go help packages”。

See also: go build, go vet.

​	另请参阅：go build、go vet。