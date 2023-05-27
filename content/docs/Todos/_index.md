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
3. fmt.Println会调用该类型的`String() string`方法?
4. 哪些fmt.函数会通过反射来调用`Error() string` 或 `String() string`方法？
5. 空标记符_始终都不是一个新的变量吗？P24 -> 是的，请参见 P36
6. 在同一个包同一个文件中，是否可以定义多个init函数？P30
7. 使用fmt.Println()没有任何实参，会打印吗？还是报错？P32
8. go 标识符可以使用多个_开头吗？-> 可以
9. 可以在32位系统上使用float64、int64、complex128？
10. const a = 1; const b = 1.2; var c = a + b; 请问是否会编译错误，c的值是多少？c的类型？-> 可以编译，c=2.2，c的类型是float64
11. 比较运算符，用于判断两个接口变量时，这两个接口变量有什么限制？怎么进行比较？
12. 比较运算符，用于两种不同的类型的变量进行比较，编译会报错吗？可以用一个变量和常量进行比较吗？若变量的类型和常量的类型不一致，是否还可以进行比较？