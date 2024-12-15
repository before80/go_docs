+++
title = "require"
date = 2024-12-15T11:07:47+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/stretchr/testify/require](https://pkg.go.dev/github.com/stretchr/testify/require)
>
> 收录该文档时间： `2024-12-15T11:07:47+08:00`
>
> 版本：[Version: v1.10.0](https://pkg.go.dev/github.com/stretchr/testify/require?tab=versions)

### Overview [¶](https://pkg.go.dev/github.com/stretchr/testify/require#pkg-overview)

- [Example Usage](https://pkg.go.dev/github.com/stretchr/testify/require#hdr-Example_Usage)
- [Assertions](https://pkg.go.dev/github.com/stretchr/testify/require#hdr-Assertions)

Package require implements the same assertions as the `assert` package but stops test execution when a test fails.

#### Example Usage [¶](https://pkg.go.dev/github.com/stretchr/testify/require#hdr-Example_Usage)

The following is a complete example using require in a standard test function:

```
import (
  "testing"
  "github.com/stretchr/testify/require"
)

func TestSomething(t *testing.T) {

  var a string = "Hello"
  var b string = "Hello"

  require.Equal(t, a, b, "The two words should be the same.")

}
```

#### Assertions [¶](https://pkg.go.dev/github.com/stretchr/testify/require#hdr-Assertions)

The `require` package have same global functions as in the `assert` package, but instead of returning a boolean result they call `t.FailNow()`.

Every assertion function also takes an optional string message as the final argument, allowing custom error messages to be appended to the message the assertion method outputs.

### Constants [¶](https://pkg.go.dev/github.com/stretchr/testify/require#pkg-constants)

This section is empty.

### Variables [¶](https://pkg.go.dev/github.com/stretchr/testify/require#pkg-variables)

This section is empty.

### Functions [¶](https://pkg.go.dev/github.com/stretchr/testify/require#pkg-functions)

#### func [Condition](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L13) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Condition)

```
func Condition(t TestingT, comp assert.Comparison, msgAndArgs ...interface{})
```

Condition uses a Comparison to assert a complex condition.

#### func [Conditionf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L24) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Conditionf)added in v1.2.0

```
func Conditionf(t TestingT, comp assert.Comparison, msg string, args ...interface{})
```

Conditionf uses a Comparison to assert a complex condition.

#### func [Contains](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L40) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Contains)

```
func Contains(t TestingT, s interface{}, contains interface{}, msgAndArgs ...interface{})
```

Contains asserts that the specified string, list(array, slice...) or map contains the specified substring or element.

```
require.Contains(t, "Hello World", "World")
require.Contains(t, ["Hello", "World"], "World")
require.Contains(t, {"Hello": "World"}, "Hello")
```

#### func [Containsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L56) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Containsf)added in v1.2.0

```
func Containsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{})
```

Containsf asserts that the specified string, list(array, slice...) or map contains the specified substring or element.

```
require.Containsf(t, "Hello World", "World", "error message %s", "formatted")
require.Containsf(t, ["Hello", "World"], "World", "error message %s", "formatted")
require.Containsf(t, {"Hello": "World"}, "Hello", "error message %s", "formatted")
```

#### func [DirExists](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L68) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#DirExists)added in v1.2.0

```
func DirExists(t TestingT, path string, msgAndArgs ...interface{})
```

DirExists checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.

#### func [DirExistsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L80) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#DirExistsf)added in v1.2.0

```
func DirExistsf(t TestingT, path string, msg string, args ...interface{})
```

DirExistsf checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.

#### func [ElementsMatch](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L95) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ElementsMatch)added in v1.2.0

```
func ElementsMatch(t TestingT, listA interface{}, listB interface{}, msgAndArgs ...interface{})
```

ElementsMatch asserts that the specified listA(array, slice...) is equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should match.

require.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])

#### func [ElementsMatchf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L110) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ElementsMatchf)added in v1.2.0

```
func ElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{})
```

ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should match.

require.ElementsMatchf(t, [1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")

#### func [Empty](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L124) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Empty)

```
func Empty(t TestingT, object interface{}, msgAndArgs ...interface{})
```

Empty asserts that the specified object is empty. I.e. nil, "", false, 0 or either a slice or a channel with len == 0.

```
require.Empty(t, obj)
```

#### func [Emptyf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L138) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Emptyf)added in v1.2.0

```
func Emptyf(t TestingT, object interface{}, msg string, args ...interface{})
```

Emptyf asserts that the specified object is empty. I.e. nil, "", false, 0 or either a slice or a channel with len == 0.

```
require.Emptyf(t, obj, "error message %s", "formatted")
```

#### func [Equal](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L155) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Equal)

```
func Equal(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

Equal asserts that two objects are equal.

```
require.Equal(t, 123, 123)
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses). Function equality cannot be determined and will always fail.

#### func [EqualError](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L170) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#EqualError)

```
func EqualError(t TestingT, theError error, errString string, msgAndArgs ...interface{})
```

EqualError asserts that a function returned an error (i.e. not `nil`) and that it is equal to the provided error.

```
actualObj, err := SomeFunction()
require.EqualError(t, err,  expectedErrorString)
```

#### func [EqualErrorf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L185) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#EqualErrorf)added in v1.2.0

```
func EqualErrorf(t TestingT, theError error, errString string, msg string, args ...interface{})
```

EqualErrorf asserts that a function returned an error (i.e. not `nil`) and that it is equal to the provided error.

```
actualObj, err := SomeFunction()
require.EqualErrorf(t, err,  expectedErrorString, "error message %s", "formatted")
```

#### func [EqualExportedValues](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L205) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#EqualExportedValues)added in v1.8.3

```
func EqualExportedValues(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

EqualExportedValues asserts that the types of two objects are equal and their public fields are also equal. This is useful for comparing structs that have private fields that could potentially differ.

```
 type S struct {
	Exported     	int
	notExported   	int
 }
 require.EqualExportedValues(t, S{1, 2}, S{1, 3}) => true
 require.EqualExportedValues(t, S{1, 2}, S{2, 3}) => false
```

#### func [EqualExportedValuesf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L225) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#EqualExportedValuesf)added in v1.8.3

```
func EqualExportedValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{})
```

EqualExportedValuesf asserts that the types of two objects are equal and their public fields are also equal. This is useful for comparing structs that have private fields that could potentially differ.

```
 type S struct {
	Exported     	int
	notExported   	int
 }
 require.EqualExportedValuesf(t, S{1, 2}, S{1, 3}, "error message %s", "formatted") => true
 require.EqualExportedValuesf(t, S{1, 2}, S{2, 3}, "error message %s", "formatted") => false
```

#### func [EqualValues](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L239) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#EqualValues)

```
func EqualValues(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

EqualValues asserts that two objects are equal or convertible to the larger type and equal.

```
require.EqualValues(t, uint32(123), int32(123))
```

#### func [EqualValuesf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L253) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#EqualValuesf)added in v1.2.0

```
func EqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{})
```

EqualValuesf asserts that two objects are equal or convertible to the larger type and equal.

```
require.EqualValuesf(t, uint32(123), int32(123), "error message %s", "formatted")
```

#### func [Equalf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L270) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Equalf)added in v1.2.0

```
func Equalf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{})
```

Equalf asserts that two objects are equal.

```
require.Equalf(t, 123, 123, "error message %s", "formatted")
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses). Function equality cannot be determined and will always fail.

#### func [Error](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L286) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Error)

```
func Error(t TestingT, err error, msgAndArgs ...interface{})
```

Error asserts that a function returned an error (i.e. not `nil`).

```
  actualObj, err := SomeFunction()
  if require.Error(t, err) {
	   require.Equal(t, expectedError, err)
  }
```

#### func [ErrorAs](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L298) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ErrorAs)added in v1.7.0

```
func ErrorAs(t TestingT, err error, target interface{}, msgAndArgs ...interface{})
```

ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value. This is a wrapper for errors.As.

#### func [ErrorAsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L310) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ErrorAsf)added in v1.7.0

```
func ErrorAsf(t TestingT, err error, target interface{}, msg string, args ...interface{})
```

ErrorAsf asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value. This is a wrapper for errors.As.

#### func [ErrorContains](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L325) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ErrorContains)added in v1.7.1

```
func ErrorContains(t TestingT, theError error, contains string, msgAndArgs ...interface{})
```

ErrorContains asserts that a function returned an error (i.e. not `nil`) and that the error contains the specified substring.

```
actualObj, err := SomeFunction()
require.ErrorContains(t, err,  expectedErrorSubString)
```

#### func [ErrorContainsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L340) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ErrorContainsf)added in v1.7.1

```
func ErrorContainsf(t TestingT, theError error, contains string, msg string, args ...interface{})
```

ErrorContainsf asserts that a function returned an error (i.e. not `nil`) and that the error contains the specified substring.

```
actualObj, err := SomeFunction()
require.ErrorContainsf(t, err,  expectedErrorSubString, "error message %s", "formatted")
```

#### func [ErrorIs](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L352) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ErrorIs)added in v1.7.0

```
func ErrorIs(t TestingT, err error, target error, msgAndArgs ...interface{})
```

ErrorIs asserts that at least one of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func [ErrorIsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L364) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ErrorIsf)added in v1.7.0

```
func ErrorIsf(t TestingT, err error, target error, msg string, args ...interface{})
```

ErrorIsf asserts that at least one of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func [Errorf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L380) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Errorf)added in v1.2.0

```
func Errorf(t TestingT, err error, msg string, args ...interface{})
```

Errorf asserts that a function returned an error (i.e. not `nil`).

```
  actualObj, err := SomeFunction()
  if require.Errorf(t, err, "error message %s", "formatted") {
	   require.Equal(t, expectedErrorf, err)
  }
```

#### func [Eventually](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L394) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Eventually)added in v1.4.0

```
func Eventually(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{})
```

Eventually asserts that given condition will be met in waitFor time, periodically checking target function each tick.

```
require.Eventually(t, func() bool { return true; }, time.Second, 10*time.Millisecond)
```

#### func [EventuallyWithT](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L422) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#EventuallyWithT)added in v1.8.3

```
func EventuallyWithT(t TestingT, condition func(collect *assert.CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{})
```

EventuallyWithT asserts that given condition will be met in waitFor time, periodically checking target function each tick. In contrast to Eventually, it supplies a CollectT to the condition function, so that the condition function can use the CollectT to call other assertions. The condition is considered "met" if no errors are raised in a tick. The supplied CollectT collects all errors from one tick (if there are any). If the condition is not met before waitFor, the collected errors of the last tick are copied to t.

```
externalValue := false
go func() {
	time.Sleep(8*time.Second)
	externalValue = true
}()
require.EventuallyWithT(t, func(c *require.CollectT) {
	// add assertions as needed; any assertion failure will fail the current tick
	require.True(c, externalValue, "expected 'externalValue' to be true")
}, 10*time.Second, 1*time.Second, "external state has not changed to 'true'; still false")
```

#### func [EventuallyWithTf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L450) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#EventuallyWithTf)added in v1.8.3

```
func EventuallyWithTf(t TestingT, condition func(collect *assert.CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...interface{})
```

EventuallyWithTf asserts that given condition will be met in waitFor time, periodically checking target function each tick. In contrast to Eventually, it supplies a CollectT to the condition function, so that the condition function can use the CollectT to call other assertions. The condition is considered "met" if no errors are raised in a tick. The supplied CollectT collects all errors from one tick (if there are any). If the condition is not met before waitFor, the collected errors of the last tick are copied to t.

```
externalValue := false
go func() {
	time.Sleep(8*time.Second)
	externalValue = true
}()
require.EventuallyWithTf(t, func(c *require.CollectT, "error message %s", "formatted") {
	// add assertions as needed; any assertion failure will fail the current tick
	require.True(c, externalValue, "expected 'externalValue' to be true")
}, 10*time.Second, 1*time.Second, "external state has not changed to 'true'; still false")
```

#### func [Eventuallyf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L464) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Eventuallyf)added in v1.4.0

```
func Eventuallyf(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{})
```

Eventuallyf asserts that given condition will be met in waitFor time, periodically checking target function each tick.

```
require.Eventuallyf(t, func() bool { return true; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
```

#### func [Exactly](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L477) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Exactly)

```
func Exactly(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

Exactly asserts that two objects are equal in value and type.

```
require.Exactly(t, int32(123), int64(123))
```

#### func [Exactlyf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L490) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Exactlyf)added in v1.2.0

```
func Exactlyf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{})
```

Exactlyf asserts that two objects are equal in value and type.

```
require.Exactlyf(t, int32(123), int64(123), "error message %s", "formatted")
```

#### func [Fail](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L501) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Fail)

```
func Fail(t TestingT, failureMessage string, msgAndArgs ...interface{})
```

Fail reports a failure through

#### func [FailNow](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L512) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#FailNow)

```
func FailNow(t TestingT, failureMessage string, msgAndArgs ...interface{})
```

FailNow fails test

#### func [FailNowf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L523) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#FailNowf)added in v1.2.0

```
func FailNowf(t TestingT, failureMessage string, msg string, args ...interface{})
```

FailNowf fails test

#### func [Failf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L534) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Failf)added in v1.2.0

```
func Failf(t TestingT, failureMessage string, msg string, args ...interface{})
```

Failf reports a failure through

#### func [False](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L547) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#False)

```
func False(t TestingT, value bool, msgAndArgs ...interface{})
```

False asserts that the specified value is false.

```
require.False(t, myBool)
```

#### func [Falsef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L560) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Falsef)added in v1.2.0

```
func Falsef(t TestingT, value bool, msg string, args ...interface{})
```

Falsef asserts that the specified value is false.

```
require.Falsef(t, myBool, "error message %s", "formatted")
```

#### func [FileExists](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L572) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#FileExists)added in v1.2.0

```
func FileExists(t TestingT, path string, msgAndArgs ...interface{})
```

FileExists checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.

#### func [FileExistsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L584) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#FileExistsf)added in v1.2.0

```
func FileExistsf(t TestingT, path string, msg string, args ...interface{})
```

FileExistsf checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.

#### func [Greater](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L599) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Greater)added in v1.4.0

```
func Greater(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
```

Greater asserts that the first element is greater than the second

```
require.Greater(t, 2, 1)
require.Greater(t, float64(2), float64(1))
require.Greater(t, "b", "a")
```

#### func [GreaterOrEqual](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L615) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#GreaterOrEqual)added in v1.4.0

```
func GreaterOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
```

GreaterOrEqual asserts that the first element is greater than or equal to the second

```
require.GreaterOrEqual(t, 2, 1)
require.GreaterOrEqual(t, 2, 2)
require.GreaterOrEqual(t, "b", "a")
require.GreaterOrEqual(t, "b", "b")
```

#### func [GreaterOrEqualf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L631) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#GreaterOrEqualf)added in v1.4.0

```
func GreaterOrEqualf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{})
```

GreaterOrEqualf asserts that the first element is greater than or equal to the second

```
require.GreaterOrEqualf(t, 2, 1, "error message %s", "formatted")
require.GreaterOrEqualf(t, 2, 2, "error message %s", "formatted")
require.GreaterOrEqualf(t, "b", "a", "error message %s", "formatted")
require.GreaterOrEqualf(t, "b", "b", "error message %s", "formatted")
```

#### func [Greaterf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L646) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Greaterf)added in v1.4.0

```
func Greaterf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{})
```

Greaterf asserts that the first element is greater than the second

```
require.Greaterf(t, 2, 1, "error message %s", "formatted")
require.Greaterf(t, float64(2), float64(1), "error message %s", "formatted")
require.Greaterf(t, "b", "a", "error message %s", "formatted")
```

#### func [HTTPBodyContains](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L662) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPBodyContains)

```
func HTTPBodyContains(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{})
```

HTTPBodyContains asserts that a specified handler returns a body that contains a string.

```
require.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPBodyContainsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L678) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPBodyContainsf)added in v1.2.0

```
func HTTPBodyContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{})
```

HTTPBodyContainsf asserts that a specified handler returns a body that contains a string.

```
require.HTTPBodyContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPBodyNotContains](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L694) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPBodyNotContains)

```
func HTTPBodyNotContains(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{})
```

HTTPBodyNotContains asserts that a specified handler returns a body that does not contain a string.

```
require.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPBodyNotContainsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L710) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPBodyNotContainsf)added in v1.2.0

```
func HTTPBodyNotContainsf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{})
```

HTTPBodyNotContainsf asserts that a specified handler returns a body that does not contain a string.

```
require.HTTPBodyNotContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPError](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L725) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPError)

```
func HTTPError(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{})
```

HTTPError asserts that a specified handler returns an error status code.

```
require.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPErrorf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L740) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPErrorf)added in v1.2.0

```
func HTTPErrorf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{})
```

HTTPErrorf asserts that a specified handler returns an error status code.

```
require.HTTPErrorf(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPRedirect](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L755) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPRedirect)

```
func HTTPRedirect(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{})
```

HTTPRedirect asserts that a specified handler returns a redirect status code.

```
require.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPRedirectf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L770) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPRedirectf)added in v1.2.0

```
func HTTPRedirectf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{})
```

HTTPRedirectf asserts that a specified handler returns a redirect status code.

```
require.HTTPRedirectf(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPStatusCode](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L785) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPStatusCode)added in v1.6.0

```
func HTTPStatusCode(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...interface{})
```

HTTPStatusCode asserts that a specified handler returns a specified status code.

```
require.HTTPStatusCode(t, myHandler, "GET", "/notImplemented", nil, 501)
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPStatusCodef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L800) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPStatusCodef)added in v1.6.0

```
func HTTPStatusCodef(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{})
```

HTTPStatusCodef asserts that a specified handler returns a specified status code.

```
require.HTTPStatusCodef(t, myHandler, "GET", "/notImplemented", nil, 501, "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPSuccess](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L815) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPSuccess)

```
func HTTPSuccess(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{})
```

HTTPSuccess asserts that a specified handler returns a success status code.

```
require.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
```

Returns whether the assertion was successful (true) or not (false).

#### func [HTTPSuccessf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L830) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#HTTPSuccessf)added in v1.2.0

```
func HTTPSuccessf(t TestingT, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{})
```

HTTPSuccessf asserts that a specified handler returns a success status code.

```
require.HTTPSuccessf(t, myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func [Implements](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L843) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Implements)

```
func Implements(t TestingT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{})
```

Implements asserts that an object is implemented by the specified interface.

```
require.Implements(t, (*MyInterface)(nil), new(MyObject))
```

#### func [Implementsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L856) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Implementsf)added in v1.2.0

```
func Implementsf(t TestingT, interfaceObject interface{}, object interface{}, msg string, args ...interface{})
```

Implementsf asserts that an object is implemented by the specified interface.

```
require.Implementsf(t, (*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
```

#### func [InDelta](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L869) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#InDelta)

```
func InDelta(t TestingT, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{})
```

InDelta asserts that the two numerals are within delta of each other.

```
require.InDelta(t, math.Pi, 22/7.0, 0.01)
```

#### func [InDeltaMapValues](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L880) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#InDeltaMapValues)added in v1.2.0

```
func InDeltaMapValues(t TestingT, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{})
```

InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.

#### func [InDeltaMapValuesf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L891) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#InDeltaMapValuesf)added in v1.2.0

```
func InDeltaMapValuesf(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{})
```

InDeltaMapValuesf is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.

#### func [InDeltaSlice](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L902) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#InDeltaSlice)

```
func InDeltaSlice(t TestingT, expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{})
```

InDeltaSlice is the same as InDelta, except it compares two slices.

#### func [InDeltaSlicef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L913) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#InDeltaSlicef)added in v1.2.0

```
func InDeltaSlicef(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{})
```

InDeltaSlicef is the same as InDelta, except it compares two slices.

#### func [InDeltaf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L926) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#InDeltaf)added in v1.2.0

```
func InDeltaf(t TestingT, expected interface{}, actual interface{}, delta float64, msg string, args ...interface{})
```

InDeltaf asserts that the two numerals are within delta of each other.

```
require.InDeltaf(t, math.Pi, 22/7.0, 0.01, "error message %s", "formatted")
```

#### func [InEpsilon](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L937) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#InEpsilon)

```
func InEpsilon(t TestingT, expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{})
```

InEpsilon asserts that expected and actual have a relative error less than epsilon

#### func [InEpsilonSlice](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L948) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#InEpsilonSlice)

```
func InEpsilonSlice(t TestingT, expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{})
```

InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.

#### func [InEpsilonSlicef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L959) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#InEpsilonSlicef)added in v1.2.0

```
func InEpsilonSlicef(t TestingT, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{})
```

InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.

#### func [InEpsilonf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L970) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#InEpsilonf)added in v1.2.0

```
func InEpsilonf(t TestingT, expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{})
```

InEpsilonf asserts that expected and actual have a relative error less than epsilon

#### func [IsDecreasing](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L985) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#IsDecreasing)added in v1.7.0

```
func IsDecreasing(t TestingT, object interface{}, msgAndArgs ...interface{})
```

IsDecreasing asserts that the collection is decreasing

```
require.IsDecreasing(t, []int{2, 1, 0})
require.IsDecreasing(t, []float{2, 1})
require.IsDecreasing(t, []string{"b", "a"})
```

#### func [IsDecreasingf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1000) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#IsDecreasingf)added in v1.7.0

```
func IsDecreasingf(t TestingT, object interface{}, msg string, args ...interface{})
```

IsDecreasingf asserts that the collection is decreasing

```
require.IsDecreasingf(t, []int{2, 1, 0}, "error message %s", "formatted")
require.IsDecreasingf(t, []float{2, 1}, "error message %s", "formatted")
require.IsDecreasingf(t, []string{"b", "a"}, "error message %s", "formatted")
```

#### func [IsIncreasing](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1015) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#IsIncreasing)added in v1.7.0

```
func IsIncreasing(t TestingT, object interface{}, msgAndArgs ...interface{})
```

IsIncreasing asserts that the collection is increasing

```
require.IsIncreasing(t, []int{1, 2, 3})
require.IsIncreasing(t, []float{1, 2})
require.IsIncreasing(t, []string{"a", "b"})
```

#### func [IsIncreasingf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1030) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#IsIncreasingf)added in v1.7.0

```
func IsIncreasingf(t TestingT, object interface{}, msg string, args ...interface{})
```

IsIncreasingf asserts that the collection is increasing

```
require.IsIncreasingf(t, []int{1, 2, 3}, "error message %s", "formatted")
require.IsIncreasingf(t, []float{1, 2}, "error message %s", "formatted")
require.IsIncreasingf(t, []string{"a", "b"}, "error message %s", "formatted")
```

#### func [IsNonDecreasing](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1045) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#IsNonDecreasing)added in v1.7.0

```
func IsNonDecreasing(t TestingT, object interface{}, msgAndArgs ...interface{})
```

IsNonDecreasing asserts that the collection is not decreasing

```
require.IsNonDecreasing(t, []int{1, 1, 2})
require.IsNonDecreasing(t, []float{1, 2})
require.IsNonDecreasing(t, []string{"a", "b"})
```

#### func [IsNonDecreasingf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1060) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#IsNonDecreasingf)added in v1.7.0

```
func IsNonDecreasingf(t TestingT, object interface{}, msg string, args ...interface{})
```

IsNonDecreasingf asserts that the collection is not decreasing

```
require.IsNonDecreasingf(t, []int{1, 1, 2}, "error message %s", "formatted")
require.IsNonDecreasingf(t, []float{1, 2}, "error message %s", "formatted")
require.IsNonDecreasingf(t, []string{"a", "b"}, "error message %s", "formatted")
```

#### func [IsNonIncreasing](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1075) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#IsNonIncreasing)added in v1.7.0

```
func IsNonIncreasing(t TestingT, object interface{}, msgAndArgs ...interface{})
```

IsNonIncreasing asserts that the collection is not increasing

```
require.IsNonIncreasing(t, []int{2, 1, 1})
require.IsNonIncreasing(t, []float{2, 1})
require.IsNonIncreasing(t, []string{"b", "a"})
```

#### func [IsNonIncreasingf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1090) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#IsNonIncreasingf)added in v1.7.0

```
func IsNonIncreasingf(t TestingT, object interface{}, msg string, args ...interface{})
```

IsNonIncreasingf asserts that the collection is not increasing

```
require.IsNonIncreasingf(t, []int{2, 1, 1}, "error message %s", "formatted")
require.IsNonIncreasingf(t, []float{2, 1}, "error message %s", "formatted")
require.IsNonIncreasingf(t, []string{"b", "a"}, "error message %s", "formatted")
```

#### func [IsType](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1101) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#IsType)

```
func IsType(t TestingT, expectedType interface{}, object interface{}, msgAndArgs ...interface{})
```

IsType asserts that the specified objects are of the same type.

#### func [IsTypef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1112) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#IsTypef)added in v1.2.0

```
func IsTypef(t TestingT, expectedType interface{}, object interface{}, msg string, args ...interface{})
```

IsTypef asserts that the specified objects are of the same type.

#### func [JSONEq](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1125) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#JSONEq)

```
func JSONEq(t TestingT, expected string, actual string, msgAndArgs ...interface{})
```

JSONEq asserts that two JSON strings are equivalent.

```
require.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
```

#### func [JSONEqf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1138) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#JSONEqf)added in v1.2.0

```
func JSONEqf(t TestingT, expected string, actual string, msg string, args ...interface{})
```

JSONEqf asserts that two JSON strings are equivalent.

```
require.JSONEqf(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
```

#### func [Len](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1152) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Len)

```
func Len(t TestingT, object interface{}, length int, msgAndArgs ...interface{})
```

Len asserts that the specified object has specific length. Len also fails if the object has a type that len() not accept.

```
require.Len(t, mySlice, 3)
```

#### func [Lenf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1166) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Lenf)added in v1.2.0

```
func Lenf(t TestingT, object interface{}, length int, msg string, args ...interface{})
```

Lenf asserts that the specified object has specific length. Lenf also fails if the object has a type that len() not accept.

```
require.Lenf(t, mySlice, 3, "error message %s", "formatted")
```

#### func [Less](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1181) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Less)added in v1.4.0

```
func Less(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
```

Less asserts that the first element is less than the second

```
require.Less(t, 1, 2)
require.Less(t, float64(1), float64(2))
require.Less(t, "a", "b")
```

#### func [LessOrEqual](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1197) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#LessOrEqual)added in v1.4.0

```
func LessOrEqual(t TestingT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
```

LessOrEqual asserts that the first element is less than or equal to the second

```
require.LessOrEqual(t, 1, 2)
require.LessOrEqual(t, 2, 2)
require.LessOrEqual(t, "a", "b")
require.LessOrEqual(t, "b", "b")
```

#### func [LessOrEqualf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1213) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#LessOrEqualf)added in v1.4.0

```
func LessOrEqualf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{})
```

LessOrEqualf asserts that the first element is less than or equal to the second

```
require.LessOrEqualf(t, 1, 2, "error message %s", "formatted")
require.LessOrEqualf(t, 2, 2, "error message %s", "formatted")
require.LessOrEqualf(t, "a", "b", "error message %s", "formatted")
require.LessOrEqualf(t, "b", "b", "error message %s", "formatted")
```

#### func [Lessf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1228) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Lessf)added in v1.4.0

```
func Lessf(t TestingT, e1 interface{}, e2 interface{}, msg string, args ...interface{})
```

Lessf asserts that the first element is less than the second

```
require.Lessf(t, 1, 2, "error message %s", "formatted")
require.Lessf(t, float64(1), float64(2), "error message %s", "formatted")
require.Lessf(t, "a", "b", "error message %s", "formatted")
```

#### func [Negative](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1242) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Negative)added in v1.7.0

```
func Negative(t TestingT, e interface{}, msgAndArgs ...interface{})
```

Negative asserts that the specified element is negative

```
require.Negative(t, -1)
require.Negative(t, -1.23)
```

#### func [Negativef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1256) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Negativef)added in v1.7.0

```
func Negativef(t TestingT, e interface{}, msg string, args ...interface{})
```

Negativef asserts that the specified element is negative

```
require.Negativef(t, -1, "error message %s", "formatted")
require.Negativef(t, -1.23, "error message %s", "formatted")
```

#### func [Never](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1270) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Never)added in v1.5.0

```
func Never(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{})
```

Never asserts that the given condition doesn't satisfy in waitFor time, periodically checking the target function each tick.

```
require.Never(t, func() bool { return false; }, time.Second, 10*time.Millisecond)
```

#### func [Neverf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1284) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Neverf)added in v1.5.0

```
func Neverf(t TestingT, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{})
```

Neverf asserts that the given condition doesn't satisfy in waitFor time, periodically checking the target function each tick.

```
require.Neverf(t, func() bool { return false; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
```

#### func [Nil](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1297) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Nil)

```
func Nil(t TestingT, object interface{}, msgAndArgs ...interface{})
```

Nil asserts that the specified object is nil.

```
require.Nil(t, err)
```

#### func [Nilf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1310) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Nilf)added in v1.2.0

```
func Nilf(t TestingT, object interface{}, msg string, args ...interface{})
```

Nilf asserts that the specified object is nil.

```
require.Nilf(t, err, "error message %s", "formatted")
```

#### func [NoDirExists](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1322) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NoDirExists)added in v1.5.0

```
func NoDirExists(t TestingT, path string, msgAndArgs ...interface{})
```

NoDirExists checks whether a directory does not exist in the given path. It fails if the path points to an existing _directory_ only.

#### func [NoDirExistsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1334) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NoDirExistsf)added in v1.5.0

```
func NoDirExistsf(t TestingT, path string, msg string, args ...interface{})
```

NoDirExistsf checks whether a directory does not exist in the given path. It fails if the path points to an existing _directory_ only.

#### func [NoError](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1350) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NoError)

```
func NoError(t TestingT, err error, msgAndArgs ...interface{})
```

NoError asserts that a function returned no error (i.e. `nil`).

```
  actualObj, err := SomeFunction()
  if require.NoError(t, err) {
	   require.Equal(t, expectedObj, actualObj)
  }
```

#### func [NoErrorf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1366) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NoErrorf)added in v1.2.0

```
func NoErrorf(t TestingT, err error, msg string, args ...interface{})
```

NoErrorf asserts that a function returned no error (i.e. `nil`).

```
  actualObj, err := SomeFunction()
  if require.NoErrorf(t, err, "error message %s", "formatted") {
	   require.Equal(t, expectedObj, actualObj)
  }
```

#### func [NoFileExists](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1378) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NoFileExists)added in v1.5.0

```
func NoFileExists(t TestingT, path string, msgAndArgs ...interface{})
```

NoFileExists checks whether a file does not exist in a given path. It fails if the path points to an existing _file_ only.

#### func [NoFileExistsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1390) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NoFileExistsf)added in v1.5.0

```
func NoFileExistsf(t TestingT, path string, msg string, args ...interface{})
```

NoFileExistsf checks whether a file does not exist in a given path. It fails if the path points to an existing _file_ only.

#### func [NotContains](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1406) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotContains)

```
func NotContains(t TestingT, s interface{}, contains interface{}, msgAndArgs ...interface{})
```

NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the specified substring or element.

```
require.NotContains(t, "Hello World", "Earth")
require.NotContains(t, ["Hello", "World"], "Earth")
require.NotContains(t, {"Hello": "World"}, "Earth")
```

#### func [NotContainsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1422) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotContainsf)added in v1.2.0

```
func NotContainsf(t TestingT, s interface{}, contains interface{}, msg string, args ...interface{})
```

NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the specified substring or element.

```
require.NotContainsf(t, "Hello World", "Earth", "error message %s", "formatted")
require.NotContainsf(t, ["Hello", "World"], "Earth", "error message %s", "formatted")
require.NotContainsf(t, {"Hello": "World"}, "Earth", "error message %s", "formatted")
```

#### func [NotElementsMatch](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1442) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotElementsMatch)added in v1.10.0

```
func NotElementsMatch(t TestingT, listA interface{}, listB interface{}, msgAndArgs ...interface{})
```

NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should not match. This is an inverse of ElementsMatch.

require.NotElementsMatch(t, [1, 1, 2, 3], [1, 1, 2, 3]) -> false

require.NotElementsMatch(t, [1, 1, 2, 3], [1, 2, 3]) -> true

require.NotElementsMatch(t, [1, 2, 3], [1, 2, 4]) -> true

#### func [NotElementsMatchf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1462) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotElementsMatchf)added in v1.10.0

```
func NotElementsMatchf(t TestingT, listA interface{}, listB interface{}, msg string, args ...interface{})
```

NotElementsMatchf asserts that the specified listA(array, slice...) is NOT equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should not match. This is an inverse of ElementsMatch.

require.NotElementsMatchf(t, [1, 1, 2, 3], [1, 1, 2, 3], "error message %s", "formatted") -> false

require.NotElementsMatchf(t, [1, 1, 2, 3], [1, 2, 3], "error message %s", "formatted") -> true

require.NotElementsMatchf(t, [1, 2, 3], [1, 2, 4], "error message %s", "formatted") -> true

#### func [NotEmpty](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1478) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotEmpty)

```
func NotEmpty(t TestingT, object interface{}, msgAndArgs ...interface{})
```

NotEmpty asserts that the specified object is NOT empty. I.e. not nil, "", false, 0 or either a slice or a channel with len == 0.

```
if require.NotEmpty(t, obj) {
  require.Equal(t, "two", obj[1])
}
```

#### func [NotEmptyf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1494) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotEmptyf)added in v1.2.0

```
func NotEmptyf(t TestingT, object interface{}, msg string, args ...interface{})
```

NotEmptyf asserts that the specified object is NOT empty. I.e. not nil, "", false, 0 or either a slice or a channel with len == 0.

```
if require.NotEmptyf(t, obj, "error message %s", "formatted") {
  require.Equal(t, "two", obj[1])
}
```

#### func [NotEqual](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1510) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotEqual)

```
func NotEqual(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

NotEqual asserts that the specified values are NOT equal.

```
require.NotEqual(t, obj1, obj2)
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses).

#### func [NotEqualValues](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1523) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotEqualValues)added in v1.6.0

```
func NotEqualValues(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

NotEqualValues asserts that two objects are not equal even when converted to the same type

```
require.NotEqualValues(t, obj1, obj2)
```

#### func [NotEqualValuesf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1536) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotEqualValuesf)added in v1.6.0

```
func NotEqualValuesf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{})
```

NotEqualValuesf asserts that two objects are not equal even when converted to the same type

```
require.NotEqualValuesf(t, obj1, obj2, "error message %s", "formatted")
```

#### func [NotEqualf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1552) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotEqualf)added in v1.2.0

```
func NotEqualf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{})
```

NotEqualf asserts that the specified values are NOT equal.

```
require.NotEqualf(t, obj1, obj2, "error message %s", "formatted")
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses).

#### func [NotErrorAs](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1564) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotErrorAs)added in v1.10.0

```
func NotErrorAs(t TestingT, err error, target interface{}, msgAndArgs ...interface{})
```

NotErrorAs asserts that none of the errors in err's chain matches target, but if so, sets target to that error value.

#### func [NotErrorAsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1576) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotErrorAsf)added in v1.10.0

```
func NotErrorAsf(t TestingT, err error, target interface{}, msg string, args ...interface{})
```

NotErrorAsf asserts that none of the errors in err's chain matches target, but if so, sets target to that error value.

#### func [NotErrorIs](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1588) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotErrorIs)added in v1.7.0

```
func NotErrorIs(t TestingT, err error, target error, msgAndArgs ...interface{})
```

NotErrorIs asserts that none of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func [NotErrorIsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1600) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotErrorIsf)added in v1.7.0

```
func NotErrorIsf(t TestingT, err error, target error, msg string, args ...interface{})
```

NotErrorIsf asserts that none of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func [NotImplements](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1613) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotImplements)added in v1.9.0

```
func NotImplements(t TestingT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{})
```

NotImplements asserts that an object does not implement the specified interface.

```
require.NotImplements(t, (*MyInterface)(nil), new(MyObject))
```

#### func [NotImplementsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1626) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotImplementsf)added in v1.9.0

```
func NotImplementsf(t TestingT, interfaceObject interface{}, object interface{}, msg string, args ...interface{})
```

NotImplementsf asserts that an object does not implement the specified interface.

```
require.NotImplementsf(t, (*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
```

#### func [NotNil](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1639) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotNil)

```
func NotNil(t TestingT, object interface{}, msgAndArgs ...interface{})
```

NotNil asserts that the specified object is not nil.

```
require.NotNil(t, err)
```

#### func [NotNilf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1652) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotNilf)added in v1.2.0

```
func NotNilf(t TestingT, object interface{}, msg string, args ...interface{})
```

NotNilf asserts that the specified object is not nil.

```
require.NotNilf(t, err, "error message %s", "formatted")
```

#### func [NotPanics](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1665) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotPanics)

```
func NotPanics(t TestingT, f assert.PanicTestFunc, msgAndArgs ...interface{})
```

NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.

```
require.NotPanics(t, func(){ RemainCalm() })
```

#### func [NotPanicsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1678) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotPanicsf)added in v1.2.0

```
func NotPanicsf(t TestingT, f assert.PanicTestFunc, msg string, args ...interface{})
```

NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.

```
require.NotPanicsf(t, func(){ RemainCalm() }, "error message %s", "formatted")
```

#### func [NotRegexp](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1692) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotRegexp)

```
func NotRegexp(t TestingT, rx interface{}, str interface{}, msgAndArgs ...interface{})
```

NotRegexp asserts that a specified regexp does not match a string.

```
require.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
require.NotRegexp(t, "^start", "it's not starting")
```

#### func [NotRegexpf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1706) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotRegexpf)added in v1.2.0

```
func NotRegexpf(t TestingT, rx interface{}, str interface{}, msg string, args ...interface{})
```

NotRegexpf asserts that a specified regexp does not match a string.

```
require.NotRegexpf(t, regexp.MustCompile("starts"), "it's starting", "error message %s", "formatted")
require.NotRegexpf(t, "^start", "it's not starting", "error message %s", "formatted")
```

#### func [NotSame](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1722) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotSame)added in v1.5.0

```
func NotSame(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

NotSame asserts that two pointers do not reference the same object.

```
require.NotSame(t, ptr1, ptr2)
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func [NotSamef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1738) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotSamef)added in v1.5.0

```
func NotSamef(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{})
```

NotSamef asserts that two pointers do not reference the same object.

```
require.NotSamef(t, ptr1, ptr2, "error message %s", "formatted")
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func [NotSubset](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1754) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotSubset)added in v1.2.0

```
func NotSubset(t TestingT, list interface{}, subset interface{}, msgAndArgs ...interface{})
```

NotSubset asserts that the specified list(array, slice...) or map does NOT contain all elements given in the specified subset list(array, slice...) or map.

```
require.NotSubset(t, [1, 3, 4], [1, 2])
require.NotSubset(t, {"x": 1, "y": 2}, {"z": 3})
```

#### func [NotSubsetf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1770) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotSubsetf)added in v1.2.0

```
func NotSubsetf(t TestingT, list interface{}, subset interface{}, msg string, args ...interface{})
```

NotSubsetf asserts that the specified list(array, slice...) or map does NOT contain all elements given in the specified subset list(array, slice...) or map.

```
require.NotSubsetf(t, [1, 3, 4], [1, 2], "error message %s", "formatted")
require.NotSubsetf(t, {"x": 1, "y": 2}, {"z": 3}, "error message %s", "formatted")
```

#### func [NotZero](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1781) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotZero)

```
func NotZero(t TestingT, i interface{}, msgAndArgs ...interface{})
```

NotZero asserts that i is not the zero value for its type.

#### func [NotZerof](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1792) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#NotZerof)added in v1.2.0

```
func NotZerof(t TestingT, i interface{}, msg string, args ...interface{})
```

NotZerof asserts that i is not the zero value for its type.

#### func [Panics](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1805) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Panics)

```
func Panics(t TestingT, f assert.PanicTestFunc, msgAndArgs ...interface{})
```

Panics asserts that the code inside the specified PanicTestFunc panics.

```
require.Panics(t, func(){ GoCrazy() })
```

#### func [PanicsWithError](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1820) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#PanicsWithError)added in v1.5.0

```
func PanicsWithError(t TestingT, errString string, f assert.PanicTestFunc, msgAndArgs ...interface{})
```

PanicsWithError asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value is an error that satisfies the EqualError comparison.

```
require.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
```

#### func [PanicsWithErrorf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1835) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#PanicsWithErrorf)added in v1.5.0

```
func PanicsWithErrorf(t TestingT, errString string, f assert.PanicTestFunc, msg string, args ...interface{})
```

PanicsWithErrorf asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value is an error that satisfies the EqualError comparison.

```
require.PanicsWithErrorf(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func [PanicsWithValue](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1849) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#PanicsWithValue)added in v1.2.0

```
func PanicsWithValue(t TestingT, expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{})
```

PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value equals the expected panic value.

```
require.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
```

#### func [PanicsWithValuef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1863) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#PanicsWithValuef)added in v1.2.0

```
func PanicsWithValuef(t TestingT, expected interface{}, f assert.PanicTestFunc, msg string, args ...interface{})
```

PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value equals the expected panic value.

```
require.PanicsWithValuef(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func [Panicsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1876) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Panicsf)added in v1.2.0

```
func Panicsf(t TestingT, f assert.PanicTestFunc, msg string, args ...interface{})
```

Panicsf asserts that the code inside the specified PanicTestFunc panics.

```
require.Panicsf(t, func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func [Positive](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1890) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Positive)added in v1.7.0

```
func Positive(t TestingT, e interface{}, msgAndArgs ...interface{})
```

Positive asserts that the specified element is positive

```
require.Positive(t, 1)
require.Positive(t, 1.23)
```

#### func [Positivef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1904) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Positivef)added in v1.7.0

```
func Positivef(t TestingT, e interface{}, msg string, args ...interface{})
```

Positivef asserts that the specified element is positive

```
require.Positivef(t, 1, "error message %s", "formatted")
require.Positivef(t, 1.23, "error message %s", "formatted")
```

#### func [Regexp](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1918) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Regexp)

```
func Regexp(t TestingT, rx interface{}, str interface{}, msgAndArgs ...interface{})
```

Regexp asserts that a specified regexp matches a string.

```
require.Regexp(t, regexp.MustCompile("start"), "it's starting")
require.Regexp(t, "start...$", "it's not starting")
```

#### func [Regexpf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1932) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Regexpf)added in v1.2.0

```
func Regexpf(t TestingT, rx interface{}, str interface{}, msg string, args ...interface{})
```

Regexpf asserts that a specified regexp matches a string.

```
require.Regexpf(t, regexp.MustCompile("start"), "it's starting", "error message %s", "formatted")
require.Regexpf(t, "start...$", "it's not starting", "error message %s", "formatted")
```

#### func [Same](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1948) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Same)added in v1.4.0

```
func Same(t TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

Same asserts that two pointers reference the same object.

```
require.Same(t, ptr1, ptr2)
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func [Samef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1964) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Samef)added in v1.4.0

```
func Samef(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{})
```

Samef asserts that two pointers reference the same object.

```
require.Samef(t, ptr1, ptr2, "error message %s", "formatted")
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func [Subset](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1979) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Subset)added in v1.2.0

```
func Subset(t TestingT, list interface{}, subset interface{}, msgAndArgs ...interface{})
```

Subset asserts that the specified list(array, slice...) or map contains all elements given in the specified subset list(array, slice...) or map.

```
require.Subset(t, [1, 2, 3], [1, 2])
require.Subset(t, {"x": 1, "y": 2}, {"x": 1})
```

#### func [Subsetf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L1994) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Subsetf)added in v1.2.0

```
func Subsetf(t TestingT, list interface{}, subset interface{}, msg string, args ...interface{})
```

Subsetf asserts that the specified list(array, slice...) or map contains all elements given in the specified subset list(array, slice...) or map.

```
require.Subsetf(t, [1, 2, 3], [1, 2], "error message %s", "formatted")
require.Subsetf(t, {"x": 1, "y": 2}, {"x": 1}, "error message %s", "formatted")
```

#### func [True](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L2007) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#True)

```
func True(t TestingT, value bool, msgAndArgs ...interface{})
```

True asserts that the specified value is true.

```
require.True(t, myBool)
```

#### func [Truef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L2020) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Truef)added in v1.2.0

```
func Truef(t TestingT, value bool, msg string, args ...interface{})
```

Truef asserts that the specified value is true.

```
require.Truef(t, myBool, "error message %s", "formatted")
```

#### func [WithinDuration](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L2033) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#WithinDuration)

```
func WithinDuration(t TestingT, expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{})
```

WithinDuration asserts that the two times are within duration delta of each other.

```
require.WithinDuration(t, time.Now(), time.Now(), 10*time.Second)
```

#### func [WithinDurationf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L2046) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#WithinDurationf)added in v1.2.0

```
func WithinDurationf(t TestingT, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{})
```

WithinDurationf asserts that the two times are within duration delta of each other.

```
require.WithinDurationf(t, time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
```

#### func [WithinRange](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L2059) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#WithinRange)added in v1.8.0

```
func WithinRange(t TestingT, actual time.Time, start time.Time, end time.Time, msgAndArgs ...interface{})
```

WithinRange asserts that a time is within a time range (inclusive).

```
require.WithinRange(t, time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
```

#### func [WithinRangef](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L2072) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#WithinRangef)added in v1.8.0

```
func WithinRangef(t TestingT, actual time.Time, start time.Time, end time.Time, msg string, args ...interface{})
```

WithinRangef asserts that a time is within a time range (inclusive).

```
require.WithinRangef(t, time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second), "error message %s", "formatted")
```

#### func [YAMLEq](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L2083) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#YAMLEq)added in v1.4.0

```
func YAMLEq(t TestingT, expected string, actual string, msgAndArgs ...interface{})
```

YAMLEq asserts that two YAML strings are equivalent.

#### func [YAMLEqf](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L2094) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#YAMLEqf)added in v1.4.0

```
func YAMLEqf(t TestingT, expected string, actual string, msg string, args ...interface{})
```

YAMLEqf asserts that two YAML strings are equivalent.

#### func [Zero](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L2105) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Zero)

```
func Zero(t TestingT, i interface{}, msgAndArgs ...interface{})
```

Zero asserts that i is the zero value for its type.

#### func [Zerof](https://github.com/stretchr/testify/blob/v1.10.0/require/require.go#L2116) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Zerof)added in v1.2.0

```
func Zerof(t TestingT, i interface{}, msg string, args ...interface{})
```

Zerof asserts that i is the zero value for its type.

### Types [¶](https://pkg.go.dev/github.com/stretchr/testify/require#pkg-types)

#### type [Assertions](https://github.com/stretchr/testify/blob/v1.10.0/require/forward_requirements.go#L5) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions)

```
type Assertions struct {
	// contains filtered or unexported fields
}
```

Assertions provides assertion methods around the TestingT interface.

#### func [New](https://github.com/stretchr/testify/blob/v1.10.0/require/forward_requirements.go#L10) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#New)

```
func New(t TestingT) *Assertions
```

New makes a new Assertions object for the specified TestingT.

#### func (*Assertions) [Condition](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L13) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Condition)

```
func (a *Assertions) Condition(comp assert.Comparison, msgAndArgs ...interface{})
```

Condition uses a Comparison to assert a complex condition.

#### func (*Assertions) [Conditionf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L21) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Conditionf)added in v1.2.0

```
func (a *Assertions) Conditionf(comp assert.Comparison, msg string, args ...interface{})
```

Conditionf uses a Comparison to assert a complex condition.

#### func (*Assertions) [Contains](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L34) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Contains)

```
func (a *Assertions) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{})
```

Contains asserts that the specified string, list(array, slice...) or map contains the specified substring or element.

```
a.Contains("Hello World", "World")
a.Contains(["Hello", "World"], "World")
a.Contains({"Hello": "World"}, "Hello")
```

#### func (*Assertions) [Containsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L47) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Containsf)added in v1.2.0

```
func (a *Assertions) Containsf(s interface{}, contains interface{}, msg string, args ...interface{})
```

Containsf asserts that the specified string, list(array, slice...) or map contains the specified substring or element.

```
a.Containsf("Hello World", "World", "error message %s", "formatted")
a.Containsf(["Hello", "World"], "World", "error message %s", "formatted")
a.Containsf({"Hello": "World"}, "Hello", "error message %s", "formatted")
```

#### func (*Assertions) [DirExists](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L56) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.DirExists)added in v1.2.0

```
func (a *Assertions) DirExists(path string, msgAndArgs ...interface{})
```

DirExists checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.

#### func (*Assertions) [DirExistsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L65) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.DirExistsf)added in v1.2.0

```
func (a *Assertions) DirExistsf(path string, msg string, args ...interface{})
```

DirExistsf checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.

#### func (*Assertions) [ElementsMatch](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L77) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.ElementsMatch)added in v1.2.0

```
func (a *Assertions) ElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{})
```

ElementsMatch asserts that the specified listA(array, slice...) is equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should match.

a.ElementsMatch([1, 3, 2, 3], [1, 3, 3, 2])

#### func (*Assertions) [ElementsMatchf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L89) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.ElementsMatchf)added in v1.2.0

```
func (a *Assertions) ElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{})
```

ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should match.

a.ElementsMatchf([1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")

#### func (*Assertions) [Empty](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L100) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Empty)

```
func (a *Assertions) Empty(object interface{}, msgAndArgs ...interface{})
```

Empty asserts that the specified object is empty. I.e. nil, "", false, 0 or either a slice or a channel with len == 0.

```
a.Empty(obj)
```

#### func (*Assertions) [Emptyf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L111) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Emptyf)added in v1.2.0

```
func (a *Assertions) Emptyf(object interface{}, msg string, args ...interface{})
```

Emptyf asserts that the specified object is empty. I.e. nil, "", false, 0 or either a slice or a channel with len == 0.

```
a.Emptyf(obj, "error message %s", "formatted")
```

#### func (*Assertions) [Equal](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L125) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Equal)

```
func (a *Assertions) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

Equal asserts that two objects are equal.

```
a.Equal(123, 123)
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses). Function equality cannot be determined and will always fail.

#### func (*Assertions) [EqualError](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L137) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.EqualError)

```
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...interface{})
```

EqualError asserts that a function returned an error (i.e. not `nil`) and that it is equal to the provided error.

```
actualObj, err := SomeFunction()
a.EqualError(err,  expectedErrorString)
```

#### func (*Assertions) [EqualErrorf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L149) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.EqualErrorf)added in v1.2.0

```
func (a *Assertions) EqualErrorf(theError error, errString string, msg string, args ...interface{})
```

EqualErrorf asserts that a function returned an error (i.e. not `nil`) and that it is equal to the provided error.

```
actualObj, err := SomeFunction()
a.EqualErrorf(err,  expectedErrorString, "error message %s", "formatted")
```

#### func (*Assertions) [EqualExportedValues](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L166) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.EqualExportedValues)added in v1.8.3

```
func (a *Assertions) EqualExportedValues(expected interface{}, actual interface{}, msgAndArgs ...interface{})
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

#### func (*Assertions) [EqualExportedValuesf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L183) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.EqualExportedValuesf)added in v1.8.3

```
func (a *Assertions) EqualExportedValuesf(expected interface{}, actual interface{}, msg string, args ...interface{})
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

#### func (*Assertions) [EqualValues](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L194) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.EqualValues)

```
func (a *Assertions) EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

EqualValues asserts that two objects are equal or convertible to the larger type and equal.

```
a.EqualValues(uint32(123), int32(123))
```

#### func (*Assertions) [EqualValuesf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L205) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.EqualValuesf)added in v1.2.0

```
func (a *Assertions) EqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{})
```

EqualValuesf asserts that two objects are equal or convertible to the larger type and equal.

```
a.EqualValuesf(uint32(123), int32(123), "error message %s", "formatted")
```

#### func (*Assertions) [Equalf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L219) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Equalf)added in v1.2.0

```
func (a *Assertions) Equalf(expected interface{}, actual interface{}, msg string, args ...interface{})
```

Equalf asserts that two objects are equal.

```
a.Equalf(123, 123, "error message %s", "formatted")
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses). Function equality cannot be determined and will always fail.

#### func (*Assertions) [Error](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L232) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Error)

```
func (a *Assertions) Error(err error, msgAndArgs ...interface{})
```

Error asserts that a function returned an error (i.e. not `nil`).

```
  actualObj, err := SomeFunction()
  if a.Error(err) {
	   assert.Equal(t, expectedError, err)
  }
```

#### func (*Assertions) [ErrorAs](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L241) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.ErrorAs)added in v1.7.0

```
func (a *Assertions) ErrorAs(err error, target interface{}, msgAndArgs ...interface{})
```

ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value. This is a wrapper for errors.As.

#### func (*Assertions) [ErrorAsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L250) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.ErrorAsf)added in v1.7.0

```
func (a *Assertions) ErrorAsf(err error, target interface{}, msg string, args ...interface{})
```

ErrorAsf asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value. This is a wrapper for errors.As.

#### func (*Assertions) [ErrorContains](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L262) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.ErrorContains)added in v1.7.1

```
func (a *Assertions) ErrorContains(theError error, contains string, msgAndArgs ...interface{})
```

ErrorContains asserts that a function returned an error (i.e. not `nil`) and that the error contains the specified substring.

```
actualObj, err := SomeFunction()
a.ErrorContains(err,  expectedErrorSubString)
```

#### func (*Assertions) [ErrorContainsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L274) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.ErrorContainsf)added in v1.7.1

```
func (a *Assertions) ErrorContainsf(theError error, contains string, msg string, args ...interface{})
```

ErrorContainsf asserts that a function returned an error (i.e. not `nil`) and that the error contains the specified substring.

```
actualObj, err := SomeFunction()
a.ErrorContainsf(err,  expectedErrorSubString, "error message %s", "formatted")
```

#### func (*Assertions) [ErrorIs](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L283) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.ErrorIs)added in v1.7.0

```
func (a *Assertions) ErrorIs(err error, target error, msgAndArgs ...interface{})
```

ErrorIs asserts that at least one of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func (*Assertions) [ErrorIsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L292) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.ErrorIsf)added in v1.7.0

```
func (a *Assertions) ErrorIsf(err error, target error, msg string, args ...interface{})
```

ErrorIsf asserts that at least one of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func (*Assertions) [Errorf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L305) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Errorf)added in v1.2.0

```
func (a *Assertions) Errorf(err error, msg string, args ...interface{})
```

Errorf asserts that a function returned an error (i.e. not `nil`).

```
  actualObj, err := SomeFunction()
  if a.Errorf(err, "error message %s", "formatted") {
	   assert.Equal(t, expectedErrorf, err)
  }
```

#### func (*Assertions) [Eventually](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L316) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Eventually)added in v1.4.0

```
func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{})
```

Eventually asserts that given condition will be met in waitFor time, periodically checking target function each tick.

```
a.Eventually(func() bool { return true; }, time.Second, 10*time.Millisecond)
```

#### func (*Assertions) [EventuallyWithT](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L341) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.EventuallyWithT)added in v1.8.3

```
func (a *Assertions) EventuallyWithT(condition func(collect *assert.CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{})
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

#### func (*Assertions) [EventuallyWithTf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L366) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.EventuallyWithTf)added in v1.8.3

```
func (a *Assertions) EventuallyWithTf(condition func(collect *assert.CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...interface{})
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

#### func (*Assertions) [Eventuallyf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L377) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Eventuallyf)added in v1.4.0

```
func (a *Assertions) Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{})
```

Eventuallyf asserts that given condition will be met in waitFor time, periodically checking target function each tick.

```
a.Eventuallyf(func() bool { return true; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
```

#### func (*Assertions) [Exactly](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L387) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Exactly)

```
func (a *Assertions) Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

Exactly asserts that two objects are equal in value and type.

```
a.Exactly(int32(123), int64(123))
```

#### func (*Assertions) [Exactlyf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L397) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Exactlyf)added in v1.2.0

```
func (a *Assertions) Exactlyf(expected interface{}, actual interface{}, msg string, args ...interface{})
```

Exactlyf asserts that two objects are equal in value and type.

```
a.Exactlyf(int32(123), int64(123), "error message %s", "formatted")
```

#### func (*Assertions) [Fail](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L405) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Fail)

```
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...interface{})
```

Fail reports a failure through

#### func (*Assertions) [FailNow](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L413) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.FailNow)

```
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...interface{})
```

FailNow fails test

#### func (*Assertions) [FailNowf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L421) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.FailNowf)added in v1.2.0

```
func (a *Assertions) FailNowf(failureMessage string, msg string, args ...interface{})
```

FailNowf fails test

#### func (*Assertions) [Failf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L429) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Failf)added in v1.2.0

```
func (a *Assertions) Failf(failureMessage string, msg string, args ...interface{})
```

Failf reports a failure through

#### func (*Assertions) [False](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L439) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.False)

```
func (a *Assertions) False(value bool, msgAndArgs ...interface{})
```

False asserts that the specified value is false.

```
a.False(myBool)
```

#### func (*Assertions) [Falsef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L449) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Falsef)added in v1.2.0

```
func (a *Assertions) Falsef(value bool, msg string, args ...interface{})
```

Falsef asserts that the specified value is false.

```
a.Falsef(myBool, "error message %s", "formatted")
```

#### func (*Assertions) [FileExists](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L458) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.FileExists)added in v1.2.0

```
func (a *Assertions) FileExists(path string, msgAndArgs ...interface{})
```

FileExists checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.

#### func (*Assertions) [FileExistsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L467) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.FileExistsf)added in v1.2.0

```
func (a *Assertions) FileExistsf(path string, msg string, args ...interface{})
```

FileExistsf checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.

#### func (*Assertions) [Greater](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L479) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Greater)added in v1.4.0

```
func (a *Assertions) Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
```

Greater asserts that the first element is greater than the second

```
a.Greater(2, 1)
a.Greater(float64(2), float64(1))
a.Greater("b", "a")
```

#### func (*Assertions) [GreaterOrEqual](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L492) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.GreaterOrEqual)added in v1.4.0

```
func (a *Assertions) GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
```

GreaterOrEqual asserts that the first element is greater than or equal to the second

```
a.GreaterOrEqual(2, 1)
a.GreaterOrEqual(2, 2)
a.GreaterOrEqual("b", "a")
a.GreaterOrEqual("b", "b")
```

#### func (*Assertions) [GreaterOrEqualf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L505) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.GreaterOrEqualf)added in v1.4.0

```
func (a *Assertions) GreaterOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{})
```

GreaterOrEqualf asserts that the first element is greater than or equal to the second

```
a.GreaterOrEqualf(2, 1, "error message %s", "formatted")
a.GreaterOrEqualf(2, 2, "error message %s", "formatted")
a.GreaterOrEqualf("b", "a", "error message %s", "formatted")
a.GreaterOrEqualf("b", "b", "error message %s", "formatted")
```

#### func (*Assertions) [Greaterf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L517) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Greaterf)added in v1.4.0

```
func (a *Assertions) Greaterf(e1 interface{}, e2 interface{}, msg string, args ...interface{})
```

Greaterf asserts that the first element is greater than the second

```
a.Greaterf(2, 1, "error message %s", "formatted")
a.Greaterf(float64(2), float64(1), "error message %s", "formatted")
a.Greaterf("b", "a", "error message %s", "formatted")
```

#### func (*Assertions) [HTTPBodyContains](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L530) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPBodyContains)

```
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{})
```

HTTPBodyContains asserts that a specified handler returns a body that contains a string.

```
a.HTTPBodyContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPBodyContainsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L543) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPBodyContainsf)added in v1.2.0

```
func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{})
```

HTTPBodyContainsf asserts that a specified handler returns a body that contains a string.

```
a.HTTPBodyContainsf(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPBodyNotContains](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L556) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPBodyNotContains)

```
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msgAndArgs ...interface{})
```

HTTPBodyNotContains asserts that a specified handler returns a body that does not contain a string.

```
a.HTTPBodyNotContains(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPBodyNotContainsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L569) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPBodyNotContainsf)added in v1.2.0

```
func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{})
```

HTTPBodyNotContainsf asserts that a specified handler returns a body that does not contain a string.

```
a.HTTPBodyNotContainsf(myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPError](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L581) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPError)

```
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{})
```

HTTPError asserts that a specified handler returns an error status code.

```
a.HTTPError(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPErrorf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L593) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPErrorf)added in v1.2.0

```
func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{})
```

HTTPErrorf asserts that a specified handler returns an error status code.

```
a.HTTPErrorf(myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPRedirect](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L605) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPRedirect)

```
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{})
```

HTTPRedirect asserts that a specified handler returns a redirect status code.

```
a.HTTPRedirect(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPRedirectf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L617) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPRedirectf)added in v1.2.0

```
func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{})
```

HTTPRedirectf asserts that a specified handler returns a redirect status code.

```
a.HTTPRedirectf(myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPStatusCode](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L629) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPStatusCode)added in v1.6.0

```
func (a *Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...interface{})
```

HTTPStatusCode asserts that a specified handler returns a specified status code.

```
a.HTTPStatusCode(myHandler, "GET", "/notImplemented", nil, 501)
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPStatusCodef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L641) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPStatusCodef)added in v1.6.0

```
func (a *Assertions) HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{})
```

HTTPStatusCodef asserts that a specified handler returns a specified status code.

```
a.HTTPStatusCodef(myHandler, "GET", "/notImplemented", nil, 501, "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPSuccess](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L653) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPSuccess)

```
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...interface{})
```

HTTPSuccess asserts that a specified handler returns a success status code.

```
a.HTTPSuccess(myHandler, "POST", "http://www.google.com", nil)
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [HTTPSuccessf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L665) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.HTTPSuccessf)added in v1.2.0

```
func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{})
```

HTTPSuccessf asserts that a specified handler returns a success status code.

```
a.HTTPSuccessf(myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
```

Returns whether the assertion was successful (true) or not (false).

#### func (*Assertions) [Implements](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L675) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Implements)

```
func (a *Assertions) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{})
```

Implements asserts that an object is implemented by the specified interface.

```
a.Implements((*MyInterface)(nil), new(MyObject))
```

#### func (*Assertions) [Implementsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L685) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Implementsf)added in v1.2.0

```
func (a *Assertions) Implementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{})
```

Implementsf asserts that an object is implemented by the specified interface.

```
a.Implementsf((*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
```

#### func (*Assertions) [InDelta](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L695) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.InDelta)

```
func (a *Assertions) InDelta(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{})
```

InDelta asserts that the two numerals are within delta of each other.

```
a.InDelta(math.Pi, 22/7.0, 0.01)
```

#### func (*Assertions) [InDeltaMapValues](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L703) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.InDeltaMapValues)added in v1.2.0

```
func (a *Assertions) InDeltaMapValues(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{})
```

InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.

#### func (*Assertions) [InDeltaMapValuesf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L711) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.InDeltaMapValuesf)added in v1.2.0

```
func (a *Assertions) InDeltaMapValuesf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{})
```

InDeltaMapValuesf is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.

#### func (*Assertions) [InDeltaSlice](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L719) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.InDeltaSlice)

```
func (a *Assertions) InDeltaSlice(expected interface{}, actual interface{}, delta float64, msgAndArgs ...interface{})
```

InDeltaSlice is the same as InDelta, except it compares two slices.

#### func (*Assertions) [InDeltaSlicef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L727) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.InDeltaSlicef)added in v1.2.0

```
func (a *Assertions) InDeltaSlicef(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{})
```

InDeltaSlicef is the same as InDelta, except it compares two slices.

#### func (*Assertions) [InDeltaf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L737) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.InDeltaf)added in v1.2.0

```
func (a *Assertions) InDeltaf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{})
```

InDeltaf asserts that the two numerals are within delta of each other.

```
a.InDeltaf(math.Pi, 22/7.0, 0.01, "error message %s", "formatted")
```

#### func (*Assertions) [InEpsilon](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L745) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.InEpsilon)

```
func (a *Assertions) InEpsilon(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{})
```

InEpsilon asserts that expected and actual have a relative error less than epsilon

#### func (*Assertions) [InEpsilonSlice](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L753) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.InEpsilonSlice)

```
func (a *Assertions) InEpsilonSlice(expected interface{}, actual interface{}, epsilon float64, msgAndArgs ...interface{})
```

InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.

#### func (*Assertions) [InEpsilonSlicef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L761) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.InEpsilonSlicef)added in v1.2.0

```
func (a *Assertions) InEpsilonSlicef(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{})
```

InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.

#### func (*Assertions) [InEpsilonf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L769) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.InEpsilonf)added in v1.2.0

```
func (a *Assertions) InEpsilonf(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{})
```

InEpsilonf asserts that expected and actual have a relative error less than epsilon

#### func (*Assertions) [IsDecreasing](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L781) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.IsDecreasing)added in v1.7.0

```
func (a *Assertions) IsDecreasing(object interface{}, msgAndArgs ...interface{})
```

IsDecreasing asserts that the collection is decreasing

```
a.IsDecreasing([]int{2, 1, 0})
a.IsDecreasing([]float{2, 1})
a.IsDecreasing([]string{"b", "a"})
```

#### func (*Assertions) [IsDecreasingf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L793) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.IsDecreasingf)added in v1.7.0

```
func (a *Assertions) IsDecreasingf(object interface{}, msg string, args ...interface{})
```

IsDecreasingf asserts that the collection is decreasing

```
a.IsDecreasingf([]int{2, 1, 0}, "error message %s", "formatted")
a.IsDecreasingf([]float{2, 1}, "error message %s", "formatted")
a.IsDecreasingf([]string{"b", "a"}, "error message %s", "formatted")
```

#### func (*Assertions) [IsIncreasing](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L805) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.IsIncreasing)added in v1.7.0

```
func (a *Assertions) IsIncreasing(object interface{}, msgAndArgs ...interface{})
```

IsIncreasing asserts that the collection is increasing

```
a.IsIncreasing([]int{1, 2, 3})
a.IsIncreasing([]float{1, 2})
a.IsIncreasing([]string{"a", "b"})
```

#### func (*Assertions) [IsIncreasingf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L817) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.IsIncreasingf)added in v1.7.0

```
func (a *Assertions) IsIncreasingf(object interface{}, msg string, args ...interface{})
```

IsIncreasingf asserts that the collection is increasing

```
a.IsIncreasingf([]int{1, 2, 3}, "error message %s", "formatted")
a.IsIncreasingf([]float{1, 2}, "error message %s", "formatted")
a.IsIncreasingf([]string{"a", "b"}, "error message %s", "formatted")
```

#### func (*Assertions) [IsNonDecreasing](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L829) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.IsNonDecreasing)added in v1.7.0

```
func (a *Assertions) IsNonDecreasing(object interface{}, msgAndArgs ...interface{})
```

IsNonDecreasing asserts that the collection is not decreasing

```
a.IsNonDecreasing([]int{1, 1, 2})
a.IsNonDecreasing([]float{1, 2})
a.IsNonDecreasing([]string{"a", "b"})
```

#### func (*Assertions) [IsNonDecreasingf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L841) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.IsNonDecreasingf)added in v1.7.0

```
func (a *Assertions) IsNonDecreasingf(object interface{}, msg string, args ...interface{})
```

IsNonDecreasingf asserts that the collection is not decreasing

```
a.IsNonDecreasingf([]int{1, 1, 2}, "error message %s", "formatted")
a.IsNonDecreasingf([]float{1, 2}, "error message %s", "formatted")
a.IsNonDecreasingf([]string{"a", "b"}, "error message %s", "formatted")
```

#### func (*Assertions) [IsNonIncreasing](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L853) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.IsNonIncreasing)added in v1.7.0

```
func (a *Assertions) IsNonIncreasing(object interface{}, msgAndArgs ...interface{})
```

IsNonIncreasing asserts that the collection is not increasing

```
a.IsNonIncreasing([]int{2, 1, 1})
a.IsNonIncreasing([]float{2, 1})
a.IsNonIncreasing([]string{"b", "a"})
```

#### func (*Assertions) [IsNonIncreasingf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L865) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.IsNonIncreasingf)added in v1.7.0

```
func (a *Assertions) IsNonIncreasingf(object interface{}, msg string, args ...interface{})
```

IsNonIncreasingf asserts that the collection is not increasing

```
a.IsNonIncreasingf([]int{2, 1, 1}, "error message %s", "formatted")
a.IsNonIncreasingf([]float{2, 1}, "error message %s", "formatted")
a.IsNonIncreasingf([]string{"b", "a"}, "error message %s", "formatted")
```

#### func (*Assertions) [IsType](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L873) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.IsType)

```
func (a *Assertions) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{})
```

IsType asserts that the specified objects are of the same type.

#### func (*Assertions) [IsTypef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L881) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.IsTypef)added in v1.2.0

```
func (a *Assertions) IsTypef(expectedType interface{}, object interface{}, msg string, args ...interface{})
```

IsTypef asserts that the specified objects are of the same type.

#### func (*Assertions) [JSONEq](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L891) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.JSONEq)

```
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{})
```

JSONEq asserts that two JSON strings are equivalent.

```
a.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
```

#### func (*Assertions) [JSONEqf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L901) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.JSONEqf)added in v1.2.0

```
func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...interface{})
```

JSONEqf asserts that two JSON strings are equivalent.

```
a.JSONEqf(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
```

#### func (*Assertions) [Len](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L912) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Len)

```
func (a *Assertions) Len(object interface{}, length int, msgAndArgs ...interface{})
```

Len asserts that the specified object has specific length. Len also fails if the object has a type that len() not accept.

```
a.Len(mySlice, 3)
```

#### func (*Assertions) [Lenf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L923) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Lenf)added in v1.2.0

```
func (a *Assertions) Lenf(object interface{}, length int, msg string, args ...interface{})
```

Lenf asserts that the specified object has specific length. Lenf also fails if the object has a type that len() not accept.

```
a.Lenf(mySlice, 3, "error message %s", "formatted")
```

#### func (*Assertions) [Less](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L935) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Less)added in v1.4.0

```
func (a *Assertions) Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
```

Less asserts that the first element is less than the second

```
a.Less(1, 2)
a.Less(float64(1), float64(2))
a.Less("a", "b")
```

#### func (*Assertions) [LessOrEqual](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L948) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.LessOrEqual)added in v1.4.0

```
func (a *Assertions) LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
```

LessOrEqual asserts that the first element is less than or equal to the second

```
a.LessOrEqual(1, 2)
a.LessOrEqual(2, 2)
a.LessOrEqual("a", "b")
a.LessOrEqual("b", "b")
```

#### func (*Assertions) [LessOrEqualf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L961) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.LessOrEqualf)added in v1.4.0

```
func (a *Assertions) LessOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{})
```

LessOrEqualf asserts that the first element is less than or equal to the second

```
a.LessOrEqualf(1, 2, "error message %s", "formatted")
a.LessOrEqualf(2, 2, "error message %s", "formatted")
a.LessOrEqualf("a", "b", "error message %s", "formatted")
a.LessOrEqualf("b", "b", "error message %s", "formatted")
```

#### func (*Assertions) [Lessf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L973) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Lessf)added in v1.4.0

```
func (a *Assertions) Lessf(e1 interface{}, e2 interface{}, msg string, args ...interface{})
```

Lessf asserts that the first element is less than the second

```
a.Lessf(1, 2, "error message %s", "formatted")
a.Lessf(float64(1), float64(2), "error message %s", "formatted")
a.Lessf("a", "b", "error message %s", "formatted")
```

#### func (*Assertions) [Negative](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L984) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Negative)added in v1.7.0

```
func (a *Assertions) Negative(e interface{}, msgAndArgs ...interface{})
```

Negative asserts that the specified element is negative

```
a.Negative(-1)
a.Negative(-1.23)
```

#### func (*Assertions) [Negativef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L995) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Negativef)added in v1.7.0

```
func (a *Assertions) Negativef(e interface{}, msg string, args ...interface{})
```

Negativef asserts that the specified element is negative

```
a.Negativef(-1, "error message %s", "formatted")
a.Negativef(-1.23, "error message %s", "formatted")
```

#### func (*Assertions) [Never](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1006) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Never)added in v1.5.0

```
func (a *Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{})
```

Never asserts that the given condition doesn't satisfy in waitFor time, periodically checking the target function each tick.

```
a.Never(func() bool { return false; }, time.Second, 10*time.Millisecond)
```

#### func (*Assertions) [Neverf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1017) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Neverf)added in v1.5.0

```
func (a *Assertions) Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{})
```

Neverf asserts that the given condition doesn't satisfy in waitFor time, periodically checking the target function each tick.

```
a.Neverf(func() bool { return false; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
```

#### func (*Assertions) [Nil](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1027) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Nil)

```
func (a *Assertions) Nil(object interface{}, msgAndArgs ...interface{})
```

Nil asserts that the specified object is nil.

```
a.Nil(err)
```

#### func (*Assertions) [Nilf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1037) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Nilf)added in v1.2.0

```
func (a *Assertions) Nilf(object interface{}, msg string, args ...interface{})
```

Nilf asserts that the specified object is nil.

```
a.Nilf(err, "error message %s", "formatted")
```

#### func (*Assertions) [NoDirExists](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1046) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NoDirExists)added in v1.5.0

```
func (a *Assertions) NoDirExists(path string, msgAndArgs ...interface{})
```

NoDirExists checks whether a directory does not exist in the given path. It fails if the path points to an existing _directory_ only.

#### func (*Assertions) [NoDirExistsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1055) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NoDirExistsf)added in v1.5.0

```
func (a *Assertions) NoDirExistsf(path string, msg string, args ...interface{})
```

NoDirExistsf checks whether a directory does not exist in the given path. It fails if the path points to an existing _directory_ only.

#### func (*Assertions) [NoError](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1068) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NoError)

```
func (a *Assertions) NoError(err error, msgAndArgs ...interface{})
```

NoError asserts that a function returned no error (i.e. `nil`).

```
  actualObj, err := SomeFunction()
  if a.NoError(err) {
	   assert.Equal(t, expectedObj, actualObj)
  }
```

#### func (*Assertions) [NoErrorf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1081) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NoErrorf)added in v1.2.0

```
func (a *Assertions) NoErrorf(err error, msg string, args ...interface{})
```

NoErrorf asserts that a function returned no error (i.e. `nil`).

```
  actualObj, err := SomeFunction()
  if a.NoErrorf(err, "error message %s", "formatted") {
	   assert.Equal(t, expectedObj, actualObj)
  }
```

#### func (*Assertions) [NoFileExists](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1090) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NoFileExists)added in v1.5.0

```
func (a *Assertions) NoFileExists(path string, msgAndArgs ...interface{})
```

NoFileExists checks whether a file does not exist in a given path. It fails if the path points to an existing _file_ only.

#### func (*Assertions) [NoFileExistsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1099) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NoFileExistsf)added in v1.5.0

```
func (a *Assertions) NoFileExistsf(path string, msg string, args ...interface{})
```

NoFileExistsf checks whether a file does not exist in a given path. It fails if the path points to an existing _file_ only.

#### func (*Assertions) [NotContains](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1112) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotContains)

```
func (a *Assertions) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{})
```

NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the specified substring or element.

```
a.NotContains("Hello World", "Earth")
a.NotContains(["Hello", "World"], "Earth")
a.NotContains({"Hello": "World"}, "Earth")
```

#### func (*Assertions) [NotContainsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1125) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotContainsf)added in v1.2.0

```
func (a *Assertions) NotContainsf(s interface{}, contains interface{}, msg string, args ...interface{})
```

NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the specified substring or element.

```
a.NotContainsf("Hello World", "Earth", "error message %s", "formatted")
a.NotContainsf(["Hello", "World"], "Earth", "error message %s", "formatted")
a.NotContainsf({"Hello": "World"}, "Earth", "error message %s", "formatted")
```

#### func (*Assertions) [NotElementsMatch](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1142) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotElementsMatch)added in v1.10.0

```
func (a *Assertions) NotElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{})
```

NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should not match. This is an inverse of ElementsMatch.

a.NotElementsMatch([1, 1, 2, 3], [1, 1, 2, 3]) -> false

a.NotElementsMatch([1, 1, 2, 3], [1, 2, 3]) -> true

a.NotElementsMatch([1, 2, 3], [1, 2, 4]) -> true

#### func (*Assertions) [NotElementsMatchf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1159) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotElementsMatchf)added in v1.10.0

```
func (a *Assertions) NotElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{})
```

NotElementsMatchf asserts that the specified listA(array, slice...) is NOT equal to specified listB(array, slice...) ignoring the order of the elements. If there are duplicate elements, the number of appearances of each of them in both lists should not match. This is an inverse of ElementsMatch.

a.NotElementsMatchf([1, 1, 2, 3], [1, 1, 2, 3], "error message %s", "formatted") -> false

a.NotElementsMatchf([1, 1, 2, 3], [1, 2, 3], "error message %s", "formatted") -> true

a.NotElementsMatchf([1, 2, 3], [1, 2, 4], "error message %s", "formatted") -> true

#### func (*Assertions) [NotEmpty](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1172) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotEmpty)

```
func (a *Assertions) NotEmpty(object interface{}, msgAndArgs ...interface{})
```

NotEmpty asserts that the specified object is NOT empty. I.e. not nil, "", false, 0 or either a slice or a channel with len == 0.

```
if a.NotEmpty(obj) {
  assert.Equal(t, "two", obj[1])
}
```

#### func (*Assertions) [NotEmptyf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1185) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotEmptyf)added in v1.2.0

```
func (a *Assertions) NotEmptyf(object interface{}, msg string, args ...interface{})
```

NotEmptyf asserts that the specified object is NOT empty. I.e. not nil, "", false, 0 or either a slice or a channel with len == 0.

```
if a.NotEmptyf(obj, "error message %s", "formatted") {
  assert.Equal(t, "two", obj[1])
}
```

#### func (*Assertions) [NotEqual](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1198) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotEqual)

```
func (a *Assertions) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

NotEqual asserts that the specified values are NOT equal.

```
a.NotEqual(obj1, obj2)
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses).

#### func (*Assertions) [NotEqualValues](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1208) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotEqualValues)added in v1.6.0

```
func (a *Assertions) NotEqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

NotEqualValues asserts that two objects are not equal even when converted to the same type

```
a.NotEqualValues(obj1, obj2)
```

#### func (*Assertions) [NotEqualValuesf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1218) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotEqualValuesf)added in v1.6.0

```
func (a *Assertions) NotEqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{})
```

NotEqualValuesf asserts that two objects are not equal even when converted to the same type

```
a.NotEqualValuesf(obj1, obj2, "error message %s", "formatted")
```

#### func (*Assertions) [NotEqualf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1231) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotEqualf)added in v1.2.0

```
func (a *Assertions) NotEqualf(expected interface{}, actual interface{}, msg string, args ...interface{})
```

NotEqualf asserts that the specified values are NOT equal.

```
a.NotEqualf(obj1, obj2, "error message %s", "formatted")
```

Pointer variable equality is determined based on the equality of the referenced values (as opposed to the memory addresses).

#### func (*Assertions) [NotErrorAs](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1240) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotErrorAs)added in v1.10.0

```
func (a *Assertions) NotErrorAs(err error, target interface{}, msgAndArgs ...interface{})
```

NotErrorAs asserts that none of the errors in err's chain matches target, but if so, sets target to that error value.

#### func (*Assertions) [NotErrorAsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1249) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotErrorAsf)added in v1.10.0

```
func (a *Assertions) NotErrorAsf(err error, target interface{}, msg string, args ...interface{})
```

NotErrorAsf asserts that none of the errors in err's chain matches target, but if so, sets target to that error value.

#### func (*Assertions) [NotErrorIs](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1258) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotErrorIs)added in v1.7.0

```
func (a *Assertions) NotErrorIs(err error, target error, msgAndArgs ...interface{})
```

NotErrorIs asserts that none of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func (*Assertions) [NotErrorIsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1267) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotErrorIsf)added in v1.7.0

```
func (a *Assertions) NotErrorIsf(err error, target error, msg string, args ...interface{})
```

NotErrorIsf asserts that none of the errors in err's chain matches target. This is a wrapper for errors.Is.

#### func (*Assertions) [NotImplements](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1277) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotImplements)added in v1.9.0

```
func (a *Assertions) NotImplements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{})
```

NotImplements asserts that an object does not implement the specified interface.

```
a.NotImplements((*MyInterface)(nil), new(MyObject))
```

#### func (*Assertions) [NotImplementsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1287) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotImplementsf)added in v1.9.0

```
func (a *Assertions) NotImplementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{})
```

NotImplementsf asserts that an object does not implement the specified interface.

```
a.NotImplementsf((*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
```

#### func (*Assertions) [NotNil](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1297) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotNil)

```
func (a *Assertions) NotNil(object interface{}, msgAndArgs ...interface{})
```

NotNil asserts that the specified object is not nil.

```
a.NotNil(err)
```

#### func (*Assertions) [NotNilf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1307) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotNilf)added in v1.2.0

```
func (a *Assertions) NotNilf(object interface{}, msg string, args ...interface{})
```

NotNilf asserts that the specified object is not nil.

```
a.NotNilf(err, "error message %s", "formatted")
```

#### func (*Assertions) [NotPanics](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1317) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotPanics)

```
func (a *Assertions) NotPanics(f assert.PanicTestFunc, msgAndArgs ...interface{})
```

NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.

```
a.NotPanics(func(){ RemainCalm() })
```

#### func (*Assertions) [NotPanicsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1327) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotPanicsf)added in v1.2.0

```
func (a *Assertions) NotPanicsf(f assert.PanicTestFunc, msg string, args ...interface{})
```

NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.

```
a.NotPanicsf(func(){ RemainCalm() }, "error message %s", "formatted")
```

#### func (*Assertions) [NotRegexp](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1338) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotRegexp)

```
func (a *Assertions) NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{})
```

NotRegexp asserts that a specified regexp does not match a string.

```
a.NotRegexp(regexp.MustCompile("starts"), "it's starting")
a.NotRegexp("^start", "it's not starting")
```

#### func (*Assertions) [NotRegexpf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1349) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotRegexpf)added in v1.2.0

```
func (a *Assertions) NotRegexpf(rx interface{}, str interface{}, msg string, args ...interface{})
```

NotRegexpf asserts that a specified regexp does not match a string.

```
a.NotRegexpf(regexp.MustCompile("starts"), "it's starting", "error message %s", "formatted")
a.NotRegexpf("^start", "it's not starting", "error message %s", "formatted")
```

#### func (*Assertions) [NotSame](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1362) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotSame)added in v1.5.0

```
func (a *Assertions) NotSame(expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

NotSame asserts that two pointers do not reference the same object.

```
a.NotSame(ptr1, ptr2)
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func (*Assertions) [NotSamef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1375) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotSamef)added in v1.5.0

```
func (a *Assertions) NotSamef(expected interface{}, actual interface{}, msg string, args ...interface{})
```

NotSamef asserts that two pointers do not reference the same object.

```
a.NotSamef(ptr1, ptr2, "error message %s", "formatted")
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func (*Assertions) [NotSubset](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1388) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotSubset)added in v1.2.0

```
func (a *Assertions) NotSubset(list interface{}, subset interface{}, msgAndArgs ...interface{})
```

NotSubset asserts that the specified list(array, slice...) or map does NOT contain all elements given in the specified subset list(array, slice...) or map.

```
a.NotSubset([1, 3, 4], [1, 2])
a.NotSubset({"x": 1, "y": 2}, {"z": 3})
```

#### func (*Assertions) [NotSubsetf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1401) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotSubsetf)added in v1.2.0

```
func (a *Assertions) NotSubsetf(list interface{}, subset interface{}, msg string, args ...interface{})
```

NotSubsetf asserts that the specified list(array, slice...) or map does NOT contain all elements given in the specified subset list(array, slice...) or map.

```
a.NotSubsetf([1, 3, 4], [1, 2], "error message %s", "formatted")
a.NotSubsetf({"x": 1, "y": 2}, {"z": 3}, "error message %s", "formatted")
```

#### func (*Assertions) [NotZero](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1409) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotZero)

```
func (a *Assertions) NotZero(i interface{}, msgAndArgs ...interface{})
```

NotZero asserts that i is not the zero value for its type.

#### func (*Assertions) [NotZerof](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1417) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.NotZerof)added in v1.2.0

```
func (a *Assertions) NotZerof(i interface{}, msg string, args ...interface{})
```

NotZerof asserts that i is not the zero value for its type.

#### func (*Assertions) [Panics](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1427) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Panics)

```
func (a *Assertions) Panics(f assert.PanicTestFunc, msgAndArgs ...interface{})
```

Panics asserts that the code inside the specified PanicTestFunc panics.

```
a.Panics(func(){ GoCrazy() })
```

#### func (*Assertions) [PanicsWithError](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1439) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.PanicsWithError)added in v1.5.0

```
func (a *Assertions) PanicsWithError(errString string, f assert.PanicTestFunc, msgAndArgs ...interface{})
```

PanicsWithError asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value is an error that satisfies the EqualError comparison.

```
a.PanicsWithError("crazy error", func(){ GoCrazy() })
```

#### func (*Assertions) [PanicsWithErrorf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1451) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.PanicsWithErrorf)added in v1.5.0

```
func (a *Assertions) PanicsWithErrorf(errString string, f assert.PanicTestFunc, msg string, args ...interface{})
```

PanicsWithErrorf asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value is an error that satisfies the EqualError comparison.

```
a.PanicsWithErrorf("crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func (*Assertions) [PanicsWithValue](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1462) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.PanicsWithValue)added in v1.2.0

```
func (a *Assertions) PanicsWithValue(expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{})
```

PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value equals the expected panic value.

```
a.PanicsWithValue("crazy error", func(){ GoCrazy() })
```

#### func (*Assertions) [PanicsWithValuef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1473) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.PanicsWithValuef)added in v1.2.0

```
func (a *Assertions) PanicsWithValuef(expected interface{}, f assert.PanicTestFunc, msg string, args ...interface{})
```

PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that the recovered panic value equals the expected panic value.

```
a.PanicsWithValuef("crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func (*Assertions) [Panicsf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1483) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Panicsf)added in v1.2.0

```
func (a *Assertions) Panicsf(f assert.PanicTestFunc, msg string, args ...interface{})
```

Panicsf asserts that the code inside the specified PanicTestFunc panics.

```
a.Panicsf(func(){ GoCrazy() }, "error message %s", "formatted")
```

#### func (*Assertions) [Positive](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1494) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Positive)added in v1.7.0

```
func (a *Assertions) Positive(e interface{}, msgAndArgs ...interface{})
```

Positive asserts that the specified element is positive

```
a.Positive(1)
a.Positive(1.23)
```

#### func (*Assertions) [Positivef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1505) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Positivef)added in v1.7.0

```
func (a *Assertions) Positivef(e interface{}, msg string, args ...interface{})
```

Positivef asserts that the specified element is positive

```
a.Positivef(1, "error message %s", "formatted")
a.Positivef(1.23, "error message %s", "formatted")
```

#### func (*Assertions) [Regexp](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1516) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Regexp)

```
func (a *Assertions) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{})
```

Regexp asserts that a specified regexp matches a string.

```
a.Regexp(regexp.MustCompile("start"), "it's starting")
a.Regexp("start...$", "it's not starting")
```

#### func (*Assertions) [Regexpf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1527) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Regexpf)added in v1.2.0

```
func (a *Assertions) Regexpf(rx interface{}, str interface{}, msg string, args ...interface{})
```

Regexpf asserts that a specified regexp matches a string.

```
a.Regexpf(regexp.MustCompile("start"), "it's starting", "error message %s", "formatted")
a.Regexpf("start...$", "it's not starting", "error message %s", "formatted")
```

#### func (*Assertions) [Same](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1540) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Same)added in v1.4.0

```
func (a *Assertions) Same(expected interface{}, actual interface{}, msgAndArgs ...interface{})
```

Same asserts that two pointers reference the same object.

```
a.Same(ptr1, ptr2)
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func (*Assertions) [Samef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1553) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Samef)added in v1.4.0

```
func (a *Assertions) Samef(expected interface{}, actual interface{}, msg string, args ...interface{})
```

Samef asserts that two pointers reference the same object.

```
a.Samef(ptr1, ptr2, "error message %s", "formatted")
```

Both arguments must be pointer variables. Pointer variable sameness is determined based on the equality of both type and value.

#### func (*Assertions) [Subset](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1565) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Subset)added in v1.2.0

```
func (a *Assertions) Subset(list interface{}, subset interface{}, msgAndArgs ...interface{})
```

Subset asserts that the specified list(array, slice...) or map contains all elements given in the specified subset list(array, slice...) or map.

```
a.Subset([1, 2, 3], [1, 2])
a.Subset({"x": 1, "y": 2}, {"x": 1})
```

#### func (*Assertions) [Subsetf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1577) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Subsetf)added in v1.2.0

```
func (a *Assertions) Subsetf(list interface{}, subset interface{}, msg string, args ...interface{})
```

Subsetf asserts that the specified list(array, slice...) or map contains all elements given in the specified subset list(array, slice...) or map.

```
a.Subsetf([1, 2, 3], [1, 2], "error message %s", "formatted")
a.Subsetf({"x": 1, "y": 2}, {"x": 1}, "error message %s", "formatted")
```

#### func (*Assertions) [True](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1587) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.True)

```
func (a *Assertions) True(value bool, msgAndArgs ...interface{})
```

True asserts that the specified value is true.

```
a.True(myBool)
```

#### func (*Assertions) [Truef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1597) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Truef)added in v1.2.0

```
func (a *Assertions) Truef(value bool, msg string, args ...interface{})
```

Truef asserts that the specified value is true.

```
a.Truef(myBool, "error message %s", "formatted")
```

#### func (*Assertions) [WithinDuration](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1607) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.WithinDuration)

```
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...interface{})
```

WithinDuration asserts that the two times are within duration delta of each other.

```
a.WithinDuration(time.Now(), time.Now(), 10*time.Second)
```

#### func (*Assertions) [WithinDurationf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1617) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.WithinDurationf)added in v1.2.0

```
func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{})
```

WithinDurationf asserts that the two times are within duration delta of each other.

```
a.WithinDurationf(time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
```

#### func (*Assertions) [WithinRange](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1627) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.WithinRange)added in v1.8.0

```
func (a *Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...interface{})
```

WithinRange asserts that a time is within a time range (inclusive).

```
a.WithinRange(time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
```

#### func (*Assertions) [WithinRangef](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1637) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.WithinRangef)added in v1.8.0

```
func (a *Assertions) WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ...interface{})
```

WithinRangef asserts that a time is within a time range (inclusive).

```
a.WithinRangef(time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second), "error message %s", "formatted")
```

#### func (*Assertions) [YAMLEq](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1645) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.YAMLEq)added in v1.4.0

```
func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...interface{})
```

YAMLEq asserts that two YAML strings are equivalent.

#### func (*Assertions) [YAMLEqf](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1653) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.YAMLEqf)added in v1.4.0

```
func (a *Assertions) YAMLEqf(expected string, actual string, msg string, args ...interface{})
```

YAMLEqf asserts that two YAML strings are equivalent.

#### func (*Assertions) [Zero](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1661) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Zero)

```
func (a *Assertions) Zero(i interface{}, msgAndArgs ...interface{})
```

Zero asserts that i is the zero value for its type.

#### func (*Assertions) [Zerof](https://github.com/stretchr/testify/blob/v1.10.0/require/require_forward.go#L1669) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#Assertions.Zerof)added in v1.2.0

```
func (a *Assertions) Zerof(i interface{}, msg string, args ...interface{})
```

Zerof asserts that i is the zero value for its type.

#### type [BoolAssertionFunc](https://github.com/stretchr/testify/blob/v1.10.0/require/requirements.go#L23) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#BoolAssertionFunc)added in v1.2.2

```
type BoolAssertionFunc func(TestingT, bool, ...interface{})
```

BoolAssertionFunc is a common function prototype when validating a bool value. Can be useful for table driven tests.

<details tabindex="-1" id="example-BoolAssertionFunc" class="Documentation-exampleDetails js-exampleContainer" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 1rem 0px 0px; padding: 0px; vertical-align: baseline; display: block;"><summary class="Documentation-exampleDetailsHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px 0px 2rem; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); cursor: pointer; outline: none; text-decoration: none;">Example<span>&nbsp;</span><a href="https://pkg.go.dev/github.com/stretchr/testify/require#example-BoolAssertionFunc" title="Go to Example" aria-label="Go to Example" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); text-decoration: none; opacity: 0;">¶</a></summary><div class="Documentation-exampleDetailsBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline;"><textarea class="Documentation-exampleCode code" spellcheck="false" data-hint_scrollable="true" style="box-sizing: border-box; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; font-size: 0.875rem; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; background-color: var(--color-background-accented); border: var(--border); border-top-left-radius: ; border-top-right-radius: ; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px; color: var(--color-text); overflow-x: auto; padding: 0.625rem; tab-size: 4; white-space: pre; height: 29.625rem; outline: none; resize: none; width: 1263.31px; margin: 0px;"></textarea><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: -0.25rem 0px 1rem; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: 0px 0px 0.3rem 0.3rem; color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre;"><span class="Documentation-exampleOutputLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span><span class="Documentation-exampleOutput" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px 0px 0.5rem; padding: 0px; vertical-align: baseline; border-top-left-radius: 0px; border-top-right-radius: 0px;"></span></pre></div></details>

#### type [ComparisonAssertionFunc](https://github.com/stretchr/testify/blob/v1.10.0/require/requirements.go#L15) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ComparisonAssertionFunc)added in v1.2.2

```
type ComparisonAssertionFunc func(TestingT, interface{}, interface{}, ...interface{})
```

ComparisonAssertionFunc is a common function prototype when comparing two values. Can be useful for table driven tests.

<details tabindex="-1" id="example-ComparisonAssertionFunc" class="Documentation-exampleDetails js-exampleContainer" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 1rem 0px 0px; padding: 0px; vertical-align: baseline; display: block;"><summary class="Documentation-exampleDetailsHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px 0px 2rem; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); cursor: pointer; outline: none; text-decoration: none;">Example<span>&nbsp;</span><a href="https://pkg.go.dev/github.com/stretchr/testify/require#example-ComparisonAssertionFunc" title="Go to Example" aria-label="Go to Example" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); text-decoration: none; opacity: 0;">¶</a></summary><div class="Documentation-exampleDetailsBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline;"><textarea class="Documentation-exampleCode code" spellcheck="false" data-hint_scrollable="true" style="box-sizing: border-box; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; font-size: 0.875rem; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; background-color: var(--color-background-accented); border: var(--border); border-top-left-radius: ; border-top-right-radius: ; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px; color: var(--color-text); overflow-x: auto; padding: 0.625rem; tab-size: 4; white-space: pre; height: 35.875rem; outline: none; resize: none; width: 1263.31px; margin: 0px;"></textarea><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: -0.25rem 0px 1rem; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: 0px 0px 0.3rem 0.3rem; color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre;"><span class="Documentation-exampleOutputLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span><span class="Documentation-exampleOutput" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px 0px 0.5rem; padding: 0px; vertical-align: baseline; border-top-left-radius: 0px; border-top-right-radius: 0px;"></span></pre></div></details>

#### type [ErrorAssertionFunc](https://github.com/stretchr/testify/blob/v1.10.0/require/requirements.go#L27) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ErrorAssertionFunc)added in v1.2.2

```
type ErrorAssertionFunc func(TestingT, error, ...interface{})
```

ErrorAssertionFunc is a common function prototype when validating an error value. Can be useful for table driven tests.

<details tabindex="-1" id="example-ErrorAssertionFunc" class="Documentation-exampleDetails js-exampleContainer" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 1rem 0px 0px; padding: 0px; vertical-align: baseline; display: block;"><summary class="Documentation-exampleDetailsHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px 0px 2rem; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); cursor: pointer; outline: none; text-decoration: none;">Example<span>&nbsp;</span><a href="https://pkg.go.dev/github.com/stretchr/testify/require#example-ErrorAssertionFunc" title="Go to Example" aria-label="Go to Example" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); text-decoration: none; opacity: 0;">¶</a></summary><div class="Documentation-exampleDetailsBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline;"><textarea class="Documentation-exampleCode code" spellcheck="false" data-hint_scrollable="true" style="box-sizing: border-box; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; font-size: 0.875rem; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; background-color: var(--color-background-accented); border: var(--border); border-top-left-radius: ; border-top-right-radius: ; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px; color: var(--color-text); overflow-x: auto; padding: 0.625rem; tab-size: 4; white-space: pre; height: 30.875rem; outline: none; resize: none; width: 1263.31px; margin: 0px;"></textarea><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: -0.25rem 0px 1rem; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: 0px 0px 0.3rem 0.3rem; color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre;"><span class="Documentation-exampleOutputLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span><span class="Documentation-exampleOutput" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px 0px 0.5rem; padding: 0px; vertical-align: baseline; border-top-left-radius: 0px; border-top-right-radius: 0px;"></span></pre></div></details>

#### type [TestingT](https://github.com/stretchr/testify/blob/v1.10.0/require/requirements.go#L4) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#TestingT)

```
type TestingT interface {
	Errorf(format string, args ...interface{})
	FailNow()
}
```

TestingT is an interface wrapper around *testing.T

#### type [ValueAssertionFunc](https://github.com/stretchr/testify/blob/v1.10.0/require/requirements.go#L19) [¶](https://pkg.go.dev/github.com/stretchr/testify/require#ValueAssertionFunc)added in v1.2.2

```
type ValueAssertionFunc func(TestingT, interface{}, ...interface{})
```

ValueAssertionFunc is a common function prototype when validating a single value. Can be useful for table driven tests.



