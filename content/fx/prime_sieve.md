+++
title = "素数筛 - prime sieve"
weight = 1000
date = 2024-07-02T06:27:37+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++



代码如下：

> 代码摘自[https://go.dev/play/p/9U22NfrXeq](https://go.dev/play/p/9U22NfrXeq)
>
> prime: n. 素数，质数，盛年，壮年时期
>
> ​		adj. 主要的，原始的，素数的
>
> sieve /siv/  n. 筛子 vt. 筛

```go
// A concurrent prime sieve

package main

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

```

> 我：能否用更形象的说法来解释下，以上这段代码中 ch = ch1 ？
>
> ChatGPT：当然可以！让我们用一个形象的比喻来解释这段代码中的 `ch = ch1`。
>
> ### 比喻
>
> 想象你在一个流水线上的工作场景，你需要不断地筛选传送带上经过的物品。你的任务是把传送带上的物品按照某种规则筛选出来。
>
> 1. **起点：生成物品**：
>    - 有一个传送带（通道 `ch`）不停地将数字物品（从2开始的自然数）传送过来。这相当于 `Generate` 函数在向传送带 `ch` 发送数字。
> 2. **第一层筛选：筛选器1**：
>    - 你是第一层的筛选器，你负责筛选掉所有能被2整除的物品，并将剩下的物品传送到下一个传送带（新的通道 `ch1`）。
>    - 这时你发现2是一个素数，将其记录下来（打印出来）。
>    - 你在开始工作之前，需要为下一个筛选器设置一个新的传送带（`ch1`）。
>    - 工作完毕后，你告诉下一个筛选器：“这是新的传送带，继续筛选。”
> 3. **第二层筛选：筛选器2**：
>    - 第二层的筛选器接过新的传送带（`ch = ch1`），他负责筛选掉所有能被3整除的物品（从上一层筛选后的物品中）。
>    - 他发现3是一个素数，将其记录下来（打印出来）。
>    - 他也要为下一个筛选器设置一个新的传送带（`ch1`）。
> 4. **后续层筛选**：
>    - 每一层筛选器都会重复同样的过程，筛选出新的素数，设置新的传送带，将新传送带传给下一层筛选器。
>    - 每一层的筛选器只关注从上一层传送过来的物品，并且根据当前的素数进行筛选。
>
> ### 总结
>
> `ch = ch1` 就是把旧的传送带（旧的筛选器输出）换成新的传送带（新的筛选器输入），确保每一层的筛选器都能接收到前一层筛选后的结果，并继续筛选，直到所有需要筛选的层次都完成。这样，每一层筛选器都在其自己的传送带上工作，通过层层传递和筛选，最终得到所有的素数。

