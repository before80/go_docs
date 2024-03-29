+++
title = "gpool"
date = 2024-03-21T17:44:43+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：

Package gpool provides object-reusable concurrent-safe pool.

​	软件包 gpool 提供对象可重用的并发安全池。

## 常量

This section is empty.

## 变量

This section is empty.

## 函数

This section is empty.

## 类型

### type ExpireFunc

```go
type ExpireFunc func(interface{})
```

ExpireFunc Destruction function for object.

​	ExpireFunc 对象的销毁函数。

### type NewFunc

```go
type NewFunc func() (interface{}, error)
```

NewFunc Creation function for object.

​	NewFunc 对象的创建函数。

### type Pool

```go
type Pool struct {
	TTL     time.Duration               // Time To Live for pool items.
	NewFunc func() (interface{}, error) // Callback function to create pool item.
	// ExpireFunc is the function for expired items destruction.
	// This function needs to be defined when the pool items
	// need to perform additional destruction operations.
	// Eg: net.Conn, os.File, etc.
	ExpireFunc func(interface{})
	// contains filtered or unexported fields
}
```

Pool is an Object-Reusable Pool.

​	Pool 是一个对象可重用的池。

#### func New

```go
func New(ttl time.Duration, newFunc NewFunc, expireFunc ...ExpireFunc) *Pool
```

New creates and returns a new object pool. To ensure execution efficiency, the expiration time cannot be modified once it is set.

​	New 创建并返回一个新的对象池。为保证执行效率，过期时间一经设置，无法修改。

Note the expiration logic: ttl = 0 : not expired; ttl < 0 : immediate expired after use; ttl > 0 : timeout expired;

​	注意过期逻辑：ttl = 0 ： 未过期;ttl < 0：使用后立即过期; ttl > 0：超时过期;

##### Example

``` go
```

#### (*Pool) Clear

```go
func (p *Pool) Clear()
```

Clear clears pool, which means it will remove all items from pool.

​	清除池，这意味着它将从池中删除所有项目。

##### Example

``` go
```

#### (*Pool) Close

```go
func (p *Pool) Close()
```

Close closes the pool. If `p` has ExpireFunc, then it automatically closes all items using this function before it’s closed. Commonly you do not need to call this function manually.

​	关闭 关闭池。如果 `p` 具有 ExpireFunc，则它会在关闭之前使用此函数自动关闭所有项目。通常，不需要手动调用此函数。

##### Example

``` go
```

#### (*Pool) Get

```go
func (p *Pool) Get() (interface{}, error)
```

Get picks and returns an item from pool. If the pool is empty and NewFunc is defined, it creates and returns one from NewFunc.

​	从池中获取拣选并返回项目。如果池为空且定义了 NewFunc，则会从 NewFunc 创建并返回一个池。

##### Example

``` go
```

#### (*Pool) MustPut

```go
func (p *Pool) MustPut(value interface{})
```

MustPut puts an item to pool, it panics if any error occurs.

​	MustPut 将项目放入池中，如果发生任何错误，它会崩溃。

#### (*Pool) Put

```go
func (p *Pool) Put(value interface{}) error
```

Put puts an item to pool.

​	Put 将项目放入池中。

##### Example

``` go
```

#### (*Pool) Size

```go
func (p *Pool) Size() int
```

Size returns the count of available items of pool.

​	Size 返回池中可用项的计数。

##### Example

``` go
```

```go
package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/container/gpool"
)

func main() {
	type DBConn struct {
		Conn  *sql.Conn
		Limit int
	}

	dbConnPool := gpool.New(time.Hour,
		func() (interface{}, error) {
			dbConn := new(DBConn)
			dbConn.Limit = 10
			return dbConn, nil
		},
		func(i interface{}) {
			// sample : close db conn
			// i.(DBConn).Conn.Close()
		})

	conn, _ := dbConnPool.Get()
	fmt.Println(dbConnPool.Size())
	dbConnPool.MustPut(conn)
	dbConnPool.MustPut(conn)
	fmt.Println(dbConnPool.Size())

}

Output:

0
2
```







<iframe allowtransparency="true" frameborder="0" scrolling="no" class="sk_ui" src="chrome-extension://gfbliohnnapiefjpjlpjnehglfpaknnc/pages/frontend.html" title="Surfingkeys" style="left: 0px; bottom: 0px; width: 1555px; height: 0px; z-index: 2147483647;"></iframe>