+++
title = "测试"
date = 2024-02-04T21:18:56+08:00
weight = 14
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/testing/]({{< ref "/buffalo/guides/testing" >}})

# Testing 测试 

## Running Tests 运行测试 

The Buffalo test runner will make sure your test environment is correct, and run your tests.

​	Buffalo 测试运行器将确保您的测试环境正确无误，并运行您的测试。

For example, if using Pop (database), it will first try to setup your test database by using the schema from the development database. If that doesn’t exist (CI, for example), it will run the migrations against the test database.

​	例如，如果使用 Pop（数据库），它将首先尝试使用开发数据库中的架构来设置您的测试数据库。如果不存在（例如 CI），它将在测试数据库上运行迁移。

The test runner will also make sure to ignore the dreaded `vendor` directory.

​	测试运行器还将确保忽略可怕的 `vendor` 目录。

```console
$ buffalo test

dropped database authrecipe_test
created database authrecipe_test
dumped schema for authrecipe_development
loaded schema for authrecipe_test
go test -p 1 github.com/gobuffalo/authrecipe github.com/gobuffalo/authrecipe/actions github.com/gobuffalo/authrecipe/grifts github.com/gobuffalo/authrecipe/models
?   	github.com/gobuffalo/authrecipe	[no test files]
ok  	github.com/gobuffalo/authrecipe/actions	0.640s
?   	github.com/gobuffalo/authrecipe/grifts	[no test files]
ok  	github.com/gobuffalo/authrecipe/models	0.327s
```

### Execute a single test 执行单个测试 

Since **0.10.2**
自 0.10.2 起



Debugging a specific test is a difficult task, if you must execute all existing tests. You can use the `-m` flag to execute a single test method:

​	如果必须执行所有现有测试，那么调试特定测试是一项困难的任务。您可以使用 `-m` 标志来执行单个测试方法：

```console
$ buffalo test -m "FooMethod"
```

This will iterate through all packages and run any test that matches `FooMethod` in any package.

​	这将在所有包中进行迭代，并在任何包中运行与 `FooMethod` 匹配的任何测试。

Since **0.14.10**
自 0.14.10 起



To limit to one package, specify the package name:

​	要限制为一个包，请指定包名称：

```bash
$ buffalo test models -m "FooMethod"
```

## Test Suites 测试套件 

Buffalo uses the [`github.com/gobuffalo/suite`](https://github.com/gobuffalo/suite) package to create test suites.

​	Buffalo 使用 `github.com/gobuffalo/suite` 包来创建测试套件。

When running a test that is part of the test suite, the following is available to the test:

​	在运行测试套件中的一部分测试时，测试可以使用以下内容：

- The application, `as.App`.
  应用程序， `as.App` 。
- The database, `as.DB` (if using Pop).
  数据库， `as.DB` （如果使用 Pop）。
- The session, as `as.Session`.
  会话，作为 `as.Session` 。
- The [`github.com/stretchr/testify/require`](https://github.com/stretchr/testify) test assertions.
  `github.com/stretchr/testify/require` 测试断言。
- The [`github.com/gobuffalo/httptest`](https://github.com/gobuffalo/httptest) HTTP testing library.
  `github.com/gobuffalo/httptest` HTTP 测试库。 测试示例 #

## Test Example

```go
func (as *ActionSuite) Test_WidgetsResource_Create() {
  // setup a Widget model
  w := &models.Widget{Name: "My Widget"} // make a POST /widgets request
  res := as.HTML("/widgets").Post(w)
  // assert that the response status code was 302 as.Equal(302, res.Code)
  // retrieve the first Widget from the database
  err := as.DB.First(w)
  as.NoError(err)
  as.NotZero(w.ID)
  // assert the Widget title was saved correctly
  as.Equal("My Widget", w.Name)
  // assert the redirect was sent to the place expected
  as.Equal(fmt.Sprintf("/widgets/%s", w.ID), res.Location())
}
```

## Fixtures 固定装置 

Since **0.12.0**
自 0.12.0 起



Often it is useful to load sample data into the database at the start of a test. For example, you need to have a user in the database to log a person into the application, or you need some data in the database to test that a route renders responses correctly. Fixtures help us solve these problems easily.

​	通常，在测试开始时将示例数据加载到数据库中非常有用。例如，您需要在数据库中有一个用户才能将某人登录到应用程序，或者您需要在数据库中有一些数据来测试路由是否正确呈现响应。固定装置可以帮助我们轻松解决这些问题。

Fixtures are `toml` files that are placed in your `fixtures` directory. You can have as many fixture files as you like and they can also be named anyway that you like.

​	固定装置是放置在 `fixtures` 目录中的 `toml` 文件。您可以拥有任意数量的固定装置文件，并且可以根据需要对它们进行命名。

For example, if we have a simple `Widget` model that renders to a page that looks like this:

​	例如，如果我们有一个简单的 `Widget` 模型，它会呈现为如下所示的页面：

```go
type Widget struct {
  ID          uuid.UUID
  CreatedAt   time.Time
  UpdatedAt   time.Time
  Name        string
  Description string
}
```

We can create a fixture file like this:

​	我们可以创建一个这样的固定装置文件：

```toml
[[scenario]]
name = "lots of widgets"

  [[scenario.table]]
    name = "widgets"

    [[scenario.table.row]]
      id = "\<%= uuidNamed("widget") %>"
      name = "This is widget #1"
      description = "some widget body #1"
      created_at = "\<%= now() %>"
      updated_at = "\<%= now() %>"

    [[scenario.table.row]]
      id = "\<%= uuid() %>"
      name = "This is widget #2"
      description = "some widget body #2"
      created_at = "\<%= now() %>"
      updated_at = "\<%= now() %>"
```

When we run our suite, these two records will be created in our test database and we can then test against these records. All you need to do to load the fixture is to reference it by its name with `ActionSuite.LoadFixture`.

​	当我们运行套件时，这两个记录将在我们的测试数据库中创建，然后我们可以针对这些记录进行测试。您需要做的就是通过 `ActionSuite.LoadFixture` 引用其名称来加载固定装置。

```go
func (as *ActionSuite) Test_WidgetsResource_List() {
  as.LoadFixture("lots of widgets")
  res := as.HTML("/widgets").Get()

  as.Equal(200, res.Code)
  body := res.Body.String()
  as.Contains(body, "This is widget #1")
  as.Contains(body, "This is widget #2")
}
```

You can find more detailed information on fixtures in the [gobuffalo/suite repository README](https://github.com/gobuffalo/suite#fixtures-test-data).

​	您可以在 gobuffalo/suite 存储库自述文件中找到有关固定装置的更多详细信息。

## Accessing the Session 访问会话 

Being able to manipulate the session for testing is very important. Test suites in Buffalo give you access to a testing session that you can use.

​	能够操纵会话以进行测试非常重要。Buffalo 中的测试套件允许您访问可以使用的测试会话。

See https://github.com/gobuffalo/authrecipe for a more in-depth example.

​	请参阅 https://github.com/gobuffalo/authrecipe，以获取更深入的示例。

```go
func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
  // get a user from the DB

  // set the user ID onto the session
  as.Session.Set("current_user_id", user.ID)

  res := as.HTML("/").Get()
  as.Equal(200, res.Code)

  // now the user is "logged in"
  as.Contains(res.Body.String(), "Sign Out")

  // clear the session
  as.Session.Clear()
  res = as.HTML("/").Get()
  as.Equal(200, res.Code)

  // now the user is "logged out"
  as.Contains(res.Body.String(), "Sign In")
}
```

## Coverage Reports 覆盖率报告 

The following feature requires the use of **Go 1.10** or a more recent version. Go cover does not support the `./...` operator in older versions, and trying to use it will generate an error.
以下功能需要使用 Go 1.10 或更高版本。Go cover 不支持旧版本中的 `./...` 运算符，尝试使用它会生成错误。

It is possible to generate test coverage reports with buffalo by specifying the `-coverprofile` flag as follows:

​	可以通过指定 `-coverprofile` 标志来生成 buffalo 的测试覆盖率报告，如下所示：

```console
$ buffalo test -coverprofile=c.out ./...
created database authrecipe_test
loaded schema for authrecipe_test
INFO[0010] go test -p 1 -coverprofile=c.out ./...
?       github.com/gobuffalo/authrecipe [no test files]
ok      github.com/gobuffalo/authrecipe/actions 2.770s  coverage: 76.9% of statements
?       github.com/gobuffalo/authrecipe/grifts  [no test files]
ok      github.com/gobuffalo/authrecipe/models  2.609s  coverage: 71.4% of statements
```
