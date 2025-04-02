+++
title = "hmac"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/hmac@go1.24.2](https://pkg.go.dev/crypto/hmac@go1.24.2)

Package hmac implements the Keyed-Hash Message Authentication Code (HMAC) as defined in U.S. Federal Information Processing Standards Publication 198. An HMAC is a cryptographic hash that uses a key to sign a message. The receiver verifies the hash by recomputing it using the same key.

​	hmac 包实现了美国联邦信息处理标准发布的 198 中定义的密钥散列消息认证码 (HMAC)。HMAC 是一种使用密钥对消息进行签名的加密哈希。接收方使用相同的密钥重新计算哈希值来验证哈希值。

Receivers should be careful to use Equal to compare MACs in order to avoid timing side-channels:

​	接收器应小心使用 Equal 比较 MAC，以避免定时旁道：

```go
// ValidMAC reports whether messageMAC is a valid HMAC tag for message.
func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
```


## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Equal  <- go1.1

``` go
func Equal(mac1, mac2 []byte) bool
```

Equal compares two MACs for equality without leaking timing information.

​	Equal 比较两个 MAC 是否相等，而不会泄露定时信息。

### func New 

``` go
func New(h func() hash.Hash, key []byte) hash.Hash
```

New returns a new HMAC hash using the given hash.Hash type and key. New functions like sha256.New from crypto/sha256 can be used as h. h must return a new Hash every time it is called. Note that unlike other hash implementations in the standard library, the returned Hash does not implement encoding.BinaryMarshaler or encoding.BinaryUnmarshaler.

​	New 使用给定的 hash.Hash 类型和密钥返回一个新的 HMAC 哈希。New 函数（例如 crypto/sha256 中的 sha256.New）可用作 h。h 必须在每次调用时返回一个新的哈希。请注意，与标准库中的其他哈希实现不同，返回的哈希不实现 encoding.BinaryMarshaler 或 encoding.BinaryUnmarshaler。

## 类型

This section is empty.