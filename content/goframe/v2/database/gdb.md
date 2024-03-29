+++
title = "gdb"
date = 2024-03-21T17:47:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/database/gdb](https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/database/gdb)

Package gdb provides ORM features for popular relationship databases.

​	软件包 gdb 为流行的关系数据库提供 ORM 功能。

TODO use context.Context as required parameter for all DB operations.

​	TODO 使用上下文。上下文作为所有数据库操作的必需参数。

## 常量

[View Source](https://github.com/gogf/gf/blob/v2.6.4/database/gdb/gdb.go#L413)

```go
const (
	InsertOperationInsert      = "INSERT"
	InsertOperationReplace     = "REPLACE"
	InsertOperationIgnore      = "INSERT IGNORE"
	InsertOnDuplicateKeyUpdate = "ON DUPLICATE KEY UPDATE"
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/database/gdb/gdb.go#L420)

```go
const (
	SqlTypeBegin               = "DB.Begin"
	SqlTypeTXCommit            = "TX.Commit"
	SqlTypeTXRollback          = "TX.Rollback"
	SqlTypeExecContext         = "DB.ExecContext"
	SqlTypeQueryContext        = "DB.QueryContext"
	SqlTypePrepareContext      = "DB.PrepareContext"
	SqlTypeStmtExecContext     = "DB.Statement.ExecContext"
	SqlTypeStmtQueryContext    = "DB.Statement.QueryContext"
	SqlTypeStmtQueryRowContext = "DB.Statement.QueryRowContext"
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/database/gdb/gdb_func.go#L60)

```go
const (
	OrmTagForStruct    = "orm"
	OrmTagForTable     = "table"
	OrmTagForWith      = "with"
	OrmTagForWithWhere = "where"
	OrmTagForWithOrder = "order"
	OrmTagForDo        = "do"
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/database/gdb/gdb_core_config.go#L58)

```go
const (
	DefaultGroupName = "default" // Default group name.
)
```

## 变量

This section is empty.

## 函数

#### func AddConfigNode

```go
func AddConfigNode(group string, node ConfigNode)
```

AddConfigNode adds one node configuration to configuration of given group.

​	AddConfigNode 将一个节点配置添加到给定组的配置中。

#### func AddDefaultConfigGroup

```go
func AddDefaultConfigGroup(nodes ConfigGroup)
```

AddDefaultConfigGroup adds multiple node configurations to configuration of default group.

​	AddDefaultConfigGroup 将多个节点配置添加到默认组的配置中。

#### func AddDefaultConfigNode

```go
func AddDefaultConfigNode(node ConfigNode)
```

AddDefaultConfigNode adds one node configuration to configuration of default group.

​	AddDefaultConfigNode 将一个节点配置添加到默认组的配置中。

#### func CatchSQL <-2.2.0

```go
func CatchSQL(ctx context.Context, f func(ctx context.Context) error) (sqlArray []string, err error)
```

CatchSQL catches and returns all sql statements that are EXECUTED in given closure function. Be caution that, all the following sql statements should use the context object passing by function `f`.

​	CatchSQL 捕获并返回在给定闭包函数中执行的所有 sql 语句。请注意，以下所有 sql 语句都应使用 context object passing by function `f` 。

#### func FormatMultiLineSqlToSingle <-2.6.4

```go
func FormatMultiLineSqlToSingle(sqlTmp string) string
```

FormatMultiLineSqlToSingle formats sql template string into one line.

​	FormatMultiLineSqlToSingle 将 sql 模板字符串格式化为一行。

#### func FormatSqlWithArgs

```go
func FormatSqlWithArgs(sql string, args []interface{}) string
```

FormatSqlWithArgs binds the arguments to the sql string and returns a complete sql string, just for debugging.

​	FormatSqlWithArgs 将参数绑定到 sql 字符串，并返回一个完整的 sql 字符串，仅用于调试。

#### func GetDefaultGroup

```go
func GetDefaultGroup() string
```

GetDefaultGroup returns the { name of default configuration.

​	GetDefaultGroup 返回默认配置的 { 名称。

#### func GetInsertOperationByOption

```go
func GetInsertOperationByOption(option InsertOption) string
```

GetInsertOperationByOption returns proper insert option with given parameter `option`.

​	GetInsertOperationByOption 返回具有给定参数的正确插入选项 `option` 。

#### func GetPrimaryKeyCondition

```go
func GetPrimaryKeyCondition(primary string, where ...interface{}) (newWhereCondition []interface{})
```

GetPrimaryKeyCondition returns a new where condition by primary field name. The optional parameter `where` is like follows: 123 => primary=123 []int{1, 2, 3} => primary IN(1,2,3) “john” => primary=‘john’ []string{“john”, “smith”} => primary IN(‘john’,‘smith’) g.Map{“id”: g.Slice{1,2,3}} => id IN(1,2,3) g.Map{“id”: 1, “name”: “john”} => id=1 AND name=‘john’ etc.

​	GetPrimaryKeyCondition 按主字段名称返回新的 where 条件。可选参数 `where` 如下： 123 => primary=123 []int{1， 2， 3} => primary IN（1,2,3） “john” => primary='john' []string{“john”， “smith”} => primary IN（'john'，'smith'） g.Map{“id”： g.Slice{1,2,3}} => id IN（1,2,3） g.Map{“id”： 1， “name”： “john”} => id=1 AND name='john' 等。

Note that it returns the given `where` parameter directly if the `primary` is empty or length of `where` > 1.

​	请注意，如果 为 `primary` 空或长度为 `where` > 1，则直接返回给定 `where` 参数。

#### func IsConfigured

```go
func IsConfigured() bool
```

IsConfigured checks and returns whether the database configured. It returns true if any configuration exists.

​	IsConfigured 检查并返回是否配置了数据库。如果存在任何配置，则返回 true。

#### func ListItemValues

```go
func ListItemValues(list interface{}, key interface{}, subKey ...interface{}) (values []interface{})
```

ListItemValues retrieves and returns the elements of all item struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice.

​	ListItemValues 检索并返回所有项目 struct/map 的元素，并带有键 `key` 。请注意，该参数 `list` 应为包含 map 或 struct 元素的切片类型，否则它将返回一个空切片。

The parameter `list` supports types like: []map[string]interface{} []map[string]sub-map []struct []struct:sub-struct Note that the sub-map/sub-struct makes sense only if the optional parameter `subKey` is given. See gutil.ListItemValues.

​	该参数 `list` 支持以下类型： []map[string]interface{} []map[string]sub-map []struct []struct：sub-struct 请注意，仅当给定可选参数 `subKey` 时，sub-map/sub-struct 才有意义。请参见 gutil。ListItemValues。

#### func ListItemValuesUnique

```go
func ListItemValuesUnique(list interface{}, key string, subKey ...interface{}) []interface{}
```

ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice. See gutil.ListItemValuesUnique.

​	ListItemValuesUnique 检索并返回所有带有键 `key` 的结构/映射的唯一元素。请注意，该参数 `list` 应为包含 map 或 struct 元素的切片类型，否则它将返回一个空切片。请参见 gutil。ListItemValuesUnique。

#### func MapOrStructToMapDeep <-2.6.0

```go
func MapOrStructToMapDeep(value interface{}, omitempty bool) map[string]interface{}
```

MapOrStructToMapDeep converts `value` to map type recursively(if attribute struct is embedded). The parameter `value` should be type of *map/map/*struct/struct. It supports embedded struct definition for struct.

​	MapOrStructToMapDeep 递归 `value` 转换为映射类型（如果嵌入了属性结构）。参数 `value` 的类型应为 *map/map/*struct/struct。它支持 struct 的嵌入式结构定义。

#### func Register

```go
func Register(name string, driver Driver) error
```

Register registers custom database driver to gdb.

​	Register 将自定义数据库驱动程序注册到 gdb。

#### func SetConfig

```go
func SetConfig(config Config)
```

SetConfig sets the global configuration for package. It will overwrite the old configuration of package.

​	SetConfig 设置包的全局配置。它将覆盖包的旧配置。

#### func SetConfigGroup

```go
func SetConfigGroup(group string, nodes ConfigGroup)
```

SetConfigGroup sets the configuration for given group.

​	SetConfigGroup 设置给定组的配置。

#### func SetDefaultGroup

```go
func SetDefaultGroup(name string)
```

SetDefaultGroup sets the group name for default configuration.

​	SetDefaultGroup 设置默认配置的组名称。

#### func ToSQL <-2.2.0

```go
func ToSQL(ctx context.Context, f func(ctx context.Context) error) (sql string, err error)
```

ToSQL formats and returns the last one of sql statements in given closure function WITHOUT TRULY EXECUTING IT. Be caution that, all the following sql statements should use the context object passing by function `f`.

​	ToSQL 格式化并返回给定闭包函数中的最后一个 sql 语句，而没有真正执行它。请注意，以下所有 sql 语句都应使用 context object passing by function `f` 。

#### func WithDB <-2.0.5

```go
func WithDB(ctx context.Context, db DB) context.Context
```

WithDB injects given db object into context and returns a new context.

​	WithDB 将给定的 db 对象注入到上下文中并返回一个新上下文。

#### func WithTX

```go
func WithTX(ctx context.Context, tx TX) context.Context
```

WithTX injects given transaction object into context and returns a new context.

​	WithTX 将给定的事务对象注入到上下文中并返回新的上下文。

## 类型

### type CacheOption

```go
type CacheOption struct {
	// Duration is the TTL for the cache.
	// If the parameter `Duration` < 0, which means it clear the cache with given `Name`.
	// If the parameter `Duration` = 0, which means it never expires.
	// If the parameter `Duration` > 0, which means it expires after `Duration`.
	Duration time.Duration

	// Name is an optional unique name for the cache.
	// The Name is used to bind a name to the cache, which means you can later control the cache
	// like changing the `duration` or clearing the cache with specified Name.
	Name string

	// Force caches the query result whatever the result is nil or not.
	// It is used to avoid Cache Penetration.
	Force bool
}
```

CacheOption is options for model cache control in query.

​	CacheOption 是用于查询中模型缓存控制的选项。

### type CatchSQLManager <-2.2.0

```go
type CatchSQLManager struct {
	SQLArray *garray.StrArray
	DoCommit bool // DoCommit marks it will be committed to underlying driver or not.
}
```

### type ChunkHandler

```go
type ChunkHandler func(result Result, err error) bool
```

ChunkHandler is a function that is used in function Chunk, which handles given Result and error. It returns true if it wants to continue chunking, or else it returns false to stop chunking.

​	ChunkHandler 是函数 Chunk 中使用的函数，用于处理给定的 Result 和 error。如果它想继续分块，它返回 true，否则它返回 false 停止分块。

### type Config

```go
type Config map[string]ConfigGroup
```

Config is the configuration management object.

​	Config 是配置管理对象。

### type ConfigGroup

```go
type ConfigGroup []ConfigNode
```

ConfigGroup is a slice of configuration node for specified named group.

​	ConfigGroup 是指定命名组的配置节点切片。

#### func GetConfig

```go
func GetConfig(group string) ConfigGroup
```

GetConfig retrieves and returns the configuration of given group.

​	GetConfig 检索并返回给定组的配置。

### type ConfigNode

```go
type ConfigNode struct {
	Host                 string        `json:"host"`                 // Host of server, ip or domain like: 127.0.0.1, localhost
	Port                 string        `json:"port"`                 // Port, it's commonly 3306.
	User                 string        `json:"user"`                 // Authentication username.
	Pass                 string        `json:"pass"`                 // Authentication password.
	Name                 string        `json:"name"`                 // Default used database name.
	Type                 string        `json:"type"`                 // Database type: mysql, mariadb, sqlite, mssql, pgsql, oracle, clickhouse, dm.
	Link                 string        `json:"link"`                 // (Optional) Custom link information for all configuration in one single string.
	Extra                string        `json:"extra"`                // (Optional) Extra configuration according the registered third-party database driver.
	Role                 string        `json:"role"`                 // (Optional, "master" in default) Node role, used for master-slave mode: master, slave.
	Debug                bool          `json:"debug"`                // (Optional) Debug mode enables debug information logging and output.
	Prefix               string        `json:"prefix"`               // (Optional) Table prefix.
	DryRun               bool          `json:"dryRun"`               // (Optional) Dry run, which does SELECT but no INSERT/UPDATE/DELETE statements.
	Weight               int           `json:"weight"`               // (Optional) Weight for load balance calculating, it's useless if there's just one node.
	Charset              string        `json:"charset"`              // (Optional, "utf8" in default) Custom charset when operating on database.
	Protocol             string        `json:"protocol"`             // (Optional, "tcp" in default) See net.Dial for more information which networks are available.
	Timezone             string        `json:"timezone"`             // (Optional) Sets the time zone for displaying and interpreting time stamps.
	Namespace            string        `json:"namespace"`            // (Optional) Namespace for some databases. Eg, in pgsql, the `Name` acts as the `catalog`, the `NameSpace` acts as the `schema`.
	MaxIdleConnCount     int           `json:"maxIdle"`              // (Optional) Max idle connection configuration for underlying connection pool.
	MaxOpenConnCount     int           `json:"maxOpen"`              // (Optional) Max open connection configuration for underlying connection pool.
	MaxConnLifeTime      time.Duration `json:"maxLifeTime"`          // (Optional) Max amount of time a connection may be idle before being closed.
	QueryTimeout         time.Duration `json:"queryTimeout"`         // (Optional) Max query time for per dql.
	ExecTimeout          time.Duration `json:"execTimeout"`          // (Optional) Max exec time for dml.
	TranTimeout          time.Duration `json:"tranTimeout"`          // (Optional) Max exec time for a transaction.
	PrepareTimeout       time.Duration `json:"prepareTimeout"`       // (Optional) Max exec time for prepare operation.
	CreatedAt            string        `json:"createdAt"`            // (Optional) The field name of table for automatic-filled created datetime.
	UpdatedAt            string        `json:"updatedAt"`            // (Optional) The field name of table for automatic-filled updated datetime.
	DeletedAt            string        `json:"deletedAt"`            // (Optional) The field name of table for automatic-filled updated datetime.
	TimeMaintainDisabled bool          `json:"timeMaintainDisabled"` // (Optional) Disable the automatic time maintaining feature.
}
```

ConfigNode is configuration for one node.

​	ConfigNode 是一个节点的配置。

### type Core

```go
type Core struct {
	// contains filtered or unexported fields
}
```

Core is the base struct for database management.

​	Core 是数据库管理的基本结构。

#### (*Core) Begin

```go
func (c *Core) Begin(ctx context.Context) (tx TX, err error)
```

Begin starts and returns the transaction object. You should call Commit or Rollback functions of the transaction object if you no longer use the transaction. Commit or Rollback functions will also close the transaction automatically.

​	Begin 启动并返回事务对象。如果不再使用事务，则应调用事务对象的 Commit 或 Rollback 函数。Commit 或 Rollback 函数也将自动关闭事务。

#### (*Core) CheckLocalTypeForField

```go
func (c *Core) CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (LocalType, error)
```

CheckLocalTypeForField checks and returns corresponding type for given db type.

​	CheckLocalTypeForField 检查并返回给定数据库类型的相应类型。

#### (*Core) ClearCache

```go
func (c *Core) ClearCache(ctx context.Context, table string) (err error)
```

ClearCache removes cached sql result of certain table.

​	ClearCache 删除某个表的缓存 sql 结果。

#### (*Core) ClearCacheAll

```go
func (c *Core) ClearCacheAll(ctx context.Context) (err error)
```

ClearCacheAll removes all cached sql result from cache

​	ClearCacheAll 从缓存中删除所有缓存的 sql 结果

#### (*Core) ClearTableFields

```go
func (c *Core) ClearTableFields(ctx context.Context, table string, schema ...string) (err error)
```

ClearTableFields removes certain cached table fields of current configuration group.

​	ClearTableFields 删除当前配置组的某些缓存表字段。

#### (*Core) ClearTableFieldsAll

```go
func (c *Core) ClearTableFieldsAll(ctx context.Context) (err error)
```

ClearTableFieldsAll removes all cached table fields of current configuration group.

​	ClearTableFieldsAll 删除当前配置组的所有缓存表字段。

#### (*Core) Close

```go
func (c *Core) Close(ctx context.Context) (err error)
```

Close closes the database and prevents new queries from starting. Close then waits for all queries that have started processing on the server to finish.

​	关闭将关闭数据库并阻止启动新查询。然后，关闭将等待服务器上已开始处理的所有查询完成。

It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.

​	关闭数据库的情况很少见，因为数据库句柄是长期存在的，并且在许多 goroutine 之间共享。

#### (*Core) ConvertDataForRecord

```go
func (c *Core) ConvertDataForRecord(ctx context.Context, value interface{}, table string) (map[string]interface{}, error)
```

ConvertDataForRecord is a very important function, which does converting for any data that will be inserted into table/collection as a record.

​	ConvertDataForRecord 是一个非常重要的函数，它对将作为记录插入到表/集合中的任何数据进行转换。

The parameter `value` should be type of *map/map/*struct/struct. It supports embedded struct definition for struct.

​	参数 `value` 的类型应为 *map/map/*struct/struct。它支持 struct 的嵌入式结构定义。

#### (*Core) ConvertValueForField

```go
func (c *Core) ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error)
```

ConvertValueForField converts value to the type of the record field. The parameter `fieldType` is the target record field. The parameter `fieldValue` is the value that to be committed to record field.

​	ConvertValueForField 将值转换为记录字段的类型。该参数 `fieldType` 是目标记录字段。该参数 `fieldValue` 是要提交到记录字段的值。

#### (*Core) ConvertValueForLocal

```go
func (c *Core) ConvertValueForLocal(
	ctx context.Context, fieldType string, fieldValue interface{},
) (interface{}, error)
```

ConvertValueForLocal converts value to local Golang type of value according field type name from database. The parameter `fieldType` is in lower case, like: `float(5,2)`, `unsigned double(5,2)`, `decimal(10,2)`, `char(45)`, `varchar(100)`, etc.

​	ConvertValueForLocal 根据数据库中的字段类型名称将值转换为值的本地 Golang 类型。参数 `fieldType` 为小写，如： `float(5,2)` 、 `unsigned double(5,2)` `decimal(10,2)` `char(45)` `varchar(100)` 等。

#### (*Core) Ctx

```go
func (c *Core) Ctx(ctx context.Context) DB
```

Ctx is a chaining function, which creates and returns a new DB that is a shallow copy of current DB object and with given context in it. Note that this returned DB object can be used only once, so do not assign it to a global or package variable for long using.

​	Ctx 是一个链接函数，它创建并返回一个新的数据库，该数据库是当前数据库对象的浅层副本，并且其中包含给定的上下文。请注意，此返回的 DB 对象只能使用一次，因此不要将其分配给全局变量或包变量以长期使用。

#### (*Core) Delete

```go
func (c *Core) Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (result sql.Result, err error)
```

Delete does “DELETE FROM … " statement for the table.

​	Delete 执行“DELETE FROM ...“的语句。

The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc. It is commonly used with parameter `args`. Eg: “uid=10000”, “uid”, 10000 “money>? AND name like ?”, 99999, “vip_%” “status IN (?)”, g.Slice{1,2,3} “age IN(?,?)”, 18, 50 User{ Id : 1, UserName : “john”}.

​	参数 `condition` 可以是字符串/地图/gmap/slice/struct/*struct等类型。它通常与参数一起使用 `args` 。例如：“uid=10000”、“uid”、10000“钱>？AND name like ？“， 99999， ”vip_%“ ”status IN （？）“， g.Slice{1,2,3} ”age IN（?,?）“， 18， 50 User{ Id ： 1， UserName ： ”john“}.

#### (*Core) DoCommit

```go
func (c *Core) DoCommit(ctx context.Context, in DoCommitInput) (out DoCommitOutput, err error)
```

DoCommit commits current sql and arguments to underlying sql driver.

​	DoCommit 将当前 sql 和参数提交到底层 sql 驱动程序。

#### (*Core) DoDelete

```go
func (c *Core) DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result sql.Result, err error)
```

DoDelete does “DELETE FROM … " statement for the table. This function is usually used for custom interface definition, you do not need call it manually.

​	DoDelete 执行“从...“的语句。此函数通常用于自定义接口定义，无需手动调用。

#### (*Core) DoExec

```go
func (c *Core) DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result sql.Result, err error)
```

DoExec commits the sql string and its arguments to underlying driver through given link object and returns the execution result.

​	DoExec 通过给定的链接对象将 sql 字符串及其参数提交到底层驱动程序，并返回执行结果。

#### (*Core) DoFilter

```go
func (c *Core) DoFilter(ctx context.Context, link Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error)
```

DoFilter is a hook function, which filters the sql and its arguments before it’s committed to underlying driver. The parameter `link` specifies the current database connection operation object. You can modify the sql string `sql` and its arguments `args` as you wish before they’re committed to driver.

​	DoFilter 是一个钩子函数，用于在将 sql 及其参数提交到基础驱动程序之前对其进行筛选。该参数 `link` 指定当前数据库连接操作对象。在将 sql 字符串 `sql` 及其参数 `args` 提交到驱动程序之前，可以根据需要修改它们。

#### (*Core) DoInsert

```go
func (c *Core) DoInsert(ctx context.Context, link Link, table string, list List, option DoInsertOption) (result sql.Result, err error)
```

DoInsert inserts or updates data forF given table. This function is usually used for custom interface definition, you do not need call it manually. The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”})

​	DoInsert 插入或更新 F 给定表的数据。此函数通常用于自定义接口定义，无需手动调用。参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。 例如： Data（g.Map{“uid”： 10000， “name”：“john”}） 数据（g.Slice{g.Map{“uid”： 10000， “name”：“john”}， g.Map{“uid”： 20000， “name”：“smith”}）

The parameter `option` values are as follows: InsertOptionDefault: just insert, if there’s unique/primary key in the data, it returns error; InsertOptionReplace: if there’s unique/primary key in the data, it deletes it from table and inserts a new one; InsertOptionSave: if there’s unique/primary key in the data, it updates it or else inserts a new one; InsertOptionIgnore: if there’s unique/primary key in the data, it ignores the inserting;

​	参数 `option` 值如下： InsertOptionDefault：直接插入，如果数据中有唯一/主键，则返回错误;InsertOptionReplace：如果数据中有唯一键/主键，则将其从表中删除并插入新键;InsertOptionSave：如果数据中存在唯一/主键，则更新它或插入新键;InsertOptionIgnore：如果数据中有唯一/主键，则忽略插入;

#### (*Core) DoPrepare

```go
func (c *Core) DoPrepare(ctx context.Context, link Link, sql string) (stmt *Stmt, err error)
```

DoPrepare calls prepare function on given link object and returns the statement object.

​	DoPrepare 调用给定链接对象的 prepare 函数并返回语句对象。

#### (*Core) DoQuery

```go
func (c *Core) DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)
```

DoQuery commits the sql string and its arguments to underlying driver through given link object and returns the execution result.

​	DoQuery 通过给定的链接对象将 sql 字符串及其参数提交到底层驱动程序，并返回执行结果。

#### (*Core) DoSelect

```go
func (c *Core) DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)
```

DoSelect queries and returns data records from database.

​	DoSelect 从数据库中查询并返回数据记录。

#### (*Core) DoUpdate

```go
func (c *Core) DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error)
```

DoUpdate does “UPDATE … " statement for the table. This function is usually used for custom interface definition, you do not need to call it manually.

​	DoUpdate 执行“更新...“的语句。此函数通常用于自定义接口定义，无需手动调用。

#### (*Core) Exec

```go
func (c *Core) Exec(ctx context.Context, sql string, args ...interface{}) (result sql.Result, err error)
```

Exec commits one query SQL to underlying driver and returns the execution result. It is most commonly used for data inserting and updating.

​	Exec 将一个查询 SQL 提交到基础驱动程序并返回执行结果。它最常用于数据插入和更新。

#### (*Core) FilteredLink

```go
func (c *Core) FilteredLink() string
```

FilteredLink retrieves and returns filtered `linkInfo` that can be using for logging or tracing purpose.

​	FilteredLink 检索并返回可用于日志记录或跟踪目的的 `linkInfo` 筛选。

#### (*Core) FormatSqlBeforeExecuting

```go
func (c *Core) FormatSqlBeforeExecuting(sql string, args []interface{}) (newSql string, newArgs []interface{})
```

FormatSqlBeforeExecuting formats the sql string and its arguments before executing. The internal handleArguments function might be called twice during the SQL procedure, but do not worry about it, it’s safe and efficient.

​	FormatSqlBeforeExecuting 在执行之前设置 sql 字符串及其参数的格式。内部 handleArguments 函数在 SQL 过程中可能会被调用两次，但不用担心，它是安全高效的。

#### (*Core) FormatUpsert

```go
func (c *Core) FormatUpsert(columns []string, list List, option DoInsertOption) (string, error)
```

FormatUpsert formats and returns SQL clause part for upsert statement. In default implements, this function performs upsert statement for MySQL like: `INSERT INTO ... ON DUPLICATE KEY UPDATE x=VALUES(z),m=VALUES(y)...`

​	FormatUpsert 格式化并返回 upsert 语句的 SQL 子句部分。在默认实现中，此函数对 MySQL 执行 upsert 语句，如下所示： `INSERT INTO ... ON DUPLICATE KEY UPDATE x=VALUES(z),m=VALUES(y)...`

#### (*Core) GetAll

```go
func (c *Core) GetAll(ctx context.Context, sql string, args ...interface{}) (Result, error)
```

GetAll queries and returns data records from database.

​	GetAll 查询并返回数据库中的数据记录。

#### (*Core) GetArray

```go
func (c *Core) GetArray(ctx context.Context, sql string, args ...interface{}) ([]Value, error)
```

GetArray queries and returns data values as slice from database. Note that if there are multiple columns in the result, it returns just one column values randomly.

​	GetArray 查询数据值并将其作为数据库中的切片返回。请注意，如果结果中有多个列，则仅随机返回一个列值。

#### (*Core) GetCache

```go
func (c *Core) GetCache() *gcache.Cache
```

GetCache returns the internal cache object.

​	GetCache 返回内部缓存对象。

#### (*Core) GetChars

```go
func (c *Core) GetChars() (charLeft string, charRight string)
```

GetChars returns the security char for current database. It does nothing in default.

​	GetChars 返回当前数据库的安全字符。默认情况下，它不执行任何操作。

#### (*Core) GetConfig

```go
func (c *Core) GetConfig() *ConfigNode
```

GetConfig returns the current used node configuration.

​	GetConfig 返回当前使用的节点配置。

#### (*Core) GetCore

```go
func (c *Core) GetCore() *Core
```

GetCore returns the underlying *Core object.

​	GetCore 返回基础 *Core 对象。

#### (*Core) GetCount

```go
func (c *Core) GetCount(ctx context.Context, sql string, args ...interface{}) (int, error)
```

GetCount queries and returns the count from database.

​	GetCount 查询并返回数据库中的计数。

#### (*Core) GetCtx

```go
func (c *Core) GetCtx() context.Context
```

GetCtx returns the context for current DB. It returns `context.Background()` is there’s no context previously set.

​	GetCtx 返回当前数据库的上下文。它返回 `context.Background()` 的是之前没有设置的上下文。

#### (*Core) GetCtxTimeout

```go
func (c *Core) GetCtxTimeout(ctx context.Context, timeoutType int) (context.Context, context.CancelFunc)
```

GetCtxTimeout returns the context and cancel function for specified timeout type.

​	GetCtxTimeout 返回指定超时类型的上下文和 cancel 函数。

#### (*Core) GetDB

```go
func (c *Core) GetDB() DB
```

GetDB returns the underlying DB.

​	GetDB 返回基础数据库。

#### (*Core) GetDebug

```go
func (c *Core) GetDebug() bool
```

GetDebug returns the debug value.

​	GetDebug 返回调试值。

#### (*Core) GetDryRun

```go
func (c *Core) GetDryRun() bool
```

GetDryRun returns the DryRun value.

​	GetDryRun 返回 DryRun 值。

#### (*Core) GetFieldType

```go
func (c *Core) GetFieldType(ctx context.Context, fieldName, table, schema string) *TableField
```

GetFieldType retrieves and returns the field type object for certain field by name.

​	GetFieldType 按名称检索并返回特定字段的字段类型对象。

#### (*Core) GetFieldTypeStr

```go
func (c *Core) GetFieldTypeStr(ctx context.Context, fieldName, table, schema string) string
```

GetFieldTypeStr retrieves and returns the field type string for certain field by name.

​	GetFieldTypeStr 按名称检索并返回特定字段的字段类型字符串。

#### (*Core) GetGroup

```go
func (c *Core) GetGroup() string
```

GetGroup returns the group string configured.

​	GetGroup 返回配置的组字符串。

#### (*Core) GetIgnoreResultFromCtx

```go
func (c *Core) GetIgnoreResultFromCtx(ctx context.Context) bool
```

#### (*Core) GetInternalCtxDataFromCtx

```go
func (c *Core) GetInternalCtxDataFromCtx(ctx context.Context) *internalCtxData
```

#### (*Core) GetLink

```go
func (c *Core) GetLink(ctx context.Context, master bool, schema string) (Link, error)
```

GetLink creates and returns the underlying database link object with transaction checks. The parameter `master` specifies whether using the master node if master-slave configured.

​	GetLink 创建并返回具有事务检查的基础数据库链接对象。该参数 `master` 指定是否在配置了主从节点的情况下使用主节点。

#### (*Core) GetLogger

```go
func (c *Core) GetLogger() glog.ILogger
```

GetLogger returns the (logger) of the orm.

​	GetLogger 返回 orm 的 （logger）。

#### (*Core) GetOne

```go
func (c *Core) GetOne(ctx context.Context, sql string, args ...interface{}) (Record, error)
```

GetOne queries and returns one record from database.

​	GetOne 查询并返回数据库中的一条记录。

#### (*Core) GetPrefix

```go
func (c *Core) GetPrefix() string
```

GetPrefix returns the table prefix string configured.

​	GetPrefix 返回配置的表前缀字符串。

#### (*Core) GetScan

```go
func (c *Core) GetScan(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error
```

GetScan queries one or more records from database and converts them to given struct or struct array.

​	GetScan 从数据库中查询一条或多条记录，并将它们转换为给定的 struct 或 struct 数组。

If parameter `pointer` is type of struct pointer, it calls GetStruct internally for the conversion. If parameter `pointer` is type of slice, it calls GetStructs internally for conversion.

​	如果 parameter `pointer` 是结构指针的类型，则它会在内部调用 GetStruct 进行转换。如果 parameter `pointer` 是切片的类型，则它会在内部调用 GetStructs 进行转换。

#### (*Core) GetSchema

```go
func (c *Core) GetSchema() string
```

GetSchema returns the schema configured.

​	GetSchema 返回配置的架构。

#### (*Core) GetTablesWithCache

```go
func (c *Core) GetTablesWithCache() ([]string, error)
```

GetTablesWithCache retrieves and returns the table names of current database with cache.

​	GetTablesWithCache 检索并返回具有缓存的当前数据库的表名。

#### (*Core) GetValue

```go
func (c *Core) GetValue(ctx context.Context, sql string, args ...interface{}) (Value, error)
```

GetValue queries and returns the field value from database. The sql should query only one field from database, or else it returns only one field of the result.

​	GetValue 查询并返回数据库中的字段值。sql 应该只从数据库中查询一个字段，否则它只返回结果的一个字段。

#### (*Core) HasField

```go
func (c *Core) HasField(ctx context.Context, table, field string, schema ...string) (bool, error)
```

HasField determine whether the field exists in the table.

​	HasField 确定表中是否存在该字段。

#### (*Core) HasTable

```go
func (c *Core) HasTable(name string) (bool, error)
```

HasTable determine whether the table name exists in the database.

​	HasTable 确定数据库中是否存在表名。

#### (*Core) InjectIgnoreResult

```go
func (c *Core) InjectIgnoreResult(ctx context.Context) context.Context
```

#### (*Core) InjectInternalCtxData

```go
func (c *Core) InjectInternalCtxData(ctx context.Context) context.Context
```

#### (*Core) Insert

```go
func (c *Core) Insert(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)
```

Insert does “INSERT INTO …” statement for the table. If there’s already one unique record of the data in the table, it returns error.

​	Insert 执行“INSERT INTO ...”表的语句。如果表中已有一条数据的唯一记录，则返回错误。

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”})

​	参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。 例如： Data（g.Map{“uid”： 10000， “name”：“john”}） 数据（g.Slice{g.Map{“uid”： 10000， “name”：“john”}， g.Map{“uid”： 20000， “name”：“smith”}）

The parameter `batch` specifies the batch operation count when given data is slice.

​	该参数 `batch` 指定给定数据为切片时的批处理操作计数。

#### (*Core) InsertAndGetId

```go
func (c *Core) InsertAndGetId(ctx context.Context, table string, data interface{}, batch ...int) (int64, error)
```

InsertAndGetId performs action Insert and returns the last insert id that automatically generated.

​	InsertAndGetId 执行操作 Insert 并返回自动生成的最后一个插入 ID。

#### (*Core) InsertIgnore

```go
func (c *Core) InsertIgnore(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)
```

InsertIgnore does “INSERT IGNORE INTO …” statement for the table. If there’s already one unique record of the data in the table, it ignores the inserting.

​	InsertIgnore 执行“INSERT IGNORE INTO ...”表的语句。如果表中已经有一条数据的唯一记录，则会忽略插入。

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”})

​	参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。 例如： Data（g.Map{“uid”： 10000， “name”：“john”}） 数据（g.Slice{g.Map{“uid”： 10000， “name”：“john”}， g.Map{“uid”： 20000， “name”：“smith”}）

The parameter `batch` specifies the batch operation count when given data is slice.

​	该参数 `batch` 指定给定数据为切片时的批处理操作计数。

#### (*Core) IsSoftCreatedFieldName

```go
func (c *Core) IsSoftCreatedFieldName(fieldName string) bool
```

IsSoftCreatedFieldName checks and returns whether given field name is an automatic-filled created time.

​	IsSoftCreatedFieldName 检查并返回给定的字段名称是否为自动填充的创建时间。

#### (Core) MarshalJSON

```go
func (c Core) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. It just returns the pointer address.

​	MarshalJSON 实现 json 的接口 MarshalJSON。元帅。它只返回指针地址。

Note that this interface implements mainly for workaround for a json infinite loop bug of Golang version < v1.14.

​	请注意，此接口主要用于解决 Golang 版本 < v1.14 的 json 无限循环错误。

#### (*Core) Master

```go
func (c *Core) Master(schema ...string) (*sql.DB, error)
```

Master creates and returns a connection from master node if master-slave configured. It returns the default connection if master-slave not configured.

​	如果配置了主从节点，则主节点创建并返回连接。如果未配置主从连接，则返回默认连接。

#### (*Core) MasterLink

```go
func (c *Core) MasterLink(schema ...string) (Link, error)
```

MasterLink acts like function Master but with additional `schema` parameter specifying the schema for the connection. It is defined for internal usage. Also see Master.

​	MasterLink 的作用类似于函数 Master，但具有指定连接架构的附加 `schema` 参数。它被定义为供内部使用。另见师父。

#### (*Core) Model

```go
func (c *Core) Model(tableNameQueryOrStruct ...interface{}) *Model
```

Model creates and returns a new ORM model from given schema. The parameter `tableNameQueryOrStruct` can be more than one table names, and also alias name, like:

​	模型从给定模式创建并返回新的 ORM 模型。该参数 `tableNameQueryOrStruct` 可以是多个表名，也可以是别名，例如：

1. Model names: db.Model(“user”) db.Model(“user u”) db.Model(“user, user_detail”) db.Model(“user u, user_detail ud”)
   型号名称：db。Model（“user”） 数据库。Model（“user u”） 数据库。Model（“user， user_detail”） 数据库。Model（“用户 u， user_detail ud”）
2. Model name with alias: db.Model(“user”, “u”)
   别名为：db 的型号名称。模型（“用户”， “u”）
3. Model name with sub-query: db.Model(”? AS a, ? AS b”, subQuery1, subQuery2)
   带有子查询的模型名称：db。Model（“？作为，？AS b“、subQuery1、subQuery2）

#### (*Core) PingMaster

```go
func (c *Core) PingMaster() error
```

PingMaster pings the master node to check authentication or keeps the connection alive.

​	PingMaster 对主节点执行 ping 操作以检查身份验证或保持连接处于活动状态。

#### (*Core) PingSlave

```go
func (c *Core) PingSlave() error
```

PingSlave pings the slave node to check authentication or keeps the connection alive.

​	PingSlave 对从节点执行 ping 操作以检查身份验证或保持连接处于活动状态。

#### (*Core) Prepare

```go
func (c *Core) Prepare(ctx context.Context, sql string, execOnMaster ...bool) (*Stmt, error)
```

Prepare creates a prepared statement for later queries or executions. Multiple queries or executions may be run concurrently from the returned statement. The caller must call the statement’s Close method when the statement is no longer needed.

​	Prepare 为以后的查询或执行创建预准备语句。可以从返回的语句同时运行多个查询或执行。当不再需要语句时，调用方必须调用语句的 Close 方法。

The parameter `execOnMaster` specifies whether executing the sql on master node, or else it executes the sql on slave node if master-slave configured.

​	该参数 `execOnMaster` 指定是在主节点上执行 sql，还是在配置了主从节点时在从节点上执行 sql。

#### (*Core) Query

```go
func (c *Core) Query(ctx context.Context, sql string, args ...interface{}) (result Result, err error)
```

Query commits one query SQL to underlying driver and returns the execution result. It is most commonly used for data querying.

​	查询将一个查询 SQL 提交到基础驱动程序并返回执行结果。它最常用于数据查询。

#### (*Core) QuotePrefixTableName

```go
func (c *Core) QuotePrefixTableName(table string) string
```

QuotePrefixTableName adds prefix string and quotes chars for the table. It handles table string like: “user”, “user u”, “user,user_detail”, “user u, user_detail ut”, “user as u, user_detail as ut”.

​	QuotePrefixTableName 为表添加前缀字符串和引号字符。它处理表字符串，例如：“用户”、“用户 u”、“用户user_detail”、“用户 u、user_detail ut”、“用户作为 u，user_detail 作为 ut”。

Note that, this will automatically checks the table prefix whether already added, if true it does nothing to the table name, or else adds the prefix to the table name.

​	请注意，这将自动检查表前缀是否已添加，如果为 true，则不对表名执行任何操作，否则将前缀添加到表名中。

#### (*Core) QuoteString

```go
func (c *Core) QuoteString(s string) string
```

QuoteString quotes string with quote chars. Strings like: “user”, “user u”, “user,user_detail”, “user u, user_detail ut”, “u.id asc”.

​	QuoteString 带有引号字符的引号字符串。字符串如下：“user”、“user u”、“user，user_detail”、“user u， user_detail ut”、“u.id asc”。

The meaning of a `string` can be considered as part of a statement string including columns.

​	a `string` 的含义可以被视为语句字符串（包括列）的一部分。

#### (*Core) QuoteWord

```go
func (c *Core) QuoteWord(s string) string
```

QuoteWord checks given string `s` a word, if true it quotes `s` with security chars of the database and returns the quoted string; or else it returns `s` without any change.

​	QuoteWord 检查给定字符串 `s` 的单词，如果为 true，则 `s` 使用数据库的安全字符引用并返回带引号的字符串;否则，它将 `s` 返回而没有任何更改。

The meaning of a `word` can be considered as a column name.

​	a `word` 的含义可以视为列名。

#### (*Core) Raw

```go
func (c *Core) Raw(rawSql string, args ...interface{}) *Model
```

Raw creates and returns a model based on a raw sql not a table. Example:

​	raw 创建并返回基于原始 sql 而不是表的模型。例：

```
db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
```

#### (*Core) Replace

```go
func (c *Core) Replace(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)
```

Replace does “REPLACE INTO …” statement for the table. If there’s already one unique record of the data in the table, it deletes the record and inserts a new one.

​	替换执行“REPLACE INTO ...”表的语句。如果表中已有一条数据的唯一记录，则会删除该记录并插入一条新记录。

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”})

​	参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。 例如： Data（g.Map{“uid”： 10000， “name”：“john”}） 数据（g.Slice{g.Map{“uid”： 10000， “name”：“john”}， g.Map{“uid”： 20000， “name”：“smith”}）

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. If given data is type of slice, it then does batch replacing, and the optional parameter `batch` specifies the batch operation count.

​	参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。如果给定的数据是切片类型，则执行批量替换，可选参数 `batch` 指定批处理操作计数。

#### (*Core) RowsToResult

```go
func (c *Core) RowsToResult(ctx context.Context, rows *sql.Rows) (Result, error)
```

RowsToResult converts underlying data record type sql.Rows to Result type.

​	RowsToResult 转换基础数据记录类型 sql。行到结果类型。

#### (*Core) Save

```go
func (c *Core) Save(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)
```

Save does “INSERT INTO … ON DUPLICATE KEY UPDATE…” statement for the table. It updates the record if there’s primary or unique index in the saving data, or else it inserts a new record into the table.

​	保存执行“插入...在重复的密钥更新中......”表的语句。如果保存数据中有主索引或唯一索引，它将更新记录，或者将新记录插入到表中。

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”})

​	参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。 例如： Data（g.Map{“uid”： 10000， “name”：“john”}） 数据（g.Slice{g.Map{“uid”： 10000， “name”：“john”}， g.Map{“uid”： 20000， “name”：“smith”}）

If given data is type of slice, it then does batch saving, and the optional parameter `batch` specifies the batch operation count.

​	如果给定的数据是切片类型，则执行批量保存，可选参数 `batch` 指定批处理操作计数。

#### (*Core) Schema

```go
func (c *Core) Schema(schema string) *Schema
```

Schema creates and returns a schema.

​	架构创建并返回架构。

#### (*Core) SetDebug

```go
func (c *Core) SetDebug(debug bool)
```

SetDebug enables/disables the debug mode.

​	SetDebug 启用/禁用调试模式。

#### (*Core) SetDryRun

```go
func (c *Core) SetDryRun(enabled bool)
```

SetDryRun enables/disables the DryRun feature.

​	SetDryRun 启用/禁用 DryRun 功能。

#### (*Core) SetLogger

```go
func (c *Core) SetLogger(logger glog.ILogger)
```

SetLogger sets the logger for orm.

​	SetLogger 为 orm 设置记录器。

#### (*Core) SetMaxConnLifeTime

```go
func (c *Core) SetMaxConnLifeTime(d time.Duration)
```

SetMaxConnLifeTime sets the maximum amount of time a connection may be reused.

​	SetMaxConnLifeTime 设置连接可重用的最长时间。

Expired connections may be closed lazily before reuse.

​	过期的连接可能会在重用之前延迟关闭。

If d <= 0, connections are not closed due to a connection’s age.

​	如果 d <= 0，则连接不会因连接的年龄而关闭。

#### (*Core) SetMaxIdleConnCount

```go
func (c *Core) SetMaxIdleConnCount(n int)
```

SetMaxIdleConnCount sets the maximum number of connections in the idle connection pool.

​	SetMaxIdleConnCount 设置空闲连接池中的最大连接数。

If MaxOpenConns is greater than 0 but less than the new MaxIdleConns, then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.

​	如果 MaxOpenConns 大于 0 但小于新的 MaxIdleConns，则新的 MaxIdleConns 将减少以匹配 MaxOpenConns 限制。

If n <= 0, no idle connections are retained.

​	如果 n <= 0，则不保留空闲连接。

The default max idle connections is currently 2. This may change in a future release.

​	默认的最大空闲连接数当前为 2。这可能会在将来的版本中更改。

#### (*Core) SetMaxOpenConnCount

```go
func (c *Core) SetMaxOpenConnCount(n int)
```

SetMaxOpenConnCount sets the maximum number of open connections to the database.

​	SetMaxOpenConnCount 设置与数据库的最大打开连接数。

If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than MaxIdleConns, then MaxIdleConns will be reduced to match the new MaxOpenConns limit.

​	如果 MaxIdleConns 大于 0，并且新的 MaxOpenConns 小于 MaxIdleConns，则 MaxIdleConns 将减少以匹配新的 MaxOpenConns 限制。

If n <= 0, then there is no limit on the number of open connections. The default is 0 (unlimited).

​	如果 n <= 0，则打开的连接数没有限制。默认值为 0（无限制）。

#### (*Core) Slave

```go
func (c *Core) Slave(schema ...string) (*sql.DB, error)
```

Slave creates and returns a connection from slave node if master-slave configured. It returns the default connection if master-slave not configured.

​	如果配置了主从节点，则从节点创建并返回连接。如果未配置主从连接，则返回默认连接。

#### (*Core) SlaveLink

```go
func (c *Core) SlaveLink(schema ...string) (Link, error)
```

SlaveLink acts like function Slave but with additional `schema` parameter specifying the schema for the connection. It is defined for internal usage. Also see Slave.

​	SlaveLink 的作用类似于函数 Slave，但具有指定连接架构的附加 `schema` 参数。它被定义为供内部使用。另请参阅 Slave。

#### (*Core) TableFields

```go
func (c *Core) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*TableField, err error)
```

TableFields retrieves and returns the fields’ information of specified table of current schema.

​	TableFields 检索并返回当前架构的指定表的字段信息。

The parameter `link` is optional, if given nil it automatically retrieves a raw sql connection as its link to proceed necessary sql query.

​	该参数 `link` 是可选的，如果给定 nil，它会自动检索原始 sql 连接作为其链接以继续进行必要的 sql 查询。

Note that it returns a map containing the field name and its corresponding fields. As a map is unsorted, the TableField struct has an “Index” field marks its sequence in the fields.

​	请注意，它返回包含字段名称及其相应字段的映射。由于地图未排序，因此 TableField 结构有一个“索引”字段标记其在字段中的顺序。

It’s using cache feature to enhance the performance, which is never expired util the process restarts.

​	它使用缓存功能来增强性能，该功能永远不会过期，因为进程重新启动。

#### (*Core) Tables

```go
func (c *Core) Tables(ctx context.Context, schema ...string) (tables []string, err error)
```

Tables retrieves and returns the tables of current schema. It’s mainly used in cli tool chain for automatically generating the models.

​	表检索并返回当前架构的表。它主要用于 cli 工具链中，用于自动生成模型。

#### (*Core) Transaction

```go
func (c *Core) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error)
```

Transaction wraps the transaction logic using function `f`. It rollbacks the transaction and returns the error from function `f` if it returns non-nil error. It commits the transaction and returns nil if function `f` returns nil.

​	事务使用函数 `f` 包装事务逻辑。 `f` 它会回滚事务，如果函数返回非 nil 错误，则返回错误。它提交事务，如果函数 `f` 返回 nil，则返回 nil。

Note that, you should not Commit or Rollback the transaction in function `f` as it is automatically handled by this function.

​	请注意，您不应该在函数 `f` 中提交或回滚事务，因为它是由此函数自动处理的。

#### (*Core) Union

```go
func (c *Core) Union(unions ...*Model) *Model
```

Union does “(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) …” statement.

​	Union does “（SELECT xxx FROM xxx） UNION （SELECT xxx FROM xxx） ...”陈述。

#### (*Core) UnionAll

```go
func (c *Core) UnionAll(unions ...*Model) *Model
```

UnionAll does “(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) …” statement.

​	UnionAll 执行“（从 xxx 中选择 xxx）UNION ALL （SELECT XXX FROM XXX）...”陈述。

#### (*Core) Update

```go
func (c *Core) Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)
```

Update does “UPDATE … " statement for the table.

​	更新执行“更新...“的语句。

The parameter `data` can be type of string/map/gmap/struct/*struct, etc. Eg: “uid=10000”, “uid”, 10000, g.Map{“uid”: 10000, “name”:“john”}

​	参数 `data` 可以是 string/map/gmap/struct/*struct 等类型。例如： “uid=10000”， “uid”， 10000， g.Map{“uid”： 10000， “name”：“john”}

The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc. It is commonly used with parameter `args`. Eg: “uid=10000”, “uid”, 10000 “money>? AND name like ?”, 99999, “vip_%” “status IN (?)”, g.Slice{1,2,3} “age IN(?,?)”, 18, 50 User{ Id : 1, UserName : “john”}.

​	参数 `condition` 可以是字符串/地图/gmap/slice/struct/*struct等类型。它通常与参数一起使用 `args` 。例如：“uid=10000”、“uid”、10000“钱>？AND name like ？“， 99999， ”vip_%“ ”status IN （？）“， g.Slice{1,2,3} ”age IN（?,?）“， 18， 50 User{ Id ： 1， UserName ： ”john“}.

#### (*Core) With

```go
func (c *Core) With(objects ...interface{}) *Model
```

With creates and returns an ORM model based on metadata of given object.

​	With 创建并返回一个基于给定对象的元数据的 ORM 模型。

### type Counter

```go
type Counter struct {
	Field string
	Value float64
}
```

Counter is the type for update count.

​	计数器是更新计数的类型。

### type DB

```go
type DB interface {

	// Model creates and returns a new ORM model from given schema.
	// The parameter `table` can be more than one table names, and also alias name, like:
	// 1. Model names:
	//    Model("user")
	//    Model("user u")
	//    Model("user, user_detail")
	//    Model("user u, user_detail ud")
	// 2. Model name with alias: Model("user", "u")
	// Also see Core.Model.
	Model(tableNameOrStruct ...interface{}) *Model

	// Raw creates and returns a model based on a raw sql not a table.
	Raw(rawSql string, args ...interface{}) *Model

	// Schema creates and returns a schema.
	// Also see Core.Schema.
	Schema(schema string) *Schema

	// With creates and returns an ORM model based on metadata of given object.
	// Also see Core.With.
	With(objects ...interface{}) *Model

	// Open creates a raw connection object for database with given node configuration.
	// Note that it is not recommended using the function manually.
	// Also see DriverMysql.Open.
	Open(config *ConfigNode) (*sql.DB, error)

	// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
	// of current DB object and with given context in it.
	// Also see Core.Ctx.
	Ctx(ctx context.Context) DB

	// Close closes the database and prevents new queries from starting.
	// Close then waits for all queries that have started processing on the server
	// to finish.
	//
	// It is rare to Close a DB, as the DB handle is meant to be
	// long-lived and shared between many goroutines.
	Close(ctx context.Context) error

	Query(ctx context.Context, sql string, args ...interface{}) (Result, error)    // See Core.Query.
	Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) // See Core.Exec.
	Prepare(ctx context.Context, sql string, execOnMaster ...bool) (*Stmt, error)  // See Core.Prepare.

	Insert(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                               // See Core.Insert.
	InsertIgnore(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                         // See Core.InsertIgnore.
	InsertAndGetId(ctx context.Context, table string, data interface{}, batch ...int) (int64, error)                            // See Core.InsertAndGetId.
	Replace(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                              // See Core.Replace.
	Save(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)                                 // See Core.Save.
	Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error) // See Core.Update.
	Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (sql.Result, error)                   // See Core.Delete.

	DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)                                           // See Core.DoSelect.
	DoInsert(ctx context.Context, link Link, table string, data List, option DoInsertOption) (result sql.Result, err error)                        // See Core.DoInsert.
	DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error) // See Core.DoUpdate.
	DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result sql.Result, err error)                   // See Core.DoDelete.

	DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)    // See Core.DoQuery.
	DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result sql.Result, err error) // See Core.DoExec.

	DoFilter(ctx context.Context, link Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error) // See Core.DoFilter.
	DoCommit(ctx context.Context, in DoCommitInput) (out DoCommitOutput, err error)                                            // See Core.DoCommit.

	DoPrepare(ctx context.Context, link Link, sql string) (*Stmt, error) // See Core.DoPrepare.

	GetAll(ctx context.Context, sql string, args ...interface{}) (Result, error)                // See Core.GetAll.
	GetOne(ctx context.Context, sql string, args ...interface{}) (Record, error)                // See Core.GetOne.
	GetValue(ctx context.Context, sql string, args ...interface{}) (Value, error)               // See Core.GetValue.
	GetArray(ctx context.Context, sql string, args ...interface{}) ([]Value, error)             // See Core.GetArray.
	GetCount(ctx context.Context, sql string, args ...interface{}) (int, error)                 // See Core.GetCount.
	GetScan(ctx context.Context, objPointer interface{}, sql string, args ...interface{}) error // See Core.GetScan.
	Union(unions ...*Model) *Model                                                              // See Core.Union.
	UnionAll(unions ...*Model) *Model                                                           // See Core.UnionAll.

	Master(schema ...string) (*sql.DB, error) // See Core.Master.
	Slave(schema ...string) (*sql.DB, error)  // See Core.Slave.

	PingMaster() error // See Core.PingMaster.
	PingSlave() error  // See Core.PingSlave.

	Begin(ctx context.Context) (TX, error)                                           // See Core.Begin.
	Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) error // See Core.Transaction.

	GetCache() *gcache.Cache            // See Core.GetCache.
	SetDebug(debug bool)                // See Core.SetDebug.
	GetDebug() bool                     // See Core.GetDebug.
	GetSchema() string                  // See Core.GetSchema.
	GetPrefix() string                  // See Core.GetPrefix.
	GetGroup() string                   // See Core.GetGroup.
	SetDryRun(enabled bool)             // See Core.SetDryRun.
	GetDryRun() bool                    // See Core.GetDryRun.
	SetLogger(logger glog.ILogger)      // See Core.SetLogger.
	GetLogger() glog.ILogger            // See Core.GetLogger.
	GetConfig() *ConfigNode             // See Core.GetConfig.
	SetMaxIdleConnCount(n int)          // See Core.SetMaxIdleConnCount.
	SetMaxOpenConnCount(n int)          // See Core.SetMaxOpenConnCount.
	SetMaxConnLifeTime(d time.Duration) // See Core.SetMaxConnLifeTime.

	GetCtx() context.Context                                                                                 // See Core.GetCtx.
	GetCore() *Core                                                                                          // See Core.GetCore
	GetChars() (charLeft string, charRight string)                                                           // See Core.GetChars.
	Tables(ctx context.Context, schema ...string) (tables []string, err error)                               // See Core.Tables. The driver must implement this function.
	TableFields(ctx context.Context, table string, schema ...string) (map[string]*TableField, error)         // See Core.TableFields. The driver must implement this function.
	ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) // See Core.ConvertValueForField
	ConvertValueForLocal(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error) // See Core.ConvertValueForLocal
	CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (LocalType, error) // See Core.CheckLocalTypeForField
	FormatUpsert(columns []string, list List, option DoInsertOption) (string, error)                         // See Core.DoFormatUpsert
}
```

DB defines the interfaces for ORM operations.

​	DB 定义 ORM 操作的接口。

#### func DBFromCtx <-2.0.5

```go
func DBFromCtx(ctx context.Context) DB
```

DBFromCtx retrieves and returns DB object from context.

​	DBFromCtx 从上下文中检索并返回 DB 对象。

#### func Instance

```go
func Instance(name ...string) (db DB, err error)
```

Instance returns an instance for DB operations. The parameter `name` specifies the configuration group name, which is DefaultGroupName in default.

​	Instance 返回数据库操作的实例。该参数 `name` 指定配置组名称，默认为 DefaultGroupName。

#### func New

```go
func New(node ConfigNode) (db DB, err error)
```

New creates and returns an ORM object with given configuration node.

​	new 创建并返回具有给定配置节点的 ORM 对象。

#### func NewByGroup

```go
func NewByGroup(group ...string) (db DB, err error)
```

NewByGroup creates and returns an ORM object with global configurations. The parameter `name` specifies the configuration group name, which is DefaultGroupName in default.

​	NewByGroup 创建并返回具有全局配置的 ORM 对象。该参数 `name` 指定配置组名称，默认为 DefaultGroupName。

### type DoCommitInput

```go
type DoCommitInput struct {
	Db            *sql.DB
	Tx            *sql.Tx
	Stmt          *sql.Stmt
	Link          Link
	Sql           string
	Args          []interface{}
	Type          string
	IsTransaction bool
}
```

DoCommitInput is the input parameters for function DoCommit.

​	DoCommitInput 是函数 DoCommit 的输入参数。

### type DoCommitOutput

```go
type DoCommitOutput struct {
	Result    sql.Result  // Result is the result of exec statement.
	Records   []Record    // Records is the result of query statement.
	Stmt      *Stmt       // Stmt is the Statement object result for Prepare.
	Tx        TX          // Tx is the transaction object result for Begin.
	RawResult interface{} // RawResult is the underlying result, which might be sql.Result/*sql.Rows/*sql.Row.
}
```

DoCommitOutput is the output parameters for function DoCommit.

​	DoCommitOutput 是函数 DoCommit 的输出参数。

### type DoInsertOption

```go
type DoInsertOption struct {
	OnDuplicateStr string                 // Custom string for `on duplicated` statement.
	OnDuplicateMap map[string]interface{} // Custom key-value map from `OnDuplicateEx` function for `on duplicated` statement.
	OnConflict     []string               // Custom conflict key of upsert clause, if the database needs it.
	InsertOption   InsertOption           // Insert operation in constant value.
	BatchCount     int                    // Batch count for batch inserting.
}
```

DoInsertOption is the input struct for function DoInsert.

​	DoInsertOption 是函数 DoInsert 的输入结构。

### type Driver

```go
type Driver interface {
	// New creates and returns a database object for specified database server.
	New(core *Core, node *ConfigNode) (DB, error)
}
```

Driver is the interface for integrating sql drivers into package gdb.

​	驱动程序是用于将 sql 驱动程序集成到包 gdb 中的接口。

### type DriverDefault <-2.2.0

```go
type DriverDefault struct {
	*Core
}
```

DriverDefault is the default driver for mysql database, which does nothing.

​	DriverDefault 是 mysql 数据库的默认驱动程序，它不执行任何操作。

#### (*DriverDefault) New

```go
func (d *DriverDefault) New(core *Core, node *ConfigNode) (DB, error)
```

New creates and returns a database object for mysql. It implements the interface of gdb.Driver for extra database driver installation.

​	New 创建并返回 mysql 的数据库对象。它实现了 gdb 的接口。用于额外数据库驱动程序安装的驱动程序。

#### (*DriverDefault) Open

```go
func (d *DriverDefault) Open(config *ConfigNode) (db *sql.DB, err error)
```

Open creates and returns an underlying sql.DB object for mysql. Note that it converts time.Time argument to local timezone in default.

​	Open 创建并返回基础 sql。mysql 的数据库对象。请注意，它会转换时间。默认情况下，时间参数设置为本地时区。

#### (*DriverDefault) PingMaster

```go
func (d *DriverDefault) PingMaster() error
```

PingMaster pings the master node to check authentication or keeps the connection alive.

​	PingMaster 对主节点执行 ping 操作以检查身份验证或保持连接处于活动状态。

#### (*DriverDefault) PingSlave

```go
func (d *DriverDefault) PingSlave() error
```

PingSlave pings the slave node to check authentication or keeps the connection alive.

​	PingSlave 对从节点执行 ping 操作以检查身份验证或保持连接处于活动状态。

### type DriverWrapper <-2.2.0

```go
type DriverWrapper struct {
	// contains filtered or unexported fields
}
```

DriverWrapper is a driver wrapper for extending features with embedded driver.

​	DriverWrapper 是一个驱动程序包装器，用于使用嵌入式驱动程序扩展功能。

#### (*DriverWrapper) New

```go
func (d *DriverWrapper) New(core *Core, node *ConfigNode) (DB, error)
```

New creates and returns a database object for mysql. It implements the interface of gdb.Driver for extra database driver installation.

​	New 创建并返回 mysql 的数据库对象。它实现了 gdb 的接口。用于额外数据库驱动程序安装的驱动程序。

### type DriverWrapperDB <-2.2.0

```go
type DriverWrapperDB struct {
	DB
}
```

DriverWrapperDB is a DB wrapper for extending features with embedded DB.

​	DriverWrapperDB 是一个数据库包装器，用于使用嵌入式数据库扩展功能。

#### (*DriverWrapperDB) DoInsert

```go
func (d *DriverWrapperDB) DoInsert(ctx context.Context, link Link, table string, list List, option DoInsertOption) (result sql.Result, err error)
```

DoInsert inserts or updates data forF given table. This function is usually used for custom interface definition, you do not need call it manually. The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”})

​	DoInsert 插入或更新 F 给定表的数据。此函数通常用于自定义接口定义，无需手动调用。参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。 例如： Data（g.Map{“uid”： 10000， “name”：“john”}） 数据（g.Slice{g.Map{“uid”： 10000， “name”：“john”}， g.Map{“uid”： 20000， “name”：“smith”}）

The parameter `option` values are as follows: InsertOptionDefault: just insert, if there’s unique/primary key in the data, it returns error; InsertOptionReplace: if there’s unique/primary key in the data, it deletes it from table and inserts a new one; InsertOptionSave: if there’s unique/primary key in the data, it updates it or else inserts a new one; InsertOptionIgnore: if there’s unique/primary key in the data, it ignores the inserting;

​	参数 `option` 值如下： InsertOptionDefault：直接插入，如果数据中有唯一/主键，则返回错误;InsertOptionReplace：如果数据中有唯一键/主键，则将其从表中删除并插入新键;InsertOptionSave：如果数据中存在唯一/主键，则更新它或插入新键;InsertOptionIgnore：如果数据中有唯一/主键，则忽略插入;

#### (*DriverWrapperDB) Open

```go
func (d *DriverWrapperDB) Open(node *ConfigNode) (db *sql.DB, err error)
```

Open creates and returns an underlying sql.DB object for pgsql. https://pkg.go.dev/github.com/lib/pq

​	Open 创建并返回基础 sql。pgsql 的 DB 对象。https://pkg.go.dev/github.com/lib/pq

#### (*DriverWrapperDB) TableFields

```go
func (d *DriverWrapperDB) TableFields(
	ctx context.Context, table string, schema ...string,
) (fields map[string]*TableField, err error)
```

TableFields retrieves and returns the fields’ information of specified table of current schema.

​	TableFields 检索并返回当前架构的指定表的字段信息。

The parameter `link` is optional, if given nil it automatically retrieves a raw sql connection as its link to proceed necessary sql query.

​	该参数 `link` 是可选的，如果给定 nil，它会自动检索原始 sql 连接作为其链接以继续进行必要的 sql 查询。

Note that it returns a map containing the field name and its corresponding fields. As a map is unsorted, the TableField struct has an “Index” field marks its sequence in the fields.

​	请注意，它返回包含字段名称及其相应字段的映射。由于地图未排序，因此 TableField 结构有一个“索引”字段标记其在字段中的顺序。

It’s using cache feature to enhance the performance, which is never expired util the process restarts.

​	它使用缓存功能来增强性能，该功能永远不会过期，因为进程重新启动。

#### (*DriverWrapperDB) Tables

```go
func (d *DriverWrapperDB) Tables(ctx context.Context, schema ...string) (tables []string, err error)
```

Tables retrieves and returns the tables of current schema. It’s mainly used in cli tool chain for automatically generating the models.

​	表检索并返回当前架构的表。它主要用于 cli 工具链中，用于自动生成模型。

### type HookDeleteInput <-2.0.5

```go
type HookDeleteInput struct {
	Model     *Model        // Current operation Model.
	Table     string        // The table name that to be used. Update this attribute to change target table name.
	Schema    string        // The schema name that to be used. Update this attribute to change target schema name.
	Condition string        // The where condition string for deleting.
	Args      []interface{} // The arguments for sql place-holders.
	// contains filtered or unexported fields
}
```

HookDeleteInput holds the parameters for delete hook operation.

​	HookDeleteInput 保存删除挂钩操作的参数。

#### (*HookDeleteInput) Next

```go
func (h *HookDeleteInput) Next(ctx context.Context) (result sql.Result, err error)
```

Next calls the next hook handler.

​	接下来调用下一个挂钩处理程序。

### type HookFuncDelete <-2.0.5

```go
type HookFuncDelete func(ctx context.Context, in *HookDeleteInput) (result sql.Result, err error)
```

### type HookFuncInsert <-2.0.5

```go
type HookFuncInsert func(ctx context.Context, in *HookInsertInput) (result sql.Result, err error)
```

### type HookFuncSelect <-2.0.5

```go
type HookFuncSelect func(ctx context.Context, in *HookSelectInput) (result Result, err error)
```

### type HookFuncUpdate <-2.0.5

```go
type HookFuncUpdate func(ctx context.Context, in *HookUpdateInput) (result sql.Result, err error)
```

### type HookHandler <-2.0.5

```go
type HookHandler struct {
	Select HookFuncSelect
	Insert HookFuncInsert
	Update HookFuncUpdate
	Delete HookFuncDelete
}
```

HookHandler manages all supported hook functions for Model.

​	HookHandler 管理 Model 支持的所有钩子函数。

### type HookInsertInput <-2.0.5

```go
type HookInsertInput struct {
	Model  *Model         // Current operation Model.
	Table  string         // The table name that to be used. Update this attribute to change target table name.
	Schema string         // The schema name that to be used. Update this attribute to change target schema name.
	Data   List           // The data records list to be inserted/saved into table.
	Option DoInsertOption // The extra option for data inserting.
	// contains filtered or unexported fields
}
```

HookInsertInput holds the parameters for insert hook operation.

​	HookInsertInput 保存插入挂钩操作的参数。

#### (*HookInsertInput) Next

```go
func (h *HookInsertInput) Next(ctx context.Context) (result sql.Result, err error)
```

Next calls the next hook handler.

​	接下来调用下一个挂钩处理程序。

### type HookSelectInput <-2.0.5

```go
type HookSelectInput struct {
	Model  *Model        // Current operation Model.
	Table  string        // The table name that to be used. Update this attribute to change target table name.
	Schema string        // The schema name that to be used. Update this attribute to change target schema name.
	Sql    string        // The sql string that to be committed.
	Args   []interface{} // The arguments of sql.
	// contains filtered or unexported fields
}
```

HookSelectInput holds the parameters for select hook operation. Note that, COUNT statement will also be hooked by this feature, which is usually not be interesting for upper business hook handler.

​	HookSelectInput 保存用于选择挂钩操作的参数。需要注意的是，COUNT 语句也会被这个特性所钩住，这通常对上层业务钩子处理程序来说并不有趣。

#### (*HookSelectInput) Next

```go
func (h *HookSelectInput) Next(ctx context.Context) (result Result, err error)
```

Next calls the next hook handler.

​	接下来调用下一个挂钩处理程序。

### type HookUpdateInput <-2.0.5

```go
type HookUpdateInput struct {
	Model     *Model        // Current operation Model.
	Table     string        // The table name that to be used. Update this attribute to change target table name.
	Schema    string        // The schema name that to be used. Update this attribute to change target schema name.
	Data      interface{}   // Data can be type of: map[string]interface{}/string. You can use type assertion on `Data`.
	Condition string        // The where condition string for updating.
	Args      []interface{} // The arguments for sql place-holders.
	// contains filtered or unexported fields
}
```

HookUpdateInput holds the parameters for update hook operation.

​	HookUpdateInput 保存更新挂钩操作的参数。

#### (*HookUpdateInput) Next

```go
func (h *HookUpdateInput) Next(ctx context.Context) (result sql.Result, err error)
```

Next calls the next hook handler.

​	接下来调用下一个挂钩处理程序。

### type InsertOption <-2.5.0

```go
type InsertOption int
const (
	InsertOptionDefault InsertOption = iota
	InsertOptionReplace
	InsertOptionSave
	InsertOptionIgnore
)
```

### type Link

```go
type Link interface {
	QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, sql string) (*sql.Stmt, error)
	IsOnMaster() bool
	IsTransaction() bool
}
```

Link is a common database function wrapper interface. Note that, any operation using `Link` will have no SQL logging.

​	Link 是一种常见的数据库函数包装接口。请注意，任何操作都将 `Link` 没有 SQL 日志记录。

### type List

```go
type List = []Map // List is type of map array.
```

### type LocalType <-2.5.3

```go
type LocalType string
const (
	LocalTypeUndefined   LocalType = ""
	LocalTypeString      LocalType = "string"
	LocalTypeDate        LocalType = "date"
	LocalTypeDatetime    LocalType = "datetime"
	LocalTypeInt         LocalType = "int"
	LocalTypeUint        LocalType = "uint"
	LocalTypeInt64       LocalType = "int64"
	LocalTypeUint64      LocalType = "uint64"
	LocalTypeIntSlice    LocalType = "[]int"
	LocalTypeInt64Slice  LocalType = "[]int64"
	LocalTypeUint64Slice LocalType = "[]uint64"
	LocalTypeInt64Bytes  LocalType = "int64-bytes"
	LocalTypeUint64Bytes LocalType = "uint64-bytes"
	LocalTypeFloat32     LocalType = "float32"
	LocalTypeFloat64     LocalType = "float64"
	LocalTypeBytes       LocalType = "[]byte"
	LocalTypeBool        LocalType = "bool"
	LocalTypeJson        LocalType = "json"
	LocalTypeJsonb       LocalType = "jsonb"
)
```

### type Map

```go
type Map = map[string]interface{} // Map is alias of map[string]interface{}, which is the most common usage map type.
```

### type Model

```go
type Model struct {
	// contains filtered or unexported fields
}
```

Model is core struct implementing the DAO for ORM.

​	模型是实现 ORM 的 DAO 的核心结构。

#### (*Model) All

```go
func (m *Model) All(where ...interface{}) (Result, error)
```

All does “SELECT FROM …” statement for the model. It retrieves the records from table and returns the result as slice type. It returns nil if there’s no record retrieved with the given conditions from table.

​	所有“选择从...”都做模型的语句。它从表中检索记录，并以切片类型返回结果。如果没有从表中检索到具有给定条件的记录，则返回 nil。

The optional parameter `where` is the same as the parameter of Model.Where function, see Model.Where.

​	可选参数 `where` 与 Model.Where 函数的参数相同，参见 Model.Where。

#### (*Model) AllAndCount

```go
func (m *Model) AllAndCount(useFieldForCount bool) (result Result, totalCount int, err error)
```

AllAndCount retrieves all records and the total count of records from the model. If useFieldForCount is true, it will use the fields specified in the model for counting; otherwise, it will use a constant value of 1 for counting. It returns the result as a slice of records, the total count of records, and an error if any. The where parameter is an optional list of conditions to use when retrieving records.

​	AllAndCount 从模型中检索所有记录和记录总数。如果 useFieldForCount 为 true，它将使用模型中指定的字段进行计数;否则，它将使用常量值 1 进行计数。它将结果作为记录切片、记录总数和错误（如果有）返回。where 参数是检索记录时要使用的条件的可选列表。

Example:

​	例：

```go
var model Model
var result Result
var count int
where := []interface{}{"name = ?", "John"}
result, count, err := model.AllAndCount(true)
if err != nil {
    // Handle error.
}
fmt.Println(result, count)
```

#### (*Model) Args

```go
func (m *Model) Args(args ...interface{}) *Model
```

Args sets custom arguments for model operation.

​	Args 为模型操作设置自定义参数。

#### (*Model) Array

```go
func (m *Model) Array(fieldsAndWhere ...interface{}) ([]Value, error)
```

Array queries and returns data values as slice from database. Note that if there are multiple columns in the result, it returns just one column values randomly.

​	数组查询数据值并将其作为数据库中的切片返回。请注意，如果结果中有多个列，则仅随机返回一个列值。

If the optional parameter `fieldsAndWhere` is given, the fieldsAndWhere[0] is the selected fields and fieldsAndWhere[1:] is treated as where condition fields. Also see Model.Fields and Model.Where functions.

​	如果给定了可选参数 `fieldsAndWhere` ，则 fieldsAndWhere[0] 是所选字段，fieldsAndWhere[1：] 被视为 where 条件字段。另请参阅 Model.Fields 和 Model.Where 函数。

#### (*Model) As

```go
func (m *Model) As(as string) *Model
```

As sets an alias name for current table.

​	As 设置当前表的别名。

#### (*Model) Avg

```go
func (m *Model) Avg(column string) (float64, error)
```

Avg does “SELECT AVG(x) FROM …” statement for the model.

​	平均执行“SELECT AVG（x） FROM ...”模型的语句。

#### (*Model) Batch

```go
func (m *Model) Batch(batch int) *Model
```

Batch sets the batch operation number for the model.

​	Batch 设置模型的批处理操作编号。

#### (*Model) Builder

```go
func (m *Model) Builder() *WhereBuilder
```

Builder creates and returns a WhereBuilder. Please note that the builder is chain-safe.

​	Builder 创建并返回 WhereBuilder。请注意，建筑商是链条安全的。

#### (*Model) Cache

```go
func (m *Model) Cache(option CacheOption) *Model
```

Cache sets the cache feature for the model. It caches the result of the sql, which means if there’s another same sql request, it just reads and returns the result from cache, it but not committed and executed into the database.

​	缓存设置模型的缓存特征。它缓存 sql 的结果，这意味着如果有另一个相同的 sql 请求，它只是从缓存中读取并返回结果，但不会提交并执行到数据库中。

Note that, the cache feature is disabled if the model is performing select statement on a transaction.

​	请注意，如果模型对事务执行 select 语句，则缓存功能将被禁用。

#### (*Model) Chunk

```go
func (m *Model) Chunk(size int, handler ChunkHandler)
```

Chunk iterates the query result with given `size` and `handler` function.

​	Chunk 使用 given `size` 和 `handler` function 迭代查询结果。

#### (*Model) Clone

```go
func (m *Model) Clone() *Model
```

Clone creates and returns a new model which is a Clone of current model. Note that it uses deep-copy for the Clone.

​	克隆创建并返回一个新模型，该模型是当前模型的克隆。请注意，它对克隆使用深层复制。

#### (*Model) Count

```go
func (m *Model) Count(where ...interface{}) (int, error)
```

Count does “SELECT COUNT(x) FROM …” statement for the model. The optional parameter `where` is the same as the parameter of Model.Where function, see Model.Where.

​	Count 执行“SELECT COUNT（x） FROM ...”。模型的语句。可选参数 `where` 与 Model.Where 函数的参数相同，参见 Model.Where。

#### (*Model) CountColumn

```go
func (m *Model) CountColumn(column string) (int, error)
```

CountColumn does “SELECT COUNT(x) FROM …” statement for the model.

​	CountColumn 执行“SELECT COUNT（x） FROM ...”模型的语句。

#### (*Model) Ctx

```go
func (m *Model) Ctx(ctx context.Context) *Model
```

Ctx sets the context for current operation.

​	Ctx 设置当前操作的上下文。

#### (*Model) DB

```go
func (m *Model) DB(db DB) *Model
```

DB sets/changes the db object for current operation.

​	DB 设置/更改当前操作的 db 对象。

#### (*Model) Data

```go
func (m *Model) Data(data ...interface{}) *Model
```

Data sets the operation data for the model. The parameter `data` can be type of string/map/gmap/slice/struct/*struct, etc. Note that, it uses shallow value copying for `data` if `data` is type of map/slice to avoid changing it inside function. Eg: Data(“uid=10000”) Data(“uid”, 10000) Data(“uid=? AND name=?”, 10000, “john”) Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”}).

​	数据设置模型的操作数据。参数 `data` 可以是字符串/地图/gmap/slice/struct/*struct等类型。请注意，它对 `data` if `data` 是映射/切片的类型使用浅层值复制，以避免在函数中更改它。例如： Data（“uid=10000”） data（“uid”， 10000） data（“uid=？AND name=？“， 10000， ”john“） Data（g.Map{”uid“： 10000， ”name“：”john“}） Data（g.Slice{g.Map{”uid“： 10000， ”name“：”john“}， g.Map{”uid“： 20000， ”name“：”smith“}）。

#### (*Model) Decrement

```go
func (m *Model) Decrement(column string, amount interface{}) (sql.Result, error)
```

Decrement decrements a column’s value by a given amount. The parameter `amount` can be type of float or integer.

​	递减将列的值递减给定的量。参数 `amount` 可以是 float 类型或整数类型。

#### (*Model) Delete

```go
func (m *Model) Delete(where ...interface{}) (result sql.Result, err error)
```

Delete does “DELETE FROM … " statement for the model. The optional parameter `where` is the same as the parameter of Model.Where function, see Model.Where.

​	Delete 执行“DELETE FROM ...“的语句。可选参数 `where` 与 Model.Where 函数的参数相同，参见 Model.Where。

#### (*Model) Distinct

```go
func (m *Model) Distinct() *Model
```

Distinct forces the query to only return distinct results.

​	distinct 强制查询仅返回不同的结果。

#### (*Model) FieldAvg

```go
func (m *Model) FieldAvg(column string, as ...string) *Model
```

FieldAvg formats and appends commonly used field `AVG(column)` to the select fields of model.

​	FieldAvg 格式化常用字段 `AVG(column)` 并将其附加到模型的选定字段。

#### (*Model) FieldCount

```go
func (m *Model) FieldCount(column string, as ...string) *Model
```

FieldCount formats and appends commonly used field `COUNT(column)` to the select fields of model.

​	FieldCount 设置常用字段 `COUNT(column)` 的格式，并将其附加到模型的选定字段中。

#### (*Model) FieldMax

```go
func (m *Model) FieldMax(column string, as ...string) *Model
```

FieldMax formats and appends commonly used field `MAX(column)` to the select fields of model.

​	FieldMax 设置常用字段 `MAX(column)` 的格式，并将其附加到模型的选定字段中。

#### (*Model) FieldMin

```go
func (m *Model) FieldMin(column string, as ...string) *Model
```

FieldMin formats and appends commonly used field `MIN(column)` to the select fields of model.

​	FieldMin 格式化常用字段 `MIN(column)` 并将其附加到模型的选定字段中。

#### (*Model) FieldSum

```go
func (m *Model) FieldSum(column string, as ...string) *Model
```

FieldSum formats and appends commonly used field `SUM(column)` to the select fields of model.

​	FieldSum 设置常用字段 `SUM(column)` 的格式，并将其附加到模型的选定字段中。

#### (*Model) Fields

```go
func (m *Model) Fields(fieldNamesOrMapStruct ...interface{}) *Model
```

Fields appends `fieldNamesOrMapStruct` to the operation fields of the model, multiple fields joined using char ‘,’. The parameter `fieldNamesOrMapStruct` can be type of string/map/*map/struct/*struct.

​	字段附加 `fieldNamesOrMapStruct` 到模型的操作字段，多个字段使用 char '，' 连接。参数 `fieldNamesOrMapStruct` 的类型可以是 string/map/*map/struct/*struct。

Eg: Fields(“id”, “name”, “age”) Fields([]string{“id”, “name”, “age”}) Fields(map[string]interface{}{“id”:1, “name”:“john”, “age”:18}) Fields(User{ Id: 1, Name: “john”, Age: 18}).

​	蛋： fields（“id”， “name”， “before”） fields（[]string{“id”， “name”， “before”}） fields（map[string]interface{}{“id”：1， “name”：“jon”， “before”：18}） fields（user{ id： 1， name： “jon”， before： 18}）.

#### (*Model) FieldsEx

```go
func (m *Model) FieldsEx(fieldNamesOrMapStruct ...interface{}) *Model
```

FieldsEx appends `fieldNamesOrMapStruct` to the excluded operation fields of the model, multiple fields joined using char ‘,’. Note that this function supports only single table operations. The parameter `fieldNamesOrMapStruct` can be type of string/map/*map/struct/*struct.

​	FieldsEx 将多个字段附加 `fieldNamesOrMapStruct` 到模型的排除操作字段中，并使用 char '，' 连接多个字段。请注意，此函数仅支持单表操作。参数 `fieldNamesOrMapStruct` 的类型可以是 string/map/*map/struct/*struct。

Also see Fields.

​	另请参阅字段。

#### (*Model) FieldsExPrefix

```go
func (m *Model) FieldsExPrefix(prefixOrAlias string, fieldNamesOrMapStruct ...interface{}) *Model
```

FieldsExPrefix performs as function FieldsEx but add extra prefix for each field.

​	FieldsExPrefix 作为函数 FieldsEx 执行，但为每个字段添加额外的前缀。

#### (*Model) FieldsPrefix

```go
func (m *Model) FieldsPrefix(prefixOrAlias string, fieldNamesOrMapStruct ...interface{}) *Model
```

FieldsPrefix performs as function Fields but add extra prefix for each field.

​	FieldsPrefix 作为函数 Fields 执行，但为每个字段添加额外的前缀。

#### (*Model) GetCtx

```go
func (m *Model) GetCtx() context.Context
```

GetCtx returns the context for current Model. It returns `context.Background()` is there’s no context previously set.

​	GetCtx 返回当前模型的上下文。它返回 `context.Background()` 的是之前没有设置的上下文。

#### (*Model) GetFieldsExStr

```go
func (m *Model) GetFieldsExStr(fields string, prefix ...string) string
```

GetFieldsExStr retrieves and returns fields which are not in parameter `fields` from the table, joined with char ‘,’. The parameter `fields` specifies the fields that are excluded. The optional parameter `prefix` specifies the prefix for each field, eg: FieldsExStr(“id”, “u.”).

​	GetFieldsExStr `fields` 检索并返回表中不在参数中的字段，并用 char '，' 连接。该参数 `fields` 指定排除的字段。可选参数 `prefix` 指定每个字段的前缀，例如：FieldsExStr（“id”， “u.”）。

#### (*Model) GetFieldsStr

```go
func (m *Model) GetFieldsStr(prefix ...string) string
```

GetFieldsStr retrieves and returns all fields from the table, joined with char ‘,’. The optional parameter `prefix` specifies the prefix for each field, eg: GetFieldsStr(“u.”).

​	GetFieldsStr 检索并返回表中的所有字段，并用 char '，' 连接。可选参数 `prefix` 指定每个字段的前缀，例如：GetFieldsStr（“u.”）。

#### (*Model) Group

```go
func (m *Model) Group(groupBy ...string) *Model
```

Group sets the “GROUP BY” statement for the model.

​	Group 为模型设置“GROUP BY”语句。

#### (*Model) Handler

```go
func (m *Model) Handler(handlers ...ModelHandler) *Model
```

Handler calls each of `handlers` on current Model and returns a new Model. ModelHandler is a function that handles given Model and returns a new Model that is custom modified.

​	处理程序调用当前模型上的每个 `handlers` 模型并返回一个新模型。ModelHandler 是一个函数，用于处理给定的模型并返回自定义修改的新模型。

#### (*Model) HasField

```go
func (m *Model) HasField(field string) (bool, error)
```

HasField determine whether the field exists in the table.

​	HasField 确定表中是否存在该字段。

#### (*Model) Having

```go
func (m *Model) Having(having interface{}, args ...interface{}) *Model
```

Having sets the having statement for the model. The parameters of this function usage are as the same as function Where. See Where.

​	为模型设置了 having 语句。此函数用法的参数与函数 Where 相同。查看位置。

#### (*Model) Hook

```go
func (m *Model) Hook(hook HookHandler) *Model
```

Hook sets the hook functions for current model.

​	Hook 设置当前模型的钩子函数。

#### (*Model) Increment

```go
func (m *Model) Increment(column string, amount interface{}) (sql.Result, error)
```

Increment increments a column’s value by a given amount. The parameter `amount` can be type of float or integer.

​	Increment 将列的值递增给定的量。参数 `amount` 可以是 float 类型或整数类型。

#### (*Model) InnerJoin

```go
func (m *Model) InnerJoin(tableOrSubQueryAndJoinConditions ...string) *Model
```

InnerJoin does “INNER JOIN … ON …” statement on the model. The parameter `table` can be joined table and its joined condition, and also with its alias name。

​	InnerJoin 做“INNER JOIN ...在......”关于模型的声明。该参数 `table` 可以连接表及其连接条件，也可以连接其别名。

Eg: Model(“user”).InnerJoin(“user_detail”, “user_detail.uid=user.uid”) Model(“user”, “u”).InnerJoin(“user_detail”, “ud”, “ud.uid=u.uid”) Model(“user”, “u”).InnerJoin(“SELECT xxx FROM xxx”,“a”, “a.uid=u.uid”).

​	例如：model（“user”）。InnerJoin（“user_detail”， “user_detail.uid=user.uid”） Model（“用户”， “u”）.InnerJoin（“user_detail”， “ud”， “ud.uid=u.uid”） Model（“用户”， “u”）.InnerJoin（“从 xxx 中选择 xxx”，“a”， “a.uid=u.uid”）。

#### (*Model) InnerJoinOnField

```go
func (m *Model) InnerJoinOnField(table, field string) *Model
```

InnerJoinOnField performs as InnerJoin, but it joins both tables with the `same field name`.

​	InnerJoinOnField 的执行方式为 InnerJoin，但它使用 `same field name` .

Eg: Model(“order”).InnerJoinOnField(“user”, “user_id”) Model(“order”).InnerJoinOnField(“product”, “product_id”).

​	例如：model（“order”）。InnerJoinOnField（“用户”， “user_id”） Model（“order”）.InnerJoinOnField（“产品”， “product_id”）.

#### (*Model) InnerJoinOnFields

```go
func (m *Model) InnerJoinOnFields(table, firstField, operator, secondField string) *Model
```

InnerJoinOnFields performs as InnerJoin. It specifies different fields and comparison operator.

​	InnerJoinOnFields 以 InnerJoin 的形式执行。它指定不同的字段和比较运算符。

Eg: Model(“user”).InnerJoinOnFields(“order”, “id”, “=”, “user_id”) Model(“user”).InnerJoinOnFields(“order”, “id”, “>”, “user_id”) Model(“user”).InnerJoinOnFields(“order”, “id”, “<”, “user_id”)

​	例如：model（“user”）。InnerJoinOnFields（“order”， “id”， “=”， “user_id”） Model（“用户”）。InnerJoinOnFields（“order”， “id”， “>”， “user_id”） Model（“用户”）.InnerJoinOnFields（“订单”， “id”， “<”， “user_id”）

#### (*Model) Insert

```go
func (m *Model) Insert(data ...interface{}) (result sql.Result, err error)
```

Insert does “INSERT INTO …” statement for the model. The optional parameter `data` is the same as the parameter of Model.Data function, see Model.Data.

​	Insert 执行“INSERT INTO ...”模型的语句。可选参数 `data` 与 Model.Data 函数的参数相同，请参见 Model.Data。

#### (*Model) InsertAndGetId

```go
func (m *Model) InsertAndGetId(data ...interface{}) (lastInsertId int64, err error)
```

InsertAndGetId performs action Insert and returns the last insert id that automatically generated.

​	InsertAndGetId 执行操作 Insert 并返回自动生成的最后一个插入 ID。

#### (*Model) InsertIgnore

```go
func (m *Model) InsertIgnore(data ...interface{}) (result sql.Result, err error)
```

InsertIgnore does “INSERT IGNORE INTO …” statement for the model. The optional parameter `data` is the same as the parameter of Model.Data function, see Model.Data.

​	InsertIgnore 执行“INSERT IGNORE INTO ...”模型的语句。可选参数 `data` 与 Model.Data 函数的参数相同，请参见 Model.Data。

#### (*Model) LeftJoin

```go
func (m *Model) LeftJoin(tableOrSubQueryAndJoinConditions ...string) *Model
```

LeftJoin does “LEFT JOIN … ON …” statement on the model. The parameter `table` can be joined table and its joined condition, and also with its alias name.

​	LeftJoin 执行“LEFT JOIN ...在......”关于模型的声明。参数 `table` 可以是联接表及其联接条件，也可以是其别名。

Eg: Model(“user”).LeftJoin(“user_detail”, “user_detail.uid=user.uid”) Model(“user”, “u”).LeftJoin(“user_detail”, “ud”, “ud.uid=u.uid”) Model(“user”, “u”).LeftJoin(“SELECT xxx FROM xxx”,“a”, “a.uid=u.uid”).

​	例如：model（“user”）。LeftJoin（“user_detail”， “user_detail.uid=user.uid”） Model（“用户”， “u”）.LeftJoin（“user_detail”， “ud”， “ud.uid=u.uid”） Model（“用户”， “u”）.LeftJoin（“从 xxx 中选择 xxx”，“a”， “a.uid=u.uid”）。

#### (*Model) LeftJoinOnField

```go
func (m *Model) LeftJoinOnField(table, field string) *Model
```

LeftJoinOnField performs as LeftJoin, but it joins both tables with the `same field name`.

​	LeftJoinOnField 以 LeftJoin 的形式执行，但它使用 `same field name` .

Eg: Model(“order”).LeftJoinOnField(“user”, “user_id”) Model(“order”).LeftJoinOnField(“product”, “product_id”).

​	例如：model（“order”）。LeftJoinOnField（“用户”， “user_id”） Model（“order”）.LeftJoinOnField（“产品”， “product_id”）。

#### (*Model) LeftJoinOnFields

```go
func (m *Model) LeftJoinOnFields(table, firstField, operator, secondField string) *Model
```

LeftJoinOnFields performs as LeftJoin. It specifies different fields and comparison operator.

​	LeftJoinOnFields 以 LeftJoin 的形式执行。它指定不同的字段和比较运算符。

Eg: Model(“user”).LeftJoinOnFields(“order”, “id”, “=”, “user_id”) Model(“user”).LeftJoinOnFields(“order”, “id”, “>”, “user_id”) Model(“user”).LeftJoinOnFields(“order”, “id”, “<”, “user_id”)

​	例如：model（“user”）。LeftJoinOnFields（“order”， “id”， “=”， “user_id”） Model（“用户”）。LeftJoinOnFields（“order”， “id”， “>”， “user_id”） Model（“用户”）.LeftJoinOnFields（“订单”， “id”， “<”， “user_id”）

#### (*Model) Limit

```go
func (m *Model) Limit(limit ...int) *Model
```

Limit sets the “LIMIT” statement for the model. The parameter `limit` can be either one or two number, if passed two number is passed, it then sets “LIMIT limit[0],limit[1]” statement for the model, or else it sets “LIMIT limit[0]” statement.

​	Limit 设置模型的“LIMIT”语句。参数 `limit` 可以是一个或两个数字，如果传递两个数字，则为模型设置“LIMIT limit[0]，limit[1]”语句，或者设置“LIMIT limit[0]”语句。

#### (*Model) LockShared

```go
func (m *Model) LockShared() *Model
```

LockShared sets the lock in share mode for current operation.

​	LockShared 为当前操作设置共享模式下的锁定。

#### (*Model) LockUpdate

```go
func (m *Model) LockUpdate() *Model
```

LockUpdate sets the lock for update for current operation.

​	LockUpdate 为当前操作设置更新锁。

#### (*Model) Master

```go
func (m *Model) Master() *Model
```

Master marks the following operation on master node.

​	Master 在 master 节点上标记以下操作。

#### (*Model) Max

```go
func (m *Model) Max(column string) (float64, error)
```

Max does “SELECT MAX(x) FROM …” statement for the model.

​	Max 执行“SELECT MAX（x） FROM ...”模型的语句。

#### (*Model) Min

```go
func (m *Model) Min(column string) (float64, error)
```

Min does “SELECT MIN(x) FROM …” statement for the model.

​	Min 执行“SELECT MIN（x） FROM ...”模型的语句。

#### (*Model) Offset

```go
func (m *Model) Offset(offset int) *Model
```

Offset sets the “OFFSET” statement for the model. It only makes sense for some databases like SQLServer, PostgreSQL, etc.

​	Offset 设置模型的“OFFSET”语句。它只对某些数据库（如 SQLServer、PostgreSQL 等）有意义。

#### (*Model) OmitEmpty

```go
func (m *Model) OmitEmpty() *Model
```

OmitEmpty sets optionOmitEmpty option for the model, which automatically filers the data and where parameters for `empty` values.

​	OmitEmpty 为模型设置 optionOmitEmpty 选项，该选项会自动归档 `empty` 数据以及值的 where 参数。

#### (*Model) OmitEmptyData

```go
func (m *Model) OmitEmptyData() *Model
```

OmitEmptyData sets optionOmitEmptyData option for the model, which automatically filers the Data parameters for `empty` values.

​	OmitEmptyData 为模型设置 optionOmitEmptyData 选项，该选项会自动归 `empty` 档值的 Data 参数。

#### (*Model) OmitEmptyWhere

```go
func (m *Model) OmitEmptyWhere() *Model
```

OmitEmptyWhere sets optionOmitEmptyWhere option for the model, which automatically filers the Where/Having parameters for `empty` values.

​	OmitEmptyWhere 为模型设置 optionOmitEmptyWhere 选项，该选项会自动为 `empty` 值归档 Where/having 参数。

Eg:

```
Where("id", []int{}).All()             -> SELECT xxx FROM xxx WHERE 0=1
Where("name", "").All()                -> SELECT xxx FROM xxx WHERE `name`=''
OmitEmpty().Where("id", []int{}).All() -> SELECT xxx FROM xxx
OmitEmpty().("name", "").All()         -> SELECT xxx FROM xxx.
```

#### (*Model) OmitNil

```go
func (m *Model) OmitNil() *Model
```

OmitNil sets optionOmitNil option for the model, which automatically filers the data and where parameters for `nil` values.

​	OmitNil 为模型设置 optionsOmitNil 选项，该选项会自动归档 `nil` 数据以及值的参数。

#### (*Model) OmitNilData

```go
func (m *Model) OmitNilData() *Model
```

OmitNilData sets optionOmitNilData option for the model, which automatically filers the Data parameters for `nil` values.

​	OmitNilData 为模型设置 OptionOmitNilData 选项，该选项会自动为 `nil` 值提交 Data 参数。

#### (*Model) OmitNilWhere

```go
func (m *Model) OmitNilWhere() *Model
```

OmitNilWhere sets optionOmitNilWhere option for the model, which automatically filers the Where/Having parameters for `nil` values.

​	OmitNilWhere 为模型设置 optionOmitNilWhere 选项，该选项会自动为 `nil` 值归档 Where/having 参数。

#### (*Model) OnConflict

```go
func (m *Model) OnConflict(onConflict ...interface{}) *Model
```

OnConflict sets the primary key or index when columns conflicts occurs. It’s not necessary for MySQL driver.

​	OnConflict 在发生列冲突时设置主键或索引。MySQL驱动程序不是必需的。

#### (*Model) OnDuplicate

```go
func (m *Model) OnDuplicate(onDuplicate ...interface{}) *Model
```

OnDuplicate sets the operations when columns conflicts occurs. In MySQL, this is used for “ON DUPLICATE KEY UPDATE” statement. In PgSQL, this is used for “ON CONFLICT (id) DO UPDATE SET” statement. The parameter `onDuplicate` can be type of string/Raw/*Raw/map/slice. Example:

​	OnDuplicate 设置发生列冲突时的操作。在MySQL中，这用于“ON DUPLICATE KEY UPDATE”语句。在 PgSQL 中，这用于“ON CONFLICT （id） DO UPDATE SET”语句。参数 `onDuplicate` 类型可以是 string/Raw/*Raw/map/slice。例：

OnDuplicate(“nickname, age”) OnDuplicate(“nickname”, “age”)

​	OnDuplicate（“昵称，年龄”） OnDuplicate（“昵称”， “年龄”）

```
OnDuplicate(g.Map{
	  "nickname": gdb.Raw("CONCAT('name_', VALUES(`nickname`))"),
})

OnDuplicate(g.Map{
	  "nickname": "passport",
}).
```

#### (*Model) OnDuplicateEx

```go
func (m *Model) OnDuplicateEx(onDuplicateEx ...interface{}) *Model
```

OnDuplicateEx sets the excluding columns for operations when columns conflict occurs. In MySQL, this is used for “ON DUPLICATE KEY UPDATE” statement. In PgSQL, this is used for “ON CONFLICT (id) DO UPDATE SET” statement. The parameter `onDuplicateEx` can be type of string/map/slice. Example:

​	OnDuplicateEx 设置发生列冲突时操作的排除列。在MySQL中，这用于“ON DUPLICATE KEY UPDATE”语句。在 PgSQL 中，这用于“ON CONFLICT （id） DO UPDATE SET”语句。参数 `onDuplicateEx` 可以是字符串/映射/切片的类型。例：

OnDuplicateEx(“passport, password”) OnDuplicateEx(“passport”, “password”)

​	OnDuplicateEx（“护照，密码”） OnDuplicateEx（“护照”， “密码”）

```
OnDuplicateEx(g.Map{
	  "passport": "",
	  "password": "",
}).
```

#### (*Model) One

```go
func (m *Model) One(where ...interface{}) (Record, error)
```

One retrieves one record from table and returns the result as map type. It returns nil if there’s no record retrieved with the given conditions from table.

​	从表中检索一条记录，并将结果作为映射类型返回。如果没有从表中检索到具有给定条件的记录，则返回 nil。

The optional parameter `where` is the same as the parameter of Model.Where function, see Model.Where.

​	可选参数 `where` 与 Model.Where 函数的参数相同，参见 Model.Where。

#### (*Model) Order

```go
func (m *Model) Order(orderBy ...interface{}) *Model
```

Order sets the “ORDER BY” statement for the model.

​	Order 设置模型的“ORDER BY”语句。

Eg: Order(“id desc”) Order(“id”, “desc”). Order(“id desc,name asc”) Order(“id desc”).Order(“name asc”) Order(gdb.Raw(“field(id, 3,1,2)”)).

​	例如：Order（“id desc”） Order（“id”， “desc”）。Order（“id desc，name asc”） Order（“id desc”）.Order（“name asc”） Order（gdb.Raw（“field（id， 3,1,2）”））。

#### (*Model) OrderAsc

```go
func (m *Model) OrderAsc(column string) *Model
```

OrderAsc sets the “ORDER BY xxx ASC” statement for the model.

​	OrderAsc 为模型设置“ORDER BY xxx ASC”语句。

#### (*Model) OrderDesc

```go
func (m *Model) OrderDesc(column string) *Model
```

OrderDesc sets the “ORDER BY xxx DESC” statement for the model.

​	OrderDesc 为模型设置“ORDER BY xxx DESC”语句。

#### (*Model) OrderRandom

```go
func (m *Model) OrderRandom() *Model
```

OrderRandom sets the “ORDER BY RANDOM()” statement for the model.

​	OrderRandom 设置模型的“ORDER BY RANDOM（）”语句。

#### (*Model) Page

```go
func (m *Model) Page(page, limit int) *Model
```

Page sets the paging number for the model. The parameter `page` is started from 1 for paging. Note that, it differs that the Limit function starts from 0 for “LIMIT” statement.

​	Page 设置模型的分页编号。该参数 `page` 从 1 开始进行分页。请注意，对于“LIMIT”语句，Limit 函数从 0 开始是不同的。

#### (*Model) Partition

```go
func (m *Model) Partition(partitions ...string) *Model
```

Partition sets Partition name. Example: dao.User.Ctx(ctx).Partition（“p1”,“p2”,“p3”).All()

​	分区设置分区名称。示例：dao。User.Ctx（ctx）。分区（“p1”，“p2”，“p3”）。全部（）

#### (*Model) QuoteWord

```go
func (m *Model) QuoteWord(s string) string
```

QuoteWord checks given string `s` a word, if true it quotes `s` with security chars of the database and returns the quoted string; or else it returns `s` without any change.

​	QuoteWord 检查给定字符串 `s` 的单词，如果为 true，则 `s` 使用数据库的安全字符引用并返回带引号的字符串;否则，它将 `s` 返回而没有任何更改。

The meaning of a `word` can be considered as a column name.

​	a `word` 的含义可以视为列名。

#### (*Model) Raw

```go
func (m *Model) Raw(rawSql string, args ...interface{}) *Model
```

Raw sets current model as a raw sql model. Example:

​	raw 将当前模型设置为原始 sql 模型。例：

```
db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
```

See Core.Raw.

​	请参阅 Core.Raw。

#### (*Model) Replace

```go
func (m *Model) Replace(data ...interface{}) (result sql.Result, err error)
```

Replace does “REPLACE INTO …” statement for the model. The optional parameter `data` is the same as the parameter of Model.Data function, see Model.Data.

​	替换执行“REPLACE INTO ...”模型的语句。可选参数 `data` 与 Model.Data 函数的参数相同，请参见 Model.Data。

#### (*Model) RightJoin

```go
func (m *Model) RightJoin(tableOrSubQueryAndJoinConditions ...string) *Model
```

RightJoin does “RIGHT JOIN … ON …” statement on the model. The parameter `table` can be joined table and its joined condition, and also with its alias name.

​	RightJoin做“RIGHT JOIN ...在......”关于模型的声明。参数 `table` 可以是联接表及其联接条件，也可以是其别名。

Eg: Model(“user”).RightJoin(“user_detail”, “user_detail.uid=user.uid”) Model(“user”, “u”).RightJoin(“user_detail”, “ud”, “ud.uid=u.uid”) Model(“user”, “u”).RightJoin(“SELECT xxx FROM xxx”,“a”, “a.uid=u.uid”).

​	例如：model（“user”）。RightJoin（“user_detail”， “user_detail.uid=user.uid”） Model（“用户”， “u”）.RightJoin（“user_detail”， “ud”， “ud.uid=u.uid”） Model（“用户”， “u”）.RightJoin（“从 xxx 中选择 xxx”，“a”， “a.uid=u.uid”）。

#### (*Model) RightJoinOnField

```go
func (m *Model) RightJoinOnField(table, field string) *Model
```

RightJoinOnField performs as RightJoin, but it joins both tables with the `same field name`.

​	RightJoinOnField 以 RightJoin 的形式执行，但它使用 `same field name` .

Eg: Model(“order”).InnerJoinOnField(“user”, “user_id”) Model(“order”).InnerJoinOnField(“product”, “product_id”).

​	例如：model（“order”）。InnerJoinOnField（“用户”， “user_id”） Model（“order”）.InnerJoinOnField（“产品”， “product_id”）.

#### (*Model) RightJoinOnFields

```go
func (m *Model) RightJoinOnFields(table, firstField, operator, secondField string) *Model
```

RightJoinOnFields performs as RightJoin. It specifies different fields and comparison operator.

​	RightJoinOnFields 以 RightJoin 的形式执行。它指定不同的字段和比较运算符。

Eg: Model(“user”).RightJoinOnFields(“order”, “id”, “=”, “user_id”) Model(“user”).RightJoinOnFields(“order”, “id”, “>”, “user_id”) Model(“user”).RightJoinOnFields(“order”, “id”, “<”, “user_id”)

​	例如：model（“user”）。RightJoinOnFields（“order”， “id”， “=”， “user_id”） Model（“用户”）。RightJoinOnFields（“order”， “id”， “>”， “user_id”） Model（“用户”）。RightJoinOnFields（“订单”， “id”， “<”， “user_id”）

#### (*Model) Safe

```go
func (m *Model) Safe(safe ...bool) *Model
```

Safe marks this model safe or unsafe. If safe is true, it clones and returns a new model object whenever the operation done, or else it changes the attribute of current model.

​	“安全”标记此模型安全或不安全。如果 safe 为 true，则每当操作完成时，它就会克隆并返回一个新的模型对象，否则它会更改当前模型的属性。

#### (*Model) Save

```go
func (m *Model) Save(data ...interface{}) (result sql.Result, err error)
```

Save does “INSERT INTO … ON DUPLICATE KEY UPDATE…” statement for the model. The optional parameter `data` is the same as the parameter of Model.Data function, see Model.Data.

​	保存执行“插入...在重复的密钥更新中......”模型的语句。可选参数 `data` 与 Model.Data 函数的参数相同，请参见 Model.Data。

It updates the record if there’s primary or unique index in the saving data, or else it inserts a new record into the table.

​	如果保存数据中有主索引或唯一索引，它将更新记录，或者将新记录插入到表中。

#### (*Model) Scan

```go
func (m *Model) Scan(pointer interface{}, where ...interface{}) error
```

Scan automatically calls Struct or Structs function according to the type of parameter `pointer`. It calls function doStruct if `pointer` is type of *struct/**struct. It calls function doStructs if `pointer` is type of *[]struct/*[]*struct.

​	Scan 根据参数类型自动调用 Struct 或 Structs 函数 `pointer` 。如果 `pointer` 是 *struct/**struct 的类型，则调用函数 doStruct。如果 `pointer` type 为 []struct/[]*struct，则调用函数 doStructs。

The optional parameter `where` is the same as the parameter of Model.Where function, see Model.Where.

​	可选参数 `where` 与 Model.Where 函数的参数相同，参见 Model.Where。

Note that it returns sql.ErrNoRows if the given parameter `pointer` pointed to a variable that has default value and there’s no record retrieved with the given conditions from table.

​	请注意，它返回 sql。如果给定参数 `pointer` 指向具有默认值的变量，并且没有从表中检索到给定条件的记录，则为 ErrNoRows。

Example: user := new(User) err := db.Model(“user”).Where(“id”, 1).Scan(user)

​	示例：user ：= new（User） err ：= db。Model（“用户”）。其中（“id”， 1）。扫描（用户）

user := (*User)(nil) err := db.Model(“user”).Where(“id”, 1).Scan(&user)

​	用户 ：= （*User）（nil） err ：= db。Model（“用户”）。其中（“id”， 1）。扫描（&user）

users := ([]User)(nil) err := db.Model(“user”).Scan(&users)

​	用户 ：= （[]User）（nil） err ：= db.Model（“用户”）。扫描（&用户）

users := ([]*User)(nil) err := db.Model(“user”).Scan(&users).

​	用户 ：= （[]*User）（nil） err ：= db.Model（“用户”）。扫描（&用户）。

#### (*Model) ScanAndCount

```go
func (m *Model) ScanAndCount(pointer interface{}, totalCount *int, useFieldForCount bool) (err error)
```

ScanAndCount scans a single record or record array that matches the given conditions and counts the total number of records that match those conditions. If useFieldForCount is true, it will use the fields specified in the model for counting; The pointer parameter is a pointer to a struct that the scanned data will be stored in. The pointerCount parameter is a pointer to an integer that will be set to the total number of records that match the given conditions. The where parameter is an optional list of conditions to use when retrieving records.

​	ScanAndCount 扫描与给定条件匹配的单个记录或记录数组，并计算与这些条件匹配的记录总数。如果 useFieldForCount 为 true，它将使用模型中指定的字段进行计数;指针参数是指向将存储扫描数据的结构的指针。pointerCount 参数是指向整数的指针，该整数将设置为与给定条件匹配的记录总数。where 参数是检索记录时要使用的条件的可选列表。

Example:

​	例：

```go
var count int
user := new(User)
err  := db.Model("user").Where("id", 1).ScanAndCount(user,&count,true)
fmt.Println(user, count)
```

Example Join:

​	联接示例：

```go
type User struct {
	Id       int
	Passport string
	Name     string
	Age      int
}
var users []User
var count int
db.Model(table).As("u1").
	LeftJoin(tableName2, "u2", "u2.id=u1.id").
	Fields("u1.passport,u1.id,u2.name,u2.age").
	Where("u1.id<2").
	ScanAndCount(&users, &count, false)
```

#### (*Model) ScanList

```go
func (m *Model) ScanList(structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error)
```

ScanList converts `r` to struct slice which contains other complex struct attributes. Note that the parameter `listPointer` should be type of *[]struct/*[]*struct.

​	ScanList `r` 转换为包含其他复杂结构属性的结构切片。请注意，参数 `listPointer` 的类型应为 []struct/[]*struct。

See Result.ScanList.

​	请参阅 Result.ScanList。

#### (*Model) Schema

```go
func (m *Model) Schema(schema string) *Model
```

Schema sets the schema for current operation.

​	架构设置当前操作的架构。

#### (*Model) Slave

```go
func (m *Model) Slave() *Model
```

Slave marks the following operation on slave node. Note that it makes sense only if there’s any slave node configured.

​	Slave 在从节点上标记以下操作。请注意，仅当配置了任何从属节点时，它才有意义。

#### (*Model) SoftTime

```go
func (m *Model) SoftTime(option SoftTimeOption) *Model
```

SoftTime sets the SoftTimeOption to customize soft time feature for Model.

​	SoftTime 设置 SoftTimeOption 以自定义模型的软时间功能。

#### (*Model) Sum

```go
func (m *Model) Sum(column string) (float64, error)
```

Sum does “SELECT SUM(x) FROM …” statement for the model.

​	Sum 表示 “SELECT SUM（x） FROM ...”模型的语句。

#### (*Model) TX

```go
func (m *Model) TX(tx TX) *Model
```

TX sets/changes the transaction for current operation.

​	TX 设置/更改当前操作的事务。

#### (*Model) TableFields

```go
func (m *Model) TableFields(tableStr string, schema ...string) (fields map[string]*TableField, err error)
```

TableFields retrieves and returns the fields’ information of specified table of current schema.

​	TableFields 检索并返回当前架构的指定表的字段信息。

Also see DriverMysql.TableFields.

​	另请参阅 DriverMysql.TableFields。

#### (*Model) Transaction

```go
func (m *Model) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error)
```

Transaction wraps the transaction logic using function `f`. It rollbacks the transaction and returns the error from function `f` if it returns non-nil error. It commits the transaction and returns nil if function `f` returns nil.

​	事务使用函数 `f` 包装事务逻辑。 `f` 它会回滚事务，如果函数返回非 nil 错误，则返回错误。它提交事务，如果函数 `f` 返回 nil，则返回 nil。

Note that, you should not Commit or Rollback the transaction in function `f` as it is automatically handled by this function.

​	请注意，您不应该在函数 `f` 中提交或回滚事务，因为它是由此函数自动处理的。

#### (*Model) Union

```go
func (m *Model) Union(unions ...*Model) *Model
```

Union does “(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) …” statement for the model.

​	Union does “（SELECT xxx FROM xxx） UNION （SELECT xxx FROM xxx） ...”模型的语句。

#### (*Model) UnionAll

```go
func (m *Model) UnionAll(unions ...*Model) *Model
```

UnionAll does “(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) …” statement for the model.

​	UnionAll 执行“（从 xxx 中选择 xxx）UNION ALL （SELECT XXX FROM XXX）...”模型的语句。

#### (*Model) Unscoped

```go
func (m *Model) Unscoped() *Model
```

Unscoped disables the soft time feature for insert, update and delete operations.

​	Unscoped 禁用插入、更新和删除操作的软时间功能。

#### (*Model) Update

```go
func (m *Model) Update(dataAndWhere ...interface{}) (result sql.Result, err error)
```

Update does “UPDATE … " statement for the model.

​	更新执行“更新...“的语句。

If the optional parameter `dataAndWhere` is given, the dataAndWhere[0] is the updated data field, and dataAndWhere[1:] is treated as where condition fields. Also see Model.Data and Model.Where functions.

​	如果给定了可选参数 `dataAndWhere` ，则 dataAndWhere[0] 是更新的数据字段，dataAndWhere[1：] 被视为 where 条件字段。另请参阅 Model.Data 和 Model.Where 函数。

#### (*Model) UpdateAndGetAffected

```go
func (m *Model) UpdateAndGetAffected(dataAndWhere ...interface{}) (affected int64, err error)
```

UpdateAndGetAffected performs update statement and returns the affected rows number.

​	UpdateAndGetAffected 执行 update 语句并返回受影响的行号。

#### (*Model) Value

```go
func (m *Model) Value(fieldsAndWhere ...interface{}) (Value, error)
```

Value retrieves a specified record value from table and returns the result as interface type. It returns nil if there’s no record found with the given conditions from table.

​	Value 从表中检索指定的记录值，并将结果作为接口类型返回。如果在表中找不到具有给定条件的记录，则返回 nil。

If the optional parameter `fieldsAndWhere` is given, the fieldsAndWhere[0] is the selected fields and fieldsAndWhere[1:] is treated as where condition fields. Also see Model.Fields and Model.Where functions.

​	如果给定了可选参数 `fieldsAndWhere` ，则 fieldsAndWhere[0] 是所选字段，fieldsAndWhere[1：] 被视为 where 条件字段。另请参阅 Model.Fields 和 Model.Where 函数。

#### (*Model) Where

```go
func (m *Model) Where(where interface{}, args ...interface{}) *Model
```

Where sets the condition statement for the builder. The parameter `where` can be type of string/map/gmap/slice/struct/*struct, etc. Note that, if it’s called more than one times, multiple conditions will be joined into where statement using “AND”. See WhereBuilder.Where.

​	Where 设置构建器的条件语句。参数 `where` 可以是字符串/地图/gmap/slice/struct/*struct等类型。请注意，如果它被调用了多次，则多个条件将使用“AND”加入到 where 语句中。请参阅 WhereBuilder.Where。

#### (*Model) WhereBetween

```go
func (m *Model) WhereBetween(column string, min, max interface{}) *Model
```

WhereBetween builds `column BETWEEN min AND max` statement. See WhereBuilder.WhereBetween.

​	WhereBetween 构建 `column BETWEEN min AND max` 语句。请参阅 WhereBuilder.WhereBetween。

#### (*Model) WhereGT

```go
func (m *Model) WhereGT(column string, value interface{}) *Model
```

WhereGT builds `column > value` statement. See WhereBuilder.WhereGT.

​	WhereGT 构建 `column > value` 语句。请参阅 WhereBuilder.WhereGT。

#### (*Model) WhereGTE

```go
func (m *Model) WhereGTE(column string, value interface{}) *Model
```

WhereGTE builds `column >= value` statement. See WhereBuilder.WhereGTE.

​	其中 GTE 构建 `column >= value` 语句。请参阅 WhereBuilder.WhereGTE。

#### (*Model) WhereIn

```go
func (m *Model) WhereIn(column string, in interface{}) *Model
```

WhereIn builds `column IN (in)` statement. See WhereBuilder.WhereIn.

​	其中 builds `column IN (in)` 语句。请参阅 WhereBuilder.WhereIn。

#### (*Model) WhereLT

```go
func (m *Model) WhereLT(column string, value interface{}) *Model
```

WhereLT builds `column < value` statement. See WhereBuilder.WhereLT.

​	其中 LT 构建 `column < value` 语句。请参阅 WhereBuilder.WhereLT。

#### (*Model) WhereLTE

```go
func (m *Model) WhereLTE(column string, value interface{}) *Model
```

WhereLTE builds `column <= value` statement. See WhereBuilder.WhereLTE.

​	其中LTE构建 `column <= value` 语句。请参阅 WhereBuilder.WhereLTE。

#### (*Model) WhereLike

```go
func (m *Model) WhereLike(column string, like string) *Model
```

WhereLike builds `column LIKE like` statement. See WhereBuilder.WhereLike.

​	WhereLike 构建 `column LIKE like` 语句。请参阅 WhereBuilder.WhereLike。

#### (*Model) WhereNot

```go
func (m *Model) WhereNot(column string, value interface{}) *Model
```

WhereNot builds `column != value` statement. See WhereBuilder.WhereNot.

​	WhereNot 构建 `column != value` 语句。请参阅 WhereBuilder.WhereNot。

#### (*Model) WhereNotBetween

```go
func (m *Model) WhereNotBetween(column string, min, max interface{}) *Model
```

WhereNotBetween builds `column NOT BETWEEN min AND max` statement. See WhereBuilder.WhereNotBetween.

​	WhereNotBetween 构建 `column NOT BETWEEN min AND max` 语句。请参阅 WhereBuilder.WhereNotBetween。

#### (*Model) WhereNotIn

```go
func (m *Model) WhereNotIn(column string, in interface{}) *Model
```

WhereNotIn builds `column NOT IN (in)` statement. See WhereBuilder.WhereNotIn.

​	WhereNotIn 构建 `column NOT IN (in)` 语句。请参阅 WhereBuilder.WhereNotIn。

#### (*Model) WhereNotLike

```go
func (m *Model) WhereNotLike(column string, like interface{}) *Model
```

WhereNotLike builds `column NOT LIKE like` statement. See WhereBuilder.WhereNotLike.

​	WhereNotLike 构建 `column NOT LIKE like` 语句。请参阅 WhereBuilder.WhereNotLike。

#### (*Model) WhereNotNull

```go
func (m *Model) WhereNotNull(columns ...string) *Model
```

WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement. See WhereBuilder.WhereNotNull.

​	WhereNotNull 生成 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。请参阅 WhereBuilder.WhereNotNull。

#### (*Model) WhereNull

```go
func (m *Model) WhereNull(columns ...string) *Model
```

WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement. See WhereBuilder.WhereNull.

​	WhereNull 构建 `columns[0] IS NULL AND columns[1] IS NULL ...` 语句。请参阅 WhereBuilder.WhereNull。

#### (*Model) WhereOr

```go
func (m *Model) WhereOr(where interface{}, args ...interface{}) *Model
```

WhereOr adds “OR” condition to the where statement. See WhereBuilder.WhereOr.

​	WhereOr 将“OR”条件添加到 where 语句中。请参阅 WhereBuilder.WhereOr。

#### (*Model) WhereOrBetween

```go
func (m *Model) WhereOrBetween(column string, min, max interface{}) *Model
```

WhereOrBetween builds `column BETWEEN min AND max` statement in `OR` conditions. See WhereBuilder.WhereOrBetween.

​	WhereOrBetween 在 conditions 中 `OR` 构建 `column BETWEEN min AND max` 语句。请参阅 WhereBuilder.WhereOrBetween。

#### (*Model) WhereOrGT

```go
func (m *Model) WhereOrGT(column string, value interface{}) *Model
```

WhereOrGT builds `column > value` statement in `OR` conditions. See WhereBuilder.WhereOrGT.

​	WhereOrGT 在 conditions 中 `OR` 构建 `column > value` 语句。请参阅 WhereBuilder.WhereOrGT。

#### (*Model) WhereOrGTE

```go
func (m *Model) WhereOrGTE(column string, value interface{}) *Model
```

WhereOrGTE builds `column >= value` statement in `OR` conditions. See WhereBuilder.WhereOrGTE.

​	WhereOrGTE 在 conditions 中 `OR` 构建 `column >= value` 语句。请参阅 WhereBuilder.WhereOrGTE。

#### (*Model) WhereOrIn

```go
func (m *Model) WhereOrIn(column string, in interface{}) *Model
```

WhereOrIn builds `column IN (in)` statement in `OR` conditions. See WhereBuilder.WhereOrIn.

​	WhereOrIn 在 conditions 中 `OR` 构建 `column IN (in)` 语句。请参阅 WhereBuilder.WhereOrIn。

#### (*Model) WhereOrLT

```go
func (m *Model) WhereOrLT(column string, value interface{}) *Model
```

WhereOrLT builds `column < value` statement in `OR` conditions. See WhereBuilder.WhereOrLT.

​	WhereOrLT 在 conditions 中 `OR` 构建 `column < value` 语句。请参阅 WhereBuilder.WhereOrLT。

#### (*Model) WhereOrLTE

```go
func (m *Model) WhereOrLTE(column string, value interface{}) *Model
```

WhereOrLTE builds `column <= value` statement in `OR` conditions. See WhereBuilder.WhereOrLTE.

​	WhereOrLTE 在 conditions 中 `OR` 构建 `column <= value` 语句。请参阅 WhereBuilder.WhereOrLTE。

#### (*Model) WhereOrLike

```go
func (m *Model) WhereOrLike(column string, like interface{}) *Model
```

WhereOrLike builds `column LIKE like` statement in `OR` conditions. See WhereBuilder.WhereOrLike.

​	WhereOrLike 在 conditions 中 `OR` 构建 `column LIKE like` 语句。请参阅 WhereBuilder.WhereOrLike。

#### (*Model) WhereOrNot

```go
func (m *Model) WhereOrNot(column string, value interface{}) *Model
```

WhereOrNot builds `column != value` statement. See WhereBuilder.WhereOrNot.

​	WhereOrNot 构建 `column != value` 语句。请参阅 WhereBuilder.WhereOrNot。

#### (*Model) WhereOrNotBetween

```go
func (m *Model) WhereOrNotBetween(column string, min, max interface{}) *Model
```

WhereOrNotBetween builds `column NOT BETWEEN min AND max` statement in `OR` conditions. See WhereBuilder.WhereOrNotBetween.

​	WhereOrNotBetween 在 conditions 中 `OR` 构建 `column NOT BETWEEN min AND max` 语句。请参阅 WhereBuilder.WhereOrNotBetween。

#### (*Model) WhereOrNotIn

```go
func (m *Model) WhereOrNotIn(column string, in interface{}) *Model
```

WhereOrNotIn builds `column NOT IN (in)` statement. See WhereBuilder.WhereOrNotIn.

​	WhereOrNotIn 生成 `column NOT IN (in)` 语句。请参阅 WhereBuilder.WhereOrNotIn。

#### (*Model) WhereOrNotLike

```go
func (m *Model) WhereOrNotLike(column string, like interface{}) *Model
```

WhereOrNotLike builds `column NOT LIKE 'like'` statement in `OR` conditions. See WhereBuilder.WhereOrNotLike.

​	WhereOrNotLike 在 conditions 中 `OR` 构建 `column NOT LIKE 'like'` 语句。请参阅 WhereBuilder.WhereOrNotLike。

#### (*Model) WhereOrNotNull

```go
func (m *Model) WhereOrNotNull(columns ...string) *Model
```

WhereOrNotNull builds `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` statement in `OR` conditions. See WhereBuilder.WhereOrNotNull.

​	WhereOrNotNull 在 conditions 中 `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` `OR` 生成语句。请参阅 WhereBuilder.WhereOrNotNull。

#### (*Model) WhereOrNull

```go
func (m *Model) WhereOrNull(columns ...string) *Model
```

WhereOrNull builds `columns[0] IS NULL OR columns[1] IS NULL ...` statement in `OR` conditions. See WhereBuilder.WhereOrNull.

​	WhereOrNull 在 conditions 中 `OR` 构建 `columns[0] IS NULL OR columns[1] IS NULL ...` 语句。请参阅 WhereBuilder.WhereOrNull。

#### (*Model) WhereOrPrefix

```go
func (m *Model) WhereOrPrefix(prefix string, where interface{}, args ...interface{}) *Model
```

WhereOrPrefix performs as WhereOr, but it adds prefix to each field in where statement. See WhereBuilder.WhereOrPrefix.

​	WhereOrPrefix 的执行方式为 WhereOr，但它向 where 语句中的每个字段添加前缀。请参阅 WhereBuilder.WhereOrPrefix。

#### (*Model) WhereOrPrefixBetween

```go
func (m *Model) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *Model
```

WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixBetween.

​	WhereOrPrefixBetween 在 conditions 中 `OR` 构建 `prefix.column BETWEEN min AND max` 语句。请参阅 WhereBuilder.WhereOrPrefixBetween。

#### (*Model) WhereOrPrefixGT

```go
func (m *Model) WhereOrPrefixGT(prefix string, column string, value interface{}) *Model
```

WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixGT.

​	WhereOrPrefixGT 在 conditions 中 `prefix.column > value` `OR` 构建语句。请参阅 WhereBuilder.WhereOrPrefixGT。

#### (*Model) WhereOrPrefixGTE

```go
func (m *Model) WhereOrPrefixGTE(prefix string, column string, value interface{}) *Model
```

WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixGTE.

​	WhereOrPrefixGTE 在 conditions 中 `OR` 构建 `prefix.column >= value` 语句。请参阅 WhereBuilder.WhereOrPrefixGTE。

#### (*Model) WhereOrPrefixIn

```go
func (m *Model) WhereOrPrefixIn(prefix string, column string, in interface{}) *Model
```

WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixIn.

​	WhereOrPrefixIn 在 conditions 中 `OR` 生成 `prefix.column IN (in)` 语句。请参阅 WhereBuilder.WhereOrPrefixIn。

#### (*Model) WhereOrPrefixLT

```go
func (m *Model) WhereOrPrefixLT(prefix string, column string, value interface{}) *Model
```

WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixLT.

​	WhereOrPrefixLT 在 conditions 中 `OR` 构建 `prefix.column < value` 语句。请参阅 WhereBuilder.WhereOrPrefixLT。

#### (*Model) WhereOrPrefixLTE

```go
func (m *Model) WhereOrPrefixLTE(prefix string, column string, value interface{}) *Model
```

WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixLTE.

​	WhereOrPrefixLTE 在 conditions 中 `OR` 构建 `prefix.column <= value` 语句。请参阅 WhereBuilder.WhereOrPrefixLTE。

#### (*Model) WhereOrPrefixLike

```go
func (m *Model) WhereOrPrefixLike(prefix string, column string, like interface{}) *Model
```

WhereOrPrefixLike builds `prefix.column LIKE like` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixLike.

​	WhereOrPrefixLike 在 conditions 中 `OR` 构建 `prefix.column LIKE like` 语句。请参阅 WhereBuilder.WhereOrPrefixLike。

#### (*Model) WhereOrPrefixNot

```go
func (m *Model) WhereOrPrefixNot(prefix string, column string, value interface{}) *Model
```

WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixNot.

​	WhereOrPrefixNot 在 conditions 中 `OR` 构建 `prefix.column != value` 语句。请参阅 WhereBuilder.WhereOrPrefixNot。

#### (*Model) WhereOrPrefixNotBetween

```go
func (m *Model) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *Model
```

WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixNotBetween.

​	WhereOrPrefixNotBetween 在 conditions 中 `OR` 构建 `prefix.column NOT BETWEEN min AND max` 语句。请参阅 WhereBuilder.WhereOrPrefixNotBetween。

#### (*Model) WhereOrPrefixNotIn

```go
func (m *Model) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *Model
```

WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement. See WhereBuilder.WhereOrPrefixNotIn.

​	WhereOrPrefixNotIn 生成 `prefix.column NOT IN (in)` 语句。请参阅 WhereBuilder.WhereOrPrefixNotIn。

#### (*Model) WhereOrPrefixNotLike

```go
func (m *Model) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *Model
```

WhereOrPrefixNotLike builds `prefix.column NOT LIKE like` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixNotLike.

​	WhereOrPrefixNotLike 在 conditions 中 `OR` 构建 `prefix.column NOT LIKE like` 语句。请参阅 WhereBuilder.WhereOrPrefixNotLike。

#### (*Model) WhereOrPrefixNotNull

```go
func (m *Model) WhereOrPrefixNotNull(prefix string, columns ...string) *Model
```

WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixNotNull.

​	WhereOrPrefixNotNull 在 conditions 中 `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` `OR` 构建语句。请参阅 WhereBuilder.WhereOrPrefixNotNull。

#### (*Model) WhereOrPrefixNull

```go
func (m *Model) WhereOrPrefixNull(prefix string, columns ...string) *Model
```

WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixNull.

​	WhereOrPrefixNull 在 conditions 中 `OR` 构建 `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` 语句。请参阅 WhereBuilder.WhereOrPrefixNull。

#### (*Model) WhereOrf

```go
func (m *Model) WhereOrf(format string, args ...interface{}) *Model
```

WhereOrf builds `OR` condition string using fmt.Sprintf and arguments. See WhereBuilder.WhereOrf.

​	WhereOrf 使用 fmt `OR` 构建条件字符串。Sprintf 和参数。请参阅 WhereBuilder.WhereOrf。

#### (*Model) WherePrefix

```go
func (m *Model) WherePrefix(prefix string, where interface{}, args ...interface{}) *Model
```

WherePrefix performs as Where, but it adds prefix to each field in where statement. See WhereBuilder.WherePrefix.

​	WherePrefix 的执行方式为 Where，但它向 where 语句中的每个字段添加前缀。请参阅 WhereBuilder.WherePrefix。

#### (*Model) WherePrefixBetween

```go
func (m *Model) WherePrefixBetween(prefix string, column string, min, max interface{}) *Model
```

WherePrefixBetween builds `prefix.column BETWEEN min AND max` statement. See WhereBuilder.WherePrefixBetween.

​	WherePrefixBetween 构建 `prefix.column BETWEEN min AND max` 语句。请参阅 WhereBuilder.WherePrefixBetween。

#### (*Model) WherePrefixGT

```go
func (m *Model) WherePrefixGT(prefix string, column string, value interface{}) *Model
```

WherePrefixGT builds `prefix.column > value` statement. See WhereBuilder.WherePrefixGT.

​	WherePrefixGT 构建 `prefix.column > value` 语句。请参阅 WhereBuilder.WherePrefixGT。

#### (*Model) WherePrefixGTE

```go
func (m *Model) WherePrefixGTE(prefix string, column string, value interface{}) *Model
```

WherePrefixGTE builds `prefix.column >= value` statement. See WhereBuilder.WherePrefixGTE.

​	WherePrefixGTE 生成 `prefix.column >= value` 语句。请参阅 WhereBuilder.WherePrefixGTE。

#### (*Model) WherePrefixIn

```go
func (m *Model) WherePrefixIn(prefix string, column string, in interface{}) *Model
```

WherePrefixIn builds `prefix.column IN (in)` statement. See WhereBuilder.WherePrefixIn.

​	WherePrefixIn 构建 `prefix.column IN (in)` 语句。请参阅 WhereBuilder.WherePrefixIn。

#### (*Model) WherePrefixLT

```go
func (m *Model) WherePrefixLT(prefix string, column string, value interface{}) *Model
```

WherePrefixLT builds `prefix.column < value` statement. See WhereBuilder.WherePrefixLT.

​	WherePrefixLT 生成 `prefix.column < value` 语句。请参阅 WhereBuilder.WherePrefixLT。

#### (*Model) WherePrefixLTE

```go
func (m *Model) WherePrefixLTE(prefix string, column string, value interface{}) *Model
```

WherePrefixLTE builds `prefix.column <= value` statement. See WhereBuilder.WherePrefixLTE.

​	WherePrefixLTE 构建 `prefix.column <= value` 语句。请参阅 WhereBuilder.WherePrefixLTE。

#### (*Model) WherePrefixLike

```go
func (m *Model) WherePrefixLike(prefix string, column string, like interface{}) *Model
```

WherePrefixLike builds `prefix.column LIKE like` statement. See WhereBuilder.WherePrefixLike.

​	WherePrefixLike 构建语 `prefix.column LIKE like` 句。请参阅 WhereBuilder.WherePrefixLike。

#### (*Model) WherePrefixNot

```go
func (m *Model) WherePrefixNot(prefix string, column string, value interface{}) *Model
```

WherePrefixNot builds `prefix.column != value` statement. See WhereBuilder.WherePrefixNot.

​	WherePrefixNot 构建 `prefix.column != value` 语句。请参阅 WhereBuilder.WherePrefixNot。

#### (*Model) WherePrefixNotBetween

```go
func (m *Model) WherePrefixNotBetween(prefix string, column string, min, max interface{}) *Model
```

WherePrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement. See WhereBuilder.WherePrefixNotBetween.

​	WherePrefixNotBetween 构建语 `prefix.column NOT BETWEEN min AND max` 句。请参阅 WhereBuilder.WherePrefixNotBetween。

#### (*Model) WherePrefixNotIn

```go
func (m *Model) WherePrefixNotIn(prefix string, column string, in interface{}) *Model
```

WherePrefixNotIn builds `prefix.column NOT IN (in)` statement. See WhereBuilder.WherePrefixNotIn.

​	WherePrefixNotIn 构建 `prefix.column NOT IN (in)` 语句。请参阅 WhereBuilder.WherePrefixNotIn。

#### (*Model) WherePrefixNotLike

```go
func (m *Model) WherePrefixNotLike(prefix string, column string, like interface{}) *Model
```

WherePrefixNotLike builds `prefix.column NOT LIKE like` statement. See WhereBuilder.WherePrefixNotLike.

​	WherePrefixNotLike 构建语 `prefix.column NOT LIKE like` 句。请参阅 WhereBuilder.WherePrefixNotLike。

#### (*Model) WherePrefixNotNull

```go
func (m *Model) WherePrefixNotNull(prefix string, columns ...string) *Model
```

WherePrefixNotNull builds `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` statement. See WhereBuilder.WherePrefixNotNull.

​	WherePrefixNotNull 生成 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 语句。请参阅 WhereBuilder.WherePrefixNotNull。

#### (*Model) WherePrefixNull

```go
func (m *Model) WherePrefixNull(prefix string, columns ...string) *Model
```

WherePrefixNull builds `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` statement. See WhereBuilder.WherePrefixNull.

​	WherePrefixNull 构建语 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 句。请参阅 WhereBuilder.WherePrefixNull。

#### (*Model) WherePri

```go
func (m *Model) WherePri(where interface{}, args ...interface{}) *Model
```

WherePri does the same logic as Model.Where except that if the parameter `where` is a single condition like int/string/float/slice, it treats the condition as the primary key value. That is, if primary key is “id” and given `where` parameter as “123”, the WherePri function treats the condition as “id=123”, but Model.Where treats the condition as string “123”. See WhereBuilder.WherePri.

​	WherePri 执行与 Model.Where 相同的逻辑，不同之处在于如果参数 `where` 是单个条件（如 int/string/float/slice），则它将该条件视为主键值。也就是说，如果主键为“id”，并且给定 `where` 参数为“123”，则 WherePri 函数将条件视为“id=123”，但 Model.Where 将条件视为字符串“123”。请参阅 WhereBuilder.WherePri。

#### (*Model) Wheref

```go
func (m *Model) Wheref(format string, args ...interface{}) *Model
```

Wheref builds condition string using fmt.Sprintf and arguments. Note that if the number of `args` is more than the placeholder in `format`, the extra `args` will be used as the where condition arguments of the Model. See WhereBuilder.Wheref.

​	Wheref 使用 fmt 构建条件字符串。Sprintf 和参数。请注意，如果 的数量 `args` 大于 `format` 中的占位符，则额外的 `args` 将用作模型的 where 条件参数。请参阅 WhereBuilder.Wheref。

#### (*Model) With

```go
func (m *Model) With(objects ...interface{}) *Model
```

With creates and returns an ORM model based on metadata of given object. It also enables model association operations feature on given `object`. It can be called multiple times to add one or more objects to model and enable their mode association operations feature. For example, if given struct definition:

​	With 创建并返回一个基于给定对象的元数据的 ORM 模型。它还在给定 `object` 的 .可以多次调用它来添加一个或多个对象进行建模并启用其模式关联操作功能。例如，如果给定结构定义：

```go
type User struct {
	 gmeta.Meta `orm:"table:user"`
	 Id         int           `json:"id"`
	 Name       string        `json:"name"`
	 UserDetail *UserDetail   `orm:"with:uid=id"`
	 UserScores []*UserScores `orm:"with:uid=id"`
}
```

We can enable model association operations on attribute `UserDetail` and `UserScores` by:

​	我们可以通过以下方式对属性 `UserDetail` 和 `UserScores` 属性启用模型关联操作：

```
db.With(User{}.UserDetail).With(User{}.UserScores).Scan(xxx)
```

Or:

```
db.With(UserDetail{}).With(UserScores{}).Scan(xxx)
```

Or:

```
db.With(UserDetail{}, UserScores{}).Scan(xxx)
```

#### (*Model) WithAll

```go
func (m *Model) WithAll() *Model
```

WithAll enables model association operations on all objects that have “with” tag in the struct.

​	WithAll 对结构中具有“with”标记的所有对象启用模型关联操作。

### type ModelHandler

```go
type ModelHandler func(m *Model) *Model
```

ModelHandler is a function that handles given Model and returns a new Model that is custom modified.

​	ModelHandler 是一个函数，用于处理给定的模型并返回自定义修改的新模型。

### type Raw

```go
type Raw string // Raw is a raw sql that will not be treated as argument but as a direct sql part.
```

### type Record

```go
type Record map[string]Value // Record is the row record of the table.
```

#### (Record) GMap

```go
func (r Record) GMap() *gmap.StrAnyMap
```

GMap converts `r` to a gmap.

​	GMap `r` 转换为 gmap。

#### (Record) IsEmpty

```go
func (r Record) IsEmpty() bool
```

IsEmpty checks and returns whether `r` is empty.

​	IsEmpty 检查并返回是否 `r` 为空。

#### (Record) Json

```go
func (r Record) Json() string
```

Json converts `r` to JSON format content.

​	Json `r` 转换为 JSON 格式的内容。

#### (Record) Map

```go
func (r Record) Map() Map
```

Map converts `r` to map[string]interface{}.

​	Map `r` 转换为 map[string]interface{}。

#### (Record) Struct

```go
func (r Record) Struct(pointer interface{}) error
```

Struct converts `r` to a struct. Note that the parameter `pointer` should be type of *struct/**struct.

​	Struct 转换为 `r` struct。请注意，参数 `pointer` 的类型应为 *struct/**struct。

Note that it returns sql.ErrNoRows if `r` is empty.

​	请注意，它返回 sql。如果 `r` ErrNoRows 为空。

#### (Record) Xml

```go
func (r Record) Xml(rootTag ...string) string
```

Xml converts `r` to XML format content.

​	Xml `r` 转换为 XML 格式的内容。

### type Result

```go
type Result []Record // Result is the row record array.
```

#### (Result) Array

```go
func (r Result) Array(field ...string) []Value
```

Array retrieves and returns specified column values as slice. The parameter `field` is optional is the column field is only one. The default `field` is the first field name of the first item in `Result` if parameter `field` is not given.

​	数组检索指定的列值并将其作为切片返回。该参数 `field` 是可选的，因为列字段只有一个。默认 `field` 值是 if 参数 `field` 中 `Result` 第一项的第一个字段名称。

#### (Result) Chunk

```go
func (r Result) Chunk(size int) []Result
```

Chunk splits a Result into multiple Results, the size of each array is determined by `size`. The last chunk may contain less than size elements.

​	Chunk 将一个 Result 拆分为多个 Results，每个数组的大小由 `size` 决定。最后一个块可能包含小于 size 的元素。

#### (Result) IsEmpty

```go
func (r Result) IsEmpty() bool
```

IsEmpty checks and returns whether `r` is empty.

​	IsEmpty 检查并返回是否 `r` 为空。

#### (Result) Json

```go
func (r Result) Json() string
```

Json converts `r` to JSON format content.

​	Json `r` 转换为 JSON 格式的内容。

#### (Result) Len

```go
func (r Result) Len() int
```

Len returns the length of result list.

​	Len 返回结果列表的长度。

#### (Result) List

```go
func (r Result) List() List
```

List converts `r` to a List.

​	List 将 `r` 转换为 List。

#### (Result) MapKeyInt

```go
func (r Result) MapKeyInt(key string) map[int]Map
```

MapKeyInt converts `r` to a map[int]Map of which key is specified by `key`.

​	MapKeyInt `r` 转换为 map[int]Map 的键由 `key` 指定。

#### (Result) MapKeyStr

```go
func (r Result) MapKeyStr(key string) map[string]Map
```

MapKeyStr converts `r` to a map[string]Map of which key is specified by `key`.

​	MapKeyStr `r` 转换为 map[string]Map 的键由 `key` .

#### (Result) MapKeyUint

```go
func (r Result) MapKeyUint(key string) map[uint]Map
```

MapKeyUint converts `r` to a map[uint]Map of which key is specified by `key`.

​	MapKeyUint `r` 转换为 map[uint]Map 的键由 `key` 指定。

#### (Result) MapKeyValue

```go
func (r Result) MapKeyValue(key string) map[string]Value
```

MapKeyValue converts `r` to a map[string]Value of which key is specified by `key`. Note that the item value may be type of slice.

​	MapKeyValue `r` 转换为 map[string]其键的值由 `key` .请注意，项目值可以是切片类型。

#### (Result) RecordKeyInt

```go
func (r Result) RecordKeyInt(key string) map[int]Record
```

RecordKeyInt converts `r` to a map[int]Record of which key is specified by `key`.

​	RecordKeyInt `r` 转换为 map[int]由 `key` 指定其键的记录。

#### (Result) RecordKeyStr

```go
func (r Result) RecordKeyStr(key string) map[string]Record
```

RecordKeyStr converts `r` to a map[string]Record of which key is specified by `key`.

​	RecordKeyStr `r` 转换为 map[string]由 `key` 指定键的记录。

#### (Result) RecordKeyUint

```go
func (r Result) RecordKeyUint(key string) map[uint]Record
```

RecordKeyUint converts `r` to a map[uint]Record of which key is specified by `key`.

​	RecordKeyUint `r` 转换为 map[uint]由 指定的键 `key` 的记录。

#### (Result) ScanList

```go
func (r Result) ScanList(structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error)
```

ScanList converts `r` to struct slice which contains other complex struct attributes. Note that the parameter `structSlicePointer` should be type of *[]struct/*[]*struct.

​	ScanList `r` 转换为包含其他复杂结构属性的结构切片。请注意，参数 `structSlicePointer` 的类型应为 []struct/[]*struct。

Usage example 1: Normal attribute struct relation:

​	使用示例1：法线属性结构关系：

```go
type EntityUser struct {
	   Uid  int
	   Name string
}

type EntityUserDetail struct {
	   Uid     int
	   Address string
}

type EntityUserScores struct {
	   Id     int
	   Uid    int
	   Score  int
	   Course string
}

type Entity struct {
    User       *EntityUser
	   UserDetail *EntityUserDetail
	   UserScores []*EntityUserScores
}
```

var users []*Entity ScanList(&users, “User”) ScanList(&users, “User”, “uid”) ScanList(&users, “UserDetail”, “User”, “uid:Uid”) ScanList(&users, “UserScores”, “User”, “uid:Uid”) ScanList(&users, “UserScores”, “User”, “uid”)

​	var users []*实体 ScanList（&users， “用户”） ScanList（&users， “用户”， “uid”） ScanList（&users， “UserDetail”， “用户”， “uid：Uid”） ScanList（&users， “UserScores”， “User”， “uid：Uid”） ScanList（&users， “UserScores”， “User”， “uid”）

Usage example 2: Embedded attribute struct relation:

​	使用示例2：嵌入属性结构关系：

```go
type EntityUser struct {
	   Uid  int
	   Name string
}

type EntityUserDetail struct {
	   Uid     int
	   Address string
}

type EntityUserScores struct {
	   Id    int
	   Uid   int
	   Score int
}

type Entity struct {
	   EntityUser
	   UserDetail EntityUserDetail
	   UserScores []EntityUserScores
}
```

var users []*Entity ScanList(&users) ScanList(&users, “UserDetail”, “uid”) ScanList(&users, “UserScores”, “uid”)

​	var users []*实体 ScanList（&users） ScanList（&users， “UserDetail”， “uid”） ScanList（&users， “UserScores”， “uid”）

The parameters “User/UserDetail/UserScores” in the example codes specify the target attribute struct that current result will be bound to.

​	示例代码中的参数“User/UserDetail/UserScores”指定当前结果将绑定到的目标属性结构。

The “uid” in the example codes is the table field name of the result, and the “Uid” is the relational struct attribute name - not the attribute name of the bound to target. In the example codes, it’s attribute name “Uid” of “User” of entity “Entity”. It automatically calculates the HasOne/HasMany relationship with given `relation` parameter.

​	示例代码中的“uid”是结果的表字段名称，“Uid”是关系结构属性名称，而不是绑定到目标的属性名称。在示例代码中，它是实体“Entity”的“User”的属性名称“Uid”。它自动计算给定 `relation` 参数的 HasOne/HasMany 关系。

See the example or unit testing cases for clear understanding for this function.

​	请参阅示例或单元测试用例，以清楚地了解此函数。

#### (Result) Size

```go
func (r Result) Size() int
```

Size is alias of function Len.

​	size 是函数 Len 的别名。

#### (Result) Structs

```go
func (r Result) Structs(pointer interface{}) (err error)
```

Structs converts `r` to struct slice. Note that the parameter `pointer` should be type of *[]struct/*[]*struct.

​	Structs `r` 转换为结构切片。请注意，参数 `pointer` 的类型应为 []struct/[]*struct。

#### (Result) Xml

```go
func (r Result) Xml(rootTag ...string) string
```

Xml converts `r` to XML format content.

​	Xml `r` 转换为 XML 格式的内容。

### type Schema

```go
type Schema struct {
	DB
}
```

Schema is a schema object from which it can then create a Model.

​	Schema 是一个架构对象，然后它可以从中创建模型。

### type SoftTimeOption <-2.6.3

```go
type SoftTimeOption struct {
	SoftTimeType SoftTimeType // The value type for soft time field.
}
```

SoftTimeOption is the option to customize soft time feature for Model.

​	SoftTimeOption 是用于自定义模型的软时间功能的选项。

### type SoftTimeType <-2.6.3

```go
type SoftTimeType int
```

SoftTimeType custom defines the soft time field type.

​	SoftTimeType 自定义定义软时间字段类型。

```go
const (
	SoftTimeTypeAuto           SoftTimeType = 0 // (Default)Auto detect the field type by table field type.
	SoftTimeTypeTime           SoftTimeType = 1 // Using datetime as the field value.
	SoftTimeTypeTimestamp      SoftTimeType = 2 // In unix seconds.
	SoftTimeTypeTimestampMilli SoftTimeType = 3 // In unix milliseconds.
	SoftTimeTypeTimestampMicro SoftTimeType = 4 // In unix microseconds.
	SoftTimeTypeTimestampNano  SoftTimeType = 5 // In unix nanoseconds.
)
```

### type Sql

```go
type Sql struct {
	Sql           string        // SQL string(may contain reserved char '?').
	Type          string        // SQL operation type.
	Args          []interface{} // Arguments for this sql.
	Format        string        // Formatted sql which contains arguments in the sql.
	Error         error         // Execution result.
	Start         int64         // Start execution timestamp in milliseconds.
	End           int64         // End execution timestamp in milliseconds.
	Group         string        // Group is the group name of the configuration that the sql is executed from.
	Schema        string        // Schema is the schema name of the configuration that the sql is executed from.
	IsTransaction bool          // IsTransaction marks whether this sql is executed in transaction.
	RowsAffected  int64         // RowsAffected marks retrieved or affected number with current sql statement.
}
```

Sql is the sql recording struct.

​	Sql 是 sql 记录结构。

### type SqlResult

```go
type SqlResult struct {
	Result   sql.Result
	Affected int64
}
```

SqlResult is execution result for sql operations. It also supports batch operation result for rowsAffected.

​	SqlResult 是 sql 操作的执行结果。它还支持对 rowsAffected 进行批量操作结果。

#### (*SqlResult) LastInsertId

```go
func (r *SqlResult) LastInsertId() (int64, error)
```

LastInsertId returns the integer generated by the database in response to a command. Typically, this will be from an “auto increment” column when inserting a new row. Not all databases support this feature, and the syntax of such statements varies. Also, See sql.Result.

​	LastInsertId 返回数据库为响应命令而生成的整数。通常，这将来自插入新行时的“自动递增”列。并非所有数据库都支持此功能，并且此类语句的语法各不相同。另外，请参阅 sql。结果。

#### (*SqlResult) MustGetAffected

```go
func (r *SqlResult) MustGetAffected() int64
```

MustGetAffected returns the affected rows count, if any error occurs, it panics.

​	MustGetAffected 返回受影响的行计数，如果发生任何错误，它会崩溃。

#### (*SqlResult) MustGetInsertId

```go
func (r *SqlResult) MustGetInsertId() int64
```

MustGetInsertId returns the last insert id, if any error occurs, it panics.

​	MustGetInsertId 返回最后一个插入 ID，如果发生任何错误，它会崩溃。

#### (*SqlResult) RowsAffected

```go
func (r *SqlResult) RowsAffected() (int64, error)
```

RowsAffected returns the number of rows affected by an update, insert, or delete. Not every database or database driver may support this. Also, See sql.Result.

​	RowsAffected 返回受更新、插入或删除影响的行数。并非每个数据库或数据库驱动程序都支持此功能。另外，请参阅 sql。结果。

### type Stmt

```go
type Stmt struct {
	*sql.Stmt
	// contains filtered or unexported fields
}
```

Stmt is a prepared statement. A Stmt is safe for concurrent use by multiple goroutines.

​	Stmt 是一个准备好的语句。Stmt 对于多个 goroutine 并发使用是安全的。

If a Stmt is prepared on a Tx or Conn, it will be bound to a single underlying connection forever. If the Tx or Conn closes, the Stmt will become unusable and all operations will return an error. If a Stmt is prepared on a DB, it will remain usable for the lifetime of the DB. When the Stmt needs to execute on a new underlying connection, it will prepare itself on the new connection automatically.

​	如果在 Tx 或 Conn 上准备了 Stmt，它将永远绑定到单个底层连接。如果 Tx 或 Conn 关闭，则 Stmt 将变得不可用，所有操作都将返回错误。如果在数据库上准备了 Stmt，则它将在数据库的生命周期内保持可用。当 Stmt 需要在新的底层连接上执行时，它将自动在新连接上做好准备。

#### (*Stmt) Close

```go
func (s *Stmt) Close() error
```

Close closes the statement.

​	Close 关闭语句。

#### (*Stmt) Exec

```go
func (s *Stmt) Exec(args ...interface{}) (sql.Result, error)
```

Exec executes a prepared statement with the given arguments and returns a Result summarizing the effect of the statement.

​	Exec 使用给定的参数执行准备好的语句，并返回一个 Result，总结该语句的效果。

#### (*Stmt) ExecContext

```go
func (s *Stmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
```

ExecContext executes a prepared statement with the given arguments and returns a Result summarizing the effect of the statement.

​	ExecContext 使用给定的参数执行预准备语句，并返回一个 Result，总结该语句的效果。

#### (*Stmt) Query

```go
func (s *Stmt) Query(args ...interface{}) (*sql.Rows, error)
```

Query executes a prepared query statement with the given arguments and returns the query results as a *Rows.

​	Query 使用给定的参数执行准备好的查询语句，并以 *Rows 的形式返回查询结果。

#### (*Stmt) QueryContext

```go
func (s *Stmt) QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error)
```

QueryContext executes a prepared query statement with the given arguments and returns the query results as a *Rows.

​	QueryContext 使用给定的参数执行准备好的查询语句，并以 *Rows 的形式返回查询结果。

#### (*Stmt) QueryRow

```go
func (s *Stmt) QueryRow(args ...interface{}) *sql.Row
```

QueryRow executes a prepared query statement with the given arguments. If an error occurs during the execution of the statement, that error will be returned by a call to Scan on the returned *Row, which is always non-nil. If the query selects no rows, the *Row’s Scan will return ErrNoRows. Otherwise, the *Row’s Scan scans the first selected row and discards the rest.

​	QueryRow 使用给定的参数执行准备好的查询语句。如果在语句执行过程中发生错误，则对返回的 *Row 的 Scan 调用将返回该错误，该行始终为非 nil。如果查询未选择任何行，则 *Row's Scan 将返回 ErrNoRows。否则，*行的扫描将扫描第一个选定的行并丢弃其余行。

Example usage:

​	用法示例：

```go
var name string
err := nameByUseridStmt.QueryRow(id).Scan(&name)
```

#### (*Stmt) QueryRowContext

```go
func (s *Stmt) QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row
```

QueryRowContext executes a prepared query statement with the given arguments. If an error occurs during the execution of the statement, that error will be returned by a call to Scan on the returned *Row, which is always non-nil. If the query selects no rows, the *Row’s Scan will return ErrNoRows. Otherwise, the *Row’s Scan scans the first selected row and discards the rest.

​	QueryRowContext 使用给定的参数执行准备好的查询语句。如果在语句执行过程中发生错误，则对返回的 *Row 的 Scan 调用将返回该错误，该行始终为非 nil。如果查询未选择任何行，则 *Row's Scan 将返回 ErrNoRows。否则，*行的扫描将扫描第一个选定的行并丢弃其余行。

### type TX

```go
type TX interface {
	Link

	Ctx(ctx context.Context) TX
	Raw(rawSql string, args ...interface{}) *Model
	Model(tableNameQueryOrStruct ...interface{}) *Model
	With(object interface{}) *Model

	Begin() error
	Commit() error
	Rollback() error
	Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error)

	Query(sql string, args ...interface{}) (result Result, err error)
	Exec(sql string, args ...interface{}) (sql.Result, error)
	Prepare(sql string) (*Stmt, error)

	GetAll(sql string, args ...interface{}) (Result, error)
	GetOne(sql string, args ...interface{}) (Record, error)
	GetStruct(obj interface{}, sql string, args ...interface{}) error
	GetStructs(objPointerSlice interface{}, sql string, args ...interface{}) error
	GetScan(pointer interface{}, sql string, args ...interface{}) error
	GetValue(sql string, args ...interface{}) (Value, error)
	GetCount(sql string, args ...interface{}) (int64, error)

	Insert(table string, data interface{}, batch ...int) (sql.Result, error)
	InsertIgnore(table string, data interface{}, batch ...int) (sql.Result, error)
	InsertAndGetId(table string, data interface{}, batch ...int) (int64, error)
	Replace(table string, data interface{}, batch ...int) (sql.Result, error)
	Save(table string, data interface{}, batch ...int) (sql.Result, error)
	Update(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)
	Delete(table string, condition interface{}, args ...interface{}) (sql.Result, error)

	GetCtx() context.Context
	GetDB() DB
	GetSqlTX() *sql.Tx
	IsClosed() bool

	SavePoint(point string) error
	RollbackTo(point string) error
}
```

TX defines the interfaces for ORM transaction operations.

​	TX 定义了 ORM 事务操作的接口。

#### func TXFromCtx

```go
func TXFromCtx(ctx context.Context, group string) TX
```

TXFromCtx retrieves and returns transaction object from context. It is usually used in nested transaction feature, and it returns nil if it is not set previously.

​	TXFromCtx 从上下文中检索并返回事务对象。它通常用于嵌套事务功能，如果之前没有设置，则返回 nil。

### type TXCore <-2.3.0

```go
type TXCore struct {
	// contains filtered or unexported fields
}
```

TXCore is the struct for transaction management.

​	TXCore 是事务管理的结构。

#### (*TXCore) Begin

```go
func (tx *TXCore) Begin() error
```

Begin starts a nested transaction procedure.

​	Begin 启动嵌套事务过程。

#### (*TXCore) Commit

```go
func (tx *TXCore) Commit() error
```

Commit commits current transaction. Note that it releases previous saved transaction point if it’s in a nested transaction procedure, or else it commits the hole transaction.

​	Commit 提交当前事务。请注意，如果它位于嵌套事务过程中，它会释放以前保存的事务点，否则它会提交漏洞事务。

#### (*TXCore) Ctx

```go
func (tx *TXCore) Ctx(ctx context.Context) TX
```

Ctx sets the context for current transaction.

​	Ctx 设置当前事务的上下文。

#### (*TXCore) Delete

```go
func (tx *TXCore) Delete(table string, condition interface{}, args ...interface{}) (sql.Result, error)
```

Delete does “DELETE FROM … " statement for the table.

​	Delete 执行“DELETE FROM ...“的语句。

The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc. It is commonly used with parameter `args`. Eg: “uid=10000”, “uid”, 10000 “money>? AND name like ?”, 99999, “vip_%” “status IN (?)”, g.Slice{1,2,3} “age IN(?,?)”, 18, 50 User{ Id : 1, UserName : “john”}.

​	参数 `condition` 可以是字符串/地图/gmap/slice/struct/*struct等类型。它通常与参数一起使用 `args` 。例如：“uid=10000”、“uid”、10000“钱>？AND name like ？“， 99999， ”vip_%“ ”status IN （？）“， g.Slice{1,2,3} ”age IN（?,?）“， 18， 50 User{ Id ： 1， UserName ： ”john“}.

#### (*TXCore) Exec

```go
func (tx *TXCore) Exec(sql string, args ...interface{}) (sql.Result, error)
```

Exec does none query operation on transaction. See Core.Exec.

​	Exec 不对事务执行任何查询操作。请参阅 Core.Exec。

#### (*TXCore) ExecContext

```go
func (tx *TXCore) ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
```

ExecContext implements interface function Link.ExecContext.

​	ExecContext 实现接口函数 Link.ExecContext。

#### (*TXCore) GetAll

```go
func (tx *TXCore) GetAll(sql string, args ...interface{}) (Result, error)
```

GetAll queries and returns data records from database.

​	GetAll 查询并返回数据库中的数据记录。

#### (*TXCore) GetCount

```go
func (tx *TXCore) GetCount(sql string, args ...interface{}) (int64, error)
```

GetCount queries and returns the count from database.

​	GetCount 查询并返回数据库中的计数。

#### (*TXCore) GetCtx

```go
func (tx *TXCore) GetCtx() context.Context
```

GetCtx returns the context for current transaction.

​	GetCtx 返回当前事务的上下文。

#### (*TXCore) GetDB

```go
func (tx *TXCore) GetDB() DB
```

GetDB returns the DB for current transaction.

​	GetDB 返回当前事务的数据库。

#### (*TXCore) GetOne

```go
func (tx *TXCore) GetOne(sql string, args ...interface{}) (Record, error)
```

GetOne queries and returns one record from database.

​	GetOne 查询并返回数据库中的一条记录。

#### (*TXCore) GetScan

```go
func (tx *TXCore) GetScan(pointer interface{}, sql string, args ...interface{}) error
```

GetScan queries one or more records from database and converts them to given struct or struct array.

​	GetScan 从数据库中查询一条或多条记录，并将它们转换为给定的 struct 或 struct 数组。

If parameter `pointer` is type of struct pointer, it calls GetStruct internally for the conversion. If parameter `pointer` is type of slice, it calls GetStructs internally for conversion.

​	如果 parameter `pointer` 是结构指针的类型，则它会在内部调用 GetStruct 进行转换。如果 parameter `pointer` 是切片的类型，则它会在内部调用 GetStructs 进行转换。

#### (*TXCore) GetSqlTX

```go
func (tx *TXCore) GetSqlTX() *sql.Tx
```

GetSqlTX returns the underlying transaction object for current transaction.

​	GetSqlTX 返回当前事务的基础事务对象。

#### (*TXCore) GetStruct

```go
func (tx *TXCore) GetStruct(obj interface{}, sql string, args ...interface{}) error
```

GetStruct queries one record from database and converts it to given struct. The parameter `pointer` should be a pointer to struct.

​	GetStruct 从数据库中查询一条记录，并将其转换为给定的结构。该参数 `pointer` 应是指向结构的指针。

#### (*TXCore) GetStructs

```go
func (tx *TXCore) GetStructs(objPointerSlice interface{}, sql string, args ...interface{}) error
```

GetStructs queries records from database and converts them to given struct. The parameter `pointer` should be type of struct slice: []struct/[]*struct.

​	GetStructs 从数据库查询记录，并将其转换为给定的结构。参数 `pointer` 应为 struct slice 类型：[]struct/[]*struct。

#### (*TXCore) GetValue

```go
func (tx *TXCore) GetValue(sql string, args ...interface{}) (Value, error)
```

GetValue queries and returns the field value from database. The sql should query only one field from database, or else it returns only one field of the result.

​	GetValue 查询并返回数据库中的字段值。sql 应该只从数据库中查询一个字段，否则它只返回结果的一个字段。

#### (*TXCore) Insert

```go
func (tx *TXCore) Insert(table string, data interface{}, batch ...int) (sql.Result, error)
```

Insert does “INSERT INTO …” statement for the table. If there’s already one unique record of the data in the table, it returns error.

​	Insert 执行“INSERT INTO ...”表的语句。如果表中已有一条数据的唯一记录，则返回错误。

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”})

​	参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。 例如： Data（g.Map{“uid”： 10000， “name”：“john”}） 数据（g.Slice{g.Map{“uid”： 10000， “name”：“john”}， g.Map{“uid”： 20000， “name”：“smith”}）

The parameter `batch` specifies the batch operation count when given data is slice.

​	该参数 `batch` 指定给定数据为切片时的批处理操作计数。

#### (*TXCore) InsertAndGetId

```go
func (tx *TXCore) InsertAndGetId(table string, data interface{}, batch ...int) (int64, error)
```

InsertAndGetId performs action Insert and returns the last insert id that automatically generated.

​	InsertAndGetId 执行操作 Insert 并返回自动生成的最后一个插入 ID。

#### (*TXCore) InsertIgnore

```go
func (tx *TXCore) InsertIgnore(table string, data interface{}, batch ...int) (sql.Result, error)
```

InsertIgnore does “INSERT IGNORE INTO …” statement for the table. If there’s already one unique record of the data in the table, it ignores the inserting.

​	InsertIgnore 执行“INSERT IGNORE INTO ...”表的语句。如果表中已经有一条数据的唯一记录，则会忽略插入。

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”})

​	参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。 例如： Data（g.Map{“uid”： 10000， “name”：“john”}） 数据（g.Slice{g.Map{“uid”： 10000， “name”：“john”}， g.Map{“uid”： 20000， “name”：“smith”}）

The parameter `batch` specifies the batch operation count when given data is slice.

​	该参数 `batch` 指定给定数据为切片时的批处理操作计数。

#### (*TXCore) IsClosed

```go
func (tx *TXCore) IsClosed() bool
```

IsClosed checks and returns this transaction has already been committed or rolled back.

​	IsClosed 检查并返回此事务已提交或回滚。

#### (*TXCore) IsOnMaster

```go
func (tx *TXCore) IsOnMaster() bool
```

IsOnMaster implements interface function Link.IsOnMaster.

​	IsOnMaster 实现接口函数 Link.IsOnMaster。

#### (*TXCore) IsTransaction

```go
func (tx *TXCore) IsTransaction() bool
```

IsTransaction implements interface function Link.IsTransaction.

​	IsTransaction 实现接口函数 Link.IsTransaction。

#### (*TXCore) Model

```go
func (tx *TXCore) Model(tableNameQueryOrStruct ...interface{}) *Model
```

Model acts like Core.Model except it operates on transaction. See Core.Model.

​	Model 的行为类似于 Core.Model，只不过它对事务进行操作。请参阅 Core.Model。

#### (*TXCore) Prepare

```go
func (tx *TXCore) Prepare(sql string) (*Stmt, error)
```

Prepare creates a prepared statement for later queries or executions. Multiple queries or executions may be run concurrently from the returned statement. The caller must call the statement’s Close method when the statement is no longer needed.

​	Prepare 为以后的查询或执行创建预准备语句。可以从返回的语句同时运行多个查询或执行。当不再需要语句时，调用方必须调用语句的 Close 方法。

#### (*TXCore) PrepareContext

```go
func (tx *TXCore) PrepareContext(ctx context.Context, sql string) (*sql.Stmt, error)
```

PrepareContext implements interface function Link.PrepareContext.

​	PrepareContext 实现接口函数 Link.PrepareContext。

#### (*TXCore) Query

```go
func (tx *TXCore) Query(sql string, args ...interface{}) (result Result, err error)
```

Query does query operation on transaction. See Core.Query.

​	Query 对事务执行查询操作。请参阅 Core.Query。

#### (*TXCore) QueryContext

```go
func (tx *TXCore) QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
```

QueryContext implements interface function Link.QueryContext.

​	QueryContext 实现接口函数 Link.QueryContext。

#### (*TXCore) Raw

```go
func (tx *TXCore) Raw(rawSql string, args ...interface{}) *Model
```

#### (*TXCore) Replace

```go
func (tx *TXCore) Replace(table string, data interface{}, batch ...int) (sql.Result, error)
```

Replace does “REPLACE INTO …” statement for the table. If there’s already one unique record of the data in the table, it deletes the record and inserts a new one.

​	替换执行“REPLACE INTO ...”表的语句。如果表中已有一条数据的唯一记录，则会删除该记录并插入一条新记录。

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”})

​	参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。 例如： Data（g.Map{“uid”： 10000， “name”：“john”}） 数据（g.Slice{g.Map{“uid”： 10000， “name”：“john”}， g.Map{“uid”： 20000， “name”：“smith”}）

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. If given data is type of slice, it then does batch replacing, and the optional parameter `batch` specifies the batch operation count.

​	参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。如果给定的数据是切片类型，则执行批量替换，可选参数 `batch` 指定批处理操作计数。

#### (*TXCore) Rollback

```go
func (tx *TXCore) Rollback() error
```

Rollback aborts current transaction. Note that it aborts current transaction if it’s in a nested transaction procedure, or else it aborts the hole transaction.

​	回滚会中止当前事务。请注意，如果当前事务位于嵌套事务过程中，则会中止当前事务，否则会中止空洞事务。

#### (*TXCore) RollbackTo

```go
func (tx *TXCore) RollbackTo(point string) error
```

RollbackTo performs `ROLLBACK TO SAVEPOINT xxx` SQL statement that rollbacks to specified saved transaction. The parameter `point` specifies the point name that was saved previously.

​	RollbackTo 执行回滚到指定保存事务的 `ROLLBACK TO SAVEPOINT xxx` SQL 语句。该参数 `point` 指定之前保存的点名称。

#### (*TXCore) Save

```go
func (tx *TXCore) Save(table string, data interface{}, batch ...int) (sql.Result, error)
```

Save does “INSERT INTO … ON DUPLICATE KEY UPDATE…” statement for the table. It updates the record if there’s primary or unique index in the saving data, or else it inserts a new record into the table.

​	保存执行“插入...在重复的密钥更新中......”表的语句。如果保存数据中有主索引或唯一索引，它将更新记录，或者将新记录插入到表中。

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{“uid”: 10000, “name”:“john”}) Data(g.Slice{g.Map{“uid”: 10000, “name”:“john”}, g.Map{“uid”: 20000, “name”:“smith”})

​	参数 `data` 可以是 map/gmap/struct/*struct/[]map/[]struct 等类型。 例如： Data（g.Map{“uid”： 10000， “name”：“john”}） 数据（g.Slice{g.Map{“uid”： 10000， “name”：“john”}， g.Map{“uid”： 20000， “name”：“smith”}）

If given data is type of slice, it then does batch saving, and the optional parameter `batch` specifies the batch operation count.

​	如果给定的数据是切片类型，则执行批量保存，可选参数 `batch` 指定批处理操作计数。

#### (*TXCore) SavePoint

```go
func (tx *TXCore) SavePoint(point string) error
```

SavePoint performs `SAVEPOINT xxx` SQL statement that saves transaction at current point. The parameter `point` specifies the point name that will be saved to server.

​	SavePoint 执行 `SAVEPOINT xxx` SQL 语句，以保存当前点的事务。该参数 `point` 指定将保存到服务器的点名称。

#### (*TXCore) Transaction

```go
func (tx *TXCore) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error)
```

Transaction wraps the transaction logic using function `f`. It rollbacks the transaction and returns the error from function `f` if it returns non-nil error. It commits the transaction and returns nil if function `f` returns nil.

​	事务使用函数 `f` 包装事务逻辑。 `f` 它会回滚事务，如果函数返回非 nil 错误，则返回错误。它提交事务，如果函数 `f` 返回 nil，则返回 nil。

Note that, you should not Commit or Rollback the transaction in function `f` as it is automatically handled by this function.

​	请注意，您不应该在函数 `f` 中提交或回滚事务，因为它是由此函数自动处理的。

#### (*TXCore) Update

```go
func (tx *TXCore) Update(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)
```

Update does “UPDATE … " statement for the table.

​	更新执行“更新...“的语句。

The parameter `data` can be type of string/map/gmap/struct/*struct, etc. Eg: “uid=10000”, “uid”, 10000, g.Map{“uid”: 10000, “name”:“john”}

​	参数 `data` 可以是 string/map/gmap/struct/*struct 等类型。例如： “uid=10000”， “uid”， 10000， g.Map{“uid”： 10000， “name”：“john”}

The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc. It is commonly used with parameter `args`. Eg: “uid=10000”, “uid”, 10000 “money>? AND name like ?”, 99999, “vip_%” “status IN (?)”, g.Slice{1,2,3} “age IN(?,?)”, 18, 50 User{ Id : 1, UserName : “john”}.

​	参数 `condition` 可以是字符串/地图/gmap/slice/struct/*struct等类型。它通常与参数一起使用 `args` 。例如：“uid=10000”、“uid”、10000“钱>？AND name like ？“， 99999， ”vip_%“ ”status IN （？）“， g.Slice{1,2,3} ”age IN（?,?）“， 18， 50 User{ Id ： 1， UserName ： ”john“}.

#### (*TXCore) With

```go
func (tx *TXCore) With(object interface{}) *Model
```

With acts like Core.With except it operates on transaction. See Core.With.

​	与 Core.With 类似，但它在事务上运行。请参阅 Core.With。

### type TableField

```go
type TableField struct {
	Index   int         // For ordering purpose as map is unordered.
	Name    string      // Field name.
	Type    string      // Field type. Eg: 'int(10) unsigned', 'varchar(64)'.
	Null    bool        // Field can be null or not.
	Key     string      // The index information(empty if it's not an index). Eg: PRI, MUL.
	Default interface{} // Default value for the field.
	Extra   string      // Extra information. Eg: auto_increment.
	Comment string      // Field comment.
}
```

TableField is the struct for table field.

​	TableField 是表字段的结构。

### type Value

```go
type Value = *gvar.Var // Value is the field value type.
```

### type WhereBuilder <-2.1.0

```go
type WhereBuilder struct {
	// contains filtered or unexported fields
}
```

WhereBuilder holds multiple where conditions in a group.

​	WhereBuilder 在一个组中保存多个 where 条件。

#### (*WhereBuilder) Build

```go
func (b *WhereBuilder) Build() (conditionWhere string, conditionArgs []interface{})
```

Build builds current WhereBuilder and returns the condition string and parameters.

​	Build 生成当前 WhereBuilder 并返回条件字符串和参数。

#### (*WhereBuilder) Clone

```go
func (b *WhereBuilder) Clone() *WhereBuilder
```

Clone clones and returns a WhereBuilder that is a copy of current one.

​	克隆克隆并返回一个 WhereBuilder，该 WhereBuilder 是当前克隆的副本。

#### (*WhereBuilder) Where

```go
func (b *WhereBuilder) Where(where interface{}, args ...interface{}) *WhereBuilder
```

Where sets the condition statement for the builder. The parameter `where` can be type of string/map/gmap/slice/struct/*struct, etc. Note that, if it’s called more than one times, multiple conditions will be joined into where statement using “AND”. Eg: Where(“uid=10000”) Where(“uid”, 10000) Where(“money>? AND name like ?”, 99999, “vip_%”) Where(“uid”, 1).Where(“name”, “john”) Where(“status IN (?)”, g.Slice{1,2,3}) Where(“age IN(?,?)”, 18, 50) Where(User{ Id : 1, UserName : “john”}).

​	Where 设置构建器的条件语句。参数 `where` 可以是字符串/地图/gmap/slice/struct/*struct等类型。请注意，如果它被调用了多次，则多个条件将使用“AND”加入到 where 语句中。例如：Where（“uid=10000”） where（“uid”， 10000） where（“money>？AND name like ？“， 99999， ”vip_%“） Where（”uid“， 1）。where（“name”， “john”） Where（“status IN （？）”， g.Slice{1,2,3}） Where（“age IN（?,?）”， 18， 50） Where（User{ id ： 1， userName ： “john”}）.

#### (*WhereBuilder) WhereBetween

```go
func (b *WhereBuilder) WhereBetween(column string, min, max interface{}) *WhereBuilder
```

WhereBetween builds `column BETWEEN min AND max` statement.

​	WhereBetween 构建 `column BETWEEN min AND max` 语句。

#### (*WhereBuilder) WhereGT

```go
func (b *WhereBuilder) WhereGT(column string, value interface{}) *WhereBuilder
```

WhereGT builds `column > value` statement.

​	WhereGT 构建 `column > value` 语句。

#### (*WhereBuilder) WhereGTE

```go
func (b *WhereBuilder) WhereGTE(column string, value interface{}) *WhereBuilder
```

WhereGTE builds `column >= value` statement.

​	其中 GTE 构建 `column >= value` 语句。

#### (*WhereBuilder) WhereIn

```go
func (b *WhereBuilder) WhereIn(column string, in interface{}) *WhereBuilder
```

WhereIn builds `column IN (in)` statement.

​	其中 builds `column IN (in)` 语句。

#### (*WhereBuilder) WhereLT

```go
func (b *WhereBuilder) WhereLT(column string, value interface{}) *WhereBuilder
```

WhereLT builds `column < value` statement.

​	其中 LT 构建 `column < value` 语句。

#### (*WhereBuilder) WhereLTE

```go
func (b *WhereBuilder) WhereLTE(column string, value interface{}) *WhereBuilder
```

WhereLTE builds `column <= value` statement.

​	其中LTE构建 `column <= value` 语句。

#### (*WhereBuilder) WhereLike

```go
func (b *WhereBuilder) WhereLike(column string, like string) *WhereBuilder
```

WhereLike builds `column LIKE like` statement.

​	WhereLike 构建 `column LIKE like` 语句。

#### (*WhereBuilder) WhereNot

```go
func (b *WhereBuilder) WhereNot(column string, value interface{}) *WhereBuilder
```

WhereNot builds `column != value` statement.

​	WhereNot 构建 `column != value` 语句。

#### (*WhereBuilder) WhereNotBetween

```go
func (b *WhereBuilder) WhereNotBetween(column string, min, max interface{}) *WhereBuilder
```

WhereNotBetween builds `column NOT BETWEEN min AND max` statement.

​	WhereNotBetween 构建 `column NOT BETWEEN min AND max` 语句。

#### (*WhereBuilder) WhereNotIn

```go
func (b *WhereBuilder) WhereNotIn(column string, in interface{}) *WhereBuilder
```

WhereNotIn builds `column NOT IN (in)` statement.

​	WhereNotIn 构建 `column NOT IN (in)` 语句。

#### (*WhereBuilder) WhereNotLike

```go
func (b *WhereBuilder) WhereNotLike(column string, like interface{}) *WhereBuilder
```

WhereNotLike builds `column NOT LIKE like` statement.

​	WhereNotLike 构建 `column NOT LIKE like` 语句。

#### (*WhereBuilder) WhereNotNull

```go
func (b *WhereBuilder) WhereNotNull(columns ...string) *WhereBuilder
```

WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.

​	WhereNotNull 生成 `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` 语句。

#### (*WhereBuilder) WhereNull

```go
func (b *WhereBuilder) WhereNull(columns ...string) *WhereBuilder
```

WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.

​	WhereNull 构建 `columns[0] IS NULL AND columns[1] IS NULL ...` 语句。

#### (*WhereBuilder) WhereOr

```go
func (b *WhereBuilder) WhereOr(where interface{}, args ...interface{}) *WhereBuilder
```

WhereOr adds “OR” condition to the where statement.

​	WhereOr 将“OR”条件添加到 where 语句中。

#### (*WhereBuilder) WhereOrBetween

```go
func (b *WhereBuilder) WhereOrBetween(column string, min, max interface{}) *WhereBuilder
```

WhereOrBetween builds `column BETWEEN min AND max` statement in `OR` conditions.

​	WhereOrBetween 在 conditions 中 `OR` 构建 `column BETWEEN min AND max` 语句。

#### (*WhereBuilder) WhereOrGT

```go
func (b *WhereBuilder) WhereOrGT(column string, value interface{}) *WhereBuilder
```

WhereOrGT builds `column > value` statement in `OR` conditions.

​	WhereOrGT 在 conditions 中 `OR` 构建 `column > value` 语句。

#### (*WhereBuilder) WhereOrGTE

```go
func (b *WhereBuilder) WhereOrGTE(column string, value interface{}) *WhereBuilder
```

WhereOrGTE builds `column >= value` statement in `OR` conditions.

​	WhereOrGTE 在 conditions 中 `OR` 构建 `column >= value` 语句。

#### (*WhereBuilder) WhereOrIn

```go
func (b *WhereBuilder) WhereOrIn(column string, in interface{}) *WhereBuilder
```

WhereOrIn builds `column IN (in)` statement in `OR` conditions.

​	WhereOrIn 在 conditions 中 `OR` 构建 `column IN (in)` 语句。

#### (*WhereBuilder) WhereOrLT

```go
func (b *WhereBuilder) WhereOrLT(column string, value interface{}) *WhereBuilder
```

WhereOrLT builds `column < value` statement in `OR` conditions.

​	WhereOrLT 在 conditions 中 `OR` 构建 `column < value` 语句。

#### (*WhereBuilder) WhereOrLTE

```go
func (b *WhereBuilder) WhereOrLTE(column string, value interface{}) *WhereBuilder
```

WhereOrLTE builds `column <= value` statement in `OR` conditions.

​	WhereOrLTE 在 conditions 中 `OR` 构建 `column <= value` 语句。

#### (*WhereBuilder) WhereOrLike

```go
func (b *WhereBuilder) WhereOrLike(column string, like interface{}) *WhereBuilder
```

WhereOrLike builds `column LIKE 'like'` statement in `OR` conditions.

​	WhereOrLike 在 conditions 中 `OR` 构建 `column LIKE 'like'` 语句。

#### (*WhereBuilder) WhereOrNot

```go
func (b *WhereBuilder) WhereOrNot(column string, value interface{}) *WhereBuilder
```

WhereOrNot builds `column != value` statement in `OR` conditions.

​	WhereOrNot 在 conditions 中 `OR` 构建 `column != value` 语句。

#### (*WhereBuilder) WhereOrNotBetween

```go
func (b *WhereBuilder) WhereOrNotBetween(column string, min, max interface{}) *WhereBuilder
```

WhereOrNotBetween builds `column NOT BETWEEN min AND max` statement in `OR` conditions.

​	WhereOrNotBetween 在 conditions 中 `OR` 构建 `column NOT BETWEEN min AND max` 语句。

#### (*WhereBuilder) WhereOrNotIn

```go
func (b *WhereBuilder) WhereOrNotIn(column string, in interface{}) *WhereBuilder
```

WhereOrNotIn builds `column NOT IN (in)` statement.

​	WhereOrNotIn 生成 `column NOT IN (in)` 语句。

#### (*WhereBuilder) WhereOrNotLike

```go
func (b *WhereBuilder) WhereOrNotLike(column string, like interface{}) *WhereBuilder
```

WhereOrNotLike builds `column NOT LIKE like` statement in `OR` conditions.

​	WhereOrNotLike 在 conditions 中 `OR` 构建 `column NOT LIKE like` 语句。

#### (*WhereBuilder) WhereOrNotNull

```go
func (b *WhereBuilder) WhereOrNotNull(columns ...string) *WhereBuilder
```

WhereOrNotNull builds `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` statement in `OR` conditions.

​	WhereOrNotNull 在 conditions 中 `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` `OR` 生成语句。

#### (*WhereBuilder) WhereOrNull

```go
func (b *WhereBuilder) WhereOrNull(columns ...string) *WhereBuilder
```

WhereOrNull builds `columns[0] IS NULL OR columns[1] IS NULL ...` statement in `OR` conditions.

​	WhereOrNull 在 conditions 中 `OR` 构建 `columns[0] IS NULL OR columns[1] IS NULL ...` 语句。

#### (*WhereBuilder) WhereOrPrefix

```go
func (b *WhereBuilder) WhereOrPrefix(prefix string, where interface{}, args ...interface{}) *WhereBuilder
```

WhereOrPrefix performs as WhereOr, but it adds prefix to each field in where statement. Eg: WhereOrPrefix(“order”, “status”, “paid”) => WHERE xxx OR (`order`.`status`=‘paid’) WhereOrPrefix(“order”, struct{Status:“paid”, “channel”:“bank”}) => WHERE xxx OR (`order`.`status`=‘paid’ AND `order`.`channel`=‘bank’)

​	WhereOrPrefix 的执行方式为 WhereOr，但它向 where 语句中的每个字段添加前缀。例如：WhereOrPrefix（“order”， “status”， “paid”） => WHERE xxx OR （ `order` . `status` ='paid'） WhereOrPrefix（“order”， struct{Status：“paid”， “channel”：“bank”}） => WHERE xxx OR （ `order` . `status` ='paid' 和 `order` . `channel` ='银行'）

#### (*WhereBuilder) WhereOrPrefixBetween

```go
func (b *WhereBuilder) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder
```

WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in `OR` conditions.

​	WhereOrPrefixBetween 在 conditions 中 `OR` 构建 `prefix.column BETWEEN min AND max` 语句。

#### (*WhereBuilder) WhereOrPrefixGT

```go
func (b *WhereBuilder) WhereOrPrefixGT(prefix string, column string, value interface{}) *WhereBuilder
```

WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions.

​	WhereOrPrefixGT 在 conditions 中 `prefix.column > value` `OR` 构建语句。

#### (*WhereBuilder) WhereOrPrefixGTE

```go
func (b *WhereBuilder) WhereOrPrefixGTE(prefix string, column string, value interface{}) *WhereBuilder
```

WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions.

​	WhereOrPrefixGTE 在 conditions 中 `OR` 构建 `prefix.column >= value` 语句。

#### (*WhereBuilder) WhereOrPrefixIn

```go
func (b *WhereBuilder) WhereOrPrefixIn(prefix string, column string, in interface{}) *WhereBuilder
```

WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions.

​	WhereOrPrefixIn 在 conditions 中 `OR` 生成 `prefix.column IN (in)` 语句。

#### (*WhereBuilder) WhereOrPrefixLT

```go
func (b *WhereBuilder) WhereOrPrefixLT(prefix string, column string, value interface{}) *WhereBuilder
```

WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions.

​	WhereOrPrefixLT 在 conditions 中 `OR` 构建 `prefix.column < value` 语句。

#### (*WhereBuilder) WhereOrPrefixLTE

```go
func (b *WhereBuilder) WhereOrPrefixLTE(prefix string, column string, value interface{}) *WhereBuilder
```

WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions.

​	WhereOrPrefixLTE 在 conditions 中 `OR` 构建 `prefix.column <= value` 语句。

#### (*WhereBuilder) WhereOrPrefixLike

```go
func (b *WhereBuilder) WhereOrPrefixLike(prefix string, column string, like interface{}) *WhereBuilder
```

WhereOrPrefixLike builds `prefix.column LIKE 'like'` statement in `OR` conditions.

​	WhereOrPrefixLike 在 conditions 中 `OR` 构建 `prefix.column LIKE 'like'` 语句。

#### (*WhereBuilder) WhereOrPrefixNot

```go
func (b *WhereBuilder) WhereOrPrefixNot(prefix string, column string, value interface{}) *WhereBuilder
```

WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions.

​	WhereOrPrefixNot 在 conditions 中 `OR` 构建 `prefix.column != value` 语句。

#### (*WhereBuilder) WhereOrPrefixNotBetween

```go
func (b *WhereBuilder) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder
```

WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement in `OR` conditions.

​	WhereOrPrefixNotBetween 在 conditions 中 `OR` 构建 `prefix.column NOT BETWEEN min AND max` 语句。

#### (*WhereBuilder) WhereOrPrefixNotIn

```go
func (b *WhereBuilder) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder
```

WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement.

​	WhereOrPrefixNotIn 生成 `prefix.column NOT IN (in)` 语句。

#### (*WhereBuilder) WhereOrPrefixNotLike

```go
func (b *WhereBuilder) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder
```

WhereOrPrefixNotLike builds `prefix.column NOT LIKE 'like'` statement in `OR` conditions.

​	WhereOrPrefixNotLike 在 conditions 中 `OR` 构建 `prefix.column NOT LIKE 'like'` 语句。

#### (*WhereBuilder) WhereOrPrefixNotNull

```go
func (b *WhereBuilder) WhereOrPrefixNotNull(prefix string, columns ...string) *WhereBuilder
```

WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` statement in `OR` conditions.

​	WhereOrPrefixNotNull 在 conditions 中 `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` `OR` 构建语句。

#### (*WhereBuilder) WhereOrPrefixNull

```go
func (b *WhereBuilder) WhereOrPrefixNull(prefix string, columns ...string) *WhereBuilder
```

WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` statement in `OR` conditions.

​	WhereOrPrefixNull 在 conditions 中 `OR` 构建 `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` 语句。

#### (*WhereBuilder) WhereOrf

```go
func (b *WhereBuilder) WhereOrf(format string, args ...interface{}) *WhereBuilder
```

WhereOrf builds `OR` condition string using fmt.Sprintf and arguments. Eg: WhereOrf(`amount<? and status=%s`, “paid”, 100) => WHERE xxx OR `amount`<100 and status=‘paid’ WhereOrf(`amount<%d and status=%s`, 100, “paid”) => WHERE xxx OR `amount`<100 and status=‘paid’

​	WhereOrf 使用 fmt `OR` 构建条件字符串。Sprintf 和参数。例如： WhereOrf（ `amount<? and status=%s` ， “paid”， 100） => WHERE xxx OR `amount` <100 and status='paid' WhereOrf（ `amount<%d and status=%s` ， 100， “paid”） => WHERE xxx OR `amount` <100 and status='paid'

#### (*WhereBuilder) WherePrefix

```go
func (b *WhereBuilder) WherePrefix(prefix string, where interface{}, args ...interface{}) *WhereBuilder
```

WherePrefix performs as Where, but it adds prefix to each field in where statement. Eg: WherePrefix(“order”, “status”, “paid”) => WHERE `order`.`status`=‘paid’ WherePrefix(“order”, struct{Status:“paid”, “channel”:“bank”}) => WHERE `order`.`status`=‘paid’ AND `order`.`channel`=‘bank’

​	WherePrefix 的执行方式为 Where，但它向 where 语句中的每个字段添加前缀。例如：WherePrefix（“order”， “status”， “paid”） => WHERE `order` . `status` ='paid' WherePrefix（“order”， struct{Status：“paid”， “channel”：“bank”}） => WHERE `order` . `status` ='paid' 和 `order` . `channel` ='银行'

#### (*WhereBuilder) WherePrefixBetween

```go
func (b *WhereBuilder) WherePrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder
```

WherePrefixBetween builds `prefix.column BETWEEN min AND max` statement.

​	WherePrefixBetween 构建 `prefix.column BETWEEN min AND max` 语句。

#### (*WhereBuilder) WherePrefixGT

```go
func (b *WhereBuilder) WherePrefixGT(prefix string, column string, value interface{}) *WhereBuilder
```

WherePrefixGT builds `prefix.column > value` statement.

​	WherePrefixGT 构建 `prefix.column > value` 语句。

#### (*WhereBuilder) WherePrefixGTE

```go
func (b *WhereBuilder) WherePrefixGTE(prefix string, column string, value interface{}) *WhereBuilder
```

WherePrefixGTE builds `prefix.column >= value` statement.

​	WherePrefixGTE 生成 `prefix.column >= value` 语句。

#### (*WhereBuilder) WherePrefixIn

```go
func (b *WhereBuilder) WherePrefixIn(prefix string, column string, in interface{}) *WhereBuilder
```

WherePrefixIn builds `prefix.column IN (in)` statement.

​	WherePrefixIn 构建 `prefix.column IN (in)` 语句。

#### (*WhereBuilder) WherePrefixLT

```go
func (b *WhereBuilder) WherePrefixLT(prefix string, column string, value interface{}) *WhereBuilder
```

WherePrefixLT builds `prefix.column < value` statement.

​	WherePrefixLT 生成 `prefix.column < value` 语句。

#### (*WhereBuilder) WherePrefixLTE

```go
func (b *WhereBuilder) WherePrefixLTE(prefix string, column string, value interface{}) *WhereBuilder
```

WherePrefixLTE builds `prefix.column <= value` statement.

​	WherePrefixLTE 构建 `prefix.column <= value` 语句。

#### (*WhereBuilder) WherePrefixLike

```go
func (b *WhereBuilder) WherePrefixLike(prefix string, column string, like interface{}) *WhereBuilder
```

WherePrefixLike builds `prefix.column LIKE like` statement.

​	WherePrefixLike 构建语 `prefix.column LIKE like` 句。

#### (*WhereBuilder) WherePrefixNot

```go
func (b *WhereBuilder) WherePrefixNot(prefix string, column string, value interface{}) *WhereBuilder
```

WherePrefixNot builds `prefix.column != value` statement.

​	WherePrefixNot 构建 `prefix.column != value` 语句。

#### (*WhereBuilder) WherePrefixNotBetween

```go
func (b *WhereBuilder) WherePrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder
```

WherePrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement.

​	WherePrefixNotBetween 构建语 `prefix.column NOT BETWEEN min AND max` 句。

#### (*WhereBuilder) WherePrefixNotIn

```go
func (b *WhereBuilder) WherePrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder
```

WherePrefixNotIn builds `prefix.column NOT IN (in)` statement.

​	WherePrefixNotIn 构建 `prefix.column NOT IN (in)` 语句。

#### (*WhereBuilder) WherePrefixNotLike

```go
func (b *WhereBuilder) WherePrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder
```

WherePrefixNotLike builds `prefix.column NOT LIKE like` statement.

​	WherePrefixNotLike 构建语 `prefix.column NOT LIKE like` 句。

#### (*WhereBuilder) WherePrefixNotNull

```go
func (b *WhereBuilder) WherePrefixNotNull(prefix string, columns ...string) *WhereBuilder
```

WherePrefixNotNull builds `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` statement.

​	WherePrefixNotNull 生成 `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` 语句。

#### (*WhereBuilder) WherePrefixNull

```go
func (b *WhereBuilder) WherePrefixNull(prefix string, columns ...string) *WhereBuilder
```

WherePrefixNull builds `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` statement.

​	WherePrefixNull 构建语 `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` 句。

#### (*WhereBuilder) WherePri

```go
func (b *WhereBuilder) WherePri(where interface{}, args ...interface{}) *WhereBuilder
```

WherePri does the same logic as Model.Where except that if the parameter `where` is a single condition like int/string/float/slice, it treats the condition as the primary key value. That is, if primary key is “id” and given `where` parameter as “123”, the WherePri function treats the condition as “id=123”, but Model.Where treats the condition as string “123”.

​	WherePri 执行与 Model.Where 相同的逻辑，不同之处在于如果参数 `where` 是单个条件（如 int/string/float/slice），则它将该条件视为主键值。也就是说，如果主键为“id”，并且给定 `where` 参数为“123”，则 WherePri 函数将条件视为“id=123”，但 Model.Where 将条件视为字符串“123”。

#### (*WhereBuilder) Wheref

```go
func (b *WhereBuilder) Wheref(format string, args ...interface{}) *WhereBuilder
```

Wheref builds condition string using fmt.Sprintf and arguments. Note that if the number of `args` is more than the placeholder in `format`, the extra `args` will be used as the where condition arguments of the Model. Eg: Wheref(`amount<? and status=%s`, “paid”, 100) => WHERE `amount`<100 and status=‘paid’ Wheref(`amount<%d and status=%s`, 100, “paid”) => WHERE `amount`<100 and status=‘paid’

​	Wheref 使用 fmt 构建条件字符串。Sprintf 和参数。请注意，如果 的数量 `args` 大于 `format` 中的占位符，则额外的 `args` 将用作模型的 where 条件参数。例如： Wheref（ `amount<? and status=%s` ， “paid”， 100） => WHERE `amount` <100 and status='paid' Wheref（ `amount<%d and status=%s` ， 100， “paid”） => WHERE `amount` <100 and status='paid'

### type WhereHolder <-2.1.0

```go
type WhereHolder struct {
	Type     string        // Type of this holder.
	Operator int           // Operator for this holder.
	Where    interface{}   // Where parameter, which can commonly be type of string/map/struct.
	Args     []interface{} // Arguments for where parameter.
	Prefix   string        // Field prefix, eg: "user.", "order.".
}
```

WhereHolder is the holder for where condition preparing.

​	WhereHolder 是 where 条件准备的持有者。