+++
title = "gif"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/image/gif@go1.23.0](https://pkg.go.dev/image/gif@go1.23.0)

Package gif implements a GIF image decoder and encoder.

​	 gif 包实现了一个 GIF 图像解码器和编码器。

The GIF specification is at https://www.w3.org/Graphics/GIF/spec-gif89a.txt.

​	GIF 规范位于 [https://www.w3.org/Graphics/GIF/spec-gif89a.txt](https://www.w3.org/Graphics/GIF/spec-gif89a.txt)。

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

​	处置方法。

## 变量

This section is empty.

## 函数

### func Decode 

``` go 
func Decode(r io.Reader) (image.Image, error)
```

Decode reads a GIF image from r and returns the first embedded image as an image.Image.

​	Decode 从 r 读取 GIF 图像，并将第一个嵌入图像作为 image.Image 返回。

### func DecodeConfig

```go
func DecodeConfig(r io.Reader) (image.Config, error)
```

DecodeConfig returns the global color model and dimensions of a GIF image without decoding the entire image.

​	DecodeConfig 返回 GIF 图像的全局颜色模型和尺寸，而无需解码整个图像。

### func Encode <- go1.2

```go
func Encode(w io.Writer, m image.Image, o *Options) error
```

Encode writes the Image m to w in GIF format.

​	Encode 以 GIF 格式将图像 m 写入 w。

### func EncodeAll <- go1.2

```go
func EncodeAll(w io.Writer, g *GIF) error
```

EncodeAll writes the images in g to w in GIF format with the given loop count and delay between frames.

​	EncodeAll 以 GIF 格式将 g 中的图像写入 w，并具有给定的循环计数和帧之间的延迟。

## 类型

### type GIF

```go
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

​	GIF 表示存储在 GIF 文件中的可能多个图像。

#### func DecodeAll

```go
func DecodeAll(r io.Reader) (*GIF, error)
```

DecodeAll reads a GIF image from r and returns the sequential frames and timing information.

​	DecodeAll 从 r 读取 GIF 图像，并返回顺序帧和计时信息。

### type Options <- go1.2

```go
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

​	Options 是编码参数。