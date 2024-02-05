+++
title = "crypto/tls中的自动密码套件排序"
weight = 89
date = 2023-05-18T17:03:08+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Automatic cipher suite ordering in crypto/tls - crypto/tls中的自动密码套件排序

> 原文：[https://go.dev/blog/tls-cipher-suites](https://go.dev/blog/tls-cipher-suites)

Filippo Valsorda
15 September 2021

The Go standard library provides `crypto/tls`, a robust implementation of Transport Layer Security (TLS), the most important security protocol on the Internet, and the fundamental component of HTTPS. In Go 1.17 we made its configuration easier, more secure, and more efficient by automating the priority order of cipher suites.

Go标准库提供了crypto/tls，这是传输层安全（TLS）的强大实现，是互联网上最重要的安全协议，是HTTPS的基本组成部分。在Go 1.17中，我们通过自动调整密码套件的优先级顺序，使其配置更容易、更安全、更高效。

## How cipher suites work 密码套件如何工作

Cipher suites date back to TLS’s predecessor Secure Socket Layer (SSL), which [called them "cipher kinds"](https://datatracker.ietf.org/doc/html/draft-hickman-netscape-ssl-00#appendix-C.4). They are the intimidating-looking identifiers like `TLS_RSA_WITH_AES_256_CBC_SHA` and `TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256` that spell out the algorithms used to exchange keys, authenticate certificates, and encrypt records in a TLS connection.

密码套件可以追溯到TLS的前身安全套接字层（SSL），它称之为 "密码种类"。它们是看起来很吓人的标识符，如TLS_RSA_WITH_AES_256_CBC_SHA和TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256，阐明了用于在TLS连接中交换密钥、认证证书和加密记录的算法。

Cipher suites are *negotiated* during the TLS handshake: the client sends the list of cipher suites it supports in its first message, the Client Hello, and the server picks one from that list, communicating its choice to the client. The client sends the list of supported cipher suites in its own preference order, and the server is free to pick from it however it wants. Most commonly, the server will pick the first mutually supported cipher suite either in client preference order or in server preference order, based on its configuration.

密码套件是在TLS握手过程中协商的：客户端在其第一个消息中发送它所支持的密码套件列表，即 "客户端您好"，服务器从该列表中挑选一个，将其选择传达给客户端。客户端以自己的偏好顺序发送所支持的密码套件列表，而服务器可以自由地从中挑选它。最常见的是，服务器会根据其配置，按照客户端的偏好顺序或服务器的偏好顺序，选择第一个相互支持的密码套件。

Cipher suites are really only one of many negotiated parameters—supported curves/groups and signature algorithms are additionally negotiated through their own extensions—but are the most complex and famous ones, and the only ones that developers and administrators were trained over the years to have opinions on.

密码套件实际上只是许多协商参数中的一个，支持的曲线/组和签名算法是通过他们自己的扩展来协商的，但这是最复杂和最著名的参数，也是开发人员和管理员多年来被训练得有意见的唯一参数。

In TLS 1.0–1.2, all these parameters interact in a complex web of inter-dependencies: for example supported certificates depend on supported signature algorithms, supported curves, and supported cipher suites. In TLS 1.3 this was all drastically simplified: cipher suites only specify symmetric encryption algorithms, while supported curves/groups govern the key exchange and supported signature algorithms apply to the certificate.

在TLS 1.0-1.2中，所有这些参数在一个复杂的相互依赖的网络中相互作用：例如，支持的证书依赖于支持的签名算法、支持的曲线和支持的密码套件。在TLS 1.3中，这一切都被大大简化了：密码套件只指定对称加密算法，而支持的曲线/组管理密钥交换，支持的签名算法适用于证书。

## A complex choice abdicated to developers 一个复杂的选择被放弃给了开发者

Most HTTPS and TLS servers delegate the choice of cipher suites and preference order to the server operator or the applications developer. This is a complex choice that requires up-to-date and specialized knowledge for many reasons.

大多数HTTPS和TLS服务器将密码套件和偏好顺序的选择委托给服务器运营商或应用程序开发人员。这是一个复杂的选择，需要最新的和专门的知识，原因很多。

Some older cipher suites have insecure components, some require extremely careful and complex implementations to be secure, and some are only secure if the client applies certain mitigations or even has certain hardware. Beyond the security of the individual components, different cipher suites can provide drastically different security properties for the whole connection, as cipher suites without ECDHE or DHE don’t provide forward secrecy—the property that connections can’t be retroactively or passively decrypted with the certificate’s key. Finally, the choice of supported cipher suites impacts compatibility and performance, and making changes without an up-to-date understanding of the ecosystem can lead to breaking connections from legacy clients, increasing the resources consumed by the server, or draining the batteries of mobile clients.

一些旧的密码套件有不安全的组件，一些需要极其仔细和复杂的实现才是安全的，一些只有在客户端应用某些缓解措施或甚至拥有某些硬件时才是安全的。除了单个组件的安全性，不同的密码套件可以为整个连接提供截然不同的安全属性，因为没有ECDHE或DHE的密码套件不提供前向保密性--连接不能被追溯或被动地用证书的密钥解密的属性。最后，所支持的密码套件的选择会影响到兼容性和性能，在没有对生态系统进行最新了解的情况下进行改变，可能会导致传统客户端的连接中断，增加服务器消耗的资源，或耗尽移动客户端的电池。

This choice is so arcane and delicate that there are dedicated tools to guide operators, such as the excellent [Mozilla SSL Configuration Generator](https://ssl-config.mozilla.org/).

这种选择是如此神秘和微妙，以至于有专门的工具来指导操作者，如优秀的Mozilla SSL配置生成器。

How did we get here and why is it like this?

我们是如何走到这一步的，为什么会变成这样？

To start, individual cryptographic components used to break much more often. In 2011, when the BEAST attack broke CBC cipher suites in such a way that only clients could mitigate the attack, servers moved to preferring RC4, which was unaffected. In 2013, when it became clear that RC4 was broken, servers went back to CBC. When Lucky Thirteen made it clear it was extremely hard to implement CBC cipher suites due to their backwards MAC-then-encrypt design… Well, there wasn’t anything else on the table so implementations had to [carefully jump through hoops](https://www.imperialviolet.org/2013/02/04/luckythirteen.html) to implement CBC and kept [failing at that daunting task for years](https://blog.cloudflare.com/yet-another-padding-oracle-in-openssl-cbc-ciphersuites/). Configurable cipher suites and [cryptographic agility](https://www.imperialviolet.org/2016/05/16/agility.html) used to provide some reassurance that when a component broke it could be replaced on the fly.

首先，单个加密组件曾经更频繁地被破坏。2011年，当BEAST攻击破坏了CBC密码套件，以至于只有客户端可以减轻攻击，服务器转而倾向于使用不受影响的RC4。2013年，当RC4明显被破坏时，服务器又回到了CBC。当Lucky Thirteen明确表示，由于其向后的MAC-then-encrypt设计，实施CBC密码套件是非常困难的......好吧，没有任何其他的东西在桌子上，所以实施者不得不小心翼翼地跳过障碍来实施CBC，并且多年来一直在这个艰巨的任务中失败。可配置的密码套件和密码学的敏捷性曾经提供了一些保证，当一个组件损坏时，它可以被即时替换。

Modern cryptography is significantly different. Protocols can still break from time to time, but it’s rarely an individual abstracted component that fails. *None of the AEAD-based cipher suites introduced starting with TLS 1.2 in 2008 have been broken.* These days cryptographic agility is a liability: it introduces complexity that can lead to weaknesses or downgrades, and it is only necessary for performance and compliance reasons.

现代密码学则明显不同。协议仍然会时不时地发生故障，但很少是单个抽象的组件发生故障。从2008年TLS 1.2开始引入的基于AEAD的密码套件，没有一个被破解。这些天来，密码学的敏捷性是一种责任：它引入了可能导致弱点或降级的复杂性，而且它只是出于性能和合规性的原因才需要。

Patching used to be different, too. Today we acknowledge that promptly applying software patches for disclosed vulnerabilities is the cornerstone of secure software deployments, but ten years ago it was not standard practice. Changing configuration was seen as a much more rapid option to respond to vulnerable cipher suites, so the operator, through configuration, was put fully in charge of them. We now have the opposite problem: there are fully patched and updated servers that still behave weirdly, suboptimally, or insecurely, because their configurations haven’t been touched in years.

补丁曾经也是不同的。今天，我们承认及时为已披露的漏洞打上软件补丁是安全软件部署的基石，但在十年前这还不是标准做法。改变配置被认为是应对有漏洞的密码套件的一个更快速的选择，所以操作者通过配置，完全负责它们。我们现在有一个相反的问题：有一些完全打了补丁和更新的服务器仍然表现得很奇怪，不理想，或者不安全，因为他们的配置已经多年没有被触及。

Finally, it was understood that servers tended to update more slowly than clients, and therefore were less reliable judges of the best choice of cipher suite. However, it’s servers who have the last word on cipher suite selection, so the default became to make servers defer to the client preference order, instead of having strong opinions. This is still partially true: browsers managed to make automatic updates happen and are much more up-to-date than the average server. On the other hand, a number of legacy devices are now out of support and are stuck on old TLS client configurations, which often makes an up-to-date server better equipped to choose than some of its clients.

最后，人们了解到，服务器往往比客户端更新得更慢，因此对密码套件的最佳选择的判断不太可靠。然而，是服务器对密码套件的选择有最后的决定权，所以默认的做法是让服务器服从于客户端的偏好顺序，而不是有强烈的意见。这仍然是部分正确的：浏览器设法使自动更新发生，并且比一般的服务器要更新的多。另一方面，一些传统的设备现在已经不支持了，而且还停留在旧的TLS客户端配置上，这往往使一个最新的服务器比一些客户端更有能力去选择。

Regardless of how we got here, it’s a failure of cryptography engineering to require application developers and server operators to become experts in the nuances of cipher suite selection, and to stay up-to-date on the latest developments to keep their configs up-to-date. If they are deploying our security patches, that should be enough.

不管我们是如何走到这一步的，要求应用程序开发人员和服务器操作员成为密码套件选择的细微差别的专家，并保持最新的发展，以保持他们的配置是最新的，这是密码学工程的失败。如果他们部署了我们的安全补丁，这就足够了。

The Mozilla SSL Configuration Generator is great, and it should not exist.

Mozilla的SSL配置生成器很好，它不应该存在。

Is this getting any better?

这是否会变得更好？

There is good news and bad news for how things are trending in the past few years. The bad news is that ordering is getting even more nuanced, because there are sets of cipher suites that have equivalent security properties. The best choice within such a set depends on the available hardware and is hard to express in a config file. In other systems, what started as a simple list of cipher suites now depends on [more complex syntax](https://boringssl.googlesource.com/boringssl/+/c3b373bf4f4b2e2fba2578d1d5b5fe04e410f7cb/include/openssl/ssl.h#1457) or additional flags like [SSL_OP_PRIORITIZE_CHACHA](https://www.openssl.org/docs/man1.1.1/man3/SSL_CTX_clear_options.html#:~:text=session-,ssl_op_prioritize_chacha,-When).

在过去的几年里，事情的发展趋势有好消息和坏消息。坏消息是，排序变得更加细微，因为有几组密码套件具有同等的安全属性。在这样一个集合中的最佳选择取决于可用的硬件，并且很难在一个配置文件中表达。在其他系统中，开始时只是一个简单的密码套件列表，现在则取决于更复杂的语法或附加标志，如SSL_OP_PRIORITIZE_CHACHA。

The good news is that TLS 1.3 drastically simplified cipher suites, and it uses a disjoint set from TLS 1.0–1.2. All TLS 1.3 cipher suites are secure, so application developers and server operators shouldn’t have to worry about them at all. Indeed, some TLS libraries like BoringSSL and Go’s `crypto/tls` don’t allow configuring them at all.

好消息是，TLS 1.3极大地简化了密码套件，它使用的是与TLS 1.0-1.2不相干的集合。所有的TLS 1.3密码套件都是安全的，所以应用程序开发人员和服务器操作员根本不需要担心它们。事实上，一些TLS库如BoringSSL和Go的crypto/tls根本不允许配置它们。

## Go’s crypto/tls and cipher suites Go的crypto/tls和密码套件

Go does allow configuring cipher suites in TLS 1.0–1.2. Applications have always been able to set the enabled cipher suites and preference order with [`Config.CipherSuites`](https://pkg.go.dev/crypto/tls#Config.CipherSuites). Servers prioritize the client’s preference order by default, unless [`Config.PreferServerCipherSuites`](https://pkg.go.dev/crypto/tls#Config.PreferServerCipherSuites) is set.

Go确实允许在TLS 1.0-1.2中配置密码套件。应用程序一直能够通过 Config.CipherSuites 设置启用的密码套件和偏好顺序。除非Config.PreferServerCipherSuites被设置，否则服务器默认会优先考虑客户端的偏好顺序。

When we implemented TLS 1.3 in Go 1.12, [we didn’t make TLS 1.3 cipher suites configurable](https://go.dev/issue/29349), because they are a disjoint set from the TLS 1.0–1.2 ones and most importantly they are all secure, so there is no need to delegate a choice to the application. `Config.PreferServerCipherSuites` still controls which side’s preference order is used, and the local side’s preferences depend on the available hardware.

当我们在Go 1.12中实现TLS 1.3时，我们没有让TLS 1.3密码套件可配置，因为它们与TLS 1.0-1.2的密码套件是不相干的，最重要的是它们都是安全的，所以没有必要将选择权交给应用程序。Config.PreferServerCipherSuites仍然控制使用哪一方的偏好顺序，而本方的偏好取决于可用的硬件。

In Go 1.14 we [exposed supported cipher suites](https://pkg.go.dev/crypto/tls#CipherSuites), but explicitly chose to return them in a neutral order (sorted by their ID), so that we wouldn’t end up tied to representing our priority logic in terms of a static sort order.

在Go 1.14中，我们公开了支持的密码套件，但明确选择了以中性的顺序（按ID排序）返回，这样我们就不会最终被束缚在以静态排序的方式来表示我们的优先级逻辑。

In Go 1.16, we started actively [preferring ChaCha20Poly1305 cipher suites over AES-GCM on the server](https://go.dev/cl/262857) when we detect that either the client or the server lacks hardware support for AES-GCM. This is because AES-GCM is hard to implement efficiently and securely without dedicated hardware support (such as the AES-NI and CLMUL instruction sets).

在Go 1.16中，当我们检测到客户端或服务器缺乏对AES-GCM的硬件支持时，我们开始主动在服务器上选择ChaCha20Poly1305密码套件而不是AES-GCM。这是因为如果没有专门的硬件支持（如AES-NI和CLMUL指令集），AES-GCM很难有效和安全地实现。

**Go 1.17, recently released, takes over cipher suite preference ordering for all Go users.** While `Config.CipherSuites` still controls which TLS 1.0–1.2 cipher suites are enabled, it is not used for ordering, and `Config.PreferServerCipherSuites` is now ignored. Instead, `crypto/tls` [makes all ordering decisions](https://go.dev/cl/314609), based on the available cipher suites, the local hardware, and the inferred remote hardware capabilities.

最近发布的Go 1.17，接管了所有Go用户的密码套件偏好排序。虽然Config.CipherSuites仍然控制哪些TLS 1.0-1.2密码套件被启用，但它不用于排序，而且Config.PreferServerCipherSuites现在被忽略。相反，crypto/tls 会根据可用的密码套件、本地硬件和推断的远程硬件能力做出所有排序决定。

The [current TLS 1.0–1.2 ordering logic](https://cs.opensource.google/go/go/+/9d0819b27ca248f9949e7cf6bf7cb9fe7cf574e8:src/crypto/tls/cipher_suites.go;l=206-270) follows the following rules:

当前 TLS 1.0-1.2 的排序逻辑遵循以下规则：

1. ECDHE is preferred over the static RSA key exchange.ECDHE优于静态RSA密钥交换。

   The most important property of a cipher suite is enabling forward secrecy. We don’t implement "classic" finite-field Diffie-Hellman, because it’s complex, slower, weaker, and [subtly broken](https://datatracker.ietf.org/doc/draft-bartle-tls-deprecate-ffdh/) in TLS 1.0–1.2, so that means prioritizing the Elliptic Curve Diffie-Hellman key exchange over the legacy static RSA key exchange. (The latter simply encrypts the connection’s secret using the certificate’s public key, making it possible to decrypt if the certificate is compromised in the future.)密码套件最重要的属性是启用前向保密性。我们没有实现 "经典 "的有限字段Diffie-Hellman，因为它很复杂，速度较慢，较弱，而且在TLS 1.0-1.2中被巧妙地破坏了，所以这意味着优先考虑椭圆曲线Diffie-Hellman密钥交换而不是传统的静态RSA密钥交换。(后者只是使用证书的公钥对连接的秘密进行加密，如果证书在未来被破坏，就有可能被解密）。

2. AEAD modes are preferred over CBC for encryption.AEAD模式比CBC模式更适合用于加密。

   Even if we do implement partial countermeasures for Lucky13 ([my first contribution to the Go standard library, back in 2015!](https://go.dev/cl/18130)), the CBC suites are [a nightmare to get right](https://blog.cloudflare.com/yet-another-padding-oracle-in-openssl-cbc-ciphersuites/), so all other more important things being equal, we pick AES-GCM and ChaCha20Poly1305 instead.即使我们为Lucky13（我对Go标准库的第一个贡献，早在2015年！）实现了部分对策，CBC套件也是一场噩梦，所以在其他更重要的事情上，我们选择AES-GCM和ChaCha20Poly1305来代替。

3. 3DES, CBC-SHA256, and RC4 are only used if nothing else is available, in that preference order.3DES、CBC-SHA256和RC4只有在没有其他可用的情况下才会被使用，按照这个偏好顺序。

   3DES has 64-bit blocks, which makes it fundamentally vulnerable to [birthday attacks](https://sweet32.info/) given enough traffic. 3DES is listed under [`InsecureCipherSuites`](https://pkg.go.dev/crypto/tls#InsecureCipherSuites), but it’s enabled by default for compatibility. (An additional benefit of controlling preference orders is that we can afford to keep less secure cipher suites enabled by default without worrying about applications or clients selecting them except as a last resort. This is safe because there are no downgrade attacks that rely on the availability of a weaker cipher suite to attack peers that support better alternatives.)3DES有64位的块，这使得它在足够的流量下，从根本上容易受到生日攻击。3DES被列在InsecureCipherSuites下，但为了兼容，它被默认启用。(控制偏好顺序的另一个好处是，我们可以在默认情况下保持较不安全的密码套件，而不必担心应用程序或客户端选择它们，除非是最后的手段。这是安全的，因为没有降级攻击，即依靠较弱的密码套件的可用性来攻击支持更好的替代方案的对等者）。

   The CBC cipher suites are vulnerable to Lucky13-style side channel attacks and we only partially implement the [complex](https://www.imperialviolet.org/2013/02/04/luckythirteen.html) countermeasures discussed above for the SHA-1 hash, not for SHA-256. CBC-SHA1 suites have compatibility value, justifying the extra complexity, while the CBC-SHA256 ones don’t, so they are disabled by default.CBC密码套件容易受到Lucky13式的侧信道攻击，我们只为SHA-1哈希值部分实现了上面讨论的复杂对策，而不是为SHA-256。CBC-SHA1套件有兼容性价值，证明了额外的复杂性，而CBC-SHA256套件没有，所以它们默认是禁用的。

   RC4 has [practically exploitable biases](https://www.rc4nomore.com/) that can lead to plaintext recovery without side channels. It doesn’t get any worse than this, so RC4 is disabled by default.RC4有实际可利用的偏差，可以在没有侧信道的情况下导致明文恢复。没有比这更糟糕的了，所以RC4默认是禁用的。

4. ChaCha20Poly1305 is preferred over AES-GCM for encryption, unless both sides have hardware support for the latter.ChaCha20Poly1305比AES-GCM更适合用于加密，除非双方都有对后者的硬件支持。

   As we discussed above, AES-GCM is hard to implement efficiently and securely without hardware support. If we detect that there isn’t local hardware support or (on the server) that the client has not prioritized AES-GCM, we pick ChaCha20Poly1305 instead.正如我们上面所讨论的，AES-GCM在没有硬件支持的情况下很难有效和安全地实现。如果我们检测到本地没有硬件支持，或者（在服务器上）客户端没有优先考虑AES-GCM，我们会选择ChaCha20Poly1305来代替。

5. AES-128 is preferred over AES-256 for encryption.AES-128比AES-256更适合用于加密。

   AES-256 has a larger key than AES-128, which is usually good, but it also performs more rounds of the core encryption function, making it slower. (The extra rounds in AES-256 are independent of the key size change; they are an attempt to provide a wider margin against cryptanalysis.) The larger key is only useful in multi-user and post-quantum settings, which are not relevant to TLS, which generates sufficiently random IVs and has no post-quantum key exchange support. Since the larger key doesn’t have any benefit, we prefer AES-128 for its speed.AES-256比AES-128有更大的密钥，这通常是好的，但它也执行了更多轮的核心加密功能，使其更慢。(AES-256中的额外轮次与密钥大小的变化无关；它们是为了提供更广泛的余量来防止密码分析。) 更大的密钥只在多用户和后量子设置中有用，这与TLS无关，TLS生成足够随机的IV，并且没有后量子密钥交换支持。由于较大的密钥没有任何好处，我们更喜欢AES-128的速度。

[TLS 1.3’s ordering logic](https://cs.opensource.google/go/go/+/9d0819b27ca248f9949e7cf6bf7cb9fe7cf574e8:src/crypto/tls/cipher_suites.go;l=342-355) needs only the last two rules, because TLS 1.3 eliminated the problematic algorithms the first three rules are guarding against.

TLS 1.3的排序逻辑只需要最后两条规则，因为TLS 1.3消除了前三条规则所防范的问题算法。

## FAQs 常见问题

*What if a cipher suite turns out to be broken?* Just like any other vulnerability, it will be fixed in a security release for all supported Go versions. All applications need to be prepared to apply security fixes to operate securely. Historically, broken cipher suites are increasingly rare.

如果一个密码套件被发现有问题怎么办？就像其他的漏洞一样，它将在所有支持的Go版本的安全版本中被修复。所有的应用程序都需要准备好应用安全修复，以便安全地运行。从历史上看，破损的密码套件越来越少。

*Why leave enabled TLS 1.0–1.2 cipher suites configurable?* There is a meaningful tradeoff between *baseline* security and legacy compatibility to make in choosing which cipher suites to enable, and that’s a choice we can’t make ourselves without either cutting out an unacceptable slice of the ecosystem, or reducing the security guarantees of modern users.

为什么让启用的TLS 1.0-1.2密码套件可配置？在选择启用哪些密码套件时，需要在基线安全和传统兼容性之间进行有意义的权衡，如果不砍掉生态系统中不可接受的部分，或者减少现代用户的安全保障，我们就无法自己做出选择。

*Why not make TLS 1.3 cipher suites configurable?* Conversely, there is no tradeoff to make with TLS 1.3, as all its cipher suites provide strong security. This lets us leave them all enabled and pick the fastest based on the specifics of the connection without requiring the developer’s involvement.

为什么不使TLS 1.3密码套件可配置？相反，TLS 1.3没有必要做任何权衡，因为它的所有密码套件都提供了强大的安全性。这让我们可以将它们全部启用，并根据连接的具体情况选择最快的，而不需要开发者的参与。

## Key takeaways 主要启示

Starting in Go 1.17, `crypto/tls` is taking over the order in which available cipher suites are selected. With a regularly updated Go version, this is safer than letting potentially outdated clients pick the order, lets us optimize performance, and it lifts significant complexity from Go developers.

从Go 1.17开始，crypto/tls正在接管选择可用密码套件的顺序。在定期更新的 Go 版本中，这比让可能过时的客户端选择顺序更安全，让我们优化性能，并从 Go 开发者那里解除了大量的复杂性。

This is consistent with our general philosophy of making cryptographic decisions whenever we can, instead of delegating them to developers, and with our [cryptography principles](https://go.dev/design/cryptography-principles). Hopefully other TLS libraries will adopt similar changes, making delicate cipher suite configuration a thing of the past.

这与我们的总体理念是一致的，即只要有可能，我们就会做出加密决定，而不是将其委托给开发人员，这也符合我们的加密原则。希望其他TLS库也能采用类似的变化，使精细的密码套件配置成为过去。
