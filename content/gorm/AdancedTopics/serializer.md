+++
title = "Serializer"
date = 2023-10-28T14:34:57+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gorm.io/docs/serializer.html](https://gorm.io/docs/serializer.html)

Serializer is an extensible interface that allows to customize how to serialize and deserialize data with databasae.

​	序列化器是一个可扩展的接口，允许自定义如何将数据与数据库进行序列化和反序列化。

GORM provides some default serializers: `json`, `gob`, `unixtime`, here is a quick example of how to use it.

​	GORM提供了一些默认的序列化器：`json`、`gob`、`unixtime`，这里是一个如何使用它的快速示例。

``` go
type User struct {
  Name        []byte                 `gorm:"serializer:json"`
  Roles       Roles                  `gorm:"serializer:json"`
  Contracts   map[string]interface{} `gorm:"serializer:json"`
  JobInfo     Job                    `gorm:"type:bytes;serializer:gob"`
  CreatedTime int64                  `gorm:"serializer:unixtime;type:time"` // 将int存储为datetime到数据库中 store int as datetime into database
}

type Roles []string

type Job struct {
  Title    string
  Location string
  IsIntern bool
}

createdAt := time.Date(2020, 1, 1, 0, 8, 0, 0, time.UTC)
data := User{
  Name:        []byte("jinzhu"),
  Roles:       []string{"admin", "owner"},
  Contracts:   map[string]interface{}{"name": "jinzhu", "age": 10},
  CreatedTime: createdAt.Unix(),
  JobInfo: Job{
    Title:    "Developer",
    Location: "NY",
    IsIntern: false,
  },
}

DB.Create(&data)
// INSERT INTO `users` (`name`,`roles`,`contracts`,`job_info`,`created_time`) VALUES
//   ("\"amluemh1\"","[\"admin\",\"owner\"]","{\"age\":10,\"name\":\"jinzhu\"}",<gob binary>,"2020-01-01 00:08:00")

var result User
DB.First(&result, "id = ?", data.ID)
// result => User{
//   Name:        []byte("jinzhu"),
//   Roles:       []string{"admin", "owner"},
//   Contracts:   map[string]interface{}{"name": "jinzhu", "age": 10},
//   CreatedTime: createdAt.Unix(),
//   JobInfo: Job{
//     Title:    "Developer",
//     Location: "NY",
//     IsIntern: false,
//   },
// }

DB.Where(User{Name: []byte("jinzhu")}).Take(&result)
// SELECT * FROM `users` WHERE `users`.`name` = "\"amluemh1\"
```

## 注册序列化器 Register Serializer

A Serializer needs to implement how to serialize and deserialize data, so it requires to implement the the following interface

​	一个序列化器需要实现如何序列化和反序列化数据，因此它需要实现以下接口：

``` go
import "gorm.io/gorm/schema"

type SerializerInterface interface {
  Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) error
  SerializerValuerInterface
}

type SerializerValuerInterface interface {
  Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error)
}
```

For example, the default `JSONSerializer` is implemented like:

​	例如，默认的`JSONSerializer`是像这样实现的：

``` go
// JSONSerializer json serializer
type JSONSerializer struct {
}

// Scan实现序列化器接口 Scan implements serializer interface
func (JSONSerializer) Scan(ctx context.Context, field *Field, dst reflect.Value, dbValue interface{}) (err error) {
  fieldValue := reflect.New(field.FieldType)

  if dbValue != nil {
    var bytes []byte
    switch v := dbValue.(type) {
    case []byte:
      bytes = v
    case string:
      bytes = []byte(v)
    default:
      return fmt.Errorf("failed to unmarshal JSONB value: %#v", dbValue)
    }

    err = json.Unmarshal(bytes, fieldValue.Interface())
  }

  field.ReflectValueOf(ctx, dst).Set(fieldValue.Elem())
  return
}

// Scan implements serializer interface
func (JSONSerializer) Value(ctx context.Context, field *Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
  return json.Marshal(fieldValue)
}
```

And registered with the following code:

​	并使用以下代码注册：

``` go
schema.RegisterSerializer("json", JSONSerializer{})
```

After registering a serializer, you can use it with the `serializer` tag, for example:

​	在注册序列化器后，您可以使用带有`serializer`标签的标签，例如：

``` go
type User struct {
  Name []byte `gorm:"serializer:json"`
}
```

## 自定义序列化器类型 Customized Serializer Type

You can use a registered serializer with tags, you are also allowed to create a customized struct that implements the above `SerializerInterface` and use it as a field type directly, for example:

​	您可以使用已注册的序列化器与标签一起使用，您还可以直接创建一个实现了上述`SerializerInterface`的结构体并将其用作字段类型，例如：

``` go
type EncryptedString string

// ctx: 包含请求范围的值 ctx: contains request-scoped values
// field: 使用序列化器的字段，包含GORM设置和结构体标签 field: the field using the serializer, contains GORM settings, struct tags
// dst: 当前模型的值，下面的示例中的user dst: current model value, `user` in the below example
// dbValue: 数据库中当前字段的值 dbValue: current field's value in database
func (es *EncryptedString) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
  switch value := dbValue.(type) {
  case []byte:
    *es = EncryptedString(bytes.TrimPrefix(value, []byte("hello")))
  case string:
    *es = EncryptedString(strings.TrimPrefix(value, "hello"))
  default:
    return fmt.Errorf("unsupported data %#v", dbValue)
  }
  return nil
}

// ctx: 包含请求范围的值 ctx: contains request-scoped values
// field: 使用序列化器的字段，包含GORM设置和结构体标签 field: the field using the serializer, contains GORM settings, struct tags
// dst: 当前模型的值，下面的示例中的user dst: current model value, `user` in the below example
// fieldValue: dst的当前字段值 fieldValue: current field's value of the dst
func (es EncryptedString) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
  return "hello" + string(es), nil
}

type User struct {
  gorm.Model
  Password EncryptedString
}

data := User{
  Password: EncryptedString("pass"),
}

DB.Create(&data)
// INSERT INTO `serializer_structs` (`password`) VALUES ("hellopass")

var result User
DB.First(&result, "id = ?", data.ID)
// result => User{
//   Password: EncryptedString("pass"),
// }

DB.Where(User{Password: EncryptedString("pass")}).Take(&result)
// SELECT * FROM `users` WHERE `users`.`password` = "hellopass"
```