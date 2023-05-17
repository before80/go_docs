+++
title = "sort"
linkTitle = "sort"
date = 2023-05-17T09:59:21+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# sort

[https://pkg.go.dev/sort@go1.20.1](https://pkg.go.dev/sort@go1.20.1)

​	sort包提供了对切片和用户定义集合进行排序的基本功能。

#### Example
``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge 实现了基于 Age 字段对 []Person 进行 sort.Interface 排序
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
    // 排序一个 slice 有两种方法。
    // 首先，可以为 slice 类型定义一组方法，就像 ByAge 一样，
    // 然后调用 sort.Sort。在这个第一个示例中，我们使用了这种技术。
	sort.Sort(ByAge(people))
	fmt.Println(people)

    // 另一种方法是使用 sort.Slice 和自定义的 Less 函数，
    // 该函数可以作为闭包提供。
    // 在这种情况下，不需要方法。(如果它们存在，则会被忽略。)
    // 这里我们按相反的顺序重新排序：比较闭包和 ByAge.Less。
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println(people)

}
Output:

[Bob: 31 John: 42 Michael: 17 Jenny: 26]
[Michael: 17 Jenny: 26 Bob: 31 John: 42]
[John: 42 Bob: 31 Jenny: 26 Michael: 17]
```

#### Example(SortKeys)

​	ExampleSortKeys 展示了使用可编程排序标准对结构体类型进行排序的技术。

``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

// 一些类型定义以明确单位。
type earthMass float64
type au float64

// Planet 定义了太阳系物体的属性。
type Planet struct {
	name     string
	mass     earthMass
	distance au
}

// By 是一个 "less" 函数的类型，它定义了其 Planet 参数的排序顺序。
type By func(p1, p2 *Planet) bool

// Sort 是函数类型 By 的方法，它根据函数对参数切片进行排序。
func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		by:      by, // Sort 方法的接收者是定义排序顺序的函数(闭包)。
	}
	sort.Sort(ps)
}

// planetSorter 将 By 函数和要排序的 Planets 切片合并起来。
type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool // Less方法中使用的闭包
}

// Len是sort.Interface的一部分。
func (s *planetSorter) Len() int {
	return len(s.planets)
}

// Swap是sort.Interface的一部分。
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less 是 sort.Interface 的一部分。它通过在 sorter 中调用 "by" 闭包来实现。
func (s *planetSorter) Less(i, j int) bool {
	return s.by(&s.planets[i], &s.planets[j])
}

var planets = []Planet{
	{"Mercury", 0.055, 0.4},
	{"Venus", 0.815, 0.7},
	{"Earth", 1.0, 1.0},
	{"Mars", 0.107, 1.5},
}

func main() {
    // 对Planet结构进行排序的闭包。
	name := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}
	mass := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}
	distance := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}
	decreasingDistance := func(p1, p2 *Planet) bool {
		return distance(p2, p1)
	}

    // 按照各种标准对planets进行排序。
	By(name).Sort(planets)
	fmt.Println("By name:", planets)

	By(mass).Sort(planets)
	fmt.Println("By mass:", planets)

	By(distance).Sort(planets)
	fmt.Println("By distance:", planets)

	By(decreasingDistance).Sort(planets)
	fmt.Println("By decreasing distance:", planets)

}
Output:

By name: [{Earth 1 1} {Mars 0.107 1.5} {Mercury 0.055 0.4} {Venus 0.815 0.7}]
By mass: [{Mercury 0.055 0.4} {Mars 0.107 1.5} {Venus 0.815 0.7} {Earth 1 1}]
By distance: [{Mercury 0.055 0.4} {Venus 0.815 0.7} {Earth 1 1} {Mars 0.107 1.5}]
By decreasing distance: [{Mars 0.107 1.5} {Earth 1 1} {Venus 0.815 0.7} {Mercury 0.055 0.4}]
```

#### Example(SortMultiKeys)

ExampleMultiKeys demonstrates a technique for sorting a struct type using different sets of multiple fields in the comparison. We chain together "Less" functions, each of which compares a single field.

​	ExampleMultiKeys演示了一种在比较中使用不同的多字段集对结构类型进行排序的技术。我们将 "Less"函数串联起来，每个函数都对一个字段进行比较。

``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

//  Change是一个记录源代码更改的记录，记录用户、语言和增量大小。
type Change struct {
	user     string
	language string
	lines    int
}

type lessFunc func(p1, p2 *Change) bool

// multiSorter实现了Sort接口，对其中的变化进行排序。
type multiSorter struct {
	changes []Change
	less    []lessFunc
}

// Sort根据传递给OrderedBy的less函数对参数片进行排序。
func (ms *multiSorter) Sort(changes []Change) {
	ms.changes = changes
	sort.Sort(ms)
}

// OrderedBy返回一个Sorter，该Sorter使用less函数依次进行排序。
//调用其Sort方法对数据进行排序。
func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

// Len方法是sort.Interface的一部分。
func (ms *multiSorter) Len() int {
	return len(ms.changes)
}

// Swap方法是sort.Interface的一部分。
func (ms *multiSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

// Less是sort.Interface的一部分。
// 它通过循环遍历less函数来实现，
// 直到找到区分两个项(一个小于另一个)的比较为止。
// 请注意，它可能会在每次调用时调用less函数两次。
// 我们可以更改函数返回-1、0、1并减少调用次数以提高效率：
// 这是reader的一个练习。
	
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.changes[i], &ms.changes[j]
    // 尝试除最后一个比较之外的所有比较。
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
            // p < q，因此我们做出了决定。
			return true
		case less(q, p):
            // p > q，因此我们做出了决定。
			return false
		}
        // p == q; 尝试下一个比较。
	}
    // 到这里为止的所有比较都说 "相等"，因此只需返回最终比较报告的任何内容。
	return ms.less[k](p, q)
}

var changes = []Change{
	{"gri", "Go", 100},
	{"ken", "C", 150},
	{"glenda", "Go", 200},
	{"rsc", "Go", 200},
	{"r", "Go", 100},
	{"ken", "Go", 200},
	{"dmr", "C", 100},
	{"r", "C", 150},
	{"gri", "Smalltalk", 80},
}

// ExampleMultiKeys演示了一种在比较中使用不同的多字段集对结构类型进行排序的技术。我们把 "Less "函数连在一起，每个函数都比较一个字段。
// ExampleMultiKeys演示了一种使用不同的多个字段集在比较中对结构类型进行排序的技术。我们链接"Less"函数，每个函数都比较一个字段。
func main() {
    // 闭包函数，用于对Change结构排序。
	user := func(c1, c2 *Change) bool {
		return c1.user < c2.user
	}
	language := func(c1, c2 *Change) bool {
		return c1.language < c2.language
	}
	increasingLines := func(c1, c2 *Change) bool {
		return c1.lines < c2.lines
	}
	decreasingLines := func(c1, c2 *Change) bool {
		return c1.lines > c2.lines // 注意：>的顺序是向下的。
	}

    // 简单使用：按用户排序。
	OrderedBy(user).Sort(changes)
	fmt.Println("By user:", changes)

    // 更多示例。
	OrderedBy(user, increasingLines).Sort(changes)
	fmt.Println("By user,<lines:", changes)

	OrderedBy(user, decreasingLines).Sort(changes)
	fmt.Println("By user,>lines:", changes)

	OrderedBy(language, increasingLines).Sort(changes)
	fmt.Println("By language,<lines:", changes)

	OrderedBy(language, increasingLines, user).Sort(changes)
	fmt.Println("By language,<lines,user:", changes)

}
Output:

By user: [{dmr C 100} {glenda Go 200} {gri Go 100} {gri Smalltalk 80} {ken C 150} {ken Go 200} {r Go 100} {r C 150} {rsc Go 200}]
By user,<lines: [{dmr C 100} {glenda Go 200} {gri Smalltalk 80} {gri Go 100} {ken C 150} {ken Go 200} {r Go 100} {r C 150} {rsc Go 200}]
By user,>lines: [{dmr C 100} {glenda Go 200} {gri Go 100} {gri Smalltalk 80} {ken Go 200} {ken C 150} {r C 150} {r Go 100} {rsc Go 200}]
By language,<lines: [{dmr C 100} {ken C 150} {r C 150} {gri Go 100} {r Go 100} {glenda Go 200} {ken Go 200} {rsc Go 200} {gri Smalltalk 80}]
By language,<lines,user: [{dmr C 100} {ken C 150} {r C 150} {gri Go 100} {r Go 100} {glenda Go 200} {ken Go 200} {rsc Go 200} {gri Smalltalk 80}]
```

#### Example (SortWrapper)
``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

type Grams int

func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }

type Organ struct {
	Name   string
	Weight Grams
}

type Organs []*Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByName实现了sort.Interface接口，
提供Less方法并使用嵌入的Organs值的Len和Swap方法。
type ByName struct{ Organs }

func (s ByName) Less(i, j int) bool { return s.Organs[i].Name < s.Organs[j].Name }

// ByWeight实现了sort.Interface接口，
// 提供Less方法并使用嵌入的Organs值的Len和Swap方法。
type ByWeight struct{ Organs }

func (s ByWeight) Less(i, j int) bool { return s.Organs[i].Weight < s.Organs[j].Weight }

func main() {
	s := []*Organ{
		{"brain", 1340},
		{"heart", 290},
		{"liver", 1494},
		{"pancreas", 131},
		{"prostate", 62},
		{"spleen", 162},
	}

	sort.Sort(ByWeight{s})
	fmt.Println("Organs by weight:")
	printOrgans(s)

	sort.Sort(ByName{s})
	fmt.Println("Organs by name:")
	printOrgans(s)

}

func printOrgans(s []*Organ) {
	for _, o := range s {
		fmt.Printf("%-8s (%v)\n", o.Name, o.Weight)
	}
}
Output:

Organs by weight:
prostate (62g)
pancreas (131g)
spleen   (162g)
heart    (290g)
brain    (1340g)
liver    (1494g)
Organs by name:
brain    (1340g)
heart    (290g)
liver    (1494g)
pancreas (131g)
prostate (62g)
spleen   (162g)
```


## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [Find](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/search.go;l=99)  <- go1.19

``` go linenums="1"
func Find(n int, cmp func(int) int) (i int, found bool)
```

​	Find函数使用二分查找来查找并返回[0,`n`)中cmp(i)<= 0的最小索引i。如果没有这样的索引i,则Find返回i = n。如果i < n并且cmp(i)== 0,则found结果为true。Find函数仅为范围[0，n)中的i调用cmp(i)。

​	为了允许二分查找，Find要求cmp(i)> 0为范围的前导前缀，cmp(i)== 0为中间，cmp(i)< 0为范围的最终后缀。(每个子范围可能为空。)建立此条件的常规方法是将cmp(i)解释为所需目标值t与基础索引数据结构x中的条目i进行比较，当t < x [i]，t == x [i]和t> x [i]时，分别返回<0，0和> 0。

​	例如，在已排序的随机访问字符串列表中查找特定字符串的方法如下：

```go linenums="1"
i, found := sort.Find(x.Len(), func(i int) int {
    return strings.Compare(target, x.At(i))
})
if found {
    fmt.Printf("found %s at entry %d\n", target, i)
} else {
    fmt.Printf("%s not found, would insert at %d", target, i)
}
```

#### func [Float64s](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=161) 

``` go linenums="1"
func Float64s(x []float64)
```

​	Float64s函数以递增顺序对float64s切片进行排序。非数字(NaN)值在其他值之前排序。

##### Float64s Example
``` go linenums="1"
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(s)
	fmt.Println(s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	sort.Float64s(s)
	fmt.Println(s)

}
Output:

[-3.8 -1.3 0.7 2.6 5.2]
[NaN -Inf 0 +Inf]
```

#### func [Float64sAreSorted](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=171) 

``` go linenums="1"
func Float64sAreSorted(x []float64) bool
```

​	Float64sAreSorted函数报告切片x是否按递增顺序排序，其中非数字(NaN)值在任何其他值之前。

##### Float64sAreSorted Example
``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []float64{0.7, 1.3, 2.6, 3.8, 5.2} // sorted ascending
	fmt.Println(sort.Float64sAreSorted(s))

	s = []float64{5.2, 3.8, 2.6, 1.3, 0.7} // sorted descending
	fmt.Println(sort.Float64sAreSorted(s))

	s = []float64{5.2, 1.3, 0.7, 3.8, 2.6} // unsorted
	fmt.Println(sort.Float64sAreSorted(s))

}
Output:

true
false
false
```

#### func [Ints](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=157) 

``` go linenums="1"
func Ints(x []int)
```

​	Ints函数按递增顺序对int的切片进行排序。

##### Ints Example
``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println(s)
}
Output:

[1 2 3 4 5 6]
```

#### func [IntsAreSorted](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=167) 

``` go linenums="1"
func IntsAreSorted(x []int) bool
```

​	IntsAreSorted函数报告切片x是否按递增顺序排序。

##### IntsAreSorted Example
``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6} // sorted ascending
	fmt.Println(sort.IntsAreSorted(s))

	s = []int{6, 5, 4, 3, 2, 1} // sorted descending
	fmt.Println(sort.IntsAreSorted(s))

	s = []int{3, 2, 4, 1, 5} // unsorted
	fmt.Println(sort.IntsAreSorted(s))

}
Output:

true
false
false
```

#### func [IsSorted](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=99) 

``` go linenums="1"
func IsSorted(data Interface) bool
```

​	IsSorted函数判断 data 是否已经排好序。

#### func [Search](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/search.go;l=58) 

``` go linenums="1"
func Search(n int, f func(int) bool) int
```

​	Search函数使用二分查找在 [0,n) 区间内查找并`返回`满足 f(i) 为 true 的`最小索引 i`，假设在范围 [0,n) 上 f(i) == true 意味着 f(i+1) == true。也就是说，Search 要求在输入区间 [0,n) 上 f 在某些(可能为空)前缀处为 false，而在剩余部分为 true；Search 返回第一个为 true 的索引。如果没有这样的索引，则 Search 返回 n。注意，与 strings.Index 等函数不同，"未找到"返回值不是 -1。Search 只对 [0,n) 区间内的 i 调用 f(i)。

​	Search函数的一个常见用法是在排序的可索引数据结构(如数组或切片)中查找值 x 的索引 i。在这种情况下，参数 f，通常是一个闭包，捕获要搜索的值以及数据结构的索引和排序方式。

​	例如，给定升序排列的切片 data，调用 Search(len(data), func(i int) bool { return data[i] >= 23 }) 返回最小的索引 i，使得 data[i] >= 23。如果调用者想要查找 23 是否在切片中，则必须单独测试 data[i] == 23。

​	在降序排列的数据中搜索需要使用 <= 操作符而不是 >= 操作符。

​	为了完成上面的例子，以下代码尝试在升序排列的整数切片 data 中查找值 x：

```go linenums="1"
x := 23
i := sort.Search(len(data), func(i int) bool { return data[i] >= x })
if i < len(data) && data[i] == x {
	// x存在于data[i]中
} else {
	// x 不在 data 中，但是 i 是可以插入 x 的索引
}
```

​	作为一个更为玩乐的例子，以下程序可以猜测你选的数字：

``` go linenums="1"
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
```

##### Search Example

​	此示例演示如何搜索按升序排序的列表。

``` go linenums="1" hl_lines="11 11"
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}
Output:

found 6 at index 2 in [1 3 6 10 15 21 28 36 45 55]
```

##### Search Example(DescendingOrder)

​	此示例演示如何搜索按降序排序的列表。这种方法与按升序搜索列表相同，但条件颠倒了。

``` go linenums="1" hl_lines="11 11"
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{55, 45, 36, 28, 21, 15, 10, 6, 3, 1}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] <= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}
Output:

found 6 at index 7 in [55 45 36 28 21 15 10 6 3 1]
```

#### func [SearchFloat64s](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/search.go;l=131) 

``` go linenums="1"
func SearchFloat64s(a []float64, x float64) int
```

​	SearchFloat64s函数在已排序的 float64s 切片中搜索 x 并返回Search指定的索引。如果 x 不在 a 中，则返回可以插入 x 的索引(可能为 len(a))。切片必须升序排列。

##### SearchFloat64s Example

​	这个例子演示了如何在升序排序的列表中搜索 float64。

``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []float64{1.0, 2.0, 3.3, 4.6, 6.1, 7.2, 8.0}

	x := 2.0
	i := sort.SearchFloat64s(a, x)
	if a[i] == x {
		fmt.Printf("found %g at index %d in %v\n", x, i, a)
	}

	x = 0.5
	i = sort.SearchFloat64s(a, x)
	if a[i] != x {
		fmt.Printf("%g not found, can be inserted at index %d in %v\n", x, i, a)
	}
}

Output:

found 2 at index 1 in [1 2 3.3 4.6 6.1 7.2 8]
0.5 not found, can be inserted at index 0 in [1 2 3.3 4.6 6.1 7.2 8]
```

> `%g` 格式化占位符用于格式化浮点数，可以自动选择使用 %e 或 %f 进行格式化，具体是使用哪一种方式取决于数值的大小和精度。如果数值很大或者很小，将使用 %e 进行格式化，否则将使用 %f 进行格式化。例如，对于 12345.6789，使用 `%g` 进行格式化的结果可能是 `1.23456789e+04`。

#### func [SearchInts](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/search.go;l=123) 

``` go linenums="1"
func SearchInts(a []int, x int) int
```

​	SearchInts函数在一个已排序的整数切片中搜索x并返回其下标，如Search所定义的那样。如果x不存在，则返回插入x的下标(可能是len(a))。切片必须按升序排序。

##### SearchInts Example
``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3, 4, 6, 7, 8}

	x := 2
	i := sort.SearchInts(a, x)
	if a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	}

	x = 5
	i = sort.SearchInts(a, x)
	if a[i] != x {
		fmt.Printf("%d not found, can be inserted at index %d in %v\n", x, i, a)
	}
}

Output:

found 2 at index 1 in [1 2 3 4 6 7 8]
5 not found, can be inserted at index 4 in [1 2 3 4 6 7 8]
```

#### func [SearchStrings](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/search.go;l=139) 

``` go linenums="1"
func SearchStrings(a []string, x string) int
```

​	SearchStrings函数在已排序的字符串切片中搜索x并返回其下标，如Search所定义的那样。如果x不存在，则返回插入x的下标(可能是len(a))。切片必须按升序排序。

#### func [Slice](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/slice.go;l=21)  <- go1.8

``` go linenums="1"
func Slice(x any, less func(i, j int) bool)
```

​	Slice函数使用提供的less函数对切片x进行排序。如果x不是切片，则会引发panic。

​	排序不保证是稳定的：相等的元素可能会从它们的原始顺序中反转。对于稳定排序，请使用SliceStable函数。

​	less函数必须满足Interface类型的Less方法的相同要求。

##### Slice Example
``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
}
Output:

By name: [{Alice 55} {Bob 75} {Gopher 7} {Vera 24}]
By age: [{Gopher 7} {Vera 24} {Alice 55} {Bob 75}]
```

#### func [SliceIsSorted](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/slice.go;l=43)  <- go1.8

``` go linenums="1"
func SliceIsSorted(x any, less func(i, j int) bool) bool
```

​	SliceIsSorted函数报告切片x是否按照提供的less函数排序。如果x不是切片，则会引发panic。

#### func [SliceStable](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/slice.go;l=35)  <- go1.8

``` go linenums="1"
func SliceStable(x any, less func(i, j int) bool)
```

​	SliceStable函数使用提供的less函数对切片x进行排序，保持相等的元素的原始顺序。如果x不是切片，则会引发panic。

​	less函数必须满足Interface类型的Less方法的相同要求。

##### SliceStable Example
``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {

	people := []struct {
		Name string
		Age  int
	}{
		{"Alice", 25},
		{"Elizabeth", 75},
		{"Alice", 75},
		{"Bob", 75},
		{"Alice", 75},
		{"Bob", 25},
		{"Colin", 25},
		{"Elizabeth", 25},
	}

	// Sort by name, preserving original order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	// Sort by age preserving name order
	sort.SliceStable(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age,name:", people)

}
Output:

By name: [{Alice 25} {Alice 75} {Alice 75} {Bob 75} {Bob 25} {Colin 25} {Elizabeth 75} {Elizabeth 25}]
By age,name: [{Alice 25} {Bob 25} {Colin 25} {Elizabeth 25} {Alice 75} {Alice 75} {Bob 75} {Elizabeth 75}]
```

#### func [Sort](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=42) 

``` go linenums="1"
func Sort(data Interface)
```

​	Sort函数按照Less方法的规定，以升序对数据进行排序。它对data.Len进行一次调用以确定n，对data.Less和data.Swap进行`O(n*log(n))`次调用。排序不保证是稳定的。

#### func [Stable](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=207)  <- go1.2

``` go linenums="1"
func Stable(data Interface)
```

​	Stable函数按照 Less 方法定义的升序规则对 data 进行排序，同时保留相等元素的原始顺序。

​	它对 data.Len 进行一次调用来确定 n，对 data.Less 进行 `O(n*log(n))` 次调用，对 data.Swap 进行 `O(n*log(n)*log(n))` 次调用。

#### func [Strings](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=164) 

``` go linenums="1"
func Strings(x []string)
```

​	Strings函数按照升序规则对字符串切片 x 进行排序。

##### Strings Example
``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)
}
Output:

[Alpha Bravo Delta Go Gopher Grin]
```

#### func [StringsAreSorted](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=174) 

``` go linenums="1"
func StringsAreSorted(x []string) bool
```

​	StringsAreSorted函数报告字符串切片 x 是否按升序排列。

## 类型

### type [Float64Slice](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=123) 

``` go linenums="1"
type Float64Slice []float64
```

​	 Float64Slice 类型实现了 Interface 接口，适用于 []float64，按升序排列，其中 Not-a-Number(NaN)值排在其他值之前。

#### (Float64Slice) [Len](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=125) 

``` go linenums="1"
func (x Float64Slice) Len() int
```

​	Len方法返回切片长度。

#### (Float64Slice) [Less](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=133) 

``` go linenums="1"
func (x Float64Slice) Less(i, j int) bool
```

​	Less方法报告 x[i] 是否应该在 x[j] 之前，根据 sort.Interface 规定的要求排序。请注意，浮点数比较本身不是一个传递性关系：对于 Not-a-Number(NaN)值，它不会报告一致的排序。此 Less 方法实现将 NaN 值排在其他值之前，使用以下方式：

```
x[i] < x[j] || (math.IsNaN(x[i]) && !math.IsNaN(x[j]))
```

#### (Float64Slice) [Search](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/search.go;l=147) 

``` go linenums="1"
func (p Float64Slice) Search(x float64) int
```

​	Search 函数返回将 x 插入到已按升序排序的 float64 切片中的索引 i，以便使 slice[i-1] <= x <= slice[i]，或如果 x 大于所有 slice 值，返回len(slice)。它调用了 SearchFloat64s 函数来执行二分搜索。

##### Search Example

```go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := sort.Float64Slice{2.1, 3.2, 5.5, 6.6, 7.7, 8.8}
	target := 5.5

	// 在nums中搜索target
	index := nums.Search(target)

	if index < nums.Len() && nums[index] == target {
		fmt.Printf("%v 在 nums 中的索引为 %d\n", target, index)
	} else {
		fmt.Printf("%v 不在 nums 中，此时index=%d\n", target,index)
	}

	target = 5.0

	// 在nums中搜索target
	index = nums.Search(target)

	if index < nums.Len() && nums[index] == target {
		fmt.Printf("%v 在 nums 中的索引为 %d\n", target, index)
	} else {
		fmt.Printf("%v 不在 nums 中，此时index=%d\n", target,index)
	}

	target = 99.9
	// 在nums中搜索target
	index = nums.Search(target)

	if index < nums.Len() && nums[index] == target {
		fmt.Printf("%v 在 nums 中的索引为 %d\n", target, index)
	} else {
		fmt.Printf("%v 不在 nums 中，此时index=%d\n", target,index)
	}
}
Output:

5.5 在 nums 中的索引为 2
5 不在 nums 中，此时index=2
99.9 不在 nums 中，此时index=6
```



#### (Float64Slice) [Sort](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=142) 

``` go linenums="1"
func (x Float64Slice) Sort()
```

​	Sort方法是一个便捷方法：x.Sort() 调用 Sort(x)函数。

##### Sort Example

```go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	floats := []float64{5.5, 2.2, 9.9, 0.1, 8.8}
	fmt.Println("Before sorting:", floats)
	// 在Sort方法调用后，floats切片已按升序排序。
	sort.Float64Slice(floats).Sort()
	fmt.Println("After sorting:", floats)
}
Output:

Before sorting: [5.5 2.2 9.9 0.1 8.8]
After sorting: [0.1 2.2 5.5 8.8 9.9]
```



#### (Float64Slice) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=134) 

``` go linenums="1"
func (x Float64Slice) Swap(i, j int)
```

​	Swap 方法交换索引为 i 和 j 的元素。

##### Swap Example

```go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 定义一个包含浮点数的切片
	nums := []float64{2.4, 1.2, 3.6, 0.8}

	// 将切片转换为 Float64Slice 类型
	fSlice := sort.Float64Slice(nums)

	// 打印切片交换前的内容
	fmt.Println("Before swapping:", fSlice)

	// 交换第二个和第三个元素
	fSlice.Swap(1, 2)

	// 打印切片交换后的内容
	fmt.Println("After swapping:", fSlice)
}
Output:

Before swapping: [2.4 1.2 3.6 0.8]
After swapping: [2.4 3.6 1.2 0.8]
```



### type [IntSlice](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=112) 

``` go linenums="1"
type IntSlice []int
```

 	IntSlice类型将 `Interface 的方法`附加到 []int，按升序排列。

#### (IntSlice) [Len](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=114) 

``` go linenums="1"
func (x IntSlice) Len() int
```

#### (IntSlice) [Less](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=115) 

``` go linenums="1"
func (x IntSlice) Less(i, j int) bool
```

#### (IntSlice) [Search](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/search.go;l=144) 

``` go linenums="1"
func (p IntSlice) Search(x int) int
```

Search returns the result of applying SearchInts to the receiver and x.

搜索返回对接收器和x应用SearchInts的结果。

​	Search方法用于在已排序的int切片p中查找元素x，并返回x的索引值，如果x不存在，则返回x应该插入的位置(即第一个大于x的元素的索引)，p必须是按升序排序的。

#### (IntSlice) [Sort](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=119) 

``` go linenums="1"
func (x IntSlice) Sort()
```

​	Sort是一个方便的方法：x.Sort()调用Sort(x)函数。

​	Sort方法是IntSlice的方法，用于按升序对切片进行排序，使用方法为x.Sort()。

#### (IntSlice) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=116) 

``` go linenums="1"
func (x IntSlice) Swap(i, j int)
```

​	Swap方法是IntSlice的方法，用于交换切片x中索引为i和j的两个元素。

### type [Interface](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=14) 

``` go linenums="1"
type Interface interface {
    // Len方法返回集合中的元素数
	Len() int

	// Less函数判断索引i对应的元素是否应该排在索引j对应的元素前面。
    //
    // 如果Less(i, j)和Less(j, i)都为false，
    // 那么元素i和元素j被认为是相等的。
    // 在Sort中，相等的元素可能以任意顺序出现，
    // 而在Stable中，相等元素的原始顺序会被保留。
    //
    // Less函数必须描述一个可传递的排序关系:
    // - 如果Less(i, j)和Less(j, k)都为true，
    // 	 则Less(i, k)也必须为true。
    // - 如果Less(i, j)和Less(j, k)都为false，
    //   则Less(i, k)也必须为false。
    //
    // 需要注意的是，在涉及到非数字(NaN)值的情况下，
    // 浮点数比较(即对float32或float64值使用<运算符)
    // 不是可传递的排序关系。
    // 有关浮点数的正确实现，请参见Float64Slice.Less函数。
	Less(i, j int) bool

    // Swap方法交换索引i和j上的两个元素。
	Swap(i, j int)
}
```

​	Interface的一个实现可以通过这个包中的例程进行排序。这些方法通过整数索引来引用底层集合的元素。

> 注意，Less方法需要满足传递性条件：如果Less(i, j)和Less(j, k)都为真，则Less(i, k)也为真。如果Less(i, j)和Less(j, k)都为假，则Less(i, k)也为假。此外，当涉及到NaN值时，浮点数比较(即float32或float64值上的<运算符)不是传递性关系。对于浮点数值，需要使用正确的实现，例如Float64Slice.Less方法。

#### func [Reverse](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=94)  <- go1.1

``` go linenums="1"
func Reverse(data Interface) Interface
```

​	Reverse函数返回data的逆序。

##### Reverse Example
``` go linenums="1"
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}
Output:

[6 5 4 3 2 1]
```

### type [StringSlice](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=145) 

``` go linenums="1"
type StringSlice []string
```

​	StringSlice类型将`Interface的方法`附加到[]string上，按升序排序。

#### (StringSlice) [Len](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=147) 

``` go linenums="1"
func (x StringSlice) Len() int
```

​	Len方法返回StringSlice的元素数量。

#### (StringSlice) [Less](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=148) 

``` go linenums="1"
func (x StringSlice) Less(i, j int) bool
```

​	Less方法根据升序规则报告x[i]是否应在x[j]之前排序。

#### (StringSlice) [Search](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/search.go;l=150) 

``` go linenums="1"
func (p StringSlice) Search(x string) int
```

​	Search方法返回将SearchStrings应用于接收器和x的结果。

#### (StringSlice) [Sort](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=152) 

``` go linenums="1"
func (x StringSlice) Sort()
```

​	Sort方法是一个便捷方法：x.Sort() 调用Sort(x)函数。

#### (StringSlice) [Swap](https://cs.opensource.google/go/go/+/go1.20.1:src/sort/sort.go;l=149) 

``` go linenums="1"
func (x StringSlice) Swap(i, j int)
```

​	Swap方法交换索引i和j处的元素。