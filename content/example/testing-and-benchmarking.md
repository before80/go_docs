+++
title = "testing-and-benchmarking"
date = 2023-08-07T13:56:08+08:00
weight = 61
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# testing and benchmarking

> 原文：https://gobyexample.com/testing-and-benchmarking

```go
// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

```
单元测试：

```bash
testing_and_benchmarking $ go test -v
=== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/2,-2
=== RUN   TestIntMinTableDriven/0,-1
=== RUN   TestIntMinTableDriven/-1,0
--- PASS: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)

```

基准测试：

```bash
testing_and_benchmarking $ go test -bench .
goos: windows
goarch: amd64
pkg: by_example/testing_and_benchmarking
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkIntMin-12      1000000000               0.2925 ns/op
PASS
ok      by_example/testing_and_benchmarking     0.401s
```

或

```go
testing_and_benchmarking $ go test -bench=.                
ok      by_example/testing_and_benchmarking     0.052s
```

