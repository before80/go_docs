+++
title = "切片"
date = 2024-08-19T19:26:59+08:00
weight = 20
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 说明
>
> ​	以下实例代码中：
>
> ​	`verbs`的定义是：`var verbs = []string{"T", "v", "#v"} `
>
> ​	`mfp`来自：`"github.com/before80/utils/mfp"`

## 格式化动词

```go
type Person struct {
    name string
    age  int8
}

newVerbs := []string{"T", "%v", "+v", "#v"}
sl115 := []int{1, 2, 3}
sl116 := []float32{1.1, 2.2, 3.3}
sl117 := []string{"A", "B", "C"}
sl118 := []Person{{"Alice", 12}, {"Bob", 28}}
mfp.PrintFmtValWithLC("sl115", sl115, newVerbs)
mfp.PrintFmtValWithLC("sl116", sl116, newVerbs)
mfp.PrintFmtValWithLC("sl117", sl117, newVerbs)
mfp.PrintFmtValWithLC("sl118", sl118, newVerbs)
```

```
sl115:  %T -> []int | %[1 2 3] -> %v | %+v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl116:  %T -> []float32 | %[1.1 2.2 3.3] -> %v | %+v -> [1.1 2.2 3.3] | %#v -> []float32{1.1, 2.2, 3.3} | len=3 | cap=3
sl117:  %T -> []string | %[A B C] -> %v | %+v -> [A B C] | %#v -> []string{"A", "B", "C"} | len=3 | cap=3
sl118:  %T -> []main.Person | %[{Alice 12} {Bob 28}] -> %v | %+v -> [{name:Alice age:12} {name:Bob age:28}] | %#v -> []main.Person{main.Person{name:"Alice", age:12}, main.Person{name:"Bob", age:28}} | len=2 | cap=2
```



## C创建

### 1 直接创建

```go
var sl1 []int
var sl2 []int = []int{1, 2, 3}
var sl3 = []int{1, 2, 3}
sl4 := []int{1, 2, 3}
mfp.PrintFmtValWithLC("sl1", sl1, verbs)
mfp.PrintFmtValWithLC("sl2", sl2, verbs)
mfp.PrintFmtValWithLC("sl3", sl3, verbs)
mfp.PrintFmtValWithLC("sl4", sl4, verbs
```

```
sl1:    %T -> []int | %v -> [] | %#v -> []int(nil) | len=0 | cap=0
sl2:    %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl3:    %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl4:    %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
```



### 2 基于数组创建

```go
a1 := [...]int{1, 2, 3, 4, 5, 6}
mfp.PrintFmtValWithLC("a1", a1, verbs)
sl5 := a1[:]
sl6 := a1[0:]
sl7 := a1[:len(a1)]
sl8 := a1[0:len(a1)]
//sl9 := a1[0:3:2] // 报错：invalid slice indices: 2 < 3
sl10 := a1[0:3:3]
sl11 := a1[0:3:4]
sl12 := a1[0:3:5]
sl13 := a1[0:3:6]
//sl14 := a1[0:3:7] // 报错：invalid argument: index 7 out of bounds [0:7]
mfp.PrintFmtValWithLC("sl5", sl5, verbs)
mfp.PrintFmtValWithLC("sl6", sl6, verbs)
mfp.PrintFmtValWithLC("sl7", sl7, verbs)
mfp.PrintFmtValWithLC("sl8", sl8, verbs)
//mfp.PrintFmtValWithLC("sl9", sl9, verbs)
mfp.PrintFmtValWithLC("sl10", sl10, verbs)
mfp.PrintFmtValWithLC("sl11", sl11, verbs)
mfp.PrintFmtValWithLC("sl12", sl12, verbs)
mfp.PrintFmtValWithLC("sl13", sl13, verbs)
//mfp.PrintFmtValWithLC("sl14", sl14, verbs)
```

```
a1:     %T -> [6]int | %v -> [1 2 3 4 5 6] | %#v -> [6]int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
sl5:    %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
sl6:    %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
sl7:    %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
sl8:    %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
sl10:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl11:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=4
sl12:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=5
sl13:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=6
```

​	由以上示例在使用 `a[low:high:max]`获取新切片时可以看出：`max >= high >= low` ；`max`不得大于底层数组的上边界所在的索引； 新切片的长度为`high - low`，而容量为`max - low`。


### 3 用make创建

```go
sl15 := make([]int, 3)
//sl16 := make([]int, 3, 2) // 报错：invalid argument: length and capacity swapped
sl17 := make([]int, 3, 3)
sl18 := make([]int, 3, 4)
mfp.PrintFmtValWithLC("sl15", sl15, verbs)
//mfp.PrintFmtValWithLC("sl16", sl16, verbs)
mfp.PrintFmtValWithLC("sl17", sl17, verbs)
mfp.PrintFmtValWithLC("sl18", sl18, verbs)
```

```
sl15:   %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=3
sl17:   %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=3
sl18:   %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=4
```



### 4 用new创建

```go
sl19 := *new([]int) // 注意此时 sl19 为空切片，其长度和容量都为0
mfp.PrintFmtValWithLC("sl19", sl19, verbs)
sl19 = append(sl19, 1)
mfp.PrintFmtValWithLC("sl19", sl19, verbs)
sl19 = append(sl19, 2)
mfp.PrintFmtValWithLC("sl19", sl19, verbs)
sl19 = append(sl19, 3)
mfp.PrintFmtValWithLC("sl19", sl19, verbs)
```

```
sl19:   %T -> []int | %v -> [] | %#v -> []int(nil) | len=0 | cap=0
sl19:   %T -> []int | %v -> [1] | %#v -> []int{1} | len=1 | cap=1
sl19:   %T -> []int | %v -> [1 2] | %#v -> []int{1, 2} | len=2 | cap=2
sl19:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=4
```

### 5 基于已有切片创建

```go
a2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
mfp.PrintFmtValWithLC("已有数组 a2", a2, verbs)

sl20 := a2[0:6]
mfp.PrintFmtValWithLC("已有切片 sl20", sl20, verbs)

sl21 := sl20[:]
sl22 := sl20[0:]
sl23 := sl20[:len(sl20)]
sl24 := sl20[:cap(sl20)]
sl25 := sl20[0:len(sl20)]
sl26 := sl20[0:cap(sl20)]
//sl27 := sl20[0:cap(sl20)+1] // 报错：panic: runtime error: slice bounds out of range [:11] with capacity 10
sl28 := sl20[1:3]
sl29 := sl20[1:4]
sl30 := sl20[2:4]
//sl31 := sl20[2:4:2] // 报错：invalid slice indices: 2 < 4
//sl32 := sl20[2:4:3] // 报错：invalid slice indices: 3 < 4
sl33 := sl20[2:4:4]
sl34 := sl20[2:4:5]
sl35 := sl20[2:4:6]
sl36 := sl20[2:4:7]

mfp.PrintFmtValWithLC("sl21=sl20[:]", sl21, verbs)
mfp.PrintFmtValWithLC("sl22=sl20[0:]", sl22, verbs)
mfp.PrintFmtValWithLC("sl23=sl20[:len(sl20)]", sl23, verbs)
mfp.PrintFmtValWithLC("sl24=sl20[:cap(sl20)]", sl24, verbs)
mfp.PrintFmtValWithLC("sl25=[0:len(sl20)]", sl25, verbs)
mfp.PrintFmtValWithLC("sl26=[0:cap(sl20)]", sl26, verbs)
//mfp.PrintFmtValWithLC("sl27=sl20[0:cap(sl20)+1]", sl27, verbs)
mfp.PrintFmtValWithLC("sl28=sl20[1:3]", sl28, verbs)
mfp.PrintFmtValWithLC("sl29=sl20[1:4]", sl29, verbs)
mfp.PrintFmtValWithLC("sl30=sl20[2:4]", sl30, verbs)
//mfp.PrintFmtValWithLC("sl31=sl20[2:4:2]", sl31, verbs)
//mfp.PrintFmtValWithLC("sl32=sl20[2:4:3]", sl32, verbs)
mfp.PrintFmtValWithLC("sl33=sl20[2:4:4]", sl33, verbs)
mfp.PrintFmtValWithLC("sl34=sl20[2:4:5]", sl34, verbs)
mfp.PrintFmtValWithLC("sl35=sl20[2:4:6]", sl35, verbs)
mfp.PrintFmtValWithLC("sl36=sl20[2:4:7]", sl36, verbs)
```

```
已有数组 a2:    %T -> [10]int | %v -> [1 2 3 4 5 6 7 8 9 10] | %#v -> [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} | len=10 | cap=10
已有切片 sl20:  %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
sl21=sl20[:]:   %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
sl22=sl20[0:]:  %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
sl23=sl20[:len(sl20)]:  %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
sl24=sl20[:cap(sl20)]:  %T -> []int | %v -> [1 2 3 4 5 6 7 8 9 10] | %#v -> []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} | len=10 | cap=10
sl25=[0:len(sl20)]:     %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
sl26=[0:cap(sl20)]:     %T -> []int | %v -> [1 2 3 4 5 6 7 8 9 10] | %#v -> []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} | len=10 | cap=10
sl28=sl20[1:3]:         %T -> []int | %v -> [2 3] | %#v -> []int{2, 3} | len=2 | cap=9
sl29=sl20[1:4]:         %T -> []int | %v -> [2 3 4] | %#v -> []int{2, 3, 4} | len=3 | cap=9
sl30=sl20[2:4]:         %T -> []int | %v -> [3 4] | %#v -> []int{3, 4} | len=2 | cap=8
sl33=sl20[2:4:4]:       %T -> []int | %v -> [3 4] | %#v -> []int{3, 4} | len=2 | cap=2
sl34=sl20[2:4:5]:       %T -> []int | %v -> [3 4] | %#v -> []int{3, 4} | len=2 | cap=3
sl35=sl20[2:4:6]:       %T -> []int | %v -> [3 4] | %#v -> []int{3, 4} | len=2 | cap=4
sl36=sl20[2:4:7]:       %T -> []int | %v -> [3 4] | %#v -> []int{3, 4} | len=2 | cap=5
```

​	由上面给出的示例代码中的`sl24`和`sl26`，我们可以知道`sl20`这个切片的底层数组实际上就是`a2`。同时`a2`也是`sl21`到`sl36`的底层数组。

​	由以上示例在使用 `sl[low:high:max]`获取新切片时可以看出：`max >= high >= low` ；`max`不得大于底层数组的上边界所在的索引； 新切片的长度为`high - low`，而容量为`max - low`。

## U修改

### 1 修改元素

```go
sl37 := []int{1, 2, 3}
mfp.PrintFmtValWithLC("sl37", sl37, verbs)
sl37[0] = 11
mfp.PrintFmtValWithLC("sl37", sl37, verbs)
sl37[len(sl37)-1] = 33
mfp.PrintFmtValWithLC("sl37", sl37, verbs)
// 修改不存在的元素
//sl37[3] = 4 // 报错：panic: runtime error: index out of range [3] with length 3
```

```
sl37:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl37:   %T -> []int | %v -> [11 2 3] | %#v -> []int{11, 2, 3} | len=3 | cap=3
sl37:   %T -> []int | %v -> [11 2 33] | %#v -> []int{11, 2, 33} | len=3 | cap=3
```



### 2 用整个切片赋值

```go
sl38 := []int{1, 2, 3}
mfp.PrintFmtValWithLC("1 sl38", sl38, verbs)
sl38 = []int{1, 2, 3, 4}
mfp.PrintFmtValWithLC("2 sl38", sl38, verbs)
sl38 = make([]int, 5, 10)
mfp.PrintFmtValWithLC("3 sl38", sl38, verbs)
sl38 = *new([]int)
mfp.PrintFmtValWithLC("4 sl38", sl38, verbs)
sl39 := []int{1, 2, 3, 4, 5, 6}
mfp.PrintFmtValWithLC("5 sl39", sl39, verbs)
sl38 = sl39
mfp.PrintFmtValWithLC("6 sl38", sl38, verbs)
```

```
1 sl38:         %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
2 sl38:         %T -> []int | %v -> [1 2 3 4] | %#v -> []int{1, 2, 3, 4} | len=4 | cap=4
3 sl38:         %T -> []int | %v -> [0 0 0 0 0] | %#v -> []int{0, 0, 0, 0, 0} | len=5 | cap=10
4 sl38:         %T -> []int | %v -> [] | %#v -> []int(nil) | len=0 | cap=0
5 sl39:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
6 sl38:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
```

### 3 插入

#### 3.1 使用slices.Insert函数

```go
fmt.Println("从go1.21版本开始才可以使用")

sl109 := make([]int, 2, 3)
sl109 = slices.Replace(sl109, 0, 2, []int{1, 2}...)
mfp.PrintFmtValWithLC("1 sl109", sl109, verbs)
sl109 = slices.Insert(sl109, 0, []int{11, 22, 33}...)
mfp.PrintFmtValWithLC("2 sl109", sl109, verbs)

sl110 := make([]int, 2)
sl110 = slices.Replace(sl110, 0, 2, []int{1, 2}...)
mfp.PrintFmtValWithLC("1 sl110", sl110, verbs)
sl110 = slices.Insert(sl110, 0, []int{11, 22}...)
mfp.PrintFmtValWithLC("2 sl110", sl110, verbs)

sl111 := make([]int, 2)
sl111 = slices.Replace(sl111, 0, 2, []int{1, 2}...)
mfp.PrintFmtValWithLC("1 sl111", sl111, verbs)
sl111 = slices.Insert(sl111, 0, []int{11, 22, 33}...)
mfp.PrintFmtValWithLC("2 sl111", sl111, verbs)
```

```
1 sl109:        %T -> []int | %v -> [1 2] | %#v -> []int{1, 2} | len=2 | cap=3
2 sl109:        %T -> []int | %v -> [11 22 33 1 2] | %#v -> []int{11, 22, 33, 1, 2} | len=5 | cap=6
1 sl110:        %T -> []int | %v -> [1 2] | %#v -> []int{1, 2} | len=2 | cap=2
2 sl110:        %T -> []int | %v -> [11 22 1 2] | %#v -> []int{11, 22, 1, 2} | len=4 | cap=4
1 sl111:        %T -> []int | %v -> [1 2] | %#v -> []int{1, 2} | len=2 | cap=2
2 sl111:        %T -> []int | %v -> [11 22 33 1 2] | %#v -> []int{11, 22, 33, 1, 2} | len=5 | cap=6
```

​	由以上示例，我们可以发现在使用`slices.Insert`函数可以一次性插入多个新的元素，并且新生成的切片会自动进行扩容。

#### 3.2 使用slices.Replace函数

```go
fmt.Println("从go1.21版本开始才可以使用")

sl112 := make([]int, 2, 3)
sl112 = slices.Replace(sl112, 0, 2, []int{1, 2}...)
mfp.PrintFmtValWithLC("1 sl112", sl112, verbs)
sl112 = slices.Replace(sl112, 0, 0, 11)
mfp.PrintFmtValWithLC("2 sl112", sl112, verbs)
sl112 = slices.Replace(sl112, 0, 0, 111)
mfp.PrintFmtValWithLC("3 sl112", sl112, verbs)
```

```
1 sl112:        %T -> []int | %v -> [1 2] | %#v -> []int{1, 2} | len=2 | cap=3
2 sl112:        %T -> []int | %v -> [11 1 2] | %#v -> []int{11, 1, 2} | len=3 | cap=3
3 sl112:        %T -> []int | %v -> [111 11 1 2] | %#v -> []int{111, 11, 1, 2} | len=4 | cap=6
```

​	由以上示例，我们可以发现在使用`slices.Replace`函数一次只能插入1个新的元素，并且新生成的切片会自动进行扩容。

### 4 替换

#### 4.1 使用for循环

```go
sl73 := make([]int, 6, 10)
mfp.PrintFmtValWithLC("1 sl73", sl73, verbs)
// 将 sl73[0]~sl73[6]依次替换为 1~6
for k, _ := range sl73 {
    if k <= 6 {
        sl73[k] = k + 1
    }
}
mfp.PrintFmtValWithLC("2 sl73", sl73, verbs)
```

```
1 sl73:         %T -> []int | %v -> [0 0 0 0 0 0] | %#v -> []int{0, 0, 0, 0, 0, 0} | len=6 | cap=10
2 sl73:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
```

#### 4.2 使用slices.Replace函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl74 := make([]int, 6, 10)
mfp.PrintFmtValWithLC("1 sl74", sl74, verbs)
sl74 = slices.Replace(sl74, 0, 6, []int{1, 2, 3, 4, 5, 6}...)
mfp.PrintFmtValWithLC("2 sl74", sl74, verbs)
sl74 = slices.Replace(sl74, 0, 1, 111)
mfp.PrintFmtValWithLC("3 sl74", sl74, verbs)
//sl74 = slices.Replace(sl74, 0, 7, []int{1, 2, 3, 4, 5, 6}...) // 报错：panic: runtime error: slice bounds out of range [7:6]
//mfp.PrintFmtValWithLC("4 sl74", sl74, verbs)
//sl74 = slices.Replace(sl74, 0, 7, []int{1, 2, 3, 4, 5, 6, 7}...) // 报错：panic: runtime error: slice bounds out of range [7:6]
//mfp.PrintFmtValWithLC("5 sl74", sl74, verbs)
```

```
1 sl74:         %T -> []int | %v -> [0 0 0 0 0 0] | %#v -> []int{0, 0, 0, 0, 0, 0} | len=6 | cap=10
2 sl74:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
3 sl74:         %T -> []int | %v -> [111 2 3 4 5 6] | %#v -> []int{111, 2, 3, 4, 5, 6} | len=6 | cap=10
```

​	这里的Replace函数的定义为`func Replace[S ~[]E, E any](s S, i, j int, v ...E) S`，结合以上示例，可以发现， `i`和`j` 的必须是在`[0, len(S)]` （包含`0`和`len(S)`）的范围内，否则报错，实际替换不会替换`j`处的元素值。

### 5 反转

#### 5.1 使用for循环

```go
func reverseSlice(slice []int) {
	length := len(slice)
    for i := 0; i < length/2; i++ {
        j := length - 1 - i
        slice[i], slice[j] = slice[j], slice[i]
    }
}
sl76 := []int{1, 2, 3, 4, 5, 6}
mfp.PrintFmtValWithLC("1 sl76", sl76, verbs)
reverseSlice(sl76)
mfp.PrintFmtValWithLC("2 sl76", sl76, verbs)
```

```
1 sl76:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
2 sl76:         %T -> []int | %v -> [6 5 4 3 2 1] | %#v -> []int{6, 5, 4, 3, 2, 1} | len=6 | cap=6
```



#### 5.2 使用slices.Reverse函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl77 := []int{1, 2, 3, 4, 5, 6}
mfp.PrintFmtValWithLC("1 sl77", sl77, verbs)
slices.Reverse(sl77)
mfp.PrintFmtValWithLC("2 sl77", sl77, verbs)
```

```
1 sl77:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
2 sl77:         %T -> []int | %v -> [6 5 4 3 2 1] | %#v -> []int{6, 5, 4, 3, 2, 1} | len=6 | cap=6
```

### 6 移除

#### 6.1 移除未使用的容量

​	这里使用了`slices.Clip`函数，需要注意，`Clip的返回值`才是移除未使用的容量后的切片。

```go
fmt.Println("使用slices.Clip函数")
fmt.Println("从go1.21版本开始才可以使用")
sl78 := make([]int, 3, 6)
mfp.PrintFmtValWithLC("1 sl78", sl78, verbs)
sl78 = slices.Clip(sl78)
mfp.PrintFmtValWithLC("2 sl78", sl78, verbs)
```

```
1 sl78:         %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=6
2 sl78:         %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=3
```



### 7 排序

#### 7.1 使用slices.Sort函数

```go
fmt.Println("从go1.21版本开始才可以使用")	
sl82 := []float64{0, 42.12, -10.123, 8, math.NaN()}
mfp.PrintFmtValWithL("1 sl82", sl82, verbs)
slices.Sort(sl82)
mfp.PrintFmtValWithL("2 sl82", sl82, verbs)

type Person struct {
    name string
    age  int8
}

sl83 := []Person{
    {"zlx2", 30},
    {"zlx1", 32},
    {"zlx3", 29},
}

mfp.PrintFmtValWithLC("1 sl83", sl83, verbs)
//slices.Sort(sl83) // 报错：Person does not satisfy cmp.Ordered (Person missing in ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string)
//mfp.PrintFmtValWithLC("2 sl83", sl83, verbs)

sl84 := []int{2, 4, 6, 8, 1, 3, 5, 7}
mfp.PrintFmtValWithLC("1 sl84", sl84, verbs)
slices.Sort(sl84)
mfp.PrintFmtValWithLC("2 sl84", sl84, verbs)

sl85 := make([]int, 3, 6)
sl85 = slices.Replace(sl85, 0, 3, []int{2, 1, 3}...)
mfp.PrintFmtValWithLC("1 sl85", sl85, verbs)
slices.Sort(sl85)
mfp.PrintFmtValWithLC("2 sl85", sl85, verbs)
```

```
1 sl82:         %T -> []float64 | %v -> [0 42.12 -10.123 8 NaN] | %#v -> []float64{0, 42.12, -10.123, 8, NaN} | len=5
2 sl82:         %T -> []float64 | %v -> [NaN -10.123 0 8 42.12] | %#v -> []float64{NaN, -10.123, 0, 8, 42.12} | len=5
1 sl83:         %T -> []main.Person | %v -> [{zlx2 30} {zlx1 32} {zlx3 29}] | %#v -> []main.Person{main.Person{name:"zlx2", age:30}, main.Person{name:"zlx1", age:32}, main.Person{name:"zlx3", age:29}} | len=3 | cap=3
1 sl84:         %T -> []int | %v -> [2 4 6 8 1 3 5 7] | %#v -> []int{2, 4, 6, 8, 1, 3, 5, 7} | len=8 | cap=8
2 sl84:         %T -> []int | %v -> [1 2 3 4 5 6 7 8] | %#v -> []int{1, 2, 3, 4, 5, 6, 7, 8} | len=8 | cap=8
1 sl85:         %T -> []int | %v -> [2 1 3] | %#v -> []int{2, 1, 3} | len=3 | cap=6
2 sl85:         %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=6
```

#### 7.2 使用slices.SortFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")	
type Person struct {
    name string
    age  int8
}

sl86 := []Person{
    {"zlx2", 30},
    {"zlx1", 32},
    {"zlx3", 29},
}

mfp.PrintFmtValWithLC("1 sl86", sl86, verbs)
slices.SortFunc(sl86, func(a, b Person) int {
    return cmp.Compare(a.age, b.age)
})
mfp.PrintFmtValWithLC("2 sl86", sl86, verbs)

sl88 := []Person{
    {"Gopher", 13},
    {"Alice", 55},
    {"Bob", 24},
    {"Alice", 20},
}
mfp.PrintFmtValWithLC("1 sl88", sl88, verbs)
slices.SortFunc(sl88, func(a, b Person) int {
    if n := cmp.Compare(a.name, b.name); n != 0 {
        return n
    }
    // 如果 name 字段的值相等，则继续按 age 字段进行排序
    return cmp.Compare(a.age, b.age)
})
mfp.PrintFmtValWithLC("2 sl88", sl88, verbs)
```

```
1 sl86:         %T -> []main.Person | %v -> [{zlx2 30} {zlx1 32} {zlx3 29}] | %#v -> []main.Person{main.Person{name:"zlx2", age:30}, main.Person{name:"zlx1", age:32}, main.Person{name:"zlx3", age:29}} | len=3 | cap=3
2 sl86:         %T -> []main.Person | %v -> [{zlx3 29} {zlx2 30} {zlx1 32}] | %#v -> []main.Person{main.Person{name:"zlx3", age:29}, main.Person{name:"zlx2", age:30}, main.Person{name:"zlx1", age:32}} | len=3 | cap=3
1 sl88: 	%T -> []main.Person | %v -> [{Gopher 13} {Alice 55} {Bob 24} {Alice 20}] | %#v -> []main.Person{main.Person{name:"Gopher", age:13}, main.Person{name:"Alice", age:55}, main.Person{name:"Bob", age:24}, main.Person{name:"Alice", age:20}} | len=4 | cap=4
2 sl88: 	%T -> []main.Person | %v -> [{Alice 20} {Alice 55} {Bob 24} {Gopher 13}] | %#v -> []main.Person{main.Person{name:"Alice", age:20}, main.Person{name:"Alice", age:55}, main.Person{name:"Bob", age:24}, main.Person{name:"Gopher", age:13}} | len=4 | cap=4
```

#### 7.3 使用slices.SortStableFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")	

sl89 := []Person{
    {"Gopher", 13},
    {"Alice", 55},
    {"Bob", 24},
    {"Alice", 30},
    {"Alice", 20},
}
mfp.PrintFmtValWithLC("1 sl89", sl89, verbs)
slices.SortStableFunc(sl89, func(a, b Person) int {
    return cmp.Compare(a.name, b.name)
})
mfp.PrintFmtValWithLC("2 sl89", sl89, verbs)
```

```
1 sl89:         %T -> []main.Person | %v -> [{Gopher 13} {Alice 55} {Bob 24} {Alice 30} {Alice 20}] | %#v -> []main.Person{main.Person{name:"Gopher", age:13}, main.Person{name:"Alice", age:55}, main.Person{name:"Bob", age:24}, main.Person{name:"Alice", age:30}, main.Person{name:"Alice", age:20}} | len=5 | cap=5
2 sl89:         %T -> []main.Person | %v -> [{Alice 55} {Alice 30} {Alice 20} {Bob 24} {Gopher 13}] | %#v -> []main.Person{main.Person{name:"Alice", age:55}, main.Person{name:"Alice", age:30}, main.Person{name:"Alice", age:20}, main.Person{name:"Bob", age:24}, main.Person{name:"Gopher", age:13}} | len=5 | cap=5
```



## A访问

### 1 直接访问指定索引下标的元素

```go
sl40 := []int{1, 2, 3}
fmt.Println("sl40[0]=", sl40[0])
fmt.Println("sl40[1]=", sl40[1])
fmt.Println("sl40[2]=", sl40[2])
```

```
sl40[0]= 1
sl40[1]= 2
sl40[2]= 3
```

### 2 遍历切片

​	通过遍历的方式访问所需索引下标或全部索引下标的元素：

```go
for k, v := range sl40 {
    if k%2 == 0 {
        fmt.Println(k, "->", v)
    }
}
mfp.PrintHr()
for k, v := range sl40 {
    fmt.Println(k, "->", v)
}
mfp.PrintHr()
IamNaN5 := math.NaN()
sl40x := []float64{0, 42.12, -10.123, 8, IamNaN5}
for k, v := range sl40x {
    fmt.Println(k, "->", v)
}
```

```
0 -> 1
2 -> 3
------------------
0 -> 1
1 -> 2
2 -> 3
------------------
0 -> 0
1 -> 42.12
2 -> -10.123
3 -> 8
4 -> NaN
```

### 3 复制切片

#### 使用copy函数

```go
slSrc43 := make([]int, 3, 6)
slSrc43 = slices.Replace(slSrc43, 0, 3, []int{1, 2, 3}...)
mfp.PrintFmtValWithLC("1 slSrc43", slSrc43, verbs)
slDst44 := make([]int, len(slSrc43))
mfp.PrintFmtValWithLC("2 slDst44", slDst44, verbs)

copy(slDst44, slSrc43) // func copy(dst []Type, src []Type) int
mfp.PrintFmtValWithLC("3 slDst44", slDst44, verbs)
slDst44[0] = 11
fmt.Println("slDst44[0] = 11 之后")
mfp.PrintFmtValWithLC("4 slDst43", slSrc43, verbs)
mfp.PrintFmtValWithLC("5 slDst44", slDst44, verbs)
slSrc43[1] = 22
fmt.Println("slSrc43[1] = 22 之后")
mfp.PrintFmtValWithLC("6 slDst43", slSrc43, verbs)
mfp.PrintFmtValWithLC("7 slDst44", slDst44, verbs)
```

```
1 slSrc43:      %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=6
2 slDst44:      %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=3
3 slDst44:      %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
slDst44[0] = 11 之后
4 slDst43:      %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=6
5 slDst44:      %T -> []int | %v -> [11 2 3] | %#v -> []int{11, 2, 3} | len=3 | cap=3
slSrc43[1] = 22 之后
6 slDst43:      %T -> []int | %v -> [1 22 3] | %#v -> []int{1, 22, 3} | len=3 | cap=6
7 slDst44:      %T -> []int | %v -> [11 2 3] | %#v -> []int{11, 2, 3} | len=3 | cap=3
```

​	可见，`copy`函数复制后产生的切片和源切片不共用同一个底层数组！

#### 3.1 使用slices.Clone函数

```go
fmt.Println("从go1.21版本开始才可以使用")	
sl107 := make([]int, 3, 6)
sl107 = slices.Replace(sl107, 0, 3, []int{1, 2, 3}...)
mfp.PrintFmtValWithLC("1 sl107", sl107, verbs)
sl108 := slices.Clone(sl107)
mfp.PrintFmtValWithLC("2 sl108", sl108, verbs)
sl108[0] = 11
fmt.Println("sl108[0] = 11 之后")
mfp.PrintFmtValWithLC("3 sl107", sl107, verbs)
mfp.PrintFmtValWithLC("4 sl108", sl108, verbs)

sl107[1] = 22
fmt.Println("sl107[1] = 22 之后")
mfp.PrintFmtValWithLC("5 sl107", sl107, verbs)
mfp.PrintFmtValWithLC("6 sl108", sl108, verbs)
```

```
1 sl107:        %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=6
2 sl108:        %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl108[0] = 11 之后
3 sl107:        %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=6
4 sl108:        %T -> []int | %v -> [11 2 3] | %#v -> []int{11, 2, 3} | len=3 | cap=3
sl107[1] = 22 之后
5 sl107:        %T -> []int | %v -> [1 22 3] | %#v -> []int{1, 22, 3} | len=3 | cap=6
6 sl108:        %T -> []int | %v -> [11 2 3] | %#v -> []int{11, 2, 3} | len=3 | cap=3
```

​	通过以上示例，我们可以发现，`slices.Clone`函数并不会将源切片中的未使用的容量复制给新生成的切片，并且源切片和新生成的切片不共用同一个底层数组。

### 4 连接多个切片

#### 4.1 使用slices.Concat函数

```go
fmt.Println("从go1.22版本开始才可以使用")
sl94 := []int{1, 2, 3}
sl95 := []int{4, 5, 6}
sl96 := make([]int, 3, 6)
sl96 = slices.Replace(sl96, 0, 3, []int{7, 8, 9}...)
sl97 := make([]int, 3, 7)
sl97 = slices.Replace(sl97, 0, 3, []int{7, 8, 9}...)
sl98 := make([]int, 3, 8)
sl98 = slices.Replace(sl98, 0, 3, []int{7, 8, 9}...)
sl99 := slices.Concat(sl94, sl95, sl96)
sl100 := slices.Concat(sl94, sl95, sl97)
sl101 := slices.Concat(sl94, sl95, sl98)
mfp.PrintFmtValWithLC("sl94", sl94, verbs)
mfp.PrintFmtValWithLC("sl95", sl95, verbs)
mfp.PrintFmtValWithLC("sl96", sl96, verbs)
mfp.PrintFmtValWithLC("sl97", sl97, verbs)
mfp.PrintFmtValWithLC("sl98", sl98, verbs)
mfp.PrintFmtValWithLC("sl99", sl99, verbs)
mfp.PrintFmtValWithLC("sl100", sl100, verbs)
mfp.PrintFmtValWithLC("sl101", sl101, verbs)
```

```
从go1.22版本开始才可以使用
sl94:   %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
sl95:   %T -> []int | %v -> [4 5 6] | %#v -> []int{4, 5, 6} | len=3 | cap=3
sl96:   %T -> []int | %v -> [7 8 9] | %#v -> []int{7, 8, 9} | len=3 | cap=6
sl97:   %T -> []int | %v -> [7 8 9] | %#v -> []int{7, 8, 9} | len=3 | cap=7
sl98:   %T -> []int | %v -> [7 8 9] | %#v -> []int{7, 8, 9} | len=3 | cap=8
sl99:   %T -> []int | %v -> [1 2 3 4 5 6 7 8 9] | %#v -> []int{1, 2, 3, 4, 5, 6, 7, 8, 9} | len=9 | cap=10
sl100:  %T -> []int | %v -> [1 2 3 4 5 6 7 8 9] | %#v -> []int{1, 2, 3, 4, 5, 6, 7, 8, 9} | len=9 | cap=10
sl101:  %T -> []int | %v -> [1 2 3 4 5 6 7 8 9] | %#v -> []int{1, 2, 3, 4, 5, 6, 7, 8, 9} | len=9 | cap=10
```



### 5 获取相关切片属性

```go
sl41 := []int{1, 2, 3}
fmt.Println("sl41切片的长度 len(sl41)=", len(sl41))
fmt.Println("sl41切片的容量 cap(sl41)=", cap(sl41))
```

```
sl41切片的长度 len(sl41)= 3
sl41切片的容量 cap(sl41)= 3
```

### 6 获取索引

#### 6.1 使用slices.Index函数

```go
fmt.Println("从go1.21版本开始才可以使用")

sl113 := []string{"hello", "golang", "China", "World"}
fmt.Println("golang在sl113中的索引是 ", slices.Index(sl113, "golang"))
fmt.Println("China在sl113中的索引是 ", slices.Index(sl113, "China"))
fmt.Println("xyz在sl113中的索引是 ", slices.Index(sl113, "xyz"))
```

```
golang在sl113中的索引是  1
China在sl113中的索引是  2
xyz在sl113中的索引是  -1
```

​	需要注意的是，若指定的元素值并不在切片中，`slices.Index`函数返回的是`-1`。

#### 6.2 使用slices.BinarySearch函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl114 := []string{"hello", "golang", "China", "World"}
fmt.Println("未排序的sl114")
mfp.PrintFmtValWithLC("1 sl114", sl114, verbs)
i114, b114 := slices.BinarySearch(sl114, "golang")
fmt.Printf("golang 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)
i114, b114 = slices.BinarySearch(sl114, "China")
fmt.Printf("China 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)
i114, b114 = slices.BinarySearch(sl114, "xyz")
fmt.Printf("xyz 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)

mfp.PrintHr()
fmt.Println("已排序的sl114")
slices.Sort(sl114)
mfp.PrintFmtValWithLC("2 sl114", sl114, verbs)
i114, b114 = slices.BinarySearch(sl114, "golang")
fmt.Printf("golang 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)
i114, b114 = slices.BinarySearch(sl114, "China")
fmt.Printf("China 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)
i114, b114 = slices.BinarySearch(sl114, "xyz")
fmt.Printf("xyz 存在于sl114中？-> %t 所在索引是 %d\n", b114, i114)
```

```
未排序的sl114
1 sl114:        %T -> []string | %v -> [hello golang China World] | %#v -> []string{"hello", "golang", "China", "World"} | len=4 | cap=4
golang 存在于sl114中？-> false 所在索引是 4
China 存在于sl114中？-> false 所在索引是 0
xyz 存在于sl114中？-> false 所在索引是 4
------------------
已排序的sl114
2 sl114:        %T -> []string | %v -> [China World golang hello] | %#v -> []string{"China", "World", "golang", "hello"} | len=4 | cap=4
golang 存在于sl114中？-> true 所在索引是 2
China 存在于sl114中？-> true 所在索引是 0
xyz 存在于sl114中？-> false 所在索引是 4
```

​	由以上示例，我们可以发现在未排序的切片中使用`slices.BinarySearch`函数时，返回的结果都是不正确的！在查找不存在切片的中元素时，返回的索引是切片的长度，而非`-1`.

#### 6.3 使用slices.BinarySearchFunc函数

​	感觉不怎么实用，故未给出示例。

```go

```



### 7 判断是否相等

#### 7.1 是否可以使用`==`或`!=`?

​	=> 不可以！

```go
sl46 := []int{1, 2, 3}
sl47 := []int{1, 2, 3}
//fmt.Println("sl46 == sl47 -> ", sl46 == sl47) // 报错：invalid operation: sl46 == sl47 (slice can only be compared to nil)
//fmt.Println("sl46 != sl47 -> ", sl46 != sl47)// 报错：invalid operation: sl46 != sl47 (slice can only be compared to nil)
```

​	以上示例显示，在使用`==` 或 `!=` 时 切片 只可以和 `nil` 进行比较。

#### 7.2 使用slices.Equal函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl48 := []int{1, 2, 3}
sl49 := []int{1, 2, 3}
sl50 := []int{11, 2, 3}
sl51 := []int{1, 2, 3, 4}
sl48x1 := make([]int, 3, 6)
sl48x1 = slices.Replace(sl48x1, 0, 3, []int{1, 2, 3}...)
mfp.PrintFmtValWithLC("sl48x1", sl48x1, verbs)
fmt.Println("sl48 == sl49 -> ", slices.Equal(sl48, sl49))
fmt.Println("sl48 == sl50 -> ", slices.Equal(sl48, sl50))
fmt.Println("sl48 == sl51 -> ", slices.Equal(sl48, sl51))
fmt.Println("sl48 == sl48x1 -> ", slices.Equal(sl48, sl48x1))
```

```
sl48x1:         %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=6
sl48 == sl49 ->  true
sl48 == sl50 ->  false
sl48 == sl51 ->  false
sl48 == sl48x1 ->  true
```



#### 7.3 使用slices.EqualFunc函数

```go
sl52 := []int{1, 15, 8}
sl53 := []int{1, 15, 8}
sl54 := []int{11, 15, 8}
sl55 := []string{"01", "0x0f", "0o10"}
sl52x1 := make([]int, 3, 6)
sl52x1 = slices.Replace(sl52x1, 0, 3, []int{1, 15, 8}...)
mfp.PrintFmtValWithLC("sl52x1", sl52x1, verbs)
feq1 := func(e1, e2 int) bool {
    return e1 == e2
}
feq2 := func(e1 int, e2 string) bool {
    sn, err := strconv.ParseInt(e2, 0, 64)
    if err != nil {
        return false
    }
    return e1 == int(sn)
}
fmt.Println("sl52 == sl53 -> ", slices.EqualFunc(sl52, sl53, feq1))
fmt.Println("sl52 == sl54 -> ", slices.EqualFunc(sl52, sl54, feq1))
fmt.Println("sl52 == sl55 -> ", slices.EqualFunc(sl52, sl55, feq2))
fmt.Println("sl52 == sl52x1 -> ", slices.EqualFunc(sl52, sl52x1, feq1))
```

```
sl52x1:         %T -> []int | %v -> [1 15 8] | %#v -> []int{1, 15, 8} | len=3 | cap=6
sl52 == sl53 ->  true
sl52 == sl54 ->  false
sl52 == sl55 ->  true
sl52 == sl52x1 ->  true
```

#### 7.4 使用slices.Compare函数

```go
sl90 := []int{1, 2, 3}
sl91 := []int{1, 2, 3}
sl92 := []int{1, 2, 3, 4}
sl93 := []int{11, 2, 3}
fmt.Println("sl90 == sl91 ->", slices.Compare(sl90, sl91) == 0)
fmt.Println("sl90 == sl92 ->", slices.Compare(sl90, sl92) == 0)
fmt.Println("sl90 == sl93 ->", slices.Compare(sl90, sl93) == 0)
```

```
sl90 == sl91 -> true
sl90 == sl92 -> false
sl90 == sl93 -> false
```



### 8 判断是否存在

#### 	8.1 使用for循环

```go
sl56 := []int{1, 2, 3}
forFunc := func(src []int, target int) bool {
    for _, v := range src {
        if v == target {
            return true
        }
    }
    return false
}

fmt.Println("1 在 sl56中 -> ", forFunc(sl56, 1))
fmt.Println("4 在 sl56中 -> ", forFunc(sl56, 4))
```

```
1 在 sl56中 ->  true
4 在 sl56中 ->  false
```

#### 8.2 使用slices.Contains函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl57 := []int{1, 2, 3}
fmt.Println("1 在 sl57中 -> ", slices.Contains(sl57, 1))
fmt.Println("4 在 sl57中 -> ", slices.Contains(sl57, 4))
```

```
1 在 sl57中 ->  true
4 在 sl57中 ->  false
```

#### 8.3 使用slices.ContainsFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl58 := []int{0, 42, -10, 8}

fmt.Println("sl58中存在负数 -> ", slices.ContainsFunc(sl58, func(e int) bool {
    return e < 0
}))
fmt.Println("sl58中存在奇数 -> ", slices.ContainsFunc(sl58, func(e int) bool {
    return e%2 == 1
}))
fmt.Println("sl58中存在 8 -> ", slices.ContainsFunc(sl58, func(e int) bool {
    return e == 8
}))
```

```
sl58中存在负数 ->  true
sl58中存在奇数 ->  false
sl58中存在 8 ->  true
```

### 9 判断是否已排序

#### 9.1 使用slices.IsSorted函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl104 := []int{1, 2, 3}
sl105 := []int{1, 3, 2}
fmt.Println("sl104已排序？-> ", slices.IsSorted(sl104))
fmt.Println("sl105已排序？-> ", slices.IsSorted(sl105))
```

```
sl104已排序？->  true
sl105已排序？->  false
```

#### 9.2 使用slices.IsSortedFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl106 := []string{"alice", "Bob", "VERA"}
fmt.Println("sl106已排序？-> ", slices.IsSortedFunc(sl106, func(a, b string) int {
    return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
}))
```

```
sl106已排序？->  true
```



### 10 获取最大值

#### 10.1 使用for循环

```go
sl59 := []int{0, 42, -10, 8}

maxK := 0
maxV := sl59[0]
for k, v := range sl59 {
    if maxV < v {
        maxK = k
        maxV = v
    }
}
fmt.Printf("sl59中的最大值是sl59[%d]=%d\n", maxK, maxV)
```

```
sl59中的最大值是sl59[1]=42
```



#### 10.2 使用slices.Max函数

```go
fmt.Println("从go1.21版本开始才可以使用")

sl60 := []int{0, 42, -10, 8}
IamNaN := math.NaN()
sl61 := []float64{0, 42.12, -10.123, 8, IamNaN}
//sl62 := []int{0, 42, -10, 8, IamNaN} // 报错：cannot use IamNaN (variable of type float64) as int value in array or slice literal
fmt.Printf("sl60中的最大值是%d\n", slices.Max(sl60))

maxV2 := slices.Max(sl61)
fmt.Printf("sl61中的最大值是%f（%T）\n", maxV2, maxV2)
```

```
sl60中的最大值是42
sl61中的最大值是NaN（float64）
```



#### 10.3 使用slices.MaxFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl64 := []int{0, 42, -10, 8}
IamNaN2 := math.NaN()
sl65 := []float64{0, 42.12, -10.123, 8, IamNaN2}
fmt.Printf("sl64中最大值是%d\n", slices.MaxFunc(sl64, func(e1, e2 int) int {
    return cmp.Compare(e1, e2)
}))

fmt.Printf("sl65中最大值是%f\n", slices.MaxFunc(sl65, func(e1, e2 float64) int {
    return cmp.Compare(e1, e2)
}))

//sl66 := []int{}
//fmt.Printf("sl66中最大值是%d\n", slices.MaxFunc(sl66, func(e1, e2 int) int {
//	return cmp.Compare(e1, e2)
//})) // 报错：panic: slices.Max: empty list
```

```
sl64中最大值是42
sl65中最大值是42.120000
```



### 11 获取最小值

#### 11.1 使用for循环

```go
func findMin[T1, T2 cmp.Ordered](minK T1, minV T2, src []T2) (T1, T2) {
	for k, v := range src {
		if minV > v {
			minK = T1(k)
			minV = v
		}
	}
	return minK, minV
}

sl67 := []int{0, 42, -10, 8}
minK1, minV1 := findMin(0, sl67[0], sl67)
fmt.Printf("sl67中的最小值是sl67[%d]=%d\n", minK1, minV1)

sl68 := []float64{0, 42.12, -10.123, 8}
minK2, minV2 := findMin(0, sl68[0], sl68)
fmt.Printf("sl68中的最小值是sl68[%d]=%f\n", minK2, minV2)

IamNaN3 := math.NaN()
sl69 := []float64{0, 42.12, -10.123, 8, IamNaN3}
minK3, minV3 := findMin(0, sl69[0], sl69)
fmt.Printf("sl69中的最小值是sl69[%d]=%f\n", minK3, minV3)
```

```
sl67中的最小值是sl67[2]=-10
sl68中的最小值是sl68[2]=-10.123000
sl69中的最小值是sl69[2]=-10.123000
```

#### 11.2 使用slices.Min函数

```go
fmt.Println("从go1.21版本开始才可以使用")

sl70 := []int{0, 42, -10, 8}
sl71 := []float64{0, 42.12, -10.123, 8}
IamNaN4 := math.NaN()
sl72 := []float64{0, 42.12, -10.123, 8, IamNaN4}
fmt.Println("sl70中的最小值是", slices.Min(sl70))
fmt.Println("sl71中的最小值是", slices.Min(sl71))
fmt.Println("sl72中的最小值是", slices.Min(sl72))
```

```
sl70中的最小值是 -10
sl71中的最小值是 -10.123
sl72中的最小值是 NaN
```



#### 11.3 使用slices.MinFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")

sl70 := []int{0, 42, -10, 8}
sl71 := []float64{0, 42.12, -10.123, 8}
IamNaN4 := math.NaN()
sl72 := []float64{0, 42.12, -10.123, 8, IamNaN4}
fmt.Println("sl70中的最小值是", slices.MinFunc(sl70, func(a, b int) int {
	return cmp.Compare(a, b)
}))
fmt.Println("sl71中的最小值是", slices.MinFunc(sl71, func(a, b float64) int {
	return cmp.Compare(a, b)
}))
fmt.Println("sl72中的最小值是", slices.MinFunc(sl72, func(a, b float64) int {
	return cmp.Compare(a, b)
}))
```

```
sl70中的最小值是 -10
sl71中的最小值是 -10.123
sl72中的最小值是 NaN
```





#### 11.4 使用slices.Replace函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl74 := make([]int, 6, 10)
mfp.PrintFmtValWithLC("1 sl74", sl74, verbs)
sl74 = slices.Replace(sl74, 0, 6, []int{1, 2, 3, 4, 5, 6}...)
mfp.PrintFmtValWithLC("2 sl74", sl74, verbs)
//sl74 = slices.Replace(sl74, 0, 7, []int{1, 2, 3, 4, 5, 6}...) // 报错：panic: runtime error: slice bounds out of range [7:6]
//mfp.PrintFmtValWithLC("3 sl74", sl74, verbs)
//sl74 = slices.Replace(sl74, 0, 7, []int{1, 2, 3, 4, 5, 6, 7}...) // 报错：panic: runtime error: slice bounds out of range [7:6]
//mfp.PrintFmtValWithLC("4 sl74", sl74, verbs)
```

```
1 sl74:         %T -> []int | %v -> [0 0 0 0 0 0] | %#v -> []int{0, 0, 0, 0, 0, 0} | len=6 | cap=10
2 sl74:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=10
```



## D删除

### 1 是否可以删除某一元素呢？

​	=> 可以！

#### 1.1 使用append函数

```go
sl42 := []int{1, 2, 3, 4, 5, 6}
i := 3 // 需要删除元素的索引下标
mfp.PrintFmtValWithLC("1 sl42", sl42, verbs)
sl42 = append(sl42[0:i], sl42[i+1:]...) // 删除 索引为3的元素
mfp.PrintFmtValWithLC("2 sl42", sl42, verbs)
sl42 = append(sl42[0:i], sl42[i+1:]...) // 删除 当前索引为3的元素
mfp.PrintFmtValWithLC("3 sl42", sl42, verbs)
i = 0
sl42 = append(sl42[0:i], sl42[i+1:]...) // 删除 当前索引为0的元素
mfp.PrintFmtValWithLC("4 sl42", sl42, verbs)
sl42 = sl42[:len(sl42) - 1] // 删除当前的最后一个元素
mfp.PrintFmtValWithLC("5 sl42", sl42, verbs)
sl42 = sl42[1:] // 删除当前的第一个元素
mfp.PrintFmtValWithLC("6 sl42", sl42, verbs)
```

```
1 sl42:         %T -> []int | %v -> [1 2 3 4 5 6] | %#v -> []int{1, 2, 3, 4, 5, 6} | len=6 | cap=6
2 sl42:         %T -> []int | %v -> [1 2 3 5 6] | %#v -> []int{1, 2, 3, 5, 6} | len=5 | cap=6
3 sl42:         %T -> []int | %v -> [1 2 3 6] | %#v -> []int{1, 2, 3, 6} | len=4 | cap=6
4 sl42:         %T -> []int | %v -> [2 3 6] | %#v -> []int{2, 3, 6} | len=3 | cap=6
4 sl42:         %T -> []int | %v -> [2 3 6] | %#v -> []int{2, 3, 6} | len=3 | cap=6
5 sl42:         %T -> []int | %v -> [2 3] | %#v -> []int{2, 3} | len=2 | cap=6
6 sl42:         %T -> []int | %v -> [3] | %#v -> []int{3} | len=1 | cap=5
```

#### 1.2 使用slices.Delete函数

```go
fmt.Println("从go1.21版本开始才可以使用")	
sl87 := []int{1, 2, 3, 4, 5, 6}
mfp.PrintFmtValWithLC("1 sl87", sl87, verbs)
sl87 = slices.Delete(sl87, 0, 0) // 注意这里并没有删除成功
mfp.PrintFmtValWithLC("2 sl87", sl87, verbs)
sl87 = slices.Delete(sl87, 0, 1)// 这里才会删除成功
mfp.PrintFmtValWithLC("3 sl87", sl87, verbs)
```

```

```

### 1.3 是否可以批量删除一些元素

​	=> 可以 ！

​	可以使用`append`函数或`slices.Delete`函数来实现，具体代码参照删除某一元素的代码。

### 2 去重

#### 2.1 使用slices.Compact函数

> 注意
>
> ​	`slices.Compact`函数只能用于去除连续相等的元素。

```go
fmt.Println("从go1.21版本开始才可以使用")
sl102 := []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 8, 8, 1, 2, 3, 4, 5, 8}
mfp.PrintFmtValWithLC("1 sl102", sl102, verbs)
sl102 = slices.Compact(sl102)
mfp.PrintFmtValWithLC("2 sl102", sl102, verbs)
```

```
1 sl102:        %T -> []int | %v -> [0 1 1 2 2 3 3 4 4 5 5 8 8 1 2 3 4 5 8] | %#v -> []int{0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 8, 8, 1, 2, 3, 4, 5, 8} | len=19 | cap=19
2 sl102:        %T -> []int | %v -> [0 1 2 3 4 5 8 1 2 3 4 5 8] | %#v -> []int{0, 1, 2, 3, 4, 5, 8, 1, 2, 3, 4, 5, 8} | len=13 | cap=19
```

#### 2.2 使用slices.CompactFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")
sl103 := []string{"bob", "Bob", "alice", "Vera", "VERA"}
mfp.PrintFmtValWithLC("1 sl103", sl103, verbs)
sl103 = slices.CompactFunc(sl103, func(a, b string) bool {
    return strings.ToLower(a) == strings.ToLower(b)
})
mfp.PrintFmtValWithLC("2 sl103", sl103, verbs)
```

```
1 sl103:        %T -> []string | %v -> [bob Bob alice Vera VERA] | %#v -> []string{"bob", "Bob", "alice", "Vera", "VERA"} | len=5 | cap=5
2 sl103:        %T -> []string | %v -> [bob alice Vera] | %#v -> []string{"bob", "alice", "Vera"} | len=3 | cap=5
```



## 作为实参传递给函数或方法

​	在 Go 语言中，`切片是引用类型`。切片本身是一个包含指向底层数组的指针、长度和容量的数据结构。当你将一个切片赋值给另一个切片，或者将一个切片作为函数参数传递时，实际上是传递了切片的引用，而不是切片的副本。因此，对切片的修改会影响到原始切片以及引用同一底层数组的其他切片。

​	切片作为函数参数传递时，由于只是传递了切片的引用，而不是整个切片的副本，因此`在性能和内存上并不会有大开销`。即使切片的长度很大，传递切片的引用也只是传递了指向底层数组的指针、长度和容量这几个值，并不会复制整个底层数组。因此，切片作为实参通常不会产生额外的内存开销。

​	需要注意的是，如果在函数内部修改了切片的长度或容量，可能会导致底层数组重新分配内存，从而产生额外的内存开销。但是这种情况并不是切片本身作为实参引起的，而是对切片的修改引起的。

## 易混淆的知识点

​	

## 易错点

### 1 访问最后一个切片元素

​	直接用`sl[len(sl)]`访问切片`sl`的最后一个元素 => 肯定报错！

```go
sl45 := []int{1, 2, 3}
//fmt.Println(sl45[len(sl45)])   // 报错：panic: runtime error: index out of range [3] with length 3
fmt.Println(sl45[len(sl45)-1]) // 正确方式
```

```
3
```

### 2 长度和容量不一致的切片

​	长度和容量不一致时，给索引`i`的范围是`len(sl) <= i <= cap(sl)` 的元素赋值，以为可以增加切片的长度，实际却是 `panic`。 

```go
sl75 := make([]int, 3, 6)
mfp.PrintFmtValWithLC("1 sl75", sl75, verbs)
//sl75[3] = 4 // 报错：panic: runtime error: index out of range [3] with length 3
//mfp.PrintFmtValWithLC("2 sl75", sl75, verbs)
//sl75[4] = 5 // 报错：panic: runtime error: index out of range [4] with length 3
//mfp.PrintFmtValWithLC("3 sl75", sl75, verbs)
```

```
1 sl75:         %T -> []int | %v -> [0 0 0] | %#v -> []int{0, 0, 0} | len=3 | cap=6
```

### 3 使用slices.Replace函数

​	使用slices.Replace函数：`func Replace[S ~[]E, E any](s S, i, j int, v ...E) S ` 时，将 `i`和 `j`设置成一样，以为只会替换索引一处的元素值，而实际上却是往索引`i`前面插入一个新的元素值`v`。

```go
fmt.Println("错误的方式 1")
sl79 := []int{1, 2, 3}
mfp.PrintFmtValWithLC("1 sl79", sl79, verbs)
// 要修改索引0处的元素值
sl79 = slices.Replace(sl19, 0, 0, 111)
mfp.PrintFmtValWithLC("2 sl79", sl79, verbs)

fmt.Println("错误的方式 2")
sl81 := []int{1, 2, 3}
mfp.PrintFmtValWithLC("1 sl81", sl81, verbs)
fmt.Println("若 i == j == len(sl) 呢？")
sl81 = slices.Replace(sl81, 3, 3, 111)
mfp.PrintFmtValWithLC("2 sl81", sl81, verbs)

fmt.Println("正确的方式")
sl80 := []int{1, 2, 3}
mfp.PrintFmtValWithLC("1 sl80", sl80, verbs)
sl80 = slices.Replace(sl80, 0, 1, 111)
mfp.PrintFmtValWithLC("2 sl80", sl80, verbs)
```

```
错误的方式 1
1 sl79:         %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
2 sl79:         %T -> []int | %v -> [111 1 2 3] | %#v -> []int{111, 1, 2, 3} | len=4 | cap=4
错误的方式 2
1 sl81:         %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
若 i == j == len(sl) 呢？
2 sl81:         %T -> []int | %v -> [1 2 3 111] | %#v -> []int{1, 2, 3, 111} | len=4 | cap=6
正确的方式
1 sl80:         %T -> []int | %v -> [1 2 3] | %#v -> []int{1, 2, 3} | len=3 | cap=3
2 sl80:         %T -> []int | %v -> [111 2 3] | %#v -> []int{111, 2, 3} | len=3 | cap=3
```

