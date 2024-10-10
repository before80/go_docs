+++
title = "创建"
date = 2023-10-28T14:25:34+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/create.html](https://gorm.io/docs/create.html)

## 创建记录 Create Record

```go
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

result := db.Create(&user) // 将数据指针传递给Create pass pointer of data to Create

user.ID             // 返回插入数据的主键 returns inserted data's primary key
result.Error        // 返回错误 returns error
result.RowsAffected // 返回插入的记录数 returns inserted records count
```

We can also create multiple records with `Create()`:

​	我们还可以使用`Create()`创建多个记录：

```go
users := []*User{
  User{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
  User{Name: "Jackson", Age: 19, Birthday: time.Now()},
}

result := db.Create(users) // 传递一个切片来插入多行 pass a slice to insert multiple row

result.Error        // 返回错误 returns error
result.RowsAffected // 返回插入的记录数 returns inserted records count
```

> **NOTE** You cannot pass a struct to ‘create’, so you should pass a pointer to the data.
>
> **注意**：你不能将结构体直接传递给'create'，所以你应该传递数据指针。

## 创建选定字段的记录 Create Record With Selected Fields

Create a record and assign a value to the fields specified.

​	创建一个记录，并为指定的字段分配值。

```go
db.Select("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
```

Create a record and ignore the values for fields passed to omit.  

​	创建一个记录，并忽略传递给 omit 的字段的值。

```go
db.Omit("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
```

## 批量插入 Batch Insert

To efficiently insert large number of records, pass a slice to the `Create` method. GORM will generate a single SQL statement to insert all the data and backfill primary key values, hook methods will be invoked too. It will begin a **transaction** when records can be splited into multiple batches.

​	为了高效地插入大量记录，可以将切片传递给 `Create` 方法。GORM 将生成一个单一的 SQL 语句来插入所有数据并回填主键值，同时也会调用钩子方法。当记录可以分成多个批次时，它将开始一个 **事务**。

```go
var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
db.Create(&users)

for _, user := range users {
  user.ID // 1,2,3
}
```

You can specify batch size when creating with `CreateInBatches`, e.g:   

​	在创建时可以使用 `CreateInBatches` 指定批处理大小，例如：

``` go
var users = []User{{Name: "jinzhu_1"}, ...., {Name: "jinzhu_10000"}}

// batch size 100
db.CreateInBatches(users, 100)
```

Batch Insert is also supported when using [Upsert](https://gorm.io/docs/create.html#upsert) and [Create With Associations](https://gorm.io/docs/create.html#create_with_associations)

​	在使用 [Upsert](https://gorm.io/docs/create.html#upsert) 和 [Create With Associations](https://gorm.io/docs/create.html#create_with_associations) 时也支持批量插入。

> **NOTE** initialize GORM with `CreateBatchSize` option, all `INSERT` will respect this option when creating record & associations
>
> **注意**：使用 `CreateBatchSize` 选项初始化 GORM，所有的 `INSERT` 在创建记录和关联关系时都会尊重这个选项。

```go
db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
  CreateBatchSize: 1000,
})

db := db.Session(&gorm.Session{CreateBatchSize: 1000})

users = [5000]User{{Name: "jinzhu", Pets: []Pet{pet1, pet2, pet3}}...}

db.Create(&users)
// INSERT INTO users xxx (5 batches)
// INSERT INTO pets xxx (15 batches)
```

## 创建钩子 Create Hooks

GORM allows user defined hooks to be implemented for `BeforeSave`, `BeforeCreate`, `AfterSave`, `AfterCreate`. These hook method will be called when creating a record, refer [Hooks](https://gorm.io/docs/hooks.html) for details on the lifecycle

​	GORM 允许用户定义 `BeforeSave`、`BeforeCreate`、`AfterSave` 和 `AfterCreate` 的钩子。这些钩子方法将在创建记录时被调用，关于生命周期的详细信息，请参阅 [Hooks]({{< ref "/gorm/Tutorials/hooks">}})。

```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.UUID = uuid.New()

  if u.Role == "admin" {
    return errors.New("invalid role")
  }
  return
}
```

If you want to skip `Hooks` methods, you can use the `SkipHooks` session mode, for example:  

​	如果你想跳过 `Hooks` 方法，可以使用 `SkipHooks` 会话模式，例如：

```go
DB.Session(&gorm.Session{SkipHooks: true}).Create(&user)

DB.Session(&gorm.Session{SkipHooks: true}).Create(&users)

DB.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)
```

## 从映射创建 Create From Map

GORM supports create from `map[string]interface{}` and `[]map[string]interface{}{}`, e.g:

​	GORM 支持从 `map[string]interface{}` 和 `[]map[string]interface{}{}` 创建，例如：

```go
db.Model(&User{}).Create(map[string]interface{}{
  "Name": "jinzhu", "Age": 18,
})

// batch insert from `[]map[string]interface{}{}`
// 批量插入来自 `[]map[string]interface{}{}`
db.Model(&User{}).Create([]map[string]interface{}{
  {"Name": "jinzhu_1", "Age": 18},
  {"Name": "jinzhu_2", "Age": 20},
})
```

> **NOTE** When creating from map, hooks won’t be invoked, associations won’t be saved and primary key values won’t be back filled
>
> **注意** 当从映射创建时，不会调用钩子，关联关系不会被保存，主键值也不会被回填

## 从 SQL 表达式/上下文值器创建 Create From SQL Expression/Context Valuer

GORM allows insert data with SQL expression, there are two ways to achieve this goal, create from `map[string]interface{}` or [Customized Data Types](https://gorm.io/docs/data_types.html#gorm_valuer_interface), for example:

​	GORM 允许使用 SQL 表达式插入数据，有两种实现目标的方法，一种是从 `map[string]interface{}`，另一种是[自定义数据类型]({{< ref "/gorm/Tutorials/customizeDataTypes#gormvaluerinterface">}})，例如：

```go
// Create from map
db.Model(User{}).Create(map[string]interface{}{
  "Name": "jinzhu",
  "Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
})
// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"));

// Create from customized data type
type Location struct {
  X, Y int
}

// Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
  // Scan a value into struct from database driver
}

func (loc Location) GormDataType() string {
  return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
  return clause.Expr{
    SQL:  "ST_PointFromText(?)",
    Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
  }
}

type User struct {
  Name     string
  Location Location
}

db.Create(&User{
  Name:     "jinzhu",
  Location: Location{X: 100, Y: 100},
})
// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"))
```

## 高级 Advanced

### 创建关联数据 Create With Associations

When creating some data with associations, if its associations value is not zero-value, those associations will be upserted, and its `Hooks` methods will be invoked.

​	当创建一些带有关联的数据时，如果关联值不是零值，那么这些关联将被插入或更新，并且会调用其 `Hooks` 方法。

``` go
type CreditCard struct {
  gorm.Model
  Number   string
  UserID   uint
}

type User struct {
  gorm.Model
  Name       string
  CreditCard CreditCard
}

db.Create(&User{
  Name: "jinzhu",
  CreditCard: CreditCard{Number: "411111111111"}
})
// INSERT INTO `users` ...
// INSERT INTO `credit_cards` ...
```

You can skip saving associations with `Select`, `Omit`, for example:

​	你可以使用 `Select`、`Omit` 跳过保存关联，例如：

``` go
db.Omit("CreditCard").Create(&user)

// skip all associations
db.Omit(clause.Associations).Create(&user)    
```

### 默认值 Default Values

You can define default values for fields with tag `default`, for example:

​	你可以为带有标签 `default` 的字段定义默认值，例如：

``` go
type User struct {
  ID   int64
  Name string `gorm:"default:galeone"`
  Age  int64  `gorm:"default:18"`
}
```

Then the default value *will be used* when inserting into the database for [zero-value](https://tour.golang.org/basics/12) fields    

​	然后，在插入数据库时，对于[零值]({{< ref "/docs/GoTour/Basics/PackagesVariablesAndFunctions#zero-values-零值">}})字段，将使用 *默认值*。

> **NOTE** Any zero value like `0`, `''`, `false` won’t be saved into the database for those fields defined default value, you might want to use pointer type or Scanner/Valuer to avoid this, for example:
>
> **注意** 任何零值，如 `0`、`''`、`false`，不会将这些字段的定义的默认值保存到数据库中，你可能想使用指针类型或 Scanner/Valuer 来避免这种情况，例如：
>
> ```go
> type User struct {
>   gorm.Model
>   Name string
>   Age  *int           `gorm:"default:18"`
>   Active sql.NullBool `gorm:"default:true"`
> }
> ```



> **NOTE** You have to setup the `default` tag for fields having default or virtual/generated value in database, if you want to skip a default value definition when migrating, you could use `default:(-)`, for example:
>
> **注意** 你必须为数据库中具有默认值或虚拟/生成值的字段设置 `default` 标签，如果你想在迁移时跳过默认值定义，可以使用 `default:(-)`，例如：
>
> ```go
> type User struct {
>   ID        string `gorm:"default:uuid_generate_v3()"` // db func
>   FirstName string
>   LastName  string
>   Age       uint8
>   FullName  string `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
> }
> ```

When using virtual/generated value, you might need to disable its creating/updating permission, check out [Field-Level Permission](https://gorm.io/docs/models.html#field_permission)

​	在使用虚拟/生成值时，你可能需要禁用其创建/更新权限，查看 [字段级权限控制]({{< ref "/gorm/GettingStarted/declaringModels#字段级权限控制-field-level-permission">}})。

### Upsert / On Conflict

GORM provides compatible Upsert support for different databases

​	GORM 为不同的数据库提供了兼容的 Upsert 支持。

``` go
import "gorm.io/gorm/clause"

// 在冲突时不执行任何操作 Do nothing on conflict
db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

// 在 `id` 冲突时将列更新为默认值 Update columns to default value on `id` conflict
db.Clauses(clause.OnConflict{
  Columns:   []clause.Column{{Name: "id"}},
  DoUpdates: clause.Assignments(map[string]interface{}{"role": "user"}),
}).Create(&users)
// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET ***; SQL Server
// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE ***; MySQL

// 使用 SQL 表达式 Use SQL expression
db.Clauses(clause.OnConflict{
  Columns:   []clause.Column{{Name: "id"}},
  DoUpdates: clause.Assignments(map[string]interface{}{"count": gorm.Expr("GREATEST(count, VALUES(count))")}),
}).Create(&users)
// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `count`=GREATEST(count, VALUES(count));

// 在 `id` 冲突时将列更新为新值 Update columns to new value on `id` conflict
db.Clauses(clause.OnConflict{
  Columns:   []clause.Column{{Name: "id"}},
  DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
}).Create(&users)
// MERGE INTO "users" USING *** WHEN NOT MATCHED THEN INSERT *** WHEN MATCHED THEN UPDATE SET "name"="excluded"."name"; SQL Server
// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age"; PostgreSQL
// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age); MySQL

// 在冲突时更新除主键和具有 SQL 函数默认值的列之外的所有列 Update all columns to new value on conflict except primary keys and those columns having default values from sql func
db.Clauses(clause.OnConflict{
  UpdateAll: true,
}).Create(&users)
// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age", ...;
// INSERT INTO `users` *** ON DUPLICATE KEY UPDATE `name`=VALUES(name),`age`=VALUES(age), ...; MySQL
```

Also checkout `FirstOrInit`, `FirstOrCreate` on [Advanced Query](https://gorm.io/docs/advanced_query.html)

​	还可以查看 [高级查询]({{< ref "/gorm/CRUDInterface/advancedQuery">}}) 中的 `FirstOrInit` 和 `FirstOrCreate`。

Checkout [Raw SQL and SQL Builder](https://gorm.io/docs/sql_builder.html) for more details

​	查看 [原始 SQL 和 SQL 构建器]({{< ref "/gorm/CRUDInterface/rawSQLAndSQLBuilder">}}) 以获取更多详细信息。