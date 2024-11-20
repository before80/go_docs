+++
title = "utils"
date = 2024-11-20T18:02:07+08:00
weight = 80
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/utils](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/utils)
>
> 收录该文档时间：`2024-11-20T18:02:07+08:00`
>
> [Version: v0.116.2](https://pkg.go.dev/github.com/go-rod/rod@v0.116.2/lib/utils?tab=versions)

## Overview

Package utils ...

## 常量

This section is empty.

## 变量

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/utils/utils.go#L35)

``` go
var InContainer = FileExists("/.dockerenv") || FileExists("/.containerenv") ||
	os.Getenv("KUBERNETES_SERVICE_HOST") != ""
```

InContainer will be true if is inside container environment, such as docker.

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/utils/utils.go#L68)

``` go
var Panic = func(v interface{}) { panic(v) }
```

Panic is the same as the built-in panic.

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/utils/utils.go#L30)

``` go
var TestEnvs = map[string]string{
	"GODEBUG": "tracebackancestors=100",
}
```

TestEnvs for testing.

## 函数

### func AbsolutePaths <- 0.109.0

``` go
func AbsolutePaths(paths []string) []string
```

AbsolutePaths returns absolute paths of files in current working directory.

### func All <- 0.52.0

``` go
func All(actions ...func()) func()
```

All runs all actions concurrently, returns the wait function for all actions.

### func CropImage <- 0.88.5

``` go
func CropImage(bin []byte, quality, x, y, width, height int) ([]byte, error)
```

CropImage by the specified box, quality is only for jpeg bin.

### func DefaultBackoff <- 0.52.0

``` go
func DefaultBackoff(interval time.Duration) time.Duration
```

DefaultBackoff algorithm: A(n) = A(n-1) * random[1.9, 2.1).

### func Dump <- 0.54.0

``` go
func Dump(list ...interface{}) string
```

Dump values for debugging.

### func E 

``` go
func E(args ...interface{}) []interface{}
```

E if the last arg is error, panic it.

### func EscapeGoString <- 0.62.0

``` go
func EscapeGoString(s string) string
```

EscapeGoString not using encoding like base64 or gzip because of they will make git diff every large for small change.

### func Exec <- 0.52.0

``` go
func Exec(line string, rest ...string) string
```

Exec command.

### func ExecLine <- 0.97.13

``` go
func ExecLine(std bool, line string, rest ...string) string
```

ExecLine of command.

### func FileExists <- 0.52.0

``` go
func FileExists(path string) bool
```

FileExists checks if file exists, only for file, not for dir.

### func FormatCLIArgs <- 0.106.0

``` go
func FormatCLIArgs(args []string) string
```

FormatCLIArgs into one line string.

### func Mkdir <- 0.52.0

``` go
func Mkdir(path string) error
```

Mkdir makes dir recursively.

### func MustToJSON <- 0.52.0

``` go
func MustToJSON(data interface{}) string
```

MustToJSON encode data to json string.

### func MustToJSONBytes <- 0.52.0

``` go
func MustToJSONBytes(data interface{}) []byte
```

MustToJSONBytes encode data to json bytes.

### func Noop <- 0.112.9

``` go
func Noop()
```

Noop does nothing.

### func OutputFile <- 0.52.0

``` go
func OutputFile(p string, data interface{}) error
```

OutputFile auto creates file if not exists, it will try to detect the data type and auto output binary, string or json.

### func Pause <- 0.52.0

``` go
func Pause()
```

Pause the goroutine forever.

### func RandString <- 0.52.0

``` go
func RandString(l int) string
```

RandString generate random string with specified string length.

### func ReadString <- 0.52.0

``` go
func ReadString(p string) (string, error)
```

ReadString reads file as string.

### func Retry <- 0.52.0

``` go
func Retry(ctx context.Context, s Sleeper, fn func() (stop bool, err error)) error
```

Retry fn and sleeper until fn returns true or s returns error.

### func S <- 0.52.0

``` go
func S(tpl string, params ...interface{}) string
```

S Template render, the params is key-value pairs.

### func Sleep <- 0.52.0

``` go
func Sleep(seconds float64)
```

Sleep the goroutine for specified seconds, such as 2.3 seconds.

### func SplicePngVertical <- 0.114.7

``` go
func SplicePngVertical(files []ImgWithBox, format proto.PageCaptureScreenshotFormat, opt *ImgOption) ([]byte, error)
```

SplicePngVertical splice png vertically, if there is only one image, it will return the image directly. Only support png and jpeg format yet, webP is not supported because no suitable processing library was found in golang.

### func UseNode <- 0.116.1

``` go
func UseNode(std bool)
```

UseNode installs Node.js and set the bin path to PATH env var.

## 类型

### type IdleCounter <- 0.69.0

``` go
type IdleCounter struct {
	// contains filtered or unexported fields
}
```

IdleCounter is similar to sync.WaitGroup but it only resolves if no jobs for specified duration.

#### func NewIdleCounter <- 0.69.0

``` go
func NewIdleCounter(d time.Duration) *IdleCounter
```

NewIdleCounter ...

#### (*IdleCounter) Add <- 0.69.0

``` go
func (de *IdleCounter) Add()
```

Add ...

#### (*IdleCounter) Done <- 0.69.0

``` go
func (de *IdleCounter) Done()
```

Done ...

#### (*IdleCounter) Wait <- 0.69.0

``` go
func (de *IdleCounter) Wait(ctx context.Context)
```

Wait ...

#### type ImgOption <- 0.114.7

``` go
type ImgOption struct {
	Quality int
}
```

ImgOption is the option for image processing.

#### type ImgProcessor <- 0.114.7

``` go
type ImgProcessor interface {
	Encode(img image.Image, opt *ImgOption) ([]byte, error)
	Decode(file io.Reader) (image.Image, error)
}
```

ImgProcessor is the interface for image processing.

#### func NewImgProcessor <- 0.114.7

``` go
func NewImgProcessor(format proto.PageCaptureScreenshotFormat) (ImgProcessor, error)
```

NewImgProcessor create a ImgProcessor by the format.

### type ImgWithBox <- 0.114.7

``` go
type ImgWithBox struct {
	Img []byte
	Box *image.Rectangle
}
```

ImgWithBox is a image with a box, if the box is nil, it means the whole image.

### type Log <- 0.70.0

``` go
type Log func(msg ...interface{})
```

Log type for Println.

#### func MultiLogger <- 0.74.0

``` go
func MultiLogger(list ...Logger) Log
```

MultiLogger is similar to https://golang.org/pkg/io/#MultiWriter

#### (Log) Println <- 0.70.0

``` go
func (l Log) Println(msg ...interface{})
```

Println interface.

### type Logger <- 0.70.0

``` go
type Logger interface {
	// Same as fmt.Printf
	Println(vs ...interface{})
}
```

Logger interface.

``` go
var LoggerQuiet Logger = Log(func(_ ...interface{}) {})
```

LoggerQuiet does nothing.

### type MaxSleepCountError <- 0.114.8

``` go
type MaxSleepCountError struct {
	// Max count
	Max int
}
```

MaxSleepCountError type.

#### (*MaxSleepCountError) Error <- 0.114.8

``` go
func (e *MaxSleepCountError) Error() string
```

Error interface.

#### (*MaxSleepCountError) Is <- 0.114.8

``` go
func (e *MaxSleepCountError) Is(err error) bool
```

Is interface.

### type Sleeper <- 0.52.0

``` go
type Sleeper func(context.Context) error
```

Sleeper sleeps the current goroutine for sometime, returns the reason to wake, if ctx is done release resource.

### func BackoffSleeper <- 0.52.0

``` go
func BackoffSleeper(initInterval, maxInterval time.Duration, algorithm func(time.Duration) time.Duration) Sleeper
```

BackoffSleeper returns a sleeper that sleeps in a backoff manner every time get called. The sleep interval of the sleeper will grow from initInterval to maxInterval by the specified algorithm, then use maxInterval as the interval. If maxInterval is not greater than 0, the sleeper will wake immediately. If algorithm is nil, DefaultBackoff will be used.

### func CountSleeper <- 0.52.0

``` go
func CountSleeper(max int) Sleeper
```

CountSleeper wakes immediately. When counts to the max returns *ErrMaxSleepCount.

### func EachSleepers <- 0.92.0

``` go
func EachSleepers(list ...Sleeper) Sleeper
```

EachSleepers returns a sleeper wakes up when each sleeper is awake. If a sleeper returns error, it will wake up immediately.

### func RaceSleepers <- 0.92.0

``` go
func RaceSleepers(list ...Sleeper) Sleeper
```

RaceSleepers returns a sleeper wakes up when one of the sleepers wakes.
