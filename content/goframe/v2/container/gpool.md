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

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type ExpireFunc 

``` go
type ExpireFunc func(interface{})
```

ExpireFunc Destruction function for object.

#### type NewFunc 

``` go
type NewFunc func() (interface{}, error)
```

NewFunc Creation function for object.

#### type Pool 

``` go
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

##### func New 

``` go
func New(ttl time.Duration, newFunc NewFunc, expireFunc ...ExpireFunc) *Pool
```

New creates and returns a new object pool. To ensure execution efficiency, the expiration time cannot be modified once it is set.

Note the expiration logic: ttl = 0 : not expired; ttl < 0 : immediate expired after use; ttl > 0 : timeout expired;

##### Example

``` go
```
##### (*Pool) Clear 

``` go
func (p *Pool) Clear()
```

Clear clears pool, which means it will remove all items from pool.

##### Example

``` go
```
##### (*Pool) Close 

``` go
func (p *Pool) Close()
```

Close closes the pool. If `p` has ExpireFunc, then it automatically closes all items using this function before it's closed. Commonly you do not need to call this function manually.

##### Example

``` go
```
##### (*Pool) Get 

``` go
func (p *Pool) Get() (interface{}, error)
```

Get picks and returns an item from pool. If the pool is empty and NewFunc is defined, it creates and returns one from NewFunc.

##### Example

``` go
```
##### (*Pool) MustPut <-2.3.0

``` go
func (p *Pool) MustPut(value interface{})
```

MustPut puts an item to pool, it panics if any error occurs.

##### (*Pool) Put 

``` go
func (p *Pool) Put(value interface{}) error
```

Put puts an item to pool.

##### Example

``` go
```
##### (*Pool) Size 

``` go
func (p *Pool) Size() int
```

Size returns the count of available items of pool.



Example 

``` go
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