+++
title = "pem"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/encoding/pem@go1.23.0](https://pkg.go.dev/encoding/pem@go1.23.0)

Package pem implements the PEM data encoding, which originated in Privacy Enhanced Mail. The most common use of PEM encoding today is in TLS keys and certificates. See [RFC 1421](https://rfc-editor.org/rfc/rfc1421.html).

​	`pem`包实现了PEM数据编码，该编码起源于隐私增强邮件（Privacy Enhanced Mail）。如今，PEM编码最常用于TLS密钥和证书。详见[RFC 1421](https://rfc-editor.org/rfc/rfc1421.html)。


## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Encode 

``` go 
func Encode(out io.Writer, b *Block) error
```

Encode writes the PEM encoding of b to out.

​	Encode函数将`b`的PEM编码写入`out`。

#### Encode Example
``` go 
package main

import (
	"encoding/pem"
	"log"
	"os"
)

func main() {
	block := &pem.Block{
		Type: "MESSAGE",
		Headers: map[string]string{
			"Animal": "Gopher",
		},
		Bytes: []byte("test"),
	}

	if err := pem.Encode(os.Stdout, block); err != nil {
		log.Fatal(err)
	}
}
Output:

-----BEGIN MESSAGE-----
Animal: Gopher

dGVzdA==
-----END MESSAGE-----
```

### func EncodeToMemory 

``` go 
func EncodeToMemory(b *Block) []byte
```

EncodeToMemory returns the PEM encoding of b.

​	EncodeToMemory函数返回`b`的PEM编码。

If b has invalid headers and cannot be encoded, EncodeToMemory returns nil. If it is important to report details about this error case, use Encode instead.

​	如果`b`具有无效的标头且无法被编码，EncodeToMemory函数返回nil。如果报告这种错误情况的细节很重要，请使用Encode代替。

## 类型

### type Block 

``` go 
type Block struct {
	Type    string            // The type, taken from the preamble (i.e. "RSA PRIVATE KEY").// 类型，取自序言(即 "RSA PRIVATE KEY")。
	Headers map[string]string // 可选的头信息。
	Bytes   []byte            // The decoded bytes of the contents. Typically a DER encoded ASN.1 structure. // 解码后的内容字节。通常是一个DER编码的ASN.1结构。
}
```

A Block represents a PEM encoded structure.

​	Block表示一个PEM编码的结构。

The encoded form is:

​	其编码后的形式是：

```
-----BEGIN Type-----
Headers
base64-encoded Bytes
-----END Type-----
```

where Headers is a possibly empty sequence of Key: Value lines.

​	其中Headers是一个可能是空的Key: Value行序列。

#### func Decode 

``` go 
func Decode(data []byte) (p *Block, rest []byte)
```

Decode will find the next PEM formatted block (certificate, private key etc) in the input. It returns that block and the remainder of the input. If no PEM data is found, p is nil and the whole of the input is returned in rest.

​	Decode函数会在输入中查找下一个PEM格式的块（证书、私钥等）。它返回该块以及剩余的输入。如果未找到PEM数据，`p`为nil，并将整个输入返回给`rest`。

##### Decode Example

``` go 
package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

func main() {
	var pubPEMData = []byte(`
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAlRuRnThUjU8/prwYxbty
WPT9pURI3lbsKMiB6Fn/VHOKE13p4D8xgOCADpdRagdT6n4etr9atzDKUSvpMtR3
CP5noNc97WiNCggBjVWhs7szEe8ugyqF23XwpHQ6uV1LKH50m92MbOWfCtjU9p/x
qhNpQQ1AZhqNy5Gevap5k8XzRmjSldNAFZMY7Yv3Gi+nyCwGwpVtBUwhuLzgNFK/
yDtw2WcWmUU7NuC8Q6MWvPebxVtCfVp/iQU6q60yyt6aGOBkhAX0LpKAEhKidixY
nP9PNVBvxgu3XZ4P36gZV6+ummKdBVnc3NqwBLu5+CcdRdusmHPHd5pHf4/38Z3/
6qU2a/fPvWzceVTEgZ47QjFMTCTmCwNt29cvi7zZeQzjtwQgn4ipN9NibRH/Ax/q
TbIzHfrJ1xa2RteWSdFjwtxi9C20HUkjXSeI4YlzQMH0fPX6KCE7aVePTOnB69I/
a9/q96DiXZajwlpq3wFctrs1oXqBp5DVrCIj8hU2wNgB7LtQ1mCtsYz//heai0K9
PhE4X6hiE0YmeAZjR0uHl8M/5aW9xCoJ72+12kKpWAa0SFRWLy6FejNYCYpkupVJ
yecLk/4L1W0l6jQQZnWErXZYe0PNFcmwGXy1Rep83kfBRNKRy5tvocalLlwXLdUk
AIU+2GKjyT3iMuzZxxFxPFMCAwEAAQ==
-----END PUBLIC KEY-----
and some more`)

	block, rest := pem.Decode(pubPEMData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got a %T, with remaining data: %q", pub, rest)
}

Output:

Got a *rsa.PublicKey, with remaining data: "and some more"

```

