+++
title = "会话控制"
date = 2024-02-04T09:57:58+08:00
weight = 7
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/session/]({{< ref "/beego/mvcIntroduction/controllers/sessionControl" >}})

# Session control 会话控制



## Session control 会话控制

Beego has a built-in session module that supports memory, file, mysql, redis, couchbase, memcache and postgres as the save provider. Other providers can be implemented according to the interface.

​	Beego 具有一个内建的会话模块，支持内存、文件、数据库、redis、couchbase、memcache 和 postgre 作为保存方式。其他方式可以根据接口来定。

To use session in Beego switch it on in the main function:

​	在 Go 中使用会话时，在主函数中打开它：

```
web.BConfig.WebConfig.Session.SessionOn = true
```

Or it can be activated in the configuration file:

​	或者它可以在配置文件中：

```
SessionOn = true
```

After being switched on, session can used be used like this:

​	打开后，会话可以像这样使用：

```go
func (this *MainController) Get() {
	v := this.GetSession("asta")
	if v == nil {
		this.SetSession("asta", int(1))
		this.Data["num"] = 0
	} else {
		this.SetSession("asta", v.(int)+1)
		this.Data["num"] = v.(int)
	}
	this.TplName = "index.tpl"
}
```

There are several useful methods to handle session:

​	有几个有用的方法来处理会话：

- SetSession(name string, value interface{})
  SetSession(name string, value interface)
- GetSession(name string) interface{}
- DelSession(name string)
- SessionRegenerateID()
- DestroySession()
  Session()

The most commonly used methods are `SetSession`, `GetSession`, and `DelSession`.

​	最常用的方法是 `SetSession` 、 `GetSession` 和 `DelSession` 。

Custom logic can also be used:

​	也可以使用自定义逻辑：

```
sess := this.StartSession()
defer sess.SessionRelease()
```

sess object has following methods:

​	sess 对象具有以下方法：

- Set
- Get
- Delete
- SessionID
- SessionRelease
- Flush

SetSession, GetSession and DelSession methods are recommended for session operation as it will release resource automatically.

​	建议使用 SetSession、GetSession 和 DelSession 方法进行会话操作，因为它会自动释放资源。

Here are some parameters used in the Session module:

​	以下是 Session 模块中使用的一些参数：

- SessionOn

  Enables Session. Default value is `false`. Parameter name in configuration file: `SessionOn`

  ​	启用 Session。默认值为 `false` 。配置文件中的参数名称： `SessionOn`

- SessionProvider Sets Session provider. Set to `memory` by default. `File`, `mysql` and `redis` are also supported. Parameter name in configuration file: `sessionprovider`.
  SessionProvider 设置会话提供程序。默认设置为 `memory` 。还支持 `File` 、 `mysql` 和 `redis` 。配置文件中的参数名称： `sessionprovider` 。

- SessionName Sets the cookie name. Session is stored in browser’s cookies by default. The default name is beegosessionID. Parameter name in configuration file: `sessionname`.
  SessionName 设置 cookie 名称。默认情况下，会话存储在浏览器的 cookie 中。默认名称为 beegosessionID。配置文件中的参数名称： `sessionname` 。

- SessionGCMaxLifetime Sets the Session expire time. Default value is `3600s`. Parameter name in configuration file: `sessiongcmaxlifetime`.
  SessionGCMaxLifetime 设置会话过期时间。默认值为 `3600s` 。配置文件中的参数名称： `sessiongcmaxlifetime` 。

- SessionProviderConfig Sets the save path or connection string for file, mysql or redis. Default value is empty. Parameter name in configuration file: `sessionproviderconfig`.
  SessionProviderConfig 设置文件、mysql 或 redis 的保存路径或连接字符串。默认值为空。配置文件中的参数名称： `sessionproviderconfig` 。

- SessionHashFunc Sets the function used to generate sessionid. The default value is `sha1`.
  SessionHashFunc 设置用于生成 sessionid 的函数。默认值为 `sha1` 。

- SessionCookieLifeTime Sets the cookie expire time. The cookie is used to store data in client.
  SessionCookieLifeTime 设置 cookie 过期时间。cookie 用于在客户端存储数据。

## Package Installation 软件包安装

If you are not using Go modules, manual installation may be required.
如果您未使用 Go 模块，可能需要手动安装。

*Note: Beego >= 1.1.3 removed all dependencies
注意：Beego >= 1.1.3 已删除所有依赖项*

```bash
# Couchbase
go get -u github.com/beego/beego/v2/server/web/session/couchbase

# Ledis
go get -u github.com/beego/beego/v2/server/web/session/ledis

# Memcache
go get -u github.com/beego/beego/v2/server/web/session/memcache

# MySQL
go get -u github.com/beego/beego/v2/server/web/session/mysql

# Postgres
go get -u github.com/beego/beego/v2/server/web/session/postgres

# Redis
go get -u github.com/beego/beego/v2/server/web/session/redis

# Redis (cluster mode)
go get -u github.com/beego/beego/v2/server/web/session/redis_cluster

# Redis (sentinel)
go get -u github.com/beego/beego/v2/server/web/session/redis_sentinel

# SSDB
go get -u github.com/beego/beego/v2/server/web/session/ssdb
```

## Example Usage 示例用法

### Couchbase

SessionProviderConfig is connection address using [couchbase](https://github.com/couchbaselabs/go-couchbase).
SessionProviderConfig 是使用 couchbase 的连接地址。

```go
// main.go
package main

import (
  "github.com/beego/beego/v2/server/web"
  _ "github.com/beego/beego/v2/server/web/session/couchbase"
)

func init() {
  web.BConfig.WebConfig.Session.SessionOn = true
  web.BConfig.WebConfig.Session.SessionProvider = "couchbase"
  web.BConfig.WebConfig.Session.SessionProviderConfig = "http://bucketname:bucketpass@myserver:8091/"
}
```

### File

```go
// main.go
package main

import (
  "github.com/beego/beego/v2/server/web"
)

func init() {
  web.BConfig.WebConfig.Session.SessionOn = true
  web.BConfig.WebConfig.Session.SessionProvider = "file"
  web.BConfig.WebConfig.Session.SessionProviderConfig = "/tmp"
}
```

### Memcache

SessionProviderConfig is the connection address using [memcache](https://github.com/beego/memcache).
SessionProviderConfig 是使用 memcache 的连接地址。

```go
// main.go
package main

import (
  "github.com/beego/beego/v2/server/web"
  _ "github.com/beego/beego/v2/server/web/session/memcache"
)

func init() {
  web.BConfig.WebConfig.Session.SessionOn = true
  web.BConfig.WebConfig.Session.SessionProvider = "memcache"
  web.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:7080"
}
```

### MySQL

SessionProviderConfig is the connection address using [go-sql-driver](https://github.com/go-sql-driver/mysql).
SessionProviderConfig 是使用 go-sql-driver 的连接地址。

```go
// main.go
package main

import (
  "github.com/beego/beego/v2/server/web"
  _ "github.com/beego/beego/v2/server/web/session/mysql"
)

func init() {
  web.BConfig.WebConfig.Session.SessionOn = true
  web.BConfig.WebConfig.Session.SessionProvider = "mysql"
  web.BConfig.WebConfig.Session.SessionProviderConfig = "username:password@protocol(address)/dbname?param=value"
}
```

### Postgres

SessionProviderConfig is the connection address using [postgres](https://github.com/lib/pq).
SessionProviderConfig 是使用 postgres 的连接地址。

```go
// main.go
package main

import (
  "github.com/beego/beego/v2/server/web"
  _ "github.com/beego/beego/v2/server/web/session/postgres"
)

func init() {
  web.BConfig.WebConfig.Session.SessionOn = true
  web.BConfig.WebConfig.Session.SessionProvider = "postgresql"
  web.BConfig.WebConfig.Session.SessionProviderConfig = "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
}
```

### Redis

SessionProviderConfig is the connection address using [redigo](https://github.com/garyburd/redigo).
SessionProviderConfig 是使用 redigo 的连接地址。

```go
// main.go
package main

import (
  "github.com/beego/beego/v2/server/web"
  _ "github.com/beego/beego/v2/server/web/session/redis"
)

func init() {
  web.BConfig.WebConfig.Session.SessionOn = true
  web.BConfig.WebConfig.Session.SessionProvider = "redis"
  web.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
}
```

## Note: 注意：

Session uses `gob` to register objects. When using a session engine other than `memory`, objects must be registered in session before they can be used. Use `gob.Register()` to register them in `init()` function.
会话使用 `gob` 来注册对象。当使用除 `memory` 之外的会话引擎时，必须在会话中注册对象才能使用它们。使用 `gob.Register()` 在 `init()` 函数中注册它们。
