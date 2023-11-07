+++
title = "sha1"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
https://pkg.go.dev/crypto/sha1@go1.21.3

Package sha1 implements the SHA-1 hash algorithm as defined in [RFC 3174](https://rfc-editor.org/rfc/rfc3174.html).

SHA-1 is cryptographically broken and should not be used for secure applications.

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/sha1/sha1.go;l=26)

``` go
const BlockSize = 64
```

The blocksize of SHA-1 in bytes.

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/sha1/sha1.go;l=23)

``` go
const Size = 20
```

The size of a SHA-1 checksum in bytes.

## 变量

This section is empty.

## 函数

### func New 

``` go
func New() hash.Hash
```

New returns a new hash.Hash computing the SHA1 checksum. The Hash also implements encoding.BinaryMarshaler and encoding.BinaryUnmarshaler to marshal and unmarshal the internal state of the hash.

#### New Example

```go
package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	h := sha1.New()
	io.WriteString(h, "His money is twice tainted:")
	io.WriteString(h, " 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
}
Output:

59 7f 6a 54 00 10 f9 4c 15 d7 18 06 a9 9a 2c 87 10 e7 47 bd
```



#### New Example (File)

```go
package main

import (
	"crypto/sha1"
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

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("% x", h.Sum(nil))
}

Output:

2009/11/10 23:00:00 open file.txt: no such file or directory
```



### func Sum  <- go1.2

``` go
func Sum(data []byte) [Size]byte
```

Sum returns the SHA-1 checksum of the data.

#### Sum  Example

```go
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	data := []byte("This page intentionally left blank.")
	fmt.Printf("% x", sha1.Sum(data))
}
Output:

af 06 49 23 bb f2 30 15 96 aa c4 c2 73 ba 32 17 8e bc 4a 96
```



## 类型

This section is empty.