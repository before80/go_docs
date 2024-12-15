+++
title = "composition"
date = 2024-12-15T11:18:27+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://github.com/smartystreets/goconvey/wiki/Composition](https://github.com/smartystreets/goconvey/wiki/Composition)
>
> 收录该文档时间： `2024-12-15T11:18:27+08:00`

# Composition

[Edit](https://github.com/smartystreets/goconvey/wiki/Composition/_edit) [New page](https://github.com/smartystreets/goconvey/wiki/_new)

Warren Turkal edited this page on Sep 15, 2015 · [7 revisions](https://github.com/smartystreets/goconvey/wiki/Composition/_history)

Writing self-documenting tests is remarkably easy with GoConvey.

### Examples



First, take a look through the [examples folder](https://github.com/smartystreets/goconvey/tree/master/examples) to get the basic idea. We'd recommend reviewing [isolated_execution_test.go](https://github.com/smartystreets/goconvey/blob/master/convey/isolated_execution_test.go) for a more thorough understanding of how you can compose test cases.

### Functions



See [GoDoc](http://godoc.org/github.com/smartystreets/goconvey) for exported functions and assertions. You'd be most interested in the [convey](http://godoc.org/github.com/smartystreets/goconvey/convey) package.

### Quick tutorial



In your test file, import needed packages:

```
import(
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)
```



(Notice the dot-notation for the `convey` package, for convenience.)

Since GoConvey uses `go test`, set up a Test function:

```
func TestSomething(t *testing.T) {
	
}
```



To set up test cases, we use `Convey()` to define scope/context/behavior/ideas, and `So()` to make assertions. For example:

```
Convey("1 should equal 1", t, func() {
	So(1, ShouldEqual, 1)
})
```



There's a working GoConvey test. Notice that we pass in the `*testing.T` object. Only the top-level calls to `Convey()` require that. For nested calls, you must omit it. For instance:

```
Convey("Comparing two variables", t, func() {
	myVar := "Hello, world!"

	Convey(`"Asdf" should NOT equal "qwerty"`, func() {
		So("Asdf", ShouldNotEqual, "qwerty")
	})

	Convey("myVar should not be nil", func() {
		So(myVar, ShouldNotBeNil)
	})
})
```



If you haven't yet implemented a test or scope, just set its function to `nil` to [skip](https://github.com/smartystreets/goconvey/wiki/Skip) it:

```
Convey("This isn't yet implemented", nil)
```



### [Next](https://github.com/smartystreets/goconvey/wiki/Assertions)



Next, you should learn about the [standard assertions](https://github.com/smartystreets/goconvey/wiki/Assertions). You may also skip ahead to [executing tests](https://github.com/smartystreets/goconvey/wiki/Execution) or to [Skip](https://github.com/smartystreets/goconvey/wiki/Skip) to make testing more convenient.
