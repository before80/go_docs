+++
title = "新的 unique 包"
date = 2024-09-06T12:19:59+08:00
weight = 900
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# New unique package - 新的 unique 包

Michael Knyszek
27 August 2024

作者：Michael Knyszek
日期：2024 年 8 月 27 日

The standard library of Go 1.23 now includes the [new `unique` package](https://pkg.go.dev/unique). The purpose behind this package is to enable the canonicalization of comparable values. In other words, this package lets you deduplicate values so that they point to a single, canonical, unique copy, while efficiently managing the canonical copies under the hood. You might be familiar with this concept already, called [“interning”](https://en.wikipedia.org/wiki/Interning_(computer_science)). Let’s dive in to see how it works, and why it’s useful.

​	Go 1.23 的标准库现在包含了 [新的 `unique` 包](https://pkg.go.dev/unique)。该包的目的在于实现可比较值的标准化（"interning"）。换句话说，这个包允许你对值去重，使它们指向一个唯一的、标准化的副本，同时在底层高效地管理这些标准副本。你可能已经熟悉这个概念，这被称为 [“interning”](https://en.wikipedia.org/wiki/Interning_(computer_science))。让我们来看看它是如何工作的，以及它为何有用。

## 一个简单的 interning 实现 A simple implementation of interning

At a high level, interning is very simple. Take the code sample below, which deduplicates strings using just a regular map.

​	从高层来看，interning 非常简单。以下代码示例使用常规的 map 对字符串进行去重。

```go
var internPool map[string]string

// Intern returns a string that is equal to s but that may share storage with
// a string previously passed to Intern.
// Intern 返回一个与 s 相等的字符串，但该字符串可能与之前传递给 Intern 的字符串共享存储。
func Intern(s string) string {
    pooled, ok := internPool[s]
    if !ok {
        // Clone the string in case it's part of some much bigger string.
        // This should be rare, if interning is being used well.
         // 克隆字符串，以防它是某个更大字符串的一部分。
        // 如果合理使用 interning，这种情况应该很少发生。
        pooled = strings.Clone(s)
        internPool[pooled] = pooled
    }
    return pooled
}
```

This is useful for when you’re constructing a lot of strings that are likely to be duplicates, like when parsing a text format.

​	当你在构造可能会有很多重复的字符串时，这段代码非常有用，比如在解析文本格式时。

This implementation is super simple and works well enough for some cases, but it has a few problems:

​	这个实现非常简单，并且在某些情况下工作得很好，但它有几个问题：

- It never removes strings from the pool.
- 它永远不会从池中移除字符串。
- It cannot be safely used by multiple goroutines concurrently.
- 它不能被多个 goroutine 并发安全地使用。
- It only works with strings, even though the idea is quite general.
- 它只能处理字符串，尽管这个想法是通用的。

There’s also a missed opportunity in this implementation, and it’s subtle. Under the hood, [strings are immutable structures consisting of a pointer and a length](https://go.dev/blog/slices). When comparing two strings, if the pointers are not equal, then we must compare their contents to determine equality. But if we know that two strings are canonicalized, then it *is* sufficient to just check their pointers.

​	这个实现还有一个细微的缺点。底层的 [字符串是由指针和长度组成的不可变结构](https://go.dev/blog/slices)。当比较两个字符串时，如果它们的指针不相等，那么必须比较它们的内容来确定是否相等。但如果我们知道两个字符串是标准化的，那么仅仅检查它们的指针就足够了。

## 引入 `unique` 包 Enter the `unique` package

The new `unique` package introduces a function similar to `Intern` called [`Make`](https://pkg.go.dev/unique#Make).

​	新的 `unique` 包引入了一个类似于 `Intern` 的函数，叫做 [`Make`](https://pkg.go.dev/unique#Make)。

It works about the same way as `Intern`. Internally there’s also a global map ([a fast generic concurrent map](https://pkg.go.dev/internal/concurrent@go1.23.0)) and `Make` looks up the provided value in that map. But it also differs from `Intern` in two important ways. Firstly, it accepts values of any comparable type. And secondly, it returns a wrapper value, a [`Handle[T\]`](https://pkg.go.dev/unique#Handle), from which the canonical value can be retrieved.

​	它的工作方式与 `Intern` 相似。内部也有一个全局 map（[一个快速的泛型并发 map](https://pkg.go.dev/internal/concurrent@go1.23.0)），`Make` 会在该 map 中查找提供的值。但它与 `Intern` 有两个重要的不同点。首先，它接受任何可比较类型的值。其次，它返回一个包装值 [`Handle[T\]`](https://pkg.go.dev/unique#Handle)，可以通过它来获取标准化的值。

This `Handle[T]` is key to the design. A `Handle[T]` has the property that two `Handle[T]` values are equal if and only if the values used to create them are equal. What’s more, the comparison of two `Handle[T]` values is cheap: it comes down to a pointer comparison. Compared to comparing two long strings, that’s an order of magnitude cheaper!

​	这个 `Handle[T]` 是设计的关键。`Handle[T]` 具有这样的属性：只有在创建它们的值相等时，两个 `Handle[T]` 才会相等。而且，比较两个 `Handle[T]` 的代价很低：它仅需进行指针比较。与比较两个长字符串相比，这种比较的效率要高一个数量级！

So far, this is nothing you can’t do in ordinary Go code.

​	到目前为止，这些都是你可以在普通 Go 代码中实现的功能。

But `Handle[T]` also has a second purpose: as long as a `Handle[T]` exists for a value, the map will retain the canonical copy of the value. Once all `Handle[T]` values that map to a specific value are gone, the package marks that internal map entry as deletable, to be reclaimed in the near future. This sets a clear policy for when to remove entries from the map: when the canonical entries are no longer being used, then the garbage collector is free to clean them up.

​	但 `Handle[T]` 还有一个额外的作用：只要 `Handle[T]` 存在于某个值，那么 map 就会保留该值的标准副本。一旦所有映射到特定值的 `Handle[T]` 都被销毁，该包就会将该内部 map 条目标记为可删除，并在不久的将来进行回收。这为何时从 map 中移除条目制定了明确的策略：当标准化的条目不再被使用时，垃圾收集器可以自由地清理它们。

If you’ve used Lisp before, this may all sound quite familiar to you. Lisp [symbols](https://en.wikipedia.org/wiki/Symbol_(programming)) are interned strings, but not strings themselves, and all symbols’ string values are guaranteed to be in the same pool. This relationship between symbols and strings parallels the relationship between `Handle[string]` and `string`.

​	如果你曾经使用过 Lisp，这一切可能听起来很熟悉。Lisp 的 [symbols](https://en.wikipedia.org/wiki/Symbol_(programming)) 是 interned 字符串，但不是字符串本身，所有 symbols 的字符串值都保证在同一个池中。这种 symbols 和字符串的关系类似于 `Handle[string]` 和 `string` 之间的关系。

## 一个真实的例子 A real-world example

So, how might one use `unique.Make`? Look no further than the `net/netip` package in the standard library, which interns values of type `addrDetail`, part of the [`netip.Addr`](https://pkg.go.dev/net/netip#Addr) structure.

​	那么，如何使用 `unique.Make` 呢？可以看看标准库中的 `net/netip` 包，它将 `addrDetail` 类型的值进行了 intern，该类型是 [`netip.Addr`](https://pkg.go.dev/net/netip#Addr) 结构的一部分。

Below is an abridged version of the actual code from `net/netip` that uses `unique`.

​	下面是实际 `net/netip` 代码的简化版本，使用了 `unique`。

```go
// Addr represents an IPv4 or IPv6 address (with or without a scoped
// addressing zone), similar to net.IP or net.IPAddr.
// Addr 表示一个 IPv4 或 IPv6 地址（可能包含或不包含作用域地址区域），类似于 net.IP 或 net.IPAddr。
type Addr struct {
    // Other irrelevant unexported fields...
    // 其他不相关的未导出字段...

    // Details about the address, wrapped up together and canonicalized.
    // 地址的详细信息，被包装在一起并标准化。
    z unique.Handle[addrDetail]
}

// addrDetail indicates whether the address is IPv4 or IPv6, and if IPv6,
// specifies the zone name for the address.
// addrDetail 表示地址是 IPv4 还是 IPv6，如果是 IPv6，还指定了该地址的区域名称。
type addrDetail struct {
    isV6   bool   // IPv4 is false, IPv6 is true. IPv4 为 false，IPv6 为 true。
    zoneV6 string // May be != "" if IsV6 is true.  如果 IsV6 为 true，可能不为空。
}

var z6noz = unique.Make(addrDetail{isV6: true})

// WithZone returns an IP that's the same as ip but with the provided
// zone. If zone is empty, the zone is removed. If ip is an IPv4
// address, WithZone is a no-op and returns ip unchanged.
// WithZone 返回一个与 ip 相同但带有提供的区域的 IP。
// 如果 zone 为空，则删除该区域。如果 ip 是 IPv4 地址，WithZone 是无操作，并返回未更改的 ip。
func (ip Addr) WithZone(zone string) Addr {
    if !ip.Is6() {
        return ip
    }
    if zone == "" {
        ip.z = z6noz
        return ip
    }
    ip.z = unique.Make(addrDetail{isV6: true, zoneV6: zone})
    return ip
}
```

Since many IP addresses are likely to use the same zone and this zone is part of their identity, it makes a lot of sense to canonicalize them. The deduplication of zones reduces the average memory footprint of each `netip.Addr`, while the fact that they’re canonicalized means `netip.Addr` values are more efficient to compare, since comparing zone names becomes a simple pointer comparison.

​	由于许多 IP 地址可能使用相同的区域，并且该区域是其身份的一部分，因此将它们标准化是很有意义的。区域的去重减少了每个 `netip.Addr` 的平均内存占用，而由于它们是标准化的，`netip.Addr` 值的比较效率更高，因为区域名称的比较变成了简单的指针比较。

## 关于字符串 interning 的附注 A footnote about interning strings

While the `unique` package is useful, `Make` is admittedly not quite like `Intern` for strings, since the `Handle[T]` is required to keep a string from being deleted from the internal map. This means you need to modify your code to retain handles as well as strings.

​	尽管 `unique` 包很有用，但 `Make` 与 `Intern` 对字符串的处理略有不同，因为 `Handle[T]` 是保持字符串不被从内部 map 中删除的必要条件。这意味着你需要修改代码来同时保留 handle 和字符串。

But strings are special in that, although they behave like values, they actually contain pointers under the hood, as we mentioned earlier. This means that we could potentially canonicalize just the underlying storage of the string, hiding the details of a `Handle[T]` inside the string itself. So, there is still a place in the future for what I’ll call *transparent string interning*, in which strings can be interned without the `Handle[T]` type, similar to the `Intern` function but with semantics more closely resembling `Make`.

​	但字符串是特殊的，尽管它们表现得像值，但实际上在底层包含了指针，正如我们之前提到的。这意味着我们可以潜在地仅对字符串的底层存储进行标准化，将 `Handle[T]` 的细节隐藏在字符串内部。因此，未来仍然有实现 *透明字符串 interning* 的空间，在这种情况下，字符串可以像 `Intern` 函数一样被 intern，但语义更接近 `Make`。

In the meantime, `unique.Make("my string").Value()` is one possible workaround. Even though failing to retain the handle will allow the string to be deleted from `unique`’s internal map, map entries are not deleted immediately. In practice, entries will not be deleted until at least the next garbage collection completes, so this workaround still allows for some degree of deduplication in the periods between collections.

​	在此期间，`unique.Make("my string").Value()` 是一种可能的替代方案。即使没有保留 handle 会导致字符串从 `unique` 的内部 map 中删除，但 map 条目不会立即被删除。实际上，条目不会在下一次垃圾收集完成之前被删除，因此在垃圾收集之间的时间段内，这种替代方案仍然允许一定程度的去重。

## 历史背景与未来展望 Some history, and looking toward the future

The truth is that the `net/netip` package actually interned zone strings since it was first introduced. The interning package it used was an internal copy of the [go4.org/intern](https://pkg.go.dev/go4.org/intern) package. Like the `unique` package, it has a `Value` type (which looks a lot like a `Handle[T]`, pre-generics), has the notable property that entries in the internal map are removed once their handles are no longer referenced.

​	事实上，`net/netip` 包从首次引入时就已经对区域字符串进行了 intern。它使用的 intern 包是 [go4.org/intern](https://pkg.go.dev/go4.org/intern) 包的内部副本。与 `unique` 包类似，它有一个 `Value` 类型（在泛型之前，它看起来很像 `Handle[T]`），并具有这样一个显著的特点：一旦 handle 不再被引用，内部 map 中的条目将被移除。

But to achieve this behavior, it has to do some unsafe things. In particular, it makes some assumptions about the garbage collector’s behavior to implement [*weak pointers*](https://en.wikipedia.org/wiki/Weak_reference) outside the runtime. A weak pointer is a pointer that doesn’t prevent the garbage collector from reclaiming a variable; when this happens, the pointer automatically becomes nil. As it happens, weak pointers are *also* the core abstraction underlying the `unique` package.

​	但为了实现这一行为，它必须做一些不安全的事情。特别是，它需要对垃圾收集器的行为做出一些假设，以在运行时之外实现 [*弱指针*](https://en.wikipedia.org/wiki/Weak_reference)。弱指针是一种不阻止垃圾收集器回收变量的指针；当这种情况发生时，指针会自动变为 nil。事实上，弱指针也是 `unique` 包的核心抽象。

That’s right: while implementing the `unique` package, we added proper weak pointer support to the garbage collector. And after stepping through the minefield of regrettable design decisions that accompany weak pointers (like, should weak pointers track [object resurrection](https://en.wikipedia.org/wiki/Object_resurrection)? No!), we were astonished by how simple and straightforward all of it turned out to be. Astonished enough that weak pointers are now a [public proposal](https://go.dev/issue/67552).

​	没错：在实现 `unique` 包的过程中，我们为垃圾收集器添加了适当的弱指针支持。经过对弱指针设计决策所带来的问题的深思熟虑（例如，弱指针是否应该跟踪 [对象复活](https://en.wikipedia.org/wiki/Object_resurrection)？答案是否定的！），我们惊讶地发现，一切变得如此简单和直接。这个过程中的惊讶程度足以让弱指针成为 [公共提案](https://go.dev/issue/67552)。

This work also led us to reexamine finalizers, resulting in another proposal for an easier-to-use and more efficient [replacement for finalizers](https://go.dev/issue/67535). With [a hash function for comparable values](https://go.dev/issue/54670) on the way as well, the future of [building memory-efficient caches](https://go.dev/issue/67552#issuecomment-2200755798) in Go is bright!

​	这一工作还促使我们重新审视终结器（finalizer），提出了一个更易于使用且效率更高的 [终结器替代方案](https://go.dev/issue/67535) 的提案。随着 [可比较值的哈希函数](https://go.dev/issue/54670) 的推出，Go 中构建内存高效缓存的未来一片光明！
