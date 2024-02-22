+++
title = "rc4"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/rc4@go1.21.3](https://pkg.go.dev/crypto/rc4@go1.21.3)

Package rc4 implements RC4 encryption, as defined in Bruce Schneier's Applied Cryptography.

​	 rc4 包实现 RC4 加密，如 Bruce Schneier 的应用密码术中所定义。

RC4 is cryptographically broken and should not be used for secure applications.

​	RC4 在密码学上已被攻破，不应将其用于安全应用程序。


## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type Cipher 

``` go
type Cipher struct {
	// contains filtered or unexported fields
}
```

A Cipher is an instance of RC4 using a particular key.

​	Cipher是使用特定密钥的 RC4 实例。

#### func NewCipher 

``` go
func NewCipher(key []byte) (*Cipher, error)
```

NewCipher creates and returns a new Cipher. The key argument should be the RC4 key, at least 1 byte and at most 256 bytes.

​	NewCipher 创建并返回一个新密码。key 参数应为 RC4 密钥，至少 1 个字节，最多 256 个字节。

#### (*Cipher) Reset <- DEPRECATED

```go
func (c *Cipher) Reset()
```

Reset zeros the key data and makes the Cipher unusable.

​	Reset 将密钥数据清零，使密码不可用。

Deprecated: Reset can't guarantee that the key will be entirely removed from the process's memory.

​	已弃用：Reset 无法保证密钥将从进程的内存中完全删除。

#### (*Cipher) XORKeyStream 

``` go
func (c *Cipher) XORKeyStream(dst, src []byte)
```

XORKeyStream sets dst to the result of XORing src with the key stream. Dst and src must overlap entirely or not at all.

​	XORKeyStream 将 dst 设置为 src 与密钥流进行异或运算的结果。Dst 和 src 必须完全重叠或根本不重叠。

### type KeySizeError 

``` go
type KeySizeError int
```

#### (KeySizeError) Error 

``` go
func (k KeySizeError) Error() string
```