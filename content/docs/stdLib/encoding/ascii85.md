+++
title = "ascii85"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# ascii85

https://pkg.go.dev/encoding/ascii85@go1.20.1



Package ascii85 implements the ascii85 data encoding as used in the btoa tool and Adobe's PostScript and PDF document formats.



## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [Decode](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/ascii85/ascii85.go;l=186) 

``` go 
func Decode(dst, src []byte, flush bool) (ndst, nsrc int, err error)
```

Decode decodes src into dst, returning both the number of bytes written to dst and the number consumed from src. If src contains invalid ascii85 data, Decode will return the number of bytes successfully written and a CorruptInputError. Decode ignores space and control characters in src. Often, ascii85-encoded data is wrapped in <~ and ~> symbols. Decode expects these to have been stripped by the caller.

If flush is true, Decode assumes that src represents the end of the input stream and processes it completely rather than wait for the completion of another 32-bit block.

NewDecoder wraps an io.Reader interface around Decode.

#### func [Encode](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/ascii85/ascii85.go;l=27) 

``` go 
func Encode(dst, src []byte) int
```

Encode encodes src into at most MaxEncodedLen(len(src)) bytes of dst, returning the actual number of bytes written.

The encoding handles 4-byte chunks, using a special encoding for the last fragment, so Encode is not appropriate for use on individual blocks of a large data stream. Use NewEncoder() instead.

Often, ascii85-encoded data is wrapped in <~ and ~> symbols. Encode does not add these.

#### func [MaxEncodedLen](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/ascii85/ascii85.go;l=86) 

``` go 
func MaxEncodedLen(n int) int
```

MaxEncodedLen returns the maximum length of an encoding of n source bytes.

#### func [NewDecoder](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/ascii85/ascii85.go;l=243) 

``` go 
func NewDecoder(r io.Reader) io.Reader
```

NewDecoder constructs a new ascii85 stream decoder.

#### func [NewEncoder](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/ascii85/ascii85.go;l=93) 

``` go 
func NewEncoder(w io.Writer) io.WriteCloser
```

NewEncoder returns a new ascii85 stream encoder. Data written to the returned writer will be encoded and then written to w. Ascii85 encodings operate in 32-bit blocks; when finished writing, the caller must Close the returned encoder to flush any trailing partial block.

## 类型

### type [CorruptInputError](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/ascii85/ascii85.go;l=167) 

``` go 
type CorruptInputError int64
```

#### (CorruptInputError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/encoding/ascii85/ascii85.go;l=169) 

``` go 
func (e CorruptInputError) Error() string
```