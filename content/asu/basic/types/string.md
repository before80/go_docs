+++
title = "string"
date = 2024-07-13T14:04:51+08:00
weight = 300
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

## 字符串元素的类型是？

```go
package main

import "fmt"

func main() {
	s1 := "你好世界！你好中国"
	s := s1[0]

	fmt.Printf("%v,%c,%q,%b,%T,%T\n", s, s, s, s, s, &s)

	s2 := "Hello World! Hello China"
	s = s2[0]
	fmt.Printf("%v,%c,%q,%b,%T,%T\n", s, s, s, s, s, &s)
}
Output:
228,ä,'ä',11100100,uint8,*uint8
72,H,'H',1001000,uint8,*uint8

```

可见，字符串元素的类型却是`uint8` ，别名也就是`byte`！

## 反引号字符串中能引用到变量或常量吗？

```go
package main

import "fmt"

func main() {
	name1 := "zlx1"
	s1 := `你好，{name1}`
	s2 := `你好，${name1}`

	fmt.Println(name1)
	fmt.Println(s1)
	fmt.Println(s2)

	const name2 = "zlx2"
	s3 := `你好，{name2}`
	s4 := `你好，${name2}`

	fmt.Println(name2)
	fmt.Println(s3)
	fmt.Println(s4)
}

Output:
zlx1
你好，{name1}
你好，${name1}
zlx2
你好，{name2}
你好，${name2}

```

可见，反引号字符串中不能引用到变量或常量！

## 对字符串进行for range遍历

```go
package main

import "fmt"

func main() {
	s := "你好世界！你好中国"
	fmt.Printf("%2s %1s %4s %4s %5s %6s %4s\n", "i", "%c", "%x", "%X", "%v", "%U", "%q")
	for i, v := range s {
		fmt.Printf("%2d,%c,%x,%X,%v,%U,%q\n", i, v, v, v, v, v, v)
	}
}
Output:
 i %c   %x   %X    %v     %U   %q
 0,你,4f60,4F60,20320,U+4F60,'你'
 3,好,597d,597D,22909,U+597D,'好'
 6,世,4e16,4E16,19990,U+4E16,'世'
 9,界,754c,754C,30028,U+754C,'界'
12,！,ff01,FF01,65281,U+FF01,'！'
15,你,4f60,4F60,20320,U+4F60,'你'
18,好,597d,597D,22909,U+597D,'好'
21,中,4e2d,4E2D,20013,U+4E2D,'中'
24,国,56fd,56FD,22269,U+56FD,'国'
```

可见，对于非ASCII字符组成的字符串，其`i`的值并不是连续的！

## 对字符串进行len操作获取长度

```go
package main

import "fmt"

func main() {
	s := "你好世界！你好中国"
	fmt.Println(len(s)) // 27
}

```

可见，len() 所获取的长度实际是字符串中字节的个数！

## 对字符串使用简单的切片表达式

```go
package main

import "fmt"

func main() {
	s := "你好世界！你好中国"
	s1 := s[0:2]
	s2 := s[0:3]
	s3 := s[0:5]
	s4 := s[0:6]
	fmt.Println(s1) // ��
	fmt.Println(s2) // 你
	fmt.Println(s3) // 你��
	fmt.Println(s4) // 你好
}

```

可见，索引中的计数是按照字节来计算的，而非字符！

## 对字符串使用完整的切片扩展表达式

```go
package main

import "fmt"

func main() {
	s := "你好世界！你好中国"
	s1 := s[0:3:6] // 编译报错：invalid operation: 3-index slice of string
	fmt.Println(s1)
}

```

可见，完整的切片扩展表达式不能应用于字符串上！
