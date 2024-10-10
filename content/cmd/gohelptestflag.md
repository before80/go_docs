+++
title = "go help testflag"
date = 2023-05-17T09:59:21+08:00
type = "docs"
weight = 660
description = ""
isCJKLanguage = true
draft = false

+++

The 'go test' command takes both flags that apply to 'go test' itself and flags that apply to the resulting test binary.

​	`go test`命令既可以接受适用于 `go test`本身的标志，也可以接受适用于生成的测试二进制文件的标志。

Several of the flags control profiling and write an execution profile suitable for "go tool pprof"; run "go tool pprof -h" for more information. The --alloc_space, --alloc_objects, and --show_bytes options of pprof control how the information is presented.

​	其中一些标志用于控制性能分析，并生成适用于 "go tool pprof" 的执行分析文件；运行 "go tool pprof -h" 以获取更多信息。pprof 的 `--alloc_space`、`--alloc_objects` 和 `--show_bytes` 选项用于控制信息的显示方式。	

## `go test`命令可以识别并用于控制测试执行的标志

The following flags are recognized by the 'go test' command and control the execution of any test:

​	以下是 `go test`命令可以识别并用于控制测试执行的标志：

### -bench regexp    

​    Run only those benchmarks matching a regular expression.  By default, no benchmarks are run. To run all benchmarks, use '-bench .' or '-bench=.'.  The regular expression is split by unbracketed slash (/)  characters into a sequence of regular expressions, and each  part of a benchmark's identifier must match the corresponding element in the sequence, if any. Possible parents of matches  are run with b.N=1 to identify sub-benchmarks. For example,  given -bench=X/Y, top-level benchmarks matching X are run  with b.N=1 to find any sub-benchmarks matching Y, which are  then run in full.

​    仅运行与正则表达式匹配的基准测试。 默认情况下，不运行任何基准测试。要运行所有基准测试，使用 '-bench .' 或 '-bench=.'。
​    正则表达式会根据未包含在括号中的斜杠 (/) 字符进行拆分，形成一系列正则表达式，基准测试标识的每个部分必须与序列中的
​    相应元素匹配（如果有）。可能的父匹配会以 b.N=1 运行，以识别子基准测试。例如，给定 -bench=X/Y，将以 b.N=1 运行与 X 匹配的顶级基准测试，以查找与 Y 匹配的任何子基准测试，然后对其进行完整运行。

### -benchtime t      

​    Run enough iterations of each benchmark to take t, specified as a time.Duration (for example, -benchtime 1h30s). The default is 1 second (1s). The special syntax Nx means to run the benchmark N times  (for example, -benchtime 100x).

​    运行足够的基准测试迭代次数，以达到指定的时长 t，t以time.Duration形式指定（例如，-benchtime 1h30s）。
​    默认情况下，时长为 1 秒（1s）。特殊语法 Nx 表示运行基准测试 N 次（例如，-benchtime 100x）。

### -count n        

​    Run each test, benchmark, and fuzz seed n times (default 1). If -cpu is set, run n times for each GOMAXPROCS value. Examples are always run once. -count does not apply to fuzz tests matched by -fuzz.

​    运行每个测试、基准测试和模糊种子 n 次（默认为 1）。如果设置了 -cpu，则每个 GOMAXPROCS 值运行 n 次。示例总是运行一次。-count 不适用于由 -fuzz 匹配的模糊测试。

### -cover        

Enable coverage analysis.  Note that because coverage works by annotating the source code before compilation, compilation and test failures with  coverage enabled may report line numbers that don't correspond to the original sources.

​    启用覆盖率(coverage)分析。请注意，由于覆盖率分析通过在编译之前对源代码进行注释来工作，因此启用覆盖率分析的编译和测试失败可能会报告不对应于原始源代码的行号。

### -covermode set,count,atomic        

 Set the mode for coverage analysis for the package[s] being tested. The default is "set" unless -race is enabled,
    in which case it is "atomic".

​    The values:
​        set: bool: does this statement run?
​        count: int: how many times does this statement run?
​        atomic: int: count, but correct in multithreaded tests; significantly more expensive. Sets -cover.

​    为正在测试的 package[s] 设置覆盖率(coverage)分析模式。 默认为 "set"，除非启用了 -race，此时为 "atomic"。
​    这些值分别是：

​       set: bool：是否执行此语句？
​       count: int：此语句运行多少次？
​       atomic: int：计数，但在多线程测试中正确；代价较高。
​    设置 -cover。

### -coverpkg pattern1,pattern2,pattern3

  Apply coverage analysis in each test to packages matching the patterns. The default is for each test to analyze only the package being tested. See 'go help packages' for a description of package patterns.
    Sets -cover.

​    在每个测试中将覆盖率(coverage)分析应用于与模式匹配的 package。
​    默认情况下，每个测试仅分析正在测试的 package。
​    有关 package 模式的描述，请参阅 'go help packages'。
​    设置 -cover。
​    

### -cpu 1,2,4

​    Specify a list of GOMAXPROCS values for which the tests, benchmarks or fuzz tests should be executed. The default is the current value of GOMAXPROCS. -cpu does not apply to fuzz tests matched by -fuzz.

​    为测试、基准测试或模糊测试指定一组 GOMAXPROCS 值。
​    默认值为当前的 GOMAXPROCS 值。
​    -cpu 不适用于通过 -fuzz 匹配的模糊测试。

### -failfast

​    Do not start new tests after the first test failure.

​    在第一个测试失败后，不再启动新的测试。
​    

### -fullpath

​    Show full file names in the error messages.

### -fuzz regexp

​    Run the fuzz test matching the regular expression. When specified, the command line argument must match exactly one package within the main module, and regexp must match exactly one fuzz test within that package. Fuzzing will occur after tests, benchmarks, seed corpora of other fuzz tests, and examples have completed. See the Fuzzing section of the testing package documentation for details.

​    运行与正则表达式匹配的模糊测试。
​    在指定此标志时，命令行参数必须完全匹配主模块中的一个包，
​    并且 regexp 必须完全匹配该包内的一个模糊测试。
​    模糊测试将在测试、基准测试、其他模糊测试的种子语料库
​    以及示例完成后进行。
​    有关详细信息，请参阅testing包文档中的模糊测试部分。

### -fuzztime t

​    Run enough iterations of the fuzz target during fuzzing to take t, specified as a time.Duration (for example, -fuzztime 1h30s). The default is to run forever. The special syntax Nx means to run the fuzz target N times (for example, -fuzztime 1000x).

​    在模糊测试期间运行足够的模糊目标迭代，以达到指定的时长 t，
​    t以time.Duration形式指定（例如，-fuzztime 1h30s）。
​    默认值为永远运行。
​    特殊语法 Nx 表示运行模糊目标 N 次（例如，-fuzztime 1000x）。
​    

### -fuzzminimizetime t

​    Run enough iterations of the fuzz target during each minimization attempt to take t, as specified as a time.Duration (for example,  -fuzzminimizetime 30s). The default is 60s. The special syntax Nx means to run the fuzz target N times (for example, -fuzzminimizetime 100x).

​    在每次最小化尝试期间运行足够的模糊目标迭代，以达到指定的时长 t，
​    t以time.Duration形式指定（例如，-fuzzminimizetime 30s）。
​    默认为 60 秒。
​    特殊语法 Nx 表示运行模糊目标 N 次（例如，-fuzzminimizetime 100x）。

### -json

​    Log verbose output and test results in JSON. This presents the same information as the -v flag in a machine-readable format.

​    以 JSON 格式记录详细输出和测试结果。
​    这以机器可读的格式呈现与 -v 标志相同的信息。

### -list regexp

​    List tests, benchmarks, fuzz tests, or examples matching the regular expression. No tests, benchmarks, fuzz tests, or examples will be run. This will only list top-level tests. No subtest or subbenchmarks will be shown.

​    列出与正则表达式匹配的测试、基准测试、模糊测试或示例。
​    不运行任何测试、基准测试、模糊测试或示例。
​    仅列出顶级测试。不显示子测试或子基准测试。
​    

### -parallel n          

​    Allow parallel execution of test functions that call t.Parallel, and fuzz targets that call t.Parallel when running the seed corpus. The value of this flag is the maximum number of tests to run simultaneously.

While fuzzing, the value of this flag is the maximum number of subprocesses that may call the fuzz function simultaneously, regardless of  whether T.Parallel is called.

​    By default, -parallel is set to the value of GOMAXPROCS. Setting -parallel to values higher than GOMAXPROCS may cause degraded performance due to CPU contention, especially when fuzzing. Note that -parallel only applies within a single test binary. The 'go test' command may run tests for different packages in parallel as well, according to the setting of the -p flag (see 'go help build').

​    允许并行执行调用了 t.Parallel 的测试函数，
​    以及运行种子语料库时调用了 t.Parallel 的模糊目标。
​    该标志的值是要同时运行的最大测试数。
​    在模糊测试时，该标志的值是可能同时调用模糊函数的最大子进程数，
​    无论是否调用了 T.Parallel。
​    默认情况下，-parallel 设置为 GOMAXPROCS 的值。
​    将 -parallel 设置为高于 GOMAXPROCS 的值可能会导致性能降低，
​    因为 CPU 冲突，特别是在模糊测试时。
​    请注意，-parallel 仅在单个测试二进制文件中适用。
​    根据 -p 标志的设置（请参阅 'go help build'），
​    'go test' 命令也可以并行运行不同包的测试。

### -run regexp

​    Run only those tests, examples, and fuzz tests matching the regular expression. For tests, the regular expression is split by unbracketed slash (/) characters into a sequence of regular expressions, and each part of a test's identifier must match the corresponding element in the sequence, if any. Note that possible parents of matches are run too, so that -run=X/Y matches and runs and reports the result of all tests matching X, even those without sub-tests matching Y, because it must run them to look for those sub-tests.See also -skip.

​    仅运行与正则表达式匹配的测试、示例和模糊测试。
​    对于测试，正则表达式会通过未包含在括号中的斜杠（/）
​    字符拆分为一系列正则表达式，而测试的每个部分必须与序列中的
​    相应元素匹配（如果有的话）。请注意，可能的匹配父级也会运行，
​    因此 -run=X/Y 会匹配并运行 X 匹配的所有测试的结果，
​    即使没有子测试与 Y 匹配，因为必须运行它们以查找这些子测试。
​    另请参阅 -skip。
​    

### -short

​    Tell long-running tests to shorten their run time. It is off by default but set during all.bash so that installing the Go tree can run a sanity check but not spend time running exhaustive tests.

​    告诉长时间运行的测试缩短其运行时间。
​    默认情况下未启用，但在 all.bash 中设置为启用，
​    以便在运行全面测试时可以运行一次健全性检查，
​    而不会花费时间运行详尽测试。

### -shuffle off,on,N

​    Randomize the execution order of tests and benchmarks. It is off by default. If -shuffle is set to on, then it will seed the randomizer using the system clock. If -shuffle is set to an integer N, then N will be used as the seed value. In both cases, the seed will be reported for reproducibility.

​    随机化测试和基准测试的执行顺序。
​    默认情况下为off。如果将 -shuffle 设置为 on，
​    则它将使用系统时钟对随机生成器（randomizer）进行种子化。
​    如果将 -shuffle 设置为整数 N，
​    则 N 将用作种子值。在这两种情况下，种子将被报告以便进行重现。

### -skip regexp

​    Run only those tests, examples, fuzz tests, and benchmarks that do not match the regular expression. Like for -run and -bench, for tests and benchmarks, the regular expression is split by unbracketed slash (/) characters into a sequence of regular expressions, and each part of a test's identifier must match the corresponding element in the sequence, if any.

​    仅运行与正则表达式不匹配的测试、示例、模糊测试和基准测试。
​    与 -run 和 -bench 一样，对于测试和基准测试，
​    正则表达式会根据未括号化的斜杠 (/) 字符分成一系列正则表达式，
​    每个测试的标识符部分必须与序列中的相应元素匹配（如果有的话）。

### -timeout d

​    If a test binary runs longer than duration d, panic. If d is 0, the timeout is disabled. The default is 10 minutes (10m).

​    如果测试二进制运行时长超过 d，将触发 panic。
​    如果 d 为 0，则禁用超时。
​    默认值为 10 分钟（10m）。
​    

### -v

​    Verbose output: log all tests as they are run. Also print all text from Log and Logf calls even if the test succeeds.

​    详细输出：记录所有运行的测试。
​    即使测试成功，也会打印来自 Log 和 Logf 调用的所有文本。

### -vet list            

​    Configure the invocation of "go vet" during "go test" to use the comma-separated list of vet checks. If list is empty, "go test" runs "go vet" with a curated list of checks believed to be always worth addressing. If list is "off", "go test" does not run "go vet" at all.

​    配置在 "go test" 期间调用 "go vet" 
​    使用逗号分隔的 vet 检查列表。
​    如果列表为空，
​    则 "go test" 使用相信始终值得解决的检查的精选列表运行 "go vet"。
​    如果列表为 "off"，则 "go test" 根本不运行 "go vet"。



## `go test` 命令用于在执行过程中对测试进行分析的标志

The following flags are also recognized by 'go test' and can be used to profile the tests during execution:

​	`go test` 命令还识别以下标志，用于在执行过程中对测试进行分析：

### -benchmem

Print memory allocation statistics for benchmarks.

​    打印基准测试的内存分配统计信息。

### -blockprofile block.out

​    Write a goroutine blocking profile to the specified file
​    when all tests are complete.
​    Writes test binary as -c would.
​    在所有测试完成后，将 goroutine 阻塞分析写入指定的文件。
​    将测试二进制文件写入 -c 的目录。

### -blockprofilerate n       

Control the detail provided in goroutine blocking profiles by calling runtime.SetBlockProfileRate with n. See 'go doc runtime.SetBlockProfileRate'. The profiler aims to sample, on average, one blocking event every  n nanoseconds the program spends blocked. By default, if -test.blockprofile is set without this flag, all blocking events are recorded, equivalent to -test.blockprofilerate=1.

​    通过调用 runtime.SetBlockProfileRate 
​    设置 Goroutine 阻塞概要中提供的详细信息。
​    请参阅 'go doc runtime.SetBlockProfileRate'。
​    该分析器（profiler）的目标是平均每个程序被
​    阻塞的 n 纳秒时间内采样一次阻塞事件。
​    默认情况下，如果设置了 -test.blockprofile 但没有设置此标志，
​    则会记录所有阻塞事件，相当于 -test.blockprofilerate=1。
​        

### -coverprofile cover.out

Write a coverage profile to the file after all tests have passed. Sets -cover.

​    在所有测试通过后，将覆盖率（coverage）分析写入文件。设置 -cover。
​            

### -cpuprofile cpu.out

Write a CPU profile to the specified file before exiting. Writes test binary as -c would.

​    在退出之前，将 CPU 分析（profile）写入指定的文件。
​    将测试二进制文件写入 -c 的目录。

### -memprofile mem.out

Write an allocation profile to the file after all tests have passed. Writes test binary as -c would.

​    在所有测试通过后，将分配（allocation）分析写入文件。
​    将测试二进制文件写入 -c 的目录。

### -memprofilerate n

Enable more precise (and expensive) memory allocation profiles by setting runtime.MemProfileRate. See 'go doc runtime.MemProfileRate'. To profile all memory allocations, use -test.memprofilerate=1.

​    通过设置 runtime.MemProfileRate 来启用更精确（但更昂贵）
​    的内存分配分析。有关详细信息，
​    请参阅 'go doc runtime.MemProfileRate'。
​    若要分析所有内存分配，请使用 -test.memprofilerate=1。

### -mutexprofile mutex.out

Write a mutex contention profile to the specified file when all tests are complete. Writes test binary as -c would.

​    在所有测试完成后，将互斥锁争用（mutex contention）分析写入指定的文件。将测试二进制文件写入 -c 的目录。

### -mutexprofilefraction n

Sample 1 in n stack traces of goroutines holding a contended mutex.

​    对持有有争议的互斥锁的 goroutine 进行1:n的栈跟踪采样。

### -outputdir directory

​        Place output files from profiling in the specified directory,by default the directory in which "go test" is running.

​    	将性能分析的输出文件放置在指定的目录中， 默认情况下是"go test"正在运行的目录。
​        

### -trace trace.out

​        Write an execution trace to the specified file before exiting.

​    	在退出之前，将执行跟踪写入指定的文件。



Each of these flags is also recognized with an optional 'test.' prefix,as in -test.v. When invoking the generated test binary (the result of 'go test -c') directly, however, the prefix is mandatory.

​	这些标志也可以使用可选的 'test.' 前缀来识别，例如 `-test.v`。然而，当直接调用生成的测试二进制文件（由 '`go test -c`' 生成）时，前缀是必需的。

The `go test`command rewrites or removes recognized flags, as appropriate, both before and after the optional package list, before invoking the test binary.

​	`go test` 命令会在调用测试二进制文件之前和之后，根据需要在可选的包列表之前和之后重写或删除已识别的标志。

For instance, the command

​	例如，以下命令：

```cmd
    go test -v -myflag testdata -cpuprofile=prof.out -x
```

will compile the test binary and then run it as

将会编译测试二进制文件，然后作为以下方式运行：

```cmd
    pkg.test -test.v -myflag testdata -test.cpuprofile=prof.out
```

(The -x flag is removed because it applies only to the go command's execution, not to the test itself.)

（`-x` 标志已被移除，因为它仅适用于 go 命令的执行，不适用于测试本身。）

The test flags that generate profiles (other than for coverage) also leave the test binary in pkg.test for use when analyzing the profiles.

​	生成性能分析的测试标志（除了覆盖率分析）还会保留测试二进制文件 pkg.test，以供分析性能分析时使用。

When 'go test' runs a test binary, it does so from within the corresponding package's source code directory. Depending on the test, it may be necessary to do the same when invoking a generated test binary directly. Because that directory may be located within the module cache, which may be read-only and is verified by checksums, the test must not write to it or any other directory within the module unless explicitly requested by the user (such as with the -fuzz flag, which writes failures to testdata/fuzz).

​	当 `go test` 运行测试二进制文件时，它是从相应包的源代码目录内运行的。根据测试的不同，当直接调用生成的测试二进制文件时，可能需要做相同的操作。因为该目录可能位于模块缓存中，模块缓存可能是只读的，并且通过校验和进行验证，所以测试不能将其写入或写入模块内的任何其他目录，除非用户明确要求（例如，使用 `-fuzz` 标志，将失败写入 `testdata/fuzz`目录）。

The command-line package list, if present, must appear before any flag not known to the go test command. Continuing the example above, the package list would have to appear before -myflag, but could appear on either side of -v.

​	命令行 package 列表（如果存在）必须出现在 go test 命令不识别的任何标志之前。延续上面的示例，package 列表必须出现在 `-myflag` 之前，但可以出现在 `-v` 的任一侧。

When 'go test' runs in package list mode, 'go test' caches successful package test results to avoid unnecessary repeated running of tests. To disable test caching, use any test flag or argument other than the cacheable flags. The idiomatic way to disable test caching explicitly is to use -count=1.

​	当 'go test' 在 package 列表模式下运行时，它会缓存成功的包测试结果，以避免不必要地重复运行测试。要禁用测试缓存，请使用除了可缓存标志之外的任何测试标志或参数。显式禁用测试缓存的惯用方法是使用 `-count=1`。

To keep an argument for a test binary from being interpreted as a known flag or a package name, use -args (see 'go help test') which passes the remainder of the command line through to the test binary uninterpreted and unaltered.

​	为了防止测试二进制文件的参数被解释为已知标志或包名，请使用 `-args`（参见 'go help test'），该标志将命令行的其余部分不加解释地传递给测试二进制文件。

For instance, the command

​	例如，以下命令：

```cmd
go test -v -args -x -v
```

will compile the test binary and then run it as

将编译测试二进制文件，然后作为以下方式运行：

```cmd
pkg.test -test.v -x -v
```

Similarly,

类似地，

```cmd
go test -args math
```

will compile the test binary and then run it as

将编译测试二进制文件，然后作为以下方式运行：

```cmd
pkg.test math
```

In the first example, the -x and the second -v are passed through to the test binary unchanged and with no effect on the go command itself. In the second example, the argument math is passed through to the test binary, instead of being interpreted as the package list.

​	在第一个示例中，`-x` 和第二个 `-v` 被不加解释地传递给测试二进制文件，不会对 go 命令本身产生影响。在第二个示例中，参数 `math` 被不加解释地传递给测试二进制文件，而不是被解释为 package 列表。
