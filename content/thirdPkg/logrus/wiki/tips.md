+++
title = "tips"
date = 2023-06-25T09:41:44+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Tips

https://github.com/sirupsen/logrus/wiki/Tips

pavel edited this page on Jul 7, 2021 · [2 revisions](https://github.com/sirupsen/logrus/wiki/Tips/_history)

pavel 在 2021 年 7 月 7 日编辑了此页面 · [2 次修订](https://github.com/sirupsen/logrus/wiki/Tips/_history)



​	你可以设置默认的文本格式化器以显示完整的时间戳，像这样：

```go
package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	log.Println("Fulltimestamp here")
}
```