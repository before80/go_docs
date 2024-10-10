+++
title = "Time Formatting / Parsing"
date = 2023-08-07T13:51:32+08:00
weight = 48
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Time Formatting / Parsing

> 原文：https://gobyexample.com/time-formatting-parsing

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339)) // 2023-08-10T16:53:15+08:00

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")

	p(t1) // 2012-11-01 22:08:41 +0000 +0000

	p(t.Format("3:04PM"))                           // 4:53PM
	p(t.Format("Mon Jan _2 15:04:05 2006"))         // Thu Aug 10 16:53:15 2023
	p(t.Format("2006-01-02T15:04:05.999999-07:00")) // 2023-08-10T16:53:15.846298+08:00

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")

	p(t2) // 0000-01-01 20:41:00 +0000 UTC

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()) // 2023-08-10T16:53:15-00:00

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")

	p(e) // parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
}

```

