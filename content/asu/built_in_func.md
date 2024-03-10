+++
title = "built_in_func"
date = 2023-08-01T13:08:51+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# 内置函数

## append()

### 定义

```go
func append(slice []Type, elems ...Type) []Type
```



### 注意

- 只能用于切片类型的变量！

- append函数是一个[可变参数]({{< ref "/langSpec/Types#function-types-函数型">}})函数！

### 特例

​	**作为一个特例**，如果第一个实参的[核心类型]({{< ref "/langSpec/PropertiesOfTypesAndValues#core-types-核心类型">}})是`[]byte`，`append`也接受核心类型是[bytestring]({{< ref "/langSpec/PropertiesOfTypesAndValues#core-types-核心类型">}})的第二个实参，后面跟随`...` 。这种形式追加了字节切片或字符串的字节。

```go
package main

import "fmt"

func main() {
	s := append([]byte("你好世界！"), "你好中国"...)
	fmt.Printf("%q,%T\n", s, s)
}

Output:
"你好世界！你好中国",[]uint8

```



## cap()



## clear() <- go 1.21

## close()

## copy()

## complex()

## delete()

## imag()

## len()

## make()

## max() <- go 1.21

## min() <- go 1.21

## new()

## panic()

## print()

## println()

## real()

## recover()



