+++
title = "x509/pkix"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/x509/pkix@go1.23.0](https://pkg.go.dev/crypto/x509/pkix@go1.23.0)

Package pkix contains shared, low level structures used for ASN.1 parsing and serialization of X.509 certificates, CRL and OCSP.

​	 pkix 包包含用于 ASN.1 解析和 X.509 证书、CRL 和 OCSP 序列化的共享低级结构。

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type AlgorithmIdentifier 

``` go
type AlgorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.RawValue `asn1:"optional"`
}
```

AlgorithmIdentifier represents the ASN.1 structure of the same name. See [RFC 5280, section 4.1.1.2](https://rfc-editor.org/rfc/rfc5280.html#section-4.1.1.2).

​	AlgorithmIdentifier 表示同名 ASN.1 结构。请参阅 RFC 5280 第 4.1.1.2 节。

### type AttributeTypeAndValue

```go
type AttributeTypeAndValue struct {
	Type  asn1.ObjectIdentifier
	Value any
}
```

AttributeTypeAndValue mirrors the ASN.1 structure of the same name in [RFC 5280, Section 4.1.2.4](https://rfc-editor.org/rfc/rfc5280.html#section-4.1.2.4).

​	AttributeTypeAndValue 镜像 RFC 5280 第 4.1.2.4 节中同名的 ASN.1 结构。

### type AttributeTypeAndValueSET <- go1.3

```go
type AttributeTypeAndValueSET struct {
	Type  asn1.ObjectIdentifier
	Value [][]AttributeTypeAndValue `asn1:"set"`
}
```

AttributeTypeAndValueSET represents a set of ASN.1 sequences of AttributeTypeAndValue sequences from [RFC 2986](https://rfc-editor.org/rfc/rfc2986.html) (PKCS #10).

​	AttributeTypeAndValueSET 表示来自 RFC 2986 (PKCS #10) 的 AttributeTypeAndValue 序列的 ASN.1 序列集。

### type CertificateList <- DEPRECATED

```go
type CertificateList struct {
	TBSCertList        TBSCertificateList
	SignatureAlgorithm AlgorithmIdentifier
	SignatureValue     asn1.BitString
}
```

CertificateList represents the ASN.1 structure of the same name. See [RFC 5280, section 5.1](https://rfc-editor.org/rfc/rfc5280.html#section-5.1). Use Certificate.CheckCRLSignature to verify the signature.

​	CertificateList 表示同名的 ASN.1 结构。请参阅 RFC 5280 第 5.1 节。使用 Certificate.CheckCRLSignature 验证签名。

Deprecated: x509.RevocationList should be used instead.

​	已弃用：应改用 x509.RevocationList。

#### (*CertificateList) HasExpired

```go
func (certList *CertificateList) HasExpired(now time.Time) bool
```

HasExpired reports whether certList should have been updated by now.

​	HasExpired 报告 certList 是否应该已经更新。

### type Extension

```go
type Extension struct {
	Id       asn1.ObjectIdentifier
	Critical bool `asn1:"optional"`
	Value    []byte
}
```

Extension represents the ASN.1 structure of the same name. See [RFC 5280, section 4.2](https://rfc-editor.org/rfc/rfc5280.html#section-4.2).

​	Extension 表示同名的 ASN.1 结构。请参阅 RFC 5280 第 4.2 节。

### type Name

```go
type Name struct {
	Country, Organization, OrganizationalUnit []string
	Locality, Province                        []string
	StreetAddress, PostalCode                 []string
	SerialNumber, CommonName                  string

	// Names contains all parsed attributes. When parsing distinguished names,
	// this can be used to extract non-standard attributes that are not parsed
	// by this package. When marshaling to RDNSequences, the Names field is
	// ignored, see ExtraNames.
	Names []AttributeTypeAndValue

	// ExtraNames contains attributes to be copied, raw, into any marshaled
	// distinguished names. Values override any attributes with the same OID.
	// The ExtraNames field is not populated when parsing, see Names.
	ExtraNames []AttributeTypeAndValue
}
```

Name represents an X.509 distinguished name. This only includes the common elements of a DN. Note that Name is only an approximation of the X.509 structure. If an accurate representation is needed, asn1.Unmarshal the raw subject or issuer as an RDNSequence.

​	Name 表示 X.509 专有名称。这仅包括 DN 的常见元素。请注意，Name 只是 X.509 结构的近似值。如果需要准确的表示，请将原始主题或颁发者作为 RDNSequence 进行 asn1.Unmarshal。

#### (*Name) FillFromRDNSequence

```go
func (n *Name) FillFromRDNSequence(rdns *RDNSequence)
```

FillFromRDNSequence populates n from the provided RDNSequence. Multi-entry RDNs are flattened, all entries are added to the relevant n fields, and the grouping is not preserved.

​	FillFromRDNSequence 从提供的 RDNSequence 填充 n。多条目 RDN 被展平，所有条目都被添加到相关的 n 字段，并且不保留分组。

#### (Name) String <- go1.10

```go
func (n Name) String() string
```

String returns the string form of n, roughly following the [RFC 2253](https://rfc-editor.org/rfc/rfc2253.html) Distinguished Names syntax.

​	String 返回 n 的字符串形式，大致遵循 RFC 2253 专有名称语法。

#### (Name) ToRDNSequence

```go
func (n Name) ToRDNSequence() (ret RDNSequence)
```

ToRDNSequence converts n into a single RDNSequence. The following attributes are encoded as multi-value RDNs:

​	ToRDNSequence 将 n 转换为单个 RDNSequence。以下属性被编码为多值 RDN：

- Country
  国家/地区
- Organization
  组织
- OrganizationalUnit
  组织单位
- Locality
  地区
- Province
  省份
- StreetAddress
  街道地址
- PostalCode
  邮政编码

Each ExtraNames entry is encoded as an individual RDN.

​	每个 ExtraNames 条目都编码为一个单独的 RDN。

### type RDNSequence

```go
type RDNSequence []RelativeDistinguishedNameSET
```

#### (RDNSequence) String <- go1.10

```go
func (r RDNSequence) String() string
```

String returns a string representation of the sequence r, roughly following the [RFC 2253](https://rfc-editor.org/rfc/rfc2253.html) Distinguished Names syntax.

​	String 返回序列 r 的字符串表示形式，大致遵循 RFC 2253 识别名称语法。

### type RelativeDistinguishedNameSET

```go
type RelativeDistinguishedNameSET []AttributeTypeAndValue
```

### type RevokedCertificate

```go
type RevokedCertificate struct {
	SerialNumber   *big.Int
	RevocationTime time.Time
	Extensions     []Extension `asn1:"optional"`
}
```

RevokedCertificate represents the ASN.1 structure of the same name. See [RFC 5280, section 5.1](https://rfc-editor.org/rfc/rfc5280.html#section-5.1).

​	RevokedCertificate 表示同名的 ASN.1 结构。请参阅 RFC 5280 第 5.1 节。

### type TBSCertificateList <-DEPRECATED

``` go
type TBSCertificateList struct {
	Raw                 asn1.RawContent
	Version             int `asn1:"optional,default:0"`
	Signature           AlgorithmIdentifier
	Issuer              RDNSequence
	ThisUpdate          time.Time
	NextUpdate          time.Time            `asn1:"optional"`
	RevokedCertificates []RevokedCertificate `asn1:"optional"`
	Extensions          []Extension          `asn1:"tag:0,optional,explicit"`
}
```

TBSCertificateList represents the ASN.1 structure of the same name. See [RFC 5280, section 5.1](https://rfc-editor.org/rfc/rfc5280.html#section-5.1).

​	TBSCertificateList 表示同名的 ASN.1 结构。请参阅 RFC 5280 第 5.1 节。

Deprecated: x509.RevocationList should be used instead.

​	已弃用：应改用 x509.RevocationList。

