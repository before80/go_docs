+++
title = "defaults"
date = 2024-11-20T18:01:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

### Overview 

Package defaults of commonly used options parsed from environment. Check ResetWith for details.

### Index 

- [Variables](https://pkg.go.dev/github.com/go-rod/rod/lib/defaults#pkg-variables)
- [func Reset()](https://pkg.go.dev/github.com/go-rod/rod/lib/defaults#Reset)
- [func ResetWith(options string)](https://pkg.go.dev/github.com/go-rod/rod/lib/defaults#ResetWith)

### Constants 

This section is empty.

### Variables 

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L48)

``` go
var Bin string
```

Bin is the default of launcher.Launcher.Bin . Option name is "bin".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L64)

``` go
var CDP utils.Logger
```

CDP is the default of cdp.Client.Logger Option name is "cdp".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L36)

``` go
var Devtools bool
```

Devtools is the default of launcher.Launcher.Devtools . Option name is "devtools".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L40)

``` go
var Dir string
```

Dir is the default of launcher.Launcher.UserDataDir . Option name is "dir".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L56)

``` go
var LockPort int
```

LockPort is the default of launcher.Browser.LockPort Option name is "lock".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L28)

``` go
var Monitor string
```

Monitor is the default of rod.Browser.ServeMonitor . Option name is "monitor".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L44)

``` go
var Port string
```

Port is the default of launcher.Launcher.RemoteDebuggingPort . Option name is "port".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L52)

``` go
var Proxy string
```

Proxy is the default of launcher.Launcher.Proxy Option name is "proxy".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L32)

``` go
var Show bool
```

Show is the default of launcher.Launcher.Headless . Option name is "show".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L24)

``` go
var Slow time.Duration
```

Slow is the default of rod.Browser.SlowMotion . The format is same as https://golang.org/pkg/time/#ParseDuration Option name is "slow".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L19)

``` go
var Trace bool
```

Trace is the default of rod.Browser.Trace . Option name is "trace".

[View Source](https://github.com/go-rod/rod/blob/v0.116.2/lib/defaults/defaults.go#L60)

``` go
var URL string
```

URL is the default websocket url for remote control a browser. Option name is "url".

### Functions 

#### func Reset <- 0.47.0

``` go
func Reset()
```

Reset all flags to their init values.

#### func ResetWith <- 0.106.0

``` go
func ResetWith(options string)
```

ResetWith options and "-rod" command line flag. It will be called in an init() , so you don't have to call it manually. It will try to load the cli flag "-rod" and then the options, the later override the former. If you want to disable the global cli argument flag, set env DISABLE_ROD_FLAG. Values are separated by commas, key and value are separated by "=". For example:

```
go run main.go -rod=show
go run main.go -rod show,trace,slow=1s,monitor
go run main.go --rod="slow=1s,dir=path/has /space,monitor=:9223"
```

### Types 

This section is empty.
