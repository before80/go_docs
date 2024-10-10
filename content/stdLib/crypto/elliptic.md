+++
title = "elliptic"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/elliptic@go1.23.0](https://pkg.go.dev/crypto/elliptic@go1.23.0)

Package elliptic implements the standard NIST P-224, P-256, P-384, and P-521 elliptic curves over prime fields.

​	elliptic 包实现标准 NIST P-224、P-256、P-384 和 P-521 素数域上的椭圆曲线。

The P224(), P256(), P384() and P521() values are necessary to use the crypto/ecdsa package. Most other uses should migrate to the more efficient and safer crypto/ecdh package.

​	P224()、P256()、P384() 和 P521() 值对于使用 crypto/ecdsa 包是必需的。大多数其他用途都应迁移到更高效、更安全的 crypto/ecdh 包。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func GenerateKey <- DEPRECATED

``` go
func GenerateKey(curve Curve, rand io.Reader) (priv []byte, x, y *big.Int, err error)
```

GenerateKey returns a public/private key pair. The private key is generated using the given reader, which must return random data.

​	GenerateKey 返回公钥/私钥对。私钥使用给定的读取器生成，该读取器必须返回随机数据。

Deprecated: for ECDH, use the GenerateKey methods of the crypto/ecdh package; for ECDSA, use the GenerateKey function of the crypto/ecdsa package.

​	已弃用：对于 ECDH，请使用 crypto/ecdh 包的 GenerateKey 方法；对于 ECDSA，请使用 crypto/ecdsa 包的 GenerateKey 函数。

### func Marshal <- DEPRECATED

``` go
func Marshal(curve Curve, x, y *big.Int) []byte
```

Marshal converts a point on the curve into the uncompressed form specified in SEC 1, Version 2.0, Section 2.3.3. If the point is not on the curve (or is the conventional point at infinity), the behavior is undefined.

​	Marshal 将曲线上的点转换为 SEC 1 版本 2.0 第 2.3.3 节中指定的未压缩形式。如果该点不在曲线上（或为无穷大处的传统点），则行为未定义。

Deprecated: for ECDH, use the crypto/ecdh package. This function returns an encoding equivalent to that of PublicKey.Bytes in crypto/ecdh.

​	已弃用：对于 ECDH，请使用 crypto/ecdh 包。此函数返回的编码等效于 crypto/ecdh 中 PublicKey.Bytes 的编码。

### func MarshalCompressed  <- go1.15

``` go
func MarshalCompressed(curve Curve, x, y *big.Int) []byte
```

MarshalCompressed converts a point on the curve into the compressed form specified in SEC 1, Version 2.0, Section 2.3.3. If the point is not on the curve (or is the conventional point at infinity), the behavior is undefined.

​	MarshalCompressed 将曲线上的一个点转换为 SEC 1 第 2.0 版第 2.3.3 节中指定的压缩形式。如果该点不在曲线上（或为无穷远处的常规点），则行为未定义。

### func Unmarshal <- DEPRECATED

``` go
func Unmarshal(curve Curve, data []byte) (x, y *big.Int)
```

Unmarshal converts a point, serialized by Marshal, into an x, y pair. It is an error if the point is not in uncompressed form, is not on the curve, or is the point at infinity. On error, x = nil.

​	Unmarshal 将由 Marshal 序列化的点转换为 x, y 对。如果该点不在未压缩形式中、不在曲线上或为无穷远处的点，则会出错。出错时，x = nil。

Deprecated: for ECDH, use the crypto/ecdh package. This function accepts an encoding equivalent to that of the NewPublicKey methods in crypto/ecdh.

​	已弃用：对于 ECDH，请使用 crypto/ecdh 包。此函数接受与 crypto/ecdh 中 NewPublicKey 方法等效的编码。

### func UnmarshalCompressed  <- go1.15

``` go
func UnmarshalCompressed(curve Curve, data []byte) (x, y *big.Int)
```

UnmarshalCompressed converts a point, serialized by MarshalCompressed, into an x, y pair. It is an error if the point is not in compressed form, is not on the curve, or is the point at infinity. On error, x = nil.

​	UnmarshalCompressed 将由 MarshalCompressed 序列化的点转换为 x, y 对。如果该点不在压缩形式中、不在曲线上或为无穷远处的点，则会出错。出错时，x = nil。

## 类型

### type Curve 

``` go
type Curve interface {
	// Params returns the parameters for the curve.
	Params() *CurveParams

	// IsOnCurve reports whether the given (x,y) lies on the curve.
	//
	// Note: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
	// package. The NewPublicKey methods of NIST curves in crypto/ecdh accept
	// the same encoding as the Unmarshal function, and perform on-curve checks.
	IsOnCurve(x, y *big.Int) bool

	// Add returns the sum of (x1,y1) and (x2,y2).
	//
	// Note: this is a low-level unsafe API.
	Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)

	// Double returns 2*(x,y).
	//
	// Note: this is a low-level unsafe API.
	Double(x1, y1 *big.Int) (x, y *big.Int)

	// ScalarMult returns k*(x,y) where k is an integer in big-endian form.
	//
	// Note: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
	// package. Most uses of ScalarMult can be replaced by a call to the ECDH
	// methods of NIST curves in crypto/ecdh.
	ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)

	// ScalarBaseMult returns k*G, where G is the base point of the group
	// and k is an integer in big-endian form.
	//
	// Note: this is a low-level unsafe API. For ECDH, use the crypto/ecdh
	// package. Most uses of ScalarBaseMult can be replaced by a call to the
	// PrivateKey.PublicKey method in crypto/ecdh.
	ScalarBaseMult(k []byte) (x, y *big.Int)
}
```

A Curve represents a short-form Weierstrass curve with a=-3.
Curve 表示 a=-3 的短形式 Weierstrass 曲线。

The behavior of Add, Double, and ScalarMult when the input is not a point on the curve is undefined.
当输入不是曲线上的一点时，Add、Double 和 ScalarMult 的行为未定义。

Note that the conventional point at infinity (0, 0) is not considered on the curve, although it can be returned by Add, Double, ScalarMult, or ScalarBaseMult (but not the Unmarshal or UnmarshalCompressed functions).

​	请注意，常规无穷点 (0, 0) 不在曲线上，尽管它可以通过 Add、Double、ScalarMult 或 ScalarBaseMult 返回（但 Unmarshal 或 UnmarshalCompressed 函数不行）。

#### func P224 

``` go
func P224() Curve
```

P224 returns a Curve which implements NIST P-224 (FIPS 186-3, section D.2.2), also known as secp224r1. The CurveParams.Name of this Curve is “P-224”.

​	P224 返回一个实现 NIST P-224（FIPS 186-3，D.2.2 节）的曲线，也称为 secp224r1。此曲线的 CurveParams.Name 为“P-224”。

Multiple invocations of this function will return the same value, so it can be used for equality checks and switch statements.

​	多次调用此函数将返回相同的值，因此可将其用于相等性检查和 switch 语句。

The cryptographic operations are implemented using constant-time algorithms.

​	使用恒定时间算法实现加密操作。

#### func P256 

``` go
func P256() Curve
```

P256 returns a Curve which implements NIST P-256 (FIPS 186-3, section D.2.3), also known as secp256r1 or prime256v1. The CurveParams.Name of this Curve is “P-256”.

​	P256 返回一个实现 NIST P-256（FIPS 186-3，D.2.3 节）的曲线，也称为 secp256r1 或 prime256v1。此曲线的 CurveParams.Name 为“P-256”。

Multiple invocations of this function will return the same value, so it can be used for equality checks and switch statements.

​	多次调用此函数将返回相同的值，因此可将其用于相等性检查和 switch 语句。

The cryptographic operations are implemented using constant-time algorithms.

​	使用恒定时间算法实现加密操作。

#### func P384 

``` go
func P384() Curve
```

P384 returns a Curve which implements NIST P-384 (FIPS 186-3, section D.2.4), also known as secp384r1. The CurveParams.Name of this Curve is “P-384”.

​	P384 返回一个实现 NIST P-384（FIPS 186-3，D.2.4 节）的曲线，也称为 secp384r1。此曲线的 CurveParams.Name 为“P-384”。

Multiple invocations of this function will return the same value, so it can be used for equality checks and switch statements.

​	多次调用此函数将返回相同的值，因此可将其用于相等性检查和 switch 语句。

The cryptographic operations are implemented using constant-time algorithms.

​	使用恒定时间算法实现加密操作。

#### func P521 

``` go
func P521() Curve
```

P521 returns a Curve which implements NIST P-521 (FIPS 186-3, section D.2.5), also known as secp521r1. The CurveParams.Name of this Curve is “P-521”.

​	P521 返回一个实现 NIST P-521（FIPS 186-3，D.2.5 节）的曲线，也称为 secp521r1。此曲线的 CurveParams.Name 为“P-521”。

Multiple invocations of this function will return the same value, so it can be used for equality checks and switch statements.

​	多次调用此函数将返回相同的值，因此可将其用于相等性检查和 switch 语句。

The cryptographic operations are implemented using constant-time algorithms.

​	加密操作使用恒定时间算法实现。

### type CurveParams 

``` go
type CurveParams struct {
	P       *big.Int // the order of the underlying field
	N       *big.Int // the order of the base point
	B       *big.Int // the constant of the curve equation
	Gx, Gy  *big.Int // (x,y) of the base point
	BitSize int      // the size of the underlying field
	Name    string   // the canonical name of the curve
}
```

CurveParams contains the parameters of an elliptic curve and also provides a generic, non-constant time implementation of Curve.

​	CurveParams 包含椭圆曲线的参数，还提供了 Curve 的通用非恒定时间实现。

Note: Custom curves (those not returned by P224(), P256(), P384(), and P521()) are not guaranteed to provide any security property.

​	注意：自定义曲线（P224()、P256()、P384() 和 P521() 未返回的曲线）不能保证提供任何安全属性。

#### (*CurveParams) Add <- DEPRECATED

``` go
func (curve *CurveParams) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int)
```

Add implements Curve.Add.

​	Add 实现 Curve.Add。

Deprecated: the CurveParams methods are deprecated and are not guaranteed to provide any security property. For ECDH, use the crypto/ecdh package. For ECDSA, use the crypto/ecdsa package with a Curve value returned directly from P224(), P256(), P384(), or P521().

​	已弃用：CurveParams 方法已弃用，不能保证提供任何安全属性。对于 ECDH，请使用 crypto/ecdh 包。对于 ECDSA，请使用 crypto/ecdsa 包，其中 Curve 值直接从 P224()、P256()、P384() 或 P521() 返回。

#### (*CurveParams) Double <- DEPRECATED

``` go
func (curve *CurveParams) Double(x1, y1 *big.Int) (*big.Int, *big.Int)
```

Double implements Curve.Double.
Double 实现 Curve.Double。

Deprecated: the CurveParams methods are deprecated and are not guaranteed to provide any security property. For ECDH, use the crypto/ecdh package. For ECDSA, use the crypto/ecdsa package with a Curve value returned directly from P224(), P256(), P384(), or P521().

​	已弃用：CurveParams 方法已弃用，不能保证提供任何安全属性。对于 ECDH，请使用 crypto/ecdh 包。对于 ECDSA，请使用 crypto/ecdsa 包，其中 Curve 值直接从 P224()、P256()、P384() 或 P521() 返回。

#### (*CurveParams) IsOnCurve <- DEPRECATED

``` go
func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool
```

IsOnCurve implements Curve.IsOnCurve.
IsOnCurve 实现 Curve.IsOnCurve。

Deprecated: the CurveParams methods are deprecated and are not guaranteed to provide any security property. For ECDH, use the crypto/ecdh package. For ECDSA, use the crypto/ecdsa package with a Curve value returned directly from P224(), P256(), P384(), or P521().

​	已弃用：CurveParams 方法已弃用，不能保证提供任何安全属性。对于 ECDH，请使用 crypto/ecdh 包。对于 ECDSA，请使用 crypto/ecdsa 包，其中 Curve 值直接从 P224()、P256()、P384() 或 P521() 返回。

#### (*CurveParams) Params 

``` go
func (curve *CurveParams) Params() *CurveParams
```

#### (*CurveParams) ScalarBaseMult <- DEPRECATED

``` go
func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int)
```

ScalarBaseMult implements Curve.ScalarBaseMult.
ScalarBaseMult 实现 Curve.ScalarBaseMult。

Deprecated: the CurveParams methods are deprecated and are not guaranteed to provide any security property. For ECDH, use the crypto/ecdh package. For ECDSA, use the crypto/ecdsa package with a Curve value returned directly from P224(), P256(), P384(), or P521().

​	已弃用：CurveParams 方法已弃用，不能保证提供任何安全属性。对于 ECDH，请使用 crypto/ecdh 包。对于 ECDSA，请使用 crypto/ecdsa 包，其中 Curve 值直接从 P224()、P256()、P384() 或 P521() 返回。

#### (*CurveParams) ScalarMult <- DEPRECATED

``` go
func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int)
```

ScalarMult implements Curve.ScalarMult.
ScalarMult 实现 Curve.ScalarMult。

Deprecated: the CurveParams methods are deprecated and are not guaranteed to provide any security property. For ECDH, use the crypto/ecdh package. For ECDSA, use the crypto/ecdsa package with a Curve value returned directly from P224(), P256(), P384(), or P521().

​	已弃用：CurveParams 方法已弃用，不能保证提供任何安全属性。对于 ECDH，请使用 crypto/ecdh 包。对于 ECDSA，请使用 crypto/ecdsa 包，其中 Curve 值直接从 P224()、P256()、P384() 或 P521() 返回。