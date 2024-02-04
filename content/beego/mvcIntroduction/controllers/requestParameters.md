+++
title = "请求参数"
date = 2024-02-04T09:57:22+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/mvc/controller/params/](https://beego.wiki/docs/mvc/controller/params/)

# Request parameters 请求参数



## Accept parameters 接受参数

Beego will automatically parse data passed by user from GET, POST and other methods. This data can be accessed using:

&zeroWidthSpace;Beego 会自动解析用户通过 GET、POST 和其他方法传递的数据。可以使用以下方法访问这些数据：

- GetString(key string) string
- GetStrings(key string) []string
- GetInt(key string) (int, error)
- GetInt8(key string) (int8, error)
- GetInt16(key string) (int16, error)
- GetInt32(key string) (int32, error)
- GetInt64(key string) (int64, error)
- GetUint8(key string) (uint8, error)
- GetUint16(key string) (uint16, error)
- GetUint32(key string) (uint32, error)
- GetUint64(key string) (uint64, error)
- GetBool(key string) (bool, error)
- GetFloat(key string) (float64, error)

For example:

&zeroWidthSpace;例如：

```go
func (this *MainController) Post() {
	jsoninfo := this.GetString("jsoninfo")
	if jsoninfo == "" {
		this.Ctx.WriteString("jsoninfo is empty")
		return
	}
}
```

More information about the request can be retrieved by accessing `this.Ctx.Request`. For more details see [Request](http://gowalker.org/net/http#Request).

&zeroWidthSpace;可以通过访问 `this.Ctx.Request` 来检索有关请求的更多信息。有关更多详细信息，请参阅请求。

## Parse to struct 解析为结构

Data submitted from a form may be assigned to a struct by mapping struct fields to the form’s input elements and parsing all data into a struct.

&zeroWidthSpace;从表单提交的数据可以通过将结构域映射到表单的输入元素并将所有数据解析为结构来分配给结构。

Define struct:

&zeroWidthSpace;定义结构：

```go
type User struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}
```

Define form:

&zeroWidthSpace;定义形式：

```
<form id="user">
	name：<input name="username" type="text" />
	age：<input name="age" type="text" />
	email：<input name="Email" type="text" />
	<input type="submit" value="submit" />
</form>
```

Parsing in Controller:

&zeroWidthSpace;在控制器中解析：

```go
func (this *MainController) Post() {
	u := User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
	}
}
```

Notes:

&zeroWidthSpace;注释：

- The same tag is used for the definition of structTag form and [renderform method](https://beego.wiki/docs/mvc/view/view#renderform).
  相同的标签用于定义 structTag 形式和 renderform 方法。
- If there is a form tag after the key while defining the struct, the value in the form which has the same name as that tag will be assigned. Otherwise, the value in the form which has the same name as that field name will be assigned. In the above example, Form values username and age will be assigned to Name and Age in user struct and Email will be assigned to Email in struct.
  如果在定义结构时键后有一个 form 标签，则将分配具有与该标签相同的名称的 form 中的值。否则，将分配具有与该字段名称相同的名称的 form 中的值。在上面的示例中，表单值 username 和 age 将分配给 user 结构中的 Name 和 Age，Email 将分配给结构中的 Email。
- While calling the method ParseForm of the Controller the parameter passed in must be a pointer to a struct. Otherwise, the assignment will fail and will return a `xx must be a struct pointer` error.
  在调用控制器的 ParseForm 方法时，传入的参数必须是结构的指针。否则，分配将失败并返回 `xx must be a struct pointer` 错误。
- Fields can be ignored by using lowercase for that field or by using `-` as the value of the tag.
  可以通过使用小写字母表示该字段或使用 `-` 作为标签的值来忽略字段。

## Automatic Parameter Routing 参数路由

Automatic parameter routing removes the need for boilerplate code like `this.GetString(..)`, `this.GetInt(..)` etc. Instead https parameters are injected directly as method parameters and the method return values are rendered as http responses. This works in conjunction with annotations to create a seamless integration.

&zeroWidthSpace;参数消除了对样板代码（如 `this.GetString(..)` 、 `this.GetInt(..)` 等）的需求。相反，https 参数直接作为方法参数，方法返回值呈现为 http 响应。这与注释结合使用，可以实现无缝集成。

### How does it work? 它是如何工作的？

Start by defining a regular controller method with a `@router` annotation and add parameters to the method signature

&zeroWidthSpace;通过使用 `@router` 注释定义一个常规控制器方法，并向方法签名添加参数

```go
// @router /tasks
func (c *TaskController) MyMethod(id int) {
...
}
```

When an http request comes in that matches the defined routing Beego will scan the parameters in the method signature and try to find matching http request paramters, where method parameter name is the http request parameter name. Beego will then convert them to the correct parameter type and pass them to your method. By default Beego will look for parameters in the quey string (when using `GET`) or form data (when using `POST`). If your routing definition contains parameters Beego will automatically search for them in the path:

&zeroWidthSpace;当一个匹配定义路由的 http 请求进来时，Beego 将扫描方法签名中的参数，并尝试查找匹配的 http 请求参数，其中方法参数名称是 http 请求参数名称。然后，Beego 会将它们转换为正确的参数类型，并将其传递给您的方法。默认情况下，Beego 将在查询字符串（使用 `GET` 时）或表单数据（使用 `POST` 时）中查找参数。如果您的路由定义包含参数，Beego 将自动在路径中搜索它们：

```go
// @router /task/:id
func (c *TaskController) MyMethod(id int) {
...
}
```

Annotations can also be used to indicate a parameter is passed in a header or in the request body. Bego will search for it accordingly.

&zeroWidthSpace;注释还可用于指示参数是在标头中还是在请求正文中传递的。Bego 将相应地搜索它。

If a parameter is not found in the http request it will be passed to your controller method as a zero value (i.e. 0 for int, false for bool etc.). If a default value for that parameter has been defined in annotations, Beego will pass that default value if it is missing. To differentiate between missing parameters and default values define the parameter as a pointer, e.g.:

&zeroWidthSpace;如果在 http 请求中找不到参数，它将作为零值传递给您的控制器方法（即 int 为 0，bool 为 false 等）。如果在注释中定义了该参数的默认值，则 Beego 将在缺少该值时传递该默认值。为了区分缺失的参数和默认值，将参数定义为指针，例如：

```go
// @router /tasks
func (c *TaskController) MyMethod(id *int) {
...
}
```

If the parameter in the above case was missing, `id` would be null. If the parameter exists and equals to zero, `id` would be 0. When using annotations to create swagger documentation a parameter can be marked as `required`. If the parameter is missing in the request a `400 Bad Request` error will be returned to the client:

&zeroWidthSpace;如果上述情况中的参数缺失， `id` 将为 null。如果参数存在且等于零， `id` 将为 0。在使用注释创建 swagger 文档时，可以将参数标记为 `required` 。如果请求中缺少参数，将向客户端返回 `400 Bad Request` 错误：

```go
// @Param   id     query   int true       "task id"
// @router /tasks
func (c *TaskController) MyMethod(id *int) {
...
}
```

If Beego can not convert the parameter to the requested type (i.e. if a string is passed that can not be parsed as an integer) an error will be returned to the client.

&zeroWidthSpace;如果 Beego 无法将参数转换为请求的类型（即，如果传递的字符串无法解析为整数），将向客户端返回错误。

The following table shows which types are supported and how they are parsed:

&zeroWidthSpace;下表显示了支持哪些类型以及如何解析这些类型：

| Data Type 数据类型                                           | Location 位置           | Example 示例                                                 | Comment 注释                                                 |
| ------------------------------------------------------------ | ----------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| int, int64, uint etc. int、int64、uint 等                    | anywhere 任意位置       | “1”,"-100" “1”，“-100”                                       | Uses `strconv.Atoi(value)` 在任何地方使用 `strconv.Atoi(value)` |
| float32,float64 float32、float64                             | anywhere “1.5”, “-3.5”  | “1.5”, “-3.5” 在任何地方使用                                 | Uses `strconv.ParseFloat()` “1”, “T”, “false”                |
| bool                                                         | anywhere 在任何地方使用 | “1”, “T”, “false” time.Time                                  | Uses `strconv.ParseBool()` “2017-01-01” “2017-01-01T00:00:00Z” |
| time.Time                                                    | anywhere                | “2017-01-01” “2017-01-01T00:00:00Z”                          | Uses RFC3339 or short date format (`"2006-01-02"`) when parsing 解析 |
| []string, []int etc. []string, []int 等时，使用 RFC3339 或短日期格式 ( )。 | query 查询              | “A,B,C” “1,2,3”                                              | Any type is supported as a slice. When it is located in the query string, it is parsed as a comma separated list 任何类型都支持作为切片。当它位于查询字符串中时，它将被解析为逗号分隔的列表 |
| []string, []int etc. []string, []int 等。                    | body 正文               | [“A”,“B”,“C”] [1,2,3]                                        | When slices are located in the request body they are parsed as JSON arrays 当切片位于请求正文中时，它们将被解析为 JSON 数组 |
| []byte                                                       | anywhere                | “ABC”                                                        | byte[] is not treated as an array but as a string byte[] 不被视为数组，而被视为字符串 |
| *int, *string, *float etc. *int、*string、*float 等          | anywhere                | Pointers will receive null if the parameter is missing from the request otherwise, it will behave the same as defined in the other rows 如果请求中缺少参数，指针将接收 null，否则，它将与其他行中定义的行为相同 |                                                              |
| structs, all others 结构体，所有其他结构体                   | anywhere                | {“X”:“Y”}                                                    | structs and other types (e.g. maps) are always parsed as JSON using `json.Unmarshal()` 使用 `json.Unmarshal()` 始终将结构体和其他类型（例如映射）解析为 JSON |

### How are method return values handled? 方法返回值如何处理？

Method return values are handled automatically in the same manner as parameters. A method can have one or more return values and Beego will render all of them to the response. The best practice is to define one result as a ‘regular’ type (i.e. a map, a struct or any other data type) and another as an error data type:

&zeroWidthSpace;方法返回值以与参数相同的方式自动处理。一个方法可以有一个或更多个返回值，Beego 将把它们全部呈现到相应内容中。最佳做法是将一个结果定义为“常规”类型（即映射、结构或任何其他数据类型），另一个定义为错误数据类型：

```go
// @Param   id     query   int true       "task id"
// @router /tasks
func (c *TaskController) MyMethod(id *int) (*MyModel, error) {
...
}
```

In the code above the method can return three different results:

&zeroWidthSpace;在上面的代码中，该方法可以返回三个不同的结果：

- Only `MyModel` (nil `error`)
  仅 `MyModel` （nil `error` ）
- Only `error` (nil `MyModel`)
  仅 `error` （nil `MyModel` ）
- Both `MyModel` and `error`
  `MyModel` 和 `error` 同时存在

When a regular type is returned it is rendered directly as JSON, and when an error is returned it is rendered as an http status code. Beego will handle all cases correctly and supports returning both response body and http error if both values are non-nil.

&zeroWidthSpace;当返回常规类型时，它将直接呈现为 JSON，当返回错误时，它将呈现为 http 状态代码。Beego 将正确处理所有情况，并支持返回相应内容和 http 错误（如果两个值均为非 nil）。

A few helper types will return common http status codes easily. For example, `404 Not Found`, `302 Redirect` or other http status codes like in the following example:

&zeroWidthSpace;一些帮助程序类型将很容易地返回常见的 http 状态代码。例如， `404 Not Found` 、 `302 Redirect` 或其他 http 状态代码，如下例所示：

```go
func (c *TaskController) MyMethod(id *int) (*MyModel, error) {
  if /* not found */ {
    return nil, context.NotFound
  } else if /* some error */ {
    return nil, context.StatusCode(401)
  } else {
  	return &MyModel{}, nil
  }
}
```

### How annotations work in conjuction with method parameters? 注释如何与方法参数一起使用？

Automatic Parameter Routing works best together with `@Param` annotations. The following features are supported with annotations:

&zeroWidthSpace;自动参数路由最适合与 `@Param` 注释一起使用。注释支持以下功能：

- If a parameter is marked as required, Beego will return an error if the parameter is not present in the http request:
  如果某个参数被标记为必需，如果该参数不存在于 http 请求中，Beego 将返回一个错误：

```go
// @Param   brand_id    query   int true       "brand id"
```

(the `true` option in the annotation above indicates that brand_id is a required parameter)

&zeroWidthSpace;（上面的注释中的 `true` 选项表示 brand_id 是一个必需参数）

- If a parameter has a default value and it does not exist in the http request, Beego will pass that default value to the method:
  如果某个参数具有默认值，并且它不存在于 http 请求中，Beego 将把该默认值传递给该方法：

```go
// @Param   brand_id    query   int false  5  "brand id"
```

(the `5` in the annotation above indicates that this is the default value for that parameter)

&zeroWidthSpace;（上面的注释中的 `5` 表示这是该参数的默认值）

- The location parameter in the annotation indicates where beego will search for that parameter in the request (i.e. query, header, body etc.)
  注释中的 location 参数表示 beego 将在请求中搜索该参数的位置（即查询、标头、正文等）

```go
// @Param   brand_id    path   	int 	true  "brand id"
// @Param   category    query 	string	false "category" 
// @Param   token	header  string	false "auth token"
// @Param   task	body	{models.Task} false "the task object"
```

- If a parameter name in the http request is different from the method parameter name, you can “redirect” the parameter using the `=>` notation. This is useful, for example, When a header name is `X-Token` and the method parameter is named `x_token`:
  如果 http 请求中的参数名称与方法参数名称不同，您可以使用 `=>` 符号“重定向”该参数。例如，当标头名称为 `X-Token` 而方法参数名称为 `x_token` 时，这很有用：

```go
// @Param   X-Token=>x_token	header  string	false "auth token"
```

- A parameter swagger data type can be inferred from the method to make maintainance easier. Use the `auto` data type and Beego will generate the correct swagger documentation:
  可以从方法中推断出参数 swagger 数据类型，以便于维护。使用 `auto` 数据类型，Beego 将生成正确的 swagger 文档：

```go
// @Param   id     query   auto true       "task id"
// @router /tasks
func (c *TaskController) MyMethod(id int) (*MyModel, error) {
...
}
```

## Retrieving data from request body 从请求正文中检索数据

In API application development always use `JSON` or `XML` as the data type. To retrieve the data from the request body:

&zeroWidthSpace;在 API 应用程序开发中，始终使用 `JSON` 或 `XML` 作为数据类型。要从请求正文中检索数据：

1. Set `copyrequestbody = true` in configuration file.
   在配置文件中设置 `copyrequestbody = true` 。
2. Then in the Controller you can
   然后在控制器中可以

```go
func (this *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	this.Data["json"] = map[string]interface{}{"ObjectId": objectid }
	this.ServeJSON()
}
```

## Uploading files 上传文件

To upload files with Beego set attribute `enctype="multipart/form-data"` in your form.

&zeroWidthSpace;要使用 Beego 上传文件，请在表单中设置属性 `enctype="multipart/form-data"` 。

Usually an uploaded file is stored in the system memory, but if the file size is larger than the memory size limitation in the configuration file, the file will be stored in a temporary file. The default memory size is 64M but can be changed using (bit shift):

&zeroWidthSpace;通常，上传的文件存储在系统内存中，但如果文件大小大于配置文件中的内存大小限制，则文件将存储在临时文件中。默认内存大小为 64M，但可以使用（位移）进行更改：

```
beego.MaxMemory = 1<<22
```

Or it can be set manualy in the configuration file (bit shift):

&zeroWidthSpace;或者可以在配置文件中手动设置（位移）：

```
maxmemory = 1<<22
```

In v2.x, there is another parameter `MaxUploadSize` used to limit the max size of uploading files.

&zeroWidthSpace;在 v2.x 中，还有另一个参数 `MaxUploadSize` 用于限制上传文件的大小。

If you upload multiple files in one request, it limits the sum size of those files.

&zeroWidthSpace;如果在一个请求中上传多个文件，它会限制这些文件的总大小。

Usually, `web.BConfig.MaxMemory` should be less than `web.BConfig.MaxUploadSize`:

&zeroWidthSpace;通常， `web.BConfig.MaxMemory` 应小于 `web.BConfig.MaxUploadSize` ：

1. if file size < `MaxMemory`, handling file in memory;
   如果文件大小 < `MaxMemory` , 内存中处理文件；
2. `MaxMemory` < file size < `MaxUploadSize`, handling file by using temporary directory.
   `MaxMemory` < 文件大小 < `MaxUploadSize` , 使用临时目录处理文件。
3. file size > `MaxUploadSize`, return 413;
   文件大小 > `MaxUploadSize` , 返回 413；

Beego provides three functions to handle file uploads:

&zeroWidthSpace;Beego 提供了三个函数来处理文件上传：

- GetFile(key string) (multipart.File, *multipart.FileHeader, error)

This method is used to read the file name `the_file` from form and return the information. The uploaded file can then be processed based on this information, such as filter or save the file.

&zeroWidthSpace;此方法用于从表单中读取文件名 `the_file` 并返回信息。然后可以根据此信息处理上传的文件，例如过滤或保存文件。

- GetFiles(key string) ([]*multipart.FileHeader, error)

This method returns all the multi-upload files:

&zeroWidthSpace;此方法返回所有多上传文件：

```go
func (m *MainController) Post() {
	// 'files' is the name of the multipart form input
	files, err := m.GetFiles("files")
	if err != nil {
		logger.Error(err.Error())
	}
	... do something with files
```

- SaveToFile(fromfile, tofile string) error

This method implements the saving function based on the method `GetFile`

&zeroWidthSpace;此方法基于方法 `GetFile` 实现保存功能

Here is an example of saving a file:

&zeroWidthSpace;以下是如何保存文件的示例：

```go
func (this *MainController) Post() {
	this.SaveToFile("the_file","/var/www/uploads/uploaded_file.txt")
}
```

## Data Bind 数据绑定

Data bind lets the user bind the request data to a variable, the request url as follows:

&zeroWidthSpace;数据绑定允许用户将请求数据绑定到变量，请求 URL 如下：

```
?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie
var id int
ctx.Input.Bind(&id, "id")  // id ==123

var isok bool
ctx.Input.Bind(&isok, "isok")  // isok ==true

var ft float64
ctx.Input.Bind(&ft, "ft")  // ft ==1.2

ol := make([]int, 0, 2)
ctx.Input.Bind(&ol, "ol")  // ol ==[1 2]

ul := make([]string, 0, 2)
ctx.Input.Bind(&ul, "ul")  // ul ==[str array]

user struct{Name}
ctx.Input.Bind(&user, "user")  // user =={Name:"astaxie"}
```