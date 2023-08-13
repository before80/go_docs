+++
title = "error"
date = 2023-08-07T13:36:12+08:00
weight = 23
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# error

```go
package main

import (
	"errors"
	"fmt"
)

// const E0 = errors.New("can't open file") // 编译报错：errors.New("can't open file") (value of type error) is not constant
var e1 = errors.New("can't open file")

type denominatorErr struct {
	denominator int
	err         string
}

func (e *denominatorErr) Error() string {
	return fmt.Sprintf("%d - %s", e.denominator, e.err)
}

func Divide1(a, b int) (quotient, remainder int, err error) {
	if b == 0 {
		return 0, 0, &denominatorErr{0, "The divisor can't be 0."}
	}

	return a / b, a % b, nil
}

func Divide2(a, b int) (quotient, remainder int, err error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("happen: %w", &denominatorErr{0, "The divisor can't be 0."})
	}

	return a / b, a % b, nil
}

func main() {
	fmt.Printf("%v|%T|%q\n", e1, e1, e1) // can't open file|*errors.errorString|"can't open file"happen: can't open file

	e2 := fmt.Errorf("happen: %w", e1)
	fmt.Printf("%v|%T|%q\n", e2, e2, e2) // happen: can't open file|*fmt.wrapError|"happen: can't open file"

	fmt.Println(errors.Is(e2, e1)) //true
	if errors.Is(e2, e1) {
		fmt.Println("在e2中匹配到e1的错误") // 在e2中匹配到e1的错误
	}

	_, _, err := Divide1(2, 0)
	fmt.Printf("%v|%T|%q\n", err, err, err) // 0 - The divisor can't be 0.|*main.denominatorErr|"0 - The divisor can't be 0."
	if deErr, ok := err.(*denominatorErr); ok {
		fmt.Println(deErr.denominator) // 0
		fmt.Println(deErr.err)         // The divisor can't be 0.
	}

	_, _, err = Divide2(2, 0)
	fmt.Printf("%v|%T|%q\n", err, err, err) // happen: 0 - The divisor can't be 0.|*fmt.wrapError|"happen: 0 - The divisor can't be 0."
	//var targetErr denominatorErr // panic: errors: *target must be interface or implement error
	var targetErr *denominatorErr
	if errors.As(err, &targetErr) {
		fmt.Println(targetErr.denominator) // 0
		fmt.Println(targetErr.err)         // The divisor can't be 0.
	}

	fmt.Println(Divide1(3, 2)) // 1 1 <nil>
	fmt.Println(Divide2(3, 2)) // 1 1 <nil>

}

```

