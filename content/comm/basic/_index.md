+++
title = "基础"
date = 2024-03-01T15:18:46+08:00
weight = -100
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

给出Go、Python、Java、Rust、C/C++、JavaScript、TypeScript、C#、Erlang、PHP、Ruby各个编程语言各自的



## 安装、配置、卸载、更新、运行代码

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}

​	详见：[/comm/Go/basic/installs]({{< ref "/comm/Go/basic/installs" >}})

{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 关键字、保留字

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}

​	Go目前有25个关键字。

来源：[https://go.dev/ref/spec#Keywords](https://go.dev/ref/spec#Keywords)

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```



{{% /tab  %}}

{{% tab header="Python" %}}

​	Python目前有35个关键字。

​	来源：[https://docs.python.org/zh-cn/3.12/reference/lexical_analysis.html#keywords](https://docs.python.org/zh-cn/3.12/reference/lexical_analysis.html#keywords)

```python
False      await      else       import     pass
None       break      except     in         raise
True       class      finally    is         return
and        continue   for        lambda     try
as         def        from       nonlocal   while
assert     del        global     not        with
async      elif       if         or         yield
```

{{% /tab  %}}

{{% tab header="Java" %}}

来源：[https://docs.oracle.com/javase/specs/jls/se21/html/jls-3.html#jls-3.9](https://docs.oracle.com/javase/specs/jls/se21/html/jls-3.html#jls-3.9)

*ReservedKeyword:*

```java
abstract   continue   for          new         switch
assert     default    if           package     synchronized
boolean    do         goto         private     this
break      double     implements   protected   throw
byte       else       import       public      throws
case       enum       instanceof   return      transient
catch      extends    int          short       try
char       final      interface    static      void
class      finally    long         strictfp    volatile
const      float      native       super       while
_ (underscore)
```

*ContextualKeyword:*

```java
exports      opens      requires     uses   yield
module       permits    sealed       var         
non-sealed   provides   to           when        
open         record     transitive   with    
```

{{% /tab  %}}

{{% tab header="Rust" %}}

来源：[https://doc.rust-lang.org/reference/keywords.html](https://doc.rust-lang.org/reference/keywords.html)



{{% /tab  %}}

{{% tab header="C/C++" %}}

来源：[https://en.cppreference.com/w/cpp/keyword](https://en.cppreference.com/w/cpp/keyword)



{{% /tab  %}}

{{% tab header="JavaScript" %}}

来源：[https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar#keywords](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar#keywords)

{{% /tab  %}}

{{% tab header="TypeScript" %}}

来源：

{{% /tab  %}}

{{% tab header="C#" %}}

来源：

{{% /tab  %}}

{{% tab header="Erlang" %}}

来源：

{{% /tab  %}}

{{% tab header="PHP" %}}

来源：[https://www.php.net/manual/en/reserved.keywords.php](https://www.php.net/manual/en/reserved.keywords.php)

```php
__halt_compiler()	abstract	and	array()	as
break	callable	case	catch	class
clone	const	continue	declare	default
die()	do	echo	else	elseif
empty()	enddeclare	endfor	endforeach	endif
endswitch	endwhile	eval()	exit()	extends
final	finally	fn (as of PHP 7.4)	for	foreach
function	global	goto	if	implements
include	include_once	instanceof	insteadof	interface
isset()	list()	match (as of PHP 8.0)	namespace	new
or	print	private	protected	public
readonly (as of PHP 8.1.0) *	require	require_once	return	static
switch	throw	trait	try	unset()
use	var	while	xor	yield
yield from
```

{{% /tab  %}}

{{% tab header="Ruby" %}}

来源：

{{% /tab  %}}

{{< /tabpane >}}



## 查找可用模块

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}

- [https://go.libhunt.com](https://go.libhunt.com/)：一个 Go 语言库的集合，方便开发者查找和比较不同的库。

- [https://pkg.go.dev](https://pkg.go.dev/)：官方

{{% /tab  %}}

{{% tab header="Python" %}}

- [https://pypi.org](https://pypi.org/)：Python 的官方包索引，包含了成千上万的第三方库和模块。

- [https://awesome-python.com](https://awesome-python.com)：一个精选的 Python 库和资源列表，涵盖了各个领域的工具和库。

- [https://docs.python.org/zh-cn](https://docs.python.org/zh-cn/3.12/library/index.html)：Python 标准库的官方文档，详细介绍了 Python 内置的模块和功能。

{{% /tab  %}}

{{% tab header="Java" %}}

- [https://central.sonatype.com](https://central.sonatype.com/ ) 
- [https://java.libhunt.com](https://java.libhunt.com/)

{{% /tab  %}}

{{% tab header="Rust" %}}

- [https://crates.io](https://crates.io/)：Rust 的官方包注册中心，包含了众多的 Rust 库和工具。

- [https://rust.libhunt.com](https://rust.libhunt.com/)：一个精选的 Rust 库和资源列表，涵盖了各个领域的工具和库。

{{% /tab  %}}

{{% tab header="C/C++" %}}

- [GitHub](https://github.com/)：全球最大的代码托管平台，包含了大量的 C 和 C++ 项目和库。

- [https://sourceforge.net](https://sourceforge.net/)：一个老牌的开源软件开发平台，提供了众多的 C/C++ 项目。

- [https://faraz.work/awesome-cpp](https://faraz.work/awesome-cpp/)：一个精选的 C++ 库和资源列表，涵盖了各个领域的工具和库。

- [https://cpp.libhunt.com](https://cpp.libhunt.com/)

{{% /tab  %}}

{{% tab header="JavaScript" %}}

- [https://www.npmjs.com](https://www.npmjs.com/)：JavaScript 的包管理器，包含了大量的前端和后端库。

- [https://js.libhunt.com](https://js.libhunt.com/)

{{% /tab  %}}

{{% tab header="TypeScript" %}}

- [https://www.libhunt.com/l/typescript](https://www.libhunt.com/l/typescript)

{{% /tab  %}}

{{% tab header="C#" %}}

- [https://www.nuget.org](https://www.nuget.org/)：C# 的官方包管理器，包含了大量的 .NET 库和工具。
- [https://www.libhunt.com/l/c-sharp](https://www.libhunt.com/l/c-sharp)

{{% /tab  %}}

{{% tab header="Erlang" %}}

- [https://hex.pm](https://hex.pm/)：Erlang 和 Elixir 的包管理器，提供了众多的库和工具。
- [https://www.libhunt.com/l/erlang](https://www.libhunt.com/l/erlang)

{{% /tab  %}}

{{% tab header="PHP" %}}

- [https://packagist.org](https://packagist.org/)：PHP 的官方包仓库，包含了大量的 PHP 库和框架。
- [https://php.libhunt.com](https://php.libhunt.com/)



{{% /tab  %}}

{{% tab header="Ruby" %}}

- [https://rubygems.org](https://rubygems.org/)：Ruby 的官方包管理器，提供了众多的 Ruby 库和工具。
- [https://ruby.libhunt.com](https://ruby.libhunt.com/)



{{% /tab  %}}

{{< /tabpane >}}

## 模块管理工具

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}

​	Go 语言的模块管理工具是 **Go Modules**，自 Go 1.11 版本起引入，并在 Go 1.13 版本后成为官方推荐的标准。

**Go Modules 的核心文件**

- **go.mod：** 定义了模块的名称、Go 版本以及项目的依赖关系（可以通过`go mod init <模块路径>`自动生成文件，并可通过`go mod tidy`、`go mod edit`、`go get`等命令进行编辑）。
- **go.sum：** 记录了每个依赖包的哈希值，用于校验依赖包的完整性（通过`go mod tidy`或`go build`等命令自动生成）。

**常用的 Go Modules 命令**

| 命令                 | 描述                                                     | 示例                                                  |
| -------------------- | -------------------------------------------------------- | ----------------------------------------------------- |
| `go mod init`        | 初始化一个新的模块，创建 `go.mod` 文件。                 | `go mod init github.com/username/projectname`         |
| `go mod tidy`        | 清理 `go.mod` 和 `go.sum` 文件，移除不再需要的依赖项。   | `go mod tidy`                                         |
| `go mod vendor`      | 将所有依赖项复制到 `vendor` 目录，方便离线构建。         | `go mod vendor`                                       |
| `go get`             | 下载并安装指定的包及其依赖项。                           | `go get github.com/some/package`                      |
| `go get -u`          | 更新指定的包及其依赖项到最新版本。                       | `go get -u github.com/some/package`                   |
| `go get -d`          | 仅下载指定的包及其依赖项，不进行安装。                   | `go get -d github.com/some/package`                   |
| `go list -m all`     | 列出当前模块的所有依赖项及其版本。                       | `go list -m all`                                      |
| `go mod verify`      | 校验 `go.mod` 和 `go.sum` 文件中的依赖项是否与实际一致。 | `go mod verify`                                       |
| `go clean -modcache` | 清理模块缓存，释放磁盘空间。                             | `go clean -modcache`                                  |
| `go build`           | 编译当前模块的所有包。                                   | `go build`                                            |
| `go test`            | 运行当前模块的所有测试。                                 | `go test`                                             |
| `go run`             | 编译并运行指定的 Go 源文件。                             | `go run main.go`                                      |
| `go install`         | 编译并安装指定的包或可执行文件。                         | `go install github.com/some/package`                  |
| `go mod graph`       | 打印模块的依赖图。                                       | `go mod graph`                                        |
| `go mod edit`        | 编辑 `go.mod` 文件，例如添加、删除或修改依赖项。         | `go mod edit -require=github.com/some/package@v1.2.3` |
| `go mod download`    | 下载 `go.mod` 中列出的所有依赖项到本地缓存。             | `go mod download`                                     |
| `go mod why`         | 解释为什么需要某个包或模块。                             | `go mod why github.com/some/package`                  |
| `go mod graph`       | 打印模块的依赖图。                                       | `go mod graph`                                        |
| `go mod edit`        | 编辑 `go.mod` 文件，例如添加、删除或修改依赖项。         | `go mod edit -require=github.com/some/package@v1.2.3` |
| `go mod vendor`      | 将所有依赖项复制到 `vendor` 目录，方便离线构建。         | `go mod vendor`                                       |

{{% /tab  %}}

{{% tab header="Python" %}}

​	`pip`是Python自带的，无需安装，若有新版本只需更新。

**`pip`默认的配置文件**

（1）用户级配置文件

| 操作系统        | 配置文件路径                                                 | 说明                                                         |
| --------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **Linux/macOS** | `~/.config/pip/pip.conf`                                     | 新版本 pip 默认遵循 XDG 规范，优先使用此路径。旧路径 `~/.pip/pip.conf` 仍兼容，但已不推荐。 |
| **Windows**     | `%APPDATA%\pip\pip.ini`（即 `C:\Users\<用户名>\AppData\Roaming\pip\pip.ini`）（`%APPDATA%`表示环境变量：`$env:APPDATA`，其值为：`C:\Users\用户名\AppData\Roaming`） | 需显示隐藏文件夹才能看到 `AppData` 目录。                    |

（2）全局配置文件

| 操作系统        | 配置文件路径                 | 说明                                                       |
| --------------- | ---------------------------- | ---------------------------------------------------------- |
| **Linux/macOS** | `/etc/pip.conf`              | 需要管理员权限才能修改。                                   |
| **Windows**     | `C:\ProgramData\pip\pip.ini` | `ProgramData` 是隐藏目录，需开启 “显示隐藏文件” 才能访问。 |

**`pip`配置文件优先级**

​	`pip` 加载配置的顺序为（后者覆盖前者）：

1. **全局配置** → 2. **用户配置** → 3. **虚拟环境配置**（若有） → 4. **命令行参数**。



**`pip`的更新**

```sh
python  -m pip install --upgrade pip

# 以下是示例
PS C:\Windows\System32> pip -V
pip 24.3.1 from D:\tools\Python312\Lib\site-packages\pip (python 3.12)
PS C:\Windows\System32> python -m pip install --upgrade pip
Looking in indexes: https://mirrors.aliyun.com/pypi/simple/
Requirement already satisfied: pip in d:\tools\python312\lib\site-packages (24.3.1)
Collecting pip
  Downloading https://mirrors.aliyun.com/pypi/packages/c9/bc/b7db44f5f39f9d0494071bddae6880eb645970366d0a200022a1a93d57f5/pip-25.0.1-py3-none-any.whl (1.8 MB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 1.8/1.8 MB 11.2 MB/s eta 0:00:00
Installing collected packages: pip
  Attempting uninstall: pip
    Found existing installation: pip 24.3.1
    Uninstalling pip-24.3.1:
      Successfully uninstalled pip-24.3.1
Successfully installed pip-25.0.1
PS C:\Windows\System32>
```

**`pip`在国内常见可用的镜像源**

| 镜像名称     | URL                                                   |
| ------------ | ----------------------------------------------------- |
| **清华大学** | `https://pypi.tuna.tsinghua.edu.cn/simple`            |
| **阿里云**   | `https://mirrors.aliyun.com/pypi/simple`              |
| **腾讯云**   | `https://mirrors.cloud.tencent.com/pypi/simple`       |
| **华为云**   | `https://repo.huaweicloud.com/repository/pypi/simple` |

**`pip`的常用命令**

（1）查看包的可用版本

```sh
# 查看 numpy的可用版本（仅限 pip 23.0+）
PS D:\Docs\hugos\go_docs> pip index versions numpy
WARNING: pip index is currently an experimental command. It may be removed/changed in a future release without prior warning.
numpy (2.2.2)
Available versions: 2.2.2, 2.2.1, 2.2.0, 2.1.3, 2.1.2, 2.1.1, 2.1.0, 2.0.2, 2.0.1, 2.0.0, 1.26.4, 1.26.3, 1.26.2, 1.26.1, 1.26.0, 1.25.2, 1.25.1, 1.25.0, 1.24.4, 1.24.3, 1.24.2, 1.24.1, 1.24.0, 1.23.5, 1.23.4, 1.23.3, 1.23.2, 1.23.1, 1.23.0, 1.22.4, 1.22.3, 1.22.2, 1.22.1, 1.22.0, 1.21.1, 1.21.0, 1.20.3, 1.20.2, 1.20.1, 1.20.0, 1.19.5, 1.19.4, 1.19.3, 1.19.2, 1.19.1, 1.19.0, 1.18.5, 1.18.4, 1.18.3, 1.18.2, 1.18.1, 1.18.0, 1.17.5, 1.17.4, 1.17.3, 1.17.2, 1.17.1, 1.17.0, 1.16.6, 1.16.5, 1.16.4, 1.16.3, 1.16.2, 1.16.1, 1.16.0, 1.15.4, 1.15.3, 1.15.2, 1.15.1, 1.15.0, 1.14.6, 1.14.5, 1.14.4, 1.14.3, 1.14.2, 1.14.1, 1.14.0, 1.13.3, 1.13.1, 1.13.0, 1.12.1, 1.12.0, 1.11.3, 1.11.2, 1.11.1, 1.11.0, 1.10.4, 1.10.2, 1.10.1, 1.10.0.post2, 1.9.3, 1.9.2, 1.9.1, 1.9.0, 1.8.2, 1.8.1, 1.8.0, 1.7.2, 1.7.1, 1.7.0, 1.6.2, 1.6.1, 1.6.0, 1.5.1, 1.5.0, 1.4.1, 1.3.0
  INSTALLED: 2.2.2
  LATEST:    2.2.2
  
  
  # 也可以通过如下方式，获取numpy的最新可用版本（不实际安装）
PS D:\Docs\hugos\go_docs> pip install --upgrade numpy --dry-run
Looking in indexes: https://mirrors.aliyun.com/pypi/simple/
Collecting numpy
  Downloading https://mirrors.aliyun.com/pypi/packages/fc/84/7f801a42a67b9772a883223a0a1e12069a14626c81a732bd70aac57aebc1/numpy-2.2.2-cp312-cp312-win_amd64.whl (12.6 MB)
     ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━ 12.6/12.6 MB 31.6 MB/s eta 0:00:00
Would install numpy-2.2.2
  
  
```

（2）安装包

a.基本安装

```sh
# 安装最新版本的 requests
pip install requests

# 安装指定版本的 numpy
pip install numpy==1.21.0

# 安装本地 .whl 文件
pip install ./downloads/pandas-2.0.3-cp39-cp39-win_amd64.whl

#.whl 文件名需符合 PEP 427 的命名规范，格式为：
# {包名}-{版本}-{Python标签}-{ABI标签}-{平台标签}.whl
# 例如：pandas-2.0.3-cp39-cp39-win_amd64.whl
# Python 标签：cp39 表示适用于 Python 3.9 的 CPython 解释器。
# ABI 标签：cp39 表示与 Python 3.9 ABI 兼容。
# 平台标签：win_amd64 表示 Windows 64 位 系统。
# .whl 文件本质是一个 ZIP 压缩包，
```

b.从 Git 仓库安装

```sh
# 安装 GitHub 上的开发版库
pip install git+https://github.com/user/repo.git@branch
```

c.使用镜像加速

```sh
# 使用清华镜像源安装
pip install tensorflow -i https://pypi.tuna.tsinghua.edu.cn/simple

# 信任 HTTP 镜像源（如阿里云旧版）
pip install flask -i http://mirrors.aliyun.com/pypi/simple --trusted-host mirrors.aliyun.com
```

d.批量安装依赖

```sh
pip install -r requirements.txt
# 每行包含一个包名，可以指定版本号，也可以不指定。
# 可以使用 ==、>=、<=、>、< 等符号来指定版本范围。
# 可以使用 -e . 来安装本地项目。
# 可以使用 # 来添加注释，注释内容会被忽略。
```

（3）升级包

```sh
# 升级 requests 到最新版
pip install --upgrade requests

# 简写形式
pip install -U requests
```

（4）卸载包

```sh
#  卸载 pandsa
pip uninstall pandas
```

（5）依赖管理

```sh
# 导出当前环境所有依赖
pip freeze > requirements.txt

# 仅导出项目直接依赖（需配合 pip-tools）
pip install pip-tools
# Compiles requirements.txt from requirements.in, pyproject.toml, setup.cfg,  or setup.py specs.
# 注意不用使用 > requirements.txt，来重定向到requirements.txt
pip-compile requirements.in
# pip-compile 是 pip-tools 工具集提供的命令，安装 pip-tools 后才能使用。
# requirements.in（该文件必须存在）是一个手动维护的输入文件，仅包含项目的直接依赖（即你显式安装的包）

# 修改requirements.in文件后，重新生成新的requirements.txt
# 注意不用使用 > requirements.txt，来重定向到requirements.txt
pip-compile --upgrade requirements.in
```

（6）列出已安装的包

```sh
# 列出所有包
pip list

# 列出过期的包
pip list --outdated
```

（7）查看包详情

```sh
# 列出所有包
pip list

# 列出过期的包
pip list --outdated
```

（8）设置全局镜像源

```sh
# 永久配置清华镜像（Linux/macOS）
pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple

```

（9）查看当前配置

```sh
pip config list
```

（10）安装开发模式（可编辑模式）

```sh
# 本地开发时直接链接代码（无需重复安装）
pip install -e ./my_package
```

（11）忽略缓存安装

```sh
# 避免使用旧缓存文件
pip install --no-cache-dir torch
```

（12）限制超时时间

```sh
# 设置超时为 60 秒
pip install --timeout 60 scipy
```



{{% /tab  %}}

{{% tab header="Java" %}}

Maven 是 Java 项目的构建、依赖管理和项目管理工具，提供了丰富的命令来支持项目的各个阶段。

**Maven的安装**

1. 下载 Maven

- 访问 Maven 官网：https://maven.apache.org/download.cgi
- 在 "Files" 部分，找到 "Binary Distributions" 下的最新版本（例如 apache-maven-3.9.2-bin.zip），点击下载。

2. 解压 Maven

- 将下载的压缩包解压到你想要安装 Maven 的目录，例如：`C:\Program Files\Apache\apache-maven-3.9.2`。

3. 配置环境变量

- **设置 MAVEN_HOME 变量：**
  - 在 Windows 搜索栏中输入 "环境变量"，选择 "编辑系统环境变量"。
  - 点击 "环境变量" 按钮。
  - 在 "系统变量" 部分，点击 "新建"。
  - 变量名输入 `MAVEN_HOME`，变量值输入 Maven 的安装目录，例如 `C:\Program Files\Apache\apache-maven-3.9.2`。
  - 点击 "确定"。
- **将 Maven 的 bin 目录添加到 Path 变量：**
  - 在 "系统变量" 部分，找到名为 "Path" 的变量，选中并点击 "编辑"。
  - 点击 "新建"，输入 `%MAVEN_HOME%\bin`。
  - 点击 "确定"。

4. 验证安装

- 打开命令提示符（CMD）或 PowerShell。
- 输入 `mvn -v` 命令，按回车键。
- 如果显示 Maven 的版本信息，则表示安装成功。

5. 配置 Maven（可选）

- **设置本地仓库：**

  - 在 Maven 安装目录下，找到 `conf` 文件夹，打开 `settings.xml` 文件。

  - 找到 `<localRepository>` 标签，取消注释，并设置本地仓库路径，例如：

    ```xml
    <localRepository>D:/maven/repository</localRepository>
    ```

  - 如果没有 `repository` 文件夹，需要手动创建。

- **配置阿里云镜像（可选，加速下载）：**

  详细请参见：[https://developer.aliyun.com/mvn/guide](https://developer.aliyun.com/mvn/guide)

  - 在 `settings.xml` 文件中，找到 `<mirrors>` 标签，添加以下内容：

    ```xml
    <mirror>
      <id>aliyunmaven</id>
      <mirrorOf>*</mirrorOf>
      <name>阿里云公共仓库</name>
      <url>https://maven.aliyun.com/repository/public</url>
    </mirror>
    ```

**常用的 Maven 命令及其示例**

| 命令                               | 描述                                         | 示例                                                         |
| ---------------------------------- | -------------------------------------------- | ------------------------------------------------------------ |
| `mvn clean`                        | 清理项目，删除 `target` 目录下的所有生成物。 | `mvn clean`                                                  |
| `mvn compile`                      | 编译项目的源代码。                           | `mvn compile`                                                |
| `mvn test`                         | 运行项目的单元测试。                         | `mvn test`                                                   |
| `mvn package`                      | 打包项目，生成 JAR 或 WAR 文件。             | `mvn package`                                                |
| `mvn install`                      | 将打包后的文件安装到本地仓库。               | `mvn install`                                                |
| `mvn deploy`                       | 将打包后的文件部署到远程仓库。               | `mvn deploy`                                                 |
| `mvn validate`                     | 验证项目是否正确，检查项目的有效性。         | `mvn validate`                                               |
| `mvn verify`                       | 运行集成测试，验证项目是否满足质量标准。     | `mvn verify`                                                 |
| `mvn clean install`                | 清理项目并安装到本地仓库。                   | `mvn clean install`                                          |
| `mvn clean package`                | 清理项目并打包。                             | `mvn clean package`                                          |
| `mvn clean deploy`                 | 清理项目并部署到远程仓库。                   | `mvn clean deploy`                                           |
| `mvn site`                         | 生成项目相关信息的网站。                     | `mvn site`                                                   |
| `mvn dependency:tree`              | 打印项目的依赖关系树。                       | `mvn dependency:tree`                                        |
| `mvn dependency:list`              | 列出项目的所有依赖。                         | `mvn dependency:list`                                        |
| `mvn dependency:copy-dependencies` | 将项目的依赖复制到指定目录。                 | `mvn dependency:copy-dependencies -DoutputDirectory=libs`    |
| `mvn archetype:generate`           | 创建一个新的 Maven 项目。                    | `mvn archetype:generate -DgroupId=com.example -DartifactId=my-app -DarchetypeArtifactId=maven-archetype-quickstart` |
| `mvn clean validate`               | 清理项目并验证项目的有效性。                 | `mvn clean validate`                                         |
| `mvn clean compile`                | 清理项目并编译源代码。                       | `mvn clean compile`                                          |
| `mvn clean test`                   | 清理项目并运行单元测试。                     | `mvn clean test`                                             |
| `mvn clean package`                | 清理项目并打包。                             | `mvn clean package`                                          |
| `mvn clean install`                | 清理项目并安装到本地仓库。                   | `mvn clean install`                                          |
| `mvn clean deploy`                 | 清理项目并部署到远程仓库。                   | `mvn clean deploy`                                           |





{{% /tab  %}}

{{% tab header="Rust" %}}

**内联模块**

​	直接在文件中用 `mod module_name { ... }` 定义。模块可以包含其他模块，形成层次结构，即模块可以嵌套模块。

**文件模块**

- 文件 `src/my_module.rs` 自动成为模块 `my_module`。
- 目录 `src/my_module/` 需包含 `mod.rs` 文件作为入口。

**使用模块**

​	使用 `use` 关键字引入模块或其成员。

```rust
use my_module::my_function;

fn main() {
    my_function();
}
```

**Crate：Rust 的包和构建单元**

​	在 Rust 中，`crate` 是编译单元，可以是一个库或一个二进制可执行文件。每个 crate 都有一个根文件，通常是 `src/main.rs`（对于二进制 crate）或 `src/lib.rs`（对于库 crate）。

- **库 crate**：如果项目包含 `src/lib.rs` 文件，则该项目是一个库 crate。
- **二进制 crate**：如果项目包含 `src/main.rs` 文件，则该项目是一个二进制 crate。
- **多二进制 crate**：如果项目的 `src` 目录下包含多个源文件，Cargo 会将它们视为多个二进制 crate。

**Crate 根的具体规则**

1. **二进制 crate**（生成可执行文件）：
   - crate 根是 `src/main.rs` 文件。
   - 例如：当你运行 `cargo new my_project`，默认生成的 `src/main.rs` 就是这个 crate 的根。
2. **库 crate**（生成 `.rlib` 库文件）：
   - crate 根是 `src/lib.rs` 文件。
   - 例如：当你运行 `cargo new my_lib --lib`，生成的 `src/lib.rs` 是根。

 **包（Package）**

​	包是一个包含一个或多个 crate 的集合。每个包都有一个 `Cargo.toml` 配置文件，定义了包的元数据和依赖关系。

（1）**创建包**：使用 `cargo new` 命令创建一个新包。

```sh
cargo new my_package
```

（2）**包结构**：包的目录结构通常如下：

```sh
my_package/
├── Cargo.toml
└── src/
    └── main.rs
```

（3）**依赖管理**：在 `Cargo.toml` 文件的 `[dependencies]` 部分添加依赖项。

```toml
[dependencies]
serde = "1.0"
```

{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 文件名与模块名、包名之间的关系以及如何引入

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}

​	一个文件就是一个模块，文件名即模块名（不含 `.py` 后缀）。

​	在Python 3.3之前版本，若目录中包含 `__init__.py` 文件，则就是一个**包**。`__init__.py` 的作用是标识该目录是一个**包**，而不是普通的目录。**包**可以包含多个**模块**（文件）和**子包**（子目录）。参见[PEP 420: 隐式命名空间包](https://docs.python.org/zh-cn/3.9/whatsnew/3.3.html#pep-420-implicit-namespace-packages)

​	在Python 3.3及以上版本，即使目录中没有 `__init__.py` 文件，也可以将其作为**包**导入。这是因为 Python 现在支持 **隐式命名空间包**。但如果需要兼容旧版本 Python 或实现特殊逻辑，仍推荐显式地添加 `__init__.py` 文件。

**定义模块**

```python
# 文件名：math_utils.py
def add(a, b):
    return a + b
```

**引入模块**

```python
# 文件名：main.py
from math_utils import add
print(add(1, 2))
```

**对于`__init__.py`文件中的内容有什么要求？**

​	`__init__.py` 的内容是可选的，可以根据需求添加代码或保持为空。以下是常见用法：

**（1）空文件**

- 如果不需要任何初始化逻辑，`__init__.py` 文件可以留空。

  **示例**：

  ```python
  mypackage/
      __init__.py  # 空文件
      module1.py
      module2.py
  ```

**（2）导入子模块**

- 在 `__init__.py` 中导入子模块，以便可以直接通过包名访问子模块。

  **示例**：

  ```python
  # __init__.py
  from .module1 import func1
  from .module2 import func2
  ```

  **使用**：

  ```python
  import mypackage
  mypackage.func1()
  ```

**（3）定义 `__all__`**

- 指定当使用 `from package import *` 时导入的模块或符号。

  **示例**：

  ```
  # __init__.py
  __all__ = ["module1", "module2"]
  ```

  **使用**：

  ```python
  from mypackage import *
  ```

**（4）设置包级别变量**

- 定义一些全局变量或配置选项。

  **示例**：

  ```python
  # __init__.py
  VERSION = "1.0.0"
  ```

  **使用**：

  ```
  import mypackage
  print(mypackage.VERSION)
  ```

**（5）动态加载模块**

- 使用 `importlib` 或其他动态加载技术动态导入模块。

  **示例**：

  ```python
  import importlib
  
  def load_module(name):
      return importlib.import_module(name)
  ```

{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## main包、main函数

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 操作符和标点符号



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}

​	详见 [comm/Go/basic/operators]({{< ref "/comm/Go/basic/operators">}})

​	Go语言比较特殊的操作符有：

- `&^`以及`&^=`，按位清除，按位清除赋值
- `++`和`--`是语句，且只有后置`++`、`--`，而没有前置`++`、`--`
- `:=`用于短变量声明
- `<-`用在channel，目前有两种用法：1 用在函数参数、返回值类型；2 直接用在channel变量上，用于从接收channel变量的值，或者向channel发送指定值；
- `...`目前有两种用法：1 用在函数、方法的声明上的最后一个参数，表示可以接受多个参数，在函数和方法体中可以通过这个参数获得一个完整切片；2 用在函数或方法调用上，用于向函数、方法传递实参，目前可以在切片和字符串类型的变量或字面量上使用。

{{% /tab  %}}

{{% tab header="Python" %}}

​	 详见 [comm/Python/basic/operators]({{< ref "/comm/Python/basic/operators">}})

操作符

- **算术**：`+`, `-`, `*`, `/`, `//`（整除）, `%`（取余）, `**`（幂运算）
- **比较**：`==`, `!=`, `<`, `>`, `<=`, `>=`, `is`（对象身份比较）, `is not`（对象身份比较）, `in`（成员判断）, `not in`（成员判断）
- **逻辑**：`and`, `or`, `not`
- **位运算**：`&`, `|`, `^`, `~`, `<<`, `>>`
- **赋值**：`=`, `+=`, `-=`, `*=`, `/=`, `//=`, `%=`, `**=`
- **其他**：`:=` (海象运算符，3.8+版本), `@` (装饰器 / 矩阵乘法), `->` (函数返回值注解).

标点符号

- `( )` (元组 / 函数调用), 
- `[ ]` (列表 / 索引), 
- `{ }` (字典 / 集合), 
- `:` (切片 / 代码块), 
- `,` (分隔符), 
- `.` (属性访问), 
- `#` (注释), 
- `"`/`'` (字符串), 
- `\` (续行符), 
- `@` (装饰器).



```python
# 使用:=的示例
# 常规写法
line = input()
if line:
    print(f"你输入了: {line}")

# 海象运算符写法
if line := input():
    print(f"你输入了: {line}")

# 列表推导式中的应用
numbers = [n**2 for n in [1, 2, 3, 4, 5] if (result := n**2) > 10]
print(numbers)  # 输出: [16, 25]
print(result)  # 输出:25 (最后一个满足条件的值)


# 使用@的示例
import numpy as np

matrix1 = np.array([[1, 2], [3, 4]])
matrix2 = np.array([[5, 6], [7, 8]])

result = matrix1 @ matrix2
print(result)
# 输出:
# [[19 22]
#  [43 50]]
```



​	Python语言比较特殊的操作符有：

- and（逻辑与）Python中并没有`&&`，若使用 `&&`，会报 `SyntaxError`

- or（逻辑或）Python中并没有`||`，若使用 `||`，会报 `SyntaxError`

- not（逻辑非）Python中并没有`!`，若使用 `!`，会报 `SyntaxError`

  

{{% /tab  %}}

{{% tab header="Java" %}}

操作符

- **算术**：`+`, `-`, `*`, `/`, `%`, `++`（区分前置和后置形式）, `--`（区分前置和后置形式）
- **比较**：`==`, `!=`, `<`, `>`, `<=`, `>=`, `instanceof`（检查一个对象是否是某个类或接口的实例）
- **逻辑**：`&&`, `||`, `!`
- **位运算**：`&`, `|`, `^`（按位异或）, `~`（按位取反）, `<<`（带符号位左移，只用于整数类型）, `>>`（带符号位右移，只用于整数类型）, `>>>`（不带符号位右移，只用于整数类型）
- **赋值**：`=`, `+=`, `-=`, `*=`, `/=`, `%=`, `&=`, `|=`, `^=`, `<<=`, `>>=`, `>>>=`
- **三元**：`? :`

标点符号

- `{ }` (代码块), 
- `[ ]` (数组), 
- `( )` (方法参数), 
- `;` (语句结束), 
- `,` (分隔符), 
- `.` (对象成员访问), 
- `@` (注解), 
- `"` (字符串), 
- `'` (字符).

{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 数据类型

### 数据类型

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}

- 内置类型：
  - 布尔型：`bool`，表示真或假。
  - 数值型：
    - 整数：`int`、`int8`、`int16`、`int32`、`int64`、`uint`、`uint8`（`byte` 别名）、`uint16`、`uint32`、`uint64`、`uintptr`。
    - 浮点数：`float32`、`float64`。
    - 复数：`complex64`、`complex128`。
  - 字符串型：`string`，用于存储文本数据。
- 复合类型：
  - 数组：`[n]T`，固定长度，`n` 为长度，`T` 为元素类型。
  - 切片：`[]T`，长度可变，基于数组实现。
  - 映射：`map[K]V`，键值对集合，`K` 为键类型，`V` 为值类型。
  - 结构体：`struct`，用于组合不同类型的数据。
  - 接口：`interface`，定义一组方法签名。
- 自定义类型：
  - 使用 `type` 关键字，如 `type MyInt int` 定义 `MyInt` 为 `int` 类型的别名；`type Person struct { Name string; Age int }` 定义一个结构体类型。

```go
package main

import "fmt"

func main() {
    // 内置类型
    var b bool = true       // 布尔型
    var i int = 10         // 整型
    var i8 int8 = 8        // 8位整型
    var i16 int16 = 16      // 16位整型
    var i32 int32 = 32      // 32位整型
    var i64 int64 = 64      // 64位整型
    var ui uint = 10       // 无符号整型
    var ui8 uint8 = 8      // 8位无符号整型
    var ui16 uint16 = 16    // 16位无符号整型
    var ui32 uint32 = 32    // 32位无符号整型
    var ui64 uint64 = 64    // 64位无符号整型
    var f float32 = 3.14    // 浮点型
    var f64 float64 = 3.1415926  // 双精度浮点型
    var c complex128 = 1 + 2i // 复数型
    var s string = "hello"   // 字符串型
    var by byte = 'a'      // 字节型
    var r rune = '中'      // 符文型

    fmt.Println(b, i, i8, i16, i32, i64, ui, ui8, ui16, ui32, ui64, f, f64, c, s, by, r)
    // true 10 8 16 32 64 10 8 16 32 64 3.14 3.1415926 (1+2i) hello 97 20013

    // 复合类型
    var arr [3]int = [3]int{1, 2, 3} // 数组
    var slice []int = []int{1, 2, 3} // 切片
    var m map[string]int = map[string]int{"a": 1, "b": 2} // 映射
    type Person struct { // 结构体
        Name string
        Age  int
    }
    var p Person = Person{"Alice", 30}
    var ptr *int = &i // 指针
    var f1 func(int) int = func(x int) int { return x * 2 } // 函数类型
    type MyInterface interface { // 接口
        MyMethod()
    }

    fmt.Println(arr, slice, m, p, ptr, f1)
    // [1 2 3] [1 2 3] map[a:1 b:2] {Alice 30} 0xc0000140a0 0xc0000100c0

    // 自定义类型
    type MyInt int // 类型别名
    type MyStruct struct { // 结构体
        Name string
    }
    var myInt MyInt = 20
    var myStruct MyStruct = MyStruct{"Bob"}

    fmt.Println(myInt, myStruct)
    // 20 {Bob}
}
```

{{% /tab  %}}

{{% tab header="Python" %}}

- 内置类型：
  - 布尔型：`bool`，`True` 或 `False`。
  - 数值型：
    - 整数：`int`，任意大小的整数。
    - 浮点数：`float`，带小数点的数。
    - 复数：`complex`，如 `3 + 4j`。
  - 字符串型：`str`，用于存储文本，用单引号、双引号或三引号表示。
  - 序列类型：
    - 列表：`list`，有序可变元素集合，如 `[1, 2, "hello"]`。
    - 元组：`tuple`，有序不可变元素集合，如 `(1, 2, "world")`。
  - 集合类型：
    - 集合：`set`，无序不重复元素集合，如 `{1, 2, 3}`。
    - 冻结集合：`frozenset`，不可变的集合。
  - 映射类型：`dict`，键值对集合，如 `{"name": "Alice", "age": 25}`。
- 自定义类型：
  - 使用 `class` 关键字定义类，如 `class Person: def __init__(self, name, age): self.name = name; self.age = age`。

```python
# 内置类型
b = True  # 布尔型
i = 10  # 整型
f = 3.14  # 浮点型
c = 1 + 2j  # 复数型
s = "hello"  # 字符串型
by = b"hello"  # 字节型
l = [1, 2, 3]  # 列表
t = (1, 2, 3)  # 元组
s1 = {1, 2, 3}  # 集合
d = {"a": 1, "b": 2}  # 字典
by_arr = bytearray(b"hello")  # 字节数组

print(b, i, f, c, s, by, l, t, s1, d, by_arr)
# True 10 3.14 (1+2j) hello b'hello' [1, 2, 3] (1, 2, 3) {1, 2, 3} {'a': 1, 'b': 2} bytearray(b'hello')

# 自定义类型
class MyClass:  # 类
    pass

my_object = MyClass()

print(my_object)
# <__main__.MyClass object at 0x...>
```

{{% /tab  %}}

{{% tab header="Java" %}}

- 内置类型（基本数据类型）：
  - 布尔型：`boolean`，值为 `true` 或 `false`。
  - 数值型：
    - 整数：`byte`（8 位）、`short`（16 位）、`int`（32 位）、`long`（64 位）。
    - 浮点数：`float`（32 位）、`double`（64 位）。
    - 字符型：`char`，表示单个字符，用单引号括起来。
- 复合类型（引用类型）：
  - 类：`class`，用于封装数据和行为。
  - 接口：`interface`，定义一组方法签名。
  - 数组：`type[]`，如 `int[] numbers = new int[5];`。
  - 枚举：`enum`，自 Java 5 引入，用于定义一组常量。
- 自定义类型：
  - 使用 `class` 关键字定义类，如 `class Person { String name; int age; }`；使用 `interface` 关键字定义接口；使用 `enum` 关键字定义枚举类型。

```java
public class Main {
    public static void main(String[] args) {
        // 内置类型（原始类型）
        boolean b = true; // 布尔型
        byte by = 10; // 字节型
        short s = 20; // 短整型
        int i = 30; // 整型
        long l = 40L; // 长整型
        float f = 3.14f; // 单精度浮点型
        double d = 3.14159; // 双精度浮点型
        char c = 'a'; // 字符型

        System.out.println(b + " " + by + " " + s + " " + i + " " + l + " " + f + " " + d + " " + c);
        // true 10 20 30 40 3.14 3.14159 a

        // 复合类型（引用类型）
        int[] arr = {1, 2, 3}; // 数组
        class MyClass { // 类
            String name;
        }
        MyClass myObj = new MyClass();
        myObj.name = "Alice";
        interface MyInterface { // 接口
            void myMethod();
        }
        enum MyEnum { // 枚举
            A, B, C
        }

        System.out.println(arr[0] + " " + myObj.name + " " + MyEnum.A);
        // 1 Alice A

        // 自定义类型
        class MyOtherClass { // 类
            int value;
        }

        MyOtherClass myOtherObj = new MyOtherClass();
        myOtherObj.value = 100;

        System.out.println(myOtherObj.value);
        // 100
    }
}
```

{{% /tab  %}}

{{% tab header="Rust" %}}

- 内置类型：
  - 布尔型：`bool`，`true` 或 `false`。
  - 数值型：
    - 整数：`i8`、`i16`、`i32`、`i64`、`i128`、`isize`、`u8`、`u16`、`u32`、`u64`、`u128`、`usize`。
    - 浮点数：`f32`、`f64`。
    - 字符型：`char`，表示单个 Unicode 字符。
  - 字符串类型：
    - `str`（字符串切片），不可变借用字符串。
    - `String`，可变字符串。
- 复合类型：
  - 元组：`(type1, type2,...)`，固定长度，元素类型可以不同。
  - 数组：`[T; n]`，固定长度，`T` 为元素类型，`n` 为长度。
  - 结构体：`struct`，用于组合不同类型的数据。
  - 枚举：`enum`，用于定义多种可能的类型。
  - 联合体：`union`（不稳定特性）。
- 自定义类型：
  - 使用 `type` 关键字定义类型别名，如 `type MyInt = i32;`；使用 `struct` 关键字定义结构体，如 `struct Person { name: String, age: u32 }`；使用 `enum` 关键字定义枚举，如 `enum Color { Red, Green, Blue }`。

```rust
fn main() {
    // 内置类型
    let b: bool = true; // 布尔型
    let i: i32 = 10; // 整型
    let i8: i8 = 8; // 8位整型
    let i16: i16 = 16; // 16位整型
    let i32: i32 = 32; // 32位整型
    let i64: i64 = 64; // 64位整型
    let isize: isize = 123; // 根据平台大小而定的整型
    let u8: u8 = 8; // 8位无符号整型
    let u16: u16 = 16; // 16位无符号整型
    let u32: u32 = 32; // 32位无符号整型
    let u64: u64 = 64; // 64位无符号整型
    let usize: usize = 123; // 根据平台大小而定的无符号整型
    let f: f32 = 3.14; // 浮点型
    let f64: f64 = 3.1415926; // 双精度浮点型
    let c: char = 'a'; // 字符型
    let s: &str = "hello"; // 字符串切片
    let string: String = String::from("world"); // 字符串

    println!("{} {} {} {} {} {} {} {} {} {} {} {} {} {} {} {}", b, i, i8, i16, i32, i64, isize, u8, u16, u32, u64, usize, f, f64, c, s, string);
    // true 10 8 16 32 64 123 8 16 32 64 123 3.14 3.1415926 a hello world

    // 复合类型
    let arr: [i32; 3] = [1, 2, 3]; // 数组
    let slice: &[i32] = &arr[..]; // 切片
    let tuple: (i32, &str) = (1, "hello"); // 元组
    struct MyStruct { // 结构体
        name: String,
    }
    let my_struct: MyStruct = MyStruct { name: String::from("Alice") };
    enum MyEnum { // 枚举
        A,
        B,
    }
    let my_enum: MyEnum = MyEnum::A;
    let ptr: *const i32 = &i; // const指针
    let mut mutable_i: i32 = 20;
    let mutable_ptr: *mut i32 = &mut mutable_i;// 可变指针
    type MyType = i32; // 类型别名

    println!("{:?} {:?} {:?} {:?} {:?} {:?} {:?} {:?} {:?}", arr, slice, tuple, my_struct, my_enum, ptr, mutable_ptr, MyType::default());
    // [1, 2, 3] [1, 2, 3] (1, "hello") MyStruct { name: "Alice" } A 0x7ff7b5a041f0 0x7ff7b5a04204 0

    // 自定义类型
    struct MyOtherStruct { // 结构体
        value: i32,
    }

    let my_other_struct: MyOtherStruct = MyOtherStruct { value: 100 };

    println!("{}", my_other_struct.value);
    // 100
}
```

{{% /tab  %}}

{{% tab header="C/C++" %}}

- C 语言：

  - 内置类型：
    - 布尔型：`_Bool`（C99 引入，通常用 `stdbool.h` 头文件中的 `bool` 宏，值为 `true` 或 `false`）。
    - 数值型：
      - 整数：`char`、`short`、`int`、`long`、`long long`（C99 引入），每种又分 `signed` 和 `unsigned`。
      - 浮点数：`float`、`double`、`long double`。
    - 空类型：`void`，表示无类型。
  - 复合类型：
    - 数组：`type array_name[size];`，固定长度。
    - 结构体：`struct`，用于组合不同类型的数据。
    - 联合体：`union`，用于共享内存空间。
    - 指针：`type *pointer_name;`，存储变量地址。
  - 自定义类型：
    - 使用 `typedef` 关键字定义类型别名，如 `typedef int MyInt;`；使用 `struct` 关键字定义结构体，如 `struct Person { char name[20]; int age; };`。

- C++ 语言：

  - **内置类型**：在 C 语言基础上，布尔型为 `bool`，值为 `true` 或 `false`。
  - 复合类型：
    - 数组：`type array_name[size];`，固定长度，C++11 引入 `std::array` 更安全易用。
    - 结构体：`struct`，与 C 类似，但更面向对象。
    - 类：`class`，用于封装数据和行为，是 C++ 面向对象的核心。
    - 联合体：`union`，与 C 类似。
    - 指针：`type *pointer_name;`，存储变量地址，C++11 引入智能指针 `std::unique_ptr`、`std::shared_ptr`、`std::weak_ptr` 等。
    - 引用：`type &reference_name = variable;`，为变量的别名。
    - 枚举：`enum`，可以定义强类型枚举 `enum class`（C++11 引入）。
  - 自定义类型：
    - 使用 `typedef` 关键字定义类型别名；使用 `class` 关键字定义类，如 `class Person { public: string name; int age; };`；使用 `struct` 关键字定义结构体；使用 `enum` 或 `enum class` 定义枚举类型。

  ```c++
  #include <iostream>
  #include <string>
  #include <vector>
  
  using namespace std;
  
  int main() {
      // 内置类型（基本类型）
      bool b = true; // 布尔型
      int i = 10; // 整型
      float f = 3.14f; // 单精度浮点型
      double d = 3.14159; // 双精度浮点型
      char c = 'a'; // 字符型
      string s = "hello"; // 字符串
  
      cout << b << " " << i << " " << f << " " << d << " " << c << " " << s << endl; // 1 10 3.14 3.14159 a hello
  
      // 复合类型（派生类型）
      int arr[] = {1, 2, 3}; // 数组
      int* ptr = &i; // 指针
      struct MyStruct { // 结构体
          string name;
          int age;
      };
      MyStruct myStruct = {"Alice", 30};
      enum MyEnum { // 枚举
          A, B, C
      };
      MyEnum myEnum = MyEnum::A;
      vector<int> vec = {1, 2, 3}; // 动态数组（vector）
  
      cout << arr[0] << " " << *ptr << " " << myStruct.name << " " << myEnum << " "; // 1 10 Alice 0
      for (int val : vec) {
          cout << val << " "; // 1 2 3
      }
      cout << endl;
  
      // 自定义类型
      class MyClass { // 类
      public:
          int value;
      };
  
      MyClass myObj;
      myObj.value = 100;
  
      cout << myObj.value << endl; // 100
  
      return 0;
  }
  ```

  

{{% /tab  %}}

{{% tab header="JavaScript" %}}

- 内置类型：
  - 原始类型：
    - 布尔型：`boolean`，`true` 或 `false`。
    - 数值型：`number`，表示整数和浮点数。
    - 字符串型：`string`，用于存储文本。
    - 空值：`null`，表示空对象指针。
    - 未定义：`undefined`，表示变量未赋值。
    - 符号：`symbol`（ES6 引入），用于创建唯一标识符。
  - 对象类型：`object`，用于存储键值对集合，如 `{ name: "Alice", age: 25 }`。
- 自定义类型：
  - 使用构造函数、类（ES6 引入 `class` 语法糖，本质还是构造函数）等方式创建自定义对象，如 `class Person { constructor(name, age) { this.name = name; this.age = age; } }`。

```javascript
// 原始类型
let b = true; // 布尔型
let n = 10; // 数值型
let s = "hello"; // 字符串型
let nu = null; // 空值
let un = undefined; // 未定义
let sy = Symbol("a"); // 符号
let bi = 1234567890123456789012345678901234567890n; // BigInt

console.log(b, n, s, nu, un, sy, bi); // true 10 hello null undefined Symbol(a) 1234567890123456789012345678901234567890n

// 复合类型（对象类型）
let obj = { name: "Alice", age: 30 }; // 对象
let arr = [1, 2, 3]; // 数组
function myFunction(x, y) { return x + y; } // 函数

console.log(obj, arr, myFunction(2, 3)); // { name: 'Alice', age: 30 } [ 1, 2, 3 ] 5

// 自定义类型
class MyClass { // 类
    constructor(value) {
        this.value = value;
    }
}

let myObj = new MyClass(100);

console.log(myObj.value); // 100
```

{{% /tab  %}}

{{% tab header="TypeScript" %}}

- 内置类型：
  - 与 JavaScript 原始类型类似：`boolean`、`number`、`string`、`null`、`undefined`、`symbol`。
  - 新增：
    - 任意类型：`any`，可以是任何类型。
    - 未知类型：`unknown`，与 `any` 类似但更安全。
    - 空类型：`void`，表示函数没有返回值。
    - never 类型：`never`，表示从不会出现的值。
  - 数组类型：`type[]` 或 `Array<type>`，如 `number[]` 或 `Array<number>`。
  - 元组类型：`[type1, type2,...]`，固定长度，元素类型固定。
- 自定义类型：
  - 使用 `type` 关键字定义类型别名，如 `type MyNumber = number;`；使用 `interface` 关键字定义接口，如 `interface Person { name: string; age: number; }`；使用 `class` 关键字定义类，如 `class Person { constructor(public name: string, public age: number) {} }`。

```typescript
// 原始类型
let b: boolean = true; // 布尔型
let n: number = 10; // 数值型
let s: string = "hello"; // 字符串型
let nu: null = null; // 空值
let un: undefined = undefined; // 未定义
let sy: symbol = Symbol("a"); // 符号
let bi: bigint = 1234567890123456789012345678901234567890n; // BigInt

console.log(b, n, s, nu, un, sy, bi); // true 10 hello null undefined Symbol(a) 1234567890123456789012345678901234567890n

// 复合类型（对象类型）
let obj: { name: string; age: number } = { name: "Alice", age: 30 }; // 对象
let arr: number[] = [1, 2, 3]; // 数组
let myFunction: (x: number, y: number) => number = (x, y) => x + y; // 函数
let tuple: [string, number] = ["hello", 10]; // 元组
enum MyEnum { A, B, C } // 枚举
let myEnum: MyEnum = MyEnum.A;

console.log(obj, arr, myFunction(2, 3), tuple, myEnum); // { name: 'Alice', age: 30 } [ 1, 2, 3 ] 5 [ 'hello', 10 ] 0

// 自定义类型
class MyClass { // 类
    value: number;
    constructor(value: number) {
        this.value = value;
    }
}

interface MyInterface { // 接口
    value: number;
}

type MyType = { // 类型别名
    value: number;
};

let myObj = new MyClass(100);
let myInterfaceObj: MyInterface = { value: 200 };
let myTypeObj: MyType = { value: 300 };

console.log(myObj.value, myInterfaceObj.value, myTypeObj.value); // 100 200 300
```

{{% /tab  %}}

{{% tab header="C#" %}}

- 内置类型（值类型）：
  - 布尔型：`bool`，`true` 或 `false`。
  - 数值型：
    - 整数：`sbyte`、`byte`、`short`、`ushort`、`int`、`uint`、`long`、`ulong`。
    - 浮点数：`float`、`double`、`decimal`。
    - 字符型：`char`，表示单个 Unicode 字符。
  - 枚举类型：`enum`，用于定义一组命名常量。
  - 结构体类型：`struct`，轻量级对象，值类型。
- 内置类型（引用类型）：
  - 对象类型：`object`，所有类型的基类。
  - 字符串类型：`string`，用于存储文本。
  - 数组类型：`type[]`，如 `int[] numbers = new int[5];`。
  - 类：`class`，用于封装数据和行为。
  - 接口：`interface`，定义一组方法签名。
  - 委托：`delegate`，用于实现回调函数。
  - 事件：基于委托，用于对象间的通信。
- 自定义类型：
  - 使用 `class` 关键字定义类，如 `class Person { public string name; public int age; }`；使用 `struct` 关键字定义结构体；使用 `interface` 关键字定义接口；使用 `enum` 关键字定义枚举类型；使用 `delegate` 关键字定义委托类型。

```c#
using System;
using System.Collections.Generic;

public class Program
{
    public static void Main(string[] args)
    {
        // 值类型
        bool b = true; // 布尔型
        byte by = 10; // 字节型
        short s = 20; // 短整型
        int i = 30; // 整型
        long l = 40L; // 长整型
        float f = 3.14f; // 单精度浮点型
        double d = 3.14159; // 双精度浮点型
        char c = 'a'; // 字符型
        decimal de = 100.00M; // 十进制型
        enum MyEnum { A, B, C } // 枚举
        MyEnum myEnum = MyEnum.A;
        struct MyStruct { public string name; public int age; } // 结构体
        MyStruct myStruct = new MyStruct { name = "Alice", age = 30 };

        Console.WriteLine($"{b} {by} {s} {i} {l} {f} {d} {c} {de} {(int)myEnum} {myStruct.name} {myStruct.age}"); // True 10 20 30 40 3.14 3.14159 a 100.00 0 Alice 30

        // 引用类型
        string str = "hello"; // 字符串
        int[] arr = { 1, 2, 3 }; // 数组
        class MyClass { public int value; } // 类
        MyClass myObj = new MyClass { value = 100 };
        interface MyInterface { int MyMethod(); } // 接口
        delegate int MyDelegate(int x); // 委托

        Console.WriteLine($"{str} {arr[0]} {myObj.value}"); // hello 1 100

        // 自定义类型
        class MyOtherClass : MyClass { } // 类
        interface MyOtherInterface : MyInterface { } // 接口

        Console.WriteLine(new MyOtherClass().value); // 0
    }
}
```

{{% /tab  %}}

{{% tab header="Erlang" %}}

- 内置类型：
  - 布尔型：`true`、`false`。
  - 数值型：
    - 整数：任意精度整数。
    - 浮点数：`float`。
  - 原子：`atom`，常量，以小写字母开头或用单引号括起来，如 `ok`、`'MyAtom'`。
  - 字符串：实际上是整数列表，如 `[97, 98, 99]` 表示 `"abc"`。
  - 列表：`[element1, element2,...]`，元素类型可以不同。
  - 元组：`{element1, element2,...}`，固定长度，元素类型可以不同。
- 自定义类型：
  - 使用 `record` 语法定义记录类型，如 `-record(person, {name, age}).`，使用时 `#person{name = "Alice", age = 25}`。

```erlang
-module(main).
-export([main/0]).

main() ->
    % 基本类型
    B = true, % 布尔型
    I = 10, % 整数
    F = 3.14, % 浮点数
    A = atom, % 原子
    S = "hello", % 字符串
    Bs = <<1, 2, 3>>, % 位串
    N = nil, % 空列表

    io:format("~p ~p ~p ~p ~p ~p ~p~n", [B, I, F, A, S, Bs, N]), % true 10 3.14 atom "hello" <<1,2,3>> []

    % 复合类型
    Tuple = {1, "hello", 3.14}, % 元组
    List = [1, 2, 3], % 列表
    Map = #{a => 1, b => 2}, % 映射
    Binary = <<1, 2, 3>>, % 二进制型

    io:format("~p ~p ~p ~p~n", [Tuple, List, Map, Binary]), % {1,"hello",3.14} [1,2,3] #{a => 1,b => 2} <<1,2,3>>

    % 自定义类型
    -type my_int() :: integer(). % 类型别名
    -record(person, {name = "", age = 0}). % 记录

    MyInt :: my_int() = 20.
    Person = #person{name = "Alice", age = 30}.

    io:format("~p ~p~n", [MyInt, Person]). % 20 {person, "Alice", 30}
```

{{% /tab  %}}

{{% tab header="PHP" %}}

- 内置类型：
  - 布尔型：`bool`，`true` 或 `false`。
  - 数值型：
    - 整数：`int`，表示整数。
    - 浮点数：`float`（`double` 别名），表示浮点数。
  - 字符串型：`string`，用于存储文本。
  - 数组：`array`，可以是索引数组、关联数组等多种形式，如 `$arr = [1, 2, 3];` 或 `$assocArr = ["name" => "Alice", "age" => 25];`。
  - 对象：`object`，使用 `class` 定义的类的实例。
  - 资源：`resource`，用于表示外部资源，如数据库连接等。
  - 空值：`null`，表示变量没有值。
- 自定义类型：
  - 使用 `class` 关键字定义类，如 `class Person { public $name; public $age; }`。

```php
<?php
// 标量类型
$b = true; // 布尔型
$i = 10; // 整型
$f = 3.14; // 浮点型
$s = "hello"; // 字符串型

echo $b . " " . $i . " " . $f . " " . $s . "\n"; // 1 10 3.14 hello

// 复合类型
$arr = array(1, 2, 3); // 数组
$obj = new stdClass(); // 对象
$obj->name = "Alice";

echo $arr[0] . " " . $obj->name . "\n"; // 1 Alice

// 特殊类型
$null = NULL; // NULL 类型

echo var_export($null) . "\n"; // NULL

// 自定义类型
class MyClass { // 类
    public $value;
}

$myObj = new MyClass();
$myObj->value = 100;

echo $myObj->value . "\n"; // 100
?>
```

{{% /tab  %}}

{{% tab header="Ruby" %}}

- 内置类型：
  - 布尔型：`true`、`false`。
  - 数值型：
    - 整数：`Fixnum`（小整数，范围有限）、`Bignum`（大整数，任意精度），在 Ruby 2.4 后统一为 `Integer`。
    - 浮点数：`Float`。
  - 字符串型：`String`，用于存储文本。
  - 数组：`Array`，有序元素集合，元素类型可以不同，如 `[1, "hello", true]`。
  - 哈希：`Hash`，键值对集合，如 `{name: "Alice", age: 25}`。
  - 范围：`Range`，表示一个范围，如 `1..10`。
  - 符号：`Symbol`，不可变的字符串 - like 对象，用于标识符等，如 `:name`。
  - 正则表达式：`Regexp`，用于模式匹配。
  - proc 和 lambda：`Proc`、`lambda`，用于创建可调用的代码块。
- 自定义类型：
  - 使用 `class` 关键字定义类，如 `class Person attr_accessor :name, :age def initialize(name, age) @name = name @age = age end end`。

```ruby
# 内置类型
b = true # 布尔型
i = 10 # 整型
f = 3.14 # 浮点型
s = "hello" # 字符串型
sym = :symbol # 符号

puts "#{b} #{i} #{f} #{s} #{sym}" # true 10 3.14 hello symbol

# 复合类型
arr = [1, 2, 3] # 数组
hash = { "a" => 1, "b" => 2 } # 哈希
range = 1..10 # 范围

puts "#{arr} #{hash} #{range}" # [1, 2, 3] {"a"=>1, "b"=>2} 1..10

# 自定义类型
class MyClass # 类
  def initialize(value)
    @value = value
  end
end

my_obj = MyClass.new(100)

puts my_obj.instance_variable_get(:@value) # 100
```

{{% /tab  %}}

{{< /tabpane >}}

### 类型转换

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 声明和作用域



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 常量



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 变量



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 命名规范

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 控制语句

### 判断语句（选择语句）

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 循环语句

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 内置函数

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 字符串

### 字符串字面值（string literal）

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 格式化字符串

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

### 字符串常见操作

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}



## 指针



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 模块



{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 测试

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}



## 异常和错误

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 继承

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 编码

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## 正则表达式

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}

## I/O操作

{{< tabpane text=true persist=disabled >}}

{{% tab header="Go" %}}



{{% /tab  %}}

{{% tab header="Python" %}}



{{% /tab  %}}

{{% tab header="Java" %}}



{{% /tab  %}}

{{% tab header="Rust" %}}



{{% /tab  %}}

{{% tab header="C/C++" %}}



{{% /tab  %}}

{{% tab header="JavaScript" %}}



{{% /tab  %}}

{{% tab header="TypeScript" %}}



{{% /tab  %}}

{{% tab header="C#" %}}



{{% /tab  %}}

{{% tab header="Erlang" %}}



{{% /tab  %}}

{{% tab header="PHP" %}}



{{% /tab  %}}

{{% tab header="Ruby" %}}



{{% /tab  %}}

{{< /tabpane >}}