+++
title = "从 unique 到清理和弱引用：用于效率的新低级工具"
date = 2025-03-31T14:26:24+08:00
weight = 960
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/cleanups-and-weak](https://go.dev/blog/cleanups-and-weak)

# From unique to cleanups and weak: new low-level tools for efficiency - 从 unique 到清理和弱引用：用于效率的新低级工具

Michael Knyszek  

6 March 2025  

In [last year’s blog post](https://go.dev/blog/unique) about the `unique` package, we alluded to some new features then in proposal review, and we’re excited to share that as of Go 1.24 they are now available to all Go developers. These new features are [the `runtime.AddCleanup` function](https://pkg.go.dev/runtime#AddCleanup), which queues up a function to run when an object is no longer reachable, and [the `weak.Pointer` type](https://pkg.go.dev/weak#Pointer), which safely points to an object without preventing it from being garbage collected. Together, these two features are powerful enough to build your own `unique` package! Let’s dig into what makes these features useful, and when to use them.

​	在 [去年的博客文章](https://go.dev/blog/unique) 中，我们提到了 `unique` 包，并暗示了一些当时处于提案审查中的新功能。我们很高兴地宣布，自 Go 1.24 起，这些新功能现已对所有 Go 开发者开放。这些新功能包括 [ `runtime.AddCleanup` 函数](https://pkg.go.dev/runtime#AddCleanup)，它在对象不再可达时排队执行一个函数，以及 [ `weak.Pointer` 类型](https://pkg.go.dev/weak#Pointer)，它安全地指向一个对象而不阻止其被垃圾回收。这两个功能结合起来足以让你自己构建一个 `unique` 包！让我们深入探讨这些功能为何有用，以及何时使用它们。

> Note: these new features are advanced features of the garbage collector. If you’re not already familiar with basic garbage collection concepts, we strongly recommend reading the introduction of our [garbage collector guide](https://go.dev/doc/gc-guide#Introduction).
>
> ​	注意：这些新功能是垃圾回收器的高级功能。如果您还不熟悉基本的垃圾回收概念，我们强烈建议您阅读我们的 [垃圾回收指南](https://go.dev/doc/gc-guide#Introduction) 的介绍部分。
>

## Cleanups 清理

If you’ve ever used a finalizer, then the concept of a cleanup will be familiar. A finalizer is a function, associated with an allocated object by [calling `runtime.SetFinalizer`](https://pkg.go.dev/runtime#SetFinalizer), that is later called by the garbage collector some time after the object becomes unreachable. At a high level, cleanups work the same way.

​	如果您曾经使用过终结器（finalizer），那么清理（cleanup）的概念对您来说会很熟悉。终结器是一个函数，通过 [调用 `runtime.SetFinalizer`](https://pkg.go.dev/runtime#SetFinalizer) 与分配的对象关联，并在对象变得不可达后的某个时间由垃圾回收器调用。从高层次来看，清理的运作方式相同。

Let’s consider an application that makes use of a memory-mapped file, and see how cleanups can help.

​	让我们考虑一个使用内存映射文件的应用程序，看看清理如何提供帮助。

```
//go:build unix

type MemoryMappedFile struct {
    data []byte
}

func NewMemoryMappedFile(filename string) (*MemoryMappedFile, error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    // Get the file's info; we need its size.
    fi, err := f.Stat()
    if err != nil {
        return nil, err
    }

    // Extract the file descriptor.
    conn, err := f.SyscallConn()
    if err != nil {
        return nil, err
    }
    var data []byte
    connErr := conn.Control(func(fd uintptr) {
        // Create a memory mapping backed by this file.
        data, err = syscall.Mmap(int(fd), 0, int(fi.Size()), syscall.PROT_READ, syscall.MAP_SHARED)
    })
    if connErr != nil {
        return nil, connErr
    }
    if err != nil {
        return nil, err
    }
    mf := &MemoryMappedFile{data: data}
    cleanup := func(data []byte) {
        syscall.Munmap(data) // ignore error
    }
    runtime.AddCleanup(mf, cleanup, data)
    return mf, nil
}
```

A memory-mapped file has its contents mapped to memory, in this case, the underlying data of a byte slice. Thanks to operating-system magic, reads and writes to the byte slice directly access the contents of the file. With this code, we can pass around a `*MemoryMappedFile`, and when it’s no longer referenced, the memory mapping we created will get cleaned up.

​	一个内存映射文件将其内容映射到内存中，在此例中，是一个字节切片的底层数据。得益于操作系统的魔法，对字节切片的读写直接访问文件内容。通过这段代码，我们可以传递一个 `*MemoryMappedFile`，当它不再被引用时，我们创建的内存映射将被清理。

> Notice that `runtime.AddCleanup` takes three arguments: the address of a variable to attach the cleanup to, the cleanup function itself, and an argument to the cleanup function. A key difference between this function and `runtime.SetFinalizer` is that the cleanup function takes a different argument than the object we’re attaching the cleanup to. This change fixes some problems with finalizers.
>
> ​	注意，`runtime.AddCleanup` 接受三个参数：要附加清理的变量地址、清理函数本身以及清理函数的一个参数。此函数与 `runtime.SetFinalizer` 的一个关键区别在于，清理函数接受的参数与我们附加清理的对象不同。这一变化修复了终结器的一些问题。
>

It’s no secret that [finalizers are difficult to use correctly](https://go.dev/doc/gc-guide#Common_finalizer_issues). For example, objects to which finalizers are attached must not be involved in any reference cycles (even a pointer to itself is too much!), otherwise the object will never be reclaimed and the finalizer will never run, causing a leak. Finalizers also significantly delay reclamation of memory. It takes at a minimum two full garbage collection cycles to reclaim the memory for a finalized object: one to determine that it’s unreachable, and another to determine that it’s still unreachable after the finalizer executes.

​	众所周知，[终结器很难正确使用](https://go.dev/doc/gc-guide#Common_finalizer_issues)。例如，附加了终结器的对象不得参与任何引用循环（即使是自指向的指针也不行！），否则对象将永远不会被回收，终结器也不会运行，导致泄漏。终结器还会显著延迟内存回收。回收一个终结对象的内存至少需要两个完整的垃圾回收周期：一次确定其不可达，另一次在终结器执行后确定其仍不可达。

The problem is that finalizers [resurrect the objects they’re attached to](https://en.wikipedia.org/wiki/Object_resurrection). The finalizer doesn’t run until the object is unreachable, at which point it is considered “dead.” But since the finalizer is called with a pointer to the object, the garbage collector must prevent the reclamation of that object’s memory, and instead must generate a new reference for the finalizer, making it reachable, or “live,” once more. That reference may even remain after the finalizer returns, for example if the finalizer writes it to a global variable or sends it across a channel. Object resurrection is problematic because it means the object, and everything it points to, and everything those objects point to, and so on, is reachable, even if it would otherwise have been collected as garbage.

​	问题在于，终结器 [会复活它们所附加的对象](https://en.wikipedia.org/wiki/Object_resurrection)。终结器在对象不可达之前不会运行，此时对象被认为是“死的”。但由于终结器使用指向对象的指针调用，垃圾回收器必须阻止回收该对象的内存，而必须为终结器生成一个新引用，使其再次变得可达或“活的”。该引用甚至可能在终结器返回后仍然存在，例如如果终结器将其写入全局变量或通过通道发送。对象复活是有问题的，因为这意味着对象及其指向的所有内容，以及那些对象指向的内容等等，都是可达的，即使它们本应被作为垃圾回收。

We solve both of these problems by not passing the original object to the cleanup function. First, the values the object refers to don’t need to be kept specially reachable by the garbage collector, so the object can still be reclaimed even if it’s involved in a cycle. Second, since the object is not needed for the cleanup, its memory can be reclaimed immediately.

​	我们通过不将原始对象传递给清理函数解决了这两个问题。首先，对象引用的值无需被垃圾回收器特别保持可达，因此即使对象涉及循环，也可以被回收。其次，由于清理不需要对象，其内存可以立即被回收。

## Weak pointers 弱指针

Returning to our memory-mapped file example, suppose we notice that our program frequently maps the same files over and over, from different goroutines that are unaware of each other. This is fine from a memory perspective, since all these mappings will share physical memory, but it results in lots of unnecessary system calls to map and unmap the file. This is especially bad if each goroutine reads only a small section of each file.

​	回到我们的内存映射文件示例，假设我们注意到程序频繁地反复映射相同的文件，来自彼此不知情的不同 goroutine。从内存角度看这没问题，因为所有这些映射将共享物理内存，但这导致了许多不必要的系统调用来映射和取消映射文件。如果每个 goroutine 仅读取每个文件的一小部分，这尤其糟糕。

So, let’s deduplicate the mappings by filename. (Let’s assume that our program only reads from the mappings, and the files themselves are never modified or renamed once created. Such assumptions are reasonable for system font files, for example.)

​	因此，让我们按文件名去重映射。（假设我们的程序只从映射中读取，并且文件一旦创建后不会被修改或重命名。例如，对于系统字体文件，这样的假设是合理的。）

We could maintain a map from filename to memory mapping, but then it becomes unclear when it’s safe to remove entries from that map. We could *almost* use a cleanup, if it weren’t for the fact that the map entry itself will keep the memory-mapped file object alive.

​	我们可以维护一个从文件名到内存映射的映射，但那样就不清楚何时可以安全地从该映射中移除条目。如果不是因为映射条目本身会保持内存映射文件对象的存活，我们 *几乎* 可以使用清理。

Weak pointers solve this problem. A weak pointer is a special kind of pointer that the garbage collector ignores when deciding whether an object is reachable. Go 1.24’s [new weak pointer type, `weak.Pointer`](https://pkg.go.dev/weak#Pointer), has a `Value` method that returns either a real pointer if the object is still reachable, or `nil` if it is not.

​	弱指针解决了这个问题。弱指针是一种特殊指针，垃圾回收器在决定对象是否可达时会忽略它。Go 1.24 的 [新弱指针类型 `weak.Pointer`](https://pkg.go.dev/weak#Pointer) 有一个 `Value` 方法，如果对象仍可达则返回真实指针，否则返回 `nil`。

If we instead maintain a map that only *weakly* points to the memory-mapped file, we can clean up the map entry when nobody’s using it anymore! Let’s see what this looks like.

​	如果我们改为维护一个仅 *弱引用* 指向内存映射文件的映射，我们可以在没人使用它时清理该映射条目！让我们看看这是什么样子。

```
var cache sync.Map // map[string]weak.Pointer[MemoryMappedFile]

func NewCachedMemoryMappedFile(filename string) (*MemoryMappedFile, error) {
    var newFile *MemoryMappedFile
    for {
        // Try to load an existing value out of the cache.
        value, ok := cache.Load(filename)
        if !ok {
            // No value found. Create a new mapped file if needed.
            if newFile == nil {
                var err error
                newFile, err = NewMemoryMappedFile(filename)
                if err != nil {
                    return nil, err
                }
            }

            // Try to install the new mapped file.
            wp := weak.Make(newFile)
            var loaded bool
            value, loaded = cache.LoadOrStore(filename, wp)
            if !loaded {
                runtime.AddCleanup(newFile, func(filename string) {
                    // Only delete if the weak pointer is equal. If it's not, someone
                    // else already deleted the entry and installed a new mapped file.
                    cache.CompareAndDelete(filename, wp)
                }, filename)
                return newFile, nil
            }
            // Someone got to installing the file before us.
            //
            // If it's still there when we check in a moment, we'll discard newFile
            // and it'll get cleaned up by garbage collector.
        }

        // See if our cache entry is valid.
        if mf := value.(weak.Pointer[MemoryMappedFile]).Value(); mf != nil {
            return mf, nil
        }

        // Discovered a nil entry awaiting cleanup. Eagerly delete it.
        cache.CompareAndDelete(filename, value)
    }
}
```

This example is a little complicated, but the gist is simple. We start with a global concurrent map of all the mapped files we made. `NewCachedMemoryMappedFile` consults this map for an existing mapped file, and if that fails, creates and tries to insert a new mapped file. This could of course fail as well since we’re racing with other insertions, so we need to be careful about that too, and retry. (This design has a flaw in that we might wastefully map the same file multiple times in a race, and we’ll have to throw it away via the cleanup added by `NewMemoryMappedFile`. This is probably not a big deal most of the time. Fixing it is left as an exercise for the reader.)

​	这个例子有点复杂，但核心很简单。我们从一个全局并发映射开始，记录我们创建的所有映射文件。`NewCachedMemoryMappedFile` 查询此映射以获取现有的映射文件，如果失败，则创建并尝试插入一个新的映射文件。当然，这也可能失败，因为我们与其他插入操作竞争，所以我们需要小心处理，还要重试。（这个设计有一个缺陷，在竞争中我们可能会多次浪费地映射同一文件，并通过 `NewMemoryMappedFile` 添加的清理将其丢弃。大多数时候这可能不是大问题。修复它留给读者作为练习。）

Let’s look at some useful properties of weak pointers and cleanups exploited by this code.

​	让我们看看这段代码利用的弱指针和清理的一些有用属性。

Firstly, notice that weak pointers are comparable. Not only that, weak pointers have a stable and independent identity, which remains even after the objects they point to are long gone. This is why it is safe for the cleanup function to call `sync.Map`’s `CompareAndDelete`, which compares the `weak.Pointer`, and a crucial reason this code works at all.

​	首先，注意弱指针是可比较的。不仅如此，弱指针具有稳定且独立的身份，即使它们指向的对象早已消失，这种身份仍然存在。这就是为什么清理函数调用 `sync.Map` 的 `CompareAndDelete` 是安全的，它比较 `weak.Pointer`，这是这段代码能够工作的关键原因。

Secondly, observe that we can add multiple independent cleanups to a single `MemoryMappedFile` object. This allows us to use cleanups in a composable way and use them to build generic data structures. In this particular example, it might be more efficient to combine `NewCachedMemoryMappedFile` with `NewMemoryMappedFile` and have them share a cleanup. However, the advantage of the code we wrote above is that it can be rewritten in a generic way!

​	其次，观察到我们可以为单个 `MemoryMappedFile` 对象添加多个独立的清理。这允许我们以可组合的方式使用清理，并用它们构建通用数据结构。在这个特定示例中，将 `NewCachedMemoryMappedFile` 与 `NewMemoryMappedFile` 结合起来并共享一个清理可能更高效。然而，我们上面编写的代码的优势在于它可以以通用方式重写！

```
type Cache[K comparable, V any] struct {
    create func(K) (*V, error)
    m     sync.Map
}

func NewCache[K comparable, V any](create func(K) (*V, error)) *Cache[K, V] {
    return &Cache[K, V]{create: create}
}

func (c *Cache[K, V]) Get(key K) (*V, error) {
    var newValue *V
    for {
        // Try to load an existing value out of the cache.
        value, ok := cache.Load(key)
        if !ok {
            // No value found. Create a new mapped file if needed.
            if newValue == nil {
                var err error
                newValue, err = c.create(key)
                if err != nil {
                    return nil, err
                }
            }

            // Try to install the new mapped file.
            wp := weak.Make(newValue)
            var loaded bool
            value, loaded = cache.LoadOrStore(key, wp)
            if !loaded {
                runtime.AddCleanup(newValue, func(key K) {
                    // Only delete if the weak pointer is equal. If it's not, someone
                    // else already deleted the entry and installed a new mapped file.
                    cache.CompareAndDelete(key, wp)
                }, key)
                return newValue, nil
            }
        }

        // See if our cache entry is valid.
        if mf := value.(weak.Pointer[V]).Value(); mf != nil {
            return mf, nil
        }

        // Discovered a nil entry awaiting cleanup. Eagerly delete it.
        cache.CompareAndDelete(key, value)
    }
}
```

## Caveats and future work 注意事项和未来工作

Despite our best efforts, cleanups and weak pointers can still be error-prone. To guide those considering using finalizers, cleanups, and weak pointers, we recently updated the [guide to the garbage collector](https://go.dev/doc/gc-guide#Finalizers_cleanups_and_weak_pointers) with some advice about using these features. Take a look next time you reach for them, but also carefully consider whether you need to use them at all. These are advanced tools with subtle semantics and, as the guide says, most Go code benefits from these features indirectly, not from using them directly. Stick to the use-cases where these features shine, and you’ll be alright.

​	尽管我们尽了最大努力，清理和弱指针仍可能容易出错。为了指导那些考虑使用终结器、清理和弱指针的人，我们最近更新了 [垃圾回收指南](https://go.dev/doc/gc-guide#Finalizers_cleanups_and_weak_pointers)，提供了一些关于使用这些功能的建议。下次使用它们时请查看指南，但也要仔细考虑是否真的需要使用它们。这些是具有微妙语义的高级工具，正如指南所说，大多数 Go 代码从这些功能中获益是间接的，而不是直接使用。坚持使用这些功能擅长的用例，你会没事的。

For now, we’ll call out some of the issues that you are more likely to run into.

​	现在，我们将指出一些你更可能遇到的问题。

First, the object the cleanup is attached to must be reachable from neither the cleanup function (as a captured variable) nor the argument to the cleanup function. Both of these situations result in the cleanup never executing. (In the special case of the cleanup argument being exactly the pointer passed to `runtime.AddCleanup`, `runtime.AddCleanup` will panic, as a signal to the caller that they should not use cleanups the same way as finalizers.)

​	首先，清理所附加的对象不得从清理函数（作为捕获变量）或清理函数的参数中可达。这两种情况都会导致清理永不执行。（在清理参数恰好是传递给 `runtime.AddCleanup` 的指针的特殊情况下，`runtime.AddCleanup` 会引发 panic，作为对调用者的信号，表明他们不应以与终结器相同的方式使用清理。）

Second, when weak pointers are used as map keys, the weakly referenced object must not be reachable from the corresponding map value, otherwise the object will continue to remain live. This may seem obvious when deep inside of a blog post about weak pointers, but it’s an easy subtlety to miss. This problem inspired the entire concept of an [ephemeron](https://en.wikipedia.org/wiki/Ephemeron) to resolve it, which is a potential future direction.

​	其次，当弱指针用作映射键时，弱引用的对象不得从相应的映射值中可达，否则该对象将继续保持存活。在深入探讨弱指针的博客文章中这可能看似显而易见，但这是一个容易忽略的微妙之处。这个问题启发了 [ephemeron](https://en.wikipedia.org/wiki/Ephemeron) 的整个概念来解决它，这是未来的一个潜在方向。

Thirdly, a common pattern with cleanups is that a wrapper object is needed, like we see here with our `MemoryMappedFile` example. In this particular case, you could imagine the garbage collector directly tracking the mapped memory region and passing around the inner `[]byte`. Such functionality is possible future work, and an API for it has been recently [proposed](https://go.dev/issue/70224).

​	第三，清理的一个常见模式是需要一个包装对象，就像我们在 `MemoryMappedFile` 示例中看到的那样。在这个特定情况下，你可以想象垃圾回收器直接跟踪映射的内存区域并传递内部的 `[]byte`。这样的功能是未来可能的工作，一个相关的 API 最近已被 [提议](https://go.dev/issue/70224)。

Lastly, both weak pointers and cleanups are inherently non-deterministic, their behavior depending intimately on the design and dynamics of the garbage collector. The documentation for cleanups even permits the garbage collector never to run cleanups at all. Effectively testing code that uses them can be tricky, but [it is possible](https://go.dev/doc/gc-guide#Testing_object_death).

​	最后，弱指针和清理本质上是非确定性的，它们的行为密切依赖于垃圾回收器的设计和动态。清理的文档甚至允许垃圾回收器完全不运行清理。有效测试使用它们的代码可能很棘手，但 [这是可能的](https://go.dev/doc/gc-guide#Testing_object_death)。

## Why now? 为什么现在？

Weak pointers have been brought up as a feature for Go since nearly the beginning, but for years were not prioritized by the Go team. One reason for that is that they are subtle, and the design space of weak pointers is a minefield of decisions that can make them even harder to use. Another is that weak pointers are a niche tool, while simultaneously adding complexity to the language. We already had experience with how painful `SetFinalizer` could be to use. But there are some useful programs that are not expressible without them, and the `unique` package and the reasons for its existence really emphasized that.

​	弱指针几乎从一开始就被提出作为 Go 的一个功能，但多年来未被 Go 团队优先考虑。一个原因是它们很微妙，弱指针的设计空间充满了决策雷区，可能使它们更难使用。另一个原因是弱指针是一个小众工具，同时增加了语言的复杂性。我们已经体验到使用 `SetFinalizer` 有多痛苦。但有些有用的程序没有它们无法表达，`unique` 包及其存在的理由真正强调了这一点。

With generics, the hindsight of finalizers, and insights from all the great work since done by teams in other languages like C# and Java, the designs for weak pointers and cleanups came together quickly. The desire to use weak pointers with finalizers raised additional questions, and so the design for `runtime.AddCleanup` quickly came together as well.

​	有了泛型、终结器的后见之明，以及自其他语言（如 C# 和 Java）团队所做的伟大工作以来获得的见解，弱指针和清理的设计很快就成型了。希望将弱指针与终结器一起使用提出了额外的问题，因此 `runtime.AddCleanup` 的设计也很快成型。

## Acknowledgements 致谢

I want to thank everyone in the community who contributed feedback on the proposal issues and filed bugs when the features became available. I also want to thank David Chase for thoroughly thinking through weak pointer semantics with me, and I want to thank him, Russ Cox, and Austin Clements for their help with the design of `runtime.AddCleanup`. I want to thank Carlos Amedee for his work on getting `runtime.AddCleanup` implemented, polished, landed for Go 1.24. And finally I want to thank Carlos Amedee and Ian Lance Taylor for their work replacing `runtime.SetFinalizer` with `runtime.AddCleanup` throughout the standard library for Go 1.25.

​	我要感谢社区中所有在提案问题上提供反馈并在功能可用时提交错误的人。我还要感谢 David Chase 与我一起彻底思考弱指针语义，还要感谢他、Russ Cox 和 Austin Clements 在 `runtime.AddCleanup` 设计上的帮助。我要感谢 Carlos Amedee 为实现、完善并将 `runtime.AddCleanup` 纳入 Go 1.24 所做的工作。最后，我要感谢 Carlos Amedee 和 Ian Lance Taylor 为 Go 1.25 在整个标准库中用 `runtime.AddCleanup` 替换 `runtime.SetFinalizer` 所做的工作。