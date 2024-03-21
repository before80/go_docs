+++
title = "gsession"
date = 2024-03-21T17:57:06+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gsession

Package gsession implements manager and storage features for sessions.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gsession/gsession_storage_file.go#L38)

``` go
const (
	DefaultStorageFileCryptoEnabled        = false
	DefaultStorageFileUpdateTTLInterval    = 10 * time.Second
	DefaultStorageFileClearExpiredInterval = time.Hour
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gsession/gsession_storage_redis.go#L28)

``` go
const (
	// DefaultStorageRedisLoopInterval is the interval updating TTL for session ids
	// in last duration.
	DefaultStorageRedisLoopInterval = 10 * time.Second
)
```

### Variables 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gsession/gsession_storage_file.go#L44)

``` go
var (
	DefaultStorageFilePath      = gfile.Temp("gsessions")
	DefaultStorageFileCryptoKey = []byte("Session storage file crypto key!")
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gsession/gsession.go#L16)

``` go
var (
	// ErrorDisabled is used for marking certain interface function not used.
	ErrorDisabled = gerror.NewWithOption(gerror.Option{
		Text: "this feature is disabled in this storage",
		Code: gcode.CodeNotSupported,
	})
)
```

### Functions 

##### func NewSessionId 

``` go
func NewSessionId() string
```

NewSessionId creates and returns a new and unique session id string, which is in 32 bytes.

### Types 

#### type Manager 

``` go
type Manager struct {
	// contains filtered or unexported fields
}
```

Manager for sessions.

##### func New 

``` go
func New(ttl time.Duration, storage ...Storage) *Manager
```

New creates and returns a new session manager.

##### Example

``` go
```
##### (*Manager) GetStorage 

``` go
func (m *Manager) GetStorage() Storage
```

GetStorage returns the session storage of current manager.

##### Example

``` go
```
##### (*Manager) GetTTL <-2.1.0

``` go
func (m *Manager) GetTTL() time.Duration
```

GetTTL returns the TTL of the session manager.

##### (*Manager) New 

``` go
func (m *Manager) New(ctx context.Context, sessionId ...string) *Session
```

New creates or fetches the session for given session id. The parameter `sessionId` is optional, it creates a new one if not it's passed depending on Storage.New.

##### (*Manager) SetStorage 

``` go
func (m *Manager) SetStorage(storage Storage)
```

SetStorage sets the session storage for manager.

##### Example

``` go
```
##### (*Manager) SetTTL 

``` go
func (m *Manager) SetTTL(ttl time.Duration)
```

SetTTL the TTL for the session manager.

##### Example

``` go
```
#### type Session 

``` go
type Session struct {
	// contains filtered or unexported fields
}
```

Session struct for storing single session data, which is bound to a single request. The Session struct is the interface with user, but the Storage is the underlying adapter designed interface for functionality implements.

##### (*Session) Close 

``` go
func (s *Session) Close() error
```

Close closes current session and updates its ttl in the session manager. If this session is dirty, it also exports it to storage.

NOTE that this function must be called ever after a session request done.

##### (*Session) Contains 

``` go
func (s *Session) Contains(key string) (ok bool, err error)
```

Contains checks whether key exist in the session.

##### Example

``` go
```
##### (*Session) Data 

``` go
func (s *Session) Data() (sessionData map[string]interface{}, err error)
```

Data returns all data as map. Note that it's using value copy internally for concurrent-safe purpose.

##### Example

``` go
```
##### (*Session) Get 

``` go
func (s *Session) Get(key string, def ...interface{}) (value *gvar.Var, err error)
```

Get retrieves session value with given key. It returns `def` if the key does not exist in the session if `def` is given, or else it returns nil.

##### (*Session) Id 

``` go
func (s *Session) Id() (id string, err error)
```

Id returns the session id for this session. It creates and returns a new session id if the session id is not passed in initialization.

##### Example

``` go
```
##### (*Session) IsDirty 

``` go
func (s *Session) IsDirty() bool
```

IsDirty checks whether there's any data changes in the session.

##### (*Session) MustContains 

``` go
func (s *Session) MustContains(key string) bool
```

MustContains performs as function Contains, but it panics if any error occurs.

##### (*Session) MustData 

``` go
func (s *Session) MustData() map[string]interface{}
```

MustData performs as function Data, but it panics if any error occurs.

##### (*Session) MustGet 

``` go
func (s *Session) MustGet(key string, def ...interface{}) *gvar.Var
```

MustGet performs as function Get, but it panics if any error occurs.

##### (*Session) MustId 

``` go
func (s *Session) MustId() string
```

MustId performs as function Id, but it panics if any error occurs.

##### (*Session) MustRemove 

``` go
func (s *Session) MustRemove(keys ...string)
```

MustRemove performs as function Remove, but it panics if any error occurs.

##### (*Session) MustSet 

``` go
func (s *Session) MustSet(key string, value interface{})
```

MustSet performs as function Set, but it panics if any error occurs.

##### (*Session) MustSetMap 

``` go
func (s *Session) MustSetMap(data map[string]interface{})
```

MustSetMap performs as function SetMap, but it panics if any error occurs.

##### (*Session) MustSize 

``` go
func (s *Session) MustSize() int
```

MustSize performs as function Size, but it panics if any error occurs.

##### (*Session) Remove 

``` go
func (s *Session) Remove(keys ...string) (err error)
```

Remove removes key along with its value from this session.

##### Example

``` go
```
##### (*Session) RemoveAll 

``` go
func (s *Session) RemoveAll() (err error)
```

RemoveAll deletes all key-value pairs from this session.

##### Example

``` go
```
##### (*Session) Set 

``` go
func (s *Session) Set(key string, value interface{}) (err error)
```

Set sets key-value pair to this session.

##### Example

``` go
```
##### (*Session) SetId 

``` go
func (s *Session) SetId(id string) error
```

SetId sets custom session before session starts. It returns error if it is called after session starts.

##### Example

``` go
```
##### (*Session) SetIdFunc 

``` go
func (s *Session) SetIdFunc(f func(ttl time.Duration) string) error
```

SetIdFunc sets custom session id creating function before session starts. It returns error if it is called after session starts.

##### Example

``` go
```
##### (*Session) SetMap 

``` go
func (s *Session) SetMap(data map[string]interface{}) (err error)
```

SetMap batch sets the session using map.

##### Example

``` go
```
##### (*Session) Size 

``` go
func (s *Session) Size() (size int, err error)
```

Size returns the size of the session.

##### Example

``` go
```
#### type Storage 

``` go
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

#### type StorageBase <-2.1.0

``` go
type StorageBase struct{}
```

StorageBase is a base implement for Session Storage.

##### (*StorageBase) Data <-2.1.0

``` go
func (s *StorageBase) Data(ctx context.Context, sessionId string) (sessionData map[string]interface{}, err error)
```

Data retrieves all key-value pairs as map from storage.

##### (*StorageBase) Get <-2.1.0

``` go
func (s *StorageBase) Get(ctx context.Context, sessionId string, key string) (value interface{}, err error)
```

Get retrieves certain session value with given key. It returns nil if the key does not exist in the session.

##### (*StorageBase) GetSession <-2.1.0

``` go
func (s *StorageBase) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)
```

GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.

The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded. The parameter `data` is the current old session data stored in memory, and for some storage it might be nil if memory storage is disabled.

This function is called ever when session starts.

##### (*StorageBase) GetSize <-2.1.0

``` go
func (s *StorageBase) GetSize(ctx context.Context, sessionId string) (size int, err error)
```

GetSize retrieves the size of key-value pairs from storage.

##### (*StorageBase) New <-2.1.0

``` go
func (s *StorageBase) New(ctx context.Context, ttl time.Duration) (id string, err error)
```

New creates a session id. This function can be used for custom session creation.

##### (*StorageBase) Remove <-2.1.0

``` go
func (s *StorageBase) Remove(ctx context.Context, sessionId string, key string) error
```

Remove deletes key with its value from storage.

##### (*StorageBase) RemoveAll <-2.1.0

``` go
func (s *StorageBase) RemoveAll(ctx context.Context, sessionId string) error
```

RemoveAll deletes session from storage.

##### (*StorageBase) Set <-2.1.0

``` go
func (s *StorageBase) Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error
```

Set sets key-value session pair to the storage. The parameter `ttl` specifies the TTL for the session id (not for the key-value pair).

##### (*StorageBase) SetMap <-2.1.0

``` go
func (s *StorageBase) SetMap(ctx context.Context, sessionId string, mapData map[string]interface{}, ttl time.Duration) error
```

SetMap batch sets key-value session pairs with map to the storage. The parameter `ttl` specifies the TTL for the session id(not for the key-value pair).

##### (*StorageBase) SetSession <-2.1.0

``` go
func (s *StorageBase) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error
```

SetSession updates the data map for specified session id. This function is called ever after session, which is changed dirty, is closed. This copy all session data map from memory to storage.

##### (*StorageBase) UpdateTTL <-2.1.0

``` go
func (s *StorageBase) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
```

UpdateTTL updates the TTL for specified session id. This function is called ever after session, which is not dirty, is closed. It just adds the session id to the async handling queue.

#### type StorageFile 

``` go
type StorageFile struct {
	StorageBase
	// contains filtered or unexported fields
}
```

StorageFile implements the Session Storage interface with file system.

##### func NewStorageFile 

``` go
func NewStorageFile(path string, ttl time.Duration) *StorageFile
```

NewStorageFile creates and returns a file storage object for session.

##### (*StorageFile) GetSession 

``` go
func (s *StorageFile) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (sessionData *gmap.StrAnyMap, err error)
```

GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.

The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded. The parameter `data` is the current old session data stored in memory, and for some storage it might be nil if memory storage is disabled.

This function is called ever when session starts.

##### (*StorageFile) RemoveAll 

``` go
func (s *StorageFile) RemoveAll(ctx context.Context, sessionId string) error
```

RemoveAll deletes all key-value pairs from storage.

##### (*StorageFile) SetCryptoEnabled 

``` go
func (s *StorageFile) SetCryptoEnabled(enabled bool)
```

SetCryptoEnabled enables/disables the crypto feature for session storage.

##### Example

``` go
```
##### (*StorageFile) SetCryptoKey 

``` go
func (s *StorageFile) SetCryptoKey(key []byte)
```

SetCryptoKey sets the crypto key for session storage. The crypto key is used when crypto feature is enabled.

##### Example

``` go
```
##### (*StorageFile) SetSession 

``` go
func (s *StorageFile) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error
```

SetSession updates the data map for specified session id. This function is called ever after session, which is changed dirty, is closed. This copy all session data map from memory to storage.

##### (*StorageFile) UpdateTTL 

``` go
func (s *StorageFile) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
```

UpdateTTL updates the TTL for specified session id. This function is called ever after session, which is not dirty, is closed. It just adds the session id to the async handling queue.

##### Example

``` go
```
#### type StorageMemory 

``` go
type StorageMemory struct {
	StorageBase
	// contains filtered or unexported fields
}
```

StorageMemory implements the Session Storage interface with memory.

##### func NewStorageMemory 

``` go
func NewStorageMemory() *StorageMemory
```

NewStorageMemory creates and returns a file storage object for session.

##### (*StorageMemory) GetSession 

``` go
func (s *StorageMemory) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)
```

GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.

The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded. The parameter `data` is the current old session data stored in memory, and for some storage it might be nil if memory storage is disabled.

This function is called ever when session starts.

##### (*StorageMemory) RemoveAll 

``` go
func (s *StorageMemory) RemoveAll(ctx context.Context, sessionId string) error
```

RemoveAll deletes session from storage.

##### (*StorageMemory) SetSession 

``` go
func (s *StorageMemory) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error
```

SetSession updates the data map for specified session id. This function is called ever after session, which is changed dirty, is closed. This copy all session data map from memory to storage.

##### (*StorageMemory) UpdateTTL 

``` go
func (s *StorageMemory) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
```

UpdateTTL updates the TTL for specified session id. This function is called ever after session, which is not dirty, is closed. It just adds the session id to the async handling queue.

#### type StorageRedis 

``` go
type StorageRedis struct {
	StorageBase
	// contains filtered or unexported fields
}
```

StorageRedis implements the Session Storage interface with redis.

##### func NewStorageRedis 

``` go
func NewStorageRedis(redis *gredis.Redis, prefix ...string) *StorageRedis
```

NewStorageRedis creates and returns a redis storage object for session.

##### (*StorageRedis) GetSession 

``` go
func (s *StorageRedis) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)
```

GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.

The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded. The parameter `data` is the current old session data stored in memory, and for some storage it might be nil if memory storage is disabled.

This function is called ever when session starts.

##### (*StorageRedis) RemoveAll 

``` go
func (s *StorageRedis) RemoveAll(ctx context.Context, sessionId string) error
```

RemoveAll deletes all key-value pairs from storage.

##### Example

``` go
```
##### (*StorageRedis) SetSession 

``` go
func (s *StorageRedis) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error
```

SetSession updates the data map for specified session id. This function is called ever after session, which is changed dirty, is closed. This copy all session data map from memory to storage.

##### (*StorageRedis) UpdateTTL 

``` go
func (s *StorageRedis) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
```

UpdateTTL updates the TTL for specified session id. This function is called ever after session, which is not dirty, is closed. It just adds the session id to the async handling queue.

##### Example

``` go
```
#### type StorageRedisHashTable 

``` go
type StorageRedisHashTable struct {
	StorageBase
	// contains filtered or unexported fields
}
```

StorageRedisHashTable implements the Session Storage interface with redis hash table.

##### func NewStorageRedisHashTable 

``` go
func NewStorageRedisHashTable(redis *gredis.Redis, prefix ...string) *StorageRedisHashTable
```

NewStorageRedisHashTable creates and returns a redis hash table storage object for session.

##### (*StorageRedisHashTable) Data 

``` go
func (s *StorageRedisHashTable) Data(ctx context.Context, sessionId string) (data map[string]interface{}, err error)
```

Data retrieves all key-value pairs as map from storage.

##### Example

``` go
```
##### (*StorageRedisHashTable) Get 

``` go
func (s *StorageRedisHashTable) Get(ctx context.Context, sessionId string, key string) (value interface{}, err error)
```

Get retrieves session value with given key. It returns nil if the key does not exist in the session.

##### Example

``` go
```
##### (*StorageRedisHashTable) GetSession 

``` go
func (s *StorageRedisHashTable) GetSession(ctx context.Context, sessionId string, ttl time.Duration) (*gmap.StrAnyMap, error)
```

GetSession returns the session data as *gmap.StrAnyMap for given session id from storage.

The parameter `ttl` specifies the TTL for this session, and it returns nil if the TTL is exceeded. The parameter `data` is the current old session data stored in memory, and for some storage it might be nil if memory storage is disabled.

This function is called ever when session starts.

##### Example

``` go
```
##### (*StorageRedisHashTable) GetSize 

``` go
func (s *StorageRedisHashTable) GetSize(ctx context.Context, sessionId string) (size int, err error)
```

GetSize retrieves the size of key-value pairs from storage.

##### Example

``` go
```
##### (*StorageRedisHashTable) Remove 

``` go
func (s *StorageRedisHashTable) Remove(ctx context.Context, sessionId string, key string) error
```

Remove deletes key with its value from storage.

##### Example

``` go
```
##### (*StorageRedisHashTable) RemoveAll 

``` go
func (s *StorageRedisHashTable) RemoveAll(ctx context.Context, sessionId string) error
```

RemoveAll deletes all key-value pairs from storage.

##### Example

``` go
```
##### (*StorageRedisHashTable) Set 

``` go
func (s *StorageRedisHashTable) Set(ctx context.Context, sessionId string, key string, value interface{}, ttl time.Duration) error
```

Set sets key-value session pair to the storage. The parameter `ttl` specifies the TTL for the session id (not for the key-value pair).

##### (*StorageRedisHashTable) SetMap 

``` go
func (s *StorageRedisHashTable) SetMap(ctx context.Context, sessionId string, data map[string]interface{}, ttl time.Duration) error
```

SetMap batch sets key-value session pairs with map to the storage. The parameter `ttl` specifies the TTL for the session id(not for the key-value pair).

##### (*StorageRedisHashTable) SetSession 

``` go
func (s *StorageRedisHashTable) SetSession(ctx context.Context, sessionId string, sessionData *gmap.StrAnyMap, ttl time.Duration) error
```

SetSession updates the data map for specified session id. This function is called ever after session, which is changed dirty, is closed. This copy all session data map from memory to storage.

##### Example

``` go
```
##### (*StorageRedisHashTable) UpdateTTL 

``` go
func (s *StorageRedisHashTable) UpdateTTL(ctx context.Context, sessionId string, ttl time.Duration) error
```

UpdateTTL updates the TTL for specified session id. This function is called ever after session, which is not dirty, is closed. It just adds the session id to the async handling queue.

Example UpdateTTL

``` go
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

