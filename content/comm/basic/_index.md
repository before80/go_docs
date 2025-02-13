+++
title = "基础"
date = 2024-03-01T15:18:46+08:00
weight = -100
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

Go、Python、Java、Rust、C/C++、JavaScript、TypeScript、C#、Erlang、PHP、Ruby

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



{{% /tab  %}}

{{% tab header="Python" %}}

`pip`是Python自带的，无需安装，若有新版本只需更新。

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

查看包的可用版本

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

安装包

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

升级包

```sh
# 升级 requests 到最新版
pip install --upgrade requests

# 简写形式
pip install -U requests
```

卸载包

```sh
#  卸载 pandsa
pip uninstall pandas
```

依赖管理

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

列出已安装的包

```sh
# 列出所有包
pip list

# 列出过期的包
pip list --outdated
```

查看包详情

```sh
# 列出所有包
pip list

# 列出过期的包
pip list --outdated
```

设置全局镜像源

```sh
# 永久配置清华镜像（Linux/macOS）
pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple

```

查看当前配置

```sh
pip config list
```

安装开发模式（可编辑模式）

```sh
# 本地开发时直接链接代码（无需重复安装）
pip install -e ./my_package
```

忽略缓存安装

```sh
# 避免使用旧缓存文件
pip install --no-cache-dir torch
```

限制超时时间

```sh
# 设置超时为 60 秒
pip install --timeout 60 scipy
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