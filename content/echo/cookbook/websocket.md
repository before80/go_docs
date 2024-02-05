+++
title = "websocket"
date = 2023-07-09T22:05:29+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# WebSocket

> 原文：[https://echo.labstack.com/docs/cookbook/websocket](https://echo.labstack.com/docs/cookbook/websocket)

## Using net WebSocket

### Server

cookbook/websocket/net/server.go

```go
package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
```



## Using gorilla WebSocket

### Server

cookbook/websocket/gorilla/server.go

```go
package main

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{}
)

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/ws", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
```



## Client

cookbook/websocket/public/index.html

```html
<!doctype html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <title>WebSocket</title>
</head>

<body>
  <p id="output"></p>

  <script>
    var loc = window.location;
    var uri = 'ws:';

    if (loc.protocol === 'https:') {
      uri = 'wss:';
    }
    uri += '//' + loc.host;
    uri += loc.pathname + 'ws';

    ws = new WebSocket(uri)

    ws.onopen = function() {
      console.log('Connected')
    }

    ws.onmessage = function(evt) {
      var out = document.getElementById('output');
      out.innerHTML += evt.data + '<br>';
    }

    setInterval(function() {
      ws.send('Hello, Server!');
    }, 1000);
  </script>
</body>

</html>
```



## Output

```sh
Hello, Client!
Hello, Client!
Hello, Client!
Hello, Client!
Hello, Client!
```



```sh
Hello, Server!
Hello, Server!
Hello, Server!
Hello, Server!
Hello, Server!
```