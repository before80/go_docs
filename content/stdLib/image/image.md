+++
title = "image"
date = 2023-05-17T11:11:20+08:00
type = "docs"
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
``` go 
```

##### Example
``` go 
```








## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=11)

``` go 
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

``` go 
var ErrFormat = errors.New("image: unknown format")
```

ErrFormat indicates that decoding encountered an unknown format.

## 函数

#### func RegisterFormat 

``` go 
func RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error))
```

RegisterFormat registers an image format for use by Decode. Name is the name of the format, like "jpeg" or "png". Magic is the magic prefix that identifies the format's encoding. The magic string can contain "?" wildcards that each match any one byte. Decode is the function that decodes the encoded image. DecodeConfig is the function that decodes just its configuration.

## 类型

### type Alpha 

``` go 
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

#### func NewAlpha 

``` go 
func NewAlpha(r Rectangle) *Alpha
```

NewAlpha returns a new Alpha image with the given bounds.

#### (*Alpha) AlphaAt  <- go1.4

``` go 
func (p *Alpha) AlphaAt(x, y int) color.Alpha
```

#### (*Alpha) At 

``` go 
func (p *Alpha) At(x, y int) color.Color
```

#### (*Alpha) Bounds 

``` go 
func (p *Alpha) Bounds() Rectangle
```

#### (*Alpha) ColorModel 

``` go 
func (p *Alpha) ColorModel() color.Model
```

#### (*Alpha) Opaque 

``` go 
func (p *Alpha) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Alpha) PixOffset 

``` go 
func (p *Alpha) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*Alpha) RGBA64At  <- go1.17

``` go 
func (p *Alpha) RGBA64At(x, y int) color.RGBA64
```

#### (*Alpha) Set 

``` go 
func (p *Alpha) Set(x, y int, c color.Color)
```

#### (*Alpha) SetAlpha 

``` go 
func (p *Alpha) SetAlpha(x, y int, c color.Alpha)
```

#### (*Alpha) SetRGBA64  <- go1.17

``` go 
func (p *Alpha) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Alpha) SubImage 

``` go 
func (p *Alpha) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type Alpha16 

``` go 
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

#### func NewAlpha16 

``` go 
func NewAlpha16(r Rectangle) *Alpha16
```

NewAlpha16 returns a new Alpha16 image with the given bounds.

#### (*Alpha16) Alpha16At  <- go1.4

``` go 
func (p *Alpha16) Alpha16At(x, y int) color.Alpha16
```

#### (*Alpha16) At 

``` go 
func (p *Alpha16) At(x, y int) color.Color
```

#### (*Alpha16) Bounds 

``` go 
func (p *Alpha16) Bounds() Rectangle
```

#### (*Alpha16) ColorModel 

``` go 
func (p *Alpha16) ColorModel() color.Model
```

#### (*Alpha16) Opaque 

``` go 
func (p *Alpha16) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Alpha16) PixOffset 

``` go 
func (p *Alpha16) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*Alpha16) RGBA64At  <- go1.17

``` go 
func (p *Alpha16) RGBA64At(x, y int) color.RGBA64
```

#### (*Alpha16) Set 

``` go 
func (p *Alpha16) Set(x, y int, c color.Color)
```

#### (*Alpha16) SetAlpha16 

``` go 
func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16)
```

#### (*Alpha16) SetRGBA64  <- go1.17

``` go 
func (p *Alpha16) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Alpha16) SubImage 

``` go 
func (p *Alpha16) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type CMYK  <- go1.5

``` go 
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

#### func NewCMYK  <- go1.5

``` go 
func NewCMYK(r Rectangle) *CMYK
```

NewCMYK returns a new CMYK image with the given bounds.

#### (*CMYK) At  <- go1.5

``` go 
func (p *CMYK) At(x, y int) color.Color
```

#### (*CMYK) Bounds  <- go1.5

``` go 
func (p *CMYK) Bounds() Rectangle
```

#### (*CMYK) CMYKAt  <- go1.5

``` go 
func (p *CMYK) CMYKAt(x, y int) color.CMYK
```

#### (*CMYK) ColorModel  <- go1.5

``` go 
func (p *CMYK) ColorModel() color.Model
```

#### (*CMYK) Opaque  <- go1.5

``` go 
func (p *CMYK) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*CMYK) PixOffset  <- go1.5

``` go 
func (p *CMYK) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*CMYK) RGBA64At  <- go1.17

``` go 
func (p *CMYK) RGBA64At(x, y int) color.RGBA64
```

#### (*CMYK) Set  <- go1.5

``` go 
func (p *CMYK) Set(x, y int, c color.Color)
```

#### (*CMYK) SetCMYK  <- go1.5

``` go 
func (p *CMYK) SetCMYK(x, y int, c color.CMYK)
```

#### (*CMYK) SetRGBA64  <- go1.17

``` go 
func (p *CMYK) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*CMYK) SubImage  <- go1.5

``` go 
func (p *CMYK) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type Config 

``` go 
type Config struct {
	ColorModel    color.Model
	Width, Height int
}
```

Config holds an image's color model and dimensions.

#### func DecodeConfig 

``` go 
func DecodeConfig(r io.Reader) (Config, string, error)
```

DecodeConfig decodes the color model and dimensions of an image that has been encoded in a registered format. The string returned is the format name used during format registration. Format registration is typically done by an init function in the codec-specific package.

### type Gray 

``` go 
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

#### func NewGray 

``` go 
func NewGray(r Rectangle) *Gray
```

NewGray returns a new Gray image with the given bounds.

#### (*Gray) At 

``` go 
func (p *Gray) At(x, y int) color.Color
```

#### (*Gray) Bounds 

``` go 
func (p *Gray) Bounds() Rectangle
```

#### (*Gray) ColorModel 

``` go 
func (p *Gray) ColorModel() color.Model
```

#### (*Gray) GrayAt  <- go1.4

``` go 
func (p *Gray) GrayAt(x, y int) color.Gray
```

#### (*Gray) Opaque 

``` go 
func (p *Gray) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Gray) PixOffset 

``` go 
func (p *Gray) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*Gray) RGBA64At  <- go1.17

``` go 
func (p *Gray) RGBA64At(x, y int) color.RGBA64
```

#### (*Gray) Set 

``` go 
func (p *Gray) Set(x, y int, c color.Color)
```

#### (*Gray) SetGray 

``` go 
func (p *Gray) SetGray(x, y int, c color.Gray)
```

#### (*Gray) SetRGBA64  <- go1.17

``` go 
func (p *Gray) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Gray) SubImage 

``` go 
func (p *Gray) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type Gray16 

``` go 
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

#### func NewGray16 

``` go 
func NewGray16(r Rectangle) *Gray16
```

NewGray16 returns a new Gray16 image with the given bounds.

#### (*Gray16) At 

``` go 
func (p *Gray16) At(x, y int) color.Color
```

#### (*Gray16) Bounds 

``` go 
func (p *Gray16) Bounds() Rectangle
```

#### (*Gray16) ColorModel 

``` go 
func (p *Gray16) ColorModel() color.Model
```

#### (*Gray16) Gray16At  <- go1.4

``` go 
func (p *Gray16) Gray16At(x, y int) color.Gray16
```

#### (*Gray16) Opaque 

``` go 
func (p *Gray16) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Gray16) PixOffset 

``` go 
func (p *Gray16) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*Gray16) RGBA64At  <- go1.17

``` go 
func (p *Gray16) RGBA64At(x, y int) color.RGBA64
```

#### (*Gray16) Set 

``` go 
func (p *Gray16) Set(x, y int, c color.Color)
```

#### (*Gray16) SetGray16 

``` go 
func (p *Gray16) SetGray16(x, y int, c color.Gray16)
```

#### (*Gray16) SetRGBA64  <- go1.17

``` go 
func (p *Gray16) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Gray16) SubImage 

``` go 
func (p *Gray16) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type Image 

``` go 
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

#### func Decode 

``` go 
func Decode(r io.Reader) (Image, string, error)
```

Decode decodes an image that has been encoded in a registered format. The string returned is the format name used during format registration. Format registration is typically done by an init function in the codec- specific package.

### type NRGBA 

``` go 
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

#### func NewNRGBA 

``` go 
func NewNRGBA(r Rectangle) *NRGBA
```

NewNRGBA returns a new NRGBA image with the given bounds.

#### (*NRGBA) At 

``` go 
func (p *NRGBA) At(x, y int) color.Color
```

#### (*NRGBA) Bounds 

``` go 
func (p *NRGBA) Bounds() Rectangle
```

#### (*NRGBA) ColorModel 

``` go 
func (p *NRGBA) ColorModel() color.Model
```

#### (*NRGBA) NRGBAAt  <- go1.4

``` go 
func (p *NRGBA) NRGBAAt(x, y int) color.NRGBA
```

#### (*NRGBA) Opaque 

``` go 
func (p *NRGBA) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*NRGBA) PixOffset 

``` go 
func (p *NRGBA) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*NRGBA) RGBA64At  <- go1.17

``` go 
func (p *NRGBA) RGBA64At(x, y int) color.RGBA64
```

#### (*NRGBA) Set 

``` go 
func (p *NRGBA) Set(x, y int, c color.Color)
```

#### (*NRGBA) SetNRGBA 

``` go 
func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA)
```

#### (*NRGBA) SetRGBA64  <- go1.17

``` go 
func (p *NRGBA) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*NRGBA) SubImage 

``` go 
func (p *NRGBA) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type NRGBA64 

``` go 
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

#### func NewNRGBA64 

``` go 
func NewNRGBA64(r Rectangle) *NRGBA64
```

NewNRGBA64 returns a new NRGBA64 image with the given bounds.

#### (*NRGBA64) At 

``` go 
func (p *NRGBA64) At(x, y int) color.Color
```

#### (*NRGBA64) Bounds 

``` go 
func (p *NRGBA64) Bounds() Rectangle
```

#### (*NRGBA64) ColorModel 

``` go 
func (p *NRGBA64) ColorModel() color.Model
```

#### (*NRGBA64) NRGBA64At  <- go1.4

``` go 
func (p *NRGBA64) NRGBA64At(x, y int) color.NRGBA64
```

#### (*NRGBA64) Opaque 

``` go 
func (p *NRGBA64) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*NRGBA64) PixOffset 

``` go 
func (p *NRGBA64) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*NRGBA64) RGBA64At  <- go1.17

``` go 
func (p *NRGBA64) RGBA64At(x, y int) color.RGBA64
```

#### (*NRGBA64) Set 

``` go 
func (p *NRGBA64) Set(x, y int, c color.Color)
```

#### (*NRGBA64) SetNRGBA64 

``` go 
func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64)
```

#### (*NRGBA64) SetRGBA64  <- go1.17

``` go 
func (p *NRGBA64) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*NRGBA64) SubImage 

``` go 
func (p *NRGBA64) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type NYCbCrA  <- go1.6

``` go 
type NYCbCrA struct {
	YCbCr
	A       []uint8
	AStride int
}
```

NYCbCrA is an in-memory image of non-alpha-premultiplied Y'CbCr-with-alpha colors. A and AStride are analogous to the Y and YStride fields of the embedded YCbCr.

#### func NewNYCbCrA  <- go1.6

``` go 
func NewNYCbCrA(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *NYCbCrA
```

NewNYCbCrA returns a new NYCbCrA image with the given bounds and subsample ratio.

#### (*NYCbCrA) AOffset  <- go1.6

``` go 
func (p *NYCbCrA) AOffset(x, y int) int
```

AOffset returns the index of the first element of A that corresponds to the pixel at (x, y).

#### (*NYCbCrA) At  <- go1.6

``` go 
func (p *NYCbCrA) At(x, y int) color.Color
```

#### (*NYCbCrA) ColorModel  <- go1.6

``` go 
func (p *NYCbCrA) ColorModel() color.Model
```

#### (*NYCbCrA) NYCbCrAAt  <- go1.6

``` go 
func (p *NYCbCrA) NYCbCrAAt(x, y int) color.NYCbCrA
```

#### (*NYCbCrA) Opaque  <- go1.6

``` go 
func (p *NYCbCrA) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*NYCbCrA) RGBA64At  <- go1.17

``` go 
func (p *NYCbCrA) RGBA64At(x, y int) color.RGBA64
```

#### (*NYCbCrA) SubImage  <- go1.6

``` go 
func (p *NYCbCrA) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type Paletted 

``` go 
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

#### func NewPaletted 

``` go 
func NewPaletted(r Rectangle, p color.Palette) *Paletted
```

NewPaletted returns a new Paletted image with the given width, height and palette.

#### (*Paletted) At 

``` go 
func (p *Paletted) At(x, y int) color.Color
```

#### (*Paletted) Bounds 

``` go 
func (p *Paletted) Bounds() Rectangle
```

#### (*Paletted) ColorIndexAt 

``` go 
func (p *Paletted) ColorIndexAt(x, y int) uint8
```

#### (*Paletted) ColorModel 

``` go 
func (p *Paletted) ColorModel() color.Model
```

#### (*Paletted) Opaque 

``` go 
func (p *Paletted) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Paletted) PixOffset 

``` go 
func (p *Paletted) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*Paletted) RGBA64At  <- go1.17

``` go 
func (p *Paletted) RGBA64At(x, y int) color.RGBA64
```

#### (*Paletted) Set 

``` go 
func (p *Paletted) Set(x, y int, c color.Color)
```

#### (*Paletted) SetColorIndex 

``` go 
func (p *Paletted) SetColorIndex(x, y int, index uint8)
```

#### (*Paletted) SetRGBA64  <- go1.17

``` go 
func (p *Paletted) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Paletted) SubImage 

``` go 
func (p *Paletted) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type PalettedImage 

``` go 
type PalettedImage interface {
	// ColorIndexAt returns the palette index of the pixel at (x, y).
	ColorIndexAt(x, y int) uint8
	Image
}
```

PalettedImage is an image whose colors may come from a limited palette. If m is a PalettedImage and m.ColorModel() returns a color.Palette p, then m.At(x, y) should be equivalent to p[m.ColorIndexAt(x, y)]. If m's color model is not a color.Palette, then ColorIndexAt's behavior is undefined.

### type Point 

``` go 
type Point struct {
	X, Y int
}
```

A Point is an X, Y coordinate pair. The axes increase right and down.

``` go 
var ZP Point
```

ZP is the zero Point.

Deprecated: Use a literal image.Point{} instead.

#### func Pt 

``` go 
func Pt(X, Y int) Point
```

Pt is shorthand for Point{X, Y}.

#### (Point) Add 

``` go 
func (p Point) Add(q Point) Point
```

Add returns the vector p+q.

#### (Point) Div 

``` go 
func (p Point) Div(k int) Point
```

Div returns the vector p/k.

#### (Point) Eq 

``` go 
func (p Point) Eq(q Point) bool
```

Eq reports whether p and q are equal.

#### (Point) In 

``` go 
func (p Point) In(r Rectangle) bool
```

In reports whether p is in r.

#### (Point) Mod 

``` go 
func (p Point) Mod(r Rectangle) Point
```

Mod returns the point q in r such that p.X-q.X is a multiple of r's width and p.Y-q.Y is a multiple of r's height.

#### (Point) Mul 

``` go 
func (p Point) Mul(k int) Point
```

Mul returns the vector p*k.

#### (Point) String 

``` go 
func (p Point) String() string
```

String returns a string representation of p like "(3,4)".

#### (Point) Sub 

``` go 
func (p Point) Sub(q Point) Point
```

Sub returns the vector p-q.

### type RGBA 

``` go 
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

#### func NewRGBA 

``` go 
func NewRGBA(r Rectangle) *RGBA
```

NewRGBA returns a new RGBA image with the given bounds.

#### (*RGBA) At 

``` go 
func (p *RGBA) At(x, y int) color.Color
```

#### (*RGBA) Bounds 

``` go 
func (p *RGBA) Bounds() Rectangle
```

#### (*RGBA) ColorModel 

``` go 
func (p *RGBA) ColorModel() color.Model
```

#### (*RGBA) Opaque 

``` go 
func (p *RGBA) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*RGBA) PixOffset 

``` go 
func (p *RGBA) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*RGBA) RGBA64At  <- go1.17

``` go 
func (p *RGBA) RGBA64At(x, y int) color.RGBA64
```

#### (*RGBA) RGBAAt  <- go1.4

``` go 
func (p *RGBA) RGBAAt(x, y int) color.RGBA
```

#### (*RGBA) Set 

``` go 
func (p *RGBA) Set(x, y int, c color.Color)
```

#### (*RGBA) SetRGBA 

``` go 
func (p *RGBA) SetRGBA(x, y int, c color.RGBA)
```

#### (*RGBA) SetRGBA64  <- go1.17

``` go 
func (p *RGBA) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*RGBA) SubImage 

``` go 
func (p *RGBA) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type RGBA64 

``` go 
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

#### func NewRGBA64 

``` go 
func NewRGBA64(r Rectangle) *RGBA64
```

NewRGBA64 returns a new RGBA64 image with the given bounds.

#### (*RGBA64) At 

``` go 
func (p *RGBA64) At(x, y int) color.Color
```

#### (*RGBA64) Bounds 

``` go 
func (p *RGBA64) Bounds() Rectangle
```

#### (*RGBA64) ColorModel 

``` go 
func (p *RGBA64) ColorModel() color.Model
```

#### (*RGBA64) Opaque 

``` go 
func (p *RGBA64) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*RGBA64) PixOffset 

``` go 
func (p *RGBA64) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

#### (*RGBA64) RGBA64At  <- go1.4

``` go 
func (p *RGBA64) RGBA64At(x, y int) color.RGBA64
```

#### (*RGBA64) Set 

``` go 
func (p *RGBA64) Set(x, y int, c color.Color)
```

#### (*RGBA64) SetRGBA64 

``` go 
func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*RGBA64) SubImage 

``` go 
func (p *RGBA64) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

### type RGBA64Image  <- go1.17

``` go 
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

### type Rectangle 

``` go 
type Rectangle struct {
	Min, Max Point
}
```

A Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y. It is well-formed if Min.X <= Max.X and likewise for Y. Points are always well-formed. A rectangle's methods always return well-formed outputs for well-formed inputs.

A Rectangle is also an Image whose bounds are the rectangle itself. At returns color.Opaque for points in the rectangle and color.Transparent otherwise.

``` go 
var ZR Rectangle
```

ZR is the zero Rectangle.

Deprecated: Use a literal image.Rectangle{} instead.

#### func Rect 

``` go 
func Rect(x0, y0, x1, y1 int) Rectangle
```

Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}. The returned rectangle has minimum and maximum coordinates swapped if necessary so that it is well-formed.

#### (Rectangle) Add 

``` go 
func (r Rectangle) Add(p Point) Rectangle
```

Add returns the rectangle r translated by p.

#### (Rectangle) At  <- go1.5

``` go 
func (r Rectangle) At(x, y int) color.Color
```

At implements the Image interface.

#### (Rectangle) Bounds  <- go1.5

``` go 
func (r Rectangle) Bounds() Rectangle
```

Bounds implements the Image interface.

#### (Rectangle) Canon 

``` go 
func (r Rectangle) Canon() Rectangle
```

Canon returns the canonical version of r. The returned rectangle has minimum and maximum coordinates swapped if necessary so that it is well-formed.

#### (Rectangle) ColorModel  <- go1.5

``` go 
func (r Rectangle) ColorModel() color.Model
```

ColorModel implements the Image interface.

#### (Rectangle) Dx 

``` go 
func (r Rectangle) Dx() int
```

Dx returns r's width.

#### (Rectangle) Dy 

``` go 
func (r Rectangle) Dy() int
```

Dy returns r's height.

#### (Rectangle) Empty 

``` go 
func (r Rectangle) Empty() bool
```

Empty reports whether the rectangle contains no points.

#### (Rectangle) Eq 

``` go 
func (r Rectangle) Eq(s Rectangle) bool
```

Eq reports whether r and s contain the same set of points. All empty rectangles are considered equal.

#### (Rectangle) In 

``` go 
func (r Rectangle) In(s Rectangle) bool
```

In reports whether every point in r is in s.

#### (Rectangle) Inset 

``` go 
func (r Rectangle) Inset(n int) Rectangle
```

Inset returns the rectangle r inset by n, which may be negative. If either of r's dimensions is less than 2*n then an empty rectangle near the center of r will be returned.

#### (Rectangle) Intersect 

``` go 
func (r Rectangle) Intersect(s Rectangle) Rectangle
```

Intersect returns the largest rectangle contained by both r and s. If the two rectangles do not overlap then the zero rectangle will be returned.

#### (Rectangle) Overlaps 

``` go 
func (r Rectangle) Overlaps(s Rectangle) bool
```

Overlaps reports whether r and s have a non-empty intersection.

#### (Rectangle) RGBA64At  <- go1.17

``` go 
func (r Rectangle) RGBA64At(x, y int) color.RGBA64
```

RGBA64At implements the RGBA64Image interface.

#### (Rectangle) Size 

``` go 
func (r Rectangle) Size() Point
```

Size returns r's width and height.

#### (Rectangle) String 

``` go 
func (r Rectangle) String() string
```

String returns a string representation of r like "(3,4)-(6,5)".

#### (Rectangle) Sub 

``` go 
func (r Rectangle) Sub(p Point) Rectangle
```

Sub returns the rectangle r translated by -p.

#### (Rectangle) Union 

``` go 
func (r Rectangle) Union(s Rectangle) Rectangle
```

Union returns the smallest rectangle that contains both r and s.

### type Uniform 

``` go 
type Uniform struct {
	C color.Color
}
```

Uniform is an infinite-sized Image of uniform color. It implements the color.Color, color.Model, and Image interfaces.

#### func NewUniform 

``` go 
func NewUniform(c color.Color) *Uniform
```

NewUniform returns a new Uniform image of the given color.

#### (*Uniform) At 

``` go 
func (c *Uniform) At(x, y int) color.Color
```

#### (*Uniform) Bounds 

``` go 
func (c *Uniform) Bounds() Rectangle
```

#### (*Uniform) ColorModel 

``` go 
func (c *Uniform) ColorModel() color.Model
```

#### (*Uniform) Convert 

``` go 
func (c *Uniform) Convert(color.Color) color.Color
```

#### (*Uniform) Opaque 

``` go 
func (c *Uniform) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

#### (*Uniform) RGBA 

``` go 
func (c *Uniform) RGBA() (r, g, b, a uint32)
```

#### (*Uniform) RGBA64At  <- go1.17

``` go 
func (c *Uniform) RGBA64At(x, y int) color.RGBA64
```

### type YCbCr 

``` go 
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

#### func NewYCbCr 

``` go 
func NewYCbCr(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *YCbCr
```

NewYCbCr returns a new YCbCr image with the given bounds and subsample ratio.

#### (*YCbCr) At 

``` go 
func (p *YCbCr) At(x, y int) color.Color
```

#### (*YCbCr) Bounds 

``` go 
func (p *YCbCr) Bounds() Rectangle
```

#### (*YCbCr) COffset 

``` go 
func (p *YCbCr) COffset(x, y int) int
```

COffset returns the index of the first element of Cb or Cr that corresponds to the pixel at (x, y).

#### (*YCbCr) ColorModel 

``` go 
func (p *YCbCr) ColorModel() color.Model
```

#### (*YCbCr) Opaque 

``` go 
func (p *YCbCr) Opaque() bool
```

#### (*YCbCr) RGBA64At  <- go1.17

``` go 
func (p *YCbCr) RGBA64At(x, y int) color.RGBA64
```

#### (*YCbCr) SubImage 

``` go 
func (p *YCbCr) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

#### (*YCbCr) YCbCrAt  <- go1.4

``` go 
func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr
```

#### (*YCbCr) YOffset 

``` go 
func (p *YCbCr) YOffset(x, y int) int
```

YOffset returns the index of the first element of Y that corresponds to the pixel at (x, y).

### type YCbCrSubsampleRatio 

``` go 
type YCbCrSubsampleRatio int
```

YCbCrSubsampleRatio is the chroma subsample ratio used in a YCbCr image.

``` go 
const (
	YCbCrSubsampleRatio444 YCbCrSubsampleRatio = iota
	YCbCrSubsampleRatio422
	YCbCrSubsampleRatio420
	YCbCrSubsampleRatio440
	YCbCrSubsampleRatio411
	YCbCrSubsampleRatio410
)
```

#### (YCbCrSubsampleRatio) String 

``` go 
func (s YCbCrSubsampleRatio) String() string
```