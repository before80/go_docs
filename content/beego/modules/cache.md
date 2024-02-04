+++
title = "cache 模块"
date = 2024-02-04T09:31:15+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/module/cache/](https://beego.wiki/docs/module/cache/)

# Cache Module 缓存模块

Beego’s cache module is used for caching data, inspired by `database/sql`. It supports four cache providers: file, memcache, memory and redis. You can install it by:

&zeroWidthSpace;Beego 的缓存模块用于缓存数据，灵感来自 `database/sql` 。它支持四种缓存提供程序：文件、memcache、内存和 redis。您可以通过以下方式安装它：

```
github.com/beego/beego/v2/client/cache
```

If you use the `memcache` or `redis` provider, you should first install:

&zeroWidthSpace;如果您使用 `memcache` 或 `redis` 提供程序，则应首先安装：

```
go get -u github.com/beego/beego/v2/client/cache/memcache
```

and then import: 
&zeroWidthSpace;然后导入：

```
import _ "github.com/beego/beego/v2/client/cache/memcache"
```

## Basic Usage 基本用法

First step is importing the package:

&zeroWidthSpace;第一步是导入包：

```
import (
	"github.com/beego/beego/v2/client/cache"
)
```

Then initialize a global variable object:

&zeroWidthSpace;然后初始化一个全局变量对象：

```
bm, err := cache.NewCache("memory", `{"interval":60}`)
```

Then we can use `bm` to modify the cache:

&zeroWidthSpace;然后我们可以使用 `bm` 来修改缓存：

```
bm.Put("astaxie", 1, 10*time.Second)
bm.Get("astaxie")
bm.IsExist("astaxie")
bm.Delete("astaxie")
```

## Provider Settings 提供程序设置

Here is how to configure the four providers:

&zeroWidthSpace;以下是配置四个提供程序的方法：

- memory

  `interval` stands for GC time, which means the cache will be cleared every 60s:

  &zeroWidthSpace; `interval` 代表 GC 时间，这意味着缓存每 60 秒清除一次：

  ```
    {"interval":60}
  ```

- file

  ```
    {"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}
  ```

- redis

  redis uses [redigo](https://github.com/garyburd/redigo/tree/master/redis) 
  &zeroWidthSpace;redis 使用 redigo

  ```
    {"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}
  ```

  - key: the Redis collection name
    key：Redis 集合名称
  - conn: Redis connection info
    conn：Redis 连接信息
  - dbNum: Select the DB having the specified zero-based numeric index.
    dbNum：选择具有指定基于零的数字索引的数据库。
  - password: the password for connecting password-protected Redis server
    password：用于连接受密码保护的 Redis 服务器的密码

- memcache

  memcache uses [vitess](http://code.google.com/p/vitess/go/memcache) 
  &zeroWidthSpace;memcache 使用 vitess

  ```
    {"conn":"127.0.0.1:11211"}
  ```

## Creating your own provider 创建您自己的提供程序

The cache module uses the Cache interface, so you can create your own cache provider by implementing this interface and registering it.

&zeroWidthSpace;缓存模块使用 Cache 接口，因此您可以通过实现此接口并注册它来创建自己的缓存提供程序。

```go
type Cache interface {
	// Get a cached value by key.
	Get(ctx context.Context, key string) (interface{}, error)
	// GetMulti is a batch version of Get.
	GetMulti(ctx context.Context, keys []string) ([]interface{}, error)
	// Set a cached value with key and expire time.
	Put(ctx context.Context, key string, val interface{}, timeout time.Duration) error
	// Delete cached value by key.
	Delete(ctx context.Context, key string) error
	// Increment a cached int value by key, as a counter.
	Incr(ctx context.Context, key string) error
	// Decrement a cached int value by key, as a counter.
	Decr(ctx context.Context, key string) error
	// Check if a cached value exists or not.
	IsExist(ctx context.Context, key string) (bool, error)
	// Clear all cache.
	ClearAll(ctx context.Context) error
	// Start gc routine based on config string settings.
	StartAndGC(config string) error
}
```

Register your provider:

&zeroWidthSpace;注册您的提供程序：

```
func init() {
	cache.Register("myowncache", NewOwnCache())
}
```