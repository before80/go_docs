+++
title = "iter"
date = 2024-09-05T11:11:18+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/iter@go1.23.0](https://pkg.go.dev/iter@go1.23.0)

> 注意
>
> ​	从go1.23.0开始才有该包。

Package iter provides basic definitions and operations related to iterators over sequences.

​	`iter` 包提供了与序列迭代器相关的基本定义和操作。

## 迭代器 Iterators 

An iterator is a function that passes successive elements of a sequence to a callback function, conventionally named yield. The function stops either when the sequence is finished or when yield returns false, indicating to stop the iteration early. This package defines [Seq](https://pkg.go.dev/iter@go1.23.0#Seq) and [Seq2](https://pkg.go.dev/iter@go1.23.0#Seq2) (pronounced like seek—the first syllable of sequence) as shorthands for iterators that pass 1 or 2 values per sequence element to yield:

​	迭代器是一个函数，它将序列中的后续元素传递给一个回调函数，通常命名为 `yield`。当序列结束或者 `yield` 返回 `false` 时，函数会停止迭代，表示提前停止迭代。此包定义了 [Seq](https://pkg.go.dev/iter@go1.23.0#Seq) 和 [Seq2](https://pkg.go.dev/iter@go1.23.0#Seq2)（读音类似于 sequence 的前两个音节 “seek”），它们分别代表传递 1 个或 2 个值给 `yield` 的迭代器：

```
type (
	Seq[V any]     func(yield func(V) bool)
	Seq2[K, V any] func(yield func(K, V) bool)
)
```

Seq2 represents a sequence of paired values, conventionally key-value or index-value pairs.

​	`Seq2` 表示一个成对值的序列，通常是键值对或索引值对。

Yield returns true if the iterator should continue with the next element in the sequence, false if it should stop.

​	`yield` 返回 `true` 时，迭代器会继续传递序列中的下一个元素；返回 `false` 时则停止迭代。

Iterator functions are most often called by a range loop, as in:

​	迭代器函数通常在 `range` 循环中调用，如下所示：

```
func PrintAll[V any](seq iter.Seq[V]) {
	for v := range seq {
		fmt.Println(v)
	}
}
```

## 命名约定 Naming Conventions 

Iterator functions and methods are named for the sequence being walked:

​	迭代器函数和方法通常根据所迭代的序列命名：

```
// All returns an iterator over all elements in s.
// All 返回序列中所有元素的迭代器。
func (s *Set[V]) All() iter.Seq[V]
```

The iterator method on a collection type is conventionally named All, because it iterates a sequence of all the values in the collection.

​	集合类型的迭代器方法通常命名为 `All`，因为它迭代集合中的所有值。

For a type containing multiple possible sequences, the iterator's name can indicate which sequence is being provided:

​	对于包含多个可能序列的类型，迭代器的名称可以指示提供的序列：

```
// Cities returns an iterator over the major cities in the country.
// Cities 返回国家中主要城市的迭代器。
func (c *Country) Cities() iter.Seq[*City]

// Languages returns an iterator over the official spoken languages of the country.
// Languages 返回国家的官方语言的迭代器。
func (c *Country) Languages() iter.Seq[string]
```

If an iterator requires additional configuration, the constructor function can take additional configuration arguments:

​	如果迭代器需要额外的配置，构造函数可以接受额外的配置参数：

```
// Scan returns an iterator over key-value pairs with min ≤ key ≤ max.
// Scan 返回键值对的迭代器，满足 min ≤ key ≤ max。
func (m *Map[K, V]) Scan(min, max K) iter.Seq2[K, V]

// Split returns an iterator over the (possibly-empty) substrings of s
// separated by sep.
// Split 返回由 sep 分隔的 s 的（可能为空的）子字符串的迭代器。
func Split(s, sep string) iter.Seq[string]
```

When there are multiple possible iteration orders, the method name may indicate that order:

​	当有多个可能的迭代顺序时，方法名可能会指示该顺序：

```
// All returns an iterator over the list from head to tail.
// All 返回从头到尾迭代列表的迭代器。
func (l *List[V]) All() iter.Seq[V]

// Backward returns an iterator over the list from tail to head.
// Backward 返回从尾到头迭代列表的迭代器。
func (l *List[V]) Backward() iter.Seq[V]

// Preorder returns an iterator over all nodes of the syntax tree
// beneath (and including) the specified root, in depth-first preorder,
// visiting a parent node before its children.
// Preorder 返回遍历语法树中指定根节点及其子节点的迭代器，深度优先先序遍历，先访问父节点，再访问子节点。
func Preorder(root Node) iter.Seq[Node]
```

## 单次使用迭代器 Single-Use Iterators 

Most iterators provide the ability to walk an entire sequence: when called, the iterator does any setup necessary to start the sequence, then calls yield on successive elements of the sequence, and then cleans up before returning. Calling the iterator again walks the sequence again.

​	大多数迭代器提供遍历整个序列的能力：当被调用时，迭代器会进行必要的设置以开始序列，然后依次调用 `yield` 传递序列元素，最后在返回之前进行清理。再次调用迭代器会再次遍历序列。

Some iterators break that convention, providing the ability to walk a sequence only once. These “single-use iterators” typically report values from a data stream that cannot be rewound to start over. Calling the iterator again after stopping early may continue the stream, but calling it again after the sequence is finished will yield no values at all. Doc comments for functions or methods that return single-use iterators should document this fact:

​	有些迭代器打破了这个约定，只能遍历序列一次。这些“单次使用的迭代器”通常从无法重新开始的数据流中报告值。提前停止迭代后再次调用迭代器可能会继续数据流，但在序列结束后再次调用它将不会产生任何值。返回单次使用迭代器的函数或方法的文档注释应该记录这一事实：

```
// Lines returns an iterator over lines read from r.
// It returns a single-use iterator.
// Lines 返回从 r 读取的行的迭代器。
// 它返回一个单次使用的迭代器。
func (r *Reader) Lines() iter.Seq[string]
```

## 拉取值 Pulling Values 

Functions and methods that accept or return iterators should use the standard [Seq](https://pkg.go.dev/iter@go1.23.0#Seq) or [Seq2](https://pkg.go.dev/iter@go1.23.0#Seq2) types, to ensure compatibility with range loops and other iterator adapters. The standard iterators can be thought of as “push iterators”, which push values to the yield function.

​	接受或返回迭代器的函数和方法应该使用标准的 [Seq](https://pkg.go.dev/iter@go1.23.0#Seq) 或 [Seq2](https://pkg.go.dev/iter@go1.23.0#Seq2) 类型，以确保与 `range` 循环和其他迭代器适配器的兼容性。标准迭代器可以视为“推迭代器”，将值推送到 `yield` 函数。

Sometimes a range loop is not the most natural way to consume values of the sequence. In this case, [Pull](https://pkg.go.dev/iter@go1.23.0#Pull) converts a standard push iterator to a “pull iterator”, which can be called to pull one value at a time from the sequence. [Pull](https://pkg.go.dev/iter@go1.23.0#Pull) starts an iterator and returns a pair of functions—next and stop—which return the next value from the iterator and stop it, respectively.

​	有时 `range` 循环并不是消费序列值的最自然方式。在这种情况下， [Pull](https://pkg.go.dev/iter@go1.23.0#Pull) 可以将标准的推迭代器转换为“拉迭代器”，该迭代器可以通过调用一个函数来从序列中一次拉取一个值。[Pull](https://pkg.go.dev/iter@go1.23.0#Pull) 启动一个迭代器，并返回一对函数——`next` 和 `stop`，它们分别返回下一个值并停止迭代。

For example:

​	例如：

```
// Pairs returns an iterator over successive pairs of values from seq.
// Pairs 返回 seq 中连续值对的迭代器。
func Pairs[V any](seq iter.Seq[V]) iter.Seq2[V, V] {
	return func(yield func(V, V) bool) {
		next, stop := iter.Pull(seq)
		defer stop()
		for {
			v1, ok1 := next()
			if !ok1 {
				return
			}
			v2, ok2 := next()
			// If ok2 is false, v2 should be the
			// zero value; yield one last pair.
			// 如果 ok2 为 false，v2 应该是零值；
			// 传递最后一对值。
			if !yield(v1, v2) {
				return
			}
			if !ok2 {
				return
			}
		}
	}
}
```

If clients do not consume the sequence to completion, they must call stop, which allows the iterator function to finish and return. As shown in the example, the conventional way to ensure this is to use defer.

​	如果客户端没有完全消费序列，必须调用 `stop` 来结束迭代，这使得迭代器函数可以完成并返回。如示例中所示，确保这一点的常规方法是使用 `defer`。

## 标准库使用情况 Standard Library Usage 

A few packages in the standard library provide iterator-based APIs, most notably the [maps](https://pkg.go.dev/maps) and [slices](https://pkg.go.dev/slices) packages. For example, [maps.Keys](https://pkg.go.dev/maps#Keys) returns an iterator over the keys of a map, while [slices.Sorted](https://pkg.go.dev/slices#Sorted) collects the values of an iterator into a slice, sorts them, and returns the slice, so to iterate over the sorted keys of a map:

​	标准库中的一些包提供了基于迭代器的 API，最著名的有 [maps](https://pkg.go.dev/maps) 和 [slices](https://pkg.go.dev/slices) 包。例如， [maps.Keys](https://pkg.go.dev/maps#Keys) 返回一个映射键的迭代器，而 [slices.Sorted](https://pkg.go.dev/slices#Sorted) 收集迭代器的值到切片中，将其排序并返回切片，因此可以迭代排序后的映射键：

```
for _, key := range slices.Sorted(maps.Keys(m)) {
	...
}
```

## 变更 Mutation 

Iterators provide only the values of the sequence, not any direct way to modify it. If an iterator wishes to provide a mechanism for modifying a sequence during iteration, the usual approach is to define a position type with the extra operations and then provide an iterator over positions.

​	迭代器仅提供序列值，而不提供任何直接修改序列的方式。如果迭代器希望在迭代过程中提供修改序列的机制，通常的方法是定义一个带有额外操作的位置类型，然后提供该类型的位置的迭代器。

For example, a tree implementation might provide:

​	例如，树实现可能提供：

```
// Positions returns an iterator over positions in the sequence.
// Positions 返回序列中位置的迭代器。
func (t *Tree[V]) Positions() iter.Seq[*Pos]

// A Pos represents a position in the sequence.
// It is only valid during the yield call it is passed to.
// Pos 表示序列中的一个位置。
// 它仅在传递给 `yield` 调用期间有效。
type Pos[V any] struct { ... }

// Pos returns the value at the cursor.
// Pos 返回光标处的值。
func (p *Pos[V]) Value() V

// Delete deletes the value at this point in the iteration.
// Delete 删除迭代过程中当前位置的值。
func (p *Pos[V]) Delete()

// Set changes the value v at the cursor.
// Set 修改光标处的值为 v。
func (p *Pos[V]) Set(v V)
```

And then a client could delete boring values from the tree using:

​	然后客户端可以使用如下方式删除树中无聊的值：

```
for p := range t.Positions() {
	if boring(p.Value()) {
		p.Delete()
	}
}
```

## 常量

This section is empty.

## 变量 

This section is empty.

## 函数

### func Pull 

```
func Pull[V any](seq Seq[V]) (next func() (V, bool), stop func())
```

Pull converts the “push-style” iterator sequence seq into a “pull-style” iterator accessed by the two functions next and stop.

​	Pull 将“推送式”的迭代器序列 `seq` 转换为通过两个函数 `next` 和 `stop` 访问的“拉取式”迭代器。

Next returns the next value in the sequence and a boolean indicating whether the value is valid. When the sequence is over, next returns the zero V and false. It is valid to call next after reaching the end of the sequence or after calling stop. These calls will continue to return the zero V and false.

​	`next` 返回序列中的下一个值以及一个布尔值，指示该值是否有效。当序列结束时，`next` 返回零值 `V` 和 `false`。在到达序列末尾或调用 `stop` 之后调用 `next` 是合法的。这些调用将继续返回零值 `V` 和 `false`。

Stop ends the iteration. It must be called when the caller is no longer interested in next values and next has not yet signaled that the sequence is over (with a false boolean return). It is valid to call stop multiple times and when next has already returned false. Typically, callers should “defer stop()”.

​	`stop` 结束迭代。当调用者不再需要 `next` 的值并且 `next` 尚未发出序列结束的信号（通过返回 `false` 布尔值）时，必须调用 `stop`。可以多次调用 `stop`，以及当 `next` 已经返回 `false` 时调用。通常，调用者应使用 `defer stop()`。

It is an error to call next or stop from multiple goroutines simultaneously.

​	在多个 goroutine 中同时调用 `next` 或 `stop` 是错误的。

If the iterator panics during a call to next (or stop), then next (or stop) itself panics with the same value.

​	如果迭代器在调用 `next`（或 `stop`）期间发生恐慌，那么 `next`（或 `stop`）本身也会以相同的值发生恐慌。

### func Pull2 

```
func Pull2[K, V any](seq Seq2[K, V]) (next func() (K, V, bool), stop func())
```

Pull2 converts the “push-style” iterator sequence seq into a “pull-style” iterator accessed by the two functions next and stop.

​	Pull2 将“推送式”的迭代器序列 `seq` 转换为通过两个函数 `next` 和 `stop` 访问的“拉取式”迭代器。

Next returns the next pair in the sequence and a boolean indicating whether the pair is valid. When the sequence is over, next returns a pair of zero values and false. It is valid to call next after reaching the end of the sequence or after calling stop. These calls will continue to return a pair of zero values and false.

​	`next` 返回序列中的下一个键值对以及一个布尔值，指示该对是否有效。当序列结束时，`next` 返回一对零值和 `false`。在到达序列末尾或调用 `stop` 之后调用 `next` 是合法的。这些调用将继续返回一对零值和 `false`。

Stop ends the iteration. It must be called when the caller is no longer interested in next values and next has not yet signaled that the sequence is over (with a false boolean return). It is valid to call stop multiple times and when next has already returned false. Typically, callers should “defer stop()”.

​	`stop` 结束迭代。当调用者不再需要 `next` 的值并且 `next` 尚未发出序列结束的信号（通过返回 `false` 布尔值）时，必须调用 `stop`。可以多次调用 `stop`，以及当 `next` 已经返回 `false` 时调用。通常，调用者应使用 `defer stop()`。

It is an error to call next or stop from multiple goroutines simultaneously.

​	在多个 goroutine 中同时调用 `next` 或 `stop` 是错误的。

If the iterator panics during a call to next (or stop), then next (or stop) itself panics with the same value.

​	如果迭代器在调用 `next`（或 `stop`）期间发生恐慌，那么 `next`（或 `stop`）本身也会以相同的值发生恐慌。

## 类型

### type Seq 

```
type Seq[V any] func(yield func(V) bool)
```

Seq is an iterator over sequences of individual values. When called as seq(yield), seq calls yield(v) for each value v in the sequence, stopping early if yield returns false. See the [iter](https://pkg.go.dev/iter@go1.23.0) package documentation for more details.

​	Seq 是一个遍历单个值序列的迭代器。当被调用为 `seq(yield)` 时，`seq` 会为序列中的每个值 `v` 调用 `yield(v)`，如果 `yield` 返回 `false` 则提前停止。详情请参见 [iter](https://pkg.go.dev/iter@go1.23.0) 包文档。

### type Seq2 

```
type Seq2[K, V any] func(yield func(K, V) bool)
```

Seq2 is an iterator over sequences of pairs of values, most commonly key-value pairs. When called as seq(yield), seq calls yield(k, v) for each pair (k, v) in the sequence, stopping early if yield returns false. See the [iter](https://pkg.go.dev/iter@go1.23.0) package documentation for more details.

​	Seq2 是一个遍历键值对序列的迭代器，最常见的是键值对。当被调用为 `seq(yield)` 时，`seq` 会为序列中的每对键值 `(k, v)` 调用 `yield(k, v)`，如果 `yield` 返回 `false` 则提前停止。详情请参见 [iter](https://pkg.go.dev/iter@go1.23.0) 包文档。
