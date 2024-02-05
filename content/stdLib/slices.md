+++
title = "slices"
date = 2023-11-05T14:26:18+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

> 原文：[https://pkg.go.dev/slices](https://pkg.go.dev/slices)

## 概述 

Package slices defines various functions useful with slices of any type.

​	`slices` 包定义了各种类型切片的有用函数。

## 常量

This section is empty.

## 常量

This section is empty.

## 函数

### func BinarySearch 

``` go
func BinarySearch[S ~[]E, E cmp.Ordered](x S, target E) (int, bool)
```

BinarySearch searches for target in a sorted slice and returns the position where target is found, or the position where target would appear in the sort order; it also returns a bool saying whether the target is really found in the slice. The slice must be sorted in increasing order.

​	`BinarySearch` 在已排序的切片中搜索目标，并返回找到目标的位置，或者在排序顺序中目标应该出现的位置；它还返回一个布尔值，表示目标是否真正在切片中找到。切片必须以递增顺序排序。

#### BinarySearch Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	n, found := slices.BinarySearch(names, "Vera")
	fmt.Println("Vera:", n, found)
	n, found = slices.BinarySearch(names, "Bill")
	fmt.Println("Bill:", n, found)
}
Output:

Vera: 2 true
Bill: 1 false
```
### func BinarySearchFunc 

``` go
func BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool)
```

BinarySearchFunc works like [BinarySearch](https://pkg.go.dev/slices#BinarySearch), but uses a custom comparison function. The slice must be sorted in increasing order, where "increasing" is defined by cmp. cmp should return 0 if the slice element matches the target, a negative number if the slice element precedes the target, or a positive number if the slice element follows the target. cmp must implement the same ordering as the slice, such that if cmp(a, t) < 0 and cmp(b, t) >= 0, then a must precede b in the slice.

​	`BinarySearchFunc` 函数的工作原理类似于 [BinarySearch](#binarySearch)函数，但使用自定义的比较函数。切片必须以递增顺序排序，其中“递增（increasing）”由 `cmp` 定义。如果切片元素与目标匹配，则 `cmp` 应返回 0；如果切片元素在目标之前，则返回一个负数；如果切片元素在目标之后，则返回一个正数。`cmp` 必须实现与切片相同的排序顺序，以便如果 `cmp(a, t) < 0` 和 `cmp(b, t) >= 0`，那么 `a` 必须在切片中出现在 `b` 之前。

#### BinarySearchFunc Example

``` go
package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 55},
		{"Bob", 24},
		{"Gopher", 13},
	}
	n, found := slices.BinarySearchFunc(people, Person{"Bob", 0}, func(a, b Person) int {
		return cmp.Compare(a.Name, b.Name)
	})
	fmt.Println("Bob:", n, found)
}
Output:

Bob: 1 true
```
### func Clip 

``` go
func Clip[S ~[]E, E any](s S) S
```

Clip removes unused capacity from the slice, returning s[:len(s):len(s)].

​	`Clip` 函数从切片中移除未使用的容量，并返回 `s[:len(s):len(s)]`。

### func Clone 

``` go
func Clone[S ~[]E, E any](s S) S
```

Clone returns a copy of the slice. The elements are copied using assignment, so this is a shallow clone.

​	`Clone` 函数返回切片的副本。元素是通过赋值复制的，因此这是`浅复制`。

### func Compact 

``` go
func Compact[S ~[]E, E comparable](s S) S
```

Compact replaces consecutive runs of equal elements with a single copy. This is like the uniq command found on Unix. Compact modifies the contents of the slice s and returns the modified slice, which may have a smaller length. When Compact discards m elements in total, it might not modify the elements s[len(s)-m:len(s)]. If those elements contain pointers you might consider zeroing those elements so that objects they reference can be garbage collected.

​	这个函数`Compact`的作用是将连续相等的元素替换为单个副本，类似于Unix中的uniq命令。它修改切片`s`的内容并返回修改后的切片，返回切片的长度可能会变小。当`Compact`总共丢弃m个元素时，它可能不会修改元素`s[len(s)-m:len(s)]`。如果这些元素包含指针，您可能需要将这些元素置零，以便垃圾回收可以收集它们引用的对象。

#### Compact Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	seq := []int{0, 1, 1, 2, 3, 5, 8}
	seq = slices.Compact(seq)
	fmt.Println(seq)
}
Output:

[0 1 2 3 5 8]
```
### func CompactFunc 

``` go
func CompactFunc[S ~[]E, E any](s S, eq func(E, E) bool) S
```

CompactFunc is like [Compact](https://pkg.go.dev/slices#Compact) but uses an equality function to compare elements. For runs of elements that compare equal, CompactFunc keeps the first one.

​	`CompactFunc`函数类似于[Compact](#compact)，但它使用一个等价函数来比较元素。对于比较相等的元素序列，`CompactFunc`保留第一个元素。

#### CompactFunc Example

``` go
package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"bob", "Bob", "alice", "Vera", "VERA"}
	names = slices.CompactFunc(names, func(a, b string) bool {
		return strings.ToLower(a) == strings.ToLower(b)
	})
	fmt.Println(names)
}
Output:

[bob alice Vera]
```
### func Compare 

``` go
func Compare[S ~[]E, E cmp.Ordered](s1, s2 S) int
```

Compare compares the elements of s1 and s2, using [cmp.Compare](https://pkg.go.dev/cmp#Compare) on each pair of elements. The elements are compared sequentially, starting at index 0, until one element is not equal to the other. The result of comparing the first non-matching elements is returned. If both slices are equal until one of them ends, the shorter slice is considered less than the longer one. The result is 0 if s1 == s2, -1 if s1 < s2, and +1 if s1 > s2.

​	Compare函数比较s1和s2的元素，使用[cmp.Compare](https://pkg.go.dev/cmp#Compare)来比较每一对元素。元素按顺序进行比较，从索引0开始，直到找到不相等的元素。返回比较第一个不匹配元素的结果。如果两个切片相等直到其中一个切片结束，那么较短的切片被认为是小于较长的切片。如果s1等于s2，结果为0；如果s1小于s2，结果为-1；如果s1大于s2，结果为+1。

#### Compare Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	fmt.Println("Equal:", slices.Compare(names, []string{"Alice", "Bob", "Vera"}))
	fmt.Println("V < X:", slices.Compare(names, []string{"Alice", "Bob", "Xena"}))
	fmt.Println("V > C:", slices.Compare(names, []string{"Alice", "Bob", "Cat"}))
	fmt.Println("3 > 2:", slices.Compare(names, []string{"Alice", "Bob"}))
}
Output:

Equal: 0
V < X: -1
V > C: 1
3 > 2: 1
```
### func CompareFunc 

``` go
func CompareFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, cmp func(E1, E2) int) int
```

CompareFunc is like [Compare](https://pkg.go.dev/slices#Compare) but uses a custom comparison function on each pair of elements. The result is the first non-zero result of cmp; if cmp always returns 0 the result is 0 if len(s1) == len(s2), -1 if len(s1) < len(s2), and +1 if len(s1) > len(s2).

#### CompareFunc Example

``` go
package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	numbers := []int{0, 43, 8}
	strings := []string{"0", "0", "8"}
	result := slices.CompareFunc(numbers, strings, func(n int, s string) int {
		sn, err := strconv.Atoi(s)
		if err != nil {
			return 1
		}
		return cmp.Compare(n, sn)
	})
	fmt.Println(result)
}
Output:

1
```
### func Contains 

``` go
func Contains[S ~[]E, E comparable](s S, v E) bool
```

Contains reports whether v is present in s.

### func ContainsFunc 

``` go
func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool
```

ContainsFunc reports whether at least one element e of s satisfies f(e).

#### ContainsFunc Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	hasNegative := slices.ContainsFunc(numbers, func(n int) bool {
		return n < 0
	})
	fmt.Println("Has a negative:", hasNegative)
	hasOdd := slices.ContainsFunc(numbers, func(n int) bool {
		return n%2 != 0
	})
	fmt.Println("Has an odd number:", hasOdd)
}
Output:

Has a negative: true
Has an odd number: false
```
### func Delete 

``` go
func Delete[S ~[]E, E any](s S, i, j int) S
```

Delete removes the elements s[i:j] from s, returning the modified slice. Delete panics if s[i:j] is not a valid slice of s. Delete is O(len(s)-j), so if many items must be deleted, it is better to make a single call deleting them all together than to delete one at a time. Delete might not modify the elements s[len(s)-(j-i):len(s)]. If those elements contain pointers you might consider zeroing those elements so that objects they reference can be garbage collected.

#### Delete Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	letters := []string{"a", "b", "c", "d", "e"}
	letters = slices.Delete(letters, 1, 4)
	fmt.Println(letters)
}
Output:

[a e]
```
### func DeleteFunc 

``` go
func DeleteFunc[S ~[]E, E any](s S, del func(E) bool) S
```

DeleteFunc removes any elements from s for which del returns true, returning the modified slice. When DeleteFunc removes m elements, it might not modify the elements s[len(s)-m:len(s)]. If those elements contain pointers you might consider zeroing those elements so that objects they reference can be garbage collected.

#### DeleteFunc Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	seq := []int{0, 1, 1, 2, 3, 5, 8}
	seq = slices.DeleteFunc(seq, func(n int) bool {
		return n%2 != 0 // delete the odd numbers
	})
	fmt.Println(seq)
}
Output:

[0 2 8]
```
### func Equal 

``` go
func Equal[S ~[]E, E comparable](s1, s2 S) bool
```

Equal reports whether two slices are equal: the same length and all elements equal. If the lengths are different, Equal returns false. Otherwise, the elements are compared in increasing index order, and the comparison stops at the first unequal pair. Floating point NaNs are not considered equal.

#### Equal Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, 8}
	fmt.Println(slices.Equal(numbers, []int{0, 42, 8}))
	fmt.Println(slices.Equal(numbers, []int{10}))
}
Output:

true
false
```
### func EqualFunc 

``` go
func EqualFunc[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, s2 S2, eq func(E1, E2) bool) bool
```

EqualFunc reports whether two slices are equal using an equality function on each pair of elements. If the lengths are different, EqualFunc returns false. Otherwise, the elements are compared in increasing index order, and the comparison stops at the first index for which eq returns false.

#### EqualFunc Example

``` go
package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	numbers := []int{0, 42, 8}
	strings := []string{"000", "42", "0o10"}
	equal := slices.EqualFunc(numbers, strings, func(n int, s string) bool {
		sn, err := strconv.ParseInt(s, 0, 64)
		if err != nil {
			return false
		}
		return n == int(sn)
	})
	fmt.Println(equal)
}
Output:

true
```
### func Grow 

``` go
func Grow[S ~[]E, E any](s S, n int) S
```

Grow increases the slice's capacity, if necessary, to guarantee space for another n elements. After Grow(n), at least n elements can be appended to the slice without another allocation. If n is negative or too large to allocate the memory, Grow panics.

### func Index 

``` go
func Index[S ~[]E, E comparable](s S, v E) int
```

Index returns the index of the first occurrence of v in s, or -1 if not present.

#### Index Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, 8}
	fmt.Println(slices.Index(numbers, 8))
	fmt.Println(slices.Index(numbers, 7))
}
Output:

2
-1
```
### func IndexFunc 

``` go
func IndexFunc[S ~[]E, E any](s S, f func(E) bool) int
```

IndexFunc returns the first index i satisfying f(s[i]), or -1 if none do.

#### IndexFunc Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	i := slices.IndexFunc(numbers, func(n int) bool {
		return n < 0
	})
	fmt.Println("First negative at index", i)
}
Output:

First negative at index 2
```
### func Insert 

``` go
func Insert[S ~[]E, E any](s S, i int, v ...E) S
```

Insert inserts the values v... into s at index i, returning the modified slice. The elements at s[i:] are shifted up to make room. In the returned slice r, r[i] == v[0], and r[i+len(v)] == value originally at r[i]. Insert panics if i is out of range. This function is O(len(s) + len(v)).

#### Insert Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	names = slices.Insert(names, 1, "Bill", "Billie")
	names = slices.Insert(names, len(names), "Zac")
	fmt.Println(names)
}
Output:

[Alice Bill Billie Bob Vera Zac]
```
### func IsSorted 

``` go
func IsSorted[S ~[]E, E cmp.Ordered](x S) bool
```

IsSorted reports whether x is sorted in ascending order.

#### IsSorted Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(slices.IsSorted([]string{"Alice", "Bob", "Vera"}))
	fmt.Println(slices.IsSorted([]int{0, 2, 1}))
}
Output:

true
false
```
### func IsSortedFunc 

``` go
func IsSortedFunc[S ~[]E, E any](x S, cmp func(a, b E) int) bool
```

IsSortedFunc reports whether x is sorted in ascending order, with cmp as the comparison function as defined by [SortFunc](https://pkg.go.dev/slices#SortFunc).

#### IsSortedFunc Example

``` go
package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"alice", "Bob", "VERA"}
	isSortedInsensitive := slices.IsSortedFunc(names, func(a, b string) int {
		return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	fmt.Println(isSortedInsensitive)
	fmt.Println(slices.IsSorted(names))
}
Output:

true
false
```
### func Max 

``` go
func Max[S ~[]E, E cmp.Ordered](x S) E
```

Max returns the maximal value in x. It panics if x is empty. For floating-point E, Max propagates NaNs (any NaN value in x forces the output to be NaN).

#### Max Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	fmt.Println(slices.Max(numbers))
}
Output:

42
```
### func MaxFunc 

``` go
func MaxFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E
```

MaxFunc returns the maximal value in x, using cmp to compare elements. It panics if x is empty. If there is more than one maximal element according to the cmp function, MaxFunc returns the first one.

#### MaxFunc Example

``` go
package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Gopher", 13},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 55},
	}
	firstOldest := slices.MaxFunc(people, func(a, b Person) int {
		return cmp.Compare(a.Age, b.Age)
	})
	fmt.Println(firstOldest.Name)
}
Output:

Alice
```
### func Min 

``` go
func Min[S ~[]E, E cmp.Ordered](x S) E
```

Min returns the minimal value in x. It panics if x is empty. For floating-point numbers, Min propagates NaNs (any NaN value in x forces the output to be NaN).

#### Min  Example

```go
package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{0, 42, -10, 8}
	fmt.Println(slices.Min(numbers))
}
Output:

-10
```

### func MinFunc 

``` go
func MinFunc[S ~[]E, E any](x S, cmp func(a, b E) int) E
```

MinFunc returns the minimal value in x, using cmp to compare elements. It panics if x is empty. If there is more than one minimal element according to the cmp function, MinFunc returns the first one.

#### MinFunc Example

``` go
package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Gopher", 13},
		{"Bob", 5},
		{"Vera", 24},
		{"Bill", 5},
	}
	firstYoungest := slices.MinFunc(people, func(a, b Person) int {
		return cmp.Compare(a.Age, b.Age)
	})
	fmt.Println(firstYoungest.Name)
}
Output:

Bob
```
### func Replace 

``` go
func Replace[S ~[]E, E any](s S, i, j int, v ...E) S
```

Replace replaces the elements s[i:j] by the given v, and returns the modified slice. Replace panics if s[i:j] is not a valid slice of s.

#### Replace Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera", "Zac"}
	names = slices.Replace(names, 1, 3, "Bill", "Billie", "Cat")
	fmt.Println(names)
}
Output:

[Alice Bill Billie Cat Zac]
```
### func Reverse 

``` go
func Reverse[S ~[]E, E any](s S)
```

Reverse reverses the elements of the slice in place.

#### Reverse Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"alice", "Bob", "VERA"}
	slices.Reverse(names)
	fmt.Println(names)
}
Output:

[VERA Bob alice]
```
### func Sort 

``` go
func Sort[S ~[]E, E cmp.Ordered](x S)
```

Sort sorts a slice of any ordered type in ascending order. When sorting floating-point numbers, NaNs are ordered before other values.

#### Sort Example

``` go
package main

import (
	"fmt"
	"slices"
)

func main() {
	smallInts := []int8{0, 42, -10, 8}
	slices.Sort(smallInts)
	fmt.Println(smallInts)
}
Output:

[-10 0 8 42]
```
### func SortFunc 

``` go
func SortFunc[S ~[]E, E any](x S, cmp func(a, b E) int)
```

SortFunc sorts the slice x in ascending order as determined by the cmp function. This sort is not guaranteed to be stable. cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b.

SortFunc requires that cmp is a strict weak ordering. See https://en.wikipedia.org/wiki/Weak_ordering#Strict_weak_orderings.

#### SortFunc Example (CaseInsensitive)

``` go
package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	names := []string{"Bob", "alice", "VERA"}
	slices.SortFunc(names, func(a, b string) int {
		return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	fmt.Println(names)
}
Output:

[alice Bob VERA]
```
#### SortFunc Example (MultiField)

``` go
package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Gopher", 13},
		{"Alice", 55},
		{"Bob", 24},
		{"Alice", 20},
	}
	slices.SortFunc(people, func(a, b Person) int {
		if n := cmp.Compare(a.Name, b.Name); n != 0 {
			return n
		}
		// If names are equal, order by age
		return cmp.Compare(a.Age, b.Age)
	})
	fmt.Println(people)
}
Output:

[{Alice 20} {Alice 55} {Bob 24} {Gopher 13}]
```
### func SortStableFunc 

``` go
func SortStableFunc[S ~[]E, E any](x S, cmp func(a, b E) int)
```

SortStableFunc sorts the slice x while keeping the original order of equal elements, using cmp to compare elements in the same way as [SortFunc](https://pkg.go.dev/slices#SortFunc).

#### SortStableFunc Example

``` go
package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Gopher", 13},
		{"Alice", 20},
		{"Bob", 24},
		{"Alice", 55},
	}
	// Stable sort by name, keeping age ordering of Alices intact
	slices.SortStableFunc(people, func(a, b Person) int {
		return cmp.Compare(a.Name, b.Name)
	})
	fmt.Println(people)
}

```
### Types 

This section is empty.