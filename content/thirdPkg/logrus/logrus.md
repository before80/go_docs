+++
title = "logrus文档"
date = 2023-06-05T11:23:39+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Logrus 

> 原文：[https://pkg.go.dev/github.com/sirupsen/logrus](https://pkg.go.dev/github.com/sirupsen/logrus)
>
> 版本：v1.9.3 
>
> 发布时间：2023.5.21

# ![img](logrus_img/hTeVwmJ.png)

Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.

**Logrus is in maintenance-mode.** We will not be introducing new features. It's simply too hard to do in a way that won't break many people's projects, which is the last thing you want from your Logging library (again...).

This does not mean Logrus is dead. Logrus will continue to be maintained for security, (backwards compatible) bug fixes, and performance (where we are limited by the interface).

I believe Logrus' biggest contribution is to have played a part in today's widespread use of structured logging in Golang. There doesn't seem to be a reason to do a major, breaking iteration into Logrus V2, since the fantastic Go community has built those independently. Many fantastic alternatives have sprung up. Logrus would look like those, had it been re-designed with what we know about structured logging in Go today. Check out, for example, [Zerolog](https://github.com/rs/zerolog), [Zap](https://github.com/uber-go/zap), and [Apex](https://github.com/apex/log).

**Seeing weird case-sensitive problems?** It's in the past been possible to import Logrus as both upper- and lower-case. Due to the Go package environment, this caused issues in the community and we needed a standard. Some environments experienced problems with the upper-case variant, so the lower-case was decided. Everything using `logrus` will need to use the lower-case: `github.com/sirupsen/logrus`. Any package that isn't, should be changed.

To fix Glide, see [these comments](https://github.com/sirupsen/logrus/issues/553#issuecomment-306591437). For an in-depth explanation of the casing issue, see [this comment](https://github.com/sirupsen/logrus/issues/570#issuecomment-313933276).

Nicely color-coded in development (when a TTY is attached, otherwise just plain text):

![Colored](logrus_img/PY7qMwd.png)

With `log.SetFormatter(&log.JSONFormatter{})`, for easy parsing by logstash or Splunk:

```
{"animal":"walrus","level":"info","msg":"A group of walrus emerges from the
ocean","size":10,"time":"2014-03-10 19:57:38.562264131 -0400 EDT"}

{"level":"warning","msg":"The group's number increased tremendously!",
"number":122,"omg":true,"time":"2014-03-10 19:57:38.562471297 -0400 EDT"}

{"animal":"walrus","level":"info","msg":"A giant walrus appears!",
"size":10,"time":"2014-03-10 19:57:38.562500591 -0400 EDT"}

{"animal":"walrus","level":"info","msg":"Tremendously sized cow enters the ocean.",
"size":9,"time":"2014-03-10 19:57:38.562527896 -0400 EDT"}

{"level":"fatal","msg":"The ice breaks!","number":100,"omg":true,
"time":"2014-03-10 19:57:38.562543128 -0400 EDT"}
```

With the default `log.SetFormatter(&log.TextFormatter{})` when a TTY is not attached, the output is compatible with the [logfmt](http://godoc.org/github.com/kr/logfmt) format:

```
time="2015-03-26T01:27:38-04:00" level=debug msg="Started observing beach" animal=walrus number=8
time="2015-03-26T01:27:38-04:00" level=info msg="A group of walrus emerges from the ocean" animal=walrus size=10
time="2015-03-26T01:27:38-04:00" level=warning msg="The group's number increased tremendously!" number=122 omg=true
time="2015-03-26T01:27:38-04:00" level=debug msg="Temperature changes" temperature=-4
time="2015-03-26T01:27:38-04:00" level=panic msg="It's over 9000!" animal=orca size=9009
time="2015-03-26T01:27:38-04:00" level=fatal msg="The ice breaks!" err=&{0x2082280c0 map[animal:orca size:9009] 2015-03-26 01:27:38.441574009 -0400 EDT panic It's over 9000!} number=100 omg=true
```

To ensure this behaviour even if a TTY is attached, set your formatter as follows:

```
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
```

###### Logging Method Name

If you wish to add the calling method as a field, instruct the logger via:

```
log.SetReportCaller(true)
```

This adds the caller as 'method' like so:

```
{"animal":"penguin","level":"fatal","method":"github.com/sirupsen/arcticcreatures.migrate","msg":"a penguin swims by",
"time":"2014-03-10 19:57:38.562543129 -0400 EDT"}
time="2015-03-26T01:27:38-04:00" level=fatal method=github.com/sirupsen/arcticcreatures.migrate msg="a penguin swims by" animal=penguin
```

Note that this does add measurable overhead - the cost will depend on the version of Go, but is between 20 and 40% in recent tests with 1.6 and 1.7. You can validate this in your environment via benchmarks:

```
go test -bench=.*CallerTracing
```

###### Case-sensitivity

The organization's name was changed to lower-case--and this will not be changed back. If you are getting import conflicts due to case sensitivity, please use the lower-case import: `github.com/sirupsen/logrus`.

###### Example

The simplest way to use Logrus is simply the package-level exported logger:

``` go
package main

import (
  log "github.com/sirupsen/logrus"
)

func main() {
  log.WithFields(log.Fields{
    "animal": "walrus",
  }).Info("A walrus appears")
}
```

Note that it's completely api-compatible with the stdlib logger, so you can replace your `log` imports everywhere with `log "github.com/sirupsen/logrus"` and you'll now have the flexibility of Logrus. You can customize it all you want:

``` go
package main

import (
  "os"
  log "github.com/sirupsen/logrus"
)

func init() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  log.SetLevel(log.WarnLevel)
}

func main() {
  log.WithFields(log.Fields{
    "animal": "walrus",
    "size":   10,
  }).Info("A group of walrus emerges from the ocean")

  log.WithFields(log.Fields{
    "omg":    true,
    "number": 122,
  }).Warn("The group's number increased tremendously!")

  log.WithFields(log.Fields{
    "omg":    true,
    "number": 100,
  }).Fatal("The ice breaks!")

  // A common pattern is to re-use fields between logging statements by re-using
  // the logrus.Entry returned from WithFields()
  contextLogger := log.WithFields(log.Fields{
    "common": "this is a common field",
    "other": "I also should be logged always",
  })

  contextLogger.Info("I'll be logged with common and other field")
  contextLogger.Info("Me too")
}
```

For more advanced usage such as logging to multiple locations from the same application, you can also create an instance of the `logrus` Logger:

``` go
package main

import (
  "os"
  "github.com/sirupsen/logrus"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func main() {
  // The API for setting attributes is a little different than the package level
  // exported logger. See Godoc.
  log.Out = os.Stdout

  // You could set this to any `io.Writer` such as a file
  // file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  // if err == nil {
  //  log.Out = file
  // } else {
  //  log.Info("Failed to log to file, using default stderr")
  // }

  log.WithFields(logrus.Fields{
    "animal": "walrus",
    "size":   10,
  }).Info("A group of walrus emerges from the ocean")
}
```

###### Fields

Logrus encourages careful, structured logging through logging fields instead of long, unparseable error messages. For example, instead of: `log.Fatalf("Failed to send event %s to topic %s with key %d")`, you should log the much more discoverable:

```
log.WithFields(log.Fields{
  "event": event,
  "topic": topic,
  "key": key,
}).Fatal("Failed to send event")
```

We've found this API forces you to think about logging in a way that produces much more useful logging messages. We've been in countless situations where just a single added field to a log statement that was already there would've saved us hours. The `WithFields` call is optional.

In general, with Logrus using any of the `printf`-family functions should be seen as a hint you should add a field, however, you can still use the `printf`-family functions with Logrus.

###### Default Fields

Often it's helpful to have fields *always* attached to log statements in an application or parts of one. For example, you may want to always log the `request_id` and `user_ip` in the context of a request. Instead of writing `log.WithFields(log.Fields{"request_id": request_id, "user_ip": user_ip})` on every line, you can create a `logrus.Entry` to pass around instead:

```
requestLogger := log.WithFields(log.Fields{"request_id": request_id, "user_ip": user_ip})
requestLogger.Info("something happened on that request") # will log request_id and user_ip
requestLogger.Warn("something not great happened")
```

###### Hooks

You can add hooks for logging levels. For example to send errors to an exception tracking service on `Error`, `Fatal` and `Panic`, info to StatsD or log to multiple places simultaneously, e.g. syslog.

Logrus comes with [built-in hooks](https://github.com/sirupsen/logrus/blob/v1.9.3/hooks). Add those, or your custom hook, in `init`:

```
import (
  log "github.com/sirupsen/logrus"
  "gopkg.in/gemnasium/logrus-airbrake-hook.v2" // the package is named "airbrake"
  logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
  "log/syslog"
)

func init() {

  // Use the Airbrake hook to report errors that have Error severity or above to
  // an exception tracker. You can create custom hooks, see the Hooks section.
  log.AddHook(airbrake.NewHook(123, "xyz", "production"))

  hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")
  if err != nil {
    log.Error("Unable to connect to local syslog daemon")
  } else {
    log.AddHook(hook)
  }
}
```

Note: Syslog hook also support connecting to local syslog (Ex. "/dev/log" or "/var/run/syslog" or "/var/run/log"). For the detail, please check the [syslog hook README](https://github.com/sirupsen/logrus/blob/v1.9.3/hooks/syslog/README.md).

A list of currently known service hooks can be found in this wiki [page](https://github.com/sirupsen/logrus/wiki/Hooks)

###### Level logging

Logrus has seven logging levels: Trace, Debug, Info, Warning, Error, Fatal and Panic.

```
log.Trace("Something very low level.")
log.Debug("Useful debugging information.")
log.Info("Something noteworthy happened!")
log.Warn("You should probably take a look at this.")
log.Error("Something failed but I'm not quitting.")
// Calls os.Exit(1) after logging
log.Fatal("Bye.")
// Calls panic() after logging
log.Panic("I'm bailing.")
```

You can set the logging level on a `Logger`, then it will only log entries with that severity or anything above it:

```
// Will log anything that is info or above (warn, error, fatal, panic). Default.
log.SetLevel(log.InfoLevel)
```

It may be useful to set `log.Level = logrus.DebugLevel` in a debug or verbose environment if your application has that.

Note: If you want different log levels for global (`log.SetLevel(...)`) and syslog logging, please check the [syslog hook README](https://github.com/sirupsen/logrus/blob/v1.9.3/hooks/syslog/README.md).

###### Entries

Besides the fields added with `WithField` or `WithFields` some fields are automatically added to all logging events:

1. `time`. The timestamp when the entry was created.
2. `msg`. The logging message passed to `{Info,Warn,Error,Fatal,Panic}` after the `AddFields` call. E.g. `Failed to send event.`
3. `level`. The logging level. E.g. `info`.

###### Environments

Logrus has no notion of environment.

If you wish for hooks and formatters to only be used in specific environments, you should handle that yourself. For example, if your application has a global variable `Environment`, which is a string representation of the environment you could do:

```
import (
  log "github.com/sirupsen/logrus"
)

func init() {
  // do something here to set environment depending on an environment variable
  // or command-line flag
  if Environment == "production" {
    log.SetFormatter(&log.JSONFormatter{})
  } else {
    // The TextFormatter is default, you don't actually have to do this.
    log.SetFormatter(&log.TextFormatter{})
  }
}
```

This configuration is how `logrus` was intended to be used, but JSON in production is mostly only useful if you do log aggregation with tools like Splunk or Logstash.

###### Formatters

The built-in logging formatters are:

- ```
  logrus.TextFormatter
  ```

  . Logs the event in colors if stdout is a tty, otherwise without colors.

  - *Note:* to force colored output when there is no TTY, set the `ForceColors` field to `true`. To force no colored output even if there is a TTY set the `DisableColors` field to `true`. For Windows, see [github.com/mattn/go-colorable](https://github.com/mattn/go-colorable).
  - When colors are enabled, levels are truncated to 4 characters by default. To disable truncation set the `DisableLevelTruncation` field to `true`.
  - When outputting to a TTY, it's often helpful to visually scan down a column where all the levels are the same width. Setting the `PadLevelText` field to `true` enables this behavior, by adding padding to the level text.
  - All options are listed in the [generated docs](https://godoc.org/github.com/sirupsen/logrus#TextFormatter).

- ```
  logrus.JSONFormatter
  ```

  . Logs fields as JSON.

  - All options are listed in the [generated docs](https://godoc.org/github.com/sirupsen/logrus#JSONFormatter).

Third party logging formatters:

- [`FluentdFormatter`](https://github.com/joonix/log). Formats entries that can be parsed by Kubernetes and Google Container Engine.
- [`GELF`](https://github.com/fabienm/go-logrus-formatters). Formats entries so they comply to Graylog's [GELF 1.1 specification](http://docs.graylog.org/en/2.4/pages/gelf.html).
- [`logstash`](https://github.com/bshuster-repo/logrus-logstash-hook). Logs fields as [Logstash](http://logstash.net/) Events.
- [`prefixed`](https://github.com/x-cray/logrus-prefixed-formatter). Displays log entry source along with alternative layout.
- [`zalgo`](https://github.com/aybabtme/logzalgo). Invoking the Power of Zalgo.
- [`nested-logrus-formatter`](https://github.com/antonfisher/nested-logrus-formatter). Converts logrus fields to a nested structure.
- [`powerful-logrus-formatter`](https://github.com/zput/zxcTool). get fileName, log's line number and the latest function's name when print log; Sava log to files.
- [`caption-json-formatter`](https://github.com/nolleh/caption_json_formatter). logrus's message json formatter with human-readable caption added.

You can define your formatter by implementing the `Formatter` interface, requiring a `Format` method. `Format` takes an `*Entry`. `entry.Data` is a `Fields` type (`map[string]interface{}`) with all your fields as well as the default ones (see Entries section above):

``` go
type MyJSONFormatter struct {
}

log.SetFormatter(new(MyJSONFormatter))

func (f *MyJSONFormatter) Format(entry *Entry) ([]byte, error) {
  // Note this doesn't include Time, Level and Message which are available on
  // the Entry. Consult `godoc` on information about those fields or read the
  // source of the official loggers.
  serialized, err := json.Marshal(entry.Data)
    if err != nil {
      return nil, fmt.Errorf("Failed to marshal fields to JSON, %w", err)
    }
  return append(serialized, '\n'), nil
}
```

###### Logger as an `io.Writer`

Logrus can be transformed into an `io.Writer`. That writer is the end of an `io.Pipe` and it is your responsibility to close it.

```
w := logger.Writer()
defer w.Close()

srv := http.Server{
    // create a stdlib log.Logger that writes to
    // logrus.Logger.
    ErrorLog: log.New(w, "", 0),
}
```

Each line written to that writer will be printed the usual way, using formatters and hooks. The level for those entries is `info`.

This means that we can override the standard library logger easily:

```
logger := logrus.New()
logger.Formatter = &logrus.JSONFormatter{}

// Use logrus for standard log output
// Note that `log` here references stdlib's log
// Not logrus imported under the name `log`.
log.SetOutput(logger.Writer())
```

###### Rotation

Log rotation is not provided with Logrus. Log rotation should be done by an external program (like `logrotate(8)`) that can compress and delete old log entries. It should not be a feature of the application-level logger.

###### Tools

| Tool                                                         | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [Logrus Mate](https://github.com/gogap/logrus_mate)          | Logrus mate is a tool for Logrus to manage loggers, you can initial logger's level, hook and formatter by config file, the logger will be generated with different configs in different environments. |
| [Logrus Viper Helper](https://github.com/heirko/go-contrib/tree/master/logrusHelper) | An Helper around Logrus to wrap with spf13/Viper to load configuration with fangs! And to simplify Logrus configuration use some behavior of [Logrus Mate](https://github.com/gogap/logrus_mate). [sample](https://github.com/heirko/iris-contrib/raw/master/middleware/logrus-logger/example) |

###### Testing

Logrus has a built in facility for asserting the presence of log messages. This is implemented through the `test` hook and provides:

- decorators for existing logger (`test.NewLocal` and `test.NewGlobal`) which basically just adds the `test` hook
- a test logger (`test.NewNullLogger`) that just records log messages (and does not output any):

```
import(
  "github.com/sirupsen/logrus"
  "github.com/sirupsen/logrus/hooks/test"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSomething(t*testing.T){
  logger, hook := test.NewNullLogger()
  logger.Error("Helloerror")

  assert.Equal(t, 1, len(hook.Entries))
  assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
  assert.Equal(t, "Helloerror", hook.LastEntry().Message)

  hook.Reset()
  assert.Nil(t, hook.LastEntry())
}
```

###### Fatal handlers

Logrus can register one or more functions that will be called when any `fatal` level message is logged. The registered handlers will be executed before logrus performs an `os.Exit(1)`. This behavior may be helpful if callers need to gracefully shutdown. Unlike a `panic("Something went wrong...")` call which can be intercepted with a deferred `recover` a call to `os.Exit(1)` can not be intercepted.

```
...
handler := func() {
  // gracefully shutdown something...
}
logrus.RegisterExitHandler(handler)
...
```

###### Thread safety

By default, Logger is protected by a mutex for concurrent writes. The mutex is held when calling hooks and writing logs. If you are sure such locking is not needed, you can call logger.SetNoLock() to disable the locking.

Situation when locking is not needed includes:

- You have no hooks registered, or hooks calling is already thread-safe.

- Writing to logger.Out is already thread-safe, for example:

  1. logger.Out is protected by locks.

  2. logger.Out is an os.File handler opened with `O_APPEND` flag, and every write is smaller than 4k. (This allows multi-thread/multi-process writing)

     (Refer to http://www.notthewizard.com/2014/06/17/are-files-appends-really-atomic/)

Collapse ▴

## ![img](logrus_img/code_gm_grey_24dp.svg+xml) Documentation 

[Rendered for](https://go.dev/about#build-context)                   linux/amd64                   windows/amd64                   darwin/amd64                   js/wasm                

### Overview 

Package logrus is a structured logger for Go, completely API compatible with the standard library logger.

The simplest way to use Logrus is simply the package-level exported logger:

``` go
package main

import (
  log "github.com/sirupsen/logrus"
)

func main() {
  log.WithFields(log.Fields{
    "animal": "walrus",
    "number": 1,
    "size":   10,
  }).Info("A walrus appears")
}
```

Output:

```
time="2015-09-07T08:48:33Z" level=info msg="A walrus appears" animal=walrus number=1 size=10
```

For a full guide visit https://github.com/sirupsen/logrus

##### Example (Basic)

``` go
```
##### Example (Hook)

``` go

```
### Constants 

[View Source](https://github.com/sirupsen/logrus/blob/v1.9.3/formatter.go#L6)

``` go
const (
	FieldKeyMsg         = "msg"
	FieldKeyLevel       = "level"
	FieldKeyTime        = "time"
	FieldKeyLogrusError = "logrus_error"
	FieldKeyFunc        = "func"
	FieldKeyFile        = "file"
)
```

Default key names for the default fields

### Variables 

[View Source](https://github.com/sirupsen/logrus/blob/v1.9.3/logrus.go#L81)

```
var AllLevels = []Level{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}
```

A constant exposing all logging levels

[View Source](https://github.com/sirupsen/logrus/blob/v1.9.3/entry.go#L38)

```
var ErrorKey = "error"
```

Defines the key when adding errors using WithError.

### Functions 

#### func AddHook <-0.4.0

``` go
func AddHook(hook Hook)
```

AddHook adds a hook to the standard logger hooks.

#### func Debug 

``` go
func Debug(args ...interface{})
```

Debug logs a message at level Debug on the standard logger.

#### func DebugFn <-1.7.0

``` go
func DebugFn(fn LogFunction)
```

DebugFn logs a message from a func at level Debug on the standard logger.

#### func Debugf <-0.4.1

``` go
func Debugf(format string, args ...interface{})
```

Debugf logs a message at level Debug on the standard logger.

#### func Debugln <-0.4.1

``` go
func Debugln(args ...interface{})
```

Debugln logs a message at level Debug on the standard logger.

#### func DeferExitHandler <-1.4.0

``` go
func DeferExitHandler(handler func())
```

DeferExitHandler prepends a Logrus Exit handler to the list of handlers, call logrus.Exit to invoke all handlers. The handlers will also be invoked when any Fatal log entry is made.

This method is useful when a caller wishes to use logrus to log a fatal message but also needs to gracefully shutdown. An example usecase could be closing database connections, or sending a alert that the application is closing.

#### func Error 

``` go
func Error(args ...interface{})
```

Error logs a message at level Error on the standard logger.

#### func ErrorFn <-1.7.0

``` go
func ErrorFn(fn LogFunction)
```

ErrorFn logs a message from a func at level Error on the standard logger.

#### func Errorf <-0.4.1

``` go
func Errorf(format string, args ...interface{})
```

Errorf logs a message at level Error on the standard logger.

#### func Errorln <-0.4.1

``` go
func Errorln(args ...interface{})
```

Errorln logs a message at level Error on the standard logger.

#### func Exit <-0.11.0

``` go
func Exit(code int)
```

Exit runs all the Logrus atexit handlers and then terminates the program using os.Exit(code)

#### func Fatal 

``` go
func Fatal(args ...interface{})
```

Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.

#### func FatalFn <-1.7.0

``` go
func FatalFn(fn LogFunction)
```

FatalFn logs a message from a func at level Fatal on the standard logger then the process will exit with status set to 1.

#### func Fatalf <-0.4.1

``` go
func Fatalf(format string, args ...interface{})
```

Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.

#### func Fatalln <-0.4.1

``` go
func Fatalln(args ...interface{})
```

Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.

#### func Info 

``` go
func Info(args ...interface{})
```

Info logs a message at level Info on the standard logger.

#### func InfoFn <-1.7.0

``` go
func InfoFn(fn LogFunction)
```

InfoFn logs a message from a func at level Info on the standard logger.

#### func Infof <-0.4.1

``` go
func Infof(format string, args ...interface{})
```

Infof logs a message at level Info on the standard logger.

#### func Infoln <-0.4.1

``` go
func Infoln(args ...interface{})
```

Infoln logs a message at level Info on the standard logger.

#### func IsLevelEnabled <-1.1.0

``` go
func IsLevelEnabled(level Level) bool
```

IsLevelEnabled checks if the log level of the standard logger is greater than the level param

#### func Panic 

``` go
func Panic(args ...interface{})
```

Panic logs a message at level Panic on the standard logger.

#### func PanicFn <-1.7.0

``` go
func PanicFn(fn LogFunction)
```

PanicFn logs a message from a func at level Panic on the standard logger.

#### func Panicf <-0.4.1

``` go
func Panicf(format string, args ...interface{})
```

Panicf logs a message at level Panic on the standard logger.

#### func Panicln <-0.4.1

``` go
func Panicln(args ...interface{})
```

Panicln logs a message at level Panic on the standard logger.

#### func Print <-0.4.1

``` go
func Print(args ...interface{})
```

Print logs a message at level Info on the standard logger.

#### func PrintFn <-1.7.0

``` go
func PrintFn(fn LogFunction)
```

PrintFn logs a message from a func at level Info on the standard logger.

#### func Printf <-0.4.1

``` go
func Printf(format string, args ...interface{})
```

Printf logs a message at level Info on the standard logger.

#### func Println <-0.4.1

``` go
func Println(args ...interface{})
```

Println logs a message at level Info on the standard logger.

#### func RegisterExitHandler <-0.11.0

``` go
func RegisterExitHandler(handler func())
```

RegisterExitHandler appends a Logrus Exit handler to the list of handlers, call logrus.Exit to invoke all handlers. The handlers will also be invoked when any Fatal log entry is made.

This method is useful when a caller wishes to use logrus to log a fatal message but also needs to gracefully shutdown. An example usecase could be closing database connections, or sending a alert that the application is closing.

#### func SetBufferPool <-1.7.0

``` go
func SetBufferPool(bp BufferPool)
```

SetBufferPool allows to replace the default logrus buffer pool to better meets the specific needs of an application.

#### func SetFormatter <-0.4.0

``` go
func SetFormatter(formatter Formatter)
```

SetFormatter sets the standard logger formatter.

#### func SetLevel <-0.4.0

``` go
func SetLevel(level Level)
```

SetLevel sets the standard logger level.

#### func SetOutput <-0.4.0

``` go
func SetOutput(out io.Writer)
```

SetOutput sets the standard logger output.

#### func SetReportCaller <-1.2.0

``` go
func SetReportCaller(include bool)
```

SetReportCaller sets whether the standard logger will include the calling method as a field.

#### func Trace <-1.2.0

``` go
func Trace(args ...interface{})
```

Trace logs a message at level Trace on the standard logger.

#### func TraceFn <-1.7.0

``` go
func TraceFn(fn LogFunction)
```

TraceFn logs a message from a func at level Trace on the standard logger.

#### func Tracef <-1.2.0

``` go
func Tracef(format string, args ...interface{})
```

Tracef logs a message at level Trace on the standard logger.

#### func Traceln <-1.2.0

``` go
func Traceln(args ...interface{})
```

Traceln logs a message at level Trace on the standard logger.

#### func Warn 

``` go
func Warn(args ...interface{})
```

Warn logs a message at level Warn on the standard logger.

#### func WarnFn <-1.7.0

``` go
func WarnFn(fn LogFunction)
```

WarnFn logs a message from a func at level Warn on the standard logger.

#### func Warnf <-0.4.1

``` go
func Warnf(format string, args ...interface{})
```

Warnf logs a message at level Warn on the standard logger.

#### func Warning <-0.4.1

``` go
func Warning(args ...interface{})
```

Warning logs a message at level Warn on the standard logger.

#### func WarningFn <-1.7.0

``` go
func WarningFn(fn LogFunction)
```

WarningFn logs a message from a func at level Warn on the standard logger.

#### func Warningf <-0.4.1

``` go
func Warningf(format string, args ...interface{})
```

Warningf logs a message at level Warn on the standard logger.

#### func Warningln <-0.4.1

``` go
func Warningln(args ...interface{})
```

Warningln logs a message at level Warn on the standard logger.

#### func Warnln <-0.4.1

``` go
func Warnln(args ...interface{})
```

Warnln logs a message at level Warn on the standard logger.

### Types 

#### type BufferPool <-1.7.0

``` go
type BufferPool interface {
	Put(*bytes.Buffer)
	Get() *bytes.Buffer
}
```

#### type Entry 

``` go
type Entry struct {
	Logger *Logger

	// Contains all the fields set by the user.
	Data Fields

	// Time at which the log entry was created
	Time time.Time

	// Level the log entry was logged at: Trace, Debug, Info, Warn, Error, Fatal or Panic
	// This field will be set on entry firing and the value will be equal to the one in Logger struct field.
	Level Level

	// Calling method, with package name
	Caller *runtime.Frame

	// Message passed to Trace, Debug, Info, Warn, Error, Fatal or Panic
	Message string

	// When formatter is called in entry.log(), a Buffer may be set to entry
	Buffer *bytes.Buffer

	// Contains the context set by the user. Useful for hook processing etc.
	Context context.Context
	// contains filtered or unexported fields
}
```

An entry is the final or intermediate Logrus logging entry. It contains all the fields passed with WithField{,s}. It's finally logged when Trace, Debug, Info, Warn, Error, Fatal or Panic is called on it. These objects can be reused and passed around as much as you wish to avoid field duplication.

#### func NewEntry 

``` go
func NewEntry(logger *Logger) *Entry
```

#### func WithContext <-1.4.0

``` go
func WithContext(ctx context.Context) *Entry
```

WithContext creates an entry from the standard logger and adds a context to it.

#### func WithError <-0.8.7

``` go
func WithError(err error) *Entry
```

WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.

#### func WithField <-0.4.0

``` go
func WithField(key string, value interface{}) *Entry
```

WithField creates an entry from the standard logger and adds a field to it. If you want multiple fields, use `WithFields`.

Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal or Panic on the Entry it returns.

#### func WithFields <-0.4.0

``` go
func WithFields(fields Fields) *Entry
```

WithFields creates an entry from the standard logger and adds multiple fields to it. This is simply a helper for `WithField`, invoking it once for each field.

Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal or Panic on the Entry it returns.

#### func WithTime <-1.0.6

``` go
func WithTime(t time.Time) *Entry
```

WithTime creates an entry from the standard logger and overrides the time of logs generated with it.

Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal or Panic on the Entry it returns.

#### (*Entry) Bytes <-1.5.0

``` go
func (entry *Entry) Bytes() ([]byte, error)
```

Returns the bytes representation of this entry from the formatter.

#### (*Entry) Debug 

``` go
func (entry *Entry) Debug(args ...interface{})
```

#### (*Entry) Debugf 

``` go
func (entry *Entry) Debugf(format string, args ...interface{})
```

#### (*Entry) Debugln 

``` go
func (entry *Entry) Debugln(args ...interface{})
```

#### (*Entry) Dup <-1.7.1

``` go
func (entry *Entry) Dup() *Entry
```

#### (*Entry) Error 

``` go
func (entry *Entry) Error(args ...interface{})
```

#### (*Entry) Errorf 

``` go
func (entry *Entry) Errorf(format string, args ...interface{})
```

#### (*Entry) Errorln 

``` go
func (entry *Entry) Errorln(args ...interface{})
```

#### (*Entry) Fatal 

``` go
func (entry *Entry) Fatal(args ...interface{})
```

#### (*Entry) Fatalf 

``` go
func (entry *Entry) Fatalf(format string, args ...interface{})
```

#### (*Entry) Fatalln 

``` go
func (entry *Entry) Fatalln(args ...interface{})
```

#### func (Entry) [HasCaller](https://github.com/sirupsen/logrus/blob/v1.9.3/entry.go#L215) <-1.2.0

``` go
func (entry Entry) HasCaller() (has bool)
```

#### (*Entry) Info 

``` go
func (entry *Entry) Info(args ...interface{})
```

#### (*Entry) Infof 

``` go
func (entry *Entry) Infof(format string, args ...interface{})
```

#### (*Entry) Infoln 

``` go
func (entry *Entry) Infoln(args ...interface{})
```

#### (*Entry) Log <-1.3.0

``` go
func (entry *Entry) Log(level Level, args ...interface{})
```

Log will log a message at the level given as parameter. Warning: using Log at Panic or Fatal level will not respectively Panic nor Exit. For this behaviour Entry.Panic or Entry.Fatal should be used instead.

#### (*Entry) Logf <-1.3.0

``` go
func (entry *Entry) Logf(level Level, format string, args ...interface{})
```

#### (*Entry) Logln <-1.3.0

``` go
func (entry *Entry) Logln(level Level, args ...interface{})
```

#### (*Entry) Panic 

``` go
func (entry *Entry) Panic(args ...interface{})
```

#### (*Entry) Panicf 

``` go
func (entry *Entry) Panicf(format string, args ...interface{})
```

#### (*Entry) Panicln 

``` go
func (entry *Entry) Panicln(args ...interface{})
```

#### (*Entry) Print 

``` go
func (entry *Entry) Print(args ...interface{})
```

#### (*Entry) Printf 

``` go
func (entry *Entry) Printf(format string, args ...interface{})
```

#### (*Entry) Println 

``` go
func (entry *Entry) Println(args ...interface{})
```

#### (*Entry) String 

``` go
func (entry *Entry) String() (string, error)
```

Returns the string representation from the reader and ultimately the formatter.

#### (*Entry) Trace <-1.2.0

``` go
func (entry *Entry) Trace(args ...interface{})
```

#### (*Entry) Tracef <-1.2.0

``` go
func (entry *Entry) Tracef(format string, args ...interface{})
```

#### (*Entry) Traceln <-1.2.0

``` go
func (entry *Entry) Traceln(args ...interface{})
```

#### (*Entry) Warn 

``` go
func (entry *Entry) Warn(args ...interface{})
```

#### (*Entry) Warnf 

``` go
func (entry *Entry) Warnf(format string, args ...interface{})
```

#### (*Entry) Warning <-0.6.3

``` go
func (entry *Entry) Warning(args ...interface{})
```

#### (*Entry) Warningf 

``` go
func (entry *Entry) Warningf(format string, args ...interface{})
```

#### (*Entry) Warningln 

``` go
func (entry *Entry) Warningln(args ...interface{})
```

#### (*Entry) Warnln 

``` go
func (entry *Entry) Warnln(args ...interface{})
```

#### (*Entry) WithContext <-1.4.0

``` go
func (entry *Entry) WithContext(ctx context.Context) *Entry
```

Add a context to the Entry.

#### (*Entry) WithError <-0.8.7

``` go
func (entry *Entry) WithError(err error) *Entry
```

Add an error as single field (using the key defined in ErrorKey) to the Entry.

#### (*Entry) WithField 

``` go
func (entry *Entry) WithField(key string, value interface{}) *Entry
```

Add a single field to the Entry.

#### (*Entry) WithFields 

``` go
func (entry *Entry) WithFields(fields Fields) *Entry
```

Add a map of fields to the Entry.

#### (*Entry) WithTime <-1.0.6

``` go
func (entry *Entry) WithTime(t time.Time) *Entry
```

Overrides the time of the Entry.

#### (*Entry) Writer <-0.11.5

``` go
func (entry *Entry) Writer() *io.PipeWriter
```

Writer returns an io.Writer that writes to the logger at the info log level

#### (*Entry) WriterLevel <-0.11.5

``` go
func (entry *Entry) WriterLevel(level Level) *io.PipeWriter
```

WriterLevel returns an io.Writer that writes to the logger at the given log level

#### type Ext1FieldLogger <-1.2.0

``` go
type Ext1FieldLogger interface {
	FieldLogger
	Tracef(format string, args ...interface{})
	Trace(args ...interface{})
	Traceln(args ...interface{})
}
```

Ext1FieldLogger (the first extension to FieldLogger) is superfluous, it is here for consistancy. Do not use. Use Logger or Entry instead.

#### type FieldLogger <-0.10.0

``` go
type FieldLogger interface {
	WithField(key string, value interface{}) *Entry
	WithFields(fields Fields) *Entry
	WithError(err error) *Entry

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}
```

The FieldLogger interface generalizes the Entry and Logger types

#### type FieldMap <-0.11.1

``` go
type FieldMap map[fieldKey]string
```

FieldMap allows customization of the key names for default fields.

#### type Fields 

``` go
type Fields map[string]interface{}
```

Fields type, used to pass to `WithFields`.

#### type Formatter 

``` go
type Formatter interface {
	Format(*Entry) ([]byte, error)
}
```

The Formatter interface is used to implement a custom Formatter. It takes an `Entry`. It exposes all the fields, including the default ones:

\* `entry.Data["msg"]`. The message passed from Info, Warn, Error .. * `entry.Data["time"]`. The timestamp. * `entry.Data["level"]. The level the entry was logged at.

Any additional fields added with `WithField` or `WithFields` are also in `entry.Data`. Format is expected to return an array of bytes which are then logged to `logger.Out`.

#### type Hook 

``` go
type Hook interface {
	Levels() []Level
	Fire(*Entry) error
}
```

A hook to be fired when logging on the logging levels returned from `Levels()` on your implementation of the interface. Note that this is not fired in a goroutine or a channel with workers, you should handle such functionality yourself if your call is non-blocking and you don't wish for the logging calls for levels returned from `Levels()` to block.

#### type JSONFormatter 

``` go
type JSONFormatter struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	// The format to use is the same than for time.Format or time.Parse from the standard
	// library.
	// The standard Library already provides a set of predefined format.
	TimestampFormat string

	// DisableTimestamp allows disabling automatic timestamps in output
	DisableTimestamp bool

	// DisableHTMLEscape allows disabling html escaping in output
	DisableHTMLEscape bool

	// DataKey allows users to put all the log entry parameters into a nested dictionary at a given key.
	DataKey string

	// FieldMap allows users to customize the names of keys for default fields.
	// As an example:
	// formatter := &JSONFormatter{
	//   	FieldMap: FieldMap{
	// 		 FieldKeyTime:  "@timestamp",
	// 		 FieldKeyLevel: "@level",
	// 		 FieldKeyMsg:   "@message",
	// 		 FieldKeyFunc:  "@caller",
	//    },
	// }
	FieldMap FieldMap

	// CallerPrettyfier can be set by the user to modify the content
	// of the function and file keys in the json data when ReportCaller is
	// activated. If any of the returned value is the empty string the
	// corresponding key will be removed from json fields.
	CallerPrettyfier func(*runtime.Frame) (function string, file string)

	// PrettyPrint will indent all json logs
	PrettyPrint bool
}
```

JSONFormatter formats logs into parsable json

#### (*JSONFormatter) Format 

``` go
func (f *JSONFormatter) Format(entry *Entry) ([]byte, error)
```

Format renders a single log entry

#### type Level 

``` go
type Level uint32
```

Level type

``` go
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)
```

These are the different logging levels. You can set the logging level to log on your instance of logger, obtained with `logrus.New()`.

#### func GetLevel <-0.6.1

``` go
func GetLevel() Level
```

GetLevel returns the standard logger level.

#### func ParseLevel <-0.6.0

``` go
func ParseLevel(lvl string) (Level, error)
```

ParseLevel takes a string level and returns the Logrus log level constant.

#### func (Level) [MarshalText](https://github.com/sirupsen/logrus/blob/v1.9.3/logrus.go#L59) <-1.3.0

``` go
func (level Level) MarshalText() ([]byte, error)
```

#### func (Level) [String](https://github.com/sirupsen/logrus/blob/v1.9.3/logrus.go#L16) <-0.4.0

``` go
func (level Level) String() string
```

Convert the Level to a string. E.g. PanicLevel becomes "panic".

#### (*Level) UnmarshalText <-1.2.0

``` go
func (level *Level) UnmarshalText(text []byte) error
```

UnmarshalText implements encoding.TextUnmarshaler.

#### type LevelHooks <-0.8.3

``` go
type LevelHooks map[Level][]Hook
```

Internal type for storing the hooks on a logger instance.

#### func (LevelHooks) [Add](https://github.com/sirupsen/logrus/blob/v1.9.3/hooks.go#L18) <-0.8.3

``` go
func (hooks LevelHooks) Add(hook Hook)
```

Add a hook to an instance of logger. This is called with `log.Hooks.Add(new(MyHook))` where `MyHook` implements the `Hook` interface.

#### func (LevelHooks) [Fire](https://github.com/sirupsen/logrus/blob/v1.9.3/hooks.go#L26) <-0.8.3

``` go
func (hooks LevelHooks) Fire(level Level, entry *Entry) error
```

Fire all the hooks for the passed level. Used by `entry.log` to fire appropriate hooks for a log entry.

#### type LogFunction <-1.7.0

``` go
type LogFunction func() []interface{}
```

LogFunction For big messages, it can be more efficient to pass a function and only call it if the log level is actually enables rather than generating the log message and then checking if the level is enabled

#### type Logger 

``` go
type Logger struct {
	// The logs are `io.Copy`'d to this in a mutex. It's common to set this to a
	// file, or leave it default which is `os.Stderr`. You can also set this to
	// something more adventurous, such as logging to Kafka.
	Out io.Writer
	// Hooks for the logger instance. These allow firing events based on logging
	// levels and log entries. For example, to send errors to an error tracking
	// service, log to StatsD or dump the core on fatal errors.
	Hooks LevelHooks
	// All log entries pass through the formatter before logged to Out. The
	// included formatters are `TextFormatter` and `JSONFormatter` for which
	// TextFormatter is the default. In development (when a TTY is attached) it
	// logs with colors, but to a file it wouldn't. You can easily implement your
	// own that implements the `Formatter` interface, see the `README` or included
	// formatters for examples.
	Formatter Formatter

	// Flag for whether to log caller info (off by default)
	ReportCaller bool

	// The logging level the logger should log at. This is typically (and defaults
	// to) `logrus.Info`, which allows Info(), Warn(), Error() and Fatal() to be
	// logged.
	Level Level

	// Function to exit the application, defaults to `os.Exit()`
	ExitFunc exitFunc
	// The buffer pool used to format the log. If it is nil, the default global
	// buffer pool will be used.
	BufferPool BufferPool
	// contains filtered or unexported fields
}
```

#### func New 

``` go
func New() *Logger
```

Creates a new logger. Configuration should be set by changing `Formatter`, `Out` and `Hooks` directly on the default logger instance. You can also just instantiate your own:

```
var log = &logrus.Logger{
  Out: os.Stderr,
  Formatter: new(logrus.TextFormatter),
  Hooks: make(logrus.LevelHooks),
  Level: logrus.DebugLevel,
}
```

It's recommended to make this a global instance called `log`.

#### func StandardLogger <-0.6.5

``` go
func StandardLogger() *Logger
```

#### (*Logger) AddHook <-1.0.4

``` go
func (logger *Logger) AddHook(hook Hook)
```

AddHook adds a hook to the logger hooks.

#### (*Logger) Debug 

``` go
func (logger *Logger) Debug(args ...interface{})
```

#### (*Logger) DebugFn <-1.7.0

``` go
func (logger *Logger) DebugFn(fn LogFunction)
```

#### (*Logger) Debugf 

``` go
func (logger *Logger) Debugf(format string, args ...interface{})
```

#### (*Logger) Debugln 

``` go
func (logger *Logger) Debugln(args ...interface{})
```

#### (*Logger) Error 

``` go
func (logger *Logger) Error(args ...interface{})
```

#### (*Logger) ErrorFn <-1.7.0

``` go
func (logger *Logger) ErrorFn(fn LogFunction)
```

#### (*Logger) Errorf 

``` go
func (logger *Logger) Errorf(format string, args ...interface{})
```

#### (*Logger) Errorln 

``` go
func (logger *Logger) Errorln(args ...interface{})
```

#### (*Logger) Exit <-1.2.0

``` go
func (logger *Logger) Exit(code int)
```

#### (*Logger) Fatal 

``` go
func (logger *Logger) Fatal(args ...interface{})
```

#### (*Logger) FatalFn <-1.7.0

``` go
func (logger *Logger) FatalFn(fn LogFunction)
```

#### (*Logger) Fatalf 

``` go
func (logger *Logger) Fatalf(format string, args ...interface{})
```

#### (*Logger) Fatalln 

``` go
func (logger *Logger) Fatalln(args ...interface{})
```

#### (*Logger) GetLevel <-1.1.0

``` go
func (logger *Logger) GetLevel() Level
```

GetLevel returns the logger level.

#### (*Logger) Info 

``` go
func (logger *Logger) Info(args ...interface{})
```

#### (*Logger) InfoFn <-1.7.0

``` go
func (logger *Logger) InfoFn(fn LogFunction)
```

#### (*Logger) Infof 

``` go
func (logger *Logger) Infof(format string, args ...interface{})
```

#### (*Logger) Infoln 

``` go
func (logger *Logger) Infoln(args ...interface{})
```

#### (*Logger) IsLevelEnabled <-1.1.0

``` go
func (logger *Logger) IsLevelEnabled(level Level) bool
```

IsLevelEnabled checks if the log level of the logger is greater than the level param

#### (*Logger) Log <-1.3.0

``` go
func (logger *Logger) Log(level Level, args ...interface{})
```

Log will log a message at the level given as parameter. Warning: using Log at Panic or Fatal level will not respectively Panic nor Exit. For this behaviour Logger.Panic or Logger.Fatal should be used instead.

#### (*Logger) LogFn <-1.7.0

``` go
func (logger *Logger) LogFn(level Level, fn LogFunction)
```

#### (*Logger) Logf <-1.3.0

``` go
func (logger *Logger) Logf(level Level, format string, args ...interface{})
```

#### (*Logger) Logln <-1.3.0

``` go
func (logger *Logger) Logln(level Level, args ...interface{})
```

#### (*Logger) Panic 

``` go
func (logger *Logger) Panic(args ...interface{})
```

#### (*Logger) PanicFn <-1.7.0

``` go
func (logger *Logger) PanicFn(fn LogFunction)
```

#### (*Logger) Panicf 

``` go
func (logger *Logger) Panicf(format string, args ...interface{})
```

#### (*Logger) Panicln 

``` go
func (logger *Logger) Panicln(args ...interface{})
```

#### (*Logger) Print 

``` go
func (logger *Logger) Print(args ...interface{})
```

#### (*Logger) PrintFn <-1.7.0

``` go
func (logger *Logger) PrintFn(fn LogFunction)
```

#### (*Logger) Printf 

``` go
func (logger *Logger) Printf(format string, args ...interface{})
```

#### (*Logger) Println 

``` go
func (logger *Logger) Println(args ...interface{})
```

#### (*Logger) ReplaceHooks <-1.1.0

``` go
func (logger *Logger) ReplaceHooks(hooks LevelHooks) LevelHooks
```

ReplaceHooks replaces the logger hooks and returns the old ones

#### (*Logger) SetBufferPool <-1.8.2

``` go
func (logger *Logger) SetBufferPool(pool BufferPool)
```

SetBufferPool sets the logger buffer pool.

#### (*Logger) SetFormatter <-1.1.0

``` go
func (logger *Logger) SetFormatter(formatter Formatter)
```

SetFormatter sets the logger formatter.

#### (*Logger) SetLevel <-1.0.3

``` go
func (logger *Logger) SetLevel(level Level)
```

SetLevel sets the logger level.

#### (*Logger) SetNoLock <-0.11.0

``` go
func (logger *Logger) SetNoLock()
```

When file is opened with appending mode, it's safe to write concurrently to a file (within 4k message on Linux). In these cases user can choose to disable the lock.

#### (*Logger) SetOutput <-1.0.6

``` go
func (logger *Logger) SetOutput(output io.Writer)
```

SetOutput sets the logger output.

#### (*Logger) SetReportCaller <-1.2.0

``` go
func (logger *Logger) SetReportCaller(reportCaller bool)
```

#### (*Logger) Trace <-1.2.0

``` go
func (logger *Logger) Trace(args ...interface{})
```

#### (*Logger) TraceFn <-1.7.0

``` go
func (logger *Logger) TraceFn(fn LogFunction)
```

#### (*Logger) Tracef <-1.2.0

``` go
func (logger *Logger) Tracef(format string, args ...interface{})
```

#### (*Logger) Traceln <-1.2.0

``` go
func (logger *Logger) Traceln(args ...interface{})
```

#### (*Logger) Warn 

``` go
func (logger *Logger) Warn(args ...interface{})
```

#### (*Logger) WarnFn <-1.7.0

``` go
func (logger *Logger) WarnFn(fn LogFunction)
```

#### (*Logger) Warnf 

``` go
func (logger *Logger) Warnf(format string, args ...interface{})
```

#### (*Logger) Warning 

``` go
func (logger *Logger) Warning(args ...interface{})
```

#### (*Logger) WarningFn <-1.7.0

``` go
func (logger *Logger) WarningFn(fn LogFunction)
```

#### (*Logger) Warningf 

``` go
func (logger *Logger) Warningf(format string, args ...interface{})
```

#### (*Logger) Warningln 

``` go
func (logger *Logger) Warningln(args ...interface{})
```

#### (*Logger) Warnln 

``` go
func (logger *Logger) Warnln(args ...interface{})
```

#### (*Logger) WithContext <-1.4.0

``` go
func (logger *Logger) WithContext(ctx context.Context) *Entry
```

Add a context to the log entry.

#### (*Logger) WithError <-0.9.0

``` go
func (logger *Logger) WithError(err error) *Entry
```

Add an error as single field to the log entry. All it does is call `WithError` for the given `error`.

#### (*Logger) WithField 

``` go
func (logger *Logger) WithField(key string, value interface{}) *Entry
```

WithField allocates a new entry and adds a field to it. Debug, Print, Info, Warn, Error, Fatal or Panic must be then applied to this new returned entry. If you want multiple fields, use `WithFields`.

#### (*Logger) WithFields 

``` go
func (logger *Logger) WithFields(fields Fields) *Entry
```

Adds a struct of fields to the log entry. All it does is call `WithField` for each `Field`.

#### (*Logger) WithTime <-1.0.6

``` go
func (logger *Logger) WithTime(t time.Time) *Entry
```

Overrides the time of the log entry.

#### (*Logger) Writer <-0.6.5

``` go
func (logger *Logger) Writer() *io.PipeWriter
```

Writer at INFO level. See WriterLevel for details.

##### Example (HttpServer)

``` go
```
##### Example (Stdlib)

``` go
```
#### (*Logger) WriterLevel <-0.11.0

``` go
func (logger *Logger) WriterLevel(level Level) *io.PipeWriter
```

WriterLevel returns an io.Writer that can be used to write arbitrary text to the logger at the given log level. Each line written to the writer will be printed in the usual way using formatters and hooks. The writer is part of an io.Pipe and it is the callers responsibility to close the writer when done. This can be used to override the standard library logger easily.

#### type MutexWrap <-0.11.0

``` go
type MutexWrap struct {
	// contains filtered or unexported fields
}
```

#### (*MutexWrap) Disable <-0.11.0

``` go
func (mw *MutexWrap) Disable()
```

#### (*MutexWrap) Lock <-0.11.0

``` go
func (mw *MutexWrap) Lock()
```

#### (*MutexWrap) Unlock <-0.11.0

``` go
func (mw *MutexWrap) Unlock()
```

#### type StdLogger 

``` go
type StdLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}
```

StdLogger is what your logrus-enabled library should take, that way it'll accept a stdlib logger and a logrus logger. There's no standard interface, this is the closest we get, unfortunately.

#### type TextFormatter 

``` go
type TextFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors bool

	// Force disabling colors.
	DisableColors bool

	// Force quoting of all values
	ForceQuote bool

	// DisableQuote disables quoting for all values.
	// DisableQuote will have a lower priority than ForceQuote.
	// If both of them are set to true, quote will be forced on all values.
	DisableQuote bool

	// Override coloring based on CLICOLOR and CLICOLOR_FORCE. - https://bixense.com/clicolors/
	EnvironmentOverrideColors bool

	// Disable timestamp logging. useful when output is redirected to logging
	// system that already adds timestamps.
	DisableTimestamp bool

	// Enable logging the full timestamp when a TTY is attached instead of just
	// the time passed since beginning of execution.
	FullTimestamp bool

	// TimestampFormat to use for display when a full timestamp is printed.
	// The format to use is the same than for time.Format or time.Parse from the standard
	// library.
	// The standard Library already provides a set of predefined format.
	TimestampFormat string

	// The fields are sorted by default for a consistent output. For applications
	// that log extremely frequently and don't use the JSON formatter this may not
	// be desired.
	DisableSorting bool

	// The keys sorting function, when uninitialized it uses sort.Strings.
	SortingFunc func([]string)

	// Disables the truncation of the level text to 4 characters.
	DisableLevelTruncation bool

	// PadLevelText Adds padding the level text so that all the levels output at the same length
	// PadLevelText is a superset of the DisableLevelTruncation option
	PadLevelText bool

	// QuoteEmptyFields will wrap empty fields in quotes if true
	QuoteEmptyFields bool

	// FieldMap allows users to customize the names of keys for default fields.
	// As an example:
	// formatter := &TextFormatter{
	//     FieldMap: FieldMap{
	//         FieldKeyTime:  "@timestamp",
	//         FieldKeyLevel: "@level",
	//         FieldKeyMsg:   "@message"}}
	FieldMap FieldMap

	// CallerPrettyfier can be set by the user to modify the content
	// of the function and file keys in the data when ReportCaller is
	// activated. If any of the returned value is the empty string the
	// corresponding key will be removed from fields.
	CallerPrettyfier func(*runtime.Frame) (function string, file string)
	// contains filtered or unexported fields
}
```

TextFormatter formats logs into text

#### (*TextFormatter) Format 

``` go
func (f *TextFormatter) Format(entry *Entry) ([]byte, error)
```

Format renders a single log entry