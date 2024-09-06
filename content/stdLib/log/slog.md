+++
title = "slog"
date = 2023-11-05T14:27:24+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文：[https://pkg.go.dev/slog@go1.23.0](https://pkg.go.dev/slog@go1.23.0)

> 注意
>
> ​	从go1.21.0开始才有该包。

## Overview 

Package slog provides structured logging, in which log records include a message, a severity level, and various other attributes expressed as key-value pairs.

​	`slog` 包提供了结构化日志记录功能，其中日志记录包含消息、严重级别以及其他以键值对形式表达的属性。

It defines a type, [Logger](https://pkg.go.dev/log/slog@go1.21.3#Logger), which provides several methods (such as [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info) and [Logger.Error](https://pkg.go.dev/log/slog@go1.21.3#Logger.Error)) for reporting events of interest.

​	它定义了一种类型，[Logger](https://pkg.go.dev/log/slog@go1.21.3#Logger)，该类型提供了多个方法（如 [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info) 和 [Logger.Error](https://pkg.go.dev/log/slog@go1.21.3#Logger.Error)）用于报告重要事件。

Each Logger is associated with a [Handler](https://pkg.go.dev/log/slog@go1.21.3#Handler). A Logger output method creates a [Record](https://pkg.go.dev/log/slog@go1.21.3#Record) from the method arguments and passes it to the Handler, which decides how to handle it. There is a default Logger accessible through top-level functions (such as [Info](https://pkg.go.dev/log/slog@go1.21.3#Info) and [Error](https://pkg.go.dev/log/slog@go1.21.3#Error)) that call the corresponding Logger methods.

​	每个 Logger 都与一个 [Handler](https://pkg.go.dev/log/slog@go1.21.3#Handler) 相关联。Logger 的输出方法会根据方法参数创建一个 [Record](https://pkg.go.dev/log/slog@go1.21.3#Record)，并将其传递给 Handler，由 Handler 决定如何处理。可以通过顶级函数（如 [Info](https://pkg.go.dev/log/slog@go1.21.3#Info) 和 [Error](https://pkg.go.dev/log/slog@go1.21.3#Error)）访问默认 Logger，这些函数会调用相应的 Logger 方法。

A log record consists of a time, a level, a message, and a set of key-value pairs, where the keys are strings and the values may be of any type. As an example,

​	日志记录由时间、级别、消息和一组键值对组成，其中键是字符串，值可以是任何类型。举个例子：

```
slog.Info("hello", "count", 3)
```

creates a record containing the time of the call, a level of Info, the message "hello", and a single pair with key "count" and value 3.

​	这将创建一条包含调用时间、Info 级别、消息 "hello" 以及键 "count" 和值 3 的记录。

The [Info](https://pkg.go.dev/log/slog@go1.21.3#Info) top-level function calls the [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info) method on the default Logger. In addition to [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info), there are methods for Debug, Warn and Error levels. Besides these convenience methods for common levels, there is also a [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) method which takes the level as an argument. Each of these methods has a corresponding top-level function that uses the default logger.

​	顶级函数 [Info](https://pkg.go.dev/log/slog@go1.21.3#Info) 调用了默认 Logger 的 [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info) 方法。除了 [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info) 之外，还有用于 Debug、Warn 和 Error 级别的方法。除了这些常见级别的便捷方法外，还有一个 [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) 方法，该方法将级别作为参数。每种方法都有一个对应的顶级函数，这些函数使用默认 Logger。

The default handler formats the log record's message, time, level, and attributes as a string and passes it to the [log](https://pkg.go.dev/log) package.

​	默认的 Handler 会将日志记录的消息、时间、级别和属性格式化为字符串，并将其传递给 [log](https://pkg.go.dev/log) 包。

```
2022/11/08 15:28:26 INFO hello count=3
```

For more control over the output format, create a logger with a different handler. This statement uses [New](https://pkg.go.dev/log/slog@go1.21.3#New) to create a new logger with a TextHandler that writes structured records in text form to standard error:

​	要对输出格式进行更细致的控制，可以创建一个具有不同处理器的 Logger。以下语句使用 [New](https://pkg.go.dev/log/slog@go1.21.3#New) 创建了一个使用 TextHandler 的新 Logger，该处理器将结构化记录以文本形式写入标准错误输出：

```
logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
```

[TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) output is a sequence of key=value pairs, easily and unambiguously parsed by machine. This statement:

​	[TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) 的输出是一系列键值对，机器可以轻松且明确地解析。这条语句：

```
logger.Info("hello", "count", 3)
```

produces this output:

​	会生成以下输出：

```
time=2022-11-08T15:28:26.000-05:00 level=INFO msg=hello count=3
```

The package also provides [JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler), whose output is line-delimited JSON:

​	该包还提供了 [JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler)，其输出为行分隔的 JSON：

```
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("hello", "count", 3)
```

produces this output:

​	生成以下输出：

```
{"time":"2022-11-08T15:28:26.000000000-05:00","level":"INFO","msg":"hello","count":3}
```

Both [TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) and [JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler) can be configured with [HandlerOptions](https://pkg.go.dev/log/slog@go1.21.3#HandlerOptions). There are options for setting the minimum level (see Levels, below), displaying the source file and line of the log call, and modifying attributes before they are logged.

​	[TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) 和 [JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler) 都可以通过 [HandlerOptions](https://pkg.go.dev/log/slog@go1.21.3#HandlerOptions) 进行配置。可以选择设置最小日志级别（参见下文的级别部分）、显示日志调用的源文件和行号，以及在日志记录之前修改属性。

Setting a logger as the default with

​	通过以下语句将 Logger 设置为默认 Logger：

```
slog.SetDefault(logger)
```

will cause the top-level functions like [Info](https://pkg.go.dev/log/slog@go1.21.3#Info) to use it. [SetDefault](https://pkg.go.dev/log/slog@go1.21.3#SetDefault) also updates the default logger used by the [log](https://pkg.go.dev/log) package, so that existing applications that use [log.Printf](https://pkg.go.dev/log#Printf) and related functions will send log records to the logger's handler without needing to be rewritten.

​	这样，像 [Info](https://pkg.go.dev/log/slog@go1.21.3#Info) 这样的顶级函数将会使用它。[SetDefault](https://pkg.go.dev/log/slog@go1.21.3#SetDefault) 还会更新 [log](https://pkg.go.dev/log) 包使用的默认 Logger，因此现有使用 [log.Printf](https://pkg.go.dev/log#Printf) 和相关函数的应用程序无需重写代码即可将日志记录发送到 Logger 的处理器。

Some attributes are common to many log calls. For example, you may wish to include the URL or trace identifier of a server request with all log events arising from the request. Rather than repeat the attribute with every log call, you can use [Logger.With](https://pkg.go.dev/log/slog@go1.21.3#Logger.With) to construct a new Logger containing the attributes:

​	有些属性在许多日志调用中都是通用的。例如，您可能希望将服务器请求的 URL 或跟踪标识符包含在与该请求相关的所有日志事件中。与其在每次日志调用中重复该属性，不如使用 [Logger.With](https://pkg.go.dev/log/slog@go1.21.3#Logger.With) 构造一个包含该属性的新 Logger：

```
logger2 := logger.With("url", r.URL)
```

The arguments to With are the same key-value pairs used in [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info). The result is a new Logger with the same handler as the original, but additional attributes that will appear in the output of every call.

​	With 的参数与 [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info) 使用的键值对相同。结果是一个新的 Logger，具有与原始 Logger 相同的处理器，但额外的属性将在每次调用的输出中出现。

### 日志级别 Levels 

A [Level](https://pkg.go.dev/log/slog@go1.21.3#Level) is an integer representing the importance or severity of a log event. The higher the level, the more severe the event. This package defines constants for the most common levels, but any int can be used as a level.

​	一个 [日志级别](https://pkg.go.dev/log/slog@go1.21.3#Level) 是表示日志事件重要性或严重性的整数。级别越高，事件越严重。此包定义了常见级别的常量，但可以使用任何整数作为级别。

In an application, you may wish to log messages only at a certain level or greater. One common configuration is to log messages at Info or higher levels, suppressing debug logging until it is needed. The built-in handlers can be configured with the minimum level to output by setting [HandlerOptions.Level]. The program's `main` function typically does this. The default value is LevelInfo.

​	在应用程序中，您可能希望仅记录某个级别或更高的消息。一种常见的配置是记录 Info 级别或更高级别的消息，直到需要时再启用调试日志。可以通过设置 [HandlerOptions.Level] 来配置内置处理程序的输出最小级别，通常在程序的 `main` 函数中完成。默认值为 LevelInfo。

Setting the [HandlerOptions.Level] field to a [Level](https://pkg.go.dev/log/slog@go1.21.3#Level) value fixes the handler's minimum level throughout its lifetime. Setting it to a [LevelVar](https://pkg.go.dev/log/slog@go1.21.3#LevelVar) allows the level to be varied dynamically. A LevelVar holds a Level and is safe to read or write from multiple goroutines. To vary the level dynamically for an entire program, first initialize a global LevelVar:

​	将 [HandlerOptions.Level] 字段设置为 [Level](https://pkg.go.dev/log/slog@go1.21.3#Level) 值后，处理程序的最小级别将在其整个生命周期内固定不变。将其设置为 [LevelVar](https://pkg.go.dev/log/slog@go1.21.3#LevelVar) 则允许动态更改级别。LevelVar 持有一个 Level，并且可以安全地从多个 goroutine 中读写。要为整个程序动态调整日志级别，首先初始化一个全局 LevelVar：

``` go
var programLevel = new(slog.LevelVar) // Info by default
```

Then use the LevelVar to construct a handler, and make it the default:

​	然后使用该 LevelVar 构建一个处理程序，并将其设为默认处理程序：

```
h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
slog.SetDefault(slog.New(h))
```

Now the program can change its logging level with a single statement:

​	现在，程序可以通过一条语句更改日志级别：

```
programLevel.Set(slog.LevelDebug)
```

### 组 Groups 

Attributes can be collected into groups. A group has a name that is used to qualify the names of its attributes. How this qualification is displayed depends on the handler. [TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) separates the group and attribute names with a dot. [JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler) treats each group as a separate JSON object, with the group name as the key.

​	属性可以被收集到组中。组有一个名称，用于限定其属性名称。显示此限定的方式取决于处理程序。[TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) 用点号分隔组名和属性名。[JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler) 将每个组视为一个单独的 JSON 对象，以组名作为键。

Use [Group](https://pkg.go.dev/log/slog@go1.21.3#Group) to create a Group attribute from a name and a list of key-value pairs:

​	使用 [Group](https://pkg.go.dev/log/slog@go1.21.3#Group) 可以通过名称和键值对列表创建一个组属性：

```
slog.Group("request",
    "method", r.Method,
    "url", r.URL)
```

TextHandler would display this group as

​	TextHandler 会将该组显示为：

```
request.method=GET request.url=http://example.com
```

JSONHandler would display it as

​	而 JSONHandler 会显示为：

```
"request":{"method":"GET","url":"http://example.com"}
```

Use [Logger.WithGroup](https://pkg.go.dev/log/slog@go1.21.3#Logger.WithGroup) to qualify all of a Logger's output with a group name. Calling WithGroup on a Logger results in a new Logger with the same Handler as the original, but with all its attributes qualified by the group name.

​	使用 [Logger.WithGroup](https://pkg.go.dev/log/slog@go1.21.3#Logger.WithGroup) 可以将所有日志输出与组名关联。调用 Logger 的 WithGroup 会生成一个新 Logger，具有与原始 Logger 相同的处理程序，但所有属性都会加上组名作为前缀。

This can help prevent duplicate attribute keys in large systems, where subsystems might use the same keys. Pass each subsystem a different Logger with its own group name so that potential duplicates are qualified:

​	这可以帮助避免在大型系统中出现重复的属性键，例如子系统可能会使用相同的键。为每个子系统提供不同的 Logger，并为它们指定各自的组名，从而避免可能的重复：

```
logger := slog.Default().With("id", systemID)
parserLogger := logger.WithGroup("parser")
parseInput(input, parserLogger)
```

When parseInput logs with parserLogger, its keys will be qualified with "parser", so even if it uses the common key "id", the log line will have distinct keys.

​	当 parseInput 使用 parserLogger 进行日志记录时，其键将带有 "parser" 前缀，因此即使使用了常见键 "id"，日志行中的键也会是唯一的。

### 上下文 Contexts 

Some handlers may wish to include information from the [context.Context](https://pkg.go.dev/context#Context) that is available at the call site. One example of such information is the identifier for the current span when tracing is enabled.

​	某些处理程序可能希望包括来自调用站点的 [context.Context](https://pkg.go.dev/context#Context) 中的信息。例如，启用追踪时，可能需要记录当前 span 的标识符。

The [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) and [Logger.LogAttrs](https://pkg.go.dev/log/slog@go1.21.3#Logger.LogAttrs) methods take a context as a first argument, as do their corresponding top-level functions.

​	[Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) 和 [Logger.LogAttrs](https://pkg.go.dev/log/slog@go1.21.3#Logger.LogAttrs) 方法的第一个参数为 context，它们对应的顶层函数也同样如此。

Although the convenience methods on Logger (Info and so on) and the corresponding top-level functions do not take a context, the alternatives ending in "Context" do. For example,

​	尽管 Logger 的便捷方法（如 Info 等）和相应的顶层函数不接受 context 参数，但以 "Context" 结尾的方法则可以。例如：

```
slog.InfoContext(ctx, "message")
```

It is recommended to pass a context to an output method if one is available.

​	如果上下文可用，建议将其传递给日志输出方法。

### 属性和值 Attrs and Values 

An [Attr](https://pkg.go.dev/log/slog@go1.21.3#Attr) is a key-value pair. The Logger output methods accept Attrs as well as alternating keys and values. The statement

​	一个 [属性](https://pkg.go.dev/log/slog@go1.21.3#Attr) 是一个键值对。Logger 的输出方法既接受属性，也接受交替的键和值。以下语句：

```
slog.Info("hello", slog.Int("count", 3))
```

behaves the same as

​	与此语句行为相同：

```
slog.Info("hello", "count", 3)
```

There are convenience constructors for [Attr](https://pkg.go.dev/log/slog@go1.21.3#Attr) such as [Int](https://pkg.go.dev/log/slog@go1.21.3#Int), [String](https://pkg.go.dev/log/slog@go1.21.3#String), and [Bool](https://pkg.go.dev/log/slog@go1.21.3#Bool) for common types, as well as the function [Any](https://pkg.go.dev/log/slog@go1.21.3#Any) for constructing Attrs of any type.

​	对于常见类型，有用于构造 [属性](https://pkg.go.dev/log/slog@go1.21.3#Attr) 的便捷构造器，例如 [Int](https://pkg.go.dev/log/slog@go1.21.3#Int)、[String](https://pkg.go.dev/log/slog@go1.21.3#String) 和 [Bool](https://pkg.go.dev/log/slog@go1.21.3#Bool)，也有用于构造任何类型属性的函数 [Any](https://pkg.go.dev/log/slog@go1.21.3#Any)。

The value part of an Attr is a type called [Value](https://pkg.go.dev/log/slog@go1.21.3#Value). Like an [any], a Value can hold any Go value, but it can represent typical values, including all numbers and strings, without an allocation.

​	属性的值部分是一个称为 [Value](https://pkg.go.dev/log/slog@go1.21.3#Value) 的类型。类似于 [any]，Value 可以包含任何 Go 值，但它能够无分配地表示典型值，包括所有数字和字符串。

For the most efficient log output, use [Logger.LogAttrs](https://pkg.go.dev/log/slog@go1.21.3#Logger.LogAttrs). It is similar to [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) but accepts only Attrs, not alternating keys and values; this allows it, too, to avoid allocation.

​	为了实现最有效的日志输出，使用 [Logger.LogAttrs](https://pkg.go.dev/log/slog@go1.21.3#Logger.LogAttrs)。它类似于 [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log)，但仅接受属性，不接受交替的键和值；这也使其能够避免分配。

The call

​	以下调用：

```
logger.LogAttrs(ctx, slog.LevelInfo, "hello", slog.Int("count", 3))
```

is the most efficient way to achieve the same output as

​	是实现与以下语句相同输出的最有效方式：

```
slog.Info("hello", "count", 3)
```

### 自定义类型的日志行为 Customizing a type's logging behavior 

If a type implements the [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer) interface, the [Value](https://pkg.go.dev/log/slog@go1.21.3#Value) returned from its LogValue method is used for logging. You can use this to control how values of the type appear in logs. For example, you can redact secret information like passwords, or gather a struct's fields in a Group. See the examples under [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer) for details.

​	如果某个类型实现了 [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer) 接口，那么其 `LogValue` 方法返回的 [Value](https://pkg.go.dev/log/slog@go1.21.3#Value) 将用于日志记录。你可以利用这一点来控制该类型的值在日志中的显示方式。例如，你可以对密码等敏感信息进行屏蔽，或者将结构体的字段收集到一个 Group 中。请参阅 [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer) 下的示例以获取更多细节。

A LogValue method may return a Value that itself implements [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer). The [Value.Resolve](https://pkg.go.dev/log/slog@go1.21.3#Value.Resolve) method handles these cases carefully, avoiding infinite loops and unbounded recursion. Handler authors and others may wish to use Value.Resolve instead of calling LogValue directly.

​	`LogValue` 方法可能会返回一个自身也实现了 [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer) 接口的 `Value`。在这种情况下，使用 [Value.Resolve](https://pkg.go.dev/log/slog@go1.21.3#Value.Resolve) 方法来处理这些情况，从而避免无限循环和无限递归。处理程序作者和其他人可以选择使用 `Value.Resolve`，而不是直接调用 `LogValue`。

### 包装输出方法 Wrapping output methods 

The logger functions use reflection over the call stack to find the file name and line number of the logging call within the application. This can produce incorrect source information for functions that wrap slog. For instance, if you define this function in file mylog.go:

​	日志记录函数通过对调用栈进行反射来查找应用程序中日志调用的文件名和行号。这可能会导致包装 slog 的函数产生不正确的源信息。例如，如果你在文件 `mylog.go` 中定义了以下函数：

``` go
func Infof(format string, args ...any) {
    slog.Default().Info(fmt.Sprintf(format, args...))
}
```

and you call it like this in main.go:

​	并在 `main.go` 中这样调用它：

```
Infof(slog.Default(), "hello, %s", "world")
```

then slog will report the source file as mylog.go, not main.go.

​	那么 slog 会将源文件报告为 `mylog.go`，而不是 `main.go`。

A correct implementation of Infof will obtain the source location (pc) and pass it to NewRecord. The Infof function in the package-level example called "wrapping" demonstrates how to do this.

​	正确的 `Infof` 实现应获取源位置（`pc`）并将其传递给 `NewRecord`。在名为 "wrapping" 的包级示例中，`Infof` 函数演示了如何实现这一点。

### 处理记录 Working with Records 

Sometimes a Handler will need to modify a Record before passing it on to another Handler or backend. A Record contains a mixture of simple public fields (e.g. Time, Level, Message) and hidden fields that refer to state (such as attributes) indirectly. This means that modifying a simple copy of a Record (e.g. by calling [Record.Add](https://pkg.go.dev/log/slog@go1.21.3#Record.Add) or [Record.AddAttrs](https://pkg.go.dev/log/slog@go1.21.3#Record.AddAttrs) to add attributes) may have unexpected effects on the original. Before modifying a Record, use [Record.Clone](https://pkg.go.dev/log/slog@go1.21.3#Record.Clone) to create a copy that shares no state with the original, or create a new Record with [NewRecord](https://pkg.go.dev/log/slog@go1.21.3#NewRecord) and build up its Attrs by traversing the old ones with [Record.Attrs](https://pkg.go.dev/log/slog@go1.21.3#Record.Attrs).

​	有时，处理程序需要在将记录传递给其他处理程序或后端之前对其进行修改。记录包含简单的公共字段（如 `Time`、`Level`、`Message`）和间接引用状态的隐藏字段（如属性）。这意味着简单复制记录（例如，通过调用 [Record.Add](https://pkg.go.dev/log/slog@go1.21.3#Record.Add) 或 [Record.AddAttrs](https://pkg.go.dev/log/slog@go1.21.3#Record.AddAttrs) 来添加属性）可能会对原始记录产生意想不到的影响。在修改记录之前，使用 [Record.Clone](https://pkg.go.dev/log/slog@go1.21.3#Record.Clone) 来创建一个不与原始记录共享状态的副本，或使用 [NewRecord](https://pkg.go.dev/log/slog@go1.21.3#NewRecord) 创建新记录，并通过遍历旧的记录来构建它的属性。

### 性能考虑 Performance considerations 

If profiling your application demonstrates that logging is taking significant time, the following suggestions may help.

​	如果性能分析表明日志记录占用了大量时间，可以参考以下建议。

If many log lines have a common attribute, use [Logger.With](https://pkg.go.dev/log/slog@go1.21.3#Logger.With) to create a Logger with that attribute. The built-in handlers will format that attribute only once, at the call to [Logger.With](https://pkg.go.dev/log/slog@go1.21.3#Logger.With). The [Handler](https://pkg.go.dev/log/slog@go1.21.3#Handler) interface is designed to allow that optimization, and a well-written Handler should take advantage of it.

​	如果许多日志行具有公共属性，请使用 [Logger.With](https://pkg.go.dev/log/slog@go1.21.3#Logger.With) 创建带有该属性的 `Logger`。内置处理程序只会在调用 [Logger.With](https://pkg.go.dev/log/slog@go1.21.3#Logger.With) 时格式化该属性。`Handler` 接口被设计为允许该优化，编写良好的处理程序应充分利用这一点。

The arguments to a log call are always evaluated, even if the log event is discarded. If possible, defer computation so that it happens only if the value is actually logged. For example, consider the call

​	日志调用的参数总是会被计算，即使日志事件被丢弃也是如此。如果可能，请推迟计算，直到该值实际需要记录。例如，考虑以下调用：

```
slog.Info("starting request", "url", r.URL.String())  // may compute String unnecessarily 可能不必要地计算了 String
```

The URL.String method will be called even if the logger discards Info-level events. Instead, pass the URL directly:

​	即使日志丢弃了 `Info` 级别的事件，`URL.String` 方法仍然会被调用。相反，直接传递 `URL`：

```
slog.Info("starting request", "url", &r.URL) // calls URL.String only if needed 只有在需要时才会调用 URL.String
```

The built-in [TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) will call its String method, but only if the log event is enabled. Avoiding the call to String also preserves the structure of the underlying value. For example [JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler) emits the components of the parsed URL as a JSON object. If you want to avoid eagerly paying the cost of the String call without causing the handler to potentially inspect the structure of the value, wrap the value in a fmt.Stringer implementation that hides its Marshal methods.

​	内置的 [TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) 只会在日志事件启用时调用其 `String` 方法。避免调用 `String` 还可以保留底层值的结构。例如，`JSONHandler` 会将解析后的 URL 组件作为 JSON 对象发出。如果你想避免急切地支付 `String` 调用的开销，同时又不希望处理程序检查值的结构，可以将该值包装在 `fmt.Stringer` 实现中，隐藏其 `Marshal` 方法。

You can also use the [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer) interface to avoid unnecessary work in disabled log calls. Say you need to log some expensive value:

​	你还可以使用 [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer) 接口来避免在禁用的日志调用中执行不必要的工作。例如，如果你需要记录一些耗时的值：

```
slog.Debug("frobbing", "value", computeExpensiveValue(arg))
```

Even if this line is disabled, computeExpensiveValue will be called. To avoid that, define a type implementing LogValuer:

​	即使该行被禁用，`computeExpensiveValue` 也会被调用。为了避免这种情况，定义一个实现 `LogValuer` 的类型：

``` go
type expensive struct { arg int }

func (e expensive) LogValue() slog.Value {
    return slog.AnyValue(computeExpensiveValue(e.arg))
}
```

Then use a value of that type in log calls:

​	然后在日志调用中使用该类型的值：

```
slog.Debug("frobbing", "value", expensive{arg})
```

Now computeExpensiveValue will only be called when the line is enabled.

​	这样，`computeExpensiveValue` 只会在该行启用时被调用。

The built-in handlers acquire a lock before calling [io.Writer.Write](https://pkg.go.dev/io#Writer.Write) to ensure that each record is written in one piece. User-defined handlers are responsible for their own locking.

​	内置处理程序在调用 [io.Writer.Write](https://pkg.go.dev/io#Writer.Write) 之前会获取锁，以确保每条记录都完整写入。用户定义的处理程序则需自行负责锁定。

### 编写处理程序 Writing a handler 

For a guide to writing a custom handler, see https://golang.org/s/slog-handler-guide.

​	关于如何编写自定义处理程序的指南，请参见 https://golang.org/s/slog-handler-guide。

## Example (Wrapping)

``` go
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Infof is an example of a user-defined logging function that wraps slog.
// The log record contains the source position of the caller of Infof.
// Infof 是一个用户自定义的日志记录函数示例，它包装了 slog。
// 日志记录包含 Infof 调用者的源位置。
func Infof(logger *slog.Logger, format string, args ...any) {
	if !logger.Enabled(context.Background(), slog.LevelInfo) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:]) // skip [Callers, Infof]
	r := slog.NewRecord(time.Now(), slog.LevelInfo, fmt.Sprintf(format, args...), pcs[0])
	_ = logger.Handler().Handle(context.Background(), r)
}

func main() {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		// Remove time.
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}
		// Remove the directory from the source's filename.
        // 从源文件名中移除目录。
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, ReplaceAttr: replace}))
	Infof(logger, "message, %s", "formatted")

}
Output:

level=INFO source=example_wrap_test.go:43 msg="message, formatted"
```


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.21.3:src/log/slog/handler.go;l=170)

``` go
const (
	// TimeKey is the key used by the built-in handlers for the time
	// when the log method is called. The associated Value is a [time.Time].
    // TimeKey 是内置处理程序用于记录日志调用时间的键。
	// 关联的值是 [time.Time]。
	TimeKey = "time"
	// LevelKey is the key used by the built-in handlers for the level
	// of the log call. The associated value is a [Level].
    // LevelKey 是内置处理程序用于记录日志调用级别的键。
	// 关联的值是 [Level]。
	LevelKey = "level"
	// MessageKey is the key used by the built-in handlers for the
	// message of the log call. The associated value is a string.
    // MessageKey 是内置处理程序用于记录日志消息的键。
	// 关联的值是字符串。
	MessageKey = "msg"
	// SourceKey is the key used by the built-in handlers for the source file
	// and line of the log call. The associated value is a string.
    // SourceKey 是内置处理程序用于记录日志调用的源文件和行号的键。
	// 关联的值是字符串。
	SourceKey = "source"
)
```

Keys for "built-in" attributes.

​	这是“内置”属性的键。

## 变量

This section is empty.

## Functions 

### func Debug 

``` go
func Debug(msg string, args ...any)
```

Debug calls Logger.Debug on the default logger.

​	`Debug` 调用默认日志记录器的 `Logger.Debug` 方法。

### func DebugContext 

``` go
func DebugContext(ctx context.Context, msg string, args ...any)
```

DebugContext calls Logger.DebugContext on the default logger.

​	`DebugContext` 调用默认日志记录器的 `Logger.DebugContext` 方法。

### func Error 

``` go
func Error(msg string, args ...any)
```

Error calls Logger.Error on the default logger.

​	`Error` 调用默认日志记录器的 `Logger.Error` 方法。

### func ErrorContext 

``` go
func ErrorContext(ctx context.Context, msg string, args ...any)
```

ErrorContext calls Logger.ErrorContext on the default logger.

​	`ErrorContext` 调用默认日志记录器的 `Logger.ErrorContext` 方法。

### func Info 

``` go
func Info(msg string, args ...any)
```

Info calls Logger.Info on the default logger.

​	`Info` 调用默认日志记录器的 `Logger.Info` 方法。

### func InfoContext 

``` go
func InfoContext(ctx context.Context, msg string, args ...any)
```

InfoContext calls Logger.InfoContext on the default logger.

​	`InfoContext` 调用默认日志记录器的 `Logger.InfoContext` 方法。

### func Log 

``` go
func Log(ctx context.Context, level Level, msg string, args ...any)
```

Log calls Logger.Log on the default logger.

​	`Log` 调用默认日志记录器的 `Logger.Log` 方法。

### func LogAttrs 

``` go
func LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)
```

LogAttrs calls Logger.LogAttrs on the default logger.

​	`LogAttrs` 调用默认日志记录器的 `Logger.LogAttrs` 方法。

### func NewLogLogger 

``` go
func NewLogLogger(h Handler, level Level) *log.Logger
```

NewLogLogger returns a new log.Logger such that each call to its Output method dispatches a Record to the specified handler. The logger acts as a bridge from the older log API to newer structured logging handlers.

​	`NewLogLogger` 返回一个新的 `log.Logger`，每次调用其 `Output` 方法时都会将记录分派给指定的处理程序。此日志记录器充当旧日志 API 与新结构化日志处理程序之间的桥梁。

### func SetDefault 

``` go
func SetDefault(l *Logger)
```

SetDefault makes l the default Logger. After this call, output from the log package's default Logger (as with [log.Print](https://pkg.go.dev/log#Print), etc.) will be logged at LevelInfo using l's Handler.

​	`SetDefault` 将 `l` 设置为默认日志记录器。此调用后，日志包的默认日志记录器输出（如 [log.Print](https://pkg.go.dev/log#Print) 等）将使用 `l` 的处理程序在 `LevelInfo` 级别进行记录。

### func Warn 

``` go
func Warn(msg string, args ...any)
```

Warn calls Logger.Warn on the default logger.

​	Warn 使用默认的日志记录器调用 Logger.Warn。

### func WarnContext 

``` go
func WarnContext(ctx context.Context, msg string, args ...any)
```

WarnContext calls Logger.WarnContext on the default logger.

​	WarnContext 使用默认的日志记录器调用 Logger.WarnContext，并支持上下文传递。

## Types 

### type Attr 

``` go
type Attr struct {
	Key   string
	Value Value
}
```

An Attr is a key-value pair.

​	Attr 是一个键值对。

#### func Any 

``` go
func Any(key string, value any) Attr
```

Any returns an Attr for the supplied value. See [AnyValue](https://pkg.go.dev/log/slog@go1.21.3#AnyValue) for how values are treated.

​	Any 返回一个给定值的 Attr。参考 [AnyValue](https://pkg.go.dev/log/slog@go1.21.3#AnyValue) 查看如何处理这些值。

#### func Bool 

``` go
func Bool(key string, v bool) Attr
```

Bool returns an Attr for a bool.

​	Bool 返回一个布尔值的 Attr。

#### func Duration 

``` go
func Duration(key string, v time.Duration) Attr
```

Duration returns an Attr for a time.Duration.

​	Duration 返回一个 time.Duration 类型值的 Attr。

#### func Float64 

``` go
func Float64(key string, v float64) Attr
```

Float64 returns an Attr for a floating-point number.

​	Float64 返回一个浮点数的 Attr。

#### func Group 

``` go
func Group(key string, args ...any) Attr
```

Group returns an Attr for a Group Value. The first argument is the key; the remaining arguments are converted to Attrs as in [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log).

​	Group 返回一个组值的 Attr。第一个参数是键，其余参数将按照 [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) 的方式转换为 Attr。

Use Group to collect several key-value pairs under a single key on a log line, or as the result of LogValue in order to log a single value as multiple Attrs.

​	使用 Group 来将多个键值对归类在一个日志行的单一键下，或者作为 LogValue 的结果，将单个值记录为多个 Attr。

##### Example Group

``` go
package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	r, _ := http.NewRequest("GET", "localhost", nil)
	// ...

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey && len(groups) == 0 {
					return slog.Attr{}
				}
				return a
			},
		}),
	)
	logger.Info("finished",
		slog.Group("req",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String())),
		slog.Int("status", http.StatusOK),
		slog.Duration("duration", time.Second))

}
Output:

level=INFO msg=finished req.method=GET req.url=localhost status=200 duration=1s
```
#### func Int 

``` go
func Int(key string, value int) Attr
```

Int converts an int to an int64 and returns an Attr with that value.

​	Int 将 int 转换为 int64 并返回一个带有该值的 Attr。

#### func Int64 

``` go
func Int64(key string, value int64) Attr
```

Int64 returns an Attr for an int64.

​	Int64 返回一个 int64 类型值的 Attr。

#### func String 

``` go
func String(key, value string) Attr
```

String returns an Attr for a string value.

​	String 返回一个字符串值的 Attr。

#### func Time 

``` go
func Time(key string, v time.Time) Attr
```

Time returns an Attr for a time.Time. It discards the monotonic portion.

​	Time 返回一个 time.Time 类型值的 Attr。它会丢弃时间的单调部分。

#### func Uint64 

``` go
func Uint64(key string, v uint64) Attr
```

Uint64 returns an Attr for a uint64.

​	Uint64 返回一个 uint64 类型值的 Attr。

#### (Attr) Equal 

``` go
func (a Attr) Equal(b Attr) bool
```

Equal reports whether a and b have equal keys and values.

​	Equal 判断 a 和 b 是否具有相同的键和值。

#### (Attr) String 

``` go
func (a Attr) String() string
```

### type Handler 

``` go
type Handler interface {
	// Enabled reports whether the handler handles records at the given level.
	// The handler ignores records whose level is lower.
	// It is called early, before any arguments are processed,
	// to save effort if the log event should be discarded.
	// If called from a Logger method, the first argument is the context
	// passed to that method, or context.Background() if nil was passed
	// or the method does not take a context.
	// The context is passed so Enabled can use its values
	// to make a decision.
    // Enabled 判断处理器是否处理给定级别的记录。
	// 处理器会忽略较低级别的记录。
	// 这个方法在早期调用，以便在日志事件应该被丢弃时节省处理资源。
	// 如果从 Logger 方法中调用，第一个参数是传递给该方法的上下文，
	// 或 context.Background()（如果传递了 nil 或方法不接受上下文）。
	// 上下文被传递进来是为了让 Enabled 使用其值来做出决策。
	Enabled(context.Context, Level) bool

	// Handle handles the Record.
	// It will only be called when Enabled returns true.
	// The Context argument is as for Enabled.
	// It is present solely to provide Handlers access to the context's values.
	// Canceling the context should not affect record processing.
	// (Among other things, log messages may be necessary to debug a
	// cancellation-related problem.)
	//
	// Handle methods that produce output should observe the following rules:
	//   - If r.Time is the zero time, ignore the time.
	//   - If r.PC is zero, ignore it.
	//   - Attr's values should be resolved.
	//   - If an Attr's key and value are both the zero value, ignore the Attr.
	//     This can be tested with attr.Equal(Attr{}).
	//   - If a group's key is empty, inline the group's Attrs.
	//   - If a group has no Attrs (even if it has a non-empty key),
	//     ignore it.
    // Handle 处理 Record。
	// 它只会在 Enabled 返回 true 时被调用。
	// Context 参数与 Enabled 的参数一样。
	// 它仅用于为处理器提供访问上下文值的机会。
	// 取消上下文不应影响记录处理。
	// （例如，日志消息可能是调试与取消相关问题所必需的。）
	//
	// Handle 方法产生输出时应该遵守以下规则：
	//   - 如果 r.Time 是零时间，忽略时间。
	//   - 如果 r.PC 为零，忽略它。
	//   - Attr 的值应已解析。
	//   - 如果 Attr 的键和值都是零值，忽略该 Attr。
	//     这可以通过 attr.Equal(Attr{}) 来测试。
	//   - 如果组的键为空，内联组的 Attrs。
	//   - 如果组没有 Attrs（即使它有一个非空键），
	//     也忽略它。
	Handle(context.Context, Record) error

	// WithAttrs returns a new Handler whose attributes consist of
	// both the receiver's attributes and the arguments.
	// The Handler owns the slice: it may retain, modify or discard it.
    // WithAttrs 返回一个新的 Handler，其属性包含接收者的属性和参数。
	// Handler 拥有这个切片：它可以保留、修改或丢弃它。
	WithAttrs(attrs []Attr) Handler

	// WithGroup returns a new Handler with the given group appended to
	// the receiver's existing groups.
	// The keys of all subsequent attributes, whether added by With or in a
	// Record, should be qualified by the sequence of group names.
	//
	// How this qualification happens is up to the Handler, so long as
	// this Handler's attribute keys differ from those of another Handler
	// with a different sequence of group names.
	//
	// A Handler should treat WithGroup as starting a Group of Attrs that ends
	// at the end of the log event. That is,
	//
	//     logger.WithGroup("s").LogAttrs(level, msg, slog.Int("a", 1), slog.Int("b", 2))
	//
	// should behave like
	//
	//     logger.LogAttrs(level, msg, slog.Group("s", slog.Int("a", 1), slog.Int("b", 2)))
	//
	// If the name is empty, WithGroup returns the receiver.
    // WithGroup 返回一个新的 Handler，其中给定组被附加到接收者的现有组中。
	// 所有后续属性的键，无论是通过 With 添加的还是记录中的，
	// 都应该通过组名序列进行限定。
	//
	// 这种限定方式由 Handler 决定，只要这个 Handler 的属性键与另一个 Handler
	// 使用不同组名序列的属性键不同即可。
	//
	// Handler 应该将 WithGroup 视为在日志事件结束时开始的一组 Attrs。
	// 也就是说，
	//
	//     logger.WithGroup("s").LogAttrs(level, msg, slog.Int("a", 1), slog.Int("b", 2))
	//
	// 应该表现得像
	//
	//     logger.LogAttrs(level, msg, slog.Group("s", slog.Int("a", 1), slog.Int("b", 2)))
	//
	// 如果名称为空，WithGroup 返回接收者。
	WithGroup(name string) Handler
}
```

A Handler handles log records produced by a Logger..

​	Handler 处理由 Logger 生成的日志记录。

A typical handler may print log records to standard error, or write them to a file or database, or perhaps augment them with additional attributes and pass them on to another handler.

​	典型的处理器可能会将日志记录打印到标准错误输出、写入文件或数据库，或者为其添加额外属性并传递给另一个处理器。

Any of the Handler's methods may be called concurrently with itself or with other methods. It is the responsibility of the Handler to manage this concurrency.

​	Handler 的任何方法都可以与自身或其他方法并发调用。管理这种并发性是 Handler 的责任。

Users of the slog package should not invoke Handler methods directly. They should use the methods of [Logger](https://pkg.go.dev/log/slog@go1.21.3#Logger) instead.

​	`slog` 包的使用者不应该直接调用 Handler 方法，而应使用 [Logger](https://pkg.go.dev/log/slog@go1.21.3#Logger) 的方法。

#### Example (LevelHandler)

This example shows how to Use a LevelHandler to change the level of an existing Handler while preserving its other behavior.

​	这个示例展示了如何使用 LevelHandler 来更改现有 Handler 的级别，同时保留其其他行为。

This example demonstrates increasing the log level to reduce a logger's output.

​	此示例演示了如何提高日志级别以减少日志输出。

Another typical use would be to decrease the log level (to LevelDebug, say) during a part of the program that was suspected of containing a bug.

​	另一个典型的用法是在程序的某个部分（例如，怀疑存在 bug 的部分）降低日志级别（比如降低到 LevelDebug）。

``` go
package main

import (
	"context"
	"log/slog"
	"log/slog/internal/slogtest"
	"os"
)

// A LevelHandler wraps a Handler with an Enabled method
// that returns false for levels below a minimum.
type LevelHandler struct {
	level   slog.Leveler
	handler slog.Handler
}

// NewLevelHandler returns a LevelHandler with the given level.
// All methods except Enabled delegate to h.
func NewLevelHandler(level slog.Leveler, h slog.Handler) *LevelHandler {
	// Optimization: avoid chains of LevelHandlers.
	if lh, ok := h.(*LevelHandler); ok {
		h = lh.Handler()
	}
	return &LevelHandler{level, h}
}

// Enabled implements Handler.Enabled by reporting whether
// level is at least as large as h's level.
func (h *LevelHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level.Level()
}

// Handle implements Handler.Handle.
func (h *LevelHandler) Handle(ctx context.Context, r slog.Record) error {
	return h.handler.Handle(ctx, r)
}

// WithAttrs implements Handler.WithAttrs.
func (h *LevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return NewLevelHandler(h.level, h.handler.WithAttrs(attrs))
}

// WithGroup implements Handler.WithGroup.
func (h *LevelHandler) WithGroup(name string) slog.Handler {
	return NewLevelHandler(h.level, h.handler.WithGroup(name))
}

// Handler returns the Handler wrapped by h.
func (h *LevelHandler) Handler() slog.Handler {
	return h.handler
}

// This example shows how to Use a LevelHandler to change the level of an
// existing Handler while preserving its other behavior.
//
// This example demonstrates increasing the log level to reduce a logger's
// output.
//
// Another typical use would be to decrease the log level (to LevelDebug, say)
// during a part of the program that was suspected of containing a bug.
func main() {
	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime})
	logger := slog.New(NewLevelHandler(slog.LevelWarn, th))
	logger.Info("not printed")
	logger.Warn("printed")

}
Output:

level=WARN msg=printed
```
### type HandlerOptions 

``` go
type HandlerOptions struct {
	// AddSource causes the handler to compute the source code position
	// of the log statement and add a SourceKey attribute to the output.
    // AddSource 导致处理器计算日志语句的源代码位置
	// 并将 SourceKey 属性添加到输出中。
	AddSource bool

	// Level reports the minimum record level that will be logged.
	// The handler discards records with lower levels.
	// If Level is nil, the handler assumes LevelInfo.
	// The handler calls Level.Level for each record processed;
	// to adjust the minimum level dynamically, use a LevelVar.
    // Level 报告将被记录的最低记录级别。
	// 处理器丢弃较低级别的记录。
	// 如果 Level 为 nil，处理器假定为 LevelInfo。
	// 处理器会在处理每个记录时调用 Level.Level；
	// 要动态调整最小级别，请使用 LevelVar。
	Level Leveler

	// ReplaceAttr is called to rewrite each non-group attribute before it is logged.
	// The attribute's value has been resolved (see [Value.Resolve]).
	// If ReplaceAttr returns a zero Attr, the attribute is discarded.
	//
	// The built-in attributes with keys "time", "level", "source", and "msg"
	// are passed to this function, except that time is omitted
	// if zero, and source is omitted if AddSource is false.
	//
	// The first argument is a list of currently open groups that contain the
	// Attr. It must not be retained or modified. ReplaceAttr is never called
	// for Group attributes, only their contents. For example, the attribute
	// list
	//
	//     Int("a", 1), Group("g", Int("b", 2)), Int("c", 3)
	//
	// results in consecutive calls to ReplaceAttr with the following arguments:
	//
	//     nil, Int("a", 1)
	//     []string{"g"}, Int("b", 2)
	//     nil, Int("c", 3)
	//
	// ReplaceAttr can be used to change the default keys of the built-in
	// attributes, convert types (for example, to replace a `time.Time` with the
	// integer seconds since the Unix epoch), sanitize personal information, or
	// remove attributes from the output.
    // ReplaceAttr 被调用以在日志记录之前重写每个非组属性。
	// 属性的值已解析（请参见 [Value.Resolve]）。
	// 如果 ReplaceAttr 返回零值 Attr，则该属性将被丢弃。
	//
	// 具有键 "time"、"level"、"source" 和 "msg" 的内置属性
	// 将传递给此函数，除非 time 为零或 AddSource 为 false 时 source 被忽略。
	//
	// 第一个参数是当前包含 Attr 的已打开组的列表。
	// 它不能被保留或修改。ReplaceAttr 永远不会对 Group 属性调用，
	// 只会调用其内容。例如，属性列表
	//
	//     Int("a", 1), Group("g", Int("b", 2)), Int("c", 3)
	//
	// 会连续调用 ReplaceAttr，参数如下：
	//
	//     nil, Int("a", 1)
	//     []string{"g"}, Int("b", 2)
	//     nil, Int("c", 3)
	//
	// ReplaceAttr 可用于更改内置属性的默认键、转换类型（例如，将 `time.Time` 替换为自 Unix 纪元以来的整数秒数）、清理个人信息，或从输出中删除属性。
	ReplaceAttr func(groups []string, a Attr) Attr
}
```

HandlerOptions are options for a TextHandler or JSONHandler. A zero HandlerOptions consists entirely of default values.

​	HandlerOptions 是 TextHandler 或 JSONHandler 的选项。零值的 HandlerOptions 全部由默认值组成。

#### Example (CustomLevels)

​	This example demonstrates using custom log levels and custom log level names. In addition to the default log levels, it introduces Trace, Notice, and Emergency levels. The ReplaceAttr changes the way levels are printed for both the standard log levels and the custom log levels.

​	这个示例演示了如何使用自定义日志级别和自定义日志级别名称。除了默认的日志级别之外，它还引入了 Trace、Notice 和 Emergency 级别。ReplaceAttr 更改了标准日志级别和自定义日志级别的打印方式。

``` go
package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	// Exported constants from a custom logging package.
	const (
		LevelTrace     = slog.Level(-8)
		LevelDebug     = slog.LevelDebug
		LevelInfo      = slog.LevelInfo
		LevelNotice    = slog.Level(2)
		LevelWarning   = slog.LevelWarn
		LevelError     = slog.LevelError
		LevelEmergency = slog.Level(12)
	)

	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// Set a custom level to show all log output. The default value is
		// LevelInfo, which would drop Debug and Trace logs.
		Level: LevelTrace,

		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Remove time from the output for predictable test output.
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}

			// Customize the name of the level key and the output string, including
			// custom level values.
			if a.Key == slog.LevelKey {
				// Rename the level key from "level" to "sev".
				a.Key = "sev"

				// Handle custom level values.
				level := a.Value.Any().(slog.Level)

				// This could also look up the name from a map or other structure, but
				// this demonstrates using a switch statement to rename levels. For
				// maximum performance, the string values should be constants, but this
				// example uses the raw strings for readability.
				switch {
				case level < LevelDebug:
					a.Value = slog.StringValue("TRACE")
				case level < LevelInfo:
					a.Value = slog.StringValue("DEBUG")
				case level < LevelNotice:
					a.Value = slog.StringValue("INFO")
				case level < LevelWarning:
					a.Value = slog.StringValue("NOTICE")
				case level < LevelError:
					a.Value = slog.StringValue("WARNING")
				case level < LevelEmergency:
					a.Value = slog.StringValue("ERROR")
				default:
					a.Value = slog.StringValue("EMERGENCY")
				}
			}

			return a
		},
	})

	logger := slog.New(th)
	ctx := context.Background()
	logger.Log(ctx, LevelEmergency, "missing pilots")
	logger.Error("failed to start engines", "err", "missing fuel")
	logger.Warn("falling back to default value")
	logger.Log(ctx, LevelNotice, "all systems are running")
	logger.Info("initiating launch")
	logger.Debug("starting background job")
	logger.Log(ctx, LevelTrace, "button clicked")

}
Output:

sev=EMERGENCY msg="missing pilots"
sev=ERROR msg="failed to start engines" err="missing fuel"
sev=WARNING msg="falling back to default value"
sev=NOTICE msg="all systems are running"
sev=INFO msg="initiating launch"
sev=DEBUG msg="starting background job"
sev=TRACE msg="button clicked"
```
### type JSONHandler 

``` go
type JSONHandler struct {
	// contains filtered or unexported fields
}
```

JSONHandler is a Handler that writes Records to an io.Writer as line-delimited JSON objects.

​	JSONHandler 是一种 Handler，用于将记录写入 io.Writer，格式为行分隔的 JSON 对象。

#### func NewJSONHandler 

``` go
func NewJSONHandler(w io.Writer, opts *HandlerOptions) *JSONHandler
```

NewJSONHandler creates a JSONHandler that writes to w, using the given options. If opts is nil, the default options are used.

​	NewJSONHandler 创建一个 JSONHandler，将输出写入 w，使用给定的选项。如果 opts 为 nil，则使用默认选项。

#### (*JSONHandler) Enabled 

``` go
func (h *JSONHandler) Enabled(_ context.Context, level Level) bool
```

Enabled reports whether the handler handles records at the given level. The handler ignores records whose level is lower.

​	Enabled 用于报告处理器是否处理给定级别的记录。处理器会忽略低于该级别的记录。

#### (*JSONHandler) Handle 

``` go
func (h *JSONHandler) Handle(_ context.Context, r Record) error
```

Handle formats its argument Record as a JSON object on a single line.

​	Handle 将其参数 Record 格式化为单行的 JSON 对象。

If the Record's time is zero, the time is omitted. Otherwise, the key is "time" and the value is output as with json.Marshal.

​	如果记录的时间为零，则省略时间。否则，键为 "time"，值的输出格式与 json.Marshal 相同。

If the Record's level is zero, the level is omitted. Otherwise, the key is "level" and the value of [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String) is output.

​	如果记录的级别为零，则省略级别。否则，键为 "level"，值为 [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String) 的输出结果。

If the AddSource option is set and source information is available, the key is "source", and the value is a record of type [Source](https://pkg.go.dev/log/slog@go1.21.3#Source).

​	如果设置了 AddSource 选项且源信息可用，则键为 "source"，值为类型为 [Source](https://pkg.go.dev/log/slog@go1.21.3#Source) 的记录。

The message's key is "msg".

​	消息的键为 "msg"。

To modify these or other attributes, or remove them from the output, use [HandlerOptions.ReplaceAttr].

​	要修改这些或其他属性，或从输出中删除它们，请使用 [HandlerOptions.ReplaceAttr]。

Values are formatted as with an [encoding/json.Encoder](https://pkg.go.dev/encoding/json#Encoder) with SetEscapeHTML(false), with two exceptions.

​	值的格式与 [encoding/json.Encoder](https://pkg.go.dev/encoding/json#Encoder) 使用 SetEscapeHTML(false) 时的格式相同，但有两个例外。

First, an Attr whose Value is of type error is formatted as a string, by calling its Error method. Only errors in Attrs receive this special treatment, not errors embedded in structs, slices, maps or other data structures that are processed by the encoding/json package.

​	首先，类型为 error 的 Attr 的值将被格式化为字符串，通过调用其 Error 方法。只有 Attr 中的错误会得到这种特殊处理，而不是嵌入到结构体、切片、映射或其他数据结构中的错误，这些数据结构是由 encoding/json 包处理的。

Second, an encoding failure does not cause Handle to return an error. Instead, the error message is formatted as a string.

​	其次，编码失败不会导致 Handle 返回错误。相反，错误消息将被格式化为字符串。

Each call to Handle results in a single serialized call to io.Writer.Write.

​	每次调用 Handle 都会导致单个序列化调用 io.Writer.Write。

#### (*JSONHandler) WithAttrs 

``` go
func (h *JSONHandler) WithAttrs(attrs []Attr) Handler
```

WithAttrs returns a new JSONHandler whose attributes consists of h's attributes followed by attrs.

​	WithAttrs 返回一个新的 JSONHandler，其属性包括 h 的属性和附加的 attrs。

#### (*JSONHandler) WithGroup 

``` go
func (h *JSONHandler) WithGroup(name string) Handler
```

### type Kind 

``` go
type Kind int
```

Kind is the kind of a Value.

​	Kind 表示 Value 的类型。

``` go
const (
	KindAny Kind = iota
	KindBool
	KindDuration
	KindFloat64
	KindInt64
	KindString
	KindTime
	KindUint64
	KindGroup
	KindLogValuer
)
```

#### (Kind) String 

``` go
func (k Kind) String() string
```

### type Level 

``` go
type Level int
```

A Level is the importance or severity of a log event. The higher the level, the more important or severe the event.

​	Level 表示日志事件的重要性或严重性。级别越高，事件越重要或严重。

``` go
const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)
```

Level numbers are inherently arbitrary, but we picked them to satisfy three constraints. Any system can map them to another numbering scheme if it wishes.

​	Level 数字本质上是任意的，但我们选择了这些数字以满足三个条件。任何系统都可以根据需要将它们映射到另一个编号方案。

First, we wanted the default level to be Info, Since Levels are ints, Info is the default value for int, zero.

​	首先，我们希望默认级别为 Info。由于 Level 是整数，因此 Info 是整数的默认值，即 0。

Second, we wanted to make it easy to use levels to specify logger verbosity. Since a larger level means a more severe event, a logger that accepts events with smaller (or more negative) level means a more verbose logger. Logger verbosity is thus the negation of event severity, and the default verbosity of 0 accepts all events at least as severe as INFO.

​	其次，我们希望能够轻松使用级别来指定记录器的详细程度。由于较大的级别表示事件更严重，接受较小（或更负）级别事件的记录器意味着它是一个更详细的记录器。因此，记录器的详细程度是事件严重性的反转，默认详细程度 0 接受所有至少与 INFO 同样严重的事件。

Third, we wanted some room between levels to accommodate schemes with named levels between ours. For example, Google Cloud Logging defines a Notice level between Info and Warn. Since there are only a few of these intermediate levels, the gap between the numbers need not be large. Our gap of 4 matches OpenTelemetry's mapping. Subtracting 9 from an OpenTelemetry level in the DEBUG, INFO, WARN and ERROR ranges converts it to the corresponding slog Level range. OpenTelemetry also has the names TRACE and FATAL, which slog does not. But those OpenTelemetry levels can still be represented as slog Levels by using the appropriate integers.

​	第三，我们希望在级别之间有一些空间，以适应我们命名级别之间的方案。例如，Google Cloud Logging 在 Info 和 Warn 之间定义了一个 Notice 级别。由于这些中间级别很少，数字之间的间隔不需要太大。我们选择的间隔 4 与 OpenTelemetry 的映射一致。将 OpenTelemetry 级别中的 DEBUG、INFO、WARN 和 ERROR 范围减去 9 后可以转换为相应的 slog 级别范围。OpenTelemetry 还有 TRACE 和 FATAL 名称，而 slog 没有这些名称。但这些 OpenTelemetry 级别仍然可以通过使用适当的整数表示为 slog 级别。

Names for common levels.

​	常见级别的名称。

#### func SetLogLoggerLevel <- go1.22.0

``` go
func SetLogLoggerLevel(level Level) (oldLevel Level)
```

SetLogLoggerLevel controls the level for the bridge to the [log](https://pkg.go.dev/log) package.

​	SetLogLoggerLevel 控制连接到 [log](https://pkg.go.dev/log) 包的桥接器的级别。

Before [SetDefault](https://pkg.go.dev/log/slog#SetDefault) is called, slog top-level logging functions call the default [log.Logger](https://pkg.go.dev/log#Logger). In that mode, SetLogLoggerLevel sets the minimum level for those calls. By default, the minimum level is Info, so calls to [Debug](https://pkg.go.dev/log/slog#Debug) (as well as top-level logging calls at lower levels) will not be passed to the log.Logger. After calling

​	在调用 [SetDefault](https://pkg.go.dev/log/slog#SetDefault) 之前，slog 顶级日志记录函数会调用默认的 [log.Logger](https://pkg.go.dev/log#Logger)。在该模式下，SetLogLoggerLevel 设置这些调用的最小级别。默认情况下，最小级别为 Info，因此 [Debug](https://pkg.go.dev/log/slog#Debug) 级别（以及较低级别的顶级日志记录调用）不会传递给 log.Logger。调用

```
slog.SetLogLoggerLevel(slog.LevelDebug)
```

calls to [Debug](https://pkg.go.dev/log/slog#Debug) will be passed to the log.Logger.

​	之后，[Debug](https://pkg.go.dev/log/slog#Debug) 的调用将传递给 log.Logger。

After [SetDefault](https://pkg.go.dev/log/slog#SetDefault) is called, calls to the default [log.Logger](https://pkg.go.dev/log#Logger) are passed to the slog default handler. In that mode, SetLogLoggerLevel sets the level at which those calls are logged. That is, after calling

​	在调用 [SetDefault](https://pkg.go.dev/log/slog#SetDefault) 之后，对默认 [log.Logger](https://pkg.go.dev/log#Logger) 的调用将传递给 slog 默认处理程序。在该模式下，SetLogLoggerLevel 设置记录这些调用的级别。也就是说，调用

```
slog.SetLogLoggerLevel(slog.LevelDebug)
```

A call to [log.Printf](https://pkg.go.dev/log#Printf) will result in output at level [LevelDebug](https://pkg.go.dev/log/slog#LevelDebug).

​	之后，调用 [log.Printf](https://pkg.go.dev/log#Printf) 将导致以 [LevelDebug](https://pkg.go.dev/log/slog#LevelDebug) 级别输出。

SetLogLoggerLevel returns the previous value.

​	SetLogLoggerLevel 返回先前的值。

##### Example  (Log)

This example shows how to use slog.SetLogLoggerLevel to change the minimal level of the internal default handler for slog package before calling slog.SetDefault.

​	此示例显示如何使用 slog.SetLogLoggerLevel 在调用 slog.SetDefault 之前更改 slog 包内部默认处理程序的最小级别。

```go
package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	defer log.SetFlags(log.Flags()) // revert changes after the example
	log.SetFlags(0)
	defer log.SetOutput(log.Writer()) // revert changes after the example
	log.SetOutput(os.Stdout)

	// Default logging level is slog.LevelInfo.
	log.Print("log debug") // log debug
	slog.Debug("debug")    // no output
	slog.Info("info")      // INFO info

	// Set the default logging level to slog.LevelDebug.
	currentLogLevel := slog.SetLogLoggerLevel(slog.LevelDebug)
	defer slog.SetLogLoggerLevel(currentLogLevel) // revert changes after the example

	log.Print("log debug") // log debug
	slog.Debug("debug")    // DEBUG debug
	slog.Info("info")      // INFO info

}
Output:

log debug
INFO info
log debug
DEBUG debug
INFO info
```

##### Example (Slog)

This example shows how to use slog.SetLogLoggerLevel to change the minimal level of the internal writer that uses the custom handler for log package after calling slog.SetDefault.

​	此示例显示如何使用 slog.SetLogLoggerLevel 在调用 slog.SetDefault 之后更改使用自定义处理程序的日志包内部记录器的最小级别。

```go
package main

import (
	"log"
	"log/slog"
	"log/slog/internal/slogtest"
	"os"
)

func main() {
	// Set the default logging level to slog.LevelError.
	currentLogLevel := slog.SetLogLoggerLevel(slog.LevelError)
	defer slog.SetLogLoggerLevel(currentLogLevel) // revert changes after the example

	defer slog.SetDefault(slog.Default()) // revert changes after the example
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime})))

	log.Print("error") // level=ERROR msg=error

}
Output:

level=ERROR msg=error
```



#### (Level) Level 

``` go
func (l Level) Level() Level
```

Level returns the receiver. It implements Leveler.

​	Level 返回接收者。它实现了 Leveler 接口。

#### (Level) MarshalJSON 

``` go
func (l Level) MarshalJSON() ([]byte, error)
```

MarshalJSON implements [encoding/json.Marshaler](https://pkg.go.dev/encoding/json#Marshaler) by quoting the output of [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String).

​	MarshalJSON 实现了 [encoding/json.Marshaler](https://pkg.go.dev/encoding/json#Marshaler)，它会引用 [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String) 的输出。

#### (Level) MarshalText 

``` go
func (l Level) MarshalText() ([]byte, error)
```

MarshalText implements [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler) by calling [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String).

​	MarshalText 通过调用 [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String) 实现了 [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler)。

#### (Level) String 

``` go
func (l Level) String() string
```

String returns a name for the level. If the level has a name, then that name in uppercase is returned. If the level is between named values, then an integer is appended to the uppercased name. Examples:

​	String 返回该级别的名称。如果该级别有名称，则返回该名称的大写形式。如果该级别位于命名值之间，则会在大写名称后附加一个整数。示例：

```
LevelWarn.String() => "WARN"
(LevelInfo+2).String() => "INFO+2"
```

#### (*Level) UnmarshalJSON 

``` go
func (l *Level) UnmarshalJSON(data []byte) error
```

UnmarshalJSON implements [encoding/json.Unmarshaler](https://pkg.go.dev/encoding/json#Unmarshaler) It accepts any string produced by [Level.MarshalJSON](https://pkg.go.dev/log/slog@go1.21.3#Level.MarshalJSON), ignoring case. It also accepts numeric offsets that would result in a different string on output. For example, "Error-8" would marshal as "INFO".

​	UnmarshalJSON 实现了 [encoding/json.Unmarshaler](https://pkg.go.dev/encoding/json#Unmarshaler)。它接受由 [Level.MarshalJSON](https://pkg.go.dev/log/slog@go1.21.3#Level.MarshalJSON) 生成的任何字符串，忽略大小写。它还接受将导致不同输出字符串的数值偏移。例如，"Error-8" 将会被序列化为 "INFO"。

#### (*Level) UnmarshalText 

``` go
func (l *Level) UnmarshalText(data []byte) error
```

UnmarshalText implements [encoding.TextUnmarshaler](https://pkg.go.dev/encoding#TextUnmarshaler). It accepts any string produced by [Level.MarshalText](https://pkg.go.dev/log/slog@go1.21.3#Level.MarshalText), ignoring case. It also accepts numeric offsets that would result in a different string on output. For example, "Error-8" would marshal as "INFO".

​	UnmarshalText 通过调用 [Level.UnmarshalText](https://pkg.go.dev/log/slog@go1.21.3#Level.UnmarshalText) 实现了 [encoding.TextUnmarshaler](https://pkg.go.dev/encoding#TextUnmarshaler)。它接受由 [Level.MarshalText](https://pkg.go.dev/log/slog@go1.21.3#Level.MarshalText) 生成的任何字符串，忽略大小写。它还接受将导致不同输出字符串的数值偏移。例如，"Error-8" 将会被序列化为 "INFO"。

### type LevelVar 

``` go
type LevelVar struct {
	// contains filtered or unexported fields
}
```

A LevelVar is a Level variable, to allow a Handler level to change dynamically. It implements Leveler as well as a Set method, and it is safe for use by multiple goroutines. The zero LevelVar corresponds to LevelInfo.

​	LevelVar 是一个 Level 变量，用于动态更改 Handler 的级别。它实现了 Leveler 接口以及 Set 方法，且可安全用于多个协程。初始 LevelVar 对应于 LevelInfo。

#### (*LevelVar) Level 

``` go
func (v *LevelVar) Level() Level
```

Level returns v's level.

​	Level 返回 v 的级别。

#### (*LevelVar) MarshalText 

``` go
func (v *LevelVar) MarshalText() ([]byte, error)
```

MarshalText implements [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler) by calling [Level.MarshalText](https://pkg.go.dev/log/slog@go1.21.3#Level.MarshalText).

​	MarshalText 通过调用 [Level.MarshalText](https://pkg.go.dev/log/slog@go1.21.3#Level.MarshalText) 实现了 [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler)。

#### (*LevelVar) Set 

``` go
func (v *LevelVar) Set(l Level)
```

Set sets v's level to l.

​	Set 设置 v 的级别为 l。

#### (*LevelVar) String 

``` go
func (v *LevelVar) String() string
```

#### (*LevelVar) UnmarshalText 

``` go
func (v *LevelVar) UnmarshalText(data []byte) error
```

UnmarshalText implements [encoding.TextUnmarshaler](https://pkg.go.dev/encoding#TextUnmarshaler) by calling [Level.UnmarshalText](https://pkg.go.dev/log/slog@go1.21.3#Level.UnmarshalText).

​	UnmarshalText 通过调用 [Level.UnmarshalText](https://pkg.go.dev/log/slog@go1.21.3#Level.UnmarshalText) 实现了 [encoding.TextUnmarshaler](https://pkg.go.dev/encoding#TextUnmarshaler)。

### type Leveler 

``` go
type Leveler interface {
	Level() Level
}
```

A Leveler provides a Level value.

​	Leveler 提供了一个 Level 值。As Level itself implements Leveler, clients typically supply a Level value wherever a Leveler is needed, such as in HandlerOptions. Clients who need to vary the level dynamically can provide a more complex Leveler implementation such as *LevelVar.

​	由于 Level 本身实现了 Leveler，客户端通常在需要 Leveler 的地方提供一个 Level 值，例如在 HandlerOptions 中。需要动态变化级别的客户端可以提供更复杂的 Leveler 实现，如 *LevelVar。

### type LogValuer 

``` go
type LogValuer interface {
	LogValue() Value
}
```

A LogValuer is any Go value that can convert itself into a Value for logging.

​	LogValuer 是任何可以将自身转换为日志记录 Value 的 Go 值。

This mechanism may be used to defer expensive operations until they are needed, or to expand a single value into a sequence of components.

​	这个机制可以用于延迟执行昂贵的操作，直到需要时，或者将单个值展开为一系列组件。

#### Example (Group)

``` go
```
#### Example (Secret)

``` go
```
### type Logger 

``` go
type Logger struct {
	// contains filtered or unexported fields
}
```

A Logger records structured information about each call to its Log, Debug, Info, Warn, and Error methods. For each call, it creates a Record and passes it to a Handler.

​	Logger 记录每次对其 Log、Debug、Info、Warn 和 Error 方法的调用的结构化信息。对于每次调用，它都会创建一个 Record 并将其传递给 Handler。

To create a new Logger, call [New](https://pkg.go.dev/log/slog@go1.21.3#New) or a Logger method that begins "With".

​	要创建一个新的 Logger，请调用 [New](https://pkg.go.dev/log/slog@go1.21.3#New) 或者一个以 "With" 开头的 Logger 方法。

#### func Default 

``` go
func Default() *Logger
```

Default returns the default Logger.

​	Default 返回默认的 Logger。

#### func New 

``` go
func New(h Handler) *Logger
```

New creates a new Logger with the given non-nil Handler.

​	New 创建一个带有给定非空 Handler 的新 Logger。

#### func With 

``` go
func With(args ...any) *Logger
```

With calls Logger.With on the default logger.

​	With 在默认的 Logger 上调用 Logger.With。

#### (*Logger) Debug 

``` go
func (l *Logger) Debug(msg string, args ...any)
```

Debug logs at LevelDebug.

​	Debug 在 LevelDebug 级别记录日志。

#### (*Logger) DebugContext 

``` go
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any)
```

DebugContext logs at LevelDebug with the given context.

​	DebugContext 使用给定的上下文在 LevelDebug 级别记录日志。

#### (*Logger) Enabled 

``` go
func (l *Logger) Enabled(ctx context.Context, level Level) bool
```

Enabled reports whether l emits log records at the given context and level.

​	Enabled 报告 l 是否在给定的上下文和级别上发出日志记录。

#### (*Logger) Error 

``` go
func (l *Logger) Error(msg string, args ...any)
```

Error logs at LevelError.

​	Error 在 LevelError 级别记录日志。

#### (*Logger) ErrorContext 

``` go
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any)
```

ErrorContext logs at LevelError with the given context.

​	ErrorContext 使用给定的上下文在 LevelError 级别记录日志。

#### (*Logger) Handler 

``` go
func (l *Logger) Handler() Handler
```

Handler returns l's Handler.

​	Handler 返回 l 的 Handler。

#### (*Logger) Info 

``` go
func (l *Logger) Info(msg string, args ...any)
```

Info logs at LevelInfo.

​	Info 在 LevelInfo 级别记录日志。

#### (*Logger) InfoContext 

``` go
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any)
```

InfoContext logs at LevelInfo with the given context.

​	InfoContext 使用给定的上下文在 LevelInfo 级别记录日志。

#### (*Logger) Log 

``` go
func (l *Logger) Log(ctx context.Context, level Level, msg string, args ...any)
```

Log emits a log record with the current time and the given level and message. The Record's Attrs consist of the Logger's attributes followed by the Attrs specified by args.

​	Log 记录一条包含当前时间的日志，并使用给定的级别和消息。Record 的属性由 Logger 的属性和由 args 指定的属性组成。

The attribute arguments are processed as follows:

​	属性参数的处理如下：

- If an argument is an Attr, it is used as is.
- 如果参数是 Attr，则按原样使用。
- If an argument is a string and this is not the last argument, the following argument is treated as the value and the two are combined into an Attr.
- 如果参数是字符串且不是最后一个参数，则将后面的参数视为值，并将两者组合成一个 Attr。
- Otherwise, the argument is treated as a value with key "!BADKEY".
- 否则，将参数视为键为 "!BADKEY" 的值。

#### (*Logger) LogAttrs 

``` go
func (l *Logger) LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)
```

LogAttrs is a more efficient version of [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) that accepts only Attrs.

​	LogAttrs 是 [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) 的更高效版本，只接受 Attrs。

#### (*Logger) Warn 

``` go
func (l *Logger) Warn(msg string, args ...any)
```

Warn logs at LevelWarn.

​	Warn 在 LevelWarn 级别记录日志。

#### (*Logger) WarnContext 

``` go
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any)
```

WarnContext logs at LevelWarn with the given context.

​	WarnContext 使用给定的上下文在 LevelWarn 级别记录日志。

#### (*Logger) With 

``` go
func (l *Logger) With(args ...any) *Logger
```

With returns a Logger that includes the given attributes in each output operation. Arguments are converted to attributes as if by [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log).

​	With 返回一个 Logger，它在每次输出操作中包含给定的属性。参数将像 [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) 一样被转换为属性。

#### (*Logger) WithGroup 

``` go
func (l *Logger) WithGroup(name string) *Logger
```

WithGroup returns a Logger that starts a group, if name is non-empty. The keys of all attributes added to the Logger will be qualified by the given name. (How that qualification happens depends on the [Handler.WithGroup] method of the Logger's Handler.)

​	WithGroup 返回一个 Logger，如果 name 非空，则启动一个组。添加到 Logger 的所有属性的键都将由给定的名称限定。（如何限定取决于 Logger 的 Handler 的 [Handler.WithGroup] 方法。）

If name is empty, WithGroup returns the receiver.

​	如果 name 为空，WithGroup 返回接收者。

### type Record 

``` go
type Record struct {
	// The time at which the output method (Log, Info, etc.) was called.
    // 输出方法（Log、Info 等）被调用的时间。
	Time time.Time

	// The log message.
    // 日志消息。
	Message string

	// The level of the event.
    // 事件的级别。
	Level Level

	// The program counter at the time the record was constructed, as determined
	// by runtime.Callers. If zero, no program counter is available.
	//
	// The only valid use for this value is as an argument to
	// [runtime.CallersFrames]. In particular, it must not be passed to
	// [runtime.FuncForPC].
    // 构造记录时的程序计数器，由 runtime.Callers 确定。如果为零，则没有可用的程序计数器。
	//
	// 这个值唯一有效的用途是作为 [runtime.CallersFrames] 的参数。特别是，它不能传递给 [runtime.FuncForPC]。
	PC uintptr
	// contains filtered or unexported fields
}
```

A Record holds information about a log event. Copies of a Record share state. Do not modify a Record after handing out a copy to it. Call [NewRecord](https://pkg.go.dev/log/slog@go1.21.3#NewRecord) to create a new Record. Use [Record.Clone](https://pkg.go.dev/log/slog@go1.21.3#Record.Clone) to create a copy with no shared state.

​	Record 保存关于日志事件的信息。Record 的副本共享状态。在分发副本后不要修改 Record。调用 [NewRecord](https://pkg.go.dev/log/slog@go1.21.3#NewRecord) 来创建一个新的 Record。使用 [Record.Clone](https://pkg.go.dev/log/slog@go1.21.3#Record.Clone) 创建一个不共享状态的副本。

#### func NewRecord 

``` go
func NewRecord(t time.Time, level Level, msg string, pc uintptr) Record
```

NewRecord creates a Record from the given arguments. Use [Record.AddAttrs](https://pkg.go.dev/log/slog@go1.21.3#Record.AddAttrs) to add attributes to the Record.

​	NewRecord 使用给定的参数创建一个 Record。使用 [Record.AddAttrs](https://pkg.go.dev/log/slog@go1.21.3#Record.AddAttrs) 向 Record 添加属性。

NewRecord is intended for logging APIs that want to support a [Handler](https://pkg.go.dev/log/slog@go1.21.3#Handler) as a backend.

​	NewRecord 适用于希望支持 [Handler](https://pkg.go.dev/log/slog@go1.21.3#Handler) 作为后端的日志记录 API。

#### (*Record) Add 

``` go
func (r *Record) Add(args ...any)
```

Add converts the args to Attrs as described in [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log), then appends the Attrs to the Record's list of Attrs. It omits empty groups.

​	Add 将 args 转换为 Attrs，如 [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) 中所述，然后将这些 Attrs 附加到 Record 的属性列表中。它会忽略空组。

#### (*Record) AddAttrs 

``` go
func (r *Record) AddAttrs(attrs ...Attr)
```

AddAttrs appends the given Attrs to the Record's list of Attrs. It omits empty groups.

​	AddAttrs 将给定的 Attrs 附加到 Record 的属性列表中。它会忽略空组。

#### (Record) Attrs 

``` go
func (r Record) Attrs(f func(Attr) bool)
```

Attrs calls f on each Attr in the Record. Iteration stops if f returns false.

​	Attrs 对 Record 中的每个 Attr 调用 f。如果 f 返回 false，迭代停止。

#### (Record) Clone 

``` go
func (r Record) Clone() Record
```

Clone returns a copy of the record with no shared state. The original record and the clone can both be modified without interfering with each other.

​	Clone 返回一个没有共享状态的记录副本。原始记录和副本可以在不互相干扰的情况下进行修改。

#### (Record) NumAttrs 

``` go
func (r Record) NumAttrs() int
```

NumAttrs returns the number of attributes in the Record.

​	NumAttrs 返回 Record 中属性的数量。

### type Source 

``` go
type Source struct {
	// Function is the package path-qualified function name containing the
	// source line. If non-empty, this string uniquely identifies a single
	// function in the program. This may be the empty string if not known.
    // Function 是包含源代码行的包路径限定函数名。如果非空，此字符串唯一标识程序中的单个函数。如果未知，这可能是空字符串。
	Function string `json:"function"`
	// File and Line are the file name and line number (1-based) of the source
	// line. These may be the empty string and zero, respectively, if not known.
    // File 和 Line 分别是源代码行的文件名和行号（基于 1）。如果未知，这些可能是空字符串和零。
	File string `json:"file"`
	Line int    `json:"line"`
}
```

Source describes the location of a line of source code.

​	Source 描述源代码行的位置。

### type TextHandler 

``` go
type TextHandler struct {
	// contains filtered or unexported fields
}
```

TextHandler is a Handler that writes Records to an io.Writer as a sequence of key=value pairs separated by spaces and followed by a newline.

​	TextHandler 是一个 Handler，它将 Records 作为键=值对序列写入 io.Writer，每个对之间用空格分隔，并以换行符结尾。

#### func NewTextHandler 

``` go
func NewTextHandler(w io.Writer, opts *HandlerOptions) *TextHandler
```

NewTextHandler creates a TextHandler that writes to w, using the given options. If opts is nil, the default options are used.

​	NewTextHandler 创建一个 TextHandler，将日志记录写入 w，使用给定的选项。如果 opts 为 nil，则使用默认选项。

#### (*TextHandler) Enabled 

``` go
func (h *TextHandler) Enabled(_ context.Context, level Level) bool
```

Enabled reports whether the handler handles records at the given level. The handler ignores records whose level is lower.

​	Enabled 报告处理器是否处理给定级别的记录。处理器会忽略级别较低的记录。

#### (*TextHandler) Handle 

``` go
func (h *TextHandler) Handle(_ context.Context, r Record) error
```

Handle formats its argument Record as a single line of space-separated key=value items.

​	Handle 将其参数 Record 格式化为一行由空格分隔的key=value项。

If the Record's time is zero, the time is omitted. Otherwise, the key is "time" and the value is output in RFC3339 format with millisecond precision.

​	如果 Record 的时间为零，则省略时间。否则，键为 "time"，值以毫秒精度输出为 RFC3339 格式。

If the Record's level is zero, the level is omitted. Otherwise, the key is "level" and the value of [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String) is output.

​	如果 Record 的级别为零，则省略级别。否则，键为 "level"，值为 [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String) 的输出。

If the AddSource option is set and source information is available, the key is "source" and the value is output as FILE:LINE.

​	如果设置了 AddSource 选项并且有可用的源信息，键为 "source"，值以 FILE:LINE 的格式输出。

The message's key is "msg".

​	消息的键为 "msg"。

To modify these or other attributes, or remove them from the output, use [HandlerOptions.ReplaceAttr].

​	要修改这些或其他属性，或将它们从输出中删除，请使用 [HandlerOptions.ReplaceAttr]。

If a value implements [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler), the result of MarshalText is written. Otherwise, the result of fmt.Sprint is written.

​	如果值实现了 [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler)，则写入 MarshalText 的结果。否则，写入 fmt.Sprint 的结果。

Keys and values are quoted with [strconv.Quote](https://pkg.go.dev/strconv#Quote) if they contain Unicode space characters, non-printing characters, '"' or '='.

​	如果键或值包含 Unicode 空格字符、不可打印字符、'"' 或 '='，则使用 [strconv.Quote](https://pkg.go.dev/strconv#Quote) 进行引用。

Keys inside groups consist of components (keys or group names) separated by dots. No further escaping is performed. Thus there is no way to determine from the key "a.b.c" whether there are two groups "a" and "b" and a key "c", or a single group "a.b" and a key "c", or single group "a" and a key "b.c". If it is necessary to reconstruct the group structure of a key even in the presence of dots inside components, use [HandlerOptions.ReplaceAttr] to encode that information in the key.

​	组内的键由组件（键或组名）通过点分隔组成。不会进行进一步的转义。因此，从键 "a.b.c" 无法确定是两个组 "a" 和 "b" 以及一个键 "c"，还是一个组 "a.b" 和一个键 "c"，或者是一个组 "a" 和一个键 "b.c"。如果需要在组件内部存在点的情况下重建键的组结构，可以使用 [HandlerOptions.ReplaceAttr] 来在键中编码该信息。

Each call to Handle results in a single serialized call to io.Writer.Write.

​	每次调用 Handle 方法都会导致一次对 io.Writer.Write 的单次序列化调用。

#### (*TextHandler) WithAttrs 

``` go
func (h *TextHandler) WithAttrs(attrs []Attr) Handler
```

WithAttrs returns a new TextHandler whose attributes consists of h's attributes followed by attrs.

​	WithAttrs 返回一个新的 TextHandler，其属性由 h 的属性加上 attrs 组成。

#### (*TextHandler) WithGroup 

``` go
func (h *TextHandler) WithGroup(name string) Handler
```

### type Value 

``` go
type Value struct {
	// contains filtered or unexported fields
}
```

A Value can represent any Go value, but unlike type any, it can represent most small values without an allocation. The zero Value corresponds to nil.

​	Value 可以表示任何 Go 值，但与类型 any 不同，它可以在不进行分配的情况下表示大多数小值。零值的 Value 对应于 nil。

#### func AnyValue 

``` go
func AnyValue(v any) Value
```

AnyValue returns a Value for the supplied value.

​	AnyValue 返回提供值的 Value。

If the supplied value is of type Value, it is returned unmodified.

​	如果提供的值是 Value 类型，则返回未修改的值。

Given a value of one of Go's predeclared string, bool, or (non-complex) numeric types, AnyValue returns a Value of kind String, Bool, Uint64, Int64, or Float64. The width of the original numeric type is not preserved.

​	对于 Go 的预定义字符串、布尔值或（非复数）数字类型的值，AnyValue 返回类型为 String、Bool、Uint64、Int64 或 Float64 的 Value。原始数字类型的宽度不会被保留。

Given a time.Time or time.Duration value, AnyValue returns a Value of kind KindTime or KindDuration. The monotonic time is not preserved.

​	对于 time.Time 或 time.Duration 值，AnyValue 返回类型为 KindTime 或 KindDuration 的 Value。单调时间不会被保留。

For nil, or values of all other types, including named types whose underlying type is numeric, AnyValue returns a value of kind KindAny.

​	对于 nil 或所有其他类型的值，包括基础类型为数字的命名类型，AnyValue 返回类型为 KindAny 的 Value。

#### func BoolValue 

``` go
func BoolValue(v bool) Value
```

BoolValue returns a Value for a bool.

​	BoolValue 返回布尔值的 Value。

#### func DurationValue 

``` go
func DurationValue(v time.Duration) Value
```

DurationValue returns a Value for a time.Duration.

​	DurationValue 返回时间间隔的 Value。

#### func Float64Value 

``` go
func Float64Value(v float64) Value
```

Float64Value returns a Value for a floating-point number.

​	Float64Value 返回浮点数的 Value。

#### func GroupValue 

``` go
func GroupValue(as ...Attr) Value
```

GroupValue returns a new Value for a list of Attrs. The caller must not subsequently mutate the argument slice.

​	GroupValue 返回一个新的 Value，用于表示 Attrs 列表。调用者不得随后修改参数切片。

#### func Int64Value 

``` go
func Int64Value(v int64) Value
```

Int64Value returns a Value for an int64.

​	Int64Value 返回 int64 的 Value。

#### func IntValue 

``` go
func IntValue(v int) Value
```

IntValue returns a Value for an int.

​	IntValue 返回 int 的 Value。

#### func StringValue 

``` go
func StringValue(value string) Value
```

StringValue returns a new Value for a string.

​	StringValue 返回字符串的 Value。

#### func TimeValue 

``` go
func TimeValue(v time.Time) Value
```

TimeValue returns a Value for a time.Time. It discards the monotonic portion.

​	TimeValue 返回 time.Time 的 Value。它会丢弃单调部分。

#### func Uint64Value 

``` go
func Uint64Value(v uint64) Value
```

Uint64Value returns a Value for a uint64.

​	Uint64Value 返回 uint64 的 Value。

#### (Value) Any 

``` go
func (v Value) Any() any
```

Any returns v's value as an any.

​	Any 返回 v 的值作为 any。

#### (Value) Bool 

``` go
func (v Value) Bool() bool
```

Bool returns v's value as a bool. It panics if v is not a bool.

​	Bool 返回 v 的值作为布尔值。如果 v 不是布尔值，则会引发恐慌。

#### (Value) Duration 

``` go
func (a Value) Duration() time.Duration
```

Duration returns v's value as a time.Duration. It panics if v is not a time.Duration.

​	Duration 返回 v 的值作为时间间隔。如果 v 不是时间间隔，则会引发恐慌。

#### (Value) Equal 

``` go
func (v Value) Equal(w Value) bool
```

Equal reports whether v and w represent the same Go value.

​	Equal 报告 v 和 w 是否表示相同的 Go 值。

#### (Value) Float64 

``` go
func (v Value) Float64() float64
```

Float64 returns v's value as a float64. It panics if v is not a float64.

​	Float64 返回 v 的值作为 float64。如果 v 不是 float64，则会引发恐慌。

#### (Value) Group 

``` go
func (v Value) Group() []Attr
```

Group returns v's value as a []Attr. It panics if v's Kind is not KindGroup.

​	Group 返回 v 的值作为 []Attr。如果 v 的 Kind 不是 KindGroup，则会引发恐慌。

#### (Value) Int64 

``` go
func (v Value) Int64() int64
```

Int64 returns v's value as an int64. It panics if v is not a signed integer.

​	Int64 返回 v 的值作为 int64。如果 v 不是有符号整数，则会引发恐慌。

#### (Value) Kind 

``` go
func (v Value) Kind() Kind
```

Kind returns v's Kind.

​	Kind 返回 v 的 Kind。

#### (Value) LogValuer 

``` go
func (v Value) LogValuer() LogValuer
```

LogValuer returns v's value as a LogValuer. It panics if v is not a LogValuer.

​	LogValuer 返回 v 的值作为 LogValuer。如果 v 不是 LogValuer，则会引发恐慌。

#### (Value) Resolve 

``` go
func (v Value) Resolve() (rv Value)
```

Resolve repeatedly calls LogValue on v while it implements LogValuer, and returns the result. If v resolves to a group, the group's attributes' values are not recursively resolved. If the number of LogValue calls exceeds a threshold, a Value containing an error is returned. Resolve's return value is guaranteed not to be of Kind KindLogValuer.

​	Resolve 在 v 实现 LogValuer 的情况下重复调用 LogValue，并返回结果。如果 v 解析为一个组，该组的属性值不会递归解析。如果 LogValue 调用的次数超过阈值，则返回一个包含错误的 Value。Resolve 的返回值保证不是 Kind 为 KindLogValuer 的值。

#### (Value) String 

``` go
func (v Value) String() string
```

String returns Value's value as a string, formatted like fmt.Sprint. Unlike the methods Int64, Float64, and so on, which panic if v is of the wrong kind, String never panics.

​	String 返回 Value 的值作为字符串，格式类似 fmt.Sprint。与 Int64、Float64 等方法不同，即使 v 的类型不正确，String 方法也不会引发恐慌。

#### (Value) Time 

``` go
func (v Value) Time() time.Time
```

Time returns v's value as a time.Time. It panics if v is not a time.Time.

​	Time 返回 v 的值作为 time.Time。如果 v 不是 time.Time，则会引发恐慌。

#### (Value) Uint64 

``` go
func (v Value) Uint64() uint64
```

Uint64 returns v's value as a uint64. It panics if v is not an unsigned integer.

​	Uint64 返回 v 的值作为 uint64。如果 v 不是无符号整数，则会引发恐慌。