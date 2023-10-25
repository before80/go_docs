+++
title = "context"
linkTitle = "context"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/context@go1.20.1

​	context包context 定义了 Context 类型，它在 API 边界和进程之间传递截止时间、取消信号和其他请求作用域值。

​	服务端的传入请求应该创建一个 Context，对服务端的外部调用应该接受一个 Context。它们之间的函数调用链必须传播 Context，并且可以使用 WithCancel、WithDeadline、WithTimeout 或 WithValue 创建派生的 Context 来替换它。当一个 Context 被取消时，所有从它派生的 Context 也会被取消。	

​	WithCancel、WithDeadline 和 WithTimeout 函数接受一个 Context(父级)并返回一个派生的 Context(子级)和一个 CancelFunc。调用 CancelFunc 会取消子级及其子级，移除父级对子级的引用并停止任何相关的定时器。如果不调用 CancelFunc，则会泄漏子级及其子级，直到父级被取消或定时器触发。go vet 工具检查所有控制流路径上是否使用了 CancelFuncs。

​	WithCancelCause 函数返回一个 CancelCauseFunc，它接受一个错误并将其记录为取消原因。调用取消的 Context 或任何其子级的 Cause 函数都会检索到取消原因。如果未指定原因，则 Cause(ctx) 返回与 ctx.Err() 相同的值。	

​	使用 Context 的程序应该遵循以下规则，以使接口在包之间保持一致并启用静态分析工具检查上下文传播：	

(1)不要在结构类型中存储 Context；相反，将 Context 显式传递给需要它的每个函数。Context 应该是第一个参数，通常命名为 ctx：

```go 
func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}
```

​	(2)即使函数允许，也不要传递 nil Context。如果不确定要使用哪个 Context，请传递 context.TODO。

​	(3)仅将 context Value 用于跨进程和 API 传递的请求作用域数据，而不是将可选参数传递给函数。

​	(4)同一个 Context 可以传递给在不同 goroutine 中运行的函数；Context 可以同时被多个 goroutine 安全使用。

​	有关使用 Context 的示例代码，请参见[博客《Go并发模式：Context》]({{< ref "/goBlog/2014/GoConcurrencyPatternsContext" >}})。


## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/context/context.go;l=163)

```go 
var Canceled = errors.New("context canceled")
```

​	Canceled是当上下文被取消时，Context.Err返回的错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/context/context.go;l=167)

```go 
var DeadlineExceeded error = deadlineExceededError{}
```

​	DeadlineExceeded是当上下文的截止时间过期时，Context.Err返回的错误。

## 函数

#### func Cause  <- go1.20

```go 
func Cause(c Context) error
```

​	Cause函数返回一个非nil的错误，解释为什么c被取消。c或其父级的第一个取消设置原因。如果取消是通过对CancelCauseFunc(err)的调用进行的，则Cause返回err。否则，Cause(c)返回与c.Err()相同的值。如果c尚未被取消，则Cause返回nil。

#### func WithCancel 

```go 
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

​	WithCancel函数返回parent的副本并创建一个新的Done通道。当调用返回的cancel函数或父级上下文的Done通道关闭时，返回的上下文的Done通道关闭。

​	取消此上下文会释放与之相关联的资源，因此代码应尽快调用cancel，以便在此上下文中运行的操作完成。

##### WithCancel Example

​	这个例子演示了使用可取消的上下文来防止 Goroutine 泄漏。在例子函数的结尾处，由 gen 启动的 Goroutine 将在不泄漏的情况下返回。

``` go 
package main

import (
	"context"
	"fmt"
)

func main() {
    // gen在一个单独的goroutine中生成整数，并将它们发送到返回的通道中。
	// gen的调用者在消耗完生成的整数后需要取消上下文，以免泄露gen启动的内部goroutine。
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // 返回时不泄露goroutine的信息
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们消耗完整数后再取消。

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
Output:

1
2
3
4
5
```

#### func WithCancelCause  <- go1.20

```go 
func WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc)
```

​	WithCancelCause函数类似于WithCancel但返回一个CancelCauseFunc而不是CancelFunc。使用非nil错误(the "cause")调用cancel将记录该错误在ctx中;然后可以使用Cause(ctx)检索它。使用nil调用cancel将原因设置为已取消。

例如使用：

```go 
ctx, cancel := context.WithCancelCause(parent)
cancel(myError)
ctx.Err() // returns context.Canceled
context.Cause(ctx) // returns myError
```

#### func WithDeadline 

```go 
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
```

​	WithDeadline函数返回父Context的副本，其截止时间早于或等于d。如果父Context的截止时间早于d，则WithDeadline(parent, d)在语义上等同于parent。返回的Context的Done通道在到期时关闭，当调用返回的cancel函数时关闭，或者当父Context的Done通道关闭时关闭，以先发生的事件为准。

​	取消此上下文会释放与其关联的资源，因此代码应在此Context中运行的操作完成后尽快调用cancel。

##### WithDeadline Example

此示例传递了一个带有任意deadline 的上下文，以告诉阻塞函数，它应该在到达deadline 后立即放弃它的工作。

``` go 
package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用其取消函数都是很好的做法。
    // 如果不这样做，可能会使上下文和它的parent活得比必要的时间更长。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}

Output:

context deadline exceeded
```

#### func WithTimeout 

```go 
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

​	WithTimeout函数返回WithDeadline(parent, time.Now().Add(timeout))。

​	取消此Context会释放与其关联的资源，因此代码应在此Context中运行的操作完成后尽快调用cancel：

```go 
func slowOperationWithTimeout(ctx context.Context) (Result, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()  // 如果slowOperation在超时前完成，则释放资源
	return slowOperation(ctx)
}
```

##### WithTimeout Example

此示例传递了一个带有超时的上下文，以告诉阻塞函数在超时过后应该放弃它的工作。

``` go 
package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
    //传递一个带有超时的上下文，以告诉阻塞函数在超时过后应该放弃它的工作
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}

}

Output:

context deadline exceeded
```

## 类型

### type CancelCauseFunc  <- go1.20

```go 
type CancelCauseFunc func(cause error)
```

​	CancelCauseFunc类型的行为类似于CancelFunc，但还会设置取消原因。该原因可以通过在取消的Context或其派生的Context上调用Cause来检索。

​	如果上下文已经被取消，则CancelCauseFunc不会设置取消原因。例如，如果childContext是从parentContext派生的： 

- 如果parentContext在childContext之前以cause1取消，则Cause(parentContext) == Cause(childContext) == cause1  (即 parentContext 可以影响到 childContext )
- 如果childContext在parentContext之前以cause2取消，则Cause(parentContext) == cause1，并且Cause(childContext) == cause2。(即 childContext 影响不到 parentContext )

### type CancelFunc 

```go 
type CancelFunc func()
```

​	CancelFunc类型告诉操作放弃它的工作。CancelFunc类型不等待工作停止。`多个goroutine可以同时调用CancelFunc`。在第一次调用之后，对CancelFunc的后续调用不起作用。

### type Context 

```go 
type Context interface {
    // Deadline返回代表该上下文所做的工作应该被取消的时间。
    // 如果没有设置deadline，Deadline方法返回ok==false。
    // 对Deadline的连续调用会返回相同的结果。
	Deadline() (deadline time.Time, ok bool)

    // Done返回一个通道，该通道在此上下文代表的工作应被取消时关闭。
    // 如果这个上下文永远不能被取消，则Done可能会返回nil。
    // 对Done的连续调用会返回相同的值。
    // Done通道的关闭可以异步发生，在cancel函数返回之后。
	//
	// WithCancel安排在调用cancel时关闭Done；
	// WithDeadline安排在截止时间过期时关闭Done；
	// WithTimeout安排在超时时间过去时关闭Done。
	//
    // Done用于在select语句中使用：
	//
	//	// Stream用DoSomething生成数值，并将其发送到out，
    //  // 直到DoSomething返回错误或ctx.Done被关闭。
	//  func Stream(ctx context.Context, out chan<- Value) error {
	//  	for {
	//  		v, err := DoSomething(ctx)
	//  		if err != nil {
	//  			return err
	//  		}
	//  		select {
	//  		case <-ctx.Done():
	//  			return ctx.Err()
	//  		case out <- v:
	//  		}
	//  	}
	//  }
	//
    // 参见https://blog.golang.org/pipelines，
    // (2014年的博客：Go Concurrency Patterns: Pipelines and cancellation)
    // 以了解更多关于如何使用Done通道取消的例子。
	Done() <-chan struct{}

    // 如果Done尚未关闭，则Err返回nil。
	// 如果Done已关闭，则Err返回一个非nil错误，解释原因：
// 如果上下文被取消，则为Canceled；
// 如果上下文的截止时间过去，则为DeadlineExceeded。
// 在Err返回非nil错误之后，连续调用Err将返回相同的错误。
	Err() error

    // Value返回与此上下文相关的key的值，
    // 如果没有与key相关的值，则返回nil。
    // 用相同的key连续调用Value会返回相同的结果。
	//
	// 只有在请求范围内的数据穿越进程和API时才使用上下文值，
    // 而不是将可选参数传递给函数。
	//
    // key 标识了一个上下文中的特定值。
    // 希望在Context中存储值的函数通常会在一个全局变量中分配一个key，
    // 然后使用该key作为context.WithValue和Context.Value的实参。
    // key可以是任何支持可比较的类型。
	// 包应该将key定义为一个不可导出的类型，以避免冲突。
	//
	// 定义Context的key的包应该为使用该key存储的值提供类型安全的访问器：
	//	
	// // user包定义了一个存储在Contexts中的User类型。
	// 	package user
	//
	// 	import "context"
	//
	// // User是存储在Contexts中的值的类型。
	// 	type User struct {...}
	//
	//	// key是本包中定义的键的不可导出类型。
	//  //这可以防止与其他包中定义的键发生冲突。
	// 	type key int
	//
	//	// userKey是Contexts中user.User值的key。
    //  // 它是未被导出的；客户端使用user.NewContext和user.FromContext而不是直接使用这个key。
	// 	var userKey key
	//
	//	// NewContext返回一个新的带有u值的Context。
	// 	func NewContext(ctx context.Context, u *User) context.Context {
	// 		return context.WithValue(ctx, userKey, u)
	// 	}
	//
	//	// FromContext返回存储在ctx中的User值(如果有的话)。
	// 	func FromContext(ctx context.Context) (*User, bool) {
	// 		u, ok := ctx.Value(userKey).(*User)
	// 		return u, ok
	// 	}
	Value(key any) any
}
```

​	Context传递截止时间、取消信号和其他值跨API边界。

​	Context的方法可以同时被多个goroutine调用。

#### func Background 

```go 
func Background() Context
```

​	Background函数返回一个非nil、空的Context。它永远不会被取消，没有值，也没有截止时间。它通常用在main函数、初始化和测试以及作为传入请求的顶级Context使用。

#### func TODO 

```go 
func TODO() Context
```

​	TODO函数返回一个非nil、空的Context。当不清楚使用哪个Context或尚未可用(因为周围的函数尚未扩展为接受Context参数)时，代码应使用context.TODO。

#### func WithValue 

```go 
func WithValue(parent Context, key, val any) Context
```

​	WithValue函数返回parent的副本，其中与键关联的值为val。

​	仅将context Values用于跨进程和API传递的请求范围数据，而不是将可选参数传递给函数。

​	所提供的key 必须是可比较的，不应该是字符串或任何其他内置类型，以避免不同包之间的冲突。使用WithValue的用户应该为它们自己的key 定义类型。为了避免在分配给interface{}时分配内存，context键通常具有具体类型struct{}。或者，导出的上下文key 变量的静态类型应该是指针或接口。

##### WithValue Example
``` go 
package main

import (
	"context"
	"fmt"
)

func main() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

}
Output:

found value: Go
key not found: color
```

