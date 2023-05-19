+++
title = "sync"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# sync

https://pkg.go.dev/sync@go1.20.1

​	 sync包提供了基本的同步原语，如互斥锁。除了 Once 和 WaitGroup 类型外，大多数都是用于低级库例程的。更高级别的同步最好通过通道和通信实现。

Values containing the types defined in this package should not be copied.

​	包含此包中定义的类型的值不应被复制。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Cond 

``` go 
type Cond struct {

	// 在观察或更改条件时保持 L
	L Locker
	// 包含过滤或未公开的字段
}
```

​	Cond结构体实现条件变量，是用于等待或宣布事件发生的 goroutine 等待的交汇点。

​	每个 Cond 都有一个关联的 Locker L(通常是 `*Mutex` 或 `*RWMutex`)，在更改条件和调用 Wait 方法时必须持有该锁。

​	Cond 在首次使用后不得复制。

​	Cond 在 Go 内存模型的术语中，使 Broadcast 或 Signal 调用在解除其阻塞的任何 Wait 调用"之前同步"。

​	对于许多简单的用例，用户最好使用通道而不是 Cond(Broadcast 对应于关闭通道，而 Signal 对应于在通道上发送)。

​	有关替代 sync.Cond 的更多信息，请参见 [Roberto Clapis 的高级并发模式系列](https://blogtitle.github.io/categories/concurrency/)以及 [Bryan Mills 的并发模式演讲](https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view)。

#### func NewCond 

``` go 
func NewCond(l Locker) *Cond
```

​	NewCond函数返回一个带有 Locker l 的新 Cond。

#### (*Cond) Broadcast 

``` go 
func (c *Cond) Broadcast()
```

​	Broadcast方法唤醒正在c上等待的所有goroutine。

​	调用者可以在调用时持有c.L，但不是必须的。

#### (*Cond) Signal

```
func (c *Cond) Signal()
```

​	Signal 方法会唤醒一个等待在 c 上的 goroutine，如果有的话。

​	调用 Signal 时，允许但不要求调用方持有 c.L。

​	Signal() 不会影响 goroutine 的调度优先级；如果其他 goroutine 正在尝试锁定 c.L，则它们可能会在"等待(waiting)"goroutine之前被唤醒。

#### (*Cond) Wait 

``` go 
func (c *Cond) Wait()
```

​	Wait方法原子性地解锁 c.L 并挂起调用的 goroutine 的执行。稍后恢复执行后，Wait 在返回之前锁定 c.L。与其他系统不同，除非被 Broadcast 或 Signal 唤醒，Wait 不能返回。

​	由于 Wait方法在等待时未锁定 c.L，因此调用者通常不能假设在 Wait 返回时条件为真。相反，调用者应在循环中等待：

```
c.L.Lock()
for !condition() {
    c.Wait()
}
... make use of condition ...
c.L.Unlock()
```

### type Locker 

``` go 
type Locker interface {
	Lock()
	Unlock()
}
```

​	Locker接口表示一个可以锁定和解锁的对象。

### type Map  <- go1.9

``` go 
type Map struct {
	// contains filtered or unexported fields
}
```

​	Map结构体类似于 Go map[interface{}]interface{}，但可以在多个 goroutine 中安全地使用，而无需额外的锁定或协调。负载、存储和删除以平摊常数时间运行。

​	Map类型是特定的。大多数代码应该使用普通的Go map，通过单独的锁或协调来实现更好的类型安全性，并使其更容易维护map内容的其他不变量。

​	Map类型针对两种常见用例进行了优化：(1) 给定键的条目仅被写入一次，但被多次读取，例如仅增长的缓存；或者(2) 多个goroutine读取、写入和覆盖不重叠的键集的条目。在这两种情况下，与配对单独的Mutex或RWMutex的Go map相比，使用Map可能会显著降低锁争用。

​	零值Map为空且可以直接使用。Map在第一次使用后不得复制。

​	根据Go内存模型的术语，Map安排写操作"在"观察到写操作的效果的任何读取操作"之前"，其中读取和写操作定义如下。Load，LoadAndDelete，LoadOrStore，Swap，CompareAndSwap和CompareAndDelete是读取操作; Delete，LoadAndDelete，Store和Swap是写操作;当LoadOrStore返回loaded设置为false时，它是一个写操作;当CompareAndSwap返回swapped设置为true时，它是一个写操作;当CompareAndDelete返回deleted设置为true时，它是一个写操作。

#### (*Map) CompareAndDelete  <- go1.20

``` go 
func (m *Map) CompareAndDelete(key, old any) (deleted bool)
```

​	CompareAndDelete方法会删除键为key的条目，如果它的值等于old。旧值必须是可比较类型。

​	如果map中没有当前键的值，则CompareAndDelete方法返回false(即使旧值是nil接口值)。

#### (*Map) CompareAndSwap  <- go1.20

``` go 
func (m *Map) CompareAndSwap(key, old, new any) bool
```

​	CompareAndSwap方法会交换键的旧值和新值，如果存储在map中的值等于旧值。旧值必须是可比较类型。

#### (*Map) Delete  <- go1.9

``` go 
func (m *Map) Delete(key any)
```

​	Delete方法用于删除一个键的值。

#### (*Map) Load  <- go1.9

``` go 
func (m *Map) Load(key any) (value any, ok bool)
```

​	Load方法用于获取一个键对应的值。如果键不存在，返回值value为nil，ok为false。

#### (*Map) LoadAndDelete  <- go1.15

``` go 
func (m *Map) LoadAndDelete(key any) (value any, loaded bool)
```

​	LoadAndDelete方法用于获取一个键对应的值并从Map中删除该键值对。如果键不存在，则返回值value为nil，loaded为false。

#### (*Map) LoadOrStore  <- go1.9

``` go 
func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)
```

​	LoadOrStore方法用于获取一个键对应的值。如果键存在，返回值actual为已存在的值，loaded为true。如果键不存在，会将给定的值存储并返回，此时loaded为false。

#### (*Map) Range  <- go1.9

``` go 
func (m *Map) Range(f func(key, value any) bool)
```

​	Range方法遍历Map中的每个键值对，并按顺序调用参数f。如果f返回false，则停止迭代。

​	Range方法并不一定对应Map内容的一致快照。如果在Range调用期间并发地存储或删除任何键的值(包括在f中)，则Range可能会反映该键的任何时点的任何映射。Range方法不会阻塞接收器上的其他方法；即使f本身也可以调用m上的任何方法。

​	即使f返回false，Range方法的时间复杂度可能为O(N)，其中N为Map中元素的数量。

#### (*Map) Store  <- go1.9

``` go 
func (m *Map) Store(key, value any)
```

​	Store方法用于设置一个键的值。

#### (*Map) Swap  <- go1.20

``` go 
func (m *Map) Swap(key, value any) (previous any, loaded bool)
```

​	Swap方法交换给定键的值并返回先前的值(如果存在)。loaded结果报告键是否存在。

### type Mutex 

``` go 
type Mutex struct {
	// contains filtered or unexported fields
}
```

​	Mutex类型是互斥锁。互斥锁的零值是未锁定的互斥锁。

​	Mutex在第一次使用后不能被复制。

​	根据Go内存模型的术语，第n次Unlock调用"在" m次调用Lock之前同步，其中n < m。成功调用TryLock等效于调用Lock。TryLock的失败调用根本不建立任何"在之前同步"的关系。

#### (*Mutex) Lock 

``` go 
func (m *Mutex) Lock()
```

​	Lock方法锁定m。如果锁已经在使用中，则调用goroutine会阻塞，直到互斥锁可用。

#### (*Mutex) TryLock  <- go1.18

``` go 
func (m *Mutex) TryLock() bool
```

​	TryLock方法尝试锁定m并报告是否成功。

​	请注意，尽管存在正确使用TryLock方法，但它们很少出现，并且使用TryLock方法通常是特定互斥锁使用中更深层次问题的标志。

#### (*Mutex) Unlock 

``` go 
func (m *Mutex) Unlock()
```

​	Unlock解锁m。如果在进入Unlock时未锁定m，则会引发运行时错误。

​	锁定的互斥锁与特定goroutine不相关联。允许一个goroutine锁定Mutex，然后安排另一个goroutine解锁它。

### type Once 

``` go 
type Once struct {
	// contains filtered or unexported fields
}
```

​	Once 对象只会执行一次操作。

​	在第一次使用后，Once 不能被复制。

​	在 Go 内存模型的术语中，f 函数的返回值在 any 调用 once.Do(f) 的返回值之前"同步"。

##### Example
``` go 
package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
Output:

Only once
```

#### (*Once) Do 

``` go 
func (o *Once) Do(f func())
```

​	Do方法只在第一次针对该 Once 实例调用 Do 时调用函数 f。换句话说，给定

``` go 
var once Once
```

​	如果多次调用 once.Do(f)，则仅有第一次调用会调用 f，即使 f 在每次调用中的值不同，也是如此。对于每个需要执行的函数，都需要一个新的 Once 实例。

​	Do方法旨在进行必须仅运行一次的初始化。由于 f 是零参函数，可能需要使用函数文本来捕获由 Do 调用的函数的参数：

```
config.once.Do(func() { config.init(filename) })
```

​	因为在 f 返回之前，没有 Do方法调用会返回，所以如果 f 调用 Do，它将死锁。

​	如果 f 引发 panic，Do方法将视其为已返回；将来的 Do方法调用将在不调用 f 的情况下返回。

### type Pool  <- go1.3

``` go 
type Pool struct {

	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() any
	// contains filtered or unexported fields
}
```

​	类型 Pool 是一组可临时保存和检索的对象。

​	在任何时候，存储在 Pool 中的任何项都可能被自动删除而没有通知。如果在此时 Pool 持有唯一的引用，则可能会被解除分配。

​	Pool 可以同时安全地用于多个 goroutine。

​	Pool 的目的是缓存已分配但未使用的项目以供以后重用，从而减轻垃圾回收器的压力。也就是说，它使得构建高效、线程安全的自由列表变得容易。但是，它不适用于所有自由列表。

​	Pool 的一个合适的用途是管理一组在包的并发独立客户端之间隐式共享的临时项目，并且这些项目可能会被重用。Pool 提供了一种将分配开销分摊到许多客户端的方式。

​	fmt 包是 Pool 的一个很好的用例，它维护一个动态大小的临时输出缓冲区存储。当许多 goroutine 正在积极打印时，存储会扩展，而在空闲时会缩小。

​	另一方面，作为短暂对象的一部分维护的自由列表不适合用于 Pool，因为在这种情况下，开销无法很好地分摊。在这种情况下，让这些对象实现自己的自由列表更有效率。

​	Pool 在第一次使用后不得进行复制。

​	在 Go 内存模型的术语中，Put(x) 的调用"在" Get 返回同一值 x 之前同步。同样，New 返回 x 的调用"在" Get 返回同一值 x 之前同步。

##### Example
``` go 
package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() any {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(bytes.Buffer)
	},
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func main() {
	Log(os.Stdout, "path", "/search?q=flowers")
}
Output:

2006-01-02T15:04:05Z path=/search?q=flowers
```

#### (*Pool) Get  <- go1.3

``` go 
func (p *Pool) Get() any
```

​	Get 方法从池中选择一个任意项、将其从池中移除并返回给调用者。Get 可能会选择忽略池并将其视为空。调用者不应假设 Put 传递的值与 Get 返回的值之间有任何关系。

​	如果 Get方法原本会返回 nil，且 p.New 不为 nil，则 Get 返回调用 p.New 的结果。

#### (*Pool) Put  <- go1.3

``` go 
func (p *Pool) Put(x any)
```

​	Put 方法将 x 添加到池中。

### type RWMutex 

``` go 
type RWMutex struct {
	// contains filtered or unexported fields
}
```

​	RWMutex 是一个读写互斥锁。锁可以被任意数量的读取器或单个写入器持有。RWMutex 的零值为未锁定的互斥锁。

​	RWMutex 必须在第一次使用后不进行复制。

​	如果一个 goroutine 持有 RWMutex 进行读取，并且另一个 goroutine 可能调用 Lock，则没有 goroutine 应该期望能够在初始读取锁定被释放之前获取读取锁定。特别地，这禁止了递归读取锁定。这是为了确保锁最终变为可用；一个被阻塞的 Lock 调用排除了新的读取器获取锁定。

​	按照 Go 内存模型的术语，第 n 次调用 Unlock "在之前同步"第 m 次调用 Lock，就像 Mutex 一样。对于任何对 RLock 的调用，都存在一个 n，使得第 n 次调用 Unlock "在之前同步"该对 RLock 的调用，而相应的 RUnlock 调用"在之前同步"第 n+1 次调用 Lock。

#### (*RWMutex) Lock 

``` go 
func (rw *RWMutex) Lock()
```

​	Lock方法锁定 rw 进行写入。如果锁已被锁定以供读取或写入，则 Lock 阻塞，直到锁可用。

#### (*RWMutex) RLock 

``` go 
func (rw *RWMutex) RLock()
```

​	RLock 锁定 rw 进行读取。

​	它不应该用于递归读取锁定；一个被阻塞的 Lock 调用排除了新的读取器获取锁定。请参见 RWMutex 类型的文档。

#### (*RWMutex) RLocker 

``` go 
func (rw *RWMutex) RLocker() Locker
```

​	RLocker方法返回一个Locker接口，该接口通过调用rw.RLock和rw.RUnlock实现Lock和Unlock方法。

#### (*RWMutex) RUnlock 

``` go 
func (rw *RWMutex) RUnlock()
```

​	RUnlock方法撤消单个RLock调用；它不会影响其他同时读取者。如果rw在进入RUnlock时未被读取锁定，则会出现运行时错误。

#### (*RWMutex) TryLock  <- go1.18

``` go 
func (rw *RWMutex) TryLock() bool
```

​	TryLock方法尝试为写入锁定rw，并报告是否成功。

​	请注意，尽管TryLock方法的正确使用确实存在，但TryLock方法的使用往往是特定互斥锁使用中深层问题的迹象。

#### (*RWMutex) TryRLock  <- go1.18

``` go 
func (rw *RWMutex) TryRLock() bool
```

​	TryRLock方法尝试为读取锁定rw，并报告是否成功。

​	请注意，尽管TryRLock方法的正确使用确实存在，但TryRLock的使用往往是特定互斥锁使用中深层问题的迹象。

#### (*RWMutex) Unlock 

``` go 
func (rw *RWMutex) Unlock()
```

​	Unlock方法解锁rw以进行写入。如果在进入Unlock时rw未被写入锁定，则会出现运行时错误。

​	与Mutexes一样，锁定的RWMutex不与特定的goroutine相关联。一个goroutine可以RLock(Lock)RWMutex，然后安排另一个goroutine RUnlock(Unlock)它。

### type WaitGroup 

``` go 
type WaitGroup struct {
	// contains filtered or unexported fields
}
```

​	WaitGroup结构体等待一组goroutine完成。主goroutine调用Add来设置等待的goroutine数量。然后每个goroutine都运行并在完成时调用Done。同时，可以使用Wait来阻塞，直到所有goroutine完成。

​	WaitGroup在第一次使用后不得复制。

​	在Go内存模型的术语中，Done的调用"在"解除任何Wait调用之前"同步返回它的阻塞。

##### Example
``` go 
package main

import (
	"sync"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.example.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

```

#### (*WaitGroup) Add 

``` go 
func (wg *WaitGroup) Add(delta int)
```

​	添加delta(可以为负数)到WaitGroup计数器中。如果计数器变为零，则所有在Wait上阻塞的goroutine都会被释放。如果计数器变为负数，则Add会panic。

​	请注意，在计数器为零时具有正delta的调用必须在Wait之前发生。具有负delta的调用，或者在计数器大于零时开始具有正delta的调用，可以随时发生。通常，这意味着在创建goroutine或等待其他事件的语句之前执行Add调用。如果要重用WaitGroup以等待几组独立的事件，则所有先前的Wait调用都必须返回后才能发生新的Add调用。有关WaitGroup的示例，请参见官方文档。

#### (*WaitGroup) Done 

``` go 
func (wg *WaitGroup) Done()
```

​	Done方法将WaitGroup计数器减一。

#### (*WaitGroup) Wait 

``` go 
func (wg *WaitGroup) Wait()
```

​	Wait方法阻塞，直到WaitGroup计数器为零。