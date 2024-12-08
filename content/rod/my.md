+++
title = "个人笔记"
date = 2024-11-23T18:19:37+08:00
weight = -10000
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

## 默认情况下会自动下载chromium浏览器

```go

```



## 设置指定的本地浏览器

​	示例

```go
	path := "/usr/bin/google-chrome"
	userDataDir := "/home/lx/.config/google-chrome"

	u := launcher.New().
		Bin(path).
		Headless(false).
		Set("window-size", "1920,1080").
		Set("user-data-dir", userDataDir). // 加载用户配置
		MustLaunch()

	// 使用 Rod 连接到本地 Chrome
	browser := rod.New().ControlURL(u).MustConnect()

	// 打开 GitHub 首页
	page := browser.MustPage("https://github.com/search?q=golang")
	// 设置窗口大小为 1920x1080
	page.MustSetViewport(1920, 1080, 1, false)
```

## 通过自动查找本地浏览器

示例

```go
	path, _ := launcher.LookPath()
	userDataDir := "/home/lx/.config/google-chrome"

	u := launcher.New().
		Bin(path).
		Headless(false).
		Set("window-size", "1920,1080").
		Set("user-data-dir", userDataDir). // 加载用户配置
		MustLaunch()

	// 使用 Rod 连接到本地 Chrome
	browser := rod.New().ControlURL(u).MustConnect()

	// 打开 GitHub 首页
	page := browser.MustPage("https://github.com/search?q=golang")
	// 设置窗口大小为 1920x1080
	page.MustSetViewport(1920, 1080, 1, false)
```

> launcher.LookPath函数的源码如下：
>
> ```go
> // LookPath searches for the browser executable from often used paths on current operating system.
> func LookPath() (found string, has bool) {
> 	list := map[string][]string{
> 		"darwin": {
> 			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
> 			"/Applications/Chromium.app/Contents/MacOS/Chromium",
> 			"/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge",
> 			"/Applications/Google Chrome Canary.app/Contents/MacOS/Google Chrome Canary",
> 			"/usr/bin/google-chrome-stable",
> 			"/usr/bin/google-chrome",
> 			"/usr/bin/chromium",
> 			"/usr/bin/chromium-browser",
> 		},
> 		"linux": {
> 			"chrome",
> 			"google-chrome",
> 			"/usr/bin/google-chrome",
> 			"microsoft-edge",
> 			"/usr/bin/microsoft-edge",
> 			"chromium",
> 			"chromium-browser",
> 			"/usr/bin/google-chrome-stable",
> 			"/usr/bin/chromium",
> 			"/usr/bin/chromium-browser",
> 			"/snap/bin/chromium",
> 			"/data/data/com.termux/files/usr/bin/chromium-browser",
> 		},
> 		"openbsd": {
> 			"chrome",
> 			"chromium",
> 		},
> 		"windows": append([]string{"chrome", "edge"}, expandWindowsExePaths(
> 			`Google\Chrome\Application\chrome.exe`,
> 			`Chromium\Application\chrome.exe`,
> 			`Microsoft\Edge\Application\msedge.exe`,
> 		)...),
> 	}[runtime.GOOS]
> 
> 	for _, path := range list {
> 		var err error
> 		found, err = exec.LookPath(path)
> 		has = err == nil
> 		if has {
> 			break
> 		}
> 	}
> 
> 	return
> }
> ```
>
> ​	在遍历以上函数中路径切片时，若能找到浏览器，则不再进一步进行查找。也就是说，路径优先级有很大关系！

## 清除指定网站的Cookies

​	示例

```go
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// 打开页面
	page := browser.MustPage("https://baidu.com")
	page.MustWaitLoad()

	// 使用 JavaScript 获取 cookies
	cookies1 := page.MustEval(`() => document.cookie`).String()
	fmt.Println("Cookies:", cookies1)

	// 获取当前页面的 cookies
	cookies, err := proto.NetworkGetCookies{
		Urls: []string{"https://baidu.com"},
	}.Call(page)
	if err != nil {
		panic(err)
	}
	// 打印当前页面的 cookies
	for _, cookie := range cookies.Cookies {
		fmt.Printf("1 Name: %s, Value: %s, Domain: %s\n", cookie.Name, cookie.Value, cookie.Domain)
	}
	// 清除所有 cookies
	err = proto.NetworkClearBrowserCookies{}.Call(page)
	if err != nil {
		panic(err)
	}

	println("Cookies cleared")

	cookies, err = proto.NetworkGetCookies{
		Urls: []string{"https://baidu.com"},
	}.Call(page)
	if err != nil {
		panic(err)
	}
	// 打印当前页面的 cookies
	for _, cookie := range cookies.Cookies {
		fmt.Printf("2 Name: %s, Value: %s, Domain: %s\n", cookie.Name, cookie.Value, cookie.Domain)
	}
```

## 劫持Js并替换其内容

```go
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// 打开页面
	page := browser.MustPage("https://baidu.com")

	// 自定义的替换内容
	customJS := `
		console.log("This is a custom JS file");
		window.customInjected = true;
	`

	// 启用网络拦截
	_ = proto.FetchEnable{
		Patterns: []*proto.FetchRequestPattern{
			{
				URLPattern: "*.js", // 匹配所有 JS 文件
			},
		},
	}.Call(page)

	// 监听并处理网络请求
	go page.EachEvent(func(e *proto.FetchRequestPaused) {
		// 检查是否为目标 JS 文件
		if e.Request.URL == "https://hector.baidu.com/a.js" {
			fmt.Println("found the js")
			_ = proto.FetchFulfillRequest{
				RequestID:    e.RequestID,
				ResponseCode: 200,
				Body:         []byte(customJS),
			}.Call(page)
		} else {
			// 继续加载其他资源
			_ = proto.FetchContinueRequest{
				RequestID: e.RequestID,
			}.Call(page)
		}
	})()

	// 加载页面
	page.MustWaitLoad()

	// 检查替换结果
	res := page.MustEval("() => window.customInjected").Bool()
	fmt.Println("Custom JS Injected:", res)
```

