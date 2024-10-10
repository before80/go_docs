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
// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now) // 2023-08-26 10:13:42.8992275 +0800 CST m=+0.006364701

	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	then := time.Date(
		2023, 8, 10, 15, 23, 58, 651387237, beijing)
	p(then) // 2023-08-10 15:23:58.651387237 +0800 Beijing Time

	p(then.Year())       // 2023
	p(then.Month())      // August
	p(then.Day())        // 10
	p(then.Hour())       // 15
	p(then.Minute())     // 23
	p(then.Second())     // 58
	p(then.Nanosecond()) // 651387237
	p(then.Location())   // Beijing Time

	p(then.Weekday()) // Thursday

	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	diff := now.Sub(then)
	p(diff) // 378h49m44.247840263s

	p(diff.Hours())       // 0.6536735001286111
	p(diff.Minutes())     // 378.8289577334064
	p(diff.Seconds())     // 22729.7374640043
	p(diff.Nanoseconds()) // 1.363784247840263e+06

	p(then.Add(diff))  // 2023-08-26 10:13:42.8992275 +0800 Beijing Time
	p(then.Add(-diff)) // 2023-07-25 20:34:14.403546974 +0800 Beijing Time
}
```

