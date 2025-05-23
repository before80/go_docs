+++
title = "starlark"
date = 2024-12-09T07:59:50+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/cli/starlark.md](https://github.com/go-delve/delve/blob/master/Documentation/cli/starlark.md)
>
> 收录该文档时间： `2024-12-09T07:59:50+08:00`

# Introduction



Passing a file with the .star extension to the `source` command will cause delve to interpret it as a starlark script.

​	将扩展名为 `.star` 的文件传递给 `source` 命令，将导致 delve 将其解释为 Starlark 脚本。

Starlark is a dialect of python, a [specification of its syntax can be found here](https://github.com/google/starlark-go/blob/master/doc/spec.md).

​	Starlark 是 Python 的一个方言， [它的语法规范可以在这里找到](https://github.com/google/starlark-go/blob/master/doc/spec.md)。

In addition to the normal starlark built-ins delve defines [a number of global functions](https://github.com/go-delve/delve/blob/master/Documentation/cli/starlark.md#Starlark-built-ins) that can be used to interact with the debugger.

​	除了正常的 Starlark 内置函数，delve 还定义了 [一些全局函数](https://github.com/go-delve/delve/blob/master/Documentation/cli/starlark.md#Starlark-built-ins)，可以用来与调试器交互。

After the file has been evaluated delve will bind any function starting with `command_` to a command-line command: for example `command_goroutines_wait_reason` will be bound to `goroutines_wait_reason`.

​	在文件评估后，delve 会将所有以 `command_` 开头的函数绑定到命令行命令上：例如，`command_goroutines_wait_reason` 会绑定到 `goroutines_wait_reason`。

Then if a function named `main` exists it will be executed.

​	然后，如果存在名为 `main` 的函数，它将被执行。

Global functions with a name that begins with a capital letter will be available to other scripts.

​	以大写字母开头的全局函数将对其他脚本可用。

# Starlark 内置函数 Starlark built-ins



| Function                                                     | API Call                                                     |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| amend_breakpoint(Breakpoint)                                 | Equivalent to API call [AmendBreakpoint](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.AmendBreakpoint) |
| ancestors(GoroutineID, NumAncestors, Depth)                  | Equivalent to API call [Ancestors](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Ancestors) |
| attached_to_existing_process()                               | Equivalent to API call [AttachedToExistingProcess](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.AttachedToExistingProcess) |
| build_id()                                                   | Equivalent to API call [BuildID](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.BuildID) |
| cancel_next()                                                | Equivalent to API call [CancelNext](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.CancelNext) |
| checkpoint(Where)                                            | Equivalent to API call [Checkpoint](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Checkpoint) |
| clear_breakpoint(Id, Name)                                   | Equivalent to API call [ClearBreakpoint](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ClearBreakpoint) |
| clear_checkpoint(ID)                                         | Equivalent to API call [ClearCheckpoint](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ClearCheckpoint) |
| raw_command(Name, ThreadID, GoroutineID, ReturnInfoLoadConfig, Expr, UnsafeCall) | Equivalent to API call [Command](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Command) |
| create_breakpoint(Breakpoint, LocExpr, SubstitutePathRules, Suspended) | Equivalent to API call [CreateBreakpoint](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.CreateBreakpoint) |
| create_ebpf_tracepoint(FunctionName)                         | Equivalent to API call [CreateEBPFTracepoint](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.CreateEBPFTracepoint) |
| create_watchpoint(Scope, Expr, Type)                         | Equivalent to API call [CreateWatchpoint](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.CreateWatchpoint) |
| debug_info_directories(Set, List)                            | Equivalent to API call [DebugInfoDirectories](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.DebugInfoDirectories) |
| detach(Kill)                                                 | Equivalent to API call [Detach](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Detach) |
| disassemble(Scope, StartPC, EndPC, Flavour)                  | Equivalent to API call [Disassemble](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Disassemble) |
| dump_cancel()                                                | Equivalent to API call [DumpCancel](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.DumpCancel) |
| dump_start(Destination)                                      | Equivalent to API call [DumpStart](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.DumpStart) |
| dump_wait(Wait)                                              | Equivalent to API call [DumpWait](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.DumpWait) |
| eval(Scope, Expr, Cfg)                                       | Equivalent to API call [Eval](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Eval) |
| examine_memory(Address, Length)                              | Equivalent to API call [ExamineMemory](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ExamineMemory) |
| find_location(Scope, Loc, IncludeNonExecutableLines, SubstitutePathRules) | Equivalent to API call [FindLocation](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.FindLocation) |
| follow_exec(Enable, Regex)                                   | Equivalent to API call [FollowExec](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.FollowExec) |
| follow_exec_enabled()                                        | Equivalent to API call [FollowExecEnabled](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.FollowExecEnabled) |
| function_return_locations(FnName)                            | Equivalent to API call [FunctionReturnLocations](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.FunctionReturnLocations) |
| get_breakpoint(Id, Name)                                     | Equivalent to API call [GetBreakpoint](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.GetBreakpoint) |
| get_buffered_tracepoints()                                   | Equivalent to API call [GetBufferedTracepoints](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.GetBufferedTracepoints) |
| get_thread(Id)                                               | Equivalent to API call [GetThread](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.GetThread) |
| guess_substitute_path(Args)                                  | Equivalent to API call [GuessSubstitutePath](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.GuessSubstitutePath) |
| is_multiclient()                                             | Equivalent to API call [IsMulticlient](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.IsMulticlient) |
| last_modified()                                              | Equivalent to API call [LastModified](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.LastModified) |
| breakpoints(All)                                             | Equivalent to API call [ListBreakpoints](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListBreakpoints) |
| checkpoints()                                                | Equivalent to API call [ListCheckpoints](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListCheckpoints) |
| dynamic_libraries()                                          | Equivalent to API call [ListDynamicLibraries](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListDynamicLibraries) |
| function_args(Scope, Cfg)                                    | Equivalent to API call [ListFunctionArgs](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListFunctionArgs) |
| functions(Filter, FollowCalls)                               | Equivalent to API call [ListFunctions](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListFunctions) |
| goroutines(Start, Count, Filters, GoroutineGroupingOptions, EvalScope) | Equivalent to API call [ListGoroutines](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListGoroutines) |
| local_vars(Scope, Cfg)                                       | Equivalent to API call [ListLocalVars](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListLocalVars) |
| package_vars(Filter, Cfg)                                    | Equivalent to API call [ListPackageVars](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListPackageVars) |
| packages_build_info(IncludeFiles, Filter)                    | Equivalent to API call [ListPackagesBuildInfo](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListPackagesBuildInfo) |
| registers(ThreadID, IncludeFp, Scope)                        | Equivalent to API call [ListRegisters](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListRegisters) |
| sources(Filter)                                              | Equivalent to API call [ListSources](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListSources) |
| targets()                                                    | Equivalent to API call [ListTargets](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListTargets) |
| threads()                                                    | Equivalent to API call [ListThreads](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListThreads) |
| types(Filter)                                                | Equivalent to API call [ListTypes](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ListTypes) |
| process_pid()                                                | Equivalent to API call [ProcessPid](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ProcessPid) |
| recorded()                                                   | Equivalent to API call [Recorded](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Recorded) |
| restart(Position, ResetArgs, NewArgs, Rerecord, Rebuild, NewRedirects) | Equivalent to API call [Restart](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Restart) |
| set_expr(Scope, Symbol, Value)                               | Equivalent to API call [Set](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Set) |
| stacktrace(Id, Depth, Full, Defers, Opts, Cfg)               | Equivalent to API call [Stacktrace](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Stacktrace) |
| state(NonBlocking)                                           | Equivalent to API call [State](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.State) |
| toggle_breakpoint(Id, Name)                                  | Equivalent to API call [ToggleBreakpoint](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.ToggleBreakpoint) |
| dlv_command(command)                                         | Executes the specified command as if typed at the dlv_prompt |
| read_file(path)                                              | Reads the file as a string                                   |
| write_file(path, contents)                                   | Writes string to a file                                      |
| cur_scope()                                                  | Returns the current evaluation scope                         |
| default_load_config()                                        | Returns the current default load configuration               |

In addition to these built-ins, the [time](https://pkg.go.dev/go.starlark.net/lib/time#pkg-variables) library from the starlark-go project is also available to scripts.

​	除了这些内置函数外，starlark-go 项目的 [time](https://pkg.go.dev/go.starlark.net/lib/time#pkg-variables) 库也可以在脚本中使用。

## 我应该使用 raw_command 还是 dlv_command？Should I use raw_command or dlv_command?



There are two ways to resume the execution of the target program:

​	有两种方法可以恢复目标程序的执行：

```
raw_command("continue")
dlv_command("continue")
```



The first one maps to the API call [Command](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Command). As such all the caveats explained in the [Client HowTo](https://github.com/go-delve/delve/blob/master/Documentation/api/ClientHowto.md).

​	第一个方法映射到 API 调用 [Command](https://pkg.go.dev/github.com/go-delve/delve/service/rpc2#RPCServer.Command)。因此，所有在 [Client HowTo](https://github.com/go-delve/delve/blob/master/Documentation/api/ClientHowto.md) 中解释的注意事项都适用。

The latter is equivalent to typing `continue` to the `(dlv)` command line and should do what you expect.

​	后者相当于在 `(dlv)` 命令行中输入 `continue`，应该能够按预期执行。

In general `dlv_command("continue")` should be preferred, unless the behavior you wish to produces diverges significantly from that of the command line's `continue`.

​	一般来说，应该首选 `dlv_command("continue")`，除非你希望的行为与命令行中的 `continue` 显著不同。

# 创建新命令 Creating new commands



Any global function with a name starting with `command_` will be made available as a command line command. If the function has a single argument named `args` all arguments passed on the command line will be passed to the function as a single string.

​	任何名称以 `command_` 开头的全局函数将作为命令行命令提供。如果该函数有一个名为 `args` 的单个参数，则命令行上传递的所有参数将作为一个字符串传递给该函数。

Otherwise arguments passed on the command line are interpreted as starlark expressions. See the [expression arguments](https://github.com/go-delve/delve/blob/master/Documentation/cli/starlark.md#expression-arguments) example.

​	否则，命令行上传递的参数将被解释为 Starlark 表达式。请参阅 [expression arguments](https://github.com/go-delve/delve/blob/master/Documentation/cli/starlark.md#expression-arguments) 示例。

If the command function has a doc string it will be used as a help message.

​	如果命令函数有文档字符串，它将作为帮助信息使用。

# 使用变量 Working with variables



Variables of the target program can be accessed using `local_vars`, `function_args` or the `eval` functions. Each variable will be returned as a [Variable](https://pkg.go.dev/github.com/go-delve/delve/service/api#Variable) struct, with one special field: `Value`.

​	可以通过 `local_vars`、`function_args` 或 `eval` 函数访问目标程序的变量。每个变量将作为一个 [Variable](https://pkg.go.dev/github.com/go-delve/delve/service/api#Variable) 结构返回，其中有一个特殊字段：`Value`。

## Variable.Value



The `Value` field will return the value of the target variable converted to a starlark value:

​	`Value` 字段将返回目标变量的值，并将其转换为 Starlark 值：

- integers, floating point numbers and strings are represented by equivalent starlark values
  - 整数、浮点数和字符串由相应的 Starlark 值表示

- structs are represented as starlark dictionaries
  - 结构体由 Starlark 字典表示

- slices and arrays are represented by starlark lists
  - 切片和数组由 Starlark 列表表示

- maps are represented by starlark dicts
  - 映射由 Starlark 字典表示

- pointers and interfaces are represented by a one-element starlark list containing the value they point to
  - 指针和接口由一个包含它们指向值的单元素 Starlark 列表表示


For example, given this variable in the target program:

​	例如，给定目标程序中的以下变量：

```
type astruct struct {
	A int
	B int
}

s2 := []astruct{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}, {13, 14}, {15, 16}}
```



The following is possible:

​	以下操作是可能的：

```
>>> s2 = eval(None, "s2").Variable
>>> s2.Value[0]                                     # access of a slice item by index 通过索引访问切片项
main.astruct {A: 1, B: 2}
>>> a = s2.Value[1]
>>> a.Value.A                                       # access to a struct field 访问结构体字段
3
>>> a.Value.A + 10                            # arithmetic on the value of s2[1].X 对 s2[1].X 的算术运算
13
>>> a.Value["B"]                                    # access to a struct field, using dictionary syntax 使用字典语法访问结构体字段
4
```



For more examples see the [linked list example](https://github.com/go-delve/delve/blob/master/Documentation/cli/starlark.md#Print-all-elements-of-a-linked-list) below.

​	更多示例请参见 [链表示例](https://github.com/go-delve/delve/blob/master/Documentation/cli/starlark.md#Print-all-elements-of-a-linked-list)。

# Examples



## 列出 goroutines 并创建自定义命令 Listing goroutines and making custom commands



Create a `goroutine_start_line` command that prints the starting line of each goroutine, sets `gsl` as an alias:

​	创建一个 `goroutine_start_line` 命令，打印每个 goroutine 的起始行，并将 `gsl` 设置为别名：

```
def command_goroutine_start_line(args):
	gs = goroutines().Goroutines
	for g in gs:
		line = read_file(g.StartLoc.File).splitlines()[g.StartLoc.Line-1].strip()
		print(g.ID, "\t", g.StartLoc.File + ":" + str(g.StartLoc.Line), "\t", line)

def main():
	dlv_command("config alias goroutine_start_line gsl")
```



Use it like this:

​	用法如下：

```
(dlv) source goroutine_start_line.star
(dlv) goroutine_start_line
1 	 /usr/local/go/src/runtime/proc.go:110 	 func main() {
2 	 /usr/local/go/src/runtime/proc.go:242 	 func forcegchelper() {
17 	 /usr/local/go/src/runtime/mgcsweep.go:64 	 func bgsweep(c chan int) {
18 	 /usr/local/go/src/runtime/mfinal.go:161 	 func runfinq() {
(dlv) gsl
1 	 /usr/local/go/src/runtime/proc.go:110 	 func main() {
2 	 /usr/local/go/src/runtime/proc.go:242 	 func forcegchelper() {
17 	 /usr/local/go/src/runtime/mgcsweep.go:64 	 func bgsweep(c chan int) {
18 	 /usr/local/go/src/runtime/mfinal.go:161 	 func runfinq() {
```



## 表达式参数 Expression arguments



After evaluating this script:

​	评估以下脚本后：

```
def command_echo(args):
	print(args)

def command_echo_expr(a, b, c):
	print("a", a, "b", b, "c", c)
```



The first command, `echo`, takes its arguments as a single string, while for `echo_expr` it will be possible to pass starlark expression as arguments:

​	第一个命令 `echo` 将其参数作为一个字符串接收，而 `echo_expr` 则可以作为参数传递 Starlark 表达式：

```
(dlv) echo 2+2, 2-1, 2*3
"2+2, 2-1, 2*3"
(dlv) echo_expr 2+2, 2-1, 2*3
a 4 b 1 c 6
```



## 创建断点 Creating breakpoints



Set a breakpoint on all private methods of package `main`:

​	为 `main` 包中的所有私有方法设置断点：

```
def main():
	for f in functions().Funcs:
		v = f.split('.')
		if len(v) != 2:
			continue
		if v[0] != "main":
			continue
		if v[1][0] >= 'a' and v[1][0] <= 'z':
			create_breakpoint({ "FunctionName": f, "Line": -1 }) # see documentation of RPCServer.CreateBreakpoint
```



## 切换 goroutines - Switching goroutines



Create a command, `switch_to_main_goroutine`, that searches for a goroutine running a function in the main package and switches to it:

​	创建一个命令 `switch_to_main_goroutine`，它搜索正在运行 `main` 包中函数的 goroutine 并切换到它：

```
def command_switch_to_main_goroutine(args):
	for g in goroutines().Goroutines:
		if g.currentLoc.function != None and g.currentLoc.function.name.startswith("main."):
			print("switching to:", g.id)
			raw_command("switchGoroutine", GoroutineID=g.id)
			break
```



## 列出 goroutines - Listing goroutines



Create a command, "goexcl", that lists all goroutines excluding the ones stopped on a specified function.

​	创建一个命令 `goexcl`，它列出所有 goroutines，排除在指定函数中停止的 goroutines。

```
def command_goexcl(args):
	"""Prints all goroutines not stopped in the function passed as argument."""
	excluded = 0
	start = 0
	while start >= 0:
		gr = goroutines(start, 10)
		start = gr.Nextg
		for g in gr.Goroutines:
			fn = g.UserCurrentLoc.Function
			if fn == None:
				print("Goroutine", g.ID, "User:", g.UserCurrentLoc.File, g.UserCurrentLoc.Line)
			elif fn.Name_ != args:
				print("Goroutine", g.ID, "User:", g.UserCurrentLoc.File, g.UserCurrentLoc.Line, fn.Name_)
			else:
				excluded = excluded + 1
	print("Excluded", excluded, "goroutines")
```



Usage:

​	用法：

```
(dlv) goexcl main.somefunc
```



prints all goroutines that are not stopped inside `main.somefunc`.

​	打印所有未在 `main.somefunc` 中停止的 goroutines。

## 反复执行目标，直到命中断点 Repeatedly executing the target until a breakpoint is hit.



Repeatedly call continue and restart until the target hits a breakpoint.

​	反复调用 `continue` 和 `restart`，直到目标命中断点。

```
def command_flaky(args):
	"Repeatedly runs program until a breakpoint is hit"
	while True:
		if dlv_command("continue") == None:
			break
		dlv_command("restart")
```



## 打印链表中的所有元素 Print all elements of a linked list



```
def command_linked_list(args):
	"""Prints the contents of a linked list.
	
	linked_list <var_name> <next_field_name> <max_depth>

Prints up to max_depth elements of the linked list variable 'var_name' using 'next_field_name' as the name of the link field.
"""
	var_name, next_field_name, max_depth = args.split(" ")
	max_depth = int(max_depth)
	next_name = var_name
	v = eval(None, var_name).Variable.Value
	for i in range(0, max_depth):
		print(str(i)+":",v)
		if v[0] == None:
			break
		v = v[next_field_name]
```



## 查找匹配谓词的数组元素 Find an array element matching a predicate



```
def command_find_array(arr, pred):
	"""Calls pred for each element of the array or slice 'arr' returns the index of the first element for which pred returns true.
	
	find_arr <arr> <pred>
	
Example use (find the first element of slice 's2' with field A equal to 5):
	
	find_arr "s2", lambda x: x.A == 5
"""
	arrv = eval(None, arr).Variable
	for i in range(0, arrv.Len):
		v = arrv.Value[i]
		if pred(v):
			print("found", i)
			return

	print("not found")
```



## 重新运行程序，直到失败或命中断点 Rerunning a program until it fails or hits a breakpoint



```
def command_flaky(args):
	"Continues and restarts the target program repeatedly (re-recording it on the rr backend), until a breakpoint is hit"
	count = 1
	while True:
		if dlv_command("continue") == None:
			break
		print("restarting", count, "...")
		count = count+1
		restart(Rerecord=True)
```



## 作为参数传递结构体 Passing a struct as an argument



Struct literals can be passed to built-ins as Starlark dictionaries. For example, the following snippet passes in an [api.EvalScope](https://pkg.go.dev/github.com/go-delve/delve/service/api#EvalScope) and [api.LoadConfig](https://pkg.go.dev/github.com/go-delve/delve/service/api#LoadConfig) to the `eval` built-in. `None` can be passed for optional arguments, and trailing optional arguments can be elided completely.

​	结构体字面量可以作为 Starlark 字典传递给内置函数。例如，以下代码将 [api.EvalScope](https://pkg.go.dev/github.com/go-delve/delve/service/api#EvalScope) 和 [api.LoadConfig](https://pkg.go.dev/github.com/go-delve/delve/service/api#LoadConfig) 传递给 `eval` 内置函数。`None` 可以传递给可选参数，尾随的可选参数可以完全省略。

```python
var = eval(
        {"GoroutineID": 42, "Frame": 5},
        "myVar",
        {"FollowPointers":True, "MaxVariableRecurse":2, "MaxStringLen":100, "MaxArrayValues":10, "MaxStructFields":100}
      )
```
