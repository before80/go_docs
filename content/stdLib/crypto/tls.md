+++
title = "tls"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/crypto/tls@go1.21.3](https://pkg.go.dev/crypto/tls@go1.21.3)

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

## 类型

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

### type Config 

``` go
type Config struct {
	// Rand provides the source of entropy for nonces and RSA blinding.
	// If Rand is nil, TLS uses the cryptographic random reader in package
	// crypto/rand.
	// The Reader must be safe for use by multiple goroutines.
	Rand io.Reader

	// Time returns the current time as the number of seconds since the epoch.
	// If Time is nil, TLS uses time.Now.
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
	Certificates []Certificate

	// NameToCertificate maps from a certificate name to an element of
	// Certificates. Note that a certificate name can be of the form
	// '*.example.com' and so doesn't have to be a domain name as such.
	//
	// Deprecated: NameToCertificate only allows associating a single
	// certificate with a given name. Leave this field nil to let the library
	// select the first compatible chain from Certificates.
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
	GetConfigForClient func(*ClientHelloInfo) (*Config, error)

	// VerifyPeerCertificate, if not nil, is called after normal
	// certificate verification by either a TLS client or server. It
	// receives the raw ASN.1 certificates provided by the peer and also
	// any verified chains that normal processing found. If it returns a
	// non-nil error, the handshake is aborted and that error results.
	//
	// If normal verification fails then the handshake will abort before
	// considering this callback. If normal verification is disabled by
	// setting InsecureSkipVerify, or (for a server) when ClientAuth is
	// RequestClientCert or RequireAnyClientCert, then this callback will
	// be considered but the verifiedChains argument will always be nil.
	//
	// verifiedChains and its contents should not be modified.
	VerifyPeerCertificate func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error

	// VerifyConnection, if not nil, is called after normal certificate
	// verification and after VerifyPeerCertificate by either a TLS client
	// or server. If it returns a non-nil error, the handshake is aborted
	// and that error results.
	//
	// If normal verification fails then the handshake will abort before
	// considering this callback. This callback will run for all connections
	// regardless of InsecureSkipVerify or ClientAuth settings.
	VerifyConnection func(ConnectionState) error

	// RootCAs defines the set of root certificate authorities
	// that clients use when verifying server certificates.
	// If RootCAs is nil, TLS uses the host's root CA set.
	RootCAs *x509.CertPool

	// NextProtos is a list of supported application level protocols, in
	// order of preference. If both peers support ALPN, the selected
	// protocol will be one from this list, and the connection will fail
	// if there is no mutually supported protocol. If NextProtos is empty
	// or the peer doesn't support ALPN, the connection will succeed and
	// ConnectionState.NegotiatedProtocol will be empty.
	NextProtos []string

	// ServerName is used to verify the hostname on the returned
	// certificates unless InsecureSkipVerify is given. It is also included
	// in the client's handshake to support virtual hosting unless it is
	// an IP address.
	ServerName string

	// ClientAuth determines the server's policy for
	// TLS Client Authentication. The default is NoClientCert.
	ClientAuth ClientAuthType

	// ClientCAs defines the set of root certificate authorities
	// that servers use if required to verify a client certificate
	// by the policy in ClientAuth.
	ClientCAs *x509.CertPool

	// InsecureSkipVerify controls whether a client verifies the server's
	// certificate chain and host name. If InsecureSkipVerify is true, crypto/tls
	// accepts any certificate presented by the server and any host name in that
	// certificate. In this mode, TLS is susceptible to machine-in-the-middle
	// attacks unless custom verification is used. This should be used only for
	// testing or in combination with VerifyConnection or VerifyPeerCertificate.
	InsecureSkipVerify bool

	// CipherSuites is a list of enabled TLS 1.0–1.2 cipher suites. The order of
	// the list is ignored. Note that TLS 1.3 ciphersuites are not configurable.
	//
	// If CipherSuites is nil, a safe default list is used. The default cipher
	// suites might change over time.
	CipherSuites []uint16

	// PreferServerCipherSuites is a legacy field and has no effect.
	//
	// It used to control whether the server would follow the client's or the
	// server's preference. Servers now select the best mutually supported
	// cipher suite based on logic that takes into account inferred client
	// hardware, server hardware, and security.
	//
	// Deprecated: PreferServerCipherSuites is ignored.
	PreferServerCipherSuites bool

	// SessionTicketsDisabled may be set to true to disable session ticket and
	// PSK (resumption) support. Note that on clients, session ticket support is
	// also disabled if ClientSessionCache is nil.
	SessionTicketsDisabled bool

	// SessionTicketKey is used by TLS servers to provide session resumption.
	// See RFC 5077 and the PSK mode of RFC 8446. If zero, it will be filled
	// with random data before the first server handshake.
	//
	// Deprecated: if this field is left at zero, session ticket keys will be
	// automatically rotated every day and dropped after seven days. For
	// customizing the rotation schedule or synchronizing servers that are
	// terminating connections for the same host, use SetSessionTicketKeys.
	SessionTicketKey [32]byte

	// ClientSessionCache is a cache of ClientSessionState entries for TLS
	// session resumption. It is only used by clients.
	ClientSessionCache ClientSessionCache

	// MinVersion contains the minimum TLS version that is acceptable.
	//
	// By default, TLS 1.2 is currently used as the minimum when acting as a
	// client, and TLS 1.0 when acting as a server. TLS 1.0 is the minimum
	// supported by this package, both as a client and as a server.
	//
	// The client-side default can temporarily be reverted to TLS 1.0 by
	// including the value "x509sha1=1" in the GODEBUG environment variable.
	// Note that this option will be removed in Go 1.19 (but it will still be
	// possible to set this field to VersionTLS10 explicitly).
	MinVersion uint16

	// MaxVersion contains the maximum TLS version that is acceptable.
	//
	// By default, the maximum version supported by this package is used,
	// which is currently TLS 1.3.
	MaxVersion uint16

	// CurvePreferences contains the elliptic curves that will be used in
	// an ECDHE handshake, in preference order. If empty, the default will
	// be used. The client will use the first preference as the type for
	// its key share in TLS 1.3. This may change in the future.
	CurvePreferences []CurveID

	// DynamicRecordSizingDisabled disables adaptive sizing of TLS records.
	// When true, the largest possible TLS record size is always used. When
	// false, the size of TLS records may be adjusted in an attempt to
	// improve latency.
	DynamicRecordSizingDisabled bool

	// Renegotiation controls what types of renegotiation are supported.
	// The default, none, is correct for the vast majority of applications.
	Renegotiation RenegotiationSupport

	// KeyLogWriter optionally specifies a destination for TLS master secrets
	// in NSS key log format that can be used to allow external programs
	// such as Wireshark to decrypt TLS connections.
	// See https://developer.mozilla.org/en-US/docs/Mozilla/Projects/NSS/Key_Log_Format.
	// Use of KeyLogWriter compromises security and should only be
	// used for debugging.
	KeyLogWriter io.Writer
	// contains filtered or unexported fields
}
```

A Config structure is used to configure a TLS client or server. After one has been passed to a TLS function it must not be modified. A Config may be reused; the tls package will also not modify it.

​	Config 结构用于配置 TLS 客户端或服务器。将一个结构传递给 TLS 函数后，不得对其进行修改。Config 可以重复使用；tls 包也不会修改它。

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



#### func (*Config) BuildNameToCertificate <- DEPRECATED

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

​	TLS 1.3 中未定义重新协商。

``` go
const (
	// RenegotiateNever disables renegotiation.
	RenegotiateNever RenegotiationSupport = iota

	// RenegotiateOnceAsClient allows a remote server to request
	// renegotiation once per connection.
	RenegotiateOnceAsClient

	// RenegotiateFreelyAsClient allows a remote server to repeatedly
	// request renegotiation.
	RenegotiateFreelyAsClient
)
```

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