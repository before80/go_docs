package main

import (
	"fmt"
	"sync"
	"time"
)

func forRange(withSleep bool) {
	a := []int{1, 2, 3, 4, 5, 6}
	var wg sync.WaitGroup
	for _, v := range a {
		go func() {
			wg.Add(1)
			if withSleep {
				time.Sleep(time.Second)
			}
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	fmt.Println("withSleep = false")
	forRange(false)
	time.Sleep(time.Second)
	fmt.Println("withSleep = true")
	forRange(true)
}
