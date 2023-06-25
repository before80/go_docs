+++
title = "writer"
date = 2023-06-25T09:29:23+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Writer Hooks for Logrus

https://github.com/sirupsen/logrus/tree/v1.9.3/hooks/writer

Send logs of given levels to any object with `io.Writer` interface.

## Usage

If you want for example send high level logs to `Stderr` and logs of normal execution to `Stdout`, you could do it like this:

```
package main

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

func main() {
	log.SetOutput(ioutil.Discard) // Send all logs to nowhere by default

	log.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})
	log.AddHook(&writer.Hook{ // Send info and debug logs to stdout
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.InfoLevel,
			log.DebugLevel,
		},
	})
	log.Info("This will go to stdout")
	log.Warn("This will go to stderr")
}
```