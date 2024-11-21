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

​	**InContainer** 如果在容器环境（如 Docker）中，则为 true。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/utils/utils.go#L68)

``` go
var Panic = func(v interface{}) { panic(v) }
```

Panic is the same as the built-in panic.

​	**Panic** 与内置的 `panic` 相同。

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/utils/utils.go#L30)

``` go
var TestEnvs = map[string]string{
	"GODEBUG": "tracebackancestors=100",
}
```

TestEnvs for testing.

​	**TestEnvs** 用于测试。

## 函数

### func AbsolutePaths <- 0.109.0

``` go
func AbsolutePaths(paths []string) []string
```

AbsolutePaths returns absolute paths of files in current working directory.

​	**AbsolutePaths** 返回当前工作目录中文件的绝对路径。

### func All <- 0.52.0

``` go
func All(actions ...func()) func()
```

All runs all actions concurrently, returns the wait function for all actions.

​	**All** 并发运行所有操作，并返回一个等待所有操作完成的函数。

### func CropImage <- 0.88.5

``` go
func CropImage(bin []byte, quality, x, y, width, height int) ([]byte, error)
```

CropImage by the specified box, quality is only for jpeg bin.

​	**CropImage** 按指定区域裁剪图像，`quality` 仅适用于 JPEG 格式。

### func DefaultBackoff <- 0.52.0

``` go
func DefaultBackoff(interval time.Duration) time.Duration
```

DefaultBackoff algorithm: A(n) = A(n-1) * random[1.9, 2.1).

​	**DefaultBackoff** 的算法：A(n) = A(n-1) * random[1.9, 2.1)。

### func Dump <- 0.54.0

``` go
func Dump(list ...interface{}) string
```

Dump values for debugging.

​	**Dump** 调试用的值输出。

### func E 

``` go
func E(args ...interface{}) []interface{}
```

E if the last arg is error, panic it.

​	**E** 如果最后一个参数是 `error`，则抛出 `panic`。

### func EscapeGoString <- 0.62.0

``` go
func EscapeGoString(s string) string
```

EscapeGoString not using encoding like base64 or gzip because of they will make git diff every large for small change.

​	**EscapeGoString** 不使用 Base64 或 Gzip 编码，以避免小改动导致大的 Git 差异。

### func Exec <- 0.52.0

``` go
func Exec(line string, rest ...string) string
```

Exec command.

​	**Exec** 执行命令。

### func ExecLine <- 0.97.13

``` go
func ExecLine(std bool, line string, rest ...string) string
```

ExecLine of command.

​	**ExecLine** 执行命令。

### func FileExists <- 0.52.0

``` go
func FileExists(path string) bool
```

FileExists checks if file exists, only for file, not for dir.

​	**FileExists** 检查文件是否存在，仅适用于文件，不适用于目录。

### func FormatCLIArgs <- 0.106.0

``` go
func FormatCLIArgs(args []string) string
```

FormatCLIArgs into one line string.

​	**FormatCLIArgs** 将命令行参数格式化为单行字符串。

### func Mkdir <- 0.52.0

``` go
func Mkdir(path string) error
```

Mkdir makes dir recursively.

​	**Mkdir** 递归创建目录。

### func MustToJSON <- 0.52.0

``` go
func MustToJSON(data interface{}) string
```

MustToJSON encode data to json string.

​	**MustToJSON** 将数据编码为 JSON 字符串。

### func MustToJSONBytes <- 0.52.0

``` go
func MustToJSONBytes(data interface{}) []byte
```

MustToJSONBytes encode data to json bytes.

​	**MustToJSONBytes** 将数据编码为 JSON 字节流。

### func Noop <- 0.112.9

``` go
func Noop()
```

Noop does nothing.

​	**Noop** 不执行任何操作。

### func OutputFile <- 0.52.0

``` go
func OutputFile(p string, data interface{}) error
```

OutputFile auto creates file if not exists, it will try to detect the data type and auto output binary, string or json.

​	**OutputFile** 如果文件不存在，会自动创建，并根据数据类型输出二进制、字符串或 JSON。

### func Pause <- 0.52.0

``` go
func Pause()
```

Pause the goroutine forever.

​	**Pause** 使当前 goroutine 永久暂停。

### func RandString <- 0.52.0

``` go
func RandString(l int) string
```

RandString generate random string with specified string length.

​	**RandString** 生成指定长度的随机字符串。

### func ReadString <- 0.52.0

``` go
func ReadString(p string) (string, error)
```

ReadString reads file as string.

​	**ReadString** 将文件内容读取为字符串。

### func Retry <- 0.52.0

``` go
func Retry(ctx context.Context, s Sleeper, fn func() (stop bool, err error)) error
```

Retry fn and sleeper until fn returns true or s returns error.

​	**Retry** 按 `fn` 和 `s` 逻辑重试，直到 `fn` 返回 `true` 或 `s` 返回错误。

### func S <- 0.52.0

``` go
func S(tpl string, params ...interface{}) string
```

S Template render, the params is key-value pairs.

​	**S** 模板渲染，`params` 是键值对。

### func Sleep <- 0.52.0

``` go
func Sleep(seconds float64)
```

Sleep the goroutine for specified seconds, such as 2.3 seconds.

​	**Sleep** 暂停当前 goroutine 指定的秒数，例如 2.3 秒。

### func SplicePngVertical <- 0.114.7

``` go
func SplicePngVertical(files []ImgWithBox, format proto.PageCaptureScreenshotFormat, opt *ImgOption) ([]byte, error)
```

SplicePngVertical splice png vertically, if there is only one image, it will return the image directly. Only support png and jpeg format yet, webP is not supported because no suitable processing library was found in golang.

​	**SplicePngVertical** 垂直拼接 PNG 图像，只有一张图时直接返回，支持 PNG 和 JPEG 格式。

### func UseNode <- 0.116.1

``` go
func UseNode(std bool)
```

UseNode installs Node.js and set the bin path to PATH env var.

​	**UseNode** 安装 Node.js 并将二进制路径添加到 PATH 环境变量。

## 类型

### type IdleCounter <- 0.69.0

``` go
type IdleCounter struct {
	// contains filtered or unexported fields
}
```

IdleCounter is similar to sync.WaitGroup but it only resolves if no jobs for specified duration.

​	**IdleCounter** 类似于 `sync.WaitGroup`，但仅在指定时间内无任务时解析。

#### func NewIdleCounter <- 0.69.0

``` go
func NewIdleCounter(d time.Duration) *IdleCounter
```

NewIdleCounter ...

​	**NewIdleCounter** 创建 `IdleCounter`。

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

​	**ImgOption** 图像处理选项。

#### type ImgProcessor <- 0.114.7

``` go
type ImgProcessor interface {
	Encode(img image.Image, opt *ImgOption) ([]byte, error)
	Decode(file io.Reader) (image.Image, error)
}
```

ImgProcessor is the interface for image processing.

​	**ImgProcessor** 图像处理接口。

#### func NewImgProcessor <- 0.114.7

``` go
func NewImgProcessor(format proto.PageCaptureScreenshotFormat) (ImgProcessor, error)
```

NewImgProcessor create a ImgProcessor by the format.

​	**NewImgProcessor** 根据格式创建 `ImgProcessor`。

### type ImgWithBox <- 0.114.7

``` go
type ImgWithBox struct {
	Img []byte
	Box *image.Rectangle
}
```

ImgWithBox is a image with a box, if the box is nil, it means the whole image.

​	**ImgWithBox** 表示带有框的图像，如果框为 `nil`，表示整个图像。

### type Log <- 0.70.0

``` go
type Log func(msg ...interface{})
```

Log type for Println.

​	**Log** 用于 `Println` 的类型。

#### func MultiLogger <- 0.74.0

``` go
func MultiLogger(list ...Logger) Log
```

MultiLogger is similar to https://golang.org/pkg/io/#MultiWriter

​	**MultiLogger** 是 https://golang.org/pkg/io/#MultiWriter 的简化版本。

#### (Log) Println <- 0.70.0

``` go
func (l Log) Println(msg ...interface{})
```

Println interface.

​	**Println** 接口。

### type Logger <- 0.70.0

``` go
type Logger interface {
	// Same as fmt.Printf
    // 与 fmt.Printf 相同
	Println(vs ...interface{})
}
```

Logger interface.

​	**Logger** 接口。

``` go
var LoggerQuiet Logger = Log(func(_ ...interface{}) {})
```

LoggerQuiet does nothing.

​	**LoggerQuiet** 不执行任何操作。

### type MaxSleepCountError <- 0.114.8

``` go
type MaxSleepCountError struct {
	// Max count
    // 最大计数
	Max int
}
```

MaxSleepCountError type.

​	**MaxSleepCountError** 类型。

#### (*MaxSleepCountError) Error <- 0.114.8

``` go
func (e *MaxSleepCountError) Error() string
```

Error interface.

​	实现 **Error** 接口。

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

​	**Sleeper** 使当前 goroutine 暂停一段时间，返回唤醒原因。如果上下文完成则释放资源。

### func BackoffSleeper <- 0.52.0

``` go
func BackoffSleeper(initInterval, maxInterval time.Duration, algorithm func(time.Duration) time.Duration) Sleeper
```

BackoffSleeper returns a sleeper that sleeps in a backoff manner every time get called. The sleep interval of the sleeper will grow from initInterval to maxInterval by the specified algorithm, then use maxInterval as the interval. If maxInterval is not greater than 0, the sleeper will wake immediately. If algorithm is nil, DefaultBackoff will be used.

​	**BackoffSleeper** 返回一个按照回退方式暂停的 `Sleeper`。暂停时间从 `initInterval` 增长到 `maxInterval`，并按照指定算法更新。如果 `maxInterval` 小于等于 0，则立即唤醒。如果算法为 `nil`，将使用默认的 **DefaultBackoff**。

### func CountSleeper <- 0.52.0

``` go
func CountSleeper(max int) Sleeper
```

CountSleeper wakes immediately. When counts to the max returns *ErrMaxSleepCount.

​	**CountSleeper** 立即唤醒。当计数达到最大值时，返回 *ErrMaxSleepCount。

### func EachSleepers <- 0.92.0

``` go
func EachSleepers(list ...Sleeper) Sleeper
```

EachSleepers returns a sleeper wakes up when each sleeper is awake. If a sleeper returns error, it will wake up immediately.

​	**EachSleepers** 返回一个 `Sleeper`，当所有 `Sleeper` 唤醒时会唤醒。如果其中一个 `Sleeper` 返回错误，它会立即唤醒。

### func RaceSleepers <- 0.92.0

``` go
func RaceSleepers(list ...Sleeper) Sleeper
```

RaceSleepers returns a sleeper wakes up when one of the sleepers wakes.

​	**RaceSleepers** 返回一个 `Sleeper`，当任意一个 `Sleeper` 唤醒时会唤醒。
