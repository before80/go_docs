+++
title = "image"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# image

https://pkg.go.dev/image@go1.20.1



Package image implements a basic 2-D image library.

The fundamental interface is called Image. An Image contains colors, which are described in the image/color package.

Values of the Image interface are created either by calling functions such as NewRGBA and NewPaletted, or by calling Decode on an io.Reader containing image data in a format such as GIF, JPEG or PNG. Decoding any particular image format requires the prior registration of a decoder function. Registration is typically automatic as a side effect of initializing that format's package so that, to decode a PNG image, it suffices to have

```
import _ "image/png"
```

in a program's main package. The _ means to import a package purely for its initialization side effects.

See "The Go image package" for more details: https://golang.org/doc/articles/image_package.html

##### Example
``` go linenums="1"
```

##### Example
``` go linenums="1"
```








## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=11)

``` go linenums="1"
var (
	// Black is an opaque black uniform image.
	Black = NewUniform(color.Black)
	// White is an opaque white uniform image.
	White = NewUniform(color.White)
	// Transparent is a fully transparent uniform image.
	Transparent = NewUniform(color.Transparent)
	// Opaque is a fully opaque uniform image.
	Opaque = NewUniform(color.Opaque)
)
```

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/image/format.go;l=16)

``` go linenums="1"
var ErrFormat = errors.New("image: unknown format")
```

ErrFormat indicates that decoding encountered an unknown format.

## 函数

#### func [RegisterFormat](https://cs.opensource.google/go/go/+/go1.20.1:src/image/format.go;l=37) 

``` go linenums="1"
func RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error))
```

RegisterFormat registers an image format for use by Decode. Name is the name of the format, like "jpeg" or "png". Magic is the magic prefix that identifies the format's encoding. The magic string can contain "?" wildcards that each match any one byte. Decode is the function that decodes the encoded image. DecodeConfig is the function that decodes just its configuration.

## 类型

### type [Alpha](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=610) 

``` go linenums="1"
type Alpha struct {
	// Pix holds the image's pixels, as alpha values. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}
```

Alpha is an in-memory image whose At method returns color.Alpha values.

#### func [NewAlpha](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=709) 

``` go linenums="1"
func NewAlpha(r Rectangle) *Alpha
```

NewAlpha returns a new Alpha image with the given bounds.

#### (*Alpha) [AlphaAt](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=634)  <- go1.4

``` go linenums="1"
func (p *Alpha) AlphaAt(x, y int) color.Alpha
```

#### (*Alpha) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=624) 

``` go linenums="1"
func (p *Alpha) At(x, y int) color.Color
```

#### (*Alpha) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=622) 

``` go linenums="1"
func (p *Alpha) Bounds() Rectangle
```

#### (*Alpha) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=620) 

``` go linenums="1"
func (p *Alpha) ColorModel() color.Model
```

#### (*Alpha) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=691) 

``` go linenums="1"
func (p *Alpha) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Alpha) [PixOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=644) 

``` go linenums="1"
func (p *Alpha) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*Alpha) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=628)  <- go1.17

``` go linenums="1"
func (p *Alpha) RGBA64At(x, y int) color.RGBA64
```

#### (*Alpha) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=648) 

``` go linenums="1"
func (p *Alpha) Set(x, y int, c color.Color)
```

#### (*Alpha) [SetAlpha](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=664) 

``` go linenums="1"
func (p *Alpha) SetAlpha(x, y int, c color.Alpha)
```

#### (*Alpha) [SetRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=656)  <- go1.17

``` go linenums="1"
func (p *Alpha) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Alpha) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=674) 

``` go linenums="1"
func (p *Alpha) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [Alpha16](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=718) 

``` go linenums="1"
type Alpha16 struct {
	// Pix holds the image's pixels, as alpha values in big-endian format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}
```

Alpha16 is an in-memory image whose At method returns color.Alpha16 values.

#### func [NewAlpha16](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=820) 

``` go linenums="1"
func NewAlpha16(r Rectangle) *Alpha16
```

NewAlpha16 returns a new Alpha16 image with the given bounds.

#### (*Alpha16) [Alpha16At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=741)  <- go1.4

``` go linenums="1"
func (p *Alpha16) Alpha16At(x, y int) color.Alpha16
```

#### (*Alpha16) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=732) 

``` go linenums="1"
func (p *Alpha16) At(x, y int) color.Color
```

#### (*Alpha16) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=730) 

``` go linenums="1"
func (p *Alpha16) Bounds() Rectangle
```

#### (*Alpha16) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=728) 

``` go linenums="1"
func (p *Alpha16) ColorModel() color.Model
```

#### (*Alpha16) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=802) 

``` go linenums="1"
func (p *Alpha16) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Alpha16) [PixOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=751) 

``` go linenums="1"
func (p *Alpha16) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*Alpha16) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=736)  <- go1.17

``` go linenums="1"
func (p *Alpha16) RGBA64At(x, y int) color.RGBA64
```

#### (*Alpha16) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=755) 

``` go linenums="1"
func (p *Alpha16) Set(x, y int, c color.Color)
```

#### (*Alpha16) [SetAlpha16](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=774) 

``` go linenums="1"
func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16)
```

#### (*Alpha16) [SetRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=765)  <- go1.17

``` go linenums="1"
func (p *Alpha16) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Alpha16) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=785) 

``` go linenums="1"
func (p *Alpha16) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [CMYK](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1026)  <- go1.5

``` go linenums="1"
type CMYK struct {
	// Pix holds the image's pixels, in C, M, Y, K order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}
```

CMYK is an in-memory image whose At method returns color.CMYK values.

#### func [NewCMYK](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1126)  <- go1.5

``` go linenums="1"
func NewCMYK(r Rectangle) *CMYK
```

NewCMYK returns a new CMYK image with the given bounds.

#### (*CMYK) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1040)  <- go1.5

``` go linenums="1"
func (p *CMYK) At(x, y int) color.Color
```

#### (*CMYK) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1038)  <- go1.5

``` go linenums="1"
func (p *CMYK) Bounds() Rectangle
```

#### (*CMYK) [CMYKAt](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1049)  <- go1.5

``` go linenums="1"
func (p *CMYK) CMYKAt(x, y int) color.CMYK
```

#### (*CMYK) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1036)  <- go1.5

``` go linenums="1"
func (p *CMYK) ColorModel() color.Model
```

#### (*CMYK) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1121)  <- go1.5

``` go linenums="1"
func (p *CMYK) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*CMYK) [PixOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1060)  <- go1.5

``` go linenums="1"
func (p *CMYK) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*CMYK) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1044)  <- go1.17

``` go linenums="1"
func (p *CMYK) RGBA64At(x, y int) color.RGBA64
```

#### (*CMYK) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1064)  <- go1.5

``` go linenums="1"
func (p *CMYK) Set(x, y int, c color.Color)
```

#### (*CMYK) [SetCMYK](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1090)  <- go1.5

``` go linenums="1"
func (p *CMYK) SetCMYK(x, y int, c color.CMYK)
```

#### (*CMYK) [SetRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1077)  <- go1.17

``` go linenums="1"
func (p *CMYK) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*CMYK) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1104)  <- go1.5

``` go linenums="1"
func (p *CMYK) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [Config](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=31) 

``` go linenums="1"
type Config struct {
	ColorModel    color.Model
	Width, Height int
}
```

Config holds an image's color model and dimensions.

#### func [DecodeConfig](https://cs.opensource.google/go/go/+/go1.20.1:src/image/format.go;l=101) 

``` go linenums="1"
func DecodeConfig(r io.Reader) (Config, string, error)
```

DecodeConfig decodes the color model and dimensions of an image that has been encoded in a registered format. The string returned is the format name used during format registration. Format registration is typically done by an init function in the codec-specific package.

### type [Gray](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=829) 

``` go linenums="1"
type Gray struct {
	// Pix holds the image's pixels, as gray values. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}
```

Gray is an in-memory image whose At method returns color.Gray values.

#### func [NewGray](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=917) 

``` go linenums="1"
func NewGray(r Rectangle) *Gray
```

NewGray returns a new Gray image with the given bounds.

#### (*Gray) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=843) 

``` go linenums="1"
func (p *Gray) At(x, y int) color.Color
```

#### (*Gray) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=841) 

``` go linenums="1"
func (p *Gray) Bounds() Rectangle
```

#### (*Gray) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=839) 

``` go linenums="1"
func (p *Gray) ColorModel() color.Model
```

#### (*Gray) [GrayAt](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=853)  <- go1.4

``` go linenums="1"
func (p *Gray) GrayAt(x, y int) color.Gray
```

#### (*Gray) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=912) 

``` go linenums="1"
func (p *Gray) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Gray) [PixOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=863) 

``` go linenums="1"
func (p *Gray) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*Gray) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=847)  <- go1.17

``` go linenums="1"
func (p *Gray) RGBA64At(x, y int) color.RGBA64
```

#### (*Gray) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=867) 

``` go linenums="1"
func (p *Gray) Set(x, y int, c color.Color)
```

#### (*Gray) [SetGray](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=885) 

``` go linenums="1"
func (p *Gray) SetGray(x, y int, c color.Gray)
```

#### (*Gray) [SetRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=875)  <- go1.17

``` go linenums="1"
func (p *Gray) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Gray) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=895) 

``` go linenums="1"
func (p *Gray) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [Gray16](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=926) 

``` go linenums="1"
type Gray16 struct {
	// Pix holds the image's pixels, as gray values in big-endian format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}
```

Gray16 is an in-memory image whose At method returns color.Gray16 values.

#### func [NewGray16](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1017) 

``` go linenums="1"
func NewGray16(r Rectangle) *Gray16
```

NewGray16 returns a new Gray16 image with the given bounds.

#### (*Gray16) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=940) 

``` go linenums="1"
func (p *Gray16) At(x, y int) color.Color
```

#### (*Gray16) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=938) 

``` go linenums="1"
func (p *Gray16) Bounds() Rectangle
```

#### (*Gray16) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=936) 

``` go linenums="1"
func (p *Gray16) ColorModel() color.Model
```

#### (*Gray16) [Gray16At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=949)  <- go1.4

``` go linenums="1"
func (p *Gray16) Gray16At(x, y int) color.Gray16
```

#### (*Gray16) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1012) 

``` go linenums="1"
func (p *Gray16) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Gray16) [PixOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=959) 

``` go linenums="1"
func (p *Gray16) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*Gray16) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=944)  <- go1.17

``` go linenums="1"
func (p *Gray16) RGBA64At(x, y int) color.RGBA64
```

#### (*Gray16) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=963) 

``` go linenums="1"
func (p *Gray16) Set(x, y int, c color.Color)
```

#### (*Gray16) [SetGray16](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=984) 

``` go linenums="1"
func (p *Gray16) SetGray16(x, y int, c color.Gray16)
```

#### (*Gray16) [SetRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=973)  <- go1.17

``` go linenums="1"
func (p *Gray16) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Gray16) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=995) 

``` go linenums="1"
func (p *Gray16) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [Image](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=38) 

``` go linenums="1"
type Image interface {
	// ColorModel returns the Image's color model.
	ColorModel() color.Model
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	Bounds() Rectangle
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	At(x, y int) color.Color
}
```

Image is a finite rectangular grid of color.Color values taken from a color model.

#### func [Decode](https://cs.opensource.google/go/go/+/go1.20.1:src/image/format.go;l=87) 

``` go linenums="1"
func Decode(r io.Reader) (Image, string, error)
```

Decode decodes an image that has been encoded in a registered format. The string returned is the format name used during format registration. Format registration is typically done by an init function in the codec- specific package.

### type [NRGBA](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=339) 

``` go linenums="1"
type NRGBA struct {
	// Pix holds the image's pixels, in R, G, B, A order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}
```

NRGBA is an in-memory image whose At method returns color.NRGBA values.

#### func [NewNRGBA](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=457) 

``` go linenums="1"
func NewNRGBA(r Rectangle) *NRGBA
```

NewNRGBA returns a new NRGBA image with the given bounds.

#### (*NRGBA) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=353) 

``` go linenums="1"
func (p *NRGBA) At(x, y int) color.Color
```

#### (*NRGBA) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=351) 

``` go linenums="1"
func (p *NRGBA) Bounds() Rectangle
```

#### (*NRGBA) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=349) 

``` go linenums="1"
func (p *NRGBA) ColorModel() color.Model
```

#### (*NRGBA) [NRGBAAt](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=362)  <- go1.4

``` go linenums="1"
func (p *NRGBA) NRGBAAt(x, y int) color.NRGBA
```

#### (*NRGBA) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=439) 

``` go linenums="1"
func (p *NRGBA) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*NRGBA) [PixOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=373) 

``` go linenums="1"
func (p *NRGBA) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*NRGBA) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=357)  <- go1.17

``` go linenums="1"
func (p *NRGBA) RGBA64At(x, y int) color.RGBA64
```

#### (*NRGBA) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=377) 

``` go linenums="1"
func (p *NRGBA) Set(x, y int, c color.Color)
```

#### (*NRGBA) [SetNRGBA](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=408) 

``` go linenums="1"
func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA)
```

#### (*NRGBA) [SetRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=390)  <- go1.17

``` go linenums="1"
func (p *NRGBA) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*NRGBA) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=422) 

``` go linenums="1"
func (p *NRGBA) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [NRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=466) 

``` go linenums="1"
type NRGBA64 struct {
	// Pix holds the image's pixels, in R, G, B, A order and big-endian format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*8].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}
```

NRGBA64 is an in-memory image whose At method returns color.NRGBA64 values.

#### func [NewNRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=601) 

``` go linenums="1"
func NewNRGBA64(r Rectangle) *NRGBA64
```

NewNRGBA64 returns a new NRGBA64 image with the given bounds.

#### (*NRGBA64) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=480) 

``` go linenums="1"
func (p *NRGBA64) At(x, y int) color.Color
```

#### (*NRGBA64) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=478) 

``` go linenums="1"
func (p *NRGBA64) Bounds() Rectangle
```

#### (*NRGBA64) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=476) 

``` go linenums="1"
func (p *NRGBA64) ColorModel() color.Model
```

#### (*NRGBA64) [NRGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=489)  <- go1.4

``` go linenums="1"
func (p *NRGBA64) NRGBA64At(x, y int) color.NRGBA64
```

#### (*NRGBA64) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=583) 

``` go linenums="1"
func (p *NRGBA64) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*NRGBA64) [PixOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=505) 

``` go linenums="1"
func (p *NRGBA64) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*NRGBA64) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=484)  <- go1.17

``` go linenums="1"
func (p *NRGBA64) RGBA64At(x, y int) color.RGBA64
```

#### (*NRGBA64) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=509) 

``` go linenums="1"
func (p *NRGBA64) Set(x, y int, c color.Color)
```

#### (*NRGBA64) [SetNRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=548) 

``` go linenums="1"
func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64)
```

#### (*NRGBA64) [SetRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=526)  <- go1.17

``` go linenums="1"
func (p *NRGBA64) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*NRGBA64) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=566) 

``` go linenums="1"
func (p *NRGBA64) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [NYCbCrA](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=205)  <- go1.6

``` go linenums="1"
type NYCbCrA struct {
	YCbCr
	A       []uint8
	AStride int
}
```

NYCbCrA is an in-memory image of non-alpha-premultiplied Y'CbCr-with-alpha colors. A and AStride are analogous to the Y and YStride fields of the embedded YCbCr.

#### func [NewNYCbCrA](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=299)  <- go1.6

``` go linenums="1"
func NewNYCbCrA(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *NYCbCrA
```

NewNYCbCrA returns a new NYCbCrA image with the given bounds and subsample ratio.

#### (*NYCbCrA) [AOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=243)  <- go1.6

``` go linenums="1"
func (p *NYCbCrA) AOffset(x, y int) int
```

AOffset returns the index of the first element of A that corresponds to the pixel at (x, y).

#### (*NYCbCrA) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=215)  <- go1.6

``` go linenums="1"
func (p *NYCbCrA) At(x, y int) color.Color
```

#### (*NYCbCrA) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=211)  <- go1.6

``` go linenums="1"
func (p *NYCbCrA) ColorModel() color.Model
```

#### (*NYCbCrA) [NYCbCrAAt](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=224)  <- go1.6

``` go linenums="1"
func (p *NYCbCrA) NYCbCrAAt(x, y int) color.NYCbCrA
```

#### (*NYCbCrA) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=280)  <- go1.6

``` go linenums="1"
func (p *NYCbCrA) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*NYCbCrA) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=219)  <- go1.17

``` go linenums="1"
func (p *NYCbCrA) RGBA64At(x, y int) color.RGBA64
```

#### (*NYCbCrA) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=249)  <- go1.6

``` go linenums="1"
func (p *NYCbCrA) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [Paletted](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1135) 

``` go linenums="1"
type Paletted struct {
	// Pix holds the image's pixels, as palette indices. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*1].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
	// Palette is the image's palette.
	Palette color.Palette
}
```

Paletted is an in-memory image of uint8 indices into a given palette.

#### func [NewPaletted](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1266) 

``` go linenums="1"
func NewPaletted(r Rectangle, p color.Palette) *Paletted
```

NewPaletted returns a new Paletted image with the given width, height and palette.

#### (*Paletted) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1151) 

``` go linenums="1"
func (p *Paletted) At(x, y int) color.Color
```

#### (*Paletted) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1149) 

``` go linenums="1"
func (p *Paletted) Bounds() Rectangle
```

#### (*Paletted) [ColorIndexAt](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1204) 

``` go linenums="1"
func (p *Paletted) ColorIndexAt(x, y int) uint8
```

#### (*Paletted) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1147) 

``` go linenums="1"
func (p *Paletted) ColorModel() color.Model
```

#### (*Paletted) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1242) 

``` go linenums="1"
func (p *Paletted) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Paletted) [PixOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1184) 

``` go linenums="1"
func (p *Paletted) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*Paletted) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1162)  <- go1.17

``` go linenums="1"
func (p *Paletted) RGBA64At(x, y int) color.RGBA64
```

#### (*Paletted) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1188) 

``` go linenums="1"
func (p *Paletted) Set(x, y int, c color.Color)
```

#### (*Paletted) [SetColorIndex](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1212) 

``` go linenums="1"
func (p *Paletted) SetColorIndex(x, y int, index uint8)
```

#### (*Paletted) [SetRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1196)  <- go1.17

``` go linenums="1"
func (p *Paletted) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Paletted) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=1222) 

``` go linenums="1"
func (p *Paletted) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [PalettedImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=66) 

``` go linenums="1"
type PalettedImage interface {
	// ColorIndexAt returns the palette index of the pixel at (x, y).
	ColorIndexAt(x, y int) uint8
	Image
}
```

PalettedImage is an image whose colors may come from a limited palette. If m is a PalettedImage and m.ColorModel() returns a color.Palette p, then m.At(x, y) should be equivalent to p[m.ColorIndexAt(x, y)]. If m's color model is not a color.Palette, then ColorIndexAt's behavior is undefined.

### type [Point](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=14) 

``` go linenums="1"
type Point struct {
	X, Y int
}
```

A Point is an X, Y coordinate pair. The axes increase right and down.

``` go linenums="1"
var ZP Point
```

ZP is the zero Point.

Deprecated: Use a literal image.Point{} instead.

#### func [Pt](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=76) 

``` go linenums="1"
func Pt(X, Y int) Point
```

Pt is shorthand for Point{X, Y}.

#### (Point) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=24) 

``` go linenums="1"
func (p Point) Add(q Point) Point
```

Add returns the vector p+q.

#### (Point) [Div](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=39) 

``` go linenums="1"
func (p Point) Div(k int) Point
```

Div returns the vector p/k.

#### (Point) [Eq](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=66) 

``` go linenums="1"
func (p Point) Eq(q Point) bool
```

Eq reports whether p and q are equal.

#### (Point) [In](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=44) 

``` go linenums="1"
func (p Point) In(r Rectangle) bool
```

In reports whether p is in r.

#### (Point) [Mod](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=51) 

``` go linenums="1"
func (p Point) Mod(r Rectangle) Point
```

Mod returns the point q in r such that p.X-q.X is a multiple of r's width and p.Y-q.Y is a multiple of r's height.

#### (Point) [Mul](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=34) 

``` go linenums="1"
func (p Point) Mul(k int) Point
```

Mul returns the vector p*k.

#### (Point) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=19) 

``` go linenums="1"
func (p Point) String() string
```

String returns a string representation of p like "(3,4)".

#### (Point) [Sub](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=29) 

``` go linenums="1"
func (p Point) Sub(q Point) Point
```

Sub returns the vector p-q.

### type [RGBA](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=88) 

``` go linenums="1"
type RGBA struct {
	// Pix holds the image's pixels, in R, G, B, A order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}
```

RGBA is an in-memory image whose At method returns color.RGBA values.

#### func [NewRGBA](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=213) 

``` go linenums="1"
func NewRGBA(r Rectangle) *RGBA
```

NewRGBA returns a new RGBA image with the given bounds.

#### (*RGBA) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=102) 

``` go linenums="1"
func (p *RGBA) At(x, y int) color.Color
```

#### (*RGBA) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=100) 

``` go linenums="1"
func (p *RGBA) Bounds() Rectangle
```

#### (*RGBA) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=98) 

``` go linenums="1"
func (p *RGBA) ColorModel() color.Model
```

#### (*RGBA) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=195) 

``` go linenums="1"
func (p *RGBA) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*RGBA) [PixOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=135) 

``` go linenums="1"
func (p *RGBA) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*RGBA) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=106)  <- go1.17

``` go linenums="1"
func (p *RGBA) RGBA64At(x, y int) color.RGBA64
```

#### (*RGBA) [RGBAAt](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=124)  <- go1.4

``` go linenums="1"
func (p *RGBA) RGBAAt(x, y int) color.RGBA
```

#### (*RGBA) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=139) 

``` go linenums="1"
func (p *RGBA) Set(x, y int, c color.Color)
```

#### (*RGBA) [SetRGBA](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=164) 

``` go linenums="1"
func (p *RGBA) SetRGBA(x, y int, c color.RGBA)
```

#### (*RGBA) [SetRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=152)  <- go1.17

``` go linenums="1"
func (p *RGBA) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*RGBA) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=178) 

``` go linenums="1"
func (p *RGBA) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [RGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=222) 

``` go linenums="1"
type RGBA64 struct {
	// Pix holds the image's pixels, in R, G, B, A order and big-endian format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*8].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}
```

RGBA64 is an in-memory image whose At method returns color.RGBA64 values.

#### func [NewRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=330) 

``` go linenums="1"
func NewRGBA64(r Rectangle) *RGBA64
```

NewRGBA64 returns a new RGBA64 image with the given bounds.

#### (*RGBA64) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=236) 

``` go linenums="1"
func (p *RGBA64) At(x, y int) color.Color
```

#### (*RGBA64) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=234) 

``` go linenums="1"
func (p *RGBA64) Bounds() Rectangle
```

#### (*RGBA64) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=232) 

``` go linenums="1"
func (p *RGBA64) ColorModel() color.Model
```

#### (*RGBA64) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=312) 

``` go linenums="1"
func (p *RGBA64) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*RGBA64) [PixOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=256) 

``` go linenums="1"
func (p *RGBA64) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*RGBA64) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=240)  <- go1.4

``` go linenums="1"
func (p *RGBA64) RGBA64At(x, y int) color.RGBA64
```

#### (*RGBA64) [Set](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=260) 

``` go linenums="1"
func (p *RGBA64) Set(x, y int, c color.Color)
```

#### (*RGBA64) [SetRGBA64](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=277) 

``` go linenums="1"
func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*RGBA64) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=295) 

``` go linenums="1"
func (p *RGBA64) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type [RGBA64Image](https://cs.opensource.google/go/go/+/go1.20.1:src/image/image.go;l=52)  <- go1.17

``` go linenums="1"
type RGBA64Image interface {
	// RGBA64At returns the RGBA64 color of the pixel at (x, y). It is
	// equivalent to calling At(x, y).RGBA() and converting the resulting
	// 32-bit return values to a color.RGBA64, but it can avoid allocations
	// from converting concrete color types to the color.Color interface type.
	RGBA64At(x, y int) color.RGBA64
	Image
}
```

RGBA64Image is an Image whose pixels can be converted directly to a color.RGBA64.

### type [Rectangle](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=88) 

``` go linenums="1"
type Rectangle struct {
	Min, Max Point
}
```

A Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y. It is well-formed if Min.X <= Max.X and likewise for Y. Points are always well-formed. A rectangle's methods always return well-formed outputs for well-formed inputs.

A Rectangle is also an Image whose bounds are the rectangle itself. At returns color.Opaque for points in the rectangle and color.Transparent otherwise.

``` go linenums="1"
var ZR Rectangle
```

ZR is the zero Rectangle.

Deprecated: Use a literal image.Rectangle{} instead.

#### func [Rect](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=275) 

``` go linenums="1"
func Rect(x0, y0, x1, y1 int) Rectangle
```

Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}. The returned rectangle has minimum and maximum coordinates swapped if necessary so that it is well-formed.

#### (Rectangle) [Add](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=116) 

``` go linenums="1"
func (r Rectangle) Add(p Point) Rectangle
```

Add returns the rectangle r translated by p.

#### (Rectangle) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=242)  <- go1.5

``` go linenums="1"
func (r Rectangle) At(x, y int) color.Color
```

At implements the Image interface.

#### (Rectangle) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=258)  <- go1.5

``` go linenums="1"
func (r Rectangle) Bounds() Rectangle
```

Bounds implements the Image interface.

#### (Rectangle) [Canon](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=231) 

``` go linenums="1"
func (r Rectangle) Canon() Rectangle
```

Canon returns the canonical version of r. The returned rectangle has minimum and maximum coordinates swapped if necessary so that it is well-formed.

#### (Rectangle) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=263)  <- go1.5

``` go linenums="1"
func (r Rectangle) ColorModel() color.Model
```

ColorModel implements the Image interface.

#### (Rectangle) [Dx](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=98) 

``` go linenums="1"
func (r Rectangle) Dx() int
```

Dx returns r's width.

#### (Rectangle) [Dy](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=103) 

``` go linenums="1"
func (r Rectangle) Dy() int
```

Dy returns r's height.

#### (Rectangle) [Empty](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=201) 

``` go linenums="1"
func (r Rectangle) Empty() bool
```

Empty reports whether the rectangle contains no points.

#### (Rectangle) [Eq](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=207) 

``` go linenums="1"
func (r Rectangle) Eq(s Rectangle) bool
```

Eq reports whether r and s contain the same set of points. All empty rectangles are considered equal.

#### (Rectangle) [In](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=219) 

``` go linenums="1"
func (r Rectangle) In(s Rectangle) bool
```

In reports whether every point in r is in s.

#### (Rectangle) [Inset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=134) 

``` go linenums="1"
func (r Rectangle) Inset(n int) Rectangle
```

Inset returns the rectangle r inset by n, which may be negative. If either of r's dimensions is less than 2*n then an empty rectangle near the center of r will be returned.

#### (Rectangle) [Intersect](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=154) 

``` go linenums="1"
func (r Rectangle) Intersect(s Rectangle) Rectangle
```

Intersect returns the largest rectangle contained by both r and s. If the two rectangles do not overlap then the zero rectangle will be returned.

#### (Rectangle) [Overlaps](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=212) 

``` go linenums="1"
func (r Rectangle) Overlaps(s Rectangle) bool
```

Overlaps reports whether r and s have a non-empty intersection.

#### (Rectangle) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=250)  <- go1.17

``` go linenums="1"
func (r Rectangle) RGBA64At(x, y int) color.RGBA64
```

RGBA64At implements the RGBA64Image interface.

#### (Rectangle) [Size](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=108) 

``` go linenums="1"
func (r Rectangle) Size() Point
```

Size returns r's width and height.

#### (Rectangle) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=93) 

``` go linenums="1"
func (r Rectangle) String() string
```

String returns a string representation of r like "(3,4)-(6,5)".

#### (Rectangle) [Sub](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=124) 

``` go linenums="1"
func (r Rectangle) Sub(p Point) Rectangle
```

Sub returns the rectangle r translated by -p.

#### (Rectangle) [Union](https://cs.opensource.google/go/go/+/go1.20.1:src/image/geom.go;l=178) 

``` go linenums="1"
func (r Rectangle) Union(s Rectangle) Rectangle
```

Union returns the smallest rectangle that contains both r and s.

### type [Uniform](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=24) 

``` go linenums="1"
type Uniform struct {
	C color.Color
}
```

Uniform is an infinite-sized Image of uniform color. It implements the color.Color, color.Model, and Image interfaces.

#### func [NewUniform](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=56) 

``` go linenums="1"
func NewUniform(c color.Color) *Uniform
```

NewUniform returns a new Uniform image of the given color.

#### (*Uniform) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=42) 

``` go linenums="1"
func (c *Uniform) At(x, y int) color.Color
```

#### (*Uniform) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=40) 

``` go linenums="1"
func (c *Uniform) Bounds() Rectangle
```

#### (*Uniform) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=32) 

``` go linenums="1"
func (c *Uniform) ColorModel() color.Model
```

#### (*Uniform) [Convert](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=36) 

``` go linenums="1"
func (c *Uniform) Convert(color.Color) color.Color
```

#### (*Uniform) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=50) 

``` go linenums="1"
func (c *Uniform) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Uniform) [RGBA](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=28) 

``` go linenums="1"
func (c *Uniform) RGBA() (r, g, b, a uint32)
```

#### (*Uniform) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=44)  <- go1.17

``` go linenums="1"
func (c *Uniform) RGBA64At(x, y int) color.RGBA64
```

### type [YCbCr](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=55) 

``` go linenums="1"
type YCbCr struct {
	Y, Cb, Cr      []uint8
	YStride        int
	CStride        int
	SubsampleRatio YCbCrSubsampleRatio
	Rect           Rectangle
}
```

YCbCr is an in-memory image of Y'CbCr colors. There is one Y sample per pixel, but each Cb and Cr sample can span one or more pixels. YStride is the Y slice index delta between vertically adjacent pixels. CStride is the Cb and Cr slice index delta between vertically adjacent pixels that map to separate chroma samples. It is not an absolute requirement, but YStride and len(Y) are typically multiples of 8, and:

```
For 4:4:4, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1.
For 4:2:2, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2.
For 4:2:0, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4.
For 4:4:0, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2.
For 4:1:1, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4.
For 4:1:0, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8.
```

#### func [NewYCbCr](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=175) 

``` go linenums="1"
func NewYCbCr(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *YCbCr
```

NewYCbCr returns a new YCbCr image with the given bounds and subsample ratio.

#### (*YCbCr) [At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=71) 

``` go linenums="1"
func (p *YCbCr) At(x, y int) color.Color
```

#### (*YCbCr) [Bounds](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=67) 

``` go linenums="1"
func (p *YCbCr) Bounds() Rectangle
```

#### (*YCbCr) [COffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=101) 

``` go linenums="1"
func (p *YCbCr) COffset(x, y int) int
```

COffset returns the index of the first element of Cb or Cr that corresponds to the pixel at (x, y).

#### (*YCbCr) [ColorModel](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=63) 

``` go linenums="1"
func (p *YCbCr) ColorModel() color.Model
```

#### (*YCbCr) [Opaque](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=143) 

``` go linenums="1"
func (p *YCbCr) Opaque() bool
```

#### (*YCbCr) [RGBA64At](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=75)  <- go1.17

``` go linenums="1"
func (p *YCbCr) RGBA64At(x, y int) color.RGBA64
```

#### (*YCbCr) [SubImage](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=120) 

``` go linenums="1"
func (p *YCbCr) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

#### (*YCbCr) [YCbCrAt](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=80)  <- go1.4

``` go linenums="1"
func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr
```

#### (*YCbCr) [YOffset](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=95) 

``` go linenums="1"
func (p *YCbCr) YOffset(x, y int) int
```

YOffset returns the index of the first element of Y that corresponds to the pixel at (x, y).

### type [YCbCrSubsampleRatio](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=12) 

``` go linenums="1"
type YCbCrSubsampleRatio int
```

YCbCrSubsampleRatio is the chroma subsample ratio used in a YCbCr image.

``` go linenums="1"
const (
	YCbCrSubsampleRatio444 YCbCrSubsampleRatio = iota
	YCbCrSubsampleRatio422
	YCbCrSubsampleRatio420
	YCbCrSubsampleRatio440
	YCbCrSubsampleRatio411
	YCbCrSubsampleRatio410
)
```

#### (YCbCrSubsampleRatio) [String](https://cs.opensource.google/go/go/+/go1.20.1:src/image/ycbcr.go;l=23) 

``` go linenums="1"
func (s YCbCrSubsampleRatio) String() string
```