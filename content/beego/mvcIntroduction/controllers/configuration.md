+++
title = "配置"
date = 2024-02-04T09:55:24+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/config/](https://beego.wiki/docs/mvc/controller/config/)

# Configuration 配置



## Configuration 配置

By default the Beego configuration file uses the INI format. Other supported formats include XML, JSON, and YAML.

&zeroWidthSpace;默认情况下，Beego 配置文件使用 INI 格式。其他支持的格式包括 XML、JSON 和 YAML。

## Default configurations parsing 默认配置解析

Beego will parse the `conf/app.conf` file by default.

&zeroWidthSpace;Beego 默认会解析 `conf/app.conf` 文件。

Many default variables can be initialized in this file:

&zeroWidthSpace;在这个文件中可以初始化许多默认变量：

```
appname = beepkg
httpaddr = "127.0.0.1"
httpport = 9090
runmode ="dev"
autorender = false
recoverpanic = false
viewspath = "myview"
```

These configurations will replace Beego’s default values.

&zeroWidthSpace;这些配置会替换 Beego 的默认值。

Other application specific values can also be set using this file, such as database connection details:

&zeroWidthSpace;也可以使用此文件设置其他应用程序特定的值，例如数据库连接详细信息：

```
mysqluser = "root"
mysqlpass = "rootpass"
mysqlurls = "127.0.0.1"
mysqldb   = "beego"
```

These configurations can be accessed like this:

&zeroWidthSpace;可以像这样访问这些配置：

```
beego.AppConfig.String("mysqluser")
beego.AppConfig.String("mysqlpass")
beego.AppConfig.String("mysqlurls")
beego.AppConfig.String("mysqldb")
```

AppConfig’s methods:

&zeroWidthSpace;AppConfig 的方法：

- Set(key, val string) error
- String(key string) string
- Strings(key string) []string
- Int(key string) (int, error)
- Int64(key string) (int64, error)
- Bool(key string) (bool, error)
- Float(key string) (float64, error)
- DefaultString(key string, defaultVal string) string
- DefaultStrings(key string, defaultVal []string)
- DefaultInt(key string, defaultVal int) int
- DefaultInt64(key string, defaultVal int64) int64
- DefaultBool(key string, defaultVal bool) bool
- DefaultFloat(key string, defaultVal float64) float64
- DIY(key string) (interface{}, error)
- GetSection(section string) (map[string]string, error)
- SaveConfigFile(filename string) error

When using the INI format the key supports the section::key pattern.

&zeroWidthSpace;使用 INI 格式时，键支持 section::key 模式。

The Default* methods can be used to return default values if the config file cannot be read.

&zeroWidthSpace;如果无法读取配置文件，可以使用 Default* 方法返回默认值。

### Configurations for Different Environments 不同环境的配置

Configurations for different runmodes can be set under their own sections. Beego will take the configurations of the current runmode by default. For example:

&zeroWidthSpace;不同运行模式的配置可以设置在它们自己的部分下。默认情况下，Beego 将采用当前运行模式的配置。例如：

appname = beepkg httpaddr = “127.0.0.1” httpport = 9090 runmode =“dev” autorender = false recoverpanic = false viewspath = “myview”

[dev] httpport = 8080 [prod] httpport = 8088 [test] httpport = 8888

&zeroWidthSpace;[开发] httpport = 8080 [生产] httpport = 8088 [测试] httpport = 8888

The configurations above set up httpport for dev, prod and test environments. Beego will take httpport = 8080 for the current runmode “dev”.

&zeroWidthSpace;上面的配置为 dev、prod 和 test 环境设置了 httpport。Beego 将为当前运行模式“dev”获取 httpport = 8080。

To get config to operate under a different runmode use “runmode::key”. For example:

&zeroWidthSpace;若要让配置在不同的运行模式下运行，请使用“runmode::key”。例如：

```
beego.AppConfig.String("dev::mysqluser")
```

For custom configs use `beego.GetConfig(typ, key string)` to get the config.

&zeroWidthSpace;对于自定义配置，请使用 `beego.GetConfig(typ, key string)` 获取配置。

### Multiple config files 多个配置文件

The INI config file supports `include` to including multiple config files.

&zeroWidthSpace;INI 配置文件支持 `include` 以包括多个配置文件。

app.conf

```
appname = beepkg
httpaddr = "127.0.0.1"
httpport = 9090

include "app2.conf"
```

app2.conf

```
runmode ="dev"
autorender = false
recoverpanic = false
viewspath = "myview"

[dev]
httpport = 8080
[prod]
httpport = 8088
[test]
httpport = 8888
```

### Beego default config variables Beego 默认配置变量

Beego includes many configurable variables. These can be configured and overwritten in `conf/app.conf`.

&zeroWidthSpace;Beego 包含许多可配置变量。这些变量可以在 `conf/app.conf` 中进行配置和覆盖。

#### Basic config 基本配置

```go
// now only support ini, next will support json.
func parseConfig(appConfigPath string) (err error) {
	AppConfig, err = newAppConfig(appConfigProvider, appConfigPath)
	if err != nil {
		return err
	}
	return assignConfig(AppConfig)
}
```

- LoadAppConfig The file format of LoadAppConfig. By default this is `ini`. Other valid formats include `xml`, `yaml`, and `json`. The application configuration file path. By default this is `conf/app.conf`.

  &zeroWidthSpace;LoadAppConfig LoadAppConfig 的文件格式。默认情况下，这是 `ini` 。其他有效格式包括 `xml` 、 `yaml` 和 `json` 。应用程序配置文件路径。默认情况下，这是 `conf/app.conf` 。

  `beego.LoadAppConfig("yaml", "conf/app.conf")`

#### App config 应用程序配置

- AppName

  The application name. By default this is “Beego”. If the application is created by `bee new project_name` it will be set to project_name.

  &zeroWidthSpace;应用程序名称。默认情况下，这是“Beego”。如果应用程序由 `bee new project_name` 创建，则将其设置为 project_name。

  `beego.BConfig.AppName = "beego"`

- RunMode

  The application mode. By default this is set to `dev`. Other valid modes include `prod` and `test`. In `dev` mode user friendly error pages will be shown. In `prod` mode user friendly error pages will not be rendered.

  &zeroWidthSpace;应用程序模式。默认情况下，将其设置为 `dev` 。其他有效模式包括 `prod` 和 `test` 。在 `dev` 模式下，将显示用户友好的错误页面。在 `prod` 模式下，不会呈现用户友好的错误页面。

  `beego.BConfig.RunMode = "dev"`

- RouterCaseSensitive

  Set case sensitivity for the router. By default this value is true.

  &zeroWidthSpace;为路由器设置大小写敏感性。默认情况下，此值为 true。

  `beego.BConfig.RouterCaseSensitive = true`

- ServerName

  The Beego server name. By default this name is `beego`.

  &zeroWidthSpace;Beego 服务器名称。默认情况下，此名称为 `beego` 。

  `beego.BConfig.ServerName = "beego"`

- RecoverPanic

  When active the application will recover from exceptions without exiting the application. By default this is set to true.

  &zeroWidthSpace;当应用程序处于活动状态时，它将从异常中恢复，而不会退出应用程序。默认情况下，此项设置为 true。

  `beego.BConfig.RecoverPanic = true`

- CopyRequestBody

  Toggle copying of raw request body in context. By default this is false except for GET, HEAD or file uploading.

  &zeroWidthSpace;在上下文中切换复制原始请求正文。默认情况下，此项为 false，但 GET、HEAD 或文件上传除外。

  `beego.BConfig.CopyRequestBody = false`

- EnableGzip

  Enable Gzip. By default this is false. If Gzip is enabled the output of templates will be compressed by Gzip or zlib according to the `Accept-Encoding` setting of the browser.

  &zeroWidthSpace;启用 Gzip。默认情况下，此项为 false。如果启用了 Gzip，模板的输出将根据浏览器的 `Accept-Encoding` 设置由 Gzip 或 zlib 压缩。

  `beego.BConfig.EnableGzip = false`

  Further properties can be configured as below:

  &zeroWidthSpace;其他属性可以按如下方式配置：

  `gzipCompressLevel = 9` Sets the compression level used for deflate compression(0-9). By default is 9 (best speed).

  &zeroWidthSpace; `gzipCompressLevel = 9` 设置用于压缩压缩的压缩级别 (0-9)。默认值为 9（最佳速度）。

  `gzipMinLength = 256` Original content will only be compressed if length is either unknown or greater than gzipMinLength. The default length is 20B.

  &zeroWidthSpace; `gzipMinLength = 256` 仅当长度未知或大于 gzipMinLength 时，才会压缩原始内容。默认长度为 20B。

  `includedMethods = get;post` List of HTTP methods to compress. By default only GET requests are compressed.

  &zeroWidthSpace; `includedMethods = get;post` 要压缩的 HTTP 方法列表。默认情况下，仅压缩 GET 请求。

- MaxMemory

  Sets the memory cache size for file uploading. By default this is `1 << 26`(64M).

  &zeroWidthSpace;设置文件上传的内存缓存大小。默认值为 `1 << 26` (64M)。

  `beego.BConfig.MaxMemory = 1 << 26`

- EnableErrorsShow

  Toggles the display of error messages. By default this is True.

  &zeroWidthSpace;切换错误消息的显示。默认值为 True。

  `beego.BConfig.EnableErrorsShow = true`

- EnableErrorsRender

  Toggles rendering error messages. By default this is set to True. User friendly error pages will not be rendered even in dev `RunMode` if this value is false.

  &zeroWidthSpace;切换渲染错误消息。默认设置为 True。如果此值为 false，即使在 dev `RunMode` 中也不会渲染用户友好的错误页面。

#### Web config Web 配置

- AutoRender

  Enable auto render. By default this is True. This value should be set to false for API applications, as there is no need to render templates.

  &zeroWidthSpace;启用自动渲染。默认值为 True。对于 API 应用程序，此值应设置为 false，因为无需渲染模板。

  `beego.BConfig.WebConfig.AutoRender = true`

- EnableDocs

  Enable Docs. By default this is False.

  &zeroWidthSpace;启用文档。默认情况下为 False。

  `beego.BConfig.WebConfig.EnableDocs = false`

- FlashName

  Sets the Flash Cookie name. By default this is `BEEGO_FLASH`.

  &zeroWidthSpace;设置 Flash Cookie 名称。默认情况下为 `BEEGO_FLASH` 。

  `beego.BConfig.WebConfig.FlashName = "BEEGO_FLASH"`

- FlashSeperator

  Set the Flash data separator. By default this is `BEEGOFLASH`.

  &zeroWidthSpace;设置 Flash 数据分隔符。默认情况下为 `BEEGOFLASH` 。

  `beego.BConfig.WebConfig.FlashSeperator = "BEEGOFLASH"`

- DirectoryIndex

  Enable listing of the static directory. By default this is False and will return a 403 error.

  &zeroWidthSpace;启用静态目录的列表。默认情况下为 False，并将返回 403 错误。

  `beego.BConfig.WebConfig.DirectoryIndex = false`

- StaticDir

  Sets the static file dir(s). By default this is `static`.

  &zeroWidthSpace;设置静态文件目录。默认情况下为 `static` 。

  1. Single dir, `StaticDir = download`. Same as `beego.SetStaticPath("/download","download")`
     单个目录， `StaticDir = download` 。与 `beego.SetStaticPath("/download","download")`
  2. Multiple dirs, `StaticDir = download:down download2:down2`. Same as `beego.SetStaticPath("/download","down")` and `beego.SetStaticPath("/download2","down2")`

  
  &zeroWidthSpace;2. 多个目录， `StaticDir = download:down download2:down2` 。与 `beego.SetStaticPath("/download","down")` 和 `beego.SetStaticPath("/download2","down2")` 相同

  `beego.BConfig.WebConfig.StaticDir = map[string]string{"download":"download"}`

- StaticExtensionsToGzip

  Sets a list of file extensions that will support compression by Gzip. The formats `.css` and `.js` are supported by default.

  &zeroWidthSpace;设置支持 Gzip 压缩的文件扩展名列表。默认情况下支持 `.css` 和 `.js` 格式。

  `beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}`

  Same as in config file StaticExtensionsToGzip = .css, .js

  &zeroWidthSpace;与配置文件中的 StaticExtensionsToGzip = .css, .js 相同

- TemplateLeft

  Left mark of the template, `{{` by default.

  &zeroWidthSpace;模板的左标记，默认情况下为 `{{` 。

  `beego.BConfig.WebConfig.TemplateLeft = "{{"`

- TemplateRight

  Right mark of the template, `}}` by default.

  &zeroWidthSpace;模板的右标记，默认情况下为 `}}` 。

  `beego.BConfig.WebConfig.TemplateRight = "}}"`

- ViewsPath

  Set the location of template files. This is set to `views` by default.

  &zeroWidthSpace;设置模板文件的位置。默认情况下设置为 `views` 。

  `beego.BConfig.WebConfig.ViewsPath = "views"`

- EnableXSRF Enable XSRF

  &zeroWidthSpace;EnableXSRF 启用 XSRF

  `beego.BConfig.WebConfig.EnableXSRF = false`

- XSRFKEY

  Set the XSRF key. By default this is `beegoxsrf`.

  &zeroWidthSpace;设置 XSRF 密钥。默认情况下，这是 `beegoxsrf` 。

  `beego.BConfig.WebConfig.XSRFKEY = "beegoxsrf"`

- XSRFExpire

  Set the XSRF expire time. By default this is set to `0`.

  &zeroWidthSpace;设置 XSRF 过期时间。默认情况下，将其设置为 `0` 。

  `beego.BConfig.WebConfig.XSRFExpire = 0`

- CommentRouterPath

  Beego scan `CommentRouterPath` to auto generate router, the default value is `controllers`。 `beego.BConfig.WebConfig.CommentRouterPath = "controllers"`

  &zeroWidthSpace;Beego 扫描 `CommentRouterPath` 以自动生成路由器，默认值为 `controllers` 。 `beego.BConfig.WebConfig.CommentRouterPath = "controllers"`

#### HTTP Server config HTTP 服务器配置

- Graceful

  Enable graceful shutdown. By default this is False.

  &zeroWidthSpace;启用优雅关机。默认情况下，这是 False。

  `beego.BConfig.Listen.Graceful = false`

- ServerTimeOut

  Set the http timeout. By default thi is ‘0’, no timeout.

  &zeroWidthSpace;设置 http 超时时间。默认情况下为“0”，无超时时间。

  `beego.BConfig.Listen.ServerTimeOut = 0`

- ListenTCP4

  Set the address type. default is `tcp6` but we can set it to true to force use `TCP4`.

&zeroWidthSpace;设置地址类型。默认值为 `tcp6` ，但我们可以将其设置为 true 以强制使用 `TCP4` 。

```
`beego.BConfig.Listen.ListenTCP4 = true`
```

- EnableHTTP

  Enable HTTP listen. By default this is set to True.

  &zeroWidthSpace;启用 HTTP 监听。默认情况下，此项设置为 True。

  `beego.BConfig.Listen.EnableHTTP = true`

- HTTPAddr

  Set the address the app listens to. By default this value is empty and the app will listen to all IPs.

  &zeroWidthSpace;设置应用程序监听的地址。默认情况下，此值为空，应用程序将监听所有 IP。

  `beego.BConfig.Listen.HTTPAddr = ""`

- HTTPPort

  Set the port the app listens on. By default this is 8080

  &zeroWidthSpace;设置应用程序监听的端口。默认情况下，此端口为 8080

  `beego.BConfig.Listen.HTTPPort = 8080`

- EnableHTTPS

  Enable HTTPS. By default this is False. When enabled `HTTPSCertFile` and `HTTPSKeyFile` must also be set.

  &zeroWidthSpace;启用 HTTPS。默认情况下，此值为 False。启用时，还必须设置 `HTTPSCertFile` 和 `HTTPSKeyFile` 。

  `beego.BConfig.Listen.EnableHTTPS = false`

- HTTPSAddr

  Set the address the app listens to. Default is empty and the app will listen to all IPs.

  &zeroWidthSpace;设置应用监听的地址。默认值为空，应用将监听所有 IP。

  `beego.BConfig.Listen.HTTPSAddr = ""`

- HTTPSPort

  Set the port the app listens on. By default this is 10443

  &zeroWidthSpace;设置应用监听的端口。默认值为 10443

  `beego.BConfig.Listen.HTTPSPort = 10443`

- HTTPSCertFile

  Set the SSL cert path. By default this value is empty.

  &zeroWidthSpace;设置 SSL 证书路径。默认情况下，此值为空。

  `beego.BConfig.Listen.HTTPSCertFile = "conf/ssl.crt"`

- HTTPSKeyFile

  Set the SSL key path. By default this value is empty.

  &zeroWidthSpace;设置 SSL 密钥路径。默认情况下，此值为空。

  `beego.BConfig.Listen.HTTPSKeyFile = "conf/ssl.key"`

- EnableAdmin

  Enable supervisor module. By default this is False.

  &zeroWidthSpace;启用管理员模块。默认情况下，此项为 False。

  `beego.BConfig.Listen.EnableAdmin = false`

- AdminAddr

  Set the address the admin app listens to. By default this is blank and the app will listen to any IP.

  &zeroWidthSpace;设置管理员应用程序侦听的地址。默认情况下，此项为空，应用程序将侦听任何 IP。

  `beego.BConfig.Listen.AdminAddr = ""`

- AdminPort

  Set the port the admin app listens on. By default this is 8088.

  &zeroWidthSpace;设置管理员应用程序侦听的端口。默认情况下，此项为 8088。

  `beego.BConfig.Listen.AdminPort = 8088`

- EnableFcgi

  Enable fastcgi. By default this is False.

  &zeroWidthSpace;启用 fastcgi。默认情况下，此项为 False。

  `beego.BConfig.Listen.EnableFcgi = false`

- EnableStdIo

  Enable fastcgi standard I/O or not. By default this is False.

  &zeroWidthSpace;启用或不启用 fastcgi 标准 I/O。默认情况下，此项为 False。

  `beego.BConfig.Listen.EnableStdIo = false`

#### Session config 会话配置

- SessionOn

  Enable session. By default this is False.

  &zeroWidthSpace;启用会话。默认情况下，这是 False。

  `beego.BConfig.WebConfig.Session.SessionOn = false`

- SessionProvider

  Set the session provider. By default this is `memory`.

  &zeroWidthSpace;设置会话提供程序。默认情况下，这是 `memory` 。

  `beego.BConfig.WebConfig.Session.SessionProvider = "memory"`

- SessionName

  Set the session cookie name stored in the browser. By default this is `beegosessionID`.

  &zeroWidthSpace;设置存储在浏览器中的会话 Cookie 名称。默认情况下，这是 `beegosessionID` 。

  `beego.BConfig.WebConfig.Session.SessionName = "beegosessionID"`

- SessionGCMaxLifetime

  Set the session expire time. By default this is 3600s.

  &zeroWidthSpace;设置会话过期时间。默认情况下，这是 3600 秒。

  `beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600`

- SessionProviderConfig

  Set the session provider config. Different providers can require different config settings. Please see [session](https://beego.wiki/docs/module/session) for more information.

  &zeroWidthSpace;设置会话提供程序配置。不同的提供程序可能需要不同的配置设置。有关更多信息，请参阅会话。

- SessionCookieLifeTime

  Set the valid expiry time of the cookie in browser for session. By default this is 3600s.

  &zeroWidthSpace;设置会话在浏览器中 cookie 的有效过期时间。默认情况下，这是 3600 秒。

  `beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600`

- SessionAutoSetCookie

  Enable SetCookie. By default this is True.

  &zeroWidthSpace;启用 SetCookie。默认情况下，这是 True。

  `beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true`

- SessionDomain

  Set the session cookie domain. By default this is empty.

  &zeroWidthSpace;设置会话 cookie 域。默认情况下，这是空的。

  `beego.BConfig.WebConfig.Session.SessionDomain = ""`

#### Log config 日志配置

```
See [logs module](en-US/module/logs) for more information.
```

- AccessLogs

  Enable output access logs. By default these logs will not be output under ‘prod’ mode.

  &zeroWidthSpace;启用输出访问日志。默认情况下，这些日志不会在“prod”模式下输出。

  `beego.BConfig.Log.AccessLogs = false`

- FileLineNum

  Toggle printing line numbers. By default this is True. This config is not supported in config file.

  &zeroWidthSpace;切换打印行号。默认情况下为 True。此配置不受配置文件支持。

  `beego.BConfig.Log.FileLineNum = true`

- Outputs

  &zeroWidthSpace;输出

  Log outputs config. This config is not supported in config file.

  &zeroWidthSpace;日志输出配置。此配置不受配置文件支持。

  `beego.BConfig.Log.Outputs = map[string]string{"console": ""}`

  or

  `beego.BConfig.Log.Outputs["console"] = ""`