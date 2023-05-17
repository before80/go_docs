+++
title = "color/palette"
date = 2023-05-17T11:11:20+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# palette

https://pkg.go.dev/image/color/palette@go1.20.1



Package palette provides standard color palettes.



## 常量 

This section is empty.

## 变量

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/image/color/palette/palette.go;l=23)

``` go linenums="1"
var Plan9 = []color.Color{}/* 256 elements not displayed */
```

Plan9 is a 256-color palette that partitions the 24-bit RGB space into 4×4×4 subdivision, with 4 shades in each subcube. Compared to the WebSafe, the idea is to reduce the color resolution by dicing the color cube into fewer cells, and to use the extra space to increase the intensity resolution. This results in 16 gray shades (4 gray subcubes with 4 samples in each), 13 shades of each primary and secondary color (3 subcubes with 4 samples plus black) and a reasonable selection of colors covering the rest of the color cube. The advantage is better representation of continuous tones.

This palette was used in the Plan 9 Operating System, described at https://9p.io/magic/man2html/6/color

[View Source](https://cs.opensource.google/go/go/+/go1.20.1:src/image/color/palette/palette.go;l=286)

``` go linenums="1"
var WebSafe = []color.Color{}/* 216 elements not displayed */
```

WebSafe is a 216-color palette that was popularized by early versions of Netscape Navigator. It is also known as the Netscape Color Cube.

See https://en.wikipedia.org/wiki/Web_colors#Web-safe_colors for details.

## 函数

This section is empty.

## 类型

This section is empty.