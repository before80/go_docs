+++
title = "gsession"
date = 2024-03-21T17:57:06+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gsession](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gsession)

Package gsession implements manager and storage features for sessions.

​	软件包 gsession 实现了会话的管理器和存储功能。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gsession/gsession_storage_file.go#L38)

```go
const (
	DefaultStorageFileCryptoEnabled        = false
	DefaultStorageFileUpdateTTLInterval    = 10 * time.Second
	DefaultStorageFileClearExpiredInterval = time.Hour
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gsession/gsession_storage_redis.go#L28)

```go
const (
	// DefaultStorageRedisLoopInterval is the interval updating TTL for session ids
	// in last duration.
	DefaultStorageRedisLoopInterval = 10 * time.Second
)
```

## 变量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gsession/gsession_storage_file.go#L44)

```go
var (
	DefaultStorageFilePath      = gfile.Temp("gsessions")
	DefaultStorageFileCryptoKey = []byte("Session storage file crypto key!")
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gsession/gsession.go#L16)

```go
var (
	// ErrorDisabled is used for marking certain interface function not used.
	ErrorDisabled = gerror.NewWithOption(gerror.Option{
		Text: "this feature is disabled in this storage",
		Code: gcode.CodeNotSupported,
	})
)
```

## 函数

#### func NewSessionId

```go
func NewSessionId() string
```

NewSessionId creates and returns a new and unique session id string, which is in 32 bytes.

​	NewSessionId 创建并返回一个新的唯一会话 ID 字符串，该字符串以 32 个字节为单位。

## 类型

### type Manager

```go
type Manager struct {
	// contains filtered or unexported fields
}
```

Manager for sessions.

​	会话经理。

#### func New

```go
func New(ttl time.Duration, storage ...Storage) *Manager
```

New creates and returns a new session manager.

​	new 创建并返回新的会话管理器。

##### Example

``` go
```

#### (*Manager) GetStorage

```go
func (m *Manager) GetStorage() Storage
```

GetStorage returns the session storage of current manager.

​	GetStorage 返回当前管理器的会话存储。

##### Example

``` go
```

#### (*Manager) GetTTL

```go
func (m *Manager) GetTTL() time.Duration
```

GetTTL returns the TTL of the session manager.

​	GetTTL 返回会话管理器的 TTL。

#### (*Manager) New

```go
func (m *Manager) New(ctx context.Context, sessionId ...string) *Session
```

New creates or fetches the session for given session id. The parameter `sessionId` is optional, it creates a new one if not it’s passed depending on Storage.New.

​	new 为给定会话 ID 创建或获取会话。该参数 `sessionId` 是可选的，如果不是，它会创建一个新参数，具体取决于 Storage.New。

#### (*Manager) SetStorage

```go
func (m *Manager) SetStorage(storage Storage)
```

SetStorage sets the session storage for manager.

​	SetStorage 设置管理器的会话存储。

##### Example

``` go
```

#### (*Manager) SetTTL

```go
func (m *Manager) SetTTL(ttl time.Duration)
```

SetTTL the TTL for the session manager.

​	SetTTL 会话管理器的 TTL。

##### Example

``` go
```

### type Session

```go
type Session struct {
	// contains filtered or unexported fields
}
```

Session struct for storing single session data, which is bound to a single request. The Session struct is the interface with user, but the Storage is the underlying adapter designed interface for functionality implements.

​	用于存储绑定到单个请求的单个会话数据的会话结构。Session 结构是与用户的接口，而 Storage 是为功能实现设计的底层适配器接口。

#### (*Session) Close

```go
func (s *Session) Close() error
```

Close closes current session and updates its ttl in the session manager. If this session is dirty, it also exports it to storage.

​	关闭关闭当前会话并在会话管理器中更新其 ttl。如果此会话是脏的，它还会将其导出到存储。

NOTE that this function must be called ever after a session request done.

​	请注意，必须在会话请求完成后调用此函数。

#### (*Session) Contains

```go
func (s *Session) Contains(key string) (ok bool, err error)
```

Contains checks whether key exist in the session.

​	包含检查会话中是否存在密钥。

##### Example

``` go
```

#### (*Session) Data

```go
func (s *Session) Data() (sessionData map[string]interface{}, err error)
```

Data returns all data as map. Note that it’s using value copy internally for concurrent-safe purpose.

​	数据以地图的形式返回所有数据。请注意，它在内部使用值复制以实现并发安全目的。

##### Example

``` go
```

#### (*Session) Get

```go
func (s *Session) Get(key string, def ...interface{}) (value *gvar.Var, err error)
```

Get retrieves session value with given key. It returns `def` if the key does not exist in the session if `def` is given, or else it returns nil.

​	Get 使用给定键检索会话值。如果给定的密钥在会话 `def` 中不存在，则返回 `def` ，否则返回 nil。

#### (*Session) Id

```go
func (s *Session) Id() (id string, err error)
```

Id returns the session id for this session. It creates and returns a new session id if the session id is not passed in initialization.

​	Id 返回此会话的会话 ID。如果在初始化中未传递会话 ID，则会创建并返回新的会话 ID。

##### Example

``` go
```

#### (*Session) IsDirty

```go
func (s *Session) IsDirty() bool
```

IsDirty checks whether there’s any data changes in the session.

​	IsDirty 检查会话中是否有任何数据更改。

#### (*Session) MustContains

```go
func (s *Session) MustContains(key string) bool
```

MustContains performs as function Contains, but it panics if any error occurs.

​	MustContains 与函数 Contains 一样执行，但如果发生任何错误，它会崩溃。

#### (*Session) MustData

```go
func (s *Session) MustData() map[string]interface{}
```

MustData performs as function Data, but it panics if any error occurs.

​	MustData 作为函数 Data 执行，但如果发生任何错误，它会崩溃。

#### (*Session) MustGet

```go
func (s *Session) MustGet(key string, def ...interface{}) *gvar.Var
```

MustGet performs as function Get, but it panics if any error occurs.

​	MustGet 作为函数 Get 执行，但如果发生任何错误，它会崩溃。

#### (*Session) MustId

```go
func (s *Session) MustId() string
```

MustId performs as function Id, but it panics if any error occurs.

​	MustId 作为函数 Id 执行，但如果发生任何错误，它会崩溃。

#### (*Session) MustRemove

```go
func (s *Session) MustRemove(keys ...string)
```

MustRemove performs as function Remove, but it panics if any error occurs.

​	MustRemove 作为函数 Remove 执行，但如果发生任何错误，它会崩溃。

#### (*Session) MustSet

```go
func (s *Session) MustSet(key string, value interface{})
```

MustSet performs as function Set, but it panics if any error occurs.

​	MustSet 作为函数 Set 执行，但如果发生任何错误，它会崩溃。

#### (*Session) MustSetMap

```go
func (s *Session) MustSetMap(data map[string]interface{})
```

MustSetMap performs as function SetMap, but it panics if any error occurs.

​	MustSetMap 作为函数 SetMap 执行，但如果发生任何错误，它会崩溃。

#### (*Session) MustSize

```go
func (s *Session) MustSize() int
```

MustSize performs as function Size, but it panics if any error occurs.

​	MustSize 作为函数 Size 执行，但如果发生任何错误，它会崩溃。

#### (*Session) Remove

```go
func (s *Session) Remove(keys ...string) (err error)
```

Remove removes key along with its value from this session.

​	“删除”（Remove） 从此会话中删除键及其值。

##### Example

``` go
```

#### (*Session) RemoveAll

```go
func (s *Session) RemoveAll() (err error)
```

RemoveAll deletes all key-value pairs from this session.

​	RemoveAll 从此会话中删除所有键值对。

##### Example

``` go
```

#### (*Session) Set

```go
func (s *Session) Set(key string, value interface{}) (err error)
```

Set sets key-value pair to this session.

​	Set 将键值对设置为此会话。

##### Example

``` go
```

#### (*Session) SetId

```go
func (s *Session) SetId(id string) error
```

SetId sets custom session before session starts. It returns error if it is called after session starts.

​	SetId 在会话开始之前设置自定义会话。如果在会话启动后调用它，则返回错误。

##### Example

``` go
```

#### (*Session) SetIdFunc

```go
func (s *Session) SetIdFunc(f func(ttl time.Duration) string) error
```

SetIdFunc sets custom session id creating function before session starts. It returns error if it is called after session starts.

​	SetIdFunc 在会话开始前设置自定义会话 ID 创建函数。如果在会话启动后调用它，则返回错误。

##### Example

``` go
```

#### (*Session) SetMap

```go
func (s *Session) SetMap(data map[string]interface{}) (err error)
```

SetMap batch sets the session using map.

​	SetMap 批处理使用 map 设置会话。

##### Example

``` go
```

#### (*Session) Size

```go
func (s *Session) Size() (size int, err error)
```

Size returns the size of the session.

​	Size 返回会话的大小。

##### Example

``` go
```

### type Storage

```go
type Storage interface {
	// New creates a custom session id.
	// This function can be used for custom session creation.
	New(ctx context.Context, ttl time.Duration) (sessionId string, err error)

	// Get retrieves and returns certain session value with given key.
	// It returns nil if the key does not exist in the session.
	Get(ctx context.Context, sessionId string, key string) (value interface{}, err error)

	// GetSize retrieves and returns the size of key-value pairs from storage.
	GetSize(ctx context.Context, sessionId string) (size int, err error)

	// Data retrieves all key-value pairs as map from storage.
	Data(ctx context.Context, sessionId string) (sessionData map[string]interface{}, err error)

	// Set sets one key-value session pair to the storage.
	// The parameter `ttl` specifies the TTL for the session id.
	Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error

	// SetMap batch sets key-value session pairs as map to the storage.
	// The parameter `ttl` specifies the TTL for the session id.
	SetMap(ctx context.Context, sessionId string, mapData map[string]interface{}, ttl time.Duration) error

	// Remove deletes key-value pair from specified session from storage.
	Remove(ctx context.Context, sessionId string, key string) error

	// RemoveAll deletes session from storage.
	RemoveAll(ctx context.Context, sessionId string) error

	// GetSession returns the session data as `*gmap.StrAnyMap` for given session from storage.
	//
	// The parameter `ttl` specifies the TTL for this session.
	// The parameter `data` is the current old session data stored in memory,
	// and for some storage it might be nil if memory storage is disabled.
	//
	// This function is called ever when session starts.
	// It returns nil if the session does not exist or its TTL is expired.
	GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)

	// SetSession updates the data for specified session id.
	// This function is called ever after session, which is changed dirty, is closed.
	// This copy all session data map from memory to storage.
	SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error

	// UpdateTTL updates the TTL for specified session id.
	// This function is called ever after session, which is not dirty, is closed.
	UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
}
```

Storage is the interface definition for session storage.

​	存储是会话存储的接口定义。

### type StorageBase <-2.1.0

```go
type StorageBase struct{}
```

StorageBase is a base implement for Session Storage.

​	StorageBase 是会话存储的基本实现。

#### (*StorageBase) Data

```go
func (s *StorageBase) Data(ctx context.Context, sessionId string) (sessionData map[string]interface{}, err error)
```

Data retrieves all key-value pairs as map from storage.

​	数据从存储中检索所有键值对作为映射。

#### (*StorageBase) Get

```go
func (s *StorageBase) Get(ctx context.Context, sessionId string, key string) (value interface{}, err error)
```

Get retrieves certain session value with given key. It returns nil if the key does not exist in the session.

​	Get 使用给定的密钥检索特定的会话值。如果会话中不存在密钥，则返回 nil。

#### (*StorageBase) GetSession

```go
func (s *StorageBase) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)
```

GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.

​	GetSession 以 *gmap 格式返回会话数据。StrAnyMap 用于存储中的给定会话 ID。

The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded. The parameter `data` is the current old session data stored in memory, and for some storage it might be nil if memory storage is disabled.

​	该参数 `ttl` 指定此会话的 TTL，如果超过 TTL，则返回 nil。该参数 `data` 是存储在内存中的当前旧会话数据，对于某些存储，如果禁用内存存储，则该参数可能为零。

This function is called ever when session starts.

​	此函数在会话开始时调用。

#### (*StorageBase) GetSize

```go
func (s *StorageBase) GetSize(ctx context.Context, sessionId string) (size int, err error)
```

GetSize retrieves the size of key-value pairs from storage.

​	GetSize 从存储中检索键值对的大小。

#### (*StorageBase) New

```go
func (s *StorageBase) New(ctx context.Context, ttl time.Duration) (id string, err error)
```

New creates a session id. This function can be used for custom session creation.

​	New 将创建会话 ID。此函数可用于创建自定义会话。

#### (*StorageBase) Remove

```go
func (s *StorageBase) Remove(ctx context.Context, sessionId string, key string) error
```

Remove deletes key with its value from storage.

​	从存储中删除具有其值的删除键。

#### (*StorageBase) RemoveAll

```go
func (s *StorageBase) RemoveAll(ctx context.Context, sessionId string) error
```

RemoveAll deletes session from storage.

​	RemoveAll 从存储中删除会话。

#### (*StorageBase) Set

```go
func (s *StorageBase) Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error
```

Set sets key-value session pair to the storage. The parameter `ttl` specifies the TTL for the session id (not for the key-value pair).

​	Set 将键值会话对设置为存储。该参数 `ttl` 指定会话 ID 的 TTL（而不是键值对的 TTL）。

#### (*StorageBase) SetMap

```go
func (s *StorageBase) SetMap(ctx context.Context, sessionId string, mapData map[string]interface{}, ttl time.Duration) error
```

SetMap batch sets key-value session pairs with map to the storage. The parameter `ttl` specifies the TTL for the session id(not for the key-value pair).

​	SetMap 批处理键值会话对，并映射到存储。该参数 `ttl` 指定会话 ID 的 TTL（而不是键值对的 TTL）。

#### (*StorageBase) SetSession

```go
func (s *StorageBase) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error
```

SetSession updates the data map for specified session id. This function is called ever after session, which is changed dirty, is closed. This copy all session data map from memory to storage.

​	SetSession 更新指定会话 ID 的数据映射。此函数在会话后调用，该会话已更改为脏，已关闭。这会将所有会话数据映射从内存复制到存储。

#### (*StorageBase) UpdateTTL

```go
func (s *StorageBase) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
```

UpdateTTL updates the TTL for specified session id. This function is called ever after session, which is not dirty, is closed. It just adds the session id to the async handling queue.

​	UpdateTTL 更新指定会话 ID 的 TTL。此函数在会话关闭后调用，该会话不脏。它只是将会话 ID 添加到异步处理队列。

### type StorageFile

```go
type StorageFile struct {
	StorageBase
	// contains filtered or unexported fields
}
```

StorageFile implements the Session Storage interface with file system.

​	StorageFile 使用文件系统实现会话存储接口。

#### func NewStorageFile

```go
func NewStorageFile(path string, ttl time.Duration) *StorageFile
```

NewStorageFile creates and returns a file storage object for session.

​	NewStorageFile 创建并返回会话的文件存储对象。

#### (*StorageFile) GetSession

```go
func (s *StorageFile) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (sessionData *gmap.StrAnyMap, err error)
```

GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.

​	GetSession 以 *gmap 格式返回会话数据。StrAnyMap 用于存储中的给定会话 ID。

The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded. The parameter `data` is the current old session data stored in memory, and for some storage it might be nil if memory storage is disabled.

​	该参数 `ttl` 指定此会话的 TTL，如果超过 TTL，则返回 nil。该参数 `data` 是存储在内存中的当前旧会话数据，对于某些存储，如果禁用内存存储，则该参数可能为零。

This function is called ever when session starts.

​	此函数在会话开始时调用。

#### (*StorageFile) RemoveAll

```go
func (s *StorageFile) RemoveAll(ctx context.Context, sessionId string) error
```

RemoveAll deletes all key-value pairs from storage.

​	RemoveAll 从存储中删除所有键值对。

#### (*StorageFile) SetCryptoEnabled

```go
func (s *StorageFile) SetCryptoEnabled(enabled bool)
```

SetCryptoEnabled enables/disables the crypto feature for session storage.

​	SetCryptoEnabled 启用/禁用会话存储的加密功能。

##### Example

``` go
```

#### (*StorageFile) SetCryptoKey

```go
func (s *StorageFile) SetCryptoKey(key []byte)
```

SetCryptoKey sets the crypto key for session storage. The crypto key is used when crypto feature is enabled.

​	SetCryptoKey 设置会话存储的加密密钥。启用加密功能时使用加密密钥。

##### Example

``` go
```

#### (*StorageFile) SetSession

```go
func (s *StorageFile) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error
```

SetSession updates the data map for specified session id. This function is called ever after session, which is changed dirty, is closed. This copy all session data map from memory to storage.

​	SetSession 更新指定会话 ID 的数据映射。此函数在会话后调用，该会话已更改为脏，已关闭。这会将所有会话数据映射从内存复制到存储。

#### (*StorageFile) UpdateTTL

```go
func (s *StorageFile) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
```

UpdateTTL updates the TTL for specified session id. This function is called ever after session, which is not dirty, is closed. It just adds the session id to the async handling queue.

​	UpdateTTL 更新指定会话 ID 的 TTL。此函数在会话关闭后调用，该会话不脏。它只是将会话 ID 添加到异步处理队列。

##### Example

``` go
```

### type StorageMemory

```go
type StorageMemory struct {
	StorageBase
	// contains filtered or unexported fields
}
```

StorageMemory implements the Session Storage interface with memory.

​	StorageMemory 使用内存实现会话存储接口。

#### func NewStorageMemory

```go
func NewStorageMemory() *StorageMemory
```

NewStorageMemory creates and returns a file storage object for session.

​	NewStorageMemory 创建并返回会话的文件存储对象。

#### (*StorageMemory) GetSession

```go
func (s *StorageMemory) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)
```

GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.

​	GetSession 以 *gmap 格式返回会话数据。StrAnyMap 用于存储中的给定会话 ID。

The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded. The parameter `data` is the current old session data stored in memory, and for some storage it might be nil if memory storage is disabled.

​	该参数 `ttl` 指定此会话的 TTL，如果超过 TTL，则返回 nil。该参数 `data` 是存储在内存中的当前旧会话数据，对于某些存储，如果禁用内存存储，则该参数可能为零。

This function is called ever when session starts.

​	此函数在会话开始时调用。

#### (*StorageMemory) RemoveAll

```go
func (s *StorageMemory) RemoveAll(ctx context.Context, sessionId string) error
```

RemoveAll deletes session from storage.

​	RemoveAll 从存储中删除会话。

#### (*StorageMemory) SetSession

```go
func (s *StorageMemory) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error
```

SetSession updates the data map for specified session id. This function is called ever after session, which is changed dirty, is closed. This copy all session data map from memory to storage.

​	SetSession 更新指定会话 ID 的数据映射。此函数在会话后调用，该会话已更改为脏，已关闭。这会将所有会话数据映射从内存复制到存储。

#### (*StorageMemory) UpdateTTL

```go
func (s *StorageMemory) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
```

UpdateTTL updates the TTL for specified session id. This function is called ever after session, which is not dirty, is closed. It just adds the session id to the async handling queue.

​	UpdateTTL 更新指定会话 ID 的 TTL。此函数在会话关闭后调用，该会话不脏。它只是将会话 ID 添加到异步处理队列。

### type StorageRedis

```go
type StorageRedis struct {
	StorageBase
	// contains filtered or unexported fields
}
```

StorageRedis implements the Session Storage interface with redis.

​	StorageRedis 使用 redis 实现会话存储接口。

#### func NewStorageRedis

```go
func NewStorageRedis(redis *gredis.Redis, prefix ...string) *StorageRedis
```

NewStorageRedis creates and returns a redis storage object for session.

​	NewStorageRedis 创建并返回会话的 redis 存储对象。

#### (*StorageRedis) GetSession

```go
func (s *StorageRedis) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)
```

GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.

​	GetSession 以 *gmap 格式返回会话数据。StrAnyMap 用于存储中的给定会话 ID。

The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded. The parameter `data` is the current old session data stored in memory, and for some storage it might be nil if memory storage is disabled.

​	该参数 `ttl` 指定此会话的 TTL，如果超过 TTL，则返回 nil。该参数 `data` 是存储在内存中的当前旧会话数据，对于某些存储，如果禁用内存存储，则该参数可能为零。

This function is called ever when session starts.

​	此函数在会话开始时调用。

#### (*StorageRedis) RemoveAll

```go
func (s *StorageRedis) RemoveAll(ctx context.Context, sessionId string) error
```

RemoveAll deletes all key-value pairs from storage.

​	RemoveAll 从存储中删除所有键值对。

##### Example

``` go
```

#### (*StorageRedis) SetSession

```go
func (s *StorageRedis) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error
```

SetSession updates the data map for specified session id. This function is called ever after session, which is changed dirty, is closed. This copy all session data map from memory to storage.

​	SetSession 更新指定会话 ID 的数据映射。此函数在会话后调用，该会话已更改为脏，已关闭。这会将所有会话数据映射从内存复制到存储。

#### (*StorageRedis) UpdateTTL

```go
func (s *StorageRedis) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
```

UpdateTTL updates the TTL for specified session id. This function is called ever after session, which is not dirty, is closed. It just adds the session id to the async handling queue.

​	UpdateTTL 更新指定会话 ID 的 TTL。此函数在会话关闭后调用，该会话不脏。它只是将会话 ID 添加到异步处理队列。

##### Example

``` go
```

### type StorageRedisHashTable

```go
type StorageRedisHashTable struct {
	StorageBase
	// contains filtered or unexported fields
}
```

StorageRedisHashTable implements the Session Storage interface with redis hash table.

​	StorageRedisHashTable 使用 redis 哈希表实现会话存储接口。

#### func NewStorageRedisHashTable

```go
func NewStorageRedisHashTable(redis *gredis.Redis, prefix ...string) *StorageRedisHashTable
```

NewStorageRedisHashTable creates and returns a redis hash table storage object for session.

​	NewStorageRedisHashTable 创建并返回会话的 redis 哈希表存储对象。

#### (*StorageRedisHashTable) Data

```go
func (s *StorageRedisHashTable) Data(ctx context.Context, sessionId string) (data map[string]interface{}, err error)
```

Data retrieves all key-value pairs as map from storage.

​	数据从存储中检索所有键值对作为映射。

##### Example

``` go
```

#### (*StorageRedisHashTable) Get

```go
func (s *StorageRedisHashTable) Get(ctx context.Context, sessionId string, key string) (value interface{}, err error)
```

Get retrieves session value with given key. It returns nil if the key does not exist in the session.

​	Get 使用给定键检索会话值。如果会话中不存在密钥，则返回 nil。

##### Example

``` go
```

#### (*StorageRedisHashTable) GetSession

```go
func (s *StorageRedisHashTable) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)
```

GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.

​	GetSession 以 *gmap 格式返回会话数据。StrAnyMap 用于存储中的给定会话 ID。

The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded. The parameter `data` is the current old session data stored in memory, and for some storage it might be nil if memory storage is disabled.

​	该参数 `ttl` 指定此会话的 TTL，如果超过 TTL，则返回 nil。该参数 `data` 是存储在内存中的当前旧会话数据，对于某些存储，如果禁用内存存储，则该参数可能为零。

This function is called ever when session starts.

​	此函数在会话开始时调用。

##### Example

``` go
```

#### (*StorageRedisHashTable) GetSize

```go
func (s *StorageRedisHashTable) GetSize(ctx context.Context, sessionId string) (size int, err error)
```

GetSize retrieves the size of key-value pairs from storage.

​	GetSize 从存储中检索键值对的大小。

##### Example

``` go
```

#### (*StorageRedisHashTable) Remove

```go
func (s *StorageRedisHashTable) Remove(ctx context.Context, sessionId string, key string) error
```

Remove deletes key with its value from storage.

​	从存储中删除具有其值的删除键。

##### Example

``` go
```

#### (*StorageRedisHashTable) RemoveAll

```go
func (s *StorageRedisHashTable) RemoveAll(ctx context.Context, sessionId string) error
```

RemoveAll deletes all key-value pairs from storage.

​	RemoveAll 从存储中删除所有键值对。

##### Example

``` go
```

#### (*StorageRedisHashTable) Set

```go
func (s *StorageRedisHashTable) Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error
```

Set sets key-value session pair to the storage. The parameter `ttl` specifies the TTL for the session id (not for the key-value pair).

​	Set 将键值会话对设置为存储。该参数 `ttl` 指定会话 ID 的 TTL（而不是键值对的 TTL）。

#### (*StorageRedisHashTable) SetMap

```go
func (s *StorageRedisHashTable) SetMap(ctx context.Context, sessionId string, data map[string]interface{}, ttl time.Duration) error
```

SetMap batch sets key-value session pairs with map to the storage. The parameter `ttl` specifies the TTL for the session id(not for the key-value pair).

​	SetMap 批处理键值会话对，并映射到存储。该参数 `ttl` 指定会话 ID 的 TTL（而不是键值对的 TTL）。

#### (*StorageRedisHashTable) SetSession

```go
func (s *StorageRedisHashTable) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error
```

SetSession updates the data map for specified session id. This function is called ever after session, which is changed dirty, is closed. This copy all session data map from memory to storage.

​	SetSession 更新指定会话 ID 的数据映射。此函数在会话后调用，该会话已更改为脏，已关闭。这会将所有会话数据映射从内存复制到存储。

##### Example

``` go
```

#### (*StorageRedisHashTable) UpdateTTL

```go
func (s *StorageRedisHashTable) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
```

UpdateTTL updates the TTL for specified session id. This function is called ever after session, which is not dirty, is closed. It just adds the session id to the async handling queue.

​	UpdateTTL 更新指定会话 ID 的 TTL。此函数在会话关闭后调用，该会话不脏。它只是将会话 ID 添加到异步处理队列。

Example UpdateTTL

```go
package main

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gsession"
)

func main() {
	storage := gsession.NewStorageRedisHashTable(g.Redis())

	err := storage.UpdateTTL(gctx.New(), "id", time.Second)

	fmt.Println(err)

	// May Output:
	// redis adapter is not set, missing configuration or adapter register? possible reference: https://github.com/gogf/gf/tree/master/contrib/nosql/redis
}
Output:
```



