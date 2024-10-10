+++
title = "map"
date = 2024-07-13T14:05:04+08:00
weight = 600
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

## 关于map不得不知的知识点

1. map类型的默认值是`nil`；

  ```go
  var players map[string]int8
  // map[], map[string]int8, 0, true
  fmt.Printf("%[1]v, %[1]T, %d, %t\n", players, len(players), players == nil) 
  ```

2. 当从map中请求不存在的键时，返回的是map值类型的零值：

  ```go
  var players map[string]int8
  num := players["Durant"]
  fmt.Printf("num=%[1]v,%[1]T\n", num) // num=0,int8
  ```

3. 通过所接收的第二返回值，来判断所指定的键是否存在于map中：

  ```go
  players := map[string]int8{
      "Curry":  4,
      "LeBron": 6,
  }
  
  num, ok := players["Durant"]
  fmt.Println(num, ok) // 0 false
  num, ok = players["LeBron"]
  fmt.Println(num, ok) // 6 true
  ```

  

4. map类型是不可比较的类型，即不能将两个map类型的值（就算是这两个map类型值是同一变量，也不能）用在比较运算符中，否则编译报错，但map类型可以使用`==`比较运算符与`nil`进行比较，也仅仅可以用`==`比较运算符！

   ```go
   players1 := map[string]int8{
       "Curry":  4,
       "LeBron": 6,
   }
   
   players2 := map[string]int8{
       "Curry":  4,
       "LeBron": 6,
   }
   
   teams := map[string]string{
       "Warriors": "Golden State",
       "Lakers":   "Los Angeles",
   }
   
   fmt.Printf("%t\n", players1 == nil)                    // false
   fmt.Printf("%t\n", players1 == players1)               // invalid operation: players1 == players1 (map can only be compared to nil)
   fmt.Printf("%t\n", players1 == players2)               // invalid operation: players1 == players2 (map can only be compared to nil)
   fmt.Printf("%t\n", players1 == teams)                  // invalid operation: players1 == teams (mismatched types map[string]int8 and map[string]string)
   fmt.Printf("%t\n", players1 >= teams)                  // invalid operation: players1 == teams (mismatched types map[string]int8 and map[string]string)
   fmt.Printf("%t,%t\n", players1 > nil, players1 >= nil) // invalid operation: players1 > nil (operator > not defined on map) 以及 invalid operation: players1 >= nil (operator >= not defined on map)
   fmt.Printf("%t,%t\n", players1 < nil, players1 <= nil) // invalid operation: players1 > nil (operator < not defined on map) 以及 invalid operation: players1 <= nil (operator <= not defined on map)
   
   players1 = nil
   fmt.Printf("%t\n", players1 == nil)                    // true
   fmt.Printf("%t,%t\n", players1 > nil, players1 >= nil) // invalid operation: players1 > nil (operator > not defined on map) 以及 invalid operation: players1 >= nil (operator >= not defined on map)
   fmt.Printf("%t,%t\n", players1 < nil, players1 <= nil) // invalid operation: players1 > nil (operator < not defined on map) 以及 invalid operation: players1 <= nil (operator <= not defined on map)
   ```

   

5. map的键必须是可比较的类型（例如：布尔值、数字、字符串、指针、通道、由可比较类型组成的数组、字段均为可比较类型的结构体）；

6. 对未初始化的map进行键赋值操作，将引发运行时panic：

  ```go
  var players map[string]int8
  players["Curry"] = 4 // panic: assignment to entry in nil map
  ```

7. 使用map字面量时，每一行的键值对后面的`,`都是不能省略的(即使是最后一行也是如此)，否则编译时报错：

  ```go
  players = map[string]int8{
      "Curry":  4 // syntax error: unexpected newline in composite literal; possibly missing comma or }
      "LeBron": 6,
  }
  
  players = map[string]int8{
      "Curry":  4,
      "LeBron": 6  // syntax error: unexpected newline in composite literal; possibly missing comma or }
  }
  ```

  

8. 内置函数`delete`只能用于map类型，当map是`nil`或map中无任何键值对，delete相当于空操作（no-op）；

9. 内置函数`delete`用于map时，一次只能删除一个键；

10. 内置函数`cap`不能用于获取map的容量，理论上map的能够容纳无限个键值对，若使用`cap`对map进行操作，将导致编译错误：

  ```go
  fmt.Println(cap(players)) // invalid argument: players (variable of type map[string]int8) for cap
  ```

11. 内置函数`make`初始化map时，可以传递map的长度，但不可传递map的容量，否则将导致编译错误：

   ```go
   equipments := make(map[string]float64, 3) // invalid operation: make(map[string]float64, 3, 3) expects 1 or 2 arguments; found 3
   ```

12. 内置函数`make`初始化map后，该map非`nil`：

   ```go
   equipments := make(map[string]float64, 3)
   fmt.Printf("%[1]v, %[1]T, %d, %t\n", equipments, len(equipments), equipments == nil) // map[], map[string]float64, 0, false
   ```

   

13. 完全删除一个map中存储的值，可以直接赋值`nil`；

    





## 初始化map

### 使用短变量语法

```go
users := map[string]int8{
    "Curry": "4",
    "LeBron": "6",
    "Durant": "7"
}
```



### 使用make函数

## 未初始化的map



## map的键

​	必须是可比较的类型。

> 以下是comparable接口的相关信息：
> 
> ```go
> type comparable interface{ comparable } 
> ```
>
> comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types). The comparable interface may only be used as a type parameter constraint, not as the type of a variable.
>
> ​	 comparable是由所有可比较类型(布尔值、数字、字符串、指针、通道、由可比较类型组成的数组、字段均为可比较类型的结构体)实现的接口。`可比较接口只能用作类型参数约束，而不能作为变量类型`。

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

### 示例1

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

### 示例2

​	在函数间通过指针的方式传递sync.Map

```go
package main

import (
	"fmt"
	"sync"
)

func WriteMap(m *sync.Map) {
	m.Store(1, 1)
	m.Store(2, 2)
}

func ReadMap(m *sync.Map) {
	m.Range(func(k, v any) bool {
		kd, ok1 := k.(int)
		vd, ok2 := v.(int)
		if ok1 && ok2 {
			fmt.Println(kd, "->", vd)
		}
		return true
	})
}

func main() {
	var m sync.Map
	WriteMap(&m)
	ReadMap(&m)
}

```

怎么获取?
