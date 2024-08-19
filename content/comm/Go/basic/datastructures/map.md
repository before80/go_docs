+++
title = "map"
date = 2024-08-19T19:27:33+08:00
weight = 1
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

## C创建

### 1 直接创建

```go
var m1 map[int]int
var m2 map[string]int = map[string]int{"A": 1, "B": 2}
var m3 = map[string]int{"A": 1, "B": 2}
m4 := map[string]int{"A": 1, "B": 2}
mfp.PrintFmtValWithL("m1", m1, verbs)
mfp.PrintFmtValWithL("m2", m2, verbs)
mfp.PrintFmtValWithL("m3", m3, verbs)
mfp.PrintFmtValWithL("m4", m4, verbs)
```

```
m1:     %T -> map[int]int | %v -> map[] | %#v -> map[int]int(nil) | len=0
m2:     %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
m3:     %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
m4:     %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
```



### 2 用make创建

```go
m5 := make(map[string]int)
m6 := make(map[string]int, 3)
//m7 := make(map[string]int, 3, 3) // 报错：invalid operation: make(map[string]int, 3, 3) expects 1 or 2 arguments; found 3
mfp.PrintFmtValWithL("1 m5", m5, verbs)
mfp.PrintFmtValWithL("2 m6", m6, verbs)
//mfp.PrintFmtValWithL("m7", m7, verbs)
```

```
1 m5:   %T -> map[string]int | %v -> map[] | %#v -> map[string]int{} | len=0
2 m6:   %T -> map[string]int | %v -> map[] | %#v -> map[string]int{} | len=0
```



### 3 用new创建

```go
m7 := *new(map[string]int)
mfp.PrintFmtValWithL("m7", m7, verbs)

//m7["A"] = 1 // 报错：panic: assignment to entry in nil map
//mfp.PrintFmtValWithL("m7", m7, verbs)

m7 = map[string]int{"A": 1}
mfp.PrintFmtValWithL("m7", m7, verbs)
```

```
m7:     %T -> map[string]int | %v -> map[] | %#v -> map[string]int(nil) | len=0
m7:     %T -> map[string]int | %v -> map[A:1] | %#v -> map[string]int{"A":1} | len=1
```



## U修改

### 1 修改元素

```
m9 := map[string]int{"A": 1, "B": 2, "C": 3}
mfp.PrintFmtValWithL("1 m9", m9, verbs)
m9["A"] = 11
mfp.PrintFmtValWithL("2 m9", m9, verbs)
m9["D"] = 4 // 修改不存在的Key
mfp.PrintFmtValWithL("3 m9", m9, verbs)
```

```
1 m9:   %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
2 m9:   %T -> map[string]int | %v -> map[A:11 B:2 C:3] | %#v -> map[string]int{"A":11, "B":2, "C":3} | len=3
3 m9:   %T -> map[string]int | %v -> map[A:11 B:2 C:3 D:4] | %#v -> map[string]int{"A":11, "B":2, "C":3, "D":4} | len=4
```



### 2 用整个map赋值

```go
m10 := map[string]int{"A": 1, "B": 2, "C": 3}
mfp.PrintFmtValWithL("1 m10", m10, verbs)
m10 = map[string]int{"A": 11, "B": 22, "C": 33, "D": 44}
mfp.PrintFmtValWithL("2 m10", m10, verbs)
m11 := map[string]int{"A": 111, "B": 222, "C": 333, "D": 444}
m10 = m11
mfp.PrintFmtValWithL("3 m10", m10, verbs)
m11["A"] = 1
mfp.PrintFmtValWithL("4 m10", m10, verbs)
```

```
1 m10:  %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
2 m10:  %T -> map[string]int | %v -> map[A:11 B:22 C:33 D:44] | %#v -> map[string]int{"A":11, "B":22, "C":33, "D":44} | len=4
3 m10:  %T -> map[string]int | %v -> map[A:111 B:222 C:333 D:444] | %#v -> map[string]int{"A":111, "B":222, "C":333, "D":444} | len=4
4 m10:  %T -> map[string]int | %v -> map[A:1 B:222 C:333 D:444] | %#v -> map[string]int{"A":1, "B":222, "C":333, "D":444} | len=4
```



## A访问

### 1 直接访问指定Key的元素

```go
m12 := map[string]int{"A": 1, "B": 2, "C": 3}
fmt.Println(m12["A"])
fmt.Println(m12["B"])
fmt.Println(m12["C"])
fmt.Println(m12["D"])// 访问不存在的Key
```

```
1
2
3
0
```



### 2 遍历map

```go
for k,v := range m12 {
    fmt.Println(k,"->", v)
}
```

```
A -> 1
B -> 2
C -> 3
```

​	需要注意的是，遍历是无序的，每一次的遍历顺序都有可能不同！

### 3 复制map

#### 3.1 使用maps.Clone函数

```go
fmt.Println("从go1.21版本开始才可以使用")

fmt.Println("使用maps.Clone函数")
m13 := map[string]int{"A": 1, "B": 2, "C": 3}
mfp.PrintFmtValWithL("1 m13", m13, verbs)
m14 := maps.Clone(m13)
mfp.PrintFmtValWithL("2 m14", m14, verbs)

m13["A"] = 11
fmt.Println(`修改 m13["A"] = 11`)
mfp.PrintFmtValWithL("3 m13", m13, verbs)
mfp.PrintFmtValWithL("4 m14", m14, verbs)

m14["B"] = 22
fmt.Println(`修改 m14["B"] = 22`)
mfp.PrintFmtValWithL("5 m13", m13, verbs)
mfp.PrintFmtValWithL("6 m14", m14, verbs)
```

```
从go1.21版本开始可使用
使用maps.Clone函数
1 m13:  %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
2 m14:  %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
修改 m13["A"] = 11
3 m13:  %T -> map[string]int | %v -> map[A:11 B:2 C:3] | %#v -> map[string]int{"A":11, "B":2, "C":3} | len=3
4 m14:  %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
修改 m14["B"] = 22
5 m13:  %T -> map[string]int | %v -> map[A:11 B:2 C:3] | %#v -> map[string]int{"A":11, "B":2, "C":3} | len=3
6 m14:  %T -> map[string]int | %v -> map[A:1 B:22 C:3] | %#v -> map[string]int{"A":1, "B":22, "C":3} | len=3
```

​	由以上示例，我们可以发现使用maps.Clone函数生成的新的map和源map在数据操作上互不影响。

#### 3.2 使用maps.Copy函数

```go
fmt.Println("从go1.21版本开始才可以使用")

m15 := map[string]int{"A": 1, "B": 2}
m16 := map[string]int{"A": 11, "C": 33}
fmt.Println(`使用Copy函数前`)
mfp.PrintFmtValWithL("m15", m15, verbs)
mfp.PrintFmtValWithL("m16", m16, verbs)
maps.Copy(m16, m15) // func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2)

fmt.Println(`使用Copy函数后`)
mfp.PrintFmtValWithL("m15", m15, verbs)
mfp.PrintFmtValWithL("m16", m16, verbs)

m15["A"] = 111
fmt.Println(`修改 m15["A"] = 111`)
mfp.PrintFmtValWithL("m15", m15, verbs)
mfp.PrintFmtValWithL("m16", m16, verbs)

m16["B"] = 222
fmt.Println(`修改 m16["B"] = 222`)
mfp.PrintFmtValWithL("m15", m15, verbs)
mfp.PrintFmtValWithL("m16", m16, verbs)
```

```
使用Copy函数前
m15:    %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
m16:    %T -> map[string]int | %v -> map[A:11 C:33] | %#v -> map[string]int{"A":11, "C":33} | len=2
使用Copy函数后
m15:    %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
m16:    %T -> map[string]int | %v -> map[A:1 B:2 C:33] | %#v -> map[string]int{"A":1, "B":2, "C":33} | len=3
修改 m15["A"] = 111
m15:    %T -> map[string]int | %v -> map[A:111 B:2] | %#v -> map[string]int{"A":111, "B":2} | len=2
m16:    %T -> map[string]int | %v -> map[A:1 B:2 C:33] | %#v -> map[string]int{"A":1, "B":2, "C":33} | len=3
修改 m16["B"] = 222
m15:    %T -> map[string]int | %v -> map[A:111 B:2] | %#v -> map[string]int{"A":111, "B":2} | len=2
m16:    %T -> map[string]int | %v -> map[A:1 B:222 C:33] | %#v -> map[string]int{"A":1, "B":222, "C":33} | len=3
```

​	由以上示例，我们可以发现使用maps.Copy函数后目的map和源map在数据操作上互不影响。

### 4 获取相关map属性

```go
fmt.Println("m12 map的长度 len(m12)=", len(m12))
```

```
m12 map的长度 len(m12)= 3
```

### 5 判断相等

#### 5.1 是否可以使用==或 !=？

​	=> 不可以！

```go
m18 := map[string]int{"A": 1, "B": 2, "C": 3}
m19 := map[string]int{"A": 1, "B": 2, "C": 3}
//fmt.Println("m18 == m19 -> ", m18 == m19) // 报错：invalid operation: m18 == m19 (map can only be compared to nil)
//fmt.Println("m18 != m19 -> ", m18 != m19) // 报错：invalid operation: m18 != m19 (map can only be compared to nil)
```

​	以上示例显示，在使用`==` 或 `!=` 时 map 只可以和 `nil` 进行比较。

#### 5.2 使用maps.Equal函数

```go
fmt.Println("从go1.21版本开始才可以使用")

m20 := map[string]int{"A": 1, "B": 2}
m21 := map[string]int{"A": 1, "B": 2}
fmt.Println("m20 == m21 ->", maps.Equal(m20, m21))

m22 := map[string]int{"A": 11, "B": 2}
fmt.Println("m20 == m22 ->", maps.Equal(m20, m22))

m23 := map[string]int{"A": 1, "B": 2, "C": 3}
fmt.Println("m20 == m23 ->", maps.Equal(m20, m23))
```

```
m20 == m21 -> true
m20 == m22 -> false
m20 == m23 -> false
```

#### 5.3 使用maps.EqualFunc函数

```go
fmt.Println("从go1.21版本开始才可以使用")
m24 := map[string]int{"A": 1, "B": 2}
m25 := map[string]int{"A": 1, "B": 2}
fmt.Println("m24 == m25 -> ", maps.EqualFunc(m24, m25, func(v1 int, v2 int) bool {
    if v1 == v2 {
        return true
    }
    return false
}))
```

```
m24 == m25 ->  true
```



## D删除

### 1 是否可以删除map中的某一元素？

​	=> 可以！

#### 1.1 使用delete函数

```go
m8 := map[string]int{"A": 1, "B": 2, "C": 3}
mfp.PrintFmtValWithL("m8", m8, verbs)
delete(m8, "A")
mfp.PrintFmtValWithL("m8", m8, verbs)
delete(m8, "A") // 重复删除，也不会报错
mfp.PrintFmtValWithL("m8", m8, verbs)
delete(m8, "B")
mfp.PrintFmtValWithL("m8", m8, verbs)
delete(m8, "C")
mfp.PrintFmtValWithL("m8", m8, verbs)
```

```
m8:     %T -> map[string]int | %v -> map[A:1 B:2 C:3] | %#v -> map[string]int{"A":1, "B":2, "C":3} | len=3
m8:     %T -> map[string]int | %v -> map[B:2 C:3] | %#v -> map[string]int{"B":2, "C":3} | len=2
m8:     %T -> map[string]int | %v -> map[B:2 C:3] | %#v -> map[string]int{"B":2, "C":3} | len=2
m8:     %T -> map[string]int | %v -> map[C:3] | %#v -> map[string]int{"C":3} | len=1
m8:     %T -> map[string]int | %v -> map[] | %#v -> map[string]int{} | len=0
```

#### 1.2 使用maps.DeleteFunc函数

```go
m17 := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}
fmt.Println("使用maps.DeleteFunc函数前")
mfp.PrintFmtValWithL("m17", m17, verbs)
maps.DeleteFunc(m17, func(k string, v int) bool {
    if v%2 == 1 {
        return true
    }
    return false
})

fmt.Println("使用maps.DeleteFunc函数后")
mfp.PrintFmtValWithL("m17", m17, verbs)
```

```
使用maps.DeleteFunc函数前
m17:    %T -> map[string]int | %v -> map[A:1 B:2 C:3 D:4] | %#v -> map[string]int{"A":1, "B":2, "C":3, "D":4} | len=4
使用maps.DeleteFunc函数后
m17:    %T -> map[string]int | %v -> map[B:2 D:4] | %#v -> map[string]int{"B":2, "D":4} | len=2
```



## 作为实参传递给函数或方法

​	在 Go 语言中，`map 是引用类型`。当你将一个 map 赋值给另一个变量，或者将一个 map 作为函数参数传递时，实际上是传递了 map 的引用，而不是整个 map 的副本。因此，对 map 的修改会影响到原始 map 以及引用同一个 map 的其他变量。

​	`map 作为函数参数传递时，并不会产生大的性能和内存开销`。与切片类似，虽然 map 可能包含大量的键值对，但传递 map 的引用只是传递了指向底层数据结构的指针，而不是复制整个底层数据结构。因此，map 作为实参传递通常不会产生额外的内存开销。

​	需要注意的是，在并发编程中，对 map 的并发访问可能会导致竞态条件，因此在多个 goroutine 中共享 map 时，需要使用适当的同步机制（例如 sync.Mutex 或 sync.RWMutex）来保护 map 的访问。

## 易混淆的知识点



## 易错点

### 1 直接对new函数创建的map进行key操作

​	=> 直接报错！

```go
m26 := *new(map[string]int)
mfp.PrintFmtValWithL("1 m26", m26, verbs)
//m26["A"] = 1 // 报错：panic: assignment to entry in nil map
m26 = map[string]int{"A": 1} // 正确方式
mfp.PrintFmtValWithL("2 m26", m26, verbs)
m26["B"] = 2
mfp.PrintFmtValWithL("3 m26", m26, verbs)
```

```
1 m26:  %T -> map[string]int | %v -> map[] | %#v -> map[string]int(nil) | len=0
2 m26:  %T -> map[string]int | %v -> map[A:1] | %#v -> map[string]int{"A":1} | len=1
3 m26:  %T -> map[string]int | %v -> map[A:1 B:2] | %#v -> map[string]int{"A":1, "B":2} | len=2
```

### 2 以为可以使用copy内置函数来复制一个map

```go
m27 := map[string]int{"A": 1}
m28 := make(map[string]int, 1)
//copy(m28, m27) // 报错：invalid argument: copy expects slice arguments; found m28 (variable of type map[string]int) and m27 (variable of type map[string]int)
```

