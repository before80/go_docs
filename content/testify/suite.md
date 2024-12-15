+++
title = "suite"
date = 2024-12-15T11:08:05+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/stretchr/testify/suite](https://pkg.go.dev/github.com/stretchr/testify/suite)
>
> 收录该文档时间： `2024-12-15T11:08:05+08:00`
>
> 版本：[Version: v1.10.0](https://pkg.go.dev/github.com/stretchr/testify/suite?tab=versions)

### Overview [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#pkg-overview)

Package suite contains logic for creating testing suite structs and running the methods on those structs as tests. The most useful piece of this package is that you can create setup/teardown methods on your testing suites, which will run before/after the whole suite or individual tests (depending on which interface(s) you implement).

The suite package does not support parallel tests. See [issue 934](https://github.com/stretchr/testify/issues/934).

A testing suite is usually built by first extending the built-in suite functionality from suite.Suite in testify. Alternatively, you could reproduce that logic on your own if you wanted (you just need to implement the TestingSuite interface from suite/interfaces.go).

After that, you can implement any of the interfaces in suite/interfaces.go to add setup/teardown functionality to your suite, and add any methods that start with "Test" to add tests. Methods that do not match any suite interfaces and do not begin with "Test" will not be run by testify, and can safely be used as helper methods.

Once you've built your testing suite, you need to run the suite (using suite.Run from testify) inside any function that matches the identity that "go test" is already looking for (i.e. func(*testing.T)).

Regular expression to select test suites specified command-line argument "-run". Regular expression to select the methods of test suites specified command-line argument "-m". Suite object has assertion methods.

A crude example:

```
// Basic imports
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
    suite.Suite
    VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *ExampleTestSuite) SetupTest() {
    suite.VariableThatShouldStartAtFive = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *ExampleTestSuite) TestExample() {
    assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
    suite.Equal(5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
    suite.Run(t, new(ExampleTestSuite))
}
```

### Index [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#pkg-index)

- [func Run(t *testing.T, suite TestingSuite)](https://pkg.go.dev/github.com/stretchr/testify/suite#Run)
- [type AfterTest](https://pkg.go.dev/github.com/stretchr/testify/suite#AfterTest)
- [type BeforeTest](https://pkg.go.dev/github.com/stretchr/testify/suite#BeforeTest)
- [type SetupAllSuite](https://pkg.go.dev/github.com/stretchr/testify/suite#SetupAllSuite)
- [type SetupSubTest](https://pkg.go.dev/github.com/stretchr/testify/suite#SetupSubTest)
- [type SetupTestSuite](https://pkg.go.dev/github.com/stretchr/testify/suite#SetupTestSuite)
- [type Suite](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite)
- - [func (suite *Suite) Assert() *assert.Assertions](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.Assert)
  - [func (suite *Suite) Require() *require.Assertions](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.Require)
  - [func (suite *Suite) Run(name string, subtest func()) bool](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.Run)
  - [func (suite *Suite) SetS(s TestingSuite)](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.SetS)
  - [func (suite *Suite) SetT(t *testing.T)](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.SetT)
  - [func (suite *Suite) T() *testing.T](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.T)
- [type SuiteInformation](https://pkg.go.dev/github.com/stretchr/testify/suite#SuiteInformation)
- - [func (s SuiteInformation) Passed() bool](https://pkg.go.dev/github.com/stretchr/testify/suite#SuiteInformation.Passed)
- [type TearDownAllSuite](https://pkg.go.dev/github.com/stretchr/testify/suite#TearDownAllSuite)
- [type TearDownSubTest](https://pkg.go.dev/github.com/stretchr/testify/suite#TearDownSubTest)
- [type TearDownTestSuite](https://pkg.go.dev/github.com/stretchr/testify/suite#TearDownTestSuite)
- [type TestInformation](https://pkg.go.dev/github.com/stretchr/testify/suite#TestInformation)
- [type TestingSuite](https://pkg.go.dev/github.com/stretchr/testify/suite#TestingSuite)
- [type WithStats](https://pkg.go.dev/github.com/stretchr/testify/suite#WithStats)

### Constants [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#pkg-constants)

This section is empty.

### Variables [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#pkg-variables)

This section is empty.

### Functions [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#pkg-functions)

#### func [Run](https://github.com/stretchr/testify/blob/v1.10.0/suite/suite.go#L121) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#Run)

```
func Run(t *testing.T, suite TestingSuite)
```

Run takes a testing suite and runs all of the tests attached to it.

### Types [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#pkg-types)

#### type [AfterTest](https://github.com/stretchr/testify/blob/v1.10.0/suite/interfaces.go#L45) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#AfterTest)added in v1.2.0

```
type AfterTest interface {
	AfterTest(suiteName, testName string)
}
```

AfterTest has a function to be executed right after the test finishes and receives the suite and test names as input

#### type [BeforeTest](https://github.com/stretchr/testify/blob/v1.10.0/suite/interfaces.go#L39) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#BeforeTest)added in v1.2.0

```
type BeforeTest interface {
	BeforeTest(suiteName, testName string)
}
```

BeforeTest has a function to be executed right before the test starts and receives the suite and test names as input

#### type [SetupAllSuite](https://github.com/stretchr/testify/blob/v1.10.0/suite/interfaces.go#L15) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#SetupAllSuite)

```
type SetupAllSuite interface {
	SetupSuite()
}
```

SetupAllSuite has a SetupSuite method, which will run before the tests in the suite are run.

#### type [SetupSubTest](https://github.com/stretchr/testify/blob/v1.10.0/suite/interfaces.go#L58) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#SetupSubTest)added in v1.8.2

```
type SetupSubTest interface {
	SetupSubTest()
}
```

SetupSubTest has a SetupSubTest method, which will run before each subtest in the suite.

#### type [SetupTestSuite](https://github.com/stretchr/testify/blob/v1.10.0/suite/interfaces.go#L21) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#SetupTestSuite)

```
type SetupTestSuite interface {
	SetupTest()
}
```

SetupTestSuite has a SetupTest method, which will run before each test in the suite.

#### type [Suite](https://github.com/stretchr/testify/blob/v1.10.0/suite/suite.go#L23) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite)

```
type Suite struct {
	*assert.Assertions
	// contains filtered or unexported fields
}
```

Suite is a basic testing suite with methods for storing and retrieving the current *testing.T context.

#### func (*Suite) [Assert](https://github.com/stretchr/testify/blob/v1.10.0/suite/suite.go#L71) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.Assert)

```
func (suite *Suite) Assert() *assert.Assertions
```

Assert returns an assert context for suite. Normally, you can call `suite.NoError(expected, actual)`, but for situations where the embedded methods are overridden (for example, you might want to override assert.Assertions with require.Assertions), this method is provided so you can call `suite.Assert().NoError()`.

#### func (*Suite) [Require](https://github.com/stretchr/testify/blob/v1.10.0/suite/suite.go#L57) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.Require)

```
func (suite *Suite) Require() *require.Assertions
```

Require returns a require context for suite.

#### func (*Suite) [Run](https://github.com/stretchr/testify/blob/v1.10.0/suite/suite.go#L98) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.Run)added in v1.3.0

```
func (suite *Suite) Run(name string, subtest func()) bool
```

Run provides suite functionality around golang subtests. It should be called in place of t.Run(name, func(t *testing.T)) in test suite code. The passed-in func will be executed as a subtest with a fresh instance of t. Provides compatibility with go test pkg -run TestSuite/TestName/SubTestName.

#### func (*Suite) [SetS](https://github.com/stretchr/testify/blob/v1.10.0/suite/suite.go#L52) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.SetS)added in v1.8.2

```
func (suite *Suite) SetS(s TestingSuite)
```

SetS needs to set the current test suite as parent to get access to the parent methods

#### func (*Suite) [SetT](https://github.com/stretchr/testify/blob/v1.10.0/suite/suite.go#L42) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.SetT)

```
func (suite *Suite) SetT(t *testing.T)
```

SetT sets the current *testing.T context.

#### func (*Suite) [T](https://github.com/stretchr/testify/blob/v1.10.0/suite/suite.go#L35) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#Suite.T)

```
func (suite *Suite) T() *testing.T
```

T retrieves the current *testing.T context.

#### type [SuiteInformation](https://github.com/stretchr/testify/blob/v1.10.0/suite/stats.go#L6) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#SuiteInformation)added in v1.6.0

```
type SuiteInformation struct {
	Start, End time.Time
	TestStats  map[string]*TestInformation
}
```

SuiteInformation stats stores stats for the whole suite execution.

#### func (SuiteInformation) [Passed](https://github.com/stretchr/testify/blob/v1.10.0/suite/stats.go#L38) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#SuiteInformation.Passed)added in v1.6.0

```
func (s SuiteInformation) Passed() bool
```

#### type [TearDownAllSuite](https://github.com/stretchr/testify/blob/v1.10.0/suite/interfaces.go#L27) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#TearDownAllSuite)

```
type TearDownAllSuite interface {
	TearDownSuite()
}
```

TearDownAllSuite has a TearDownSuite method, which will run after all the tests in the suite have been run.

#### type [TearDownSubTest](https://github.com/stretchr/testify/blob/v1.10.0/suite/interfaces.go#L64) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#TearDownSubTest)added in v1.8.2

```
type TearDownSubTest interface {
	TearDownSubTest()
}
```

TearDownSubTest has a TearDownSubTest method, which will run after each subtest in the suite have been run.

#### type [TearDownTestSuite](https://github.com/stretchr/testify/blob/v1.10.0/suite/interfaces.go#L33) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#TearDownTestSuite)

```
type TearDownTestSuite interface {
	TearDownTest()
}
```

TearDownTestSuite has a TearDownTest method, which will run after each test in the suite.

#### type [TestInformation](https://github.com/stretchr/testify/blob/v1.10.0/suite/stats.go#L12) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#TestInformation)added in v1.6.0

```
type TestInformation struct {
	TestName   string
	Start, End time.Time
	Passed     bool
}
```

TestInformation stores information about the execution of each test.

#### type [TestingSuite](https://github.com/stretchr/testify/blob/v1.10.0/suite/interfaces.go#L7) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#TestingSuite)

```
type TestingSuite interface {
	T() *testing.T
	SetT(*testing.T)
	SetS(suite TestingSuite)
}
```

TestingSuite can store and return the current *testing.T context generated by 'go test'.

#### type [WithStats](https://github.com/stretchr/testify/blob/v1.10.0/suite/interfaces.go#L52) [¶](https://pkg.go.dev/github.com/stretchr/testify/suite#WithStats)added in v1.6.0

```
type WithStats interface {
	HandleStats(suiteName string, stats *SuiteInformation)
}
```

WithStats implements HandleStats, a function that will be executed when a test suite is finished. The stats contain information about the execution of that suite and its tests.
