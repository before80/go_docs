+++
title = "png"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# png

https://pkg.go.dev/image/png@go1.20.1



Package png implements a PNG image decoder and encoder.

The PNG specification is at https://www.w3.org/TR/PNG/.








## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

#### func [Decode](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/reader.go;l=976) 

``` go linenums="1"
func Decode(r io.Reader) (image.Image, error)
```

Decode reads a PNG image from r and returns it as an image.Image. The type of Image returned depends on the PNG contents.

##### Example
``` go linenums="1"
```

#### func [DecodeConfig](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/reader.go;l=1000) 

``` go linenums="1"
func DecodeConfig(r io.Reader) (image.Config, error)
```

DecodeConfig returns the color model and dimensions of a PNG image without decoding the entire image.

#### func [Encode](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/writer.go;l=590) 

``` go linenums="1"
func Encode(w io.Writer, m image.Image) error
```

Encode writes the Image m to w in PNG format. Any Image may be encoded, but images that are not image.NRGBA might be encoded lossily.

##### Example
``` go linenums="1"
```

## 类型

### type [CompressionLevel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/writer.go;l=55)  <- go1.4

``` go linenums="1"
type CompressionLevel int
```

CompressionLevel indicates the compression level.

``` go linenums="1"
const (
	DefaultCompression CompressionLevel = 0
	NoCompression      CompressionLevel = -1
	BestSpeed          CompressionLevel = -2
	BestCompression    CompressionLevel = -3
)
```

### type [Encoder](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/writer.go;l=19)  <- go1.4

``` go linenums="1"
type Encoder struct {
	CompressionLevel CompressionLevel

	// BufferPool optionally specifies a buffer pool to get temporary
	// EncoderBuffers when encoding an image.
	BufferPool EncoderBufferPool
}
```

Encoder configures encoding PNG images.

#### (*Encoder) [Encode](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/writer.go;l=596)  <- go1.4

``` go linenums="1"
func (enc *Encoder) Encode(w io.Writer, m image.Image) error
```

Encode writes the Image m to w in PNG format.

### type [EncoderBuffer](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/writer.go;l=36)  <- go1.9

``` go linenums="1"
type EncoderBuffer encoder
```

EncoderBuffer holds the buffers used for encoding PNG images.

### type [EncoderBufferPool](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/writer.go;l=30)  <- go1.9

``` go linenums="1"
type EncoderBufferPool interface {
	Get() *EncoderBuffer
	Put(*EncoderBuffer)
}
```

EncoderBufferPool is an interface for getting and returning temporary instances of the EncoderBuffer struct. This can be used to reuse buffers when encoding multiple images.

### type [FormatError](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/reader.go;l=128) 

``` go linenums="1"
type FormatError string
```

A FormatError reports that the input is not a valid PNG.

#### (FormatError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/reader.go;l=130) 

``` go linenums="1"
func (e FormatError) Error() string
```

### type [UnsupportedError](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/reader.go;l=135) 

``` go linenums="1"
type UnsupportedError string
```

An UnsupportedError reports that the input uses a valid but unimplemented PNG feature.

#### (UnsupportedError) [Error](https://cs.opensource.google/go/go/+/go1.20.1:src/image/png/reader.go;l=137) 

``` go linenums="1"
func (e UnsupportedError) Error() string
```