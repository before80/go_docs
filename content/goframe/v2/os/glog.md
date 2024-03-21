+++
title = "glog"
date = 2024-03-21T17:56:00+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/glog

Package glog implements powerful and easy-to-use leveled logging functionality.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/glog/glog_logger.go#L50)

``` go
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

``` go
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

``` go
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

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/glog/glog_logger_level.go#L18)

``` go
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

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/glog/glog_instance.go#L11)

``` go
const (
	// DefaultName is the default group name for instance usage.
	DefaultName = "default"
)
```

### Variables 

This section is empty.

### Functions 

##### func Critical 

``` go
func Critical(ctx context.Context, v ...interface{})
```

Critical prints the logging content with [CRIT] header and newline. It also prints caller stack info if stack feature is enabled.

##### func Criticalf 

``` go
func Criticalf(ctx context.Context, format string, v ...interface{})
```

Criticalf prints the logging content with [CRIT] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

##### func Debug 

``` go
func Debug(ctx context.Context, v ...interface{})
```

Debug prints the logging content with [DEBU] header and newline.

##### func Debugf 

``` go
func Debugf(ctx context.Context, format string, v ...interface{})
```

Debugf prints the logging content with [DEBU] header, custom format and newline.

##### func Error 

``` go
func Error(ctx context.Context, v ...interface{})
```

Error prints the logging content with [ERRO] header and newline. It also prints caller stack info if stack feature is enabled.

##### func Errorf 

``` go
func Errorf(ctx context.Context, format string, v ...interface{})
```

Errorf prints the logging content with [ERRO] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

##### func Fatal 

``` go
func Fatal(ctx context.Context, v ...interface{})
```

Fatal prints the logging content with [FATA] header and newline, then exit the current process.

##### func Fatalf 

``` go
func Fatalf(ctx context.Context, format string, v ...interface{})
```

Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.

##### func GetCtxKeys 

``` go
func GetCtxKeys() []interface{}
```

GetCtxKeys retrieves and returns the context keys for logging.

##### func GetFlags 

``` go
func GetFlags() int
```

GetFlags returns the flags of defaultLogger.

##### func GetLevel 

``` go
func GetLevel() int
```

GetLevel returns the default logging level value.

##### func GetLevelPrefix 

``` go
func GetLevelPrefix(level int) string
```

GetLevelPrefix returns the prefix string for specified level.

##### func GetPath 

``` go
func GetPath() string
```

GetPath returns the logging directory path for file logging. It returns empty string if no directory path set.

##### func GetStack 

``` go
func GetStack(skip ...int) string
```

GetStack returns the caller stack content, the optional parameter `skip` specify the skipped stack offset from the end point.

##### func GetWriter 

``` go
func GetWriter() io.Writer
```

GetWriter returns the customized writer object, which implements the io.Writer interface. It returns nil if no customized writer set.

##### func HandlerJson <-2.1.0

``` go
func HandlerJson(ctx context.Context, in *HandlerInput)
```

HandlerJson is a handler for output logging content as a single json string.

##### func HandlerStructure <-2.5.3

``` go
func HandlerStructure(ctx context.Context, in *HandlerInput)
```

HandlerStructure is a handler for output logging content as a structured string.

##### func Info 

``` go
func Info(ctx context.Context, v ...interface{})
```

Info prints the logging content with [INFO] header and newline.

##### func Infof 

``` go
func Infof(ctx context.Context, format string, v ...interface{})
```

Infof prints the logging content with [INFO] header, custom format and newline.

##### func Notice 

``` go
func Notice(ctx context.Context, v ...interface{})
```

Notice prints the logging content with [NOTI] header and newline. It also prints caller stack info if stack feature is enabled.

##### func Noticef 

``` go
func Noticef(ctx context.Context, format string, v ...interface{})
```

Noticef prints the logging content with [NOTI] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

##### func Panic 

``` go
func Panic(ctx context.Context, v ...interface{})
```

Panic prints the logging content with [PANI] header and newline, then panics.

##### func Panicf 

``` go
func Panicf(ctx context.Context, format string, v ...interface{})
```

Panicf prints the logging content with [PANI] header, custom format and newline, then panics.

##### func Print 

``` go
func Print(ctx context.Context, v ...interface{})
```

Print prints `v` with newline using fmt.Sprintln. The parameter `v` can be multiple variables.

##### func PrintStack 

``` go
func PrintStack(ctx context.Context, skip ...int)
```

PrintStack prints the caller stack, the optional parameter `skip` specify the skipped stack offset from the end point.

##### func Printf 

``` go
func Printf(ctx context.Context, format string, v ...interface{})
```

Printf prints `v` with format `format` using fmt.Sprintf. The parameter `v` can be multiple variables.

##### func SetAsync 

``` go
func SetAsync(enabled bool)
```

SetAsync enables/disables async logging output feature for default defaultLogger.

##### func SetConfig 

``` go
func SetConfig(config Config) error
```

SetConfig set configurations for the defaultLogger.

##### func SetConfigWithMap 

``` go
func SetConfigWithMap(m map[string]interface{}) error
```

SetConfigWithMap set configurations with map for the defaultLogger.

##### func SetCtxKeys 

``` go
func SetCtxKeys(keys ...interface{})
```

SetCtxKeys sets the context keys for defaultLogger. The keys is used for retrieving values from context and printing them to logging content.

Note that multiple calls of this function will overwrite the previous set context keys.

##### func SetDebug 

``` go
func SetDebug(debug bool)
```

SetDebug enables/disables the debug level for default defaultLogger. The debug level is enabled in default.

##### func SetDefaultHandler <-2.1.0

``` go
func SetDefaultHandler(handler Handler)
```

SetDefaultHandler sets default handler for package.

##### func SetDefaultLogger 

``` go
func SetDefaultLogger(l *Logger)
```

SetDefaultLogger sets the default logger for package glog. Note that there might be concurrent safety issue if calls this function in different goroutines.

##### func SetFile 

``` go
func SetFile(pattern string)
```

SetFile sets the file name `pattern` for file logging. Datetime pattern can be used in `pattern`, eg: access-{Ymd}.log. The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log

##### func SetFlags 

``` go
func SetFlags(flags int)
```

SetFlags sets extra flags for logging output features.

##### func SetHandlers 

``` go
func SetHandlers(handlers ...Handler)
```

SetHandlers sets the logging handlers for default defaultLogger.

##### func SetHeaderPrint 

``` go
func SetHeaderPrint(enabled bool)
```

SetHeaderPrint sets whether output header of the logging contents, which is true in default.

##### func SetLevel 

``` go
func SetLevel(level int)
```

SetLevel sets the default logging level.

##### func SetLevelPrefix 

``` go
func SetLevelPrefix(level int, prefix string)
```

SetLevelPrefix sets the prefix string for specified level.

##### func SetLevelPrefixes 

``` go
func SetLevelPrefixes(prefixes map[int]string)
```

SetLevelPrefixes sets the level to prefix string mapping for the defaultLogger.

##### func SetLevelStr 

``` go
func SetLevelStr(levelStr string) error
```

SetLevelStr sets the logging level by level string.

##### func SetPath 

``` go
func SetPath(path string) error
```

SetPath sets the directory path for file logging.

##### func SetPrefix 

``` go
func SetPrefix(prefix string)
```

SetPrefix sets prefix string for every logging content. Prefix is part of header, which means if header output is shut, no prefix will be output.

##### func SetStack 

``` go
func SetStack(enabled bool)
```

SetStack enables/disables the stack feature in failure logging outputs.

##### func SetStdoutPrint 

``` go
func SetStdoutPrint(enabled bool)
```

SetStdoutPrint sets whether ouptput the logging contents to stdout, which is true in default.

##### func SetWriter 

``` go
func SetWriter(writer io.Writer)
```

SetWriter sets the customized logging `writer` for logging. The `writer` object should implements the io.Writer interface. Developer can use customized logging `writer` to redirect logging output to another service, eg: kafka, mysql, mongodb, etc.

##### func SetWriterColorEnable 

``` go
func SetWriterColorEnable(enabled bool)
```

SetWriterColorEnable sets the file logging with color

##### func Warning 

``` go
func Warning(ctx context.Context, v ...interface{})
```

Warning prints the logging content with [WARN] header and newline. It also prints caller stack info if stack feature is enabled.

##### func Warningf 

``` go
func Warningf(ctx context.Context, format string, v ...interface{})
```

Warningf prints the logging content with [WARN] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

### Types 

#### type Config 

``` go
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

##### func DefaultConfig 

``` go
func DefaultConfig() Config
```

DefaultConfig returns the default configuration for logger.

#### type Handler 

``` go
type Handler func(ctx context.Context, in *HandlerInput)
```

Handler is function handler for custom logging content outputs.

##### func GetDefaultHandler <-2.1.0

``` go
func GetDefaultHandler() Handler
```

GetDefaultHandler returns the default handler of package.

#### type HandlerInput 

``` go
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

##### (*HandlerInput) Next 

``` go
func (in *HandlerInput) Next(ctx context.Context)
```

Next calls the next logging handler in middleware way.

##### (*HandlerInput) String 

``` go
func (in *HandlerInput) String(withColor ...bool) string
```

String returns the logging content formatted by default logging handler.

#### type HandlerOutputJson <-2.1.0

``` go
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

#### type ILogger <-2.1.2

``` go
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

#### type Logger 

``` go
type Logger struct {
	// contains filtered or unexported fields
}
```

Logger is the struct for logging management.

##### func Async 

``` go
func Async(enabled ...bool) *Logger
```

Async is a chaining function, which enables/disables async logging output feature.

##### func Cat 

``` go
func Cat(category string) *Logger
```

Cat is a chaining function, which sets the category to `category` for current logging content output.

##### func DefaultLogger 

``` go
func DefaultLogger() *Logger
```

DefaultLogger returns the default logger.

##### func Expose 

``` go
func Expose() *Logger
```

Expose returns the default logger of package glog.

##### func File 

``` go
func File(pattern string) *Logger
```

File is a chaining function, which sets file name `pattern` for the current logging content output.

##### func Header 

``` go
func Header(enabled ...bool) *Logger
```

Header is a chaining function, which enables/disables log header for the current logging content output. It's enabled in default.

##### func Instance 

``` go
func Instance(name ...string) *Logger
```

Instance returns an instance of Logger with default settings. The parameter `name` is the name for the instance.

##### func Level 

``` go
func Level(level int) *Logger
```

Level is a chaining function, which sets logging level for the current logging content output.

##### func LevelStr 

``` go
func LevelStr(levelStr string) *Logger
```

LevelStr is a chaining function, which sets logging level for the current logging content output using level string.

##### func Line 

``` go
func Line(long ...bool) *Logger
```

Line is a chaining function, which enables/disables printing its caller file along with its line number. The parameter `long` specified whether print the long absolute file path, eg: /a/b/c/d.go:23.

##### func New 

``` go
func New() *Logger
```

New creates and returns a custom logger.

##### func NewWithWriter 

``` go
func NewWithWriter(writer io.Writer) *Logger
```

NewWithWriter creates and returns a custom logger with io.Writer.

##### func Path 

``` go
func Path(path string) *Logger
```

Path is a chaining function, which sets the directory path to `path` for current logging content output.

##### func Skip 

``` go
func Skip(skip int) *Logger
```

Skip is a chaining function, which sets stack skip for the current logging content output. It also affects the caller file path checks when line number printing enabled.

##### func Stack 

``` go
func Stack(enabled bool, skip ...int) *Logger
```

Stack is a chaining function, which sets stack options for the current logging content output .

##### func StackWithFilter 

``` go
func StackWithFilter(filter string) *Logger
```

StackWithFilter is a chaining function, which sets stack filter for the current logging content output .

##### func Stdout 

``` go
func Stdout(enabled ...bool) *Logger
```

Stdout is a chaining function, which enables/disables stdout for the current logging content output. It's enabled in default.

##### func To 

``` go
func To(writer io.Writer) *Logger
```

To is a chaining function, which redirects current logging content output to the sepecified `writer`.

##### (*Logger) AppendCtxKeys 

``` go
func (l *Logger) AppendCtxKeys(keys ...interface{})
```

AppendCtxKeys appends extra keys to logger. It ignores the key if it is already appended to the logger previously.

##### (*Logger) Async 

``` go
func (l *Logger) Async(enabled ...bool) *Logger
```

Async is a chaining function, which enables/disables async logging output feature.

##### (*Logger) Cat 

``` go
func (l *Logger) Cat(category string) *Logger
```

Cat is a chaining function, which sets the category to `category` for current logging content output. Param `category` can be hierarchical, eg: module/user.

##### (*Logger) Clone 

``` go
func (l *Logger) Clone() *Logger
```

Clone returns a new logger, which a `shallow copy` of the current logger. Note that the attribute `config` of the cloned one is the shallow copy of current one.

##### (*Logger) Critical 

``` go
func (l *Logger) Critical(ctx context.Context, v ...interface{})
```

Critical prints the logging content with [CRIT] header and newline. It also prints caller stack info if stack feature is enabled.

##### (*Logger) Criticalf 

``` go
func (l *Logger) Criticalf(ctx context.Context, format string, v ...interface{})
```

Criticalf prints the logging content with [CRIT] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

##### (*Logger) Debug 

``` go
func (l *Logger) Debug(ctx context.Context, v ...interface{})
```

Debug prints the logging content with [DEBU] header and newline.

##### (*Logger) Debugf 

``` go
func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{})
```

Debugf prints the logging content with [DEBU] header, custom format and newline.

##### (*Logger) Error 

``` go
func (l *Logger) Error(ctx context.Context, v ...interface{})
```

Error prints the logging content with [ERRO] header and newline. It also prints caller stack info if stack feature is enabled.

##### (*Logger) Errorf 

``` go
func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{})
```

Errorf prints the logging content with [ERRO] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

##### (*Logger) Fatal 

``` go
func (l *Logger) Fatal(ctx context.Context, v ...interface{})
```

Fatal prints the logging content with [FATA] header and newline, then exit the current process.

##### (*Logger) Fatalf 

``` go
func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{})
```

Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.

##### (*Logger) File 

``` go
func (l *Logger) File(file string) *Logger
```

File is a chaining function, which sets file name `pattern` for the current logging content output.

##### (*Logger) GetConfig 

``` go
func (l *Logger) GetConfig() Config
```

GetConfig returns the configuration of current Logger.

##### (*Logger) GetCtxKeys 

``` go
func (l *Logger) GetCtxKeys() []interface{}
```

GetCtxKeys retrieves and returns the context keys for logging.

##### (*Logger) GetFlags 

``` go
func (l *Logger) GetFlags() int
```

GetFlags returns the flags of logger.

##### (*Logger) GetLevel 

``` go
func (l *Logger) GetLevel() int
```

GetLevel returns the logging level value.

##### (*Logger) GetLevelPrefix 

``` go
func (l *Logger) GetLevelPrefix(level int) string
```

GetLevelPrefix returns the prefix string for specified level.

##### (*Logger) GetPath 

``` go
func (l *Logger) GetPath() string
```

GetPath returns the logging directory path for file logging. It returns empty string if no directory path set.

##### (*Logger) GetStack 

``` go
func (l *Logger) GetStack(skip ...int) string
```

GetStack returns the caller stack content, the optional parameter `skip` specify the skipped stack offset from the end point.

##### (*Logger) GetWriter 

``` go
func (l *Logger) GetWriter() io.Writer
```

GetWriter returns the customized writer object, which implements the io.Writer interface. It returns nil if no writer previously set.

##### (*Logger) Header 

``` go
func (l *Logger) Header(enabled ...bool) *Logger
```

Header is a chaining function, which enables/disables log header for the current logging content output. It's enabled in default.

##### (*Logger) Info 

``` go
func (l *Logger) Info(ctx context.Context, v ...interface{})
```

Info prints the logging content with [INFO] header and newline.

##### (*Logger) Infof 

``` go
func (l *Logger) Infof(ctx context.Context, format string, v ...interface{})
```

Infof prints the logging content with [INFO] header, custom format and newline.

##### (*Logger) Level 

``` go
func (l *Logger) Level(level int) *Logger
```

Level is a chaining function, which sets logging level for the current logging content output.

##### (*Logger) LevelStr 

``` go
func (l *Logger) LevelStr(levelStr string) *Logger
```

LevelStr is a chaining function, which sets logging level for the current logging content output using level string.

##### (*Logger) Line 

``` go
func (l *Logger) Line(long ...bool) *Logger
```

Line is a chaining function, which enables/disables printing its caller file path along with its line number. The parameter `long` specified whether print the long absolute file path, eg: /a/b/c/d.go:23, or else short one: d.go:23.

##### (*Logger) Notice 

``` go
func (l *Logger) Notice(ctx context.Context, v ...interface{})
```

Notice prints the logging content with [NOTI] header and newline. It also prints caller stack info if stack feature is enabled.

##### (*Logger) Noticef 

``` go
func (l *Logger) Noticef(ctx context.Context, format string, v ...interface{})
```

Noticef prints the logging content with [NOTI] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

##### (*Logger) Panic 

``` go
func (l *Logger) Panic(ctx context.Context, v ...interface{})
```

Panic prints the logging content with [PANI] header and newline, then panics.

##### (*Logger) Panicf 

``` go
func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{})
```

Panicf prints the logging content with [PANI] header, custom format and newline, then panics.

##### (*Logger) Path 

``` go
func (l *Logger) Path(path string) *Logger
```

Path is a chaining function, which sets the directory path to `path` for current logging content output.

Note that the parameter `path` is a directory path, not a file path.

##### (*Logger) Print 

``` go
func (l *Logger) Print(ctx context.Context, v ...interface{})
```

Print prints `v` with newline using fmt.Sprintln. The parameter `v` can be multiple variables.

##### (*Logger) PrintStack 

``` go
func (l *Logger) PrintStack(ctx context.Context, skip ...int)
```

PrintStack prints the caller stack, the optional parameter `skip` specify the skipped stack offset from the end point.

##### (*Logger) Printf 

``` go
func (l *Logger) Printf(ctx context.Context, format string, v ...interface{})
```

Printf prints `v` with format `format` using fmt.Sprintf. The parameter `v` can be multiple variables.

##### (*Logger) SetAsync 

``` go
func (l *Logger) SetAsync(enabled bool)
```

SetAsync enables/disables async logging output feature.

##### (*Logger) SetConfig 

``` go
func (l *Logger) SetConfig(config Config) error
```

SetConfig set configurations for the logger.

##### (*Logger) SetConfigWithMap 

``` go
func (l *Logger) SetConfigWithMap(m map[string]interface{}) error
```

SetConfigWithMap set configurations with map for the logger.

##### (*Logger) SetCtxKeys 

``` go
func (l *Logger) SetCtxKeys(keys ...interface{})
```

SetCtxKeys sets the context keys for logger. The keys is used for retrieving values from context and printing them to logging content.

Note that multiple calls of this function will overwrite the previous set context keys.

##### (*Logger) SetDebug 

``` go
func (l *Logger) SetDebug(debug bool)
```

SetDebug enables/disables the debug level for logger. The debug level is enabled in default.

##### (*Logger) SetFile 

``` go
func (l *Logger) SetFile(pattern string)
```

SetFile sets the file name `pattern` for file logging. Datetime pattern can be used in `pattern`, eg: access-{Ymd}.log. The default file name pattern is: Y-m-d.log, eg: 2018-01-01.log

##### (*Logger) SetFlags 

``` go
func (l *Logger) SetFlags(flags int)
```

SetFlags sets extra flags for logging output features.

##### (*Logger) SetHandlers 

``` go
func (l *Logger) SetHandlers(handlers ...Handler)
```

SetHandlers sets the logging handlers for current logger.

##### (*Logger) SetHeaderPrint 

``` go
func (l *Logger) SetHeaderPrint(enabled bool)
```

SetHeaderPrint sets whether output header of the logging contents, which is true in default.

##### (*Logger) SetLevel 

``` go
func (l *Logger) SetLevel(level int)
```

SetLevel sets the logging level. Note that levels ` LEVEL_CRIT | LEVEL_PANI | LEVEL_FATA ` cannot be removed for logging content, which are automatically added to levels.

##### (*Logger) SetLevelPrefix 

``` go
func (l *Logger) SetLevelPrefix(level int, prefix string)
```

SetLevelPrefix sets the prefix string for specified level.

##### (*Logger) SetLevelPrefixes 

``` go
func (l *Logger) SetLevelPrefixes(prefixes map[int]string)
```

SetLevelPrefixes sets the level to prefix string mapping for the logger.

##### (*Logger) SetLevelPrint <-2.3.0

``` go
func (l *Logger) SetLevelPrint(enabled bool)
```

SetLevelPrint sets whether output level string of the logging contents, which is true in default.

##### (*Logger) SetLevelStr 

``` go
func (l *Logger) SetLevelStr(levelStr string) error
```

SetLevelStr sets the logging level by level string.

##### (*Logger) SetPath 

``` go
func (l *Logger) SetPath(path string) error
```

SetPath sets the directory path for file logging.

##### (*Logger) SetPrefix 

``` go
func (l *Logger) SetPrefix(prefix string)
```

SetPrefix sets prefix string for every logging content. Prefix is part of header, which means if header output is shut, no prefix will be output.

##### (*Logger) SetStack 

``` go
func (l *Logger) SetStack(enabled bool)
```

SetStack enables/disables the stack feature in failure logging outputs.

##### (*Logger) SetStackFilter 

``` go
func (l *Logger) SetStackFilter(filter string)
```

SetStackFilter sets the stack filter from the end point.

##### (*Logger) SetStackSkip 

``` go
func (l *Logger) SetStackSkip(skip int)
```

SetStackSkip sets the stack offset from the end point.

##### (*Logger) SetStdoutColorDisabled 

``` go
func (l *Logger) SetStdoutColorDisabled(disabled bool)
```

SetStdoutColorDisabled disables stdout logging with color.

##### (*Logger) SetStdoutPrint 

``` go
func (l *Logger) SetStdoutPrint(enabled bool)
```

SetStdoutPrint sets whether output the logging contents to stdout, which is true in default.

##### (*Logger) SetTimeFormat <-2.4.2

``` go
func (l *Logger) SetTimeFormat(timeFormat string)
```

SetTimeFormat sets the time format for the logging time.

##### (*Logger) SetWriter 

``` go
func (l *Logger) SetWriter(writer io.Writer)
```

SetWriter sets the customized logging `writer` for logging. The `writer` object should implement the io.Writer interface. Developer can use customized logging `writer` to redirect logging output to another service, eg: kafka, mysql, mongodb, etc.

##### (*Logger) SetWriterColorEnable 

``` go
func (l *Logger) SetWriterColorEnable(enabled bool)
```

SetWriterColorEnable enables file/writer logging with color.

##### (*Logger) Skip 

``` go
func (l *Logger) Skip(skip int) *Logger
```

Skip is a chaining function, which sets stack skip for the current logging content output. It also affects the caller file path checks when line number printing enabled.

##### (*Logger) Stack 

``` go
func (l *Logger) Stack(enabled bool, skip ...int) *Logger
```

Stack is a chaining function, which sets stack options for the current logging content output .

##### (*Logger) StackWithFilter 

``` go
func (l *Logger) StackWithFilter(filter string) *Logger
```

StackWithFilter is a chaining function, which sets stack filter for the current logging content output .

##### (*Logger) Stdout 

``` go
func (l *Logger) Stdout(enabled ...bool) *Logger
```

Stdout is a chaining function, which enables/disables stdout for the current logging content output. It's enabled in default.

##### (*Logger) To 

``` go
func (l *Logger) To(writer io.Writer) *Logger
```

To is a chaining function, which redirects current logging content output to the specified `writer`.

##### (*Logger) Warning 

``` go
func (l *Logger) Warning(ctx context.Context, v ...interface{})
```

Warning prints the logging content with [WARN] header and newline. It also prints caller stack info if stack feature is enabled.

##### (*Logger) Warningf 

``` go
func (l *Logger) Warningf(ctx context.Context, format string, v ...interface{})
```

Warningf prints the logging content with [WARN] header, custom format and newline. It also prints caller stack info if stack feature is enabled.

##### (*Logger) Write 

``` go
func (l *Logger) Write(p []byte) (n int, err error)
```

Write implements the io.Writer interface. It just prints the content using Print.