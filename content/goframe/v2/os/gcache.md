+++
title = "gcache"
date = 2024-03-21T17:54:37+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcache

Package gcache provides kinds of cache management for process.

It provides a concurrent-safe in-memory cache adapter for process in default.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gcache/gcache.go#L22)

``` go
const (
	DurationNoExpire = time.Duration(0) // Expire duration that never expires.
)
```

### Variables 

This section is empty.

### Functions 

##### func Contains 

``` go
func Contains(ctx context.Context, key interface{}) (bool, error)
```

Contains checks and returns true if `key` exists in the cache, or else returns false.

##### func Data 

``` go
func Data(ctx context.Context) (map[interface{}]interface{}, error)
```

Data returns a copy of all key-value pairs in the cache as map type. Note that this function may lead lots of memory usage, you can implement this function if necessary.

##### func Get 

``` go
func Get(ctx context.Context, key interface{}) (*gvar.Var, error)
```

Get retrieves and returns the associated value of given `key`. It returns nil if it does not exist, or its value is nil, or it's expired. If you would like to check if the `key` exists in the cache, it's better using function Contains.

##### func GetExpire 

``` go
func GetExpire(ctx context.Context, key interface{}) (time.Duration, error)
```

GetExpire retrieves and returns the expiration of `key` in the cache.

Note that, It returns 0 if the `key` does not expire. It returns -1 if the `key` does not exist in the cache.

##### func GetOrSet 

``` go
func GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (*gvar.Var, error)
```

GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and returns `value` if `key` does not exist in the cache. The key-value pair expires after `duration`.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

##### func GetOrSetFunc 

``` go
func GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error)
```

GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

##### func GetOrSetFuncLock 

``` go
func GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error)
```

GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

Note that it differs from function `GetOrSetFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

##### func KeyStrings 

``` go
func KeyStrings(ctx context.Context) ([]string, error)
```

KeyStrings returns all keys in the cache as string slice.

##### func Keys 

``` go
func Keys(ctx context.Context) ([]interface{}, error)
```

Keys returns all keys in the cache as slice.

##### func MustContains 

``` go
func MustContains(ctx context.Context, key interface{}) bool
```

MustContains acts like Contains, but it panics if any error occurs.

##### func MustData 

``` go
func MustData(ctx context.Context) map[interface{}]interface{}
```

MustData acts like Data, but it panics if any error occurs.

##### func MustGet 

``` go
func MustGet(ctx context.Context, key interface{}) *gvar.Var
```

MustGet acts like Get, but it panics if any error occurs.

##### func MustGetExpire 

``` go
func MustGetExpire(ctx context.Context, key interface{}) time.Duration
```

MustGetExpire acts like GetExpire, but it panics if any error occurs.

##### func MustGetOrSet 

``` go
func MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var
```

MustGetOrSet acts like GetOrSet, but it panics if any error occurs.

##### func MustGetOrSetFunc 

``` go
func MustGetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var
```

MustGetOrSetFunc acts like GetOrSetFunc, but it panics if any error occurs.

##### func MustGetOrSetFuncLock 

``` go
func MustGetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var
```

MustGetOrSetFuncLock acts like GetOrSetFuncLock, but it panics if any error occurs.

##### func MustKeyStrings 

``` go
func MustKeyStrings(ctx context.Context) []string
```

MustKeyStrings acts like KeyStrings, but it panics if any error occurs.

##### func MustKeys 

``` go
func MustKeys(ctx context.Context) []interface{}
```

MustKeys acts like Keys, but it panics if any error occurs.

##### func MustSize 

``` go
func MustSize(ctx context.Context) int
```

MustSize acts like Size, but it panics if any error occurs.

##### func MustValues 

``` go
func MustValues(ctx context.Context) []interface{}
```

MustValues acts like Values, but it panics if any error occurs.

##### func Remove 

``` go
func Remove(ctx context.Context, keys ...interface{}) (value *gvar.Var, err error)
```

Remove deletes one or more keys from cache, and returns its value. If multiple keys are given, it returns the value of the last deleted item.

##### func Removes 

``` go
func Removes(ctx context.Context, keys []interface{}) error
```

Removes deletes `keys` in the cache.

##### func Set 

``` go
func Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error
```

Set sets cache with `key`-`value` pair, which is expired after `duration`.

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

##### func SetIfNotExist 

``` go
func SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error)
```

SetIfNotExist sets cache with `key`-`value` pair which is expired after `duration` if `key` does not exist in the cache. It returns true the `key` does not exist in the cache, and it sets `value` successfully to the cache, or else it returns false.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

##### func SetIfNotExistFunc 

``` go
func SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error)
```

SetIfNotExistFunc sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

The parameter `value` can be type of `func() interface{}`, but it does nothing if its result is nil.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

##### func SetIfNotExistFuncLock 

``` go
func SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error)
```

SetIfNotExistFuncLock sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

Note that it differs from function `SetIfNotExistFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

##### func SetMap 

``` go
func SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error
```

SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

##### func Size 

``` go
func Size(ctx context.Context) (int, error)
```

Size returns the number of items in the cache.

##### func Update 

``` go
func Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error)
```

Update updates the value of `key` without changing its expiration and returns the old value. The returned value `exist` is false if the `key` does not exist in the cache.

It deletes the `key` if given `value` is nil. It does nothing if `key` does not exist in the cache.

##### func UpdateExpire 

``` go
func UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)
```

UpdateExpire updates the expiration of `key` and returns the old expiration duration value.

It returns -1 and does nothing if the `key` does not exist in the cache. It deletes the `key` if `duration` < 0.

##### func Values 

``` go
func Values(ctx context.Context) ([]interface{}, error)
```

Values returns all values in the cache as slice.

### Types 

#### type Adapter 

``` go
type Adapter interface {
	// Set sets cache with `key`-`value` pair, which is expired after `duration`.
	//
	// It does not expire if `duration` == 0.
	// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
	Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error

	// SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.
	//
	// It does not expire if `duration` == 0.
	// It deletes the keys of `data` if `duration` < 0 or given `value` is nil.
	SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error

	// SetIfNotExist sets cache with `key`-`value` pair which is expired after `duration`
	// if `key` does not exist in the cache. It returns true the `key` does not exist in the
	// cache, and it sets `value` successfully to the cache, or else it returns false.
	//
	// It does not expire if `duration` == 0.
	// It deletes the `key` if `duration` < 0 or given `value` is nil.
	SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error)

	// SetIfNotExistFunc sets `key` with result of function `f` and returns true
	// if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.
	//
	// The parameter `value` can be type of `func() interface{}`, but it does nothing if its
	// result is nil.
	//
	// It does not expire if `duration` == 0.
	// It deletes the `key` if `duration` < 0 or given `value` is nil.
	SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)

	// SetIfNotExistFuncLock sets `key` with result of function `f` and returns true
	// if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.
	//
	// It does not expire if `duration` == 0.
	// It deletes the `key` if `duration` < 0 or given `value` is nil.
	//
	// Note that it differs from function `SetIfNotExistFunc` is that the function `f` is executed within
	// writing mutex lock for concurrent safety purpose.
	SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)

	// Get retrieves and returns the associated value of given `key`.
	// It returns nil if it does not exist, or its value is nil, or it's expired.
	// If you would like to check if the `key` exists in the cache, it's better using function Contains.
	Get(ctx context.Context, key interface{}) (*gvar.Var, error)

	// GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and
	// returns `value` if `key` does not exist in the cache. The key-value pair expires
	// after `duration`.
	//
	// It does not expire if `duration` == 0.
	// It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing
	// if `value` is a function and the function result is nil.
	GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error)

	// GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of
	// function `f` and returns its result if `key` does not exist in the cache. The key-value
	// pair expires after `duration`.
	//
	// It does not expire if `duration` == 0.
	// It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing
	// if `value` is a function and the function result is nil.
	GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)

	// GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with result of
	// function `f` and returns its result if `key` does not exist in the cache. The key-value
	// pair expires after `duration`.
	//
	// It does not expire if `duration` == 0.
	// It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing
	// if `value` is a function and the function result is nil.
	//
	// Note that it differs from function `GetOrSetFunc` is that the function `f` is executed within
	// writing mutex lock for concurrent safety purpose.
	GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)

	// Contains checks and returns true if `key` exists in the cache, or else returns false.
	Contains(ctx context.Context, key interface{}) (bool, error)

	// Size returns the number of items in the cache.
	Size(ctx context.Context) (size int, err error)

	// Data returns a copy of all key-value pairs in the cache as map type.
	// Note that this function may lead lots of memory usage, you can implement this function
	// if necessary.
	Data(ctx context.Context) (data map[interface{}]interface{}, err error)

	// Keys returns all keys in the cache as slice.
	Keys(ctx context.Context) (keys []interface{}, err error)

	// Values returns all values in the cache as slice.
	Values(ctx context.Context) (values []interface{}, err error)

	// Update updates the value of `key` without changing its expiration and returns the old value.
	// The returned value `exist` is false if the `key` does not exist in the cache.
	//
	// It deletes the `key` if given `value` is nil.
	// It does nothing if `key` does not exist in the cache.
	Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error)

	// UpdateExpire updates the expiration of `key` and returns the old expiration duration value.
	//
	// It returns -1 and does nothing if the `key` does not exist in the cache.
	// It deletes the `key` if `duration` < 0.
	UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)

	// GetExpire retrieves and returns the expiration of `key` in the cache.
	//
	// Note that,
	// It returns 0 if the `key` does not expire.
	// It returns -1 if the `key` does not exist in the cache.
	GetExpire(ctx context.Context, key interface{}) (time.Duration, error)

	// Remove deletes one or more keys from cache, and returns its value.
	// If multiple keys are given, it returns the value of the last deleted item.
	Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error)

	// Clear clears all data of the cache.
	// Note that this function is sensitive and should be carefully used.
	Clear(ctx context.Context) error

	// Close closes the cache if necessary.
	Close(ctx context.Context) error
}
```

Adapter is the core adapter for cache features implements.

Note that the implementer itself should guarantee the concurrent safety of these functions.

##### func NewAdapterMemory 

``` go
func NewAdapterMemory(lruCap ...int) Adapter
```

NewAdapterMemory creates and returns a new memory cache object.

##### func NewAdapterRedis 

``` go
func NewAdapterRedis(redis *gredis.Redis) Adapter
```

NewAdapterRedis creates and returns a new memory cache object.

#### type AdapterMemory 

``` go
type AdapterMemory struct {
	// contains filtered or unexported fields
}
```

AdapterMemory is an adapter implements using memory.

##### (*AdapterMemory) Clear 

``` go
func (c *AdapterMemory) Clear(ctx context.Context) error
```

Clear clears all data of the cache. Note that this function is sensitive and should be carefully used.

##### (*AdapterMemory) Close 

``` go
func (c *AdapterMemory) Close(ctx context.Context) error
```

Close closes the cache.

##### (*AdapterMemory) Contains 

``` go
func (c *AdapterMemory) Contains(ctx context.Context, key interface{}) (bool, error)
```

Contains checks and returns true if `key` exists in the cache, or else returns false.

##### (*AdapterMemory) Data 

``` go
func (c *AdapterMemory) Data(ctx context.Context) (map[interface{}]interface{}, error)
```

Data returns a copy of all key-value pairs in the cache as map type.

##### (*AdapterMemory) Get 

``` go
func (c *AdapterMemory) Get(ctx context.Context, key interface{}) (*gvar.Var, error)
```

Get retrieves and returns the associated value of given `key`. It returns nil if it does not exist, or its value is nil, or it's expired. If you would like to check if the `key` exists in the cache, it's better using function Contains.

##### (*AdapterMemory) GetExpire 

``` go
func (c *AdapterMemory) GetExpire(ctx context.Context, key interface{}) (time.Duration, error)
```

GetExpire retrieves and returns the expiration of `key` in the cache.

Note that, It returns 0 if the `key` does not expire. It returns -1 if the `key` does not exist in the cache.

##### (*AdapterMemory) GetOrSet 

``` go
func (c *AdapterMemory) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (*gvar.Var, error)
```

GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and returns `value` if `key` does not exist in the cache. The key-value pair expires after `duration`.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

##### (*AdapterMemory) GetOrSetFunc 

``` go
func (c *AdapterMemory) GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error)
```

GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

##### (*AdapterMemory) GetOrSetFuncLock 

``` go
func (c *AdapterMemory) GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error)
```

GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

Note that it differs from function `GetOrSetFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

##### (*AdapterMemory) Keys 

``` go
func (c *AdapterMemory) Keys(ctx context.Context) ([]interface{}, error)
```

Keys returns all keys in the cache as slice.

##### (*AdapterMemory) Remove 

``` go
func (c *AdapterMemory) Remove(ctx context.Context, keys ...interface{}) (*gvar.Var, error)
```

Remove deletes one or more keys from cache, and returns its value. If multiple keys are given, it returns the value of the last deleted item.

##### (*AdapterMemory) Set 

``` go
func (c *AdapterMemory) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error
```

Set sets cache with `key`-`value` pair, which is expired after `duration`.

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

##### (*AdapterMemory) SetIfNotExist 

``` go
func (c *AdapterMemory) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error)
```

SetIfNotExist sets cache with `key`-`value` pair which is expired after `duration` if `key` does not exist in the cache. It returns true the `key` does not exist in the cache, and it sets `value` successfully to the cache, or else it returns false.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

##### (*AdapterMemory) SetIfNotExistFunc 

``` go
func (c *AdapterMemory) SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error)
```

SetIfNotExistFunc sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

The parameter `value` can be type of `func() interface{}`, but it does nothing if its result is nil.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

##### (*AdapterMemory) SetIfNotExistFuncLock 

``` go
func (c *AdapterMemory) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error)
```

SetIfNotExistFuncLock sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

Note that it differs from function `SetIfNotExistFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

##### (*AdapterMemory) SetMap 

``` go
func (c *AdapterMemory) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error
```

SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

##### (*AdapterMemory) Size 

``` go
func (c *AdapterMemory) Size(ctx context.Context) (size int, err error)
```

Size returns the size of the cache.

##### (*AdapterMemory) Update 

``` go
func (c *AdapterMemory) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error)
```

Update updates the value of `key` without changing its expiration and returns the old value. The returned value `exist` is false if the `key` does not exist in the cache.

It deletes the `key` if given `value` is nil. It does nothing if `key` does not exist in the cache.

##### (*AdapterMemory) UpdateExpire 

``` go
func (c *AdapterMemory) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)
```

UpdateExpire updates the expiration of `key` and returns the old expiration duration value.

It returns -1 and does nothing if the `key` does not exist in the cache. It deletes the `key` if `duration` < 0.

##### (*AdapterMemory) Values 

``` go
func (c *AdapterMemory) Values(ctx context.Context) ([]interface{}, error)
```

Values returns all values in the cache as slice.

#### type AdapterRedis 

``` go
type AdapterRedis struct {
	// contains filtered or unexported fields
}
```

AdapterRedis is the gcache adapter implements using Redis server.

##### (*AdapterRedis) Clear 

``` go
func (c *AdapterRedis) Clear(ctx context.Context) (err error)
```

Clear clears all data of the cache. Note that this function is sensitive and should be carefully used. It uses `FLUSHDB` command in redis server, which might be disabled in server.

##### (*AdapterRedis) Close 

``` go
func (c *AdapterRedis) Close(ctx context.Context) error
```

Close closes the cache.

##### (*AdapterRedis) Contains 

``` go
func (c *AdapterRedis) Contains(ctx context.Context, key interface{}) (bool, error)
```

Contains checks and returns true if `key` exists in the cache, or else returns false.

##### (*AdapterRedis) Data 

``` go
func (c *AdapterRedis) Data(ctx context.Context) (map[interface{}]interface{}, error)
```

Data returns a copy of all key-value pairs in the cache as map type. Note that this function may lead lots of memory usage, you can implement this function if necessary.

##### (*AdapterRedis) Get 

``` go
func (c *AdapterRedis) Get(ctx context.Context, key interface{}) (*gvar.Var, error)
```

Get retrieves and returns the associated value of given <key>. It returns nil if it does not exist or its value is nil.

##### (*AdapterRedis) GetExpire 

``` go
func (c *AdapterRedis) GetExpire(ctx context.Context, key interface{}) (time.Duration, error)
```

GetExpire retrieves and returns the expiration of `key` in the cache.

Note that, It returns 0 if the `key` does not expire. It returns -1 if the `key` does not exist in the cache.

##### (*AdapterRedis) GetOrSet 

``` go
func (c *AdapterRedis) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error)
```

GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and returns `value` if `key` does not exist in the cache. The key-value pair expires after `duration`.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

##### (*AdapterRedis) GetOrSetFunc 

``` go
func (c *AdapterRedis) GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)
```

GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

##### (*AdapterRedis) GetOrSetFuncLock 

``` go
func (c *AdapterRedis) GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)
```

GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

Note that it differs from function `GetOrSetFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

##### (*AdapterRedis) Keys 

``` go
func (c *AdapterRedis) Keys(ctx context.Context) ([]interface{}, error)
```

Keys returns all keys in the cache as slice.

##### (*AdapterRedis) Remove 

``` go
func (c *AdapterRedis) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error)
```

Remove deletes the one or more keys from cache, and returns its value. If multiple keys are given, it returns the value of the deleted last item.

##### (*AdapterRedis) Set 

``` go
func (c *AdapterRedis) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (err error)
```

Set sets cache with `key`-`value` pair, which is expired after `duration`.

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

##### (*AdapterRedis) SetIfNotExist 

``` go
func (c *AdapterRedis) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error)
```

SetIfNotExist sets cache with `key`-`value` pair which is expired after `duration` if `key` does not exist in the cache. It returns true the `key` does not exist in the cache, and it sets `value` successfully to the cache, or else it returns false.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

##### (*AdapterRedis) SetIfNotExistFunc 

``` go
func (c *AdapterRedis) SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)
```

SetIfNotExistFunc sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

The parameter `value` can be type of `func() interface{}`, but it does nothing if its result is nil.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

##### (*AdapterRedis) SetIfNotExistFuncLock 

``` go
func (c *AdapterRedis) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)
```

SetIfNotExistFuncLock sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

Note that it differs from function `SetIfNotExistFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

##### (*AdapterRedis) SetMap 

``` go
func (c *AdapterRedis) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error
```

SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

##### (*AdapterRedis) Size 

``` go
func (c *AdapterRedis) Size(ctx context.Context) (size int, err error)
```

Size returns the number of items in the cache.

##### (*AdapterRedis) Update 

``` go
func (c *AdapterRedis) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error)
```

Update updates the value of `key` without changing its expiration and returns the old value. The returned value `exist` is false if the `key` does not exist in the cache.

It deletes the `key` if given `value` is nil. It does nothing if `key` does not exist in the cache.

##### (*AdapterRedis) UpdateExpire 

``` go
func (c *AdapterRedis) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)
```

UpdateExpire updates the expiration of `key` and returns the old expiration duration value.

It returns -1 and does nothing if the `key` does not exist in the cache. It deletes the `key` if `duration` < 0.

##### (*AdapterRedis) Values 

``` go
func (c *AdapterRedis) Values(ctx context.Context) ([]interface{}, error)
```

Values returns all values in the cache as slice.

#### type Cache 

``` go
type Cache struct {
	// contains filtered or unexported fields
}
```

Cache struct.

##### func New 

``` go
func New(lruCap ...int) *Cache
```

New creates and returns a new cache object using default memory adapter. Note that the LRU feature is only available using memory adapter.

##### Example

``` go
```
##### func NewWithAdapter 

``` go
func NewWithAdapter(adapter Adapter) *Cache
```

NewWithAdapter creates and returns a Cache object with given Adapter implements.

##### (*Cache) GetAdapter 

``` go
func (c *Cache) GetAdapter() Adapter
```

GetAdapter returns the adapter that is set in current Cache.

##### Example

``` go
```
##### (*Cache) KeyStrings 

``` go
func (c *Cache) KeyStrings(ctx context.Context) ([]string, error)
```

KeyStrings returns all keys in the cache as string slice.

##### Example

``` go
```
##### (*Cache) MustContains 

``` go
func (c *Cache) MustContains(ctx context.Context, key interface{}) bool
```

MustContains acts like Contains, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) MustData 

``` go
func (c *Cache) MustData(ctx context.Context) map[interface{}]interface{}
```

MustData acts like Data, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) MustGet 

``` go
func (c *Cache) MustGet(ctx context.Context, key interface{}) *gvar.Var
```

MustGet acts like Get, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) MustGetExpire 

``` go
func (c *Cache) MustGetExpire(ctx context.Context, key interface{}) time.Duration
```

MustGetExpire acts like GetExpire, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) MustGetOrSet 

``` go
func (c *Cache) MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var
```

MustGetOrSet acts like GetOrSet, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) MustGetOrSetFunc 

``` go
func (c *Cache) MustGetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var
```

MustGetOrSetFunc acts like GetOrSetFunc, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) MustGetOrSetFuncLock 

``` go
func (c *Cache) MustGetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var
```

MustGetOrSetFuncLock acts like GetOrSetFuncLock, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) MustKeyStrings 

``` go
func (c *Cache) MustKeyStrings(ctx context.Context) []string
```

MustKeyStrings acts like KeyStrings, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) MustKeys 

``` go
func (c *Cache) MustKeys(ctx context.Context) []interface{}
```

MustKeys acts like Keys, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) MustSize 

``` go
func (c *Cache) MustSize(ctx context.Context) int
```

MustSize acts like Size, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) MustValues 

``` go
func (c *Cache) MustValues(ctx context.Context) []interface{}
```

MustValues acts like Values, but it panics if any error occurs.

##### Example

``` go
```
##### (*Cache) Removes 

``` go
func (c *Cache) Removes(ctx context.Context, keys []interface{}) error
```

Removes deletes `keys` in the cache.

##### Example

``` go
```
##### (*Cache) SetAdapter 

``` go
func (c *Cache) SetAdapter(adapter Adapter)
```

SetAdapter changes the adapter for this cache. Be very note that, this setting function is not concurrent-safe, which means you should not call this setting function concurrently in multiple goroutines.

##### Example

``` go
```
#### type Func 

``` go
type Func func(ctx context.Context) (value interface{}, err error)
```

Func is the cache function that calculates and returns the value.