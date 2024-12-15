+++
title = "assert"
date = 2024-12-15T11:07:36+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/stretchr/testify/assert](https://pkg.go.dev/github.com/stretchr/testify/assert)
>
> 收录该文档时间： `2024-12-15T11:07:36+08:00`
>
> 版本：[Version: v1.10.0](https://pkg.go.dev/github.com/stretchr/testify/assert?tab=versions)

### Overview [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#pkg-overview)

- [Example Usage](https://pkg.go.dev/github.com/stretchr/testify/assert#hdr-Example_Usage)
- [Assertions](https://pkg.go.dev/github.com/stretchr/testify/assert#hdr-Assertions)

Package assert provides a set of comprehensive testing tools for use with the normal Go testing system.

#### Example Usage [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#hdr-Example_Usage)

The following is a complete example using assert in a standard test function:

```
import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

  var a string = "Hello"
  var b string = "Hello"

  assert.Equal(t, a, b, "The two words should be the same.")

}
```

if you assert many times, use the format below:

```
import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
  assert := assert.New(t)

  var a string = "Hello"
  var b string = "Hello"

  assert.Equal(a, b, "The two words should be the same.")
}
```

#### Assertions [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#hdr-Assertions)

Assertions allow you to easily write test code, and are global funcs in the `assert` package. All assertion functions take, as the first argument, the `*testing.T` object provided by the testing framework. This allows the assertion funcs to write the failings and other details to the correct place.

Every assertion function also takes an optional string message as the final argument, allowing custom error messages to be appended to the message the assertion method outputs.

### Constants [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#pkg-constants)

This section is empty.

### Variables [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#pkg-variables)

[View Source](https://github.com/stretchr/testify/blob/v1.10.0/assert/errors.go#L10)

```
var AnError = errors.New("assert.AnError general error for testing")
```

AnError is an error instance useful for testing. If the code does not care about error specifics, and only needs to return the error for example, this error should be used to make the test code more readable.

### Functions [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#pkg-functions)

#### func [CallerInfo](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L212) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#CallerInfo)

```
func CallerInfo() []string
```

CallerInfo returns an array of strings containing the file and line number of each stack frame leading from the current test to the assert call that failed.

#### func [Condition](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1212) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Condition)

```
func Condition(t TestingT, comp Comparison, msgAndArgs ...interface{}) bool
```

Condition uses a Comparison to assert a complex condition.

#### func [Conditionf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L12) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Conditionf)added in v1.2.0

```
func Conditionf(t TestingT, comp Comparison, msg string, args ...interface{}) bool
```

Conditionf uses a Comparison to assert a complex condition.

#### func [Contains](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L927) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Contains)

```
func Contains(t TestingT, s, contains interface{}, msgAndArgs ...interface{}) bool
```

Contains asserts that the specified string, list(array, slice...) or map contains the specified substring or element.

```
assert.Contains(t, "Hello World", "World")
assert.Contains(t, ["Hello", "World"], "World")
assert.Contains(t, {"Hello": "World"}, "Hello")
```

#### func [Containsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L25) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Containsf)added in v1.2.0

```
func Containsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{}) bool
```

Containsf asserts that the specified string, list(array, slice...) or map contains the specified substring or element.

```
assert.Containsf(t, "Hello World", "World", "error message %s", "formatted")
assert.Containsf(t, ["Hello", "World"], "World", "error message %s", "formatted")
assert.Containsf(t, {"Hello": "World"}, "Hello", "error message %s", "formatted")
```

#### func [DirExists](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1768) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#DirExists)added in v1.2.0

```
func DirExists(t TestingT, path string, msgAndArgs ...interface{}) bool
```

DirExists checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.

#### func [DirExistsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L34) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#DirExistsf)added in v1.2.0

```
func DirExistsf(t TestingT, path string, msg string, args ...interface{}) bool
```

DirExistsf checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.

#### func [ElementsMatch](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1087) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ElementsMatch)added in v1.2.0

```
func ElementsMatch(t TestingT, listA, listB interface{}, msgAndArgs ...interface{}) (ok bool)
```

ElementsMatch asserts that the specified listA(array, slice...) is equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should match.

assert.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])

#### func [ElementsMatchf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L46) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ElementsMatchf)added in v1.2.0

```
func ElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{}) bool
```

ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should match.

assert.ElementsMatchf(t, [1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")

#### func [Empty](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L749) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Empty)

```
func Empty(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
```

Empty asserts that the specified object is empty. I.e. nil, "", false, 0 or either a slice or a channel with len == 0.

```
assert.Empty(t, obj)
```

#### func [Emptyf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L57) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Emptyf)added in v1.2.0

```
func Emptyf(t TestingT, object interface{}, msg string, args ...interface{}) bool
```

Emptyf asserts that the specified object is empty. I.e. nil, "", false, 0 or either a slice or a channel with len == 0.

```
assert.Emptyf(t, obj, "error message %s", "formatted")
```

#### func [Equal](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L460) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Equal)

```
func Equal(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool
```

Equal asserts that two objects are equal.

```
assert.Equal(t, 123, 123)
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses). Function equality cannot be determined and will always fail.

#### func [EqualError](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1614) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#EqualError)

```
func EqualError(t TestingT, theError error, errString string, msgAndArgs ...interface{}) bool
```

EqualError asserts that a function returned an error (i.e. not `nil`) and that it is equal to the provided error.

```
actualObj, err := SomeFunction()
assert.EqualError(t, err,  expectedErrorString)
```

#### func [EqualErrorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L83) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#EqualErrorf)added in v1.2.0

```
func EqualErrorf(t TestingT, theError error, errString string, msg string, args ...interface{}) bool
```

EqualErrorf asserts that a function returned an error (i.e. not `nil`) and that it is equal to the provided error.

```
actualObj, err := SomeFunction()
assert.EqualErrorf(t, err,  expectedErrorString, "error message %s", "formatted")
```

#### func [EqualExportedValues](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L626) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#EqualExportedValues)added in v1.8.3

```
func EqualExportedValues(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool
```

EqualExportedValues asserts that the types of two objects are equal and their public fields are also equal. This is useful for comparing structs that have private fields that could potentially differ.

```
 type S struct {
	Exported     	int
	notExported   	int
 }
 assert.EqualExportedValues(t, S{1, 2}, S{1, 3}) => true
 assert.EqualExportedValues(t, S{1, 2}, S{2, 3}) => false
```

#### func [EqualExportedValuesf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L100) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#EqualExportedValuesf)added in v1.8.3

```
func EqualExportedValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

EqualExportedValuesf asserts that the types of two objects are equal and their public fields are also equal. This is useful for comparing structs that have private fields that could potentially differ.

```
 type S struct {
	Exported     	int
	notExported   	int
 }
 assert.EqualExportedValuesf(t, S{1, 2}, S{1, 3}, "error message %s", "formatted") => true
 assert.EqualExportedValuesf(t, S{1, 2}, S{2, 3}, "error message %s", "formatted") => false
```

#### func [EqualValues](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L599) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#EqualValues)

```
func EqualValues(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool
```

EqualValues asserts that two objects are equal or convertible to the larger type and equal.

```
assert.EqualValues(t, uint32(123), int32(123))
```

#### func [EqualValuesf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L111) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#EqualValuesf)added in v1.2.0

```
func EqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

EqualValuesf asserts that two objects are equal or convertible to the larger type and equal.

```
assert.EqualValuesf(t, uint32(123), int32(123), "error message %s", "formatted")
```

#### func [Equalf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L71) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Equalf)added in v1.2.0

```
func Equalf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

Equalf asserts that two objects are equal.

```
assert.Equalf(t, 123, 123, "error message %s", "formatted")
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses). Function equality cannot be determined and will always fail.

#### func [Error](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1598) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Error)

```
func Error(t TestingT, err error, msgAndArgs ...interface{}) bool
```

Error asserts that a function returned an error (i.e. not `nil`).

```
  actualObj, err := SomeFunction()
  if assert.Error(t, err) {
	   assert.Equal(t, expectedError, err)
  }
```

#### func [ErrorAs](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L2138) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ErrorAs)added in v1.7.0

```
func ErrorAs(t TestingT, err error, target interface{}, msgAndArgs ...interface{}) bool
```

ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value. This is a wrapper for errors.As.

#### func [ErrorAsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L133) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ErrorAsf)added in v1.7.0

```
func ErrorAsf(t TestingT, err error, target interface{}, msg string, args ...interface{}) bool
```

ErrorAsf asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value. This is a wrapper for errors.As.

#### func [ErrorContains](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1637) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ErrorContains)added in v1.7.1

```
func ErrorContains(t TestingT, theError error, contains string, msgAndArgs ...interface{}) bool
```

ErrorContains asserts that a function returned an error (i.e. not `nil`) and that the error contains the specified substring.

```
actualObj, err := SomeFunction()
assert.ErrorContains(t, err,  expectedErrorSubString)
```

#### func [ErrorContainsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L145) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ErrorContainsf)added in v1.7.1

```
func ErrorContainsf(t TestingT, theError error, contains string, msg string, args ...interface{}) bool
```

ErrorContainsf asserts that a function returned an error (i.e. not `nil`) and that the error contains the specified substring.

```
actualObj, err := SomeFunction()
assert.ErrorContainsf(t, err,  expectedErrorSubString, "error message %s", "formatted")
```

#### func [ErrorIs](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L2092) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ErrorIs)added in v1.7.0

```
func ErrorIs(t TestingT, err, target error, msgAndArgs ...interface{}) bool
```

ErrorIs asserts that at least one of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func [ErrorIsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L154) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ErrorIsf)added in v1.7.0

```
func ErrorIsf(t TestingT, err error, target error, msg string, args ...interface{}) bool
```

ErrorIsf asserts that at least one of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func [Errorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L124) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Errorf)added in v1.2.0

```
func Errorf(t TestingT, err error, msg string, args ...interface{}) bool
```

Errorf asserts that a function returned an error (i.e. not `nil`).

```
  actualObj, err := SomeFunction()
  if assert.Errorf(t, err, "error message %s", "formatted") {
	   assert.Equal(t, expectedErrorf, err)
  }
```

#### func [Eventually](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1930) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Eventually)added in v1.4.0

```
func Eventually(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool
```

Eventually asserts that given condition will be met in waitFor time, periodically checking target function each tick.

```
assert.Eventually(t, func() bool { return true; }, time.Second, 10*time.Millisecond)
```

#### func [EventuallyWithT](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L2016) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#EventuallyWithT)added in v1.8.3

```
func EventuallyWithT(t TestingT, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool
```

EventuallyWithT asserts that given condition will be met in waitFor time, periodically checking target function each tick. In contrast to Eventually, it supplies a CollectT to the condition function, so that the condition function can use the CollectT to call other assertions. The condition is considered "met" if no errors are raised in a tick. The supplied CollectT collects all errors from one tick (if there are any). If the condition is not met before waitFor, the collected errors of the last tick are copied to t.

```
externalValue := false
go func() {
	time.Sleep(8*time.Second)
	externalValue = true
}()
assert.EventuallyWithT(t, func(c *assert.CollectT) {
	// add assertions as needed; any assertion failure will fail the current tick
	assert.True(c, externalValue, "expected 'externalValue' to be true")
}, 10*time.Second, 1*time.Second, "external state has not changed to 'true'; still false")
```

#### func [EventuallyWithTf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L190) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#EventuallyWithTf)added in v1.8.3

```
func EventuallyWithTf(t TestingT, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool
```

EventuallyWithTf asserts that given condition will be met in waitFor time, periodically checking target function each tick. In contrast to Eventually, it supplies a CollectT to the condition function, so that the condition function can use the CollectT to call other assertions. The condition is considered "met" if no errors are raised in a tick. The supplied CollectT collects all errors from one tick (if there are any). If the condition is not met before waitFor, the collected errors of the last tick are copied to t.

```
externalValue := false
go func() {
	time.Sleep(8*time.Second)
	externalValue = true
}()
assert.EventuallyWithTf(t, func(c *assert.CollectT, "error message %s", "formatted") {
	// add assertions as needed; any assertion failure will fail the current tick
	assert.True(c, externalValue, "expected 'externalValue' to be true")
}, 10*time.Second, 1*time.Second, "external state has not changed to 'true'; still false")
```

#### func [Eventuallyf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L165) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Eventuallyf)added in v1.4.0

```
func Eventuallyf(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool
```

Eventuallyf asserts that given condition will be met in waitFor time, periodically checking target function each tick.

```
assert.Eventuallyf(t, func() bool { return true; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
```

#### func [Exactly](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L655) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Exactly)

```
func Exactly(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool
```

Exactly asserts that two objects are equal in value and type.

```
assert.Exactly(t, int32(123), int64(123))
```

#### func [Exactlyf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L200) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Exactlyf)added in v1.2.0

```
func Exactlyf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

Exactlyf asserts that two objects are equal in value and type.

```
assert.Exactlyf(t, int32(123), int64(123), "error message %s", "formatted")
```

#### func [Fail](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L348) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Fail)

```
func Fail(t TestingT, failureMessage string, msgAndArgs ...interface{}) bool
```

Fail reports a failure through

#### func [FailNow](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L327) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#FailNow)

```
func FailNow(t TestingT, failureMessage string, msgAndArgs ...interface{}) bool
```

FailNow fails test

#### func [FailNowf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L216) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#FailNowf)added in v1.2.0

```
func FailNowf(t TestingT, failureMessage string, msg string, args ...interface{}) bool
```

FailNowf fails test

#### func [Failf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L208) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Failf)added in v1.2.0

```
func Failf(t TestingT, failureMessage string, msg string, args ...interface{}) bool
```

Failf reports a failure through

#### func [False](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L828) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#False)

```
func False(t TestingT, value bool, msgAndArgs ...interface{}) bool
```

False asserts that the specified value is false.

```
assert.False(t, myBool)
```

#### func [Falsef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L226) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Falsef)added in v1.2.0

```
func Falsef(t TestingT, value bool, msg string, args ...interface{}) bool
```

Falsef asserts that the specified value is false.

```
assert.Falsef(t, myBool, "error message %s", "formatted")
```

#### func [FileExists](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1733) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#FileExists)added in v1.2.0

```
func FileExists(t TestingT, path string, msgAndArgs ...interface{}) bool
```

FileExists checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.

#### func [FileExistsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L235) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#FileExistsf)added in v1.2.0

```
func FileExistsf(t TestingT, path string, msg string, args ...interface{}) bool
```

FileExistsf checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.

#### func [Greater](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_compare.go#L389) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Greater)added in v1.4.0

```
func Greater(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool
```

Greater asserts that the first element is greater than the second

```
assert.Greater(t, 2, 1)
assert.Greater(t, float64(2), float64(1))
assert.Greater(t, "b", "a")
```

#### func [GreaterOrEqual](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_compare.go#L402) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#GreaterOrEqual)added in v1.4.0

```
func GreaterOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool
```

GreaterOrEqual asserts that the first element is greater than or equal to the second

```
assert.GreaterOrEqual(t, 2, 1)
assert.GreaterOrEqual(t, 2, 2)
assert.GreaterOrEqual(t, "b", "a")
assert.GreaterOrEqual(t, "b", "b")
```

#### func [GreaterOrEqualf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L260) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#GreaterOrEqualf)added in v1.4.0

```
func GreaterOrEqualf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool
```

GreaterOrEqualf asserts that the first element is greater than or equal to the second

```
assert.GreaterOrEqualf(t, 2, 1, "error message %s", "formatted")
assert.GreaterOrEqualf(t, 2, 2, "error message %s", "formatted")
assert.GreaterOrEqualf(t, "b", "a", "error message %s", "formatted")
assert.GreaterOrEqualf(t, "b", "b", "error message %s", "formatted")
```

#### func [Greaterf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L247) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Greaterf)added in v1.4.0

```
func Greaterf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool
```

Greaterf asserts that the first element is greater than the second

```
assert.Greaterf(t, 2, 1, "error message %s", "formatted")
assert.Greaterf(t, float64(2), float64(1), "error message %s", "formatted")
assert.Greaterf(t, "b", "a", "error message %s", "formatted")
```

#### func [HTTPBody](https://github.com/stretchr/testify/blob/v1.10.0/assert/http_assertions.go#L114) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPBody)

```
func HTTPBody(handler http.HandlerFunc, method, url string, values url.Values) string
```

HTTPBody is a helper that returns HTTP body of the response. It returns empty string if building a new request fails.

#### func [HTTPBodyContains](https://github.com/stretchr/testify/blob/v1.10.0/assert/http_assertions.go#L133) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPBodyContains)

```
func HTTPBodyContains(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool
```

HTTPBodyContains asserts that a specified handler returns a body that contains a string.

```
assert.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPBodyContainsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L273) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPBodyContainsf)added in v1.2.0

```
func HTTPBodyContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) bool
```

HTTPBodyContainsf asserts that a specified handler returns a body that contains a string.

```
assert.HTTPBodyContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPBodyNotContains](https://github.com/stretchr/testify/blob/v1.10.0/assert/http_assertions.go#L153) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPBodyNotContains)

```
func HTTPBodyNotContains(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool
```

HTTPBodyNotContains asserts that a specified handler returns a body that does not contain a string.

```
assert.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPBodyNotContainsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L286) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPBodyNotContainsf)added in v1.2.0

```
func HTTPBodyNotContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) bool
```

HTTPBodyNotContainsf asserts that a specified handler returns a body that does not contain a string.

```
assert.HTTPBodyNotContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPError](https://github.com/stretchr/testify/blob/v1.10.0/assert/http_assertions.go#L73) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPError)

```
func HTTPError(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool
```

HTTPError asserts that a specified handler returns an error status code.

```
assert.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPErrorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L298) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPErrorf)added in v1.2.0

```
func HTTPErrorf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool
```

HTTPErrorf asserts that a specified handler returns an error status code.

```
assert.HTTPErrorf(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPRedirect](https://github.com/stretchr/testify/blob/v1.10.0/assert/http_assertions.go#L51) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPRedirect)

```
func HTTPRedirect(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool
```

HTTPRedirect asserts that a specified handler returns a redirect status code.

```
assert.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPRedirectf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L310) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPRedirectf)added in v1.2.0

```
func HTTPRedirectf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool
```

HTTPRedirectf asserts that a specified handler returns a redirect status code.

```
assert.HTTPRedirectf(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPStatusCode](https://github.com/stretchr/testify/blob/v1.10.0/assert/http_assertions.go#L95) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPStatusCode)added in v1.6.0

```
func HTTPStatusCode(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) bool
```

HTTPStatusCode asserts that a specified handler returns a specified status code.

```
assert.HTTPStatusCode(t, myHandler, "GET", "/notImplemented", nil, 501)
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPStatusCodef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L322) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPStatusCodef)added in v1.6.0

```
func HTTPStatusCodef(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{}) bool
```

HTTPStatusCodef asserts that a specified handler returns a specified status code.

```
assert.HTTPStatusCodef(t, myHandler, "GET", "/notImplemented", nil, 501, "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPSuccess](https://github.com/stretchr/testify/blob/v1.10.0/assert/http_assertions.go#L29) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPSuccess)

```
func HTTPSuccess(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool
```

HTTPSuccess asserts that a specified handler returns a success status code.

```
assert.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPSuccessf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L334) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#HTTPSuccessf)added in v1.2.0

```
func HTTPSuccessf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool
```

HTTPSuccessf asserts that a specified handler returns a success status code.

```
assert.HTTPSuccessf(t, myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func [Implements](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L405) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Implements)

```
func Implements(t TestingT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool
```

Implements asserts that an object is implemented by the specified interface.

```
assert.Implements(t, (*MyInterface)(nil), new(MyObject))
```

#### func [Implementsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L344) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Implementsf)added in v1.2.0

```
func Implementsf(t TestingT, interfaceObject interface{}, object interface{}, msg string, args ...interface{}) bool
```

Implementsf asserts that an object is implemented by the specified interface.

```
assert.Implementsf(t, (*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
```

#### func [InDelta](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1395) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#InDelta)

```
func InDelta(t TestingT, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool
```

InDelta asserts that the two numerals are within delta of each other.

```
assert.InDelta(t, math.Pi, 22/7.0, 0.01)
```

#### func [InDeltaMapValues](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1452) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#InDeltaMapValues)added in v1.2.0

```
func InDeltaMapValues(t TestingT, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool
```

InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.

#### func [InDeltaMapValuesf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L362) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#InDeltaMapValuesf)added in v1.2.0

```
func InDeltaMapValuesf(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool
```

InDeltaMapValuesf is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.

#### func [InDeltaSlice](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1428) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#InDeltaSlice)

```
func InDeltaSlice(t TestingT, expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool
```

InDeltaSlice is the same as InDelta, except it compares two slices.

#### func [InDeltaSlicef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L370) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#InDeltaSlicef)added in v1.2.0

```
func InDeltaSlicef(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool
```

InDeltaSlicef is the same as InDelta, except it compares two slices.

#### func [InDeltaf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L354) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#InDeltaf)added in v1.2.0

```
func InDeltaf(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool
```

InDeltaf asserts that the two numerals are within delta of each other.

```
assert.InDeltaf(t, math.Pi, 22/7.0, 0.01, "error message %s", "formatted")
```

#### func [InEpsilon](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1518) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#InEpsilon)

```
func InEpsilon(t TestingT, expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool
```

InEpsilon asserts that expected and actual have a relative error less than epsilon

#### func [InEpsilonSlice](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1541) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#InEpsilonSlice)

```
func InEpsilonSlice(t TestingT, expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool
```

InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.

#### func [InEpsilonSlicef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L386) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#InEpsilonSlicef)added in v1.2.0

```
func InEpsilonSlicef(t TestingT, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool
```

InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.

#### func [InEpsilonf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L378) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#InEpsilonf)added in v1.2.0

```
func InEpsilonf(t TestingT, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool
```

InEpsilonf asserts that expected and actual have a relative error less than epsilon

#### func [IsDecreasing](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_order.go#L70) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#IsDecreasing)added in v1.7.0

```
func IsDecreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
```

IsDecreasing asserts that the collection is decreasing

```
assert.IsDecreasing(t, []int{2, 1, 0})
assert.IsDecreasing(t, []float{2, 1})
assert.IsDecreasing(t, []string{"b", "a"})
```

#### func [IsDecreasingf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L398) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#IsDecreasingf)added in v1.7.0

```
func IsDecreasingf(t TestingT, object interface{}, msg string, args ...interface{}) bool
```

IsDecreasingf asserts that the collection is decreasing

```
assert.IsDecreasingf(t, []int{2, 1, 0}, "error message %s", "formatted")
assert.IsDecreasingf(t, []float{2, 1}, "error message %s", "formatted")
assert.IsDecreasingf(t, []string{"b", "a"}, "error message %s", "formatted")
```

#### func [IsIncreasing](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_order.go#L52) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#IsIncreasing)added in v1.7.0

```
func IsIncreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
```

IsIncreasing asserts that the collection is increasing

```
assert.IsIncreasing(t, []int{1, 2, 3})
assert.IsIncreasing(t, []float{1, 2})
assert.IsIncreasing(t, []string{"a", "b"})
```

#### func [IsIncreasingf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L410) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#IsIncreasingf)added in v1.7.0

```
func IsIncreasingf(t TestingT, object interface{}, msg string, args ...interface{}) bool
```

IsIncreasingf asserts that the collection is increasing

```
assert.IsIncreasingf(t, []int{1, 2, 3}, "error message %s", "formatted")
assert.IsIncreasingf(t, []float{1, 2}, "error message %s", "formatted")
assert.IsIncreasingf(t, []string{"a", "b"}, "error message %s", "formatted")
```

#### func [IsNonDecreasing](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_order.go#L79) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#IsNonDecreasing)added in v1.7.0

```
func IsNonDecreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
```

IsNonDecreasing asserts that the collection is not decreasing

```
assert.IsNonDecreasing(t, []int{1, 1, 2})
assert.IsNonDecreasing(t, []float{1, 2})
assert.IsNonDecreasing(t, []string{"a", "b"})
```

#### func [IsNonDecreasingf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L422) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#IsNonDecreasingf)added in v1.7.0

```
func IsNonDecreasingf(t TestingT, object interface{}, msg string, args ...interface{}) bool
```

IsNonDecreasingf asserts that the collection is not decreasing

```
assert.IsNonDecreasingf(t, []int{1, 1, 2}, "error message %s", "formatted")
assert.IsNonDecreasingf(t, []float{1, 2}, "error message %s", "formatted")
assert.IsNonDecreasingf(t, []string{"a", "b"}, "error message %s", "formatted")
```

#### func [IsNonIncreasing](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_order.go#L61) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#IsNonIncreasing)added in v1.7.0

```
func IsNonIncreasing(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
```

IsNonIncreasing asserts that the collection is not increasing

```
assert.IsNonIncreasing(t, []int{2, 1, 1})
assert.IsNonIncreasing(t, []float{2, 1})
assert.IsNonIncreasing(t, []string{"b", "a"})
```

#### func [IsNonIncreasingf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L434) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#IsNonIncreasingf)added in v1.7.0

```
func IsNonIncreasingf(t TestingT, object interface{}, msg string, args ...interface{}) bool
```

IsNonIncreasingf asserts that the collection is not increasing

```
assert.IsNonIncreasingf(t, []int{2, 1, 1}, "error message %s", "formatted")
assert.IsNonIncreasingf(t, []float{2, 1}, "error message %s", "formatted")
assert.IsNonIncreasingf(t, []string{"b", "a"}, "error message %s", "formatted")
```

#### func [IsType](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L441) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#IsType)

```
func IsType(t TestingT, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool
```

IsType asserts that the specified objects are of the same type.

#### func [IsTypef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L442) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#IsTypef)added in v1.2.0

```
func IsTypef(t TestingT, expectedType interface{}, object interface{}, msg string, args ...interface{}) bool
```

IsTypef asserts that the specified objects are of the same type.

#### func [JSONEq](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1807) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#JSONEq)

```
func JSONEq(t TestingT, expected string, actual string, msgAndArgs ...interface{}) bool
```

JSONEq asserts that two JSON strings are equivalent.

```
assert.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
```

#### func [JSONEqf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L452) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#JSONEqf)added in v1.2.0

```
func JSONEqf(t TestingT, expected string, actual string, msg string, args ...interface{}) bool
```

JSONEqf asserts that two JSON strings are equivalent.

```
assert.JSONEqf(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
```

#### func [Len](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L795) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Len)

```
func Len(t TestingT, object interface{}, length int, msgAndArgs ...interface{}) bool
```

Len asserts that the specified object has specific length. Len also fails if the object has a type that len() not accept.

```
assert.Len(t, mySlice, 3)
```

#### func [Lenf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L463) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Lenf)added in v1.2.0

```
func Lenf(t TestingT, object interface{}, length int, msg string, args ...interface{}) bool
```

Lenf asserts that the specified object has specific length. Lenf also fails if the object has a type that len() not accept.

```
assert.Lenf(t, mySlice, 3, "error message %s", "formatted")
```

#### func [Less](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_compare.go#L414) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Less)added in v1.4.0

```
func Less(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool
```

Less asserts that the first element is less than the second

```
assert.Less(t, 1, 2)
assert.Less(t, float64(1), float64(2))
assert.Less(t, "a", "b")
```

#### func [LessOrEqual](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_compare.go#L427) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#LessOrEqual)added in v1.4.0

```
func LessOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool
```

LessOrEqual asserts that the first element is less than or equal to the second

```
assert.LessOrEqual(t, 1, 2)
assert.LessOrEqual(t, 2, 2)
assert.LessOrEqual(t, "a", "b")
assert.LessOrEqual(t, "b", "b")
```

#### func [LessOrEqualf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L488) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#LessOrEqualf)added in v1.4.0

```
func LessOrEqualf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool
```

LessOrEqualf asserts that the first element is less than or equal to the second

```
assert.LessOrEqualf(t, 1, 2, "error message %s", "formatted")
assert.LessOrEqualf(t, 2, 2, "error message %s", "formatted")
assert.LessOrEqualf(t, "a", "b", "error message %s", "formatted")
assert.LessOrEqualf(t, "b", "b", "error message %s", "formatted")
```

#### func [Lessf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L475) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Lessf)added in v1.4.0

```
func Lessf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool
```

Lessf asserts that the first element is less than the second

```
assert.Lessf(t, 1, 2, "error message %s", "formatted")
assert.Lessf(t, float64(1), float64(2), "error message %s", "formatted")
assert.Lessf(t, "a", "b", "error message %s", "formatted")
```

#### func [Negative](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_compare.go#L450) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Negative)added in v1.7.0

```
func Negative(t TestingT, e interface{}, msgAndArgs ...interface{}) bool
```

Negative asserts that the specified element is negative

```
assert.Negative(t, -1)
assert.Negative(t, -1.23)
```

#### func [Negativef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L499) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Negativef)added in v1.7.0

```
func Negativef(t TestingT, e interface{}, msg string, args ...interface{}) bool
```

Negativef asserts that the specified element is negative

```
assert.Negativef(t, -1, "error message %s", "formatted")
assert.Negativef(t, -1.23, "error message %s", "formatted")
```

#### func [Never](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L2061) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Never)added in v1.5.0

```
func Never(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool
```

Never asserts that the given condition doesn't satisfy in waitFor time, periodically checking the target function each tick.

```
assert.Never(t, func() bool { return false; }, time.Second, 10*time.Millisecond)
```

#### func [Neverf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L510) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Neverf)added in v1.5.0

```
func Neverf(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool
```

Neverf asserts that the given condition doesn't satisfy in waitFor time, periodically checking the target function each tick.

```
assert.Neverf(t, func() bool { return false; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
```

#### func [Nil](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L706) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Nil)

```
func Nil(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
```

Nil asserts that the specified object is nil.

```
assert.Nil(t, err)
```

#### func [Nilf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L520) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Nilf)added in v1.2.0

```
func Nilf(t TestingT, object interface{}, msg string, args ...interface{}) bool
```

Nilf asserts that the specified object is nil.

```
assert.Nilf(t, err, "error message %s", "formatted")
```

#### func [NoDirExists](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1787) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NoDirExists)added in v1.5.0

```
func NoDirExists(t TestingT, path string, msgAndArgs ...interface{}) bool
```

NoDirExists checks whether a directory does not exist in the given path. It fails if the path points to an existing _directory_ only.

#### func [NoDirExistsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L529) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NoDirExistsf)added in v1.5.0

```
func NoDirExistsf(t TestingT, path string, msg string, args ...interface{}) bool
```

NoDirExistsf checks whether a directory does not exist in the given path. It fails if the path points to an existing _directory_ only.

#### func [NoError](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1581) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NoError)

```
func NoError(t TestingT, err error, msgAndArgs ...interface{}) bool
```

NoError asserts that a function returned no error (i.e. `nil`).

```
  actualObj, err := SomeFunction()
  if assert.NoError(t, err) {
	   assert.Equal(t, expectedObj, actualObj)
  }
```

#### func [NoErrorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L542) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NoErrorf)added in v1.2.0

```
func NoErrorf(t TestingT, err error, msg string, args ...interface{}) bool
```

NoErrorf asserts that a function returned no error (i.e. `nil`).

```
  actualObj, err := SomeFunction()
  if assert.NoErrorf(t, err, "error message %s", "formatted") {
	   assert.Equal(t, expectedObj, actualObj)
  }
```

#### func [NoFileExists](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1752) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NoFileExists)added in v1.5.0

```
func NoFileExists(t TestingT, path string, msgAndArgs ...interface{}) bool
```

NoFileExists checks whether a file does not exist in a given path. It fails if the path points to an existing _file_ only.

#### func [NoFileExistsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L551) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NoFileExistsf)added in v1.5.0

```
func NoFileExistsf(t TestingT, path string, msg string, args ...interface{}) bool
```

NoFileExistsf checks whether a file does not exist in a given path. It fails if the path points to an existing _file_ only.

#### func [NotContains](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L950) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotContains)

```
func NotContains(t TestingT, s, contains interface{}, msgAndArgs ...interface{}) bool
```

NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the specified substring or element.

```
assert.NotContains(t, "Hello World", "Earth")
assert.NotContains(t, ["Hello", "World"], "Earth")
assert.NotContains(t, {"Hello": "World"}, "Earth")
```

#### func [NotContainsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L564) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotContainsf)added in v1.2.0

```
func NotContainsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{}) bool
```

NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the specified substring or element.

```
assert.NotContainsf(t, "Hello World", "Earth", "error message %s", "formatted")
assert.NotContainsf(t, ["Hello", "World"], "Earth", "error message %s", "formatted")
assert.NotContainsf(t, {"Hello": "World"}, "Earth", "error message %s", "formatted")
```

#### func [NotElementsMatch](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1188) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotElementsMatch)added in v1.10.0

```
func NotElementsMatch(t TestingT, listA, listB interface{}, msgAndArgs ...interface{}) (ok bool)
```

NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should not match. This is an inverse of ElementsMatch.

assert.NotElementsMatch(t, [1, 1, 2, 3], [1, 1, 2, 3]) -> false

assert.NotElementsMatch(t, [1, 1, 2, 3], [1, 2, 3]) -> true

assert.NotElementsMatch(t, [1, 2, 3], [1, 2, 4]) -> true

#### func [NotElementsMatchf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L581) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotElementsMatchf)added in v1.10.0

```
func NotElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{}) bool
```

NotElementsMatchf asserts that the specified listA(array, slice...) is NOT equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should not match. This is an inverse of ElementsMatch.

assert.NotElementsMatchf(t, [1, 1, 2, 3], [1, 1, 2, 3], "error message %s", "formatted") -> false

assert.NotElementsMatchf(t, [1, 1, 2, 3], [1, 2, 3], "error message %s", "formatted") -> true

assert.NotElementsMatchf(t, [1, 2, 3], [1, 2, 4], "error message %s", "formatted") -> true

#### func [NotEmpty](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L768) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotEmpty)

```
func NotEmpty(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
```

NotEmpty asserts that the specified object is NOT empty. I.e. not nil, "", false, 0 or either a slice or a channel with len == 0.

```
if assert.NotEmpty(t, obj) {
  assert.Equal(t, "two", obj[1])
}
```

#### func [NotEmptyf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L594) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotEmptyf)added in v1.2.0

```
func NotEmptyf(t TestingT, object interface{}, msg string, args ...interface{}) bool
```

NotEmptyf asserts that the specified object is NOT empty. I.e. not nil, "", false, 0 or either a slice or a channel with len == 0.

```
if assert.NotEmptyf(t, obj, "error message %s", "formatted") {
  assert.Equal(t, "two", obj[1])
}
```

#### func [NotEqual](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L846) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotEqual)

```
func NotEqual(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool
```

NotEqual asserts that the specified values are NOT equal.

```
assert.NotEqual(t, obj1, obj2)
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses).

#### func [NotEqualValues](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L866) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotEqualValues)added in v1.6.0

```
func NotEqualValues(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool
```

NotEqualValues asserts that two objects are not equal even when converted to the same type

```
assert.NotEqualValues(t, obj1, obj2)
```

#### func [NotEqualValuesf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L617) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotEqualValuesf)added in v1.6.0

```
func NotEqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

NotEqualValuesf asserts that two objects are not equal even when converted to the same type

```
assert.NotEqualValuesf(t, obj1, obj2, "error message %s", "formatted")
```

#### func [NotEqualf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L607) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotEqualf)added in v1.2.0

```
func NotEqualf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

NotEqualf asserts that the specified values are NOT equal.

```
assert.NotEqualf(t, obj1, obj2, "error message %s", "formatted")
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses).

#### func [NotErrorAs](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L2156) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotErrorAs)added in v1.10.0

```
func NotErrorAs(t TestingT, err error, target interface{}, msgAndArgs ...interface{}) bool
```

NotErrorAs asserts that none of the errors in err's chain matches target, but if so, sets target to that error value.

#### func [NotErrorAsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L626) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotErrorAsf)added in v1.10.0

```
func NotErrorAsf(t TestingT, err error, target interface{}, msg string, args ...interface{}) bool
```

NotErrorAsf asserts that none of the errors in err's chain matches target, but if so, sets target to that error value.

#### func [NotErrorIs](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L2115) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotErrorIs)added in v1.7.0

```
func NotErrorIs(t TestingT, err, target error, msgAndArgs ...interface{}) bool
```

NotErrorIs asserts that none of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func [NotErrorIsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L635) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotErrorIsf)added in v1.7.0

```
func NotErrorIsf(t TestingT, err error, target error, msg string, args ...interface{}) bool
```

NotErrorIsf asserts that none of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func [NotImplements](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L424) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotImplements)added in v1.9.0

```
func NotImplements(t TestingT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool
```

NotImplements asserts that an object does not implement the specified interface.

```
assert.NotImplements(t, (*MyInterface)(nil), new(MyObject))
```

#### func [NotImplementsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L645) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotImplementsf)added in v1.9.0

```
func NotImplementsf(t TestingT, interfaceObject interface{}, object interface{}, msg string, args ...interface{}) bool
```

NotImplementsf asserts that an object does not implement the specified interface.

```
assert.NotImplementsf(t, (*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
```

#### func [NotNil](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L674) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotNil)

```
func NotNil(t TestingT, object interface{}, msgAndArgs ...interface{}) bool
```

NotNil asserts that the specified object is not nil.

```
assert.NotNil(t, err)
```

#### func [NotNilf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L655) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotNilf)added in v1.2.0

```
func NotNilf(t TestingT, object interface{}, msg string, args ...interface{}) bool
```

NotNilf asserts that the specified object is not nil.

```
assert.NotNilf(t, err, "error message %s", "formatted")
```

#### func [NotPanics](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1305) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotPanics)

```
func NotPanics(t TestingT, f PanicTestFunc, msgAndArgs ...interface{}) bool
```

NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.

```
assert.NotPanics(t, func(){ RemainCalm() })
```

#### func [NotPanicsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L665) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotPanicsf)added in v1.2.0

```
func NotPanicsf(t TestingT, f PanicTestFunc, msg string, args ...interface{}) bool
```

NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.

```
assert.NotPanicsf(t, func(){ RemainCalm() }, "error message %s", "formatted")
```

#### func [NotRegexp](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1695) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotRegexp)

```
func NotRegexp(t TestingT, rx interface{}, str interface{}, msgAndArgs ...interface{}) bool
```

NotRegexp asserts that a specified regexp does not match a string.

```
assert.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
assert.NotRegexp(t, "^start", "it's not starting")
```

#### func [NotRegexpf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L676) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotRegexpf)added in v1.2.0

```
func NotRegexpf(t TestingT, rx interface{}, str interface{}, msg string, args ...interface{}) bool
```

NotRegexpf asserts that a specified regexp does not match a string.

```
assert.NotRegexpf(t, regexp.MustCompile("starts"), "it's starting", "error message %s", "formatted")
assert.NotRegexpf(t, "^start", "it's not starting", "error message %s", "formatted")
```

#### func [NotSame](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L526) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotSame)added in v1.5.0

```
func NotSame(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool
```

NotSame asserts that two pointers do not reference the same object.

```
assert.NotSame(t, ptr1, ptr2)
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func [NotSamef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L689) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotSamef)added in v1.5.0

```
func NotSamef(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

NotSamef asserts that two pointers do not reference the same object.

```
assert.NotSamef(t, ptr1, ptr2, "error message %s", "formatted")
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func [NotSubset](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1030) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotSubset)added in v1.2.0

```
func NotSubset(t TestingT, list, subset interface{}, msgAndArgs ...interface{}) (ok bool)
```

NotSubset asserts that the specified list(array, slice...) or map does NOT contain all elements given in the specified subset list(array, slice...) or map.

```
assert.NotSubset(t, [1, 3, 4], [1, 2])
assert.NotSubset(t, {"x": 1, "y": 2}, {"z": 3})
```

#### func [NotSubsetf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L702) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotSubsetf)added in v1.2.0

```
func NotSubsetf(t TestingT, list interface{}, subset interface{}, msg string, args ...interface{}) bool
```

NotSubsetf asserts that the specified list(array, slice...) or map does NOT contain all elements given in the specified subset list(array, slice...) or map.

```
assert.NotSubsetf(t, [1, 3, 4], [1, 2], "error message %s", "formatted")
assert.NotSubsetf(t, {"x": 1, "y": 2}, {"z": 3}, "error message %s", "formatted")
```

#### func [NotZero](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1721) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotZero)

```
func NotZero(t TestingT, i interface{}, msgAndArgs ...interface{}) bool
```

NotZero asserts that i is not the zero value for its type.

#### func [NotZerof](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L710) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#NotZerof)added in v1.2.0

```
func NotZerof(t TestingT, i interface{}, msg string, args ...interface{}) bool
```

NotZerof asserts that i is not the zero value for its type.

#### func [ObjectsAreEqual](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L64) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ObjectsAreEqual)

```
func ObjectsAreEqual(expected, actual interface{}) bool
```

ObjectsAreEqual determines if two objects are considered equal.

This function does no assertion of any kind.

#### func [ObjectsAreEqualValues](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L164) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ObjectsAreEqualValues)

```
func ObjectsAreEqualValues(expected, actual interface{}) bool
```

ObjectsAreEqualValues gets whether two objects are equal, or if their values are equal.

<details class="Documentation-deprecatedDetails js-deprecatedDetails" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; display: block; color: var(--color-text-subtle);"><summary style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; list-style: none; opacity: 1;"><h4 tabindex="-1" id="ObjectsExportedFieldsAreEqual" data-kind="function" class="Documentation-functionHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 600; font-stretch: inherit; line-height: 1.25em; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1.125rem; margin: 1.5rem 0px 0.5rem; padding: 0px; vertical-align: baseline; word-break: break-word; align-items: baseline; display: flex; justify-content: space-between;"><span class="Documentation-deprecatedTitle" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; align-items: center; display: flex; gap: 0.5rem;">func<a class="Documentation-source" href="https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L156" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 1;">ObjectsExportedFieldsAreEqual</a><span class="Documentation-deprecatedTag" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: 1.375; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.75rem; margin: 0px; padding: 0.125rem 0.25rem; vertical-align: middle; background-color: var(--color-border); border-radius: 0.125rem; color: var(--color-text-inverted); text-transform: uppercase;">deprecated</span><span class="Documentation-deprecatedBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.87rem; margin: 0px 0.5rem 0px 0.25rem; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></span><span class="Documentation-sinceVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.9375rem; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"><span class="Documentation-sinceVersionLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 15px; margin: 0px; padding: 0px; vertical-align: baseline;">added in</span><span>&nbsp;</span><span class="Documentation-sinceVersionVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 15px; margin: 0px; padding: 0px; vertical-align: baseline;">v1.8.3</span></span></h4></summary><div class="go-Message go-Message--warning Documentation-deprecatedItemBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 1rem 1rem 0.5rem; vertical-align: baseline; color: var(--gray-1); width: 1263.31px; background-color: var(--color-background-warning);"><div class="Documentation-declaration" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre-wrap; scroll-padding-top: calc(var(--js-sticky-header-height, 3.5rem) + .75rem); word-break: break-all; overflow-wrap: break-word;"><a href="https://pkg.go.dev/builtin#bool" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></pre></div><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"><a href="https://pkg.go.dev/github.com/stretchr/testify/assert#EqualExportedValues" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></p></div></details>

#### func [Panics](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1248) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Panics)

```
func Panics(t TestingT, f PanicTestFunc, msgAndArgs ...interface{}) bool
```

Panics asserts that the code inside the specified PanicTestFunc panics.

```
assert.Panics(t, func(){ GoCrazy() })
```

#### func [PanicsWithError](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1285) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#PanicsWithError)added in v1.5.0

```
func PanicsWithError(t TestingT, errString string, f PanicTestFunc, msgAndArgs ...interface{}) bool
```

PanicsWithError asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value is an error that satisfies the EqualError comparison.

```
assert.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
```

#### func [PanicsWithErrorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L732) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#PanicsWithErrorf)added in v1.5.0

```
func PanicsWithErrorf(t TestingT, errString string, f PanicTestFunc, msg string, args ...interface{}) bool
```

PanicsWithErrorf asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value is an error that satisfies the EqualError comparison.

```
assert.PanicsWithErrorf(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func [PanicsWithValue](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1264) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#PanicsWithValue)added in v1.2.0

```
func PanicsWithValue(t TestingT, expected interface{}, f PanicTestFunc, msgAndArgs ...interface{}) bool
```

PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value equals the expected panic value.

```
assert.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
```

#### func [PanicsWithValuef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L743) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#PanicsWithValuef)added in v1.2.0

```
func PanicsWithValuef(t TestingT, expected interface{}, f PanicTestFunc, msg string, args ...interface{}) bool
```

PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value equals the expected panic value.

```
assert.PanicsWithValuef(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func [Panicsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L720) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Panicsf)added in v1.2.0

```
func Panicsf(t TestingT, f PanicTestFunc, msg string, args ...interface{}) bool
```

Panicsf asserts that the code inside the specified PanicTestFunc panics.

```
assert.Panicsf(t, func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func [Positive](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_compare.go#L438) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Positive)added in v1.7.0

```
func Positive(t TestingT, e interface{}, msgAndArgs ...interface{}) bool
```

Positive asserts that the specified element is positive

```
assert.Positive(t, 1)
assert.Positive(t, 1.23)
```

#### func [Positivef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L754) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Positivef)added in v1.7.0

```
func Positivef(t TestingT, e interface{}, msg string, args ...interface{}) bool
```

Positivef asserts that the specified element is positive

```
assert.Positivef(t, 1, "error message %s", "formatted")
assert.Positivef(t, 1.23, "error message %s", "formatted")
```

#### func [Regexp](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1677) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Regexp)

```
func Regexp(t TestingT, rx interface{}, str interface{}, msgAndArgs ...interface{}) bool
```

Regexp asserts that a specified regexp matches a string.

```
assert.Regexp(t, regexp.MustCompile("start"), "it's starting")
assert.Regexp(t, "start...$", "it's not starting")
```

#### func [Regexpf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L765) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Regexpf)added in v1.2.0

```
func Regexpf(t TestingT, rx interface{}, str interface{}, msg string, args ...interface{}) bool
```

Regexpf asserts that a specified regexp matches a string.

```
assert.Regexpf(t, regexp.MustCompile("start"), "it's starting", "error message %s", "formatted")
assert.Regexpf(t, "start...$", "it's not starting", "error message %s", "formatted")
```

#### func [Same](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L500) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Same)added in v1.4.0

```
func Same(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool
```

Same asserts that two pointers reference the same object.

```
assert.Same(t, ptr1, ptr2)
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func [Samef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L778) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Samef)added in v1.4.0

```
func Samef(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

Samef asserts that two pointers reference the same object.

```
assert.Samef(t, ptr1, ptr2, "error message %s", "formatted")
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func [Subset](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L972) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Subset)added in v1.2.0

```
func Subset(t TestingT, list, subset interface{}, msgAndArgs ...interface{}) (ok bool)
```

Subset asserts that the specified list(array, slice...) or map contains all elements given in the specified subset list(array, slice...) or map.

```
assert.Subset(t, [1, 2, 3], [1, 2])
assert.Subset(t, {"x": 1, "y": 2}, {"x": 1})
```

#### func [Subsetf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L790) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Subsetf)added in v1.2.0

```
func Subsetf(t TestingT, list interface{}, subset interface{}, msg string, args ...interface{}) bool
```

Subsetf asserts that the specified list(array, slice...) or map contains all elements given in the specified subset list(array, slice...) or map.

```
assert.Subsetf(t, [1, 2, 3], [1, 2], "error message %s", "formatted")
assert.Subsetf(t, {"x": 1, "y": 2}, {"x": 1}, "error message %s", "formatted")
```

#### func [True](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L813) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#True)

```
func True(t TestingT, value bool, msgAndArgs ...interface{}) bool
```

True asserts that the specified value is true.

```
assert.True(t, myBool)
```

#### func [Truef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L800) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Truef)added in v1.2.0

```
func Truef(t TestingT, value bool, msg string, args ...interface{}) bool
```

Truef asserts that the specified value is true.

```
assert.Truef(t, myBool, "error message %s", "formatted")
```

#### func [WithinDuration](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1320) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#WithinDuration)

```
func WithinDuration(t TestingT, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool
```

WithinDuration asserts that the two times are within duration delta of each other.

```
assert.WithinDuration(t, time.Now(), time.Now(), 10*time.Second)
```

#### func [WithinDurationf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L810) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#WithinDurationf)added in v1.2.0

```
func WithinDurationf(t TestingT, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) bool
```

WithinDurationf asserts that the two times are within duration delta of each other.

```
assert.WithinDurationf(t, time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
```

#### func [WithinRange](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1336) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#WithinRange)added in v1.8.0

```
func WithinRange(t TestingT, actual, start, end time.Time, msgAndArgs ...interface{}) bool
```

WithinRange asserts that a time is within a time range (inclusive).

```
assert.WithinRange(t, time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
```

#### func [WithinRangef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L820) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#WithinRangef)added in v1.8.0

```
func WithinRangef(t TestingT, actual time.Time, start time.Time, end time.Time, msg string, args ...interface{}) bool
```

WithinRangef asserts that a time is within a time range (inclusive).

```
assert.WithinRangef(t, time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second), "error message %s", "formatted")
```

#### func [YAMLEq](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1825) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#YAMLEq)added in v1.4.0

```
func YAMLEq(t TestingT, expected string, actual string, msgAndArgs ...interface{}) bool
```

YAMLEq asserts that two YAML strings are equivalent.

#### func [YAMLEqf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L828) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#YAMLEqf)added in v1.4.0

```
func YAMLEqf(t TestingT, expected string, actual string, msg string, args ...interface{}) bool
```

YAMLEqf asserts that two YAML strings are equivalent.

#### func [Zero](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1710) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Zero)

```
func Zero(t TestingT, i interface{}, msgAndArgs ...interface{}) bool
```

Zero asserts that i is the zero value for its type.

#### func [Zerof](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_format.go#L836) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Zerof)added in v1.2.0

```
func Zerof(t TestingT, i interface{}, msg string, args ...interface{}) bool
```

Zerof asserts that i is the zero value for its type.

### Types [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#pkg-types)

#### type [Assertions](https://github.com/stretchr/testify/blob/v1.10.0/assert/forward_assertions.go#L5) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions)

```
type Assertions struct {
	// contains filtered or unexported fields
}
```

Assertions provides assertion methods around the TestingT interface.

#### func [New](https://github.com/stretchr/testify/blob/v1.10.0/assert/forward_assertions.go#L10) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#New)

```
func New(t TestingT) *Assertions
```

New makes a new Assertions object for the specified TestingT.

#### func (*Assertions) [Condition](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L12) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Condition)

```
func (a *Assertions) Condition(comp Comparison, msgAndArgs ...interface{}) bool
```

Condition uses a Comparison to assert a complex condition.

#### func (*Assertions) [Conditionf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L20) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Conditionf)added in v1.2.0

```
func (a *Assertions) Conditionf(comp Comparison, msg string, args ...interface{}) bool
```

Conditionf uses a Comparison to assert a complex condition.

#### func (*Assertions) [Contains](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L33) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Contains)

```
func (a *Assertions) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) bool
```

Contains asserts that the specified string, list(array, slice...) or map contains the specified substring or element.

```
a.Contains("Hello World", "World")
a.Contains(["Hello", "World"], "World")
a.Contains({"Hello": "World"}, "Hello")
```

#### func (*Assertions) [Containsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L46) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Containsf)added in v1.2.0

```
func (a *Assertions) Containsf(s interface{}, contains interface{}, msg string, args ...interface{}) bool
```

Containsf asserts that the specified string, list(array, slice...) or map contains the specified substring or element.

```
a.Containsf("Hello World", "World", "error message %s", "formatted")
a.Containsf(["Hello", "World"], "World", "error message %s", "formatted")
a.Containsf({"Hello": "World"}, "Hello", "error message %s", "formatted")
```

#### func (*Assertions) [DirExists](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L55) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.DirExists)added in v1.2.0

```
func (a *Assertions) DirExists(path string, msgAndArgs ...interface{}) bool
```

DirExists checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.

#### func (*Assertions) [DirExistsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L64) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.DirExistsf)added in v1.2.0

```
func (a *Assertions) DirExistsf(path string, msg string, args ...interface{}) bool
```

DirExistsf checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.

#### func (*Assertions) [ElementsMatch](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L76) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.ElementsMatch)added in v1.2.0

```
func (a *Assertions) ElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) bool
```

ElementsMatch asserts that the specified listA(array, slice...) is equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should match.

a.ElementsMatch([1, 3, 2, 3], [1, 3, 3, 2])

#### func (*Assertions) [ElementsMatchf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L88) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.ElementsMatchf)added in v1.2.0

```
func (a *Assertions) ElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{}) bool
```

ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should match.

a.ElementsMatchf([1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")

#### func (*Assertions) [Empty](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L99) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Empty)

```
func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{}) bool
```

Empty asserts that the specified object is empty. I.e. nil, "", false, 0 or either a slice or a channel with len == 0.

```
a.Empty(obj)
```

#### func (*Assertions) [Emptyf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L110) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Emptyf)added in v1.2.0

```
func (a *Assertions) Emptyf(object interface{}, msg string, args ...interface{}) bool
```

Emptyf asserts that the specified object is empty. I.e. nil, "", false, 0 or either a slice or a channel with len == 0.

```
a.Emptyf(obj, "error message %s", "formatted")
```

#### func (*Assertions) [Equal](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L124) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Equal)

```
func (a *Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```

Equal asserts that two objects are equal.

```
a.Equal(123, 123)
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses). Function equality cannot be determined and will always fail.

#### func (*Assertions) [EqualError](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L136) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.EqualError)

```
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{}) bool
```

EqualError asserts that a function returned an error (i.e. not `nil`) and that it is equal to the provided error.

```
actualObj, err := SomeFunction()
a.EqualError(err,  expectedErrorString)
```

#### func (*Assertions) [EqualErrorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L148) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.EqualErrorf)added in v1.2.0

```
func (a *Assertions) EqualErrorf(theError error, errString string, msg string, args ...interface{}) bool
```

EqualErrorf asserts that a function returned an error (i.e. not `nil`) and that it is equal to the provided error.

```
actualObj, err := SomeFunction()
a.EqualErrorf(err,  expectedErrorString, "error message %s", "formatted")
```

#### func (*Assertions) [EqualExportedValues](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L165) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.EqualExportedValues)added in v1.8.3

```
func (a *Assertions) EqualExportedValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```

EqualExportedValues asserts that the types of two objects are equal and their public fields are also equal. This is useful for comparing structs that have private fields that could potentially differ.

```
 type S struct {
	Exported     	int
	notExported   	int
 }
 a.EqualExportedValues(S{1, 2}, S{1, 3}) => true
 a.EqualExportedValues(S{1, 2}, S{2, 3}) => false
```

#### func (*Assertions) [EqualExportedValuesf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L182) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.EqualExportedValuesf)added in v1.8.3

```
func (a *Assertions) EqualExportedValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

EqualExportedValuesf asserts that the types of two objects are equal and their public fields are also equal. This is useful for comparing structs that have private fields that could potentially differ.

```
 type S struct {
	Exported     	int
	notExported   	int
 }
 a.EqualExportedValuesf(S{1, 2}, S{1, 3}, "error message %s", "formatted") => true
 a.EqualExportedValuesf(S{1, 2}, S{2, 3}, "error message %s", "formatted") => false
```

#### func (*Assertions) [EqualValues](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L193) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.EqualValues)

```
func (a *Assertions) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```

EqualValues asserts that two objects are equal or convertible to the larger type and equal.

```
a.EqualValues(uint32(123), int32(123))
```

#### func (*Assertions) [EqualValuesf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L204) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.EqualValuesf)added in v1.2.0

```
func (a *Assertions) EqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

EqualValuesf asserts that two objects are equal or convertible to the larger type and equal.

```
a.EqualValuesf(uint32(123), int32(123), "error message %s", "formatted")
```

#### func (*Assertions) [Equalf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L218) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Equalf)added in v1.2.0

```
func (a *Assertions) Equalf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

Equalf asserts that two objects are equal.

```
a.Equalf(123, 123, "error message %s", "formatted")
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses). Function equality cannot be determined and will always fail.

#### func (*Assertions) [Error](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L231) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Error)

```
func (a *Assertions) Error(err error, msgAndArgs ...interface{}) bool
```

Error asserts that a function returned an error (i.e. not `nil`).

```
  actualObj, err := SomeFunction()
  if a.Error(err) {
	   assert.Equal(t, expectedError, err)
  }
```

#### func (*Assertions) [ErrorAs](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L240) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.ErrorAs)added in v1.7.0

```
func (a *Assertions) ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) bool
```

ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value. This is a wrapper for errors.As.

#### func (*Assertions) [ErrorAsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L249) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.ErrorAsf)added in v1.7.0

```
func (a *Assertions) ErrorAsf(err error, target interface{}, msg string, args ...interface{}) bool
```

ErrorAsf asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value. This is a wrapper for errors.As.

#### func (*Assertions) [ErrorContains](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L261) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.ErrorContains)added in v1.7.1

```
func (a *Assertions) ErrorContains(theError error, contains string, msgAndArgs ...interface{}) bool
```

ErrorContains asserts that a function returned an error (i.e. not `nil`) and that the error contains the specified substring.

```
actualObj, err := SomeFunction()
a.ErrorContains(err,  expectedErrorSubString)
```

#### func (*Assertions) [ErrorContainsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L273) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.ErrorContainsf)added in v1.7.1

```
func (a *Assertions) ErrorContainsf(theError error, contains string, msg string, args ...interface{}) bool
```

ErrorContainsf asserts that a function returned an error (i.e. not `nil`) and that the error contains the specified substring.

```
actualObj, err := SomeFunction()
a.ErrorContainsf(err,  expectedErrorSubString, "error message %s", "formatted")
```

#### func (*Assertions) [ErrorIs](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L282) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.ErrorIs)added in v1.7.0

```
func (a *Assertions) ErrorIs(err error, target error, msgAndArgs ...interface{}) bool
```

ErrorIs asserts that at least one of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func (*Assertions) [ErrorIsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L291) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.ErrorIsf)added in v1.7.0

```
func (a *Assertions) ErrorIsf(err error, target error, msg string, args ...interface{}) bool
```

ErrorIsf asserts that at least one of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func (*Assertions) [Errorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L304) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Errorf)added in v1.2.0

```
func (a *Assertions) Errorf(err error, msg string, args ...interface{}) bool
```

Errorf asserts that a function returned an error (i.e. not `nil`).

```
  actualObj, err := SomeFunction()
  if a.Errorf(err, "error message %s", "formatted") {
	   assert.Equal(t, expectedErrorf, err)
  }
```

#### func (*Assertions) [Eventually](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L315) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Eventually)added in v1.4.0

```
func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool
```

Eventually asserts that given condition will be met in waitFor time, periodically checking target function each tick.

```
a.Eventually(func() bool { return true; }, time.Second, 10*time.Millisecond)
```

#### func (*Assertions) [EventuallyWithT](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L340) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.EventuallyWithT)added in v1.8.3

```
func (a *Assertions) EventuallyWithT(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool
```

EventuallyWithT asserts that given condition will be met in waitFor time, periodically checking target function each tick. In contrast to Eventually, it supplies a CollectT to the condition function, so that the condition function can use the CollectT to call other assertions. The condition is considered "met" if no errors are raised in a tick. The supplied CollectT collects all errors from one tick (if there are any). If the condition is not met before waitFor, the collected errors of the last tick are copied to t.

```
externalValue := false
go func() {
	time.Sleep(8*time.Second)
	externalValue = true
}()
a.EventuallyWithT(func(c *assert.CollectT) {
	// add assertions as needed; any assertion failure will fail the current tick
	assert.True(c, externalValue, "expected 'externalValue' to be true")
}, 10*time.Second, 1*time.Second, "external state has not changed to 'true'; still false")
```

#### func (*Assertions) [EventuallyWithTf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L365) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.EventuallyWithTf)added in v1.8.3

```
func (a *Assertions) EventuallyWithTf(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool
```

EventuallyWithTf asserts that given condition will be met in waitFor time, periodically checking target function each tick. In contrast to Eventually, it supplies a CollectT to the condition function, so that the condition function can use the CollectT to call other assertions. The condition is considered "met" if no errors are raised in a tick. The supplied CollectT collects all errors from one tick (if there are any). If the condition is not met before waitFor, the collected errors of the last tick are copied to t.

```
externalValue := false
go func() {
	time.Sleep(8*time.Second)
	externalValue = true
}()
a.EventuallyWithTf(func(c *assert.CollectT, "error message %s", "formatted") {
	// add assertions as needed; any assertion failure will fail the current tick
	assert.True(c, externalValue, "expected 'externalValue' to be true")
}, 10*time.Second, 1*time.Second, "external state has not changed to 'true'; still false")
```

#### func (*Assertions) [Eventuallyf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L376) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Eventuallyf)added in v1.4.0

```
func (a *Assertions) Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool
```

Eventuallyf asserts that given condition will be met in waitFor time, periodically checking target function each tick.

```
a.Eventuallyf(func() bool { return true; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
```

#### func (*Assertions) [Exactly](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L386) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Exactly)

```
func (a *Assertions) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```

Exactly asserts that two objects are equal in value and type.

```
a.Exactly(int32(123), int64(123))
```

#### func (*Assertions) [Exactlyf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L396) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Exactlyf)added in v1.2.0

```
func (a *Assertions) Exactlyf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

Exactlyf asserts that two objects are equal in value and type.

```
a.Exactlyf(int32(123), int64(123), "error message %s", "formatted")
```

#### func (*Assertions) [Fail](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L404) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Fail)

```
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface{}) bool
```

Fail reports a failure through

#### func (*Assertions) [FailNow](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L412) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.FailNow)

```
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...interface{}) bool
```

FailNow fails test

#### func (*Assertions) [FailNowf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L420) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.FailNowf)added in v1.2.0

```
func (a *Assertions) FailNowf(failureMessage string, msg string, args ...interface{}) bool
```

FailNowf fails test

#### func (*Assertions) [Failf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L428) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Failf)added in v1.2.0

```
func (a *Assertions) Failf(failureMessage string, msg string, args ...interface{}) bool
```

Failf reports a failure through

#### func (*Assertions) [False](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L438) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.False)

```
func (a *Assertions) False(value bool, msgAndArgs ...interface{}) bool
```

False asserts that the specified value is false.

```
a.False(myBool)
```

#### func (*Assertions) [Falsef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L448) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Falsef)added in v1.2.0

```
func (a *Assertions) Falsef(value bool, msg string, args ...interface{}) bool
```

Falsef asserts that the specified value is false.

```
a.Falsef(myBool, "error message %s", "formatted")
```

#### func (*Assertions) [FileExists](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L457) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.FileExists)added in v1.2.0

```
func (a *Assertions) FileExists(path string, msgAndArgs ...interface{}) bool
```

FileExists checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.

#### func (*Assertions) [FileExistsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L466) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.FileExistsf)added in v1.2.0

```
func (a *Assertions) FileExistsf(path string, msg string, args ...interface{}) bool
```

FileExistsf checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.

#### func (*Assertions) [Greater](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L478) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Greater)added in v1.4.0

```
func (a *Assertions) Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool
```

Greater asserts that the first element is greater than the second

```
a.Greater(2, 1)
a.Greater(float64(2), float64(1))
a.Greater("b", "a")
```

#### func (*Assertions) [GreaterOrEqual](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L491) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.GreaterOrEqual)added in v1.4.0

```
func (a *Assertions) GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool
```

GreaterOrEqual asserts that the first element is greater than or equal to the second

```
a.GreaterOrEqual(2, 1)
a.GreaterOrEqual(2, 2)
a.GreaterOrEqual("b", "a")
a.GreaterOrEqual("b", "b")
```

#### func (*Assertions) [GreaterOrEqualf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L504) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.GreaterOrEqualf)added in v1.4.0

```
func (a *Assertions) GreaterOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool
```

GreaterOrEqualf asserts that the first element is greater than or equal to the second

```
a.GreaterOrEqualf(2, 1, "error message %s", "formatted")
a.GreaterOrEqualf(2, 2, "error message %s", "formatted")
a.GreaterOrEqualf("b", "a", "error message %s", "formatted")
a.GreaterOrEqualf("b", "b", "error message %s", "formatted")
```

#### func (*Assertions) [Greaterf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L516) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Greaterf)added in v1.4.0

```
func (a *Assertions) Greaterf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool
```

Greaterf asserts that the first element is greater than the second

```
a.Greaterf(2, 1, "error message %s", "formatted")
a.Greaterf(float64(2), float64(1), "error message %s", "formatted")
a.Greaterf("b", "a", "error message %s", "formatted")
```

#### func (*Assertions) [HTTPBodyContains](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L529) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPBodyContains)

```
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool
```

HTTPBodyContains asserts that a specified handler returns a body that contains a string.

```
a.HTTPBodyContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPBodyContainsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L542) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPBodyContainsf)added in v1.2.0

```
func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) bool
```

HTTPBodyContainsf asserts that a specified handler returns a body that contains a string.

```
a.HTTPBodyContainsf(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPBodyNotContains](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L555) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPBodyNotContains)

```
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool
```

HTTPBodyNotContains asserts that a specified handler returns a body that does not contain a string.

```
a.HTTPBodyNotContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPBodyNotContainsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L568) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPBodyNotContainsf)added in v1.2.0

```
func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) bool
```

HTTPBodyNotContainsf asserts that a specified handler returns a body that does not contain a string.

```
a.HTTPBodyNotContainsf(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPError](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L580) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPError)

```
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) bool
```

HTTPError asserts that a specified handler returns an error status code.

```
a.HTTPError(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPErrorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L592) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPErrorf)added in v1.2.0

```
func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool
```

HTTPErrorf asserts that a specified handler returns an error status code.

```
a.HTTPErrorf(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPRedirect](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L604) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPRedirect)

```
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) bool
```

HTTPRedirect asserts that a specified handler returns a redirect status code.

```
a.HTTPRedirect(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPRedirectf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L616) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPRedirectf)added in v1.2.0

```
func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool
```

HTTPRedirectf asserts that a specified handler returns a redirect status code.

```
a.HTTPRedirectf(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPStatusCode](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L628) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPStatusCode)added in v1.6.0

```
func (a *Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) bool
```

HTTPStatusCode asserts that a specified handler returns a specified status code.

```
a.HTTPStatusCode(myHandler, "GET", "/notImplemented", nil, 501)
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPStatusCodef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L640) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPStatusCodef)added in v1.6.0

```
func (a *Assertions) HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{}) bool
```

HTTPStatusCodef asserts that a specified handler returns a specified status code.

```
a.HTTPStatusCodef(myHandler, "GET", "/notImplemented", nil, 501, "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPSuccess](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L652) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPSuccess)

```
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{}) bool
```

HTTPSuccess asserts that a specified handler returns a success status code.

```
a.HTTPSuccess(myHandler, "POST", "http://www.google.com", nil)
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPSuccessf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L664) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.HTTPSuccessf)added in v1.2.0

```
func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool
```

HTTPSuccessf asserts that a specified handler returns a success status code.

```
a.HTTPSuccessf(myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [Implements](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L674) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Implements)

```
func (a *Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool
```

Implements asserts that an object is implemented by the specified interface.

```
a.Implements((*MyInterface)(nil), new(MyObject))
```

#### func (*Assertions) [Implementsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L684) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Implementsf)added in v1.2.0

```
func (a *Assertions) Implementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) bool
```

Implementsf asserts that an object is implemented by the specified interface.

```
a.Implementsf((*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
```

#### func (*Assertions) [InDelta](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L694) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.InDelta)

```
func (a *Assertions) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) bool
```

InDelta asserts that the two numerals are within delta of each other.

```
a.InDelta(math.Pi, 22/7.0, 0.01)
```

#### func (*Assertions) [InDeltaMapValues](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L702) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.InDeltaMapValues)added in v1.2.0

```
func (a *Assertions) InDeltaMapValues(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) bool
```

InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.

#### func (*Assertions) [InDeltaMapValuesf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L710) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.InDeltaMapValuesf)added in v1.2.0

```
func (a *Assertions) InDeltaMapValuesf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool
```

InDeltaMapValuesf is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.

#### func (*Assertions) [InDeltaSlice](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L718) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.InDeltaSlice)

```
func (a *Assertions) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{}) bool
```

InDeltaSlice is the same as InDelta, except it compares two slices.

#### func (*Assertions) [InDeltaSlicef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L726) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.InDeltaSlicef)added in v1.2.0

```
func (a *Assertions) InDeltaSlicef(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool
```

InDeltaSlicef is the same as InDelta, except it compares two slices.

#### func (*Assertions) [InDeltaf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L736) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.InDeltaf)added in v1.2.0

```
func (a *Assertions) InDeltaf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool
```

InDeltaf asserts that the two numerals are within delta of each other.

```
a.InDeltaf(math.Pi, 22/7.0, 0.01, "error message %s", "formatted")
```

#### func (*Assertions) [InEpsilon](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L744) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.InEpsilon)

```
func (a *Assertions) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool
```

InEpsilon asserts that expected and actual have a relative error less than epsilon

#### func (*Assertions) [InEpsilonSlice](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L752) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.InEpsilonSlice)

```
func (a *Assertions) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool
```

InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.

#### func (*Assertions) [InEpsilonSlicef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L760) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.InEpsilonSlicef)added in v1.2.0

```
func (a *Assertions) InEpsilonSlicef(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool
```

InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.

#### func (*Assertions) [InEpsilonf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L768) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.InEpsilonf)added in v1.2.0

```
func (a *Assertions) InEpsilonf(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool
```

InEpsilonf asserts that expected and actual have a relative error less than epsilon

#### func (*Assertions) [IsDecreasing](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L780) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.IsDecreasing)added in v1.7.0

```
func (a *Assertions) IsDecreasing(object interface{}, msgAndArgs ...interface{}) bool
```

IsDecreasing asserts that the collection is decreasing

```
a.IsDecreasing([]int{2, 1, 0})
a.IsDecreasing([]float{2, 1})
a.IsDecreasing([]string{"b", "a"})
```

#### func (*Assertions) [IsDecreasingf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L792) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.IsDecreasingf)added in v1.7.0

```
func (a *Assertions) IsDecreasingf(object interface{}, msg string, args ...interface{}) bool
```

IsDecreasingf asserts that the collection is decreasing

```
a.IsDecreasingf([]int{2, 1, 0}, "error message %s", "formatted")
a.IsDecreasingf([]float{2, 1}, "error message %s", "formatted")
a.IsDecreasingf([]string{"b", "a"}, "error message %s", "formatted")
```

#### func (*Assertions) [IsIncreasing](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L804) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.IsIncreasing)added in v1.7.0

```
func (a *Assertions) IsIncreasing(object interface{}, msgAndArgs ...interface{}) bool
```

IsIncreasing asserts that the collection is increasing

```
a.IsIncreasing([]int{1, 2, 3})
a.IsIncreasing([]float{1, 2})
a.IsIncreasing([]string{"a", "b"})
```

#### func (*Assertions) [IsIncreasingf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L816) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.IsIncreasingf)added in v1.7.0

```
func (a *Assertions) IsIncreasingf(object interface{}, msg string, args ...interface{}) bool
```

IsIncreasingf asserts that the collection is increasing

```
a.IsIncreasingf([]int{1, 2, 3}, "error message %s", "formatted")
a.IsIncreasingf([]float{1, 2}, "error message %s", "formatted")
a.IsIncreasingf([]string{"a", "b"}, "error message %s", "formatted")
```

#### func (*Assertions) [IsNonDecreasing](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L828) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.IsNonDecreasing)added in v1.7.0

```
func (a *Assertions) IsNonDecreasing(object interface{}, msgAndArgs ...interface{}) bool
```

IsNonDecreasing asserts that the collection is not decreasing

```
a.IsNonDecreasing([]int{1, 1, 2})
a.IsNonDecreasing([]float{1, 2})
a.IsNonDecreasing([]string{"a", "b"})
```

#### func (*Assertions) [IsNonDecreasingf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L840) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.IsNonDecreasingf)added in v1.7.0

```
func (a *Assertions) IsNonDecreasingf(object interface{}, msg string, args ...interface{}) bool
```

IsNonDecreasingf asserts that the collection is not decreasing

```
a.IsNonDecreasingf([]int{1, 1, 2}, "error message %s", "formatted")
a.IsNonDecreasingf([]float{1, 2}, "error message %s", "formatted")
a.IsNonDecreasingf([]string{"a", "b"}, "error message %s", "formatted")
```

#### func (*Assertions) [IsNonIncreasing](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L852) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.IsNonIncreasing)added in v1.7.0

```
func (a *Assertions) IsNonIncreasing(object interface{}, msgAndArgs ...interface{}) bool
```

IsNonIncreasing asserts that the collection is not increasing

```
a.IsNonIncreasing([]int{2, 1, 1})
a.IsNonIncreasing([]float{2, 1})
a.IsNonIncreasing([]string{"b", "a"})
```

#### func (*Assertions) [IsNonIncreasingf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L864) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.IsNonIncreasingf)added in v1.7.0

```
func (a *Assertions) IsNonIncreasingf(object interface{}, msg string, args ...interface{}) bool
```

IsNonIncreasingf asserts that the collection is not increasing

```
a.IsNonIncreasingf([]int{2, 1, 1}, "error message %s", "formatted")
a.IsNonIncreasingf([]float{2, 1}, "error message %s", "formatted")
a.IsNonIncreasingf([]string{"b", "a"}, "error message %s", "formatted")
```

#### func (*Assertions) [IsType](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L872) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.IsType)

```
func (a *Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool
```

IsType asserts that the specified objects are of the same type.

#### func (*Assertions) [IsTypef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L880) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.IsTypef)added in v1.2.0

```
func (a *Assertions) IsTypef(expectedType interface{}, object interface{}, msg string, args ...interface{}) bool
```

IsTypef asserts that the specified objects are of the same type.

#### func (*Assertions) [JSONEq](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L890) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.JSONEq)

```
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{}) bool
```

JSONEq asserts that two JSON strings are equivalent.

```
a.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
```

#### func (*Assertions) [JSONEqf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L900) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.JSONEqf)added in v1.2.0

```
func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...interface{}) bool
```

JSONEqf asserts that two JSON strings are equivalent.

```
a.JSONEqf(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
```

#### func (*Assertions) [Len](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L911) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Len)

```
func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{}) bool
```

Len asserts that the specified object has specific length. Len also fails if the object has a type that len() not accept.

```
a.Len(mySlice, 3)
```

#### func (*Assertions) [Lenf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L922) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Lenf)added in v1.2.0

```
func (a *Assertions) Lenf(object interface{}, length int, msg string, args ...interface{}) bool
```

Lenf asserts that the specified object has specific length. Lenf also fails if the object has a type that len() not accept.

```
a.Lenf(mySlice, 3, "error message %s", "formatted")
```

#### func (*Assertions) [Less](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L934) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Less)added in v1.4.0

```
func (a *Assertions) Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool
```

Less asserts that the first element is less than the second

```
a.Less(1, 2)
a.Less(float64(1), float64(2))
a.Less("a", "b")
```

#### func (*Assertions) [LessOrEqual](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L947) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.LessOrEqual)added in v1.4.0

```
func (a *Assertions) LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool
```

LessOrEqual asserts that the first element is less than or equal to the second

```
a.LessOrEqual(1, 2)
a.LessOrEqual(2, 2)
a.LessOrEqual("a", "b")
a.LessOrEqual("b", "b")
```

#### func (*Assertions) [LessOrEqualf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L960) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.LessOrEqualf)added in v1.4.0

```
func (a *Assertions) LessOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool
```

LessOrEqualf asserts that the first element is less than or equal to the second

```
a.LessOrEqualf(1, 2, "error message %s", "formatted")
a.LessOrEqualf(2, 2, "error message %s", "formatted")
a.LessOrEqualf("a", "b", "error message %s", "formatted")
a.LessOrEqualf("b", "b", "error message %s", "formatted")
```

#### func (*Assertions) [Lessf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L972) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Lessf)added in v1.4.0

```
func (a *Assertions) Lessf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool
```

Lessf asserts that the first element is less than the second

```
a.Lessf(1, 2, "error message %s", "formatted")
a.Lessf(float64(1), float64(2), "error message %s", "formatted")
a.Lessf("a", "b", "error message %s", "formatted")
```

#### func (*Assertions) [Negative](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L983) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Negative)added in v1.7.0

```
func (a *Assertions) Negative(e interface{}, msgAndArgs ...interface{}) bool
```

Negative asserts that the specified element is negative

```
a.Negative(-1)
a.Negative(-1.23)
```

#### func (*Assertions) [Negativef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L994) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Negativef)added in v1.7.0

```
func (a *Assertions) Negativef(e interface{}, msg string, args ...interface{}) bool
```

Negativef asserts that the specified element is negative

```
a.Negativef(-1, "error message %s", "formatted")
a.Negativef(-1.23, "error message %s", "formatted")
```

#### func (*Assertions) [Never](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1005) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Never)added in v1.5.0

```
func (a *Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool
```

Never asserts that the given condition doesn't satisfy in waitFor time, periodically checking the target function each tick.

```
a.Never(func() bool { return false; }, time.Second, 10*time.Millisecond)
```

#### func (*Assertions) [Neverf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1016) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Neverf)added in v1.5.0

```
func (a *Assertions) Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool
```

Neverf asserts that the given condition doesn't satisfy in waitFor time, periodically checking the target function each tick.

```
a.Neverf(func() bool { return false; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
```

#### func (*Assertions) [Nil](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1026) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Nil)

```
func (a *Assertions) Nil(object interface{}, msgAndArgs ...interface{}) bool
```

Nil asserts that the specified object is nil.

```
a.Nil(err)
```

#### func (*Assertions) [Nilf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1036) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Nilf)added in v1.2.0

```
func (a *Assertions) Nilf(object interface{}, msg string, args ...interface{}) bool
```

Nilf asserts that the specified object is nil.

```
a.Nilf(err, "error message %s", "formatted")
```

#### func (*Assertions) [NoDirExists](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1045) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NoDirExists)added in v1.5.0

```
func (a *Assertions) NoDirExists(path string, msgAndArgs ...interface{}) bool
```

NoDirExists checks whether a directory does not exist in the given path. It fails if the path points to an existing _directory_ only.

#### func (*Assertions) [NoDirExistsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1054) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NoDirExistsf)added in v1.5.0

```
func (a *Assertions) NoDirExistsf(path string, msg string, args ...interface{}) bool
```

NoDirExistsf checks whether a directory does not exist in the given path. It fails if the path points to an existing _directory_ only.

#### func (*Assertions) [NoError](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1067) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NoError)

```
func (a *Assertions) NoError(err error, msgAndArgs ...interface{}) bool
```

NoError asserts that a function returned no error (i.e. `nil`).

```
  actualObj, err := SomeFunction()
  if a.NoError(err) {
	   assert.Equal(t, expectedObj, actualObj)
  }
```

#### func (*Assertions) [NoErrorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1080) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NoErrorf)added in v1.2.0

```
func (a *Assertions) NoErrorf(err error, msg string, args ...interface{}) bool
```

NoErrorf asserts that a function returned no error (i.e. `nil`).

```
  actualObj, err := SomeFunction()
  if a.NoErrorf(err, "error message %s", "formatted") {
	   assert.Equal(t, expectedObj, actualObj)
  }
```

#### func (*Assertions) [NoFileExists](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1089) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NoFileExists)added in v1.5.0

```
func (a *Assertions) NoFileExists(path string, msgAndArgs ...interface{}) bool
```

NoFileExists checks whether a file does not exist in a given path. It fails if the path points to an existing _file_ only.

#### func (*Assertions) [NoFileExistsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1098) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NoFileExistsf)added in v1.5.0

```
func (a *Assertions) NoFileExistsf(path string, msg string, args ...interface{}) bool
```

NoFileExistsf checks whether a file does not exist in a given path. It fails if the path points to an existing _file_ only.

#### func (*Assertions) [NotContains](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1111) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotContains)

```
func (a *Assertions) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) bool
```

NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the specified substring or element.

```
a.NotContains("Hello World", "Earth")
a.NotContains(["Hello", "World"], "Earth")
a.NotContains({"Hello": "World"}, "Earth")
```

#### func (*Assertions) [NotContainsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1124) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotContainsf)added in v1.2.0

```
func (a *Assertions) NotContainsf(s interface{}, contains interface{}, msg string, args ...interface{}) bool
```

NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the specified substring or element.

```
a.NotContainsf("Hello World", "Earth", "error message %s", "formatted")
a.NotContainsf(["Hello", "World"], "Earth", "error message %s", "formatted")
a.NotContainsf({"Hello": "World"}, "Earth", "error message %s", "formatted")
```

#### func (*Assertions) [NotElementsMatch](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1141) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotElementsMatch)added in v1.10.0

```
func (a *Assertions) NotElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{}) bool
```

NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should not match. This is an inverse of ElementsMatch.

a.NotElementsMatch([1, 1, 2, 3], [1, 1, 2, 3]) -> false

a.NotElementsMatch([1, 1, 2, 3], [1, 2, 3]) -> true

a.NotElementsMatch([1, 2, 3], [1, 2, 4]) -> true

#### func (*Assertions) [NotElementsMatchf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1158) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotElementsMatchf)added in v1.10.0

```
func (a *Assertions) NotElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{}) bool
```

NotElementsMatchf asserts that the specified listA(array, slice...) is NOT equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should not match. This is an inverse of ElementsMatch.

a.NotElementsMatchf([1, 1, 2, 3], [1, 1, 2, 3], "error message %s", "formatted") -> false

a.NotElementsMatchf([1, 1, 2, 3], [1, 2, 3], "error message %s", "formatted") -> true

a.NotElementsMatchf([1, 2, 3], [1, 2, 4], "error message %s", "formatted") -> true

#### func (*Assertions) [NotEmpty](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1171) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotEmpty)

```
func (a *Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{}) bool
```

NotEmpty asserts that the specified object is NOT empty. I.e. not nil, "", false, 0 or either a slice or a channel with len == 0.

```
if a.NotEmpty(obj) {
  assert.Equal(t, "two", obj[1])
}
```

#### func (*Assertions) [NotEmptyf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1184) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotEmptyf)added in v1.2.0

```
func (a *Assertions) NotEmptyf(object interface{}, msg string, args ...interface{}) bool
```

NotEmptyf asserts that the specified object is NOT empty. I.e. not nil, "", false, 0 or either a slice or a channel with len == 0.

```
if a.NotEmptyf(obj, "error message %s", "formatted") {
  assert.Equal(t, "two", obj[1])
}
```

#### func (*Assertions) [NotEqual](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1197) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotEqual)

```
func (a *Assertions) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```

NotEqual asserts that the specified values are NOT equal.

```
a.NotEqual(obj1, obj2)
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses).

#### func (*Assertions) [NotEqualValues](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1207) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotEqualValues)added in v1.6.0

```
func (a *Assertions) NotEqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```

NotEqualValues asserts that two objects are not equal even when converted to the same type

```
a.NotEqualValues(obj1, obj2)
```

#### func (*Assertions) [NotEqualValuesf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1217) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotEqualValuesf)added in v1.6.0

```
func (a *Assertions) NotEqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

NotEqualValuesf asserts that two objects are not equal even when converted to the same type

```
a.NotEqualValuesf(obj1, obj2, "error message %s", "formatted")
```

#### func (*Assertions) [NotEqualf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1230) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotEqualf)added in v1.2.0

```
func (a *Assertions) NotEqualf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

NotEqualf asserts that the specified values are NOT equal.

```
a.NotEqualf(obj1, obj2, "error message %s", "formatted")
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses).

#### func (*Assertions) [NotErrorAs](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1239) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotErrorAs)added in v1.10.0

```
func (a *Assertions) NotErrorAs(err error, target interface{}, msgAndArgs ...interface{}) bool
```

NotErrorAs asserts that none of the errors in err's chain matches target, but if so, sets target to that error value.

#### func (*Assertions) [NotErrorAsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1248) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotErrorAsf)added in v1.10.0

```
func (a *Assertions) NotErrorAsf(err error, target interface{}, msg string, args ...interface{}) bool
```

NotErrorAsf asserts that none of the errors in err's chain matches target, but if so, sets target to that error value.

#### func (*Assertions) [NotErrorIs](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1257) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotErrorIs)added in v1.7.0

```
func (a *Assertions) NotErrorIs(err error, target error, msgAndArgs ...interface{}) bool
```

NotErrorIs asserts that none of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func (*Assertions) [NotErrorIsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1266) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotErrorIsf)added in v1.7.0

```
func (a *Assertions) NotErrorIsf(err error, target error, msg string, args ...interface{}) bool
```

NotErrorIsf asserts that none of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func (*Assertions) [NotImplements](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1276) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotImplements)added in v1.9.0

```
func (a *Assertions) NotImplements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool
```

NotImplements asserts that an object does not implement the specified interface.

```
a.NotImplements((*MyInterface)(nil), new(MyObject))
```

#### func (*Assertions) [NotImplementsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1286) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotImplementsf)added in v1.9.0

```
func (a *Assertions) NotImplementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) bool
```

NotImplementsf asserts that an object does not implement the specified interface.

```
a.NotImplementsf((*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
```

#### func (*Assertions) [NotNil](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1296) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotNil)

```
func (a *Assertions) NotNil(object interface{}, msgAndArgs ...interface{}) bool
```

NotNil asserts that the specified object is not nil.

```
a.NotNil(err)
```

#### func (*Assertions) [NotNilf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1306) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotNilf)added in v1.2.0

```
func (a *Assertions) NotNilf(object interface{}, msg string, args ...interface{}) bool
```

NotNilf asserts that the specified object is not nil.

```
a.NotNilf(err, "error message %s", "formatted")
```

#### func (*Assertions) [NotPanics](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1316) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotPanics)

```
func (a *Assertions) NotPanics(f PanicTestFunc, msgAndArgs ...interface{}) bool
```

NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.

```
a.NotPanics(func(){ RemainCalm() })
```

#### func (*Assertions) [NotPanicsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1326) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotPanicsf)added in v1.2.0

```
func (a *Assertions) NotPanicsf(f PanicTestFunc, msg string, args ...interface{}) bool
```

NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.

```
a.NotPanicsf(func(){ RemainCalm() }, "error message %s", "formatted")
```

#### func (*Assertions) [NotRegexp](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1337) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotRegexp)

```
func (a *Assertions) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool
```

NotRegexp asserts that a specified regexp does not match a string.

```
a.NotRegexp(regexp.MustCompile("starts"), "it's starting")
a.NotRegexp("^start", "it's not starting")
```

#### func (*Assertions) [NotRegexpf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1348) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotRegexpf)added in v1.2.0

```
func (a *Assertions) NotRegexpf(rx interface{}, str interface{}, msg string, args ...interface{}) bool
```

NotRegexpf asserts that a specified regexp does not match a string.

```
a.NotRegexpf(regexp.MustCompile("starts"), "it's starting", "error message %s", "formatted")
a.NotRegexpf("^start", "it's not starting", "error message %s", "formatted")
```

#### func (*Assertions) [NotSame](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1361) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotSame)added in v1.5.0

```
func (a *Assertions) NotSame(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```

NotSame asserts that two pointers do not reference the same object.

```
a.NotSame(ptr1, ptr2)
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func (*Assertions) [NotSamef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1374) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotSamef)added in v1.5.0

```
func (a *Assertions) NotSamef(expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

NotSamef asserts that two pointers do not reference the same object.

```
a.NotSamef(ptr1, ptr2, "error message %s", "formatted")
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func (*Assertions) [NotSubset](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1387) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotSubset)added in v1.2.0

```
func (a *Assertions) NotSubset(list interface{}, subset interface{}, msgAndArgs ...interface{}) bool
```

NotSubset asserts that the specified list(array, slice...) or map does NOT contain all elements given in the specified subset list(array, slice...) or map.

```
a.NotSubset([1, 3, 4], [1, 2])
a.NotSubset({"x": 1, "y": 2}, {"z": 3})
```

#### func (*Assertions) [NotSubsetf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1400) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotSubsetf)added in v1.2.0

```
func (a *Assertions) NotSubsetf(list interface{}, subset interface{}, msg string, args ...interface{}) bool
```

NotSubsetf asserts that the specified list(array, slice...) or map does NOT contain all elements given in the specified subset list(array, slice...) or map.

```
a.NotSubsetf([1, 3, 4], [1, 2], "error message %s", "formatted")
a.NotSubsetf({"x": 1, "y": 2}, {"z": 3}, "error message %s", "formatted")
```

#### func (*Assertions) [NotZero](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1408) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotZero)

```
func (a *Assertions) NotZero(i interface{}, msgAndArgs ...interface{}) bool
```

NotZero asserts that i is not the zero value for its type.

#### func (*Assertions) [NotZerof](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1416) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.NotZerof)added in v1.2.0

```
func (a *Assertions) NotZerof(i interface{}, msg string, args ...interface{}) bool
```

NotZerof asserts that i is not the zero value for its type.

#### func (*Assertions) [Panics](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1426) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Panics)

```
func (a *Assertions) Panics(f PanicTestFunc, msgAndArgs ...interface{}) bool
```

Panics asserts that the code inside the specified PanicTestFunc panics.

```
a.Panics(func(){ GoCrazy() })
```

#### func (*Assertions) [PanicsWithError](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1438) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.PanicsWithError)added in v1.5.0

```
func (a *Assertions) PanicsWithError(errString string, f PanicTestFunc, msgAndArgs ...interface{}) bool
```

PanicsWithError asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value is an error that satisfies the EqualError comparison.

```
a.PanicsWithError("crazy error", func(){ GoCrazy() })
```

#### func (*Assertions) [PanicsWithErrorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1450) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.PanicsWithErrorf)added in v1.5.0

```
func (a *Assertions) PanicsWithErrorf(errString string, f PanicTestFunc, msg string, args ...interface{}) bool
```

PanicsWithErrorf asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value is an error that satisfies the EqualError comparison.

```
a.PanicsWithErrorf("crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func (*Assertions) [PanicsWithValue](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1461) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.PanicsWithValue)added in v1.2.0

```
func (a *Assertions) PanicsWithValue(expected interface{}, f PanicTestFunc, msgAndArgs ...interface{}) bool
```

PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value equals the expected panic value.

```
a.PanicsWithValue("crazy error", func(){ GoCrazy() })
```

#### func (*Assertions) [PanicsWithValuef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1472) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.PanicsWithValuef)added in v1.2.0

```
func (a *Assertions) PanicsWithValuef(expected interface{}, f PanicTestFunc, msg string, args ...interface{}) bool
```

PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value equals the expected panic value.

```
a.PanicsWithValuef("crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func (*Assertions) [Panicsf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1482) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Panicsf)added in v1.2.0

```
func (a *Assertions) Panicsf(f PanicTestFunc, msg string, args ...interface{}) bool
```

Panicsf asserts that the code inside the specified PanicTestFunc panics.

```
a.Panicsf(func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func (*Assertions) [Positive](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1493) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Positive)added in v1.7.0

```
func (a *Assertions) Positive(e interface{}, msgAndArgs ...interface{}) bool
```

Positive asserts that the specified element is positive

```
a.Positive(1)
a.Positive(1.23)
```

#### func (*Assertions) [Positivef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1504) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Positivef)added in v1.7.0

```
func (a *Assertions) Positivef(e interface{}, msg string, args ...interface{}) bool
```

Positivef asserts that the specified element is positive

```
a.Positivef(1, "error message %s", "formatted")
a.Positivef(1.23, "error message %s", "formatted")
```

#### func (*Assertions) [Regexp](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1515) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Regexp)

```
func (a *Assertions) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool
```

Regexp asserts that a specified regexp matches a string.

```
a.Regexp(regexp.MustCompile("start"), "it's starting")
a.Regexp("start...$", "it's not starting")
```

#### func (*Assertions) [Regexpf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1526) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Regexpf)added in v1.2.0

```
func (a *Assertions) Regexpf(rx interface{}, str interface{}, msg string, args ...interface{}) bool
```

Regexpf asserts that a specified regexp matches a string.

```
a.Regexpf(regexp.MustCompile("start"), "it's starting", "error message %s", "formatted")
a.Regexpf("start...$", "it's not starting", "error message %s", "formatted")
```

#### func (*Assertions) [Same](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1539) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Same)added in v1.4.0

```
func (a *Assertions) Same(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool
```

Same asserts that two pointers reference the same object.

```
a.Same(ptr1, ptr2)
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func (*Assertions) [Samef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1552) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Samef)added in v1.4.0

```
func (a *Assertions) Samef(expected interface{}, actual interface{}, msg string, args ...interface{}) bool
```

Samef asserts that two pointers reference the same object.

```
a.Samef(ptr1, ptr2, "error message %s", "formatted")
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func (*Assertions) [Subset](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1564) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Subset)added in v1.2.0

```
func (a *Assertions) Subset(list interface{}, subset interface{}, msgAndArgs ...interface{}) bool
```

Subset asserts that the specified list(array, slice...) or map contains all elements given in the specified subset list(array, slice...) or map.

```
a.Subset([1, 2, 3], [1, 2])
a.Subset({"x": 1, "y": 2}, {"x": 1})
```

#### func (*Assertions) [Subsetf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1576) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Subsetf)added in v1.2.0

```
func (a *Assertions) Subsetf(list interface{}, subset interface{}, msg string, args ...interface{}) bool
```

Subsetf asserts that the specified list(array, slice...) or map contains all elements given in the specified subset list(array, slice...) or map.

```
a.Subsetf([1, 2, 3], [1, 2], "error message %s", "formatted")
a.Subsetf({"x": 1, "y": 2}, {"x": 1}, "error message %s", "formatted")
```

#### func (*Assertions) [True](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1586) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.True)

```
func (a *Assertions) True(value bool, msgAndArgs ...interface{}) bool
```

True asserts that the specified value is true.

```
a.True(myBool)
```

#### func (*Assertions) [Truef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1596) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Truef)added in v1.2.0

```
func (a *Assertions) Truef(value bool, msg string, args ...interface{}) bool
```

Truef asserts that the specified value is true.

```
a.Truef(myBool, "error message %s", "formatted")
```

#### func (*Assertions) [WithinDuration](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1606) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.WithinDuration)

```
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool
```

WithinDuration asserts that the two times are within duration delta of each other.

```
a.WithinDuration(time.Now(), time.Now(), 10*time.Second)
```

#### func (*Assertions) [WithinDurationf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1616) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.WithinDurationf)added in v1.2.0

```
func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) bool
```

WithinDurationf asserts that the two times are within duration delta of each other.

```
a.WithinDurationf(time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
```

#### func (*Assertions) [WithinRange](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1626) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.WithinRange)added in v1.8.0

```
func (a *Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...interface{}) bool
```

WithinRange asserts that a time is within a time range (inclusive).

```
a.WithinRange(time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
```

#### func (*Assertions) [WithinRangef](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1636) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.WithinRangef)added in v1.8.0

```
func (a *Assertions) WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ...interface{}) bool
```

WithinRangef asserts that a time is within a time range (inclusive).

```
a.WithinRangef(time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second), "error message %s", "formatted")
```

#### func (*Assertions) [YAMLEq](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1644) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.YAMLEq)added in v1.4.0

```
func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...interface{}) bool
```

YAMLEq asserts that two YAML strings are equivalent.

#### func (*Assertions) [YAMLEqf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1652) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.YAMLEqf)added in v1.4.0

```
func (a *Assertions) YAMLEqf(expected string, actual string, msg string, args ...interface{}) bool
```

YAMLEqf asserts that two YAML strings are equivalent.

#### func (*Assertions) [Zero](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1660) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Zero)

```
func (a *Assertions) Zero(i interface{}, msgAndArgs ...interface{}) bool
```

Zero asserts that i is the zero value for its type.

#### func (*Assertions) [Zerof](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_forward.go#L1668) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Assertions.Zerof)added in v1.2.0

```
func (a *Assertions) Zerof(i interface{}, msg string, args ...interface{}) bool
```

Zerof asserts that i is the zero value for its type.

#### type [BoolAssertionFunc](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L44) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#BoolAssertionFunc)added in v1.2.2

```
type BoolAssertionFunc func(TestingT, bool, ...interface{}) bool
```

BoolAssertionFunc is a common function prototype when validating a bool value. Can be useful for table driven tests.

<details tabindex="-1" id="example-BoolAssertionFunc" class="Documentation-exampleDetails js-exampleContainer" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 1rem 0px 0px; padding: 0px; vertical-align: baseline; display: block;"><summary class="Documentation-exampleDetailsHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px 0px 2rem; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); cursor: pointer; outline: none; text-decoration: none;">Example<span>&nbsp;</span><a href="https://pkg.go.dev/github.com/stretchr/testify/assert#example-BoolAssertionFunc" title="Go to Example" aria-label="Go to Example" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); text-decoration: none; opacity: 0;">¶</a></summary><div class="Documentation-exampleDetailsBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline;"><textarea class="Documentation-exampleCode code" spellcheck="false" style="box-sizing: border-box; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; font-size: 0.875rem; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; background-color: var(--color-background-accented); border: var(--border); border-top-left-radius: ; border-top-right-radius: ; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px; color: var(--color-text); overflow-x: auto; padding: 0.625rem; tab-size: 4; white-space: pre; height: 29.625rem; outline: none; resize: none; width: 1263.31px; margin: 0px;"></textarea><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: -0.25rem 0px 1rem; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: 0px 0px 0.3rem 0.3rem; color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre;"><span class="Documentation-exampleOutputLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span><span class="Documentation-exampleOutput" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px 0px 0.5rem; padding: 0px; vertical-align: baseline; border-top-left-radius: 0px; border-top-right-radius: 0px;"></span></pre></div></details>

#### type [CollectT](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1960) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#CollectT)added in v1.8.3

```
type CollectT struct {
	// contains filtered or unexported fields
}
```

CollectT implements the TestingT interface and collects all errors.

<details class="Documentation-deprecatedDetails js-deprecatedDetails" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; display: block; color: var(--color-text-subtle);"><summary style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; list-style: none; opacity: 1;"><h4 tabindex="-1" id="CollectT.Copy" data-kind="method" class="Documentation-typeMethodHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 600; font-stretch: inherit; line-height: 1.25em; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1.125rem; margin: 1.5rem 0px 0.5rem; padding: 0px; vertical-align: baseline; word-break: break-word; align-items: baseline; display: flex; justify-content: space-between;"><span class="Documentation-deprecatedTitle" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; align-items: center; display: flex; gap: 0.5rem;">func (*CollectT)<a class="Documentation-source" href="https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1984" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 1;">Copy</a><span class="Documentation-deprecatedTag" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: 1.375; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.75rem; margin: 0px; padding: 0.125rem 0.25rem; vertical-align: middle; background-color: var(--color-border); border-radius: 0.125rem; color: var(--color-text-inverted); text-transform: uppercase;">deprecated</span><span class="Documentation-deprecatedBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.87rem; margin: 0px 0.5rem 0px 0.25rem; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></span><span class="Documentation-sinceVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.9375rem; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"><span class="Documentation-sinceVersionLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 15px; margin: 0px; padding: 0px; vertical-align: baseline;">added in</span><span>&nbsp;</span><span class="Documentation-sinceVersionVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 15px; margin: 0px; padding: 0px; vertical-align: baseline;">v1.8.3</span></span></h4></summary><div class="go-Message go-Message--warning Documentation-deprecatedItemBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 1rem 1rem 0.5rem; vertical-align: baseline; color: var(--gray-1); width: 1263.31px; background-color: var(--color-background-warning);"><div class="Documentation-declaration" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre-wrap; scroll-padding-top: calc(var(--js-sticky-header-height, 3.5rem) + .75rem); word-break: break-all; overflow-wrap: break-word;"><a href="https://pkg.go.dev/github.com/stretchr/testify/assert#CollectT" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/github.com/stretchr/testify/assert#TestingT" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></pre></div><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p></div></details>

#### func (*CollectT) [Errorf](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1968) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#CollectT.Errorf)added in v1.8.3

```
func (c *CollectT) Errorf(format string, args ...interface{})
```

Errorf collects the error.

#### func (*CollectT) [FailNow](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1973) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#CollectT.FailNow)added in v1.8.3

```
func (c *CollectT) FailNow()
```

FailNow stops execution by calling runtime.Goexit.

<details class="Documentation-deprecatedDetails js-deprecatedDetails" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; display: block; color: var(--color-text-subtle);"><summary style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; list-style: none; opacity: 1;"><h4 tabindex="-1" id="CollectT.Reset" data-kind="method" class="Documentation-typeMethodHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 600; font-stretch: inherit; line-height: 1.25em; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1.125rem; margin: 1.5rem 0px 0.5rem; padding: 0px; vertical-align: baseline; word-break: break-word; align-items: baseline; display: flex; justify-content: space-between;"><span class="Documentation-deprecatedTitle" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; align-items: center; display: flex; gap: 0.5rem;">func (*CollectT)<a class="Documentation-source" href="https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1979" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 1;">Reset</a><span class="Documentation-deprecatedTag" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: 1.375; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.75rem; margin: 0px; padding: 0.125rem 0.25rem; vertical-align: middle; background-color: var(--color-border); border-radius: 0.125rem; color: var(--color-text-inverted); text-transform: uppercase;">deprecated</span><span class="Documentation-deprecatedBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.87rem; margin: 0px 0.5rem 0px 0.25rem; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></span><span class="Documentation-sinceVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.9375rem; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"><span class="Documentation-sinceVersionLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 15px; margin: 0px; padding: 0px; vertical-align: baseline;">added in</span><span>&nbsp;</span><span class="Documentation-sinceVersionVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 15px; margin: 0px; padding: 0px; vertical-align: baseline;">v1.8.3</span></span></h4></summary><div class="go-Message go-Message--warning Documentation-deprecatedItemBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 1rem 1rem 0.5rem; vertical-align: baseline; color: var(--gray-1); width: 1263.31px; background-color: var(--color-background-warning);"><div class="Documentation-declaration" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre-wrap; scroll-padding-top: calc(var(--js-sticky-header-height, 3.5rem) + .75rem); word-break: break-all; overflow-wrap: break-word;"><a href="https://pkg.go.dev/github.com/stretchr/testify/assert#CollectT" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></pre></div><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p></div></details>

<details class="Documentation-deprecatedDetails js-deprecatedDetails" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; display: block; color: var(--color-text-subtle);"><summary style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; list-style: none; opacity: 1;"><h4 tabindex="-1" id="CompareType" data-kind="type" class="Documentation-typeHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 600; font-stretch: inherit; line-height: 1.25em; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1.125rem; margin: 1.5rem 0px 0.5rem; padding: 0px; vertical-align: baseline; word-break: break-word; align-items: baseline; display: flex; justify-content: space-between;"><span class="Documentation-deprecatedTitle" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; align-items: center; display: flex; gap: 0.5rem;">type<a class="Documentation-source" href="https://github.com/stretchr/testify/blob/v1.10.0/assert/assertion_compare.go#L11" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 1;">CompareType</a><span class="Documentation-deprecatedTag" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: 1.375; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.75rem; margin: 0px; padding: 0.125rem 0.25rem; vertical-align: middle; background-color: var(--color-border); border-radius: 0.125rem; color: var(--color-text-inverted); text-transform: uppercase;">deprecated</span><span class="Documentation-deprecatedBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.87rem; margin: 0px 0.5rem 0px 0.25rem; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></span><span class="Documentation-sinceVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.9375rem; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"><span class="Documentation-sinceVersionLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 15px; margin: 0px; padding: 0px; vertical-align: baseline;">added in</span><span>&nbsp;</span><span class="Documentation-sinceVersionVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 15px; margin: 0px; padding: 0px; vertical-align: baseline;">v1.6.0</span></span></h4></summary><div class="go-Message go-Message--warning Documentation-deprecatedItemBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 1rem 1rem 0.5rem; vertical-align: baseline; color: var(--gray-1); width: 1263.31px; background-color: var(--color-background-warning);"><div class="Documentation-declaration" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre; scroll-padding-top: calc(var(--js-sticky-header-height, 3.5rem) + .75rem);"></pre></div><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p></div></details>

#### type [Comparison](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L55) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#Comparison)

```
type Comparison func() (success bool)
```

Comparison is a custom function that returns true on success and false on failure

#### type [ComparisonAssertionFunc](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L36) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ComparisonAssertionFunc)added in v1.2.2

```
type ComparisonAssertionFunc func(TestingT, interface{}, interface{}, ...interface{}) bool
```

ComparisonAssertionFunc is a common function prototype when comparing two values. Can be useful for table driven tests.

<details tabindex="-1" id="example-ComparisonAssertionFunc" class="Documentation-exampleDetails js-exampleContainer" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 1rem 0px 0px; padding: 0px; vertical-align: baseline; display: block;"><summary class="Documentation-exampleDetailsHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px 0px 2rem; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); cursor: pointer; outline: none; text-decoration: none;">Example<span>&nbsp;</span><a href="https://pkg.go.dev/github.com/stretchr/testify/assert#example-ComparisonAssertionFunc" title="Go to Example" aria-label="Go to Example" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); text-decoration: none; opacity: 0;">¶</a></summary><div class="Documentation-exampleDetailsBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline;"><textarea class="Documentation-exampleCode code" spellcheck="false" style="box-sizing: border-box; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; font-size: 0.875rem; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; background-color: var(--color-background-accented); border: var(--border); border-top-left-radius: ; border-top-right-radius: ; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px; color: var(--color-text); overflow-x: auto; padding: 0.625rem; tab-size: 4; white-space: pre; height: 35.875rem; outline: none; resize: none; width: 1263.31px; margin: 0px;"></textarea><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: -0.25rem 0px 1rem; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: 0px 0px 0.3rem 0.3rem; color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre;"><span class="Documentation-exampleOutputLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span><span class="Documentation-exampleOutput" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px 0px 0.5rem; padding: 0px; vertical-align: baseline; border-top-left-radius: 0px; border-top-right-radius: 0px;"></span></pre></div></details>

#### type [ErrorAssertionFunc](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L48) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ErrorAssertionFunc)added in v1.2.2

```
type ErrorAssertionFunc func(TestingT, error, ...interface{}) bool
```

ErrorAssertionFunc is a common function prototype when validating an error value. Can be useful for table driven tests.

<details tabindex="-1" id="example-ErrorAssertionFunc" class="Documentation-exampleDetails js-exampleContainer" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 1rem 0px 0px; padding: 0px; vertical-align: baseline; display: block;"><summary class="Documentation-exampleDetailsHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px 0px 2rem; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); cursor: pointer; outline: none; text-decoration: none;">Example<span>&nbsp;</span><a href="https://pkg.go.dev/github.com/stretchr/testify/assert#example-ErrorAssertionFunc" title="Go to Example" aria-label="Go to Example" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); text-decoration: none; opacity: 0;">¶</a></summary><div class="Documentation-exampleDetailsBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline;"><textarea class="Documentation-exampleCode code" spellcheck="false" style="box-sizing: border-box; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; font-size: 0.875rem; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; background-color: var(--color-background-accented); border: var(--border); border-top-left-radius: ; border-top-right-radius: ; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px; color: var(--color-text); overflow-x: auto; padding: 0.625rem; tab-size: 4; white-space: pre; height: 30.875rem; outline: none; resize: none; width: 1263.31px; margin: 0px;"></textarea><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: -0.25rem 0px 1rem; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: 0px 0px 0.3rem 0.3rem; color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre;"><span class="Documentation-exampleOutputLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span><span class="Documentation-exampleOutput" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px 0px 0.5rem; padding: 0px; vertical-align: baseline; border-top-left-radius: 0px; border-top-right-radius: 0px;"></span></pre></div></details>

#### type [PanicAssertionFunc](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L52) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#PanicAssertionFunc)added in v1.10.0

```
type PanicAssertionFunc = func(t TestingT, f PanicTestFunc, msgAndArgs ...interface{}) bool
```

PanicAssertionFunc is a common function prototype when validating a panic value. Can be useful for table driven tests.

<details tabindex="-1" id="example-PanicAssertionFunc" class="Documentation-exampleDetails js-exampleContainer" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 1rem 0px 0px; padding: 0px; vertical-align: baseline; display: block;"><summary class="Documentation-exampleDetailsHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px 0px 2rem; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); cursor: pointer; outline: none; text-decoration: none;">Example<span>&nbsp;</span><a href="https://pkg.go.dev/github.com/stretchr/testify/assert#example-PanicAssertionFunc" title="Go to Example" aria-label="Go to Example" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); text-decoration: none; opacity: 0;">¶</a></summary><div class="Documentation-exampleDetailsBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline;"><textarea class="Documentation-exampleCode code" spellcheck="false" style="box-sizing: border-box; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; font-size: 0.875rem; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; background-color: var(--color-background-accented); border: var(--border); border-top-left-radius: ; border-top-right-radius: ; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px; color: var(--color-text); overflow-x: auto; padding: 0.625rem; tab-size: 4; white-space: pre; height: 22.125rem; outline: none; resize: none; width: 1263.31px; margin: 0px;"></textarea><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: -0.25rem 0px 1rem; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: 0px 0px 0.3rem 0.3rem; color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre;"><span class="Documentation-exampleOutputLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span><span class="Documentation-exampleOutput" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px 0px 0.5rem; padding: 0px; vertical-align: baseline; border-top-left-radius: 0px; border-top-right-radius: 0px;"></span></pre></div></details>

#### type [PanicTestFunc](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L1225) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#PanicTestFunc)

```
type PanicTestFunc func()
```

PanicTestFunc defines a func that should be passed to the assert.Panics and assert.NotPanics methods, and represents a simple func that takes no arguments, and returns nothing.

#### type [TestingT](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L30) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#TestingT)

```
type TestingT interface {
	Errorf(format string, args ...interface{})
}
```

TestingT is an interface wrapper around *testing.T

#### type [ValueAssertionFunc](https://github.com/stretchr/testify/blob/v1.10.0/assert/assertions.go#L40) [¶](https://pkg.go.dev/github.com/stretchr/testify/assert#ValueAssertionFunc)added in v1.2.2

```
type ValueAssertionFunc func(TestingT, interface{}, ...interface{}) bool
```

ValueAssertionFunc is a common function prototype when validating a single value. Can be useful for table driven tests.



