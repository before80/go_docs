+++
title = "md5"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/md5@go1.23.0](https://pkg.go.dev/crypto/md5@go1.23.0)

Package md5 implements the MD5 hash algorithm as defined in [RFC 1321](https://rfc-editor.org/rfc/rfc1321.html).

​	Package md5 实现 RFC 1321 中定义的 MD5 哈希算法。

MD5 is cryptographically broken and should not be used for secure applications.

​	MD5 在密码学上已遭破解，不应在安全应用程序中使用。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/md5/md5.go;l=28)

``` go
const BlockSize = 64
```

The blocksize of MD5 in bytes.

​	MD5 的块大小（以字节为单位）。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/md5/md5.go;l=25)

``` go
const Size = 16
```

The size of an MD5 checksum in bytes.

​	MD5 校验和的大小（以字节为单位）。

## 变量

This section is empty.

## 函数

### func New 

``` go
func New() hash.Hash
```

New returns a new hash.Hash computing the MD5 checksum. The Hash also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

​	New 返回一个新的 hash.Hash 计算 MD5 校验和。Hash 还实现了 encoding.BinaryMarshaler 和 encoding.BinaryUnmarshaler 来编组和取消编组哈希的内部状态。

#### New Example

```go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x", h.Sum(nil))
}
Output:

e2c569be17396eca2a2e3c11578123ed
```



#### New Example (File)

```go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))
}
Output:
```



### func Sum  <- go1.2

``` go
func Sum(data []byte) [Size]byte
```

Sum returns the MD5 checksum of the data.

​	Sum 返回数据的 MD5 校验和。

#### Sum  Example

```go
package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("These pretzels are making me thirsty.")
	fmt.Printf("%x", md5.Sum(data))
}
Output:

b0804ec967f48520697662a204f5fe72
```



## 类型

This section is empty.