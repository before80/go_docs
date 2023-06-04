+++
title = "png"
date = 2023-05-17T11:11:20+08:00
type = "docs"
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

#### func Decode 

``` go 
func Decode(r io.Reader) (image.Image, error)
```

Decode reads a PNG image from r and returns it as an image.Image. The type of Image returned depends on the PNG contents.

##### Example
``` go 
```

#### func DecodeConfig 

``` go 
func DecodeConfig(r io.Reader) (image.Config, error)
```

DecodeConfig returns the color model and dimensions of a PNG image without decoding the entire image.

#### func Encode 

``` go 
func Encode(w io.Writer, m image.Image) error
```

Encode writes the Image m to w in PNG format. Any Image may be encoded, but images that are not image.NRGBA might be encoded lossily.

##### Example
``` go 
```

## 类型

### type CompressionLevel  <- go1.4

``` go 
type CompressionLevel int
```

CompressionLevel indicates the compression level.

``` go 
const (
	DefaultCompression CompressionLevel = 0
	NoCompression      CompressionLevel = -1
	BestSpeed          CompressionLevel = -2
	BestCompression    CompressionLevel = -3
)
```

### type Encoder  <- go1.4

``` go 
type Encoder struct {
	CompressionLevel CompressionLevel

	// BufferPool optionally specifies a buffer pool to get temporary
	// EncoderBuffers when encoding an image.
	BufferPool EncoderBufferPool
}
```

Encoder configures encoding PNG images.

#### (*Encoder) Encode  <- go1.4

``` go 
func (enc *Encoder) Encode(w io.Writer, m image.Image) error
```

Encode writes the Image m to w in PNG format.

### type EncoderBuffer  <- go1.9

``` go 
type EncoderBuffer encoder
```

EncoderBuffer holds the buffers used for encoding PNG images.

### type EncoderBufferPool  <- go1.9

``` go 
type EncoderBufferPool interface {
	Get() *EncoderBuffer
	Put(*EncoderBuffer)
}
```

EncoderBufferPool is an interface for getting and returning temporary instances of the EncoderBuffer struct. This can be used to reuse buffers when encoding multiple images.

### type FormatError 

``` go 
type FormatError string
```

A FormatError reports that the input is not a valid PNG.

#### (FormatError) Error 

``` go 
func (e FormatError) Error() string
```

### type UnsupportedError 

``` go 
type UnsupportedError string
```

An UnsupportedError reports that the input uses a valid but unimplemented PNG feature.

#### (UnsupportedError) Error 

``` go 
func (e UnsupportedError) Error() string
```