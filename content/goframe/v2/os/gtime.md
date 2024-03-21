+++
title = "gtime"
date = 2024-03-21T17:57:36+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gtime

Package gtime provides functionality for measuring and displaying time.

This package should keep much less dependencies with other packages.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gtime/gtime.go#L27)

``` go
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

### Variables 

This section is empty.

### Functions 

##### func Date 

``` go
func Date() string
```

Date returns current date in string like "2006-01-02".

##### Example

``` go
```
##### func Datetime 

``` go
func Datetime() string
```

Datetime returns current datetime in string like "2006-01-02 15:04:05".

##### Example

``` go
```
##### func FuncCost 

``` go
func FuncCost(f func()) time.Duration
```

FuncCost calculates the cost time of function `f` in nanoseconds.

##### func ISO8601 

``` go
func ISO8601() string
```

ISO8601 returns current datetime in ISO8601 format like "2006-01-02T15:04:05-07:00".

##### Example

``` go
```
##### func ParseDuration 

``` go
func ParseDuration(s string) (duration time.Duration, err error)
```

ParseDuration parses a duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h", "1d" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h", "d".

Very note that it supports unit "d" more than function time.ParseDuration.

##### Example

``` go
```
##### func RFC822 

``` go
func RFC822() string
```

RFC822 returns current datetime in RFC822 format like "Mon, 02 Jan 06 15:04 MST".

##### Example

``` go
```
##### func SetTimeZone 

``` go
func SetTimeZone(zone string) (err error)
```

SetTimeZone sets the time zone for current whole process. The parameter `zone` is an area string specifying corresponding time zone, eg: Asia/Shanghai.

PLEASE VERY NOTE THAT: 1. This should be called before package "time" import. 2. This function should be called once. 3. Please refer to issue: https://github.com/golang/go/issues/34814

##### Example

``` go
```
##### func Timestamp 

``` go
func Timestamp() int64
```

Timestamp retrieves and returns the timestamp in seconds.

##### Example

``` go
```
##### func TimestampMicro 

``` go
func TimestampMicro() int64
```

TimestampMicro retrieves and returns the timestamp in microseconds.

##### Example

``` go
```
##### func TimestampMicroStr 

``` go
func TimestampMicroStr() string
```

TimestampMicroStr is a convenience method which retrieves and returns the timestamp in microseconds as string.

##### func TimestampMilli 

``` go
func TimestampMilli() int64
```

TimestampMilli retrieves and returns the timestamp in milliseconds.

##### Example

``` go
```
##### func TimestampMilliStr 

``` go
func TimestampMilliStr() string
```

TimestampMilliStr is a convenience method which retrieves and returns the timestamp in milliseconds as string.

##### func TimestampNano 

``` go
func TimestampNano() int64
```

TimestampNano retrieves and returns the timestamp in nanoseconds.

##### Example

``` go
```
##### func TimestampNanoStr 

``` go
func TimestampNanoStr() string
```

TimestampNanoStr is a convenience method which retrieves and returns the timestamp in nanoseconds as string.

##### func TimestampStr 

``` go
func TimestampStr() string
```

TimestampStr is a convenience method which retrieves and returns the timestamp in seconds as string.

##### Example

``` go
```
### Types 

#### type Time 

``` go
type Time struct {
	// contains filtered or unexported fields
}
```

Time is a wrapper for time.Time for additional features.

##### func ConvertZone 

``` go
func ConvertZone(strTime string, toZone string, fromZone ...string) (*Time, error)
```

ConvertZone converts time in string `strTime` from `fromZone` to `toZone`. The parameter `fromZone` is unnecessary, it is current time zone in default.

##### Example

``` go
```
##### func New 

``` go
func New(param ...interface{}) *Time
```

New creates and returns a Time object with given parameter. The optional parameter is the time object which can be type of: time.Time/*time.Time, string or integer. Example: New("2024-10-29") New(1390876568) New(t) // The t is type of time.Time.

##### func NewFromStr 

``` go
func NewFromStr(str string) *Time
```

NewFromStr creates and returns a Time object with given string. Note that it returns nil if there's error occurs.

##### Example

``` go
```
##### func NewFromStrFormat 

``` go
func NewFromStrFormat(str string, format string) *Time
```

NewFromStrFormat creates and returns a Time object with given string and custom format like: Y-m-d H:i:s. Note that it returns nil if there's error occurs.

##### Example

``` go
```
##### func NewFromStrLayout 

``` go
func NewFromStrLayout(str string, layout string) *Time
```

NewFromStrLayout creates and returns a Time object with given string and stdlib layout like: 2006-01-02 15:04:05. Note that it returns nil if there's error occurs.

##### Example

``` go
```
##### func NewFromTime 

``` go
func NewFromTime(t time.Time) *Time
```

NewFromTime creates and returns a Time object with given time.Time object.

##### Example

``` go
```
##### func NewFromTimeStamp 

``` go
func NewFromTimeStamp(timestamp int64) *Time
```

NewFromTimeStamp creates and returns a Time object with given timestamp, which can be in seconds to nanoseconds. Eg: 1600443866 and 1600443866199266000 are both considered as valid timestamp number.

##### Example

``` go
```
##### func Now 

``` go
func Now() *Time
```

Now creates and returns a time object of now.

##### Example

``` go
```
##### func ParseTimeFromContent 

``` go
func ParseTimeFromContent(content string, format ...string) *Time
```

ParseTimeFromContent retrieves time information for content string, it then parses and returns it as *Time object. It returns the first time information if there are more than one time string in the content. It only retrieves and parses the time information with given first matched `format` if it's passed.

##### func StrToTime 

``` go
func StrToTime(str string, format ...string) (*Time, error)
```

StrToTime converts string to *Time object. It also supports timestamp string. The parameter `format` is unnecessary, which specifies the format for converting like "Y-m-d H:i:s". If `format` is given, it acts as same as function StrToTimeFormat. If `format` is not given, it converts string as a "standard" datetime string. Note that, it fails and returns error if there's no date string in `str`.

##### Example

``` go
```
##### func StrToTimeFormat 

``` go
func StrToTimeFormat(str string, format string) (*Time, error)
```

StrToTimeFormat parses string `str` to *Time object with given format `format`. The parameter `format` is like "Y-m-d H:i:s".

##### Example

``` go
```
##### func StrToTimeLayout 

``` go
func StrToTimeLayout(str string, layout string) (*Time, error)
```

StrToTimeLayout parses string `str` to *Time object with given format `layout`. The parameter `layout` is in stdlib format like "2006-01-02 15:04:05".

##### Example

``` go
```
##### (*Time) Add 

``` go
func (t *Time) Add(d time.Duration) *Time
```

Add adds the duration to current time.

##### Example

``` go
```
##### (*Time) AddDate 

``` go
func (t *Time) AddDate(years int, months int, days int) *Time
```

AddDate adds year, month and day to the time.

##### Example

``` go
```
##### (*Time) AddStr 

``` go
func (t *Time) AddStr(duration string) (*Time, error)
```

AddStr parses the given duration as string and adds it to current time.

##### Example

``` go
```
##### (*Time) After 

``` go
func (t *Time) After(u *Time) bool
```

After reports whether the time instant t is after u.

##### Example

``` go
```
##### (*Time) Before 

``` go
func (t *Time) Before(u *Time) bool
```

Before reports whether the time instant t is before u.

##### Example

``` go
```
##### (*Time) Clone 

``` go
func (t *Time) Clone() *Time
```

Clone returns a new Time object which is a clone of current time object.

##### (*Time) DayOfYear 

``` go
func (t *Time) DayOfYear() int
```

DayOfYear checks and returns the position of the day for the year.

##### Example

``` go
```
##### (*Time) DaysInMonth 

``` go
func (t *Time) DaysInMonth() int
```

DaysInMonth returns the day count of current month.

##### Example

``` go
```
##### (*Time) DeepCopy <-2.1.0

``` go
func (t *Time) DeepCopy() interface{}
```

DeepCopy implements interface for deep copy of current type.

##### (*Time) EndOfDay 

``` go
func (t *Time) EndOfDay(withNanoPrecision ...bool) *Time
```

EndOfDay clones and returns a new time which is the end of day the and its time is set to 23:59:59.

##### Example

``` go
```
##### (*Time) EndOfHalf 

``` go
func (t *Time) EndOfHalf(withNanoPrecision ...bool) *Time
```

EndOfHalf clones and returns a new time which is the end of the half year and its time is set to 23:59:59.

##### Example

``` go
```
##### (*Time) EndOfHour 

``` go
func (t *Time) EndOfHour(withNanoPrecision ...bool) *Time
```

EndOfHour clones and returns a new time of which the minutes and seconds are both set to 59.

##### Example

``` go
```
##### (*Time) EndOfMinute 

``` go
func (t *Time) EndOfMinute(withNanoPrecision ...bool) *Time
```

EndOfMinute clones and returns a new time of which the seconds is set to 59.

##### Example

``` go
```
##### (*Time) EndOfMonth 

``` go
func (t *Time) EndOfMonth(withNanoPrecision ...bool) *Time
```

EndOfMonth clones and returns a new time which is the end of the month and its time is set to 23:59:59.

##### Example

``` go
```
##### (*Time) EndOfQuarter 

``` go
func (t *Time) EndOfQuarter(withNanoPrecision ...bool) *Time
```

EndOfQuarter clones and returns a new time which is end of the quarter and its time is set to 23:59:59.

##### Example

``` go
```
##### (*Time) EndOfWeek 

``` go
func (t *Time) EndOfWeek(withNanoPrecision ...bool) *Time
```

EndOfWeek clones and returns a new time which is the end of week and its time is set to 23:59:59.

##### Example

``` go
```
##### (*Time) EndOfYear 

``` go
func (t *Time) EndOfYear(withNanoPrecision ...bool) *Time
```

EndOfYear clones and returns a new time which is the end of the year and its time is set to 23:59:59.

##### Example

``` go
```
##### (*Time) Equal 

``` go
func (t *Time) Equal(u *Time) bool
```

Equal reports whether t and u represent the same time instant. Two times can be equal even if they are in different locations. For example, 6:00 +0200 CEST and 4:00 UTC are Equal. See the documentation on the Time type for the pitfalls of using == with Time values; most code should use Equal instead.

##### Example

``` go
```
##### (*Time) Format 

``` go
func (t *Time) Format(format string) string
```

Format formats and returns the formatted result with custom `format`. Refer method Layout, if you want to follow stdlib layout.

##### Example

``` go
```
##### (*Time) FormatNew 

``` go
func (t *Time) FormatNew(format string) *Time
```

FormatNew formats and returns a new Time object with given custom `format`.

##### Example

``` go
```
##### (*Time) FormatTo 

``` go
func (t *Time) FormatTo(format string) *Time
```

FormatTo formats `t` with given custom `format`.

##### Example

``` go
```
##### (*Time) ISO8601 

``` go
func (t *Time) ISO8601() string
```

ISO8601 formats the time as ISO8601 and returns it as string.

##### (*Time) IsLeapYear 

``` go
func (t *Time) IsLeapYear() bool
```

IsLeapYear checks whether the time is leap year.

##### Example

``` go
```
##### (*Time) IsZero 

``` go
func (t *Time) IsZero() bool
```

IsZero reports whether t represents the zero time instant, January 1, year 1, 00:00:00 UTC.

##### Example

``` go
```
##### (*Time) Layout 

``` go
func (t *Time) Layout(layout string) string
```

Layout formats the time with stdlib layout and returns the formatted result.

##### Example

``` go
```
##### (*Time) LayoutNew 

``` go
func (t *Time) LayoutNew(layout string) *Time
```

LayoutNew formats the time with stdlib layout and returns the new Time object.

##### Example

``` go
```
##### (*Time) LayoutTo 

``` go
func (t *Time) LayoutTo(layout string) *Time
```

LayoutTo formats `t` with stdlib layout.

##### Example

``` go
```
##### (*Time) Local 

``` go
func (t *Time) Local() *Time
```

Local converts the time to local timezone.

##### (Time) MarshalJSON 

``` go
func (t Time) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. Note that, DO NOT use `(t *Time) MarshalJSON() ([]byte, error)` as it looses interface implement of `MarshalJSON` for struct of Time.

##### Example

``` go
```
##### (*Time) Microsecond 

``` go
func (t *Time) Microsecond() int
```

Microsecond returns the microsecond offset within the second specified by t, in the range [0, 999999].

##### (*Time) Millisecond 

``` go
func (t *Time) Millisecond() int
```

Millisecond returns the millisecond offset within the second specified by t, in the range [0, 999].

##### (*Time) Month 

``` go
func (t *Time) Month() int
```

Month returns the month of the year specified by t.

##### Example

``` go
```
##### (*Time) Nanosecond 

``` go
func (t *Time) Nanosecond() int
```

Nanosecond returns the nanosecond offset within the second specified by t, in the range [0, 999999999].

##### (*Time) NoValidation 

``` go
func (t *Time) NoValidation()
```

NoValidation marks this struct object will not be validated by package gvalid.

##### (*Time) RFC822 

``` go
func (t *Time) RFC822() string
```

RFC822 formats the time as RFC822 and returns it as string.

##### (*Time) Round 

``` go
func (t *Time) Round(d time.Duration) *Time
```

Round returns the result of rounding t to the nearest multiple of d (since the zero time). The rounding behavior for halfway values is to round up. If d <= 0, Round returns t stripped of any monotonic clock reading but otherwise unchanged.

Round operates on the time as an absolute duration since the zero time; it does not operate on the presentation form of the time. Thus, Round(Hour) may return a time with a non-zero minute, depending on the time's Location.

##### Example

``` go
```
##### (*Time) Scan 

``` go
func (t *Time) Scan(value interface{}) error
```

Scan implements interface used by Scan in package database/sql for Scanning value from database to local golang variable.

##### (*Time) Second 

``` go
func (t *Time) Second() int
```

Second returns the second offset within the minute specified by t, in the range [0, 59].

##### Example

``` go
```
##### (*Time) StartOfDay 

``` go
func (t *Time) StartOfDay() *Time
```

StartOfDay clones and returns a new time which is the start of day, its time is set to 00:00:00.

##### Example

``` go
```
##### (*Time) StartOfHalf 

``` go
func (t *Time) StartOfHalf() *Time
```

StartOfHalf clones and returns a new time which is the first day of the half year and its time is set to 00:00:00.

##### Example

``` go
```
##### (*Time) StartOfHour 

``` go
func (t *Time) StartOfHour() *Time
```

StartOfHour clones and returns a new time of which the hour, minutes and seconds are set to 0.

##### Example

``` go
```
##### (*Time) StartOfMinute 

``` go
func (t *Time) StartOfMinute() *Time
```

StartOfMinute clones and returns a new time of which the seconds is set to 0.

##### Example

``` go
```
##### (*Time) StartOfMonth 

``` go
func (t *Time) StartOfMonth() *Time
```

StartOfMonth clones and returns a new time which is the first day of the month and its is set to 00:00:00

##### (*Time) StartOfQuarter 

``` go
func (t *Time) StartOfQuarter() *Time
```

StartOfQuarter clones and returns a new time which is the first day of the quarter and its time is set to 00:00:00.

##### Example

``` go
```
##### (*Time) StartOfWeek 

``` go
func (t *Time) StartOfWeek() *Time
```

StartOfWeek clones and returns a new time which is the first day of week and its time is set to 00:00:00.

##### Example

``` go
```
##### (*Time) StartOfYear 

``` go
func (t *Time) StartOfYear() *Time
```

StartOfYear clones and returns a new time which is the first day of the year and its time is set to 00:00:00.

##### Example

``` go
```
##### (*Time) String 

``` go
func (t *Time) String() string
```

String returns current time object as string.

##### Example

``` go
```
##### (*Time) Sub 

``` go
func (t *Time) Sub(u *Time) time.Duration
```

Sub returns the duration t-u. If the result exceeds the maximum (or minimum) value that can be stored in a Duration, the maximum (or minimum) duration will be returned. To compute t-d for a duration d, use t.Add(-d).

##### Example

``` go
```
##### (*Time) Timestamp 

``` go
func (t *Time) Timestamp() int64
```

Timestamp returns the timestamp in seconds.

##### Example

``` go
```
##### (*Time) TimestampMicro 

``` go
func (t *Time) TimestampMicro() int64
```

TimestampMicro returns the timestamp in microseconds.

##### Example

``` go
```
##### (*Time) TimestampMicroStr 

``` go
func (t *Time) TimestampMicroStr() string
```

TimestampMicroStr is a convenience method which retrieves and returns the timestamp in microseconds as string.

##### (*Time) TimestampMilli 

``` go
func (t *Time) TimestampMilli() int64
```

TimestampMilli returns the timestamp in milliseconds.

##### Example

``` go
```
##### (*Time) TimestampMilliStr 

``` go
func (t *Time) TimestampMilliStr() string
```

TimestampMilliStr is a convenience method which retrieves and returns the timestamp in milliseconds as string.

##### (*Time) TimestampNano 

``` go
func (t *Time) TimestampNano() int64
```

TimestampNano returns the timestamp in nanoseconds.

##### Example

``` go
```
##### (*Time) TimestampNanoStr 

``` go
func (t *Time) TimestampNanoStr() string
```

TimestampNanoStr is a convenience method which retrieves and returns the timestamp in nanoseconds as string.

##### (*Time) TimestampStr 

``` go
func (t *Time) TimestampStr() string
```

TimestampStr is a convenience method which retrieves and returns the timestamp in seconds as string.

##### Example

``` go
```
##### (*Time) ToLocation 

``` go
func (t *Time) ToLocation(location *time.Location) *Time
```

ToLocation converts current time to specified location.

##### (*Time) ToZone 

``` go
func (t *Time) ToZone(zone string) (*Time, error)
```

ToZone converts current time to specified zone like: Asia/Shanghai.

##### Example

``` go
```
##### (*Time) Truncate 

``` go
func (t *Time) Truncate(d time.Duration) *Time
```

Truncate returns the result of rounding t down to a multiple of d (since the zero time). If d <= 0, Truncate returns t stripped of any monotonic clock reading but otherwise unchanged.

Truncate operates on the time as an absolute duration since the zero time; it does not operate on the presentation form of the time. Thus, Truncate(Hour) may return a time with a non-zero minute, depending on the time's Location.

##### Example

``` go
```
##### (*Time) UTC 

``` go
func (t *Time) UTC() *Time
```

UTC converts current time to UTC timezone.

##### (*Time) UnmarshalJSON 

``` go
func (t *Time) UnmarshalJSON(b []byte) error
```

UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.

##### Example

``` go
```
##### (*Time) UnmarshalText 

``` go
func (t *Time) UnmarshalText(data []byte) error
```

UnmarshalText implements the encoding.TextUnmarshaler interface. Note that it overwrites the same implementer of `time.Time`.

##### (*Time) Value 

``` go
func (t *Time) Value() (driver.Value, error)
```

Value is the interface providing the Value method for package database/sql/driver for retrieving value from golang variable to database.

##### (*Time) WeeksOfYear 

``` go
func (t *Time) WeeksOfYear() int
```

WeeksOfYear returns the point of current week for the year.

Example WeekOfYear

``` go
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

