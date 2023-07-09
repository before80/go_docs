+++
title = "jaeger"
date = 2023-07-09T21:55:04+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Jaeger

https://echo.labstack.com/docs/middleware/jaeger

NOTE

Echo community contribution

Trace requests on Echo framework with Jaeger Tracing Middleware.

## Usage

```go
package main
import (
    "github.com/labstack/echo-contrib/jaegertracing"
    "github.com/labstack/echo/v4"
)
func main() {
    e := echo.New()
    // Enable tracing middleware
    c := jaegertracing.New(e, nil)
    defer c.Close()

    e.Logger.Fatal(e.Start(":1323"))
}
```



Enabling the tracing middleware creates a tracer and a root tracing span for every request.

## Custom Configuration

By default, traces are sent to `localhost` Jaeger agent instance. To configure an external Jaeger, start your application with environment variables.

### Usage

```bash
$ JAEGER_AGENT_HOST=192.168.1.10 JAEGER_AGENT_PORT=6831 ./myserver
```



The tracer can be initialized with values coming from environment variables. None of the env vars are required and all of them can be overriden via direct setting of the property on the configuration object.

| Property                         | Description                                                  |
| -------------------------------- | ------------------------------------------------------------ |
| JAEGER_SERVICE_NAME              | The service name                                             |
| JAEGER_AGENT_HOST                | The hostname for communicating with agent via UDP            |
| JAEGER_AGENT_PORT                | The port for communicating with agent via UDP                |
| JAEGER_ENDPOINT                  | The HTTP endpoint for sending spans directly to a collector, i.e. http://jaeger-collector:14268/api/traces |
| JAEGER_USER                      | Username to send as part of "Basic" authentication to the collector endpoint |
| JAEGER_PASSWORD                  | Password to send as part of "Basic" authentication to the collector endpoint |
| JAEGER_REPORTER_LOG_SPANS        | Whether the reporter should also log the spans               |
| JAEGER_REPORTER_MAX_QUEUE_SIZE   | The reporter's maximum queue size                            |
| JAEGER_REPORTER_FLUSH_INTERVAL   | The reporter's flush interval, with units, e.g. "500ms" or "2s" ([valid units][timeunits]) |
| JAEGER_SAMPLER_TYPE              | The sampler type                                             |
| JAEGER_SAMPLER_PARAM             | The sampler parameter (number)                               |
| JAEGER_SAMPLER_MANAGER_HOST_PORT | The HTTP endpoint when using the remote sampler, i.e. http://jaeger-agent:5778/sampling |
| JAEGER_SAMPLER_MAX_OPERATIONS    | The maximum number of operations that the sampler will keep track of |
| JAEGER_SAMPLER_REFRESH_INTERVAL  | How often the remotely controlled sampler will poll jaeger-agent for the appropriate sampling strategy, with units, e.g. "1m" or "30s" ([valid units][timeunits]) |
| JAEGER_TAGS                      | A comma separated list of `name = value` tracer level tags, which get added to all reported spans. The value can also refer to an environment variable using the format `${envVarName:default}`, where the `:default` is optional, and identifies a value to be used if the environment variable cannot be found |
| JAEGER_DISABLED                  | Whether the tracer is disabled or not. If true, the default `opentracing.NoopTracer` is used. |
| JAEGER_RPC_METRICS               | Whether to store RPC metrics                                 |

By default, the client sends traces via UDP to the agent at `localhost:6831`. Use `JAEGER_AGENT_HOST` and `JAEGER_AGENT_PORT` to send UDP traces to a different `host:port`. If `JAEGER_ENDPOINT` is set, the client sends traces to the endpoint via `HTTP`, making the `JAEGER_AGENT_HOST` and `JAEGER_AGENT_PORT` unused. If `JAEGER_ENDPOINT` is secured, HTTP basic authentication can be performed by setting the `JAEGER_USER` and `JAEGER_PASSWORD` environment variables.

### Skipping URL(s)

A middleware skipper can be passed to avoid tracing spans to certain URL(s).

*Usage*

```go
package main
import (
    "strings"
    "github.com/labstack/echo-contrib/jaegertracing"
    "github.com/labstack/echo/v4"
)

// urlSkipper ignores metrics route on some middleware
func urlSkipper(c echo.Context) bool {
    if strings.HasPrefix(c.Path(), "/testurl") {
        return true
    }
    return false
}

func main() {
    e := echo.New()
    // Enable tracing middleware
    c := jaegertracing.New(e, urlSkipper)
    defer c.Close()

    e.Logger.Fatal(e.Start(":1323"))
}
```



### TraceFunction

This is a wrapper function that can be used to seamlessly add a span for the duration of the invoked function. There is no need to change function arguments.

*Usage*

```go
package main
import (
    "github.com/labstack/echo-contrib/jaegertracing"
    "github.com/labstack/echo/v4"
    "net/http"
    "time"
)
func main() {
    e := echo.New()
    // Enable tracing middleware
    c := jaegertracing.New(e, nil)
    defer c.Close()
    e.GET("/", func(c echo.Context) error {
        // Wrap slowFunc on a new span to trace it's execution passing the function arguments
        jaegertracing.TraceFunction(c, slowFunc, "Test String")
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}

// A function to be wrapped. No need to change it's arguments due to tracing
func slowFunc(s string) {
    time.Sleep(200 * time.Millisecond)
    return
}
```



### xxxxxxxxxx23 1// Root level middleware2e.Use(middleware.Logger())3e.Use(middleware.Recover())4​5// Group level middleware6g := e.Group("/admin")7g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {8  if username == "joe" && password == "secret" {9    return true, nil10  }11  return false, nil12}))13​14// Route level middleware15track := func(next echo.HandlerFunc) echo.HandlerFunc {16    return func(c echo.Context) error {17        println("request to /users")18        return next(c)19    }20}21e.GET("/users", func(c echo.Context) error {22    return c.String(http.StatusOK, "/users")23}, track)go

For more control over the Span, the function `CreateChildSpan` can be called giving control on data to be appended to the span like log messages, baggages and tags.

*Usage*

```go
package main
import (
    "github.com/labstack/echo-contrib/jaegertracing"
    "github.com/labstack/echo/v4"
)
func main() {
    e := echo.New()
    // Enable tracing middleware
    c := jaegertracing.New(e, nil)
    defer c.Close()
    e.GET("/", func(c echo.Context) error {
        // Do something before creating the child span
        time.Sleep(40 * time.Millisecond)
        sp := jaegertracing.CreateChildSpan(c, "Child span for additional processing")
        defer sp.Finish()
        sp.LogEvent("Test log")
        sp.SetBaggageItem("Test baggage", "baggage")
        sp.SetTag("Test tag", "New Tag")
        time.Sleep(100 * time.Millisecond)
        return c.String(http.StatusOK, "Hello, World!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}
```



## References

- Opentracing Library: https://github.com/opentracing/opentracing-go
- Jaeger configuration: https://github.com/jaegertracing/jaeger-client-go#environment-variables