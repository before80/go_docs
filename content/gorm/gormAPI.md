+++
title = "gorm API"
date = 2023-10-30T18:37:28+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++



The fantastic ORM library for Golang, aims to be developer friendly.

## Overview

- Full-Featured ORM
- Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism, Single-table inheritance)
- Hooks (Before/After Create/Save/Update/Delete/Find)
- Eager loading with `Preload`, `Joins`
- Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
- Context, Prepared Statement Mode, DryRun Mode
- Batch Insert, FindInBatches, Find To Map
- SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, NamedArg, Search/Update/Create with SQL Expr
- Composite Primary Key
- Auto Migrations
- Logger
- Extendable, flexible plugin API: Database Resolver (Multiple Databases, Read/Write Splitting) / Prometheus…
- Every feature comes with tests
- Developer Friendly

## Getting Started

- GORM Guides [https://gorm.io](https://gorm.io/)
- Gen Guides https://gorm.io/gen/index.html

## Contributing

[You can help to deliver a better GORM, check out things you can do](https://gorm.io/contribute.html)

## Contributors

[Thank you](https://github.com/go-gorm/gorm/graphs/contributors) for contributing to the GORM framework!

## License

© Jinzhu, 2013~time.Now

Released under the [MIT License](https://github.com/go-gorm/gorm/raw/master/License)



## Constants 

This section is empty.

## Variables 

[View Source](https://github.com/go-gorm/gorm/blob/v1.25.5/errors.go#L9)

``` go
var (
	// ErrRecordNotFound record not found error
	ErrRecordNotFound = logger.ErrRecordNotFound
	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction = errors.New("invalid transaction")
	// ErrNotImplemented not implemented
	ErrNotImplemented = errors.New("not implemented")
	// ErrMissingWhereClause missing where clause
	ErrMissingWhereClause = errors.New("WHERE conditions required")
	// ErrUnsupportedRelation unsupported relations
	ErrUnsupportedRelation = errors.New("unsupported relations")
	// ErrPrimaryKeyRequired primary keys required
	ErrPrimaryKeyRequired = errors.New("primary key required")
	// ErrModelValueRequired model value required
	ErrModelValueRequired = errors.New("model value required")
	// ErrModelAccessibleFieldsRequired model accessible fields required
	ErrModelAccessibleFieldsRequired = errors.New("model accessible fields required")
	// ErrSubQueryRequired sub query required
	ErrSubQueryRequired = errors.New("sub query required")
	// ErrInvalidData unsupported data
	ErrInvalidData = errors.New("unsupported data")
	// ErrUnsupportedDriver unsupported driver
	ErrUnsupportedDriver = errors.New("unsupported driver")
	// ErrRegistered registered
	ErrRegistered = errors.New("registered")
	// ErrInvalidField invalid field
	ErrInvalidField = errors.New("invalid field")
	// ErrEmptySlice empty slice found
	ErrEmptySlice = errors.New("empty slice found")
	// ErrDryRunModeUnsupported dry run mode unsupported
	ErrDryRunModeUnsupported = errors.New("dry run mode unsupported")
	// ErrInvalidDB invalid db
	ErrInvalidDB = errors.New("invalid db")
	// ErrInvalidValue invalid value
	ErrInvalidValue = errors.New("invalid value, should be pointer to struct or slice")
	// ErrInvalidValueOfLength invalid values do not match length
	ErrInvalidValueOfLength = errors.New("invalid association values, length doesn't match")
	// ErrPreloadNotAllowed preload is not allowed when count is used
	ErrPreloadNotAllowed = errors.New("preload is not allowed when count is used")
	// ErrDuplicatedKey occurs when there is a unique key constraint violation
	ErrDuplicatedKey = errors.New("duplicated key not allowed")
	// ErrForeignKeyViolated occurs when there is a foreign key constraint violation
	ErrForeignKeyViolated = errors.New("violates foreign key constraint")
)
```

## Functions 

### func Expr 

``` go
func Expr(expr string, args ...interface{}) clause.Expr
```

Expr returns clause.Expr, which can be used to pass SQL expression as params

### func Scan 

``` go
func Scan(rows Rows, db *DB, mode ScanMode)
```

Scan scan rows into db statement

## Types 

### type Association 

``` go
type Association struct {
	DB           *DB
	Relationship *schema.Relationship
	Unscope      bool
	Error        error
}
```

Association Mode contains some helper methods to handle relationship things easily.

(*Association) Append 

``` go
func (association *Association) Append(values ...interface{}) error
```

#### (*Association) Clear 

``` go
func (association *Association) Clear() error
```

#### (*Association) Count 

``` go
func (association *Association) Count() (count int64)
```

#### (*Association) Delete 

``` go
func (association *Association) Delete(values ...interface{}) error
```

#### (*Association) Find 

``` go
func (association *Association) Find(out interface{}, conds ...interface{}) error
```

#### (*Association) Replace 

``` go
func (association *Association) Replace(values ...interface{}) error
```

#### (*Association) Unscoped <- 1.25.1

``` go
func (association *Association) Unscoped() *Association
```

### type ColumnType <- 1.20.5

``` go
type ColumnType interface {
	Name() string
	DatabaseTypeName() string                 // varchar
	ColumnType() (columnType string, ok bool) // varchar(64)
	PrimaryKey() (isPrimaryKey bool, ok bool)
	AutoIncrement() (isAutoIncrement bool, ok bool)
	Length() (length int64, ok bool)
	DecimalSize() (precision int64, scale int64, ok bool)
	Nullable() (nullable bool, ok bool)
	Unique() (unique bool, ok bool)
	ScanType() reflect.Type
	Comment() (value string, ok bool)
	DefaultValue() (value string, ok bool)
}
```

ColumnType column type interface

### type Config 

``` go
type Config struct {
	// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
	// You can disable it by setting `SkipDefaultTransaction` to true
	SkipDefaultTransaction bool
	// NamingStrategy tables, columns naming strategy
	NamingStrategy schema.Namer
	// FullSaveAssociations full save associations
	FullSaveAssociations bool
	// Logger
	Logger logger.Interface
	// NowFunc the function to be used when creating a new timestamp
	NowFunc func() time.Time
	// DryRun generate sql without execute
	DryRun bool
	// PrepareStmt executes the given query in cached statement
	PrepareStmt bool
	// DisableAutomaticPing
	DisableAutomaticPing bool
	// DisableForeignKeyConstraintWhenMigrating
	DisableForeignKeyConstraintWhenMigrating bool
	// IgnoreRelationshipsWhenMigrating
	IgnoreRelationshipsWhenMigrating bool
	// DisableNestedTransaction disable nested transaction
	DisableNestedTransaction bool
	// AllowGlobalUpdate allow global update
	AllowGlobalUpdate bool
	// QueryFields executes the SQL query with all fields of the table
	QueryFields bool
	// CreateBatchSize default create batch size
	CreateBatchSize int
	// TranslateError enabling error translation
	TranslateError bool

	// ClauseBuilders clause builder
	ClauseBuilders map[string]clause.ClauseBuilder
	// ConnPool db conn pool
	ConnPool ConnPool
	// Dialector database dialector
	Dialector
	// Plugins registered plugins
	Plugins map[string]Plugin
	// contains filtered or unexported fields
}
```

Config GORM config

#### (*Config) AfterInitialize <- 1.21.0

``` go
func (c *Config) AfterInitialize(db *DB) error
```

AfterInitialize initialize plugins after db connected

#### (*Config) Apply <- 1.21.0

``` go
func (c *Config) Apply(config *Config) error
```

Apply update config to new config

### type ConnPool 

``` go
type ConnPool interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}
```

ConnPool db conns pool interface

### type ConnPoolBeginner <- 0.2.3

``` go
type ConnPoolBeginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (ConnPool, error)
}
```

ConnPoolBeginner conn pool beginner

### type DB 

``` go
type DB struct {
	*Config
	Error        error
	RowsAffected int64
	Statement    *Statement
	// contains filtered or unexported fields
}
```

DB GORM DB definition

#### func Open 

``` go
func Open(dialector Dialector, opts ...Option) (db *DB, err error)
```

Open initialize db session based on dialector

#### (*DB) AddError 

``` go
func (db *DB) AddError(err error) error
```

AddError add error to db

#### (*DB) Assign 

``` go
func (db *DB) Assign(attrs ...interface{}) (tx *DB)
```

Assign provide attributes used in [FirstOrCreate](https://gorm.io/docs/advanced_query.html#FirstOrCreate) or [FirstOrInit](https://gorm.io/docs/advanced_query.html#FirstOrInit)

Assign adds attributes even if the record is found. If using FirstOrCreate, this means that records will be updated even if they are found.

```go
// assign an email regardless of if the record is not found
db.Where(User{Name: "non_existing"}).Assign(User{Email: "fake@fake.org"}).FirstOrInit(&user)
// user -> User{Name: "non_existing", Email: "fake@fake.org"}

// assign email regardless of if record is found
db.Where(User{Name: "jinzhu"}).Assign(User{Email: "fake@fake.org"}).FirstOrInit(&user)
// user -> User{Name: "jinzhu", Age: 20, Email: "fake@fake.org"}
```

#### (*DB) Association 

``` go
func (db *DB) Association(column string) *Association
```

#### (*DB) Attrs 

``` go
func (db *DB) Attrs(attrs ...interface{}) (tx *DB)
```

Attrs provide attributes used in [FirstOrCreate](https://gorm.io/docs/advanced_query.html#FirstOrCreate) or [FirstOrInit](https://gorm.io/docs/advanced_query.html#FirstOrInit)

Attrs only adds attributes if the record is not found.

```go
// assign an email if the record is not found
db.Where(User{Name: "non_existing"}).Attrs(User{Email: "fake@fake.org"}).FirstOrInit(&user)
// user -> User{Name: "non_existing", Email: "fake@fake.org"}

// assign an email if the record is not found, otherwise ignore provided email
db.Where(User{Name: "jinzhu"}).Attrs(User{Email: "fake@fake.org"}).FirstOrInit(&user)
// user -> User{Name: "jinzhu", Age: 20}
```

#### (*DB) AutoMigrate 

``` go
func (db *DB) AutoMigrate(dst ...interface{}) error
```

AutoMigrate run auto migration for given models

#### (*DB) Begin 

``` go
func (db *DB) Begin(opts ...*sql.TxOptions) *DB
```

Begin begins a transaction with any transaction options opts

#### (*DB) Callback 

``` go
func (db *DB) Callback() *callbacks
```

Callback returns callback manager

#### (*DB) Clauses 

``` go
func (db *DB) Clauses(conds ...clause.Expression) (tx *DB)
```

Clauses Add clauses

This supports both standard clauses (clause.OrderBy, clause.Limit, clause.Where) and more advanced techniques like specifying lock strength and optimizer hints. See the [docs](https://gorm.io/docs/sql_builder.html#Clauses) for more depth.

```go
// add a simple limit clause
db.Clauses(clause.Limit{Limit: 1}).Find(&User{})
// tell the optimizer to use the `idx_user_name` index
db.Clauses(hints.UseIndex("idx_user_name")).Find(&User{})
// specify the lock strength to UPDATE
db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)
```

#### (*DB) Commit 

``` go
func (db *DB) Commit() *DB
```

Commit commits the changes in a transaction

#### (*DB) Connection <- 1.22.5

``` go
func (db *DB) Connection(fc func(tx *DB) error) (err error)
```

Connection uses a db connection to execute an arbitrary number of commands in fc. When finished, the connection is returned to the connection pool.

#### (*DB) Count 

``` go
func (db *DB) Count(count *int64) (tx *DB)
```

#### (*DB) Create 

``` go
func (db *DB) Create(value interface{}) (tx *DB)
```

Create inserts value, returning the inserted data's primary key in value's id

#### (*DB) CreateInBatches <- 1.20.7

``` go
func (db *DB) CreateInBatches(value interface{}, batchSize int) (tx *DB)
```

CreateInBatches inserts value in batches of batchSize

#### (*DB) DB <- 0.2.7

``` go
func (db *DB) DB() (*sql.DB, error)
```

DB returns `*sql.DB`

#### (*DB) Debug 

``` go
func (db *DB) Debug() (tx *DB)
```

Debug start debug mode

#### (*DB) Delete 

``` go
func (db *DB) Delete(value interface{}, conds ...interface{}) (tx *DB)
```

Delete deletes value matching given conditions. If value contains primary key it is included in the conditions. If value includes a deleted_at field, then Delete performs a soft delete instead by setting deleted_at with the current time if null.

#### (*DB) Distinct <- 0.2.3

``` go
func (db *DB) Distinct(args ...interface{}) (tx *DB)
```

Distinct specify distinct fields that you want querying

```go
// Select distinct names of users
db.Distinct("name").Find(&results)
// Select distinct name/age pairs from users
db.Distinct("name", "age").Find(&results)
```

#### (*DB) Exec 

``` go
func (db *DB) Exec(sql string, values ...interface{}) (tx *DB)
```

Exec executes raw sql

#### (*DB) Find 

``` go
func (db *DB) Find(dest interface{}, conds ...interface{}) (tx *DB)
```

Find finds all records matching given conditions conds

#### (*DB) FindInBatches <- 0.2.6

``` go
func (db *DB) FindInBatches(dest interface{}, batchSize int, fc func(tx *DB, batch int) error) *DB
```

FindInBatches finds all records in batches of batchSize

#### (*DB) First 

``` go
func (db *DB) First(dest interface{}, conds ...interface{}) (tx *DB)
```

First finds the first record ordered by primary key, matching given conditions conds

#### (*DB) FirstOrCreate 

``` go
func (db *DB) FirstOrCreate(dest interface{}, conds ...interface{}) (tx *DB)
```

FirstOrCreate finds the first matching record, otherwise if not found creates a new instance with given conds. Each conds must be a struct or map.

Using FirstOrCreate in conjunction with Assign will result in an update to the database even if the record exists.

```go
// assign an email if the record is not found
result := db.Where(User{Name: "non_existing"}).Attrs(User{Email: "fake@fake.org"}).FirstOrCreate(&user)
// user -> User{Name: "non_existing", Email: "fake@fake.org"}
// result.RowsAffected -> 1

// assign email regardless of if record is found
result := db.Where(User{Name: "jinzhu"}).Assign(User{Email: "fake@fake.org"}).FirstOrCreate(&user)
// user -> User{Name: "jinzhu", Age: 20, Email: "fake@fake.org"}
// result.RowsAffected -> 1
```

#### (*DB) FirstOrInit 

``` go
func (db *DB) FirstOrInit(dest interface{}, conds ...interface{}) (tx *DB)
```

FirstOrInit finds the first matching record, otherwise if not found initializes a new instance with given conds. Each conds must be a struct or map.

FirstOrInit never modifies the database. It is often used with Assign and Attrs.

```go
// assign an email if the record is not found
db.Where(User{Name: "non_existing"}).Attrs(User{Email: "fake@fake.org"}).FirstOrInit(&user)
// user -> User{Name: "non_existing", Email: "fake@fake.org"}

// assign email regardless of if record is found
db.Where(User{Name: "jinzhu"}).Assign(User{Email: "fake@fake.org"}).FirstOrInit(&user)
// user -> User{Name: "jinzhu", Age: 20, Email: "fake@fake.org"}
```

#### (*DB) Get 

``` go
func (db *DB) Get(key string) (interface{}, bool)
```

Get get value with key from current db instance's context

#### (*DB) Group 

``` go
func (db *DB) Group(name string) (tx *DB)
```

Group specify the group method on the find

```
// Select the sum age of users with given names
db.Model(&User{}).Select("name, sum(age) as total").Group("name").Find(&results)
```

#### (*DB) Having 

``` go
func (db *DB) Having(query interface{}, args ...interface{}) (tx *DB)
```

Having specify HAVING conditions for GROUP BY

```
// Select the sum age of users with name jinzhu
db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "jinzhu").Find(&result)
```

#### (*DB) InnerJoins <- 1.24.3

``` go
func (db *DB) InnerJoins(query string, args ...interface{}) (tx *DB)
```

InnerJoins specify inner joins conditions db.InnerJoins("Account").Find(&user)

#### (*DB) InstanceGet 

``` go
func (db *DB) InstanceGet(key string) (interface{}, bool)
```

InstanceGet get value with key from current db instance's context

#### (*DB) InstanceSet 

``` go
func (db *DB) InstanceSet(key string, value interface{}) *DB
```

InstanceSet store value with key into current db instance's context

#### (*DB) Joins 

``` go
func (db *DB) Joins(query string, args ...interface{}) (tx *DB)
```

Joins specify Joins conditions

```go
db.Joins("Account").Find(&user)
db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
db.Joins("Account", DB.Select("id").Where("user_id = users.id AND name = ?", "someName").Model(&Account{}))
```

#### (*DB) Last 

``` go
func (db *DB) Last(dest interface{}, conds ...interface{}) (tx *DB)
```

Last finds the last record ordered by primary key, matching given conditions conds

#### (*DB) Limit 

``` go
func (db *DB) Limit(limit int) (tx *DB)
```

Limit specify the number of records to be retrieved

Limit conditions can be cancelled by using `Limit(-1)`.

```go
// retrieve 3 users
db.Limit(3).Find(&users)
// retrieve 3 users into users1, and all users into users2
db.Limit(3).Find(&users1).Limit(-1).Find(&users2)
```

#### (*DB) Migrator 

``` go
func (db *DB) Migrator() Migrator
```

Migrator returns migrator

#### (*DB) Model 

``` go
func (db *DB) Model(value interface{}) (tx *DB)
```

Model specify the model you would like to run db operations

```go
// update all users's name to `hello`
db.Model(&User{}).Update("name", "hello")
// if user's primary key is non-blank, will use it as condition, then will only update that user's name to `hello`
db.Model(&user).Update("name", "hello")
```

#### (*DB) Not 

``` go
func (db *DB) Not(query interface{}, args ...interface{}) (tx *DB)
```

Not add NOT conditions

Not works similarly to where, and has the same syntax.

```
// Find the first user with name not equal to jinzhu
db.Not("name = ?", "jinzhu").First(&user)
```

#### (*DB) Offset 

``` go
func (db *DB) Offset(offset int) (tx *DB)
```

Offset specify the number of records to skip before starting to return the records

Offset conditions can be cancelled by using `Offset(-1)`.

```go
// select the third user
db.Offset(2).First(&user)
// select the first user by cancelling an earlier chained offset
db.Offset(5).Offset(-1).First(&user)
```

#### (*DB) Omit 

``` go
func (db *DB) Omit(columns ...string) (tx *DB)
```

Omit specify fields that you want to ignore when creating, updating and querying

#### (*DB) Or 

``` go
func (db *DB) Or(query interface{}, args ...interface{}) (tx *DB)
```

Or add OR conditions

Or is used to chain together queries with an OR.

```go
// Find the first user with name equal to jinzhu or john
db.Where("name = ?", "jinzhu").Or("name = ?", "john").First(&user)
```

#### (*DB) Order 

``` go
func (db *DB) Order(value interface{}) (tx *DB)
```

Order specify order when retrieving records from database

```go
db.Order("name DESC")
db.Order(clause.OrderByColumn{Column: clause.Column{Name: "name"}, Desc: true})
```

#### (*DB) Pluck 

``` go
func (db *DB) Pluck(column string, dest interface{}) (tx *DB)
```

Pluck queries a single column from a model, returning in the slice dest. E.g.:

``` go
var ages []int64
db.Model(&users).Pluck("age", &ages)
```

#### (*DB) Preload 

``` go
func (db *DB) Preload(query string, args ...interface{}) (tx *DB)
```

Preload preload associations with given conditions

```go
// get all users, and preload all non-cancelled orders
db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
```

#### (*DB) Raw 

``` go
func (db *DB) Raw(sql string, values ...interface{}) (tx *DB)
```

#### (*DB) Rollback 

``` go
func (db *DB) Rollback() *DB
```

Rollback rollbacks the changes in a transaction

#### (*DB) RollbackTo <- 0.2.9

``` go
func (db *DB) RollbackTo(name string) *DB
```

#### (*DB) Row 

``` go
func (db *DB) Row() *sql.Row
```

#### (*DB) Rows 

``` go
func (db *DB) Rows() (*sql.Rows, error)
```

#### (*DB) Save 

``` go
func (db *DB) Save(value interface{}) (tx *DB)
```

Save updates value in database. If value doesn't contain a matching primary key, value is inserted.

#### (*DB) SavePoint <- 0.2.9

``` go
func (db *DB) SavePoint(name string) *DB
```

#### (*DB) Scan 

``` go
func (db *DB) Scan(dest interface{}) (tx *DB)
```

Scan scans selected value to the struct dest

#### (*DB) ScanRows 

``` go
func (db *DB) ScanRows(rows *sql.Rows, dest interface{}) error
```

#### (*DB) Scopes 

``` go
func (db *DB) Scopes(funcs ...func(*DB) *DB) (tx *DB)
```

Scopes pass current database connection to arguments `func(DB) DB`, which could be used to add conditions dynamically

``` go
func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
    return db.Where("amount > ?", 1000)
}

func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
    return func (db *gorm.DB) *gorm.DB {
        return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
    }
}

db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
```

#### (*DB) Select 

``` go
func (db *DB) Select(query interface{}, args ...interface{}) (tx *DB)
```

Select specify fields that you want when querying, creating, updating

Use Select when you only want a subset of the fields. By default, GORM will select all fields. Select accepts both string arguments and arrays.

```go
// Select name and age of user using multiple arguments
db.Select("name", "age").Find(&users)
// Select name and age of user using an array
db.Select([]string{"name", "age"}).Find(&users)
```

#### (*DB) Session 

``` go
func (db *DB) Session(config *Session) *DB
```

Session create new db session

#### (*DB) Set 

``` go
func (db *DB) Set(key string, value interface{}) *DB
```

Set store value with key into current db instance's context

#### (*DB) SetupJoinTable 

``` go
func (db *DB) SetupJoinTable(model interface{}, field string, joinTable interface{}) error
```

SetupJoinTable setup join table schema

#### (*DB) Table 

``` go
func (db *DB) Table(name string, args ...interface{}) (tx *DB)
```

Table specify the table you would like to run db operations

```
// Get a user
db.Table("users").Take(&result)
```

#### (*DB) Take 

``` go
func (db *DB) Take(dest interface{}, conds ...interface{}) (tx *DB)
```

Take finds the first record returned by the database in no specified order, matching given conditions conds

#### (*DB) ToSQL <- 1.22.3

``` go
func (db *DB) ToSQL(queryFn func(tx *DB) *DB) string
```

ToSQL for generate SQL string.

```go
db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(&User{}).Where(&User{Name: "foo", Age: 20})
			.Limit(10).Offset(5)
			.Order("name ASC")
			.First(&User{})
})
```

#### (*DB) Transaction 

``` go
func (db *DB) Transaction(fc func(tx *DB) error, opts ...*sql.TxOptions) (err error)
```

Transaction start a transaction as a block, return error will rollback, otherwise to commit. Transaction executes an arbitrary number of commands in fc within a transaction. On success the changes are committed; if an error occurs they are rolled back.

#### (*DB) Unscoped 

``` go
func (db *DB) Unscoped() (tx *DB)
```

#### (*DB) Update 

``` go
func (db *DB) Update(column string, value interface{}) (tx *DB)
```

Update updates column with value using callbacks. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields

#### (*DB) UpdateColumn 

``` go
func (db *DB) UpdateColumn(column string, value interface{}) (tx *DB)
```

#### (*DB) UpdateColumns 

``` go
func (db *DB) UpdateColumns(values interface{}) (tx *DB)
```

#### (*DB) Updates 

``` go
func (db *DB) Updates(values interface{}) (tx *DB)
```

Updates updates attributes using callbacks. values must be a struct or map. Reference: https://gorm.io/docs/update.html#Update-Changed-Fields

#### (*DB) Use <- 0.2.13

``` go
func (db *DB) Use(plugin Plugin) error
```

Use use plugin

#### (*DB) Where 

``` go
func (db *DB) Where(query interface{}, args ...interface{}) (tx *DB)
```

Where add conditions

See the [docs](https://gorm.io/docs/query.html#Conditions) for details on the various formats that where clauses can take. By default, where clauses chain with AND.

```go
// Find the first user with name jinzhu
db.Where("name = ?", "jinzhu").First(&user)
// Find the first user with name jinzhu and age 20
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
// Find the first user with name jinzhu and age not equal to 20
db.Where("name = ?", "jinzhu").Where("age <> ?", "20").First(&user)
```

#### (*DB) WithContext 

``` go
func (db *DB) WithContext(ctx context.Context) *DB
```

WithContext change current instance db's context to ctx

### type DeletedAt 

``` go
type DeletedAt sql.NullTime
```

#### (DeletedAt) DeleteClauses 

``` go
func (DeletedAt) DeleteClauses(f *schema.Field) []clause.Interface
```

#### (DeletedAt) MarshalJSON <- 1.20.3

``` go
func (n DeletedAt) MarshalJSON() ([]byte, error)
```

#### (DeletedAt) QueryClauses 

``` go
func (DeletedAt) QueryClauses(f *schema.Field) []clause.Interface
```

#### (*DeletedAt) Scan 

``` go
func (n *DeletedAt) Scan(value interface{}) error
```

Scan implements the Scanner interface.

#### (*DeletedAt) UnmarshalJSON <- 1.20.3

``` go
func (n *DeletedAt) UnmarshalJSON(b []byte) error
```

#### (DeletedAt) UpdateClauses <- 1.21.11

``` go
func (DeletedAt) UpdateClauses(f *schema.Field) []clause.Interface
```

#### (DeletedAt) Value 

``` go
func (n DeletedAt) Value() (driver.Value, error)
```

Value implements the driver Valuer interface.

### type Dialector 

``` go
type Dialector interface {
	Name() string
	Initialize(*DB) error
	Migrator(db *DB) Migrator
	DataTypeOf(*schema.Field) string
	DefaultValueOf(*schema.Field) clause.Expression
	BindVarTo(writer clause.Writer, stmt *Statement, v interface{})
	QuoteTo(clause.Writer, string)
	Explain(sql string, vars ...interface{}) string
}
```

Dialector GORM database dialector

### type ErrorTranslator <- 1.25.0

``` go
type ErrorTranslator interface {
	Translate(err error) error
}
```

### type GetDBConnector <- 1.21.5

``` go
type GetDBConnector interface {
	GetDBConn() (*sql.DB, error)
}
```

GetDBConnector SQL db connector

### type Index <- 1.23.7

``` go
type Index interface {
	Table() string
	Name() string
	Columns() []string
	PrimaryKey() (isPrimaryKey bool, ok bool)
	Unique() (unique bool, ok bool)
	Option() string
}
```

### type Migrator 

``` go
type Migrator interface {
	// AutoMigrate
	AutoMigrate(dst ...interface{}) error

	// Database
	CurrentDatabase() string
	FullDataTypeOf(*schema.Field) clause.Expr
	GetTypeAliases(databaseTypeName string) []string

	// Tables
	CreateTable(dst ...interface{}) error
	DropTable(dst ...interface{}) error
	HasTable(dst interface{}) bool
	RenameTable(oldName, newName interface{}) error
	GetTables() (tableList []string, err error)
	TableType(dst interface{}) (TableType, error)

	// Columns
	AddColumn(dst interface{}, field string) error
	DropColumn(dst interface{}, field string) error
	AlterColumn(dst interface{}, field string) error
	MigrateColumn(dst interface{}, field *schema.Field, columnType ColumnType) error
	HasColumn(dst interface{}, field string) bool
	RenameColumn(dst interface{}, oldName, field string) error
	ColumnTypes(dst interface{}) ([]ColumnType, error)

	// Views
	CreateView(name string, option ViewOption) error
	DropView(name string) error

	// Constraints
	CreateConstraint(dst interface{}, name string) error
	DropConstraint(dst interface{}, name string) error
	HasConstraint(dst interface{}, name string) bool

	// Indexes
	CreateIndex(dst interface{}, name string) error
	DropIndex(dst interface{}, name string) error
	HasIndex(dst interface{}, name string) bool
	RenameIndex(dst interface{}, oldName, newName string) error
	GetIndexes(dst interface{}) ([]Index, error)
}
```

Migrator migrator interface

### type Model 

``` go
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
}
```

Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt It may be embedded into your model or you may build your own model without it

``` go
type User struct {
  gorm.Model
}
```

### type Option <- 1.21.0

``` go
type Option interface {
	Apply(*Config) error
	AfterInitialize(*DB) error
}
```

Option gorm option interface

### type ParamsFilter <- 1.24.3

``` go
type ParamsFilter interface {
	ParamsFilter(ctx context.Context, sql string, params ...interface{}) (string, []interface{})
}
```

### type Plugin <- 0.2.13

``` go
type Plugin interface {
	Name() string
	Initialize(*DB) error
}
```

Plugin GORM plugin interface

### type PreparedStmtDB <- 0.2.3

``` go
type PreparedStmtDB struct {
	Stmts       map[string]*Stmt
	PreparedSQL []string
	Mux         *sync.RWMutex
	ConnPool
}
```

#### func NewPreparedStmtDB <- 1.25.2

``` go
func NewPreparedStmtDB(connPool ConnPool) *PreparedStmtDB
```

#### (*PreparedStmtDB) BeginTx <- 0.2.3

``` go
func (db *PreparedStmtDB) BeginTx(ctx context.Context, opt *sql.TxOptions) (ConnPool, error)
```

#### (*PreparedStmtDB) Close <- 0.2.13

``` go
func (db *PreparedStmtDB) Close()
```

#### (*PreparedStmtDB) ExecContext <- 0.2.3

``` go
func (db *PreparedStmtDB) ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error)
```

#### (*PreparedStmtDB) GetDBConn <- 1.21.6

``` go
func (db *PreparedStmtDB) GetDBConn() (*sql.DB, error)
```

#### (*PreparedStmtDB) QueryContext <- 0.2.3

``` go
func (db *PreparedStmtDB) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error)
```

#### (*PreparedStmtDB) QueryRowContext <- 0.2.3

``` go
func (db *PreparedStmtDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
```

#### (*PreparedStmtDB) Reset <- 1.24.1

``` go
func (sdb *PreparedStmtDB) Reset()
```

### type PreparedStmtTX <- 0.2.3

``` go
type PreparedStmtTX struct {
	Tx
	PreparedStmtDB *PreparedStmtDB
}
```

#### (*PreparedStmtTX) Commit <- 1.20.1

``` go
func (tx *PreparedStmtTX) Commit() error
```

#### (*PreparedStmtTX) ExecContext <- 0.2.3

``` go
func (tx *PreparedStmtTX) ExecContext(ctx context.Context, query string, args ...interface{}) (result sql.Result, err error)
```

#### (*PreparedStmtTX) GetDBConn <- 1.25.4

``` go
func (db *PreparedStmtTX) GetDBConn() (*sql.DB, error)
```

#### (*PreparedStmtTX) QueryContext <- 0.2.3

``` go
func (tx *PreparedStmtTX) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error)
```

#### (*PreparedStmtTX) QueryRowContext <- 0.2.3

``` go
func (tx *PreparedStmtTX) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
```

#### (*PreparedStmtTX) Rollback <- 1.20.1

``` go
func (tx *PreparedStmtTX) Rollback() error
```

### type Rows <- 1.23.4

``` go
type Rows interface {
	Columns() ([]string, error)
	ColumnTypes() ([]*sql.ColumnType, error)
	Next() bool
	Scan(dest ...interface{}) error
	Err() error
	Close() error
}
```

Rows rows interface

### type SavePointerDialectorInterface <- 0.2.9

``` go
type SavePointerDialectorInterface interface {
	SavePoint(tx *DB, name string) error
	RollbackTo(tx *DB, name string) error
}
```

SavePointerDialectorInterface save pointer interface

### type ScanMode <- 1.22.0

``` go
type ScanMode uint8
```

ScanMode scan data mode

``` go
const (
	ScanInitialized         ScanMode = 1 << 0 // 1
	ScanUpdate              ScanMode = 1 << 1 // 2
	ScanOnConflictDoNothing ScanMode = 1 << 2 // 4
)
```

scan modes

### type Session 

``` go
type Session struct {
	DryRun                   bool
	PrepareStmt              bool
	NewDB                    bool
	Initialized              bool
	SkipHooks                bool
	SkipDefaultTransaction   bool
	DisableNestedTransaction bool
	AllowGlobalUpdate        bool
	FullSaveAssociations     bool
	QueryFields              bool
	Context                  context.Context
	Logger                   logger.Interface
	NowFunc                  func() time.Time
	CreateBatchSize          int
}
```

Session session config when create session with Session() method

### type SoftDeleteDeleteClause <- 0.2.32

``` go
type SoftDeleteDeleteClause struct {
	ZeroValue sql.NullString
	Field     *schema.Field
}
```

#### (SoftDeleteDeleteClause) Build <- 0.2.32

``` go
func (sd SoftDeleteDeleteClause) Build(clause.Builder)
```

#### (SoftDeleteDeleteClause) MergeClause <- 0.2.32

``` go
func (sd SoftDeleteDeleteClause) MergeClause(*clause.Clause)
```

#### (SoftDeleteDeleteClause) ModifyStatement <- 0.2.32

``` go
func (sd SoftDeleteDeleteClause) ModifyStatement(stmt *Statement)
```

#### (SoftDeleteDeleteClause) Name <- 0.2.32

``` go
func (sd SoftDeleteDeleteClause) Name() string
```

### type SoftDeleteQueryClause <- 0.2.32

``` go
type SoftDeleteQueryClause struct {
	ZeroValue sql.NullString
	Field     *schema.Field
}
```

#### (SoftDeleteQueryClause) Build <- 0.2.32

``` go
func (sd SoftDeleteQueryClause) Build(clause.Builder)
```

#### (SoftDeleteQueryClause) MergeClause <- 0.2.32

``` go
func (sd SoftDeleteQueryClause) MergeClause(*clause.Clause)
```

#### (SoftDeleteQueryClause) ModifyStatement <- 0.2.32

``` go
func (sd SoftDeleteQueryClause) ModifyStatement(stmt *Statement)
```

#### (SoftDeleteQueryClause) Name <- 0.2.32

``` go
func (sd SoftDeleteQueryClause) Name() string
```

### type SoftDeleteUpdateClause <- 1.21.11

``` go
type SoftDeleteUpdateClause struct {
	ZeroValue sql.NullString
	Field     *schema.Field
}
```

#### (SoftDeleteUpdateClause) Build <- 1.21.11

``` go
func (sd SoftDeleteUpdateClause) Build(clause.Builder)
```

#### (SoftDeleteUpdateClause) MergeClause <- 1.21.11

``` go
func (sd SoftDeleteUpdateClause) MergeClause(*clause.Clause)
```

#### (SoftDeleteUpdateClause) ModifyStatement <- 1.21.11

``` go
func (sd SoftDeleteUpdateClause) ModifyStatement(stmt *Statement)
```

#### (SoftDeleteUpdateClause) Name <- 1.21.11

``` go
func (sd SoftDeleteUpdateClause) Name() string
```

### type Statement 

``` go
type Statement struct {
	*DB
	TableExpr            *clause.Expr
	Table                string
	Model                interface{}
	Unscoped             bool
	Dest                 interface{}
	ReflectValue         reflect.Value
	Clauses              map[string]clause.Clause
	BuildClauses         []string
	Distinct             bool
	Selects              []string // selected columns
	Omits                []string // omit columns
	Joins                []join
	Preloads             map[string][]interface{}
	Settings             sync.Map
	ConnPool             ConnPool
	Schema               *schema.Schema
	Context              context.Context
	RaiseErrorOnNotFound bool
	SkipHooks            bool
	SQL                  strings.Builder
	Vars                 []interface{}
	CurDestIndex         int
	// contains filtered or unexported fields
}
```

Statement statement

#### (*Statement) AddClause 

``` go
func (stmt *Statement) AddClause(v clause.Interface)
```

AddClause add clause

#### (*Statement) AddClauseIfNotExists 

``` go
func (stmt *Statement) AddClauseIfNotExists(v clause.Interface)
```

AddClauseIfNotExists add clause if not exists

#### (*Statement) AddVar 

``` go
func (stmt *Statement) AddVar(writer clause.Writer, vars ...interface{})
```

AddVar add var

#### (*Statement) Build 

``` go
func (stmt *Statement) Build(clauses ...string)
```

Build build sql with clauses names

#### (*Statement) BuildCondition <- 0.2.5

``` go
func (stmt *Statement) BuildCondition(query interface{}, args ...interface{}) []clause.Expression
```

BuildCondition build condition

#### (*Statement) Changed <- 0.2.19

``` go
func (stmt *Statement) Changed(fields ...string) bool
```

Changed check model changed or not when updating

#### (*Statement) Parse 

``` go
func (stmt *Statement) Parse(value interface{}) (err error)
```

#### (*Statement) ParseWithSpecialTableName <- 1.22.0

``` go
func (stmt *Statement) ParseWithSpecialTableName(value interface{}, specialTableName string) (err error)
```

#### (*Statement) Quote 

``` go
func (stmt *Statement) Quote(field interface{}) string
```

Quote returns quoted value

#### (*Statement) QuoteTo 

``` go
func (stmt *Statement) QuoteTo(writer clause.Writer, field interface{})
```

QuoteTo write quoted value to writer

#### (*Statement) SelectAndOmitColumns <- 0.2.19

``` go
func (stmt *Statement) SelectAndOmitColumns(requireCreate, requireUpdate bool) (map[string]bool, bool)
```

SelectAndOmitColumns get select and omit columns, select -> true, omit -> false

#### (*Statement) SetColumn <- 0.2.19

``` go
func (stmt *Statement) SetColumn(name string, value interface{}, fromCallbacks ...bool)
```

SetColumn set column's value

```
stmt.SetColumn("Name", "jinzhu") // Hooks Method
stmt.SetColumn("Name", "jinzhu", true) // Callbacks Method
```

#### (*Statement) WriteByte 

``` go
func (stmt *Statement) WriteByte(c byte) error
```

WriteByte write byte

#### (*Statement) WriteQuoted 

``` go
func (stmt *Statement) WriteQuoted(value interface{})
```

WriteQuoted write quoted value

#### (*Statement) WriteString 

``` go
func (stmt *Statement) WriteString(str string) (int, error)
```

WriteString write string

### type StatementModifier 

``` go
type StatementModifier interface {
	ModifyStatement(*Statement)
}
```

StatementModifier statement modifier interface

### type Stmt <- 1.20.11

``` go
type Stmt struct {
	*sql.Stmt
	Transaction bool
	// contains filtered or unexported fields
}
```

### type TableType <- 1.25.1

``` go
type TableType interface {
	Schema() string
	Name() string
	Type() string
	Comment() (comment string, ok bool)
}
```

TableType table type interface

### type Tx <- 1.23.2

``` go
type Tx interface {
	ConnPool
	TxCommitter
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}
```

Tx sql.Tx interface

### type TxBeginner 

``` go
type TxBeginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}
```

TxBeginner tx beginner

### type TxCommitter <- 0.2.3

``` go
type TxCommitter interface {
	Commit() error
	Rollback() error
}
```

TxCommitter tx committer

### type Valuer <- 0.2.38

``` go
type Valuer interface {
	GormValue(context.Context, *DB) clause.Expr
}
```

Valuer gorm valuer interface

### type ViewOption 

``` go
type ViewOption struct {
	Replace     bool   // If true, exec `CREATE`. If false, exec `CREATE OR REPLACE`
	CheckOption string // optional. e.g. `WITH [ CASCADED | LOCAL ] CHECK OPTION`
	Query       *DB    // required subquery.
}
```

ViewOption view option