+++
title = "使用 Swiss Tables 加速 Go 的映射"
date = 2025-03-31T14:24:48+08:00
weight = 970
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://go.dev/blog/swisstable](https://go.dev/blog/swisstable)

# Faster Go maps with Swiss Tables  - 使用 Swiss Tables 加速 Go 的映射

Michael Pratt  

26 February 2025  

The hash table is a central data structure in computer science, and it provides the implementation for the map type in many languages, including Go.

​	哈希表是计算机科学中的核心数据结构，它为包括 Go 在内的许多语言中的映射类型提供了实现。

The concept of a hash table was [first described](https://spectrum.ieee.org/hans-peter-luhn-and-the-birth-of-the-hashing-algorithm) by Hans Peter Luhn in 1953 in an internal IBM memo that suggested speeding up search by placing items into “buckets” and using a linked list for overflow when buckets already contain an item. Today we would call this a [hash table using chaining](https://en.wikipedia.org/wiki/Hash_table#Separate_chaining).

​	哈希表的概念最早由 Hans Peter Luhn 于 1953 年在 IBM 内部备忘录中 [首次描述](https://spectrum.ieee.org/hans-peter-luhn-and-the-birth-of-the-hashing-algorithm)，他建议通过将项放入“桶”中并在桶已包含项时使用链表处理溢出来加速搜索。今天我们称之为 [使用链式法的哈希表](https://en.wikipedia.org/wiki/Hash_table#Separate_chaining)。

In 1954, Gene M. Amdahl, Elaine M. McGraw, and Arthur L. Samuel first used an “open addressing” scheme when programming the IBM 701. When a bucket already contains an item, the new item is placed in the next empty bucket. This idea was formalized and published in 1957 by W. Wesley Peterson in [“Addressing for Random-Access Storage”](https://ieeexplore.ieee.org/document/5392733). Today we would call this a [hash table using open addressing with linear probing](https://en.wikipedia.org/wiki/Hash_table#Open_addressing).

​	1954 年，Gene M. Amdahl、Elaine M. McGraw 和 Arthur L. Samuel 在编程 IBM 701 时首次使用了“开放寻址”方案。当一个桶已包含项时，新项被放入下一个空桶。这一想法由 W. Wesley Peterson 于 1957 年在 [“随机访问存储的寻址”](https://ieeexplore.ieee.org/document/5392733) 中正式化和发表。今天我们称之为 [使用开放寻址和线性探测的哈希表](https://en.wikipedia.org/wiki/Hash_table#Open_addressing)。

With data structures that have been around this long, it’s easy to think that they must be “done”; that we know everything there is to know about them and they can’t be improved anymore. That’s not true! Computer science research continues to make advancements in fundamental algorithms, both in terms of algorithmic complexity and taking advantage of modern CPU hardware. For example, Go 1.19 [switched the `sort` package](https://go.dev/doc/go1.19#sortpkgsort) from a traditional quicksort, to [pattern-defeating quicksort](https://arxiv.org/pdf/2106.05123.pdf), a novel sorting algorithm from Orson R. L. Peters, first described in 2015.

​	对于存在如此长时间的数据结构，很容易认为它们已经“完成”；我们已经了解了所有相关知识，无法再改进。但事实并非如此！计算机科学研究在基础算法上持续取得进展，无论是在算法复杂性还是利用现代 CPU 硬件方面。例如，Go 1.19 将 `sort` 包 [从传统快速排序切换](https://go.dev/doc/go1.19#sortpkgsort) 到 Orson R. L. Peters 于 2015 年首次描述的 [模式击败快速排序](https://arxiv.org/pdf/2106.05123.pdf)，这是一种新颖的排序算法。

Like sorting algorithms, hash table data structures continue to see improvements. In 2017, Sam Benzaquen, Alkis Evlogimenos, Matt Kulukundis, and Roman Perepelitsa at Google presented [a new C++ hash table design](https://www.youtube.com/watch?v=ncHmEUmJZf4), dubbed “Swiss Tables”. In 2018, their implementation was [open sourced in the Abseil C++ library](https://abseil.io/blog/20180927-swisstables).

​	与排序算法类似，哈希表数据结构也在不断改进。2017 年，谷歌的 Sam Benzaquen、Alkis Evlogimenos、Matt Kulukundis 和 Roman Perepelitsa 提出了 [一种新的 C++ 哈希表设计](https://www.youtube.com/watch?v=ncHmEUmJZf4)，称为“Swiss Tables”。2018 年，他们的实现 [在 Abseil C++ 库中开源](https://abseil.io/blog/20180927-swisstables)。

Go 1.24 includes a completely new implementation of the built-in map type, based on the Swiss Table design. In this blog post we’ll look at how Swiss Tables improve upon traditional hash tables, and at some of the unique challenges in bringing the Swiss Table design to Go’s maps.

​	Go 1.24 包含了基于 Swiss Table 设计的内置映射类型的全新实现。在这篇博客文章中，我们将探讨 Swiss Tables 如何改进传统哈希表，以及将 Swiss Table 设计引入 Go 映射时面临的一些独特挑战。

## Open-addressed hash table 开放寻址哈希表

Swiss Tables are a form of open-addressed hash table, so let’s do a quick overview of how a basic open-addressed hash table works.

​	Swiss Tables 是一种开放寻址哈希表形式，因此让我们快速概述一下基本的开放寻址哈希表是如何工作的。

In an open-addressed hash table, all items are stored in a single backing array. We’ll call each location in the array a *slot*. The slot to which a key belongs is primarily determined by the *hash function*, `hash(key)`. The hash function maps each key to an integer, where the same key always maps to the same integer, and different keys ideally follow a uniform random distribution of integers. The defining feature of open-addressed hash tables is that they resolve collisions by storing the key elsewhere in the backing array. So, if the slot is already full (a *collision*), then a *probe sequence* is used to consider other slots until an empty slot is found. Let’s take a look at a sample hash table to see how this works.

​	在开放寻址哈希表中，所有项都存储在单个后备数组中。我们将数组中的每个位置称为 *槽*。键所属的槽主要由 *哈希函数* `hash(key)` 确定。哈希函数将每个键映射到一个整数，同一键始终映射到同一整数，不同键理想上遵循均匀随机分布的整数。开放寻址哈希表的定义特征是通过将键存储在后备数组中的其他位置来解决冲突。因此，如果槽已满（即发生 *冲突*），则使用 *探测序列* 考虑其他槽，直到找到一个空槽。让我们看一个示例哈希表来了解其工作原理。

### Example 示例

Below you can see a 16-slot backing array for a hash table, and the key (if any) stored in each slot. The values are not shown, as they are not relevant to this example.

​	下面是一个哈希表的 16 槽后备数组，以及每个槽中存储的键（如果有的话）。值未显示，因为它们与此示例无关。

| Slot |  0   |  1   |  2   |  3   |  4   |  5   |  6   |  7   |  8   |  9   |  10  |  11  |  12  |  13  |  14  |  15  |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| Key  |      |      |      |  56  |  32  |  21  |      |      |      |      |      |  78  |      |      |      |      |

To insert a new key, we use the hash function to select a slot. Since there are only 16 slots, we need to restrict to this range, so we’ll use `hash(key) % 16` as the target slot. Suppose we want to insert key `98` and `hash(98) % 16 = 7`. Slot 7 is empty, so we simply insert 98 there. On the other hand, suppose we want to insert key `25` and `hash(25) % 16 = 3`. Slot 3 is a collision because it already contains key 56. Thus we cannot insert here.

​	要插入一个新键，我们使用哈希函数选择一个槽。由于只有 16 个槽，我们需要限制在这个范围内，因此我们将使用 `hash(key) % 16` 作为目标槽。假设我们要插入键 `98`，且 `hash(98) % 16 = 7`。槽 7 是空的，因此我们只需将 98 插入那里。另一方面，假设我们要插入键 `25`，且 `hash(25) % 16 = 3`。槽 3 发生冲突，因为它已包含键 56。因此我们无法在此插入。

We use a probe sequence to find another slot. There are a variety of well-known probe sequences. The original and most straightforward probe sequence is *linear probing*, which simply tries successive slots in order.

​	我们使用探测序列找到另一个槽。有多种众所周知的探测序列。最初且最直接的探测序列是 *线性探测*，它只是按顺序尝试连续的槽。

So, in the `hash(25) % 16 = 3` example, since slot 3 is in use, we would consider slot 4 next, which is also in use. So too is slot 5. Finally, we’d get to empty slot 6, where we’d store key 25.

​	因此，在 `hash(25) % 16 = 3` 的示例中，由于槽 3 已被使用，我们接下来会考虑槽 4，它也被使用。槽 5 也是如此。最后，我们会到达空的槽 6，在那里存储键 25。

Lookup follows the same approach. A lookup of key 25 would start at slot 3, check whether it contains key 25 (it does not), and then continue linear probing until it finds key 25 in slot 6.

​	查找遵循相同的方法。查找键 25 将从槽 3 开始，检查它是否包含键 25（它没有），然后继续线性探测，直到在槽 6 中找到键 25。

This example uses a backing array with 16 slots. What happens if we insert more than 16 elements? If the hash table runs out of space, it will grow, usually by doubling the size of the backing array. All existing entries are reinserted into the new backing array.

​	此示例使用了一个 16 槽的后备数组。如果我们插入超过 16 个元素会怎样？如果哈希表空间不足，它会增长，通常通过将后备数组的大小加倍。所有现有条目都会重新插入到新的后备数组中。

Open-addressed hash tables don’t actually wait until the backing array is completely full to grow because as the array gets more full, the average length of each probe sequence increases. In the example above using key 25, we must probe 4 different slots to find an empty slot. If the array had only one empty slot, the worst case probe length would be O(n). That is, you may need to scan the entire array. The proportion of used slots is called the *load factor*, and most hash tables define a *maximum load factor* (typically 70-90%) at which point they will grow to avoid the extremely long probe sequences of very full hash tables.

​	开放寻址哈希表实际上不会等到后备数组完全满才增长，因为随着数组越来越满，每个探测序列的平均长度会增加。在上面使用键 25 的示例中，我们必须探测 4 个不同的槽才能找到一个空槽。如果数组只有一个空槽，最坏情况下的探测长度将是 O(n)。也就是说，你可能需要扫描整个数组。已使用槽的比例称为 *负载因子*，大多数哈希表定义了一个 *最大负载因子*（通常为 70-90%），此时它们会增长以避免非常满的哈希表带来的极长探测序列。

## Swiss Table - Swiss 表

The Swiss Table [design](https://abseil.io/about/design/swisstables) is also a form of open-addressed hash table. Let’s see how it improves over a traditional open-addressed hash table. We still have a single backing array for storage, but we will break the array into logical *groups* of 8 slots each. (Larger group sizes are possible as well. More on that below.)

​	Swiss Table [设计](https://abseil.io/about/design/swisstables) 也是一种开放寻址哈希表形式。让我们看看它如何改进传统开放寻址哈希表。我们仍然使用单个后备数组进行存储，但我们会将数组分成每组 8 个槽的逻辑 *组*。（更大的组大小也是可能的，详情见下文。）

In addition, each group has a 64-bit *control word* for metadata. Each of the 8 bytes in the control word corresponds to one of the slots in the group. The value of each byte denotes whether that slot is empty, deleted, or in use. If it is in use, the byte contains the lower 7 bits of the hash for that slot’s key (called `h2`).

​	此外，每组都有一个 64 位 *控制字* 用于元数据。控制字中的 8 个字节每个对应组中的一个槽。每个字节的值表示该槽是空的、已删除的还是正在使用的。如果正在使用，该字节包含该槽键的哈希值的低 7 位（称为 `h2`）。

|      | Group 0 | Group 1 |      |      |      |      |      |      |      |      |      |      |      |      |      |      |
| :--: | :-----: | :-----: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| Slot |    0    |    1    |  2   |  3   |  4   |  5   |  6   |  7   |  0   |  1   |  2   |  3   |  4   |  5   |  6   |  7   |
| Key  |   56    |   32    |  21  |      |      |      |      |      |  78  |      |      |      |      |      |      |      |

|      | 64-bit control word 0 | 64-bit control word 1 |      |      |      |      |      |      |      |      |      |      |      |      |      |      |
| :--: | :-------------------: | :-------------------: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| Slot |           0           |           1           |  2   |  3   |  4   |  5   |  6   |  7   |  0   |  1   |  2   |  3   |  4   |  5   |  6   |  7   |
|  h2  |          23           |          89           |  50  |      |      |      |      |      |  47  |      |      |      |      |      |      |      |

Insertion works as follows:

​	插入过程如下：

1. Compute `hash(key)` and break the hash into two parts: the upper 57-bits (called `h1`) and the lower 7 bits (called `h2`). 计算 `hash(key)` 并将哈希值分为两部分：高 57 位（称为 `h1`）和低 7 位（称为 `h2`）。
2. The upper bits (`h1`) are used to select the first group to consider: `h1 % 2` in this case, as there are only 2 groups.高位（`h1`）用于选择要考虑的第一个组：在此例中为 `h1 % 2`，因为只有 2 个组。
3. Within a group, all slots are equally eligible to hold the key. We must first determine whether any slot already contains this key, in which case this is an update rather than a new insertion. 在一个组内，所有槽都同样有资格持有该键。我们必须首先确定是否有任何槽已经包含此键，在这种情况下，这是更新而不是新插入。
4. If no slot contains the key, then we look for an empty slot to place this key 如果没有槽包含该键，那么我们寻找一个空槽来放置此键。
5. If no slot is empty, then we continue the probe sequence by searching the next group. 如果没有空槽，则通过搜索下一个组继续探测序列。


Lookup follows the same basic process. If we find an empty slot in step 4, then we know an insertion would have used this slot and can stop the search.

​	查找遵循相同的基本过程。如果我们在步骤 4 中找到一个空槽，那么我们知道插入会使用这个槽，可以停止搜索。

Step 3 is where the Swiss Table magic happens. We need to check whether any slot in a group contains the desired key. Naively, we could just do a linear scan and compare all 8 keys. However, the control word lets us do this more efficiently. Each byte contains the lower 7 bits of the hash (`h2`) for that slot. If we determine which bytes of the control word contain the `h2` we are looking for, we’ll have a set of candidate matches.

​	步骤 3 是 Swiss Table 魔力的体现。我们需要检查组中的任何槽是否包含所需的键。简单地，我们可以进行线性扫描并比较所有 8 个键。然而，控制字让我们更高效地做到这一点。每个字节包含该槽的哈希值的低 7 位（`h2`）。如果我们确定控制字中的哪些字节包含我们寻找的 `h2`，我们将得到一组候选匹配。

In other words, we want to do a byte-by-byte equality comparison within the control word. For example, if we are looking for key 32, where `h2 = 89`, the operation we want looks like so.

​	换句话说，我们希望在控制字内进行逐字节的相等比较。例如，如果我们寻找键 32，其中 `h2 = 89`，我们想要的操作如下所示。

| **Test word** **测试字** | 89   | 89   | 89   | 89   | 89   | 89   | 89   | 89   |
| ------------------------ | ---- | ---- | ---- | ---- | ---- | ---- | ---- | ---- |
| **Comparison 比较  **    | ==   | ==   | ==   | ==   | ==   | ==   | ==   | ==   |
| **Control word 控制字**  | 23   | 89   | 50   | -    | -    | -    | -    | -    |
| **Result 结果**          | 0    | 1    | 0    | 0    | 0    | 0    | 0    | 0    |

This is an operation supported by [SIMD](https://en.wikipedia.org/wiki/Single_instruction,_multiple_data) hardware, where a single instruction performs parallel operations on independent values within a larger value (*vector*). In this case, we [can implement this operation](https://cs.opensource.google/go/go/+/master:src/internal/runtime/maps/group.go;drc=a08984bc8f2acacebeeadf7445ecfb67b7e7d7b1;l=155?ss=go) using a set of standard arithmetic and bitwise operations when special SIMD hardware is not available.

​	这是 [SIMD](https://en.wikipedia.org/wiki/Single_instruction,_multiple_data) 硬件支持的操作，其中单一指令对较大值（*向量*）内的独立值执行并行操作。在这种情况下，当特殊 SIMD 硬件不可用时，我们 [可以使用一组标准算术和位运算实现此操作](https://cs.opensource.google/go/go/+/master:src/internal/runtime/maps/group.go;drc=a08984bc8f2acacebeeadf7445ecfb67b7e7d7b1;l=155?ss=go)。

The result is a set of candidate slots. Slots where `h2` does not match do not have a matching key, so they can be skipped. Slots where `h2` does match are potential matches, but we must still check the entire key, as there is potential for collisions (1/128 probability of collision with a 7-bit hash, so still quite low).

​	结果是一组候选槽。`h2` 不匹配的槽没有匹配的键，因此可以跳过。`h2` 匹配的槽是潜在匹配，但我们仍需检查整个键，因为存在冲突的可能性（7 位哈希的冲突概率为 1/128，因此仍然很低）。

This operation is very powerful, as we have effectively performed 8 steps of a probe sequence at once, in parallel. This speeds up lookup and insertion by reducing the average number of comparisons we need to perform. This improvement to probing behavior allowed both the Abseil and Go implementations to increase the maximum load factor of Swiss Table maps compared to prior maps, which lowers the average memory footprint.

​	此操作非常强大，因为我们有效地一次性并行执行了探测序列的 8 个步骤。这通过减少我们需要执行的平均比较次数来加速查找和插入。对探测行为的这种改进使 Abseil 和 Go 实现能够增加 Swiss Table 映射的最大负载因子，相较于之前的映射，这降低了平均内存占用。

## Go challenges - Go 挑战

Go’s built-in map type has some unusual properties that pose additional challenges to adopting a new map design. Two were particularly tricky to deal with.

​	Go 的内置映射类型具有一些不寻常的特性，给采用新的映射设计带来了额外的挑战。其中两个尤其棘手。

### Incremental growth 增量增长

When a hash table reaches its maximum load factor, it needs to grow the backing array. Typically this means the next insertion doubles the size of the array, and copies all entries to the new array. Imagine inserting into a map with 1GB of entries. Most insertions are very fast, but the one insertion that needs to grow the map from 1GB to 2GB will need to copy 1GB of entries, which will take a long time.

​	当哈希表达到其最大负载因子时，它需要扩展后备数组。通常这意味着下一次插入会将数组大小加倍，并将所有条目复制到新数组中。想象插入一个包含 1GB 条目的映射。大多数插入都很快，但需要将映射从 1GB 扩展到 2GB 的那次插入需要复制 1GB 的条目，这将花费很长时间。

Go is frequently used for latency-sensitive servers, so we don’t want operations on built-in types that can have arbitrarily large impact on tail latency. Instead, Go maps grow incrementally, so that each insertion has an upper bound on the amount of growth work it must do. This bounds the latency impact of a single map insertion.

​	Go 经常用于对延迟敏感的服务器，因此我们不希望内置类型的操作对尾部延迟产生任意大的影响。相反，Go 映射是增量增长的，因此每次插入对必须完成的增长工作量有一个上限。这限制了单次映射插入的延迟影响。

Unfortunately, the Abseil (C++) Swiss Table design assumes all at once growth, and the probe sequence depends on the total group count, making it difficult to break up.

​	不幸的是，Abseil（C++）Swiss Table 设计假定一次性增长，探测序列依赖于总组数，这使得分解变得困难。

Go’s built-in map addresses this with another layer of indirection by splitting each map into multiple Swiss Tables. Rather than a single Swiss Table implementing the entire map, each map consists of one or more independent tables that cover a subset of the key space. An individual table stores a maximum of 1024 entries. A variable number of upper bits in the hash are used to select which table a key belongs to. This is a form of [*extendible hashing*](https://en.wikipedia.org/wiki/Extendible_hashing), where the number of bits used increases as needed to differentiate the total number of tables.

​	Go 的内置映射通过将每个映射拆分为多个 Swiss Tables 并增加另一层间接性来解决此问题。不是单个 Swiss Table 实现整个映射，每个映射由一个或多个覆盖键空间子集的独立表组成。单个表最多存储 1024 个条目。哈希中的可变数量的高位用于选择键所属的表。这是一种 [*可扩展哈希*](https://en.wikipedia.org/wiki/Extendible_hashing) 形式，其中使用的位数根据需要增加以区分表的总数。

During insertion, if an individual table needs to grow, it will do so all at once, but other tables are unaffected. Thus the upper bound for a single insertion is the latency of growing a 1024-entry table into two 1024-entry tables, copying 1024 entries.

​	在插入期间，如果单个表需要增长，它会一次性完成，但其他表不受影响。因此，单次插入的上限是将一个 1024 条目表扩展为两个 1024 条目表并复制 1024 个条目的延迟。

### Modification during iteration 迭代期间的修改

Many hash table designs, including Abseil’s Swiss Tables, forbid modifying the map during iteration. The Go language spec [explicitly allows](https://go.dev/ref/spec#For_statements:~:text=The iteration order,iterations is 0.) modifications during iteration, with the following semantics:

​	许多哈希表设计，包括 Abseil 的 Swiss Tables，禁止在迭代期间修改映射。Go 语言规范 [明确允许](https://go.dev/ref/spec#For_statements:~:text=The iteration order,iterations is 0.) 在迭代期间修改，具有以下语义：

- If an entry is deleted before it is reached, it will not be produced. 如果在到达之前删除了一个条目，它将不会被生成。

- If an entry is updated before it is reached, the updated value will be produced. 如果在到达之前更新了一个条目，将生成更新后的值。

- If a new entry is added, it may or may not be produced. 如果添加了一个新条目，它可能被生成，也可能不被生成。

  

A typical approach to hash table iteration is to simply walk through the backing array and produce values in the order they are laid out in memory. This approach runs afoul of the above semantics, most notably because insertions may grow the map, which would shuffle the memory layout.

​	哈希表迭代的典型方法是简单地遍历后备数组，并按内存布局顺序生成值。这种方法与上述语义相冲突，最明显的是因为插入可能会扩展映射，这会打乱内存布局。

We can avoid the impact of shuffle during growth by having the iterator keep a reference to the table it is currently iterating over. If that table grows during iteration, we keep using the old version of the table and thus continue to deliver keys in the order of the old memory layout.

​	我们可以通过让迭代器保留对当前正在迭代的表的引用来避免增长期间打乱的影响。如果该表在迭代期间增长，我们继续使用旧版本的表，从而按旧内存布局的顺序继续传递键。

Does this work with the above semantics? New entries added after growth will be missed entirely, as they are only added to the grown table, not the old table. That’s fine, as the semantics allow new entries not to be produced. Updates and deletions are a problem, though: using the old table could produce stale or deleted entries.

​	这与上述语义兼容吗？增长后添加的新条目将完全被错过，因为它们只添加到扩展后的表中，而不是旧表中。这没问题，因为语义允许新条目不被生成。然而，更新和删除是个问题：使用旧表可能会生成过时或已删除的条目。

This edge case is addressed by using the old table only to determine the iteration order. Before actually returning the entry, we consult the grown table to determine whether the entry still exists, and to retrieve the latest value.

​	这个边缘情况通过仅使用旧表确定迭代顺序来解决。在实际返回条目之前，我们查阅扩展后的表以确定条目是否仍然存在，并检索最新值。

This covers all the core semantics, though there are even more small edge cases not covered here. Ultimately, the permissiveness of Go maps with iteration results in iteration being the most complex part of Go’s map implementation.

这涵盖了所有核心语义，尽管这里未涉及更多小的边缘情况。最终，Go 映射对迭代的宽容性导致迭代成为 Go 映射实现中最复杂的部分。

## Future work 未来工作

In [microbenchmarks](https://go.dev/issue/54766#issuecomment-2542444404), map operations are up to 60% faster than in Go 1.23. Exact performance improvement varies quite a bit due to the wide variety of operations and uses of maps, and some edge cases do regress compared to Go 1.23. Overall, in full application benchmarks, we found a geometric mean CPU time improvement of around 1.5%.

​	在 [微基准测试](https://go.dev/issue/54766#issuecomment-2542444404) 中，映射操作比 Go 1.23 快了高达 60%。由于映射操作和使用的多样性，确切的性能改进差异很大，一些边缘情况与 Go 1.23 相比有所退步。总体而言，在完整应用程序基准测试中，我们发现 CPU 时间的几何平均改进约为 1.5%。

There are more map improvements we want to investigate for future Go releases. For example, we may be able to [increase the locality of](https://go.dev/issue/70835) operations on maps that are not in the CPU cache.

​	我们希望为未来的 Go 版本调查更多映射改进。例如，我们可能能够 [增加不在 CPU 缓存中的映射操作的局部性](https://go.dev/issue/70835)。

We could also further improve the control word comparisons. As described above, we have a portable implementation using standard arithmetic and bitwise operations. However, some architectures have SIMD instructions that perform this sort of comparison directly. Go 1.24 already uses 8-byte SIMD instructions for amd64, but we could extend support to other architectures. More importantly, while standard instructions operate on up to 8-byte words, SIMD instructions nearly always support at least 16-byte words. This means we could increase the group size to 16 slots, and perform 16 hash comparisons in parallel instead of 8. This would further decrease the average number of probes required for lookups.

​	我们还可以进一步改进控制字比较。如上所述，我们有一个使用标准算术和位运算的可移植实现。然而，一些架构具有直接执行此类比较的 SIMD 指令。Go 1.24 已为 amd64 使用 8 字节 SIMD 指令，但我们可以扩展到其他架构的支持。更重要的是，虽然标准指令操作最多 8 字节字，但 SIMD 指令几乎总是支持至少 16 字节字。这意味着我们可以将组大小增加到 16 个槽，并行执行 16 个哈希比较而不是 8 个。这将进一步减少查找所需的平均探测次数。

## Acknowledgements 致谢

A Swiss Table-based Go map implementation has been a long time coming, and involved many contributors. I want to thank YunHao Zhang ([@zhangyunhao116](https://github.com/zhangyunhao116)), PJ Malloy ([@thepudds](https://github.com/thepudds)), and [@andy-wm-arthur](https://github.com/andy-wm-arthur) for building initial versions of a Go Swiss Table implementation. Peter Mattis ([@petermattis](https://github.com/petermattis)) combined these ideas with solutions to the Go challenges above to build [`github.com/cockroachdb/swiss`](https://pkg.go.dev/github.com/cockroachdb/swiss), a Go-spec compliant Swiss Table implementation. The Go 1.24 built-in map implementation is heavily based on Peter’s work. Thank you to everyone in the community that contributed!

​	基于 Swiss Table 的 Go 映射实现历经多年，涉及众多贡献者。我要感谢 YunHao Zhang（[@zhangyunhao116](https://github.com/zhangyunhao116)）、PJ Malloy（[@thepudds](https://github.com/thepudds)）和 [@andy-wm-arthur](https://github.com/andy-wm-arthur) 构建了 Go Swiss Table 实现的初始版本。Peter Mattis（[@petermattis](https://github.com/petermattis)）将这些想法与上述 Go 挑战的解决方案结合，构建了 [`github.com/cockroachdb/swiss`](https://pkg.go.dev/github.com/cockroachdb/swiss)，一个符合 Go 规范的 Swiss Table 实现。Go 1.24 的内置映射实现很大程度上基于 Peter 的工作。感谢社区中所有贡献者！