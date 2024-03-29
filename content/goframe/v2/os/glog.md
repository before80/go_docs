+++
title = "glog"
date = 2024-03-21T17:56:00+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/glog](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/glog)

Package glog implements powerful and easy-to-use leveled logging functionality.

​	软件包 glog 实现了强大且易于使用的分级日志记录功能。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/glog/glog_logger.go#L50)

```go
const (
	F_ASYNC      = 1 << iota // Print logging content asynchronously。
	F_FILE_LONG              // Print full file name and line number: /a/b/c/d.go:23.
	F_FILE_SHORT             // Print final file name element and line number: d.go:23. overrides F_FILE_LONG.
	F_TIME_DATE              // Print the date in the local time zone: 2009-01-23.
	F_TIME_TIME              // Print the time in the local time zone: 01:23:23.
	F_TIME_MILLI             // Print the time with milliseconds in the local time zone: 01:23:23.675.
	F_CALLER_FN              // Print Caller function name and package: main.main
	F_TIME_STD   = F_TIME_DATE | F_TIME_MILLI
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/glog/glog_logger_color.go#L11)

```go
const (
	COLOR_BLACK = 30 + iota
	COLOR_RED
	COLOR_GREEN
	COLOR_YELLOW
	COLOR_BLUE
	COLOR_MAGENTA
	COLOR_CYAN
	COLOR_WHITE
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/glog/glog_logger_color.go#L23)

```go
const (
	COLOR_HI_BLACK = 90 + iota
	COLOR_HI_RED
	COLOR_HI_GREEN
	COLOR_HI_YELLOW
	COLOR_HI_BLUE
	COLOR_HI_MAGENTA
	COLOR_HI_CYAN
	COLOR_HI_WHITE
)
```

Foreground Hi-Intensity text colors

​	前景高强度文本颜色

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/glog/glog_logger_level.go#L18)

```go
const (
	LEVEL_ALL  = LEVEL_DEBU | LEVEL_INFO | LEVEL_NOTI | LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
	LEVEL_DEV  = LEVEL_ALL
	LEVEL_PROD = LEVEL_WARN | LEVEL_ERRO | LEVEL_CRIT
	LEVEL_NONE = 0
	LEVEL_DEBU = 1 << iota // 16
	LEVEL_INFO             // 32
	LEVEL_NOTI             // 64
	LEVEL_WARN             // 128
	LEVEL_ERRO             // 256
	LEVEL_CRIT             // 512
	LEVEL_PANI             // 1024
	LEVEL_FATA             // 2048
)
```

Note that the LEVEL_PANI and LEVEL_FATA levels are not used for logging output, but for prefix configurations.

​	请注意，LEVEL_PANI 和 LEVEL_FATA 级别不用于日志记录输出，而是用于前缀配置。

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/glog/glog_instance.go#L11)

```go
const (
	// DefaultName is the default group name for instance usage.
	DefaultName = "default"
)
```

## 变量

This section is empty.

## 函数

#### func Critical

```go
func Critical(ctx context.Context, v ...interface{})
```

Critical prints the logging content with [CRIT] header and newline. It also prints caller stack info if stack feature is enabled.

​	严重版使用 [CRIT] 标头和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### func Criticalf

```go
func Criticalf(ctx context.Context, format string, v ...interface{})
```

Criticalf prints the logging content with [CRIT] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

​	Criticalf 使用 [CRIT] 标头、自定义格式和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### func Debug

```go
func Debug(ctx context.Context, v ...interface{})
```

Debug prints the logging content with [DEBU] header and newline.

​	Debug 使用 [DEBU] 标头和换行符打印日志记录内容。

#### func Debugf

```go
func Debugf(ctx context.Context, format string, v ...interface{})
```

Debugf prints the logging content with [DEBU] header, custom format and newline.

​	Debugf 使用 [DEBU] 标头、自定义格式和换行符打印日志记录内容。

#### func Error

```go
func Error(ctx context.Context, v ...interface{})
```

Error prints the logging content with [ERRO] header and newline. It also prints caller stack info if stack feature is enabled.

​	错误打印带有 [ERRO] 标头和换行符的日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### func Errorf

```go
func Errorf(ctx context.Context, format string, v ...interface{})
```

Errorf prints the logging content with [ERRO] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

​	Errorf 使用 [ERRO] 标头、自定义格式和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### func Fatal

```go
func Fatal(ctx context.Context, v ...interface{})
```

Fatal prints the logging content with [FATA] header and newline, then exit the current process.

​	Fatal 使用 [FATA] 标头和换行符打印日志记录内容，然后退出当前进程。

#### func Fatalf

```go
func Fatalf(ctx context.Context, format string, v ...interface{})
```

Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.

​	Fatalf 使用 [FATA] 标头、自定义格式和换行符打印日志记录内容，然后退出当前进程。

#### func GetCtxKeys

```go
func GetCtxKeys() []interface{}
```

GetCtxKeys retrieves and returns the context keys for logging.

​	GetCtxKeys 检索并返回用于日志记录的上下文键。

#### func GetFlags

```go
func GetFlags() int
```

GetFlags returns the flags of defaultLogger.

​	GetFlags 返回 defaultLogger 的标志。

#### func GetLevel

```go
func GetLevel() int
```

GetLevel returns the default logging level value.

​	GetLevel 返回默认日志记录级别值。

#### func GetLevelPrefix

```go
func GetLevelPrefix(level int) string
```

GetLevelPrefix returns the prefix string for specified level.

​	GetLevelPrefix 返回指定级别的前缀字符串。

#### func GetPath

```go
func GetPath() string
```

GetPath returns the logging directory path for file logging. It returns empty string if no directory path set.

​	GetPath 返回文件日志记录的日志记录目录路径。如果未设置目录路径，则返回空字符串。

#### func GetStack

```go
func GetStack(skip ...int) string
```

GetStack returns the caller stack content, the optional parameter `skip` specify the skipped stack offset from the end point.

​	GetStack 返回调用方堆栈内容，可选参数 `skip` 指定从端点跳过的堆栈偏移量。

#### func GetWriter

```go
func GetWriter() io.Writer
```

GetWriter returns the customized writer object, which implements the io.Writer interface. It returns nil if no customized writer set.

​	GetWriter 返回自定义的 writer 对象，该对象实现 io。编写器界面。如果没有自定义编写器集，则返回 nil。

#### func HandlerJson <-2.1.0

```go
func HandlerJson(ctx context.Context, in *HandlerInput)
```

HandlerJson is a handler for output logging content as a single json string.

​	HandlerJson 是一个处理程序，用于将日志记录内容输出为单个 json 字符串。

#### func HandlerStructure <-2.5.3

```go
func HandlerStructure(ctx context.Context, in *HandlerInput)
```

HandlerStructure is a handler for output logging content as a structured string.

​	HandlerStructure 是将内容记录为结构化字符串的输出处理程序。

#### func Info

```go
func Info(ctx context.Context, v ...interface{})
```

Info prints the logging content with [INFO] header and newline.

​	Info 使用 [INFO] 标头和换行符打印日志记录内容。

#### func Infof

```go
func Infof(ctx context.Context, format string, v ...interface{})
```

Infof prints the logging content with [INFO] header, custom format and newline.

​	Infof 使用 [INFO] 标头、自定义格式和换行符打印日志记录内容。

#### func Notice

```go
func Notice(ctx context.Context, v ...interface{})
```

Notice prints the logging content with [NOTI] header and newline. It also prints caller stack info if stack feature is enabled.

​	注意使用 [NOTI] 标头和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### func Noticef

```go
func Noticef(ctx context.Context, format string, v ...interface{})
```

Noticef prints the logging content with [NOTI] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

​	Noticef 使用 [NOTI] 标头、自定义格式和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### func Panic

```go
func Panic(ctx context.Context, v ...interface{})
```

Panic prints the logging content with [PANI] header and newline, then panics.

​	Panic 使用 [PANI] 标头和换行符打印日志记录内容，然后出现 panic。

#### func Panicf

```go
func Panicf(ctx context.Context, format string, v ...interface{})
```

Panicf prints the logging content with [PANI] header, custom format and newline, then panics.

​	Panicf 使用 [PANI] 标头、自定义格式和换行符打印日志记录内容，然后 panic。

#### func Print

```go
func Print(ctx context.Context, v ...interface{})
```

Print prints `v` with newline using fmt.Sprintln. The parameter `v` can be multiple variables.

​	使用 fmt `v` 用换行符打印打印。斯普林。该参数 `v` 可以是多个变量。

#### func PrintStack

```go
func PrintStack(ctx context.Context, skip ...int)
```

PrintStack prints the caller stack, the optional parameter `skip` specify the skipped stack offset from the end point.

​	PrintStack 打印调用方堆栈，可选参数 `skip` 指定从端点跳过的堆栈偏移量。

#### func Printf

```go
func Printf(ctx context.Context, format string, v ...interface{})
```

Printf prints `v` with format `format` using fmt.Sprintf. The parameter `v` can be multiple variables.

​	Printf 使用 fmt `v` 以格式 `format` 打印。斯普林特夫。该参数 `v` 可以是多个变量。

#### func SetAsync

```go
func SetAsync(enabled bool)
```

SetAsync enables/disables async logging output feature for default defaultLogger.

​	SetAsync 为默认 defaultLogger 启用/禁用异步日志记录输出功能。

#### func SetConfig

```go
func SetConfig(config Config) error
```

SetConfig set configurations for the defaultLogger.

​	SetConfig 为 defaultLogger 设置配置。

#### func SetConfigWithMap

```go
func SetConfigWithMap(m map[string]interface{}) error
```

SetConfigWithMap set configurations with map for the defaultLogger.

​	SetConfigWithMap 使用 defaultLogger 的 map 设置配置。

#### func SetCtxKeys

```go
func SetCtxKeys(keys ...interface{})
```

SetCtxKeys sets the context keys for defaultLogger. The keys is used for retrieving values from context and printing them to logging content.

​	SetCtxKeys 设置 defaultLogger 的上下文键。这些键用于从上下文中检索值并将其打印到日志记录内容。

Note that multiple calls of this function will overwrite the previous set context keys.

​	请注意，此函数的多次调用将覆盖之前设置的上下文键。

#### func SetDebug

```go
func SetDebug(debug bool)
```

SetDebug enables/disables the debug level for default defaultLogger. The debug level is enabled in default.

​	SetDebug 启用/禁用默认 defaultLogger 的调试级别。默认情况下，调试级别处于启用状态。

#### func SetDefaultHandler <-2.1.0

```go
func SetDefaultHandler(handler Handler)
```

SetDefaultHandler sets default handler for package.

​	SetDefaultHandler 设置包的默认处理程序。

#### func SetDefaultLogger

```go
func SetDefaultLogger(l *Logger)
```

SetDefaultLogger sets the default logger for package glog. Note that there might be concurrent safety issue if calls this function in different goroutines.

​	SetDefaultLogger 设置包 glog 的默认记录器。请注意，如果在不同的 goroutine 中调用此函数，则可能存在并发安全问题。

#### func SetFile

```go
func SetFile(pattern string)
```

SetFile sets the file name `pattern` for file logging. Datetime pattern can be used in `pattern`, eg: access-{Ymd}.log. The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log

​	SetFile 设置文件日志记录的文件名 `pattern` 。日期时间模式可用于 `pattern` ，例如：access-{Ymd}.log。默认文件名模式为：Y-m-d.log，例如：2018-01-01.log

#### func SetFlags

```go
func SetFlags(flags int)
```

SetFlags sets extra flags for logging output features.

​	SetFlags 设置用于记录输出要素的额外标志。

#### func SetHandlers

```go
func SetHandlers(handlers ...Handler)
```

SetHandlers sets the logging handlers for default defaultLogger.

​	SetHandlers 设置默认 defaultLogger 的日志记录处理程序。

#### func SetHeaderPrint

```go
func SetHeaderPrint(enabled bool)
```

SetHeaderPrint sets whether output header of the logging contents, which is true in default.

​	SetHeaderPrint 设置日志记录内容是否输出标头，默认为 true。

#### func SetLevel

```go
func SetLevel(level int)
```

SetLevel sets the default logging level.

​	SetLevel 设置默认日志记录级别。

#### func SetLevelPrefix

```go
func SetLevelPrefix(level int, prefix string)
```

SetLevelPrefix sets the prefix string for specified level.

​	SetLevelPrefix 设置指定级别的前缀字符串。

#### func SetLevelPrefixes

```go
func SetLevelPrefixes(prefixes map[int]string)
```

SetLevelPrefixes sets the level to prefix string mapping for the defaultLogger.

​	SetLevelPrefixes 将级别设置为 defaultLogger 的前缀字符串映射。

#### func SetLevelStr

```go
func SetLevelStr(levelStr string) error
```

SetLevelStr sets the logging level by level string.

​	SetLevelStr 按级别字符串设置日志记录级别。

#### func SetPath

```go
func SetPath(path string) error
```

SetPath sets the directory path for file logging.

​	SetPath 设置文件日志记录的目录路径。

#### func SetPrefix

```go
func SetPrefix(prefix string)
```

SetPrefix sets prefix string for every logging content. Prefix is part of header, which means if header output is shut, no prefix will be output.

​	SetPrefix 为每个日志记录内容设置前缀字符串。前缀是标头的一部分，这意味着如果标头输出关闭，则不会输出任何前缀。

#### func SetStack

```go
func SetStack(enabled bool)
```

SetStack enables/disables the stack feature in failure logging outputs.

​	SetStack 在故障日志记录输出中启用/禁用堆栈功能。

#### func SetStdoutPrint

```go
func SetStdoutPrint(enabled bool)
```

SetStdoutPrint sets whether ouptput the logging contents to stdout, which is true in default.

​	SetStdoutPrint 设置是否将日志记录内容输出到 stdout，默认为 true。

#### func SetWriter

```go
func SetWriter(writer io.Writer)
```

SetWriter sets the customized logging `writer` for logging. The `writer` object should implements the io.Writer interface. Developer can use customized logging `writer` to redirect logging output to another service, eg: kafka, mysql, mongodb, etc.

​	SetWriter 设置用于日志记录的自定义日志记录 `writer` 。该 `writer` 对象应实现 io。编写器界面。开发人员可以使用自定义日志记录 `writer` 将日志记录输出重定向到另一个服务，例如：kafka、mysql、mongodb 等。

#### func SetWriterColorEnable

```go
func SetWriterColorEnable(enabled bool)
```

SetWriterColorEnable sets the file logging with color

​	SetWriterColorEnable 使用颜色设置文件日志记录

#### func Warning

```go
func Warning(ctx context.Context, v ...interface{})
```

Warning prints the logging content with [WARN] header and newline. It also prints caller stack info if stack feature is enabled.

​	警告打印带有 [WARN] 标头和换行符的日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### func Warningf

```go
func Warningf(ctx context.Context, format string, v ...interface{})
```

Warningf prints the logging content with [WARN] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

​	Warningf 使用 [WARN] 标头、自定义格式和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

## 类型

### type Config

```go
type Config struct {
	Handlers             []Handler      `json:"-"`                    // Logger handlers which implement feature similar as middleware.
	Writer               io.Writer      `json:"-"`                    // Customized io.Writer.
	Flags                int            `json:"flags"`                // Extra flags for logging output features.
	TimeFormat           string         `json:"timeFormat"`           // Logging time format
	Path                 string         `json:"path"`                 // Logging directory path.
	File                 string         `json:"file"`                 // Format pattern for logging file.
	Level                int            `json:"level"`                // Output level.
	Prefix               string         `json:"prefix"`               // Prefix string for every logging content.
	StSkip               int            `json:"stSkip"`               // Skipping count for stack.
	StStatus             int            `json:"stStatus"`             // Stack status(1: enabled - default; 0: disabled)
	StFilter             string         `json:"stFilter"`             // Stack string filter.
	CtxKeys              []interface{}  `json:"ctxKeys"`              // Context keys for logging, which is used for value retrieving from context.
	HeaderPrint          bool           `json:"header"`               // Print header or not(true in default).
	StdoutPrint          bool           `json:"stdout"`               // Output to stdout or not(true in default).
	LevelPrint           bool           `json:"levelPrint"`           // Print level format string or not(true in default).
	LevelPrefixes        map[int]string `json:"levelPrefixes"`        // Logging level to its prefix string mapping.
	RotateSize           int64          `json:"rotateSize"`           // Rotate the logging file if its size > 0 in bytes.
	RotateExpire         time.Duration  `json:"rotateExpire"`         // Rotate the logging file if its mtime exceeds this duration.
	RotateBackupLimit    int            `json:"rotateBackupLimit"`    // Max backup for rotated files, default is 0, means no backups.
	RotateBackupExpire   time.Duration  `json:"rotateBackupExpire"`   // Max expires for rotated files, which is 0 in default, means no expiration.
	RotateBackupCompress int            `json:"rotateBackupCompress"` // Compress level for rotated files using gzip algorithm. It's 0 in default, means no compression.
	RotateCheckInterval  time.Duration  `json:"rotateCheckInterval"`  // Asynchronously checks the backups and expiration at intervals. It's 1 hour in default.
	StdoutColorDisabled  bool           `json:"stdoutColorDisabled"`  // Logging level prefix with color to writer or not (false in default).
	WriterColorEnable    bool           `json:"writerColorEnable"`    // Logging level prefix with color to writer or not (false in default).
	// contains filtered or unexported fields
}
```

Config is the configuration object for logger.

​	Config 是 logger 的配置对象。

#### func DefaultConfig

```go
func DefaultConfig() Config
```

DefaultConfig returns the default configuration for logger.

​	DefaultConfig 返回记录器的默认配置。

### type Handler

```go
type Handler func(ctx context.Context, in *HandlerInput)
```

Handler is function handler for custom logging content outputs.

​	Handler 是用于自定义日志记录内容输出的函数处理程序。

#### func GetDefaultHandler <-2.1.0

```go
func GetDefaultHandler() Handler
```

GetDefaultHandler returns the default handler of package.

​	GetDefaultHandler 返回包的默认处理程序。

### type HandlerInput

```go
type HandlerInput struct {
	Logger      *Logger       // Current Logger object.
	Buffer      *bytes.Buffer // Buffer for logging content outputs.
	Time        time.Time     // Logging time, which is the time that logging triggers.
	TimeFormat  string        // Formatted time string, like "2016-01-09 12:00:00".
	Color       int           // Using color, like COLOR_RED, COLOR_BLUE, etc. Eg: 34
	Level       int           // Using level, like LEVEL_INFO, LEVEL_ERRO, etc. Eg: 256
	LevelFormat string        // Formatted level string, like "DEBU", "ERRO", etc. Eg: ERRO
	CallerFunc  string        // The source function name that calls logging, only available if F_CALLER_FN set.
	CallerPath  string        // The source file path and its line number that calls logging, only available if F_FILE_SHORT or F_FILE_LONG set.
	CtxStr      string        // The retrieved context value string from context, only available if Config.CtxKeys configured.
	TraceId     string        // Trace id, only available if OpenTelemetry is enabled.
	Prefix      string        // Custom prefix string for logging content.
	Content     string        // Content is the main logging content without error stack string produced by logger.
	Values      []any         // The passed un-formatted values array to logger.
	Stack       string        // Stack string produced by logger, only available if Config.StStatus configured.
	IsAsync     bool          // IsAsync marks it is in asynchronous logging.
	// contains filtered or unexported fields
}
```

HandlerInput is the input parameter struct for logging Handler.

​	HandlerInput 是用于记录 Handler 的输入参数结构。

#### (*HandlerInput) Next

```go
func (in *HandlerInput) Next(ctx context.Context)
```

Next calls the next logging handler in middleware way.

​	接下来以中间件方式调用下一个日志记录处理程序。

#### (*HandlerInput) String

```go
func (in *HandlerInput) String(withColor ...bool) string
```

String returns the logging content formatted by default logging handler.

​	String 返回由默认日志记录处理程序格式化的日志记录内容。

### type HandlerOutputJson <-2.1.0

```go
type HandlerOutputJson struct {
	Time       string `json:""`           // Formatted time string, like "2016-01-09 12:00:00".
	TraceId    string `json:",omitempty"` // Trace id, only available if tracing is enabled.
	CtxStr     string `json:",omitempty"` // The retrieved context value string from context, only available if Config.CtxKeys configured.
	Level      string `json:""`           // Formatted level string, like "DEBU", "ERRO", etc. Eg: ERRO
	CallerPath string `json:",omitempty"` // The source file path and its line number that calls logging, only available if F_FILE_SHORT or F_FILE_LONG set.
	CallerFunc string `json:",omitempty"` // The source function name that calls logging, only available if F_CALLER_FN set.
	Prefix     string `json:",omitempty"` // Custom prefix string for logging content.
	Content    string `json:""`           // Content is the main logging content, containing error stack string produced by logger.
	Stack      string `json:",omitempty"` // Stack string produced by logger, only available if Config.StStatus configured.
}
```

HandlerOutputJson is the structure outputting logging content as single json.

​	HandlerOutputJson 是将日志内容输出为单个 json 的结构。

### type ILogger <-2.1.2

```go
type ILogger interface {
	Print(ctx context.Context, v ...interface{})
	Printf(ctx context.Context, format string, v ...interface{})
	Debug(ctx context.Context, v ...interface{})
	Debugf(ctx context.Context, format string, v ...interface{})
	Info(ctx context.Context, v ...interface{})
	Infof(ctx context.Context, format string, v ...interface{})
	Notice(ctx context.Context, v ...interface{})
	Noticef(ctx context.Context, format string, v ...interface{})
	Warning(ctx context.Context, v ...interface{})
	Warningf(ctx context.Context, format string, v ...interface{})
	Error(ctx context.Context, v ...interface{})
	Errorf(ctx context.Context, format string, v ...interface{})
	Critical(ctx context.Context, v ...interface{})
	Criticalf(ctx context.Context, format string, v ...interface{})
	Panic(ctx context.Context, v ...interface{})
	Panicf(ctx context.Context, format string, v ...interface{})
	Fatal(ctx context.Context, v ...interface{})
	Fatalf(ctx context.Context, format string, v ...interface{})
}
```

ILogger is the API interface for logger.

​	ILogger 是记录器的 API 接口。

### type Logger

```go
type Logger struct {
	// contains filtered or unexported fields
}
```

Logger is the struct for logging management.

​	记录器是日志记录管理的结构。

#### func Async

```go
func Async(enabled ...bool) *Logger
```

Async is a chaining function, which enables/disables async logging output feature.

​	Async 是一种链接功能，用于启用/禁用异步日志记录输出功能。

#### func Cat

```go
func Cat(category string) *Logger
```

Cat is a chaining function, which sets the category to `category` for current logging content output.

​	Cat 是一个链接函数，它将当前日志记录内容输出的类别 `category` 设置为。

#### func DefaultLogger

```go
func DefaultLogger() *Logger
```

DefaultLogger returns the default logger.

​	DefaultLogger 返回默认记录器。

#### func Expose

```go
func Expose() *Logger
```

Expose returns the default logger of package glog.

​	Expose 返回包 glog 的默认记录器。

#### func File

```go
func File(pattern string) *Logger
```

File is a chaining function, which sets file name `pattern` for the current logging content output.

​	File 是一个链接函数，用于设置当前日志记录内容输出的文件名 `pattern` 。

#### func Header

```go
func Header(enabled ...bool) *Logger
```

Header is a chaining function, which enables/disables log header for the current logging content output. It’s enabled in default.

​	标头是一种链接功能，用于启用/禁用当前日志记录内容输出的日志标头。默认情况下，它处于启用状态。

#### func Instance

```go
func Instance(name ...string) *Logger
```

Instance returns an instance of Logger with default settings. The parameter `name` is the name for the instance.

​	Instance 返回具有默认设置的 Logger 实例。该参数 `name` 是实例的名称。

#### func Level

```go
func Level(level int) *Logger
```

Level is a chaining function, which sets logging level for the current logging content output.

​	Level 是一个链接函数，用于设置当前日志记录内容输出的日志记录级别。

#### func LevelStr

```go
func LevelStr(levelStr string) *Logger
```

LevelStr is a chaining function, which sets logging level for the current logging content output using level string.

​	LevelStr 是一个链接函数，它使用级别字符串为当前日志记录内容输出设置日志记录级别。

#### func Line

```go
func Line(long ...bool) *Logger
```

Line is a chaining function, which enables/disables printing its caller file along with its line number. The parameter `long` specified whether print the long absolute file path, eg: /a/b/c/d.go:23.

​	Line 是一个链接函数，它启用/禁用打印其调用方文件及其行号。该参数 `long` 指定是否打印长绝对文件路径，例如：/a/b/c/d.go：23。

#### func New

```go
func New() *Logger
```

New creates and returns a custom logger.

​	New 创建并返回自定义记录器。

#### func NewWithWriter

```go
func NewWithWriter(writer io.Writer) *Logger
```

NewWithWriter creates and returns a custom logger with io.Writer.

​	NewWithWriter 使用 io 创建并返回自定义记录器。作家。

#### func Path

```go
func Path(path string) *Logger
```

Path is a chaining function, which sets the directory path to `path` for current logging content output.

​	Path 是一个链接函数，它将当前日志记录内容输出的目录路径 `path` 设置为。

#### func Skip

```go
func Skip(skip int) *Logger
```

Skip is a chaining function, which sets stack skip for the current logging content output. It also affects the caller file path checks when line number printing enabled.

​	Skip 是一个链接函数，用于设置当前日志记录内容输出的堆栈跳过。它还会影响启用行号打印时的调用方文件路径检查。

#### func Stack

```go
func Stack(enabled bool, skip ...int) *Logger
```

Stack is a chaining function, which sets stack options for the current logging content output .

​	Stack 是一个链接函数，它为当前日志内容输出设置堆栈选项。

#### func StackWithFilter

```go
func StackWithFilter(filter string) *Logger
```

StackWithFilter is a chaining function, which sets stack filter for the current logging content output .

​	StackWithFilter 是一个链接函数，用于设置当前日志记录内容输出的堆栈过滤器。

#### func Stdout

```go
func Stdout(enabled ...bool) *Logger
```

Stdout is a chaining function, which enables/disables stdout for the current logging content output. It’s enabled in default.

​	Stdout 是一个链接函数，用于启用/禁用当前日志记录内容输出的 stdout。默认情况下，它处于启用状态。

#### func To

```go
func To(writer io.Writer) *Logger
```

To is a chaining function, which redirects current logging content output to the sepecified `writer`.

​	To 是一个链接函数，它将当前日志记录内容输出重定向到已 `writer` 分离的 .

#### (*Logger) AppendCtxKeys

```go
func (l *Logger) AppendCtxKeys(keys ...interface{})
```

AppendCtxKeys appends extra keys to logger. It ignores the key if it is already appended to the logger previously.

​	AppendCtxKeys 将额外的密钥追加到记录器。如果该密钥之前已附加到记录器，则该密钥将忽略该密钥。

#### (*Logger) Async

```go
func (l *Logger) Async(enabled ...bool) *Logger
```

Async is a chaining function, which enables/disables async logging output feature.

​	Async 是一种链接功能，用于启用/禁用异步日志记录输出功能。

#### (*Logger) Cat

```go
func (l *Logger) Cat(category string) *Logger
```

Cat is a chaining function, which sets the category to `category` for current logging content output. Param `category` can be hierarchical, eg: module/user.

​	Cat 是一个链接函数，它将当前日志记录内容输出的类别 `category` 设置为。参数 `category` 可以是分层的，例如：module/user。

#### (*Logger) Clone

```go
func (l *Logger) Clone() *Logger
```

Clone returns a new logger, which a `shallow copy` of the current logger. Note that the attribute `config` of the cloned one is the shallow copy of current one.

​	Clone 返回一个新的记录器，该记录器 `shallow copy` 是当前记录器的。请注意，克隆的属性 `config` 是当前克隆的浅层副本。

#### (*Logger) Critical

```go
func (l *Logger) Critical(ctx context.Context, v ...interface{})
```

Critical prints the logging content with [CRIT] header and newline. It also prints caller stack info if stack feature is enabled.

​	严重版使用 [CRIT] 标头和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### (*Logger) Criticalf

```go
func (l *Logger) Criticalf(ctx context.Context, format string, v ...interface{})
```

Criticalf prints the logging content with [CRIT] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

​	Criticalf 使用 [CRIT] 标头、自定义格式和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### (*Logger) Debug

```go
func (l *Logger) Debug(ctx context.Context, v ...interface{})
```

Debug prints the logging content with [DEBU] header and newline.

​	Debug 使用 [DEBU] 标头和换行符打印日志记录内容。

#### (*Logger) Debugf

```go
func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{})
```

Debugf prints the logging content with [DEBU] header, custom format and newline.

​	Debugf 使用 [DEBU] 标头、自定义格式和换行符打印日志记录内容。

#### (*Logger) Error

```go
func (l *Logger) Error(ctx context.Context, v ...interface{})
```

Error prints the logging content with [ERRO] header and newline. It also prints caller stack info if stack feature is enabled.

​	错误打印带有 [ERRO] 标头和换行符的日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### (*Logger) Errorf

```go
func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{})
```

Errorf prints the logging content with [ERRO] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

​	Errorf 使用 [ERRO] 标头、自定义格式和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### (*Logger) Fatal

```go
func (l *Logger) Fatal(ctx context.Context, v ...interface{})
```

Fatal prints the logging content with [FATA] header and newline, then exit the current process.

​	Fatal 使用 [FATA] 标头和换行符打印日志记录内容，然后退出当前进程。

#### (*Logger) Fatalf

```go
func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{})
```

Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.

​	Fatalf 使用 [FATA] 标头、自定义格式和换行符打印日志记录内容，然后退出当前进程。

#### (*Logger) File

```go
func (l *Logger) File(file string) *Logger
```

File is a chaining function, which sets file name `pattern` for the current logging content output.

​	File 是一个链接函数，用于设置当前日志记录内容输出的文件名 `pattern` 。

#### (*Logger) GetConfig

```go
func (l *Logger) GetConfig() Config
```

GetConfig returns the configuration of current Logger.

​	GetConfig 返回当前 Logger 的配置。

#### (*Logger) GetCtxKeys

```go
func (l *Logger) GetCtxKeys() []interface{}
```

GetCtxKeys retrieves and returns the context keys for logging.

​	GetCtxKeys 检索并返回用于日志记录的上下文键。

#### (*Logger) GetFlags

```go
func (l *Logger) GetFlags() int
```

GetFlags returns the flags of logger.

​	GetFlags 返回记录器的标志。

#### (*Logger) GetLevel

```go
func (l *Logger) GetLevel() int
```

GetLevel returns the logging level value.

​	GetLevel 返回日志记录级别值。

#### (*Logger) GetLevelPrefix

```go
func (l *Logger) GetLevelPrefix(level int) string
```

GetLevelPrefix returns the prefix string for specified level.

​	GetLevelPrefix 返回指定级别的前缀字符串。

#### (*Logger) GetPath

```go
func (l *Logger) GetPath() string
```

GetPath returns the logging directory path for file logging. It returns empty string if no directory path set.

​	GetPath 返回文件日志记录的日志记录目录路径。如果未设置目录路径，则返回空字符串。

#### (*Logger) GetStack

```go
func (l *Logger) GetStack(skip ...int) string
```

GetStack returns the caller stack content, the optional parameter `skip` specify the skipped stack offset from the end point.

​	GetStack 返回调用方堆栈内容，可选参数 `skip` 指定从端点跳过的堆栈偏移量。

#### (*Logger) GetWriter

```go
func (l *Logger) GetWriter() io.Writer
```

GetWriter returns the customized writer object, which implements the io.Writer interface. It returns nil if no writer previously set.

​	GetWriter 返回自定义的 writer 对象，该对象实现 io。编写器界面。如果之前未设置编写器，则返回 nil。

#### (*Logger) Header

```go
func (l *Logger) Header(enabled ...bool) *Logger
```

Header is a chaining function, which enables/disables log header for the current logging content output. It’s enabled in default.

​	标头是一种链接功能，用于启用/禁用当前日志记录内容输出的日志标头。默认情况下，它处于启用状态。

#### (*Logger) Info

```go
func (l *Logger) Info(ctx context.Context, v ...interface{})
```

Info prints the logging content with [INFO] header and newline.

​	Info 使用 [INFO] 标头和换行符打印日志记录内容。

#### (*Logger) Infof

```go
func (l *Logger) Infof(ctx context.Context, format string, v ...interface{})
```

Infof prints the logging content with [INFO] header, custom format and newline.

​	Infof 使用 [INFO] 标头、自定义格式和换行符打印日志记录内容。

#### (*Logger) Level

```go
func (l *Logger) Level(level int) *Logger
```

Level is a chaining function, which sets logging level for the current logging content output.

​	Level 是一个链接函数，用于设置当前日志记录内容输出的日志记录级别。

#### (*Logger) LevelStr

```go
func (l *Logger) LevelStr(levelStr string) *Logger
```

LevelStr is a chaining function, which sets logging level for the current logging content output using level string.

​	LevelStr 是一个链接函数，它使用级别字符串为当前日志记录内容输出设置日志记录级别。

#### (*Logger) Line

```go
func (l *Logger) Line(long ...bool) *Logger
```

Line is a chaining function, which enables/disables printing its caller file path along with its line number. The parameter `long` specified whether print the long absolute file path, eg: /a/b/c/d.go:23, or else short one: d.go:23.

​	Line 是一个链接函数，它启用/禁用打印其调用方文件路径及其行号。该参数 `long` 指定是打印长绝对文件路径，例如：/a/b/c/d.go：23，还是短路径：d.go：23。

#### (*Logger) Notice

```go
func (l *Logger) Notice(ctx context.Context, v ...interface{})
```

Notice prints the logging content with [NOTI] header and newline. It also prints caller stack info if stack feature is enabled.

​	注意使用 [NOTI] 标头和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### (*Logger) Noticef

```go
func (l *Logger) Noticef(ctx context.Context, format string, v ...interface{})
```

Noticef prints the logging content with [NOTI] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

​	Noticef 使用 [NOTI] 标头、自定义格式和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### (*Logger) Panic

```go
func (l *Logger) Panic(ctx context.Context, v ...interface{})
```

Panic prints the logging content with [PANI] header and newline, then panics.

​	Panic 使用 [PANI] 标头和换行符打印日志记录内容，然后出现 panic。

#### (*Logger) Panicf

```go
func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{})
```

Panicf prints the logging content with [PANI] header, custom format and newline, then panics.

​	Panicf 使用 [PANI] 标头、自定义格式和换行符打印日志记录内容，然后 panic。

#### (*Logger) Path

```go
func (l *Logger) Path(path string) *Logger
```

Path is a chaining function, which sets the directory path to `path` for current logging content output.

​	Path 是一个链接函数，它将当前日志记录内容输出的目录路径 `path` 设置为。

Note that the parameter `path` is a directory path, not a file path.

​	请注意，该参数 `path` 是目录路径，而不是文件路径。

#### (*Logger) Print

```go
func (l *Logger) Print(ctx context.Context, v ...interface{})
```

Print prints `v` with newline using fmt.Sprintln. The parameter `v` can be multiple variables.

​	使用 fmt `v` 用换行符打印打印。斯普林。该参数 `v` 可以是多个变量。

#### (*Logger) PrintStack

```go
func (l *Logger) PrintStack(ctx context.Context, skip ...int)
```

PrintStack prints the caller stack, the optional parameter `skip` specify the skipped stack offset from the end point.

​	PrintStack 打印调用方堆栈，可选参数 `skip` 指定从端点跳过的堆栈偏移量。

#### (*Logger) Printf

```go
func (l *Logger) Printf(ctx context.Context, format string, v ...interface{})
```

Printf prints `v` with format `format` using fmt.Sprintf. The parameter `v` can be multiple variables.

​	Printf 使用 fmt `v` 以格式 `format` 打印。斯普林特夫。该参数 `v` 可以是多个变量。

#### (*Logger) SetAsync

```go
func (l *Logger) SetAsync(enabled bool)
```

SetAsync enables/disables async logging output feature.

​	SetAsync 启用/禁用异步日志记录输出功能。

#### (*Logger) SetConfig

```go
func (l *Logger) SetConfig(config Config) error
```

SetConfig set configurations for the logger.

​	SetConfig 设置记录器的配置。

#### (*Logger) SetConfigWithMap

```go
func (l *Logger) SetConfigWithMap(m map[string]interface{}) error
```

SetConfigWithMap set configurations with map for the logger.

​	SetConfigWithMap 使用记录器的 map 设置配置。

#### (*Logger) SetCtxKeys

```go
func (l *Logger) SetCtxKeys(keys ...interface{})
```

SetCtxKeys sets the context keys for logger. The keys is used for retrieving values from context and printing them to logging content.

​	SetCtxKeys 设置记录器的上下文键。这些键用于从上下文中检索值并将其打印到日志记录内容。

Note that multiple calls of this function will overwrite the previous set context keys.

​	请注意，此函数的多次调用将覆盖之前设置的上下文键。

#### (*Logger) SetDebug

```go
func (l *Logger) SetDebug(debug bool)
```

SetDebug enables/disables the debug level for logger. The debug level is enabled in default.

​	SetDebug 启用/禁用记录器的调试级别。默认情况下，调试级别处于启用状态。

#### (*Logger) SetFile

```go
func (l *Logger) SetFile(pattern string)
```

SetFile sets the file name `pattern` for file logging. Datetime pattern can be used in `pattern`, eg: access-{Ymd}.log. The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log

​	SetFile 设置文件日志记录的文件名 `pattern` 。日期时间模式可用于 `pattern` ，例如：access-{Ymd}.log。默认文件名模式为：Y-m-d.log，例如：2018-01-01.log

#### (*Logger) SetFlags

```go
func (l *Logger) SetFlags(flags int)
```

SetFlags sets extra flags for logging output features.

​	SetFlags 设置用于记录输出要素的额外标志。

#### (*Logger) SetHandlers

```go
func (l *Logger) SetHandlers(handlers ...Handler)
```

SetHandlers sets the logging handlers for current logger.

​	SetHandlers 设置当前记录器的日志记录处理程序。

#### (*Logger) SetHeaderPrint

```go
func (l *Logger) SetHeaderPrint(enabled bool)
```

SetHeaderPrint sets whether output header of the logging contents, which is true in default.

​	SetHeaderPrint 设置日志记录内容是否输出标头，默认为 true。

#### (*Logger) SetLevel

```go
func (l *Logger) SetLevel(level int)
```

SetLevel sets the logging level. Note that levels `LEVEL_CRIT | LEVEL_PANI | LEVEL_FATA` cannot be removed for logging content, which are automatically added to levels.

​	SetLevel 设置日志记录级别。请注意，无法删除记录内容的级别 `LEVEL_CRIT | LEVEL_PANI | LEVEL_FATA` ，这些内容会自动添加到级别中。

#### (*Logger) SetLevelPrefix

```go
func (l *Logger) SetLevelPrefix(level int, prefix string)
```

SetLevelPrefix sets the prefix string for specified level.

​	SetLevelPrefix 设置指定级别的前缀字符串。

#### (*Logger) SetLevelPrefixes

```go
func (l *Logger) SetLevelPrefixes(prefixes map[int]string)
```

SetLevelPrefixes sets the level to prefix string mapping for the logger.

​	SetLevelPrefixes 将级别设置为记录器的前缀字符串映射。

#### (*Logger) SetLevelPrint

```go
func (l *Logger) SetLevelPrint(enabled bool)
```

SetLevelPrint sets whether output level string of the logging contents, which is true in default.

​	SetLevelPrint 设置日志记录内容是否输出级别字符串，默认为 true。

#### (*Logger) SetLevelStr

```go
func (l *Logger) SetLevelStr(levelStr string) error
```

SetLevelStr sets the logging level by level string.

​	SetLevelStr 按级别字符串设置日志记录级别。

#### (*Logger) SetPath

```go
func (l *Logger) SetPath(path string) error
```

SetPath sets the directory path for file logging.

​	SetPath 设置文件日志记录的目录路径。

#### (*Logger) SetPrefix

```go
func (l *Logger) SetPrefix(prefix string)
```

SetPrefix sets prefix string for every logging content. Prefix is part of header, which means if header output is shut, no prefix will be output.

​	SetPrefix 为每个日志记录内容设置前缀字符串。前缀是标头的一部分，这意味着如果标头输出关闭，则不会输出任何前缀。

#### (*Logger) SetStack

```go
func (l *Logger) SetStack(enabled bool)
```

SetStack enables/disables the stack feature in failure logging outputs.

​	SetStack 在故障日志记录输出中启用/禁用堆栈功能。

#### (*Logger) SetStackFilter

```go
func (l *Logger) SetStackFilter(filter string)
```

SetStackFilter sets the stack filter from the end point.

​	SetStackFilter 从端点设置堆栈筛选器。

#### (*Logger) SetStackSkip

```go
func (l *Logger) SetStackSkip(skip int)
```

SetStackSkip sets the stack offset from the end point.

​	SetStackSkip 设置堆栈与端点的偏移量。

#### (*Logger) SetStdoutColorDisabled

```go
func (l *Logger) SetStdoutColorDisabled(disabled bool)
```

SetStdoutColorDisabled disables stdout logging with color.

​	SetStdoutColorDisabled 禁用带有颜色的 stdout 日志记录。

#### (*Logger) SetStdoutPrint

```go
func (l *Logger) SetStdoutPrint(enabled bool)
```

SetStdoutPrint sets whether output the logging contents to stdout, which is true in default.

​	SetStdoutPrint 设置是否将日志记录内容输出到 stdout，默认为 true。

#### (*Logger) SetTimeFormat

```go
func (l *Logger) SetTimeFormat(timeFormat string)
```

SetTimeFormat sets the time format for the logging time.

​	SetTimeFormat 设置日志记录时间的时间格式。

#### (*Logger) SetWriter

```go
func (l *Logger) SetWriter(writer io.Writer)
```

SetWriter sets the customized logging `writer` for logging. The `writer` object should implement the io.Writer interface. Developer can use customized logging `writer` to redirect logging output to another service, eg: kafka, mysql, mongodb, etc.

​	SetWriter 设置用于日志记录的自定义日志记录 `writer` 。该 `writer` 对象应实现 io。编写器界面。开发人员可以使用自定义日志记录 `writer` 将日志记录输出重定向到另一个服务，例如：kafka、mysql、mongodb 等。

#### (*Logger) SetWriterColorEnable

```go
func (l *Logger) SetWriterColorEnable(enabled bool)
```

SetWriterColorEnable enables file/writer logging with color.

​	SetWriterColorEnable 启用带有颜色的文件/写入器日志记录。

#### (*Logger) Skip

```go
func (l *Logger) Skip(skip int) *Logger
```

Skip is a chaining function, which sets stack skip for the current logging content output. It also affects the caller file path checks when line number printing enabled.

​	Skip 是一个链接函数，用于设置当前日志记录内容输出的堆栈跳过。它还会影响启用行号打印时的调用方文件路径检查。

#### (*Logger) Stack

```go
func (l *Logger) Stack(enabled bool, skip ...int) *Logger
```

Stack is a chaining function, which sets stack options for the current logging content output .

​	Stack 是一个链接函数，它为当前日志内容输出设置堆栈选项。

#### (*Logger) StackWithFilter

```go
func (l *Logger) StackWithFilter(filter string) *Logger
```

StackWithFilter is a chaining function, which sets stack filter for the current logging content output .

​	StackWithFilter 是一个链接函数，用于设置当前日志记录内容输出的堆栈过滤器。

#### (*Logger) Stdout

```go
func (l *Logger) Stdout(enabled ...bool) *Logger
```

Stdout is a chaining function, which enables/disables stdout for the current logging content output. It’s enabled in default.

​	Stdout 是一个链接函数，用于启用/禁用当前日志记录内容输出的 stdout。默认情况下，它处于启用状态。

#### (*Logger) To

```go
func (l *Logger) To(writer io.Writer) *Logger
```

To is a chaining function, which redirects current logging content output to the specified `writer`.

​	To 是一个链接函数，它将当前日志记录内容输出重定向到指定的 `writer` 。

#### (*Logger) Warning

```go
func (l *Logger) Warning(ctx context.Context, v ...interface{})
```

Warning prints the logging content with [WARN] header and newline. It also prints caller stack info if stack feature is enabled.

​	警告打印带有 [WARN] 标头和换行符的日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### (*Logger) Warningf

```go
func (l *Logger) Warningf(ctx context.Context, format string, v ...interface{})
```

Warningf prints the logging content with [WARN] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

​	Warningf 使用 [WARN] 标头、自定义格式和换行符打印日志记录内容。如果启用了堆栈功能，它还会打印调用方堆栈信息。

#### (*Logger) Write

```go
func (l *Logger) Write(p []byte) (n int, err error)
```

Write implements the io.Writer interface. It just prints the content using Print.

​	Write 实现 io。编写器界面。它只是使用打印打印内容。