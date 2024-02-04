+++
title = "session 模块"
date = 2024-02-04T09:25:40+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/module/session/]({{< ref "/beego/modules/session" >}})

# Session Module 会话模块



> Notes: 备注：
>
> - This document is for using `session` as a standalone module in other projects. If you are using `session` with Beego, please check here [session control]({{< ref "/beego/mvcIntroduction/controllers/sessionControl" >}})*
>   本文档用于在其他项目中将 `session` 作为独立模块使用。如果您将 `session` 与 Beego 一起使用，请在此处查看会话控制*

## Introduction to Session Module 会话模块简介

The session module is used to store user data between different requests. It only supports saving the session id into a cookie, so if the client doesn’t support cookies, it won’t work.

​	会话模块用于在不同请求之间存储用户数据。它仅支持将会话 ID 保存到 cookie 中，因此如果客户端不支持 cookie，则它将不起作用。

It is inspired by `database/sql`, which means: one interface, multiple implementations. By default it supports four saving providers: memory, file, redis and mysql.

​	它受 `database/sql` 的启发，这意味着：一个接口，多种实现。默认情况下，它支持四种保存提供程序：内存、文件、redis 和 mysql。

Install session module: 
​	安装会话模块：

```
go get github.com/beego/beego/v2/server/web/session
```

## Basic Usage: 基本用法：

Import package first: 
​	首先导入包：

```
import (
	"github.com/beego/beego/v2/server/web/session"
)
```

Then initialize a global variable as the session manager:

​	然后初始化一个全局变量作为会话管理器：

```
var globalSessions *session.Manager
```

Then initialize data in your main function:

​	然后在主函数中初始化数据：

```
func init() {
	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
	go globalSessions.GC()
}
```

Parameters of NewManager:

​	NewManager 的参数：

1. Saving provider name: memory, file, mysql, redis
   保存提供程序名称：memory、file、mysql、redis

2. A JSON string that contains the config information.

   
   包含配置信息的 JSON 字符串。

   1. cookieName: Cookie name of session id saved on the client
      cookieName：保存在客户端的会话 ID 的 Cookie 名称
   2. enableSetCookie, omitempty: Whether to enable SetCookie, omitempty
      enableSetCookie, omitempty：是否启用 SetCookie，omitempty
   3. gclifetime: The interval of GC.
      gclifetime：GC 的间隔。
   4. maxLifetime: Expiration time of data saved on the server
      maxLifetime: 服务器上保存的数据的过期时间
   5. secure: Enable https or not. There is `cookie.Secure` while configure cookie.
      secure: 是否启用 https。在配置 cookie 时有 `cookie.Secure` 。
   6. sessionIDHashFunc: SessionID generator function. `sha1` by default.
      sessionIDHashFunc: SessionID 生成器函数。默认情况下为 `sha1` 。
   7. sessionIDHashKey: Hash key.
      sessionIDHashKey: 哈希密钥。
   8. cookieLifeTime: Cookie expiration time on the client. 0 by default, which means life time of browser.
      cookieLifeTime: 客户端上的 Cookie 过期时间。默认值为 0，表示浏览器的生命周期。
   9. providerConfig: Provider-specific config. See below for more information.
      providerConfig: 特定于提供程序的配置。有关更多信息，请参见下文。

Then we can use session in our code:

​	然后我们可以在代码中使用 session：

```
func login(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	username := sess.Get("username")
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		sess.Set("username", r.Form["username"])
	}
}
```

Here are the methods of globalSessions:

​	以下是 globalSessions 的方法：

- `SessionStart` Return session object based on current request.
  `SessionStart` 根据当前请求返回会话对象。
- `SessionDestroy` Destroy current session object.
  `SessionDestroy` 销毁当前会话对象。
- `SessionRegenerateId` Regenerate a new sessionID.
  `SessionRegenerateId` 重新生成新的 sessionID。
- `GetActiveSession` Get active session user.
  `GetActiveSession` 获取活动会话用户。
- `SetHashFunc` Set sessionID generator function.
  `SetHashFunc` 设置 sessionID 生成器函数。
- `SetSecure` Enable Secure cookie or not.
  `SetSecure` 启用安全 cookie 或不启用。

The returned session object is a Interface. Here are the methods:

​	返回的会话对象是一个接口。以下是方法：

- `Set(key, value interface{}) error`
- `Get(key interface{}) interface{}`
- `Delete(key interface{}) error`
- `SessionID() string`
- `SessionRelease()`
- `Flush() error`

## Saving Provider Config 保存提供程序配置

We’ve already seen configuration of `memory` provider. Here is the configuration of the others:

​	我们已经看到了 `memory` 提供程序的配置。以下是其他配置：

- `mysql`:

  All the parameters are the same as memory’s except the fourth parameter, e.g.:

  ​	所有参数都与内存相同，除了第四个参数，例如：

  ```
    username:password@protocol(address)/dbname?param=value
  ```

  For details see the [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql#dsn-data-source-name) documentation.

  ​	有关详细信息，请参阅 go-sql-driver/mysql 文档。

- `redis`:

  Connection config: address,pool,password

  ​	连接配置：地址、池、密码

  ```
    127.0.0.1:6379,100,astaxie
  ```

- `file`:

  The session save path. Create new files in two levels by default. E.g.: if sessionID is `xsnkjklkjjkh27hjh78908` the file will be saved as `./tmp/x/s/xsnkjklkjjkh27hjh78908`

  ​	会话保存路径。默认情况下，在两级目录中创建新文件。例如：如果 sessionID 为 `xsnkjklkjjkh27hjh78908` ，则文件将保存为 `./tmp/x/s/xsnkjklkjjkh27hjh78908`

  ```
    ./tmp
  ```

## Creating a new provider 创建新提供程序[ ](https://beego.wiki/docs/module/session/#creating-a-new-provider)

Sometimes you need to create your own session provider. The Session module uses interfaces, so you can implement this interface to create your own provider easily.

​	有时您需要创建自己的会话提供程序。Session 模块使用接口，因此您可以实现此接口以轻松创建自己的提供程序。

```go
// Store contains all data for one session process with specific id.
type Store interface {
	Set(ctx context.Context, key, value interface{}) error     //set session value
	Get(ctx context.Context, key interface{}) interface{}      //get session value
	Delete(ctx context.Context, key interface{}) error         //delete session value
	SessionID(ctx context.Context) string                      //back current sessionID
	SessionRelease(ctx context.Context, w http.ResponseWriter) // release the resource & save data to provider & return the data
	Flush(ctx context.Context) error                           //delete all data
}

// Provider contains global session methods and saved SessionStores.
// it can operate a SessionStore by its id.
type Provider interface {
	SessionInit(ctx context.Context, gclifetime int64, config string) error
	SessionRead(ctx context.Context, sid string) (Store, error)
	SessionExist(ctx context.Context, sid string) (bool, error)
	SessionRegenerate(ctx context.Context, oldsid, sid string) (Store, error)
	SessionDestroy(ctx context.Context, sid string) error
	SessionAll(ctx context.Context) int // get all active session
	SessionGC(ctx context.Context)
}
```

Finally, register your provider:

​	最后，注册您的提供程序：

```
func init() {
	// ownadapter is an instance of session.Provider
	session.Register("own", ownadapter)
}
```
