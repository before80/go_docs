+++
title = "tls"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/tls@go1.24.2](https://pkg.go.dev/crypto/tls@go1.24.2)

Package tls partially implements TLS 1.2, as specified in [RFC 5246](https://rfc-editor.org/rfc/rfc5246.html), and TLS 1.3, as specified in [RFC 8446](https://rfc-editor.org/rfc/rfc8446.html).

​	tls 包部分实现了 TLS 1.2，如 [RFC 5246](https://rfc-editor.org/rfc/rfc5246.html) 中所述，以及 TLS 1.3，如 [RFC 8446](https://rfc-editor.org/rfc/rfc8446.html) 中所述。


## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/tls/cipher_suites.go;l=664)

``` go
const (
	// TLS 1.0 - 1.2 cipher suites.
	TLS_RSA_WITH_RC4_128_SHA                      uint16 = 0x0005
	TLS_RSA_WITH_3DES_EDE_CBC_SHA                 uint16 = 0x000a
	TLS_RSA_WITH_AES_128_CBC_SHA                  uint16 = 0x002f
	TLS_RSA_WITH_AES_256_CBC_SHA                  uint16 = 0x0035
	TLS_RSA_WITH_AES_128_CBC_SHA256               uint16 = 0x003c
	TLS_RSA_WITH_AES_128_GCM_SHA256               uint16 = 0x009c
	TLS_RSA_WITH_AES_256_GCM_SHA384               uint16 = 0x009d
	TLS_ECDHE_ECDSA_WITH_RC4_128_SHA              uint16 = 0xc007
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA          uint16 = 0xc009
	TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA          uint16 = 0xc00a
	TLS_ECDHE_RSA_WITH_RC4_128_SHA                uint16 = 0xc011
	TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA           uint16 = 0xc012
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA            uint16 = 0xc013
	TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA            uint16 = 0xc014
	TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256       uint16 = 0xc023
	TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256         uint16 = 0xc027
	TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256         uint16 = 0xc02f
	TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256       uint16 = 0xc02b
	TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384         uint16 = 0xc030
	TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384       uint16 = 0xc02c
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256   uint16 = 0xcca8
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256 uint16 = 0xcca9

	// TLS 1.3 cipher suites.
	TLS_AES_128_GCM_SHA256       uint16 = 0x1301
	TLS_AES_256_GCM_SHA384       uint16 = 0x1302
	TLS_CHACHA20_POLY1305_SHA256 uint16 = 0x1303

	// TLS_FALLBACK_SCSV isn't a standard cipher suite but an indicator
	// that the client is doing version fallback. See RFC 7507.
	TLS_FALLBACK_SCSV uint16 = 0x5600

	// Legacy names for the corresponding cipher suites with the correct _SHA256
	// suffix, retained for backward compatibility.
	TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305   = TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256
	TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305 = TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
)
```

A list of cipher suite IDs that are, or have been, implemented by this package.

​	此软件包实现或曾经实现的密码套件 ID 列表。

See https://www.iana.org/assignments/tls-parameters/tls-parameters.xml

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/tls/common.go;l=28)

``` go
const (
	VersionTLS10 = 0x0301
	VersionTLS11 = 0x0302
	VersionTLS12 = 0x0303
	VersionTLS13 = 0x0304

	// Deprecated: SSLv3 is cryptographically broken, and is no longer
	// supported by this package. See golang.org/issue/32716.
	VersionSSL30 = 0x0300
)
```

## 变量

This section is empty.

## 函数

### func CipherSuiteName  <- go1.14

``` go
func CipherSuiteName(id uint16) string
```

CipherSuiteName returns the standard name for the passed cipher suite ID (e.g. "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"), or a fallback representation of the ID value if the cipher suite is not implemented by this package.

​	CipherSuiteName 返回传递的密码套件 ID 的标准名称（例如 “TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256”），或者如果此软件包未实现密码套件，则返回 ID 值的备用表示形式。

### func Listen 

``` go
func Listen(network, laddr string, config *Config) (net.Listener, error)
```

Listen creates a TLS listener accepting connections on the given network address using net.Listen. The configuration config must be non-nil and must include at least one certificate or else set GetCertificate.

​	Listen 使用 net.Listen 创建一个 TLS 侦听器，该侦听器接受给定网络地址上的连接。配置 config 必须为非 nil，并且必须至少包含一个证书，否则设置 GetCertificate。

### func NewListener 

``` go
func NewListener(inner net.Listener, config *Config) net.Listener
```

NewListener creates a Listener which accepts connections from an inner Listener and wraps each connection with Server. The configuration config must be non-nil and must include at least one certificate or else set GetCertificate.

​	NewListener 创建一个侦听器，该侦听器接受来自内部侦听器的连接，并将每个连接包装到 Server 中。配置 config 必须为非 nil，并且必须至少包含一个证书，否则设置 GetCertificate。

### func VersionName <- go1.21.0

```go
func VersionName(version uint16) string
```

VersionName returns the name for the provided TLS version number (e.g. "TLS 1.3"), or a fallback representation of the value if the version is not implemented by this package.

​	VersionName 返回所提供的 TLS 版本号的名称（例如 "TLS 1.3"），如果该版本未被此包实现，则返回该值的替代表示形式。

## 类型

### type AlertError <-go1.21.0

```go
type AlertError uint8
```

An AlertError is a TLS alert.

When using a QUIC transport, QUICConn methods will return an error which wraps AlertError rather than sending a TLS alert.

#### (AlertError) Error <-go1.21.0

```go
func (e AlertError) Error() string
```

### type Certificate 

``` go
type Certificate struct {
	Certificate [][]byte
	// PrivateKey contains the private key corresponding to the public key in
	// Leaf. This must implement crypto.Signer with an RSA, ECDSA or Ed25519 PublicKey.
	// For a server up to TLS 1.2, it can also implement crypto.Decrypter with
	// an RSA PublicKey.
	PrivateKey crypto.PrivateKey
	// SupportedSignatureAlgorithms is an optional list restricting what
	// signature algorithms the PrivateKey can be used for.
	SupportedSignatureAlgorithms []SignatureScheme
	// OCSPStaple contains an optional OCSP response which will be served
	// to clients that request it.
	OCSPStaple []byte
	// SignedCertificateTimestamps contains an optional list of Signed
	// Certificate Timestamps which will be served to clients that request it.
	SignedCertificateTimestamps [][]byte
	// Leaf is the parsed form of the leaf certificate, which may be initialized
	// using x509.ParseCertificate to reduce per-handshake processing. If nil,
	// the leaf certificate will be parsed as needed.
	Leaf *x509.Certificate
}
```

A Certificate is a chain of one or more certificates, leaf first.

​	Certificate 是由一个或多个证书（叶证书排在第一位）组成的链。

#### func LoadX509KeyPair 

``` go
func LoadX509KeyPair(certFile, keyFile string) (Certificate, error)
```

LoadX509KeyPair reads and parses a public/private key pair from a pair of files. The files must contain PEM encoded data. The certificate file may contain intermediate certificates following the leaf certificate to form a certificate chain. On successful return, Certificate.Leaf will be nil because the parsed form of the certificate is not retained.

​	LoadX509KeyPair 从一对文件中读取和解析公钥/私钥对。这些文件必须包含 PEM 编码的数据。证书文件可能包含叶证书后面的中间证书，以形成证书链。成功返回后，Certificate.Leaf 将为 nil，因为不会保留证书的解析形式。

Before Go 1.23 Certificate.Leaf was left nil, and the parsed certificate was discarded. This behavior can be re-enabled by setting "x509keypairleaf=0" in the GODEBUG environment variable.

​	Go 1.23 之前Certificate.Leaf为nil，并丢弃已解析的证书。可以通过在GODEBUG环境变量中设置 "x509keypairleaf=0"来重新启用此行为。

##### LoadX509KeyPair Example

```go
package main

import (
	"crypto/tls"
	"log"
)

func main() {
	cert, err := tls.LoadX509KeyPair("testdata/example-cert.pem", "testdata/example-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":2000", cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = listener
}

Output:
```



#### func X509KeyPair 

``` go
func X509KeyPair(certPEMBlock, keyPEMBlock []byte) (Certificate, error)
```

X509KeyPair parses a public/private key pair from a pair of PEM encoded data. On successful return, Certificate.Leaf will be nil because the parsed form of the certificate is not retained.

​	X509KeyPair 从一对 PEM 编码数据中解析公钥/私钥对。成功返回后，Certificate.Leaf 将为 nil，因为不会保留证书的解析形式。

Before Go 1.23 Certificate.Leaf was left nil, and the parsed certificate was discarded. This behavior can be re-enabled by setting "x509keypairleaf=0" in the GODEBUG environment variable.

​	Go 1.23 之前Certificate.Leaf为nil，并丢弃已解析的证书。可以通过在GODEBUG环境变量中设置 "x509keypairleaf=0"来重新启用此行为。

##### X509KeyPair Example

```go
package main

import (
	"crypto/tls"
	"log"
)

func main() {
	certPem := []byte(`-----BEGIN CERTIFICATE-----
MIIBhTCCASugAwIBAgIQIRi6zePL6mKjOipn+dNuaTAKBggqhkjOPQQDAjASMRAw
DgYDVQQKEwdBY21lIENvMB4XDTE3MTAyMDE5NDMwNloXDTE4MTAyMDE5NDMwNlow
EjEQMA4GA1UEChMHQWNtZSBDbzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD0d
7VNhbWvZLWPuj/RtHFjvtJBEwOkhbN/BnnE8rnZR8+sbwnc/KhCk3FhnpHZnQz7B
5aETbbIgmuvewdjvSBSjYzBhMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggr
BgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdEQQiMCCCDmxvY2FsaG9zdDo1
NDUzgg4xMjcuMC4wLjE6NTQ1MzAKBggqhkjOPQQDAgNIADBFAiEA2zpJEPQyz6/l
Wf86aX6PepsntZv2GYlA5UpabfT2EZICICpJ5h/iI+i341gBmLiAFQOyTDT+/wQc
6MF9+Yw1Yy0t
-----END CERTIFICATE-----`)
	keyPem := []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIrYSSNQFaA2Hwf1duRSxKtLYX5CB04fSeQ6tF1aY/PuoAoGCCqGSM49
AwEHoUQDQgAEPR3tU2Fta9ktY+6P9G0cWO+0kETA6SFs38GecTyudlHz6xvCdz8q
EKTcWGekdmdDPsHloRNtsiCa697B2O9IFA==
-----END EC PRIVATE KEY-----`)
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":2000", cfg)
	if err != nil {
		log.Fatal(err)
	}
	_ = listener
}

```



##### X509KeyPair Example (HttpServer)

```go
package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func main() {
	certPem := []byte(`-----BEGIN CERTIFICATE-----
MIIBhTCCASugAwIBAgIQIRi6zePL6mKjOipn+dNuaTAKBggqhkjOPQQDAjASMRAw
DgYDVQQKEwdBY21lIENvMB4XDTE3MTAyMDE5NDMwNloXDTE4MTAyMDE5NDMwNlow
EjEQMA4GA1UEChMHQWNtZSBDbzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD0d
7VNhbWvZLWPuj/RtHFjvtJBEwOkhbN/BnnE8rnZR8+sbwnc/KhCk3FhnpHZnQz7B
5aETbbIgmuvewdjvSBSjYzBhMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggr
BgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdEQQiMCCCDmxvY2FsaG9zdDo1
NDUzgg4xMjcuMC4wLjE6NTQ1MzAKBggqhkjOPQQDAgNIADBFAiEA2zpJEPQyz6/l
Wf86aX6PepsntZv2GYlA5UpabfT2EZICICpJ5h/iI+i341gBmLiAFQOyTDT+/wQc
6MF9+Yw1Yy0t
-----END CERTIFICATE-----`)
	keyPem := []byte(`-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIrYSSNQFaA2Hwf1duRSxKtLYX5CB04fSeQ6tF1aY/PuoAoGCCqGSM49
AwEHoUQDQgAEPR3tU2Fta9ktY+6P9G0cWO+0kETA6SFs38GecTyudlHz6xvCdz8q
EKTcWGekdmdDPsHloRNtsiCa697B2O9IFA==
-----END EC PRIVATE KEY-----`)
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	srv := &http.Server{
		TLSConfig:    cfg,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	log.Fatal(srv.ListenAndServeTLS("", ""))
}

Output:
```



### type CertificateRequestInfo  <- go1.8

``` go
type CertificateRequestInfo struct {
	// AcceptableCAs contains zero or more, DER-encoded, X.501
	// Distinguished Names. These are the names of root or intermediate CAs
	// that the server wishes the returned certificate to be signed by. An
	// empty slice indicates that the server has no preference.
	AcceptableCAs [][]byte

	// SignatureSchemes lists the signature schemes that the server is
	// willing to verify.
	SignatureSchemes []SignatureScheme

	// Version is the TLS version that was negotiated for this connection.
	Version uint16
	// contains filtered or unexported fields
}

```

CertificateRequestInfo contains information from a server's CertificateRequest message, which is used to demand a certificate and proof of control from a client.

​	CertificateRequestInfo 包含来自服务器的 CertificateRequest 消息的信息，该消息用于向客户端索要证书和控制证明。

#### (*CertificateRequestInfo) Context  <- go1.17

``` go
func (c *CertificateRequestInfo) Context() context.Context
```

Context returns the context of the handshake that is in progress. This context is a child of the context passed to HandshakeContext, if any, and is canceled when the handshake concludes.

​	Context 返回正在进行的握手的上下文。此上下文是传递给 HandshakeContext 的上下文的子级（如果有），并在握手结束时取消。

#### (*CertificateRequestInfo) SupportsCertificate  <- go1.14

``` go
func (cri *CertificateRequestInfo) SupportsCertificate(c *Certificate) error
```

SupportsCertificate returns nil if the provided certificate is supported by the server that sent the CertificateRequest. Otherwise, it returns an error describing the reason for the incompatibility.

如果服务器发送的证书受支持，SupportsCertificate 返回 nil。否则，它会返回一个描述不兼容原因的错误。

### type CertificateVerificationError  <- go1.20

``` go
type CertificateVerificationError struct {
	// UnverifiedCertificates and its contents should not be modified.
	UnverifiedCertificates []*x509.Certificate
	Err                    error
}
```

CertificateVerificationError is returned when certificate verification fails during the handshake.

​	当握手期间证书验证失败时，会返回 CertificateVerificationError。

#### (*CertificateVerificationError) Error  <- go1.20

``` go
func (e *CertificateVerificationError) Error() string
```

#### (*CertificateVerificationError) Unwrap  <- go1.20

``` go
func (e *CertificateVerificationError) Unwrap() error
```

### type CipherSuite  <- go1.14

``` go
type CipherSuite struct {
	ID   uint16
	Name string

	// Supported versions is the list of TLS protocol versions that can
	// negotiate this cipher suite.
	SupportedVersions []uint16

	// Insecure is true if the cipher suite has known security issues
	// due to its primitives, design, or implementation.
	Insecure bool
}
```

CipherSuite is a TLS cipher suite. Note that most functions in this package accept and expose cipher suite IDs instead of this type.

​	CipherSuite 是一个 TLS 密码套件。请注意，此程序包中的大多数函数接受并公开密码套件 ID，而不是此类型。

#### func CipherSuites  <- go1.14

``` go
func CipherSuites() []*CipherSuite
```

CipherSuites returns a list of cipher suites currently implemented by this package, excluding those with security issues, which are returned by InsecureCipherSuites.

​	CipherSuites 返回此程序包当前实现的密码套件列表，不包括存在安全问题的密码套件，这些密码套件由 InsecureCipherSuites 返回。

The list is sorted by ID. Note that the default cipher suites selected by this package might depend on logic that can't be captured by a static list, and might not match those returned by this function.

​	该列表按 ID 排序。请注意，此软件包选择的默认密码套件可能取决于无法通过静态列表捕获的逻辑，并且可能与此函数返回的密码套件不匹配。

#### func InsecureCipherSuites  <- go1.14

``` go
func InsecureCipherSuites() []*CipherSuite
```

InsecureCipherSuites returns a list of cipher suites currently implemented by this package and which have security issues.

​	InsecureCipherSuites 返回此软件包当前实现且存在安全问题的密码套件列表。

Most applications should not use the cipher suites in this list, and should only use those returned by CipherSuites.

​	大多数应用程序不应使用此列表中的密码套件，而应仅使用 CipherSuites 返回的密码套件。

### type ClientAuthType 

``` go
type ClientAuthType int
```

ClientAuthType declares the policy the server will follow for TLS Client Authentication.

​	ClientAuthType 声明服务器将遵循的 TLS 客户端身份验证策略。

``` go
const (
	// NoClientCert indicates that no client certificate should be requested
	// during the handshake, and if any certificates are sent they will not
	// be verified.
	NoClientCert ClientAuthType = iota
	// RequestClientCert indicates that a client certificate should be requested
	// during the handshake, but does not require that the client send any
	// certificates.
	RequestClientCert
	// RequireAnyClientCert indicates that a client certificate should be requested
	// during the handshake, and that at least one certificate is required to be
	// sent by the client, but that certificate is not required to be valid.
	RequireAnyClientCert
	// VerifyClientCertIfGiven indicates that a client certificate should be requested
	// during the handshake, but does not require that the client sends a
	// certificate. If the client does send a certificate it is required to be
	// valid.
	VerifyClientCertIfGiven
	// RequireAndVerifyClientCert indicates that a client certificate should be requested
	// during the handshake, and that at least one valid certificate is required
	// to be sent by the client.
	RequireAndVerifyClientCert
)
```

#### (ClientAuthType) String  <- go1.15

``` go
func (i ClientAuthType) String() string
```

### type ClientHelloInfo  <- go1.4

``` go
type ClientHelloInfo struct {
	// CipherSuites lists the CipherSuites supported by the client (e.g.
	// TLS_AES_128_GCM_SHA256, TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256).
	CipherSuites []uint16

	// ServerName indicates the name of the server requested by the client
	// in order to support virtual hosting. ServerName is only set if the
	// client is using SNI (see RFC 4366, Section 3.1).
	ServerName string

	// SupportedCurves lists the elliptic curves supported by the client.
	// SupportedCurves is set only if the Supported Elliptic Curves
	// Extension is being used (see RFC 4492, Section 5.1.1).
	SupportedCurves []CurveID

	// SupportedPoints lists the point formats supported by the client.
	// SupportedPoints is set only if the Supported Point Formats Extension
	// is being used (see RFC 4492, Section 5.1.2).
	SupportedPoints []uint8

	// SignatureSchemes lists the signature and hash schemes that the client
	// is willing to verify. SignatureSchemes is set only if the Signature
	// Algorithms Extension is being used (see RFC 5246, Section 7.4.1.4.1).
	SignatureSchemes []SignatureScheme

	// SupportedProtos lists the application protocols supported by the client.
	// SupportedProtos is set only if the Application-Layer Protocol
	// Negotiation Extension is being used (see RFC 7301, Section 3.1).
	//
	// Servers can select a protocol by setting Config.NextProtos in a
	// GetConfigForClient return value.
	SupportedProtos []string

	// SupportedVersions lists the TLS versions supported by the client.
	// For TLS versions less than 1.3, this is extrapolated from the max
	// version advertised by the client, so values other than the greatest
	// might be rejected if used.
	SupportedVersions []uint16

	// Conn is the underlying net.Conn for the connection. Do not read
	// from, or write to, this connection; that will cause the TLS
	// connection to fail.
	Conn net.Conn
	// contains filtered or unexported fields
}
```

ClientHelloInfo contains information from a ClientHello message in order to guide application logic in the GetCertificate and GetConfigForClient callbacks.

​	ClientHelloInfo 包含 ClientHello 消息中的信息，以便在 GetCertificate 和 GetConfigForClient 回调中指导应用程序逻辑。

#### (*ClientHelloInfo) Context  <- go1.17

``` go
func (c *ClientHelloInfo) Context() context.Context
```

Context returns the context of the handshake that is in progress. This context is a child of the context passed to HandshakeContext, if any, and is canceled when the handshake concludes.

​	Context 返回正在进行的握手的上下文。此上下文是传递给 HandshakeContext 的上下文的子级（如果存在），并且在握手结束时取消。

#### (*ClientHelloInfo) SupportsCertificate  <- go1.14

``` go
func (chi *ClientHelloInfo) SupportsCertificate(c *Certificate) error
```

SupportsCertificate returns nil if the provided certificate is supported by the client that sent the ClientHello. Otherwise, it returns an error describing the reason for the incompatibility.

​	如果发送 ClientHello 的客户端支持提供的证书，则 SupportsCertificate 返回 nil。否则，它将返回描述不兼容原因的错误。

If this ClientHelloInfo was passed to a GetConfigForClient or GetCertificate callback, this method will take into account the associated Config. Note that if GetConfigForClient returns a different Config, the change can't be accounted for by this method.

​	如果此 ClientHelloInfo 传递给 GetConfigForClient 或 GetCertificate 回调，此方法将考虑关联的 Config。请注意，如果 GetConfigForClient 返回不同的 Config，则此方法无法解释该更改。

This function will call x509.ParseCertificate unless c.Leaf is set, which can incur a significant performance cost.

​	此函数将调用 x509.ParseCertificate，除非设置了 c.Leaf，这可能会产生巨大的性能开销。

### type ClientSessionCache  <- go1.3

``` go
type ClientSessionCache interface {
	// Get searches for a ClientSessionState associated with the given key.
	// On return, ok is true if one was found.
	Get(sessionKey string) (session *ClientSessionState, ok bool)

	// Put adds the ClientSessionState to the cache with the given key. It might
	// get called multiple times in a connection if a TLS 1.3 server provides
	// more than one session ticket. If called with a nil *ClientSessionState,
	// it should remove the cache entry.
	Put(sessionKey string, cs *ClientSessionState)
}
```

ClientSessionCache is a cache of ClientSessionState objects that can be used by a client to resume a TLS session with a given server. ClientSessionCache implementations should expect to be called concurrently from different goroutines. Up to TLS 1.2, only ticket-based resumption is supported, not SessionID-based resumption. In TLS 1.3 they were merged into PSK modes, which are supported via this interface.

​	ClientSessionCache 是一个 ClientSessionState 对象的缓存，客户端可以使用它来恢复与给定服务器的 TLS 会话。ClientSessionCache 实现应该期望从不同的 goroutine 并发调用。在 TLS 1.2 中，仅支持基于票证的恢复，不支持基于 SessionID 的恢复。在 TLS 1.3 中，它们被合并到 PSK 模式中，该模式通过此接口支持。

#### func NewLRUClientSessionCache  <- go1.3

``` go
func NewLRUClientSessionCache(capacity int) ClientSessionCache
```

NewLRUClientSessionCache returns a ClientSessionCache with the given capacity that uses an LRU strategy. If capacity is < 1, a default capacity is used instead.

​	NewLRUClientSessionCache 返回一个具有给定容量的 ClientSessionCache，该容量使用 LRU 策略。如果容量 < 1，则使用默认容量。

### type ClientSessionState  <- go1.3

``` go
type ClientSessionState struct {
	// contains filtered or unexported fields
}
```

ClientSessionState contains the state needed by clients to resume TLS sessions.

​	ClientSessionState 包含客户端恢复 TLS 会话所需的状态。

#### func NewResumptionState <-go1.21.0

``` go
func NewResumptionState(ticket []byte, state *SessionState) (*ClientSessionState, error)
```

NewResumptionState returns a state value that can be returned by `ClientSessionCache.Get` to resume a previous session.

​	`NewResumptionState` 返回一个状态值，该状态值可由`ClientSessionCache.Get`返回，以恢复先前的会话。

state needs to be returned by [ParseSessionState](https://pkg.go.dev/crypto/tls@go1.23.0#ParseSessionState), and the ticket and session state must have been returned by [ClientSessionState.ResumptionState](https://pkg.go.dev/crypto/tls@go1.23.0#ClientSessionState.ResumptionState).

​	`state` 需要由 [ParseSessionState](https://pkg.go.dev/crypto/tls@go1.23.0#ParseSessionState) 返回，而 `ticket` 和会话状态必须由 [ClientSessionState.ResumptionState](https://pkg.go.dev/crypto/tls@go1.23.0#ClientSessionState.ResumptionState) 返回。

#### (*ClientSessionState) ResumptionState <-go1.21.0

``` go
func (cs *ClientSessionState) ResumptionState() (ticket []byte, state *SessionState, err error)
```

ResumptionState returns the session ticket sent by the server (also known as the session's identity) and the state necessary to resume this session.

​	`ResumptionState` 返回由服务器发送的会话票据（也称为会话的标识符）以及恢复此会话所需的状态。

It can be called by `ClientSessionCache.Put`to serialize (with [SessionState.Bytes](https://pkg.go.dev/crypto/tls@go1.23.0#SessionState.Bytes)) and store the session.

​	它可以被 `ClientSessionCache.Put` 调用，以序列化（使用 [SessionState.Bytes](https://pkg.go.dev/crypto/tls@go1.23.0#SessionState.Bytes)）并存储会话。

### type Config 

``` go
type Config struct {
	// Rand provides the source of entropy for nonces and RSA blinding.
	// If Rand is nil, TLS uses the cryptographic random reader in package
	// crypto/rand.
	// The Reader must be safe for use by multiple goroutines.
    // Rand 提供用于生成随机数和 RSA 盲签的熵源。
	// 如果 Rand 为 nil，TLS 将使用 `crypto/rand` 包中的加密随机数生成器。
	// 该 Reader 必须能够安全地被多个 goroutines 使用。
	Rand io.Reader

	// Time returns the current time as the number of seconds since the epoch.
	// If Time is nil, TLS uses time.Now.
    // Time 返回从 epoch 开始的秒数作为当前时间。
	// 如果 Time 为 nil，TLS 将使用 time.Now。
	Time func() time.Time

	// Certificates contains one or more certificate chains to present to the
	// other side of the connection. The first certificate compatible with the
	// peer's requirements is selected automatically.
	//
	// Server configurations must set one of Certificates, GetCertificate or
	// GetConfigForClient. Clients doing client-authentication may set either
	// Certificates or GetClientCertificate.
	//
	// Note: if there are multiple Certificates, and they don't have the
	// optional field Leaf set, certificate selection will incur a significant
	// per-handshake performance cost.
    // Certificates 包含一个或多个要向对方连接展示的证书链。
	// 系统会自动选择与对方要求兼容的第一个证书。
	//
	// 服务器配置必须设置 `Certificates`、`GetCertificate` 或 `GetConfigForClient` 中的一个。
	// 执行客户端认证的客户端可以设置 `Certificates` 或 `GetClientCertificate` 中的一个。
	//
	// 注意：如果有多个证书，且它们未设置可选字段 Leaf，证书选择操作将在每次握手中导致显著的性能消耗。
	Certificates []Certificate

	// NameToCertificate maps from a certificate name to an element of
	// Certificates. Note that a certificate name can be of the form
	// '*.example.com' and so doesn't have to be a domain name as such.
	//
	// Deprecated: NameToCertificate only allows associating a single
	// certificate with a given name. Leave this field nil to let the library
	// select the first compatible chain from Certificates.
    // NameToCertificate 映射证书名称到 Certificates 中的元素。
	// 请注意，证书名称可以是 '*.example.com' 形式，因此不必一定是域名。
	//
	// 已弃用：NameToCertificate 只允许将单个证书与给定名称关联。
	// 将此字段留空以让库从 Certificates 中选择第一个兼容链。
	NameToCertificate map[string]*Certificate

	// GetCertificate returns a Certificate based on the given
	// ClientHelloInfo. It will only be called if the client supplies SNI
	// information or if Certificates is empty.
	//
	// If GetCertificate is nil or returns nil, then the certificate is
	// retrieved from NameToCertificate. If NameToCertificate is nil, the
	// best element of Certificates will be used.
	//
	// Once a Certificate is returned it should not be modified.
    // GetCertificate 根据给定的 ClientHelloInfo 返回证书。
	// 只有在客户端提供 SNI 信息或 Certificates 为空时才会调用此函数。
	//
	// 如果 GetCertificate 为 nil 或返回 nil，则从 NameToCertificate 检索证书。
	// 如果 NameToCertificate 为 nil，则会使用 Certificates 中的最佳元素。
	//
	// 一旦返回证书，它就不应被修改。
	GetCertificate func(*ClientHelloInfo) (*Certificate, error)

	// GetClientCertificate, if not nil, is called when a server requests a
	// certificate from a client. If set, the contents of Certificates will
	// be ignored.
	//
	// If GetClientCertificate returns an error, the handshake will be
	// aborted and that error will be returned. Otherwise
	// GetClientCertificate must return a non-nil Certificate. If
	// Certificate.Certificate is empty then no certificate will be sent to
	// the server. If this is unacceptable to the server then it may abort
	// the handshake.
	//
	// GetClientCertificate may be called multiple times for the same
	// connection if renegotiation occurs or if TLS 1.3 is in use.
	//
	// Once a Certificate is returned it should not be modified.
    // GetClientCertificate，如果不为 nil，当服务器请求客户端提供证书时调用。
	// 如果设置了此选项，则会忽略 Certificates 的内容。
	//
	// 如果 GetClientCertificate 返回错误，握手将中止并返回该错误。
	// 否则，GetClientCertificate 必须返回非空证书。
	// 如果 Certificate.Certificate 为空，则不会向服务器发送证书。
	// 如果服务器无法接受这一点，它可能会中止握手。
	//
	// 如果发生重新协商或使用 TLS 1.3，则可能会为同一连接多次调用 GetClientCertificate。
	//
	// 一旦返回证书，它就不应被修改。
	GetClientCertificate func(*CertificateRequestInfo) (*Certificate, error)

	// GetConfigForClient, if not nil, is called after a ClientHello is
	// received from a client. It may return a non-nil Config in order to
	// change the Config that will be used to handle this connection. If
	// the returned Config is nil, the original Config will be used. The
	// Config returned by this callback may not be subsequently modified.
	//
	// If GetConfigForClient is nil, the Config passed to Server() will be
	// used for all connections.
	//
	// If SessionTicketKey was explicitly set on the returned Config, or if
	// SetSessionTicketKeys was called on the returned Config, those keys will
	// be used. Otherwise, the original Config keys will be used (and possibly
	// rotated if they are automatically managed).
    // GetConfigForClient，如果不为 nil，则在从客户端接收到 ClientHello 后调用。
	// 它可能会返回非空 Config 以更改将用于处理此连接的 Config。
	// 如果返回的 Config 为 nil，则将使用原始 Config。
	// 由此回调返回的 Config 不能在之后被修改。
	//
	// 如果 GetConfigForClient 为 nil，则 Server() 传递的 Config 将用于所有连接。
	//
	// 如果在返回的 Config 上显式设置了 SessionTicketKey，或者调用了 SetSessionTicketKeys，
	// 则将使用这些密钥。否则，将使用原始 Config 密钥（如果它们是自动管理的，可能会旋转）。
	GetConfigForClient func(*ClientHelloInfo) (*Config, error)

	// VerifyPeerCertificate, if not nil, is called after normal
	// certificate verification by either a TLS client or server. It
	// receives the raw ASN.1 certificates provided by the peer and also
	// any verified chains that normal processing found. If it returns a
	// non-nil error, the handshake is aborted and that error results.
	//
	// If normal verification fails then the handshake will abort before
	// considering this callback. If normal verification is disabled (on the
	// client when InsecureSkipVerify is set, or on a server when ClientAuth is
	// RequestClientCert or RequireAnyClientCert), then this callback will be
	// considered but the verifiedChains argument will always be nil. When
	// ClientAuth is NoClientCert, this callback is not called on the server.
	// rawCerts may be empty on the server if ClientAuth is RequestClientCert or
	// VerifyClientCertIfGiven.
	//
	// This callback is not invoked on resumed connections, as certificates are
	// not re-verified on resumption.
	//
	// verifiedChains and its contents should not be modified.
    // VerifyPeerCertificate，如果不为 nil，则在 TLS 客户端或服务器的正常证书验证之后调用。
	// 它接收由对等方提供的原始 ASN.1 证书，以及正常处理过程中发现的任何已验证链。
	// 如果返回非空错误，则握手中止并导致该错误。
	//
	// 如果正常验证失败，握手将在考虑此回调之前中止。
	// 如果正常验证被禁用（当客户端设置了 InsecureSkipVerify，或服务器的 ClientAuth 为 RequestClientCert 或 RequireAnyClientCert），
	// 则会考虑此回调，但 verifiedChains 参数将始终为 nil。
	// 当 ClientAuth 为 NoClientCert 时，服务器不会调用此回调。
	// 如果 ClientAuth 是 RequestClientCert 或 VerifyClientCertIfGiven，服务器上的 rawCerts 可能为空。
	//
	// 在恢复的连接上不会调用此回调，因为恢复时不会重新验证证书。
	//
	// verifiedChains 及其内容不应被修改。
	VerifyPeerCertificate func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

	// VerifyConnection, if not nil, is called after normal certificate
	// verification and after VerifyPeerCertificate by either a TLS client
	// or server. If it returns a non-nil error, the handshake is aborted
	// and that error results.
	//
	// If normal verification fails then the handshake will abort before
	// considering this callback. This callback will run for all connections,
	// including resumptions, regardless of InsecureSkipVerify or ClientAuth
	// settings.
   // VerifyConnection，如果不为 nil，则在正常证书验证之后以及 VerifyPeerCertificate 之后由 TLS 客户端或服务器调用。
	// 如果返回非空错误，则握手中止并导致该错误。
	//
	// 如果正常验证失败，握手将在考虑此回调之前中止。
	// 此回调将为所有连接运行，包括恢复连接，无论 InsecureSkipVerify 或 ClientAuth 设置如何。
	VerifyConnection func(ConnectionState) error

	// RootCAs defines the set of root certificate authorities
	// that clients use when verifying server certificates.
	// If RootCAs is nil, TLS uses the host's root CA set.
    // RootCAs 定义客户端在验证服务器证书时使用的根证书颁发机构集合。
	// 如果 RootCAs 为 nil，则 TLS 使用主机的根 CA 集合。
	RootCAs *x509.CertPool

	// NextProtos is a list of supported application level protocols, in
	// order of preference. If both peers support ALPN, the selected
	// protocol will be one from this list, and the connection will fail
	// if there is no mutually supported protocol. If NextProtos is empty
	// or the peer doesn't support ALPN, the connection will succeed and
	// ConnectionState.NegotiatedProtocol will be empty.
    // NextProtos 是支持的应用层协议的列表，按优先顺序排列。
	// 如果双方都支持 ALPN，则选择的协议将是此列表中的一个，并且如果没有共同支持的协议，连接将失败。
	// 如果 NextProtos 为空或对等方不支持 ALPN，连接将成功并且 ConnectionState.NegotiatedProtocol 将为空。
	NextProtos []string

	// ServerName is used to verify the hostname on the returned
	// certificates unless InsecureSkipVerify is given. It is also included
	// in the client's handshake to support virtual hosting unless it is
	// an IP address.
    // ServerName 用于验证返回的证书上的主机名，除非设置了 InsecureSkipVerify。
	// 它还包含在客户端的握手中，以支持虚拟主机，除非它是 IP 地址。
	ServerName string

	// ClientAuth determines the server's policy for
	// TLS Client Authentication. The default is NoClientCert.
    // ClientAuth 决定服务器的 TLS 客户端认证策略。默认值为 NoClientCert。
	ClientAuth ClientAuthType

	// ClientCAs defines the set of root certificate authorities
	// that servers use if required to verify a client certificate
	// by the policy in ClientAuth.
    // ClientCAs 定义服务器在根据 ClientAuth 策略需要验证客户端证书时使用的根证书颁发机构集合。
	ClientCAs *x509.CertPool

	// InsecureSkipVerify controls whether a client verifies the server's
	// certificate chain and host name. If InsecureSkipVerify is true, crypto/tls
	// accepts any certificate presented by the server and any host name in that
	// certificate. In this mode, TLS is susceptible to machine-in-the-middle
	// attacks unless custom verification is used. This should be used only for
	// testing or in combination with VerifyConnection or VerifyPeerCertificate.
    // InsecureSkipVerify 控制客户端是否验证服务器的证书链和主机名。
	// 如果 InsecureSkipVerify 为 true，crypto/tls 将接受服务器提供的任何证书和证书中的任何主机名。
	// 在此模式下，除非使用自定义验证，否则 TLS 易受中间人攻击。
	// 仅应在测试中或与 VerifyConnection 或 VerifyPeerCertificate 结合使用时使用。
	InsecureSkipVerify bool

	// CipherSuites is a list of enabled TLS 1.0–1.2 cipher suites. The order of
	// the list is ignored. Note that TLS 1.3 ciphersuites are not configurable.
	//
	// If CipherSuites is nil, a safe default list is used. The default cipher
	// suites might change over time. In Go 1.22 RSA key exchange based cipher
	// suites were removed from the default list, but can be re-added with the
	// GODEBUG setting tlsrsakex=1. In Go 1.23 3DES cipher suites were removed
	// from the default list, but can be re-added with the GODEBUG setting
	// tls3des=1.
    // CipherSuites 是已启用的 TLS 1.0–1.2 密码套件列表。列表的顺序将被忽略。
	// 请注意，TLS 1.3 密码套件是不可配置的。
	//
	// 如果 CipherSuites 为 nil，则使用安全的默认列表。默认的密码套件可能会随着时间的推移而更改。
	// 在 Go 1.22 中，基于 RSA 密钥交换的密码套件已从默认列表中移除，但可以通过设置 GODEBUG 为 tlsrsakex=1 来重新添加。
	// 在 Go 1.23 中，3DES 密码套件已从默认列表中移除，但可以通过设置 GODEBUG 为 tls3des=1 来重新添加。
	CipherSuites []uint16

	// PreferServerCipherSuites is a legacy field and has no effect.
	//
	// It used to control whether the server would follow the client's or the
	// server's preference. Servers now select the best mutually supported
	// cipher suite based on logic that takes into account inferred client
	// hardware, server hardware, and security.
	//
	// Deprecated: PreferServerCipherSuites is ignored.
    // PreferServerCipherSuites 是一个遗留字段，没有效果。
	//
	// 它曾经控制服务器是否遵循客户端或服务器的首选项。
	// 现在，服务器将根据考虑到客户端硬件、服务器硬件和安全性的逻辑选择最兼容的密码套件。
	//
	// 已弃用：PreferServerCipherSuites 被忽略。
	PreferServerCipherSuites bool

	// SessionTicketsDisabled may be set to true to disable session ticket and
	// PSK (resumption) support. Note that on clients, session ticket support is
	// also disabled if ClientSessionCache is nil.
    // SessionTicketsDisabled 可以设置为 true 以禁用会话票据和 PSK（恢复）支持。
	// 请注意，在客户端，如果 ClientSessionCache 为 nil，则会话票据支持也将被禁用。
	SessionTicketsDisabled bool

	// SessionTicketKey is used by TLS servers to provide session resumption.
	// See RFC 5077 and the PSK mode of RFC 8446. If zero, it will be filled
	// with random data before the first server handshake.
	//
	// Deprecated: if this field is left at zero, session ticket keys will be
	// automatically rotated every day and dropped after seven days. For
	// customizing the rotation schedule or synchronizing servers that are
	// terminating connections for the same host, use SetSessionTicketKeys.
    // SessionTicketKey 由 TLS 服务器用于提供会话恢复。请参见 RFC 5077 和 RFC 8446 的 PSK 模式。
	// 如果为零，在首次服务器握手之前会用随机数据填充它。
	//
	// 已弃用：如果此字段保持为零，则会话票据密钥将每天自动旋转一次，并在七天后丢弃。
	// 要自定义旋转计划或同步终止同一主机连接的服务器，请使用 SetSessionTicketKeys。
	SessionTicketKey [32]byte

	// ClientSessionCache is a cache of ClientSessionState entries for TLS
	// session resumption. It is only used by clients.
    // ClientSessionCache 是 TLS 会话恢复的 ClientSessionState 条目的缓存。它仅由客户端使用。
	ClientSessionCache ClientSessionCache

	// UnwrapSession is called on the server to turn a ticket/identity
	// previously produced by [WrapSession] into a usable session.
	//
	// UnwrapSession will usually either decrypt a session state in the ticket
	// (for example with [Config.EncryptTicket]), or use the ticket as a handle
	// to recover a previously stored state. It must use [ParseSessionState] to
	// deserialize the session state.
	//
	// If UnwrapSession returns an error, the connection is terminated. If it
	// returns (nil, nil), the session is ignored. crypto/tls may still choose
	// not to resume the returned session.
    // UnwrapSession 在服务器上调用，以将之前由 [WrapSession] 生成的票据/身份转换为可用会话。
	//
	// UnwrapSession 通常会解密票据中的会话状态（例如使用 [Config.EncryptTicket]），
	// 或者使用票据作为句柄来恢复之前存储的状态。它必须使用 [ParseSessionState] 来反序列化会话状态。
	//
	// 如果 UnwrapSession 返回错误，连接将终止。如果它返回 (nil, nil)，会话将被忽略。
	// crypto/tls 可能仍会选择不恢复返回的会话。
	UnwrapSession func(identity []byte, cs ConnectionState) (*SessionState, error)

	// WrapSession is called on the server to produce a session ticket/identity.
	//
	// WrapSession must serialize the session state with [SessionState.Bytes].
	// It may then encrypt the serialized state (for example with
	// [Config.DecryptTicket]) and use it as the ticket, or store the state and
	// return a handle for it.
	//
	// If WrapSession returns an error, the connection is terminated.
	//
	// Warning: the return value will be exposed on the wire and to clients in
	// plaintext. The application is in charge of encrypting and authenticating
	// it (and rotating keys) or returning high-entropy identifiers. Failing to
	// do so correctly can compromise current, previous, and future connections
	// depending on the protocol version.
    // WrapSession 在服务器上调用以生成会话票据或 PSK 的身份。
	// 当 crypto/tls 生成会话票据时，它将使用此功能包装会话状态（例如使用 [Config.EncryptTicket] 加密）。
	//
	// 返回的切片将用于 [ClientSessionState.Identity]，并将传递给 [Config.UnwrapSession]。
	// 如果 UnwrapSession 返回 nil，恢复将失败，并且将像通常一样执行完整握手。
	// crypto/tls 可能仍会选择不发送票据（例如，因为它在 TLS 1.3 中不支持恢复客户端证书）。
	//
	// 如果 WrapSession 返回错误，连接将终止。
	WrapSession func(ConnectionState, *SessionState) ([]byte, error)

	// MinVersion contains the minimum TLS version that is acceptable.
	//
	// By default, TLS 1.2 is currently used as the minimum. TLS 1.0 is the
	// minimum supported by this package.
	//
	// The server-side default can be reverted to TLS 1.0 by including the value
	// "tls10server=1" in the GODEBUG environment variable.
    // MinVersion 包含可接受的最低 TLS 版本。
	//
	// 默认情况下，TLS 1.2 当前作为最低版本使用。此包支持的最低版本是 TLS 1.0。
	//
	// 通过在 GODEBUG 环境变量中包含值 "tls10server=1"，可以将服务器端默认值恢复为 TLS 1.0。
	MinVersion uint16

	// MaxVersion contains the maximum TLS version that is acceptable.
	//
	// By default, the maximum version supported by this package is used,
	// which is currently TLS 1.3.
    // MaxVersion 包含可接受的最高 TLS 版本。
	//
	// 默认情况下，使用该包支持的最高版本，当前为 TLS 1.3。
	MaxVersion uint16

	// CurvePreferences contains the elliptic curves that will be used in
	// an ECDHE handshake, in preference order. If empty, the default will
	// be used. The client will use the first preference as the type for
	// its key share in TLS 1.3. This may change in the future.
	//
	// From Go 1.23, the default includes the X25519Kyber768Draft00 hybrid
	// post-quantum key exchange. To disable it, set CurvePreferences explicitly
	// or use the GODEBUG=tlskyber=0 environment variable.
    // CurvePreferences 包含将在 ECDHE 握手中使用的椭圆曲线，按优先顺序排列。
	// 如果为空，将使用默认值。客户端将使用第一个首选项作为 TLS 1.3 中其密钥共享的类型。
	// 未来可能会有所变化。
	//
	// 从 Go 1.23 开始，默认包括 X25519Kyber768Draft00 混合后量子密钥交换。
	// 要禁用它，请显式设置 CurvePreferences 或使用 GODEBUG=tlskyber=0 环境变量。
    // From Go 1.24, the default includes the [X25519MLKEM768] hybrid
	// post-quantum key exchange. To disable it, set CurvePreferences explicitly
	// or use the GODEBUG=tlsmlkem=0 environment variable.
	CurvePreferences []CurveID

	// DynamicRecordSizingDisabled disables adaptive sizing of TLS records.
	// When true, the largest possible TLS record size is always used. When
	// false, the size of TLS records may be adjusted in an attempt to
	// improve latency.
    // DynamicRecordSizingDisabled 禁用 TLS 记录的自适应大小调整。
	// 当为 true 时，将始终使用最大的 TLS 记录大小。为 false 时，
	// 记录大小可能会调整以尝试改善延迟。
	DynamicRecordSizingDisabled bool

	// Renegotiation controls what types of renegotiation are supported.
	// The default, none, is correct for the vast majority of applications.
    // Renegotiation 控制支持的重新协商类型。
	// 默认设置为不允许重新协商，这是绝大多数应用程序的正确选择。
	Renegotiation RenegotiationSupport

	// KeyLogWriter optionally specifies a destination for TLS master secrets
	// in NSS key log format that can be used to allow external programs
	// such as Wireshark to decrypt TLS connections.
	// See https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Format.
	// Use of KeyLogWriter compromises security and should only be
	// used for debugging.
    // KeyLogWriter 可选地指定用于记录 TLS 主密钥的目的地，
	// 以 NSS 密钥日志格式，这些密钥日志可用于让外部程序如 Wireshark 解密 TLS 连接。
	// 参见 https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Format。
	// 使用 KeyLogWriter 会降低安全性，应仅用于调试。
	KeyLogWriter io.Writer

	// EncryptedClientHelloConfigList is a serialized ECHConfigList. If
	// provided, clients will attempt to connect to servers using Encrypted
	// Client Hello (ECH) using one of the provided ECHConfigs. Servers
	// currently ignore this field.
	//
	// If the list contains no valid ECH configs, the handshake will fail
	// and return an error.
	//
	// If EncryptedClientHelloConfigList is set, MinVersion, if set, must
	// be VersionTLS13.
	//
	// When EncryptedClientHelloConfigList is set, the handshake will only
	// succeed if ECH is sucessfully negotiated. If the server rejects ECH,
	// an ECHRejectionError error will be returned, which may contain a new
	// ECHConfigList that the server suggests using.
	//
	// How this field is parsed may change in future Go versions, if the
	// encoding described in the final Encrypted Client Hello RFC changes.
    // EncryptedClientHelloConfigList 是序列化的 ECHConfigList。
	// 如果提供，客户端将尝试使用提供的 ECHConfigs 中的一个通过加密客户端 Hello (ECH) 连接到服务器。
	// 服务器当前忽略该字段。
	//
	// 如果列表中没有有效的 ECH 配置，握手将失败并返回错误。
	//
	// 如果设置了 EncryptedClientHelloConfigList，则设置的 MinVersion 必须为 VersionTLS13。
	//
	// 当设置了 EncryptedClientHelloConfigList 时，握手只有在成功协商 ECH 的情况下才会成功。
	// 如果服务器拒绝 ECH，将返回 ECHRejectionError 错误，其中可能包含服务器建议使用的新 ECHConfigList。
	//
	// 如果最终的 Encrypted Client Hello RFC 中描述的编码发生变化，
	// 那么该字段的解析方式可能会在未来的 Go 版本中发生变化。
	EncryptedClientHelloConfigList []byte

	// EncryptedClientHelloRejectionVerify, if not nil, is called when ECH is
	// rejected, in order to verify the ECH provider certificate in the outer
	// Client Hello. If it returns a non-nil error, the handshake is aborted and
	// that error results.
	//
	// Unlike VerifyPeerCertificate and VerifyConnection, normal certificate
	// verification will not be performed before calling
	// EncryptedClientHelloRejectionVerify.
	//
	// If EncryptedClientHelloRejectionVerify is nil and ECH is rejected, the
	// roots in RootCAs will be used to verify the ECH providers public
	// certificate. VerifyPeerCertificate and VerifyConnection are not called
	// when ECH is rejected, even if set, and InsecureSkipVerify is ignored.
    // EncryptedClientHelloRejectionVerify，如果不为 nil，则在 ECH 被拒绝时调用，
	// 以验证外部客户端 Hello 中的 ECH 提供者证书。
	// 如果返回非 nil 错误，握手将中止并导致该错误。
	//
	// 与 VerifyPeerCertificate 和 VerifyConnection 不同，在调用
	// EncryptedClientHelloRejectionVerify 之前不会执行正常的证书验证。
	//
	// 如果 EncryptedClientHelloRejectionVerify 为 nil 且 ECH 被拒绝，
	// 则将使用 RootCAs 中的根来验证 ECH 提供者的公钥证书。
	// 即使设置了 VerifyPeerCertificate 和 VerifyConnection，当 ECH 被拒绝时也不会调用，
	// 并且 InsecureSkipVerify 将被忽略。
	EncryptedClientHelloRejectionVerify func(ConnectionState) error
	// contains filtered or unexported fields
}
```

A Config structure is used to configure a TLS client or server. After one has been passed to a TLS function it must not be modified. A Config may be reused; the tls package will also not modify it.

​	Config 结构用于配置 TLS 客户端或服务器。在将其传递给 TLS 函数后，不得对其进行修改。Config 可以重复使用；tls 包也不会修改它。

#### Example (KeyLogWriter)

```go
package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

// zeroSource is an io.Reader that returns an unlimited number of zero bytes.
type zeroSource struct{}

func (zeroSource) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 0
	}

	return len(b), nil
}

func main() {
	// Debugging TLS applications by decrypting a network traffic capture.

	// WARNING: Use of KeyLogWriter compromises security and should only be
	// used for debugging.

	// Dummy test HTTP server for the example with insecure random so output is
	// reproducible.
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	server.TLS = &tls.Config{
		Rand: zeroSource{}, // for example only; don't do this.
	}
	server.StartTLS()
	defer server.Close()

	// Typically the log would go to an open file:
	// w, err := os.OpenFile("tls-secrets.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	w := os.Stdout

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				KeyLogWriter: w,

				Rand:               zeroSource{}, // for reproducible output; don't do this.
				InsecureSkipVerify: true,         // test server certificate is not trusted.
			},
		},
	}
	resp, err := client.Get(server.URL)
	if err != nil {
		log.Fatalf("Failed to get URL: %v", err)
	}
	resp.Body.Close()

	// The resulting file can be used with Wireshark to decrypt the TLS
	// connection by setting (Pre)-Master-Secret log filename in SSL Protocol
	// preferences.
}

```



#### Example (VerifyConnection)

```go
package main

import (
	"crypto/tls"
	"crypto/x509"
)

func main() {
	// VerifyConnection can be used to replace and customize connection
	// verification. This example shows a VerifyConnection implementation that
	// will be approximately equivalent to what crypto/tls does normally to
	// verify the peer's certificate.

	// Client side configuration.
	_ = &tls.Config{
		// Set InsecureSkipVerify to skip the default validation we are
		// replacing. This will not disable VerifyConnection.
		InsecureSkipVerify: true,
		VerifyConnection: func(cs tls.ConnectionState) error {
			opts := x509.VerifyOptions{
				DNSName:       cs.ServerName,
				Intermediates: x509.NewCertPool(),
			}
			for _, cert := range cs.PeerCertificates[1:] {
				opts.Intermediates.AddCert(cert)
			}
			_, err := cs.PeerCertificates[0].Verify(opts)
			return err
		},
	}

	// Server side configuration.
	_ = &tls.Config{
		// Require client certificates (or VerifyConnection will run anyway and
		// panic accessing cs.PeerCertificates[0]) but don't verify them with the
		// default verifier. This will not disable VerifyConnection.
		ClientAuth: tls.RequireAnyClientCert,
		VerifyConnection: func(cs tls.ConnectionState) error {
			opts := x509.VerifyOptions{
				DNSName:       cs.ServerName,
				Intermediates: x509.NewCertPool(),
				KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			}
			for _, cert := range cs.PeerCertificates[1:] {
				opts.Intermediates.AddCert(cert)
			}
			_, err := cs.PeerCertificates[0].Verify(opts)
			return err
		},
	}

	// Note that when certificates are not handled by the default verifier
	// ConnectionState.VerifiedChains will be nil.
}
Output:
```



####  (*Config) BuildNameToCertificate <- DEPRECATED

```go
func (c *Config) BuildNameToCertificate()
```

BuildNameToCertificate parses c.Certificates and builds c.NameToCertificate from the CommonName and SubjectAlternateName fields of each of the leaf certificates.

​	BuildNameToCertificate 解析 c.Certificates 并从每个叶证书的 CommonName 和 SubjectAlternateName 字段构建 c.NameToCertificate。

Deprecated: NameToCertificate only allows associating a single certificate with a given name. Leave that field nil to let the library select the first compatible chain from Certificates.
	已弃用：NameToCertificate 仅允许将单个证书与给定名称相关联。将该字段保留为 nil，以便库从 Certificates 中选择第一个兼容的链。

#### (*Config) Clone  <- go1.8

``` go
func (c *Config) Clone() *Config
```

Clone returns a shallow clone of c or nil if c is nil. It is safe to clone a Config that is being used concurrently by a TLS client or server.

​	如果 c 为 nil，则 Clone 返回 c 的浅层克隆或 nil。可以克隆正在被 TLS 客户端或服务器并发使用的 Config。

#### (*Config) DecryptTicket <-go1.21.0

``` go
func (c *Config) DecryptTicket(identity []byte, cs ConnectionState) (*SessionState, error)
```

DecryptTicket decrypts a ticket encrypted by [Config.EncryptTicket](https://pkg.go.dev/crypto/tls@go1.23.0#Config.EncryptTicket). It can be used as a [Config.UnwrapSession] implementation.

​	DecryptTicket 用于解密由 [Config.EncryptTicket](https://pkg.go.dev/crypto/tls@go1.23.0#Config.EncryptTicket) 加密的票据。它可以用作 [Config.UnwrapSession] 的实现。

If the ticket can't be decrypted or parsed, DecryptTicket returns (nil, nil).

​	如果票据无法解密或解析，DecryptTicket 返回 (nil, nil)。

#### (*Config) EncryptTicket <-go1.21.0

``` go
func (c *Config) EncryptTicket(cs ConnectionState, ss *SessionState) ([]byte, error)
```

EncryptTicket encrypts a ticket with the [Config](https://pkg.go.dev/crypto/tls@go1.23.0#Config)'s configured (or default) session ticket keys. It can be used as a [Config.WrapSession] implementation.

​	EncryptTicket 使用 [Config](https://pkg.go.dev/crypto/tls@go1.23.0#Config) 配置的（或默认的）会话票据密钥加密票据。它可以用作 [Config.WrapSession] 的实现。

#### (*Config) SetSessionTicketKeys  <- go1.5

``` go
func (c *Config) SetSessionTicketKeys(keys [][32]byte)
```

SetSessionTicketKeys updates the session ticket keys for a server.

​	SetSessionTicketKeys 更新服务器的会话票证密钥。

The first key will be used when creating new tickets, while all keys can be used for decrypting tickets. It is safe to call this function while the server is running in order to rotate the session ticket keys. The function will panic if keys is empty.

​	创建新票证时将使用第一个密钥，而所有密钥均可用于解密票证。在服务器运行时调用此函数以轮换会话票证密钥是安全的。如果 keys 为空，该函数将引发 panic。

Calling this function will turn off automatic session ticket key rotation.

​	调用此函数将关闭自动会话票证密钥轮换。

If multiple servers are terminating connections for the same host they should all have the same session ticket keys. If the session ticket keys leaks, previously recorded and future TLS connections using those keys might be compromised.

​	如果多个服务器正在终止同一主机的连接，则它们都应具有相同的会话票证密钥。如果会话票证密钥泄露，则使用这些密钥记录的先前和未来的 TLS 连接可能会受到损害。

### type Conn 

``` go
type Conn struct {
	// contains filtered or unexported fields
}
```

A Conn represents a secured connection. It implements the net.Conn interface.

​	Conn 表示一个安全连接。它实现了 net.Conn 接口。

#### func Client 

``` go
func Client(conn net.Conn, config *Config) *Conn
```

Client returns a new TLS client side connection using conn as the underlying transport. The config cannot be nil: users must set either ServerName or InsecureSkipVerify in the config.

​	Client 使用 conn 作为底层传输返回一个新的 TLS 客户端连接。config 不能为 nil：用户必须在 config 中设置 ServerName 或 InsecureSkipVerify。

#### func Dial 

``` go
func Dial(network, addr string, config *Config) (*Conn, error)
```

Dial connects to the given network address using net.Dial and then initiates a TLS handshake, returning the resulting TLS connection. Dial interprets a nil configuration as equivalent to the zero configuration; see the documentation of Config for the defaults.

​	Dial 使用 net.Dial 连接到给定的网络地址，然后启动 TLS 握手，返回生成的 TLS 连接。Dial 将 nil 配置解释为等同于零配置；有关默认值，请参阅 Config 的文档。

##### Dial Example

```go
package main

import (
	"crypto/tls"
	"crypto/x509"
)

func main() {
	// Connecting with a custom root-certificate set.

	const rootPEM = `
-- GlobalSign Root R2, valid until Dec 15, 2021
-----BEGIN CERTIFICATE-----
MIIDujCCAqKgAwIBAgILBAAAAAABD4Ym5g0wDQYJKoZIhvcNAQEFBQAwTDEgMB4G
A1UECxMXR2xvYmFsU2lnbiBSb290IENBIC0gUjIxEzARBgNVBAoTCkdsb2JhbFNp
Z24xEzARBgNVBAMTCkdsb2JhbFNpZ24wHhcNMDYxMjE1MDgwMDAwWhcNMjExMjE1
MDgwMDAwWjBMMSAwHgYDVQQLExdHbG9iYWxTaWduIFJvb3QgQ0EgLSBSMjETMBEG
A1UEChMKR2xvYmFsU2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbjCCASIwDQYJKoZI
hvcNAQEBBQADggEPADCCAQoCggEBAKbPJA6+Lm8omUVCxKs+IVSbC9N/hHD6ErPL
v4dfxn+G07IwXNb9rfF73OX4YJYJkhD10FPe+3t+c4isUoh7SqbKSaZeqKeMWhG8
eoLrvozps6yWJQeXSpkqBy+0Hne/ig+1AnwblrjFuTosvNYSuetZfeLQBoZfXklq
tTleiDTsvHgMCJiEbKjNS7SgfQx5TfC4LcshytVsW33hoCmEofnTlEnLJGKRILzd
C9XZzPnqJworc5HGnRusyMvo4KD0L5CLTfuwNhv2GXqF4G3yYROIXJ/gkwpRl4pa
zq+r1feqCapgvdzZX99yqWATXgAByUr6P6TqBwMhAo6CygPCm48CAwEAAaOBnDCB
mTAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUm+IH
V2ccHsBqBt5ZtJot39wZhi4wNgYDVR0fBC8wLTAroCmgJ4YlaHR0cDovL2NybC5n
bG9iYWxzaWduLm5ldC9yb290LXIyLmNybDAfBgNVHSMEGDAWgBSb4gdXZxwewGoG
3lm0mi3f3BmGLjANBgkqhkiG9w0BAQUFAAOCAQEAmYFThxxol4aR7OBKuEQLq4Gs
J0/WwbgcQ3izDJr86iw8bmEbTUsp9Z8FHSbBuOmDAGJFtqkIk7mpM0sYmsL4h4hO
291xNBrBVNpGP+DTKqttVCL1OmLNIG+6KYnX3ZHu01yiPqFbQfXf5WRDLenVOavS
ot+3i9DAgBkcRcAtjOj4LaR0VknFBbVPFd5uRHg5h6h+u/N5GJG79G+dwfCMNYxd
AfvDbbnvRG15RjF+Cv6pgsH/76tuIMRQyV+dTZsXjAzlAcmgQWpzU/qlULRuJQ/7
TBj0/VLZjmmx6BEP3ojY+x1J96relc8geMJgEtslQIxq/H5COEBkEveegeGTLg==
-----END CERTIFICATE-----`

	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	conn, err := tls.Dial("tcp", "mail.google.com:443", &tls.Config{
		RootCAs: roots,
	})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	conn.Close()
}
Output:
```



#### func DialWithDialer  <- go1.3

``` go
func DialWithDialer(dialer *net.Dialer, network, addr string, config *Config) (*Conn, error)
```

DialWithDialer connects to the given network address using dialer.Dial and then initiates a TLS handshake, returning the resulting TLS connection. Any timeout or deadline given in the dialer apply to connection and TLS handshake as a whole.

​	DialWithDialer 使用 dialer.Dial 连接到给定的网络地址，然后启动 TLS 握手，返回生成的 TLS 连接。dialer 中给出的任何超时或截止时间都适用于连接和 TLS 握手。

DialWithDialer interprets a nil configuration as equivalent to the zero configuration; see the documentation of Config for the defaults.

​	DialWithDialer 将 nil 配置解释为等同于零配置；有关默认值，请参阅 Config 的文档。

DialWithDialer uses context.Background internally; to specify the context, use Dialer.DialContext with NetDialer set to the desired dialer.

​	DialWithDialer 在内部使用 context.Background；要指定上下文，请将 Dialer.DialContext 与 NetDialer 设置为所需的拨号器。

#### func Server 

``` go
func Server(conn net.Conn, config *Config) *Conn
```

Server returns a new TLS server side connection using conn as the underlying transport. The configuration config must be non-nil and must include at least one certificate or else set GetCertificate.

​	Server 使用 conn 作为底层传输返回新的 TLS 服务器端连接。配置 config 必须为非 nil，并且必须至少包含一个证书，否则设置 GetCertificate。

#### (*Conn) Close 

``` go
func (c *Conn) Close() error
```

Close closes the connection.

​	Close 关闭连接。

#### (*Conn) CloseWrite  <- go1.8

``` go
func (c *Conn) CloseWrite() error
```

CloseWrite shuts down the writing side of the connection. It should only be called once the handshake has completed and does not call CloseWrite on the underlying connection. Most callers should just use Close.

​	CloseWrite 关闭连接的写入端。它只应在握手完成后调用，并且不会对底层连接调用 CloseWrite。大多数调用者应仅使用 Close。

#### (*Conn) ConnectionState 

``` go
func (c *Conn) ConnectionState() ConnectionState
```

ConnectionState returns basic TLS details about the connection.

​	ConnectionState 返回有关连接的基本 TLS 详细信息。

#### (*Conn) Handshake 

``` go
func (c *Conn) Handshake() error
```

Handshake runs the client or server handshake protocol if it has not yet been run.

​	如果尚未运行客户端或服务器握手协议，Handshake 将运行该协议。

Most uses of this package need not call Handshake explicitly: the first Read or Write will call it automatically.

​	大多数使用此软件包的情况不需要显式调用 Handshake：第一个 Read 或 Write 将自动调用它。

For control over canceling or setting a timeout on a handshake, use HandshakeContext or the Dialer’s DialContext method instead.

​	要控制取消或设置握手超时，请改用 HandshakeContext 或 Dialer 的 DialContext 方法。

#### (*Conn) HandshakeContext  <- go1.17

``` go
func (c *Conn) HandshakeContext(ctx context.Context) error
```

HandshakeContext runs the client or server handshake protocol if it has not yet been run.

​	如果尚未运行客户端或服务器握手协议，HandshakeContext 将运行该协议。

The provided Context must be non-nil. If the context is canceled before the handshake is complete, the handshake is interrupted and an error is returned. Once the handshake has completed, cancellation of the context will not affect the connection.

​	提供的 Context 必须是非 nil。如果在握手完成之前取消了上下文，则会中断握手并返回错误。握手完成后，取消上下文不会影响连接。

Most uses of this package need not call HandshakeContext explicitly: the first Read or Write will call it automatically.

​	大多数使用此软件包的情况不需要显式调用 HandshakeContext：第一次 Read 或 Write 会自动调用它。

#### (*Conn) LocalAddr 

``` go
func (c *Conn) LocalAddr() net.Addr
```

LocalAddr returns the local network address.

​	LocalAddr 返回本地网络地址。

#### (*Conn) NetConn  <- go1.18

``` go
func (c *Conn) NetConn() net.Conn
```

NetConn returns the underlying connection that is wrapped by c. Note that writing to or reading from this connection directly will corrupt the TLS session.

​	NetConn 返回 c 包装的基础连接。请注意，直接向此连接写入或从中读取会损坏 TLS 会话。

#### (*Conn) OCSPResponse 

``` go
func (c *Conn) OCSPResponse() []byte
```

OCSPResponse returns the stapled OCSP response from the TLS server, if any. (Only valid for client connections.)

​	OCSPResponse 返回来自 TLS 服务器的装订 OCSP 响应（如果有）。（仅对客户端连接有效。）

#### (*Conn) Read 

``` go
func (c *Conn) Read(b []byte) (int, error)
```

Read reads data from the connection.

​	Read 从连接中读取数据。

As Read calls Handshake, in order to prevent indefinite blocking a deadline must be set for both Read and Write before Read is called when the handshake has not yet completed. See SetDeadline, SetReadDeadline, and SetWriteDeadline.

​	由于 Read 调用 Handshake，为了防止无限期阻塞，必须在 Read 调用时为 Read 和 Write 都设置一个截止时间，而此时握手尚未完成。请参阅 SetDeadline、SetReadDeadline 和 SetWriteDeadline。

#### (*Conn) RemoteAddr 

``` go
func (c *Conn) RemoteAddr() net.Addr
```

RemoteAddr returns the remote network address.

RemoteAddr 返回远程网络地址。

#### (*Conn) SetDeadline

```go
func (c *Conn) SetDeadline(t time.Time) error
```

SetDeadline sets the read and write deadlines associated with the connection. A zero value for t means Read and Write will not time out. After a Write has timed out, the TLS state is corrupt and all future writes will return the same error.

​	SetDeadline 设置与连接关联的读写截止时间。t 的值为零表示 Read 和 Write 不会超时。在 Write 超时后，TLS 状态损坏，所有后续写入都将返回相同的错误。

#### (*Conn) SetReadDeadline

```go
func (c *Conn) SetReadDeadline(t time.Time) error
```

SetReadDeadline sets the read deadline on the underlying connection. A zero value for t means Read will not time out.

​	SetReadDeadline 设置底层连接的读取截止时间。t 的值为零表示 Read 不会超时。

#### (*Conn) SetWriteDeadline

```go
func (c *Conn) SetWriteDeadline(t time.Time) error
```

SetWriteDeadline sets the write deadline on the underlying connection. A zero value for t means Write will not time out. After a Write has timed out, the TLS state is corrupt and all future writes will return the same error.

​	SetWriteDeadline 设置底层连接的写入截止时间。t 的值为零表示 Write 不会超时。在 Write 超时后，TLS 状态损坏，所有后续写入都将返回相同的错误。

#### (*Conn) VerifyHostname

```go
func (c *Conn) VerifyHostname(host string) error
```

VerifyHostname checks that the peer certificate chain is valid for connecting to host. If so, it returns nil; if not, it returns an error describing the problem.

​	VerifyHostname 检查对等证书链是否有效，以便连接到主机。如果是，则返回 nil；如果不是，则返回描述问题的错误。

#### (*Conn) Write

```go
func (c *Conn) Write(b []byte) (int, error)
```

Write writes data to the connection.

​	Write 将数据写入连接。

As Write calls Handshake, in order to prevent indefinite blocking a deadline must be set for both Read and Write before Write is called when the handshake has not yet completed. See SetDeadline, SetReadDeadline, and SetWriteDeadline.

​	由于 Write 调用 Handshake，因此为了防止无限期阻塞，必须在调用 Write 之前为 Read 和 Write 设置一个截止时间，而此时握手尚未完成。请参阅 SetDeadline、SetReadDeadline 和 SetWriteDeadline。

### type ConnectionState 

``` go
type ConnectionState struct {
	// Version is the TLS version used by the connection (e.g. VersionTLS12).
	Version uint16

	// HandshakeComplete is true if the handshake has concluded.
	HandshakeComplete bool

	// DidResume is true if this connection was successfully resumed from a
	// previous session with a session ticket or similar mechanism.
	DidResume bool

	// CipherSuite is the cipher suite negotiated for the connection (e.g.
	// TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, TLS_AES_128_GCM_SHA256).
	CipherSuite uint16

	// NegotiatedProtocol is the application protocol negotiated with ALPN.
	NegotiatedProtocol string

	// NegotiatedProtocolIsMutual used to indicate a mutual NPN negotiation.
	//
	// Deprecated: this value is always true.
	NegotiatedProtocolIsMutual bool

	// ServerName is the value of the Server Name Indication extension sent by
	// the client. It's available both on the server and on the client side.
	ServerName string

	// PeerCertificates are the parsed certificates sent by the peer, in the
	// order in which they were sent. The first element is the leaf certificate
	// that the connection is verified against.
	//
	// On the client side, it can't be empty. On the server side, it can be
	// empty if Config.ClientAuth is not RequireAnyClientCert or
	// RequireAndVerifyClientCert.
	//
	// PeerCertificates and its contents should not be modified.
	PeerCertificates []*x509.Certificate

	// VerifiedChains is a list of one or more chains where the first element is
	// PeerCertificates[0] and the last element is from Config.RootCAs (on the
	// client side) or Config.ClientCAs (on the server side).
	//
	// On the client side, it's set if Config.InsecureSkipVerify is false. On
	// the server side, it's set if Config.ClientAuth is VerifyClientCertIfGiven
	// (and the peer provided a certificate) or RequireAndVerifyClientCert.
	//
	// VerifiedChains and its contents should not be modified.
	VerifiedChains [][]*x509.Certificate

	// SignedCertificateTimestamps is a list of SCTs provided by the peer
	// through the TLS handshake for the leaf certificate, if any.
	SignedCertificateTimestamps [][]byte

	// OCSPResponse is a stapled Online Certificate Status Protocol (OCSP)
	// response provided by the peer for the leaf certificate, if any.
	OCSPResponse []byte

	// TLSUnique contains the "tls-unique" channel binding value (see RFC 5929,
	// Section 3). This value will be nil for TLS 1.3 connections and for all
	// resumed connections.
	//
	// Deprecated: there are conditions in which this value might not be unique
	// to a connection. See the Security Considerations sections of RFC 5705 and
	// RFC 7627, and https://mitls.org/pages/attacks/3SHAKE#channelbindings.
	TLSUnique []byte
	// contains filtered or unexported fields
}
```

ConnectionState records basic TLS details about the connection.

​	ConnectionState 记录有关连接的基本 TLS 详细信息。

#### (*ConnectionState) ExportKeyingMaterial  <- go1.11

``` go
func (cs *ConnectionState) ExportKeyingMaterial(label string, context []byte, length int) ([]byte, error)
```

ExportKeyingMaterial returns length bytes of exported key material in a new slice as defined in [RFC 5705](https://rfc-editor.org/rfc/rfc5705.html). If context is nil, it is not used as part of the seed. If the connection was set to allow renegotiation via Config.Renegotiation, this function will return an error.

​	ExportKeyingMaterial 以新切片形式返回长度为字节的导出密钥材料，如 RFC 5705 中所定义。如果 context 为 nil，则不会将其用作种子的一部分。如果通过 Config.Renegotiation 将连接设置为允许重新协商，则此函数将返回错误。

Exporting key material without Extended Master Secret or TLS 1.3 was disabled in Go 1.22 due to security issues (see the Security Considerations sections of [RFC 5705](https://rfc-editor.org/rfc/rfc5705.html) and [RFC 7627](https://rfc-editor.org/rfc/rfc7627.html)), but can be re-enabled with the GODEBUG setting tlsunsafeekm=1.

​	由于安全问题，在Go 1.22中禁用了没有Extended Master Secret或TLS 1.3的导出密钥材料(参见[RFC 5705](https://rfc-editor.org/rfc/rfc5705.html)和[RFC 7627](https://rfc-editor.org/rfc/rfc7627.html)的安全注意事项部分)，但可以通过GODEBUG设置tlsunsafeekm=1重新启用。

### type CurveID  <- go1.3

``` go
type CurveID uint16
```

CurveID is the type of a TLS identifier for an elliptic curve. See https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8.

​	CurveID 是椭圆曲线的 TLS 标识符的类型。请参阅 https://www.iana.org/assignments/tls-parameters/tls-parameters.xml#tls-parameters-8。

In TLS 1.3, this type is called NamedGroup, but at this time this library only supports Elliptic Curve based groups. See [RFC 8446, Section 4.2.7](https://rfc-editor.org/rfc/rfc8446.html#section-4.2.7).

​	在 TLS 1.3 中，此类型称为 NamedGroup，但此时此库仅支持基于椭圆曲线的组。请参阅 RFC 8446，第 4.2.7 节。

``` go
const (
	CurveP256 CurveID = 23
	CurveP384 CurveID = 24
	CurveP521 CurveID = 25
	X25519    CurveID = 29
)
```

#### (CurveID) String  <- go1.15

``` go
func (i CurveID) String() string
```

### type Dialer  <- go1.15

``` go
type Dialer struct {
	// NetDialer is the optional dialer to use for the TLS connections'
	// underlying TCP connections.
	// A nil NetDialer is equivalent to the net.Dialer zero value.
	NetDialer *net.Dialer

	// Config is the TLS configuration to use for new connections.
	// A nil configuration is equivalent to the zero
	// configuration; see the documentation of Config for the
	// defaults.
	Config *Config
}
```

Dialer dials TLS connections given a configuration and a Dialer for the underlying connection.

​	Dialer 为给定的配置和底层连接的 Dialer 拨打 TLS 连接。

#### (*Dialer) Dial  <- go1.15

``` go
func (d *Dialer) Dial(network, addr string) (net.Conn, error)
```

Dial connects to the given network address and initiates a TLS handshake, returning the resulting TLS connection.

​	Dial 连接到给定的网络地址并启动 TLS 握手，返回生成的 TLS 连接。

The returned Conn, if any, will always be of type *Conn.

​	返回的 Conn（如果有）始终为 *Conn 类型。

Dial uses context.Background internally; to specify the context, use DialContext.

​	Dial 在内部使用 context.Background；要指定上下文，请使用 DialContext。

#### (*Dialer) DialContext  <- go1.15

``` go
func (d *Dialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error)
```

DialContext connects to the given network address and initiates a TLS handshake, returning the resulting TLS connection.

​	DialContext 连接到给定的网络地址并启动 TLS 握手，返回生成的 TLS 连接。

The provided Context must be non-nil. If the context expires before the connection is complete, an error is returned. Once successfully connected, any expiration of the context will not affect the connection.

​	提供的 Context 必须是非 nil。如果在连接完成之前上下文过期，则会返回一个错误。一旦成功连接，上下文的任何过期都不会影响连接。

The returned Conn, if any, will always be of type *Conn.

​	返回的 Conn（如果有）将始终为 *Conn 类型。

### type ECHRejectionError <- go1.23.0

``` go
type ECHRejectionError struct {
	RetryConfigList []byte
}
```

ECHRejectionError is the error type returned when ECH is rejected by a remote server. If the server offered a ECHConfigList to use for retries, the RetryConfigList field will contain this list.

​	当远程服务器拒绝 ECH 时，将返回 ECHRejectionError 错误类型。如果服务器提供了用于重试的 ECHConfigList，RetryConfigList 字段将包含此列表。

The client may treat an ECHRejectionError with an empty set of RetryConfigs as a secure signal from the server.

​	客户端可以将没有 RetryConfigs 的 ECHRejectionError 视为来自服务器的安全信号。

#### (*ECHRejectionError) Error <- go1.23.0

``` go
func (e *ECHRejectionError) Error() string
```

### type EncryptedClientHelloKey <- 1.24.0

```go
type EncryptedClientHelloKey struct {
	// Config should be a marshalled ECHConfig associated with PrivateKey. This
	// must match the config provided to clients byte-for-byte. The config
	// should only specify the DHKEM(X25519, HKDF-SHA256) KEM ID (0x0020), the
	// HKDF-SHA256 KDF ID (0x0001), and a subset of the following AEAD IDs:
	// AES-128-GCM (0x0000), AES-256-GCM (0x0001), ChaCha20Poly1305 (0x0002).
	Config []byte
	// PrivateKey should be a marshalled private key. Currently, we expect
	// this to be the output of [ecdh.PrivateKey.Bytes].
	PrivateKey []byte
	// SendAsRetry indicates if Config should be sent as part of the list of
	// retry configs when ECH is requested by the client but rejected by the
	// server.
	SendAsRetry bool
}
```

EncryptedClientHelloKey holds a private key that is associated with a specific ECH config known to a client.

### type QUICConfig <-go1.21.0

``` go
type QUICConfig struct {
	TLSConfig *Config

	// EnableSessionEvents may be set to true to enable the
	// [QUICStoreSession] and [QUICResumeSession] events for client connections.
	// When this event is enabled, sessions are not automatically
	// stored in the client session cache.
	// The application should use [QUICConn.StoreSession] to store sessions.
	// EnableSessionEvents 可以设置为 true 以启用客户端连接的 [QUICStoreSession] 和 [QUICResumeSession] 事件。
	// 当此事件启用时，客户端会话缓存中不会自动存储会话。
	// 应用程序应使用 [QUICConn.StoreSession] 来存储会话。
	EnableSessionEvents bool
}
```

A QUICConfig configures a [QUICConn](https://pkg.go.dev/crypto/tls@go1.23.0#QUICConn).

​	QUICConfig 配置了一个 [QUICConn](https://pkg.go.dev/crypto/tls@go1.23.0#QUICConn)。

### type QUICConn <-go1.21.0

``` go
type QUICConn struct {
	// contains filtered or unexported fields
}
```

A QUICConn represents a connection which uses a QUIC implementation as the underlying transport as described in [RFC 9001](https://rfc-editor.org/rfc/rfc9001.html).

​	QUICConn 表示使用 QUIC 实现作为基础传输协议的连接，具体描述请参见 [RFC 9001](https://rfc-editor.org/rfc/rfc9001.html)。

Methods of QUICConn are not safe for concurrent use.

​	QUICConn 的方法不安全，无法并发使用。

#### func QUICClient <-go1.21.0

``` go
func QUICClient(config *QUICConfig) *QUICConn
```

QUICClient returns a new TLS client side connection using QUICTransport as the underlying transport. The config cannot be nil.

​	QUICClient 返回一个使用 QUICTransport 作为底层传输协议的新 TLS 客户端连接。config 不能为空。

The config's MinVersion must be at least TLS 1.3.

​	config 的 MinVersion 必须至少为 TLS 1.3。

#### func QUICServer <-go1.21.0

``` go
func QUICServer(config *QUICConfig) *QUICConn
```

QUICServer returns a new TLS server side connection using QUICTransport as the underlying transport. The config cannot be nil.

​	QUICServer 返回一个使用 QUICTransport 作为底层传输协议的新 TLS 服务器端连接。config 不能为空。

The config's MinVersion must be at least TLS 1.3.

​	config 的 MinVersion 必须至少为 TLS 1.3。

#### (*QUICConn) Close <-go1.21.0

``` go
func (q *QUICConn) Close() error
```

Close closes the connection and stops any in-progress handshake.

​	Close 关闭连接并停止任何正在进行的握手。

#### (*QUICConn) ConnectionState <-go1.21.0

``` go
func (q *QUICConn) ConnectionState() ConnectionState
```

ConnectionState returns basic TLS details about the connection.

​	ConnectionState 返回有关连接的基本 TLS 详细信息。

#### (*QUICConn) HandleData <-go1.21.0

``` go
func (q *QUICConn) HandleData(level QUICEncryptionLevel, data []byte) error
```

HandleData handles handshake bytes received from the peer. It may produce connection events, which may be read with [QUICConn.NextEvent](https://pkg.go.dev/crypto/tls@go1.23.0#QUICConn.NextEvent).

​	HandleData 处理从对等方接收到的握手字节。它可能会生成连接事件，这些事件可以通过 [QUICConn.NextEvent](https://pkg.go.dev/crypto/tls@go1.23.0#QUICConn.NextEvent) 读取。

#### (*QUICConn) NextEvent <-go1.21.0

``` go
func (q *QUICConn) NextEvent() QUICEvent
```

NextEvent returns the next event occurring on the connection. It returns an event with a Kind of [QUICNoEvent](https://pkg.go.dev/crypto/tls@go1.23.0#QUICNoEvent) when no events are available.

​	NextEvent 返回连接上发生的下一个事件。当没有可用事件时，它会返回一个 Kind 为 [QUICNoEvent](https://pkg.go.dev/crypto/tls@go1.23.0#QUICNoEvent) 的事件。

#### (*QUICConn) SendSessionTicket <-go1.21.0

``` go
func (q *QUICConn) SendSessionTicket(opts QUICSessionTicketOptions) error
```

SendSessionTicket sends a session ticket to the client. It produces connection events, which may be read with [QUICConn.NextEvent](https://pkg.go.dev/crypto/tls@go1.23.0#QUICConn.NextEvent). Currently, it can only be called once.

​	SendSessionTicket 向客户端发送会话票据。它会生成连接事件，这些事件可以通过 [QUICConn.NextEvent](https://pkg.go.dev/crypto/tls@go1.23.0#QUICConn.NextEvent) 读取。目前它只能调用一次。

#### (*QUICConn) SetTransportParameters <-go1.21.0

``` go
func (q *QUICConn) SetTransportParameters(params []byte)
```

SetTransportParameters sets the transport parameters to send to the peer.

​	SetTransportParameters 设置要发送给对等方的传输参数。

Server connections may delay setting the transport parameters until after receiving the client's transport parameters. See [QUICTransportParametersRequired](https://pkg.go.dev/crypto/tls@go1.23.0#QUICTransportParametersRequired).

​	服务器连接可能会延迟设置传输参数，直到收到客户端的传输参数。参见 [QUICTransportParametersRequired](https://pkg.go.dev/crypto/tls@go1.23.0#QUICTransportParametersRequired)。

#### (*QUICConn) Start <-go1.21.0

``` go
func (q *QUICConn) Start(ctx context.Context) error
```

Start starts the client or server handshake protocol. It may produce connection events, which may be read with [QUICConn.NextEvent](https://pkg.go.dev/crypto/tls@go1.23.0#QUICConn.NextEvent).

​	Start 启动客户端或服务器握手协议。它可能会生成连接事件，这些事件可以通过 [QUICConn.NextEvent](https://pkg.go.dev/crypto/tls@go1.23.0#QUICConn.NextEvent) 读取。

Start must be called at most once.

​	Start 最多只能调用一次。

#### (*QUICConn) StoreSession <-go1.23.0

``` go
func (q *QUICConn) StoreSession(session *SessionState) error
```

StoreSession stores a session previously received in a QUICStoreSession event in the ClientSessionCache. The application may process additional events or modify the SessionState before storing the session.

​	StoreSession 将在 QUICStoreSession 事件中接收到的会话存储在 ClientSessionCache 中。应用程序可以在存储会话之前处理其他事件或修改 SessionState。

### type QUICEncryptionLevel <-go1.21.0

``` go
type QUICEncryptionLevel int
```

QUICEncryptionLevel represents a QUIC encryption level used to transmit handshake messages.

​	QUICEncryptionLevel 表示用于传输握手消息的 QUIC 加密级别。

#### (QUICEncryptionLevel) String <-go1.21.0

``` go
func (l QUICEncryptionLevel) String() string
```

### type QUICEvent <-go1.21.0

``` go
type QUICEvent struct {
	Kind QUICEventKind

	// Set for QUICSetReadSecret, QUICSetWriteSecret, and QUICWriteData.
	// 针对 QUICSetReadSecret、QUICSetWriteSecret 和 QUICWriteData 设置。
	Level QUICEncryptionLevel

	// Set for QUICTransportParameters, QUICSetReadSecret, QUICSetWriteSecret, and QUICWriteData.
	// The contents are owned by crypto/tls, and are valid until the next NextEvent call.
	// 针对 QUICTransportParameters、QUICSetReadSecret、QUICSetWriteSecret 和 QUICWriteData 设置。
	// 内容由 crypto/tls 拥有，并在下一次调用 NextEvent 前有效。
	Data []byte

	// Set for QUICSetReadSecret and QUICSetWriteSecret.
	// 针对 QUICSetReadSecret 和 QUICSetWriteSecret 设置。
	Suite uint16

	// Set for QUICResumeSession and QUICStoreSession.
	// 针对 QUICResumeSession 和 QUICStoreSession 设置。
	SessionState *SessionState
}
```

A QUICEvent is an event occurring on a QUIC connection.

​	QUICEvent 是 QUIC 连接上发生的事件。

The type of event is specified by the Kind field. The contents of the other fields are kind-specific.

​	事件类型由 Kind 字段指定，其他字段的内容与类型相关。

### type QUICEventKind <-go1.21.0

``` go
type QUICEventKind int
```

A QUICEventKind is a type of operation on a QUIC connection.

​	QUICEventKind 是 QUIC 连接上的操作类型。

``` go
const (
	// QUICNoEvent indicates that there are no events available.
	// QUICNoEvent 表示没有可用事件。
	QUICNoEvent QUICEventKind = iota

	// QUICSetReadSecret and QUICSetWriteSecret provide the read and write
	// secrets for a given encryption level.
	// QUICEvent.Level, QUICEvent.Data, and QUICEvent.Suite are set.
	//
	// Secrets for the Initial encryption level are derived from the initial
	// destination connection ID, and are not provided by the QUICConn.
	// QUICSetReadSecret 和 QUICSetWriteSecret 提供给定加密级别的读取和写入密钥。
	// QUICEvent.Level、QUICEvent.Data 和 QUICEvent.Suite 已设置。
	//
	// 初始加密级别的密钥从初始目标连接 ID 派生，不由 QUICConn 提供。
	QUICSetReadSecret
	QUICSetWriteSecret

	// QUICWriteData provides data to send to the peer in CRYPTO frames.
	// QUICEvent.Data is set.
	// QUICWriteData 提供要在 CRYPTO 帧中发送给对等方的数据。
	// QUICEvent.Data 已设置。
	QUICWriteData

	// QUICTransportParameters provides the peer's QUIC transport parameters.
	// QUICEvent.Data is set.
	// QUICTransportParameters 提供对等方的 QUIC 传输参数。
	// QUICEvent.Data 已设置。
	QUICTransportParameters

	// QUICTransportParametersRequired indicates that the caller must provide
	// QUIC transport parameters to send to the peer. The caller should set
	// the transport parameters with QUICConn.SetTransportParameters and call
	// QUICConn.NextEvent again.
	//
	// If transport parameters are set before calling QUICConn.Start, the
	// connection will never generate a QUICTransportParametersRequired event.
	// QUICTransportParametersRequired 表示调用者必须提供要发送给对等方的 QUIC 传输参数。
	// 调用者应使用 QUICConn.SetTransportParameters 设置传输参数并再次调用 QUICConn.NextEvent。
	//
	// 如果在调用 QUICConn.Start 之前设置了传输参数，连接将永远不会生成 QUICTransportParametersRequired 事件。
	QUICTransportParametersRequired

	// QUICRejectedEarlyData indicates that the server rejected 0-RTT data even
	// if we offered it. It's returned before QUICEncryptionLevelApplication
	// keys are returned.
	// This event only occurs on client connections.
	// QUICRejectedEarlyData 表示服务器拒绝了 0-RTT 数据，即使我们提供了它。它在返回 QUICEncryptionLevelApplication 密钥之前返回。
	// 该事件仅发生在客户端连接上。
	QUICRejectedEarlyData

	// QUICHandshakeDone indicates that the TLS handshake has completed.
	// QUICHandshakeDone 表示 TLS 握手已完成。
	QUICHandshakeDone

	// QUICResumeSession indicates that a client is attempting to resume a previous session.
	// [QUICEvent.SessionState] is set.
	//
	// For client connections, this event occurs when the session ticket is selected.
	// For server connections, this event occurs when receiving the client's session ticket.
	//
	// The application may set [QUICEvent.SessionState.EarlyData] to false before the
	// next call to [QUICConn.NextEvent] to decline 0-RTT even if the session supports it.
	// QUICResumeSession 表示客户端正在尝试恢复先前的会话。
	// [QUICEvent.SessionState] 已设置。
	//
	// 对于客户端连接，此事件在选择会话票据时发生。
	// 对于服务器连接，此事件在接收客户端的会话票据时发生。
	//
	// 应用程序可以在下一次调用 [QUICConn.NextEvent] 之前将 [QUICEvent.SessionState.EarlyData] 设置为 false，以拒绝 0-RTT，即使会话支持它。
	QUICResumeSession

	// QUICStoreSession indicates that the server has provided state permitting
	// the client to resume the session.
	// [QUICEvent.SessionState] is set.
	// The application should use [QUICConn.StoreSession] session to store the [SessionState].
	// The application may modify the [SessionState] before storing it.
	// This event only occurs on client connections.
	// QUICStoreSession 表示服务器已提供状态，允许客户端恢复会话。
	// [QUICEvent.SessionState] 已设置。
	// 应用程序应使用 [QUICConn.StoreSession] 存储 [SessionState]。应用程序可以在存储之前修改 [SessionState]。
	// 该事件仅发生在客户端连接上。
	QUICStoreSession
)
```

### type QUICSessionTicketOptions <-go1.21.0

``` go
type QUICSessionTicketOptions struct {
	// EarlyData specifies whether the ticket may be used for 0-RTT.
	// EarlyData指定票据是否可以用于0-RTT。
	EarlyData bool
	Extra     [][]byte
}
```



### type RecordHeaderError  <- go1.6

``` go
type RecordHeaderError struct {
	// Msg contains a human readable string that describes the error.
	Msg string
	// RecordHeader contains the five bytes of TLS record header that
	// triggered the error.
	RecordHeader [5]byte
	// Conn provides the underlying net.Conn in the case that a client
	// sent an initial handshake that didn't look like TLS.
	// It is nil if there's already been a handshake or a TLS alert has
	// been written to the connection.
	Conn net.Conn
}
```

RecordHeaderError is returned when a TLS record header is invalid.

​	当 TLS 记录头无效时，将返回 RecordHeaderError。

#### (RecordHeaderError) Error  <- go1.6

``` go
func (e RecordHeaderError) Error() string
```

### type RenegotiationSupport  <- go1.7

``` go
type RenegotiationSupport int
```

RenegotiationSupport enumerates the different levels of support for TLS renegotiation. TLS renegotiation is the act of performing subsequent handshakes on a connection after the first. This significantly complicates the state machine and has been the source of numerous, subtle security issues. Initiating a renegotiation is not supported, but support for accepting renegotiation requests may be enabled.

​	RenegotiationSupport 枚举了对 TLS 重新协商的不同级别的支持。TLS 重新协商是在第一次握手后对连接执行后续握手的行为。这极大地复杂化了状态机，并且一直是许多微妙的安全问题的根源。不支持发起重新协商，但可以启用对接受重新协商请求的支持。

Even when enabled, the server may not change its identity between handshakes (i.e. the leaf certificate must be the same). Additionally, concurrent handshake and application data flow is not permitted so renegotiation can only be used with protocols that synchronise with the renegotiation, such as HTTPS.

​	即使启用，服务器也可能不会在握手之间更改其标识（即叶证书必须相同）。此外，不允许并发握手和应用程序数据流，因此重新协商只能与与重新协商同步的协议（例如 HTTPS）一起使用。

Renegotiation is not defined in TLS 1.3.

​	重新协商在 TLS 1.3 中未定义。

``` go
const (
	// RenegotiateNever disables renegotiation.
    // RenegotiateNever 禁用重新协商。
	RenegotiateNever RenegotiationSupport = iota

	// RenegotiateOnceAsClient allows a remote server to request
	// renegotiation once per connection.
    // RenegotiateOnceAsClient 允许远程服务器每个连接请求一次重新协商。
	RenegotiateOnceAsClient

	// RenegotiateFreelyAsClient allows a remote server to repeatedly
	// request renegotiation.
    // RenegotiateFreelyAsClient 允许远程服务器反复请求重新协商。
	RenegotiateFreelyAsClient
)
```

### type SessionState <-go1.21.0

``` go
type SessionState struct {

	// Extra is ignored by crypto/tls, but is encoded by [SessionState.Bytes]
	// and parsed by [ParseSessionState].
	//
	// This allows [Config.UnwrapSession]/[Config.WrapSession] and
	// [ClientSessionCache] implementations to store and retrieve additional
	// data alongside this session.
	//
	// To allow different layers in a protocol stack to share this field,
	// applications must only append to it, not replace it, and must use entries
	// that can be recognized even if out of order (for example, by starting
	// with an id and version prefix).
	// Extra 被 crypto/tls 忽略，但会被 [SessionState.Bytes] 编码，
	// 并通过 [ParseSessionState] 解析。
	//
	// 这允许 [Config.UnwrapSession]/[Config.WrapSession] 和
	// [ClientSessionCache] 实现存储和检索与该会话相关的其他数据。
	//
	// 为了允许协议栈中的不同层共享此字段，
	// 应用程序只能附加数据，而不能替换它，并且必须使用即使顺序混乱也能识别的条目
	// （例如，以 id 和版本前缀开头）。
	Extra [][]byte

	// EarlyData indicates whether the ticket can be used for 0-RTT in a QUIC
	// connection. The application may set this to false if it is true to
	// decline to offer 0-RTT even if supported.
	// EarlyData 表示该票据是否可以在 QUIC 连接中用于 0-RTT。
	// 如果该值为 true，应用程序可以将其设置为 false，
	// 即使支持 0-RTT 也可以选择不提供它。
	EarlyData bool
	// contains filtered or unexported fields
}
```

A SessionState is a resumable session.

​	SessionState 是一个可恢复的会话。

#### func ParseSessionState <-go1.21.0

``` go
func ParseSessionState(data []byte) (*SessionState, error)
```

ParseSessionState parses a [SessionState](https://pkg.go.dev/crypto/tls@go1.23.0#SessionState) encoded by [SessionState.Bytes](https://pkg.go.dev/crypto/tls@go1.23.0#SessionState.Bytes).

​	ParseSessionState 解析由 [SessionState.Bytes](https://pkg.go.dev/crypto/tls@go1.23.0#SessionState.Bytes) 编码的 [SessionState](https://pkg.go.dev/crypto/tls@go1.23.0#SessionState)。

#### (*SessionState) Bytes <-go1.21.0

``` go
func (s *SessionState) Bytes() ([]byte, error)
```

Bytes encodes the session, including any private fields, so that it can be parsed by [ParseSessionState](https://pkg.go.dev/crypto/tls@go1.23.0#ParseSessionState). The encoding contains secret values critical to the security of future and possibly past sessions.

​	Bytes 对会话进行编码，包括任何私有字段，以便可以通过 [ParseSessionState](https://pkg.go.dev/crypto/tls@go1.23.0#ParseSessionState) 解析。编码包含对未来甚至过去会话的安全性至关重要的秘密值。

The specific encoding should be considered opaque and may change incompatibly between Go versions.

​	具体的编码应被视为不透明的，并且可能在不同的 Go 版本之间不兼容地更改。

### type SignatureScheme  <- go1.8

``` go
type SignatureScheme uint16
```

SignatureScheme identifies a signature algorithm supported by TLS. See [RFC 8446, Section 4.2.3](https://rfc-editor.org/rfc/rfc8446.html#section-4.2.3).

​	SignatureScheme 标识 TLS 支持的签名算法。请参阅 RFC 8446 第 4.2.3 节。

``` go
const (
	// RSASSA-PKCS1-v1_5 algorithms.
	PKCS1WithSHA256 SignatureScheme = 0x0401
	PKCS1WithSHA384 SignatureScheme = 0x0501
	PKCS1WithSHA512 SignatureScheme = 0x0601

	// RSASSA-PSS algorithms with public key OID rsaEncryption.
	PSSWithSHA256 SignatureScheme = 0x0804
	PSSWithSHA384 SignatureScheme = 0x0805
	PSSWithSHA512 SignatureScheme = 0x0806

	// ECDSA algorithms. Only constrained to a specific curve in TLS 1.3.
	ECDSAWithP256AndSHA256 SignatureScheme = 0x0403
	ECDSAWithP384AndSHA384 SignatureScheme = 0x0503
	ECDSAWithP521AndSHA512 SignatureScheme = 0x0603

	// EdDSA algorithms.
	Ed25519 SignatureScheme = 0x0807

	// Legacy signature and hash algorithms for TLS 1.2.
	PKCS1WithSHA1 SignatureScheme = 0x0201
	ECDSAWithSHA1 SignatureScheme = 0x0203
)
```

#### (SignatureScheme) String  <- go1.15

``` go
func (i SignatureScheme) String() string
```

## Notes

## Bugs

- The crypto/tls package only implements some countermeasures against Lucky13 attacks on CBC-mode encryption, and only on SHA1 variants. See http://www.isg.rhul.ac.uk/tls/TLStiming.pdf and https://www.imperialviolet.org/2013/02/04/luckythirteen.html.
- crypto/tls 包仅对 CBC 模式加密中的 Lucky13 攻击实施了一些对策，并且仅对 SHA1 变体实施。请参阅 http://www.isg.rhul.ac.uk/tls/TLStiming.pdf 和 https://www.imperialviolet.org/2013/02/04/luckythirteen.html。