+++
title = "提示"
date = 2023-10-28T14:35:35+08:00
weight = 6
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

[https://gorm.io/docs/hints.html](https://gorm.io/docs/hints.html)

GORM provides optimizer/index/comment hints support

​	GORM提供了优化器/索引/注释提示支持。

[https://github.com/go-gorm/hints](https://github.com/go-gorm/hints)

## 优化器提示 Optimizer Hints

``` go
import "gorm.io/hints"

db.Clauses(hints.New("hint")).Find(&User{})
// SELECT * /*+ hint */ FROM `users`
```

## 索引提示 Index Hints

``` go
import "gorm.io/hints"

db.Clauses(hints.UseIndex("idx_user_name")).Find(&User{})
// SELECT * FROM `users` USE INDEX (`idx_user_name`)

db.Clauses(hints.ForceIndex("idx_user_name", "idx_user_id").ForJoin()).Find(&User{})
// SELECT * FROM `users` FORCE INDEX FOR JOIN (`idx_user_name`,`idx_user_id`)"

db.Clauses(
  hints.ForceIndex("idx_user_name", "idx_user_id").ForOrderBy(),
  hints.IgnoreIndex("idx_user_name").ForGroupBy(),
).Find(&User{})
// SELECT * FROM `users` FORCE INDEX FOR ORDER BY (`idx_user_name`,`idx_user_id`) IGNORE INDEX FOR GROUP BY (`idx_user_name`)"
```

## 注释提示 Comment Hints

``` go
import "gorm.io/hints"

db.Clauses(hints.Comment("select", "master")).Find(&User{})
// SELECT /*master*/ * FROM `users`;

db.Clauses(hints.CommentBefore("insert", "node2")).Create(&user)
// /*node2*/ INSERT INTO `users` ...;

db.Clauses(hints.CommentAfter("select", "node2")).Create(&user)
// /*node2*/ INSERT INTO `users` ...;

db.Clauses(hints.CommentAfter("where", "hint")).Find(&User{}, "id = ?", 1)
// SELECT * FROM `users` WHERE id = ? /* hint */
```

