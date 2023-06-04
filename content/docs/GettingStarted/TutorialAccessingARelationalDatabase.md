+++
title = "教程：访问关系型数据库"
weight = 9
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false
+++
# Tutorial: Accessing a relational database - 教程：访问关系型数据库

> 原文：[https://go.dev/doc/tutorial/database-access](https://go.dev/doc/tutorial/database-access)

​	本教程介绍了用Go和其标准库中的`database/sql`包访问关系数据库的基本知识。

​	如果您对Go及其工具有基本的了解，那么您将从本教程中获益匪浅。如果这是您第一次接触Go，请参阅[Tutorial: Get started with Go](../TutorialGetStartedWithGo) 的快速介绍。

​	您将使用的`database/sql`包包含用于连接数据库、执行事务、取消正在进行的操作等的类型和函数。关于使用该包的更多细节，请参阅[访问数据库](../../UsingAndUnderstandingGo/AccessingDatabases/AccessingRelationalDatabases)。

​	在本教程中，您将创建一个数据库，然后编写代码来访问数据库。您的示例项目将是一个关于复古爵士乐唱片（vintage jazz records）的数据存储库。

在本教程中，您将通过以下几个部分进行学习：

1. 为您的代码创建一个文件夹。
2. 建立一个数据库。
3. 导入数据库驱动程序。
4. 获得数据库句柄并连接。
5. 查询多行（记录）。
6. 查询单行（记录）。
7. 添加数据。

注意：关于其他教程，请参见 [Tutorials](../Tutorials)。

## 先决条件

- 安装[MySQL](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)关系型数据库管理系统（DBMS）。
- 安装 Go。有关安装说明，请参阅[Installing Go](../InstallingGo)。
- 编辑代码的工具。您拥有的任何文本编辑器都可以工作。
- 命令终端。在 Linux 和 Mac 上使用任何终端，以及在 Windows 上使用 `PowerShell` 或 `cmd`，Go 都能很好地工作。

## 为您的代码创建一个文件夹

首先，为您要写的代码创建一个文件夹。

a. 打开一个命令提示符，切换到您的主目录。

在Linux或Mac上：

```shell
$ cd
```

在Windows上：

```shell
C:\> cd %HOMEPATH%
```

​	在本教程的其余部分，我们将显示`$`作为提示符。我们使用的命令在Windows上也可以使用。

b. 在命令提示符下，为您的代码创建一个名为`data-access`的目录。

```shell
$ mkdir data-access
$ cd data-access
```

c. 创建一个模块，您可以在其中管理您将在本教程中添加的依赖项。

​	运行`go mod init`命令，给它您的新代码的模块路径。

```shell
$ go mod init example/data-access
go: creating new go.mod: module example/data-access
```

​	该命令创建了一个`go.mod`文件，您添加的依赖项将被列在其中以便追踪。更多信息，请务必参阅[管理依赖项](../../UsingAndUnderstandingGo/ManagingDependencies)。

> 注意：在实际开发中，您会根据自己的需要指定一个更具体的模块路径。更多信息，请参见[管理依赖项](../../UsingAndUnderstandingGo/ManagingDependencies#naming-a-module)。

接下来，您将创建一个数据库。

## 建立一个数据库

​	在这一步，您将创建将要使用的数据库。您将使用DBMS本身的CLI来创建数据库和表，以及添加数据。

​	您将创建一个数据库，其中包含关于黑胶唱片的复古爵士乐（vintage jazz recordings on vinyl）的数据。

​	这里的代码使用[MySQL CLI](https://dev.mysql.com/doc/refman/8.0/en/mysql.html)，但大多数DBMS都有自己的CLI，具有类似的功能。

1. 打开一个新的命令提示符。

2. 在命令行中，登录到您的DBMS，如下面MySQL的例子。

   ```shell
   $ mysql -u root -p
   Enter password:
   
   mysql>
   ```

3. 在`mysql`命令提示符下，创建一个数据库。

   ```shell
   mysql> create database recordings;
   ```

4. 使用您刚创建的数据库，以便可以添加数据表。

   ```shell
   mysql> use recordings;
   Database changed
   ```

5. 在文本编辑器的`data-access`文件夹中，创建一个名为`create-tables.sql`的文件，以保存用于添加数据表的SQL脚本。

6. 在该文件中，粘贴以下SQL代码，然后保存该文件。

   ```mysql title="create-tables.sql" linenums="1"
   DROP TABLE IF EXISTS album;
   CREATE TABLE album (
     id         INT AUTO_INCREMENT NOT NULL,
     title      VARCHAR(128) NOT NULL,
     artist     VARCHAR(255) NOT NULL,
     price      DECIMAL(5,2) NOT NULL,
     PRIMARY KEY (`id`)
   );
   
   INSERT INTO album
     (title, artist, price)
   VALUES
     ('Blue Train', 'John Coltrane', 56.99),
     ('Giant Steps', 'John Coltrane', 63.99),
     ('Jeru', 'Gerry Mulligan', 17.99),
     ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
   ```

   在这个SQL代码中，您：

   - 删除(drop)一个叫做`album`的表。如果您想重新开始处理这个表，可以先执行这条命令，这样就可以更容易地重新运行这个脚本。
   - 创建一个有四列的`album`表：`title`、`artist`和`price`。每一行的`id`值都是由DBMS自动创建的。
   - 添加四行记录。
   
7. `mysql`命令提示符下，运行您刚刚创建的脚本。

   您将使用以下形式的 `source` 命令：

   ```shell
   mysql> source /path/to/create-tables.sql
   ```

8. 在您的DBMS命令提示符下，使用`SELECT`语句来验证您已经成功创建了有数据的表。

   ```shell
   mysql> select * from album;
   +----+---------------+----------------+-------+
   | id | title         | artist         | price |
   +----+---------------+----------------+-------+
   |  1 | Blue Train    | John Coltrane  | 56.99 |
   |  2 | Giant Steps   | John Coltrane  | 63.99 |
   |  3 | Jeru          | Gerry Mulligan | 17.99 |
   |  4 | Sarah Vaughan | Sarah Vaughan  | 34.98 |
   +----+---------------+----------------+-------+
   4 rows in set (0.00 sec)
   ```

接下来，您将编写一些 Go 代码进行连接，以便可以查询。

## 查找并导入数据库驱动

​	现在您已经有了一个带有一些数据的数据库，开始编写您的Go代码吧。

​	找到并导入一个数据库驱动程序，它可以将您通过`database/sql`包中的函数提出的请求转化为数据库能够理解的请求。

a. 在您的浏览器中，访问[SQLDrivers](https://github.com/golang/go/wiki/SQLDrivers) wiki 页面，以确定您可以使用的驱动程序。

​	使用该页面上的列表来确定您要使用的驱动程序。在本教程中访问MySQL时，您将使用[Go-MySQL-Driver](https://github.com/go-sql-driver/mysql/)。

b. 注意驱动程序的包名 —— 这里是`github.com/go-sql-driver/mysql`。

c. 使用您的文本编辑器，创建一个文件来编写您的Go代码，并将该文件保存为`main.go`，放在您先前创建的`data-access`目录中。

d. 在`main.go`中，粘贴以下代码以导入驱动程序包。

```go title="main.go" linenums="1"
package main

import "github.com/go-sql-driver/mysql"
```

在这段代码中，您

- 将代码添加到`main`包中，以便可以独立执行。
- 导入MySQL驱动程序`github.com/go-sql-driver/mysql`。

导入驱动程序后，您将开始写代码来访问数据库。

## 获取数据库句柄并连接

现在编写一些Go代码，使用数据库句柄进行数据库访问。

您将使用一个指向 `sql.DB` 结构的指针，它代表对特定数据库的访问。

#### 编写代码

a. 在`main.go`中，在您刚刚添加的`import`代码下面，粘贴以下Go代码来创建一个数据库句柄。

```go linenums="1"
var db *sql.DB

func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }
    // Get a database handle.
    var err error
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}
```

在这段代码中，您：

- 声明一个类型为[*sql.DB]({{< ref "/stdLib/database/sql#type-db" >}})的`db`变量。这是您的数据库句柄。

  ​	使`db`成为一个全局变量可以简化这个示例。在生产环境中，您会避免使用全局变量，比如把变量传递给需要它的函数，或者把它包装在一个结构中。

- 使用MySQL驱动的[Config](https://pkg.go.dev/github.com/go-sql-driver/mysql#Config) —— 以及该类型的[FormatDSN](https://pkg.go.dev/github.com/go-sql-driver/mysql#Config.FormatDSN) —— 来收集连接属性并将其格式化为连接字符串的`DSN`。

  `Config`结构使得代码比连接字符串更容易阅读。

- 调用 [sql.Open]({{< ref "/stdLib/database/sql#func-open" >}}) 来初始化 `db` 变量，传递 `FormatDSN` 的返回值。

- 检查`sql.Open`是否有错误。如果数据库连接细节格式不正确，它可能会失败。

  为了简化代码，您要调用`log.Fatal`来结束执行，并将错误打印到控制台。在生产代码中，您会希望以一种更优雅的方式来处理错误。

- 调用[DB.Ping]({{< ref "/stdLib/database/sql#db-ping----go11" >}})来确认连接到数据库是否有效。在运行时，`sql.Open`可能不会立即连接，这取决于驱动程序。您在这里使用`Ping`来确认`database/sql`包在需要时可以连接。

- 检查`Ping`是否有错误，以防连接失败。

- 如果`Ping`连接成功，则打印一条信息。

b. 在`main.go`文件的顶部，就在包声明的下面，导入您需要的包，以支持您刚刚写的代码。

现在文件的顶部应该是这样的：

```go title="main.go" linenums="1"
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)
```

c. 保存`main.go`。



#### 运行代码

a. 开始跟踪MySQL驱动模块作为一个依赖项。

​	使用`go get`来添加`github.com/go-sql-driver/mysql`模块作为您自己模块的依赖。使用`点参数`表示 "获取当前目录下代码的依赖项"。

```shell
$ go get .
go get: added github.com/go-sql-driver/mysql v1.6.0
```

​	Go下载了这个依赖项，因为您在上一步的`import`声明中加入了它。关于依赖项跟踪的更多信息，请参见[添加依赖项](../../UsingAndUnderstandingGo/ManagingDependencies#添加一个依赖项)。

b. 在命令提示符下，设置 `DBUSER` 和 `DBPASS` 环境变量供 Go 程序使用。

在Linux或Mac上：

```
$ export DBUSER=username
$ export DBPASS=password
```

在Windows上：

```shell
C:\Users\you\data-access> set DBUSER=username
C:\Users\you\data-access> set DBPASS=password
```

c. 在包含`main.go`的目录中的命令行中，通过输入`go run`和一个`点参数`来运行代码，表示 "在当前目录中运行软件包"。

```shell
$ go run .
Connected!
```

您可以连接了! 接下来，您将查询一些数据。

## 查询多行（记录）

​	在本节中，您将使用Go来执行一个旨在返回多行的SQL查询。

​	对于可能返回多行的SQL语句，您可以使用`database/sql`包中的`Query`方法，然后循环浏览它所返回的行。(您将在稍后的[单行查询](#查询单行记录)一节中学习如何查询单行。)

#### 编写代码

a. 在`main.go`中，紧挨着`func main`的上方，粘贴以下`Album`结构的定义。您将用它来保存从查询中返回的行数据。

```go linenums="1"
type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}
```

b. 在`func main`下面，粘贴以下`associatesByArtist`函数来查询数据库。

```go linenums="1"
// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album

    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}
```

在这段代码中，您：

- 声明一个您定义的专辑类型的`albums`切片。这将保存来自返回行的数据。结构字段名和类型与数据库列名和类型相对应。

- 使用[DB.Query]({{< ref "/stdLib/database/sql#db-query" >}}) 执行一个`SELECT`语句来查询具有指定艺术家名字的专辑。

  ​	`Query`的第一个参数是SQL语句。在该参数之后，您可以传递零个或多个任何类型的参数。这些参数为您提供了在SQL语句中指定参数值的地方。通过将SQL语句与参数值分开（而不是用例如`fmt.Sprintf`连接），您可以使`database/sql`包将参数值与SQL文本分开发送，从而消除任何SQL注入的风险。

- 推迟关闭`rows`，这样当函数退出时，它持有的任何资源都会被释放。

- 循环返回的行，使用`Rows.Scan`将每行的列值分配给`Album`结构字段。

  ​	`Scan`需要一个指向Go值的指针列表，列值将被写入其中。在这里，您把指针指向`alb`变量中的（用`&`操作符创建的）字段，。`Scan`通过这些指针来更新结构字段。

- 在循环中，检查将列值扫描到结构字段中的错误。

- 在循环中，将新的`alb`追加到`albums`切片中。

- 循环结束后，使用`rows.Err`检查整个查询是否有错误。注意，如果查询本身失败了，检查这里的错误是发现结果不完整的唯一方法。

c. 更新您的`main`函数以调用 `albumsByArtist`。

​	在`func main`的结尾处，添加以下代码。

```go linenums="1"
albums, err := albumsByArtist("John Coltrane")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Albums found: %v\n", albums)
```

在新的代码中，您现在：

- 调用您添加的 `albumsByArtist` 函数，将其返回值分配给一个新的`albums`变量。
- 打印结果。

#### 运行该代码

从包含`main.go`的目录中的命令行，运行代码。

```shell
$ go run .
Connected!
Albums found: [{1 Blue Train John Coltrane 56.99} {2 Giant Steps John Coltrane 63.99}]
```

接下来，您将查询单行。

## 查询单行（记录）

​	在本节中，您将使用Go来查询数据库中的单行。

​	对于最多返回一行的 SQL 语句，您可以使用`QueryRow`，这比使用`Query`循环更简单。

#### 编写代码

a. 在 `albumsByArtist` 下方，粘贴以下 `albumByID` 函数。

```go linenums="1"
// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
    // An album to hold data from the returned row.
    var alb Album

    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumsById %d: no such album", id)
        }
        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return alb, nil
}
```

在这段代码中，您：

- 使用`DB.QueryRow`来执行一个`SELECT`语句，查询具有指定ID的专辑。

  它返回一个`sql.Row`。为了简化调用代码（您的代码！），`QueryRow`并不返回错误。相反，它安排稍后从`Rows.Scan`返回任何查询错误（如`sql.ErrNoRows`）。

- 使用[Row.Scan]({{< ref "/stdLib/database/sql#row-scan" >}})将列值复制到结构字段。

-  检查来自`Scan`的错误。

  ​	特殊错误`sql.ErrNoRows`表示查询没有返回任何行。通常情况下，这个错误值得用更具体的文本来代替，比如这里的 "no such album"。

b. 更新`main`以调用 `albumByID`。

在`func main`的末尾，添加以下代码。

```go linenums="1"
// Hard-code ID 2 here to test the query.
alb, err := albumByID(2)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Album found: %v\n", alb)
```

在新的代码中，您现在：

- 调用您添加的 `albumByID` 函数。
- 打印返回的album 的ID。

#### 运行该代码

从包含main.go的目录中的命令行，运行代码。

```shell
$ go run .
Connected!
Albums found: [{1 Blue Train John Coltrane 56.99} {2 Giant Steps John Coltrane 63.99}]
Album found: {2 Giant Steps John Coltrane 63.99}
```

接下来，您将在数据库中添加一个 album 。

## 添加数据

​	在本节中，您将使用Go来执行一个SQL `INSERT`语句，向数据库添加新的行。

​	您已经看到了如何在返回数据的SQL语句中使用`Query`和`QueryRow`。要执行不返回数据的SQL语句，您可以使用`Exec`。

#### 编写代码

a. 在 `albumByID` 下面，粘贴以下 `addAlbum` 函数，在数据库中插入一个新`album` ，然后保存 main.go。

```go linenums="1"
// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}
```

在这段代码中，您：

- 使用[DB.Exec]({{< ref "/stdLib/database/sql#db-exec" >}})来执行一个`INSERT`语句。

  像`Query`一样，`Exec`接收一个SQL语句，后面是该SQL语句的参数值。

- 检查`INSERT`的尝试是否有错误。

- 使用`Result.LastInsertId`检索插入的数据库行的ID。

- 检查从检索ID的尝试中是否有错误。

b. 更新main以调用新的`addAlbum`函数。

在`func main`的末尾，添加以下代码。

```go linenums="1"
albID, err := addAlbum(Album{
    Title:  "The Modern Sound of Betty Carter",
    Artist: "Betty Carter",
    Price:  49.99,
})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("ID of added album: %v\n", albID)
```

在新的代码中，您现在：

- 用一个新`album `调用`addAlbum`，将您要添加的`album` 的ID分配给`albID`变量。

#### 运行该代码

从包含`main.go`的目录中的命令行，运行代码。

```shell
$ go run .
Connected!
Albums found: [{1 Blue Train John Coltrane 56.99} {2 Giant Steps John Coltrane 63.99}]
Album found: {2 Giant Steps John Coltrane 63.99}
ID of added album: 5
```

## 总结

祝贺您！您刚刚使用 Go 对关系数据库进行了简单的操作。

建议的下一个主题：

- 请看一下数据访问（Accessing Databases下的）指南，其中包含更多关于这里涉及到的主题的更多信息。
- 如果您是Go的新手，您会发现[Effective Go](../../UsingAndUnderstandingGo/EffectiveGo)和[How to write Go code](../HowToWriteGoCode)中描述了有用的最佳实践。
- [go Tour](https://go.dev/tour/)是对Go基础知识的一个很好的逐步介绍。

## 完整的代码

本节包含您通过本教程构建的应用程序的代码。

```go title="main.go" linenums="1"
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

    albums, err := albumsByArtist("John Coltrane")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Albums found: %v\n", albums)

    // Hard-code ID 2 here to test the query.
    alb, err := albumByID(2)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Album found: %v\n", alb)

    albID, err := addAlbum(Album{
        Title:  "The Modern Sound of Betty Carter",
        Artist: "Betty Carter",
        Price:  49.99,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("ID of added album: %v\n", albID)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album

    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
    // An album to hold data from the returned row.
    var alb Album

    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
        if err == sql.ErrNoRows {
            return alb, fmt.Errorf("albumsById %d: no such album", id)
        }
        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }
    return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
    result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
    return id, nil
}
```