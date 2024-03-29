+++
title = "gdebug"
date = 2024-03-21T17:47:51+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/database/gredis](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/database/gredis)

Package gredis provides convenient client for redis server.

​	软件包 gredis 为 redis 服务器提供了便捷的客户端。

Redis Client.

​	Redis 客户端。

Redis Commands Official: https://redis.io/commands

​	Redis 命令官方：https://redis.io/commands

Redis Chinese Documentation: http://redisdoc.com/

​	Redis 中文文档：http://redisdoc.com/

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/database/gredis/gredis_config.go#L48)

```go
const (
	DefaultGroupName = "default" // Default configuration group name.
)
```

## 变量

This section is empty.

## 函数

#### func ClearConfig

```go
func ClearConfig()
```

ClearConfig removes all configurations of redis.

​	ClearConfig 删除 redis 的所有配置。

#### func RegisterAdapterFunc <-2.3.0

```go
func RegisterAdapterFunc(adapterFunc AdapterFunc)
```

RegisterAdapterFunc registers default function creating redis adapter.

​	RegisterAdapterFunc 注册创建 redis 适配器的默认函数。

#### func RemoveConfig

```go
func RemoveConfig(name ...string)
```

RemoveConfig removes the global configuration with specified group. If `name` is not passed, it removes configuration of the default group name.

​	RemoveConfig 删除具有指定组的全局配置。如果 `name` 未传递，则删除默认组名称的配置。

#### func SetConfig

```go
func SetConfig(config *Config, name ...string)
```

SetConfig sets the global configuration for specified group. If `name` is not passed, it sets configuration for the default group name.

​	SetConfig 设置指定组的全局配置。如果 `name` 未传递，它将设置默认组名称的配置。

#### func SetConfigByMap

```go
func SetConfigByMap(m map[string]interface{}, name ...string) error
```

SetConfigByMap sets the global configuration for specified group with map. If `name` is not passed, it sets configuration for the default group name.

​	SetConfigByMap 使用 map 设置指定组的全局配置。如果 `name` 未传递，它将设置默认组名称的配置。

## 类型

### type Adapter

```go
type Adapter interface {
	AdapterGroup
	AdapterOperation
}
```

Adapter is an interface for universal redis operations.

​	适配器是用于通用 redis 操作的接口。

### type AdapterFunc <-2.3.0

```go
type AdapterFunc func(config *Config) Adapter
```

AdapterFunc is the function creating redis adapter.

​	AdapterFunc 是创建 redis 适配器的函数。

### type AdapterGroup <-2.3.0

```go
type AdapterGroup interface {
	GroupGeneric() IGroupGeneric
	GroupHash() IGroupHash
	GroupList() IGroupList
	GroupPubSub() IGroupPubSub
	GroupScript() IGroupScript
	GroupSet() IGroupSet
	GroupSortedSet() IGroupSortedSet
	GroupString() IGroupString
}
```

AdapterGroup is an interface managing group operations for redis.

​	AdapterGroup 是管理 redis 的组操作的接口。

### type AdapterOperation <-2.6.2

```go
type AdapterOperation interface {
	// Do send a command to the server and returns the received reply.
	// It uses json.Marshal for struct/slice/map type values before committing them to redis.
	Do(ctx context.Context, command string, args ...interface{}) (*gvar.Var, error)

	// Conn retrieves and returns a connection object for continuous operations.
	// Note that you should call Close function manually if you do not use this connection any further.
	Conn(ctx context.Context) (conn Conn, err error)

	// Close closes current redis client, closes its connection pool and releases all its related resources.
	Close(ctx context.Context) (err error)
}
```

AdapterOperation is the core operation functions for redis. These functions can be easily overwritten by custom implements.

​	AdapterOperation 是 redis 的核心操作函数。这些函数可以很容易地被自定义实现覆盖。

### type Config

```go
type Config struct {
	// Address It supports single and cluster redis server. Multiple addresses joined with char ','. Eg: 192.168.1.1:6379, 192.168.1.2:6379.
	Address         string        `json:"address"`
	Db              int           `json:"db"`              // Redis db.
	User            string        `json:"user"`            // Username for AUTH.
	Pass            string        `json:"pass"`            // Password for AUTH.
	SentinelUser    string        `json:"sentinel_user"`   // Username for sentinel AUTH.
	SentinelPass    string        `json:"sentinel_pass"`   // Password for sentinel AUTH.
	MinIdle         int           `json:"minIdle"`         // Minimum number of connections allowed to be idle (default is 0)
	MaxIdle         int           `json:"maxIdle"`         // Maximum number of connections allowed to be idle (default is 10)
	MaxActive       int           `json:"maxActive"`       // Maximum number of connections limit (default is 0 means no limit).
	MaxConnLifetime time.Duration `json:"maxConnLifetime"` // Maximum lifetime of the connection (default is 30 seconds, not allowed to be set to 0)
	IdleTimeout     time.Duration `json:"idleTimeout"`     // Maximum idle time for connection (default is 10 seconds, not allowed to be set to 0)
	WaitTimeout     time.Duration `json:"waitTimeout"`     // Timed out duration waiting to get a connection from the connection pool.
	DialTimeout     time.Duration `json:"dialTimeout"`     // Dial connection timeout for TCP.
	ReadTimeout     time.Duration `json:"readTimeout"`     // Read timeout for TCP. DO NOT set it if not necessary.
	WriteTimeout    time.Duration `json:"writeTimeout"`    // Write timeout for TCP.
	MasterName      string        `json:"masterName"`      // Used in Redis Sentinel mode.
	TLS             bool          `json:"tls"`             // Specifies whether TLS should be used when connecting to the server.
	TLSSkipVerify   bool          `json:"tlsSkipVerify"`   // Disables server name verification when connecting over TLS.
	TLSConfig       *tls.Config   `json:"-"`               // TLS Config to use. When set TLS will be negotiated.
	SlaveOnly       bool          `json:"slaveOnly"`       // Route all commands to slave read-only nodes.
	Cluster         bool          `json:"cluster"`         // Specifies whether cluster mode be used.
	Protocol        int           `json:"protocol"`        // Specifies the RESP version (Protocol 2 or 3.)
}
```

Config is redis configuration.

​	Config 是 redis 配置。

#### func ConfigFromMap

```go
func ConfigFromMap(m map[string]interface{}) (config *Config, err error)
```

ConfigFromMap parses and returns config from given map.

​	ConfigFromMap 解析并返回给定映射中的配置。

#### func GetConfig

```go
func GetConfig(name ...string) (config *Config, ok bool)
```

GetConfig returns the global configuration with specified group name. If `name` is not passed, it returns configuration of the default group name.

​	GetConfig 返回具有指定组名称的全局配置。如果 `name` 未传递，则返回默认组名称的配置。

### type Conn

```go
type Conn interface {
	ConnCommand

	// Do send a command to the server and returns the received reply.
	// It uses json.Marshal for struct/slice/map type values before committing them to redis.
	Do(ctx context.Context, command string, args ...interface{}) (result *gvar.Var, err error)

	// Close puts the connection back to connection pool.
	Close(ctx context.Context) (err error)
}
```

Conn is an interface of a connection from universal redis client.

​	Conn 是来自通用 redis 客户端的连接接口。

### type ConnCommand <-2.3.0

```go
type ConnCommand interface {
	// Subscribe subscribes the client to the specified channels.
	// https://redis.io/commands/subscribe/
	Subscribe(ctx context.Context, channel string, channels ...string) ([]*Subscription, error)

	// PSubscribe subscribes the client to the given patterns.
	//
	// Supported glob-style patterns:
	// - h?llo subscribes to hello, hallo and hxllo
	// - h*llo subscribes to hllo and heeeello
	// - h[ae]llo subscribes to hello and hallo, but not hillo
	//
	// Use \ to escape special characters if you want to match them verbatim.
	//
	// https://redis.io/commands/psubscribe/
	PSubscribe(ctx context.Context, pattern string, patterns ...string) ([]*Subscription, error)

	// ReceiveMessage receives a single message of subscription from the Redis server.
	ReceiveMessage(ctx context.Context) (*Message, error)

	// Receive receives a single reply as gvar.Var from the Redis server.
	Receive(ctx context.Context) (result *gvar.Var, err error)
}
```

ConnCommand is an interface managing some operations bound to certain connection.

​	ConnCommand 是一个接口，用于管理绑定到某个连接的某些操作。

### type CopyOption <-2.3.0

```go
type CopyOption struct {
	DB      int  // DB option allows specifying an alternative logical database index for the destination key.
	REPLACE bool // REPLACE option removes the destination key before copying the value to it.
}
```

CopyOption provides options for function Copy.

​	CopyOption 为函数 Copy 提供选项。

### type ExpireOption <-2.3.0

```go
type ExpireOption struct {
	NX bool // NX -- Set expiry only when the key has no expiry
	XX bool // XX -- Set expiry only when the key has an existing expiry
	GT bool // GT -- Set expiry only when the new expiry is greater than current one
	LT bool // LT -- Set expiry only when the new expiry is less than current one
}
```

ExpireOption provides options for function Expire.

​	ExpireOption 为函数 Expire 提供选项。

### type FlushOp <-2.3.0

```go
type FlushOp string
const (
	FlushAsync FlushOp = "ASYNC" // ASYNC: flushes the databases asynchronously
	FlushSync  FlushOp = "SYNC"  // SYNC: flushes the databases synchronously
)
```

### type GetEXOption <-2.3.0

```go
type GetEXOption struct {
	TTLOption
	Persist bool // Persist -- Remove the time to live associated with the key.
}
```

GetEXOption provides extra option for GetEx function.

​	GetEXOption 为 GetEx 函数提供了额外的选项。

### type IGroupGeneric <-2.3.0

```go
type IGroupGeneric interface {
	Copy(ctx context.Context, source, destination string, option ...CopyOption) (int64, error)
	Exists(ctx context.Context, keys ...string) (int64, error)
	Type(ctx context.Context, key string) (string, error)
	Unlink(ctx context.Context, keys ...string) (int64, error)
	Rename(ctx context.Context, key, newKey string) error
	RenameNX(ctx context.Context, key, newKey string) (int64, error)
	Move(ctx context.Context, key string, db int) (int64, error)
	Del(ctx context.Context, keys ...string) (int64, error)
	RandomKey(ctx context.Context) (string, error)
	DBSize(ctx context.Context) (int64, error)
	Keys(ctx context.Context, pattern string) ([]string, error)
	FlushDB(ctx context.Context, option ...FlushOp) error
	FlushAll(ctx context.Context, option ...FlushOp) error
	Expire(ctx context.Context, key string, seconds int64, option ...ExpireOption) (int64, error)
	ExpireAt(ctx context.Context, key string, time time.Time, option ...ExpireOption) (int64, error)
	ExpireTime(ctx context.Context, key string) (*gvar.Var, error)
	TTL(ctx context.Context, key string) (int64, error)
	Persist(ctx context.Context, key string) (int64, error)
	PExpire(ctx context.Context, key string, milliseconds int64, option ...ExpireOption) (int64, error)
	PExpireAt(ctx context.Context, key string, time time.Time, option ...ExpireOption) (int64, error)
	PExpireTime(ctx context.Context, key string) (*gvar.Var, error)
	PTTL(ctx context.Context, key string) (int64, error)
}
```

IGroupGeneric manages generic redis operations. Implements see redis.GroupGeneric.

​	IGroupGeneric 管理泛型 redis 操作。实现见 redis。GroupGeneric。

### type IGroupHash <-2.3.0

```go
type IGroupHash interface {
	HSet(ctx context.Context, key string, fields map[string]interface{}) (int64, error)
	HSetNX(ctx context.Context, key, field string, value interface{}) (int64, error)
	HGet(ctx context.Context, key, field string) (*gvar.Var, error)
	HStrLen(ctx context.Context, key, field string) (int64, error)
	HExists(ctx context.Context, key, field string) (int64, error)
	HDel(ctx context.Context, key string, fields ...string) (int64, error)
	HLen(ctx context.Context, key string) (int64, error)
	HIncrBy(ctx context.Context, key, field string, increment int64) (int64, error)
	HIncrByFloat(ctx context.Context, key, field string, increment float64) (float64, error)
	HMSet(ctx context.Context, key string, fields map[string]interface{}) error
	HMGet(ctx context.Context, key string, fields ...string) (gvar.Vars, error)
	HKeys(ctx context.Context, key string) ([]string, error)
	HVals(ctx context.Context, key string) (gvar.Vars, error)
	HGetAll(ctx context.Context, key string) (*gvar.Var, error)
}
```

IGroupHash manages redis hash operations. Implements see redis.GroupHash.

​	IGroupHash 管理 redis 哈希操作。实现见 redis。GroupHash 中。

### type IGroupList <-2.3.0

```go
type IGroupList interface {
	LPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	LPushX(ctx context.Context, key string, element interface{}, elements ...interface{}) (int64, error)
	RPush(ctx context.Context, key string, values ...interface{}) (int64, error)
	RPushX(ctx context.Context, key string, value interface{}) (int64, error)
	LPop(ctx context.Context, key string, count ...int) (*gvar.Var, error)
	RPop(ctx context.Context, key string, count ...int) (*gvar.Var, error)
	LRem(ctx context.Context, key string, count int64, value interface{}) (int64, error)
	LLen(ctx context.Context, key string) (int64, error)
	LIndex(ctx context.Context, key string, index int64) (*gvar.Var, error)
	LInsert(ctx context.Context, key string, op LInsertOp, pivot, value interface{}) (int64, error)
	LSet(ctx context.Context, key string, index int64, value interface{}) (*gvar.Var, error)
	LRange(ctx context.Context, key string, start, stop int64) (gvar.Vars, error)
	LTrim(ctx context.Context, key string, start, stop int64) error
	BLPop(ctx context.Context, timeout int64, keys ...string) (gvar.Vars, error)
	BRPop(ctx context.Context, timeout int64, keys ...string) (gvar.Vars, error)
	RPopLPush(ctx context.Context, source, destination string) (*gvar.Var, error)
	BRPopLPush(ctx context.Context, source, destination string, timeout int64) (*gvar.Var, error)
}
```

IGroupList manages redis list operations. Implements see redis.GroupList.

​	IGroupList 管理 redis 列表操作。实现见 redis。GroupList。

### type IGroupPubSub <-2.3.0

```go
type IGroupPubSub interface {
	Publish(ctx context.Context, channel string, message interface{}) (int64, error)
	Subscribe(ctx context.Context, channel string, channels ...string) (Conn, []*Subscription, error)
	PSubscribe(ctx context.Context, pattern string, patterns ...string) (Conn, []*Subscription, error)
}
```

IGroupPubSub manages redis pub/sub operations. Implements see redis.GroupPubSub.

​	IGroupPubSub 管理 redis 发布/订阅操作。实现见 redis。GroupPubSub 中。

### type IGroupScript <-2.3.0

```go
type IGroupScript interface {
	Eval(ctx context.Context, script string, numKeys int64, keys []string, args []interface{}) (*gvar.Var, error)
	EvalSha(ctx context.Context, sha1 string, numKeys int64, keys []string, args []interface{}) (*gvar.Var, error)
	ScriptLoad(ctx context.Context, script string) (string, error)
	ScriptExists(ctx context.Context, sha1 string, sha1s ...string) (map[string]bool, error)
	ScriptFlush(ctx context.Context, option ...ScriptFlushOption) error
	ScriptKill(ctx context.Context) error
}
```

IGroupScript manages redis script operations. Implements see redis.GroupScript.

​	IGroupScript 管理 redis 脚本操作。实现见 redis。GroupScript的。

### type IGroupSet <-2.3.0

```go
type IGroupSet interface {
	SAdd(ctx context.Context, key string, member interface{}, members ...interface{}) (int64, error)
	SIsMember(ctx context.Context, key string, member interface{}) (int64, error)
	SPop(ctx context.Context, key string, count ...int) (*gvar.Var, error)
	SRandMember(ctx context.Context, key string, count ...int) (*gvar.Var, error)
	SRem(ctx context.Context, key string, member interface{}, members ...interface{}) (int64, error)
	SMove(ctx context.Context, source, destination string, member interface{}) (int64, error)
	SCard(ctx context.Context, key string) (int64, error)
	SMembers(ctx context.Context, key string) (gvar.Vars, error)
	SMIsMember(ctx context.Context, key, member interface{}, members ...interface{}) ([]int, error)
	SInter(ctx context.Context, key string, keys ...string) (gvar.Vars, error)
	SInterStore(ctx context.Context, destination string, key string, keys ...string) (int64, error)
	SUnion(ctx context.Context, key string, keys ...string) (gvar.Vars, error)
	SUnionStore(ctx context.Context, destination, key string, keys ...string) (int64, error)
	SDiff(ctx context.Context, key string, keys ...string) (gvar.Vars, error)
	SDiffStore(ctx context.Context, destination string, key string, keys ...string) (int64, error)
}
```

IGroupSet manages redis set operations. Implements see redis.GroupSet.

​	IGroupSet 管理 redis 集操作。实现见 redis。GroupSet。

### type IGroupSortedSet <-2.3.0

```go
type IGroupSortedSet interface {
	ZAdd(ctx context.Context, key string, option *ZAddOption, member ZAddMember, members ...ZAddMember) (*gvar.Var, error)
	ZScore(ctx context.Context, key string, member interface{}) (float64, error)
	ZIncrBy(ctx context.Context, key string, increment float64, member interface{}) (float64, error)
	ZCard(ctx context.Context, key string) (int64, error)
	ZCount(ctx context.Context, key string, min, max string) (int64, error)
	ZRange(ctx context.Context, key string, start, stop int64, option ...ZRangeOption) (gvar.Vars, error)
	ZRevRange(ctx context.Context, key string, start, stop int64, option ...ZRevRangeOption) (*gvar.Var, error)
	ZRank(ctx context.Context, key string, member interface{}) (int64, error)
	ZRevRank(ctx context.Context, key string, member interface{}) (int64, error)
	ZRem(ctx context.Context, key string, member interface{}, members ...interface{}) (int64, error)
	ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error)
	ZRemRangeByScore(ctx context.Context, key string, min, max string) (int64, error)
	ZRemRangeByLex(ctx context.Context, key string, min, max string) (int64, error)
	ZLexCount(ctx context.Context, key, min, max string) (int64, error)
}
```

IGroupSortedSet manages redis sorted set operations. Implements see redis.GroupSortedSet.

​	IGroupSortedSet 管理 redis 排序集操作。实现见 redis。GroupSortedSet。

### type IGroupString <-2.3.0

```go
type IGroupString interface {
	Set(ctx context.Context, key string, value interface{}, option ...SetOption) (*gvar.Var, error)
	SetNX(ctx context.Context, key string, value interface{}) (bool, error)
	SetEX(ctx context.Context, key string, value interface{}, ttlInSeconds int64) error
	Get(ctx context.Context, key string) (*gvar.Var, error)
	GetDel(ctx context.Context, key string) (*gvar.Var, error)
	GetEX(ctx context.Context, key string, option ...GetEXOption) (*gvar.Var, error)
	GetSet(ctx context.Context, key string, value interface{}) (*gvar.Var, error)
	StrLen(ctx context.Context, key string) (int64, error)
	Append(ctx context.Context, key string, value string) (int64, error)
	SetRange(ctx context.Context, key string, offset int64, value string) (int64, error)
	GetRange(ctx context.Context, key string, start, end int64) (string, error)
	Incr(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, increment int64) (int64, error)
	IncrByFloat(ctx context.Context, key string, increment float64) (float64, error)
	Decr(ctx context.Context, key string) (int64, error)
	DecrBy(ctx context.Context, key string, decrement int64) (int64, error)
	MSet(ctx context.Context, keyValueMap map[string]interface{}) error
	MSetNX(ctx context.Context, keyValueMap map[string]interface{}) (bool, error)
	MGet(ctx context.Context, keys ...string) (map[string]*gvar.Var, error)
}
```

IGroupString manages redis string operations. Implements see redis.GroupString.

​	IGroupString 管理 redis 字符串操作。实现见 redis。GroupString 中。

### type LInsertOp <-2.3.0

```go
type LInsertOp string
```

LInsertOp defines the operation name for function LInsert.

​	LInsertOp 定义函数 LInsert 的操作名称。

```go
const (
	LInsertBefore LInsertOp = "BEFORE"
	LInsertAfter  LInsertOp = "AFTER"
)
```

### type Message

```go
type Message struct {
	Channel      string
	Pattern      string
	Payload      string
	PayloadSlice []string
}
```

Message received as result of a PUBLISH command issued by another client.

​	由于另一个客户端发出的 PUBLISH 命令而收到的消息。

### type Redis

```go
type Redis struct {
	// contains filtered or unexported fields
}
```

Redis client.

​	Redis 客户端。

#### func Instance

```go
func Instance(name ...string) *Redis
```

Instance returns an instance of redis client with specified group. The `name` param is unnecessary, if `name` is not passed, it returns a redis instance with default configuration group.

​	实例返回具有指定组的 redis 客户端实例。 `name` 参数是不必要的，如果 `name` 未传递，则返回具有默认配置组的 redis 实例。

#### func New

```go
func New(config ...*Config) (*Redis, error)
```

New creates and returns a redis client. It creates a default redis adapter of go-redis.

​	New 创建并返回 redis 客户端。它创建一个默认的 redis 适配器 go-redis。

#### func NewWithAdapter

```go
func NewWithAdapter(adapter Adapter) (*Redis, error)
```

NewWithAdapter creates and returns a redis client with given adapter.

​	NewWithAdapter 创建并返回具有给定适配器的 redis 客户端。

#### (*Redis) Close

```go
func (r *Redis) Close(ctx context.Context) error
```

Close closes current redis client, closes its connection pool and releases all its related resources.

​	关闭将关闭当前 redis 客户端，关闭其连接池并释放其所有相关资源。

#### (*Redis) Conn

```go
func (r *Redis) Conn(ctx context.Context) (Conn, error)
```

Conn retrieves and returns a connection object for continuous operations. Note that you should call Close function manually if you do not use this connection any further.

​	Conn 检索并返回连接对象以进行连续操作。请注意，如果不再使用此连接，则应手动调用 Close 函数。

#### (*Redis) Do

```go
func (r *Redis) Do(ctx context.Context, command string, args ...interface{}) (*gvar.Var, error)
```

Do send a command to the server and returns the received reply. It uses json.Marshal for struct/slice/map type values before committing them to redis.

​	务必向服务器发送命令并返回收到的回复。它使用 json。在将 struct/slice/map 类型值提交到 redis 之前对其进行封送。

#### (*Redis) GetAdapter

```go
func (r *Redis) GetAdapter() Adapter
```

GetAdapter returns the adapter that is set in current redis client.

​	GetAdapter 返回在当前 redis 客户端中设置的适配器。

#### (*Redis) MustConn

```go
func (r *Redis) MustConn(ctx context.Context) Conn
```

MustConn performs as function Conn, but it panics if any error occurs internally.

​	MustConn 作为函数 Conn 执行，但如果内部发生任何错误，它会崩溃。

#### (*Redis) MustDo

```go
func (r *Redis) MustDo(ctx context.Context, command string, args ...interface{}) *gvar.Var
```

MustDo performs as function Do, but it panics if any error occurs internally.

​	MustDo 与函数 Do 一样执行，但如果内部发生任何错误，它会崩溃。

#### (*Redis) SetAdapter

```go
func (r *Redis) SetAdapter(adapter Adapter)
```

SetAdapter changes the underlying adapter with custom adapter for current redis client.

​	SetAdapter 使用当前 redis 客户端的自定义适配器更改基础适配器。

### type ScriptFlushOption <-2.3.0

```go
type ScriptFlushOption struct {
	SYNC  bool // SYNC  flushes the cache synchronously.
	ASYNC bool // ASYNC flushes the cache asynchronously.
}
```

ScriptFlushOption provides options for function ScriptFlush.

​	ScriptFlushOption 提供函数 ScriptFlush 的选项。

### type SetOption <-2.3.0

```go
type SetOption struct {
	TTLOption
	NX bool // Only set the key if it does not already exist.
	XX bool // Only set the key if it already exists.

	// Return the old string stored at key, or nil if key did not exist.
	// An error is returned and SET aborted if the value stored at key is not a string.
	Get bool
}
```

SetOption provides extra option for Set function.

​	SetOption 为 Set 函数提供了额外的选项。

### type Subscription

```go
type Subscription struct {
	Kind    string // Can be "subscribe", "unsubscribe", "psubscribe" or "punsubscribe".
	Channel string // Channel name we have subscribed to.
	Count   int    // Number of channels we are currently subscribed to.
}
```

Subscription received after a successful subscription to channel.

​	成功订阅频道后收到的订阅。

#### (*Subscription) String

```go
func (m *Subscription) String() string
```

String converts current object to a readable string.

​	String 将当前对象转换为可读字符串。

### type TTLOption <-2.3.0

```go
type TTLOption struct {
	EX      *int64 // EX seconds -- Set the specified expire time, in seconds.
	PX      *int64 // PX milliseconds -- Set the specified expire time, in milliseconds.
	EXAT    *int64 // EXAT timestamp-seconds -- Set the specified Unix time at which the key will expire, in seconds.
	PXAT    *int64 // PXAT timestamp-milliseconds -- Set the specified Unix time at which the key will expire, in milliseconds.
	KeepTTL bool   // Retain the time to live associated with the key.
}
```

TTLOption provides extra option for TTL related functions.

​	TTLOption 为 TTL 相关功能提供了额外的选项。

### type ZAddMember <-2.3.0

```go
type ZAddMember struct {
	Score  float64
	Member interface{}
}
```

ZAddMember is element struct for set.

​	ZAddMember 是 set 的元素结构。

### type ZAddOption <-2.3.0

```go
type ZAddOption struct {
	XX bool // Only update elements that already exist. Don't add new elements.
	NX bool // Only add new elements. Don't update already existing elements.
	// Only update existing elements if the new score is less than the current score.
	// This flag doesn't prevent adding new elements.
	LT bool

	// Only update existing elements if the new score is greater than the current score.
	// This flag doesn't prevent adding new elements.
	GT bool

	// Modify the return value from the number of new elements added, to the total number of elements changed (CH is an abbreviation of changed).
	// Changed elements are new elements added and elements already existing for which the score was updated.
	// So elements specified in the command line having the same score as they had in the past are not counted.
	// Note: normally the return value of ZAdd only counts the number of new elements added.
	CH bool

	// When this option is specified ZAdd acts like ZIncrBy. Only one score-element pair can be specified in this mode.
	INCR bool
}
```

ZAddOption provides options for function ZAdd.

​	ZAddOption 为函数 ZAdd 提供选项。

### type ZRangeOption <-2.3.0

```go
type ZRangeOption struct {
	ByScore bool
	ByLex   bool
	// The optional REV argument reverses the ordering, so elements are ordered from highest to lowest score,
	// and score ties are resolved by reverse lexicographical ordering.
	Rev   bool
	Limit *ZRangeOptionLimit
	// The optional WithScores argument supplements the command's reply with the scores of elements returned.
	WithScores bool
}
```

ZRangeOption provides extra option for ZRange function.

​	ZRangeOption 为 ZRange 函数提供了额外的选项。

### type ZRangeOptionLimit <-2.3.0

```go
type ZRangeOptionLimit struct {
	Offset *int
	Count  *int
}
```

ZRangeOptionLimit provides LIMIT argument for ZRange function. The optional LIMIT argument can be used to obtain a sub-range from the matching elements (similar to SELECT LIMIT offset, count in SQL). A negative `Count` returns all elements from the `Offset`.

​	ZRangeOptionLimit 为 ZRange 函数提供 LIMIT 参数。可选的 LIMIT 参数可用于从匹配元素中获取子范围（类似于 SQL 中的 SELECT LIMIT offset、count）。否定 `Count` 返回 `Offset` 中的所有元素。

### type ZRevRangeOption <-2.3.0

```go
type ZRevRangeOption struct {
	WithScores bool
}
```

ZRevRangeOption provides options for function ZRevRange.

​	ZRevRangeOption 提供函数 ZRevRange 的选项。