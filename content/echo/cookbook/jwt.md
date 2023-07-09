+++
title = "jwt"
date = 2023-07-09T22:03:38+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# JWT

https://echo.labstack.com/docs/cookbook/jwt

[JWT middleware](https://echo.labstack.com/docs/middleware/jwt) configuration can be found [here](https://echo.labstack.com/docs/middleware/jwt#configuration).

This is cookbook for:

- JWT authentication using HS256 algorithm.
- JWT is retrieved from `Authorization` request header.

## xxxxxxxxxx36 1<!DOCTYPE html>2<html>3​4<head>5    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />6    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css">7    <title>JSONP</title>8    <script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1/jquery.min.js"></script>9    <script type="text/javascript">10        var host_prefix = 'http://localhost:1323';11        $(document).ready(function() {12            // JSONP version - add 'callback=?' to the URL - fetch the JSONP response to the request13            $("#jsonp-button").click(function(e) {14                e.preventDefault();15                // The only difference on the client end is the addition of 'callback=?' to the URL16                var url = host_prefix + '/jsonp?callback=?';17                $.getJSON(url, function(jsonp) {18                    console.log(jsonp);19                    $("#jsonp-response").html(JSON.stringify(jsonp, null, 2));20                });21            });22        });23    </script>24​25</head>26​27<body>28    <div class="container" style="margin-top: 50px;">29        <input type="button" class="btn btn-primary btn-lg" id="jsonp-button" value="Get JSONP response">30        <p>31            <pre id="jsonp-response"></pre>32        </p>33    </div>34</body>35​36</html>html

### Using custom claims

cookbook/jwt/custom-claims/server.go

```go
package main

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Throws unauthorized error
	if username != "jon" || password != "shhh!" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))
	r.GET("", restricted)

	e.Logger.Fatal(e.Start(":1323"))
}
```



### Using a user-defined KeyFunc

cookbook/jwt/user-defined-keyfunc/server.go

```go
package main

import (
	"context"
	"errors"
	"fmt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lestrrat-go/jwx/jwk"
)

func getKey(token *jwt.Token) (interface{}, error) {

	// For a demonstration purpose, Google Sign-in is used.
	// https://developers.google.com/identity/sign-in/web/backend-auth
	//
	// This user-defined KeyFunc verifies tokens issued by Google Sign-In.
	//
	// Note: In this example, it downloads the keyset every time the restricted route is accessed.
	keySet, err := jwk.Fetch(context.Background(), "https://www.googleapis.com/oauth2/v3/certs")
	if err != nil {
		return nil, err
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have a key ID in the kid field")
	}

	key, found := keySet.LookupKeyID(keyID)

	if !found {
		return nil, fmt.Errorf("unable to find key %q", keyID)
	}

	var pubkey interface{}
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}

	return pubkey, nil
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Unauthenticated route
	e.GET("/", accessible)

	// Restricted group
	r := e.Group("/restricted")
	{
		config := echojwt.Config{
			KeyFunc: getKey,
		}
		r.Use(echojwt.WithConfig(config))
		r.GET("", restricted)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
```



## Client

### Login

Login using username and password to retrieve a token.

```sh
curl -X POST -d 'username=jon' -d 'password=shhh!' localhost:1323/login
```



### Response

```js
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"
}
```



### Request

Request a restricted resource using the token in `Authorization` request header.

```sh
curl localhost:1323/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjE5NTcxMzZ9.RB3arc4-OyzASAaUhC2W3ReWaXAt_z2Fd3BN4aWTgEY"
```



### Response

```sh
Welcome Jon Snow!
```