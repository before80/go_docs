+++
title = "slogtest"
date = 2023-11-05T14:33:11+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

https://pkg.go.dev/testing/slogtest@go1.21.3

## Overview 

Package slogtest implements support for testing implementations of log/slog.Handler.

### Example (Parsing)

This example demonstrates one technique for testing a handler with this package. The handler is given a [bytes.Buffer](https://pkg.go.dev/bytes#Buffer) to write to, and each line of the resulting output is parsed. For JSON output, [encoding/json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) produces a result in the desired format when given a pointer to a map[string]any.

``` go
package main

import (
	"bytes"
	"encoding/json"
	"log"
	"log/slog"
	"testing/slogtest"
)

func main() {
	var buf bytes.Buffer
	h := slog.NewJSONHandler(&buf, nil)

	results := func() []map[string]any {
		var ms []map[string]any
		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				panic(err) // In a real test, use t.Fatal.
			}
			ms = append(ms, m)
		}
		return ms
	}
	err := slogtest.TestHandler(h, results)
	if err != nil {
		log.Fatal(err)
	}

}

```
## Constants 

This section is empty.

## Variables 

This section is empty.

## Functions 

### func TestHandler 

``` go
func TestHandler(h slog.Handler, results func() []map[string]any) error
```

TestHandler tests a [slog.Handler](https://pkg.go.dev/log/slog#Handler). If TestHandler finds any misbehaviors, it returns an error for each, combined into a single error with errors.Join.

TestHandler installs the given Handler in a [slog.Logger](https://pkg.go.dev/log/slog#Logger) and makes several calls to the Logger's output methods.

The results function is invoked after all such calls. It should return a slice of map[string]any, one for each call to a Logger output method. The keys and values of the map should correspond to the keys and values of the Handler's output. Each group in the output should be represented as its own nested map[string]any. The standard keys slog.TimeKey, slog.LevelKey and slog.MessageKey should be used.

If the Handler outputs JSON, then calling [encoding/json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) with a `map[string]any` will create the right data structure.

If a Handler intentionally drops an attribute that is checked by a test, then the results function should check for its absence and add it to the map it returns.

## Types 

This section is empty.