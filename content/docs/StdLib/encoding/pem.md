+++
title = "pem"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# pem

https://pkg.go.dev/encoding/pem@go1.20.1



Package pem implements the PEM data encoding, which originated in Privacy Enhanced Mail. The most common use of PEM encoding today is in TLS keys and certificates. See [RFC 1421](https://rfc-editor.org/rfc/rfc1421.html).

包pem实现了PEM数据编码，它起源于隐私增强邮件。目前PEM编码最常见的用途是在TLS密钥和证书中。参见RFC 1421。






## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func Encode 

``` go 
func Encode(out io.Writer, b *Block) error
```

Encode writes the PEM encoding of b to out.

Encode将b的PEM编码写到out。

##### Example
``` go 
```

#### func EncodeToMemory 

``` go 
func EncodeToMemory(b *Block) []byte
```

EncodeToMemory returns the PEM encoding of b.

EncodeToMemory返回b的PEM编码。

If b has invalid headers and cannot be encoded, EncodeToMemory returns nil. If it is important to report details about this error case, use Encode instead.

如果b有无效的头文件并且不能被编码，EncodeToMemory返回nil。如果报告这种错误情况的细节很重要，请使用Encode代替。

## 类型

### type Block 

``` go 
type Block struct {
	Type    string            // The type, taken from the preamble (i.e. "RSA PRIVATE KEY").// 类型，取自序言(即 "RSA PRIVATE KEY")。
	Headers map[string]string // Optional headers.// 可选的头信息。
	Bytes   []byte            // The decoded bytes of the contents. Typically a DER encoded ASN.1 structure. // 解码后的内容字节数。通常是一个DER编码的ASN.1结构。
}
```

A Block represents a PEM encoded structure.

一个Block代表一个PEM编码的结构。

The encoded form is:

编码后的形式是：

```
-----BEGIN Type-----
Headers
base64-encoded Bytes
-----END Type-----
```

where Headers is a possibly empty sequence of Key: Value lines.

其中Headers是一个可能是空的Key.Value的序列。值的行。

#### func Decode 

``` go 
func Decode(data []byte) (p *Block, rest []byte)
```

Decode will find the next PEM formatted block (certificate, private key etc) in the input. It returns that block and the remainder of the input. If no PEM data is found, p is nil and the whole of the input is returned in rest.

Decode将在输入中找到下一个PEM格式的块(证书、私钥等)。它返回该块和输入的剩余部分。如果没有找到PEM数据，p是nil，整个输入将被返回到rest中。

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

```

