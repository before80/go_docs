+++
title = "channel"
date = 2023-08-07T13:36:37+08:00
weight = 25
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# channel

```go
package main

import "fmt"

func SendData(ch chan<- int) {
	// 在函数内部只能对 ch 进行发送数据
	ch <- 1
}

func ReceiveData(ch <-chan int) int {
	// 在函数内部只能从 ch 进行接收数据
	return <-ch
}

func main() {
	//nil channel
	var ch1 chan int
	fmt.Printf("%v,%T,len=%d\n", ch1, ch1, len(ch1)) // <nil>,chan int,len=0

	//无缓冲channel
	fmt.Println("----------no cache channel --------------")
	fmt.Println("----------no cache ch2 --------------")
	ch2 := make(chan int)
	fmt.Printf("%v,%T\n", ch2, ch2) // 0x...,chan int
	go func() {
		ch2 <- 2
	}()
	d2 := <-ch2
	close(ch2)
	fmt.Println(d2) // 2

	fmt.Println("----------no cache ch21 --------------")
	// 报错：fatal error: all goroutines are asleep - deadlock!
	//ch21 := make(chan int)
	//ch21 <- 21
	//fmt.Println(<-ch21)

	fmt.Println("----------no cache ch3 --------------")
	ch3 := make(chan int)
	fmt.Printf("%v,%T\n", ch3, ch3) // 0x...,chan int
	go func() {
		d3 := <-ch3
		fmt.Println(d3) // 3
	}()
	ch3 <- 3
	close(ch3)

	fmt.Println("----------no cache ch4 --------------")
	// 报错：fatal error: all goroutines are asleep - deadlock!
	//ch4 := make(chan int)
	//fmt.Printf("%v,%T\n", ch4, ch4) // 0x...,chan int
	//ch4 <- 4
	//go func() {
	//	d4 := <-ch4
	//	fmt.Println(d4)
	//}()
	//close(ch4)

	fmt.Println("----------no cache ch5 --------------")
	// 报错：fatal error: all goroutines are asleep - deadlock!
	//ch5 := *new(chan int)
	//fmt.Printf("%v,%T\n", ch5, ch5)   // <nil>,chan int
	//fmt.Printf("%v,%T\n", &ch5, &ch5) // 0x...,*chan int
	//go func() {
	//	d5 := <-ch5
	//	fmt.Println(d5)
	//}()
	//ch5 <- 5
	//close(ch5)

	fmt.Println("----------no cache ch6 --------------")
	// 报错：fatal error: all goroutines are asleep - deadlock!
	//ch6 := *new(chan int)
	//fmt.Printf("%v,%T\n", ch6, ch6)   // <nil>,chan int
	//fmt.Printf("%v,%T\n", &ch6, &ch6) // 0x...,*chan int
	//go func() {
	//	ch6 <- 6
	//}()
	//d6 := <-ch6
	//fmt.Println(d6)
	//close(ch6)

	fmt.Println("----------no cache ch7 --------------")
	// 报错：fatal error: all goroutines are asleep - deadlock!
	//ch7 := new(chan int)
	//fmt.Printf("%v,%T\n", ch7, ch7)   // 0x...,*chan int
	//fmt.Printf("%v,%T\n", &ch7, &ch7) // 0x...,**chan int
	//fmt.Printf("%v,%T\n", *ch7, *ch7) // <nil>,chan int
	//go func() {
	//	*ch7 <- 7
	//}()
	//d7 := <-*ch7
	//fmt.Println(d7)
	//close(*ch7)

	fmt.Println("----------no cache ch8 --------------")
	// 报错：fatal error: all goroutines are asleep - deadlock!
	//ch8 := new(chan int)
	//fmt.Printf("%v,%T\n", ch8, ch8)   // 0x...,*chan int
	//fmt.Printf("%v,%T\n", &ch8, &ch8) // 0x...,**chan int
	//go func() {
	//	(*ch8) <- 8
	//  close(*ch8)
	//}()
	//d8 := <-(*ch8)
	//fmt.Println(d8)

	fmt.Println("----------no cache ch9 --------------")
	ch9 := *new(chan int)
	fmt.Printf("%v,%T\n", ch9, ch9)   // <nil>,chan int
	fmt.Printf("%v,%T\n", &ch9, &ch9) // 0x...,*chan int

	ch9 = make(chan int) // 需要有重新赋值

	fmt.Printf("%v,%T\n", ch9, ch9) // 0x...,chan int
	go func() {
		ch9 <- 9
		close(ch9)
	}()
	d9 := <-ch9
	fmt.Println(d9) // 9

	fmt.Println("----------no cache ch10 --------------")
	ch10 := make(chan int)
	go func(ch chan int) {
		for i := 0; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch10)

    fmt.Println("----------no cache ch10 for range --------------")
	for d := range ch10 {
		fmt.Println(d)
	}

	//有缓冲channel
	fmt.Println("---------- cache channel --------------")
	fmt.Println("---------- cache ch11 --------------")
	ch11 := make(chan int, 1)
	fmt.Printf("%v,%T,len=%d\n", ch11, ch11, len(ch11)) // 0x...,chan int,len=0

	ch11 <- 1
	fmt.Printf("%v,%T,len=%d\n", ch11, ch11, len(ch11)) // 0x...,chan int,len=1
	fmt.Println(<-ch11)                                 // 1
	close(ch11)

	// 约束函数参数中的channel
	fmt.Println("--------------- 约束函数参数中的channel ---------------")
	ch12 := make(chan int)
	go SendData(ch12)
	fmt.Println(ReceiveData(ch12)) // 1

}

```

