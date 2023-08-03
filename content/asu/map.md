+++
title = "map"
weight = 92
date = 2023-06-12T16:06:15+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# map

## 对原生map进行并发写会触发panic

```go
package main

import (
	"fmt"
	"time"
)

func WriteToMap(m map[int]int, j int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("w panic"+string(j)+":", r)
		}
	}()

	for i := 0; i <= 10000; i++ {
		m[i] = i + j
	}
}

func main() {
	m := make(map[int]int)

	go WriteToMap(m, 1)
	go WriteToMap(m, 2)
	go WriteToMap(m, 3)

	time.Sleep(20 * time.Second)
	fmt.Println(m)
}

会报类似以下的错误：
fatal error: concurrent map writes
fatal error: concurrent map writes

```

## 对原生map进行并发读却不会触发panic

```go
package main

import (
	"fmt"
	"time"
)

func WriteToMap(m map[int]int, j int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("w panic"+string(j)+":", r)
		}
	}()

	for i := 0; i <= 10000; i++ {
		m[i] = i + j
	}
}

func ReadFromMap(m map[int]int, j int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("r panic"+string(j)+":", r)
		}
	}()
	for i := 0; i <= 10000; i++ {
		fmt.Println("j=", j, "->读取m[", i, "]=", m[i])
	}
}

func main() {
	m := make(map[int]int)

	WriteToMap(m, 1)

	for i := 0; i <= 1000; i++ {
		go ReadFromMap(m, i)
	}

	time.Sleep(20 * time.Second)
}

```

## 对原生map进行并发读写会触发panic

```go
package main

import (
	"fmt"
	"time"
)

func WriteToMap(m map[int]int, j int) {
    defer func () {
        if r := recover(); r != nil {
		fmt.Println("w panic"+string(j)+":", r)
	}
    }()
	
	for i := 0; i <= 10000; i++ {
		m[i] = i + j
	}
}

func ReadFromMap(m map[int]int, j int) {
	if r := recover(); r != nil {
		fmt.Println("r panic"+string(j)+":", r)
	}
	for i := 0; i <= 10000; i++ {
		fmt.Println("j=", j, "->读取m[", i, "]=", m[i])
	}
}

func main() {
	m := make(map[int]int)

	go WriteToMap(m, 1)
	go WriteToMap(m, 2)
	go WriteToMap(m, 3)
	go WriteToMap(m, 4)

	go ReadFromMap(m, 1)
	go ReadFromMap(m, 2)
	go ReadFromMap(m, 3)
	go ReadFromMap(m, 4)

	time.Sleep(20 * time.Second)
}

会报类似以下的错误：
fatal error: concurrent map writes
fatal error: concurrent map read and map write
fatal error: concurrent map read and map write

```

## 对原生map的并发读写，有什么改进的方法？

-> 改用sync.Map。一方面在并发读写时不会触发panic，另一方面在读多写少的场景下性能优于原生map。

​	但，sync.Map 因实现采用两个原生的map来实现读写分离，会使用更多的内存对象，对于GC会产生更多压力，因此在内存紧缺或GC性能要求很高的系统应尽量避免使用sync.Map。

```go

```

​	怎么给出测试，说明性能优于原生map？

```

```



## 对原生nil map 进行写入会触发panic

```go
package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic", r)
		}
	}()
	var m map[int]int
	fmt.Println(m == nil)
	m[0] = 0
}
Output:
true
panic assignment to entry in nil map
```



## 使用sync.Map的示例

```go
package main

import (
	"fmt"
	"sync"
)

func GetType(d any) string {
	switch d.(type) {
	case int:
		return "int"
	case uint:
		return "uint"
	case string:
		return "string"
	case float32:
		return "float32"
	case float64:
		return "float64"
	default:
		return "未知类型"
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic", r)
		}
	}()
	var m sync.Map
	fmt.Printf("%T\n", m)  // sync.Map
	fmt.Printf("%v\n", m)  // {{0 0} {[] {} <nil>} map[] 0}
	fmt.Printf("%#v\n", m) // sync.Map{mu:sync.Mutex{state:0, sema:0x0}, read:atomic.Pointer[sync.readOnly]{_:[0]*sync.readOnly{}, _:atomic.noCopy{}, v:(unsafe.Pointer)(nil)}, dirty:map[interface {}]*sync.entry(nil), misses:0}

	// 添加 k=int,v=int 的键值对
	m.Store(0, 0)

	// 添加 k=int,v=string 的键值对
	m.Store(1, "1")

	// 添加 k=string,v=int 的键值对
	m.Store("2", 2)

	// 添加 k=string,v=string 的键值对
	m.Store("3", "3")

	// 查询
	d, ok := m.Load(0)
	//var dd int
	//dd = d // 编译报错：cannot use d (variable of type any) as int value in assignment: need type assertion
	//fmt.Println(dd)
	fmt.Printf("%v,%T,%t\n", d, d, ok) // 0,int,true
	d, ok = m.Load(1)
	fmt.Printf("%v,%T,%t\n", d, d, ok) // 1,string,true
	d, ok = m.Load("2")
	fmt.Printf("%v,%T,%t\n", d, d, ok) // 2,int,true
	d, ok = m.Load("3")
	fmt.Printf("%v,%T,%t\n", d, d, ok) // 3,string,true

	// 遍历
	m.Range(func(k, v any) bool {
		kt := GetType(k)
		vt := GetType(v)
		fmt.Println("键类型是", kt, ",值类型是", vt, " ->", k, ":", v)
		return true
	})

	// 删除
	// 删除存在的键
	m.Delete(0)

	// 删除不存在的键
	m.Delete("A")

	// 使用 CompareAndDelete 方法
	// 对于已经删除的键
	fmt.Println(m.CompareAndDelete(0, 1)) // false
	// 对于还存在的键，给出的旧值和存储的一致
	fmt.Println(m.CompareAndDelete(1, "1")) // true
	// 对于还存在的键，给出的旧值和存储的不一致
	fmt.Println(m.CompareAndDelete(1, "11")) // false

	// 对于还存在的键，给出的旧值和存储的不一致
	fmt.Println(m.CompareAndDelete("2", 22)) // false
	// 对于还存在的键，给出的旧值和存储的一致
	fmt.Println(m.CompareAndDelete("2", 2)) // true

	// 再次遍历，会发现只剩下： 一个键值对了
	m.Range(func(k, v any) bool {
		kt := GetType(k)
		vt := GetType(v)
		fmt.Println("键类型是", kt, ",值类型是", vt, " ->", k, ":", v)
		return true
	})

	// 使用 LoadOrStore  方法
	fmt.Println(m.LoadOrStore(0, 0))      // 0 false
	fmt.Println(m.LoadOrStore(1, "1"))    // 1 false
	fmt.Println(m.LoadOrStore("2", 22))   // 22 false
	fmt.Println(m.LoadOrStore("3", "33")) // 3 true <- 注意这里的 33 最终并没有被写入到m中

	// 再次遍历
	m.Range(func(k, v any) bool {
		kt := GetType(k)
		vt := GetType(v)
		fmt.Println("键类型是", kt, ",值类型是", vt, " ->", k, ":", v)
		return true
	})

	// 使用 LoadAndDelete  方法
	fmt.Println(m.LoadAndDelete("3")) // 3 true
	fmt.Println(m.LoadAndDelete("3")) // <nil> false

	// 使用 CompareAndSwap 方法
	fmt.Println(m.CompareAndSwap("3", "", "3")) // false
	fmt.Println(m.CompareAndSwap("3", "", "3")) // false
	fmt.Println(m.CompareAndSwap("2", 2, 22))   // false
	fmt.Println(m.CompareAndSwap("2", 22, 2))   // true

	// 再次遍历
	m.Range(func(k, v any) bool {
		kt := GetType(k)
		vt := GetType(v)
		fmt.Println("键类型是", kt, ",值类型是", vt, " ->", k, ":", v)
		return true
	})
}

```



怎么获取
