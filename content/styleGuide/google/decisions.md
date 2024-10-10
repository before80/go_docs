+++
title = "决策"
date = 2024-01-22T10:01:26+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# Go Style Decisions - Go 风格决策

原文：[https://google.github.io/styleguide/go/decisions](https://google.github.io/styleguide/go/decisions)

> **Note:** This is part of a series of documents that outline [Go Style](https://google.github.io/styleguide/go/index) at Google. This document is **[normative](https://google.github.io/styleguide/go/index#normative) but not [canonical](https://google.github.io/styleguide/go/index#canonical)**, and is subordinate to the [core style guide](https://google.github.io/styleguide/go/guide). See [the overview](https://google.github.io/styleguide/go/index#about) for more information.
>
> ​	注意：这是概述 Google 中 Go 风格的一系列文档的一部分。此文档具有规范性，但不是规范性的，并且从属于核心风格指南。有关更多信息，请参阅概述。



## About 关于

This document contains style decisions intended to unify and provide standard guidance, explanations, and examples for the advice given by the Go readability mentors.

​	此文档包含旨在统一并为 Go 可读性导师提供的建议提供标准指导、说明和示例的风格决策。

This document is **not exhaustive** and will grow over time. In cases where [the core style guide](https://google.github.io/styleguide/go/guide) contradicts the advice given here, **the style guide takes precedence**, and this document should be updated accordingly.

​	此文档并不详尽，并且会随着时间的推移而增长。在核心风格指南与此处提供的建议相矛盾的情况下，风格指南优先，并且此文档应相应更新。

See [the Overview](https://google.github.io/styleguide/go#about) for the full set of Go Style documents.

​	有关完整的 Go 风格文档集，请参阅概述。

The following sections have moved from style decisions to another part of the guide:

​	以下部分已从风格决策移至指南的另一部分：

- **MixedCaps**: see [guide#mixed-caps](https://google.github.io/styleguide/go/guide#mixed-caps)

  ​	MixedCaps：请参阅 guide#mixed-caps

- **Formatting**: see [guide#formatting](https://google.github.io/styleguide/go/guide#formatting)

  ​	格式化：请参阅 guide#formatting

- **Line Length**: see [guide#line-length](https://google.github.io/styleguide/go/guide#line-length)

  ​	行长：请参阅 guide#line-length



## Naming 命名

See the naming section within [the core style guide](https://google.github.io/styleguide/go/guide#naming) for overarching guidance on naming. The following sections provide further clarification on specific areas within naming.

​	有关命名的总体指导，请参阅核心风格指南中的命名部分。以下部分提供了有关命名中特定领域的进一步说明。



### Underscores 下划线

Names in Go should in general not contain underscores. There are three exceptions to this principle:

​	Go 中的名称通常不应包含下划线。此原则有三个例外：

1. Package names that are only imported by generated code may contain underscores. See [package names](https://google.github.io/styleguide/go/decisions#package-names) for more detail around how to choose multi-word package names.
   仅由生成代码导入的包名称可能包含下划线。有关如何选择多词包名称的更多详细信息，请参阅包名称。
2. Test, Benchmark and Example function names within `*_test.go` files may include underscores.
   `*_test.go` 文件中的测试、基准和示例函数名称可能包含下划线。
3. Low-level libraries that interoperate with the operating system or cgo may reuse identifiers, as is done in [`syscall`](https://pkg.go.dev/syscall#pkg-constants). This is expected to be very rare in most codebases.
   与操作系统或 cgo 相互操作的底层库可能会重用标识符，就像在 `syscall` 中所做的那样。预计在大多数代码库中这种情况非常罕见。



### Package names 包名称



Go package names should be short and contain only lowercase letters. A package name composed of multiple words should be left unbroken in all lowercase. For example, the package [`tabwriter`](https://pkg.go.dev/text/tabwriter) is not named `tabWriter`, `TabWriter`, or `tab_writer`.

​	Go 包名称应简短，且仅包含小写字母。由多个单词组成的包名称应保持完整，全部采用小写。例如，包 `tabwriter` 的名称不是 `tabWriter` 、 `TabWriter` 或 `tab_writer` 。

Avoid selecting package names that are likely to be [shadowed](https://google.github.io/styleguide/go/best-practices#shadowing) by commonly used local variable names. For example, `usercount` is a better package name than `count`, since `count` is a commonly used variable name.

​	避免选择可能被常用的局部变量名称所掩盖的包名称。例如， `usercount` 是比 `count` 更好的包名称，因为 `count` 是一个常用的变量名称。

Go package names should not have underscores. If you need to import a package that does have one in its name (usually from generated or third party code), it must be renamed at import time to a name that is suitable for use in Go code.

​	Go 包名不应包含下划线。如果您需要导入一个名称中包含下划线的包（通常来自生成的或第三方代码），则必须在导入时将其重命名为适合在 Go 代码中使用的名称。

An exception to this is that package names that are only imported by generated code may contain underscores. Specific examples include:

​	对此的一个例外是，仅由生成的代码导入的包名可以包含下划线。具体示例包括：

- Using the `_test` suffix for an external test package, for example an integration test

  ​	对外部测试包使用 `_test` 后缀，例如集成测试

- Using the `_test` suffix for [package-level documentation examples](https://go.dev/blog/examples)

  ​	对包级文档示例使用 `_test` 后缀

Avoid uninformative package names like `util`, `utility`, `common`, `helper`, and so on. See more about [so-called “utility packages”](https://google.github.io/styleguide/go/best-practices#util-packages).

​	避免使用无意义的包名，如 `util` 、 `utility` 、 `common` 、 `helper` 等。请参阅有关所谓的“实用程序包”的更多信息。

When an imported package is renamed (e.g. `import foopb "path/to/foo_go_proto"`), the local name for the package must comply with the rules above, as the local name dictates how the symbols in the package are referenced in the file. If a given import is renamed in multiple files, particularly in the same or nearby packages, the same local name should be used wherever possible for consistency.

​	当重命名导入的包时（例如 `import foopb "path/to/foo_go_proto"` ），包的本地名称必须遵守上述规则，因为本地名称决定了如何引用文件中包中的符号。如果在多个文件中重命名给定的导入，尤其是在相同或相邻的包中，则应尽可能使用相同的本地名称以保持一致性。

See also: [Go blog post about package names](https://go.dev/blog/package-names).

​	另请参阅：有关包名的 Go 博客文章。



### Receiver names 接收者名称



[Receiver](https://golang.org/ref/spec#Method_declarations) variable names must be:

​	接收者变量名称必须：

- Short (usually one or two letters in length)
  简短（通常长度为一个或两个字母）
- Abbreviations for the type itself
  类型的缩写本身
- Applied consistently to every receiver for that type
  对该类型的每个接收器一致应用

| Long Name 长名称            | Better Name 更好的名称    |
| --------------------------- | ------------------------- |
| `func (tray Tray)`          | `func (t Tray)`           |
| `func (info *ResearchInfo)` | `func (ri *ResearchInfo)` |
| `func (this *ReportWriter)` | `func (w *ReportWriter)`  |
| `func (self *Scanner)`      | `func (s *Scanner)`       |



### Constant names 常量名称

Constant names must use [MixedCaps](https://google.github.io/styleguide/go/guide#mixed-caps) like all other names in Go. ([Exported](https://tour.golang.org/basics/3) constants start with uppercase, while unexported constants start with lowercase.) This applies even when it breaks conventions in other languages. Constant names should not be a derivative of their values and should instead explain what the value denotes.

​	常量名称必须像 Go 中所有其他名称一样使用 MixedCaps。（导出的常量以大写字母开头，而未导出的常量以小写字母开头。）即使它打破了其他语言中的约定，这也适用。常量名称不应是其值的派生词，而应解释该值表示什么。

``` go
// Good:
const MaxPacketSize = 512

const (
    ExecuteBit = 1 << iota
    WriteBit
    ReadBit
)
```

Do not use non-MixedCaps constant names or constants with a `K` prefix.

​	不要使用非 MixedCaps 常量名或带有 `K` 前缀的常量。

``` go
// Bad:
const MAX_PACKET_SIZE = 512
const kMaxBufferSize = 1024
const KMaxUsersPergroup = 500
```

Name constants based on their role, not their values. If a constant does not have a role apart from its value, then it is unnecessary to define it as a constant.

​	根据角色而不是值命名常量。如果常量除了其值之外没有角色，那么定义它为常量是没有必要的。

``` go
// Bad:
const Twelve = 12

const (
    UserNameColumn = "username"
    GroupColumn    = "group"
)
```



### Initialisms 首字母缩写词



Words in names that are initialisms or acronyms (e.g., `URL` and `NATO`) should have the same case. `URL` should appear as `URL` or `url` (as in `urlPony`, or `URLPony`), never as `Url`. This also applies to `ID` when it is short for “identifier”; write `appID` instead of `appId`.

​	名称中是首字母缩写词或缩略词的单词（例如， `URL` 和 `NATO` ）应采用相同的大小写。 `URL` 应显示为 `URL` 或 `url` （如 `urlPony` 或 `URLPony` 中所示），绝不应显示为 `Url` 。当 `ID` 是“标识符”的缩写时，此规则也适用；请编写 `appID` ，而不是 `appId` 。

- In names with multiple initialisms (e.g. `XMLAPI` because it contains `XML` and `API`), each letter within a given initialism should have the same case, but each initialism in the name does not need to have the same case.
  在具有多个首字母缩写词的名称中（例如 `XMLAPI` ，因为它包含 `XML` 和 `API` ），给定首字母缩写词中的每个字母应采用相同的大小写，但名称中的每个首字母缩写词不必采用相同的大小写。
- In names with an initialism containing a lowercase letter (e.g. `DDoS`, `iOS`, `gRPC`), the initialism should appear as it would in standard prose, unless you need to change the first letter for the sake of [exportedness](https://golang.org/ref/spec#Exported_identifiers). In these cases, the entire initialism should be the same case (e.g. `ddos`, `IOS`, `GRPC`).
  在名称中包含小写字母的首字母缩写词（例如 `DDoS` 、 `iOS` 、 `gRPC` ）中，首字母缩写词应按标准散文中的形式显示，除非您需要为了可导出性而更改第一个字母。在这些情况下，整个首字母缩写词应采用相同的大小写（例如 `ddos` 、 `IOS` 、 `GRPC` ）。

| Initialism(s) 首字母缩写词 | Scope 范围          | Correct 正确 | Incorrect 不正确                       |
| -------------------------- | ------------------- | ------------ | -------------------------------------- |
| XML API                    | Exported 导出的     | `XMLAPI`     | `XmlApi`, `XMLApi`, `XmlAPI`, `XMLapi` |
| XML API                    | Unexported 未导出的 | `xmlAPI`     | `xmlapi`, `xmlApi`                     |
| iOS                        | Exported 导出的     | `IOS`        | `Ios`, `IoS`                           |
| iOS                        | Unexported 未导出的 | `iOS`        | `ios`                                  |
| gRPC                       | Exported 导出的     | `GRPC`       | `Grpc`                                 |
| gRPC                       | Unexported 未导出的 | `gRPC`       | `grpc`                                 |
| DDoS                       | Exported 导出       | `DDoS`       | `DDOS`, `Ddos`                         |
| DDoS                       | Unexported 未导出   | `ddos`       | `dDoS`, `dDOS`                         |



### Getters 获取器



Function and method names should not use a `Get` or `get` prefix, unless the underlying concept uses the word “get” (e.g. an HTTP GET). Prefer starting the name with the noun directly, for example use `Counts` over `GetCounts`.

​	函数和方法名称不应使用前缀 `Get` 或 `get` ，除非底层概念使用单词“get”（例如 HTTP GET）。最好直接以名词开头，例如使用 `Counts` 而不是 `GetCounts` 。

If the function involves performing a complex computation or executing a remote call, a different word like `Compute` or `Fetch` can be used in place of `Get`, to make it clear to a reader that the function call may take time and could block or fail.

​	如果函数涉及执行复杂计算或执行远程调用，可以使用 `Compute` 或 `Fetch` 等其他单词代替 `Get` ，以便向读者明确说明函数调用可能需要时间，并且可能会阻塞或失败。



### Variable names 变量名称



The general rule of thumb is that the length of a name should be proportional to the size of its scope and inversely proportional to the number of times that it is used within that scope. A variable created at file scope may require multiple words, whereas a variable scoped to a single inner block may be a single word or even just a character or two, to keep the code clear and avoid extraneous information.

​	一般经验法则是，名称的长度应与其作用域的大小成正比，而与其在该作用域内使用次数成反比。在文件作用域中创建的变量可能需要多个单词，而作用域为单个内部块的变量可能是一个单词，甚至只是几个字符，以保持代码清晰并避免无关信息。

Here is a rough baseline. These numeric guidelines are not strict rules. Apply judgement based on context, [clarity](https://google.github.io/styleguide/go/guide#clarity), and [concision](https://google.github.io/styleguide/go/guide#concision).

​	以下是一个粗略的基准。这些数字准则并非严格的规则。应根据上下文、清晰度和简洁性进行判断。

- A small scope is one in which one or two small operations are performed, say 1-7 lines.
  小作用域是指执行一两个小操作的作用域，例如 1-7 行。
- A medium scope is a few small or one large operation, say 8-15 lines.
  中等范围是几个小操作或一个大操作，比如 8-15 行。
- A large scope is one or a few large operations, say 15-25 lines.
  大范围是一个或几个大操作，比如 15-25 行。
- A very large scope is anything that spans more than a page (say, more than 25 lines).
  非常大的范围是跨越一页以上的内容（比如超过 25 行）。

A name that might be perfectly clear (e.g., `c` for a counter) within a small scope could be insufficient in a larger scope and would require clarification to remind the reader of its purpose further along in the code. A scope in which there are many variables, or variables that represent similar values or concepts, may necessitate longer variable names than the scope suggests.

​	在小范围内可能非常清晰的名称（例如，计数器的 `c` ）在更大的范围内可能不够用，需要进行说明以提醒读者其在代码中的用途。包含许多变量或表示相似值或概念的变量的范围可能需要比范围建议的更长的变量名称。

The specificity of the concept can also help to keep a variable’s name concise. For example, assuming there is only a single database in use, a short variable name like `db` that might normally be reserved for very small scopes may remain perfectly clear even if the scope is very large. In this case, a single word `database` is likely acceptable based on the size of the scope, but is not required as `db` is a very common shortening for the word with few alternate interpretations.

​	概念的特殊性也有助于保持变量名称的简洁。例如，假设只使用一个数据库，像 `db` 这样的短变量名称（通常可能为非常小的范围保留）即使范围非常大也可能非常清晰。在这种情况下，根据范围的大小，一个单词 `database` 可能可以接受，但不是必需的，因为 `db` 是该单词的非常常见的缩写，几乎没有其他解释。

The name of a local variable should reflect what it contains and how it is being used in the current context, rather than where the value originated. For example, it is often the case that the best local variable name is not the same as the struct or protocol buffer field name.

​	局部变量的名称应反映其包含的内容以及在当前上下文中如何使用它，而不是值的来源。例如，最佳局部变量名称通常与结构或协议缓冲区字段名称不同。

In general:

​	一般来说：

- Single-word names like `count` or `options` are a good starting point.
  像 `count` 或 `options` 这样的单字名称是一个良好的起点。

- Additional words can be added to disambiguate similar names, for example `userCount` and `projectCount`.
  可以添加其他字词来消除歧义，例如 `userCount` 和 `projectCount` 。

- Do not simply drop letters to save typing. For example `Sandbox` is preferred over `Sbx`, particularly for exported names.
  不要仅仅为了节省打字而省略字母。例如， `Sandbox` 优于 `Sbx` ，特别是对于导出的名称。

- Omit

   

  types and type-like words

   

  from most variable names.

  
  从大多数变量名称中省略类型和类似类型的字词。

  - For a number, `userCount` is a better name than `numUsers` or `usersInt`.
    对于一个数字， `userCount` 是比 `numUsers` 或 `usersInt` 更好的名称。
  - For a slice, `users` is a better name than `userSlice`.
    对于一个切片， `users` 是比 `userSlice` 更好的名称。
  - It is acceptable to include a type-like qualifier if there are two versions of a value in scope, for example you might have an input stored in `ageString` and use `age` for the parsed value.
    如果作用域中存在两个版本的值，则可以包含一个类型限定符，例如，您可能在 `ageString` 中存储了一个输入，并使用 `age` 作为解析后的值。

- Omit words that are clear from the [surrounding context](https://google.github.io/styleguide/go/decisions#repetitive-in-context). For example, in the implementation of a `UserCount` method, a local variable called `userCount` is probably redundant; `count`, `users`, or even `c` are just as readable.
  省略周围上下文中清楚的单词。例如，在 `UserCount` 方法的实现中，名为 `userCount` 的局部变量可能是多余的； `count` 、 `users` ，甚至 `c` 都一样易读。



#### Single-letter variable names 单字母变量名

Single-letter variable names can be a useful tool to minimize [repetition](https://google.github.io/styleguide/go/decisions#repetition), but can also make code needlessly opaque. Limit their use to instances where the full word is obvious and where it would be repetitive for it to appear in place of the single-letter variable.

​	单字母变量名可以成为一个有用的工具来最小化重复，但也可以使代码不必要地不透明。将其用法限制在全词显而易见且重复出现在单字母变量位置的情况。

In general:

​	一般来说：

- For a [method receiver variable](https://google.github.io/styleguide/go/decisions#receiver-names), a one-letter or two-letter name is preferred.
  对于方法接收器变量，首选一个字母或两个字母的名称。

- Using familiar variable names for common types is often helpful:

  
  通常，对常见类型使用熟悉的变量名会有所帮助：

  - `r` for an `io.Reader` or `*http.Request`
    `r` 表示 `io.Reader` 或 `*http.Request`
  - `w` for an `io.Writer` or `http.ResponseWriter`
    `w` 表示 `io.Writer` 或 `http.ResponseWriter`

- Single-letter identifiers are acceptable as integer loop variables, particularly for indices (e.g., `i`) and coordinates (e.g., `x` and `y`).
  单字母标识符可作为整数循环变量，特别是对于索引（例如， `i` ）和坐标（例如， `x` 和 `y` ）。

- Abbreviations can be acceptable loop identifiers when the scope is short, for example `for _, n := range nodes { ... }`.
  当范围较短时，缩写可以作为可接受的循环标识符，例如 `for _, n := range nodes { ... }` 。



### Repetition 重复

A piece of Go source code should avoid unnecessary repetition. One common source of this is repetitive names, which often include unnecessary words or repeat their context or type. Code itself can also be unnecessarily repetitive if the same or a similar code segment appears multiple times in close proximity.

​	一段 Go 源代码应避免不必要的重复。其中一个常见来源是重复的名称，其中通常包括不必要的单词或重复其上下文或类型。如果相同或类似的代码段在近距离多次出现，则代码本身也可能不必要地重复。

Repetitive naming can come in many forms, including:

​	重复命名可以有多种形式，包括：



#### Package vs. exported symbol name 包与导出符号名称

When naming exported symbols, the name of the package is always visible outside your package, so redundant information between the two should be reduced or eliminated. If a package exports only one type and it is named after the package itself, the canonical name for the constructor is `New` if one is required.

​	在命名导出符号时，包的名称始终在包外可见，因此应减少或消除两者之间的冗余信息。如果一个包只导出一个类型，并且它以包本身命名，那么如果需要构造函数，则其规范名称为 `New` 。

> **Examples:** Repetitive Name -> Better Name
>
> ​	示例：重复名称 -> 更好的名称
>
> - `widget.NewWidget` -> `widget.New`
> - `widget.NewWidgetWithName` -> `widget.NewWithName`
> - `db.LoadFromDatabase` -> `db.Load`
> - `goatteleportutil.CountGoatsTeleported` -> `gtutil.CountGoatsTeleported` or `goatteleport.Count`
>   `goatteleportutil.CountGoatsTeleported` -> `gtutil.CountGoatsTeleported` 或 `goatteleport.Count`
> - `myteampb.MyTeamMethodRequest` -> `mtpb.MyTeamMethodRequest` or `myteampb.MethodRequest`
>   `myteampb.MyTeamMethodRequest` -> `mtpb.MyTeamMethodRequest` 或 `myteampb.MethodRequest`



#### Variable name vs. type 变量名与类型

The compiler always knows the type of a variable, and in most cases it is also clear to the reader what type a variable is by how it is used. It is only necessary to clarify the type of a variable if its value appears twice in the same scope.

​	编译器始终知道变量的类型，并且在大多数情况下，读者也可以通过变量的使用方式清楚地知道变量的类型。只有当变量的值在同一作用域中出现两次时，才需要明确变量的类型。

| Repetitive Name 重复名称      | Better Name 更好的名称 |
| ----------------------------- | ---------------------- |
| `var numUsers int`            | `var users int`        |
| `var nameString string`       | `var name string`      |
| `var primaryProject *Project` | `var primary *Project` |

If the value appears in multiple forms, this can be clarified either with an extra word like `raw` and `parsed` or with the underlying representation:

​	如果值以多种形式出现，则可以使用诸如 `raw` 和 `parsed` 之类的额外单词或使用基础表示法来阐明：

``` go
// Good:
limitStr := r.FormValue("limit")
limit, err := strconv.Atoi(limitStr)
// Good:
limitRaw := r.FormValue("limit")
limit, err := strconv.Atoi(limitRaw)
```



#### External context vs. local names 外部上下文与局部名称

Names that include information from their surrounding context often create extra noise without benefit. The package name, method name, type name, function name, import path, and even filename can all provide context that automatically qualifies all names within.

​	包含来自其周围上下文的信息的名称通常会产生额外的噪音而没有好处。包名称、方法名称、类型名称、函数名称、导入路径，甚至文件名都可以提供自动限定其中所有名称的上下文。

``` go
// Bad:
// In package "ads/targeting/revenue/reporting"
type AdsTargetingRevenueReport struct{}

func (p *Project) ProjectName() string
// Good:
// In package "ads/targeting/revenue/reporting"
type Report struct{}

func (p *Project) Name() string
// Bad:
// In package "sqldb"
type DBConnection struct{}
// Good:
// In package "sqldb"
type Connection struct{}
// Bad:
// In package "ads/targeting"
func Process(in *pb.FooProto) *Report {
    adsTargetingID := in.GetAdsTargetingID()
}
// Good:
// In package "ads/targeting"
func Process(in *pb.FooProto) *Report {
    id := in.GetAdsTargetingID()
}
```

Repetition should generally be evaluated in the context of the user of the symbol, rather than in isolation. For example, the following code has lots of names that may be fine in some circumstances, but redundant in context:

​	重复通常应在符号使用者的上下文中进行评估，而不是孤立地进行评估。例如，以下代码有很多名称，在某些情况下可能很好，但在上下文中是多余的：

``` go
// Bad:
func (db *DB) UserCount() (userCount int, err error) {
    var userCountInt64 int64
    if dbLoadError := db.LoadFromDatabase("count(distinct users)", &userCountInt64); dbLoadError != nil {
        return 0, fmt.Errorf("failed to load user count: %s", dbLoadError)
    }
    userCount = int(userCountInt64)
    return userCount, nil
}
```

Instead, information about names that are clear from context or usage can often be omitted:

​	相反，通常可以省略从上下文或用法中清楚的名称信息：

``` go
// Good:
func (db *DB) UserCount() (int, error) {
    var count int64
    if err := db.Load("count(distinct users)", &count); err != nil {
        return 0, fmt.Errorf("failed to load user count: %s", err)
    }
    return int(count), nil
}
```



## Commentary 注释

The conventions around commentary (which include what to comment, what style to use, how to provide runnable examples, etc.) are intended to support the experience of reading the documentation of a public API. See [Effective Go](http://golang.org/doc/effective_go.html#commentary) for more information.

​	有关注释的约定（包括要注释的内容、要使用的样式、如何提供可运行的示例等）旨在支持阅读公共 API 文档的体验。有关更多信息，请参阅 Effective Go。

The best practices document’s section on [documentation conventions](https://google.github.io/styleguide/go/best-practices#documentation-conventions) discusses this further.

​	最佳实践文档中有关文档约定的部分对此进行了进一步讨论。

**Best Practice:** Use [doc preview](https://google.github.io/styleguide/go/best-practices#documentation-preview) during development and code review to see whether the documentation and runnable examples are useful and are presented the way you expect them to be.

​	最佳实践：在开发和代码审查期间使用文档预览，以查看文档和可运行示例是否有用，以及是否按照您期望的方式呈现。

**Tip:** Godoc uses very little special formatting; lists and code snippets should usually be indented to avoid linewrapping. Apart from indentation, decoration should generally be avoided.

​	提示：Godoc 使用非常少的特殊格式；列表和代码段通常应缩进以避免换行。除了缩进之外，通常应避免装饰。



### Comment line length 注释行长度

Ensure that commentary is readable from source even on narrow screens.

​	确保即使在窄屏幕上也能从源代码中读取注释。

When a comment gets too long, it is recommended to wrap it into multiple single-line comments. When possible, aim for comments that will read well on an 80-column wide terminal, however this is not a hard cut-off; there is no fixed line length limit for comments in Go. The standard library, for example, often chooses to break a comment based on punctuation, which sometimes leaves the individual lines closer to the 60-70 character mark.

​	当注释过长时，建议将其包装成多行单行注释。如果可能，请尝试在 80 列宽的终端上良好显示的注释，但这并不是严格的限制；Go 中的注释没有固定的行长限制。例如，标准库通常选择根据标点符号来中断注释，这有时会使各个行的字符数接近 60-70 个。

There is plenty of existing code in which comments exceed 80 characters in length. This guidance should not be used as a justification to change such code as part of a readability review (see [consistency](https://google.github.io/styleguide/go/guide#consistency)), though teams are encouraged to opportunistically update comments to follow this guideline as a part of other refactors. The primary goal of this guideline is to ensure that all Go readability mentors make the same recommendation when and if recommendations are made.

​	有很多现有代码中的注释长度超过 80 个字符。此指南不应作为在可读性审查中更改此类代码的理由（请参阅一致性），尽管鼓励团队在其他重构中适时更新注释以遵循此指南。此指南的主要目的是确保所有 Go 可读性导师在提出建议时和提出建议后提出相同的建议。

See this [post from The Go Blog on documentation](https://blog.golang.org/godoc-documenting-go-code) for more on commentary.

​	请参阅 The Go Blog 上的这篇关于注释的博文，了解更多信息。

``` go
// Good:
// This is a comment paragraph.
// The length of individual lines doesn't matter in Godoc;
// but the choice of wrapping makes it easy to read on narrow screens.
//
// Don't worry too much about the long URL:
// https://supercalifragilisticexpialidocious.example.com:8080/Animalia/Chordata/Mammalia/Rodentia/Geomyoidea/Geomyidae/
//
// Similarly, if you have other information that is made awkward
// by too many line breaks, use your judgment and include a long line
// if it helps rather than hinders.
```

Avoid comments that will wrap repeatedly on small screens, which is a poor reader experience.

​	避免在小屏幕上反复换行的注释，因为这会带来糟糕的阅读体验。

```
# Bad:
// This is a comment paragraph. The length of individual lines doesn't matter in
Godoc;
// but the choice of wrapping causes jagged lines on narrow screens or in code
review,
// which can be annoying, especially when in a comment block that will wrap
repeatedly.
//
// Don't worry too much about the long URL:
// https://supercalifragilisticexpialidocious.example.com:8080/Animalia/Chordata/Mammalia/Rodentia/Geomyoidea/Geomyidae/
```



### Doc comments 文档注释



All top-level exported names must have doc comments, as should unexported type or function declarations with unobvious behavior or meaning. These comments should be [full sentences](https://google.github.io/styleguide/go/decisions#comment-sentences) that begin with the name of the object being described. An article (“a”, “an”, “the”) can precede the name to make it read more naturally.

​	所有顶级导出名称必须有文档注释，未导出的类型或函数声明也应如此，这些声明的行为或含义不明显。这些注释应为以所描述对象的名称开头的完整句子。一个冠词（“a”、“an”、“the”）可以放在名称前面，使其读起来更自然。

``` go
// Good:
// A Request represents a request to run a command.
type Request struct { ...

// Encode writes the JSON encoding of req to w.
func Encode(w io.Writer, req *Request) { ...
```

Doc comments appear in [Godoc](https://pkg.go.dev/) and are surfaced by IDEs, and therefore should be written for anyone using the package.

​	文档注释显示在 Godoc 中，并由 IDE 浮出水面，因此应为使用该软件包的任何人编写。

A documentation comment applies to the following symbol, or the group of fields if it appears in a struct.

​	文档注释适用于以下符号，或如果它出现在结构中，则适用于字段组。

``` go
// Good:
// Options configure the group management service.
type Options struct {
    // General setup:
    Name  string
    Group *FooGroup

    // Dependencies:
    DB *sql.DB

    // Customization:
    LargeGroupThreshold int // optional; default: 10
    MinimumMembers      int // optional; default: 2
}
```

**Best Practice:** If you have doc comments for unexported code, follow the same custom as if it were exported (namely, starting the comment with the unexported name). This makes it easy to export it later by simply replacing the unexported name with the newly-exported one across both comments and code.

​	最佳做法：如果您有未导出代码的文档注释，请按照与导出代码相同的习惯（即以未导出名称开头注释）进行操作。这使得以后只需在注释和代码中将未导出名称替换为新导出的名称即可轻松导出它。



### Comment sentences 注释句子



Comments that are complete sentences should be capitalized and punctuated like standard English sentences. (As an exception, it is okay to begin a sentence with an uncapitalized identifier name if it is otherwise clear. Such cases are probably best done only at the beginning of a paragraph.)

​	完整的句子注释应大写并标点，就像标准英语句子一样。（作为例外，如果句子以小写的标识符名称开头，并且很清楚，则可以接受。此类情况可能最好只在段落开头执行。）

Comments that are sentence fragments have no such requirements for punctuation or capitalization.

​	句子片段形式的注释对标点符号或大写字母没有此类要求。

[Documentation comments](https://google.github.io/styleguide/go/decisions#doc-comments) should always be complete sentences, and as such should always be capitalized and punctuated. Simple end-of-line comments (especially for struct fields) can be simple phrases that assume the field name is the subject.

​	文档注释应始终是完整的句子，因此应始终大写并标点符号。简单的行尾注释（尤其是对于结构字段）可以是简单的短语，假设字段名称是主题。

``` go
// Good:
// A Server handles serving quotes from the collected works of Shakespeare.
type Server struct {
    // BaseDir points to the base directory under which Shakespeare's works are stored.
    //
    // The directory structure is expected to be the following:
    //   {BaseDir}/manifest.json
    //   {BaseDir}/{name}/{name}-part{number}.txt
    BaseDir string

    WelcomeMessage  string // displayed when user logs in
    ProtocolVersion string // checked against incoming requests
    PageLength      int    // lines per page when printing (optional; default: 20)
}
```



### Examples 示例



Packages should clearly document their intended usage. Try to provide a [runnable example](http://blog.golang.org/examples); examples show up in Godoc. Runnable examples belong in the test file, not the production source file. See this example ([Godoc](https://pkg.go.dev/time#example-Duration), [source](https://cs.opensource.google/go/go/+/HEAD:src/time/example_test.go)).

​	软件包应清楚地记录其预期用法。尝试提供可运行的示例；示例显示在 Godoc 中。可运行的示例应属于测试文件，而不是生产源文件。请参阅此示例（Godoc，源代码）。

If it isn’t feasible to provide a runnable example, example code can be provided within code comments. As with other code and command-line snippets in comments, it should follow standard formatting conventions.

​	如果无法提供可运行的示例，则可以在代码注释中提供示例代码。与注释中的其他代码和命令行片段一样，它应遵循标准格式约定。



### Named result parameters 命名结果参数



When naming parameters, consider how function signatures appear in Godoc. The name of the function itself and the type of the result parameters are often sufficiently clear.

​	在命名参数时，请考虑函数签名在 Godoc 中的显示方式。函数本身的名称和结果参数的类型通常足够清楚。

``` go
// Good:
func (n *Node) Parent1() *Node
func (n *Node) Parent2() (*Node, error)
```

If a function returns two or more parameters of the same type, adding names can be useful.

​	如果函数返回两个或多个相同类型的参数，则添加名称可能会有用。

``` go
// Good:
func (n *Node) Children() (left, right *Node, err error)
```

If the caller must take action on particular result parameters, naming them can help suggest what the action is:

​	如果调用者必须对特定结果参数执行操作，则对它们进行命名可以帮助建议操作是什么：

``` go
// Good:
// WithTimeout returns a context that will be canceled no later than d duration
// from now.
//
// The caller must arrange for the returned cancel function to be called when
// the context is no longer needed to prevent a resource leak.
func WithTimeout(parent Context, d time.Duration) (ctx Context, cancel func())
```

In the code above, cancellation is a particular action a caller must take. However, were the result parameters written as `(Context, func())` alone, it would be unclear what is meant by “cancel function”.

​	在上面的代码中，取消是一个调用者必须执行的特定操作。但是，如果结果参数仅写为 `(Context, func())` ，则“取消函数”的含义就不清楚了。

Don’t use named result parameters when the names produce [unnecessary repetition](https://google.github.io/styleguide/go/decisions#repetitive-with-type).

​	当名称导致不必要的重复时，不要使用命名结果参数。

``` go
// Bad:
func (n *Node) Parent1() (node *Node)
func (n *Node) Parent2() (node *Node, err error)
```

Don’t name result parameters in order to avoid declaring a variable inside the function. This practice results in unnecessary API verbosity at the cost of minor implementation brevity.

​	不要命名结果参数，以避免在函数内部声明变量。这种做法会导致不必要的 API 冗长，而实现简洁性却很小。

[Naked returns](https://tour.golang.org/basics/7) are acceptable only in a small function. Once it’s a medium-sized function, be explicit with your returned values. Similarly, do not name result parameters just because it enables you to use naked returns. [Clarity](https://google.github.io/styleguide/go/guide#clarity) is always more important than saving a few lines in your function.

​	裸返回仅在小函数中可以接受。一旦它成为一个中等大小的函数，就要明确你的返回值。同样，不要仅仅因为这样可以让你使用裸返回就命名结果参数。清晰性始终比在函数中节省几行代码更重要。

It is always acceptable to name a result parameter if its value must be changed in a deferred closure.

​	如果必须在延迟闭包中更改结果参数的值，则始终可以为其命名。

> **Tip:** Types can often be clearer than names in function signatures. [GoTip #38: Functions as Named Types](https://google.github.io/styleguide/go/index.html#gotip) demonstrates this.
>
> ​	提示：类型通常比函数签名中的名称更清晰。GoTip #38：函数作为命名类型对此进行了演示。
>
> In, [`WithTimeout`](https://pkg.go.dev/context#WithTimeout) above, the real code uses a [`CancelFunc`](https://pkg.go.dev/context#CancelFunc) instead of a raw `func()` in the result parameter list and requires little effort to document.
>
> ​	在上面的 `WithTimeout` 中，实际代码在结果参数列表中使用 `CancelFunc` 而不是原始 `func()` ，并且只需要很少的精力来记录。



### Package comments 包注释



Package comments must appear immediately above the package clause with no blank line between the comment and the package name. Example:

​	包注释必须紧跟在包语句上方，注释与包名之间不能有空行。示例：

``` go
// Good:
// Package math provides basic constants and mathematical functions.
//
// This package does not guarantee bit-identical results across architectures.
package math
```

There must be a single package comment per package. If a package is composed of multiple files, exactly one of the files should have a package comment.

​	每个包必须有一个包注释。如果一个包由多个文件组成，则其中恰好一个文件应具有包注释。

Comments for `main` packages have a slightly different form, where the name of the `go_binary` rule in the BUILD file takes the place of the package name.

​	 `main` 包的注释形式略有不同，其中 BUILD 文件中的 `go_binary` 规则的名称取代了包名。

``` go
// Good:
// The seed_generator command is a utility that generates a Finch seed file
// from a set of JSON study configs.
package main
```

Other styles of comment are fine as long as the name of the binary is exactly as written in the BUILD file. When the binary name is the first word, capitalizing it is required even though it does not strictly match the spelling of the command-line invocation.

​	只要二进制文件的名称与 BUILD 文件中编写的名称完全一致，其他样式的注释都是可以的。当二进制文件名是第一个单词时，即使它与命令行调用的拼写不完全匹配，也需要将其大写。

``` go
// Good:
// Binary seed_generator ...
// Command seed_generator ...
// Program seed_generator ...
// The seed_generator command ...
// The seed_generator program ...
// Seed_generator ...
```

Tips:

​	提示：

- Example command-line invocations and API usage can be useful documentation. For Godoc formatting, indent the comment lines containing code.

  ​	示例命令行调用和 API 用法可以作为有用的文档。对于 Godoc 格式，缩进包含代码的注释行。

- If there is no obvious primary file or if the package comment is extraordinarily long, it is acceptable to put the doc comment in a file named `doc.go` with only the comment and the package clause.

  ​	如果没有明显的首要文件，或者包注释特别长，则可以将文档注释放在一个名为 `doc.go` 的文件中，其中仅包含注释和包语句。

- Multiline comments can be used instead of multiple single-line comments. This is primarily useful if the documentation contains sections which may be useful to copy and paste from the source file, as with sample command-lines (for binaries) and template examples.

  ​	多行注释可用于代替多个单行注释。如果文档包含可能对从源文件中复制和粘贴有用的部分，这尤其有用，例如示例命令行（对于二进制文件）和模板示例。

  ```
  // Good:
  /*
  The seed_generator command is a utility that generates a Finch seed file
  from a set of JSON study configs.
  
      seed_generator *.json | base64 > finch-seed.base64
  */
  package template
  ```

- Comments intended for maintainers and that apply to the whole file are typically placed after import declarations. These are not surfaced in Godoc and are not subject to the rules above on package comments.

  ​	适用于维护人员且适用于整个文件的注释通常放在导入声明之后。这些不会在 Godoc 中显示，也不受上述有关程序包注释的规则约束。



## Imports 导入





### Import renaming 导入重命名

Imports should only be renamed to avoid a name collision with other imports. (A corollary of this is that [good package names](https://google.github.io/styleguide/go/decisions#package-names) should not require renaming.) In the event of a name collision, prefer to rename the most local or project-specific import. Local names (aliases) for packages must follow [the guidance around package naming](https://google.github.io/styleguide/go/decisions#package-names), including the prohibition on the use of underscores and capital letters.

​	仅应重命名导入以避免与其他导入发生名称冲突。（由此推论，好的程序包名称不应需要重命名。）如果发生名称冲突，最好重命名最本地或项目特定的导入。程序包的本地名称（别名）必须遵循程序包命名指南，包括禁止使用下划线和大写字母。

Generated protocol buffer packages must be renamed to remove underscores from their names, and their aliases must have a `pb` suffix. See [proto and stub best practices](https://google.github.io/styleguide/go/best-practices#import-protos) for more information.

​	生成的协议缓冲区程序包必须重命名以从其名称中删除下划线，并且其别名必须具有 `pb` 后缀。有关更多信息，请参阅 proto 和 stub 最佳实践。

``` go
// Good:
import (
    fspb "path/to/package/foo_service_go_proto"
)
```

Imports that have package names with no useful identifying information (e.g. `package v1`) should be renamed to include the previous path component. The rename must be consistent with other local files importing the same package and may include the version number.

​	导入的包名称没有有用的标识信息（例如 `package v1` ），应重命名以包含以前的路径组件。重命名必须与导入相同包的其他本地文件一致，并且可能包括版本号。

**Note:** It is preferred to rename the package to conform with [good package names](https://google.github.io/styleguide/go/decisions#package-names), but that is often not feasible for packages in vendored directories.

​	注意：最好根据良好的包名称重命名包，但这对于供应商目录中的包通常不可行。

``` go
// Good:
import (
    core "github.com/kubernetes/api/core/v1"
    meta "github.com/kubernetes/apimachinery/pkg/apis/meta/v1beta1"
)
```

If you need to import a package whose name collides with a common local variable name that you want to use (e.g. `url`, `ssh`) and you wish to rename the package, the preferred way to do so is with the `pkg` suffix (e.g. `urlpkg`). Note that it is possible to shadow a package with a local variable; this rename is only necessary if the package still needs to be used when such a variable is in scope.

​	如果您需要导入一个包，其名称与您想要使用的常见局部变量名称冲突（例如 `url` 、 `ssh` ），并且您希望重命名该包，则首选的方法是使用 `pkg` 后缀（例如 `urlpkg` ）。请注意，可以使用局部变量隐藏包；只有当此类变量在作用域中时仍需要使用该包时，才需要进行此重命名。



### Import grouping 导入分组

Imports should be organized in two groups:

​	导入应分为两组：

- Standard library packages

  ​	标准库包

- Other (project and vendored) packages

  ​	其他（项目和供应商）包

``` go
// Good:
package main

import (
    "fmt"
    "hash/adler32"
    "os"

    "github.com/dsnet/compress/flate"
    "golang.org/x/text/encoding"
    "google.golang.org/protobuf/proto"
    foopb "myproj/foo/proto/proto"
    _ "myproj/rpc/protocols/dial"
    _ "myproj/security/auth/authhooks"
)
```

It is acceptable to split the project packages into multiple groups, for example if you want a separate group for renamed, imported-only-for-side-effects or another special group of imports.

​	可以将项目包拆分为多个组，例如，如果您想要一个单独的组用于重命名、仅导入以产生副作用或其他特殊组的导入。

``` go
// Good:
package main

import (
    "fmt"
    "hash/adler32"
    "os"


    "github.com/dsnet/compress/flate"
    "golang.org/x/text/encoding"
    "google.golang.org/protobuf/proto"

    foopb "myproj/foo/proto/proto"

    _ "myproj/rpc/protocols/dial"
    _ "myproj/security/auth/authhooks"
)
```

**Note:** Maintaining optional groups - splitting beyond what is necessary for the mandatory separation between standard library and Google imports - is not supported by the [goimports](https://google.github.io/styleguide/go/golang.org/x/tools/cmd/goimports) tool. Additional import subgroups require attention on the part of both authors and reviewers to maintain in a conforming state.

​	注意：维护可选组（拆分超出标准库和 Google 导入之间的强制分离所必需的内容）不受 goimports 工具支持。其他导入子组需要作者和审阅者共同关注，以保持符合状态。

Google programs that are also AppEngine apps should have a separate group for AppEngine imports.

​	同时也是 AppEngine 应用的 Google 程序应为 AppEngine 导入设置一个单独的组。

Gofmt takes care of sorting each group by import path. However, it does not automatically separate imports into groups. The popular [goimports](https://google.github.io/styleguide/go/golang.org/x/tools/cmd/goimports) tool combines Gofmt and import management, separating imports into groups based on the decision above. It is permissible to let [goimports](https://google.github.io/styleguide/go/golang.org/x/tools/cmd/goimports) manage import arrangement entirely, but as a file is revised its import list must remain internally consistent.

​	Gofmt 会按导入路径对每个组进行排序。但是，它不会自动将导入内容分成组。流行的 goimports 工具结合了 Gofmt 和导入管理，根据上述决策将导入内容分成组。允许 goimports 完全管理导入安排，但随着文件经过修订，其导入列表必须保持内部一致。



### Import “blank” (`import _`) 导入“空白”（ `import _` ）



Packages that are imported only for their side effects (using the syntax `import _ "package"`) may only be imported in a main package, or in tests that require them.

​	仅因其副作用而导入的软件包（使用语法 `import _ "package"` ）只能在主软件包中导入，或在需要它们的测试中导入。

Some examples of such packages include:

​	此类软件包的一些示例包括：

- [time/tzdata](https://pkg.go.dev/time/tzdata)

- [image/jpeg](https://pkg.go.dev/image/jpeg) in image processing code

  ​	图像处理代码中的 image/jpeg

Avoid blank imports in library packages, even if the library indirectly depends on them. Constraining side-effect imports to the main package helps control dependencies, and makes it possible to write tests that rely on a different import without conflict or wasted build costs.

​	避免在库包中进行空白导入，即使库间接依赖于它们。将副作用导入限制在主包中有助于控制依赖关系，并使编写依赖于不同导入的测试成为可能，而不会发生冲突或浪费构建成本。

The following are the only exceptions to this rule:

​	以下情况是此规则的唯一例外：

- You may use a blank import to bypass the check for disallowed imports in the [nogo static checker](https://github.com/bazelbuild/rules_go/blob/master/go/nogo.rst).

  ​	您可能使用空白导入来绕过 nogo 静态检查器中不允许导入的检查。

- You may use a blank import of the [embed](https://pkg.go.dev/embed) package in a source file which uses the `//go:embed` compiler directive.

  ​	您可能在使用 `//go:embed` 编译器指令的源文件中使用 embed 包的空白导入。

**Tip:** If you create a library package that indirectly depends on a side-effect import in production, document the intended usage.

​	提示：如果您创建间接依赖于生产中副作用导入的库包，请记录预期用法。



### Import “dot” (`import .`) 导入“点”（ `import .` ）



The `import .` form is a language feature that allows bringing identifiers exported from another package to the current package without qualification. See the [language spec](https://go.dev/ref/spec#Import_declarations) for more.

​	 `import .` 形式是一种语言特性，它允许将从另一个包导出的标识符引入当前包，而无需限定。有关更多信息，请参阅语言规范。

Do **not** use this feature in the Google codebase; it makes it harder to tell where the functionality is coming from.

​	请勿在 Google 代码库中使用此功能；这会让人更难知道功能来自哪里。

``` go
// Bad:
package foo_test

import (
    "bar/testutil" // also imports "foo"
    . "foo"
)

var myThing = Bar() // Bar defined in package foo; no qualification needed.
// Good:
package foo_test

import (
    "bar/testutil" // also imports "foo"
    "foo"
)

var myThing = foo.Bar()
```



## Errors 错误



### Returning errors 返回错误



Use `error` to signal that a function can fail. By convention, `error` is the last result parameter.

​	使用 `error` 来表示函数可能失败。根据惯例， `error` 是最后一个结果参数。

``` go
// Good:
func Good() error { /* ... */ }
```

Returning a `nil` error is the idiomatic way to signal a successful operation that could otherwise fail. If a function returns an error, callers must treat all non-error return values as unspecified unless explicitly documented otherwise. Commonly, the non-error return values are their zero values, but this cannot be assumed.

​	返回 `nil` 错误是表示可能失败的成功操作的惯用方式。如果函数返回错误，调用者必须将所有非错误返回值视为未指定，除非另有明确说明。通常，非错误返回值是其零值，但不能假设这一点。

``` go
// Good:
func GoodLookup() (*Result, error) {
    // ...
    if err != nil {
        return nil, err
    }
    return res, nil
}
```

Exported functions that return errors should return them using the `error` type. Concrete error types are susceptible to subtle bugs: a concrete `nil` pointer can get wrapped into an interface and thus become a non-nil value (see the [Go FAQ entry on the topic](https://golang.org/doc/faq#nil_error)).

​	返回错误的导出函数应使用 `error` 类型返回错误。具体的错误类型容易出现细微的错误：一个具体的 `nil` 指针可以被包装成一个接口，从而成为一个非零值（请参阅 Go FAQ 条目了解该主题）。

``` go
// Bad:
func Bad() *os.PathError { /*...*/ }
```

**Tip**: A function that takes a `context.Context` argument should usually return an `error` so that the caller can determine if the context was cancelled while the function was running.

​	提示：通常，接受 `context.Context` 参数的函数应返回 `error` ，以便调用者可以确定在函数运行时是否取消了上下文。



### Error strings 错误字符串



Error strings should not be capitalized (unless beginning with an exported name, a proper noun or an acronym) and should not end with punctuation. This is because error strings usually appear within other context before being printed to the user.

​	错误字符串不应大写（除非以导出名称、专有名词或首字母缩写词开头），也不应以标点符号结尾。这是因为错误字符串通常出现在其他上下文中，然后才打印给用户。

``` go
// Bad:
err := fmt.Errorf("Something bad happened.")
// Good:
err := fmt.Errorf("something bad happened")
```

On the other hand, the style for the full displayed message (logging, test failure, API response, or other UI) depends, but should typically be capitalized.

​	另一方面，完整显示消息的样式（日志记录、测试失败、API 响应或其他 UI）取决于具体情况，但通常应大写。

``` go
// Good:
log.Infof("Operation aborted: %v", err)
log.Errorf("Operation aborted: %v", err)
t.Errorf("Op(%q) failed unexpectedly; err=%v", args, err)
```



### Handle errors 处理错误



Code that encounters an error should make a deliberate choice about how to handle it. It is not usually appropriate to discard errors using `_` variables. If a function returns an error, do one of the following:

​	遇到错误的代码应慎重选择处理方式。通常不适合使用 `_` 变量来丢弃错误。如果函数返回错误，请执行以下操作之一：

- Handle and address the error immediately.
  立即处理并解决错误。
- Return the error to the caller.
  将错误返回给调用方。
- In exceptional situations, call [`log.Fatal`](https://pkg.go.dev/github.com/golang/glog#Fatal) or (if absolutely necessary) `panic`.
  在特殊情况下，调用 `log.Fatal` 或（如果绝对必要） `panic` 。

**Note:** `log.Fatalf` is not the standard library log. See [#logging].

​	注意： `log.Fatalf` 不是标准库日志。请参阅 [#logging]。

In the rare circumstance where it is appropriate to ignore or discard an error (for example a call to [`(*bytes.Buffer).Write`](https://pkg.go.dev/bytes#Buffer.Write) that is documented to never fail), an accompanying comment should explain why this is safe.

​	在极少数情况下，忽略或丢弃错误是合适的（例如，调用 `(*bytes.Buffer).Write` ，据记录它绝不会失败），应附带注释说明这样做是安全的。

``` go
// Good:
var b *bytes.Buffer

n, _ := b.Write(p) // never returns a non-nil error
```

For more discussion and examples of error handling, see [Effective Go](http://golang.org/doc/effective_go.html#errors) and [best practices](https://google.github.io/styleguide/go/best-practices.html#error-handling).

​	有关错误处理的更多讨论和示例，请参阅 Effective Go 和最佳做法。



### In-band errors 带内错误



In C and similar languages, it is common for functions to return values like -1, null, or the empty string to signal errors or missing results. This is known as in-band error handling.

​	在 C 和类似语言中，函数通常会返回 -1、null 或空字符串等值来表示错误或结果缺失。这称为带内错误处理。

``` go
// Bad:
// Lookup returns the value for key or -1 if there is no mapping for key.
func Lookup(key string) int
```

Failing to check for an in-band error value can lead to bugs and can attribute errors to the wrong function.

​	未能检查带内错误值可能会导致错误，并可能将错误归因于错误的函数。

``` go
// Bad:
// The following line returns an error that Parse failed for the input value,
// whereas the failure was that there is no mapping for missingKey.
return Parse(Lookup(missingKey))
```

Go’s support for multiple return values provides a better solution (see the [Effective Go section on multiple returns](http://golang.org/doc/effective_go.html#multiple-returns)). Instead of requiring clients to check for an in-band error value, a function should return an additional value to indicate whether its other return values are valid. This return value may be an error or a boolean when no explanation is needed, and should be the final return value.

​	Go 对多返回值的支持提供了一个更好的解决方案（请参阅有关多返回值的 Go 实用指南部分）。函数应返回一个附加值来指示其其他返回值是否有效，而不要求客户端检查带内错误值。此返回值在不需要解释时可以是错误或布尔值，并且应该是最终返回值。

``` go
// Good:
// Lookup returns the value for key or ok=false if there is no mapping for key.
func Lookup(key string) (value string, ok bool)
```

This API prevents the caller from incorrectly writing `Parse(Lookup(key))` which causes a compile-time error, since `Lookup(key)` has 2 outputs.

​	此 API 可防止调用者错误地编写 `Parse(Lookup(key))` ，因为 `Lookup(key)` 有 2 个输出，这会导致编译时错误。

Returning errors in this way encourages more robust and explicit error handling:

​	以这种方式返回错误有助于实现更健壮且显式的错误处理：

``` go
// Good:
value, ok := Lookup(key)
if !ok {
    return fmt.Errorf("no value for %q", key)
}
return Parse(value)
```

Some standard library functions, like those in package `strings`, return in-band error values. This greatly simplifies string-manipulation code at the cost of requiring more diligence from the programmer. In general, Go code in the Google codebase should return additional values for errors.

​	一些标准库函数（如包 `strings` 中的函数）返回带内错误值。这极大地简化了字符串操作代码，但需要程序员更加勤勉。通常，Google 代码库中的 Go 代码应为错误返回附加值。



### Indent error flow 缩进错误流



Handle errors before proceeding with the rest of your code. This improves the readability of the code by enabling the reader to find the normal path quickly. This same logic applies to any block which tests a condition then ends in a terminal condition (e.g., `return`, `panic`, `log.Fatal`).

​	在继续执行代码的其余部分之前处理错误。这通过使读者能够快速找到正常路径来提高代码的可读性。此逻辑同样适用于测试条件然后以终端条件结束的任何块（例如， `return` 、 `panic` 、 `log.Fatal` ）。

Code that runs if the terminal condition is not met should appear after the `if` block, and should not be indented in an `else` clause.

​	如果未满足终端条件，则运行的代码应出现在 `if` 块之后，并且不应缩进 `else` 子句。

``` go
// Good:
if err != nil {
    // error handling
    return // or continue, etc.
}
// normal code
// Bad:
if err != nil {
    // error handling
} else {
    // normal code that looks abnormal due to indentation
}
```

> **Tip:** If you are using a variable for more than a few lines of code, it is generally not worth using the `if`-with-initializer style. In these cases, it is usually better to move the declaration out and use a standard `if` statement:
>
> ​	提示：如果您将变量用于多行代码，通常不值得使用 `if` -with-initializer样式。在这些情况下，通常最好将声明移出并使用标准 `if` 语句：
>
> ```
> // Good:
> x, err := f()
> if err != nil {
>   // error handling
>   return
> }
> // lots of code that uses x
> // across multiple lines
> // Bad:
> if x, err := f(); err != nil {
>   // error handling
>   return
> } else {
>   // lots of code that uses x
>   // across multiple lines
> }
> ```

See [Go Tip #1: Line of Sight](https://google.github.io/styleguide/go/index.html#gotip) and [TotT: Reduce Code Complexity by Reducing Nesting](https://testing.googleblog.com/2017/06/code-health-reduce-nesting-reduce.html) for more details.

​	有关更多详细信息，请参阅 Go 提示 #1：视线和 TotT：通过减少嵌套来降低代码复杂性。



## Language 语言



### Literal formatting 文字格式化

Go has an exceptionally powerful [composite literal syntax](https://golang.org/ref/spec#Composite_literals), with which it is possible to express deeply-nested, complicated values in a single expression. Where possible, this literal syntax should be used instead of building values field-by-field. The `gofmt` formatting for literals is generally quite good, but there are some additional rules for keeping these literals readable and maintainable.

​	Go 具有非常强大的复合字面量语法，可以使用它在单个表达式中表示深度嵌套的复杂值。在可能的情况下，应使用此字面量语法，而不是逐字段构建值。字面量的 `gofmt` 格式通常非常好，但还有一些其他规则可以保持这些字面量可读且可维护。



#### Field names 字段名称

Struct literals should usually specify **field names** for types defined outside the current package.

​	结构字面量通常应为当前包外部定义的类型指定字段名称。

- Include field names for types from other packages.

  ​	包含来自其他包的类型的字段名称。

  ```
  // Good:
  good := otherpkg.Type{A: 42}
  ```

  The position of fields in a struct and the full set of fields (both of which are necessary to get right when field names are omitted) are not usually considered to be part of a struct’s public API; specifying the field name is needed to avoid unnecessary coupling.

  ​	结构中的字段位置和字段的完整集合（在省略字段名称时，两者都是正确的）通常不被视为结构的公共 API 的一部分；需要指定字段名称以避免不必要的耦合。

  ```
  // Bad:
  // https://pkg.go.dev/encoding/csv#Reader
  r := csv.Reader{',', '#', 4, false, false, false, false}
  ```

  Field names may be omitted within small, simple structs whose composition and order are documented as being stable.

  ​	可以在其组成和顺序被记录为稳定的简单小结构中省略字段名称。

  ```
  // Good:
  okay := image.Point{42, 54}
  also := image.Point{X: 42, Y: 54}
  ```

- For package-local types, field names are optional.

  ​	对于包本地类型，字段名称是可选的。

  ```
  // Good:
  okay := Type{42}
  also := internalType{4, 2}
  ```

  Field names should still be used if it makes the code clearer, and it is very common to do so. For example, a struct with a large number of fields should almost always be initialized with field names.

  ​	如果它使代码更清晰，仍然应该使用字段名称，这样做非常普遍。例如，具有大量字段的结构几乎总是应该使用字段名称进行初始化。

  ```
  // Good:
  okay := StructWithLotsOfFields{
    field1: 1,
    field2: "two",
    field3: 3.14,
    field4: true,
  }
  ```



#### Matching braces 匹配大括号

The closing half of a brace pair should always appear on a line with the same amount of indentation as the opening brace. One-line literals necessarily have this property. When the literal spans multiple lines, maintaining this property keeps the brace matching for literals the same as brace matching for common Go syntactic constructs like functions and `if` statements.

​	大括号对的闭合部分应始终出现在与开括号具有相同缩进量的行上。单行文字必然具有此属性。当文字跨越多行时，保持此属性可使文字的大括号匹配与函数和 `if` 语句等常见 Go 语法结构的大括号匹配保持一致。

The most common mistake in this area is putting the closing brace on the same line as a value in a multi-line struct literal. In these cases, the line should end with a comma and the closing brace should appear on the next line.

​	此领域的常见错误是将闭合大括号放在多行结构文字中的值所在的行上。在这些情况下，该行应以逗号结尾，闭合大括号应出现在下一行。

``` go
// Good:
good := []*Type{{Key: "value"}}
// Good:
good := []*Type{
    {Key: "multi"},
    {Key: "line"},
}
// Bad:
bad := []*Type{
    {Key: "multi"},
    {Key: "line"}}
// Bad:
bad := []*Type{
    {
        Key: "value"},
}
```



#### Cuddled braces 紧缩大括号

Dropping whitespace between braces (aka “cuddling” them) for slice and array literals is only permitted when both of the following are true.

​	仅当满足以下两个条件时，才允许省略大括号之间的空格（又称“紧缩”大括号），用于切片和数组文字。

- The [indentation matches](https://google.github.io/styleguide/go/decisions#literal-matching-braces)
  缩进匹配
- The inner values are also literals or proto builders (i.e. not a variable or other expression)
  内部值也是文字或 proto 构建器（即不是变量或其他表达式）

``` go
// Good:
good := []*Type{
    { // Not cuddled
        Field: "value",
    },
    {
        Field: "value",
    },
}
// Good:
good := []*Type{{ // Cuddled correctly
    Field: "value",
}, {
    Field: "value",
}}
// Good:
good := []*Type{
    first, // Can't be cuddled
    {Field: "second"},
}
// Good:
okay := []*pb.Type{pb.Type_builder{
    Field: "first", // Proto Builders may be cuddled to save vertical space
}.Build(), pb.Type_builder{
    Field: "second",
}.Build()}
// Bad:
bad := []*Type{
    first,
    {
        Field: "second",
    }}
```



#### Repeated type names 重复的类型名称

Repeated type names may be omitted from slice and map literals. This can be helpful in reducing clutter. A reasonable occasion for repeating the type names explicitly is when dealing with a complex type that is not common in your project, when the repetitive type names are on lines that are far apart and can remind the reader of the context.

​	重复的类型名称可以从切片和映射字面量中省略。这有助于减少混乱。显式重复类型名称的合理时机是在处理项目中不常见的复杂类型时，当重复的类型名称位于相距较远的行中时，可以提醒读者上下文。

``` go
// Good:
good := []*Type{
    {A: 42},
    {A: 43},
}
// Bad:
repetitive := []*Type{
    &Type{A: 42},
    &Type{A: 43},
}
// Good:
good := map[Type1]*Type2{
    {A: 1}: {B: 2},
    {A: 3}: {B: 4},
}
// Bad:
repetitive := map[Type1]*Type2{
    Type1{A: 1}: &Type2{B: 2},
    Type1{A: 3}: &Type2{B: 4},
}
```

**Tip:** If you want to remove repetitive type names in struct literals, you can run `gofmt -s`.

​	提示：如果您想在结构字面量中删除重复的类型名称，可以运行 `gofmt -s` 。



#### Zero-value fields 零值字段

[Zero-value](https://golang.org/ref/spec#The_zero_value) fields may be omitted from struct literals when clarity is not lost as a result.

​	当不会因此而失去清晰度时，零值字段可以从结构字面量中省略。

Well-designed APIs often employ zero-value construction for enhanced readability. For example, omitting the three zero-value fields from the following struct draws attention to the only option that is being specified.

​	设计良好的 API 通常采用零值构造来提高可读性。例如，从以下结构中省略三个零值字段会将注意力吸引到正在指定的唯一选项上。

``` go
// Bad:
import (
  "github.com/golang/leveldb"
  "github.com/golang/leveldb/db"
)

ldb := leveldb.Open("/my/table", &db.Options{
    BlockSize: 1<<16,
    ErrorIfDBExists: true,

    // These fields all have their zero values.
    BlockRestartInterval: 0,
    Comparer: nil,
    Compression: nil,
    FileSystem: nil,
    FilterPolicy: nil,
    MaxOpenFiles: 0,
    WriteBufferSize: 0,
    VerifyChecksums: false,
})
// Good:
import (
  "github.com/golang/leveldb"
  "github.com/golang/leveldb/db"
)

ldb := leveldb.Open("/my/table", &db.Options{
    BlockSize: 1<<16,
    ErrorIfDBExists: true,
})
```

Structs within table-driven tests often benefit from [explicit field names](https://google.github.io/styleguide/go/decisions#literal-field-names), especially when the test struct is not trivial. This allows the author to omit the zero-valued fields entirely when the fields in question are not related to the test case. For example, successful test cases should omit any error-related or failure-related fields. In cases where the zero value is necessary to understand the test case, such as testing for zero or `nil` inputs, the field names should be specified.

​	表驱动测试中的结构通常受益于显式字段名，尤其是在测试结构不是微不足道的情况下。这允许作者在字段与测试用例无关时完全省略零值字段。例如，成功的测试用例应省略任何与错误相关或与失败相关的字段。在零值对于理解测试用例是必要的情况下，例如测试零或 `nil` 输入，应指定字段名。

**Concise
简洁**

```
tests := []struct {
    input      string
    wantPieces []string
    wantErr    error
}{
    {
        input:      "1.2.3.4",
        wantPieces: []string{"1", "2", "3", "4"},
    },
    {
        input:   "hostname",
        wantErr: ErrBadHostname,
    },
}
```

**Explicit
显式**

```
tests := []struct {
    input    string
    wantIPv4 bool
    wantIPv6 bool
    wantErr  bool
}{
    {
        input:    "1.2.3.4",
        wantIPv4: true,
        wantIPv6: false,
    },
    {
        input:    "1:2::3:4",
        wantIPv4: false,
        wantIPv6: true,
    },
    {
        input:    "hostname",
        wantIPv4: false,
        wantIPv6: false,
        wantErr:  true,
    },
}
```



### Nil slices 空切片

For most purposes, there is no functional difference between `nil` and the empty slice. Built-in functions like `len` and `cap` behave as expected on `nil` slices.

​	对于大多数目的， `nil` 和空切片之间没有功能差异。内置函数（如 `len` 和 `cap` ）在 `nil` 切片上按预期行为。

``` go
// Good:
import "fmt"

var s []int         // nil

fmt.Println(s)      // []
fmt.Println(len(s)) // 0
fmt.Println(cap(s)) // 0
for range s {...}   // no-op

s = append(s, 42)
fmt.Println(s)      // [42]
```

If you declare an empty slice as a local variable (especially if it can be the source of a return value), prefer the nil initialization to reduce the risk of bugs by callers.

​	如果您将空切片声明为局部变量（尤其是它可以是返回值的来源时），请优先使用 nil 初始化，以降低调用者出现错误的风险。

``` go
// Good:
var t []string
// Bad:
t := []string{}
```

Do not create APIs that force their clients to make distinctions between nil and the empty slice.

​	不要创建强制其客户端区分 nil 和空切片的 API。

``` go
// Good:
// Ping pings its targets.
// Returns hosts that successfully responded.
func Ping(hosts []string) ([]string, error) { ... }
// Bad:
// Ping pings its targets and returns a list of hosts
// that successfully responded. Can be empty if the input was empty.
// nil signifies that a system error occurred.
func Ping(hosts []string) []string { ... }
```

When designing interfaces, avoid making a distinction between a `nil` slice and a non-`nil`, zero-length slice, as this can lead to subtle programming errors. This is typically accomplished by using `len` to check for emptiness, rather than `== nil`.

​	在设计接口时，避免区分 `nil` 切片和非 `nil` 、零长度切片，因为这会导致微妙的编程错误。通常通过使用 `len` 来检查是否为空，而不是 `== nil` 来实现这一点。

This implementation accepts both `nil` and zero-length slices as “empty”:

​	此实现将 `nil` 和零长度切片都接受为“空”：

``` go
// Good:
// describeInts describes s with the given prefix, unless s is empty.
func describeInts(prefix string, s []int) {
    if len(s) == 0 {
        return
    }
    fmt.Println(prefix, s)
}
```

Instead of relying on the distinction as a part of the API:

​	而不是依赖于 API 中的区分：

``` go
// Bad:
func maybeInts() []int { /* ... */ }

// describeInts describes s with the given prefix; pass nil to skip completely.
func describeInts(prefix string, s []int) {
  // The behavior of this function unintentionally changes depending on what
  // maybeInts() returns in 'empty' cases (nil or []int{}).
  if s == nil {
    return
  }
  fmt.Println(prefix, s)
}

describeInts("Here are some ints:", maybeInts())
```

See [in-band errors](https://google.github.io/styleguide/go/decisions#in-band-errors) for further discussion.

​	有关进一步讨论，请参阅带内错误。



### Indentation confusion 缩进混淆

Avoid introducing a line break if it would align the rest of the line with an indented code block. If this is unavoidable, leave a space to separate the code in the block from the wrapped line.

​	如果换行会使其余行与缩进代码块对齐，请避免换行。如果无法避免，请留一个空格将块中的代码与换行隔开。

``` go
// Bad:
if longCondition1 && longCondition2 &&
    // Conditions 3 and 4 have the same indentation as the code within the if.
    longCondition3 && longCondition4 {
    log.Info("all conditions met")
}
```

See the following sections for specific guidelines and examples:

​	有关具体指南和示例，请参阅以下部分：

- [Function formatting
  函数格式化](https://google.github.io/styleguide/go/decisions#func-formatting)
- [Conditionals and loops
  条件和循环](https://google.github.io/styleguide/go/decisions#conditional-formatting)
- [Literal formatting
  文字格式化](https://google.github.io/styleguide/go/decisions#literal-formatting)



### Function formatting 函数格式化

The signature of a function or method declaration should remain on a single line to avoid [indentation confusion](https://google.github.io/styleguide/go/decisions#indentation-confusion).

​	函数或方法声明的签名应保留在单行上，以避免缩进混淆。

Function argument lists can make some of the longest lines in a Go source file. However, they precede a change in indentation, and therefore it is difficult to break the line in a way that does not make subsequent lines look like part of the function body in a confusing way:

​	函数参数列表可能会使 Go 源文件中的一些最长行。但是，它们在缩进发生变化之前，因此很难以一种不会使后续行看起来像函数体一部分的方式来换行，这会造成混淆：

``` go
// Bad:
func (r *SomeType) SomeLongFunctionName(foo1, foo2, foo3 string,
    foo4, foo5, foo6 int) {
    foo7 := bar(foo1)
    // ...
}
```

See [best practices](https://google.github.io/styleguide/go/best-practices#funcargs) for a few options for shortening the call sites of functions that would otherwise have many arguments.

​	请参阅最佳实践，了解一些缩短函数调用站点的选项，这些函数调用站点在其他情况下会有很多参数。

``` go
// Good:
good := foo.Call(long, CallOptions{
    Names:   list,
    Of:      of,
    The:     parameters,
    Func:    all,
    Args:    on,
    Now:     separate,
    Visible: lines,
})
// Bad:
bad := foo.Call(
    long,
    list,
    of,
    parameters,
    all,
    on,
    separate,
    lines,
)
```

Lines can often be shortened by factoring out local variables.

​	通常可以通过分解局部变量来缩短行。

``` go
// Good:
local := helper(some, parameters, here)
good := foo.Call(list, of, parameters, local)
```

Similarly, function and method calls should not be separated based solely on line length.

​	同样，函数和方法调用不应仅基于行长度而分开。

``` go
// Good:
good := foo.Call(long, list, of, parameters, all, on, one, line)
// Bad:
bad := foo.Call(long, list, of, parameters,
    with, arbitrary, line, breaks)
```

Do not add comments to specific function parameters. Instead, use an [option struct](https://google.github.io/styleguide/go/best-practices#option-structure) or add more detail to the function documentation.

​	不要向特定函数参数添加注释。相反，请使用选项结构或向函数文档添加更多详细信息。

``` go
// Good:
good := server.New(ctx, server.Options{Port: 42})
// Bad:
bad := server.New(
    ctx,
    42, // Port
)
```

If call-sites are uncomfortably long, consider refactoring:

​	如果调用站点太长，请考虑重构：

``` go
// Good:
// Sometimes variadic arguments can be factored out
replacements := []string{
    "from", "to", // related values can be formatted adjacent to one another
    "source", "dest",
    "original", "new",
}

// Use the replacement struct as inputs to NewReplacer.
replacer := strings.NewReplacer(replacements...)
```

If the API cannot be changed or if the local call is unusual (whether or not the call is too long), it is always permissible to add line breaks if it aids in understanding the call.

​	如果无法更改 API 或本地调用不寻常（无论调用是否太长），如果它有助于理解调用，则始终允许添加换行符。

``` go
// Good:
canvas.RenderCube(cube,
    x0, y0, z0,
    x0, y0, z1,
    x0, y1, z0,
    x0, y1, z1,
    x1, y0, z0,
    x1, y0, z1,
    x1, y1, z0,
    x1, y1, z1,
)
```

Note that the lines in the above example are not wrapped at a specific column boundary but are grouped based on co-ordinate triples.

​	请注意，上述示例中的行不是在特定列边界处换行的，而是根据坐标三元组进行分组的。

Long string literals within functions should not be broken for the sake of line length. For functions that include such strings, a line break can be added after the string format, and the arguments can be provided on the next or subsequent lines. The decision about where the line breaks should go is best made based on semantic groupings of inputs, rather than based purely on line length.

​	函数中的长字符串文字不应为了行长而中断。对于包含此类字符串的函数，可以在字符串格式后添加一个换行符，并在下一行或后续行中提供参数。关于换行符应放在何处，最好根据输入的语义分组来做出决定，而不是纯粹基于行长。

``` go
// Good:
log.Warningf("Database key (%q, %d, %q) incompatible in transaction started by (%q, %d, %q)",
    currentCustomer, currentOffset, currentKey,
    txCustomer, txOffset, txKey)
// Bad:
log.Warningf("Database key (%q, %d, %q) incompatible in"+
    " transaction started by (%q, %d, %q)",
    currentCustomer, currentOffset, currentKey, txCustomer,
    txOffset, txKey)
```



### Conditionals and loops 条件和循环

An `if` statement should not be line broken; multi-line `if` clauses can lead to [indentation confusion](https://google.github.io/styleguide/go/decisions#indentation-confusion).

​	 `if` 语句不应换行；多行 `if` 子句可能导致缩进混乱。

``` go
// Bad:
// The second if statement is aligned with the code within the if block, causing
// indentation confusion.
if db.CurrentStatusIs(db.InTransaction) &&
    db.ValuesEqual(db.TransactionKey(), row.Key()) {
    return db.Errorf(db.TransactionError, "query failed: row (%v): key does not match transaction key", row)
}
```

If the short-circuit behavior is not required, the boolean operands can be extracted directly:

​	如果不需要短路行为，则可以直接提取布尔运算符：

``` go
// Good:
inTransaction := db.CurrentStatusIs(db.InTransaction)
keysMatch := db.ValuesEqual(db.TransactionKey(), row.Key())
if inTransaction && keysMatch {
    return db.Error(db.TransactionError, "query failed: row (%v): key does not match transaction key", row)
}
```

There may also be other locals that can be extracted, especially if the conditional is already repetitive:

​	也可能还有其他可以提取的局部变量，尤其是在条件语句已经重复的情况下：

``` go
// Good:
uid := user.GetUniqueUserID()
if db.UserIsAdmin(uid) || db.UserHasPermission(uid, perms.ViewServerConfig) || db.UserHasPermission(uid, perms.CreateGroup) {
    // ...
}
// Bad:
if db.UserIsAdmin(user.GetUniqueUserID()) || db.UserHasPermission(user.GetUniqueUserID(), perms.ViewServerConfig) || db.UserHasPermission(user.GetUniqueUserID(), perms.CreateGroup) {
    // ...
}
```

`if` statements that contain closures or multi-line struct literals should ensure that the [braces match](https://google.github.io/styleguide/go/decisions#literal-matching-braces) to avoid [indentation confusion](https://google.github.io/styleguide/go/decisions#indentation-confusion).

​	包含闭包或多行结构文字的 `if` 语句应确保大括号匹配，以避免缩进混乱。

``` go
// Good:
if err := db.RunInTransaction(func(tx *db.TX) error {
    return tx.Execute(userUpdate, x, y, z)
}); err != nil {
    return fmt.Errorf("user update failed: %s", err)
}
// Good:
if _, err := client.Update(ctx, &upb.UserUpdateRequest{
    ID:   userID,
    User: user,
}); err != nil {
    return fmt.Errorf("user update failed: %s", err)
}
```

Similarly, don’t try inserting artificial linebreaks into `for` statements. You can always let the line simply be long if there is no elegant way to refactor it:

​	同样，不要尝试在 `for` 语句中插入人工换行符。如果无法优雅地重构它，您可以始终让该行保持较长：

``` go
// Good:
for i, max := 0, collection.Size(); i < max && !collection.HasPendingWriters(); i++ {
    // ...
}
```

Often, though, there is:

​	不过，通常情况下，还是有办法的：

``` go
// Good:
for i, max := 0, collection.Size(); i < max; i++ {
    if collection.HasPendingWriters() {
        break
    }
    // ...
}
```

`switch` and `case` statements should also remain on a single line.

​	 `switch` 和 `case` 语句也应保留在单行上。

``` go
// Good:
switch good := db.TransactionStatus(); good {
case db.TransactionStarting, db.TransactionActive, db.TransactionWaiting:
    // ...
case db.TransactionCommitted, db.NoTransaction:
    // ...
default:
    // ...
}
// Bad:
switch bad := db.TransactionStatus(); bad {
case db.TransactionStarting,
    db.TransactionActive,
    db.TransactionWaiting:
    // ...
case db.TransactionCommitted,
    db.NoTransaction:
    // ...
default:
    // ...
}
```

If the line is excessively long, indent all cases and separate them with a blank line to avoid [indentation confusion](https://google.github.io/styleguide/go/decisions#indentation-confusion):

​	如果行过长，请缩进所有情况并用空行分隔，以避免缩进混淆：

``` go
// Good:
switch db.TransactionStatus() {
case
    db.TransactionStarting,
    db.TransactionActive,
    db.TransactionWaiting,
    db.TransactionCommitted:

    // ...
case db.NoTransaction:
    // ...
default:
    // ...
}
```

In conditionals comparing a variable to a constant, place the variable value on the left hand side of the equality operator:

​	在将变量与常量进行比较的条件语句中，将变量值放在等式运算符的左侧：

``` go
// Good:
if result == "foo" {
  // ...
}
```

Instead of the less clear phrasing where the constant comes first ([“Yoda style conditionals”](https://en.wikipedia.org/wiki/Yoda_conditions)):

​	而不是使用常量在前（“Yoda 风格条件语句”）这种不太清晰的表述方式：

``` go
// Bad:
if "foo" == result {
  // ...
}
```



### Copying 复制



To avoid unexpected aliasing and similar bugs, be careful when copying a struct from another package. For example, synchronization objects such as `sync.Mutex` must not be copied.

​	为了避免意外别名和类似的错误，在从另一个包复制结构时要小心。例如，诸如 `sync.Mutex` 的同步对象不得被复制。

The `bytes.Buffer` type contains a `[]byte` slice and, as an optimization for small strings, a small byte array to which the slice may refer. If you copy a `Buffer`, the slice in the copy may alias the array in the original, causing subsequent method calls to have surprising effects.

​	 `bytes.Buffer` 类型包含一个 `[]byte` 切片，并且作为对小字符串的优化，包含一个切片可能引用的字节数组。如果您复制一个 `Buffer` ，则副本中的切片可能会别名原始数组，从而导致后续方法调用产生令人惊讶的效果。

In general, do not copy a value of type `T` if its methods are associated with the pointer type, `*T`.

​	通常，如果 `T` 类型的字段与指针类型 `*T` 相关联，则不要复制该类型的变量。

``` go
// Bad:
b1 := bytes.Buffer{}
b2 := b1
```

Invoking a method that takes a value receiver can hide the copy. When you author an API, you should generally take and return pointer types if your structs contain fields that should not be copied.

​	调用采用值接收器的方法可以隐藏副本。当您编写 API 时，如果您的结构包含不应复制的字段，则通常应采用并返回指针类型。

These are acceptable:

​	这些是可以接受的：

``` go
// Good:
type Record struct {
  buf bytes.Buffer
  // other fields omitted
}

func New() *Record {...}

func (r *Record) Process(...) {...}

func Consumer(r *Record) {...}
```

But these are usually wrong:

​	但这些通常是错误的：

``` go
// Bad:
type Record struct {
  buf bytes.Buffer
  // other fields omitted
}


func (r Record) Process(...) {...} // Makes a copy of r.buf

func Consumer(r Record) {...} // Makes a copy of r.buf
```

This guidance also applies to copying `sync.Mutex`.

​	此指南也适用于复制 `sync.Mutex` 。



### Don’t panic 不要惊慌



Do not use `panic` for normal error handling. Instead, use `error` and multiple return values. See the [Effective Go section on errors](http://golang.org/doc/effective_go.html#errors).

​	不要将 `panic` 用于常规错误处理。相反，请使用 `error` 和多个返回值。请参阅有关错误的有效 Go 部分。

Within `package main` and initialization code, consider [`log.Exit`](https://pkg.go.dev/github.com/golang/glog#Exit) for errors that should terminate the program (e.g., invalid configuration), as in many of these cases a stack trace will not help the reader. Please note that [`log.Exit`](https://pkg.go.dev/github.com/golang/glog#Exit) calls [`os.Exit`](https://pkg.go.dev/os#Exit) and any deferred functions will not be run.

​	在 `package main` 和初始化代码中，考虑将 `log.Exit` 用于应终止程序的错误（例如，无效配置），因为在许多此类情况下，堆栈跟踪对读者没有帮助。请注意， `log.Exit` 调用 `os.Exit` ，并且不会运行任何延迟函数。

For errors that indicate “impossible” conditions, namely bugs that should always be caught during code review and/or testing, a function may reasonably return an error or call [`log.Fatal`](https://pkg.go.dev/github.com/golang/glog#Fatal).

​	对于指示“不可能”条件的错误，即在代码审查和/或测试期间应始终捕获的错误，函数可以合理地返回错误或调用 `log.Fatal` 。

**Note:** `log.Fatalf` is not the standard library log. See [#logging].

​	注意： `log.Fatalf` 不是标准库日志。请参阅 [#logging]。



### Must functions 必须函数

Setup helper functions that stop the program on failure follow the naming convention `MustXYZ` (or `mustXYZ`). In general, they should only be called early on program startup, not on things like user input where normal Go error handling is preferred.

​	设置在失败时停止程序的帮助器函数遵循命名约定 `MustXYZ` （或 `mustXYZ` ）。通常，它们应该只在程序启动早期调用，而不是在诸如用户输入之类的事情上，在这些事情上更喜欢正常的 Go 错误处理。

This often comes up for functions called to initialize package-level variables exclusively at [package initialization time](https://golang.org/ref/spec#Package_initialization) (e.g. [template.Must](https://golang.org/pkg/text/template/#Must) and [regexp.MustCompile](https://golang.org/pkg/regexp/#MustCompile)).

​	这通常适用于在程序包初始化时专门用于初始化程序包级变量的函数（例如 template.Must 和 regexp.MustCompile）。

``` go
// Good:
func MustParse(version string) *Version {
    v, err := Parse(version)
    if err != nil {
        log.Fatalf("MustParse(%q) = _, %v", version, err)
    }
    return v
}

// Package level "constant". If we wanted to use `Parse`, we would have had to
// set the value in `init`.
var DefaultVersion = MustParse("1.2.3")
```

**Note:** `log.Fatalf` is not the standard library log. See [#logging].

​	注意： `log.Fatalf` 不是标准库日志。请参阅 [#logging]。

The same convention may be used in test helpers that only stop the current test (using `t.Fatal`). Such helpers are often convenient in creating test values, for example in struct fields of [table driven tests](https://google.github.io/styleguide/go/decisions#table-driven-tests), as functions that return errors cannot be directly assigned to a struct field.

​	在仅停止当前测试（使用 `t.Fatal` ）的测试帮助器中可以使用相同的约定。此类帮助器通常在创建测试值时很方便，例如在表驱动测试的结构字段中，因为返回错误的函数无法直接分配给结构字段。

``` go
// Good:
func mustMarshalAny(t *testing.T, m proto.Message) *anypb.Any {
  t.Helper()
  any, err := anypb.New(m)
  if err != nil {
    t.Fatalf("MustMarshalAny(t, m) = %v; want %v", err, nil)
  }
  return any
}

func TestCreateObject(t *testing.T) {
  tests := []struct{
    desc string
    data *anypb.Any
  }{
    {
      desc: "my test case",
      // Creating values directly within table driven test cases.
      data: mustMarshalAny(t, mypb.Object{}),
    },
    // ...
  }
  // ...
}
```

In both of these cases, the value of this pattern is that the helpers can be called in a “value” context. These helpers should not be called in places where it’s difficult to ensure an error would be caught or in a context where an error should be [checked](https://google.github.io/styleguide/go/decisions#handle-errors) (e.g., in many request handlers). For constant inputs, this allows tests to easily ensure that the `Must` arguments are well-formed, and for non-constant inputs it permits tests to validate that errors are [properly handled or propagated](https://google.github.io/styleguide/go/best-practices#error-handling).

​	在这两种情况下，此模式的价值在于可以在“值”上下文中调用帮助器。不应在难以确保捕获错误或应检查错误的上下文中（例如，在许多请求处理程序中）调用这些帮助器。对于常量输入，这允许测试轻松确保 `Must` 参数格式正确，对于非常量输入，它允许测试验证错误是否得到正确处理或传播。

Where `Must` functions are used in a test, they should generally be [marked as a test helper](https://google.github.io/styleguide/go/decisions#mark-test-helpers) and call `t.Fatal` on error (see [error handling in test helpers](https://google.github.io/styleguide/go/best-practices#test-helper-error-handling) for more considerations of using that).

​	在测试中使用 `Must` 函数时，通常应将其标记为测试帮助器，并在出错时调用 `t.Fatal` （有关使用该方法的更多注意事项，请参阅测试帮助器中的错误处理）。

They should not be used when [ordinary error handling](https://google.github.io/styleguide/go/best-practices#error-handling) is possible (including with some refactoring):

​	在可能进行普通错误处理时不应使用它们（包括进行一些重构）：

``` go
// Bad:
func Version(o *servicepb.Object) (*version.Version, error) {
    // Return error instead of using Must functions.
    v := version.MustParse(o.GetVersionString())
    return dealiasVersion(v)
}
```



### Goroutine lifetimes Goroutine 生命周期



When you spawn goroutines, make it clear when or whether they exit.

​	当您生成 goroutine 时，要明确它们何时或是否退出。

Goroutines can leak by blocking on channel sends or receives. The garbage collector will not terminate a goroutine even if the channels it is blocked on are unreachable.

​	Goroutine 可能会因在通道发送或接收时阻塞而泄漏。即使 goroutine 阻塞的通道不可达，垃圾回收器也不会终止该 goroutine。

Even when goroutines do not leak, leaving them in-flight when they are no longer needed can cause other subtle and hard-to-diagnose problems. Sending on a channel that has been closed causes a panic.

​	即使 goroutine 没有泄漏，在不再需要时让它们处于飞行状态也可能导致其他难以诊断的微妙问题。向已关闭的通道发送数据会导致 panic。

``` go
// Bad:
ch := make(chan int)
ch <- 42
close(ch)
ch <- 13 // panic
```

Modifying still-in-use inputs “after the result isn’t needed” can lead to data races. Leaving goroutines in-flight for arbitrarily long can lead to unpredictable memory usage.

​	在“结果不再需要”后修改仍在使用的输入可能会导致数据竞争。让 goroutine 处于飞行状态的时间过长可能会导致内存使用情况不可预测。

Concurrent code should be written such that the goroutine lifetimes are obvious. Typically this will mean keeping synchronization-related code constrained within the scope of a function and factoring out the logic into [synchronous functions](https://google.github.io/styleguide/go/decisions#synchronous-functions). If the concurrency is still not obvious, it is important to document when and why the goroutines exit.

​	应编写并发代码，以便 goroutine 的生命周期显而易见。通常，这意味着将与同步相关的代码限制在函数的范围内，并将逻辑分解为同步函数。如果并发性仍然不明显，则务必记录 goroutine 何时以及为何退出。

Code that follows best practices around context usage often helps make this clear. It is conventionally managed with a `context.Context`:

​	遵循有关上下文使用情况的最佳实践的代码通常有助于明确这一点。通常使用 `context.Context` 管理：

``` go
// Good:
func (w *Worker) Run(ctx context.Context) error {
    // ...
    for item := range w.q {
        // process returns at latest when the context is cancelled.
        go process(ctx, item)
    }
    // ...
}
```

There are other variants of the above that use raw signal channels like `chan struct{}`, synchronized variables, [condition variables](https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view), and more. The important part is that the goroutine’s end is evident for subsequent maintainers.

​	上述内容还有其他变体，它们使用原始信号通道，如 `chan struct{}` 、同步变量、条件变量等。重要的是，后续维护人员可以明显看出 goroutine 的结束。

In contrast, the following code is careless about when its spawned goroutines finish:

​	相比之下，以下代码对其派生的 goroutine 何时结束并不关心：

``` go
// Bad:
func (w *Worker) Run() {
    // ...
    for item := range w.q {
        // process returns when it finishes, if ever, possibly not cleanly
        // handling a state transition or termination of the Go program itself.
        go process(item)
    }
    // ...
}
```

This code may look OK, but there are several underlying problems:

​	此代码可能看起来没问题，但存在几个潜在问题：

- The code probably has undefined behavior in production, and the program may not terminate cleanly, even if the operating system releases the resources.

  ​	此代码在生产中可能具有未定义的行为，即使操作系统释放了资源，程序也可能无法干净地终止。

- The code is difficult to test meaningfully due to the code’s indeterminate lifecycle.

  ​	由于代码的生命周期不确定，因此很难有意义地测试代码。

- The code may leak resources as described above.

  ​	如上所述，代码可能会泄漏资源。

See also:

​	另请参阅：

- [Never start a goroutine without knowing how it will stop
  切勿在不知道如何停止 goroutine 的情况下启动它](https://dave.cheney.net/2016/12/22/never-start-a-goroutine-without-knowing-how-it-will-stop)
- Rethinking Classical Concurrency Patterns: [slides](https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view), [video](https://www.youtube.com/watch?v=5zXAHh5tJqQ)
  重新思考经典并发模式：幻灯片、视频
- [When Go programs end
  Go 程序何时结束](https://changelog.com/gotime/165)



### Interfaces 接口



Go interfaces generally belong in the package that *consumes* values of the interface type, not a package that *implements* the interface type. The implementing package should return concrete (usually pointer or struct) types. That way, new methods can be added to implementations without requiring extensive refactoring. See [GoTip #49: Accept Interfaces, Return Concrete Types](https://google.github.io/styleguide/go/index.html#gotip) for more details.

​	Go 接口通常属于使用接口类型值的包，而不是实现接口类型的包。实现包应返回具体类型（通常是指针或结构）。这样，可以向实现中添加新方法，而无需进行大规模重构。有关更多详细信息，请参阅 GoTip #49：接受接口，返回具体类型。

Do not export a [test double](https://abseil.io/resources/swe-book/html/ch13.html#techniques_for_using_test_doubles) implementation of an interface from an API that consumes it. Instead, design the API so that it can be tested using the [public API](https://abseil.io/resources/swe-book/html/ch12.html#test_via_public_apis) of the [real implementation](https://google.github.io/styleguide/go/best-practices#use-real-transports). See [GoTip #42: Authoring a Stub for Testing](https://google.github.io/styleguide/go/index.html#gotip) for more details. Even when it is not feasible to use the real implementation, it may not be necessary to introduce an interface fully covering all methods in the real type; the consumer can create an interface containing only the methods it needs, as demonstrated in [GoTip #78: Minimal Viable Interfaces](https://google.github.io/styleguide/go/index.html#gotip).

​	不要从使用它的 API 中导出接口的测试双重实现。相反，设计 API，以便可以使用真实实现的公共 API 对其进行测试。有关更多详细信息，请参阅 GoTip #42：编写用于测试的存根。即使无法使用真实实现，也可能不必引入一个完全涵盖真实类型中所有方法的接口；使用者可以创建一个仅包含其所需方法的接口，如 GoTip #78：最小可行接口中所示。

To test packages that use Stubby RPC clients, use a real client connection. If a real server cannot be run in the test, Google’s internal practice is to obtain a real client connection to a local [test double](https://abseil.io/resources/swe-book/html/ch13.html#basic_concepts) using the internal rpctest package (coming soon!).

​	要测试使用 Stubby RPC 客户端的包，请使用真正的客户端连接。如果无法在测试中运行真正的服务器，Google 的内部做法是使用内部 rpctest 包（即将推出！）获取到本地测试双重的真正客户端连接。

Do not define interfaces before they are used (see [TotT: Code Health: Eliminate YAGNI Smells](https://testing.googleblog.com/2017/08/code-health-eliminate-yagni-smells.html) ). Without a realistic example of usage, it is too difficult to see whether an interface is even necessary, let alone what methods it should contain.

​	在使用之前不要定义接口（参见 TotT：代码健康：消除 YAGNI 气味）。如果没有实际的使用示例，很难看出是否需要接口，更不用说它应该包含哪些方法了。

Do not use interface-typed parameters if the users of the package do not need to pass different types for them.

​	如果软件包的用户不需要为接口类型参数传递不同的类型，则不要使用接口类型参数。

Do not export interfaces that the users of the package do not need.

​	不要导出软件包的用户不需要的接口。

**TODO:** Write a more in-depth doc on interfaces and link to it here.

​	TODO：编写有关接口的更深入的文档并链接到此处。

``` go
// Good:
package consumer // consumer.go

type Thinger interface { Thing() bool }

func Foo(t Thinger) string { ... }
// Good:
package consumer // consumer_test.go

type fakeThinger struct{ ... }
func (t fakeThinger) Thing() bool { ... }
...
if Foo(fakeThinger{...}) == "x" { ... }
// Bad:
package producer

type Thinger interface { Thing() bool }

type defaultThinger struct{ ... }
func (t defaultThinger) Thing() bool { ... }

func NewThinger() Thinger { return defaultThinger{ ... } }
// Good:
package producer

type Thinger struct{ ... }
func (t Thinger) Thing() bool { ... }

func NewThinger() Thinger { return Thinger{ ... } }
```



### Generics 泛型

Generics (formally called “[Type Parameters](https://go.dev/design/43651-type-parameters)”) are allowed where they fulfill your business requirements. In many applications, a conventional approach using existing language features (slices, maps, interfaces, and so on) works just as well without the added complexity, so be wary of premature use. See the discussion on [least mechanism](https://google.github.io/styleguide/go/guide#least-mechanism).

​	泛型（正式名称为“类型参数”）允许在满足业务需求的情况下使用。在许多应用程序中，使用现有语言特性（切片、映射、接口等）的传统方法同样有效，而不会增加复杂性，因此要小心过早使用。请参阅有关最小机制的讨论。

When introducing an exported API that uses generics, make sure it is suitably documented. It’s highly encouraged to include motivating runnable [examples](https://google.github.io/styleguide/go/decisions#examples).

​	在引入使用泛型的导出 API 时，请确保对其进行适当的记录。强烈建议包含激励性的可运行示例。

Do not use generics just because you are implementing an algorithm or data structure that does not care about the type of its member elements. If there is only one type being instantiated in practice, start by making your code work on that type without using generics at all. Adding polymorphism later will be straightforward compared to removing abstraction that is found to be unnecessary.

​	不要仅仅因为你要实现一个不在乎其成员元素类型的算法或数据结构而使用泛型。如果在实践中只实例化一种类型，那么首先让你的代码在不使用泛型的情况下适用于该类型。与移除被发现不必要的抽象相比，稍后添加多态性将非常简单。

Do not use generics to invent domain-specific languages (DSLs). In particular, refrain from introducing error-handling frameworks that might put a significant burden on readers. Instead prefer established [error handling](https://google.github.io/styleguide/go/decisions#errors) practices. For testing, be especially wary of introducing [assertion libraries](https://google.github.io/styleguide/go/decisions#assert) or frameworks that result in less useful [test failures](https://google.github.io/styleguide/go/decisions#useful-test-failures).

​	不要使用泛型来发明特定于领域的语言 (DSL)。特别是，不要引入可能给读者带来重大负担的错误处理框架。相反，更喜欢已建立的错误处理实践。对于测试，特别注意不要引入断言库或框架，因为这会导致不太有用的测试失败。

In general:

​	一般来说：

- [Write code, don’t design types](https://www.youtube.com/watch?v=Pa_e9EeCdy8&t=1250s). From a GopherCon talk by Robert Griesemer and Ian Lance Taylor.
  编写代码，不要设计类型。摘自 Robert Griesemer 和 Ian Lance Taylor 在 GopherCon 上的演讲。
- If you have several types that share a useful unifying interface, consider modeling the solution using that interface. Generics may not be needed.
  如果你有几种共享有用的统一接口的类型，请考虑使用该接口对解决方案进行建模。可能不需要泛型。
- Otherwise, instead of relying on the `any` type and excessive [type switching](https://tour.golang.org/methods/16), consider generics.
  否则，请考虑使用泛型，而不是依赖于 `any` 类型和过多的类型转换。

See also:

​	另请参阅：

- [Using Generics in Go](https://www.youtube.com/watch?v=nr8EpUO9jhw), talk by Ian Lance Taylor

  ​	Ian Lance Taylor 的演讲《在 Go 中使用泛型》

- [Generics tutorial](https://go.dev/doc/tutorial/generics) on Go’s webpage

  ​	Go 网页上的泛型教程



### Pass values 传递值



Do not pass pointers as function arguments just to save a few bytes. If a function reads its argument `x` only as `*x` throughout, then the argument shouldn’t be a pointer. Common instances of this include passing a pointer to a string (`*string`) or a pointer to an interface value (`*io.Reader`). In both cases, the value itself is a fixed size and can be passed directly.

​	不要仅仅为了节省几个字节就将指针作为函数参数传递。如果一个函数始终只将它的参数 `x` 作为 `*x` 来读取，那么该参数就不应该是一个指针。常见的实例包括传递一个指向字符串的指针 ( `*string` ) 或一个指向接口值的指针 ( `*io.Reader` )。在这两种情况下，值本身都是固定大小的，可以直接传递。

This advice does not apply to large structs, or even small structs that may increase in size. In particular, protocol buffer messages should generally be handled by pointer rather than by value. The pointer type satisfies the `proto.Message` interface (accepted by `proto.Marshal`, `protocmp.Transform`, etc.), and protocol buffer messages can be quite large and often grow larger over time.

​	此建议不适用于大型结构体，甚至可能增大的小型结构体。特别是，协议缓冲区消息通常应通过指针而不是通过值来处理。指针类型满足 `proto.Message` 接口（由 `proto.Marshal` 、 `protocmp.Transform` 等接受），并且协议缓冲区消息可能非常大，并且随着时间的推移通常会变得更大。



### Receiver type 接收器类型



A [method receiver](https://golang.org/ref/spec#Method_declarations) can be passed either as a value or a pointer, just as if it were a regular function parameter. The choice between the two is based on which [method set(s)](https://golang.org/ref/spec#Method_sets) the method should be a part of.

​	方法接收器可以作为值或指针传递，就像它是常规函数参数一样。两者之间的选择取决于该方法应属于哪个方法集。

**Correctness wins over speed or simplicity.** There are cases where you must use a pointer value. In other cases, pick pointers for large types or as future-proofing if you don’t have a good sense of how the code will grow, and use values for simple [plain old data](https://en.wikipedia.org/wiki/Passive_data_structure).

​	正确性胜过速度或简单性。在某些情况下，您必须使用指针值。在其他情况下，如果无法很好地了解代码的增长方式，请选择指针作为大型类型或作为未来证明，并对简单普通旧数据使用值。

The list below spells out each case in further detail:

​	以下列表更详细地说明了每种情况：

- If the receiver is a slice and the method doesn’t reslice or reallocate the slice, use a value rather than a pointer.

  ​	如果接收者是切片并且该方法不会重新切片或重新分配切片，请使用值而不是指针。

  ```
  // Good:
  type Buffer []byte
  
  func (b Buffer) Len() int { return len(b) }
  ```

- If the method needs to mutate the receiver, the receiver must be a pointer.

  ​	如果该方法需要改变接收者，则接收者必须是指针。

  ```
  // Good:
  type Counter int
  
  func (c *Counter) Inc() { *c++ }
  
  // See https://pkg.go.dev/container/heap.
  type Queue []Item
  
  func (q *Queue) Push(x Item) { *q = append([]Item{x}, *q...) }
  ```

- If the receiver is a struct containing fields that [cannot safely be copied](https://google.github.io/styleguide/go/decisions#copying), use a pointer receiver. Common examples are [`sync.Mutex`](https://pkg.go.dev/sync#Mutex) and other synchronization types.

  ​	如果接收者是包含无法安全复制的字段的结构，请使用指针接收者。常见的示例是 `sync.Mutex` 和其他同步类型。

  ```
  // Good:
  type Counter struct {
      mu    sync.Mutex
      total int
  }
  
  func (c *Counter) Inc() {
      c.mu.Lock()
      defer c.mu.Unlock()
      c.total++
  }
  ```

  **Tip:** Check the type’s [Godoc](https://pkg.go.dev/time#example-Duration) for information about whether it is safe or unsafe to copy.

  ​	提示：检查类型的 Godoc 以获取有关复制是否安全或不安全的信息。

- If the receiver is a “large” struct or array, a pointer receiver may be more efficient. Passing a struct is equivalent to passing all of its fields or elements as arguments to the method. If that seems too large to [pass by value](https://google.github.io/styleguide/go/decisions#pass-values), a pointer is a good choice.

  ​	如果接收者是“大型”结构或数组，则指针接收者可能更有效。传递结构等同于将它的所有字段或元素作为参数传递给该方法。如果传递值时看起来太大，则指针是一个不错的选择。

- For methods that will call or run concurrently with other functions that modify the receiver, use a value if those modifications should not be visible to your method; otherwise use a pointer.

  ​	对于将调用或与修改接收者的其他函数同时运行的方法，如果这些修改不应该对您的方法可见，则使用一个值；否则使用一个指针。

- If the receiver is a struct or array, any of whose elements is a pointer to something that may be mutated, prefer a pointer receiver to make the intention of mutability clear to the reader.

  ​	如果接收者是一个结构或数组，其中任何元素都是指向可能发生突变的某个内容的指针，则最好使用一个指针接收者，以便向读者明确可变性的意图。

  ```
  // Good:
  type Counter struct {
      m *Metric
  }
  
  func (c *Counter) Inc() {
      c.m.Add(1)
  }
  ```

- If the receiver is a [built-in type](https://pkg.go.dev/builtin), such as an integer or a string, that does not need to be modified, use a value.

  ​	如果接收者是一个内置类型，例如整数或字符串，不需要修改，则使用一个值。

  ```
  // Good:
  type User string
  
  func (u User) String() { return string(u) }
  ```

- If the receiver is a map, function, or channel, use a value rather than a pointer.

  ​	如果接收者是一个映射、函数或通道，则使用一个值而不是一个指针。

  ```
  // Good:
  // See https://pkg.go.dev/net/http#Header.
  type Header map[string][]string
  
  func (h Header) Add(key, value string) { /* omitted */ }
  ```

- If the receiver is a “small” array or struct that is naturally a value type with no mutable fields and no pointers, a value receiver is usually the right choice.

  ​	如果接收者是一个“小”数组或结构，它自然是一个值类型，没有可变字段也没有指针，则一个值接收者通常是正确的选择。

  ```
  // Good:
  // See https://pkg.go.dev/time#Time.
  type Time struct { /* omitted */ }
  
  func (t Time) Add(d Duration) Time { /* omitted */ }
  ```

- When in doubt, use a pointer receiver.

  ​	如有疑问，请使用一个指针接收者。

As a general guideline, prefer to make the methods for a type either all pointer methods or all value methods.

​	作为一般准则，最好使一个类型的方法全部都是指针方法或全部都是值方法。

**Note:** There is a lot of misinformation about whether passing a value or a pointer to a function can affect performance. The compiler can choose to pass pointers to values on the stack as well as copying values on the stack, but these considerations should not outweigh the readability and correctness of the code in most circumstances. When the performance does matter, it is important to profile both approaches with a realistic benchmark before deciding that one approach outperforms the other.

​	注意：关于传递值还是指针给函数是否会影响性能，有很多错误的信息。编译器可以选择将指针传递给堆栈上的值，也可以将值复制到堆栈上，但在大多数情况下，这些考虑因素不应超过代码的可读性和正确性。当性能很重要时，在决定一种方法优于另一种方法之前，使用现实的基准对两种方法进行分析非常重要。



### `switch` and `break` `switch` 和 `break`



Do not use `break` statements without target labels at the ends of `switch` clauses; they are redundant. Unlike in C and Java, `switch` clauses in Go automatically break, and a `fallthrough` statement is needed to achieve the C-style behavior. Use a comment rather than `break` if you want to clarify the purpose of an empty clause.

​	不要在 `switch` 子句的末尾使用没有目标标签的 `break` 语句；它们是多余的。与 C 和 Java 不同，Go 中的 `switch` 子句会自动中断，需要 `fallthrough` 语句才能实现 C 风格的行为。如果您想阐明空子句的目的，请使用注释而不是 `break` 。

``` go
// Good:
switch x {
case "A", "B":
    buf.WriteString(x)
case "C":
    // handled outside of the switch statement
default:
    return fmt.Errorf("unknown value: %q", x)
}
// Bad:
switch x {
case "A", "B":
    buf.WriteString(x)
    break // this break is redundant
case "C":
    break // this break is redundant
default:
    return fmt.Errorf("unknown value: %q", x)
}
```

> **Note:** If a `switch` clause is within a `for` loop, using `break` within `switch` does not exit the enclosing `for` loop.
>
> ​	注意：如果 `switch` 子句位于 `for` 循环中，则在 `switch` 中使用 `break` 不会退出封闭的 `for` 循环。
>
> ```
> for {
>   switch x {
>   case "A":
>      break // exits the switch, not the loop
>   }
> }
> ```
>
> To escape the enclosing loop, use a label on the `for` statement:
>
> ​	要退出封闭循环，请在 `for` 语句上使用标签：
>
> ```
> loop:
>   for {
>     switch x {
>     case "A":
>        break loop // exits the loop
>     }
>   }
> ```



### Synchronous functions 同步函数



Synchronous functions return their results directly and finish any callbacks or channel operations before returning. Prefer synchronous functions over asynchronous functions.

​	同步函数直接返回其结果，并在返回之前完成任何回调或通道操作。优先使用同步函数，而不是异步函数。

Synchronous functions keep goroutines localized within a call. This helps to reason about their lifetimes, and avoid leaks and data races. Synchronous functions are also easier to test, since the caller can pass an input and check the output without the need for polling or synchronization.

​	同步函数将 goroutine 本地化在一个调用中。这有助于推理它们的生存期，并避免泄漏和数据竞争。同步函数也更容易测试，因为调用者可以传递一个输入并检查输出，而无需轮询或同步。

If necessary, the caller can add concurrency by calling the function in a separate goroutine. However, it is quite difficult (sometimes impossible) to remove unnecessary concurrency at the caller side.

​	如果需要，调用者可以通过在单独的 goroutine 中调用函数来添加并发。但是，在调用者端移除不必要的并发非常困难（有时是不可能的）。

See also:

​	另请参阅：

- “Rethinking Classical Concurrency Patterns”, talk by Bryan Mills: [slides](https://drive.google.com/file/d/1nPdvhB0PutEJzdCq5ms6UI58dp50fcAN/view), [video](https://www.youtube.com/watch?v=5zXAHh5tJqQ)
  “重新思考经典并发模式”，Bryan Mills 的演讲：幻灯片、视频



### Type aliases 类型别名



Use a *type definition*, `type T1 T2`, to define a new type. Use a [*type alias*](http://golang.org/ref/spec#Type_declarations), `type T1 = T2`, to refer to an existing type without defining a new type. Type aliases are rare; their primary use is to aid migrating packages to new source code locations. Don’t use type aliasing when it is not needed.

​	使用类型定义 `type T1 T2` 来定义一个新类型。使用类型别名 `type T1 = T2` 来引用现有类型，而无需定义新类型。类型别名很少见；它们的主要用途是帮助将包迁移到新的源代码位置。在不需要时不要使用类型别名。



### Use %q 使用 %q



Go’s format functions (`fmt.Printf` etc.) have a `%q` verb which prints strings inside double-quotation marks.

​	Go 的格式化函数（ `fmt.Printf` 等）有一个 `%q` 动词，它将字符串打印在双引号内。

``` go
// Good:
fmt.Printf("value %q looks like English text", someText)
```

Prefer using `%q` over doing the equivalent manually, using `%s`:

​	优先使用 `%q` 而不是手动执行等效操作，使用 `%s` ：

``` go
// Bad:
fmt.Printf("value \"%s\" looks like English text", someText)
// Avoid manually wrapping strings with single-quotes too:
fmt.Printf("value '%s' looks like English text", someText)
```

Using `%q` is recommended in output intended for humans where the input value could possibly be empty or contain control characters. It can be very hard to notice a silent empty string, but `""` stands out clearly as such.

​	建议在供人类阅读的输出中使用 `%q` ，其中输入值可能为空或包含控制字符。很难注意到一个静默的空字符串，但 `""` 会清楚地显示出来。



### Use any 使用任何

Go 1.18 introduces an `any` type as an [alias](https://go.googlesource.com/proposal/+/master/design/18130-type-alias.md) to `interface{}`. Because it is an alias, `any` is equivalent to `interface{}` in many situations and in others it is easily interchangeable via an explicit conversion. Prefer to use `any` in new code.

​	Go 1.18 引入 `any` 类型作为 `interface{}` 的别名。由于它是一个别名，因此在许多情况下 `any` 等同于 `interface{}` ，而在其他情况下，它可以通过显式转换轻松互换。更喜欢在新的代码中使用 `any` 。

## Common libraries 常用库



### Flags 标志



Go programs in the Google codebase use an internal variant of the [standard `flag` package](https://golang.org/pkg/flag/). It has a similar interface but interoperates well with internal Google systems. Flag names in Go binaries should prefer to use underscores to separate words, though the variables that hold a flag’s value should follow the standard Go name style ([mixed caps](https://google.github.io/styleguide/go/guide#mixed-caps)). Specifically, the flag name should be in snake case, and the variable name should be the equivalent name in camel case.

​	Google 代码库中的 Go 程序使用标准 `flag` 包的内部变体。它具有类似的界面，但可以很好地与 Google 内部系统互操作。Go 二进制文件中的标志名称应优先使用下划线分隔单词，而保存标志值的变量应遵循标准的 Go 名称样式（混合大小写）。具体来说，标志名称应采用蛇形大小写，变量名称应采用等效的驼峰大小写。

``` go
// Good:
var (
    pollInterval = flag.Duration("poll_interval", time.Minute, "Interval to use for polling.")
)
// Bad:
var (
    poll_interval = flag.Int("pollIntervalSeconds", 60, "Interval to use for polling in seconds.")
)
```

Flags must only be defined in `package main` or equivalent.

​	标志只能在 `package main` 或等效项中定义。

General-purpose packages should be configured using Go APIs, not by punching through to the command-line interface; don’t let importing a library export new flags as a side effect. That is, prefer explicit function arguments or struct field assignment or much less frequently and under the strictest of scrutiny exported global variables. In the extremely rare case that it is necessary to break this rule, the flag name must clearly indicate the package that it configures.

​	通用软件包应使用 Go API 进行配置，而不是通过命令行界面进行配置；不要让导入库导出新标志作为副作用。也就是说，更喜欢显式函数参数或结构字段赋值，或者更少地使用经过最严格审查的导出全局变量。在极少数情况下，有必要打破此规则，标志名称必须清楚地表明它配置的软件包。

If your flags are global variables, place them in their own `var` group, following the imports section.

​	如果您的标志是全局变量，请将它们放在自己的 `var` 组中，位于导入部分之后。

There is additional discussion around best practices for creating [complex CLIs](https://google.github.io/styleguide/go/best-practices#complex-clis) with subcommands.

​	围绕创建具有子命令的复杂 CLI 的最佳实践还有其他讨论。

See also:

​	另请参阅：

- [Tip of the Week #45: Avoid Flags, Especially in Library Code
  本周提示 #45：避免使用标志，尤其是在库代码中](https://abseil.io/tips/45)
- [Go Tip #10: Configuration Structs and Flags
  Go 提示 #10：配置结构和标志](https://google.github.io/styleguide/go/index.html#gotip)
- [Go Tip #80: Dependency Injection Principles
  Go 提示 #80：依赖注入原则](https://google.github.io/styleguide/go/index.html#gotip)



### Logging 日志记录

Go programs in the Google codebase use a variant of the [standard `log` package](https://pkg.go.dev/log). It has a similar but more powerful interface and interoperates well with internal Google systems. An open source version of this library is available as [package `glog`](https://pkg.go.dev/github.com/golang/glog), and open source Google projects may use that, but this guide refers to it as `log` throughout.

​	Google 代码库中的 Go 程序使用标准 `log` 包的变体。它具有相似但更强大的界面，并且可以很好地与内部 Google 系统互操作。此库的开源版本可用作软件包 `glog` ，开源 Google 项目可以使用它，但本指南始终将其称为 `log` 。

**Note:** For abnormal program exits, this library uses `log.Fatal` to abort with a stacktrace, and `log.Exit` to stop without one. There is no `log.Panic` function as in the standard library.

​	注意：对于异常程序退出，此库使用 `log.Fatal` 带堆栈跟踪中止，使用 `log.Exit` 不带堆栈跟踪中止。标准库中没有 `log.Panic` 函数。

**Tip:** `log.Info(v)` is equivalent `log.Infof("%v", v)`, and the same goes for other logging levels. Prefer the non-formatting version when you have no formatting to do.

​	提示： `log.Info(v)` 等同于 `log.Infof("%v", v)` ，其他日志记录级别也是如此。在没有格式化操作时，优先使用非格式化版本。

See also:

​	另请参阅：

- Best practices on [logging errors](https://google.github.io/styleguide/go/best-practices#error-logging) and [custom verbosily levels](https://google.github.io/styleguide/go/best-practices#vlog)
  关于记录错误和自定义详细级别日志的最佳做法
- When and how to use the log package to [stop the program](https://google.github.io/styleguide/go/best-practices#checks-and-panics)
  何时以及如何使用日志包停止程序



### Contexts 上下文



Values of the [`context.Context`](https://pkg.go.dev/context) type carry security credentials, tracing information, deadlines, and cancellation signals across API and process boundaries. Unlike C++ and Java, which in the Google codebase use thread-local storage, Go programs pass contexts explicitly along the entire function call chain from incoming RPCs and HTTP requests to outgoing requests.

​	 `context.Context` 类型的变量会跨 API 和进程边界携带安全凭据、跟踪信息、截止时间和取消信号。与在 Google 代码库中使用线程本地存储的 C++ 和 Java 不同，Go 程序会通过整个函数调用链显式传递上下文，从传入的 RPC 和 HTTP 请求到传出的请求。

When passed to a function or method, `context.Context` is always the first parameter.

​	传递给函数或方法时， `context.Context` 始终是第一个参数。

``` go
func F(ctx context.Context /* other arguments */) {}
```

Exceptions are:

​	例外情况：

- In an HTTP handler, where the context comes from [`req.Context()`](https://pkg.go.dev/net/http#Request.Context).
  在 HTTP 处理程序中，上下文来自 `req.Context()` 。

- In streaming RPC methods, where the context comes from the stream.

  ​	在流式 RPC 方法中，上下文来自流。

  Code using gRPC streaming accesses a context from a `Context()` method in the generated server type, which implements `grpc.ServerStream`. See [gRPC Generated Code documentation](https://grpc.io/docs/languages/go/generated-code/).

  ​	使用 gRPC 流式传输的代码从生成的服务器类型中的 `Context()` 方法访问上下文，该方法实现了 `grpc.ServerStream` 。请参阅 gRPC 生成的代码文档。

- In entrypoint functions (see below for examples of such functions), use [`context.Background()`](https://pkg.go.dev/context/#Background).

  ​	在入口点函数中（有关此类函数的示例，请参见下文），请使用 `context.Background()` 。

  - In binary targets: `main`
    在二进制目标中： `main`
  - In general purpose code and libraries: `init`
    在通用代码和库中： `init`
  - In tests: `TestXXX`, `BenchmarkXXX`, `FuzzXXX`
    在测试中： `TestXXX` 、 `BenchmarkXXX` 、 `FuzzXXX`

> **Note**: It is very rare for code in the middle of a callchain to require creating a base context of its own using `context.Background()`. Always prefer taking a context from your caller, unless it’s the wrong context.
>
> ​	注意：在调用链中间的代码需要使用 `context.Background()` 创建自己的基本上下文的情况非常罕见。除非是错误的上下文，否则始终优先从调用者那里获取上下文。
>
> You may come across server libraries (the implementation of Stubby, gRPC, or HTTP in Google’s server framework for Go) that construct a fresh context object per request. These contexts are immediately filled with information from the incoming request, so that when passed to the request handler, the context’s attached values have been propagated to it across the network boundary from the client caller. Moreover, these contexts’ lifetimes are scoped to that of the request: when the request is finished, the context is cancelled.
>
> ​	您可能会遇到服务器库（Google 的 Go 服务器框架中 Stubby、gRPC 或 HTTP 的实现），这些库会为每个请求构建一个新的上下文对象。这些上下文会立即填充来自传入请求的信息，以便在传递给请求处理程序时，上下文的附加值已通过网络边界从客户端调用者传播到它。此外，这些上下文的生命周期范围限定为请求的生命周期：当请求完成时，上下文将被取消。
>
> Unless you are implementing a server framework, you shouldn’t create contexts with `context.Background()` in library code. Instead, prefer using context detachment, which is mentioned below, if there is an existing context available. If you think you do need `context.Background()` outside of entrypoint functions, discuss it with the Google Go style mailing list before committing to an implementation.
>
> ​	除非您正在实现服务器框架，否则不应在库代码中使用 `context.Background()` 创建上下文。相反，如果存在可用上下文，请优先使用下面提到的上下文分离。如果您认为确实需要在入口点函数之外使用 `context.Background()` ，请在提交实现之前与 Google Go 风格邮件列表讨论。

The convention that `context.Context` comes first in functions also applies to test helpers.

​	 `context.Context` 在函数中排在第一位的惯例也适用于测试帮助程序。

``` go
// Good:
func readTestFile(ctx context.Context, t *testing.T, path string) string {}
```

Do not add a context member to a struct type. Instead, add a context parameter to each method on the type that needs to pass it along. The one exception is for methods whose signature must match an interface in the standard library or in a third party library outside Google’s control. Such cases are very rare, and should be discussed with the Google Go style mailing list before implementation and readability review.

​	不要向结构类型添加上下文成员。相反，向需要传递上下文的类型上的每个方法添加一个上下文参数。唯一的例外是签名必须与 Google 无法控制的标准库或第三方库中的接口匹配的方法。这种情况非常罕见，在实现和可读性审查之前应与 Google Go 风格邮件列表讨论。

Code in the Google codebase that must spawn background operations which can run after the parent context has been cancelled can use an internal package for detachment. Follow [issue #40221](https://github.com/golang/go/issues/40221) for discussions on an open source alternative.

​	Google 代码库中的代码必须生成可以在父上下文被取消后运行的后台操作，可以使用内部包进行分离。关注问题 #40221 以讨论开源替代方案。

Since contexts are immutable, it is fine to pass the same context to multiple calls that share the same deadline, cancellation signal, credentials, parent trace, and so on.

​	由于上下文是不可变的，因此将相同的上下文传递给共享相同截止时间、取消信号、凭据、父跟踪等的多个调用是正常的。

See also:

​	另请参阅：

- [Contexts and structs
  上下文和结构](https://go.dev/blog/context-and-structs)



#### Custom contexts 自定义上下文

Do not create custom context types or use interfaces other than `context.Context` in function signatures. There are no exceptions to this rule.

​	不要创建自定义上下文类型或在函数签名中使用 `context.Context` 以外的接口。此规则没有任何例外。

Imagine if every team had a custom context. Every function call from package `p` to package `q` would have to determine how to convert a `p.Context` to a `q.Context`, for all pairs of packages `p` and `q`. This is impractical and error-prone for humans, and it makes automated refactorings that add context parameters nearly impossible.

​	想象一下，如果每个团队都有一个自定义上下文。从包 `p` 到包 `q` 的每个函数调用都必须确定如何将 `p.Context` 转换为 `q.Context` ，适用于所有包 `p` 和 `q` 的对。这对人类来说不切实际且容易出错，并且它使添加上下文参数的自动化重构几乎不可能。

If you have application data to pass around, put it in a parameter, in the receiver, in globals, or in a `Context` value if it truly belongs there. Creating your own context type is not acceptable since it undermines the ability of the Go team to make Go programs work properly in production.

​	如果您有要传递的应用程序数据，请将其放在参数中、接收器中、全局变量中或 `Context` 值中（如果它确实属于那里）。创建您自己的上下文类型是不可接受的，因为它破坏了 Go 团队使 Go 程序在生产中正常工作的可能性。



### crypto/rand



Do not use package `math/rand` to generate keys, even throwaway ones. If unseeded, the generator is completely predictable. Seeded with `time.Nanoseconds()`, there are just a few bits of entropy. Instead, use `crypto/rand`’s Reader, and if you need text, print to hexadecimal or base64.

​	不要使用包 `math/rand` 来生成密钥，即使是临时密钥。如果未播种，生成器完全可预测。使用 `time.Nanoseconds()` 播种，只有几比特熵。相反，使用 `crypto/rand` 的 Reader，如果您需要文本，请打印为十六进制或 base64。

``` go
// Good:
import (
    "crypto/rand"
    // "encoding/base64"
    // "encoding/hex"
    "fmt"

    // ...
)

func Key() string {
    buf := make([]byte, 16)
    if _, err := rand.Read(buf); err != nil {
        log.Fatalf("Out of randomness, should never happen: %v", err)
    }
    return fmt.Sprintf("%x", buf)
    // or hex.EncodeToString(buf)
    // or base64.StdEncoding.EncodeToString(buf)
}
```

**Note:** `log.Fatalf` is not the standard library log. See [#logging].

​	注意： `log.Fatalf` 不是标准库日志。请参阅 [#logging]。



## Useful test failures 有用的测试失败



It should be possible to diagnose a test’s failure without reading the test’s source. Tests should fail with helpful messages detailing:

​	应该可以诊断测试的失败，而无需阅读测试的源代码。测试应该以详细说明以下内容的帮助消息失败：

- What caused the failure
  导致失败的原因
- What inputs resulted in an error
  导致错误的输入
- The actual result
  实际结果
- What was expected
  预期结果

Specific conventions for achieving this goal are outlined below.

​	实现此目标的具体约定概述如下。



### Assertion libraries 断言库



Do not create “assertion libraries” as helpers for testing.

​	不要创建“断言库”作为测试的辅助工具。

Assertion libraries are libraries that attempt to combine the validation and production of failure messages within a test (though the same pitfalls can apply to other test helpers as well). For more on the distinction between test helpers and assertion libraries, see [best practices](https://google.github.io/styleguide/go/best-practices#test-functions).

​	断言库是尝试在测试中结合验证和生成失败消息的库（尽管同样的陷阱也适用于其他测试辅助工具）。有关测试辅助工具和断言库之间区别的更多信息，请参阅最佳做法。

``` go
// Bad:
var obj BlogPost

assert.IsNotNil(t, "obj", obj)
assert.StringEq(t, "obj.Type", obj.Type, "blogPost")
assert.IntEq(t, "obj.Comments", obj.Comments, 2)
assert.StringNotEq(t, "obj.Body", obj.Body, "")
```

Assertion libraries tend to either stop the test early (if `assert` calls `t.Fatalf` or `panic`) or omit relevant information about what the test got right:

​	断言库往往会过早地停止测试（如果 `assert` 调用 `t.Fatalf` 或 `panic` ）或省略有关测试正确内容的相关信息：

``` go
// Bad:
package assert

func IsNotNil(t *testing.T, name string, val interface{}) {
    if val == nil {
        t.Fatalf("data %s = nil, want not nil", name)
    }
}

func StringEq(t *testing.T, name, got, want string) {
    if got != want {
        t.Fatalf("data %s = %q, want %q", name, got, want)
    }
}
```

Complex assertion functions often do not provide [useful failure messages](https://google.github.io/styleguide/go/decisions#useful-test-failures) and context that exists within the test function. Too many assertion functions and libraries lead to a fragmented developer experience: which assertion library should I use, what style of output format should it emit, etc.? Fragmentation produces unnecessary confusion, especially for library maintainers and authors of large-scale changes, who are responsible for fixing potential downstream breakages. Instead of creating a domain-specific language for testing, use Go itself.

​	复杂的断言函数通常不会提供有用的失败消息和测试函数中存在的上下文。过多的断言函数和库会导致开发人员体验支离破碎：我应该使用哪个断言库，它应该发出什么样式的输出格式等？这种支离破碎会产生不必要的混乱，尤其是对于库维护者和大规模更改的作者而言，他们负责修复潜在的下游中断。不要创建用于测试的特定于领域的语言，而应使用 Go 本身。

Assertion libraries often factor out comparisons and equality checks. Prefer using standard libraries such as [`cmp`](https://pkg.go.dev/github.com/google/go-cmp/cmp) and [`fmt`](https://golang.org/pkg/fmt/) instead:

​	断言库通常会分解比较和相等性检查。最好使用标准库，例如 `cmp` 和 `fmt` ：

``` go
// Good:
var got BlogPost

want := BlogPost{
    Comments: 2,
    Body:     "Hello, world!",
}

if !cmp.Equal(got, want) {
    t.Errorf("blog post = %v, want = %v", got, want)
}
```

For more domain-specific comparison helpers, prefer returning a value or an error that can be used in the test’s failure message instead of passing `*testing.T` and calling its error reporting methods:

​	对于更具体的领域比较帮助程序，最好返回一个值或一个错误，该值或错误可用于测试的失败消息中，而不是传递 `*testing.T` 并调用其错误报告方法：

``` go
// Good:
func postLength(p BlogPost) int { return len(p.Body) }

func TestBlogPost_VeritableRant(t *testing.T) {
    post := BlogPost{Body: "I am Gunnery Sergeant Hartman, your senior drill instructor."}

    if got, want := postLength(post), 60; got != want {
        t.Errorf("length of post = %v, want %v", got, want)
    }
}
```

**Best Practice:** Were `postLength` non-trivial, it would make sense to test it directly, independently of any tests that use it.

​	最佳做法：如果 `postLength` 非平凡，那么直接测试它是有意义的，独立于任何使用它的测试。

See also:

​	另请参阅：

- [Equality comparison and diffs
  相等比较和差异](https://google.github.io/styleguide/go/decisions#types-of-equality)
- [Print diffs
  打印差异](https://google.github.io/styleguide/go/decisions#print-diffs)
- For more on the distinction between test helpers and assertion helpers, see [best practices](https://google.github.io/styleguide/go/best-practices#test-functions)
  有关测试帮助程序和断言帮助程序之间区别的更多信息，请参阅最佳做法



### Identify the function 识别函数

In most tests, failure messages should include the name of the function that failed, even though it seems obvious from the name of the test function. Specifically, your failure message should be `YourFunc(%v) = %v, want %v` instead of just `got %v, want %v`.

​	在大多数测试中，失败消息应包括失败函数的名称，即使从测试函数的名称中可以明显看出。具体来说，您的失败消息应为 `YourFunc(%v) = %v, want %v` ，而不仅仅是 `got %v, want %v` 。



### Identify the input 识别输入

In most tests, failure messages should include the function inputs if they are short. If the relevant properties of the inputs are not obvious (for example, because the inputs are large or opaque), you should name your test cases with a description of what’s being tested and print the description as part of your error message.

​	在大多数测试中，如果失败消息的函数输入很短，则应包括这些输入。如果输入的相关属性不明显（例如，因为输入很大或不透明），则应使用正在测试内容的说明来命名测试用例，并将说明作为错误消息的一部分打印出来。



### Got before want 先得到再想要

Test outputs should include the actual value that the function returned before printing the value that was expected. A standard format for printing test outputs is `YourFunc(%v) = %v, want %v`. Where you would write “actual” and “expected”, prefer using the words “got” and “want”, respectively.

​	测试输出应包括函数返回的实际值，然后再打印预期的值。打印测试输出的标准格式是 `YourFunc(%v) = %v, want %v` 。在您要写“实际”和“预期”的地方，最好分别使用“得到”和“想要”这两个词。

For diffs, directionality is less apparent, and as such it is important to include a key to aid in interpreting the failure. See the [section on printing diffs](https://google.github.io/styleguide/go/decisions#print-diffs). Whichever diff order you use in your failure messages, you should explicitly indicate it as a part of the failure message, because existing code is inconsistent about the ordering.

​	对于差异，方向性不太明显，因此包含一个有助于解释失败的关键非常重要。请参阅有关打印差异的部分。无论您在失败消息中使用哪种差异顺序，都应将其明确指示为失败消息的一部分，因为现有代码在排序方面不一致。



### Full structure comparisons 完整结构比较

If your function returns a struct (or any data type with multiple fields such as slices, arrays, and maps), avoid writing test code that performs a hand-coded field-by-field comparison of the struct. Instead, construct the data that you’re expecting your function to return, and compare directly using a [deep comparison](https://google.github.io/styleguide/go/decisions#types-of-equality).

​	如果您的函数返回一个结构（或任何具有多个字段的数据类型，例如切片、数组和映射），请避免编写执行结构的手工字段对字段比较的测试代码。相反，构建您希望函数返回的数据，并使用深度比较直接进行比较。

**Note:** This does not apply if your data contains irrelevant fields that obscure the intention of the test.

​	注意：如果您的数据包含模糊测试意图的不相关字段，则此方法不适用。

If your struct needs to be compared for approximate (or equivalent kind of semantic) equality or it contains fields that cannot be compared for equality (e.g., if one of the fields is an `io.Reader`), tweaking a [`cmp.Diff`](https://pkg.go.dev/github.com/google/go-cmp/cmp#Diff) or [`cmp.Equal`](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal) comparison with [`cmpopts`](https://pkg.go.dev/github.com/google/go-cmp/cmp/cmpopts) options such as [`cmpopts.IgnoreInterfaces`](https://pkg.go.dev/github.com/google/go-cmp/cmp/cmpopts#IgnoreInterfaces) may meet your needs ([example](https://play.golang.org/p/vrCUNVfxsvF)).

​	如果您的结构需要比较近似（或等效语义）相等性，或者它包含无法比较相等性的字段（例如，如果其中一个字段是 `io.Reader` ），则调整 `cmp.Diff` 或 `cmp.Equal` 与 `cmpopts` 选项（例如 `cmpopts.IgnoreInterfaces` ）的比较可能会满足您的需求（示例）。

If your function returns multiple return values, you don’t need to wrap those in a struct before comparing them. Just compare the return values individually and print them.

​	如果您的函数返回多个返回值，则无需在比较它们之前将它们包装在结构中。只需分别比较返回值并打印它们。

``` go
// Good:
val, multi, tail, err := strconv.UnquoteChar(`\"Fran & Freddie's Diner\"`, '"')
if err != nil {
  t.Fatalf(...)
}
if val != `"` {
  t.Errorf(...)
}
if multi {
  t.Errorf(...)
}
if tail != `Fran & Freddie's Diner"` {
  t.Errorf(...)
}
```



### Compare stable results 比较稳定结果

Avoid comparing results that may depend on output stability of a package that you do not own. Instead, the test should compare on semantically relevant information that is stable and resistant to changes in dependencies. For functionality that returns a formatted string or serialized bytes, it is generally not safe to assume that the output is stable.

​	避免比较可能取决于您不拥有的软件包的输出稳定性的结果。相反，测试应比较语义相关的信息，这些信息是稳定的并且能够抵抗依赖项的变化。对于返回格式化字符串或序列化字节的功能，通常不安全地假设输出是稳定的。

For example, [`json.Marshal`](https://golang.org/pkg/encoding/json/#Marshal) can change (and has changed in the past) the specific bytes that it emits. Tests that perform string equality on the JSON string may break if the `json` package changes how it serializes the bytes. Instead, a more robust test would parse the contents of the JSON string and ensure that it is semantically equivalent to some expected data structure.

​	例如， `json.Marshal` 可以更改（并且过去已经更改）它发出的特定字节。如果 `json` 包更改了它序列化字节的方式，则对 JSON 字符串执行字符串相等性的测试可能会中断。相反，更健壮的测试将解析 JSON 字符串的内容并确保它在语义上等同于某些预期的数据结构。



### Keep going 继续进行

Tests should keep going for as long as possible, even after a failure, in order to print out all of the failed checks in a single run. This way, a developer who is fixing the failing test doesn’t have to re-run the test after fixing each bug to find the next bug.

​	即使在发生故障后，测试也应尽可能长时间地继续进行，以便在单次运行中打印出所有失败的检查。这样，修复失败测试的开发人员不必在修复每个错误后重新运行测试以查找下一个错误。

Prefer calling `t.Error` over `t.Fatal` for reporting a mismatch. When comparing several different properties of a function’s output, use `t.Error` for each of those comparisons.

​	首选调用 `t.Error` 而不是 `t.Fatal` 来报告不匹配。比较函数输出的多个不同属性时，对每个比较使用 `t.Error` 。

Calling `t.Fatal` is primarily useful for reporting an unexpected error condition, when subsequent comparison failures are not going to be meaningful.

​	调用 `t.Fatal` 主要用于报告意外的错误情况，此时后续比较失败将没有意义。

For table-driven test, consider using subtests and use `t.Fatal` rather than `t.Error` and `continue`. See also [GoTip #25: Subtests: Making Your Tests Lean](https://google.github.io/styleguide/go/index.html#gotip).

​	对于表驱动测试，请考虑使用子测试并使用 `t.Fatal` 而不是 `t.Error` 和 `continue` 。另请参阅 GoTip #25：子测试：精简您的测试。

**Best practice:** For more discussion about when `t.Fatal` should be used, see [best practices](https://google.github.io/styleguide/go/best-practices#t-fatal).

​	最佳实践：有关何时应使用 `t.Fatal` 的更多讨论，请参阅最佳实践。



### Equality comparison and diffs 相等比较和差异

The `==` operator evaluates equality using [language-defined comparisons](http://golang.org/ref/spec#Comparison_operators). Scalar values (numbers, booleans, etc) are compared based on their values, but only some structs and interfaces can be compared in this way. Pointers are compared based on whether they point to the same variable, rather than based on the equality of the values to which they point.

​	 `==` 运算符使用语言定义的比较来评估相等性。标量值（数字、布尔值等）根据其值进行比较，但只有某些结构和接口才能以这种方式进行比较。指针的比较基于它们是否指向同一个变量，而不是基于它们指向的值的相等性。

The [`cmp`](https://pkg.go.dev/github.com/google/go-cmp/cmp) package can compare more complex data structures not appropriately handled by `==`, such as slices. Use [`cmp.Equal`](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal) for equality comparison and [`cmp.Diff`](https://pkg.go.dev/github.com/google/go-cmp/cmp#Diff) to obtain a human-readable diff between objects.

​	 `cmp` 包可以比较 `==` 未适当地处理的更复杂的数据结构，例如切片。使用 `cmp.Equal` 进行相等比较，并使用 `cmp.Diff` 获取对象之间可供人类阅读的差异。

``` go
// Good:
want := &Doc{
    Type:     "blogPost",
    Comments: 2,
    Body:     "This is the post body.",
    Authors:  []string{"isaac", "albert", "emmy"},
}
if !cmp.Equal(got, want) {
    t.Errorf("AddPost() = %+v, want %+v", got, want)
}
```

As a general-purpose comparison library, `cmp` may not know how to compare certain types. For example, it can only compare protocol buffer messages if passed the [`protocmp.Transform`](https://pkg.go.dev/google.golang.org/protobuf/testing/protocmp#Transform) option.

​	作为通用比较库， `cmp` 可能不知道如何比较某些类型。例如，它只能在传递 `protocmp.Transform` 选项时比较协议缓冲区消息。

``` go
// Good:
if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
    t.Errorf("Foo() returned unexpected difference in protobuf messages (-want +got):\n%s", diff)
}
```

Although the `cmp` package is not part of the Go standard library, it is maintained by the Go team and should produce stable equality results over time. It is user-configurable and should serve most comparison needs.

​	尽管 `cmp` 包不属于 Go 标准库，但它由 Go 团队维护，并且应该会随着时间的推移产生稳定的相等结果。它是用户可配置的，应该可以满足大多数比较需求。

Existing code may make use of the following older libraries, and may continue using them for consistency:

​	现有代码可能使用以下较旧的库，并且可能继续使用它们以保持一致性：

- [`pretty`](https://pkg.go.dev/github.com/kylelemons/godebug/pretty) produces aesthetically pleasing difference reports. However, it quite deliberately considers values that have the same visual representation as equal. In particular, `pretty` does not catch differences between nil slices and empty ones, is not sensitive to different interface implementations with identical fields, and it is possible to use a nested map as the basis for comparison with a struct value. It also serializes the entire value into a string before producing a diff, and as such is not a good choice for comparing large values. By default, it compares unexported fields, which makes it sensitive to changes in implementation details in your dependencies. For this reason, it is not appropriate to use `pretty` on protobuf messages.
  `pretty` 会生成美观的差异报告。但是，它非常刻意地将具有相同视觉表示的值视为相等。特别是， `pretty` 不会捕获 nil 切片和空切片之间的差异，对具有相同字段的不同接口实现不敏感，并且可以使用嵌套映射作为与结构值进行比较的基础。它还会在生成差异之前将整个值序列化为字符串，因此不适合比较大值。默认情况下，它会比较未导出的字段，这使其对依赖项中的实现细节更改很敏感。因此，在 protobuf 消息中使用 `pretty` 不合适。

Prefer using `cmp` for new code, and it is worth considering updating older code to use `cmp` where and when it is practical to do so.

​	对于新代码，最好使用 `cmp` ，并且值得考虑在实际可行的情况下更新较旧的代码以使用 `cmp` 。

Older code may use the standard library `reflect.DeepEqual` function to compare complex structures. `reflect.DeepEqual` should not be used for checking equality, as it is sensitive to changes in unexported fields and other implementation details. Code that is using `reflect.DeepEqual` should be updated to one of the above libraries.

​	旧代码可能会使用标准库 `reflect.DeepEqual` 函数来比较复杂结构。 `reflect.DeepEqual` 不应用于检查相等性，因为它对未导出的字段和其他实现细节的更改很敏感。使用 `reflect.DeepEqual` 的代码应更新到上述库之一。

**Note:** The `cmp` package is designed for testing, rather than production use. As such, it may panic when it suspects that a comparison is performed incorrectly to provide instruction to users on how to improve the test to be less brittle. Given cmp’s propensity towards panicking, it makes it unsuitable for code that is used in production as a spurious panic may be fatal.

​	注意： `cmp` 包专为测试而设计，而不是生产用途。因此，当它怀疑比较执行不正确时，可能会引发恐慌，以便向用户提供有关如何改进测试以减少脆弱性的说明。鉴于 cmp 倾向于恐慌，这使得它不适合用于生产中的代码，因为虚假恐慌可能是致命的。



### Level of detail 详细程度

The conventional failure message, which is suitable for most Go tests, is `YourFunc(%v) = %v, want %v`. However, there are cases that may call for more or less detail:

​	适用于大多数 Go 测试的常规失败消息是 `YourFunc(%v) = %v, want %v` 。但是，在某些情况下可能需要更多或更少的详细信息：

- Tests performing complex interactions should describe the interactions too. For example, if the same `YourFunc` is called several times, identify which call failed the test. If it’s important to know any extra state of the system, include that in the failure output (or at least in the logs).
  执行复杂交互的测试也应描述交互。例如，如果多次调用相同的 `YourFunc` ，请确定哪个调用使测试失败。如果了解系统的任何额外状态很重要，请将其包含在失败输出中（或至少包含在日志中）。
- If the data is a complex struct with significant boilerplate, it is acceptable to describe only the important parts in the message, but do not overly obscure the data.
  如果数据是具有大量样板代码的复杂结构，则仅描述消息中的重要部分是可以接受的，但不要过度模糊数据。
- Setup failures do not require the same level of detail. If a test helper populates a Spanner table but Spanner was down, you probably don’t need to include which test input you were going to store in the database. `t.Fatalf("Setup: Failed to set up test database: %s", err)` is usually helpful enough to resolve the issue.
  设置失败不需要相同的详细程度。如果测试助手填充了 Spanner 表，但 Spanner 已关闭，则可能不需要包含要存储在数据库中的测试输入。通常足以解决问题。

**Tip:** Make your failure mode trigger during development. Review what the failure message looks like and whether a maintainer can effectively deal with the failure.

​	提示：在开发期间触发故障模式。查看故障消息的显示方式以及维护人员是否可以有效地处理故障。

There are some techniques for reproducing test inputs and outputs clearly:

​	有一些技术可以清晰地重现测试输入和输出：

- When printing string data, [`%q` is often useful](https://google.github.io/styleguide/go/decisions#use-percent-q) to emphasize that the value is important and to more easily spot bad values.
  打印字符串数据时， `%q` 通常可用于强调该值很重要，并更轻松地发现错误值。
- When printing (small) structs, `%+v` can be more useful than `%v`.
  打印（小）结构时， `%+v` 可能比 `%v` 更有用。
- When validation of larger values fails, [printing a diff](https://google.github.io/styleguide/go/decisions#print-diffs) can make it easier to understand the failure.
  当较大值的验证失败时，打印差异可以更容易地理解失败。



### Print diffs 打印差异

If your function returns large output then it can be hard for someone reading the failure message to find the differences when your test fails. Instead of printing both the returned value and the wanted value, make a diff.

​	如果您的函数返回大量输出，那么当测试失败时，阅读失败消息的人可能很难找到差异。不要同时打印返回值和所需值，而是进行差异化。

To compute diffs for such values, `cmp.Diff` is preferred, particularly for new tests and new code, but other tools may be used. See [types of equality](https://google.github.io/styleguide/go/decisions#types-of-equality) for guidance regarding the strengths and weaknesses of each function.

​	要计算此类值的差异， `cmp.Diff` 是首选，特别是对于新测试和新代码，但可以使用其他工具。有关每个函数的优缺点的指导，请参阅相等类型。

- [`cmp.Diff`](https://pkg.go.dev/github.com/google/go-cmp/cmp#Diff)
- [`pretty.Compare`](https://pkg.go.dev/github.com/kylelemons/godebug/pretty#Compare)

You can use the [`diff`](https://pkg.go.dev/github.com/kylelemons/godebug/diff) package to compare multi-line strings or lists of strings. You can use this as a building block for other kinds of diffs.

​	您可以使用 `diff` 包来比较多行字符串或字符串列表。您可以将其用作其他类型差异的基础。

Add some text to your failure message explaining the direction of the diff.

​	在您的失败消息中添加一些文本，说明差异的方向。

- Something like `diff (-want +got)` is good when you’re using the `cmp`, `pretty`, and `diff` packages (if you pass `(want, got)` to the function), because the `-` and `+` that you add to your format string will match the `-` and `+` that actually appear at the beginning of the diff lines. If you pass `(got, want)` to your function, the correct key would be `(-got +want)` instead.

  ​	当您使用 `cmp` 、 `pretty` 和 `diff` 包时，类似 `diff (-want +got)` 的内容很好（如果您将 `(want, got)` 传递给函数），因为您添加到格式字符串中的 `-` 和 `+` 将匹配实际出现在差异行开头的 `-` 和 `+` 。如果您将 `(got, want)` 传递给您的函数，那么正确的键将是 `(-got +want)` 。

- The `messagediff` package uses a different output format, so the message `diff (want -> got)` is appropriate when you’re using it (if you pass `(want, got)` to the function), because the direction of the arrow will match the direction of the arrow in the “modified” lines.

  ​	由于 `messagediff` 包使用不同的输出格式，因此在使用它时（如果将 `(want, got)` 传递给函数），消息 `diff (want -> got)` 是合适的，因为箭头的方向将与“已修改”行中箭头的方向匹配。

The diff will span multiple lines, so you should print a newline before you print the diff.

​	diff 将跨越多行，因此您应该在打印 diff 之前打印一个换行符。



### Test error semantics 测试错误语义

When a unit test performs string comparisons or uses a vanilla `cmp` to check that particular kinds of errors are returned for particular inputs, you may find that your tests are brittle if any of those error messages are reworded in the future. Since this has the potential to turn your unit test into a change detector (see [TotT: Change-Detector Tests Considered Harmful](https://testing.googleblog.com/2015/01/testing-on-toilet-change-detector-tests.html) ), don’t use string comparison to check what type of error your function returns. However, it is permissible to use string comparisons to check that error messages coming from the package under test satisfy certain properties, for example, that it includes the parameter name.

​	当单元测试执行字符串比较或使用普通的 `cmp` 来检查针对特定输入返回特定类型的错误时，您可能会发现，如果将来对其中任何错误消息进行重新表述，您的测试就会变得脆弱。由于这有可能将您的单元测试变成更改检测器（请参阅 TotT：更改检测器测试被认为有害），因此不要使用字符串比较来检查您的函数返回的错误类型。但是，允许使用字符串比较来检查来自被测包的错误消息是否满足某些属性，例如，它是否包含参数名称。

Error values in Go typically have a component intended for human eyes and a component intended for semantic control flow. Tests should seek to only test semantic information that can be reliably observed, rather than display information that is intended for human debugging, as this is often subject to future changes. For guidance on constructing errors with semantic meaning see [best-practices regarding errors](https://google.github.io/styleguide/go/best-practices#error-handling). If an error with insufficient semantic information is coming from a dependency outside your control, consider filing a bug against the owner to help improve the API, rather than relying on parsing the error message.

​	Go 中的错误值通常包含一个面向人眼的部分和一个面向语义控制流的部分。测试应仅测试可以可靠观察到的语义信息，而不是显示面向人工调试的信息，因为这些信息通常会发生变化。有关构建具有语义含义的错误的指导，请参阅有关错误的最佳做法。如果语义信息不足的错误来自您无法控制的依赖项，请考虑向所有者提交错误报告以帮助改进 API，而不是依赖于解析错误消息。

Within unit tests, it is common to only care whether an error occurred or not. If so, then it is sufficient to only test whether the error was non-nil when you expected an error. If you would like to test that the error semantically matches some other error, then consider using `cmp` with [`cmpopts.EquateErrors`](https://pkg.go.dev/github.com/google/go-cmp/cmp/cmpopts#EquateErrors).

​	在单元测试中，通常只关心是否发生错误。如果是，那么当您预期出现错误时，仅测试错误是否为非 nil 就足够了。如果您想测试错误在语义上是否与其他错误匹配，请考虑将 `cmp` 与 `cmpopts.EquateErrors` 一起使用。

> **Note:** If a test uses [`cmpopts.EquateErrors`](https://pkg.go.dev/github.com/google/go-cmp/cmp/cmpopts#EquateErrors) but all of its `wantErr` values are either `nil` or `cmpopts.AnyError`, then using `cmp` is [unnecessary mechanism](https://google.github.io/styleguide/go/guide#least-mechanism). Simplify the code by making the want field a `bool`. You can then use a simple comparison with `!=`.
>
> ​	注意：如果测试使用 `cmpopts.EquateErrors` 但其所有 `wantErr` 值都是 `nil` 或 `cmpopts.AnyError` ，那么使用 `cmp` 就是不必要的机制。通过使 want 字段成为 `bool` 来简化代码。然后，您可以使用 `!=` 进行简单的比较。
>
> ```
> // Good:
> gotErr := f(test.input) != nil
> if gotErr != test.wantErr {
>     t.Errorf("f(%q) returned err = %v, want error presence = %v", test.input, gotErr, test.wantErr)
> }
> ```

See also [GoTip #13: Designing Errors for Checking](https://google.github.io/styleguide/go/index.html#gotip).

​	另请参阅 GoTip #13：设计用于检查的错误。



## Test structure 测试结构



### Subtests 子测试

The standard Go testing library offers a facility to [define subtests](https://pkg.go.dev/testing#hdr-Subtests_and_Sub_benchmarks). This allows flexibility in setup and cleanup, controlling parallelism, and test filtering. Subtests can be useful (particularly for table-driven tests), but using them is not mandatory. See also the [Go blog post about subtests](https://blog.golang.org/subtests).

​	标准 Go 测试库提供了一个定义子测试的功能。这允许在设置和清理、控制并行性和测试过滤方面具有灵活性。子测试可能很有用（特别是对于表驱动的测试），但使用它们不是强制性的。另请参阅有关子测试的 Go 博客文章。

Subtests should not depend on the execution of other cases for success or initial state, because subtests are expected to be able to be run individually with using `go test -run` flags or with Bazel [test filter](https://bazel.build/docs/user-manual#test-filter) expressions.

​	子测试不应依赖于其他用例的执行才能成功或初始状态，因为子测试应能够使用 `go test -run` 标志或 Bazel 测试筛选器表达式单独运行。



#### Subtest names 子测试名称

Name your subtest such that it is readable in test output and useful on the command line for users of test filtering. When you use `t.Run` to create a subtest, the first argument is used as a descriptive name for the test. To ensure that test results are legible to humans reading the logs, choose subtest names that will remain useful and readable after escaping. Think of subtest names more like a function identifier than a prose description. The test runner replaces spaces with underscores, and escapes non-printing characters. If your test data benefits from a longer description, consider putting the description in a separate field (perhaps to be printed using `t.Log` or alongside failure messages).

​	为您的子测试命名，以便在测试输出中可读，并且对使用测试过滤的用户在命令行中很有用。当您使用 `t.Run` 创建子测试时，第一个参数用作测试的描述性名称。为了确保测试结果对阅读日志的人类来说是可读的，请选择在转义后仍然有用且可读的子测试名称。将子测试名称视为函数标识符，而不是散文描述。测试运行器用下划线替换空格，并转义不可打印的字符。如果您的测试数据受益于更长的描述，请考虑将描述放在单独的字段中（可能使用 `t.Log` 或与失败消息一起打印）。

Subtests may be run individually using flags to the [Go test runner](https://golang.org/cmd/go/#hdr-Testing_flags) or Bazel [test filter](https://bazel.build/docs/user-manual#test-filter), so choose descriptive names that are also easy to type.

​	可以使用 Go 测试运行器或 Bazel 测试过滤器的标志单独运行子测试，因此请选择既具有描述性又易于键入的名称。

> **Warning:** Slash characters are particularly unfriendly in subtest names, since they have [special meaning for test filters](https://blog.golang.org/subtests#:~:text=Perhaps a bit,match any tests).
>
> ​	警告：斜杠字符在子测试名称中特别不友好，因为它们对测试过滤器具有特殊含义。
>
> > ```
> > # Bad:
> > # Assuming TestTime and t.Run("America/New_York", ...)
> > bazel test :mytest --test_filter="Time/New_York"    # Runs nothing!
> > bazel test :mytest --test_filter="Time//New_York"   # Correct, but awkward.
> > ```

To [identify the inputs](https://google.github.io/styleguide/go/decisions#identify-the-input) of the function, include them in the test’s failure messages, where they won’t be escaped by the test runner.

​	要识别函数的输入，请将它们包含在测试的失败消息中，测试运行器不会转义这些消息。

``` go
// Good:
func TestTranslate(t *testing.T) {
    data := []struct {
        name, desc, srcLang, dstLang, srcText, wantDstText string
    }{
        {
            name:        "hu=en_bug-1234",
            desc:        "regression test following bug 1234. contact: cleese",
            srcLang:     "hu",
            srcText:     "cigarettát és egy öngyújtót kérek",
            dstLang:     "en",
            wantDstText: "cigarettes and a lighter please",
        }, // ...
    }
    for _, d := range data {
        t.Run(d.name, func(t *testing.T) {
            got := Translate(d.srcLang, d.dstLang, d.srcText)
            if got != d.wantDstText {
                t.Errorf("%s\nTranslate(%q, %q, %q) = %q, want %q",
                    d.desc, d.srcLang, d.dstLang, d.srcText, got, d.wantDstText)
            }
        })
    }
}
```

Here are a few examples of things to avoid:

​	以下是一些需要避免的事情的示例：

``` go
// Bad:
// Too wordy.
t.Run("check that there is no mention of scratched records or hovercrafts", ...)
// Slashes cause problems on the command line.
t.Run("AM/PM confusion", ...)
```



### Table-driven tests 表格驱动测试

Use table-driven tests when many different test cases can be tested using similar testing logic.

​	当可以使用类似的测试逻辑测试许多不同的测试用例时，使用表驱动测试。

- When testing whether the actual output of a function is equal to the expected output. For example, the many [tests of `fmt.Sprintf`](https://cs.opensource.google/go/go/+/master:src/fmt/fmt_test.go) or the minimal snippet below.
  当测试函数的实际输出是否等于预期输出时。例如， `fmt.Sprintf` 的许多测试或下面的最小片段。
- When testing whether the outputs of a function always conform to the same set of invariants. For example, [tests for `net.Dial`](https://cs.opensource.google/go/go/+/master:src/net/dial_test.go;l=318;drc=5b606a9d2b7649532fe25794fa6b99bd24e7697c).
  当测试函数的输出是否始终符合同一组不变式时。例如， `net.Dial` 的测试。

Here is the minimal structure of a table-driven test, copied from the standard `strings` library. If needed, you may use different names, move the test slice into the test function, or add extra facilities such as subtests or setup and cleanup functions. Always keep [useful test failures](https://google.github.io/styleguide/go/decisions#useful-test-failures) in mind.

​	以下是表驱动测试的最小结构，从标准 `strings` 库中复制而来。如果需要，您可以使用不同的名称，将测试切片移入测试函数，或添加子测试或设置和清理函数等额外功能。始终牢记有用的测试失败。

``` go
// Good:
var compareTests = []struct {
    a, b string
    i    int
}{
    {"", "", 0},
    {"a", "", 1},
    {"", "a", -1},
    {"abc", "abc", 0},
    {"ab", "abc", -1},
    {"abc", "ab", 1},
    {"x", "ab", 1},
    {"ab", "x", -1},
    {"x", "a", 1},
    {"b", "x", -1},
    // test runtime·memeq's chunked implementation
    {"abcdefgh", "abcdefgh", 0},
    {"abcdefghi", "abcdefghi", 0},
    {"abcdefghi", "abcdefghj", -1},
}

func TestCompare(t *testing.T) {
    for _, tt := range compareTests {
        cmp := Compare(tt.a, tt.b)
        if cmp != tt.i {
            t.Errorf(`Compare(%q, %q) = %v`, tt.a, tt.b, cmp)
        }
    }
}
```

**Note**: The failure messages in this example above fulfill the guidance to [identify the function](https://google.github.io/styleguide/go/decisions#identify-the-function) and [identify the input](https://google.github.io/styleguide/go/decisions#identify-the-input). There’s no need to [identify the row numerically](https://google.github.io/styleguide/go/decisions#table-tests-identifying-the-row).

​	注意：此示例中的失败消息满足了识别函数和识别输入的指导。无需按数字识别行。

When some test cases need to be checked using different logic from other test cases, it is more appropriate to write multiple test functions, as explained in [GoTip #50: Disjoint Table Tests](https://google.github.io/styleguide/go/index.html#gotip). The logic of your test code can get difficult to understand when each entry in a table has its own different conditional logic to check each output for its inputs. If test cases have different logic but identical setup, a sequence of [subtests](https://google.github.io/styleguide/go/decisions#subtests) within a single test function might make sense.

​	当某些测试用例需要使用不同于其他测试用例的逻辑进行检查时，编写多个测试函数会更合适，如 GoTip #50：分离表测试中所述。当表中的每个条目都有自己的不同条件逻辑来检查其输入的每个输出时，测试代码的逻辑可能会变得难以理解。如果测试用例具有不同的逻辑但设置相同，则单个测试函数中的一系列子测试可能是有意义的。

You can combine table-driven tests with multiple test functions. For example, when testing that a function’s output exactly matches the expected output and that the function returns a non-nil error for an invalid input, then writing two separate table-driven test functions is the best approach: one for normal non-error outputs, and one for error outputs.

​	您可以将表驱动测试与多个测试函数结合使用。例如，在测试函数的输出与预期输出完全匹配以及该函数为无效输入返回非 nil 错误时，编写两个单独的表驱动测试函数是最佳方法：一个用于正常的无错误输出，另一个用于错误输出。



#### Data-driven test cases 数据驱动的测试用例

Table test rows can sometimes become complicated, with the row values dictating conditional behavior inside the test case. The extra clarity from the duplication between the test cases is necessary for readability.

​	表测试行有时会变得复杂，行值决定了测试用例中的条件行为。为了可读性，测试用例之间的重复提供了额外的清晰度。

``` go
// Good:
type decodeCase struct {
    name   string
    input  string
    output string
    err    error
}

func TestDecode(t *testing.T) {
    // setupCodex is slow as it creates a real Codex for the test.
    codex := setupCodex(t)

    var tests []decodeCase // rows omitted for brevity

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            output, err := Decode(test.input, codex)
            if got, want := output, test.output; got != want {
                t.Errorf("Decode(%q) = %v, want %v", test.input, got, want)
            }
            if got, want := err, test.err; !cmp.Equal(got, want) {
                t.Errorf("Decode(%q) err %q, want %q", test.input, got, want)
            }
        })
    }
}

func TestDecodeWithFake(t *testing.T) {
    // A fakeCodex is a fast approximation of a real Codex.
    codex := newFakeCodex()

    var tests []decodeCase // rows omitted for brevity

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            output, err := Decode(test.input, codex)
            if got, want := output, test.output; got != want {
                t.Errorf("Decode(%q) = %v, want %v", test.input, got, want)
            }
            if got, want := err, test.err; !cmp.Equal(got, want) {
                t.Errorf("Decode(%q) err %q, want %q", test.input, got, want)
            }
        })
    }
}
```

In the counterexample below, note how hard it is to distinguish between which type of `Codex` is used per test case in the case setup. (The highlighted parts run afoul of the advice from [TotT: Data Driven Traps!](https://testing.googleblog.com/2008/09/tott-data-driven-traps.html) .)

​	在下面的反例中，请注意区分在用例设置中每个测试用例使用的 `Codex` 类型有多难。（突出显示的部分违背了 TotT 的建议：数据驱动陷阱！）

``` go
// Bad:
type decodeCase struct {
  name   string
  input  string
  codex  testCodex
  output string
  err    error
}

type testCodex int

const (
  fake testCodex = iota
  prod
)

func TestDecode(t *testing.T) {
  var tests []decodeCase // rows omitted for brevity

  for _, test := tests {
    t.Run(test.name, func(t *testing.T) {
      var codex Codex
      switch test.codex {
      case fake:
        codex = newFakeCodex()
      case prod:
        codex = setupCodex(t)
      default:
        t.Fatalf("unknown codex type: %v", codex)
      }
      output, err := Decode(test.input, codex)
      if got, want := output, test.output; got != want {
        t.Errorf("Decode(%q) = %q, want %q", test.input, got, want)
      }
      if got, want := err, test.err; !cmp.Equal(got, want) {
        t.Errorf("Decode(%q) err %q, want %q", test.input, got, want)
      }
    })
  }
}
```



#### Identifying the row 识别行

Do not use the index of the test in the test table as a substitute for naming your tests or printing the inputs. Nobody wants to go through your test table and count the entries in order to figure out which test case is failing.

​	不要使用测试表中测试的索引作为测试的命名或打印输入的替代。没有人愿意通读您的测试表并计算条目，以找出哪个测试用例失败。

``` go
// Bad:
tests := []struct {
    input, want string
}{
    {"hello", "HELLO"},
    {"wORld", "WORLD"},
}
for i, d := range tests {
    if strings.ToUpper(d.input) != d.want {
        t.Errorf("failed on case #%d", i)
    }
}
```

Add a test description to your test struct and print it along failure messages. When using subtests, your subtest name should be effective in identifying the row.

​	向测试结构添加测试说明，并将其与失败消息一起打印。使用子测试时，子测试名称应有效识别行。

**Important:** Even though `t.Run` scopes the output and execution, you must always [identify the input](https://google.github.io/styleguide/go/decisions#identify-the-input). The table test row names must follow the [subtest naming](https://google.github.io/styleguide/go/decisions#subtest-names) guidance.

​	重要提示：即使 `t.Run` 限定了输出和执行，您也必须始终识别输入。表测试行名称必须遵循子测试命名指南。



### Test helpers 测试助手

A test helper is a function that performs a setup or cleanup task. All failures that occur in test helpers are expected to be failures of the environment (not from the code under test) — for example when a test database cannot be started because there are no more free ports on this machine.

​	测试助手是一个执行设置或清理任务的函数。在测试助手程序中发生的所有故障都应是环境故障（而不是被测代码的故障）——例如，当无法启动测试数据库时，因为此计算机上没有更多可用端口。

If you pass a `*testing.T`, call [`t.Helper`](https://pkg.go.dev/testing#T.Helper) to attribute failures in the test helper to the line where the helper is called. This parameter should come after a [context](https://google.github.io/styleguide/go/decisions#contexts) parameter, if present, and before any remaining parameters.

​	如果传递 `*testing.T` ，请调用 `t.Helper` 将测试助手中的失败归因于调用助手的行。此参数应在上下文参数（如果存在）之后，在任何剩余参数之前。

``` go
// Good:
func TestSomeFunction(t *testing.T) {
    golden := readFile(t, "testdata/golden-result.txt")
    // ... tests against golden ...
}

// readFile returns the contents of a data file.
// It must only be called from the same goroutine as started the test.
func readFile(t *testing.T, filename string) string {
    t.Helper()
    contents, err := runfiles.ReadFile(filename)
    if err != nil {
        t.Fatal(err)
    }
    return string(contents)
}
```

Do not use this pattern when it obscures the connection between a test failure and the conditions that led to it. Specifically, the guidance about [assert libraries](https://google.github.io/styleguide/go/decisions#assert) still applies, and [`t.Helper`](https://pkg.go.dev/testing#T.Helper) should not be used to implement such libraries.

​	当它模糊了测试失败与导致失败的条件之间的联系时，请勿使用此模式。具体来说，关于断言库的指导仍然适用，并且不应使用 `t.Helper` 来实现此类库。

**Tip:** For more on the distinction between test helpers and assertion helpers, see [best practices](https://google.github.io/styleguide/go/best-practices#test-functions).

​	提示：有关测试助手和断言助手之间区别的更多信息，请参阅最佳做法。

Although the above refers to `*testing.T`, much of the advice stays the same for benchmark and fuzz helpers.

​	尽管上述内容是指 `*testing.T` ，但许多建议对基准和模糊助手来说仍然相同。



### Test package 测试包





#### Tests in the same package 同一包中的测试

Tests may be defined in the same package as the code being tested.

​	可以在与被测代码相同的包中定义测试。

To write a test in the same package:

​	要在同一包中编写测试，请执行以下操作：

- Place the tests in a `foo_test.go` file
  将测试放在 `foo_test.go` 文件中
- Use `package foo` for the test file
  对测试文件使用 `package foo`
- Do not explicitly import the package to be tested
  不要显式导入要测试的包

```build
// Good:
go_library(
    name = "foo",
    srcs = ["foo.go"],
    deps = [
        ...
    ],
)

go_test(
    name = "foo_test",
    size = "small",
    srcs = ["foo_test.go"],
    library = ":foo",
    deps = [
        ...
    ],
)
```

A test in the same package can access unexported identifiers in the package. This may enable better test coverage and more concise tests. Be aware that any [examples](https://google.github.io/styleguide/go/decisions#examples) declared in the test will not have the package names that a user will need in their code.

​	同一包中的测试可以访问包中未导出的标识符。这可能会实现更好的测试覆盖率和更简洁的测试。请注意，在测试中声明的任何示例都不会具有用户在其代码中需要的包名。



#### Tests in a different package 不同包中的测试

It is not always appropriate or even possible to define a test in the same package as the code being tested. In these cases, use a package name with the `_test` suffix. This is an exception to the “no underscores” rule to [package names](https://google.github.io/styleguide/go/decisions#package-names). For example:

​	将测试定义在与被测代码相同的包中并不总是适当的，甚至是不可能的。在这些情况下，请使用带有 `_test` 后缀的包名。这是包名“无下划线”规则的例外。例如：

- If an integration test does not have an obvious library that it belongs to

  ​	如果集成测试没有它所属的明显库

  ```
  // Good:
  package gmailintegration_test
  
  import "testing"
  ```

- If defining the tests in the same package results in circular dependencies

  ​	如果在同一包中定义测试会导致循环依赖

  ```
  // Good:
  package fireworks_test
  
  import (
    "fireworks"
    "fireworkstestutil" // fireworkstestutil also imports fireworks
  )
  ```



### Use package `testing` 使用包 `testing`

The Go standard library provides the [`testing` package](https://pkg.go.dev/testing). This is the only testing framework permitted for Go code in the Google codebase. In particular, [assertion libraries](https://google.github.io/styleguide/go/decisions#assert) and third-party testing frameworks are not allowed.

​	Go 标准库提供了 `testing` 包。这是 Google 代码库中 Go 代码唯一允许的测试框架。特别是，不允许使用断言库和第三方测试框架。

The `testing` package provides a minimal but complete set of functionality for writing good tests:

​	 `testing` 包提供了一组最少但完整的编写良好测试的功能：

- Top-level tests
  顶级测试
- Benchmarks
  基准
- [Runnable examples
  可运行示例](https://blog.golang.org/examples)
- Subtests
  子测试
- Logging
  日志记录
- Failures and fatal failures
  失败和致命失败

These are designed to work cohesively with core language features like [composite literal](https://go.dev/ref/spec#Composite_literals) and [if-with-initializer](https://go.dev/ref/spec#If_statements) syntax to enable test authors to write [clear, readable, and maintainable tests].

​	这些旨在与复合字面量和 if-with-initializer 语法等核心语言特性协同工作，使测试编写者能够编写[清晰、可读且可维护的测试]。



## Non-decisions 未做决定

A style guide cannot enumerate positive prescriptions for all matters, nor can it enumerate all matters about which it does not offer an opinion. That said, here are a few things where the readability community has previously debated and has not achieved consensus about.

​	风格指南无法对所有事项列举积极的规定，也无法列举所有它不提供意见的事项。话虽如此，这里有一些事情是可读性社区以前争论过但尚未达成共识的。

- **Local variable initialization with zero value**. `var i int` and `i := 0` are equivalent. See also [initialization best practices](https://google.github.io/styleguide/go/best-practices#vardeclinitialization).
  用零值初始化局部变量。 `var i int` 和 `i := 0` 是等价的。另请参阅初始化最佳做法。
- **Empty composite literal vs. `new` or `make`**. `&File{}` and `new(File)` are equivalent. So are `map[string]bool{}` and `make(map[string]bool)`. See also [composite declaration best practices](https://google.github.io/styleguide/go/best-practices#vardeclcomposite).
  空复合字面量与 `new` 或 `make` 相比。 `&File{}` 和 `new(File)` 是等价的。同样， `map[string]bool{}` 和 `make(map[string]bool)` 也是等价的。另请参阅复合声明最佳做法。
- **got, want argument ordering in cmp.Diff calls**. Be locally consistent, and [include a legend](https://google.github.io/styleguide/go/decisions#print-diffs) in your failure message.
  got, 希望在 cmp.Diff 调用中对参数进行排序。保持本地一致性，并在您的失败消息中包含一个图例。
- **`errors.New` vs `fmt.Errorf` on non-formatted strings**. `errors.New("foo")` and `fmt.Errorf("foo")` may be used interchangeably.
  `errors.New` 与 `fmt.Errorf` 在非格式化字符串上。 `errors.New("foo")` 和 `fmt.Errorf("foo")` 可以互换使用。

If there are special circumstances where they come up again, the readability mentor might make an optional comment, but in general the author is free to pick the style they prefer in the given situation.

​	如果有特殊情况再次出现，可读性导师可能会发表可选评论，但通常作者可以自由选择在给定情况下他们喜欢的风格。

Naturally, if anything not covered by the style guide does need more discussion, authors are welcome to ask – either in the specific review, or on internal message boards.

​	当然，如果风格指南未涵盖的任何内容确实需要更多讨论，作者可以随时提问——无论是在特定评论中还是在内部留言板上。