+++
title = "gcache"
date = 2024-03-21T17:54:37+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcache](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gcache)

Package gcache provides kinds of cache management for process.

​	软件包 gcache 为进程提供了各种缓存管理。

It provides a concurrent-safe in-memory cache adapter for process in default.

​	默认情况下，它为进程提供并发安全的内存中缓存适配器。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gcache/gcache.go#L22)

```go
const (
	DurationNoExpire = time.Duration(0) // Expire duration that never expires.
)
```

## 变量

This section is empty.

## 函数

#### func Contains

```go
func Contains(ctx context.Context, key interface{}) (bool, error)
```

Contains checks and returns true if `key` exists in the cache, or else returns false.

​	包含检查并返回 true（如果 `key` 缓存中存在），否则返回 false。

#### func Data

```go
func Data(ctx context.Context) (map[interface{}]interface{}, error)
```

Data returns a copy of all key-value pairs in the cache as map type. Note that this function may lead lots of memory usage, you can implement this function if necessary.

​	数据以映射类型返回缓存中所有键值对的副本。请注意，此函数可能会导致大量内存使用，如有必要，可以实现此函数。

#### func Get

```go
func Get(ctx context.Context, key interface{}) (*gvar.Var, error)
```

Get retrieves and returns the associated value of given `key`. It returns nil if it does not exist, or its value is nil, or it’s expired. If you would like to check if the `key` exists in the cache, it’s better using function Contains.

​	Get 检索并返回给定 `key` 的关联值。如果它不存在，或者它的值为 nil，或者它已过期，则返回 nil。如果您想检查缓存中是否存在， `key` 最好使用函数 Contains。

#### func GetExpire

```go
func GetExpire(ctx context.Context, key interface{}) (time.Duration, error)
```

GetExpire retrieves and returns the expiration of `key` in the cache.

​	GetExpire 检索并返回缓存中的过 `key` 期时间。

Note that, It returns 0 if the `key` does not expire. It returns -1 if the `key` does not exist in the cache.

​	请注意，如果 不 `key` 过期，则返回 0。如果缓存中不存在， `key` 则返回 -1。

#### func GetOrSet

```go
func GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (*gvar.Var, error)
```

GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and returns `value` if `key` does not exist in the cache. The key-value pair expires after `duration`.

​	GetOrSet 检索并返回 的 `key` 值，或者 sets `key` - `value` pair 的值，如果 `key` 缓存中不存在则返回 `value` 。键值对在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil，但如果 `value` 是一个函数并且函数结果为 nil，则它不执行任何操作。

#### func GetOrSetFunc

```go
func GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error)
```

GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

​	GetOrSetFunc 检索并返回 `key` 的值，或具有函数 `f` 结果的集合 `key` ，如果 `key` 缓存中不存在，则返回其结果。键值对在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil，但如果 `value` 是一个函数并且函数结果为 nil，则它不执行任何操作。

#### func GetOrSetFuncLock

```go
func GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error)
```

GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

​	GetOrSetFuncLock 检索并返回 `key` 的值，或使用 result of 函数 `f` 进行设置 `key` ，如果 `key` 缓存中不存在，则返回其结果。键值对在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil，但如果 `value` 是一个函数并且函数结果为 nil，则它不执行任何操作。

Note that it differs from function `GetOrSetFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

​	请注意，它与函数 `GetOrSetFunc` 的不同之处在于，出于并发安全目的，该函数 `f` 是在写入互斥锁的情况下执行的。

#### func KeyStrings

```go
func KeyStrings(ctx context.Context) ([]string, error)
```

KeyStrings returns all keys in the cache as string slice.

​	KeyStrings 将缓存中的所有键作为字符串切片返回。

#### func Keys

```go
func Keys(ctx context.Context) ([]interface{}, error)
```

Keys returns all keys in the cache as slice.

​	Keys 将缓存中的所有键作为切片返回。

#### func MustContains

```go
func MustContains(ctx context.Context, key interface{}) bool
```

MustContains acts like Contains, but it panics if any error occurs.

​	MustContains 的行为类似于 Contains，但如果发生任何错误，它会崩溃。

#### func MustData

```go
func MustData(ctx context.Context) map[interface{}]interface{}
```

MustData acts like Data, but it panics if any error occurs.

​	MustData 的行为类似于 Data，但如果发生任何错误，它就会崩溃。

#### func MustGet

```go
func MustGet(ctx context.Context, key interface{}) *gvar.Var
```

MustGet acts like Get, but it panics if any error occurs.

​	MustGet 的行为类似于 Get，但如果发生任何错误，它会崩溃。

#### func MustGetExpire

```go
func MustGetExpire(ctx context.Context, key interface{}) time.Duration
```

MustGetExpire acts like GetExpire, but it panics if any error occurs.

​	MustGetExpire 的行为类似于 GetExpire，但如果发生任何错误，它会崩溃。

#### func MustGetOrSet

```go
func MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var
```

MustGetOrSet acts like GetOrSet, but it panics if any error occurs.

​	MustGetOrSet 的行为类似于 GetOrSet，但如果发生任何错误，它会崩溃。

#### func MustGetOrSetFunc

```go
func MustGetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var
```

MustGetOrSetFunc acts like GetOrSetFunc, but it panics if any error occurs.

​	MustGetOrSetFunc 的行为类似于 GetOrSetFunc，但如果发生任何错误，它会崩溃。

#### func MustGetOrSetFuncLock

```go
func MustGetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var
```

MustGetOrSetFuncLock acts like GetOrSetFuncLock, but it panics if any error occurs.

​	MustGetOrSetFuncLock 的行为类似于 GetOrSetFuncLock，但如果发生任何错误，它会崩溃。

#### func MustKeyStrings

```go
func MustKeyStrings(ctx context.Context) []string
```

MustKeyStrings acts like KeyStrings, but it panics if any error occurs.

​	MustKeyStrings 的行为类似于 KeyStrings，但如果发生任何错误，它会崩溃。

#### func MustKeys

```go
func MustKeys(ctx context.Context) []interface{}
```

MustKeys acts like Keys, but it panics if any error occurs.

​	MustKeys 的行为类似于 Keys，但如果发生任何错误，它会崩溃。

#### func MustSize

```go
func MustSize(ctx context.Context) int
```

MustSize acts like Size, but it panics if any error occurs.

​	MustSize 的作用类似于 Size，但如果发生任何错误，它就会崩溃。

#### func MustValues

```go
func MustValues(ctx context.Context) []interface{}
```

MustValues acts like Values, but it panics if any error occurs.

​	MustValues 的作用类似于 Values，但如果发生任何错误，它就会崩溃。

#### func Remove

```go
func Remove(ctx context.Context, keys ...interface{}) (value *gvar.Var, err error)
```

Remove deletes one or more keys from cache, and returns its value. If multiple keys are given, it returns the value of the last deleted item.

​	Remove 从缓存中删除一个或多个键，并返回其值。如果给定了多个键，则返回上次删除的项目的值。

#### func Removes

```go
func Removes(ctx context.Context, keys []interface{}) error
```

Removes deletes `keys` in the cache.

​	删除缓存 `keys` 中的删除内容。

#### func Set

```go
func Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error
```

Set sets cache with `key`-`value` pair, which is expired after `duration`.

​	设置 cache with `key` - `value` pair，在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `data` 如果 `duration` < 0 或给定 `value` 为 nil，则删除键。

#### func SetIfNotExist

```go
func SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error)
```

SetIfNotExist sets cache with `key`-`value` pair which is expired after `duration` if `key` does not exist in the cache. It returns true the `key` does not exist in the cache, and it sets `value` successfully to the cache, or else it returns false.

​	SetIfNotExist 将缓存设置为 `key` - 对，如果缓存中不存在，则在 `duration` if `key` `value` 之后过期。它返回 true the `key` does not exist in the cache，并成功设置为 `value` 缓存，否则返回 false。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil。

#### func SetIfNotExistFunc

```go
func SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error)
```

SetIfNotExistFunc sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

​	SetIfNotExistFunc 设置 `key` 函数 `f` 的结果，如果 `key` 缓存中不存在，则返回 true，否则不执行任何操作，如果 `key` 已存在，则返回 false。

The parameter `value` can be type of `func() interface{}`, but it does nothing if its result is nil.

​	参数 `value` 的类型可以是 `func() interface{}` ，但如果其结果为 nil，则不执行任何操作。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil。

#### func SetIfNotExistFuncLock

```go
func SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error)
```

SetIfNotExistFuncLock sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

​	SetIfNotExistFuncLock 设置 `key` 函数 `f` 的结果，如果 `key` 缓存中不存在，则返回 true，否则不执行任何操作，如果已存在，则 `key` 返回 false。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil。

Note that it differs from function `SetIfNotExistFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

​	请注意，它与函数 `SetIfNotExistFunc` 的不同之处在于，出于并发安全目的，该函数 `f` 是在写入互斥锁的情况下执行的。

#### func SetMap

```go
func SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error
```

SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.

​	SetMap 按 `data` map 批量设置键值对缓存，该缓存在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `data` 如果 `duration` < 0 或给定 `value` 为 nil，则删除键。

#### func Size

```go
func Size(ctx context.Context) (int, error)
```

Size returns the number of items in the cache.

​	Size 返回缓存中的项数。

#### func Update

```go
func Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error)
```

Update updates the value of `key` without changing its expiration and returns the old value. The returned value `exist` is false if the `key` does not exist in the cache.

​	Update 在不更改其过期时间的情况下更新 的 `key` 值，并返回旧值。如果缓存中不存在， `key` 则返回的值 `exist` 为 false。

It deletes the `key` if given `value` is nil. It does nothing if `key` does not exist in the cache.

​	它删除了 `key` if given `value` is nil。如果 `key` 缓存中不存在，则它不执行任何操作。

#### func UpdateExpire

```go
func UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)
```

UpdateExpire updates the expiration of `key` and returns the old expiration duration value.

​	UpdateExpire 更新过 `key` 期时间并返回旧的过期持续时间值。

It returns -1 and does nothing if the `key` does not exist in the cache. It deletes the `key` if `duration` < 0.

​	它返回 -1，如果缓存中不存在， `key` 则不执行任何操作。它删除了 `key` if `duration` < 0。

#### func Values

```go
func Values(ctx context.Context) ([]interface{}, error)
```

Values returns all values in the cache as slice.

​	Values 将缓存中的所有值作为切片返回。

## 类型

### type Adapter

```go
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

​	适配器是缓存功能实现的核心适配器。

Note that the implementer itself should guarantee the concurrent safety of these functions.

​	请注意，实现者本身应保证这些功能的并发安全性。

#### func NewAdapterMemory

```go
func NewAdapterMemory(lruCap ...int) Adapter
```

NewAdapterMemory creates and returns a new memory cache object.

​	NewAdapterMemory 创建并返回新的内存缓存对象。

#### func NewAdapterRedis

```go
func NewAdapterRedis(redis *gredis.Redis) Adapter
```

NewAdapterRedis creates and returns a new memory cache object.

​	NewAdapterRedis 创建并返回新的内存缓存对象。

### type AdapterMemory

```go
type AdapterMemory struct {
	// contains filtered or unexported fields
}
```

AdapterMemory is an adapter implements using memory.

​	AdapterMemory 是使用内存实现的适配器。

#### (*AdapterMemory) Clear

```go
func (c *AdapterMemory) Clear(ctx context.Context) error
```

Clear clears all data of the cache. Note that this function is sensitive and should be carefully used.

​	清除缓存中的所有数据。请注意，此功能很敏感，应谨慎使用。

#### (*AdapterMemory) Close

```go
func (c *AdapterMemory) Close(ctx context.Context) error
```

Close closes the cache.

​	关闭 关闭缓存。

#### (*AdapterMemory) Contains

```go
func (c *AdapterMemory) Contains(ctx context.Context, key interface{}) (bool, error)
```

Contains checks and returns true if `key` exists in the cache, or else returns false.

​	包含检查并返回 true（如果 `key` 缓存中存在），否则返回 false。

#### (*AdapterMemory) Data

```go
func (c *AdapterMemory) Data(ctx context.Context) (map[interface{}]interface{}, error)
```

Data returns a copy of all key-value pairs in the cache as map type.

​	数据以映射类型返回缓存中所有键值对的副本。

#### (*AdapterMemory) Get

```go
func (c *AdapterMemory) Get(ctx context.Context, key interface{}) (*gvar.Var, error)
```

Get retrieves and returns the associated value of given `key`. It returns nil if it does not exist, or its value is nil, or it’s expired. If you would like to check if the `key` exists in the cache, it’s better using function Contains.

​	Get 检索并返回给定 `key` 的关联值。如果它不存在，或者它的值为 nil，或者它已过期，则返回 nil。如果您想检查缓存中是否存在， `key` 最好使用函数 Contains。

#### (*AdapterMemory) GetExpire

```go
func (c *AdapterMemory) GetExpire(ctx context.Context, key interface{}) (time.Duration, error)
```

GetExpire retrieves and returns the expiration of `key` in the cache.

​	GetExpire 检索并返回缓存中的过 `key` 期时间。

Note that, It returns 0 if the `key` does not expire. It returns -1 if the `key` does not exist in the cache.

​	请注意，如果 不 `key` 过期，则返回 0。如果缓存中不存在， `key` 则返回 -1。

#### (*AdapterMemory) GetOrSet

```go
func (c *AdapterMemory) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (*gvar.Var, error)
```

GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and returns `value` if `key` does not exist in the cache. The key-value pair expires after `duration`.

​	GetOrSet 检索并返回 的 `key` 值，或者 sets `key` - `value` pair 的值，如果 `key` 缓存中不存在则返回 `value` 。键值对在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil，但如果 `value` 是一个函数并且函数结果为 nil，则它不执行任何操作。

#### (*AdapterMemory) GetOrSetFunc

```go
func (c *AdapterMemory) GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error)
```

GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

​	GetOrSetFunc 检索并返回 `key` 的值，或具有函数 `f` 结果的集合 `key` ，如果 `key` 缓存中不存在，则返回其结果。键值对在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil，但如果 `value` 是一个函数并且函数结果为 nil，则它不执行任何操作。

#### (*AdapterMemory) GetOrSetFuncLock

```go
func (c *AdapterMemory) GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (*gvar.Var, error)
```

GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

​	GetOrSetFuncLock 检索并返回 `key` 的值，或使用 result of 函数 `f` 进行设置 `key` ，如果 `key` 缓存中不存在，则返回其结果。键值对在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil，但如果 `value` 是一个函数并且函数结果为 nil，则它不执行任何操作。

Note that it differs from function `GetOrSetFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

​	请注意，它与函数 `GetOrSetFunc` 的不同之处在于，出于并发安全目的，该函数 `f` 是在写入互斥锁的情况下执行的。

#### (*AdapterMemory) Keys

```go
func (c *AdapterMemory) Keys(ctx context.Context) ([]interface{}, error)
```

Keys returns all keys in the cache as slice.

​	Keys 将缓存中的所有键作为切片返回。

#### (*AdapterMemory) Remove

```go
func (c *AdapterMemory) Remove(ctx context.Context, keys ...interface{}) (*gvar.Var, error)
```

Remove deletes one or more keys from cache, and returns its value. If multiple keys are given, it returns the value of the last deleted item.

​	Remove 从缓存中删除一个或多个键，并返回其值。如果给定了多个键，则返回上次删除的项目的值。

#### (*AdapterMemory) Set

```go
func (c *AdapterMemory) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error
```

Set sets cache with `key`-`value` pair, which is expired after `duration`.

​	设置 cache with `key` - `value` pair，在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `data` 如果 `duration` < 0 或给定 `value` 为 nil，则删除键。

#### (*AdapterMemory) SetIfNotExist

```go
func (c *AdapterMemory) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error)
```

SetIfNotExist sets cache with `key`-`value` pair which is expired after `duration` if `key` does not exist in the cache. It returns true the `key` does not exist in the cache, and it sets `value` successfully to the cache, or else it returns false.

​	SetIfNotExist 将缓存设置为 `key` - 对，如果缓存中不存在，则在 `duration` if `key` `value` 之后过期。它返回 true the `key` does not exist in the cache，并成功设置为 `value` 缓存，否则返回 false。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil。

#### (*AdapterMemory) SetIfNotExistFunc

```go
func (c *AdapterMemory) SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error)
```

SetIfNotExistFunc sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

​	SetIfNotExistFunc 设置 `key` 函数 `f` 的结果，如果 `key` 缓存中不存在，则返回 true，否则不执行任何操作，如果 `key` 已存在，则返回 false。

The parameter `value` can be type of `func() interface{}`, but it does nothing if its result is nil.

​	参数 `value` 的类型可以是 `func() interface{}` ，但如果其结果为 nil，则不执行任何操作。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil。

#### (*AdapterMemory) SetIfNotExistFuncLock

```go
func (c *AdapterMemory) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (bool, error)
```

SetIfNotExistFuncLock sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

​	SetIfNotExistFuncLock 设置 `key` 函数 `f` 的结果，如果 `key` 缓存中不存在，则返回 true，否则不执行任何操作，如果已存在，则 `key` 返回 false。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil。

Note that it differs from function `SetIfNotExistFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

​	请注意，它与函数 `SetIfNotExistFunc` 的不同之处在于，出于并发安全目的，该函数 `f` 是在写入互斥锁的情况下执行的。

#### (*AdapterMemory) SetMap

```go
func (c *AdapterMemory) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error
```

SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.

​	SetMap 按 `data` map 批量设置键值对缓存，该缓存在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `data` 如果 `duration` < 0 或给定 `value` 为 nil，则删除键。

#### (*AdapterMemory) Size

```go
func (c *AdapterMemory) Size(ctx context.Context) (size int, err error)
```

Size returns the size of the cache.

​	Size 返回缓存的大小。

#### (*AdapterMemory) Update

```go
func (c *AdapterMemory) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error)
```

Update updates the value of `key` without changing its expiration and returns the old value. The returned value `exist` is false if the `key` does not exist in the cache.

​	Update 在不更改其过期时间的情况下更新 的 `key` 值，并返回旧值。如果缓存中不存在， `key` 则返回的值 `exist` 为 false。

It deletes the `key` if given `value` is nil. It does nothing if `key` does not exist in the cache.

​	它删除了 `key` if given `value` is nil。如果 `key` 缓存中不存在，则它不执行任何操作。

#### (*AdapterMemory) UpdateExpire

```go
func (c *AdapterMemory) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)
```

UpdateExpire updates the expiration of `key` and returns the old expiration duration value.

​	UpdateExpire 更新过 `key` 期时间并返回旧的过期持续时间值。

It returns -1 and does nothing if the `key` does not exist in the cache. It deletes the `key` if `duration` < 0.

​	它返回 -1，如果缓存中不存在， `key` 则不执行任何操作。它删除了 `key` if `duration` < 0。

#### (*AdapterMemory) Values

```go
func (c *AdapterMemory) Values(ctx context.Context) ([]interface{}, error)
```

Values returns all values in the cache as slice.

​	Values 将缓存中的所有值作为切片返回。

### type AdapterRedis

```go
type AdapterRedis struct {
	// contains filtered or unexported fields
}
```

AdapterRedis is the gcache adapter implements using Redis server.

​	AdapterRedis 是使用 Redis 服务器实现的 gcache 适配器。

#### (*AdapterRedis) Clear

```go
func (c *AdapterRedis) Clear(ctx context.Context) (err error)
```

Clear clears all data of the cache. Note that this function is sensitive and should be carefully used. It uses `FLUSHDB` command in redis server, which might be disabled in server.

​	清除缓存中的所有数据。请注意，此功能很敏感，应谨慎使用。它在 redis 服务器中使用 `FLUSHDB` 命令，该命令可能在服务器中被禁用。

#### (*AdapterRedis) Close

```go
func (c *AdapterRedis) Close(ctx context.Context) error
```

Close closes the cache.

​	关闭 关闭缓存。

#### (*AdapterRedis) Contains

```go
func (c *AdapterRedis) Contains(ctx context.Context, key interface{}) (bool, error)
```

Contains checks and returns true if `key` exists in the cache, or else returns false.

​	包含检查并返回 true（如果 `key` 缓存中存在），否则返回 false。

#### (*AdapterRedis) Data

```go
func (c *AdapterRedis) Data(ctx context.Context) (map[interface{}]interface{}, error)
```

Data returns a copy of all key-value pairs in the cache as map type. Note that this function may lead lots of memory usage, you can implement this function if necessary.

​	数据以映射类型返回缓存中所有键值对的副本。请注意，此函数可能会导致大量内存使用，如有必要，可以实现此函数。

#### (*AdapterRedis) Get

```go
func (c *AdapterRedis) Get(ctx context.Context, key interface{}) (*gvar.Var, error)
```

Get retrieves and returns the associated value of given . It returns nil if it does not exist or its value is nil.

​	Get 检索并返回给定的关联值。如果它不存在或其值为 nil，则返回 nil。

#### (*AdapterRedis) GetExpire

```go
func (c *AdapterRedis) GetExpire(ctx context.Context, key interface{}) (time.Duration, error)
```

GetExpire retrieves and returns the expiration of `key` in the cache.

​	GetExpire 检索并返回缓存中的过 `key` 期时间。

Note that, It returns 0 if the `key` does not expire. It returns -1 if the `key` does not exist in the cache.

​	请注意，如果 不 `key` 过期，则返回 0。如果缓存中不存在， `key` 则返回 -1。

#### (*AdapterRedis) GetOrSet

```go
func (c *AdapterRedis) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error)
```

GetOrSet retrieves and returns the value of `key`, or sets `key`-`value` pair and returns `value` if `key` does not exist in the cache. The key-value pair expires after `duration`.

​	GetOrSet 检索并返回 的 `key` 值，或者 sets `key` - `value` pair 的值，如果 `key` 缓存中不存在则返回 `value` 。键值对在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil，但如果 `value` 是一个函数并且函数结果为 nil，则它不执行任何操作。

#### (*AdapterRedis) GetOrSetFunc

```go
func (c *AdapterRedis) GetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)
```

GetOrSetFunc retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

​	GetOrSetFunc 检索并返回 `key` 的值，或具有函数 `f` 结果的集合 `key` ，如果 `key` 缓存中不存在，则返回其结果。键值对在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil，但如果 `value` 是一个函数并且函数结果为 nil，则它不执行任何操作。

#### (*AdapterRedis) GetOrSetFuncLock

```go
func (c *AdapterRedis) GetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (result *gvar.Var, err error)
```

GetOrSetFuncLock retrieves and returns the value of `key`, or sets `key` with result of function `f` and returns its result if `key` does not exist in the cache. The key-value pair expires after `duration`.

​	GetOrSetFuncLock 检索并返回 `key` 的值，或使用 result of 函数 `f` 进行设置 `key` ，如果 `key` 缓存中不存在，则返回其结果。键值对在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil, but it does nothing if `value` is a function and the function result is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil，但如果 `value` 是一个函数并且函数结果为 nil，则它不执行任何操作。

Note that it differs from function `GetOrSetFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

​	请注意，它与函数 `GetOrSetFunc` 的不同之处在于，出于并发安全目的，该函数 `f` 是在写入互斥锁的情况下执行的。

#### (*AdapterRedis) Keys

```go
func (c *AdapterRedis) Keys(ctx context.Context) ([]interface{}, error)
```

Keys returns all keys in the cache as slice.

​	Keys 将缓存中的所有键作为切片返回。

#### (*AdapterRedis) Remove

```go
func (c *AdapterRedis) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error)
```

Remove deletes the one or more keys from cache, and returns its value. If multiple keys are given, it returns the value of the deleted last item.

​	Remove 从缓存中删除一个或多个键，并返回其值。如果给定了多个键，则返回已删除最后一项的值。

#### (*AdapterRedis) Set

```go
func (c *AdapterRedis) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (err error)
```

Set sets cache with `key`-`value` pair, which is expired after `duration`.

​	设置 cache with `key` - `value` pair，在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `data` 如果 `duration` < 0 或给定 `value` 为 nil，则删除键。

#### (*AdapterRedis) SetIfNotExist

```go
func (c *AdapterRedis) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (bool, error)
```

SetIfNotExist sets cache with `key`-`value` pair which is expired after `duration` if `key` does not exist in the cache. It returns true the `key` does not exist in the cache, and it sets `value` successfully to the cache, or else it returns false.

​	SetIfNotExist 将缓存设置为 `key` - 对，如果缓存中不存在，则在 `duration` if `key` `value` 之后过期。它返回 true the `key` does not exist in the cache，并成功设置为 `value` 缓存，否则返回 false。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil。

#### (*AdapterRedis) SetIfNotExistFunc

```go
func (c *AdapterRedis) SetIfNotExistFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)
```

SetIfNotExistFunc sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

​	SetIfNotExistFunc 设置 `key` 函数 `f` 的结果，如果 `key` 缓存中不存在，则返回 true，否则不执行任何操作，如果 `key` 已存在，则返回 false。

The parameter `value` can be type of `func() interface{}`, but it does nothing if its result is nil.

​	参数 `value` 的类型可以是 `func() interface{}` ，但如果其结果为 nil，则不执行任何操作。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil。

#### (*AdapterRedis) SetIfNotExistFuncLock

```go
func (c *AdapterRedis) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) (ok bool, err error)
```

SetIfNotExistFuncLock sets `key` with result of function `f` and returns true if `key` does not exist in the cache, or else it does nothing and returns false if `key` already exists.

​	SetIfNotExistFuncLock 设置 `key` 函数 `f` 的结果，如果 `key` 缓存中不存在，则返回 true，否则不执行任何操作，如果已存在，则 `key` 返回 false。

It does not expire if `duration` == 0. It deletes the `key` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `key` 它删除了 if `duration` < 0 或 given `value` 为 nil。

Note that it differs from function `SetIfNotExistFunc` is that the function `f` is executed within writing mutex lock for concurrent safety purpose.

​	请注意，它与函数 `SetIfNotExistFunc` 的不同之处在于，出于并发安全目的，该函数 `f` 是在写入互斥锁的情况下执行的。

#### (*AdapterRedis) SetMap

```go
func (c *AdapterRedis) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error
```

SetMap batch sets cache with key-value pairs by `data` map, which is expired after `duration`.

​	SetMap 按 `data` map 批量设置键值对缓存，该缓存在 `duration` 之后过期。

It does not expire if `duration` == 0. It deletes the keys of `data` if `duration` < 0 or given `value` is nil.

​	如果 `duration` == 0，它不会过期。 `data` 如果 `duration` < 0 或给定 `value` 为 nil，则删除键。

#### (*AdapterRedis) Size

```go
func (c *AdapterRedis) Size(ctx context.Context) (size int, err error)
```

Size returns the number of items in the cache.

​	Size 返回缓存中的项数。

#### (*AdapterRedis) Update

```go
func (c *AdapterRedis) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error)
```

Update updates the value of `key` without changing its expiration and returns the old value. The returned value `exist` is false if the `key` does not exist in the cache.

​	Update 在不更改其过期时间的情况下更新 的 `key` 值，并返回旧值。如果缓存中不存在， `key` 则返回的值 `exist` 为 false。

It deletes the `key` if given `value` is nil. It does nothing if `key` does not exist in the cache.

​	它删除了 `key` if given `value` is nil。如果 `key` 缓存中不存在，则它不执行任何操作。

#### (*AdapterRedis) UpdateExpire

```go
func (c *AdapterRedis) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error)
```

UpdateExpire updates the expiration of `key` and returns the old expiration duration value.

​	UpdateExpire 更新过 `key` 期时间并返回旧的过期持续时间值。

It returns -1 and does nothing if the `key` does not exist in the cache. It deletes the `key` if `duration` < 0.

​	它返回 -1，如果缓存中不存在， `key` 则不执行任何操作。它删除了 `key` if `duration` < 0。

#### (*AdapterRedis) Values

```go
func (c *AdapterRedis) Values(ctx context.Context) ([]interface{}, error)
```

Values returns all values in the cache as slice.

​	Values 将缓存中的所有值作为切片返回。

### type Cache

```go
type Cache struct {
	// contains filtered or unexported fields
}
```

Cache struct.

​	缓存结构。

#### func New

```go
func New(lruCap ...int) *Cache
```

New creates and returns a new cache object using default memory adapter. Note that the LRU feature is only available using memory adapter.

​	new 使用默认内存适配器创建并返回新的缓存对象。请注意，LRU 功能仅适用于内存适配器。

##### Example

``` go
```

#### func NewWithAdapter

```go
func NewWithAdapter(adapter Adapter) *Cache
```

NewWithAdapter creates and returns a Cache object with given Adapter implements.

​	NewWithAdapter 创建并返回具有给定 Adapter 实现的 Cache 对象。

#### (*Cache) GetAdapter

```go
func (c *Cache) GetAdapter() Adapter
```

GetAdapter returns the adapter that is set in current Cache.

​	GetAdapter 返回在当前缓存中设置的适配器。

##### Example

``` go
```

#### (*Cache) KeyStrings

```go
func (c *Cache) KeyStrings(ctx context.Context) ([]string, error)
```

KeyStrings returns all keys in the cache as string slice.

​	KeyStrings 将缓存中的所有键作为字符串切片返回。

##### Example

``` go
```

#### (*Cache) MustContains

```go
func (c *Cache) MustContains(ctx context.Context, key interface{}) bool
```

MustContains acts like Contains, but it panics if any error occurs.

​	MustContains 的行为类似于 Contains，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Cache) MustData

```go
func (c *Cache) MustData(ctx context.Context) map[interface{}]interface{}
```

MustData acts like Data, but it panics if any error occurs.

​	MustData 的行为类似于 Data，但如果发生任何错误，它就会崩溃。

##### Example

``` go
```

#### (*Cache) MustGet

```go
func (c *Cache) MustGet(ctx context.Context, key interface{}) *gvar.Var
```

MustGet acts like Get, but it panics if any error occurs.

​	MustGet 的行为类似于 Get，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Cache) MustGetExpire

```go
func (c *Cache) MustGetExpire(ctx context.Context, key interface{}) time.Duration
```

MustGetExpire acts like GetExpire, but it panics if any error occurs.

​	MustGetExpire 的行为类似于 GetExpire，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Cache) MustGetOrSet

```go
func (c *Cache) MustGetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) *gvar.Var
```

MustGetOrSet acts like GetOrSet, but it panics if any error occurs.

​	MustGetOrSet 的行为类似于 GetOrSet，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Cache) MustGetOrSetFunc

```go
func (c *Cache) MustGetOrSetFunc(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var
```

MustGetOrSetFunc acts like GetOrSetFunc, but it panics if any error occurs.

​	MustGetOrSetFunc 的行为类似于 GetOrSetFunc，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Cache) MustGetOrSetFuncLock

```go
func (c *Cache) MustGetOrSetFuncLock(ctx context.Context, key interface{}, f Func, duration time.Duration) *gvar.Var
```

MustGetOrSetFuncLock acts like GetOrSetFuncLock, but it panics if any error occurs.

​	MustGetOrSetFuncLock 的行为类似于 GetOrSetFuncLock，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Cache) MustKeyStrings

```go
func (c *Cache) MustKeyStrings(ctx context.Context) []string
```

MustKeyStrings acts like KeyStrings, but it panics if any error occurs.

​	MustKeyStrings 的行为类似于 KeyStrings，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Cache) MustKeys

```go
func (c *Cache) MustKeys(ctx context.Context) []interface{}
```

MustKeys acts like Keys, but it panics if any error occurs.

​	MustKeys 的行为类似于 Keys，但如果发生任何错误，它会崩溃。

##### Example

``` go
```

#### (*Cache) MustSize

```go
func (c *Cache) MustSize(ctx context.Context) int
```

MustSize acts like Size, but it panics if any error occurs.

​	MustSize 的作用类似于 Size，但如果发生任何错误，它就会崩溃。

##### Example

``` go
```

#### (*Cache) MustValues

```go
func (c *Cache) MustValues(ctx context.Context) []interface{}
```

MustValues acts like Values, but it panics if any error occurs.

​	MustValues 的作用类似于 Values，但如果发生任何错误，它就会崩溃。

##### Example

``` go
```

#### (*Cache) Removes

```go
func (c *Cache) Removes(ctx context.Context, keys []interface{}) error
```

Removes deletes `keys` in the cache.

​	删除缓存 `keys` 中的删除内容。

##### Example

``` go
```

#### (*Cache) SetAdapter

```go
func (c *Cache) SetAdapter(adapter Adapter)
```

SetAdapter changes the adapter for this cache. Be very note that, this setting function is not concurrent-safe, which means you should not call this setting function concurrently in multiple goroutines.

​	SetAdapter 更改此缓存的适配器。需要注意的是，此设置函数不是并发安全的，这意味着您不应该在多个 goroutine 中并发调用此设置函数。

##### Example

``` go
```

### type Func

```go
type Func func(ctx context.Context) (value interface{}, err error)
```

Func is the cache function that calculates and returns the value.

​	Func 是计算并返回值的缓存函数。