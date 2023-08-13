+++
title = "time"
date = 2023-08-07T13:51:08+08:00
weight = 46
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# time

> 原文：https://gobyexample.com/time

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	then := time.Date(
		2023, 8, 10, 15, 23, 58, 651387237, beijing)
	p(then) // 2023-08-10 15:57:00.0450508 +0800 CST m=+0.005236201

	p(then.Year())       // 2023-08-10 15:23:58.651387237 +0000 UTC
	p(then.Month())      // 2023
	p(then.Day())        // August
	p(then.Hour())       // 10
	p(then.Minute())     // 15
	p(then.Second())     // 58
	p(then.Nanosecond()) // 651387237
	p(then.Location())   // Beijing Time

	p(then.Weekday()) // Thursday

	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	diff := now.Sub(then)
	p(diff) // 39m13.224600463s

	p(diff.Hours())       // 0.6536735001286111
	p(diff.Minutes())     // 39.220410007716666
	p(diff.Seconds())     // 2353.224600463
	p(diff.Nanoseconds()) // 2353224600463

	p(then.Add(diff)) // 2023-08-10 16:03:11.8759877 +0800 Beijing Time
	p(then.Add(-diff)) // 2023-08-10 14:44:45.426786774 +0800 Beijing Time
}

```

