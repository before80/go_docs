+++
title = "todos"
date = 2023-05-27T08:36:49+08:00
description = ""
isCJKLanguage = true
draft = false
+++

# TODO 

2023年5月27日08:37:31 -> Go语言程序设计 - 英
1. io/ioutil包中的函数和方法可以使用什么来替换。
2. fmt.Println在打印错误时，是否会调用该类型中的`Error() string` 方法？
3. fmt.Println什么时候会调用该类型的`String() string`方法?
4. 哪些fmt函数会通过反射来调用`Error() string` 或 `String() string`方法？
5. 空标记符`_`始终都不是一个新的变量吗？P24 -> 是的，请参见 P36
6. 在同一个包同一个文件中，是否可以定义多个init函数？P30
7. 使用fmt.Println()没有任何实参，会打印吗？还是报错？P32
8. go 标识符可以使用多个_开头吗？-> 可以
9. 可以在32位系统上使用float64、int64、complex128？
10. const a = 1; const b = 1.2; var c = a + b; 请问是否会编译错误，c的值是多少？c的类型？-> 可以编译，c=2.2，c的类型是float64。
11. 比较运算符，用于判断两个接口变量时，这两个接口变量有什么限制？怎么进行比较？
12. 比较运算符，用于两种不同的类型的变量进行比较，编译会报错吗？可以用一个变量和常量进行比较吗？若变量的类型和常量的类型不一致，是否还可以进行比较？

2023年5月29日09:58:23

1. 浮点数计算精度问题的解决方法？

2. 浮点数比较存在的问题？

3. 浮点数避免除数为0的判断方法 `if y != 0.0 { ... }` 中的 0.0 是否可以替换成 0？P50

4. fmt.Printf()、fmt.Sprintf()中的格式化字符中使用了`*`，按照 [浮点数和复数成分]({{< ref "/stdLib/fmt#浮点数和复数成分">}}) 中描述的 会使用下一个操作数，但必须是`int`类型的值，请问，若传入是`int8`、`uint`等类型会发生什么错误？=> 不会出错！！

5. 使用`int[8|16|32] (float32或float64)`，若浮点数的值超出了目标整型的的范围，那么得到的结果值将是不可预期的！！！请给出示例 P50

6. 怎么理解go的复数类型是定长的？P53

7. 浮点数 `0.1 + 0.2 == 0.3` ？

8. ```go
   	a := 0.1
   	b := 0.2
   	c := 0.3
   
   	fmt.Println((0.1 + 0.2) == 0.3) // true
   	fmt.Println((a + b) == c) // false
   
   	var (
   		e float32 = 0.1
   		f float32 = 0.2
   		g float32 = 0.3
   	)
   
   	fmt.Println((e + f) == g) // true
   
   	var (
   		h float64 = 0.1
   		i float64 = 0.2
   		j float64 = 0.3
   	)
   
   	fmt.Println((h + i) == j) // false
   ```

   为什么一个是true，一个是false？

9. `chars := []rune` 可以使用 `string(chars)` 转换成字符串，相反，`s := string` 可以使用 `chars := []rune(s)`，将[]rune转换成字符串。那 若是自定义类型，是否还可以使用类似的方法进行相互转换，不行的情况下，需要怎么处理？ => 可以

   ```go
   	type MyStr string
   
   	var str MyStr = "中国a美国b"
   
   	chars := []rune(str)
   	fmt.Println(chars) // [20013 22269 97 32654 22269 98]
   	str1 := MyStr(chars)
   	fmt.Println(str1) // 中国a美国b
   
   ```

10. for ... range 能够用于循环哪些类型的值？=> 字符串，切片，数组，map，通道。除了这些是否还有其他类型？自定义类型是否可以？

11. P67 提到[]byte(string)转换非常快（O(1)），因为在底层[]byte可以简单地引用字符串的底层字节而无需复制。同样，其逆向转换string([]byte)的原理也类似，其底层字节也无需复制，因此其代价也是O(1)。这些在go源码的什么位置？

12. s[len(s) + 1]，返回什么？

13. P71 可以使用fmt.Fprint()函数和fmt.Fprintf()函数来输出到给定的io.Writer（如一个文件）。请问文件在go源码中的哪里体现出实现了io.Writer接口？

