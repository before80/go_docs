+++
title = "gif"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/image/gif@go1.21.3](https://pkg.go.dev/image/gif@go1.21.3)

Package gif implements a GIF image decoder and encoder.

The GIF specification is at https://www.w3.org/Graphics/GIF/spec-gif89a.txt.

## 常量 

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/image/gif/reader.go;l=45)

``` go 
const (
	DisposalNone       = 0x01
	DisposalBackground = 0x02
	DisposalPrevious   = 0x03
)
```

Disposal Methods.

## 变量

This section is empty.

## 函数

### func Decode 

``` go 
func Decode(r io.Reader) (image.Image, error)
```

Decode reads a GIF image from r and returns the first embedded image as an image.Image.

### func DecodeConfig 

``` go 
func DecodeConfig(r io.Reader) (image.Config, error)
```

DecodeConfig returns the global color model and dimensions of a GIF image without decoding the entire image.

### func Encode  <- go1.2

``` go 
func Encode(w io.Writer, m image.Image, o *Options) error
```

Encode writes the Image m to w in GIF format.

### func EncodeAll  <- go1.2

``` go 
func EncodeAll(w io.Writer, g *GIF) error
```

EncodeAll writes the images in g to w in GIF format with the given loop count and delay between frames.

## 类型

### type GIF 

``` go 
type GIF struct {
	Image []*image.Paletted // The successive images.
	Delay []int             // The successive delay times, one per frame, in 100ths of a second.
	// LoopCount controls the number of times an animation will be
	// restarted during display.
	// A LoopCount of 0 means to loop forever.
	// A LoopCount of -1 means to show each frame only once.
	// Otherwise, the animation is looped LoopCount+1 times.
	LoopCount int
	// Disposal is the successive disposal methods, one per frame. For
	// backwards compatibility, a nil Disposal is valid to pass to EncodeAll,
	// and implies that each frame's disposal method is 0 (no disposal
	// specified).
	Disposal []byte
	// Config is the global color table (palette), width and height. A nil or
	// empty-color.Palette Config.ColorModel means that each frame has its own
	// color table and there is no global color table. Each frame's bounds must
	// be within the rectangle defined by the two points (0, 0) and
	// (Config.Width, Config.Height).
	//
	// For backwards compatibility, a zero-valued Config is valid to pass to
	// EncodeAll, and implies that the overall GIF's width and height equals
	// the first frame's bounds' Rectangle.Max point.
	Config image.Config
	// BackgroundIndex is the background index in the global color table, for
	// use with the DisposalBackground disposal method.
	BackgroundIndex byte
}
```

GIF represents the possibly multiple images stored in a GIF file.

#### func DecodeAll 

``` go 
func DecodeAll(r io.Reader) (*GIF, error)
```

DecodeAll reads a GIF image from r and returns the sequential frames and timing information.

### type Options  <- go1.2

``` go 
type Options struct {
	// NumColors is the maximum number of colors used in the image.
	// It ranges from 1 to 256.
	NumColors int

	// Quantizer is used to produce a palette with size NumColors.
	// palette.Plan9 is used in place of a nil Quantizer.
	Quantizer draw.Quantizer

	// Drawer is used to convert the source image to the desired palette.
	// draw.FloydSteinberg is used in place of a nil Drawer.
	Drawer draw.Drawer
}
```

Options are the encoding parameters.