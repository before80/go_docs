+++
title = "go maps in action"
weight = 17
date = 2023-05-18T17:03:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# Go maps in action

https://go.dev/blog/maps

Andrew Gerrand
6 February 2013

## Introduction 简介

One of the most useful data structures in computer science is the hash table. Many hash table implementations exist with varying properties, but in general they offer fast lookups, adds, and deletes. Go provides a built-in map type that implements a hash table.

计算机科学中最有用的数据结构之一是哈希表。许多哈希表的实现都有不同的属性，但总的来说，它们提供了快速查找、添加和删除的功能。Go提供了一个内置的map类型来实现哈希表。

## Declaration and initialization 声明和初始化

A Go map type looks like this:

Go的映射类型看起来像这样：

```go linenums="1"
map[KeyType]ValueType
```

where `KeyType` may be any type that is [comparable](https://go.dev/ref/spec#Comparison_operators) (more on this later), and `ValueType` may be any type at all, including another map!

其中 KeyType 可以是任何可比较的类型（后面会有更多介绍），ValueType 可以是任何类型，包括另一个 map!

This variable `m` is a map of string keys to int values:

这个变量m是一个字符串键到int值的映射：

```go linenums="1"
var m map[string]int
```

Map types are reference types, like pointers or slices, and so the value of `m` above is `nil`; it doesn’t point to an initialized map. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don’t do that. To initialize a map, use the built in `make` function:

地图类型是引用类型，就像指针或片断一样，所以上面m的值是nil；它没有指向一个初始化的地图。一个nil地图在读取时表现得像一个空地图，但试图向一个nil地图写入会引起运行时的恐慌；不要这样做。要初始化一个地图，请使用内置的make函数。

```go linenums="1"
m = make(map[string]int)
```

The `make` function allocates and initializes a hash map data structure and returns a map value that points to it. The specifics of that data structure are an implementation detail of the runtime and are not specified by the language itself. In this article we will focus on the *use* of maps, not their implementation.

make函数分配和初始化了一个哈希地图数据结构，并返回一个指向它的地图值。该数据结构的具体细节是运行时的一个实现细节，而不是由语言本身指定。在这篇文章中，我们将重点讨论地图的使用，而不是它们的实现。

## Working with maps 使用地图工作

Go provides a familiar syntax for working with maps. This statement sets the key `"route"` to the value `66`:

Go提供了一种熟悉的语法来处理地图。此语句将键 "route "设置为值66：

```go linenums="1"
m["route"] = 66
```

This statement retrieves the value stored under the key `"route"` and assigns it to a new variable i:

这条语句检索存储在键 "route "下的值，并将其分配给一个新的变量i：

```go linenums="1"
i := m["route"]
```

If the requested key doesn’t exist, we get the value type’s *zero value*. In this case the value type is `int`, so the zero value is `0`:

如果请求的键不存在，我们会得到值类型的零值。在这个例子中，值类型是int，所以零值是0：

```go linenums="1"
j := m["root"]
// j == 0
```

The built in `len` function returns on the number of items in a map:

内置的len函数会返回地图中的项数：

```go linenums="1"
n := len(m)
```

The built in `delete` function removes an entry from the map:

内置的delete函数可以从地图中删除一个条目：

```go linenums="1"
delete(m, "route")
```

The `delete` function doesn’t return anything, and will do nothing if the specified key doesn’t exist.

`delete`函数不返回任何东西，如果指定的键不存在，也不会做任何事情。

A two-value assignment tests for the existence of a key:

```go linenums="1"
i, ok := m["route"]
```

In this statement, the first value (`i`) is assigned the value stored under the key `"route"`. If that key doesn’t exist, `i` is the value type’s zero value (`0`). The second value (`ok`) is a `bool` that is `true` if the key exists in the map, and `false` if not.

在这个语句中，第一个值（i）被分配给存储在键 "route "下的值。如果该键不存在，i就是该值类型的零值（0）。第二个值（ok）是一个bool，如果该键在地图中存在，则为真，如果不存在，则为假。

To test for a key without retrieving the value, use an underscore in place of the first value:

要测试一个键而不检索其值，可以用下划线代替第一个值：

```go linenums="1"
_, ok := m["route"]
```

To iterate over the contents of a map, use the `range` keyword:

要遍历一个地图的内容，请使用`range`关键字：

```go linenums="1"
for key, value := range m {
    fmt.Println("Key:", key, "Value:", value)
}
```

To initialize a map with some data, use a map literal:

要用一些数据来初始化一个地图，请使用一个地图字面：

```go linenums="1"
commits := map[string]int{
    "rsc": 3711,
    "r":   2138,
    "gri": 1908,
    "adg": 912,
}
```

The same syntax may be used to initialize an empty map, which is functionally identical to using the `make` function:

同样的语法可以用来初始化一个空的地图，这在功能上与使用make函数相同：

```go linenums="1"
m = map[string]int{}
```

## Exploiting zero values 利用零值

It can be convenient that a map retrieval yields a zero value when the key is not present.

当键不存在时，地图检索会产生一个零值，这很方便。

For instance, a map of boolean values can be used as a set-like data structure (recall that the zero value for the boolean type is false). This example traverses a linked list of `Nodes` and prints their values. It uses a map of `Node` pointers to detect cycles in the list.

例如，一个布尔值的地图可以被用作一个类似于集合的数据结构（记得布尔类型的零值是false）。这个例子遍历了一个Nodes的链接列表，并打印了它们的值。它使用一个Node指针地图来检测列表中的循环。

```go linenums="1"
    type Node struct {
        Next  *Node
        Value interface{}
    }
    var first *Node

    visited := make(map[*Node]bool)
    for n := first; n != nil; n = n.Next {
        if visited[n] {
            fmt.Println("cycle detected")
            break
        }
        visited[n] = true
        fmt.Println(n.Value)
    }
```

The expression `visited[n]` is `true` if `n` has been visited, or `false` if `n` is not present. There’s no need to use the two-value form to test for the presence of `n` in the map; the zero value default does it for us.

如果n已经被访问，表达式visited[n]为真，如果n不存在，则为假。没有必要使用双值形式来测试地图中是否存在n；零值默认为我们做了这件事。

Another instance of helpful zero values is a map of slices. Appending to a nil slice just allocates a new slice, so it’s a one-liner to append a value to a map of slices; there’s no need to check if the key exists. In the following example, the slice people is populated with `Person` values. Each `Person` has a `Name` and a slice of Likes. The example creates a map to associate each like with a slice of people that like it.

另一个有用的零值的例子是一个分片的地图。向一个无值的片子追加只是分配一个新的片子，所以向一个片子的地图追加一个值是一个单行代码；不需要检查键是否存在。在下面的例子中，片断people被填充了Person的值。每个人都有一个名字和一个 "喜欢 "的片断。这个例子创建了一个地图，将每个 "喜欢 "与喜欢它的人的片断联系起来。

```go linenums="1"
    type Person struct {
        Name  string
        Likes []string
    }
    var people []*Person

    likes := make(map[string][]*Person)
    for _, p := range people {
        for _, l := range p.Likes {
            likes[l] = append(likes[l], p)
        }
    }
```

To print a list of people who like cheese:

要打印一个喜欢奶酪的人的列表：

```go linenums="1"
    for _, p := range likes["cheese"] {
        fmt.Println(p.Name, "likes cheese.")
    }
```

To print the number of people who like bacon:

要打印喜欢培根的人的数量：

```go linenums="1"
    fmt.Println(len(likes["bacon"]), "people like bacon.")
```

Note that since both range and len treat a nil slice as a zero-length slice, these last two examples will work even if nobody likes cheese or bacon (however unlikely that may be).

请注意，由于range和len都将nil片断视为零长度的片断，所以即使没有人喜欢奶酪或培根（无论多么不可能），最后两个例子也会有效。

## Key types 键类型

As mentioned earlier, map keys may be of any type that is comparable. The [language spec](https://go.dev/ref/spec#Comparison_operators) defines this precisely, but in short, comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types. Notably absent from the list are slices, maps, and functions; these types cannot be compared using `==`, and may not be used as map keys.

如前所述，地图键可以是任何可比较的类型。语言规范对此有精确的定义，但简而言之，可比较的类型是布尔型、数字型、字符串型、指针型、通道型和接口型，以及只包含这些类型的结构或数组。值得注意的是，列表中没有切片、地图和函数；这些类型不能使用==进行比较，也不能作为地图键使用。

It’s obvious that strings, ints, and other basic types should be available as map keys, but perhaps unexpected are struct keys. Struct can be used to key data by multiple dimensions. For example, this map of maps could be used to tally web page hits by country:

很明显，字符串、ints和其他基本类型应该可以作为地图键，但也许意想不到的是结构键。结构可以用来按多个维度对数据进行键控。例如，这个地图可以用来按国家统计网页点击率：

```go linenums="1"
hits := make(map[string]map[string]int)
```

This is map of string to (map of `string` to `int`). Each key of the outer map is the path to a web page with its own inner map. Each inner map key is a two-letter country code. This expression retrieves the number of times an Australian has loaded the documentation page:

这是将字符串映射为（map of string to int）。外层地图的每个键是一个网页的路径，有自己的内部地图。每个内层地图的键是一个两个字母的国家代码。这个表达式检索了一个澳大利亚人加载文档页面的次数。

```go linenums="1"
n := hits["/doc/"]["au"]
```

Unfortunately, this approach becomes unwieldy when adding data, as for any given outer key you must check if the inner map exists, and create it if needed:

不幸的是，这种方法在添加数据时变得不方便，因为对于任何给定的外键，你必须检查内部地图是否存在，如果需要的话，还要创建它：

```go linenums="1"
func add(m map[string]map[string]int, path, country string) {
    mm, ok := m[path]
    if !ok {
        mm = make(map[string]int)
        m[path] = mm
    }
    mm[country]++
}
add(hits, "/doc/", "au")
```

On the other hand, a design that uses a single map with a struct key does away with all that complexity:

另一方面，一个使用单一地图与结构键的设计可以消除所有这些复杂性：

```go linenums="1"
type Key struct {
    Path, Country string
}
hits := make(map[Key]int)
```

When a Vietnamese person visits the home page, incrementing (and possibly creating) the appropriate counter is a one-liner:

当一个越南人访问主页时，递增（以及可能创建）适当的计数器是一个单行代码：

```go linenums="1"
hits[Key{"/", "vn"}]++
```

And it’s similarly straightforward to see how many Swiss people have read the spec:

同样，也可以直接查看有多少瑞士人阅读过本规范：

```go linenums="1"
n := hits[Key{"/ref/spec", "ch"}]
```

## Concurrency 并发性

[Maps are not safe for concurrent use](https://go.dev/doc/faq#atomic_maps): it’s not defined what happens when you read and write to them simultaneously. If you need to read from and write to a map from concurrently executing goroutines, the accesses must be mediated by some kind of synchronization mechanism. One common way to protect maps is with [sync.RWMutex](https://go.dev/pkg/sync/#RWMutex).

地图对于并发使用是不安全的：它没有定义当你同时读和写它们时会发生什么。如果你需要从同时执行的goroutine中读取和写入一个地图，那么这些访问必须由某种同步机制来调解。一种常见的保护地图的方法是使用sync.RWMutex。

This statement declares a `counter` variable that is an anonymous struct containing a map and an embedded `sync.RWMutex`.

这个语句声明了一个计数器变量，它是一个匿名结构，包含一个地图和一个嵌入式的sync.RWMutex。

```go linenums="1"
var counter = struct{
    sync.RWMutex
    m map[string]int
}{m: make(map[string]int)}
```

To read from the counter, take the read lock:

要从计数器中读出数据，需要取得读锁：

```go linenums="1"
counter.RLock()
n := counter.m["some_key"]
counter.RUnlock()
fmt.Println("some_key:", n)
```

To write to the counter, take the write lock:

要对计数器进行写操作，需要取得写锁：

```go linenums="1"
counter.Lock()
counter.m["some_key"]++
counter.Unlock()
```

## Iteration order 迭代顺序

When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed to be the same from one iteration to the next. If you require a stable iteration order you must maintain a separate data structure that specifies that order. This example uses a separate sorted slice of keys to print a `map[int]string` in key order:

当用一个范围循环在地图上迭代时，迭代顺序没有被指定，也不能保证每次迭代都是一样的。如果你需要一个稳定的迭代顺序，你必须维护一个单独的数据结构来指定这个顺序。这个例子使用一个单独的键的排序片，按照键的顺序打印map[int]字符串：

```go linenums="1"
import "sort"

var m map[int]string
var keys []int
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println("Key:", k, "Value:", m[k])
}
```

