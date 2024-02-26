+++
title = "下载并安装 Go"
weight = 1
date = 2023-05-18T16:35:08+08:00
description = ""
isCJKLanguage = true
draft = false

+++
# Download and install 下载并安装

> 原文：[https://go.dev/doc/install](https://go.dev/doc/install)

按照这里描述的步骤，快速下载并安装Go。

对于其他有关安装的内容，您可能会感兴趣

- [管理 Go 的安装](../ManagingGoInstallations) —— 如何安装多个版本和卸载。
- [从源码安装Go](../InstallingGoFromSource) —— 如何检出源码，在自己的机器上构建它们，并运行它们。

## 1. Go download.

点击下面的按钮，下载Go的安装程序。

{{< tabpane text=true >}}

{{< tab header="Linux" >}}

[Download Go for Linux go1.19.3.linux-amd64.tar.gz (142 MB)](https://go.dev/dl/go1.19.3.windows-amd64.msi)

{{< /tab  >}}

{{< tab header="Mac" >}}

[Download Go for Mac go1.19.3.darwin-amd64.pkg (145 MB)](https://go.dev/dl/go1.19.3.darwin-amd64.pkg)

{{< /tab  >}}

{{< tab header="Windows" >}}

[Download Go for Windows go1.19.3.windows-amd64.msi (135 MB)](https://go.dev/dl/go1.19.3.windows-amd64.msi)

{{< /tab  >}}

{{< /tabpane >}}



在这里没有看到您的操作系统？试试[其他的下载](https://go.dev/dl/)。

注意：默认情况下，`go` 命令使用 Google 运行的 `Go 模块镜像`和 `Go 校验数据库`下载和验证模块。[了解更多](https://go.dev/dl)。

## 2. Go install.

选择以下计算机操作系统的选项卡，然后按照其安装说明进行安装。

{{< tabpane text=true >}}

{{< tab header="Linux" >}}

1. 通过删除`/usr/local/go`文件夹（如果存在）来删除之前的任何Go安装，然后将刚刚下载的存档解压到`/usr/local`，在`/usr/local/go`中创建一个新的Go树：
```sh
$ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.3.linux-amd64.tar.gz
```
(您可能需要以root身份或通过sudo运行该命令）。

   不要将存档解压到现有的`/usr/local/go` 树中。众所周知，这样做会导致Go安装失败。

2. 在`PATH`环境变量中加入`/usr/local/go/bin`。
   您可以通过在 `$HOME/.profile` 或 `/etc/profile`（对于全系统的安装）中添加以下一行来实现。

```bash
export PATH=$PATH:/usr/local/go/bin
```

>  **注意**：对配置文件所做的修改可能要到您下次登录电脑时才会应用。要立即应用这些更改，只需直接运行 shell 命令或使用 `source $HOME/.profile` 等命令从配置文件中执行这些命令。
>

3. 通过打开命令提示符并输入以下命令来验证您已经安装了Go:
```bash
$ go version
```
4. 确认该命令打印出已安装的Go的版本。

{{< /tab >}}

{{< tab header="Mac" >}}

1. 打开您下载的包文件，按照提示安装Go。
   
   该包将Go发行版安装到`/usr/local/go`。该包应该把 `/usr/local/go/bin` 目录放到您的 `PATH` 环境变量中。您可能需要重新启动任何打开的终端会话，以使该变化生效。

2. 通过打开命令提示符并输入以下命令来验证您已经安装了Go：

    ```bash
    $ go version
    ```

3. 确认该命令打印出已安装的Go的版本。

{{< /tab >}}

{{< tab header="Windows" >}}

1. 打开您下载的`MSI`文件，按照提示安装Go。默认情况下，安装程序将把Go安装到`Program Files`或`Program Files (x86)`。您可以根据需要改变位置。安装后，您需要关闭并重新打开任何打开的命令提示符，以便安装程序对环境的改变反映在命令提示符上。
2. 验证您是否已经安装了Go。
> 1. 在Windows中，点击 "开始"菜单。 
>
> 2. 在菜单的搜索框中，输入cmd，然后按回车键。    
>
> 3. 在出现的命令提示符窗口中，输入以下命令:
>    ``` shell
>    $ go version
>    ```
> 4. 确认该命令打印了已安装的Go的版本。

{{< /tab >}}

{{< /tabpane >}}




## 3. Go code.

您已经准备好了! 访问[入门教程](../TutorialGetStartedWithGo)，编写一些简单的Go代码。大约需要10分钟完成。