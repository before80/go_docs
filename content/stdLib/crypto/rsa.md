+++
title = "rsa"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/rsa@go1.23.0](https://pkg.go.dev/crypto/rsa@go1.23.0)

Package rsa implements RSA encryption as specified in PKCS #1 and [RFC 8017](https://rfc-editor.org/rfc/rfc8017.html).

​	RSA 包实现了 PKCS #1 和 RFC 8017 中指定的 RSA 加密。

RSA is a single, fundamental operation that is used in this package to implement either public-key encryption or public-key signatures.

​	RSA 是一个单一的基本操作，用于此包中实现公钥加密或公钥签名。

The original specification for encryption and signatures with RSA is PKCS #1 and the terms “RSA encryption” and “RSA signatures” by default refer to PKCS #1 version 1.5. However, that specification has flaws and new designs should use version 2, usually called by just OAEP and PSS, where possible.

​	RSA 加密和签名的原始规范是 PKCS #1，术语“RSA 加密”和“RSA 签名”默认情况下是指 PKCS #1 版本 1.5。但是，该规范存在缺陷，新设计应尽可能使用版本 2，通常称为 OAEP 和 PSS。

Two sets of interfaces are included in this package. When a more abstract interface isn’t necessary, there are functions for encrypting/decrypting with v1.5/OAEP and signing/verifying with v1.5/PSS. If one needs to abstract over the public key primitive, the PrivateKey type implements the Decrypter and Signer interfaces from the crypto package.

​	此包中包含两组接口。当不需要更抽象的接口时，有用于使用 v1.5/OAEP 加密/解密以及使用 v1.5/PSS 签名/验证的函数。如果需要对公钥原语进行抽象，则 PrivateKey 类型会实现 crypto 包中的 Decrypter 和 Signer 接口。

Operations in this package are implemented using constant-time algorithms, except for [GenerateKey](https://pkg.go.dev/crypto/rsa@go1.20.1#GenerateKey), [PrivateKey.Precompute](https://pkg.go.dev/crypto/rsa@go1.20.1#PrivateKey.Precompute), and [PrivateKey.Validate](https://pkg.go.dev/crypto/rsa@go1.20.1#PrivateKey.Validate). Every other operation only leaks the bit size of the involved values, which all depend on the selected key size.

​	此软件包中的操作使用恒定时间算法实现，GenerateKey、PrivateKey.Precompute 和 PrivateKey.Validate 除外。其他每个操作仅泄露所涉及值的位大小，所有这些都取决于所选密钥大小。

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/rsa/pss.go;l=247)

``` go
const (
	// PSSSaltLengthAuto causes the salt in a PSS signature to be as large
	// as possible when signing, and to be auto-detected when verifying.
	PSSSaltLengthAuto = 0
	// PSSSaltLengthEqualsHash causes the salt length to equal the length
	// of the hash used in the signature.
	PSSSaltLengthEqualsHash = -1
)
```

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/rsa/rsa.go;l=549)

``` go
var ErrDecryption = errors.New("crypto/rsa: decryption error")
```

ErrDecryption represents a failure to decrypt a message. It is deliberately vague to avoid adaptive attacks.

​	ErrDecryption 表示无法解密消息。它故意模糊以避免自适应攻击。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/rsa/rsa.go;l=453)

``` go
var ErrMessageTooLong = errors.New("crypto/rsa: message too long for RSA key size")
```

ErrMessageTooLong is returned when attempting to encrypt or sign a message which is too large for the size of the key. When using SignPSS, this can also be returned if the size of the salt is too large.

​	当尝试加密或签名对于密钥大小而言过大的消息时，将返回 ErrMessageTooLong。使用 SignPSS 时，如果 salt 的大小过大，也会返回此错误。

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/rsa/rsa.go;l=553)

``` go
var ErrVerification = errors.New("crypto/rsa: verification error")
```

ErrVerification represents a failure to verify a signature. It is deliberately vague to avoid adaptive attacks.

​	ErrVerification 表示无法验证签名。它故意模糊以避免自适应攻击。

## 函数

### func DecryptOAEP 

``` go
func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
```

DecryptOAEP decrypts ciphertext using RSA-OAEP.

​	DecryptOAEP 使用 RSA-OAEP 解密密文。

OAEP is parameterised by a hash function that is used as a random oracle. Encryption and decryption of a given message must use the same hash function and sha256.New() is a reasonable choice.

​	OAEP 由用作随机预言机的哈希函数参数化。给定消息的加密和解密必须使用相同的哈希函数，sha256.New() 是一个合理的选择。

The random parameter is legacy and ignored, and it can be as nil.

​	random 参数是旧版且被忽略的，它可以为 nil。

The label parameter must match the value given when encrypting. See EncryptOAEP for details.

​	label 参数必须与加密时给定的值匹配。有关详细信息，请参阅 EncryptOAEP。

#### DecryptOAEP Example

```go
ciphertext, _ := hex.DecodeString("4d1ee10e8f286390258c51a5e80802844c3e6358ad6690b7285218a7c7ed7fc3a4c7b950fbd04d4b0239cc060dcc7065ca6f84c1756deb71ca5685cadbb82be025e16449b905c568a19c088a1abfad54bf7ecc67a7df39943ec511091a34c0f2348d04e058fcff4d55644de3cd1d580791d4524b92f3e91695582e6e340a1c50b6c6d78e80b4e42c5b4d45e479b492de42bbd39cc642ebb80226bb5200020d501b24a37bcc2ec7f34e596b4fd6b063de4858dbf5a4e3dd18e262eda0ec2d19dbd8e890d672b63d368768360b20c0b6b8592a438fa275e5fa7f60bef0dd39673fd3989cc54d2cb80c08fcd19dacbc265ee1c6014616b0e04ea0328c2a04e73460")
label := []byte("orders")

plaintext, err := rsa.DecryptOAEP(sha256.New(), nil, test2048Key, ciphertext, label)
if err != nil {
	fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
	return
}

fmt.Printf("Plaintext: %s\n", string(plaintext))

// Remember that encryption only provides confidentiality. The
// ciphertext should be signed before authenticity is assumed and, even
// then, consider that messages might be reordered.


Output:
```



### func DecryptPKCS1v15 

``` go
func DecryptPKCS1v15(random io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)
```

DecryptPKCS1v15 decrypts a plaintext using RSA and the padding scheme from PKCS #1 v1.5. The random parameter is legacy and ignored, and it can be as nil.

​	DecryptPKCS1v15 使用 RSA 和 PKCS #1 v1.5 中的填充方案解密明文。random 参数已过时且被忽略，它可以为 nil。

Note that whether this function returns an error or not discloses secret information. If an attacker can cause this function to run repeatedly and learn whether each instance returned an error then they can decrypt and forge signatures as if they had the private key. See DecryptPKCS1v15SessionKey for a way of solving this problem.

​	请注意，此函数是否返回错误都会泄露机密信息。如果攻击者可以导致此函数重复运行并了解每个实例是否返回错误，那么他们就可以解密并伪造签名，就好像他们拥有私钥一样。请参阅 DecryptPKCS1v15SessionKey 以了解解决此问题的方法。

### func DecryptPKCS1v15SessionKey 

``` go
func DecryptPKCS1v15SessionKey(random io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) error
```

DecryptPKCS1v15SessionKey decrypts a session key using RSA and the padding scheme from PKCS #1 v1.5. The random parameter is legacy and ignored, and it can be as nil. It returns an error if the ciphertext is the wrong length or if the ciphertext is greater than the public modulus. Otherwise, no error is returned. If the padding is valid, the resulting plaintext message is copied into key. Otherwise, key is unchanged. These alternatives occur in constant time. It is intended that the user of this function generate a random session key beforehand and continue the protocol with the resulting value. This will remove any possibility that an attacker can learn any information about the plaintext. See "Chosen Ciphertext Attacks Against Protocols Based on the RSA Encryption Standard PKCS #1", Daniel Bleichenbacher, Advances in Cryptology (Crypto '98).

​	DecryptPKCS1v15SessionKey 使用 RSA 和 PKCS #1 v1.5 的填充方案解密会话密钥。random 参数是旧版参数，已被忽略，可以为 nil。如果密文长度错误或密文大于公有模数，则会返回错误。否则，不会返回错误。如果填充有效，则将所得明文消息复制到 key 中。否则，key 保持不变。这些替代方案在恒定时间内发生。此函数的用户应事先生成随机会话密钥，并使用所得值继续执行协议。这将消除攻击者了解有关明文信息的任何可能。请参阅“针对基于 RSA 加密标准 PKCS #1 的协议的选定密文攻击”，Daniel Bleichenbacher，《密码学进展》(Crypto '98)。

Note that if the session key is too small then it may be possible for an attacker to brute-force it. If they can do that then they can learn whether a random value was used (because it’ll be different for the same ciphertext) and thus whether the padding was correct. This defeats the point of this function. Using at least a 16-byte key will protect against this attack.

​	请注意，如果会话密钥太小，攻击者可能会暴力破解它。如果他们能够做到这一点，那么他们就可以了解是否使用了随机值（因为对于相同的密文，它将不同），从而了解填充是否正确。这违背了此功能的初衷。使用至少 16 字节的密钥将防止此攻击。

#### DecryptPKCS1v15SessionKey Example

RSA is able to encrypt only a very limited amount of data. In order to encrypt reasonable amounts of data a hybrid scheme is commonly used: RSA is used to encrypt a key for a symmetric primitive like AES-GCM.

​	RSA 只能加密非常有限数量的数据。为了加密合理数量的数据，通常使用混合方案：RSA 用于加密对称基元（如 AES-GCM）的密钥。

Before encrypting, data is “padded” by embedding it in a known structure. This is done for a number of reasons, but the most obvious is to ensure that the value is large enough that the exponentiation is larger than the modulus. (Otherwise it could be decrypted with a square-root.)

​	在加密之前，数据通过将其嵌入到已知结构中来“填充”。这样做有许多原因，但最明显的原因是确保该值足够大，以使指数大于模数。（否则可以用平方根解密。）

In these designs, when using PKCS #1 v1.5, it’s vitally important to avoid disclosing whether the received RSA message was well-formed (that is, whether the result of decrypting is a correctly padded message) because this leaks secret information. DecryptPKCS1v15SessionKey is designed for this situation and copies the decrypted, symmetric key (if well-formed) in constant-time over a buffer that contains a random key. Thus, if the RSA result isn’t well-formed, the implementation uses a random key in constant time.

​	在这些设计中，在使用 PKCS #1 v1.5 时，避免泄露收到的 RSA 消息是否格式正确（即解密结果是否为正确填充的消息）非常重要，因为这会泄露秘密信息。DecryptPKCS1v15SessionKey 针对此情况而设计，并以恒定时间将解密的对称密钥（如果格式正确）复制到包含随机密钥的缓冲区中。因此，如果 RSA 结果格式不正确，则实现会以恒定时间使用随机密钥。

```go
// The hybrid scheme should use at least a 16-byte symmetric key. Here
// we read the random key that will be used if the RSA decryption isn't
// well-formed.
key := make([]byte, 32)
if _, err := rand.Read(key); err != nil {
	panic("RNG failure")
}

rsaCiphertext, _ := hex.DecodeString("aabbccddeeff")

if err := rsa.DecryptPKCS1v15SessionKey(nil, rsaPrivateKey, rsaCiphertext, key); err != nil {
	// Any errors that result will be "public” – meaning that they
	// can be determined without any secret information. (For
	// instance, if the length of key is impossible given the RSA
	// public key.)
	fmt.Fprintf(os.Stderr, "Error from RSA decryption: %s\n", err)
	return
}

// Given the resulting key, a symmetric scheme can be used to decrypt a
// larger ciphertext.
block, err := aes.NewCipher(key)
if err != nil {
	panic("aes.NewCipher failed: " + err.Error())
}

// Since the key is random, using a fixed nonce is acceptable as the
// (key, nonce) pair will still be unique, as required.
var zeroNonce [12]byte
aead, err := cipher.NewGCM(block)
if err != nil {
	panic("cipher.NewGCM failed: " + err.Error())
}
ciphertext, _ := hex.DecodeString("00112233445566")
plaintext, err := aead.Open(nil, zeroNonce[:], ciphertext, nil)
if err != nil {
	// The RSA ciphertext was badly formed; the decryption will
	// fail here because the AES-GCM key will be incorrect.
	fmt.Fprintf(os.Stderr, "Error decrypting: %s\n", err)
	return
}

fmt.Printf("Plaintext: %s\n", string(plaintext))


Output:
```



### func EncryptOAEP 

``` go
func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)
```

EncryptOAEP encrypts the given message with RSA-OAEP.

​	EncryptOAEP 使用 RSA-OAEP 加密给定消息。

OAEP is parameterised by a hash function that is used as a random oracle. Encryption and decryption of a given message must use the same hash function and sha256.New() is a reasonable choice.

​	OAEP 由用作随机预言机的哈希函数参数化。给定消息的加密和解密必须使用相同的哈希函数，sha256.New() 是一个合理的选择。

The random parameter is used as a source of entropy to ensure that encrypting the same message twice doesn’t result in the same ciphertext.

​	随机参数用作熵源，以确保两次加密相同的消息不会导致相同的密文。

The label parameter may contain arbitrary data that will not be encrypted, but which gives important context to the message. For example, if a given public key is used to encrypt two types of messages then distinct label values could be used to ensure that a ciphertext for one purpose cannot be used for another by an attacker. If not required it can be empty.

​	label 参数可能包含任意数据，这些数据不会被加密，但会为消息提供重要的上下文。例如，如果给定的公钥用于加密两种类型的消息，那么可以使用不同的标签值来确保攻击者无法将一种用途的密文用于另一种用途。如果不需要，它可以为空。

The message must be no longer than the length of the public modulus minus twice the hash length, minus a further 2.

​	消息的长度不得超过公钥模数的长度减去哈希长度的两倍，再减去 2。

#### EncryptOAEP Example

```go
secretMessage := []byte("send reinforcements, we're going to advance")
label := []byte("orders")

// crypto/rand.Reader is a good source of entropy for randomizing the
// encryption function.
rng := rand.Reader

ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, &test2048Key.PublicKey, secretMessage, label)
if err != nil {
	fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
	return
}

// Since encryption is a randomized function, ciphertext will be
// different each time.
fmt.Printf("Ciphertext: %x\n", ciphertext)

Output:
```



### func EncryptPKCS1v15 

``` go
func EncryptPKCS1v15(random io.Reader, pub *PublicKey, msg []byte) ([]byte, error)
```

EncryptPKCS1v15 encrypts the given message with RSA and the padding scheme from PKCS #1 v1.5. The message must be no longer than the length of the public modulus minus 11 bytes.

​	EncryptPKCS1v15 使用 RSA 和 PKCS #1 v1.5 中的填充方案加密给定消息。消息的长度不得超过公钥模数的长度减去 11 个字节。

The random parameter is used as a source of entropy to ensure that encrypting the same message twice doesn’t result in the same ciphertext.

​	random 参数用作熵源，以确保两次加密相同的消息不会产生相同的密文。

WARNING: use of this function to encrypt plaintexts other than session keys is dangerous. Use RSA OAEP in new protocols.

​	警告：将此函数用于加密除会话密钥以外的明文是危险的。在新的协议中使用 RSA OAEP。

### func SignPKCS1v15 

``` go
func SignPKCS1v15(random io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)
```

SignPKCS1v15 calculates the signature of hashed using RSASSA-PKCS1-V1_5-SIGN from RSA PKCS #1 v1.5. Note that hashed must be the result of hashing the input message using the given hash function. If hash is zero, hashed is signed directly. This isn't advisable except for interoperability.

​	SignPKCS1v15 使用 RSASSA-PKCS1-V1_5-SIGN 从 RSA PKCS #1 v1.5 计算哈希的签名。请注意，hashed 必须是使用给定的哈希函数对输入消息进行哈希的结果。如果哈希为零，则直接对 hashed 进行签名。除了为了实现互操作性之外，不建议这样做。

The random parameter is legacy and ignored, and it can be as nil.

​	random 参数是旧参数，会被忽略，可以为 nil。

This function is deterministic. Thus, if the set of possible messages is small, an attacker may be able to build a map from messages to signatures and identify the signed messages. As ever, signatures provide authenticity, not confidentiality.

​	此函数是确定性的。因此，如果可能的消息集很小，攻击者可能能够构建从消息到签名的映射并识别已签名的消息。一如既往，签名提供真实性，而不是机密性。

#### SignPKCS1v15 Example

```go
message := []byte("message to be signed")

// Only small messages can be signed directly; thus the hash of a
// message, rather than the message itself, is signed. This requires
// that the hash function be collision resistant. SHA-256 is the
// least-strong hash function that should be used for this at the time
// of writing (2016).
hashed := sha256.Sum256(message)

signature, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA256, hashed[:])
if err != nil {
	fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
	return
}

fmt.Printf("Signature: %x\n", signature)

Output:
```



### func SignPSS  <- go1.2

``` go
func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, digest []byte, opts *PSSOptions) ([]byte, error)
```

SignPSS calculates the signature of digest using PSS.

​	SignPSS 使用 PSS 计算摘要的签名。

digest must be the result of hashing the input message using the given hash function. The opts argument may be nil, in which case sensible defaults are used. If opts.Hash is set, it overrides hash.

​	digest 必须是使用给定的哈希函数对输入消息进行哈希的结果。opts 参数可以为 nil，在这种情况下，将使用合理的默认值。如果设置了 opts.Hash，则它将覆盖哈希。

### func VerifyPKCS1v15 

``` go
func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
```

VerifyPKCS1v15 verifies an RSA PKCS #1 v1.5 signature. hashed is the result of hashing the input message using the given hash function and sig is the signature. A valid signature is indicated by returning a nil error. If hash is zero then hashed is used directly. This isn't advisable except for interoperability.

​	VerifyPKCS1v15 验证 RSA PKCS #1 v1.5 签名。hashed 是使用给定的哈希函数对输入消息进行哈希的结果，sig 是签名。返回 nil 错误表示签名有效。如果 hash 为零，则直接使用 hashed。除了为了实现互操作性之外，不建议这样做。

#### VerifyPKCS1v15 Example

```go
message := []byte("message to be signed")
signature, _ := hex.DecodeString("ad2766728615cc7a746cc553916380ca7bfa4f8983b990913bc69eb0556539a350ff0f8fe65ddfd3ebe91fe1c299c2fac135bc8c61e26be44ee259f2f80c1530")

// Only small messages can be signed directly; thus the hash of a
// message, rather than the message itself, is signed. This requires
// that the hash function be collision resistant. SHA-256 is the
// least-strong hash function that should be used for this at the time
// of writing (2016).
hashed := sha256.Sum256(message)

err := rsa.VerifyPKCS1v15(&rsaPrivateKey.PublicKey, crypto.SHA256, hashed[:], signature)
if err != nil {
	fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
	return
}

// signature is a valid signature of message from the public key.

Output:
```



### func VerifyPSS  <- go1.2

``` go
func VerifyPSS(pub *PublicKey, hash crypto.Hash, digest []byte, sig []byte, opts *PSSOptions) error
```

VerifyPSS verifies a PSS signature.

​	VerifyPSS 验证 PSS 签名。

A valid signature is indicated by returning a nil error. digest must be the result of hashing the input message using the given hash function. The opts argument may be nil, in which case sensible defaults are used. opts.Hash is ignored.

​	返回 nil 错误表示签名有效。digest 必须是使用给定的哈希函数对输入消息进行哈希的结果。opts 参数可以为 nil，在这种情况下，将使用合理的默认值。opts.Hash 被忽略。

## 类型

### type CRTValue 

``` go
type CRTValue struct {
	Exp   *big.Int // D mod (prime-1).
	Coeff *big.Int // R·Coeff ≡ 1 mod Prime.
	R     *big.Int // product of primes prior to this (inc p and q).
}
```

CRTValue contains the precomputed Chinese remainder theorem values.

​	CRTValue 包含预先计算的中国剩余定理值。

### type OAEPOptions  <- go1.5

``` go
type OAEPOptions struct {
	// Hash is the hash function that will be used when generating the mask.
	Hash crypto.Hash

	// MGFHash is the hash function used for MGF1.
	// If zero, Hash is used instead.
	MGFHash crypto.Hash

	// Label is an arbitrary byte string that must be equal to the value
	// used when encrypting.
	Label []byte
}
```

OAEPOptions is an interface for passing options to OAEP decryption using the crypto.Decrypter interface.

​	OAEPOptions 是一个接口，用于使用 crypto.Decrypter 接口将选项传递给 OAEP 解密。

### type PKCS1v15DecryptOptions  <- go1.5

``` go
type PKCS1v15DecryptOptions struct {
	// SessionKeyLen is the length of the session key that is being
	// decrypted. If not zero, then a padding error during decryption will
	// cause a random plaintext of this length to be returned rather than
	// an error. These alternatives happen in constant time.
	SessionKeyLen int
}
```

PKCS1v15DecryptOptions is for passing options to PKCS #1 v1.5 decryption using the crypto.Decrypter interface.

​	PKCS1v15DecryptOptions 用于将选项传递给使用 crypto.Decrypter 接口的 PKCS #1 v1.5 解密。

### type PSSOptions  <- go1.2

``` go
type PSSOptions struct {
	// SaltLength controls the length of the salt used in the PSS signature. It
	// can either be a positive number of bytes, or one of the special
	// PSSSaltLength constants.
	SaltLength int

	// Hash is the hash function used to generate the message digest. If not
	// zero, it overrides the hash function passed to SignPSS. It's required
	// when using PrivateKey.Sign.
	Hash crypto.Hash
}
```

PSSOptions contains options for creating and verifying PSS signatures.

​	PSSOptions 包含用于创建和验证 PSS 签名的选项。

#### (*PSSOptions) HashFunc  <- go1.4

``` go
func (opts *PSSOptions) HashFunc() crypto.Hash
```

HashFunc returns opts.Hash so that PSSOptions implements crypto.SignerOpts.

​	HashFunc 返回 opts.Hash，以便 PSSOptions 实现 crypto.SignerOpts。

### type PrecomputedValues 

``` go
type PrecomputedValues struct {
	Dp, Dq *big.Int // D mod (P-1) (or mod Q-1)
	Qinv   *big.Int // Q^-1 mod P

	// CRTValues is used for the 3rd and subsequent primes. Due to a
	// historical accident, the CRT for the first two primes is handled
	// differently in PKCS #1 and interoperability is sufficiently
	// important that we mirror this.
	//
	// Note: these values are still filled in by Precompute for
	// backwards compatibility but are not used. Multi-prime RSA is very rare,
	// and is implemented by this package without CRT optimizations to limit
	// complexity.
	CRTValues []CRTValue
	// contains filtered or unexported fields
}
```

### type PrivateKey 

``` go
type PrivateKey struct {
	PublicKey            // public part.
	D         *big.Int   // private exponent
	Primes    []*big.Int // prime factors of N, has >= 2 elements.

	// Precomputed contains precomputed values that speed up RSA operations,
	// if available. It must be generated by calling PrivateKey.Precompute and
	// must not be modified.
	Precomputed PrecomputedValues
}
```

A PrivateKey represents an RSA key

​	PrivateKey 表示 RSA 密钥

#### func GenerateKey 

``` go
func GenerateKey(random io.Reader, bits int) (*PrivateKey, error)
```

GenerateKey generates an RSA keypair of the given bit size using the random source random (for example, crypto/rand.Reader).

​	GenerateKey 使用随机源 random（例如，crypto/rand.Reader）生成给定位数的 RSA 密钥对。

#### func GenerateMultiPrimeKey <-DEPRECATED

``` go
func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*PrivateKey, error)
```

GenerateMultiPrimeKey generates a multi-prime RSA keypair of the given bit size and the given random source.

​	GenerateMultiPrimeKey 生成给定位数和给定随机源的多素数 RSA 密钥对。

Table 1 in “[On the Security of Multi-prime RSA](http://www.cacr.math.uwaterloo.ca/techreports/2006/cacr2006-16.pdf)” suggests maximum numbers of primes for a given bit size.

​	“多素数 RSA 的安全性”中的表 1 建议了给定位数的最大素数数。

Although the public keys are compatible (actually, indistinguishable) from the 2-prime case, the private keys are not. Thus it may not be possible to export multi-prime private keys in certain formats or to subsequently import them into other code.

​	虽然公钥与 2 素数情况兼容（实际上无法区分），但私钥不兼容。因此，可能无法以某些格式导出多素数私钥或随后将其导入其他代码。

This package does not implement CRT optimizations for multi-prime RSA, so the keys with more than two primes will have worse performance.

​	此软件包未实现多素数 RSA 的 CRT 优化，因此具有两个以上素数的密钥的性能会更差。

Deprecated: The use of this function with a number of primes different from two is not recommended for the above security, compatibility, and performance reasons. Use GenerateKey instead.

​	已弃用：出于上述安全、兼容性和性能原因，不建议将此函数与数量不同的素数（不同于两个）一起使用。请改用 GenerateKey。

#### (*PrivateKey) Decrypt  <- go1.5

``` go
func (priv *PrivateKey) Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
```

Decrypt decrypts ciphertext with priv. If opts is nil or of type *PKCS1v15DecryptOptions then PKCS #1 v1.5 decryption is performed. Otherwise opts must have type *OAEPOptions and OAEP decryption is done.

​	Decrypt 使用 priv 解密密文。如果 opts 为 nil 或类型为 *PKCS1v15DecryptOptions，则执行 PKCS #1 v1.5 解密。否则，opts 的类型必须为 *OAEPOptions，并执行 OAEP 解密。

#### (*PrivateKey) Equal  <- go1.15

``` go
func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool
```

Equal reports whether priv and x have equivalent values. It ignores Precomputed values.

​	Equal 报告 priv 和 x 是否具有等效值。它忽略 Precomputed 值。

#### (*PrivateKey) Precompute 

``` go
func (priv *PrivateKey) Precompute()
```

Precompute performs some calculations that speed up private key operations in the future.

​	Precompute 执行一些计算，以加快将来的私钥操作。

#### (*PrivateKey) Public  <- go1.4

``` go
func (priv *PrivateKey) Public() crypto.PublicKey
```

Public returns the public key corresponding to priv.

​	Public 返回与 priv 对应的公钥。

#### (*PrivateKey) Sign  <- go1.4

``` go
func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
```

Sign signs digest with priv, reading randomness from rand. If opts is a *PSSOptions then the PSS algorithm will be used, otherwise PKCS #1 v1.5 will be used. digest must be the result of hashing the input message using opts.HashFunc().

​	Sign 使用 priv 对摘要进行签名，从 rand 读取随机数。如果 opts 是 *PSSOptions，则将使用 PSS 算法，否则将使用 PKCS #1 v1.5。摘要必须是使用 opts.HashFunc() 对输入消息进行哈希的结果。

This method implements crypto.Signer, which is an interface to support keys where the private part is kept in, for example, a hardware module. Common uses should use the Sign* functions in this package directly.

​	此方法实现 crypto.Signer，这是一个接口，用于支持私有部分（例如，硬件模块）中的密钥。常见用法应直接使用此包中的 Sign* 函数。

#### (*PrivateKey) Validate 

``` go
func (priv *PrivateKey) Validate() error
```

Validate performs basic sanity checks on the key. It returns nil if the key is valid, or else an error describing a problem.

​	Validate 对密钥执行基本健全性检查。如果密钥有效，则返回 nil，否则返回描述问题的错误。

### type PublicKey 

``` go
type PublicKey struct {
	N *big.Int // modulus
	E int      // public exponent
}
```

A PublicKey represents the public part of an RSA key.

​	PublicKey 表示 RSA 密钥的公有部分。

#### (*PublicKey) Equal  <- go1.15

``` go
func (pub *PublicKey) Equal(x crypto.PublicKey) bool
```

Equal reports whether pub and x have the same value.

​	Equal 报告 pub 和 x 是否具有相同的值。

#### (*PublicKey) Size  <- go1.11

``` go
func (pub *PublicKey) Size() int
```

Size returns the modulus size in bytes. Raw signatures and ciphertexts for or by this public key will have the same size.

​	Size 以字节为单位返回模数大小。此公钥的原始签名和密文或由此公钥生成的原始签名和密文将具有相同的大小。