+++
title = "mock"
date = 2024-12-15T11:07:54+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/stretchr/testify/mock](https://pkg.go.dev/github.com/stretchr/testify/mock)
>
> 收录该文档时间： `2024-12-15T11:07:54+08:00`
>
> 版本：[Version: v1.10.0](https://pkg.go.dev/github.com/stretchr/testify/mock?tab=versions)

### Overview [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#pkg-overview)

- [Example Usage](https://pkg.go.dev/github.com/stretchr/testify/mock#hdr-Example_Usage)

Package mock provides a system by which it is possible to mock your objects and verify calls are happening as expected.

#### Example Usage [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#hdr-Example_Usage)

The mock package provides an object, Mock, that tracks activity on another object. It is usually embedded into a test object as shown below:

```
type MyTestObject struct {
  // add a Mock object instance
  mock.Mock

  // other fields go here as normal
}
```

When implementing the methods of an interface, you wire your functions up to call the Mock.Called(args...) method, and return the appropriate values.

For example, to mock a method that saves the name and age of a person and returns the year of their birth or an error, you might write this:

```
func (o *MyTestObject) SavePersonDetails(firstname, lastname string, age int) (int, error) {
  args := o.Called(firstname, lastname, age)
  return args.Int(0), args.Error(1)
}
```

The Int, Error and Bool methods are examples of strongly typed getters that take the argument index position. Given this argument list:

```
(12, true, "Something")
```

You could read them out strongly typed like this:

```
args.Int(0)
args.Bool(1)
args.String(2)
```

For objects of your own type, use the generic Arguments.Get(index) method and make a type assertion:

```
return args.Get(0).(*MyObject), args.Get(1).(*AnotherObjectOfMine)
```

This may cause a panic if the object you are getting is nil (the type assertion will fail), in those cases you should check for nil first.



### Constants [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#pkg-constants)

[View Source](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L776)

```
const (
	// Anything is used in Diff and Assert when the argument being tested
	// shouldn't be taken into consideration.
	Anything = "mock.Anything"
)
```

### Variables [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#pkg-variables)

This section is empty.

### Functions [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#pkg-functions)

#### func [AssertExpectationsForObjects](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L592) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#AssertExpectationsForObjects)

```
func AssertExpectationsForObjects(t TestingT, testObjects ...interface{}) bool
```

AssertExpectationsForObjects asserts that everything specified with On and Return of the specified objects was in fact called as expected.

Calls may have occurred in any order.

#### func [InOrder](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L284) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#InOrder)added in v1.10.0

```
func InOrder(calls ...*Call)
```

InOrder defines the order in which the calls should be made

```
For example:

InOrder(
	Mock.On("init").Return(nil),
	Mock.On("Do").Return(nil),
)
```

#### func [MatchedBy](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L907) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#MatchedBy)

```
func MatchedBy(fn interface{}) argumentMatcher
```

MatchedBy can be used to match a mock call based on only certain properties from a complex struct or some calculation. It takes a function that will be evaluated with the called argument and will return true when there's a match and false otherwise.

Example:

```
m.On("Do", MatchedBy(func(req *http.Request) bool { return req.Host == "example.com" }))
```

fn must be a function accepting a single argument (of the expected type) which returns a bool. If fn doesn't match the required signature, MatchedBy() panics.

### Types [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#pkg-types)

<details class="Documentation-deprecatedDetails js-deprecatedDetails" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; display: block; color: var(--color-text-subtle);"><summary style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; list-style: none; opacity: 1;"><h4 tabindex="-1" id="AnythingOfTypeArgument" data-kind="type" class="Documentation-typeHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 600; font-stretch: inherit; line-height: 1.25em; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1.125rem; margin: 1.5rem 0px 0.5rem; padding: 0px; vertical-align: baseline; word-break: break-word; align-items: baseline; display: flex; justify-content: space-between;"><span class="Documentation-deprecatedTitle" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; align-items: center; display: flex; gap: 0.5rem;">type<a class="Documentation-source" href="https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L794" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 1;">AnythingOfTypeArgument</a><span class="Documentation-deprecatedTag" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: 1.375; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.75rem; margin: 0px; padding: 0.125rem 0.25rem; vertical-align: middle; background-color: var(--color-border); border-radius: 0.125rem; color: var(--color-text-inverted); text-transform: uppercase;">deprecated</span><span class="Documentation-deprecatedBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.87rem; margin: 0px 0.5rem 0px 0.25rem; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></span><span class="Documentation-sinceVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.9375rem; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></h4></summary><div class="go-Message go-Message--warning Documentation-deprecatedItemBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 1rem 1rem 0.5rem; vertical-align: baseline; color: var(--gray-1); width: 1263.31px; background-color: var(--color-background-warning);"><div class="Documentation-declaration" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre; scroll-padding-top: calc(var(--js-sticky-header-height, 3.5rem) + .75rem);"></pre></div><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"><a href="https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Diff" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Assert" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"><a href="https://pkg.go.dev/github.com/stretchr/testify/mock#AnythingOfType" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></p><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre;"></pre><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"><a href="https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.On" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></p><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre;"></pre><div class="Documentation-typeFunc" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><h4 tabindex="-1" id="AnythingOfType" data-kind="function" class="Documentation-typeFuncHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 600; font-stretch: inherit; line-height: 1.25em; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1.125rem; margin: 1.5rem 0px 0.5rem; padding: 0px; vertical-align: baseline; word-break: break-word; align-items: baseline; display: flex; justify-content: space-between;"><span style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline;"><a class="Documentation-source" href="https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L808" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 1;"></a><a class="Documentation-idLink" href="https://pkg.go.dev/github.com/stretchr/testify/mock#AnythingOfType" title="Go to AnythingOfType" aria-label="Go to AnythingOfType" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 0;"></a></span><span class="Documentation-sinceVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.9375rem; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></h4><div class="Documentation-declaration" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre-wrap; scroll-padding-top: calc(var(--js-sticky-header-height, 3.5rem) + .75rem); word-break: break-all; overflow-wrap: break-word;"><a href="https://pkg.go.dev/builtin#string" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/github.com/stretchr/testify/mock#AnythingOfTypeArgument" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></pre></div><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"><a href="https://pkg.go.dev/reflect#Type.String" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-size-adjust: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre-wrap; word-break: break-all; overflow-wrap: break-word;"></pre></div></div></details>

#### type [Arguments](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L774) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments)

```
type Arguments []interface{}
```

Arguments holds an array of method arguments or return values.

#### func (Arguments) [Assert](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L1051) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Assert)

```
func (args Arguments) Assert(t TestingT, objects ...interface{}) bool
```

Assert compares the arguments with the specified objects and fails if they do not exactly match.

#### func (Arguments) [Bool](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L1125) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Bool)

```
func (args Arguments) Bool(index int) bool
```

Bool gets the argument at the specified index. Panics if there is no argument, or if the argument is of the wrong type.

#### func (Arguments) [Diff](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L945) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Diff)

```
func (args Arguments) Diff(objects []interface{}) (string, int)
```

Diff gets a string describing the differences between the arguments and the specified objects.

Returns the diff string and number of differences found.

#### func (Arguments) [Error](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L1110) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Error)

```
func (args Arguments) Error(index int) error
```

Error gets the argument at the specified index. Panics if there is no argument, or if the argument is of the wrong type.

#### func (Arguments) [Get](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L924) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Get)

```
func (args Arguments) Get(index int) interface{}
```

Get Returns the argument at the specified index.

#### func (Arguments) [Int](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L1099) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Int)

```
func (args Arguments) Int(index int) int
```

Int gets the argument at the specified index. Panics if there is no argument, or if the argument is of the wrong type.

#### func (Arguments) [Is](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L932) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Is)

```
func (args Arguments) Is(objects ...interface{}) bool
```

Is gets whether the objects match the arguments specified.

#### func (Arguments) [String](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L1075) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.String)

```
func (args Arguments) String(indexOrNil ...int) string
```

String gets the argument at the specified index. Panics if there is no argument, or if the argument is of the wrong type.

If no index is provided, String() returns a complete string representation of the arguments.

#### type [Call](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L37) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call)

```
type Call struct {
	Parent *Mock

	// The name of the method that was or will be called.
	Method string

	// Holds the arguments of the method.
	Arguments Arguments

	// Holds the arguments that should be returned when
	// this method is called.
	ReturnArguments Arguments

	// The number of times to return the return arguments when setting
	// expectations. 0 means to always return the value.
	Repeatability int

	// Holds a channel that will be used to block the Return until it either
	// receives a message or is closed. nil means it returns immediately.
	WaitFor <-chan time.Time

	// Holds a handler used to manipulate arguments content that are passed by
	// reference. It's useful when mocking methods such as unmarshalers or
	// decoders.
	RunFn func(Arguments)

	// PanicMsg holds msg to be used to mock panic on the function call
	//  if the PanicMsg is set to a non nil string the function call will panic
	// irrespective of other settings
	PanicMsg *string
	// contains filtered or unexported fields
}
```

Call represents a method call and is used for setting expectations, as well as recording activity.

#### func (*Call) [After](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L168) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.After)

```
func (c *Call) After(d time.Duration) *Call
```

After sets how long to block until the call returns

```
Mock.On("MyMethod", arg1, arg2).After(time.Second)
```

#### func (*Call) [Maybe](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L192) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.Maybe)added in v1.2.0

```
func (c *Call) Maybe() *Call
```

Maybe allows the method call to be optional. Not calling an optional method will not cause an error while asserting expectations

#### func (*Call) [NotBefore](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L262) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.NotBefore)added in v1.8.0

```
func (c *Call) NotBefore(calls ...*Call) *Call
```

NotBefore indicates that the mock should only be called after the referenced calls have been called as expected. The referenced calls may be from the same mock instance and/or other mock instances.

```
Mock.On("Do").Return(nil).NotBefore(
    Mock.On("Init").Return(nil)
)
```

#### func (*Call) [On](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L207) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.On)

```
func (c *Call) On(methodName string, arguments ...interface{}) *Call
```

On chains a new expectation description onto the mocked interface. This allows syntax like.

```
Mock.
   On("MyMethod", 1).Return(nil).
   On("MyOtherMethod", 'a', 'b', 'c').Return(errors.New("Some Error"))
```

#### func (*Call) [Once](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L132) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.Once)

```
func (c *Call) Once() *Call
```

Once indicates that the mock should only return the value once.

```
Mock.On("MyMethod", arg1, arg2).Return(returnArg1, returnArg2).Once()
```

#### func (*Call) [Panic](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L120) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.Panic)added in v1.6.0

```
func (c *Call) Panic(msg string) *Call
```

Panic specifies if the function call should fail and the panic message

```
Mock.On("DoSomething").Panic("test panic")
```

#### func (*Call) [Return](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L108) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.Return)

```
func (c *Call) Return(returnArguments ...interface{}) *Call
```

Return specifies the return arguments for the expectation.

```
Mock.On("DoSomething").Return(errors.New("failed"))
```

#### func (*Call) [Run](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L183) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.Run)

```
func (c *Call) Run(fn func(args Arguments)) *Call
```

Run sets a handler to be called before returning. It can be used when mocking a method (such as an unmarshaler) that takes a pointer to a struct and sets properties in such struct

```
Mock.On("Unmarshal", AnythingOfType("*map[string]interface{}")).Return().Run(func(args Arguments) {
	arg := args.Get(0).(*map[string]interface{})
	arg["foo"] = "bar"
})
```

#### func (*Call) [Times](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L147) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.Times)

```
func (c *Call) Times(i int) *Call
```

Times indicates that the mock should only return the indicated number of times.

```
Mock.On("MyMethod", arg1, arg2).Return(returnArg1, returnArg2).Times(5)
```

#### func (*Call) [Twice](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L139) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.Twice)

```
func (c *Call) Twice() *Call
```

Twice indicates that the mock should only return the value twice.

```
Mock.On("MyMethod", arg1, arg2).Return(returnArg1, returnArg2).Twice()
```

#### func (*Call) [Unset](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L214) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.Unset)added in v1.7.5

```
func (c *Call) Unset() *Call
```

Unset removes a mock handler from being called.

```
test.On("func", mock.Anything).Unset()
```

#### func (*Call) [WaitUntil](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L158) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Call.WaitUntil)

```
func (c *Call) WaitUntil(w <-chan time.Time) *Call
```

WaitUntil sets the channel that will block the mock's return until its closed or a message is received.

```
Mock.On("MyMethod", arg1, arg2).WaitUntil(time.After(time.Second))
```

#### type [FunctionalOptionsArgument](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L832) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#FunctionalOptionsArgument)added in v1.8.3

```
type FunctionalOptionsArgument struct {
	// contains filtered or unexported fields
}
```

FunctionalOptionsArgument contains a list of functional options arguments expected for use when matching a list of arguments.

#### func [FunctionalOptions](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L852) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#FunctionalOptions)added in v1.8.3

```
func FunctionalOptions(values ...interface{}) *FunctionalOptionsArgument
```

FunctionalOptions returns an [FunctionalOptionsArgument](https://pkg.go.dev/github.com/stretchr/testify/mock#FunctionalOptionsArgument) object containing the expected functional-options to check for.

For example:

```
args.Assert(t, FunctionalOptions(foo.Opt1("strValue"), foo.Opt2(613)))
```

#### func (*FunctionalOptionsArgument) [String](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L837) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#FunctionalOptionsArgument.String)added in v1.8.3

```
func (f *FunctionalOptionsArgument) String() string
```

String returns the string representation of FunctionalOptionsArgument

#### type [IsTypeArgument](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L815) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#IsTypeArgument)added in v1.5.0

```
type IsTypeArgument struct {
	// contains filtered or unexported fields
}
```

IsTypeArgument is a struct that contains the type of an argument for use when type checking. This is an alternative to [AnythingOfType](https://pkg.go.dev/github.com/stretchr/testify/mock#AnythingOfType). Used in [Arguments.Diff](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Diff) and [Arguments.Assert](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Assert).

#### func [IsType](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L826) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#IsType)added in v1.5.0

```
func IsType(t interface{}) *IsTypeArgument
```

IsType returns an IsTypeArgument object containing the type to check for. You can provide a zero-value of the type to check. This is an alternative to [AnythingOfType](https://pkg.go.dev/github.com/stretchr/testify/mock#AnythingOfType). Used in [Arguments.Diff](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Diff) and [Arguments.Assert](https://pkg.go.dev/github.com/stretchr/testify/mock#Arguments.Assert).

For example:

```
args.Assert(t, IsType(""), IsType(0))
```

#### type [Mock](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L293) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock)

```
type Mock struct {
	// Represents the calls that are expected of
	// an object.
	ExpectedCalls []*Call

	// Holds the calls that were made to this mocked object.
	Calls []Call
	// contains filtered or unexported fields
}
```

Mock is the workhorse used to track activity on another object. For an example of its usage, refer to the "Example Usage" section at the top of this document.

#### func (*Mock) [AssertCalled](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L669) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.AssertCalled)

```
func (m *Mock) AssertCalled(t TestingT, methodName string, arguments ...interface{}) bool
```

AssertCalled asserts that the method was called. It can produce a false result when an argument is a pointer type and the underlying value changed after calling the mocked method.

#### func (*Mock) [AssertExpectations](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L612) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.AssertExpectations)

```
func (m *Mock) AssertExpectations(t TestingT) bool
```

AssertExpectations asserts that everything specified with On and Return was in fact called as expected. Calls may have occurred in any order.

#### func (*Mock) [AssertNotCalled](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L692) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.AssertNotCalled)

```
func (m *Mock) AssertNotCalled(t TestingT, methodName string, arguments ...interface{}) bool
```

AssertNotCalled asserts that the method was not called. It can produce a false result when an argument is a pointer type and the underlying value changed after calling the mocked method.

#### func (*Mock) [AssertNumberOfCalls](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L652) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.AssertNumberOfCalls)

```
func (m *Mock) AssertNumberOfCalls(t TestingT, methodName string, expectedCalls int) bool
```

AssertNumberOfCalls asserts that the method was called expectedCalls times.

#### func (*Mock) [Called](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L465) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.Called)

```
func (m *Mock) Called(arguments ...interface{}) Arguments
```

Called tells the mock object that a method has been called, and gets an array of arguments to return. Panics if the call is unexpected (i.e. not preceded by appropriate .On .Return() calls) If Call.WaitFor is set, blocks until the channel is closed or receives a message.

#### func (*Mock) [IsMethodCallable](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L707) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.IsMethodCallable)added in v1.6.0

```
func (m *Mock) IsMethodCallable(t TestingT, methodName string, arguments ...interface{}) bool
```

IsMethodCallable checking that the method can be called If the method was called more than `Repeatability` return false

#### func (*Mock) [MethodCalled](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L488) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.MethodCalled)added in v1.2.0

```
func (m *Mock) MethodCalled(methodName string, arguments ...interface{}) Arguments
```

MethodCalled tells the mock object that the given method has been called, and gets an array of arguments to return. Panics if the call is unexpected (i.e. not preceded by appropriate .On .Return() calls) If Call.WaitFor is set, blocks until the channel is closed or receives a message.

#### func (*Mock) [On](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L359) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.On)

```
func (m *Mock) On(methodName string, arguments ...interface{}) *Call
```

On starts a description of an expectation of the specified method being called.

```
Mock.On("MyMethod", arg1, arg2)
```

#### func (*Mock) [String](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L316) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.String)added in v1.7.1

```
func (m *Mock) String() string
```

String provides a %v format string for Mock. Note: this is used implicitly by Arguments.Diff if a Mock is passed. It exists because go's default %v formatting traverses the struct without acquiring the mutex, which is detected by go test -race.

#### func (*Mock) [Test](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L335) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.Test)added in v1.2.2

```
func (m *Mock) Test(t TestingT)
```

Test sets the test struct variable of the mock object

#### func (*Mock) [TestData](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L322) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#Mock.TestData)

```
func (m *Mock) TestData() objx.Map
```

TestData holds any data that might be useful for testing. Testify ignores this data completely allowing you to do whatever you like with it.

#### type [TestingT](https://github.com/stretchr/testify/blob/v1.10.0/mock/mock.go#L25) [¶](https://pkg.go.dev/github.com/stretchr/testify/mock#TestingT)

```
type TestingT interface {
	Logf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	FailNow()
}
```

TestingT is an interface wrapper around *testing.T
