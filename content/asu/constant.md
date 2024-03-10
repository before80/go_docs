+++
title = "常量"
weight = 20
date = 2023-06-12T10:56:55+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# 常量

```markmap
# 常量
## 常量名

## 常量值
### 布尔字面量
- true（也是预算声明的标识符）
- false（也是预算声明的标识符）
### 整数字面量
### 浮点数字面量
### 复数字面量
### 符文字面量
### 字符串字面量
### 表示常量的标识符
### 常量表达式
- 只能包含常量操作数
- 在编译时进行求值
- 常量表达式总是被精确地求值
- 常量除法或取余操作的除数一定不能为零
### 结果为常量的转换
### 内置函数的结果值
```

​	注意：常量的值是编译时就计算好的！参阅[常量表达式]({{< ref "/langSpec/Expressions#constant-expressions-常量表达式">}})。



难道常量值不能是某个标准库或第三方包的函数或方法的结果值？

验证下：

{{< tabpane text=true >}}

{{< tab header="测试1" >}}

```go
package main

import (
	"errors"
	"fmt"
	"runtime"
)

const ERR1 = fmt.Errorf("%s", "未知错误")
const ERR2 = errors.New("无法读取文件")
const LOGICALCPUs = runtime.NumCPU()

func main() {
	fmt.Println(ERR1) // .\main.go:9:14: fmt.Errorf("%s", "未知错误") (value of type error) is not constant
	fmt.Println(ERR2) // .\main.go:10:14: errors.New("无法读取文件") (value of type error) is not constant
	fmt.Println(LOGICALCPUs) // .\main.go:11:21: runtime.NumCPU() (value of type int) is not constant
}

```

{{< /tab >}}

{{< tab header="测试2" >}}

```go

```

{{< /tab >}}

{{< tab header="测试3" >}}

```go

```

{{< /tab >}}

{{< /tabpane >}}
