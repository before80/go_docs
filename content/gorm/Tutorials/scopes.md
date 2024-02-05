+++
title = "范围 Scopes"
date = 2023-10-28T14:32:25+08:00
weight = 12
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/scopes.html](https://gorm.io/docs/scopes.html)

Scopes allow you to re-use commonly used logic, the shared logic needs to be defined as type `func(*gorm.DB) *gorm.DB`

​	范围允许你重用常用逻辑，共享逻辑需要定义为类型 `func(*gorm.DB) *gorm.DB`

## 查询 Query

Scope examples for querying

​	范围示例用于查询

``` go
func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
  return db.Where("amount > ?", 1000)
}

func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
  return db.Where("pay_mode = ?", "card")
}

func PaidWithCod(db *gorm.DB) *gorm.DB {
  return db.Where("pay_mode = ?", "cod")
}

func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    return db.Scopes(AmountGreaterThan1000).Where("status IN (?)", status)
  }
}

db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&orders)
// 查找所有信用卡订单且金额大于1000 Find all credit card orders and amount greater than 1000

db.Scopes(AmountGreaterThan1000, PaidWithCod).Find(&orders)
// 查找所有COD订单且金额大于1000 Find all COD orders and amount greater than 1000

db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
// 查找所有已支付、已发货的订单且金额大于1000 Find all paid, shipped orders that amount greater than 1000
```

### 分页 Pagination

``` go
func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    q := r.URL.Query()
    page, _ := strconv.Atoi(q.Get("page"))
    if page <= 0 {
      page = 1
    }

    pageSize, _ := strconv.Atoi(q.Get("page_size"))
    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := (page - 1) * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}

db.Scopes(Paginate(r)).Find(&users)
db.Scopes(Paginate(r)).Find(&articles)
```

## 动态表名 Dynamically Table

Use `Scopes` to dynamically set the query Table

​	使用 `Scopes` 动态设置查询表名

``` go
func TableOfYear(user *User, year int) func(db *gorm.DB) *gorm.DB {
  return func(db *gorm.DB) *gorm.DB {
    tableName := user.TableName() + strconv.Itoa(year)
    return db.Table(tableName)
  }
}

DB.Scopes(TableOfYear(user, 2019)).Find(&users)
// SELECT * FROM users_2019;

DB.Scopes(TableOfYear(user, 2020)).Find(&users)
// SELECT * FROM users_2020;

// 来自不同数据库的表 Table form different database
func TableOfOrg(user *User, dbName string) func(db *gorm.DB) *gorm.DB {
  return func(db *gorm.DB) *gorm.DB {
    tableName := dbName + "." + user.TableName()
    return db.Table(tableName)
  }
}

DB.Scopes(TableOfOrg(user, "org1")).Find(&users)
// SELECT * FROM org1.users;

DB.Scopes(TableOfOrg(user, "org2")).Find(&users)
// SELECT * FROM org2.users;
```

## 更新 Updates

Scope examples for updating/deleting

​	更新/删除的范围示例

``` go
func CurOrganization(r *http.Request) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    org := r.Query("org")

    if org != "" {
      var organization Organization
      if db.Session(&Session{}).First(&organization, "name = ?", org).Error == nil {
        return db.Where("org_id = ?", organization.ID)
      }
    }

    db.AddError("invalid organization")
    return db
  }
}

db.Model(&article).Scopes(CurOrganization(r)).Update("Name", "name 1")
// UPDATE articles SET name = "name 1" WHERE org_id = 111
db.Scopes(CurOrganization(r)).Delete(&Article{})
// DELETE FROM articles WHERE org_id = 111
```