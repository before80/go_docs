+++
title = "cli"
date = 2024-12-09T07:58:54+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md)
>
> 收录该文档时间： `2024-12-09T07:58:54+08:00`


# Configuration and Command History - 配置和命令历史

If `$XDG_CONFIG_HOME` is set, then configuration and command history files are located in `$XDG_CONFIG_HOME/dlv`. Otherwise, they are located in `$HOME/.config/dlv` on Linux and `$HOME/.dlv` on other systems.

​	如果设置了 `$XDG_CONFIG_HOME`，则配置和命令历史文件位于 `$XDG_CONFIG_HOME/dlv`。否则，在 Linux 系统上，它们位于 `$HOME/.config/dlv`，在其他系统上则位于 `$HOME/.dlv`。

The configuration file `config.yml` contains all the configurable options and their default values. The command history is stored in `.dbg_history`.

​	配置文件 `config.yml` 包含所有可配置选项及其默认值。命令历史存储在 `.dbg_history` 文件中。

# Commands

## 运行程序 Running the program

Command | Description
--------|------------
[call](#call) | 恢复进程，注入函数调用（实验性!!!） Resumes process, injecting a function call (EXPERIMENTAL!!!) 
[continue](#continue) | 运行直到遇到断点或程序终止。 Run until breakpoint or program termination. 
[next](#next) | 执行下一行源码。 Step over to next source line. 
[next-instruction](#next-instruction) | 单步执行单个 CPU 指令，跳过函数调用。 Single step a single cpu instruction, skipping function calls. 
[rebuild](#rebuild) | 重新构建目标可执行文件并重新启动。如果可执行文件不是由 delve 构建的，则无法使用此命令。 Rebuild the target executable and restarts it. It does not work if the executable was not built by delve. 
[restart](#restart) | 重启进程。 Restart process. 
[rev](#rev) | 反转目标程序的执行，按指定命令执行。 Reverses the execution of the target program for the command specified. 
[rewind](#rewind) | 向后执行，直到遇到断点或记录的历史开始。 Run backwards until breakpoint or start of recorded history. 
[step](#step) | 单步执行程序。 Single step through program. 
[step-instruction](#step-instruction) | 单步执行单个 CPU 指令。 Single step a single cpu instruction. 
[stepout](#stepout) | 从当前函数中跳出。 Step out of the current function. 


## 操作断点 Manipulating breakpoints

Command | Description
--------|------------
[break](#break) | 设置一个断点。 Sets a breakpoint. 
[breakpoints](#breakpoints) | 打印活动断点的信息。 Print out info for active breakpoints. 
[clear](#clear) | 删除断点。 Deletes breakpoint. 
[clearall](#clearall) | 删除多个断点。 Deletes multiple breakpoints. 
[condition](#condition) | 设置断点条件。 Set breakpoint condition. 
[on](#on) | 在断点命中时执行命令。 Executes a command when a breakpoint is hit. 
[toggle](#toggle) | 切换断点的开启或关闭状态。 Toggles on or off a breakpoint. 
[trace](#trace) | 设置跟踪点。 Set tracepoint. 
[watch](#watch) | 设置观察点。Set watchpoint. 


## 查看程序变量和内存 Viewing program variables and memory

Command | Description
--------|------------
[args](#args) | 打印函数参数。  Print function arguments. 
[display](#display) | 每次程序停止时打印表达式的值。 Print value of an expression every time the program stops. 
[examinemem](#examinemem) | 检查给定地址处的原始内存。 Examine raw memory at the given address. 
[locals](#locals) | 打印局部变量。 Print local variables. 
[print](#print) | 计算表达式的值并打印。Evaluate an expression. 
[regs](#regs) | 打印 CPU 寄存器的内容。Print contents of CPU registers. 
[set](#set) | 修改变量的值。Changes the value of a variable. 
[vars](#vars) | 打印包变量。Print package variables. 
[whatis](#whatis) | 打印表达式的类型。Prints type of an expression. 


## 列出和切换线程和 goroutine - Listing and switching between threads and goroutines

Command | Description
--------|------------
[goroutine](#goroutine) | 显示或切换当前的 goroutine。Shows or changes current goroutine 
[goroutines](#goroutines) | 列出程序中的所有 goroutine。List program goroutines. 
[thread](#thread) | 切换到指定线程。Switch to the specified thread. 
[threads](#threads) | 打印每个被追踪线程的信息。Print out info for every traced thread. 


## 查看调用栈并选择栈帧 Viewing the call stack and selecting frames

Command | Description
--------|------------
[deferred](#deferred) | 在延迟调用的上下文中执行命令。Executes command in the context of a deferred call. 
[down](#down) | 向下移动当前栈帧。Move the current frame down. 
[frame](#frame) | 设置当前栈帧，或在不同栈帧上执行命令。Set the current frame, or execute command on a different frame. 
[stack](#stack) | 打印调用栈跟踪。Print stack trace. 
[up](#up) | 向上移动当前栈帧。Move the current frame up. 


## 其他命令 Other commands

Command | Description
--------|------------
[check](#check) | 在当前的位置创建一个检查点。Creates a checkpoint at the current position. 
[checkpoints](#checkpoints) | 打印现有检查点的信息。Print out info for existing checkpoints. 
[clear-checkpoint](#clear-checkpoint) | 删除检查点。Deletes checkpoint. 
[config](#config) | 修改配置参数。Changes configuration parameters. 
[disassemble](#disassemble) | 反汇编。Disassembler. 
[dump](#dump) | 从当前进程状态创建核心转储。Creates a core dump from the current process state 
[edit](#edit) | 在 `$DELVE_EDITOR` 或 `$EDITOR` 中打开当前文件。Open where you are in `$DELVE_EDITOR `or `$EDITOR` 
[exit](#exit) | 退出调试器。Exit the debugger. 
[funcs](#funcs) | 打印函数列表。Print list of functions. 
[help](#help) | 打印帮助信息。Prints the help message. 
[libraries](#libraries) | 列出已加载的动态库。List loaded dynamic libraries 
[list](#list) | 显示源代码。Show source code. 
[packages](#packages) | 打印包列表。Print list of packages. 
[source](#source) | 执行包含一系列 delve 命令的文件。Executes a file containing a list of delve commands 
[sources](#sources) | 打印源文件列表。Print list of source files. 
[target](#target) | 管理子进程调试。Manages child process debugging. 
[transcript](#transcript) | 将命令输出追加到文件中。Appends command output to a file. 
[types](#types) | 打印类型列表。Print list of types 

## args
Print function arguments.

​	打印函数参数。

	[goroutine <n>] [frame <m>] args [-v] [<regex>]

If regex is specified only function arguments with a name matching it will be returned. If -v is specified more information about each function argument will be shown.

​	如果指定了 regex，只有与之匹配的函数参数名称才会被返回。如果指定了 -v，会显示更多关于每个函数参数的信息。


## break
Sets a breakpoint.

​	设置断点。

	break [name] [locspec] [if <condition>]

Locspec is a location specifier in the form of:

​	Locspec 是一种位置说明符，格式如下：

  * *&lt;address> Specifies the location of memory address address. address can be specified as a decimal, hexadecimal or octal number
      * *<address> 指定内存地址 address 的位置。address 可以是十进制、十六进制或八进制数字。

  * &lt;filename>:&lt;line> Specifies the line in filename. filename can be the partial path to a file or even just the base name as long as the expression remains unambiguous.
      * <filename>:<line> 指定文件 filename 中的行号。filename 可以是文件的部分路径，甚至只需基文件名，只要表达式不产生歧义。

  * &lt;line> Specifies the line in the current file
      * <line> 指定当前文件中的行号。

  * +&lt;offset> Specifies the line offset lines after the current one
      * +<offset> 指定当前行之后的偏移行数。

  * -&lt;offset> Specifies the line offset lines before the current one
      * -<offset> 指定当前行之前的偏移行数。

  * &lt;function>[:&lt;line>] Specifies the line inside function.
      The full syntax for function is &lt;package>.(*&lt;receiver type>).&lt;function name> however the only required element is the function name,
      everything else can be omitted as long as the expression remains unambiguous. For setting a breakpoint on an init function (ex: main.init),
      the &lt;filename>:&lt;line> syntax should be used to break in the correct init function at the correct location.
      * <function>[:<line>] 指定函数内部的行号。 函数的完整语法为 <package>.(*<receiver type>).<function name>，不过仅需提供函数名，其它部分可以省略，只要表达式不产生歧义。如果需要在 init 函数上设置断点（例如：main.init），应使用 <filename>:<line> 语法来断定正确的 init 函数位置。

  * /&lt;regex>/ Specifies the location of all the functions matching regex
      * /<regex>/ 指定匹配 regex 的所有函数位置。


If locspec is omitted a breakpoint will be set on the current line.

​	如果省略 locspec，则会在当前行设置断点。

If you would like to assign a name to the breakpoint you can do so with the form:

​	如果要为断点命名，可以使用以下形式：

	break mybpname main.go:4

Finally, you can assign a condition to the newly created breakpoint by using the 'if' postfix form, like so:

​	最后，您可以使用 'if' 后缀为新创建的断点指定条件，如下所示：

	break main.go:55 if i == 5

Alternatively you can set a condition on a breakpoint after created by using the 'on' command.

​	或者，也可以在断点创建后，使用 'on' 命令设置条件。

See also: "help on", "help cond" and "help clear"

​	另请参阅： "help on"、"help cond" 和 "help clear"

Aliases: b

​	别名： b

## breakpoints
Print out info for active breakpoints.

​	打印活动断点的信息。	

	breakpoints [-a]

Specifying -a prints all physical breakpoint, including internal breakpoints.

​	指定 -a 会打印所有物理断点，包括内部断点。

Aliases: bp

​	别名： bp

## call
Resumes process, injecting a function call (EXPERIMENTAL!!!)

​		恢复进程，注入函数调用（实验性!!!）

	call [-unsafe] <function call expression>

Current limitations:

​	当前的限制：

- only pointers to stack-allocated objects can be passed as argument.
  - 只能传递指向栈分配对象的指针作为参数。

- only some automatic type conversions are supported.
  - 仅支持一些自动类型转换。

- functions can only be called on running goroutines that are not executing the runtime.
  - 只能在未执行运行时的运行中的 goroutine 上调用函数。

- the current goroutine needs to have at least 256 bytes of free space on the stack.
  - 当前的 goroutine 需要至少 256 字节的栈空间。

- functions can only be called when the goroutine is stopped at a safe point.
  - 只能在 goroutine 停止于安全点时调用函数。

- calling a function will resume execution of all goroutines.
  - 调用函数会恢复所有 goroutine 的执行。

- only supported on linux's native backend.
  - 仅支持 Linux 本地后端。




## check
Creates a checkpoint at the current position.

​	在当前的位置创建一个检查点。

	checkpoint [note]

The "note" is arbitrary text that can be used to identify the checkpoint, if it is not specified it defaults to the current filename:line position.

​	检查点是您调试会话中的标记。可以创建一个检查点并将其附加到目标程序的进程中，并在以后恢复到该点。

Aliases: checkpoint

​	**别名**: checkpoint

## checkpoints
Print out info for existing checkpoints.

​	打印现有检查点的信息。


## clear
Deletes breakpoint.

​	删除断点。

	clear <breakpoint name or id>


## clear-checkpoint
Deletes checkpoint.

​	删除检查点。

	clear-checkpoint <id>

Aliases: clearcheck

​	别名: clearcheck

## clearall
Deletes multiple breakpoints.

​	删除多个断点。

	clearall [<locspec>]

If called with the locspec argument it will delete all the breakpoints matching the locspec. If locspec is omitted all breakpoints are deleted.

​	如果带有locspec参数，则删除所有匹配该位置规范的断点。如果省略locspec，则删除所有断点。


## condition
Set breakpoint condition.

​	设置断点条件。

	condition <breakpoint name or id> <boolean expression>.
		condition <断点名称或ID> <布尔表达式>。
		
	condition -hitcount <breakpoint name or id> <operator> <argument>.
		condition -hitcount <断点名称或ID> <操作符> <参数>。
		
	condition -per-g-hitcount <breakpoint name or id> <operator> <argument>.
		condition -per-g-hitcount <断点名称或ID> <操作符> <参数>。
		
	condition -clear <breakpoint name or id>.
		condition -clear <断点名称或ID>。

Specifies that the breakpoint, tracepoint or watchpoint should break only if the boolean expression is true.

​	指定断点、追踪点或观察点只有在布尔表达式为真时才会中断。

See [Documentation/cli/expr.md](//github.com/go-delve/delve/tree/master/Documentation/cli/expr.md) for a description of supported expressions.

​	请参阅 [Documentation/cli/expr.md](http://github.com/go-delve/delve/tree/master/Documentation/cli/expr.md) 以了解支持的表达式描述。

With the `-hitcount` option a condition on the breakpoint hit count can be set, the following operators are supported

​	使用 `-hitcount` 选项，可以设置断点命中次数的条件，支持以下操作符：

	condition -hitcount bp > n
	condition -hitcount bp >= n
	condition -hitcount bp < n
	condition -hitcount bp <= n
	condition -hitcount bp == n
	condition -hitcount bp != n
	condition -hitcount bp % n

The `-per-g-hitcount` option works like -hitcount, but use per goroutine hitcount to compare with n.

​	`-per-g-hitcount` 选项的工作方式与 `-hitcount` 类似，但使用每个 goroutine 的 hitcount 与 n 进行比较。

With the `-clear` option a condition on the breakpoint can removed.

​	使用 `-clear` 选项可以移除断点上的条件。	
The '% n' form means we should stop at the breakpoint when the hitcount is a multiple of n.

​	`'% n'` 形式表示当 hitcount 是 n 的倍数时，我们应该在断点处停止。

Examples:

	cond 2 i == 10				breakpoint 2 will stop when variable i equals 10
									当变量 i 等于 10 时，断点 2 会停止
									
	cond name runtime.curg.goid == 5	breakpoint 'name' will stop only on goroutine 5
											断点 'name' 只会在 goroutine 5 停止
											
	cond -clear 2				the condition on breakpoint 2 will be removed
									断点 2 上的条件会被移除

Aliases: cond

​	别名：cond

## config
Changes configuration parameters.

​	更改配置参数。

	config -list

Show all configuration parameters.

​	显示所有配置参数。

	config -save

Saves the configuration file to disk, overwriting the current configuration file.

​	将配置文件保存到磁盘，覆盖当前的配置文件。

	config <parameter> <value>

Changes the value of a configuration parameter.

​	更改配置参数的值。

	config substitute-path <from> <to>
	config substitute-path <from>
	config substitute-path -clear
	config substitute-path -guess

Adds or removes a path substitution rule, if -clear is used all substitute-path rules are removed. Without arguments shows the current list of substitute-path rules. The -guess option causes Delve to try to guess your substitute-path configuration automatically. See also [Documentation/cli/substitutepath.md](//github.com/go-delve/delve/tree/master/Documentation/cli/substitutepath.md) for how the rules are applied.

​	添加或移除路径替换规则。如果使用 `-clear`，则会移除所有 substitute-path 规则。如果没有参数，则显示当前的 substitute-path 规则列表。 `-guess` 选项会让 Delve 尝试自动猜测你的 substitute-path 配置。 有关规则应用的更多信息，请参见 [Documentation/cli/substitutepath.md](http://github.com/go-delve/delve/tree/master/Documentation/cli/substitutepath.md)。

	config alias <command> <alias>
	config alias <alias>

Defines &lt;alias> as an alias to &lt;command> or removes an alias.

​	将 `<alias>` 定义为 `<command>` 的别名，或者移除一个别名。

	config debug-info-directories -add <path>
	config debug-info-directories -rm <path>
	config debug-info-directories -clear

Adds, removes or clears debug-info-directories.

​	添加、移除或清除 debug-info-directories。


## continue
Run until breakpoint or program termination.

​	运行直到断点或程序终止。

	continue [<locspec>]

Optional locspec argument allows you to continue until a specific location is reached. The program will halt if a breakpoint is hit before reaching the specified location.

​	可选的 locspec 参数允许你继续执行直到到达特定位置。如果在到达指定位置之前触发了断点，程序会暂停。

For example:

​	例如：

	continue main.main
	continue encoding/json.Marshal

Aliases: c

​	别名：c

## deferred
Executes command in the context of a deferred call.

​	在延迟调用的上下文中执行命令。

	deferred <n> <command>

Executes the specified command (print, args, locals) in the context of the n-th deferred call in the current frame.

​	在当前帧的第 n 个延迟调用的上下文中执行指定的命令（打印、参数、局部变量等）。


## disassemble
Disassembler.

​	反汇编。

	[goroutine <n>] [frame <m>] disassemble [-a <start> <end>] [-l <locspec>]

If no argument is specified the function being executed in the selected stack frame will be executed.

​	如果没有指定参数，将反汇编当前选中堆栈帧中正在执行的函数。

	-a <start> <end>	disassembles the specified address range
							反汇编指定的地址范围
							
	-l <locspec>		disassembles the specified function
							反汇编指定的函数

Aliases: disass

​	别名：disass

## display
Print value of an expression every time the program stops.

​	每次程序停止时打印表达式的值。

	display -a [%format] <expression>
	display -d <number>

The '-a' option adds an expression to the list of expression printed every time the program stops. The '-d' option removes the specified expression from the list.

​	`-a` 选项会将一个表达式添加到每次程序停止时打印的表达式列表中。`-d` 选项会移除指定的表达式。

If display is called without arguments it will print the value of all expression in the list.

​	如果不带参数调用 display，它会打印所有表达式的值。


## down
Move the current frame down.

​	将当前帧向下移动。

	down [<m>]
	down [<m>] <command>

Move the current frame down by &lt;m>. The second form runs the command on the given frame.

​	将当前帧向下移动 `<m>` 次。第二种形式会在给定帧上运行命令。


## dump
Creates a core dump from the current process state

​	从当前进程状态创建核心转储。

	dump <output file>

The core dump is always written in ELF, even on systems (windows, macOS) where this is not customary. For environments other than linux/amd64 threads and registers are dumped in a format that only Delve can read back.

​	核心转储总是以 ELF 格式写入，即使在一些非 Linux 系统（如 Windows 和 macOS）上，这种格式并不常见。对于非 linux/amd64 环境，线程和寄存器会以 Delve 专用格式转储。


## edit
Open where you are in `$DELVE_EDITOR` or `$EDITOR`

​	在 `$DELVE_EDITOR` 或 `$EDITOR` 中打开当前位置。

	edit [locspec]

If locspec is omitted edit will open the current source file in the editor, otherwise it will open the specified location.

​	如果省略 locspec，`edit` 将在编辑器中打开当前源文件，否则它会打开指定的位置。

Aliases: ed

​	别名：ed

## examinemem
Examine raw memory at the given address.

​	检查给定地址的原始内存。

Examine memory:

​	检查内存：

	examinemem [-fmt <format>] [-count|-len <count>] [-size <size>] <address>
	examinemem [-fmt <format>] [-count|-len <count>] [-size <size>] -x <expression>

Format represents the data format and the value is one of this list (default hex): bin(binary), oct(octal), dec(decimal), hex(hexadecimal) and raw. Length is the number of bytes (default 1) and must be less than or equal to 1000. Address is the memory location of the target to examine. Please note '-len' is deprecated by '-count and -size'. Expression can be an integer expression or pointer value of the memory location to examine.

​	`-fmt` 表示数据格式，值可以是以下之一（默认为十六进制）：bin（二进制）、oct（八进制）、dec（十进制）、hex（十六进制）和 raw。 `-len` 是字节数（默认为 1），并且必须小于或等于 1000。 `<address>` 是要检查的目标内存位置。请注意，`-len` 已被 `-count` 和 `-size` 取代。 `<expression>` 可以是整数表达式或内存位置的指针值。

For example:

​	例如：

    x -fmt hex -count 20 -size 1 0xc00008af38
    x -fmt hex -count 20 -size 1 -x 0xc00008af38 + 8
    x -fmt hex -count 20 -size 1 -x &myVar
    x -fmt hex -count 20 -size 1 -x myPtrVar

Aliases: x

​	别名：x

## exit
Exit the debugger.

​	退出调试器。		

	exit [-c]

When connected to a headless instance started with the --accept-multiclient, pass -c to resume the execution of the target process before disconnecting.

​	当连接到一个通过 `--accept-multiclient` 启动的无头实例时，使用 `-c` 可以在断开连接前恢复目标进程的执行。

Aliases: quit q

​	别名：quit q

## frame
Set the current frame, or execute command on a different frame.

​	设置当前帧，或在不同的帧上执行命令。

	frame <m>
	frame <m> <command>

The first form sets frame used by subsequent commands such as "print" or "set". The second form runs the command on the given frame.

​	第一种形式设置后续命令（如 "print" 或 "set"）使用的帧。 第二种形式会在指定的帧上执行命令。


## funcs
Print list of functions.

​	打印函数列表。

	funcs [<regex>]

If regex is specified only the functions matching it will be returned.

​	如果指定了 regex，仅返回匹配的函数。


## goroutine
Shows or changes current goroutine

​	显示或切换当前 goroutine。

	goroutine
	goroutine <id>
	goroutine <id> <command>

Called without arguments it will show information about the current goroutine.

​	不带参数调用时，将显示当前 goroutine 的信息。

Called with a single argument it will switch to the specified goroutine.

​	带一个参数调用时，将切换到指定的 goroutine。

Called with more arguments it will execute a command on the specified goroutine.

​	带更多参数调用时，将在指定的 goroutine 上执行命令。

Aliases: gr

​	别名：gr

## goroutines
List program goroutines.

​	列出程序的所有 goroutines。

	goroutines [-u|-r|-g|-s] [-t [depth]] [-l] [-with loc expr] [-without loc expr] [-group argument] [-chan expr] [-exec command]

Print out info for every goroutine. The flag controls what information is shown along with each goroutine:

​	打印每个 goroutine 的信息。标志控制显示每个 goroutine 时附加的详细信息：

	-u	displays location of topmost stackframe in user code (default)
			显示用户代码中的最上层栈帧位置（默认）
			
	-r	displays location of topmost stackframe (including frames inside private runtime functions)
			显示最上层栈帧的位置（包括私有运行时函数中的帧）
			
	-g	displays location of go instruction that created the goroutine
			显示创建该 goroutine 的 go 指令的位置
			
	-s	displays location of the start function
			显示启动函数的位置
			
	-t	displays goroutine's stacktrace (an optional depth value can be specified, default: 10)
			显示 goroutine 的堆栈跟踪（可以指定可选的深度值，默认：10）
			
	-l	displays goroutine's labels
			显示 goroutine 的标签

If no flag is specified the default is -u, i.e. the first frame within the first 30 frames that is not executing a runtime private function.

​	如果没有指定标志，默认是 `-u`，即显示在前 30 个栈帧内第一个不是运行时私有函数的栈帧。

### FILTERING

If -with or -without are specified only goroutines that match the given condition are returned.

​	如果指定了 `-with` 或 `-without`，仅返回符合条件的 goroutines。

To only display goroutines where the specified location contains (or does not contain, for -without and -wo) expr as a substring, use:

​	仅显示在指定位置包含（或不包含，使用 `-without` 或 `-wo`）指定表达式作为子字符串的 goroutines，使用：

	goroutines -with (userloc|curloc|goloc|startloc) expr
	goroutines -w (userloc|curloc|goloc|startloc) expr
	goroutines -without (userloc|curloc|goloc|startloc) expr
	goroutines -wo (userloc|curloc|goloc|startloc) expr
	
	Where:
	userloc: filter by the location of the topmost stackframe in user code
	curloc: filter by the location of the topmost stackframe (including frames inside private runtime functions)
	goloc: filter by the location of the go instruction that created the goroutine
	startloc: filter by the location of the start function
	
	其中：
	userloc: 按用户代码中最上层栈帧的位置过滤
	curloc: 按最上层栈帧的位置（包括私有运行时函数）过滤
	goloc: 按创建 goroutine 的 go 指令位置过滤
	startloc: 按起始函数的位置过滤

To only display goroutines that have (or do not have) the specified label key and value, use:

​	可以使用以下方式仅显示具有（或不具有）指定标签键和值的 goroutines：

	goroutines -with label key=value
	goroutines -without label key=value

To only display goroutines that have (or do not have) the specified label key, use:

​	可以使用以下方式仅显示具有（或不具有）指定标签键的 goroutines：

	goroutines -with label key
	goroutines -without label key

To only display goroutines that are running (or are not running) on a OS thread, use:

​	要仅显示正在运行（或不在运行）在操作系统线程上的 goroutines，使用：


	goroutines -with running
	goroutines -without running

To only display user (or runtime) goroutines, use:

​	要仅显示用户（或运行时）goroutines，使用：

	goroutines -with user
	goroutines -without user

### CHANNELS

​	
To only show goroutines waiting to send to or receive from a specific channel use:

​	要仅显示等待向特定频道发送或从中接收的 goroutines，使用：

	goroutines -chan expr

Note that 'expr' must not contain spaces.

​	请注意，'expr' 中不得包含空格。

### GROUPING

	goroutines -group (userloc|curloc|goloc|startloc|running|user)
	
	Where:
	userloc: groups goroutines by the location of the topmost stackframe in user code
	curloc: groups goroutines by the location of the topmost stackframe
	goloc: groups goroutines by the location of the go instruction that created the goroutine
	startloc: groups goroutines by the location of the start function
	running: groups goroutines by whether they are running or not
	user: groups goroutines by weather they are user or runtime goroutines
	
	其中：
	
	userloc：根据用户代码中最上层栈帧的位置分组 goroutine。
	curloc：根据最上层栈帧的位置分组 goroutine。
	goloc：根据创建该 goroutine 的 go 指令的位置分组 goroutine。
	startloc：根据起始函数的位置分组 goroutine。
	running：根据 goroutine 是否正在运行分组。
	user：根据 goroutine 是否是用户 goroutine 还是运行时 goroutine 分组。

Groups goroutines by the given location, running status or user classification, up to 5 goroutines per group will be displayed as well as the total number of goroutines in the group.

​	根据给定的位置、运行状态或用户分类分组 goroutine，每个组最多显示 5 个 goroutine，并显示组内的 goroutine 总数。

	goroutines -group label key

Groups goroutines by the value of the label with the specified key.

​	根据指定键值的标签分组 goroutine。

### EXEC

	goroutines -exec <command>

Runs the command on every goroutine.

​	在每个 goroutine 上运行命令。

Aliases: grs

​	别名：`grs`

## help
Prints the help message.

​	打印帮助信息。

	help [command]

Type "help" followed by the name of a command for more information about it.

​	输入 "help" 后跟命令名称以获取更多信息。

Aliases: h

​	别名：`h`

## libraries
List loaded dynamic libraries

​	列出已加载的动态库。


## list
Show source code.

​	显示源代码。

	[goroutine <n>] [frame <m>] list [<locspec>]

Show source around current point or provided locspec.

​	显示当前点或提供的 `locspec` 附近的源代码。

For example:

​	例如：

	frame 1 list 69
	list testvariables.go:10000
	list main.main:30
	list 40

Aliases: ls、 l

​	别名：`ls` 、`l`

## locals
Print local variables.

​	打印局部变量。

	[goroutine <n>] [frame <m>] locals [-v] [<regex>]

The name of variables that are shadowed in the current scope will be shown in parenthesis.

​	在当前作用域中被隐藏的变量名称将以括号显示。

If regex is specified only local variables with a name matching it will be returned. If -v is specified more information about each local variable will be shown.

​	如果指定了正则表达式，只会返回名称匹配的局部变量。如果使用 `-v`，则会显示每个局部变量的更多信息。


## next
Step over to next source line.

​	执行下一行源代码。

	next [count]

Optional [count] argument allows you to skip multiple lines.

​	可选的 `count` 参数允许跳过多行。

Aliases: n

​	别名：`n`

## next-instruction
Single step a single cpu instruction, skipping function calls.

​	单步执行一个 CPU 指令，跳过函数调用。

Aliases: ni、 nexti

​	别名：`ni`、 `nexti`

## on
Executes a command when a breakpoint is hit.

​	当断点被触发时执行命令。

	on <breakpoint name or id> <command>
	on <breakpoint name or id> -edit

Supported commands: print, stack, goroutine, trace and cond. 

​	支持的命令：`print`，`stack`，`goroutine`，`trace` 和 `cond`。

To convert a breakpoint into a tracepoint use:
	要将断点转换为追踪点，请使用：

	on <breakpoint name or id> trace

The command 'on &lt;bp> cond &lt;cond-arguments>' is equivalent to 'cond &lt;bp> &lt;cond-arguments>'.

​	命令 `on <bp> cond <cond-arguments>` 等价于 `cond <bp> <cond-arguments>`。

The command 'on x -edit' can be used to edit the list of commands executed when the breakpoint is hit.

​	命令 `on x -edit` 可用于编辑触发断点时执行的命令列表。


## packages
Print list of packages.

​	打印包列表。

	packages [<regex>]

If regex is specified only the packages matching it will be returned.

​	如果指定了正则表达式，则仅返回匹配的包。


## print
Evaluate an expression.

​	求值一个表达式。

	[goroutine <n>] [frame <m>] print [%format] <expression>

See [Documentation/cli/expr.md](//github.com/go-delve/delve/tree/master/Documentation/cli/expr.md) for a description of supported expressions.

​	查看 [Documentation/cli/expr.md](http://github.com/go-delve/delve/tree/master/Documentation/cli/expr.md) 以获取支持的表达式的描述。

The optional format argument is a format specifier, like the ones used by the fmt package. For example "print %x v" will print v as an hexadecimal number.

​	可选的 `format` 参数是格式说明符，类似于 fmt 包中使用的格式说明符。例如，`print %x v` 会以十六进制形式打印 `v`。

Aliases: p

​	别名：`p`

## rebuild
Rebuild the target executable and restarts it. It does not work if the executable was not built by delve.

​	重建目标可执行文件并重新启动它。如果可执行文件不是由 delve 构建的，无法使用此命令。


## regs
Print contents of CPU registers.

​	打印 CPU 寄存器的内容。

	regs [-a]

Argument -a shows more registers. Individual registers can also be displayed by 'print' and 'display'. See [Documentation/cli/expr.md](//github.com/go-delve/delve/tree/master/Documentation/cli/expr.md).

​	`-a` 参数显示更多寄存器。单个寄存器也可以通过 `print` 和 `display` 显示。请参阅 [Documentation/cli/expr.md](http://github.com/go-delve/delve/tree/master/Documentation/cli/expr.md)。


## restart
Restart process.

​	重新启动进程。

For recorded targets the command takes the following forms:

​	对于已录制的目标，命令有以下几种形式：

	restart					resets to the start of the recording
								重置到录制的开始
								
	restart [checkpoint]			resets the recording to the given checkpoint
										重置到指定的检查点
										
	restart -r [newargv...]	[redirects...]	re-records the target process
												重新录制目标进程

For live targets the command takes the following forms:

​	对于实时目标，命令有以下几种形式：

	restart [newargv...] [redirects...]	restarts the process
											重新启动进程

If newargv is omitted the process is restarted (or re-recorded) with the same argument vector. If `-noargs` is specified instead, the argument vector is cleared.

​	如果省略了 `newargv`，进程会以相同的参数向量重新启动（或重新录制）。如果指定了 `-noargs`，则清除参数向量。

A list of file redirections can be specified after the new argument list to override the redirections defined using the '`--redirect`' command line option. A syntax similar to Unix shells is used:

​	可以在新的参数列表后指定文件重定向，以覆盖通过 `--redirect` 命令行选项定义的重定向。使用类似 Unix shell 的语法：

	<input.txt	redirects the standard input of the target process from input.txt
					将目标进程的标准输入重定向到 input.txt
					
	>output.txt	redirects the standard output of the target process to output.txt
					将目标进程的标准输出重定向到 output.txt
					
	2>error.txt	redirects the standard error of the target process to error.txt
					将目标进程的标准错误重定向到 error.txt

Aliases: r

​	别名：`r`

## rev
Reverses the execution of the target program for the command specified.

​	逆向执行目标程序以执行指定的命令。

Currently, rev next, step, step-instruction and stepout commands are supported.

​	目前，支持 `rev next`，`rev step`，`rev step-instruction` 和 `rev stepout` 命令。


## rewind
Run backwards until breakpoint or start of recorded history.

​	倒退执行，直到断点或录制历史的开始。

Aliases: rw

​	别名：`rw`

## set
Changes the value of a variable.

​	更改变量的值。

	[goroutine <n>] [frame <m>] set <variable> = <value>

See [Documentation/cli/expr.md](//github.com/go-delve/delve/tree/master/Documentation/cli/expr.md) for a description of supported expressions. Only numerical variables and pointers can be changed.

​	查看 [Documentation/cli/expr.md](http://github.com/go-delve/delve/tree/master/Documentation/cli/expr.md) 以获取支持的表达式的描述。仅能更改数值变量和指针。


## source
Executes a file containing a list of delve commands

​	执行包含 delve 命令列表的文件。

	source <path>

If path ends with the .star extension it will be interpreted as a starlark script. See [Documentation/cli/starlark.md](//github.com/go-delve/delve/tree/master/Documentation/cli/starlark.md) for the syntax.

​	如果 `path` 以 `.star` 扩展名结尾，它将被解释为 Starlark 脚本。请参阅 [Documentation/cli/starlark.md](http://github.com/go-delve/delve/tree/master/Documentation/cli/starlark.md) 以了解语法。

If path is a single '-' character an interactive starlark interpreter will start instead. Type 'exit' to exit.

​	如果 `path` 是单个 `-` 字符，则将启动交互式 Starlark 解释器。输入 `exit` 以退出。


## sources
Print list of source files.

​	打印源文件列表。

	sources [<regex>]

If regex is specified only the source files matching it will be returned.

​	如果指定了正则表达式，则仅返回匹配的源文件。


## stack
Print stack trace.

​	打印栈跟踪。

	[goroutine <n>] [frame <m>] stack [<depth>] [-full] [-offsets] [-defer] [-a <n>] [-adepth <depth>] [-mode <mode>]
	
	-full		every stackframe is decorated with the value of its local variables and arguments.
					每个栈帧都装饰有其局部变量和参数的值。
					
	-offsets	prints frame offset of each frame.
					打印每个帧的帧偏移量。
					
	-defer		prints deferred function call stack for each frame.
					打印每个帧的延迟函数调用栈。
					
	-a <n>		prints stacktrace of n ancestors of the selected goroutine (target process must have tracebackancestors enabled)
					打印选定 goroutine 的 n 个祖先的栈跟踪（目标进程必须启用 tracebackancestors）。
					
	-adepth <depth>	configures depth of ancestor stacktrace
						配置祖先栈跟踪的深度。
						
	-mode <mode>	specifies the stacktrace mode, possible values are:
						指定栈跟踪模式，可能的值为：
			normal	- attempts to automatically switch between cgo frames and go frames
						尝试自动在 cgo 框架和 go 框架之间切换。
						
			simple	- disables automatic switch between cgo and go
						禁用在 cgo 和 go 之间的自动切换。
						
			fromg	- starts from the registers stored in the runtime.g struct	
						从存储在 runtime.g 结构中的寄存器开始。

Aliases: bt

​	别名：`bt`

## step
Single step through program.

​	单步执行程序。

Aliases: s

​	别名：`s`

## step-instruction
Single step a single cpu instruction.

​	单步执行一个 CPU 指令。

Aliases: si、 stepi

​	别名：`si`、 `stepi`

## stepout
Step out of the current function.

​	跳出当前函数。

Aliases: so

​	别名：`so`

## target
Manages child process debugging.

​	管理子进程调试。

	target follow-exec [-on [regex]] [-off]

Enables or disables follow exec mode. When follow exec mode Delve will automatically attach to new child processes executed by the target process. An optional regular expression can be passed to 'target follow-exec', only child processes with a command line matching the regular expression will be followed.

​	启用或禁用跟踪执行模式。当启用跟踪执行模式时，Delve 将自动附加到目标进程执行的任何新子进程。可以传递一个正则表达式给 `target follow-exec`，只有那些命令行与正则表达式匹配的子进程会被跟踪。

	target list

List currently attached processes.

​	列出当前附加的进程。

	target switch [pid]

Switches to the specified process.

​	切换到指定的进程。


## thread
Switch to the specified thread.

​	切换到指定的线程。

	thread <id>

Aliases: tr

​	别名：`tr`

## threads
Print out info for every traced thread.

​	打印每个跟踪线程的信息。


## toggle
Toggles on or off a breakpoint.

​	切换断点的开关状态。

	toggle <breakpoint name or id>


## trace
Set tracepoint.

​	设置追踪点。

	trace [name] [locspec]

A tracepoint is a breakpoint that does not stop the execution of the program, instead when the tracepoint is hit a notification is displayed. See [Documentation/cli/locspec.md](//github.com/go-delve/delve/tree/master/Documentation/cli/locspec.md) for the syntax of locspec. If locspec is omitted a tracepoint will be set on the current line.

​	追踪点是一种不停止程序执行的断点，而是在触发时显示通知。有关 `locspec` 的语法，请参见 [Documentation/cli/locspec.md](http://github.com/go-delve/delve/tree/master/Documentation/cli/locspec.md)。如果未指定 `locspec`，则将在当前行设置追踪点。

See also: "help on", "help cond" and "help clear"

​	另见："help on"、"help cond" 和 "help clear"。

Aliases: t

​	别名：`t`

## transcript
Appends command output to a file.

​	将命令输出追加到文件中。

	transcript [-t] [-x] <output file>
	transcript -off

Output of Delve's command is appended to the specified output file. If '-t' is specified and the output file exists it is truncated. If '-x' is specified output to stdout is suppressed instead.

​	Delve 命令的输出将追加到指定的输出文件中。如果指定了 `-t` 并且输出文件已存在，则会截断该文件。如果指定了 `-x`，则会抑制标准输出。

Using the -off option disables the transcript.

​	使用 `-off` 选项禁用转录。


## types
Print list of types

​	打印类型列表。

	types [<regex>]

If regex is specified only the types matching it will be returned.

​	如果指定了正则表达式，则仅返回匹配的类型。


## up
Move the current frame up.

​	将当前帧向上移动。

	up [<m>]
	up [<m>] <command>

Move the current frame up by &lt;m>. The second form runs the command on the given frame.

​	将当前帧向上移动 `<m>`。第二种形式在给定帧上运行命令。


## vars
Print package variables.

​	打印包变量。

	vars [-v] [<regex>]

If regex is specified only package variables with a name matching it will be returned. If -v is specified more information about each package variable will be shown.

​	如果指定了正则表达式，则仅返回名称匹配的包变量。如果使用 `-v`，则显示每个包变量的更多信息。


## watch
Set watchpoint.

​	设置监视点。	

	watch [-r|-w|-rw] <expr>
	
	-r	stops when the memory location is read
			当内存位置被读取时停止。
	
	-w	stops when the memory location is written
			当内存位置被写入时停止。
			
	-rw	stops when the memory location is read or written
			当内存位置被读取或写入时停止。

The memory location is specified with the same expression language used by 'print', for example:

​	内存位置使用与 `print` 命令相同的表达式语言指定。例如：

	watch v
	watch -w *(*int)(0x1400007c018)

will watch the address of variable 'v' and writes to an int at addr '0x1400007c018'.

​	这将监视变量 `v` 的地址，并写入地址为 `0x1400007c018` 的 `int` 类型。

Note that writes that do not change the value of the watched memory address might not be reported.

​	注意：如果写入操作没有改变监视内存地址的值，可能不会报告。

See also: "help print".

​	另见：`help print`。


## whatis
Prints type of an expression.

​	打印表达式的类型。

	whatis <expression>

