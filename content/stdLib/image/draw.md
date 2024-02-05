+++
title = "draw"
date = 2023-05-17T11:11:20+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++
> 原文：[https://pkg.go.dev/image/draw@go1.21.3](https://pkg.go.dev/image/draw@go1.21.3)

Package draw provides image composition functions.

See "The Go image/draw package" for an introduction to this package: https://golang.org/doc/articles/image_draw.html

## 常量 

This section is empty.

## 变量

This section is empty.

## 函数

### func Draw 

``` go 
func Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point, op Op)
```

Draw calls DrawMask with a nil mask.

### func DrawMask 

``` go 
func DrawMask(dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op)
```

DrawMask aligns r.Min in dst with sp in src and mp in mask and then replaces the rectangle r in dst with the result of a Porter-Duff composition. A nil mask is treated as opaque.

## 类型

### type Drawer  <- go1.2

``` go 
type Drawer interface {
	// Draw aligns r.Min in dst with sp in src and then replaces the
	// rectangle r in dst with the result of drawing src on dst.
	Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
}
```

Drawer contains the Draw method.

#### Example (FloydSteinberg)
``` go 
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
)

func main() {
	const width = 130
	const height = 50

	im := image.NewGray(image.Rectangle{Max: image.Point{X: width, Y: height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dist := math.Sqrt(math.Pow(float64(x-width/2), 2)/3+math.Pow(float64(y-height/2), 2)) / (height / 1.5) * 255
			var gray uint8
			if dist > 255 {
				gray = 255
			} else {
				gray = uint8(dist)
			}
			im.SetGray(x, y, color.Gray{Y: 255 - gray})
		}
	}
	pi := image.NewPaletted(im.Bounds(), []color.Color{
		color.Gray{Y: 255},
		color.Gray{Y: 160},
		color.Gray{Y: 70},
		color.Gray{Y: 35},
		color.Gray{Y: 0},
	})

	draw.FloydSteinberg.Draw(pi, im.Bounds(), im, image.ZP)
	shade := []string{" ", "░", "▒", "▓", "█"}
	for i, p := range pi.Pix {
		fmt.Print(shade[p])
		if (i+1)%width == 0 {
			fmt.Print("\n")
		}
	}
}

Output:
```

FloydSteinberg is a Drawer that is the Src Op with Floyd-Steinberg error diffusion.

### type Image 

``` go 
type Image interface {
	image.Image
	Set(x, y int, c color.Color)
}
```

Image is an image.Image with a Set method to change a single pixel.

### type Op 

``` go 
type Op int
```

Op is a Porter-Duff compositing operator.

``` go 
const (
	// Over specifies "(src in mask) over dst".
	Over Op = iota
	// Src specifies "src in mask".
	Src
)
```

#### (Op) Draw  <- go1.2

``` go 
func (op Op) Draw(dst Image, r image.Rectangle, src image.Image, sp image.Point)
```

Draw implements the Drawer interface by calling the Draw function with this Op.

### type Quantizer  <- go1.2

``` go 
type Quantizer interface {
	// Quantize appends up to cap(p) - len(p) colors to p and returns the
	// updated palette suitable for converting m to a paletted image.
	Quantize(p color.Palette, m image.Image) color.Palette
}
```

Quantizer produces a palette for an image.

### type RGBA64Image  <- go1.17

``` go 
type RGBA64Image interface {
	image.RGBA64Image
	Set(x, y int, c color.Color)
	SetRGBA64(x, y int, c color.RGBA64)
}
```

RGBA64Image extends both the Image and image.RGBA64Image interfaces with a SetRGBA64 method to change a single pixel. SetRGBA64 is equivalent to calling Set, but it can avoid allocations from converting concrete color types to the color.Color interface type.