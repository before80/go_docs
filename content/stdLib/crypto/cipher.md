+++
title = "cipher"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/cipher@go1.23.0](https://pkg.go.dev/crypto/cipher@go1.23.0)

Package cipher implements standard block cipher modes that can be wrapped around low-level block cipher implementations. See https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html and NIST Special Publication 800-38A.

​	cipher 包实现标准分组密码模式，该模式可以包装在低级分组密码实现中。请参阅 [https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html](https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html) 和 NIST 特别出版物 800-38A。


## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type AEAD  <- go1.2

```go
type AEAD interface {
	// NonceSize returns the size of the nonce that must be passed to Seal
	// and Open.
	NonceSize() int

	// Overhead returns the maximum difference between the lengths of a
	// plaintext and its ciphertext.
	Overhead() int

	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	//
	// To reuse plaintext's storage for the encrypted output, use plaintext[:0]
	// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
	Seal(dst, nonce, plaintext, additionalData []byte) []byte

	// Open decrypts and authenticates ciphertext, authenticates the
	// additional data and, if successful, appends the resulting plaintext
	// to dst, returning the updated slice. The nonce must be NonceSize()
	// bytes long and both it and the additional data must match the
	// value passed to Seal.
	//
	// To reuse ciphertext's storage for the decrypted output, use ciphertext[:0]
	// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
	//
	// Even if the function fails, the contents of dst, up to its capacity,
	// may be overwritten.
	Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error)
}
```

AEAD is a cipher mode providing authenticated encryption with associated data. For a description of the methodology, see https://en.wikipedia.org/wiki/Authenticated_encryption.

​	AEAD 是一种密码模式，提供带关联数据的经过身份验证的加密。有关方法的说明，请参阅 [https://en.wikipedia.org/wiki/Authenticated_encryption](https://en.wikipedia.org/wiki/Authenticated_encryption)。

#### func NewGCM  <- go1.2

``` go
func NewGCM(cipher Block) (AEAD, error)
```

NewGCM returns the given 128-bit, block cipher wrapped in Galois Counter Mode with the standard nonce length.

​	NewGCM 返回给定的 128 位分组密码，该密码采用标准随机数长度包装在 Galois 计数器模式中。

In general, the GHASH operation performed by this implementation of GCM is not constant-time. An exception is when the underlying Block was created by aes.NewCipher on systems with hardware support for AES. See the crypto/aes package documentation for details.

​	通常，此 GCM 实现执行的 GHASH 操作不是恒定时间的。例外情况是，当底层 Block 由 aes.NewCipher 在具有 AES 硬件支持的系统上创建时。有关详细信息，请参阅 crypto/aes 包文档。

##### Example (Decrypt)

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	ciphertext, _ := hex.DecodeString("c3aaa29f002ca75870806e44086700f62ce4d43e902b3888e23ceff797a7a471")
	nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", plaintext)
}
Output:

exampleplaintext
```



##### Example (Encrypt)

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString("6368616e676520746869732070617373776f726420746f206120736563726574")
	plaintext := []byte("exampleplaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("%x\n", ciphertext)
}
Output:
```



#### func NewGCMWithNonceSize  <- go1.5

``` go
func NewGCMWithNonceSize(cipher Block, size int) (AEAD, error)
```

NewGCMWithNonceSize returns the given 128-bit, block cipher wrapped in Galois Counter Mode, which accepts nonces of the given length. The length must not be zero.

​	NewGCMWithNonceSize 返回给定的 128 位分组密码，该密码包装在 Galois 计数器模式中，接受给定长度的随机数。长度不能为零。

Only use this function if you require compatibility with an existing cryptosystem that uses non-standard nonce lengths. All other users should use NewGCM, which is faster and more resistant to misuse.

​	仅当您需要与使用非标准随机数长度的现有密码系统兼容时才使用此功能。所有其他用户都应使用 NewGCM，它速度更快，并且更能抵抗误用。

#### func NewGCMWithTagSize  <- go1.11

``` go
func NewGCMWithTagSize(cipher Block, tagSize int) (AEAD, error)
```

NewGCMWithTagSize returns the given 128-bit, block cipher wrapped in Galois Counter Mode, which generates tags with the given length.

​	NewGCMWithTagSize 返回给定的 128 位块密码，该密码封装在伽罗瓦计数器模式中，该模式生成具有给定长度的标记。

Tag sizes between 12 and 16 bytes are allowed.

​	允许标记大小在 12 到 16 个字节之间。

Only use this function if you require compatibility with an existing cryptosystem that uses non-standard tag lengths. All other users should use NewGCM, which is more resistant to misuse.

​	仅当您需要与使用非标准标记长度的现有密码系统兼容时才使用此功能。所有其他用户都应使用 NewGCM，它更能抵抗误用。

### type Block 

``` go
type Block interface {
	// BlockSize returns the cipher's block size.
	BlockSize() int

	// Encrypt encrypts the first block in src into dst.
	// Dst and src must overlap entirely or not at all.
	Encrypt(dst, src []byte)

	// Decrypt decrypts the first block in src into dst.
	// Dst and src must overlap entirely or not at all.
	Decrypt(dst, src []byte)
}
```

A Block represents an implementation of block cipher using a given key. It provides the capability to encrypt or decrypt individual blocks. The mode implementations extend that capability to streams of blocks.

​	Block 表示使用给定密钥实现块密码。它提供了加密或解密各个块的功能。模式实现将该功能扩展到块流。

### type BlockMode 

``` go
type BlockMode interface {
	// BlockSize returns the mode's block size.
	BlockSize() int

	// CryptBlocks encrypts or decrypts a number of blocks. The length of
	// src must be a multiple of the block size. Dst and src must overlap
	// entirely or not at all.
	//
	// If len(dst) < len(src), CryptBlocks should panic. It is acceptable
	// to pass a dst bigger than src, and in that case, CryptBlocks will
	// only update dst[:len(src)] and will not touch the rest of dst.
	//
	// Multiple calls to CryptBlocks behave as if the concatenation of
	// the src buffers was passed in a single run. That is, BlockMode
	// maintains state and does not reset at each CryptBlocks call.
	CryptBlocks(dst, src []byte)
}
```

A BlockMode represents a block cipher running in a block-based mode (CBC, ECB etc).

​	BlockMode 表示在基于块的模式（CBC、ECB 等）中运行的块密码。

#### func NewCBCDecrypter 

``` go
func NewCBCDecrypter(b Block, iv []byte) BlockMode
```

NewCBCDecrypter returns a BlockMode which decrypts in cipher block chaining mode, using the given Block. The length of iv must be the same as the Block's block size and must match the iv used to encrypt the data.

​	NewCBCDecrypter 返回一个使用给定 Block 在密码块链接模式下解密的 BlockMode。iv 的长度必须与 Block 的块大小相同，并且必须与用于加密数据的 iv 匹配。

##### NewCBCDecrypter  Example

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString("73c86d43a9d700a253a96c85b0f6b03ac9792e0e757f869cca306bd3cba1c62b")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	fmt.Printf("%s\n", ciphertext)
}
Output:

exampleplaintext
```



#### func NewCBCEncrypter 

``` go
func NewCBCEncrypter(b Block, iv []byte) BlockMode
```

NewCBCEncrypter returns a BlockMode which encrypts in cipher block chaining mode, using the given Block. The length of iv must be the same as the Block's block size.

​	NewCBCEncrypter 返回一个使用给定 Block 在密码块链接模式下加密的 BlockMode。iv 的长度必须与 Block 的块大小相同。

##### NewCBCEncrypter Example

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("exampleplaintext")

	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	fmt.Printf("%x\n", ciphertext)
}
Output:
```



### type Stream 

```go
type Stream interface {
	// XORKeyStream XORs each byte in the given slice with a byte from the
	// cipher's key stream. Dst and src must overlap entirely or not at all.
	//
	// If len(dst) < len(src), XORKeyStream should panic. It is acceptable
	// to pass a dst bigger than src, and in that case, XORKeyStream will
	// only update dst[:len(src)] and will not touch the rest of dst.
	//
	// Multiple calls to XORKeyStream behave as if the concatenation of
	// the src buffers was passed in a single run. That is, Stream
	// maintains state and does not reset at each XORKeyStream call.
	XORKeyStream(dst, src []byte)
}
```

A Stream represents a stream cipher.

​	Stream 表示流密码。

#### func NewCFBDecrypter 

``` go
func NewCFBDecrypter(block Block, iv []byte) Stream
```

NewCFBDecrypter returns a Stream which decrypts with cipher feedback mode, using the given Block. The iv must be the same length as the Block's block size.

​	NewCFBDecrypter 返回一个使用给定 Block 在密码反馈模式下解密的 Stream。iv 必须与 Block 的块大小相同。

##### NewCFBDecrypter Example

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	ciphertext, _ := hex.DecodeString("7dd015f06bec7f1b8f6559dad89f4131da62261786845100056b353194ad")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("%s", ciphertext)
}
Output:

some plaintext
```



#### func NewCFBEncrypter 

``` go
func NewCFBEncrypter(block Block, iv []byte) Stream
```

NewCFBEncrypter returns a Stream which encrypts with cipher feedback mode, using the given Block. The iv must be the same length as the Block's block size.

​	NewCFBEncrypter 返回一个使用给定 Block 以密码反馈模式加密的 Stream。iv 必须与 Block 的块大小相同。

##### NewCFBEncrypter Example

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("some plaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	fmt.Printf("%x\n", ciphertext)
}
Output:
```



#### func NewCTR 

``` go
func NewCTR(block Block, iv []byte) Stream
```

NewCTR returns a Stream which encrypts/decrypts using the given Block in counter mode. The length of iv must be the same as the Block's block size.

​	NewCTR 返回一个使用给定 Block 以计数器模式加密/解密的 Stream。iv 的长度必须与 Block 的块大小相同。

##### NewCTR Example

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("some plaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	// CTR mode is the same for both encryption and decryption, so we can
	// also decrypt that ciphertext with NewCTR.

	plaintext2 := make([]byte, len(plaintext))
	stream = cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	fmt.Printf("%s\n", plaintext2)
}
Output:

some plaintext
```



#### func NewOFB 

``` go
func NewOFB(b Block, iv []byte) Stream
```

NewOFB returns a Stream that encrypts or decrypts using the block cipher b in output feedback mode. The initialization vector iv's length must be equal to b's block size.

​	NewOFB 返回一个使用块密码 b 以输出反馈模式加密或解密的 Stream。初始化向量 iv 的长度必须等于 b 的块大小。

##### NewOFB Example

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte("some plaintext")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	// OFB mode is the same for both encryption and decryption, so we can
	// also decrypt that ciphertext with NewOFB.

	plaintext2 := make([]byte, len(plaintext))
	stream = cipher.NewOFB(block, iv)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	fmt.Printf("%s\n", plaintext2)
}
Output:

some plaintext
```



### type StreamReader 

``` go
type StreamReader struct {
	S Stream
	R io.Reader
}
```

StreamReader wraps a Stream into an io.Reader. It calls XORKeyStream to process each slice of data which passes through.

​	StreamReader 将 Stream 封装到 io.Reader 中。它调用 XORKeyStream 来处理通过的每个数据切片。

#### Example

```go
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"io"
	"os"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("6368616e676520746869732070617373")

	encrypted, _ := hex.DecodeString("cf0495cc6f75dafc23948538e79904a9")
	bReader := bytes.NewReader(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	reader := &cipher.StreamReader{S: stream, R: bReader}
	// Copy the input to the output stream, decrypting as we go.
	if _, err := io.Copy(os.Stdout, reader); err != nil {
		panic(err)
	}

	// Note that this example is simplistic in that it omits any
	// authentication of the encrypted data. If you were actually to use
	// StreamReader in this manner, an attacker could flip arbitrary bits in
	// the output.

}
Output:

some secret text
```



#### (StreamReader) Read 

``` go
func (r StreamReader) Read(dst []byte) (n int, err error)
```

### type StreamWriter 

``` go
type StreamWriter struct {
	S   Stream
	W   io.Writer
	Err error // unused
}
```

StreamWriter wraps a Stream into an io.Writer. It calls XORKeyStream to process each slice of data which passes through. If any Write call returns short then the StreamWriter is out of sync and must be discarded. A StreamWriter has no internal buffering; Close does not need to be called to flush write data.

​	StreamWriter 将 Stream 封装到 io.Writer 中。它调用 XORKeyStream 来处理通过的每个数据切片。如果任何 Write 调用返回短，则 StreamWriter 将不同步，并且必须将其丢弃。StreamWriter 没有内部缓冲；无需调用 Close 来刷新写入数据。

#### Example

```go
package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString("6368616e676520746869732070617373")

	bReader := bytes.NewReader([]byte("some secret text"))

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	var out bytes.Buffer

	writer := &cipher.StreamWriter{S: stream, W: &out}
	// Copy the input to the output buffer, encrypting as we go.
	if _, err := io.Copy(writer, bReader); err != nil {
		panic(err)
	}

	// Note that this example is simplistic in that it omits any
	// authentication of the encrypted data. If you were actually to use
	// StreamReader in this manner, an attacker could flip arbitrary bits in
	// the decrypted result.

	fmt.Printf("%x\n", out.Bytes())
}
Output:

cf0495cc6f75dafc23948538e79904a9
```



#### (StreamWriter) Close 

``` go
func (w StreamWriter) Close() error
```

Close closes the underlying Writer and returns its Close return value, if the Writer is also an io.Closer. Otherwise it returns nil.

​	Close 关闭底层 Writer 并返回其 Close 返回值（如果 Writer 也是 io.Closer）。否则它返回 nil。

#### (StreamWriter) Write 

``` go
func (w StreamWriter) Write(src []byte) (n int, err error)
```