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

You can setup the default textformatter to display fulltimestamps like so

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