+++
title = "rand"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
# rand

https://pkg.go.dev/crypto/rand@go1.20.1



Package rand implements a cryptographically secure random number generator.












## 常量 ¶

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/crypto/rand/rand.go;l=20)

```
var Reader io.Reader
```

Reader is a global, shared instance of a cryptographically secure random number generator.

On Linux, FreeBSD, Dragonfly and Solaris, Reader uses getrandom(2) if available, /dev/urandom otherwise. On OpenBSD and macOS, Reader uses getentropy(2). On other Unix-like systems, Reader reads from /dev/urandom. On Windows systems, Reader uses the RtlGenRandom API. On Wasm, Reader uses the Web Crypto API.

## 函数

#### func Int 

```
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
```

Int returns a uniform random value in [0, max). It panics if max <= 0.

#### func Prime 

```
func Prime(rand io.Reader, bits int) (*big.Int, error)
```

Prime returns a number of the given bit length that is prime with high probability. Prime will return error for any error returned by rand.Read or if bits < 2.

#### func Read 

```
func Read(b []byte) (n int, err error)
```

Read is a helper function that calls Reader.Read using io.ReadFull. On return, n == len(b) if and only if err == nil.

<details tabindex="-1" id="example-Read" class="Documentation-exampleDetails js-exampleContainer" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-size: 16px; margin: 1rem 0px 0px; padding: 0px; vertical-align: baseline; display: block;"><summary class="Documentation-exampleDetailsHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-size: 16px; margin: 0px 0px 2rem; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); cursor: pointer; outline: none; text-decoration: none;">Example<span>&nbsp;</span><a href="https://pkg.go.dev/crypto/rand@go1.20.1#example-Read" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-brand-primary); text-decoration: none; opacity: 0;">¶</a></summary><div class="Documentation-exampleDetailsBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline;"><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><textarea class="Documentation-exampleCode code" spellcheck="false" style="box-sizing: border-box; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; font-size: 0.875rem; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; background-color: var(--color-background-accented); border: var(--border); border-top-left-radius: ; border-top-right-radius: ; border-bottom-right-radius: 0px; border-bottom-left-radius: 0px; color: var(--color-text); overflow-x: auto; padding: 0.625rem; tab-size: 4; white-space: pre; height: 27.125rem; outline: none; resize: none; width: 981.76px; margin: 0px;"></textarea><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-size: 0.875rem; margin: -0.25rem 0px 1rem; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: 0px 0px 0.3rem 0.3rem; color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre-wrap; word-break: break-all; overflow-wrap: break-word;"><span class="Documentation-exampleOutputLabel" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span><span class="Documentation-exampleOutput" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-size: 14px; margin: 0px 0px 0.5rem; padding: 0px; vertical-align: baseline; border-top-left-radius: 0px; border-top-right-radius: 0px;"></span></pre></div><div class="Documentation-exampleButtonsContainer" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-size: 16px; margin: 0.5rem 0px 0px; padding: 0px; vertical-align: baseline; align-items: center; display: flex; justify-content: flex-end;"><p class="Documentation-exampleError" role="alert" aria-atomic="true" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px 0.5rem 0px 0px; vertical-align: baseline; max-width: 60rem; color: var(--pink);"></p><button class="Documentation-exampleShareButton" aria-label="Share Code" style="box-sizing: border-box; border: 0.0625rem solid var(--turq-dark); font: inherit; margin: 0px 0px 0px 0.5rem; padding: 0px 1rem; vertical-align: baseline; border-radius: 0.25rem; cursor: pointer; height: 2rem; background-color: var(--white); color: var(--turq-dark);"></button><button class="Documentation-exampleFormatButton" aria-label="Format Code" style="box-sizing: border-box; border: 0.0625rem solid var(--turq-dark); font: inherit; margin: 0px 0px 0px 0.5rem; padding: 0px 1rem; vertical-align: baseline; border-radius: 0.25rem; cursor: pointer; height: 2rem; background-color: var(--white); color: var(--turq-dark);"></button><button class="Documentation-exampleRunButton" aria-label="Run Code" style="box-sizing: border-box; border: 0.0625rem solid var(--turq-dark); font: inherit; margin: 0px 0px 0px 0.5rem; padding: 0px 1rem; vertical-align: baseline; border-radius: 0.25rem; cursor: pointer; height: 2rem; background-color: var(--turq-dark); color: var(--white);"></button></div></details>

## 类型

This section is empty.