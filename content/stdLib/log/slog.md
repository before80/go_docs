+++
title = "slog"
date = 2023-11-05T14:27:24+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文：[https://pkg.go.dev/slog@go1.21.3](https://pkg.go.dev/slog@go1.21.3)

## Overview 

Package slog provides structured logging, in which log records include a message, a severity level, and various other attributes expressed as key-value pairs.

It defines a type, [Logger](https://pkg.go.dev/log/slog@go1.21.3#Logger), which provides several methods (such as [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info) and [Logger.Error](https://pkg.go.dev/log/slog@go1.21.3#Logger.Error)) for reporting events of interest.

Each Logger is associated with a [Handler](https://pkg.go.dev/log/slog@go1.21.3#Handler). A Logger output method creates a [Record](https://pkg.go.dev/log/slog@go1.21.3#Record) from the method arguments and passes it to the Handler, which decides how to handle it. There is a default Logger accessible through top-level functions (such as [Info](https://pkg.go.dev/log/slog@go1.21.3#Info) and [Error](https://pkg.go.dev/log/slog@go1.21.3#Error)) that call the corresponding Logger methods.

A log record consists of a time, a level, a message, and a set of key-value pairs, where the keys are strings and the values may be of any type. As an example,

```
slog.Info("hello", "count", 3)
```

creates a record containing the time of the call, a level of Info, the message "hello", and a single pair with key "count" and value 3.

The [Info](https://pkg.go.dev/log/slog@go1.21.3#Info) top-level function calls the [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info) method on the default Logger. In addition to [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info), there are methods for Debug, Warn and Error levels. Besides these convenience methods for common levels, there is also a [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) method which takes the level as an argument. Each of these methods has a corresponding top-level function that uses the default logger.

The default handler formats the log record's message, time, level, and attributes as a string and passes it to the [log](https://pkg.go.dev/log) package.

```
2022/11/08 15:28:26 INFO hello count=3
```

For more control over the output format, create a logger with a different handler. This statement uses [New](https://pkg.go.dev/log/slog@go1.21.3#New) to create a new logger with a TextHandler that writes structured records in text form to standard error:

```
logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
```

[TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) output is a sequence of key=value pairs, easily and unambiguously parsed by machine. This statement:

```
logger.Info("hello", "count", 3)
```

produces this output:

```
time=2022-11-08T15:28:26.000-05:00 level=INFO msg=hello count=3
```

The package also provides [JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler), whose output is line-delimited JSON:

```
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("hello", "count", 3)
```

produces this output:

```
{"time":"2022-11-08T15:28:26.000000000-05:00","level":"INFO","msg":"hello","count":3}
```

Both [TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) and [JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler) can be configured with [HandlerOptions](https://pkg.go.dev/log/slog@go1.21.3#HandlerOptions). There are options for setting the minimum level (see Levels, below), displaying the source file and line of the log call, and modifying attributes before they are logged.

Setting a logger as the default with

```
slog.SetDefault(logger)
```

will cause the top-level functions like [Info](https://pkg.go.dev/log/slog@go1.21.3#Info) to use it. [SetDefault](https://pkg.go.dev/log/slog@go1.21.3#SetDefault) also updates the default logger used by the [log](https://pkg.go.dev/log) package, so that existing applications that use [log.Printf](https://pkg.go.dev/log#Printf) and related functions will send log records to the logger's handler without needing to be rewritten.

Some attributes are common to many log calls. For example, you may wish to include the URL or trace identifier of a server request with all log events arising from the request. Rather than repeat the attribute with every log call, you can use [Logger.With](https://pkg.go.dev/log/slog@go1.21.3#Logger.With) to construct a new Logger containing the attributes:

```
logger2 := logger.With("url", r.URL)
```

The arguments to With are the same key-value pairs used in [Logger.Info](https://pkg.go.dev/log/slog@go1.21.3#Logger.Info). The result is a new Logger with the same handler as the original, but additional attributes that will appear in the output of every call.

### Levels 

A [Level](https://pkg.go.dev/log/slog@go1.21.3#Level) is an integer representing the importance or severity of a log event. The higher the level, the more severe the event. This package defines constants for the most common levels, but any int can be used as a level.

In an application, you may wish to log messages only at a certain level or greater. One common configuration is to log messages at Info or higher levels, suppressing debug logging until it is needed. The built-in handlers can be configured with the minimum level to output by setting [HandlerOptions.Level]. The program's `main` function typically does this. The default value is LevelInfo.

Setting the [HandlerOptions.Level] field to a [Level](https://pkg.go.dev/log/slog@go1.21.3#Level) value fixes the handler's minimum level throughout its lifetime. Setting it to a [LevelVar](https://pkg.go.dev/log/slog@go1.21.3#LevelVar) allows the level to be varied dynamically. A LevelVar holds a Level and is safe to read or write from multiple goroutines. To vary the level dynamically for an entire program, first initialize a global LevelVar:

``` go
var programLevel = new(slog.LevelVar) // Info by default
```

Then use the LevelVar to construct a handler, and make it the default:

```
h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
slog.SetDefault(slog.New(h))
```

Now the program can change its logging level with a single statement:

```
programLevel.Set(slog.LevelDebug)
```

### Groups 

Attributes can be collected into groups. A group has a name that is used to qualify the names of its attributes. How this qualification is displayed depends on the handler. [TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) separates the group and attribute names with a dot. [JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler) treats each group as a separate JSON object, with the group name as the key.

Use [Group](https://pkg.go.dev/log/slog@go1.21.3#Group) to create a Group attribute from a name and a list of key-value pairs:

```
slog.Group("request",
    "method", r.Method,
    "url", r.URL)
```

TextHandler would display this group as

```
request.method=GET request.url=http://example.com
```

JSONHandler would display it as

```
"request":{"method":"GET","url":"http://example.com"}
```

Use [Logger.WithGroup](https://pkg.go.dev/log/slog@go1.21.3#Logger.WithGroup) to qualify all of a Logger's output with a group name. Calling WithGroup on a Logger results in a new Logger with the same Handler as the original, but with all its attributes qualified by the group name.

This can help prevent duplicate attribute keys in large systems, where subsystems might use the same keys. Pass each subsystem a different Logger with its own group name so that potential duplicates are qualified:

```
logger := slog.Default().With("id", systemID)
parserLogger := logger.WithGroup("parser")
parseInput(input, parserLogger)
```

When parseInput logs with parserLogger, its keys will be qualified with "parser", so even if it uses the common key "id", the log line will have distinct keys.

### Contexts 

Some handlers may wish to include information from the [context.Context](https://pkg.go.dev/context#Context) that is available at the call site. One example of such information is the identifier for the current span when tracing is enabled.

The [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) and [Logger.LogAttrs](https://pkg.go.dev/log/slog@go1.21.3#Logger.LogAttrs) methods take a context as a first argument, as do their corresponding top-level functions.

Although the convenience methods on Logger (Info and so on) and the corresponding top-level functions do not take a context, the alternatives ending in "Context" do. For example,

```
slog.InfoContext(ctx, "message")
```

It is recommended to pass a context to an output method if one is available.

### Attrs and Values 

An [Attr](https://pkg.go.dev/log/slog@go1.21.3#Attr) is a key-value pair. The Logger output methods accept Attrs as well as alternating keys and values. The statement

```
slog.Info("hello", slog.Int("count", 3))
```

behaves the same as

```
slog.Info("hello", "count", 3)
```

There are convenience constructors for [Attr](https://pkg.go.dev/log/slog@go1.21.3#Attr) such as [Int](https://pkg.go.dev/log/slog@go1.21.3#Int), [String](https://pkg.go.dev/log/slog@go1.21.3#String), and [Bool](https://pkg.go.dev/log/slog@go1.21.3#Bool) for common types, as well as the function [Any](https://pkg.go.dev/log/slog@go1.21.3#Any) for constructing Attrs of any type.

The value part of an Attr is a type called [Value](https://pkg.go.dev/log/slog@go1.21.3#Value). Like an [any], a Value can hold any Go value, but it can represent typical values, including all numbers and strings, without an allocation.

For the most efficient log output, use [Logger.LogAttrs](https://pkg.go.dev/log/slog@go1.21.3#Logger.LogAttrs). It is similar to [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) but accepts only Attrs, not alternating keys and values; this allows it, too, to avoid allocation.

The call

```
logger.LogAttrs(ctx, slog.LevelInfo, "hello", slog.Int("count", 3))
```

is the most efficient way to achieve the same output as

```
slog.Info("hello", "count", 3)
```

### Customizing a type's logging behavior 

If a type implements the [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer) interface, the [Value](https://pkg.go.dev/log/slog@go1.21.3#Value) returned from its LogValue method is used for logging. You can use this to control how values of the type appear in logs. For example, you can redact secret information like passwords, or gather a struct's fields in a Group. See the examples under [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer) for details.

A LogValue method may return a Value that itself implements [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer). The [Value.Resolve](https://pkg.go.dev/log/slog@go1.21.3#Value.Resolve) method handles these cases carefully, avoiding infinite loops and unbounded recursion. Handler authors and others may wish to use Value.Resolve instead of calling LogValue directly.

### Wrapping output methods 

The logger functions use reflection over the call stack to find the file name and line number of the logging call within the application. This can produce incorrect source information for functions that wrap slog. For instance, if you define this function in file mylog.go:

``` go
func Infof(format string, args ...any) {
    slog.Default().Info(fmt.Sprintf(format, args...))
}
```

and you call it like this in main.go:

```
Infof(slog.Default(), "hello, %s", "world")
```

then slog will report the source file as mylog.go, not main.go.

A correct implementation of Infof will obtain the source location (pc) and pass it to NewRecord. The Infof function in the package-level example called "wrapping" demonstrates how to do this.

### Working with Records 

Sometimes a Handler will need to modify a Record before passing it on to another Handler or backend. A Record contains a mixture of simple public fields (e.g. Time, Level, Message) and hidden fields that refer to state (such as attributes) indirectly. This means that modifying a simple copy of a Record (e.g. by calling [Record.Add](https://pkg.go.dev/log/slog@go1.21.3#Record.Add) or [Record.AddAttrs](https://pkg.go.dev/log/slog@go1.21.3#Record.AddAttrs) to add attributes) may have unexpected effects on the original. Before modifying a Record, use [Record.Clone](https://pkg.go.dev/log/slog@go1.21.3#Record.Clone) to create a copy that shares no state with the original, or create a new Record with [NewRecord](https://pkg.go.dev/log/slog@go1.21.3#NewRecord) and build up its Attrs by traversing the old ones with [Record.Attrs](https://pkg.go.dev/log/slog@go1.21.3#Record.Attrs).

### Performance considerations 

If profiling your application demonstrates that logging is taking significant time, the following suggestions may help.

If many log lines have a common attribute, use [Logger.With](https://pkg.go.dev/log/slog@go1.21.3#Logger.With) to create a Logger with that attribute. The built-in handlers will format that attribute only once, at the call to [Logger.With](https://pkg.go.dev/log/slog@go1.21.3#Logger.With). The [Handler](https://pkg.go.dev/log/slog@go1.21.3#Handler) interface is designed to allow that optimization, and a well-written Handler should take advantage of it.

The arguments to a log call are always evaluated, even if the log event is discarded. If possible, defer computation so that it happens only if the value is actually logged. For example, consider the call

```
slog.Info("starting request", "url", r.URL.String())  // may compute String unnecessarily
```

The URL.String method will be called even if the logger discards Info-level events. Instead, pass the URL directly:

```
slog.Info("starting request", "url", &r.URL) // calls URL.String only if needed
```

The built-in [TextHandler](https://pkg.go.dev/log/slog@go1.21.3#TextHandler) will call its String method, but only if the log event is enabled. Avoiding the call to String also preserves the structure of the underlying value. For example [JSONHandler](https://pkg.go.dev/log/slog@go1.21.3#JSONHandler) emits the components of the parsed URL as a JSON object. If you want to avoid eagerly paying the cost of the String call without causing the handler to potentially inspect the structure of the value, wrap the value in a fmt.Stringer implementation that hides its Marshal methods.

You can also use the [LogValuer](https://pkg.go.dev/log/slog@go1.21.3#LogValuer) interface to avoid unnecessary work in disabled log calls. Say you need to log some expensive value:

```
slog.Debug("frobbing", "value", computeExpensiveValue(arg))
```

Even if this line is disabled, computeExpensiveValue will be called. To avoid that, define a type implementing LogValuer:

``` go
type expensive struct { arg int }

func (e expensive) LogValue() slog.Value {
    return slog.AnyValue(computeExpensiveValue(e.arg))
}
```

Then use a value of that type in log calls:

```
slog.Debug("frobbing", "value", expensive{arg})
```

Now computeExpensiveValue will only be called when the line is enabled.

The built-in handlers acquire a lock before calling [io.Writer.Write](https://pkg.go.dev/io#Writer.Write) to ensure that each record is written in one piece. User-defined handlers are responsible for their own locking.

### Writing a handler 

For a guide to writing a custom handler, see https://golang.org/s/slog-handler-guide.

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


## Constants 

[View Source](https://cs.opensource.google/go/go/+/go1.21.3:src/log/slog/handler.go;l=170)

``` go
const (
	// TimeKey is the key used by the built-in handlers for the time
	// when the log method is called. The associated Value is a [time.Time].
	TimeKey = "time"
	// LevelKey is the key used by the built-in handlers for the level
	// of the log call. The associated value is a [Level].
	LevelKey = "level"
	// MessageKey is the key used by the built-in handlers for the
	// message of the log call. The associated value is a string.
	MessageKey = "msg"
	// SourceKey is the key used by the built-in handlers for the source file
	// and line of the log call. The associated value is a string.
	SourceKey = "source"
)
```

Keys for "built-in" attributes.

## Variables 

This section is empty.

## Functions 

### func Debug 

``` go
func Debug(msg string, args ...any)
```

Debug calls Logger.Debug on the default logger.

### func DebugContext 

``` go
func DebugContext(ctx context.Context, msg string, args ...any)
```

DebugContext calls Logger.DebugContext on the default logger.

### func Error 

``` go
func Error(msg string, args ...any)
```

Error calls Logger.Error on the default logger.

### func ErrorContext 

``` go
func ErrorContext(ctx context.Context, msg string, args ...any)
```

ErrorContext calls Logger.ErrorContext on the default logger.

### func Info 

``` go
func Info(msg string, args ...any)
```

Info calls Logger.Info on the default logger.

### func InfoContext 

``` go
func InfoContext(ctx context.Context, msg string, args ...any)
```

InfoContext calls Logger.InfoContext on the default logger.

### func Log 

``` go
func Log(ctx context.Context, level Level, msg string, args ...any)
```

Log calls Logger.Log on the default logger.

### func LogAttrs 

``` go
func LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)
```

LogAttrs calls Logger.LogAttrs on the default logger.

### func NewLogLogger 

``` go
func NewLogLogger(h Handler, level Level) *log.Logger
```

NewLogLogger returns a new log.Logger such that each call to its Output method dispatches a Record to the specified handler. The logger acts as a bridge from the older log API to newer structured logging handlers.

### func SetDefault 

``` go
func SetDefault(l *Logger)
```

SetDefault makes l the default Logger. After this call, output from the log package's default Logger (as with [log.Print](https://pkg.go.dev/log#Print), etc.) will be logged at LevelInfo using l's Handler.

### func Warn 

``` go
func Warn(msg string, args ...any)
```

Warn calls Logger.Warn on the default logger.

### func WarnContext 

``` go
func WarnContext(ctx context.Context, msg string, args ...any)
```

WarnContext calls Logger.WarnContext on the default logger.

## Types 

### type Attr 

``` go
type Attr struct {
	Key   string
	Value Value
}
```

An Attr is a key-value pair.

#### func Any 

``` go
func Any(key string, value any) Attr
```

Any returns an Attr for the supplied value. See [AnyValue](https://pkg.go.dev/log/slog@go1.21.3#AnyValue) for how values are treated.

#### func Bool 

``` go
func Bool(key string, v bool) Attr
```

Bool returns an Attr for a bool.

#### func Duration 

``` go
func Duration(key string, v time.Duration) Attr
```

Duration returns an Attr for a time.Duration.

#### func Float64 

``` go
func Float64(key string, v float64) Attr
```

Float64 returns an Attr for a floating-point number.

#### func Group 

``` go
func Group(key string, args ...any) Attr
```

Group returns an Attr for a Group Value. The first argument is the key; the remaining arguments are converted to Attrs as in [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log).

Use Group to collect several key-value pairs under a single key on a log line, or as the result of LogValue in order to log a single value as multiple Attrs.

##### Example

``` go
```
#### func Int 

``` go
func Int(key string, value int) Attr
```

Int converts an int to an int64 and returns an Attr with that value.

#### func Int64 

``` go
func Int64(key string, value int64) Attr
```

Int64 returns an Attr for an int64.

#### func String 

``` go
func String(key, value string) Attr
```

String returns an Attr for a string value.

#### func Time 

``` go
func Time(key string, v time.Time) Attr
```

Time returns an Attr for a time.Time. It discards the monotonic portion.

#### func Uint64 

``` go
func Uint64(key string, v uint64) Attr
```

Uint64 returns an Attr for a uint64.

#### (Attr) Equal 

``` go
func (a Attr) Equal(b Attr) bool
```

Equal reports whether a and b have equal keys and values.

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
	Handle(context.Context, Record) error

	// WithAttrs returns a new Handler whose attributes consist of
	// both the receiver's attributes and the arguments.
	// The Handler owns the slice: it may retain, modify or discard it.
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
	WithGroup(name string) Handler
}
```

A Handler handles log records produced by a Logger..

A typical handler may print log records to standard error, or write them to a file or database, or perhaps augment them with additional attributes and pass them on to another handler.

Any of the Handler's methods may be called concurrently with itself or with other methods. It is the responsibility of the Handler to manage this concurrency.

Users of the slog package should not invoke Handler methods directly. They should use the methods of [Logger](https://pkg.go.dev/log/slog@go1.21.3#Logger) instead.

### Example (LevelHandler)

``` go
```
### type HandlerOptions 

``` go
type HandlerOptions struct {
	// AddSource causes the handler to compute the source code position
	// of the log statement and add a SourceKey attribute to the output.
	AddSource bool

	// Level reports the minimum record level that will be logged.
	// The handler discards records with lower levels.
	// If Level is nil, the handler assumes LevelInfo.
	// The handler calls Level.Level for each record processed;
	// to adjust the minimum level dynamically, use a LevelVar.
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
	ReplaceAttr func(groups []string, a Attr) Attr
}
```

HandlerOptions are options for a TextHandler or JSONHandler. A zero HandlerOptions consists entirely of default values.

#### Example (CustomLevels)

``` go
```
### type JSONHandler 

``` go
type JSONHandler struct {
	// contains filtered or unexported fields
}
```

JSONHandler is a Handler that writes Records to an io.Writer as line-delimited JSON objects.

#### func NewJSONHandler 

``` go
func NewJSONHandler(w io.Writer, opts *HandlerOptions) *JSONHandler
```

NewJSONHandler creates a JSONHandler that writes to w, using the given options. If opts is nil, the default options are used.

#### (*JSONHandler) Enabled 

``` go
func (h *JSONHandler) Enabled(_ context.Context, level Level) bool
```

Enabled reports whether the handler handles records at the given level. The handler ignores records whose level is lower.

#### (*JSONHandler) Handle 

``` go
func (h *JSONHandler) Handle(_ context.Context, r Record) error
```

Handle formats its argument Record as a JSON object on a single line.

If the Record's time is zero, the time is omitted. Otherwise, the key is "time" and the value is output as with json.Marshal.

If the Record's level is zero, the level is omitted. Otherwise, the key is "level" and the value of [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String) is output.

If the AddSource option is set and source information is available, the key is "source", and the value is a record of type [Source](https://pkg.go.dev/log/slog@go1.21.3#Source).

The message's key is "msg".

To modify these or other attributes, or remove them from the output, use [HandlerOptions.ReplaceAttr].

Values are formatted as with an [encoding/json.Encoder](https://pkg.go.dev/encoding/json#Encoder) with SetEscapeHTML(false), with two exceptions.

First, an Attr whose Value is of type error is formatted as a string, by calling its Error method. Only errors in Attrs receive this special treatment, not errors embedded in structs, slices, maps or other data structures that are processed by the encoding/json package.

Second, an encoding failure does not cause Handle to return an error. Instead, the error message is formatted as a string.

Each call to Handle results in a single serialized call to io.Writer.Write.

#### (*JSONHandler) WithAttrs 

``` go
func (h *JSONHandler) WithAttrs(attrs []Attr) Handler
```

WithAttrs returns a new JSONHandler whose attributes consists of h's attributes followed by attrs.

#### (*JSONHandler) WithGroup 

``` go
func (h *JSONHandler) WithGroup(name string) Handler
```

### type Kind 

``` go
type Kind int
```

Kind is the kind of a Value.

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

#### type Level 

``` go
type Level int
```

A Level is the importance or severity of a log event. The higher the level, the more important or severe the event.

``` go
const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)
```

Level numbers are inherently arbitrary, but we picked them to satisfy three constraints. Any system can map them to another numbering scheme if it wishes.

First, we wanted the default level to be Info, Since Levels are ints, Info is the default value for int, zero.

Second, we wanted to make it easy to use levels to specify logger verbosity. Since a larger level means a more severe event, a logger that accepts events with smaller (or more negative) level means a more verbose logger. Logger verbosity is thus the negation of event severity, and the default verbosity of 0 accepts all events at least as severe as INFO.

Third, we wanted some room between levels to accommodate schemes with named levels between ours. For example, Google Cloud Logging defines a Notice level between Info and Warn. Since there are only a few of these intermediate levels, the gap between the numbers need not be large. Our gap of 4 matches OpenTelemetry's mapping. Subtracting 9 from an OpenTelemetry level in the DEBUG, INFO, WARN and ERROR ranges converts it to the corresponding slog Level range. OpenTelemetry also has the names TRACE and FATAL, which slog does not. But those OpenTelemetry levels can still be represented as slog Levels by using the appropriate integers.

Names for common levels.

#### (Level) Level 

``` go
func (l Level) Level() Level
```

Level returns the receiver. It implements Leveler.

#### (Level) MarshalJSON 

``` go
func (l Level) MarshalJSON() ([]byte, error)
```

MarshalJSON implements [encoding/json.Marshaler](https://pkg.go.dev/encoding/json#Marshaler) by quoting the output of [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String).

#### (Level) MarshalText 

``` go
func (l Level) MarshalText() ([]byte, error)
```

MarshalText implements [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler) by calling [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String).

#### (Level) String 

``` go
func (l Level) String() string
```

String returns a name for the level. If the level has a name, then that name in uppercase is returned. If the level is between named values, then an integer is appended to the uppercased name. Examples:

```
LevelWarn.String() => "WARN"
(LevelInfo+2).String() => "INFO+2"
```

#### (*Level) UnmarshalJSON 

``` go
func (l *Level) UnmarshalJSON(data []byte) error
```

UnmarshalJSON implements [encoding/json.Unmarshaler](https://pkg.go.dev/encoding/json#Unmarshaler) It accepts any string produced by [Level.MarshalJSON](https://pkg.go.dev/log/slog@go1.21.3#Level.MarshalJSON), ignoring case. It also accepts numeric offsets that would result in a different string on output. For example, "Error-8" would marshal as "INFO".

#### (*Level) UnmarshalText 

``` go
func (l *Level) UnmarshalText(data []byte) error
```

UnmarshalText implements [encoding.TextUnmarshaler](https://pkg.go.dev/encoding#TextUnmarshaler). It accepts any string produced by [Level.MarshalText](https://pkg.go.dev/log/slog@go1.21.3#Level.MarshalText), ignoring case. It also accepts numeric offsets that would result in a different string on output. For example, "Error-8" would marshal as "INFO".

### type LevelVar 

``` go
type LevelVar struct {
	// contains filtered or unexported fields
}
```

A LevelVar is a Level variable, to allow a Handler level to change dynamically. It implements Leveler as well as a Set method, and it is safe for use by multiple goroutines. The zero LevelVar corresponds to LevelInfo.

#### (*LevelVar) Level 

``` go
func (v *LevelVar) Level() Level
```

Level returns v's level.

#### (*LevelVar) MarshalText 

``` go
func (v *LevelVar) MarshalText() ([]byte, error)
```

MarshalText implements [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler) by calling [Level.MarshalText](https://pkg.go.dev/log/slog@go1.21.3#Level.MarshalText).

#### (*LevelVar) Set 

``` go
func (v *LevelVar) Set(l Level)
```

Set sets v's level to l.

#### (*LevelVar) String 

``` go
func (v *LevelVar) String() string
```

#### (*LevelVar) UnmarshalText 

``` go
func (v *LevelVar) UnmarshalText(data []byte) error
```

UnmarshalText implements [encoding.TextUnmarshaler](https://pkg.go.dev/encoding#TextUnmarshaler) by calling [Level.UnmarshalText](https://pkg.go.dev/log/slog@go1.21.3#Level.UnmarshalText).

### type Leveler 

``` go
type Leveler interface {
	Level() Level
}
```

A Leveler provides a Level value.

As Level itself implements Leveler, clients typically supply a Level value wherever a Leveler is needed, such as in HandlerOptions. Clients who need to vary the level dynamically can provide a more complex Leveler implementation such as *LevelVar.

### type LogValuer 

``` go
type LogValuer interface {
	LogValue() Value
}
```

A LogValuer is any Go value that can convert itself into a Value for logging.

This mechanism may be used to defer expensive operations until they are needed, or to expand a single value into a sequence of components.

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

To create a new Logger, call [New](https://pkg.go.dev/log/slog@go1.21.3#New) or a Logger method that begins "With".

#### func Default 

``` go
func Default() *Logger
```

Default returns the default Logger.

#### func New 

``` go
func New(h Handler) *Logger
```

New creates a new Logger with the given non-nil Handler.

#### func With 

``` go
func With(args ...any) *Logger
```

With calls Logger.With on the default logger.

#### (*Logger) Debug 

``` go
func (l *Logger) Debug(msg string, args ...any)
```

Debug logs at LevelDebug.

#### (*Logger) DebugContext 

``` go
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any)
```

DebugContext logs at LevelDebug with the given context.

#### (*Logger) Enabled 

``` go
func (l *Logger) Enabled(ctx context.Context, level Level) bool
```

Enabled reports whether l emits log records at the given context and level.

#### (*Logger) Error 

``` go
func (l *Logger) Error(msg string, args ...any)
```

Error logs at LevelError.

#### (*Logger) ErrorContext 

``` go
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any)
```

ErrorContext logs at LevelError with the given context.

#### (*Logger) Handler 

``` go
func (l *Logger) Handler() Handler
```

Handler returns l's Handler.

#### (*Logger) Info 

``` go
func (l *Logger) Info(msg string, args ...any)
```

Info logs at LevelInfo.

#### (*Logger) InfoContext 

``` go
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any)
```

InfoContext logs at LevelInfo with the given context.

#### (*Logger) Log 

``` go
func (l *Logger) Log(ctx context.Context, level Level, msg string, args ...any)
```

Log emits a log record with the current time and the given level and message. The Record's Attrs consist of the Logger's attributes followed by the Attrs specified by args.

The attribute arguments are processed as follows:

- If an argument is an Attr, it is used as is.
- If an argument is a string and this is not the last argument, the following argument is treated as the value and the two are combined into an Attr.
- Otherwise, the argument is treated as a value with key "!BADKEY".

#### (*Logger) LogAttrs 

``` go
func (l *Logger) LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)
```

LogAttrs is a more efficient version of [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log) that accepts only Attrs.

#### (*Logger) Warn 

``` go
func (l *Logger) Warn(msg string, args ...any)
```

Warn logs at LevelWarn.

#### (*Logger) WarnContext 

``` go
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any)
```

WarnContext logs at LevelWarn with the given context.

#### (*Logger) With 

``` go
func (l *Logger) With(args ...any) *Logger
```

With returns a Logger that includes the given attributes in each output operation. Arguments are converted to attributes as if by [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log).

#### (*Logger) WithGroup 

``` go
func (l *Logger) WithGroup(name string) *Logger
```

WithGroup returns a Logger that starts a group, if name is non-empty. The keys of all attributes added to the Logger will be qualified by the given name. (How that qualification happens depends on the [Handler.WithGroup] method of the Logger's Handler.)

If name is empty, WithGroup returns the receiver.

### type Record 

``` go
type Record struct {
	// The time at which the output method (Log, Info, etc.) was called.
	Time time.Time

	// The log message.
	Message string

	// The level of the event.
	Level Level

	// The program counter at the time the record was constructed, as determined
	// by runtime.Callers. If zero, no program counter is available.
	//
	// The only valid use for this value is as an argument to
	// [runtime.CallersFrames]. In particular, it must not be passed to
	// [runtime.FuncForPC].
	PC uintptr
	// contains filtered or unexported fields
}
```

A Record holds information about a log event. Copies of a Record share state. Do not modify a Record after handing out a copy to it. Call [NewRecord](https://pkg.go.dev/log/slog@go1.21.3#NewRecord) to create a new Record. Use [Record.Clone](https://pkg.go.dev/log/slog@go1.21.3#Record.Clone) to create a copy with no shared state.

#### func NewRecord 

``` go
func NewRecord(t time.Time, level Level, msg string, pc uintptr) Record
```

NewRecord creates a Record from the given arguments. Use [Record.AddAttrs](https://pkg.go.dev/log/slog@go1.21.3#Record.AddAttrs) to add attributes to the Record.

NewRecord is intended for logging APIs that want to support a [Handler](https://pkg.go.dev/log/slog@go1.21.3#Handler) as a backend.

#### (*Record) Add 

``` go
func (r *Record) Add(args ...any)
```

Add converts the args to Attrs as described in [Logger.Log](https://pkg.go.dev/log/slog@go1.21.3#Logger.Log), then appends the Attrs to the Record's list of Attrs. It omits empty groups.

#### (*Record) AddAttrs 

``` go
func (r *Record) AddAttrs(attrs ...Attr)
```

AddAttrs appends the given Attrs to the Record's list of Attrs. It omits empty groups.

#### (Record) Attrs 

``` go
func (r Record) Attrs(f func(Attr) bool)
```

Attrs calls f on each Attr in the Record. Iteration stops if f returns false.

#### (Record) Clone 

``` go
func (r Record) Clone() Record
```

Clone returns a copy of the record with no shared state. The original record and the clone can both be modified without interfering with each other.

#### (Record) NumAttrs 

``` go
func (r Record) NumAttrs() int
```

NumAttrs returns the number of attributes in the Record.

### type Source 

``` go
type Source struct {
	// Function is the package path-qualified function name containing the
	// source line. If non-empty, this string uniquely identifies a single
	// function in the program. This may be the empty string if not known.
	Function string `json:"function"`
	// File and Line are the file name and line number (1-based) of the source
	// line. These may be the empty string and zero, respectively, if not known.
	File string `json:"file"`
	Line int    `json:"line"`
}
```

Source describes the location of a line of source code.

### type TextHandler 

``` go
type TextHandler struct {
	// contains filtered or unexported fields
}
```

TextHandler is a Handler that writes Records to an io.Writer as a sequence of key=value pairs separated by spaces and followed by a newline.

#### func NewTextHandler 

``` go
func NewTextHandler(w io.Writer, opts *HandlerOptions) *TextHandler
```

NewTextHandler creates a TextHandler that writes to w, using the given options. If opts is nil, the default options are used.

#### (*TextHandler) Enabled 

``` go
func (h *TextHandler) Enabled(_ context.Context, level Level) bool
```

Enabled reports whether the handler handles records at the given level. The handler ignores records whose level is lower.

#### (*TextHandler) Handle 

``` go
func (h *TextHandler) Handle(_ context.Context, r Record) error
```

Handle formats its argument Record as a single line of space-separated key=value items.

If the Record's time is zero, the time is omitted. Otherwise, the key is "time" and the value is output in RFC3339 format with millisecond precision.

If the Record's level is zero, the level is omitted. Otherwise, the key is "level" and the value of [Level.String](https://pkg.go.dev/log/slog@go1.21.3#Level.String) is output.

If the AddSource option is set and source information is available, the key is "source" and the value is output as FILE:LINE.

The message's key is "msg".

To modify these or other attributes, or remove them from the output, use [HandlerOptions.ReplaceAttr].

If a value implements [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler), the result of MarshalText is written. Otherwise, the result of fmt.Sprint is written.

Keys and values are quoted with [strconv.Quote](https://pkg.go.dev/strconv#Quote) if they contain Unicode space characters, non-printing characters, '"' or '='.

Keys inside groups consist of components (keys or group names) separated by dots. No further escaping is performed. Thus there is no way to determine from the key "a.b.c" whether there are two groups "a" and "b" and a key "c", or a single group "a.b" and a key "c", or single group "a" and a key "b.c". If it is necessary to reconstruct the group structure of a key even in the presence of dots inside components, use [HandlerOptions.ReplaceAttr] to encode that information in the key.

Each call to Handle results in a single serialized call to io.Writer.Write.

#### (*TextHandler) WithAttrs 

``` go
func (h *TextHandler) WithAttrs(attrs []Attr) Handler
```

WithAttrs returns a new TextHandler whose attributes consists of h's attributes followed by attrs.

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

#### func AnyValue 

``` go
func AnyValue(v any) Value
```

AnyValue returns a Value for the supplied value.

If the supplied value is of type Value, it is returned unmodified.

Given a value of one of Go's predeclared string, bool, or (non-complex) numeric types, AnyValue returns a Value of kind String, Bool, Uint64, Int64, or Float64. The width of the original numeric type is not preserved.

Given a time.Time or time.Duration value, AnyValue returns a Value of kind KindTime or KindDuration. The monotonic time is not preserved.

For nil, or values of all other types, including named types whose underlying type is numeric, AnyValue returns a value of kind KindAny.

#### func BoolValue 

``` go
func BoolValue(v bool) Value
```

BoolValue returns a Value for a bool.

#### func DurationValue 

``` go
func DurationValue(v time.Duration) Value
```

DurationValue returns a Value for a time.Duration.

#### func Float64Value 

``` go
func Float64Value(v float64) Value
```

Float64Value returns a Value for a floating-point number.

#### func GroupValue 

``` go
func GroupValue(as ...Attr) Value
```

GroupValue returns a new Value for a list of Attrs. The caller must not subsequently mutate the argument slice.

#### func Int64Value 

``` go
func Int64Value(v int64) Value
```

Int64Value returns a Value for an int64.

#### func IntValue 

``` go
func IntValue(v int) Value
```

IntValue returns a Value for an int.

#### func StringValue 

``` go
func StringValue(value string) Value
```

StringValue returns a new Value for a string.

#### func TimeValue 

``` go
func TimeValue(v time.Time) Value
```

TimeValue returns a Value for a time.Time. It discards the monotonic portion.

#### func Uint64Value 

``` go
func Uint64Value(v uint64) Value
```

Uint64Value returns a Value for a uint64.

#### (Value) Any 

``` go
func (v Value) Any() any
```

Any returns v's value as an any.

#### (Value) Bool 

``` go
func (v Value) Bool() bool
```

Bool returns v's value as a bool. It panics if v is not a bool.

#### (Value) Duration 

``` go
func (a Value) Duration() time.Duration
```

Duration returns v's value as a time.Duration. It panics if v is not a time.Duration.

#### (Value) Equal 

``` go
func (v Value) Equal(w Value) bool
```

Equal reports whether v and w represent the same Go value.

#### (Value) Float64 

``` go
func (v Value) Float64() float64
```

Float64 returns v's value as a float64. It panics if v is not a float64.

#### (Value) Group 

``` go
func (v Value) Group() []Attr
```

Group returns v's value as a []Attr. It panics if v's Kind is not KindGroup.

#### (Value) Int64 

``` go
func (v Value) Int64() int64
```

Int64 returns v's value as an int64. It panics if v is not a signed integer.

#### (Value) Kind 

``` go
func (v Value) Kind() Kind
```

Kind returns v's Kind.

#### (Value) LogValuer 

``` go
func (v Value) LogValuer() LogValuer
```

LogValuer returns v's value as a LogValuer. It panics if v is not a LogValuer.

#### (Value) Resolve 

``` go
func (v Value) Resolve() (rv Value)
```

Resolve repeatedly calls LogValue on v while it implements LogValuer, and returns the result. If v resolves to a group, the group's attributes' values are not recursively resolved. If the number of LogValue calls exceeds a threshold, a Value containing an error is returned. Resolve's return value is guaranteed not to be of Kind KindLogValuer.

#### (Value) String 

``` go
func (v Value) String() string
```

String returns Value's value as a string, formatted like fmt.Sprint. Unlike the methods Int64, Float64, and so on, which panic if v is of the wrong kind, String never panics.

#### (Value) Time 

``` go
func (v Value) Time() time.Time
```

Time returns v's value as a time.Time. It panics if v is not a time.Time.

#### (Value) Uint64 

``` go
func (v Value) Uint64() uint64
```

Uint64 returns v's value as a uint64. It panics if v is not an unsigned integer.