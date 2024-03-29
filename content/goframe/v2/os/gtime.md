+++
title = "gtime"
date = 2024-03-21T17:57:36+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gtime](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gtime)

Package gtime provides functionality for measuring and displaying time.

​	软件包 gtime 提供测量和显示时间的功能。

This package should keep much less dependencies with other packages.

​	此包应保留与其他包的依赖性要少得多。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gtime/gtime.go#L27)

```go
const (
	D  = 24 * time.Hour
	H  = time.Hour
	M  = time.Minute
	S  = time.Second
	MS = time.Millisecond
	US = time.Microsecond
	NS = time.Nanosecond
)
```

## 变量

This section is empty.

## 函数

#### func Date

```go
func Date() string
```

Date returns current date in string like “2006-01-02”.

​	Date 以字符串形式返回当前日期，例如“2006-01-02”。

##### Example

``` go
```

#### func Datetime

```go
func Datetime() string
```

Datetime returns current datetime in string like “2006-01-02 15:04:05”.

​	Datetime 以字符串形式返回当前日期时间，例如“2006-01-02 15：04：05”。

##### Example

``` go
```

#### func FuncCost

```go
func FuncCost(f func()) time.Duration
```

FuncCost calculates the cost time of function `f` in nanoseconds.

​	FuncCost 以纳秒为单位计算函数 `f` 的成本时间。

#### func ISO8601

```go
func ISO8601() string
```

ISO8601 returns current datetime in ISO8601 format like “2006-01-02T15:04:05-07:00”.

​	ISO8601以ISO8601格式返回当前日期时间，例如“2006-01-02T15：04：05-07：00”。

##### Example

``` go
```

#### func ParseDuration

```go
func ParseDuration(s string) (duration time.Duration, err error)
```

ParseDuration parses a duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as “300ms”, “-1.5h”, “1d” or “2h45m”. Valid time units are “ns”, “us” (or “µs”), “ms”, “s”, “m”, “h”, “d”.

​	ParseDuration 分析持续时间字符串。持续时间字符串是十进制数的可能有符号序列，每个十进制数都带有可选分数和单位后缀，例如“300ms”、“-1.5h”、“1d”或“2h45m”。有效时间单位为“ns”、“us”（或“μs”）、“ms”、“s”、“m”、“h”、“d”。

Very note that it supports unit “d” more than function time.ParseDuration.

​	非常注意，它支持单位“d”多于功能时间。ParseDuration。

##### Example

``` go
```

#### func RFC822

```go
func RFC822() string
```

RFC822 returns current datetime in RFC822 format like “Mon, 02 Jan 06 15:04 MST”.

​	RFC822 以 RFC822 格式返回当前日期时间，例如“Mon， 02 Jan 06 15：04 MST”。

##### Example

``` go
```

#### func SetTimeZone

```go
func SetTimeZone(zone string) (err error)
```

SetTimeZone sets the time zone for current whole process. The parameter `zone` is an area string specifying corresponding time zone, eg: Asia/Shanghai.

​	SetTimeZone 设置当前整个进程的时区。该参数 `zone` 是指定相应时区的区域字符串，例如：Asia/Shanghai。

PLEASE VERY NOTE THAT: 1. This should be called before package “time” import. 2. This function should be called once. 3. Please refer to issue: https://github.com/golang/go/issues/34814

​	请注意： 1.这应该在包“时间”导入之前调用。2. 此函数应调用一次。3. 请参考问题：https://github.com/golang/go/issues/34814

##### Example

``` go
```

#### func Timestamp

```go
func Timestamp() int64
```

Timestamp retrieves and returns the timestamp in seconds.

​	Timestamp 检索并返回时间戳（以秒为单位）。

##### Example

``` go
```

#### func TimestampMicro

```go
func TimestampMicro() int64
```

TimestampMicro retrieves and returns the timestamp in microseconds.

​	TimestampMicro 检索并返回时间戳（以微秒为单位）。

##### Example

``` go
```

#### func TimestampMicroStr

```go
func TimestampMicroStr() string
```

TimestampMicroStr is a convenience method which retrieves and returns the timestamp in microseconds as string.

​	TimestampMicroStr 是一种方便的方法，它以字符串的形式检索和返回以微秒为单位的时间戳。

#### func TimestampMilli

```go
func TimestampMilli() int64
```

TimestampMilli retrieves and returns the timestamp in milliseconds.

​	TimestampMilli 检索并返回时间戳（以毫秒为单位）。

##### Example

``` go
```

#### func TimestampMilliStr

```go
func TimestampMilliStr() string
```

TimestampMilliStr is a convenience method which retrieves and returns the timestamp in milliseconds as string.

​	TimestampMilliStr 是一种方便的方法，它检索时间戳并将其以毫秒为单位作为字符串返回。

#### func TimestampNano

```go
func TimestampNano() int64
```

TimestampNano retrieves and returns the timestamp in nanoseconds.

​	TimestampNano 检索并返回时间戳（以纳秒为单位）。

##### Example

``` go
```

#### func TimestampNanoStr

```go
func TimestampNanoStr() string
```

TimestampNanoStr is a convenience method which retrieves and returns the timestamp in nanoseconds as string.

​	TimestampNanoStr 是一种方便的方法，它检索时间戳并将其以纳秒为单位作为字符串返回。

#### func TimestampStr

```go
func TimestampStr() string
```

TimestampStr is a convenience method which retrieves and returns the timestamp in seconds as string.

​	TimestampStr 是一种方便的方法，它检索时间戳并将其以字符串形式返回（以秒为单位）。

##### Example

``` go
```

## 类型

### type Time

```go
type Time struct {
	// contains filtered or unexported fields
}
```

Time is a wrapper for time.Time for additional features.

​	时间是时间的包装。是时候使用其他功能了。

#### func ConvertZone

```go
func ConvertZone(strTime string, toZone string, fromZone ...string) (*Time, error)
```

ConvertZone converts time in string `strTime` from `fromZone` to `toZone`. The parameter `fromZone` is unnecessary, it is current time zone in default.

​	ConvertZone 将字符串 `strTime` 中的时间从 `fromZone` 转换为 `toZone` 。该参数 `fromZone` 是不必要的，默认为当前时区。

##### Example

``` go
```

#### func New

```go
func New(param ...interface{}) *Time
```

New creates and returns a Time object with given parameter. The optional parameter is the time object which can be type of: time.Time/*time.Time, string or integer. Example: New(“2024-10-29”) New(1390876568) New(t) // The t is type of time.Time.

​	New 创建并返回一个具有给定参数的 Time 对象。可选参数是 time 对象，其类型可以是：time。时间/*时间。时间、字符串或整数。示例： New（“2024-10-29”） New（1390876568） New（t） // t 是时间类型。时间。

#### func NewFromStr

```go
func NewFromStr(str string) *Time
```

NewFromStr creates and returns a Time object with given string. Note that it returns nil if there’s error occurs.

​	NewFromStr 创建并返回具有给定字符串的 Time 对象。请注意，如果发生错误，它将返回 nil。

##### Example

``` go
```

#### func NewFromStrFormat

```go
func NewFromStrFormat(str string, format string) *Time
```

NewFromStrFormat creates and returns a Time object with given string and custom format like: Y-m-d H:i:s. Note that it returns nil if there’s error occurs.

​	NewFromStrFormat 创建并返回一个具有给定字符串和自定义格式的 Time 对象，例如：Y-m-d H：i：s。请注意，如果发生错误，它将返回 nil。

##### Example

``` go
```

#### func NewFromStrLayout

```go
func NewFromStrLayout(str string, layout string) *Time
```

NewFromStrLayout creates and returns a Time object with given string and stdlib layout like: 2006-01-02 15:04:05. Note that it returns nil if there’s error occurs.

​	NewFromStrLayout 创建并返回一个具有给定字符串和 stdlib 布局的 Time 对象，如：2006-01-02 15：04：05。请注意，如果发生错误，它将返回 nil。

##### Example

``` go
```

#### func NewFromTime

```go
func NewFromTime(t time.Time) *Time
```

NewFromTime creates and returns a Time object with given time.Time object.

​	NewFromTime 创建并返回具有给定时间的 Time 对象。时间对象。

##### Example

``` go
```

#### func NewFromTimeStamp

```go
func NewFromTimeStamp(timestamp int64) *Time
```

NewFromTimeStamp creates and returns a Time object with given timestamp, which can be in seconds to nanoseconds. Eg: 1600443866 and 1600443866199266000 are both considered as valid timestamp number.

​	NewFromTimeStamp 创建并返回具有给定时间戳的 Time 对象，该时间戳可以是秒到纳秒。例如：1600443866 和 1600443866199266000 都被视为有效的时间戳编号。

##### Example

``` go
```

#### func Now

```go
func Now() *Time
```

Now creates and returns a time object of now.

​	Now 创建并返回 now 的时间对象。

##### Example

``` go
```

#### func ParseTimeFromContent

```go
func ParseTimeFromContent(content string, format ...string) *Time
```

ParseTimeFromContent retrieves time information for content string, it then parses and returns it as *Time object. It returns the first time information if there are more than one time string in the content. It only retrieves and parses the time information with given first matched `format` if it’s passed.

​	ParseTimeFromContent 检索内容字符串的时间信息，然后将其解析并作为 *Time 对象返回。如果内容中有多个时间字符串，则返回首次时间信息。它仅检索和解析给定第一个匹配 `format` 的时间信息（如果它已通过）。

#### func StrToTime

```go
func StrToTime(str string, format ...string) (*Time, error)
```

StrToTime converts string to *Time object. It also supports timestamp string. The parameter `format` is unnecessary, which specifies the format for converting like “Y-m-d H:i:s”. If `format` is given, it acts as same as function StrToTimeFormat. If `format` is not given, it converts string as a “standard” datetime string. Note that, it fails and returns error if there’s no date string in `str`.

​	StrToTime 将字符串转换为 *Time 对象。它还支持时间戳字符串。该参数 `format` 是不必要的，它指定了转换的格式，如“Y-m-d H：i：s”。如果 `format` 给定，则它的作用与函数 StrToTimeFormat 相同。如果 `format` 未给出，则将字符串转换为“标准”日期时间字符串。请注意，如果 中没有 `str` 日期字符串，则会失败并返回错误。

##### Example

``` go
```

#### func StrToTimeFormat

```go
func StrToTimeFormat(str string, format string) (*Time, error)
```

StrToTimeFormat parses string `str` to *Time object with given format `format`. The parameter `format` is like “Y-m-d H:i:s”.

​	StrToTimeFormat 使用给定格式 `format` 将字符串 `str` 解析为 *Time 对象。参数 `format` 类似于“Y-m-d H：i：s”。

##### Example

``` go
```

#### func StrToTimeLayout

```go
func StrToTimeLayout(str string, layout string) (*Time, error)
```

StrToTimeLayout parses string `str` to *Time object with given format `layout`. The parameter `layout` is in stdlib format like “2006-01-02 15:04:05”.

​	StrToTimeLayout 使用给定格式 `layout` 将字符串 `str` 解析为 *Time 对象。该参数 `layout` 采用 stdlib 格式，如“2006-01-02 15：04：05”。

##### Example

``` go
```

#### (*Time) Add

```go
func (t *Time) Add(d time.Duration) *Time
```

Add adds the duration to current time.

​	“添加”将持续时间添加到当前时间。

##### Example

``` go
```

#### (*Time) AddDate

```go
func (t *Time) AddDate(years int, months int, days int) *Time
```

AddDate adds year, month and day to the time.

​	AddDate 将年、月和日添加到时间中。

##### Example

``` go
```

#### (*Time) AddStr

```go
func (t *Time) AddStr(duration string) (*Time, error)
```

AddStr parses the given duration as string and adds it to current time.

​	AddStr 将给定的持续时间解析为字符串，并将其添加到当前时间。

##### Example

``` go
```

#### (*Time) After

```go
func (t *Time) After(u *Time) bool
```

After reports whether the time instant t is after u.

​	后报告时间瞬间t是否在u之后。

##### Example

``` go
```

#### (*Time) Before

```go
func (t *Time) Before(u *Time) bool
```

Before reports whether the time instant t is before u.

​	之前报告时间时刻 t 是否在 u 之前。

##### Example

``` go
```

#### (*Time) Clone

```go
func (t *Time) Clone() *Time
```

Clone returns a new Time object which is a clone of current time object.

​	克隆返回一个新的 Time 对象，该对象是当前时间对象的克隆。

#### (*Time) DayOfYear

```go
func (t *Time) DayOfYear() int
```

DayOfYear checks and returns the position of the day for the year.

​	DayOfYear 检查并返回当年的当天位置。

##### Example

``` go
```

#### (*Time) DaysInMonth

```go
func (t *Time) DaysInMonth() int
```

DaysInMonth returns the day count of current month.

​	DaysInMonth 返回当前月份的天数。

##### Example

``` go
```

#### (*Time) DeepCopy

```go
func (t *Time) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

​	DeepCopy实现了当前类型的深度拷贝接口。

#### (*Time) EndOfDay

```go
func (t *Time) EndOfDay(withNanoPrecision ...bool) *Time
```

EndOfDay clones and returns a new time which is the end of day the and its time is set to 23:59:59.

​	EndOfDay 克隆并返回一个新时间，即一天的结束时间，其时间设置为 23：59：59。

##### Example

``` go
```

#### (*Time) EndOfHalf

```go
func (t *Time) EndOfHalf(withNanoPrecision ...bool) *Time
```

EndOfHalf clones and returns a new time which is the end of the half year and its time is set to 23:59:59.

​	EndOfHalf 克隆并返回一个新时间，即半年的结束时间，其时间设置为 23：59：59。

##### Example

``` go
```

#### (*Time) EndOfHour

```go
func (t *Time) EndOfHour(withNanoPrecision ...bool) *Time
```

EndOfHour clones and returns a new time of which the minutes and seconds are both set to 59.

​	EndOfHour 克隆并返回一个新时间，其中分钟和秒均设置为 59。

##### Example

``` go
```

#### (*Time) EndOfMinute

```go
func (t *Time) EndOfMinute(withNanoPrecision ...bool) *Time
```

EndOfMinute clones and returns a new time of which the seconds is set to 59.

​	EndOfMinute 克隆并返回秒数设置为 59 的新时间。

##### Example

``` go
```

#### (*Time) EndOfMonth

```go
func (t *Time) EndOfMonth(withNanoPrecision ...bool) *Time
```

EndOfMonth clones and returns a new time which is the end of the month and its time is set to 23:59:59.

​	EndOfMonth 克隆并返回一个新时间，即月底，其时间设置为 23：59：59。

##### Example

``` go
```

#### (*Time) EndOfQuarter

```go
func (t *Time) EndOfQuarter(withNanoPrecision ...bool) *Time
```

EndOfQuarter clones and returns a new time which is end of the quarter and its time is set to 23:59:59.

​	EndOfQuarter 克隆并返回一个新时间，即季度末，其时间设置为 23：59：59。

##### Example

``` go
```

#### (*Time) EndOfWeek

```go
func (t *Time) EndOfWeek(withNanoPrecision ...bool) *Time
```

EndOfWeek clones and returns a new time which is the end of week and its time is set to 23:59:59.

​	EndOfWeek 克隆并返回一个新时间，即周末，其时间设置为 23：59：59。

##### Example

``` go
```

#### (*Time) EndOfYear

```go
func (t *Time) EndOfYear(withNanoPrecision ...bool) *Time
```

EndOfYear clones and returns a new time which is the end of the year and its time is set to 23:59:59.

​	EndOfYear 克隆并返回一个新时间，即年末，其时间设置为 23：59：59。

##### Example

``` go
```

#### (*Time) Equal

```go
func (t *Time) Equal(u *Time) bool
```

Equal reports whether t and u represent the same time instant. Two times can be equal even if they are in different locations. For example, 6:00 +0200 CEST and 4:00 UTC are Equal. See the documentation on the Time type for the pitfalls of using == with Time values; most code should use Equal instead.

​	相等报告 t 和 u 是否表示相同的时间时刻。即使它们位于不同的位置，两次也可以相等。例如，6：00 +0200 CEST 和 4：00 UTC 相等。请参阅有关 Time 类型的文档，了解将 == 与 Time 值一起使用的陷阱;大多数代码应改用 Equal。

##### Example

``` go
```

#### (*Time) Format

```go
func (t *Time) Format(format string) string
```

Format formats and returns the formatted result with custom `format`. Refer method Layout, if you want to follow stdlib layout.

​	格式化格式并使用自定义 `format` .如果要遵循 stdlib 布局，请参阅方法 Layout。

##### Example

``` go
```

#### (*Time) FormatNew

```go
func (t *Time) FormatNew(format string) *Time
```

FormatNew formats and returns a new Time object with given custom `format`.

​	FormatNew 格式化并返回具有给定自定义 `format` .

##### Example

``` go
```

#### (*Time) FormatTo

```go
func (t *Time) FormatTo(format string) *Time
```

FormatTo formats `t` with given custom `format`.

​	 `t` 具有给定自定义 `format` .

##### Example

``` go
```

#### (*Time) ISO8601

```go
func (t *Time) ISO8601() string
```

ISO8601 formats the time as ISO8601 and returns it as string.

​	ISO8601 将时间格式化为 ISO8601 并将其作为字符串返回。

#### (*Time) IsLeapYear

```go
func (t *Time) IsLeapYear() bool
```

IsLeapYear checks whether the time is leap year.

​	IsLeapYear 检查时间是否为闰年。

##### Example

``` go
```

#### (*Time) IsZero

```go
func (t *Time) IsZero() bool
```

IsZero reports whether t represents the zero time instant, January 1, year 1, 00:00:00 UTC.

​	IsZero 报告 t 是否表示零时间时刻，即 1 年 1 月 1 日 00：00：00 UTC。

##### Example

``` go
```

#### (*Time) Layout

```go
func (t *Time) Layout(layout string) string
```

Layout formats the time with stdlib layout and returns the formatted result.

​	Layout 使用 stdlib layout 格式化时间并返回格式化结果。

##### Example

``` go
```

#### (*Time) LayoutNew

```go
func (t *Time) LayoutNew(layout string) *Time
```

LayoutNew formats the time with stdlib layout and returns the new Time object.

​	LayoutNew 使用 stdlib 布局格式化时间，并返回新的 Time 对象。

##### Example

``` go
```

#### (*Time) LayoutTo

```go
func (t *Time) LayoutTo(layout string) *Time
```

LayoutTo formats `t` with stdlib layout.

​	 `t` 使用 stdlib 布局的 LayoutTo 格式。

##### Example

``` go
```

#### (*Time) Local

```go
func (t *Time) Local() *Time
```

Local converts the time to local timezone.

​	Local 将时间转换为本地时区。

#### (Time) MarshalJSON

```go
func (t Time) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. Note that, DO NOT use `(t *Time) MarshalJSON() ([]byte, error)` as it looses interface implement of `MarshalJSON` for struct of Time.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。请注意，不要使用 `(t *Time) MarshalJSON() ([]byte, error)` ，因为它会丢失 `MarshalJSON` for struct of Time 的接口实现。

##### Example

``` go
```

#### (*Time) Microsecond

```go
func (t *Time) Microsecond() int
```

Microsecond returns the microsecond offset within the second specified by t, in the range [0, 999999].

​	微秒返回 t 指定的秒内的微秒偏移量，范围为 [0， 999999]。

#### (*Time) Millisecond

```go
func (t *Time) Millisecond() int
```

Millisecond returns the millisecond offset within the second specified by t, in the range [0, 999].

​	毫秒返回 t 指定的秒内的毫秒偏移量，范围为 [0， 999]。

#### (*Time) Month

```go
func (t *Time) Month() int
```

Month returns the month of the year specified by t.

​	month 返回 t 指定的年份中的月份。

##### Example

``` go
```

#### (*Time) Nanosecond

```go
func (t *Time) Nanosecond() int
```

Nanosecond returns the nanosecond offset within the second specified by t, in the range [0, 999999999].

​	纳秒返回 t 指定的秒内的纳秒偏移量，范围为 [0， 999999999]。

#### (*Time) NoValidation

```go
func (t *Time) NoValidation()
```

NoValidation marks this struct object will not be validated by package gvalid.

​	NoValidation 标记此结构对象不会被包 gvalid 验证。

#### (*Time) RFC822

```go
func (t *Time) RFC822() string
```

RFC822 formats the time as RFC822 and returns it as string.

​	RFC822 将时间格式化为 RFC822，并将其作为字符串返回。

#### (*Time) Round

```go
func (t *Time) Round(d time.Duration) *Time
```

Round returns the result of rounding t to the nearest multiple of d (since the zero time). The rounding behavior for halfway values is to round up. If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.

​	Round 返回将 t 舍入到 d 的最接近倍数的结果（自零时间以来）。中途值的舍入行为是舍入。如果 d <= 0，则 Round 返回 t 去除了任何单调时钟读数，但其他方面保持不变。

Round operates on the time as an absolute duration since the zero time; it does not operate on the presentation form of the time. Thus, Round(Hour) may return a time with a non-zero minute, depending on the time’s Location.

​	Round 在时间上作为自零时间以来的绝对持续时间运行;它不对当时的演示形式进行操作。因此，Round（Hour） 可能会返回一个非零分钟的时间，具体取决于时间的位置。

##### Example

``` go
```

#### (*Time) Scan

```go
func (t *Time) Scan(value interface{}) error
```

Scan implements interface used by Scan in package database/sql for Scanning value from database to local golang variable.

​	Scan 实现了 Scan 在包 database/sql 中使用的接口，用于将数据库的值扫描到本地 golang 变量。

#### (*Time) Second

```go
func (t *Time) Second() int
```

Second returns the second offset within the minute specified by t, in the range [0, 59].

​	Second 返回 t 指定的分钟内的第二个偏移量，范围为 [0， 59]。

##### Example

``` go
```

#### (*Time) StartOfDay

```go
func (t *Time) StartOfDay() *Time
```

StartOfDay clones and returns a new time which is the start of day, its time is set to 00:00:00.

​	StartOfDay 克隆并返回一个新时间，即一天的开始，其时间设置为 00：00：00。

##### Example

``` go
```

#### (*Time) StartOfHalf

```go
func (t *Time) StartOfHalf() *Time
```

StartOfHalf clones and returns a new time which is the first day of the half year and its time is set to 00:00:00.

​	StartOfHalf 克隆并返回一个新时间，即半年的第一天，其时间设置为 00：00：00。

##### Example

``` go
```

#### (*Time) StartOfHour

```go
func (t *Time) StartOfHour() *Time
```

StartOfHour clones and returns a new time of which the hour, minutes and seconds are set to 0.

​	StartOfHour 克隆并返回一个新时间，其中的小时、分钟和秒设置为 0。

##### Example

``` go
```

#### (*Time) StartOfMinute

```go
func (t *Time) StartOfMinute() *Time
```

StartOfMinute clones and returns a new time of which the seconds is set to 0.

​	StartOfMinute 克隆并返回秒数设置为 0 的新时间。

##### Example

``` go
```

#### (*Time) StartOfMonth

```go
func (t *Time) StartOfMonth() *Time
```

StartOfMonth clones and returns a new time which is the first day of the month and its is set to 00:00:00

​	StartOfMonth 克隆并返回一个新时间，即该月的第一天，其设置为 00：00：00

#### (*Time) StartOfQuarter

```go
func (t *Time) StartOfQuarter() *Time
```

StartOfQuarter clones and returns a new time which is the first day of the quarter and its time is set to 00:00:00.

​	StartOfQuarter 克隆并返回一个新时间，即该季度的第一天，其时间设置为 00：00：00。

##### Example

``` go
```

#### (*Time) StartOfWeek

```go
func (t *Time) StartOfWeek() *Time
```

StartOfWeek clones and returns a new time which is the first day of week and its time is set to 00:00:00.

​	StartOfWeek 克隆并返回一个新时间，即一周的第一天，其时间设置为 00：00：00。

##### Example

``` go
```

#### (*Time) StartOfYear

```go
func (t *Time) StartOfYear() *Time
```

StartOfYear clones and returns a new time which is the first day of the year and its time is set to 00:00:00.

​	StartOfYear 克隆并返回一个新时间，即一年中的第一天，其时间设置为 00：00：00。

##### Example

``` go
```

#### (*Time) String

```go
func (t *Time) String() string
```

String returns current time object as string.

​	String 以字符串形式返回当前时间对象。

##### Example

``` go
```

#### (*Time) Sub

```go
func (t *Time) Sub(u *Time) time.Duration
```

Sub returns the duration t-u. If the result exceeds the maximum (or minimum) value that can be stored in a Duration, the maximum (or minimum) duration will be returned. To compute t-d for a duration d, use t.Add(-d).

​	Sub 返回持续时间 t-u。如果结果超过持续时间中可以存储的最大（或最小）值，则将返回最大（或最小）持续时间。要计算持续时间为 d 的 t-d，请使用 t.Add（-d）。

##### Example

``` go
```

#### (*Time) Timestamp

```go
func (t *Time) Timestamp() int64
```

Timestamp returns the timestamp in seconds.

​	Timestamp 返回时间戳（以秒为单位）。

##### Example

``` go
```

#### (*Time) TimestampMicro

```go
func (t *Time) TimestampMicro() int64
```

TimestampMicro returns the timestamp in microseconds.

​	TimestampMicro 返回时间戳（以微秒为单位）。

##### Example

``` go
```

#### (*Time) TimestampMicroStr

```go
func (t *Time) TimestampMicroStr() string
```

TimestampMicroStr is a convenience method which retrieves and returns the timestamp in microseconds as string.

​	TimestampMicroStr 是一种方便的方法，它以字符串的形式检索和返回以微秒为单位的时间戳。

#### (*Time) TimestampMilli

```go
func (t *Time) TimestampMilli() int64
```

TimestampMilli returns the timestamp in milliseconds.

​	TimestampMilli 以毫秒为单位返回时间戳。

##### Example

``` go
```

#### (*Time) TimestampMilliStr

```go
func (t *Time) TimestampMilliStr() string
```

TimestampMilliStr is a convenience method which retrieves and returns the timestamp in milliseconds as string.

​	TimestampMilliStr 是一种方便的方法，它检索时间戳并将其以毫秒为单位作为字符串返回。

#### (*Time) TimestampNano

```go
func (t *Time) TimestampNano() int64
```

TimestampNano returns the timestamp in nanoseconds.

​	TimestampNano 以纳秒为单位返回时间戳。

##### Example

``` go
```

#### (*Time) TimestampNanoStr

```go
func (t *Time) TimestampNanoStr() string
```

TimestampNanoStr is a convenience method which retrieves and returns the timestamp in nanoseconds as string.

​	TimestampNanoStr 是一种方便的方法，它检索时间戳并将其以纳秒为单位作为字符串返回。

#### (*Time) TimestampStr

```go
func (t *Time) TimestampStr() string
```

TimestampStr is a convenience method which retrieves and returns the timestamp in seconds as string.

​	TimestampStr 是一种方便的方法，它检索时间戳并将其以字符串形式返回（以秒为单位）。

##### Example

``` go
```

#### (*Time) ToLocation

```go
func (t *Time) ToLocation(location *time.Location) *Time
```

ToLocation converts current time to specified location.

​	ToLocation 将当前时间转换为指定位置。

#### (*Time) ToZone

```go
func (t *Time) ToZone(zone string) (*Time, error)
```

ToZone converts current time to specified zone like: Asia/Shanghai.

​	ToZone 将当前时间转换为指定区域，例如：亚洲/上海。

##### Example

``` go
```

#### (*Time) Truncate

```go
func (t *Time) Truncate(d time.Duration) *Time
```

Truncate returns the result of rounding t down to a multiple of d (since the zero time). If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.

​	截断返回将 t 向下舍入为 d 的倍数的结果（自零时间以来）。如果 d <= 0，则截断返回 t 去除了任何单调时钟读数，但其他方面保持不变。

Truncate operates on the time as an absolute duration since the zero time; it does not operate on the presentation form of the time. Thus, Truncate(Hour) may return a time with a non-zero minute, depending on the time’s Location.

​	截断在时间上作为自零时间以来的绝对持续时间;它不对当时的演示形式进行操作。因此，Truncate（Hour） 可能会返回非零分钟的时间，具体取决于时间的位置。

##### Example

``` go
```

#### (*Time) UTC

```go
func (t *Time) UTC() *Time
```

UTC converts current time to UTC timezone.

​	UTC 将当前时间转换为 UTC 时区。

#### (*Time) UnmarshalJSON

```go
func (t *Time) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

​	UnmarshalJSON 实现 json 的接口 UnmarshalJSON。元帅。

##### Example

``` go
```

#### (*Time) UnmarshalText

```go
func (t *Time) UnmarshalText(data []byte) error
```

UnmarshalText implements the encoding.TextUnmarshaler interface. Note that it overwrites the same implementer of `time.Time`.

​	UnmarshalText 实现编码。TextUnmarshaler 接口。请注意，它覆盖了 的 `time.Time` 同一实现器。

#### (*Time) Value

```go
func (t *Time) Value() (driver.Value, error)
```

Value is the interface providing the Value method for package database/sql/driver for retrieving value from golang variable to database.

​	Value 是为 package database/sql/driver 提供 Value 方法的接口，用于从 golang 变量检索到数据库的值。

#### (*Time) WeeksOfYear

```go
func (t *Time) WeeksOfYear() int
```

WeeksOfYear returns the point of current week for the year.

​	WeeksOfYear 返回当年的当周点。

Example WeekOfYear

​	示例 WeekOfYear

```go
package main

import (
	"fmt"

	"github.com/gogf/gf/v2/os/gtime"
)

func main() {
	gt1 := gtime.New("2018-01-08 08:08:08")

	fmt.Println(gt1.WeeksOfYear())

}

Output:

2
```
