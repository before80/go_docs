+++
title = "x509/pkix"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# pkix

https://pkg.go.dev/crypto/x509/pkix@go1.20.1



Package pkix contains shared, low level structures used for ASN.1 parsing and serialization of X.509 certificates, CRL and OCSP.











  
  






## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type AlgorithmIdentifier 

```
type AlgorithmIdentifier struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.RawValue `asn1:"optional"`
}
```

AlgorithmIdentifier represents the ASN.1 structure of the same name. See [RFC 5280, section 4.1.1.2](https://rfc-editor.org/rfc/rfc5280.html#section-4.1.1.2).

### type AttributeTypeAndValue 

```
type AttributeTypeAndValue struct {
	Type  asn1.ObjectIdentifier
	Value any
}
```

AttributeTypeAndValue mirrors the ASN.1 structure of the same name in [RFC 5280, Section 4.1.2.4](https://rfc-editor.org/rfc/rfc5280.html#section-4.1.2.4).

### type AttributeTypeAndValueSET  <- go1.3

```
type AttributeTypeAndValueSET struct {
	Type  asn1.ObjectIdentifier
	Value [][]AttributeTypeAndValue `asn1:"set"`
}
```

AttributeTypeAndValueSET represents a set of ASN.1 sequences of AttributeTypeAndValue sequences from [RFC 2986](https://rfc-editor.org/rfc/rfc2986.html) (PKCS #10).

#### type CertificateList <- DEPRECATED

### type Extension 

```
type Extension struct {
	Id       asn1.ObjectIdentifier
	Critical bool `asn1:"optional"`
	Value    []byte
}
```

Extension represents the ASN.1 structure of the same name. See [RFC 5280, section 4.2](https://rfc-editor.org/rfc/rfc5280.html#section-4.2).

### type Name 

```
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

#### (*Name) FillFromRDNSequence 

```
func (n *Name) FillFromRDNSequence(rdns *RDNSequence)
```

FillFromRDNSequence populates n from the provided RDNSequence. Multi-entry RDNs are flattened, all entries are added to the relevant n fields, and the grouping is not preserved.

#### (Name) String  <- go1.10

```
func (n Name) String() string
```

String returns the string form of n, roughly following the [RFC 2253](https://rfc-editor.org/rfc/rfc2253.html) Distinguished Names syntax.

#### (Name) ToRDNSequence 

```
func (n Name) ToRDNSequence() (ret RDNSequence)
```

ToRDNSequence converts n into a single RDNSequence. The following attributes are encoded as multi-value RDNs:

- Country
- Organization
- OrganizationalUnit
- Locality
- Province
- StreetAddress
- PostalCode

Each ExtraNames entry is encoded as an individual RDN.

### type RDNSequence 

```
type RDNSequence []RelativeDistinguishedNameSET
```

#### (RDNSequence) String  <- go1.10

```
func (r RDNSequence) String() string
```

String returns a string representation of the sequence r, roughly following the [RFC 2253](https://rfc-editor.org/rfc/rfc2253.html) Distinguished Names syntax.

### type RelativeDistinguishedNameSET 

```
type RelativeDistinguishedNameSET []AttributeTypeAndValue
```

### type RevokedCertificate 

```
type RevokedCertificate struct {
	SerialNumber   *big.Int
	RevocationTime time.Time
	Extensions     []Extension `asn1:"optional"`
}
```

RevokedCertificate represents the ASN.1 structure of the same name. See [RFC 5280, section 5.1](https://rfc-editor.org/rfc/rfc5280.html#section-5.1).



