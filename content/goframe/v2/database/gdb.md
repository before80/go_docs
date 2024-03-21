+++
title = "gdb"
date = 2024-03-21T17:47:34+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/database/gdb

Package gdb provides ORM features for popular relationship databases.

TODO use context.Context as required parameter for all DB operations.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/database/gdb/gdb.go#L413)

``` go
const (
	InsertOperationInsert      = "INSERT"
	InsertOperationReplace     = "REPLACE"
	InsertOperationIgnore      = "INSERT IGNORE"
	InsertOnDuplicateKeyUpdate = "ON DUPLICATE KEY UPDATE"
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/database/gdb/gdb.go#L420)

``` go
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

``` go
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

``` go
const (
	DefaultGroupName = "default" // Default group name.
)
```

### Variables 

This section is empty.

### Functions 

##### func AddConfigNode 

``` go
func AddConfigNode(group string, node ConfigNode)
```

AddConfigNode adds one node configuration to configuration of given group.

##### func AddDefaultConfigGroup 

``` go
func AddDefaultConfigGroup(nodes ConfigGroup)
```

AddDefaultConfigGroup adds multiple node configurations to configuration of default group.

##### func AddDefaultConfigNode 

``` go
func AddDefaultConfigNode(node ConfigNode)
```

AddDefaultConfigNode adds one node configuration to configuration of default group.

##### func CatchSQL <-2.2.0

``` go
func CatchSQL(ctx context.Context, f func(ctx context.Context) error) (sqlArray []string, err error)
```

CatchSQL catches and returns all sql statements that are EXECUTED in given closure function. Be caution that, all the following sql statements should use the context object passing by function `f`.

##### func FormatMultiLineSqlToSingle <-2.6.4

``` go
func FormatMultiLineSqlToSingle(sqlTmp string) string
```

FormatMultiLineSqlToSingle formats sql template string into one line.

##### func FormatSqlWithArgs 

``` go
func FormatSqlWithArgs(sql string, args []interface{}) string
```

FormatSqlWithArgs binds the arguments to the sql string and returns a complete sql string, just for debugging.

##### func GetDefaultGroup 

``` go
func GetDefaultGroup() string
```

GetDefaultGroup returns the { name of default configuration.

##### func GetInsertOperationByOption 

``` go
func GetInsertOperationByOption(option InsertOption) string
```

GetInsertOperationByOption returns proper insert option with given parameter `option`.

##### func GetPrimaryKeyCondition 

``` go
func GetPrimaryKeyCondition(primary string, where ...interface{}) (newWhereCondition []interface{})
```

GetPrimaryKeyCondition returns a new where condition by primary field name. The optional parameter `where` is like follows: 123 => primary=123 []int{1, 2, 3} => primary IN(1,2,3) "john" => primary='john' []string{"john", "smith"} => primary IN('john','smith') g.Map{"id": g.Slice{1,2,3}} => id IN(1,2,3) g.Map{"id": 1, "name": "john"} => id=1 AND name='john' etc.

Note that it returns the given `where` parameter directly if the `primary` is empty or length of `where` > 1.

##### func IsConfigured 

``` go
func IsConfigured() bool
```

IsConfigured checks and returns whether the database configured. It returns true if any configuration exists.

##### func ListItemValues 

``` go
func ListItemValues(list interface{}, key interface{}, subKey ...interface{}) (values []interface{})
```

ListItemValues retrieves and returns the elements of all item struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice.

The parameter `list` supports types like: []map[string]interface{} []map[string]sub-map []struct []struct:sub-struct Note that the sub-map/sub-struct makes sense only if the optional parameter `subKey` is given. See gutil.ListItemValues.

##### func ListItemValuesUnique 

``` go
func ListItemValuesUnique(list interface{}, key string, subKey ...interface{}) []interface{}
```

ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`. Note that the parameter `list` should be type of slice which contains elements of map or struct, or else it returns an empty slice. See gutil.ListItemValuesUnique.

##### func MapOrStructToMapDeep <-2.6.0

``` go
func MapOrStructToMapDeep(value interface{}, omitempty bool) map[string]interface{}
```

MapOrStructToMapDeep converts `value` to map type recursively(if attribute struct is embedded). The parameter `value` should be type of *map/map/*struct/struct. It supports embedded struct definition for struct.

##### func Register 

``` go
func Register(name string, driver Driver) error
```

Register registers custom database driver to gdb.

##### func SetConfig 

``` go
func SetConfig(config Config)
```

SetConfig sets the global configuration for package. It will overwrite the old configuration of package.

##### func SetConfigGroup 

``` go
func SetConfigGroup(group string, nodes ConfigGroup)
```

SetConfigGroup sets the configuration for given group.

##### func SetDefaultGroup 

``` go
func SetDefaultGroup(name string)
```

SetDefaultGroup sets the group name for default configuration.

##### func ToSQL <-2.2.0

``` go
func ToSQL(ctx context.Context, f func(ctx context.Context) error) (sql string, err error)
```

ToSQL formats and returns the last one of sql statements in given closure function WITHOUT TRULY EXECUTING IT. Be caution that, all the following sql statements should use the context object passing by function `f`.

##### func WithDB <-2.0.5

``` go
func WithDB(ctx context.Context, db DB) context.Context
```

WithDB injects given db object into context and returns a new context.

##### func WithTX 

``` go
func WithTX(ctx context.Context, tx TX) context.Context
```

WithTX injects given transaction object into context and returns a new context.

### Types 

#### type CacheOption 

``` go
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

#### type CatchSQLManager <-2.2.0

``` go
type CatchSQLManager struct {
	SQLArray *garray.StrArray
	DoCommit bool // DoCommit marks it will be committed to underlying driver or not.
}
```

#### type ChunkHandler 

``` go
type ChunkHandler func(result Result, err error) bool
```

ChunkHandler is a function that is used in function Chunk, which handles given Result and error. It returns true if it wants to continue chunking, or else it returns false to stop chunking.

#### type Config 

``` go
type Config map[string]ConfigGroup
```

Config is the configuration management object.

#### type ConfigGroup 

``` go
type ConfigGroup []ConfigNode
```

ConfigGroup is a slice of configuration node for specified named group.

##### func GetConfig 

``` go
func GetConfig(group string) ConfigGroup
```

GetConfig retrieves and returns the configuration of given group.

#### type ConfigNode 

``` go
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

#### type Core 

``` go
type Core struct {
	// contains filtered or unexported fields
}
```

Core is the base struct for database management.

##### (*Core) Begin 

``` go
func (c *Core) Begin(ctx context.Context) (tx TX, err error)
```

Begin starts and returns the transaction object. You should call Commit or Rollback functions of the transaction object if you no longer use the transaction. Commit or Rollback functions will also close the transaction automatically.

##### (*Core) CheckLocalTypeForField <-2.1.3

``` go
func (c *Core) CheckLocalTypeForField(ctx context.Context, fieldType string, fieldValue interface{}) (LocalType, error)
```

CheckLocalTypeForField checks and returns corresponding type for given db type.

##### (*Core) ClearCache <-2.2.0

``` go
func (c *Core) ClearCache(ctx context.Context, table string) (err error)
```

ClearCache removes cached sql result of certain table.

##### (*Core) ClearCacheAll <-2.2.0

``` go
func (c *Core) ClearCacheAll(ctx context.Context) (err error)
```

ClearCacheAll removes all cached sql result from cache

##### (*Core) ClearTableFields <-2.2.0

``` go
func (c *Core) ClearTableFields(ctx context.Context, table string, schema ...string) (err error)
```

ClearTableFields removes certain cached table fields of current configuration group.

##### (*Core) ClearTableFieldsAll <-2.2.0

``` go
func (c *Core) ClearTableFieldsAll(ctx context.Context) (err error)
```

ClearTableFieldsAll removes all cached table fields of current configuration group.

##### (*Core) Close 

``` go
func (c *Core) Close(ctx context.Context) (err error)
```

Close closes the database and prevents new queries from starting. Close then waits for all queries that have started processing on the server to finish.

It is rare to Close a DB, as the DB handle is meant to be long-lived and shared between many goroutines.

##### (*Core) ConvertDataForRecord 

``` go
func (c *Core) ConvertDataForRecord(ctx context.Context, value interface{}, table string) (map[string]interface{}, error)
```

ConvertDataForRecord is a very important function, which does converting for any data that will be inserted into table/collection as a record.

The parameter `value` should be type of *map/map/*struct/struct. It supports embedded struct definition for struct.

##### (*Core) ConvertValueForField <-2.5.3

``` go
func (c *Core) ConvertValueForField(ctx context.Context, fieldType string, fieldValue interface{}) (interface{}, error)
```

ConvertValueForField converts value to the type of the record field. The parameter `fieldType` is the target record field. The parameter `fieldValue` is the value that to be committed to record field.

##### (*Core) ConvertValueForLocal <-2.1.2

``` go
func (c *Core) ConvertValueForLocal(
	ctx context.Context, fieldType string, fieldValue interface{},
) (interface{}, error)
```

ConvertValueForLocal converts value to local Golang type of value according field type name from database. The parameter `fieldType` is in lower case, like: `float(5,2)`, `unsigned double(5,2)`, `decimal(10,2)`, `char(45)`, `varchar(100)`, etc.

##### (*Core) Ctx 

``` go
func (c *Core) Ctx(ctx context.Context) DB
```

Ctx is a chaining function, which creates and returns a new DB that is a shallow copy of current DB object and with given context in it. Note that this returned DB object can be used only once, so do not assign it to a global or package variable for long using.

##### (*Core) Delete 

``` go
func (c *Core) Delete(ctx context.Context, table string, condition interface{}, args ...interface{}) (result sql.Result, err error)
```

Delete does "DELETE FROM ... " statement for the table.

The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc. It is commonly used with parameter `args`. Eg: "uid=10000", "uid", 10000 "money>? AND name like ?", 99999, "vip_%" "status IN (?)", g.Slice{1,2,3} "age IN(?,?)", 18, 50 User{ Id : 1, UserName : "john"}.

##### (*Core) DoCommit 

``` go
func (c *Core) DoCommit(ctx context.Context, in DoCommitInput) (out DoCommitOutput, err error)
```

DoCommit commits current sql and arguments to underlying sql driver.

##### (*Core) DoDelete 

``` go
func (c *Core) DoDelete(ctx context.Context, link Link, table string, condition string, args ...interface{}) (result sql.Result, err error)
```

DoDelete does "DELETE FROM ... " statement for the table. This function is usually used for custom interface definition, you do not need call it manually.

##### (*Core) DoExec 

``` go
func (c *Core) DoExec(ctx context.Context, link Link, sql string, args ...interface{}) (result sql.Result, err error)
```

DoExec commits the sql string and its arguments to underlying driver through given link object and returns the execution result.

##### (*Core) DoFilter 

``` go
func (c *Core) DoFilter(ctx context.Context, link Link, sql string, args []interface{}) (newSql string, newArgs []interface{}, err error)
```

DoFilter is a hook function, which filters the sql and its arguments before it's committed to underlying driver. The parameter `link` specifies the current database connection operation object. You can modify the sql string `sql` and its arguments `args` as you wish before they're committed to driver.

##### (*Core) DoInsert 

``` go
func (c *Core) DoInsert(ctx context.Context, link Link, table string, list List, option DoInsertOption) (result sql.Result, err error)
```

DoInsert inserts or updates data forF given table. This function is usually used for custom interface definition, you do not need call it manually. The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})

The parameter `option` values are as follows: InsertOptionDefault: just insert, if there's unique/primary key in the data, it returns error; InsertOptionReplace: if there's unique/primary key in the data, it deletes it from table and inserts a new one; InsertOptionSave: if there's unique/primary key in the data, it updates it or else inserts a new one; InsertOptionIgnore: if there's unique/primary key in the data, it ignores the inserting;

##### (*Core) DoPrepare 

``` go
func (c *Core) DoPrepare(ctx context.Context, link Link, sql string) (stmt *Stmt, err error)
```

DoPrepare calls prepare function on given link object and returns the statement object.

##### (*Core) DoQuery 

``` go
func (c *Core) DoQuery(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)
```

DoQuery commits the sql string and its arguments to underlying driver through given link object and returns the execution result.

##### (*Core) DoSelect <-2.0.5

``` go
func (c *Core) DoSelect(ctx context.Context, link Link, sql string, args ...interface{}) (result Result, err error)
```

DoSelect queries and returns data records from database.

##### (*Core) DoUpdate 

``` go
func (c *Core) DoUpdate(ctx context.Context, link Link, table string, data interface{}, condition string, args ...interface{}) (result sql.Result, err error)
```

DoUpdate does "UPDATE ... " statement for the table. This function is usually used for custom interface definition, you do not need to call it manually.

##### (*Core) Exec 

``` go
func (c *Core) Exec(ctx context.Context, sql string, args ...interface{}) (result sql.Result, err error)
```

Exec commits one query SQL to underlying driver and returns the execution result. It is most commonly used for data inserting and updating.

##### (*Core) FilteredLink <-2.1.0

``` go
func (c *Core) FilteredLink() string
```

FilteredLink retrieves and returns filtered `linkInfo` that can be using for logging or tracing purpose.

##### (*Core) FormatSqlBeforeExecuting <-2.5.2

``` go
func (c *Core) FormatSqlBeforeExecuting(sql string, args []interface{}) (newSql string, newArgs []interface{})
```

FormatSqlBeforeExecuting formats the sql string and its arguments before executing. The internal handleArguments function might be called twice during the SQL procedure, but do not worry about it, it's safe and efficient.

##### (*Core) FormatUpsert <-2.6.4

``` go
func (c *Core) FormatUpsert(columns []string, list List, option DoInsertOption) (string, error)
```

FormatUpsert formats and returns SQL clause part for upsert statement. In default implements, this function performs upsert statement for MySQL like: `INSERT INTO ... ON DUPLICATE KEY UPDATE x=VALUES(z),m=VALUES(y)...`

##### (*Core) GetAll 

``` go
func (c *Core) GetAll(ctx context.Context, sql string, args ...interface{}) (Result, error)
```

GetAll queries and returns data records from database.

##### (*Core) GetArray 

``` go
func (c *Core) GetArray(ctx context.Context, sql string, args ...interface{}) ([]Value, error)
```

GetArray queries and returns data values as slice from database. Note that if there are multiple columns in the result, it returns just one column values randomly.

##### (*Core) GetCache 

``` go
func (c *Core) GetCache() *gcache.Cache
```

GetCache returns the internal cache object.

##### (*Core) GetChars 

``` go
func (c *Core) GetChars() (charLeft string, charRight string)
```

GetChars returns the security char for current database. It does nothing in default.

##### (*Core) GetConfig 

``` go
func (c *Core) GetConfig() *ConfigNode
```

GetConfig returns the current used node configuration.

##### (*Core) GetCore 

``` go
func (c *Core) GetCore() *Core
```

GetCore returns the underlying *Core object.

##### (*Core) GetCount 

``` go
func (c *Core) GetCount(ctx context.Context, sql string, args ...interface{}) (int, error)
```

GetCount queries and returns the count from database.

##### (*Core) GetCtx 

``` go
func (c *Core) GetCtx() context.Context
```

GetCtx returns the context for current DB. It returns `context.Background()` is there's no context previously set.

##### (*Core) GetCtxTimeout 

``` go
func (c *Core) GetCtxTimeout(ctx context.Context, timeoutType int) (context.Context, context.CancelFunc)
```

GetCtxTimeout returns the context and cancel function for specified timeout type.

##### (*Core) GetDB <-2.2.0

``` go
func (c *Core) GetDB() DB
```

GetDB returns the underlying DB.

##### (*Core) GetDebug 

``` go
func (c *Core) GetDebug() bool
```

GetDebug returns the debug value.

##### (*Core) GetDryRun 

``` go
func (c *Core) GetDryRun() bool
```

GetDryRun returns the DryRun value.

##### (*Core) GetFieldType <-2.5.3

``` go
func (c *Core) GetFieldType(ctx context.Context, fieldName, table, schema string) *TableField
```

GetFieldType retrieves and returns the field type object for certain field by name.

##### (*Core) GetFieldTypeStr <-2.5.3

``` go
func (c *Core) GetFieldTypeStr(ctx context.Context, fieldName, table, schema string) string
```

GetFieldTypeStr retrieves and returns the field type string for certain field by name.

##### (*Core) GetGroup 

``` go
func (c *Core) GetGroup() string
```

GetGroup returns the group string configured.

##### (*Core) GetIgnoreResultFromCtx <-2.1.0

``` go
func (c *Core) GetIgnoreResultFromCtx(ctx context.Context) bool
```

##### (*Core) GetInternalCtxDataFromCtx <-2.1.0

``` go
func (c *Core) GetInternalCtxDataFromCtx(ctx context.Context) *internalCtxData
```

##### (*Core) GetLink <-2.0.5

``` go
func (c *Core) GetLink(ctx context.Context, master bool, schema string) (Link, error)
```

GetLink creates and returns the underlying database link object with transaction checks. The parameter `master` specifies whether using the master node if master-slave configured.

##### (*Core) GetLogger 

``` go
func (c *Core) GetLogger() glog.ILogger
```

GetLogger returns the (logger) of the orm.

##### (*Core) GetOne 

``` go
func (c *Core) GetOne(ctx context.Context, sql string, args ...interface{}) (Record, error)
```

GetOne queries and returns one record from database.

##### (*Core) GetPrefix 

``` go
func (c *Core) GetPrefix() string
```

GetPrefix returns the table prefix string configured.

##### (*Core) GetScan 

``` go
func (c *Core) GetScan(ctx context.Context, pointer interface{}, sql string, args ...interface{}) error
```

GetScan queries one or more records from database and converts them to given struct or struct array.

If parameter `pointer` is type of struct pointer, it calls GetStruct internally for the conversion. If parameter `pointer` is type of slice, it calls GetStructs internally for conversion.

##### (*Core) GetSchema 

``` go
func (c *Core) GetSchema() string
```

GetSchema returns the schema configured.

##### (*Core) GetTablesWithCache <-2.5.7

``` go
func (c *Core) GetTablesWithCache() ([]string, error)
```

GetTablesWithCache retrieves and returns the table names of current database with cache.

##### (*Core) GetValue 

``` go
func (c *Core) GetValue(ctx context.Context, sql string, args ...interface{}) (Value, error)
```

GetValue queries and returns the field value from database. The sql should query only one field from database, or else it returns only one field of the result.

##### (*Core) HasField 

``` go
func (c *Core) HasField(ctx context.Context, table, field string, schema ...string) (bool, error)
```

HasField determine whether the field exists in the table.

##### (*Core) HasTable 

``` go
func (c *Core) HasTable(name string) (bool, error)
```

HasTable determine whether the table name exists in the database.

##### (*Core) InjectIgnoreResult <-2.1.0

``` go
func (c *Core) InjectIgnoreResult(ctx context.Context) context.Context
```

##### (*Core) InjectInternalCtxData <-2.1.0

``` go
func (c *Core) InjectInternalCtxData(ctx context.Context) context.Context
```

##### (*Core) Insert 

``` go
func (c *Core) Insert(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)
```

Insert does "INSERT INTO ..." statement for the table. If there's already one unique record of the data in the table, it returns error.

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})

The parameter `batch` specifies the batch operation count when given data is slice.

##### (*Core) InsertAndGetId 

``` go
func (c *Core) InsertAndGetId(ctx context.Context, table string, data interface{}, batch ...int) (int64, error)
```

InsertAndGetId performs action Insert and returns the last insert id that automatically generated.

##### (*Core) InsertIgnore 

``` go
func (c *Core) InsertIgnore(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)
```

InsertIgnore does "INSERT IGNORE INTO ..." statement for the table. If there's already one unique record of the data in the table, it ignores the inserting.

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})

The parameter `batch` specifies the batch operation count when given data is slice.

##### (*Core) IsSoftCreatedFieldName <-2.6.4

``` go
func (c *Core) IsSoftCreatedFieldName(fieldName string) bool
```

IsSoftCreatedFieldName checks and returns whether given field name is an automatic-filled created time.

##### (Core) MarshalJSON 

``` go
func (c Core) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal. It just returns the pointer address.

Note that this interface implements mainly for workaround for a json infinite loop bug of Golang version < v1.14.

##### (*Core) Master 

``` go
func (c *Core) Master(schema ...string) (*sql.DB, error)
```

Master creates and returns a connection from master node if master-slave configured. It returns the default connection if master-slave not configured.

##### (*Core) MasterLink 

``` go
func (c *Core) MasterLink(schema ...string) (Link, error)
```

MasterLink acts like function Master but with additional `schema` parameter specifying the schema for the connection. It is defined for internal usage. Also see Master.

##### (*Core) Model 

``` go
func (c *Core) Model(tableNameQueryOrStruct ...interface{}) *Model
```

Model creates and returns a new ORM model from given schema. The parameter `tableNameQueryOrStruct` can be more than one table names, and also alias name, like:

1. Model names: db.Model("user") db.Model("user u") db.Model("user, user_detail") db.Model("user u, user_detail ud")
2. Model name with alias: db.Model("user", "u")
3. Model name with sub-query: db.Model("? AS a, ? AS b", subQuery1, subQuery2)

##### (*Core) PingMaster 

``` go
func (c *Core) PingMaster() error
```

PingMaster pings the master node to check authentication or keeps the connection alive.

##### (*Core) PingSlave 

``` go
func (c *Core) PingSlave() error
```

PingSlave pings the slave node to check authentication or keeps the connection alive.

##### (*Core) Prepare 

``` go
func (c *Core) Prepare(ctx context.Context, sql string, execOnMaster ...bool) (*Stmt, error)
```

Prepare creates a prepared statement for later queries or executions. Multiple queries or executions may be run concurrently from the returned statement. The caller must call the statement's Close method when the statement is no longer needed.

The parameter `execOnMaster` specifies whether executing the sql on master node, or else it executes the sql on slave node if master-slave configured.

##### (*Core) Query 

``` go
func (c *Core) Query(ctx context.Context, sql string, args ...interface{}) (result Result, err error)
```

Query commits one query SQL to underlying driver and returns the execution result. It is most commonly used for data querying.

##### (*Core) QuotePrefixTableName 

``` go
func (c *Core) QuotePrefixTableName(table string) string
```

QuotePrefixTableName adds prefix string and quotes chars for the table. It handles table string like: "user", "user u", "user,user_detail", "user u, user_detail ut", "user as u, user_detail as ut".

Note that, this will automatically checks the table prefix whether already added, if true it does nothing to the table name, or else adds the prefix to the table name.

##### (*Core) QuoteString 

``` go
func (c *Core) QuoteString(s string) string
```

QuoteString quotes string with quote chars. Strings like: "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".

The meaning of a `string` can be considered as part of a statement string including columns.

##### (*Core) QuoteWord 

``` go
func (c *Core) QuoteWord(s string) string
```

QuoteWord checks given string `s` a word, if true it quotes `s` with security chars of the database and returns the quoted string; or else it returns `s` without any change.

The meaning of a `word` can be considered as a column name.

##### (*Core) Raw 

``` go
func (c *Core) Raw(rawSql string, args ...interface{}) *Model
```

Raw creates and returns a model based on a raw sql not a table. Example:

```
db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
```

##### (*Core) Replace 

``` go
func (c *Core) Replace(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)
```

Replace does "REPLACE INTO ..." statement for the table. If there's already one unique record of the data in the table, it deletes the record and inserts a new one.

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. If given data is type of slice, it then does batch replacing, and the optional parameter `batch` specifies the batch operation count.

##### (*Core) RowsToResult 

``` go
func (c *Core) RowsToResult(ctx context.Context, rows *sql.Rows) (Result, error)
```

RowsToResult converts underlying data record type sql.Rows to Result type.

##### (*Core) Save 

``` go
func (c *Core) Save(ctx context.Context, table string, data interface{}, batch ...int) (sql.Result, error)
```

Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the table. It updates the record if there's primary or unique index in the saving data, or else it inserts a new record into the table.

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})

If given data is type of slice, it then does batch saving, and the optional parameter `batch` specifies the batch operation count.

##### (*Core) Schema 

``` go
func (c *Core) Schema(schema string) *Schema
```

Schema creates and returns a schema.

##### (*Core) SetDebug 

``` go
func (c *Core) SetDebug(debug bool)
```

SetDebug enables/disables the debug mode.

##### (*Core) SetDryRun 

``` go
func (c *Core) SetDryRun(enabled bool)
```

SetDryRun enables/disables the DryRun feature.

##### (*Core) SetLogger 

``` go
func (c *Core) SetLogger(logger glog.ILogger)
```

SetLogger sets the logger for orm.

##### (*Core) SetMaxConnLifeTime 

``` go
func (c *Core) SetMaxConnLifeTime(d time.Duration)
```

SetMaxConnLifeTime sets the maximum amount of time a connection may be reused.

Expired connections may be closed lazily before reuse.

If d <= 0, connections are not closed due to a connection's age.

##### (*Core) SetMaxIdleConnCount 

``` go
func (c *Core) SetMaxIdleConnCount(n int)
```

SetMaxIdleConnCount sets the maximum number of connections in the idle connection pool.

If MaxOpenConns is greater than 0 but less than the new MaxIdleConns, then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.

If n <= 0, no idle connections are retained.

The default max idle connections is currently 2. This may change in a future release.

##### (*Core) SetMaxOpenConnCount 

``` go
func (c *Core) SetMaxOpenConnCount(n int)
```

SetMaxOpenConnCount sets the maximum number of open connections to the database.

If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than MaxIdleConns, then MaxIdleConns will be reduced to match the new MaxOpenConns limit.

If n <= 0, then there is no limit on the number of open connections. The default is 0 (unlimited).

##### (*Core) Slave 

``` go
func (c *Core) Slave(schema ...string) (*sql.DB, error)
```

Slave creates and returns a connection from slave node if master-slave configured. It returns the default connection if master-slave not configured.

##### (*Core) SlaveLink 

``` go
func (c *Core) SlaveLink(schema ...string) (Link, error)
```

SlaveLink acts like function Slave but with additional `schema` parameter specifying the schema for the connection. It is defined for internal usage. Also see Slave.

##### (*Core) TableFields 

``` go
func (c *Core) TableFields(ctx context.Context, table string, schema ...string) (fields map[string]*TableField, err error)
```

TableFields retrieves and returns the fields' information of specified table of current schema.

The parameter `link` is optional, if given nil it automatically retrieves a raw sql connection as its link to proceed necessary sql query.

Note that it returns a map containing the field name and its corresponding fields. As a map is unsorted, the TableField struct has an "Index" field marks its sequence in the fields.

It's using cache feature to enhance the performance, which is never expired util the process restarts.

##### (*Core) Tables 

``` go
func (c *Core) Tables(ctx context.Context, schema ...string) (tables []string, err error)
```

Tables retrieves and returns the tables of current schema. It's mainly used in cli tool chain for automatically generating the models.

##### (*Core) Transaction 

``` go
func (c *Core) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error)
```

Transaction wraps the transaction logic using function `f`. It rollbacks the transaction and returns the error from function `f` if it returns non-nil error. It commits the transaction and returns nil if function `f` returns nil.

Note that, you should not Commit or Rollback the transaction in function `f` as it is automatically handled by this function.

##### (*Core) Union 

``` go
func (c *Core) Union(unions ...*Model) *Model
```

Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement.

##### (*Core) UnionAll 

``` go
func (c *Core) UnionAll(unions ...*Model) *Model
```

UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement.

##### (*Core) Update 

``` go
func (c *Core) Update(ctx context.Context, table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)
```

Update does "UPDATE ... " statement for the table.

The parameter `data` can be type of string/map/gmap/struct/*struct, etc. Eg: "uid=10000", "uid", 10000, g.Map{"uid": 10000, "name":"john"}

The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc. It is commonly used with parameter `args`. Eg: "uid=10000", "uid", 10000 "money>? AND name like ?", 99999, "vip_%" "status IN (?)", g.Slice{1,2,3} "age IN(?,?)", 18, 50 User{ Id : 1, UserName : "john"}.

##### (*Core) With 

``` go
func (c *Core) With(objects ...interface{}) *Model
```

With creates and returns an ORM model based on metadata of given object.

#### type Counter 

``` go
type Counter struct {
	Field string
	Value float64
}
```

Counter is the type for update count.

#### type DB 

``` go
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

##### func DBFromCtx <-2.0.5

``` go
func DBFromCtx(ctx context.Context) DB
```

DBFromCtx retrieves and returns DB object from context.

##### func Instance 

``` go
func Instance(name ...string) (db DB, err error)
```

Instance returns an instance for DB operations. The parameter `name` specifies the configuration group name, which is DefaultGroupName in default.

##### func New 

``` go
func New(node ConfigNode) (db DB, err error)
```

New creates and returns an ORM object with given configuration node.

##### func NewByGroup 

``` go
func NewByGroup(group ...string) (db DB, err error)
```

NewByGroup creates and returns an ORM object with global configurations. The parameter `name` specifies the configuration group name, which is DefaultGroupName in default.

#### type DoCommitInput 

``` go
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

#### type DoCommitOutput 

``` go
type DoCommitOutput struct {
	Result    sql.Result  // Result is the result of exec statement.
	Records   []Record    // Records is the result of query statement.
	Stmt      *Stmt       // Stmt is the Statement object result for Prepare.
	Tx        TX          // Tx is the transaction object result for Begin.
	RawResult interface{} // RawResult is the underlying result, which might be sql.Result/*sql.Rows/*sql.Row.
}
```

DoCommitOutput is the output parameters for function DoCommit.

#### type DoInsertOption 

``` go
type DoInsertOption struct {
	OnDuplicateStr string                 // Custom string for `on duplicated` statement.
	OnDuplicateMap map[string]interface{} // Custom key-value map from `OnDuplicateEx` function for `on duplicated` statement.
	OnConflict     []string               // Custom conflict key of upsert clause, if the database needs it.
	InsertOption   InsertOption           // Insert operation in constant value.
	BatchCount     int                    // Batch count for batch inserting.
}
```

DoInsertOption is the input struct for function DoInsert.

#### type Driver 

``` go
type Driver interface {
	// New creates and returns a database object for specified database server.
	New(core *Core, node *ConfigNode) (DB, error)
}
```

Driver is the interface for integrating sql drivers into package gdb.

#### type DriverDefault <-2.2.0

``` go
type DriverDefault struct {
	*Core
}
```

DriverDefault is the default driver for mysql database, which does nothing.

##### (*DriverDefault) New <-2.2.0

``` go
func (d *DriverDefault) New(core *Core, node *ConfigNode) (DB, error)
```

New creates and returns a database object for mysql. It implements the interface of gdb.Driver for extra database driver installation.

##### (*DriverDefault) Open <-2.2.0

``` go
func (d *DriverDefault) Open(config *ConfigNode) (db *sql.DB, err error)
```

Open creates and returns an underlying sql.DB object for mysql. Note that it converts time.Time argument to local timezone in default.

##### (*DriverDefault) PingMaster <-2.2.0

``` go
func (d *DriverDefault) PingMaster() error
```

PingMaster pings the master node to check authentication or keeps the connection alive.

##### (*DriverDefault) PingSlave <-2.2.0

``` go
func (d *DriverDefault) PingSlave() error
```

PingSlave pings the slave node to check authentication or keeps the connection alive.

#### type DriverWrapper <-2.2.0

``` go
type DriverWrapper struct {
	// contains filtered or unexported fields
}
```

DriverWrapper is a driver wrapper for extending features with embedded driver.

##### (*DriverWrapper) New <-2.2.0

``` go
func (d *DriverWrapper) New(core *Core, node *ConfigNode) (DB, error)
```

New creates and returns a database object for mysql. It implements the interface of gdb.Driver for extra database driver installation.

#### type DriverWrapperDB <-2.2.0

``` go
type DriverWrapperDB struct {
	DB
}
```

DriverWrapperDB is a DB wrapper for extending features with embedded DB.

##### (*DriverWrapperDB) DoInsert <-2.5.2

``` go
func (d *DriverWrapperDB) DoInsert(ctx context.Context, link Link, table string, list List, option DoInsertOption) (result sql.Result, err error)
```

DoInsert inserts or updates data forF given table. This function is usually used for custom interface definition, you do not need call it manually. The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})

The parameter `option` values are as follows: InsertOptionDefault: just insert, if there's unique/primary key in the data, it returns error; InsertOptionReplace: if there's unique/primary key in the data, it deletes it from table and inserts a new one; InsertOptionSave: if there's unique/primary key in the data, it updates it or else inserts a new one; InsertOptionIgnore: if there's unique/primary key in the data, it ignores the inserting;

##### (*DriverWrapperDB) Open <-2.2.0

``` go
func (d *DriverWrapperDB) Open(node *ConfigNode) (db *sql.DB, err error)
```

Open creates and returns an underlying sql.DB object for pgsql. https://pkg.go.dev/github.com/lib/pq

##### (*DriverWrapperDB) TableFields <-2.2.0

``` go
func (d *DriverWrapperDB) TableFields(
	ctx context.Context, table string, schema ...string,
) (fields map[string]*TableField, err error)
```

TableFields retrieves and returns the fields' information of specified table of current schema.

The parameter `link` is optional, if given nil it automatically retrieves a raw sql connection as its link to proceed necessary sql query.

Note that it returns a map containing the field name and its corresponding fields. As a map is unsorted, the TableField struct has an "Index" field marks its sequence in the fields.

It's using cache feature to enhance the performance, which is never expired util the process restarts.

##### (*DriverWrapperDB) Tables <-2.2.0

``` go
func (d *DriverWrapperDB) Tables(ctx context.Context, schema ...string) (tables []string, err error)
```

Tables retrieves and returns the tables of current schema. It's mainly used in cli tool chain for automatically generating the models.

#### type HookDeleteInput <-2.0.5

``` go
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

##### (*HookDeleteInput) Next <-2.0.5

``` go
func (h *HookDeleteInput) Next(ctx context.Context) (result sql.Result, err error)
```

Next calls the next hook handler.

#### type HookFuncDelete <-2.0.5

``` go
type HookFuncDelete func(ctx context.Context, in *HookDeleteInput) (result sql.Result, err error)
```

#### type HookFuncInsert <-2.0.5

``` go
type HookFuncInsert func(ctx context.Context, in *HookInsertInput) (result sql.Result, err error)
```

#### type HookFuncSelect <-2.0.5

``` go
type HookFuncSelect func(ctx context.Context, in *HookSelectInput) (result Result, err error)
```

#### type HookFuncUpdate <-2.0.5

``` go
type HookFuncUpdate func(ctx context.Context, in *HookUpdateInput) (result sql.Result, err error)
```

#### type HookHandler <-2.0.5

``` go
type HookHandler struct {
	Select HookFuncSelect
	Insert HookFuncInsert
	Update HookFuncUpdate
	Delete HookFuncDelete
}
```

HookHandler manages all supported hook functions for Model.

#### type HookInsertInput <-2.0.5

``` go
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

##### (*HookInsertInput) Next <-2.0.5

``` go
func (h *HookInsertInput) Next(ctx context.Context) (result sql.Result, err error)
```

Next calls the next hook handler.

#### type HookSelectInput <-2.0.5

``` go
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

##### (*HookSelectInput) Next <-2.0.5

``` go
func (h *HookSelectInput) Next(ctx context.Context) (result Result, err error)
```

Next calls the next hook handler.

#### type HookUpdateInput <-2.0.5

``` go
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

##### (*HookUpdateInput) Next <-2.0.5

``` go
func (h *HookUpdateInput) Next(ctx context.Context) (result sql.Result, err error)
```

Next calls the next hook handler.

#### type InsertOption <-2.5.0

``` go
type InsertOption int
const (
	InsertOptionDefault InsertOption = iota
	InsertOptionReplace
	InsertOptionSave
	InsertOptionIgnore
)
```

#### type Link 

``` go
type Link interface {
	QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, sql string) (*sql.Stmt, error)
	IsOnMaster() bool
	IsTransaction() bool
}
```

Link is a common database function wrapper interface. Note that, any operation using `Link` will have no SQL logging.

#### type List 

``` go
type List = []Map // List is type of map array.
```

#### type LocalType <-2.5.3

``` go
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

#### type Map 

``` go
type Map = map[string]interface{} // Map is alias of map[string]interface{}, which is the most common usage map type.
```

#### type Model 

``` go
type Model struct {
	// contains filtered or unexported fields
}
```

Model is core struct implementing the DAO for ORM.

##### (*Model) All 

``` go
func (m *Model) All(where ...interface{}) (Result, error)
```

All does "SELECT FROM ..." statement for the model. It retrieves the records from table and returns the result as slice type. It returns nil if there's no record retrieved with the given conditions from table.

The optional parameter `where` is the same as the parameter of Model.Where function, see Model.Where.

##### (*Model) AllAndCount <-2.4.2

``` go
func (m *Model) AllAndCount(useFieldForCount bool) (result Result, totalCount int, err error)
```

AllAndCount retrieves all records and the total count of records from the model. If useFieldForCount is true, it will use the fields specified in the model for counting; otherwise, it will use a constant value of 1 for counting. It returns the result as a slice of records, the total count of records, and an error if any. The where parameter is an optional list of conditions to use when retrieving records.

Example:

``` go
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

##### (*Model) Args 

``` go
func (m *Model) Args(args ...interface{}) *Model
```

Args sets custom arguments for model operation.

##### (*Model) Array 

``` go
func (m *Model) Array(fieldsAndWhere ...interface{}) ([]Value, error)
```

Array queries and returns data values as slice from database. Note that if there are multiple columns in the result, it returns just one column values randomly.

If the optional parameter `fieldsAndWhere` is given, the fieldsAndWhere[0] is the selected fields and fieldsAndWhere[1:] is treated as where condition fields. Also see Model.Fields and Model.Where functions.

##### (*Model) As 

``` go
func (m *Model) As(as string) *Model
```

As sets an alias name for current table.

##### (*Model) Avg 

``` go
func (m *Model) Avg(column string) (float64, error)
```

Avg does "SELECT AVG(x) FROM ..." statement for the model.

##### (*Model) Batch 

``` go
func (m *Model) Batch(batch int) *Model
```

Batch sets the batch operation number for the model.

##### (*Model) Builder <-2.1.0

``` go
func (m *Model) Builder() *WhereBuilder
```

Builder creates and returns a WhereBuilder. Please note that the builder is chain-safe.

##### (*Model) Cache 

``` go
func (m *Model) Cache(option CacheOption) *Model
```

Cache sets the cache feature for the model. It caches the result of the sql, which means if there's another same sql request, it just reads and returns the result from cache, it but not committed and executed into the database.

Note that, the cache feature is disabled if the model is performing select statement on a transaction.

##### (*Model) Chunk 

``` go
func (m *Model) Chunk(size int, handler ChunkHandler)
```

Chunk iterates the query result with given `size` and `handler` function.

##### (*Model) Clone 

``` go
func (m *Model) Clone() *Model
```

Clone creates and returns a new model which is a Clone of current model. Note that it uses deep-copy for the Clone.

##### (*Model) Count 

``` go
func (m *Model) Count(where ...interface{}) (int, error)
```

Count does "SELECT COUNT(x) FROM ..." statement for the model. The optional parameter `where` is the same as the parameter of Model.Where function, see Model.Where.

##### (*Model) CountColumn 

``` go
func (m *Model) CountColumn(column string) (int, error)
```

CountColumn does "SELECT COUNT(x) FROM ..." statement for the model.

##### (*Model) Ctx 

``` go
func (m *Model) Ctx(ctx context.Context) *Model
```

Ctx sets the context for current operation.

##### (*Model) DB 

``` go
func (m *Model) DB(db DB) *Model
```

DB sets/changes the db object for current operation.

##### (*Model) Data 

``` go
func (m *Model) Data(data ...interface{}) *Model
```

Data sets the operation data for the model. The parameter `data` can be type of string/map/gmap/slice/struct/*struct, etc. Note that, it uses shallow value copying for `data` if `data` is type of map/slice to avoid changing it inside function. Eg: Data("uid=10000") Data("uid", 10000) Data("uid=? AND name=?", 10000, "john") Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"}).

##### (*Model) Decrement 

``` go
func (m *Model) Decrement(column string, amount interface{}) (sql.Result, error)
```

Decrement decrements a column's value by a given amount. The parameter `amount` can be type of float or integer.

##### (*Model) Delete 

``` go
func (m *Model) Delete(where ...interface{}) (result sql.Result, err error)
```

Delete does "DELETE FROM ... " statement for the model. The optional parameter `where` is the same as the parameter of Model.Where function, see Model.Where.

##### (*Model) Distinct 

``` go
func (m *Model) Distinct() *Model
```

Distinct forces the query to only return distinct results.

##### (*Model) FieldAvg 

``` go
func (m *Model) FieldAvg(column string, as ...string) *Model
```

FieldAvg formats and appends commonly used field `AVG(column)` to the select fields of model.

##### (*Model) FieldCount 

``` go
func (m *Model) FieldCount(column string, as ...string) *Model
```

FieldCount formats and appends commonly used field `COUNT(column)` to the select fields of model.

##### (*Model) FieldMax 

``` go
func (m *Model) FieldMax(column string, as ...string) *Model
```

FieldMax formats and appends commonly used field `MAX(column)` to the select fields of model.

##### (*Model) FieldMin 

``` go
func (m *Model) FieldMin(column string, as ...string) *Model
```

FieldMin formats and appends commonly used field `MIN(column)` to the select fields of model.

##### (*Model) FieldSum 

``` go
func (m *Model) FieldSum(column string, as ...string) *Model
```

FieldSum formats and appends commonly used field `SUM(column)` to the select fields of model.

##### (*Model) Fields 

``` go
func (m *Model) Fields(fieldNamesOrMapStruct ...interface{}) *Model
```

Fields appends `fieldNamesOrMapStruct` to the operation fields of the model, multiple fields joined using char ','. The parameter `fieldNamesOrMapStruct` can be type of string/map/*map/struct/*struct.

Eg: Fields("id", "name", "age") Fields([]string{"id", "name", "age"}) Fields(map[string]interface{}{"id":1, "name":"john", "age":18}) Fields(User{ Id: 1, Name: "john", Age: 18}).

##### (*Model) FieldsEx 

``` go
func (m *Model) FieldsEx(fieldNamesOrMapStruct ...interface{}) *Model
```

FieldsEx appends `fieldNamesOrMapStruct` to the excluded operation fields of the model, multiple fields joined using char ','. Note that this function supports only single table operations. The parameter `fieldNamesOrMapStruct` can be type of string/map/*map/struct/*struct.

Also see Fields.

##### (*Model) FieldsExPrefix 

``` go
func (m *Model) FieldsExPrefix(prefixOrAlias string, fieldNamesOrMapStruct ...interface{}) *Model
```

FieldsExPrefix performs as function FieldsEx but add extra prefix for each field.

##### (*Model) FieldsPrefix 

``` go
func (m *Model) FieldsPrefix(prefixOrAlias string, fieldNamesOrMapStruct ...interface{}) *Model
```

FieldsPrefix performs as function Fields but add extra prefix for each field.

##### (*Model) GetCtx 

``` go
func (m *Model) GetCtx() context.Context
```

GetCtx returns the context for current Model. It returns `context.Background()` is there's no context previously set.

##### (*Model) GetFieldsExStr 

``` go
func (m *Model) GetFieldsExStr(fields string, prefix ...string) string
```

GetFieldsExStr retrieves and returns fields which are not in parameter `fields` from the table, joined with char ','. The parameter `fields` specifies the fields that are excluded. The optional parameter `prefix` specifies the prefix for each field, eg: FieldsExStr("id", "u.").

##### (*Model) GetFieldsStr 

``` go
func (m *Model) GetFieldsStr(prefix ...string) string
```

GetFieldsStr retrieves and returns all fields from the table, joined with char ','. The optional parameter `prefix` specifies the prefix for each field, eg: GetFieldsStr("u.").

##### (*Model) Group 

``` go
func (m *Model) Group(groupBy ...string) *Model
```

Group sets the "GROUP BY" statement for the model.

##### (*Model) Handler 

``` go
func (m *Model) Handler(handlers ...ModelHandler) *Model
```

Handler calls each of `handlers` on current Model and returns a new Model. ModelHandler is a function that handles given Model and returns a new Model that is custom modified.

##### (*Model) HasField 

``` go
func (m *Model) HasField(field string) (bool, error)
```

HasField determine whether the field exists in the table.

##### (*Model) Having 

``` go
func (m *Model) Having(having interface{}, args ...interface{}) *Model
```

Having sets the having statement for the model. The parameters of this function usage are as the same as function Where. See Where.

##### (*Model) Hook <-2.0.5

``` go
func (m *Model) Hook(hook HookHandler) *Model
```

Hook sets the hook functions for current model.

##### (*Model) Increment 

``` go
func (m *Model) Increment(column string, amount interface{}) (sql.Result, error)
```

Increment increments a column's value by a given amount. The parameter `amount` can be type of float or integer.

##### (*Model) InnerJoin 

``` go
func (m *Model) InnerJoin(tableOrSubQueryAndJoinConditions ...string) *Model
```

InnerJoin does "INNER JOIN ... ON ..." statement on the model. The parameter `table` can be joined table and its joined condition, and also with its alias name。

Eg: Model("user").InnerJoin("user_detail", "user_detail.uid=user.uid") Model("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid") Model("user", "u").InnerJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid").

##### (*Model) InnerJoinOnField 

``` go
func (m *Model) InnerJoinOnField(table, field string) *Model
```

InnerJoinOnField performs as InnerJoin, but it joins both tables with the `same field name`.

Eg: Model("order").InnerJoinOnField("user", "user_id") Model("order").InnerJoinOnField("product", "product_id").

##### (*Model) InnerJoinOnFields <-2.5.3

``` go
func (m *Model) InnerJoinOnFields(table, firstField, operator, secondField string) *Model
```

InnerJoinOnFields performs as InnerJoin. It specifies different fields and comparison operator.

Eg: Model("user").InnerJoinOnFields("order", "id", "=", "user_id") Model("user").InnerJoinOnFields("order", "id", ">", "user_id") Model("user").InnerJoinOnFields("order", "id", "<", "user_id")

##### (*Model) Insert 

``` go
func (m *Model) Insert(data ...interface{}) (result sql.Result, err error)
```

Insert does "INSERT INTO ..." statement for the model. The optional parameter `data` is the same as the parameter of Model.Data function, see Model.Data.

##### (*Model) InsertAndGetId 

``` go
func (m *Model) InsertAndGetId(data ...interface{}) (lastInsertId int64, err error)
```

InsertAndGetId performs action Insert and returns the last insert id that automatically generated.

##### (*Model) InsertIgnore 

``` go
func (m *Model) InsertIgnore(data ...interface{}) (result sql.Result, err error)
```

InsertIgnore does "INSERT IGNORE INTO ..." statement for the model. The optional parameter `data` is the same as the parameter of Model.Data function, see Model.Data.

##### (*Model) LeftJoin 

``` go
func (m *Model) LeftJoin(tableOrSubQueryAndJoinConditions ...string) *Model
```

LeftJoin does "LEFT JOIN ... ON ..." statement on the model. The parameter `table` can be joined table and its joined condition, and also with its alias name.

Eg: Model("user").LeftJoin("user_detail", "user_detail.uid=user.uid") Model("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid") Model("user", "u").LeftJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid").

##### (*Model) LeftJoinOnField 

``` go
func (m *Model) LeftJoinOnField(table, field string) *Model
```

LeftJoinOnField performs as LeftJoin, but it joins both tables with the `same field name`.

Eg: Model("order").LeftJoinOnField("user", "user_id") Model("order").LeftJoinOnField("product", "product_id").

##### (*Model) LeftJoinOnFields <-2.5.3

``` go
func (m *Model) LeftJoinOnFields(table, firstField, operator, secondField string) *Model
```

LeftJoinOnFields performs as LeftJoin. It specifies different fields and comparison operator.

Eg: Model("user").LeftJoinOnFields("order", "id", "=", "user_id") Model("user").LeftJoinOnFields("order", "id", ">", "user_id") Model("user").LeftJoinOnFields("order", "id", "<", "user_id")

##### (*Model) Limit 

``` go
func (m *Model) Limit(limit ...int) *Model
```

Limit sets the "LIMIT" statement for the model. The parameter `limit` can be either one or two number, if passed two number is passed, it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]" statement.

##### (*Model) LockShared 

``` go
func (m *Model) LockShared() *Model
```

LockShared sets the lock in share mode for current operation.

##### (*Model) LockUpdate 

``` go
func (m *Model) LockUpdate() *Model
```

LockUpdate sets the lock for update for current operation.

##### (*Model) Master 

``` go
func (m *Model) Master() *Model
```

Master marks the following operation on master node.

##### (*Model) Max 

``` go
func (m *Model) Max(column string) (float64, error)
```

Max does "SELECT MAX(x) FROM ..." statement for the model.

##### (*Model) Min 

``` go
func (m *Model) Min(column string) (float64, error)
```

Min does "SELECT MIN(x) FROM ..." statement for the model.

##### (*Model) Offset 

``` go
func (m *Model) Offset(offset int) *Model
```

Offset sets the "OFFSET" statement for the model. It only makes sense for some databases like SQLServer, PostgreSQL, etc.

##### (*Model) OmitEmpty 

``` go
func (m *Model) OmitEmpty() *Model
```

OmitEmpty sets optionOmitEmpty option for the model, which automatically filers the data and where parameters for `empty` values.

##### (*Model) OmitEmptyData 

``` go
func (m *Model) OmitEmptyData() *Model
```

OmitEmptyData sets optionOmitEmptyData option for the model, which automatically filers the Data parameters for `empty` values.

##### (*Model) OmitEmptyWhere 

``` go
func (m *Model) OmitEmptyWhere() *Model
```

OmitEmptyWhere sets optionOmitEmptyWhere option for the model, which automatically filers the Where/Having parameters for `empty` values.

Eg:

```
Where("id", []int{}).All()             -> SELECT xxx FROM xxx WHERE 0=1
Where("name", "").All()                -> SELECT xxx FROM xxx WHERE `name`=''
OmitEmpty().Where("id", []int{}).All() -> SELECT xxx FROM xxx
OmitEmpty().("name", "").All()         -> SELECT xxx FROM xxx.
```

##### (*Model) OmitNil 

``` go
func (m *Model) OmitNil() *Model
```

OmitNil sets optionOmitNil option for the model, which automatically filers the data and where parameters for `nil` values.

##### (*Model) OmitNilData 

``` go
func (m *Model) OmitNilData() *Model
```

OmitNilData sets optionOmitNilData option for the model, which automatically filers the Data parameters for `nil` values.

##### (*Model) OmitNilWhere 

``` go
func (m *Model) OmitNilWhere() *Model
```

OmitNilWhere sets optionOmitNilWhere option for the model, which automatically filers the Where/Having parameters for `nil` values.

##### (*Model) OnConflict <-2.6.4

``` go
func (m *Model) OnConflict(onConflict ...interface{}) *Model
```

OnConflict sets the primary key or index when columns conflicts occurs. It's not necessary for MySQL driver.

##### (*Model) OnDuplicate 

``` go
func (m *Model) OnDuplicate(onDuplicate ...interface{}) *Model
```

OnDuplicate sets the operations when columns conflicts occurs. In MySQL, this is used for "ON DUPLICATE KEY UPDATE" statement. In PgSQL, this is used for "ON CONFLICT (id) DO UPDATE SET" statement. The parameter `onDuplicate` can be type of string/Raw/*Raw/map/slice. Example:

OnDuplicate("nickname, age") OnDuplicate("nickname", "age")

```
OnDuplicate(g.Map{
	  "nickname": gdb.Raw("CONCAT('name_', VALUES(`nickname`))"),
})

OnDuplicate(g.Map{
	  "nickname": "passport",
}).
```

##### (*Model) OnDuplicateEx 

``` go
func (m *Model) OnDuplicateEx(onDuplicateEx ...interface{}) *Model
```

OnDuplicateEx sets the excluding columns for operations when columns conflict occurs. In MySQL, this is used for "ON DUPLICATE KEY UPDATE" statement. In PgSQL, this is used for "ON CONFLICT (id) DO UPDATE SET" statement. The parameter `onDuplicateEx` can be type of string/map/slice. Example:

OnDuplicateEx("passport, password") OnDuplicateEx("passport", "password")

```
OnDuplicateEx(g.Map{
	  "passport": "",
	  "password": "",
}).
```

##### (*Model) One 

``` go
func (m *Model) One(where ...interface{}) (Record, error)
```

One retrieves one record from table and returns the result as map type. It returns nil if there's no record retrieved with the given conditions from table.

The optional parameter `where` is the same as the parameter of Model.Where function, see Model.Where.

##### (*Model) Order 

``` go
func (m *Model) Order(orderBy ...interface{}) *Model
```

Order sets the "ORDER BY" statement for the model.

Eg: Order("id desc") Order("id", "desc"). Order("id desc,name asc") Order("id desc").Order("name asc") Order(gdb.Raw("field(id, 3,1,2)")).

##### (*Model) OrderAsc 

``` go
func (m *Model) OrderAsc(column string) *Model
```

OrderAsc sets the "ORDER BY xxx ASC" statement for the model.

##### (*Model) OrderDesc 

``` go
func (m *Model) OrderDesc(column string) *Model
```

OrderDesc sets the "ORDER BY xxx DESC" statement for the model.

##### (*Model) OrderRandom 

``` go
func (m *Model) OrderRandom() *Model
```

OrderRandom sets the "ORDER BY RANDOM()" statement for the model.

##### (*Model) Page 

``` go
func (m *Model) Page(page, limit int) *Model
```

Page sets the paging number for the model. The parameter `page` is started from 1 for paging. Note that, it differs that the Limit function starts from 0 for "LIMIT" statement.

##### (*Model) Partition <-2.5.5

``` go
func (m *Model) Partition(partitions ...string) *Model
```

Partition sets Partition name. Example: dao.User.Ctx(ctx).Partition（"p1","p2","p3").All()

##### (*Model) QuoteWord 

``` go
func (m *Model) QuoteWord(s string) string
```

QuoteWord checks given string `s` a word, if true it quotes `s` with security chars of the database and returns the quoted string; or else it returns `s` without any change.

The meaning of a `word` can be considered as a column name.

##### (*Model) Raw 

``` go
func (m *Model) Raw(rawSql string, args ...interface{}) *Model
```

Raw sets current model as a raw sql model. Example:

```
db.Raw("SELECT * FROM `user` WHERE `name` = ?", "john").Scan(&result)
```

See Core.Raw.

##### (*Model) Replace 

``` go
func (m *Model) Replace(data ...interface{}) (result sql.Result, err error)
```

Replace does "REPLACE INTO ..." statement for the model. The optional parameter `data` is the same as the parameter of Model.Data function, see Model.Data.

##### (*Model) RightJoin 

``` go
func (m *Model) RightJoin(tableOrSubQueryAndJoinConditions ...string) *Model
```

RightJoin does "RIGHT JOIN ... ON ..." statement on the model. The parameter `table` can be joined table and its joined condition, and also with its alias name.

Eg: Model("user").RightJoin("user_detail", "user_detail.uid=user.uid") Model("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid") Model("user", "u").RightJoin("SELECT xxx FROM xxx","a", "a.uid=u.uid").

##### (*Model) RightJoinOnField 

``` go
func (m *Model) RightJoinOnField(table, field string) *Model
```

RightJoinOnField performs as RightJoin, but it joins both tables with the `same field name`.

Eg: Model("order").InnerJoinOnField("user", "user_id") Model("order").InnerJoinOnField("product", "product_id").

##### (*Model) RightJoinOnFields <-2.5.3

``` go
func (m *Model) RightJoinOnFields(table, firstField, operator, secondField string) *Model
```

RightJoinOnFields performs as RightJoin. It specifies different fields and comparison operator.

Eg: Model("user").RightJoinOnFields("order", "id", "=", "user_id") Model("user").RightJoinOnFields("order", "id", ">", "user_id") Model("user").RightJoinOnFields("order", "id", "<", "user_id")

##### (*Model) Safe 

``` go
func (m *Model) Safe(safe ...bool) *Model
```

Safe marks this model safe or unsafe. If safe is true, it clones and returns a new model object whenever the operation done, or else it changes the attribute of current model.

##### (*Model) Save 

``` go
func (m *Model) Save(data ...interface{}) (result sql.Result, err error)
```

Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the model. The optional parameter `data` is the same as the parameter of Model.Data function, see Model.Data.

It updates the record if there's primary or unique index in the saving data, or else it inserts a new record into the table.

##### (*Model) Scan 

``` go
func (m *Model) Scan(pointer interface{}, where ...interface{}) error
```

Scan automatically calls Struct or Structs function according to the type of parameter `pointer`. It calls function doStruct if `pointer` is type of *struct/**struct. It calls function doStructs if `pointer` is type of *[]struct/*[]*struct.

The optional parameter `where` is the same as the parameter of Model.Where function, see Model.Where.

Note that it returns sql.ErrNoRows if the given parameter `pointer` pointed to a variable that has default value and there's no record retrieved with the given conditions from table.

Example: user := new(User) err := db.Model("user").Where("id", 1).Scan(user)

user := (*User)(nil) err := db.Model("user").Where("id", 1).Scan(&user)

users := ([]User)(nil) err := db.Model("user").Scan(&users)

users := ([]*User)(nil) err := db.Model("user").Scan(&users).

##### (*Model) ScanAndCount <-2.4.2

``` go
func (m *Model) ScanAndCount(pointer interface{}, totalCount *int, useFieldForCount bool) (err error)
```

ScanAndCount scans a single record or record array that matches the given conditions and counts the total number of records that match those conditions. If useFieldForCount is true, it will use the fields specified in the model for counting; The pointer parameter is a pointer to a struct that the scanned data will be stored in. The pointerCount parameter is a pointer to an integer that will be set to the total number of records that match the given conditions. The where parameter is an optional list of conditions to use when retrieving records.

Example:

``` go
var count int
user := new(User)
err  := db.Model("user").Where("id", 1).ScanAndCount(user,&count,true)
fmt.Println(user, count)
```

Example Join:

``` go
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

##### (*Model) ScanList 

``` go
func (m *Model) ScanList(structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error)
```

ScanList converts `r` to struct slice which contains other complex struct attributes. Note that the parameter `listPointer` should be type of *[]struct/*[]*struct.

See Result.ScanList.

##### (*Model) Schema 

``` go
func (m *Model) Schema(schema string) *Model
```

Schema sets the schema for current operation.

##### (*Model) Slave 

``` go
func (m *Model) Slave() *Model
```

Slave marks the following operation on slave node. Note that it makes sense only if there's any slave node configured.

##### (*Model) SoftTime <-2.6.3

``` go
func (m *Model) SoftTime(option SoftTimeOption) *Model
```

SoftTime sets the SoftTimeOption to customize soft time feature for Model.

##### (*Model) Sum 

``` go
func (m *Model) Sum(column string) (float64, error)
```

Sum does "SELECT SUM(x) FROM ..." statement for the model.

##### (*Model) TX 

``` go
func (m *Model) TX(tx TX) *Model
```

TX sets/changes the transaction for current operation.

##### (*Model) TableFields 

``` go
func (m *Model) TableFields(tableStr string, schema ...string) (fields map[string]*TableField, err error)
```

TableFields retrieves and returns the fields' information of specified table of current schema.

Also see DriverMysql.TableFields.

##### (*Model) Transaction 

``` go
func (m *Model) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error)
```

Transaction wraps the transaction logic using function `f`. It rollbacks the transaction and returns the error from function `f` if it returns non-nil error. It commits the transaction and returns nil if function `f` returns nil.

Note that, you should not Commit or Rollback the transaction in function `f` as it is automatically handled by this function.

##### (*Model) Union 

``` go
func (m *Model) Union(unions ...*Model) *Model
```

Union does "(SELECT xxx FROM xxx) UNION (SELECT xxx FROM xxx) ..." statement for the model.

##### (*Model) UnionAll 

``` go
func (m *Model) UnionAll(unions ...*Model) *Model
```

UnionAll does "(SELECT xxx FROM xxx) UNION ALL (SELECT xxx FROM xxx) ..." statement for the model.

##### (*Model) Unscoped 

``` go
func (m *Model) Unscoped() *Model
```

Unscoped disables the soft time feature for insert, update and delete operations.

##### (*Model) Update 

``` go
func (m *Model) Update(dataAndWhere ...interface{}) (result sql.Result, err error)
```

Update does "UPDATE ... " statement for the model.

If the optional parameter `dataAndWhere` is given, the dataAndWhere[0] is the updated data field, and dataAndWhere[1:] is treated as where condition fields. Also see Model.Data and Model.Where functions.

##### (*Model) UpdateAndGetAffected <-2.1.0

``` go
func (m *Model) UpdateAndGetAffected(dataAndWhere ...interface{}) (affected int64, err error)
```

UpdateAndGetAffected performs update statement and returns the affected rows number.

##### (*Model) Value 

``` go
func (m *Model) Value(fieldsAndWhere ...interface{}) (Value, error)
```

Value retrieves a specified record value from table and returns the result as interface type. It returns nil if there's no record found with the given conditions from table.

If the optional parameter `fieldsAndWhere` is given, the fieldsAndWhere[0] is the selected fields and fieldsAndWhere[1:] is treated as where condition fields. Also see Model.Fields and Model.Where functions.

##### (*Model) Where 

``` go
func (m *Model) Where(where interface{}, args ...interface{}) *Model
```

Where sets the condition statement for the builder. The parameter `where` can be type of string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times, multiple conditions will be joined into where statement using "AND". See WhereBuilder.Where.

##### (*Model) WhereBetween 

``` go
func (m *Model) WhereBetween(column string, min, max interface{}) *Model
```

WhereBetween builds `column BETWEEN min AND max` statement. See WhereBuilder.WhereBetween.

##### (*Model) WhereGT 

``` go
func (m *Model) WhereGT(column string, value interface{}) *Model
```

WhereGT builds `column > value` statement. See WhereBuilder.WhereGT.

##### (*Model) WhereGTE 

``` go
func (m *Model) WhereGTE(column string, value interface{}) *Model
```

WhereGTE builds `column >= value` statement. See WhereBuilder.WhereGTE.

##### (*Model) WhereIn 

``` go
func (m *Model) WhereIn(column string, in interface{}) *Model
```

WhereIn builds `column IN (in)` statement. See WhereBuilder.WhereIn.

##### (*Model) WhereLT 

``` go
func (m *Model) WhereLT(column string, value interface{}) *Model
```

WhereLT builds `column < value` statement. See WhereBuilder.WhereLT.

##### (*Model) WhereLTE 

``` go
func (m *Model) WhereLTE(column string, value interface{}) *Model
```

WhereLTE builds `column <= value` statement. See WhereBuilder.WhereLTE.

##### (*Model) WhereLike 

``` go
func (m *Model) WhereLike(column string, like string) *Model
```

WhereLike builds `column LIKE like` statement. See WhereBuilder.WhereLike.

##### (*Model) WhereNot 

``` go
func (m *Model) WhereNot(column string, value interface{}) *Model
```

WhereNot builds `column != value` statement. See WhereBuilder.WhereNot.

##### (*Model) WhereNotBetween 

``` go
func (m *Model) WhereNotBetween(column string, min, max interface{}) *Model
```

WhereNotBetween builds `column NOT BETWEEN min AND max` statement. See WhereBuilder.WhereNotBetween.

##### (*Model) WhereNotIn 

``` go
func (m *Model) WhereNotIn(column string, in interface{}) *Model
```

WhereNotIn builds `column NOT IN (in)` statement. See WhereBuilder.WhereNotIn.

##### (*Model) WhereNotLike 

``` go
func (m *Model) WhereNotLike(column string, like interface{}) *Model
```

WhereNotLike builds `column NOT LIKE like` statement. See WhereBuilder.WhereNotLike.

##### (*Model) WhereNotNull 

``` go
func (m *Model) WhereNotNull(columns ...string) *Model
```

WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement. See WhereBuilder.WhereNotNull.

##### (*Model) WhereNull 

``` go
func (m *Model) WhereNull(columns ...string) *Model
```

WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement. See WhereBuilder.WhereNull.

##### (*Model) WhereOr 

``` go
func (m *Model) WhereOr(where interface{}, args ...interface{}) *Model
```

WhereOr adds "OR" condition to the where statement. See WhereBuilder.WhereOr.

##### (*Model) WhereOrBetween 

``` go
func (m *Model) WhereOrBetween(column string, min, max interface{}) *Model
```

WhereOrBetween builds `column BETWEEN min AND max` statement in `OR` conditions. See WhereBuilder.WhereOrBetween.

##### (*Model) WhereOrGT 

``` go
func (m *Model) WhereOrGT(column string, value interface{}) *Model
```

WhereOrGT builds `column > value` statement in `OR` conditions. See WhereBuilder.WhereOrGT.

##### (*Model) WhereOrGTE 

``` go
func (m *Model) WhereOrGTE(column string, value interface{}) *Model
```

WhereOrGTE builds `column >= value` statement in `OR` conditions. See WhereBuilder.WhereOrGTE.

##### (*Model) WhereOrIn 

``` go
func (m *Model) WhereOrIn(column string, in interface{}) *Model
```

WhereOrIn builds `column IN (in)` statement in `OR` conditions. See WhereBuilder.WhereOrIn.

##### (*Model) WhereOrLT 

``` go
func (m *Model) WhereOrLT(column string, value interface{}) *Model
```

WhereOrLT builds `column < value` statement in `OR` conditions. See WhereBuilder.WhereOrLT.

##### (*Model) WhereOrLTE 

``` go
func (m *Model) WhereOrLTE(column string, value interface{}) *Model
```

WhereOrLTE builds `column <= value` statement in `OR` conditions. See WhereBuilder.WhereOrLTE.

##### (*Model) WhereOrLike 

``` go
func (m *Model) WhereOrLike(column string, like interface{}) *Model
```

WhereOrLike builds `column LIKE like` statement in `OR` conditions. See WhereBuilder.WhereOrLike.

##### (*Model) WhereOrNot <-2.4.2

``` go
func (m *Model) WhereOrNot(column string, value interface{}) *Model
```

WhereOrNot builds `column != value` statement. See WhereBuilder.WhereOrNot.

##### (*Model) WhereOrNotBetween 

``` go
func (m *Model) WhereOrNotBetween(column string, min, max interface{}) *Model
```

WhereOrNotBetween builds `column NOT BETWEEN min AND max` statement in `OR` conditions. See WhereBuilder.WhereOrNotBetween.

##### (*Model) WhereOrNotIn 

``` go
func (m *Model) WhereOrNotIn(column string, in interface{}) *Model
```

WhereOrNotIn builds `column NOT IN (in)` statement. See WhereBuilder.WhereOrNotIn.

##### (*Model) WhereOrNotLike 

``` go
func (m *Model) WhereOrNotLike(column string, like interface{}) *Model
```

WhereOrNotLike builds `column NOT LIKE 'like'` statement in `OR` conditions. See WhereBuilder.WhereOrNotLike.

##### (*Model) WhereOrNotNull 

``` go
func (m *Model) WhereOrNotNull(columns ...string) *Model
```

WhereOrNotNull builds `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` statement in `OR` conditions. See WhereBuilder.WhereOrNotNull.

##### (*Model) WhereOrNull 

``` go
func (m *Model) WhereOrNull(columns ...string) *Model
```

WhereOrNull builds `columns[0] IS NULL OR columns[1] IS NULL ...` statement in `OR` conditions. See WhereBuilder.WhereOrNull.

##### (*Model) WhereOrPrefix 

``` go
func (m *Model) WhereOrPrefix(prefix string, where interface{}, args ...interface{}) *Model
```

WhereOrPrefix performs as WhereOr, but it adds prefix to each field in where statement. See WhereBuilder.WhereOrPrefix.

##### (*Model) WhereOrPrefixBetween 

``` go
func (m *Model) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *Model
```

WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixBetween.

##### (*Model) WhereOrPrefixGT 

``` go
func (m *Model) WhereOrPrefixGT(prefix string, column string, value interface{}) *Model
```

WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixGT.

##### (*Model) WhereOrPrefixGTE 

``` go
func (m *Model) WhereOrPrefixGTE(prefix string, column string, value interface{}) *Model
```

WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixGTE.

##### (*Model) WhereOrPrefixIn 

``` go
func (m *Model) WhereOrPrefixIn(prefix string, column string, in interface{}) *Model
```

WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixIn.

##### (*Model) WhereOrPrefixLT 

``` go
func (m *Model) WhereOrPrefixLT(prefix string, column string, value interface{}) *Model
```

WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixLT.

##### (*Model) WhereOrPrefixLTE 

``` go
func (m *Model) WhereOrPrefixLTE(prefix string, column string, value interface{}) *Model
```

WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixLTE.

##### (*Model) WhereOrPrefixLike 

``` go
func (m *Model) WhereOrPrefixLike(prefix string, column string, like interface{}) *Model
```

WhereOrPrefixLike builds `prefix.column LIKE like` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixLike.

##### (*Model) WhereOrPrefixNot <-2.4.2

``` go
func (m *Model) WhereOrPrefixNot(prefix string, column string, value interface{}) *Model
```

WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixNot.

##### (*Model) WhereOrPrefixNotBetween 

``` go
func (m *Model) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *Model
```

WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixNotBetween.

##### (*Model) WhereOrPrefixNotIn 

``` go
func (m *Model) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *Model
```

WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement. See WhereBuilder.WhereOrPrefixNotIn.

##### (*Model) WhereOrPrefixNotLike 

``` go
func (m *Model) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *Model
```

WhereOrPrefixNotLike builds `prefix.column NOT LIKE like` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixNotLike.

##### (*Model) WhereOrPrefixNotNull 

``` go
func (m *Model) WhereOrPrefixNotNull(prefix string, columns ...string) *Model
```

WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixNotNull.

##### (*Model) WhereOrPrefixNull 

``` go
func (m *Model) WhereOrPrefixNull(prefix string, columns ...string) *Model
```

WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` statement in `OR` conditions. See WhereBuilder.WhereOrPrefixNull.

##### (*Model) WhereOrf 

``` go
func (m *Model) WhereOrf(format string, args ...interface{}) *Model
```

WhereOrf builds `OR` condition string using fmt.Sprintf and arguments. See WhereBuilder.WhereOrf.

##### (*Model) WherePrefix 

``` go
func (m *Model) WherePrefix(prefix string, where interface{}, args ...interface{}) *Model
```

WherePrefix performs as Where, but it adds prefix to each field in where statement. See WhereBuilder.WherePrefix.

##### (*Model) WherePrefixBetween 

``` go
func (m *Model) WherePrefixBetween(prefix string, column string, min, max interface{}) *Model
```

WherePrefixBetween builds `prefix.column BETWEEN min AND max` statement. See WhereBuilder.WherePrefixBetween.

##### (*Model) WherePrefixGT 

``` go
func (m *Model) WherePrefixGT(prefix string, column string, value interface{}) *Model
```

WherePrefixGT builds `prefix.column > value` statement. See WhereBuilder.WherePrefixGT.

##### (*Model) WherePrefixGTE 

``` go
func (m *Model) WherePrefixGTE(prefix string, column string, value interface{}) *Model
```

WherePrefixGTE builds `prefix.column >= value` statement. See WhereBuilder.WherePrefixGTE.

##### (*Model) WherePrefixIn 

``` go
func (m *Model) WherePrefixIn(prefix string, column string, in interface{}) *Model
```

WherePrefixIn builds `prefix.column IN (in)` statement. See WhereBuilder.WherePrefixIn.

##### (*Model) WherePrefixLT 

``` go
func (m *Model) WherePrefixLT(prefix string, column string, value interface{}) *Model
```

WherePrefixLT builds `prefix.column < value` statement. See WhereBuilder.WherePrefixLT.

##### (*Model) WherePrefixLTE 

``` go
func (m *Model) WherePrefixLTE(prefix string, column string, value interface{}) *Model
```

WherePrefixLTE builds `prefix.column <= value` statement. See WhereBuilder.WherePrefixLTE.

##### (*Model) WherePrefixLike 

``` go
func (m *Model) WherePrefixLike(prefix string, column string, like interface{}) *Model
```

WherePrefixLike builds `prefix.column LIKE like` statement. See WhereBuilder.WherePrefixLike.

##### (*Model) WherePrefixNot 

``` go
func (m *Model) WherePrefixNot(prefix string, column string, value interface{}) *Model
```

WherePrefixNot builds `prefix.column != value` statement. See WhereBuilder.WherePrefixNot.

##### (*Model) WherePrefixNotBetween 

``` go
func (m *Model) WherePrefixNotBetween(prefix string, column string, min, max interface{}) *Model
```

WherePrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement. See WhereBuilder.WherePrefixNotBetween.

##### (*Model) WherePrefixNotIn 

``` go
func (m *Model) WherePrefixNotIn(prefix string, column string, in interface{}) *Model
```

WherePrefixNotIn builds `prefix.column NOT IN (in)` statement. See WhereBuilder.WherePrefixNotIn.

##### (*Model) WherePrefixNotLike 

``` go
func (m *Model) WherePrefixNotLike(prefix string, column string, like interface{}) *Model
```

WherePrefixNotLike builds `prefix.column NOT LIKE like` statement. See WhereBuilder.WherePrefixNotLike.

##### (*Model) WherePrefixNotNull 

``` go
func (m *Model) WherePrefixNotNull(prefix string, columns ...string) *Model
```

WherePrefixNotNull builds `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` statement. See WhereBuilder.WherePrefixNotNull.

##### (*Model) WherePrefixNull 

``` go
func (m *Model) WherePrefixNull(prefix string, columns ...string) *Model
```

WherePrefixNull builds `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` statement. See WhereBuilder.WherePrefixNull.

##### (*Model) WherePri 

``` go
func (m *Model) WherePri(where interface{}, args ...interface{}) *Model
```

WherePri does the same logic as Model.Where except that if the parameter `where` is a single condition like int/string/float/slice, it treats the condition as the primary key value. That is, if primary key is "id" and given `where` parameter as "123", the WherePri function treats the condition as "id=123", but Model.Where treats the condition as string "123". See WhereBuilder.WherePri.

##### (*Model) Wheref 

``` go
func (m *Model) Wheref(format string, args ...interface{}) *Model
```

Wheref builds condition string using fmt.Sprintf and arguments. Note that if the number of `args` is more than the placeholder in `format`, the extra `args` will be used as the where condition arguments of the Model. See WhereBuilder.Wheref.

##### (*Model) With 

``` go
func (m *Model) With(objects ...interface{}) *Model
```

With creates and returns an ORM model based on metadata of given object. It also enables model association operations feature on given `object`. It can be called multiple times to add one or more objects to model and enable their mode association operations feature. For example, if given struct definition:

``` go
type User struct {
	 gmeta.Meta `orm:"table:user"`
	 Id         int           `json:"id"`
	 Name       string        `json:"name"`
	 UserDetail *UserDetail   `orm:"with:uid=id"`
	 UserScores []*UserScores `orm:"with:uid=id"`
}
```

We can enable model association operations on attribute `UserDetail` and `UserScores` by:

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

##### (*Model) WithAll 

``` go
func (m *Model) WithAll() *Model
```

WithAll enables model association operations on all objects that have "with" tag in the struct.

#### type ModelHandler 

``` go
type ModelHandler func(m *Model) *Model
```

ModelHandler is a function that handles given Model and returns a new Model that is custom modified.

#### type Raw 

``` go
type Raw string // Raw is a raw sql that will not be treated as argument but as a direct sql part.
```

#### type Record 

``` go
type Record map[string]Value // Record is the row record of the table.
```

##### (Record) GMap 

``` go
func (r Record) GMap() *gmap.StrAnyMap
```

GMap converts `r` to a gmap.

##### (Record) IsEmpty 

``` go
func (r Record) IsEmpty() bool
```

IsEmpty checks and returns whether `r` is empty.

##### (Record) Json 

``` go
func (r Record) Json() string
```

Json converts `r` to JSON format content.

##### (Record) Map 

``` go
func (r Record) Map() Map
```

Map converts `r` to map[string]interface{}.

##### (Record) Struct 

``` go
func (r Record) Struct(pointer interface{}) error
```

Struct converts `r` to a struct. Note that the parameter `pointer` should be type of *struct/**struct.

Note that it returns sql.ErrNoRows if `r` is empty.

##### (Record) Xml 

``` go
func (r Record) Xml(rootTag ...string) string
```

Xml converts `r` to XML format content.

#### type Result 

``` go
type Result []Record // Result is the row record array.
```

##### (Result) Array 

``` go
func (r Result) Array(field ...string) []Value
```

Array retrieves and returns specified column values as slice. The parameter `field` is optional is the column field is only one. The default `field` is the first field name of the first item in `Result` if parameter `field` is not given.

##### (Result) Chunk 

``` go
func (r Result) Chunk(size int) []Result
```

Chunk splits a Result into multiple Results, the size of each array is determined by `size`. The last chunk may contain less than size elements.

##### (Result) IsEmpty 

``` go
func (r Result) IsEmpty() bool
```

IsEmpty checks and returns whether `r` is empty.

##### (Result) Json 

``` go
func (r Result) Json() string
```

Json converts `r` to JSON format content.

##### (Result) Len 

``` go
func (r Result) Len() int
```

Len returns the length of result list.

##### (Result) List 

``` go
func (r Result) List() List
```

List converts `r` to a List.

##### (Result) MapKeyInt 

``` go
func (r Result) MapKeyInt(key string) map[int]Map
```

MapKeyInt converts `r` to a map[int]Map of which key is specified by `key`.

##### (Result) MapKeyStr 

``` go
func (r Result) MapKeyStr(key string) map[string]Map
```

MapKeyStr converts `r` to a map[string]Map of which key is specified by `key`.

##### (Result) MapKeyUint 

``` go
func (r Result) MapKeyUint(key string) map[uint]Map
```

MapKeyUint converts `r` to a map[uint]Map of which key is specified by `key`.

##### (Result) MapKeyValue 

``` go
func (r Result) MapKeyValue(key string) map[string]Value
```

MapKeyValue converts `r` to a map[string]Value of which key is specified by `key`. Note that the item value may be type of slice.

##### (Result) RecordKeyInt 

``` go
func (r Result) RecordKeyInt(key string) map[int]Record
```

RecordKeyInt converts `r` to a map[int]Record of which key is specified by `key`.

##### (Result) RecordKeyStr 

``` go
func (r Result) RecordKeyStr(key string) map[string]Record
```

RecordKeyStr converts `r` to a map[string]Record of which key is specified by `key`.

##### (Result) RecordKeyUint 

``` go
func (r Result) RecordKeyUint(key string) map[uint]Record
```

RecordKeyUint converts `r` to a map[uint]Record of which key is specified by `key`.

##### (Result) ScanList 

``` go
func (r Result) ScanList(structSlicePointer interface{}, bindToAttrName string, relationAttrNameAndFields ...string) (err error)
```

ScanList converts `r` to struct slice which contains other complex struct attributes. Note that the parameter `structSlicePointer` should be type of *[]struct/*[]*struct.

Usage example 1: Normal attribute struct relation:

``` go
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

var users []*Entity ScanList(&users, "User") ScanList(&users, "User", "uid") ScanList(&users, "UserDetail", "User", "uid:Uid") ScanList(&users, "UserScores", "User", "uid:Uid") ScanList(&users, "UserScores", "User", "uid")

Usage example 2: Embedded attribute struct relation:

``` go
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

var users []*Entity ScanList(&users) ScanList(&users, "UserDetail", "uid") ScanList(&users, "UserScores", "uid")

The parameters "User/UserDetail/UserScores" in the example codes specify the target attribute struct that current result will be bound to.

The "uid" in the example codes is the table field name of the result, and the "Uid" is the relational struct attribute name - not the attribute name of the bound to target. In the example codes, it's attribute name "Uid" of "User" of entity "Entity". It automatically calculates the HasOne/HasMany relationship with given `relation` parameter.

See the example or unit testing cases for clear understanding for this function.

##### (Result) Size 

``` go
func (r Result) Size() int
```

Size is alias of function Len.

##### (Result) Structs 

``` go
func (r Result) Structs(pointer interface{}) (err error)
```

Structs converts `r` to struct slice. Note that the parameter `pointer` should be type of *[]struct/*[]*struct.

##### (Result) Xml 

``` go
func (r Result) Xml(rootTag ...string) string
```

Xml converts `r` to XML format content.

#### type Schema 

``` go
type Schema struct {
	DB
}
```

Schema is a schema object from which it can then create a Model.

#### type SoftTimeOption <-2.6.3

``` go
type SoftTimeOption struct {
	SoftTimeType SoftTimeType // The value type for soft time field.
}
```

SoftTimeOption is the option to customize soft time feature for Model.

#### type SoftTimeType <-2.6.3

``` go
type SoftTimeType int
```

SoftTimeType custom defines the soft time field type.

``` go
const (
	SoftTimeTypeAuto           SoftTimeType = 0 // (Default)Auto detect the field type by table field type.
	SoftTimeTypeTime           SoftTimeType = 1 // Using datetime as the field value.
	SoftTimeTypeTimestamp      SoftTimeType = 2 // In unix seconds.
	SoftTimeTypeTimestampMilli SoftTimeType = 3 // In unix milliseconds.
	SoftTimeTypeTimestampMicro SoftTimeType = 4 // In unix microseconds.
	SoftTimeTypeTimestampNano  SoftTimeType = 5 // In unix nanoseconds.
)
```

#### type Sql 

``` go
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

#### type SqlResult 

``` go
type SqlResult struct {
	Result   sql.Result
	Affected int64
}
```

SqlResult is execution result for sql operations. It also supports batch operation result for rowsAffected.

##### (*SqlResult) LastInsertId 

``` go
func (r *SqlResult) LastInsertId() (int64, error)
```

LastInsertId returns the integer generated by the database in response to a command. Typically, this will be from an "auto increment" column when inserting a new row. Not all databases support this feature, and the syntax of such statements varies. Also, See sql.Result.

##### (*SqlResult) MustGetAffected 

``` go
func (r *SqlResult) MustGetAffected() int64
```

MustGetAffected returns the affected rows count, if any error occurs, it panics.

##### (*SqlResult) MustGetInsertId 

``` go
func (r *SqlResult) MustGetInsertId() int64
```

MustGetInsertId returns the last insert id, if any error occurs, it panics.

##### (*SqlResult) RowsAffected 

``` go
func (r *SqlResult) RowsAffected() (int64, error)
```

RowsAffected returns the number of rows affected by an update, insert, or delete. Not every database or database driver may support this. Also, See sql.Result.

#### type Stmt 

``` go
type Stmt struct {
	*sql.Stmt
	// contains filtered or unexported fields
}
```

Stmt is a prepared statement. A Stmt is safe for concurrent use by multiple goroutines.

If a Stmt is prepared on a Tx or Conn, it will be bound to a single underlying connection forever. If the Tx or Conn closes, the Stmt will become unusable and all operations will return an error. If a Stmt is prepared on a DB, it will remain usable for the lifetime of the DB. When the Stmt needs to execute on a new underlying connection, it will prepare itself on the new connection automatically.

##### (*Stmt) Close 

``` go
func (s *Stmt) Close() error
```

Close closes the statement.

##### (*Stmt) Exec 

``` go
func (s *Stmt) Exec(args ...interface{}) (sql.Result, error)
```

Exec executes a prepared statement with the given arguments and returns a Result summarizing the effect of the statement.

##### (*Stmt) ExecContext 

``` go
func (s *Stmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
```

ExecContext executes a prepared statement with the given arguments and returns a Result summarizing the effect of the statement.

##### (*Stmt) Query 

``` go
func (s *Stmt) Query(args ...interface{}) (*sql.Rows, error)
```

Query executes a prepared query statement with the given arguments and returns the query results as a *Rows.

##### (*Stmt) QueryContext 

``` go
func (s *Stmt) QueryContext(ctx context.Context, args ...interface{}) (*sql.Rows, error)
```

QueryContext executes a prepared query statement with the given arguments and returns the query results as a *Rows.

##### (*Stmt) QueryRow 

``` go
func (s *Stmt) QueryRow(args ...interface{}) *sql.Row
```

QueryRow executes a prepared query statement with the given arguments. If an error occurs during the execution of the statement, that error will be returned by a call to Scan on the returned *Row, which is always non-nil. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

Example usage:

``` go
var name string
err := nameByUseridStmt.QueryRow(id).Scan(&name)
```

##### (*Stmt) QueryRowContext 

``` go
func (s *Stmt) QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row
```

QueryRowContext executes a prepared query statement with the given arguments. If an error occurs during the execution of the statement, that error will be returned by a call to Scan on the returned *Row, which is always non-nil. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.

#### type TX 

``` go
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

##### func TXFromCtx 

``` go
func TXFromCtx(ctx context.Context, group string) TX
```

TXFromCtx retrieves and returns transaction object from context. It is usually used in nested transaction feature, and it returns nil if it is not set previously.

#### type TXCore <-2.3.0

``` go
type TXCore struct {
	// contains filtered or unexported fields
}
```

TXCore is the struct for transaction management.

##### (*TXCore) Begin <-2.3.0

``` go
func (tx *TXCore) Begin() error
```

Begin starts a nested transaction procedure.

##### (*TXCore) Commit <-2.3.0

``` go
func (tx *TXCore) Commit() error
```

Commit commits current transaction. Note that it releases previous saved transaction point if it's in a nested transaction procedure, or else it commits the hole transaction.

##### (*TXCore) Ctx <-2.3.0

``` go
func (tx *TXCore) Ctx(ctx context.Context) TX
```

Ctx sets the context for current transaction.

##### (*TXCore) Delete <-2.3.0

``` go
func (tx *TXCore) Delete(table string, condition interface{}, args ...interface{}) (sql.Result, error)
```

Delete does "DELETE FROM ... " statement for the table.

The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc. It is commonly used with parameter `args`. Eg: "uid=10000", "uid", 10000 "money>? AND name like ?", 99999, "vip_%" "status IN (?)", g.Slice{1,2,3} "age IN(?,?)", 18, 50 User{ Id : 1, UserName : "john"}.

##### (*TXCore) Exec <-2.3.0

``` go
func (tx *TXCore) Exec(sql string, args ...interface{}) (sql.Result, error)
```

Exec does none query operation on transaction. See Core.Exec.

##### (*TXCore) ExecContext <-2.5.2

``` go
func (tx *TXCore) ExecContext(ctx context.Context, sql string, args ...interface{}) (sql.Result, error)
```

ExecContext implements interface function Link.ExecContext.

##### (*TXCore) GetAll <-2.3.0

``` go
func (tx *TXCore) GetAll(sql string, args ...interface{}) (Result, error)
```

GetAll queries and returns data records from database.

##### (*TXCore) GetCount <-2.3.0

``` go
func (tx *TXCore) GetCount(sql string, args ...interface{}) (int64, error)
```

GetCount queries and returns the count from database.

##### (*TXCore) GetCtx <-2.3.0

``` go
func (tx *TXCore) GetCtx() context.Context
```

GetCtx returns the context for current transaction.

##### (*TXCore) GetDB <-2.3.0

``` go
func (tx *TXCore) GetDB() DB
```

GetDB returns the DB for current transaction.

##### (*TXCore) GetOne <-2.3.0

``` go
func (tx *TXCore) GetOne(sql string, args ...interface{}) (Record, error)
```

GetOne queries and returns one record from database.

##### (*TXCore) GetScan <-2.3.0

``` go
func (tx *TXCore) GetScan(pointer interface{}, sql string, args ...interface{}) error
```

GetScan queries one or more records from database and converts them to given struct or struct array.

If parameter `pointer` is type of struct pointer, it calls GetStruct internally for the conversion. If parameter `pointer` is type of slice, it calls GetStructs internally for conversion.

##### (*TXCore) GetSqlTX <-2.3.0

``` go
func (tx *TXCore) GetSqlTX() *sql.Tx
```

GetSqlTX returns the underlying transaction object for current transaction.

##### (*TXCore) GetStruct <-2.3.0

``` go
func (tx *TXCore) GetStruct(obj interface{}, sql string, args ...interface{}) error
```

GetStruct queries one record from database and converts it to given struct. The parameter `pointer` should be a pointer to struct.

##### (*TXCore) GetStructs <-2.3.0

``` go
func (tx *TXCore) GetStructs(objPointerSlice interface{}, sql string, args ...interface{}) error
```

GetStructs queries records from database and converts them to given struct. The parameter `pointer` should be type of struct slice: []struct/[]*struct.

##### (*TXCore) GetValue <-2.3.0

``` go
func (tx *TXCore) GetValue(sql string, args ...interface{}) (Value, error)
```

GetValue queries and returns the field value from database. The sql should query only one field from database, or else it returns only one field of the result.

##### (*TXCore) Insert <-2.3.0

``` go
func (tx *TXCore) Insert(table string, data interface{}, batch ...int) (sql.Result, error)
```

Insert does "INSERT INTO ..." statement for the table. If there's already one unique record of the data in the table, it returns error.

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})

The parameter `batch` specifies the batch operation count when given data is slice.

##### (*TXCore) InsertAndGetId <-2.3.0

``` go
func (tx *TXCore) InsertAndGetId(table string, data interface{}, batch ...int) (int64, error)
```

InsertAndGetId performs action Insert and returns the last insert id that automatically generated.

##### (*TXCore) InsertIgnore <-2.3.0

``` go
func (tx *TXCore) InsertIgnore(table string, data interface{}, batch ...int) (sql.Result, error)
```

InsertIgnore does "INSERT IGNORE INTO ..." statement for the table. If there's already one unique record of the data in the table, it ignores the inserting.

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})

The parameter `batch` specifies the batch operation count when given data is slice.

##### (*TXCore) IsClosed <-2.3.0

``` go
func (tx *TXCore) IsClosed() bool
```

IsClosed checks and returns this transaction has already been committed or rolled back.

##### (*TXCore) IsOnMaster <-2.5.2

``` go
func (tx *TXCore) IsOnMaster() bool
```

IsOnMaster implements interface function Link.IsOnMaster.

##### (*TXCore) IsTransaction <-2.5.2

``` go
func (tx *TXCore) IsTransaction() bool
```

IsTransaction implements interface function Link.IsTransaction.

##### (*TXCore) Model <-2.3.0

``` go
func (tx *TXCore) Model(tableNameQueryOrStruct ...interface{}) *Model
```

Model acts like Core.Model except it operates on transaction. See Core.Model.

##### (*TXCore) Prepare <-2.3.0

``` go
func (tx *TXCore) Prepare(sql string) (*Stmt, error)
```

Prepare creates a prepared statement for later queries or executions. Multiple queries or executions may be run concurrently from the returned statement. The caller must call the statement's Close method when the statement is no longer needed.

##### (*TXCore) PrepareContext <-2.5.2

``` go
func (tx *TXCore) PrepareContext(ctx context.Context, sql string) (*sql.Stmt, error)
```

PrepareContext implements interface function Link.PrepareContext.

##### (*TXCore) Query <-2.3.0

``` go
func (tx *TXCore) Query(sql string, args ...interface{}) (result Result, err error)
```

Query does query operation on transaction. See Core.Query.

##### (*TXCore) QueryContext <-2.5.2

``` go
func (tx *TXCore) QueryContext(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
```

QueryContext implements interface function Link.QueryContext.

##### (*TXCore) Raw <-2.3.0

``` go
func (tx *TXCore) Raw(rawSql string, args ...interface{}) *Model
```

##### (*TXCore) Replace <-2.3.0

``` go
func (tx *TXCore) Replace(table string, data interface{}, batch ...int) (sql.Result, error)
```

Replace does "REPLACE INTO ..." statement for the table. If there's already one unique record of the data in the table, it deletes the record and inserts a new one.

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. If given data is type of slice, it then does batch replacing, and the optional parameter `batch` specifies the batch operation count.

##### (*TXCore) Rollback <-2.3.0

``` go
func (tx *TXCore) Rollback() error
```

Rollback aborts current transaction. Note that it aborts current transaction if it's in a nested transaction procedure, or else it aborts the hole transaction.

##### (*TXCore) RollbackTo <-2.3.0

``` go
func (tx *TXCore) RollbackTo(point string) error
```

RollbackTo performs `ROLLBACK TO SAVEPOINT xxx` SQL statement that rollbacks to specified saved transaction. The parameter `point` specifies the point name that was saved previously.

##### (*TXCore) Save <-2.3.0

``` go
func (tx *TXCore) Save(table string, data interface{}, batch ...int) (sql.Result, error)
```

Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the table. It updates the record if there's primary or unique index in the saving data, or else it inserts a new record into the table.

The parameter `data` can be type of map/gmap/struct/*struct/[]map/[]struct, etc. Eg: Data(g.Map{"uid": 10000, "name":"john"}) Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})

If given data is type of slice, it then does batch saving, and the optional parameter `batch` specifies the batch operation count.

##### (*TXCore) SavePoint <-2.3.0

``` go
func (tx *TXCore) SavePoint(point string) error
```

SavePoint performs `SAVEPOINT xxx` SQL statement that saves transaction at current point. The parameter `point` specifies the point name that will be saved to server.

##### (*TXCore) Transaction <-2.3.0

``` go
func (tx *TXCore) Transaction(ctx context.Context, f func(ctx context.Context, tx TX) error) (err error)
```

Transaction wraps the transaction logic using function `f`. It rollbacks the transaction and returns the error from function `f` if it returns non-nil error. It commits the transaction and returns nil if function `f` returns nil.

Note that, you should not Commit or Rollback the transaction in function `f` as it is automatically handled by this function.

##### (*TXCore) Update <-2.3.0

``` go
func (tx *TXCore) Update(table string, data interface{}, condition interface{}, args ...interface{}) (sql.Result, error)
```

Update does "UPDATE ... " statement for the table.

The parameter `data` can be type of string/map/gmap/struct/*struct, etc. Eg: "uid=10000", "uid", 10000, g.Map{"uid": 10000, "name":"john"}

The parameter `condition` can be type of string/map/gmap/slice/struct/*struct, etc. It is commonly used with parameter `args`. Eg: "uid=10000", "uid", 10000 "money>? AND name like ?", 99999, "vip_%" "status IN (?)", g.Slice{1,2,3} "age IN(?,?)", 18, 50 User{ Id : 1, UserName : "john"}.

##### (*TXCore) With <-2.3.0

``` go
func (tx *TXCore) With(object interface{}) *Model
```

With acts like Core.With except it operates on transaction. See Core.With.

#### type TableField 

``` go
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

#### type Value 

``` go
type Value = *gvar.Var // Value is the field value type.
```

#### type WhereBuilder <-2.1.0

``` go
type WhereBuilder struct {
	// contains filtered or unexported fields
}
```

WhereBuilder holds multiple where conditions in a group.

##### (*WhereBuilder) Build <-2.1.0

``` go
func (b *WhereBuilder) Build() (conditionWhere string, conditionArgs []interface{})
```

Build builds current WhereBuilder and returns the condition string and parameters.

##### (*WhereBuilder) Clone <-2.1.0

``` go
func (b *WhereBuilder) Clone() *WhereBuilder
```

Clone clones and returns a WhereBuilder that is a copy of current one.

##### (*WhereBuilder) Where <-2.1.0

``` go
func (b *WhereBuilder) Where(where interface{}, args ...interface{}) *WhereBuilder
```

Where sets the condition statement for the builder. The parameter `where` can be type of string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times, multiple conditions will be joined into where statement using "AND". Eg: Where("uid=10000") Where("uid", 10000) Where("money>? AND name like ?", 99999, "vip_%") Where("uid", 1).Where("name", "john") Where("status IN (?)", g.Slice{1,2,3}) Where("age IN(?,?)", 18, 50) Where(User{ Id : 1, UserName : "john"}).

##### (*WhereBuilder) WhereBetween <-2.1.0

``` go
func (b *WhereBuilder) WhereBetween(column string, min, max interface{}) *WhereBuilder
```

WhereBetween builds `column BETWEEN min AND max` statement.

##### (*WhereBuilder) WhereGT <-2.1.0

``` go
func (b *WhereBuilder) WhereGT(column string, value interface{}) *WhereBuilder
```

WhereGT builds `column > value` statement.

##### (*WhereBuilder) WhereGTE <-2.1.0

``` go
func (b *WhereBuilder) WhereGTE(column string, value interface{}) *WhereBuilder
```

WhereGTE builds `column >= value` statement.

##### (*WhereBuilder) WhereIn <-2.1.0

``` go
func (b *WhereBuilder) WhereIn(column string, in interface{}) *WhereBuilder
```

WhereIn builds `column IN (in)` statement.

##### (*WhereBuilder) WhereLT <-2.1.0

``` go
func (b *WhereBuilder) WhereLT(column string, value interface{}) *WhereBuilder
```

WhereLT builds `column < value` statement.

##### (*WhereBuilder) WhereLTE <-2.1.0

``` go
func (b *WhereBuilder) WhereLTE(column string, value interface{}) *WhereBuilder
```

WhereLTE builds `column <= value` statement.

##### (*WhereBuilder) WhereLike <-2.1.0

``` go
func (b *WhereBuilder) WhereLike(column string, like string) *WhereBuilder
```

WhereLike builds `column LIKE like` statement.

##### (*WhereBuilder) WhereNot <-2.1.0

``` go
func (b *WhereBuilder) WhereNot(column string, value interface{}) *WhereBuilder
```

WhereNot builds `column != value` statement.

##### (*WhereBuilder) WhereNotBetween <-2.1.0

``` go
func (b *WhereBuilder) WhereNotBetween(column string, min, max interface{}) *WhereBuilder
```

WhereNotBetween builds `column NOT BETWEEN min AND max` statement.

##### (*WhereBuilder) WhereNotIn <-2.1.0

``` go
func (b *WhereBuilder) WhereNotIn(column string, in interface{}) *WhereBuilder
```

WhereNotIn builds `column NOT IN (in)` statement.

##### (*WhereBuilder) WhereNotLike <-2.1.0

``` go
func (b *WhereBuilder) WhereNotLike(column string, like interface{}) *WhereBuilder
```

WhereNotLike builds `column NOT LIKE like` statement.

##### (*WhereBuilder) WhereNotNull <-2.1.0

``` go
func (b *WhereBuilder) WhereNotNull(columns ...string) *WhereBuilder
```

WhereNotNull builds `columns[0] IS NOT NULL AND columns[1] IS NOT NULL ...` statement.

##### (*WhereBuilder) WhereNull <-2.1.0

``` go
func (b *WhereBuilder) WhereNull(columns ...string) *WhereBuilder
```

WhereNull builds `columns[0] IS NULL AND columns[1] IS NULL ...` statement.

##### (*WhereBuilder) WhereOr <-2.1.0

``` go
func (b *WhereBuilder) WhereOr(where interface{}, args ...interface{}) *WhereBuilder
```

WhereOr adds "OR" condition to the where statement.

##### (*WhereBuilder) WhereOrBetween <-2.1.0

``` go
func (b *WhereBuilder) WhereOrBetween(column string, min, max interface{}) *WhereBuilder
```

WhereOrBetween builds `column BETWEEN min AND max` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrGT <-2.1.0

``` go
func (b *WhereBuilder) WhereOrGT(column string, value interface{}) *WhereBuilder
```

WhereOrGT builds `column > value` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrGTE <-2.1.0

``` go
func (b *WhereBuilder) WhereOrGTE(column string, value interface{}) *WhereBuilder
```

WhereOrGTE builds `column >= value` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrIn <-2.1.0

``` go
func (b *WhereBuilder) WhereOrIn(column string, in interface{}) *WhereBuilder
```

WhereOrIn builds `column IN (in)` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrLT <-2.1.0

``` go
func (b *WhereBuilder) WhereOrLT(column string, value interface{}) *WhereBuilder
```

WhereOrLT builds `column < value` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrLTE <-2.1.0

``` go
func (b *WhereBuilder) WhereOrLTE(column string, value interface{}) *WhereBuilder
```

WhereOrLTE builds `column <= value` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrLike <-2.1.0

``` go
func (b *WhereBuilder) WhereOrLike(column string, like interface{}) *WhereBuilder
```

WhereOrLike builds `column LIKE 'like'` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrNot <-2.4.2

``` go
func (b *WhereBuilder) WhereOrNot(column string, value interface{}) *WhereBuilder
```

WhereOrNot builds `column != value` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrNotBetween <-2.1.0

``` go
func (b *WhereBuilder) WhereOrNotBetween(column string, min, max interface{}) *WhereBuilder
```

WhereOrNotBetween builds `column NOT BETWEEN min AND max` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrNotIn <-2.1.0

``` go
func (b *WhereBuilder) WhereOrNotIn(column string, in interface{}) *WhereBuilder
```

WhereOrNotIn builds `column NOT IN (in)` statement.

##### (*WhereBuilder) WhereOrNotLike <-2.1.0

``` go
func (b *WhereBuilder) WhereOrNotLike(column string, like interface{}) *WhereBuilder
```

WhereOrNotLike builds `column NOT LIKE like` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrNotNull <-2.1.0

``` go
func (b *WhereBuilder) WhereOrNotNull(columns ...string) *WhereBuilder
```

WhereOrNotNull builds `columns[0] IS NOT NULL OR columns[1] IS NOT NULL ...` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrNull <-2.1.0

``` go
func (b *WhereBuilder) WhereOrNull(columns ...string) *WhereBuilder
```

WhereOrNull builds `columns[0] IS NULL OR columns[1] IS NULL ...` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefix <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefix(prefix string, where interface{}, args ...interface{}) *WhereBuilder
```

WhereOrPrefix performs as WhereOr, but it adds prefix to each field in where statement. Eg: WhereOrPrefix("order", "status", "paid") => WHERE xxx OR (`order`.`status`='paid') WhereOrPrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE xxx OR (`order`.`status`='paid' AND `order`.`channel`='bank')

##### (*WhereBuilder) WhereOrPrefixBetween <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder
```

WhereOrPrefixBetween builds `prefix.column BETWEEN min AND max` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixGT <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixGT(prefix string, column string, value interface{}) *WhereBuilder
```

WhereOrPrefixGT builds `prefix.column > value` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixGTE <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixGTE(prefix string, column string, value interface{}) *WhereBuilder
```

WhereOrPrefixGTE builds `prefix.column >= value` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixIn <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixIn(prefix string, column string, in interface{}) *WhereBuilder
```

WhereOrPrefixIn builds `prefix.column IN (in)` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixLT <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixLT(prefix string, column string, value interface{}) *WhereBuilder
```

WhereOrPrefixLT builds `prefix.column < value` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixLTE <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixLTE(prefix string, column string, value interface{}) *WhereBuilder
```

WhereOrPrefixLTE builds `prefix.column <= value` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixLike <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixLike(prefix string, column string, like interface{}) *WhereBuilder
```

WhereOrPrefixLike builds `prefix.column LIKE 'like'` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixNot <-2.4.2

``` go
func (b *WhereBuilder) WhereOrPrefixNot(prefix string, column string, value interface{}) *WhereBuilder
```

WhereOrPrefixNot builds `prefix.column != value` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixNotBetween <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder
```

WhereOrPrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixNotIn <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder
```

WhereOrPrefixNotIn builds `prefix.column NOT IN (in)` statement.

##### (*WhereBuilder) WhereOrPrefixNotLike <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder
```

WhereOrPrefixNotLike builds `prefix.column NOT LIKE 'like'` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixNotNull <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixNotNull(prefix string, columns ...string) *WhereBuilder
```

WhereOrPrefixNotNull builds `prefix.columns[0] IS NOT NULL OR prefix.columns[1] IS NOT NULL ...` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrPrefixNull <-2.1.0

``` go
func (b *WhereBuilder) WhereOrPrefixNull(prefix string, columns ...string) *WhereBuilder
```

WhereOrPrefixNull builds `prefix.columns[0] IS NULL OR prefix.columns[1] IS NULL ...` statement in `OR` conditions.

##### (*WhereBuilder) WhereOrf <-2.1.0

``` go
func (b *WhereBuilder) WhereOrf(format string, args ...interface{}) *WhereBuilder
```

WhereOrf builds `OR` condition string using fmt.Sprintf and arguments. Eg: WhereOrf(`amount<? and status=%s`, "paid", 100) => WHERE xxx OR `amount`<100 and status='paid' WhereOrf(`amount<%d and status=%s`, 100, "paid") => WHERE xxx OR `amount`<100 and status='paid'

##### (*WhereBuilder) WherePrefix <-2.1.0

``` go
func (b *WhereBuilder) WherePrefix(prefix string, where interface{}, args ...interface{}) *WhereBuilder
```

WherePrefix performs as Where, but it adds prefix to each field in where statement. Eg: WherePrefix("order", "status", "paid") => WHERE `order`.`status`='paid' WherePrefix("order", struct{Status:"paid", "channel":"bank"}) => WHERE `order`.`status`='paid' AND `order`.`channel`='bank'

##### (*WhereBuilder) WherePrefixBetween <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixBetween(prefix string, column string, min, max interface{}) *WhereBuilder
```

WherePrefixBetween builds `prefix.column BETWEEN min AND max` statement.

##### (*WhereBuilder) WherePrefixGT <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixGT(prefix string, column string, value interface{}) *WhereBuilder
```

WherePrefixGT builds `prefix.column > value` statement.

##### (*WhereBuilder) WherePrefixGTE <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixGTE(prefix string, column string, value interface{}) *WhereBuilder
```

WherePrefixGTE builds `prefix.column >= value` statement.

##### (*WhereBuilder) WherePrefixIn <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixIn(prefix string, column string, in interface{}) *WhereBuilder
```

WherePrefixIn builds `prefix.column IN (in)` statement.

##### (*WhereBuilder) WherePrefixLT <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixLT(prefix string, column string, value interface{}) *WhereBuilder
```

WherePrefixLT builds `prefix.column < value` statement.

##### (*WhereBuilder) WherePrefixLTE <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixLTE(prefix string, column string, value interface{}) *WhereBuilder
```

WherePrefixLTE builds `prefix.column <= value` statement.

##### (*WhereBuilder) WherePrefixLike <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixLike(prefix string, column string, like interface{}) *WhereBuilder
```

WherePrefixLike builds `prefix.column LIKE like` statement.

##### (*WhereBuilder) WherePrefixNot <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixNot(prefix string, column string, value interface{}) *WhereBuilder
```

WherePrefixNot builds `prefix.column != value` statement.

##### (*WhereBuilder) WherePrefixNotBetween <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixNotBetween(prefix string, column string, min, max interface{}) *WhereBuilder
```

WherePrefixNotBetween builds `prefix.column NOT BETWEEN min AND max` statement.

##### (*WhereBuilder) WherePrefixNotIn <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixNotIn(prefix string, column string, in interface{}) *WhereBuilder
```

WherePrefixNotIn builds `prefix.column NOT IN (in)` statement.

##### (*WhereBuilder) WherePrefixNotLike <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixNotLike(prefix string, column string, like interface{}) *WhereBuilder
```

WherePrefixNotLike builds `prefix.column NOT LIKE like` statement.

##### (*WhereBuilder) WherePrefixNotNull <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixNotNull(prefix string, columns ...string) *WhereBuilder
```

WherePrefixNotNull builds `prefix.columns[0] IS NOT NULL AND prefix.columns[1] IS NOT NULL ...` statement.

##### (*WhereBuilder) WherePrefixNull <-2.1.0

``` go
func (b *WhereBuilder) WherePrefixNull(prefix string, columns ...string) *WhereBuilder
```

WherePrefixNull builds `prefix.columns[0] IS NULL AND prefix.columns[1] IS NULL ...` statement.

##### (*WhereBuilder) WherePri <-2.1.0

``` go
func (b *WhereBuilder) WherePri(where interface{}, args ...interface{}) *WhereBuilder
```

WherePri does the same logic as Model.Where except that if the parameter `where` is a single condition like int/string/float/slice, it treats the condition as the primary key value. That is, if primary key is "id" and given `where` parameter as "123", the WherePri function treats the condition as "id=123", but Model.Where treats the condition as string "123".

##### (*WhereBuilder) Wheref <-2.1.0

``` go
func (b *WhereBuilder) Wheref(format string, args ...interface{}) *WhereBuilder
```

Wheref builds condition string using fmt.Sprintf and arguments. Note that if the number of `args` is more than the placeholder in `format`, the extra `args` will be used as the where condition arguments of the Model. Eg: Wheref(`amount<? and status=%s`, "paid", 100) => WHERE `amount`<100 and status='paid' Wheref(`amount<%d and status=%s`, 100, "paid") => WHERE `amount`<100 and status='paid'

#### type WhereHolder <-2.1.0

``` go
type WhereHolder struct {
	Type     string        // Type of this holder.
	Operator int           // Operator for this holder.
	Where    interface{}   // Where parameter, which can commonly be type of string/map/struct.
	Args     []interface{} // Arguments for where parameter.
	Prefix   string        // Field prefix, eg: "user.", "order.".
}
```

WhereHolder is the holder for where condition preparing.