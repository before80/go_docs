+++
title = "context"
date = 2023-05-17T09:59:21+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/context@go1.21.3](https://pkg.go.dev/context@go1.21.3)

Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

​	`context`包定义了 `Context` 类型，它在 API 边界和进程之间传递截止时间、取消信号和其他请求作用域值。

Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context. The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled.

​	服务端的传入请求应该创建一个 `Context`，对服务端的外部调用应该接受一个 Context。它们之间的函数调用链必须传播 `Context`，并且可以使用 `WithCancel`、`WithDeadline`、`WithTimeout` 或 `WithValue` 创建派生的 `Context` 来替换它。当一个 `Context` 被取消时，所有从它派生的 `Context` 也会被取消。	

The WithCancel, WithDeadline, and WithTimeout functions take a Context (the parent) and return a derived Context (the child) and a CancelFunc. Calling the CancelFunc cancels the child and its children, removes the parent's reference to the child, and stops any associated timers. Failing to call the CancelFunc leaks the child and its children until the parent is canceled or the timer fires. The go vet tool checks that CancelFuncs are used on all control-flow paths.

​	`WithCancel`、`WithDeadline` 和 `WithTimeout` 函数接受一个 `Context`(父级)并返回一个派生的 Context(子级)和一个 `CancelFunc`。调用 CancelFunc 会取消子级及其子级，移除父级对子级的引用并停止任何相关的定时器。如果不调用 `CancelFunc`，则会泄漏子级及其子级，直到父级被取消或定时器触发。go vet 工具检查所有控制流路径上是否使用了 `CancelFuncs`。

The WithCancelCause function returns a CancelCauseFunc, which takes an error and records it as the cancellation cause. Calling Cause on the canceled context or any of its children retrieves the cause. If no cause is specified, Cause(ctx) returns the same value as ctx.Err().

​	`WithCancelCause` 函数返回一个 `CancelCauseFunc`，它接受一个错误并将其记录为取消原因。调用取消的 `Context` 或任何其子级的 Cause 函数都会检索到取消原因。如果未指定原因，则 `Cause(ctx)` 返回与 `ctx.Err()` 相同的值。	

Programs that use Contexts should follow these rules to keep interfaces consistent across packages and enable static analysis tools to check context propagation:

​	使用 `Context` 的程序应该遵循以下规则，以使接口在包之间保持一致并启用静态分析工具检查上下文传播：	

Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named ctx:

(1)不要在结构类型中存储 `Context`；相反，将 `Context` 显式传递给需要它的每个函数。`Context` 应该是第一个参数，通常命名为 `ctx`：

```go 
func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}
```

Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use.

(2)即使函数允许，也不要传递 nil `Context`。如果不确定要使用哪个 `Context`，请传递 `context.TODO`。

Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.

(3)仅将 context `Value` 用于跨进程和 API 传递的请求作用域数据，而不是将可选参数传递给函数。

The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.

(4)同一个 `Context` 可以传递给在不同 goroutine 中运行的函数；`Context` 可以同时被多个 goroutine 安全使用。

See https://blog.golang.org/context for example code for a server that uses Contexts.

​	有关使用 `Context` 的示例代码，请参见[博客《Go并发模式：Context》]({{< ref "/goBlog/2014/GoConcurrencyPatternsContext" >}})。


## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/context/context.go;l=163)

```go 
var Canceled = errors.New("context canceled")
```

Canceled is the error returned by Context.Err when the context is canceled.

​	`Canceled`是当上下文被取消时，Context.Err返回的错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/context/context.go;l=167)

```go 
var DeadlineExceeded error = deadlineExceededError{}
```

DeadlineExceeded is the error returned by Context.Err when the context's deadline passes.

​	`DeadlineExceeded`是当上下文的截止时间过期时，Context.Err返回的错误。

## 函数

### func AfterFunc <- go1.21.0

```go 
func AfterFunc(ctx Context, f func()) (stop func() bool)
```

AfterFunc arranges to call f in its own goroutine after ctx is done (cancelled or timed out). If ctx is already done, AfterFunc calls f immediately in its own goroutine.

​	`AfterFunc`函数安排在`ctx`完成后（被取消或超时）在其自己的goroutine中调用`f`。如果`ctx`已经完成，`AfterFunc`会立即在其自己的goroutine中调用`f`。

Multiple calls to AfterFunc on a context operate independently; one does not replace another.

​	在一个上下文上对`AfterFunc`的多次调用是独立操作的；一次调用并不会替换另一次。

Calling the returned stop function stops the association of ctx with f. It returns true if the call stopped f from being run. If stop returns false, either the context is done and f has been started in its own goroutine; or f was already stopped. The stop function does not wait for f to complete before returning. If the caller needs to know whether f is completed, it must coordinate with f explicitly.

​	调用返回的`stop`函数会停止将`ctx`与`f`关联起来。如果调用停止了`f`的运行，它会返回`true`。如果`stop`函数返回`false`，则要么是上下文已完成并且`f`已在其自己的goroutine中开始；要么`f`已经被停止。`stop`函数不会等待`f`完成就返回。如果调用者需要知道`f`是否已完成，它必须与`f`显式协调。

If ctx has a "AfterFunc(func()) func() bool" method, AfterFunc will use it to schedule the call.

​	如果`ctx`有一个"`AfterFunc(func()) func() bool`"方法，`AfterFunc`将使用它来安排调用。

#### Example (Cond)

This example uses AfterFunc to define a function which waits on a sync.Cond, stopping the wait when a context is canceled.

​	此示例使用`AfterFunc`定义了一个函数，该函数等待一个`sync.Cond`，当上下文被取消时停止等待。

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	waitOnCond := func(ctx context.Context, cond *sync.Cond, conditionMet func() bool) error {
		stopf := context.AfterFunc(ctx, func() {
			// We need to acquire cond.L here to be sure that the Broadcast
			// below won't occur before the call to Wait, which would result
			// in a missed signal (and deadlock).
            // 我们需要在这里获取 cond.L，以确保在调用 Wait 之前不会发生下面的 Broadcast，
			// 其将导致信号丢失（和死锁）。
			cond.L.Lock()
			defer cond.L.Unlock()
            
			// If multiple goroutines are waiting on cond simultaneously,
			// we need to make sure we wake up exactly this one.
			// That means that we need to Broadcast to all of the goroutines,
			// which will wake them all up.
            // 如果多个 goroutine 同时在 cond 上等待，
			// 我们需要确保我们只唤醒这一个。
			// 这意味着我们需要 Broadcast 到所有 goroutine，
			// 这将唤醒它们所有。
			//            
			// If there are N concurrent calls to waitOnCond, each of the goroutines
			// will spuriously wake up O(N) other goroutines that aren't ready yet,
			// so this will cause the overall CPU cost to be O(N²).
            // 如果有 N 个并发的 waitOnCond 调用，每个 goroutine
			// 都会让 O(N) 其他尚未准备好的 goroutine 被虚假唤醒，
			// 所以这将导致总体 CPU 成本为 O(N²)。
			cond.Broadcast()
		})
		defer stopf()

		// Since the wakeups are using Broadcast instead of Signal, this call to
		// Wait may unblock due to some other goroutine's context becoming done,
		// so to be sure that ctx is actually done we need to check it in a loop.
        // 由于唤醒使用的是 Broadcast 而不是 Signal，所以调用 Wait 可能由于其他 goroutine 的上下文完成而取消，
		// 因此为了确保 ctx 实际上已完成，我们需要在循环中检查它。
		for !conditionMet() {
			cond.Wait()
			if ctx.Err() != nil {
				return ctx.Err()
			}
		}

		return nil
	}

	cond := sync.NewCond(new(sync.Mutex))

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
			defer cancel()

			cond.L.Lock()
			defer cond.L.Unlock()

			err := waitOnCond(ctx, cond, func() bool { return false })
			fmt.Println(err)
		}()
	}
	wg.Wait()

}
Output:

context deadline exceeded
context deadline exceeded
context deadline exceeded
context deadline exceeded
```

#### Example (Connection)

This example uses AfterFunc to define a function which reads from a net.Conn, stopping the read when a context is canceled.

​	此示例使用`AfterFunc`定义了一个函数，该函数从一个`net.Conn`读取，当上下文被取消时停止读取。

```go
package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	readFromConn := func(ctx context.Context, conn net.Conn, b []byte) (n int, err error) {
		stopc := make(chan struct{})
		stop := context.AfterFunc(ctx, func() {
			conn.SetReadDeadline(time.Now())
			close(stopc)
		})
		n, err = conn.Read(b)
		if !stop() {
			// The AfterFunc was started.
			// Wait for it to complete, and reset the Conn's deadline.
            // AfterFunc已经启动。
			// 等待它完成，并重置Conn的截止时间。
			<-stopc
			conn.SetReadDeadline(time.Time{})
			return n, ctx.Err()
		}
		return n, err
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	conn, err := net.Dial(listener.Addr().Network(), listener.Addr().String())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	b := make([]byte, 1024)
	_, err = readFromConn(ctx, conn, b)
	fmt.Println(err)

}
Output:

context deadline exceeded
```

#### Example (Merge)

This example uses AfterFunc to define a function which combines the cancellation signals of two Contexts.

​	这个示例使用`AfterFunc`定义了一个函数，该函数组合了两个上下文的取消信号。

```go
package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	// mergeCancel returns a context that contains the values of ctx,
	// and which is canceled when either ctx or cancelCtx is canceled.
    // mergeCancel 返回一个包含ctx值的上下文，当ctx或cancelCtx被取消时，该上下文也将被取消。
	mergeCancel := func(ctx, cancelCtx context.Context) (context.Context, context.CancelFunc) {
		ctx, cancel := context.WithCancelCause(ctx)
		stop := context.AfterFunc(cancelCtx, func() {
			cancel(context.Cause(cancelCtx))
		})
		return ctx, func() {
			stop()
			cancel(context.Canceled)
		}
	}

	ctx1, cancel1 := context.WithCancelCause(context.Background())
	defer cancel1(errors.New("ctx1 canceled"))

	ctx2, cancel2 := context.WithCancelCause(context.Background())

	mergedCtx, mergedCancel := mergeCancel(ctx1, ctx2)
	defer mergedCancel()

	cancel2(errors.New("ctx2 canceled"))
	<-mergedCtx.Done()
	fmt.Println(context.Cause(mergedCtx))

}

```



### func Cause  <- go1.20

```go 
func Cause(c Context) error
```

Cause returns a non-nil error explaining why c was canceled. The first cancellation of c or one of its parents sets the cause. If that cancellation happened via a call to CancelCauseFunc(err), then Cause returns err. Otherwise Cause(c) returns the same value as c.Err(). Cause returns nil if c has not been canceled yet.

​	`Cause`函数返回一个非nil的错误，解释为什么`c`被取消。c或其父级的第一个取消设置原因。如果取消是通过对`CancelCauseFunc(err)`的调用进行的，则`Cause`返回err。否则，`Cause(c)`返回与`c.Err()`相同的值。如果`c`尚未被取消，则`Cause`返回nil。

### func WithCancel 

```go 
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```

WithCancel returns a copy of parent with a new Done channel. The returned context's Done channel is closed when the returned cancel function is called or when the parent context's Done channel is closed, whichever happens first.

​	`WithCancel`函数返回`parent`的副本并创建一个新的`Done`通道。当调用返回的`cancel`函数或父级上下文的Done通道关闭时，返回的上下文的`Done`通道关闭。

Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete.

​	取消此上下文会释放与之相关联的资源，因此代码应尽快调用`cancel`，以便在此上下文中运行的操作完成。

#### WithCancel Example

This example demonstrates the use of a cancelable context to prevent a goroutine leak. By the end of the example function, the goroutine started by gen will return without leaking.

​	这个例子演示了使用可取消的上下文来防止 Goroutine 泄漏。在例子函数的结尾处，由 gen 启动的 Goroutine 将在不泄漏的情况下返回。

``` go 
package main

import (
	"context"
	"fmt"
)

func main() {
    // gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
    // gen在一个单独的goroutine中生成整数，并将它们发送到返回的通道中。
	// gen的调用者在消耗完生成的整数后需要取消上下文，以免泄露gen启动的内部goroutine。
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // 返回时不泄露goroutine的信息 cancel when we are finished consuming integers
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

### func WithCancelCause  <- go1.20

```go 
func WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc)
```

WithCancelCause behaves like WithCancel but returns a CancelCauseFunc instead of a CancelFunc. Calling cancel with a non-nil error (the "cause") records that error in ctx; it can then be retrieved using Cause(ctx). Calling cancel with nil sets the cause to Canceled.

​	`WithCancelCause`函数类似于`WithCancel`但返回一个`CancelCauseFunc`而不是`CancelFunc`。使用非`nil`错误(the "cause")调用`cancel`将记录该错误在ctx中;然后可以使用`Cause(ctx)`检索它。使用nil调用`cancel`将原因设置为已取消。

Example use:

例如使用：

```go 
ctx, cancel := context.WithCancelCause(parent)
cancel(myError)
ctx.Err() // returns context.Canceled
context.Cause(ctx) // returns myError
```

### func WithDeadline 

```go 
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
```

WithDeadline returns a copy of the parent context with the deadline adjusted to be no later than d. If the parent's deadline is already earlier than d, WithDeadline(parent, d) is semantically equivalent to parent. The returned context's Done channel is closed when the deadline expires, when the returned cancel function is called, or when the parent context's Done channel is closed, whichever happens first.

​	`WithDeadline`函数返回父`Context`的副本，其截止时间早于或等于`d`。如果父`Context`的截止时间早于`d`，则`WithDeadline(parent, d)`在语义上等同于`parent`。返回的Context的`Done`通道在到期时关闭，当调用返回的cancel函数时关闭，或者当父Context的`Done`通道关闭时关闭，以先发生的事件为准。

Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete.

​	取消此上下文会释放与其关联的资源，因此代码应在此Context中运行的操作完成后尽快调用`cancel`。

#### WithDeadline Example

This example passes a context with an arbitrary deadline to tell a blocking function that it should abandon its work as soon as it gets to it.

​	此示例传递了一个带有任意deadline 的上下文，以告诉阻塞函数，它应该在到达deadline 后立即放弃它的工作。

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

    // Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
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

### func WithDeadlineCause <-go1.21.0

```go
func WithDeadlineCause(parent Context, d time.Time, cause error) (Context, CancelFunc)
```

WithDeadlineCause behaves like [WithDeadline](https://pkg.go.dev/context@go1.21.3#WithDeadline) but also sets the cause of the returned Context when the deadline is exceeded. The returned [CancelFunc](https://pkg.go.dev/context@go1.21.3#CancelFunc) does not set the cause.

​	`WithDeadlineCause`的行为类似于[WithDeadline](https://pkg.go.dev/context@go1.21.3#WithDeadline)，但当截止时间超过时，还会设置返回的上下文的原因。返回的[CancelFunc](https://pkg.go.dev/context@go1.21.3#CancelFunc)不会设置原因。

### func WithTimeout 

```go 
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```

WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).

​	`WithTimeout`函数返回`WithDeadline(parent, time.Now().Add(timeout))`。

Canceling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete:

​	取消此Context会释放与其关联的资源，因此代码应在此Context中运行的操作完成后尽快调用cancel：

```go 
func slowOperationWithTimeout(ctx context.Context) (Result, error) {
	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()  // 如果slowOperation在超时前完成，则释放资源 releases resources if slowOperation completes before timeout elapses
	return slowOperation(ctx)
}
```

#### WithTimeout Example

This example passes a context with a timeout to tell a blocking function that it should abandon its work after the timeout elapses.

​	此示例传递了一个带有超时的上下文，以告诉阻塞函数在超时过后应该放弃它的工作。

``` go 
package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 1 * time.Millisecond

func main() {
    // Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
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

### func WithTimeoutCause <-go1.21.0

```go
func WithTimeoutCause(parent Context, timeout time.Duration, cause error) (Context, CancelFunc)
```

WithTimeoutCause behaves like [WithTimeout](https://pkg.go.dev/context@go1.21.3#WithTimeout) but also sets the cause of the returned Context when the timeout expires. The returned [CancelFunc](https://pkg.go.dev/context@go1.21.3#CancelFunc) does not set the cause.

​	`WithTimeoutCause`的行为类似于[WithTimeout](https://pkg.go.dev/context@go1.21.3#WithTimeout)，但当超时时间到期时，还会设置返回的上下文的原因。返回的[CancelFunc](https://pkg.go.dev/context@go1.21.3#CancelFunc)不会设置原因。

## 类型

### type CancelCauseFunc  <- go1.20

```go 
type CancelCauseFunc func(cause error)
```

A CancelCauseFunc behaves like a CancelFunc but additionally sets the cancellation cause. This cause can be retrieved by calling Cause on the canceled Context or on any of its derived Contexts.

​	`CancelCauseFunc`类型的行为类似于`CancelFunc`，但还会设置取消原因。该原因可以通过在取消的Context或其派生的Context上调用`Cause`来检索。

If the context has already been canceled, CancelCauseFunc does not set the cause. For example, if childContext is derived from parentContext:

​	如果上下文已经被取消，则`CancelCauseFunc`不会设置取消原因。例如，如果`childContext`是从`parentContext`派生的： 

- if parentContext is canceled with cause1 before childContext is canceled with cause2, then Cause(parentContext) == Cause(childContext) == cause1

- 如果`parentContext`在`childContext`之前以`cause1`取消，则`Cause(parentContext) == Cause(childContext) == cause1`  (即 parentContext 可以影响到 childContext )
- if childContext is canceled with cause2 before parentContext is canceled with cause1, then Cause(parentContext) == cause1 and Cause(childContext) == cause2
- 如果`childContext`在`parentContext`之前以`cause2`取消，则`Cause(parentContext) == cause1`，并且`Cause(childContext) == cause2`。(即 `childContext` 影响不到 `parentContext` )

### type CancelFunc 

```go 
type CancelFunc func()
```

A CancelFunc tells an operation to abandon its work. A CancelFunc does not wait for the work to stop. A CancelFunc may be called by multiple goroutines simultaneously. After the first call, subsequent calls to a CancelFunc do nothing.

​	`CancelFunc`类型告诉操作放弃它的工作。`CancelFunc`类型不等待工作停止。`多个goroutine可以同时调用CancelFunc`。在第一次调用之后，对`CancelFunc`的后续调用不起作用。

### type Context 

```go 
type Context interface {
    // Deadline returns the time when work done on behalf of this context
	// should be canceled. Deadline returns ok==false when no deadline is
	// set. Successive calls to Deadline return the same results.
    // Deadline返回代表该上下文所做的工作应该被取消的时间。
    // 如果没有设置deadline，Deadline方法返回ok==false。
    // 对Deadline的连续调用会返回相同的结果。
	Deadline() (deadline time.Time, ok bool)

    // Done returns a channel that's closed when work done on behalf of this
	// context should be canceled. Done may return nil if this context can
	// never be canceled. Successive calls to Done return the same value.
	// The close of the Done channel may happen asynchronously,
	// after the cancel function returns.
    // Done返回一个通道，该通道在此上下文代表的工作应被取消时关闭。
    // 如果这个上下文永远不能被取消，则Done可能会返回nil。
    // 对Done的连续调用会返回相同的值。
    // Done通道的关闭可以异步发生，在cancel函数返回之后。
	//
    // WithCancel arranges for Done to be closed when cancel is called;
	// WithDeadline arranges for Done to be closed when the deadline
	// expires; WithTimeout arranges for Done to be closed when the timeout
	// elapses.
	// WithCancel安排在调用cancel时关闭Done；
	// WithDeadline安排在截止时间过期时关闭Done；
	// WithTimeout安排在超时时间过去时关闭Done。
	//
    // Done is provided for use in select statements:
    // Done用于在select语句中使用：
	//
    //  // Stream generates values with DoSomething and sends them to out
	//  // until DoSomething returns an error or ctx.Done is closed.
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
    // See https://blog.golang.org/pipelines for more examples of how to use
	// a Done channel for cancellation.
    // 参见https://blog.golang.org/pipelines，
    // (2014年的博客：Go Concurrency Patterns: Pipelines and cancellation)
    // 以了解更多关于如何使用Done通道取消的例子。
	Done() <-chan struct{}

    // If Done is not yet closed, Err returns nil.
	// If Done is closed, Err returns a non-nil error explaining why:
	// Canceled if the context was canceled
	// or DeadlineExceeded if the context's deadline passed.
	// After Err returns a non-nil error, successive calls to Err return the same error.
    // 如果Done尚未关闭，则Err返回nil。
	// 如果Done已关闭，则Err返回一个非nil错误，解释原因：
	// 如果上下文被取消，则为Canceled；
	// 如果上下文的截止时间过去，则为DeadlineExceeded。
	// 在Err返回非nil错误之后，连续调用Err将返回相同的错误。
	Err() error

    // Value returns the value associated with this context for key, or nil
	// if no value is associated with key. Successive calls to Value with
	// the same key returns the same result.
    // Value返回与此上下文相关的key的值，
    // 如果没有与key相关的值，则返回nil。
    // 用相同的key连续调用Value会返回相同的结果。
	//
    // Use context values only for request-scoped data that transits
	// processes and API boundaries, not for passing optional parameters to
	// functions.
	// 只有在请求范围内的数据穿越进程和API时才使用上下文值，
    // 而不是将可选参数传递给函数。
	//
    // A key identifies a specific value in a Context. Functions that wish
	// to store values in Context typically allocate a key in a global
	// variable then use that key as the argument to context.WithValue and
	// Context.Value. A key can be any type that supports equality;
	// packages should define keys as an unexported type to avoid
	// collisions.
    // key 标识了一个上下文中的特定值。
    // 希望在Context中存储值的函数通常会在一个全局变量中分配一个key，
    // 然后使用该key作为context.WithValue和Context.Value的实参。
    // key可以是任何支持可比较的类型。
	// 包应该将key定义为一个不可导出的类型，以避免冲突。
	//
    // Packages that define a Context key should provide type-safe accessors
	// for the values stored using that key:
	// 定义Context的key的包应该为使用该key存储的值提供类型安全的访问器：
	//	
    // // Package user defines a User type that's stored in Contexts.
	// // user包定义了一个存储在Contexts中的User类型。
	// 	package user
	//
	// 	import "context"
	//
    // 	// key is an unexported type for keys defined in this package.
	// 	// This prevents collisions with keys defined in other packages.
    // // User is the type of value stored in the Contexts.
	// // User是存储在Contexts中的值的类型。
	// 	type User struct {...}
	//
    // 	// userKey is the key for user.User values in Contexts. It is
	// 	// unexported; clients use user.NewContext and user.FromContext
	// 	// instead of using this key directly.
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
    // 	// FromContext returns the User value stored in ctx, if any.
	//	// FromContext返回存储在ctx中的User值(如果有的话)。
	// 	func FromContext(ctx context.Context) (*User, bool) {
	// 		u, ok := ctx.Value(userKey).(*User)
	// 		return u, ok
	// 	}
	Value(key any) any
}
```

A Context carries a deadline, a cancellation signal, and other values across API boundaries.

​	`Context`传递截止时间、取消信号和其他值跨API边界。

Context's methods may be called by multiple goroutines simultaneously.

​	`Context`的方法可以同时被多个goroutine调用。

#### func Background 

```go 
func Background() Context
```

Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.

​	`Background`函数返回一个非nil、空的Context。它永远不会被取消，没有值，也没有截止时间。它通常用在main函数、初始化和测试以及作为传入请求的顶级Context使用。

#### func TODO 

```go 
func TODO() Context
```

TODO returns a non-nil, empty Context. Code should use context.TODO when it's unclear which Context to use or it is not yet available (because the surrounding function has not yet been extended to accept a Context parameter).

​	`TODO`函数返回一个非nil、空的`Context`。当不清楚使用哪个Context或尚未可用(因为周围的函数尚未扩展为接受Context参数)时，代码应使用`context.TODO`。

#### func WithValue 

```go 
func WithValue(parent Context, key, val any) Context
```

WithValue returns a copy of parent in which the value associated with key is val.

​	`WithValue`函数返回`parent`的副本，其中与键关联的值为`val`。

Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.

​	仅将context Values用于跨进程和API传递的请求范围数据，而不是将可选参数传递给函数。

The provided key must be comparable and should not be of type string or any other built-in type to avoid collisions between packages using context. Users of WithValue should define their own types for keys. To avoid allocating when assigning to an interface{}, context keys often have concrete type struct{}. Alternatively, exported context key variables' static type should be a pointer or interface.

​	所提供的key 必须是可比较的，不应该是字符串或任何其他内置类型，以避免不同包之间的冲突。使用`WithValue`的用户应该为它们自己的`key` 定义类型。为了避免在分配给`interface{}`时分配内存，`context`键通常具有具体类型`struct{}`。或者，导出的上下文`key` 变量的静态类型应该是指针或接口。

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

#### func WithoutCancel <-go1.21.0

```go
func WithoutCancel(parent Context) Context
```

WithoutCancel returns a copy of parent that is not canceled when parent is canceled. The returned context returns no Deadline or Err, and its Done channel is nil. Calling [Cause](https://pkg.go.dev/context@go1.21.3#Cause) on the returned context returns nil.

​	`WithoutCancel` 函数返回 `parent` 的一个副本，当 `parent` 被取消时，它不会被取消。返回的上下文没有 `Deadline` 或 `Err`，并且其 `Done` 通道为 nil。在返回的上下文上调用 [Cause](https://pkg.go.dev/context@go1.21.3#Cause) 返回 nil。

