+++
title = "运行Go代码"
date = 2024-07-13T14:49:52+08:00
weight = 30
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

## 要运行的Go代码（demo）

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

## 在Linux上运行Go代码

### CentOS

​	假设以上代码存放于`$HOME/goprjs/demo/app.go` 文件中

（a）切换至`$HOME/goprjs/demo/`：

```bash
cd ~/goprjs/demo/
```

（b）编译后运行或直接运行：

```bash
# 编译后运行
go build app.go
./app

# 直接运行
go run app.go
```

### Ubuntu

​	假设以上代码存放于`$HOME/goprjs/demo/app.go` 文件中

（a）切换至`$HOME/goprjs/demo/`：

```bash
cd ~/goprjs/demo/
```

（b）编译后运行或直接运行：

```bash
# 编译后运行
go build app.go
./app

# 直接运行
go run app.go
```

## 在Mac上运行Go代码

> 待买一台Mac电脑，再来处理这部分内容。

## 在Windows上运行Go代码

### windows 10及以上版本

​	假设以上代码存放于`F:\goprjs\demog\app.go` 文件中

（a）切换至`F:\goprjs\demog\`（这里使用`powershell`命令行，关于如何安装`powershell`，你可以查看[https://learn.microsoft.com/en-us/powershell/scripting/install/installing-powershell-on-windows?view=powershell-7.3](https://learn.microsoft.com/en-us/powershell/scripting/install/installing-powershell-on-windows?view=powershell-7.3)）：

```bash
cd F:\goprjs\demog\
```

（b）编译后运行或直接运行：

```bash
# 编译后运行
go build app.go
./app

# 直接运行
go run app.go
```

