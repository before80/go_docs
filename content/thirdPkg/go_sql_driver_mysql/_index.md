+++
title = "go-sql-driver/mysql"
date = 2023-10-24T15:52:34+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://pkg.go.dev/github.com/go-sql-driver/mysql](https://pkg.go.dev/github.com/go-sql-driver/mysql)
>
> 版本：v1.7.1
>
> 发布时间：2023.4.25
>
> License： MPL-2.0

## Go-MySQL-Driver

A MySQL-Driver for Go's [database/sql](https://golang.org/pkg/database/sql/) package

​	这是一个用于Go的[database/sql](https://golang.org/pkg/database/sql/)包的MySQL驱动程序。

![Go-MySQL-Driver logo](_index_img/gomysql_m.png)

------

### 特点 Features

- Lightweight and [fast](https://github.com/go-sql-driver/sql-benchmark)
- 轻量级且[快速](https://github.com/go-sql-driver/sql-benchmark)
- Native Go implementation. No C-bindings, just pure Go
- 原生Go实现，无C绑定，仅使用纯Go
- Connections over TCP/IPv4, TCP/IPv6, Unix domain sockets or [custom protocols](https://godoc.org/github.com/go-sql-driver/mysql#DialFunc)
- 支持TCP/IPv4、TCP/IPv6、Unix域套接字或[自定义协议](https://godoc.org/github.com/go-sql-driver/mysql#DialFunc)连接
- Automatic handling of broken connections
- 自动处理断开连接
- Automatic Connection Pooling *(by database/sql package)*
- 自动连接池管理（由[database/sql](https://golang.org/pkg/database/sql/)包提供）
- Supports queries larger than 16MB
- 支持大于16MB的查询
- Full [`sql.RawBytes`](https://golang.org/pkg/database/sql/#RawBytes) support.
- 完全支持[sql.RawBytes]({{< ref "/stdLib/database/sql#type-rawbytes">}})功能
- Intelligent `LONG DATA` handling in prepared statements
- 预处理语句中的`LONG DATA`智能处理
- Secure `LOAD DATA LOCAL INFILE` support with file allowlisting and `io.Reader` support
- 安全支持`LOAD DATA LOCAL INFILE`，并支持文件白名单和`io.Reader`支持
- Optional `time.Time` parsing
- 可选`time.Time`解析
- Optional placeholder interpolation
- 可选占位符插值

### 要求 Requirements

- Go 1.13 or higher. We aim to support the 3 latest versions of Go.
- Go 1.13或更高版本。我们的目标是支持最新的3个版本的Go。
- MySQL (4.1+), MariaDB, Percona Server, Google CloudSQL or Sphinx (2.2.3+)
- MySQL（4.1+），MariaDB，Percona Server，Google CloudSQL或Sphinx（2.2.3+）

------

### 安装 Installation

Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH) with the [go tool](https://golang.org/cmd/go/) from shell:

​	使用shell中的[go tool]({{< ref "/cmd/go">}})将包安装到你的[$GOPATH](https://github.com/golang/go/wiki/GOPATH)：

```
$ go get -u github.com/go-sql-driver/mysql
```

Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.

​	确保在您的机器上[安装了Git](https://git-scm.com/downloads)，并在系统路径中设置了Git。

### 用法 Usage

*Go MySQL Driver* is an implementation of Go's `database/sql/driver` interface. You only need to import the driver and can use the full [`database/sql`](https://golang.org/pkg/database/sql/) API then.

​	*Go MySQL Driver* 是Go的`database/sql/driver`接口的实现。你只需要导入驱动，然后就可以使用完整的[`database/sql`]({{< ref "/stdLib/database/sql">}}) API了。

Use `mysql` as `driverName` and a valid [DSN](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-dsn-data-source-name) as `dataSourceName`:

​	使用`mysql`作为`driverName`和有效的[DSN](#dsn-data-source-name)作为`dataSourceName`：

```go
import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...

db, err := sql.Open("mysql", "user:password@/dbname")
if err != nil {
	panic(err)
}
// See "Important settings" section.
db.SetConnMaxLifetime(time.Minute * 3)
db.SetMaxOpenConns(10)
db.SetMaxIdleConns(10)
```

[Examples are available in our Wiki](https://github.com/go-sql-driver/mysql/wiki/Examples).

​	[我们的Wiki中提供了示例](https://github.com/go-sql-driver/mysql/wiki/Examples)。

#### Important settings

`db.SetConnMaxLifetime()` is required to ensure connections are closed by the driver safely before connection is closed by MySQL server, OS, or other middlewares. Since some middlewares close idle connections by 5 minutes, we recommend timeout shorter than 5 minutes. This setting helps load balancing and changing system variables too.

​	`db.SetConnMaxLifetime()` 是确保连接在MySQL服务器、操作系统或其他中间件关闭连接之前由驱动程序安全关闭的必要条件。由于某些中间件会在5分钟后关闭空闲连接，因此我们建议超时时间短于5分钟。此设置有助于负载平衡和更改系统变量。

`db.SetMaxOpenConns()` is highly recommended to limit the number of connection used by the application. There is no recommended limit number because it depends on application and MySQL server.

​	强烈建议使用 `db.SetMaxOpenConns()` 来限制应用程序使用的连接数量。没有推荐的限制数，因为这取决于应用程序和MySQL服务器。

`db.SetMaxIdleConns()` is recommended to be set same to `db.SetMaxOpenConns()`. When it is smaller than `SetMaxOpenConns()`, connections can be opened and closed much more frequently than you expect. Idle connections can be closed by the `db.SetConnMaxLifetime()`. If you want to close idle connections more rapidly, you can use `db.SetConnMaxIdleTime()` since Go 1.15.

​	建议将 `db.SetMaxIdleConns()` 设置为与 `db.SetMaxOpenConns()` 相同。当它小于 `SetMaxOpenConns()` 时，连接的打开和关闭频率可能会比您预期的要高。空闲连接可以通过 `db.SetConnMaxLifetime()` 来关闭。如果您想更快地关闭空闲连接，请使用自Go 1.15开始的 `db.SetConnMaxIdleTime()`。

#### DSN (Data Source Name)

The Data Source Name has a common format, like e.g. [PEAR DB](http://pear.php.net/manual/en/package.database.db.intro-dsn.php) uses it, but without type-prefix (optional parts marked by squared brackets):

​	数据源名称（DSN）具有常见的格式，例如[ PEAR DB](http://pear.php.net/manual/en/package.database.db.intro-dsn.php)使用它，但没有类型前缀（可选部分由方括号标记）：

```plaintext
[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
```

A DSN in its fullest form:

​	最完整形式的DSN：

```plaintext
username:password@protocol(address)/dbname?param=value
```

Except for the databasename, all values are optional. So the minimal DSN is:

​	除了数据库名称外，所有值都是可选的。因此，最小的DSN是：

```plaintext
/dbname
```

If you do not want to preselect a database, leave `dbname` empty:

​	如果您不想预先选择数据库，请将 `dbname` 留空：

```plaintext
/
```

This has the same effect as an empty DSN string:

​	这与空DSN字符串具有相同的效果：

```plaintext

```

Alternatively, [Config.FormatDSN](https://godoc.org/github.com/go-sql-driver/mysql#Config.FormatDSN) can be used to create a DSN string by filling a struct.

​	另外，可以通过填充结构体使用 [Config.FormatDSN](https://godoc.org/github.com/go-sql-driver/mysql#Config.FormatDSN) 来创建DSN字符串。

##### Password

Passwords can consist of any character. Escaping is **not** necessary.

​	密码可以由任何字符组成，不需要进行转义。

##### Protocol

See [net.Dial](https://golang.org/pkg/net/#Dial) for more information which networks are available. In general you should use an Unix domain socket if available and TCP otherwise for best performance.

​	有关可用网络的更多信息，请参见 [net.Dial](https://golang.org/pkg/net/#Dial)。一般来说，如果可用，您应该使用Unix域套接字以达到最佳性能。否则，可以使用TCP。

##### Address

For TCP and UDP networks, addresses have the form `host[:port]`. If `port` is omitted, the default port will be used. If `host` is a literal IPv6 address, it must be enclosed in square brackets. The functions [net.JoinHostPort](https://golang.org/pkg/net/#JoinHostPort) and [net.SplitHostPort](https://golang.org/pkg/net/#SplitHostPort) manipulate addresses in this form.

​	对于TCP和UDP网络，地址的格式为 `host[:port]`。如果省略了 `port`，则默认端口将被使用。如果 `host` 是字面IPv6地址，则必须将其括在方括号中。可以使用 [net.JoinHostPort](https://golang.org/pkg/net/#JoinHostPort) 和 [net.SplitHostPort](https://golang.org/pkg/net/#SplitHostPort) 函数来操作这种形式的地址。

For Unix domain sockets the address is the absolute path to the MySQL-Server-socket, e.g. `/var/run/mysqld/mysqld.sock` or `/tmp/mysql.sock`.

​	对于Unix域套接字，地址是MySQL服务器套接字的绝对路径，例如 `/var/run/mysqld/mysqld.sock` 或 `/tmp/mysql.sock`。

##### Parameters

*Parameters are case-sensitive!*

​	*参数是区分大小写的！*

Notice that any of `true`, `TRUE`, `True` or `1` is accepted to stand for a true boolean value. Not surprisingly, false can be specified as any of: `false`, `FALSE`, `False` or `0`.

​	请注意， `true`、`TRUE`、`True` 或 `1` 中的任何一个都可以表示为布尔值true。同样，false可以通过以下任何一种方式指定：`false`、`FALSE`、`False` 或 `0`。

###### allowAllFiles

```plaintext
Type:           bool
Valid Values:   true, false
Default:        false
```

`allowAllFiles=true` disables the file allowlist for `LOAD DATA LOCAL INFILE` and allows *all* files. [*Might be insecure!*](http://dev.mysql.com/doc/refman/5.7/en/load-data-local.html)

​	如果设置 `allowAllFiles=true`，将禁用 `LOAD DATA LOCAL INFILE` 的文件允许列表，并允许所有文件。 [*可能不安全！*](http://dev.mysql.com/doc/refman/5.7/en/load-data-local.html)

###### allowCleartextPasswords

```plaintext
Type:           bool
Valid Values:   true, false
Default:        false
```

`allowCleartextPasswords=true` allows using the [cleartext client side plugin](https://dev.mysql.com/doc/en/cleartext-pluggable-authentication.html) if required by an account, such as one defined with the [PAM authentication plugin](http://dev.mysql.com/doc/en/pam-authentication-plugin.html). Sending passwords in clear text may be a security problem in some configurations. To avoid problems if there is any possibility that the password would be intercepted, clients should connect to MySQL Server using a method that protects the password. Possibilities include [TLS / SSL](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-tls), IPsec, or a private network.

​	`allowCleartextPasswords=true`允许在需要时使用[明文客户端插件](https://dev.mysql.com/doc/en/cleartext-pluggable-authentication.html)，例如通过[PAM身份验证插件](http://dev.mysql.com/doc/en/pam-authentication-plugin.html)定义的账户。以明文形式发送密码可能会在某些配置中引发安全问题。为避免可能被截获密码的问题，客户端应使用保护密码的方法连接到MySQL服务器。可能的保护方法包括[TLS / SSL](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-tls)、IPsec或专用网络。

###### allowFallbackToPlaintext

```plaintext
Type:           bool
Valid Values:   true, false
Default:        false
```

`allowFallbackToPlaintext=true` acts like a `--ssl-mode=PREFERRED` MySQL client as described in [Command Options for Connecting to the Server](https://dev.mysql.com/doc/refman/5.7/en/connection-options.html#option_general_ssl-mode)

​	`allowFallbackToPlaintext=true`的行为类似于MySQL客户端的`--ssl-mode=PREFERRED`，如[连接服务器的命令选项](https://dev.mysql.com/doc/refman/5.7/en/connection-options.html#option_general_ssl-mode)所述。

###### allowNativePasswords

```plaintext
Type:           bool
Valid Values:   true, false
Default:        true
```

`allowNativePasswords=false` disallows the usage of MySQL native password method.

​	`allowNativePasswords=false`禁止使用MySQL本地密码方法。

###### allowOldPasswords

```plaintext
Type:           bool
Valid Values:   true, false
Default:        false
```

`allowOldPasswords=true` allows the usage of the insecure old password method. This should be avoided, but is necessary in some cases. See also [the old_passwords wiki page](https://github.com/go-sql-driver/mysql/wiki/old_passwords).

​	`allowOldPasswords=true`允许使用不安全的旧密码方法。这应该避免，但在某些情况下是必要的。另请参阅[旧密码的维基页面](https://github.com/go-sql-driver/mysql/wiki/old_passwords)。

###### charset

```plaintext
Type:           string
Valid Values:   <name>
Default:        none
```

Sets the charset used for client-server interaction (`"SET NAMES <value>"`). If multiple charsets are set (separated by a comma), the following charset is used if setting the charset failes. This enables for example support for `utf8mb4` ([introduced in MySQL 5.5.3](http://dev.mysql.com/doc/refman/5.5/en/charset-unicode-utf8mb4.html)) with fallback to `utf8` for older servers (`charset=utf8mb4,utf8`).

​	设置用于客户端-服务器交互的字符集 (`"SET NAMES <value>"`)。如果设置了多个字符集（用逗号分隔），则如果设置字符集失败，将使用下一个字符集。例如，这可以支持 `utf8mb4` （[在MySQL 5.5.3中引入](http://dev.mysql.com/doc/refman/5.5/en/charset-unicode-utf8mb4.html)）并且对于旧版本的服务器降级为 `utf8` （`charset=utf8mb4,utf8`）。

Usage of the `charset` parameter is discouraged because it issues additional queries to the server. Unless you need the fallback behavior, please use `collation` instead.

​	不建议使用 `charset` 参数，因为它会向服务器发送额外的查询。除非你需要降级行为，否则请使用 `collation` 参数代替。

###### checkConnLiveness

```plaintext
Type:           bool
Valid Values:   true, false
Default:        true
```

On supported platforms connections retrieved from the connection pool are checked for liveness before using them. If the check fails, the respective connection is marked as bad and the query retried with another connection. `checkConnLiveness=false` disables this liveness check of connections.

​	在支持的平台中，从连接池中获取的连接在使用之前会检查其是否存活。如果检查失败，相应的连接将被标记为不良，并且查询将使用另一个连接重试。`checkConnLiveness=false` 禁用连接的存活检查。

###### collation

```plaintext
Type:           string
Valid Values:   <name>
Default:        utf8mb4_general_ci
```

Sets the collation used for client-server interaction on connection. In contrast to `charset`, `collation` does not issue additional queries. If the specified collation is unavailable on the target server, the connection will fail.

​	设置用于客户端-服务器交互的连接字符集的排序规则。与 `charset` 不同，`collation` 不会发送额外的查询。如果目标服务器上指定的排序规则不可用，连接将失败。

A list of valid charsets for a server is retrievable with `SHOW COLLATION`.

​	可以通过 `SHOW COLLATION` 检索有效的字符集列表。

The default collation (`utf8mb4_general_ci`) is supported from MySQL 5.5. You should use an older collation (e.g. `utf8_general_ci`) for older MySQL.

​	默认排序规则 (`utf8mb4_general_ci`) 自MySQL 5.5开始支持。对于较旧的MySQL版本，您应该使用较旧的排序规则 (例如 `utf8_general_ci`)。

Collations for charset "ucs2", "utf16", "utf16le", and "utf32" can not be used ([ref](https://dev.mysql.com/doc/refman/5.7/en/charset-connection.html#charset-connection-impermissible-client-charset)).

​	无法使用 "ucs2", "utf16", "utf16le", 和 "utf32" 字符集的排序规则 ([参考](https://dev.mysql.com/doc/refman/5.7/en/charset-connection.html#charset-connection-impermissible-client-charset))。

###### clientFoundRows

```plaintext
Type:           bool
Valid Values:   true, false
Default:        false
```

`clientFoundRows=true` causes an UPDATE to return the number of matching rows instead of the number of rows changed.

​	将 `clientFoundRows` 设置为 `true` 会使 UPDATE 返回匹配行的数量而不是更改的行数。

###### columnsWithAlias

```plaintext
Type:           bool
Valid Values:   true, false
Default:        false
```

When `columnsWithAlias` is true, calls to `sql.Rows.Columns()` will return the table alias and the column name separated by a dot. For example:

​	当`columnsWithAlias`为`true`时，对`sql.Rows.Columns()`的调用将返回由点分隔的表别名和列名。例如：

```
SELECT u.id FROM users as u
```

will return `u.id` instead of just `id` if `columnsWithAlias=true`.

如果`columnsWithAlias=true`，将返回`u.id`而不是仅仅`id`。

###### interpolateParams

```plaintext
Type:           bool
Valid Values:   true, false
Default:        false
```

If `interpolateParams` is true, placeholders (`?`) in calls to `db.Query()` and `db.Exec()` are interpolated into a single query string with given parameters. This reduces the number of roundtrips, since the driver has to prepare a statement, execute it with given parameters and close the statement again with `interpolateParams=false`.

​	如果`interpolateParams`为`true`，则在调用`db.Query()`和`db.Exec()`中的占位符(`?`)将被替换为带有给定参数的单个查询字符串。这减少了往返次数，因为驱动程序需要准备一个语句，使用给定参数执行它，如果`interpolateParams=false`，则再次关闭语句。

*This can not be used together with the multibyte encodings BIG5, CP932, GB2312, GBK or SJIS. These are rejected as they may [introduce a SQL injection vulnerability](http://stackoverflow.com/a/12118602/3430118)!*

​	*此选项不能与多字节编码BIG5、CP932、GB2312、GBK或SJIS一起使用。这些编码可能会引入SQL注入漏洞，因此被拒绝使用！*

###### loc

```plaintext
Type:           string
Valid Values:   <escaped name>
Default:        UTC
```

Sets the location for time.Time values (when using `parseTime=true`). *"Local"* sets the system's location. See [time.LoadLocation](https://golang.org/pkg/time/#LoadLocation) for details.

​	当使用`parseTime=true`时，设置时间.Time值的时区。"Local"设置了系统的时区。有关详细信息，请参阅[time.LoadLocation](https://golang.org/pkg/time/#LoadLocation)。

Note that this sets the location for time.Time values but does not change MySQL's [time_zone setting](https://dev.mysql.com/doc/refman/5.5/en/time-zone-support.html). For that see the [time_zone system variable](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-system-variables), which can also be set as a DSN parameter.

​	请注意，这仅设置时间.Time值的时区，并不会更改MySQL的[time_zone设置](https://dev.mysql.com/doc/refman/5.5/en/time-zone-support.html)。有关该设置，请参见[time_zone系统变量](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-system-variables)，它也可以作为DSN参数进行设置。

Please keep in mind, that param values must be [url.QueryEscape](https://golang.org/pkg/net/url/#QueryEscape)'ed. Alternatively you can manually replace the `/` with `%2F`. For example `US/Pacific` would be `loc=US%2FPacific`.

​	请注意，参数值必须进行[url.QueryEscape](https://golang.org/pkg/net/url/#QueryEscape)转义。或者，您也可以手动将`/`替换为`%2F`。例如，`US/Pacific`将是`loc=US%2FPacific`。

###### maxAllowedPacket

```plaintext
Type:          decimal number
Default:       64*1024*1024
```

Max packet size allowed in bytes. The default value is 64 MiB and should be adjusted to match the server settings. `maxAllowedPacket=0` can be used to automatically fetch the `max_allowed_packet` variable from server *on every connection*.

​	允许的最大数据包大小（以字节为单位）。默认值为 64 MiB，应调整以匹配服务器设置。`maxAllowedPacket=0` 可以用于从服务器自动获取 `max_allowed_packet` 变量 *在每个连接上*。

###### multiStatements

```plaintext
Type:           bool
Valid Values:   true, false
Default:        false
```

Allow multiple statements in one query. While this allows batch queries, it also greatly increases the risk of SQL injections. Only the result of the first query is returned, all other results are silently discarded.

​	在一个查询中允许多个语句。虽然这允许批量查询，但也会大大增加 SQL 注入的风险。只有第一个查询的结果被返回，所有其他结果都被静默丢弃。

When `multiStatements` is used, `?` parameters must only be used in the first statement.

​	当使用 `multiStatements` 时，`?` 参数只能在第一个语句中使用。

###### parseTime

```plaintext
Type:           bool
Valid Values:   true, false
Default:        false
```

`parseTime=true` changes the output type of `DATE` and `DATETIME` values to `time.Time` instead of `[]byte` / `string` The date or datetime like `0000-00-00 00:00:00` is converted into zero value of `time.Time`.

​	`parseTime=true` 将 `DATE` 和 `DATETIME` 类型的输出更改为 `time.Time`，而不是 `[]byte` / `string`。日期或日期时间（如 `0000-00-00 00:00:00`）将被转换为 `time.Time` 的零值。

###### readTimeout

```plaintext
Type:           duration
Default:        0
```

I/O read timeout. The value must be a decimal number with a unit suffix (*"ms"*, *"s"*, *"m"*, *"h"*), such as *"30s"*, *"0.5m"* or *"1m30s"*.

​	I/O 读取超时。值必须是一个带单位后缀的十进制数字 ("ms", "s", "m", "h")，例如 "30s", "0.5m"或 "1m30s"。

###### rejectReadOnly

```plaintext
Type:           bool
Valid Values:   true, false
Default:        false
```

`rejectReadOnly=true` causes the driver to reject read-only connections. This is for a possible race condition during an automatic failover, where the mysql client gets connected to a read-only replica after the failover.

​	`rejectReadOnly=true` 将导致驱动程序拒绝只读连接。这是为了在自动故障转移期间可能发生 race condition，其中 mysql 客户端在故障转移后连接到只读副本。

Note that this should be a fairly rare case, as an automatic failover normally happens when the primary is down, and the race condition shouldn't happen unless it comes back up online as soon as the failover is kicked off. On the other hand, when this happens, a MySQL application can get stuck on a read-only connection until restarted. It is however fairly easy to reproduce, for example, using a manual failover on AWS Aurora's MySQL-compatible cluster.

​	注意，这应该是一个相当罕见的情况，因为主节点故障时通常会发生自动故障转移，只有在故障转移被触发后立即在线，竞态条件才不应该发生。另一方面，当这种情况发生时，MySQL应用程序可能会在只读连接上卡住，直到重新启动。然而，例如，在AWS Aurora的MySQL兼容集群上使用手动故障转移很容易重现这种情况。

If you are not relying on read-only transactions to reject writes that aren't supposed to happen, setting this on some MySQL providers (such as AWS Aurora) is safer for failovers.

​	如果您不依赖于只读事务来拒绝不应该发生的写入，那么在某些MySQL提供程序（如AWS Aurora）上设置这一点对于故障转移来说更安全。

Note that ERROR 1290 can be returned for a `read-only` server and this option will cause a retry for that error. However the same error number is used for some other cases. You should ensure your application will never cause an ERROR 1290 except for `read-only` mode when enabling this option.

​	请注意，对于一个处于“只读”模式的服务器，可能会返回错误1290，此选项会导致对该错误的重试。然而，相同的错误号码用于其他一些情况。您应该确保在启用此选项时，您的应用程序除了“只读”模式外不会导致错误1290。

###### serverPubKey

```plaintext
Type:           string
Valid Values:   <name>
Default:        none
```

Server public keys can be registered with [`mysql.RegisterServerPubKey`](https://godoc.org/github.com/go-sql-driver/mysql#RegisterServerPubKey), which can then be used by the assigned name in the DSN. Public keys are used to transmit encrypted data, e.g. for authentication. If the server's public key is known, it should be set manually to avoid expensive and potentially insecure transmissions of the public key from the server to the client each time it is required.

​	服务器公钥可以通过`mysql.RegisterServerPubKey`进行注册，然后可以在DSN中使用分配的名称。公钥用于传输加密数据，例如用于身份验证。如果服务器的公钥是已知的，应该手动设置以避免每次需要时从服务器到客户端传输公钥的昂贵且可能不安全。

###### timeout

```plaintext
Type:           duration
Default:        OS default
```

Timeout for establishing connections, aka dial timeout. The value must be a decimal number with a unit suffix (*"ms"*, *"s"*, *"m"*, *"h"*), such as *"30s"*, *"0.5m"* or *"1m30s"*.

​	建立连接的超时时间，即拨号超时。值必须是一个带有单位后缀的十进制数字（"ms" *、* "s" *、* "m" *、* "h" ），例如  "30s" *、* "0.5m"  或  "1m30s"。

###### tls

```plaintext
Type:           bool / string
Valid Values:   true, false, skip-verify, preferred, <name>
Default:        false
```

`tls=true` enables TLS / SSL encrypted connection to the server. Use `skip-verify` if you want to use a self-signed or invalid certificate (server side) or use `preferred` to use TLS only when advertised by the server. This is similar to `skip-verify`, but additionally allows a fallback to a connection which is not encrypted. Neither `skip-verify` nor `preferred` add any reliable security. You can use a custom TLS config after registering it with [`mysql.RegisterTLSConfig`](https://godoc.org/github.com/go-sql-driver/mysql#RegisterTLSConfig).

​	`tls=true`启用与服务器之间的TLS / SSL加密连接。如果你想使用自签名或无效证书（服务器端），请使用`skip-verify`。或者，你可以使用`preferred`仅在服务器广告时使用TLS。这类似于`skip-verify`，但额外允许回退到未加密的连接。无论是`skip-verify`还是`preferred`都没有增加可靠的安全性。你可以在通过`mysql.RegisterTLSConfig`注册后使用自定义TLS配置。

###### writeTimeout

```plaintext
Type:           duration
Default:        0
```

I/O write timeout. The value must be a decimal number with a unit suffix (*"ms"*, *"s"*, *"m"*, *"h"*), such as *"30s"*, *"0.5m"* or *"1m30s"*.

​	I/O 写入超时。值必须是一个带有单位后缀的十进制数字（ "ms" 、 "s" 、 "m" 、 "h" ），例如  "30s" 、 "0.5m"  或  "1m30s"。

###### 系统变量 System Variables

Any other parameters are interpreted as system variables:

​	任何其他参数都被解释为系统变量：

- `<boolean_var>=<value>`: `SET <boolean_var>=<value>`
- `<enum_var>=<value>`: `SET <enum_var>=<value>`
- `<string_var>=%27<value>%27`: `SET <string_var>='<value>'`

###### 规则 Rules

- The values for string variables must be quoted with `'`.
- 字符串变量的值必须用 `'` 引用。
- The values must also be [url.QueryEscape](http://golang.org/pkg/net/url/#QueryEscape)'ed! (which implies values of string variables must be wrapped with `%27`).
- 值还必须进行 [url.QueryEscape](http://golang.org/pkg/net/url/#QueryEscape)（这暗示着字符串变量的值必须用 `%27` 包裹）。

###### 例子 Examples

- `autocommit=1`: `SET autocommit=1`
- [`time_zone=%27Europe%2FParis%27`](https://dev.mysql.com/doc/refman/5.5/en/time-zone-support.html): `SET time_zone='Europe/Paris'`
- [`transaction_isolation=%27REPEATABLE-READ%27`](https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_transaction_isolation): `SET transaction_isolation='REPEATABLE-READ'`

#####  Examples

```plaintext
user@unix(/path/to/socket)/dbname
```

```plaintext
root:pw@unix(/tmp/mysql.sock)/myDatabase?loc=Local
```

```plaintext
user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
```

Treat warnings as errors by setting the system variable [`sql_mode`](https://dev.mysql.com/doc/refman/5.7/en/sql-mode.html):

​	将警告视为错误，设置系统变量`sql_mode`为`TRADITIONAL`：

```plaintext
user:password@/dbname?sql_mode=TRADITIONAL
```

TCP via IPv6:

​	通过IPv6使用TCP：

```plaintext
user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?timeout=90s&collation=utf8mb4_unicode_ci
```

TCP on a remote host, e.g. Amazon RDS:

​	在远程主机上使用TCP，例如Amazon RDS：

```plaintext
id:password@tcp(your-amazonaws-uri.com:3306)/dbname
```

Google Cloud SQL on App Engine:

​	在App Engine上使用Google Cloud SQL：

```plaintext
user:password@unix(/cloudsql/project-id:region-name:instance-name)/dbname
```

TCP using default port (3306) on localhost:

​	在本地主机上使用TCP默认端口（3306）：

```plaintext
user:password@tcp/dbname?charset=utf8mb4,utf8&sys_var=esc%40ped
```

Use the default protocol (tcp) and host (localhost:3306):

​	使用默认协议（tcp）和主机（localhost:3306）：

```plaintext
user:password@/dbname
```

No Database preselected:

​	未预选数据库：

```plaintext
user:password@/
```

#### 连接池和超时设置 Connection pool and timeouts

The connection pool is managed by Go's database/sql package. For details on how to configure the size of the pool and how long connections stay in the pool see `*DB.SetMaxOpenConns`, `*DB.SetMaxIdleConns`, and `*DB.SetConnMaxLifetime` in the [database/sql documentation](https://golang.org/pkg/database/sql/). The read, write, and dial timeouts for each individual connection are configured with the DSN parameters [`readTimeout`](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-readtimeout), [`writeTimeout`](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-writetimeout), and [`timeout`](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-timeout), respectively.

​	连接池由Go的database/sql包管理。要配置连接池的大小以及连接在池中保持活动的时间，请参阅[database/sql文档中的`*DB.SetMaxOpenConns`](https://golang.org/pkg/database/sql/)，`*DB.SetMaxIdleConns`和`*DB.SetConnMaxLifetime`。每个连接的读、写和拨号超时时间可以通过DSN参数`readTimeout`，`writeTimeout`和`timeout`分别配置。

#### `ColumnType` Support

This driver supports the [`ColumnType` interface](https://golang.org/pkg/database/sql/#ColumnType) introduced in Go 1.8, with the exception of [`ColumnType.Length()`](https://golang.org/pkg/database/sql/#ColumnType.Length), which is currently not supported. All Unsigned database type names will be returned `UNSIGNED `with `INT`, `TINYINT`, `SMALLINT`, `BIGINT`.

​	该驱动程序支持Go 1.8中引入的`ColumnType`接口，但不支持`ColumnType.Length()`方法，该方法目前不受支持。所有无符号的数据库类型名称将以`UNSIGNED `的形式返回，包括`INT`，`TINYINT`，`SMALLINT`，`BIGINT`。

#### `context.Context` Support

Go 1.8 added `database/sql` support for `context.Context`. This driver supports query timeouts and cancellation via contexts. See [context support in the database/sql package](https://golang.org/doc/go1.8#database_sql) for more details.

​	Go 1.8为`database/sql`包添加了对`context.Context`的支持。该驱动程序支持通过上下文查询超时和取消。有关更多详细信息，请参阅[database/sql包中的上下文支持](https://golang.org/doc/go1.8#database_sql)。

#### `LOAD DATA LOCAL INFILE` support

For this feature you need direct access to the package. Therefore you must change the import path (no `_`):

​	要使用此功能，您需要直接访问包。因此，您必须更改导入路径（无下划线）：

```
import "github.com/go-sql-driver/mysql"
```

Files must be explicitly allowed by registering them with `mysql.RegisterLocalFile(filepath)` (recommended) or the allowlist check must be deactivated by using the DSN parameter `allowAllFiles=true` ([*Might be insecure!*](http://dev.mysql.com/doc/refman/5.7/en/load-data-local.html)).

​	文件必须通过使用`mysql.RegisterLocalFile(filepath)`（推荐）显式允许，或者通过使用DSN参数`allowAllFiles=true`禁用允许列表检查（[*可能不安全！*](http://dev.mysql.com/doc/refman/5.7/en/load-data-local.html)）。

To use a `io.Reader` a handler function must be registered with `mysql.RegisterReaderHandler(name, handler)` which returns a `io.Reader` or `io.ReadCloser`. The Reader is available with the filepath `Reader::<name>` then. Choose different names for different handlers and `DeregisterReaderHandler` when you don't need it anymore.

​	要使用`io.Reader`，必须使用`mysql.RegisterReaderHandler(name, handler)`注册处理程序函数，该函数返回`io.Reader`或`io.ReadCloser`。然后，可以通过文件路径`Reader::<name>`访问该读取器。为不同的处理程序选择不同的名称，并在不需要时使用`DeregisterReaderHandler`取消注册。

See the [godoc of Go-MySQL-Driver](https://godoc.org/github.com/go-sql-driver/mysql) for details.

​	有关详细信息，请参阅[Go-MySQL-Driver的godoc](https://godoc.org/github.com/go-sql-driver/mysql)。

#### `time.Time` support

The default internal output type of MySQL `DATE` and `DATETIME` values is `[]byte` which allows you to scan the value into a `[]byte`, `string` or `sql.RawBytes` variable in your program.

​	MySQL的`DATE`和`DATETIME`值的默认内部输出类型是`[]byte`，这使得你可以将值扫描到程序中的`[]byte`、`string`或`sql.RawBytes`变量中。

However, many want to scan MySQL `DATE` and `DATETIME` values into `time.Time` variables, which is the logical equivalent in Go to `DATE` and `DATETIME` in MySQL. You can do that by changing the internal output type from `[]byte` to `time.Time` with the DSN parameter `parseTime=true`. You can set the default [`time.Time` location](https://golang.org/pkg/time/#Location) with the `loc` DSN parameter.

​	但是，许多人都希望将MySQL的`DATE`和`DATETIME`值扫描到`time.Time`变量中，因为这在Go中是MySQL的`DATE`和`DATETIME`的逻辑等价物。你可以通过使用DSN参数`parseTime=true`将内部输出类型从`[]byte`更改为`time.Time`来实现这一点。你可以使用`loc` DSN参数设置默认的[ `time.Time` 位置](https://golang.org/pkg/time/#Location)。

**Caution:** As of Go 1.1, this makes `time.Time` the only variable type you can scan `DATE` and `DATETIME` values into. This breaks for example [`sql.RawBytes` support](https://github.com/go-sql-driver/mysql/wiki/Examples#rawbytes).

**注意：** 从Go 1.1开始，这使得`time.Time`成为可以扫描`DATE`和`DATETIME`值的唯一变量类型。这会打破例如[ `sql.RawBytes` 支持](https://github.com/go-sql-driver/mysql/wiki/Examples#rawbytes)。

#### Unicode support

Since version 1.5 Go-MySQL-Driver automatically uses the collation `utf8mb4_general_ci` by default.

​	自版本1.5起，Go-MySQL-Driver自动使用默认的collation `utf8mb4_general_ci`。

Other collations / charsets can be set using the [`collation`](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-collation) DSN parameter.

​	可以使用[ `collation`](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-collation) DSN参数设置其他collation / charset。

Version 1.0 of the driver recommended adding `&charset=utf8` (alias for `SET NAMES utf8`) to the DSN to enable proper UTF-8 support. This is not necessary anymore. The [`collation`](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-collation) parameter should be preferred to set another collation / charset than the default.

​	驱动程序版本1.0曾建议在DSN中添加`&charset=utf8`（别名`SET NAMES utf8`）以启用正确的UTF-8支持。这不再是必需的。[ `collation`](https://pkg.go.dev/github.com/go-sql-driver/mysql#readme-collation)参数应该优先于设置除默认之外的其他collation / charset。

See http://dev.mysql.com/doc/refman/8.0/en/charset-unicode.html for more details on MySQL's Unicode support.

​	有关MySQL的Unicode支持的更多详细信息，请参见：http://dev.mysql.com/doc/refman/8.0/en/charset-unicode.html。

### Testing / Development

To run the driver tests you may need to adjust the configuration. See the [Testing Wiki-Page](https://github.com/go-sql-driver/mysql/wiki/Testing) for details.

​	要运行驱动程序的测试，你可能需要调整配置。有关详细信息，请参阅[测试Wiki页面](https://github.com/go-sql-driver/mysql/wiki/Testing)。

Go-MySQL-Driver is not feature-complete yet. Your help is very appreciated. If you want to contribute, you can work on an [open issue](https://github.com/go-sql-driver/mysql/issues?state=open) or review a [pull request](https://github.com/go-sql-driver/mysql/pulls).

​	Go-MySQL-Driver目前尚未完成所有功能。非常感谢您的帮助。如果您想做出贡献，可以处理一个[未解决的问题](https://github.com/go-sql-driver/mysql/issues?state=open)或审阅一个[pull请求](https://github.com/go-sql-driver/mysql/pulls)。

See the [Contribution Guidelines](https://github.com/go-sql-driver/mysql/blob/master/.github/CONTRIBUTING.md) for details.

​	请参阅[贡献指南](https://github.com/go-sql-driver/mysql/blob/master/.github/CONTRIBUTING.md)了解详细信息。

------

### 许可证 License

Go-MySQL-Driver is licensed under the [Mozilla Public License Version 2.0](https://raw.github.com/go-sql-driver/mysql/master/LICENSE)

​	Go-MySQL-Driver是根据[Mozilla公共许可证版本2.0](https://raw.github.com/go-sql-driver/mysql/master/LICENSE)许可的。

Mozilla summarizes the license scope as follows:

​	Mozilla对许可证范围的总结如下：

> MPL: The copyleft applies to any files containing MPLed code.
>
> MPL：复制权适用于任何包含MPL代码的文件。

That means:

这意味着：

- You can **use** the **unchanged** source code both in private and commercially.
- 您可以使用未更改的源代码进行私人或商业用途。
- When distributing, you **must publish** the source code of any **changed files** licensed under the MPL 2.0 under a) the MPL 2.0 itself or b) a compatible license (e.g. GPL 3.0 or Apache License 2.0).
- 在发布时，您必须发布任何根据MPL 2.0许可的更改文件的源代码，要么根据MPL 2.0本身发布，要么根据兼容许可证（例如GPL 3.0或Apache许可证2.0）发布。
- You **needn't publish** the source code of your library as long as the files licensed under the MPL 2.0 are **unchanged**.
- 只要根据MPL 2.0许可的文件未更改，您就不需要发布您的库的源代码。

Please read the [MPL 2.0 FAQ](https://www.mozilla.org/en-US/MPL/2.0/FAQ/) if you have further questions regarding the license.

​	如果您对许可证有任何进一步的问题，请阅读[MPL 2.0常见问题解答](https://www.mozilla.org/en-US/MPL/2.0/FAQ/)。

You can read the full terms here: [LICENSE](https://raw.github.com/go-sql-driver/mysql/master/LICENSE).

​	您可以在此处阅读全部条款：[LICENSE](https://raw.github.com/go-sql-driver/mysql/master/LICENSE)。

![Go Gopher and MySQL Dolphin](_index_img/go-mysql-driver_m.jpg)

​       

### 概述

Package mysql provides a MySQL driver for Go's database/sql package.

​	mysql包为Go的database/sql包提供了一个MySQL驱动。

The driver should be used via the database/sql package:

​	应该通过database/sql包使用该驱动：

```go
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")
```

See https://github.com/go-sql-driver/mysql#usage for details

​	详细信息请参阅[https://github.com/go-sql-driver/mysql#usage](https://github.com/go-sql-driver/mysql#usage)

### 常量 

This section is empty.

### 变量

[View Source](https://github.com/go-sql-driver/mysql/blob/v1.7.1/errors.go#L19)

``` go
var (
	ErrInvalidConn       = errors.New("invalid connection")
	ErrMalformPkt        = errors.New("malformed packet")
	ErrNoTLS             = errors.New("TLS requested but server does not support TLS")
	ErrCleartextPassword = errors.New("this user requires clear text authentication. If you still want to use it, please add 'allowCleartextPasswords=1' to your DSN")
	ErrNativePassword    = errors.New("this user requires mysql native password authentication.")
	ErrOldPassword       = errors.New("this user requires old password authentication. If you still want to use it, please add 'allowOldPasswords=1' to your DSN. See also https://github.com/go-sql-driver/mysql/wiki/old_passwords")
	ErrUnknownPlugin     = errors.New("this authentication plugin is not supported")
	ErrOldProtocol       = errors.New("MySQL server does not support required protocol 41+")
	ErrPktSync           = errors.New("commands out of sync. You can't run this command now")
	ErrPktSyncMul        = errors.New("commands out of sync. Did you run multiple statements at once?")
	ErrPktTooLarge       = errors.New("packet for query is too large. Try adjusting the `Config.MaxAllowedPacket`")
	ErrBusyBuffer        = errors.New("busy buffer")
)
```

Various errors the driver might return. Can change between driver versions.

​	驱动程序可能返回的各种错误。可以在驱动程序版本之间更改。

### 函数

#### func DeregisterLocalFile 

``` go
func DeregisterLocalFile(filePath string)
```

DeregisterLocalFile removes the given filepath from the allowlist.

​	DeregisterLocalFile 函数从允许列表中删除给定的文件路径。

#### func DeregisterReaderHandler 

``` go
func DeregisterReaderHandler(name string)
```

DeregisterReaderHandler removes the ReaderHandler function with the given name from the registry.

​	DeregisterReaderHandler 函数从注册表中删除具有给定名称的 ReaderHandler 函数。

#### func DeregisterServerPubKey <- 1.4.0

``` go
func DeregisterServerPubKey(name string)
```

DeregisterServerPubKey removes the public key registered with the given name.

​	DeregisterServerPubKey 函数删除具有给定名称的已注册公钥。

#### func DeregisterTLSConfig <- 1.1.0

``` go
func DeregisterTLSConfig(key string)
```

DeregisterTLSConfig removes the tls.Config associated with key.

​	DeregisterTLSConfig 函数删除与键关联的 tls.Config。

#### func NewConnector <- 1.5.0

``` go
func NewConnector(cfg *Config) (driver.Connector, error)
```

NewConnector returns new driver.Connector.

​	NewConnector 函数返回新的 driver.Connector。

#### func RegisterDial <- DEPRECATED

```
func RegisterDial(network string, dial DialFunc)
```

RegisterDial registers a custom dial function. It can then be used by the network address mynet(addr), where mynet is the registered new network. addr is passed as a parameter to the dial function.

​	RegisterDial 函数注册一个自定义的拨号函数。它可以通过将 mynet(addr) 作为网络地址使用，其中 mynet 是已注册的新网络，addr 是传递给拨号函数的参数。

Deprecated: users should call RegisterDialContext instead

​	已弃用：用户应改用 RegisterDialContext。

#### func RegisterDialContext <- 1.5.0

``` go
func RegisterDialContext(net string, dial DialContextFunc)
```

RegisterDialContext registers a custom dial function. It can then be used by the network address mynet(addr), where mynet is the registered new network. The current context for the connection and its address is passed to the dial function.

​	RegisterDialContext 函数注册一个自定义的拨号函数。它可以通过将 mynet(addr) 作为网络地址使用，其中 mynet 是已注册的新网络。连接的当前上下文和其地址将传递给拨号函数。

#### func RegisterLocalFile 

``` go
func RegisterLocalFile(filePath string)
```

RegisterLocalFile adds the given file to the file allowlist, so that it can be used by "LOAD DATA LOCAL INFILE <filepath>". Alternatively you can allow the use of all local files with the DSN parameter 'allowAllFiles=true'

​	RegisterLocalFile 函数将给定文件添加到允许的文件列表中，以便可以通过 "LOAD DATA LOCAL INFILE <filepath>" 使用它。或者，您可以使用 DSN 参数 'allowAllFiles=true' 来允许使用所有本地文件。

```
filePath := "/home/gopher/data.csv"
mysql.RegisterLocalFile(filePath)
err := db.Exec("LOAD DATA LOCAL INFILE '" + filePath + "' INTO TABLE foo")
if err != nil {
...
```

#### func RegisterReaderHandler 

``` go
func RegisterReaderHandler(name string, handler func() io.Reader)
```

RegisterReaderHandler registers a handler function which is used to receive a io.Reader. The Reader can be used by "LOAD DATA LOCAL INFILE Reader::<name>". If the handler returns a io.ReadCloser Close() is called when the request is finished.

​	RegisterReaderHandler 函数注册一个处理函数，该函数返回一个 io.Reader，该 Reader 可以被 "LOAD DATA LOCAL INFILE Reader::<name>" 使用。如果处理函数返回一个 io.ReadCloser，则在请求完成时调用 Close()。

```
mysql.RegisterReaderHandler("data", func() io.Reader {
	var csvReader io.Reader // Some Reader that returns CSV data
	... // Open Reader here
	return csvReader
})
err := db.Exec("LOAD DATA LOCAL INFILE 'Reader::data' INTO TABLE foo")
if err != nil {
...
```

#### func RegisterServerPubKey <- 1.4.0

``` go
func RegisterServerPubKey(name string, pubKey *rsa.PublicKey)
```

RegisterServerPubKey registers a server RSA public key which can be used to send data in a secure manner to the server without receiving the public key in a potentially insecure way from the server first. Registered keys can afterwards be used adding serverPubKey=<name> to the DSN.

​	RegisterServerPubKey 函数注册一个服务器 RSA 公钥，用于以安全方式向服务器发送数据，而无需首先从服务器接收可能不安全的公钥。注册的密钥可以在 DSN 中添加 serverPubKey=<name> 来使用。

Note: The provided rsa.PublicKey instance is exclusively owned by the driver after registering it and may not be modified.

注意：提供的 rsa.PublicKey 实例在注册后由驱动程序独家拥有，并且不允许进行修改。

```
data, err := ioutil.ReadFile("mykey.pem")
if err != nil {
	log.Fatal(err)
}

block, _ := pem.Decode(data)
if block == nil || block.Type != "PUBLIC KEY" {
	log.Fatal("failed to decode PEM block containing public key")
}

pub, err := x509.ParsePKIXPublicKey(block.Bytes)
if err != nil {
	log.Fatal(err)
}

if rsaPubKey, ok := pub.(*rsa.PublicKey); ok {
	mysql.RegisterServerPubKey("mykey", rsaPubKey)
} else {
	log.Fatal("not a RSA public key")
}
```

#### func RegisterTLSConfig <- 1.1.0

``` go
func RegisterTLSConfig(key string, config *tls.Config) error
```

RegisterTLSConfig registers a custom tls.Config to be used with sql.Open. Use the key as a value in the DSN where tls=value.

​	RegisterTLSConfig 函数用于注册自定义的 tls.Config 以供 sql.Open 使用。在 DSN 中将 key 作为 tls=value 的值使用。

Note: The provided tls.Config is exclusively owned by the driver after registering it.

​	注意：提供的 tls.Config 在注册后由驱动程序独家拥有。

```go
rootCertPool := x509.NewCertPool()
pem, err := ioutil.ReadFile("/path/ca-cert.pem")
if err != nil {
    log.Fatal(err)
}
if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
    log.Fatal("Failed to append PEM.")
}
clientCert := make([]tls.Certificate, 0, 1)
certs, err := tls.LoadX509KeyPair("/path/client-cert.pem", "/path/client-key.pem")
if err != nil {
    log.Fatal(err)
}
clientCert = append(clientCert, certs)
mysql.RegisterTLSConfig("custom", &tls.Config{
    RootCAs: rootCertPool,
    Certificates: clientCert,
})
db, err := sql.Open("mysql", "user@tcp(localhost:3306)/test?tls=custom")
```

#### func SetLogger <- 1.2.0

``` go
func SetLogger(logger Logger) error
```

SetLogger is used to set the logger for critical errors. The initial logger is os.Stderr.

​	SetLogger 函数用于设置用于记录关键错误的记录器。初始记录器为 os.Stderr。

### 类型

#### type Config <- 1.3.0

``` go
type Config struct {
	User             string            // 用户名 Username
	Passwd           string            // 密码（需要用户名） Password (requires User)
	Net              string            // 网络类型 Network type
	Addr             string            // 网络地址（需要网络类型） Network address (requires Net)
	DBName           string            // 数据库名 Database name
	Params           map[string]string // 连接参数 Connection parameters
	Collation        string            // 连接排序规则 Connection collation
	Loc              *time.Location    // 时间值的地点 Location for time.Time values
	MaxAllowedPacket int               // 允许的最大数据包大小 Max packet size allowed
	ServerPubKey     string            // 服务器公钥名称 Server public key name

	TLSConfig    string        // TLS配置名称 TLS configuration name
	TLS          *tls.Config   // TLS配置，它的优先级高于TLSConfig TLS configuration, its priority is higher than TLSConfig
	Timeout      time.Duration // 拨号超时时间 Dial timeout
	ReadTimeout  time.Duration // I/O读取超时时间 I/O read timeout
	WriteTimeout time.Duration // I/O写入超时时间 I/O write timeout

	AllowAllFiles            bool // 允许使用所有文件进行LOAD DATA LOCAL INFILE操作 Allow all files to be used with LOAD DATA LOCAL INFILE
	AllowCleartextPasswords  bool // 允许明文客户端插件 Allows the cleartext client side plugin
	AllowFallbackToPlaintext bool // 如果服务器不支持TLS，允许降级到非加密连接 Allows fallback to unencrypted connection if server does not support TLS
	AllowNativePasswords     bool // 允许使用本地密码认证方法 Allows the native password authentication method
	AllowOldPasswords        bool // 允许使用旧的不安全的密码方法 Allows the old insecure password method
	CheckConnLiveness        bool // 在使用连接之前检查连接的存活状态 Check connections for liveness before using them
	ClientFoundRows          bool // 返回匹配的行数而不是更改的行数 Return number of matching rows instead of rows changed
	ColumnsWithAlias         bool // 将表别名添加到列名前 Prepend table alias to column names
	InterpolateParams        bool // 将占位符插值到查询字符串中 Interpolate placeholders into query string
	MultiStatements          bool // 允许在单个查询中使用多个语句 Allow multiple statements in one query
	ParseTime                bool // 将时间值解析为time.Time类型 Parse time values to time.Time
	RejectReadOnly           bool // 拒绝只读连接 Reject read-only connections
	// contains filtered or unexported fields
}
```

Config is a configuration parsed from a DSN string. If a new Config is created instead of being parsed from a DSN string, the NewConfig function should be used, which sets default values.

​	Config是从DSN字符串解析出的配置。如果要从DSN字符串创建一个新的Config，而不是进行解析，应使用NewConfig函数，该函数将设置默认值。

##### func NewConfig <- 1.4.0

``` go
func NewConfig() *Config
```

NewConfig creates a new Config and sets default values.

​	NewConfig创建一个新的Config并设置默认值。

##### func ParseDSN <- 1.3.0

``` go
func ParseDSN(dsn string) (cfg *Config, err error)
```

ParseDSN parses the DSN string to a Config

​	ParseDSN解析DSN字符串到Config。

##### (*Config) Clone <- 1.5.0

``` go
func (cfg *Config) Clone() *Config
```

##### (*Config) FormatDSN <- 1.3.0

``` go
func (cfg *Config) FormatDSN() string
```

FormatDSN formats the given Config into a DSN string which can be passed to the driver.

​	FormatDSN将给定的Config格式化为可以传递给驱动程序的DSN字符串。

#### type DialContextFunc <- 1.5.0

``` go
type DialContextFunc func(ctx context.Context, addr string) (net.Conn, error)
```

DialContextFunc is a function which can be used to establish the network connection. Custom dial functions must be registered with RegisterDialContext

​	DialContextFunc是可以用于建立网络连接的函数。自定义拨号函数必须通过RegisterDialContext进行注册。

#### type DialFunc <- Deprecated

```
type DialFunc func(addr string) (net.Conn, error)
```

DialFunc is a function which can be used to establish the network connection. Custom dial functions must be registered with RegisterDial

​	DialFunc是可以用于建立网络连接的函数。自定义拨号函数必须通过RegisterDial进行注册。

Deprecated: users should register a DialContextFunc instead

​	已弃用：用户应注册一个 DialContextFunc 代替



#### type Logger <- 1.2.0

``` go
type Logger interface {
	Print(v ...interface{})
}
```

Logger is used to log critical error messages.

​	Logger接口被用于记录关键错误信息。

#### type MySQLDriver <- 1.1.0

``` go
type MySQLDriver struct{}
```

MySQLDriver is exported to make the driver directly accessible. In general the driver is used via the database/sql package.

​	MySQLDriver 结构体是为了让驱动可以直接访问而导出的。通常情况下，驱动是通过`database/sql`包来使用的。

##### (MySQLDriver) Open <- 1.1.0

``` go
func (d MySQLDriver) Open(dsn string) (driver.Conn, error)
```

Open new Connection. See https://github.com/go-sql-driver/mysql#dsn-data-source-name for how the DSN string is formatted

​	Open函数用于新建一个连接。DSN（数据源名称）字符串的格式参见 [https://github.com/go-sql-driver/mysql#dsn-data-source-name](https://github.com/go-sql-driver/mysql#dsn-data-source-name)

##### (MySQLDriver) OpenConnector <- 1.5.0

``` go
func (d MySQLDriver) OpenConnector(dsn string) (driver.Connector, error)
```

OpenConnector implements driver.DriverContext.

​	OpenConnector方法实现了driver.DriverContext接口。

#### type MySQLError 

``` go
type MySQLError struct {
	Number   uint16
	SQLState [5]byte
	Message  string
}
```

MySQLError is an error type which represents a single MySQL error

​	MySQLError是一种错误类型，它代表一个MySQL错误。

##### (*MySQLError) Error 

``` go
func (me *MySQLError) Error() string
```

##### (*MySQLError) Is <- 1.7.0

``` go
func (me *MySQLError) Is(err error) bool
```

#### type NullTime

```
type NullTime sql.NullTime
```

​	[This NullTime implementation is not driver-specific](https://pkg.go.dev/github.com/go-sql-driver/mysql#hdr-This_NullTime_implementation_is_not_driver_specific)

NullTime represents a time.Time that may be NULL. NullTime implements the Scanner interface so it can be used as a scan destination:

​	NullTime 代表一个可能是 NULL 的 `time.Time` 类型。NullTime 实现了 `Scanner` 接口，因此它可以被用作一个扫描目标：

```
var nt NullTime
err := db.QueryRow("SELECT time FROM foo WHERE id=?", id).Scan(&nt)
...
if nt.Valid {
   // use nt.Time
} else {
   // NULL value
}
```

> This NullTime implementation is not driver-specific
>
> **注意**：这个 `NullTime` 实现不是特定于驱动程序的。

Deprecated: NullTime doesn't honor the loc DSN parameter. NullTime.Scan interprets a time as UTC, not the loc DSN parameter. Use sql.NullTime instead.

​	**已弃用**：`NullTime` 不尊重 `loc` DSN 参数。`NullTime.Scan` 将时间解释为 UTC，而不是 `loc` DSN 参数。请改用 `sql.NullTime`。

##### func (*NullTime) Scan

```
func (nt *NullTime) Scan(value interface{}) (err error)
```

Scan implements the Scanner interface. The value type must be time.Time or string / []byte (formatted time-string), otherwise Scan fails.

​	Scan 实现了 Scanner 接口。value 的类型必须是 `time.Time` 或 string / []byte（格式化的时间字符串），否则 Scan 将失败。

##### func (NullTime) Value

```
func (nt NullTime) Value() (driver.Value, error)
```

Value implements the driver Valuer interface.

​	Value 实现了驱动程序的 Valuer 接口。