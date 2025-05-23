+++
title = "image"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/image@go1.24.2](https://pkg.go.dev/image@go1.24.2)

Package image implements a basic 2-D image library.

​	image 包实现了一个基本的 2-D 图像库。

The fundamental interface is called Image. An Image contains colors, which are described in the image/color package.

​	基本接口称为 Image。Image 包含颜色，这些颜色在 image/color 包中进行了描述。

Values of the Image interface are created either by calling functions such as NewRGBA and NewPaletted, or by calling Decode on an io.Reader containing image data in a format such as GIF, JPEG or PNG. Decoding any particular image format requires the prior registration of a decoder function. Registration is typically automatic as a side effect of initializing that format’s package so that, to decode a PNG image, it suffices to have

​	Image 接口的值通过调用 NewRGBA 和 NewPaletted 等函数或通过对包含 GIF、JPEG 或 PNG 等格式的图像数据的 io.Reader 调用 Decode 来创建。解码任何特定图像格式都需要预先注册解码器函数。注册通常是初始化该格式的包的副作用的自动结果，因此，要解码 PNG 图像，只需在程序的主包中使用

```go
import _ "image/png"
```

in a program’s main package. The _ means to import a package purely for its initialization side effects.

。_ 表示纯粹为了其初始化副作用而导入一个包。

See “The Go image package” for more details: https://golang.org/doc/articles/image_package.html

​	有关更多详细信息，请参阅“Go image 包”：[https://golang.org/doc/articles/image_package.html](https://golang.org/doc/articles/image_package.html)

## Example

```go
package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"strings"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// _ "image/png"
	_ "image/jpeg"
)

func main() {
	// Decode the JPEG data. If reading from file, create a reader with
	//
	// reader, err := os.Open("testdata/video-001.q50.420.jpeg")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer reader.Close()
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	// Calculate a 16-bin histogram for m's red, green, blue and alpha components.
	//
	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.
	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 12 reduces this to the range [0, 15].
			histogram[r>>12][0]++
			histogram[g>>12][1]++
			histogram[b>>12][2]++
			histogram[a>>12][3]++
		}
	}

	// Print the results.
	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	for i, x := range histogram {
		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	}
}

const data = `
/9j/4AAQSkZJRgABAQIAHAAcAAD/2wBDABALDA4MChAODQ4SERATGCgaGBYWGDEjJR0oOjM9PDkzODdA
SFxOQERXRTc4UG1RV19iZ2hnPk1xeXBkeFxlZ2P/2wBDARESEhgVGC8aGi9jQjhCY2NjY2NjY2NjY2Nj
Y2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2P/wAARCABnAJYDASIAAhEBAxEB/8QA
HwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIh
MUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVW
V1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXG
x8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQF
BgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAV
YnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOE
hYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq
8vP09fb3+Pn6/9oADAMBAAIRAxEAPwDlwKMD0pwzSiuK57QzGDxS7D6in8Y5ximnAPUfSlcq4m3ilUYp
2OKXHvRcVxnTtS7c07HNFK4DQPakC4PNOA+tOx70XAjK/So5gBGP94fzqfvUVx/qxx/EP51UXqRP4WSE
cmgjilP3jSEZqS0IO/NGDnpUiocDg/McDjvV6HTPOdVWYgsM5KcfzzQ2JySM2jp6VYu7SWzmMUwG4cgj
kMPUVBjjtTGtRu0Zopw+lFFxhinrGzuqqMsxAA9yaXFSRv5cqSEcIwYj6GpuZ30O30fSLKzhUpbpNMv3
5XGTn29BV28jt7pPLuIVljPBBFVreYx+VbqAjycgt3x14zRcNOxGyVFHQkIc/wA61exyKLbuzjdZ046d
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
`
Output:

bin               red  green   blue  alpha
0x0000-0x0fff:    364    790   7242      0
0x1000-0x1fff:    645   2967   1039      0
0x2000-0x2fff:   1072   2299    979      0
0x3000-0x3fff:    820   2266    980      0
0x4000-0x4fff:    537   1305    541      0
0x5000-0x5fff:    319    962    261      0
0x6000-0x6fff:    322    375    177      0
0x7000-0x7fff:    601    279    214      0
0x8000-0x8fff:   3478    227    273      0
0x9000-0x9fff:   2260    234    329      0
0xa000-0xafff:    921    282    373      0
0xb000-0xbfff:    321    335    397      0
0xc000-0xcfff:    229    388    298      0
0xd000-0xdfff:    260    414    277      0
0xe000-0xefff:    516    428    298      0
0xf000-0xffff:   2785   1899   1772  15450
```

## Example (DecodeConfig)

```go
package main

import (
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"strings"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// _ "image/png"
	_ "image/jpeg"
)

func main() {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	config, format, err := image.DecodeConfig(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Width:", config.Width, "Height:", config.Height, "Format:", format)
}

const data = `
/9j/4AAQSkZJRgABAQIAHAAcAAD/2wBDABALDA4MChAODQ4SERATGCgaGBYWGDEjJR0oOjM9PDkzODdA
SFxOQERXRTc4UG1RV19iZ2hnPk1xeXBkeFxlZ2P/2wBDARESEhgVGC8aGi9jQjhCY2NjY2NjY2NjY2Nj
Y2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2NjY2P/wAARCABnAJYDASIAAhEBAxEB/8QA
HwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIh
MUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVW
V1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXG
x8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQF
BgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAV
YnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOE
hYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq
8vP09fb3+Pn6/9oADAMBAAIRAxEAPwDlwKMD0pwzSiuK57QzGDxS7D6in8Y5ximnAPUfSlcq4m3ilUYp
2OKXHvRcVxnTtS7c07HNFK4DQPakC4PNOA+tOx70XAjK/So5gBGP94fzqfvUVx/qxx/EP51UXqRP4WSE
cmgjilP3jSEZqS0IO/NGDnpUiocDg/McDjvV6HTPOdVWYgsM5KcfzzQ2JySM2jp6VYu7SWzmMUwG4cgj
kMPUVBjjtTGtRu0Zopw+lFFxhinrGzuqqMsxAA9yaXFSRv5cqSEcIwYj6GpuZ30O30fSLKzhUpbpNMv3
5XGTn29BV28jt7pPLuIVljPBBFVreYx+VbqAjycgt3x14zRcNOxGyVFHQkIc/wA61exyKLbuzjdZ046d
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
`
Output:
```

## 常量

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/image/names.go;l=11)

```go
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

```go
var ErrFormat = errors.New("image: unknown format")
```

ErrFormat indicates that decoding encountered an unknown format.

​	ErrFormat 指示解码遇到未知格式。

## 函数

### func RegisterFormat

```go
func RegisterFormat(name, magic string, decode func(io.Reader) (Image, error), decodeConfig func(io.Reader) (Config, error))
```

RegisterFormat registers an image format for use by Decode. Name is the name of the format, like “jpeg” or “png”. Magic is the magic prefix that identifies the format’s encoding. The magic string can contain “?” wildcards that each match any one byte. Decode is the function that decodes the encoded image. DecodeConfig is the function that decodes just its configuration.

​	RegisterFormat 为 Decode 注册一种图像格式。Name 是格式的名称，如“jpeg”或“png”。Magic 是标识格式编码的魔术前缀。魔术字符串可以包含“?”通配符，每个通配符匹配任何一个字节。Decode 是解码编码图像的函数。DecodeConfig 是仅解码其配置的函数。

## 类型

### type Alpha

```go
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

​	Alpha 是一个内存中图像，其 At 方法返回 color.Alpha 值。

#### func NewAlpha

```go
func NewAlpha(r Rectangle) *Alpha
```

NewAlpha returns a new Alpha image with the given bounds.

​	NewAlpha 返回一个具有给定边界的 Alpha 新图像。

#### (*Alpha) AlphaAt <- go1.4

```go
func (p *Alpha) AlphaAt(x, y int) color.Alpha
```

#### (*Alpha) At

```go
func (p *Alpha) At(x, y int) color.Color
```

#### (*Alpha) Bounds

```go
func (p *Alpha) Bounds() Rectangle
```

#### (*Alpha) ColorModel

```go
func (p *Alpha) ColorModel() color.Model
```

#### (*Alpha) Opaque

```go
func (p *Alpha) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*Alpha) PixOffset

```go
func (p *Alpha) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

​	PixOffset 返回对应于 (x, y) 处像素的 Pix 的第一个元素的索引。

#### (*Alpha) RGBA64At <- go1.17

```go
func (p *Alpha) RGBA64At(x, y int) color.RGBA64
```

#### (*Alpha) Set

```go
func (p *Alpha) Set(x, y int, c color.Color)
```

#### (*Alpha) SetAlpha

```go
func (p *Alpha) SetAlpha(x, y int, c color.Alpha)
```

#### (*Alpha) SetRGBA64 <- go1.17

```go
func (p *Alpha) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Alpha) SubImage

```go
func (p *Alpha) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回值与原始图像共享像素。

### type Alpha16

```go
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

​	Alpha16 是一个内存中图像，其 At 方法返回 color.Alpha16 值。

#### func NewAlpha16

```go
func NewAlpha16(r Rectangle) *Alpha16
```

NewAlpha16 returns a new Alpha16 image with the given bounds.

​	NewAlpha16 返回一个具有给定边界的 Alpha16 新图像。

#### (*Alpha16) Alpha16At <- go1.4

```go
func (p *Alpha16) Alpha16At(x, y int) color.Alpha16
```

#### (*Alpha16) At

```go
func (p *Alpha16) At(x, y int) color.Color
```

#### (*Alpha16) Bounds

```go
func (p *Alpha16) Bounds() Rectangle
```

#### (*Alpha16) ColorModel

```go
func (p *Alpha16) ColorModel() color.Model
```

#### (*Alpha16) Opaque

```go
func (p *Alpha16) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*Alpha16) PixOffset

```go
func (p *Alpha16) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

​	PixOffset 返回对应于 (x, y) 处像素的 Pix 的第一个元素的索引。

#### (*Alpha16) RGBA64At <- go1.17

```go
func (p *Alpha16) RGBA64At(x, y int) color.RGBA64
```

#### (*Alpha16) Set

```go
func (p *Alpha16) Set(x, y int, c color.Color)
```

#### (*Alpha16) SetAlpha16

```go
func (p *Alpha16) SetAlpha16(x, y int, c color.Alpha16)
```

#### (*Alpha16) SetRGBA64 <- go1.17

```go
func (p *Alpha16) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Alpha16) SubImage

```go
func (p *Alpha16) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回值与原始图像共享像素。

### type CMYK <- go1.5

```go
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

​	CMYK 是一个内存中的图像，其 At 方法返回 color.CMYK 值。

#### func NewCMYK <- go1.5

```go
func NewCMYK(r Rectangle) *CMYK
```

NewCMYK returns a new CMYK image with the given bounds.

​	NewCMYK 返回一个具有给定边界的新的 CMYK 图像。

#### (*CMYK) At <- go1.5

```go
func (p *CMYK) At(x, y int) color.Color
```

#### (*CMYK) Bounds <- go1.5

```go
func (p *CMYK) Bounds() Rectangle
```

#### (*CMYK) CMYKAt <- go1.5

```go
func (p *CMYK) CMYKAt(x, y int) color.CMYK
```

#### (*CMYK) ColorModel <- go1.5

```go
func (p *CMYK) ColorModel() color.Model
```

#### (*CMYK) Opaque <- go1.5

```go
func (p *CMYK) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*CMYK) PixOffset <- go1.5

```go
func (p *CMYK) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

​	PixOffset 返回对应于 (x, y) 处像素的 Pix 的第一个元素的索引。

#### (*CMYK) RGBA64At <- go1.17

```go
func (p *CMYK) RGBA64At(x, y int) color.RGBA64
```

#### (*CMYK) Set <- go1.5

```go
func (p *CMYK) Set(x, y int, c color.Color)
```

#### (*CMYK) SetCMYK <- go1.5

```go
func (p *CMYK) SetCMYK(x, y int, c color.CMYK)
```

#### (*CMYK) SetRGBA64 <- go1.17

```go
func (p *CMYK) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*CMYK) SubImage <- go1.5

```go
func (p *CMYK) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回值与原始图像共享像素。

### type Config

```go
type Config struct {
	ColorModel    color.Model
	Width, Height int
}
```

Config holds an image’s color model and dimensions.

​	Config 保存图像的色彩模型和尺寸。

#### func DecodeConfig

```go
func DecodeConfig(r io.Reader) (Config, string, error)
```

DecodeConfig decodes the color model and dimensions of an image that has been encoded in a registered format. The string returned is the format name used during format registration. Format registration is typically done by an init function in the codec-specific package.

​	DecodeConfig 解码已编码为已注册格式的图像的色彩模型和尺寸。返回的字符串是在格式注册期间使用的格式名称。格式注册通常由编解码器特定包中的 init 函数完成。

### type Gray

```go
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

​	Gray 是一个内存中的图像，其 At 方法返回 color.Gray 值。

#### func NewGray

```go
func NewGray(r Rectangle) *Gray
```

NewGray returns a new Gray image with the given bounds.

​	NewGray 返回一个具有给定边界的新的 Gray 图像。

#### (*Gray) At

```go
func (p *Gray) At(x, y int) color.Color
```

#### (*Gray) Bounds 

```go
func (p *Gray) Bounds() Rectangle
```

#### (*Gray) ColorModel

```go
func (p *Gray) ColorModel() color.Model
```

#### (*Gray) GrayAt <- go1.4

```go
func (p *Gray) GrayAt(x, y int) color.Gray
```

#### (*Gray) Opaque

```go
func (p *Gray) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*Gray) PixOffset

```go
func (p *Gray) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

​	PixOffset 返回对应于 (x, y) 处像素的 Pix 的第一个元素的索引。

#### (*Gray) RGBA64At <- go1.17

```go
func (p *Gray) RGBA64At(x, y int) color.RGBA64
```

#### (*Gray) Set

```go
func (p *Gray) Set(x, y int, c color.Color)
```

#### (*Gray) SetGray

```go
func (p *Gray) SetGray(x, y int, c color.Gray)
```

#### (*Gray) SetRGBA64 <- go1.17

```go
func (p *Gray) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Gray) SubImage

```go
func (p *Gray) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回值与原始图像共享像素。

### type Gray16

```go
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

​	Gray16 是一个内存中图像，其 At 方法返回 color.Gray16 值。

#### func NewGray16

```go
func NewGray16(r Rectangle) *Gray16
```

NewGray16 returns a new Gray16 image with the given bounds.

​	NewGray16 返回具有给定边界的新的 Gray16 图像。 （*Gray16）边界

#### (*Gray16) At

```go
func (p *Gray16) At(x, y int) color.Color
```

#### (*Gray16) Bounds 

```go
func (p *Gray16) Bounds() Rectangle
```

#### (*Gray16) ColorModel

```go
func (p *Gray16) ColorModel() color.Model
```

#### (*Gray16) Gray16At <- go1.4

```go
func (p *Gray16) Gray16At(x, y int) color.Gray16
```

#### (*Gray16) Opaque

```go
func (p *Gray16) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*Gray16) PixOffset

```go
func (p *Gray16) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

​	PixOffset 返回对应于 (x, y) 处像素的 Pix 的第一个元素的索引。

#### (*Gray16) RGBA64At <- go1.17

```go
func (p *Gray16) RGBA64At(x, y int) color.RGBA64
```

#### (*Gray16) Set

```go
func (p *Gray16) Set(x, y int, c color.Color)
```

#### (*Gray16) SetGray16

```go
func (p *Gray16) SetGray16(x, y int, c color.Gray16)
```

#### (*Gray16) SetRGBA64 <- go1.17

```go
func (p *Gray16) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Gray16) SubImage

```go
func (p *Gray16) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回值与原始图像共享像素。

### type Image

```go
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

​	Image 是一个有限的矩形颜色网格。颜色值取自颜色模型。

#### func Decode

```go
func Decode(r io.Reader) (Image, string, error)
```

Decode decodes an image that has been encoded in a registered format. The string returned is the format name used during format registration. Format registration is typically done by an init function in the codec- specific package.

​	Decode 解码已按注册格式编码的图像。返回的字符串是格式注册期间使用的格式名称。格式注册通常由编解码器特定包中的 init 函数完成。

### type NRGBA

```go
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

​	NRGBA 是一个内存中图像，其 At 方法返回 color.NRGBA 值。

#### func NewNRGBA

```go
func NewNRGBA(r Rectangle) *NRGBA
```

NewNRGBA returns a new NRGBA image with the given bounds.

​	NewNRGBA 返回具有给定边界的全新 NRGBA 图像。

#### (*NRGBA) At

```go
func (p *NRGBA) At(x, y int) color.Color
```

#### (*NRGBA) Bounds

```go
func (p *NRGBA) Bounds() Rectangle
```

#### (*NRGBA) ColorModel

```go
func (p *NRGBA) ColorModel() color.Model
```

#### (*NRGBA) NRGBAAt <- go1.4

```go
func (p *NRGBA) NRGBAAt(x, y int) color.NRGBA
```

#### (*NRGBA) Opaque

```go
func (p *NRGBA) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*NRGBA) PixOffset

```go
func (p *NRGBA) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

​	PixOffset 返回对应于 (x, y) 处像素的 Pix 的第一个元素的索引。

#### (*NRGBA) RGBA64At <- go1.17

```go
func (p *NRGBA) RGBA64At(x, y int) color.RGBA64
```

#### (*NRGBA) Set

```go
func (p *NRGBA) Set(x, y int, c color.Color)
```

#### (*NRGBA) SetNRGBA

```go
func (p *NRGBA) SetNRGBA(x, y int, c color.NRGBA)
```

#### (*NRGBA) SetRGBA64 <- go1.17

```go
func (p *NRGBA) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*NRGBA) SubImage

```go
func (p *NRGBA) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回的值与原始图像共享像素。

### type NRGBA64

```go
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

​	NRGBA64 是一个内存中图像，其 At 方法返回 color.NRGBA64 值。

#### func NewNRGBA64

```go
func NewNRGBA64(r Rectangle) *NRGBA64
```

NewNRGBA64 returns a new NRGBA64 image with the given bounds.

​	NewNRGBA64 返回具有给定边界的全新 NRGBA64 图像。

#### (*NRGBA64) At

```go
func (p *NRGBA64) At(x, y int) color.Color
```

#### (*NRGBA64) Bounds

```go
func (p *NRGBA64) Bounds() Rectangle
```

#### (*NRGBA64) ColorModel

```go
func (p *NRGBA64) ColorModel() color.Model
```

#### (*NRGBA64) NRGBA64At <- go1.4

```go
func (p *NRGBA64) NRGBA64At(x, y int) color.NRGBA64
```

#### (*NRGBA64) Opaque

```go
func (p *NRGBA64) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*NRGBA64) PixOffset

```go
func (p *NRGBA64) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

​	PixOffset 返回对应于 (x, y) 处的像素的 Pix 的第一个元素的索引。

#### (*NRGBA64) RGBA64At <- go1.17

```go
func (p *NRGBA64) RGBA64At(x, y int) color.RGBA64
```

#### (*NRGBA64) Set

```go
func (p *NRGBA64) Set(x, y int, c color.Color)
```

#### (*NRGBA64) SetNRGBA64

```go
func (p *NRGBA64) SetNRGBA64(x, y int, c color.NRGBA64)
```

#### (*NRGBA64) SetRGBA64 <- go1.17

```go
func (p *NRGBA64) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*NRGBA64) SubImage

```go
func (p *NRGBA64) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回的值与原始图像共享像素。

### type NYCbCrA <- go1.6

```go
type NYCbCrA struct {
	YCbCr
	A       []uint8
	AStride int
}
```

NYCbCrA is an in-memory image of non-alpha-premultiplied Y’CbCr-with-alpha colors. A and AStride are analogous to the Y and YStride fields of the embedded YCbCr.

​	NYCbCrA 是非 alpha 预乘 Y’CbCr-with-alpha 颜色的内存中图像。A 和 AStride 与嵌入式 YCbCr 的 Y 和 YStride 字段类似。

#### func NewNYCbCrA <- go1.6

```go
func NewNYCbCrA(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *NYCbCrA
```

NewNYCbCrA returns a new NYCbCrA image with the given bounds and subsample ratio.

​	NewNYCbCrA 返回具有给定边界和子采样率的新 NYCbCrA 图像。

#### (*NYCbCrA) AOffset <- go1.6

```go
func (p *NYCbCrA) AOffset(x, y int) int
```

AOffset returns the index of the first element of A that corresponds to the pixel at (x, y).

​	AOffset 返回对应于 (x, y) 处的像素的 A 的第一个元素的索引。

#### (*NYCbCrA) At <- go1.6

```go
func (p *NYCbCrA) At(x, y int) color.Color
```

#### (*NYCbCrA) ColorModel <- go1.6

```go
func (p *NYCbCrA) ColorModel() color.Model
```

#### (*NYCbCrA) NYCbCrAAt <- go1.6

```go
func (p *NYCbCrA) NYCbCrAAt(x, y int) color.NYCbCrA
```

#### (*NYCbCrA) Opaque <- go1.6

```go
func (p *NYCbCrA) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*NYCbCrA) RGBA64At <- go1.17

```go
func (p *NYCbCrA) RGBA64At(x, y int) color.RGBA64
```

#### (*NYCbCrA) SubImage <- go1.6

```go
func (p *NYCbCrA) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回的值与原始图像共享像素。

### type Paletted

```go
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

​	Paletted 是给定调色板中 uint8 索引的内存中图像。

#### func NewPaletted

```go
func NewPaletted(r Rectangle, p color.Palette) *Paletted
```

NewPaletted returns a new Paletted image with the given width, height and palette.

​	NewPaletted 返回具有给定宽度、高度和调色板的新调色板图像。

#### (*Paletted) At

```go
func (p *Paletted) At(x, y int) color.Color
```

#### (*Paletted) Bounds

```go
func (p *Paletted) Bounds() Rectangle
```

#### (*Paletted) ColorIndexAt

```go
func (p *Paletted) ColorIndexAt(x, y int) uint8
```

#### (*Paletted) ColorModel

```go
func (p *Paletted) ColorModel() color.Model
```

#### (*Paletted) Opaque

```go
func (p *Paletted) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*Paletted) PixOffset

```go
func (p *Paletted) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

​	PixOffset 返回对应于 (x, y) 处像素的 Pix 的第一个元素的索引。

#### (*Paletted) RGBA64At <- go1.17

```go
func (p *Paletted) RGBA64At(x, y int) color.RGBA64
```

#### (*Paletted) Set

```go
func (p *Paletted) Set(x, y int, c color.Color)
```

#### (*Paletted) SetColorIndex

```go
func (p *Paletted) SetColorIndex(x, y int, index uint8)
```

#### (*Paletted) SetRGBA64 <- go1.17

```go
func (p *Paletted) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*Paletted) SubImage

```go
func (p *Paletted) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回值与原始图像共享像素。

### type PalettedImage

```go
type PalettedImage interface {
	// ColorIndexAt returns the palette index of the pixel at (x, y).
	ColorIndexAt(x, y int) uint8
	Image
}
```

PalettedImage is an image whose colors may come from a limited palette. If m is a PalettedImage and m.ColorModel() returns a color.Palette p, then m.At(x, y) should be equivalent to p[m.ColorIndexAt(x, y)]. If m’s color model is not a color.Palette, then ColorIndexAt’s behavior is undefined.

​	PalettedImage 是一个颜色可能来自有限调色板的图像。如果 m 是一个 PalettedImage，并且 m.ColorModel() 返回一个颜色。Palette p，那么 m.At(x, y) 应该等同于 p[m.ColorIndexAt(x, y)]。如果 m 的颜色模型不是一个颜色。Palette，那么 ColorIndexAt 的行为是未定义的。

### type Point

```go
type Point struct {
	X, Y int
}
```

A Point is an X, Y coordinate pair. The axes increase right and down.

​	点是一对 X、Y 坐标。轴向右和向下增加。

```go
var ZP Point
```

ZP is the zero Point.

Deprecated: Use a literal image.Point{} instead.

​	已弃用：改用文字图像。Point{}。

#### func Pt

```go
func Pt(X, Y int) Point
```

Pt is shorthand for Point{X, Y}.

​	Pt 是 Point{X, Y} 的简写。

#### (Point) Add

```go
func (p Point) Add(q Point) Point
```

Add returns the vector p+q.

​	Add 返回向量 p+q。

#### (Point) Div

```go
func (p Point) Div(k int) Point
```

Div returns the vector p/k.

#### (Point) Eq

```go
func (p Point) Eq(q Point) bool
```

Eq reports whether p and q are equal.

​	Eq 报告 p 和 q 是否相等。

#### (Point) In

```go
func (p Point) In(r Rectangle) bool
```

In reports whether p is in r.

​	In 报告 p 是否在 r 中。

#### (Point) Mod

```go
func (p Point) Mod(r Rectangle) Point
```

Mod returns the point q in r such that p.X-q.X is a multiple of r’s width and p.Y-q.Y is a multiple of r’s height.

​	Mod 返回 r 中的点 q，使得 p.X-q.X 是 r 的宽度的倍数，p.Y-q.Y 是 r 的高度的倍数。

#### (Point) Mul

```go
func (p Point) Mul(k int) Point
```

Mul returns the vector p*k.

#### (Point) String

```go
func (p Point) String() string
```

String returns a string representation of p like “(3,4)”.

​	String 返回 p 的字符串表示形式，如 “(3,4)”。

#### (Point) Sub

```go
func (p Point) Sub(q Point) Point
```

Sub returns the vector p-q.

### type RGBA

```go
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

​	RGBA 是一个内存中图像，其 At 方法返回 color.RGBA 值。

#### func NewRGBA

```go
func NewRGBA(r Rectangle) *RGBA
```

NewRGBA returns a new RGBA image with the given bounds.

​	NewRGBA 返回具有给定边界的新的 RGBA 图像。

#### (*RGBA) At

```go
func (p *RGBA) At(x, y int) color.Color
```

#### (*RGBA) Bounds

```go
func (p *RGBA) Bounds() Rectangle
```

#### (*RGBA) ColorModel

```go
func (p *RGBA) ColorModel() color.Model
```

#### (*RGBA) Opaque

```go
func (p *RGBA) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*RGBA) PixOffset

```go
func (p *RGBA) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

​	PixOffset 返回与 (x, y) 处的像素对应的 Pix 的第一个元素的索引。

#### (*RGBA) RGBA64At <- go1.17

```go
func (p *RGBA) RGBA64At(x, y int) color.RGBA64
```

#### (*RGBA) RGBAAt <- go1.4

```go
func (p *RGBA) RGBAAt(x, y int) color.RGBA
```

#### (*RGBA) Set

```go
func (p *RGBA) Set(x, y int, c color.Color)
```

#### (*RGBA) SetRGBA

```go
func (p *RGBA) SetRGBA(x, y int, c color.RGBA)
```

#### (*RGBA) SetRGBA64 <- go1.17

```go
func (p *RGBA) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*RGBA) SubImage

```go
func (p *RGBA) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回值与原始图像共享像素。

### type RGBA64

```go
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

​	RGBA64 是一个内存中的图像，其 At 方法返回 color.RGBA64 值。

#### func NewRGBA64

```go
func NewRGBA64(r Rectangle) *RGBA64
```

NewRGBA64 returns a new RGBA64 image with the given bounds.

​	NewRGBA64 返回一个具有给定边界的 RGBA64 新图像。

#### (*RGBA64) At

```go
func (p *RGBA64) At(x, y int) color.Color
```

#### (*RGBA64) Bounds

```go
func (p *RGBA64) Bounds() Rectangle
```

#### (*RGBA64) ColorModel

```go
func (p *RGBA64) ColorModel() color.Model
```

#### (*RGBA64) Opaque

```go
func (p *RGBA64) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*RGBA64) PixOffset

```go
func (p *RGBA64) PixOffset(x, y int) int
```

PixOffset returns the index of the first element of Pix that corresponds to the pixel at (x, y).

​	PixOffset 返回与 (x, y) 处的像素对应的 Pix 的第一个元素的索引。

#### (*RGBA64) RGBA64At <- go1.4

```go
func (p *RGBA64) RGBA64At(x, y int) color.RGBA64
```

#### (*RGBA64) Set

```go
func (p *RGBA64) Set(x, y int, c color.Color)
```

#### (*RGBA64) SetRGBA64

```go
func (p *RGBA64) SetRGBA64(x, y int, c color.RGBA64)
```

#### (*RGBA64) SubImage

```go
func (p *RGBA64) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回值与原始图像共享像素。

### type RGBA64Image <- go1.17

```go
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

​	RGBA64Image 是一个图像，其像素可以直接转换为 color.RGBA64。

### type Rectangle

```go
type Rectangle struct {
	Min, Max Point
}
```

A Rectangle contains the points with Min.X <= X < Max.X, Min.Y <= Y < Max.Y. It is well-formed if Min.X <= Max.X and likewise for Y. Points are always well-formed. A rectangle’s methods always return well-formed outputs for well-formed inputs.

​	矩形包含点 Min.X <= X < Max.X、Min.Y <= Y < Max.Y。如果 Min.X <= Max.X，并且 Y 也是如此，则它格式良好。点始终格式良好。对于格式良好的输入，矩形的方法始终返回格式良好的输出。

A Rectangle is also an Image whose bounds are the rectangle itself. At returns color.Opaque for points in the rectangle and color.Transparent otherwise.

​	矩形也是一个图像，其边界就是矩形本身。At 返回矩形中点的颜色。Opaque，否则返回颜色。Transparent。

```go
var ZR Rectangle
```

ZR is the zero Rectangle.

Deprecated: Use a literal image.Rectangle{} instead.

​	已弃用：改用文字图像。Rectangle{}。

#### func Rect

```go
func Rect(x0, y0, x1, y1 int) Rectangle
```

Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}. The returned rectangle has minimum and maximum coordinates swapped if necessary so that it is well-formed.

​	Rect 是 Rectangle{Pt(x0, y0), Pt(x1, y1)} 的简写。如果需要，返回的矩形具有交换的最小和最大坐标，以便它格式良好。

#### (Rectangle) Add

```go
func (r Rectangle) Add(p Point) Rectangle
```

Add returns the rectangle r translated by p.

​	Add 返回由 p 平移的矩形 r。

#### (Rectangle) At <- go1.5

```go
func (r Rectangle) At(x, y int) color.Color
```

At implements the Image interface.

​	At 实现 Image 接口。

#### (Rectangle) Bounds <- go1.5

```go
func (r Rectangle) Bounds() Rectangle
```

Bounds implements the Image interface.

​	Bounds 实现 Image 接口。

#### (Rectangle) Canon

```go
func (r Rectangle) Canon() Rectangle
```

Canon returns the canonical version of r. The returned rectangle has minimum and maximum coordinates swapped if necessary so that it is well-formed.

​	Canon 返回 r 的规范版本。如果需要，返回的矩形具有交换的最小和最大坐标，以便它格式良好。

#### (Rectangle) ColorModel <- go1.5

```go
func (r Rectangle) ColorModel() color.Model
```

ColorModel implements the Image interface.

​	ColorModel 实现 Image 接口。

#### (Rectangle) Dx

```go
func (r Rectangle) Dx() int
```

Dx returns r’s width.

#### (Rectangle) Dy

```go
func (r Rectangle) Dy() int
```

Dy returns r’s height.

​	Dy 返回 r 的高度。

#### (Rectangle) Empty

```go
func (r Rectangle) Empty() bool
```

Empty reports whether the rectangle contains no points.

​	Empty 报告矩形是否不包含任何点。

#### (Rectangle) Eq

```go
func (r Rectangle) Eq(s Rectangle) bool
```

Eq reports whether r and s contain the same set of points. All empty rectangles are considered equal.

​	Eq 报告 r 和 s 是否包含相同的点集。所有空矩形都被视为相等。

#### (Rectangle) In

```go
func (r Rectangle) In(s Rectangle) bool
```

In reports whether every point in r is in s.

​	In 报告 r 中的每个点是否都在 s 中。

#### (Rectangle) Inset

```go
func (r Rectangle) Inset(n int) Rectangle
```

Inset returns the rectangle r inset by n, which may be negative. If either of r’s dimensions is less than 2*n then an empty rectangle near the center of r will be returned.

​	Inset 返回矩形 r 向内缩进 n，n 可能为负值。如果 r 的任何一个维度小于 2*n，则将返回 r 中心附近的空矩形。

#### (Rectangle) Intersect 

```go
func (r Rectangle) Intersect(s Rectangle) Rectangle
```

Intersect returns the largest rectangle contained by both r and s. If the two rectangles do not overlap then the zero rectangle will be returned.

​	Intersect 返回 r 和 s 都包含的最大矩形。如果两个矩形不重叠，则返回零矩形。

#### (Rectangle) Overlaps

```go
func (r Rectangle) Overlaps(s Rectangle) bool
```

Overlaps reports whether r and s have a non-empty intersection.

​	Overlaps 报告 r 和 s 是否有非空交集。

#### (Rectangle) RGBA64At <- go1.17

```go
func (r Rectangle) RGBA64At(x, y int) color.RGBA64
```

RGBA64At implements the RGBA64Image interface.

​	RGBA64At 实现 RGBA64Image 接口。

#### (Rectangle) Size 

```go
func (r Rectangle) Size() Point
```

Size returns r’s width and height.

​	Size 返回 r 的宽度和高度。

#### (Rectangle) String 

```go
func (r Rectangle) String() string
```

String returns a string representation of r like “(3,4)-(6,5)”.

​	String 返回 r 的字符串表示形式，如 “(3,4)-(6,5)”。

#### (Rectangle) Sub

```go
func (r Rectangle) Sub(p Point) Rectangle
```

Sub returns the rectangle r translated by -p.

​	Sub 返回平移了 -p 的矩形 r。

#### (Rectangle) Union

```go
func (r Rectangle) Union(s Rectangle) Rectangle
```

Union returns the smallest rectangle that contains both r and s.

​	Union 返回包含 r 和 s 的最小矩形。

### type Uniform

```go
type Uniform struct {
	C color.Color
}
```

Uniform is an infinite-sized Image of uniform color. It implements the color.Color, color.Model, and Image interfaces.

​	Uniform 是一个颜色均匀的无限大小的图像。它实现了 color.Color、color.Model 和 Image 接口。

#### func NewUniform

```go
func NewUniform(c color.Color) *Uniform
```

NewUniform returns a new Uniform image of the given color.

​	NewUniform 返回一个给定颜色的新 Uniform 图像。

#### (*Uniform) At

```go
func (c *Uniform) At(x, y int) color.Color
```

#### (*Uniform) Bounds

```go
func (c *Uniform) Bounds() Rectangle
```

#### (*Uniform) ColorModel

```go
func (c *Uniform) ColorModel() color.Model
```

#### (*Uniform) Convert

```go
func (c *Uniform) Convert(color.Color) color.Color
```

#### (*Uniform) Opaque

```go
func (c *Uniform) Opaque() bool
```

Opaque scans the entire image and reports whether it is fully opaque.

​	Opaque 扫描整个图像并报告它是否完全不透明。

#### (*Uniform) RGBA

```go
func (c *Uniform) RGBA() (r, g, b, a uint32)
```

#### (*Uniform) RGBA64At <- go1.17

```go
func (c *Uniform) RGBA64At(x, y int) color.RGBA64
```

### type YCbCr

```go
type YCbCr struct {
	Y, Cb, Cr      []uint8
	YStride        int
	CStride        int
	SubsampleRatio YCbCrSubsampleRatio
	Rect           Rectangle
}
```

YCbCr is an in-memory image of Y’CbCr colors. There is one Y sample per pixel, but each Cb and Cr sample can span one or more pixels. YStride is the Y slice index delta between vertically adjacent pixels. CStride is the Cb and Cr slice index delta between vertically adjacent pixels that map to separate chroma samples. It is not an absolute requirement, but YStride and len(Y) are typically multiples of 8, and:

​	YCbCr 是 Y’CbCr 颜色的内存中图像。每个像素有一个 Y 样本，但每个 Cb 和 Cr 样本可以跨越一个或多个像素。YStride 是垂直相邻像素之间的 Y 切片索引增量。CStride 是垂直相邻像素之间的 Cb 和 Cr 切片索引增量，这些像素映射到单独的色度样本。这不是绝对要求，但 YStride 和 len(Y) 通常是 8 的倍数，并且：

```
For 4:4:4, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/1.
For 4:2:2, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/2.
For 4:2:0, CStride == YStride/2 && len(Cb) == len(Cr) == len(Y)/4.
For 4:4:0, CStride == YStride/1 && len(Cb) == len(Cr) == len(Y)/2.
For 4:1:1, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/4.
For 4:1:0, CStride == YStride/4 && len(Cb) == len(Cr) == len(Y)/8.
```

#### func NewYCbCr

```go
func NewYCbCr(r Rectangle, subsampleRatio YCbCrSubsampleRatio) *YCbCr
```

NewYCbCr returns a new YCbCr image with the given bounds and subsample ratio.

​	NewYCbCr 返回具有给定边界和子采样率的新 YCbCr 图像。

#### (*YCbCr) At

```go
func (p *YCbCr) At(x, y int) color.Color
```

#### (*YCbCr) Bounds

```go
func (p *YCbCr) Bounds() Rectangle
```

#### (*YCbCr) COffset

```go
func (p *YCbCr) COffset(x, y int) int
```

COffset returns the index of the first element of Cb or Cr that corresponds to the pixel at (x, y).

​	COffset 返回对应于 (x, y) 处像素的 Cb 或 Cr 的第一个元素的索引。

#### (*YCbCr) ColorModel

```go
func (p *YCbCr) ColorModel() color.Model
```

#### (*YCbCr) Opaque

```go
func (p *YCbCr) Opaque() bool
```

#### (*YCbCr) RGBA64At <- go1.17

```go
func (p *YCbCr) RGBA64At(x, y int) color.RGBA64
```

#### (*YCbCr) SubImage

```go
func (p *YCbCr) SubImage(r Rectangle) Image
```

SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

​	SubImage 返回一个图像，表示通过 r 可见的图像 p 的部分。返回值与原始图像共享像素。

#### (*YCbCr) YCbCrAt <- go1.4

```go
func (p *YCbCr) YCbCrAt(x, y int) color.YCbCr
```

#### (*YCbCr) YOffset

```go
func (p *YCbCr) YOffset(x, y int) int
```

YOffset returns the index of the first element of Y that corresponds to the pixel at (x, y).

​	YOffset 返回对应于像素 (x, y) 的 Y 的第一个元素的索引。

### type YCbCrSubsampleRatio

```go
type YCbCrSubsampleRatio int
```

YCbCrSubsampleRatio is the chroma subsample ratio used in a YCbCr image.

​	YCbCrSubsampleRatio 是 YCbCr 图像中使用的色度子采样率。

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