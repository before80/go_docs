+++
title = "gfpool"
date = 2024-03-21T17:55:35+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gfpool

Package gfpool provides io-reusable pool for file pointer.

### Constants 

This section is empty.

### Variables 

This section is empty.

### Functions 

This section is empty.

### Types 

#### type File 

``` go
type File struct {
	*os.File // Underlying file pointer.
	// contains filtered or unexported fields
}
```

File is an item in the pool.

##### func Get <-2.1.1

``` go
func Get(path string, flag int, perm os.FileMode, ttl ...time.Duration) (file *File)
```

Get returns a file item with given file path, flag and opening permission. It retrieves a file item from the file pointer pool after then.

##### func Open 

``` go
func Open(path string, flag int, perm os.FileMode, ttl ...time.Duration) (file *File, err error)
```

Open creates and returns a file item with given file path, flag and opening permission. It automatically creates an associated file pointer pool internally when it's called first time. It retrieves a file item from the file pointer pool after then.

##### (*File) Close 

``` go
func (f *File) Close(close ...bool) error
```

Close puts the file pointer back to the file pointer pool.

##### (*File) Stat 

``` go
func (f *File) Stat() (os.FileInfo, error)
```

Stat returns the FileInfo structure describing file.

#### type Pool 

``` go
type Pool struct {
	// contains filtered or unexported fields
}
```

Pool pointer pool.

##### func New 

``` go
func New(path string, flag int, perm os.FileMode, ttl ...time.Duration) *Pool
```

New creates and returns a file pointer pool with given file path, flag and opening permission.

Note the expiration logic: ttl = 0 : not expired; ttl < 0 : immediate expired after use; ttl > 0 : timeout expired; It is not expired in default.

##### (*Pool) Close 

``` go
func (p *Pool) Close()
```

Close closes current file pointer pool.

##### (*Pool) File 

``` go
func (p *Pool) File() (*File, error)
```

File retrieves file item from the file pointer pool and returns it. It creates one if the file pointer pool is empty. Note that it should be closed when it will never be used. When it's closed, it is not really closed the underlying file pointer but put back to the file pointer pool.