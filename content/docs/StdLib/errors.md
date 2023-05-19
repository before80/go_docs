+++
title = "error"
linkTitle = "error"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# errors

https://pkg.go.dev/errors@go1.20.1

​	 errors 包实现了操作错误的函数。

​	New函数创建的错误只包含文本消息。

​	如果e的类型具有下列方法之一：

```go 
Unwrap() error
Unwrap() []error
```

​	那么e就包装了另一个错误。如果e.Unwrap()返回非nil的错误w或包含w的切片，则我们说e包装了w。Unwrap方法返回的nil错误表示e不包装任何错误。Unwrap方法返回一个包含nil错误值的[]error是无效的。 

### 使用Unwrap方法的示例

```go 
package main

import (
	"fmt"
)

type myError1 struct {
	Name string
}

func (e myError1) Error() string {
	return fmt.Sprintf("name=%v", e.Name)
}

type myError2 struct {
	Desc string
	Err  error
}

func (e myError2) Error() string {
	return fmt.Sprintf("Desc:%v,Err:%v", e.Desc, e.Err)
}

func (e myError2) Unwrap() error {
	return e.Err
}

func main() {
	err := myError2{"zlx1", myError1{"zlx"}}
	fmt.Println(err)
	fmt.Println(err.Unwrap())

	err2 := fmt.Errorf("err2=%w", err)
	fmt.Println(err2)
}
Output:

Desc:zlx1,Err:name=zlx
name=zlx
err2=Desc:zlx1,Err:name=zlx
```

​	创建包装错误的简单方法是调用fmt.Errorf并对错误参数应用[%w](../fmt#func-errorf)占位符：

```
wrapsErr := fmt.Errorf("... %w ...", ..., err, ...)
```

​	对错误进行连续的解包将创建一棵树。Is和As函数通过先检查错误本身，然后依次检查其每个子树(先序，深度优先遍历)来检查错误的树。

​	Is函数检查其第一个参数的树，查找与第二个参数匹配的错误。它报告它是否找到匹配项。应优先使用它，而不是简单的相等性检查：

```go 
if errors.Is(err, fs.ErrExist)
```

优于

```go 
if err == fs.ErrExist
```

因为前者将成功地匹配err包装了fs.ErrExist。

​	As检查其第一个参数的树，查找可以分配给其第二个参数的错误，`该参数必须为指针`。如果成功，则执行分配并返回true。否则，它将返回false。下面这种形式是优先使用的：

``` go 
var perr *fs.PathError
if errors.As(err, &perr) {
	fmt.Println(perr.Path)
}
```

优于

```go 
if perr, ok := err.(*fs.PathError); ok {
	fmt.Println(perr.Path)
}
```

​	因为前者将成功地匹配err包装了`*fs.PathError`。

### Example
``` go 
package main

import (
	"fmt"
	"time"
)

// MyError 是一个实现了包含时间和消息的错误类型。
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.When, e.What)
}

func oops() error {
	return MyError{
		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
		"the file system has gone away",
	}
}

func main() {
	if err := oops(); err != nil {
		fmt.Println(err)
	}
}
Output:

1989-03-15 22:30:00 +0000 UTC: the file system has gone away
```

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func As  <- go1.13

``` go 
func As(err error, target any) bool
```

​	As 函数在 err 的错误树中查找第一个与 target 匹配的错误，如果找到，则将 target 设置为该错误的值并返回 true。否则，返回 false。

​	错误树包括 err 本身，以及通过反复调用 Unwrap 获得的错误。当 err 包含多个错误时，As 按深度优先遍历顺序依次检查 err 的每个子错误。

​	如果错误的具体值可以分配给 target 指向的值，则错误与 target 匹配，或者如果错误有一个 As(target) bool 方法，则该方法返回 true。在后一种情况下，As 方法负责设置 target。

​	错误类型可能会提供 As 方法，以便可以将其视为不同的错误类型。

​	如果 target 不是一个实现了 error 接口的类型或任何接口类型的非 nil 指针，则 As 函数会引发 panic。

##### As Example
``` go 
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

}
Output:

Failed at path: non-existing
```

#### func Is  <- go1.13

``` go 
func Is(err, target error) bool
```

​	Is 函数报告 err 的错误树中是否有任何错误与 target 匹配。

​	错误树包括 err 本身，以及通过反复调用 Unwrap 获得的错误。当 err 包含多个错误时，Is 按深度优先遍历顺序依次检查 err 的每个子错误。

​	如果一个错误等于 target 或者实现了一个 Is(error) bool 方法，使得 Is(target) 返回 true，则认为该错误与 target 匹配。

​	错误类型可能会提供 Is 方法，以便可以将其视为与现有错误等效。例如，如果 MyError 定义了

``` go 
func (m MyError) Is(target error) bool { return target == fs.ErrExist }
```

​	那么 Is(MyError{}, fs.ErrExist) 返回 true。标准库中的 [syscall.Errno.Is]() 就是一个示例。Is 方法只应该浅层比较 err 和 target，不要对它们中的任何一个调用 Unwrap。

##### Is Example
``` go 
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}

}
Output:

file does not exist
```

#### func Join  <- go1.20

``` go 
func Join(errs ...error) error
```

​	Join 函数返回一个包含给定错误的错误。任何 nil 错误值都会被丢弃。如果 errs 不包含任何非 nil 值，则返回 nil。该错误的格式为 errs 中每个元素的 Error 方法返回值的串联，每个字符串之间都有一个换行符。

##### Join Example
``` go 
package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err := errors.Join(err1, err2)
	fmt.Println(err)
	if errors.Is(err, err1) {
		fmt.Println("err is err1")
	}
	if errors.Is(err, err2) {
		fmt.Println("err is err2")
	}
}
Output:

err1
err2
err is err1
err is err2
```

#### func New 

``` go 
func New(text string) error
```

​	New 函数返回一个错误，该错误格式化为给定的文本。即使文本相同，每次调用 New 都会返回一个不同的错误值。

##### New Example
``` go 
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}
Output:

emit macho dwarf: elf header corrupted
```

##### New Example (Errorf)

​	fmt 包的 Errorf 函数允许我们使用该包的格式化功能来创建描述性错误消息。

``` go 
package main

import (
	"fmt"
)

func main() {
	const name, id = "bimmler", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	if err != nil {
		fmt.Print(err)
	}
}
Output:

user "bimmler" (id 17) not found
```

#### func Unwrap  <- go1.13

``` go 
func Unwrap(err error) error
```

​	Unwrap 函数返回 err 上的 Unwrap 方法的结果，如果 err 的类型包含返回错误的 Unwrap 方法。否则，Unwrap 函数返回 nil。

​	如果 Unwrap 方法返回 []error，则 Unwrap 函数返回 nil。

##### Unwrap Example
``` go 
package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: [%w]", err1)
	fmt.Println(err2)
	fmt.Println(errors.Unwrap(err2))
	// Output
	// error2: [error1]
	// error1
}
Output:

error2: [error1]
error1
```



## 类型

This section is empty.