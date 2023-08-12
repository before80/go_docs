+++
title = "base64 encoding"
date = 2023-08-07T13:53:05+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# base64 encoding

> 原文：https://gobyexample.com/base64-encoding

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc) // YWJjMTIzIT8kKiYoKSctPUB+

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec)) // abc123!?$*&()'-=@~

	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc) // YWJjMTIzIT8kKiYoKSctPUB-

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec)) // abc123!?$*&()'-=@~
}

```

