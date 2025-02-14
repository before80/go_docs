+++
title = "devices"
date = 2024-11-20T18:02:07+08:00
weight = 70
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/go-rod/rod/lib/devices](https://pkg.go.dev/github.com/go-rod/rod/lib/devices)
>
> 收录该文档时间：`2024-11-20T18:02:07+08:00`
>
> [Version: v0.116.2](https://pkg.go.dev/github.com/go-rod/rod/lib/devices?tab=versions)



## Overview 

Package devices ...

## 常量

This section is empty.

## 变量

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/devices/list.go#L5)

``` go
var (

	// IPhone4 device.
	IPhone4 = Device{
		Title:          "iPhone 4",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_2 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D257 Safari/9537.53",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  480,
				Height: 320,
			},
			Vertical: ScreenSize{
				Width:  320,
				Height: 480,
			},
		},
	}

	// IPhone5orSE device.
	IPhone5orSE = Device{
		Title:          "iPhone 5/SE",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  568,
				Height: 320,
			},
			Vertical: ScreenSize{
				Width:  320,
				Height: 568,
			},
		},
	}

	// IPhone6or7or8 device.
	IPhone6or7or8 = Device{
		Title:          "iPhone 6/7/8",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  667,
				Height: 375,
			},
			Vertical: ScreenSize{
				Width:  375,
				Height: 667,
			},
		},
	}

	// IPhone6or7or8Plus device.
	IPhone6or7or8Plus = Device{
		Title:          "iPhone 6/7/8 Plus",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 3,
			Horizontal: ScreenSize{
				Width:  736,
				Height: 414,
			},
			Vertical: ScreenSize{
				Width:  414,
				Height: 736,
			},
		},
	}

	// IPhoneX device.
	IPhoneX = Device{
		Title:          "iPhone X",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 3,
			Horizontal: ScreenSize{
				Width:  812,
				Height: 375,
			},
			Vertical: ScreenSize{
				Width:  375,
				Height: 812,
			},
		},
	}

	// BlackBerryZ30 device.
	BlackBerryZ30 = Device{
		Title:          "BlackBerry Z30",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (BB10; Touch) AppleWebKit/537.10+ (KHTML, like Gecko) Version/10.0.9.2372 Mobile Safari/537.10+",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 360,
			},
			Vertical: ScreenSize{
				Width:  360,
				Height: 640,
			},
		},
	}

	// Nexus4 device.
	Nexus4 = Device{
		Title:          "Nexus 4",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 4.4.2; Nexus 4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 384,
			},
			Vertical: ScreenSize{
				Width:  384,
				Height: 640,
			},
		},
	}

	// Nexus5 device.
	Nexus5 = Device{
		Title:          "Nexus 5",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 3,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 360,
			},
			Vertical: ScreenSize{
				Width:  360,
				Height: 640,
			},
		},
	}

	// Nexus5X device.
	Nexus5X = Device{
		Title:          "Nexus 5X",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 8.0.0; Nexus 5X Build/OPR4.170623.006) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  732,
				Height: 412,
			},
			Vertical: ScreenSize{
				Width:  412,
				Height: 732,
			},
		},
	}

	// Nexus6 device.
	Nexus6 = Device{
		Title:          "Nexus 6",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 7.1.1; Nexus 6 Build/N6F26U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 3,
			Horizontal: ScreenSize{
				Width:  732,
				Height: 412,
			},
			Vertical: ScreenSize{
				Width:  412,
				Height: 732,
			},
		},
	}

	// Nexus6P device.
	Nexus6P = Device{
		Title:          "Nexus 6P",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 8.0.0; Nexus 6P Build/OPP3.170518.006) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 3,
			Horizontal: ScreenSize{
				Width:  732,
				Height: 412,
			},
			Vertical: ScreenSize{
				Width:  412,
				Height: 732,
			},
		},
	}

	// Pixel2 device.
	Pixel2 = Device{
		Title:          "Pixel 2",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  731,
				Height: 411,
			},
			Vertical: ScreenSize{
				Width:  411,
				Height: 731,
			},
		},
	}

	// Pixel2XL device.
	Pixel2XL = Device{
		Title:          "Pixel 2 XL",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 3,
			Horizontal: ScreenSize{
				Width:  823,
				Height: 411,
			},
			Vertical: ScreenSize{
				Width:  411,
				Height: 823,
			},
		},
	}

	// LGOptimusL70 device.
	LGOptimusL70 = Device{
		Title:          "LG Optimus L70",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; U; Android 4.4.2; en-us; LGMS323 Build/KOT49I.MS32310c) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 1,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 384,
			},
			Vertical: ScreenSize{
				Width:  384,
				Height: 640,
			},
		},
	}

	// NokiaN9 device.
	NokiaN9 = Device{
		Title:          "Nokia N9",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (MeeGo; NokiaN9) AppleWebKit/534.13 (KHTML, like Gecko) NokiaBrowser/8.5.0 Mobile Safari/534.13",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 1,
			Horizontal: ScreenSize{
				Width:  854,
				Height: 480,
			},
			Vertical: ScreenSize{
				Width:  480,
				Height: 854,
			},
		},
	}

	// NokiaLumia520 device.
	NokiaLumia520 = Device{
		Title:          "Nokia Lumia 520",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0; ARM; Touch; NOKIA; Lumia 520)",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 1,
			Horizontal: ScreenSize{
				Width:  533,
				Height: 320,
			},
			Vertical: ScreenSize{
				Width:  320,
				Height: 533,
			},
		},
	}

	// MicrosoftLumia550 device.
	MicrosoftLumia550 = Device{
		Title:          "Microsoft Lumia 550",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 550) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/14.14263",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 360,
			},
			Vertical: ScreenSize{
				Width:  640,
				Height: 360,
			},
		},
	}

	// MicrosoftLumia950 device.
	MicrosoftLumia950 = Device{
		Title:          "Microsoft Lumia 950",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; Microsoft; Lumia 950) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Mobile Safari/537.36 Edge/14.14263",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 4,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 360,
			},
			Vertical: ScreenSize{
				Width:  360,
				Height: 640,
			},
		},
	}

	// GalaxySIII device.
	GalaxySIII = Device{
		Title:          "Galaxy S III",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; U; Android 4.0; en-us; GT-I9300 Build/IMM76D) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 360,
			},
			Vertical: ScreenSize{
				Width:  360,
				Height: 640,
			},
		},
	}

	// GalaxyS5 device.
	GalaxyS5 = Device{
		Title:          "Galaxy S5",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 3,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 360,
			},
			Vertical: ScreenSize{
				Width:  360,
				Height: 640,
			},
		},
	}

	// JioPhone2 device.
	JioPhone2 = Device{
		Title:          "JioPhone 2",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Mobile; LYF/F300B/LYF-F300B-001-01-15-130718-i;Android; rv:48.0) Gecko/48.0 Firefox/48.0 KAIOS/2.5",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 1,
			Horizontal: ScreenSize{
				Width:  320,
				Height: 240,
			},
			Vertical: ScreenSize{
				Width:  240,
				Height: 320,
			},
		},
	}

	// KindleFireHDX device.
	KindleFireHDX = Device{
		Title:          "Kindle Fire HDX",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; U; en-us; KFAPWI Build/JDQ39) AppleWebKit/535.19 (KHTML, like Gecko) Silk/3.13 Safari/535.19 Silk-Accelerated=true",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  1280,
				Height: 800,
			},
			Vertical: ScreenSize{
				Width:  800,
				Height: 1280,
			},
		},
	}

	// IPadMini device.
	IPadMini = Device{
		Title:          "iPad Mini",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  1024,
				Height: 768,
			},
			Vertical: ScreenSize{
				Width:  768,
				Height: 1024,
			},
		},
	}

	// IPad device.
	IPad = Device{
		Title:          "iPad",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  1024,
				Height: 768,
			},
			Vertical: ScreenSize{
				Width:  768,
				Height: 1024,
			},
		},
	}

	// IPadPro device.
	IPadPro = Device{
		Title:          "iPad Pro",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  1366,
				Height: 1024,
			},
			Vertical: ScreenSize{
				Width:  1024,
				Height: 1366,
			},
		},
	}

	// BlackberryPlayBook device.
	BlackberryPlayBook = Device{
		Title:          "Blackberry PlayBook",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (PlayBook; U; RIM Tablet OS 2.1.0; en-US) AppleWebKit/536.2+ (KHTML like Gecko) Version/7.2.1.0 Safari/536.2+",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 1,
			Horizontal: ScreenSize{
				Width:  1024,
				Height: 600,
			},
			Vertical: ScreenSize{
				Width:  600,
				Height: 1024,
			},
		},
	}

	// Nexus10 device.
	Nexus10 = Device{
		Title:          "Nexus 10",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 10 Build/MOB31T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  1280,
				Height: 800,
			},
			Vertical: ScreenSize{
				Width:  800,
				Height: 1280,
			},
		},
	}

	// Nexus7 device.
	Nexus7 = Device{
		Title:          "Nexus 7",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 7 Build/MOB30X) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  960,
				Height: 600,
			},
			Vertical: ScreenSize{
				Width:  600,
				Height: 960,
			},
		},
	}

	// GalaxyNote3 device.
	GalaxyNote3 = Device{
		Title:          "Galaxy Note 3",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; U; Android 4.3; en-us; SM-N900T Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 3,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 360,
			},
			Vertical: ScreenSize{
				Width:  360,
				Height: 640,
			},
		},
	}

	// GalaxyNoteII device.
	GalaxyNoteII = Device{
		Title:          "Galaxy Note II",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; U; Android 4.1; en-us; GT-N7100 Build/JRO03C) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 360,
			},
			Vertical: ScreenSize{
				Width:  360,
				Height: 640,
			},
		},
	}

	// LaptopWithTouch device.
	LaptopWithTouch = Device{
		Title:          "Laptop with touch",
		Capabilities:   []string{"touch"},
		UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 1,
			Horizontal: ScreenSize{
				Width:  1280,
				Height: 950,
			},
			Vertical: ScreenSize{
				Width:  950,
				Height: 1280,
			},
		},
	}

	// LaptopWithHiDPIScreen device.
	LaptopWithHiDPIScreen = Device{
		Title:          "Laptop with HiDPI screen",
		Capabilities:   []string{},
		UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  1440,
				Height: 900,
			},
			Vertical: ScreenSize{
				Width:  900,
				Height: 1440,
			},
		},
	}

	// LaptopWithMDPIScreen device.
	LaptopWithMDPIScreen = Device{
		Title:          "Laptop with MDPI screen",
		Capabilities:   []string{},
		UserAgent:      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 1,
			Horizontal: ScreenSize{
				Width:  1280,
				Height: 800,
			},
			Vertical: ScreenSize{
				Width:  800,
				Height: 1280,
			},
		},
	}

	// MotoG4 device.
	MotoG4 = Device{
		Title:          "Moto G4",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 6.0.1; Moto G (4)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 3,
			Horizontal: ScreenSize{
				Width:  640,
				Height: 360,
			},
			Vertical: ScreenSize{
				Width:  360,
				Height: 640,
			},
		},
	}

	// SurfaceDuo device.
	SurfaceDuo = Device{
		Title:          "Surface Duo",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 2,
			Horizontal: ScreenSize{
				Width:  720,
				Height: 540,
			},
			Vertical: ScreenSize{
				Width:  540,
				Height: 720,
			},
		},
	}

	// GalaxyFold device.
	GalaxyFold = Device{
		Title:          "Galaxy Fold",
		Capabilities:   []string{"touch", "mobile"},
		UserAgent:      "Mozilla/5.0 (Linux; Android 8.0; Pixel 2 Build/OPD3.170816.012) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36",
		AcceptLanguage: "en",
		Screen: Screen{
			DevicePixelRatio: 3,
			Horizontal: ScreenSize{
				Width:  653,
				Height: 280,
			},
			Vertical: ScreenSize{
				Width:  280,
				Height: 653,
			},
		},
	}
)
```

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/devices/utils.go#L4)

``` go
var Clear = Device{/* contains filtered or unexported fields */}
```

Clear is used to clear overrides.

## 函数

This section is empty.

## 类型

### type Device <- 0.62.0

``` go
type Device struct {
	Capabilities   []string
	UserAgent      string
	AcceptLanguage string
	Screen         Screen
	Title          string
	// contains filtered or unexported fields
}
```

Device represents a emulated device.

​	Device 表示一个模拟设备。

#### (Device) IsClear <- 0.84.0

``` go
func (device Device) IsClear() bool
```

IsClear type.

​	IsClear 类型。

#### (Device) Landscape <- 0.112.1

``` go
func (device Device) Landscape() Device
```

Landscape clones the device and set it to landscape mode.

​	Landscape 克隆设备并设置为横屏模式。

#### (Device) MetricsEmulation <- 0.84.0

``` go
func (device Device) MetricsEmulation() *proto.EmulationSetDeviceMetricsOverride
```

MetricsEmulation config.

​	MetricsEmulation 配置。

#### (Device) TouchEmulation <- 0.84.0

``` go
func (device Device) TouchEmulation() *proto.EmulationSetTouchEmulationEnabled
```

TouchEmulation config.

​	TouchEmulation 配置。

#### (Device) UserAgentEmulation <- 0.84.0

``` go
func (device Device) UserAgentEmulation() *proto.NetworkSetUserAgentOverride
```

UserAgentEmulation config.

​	UserAgentEmulation 配置。

### type Screen <- 0.84.0

``` go
type Screen struct {
	DevicePixelRatio float64
	Horizontal       ScreenSize
	Vertical         ScreenSize
}
```

Screen represents the screen of a device.

​	Screen 表示设备的屏幕。

### type ScreenSize <- 0.84.0

``` go
type ScreenSize struct {
	Width  int
	Height int
}
```

ScreenSize represents the size of the screen.

​	ScreenSize 表示屏幕的尺寸。
