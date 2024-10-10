+++
title = "Systemd 服务"
date = 2024-02-04T21:20:00+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/deploy/systemd/]({{< ref "/buffalo/deploy/systemdService" >}})

# Systemd Service Systemd 服务 

In this chapter, we’ll see how to setup your Buffalo app as a Systemd service. Systemd is the new standard on many GNU/Linux distributions, for running the system services.

​	在本章中，我们将了解如何将 Buffalo 应用设置为 Systemd 服务。Systemd 是许多 GNU/Linux 发行版上运行系统服务的最新标准。

It allows you to configure an application in a standard way, and manage its lifecycle with `systemctl` commands. You can refer to the [systemd man page](https://www.freedesktop.org/software/systemd/man/systemd.service.html) for further information.

​	它允许您以标准方式配置应用程序，并使用 `systemctl` 命令管理其生命周期。您可以参阅 systemd 手册页以获取更多信息。

## Install Your Buffalo App 安装 Buffalo 应用 

The first step is to place your app into the right folder: on Debian, the common place for executables installed by hand is `/usr/local/bin`. That’s where we’ll install the app.

​	第一步是将您的应用放在正确的文件夹中：在 Debian 上，手动安装的可执行文件的常见位置是 `/usr/local/bin` 。我们将在那里安装该应用。

```bash
$ sudo mv myapp /usr/local/bin
```

Ensure the rights are correctly set, and give the ownership to the user you want to use. Here, I’ll use the `root` account.

​	确保正确设置权限，并将所有权授予您想要使用的用户。在这里，我将使用 `root` 帐户。

```bash
$ sudo chown root: /usr/local/bin/myapp
$ sudo chmod +x /usr/local/bin/myapp
```

## Create the Systemd Config File 创建 Systemd 配置文件 

The systemd service files are located in `/lib/systemd/system/`, we’ll create a new `myapp.service` file there for your app.

​	systemd 服务文件位于 `/lib/systemd/system/` 中，我们将在那里为您的应用创建一个新的 `myapp.service` 文件。

```ini
[Unit]
Description=My super app

[Service]
ExecStart=/usr/local/bin/myapp
User=root
Group=root
UMask=007

[Install]
WantedBy=multi-user.target
```

Here, we create a new service with an readable name “My super app”. It’s a simple service, which will spawn a new process described with `ExecStart`: the absolute path to your app. This process will be executed as `root:root`, with a `UMask` giving rights only to the process owner (root).

​	在这里，我们创建一个具有可读名称“My super app”的新服务。这是一个简单的服务，它将生成一个用 `ExecStart` 描述的新进程：您的应用的绝对路径。此进程将作为 `root:root` 执行，其中 `UMask` 仅授予进程所有者 (root) 权限。

In the `Install` section, we just tell Systemd to wait for a ready system. If you have more requirements, you can ask Systemd to wait for a database, for instance:

​	在 `Install` 部分，我们只是告诉 Systemd 等待系统就绪。如果您有更多要求，您可以要求 Systemd 等待数据库，例如：

```ini
[Unit]
Description=My super app
After=mysql.service

[Service]
ExecStart=/usr/local/bin/myapp
User=root
Group=root
UMask=007

[Install]
WantedBy=multi-user.target
```

## Set Env Variables 设置环境变量 

The official way to handle config with Buffalo is through [environment variables]({{< ref "/buffalo/gettingStarted/configuration" >}}). Using Systemd, you can set them with an override file.

​	使用 Buffalo 处理配置的官方方法是通过环境变量。使用 Systemd，您可以使用覆盖文件设置它们。

Our override file will be located in `/etc/systemd/system/myapp.service.d/`, and be called `override.conf`.

​	我们的覆盖文件将位于 `/etc/systemd/system/myapp.service.d/` 中，并称为 `override.conf` 。

```ini
[Service]
Environment="ADDR=0.0.0.0"
Environment="GO_ENV=production"
Environment="SESSION_SECRET=kqdjmlkajdùméa]$"
```

Each `Environment` line define an environment variable for your app.

​	每行 `Environment` 为您的应用定义一个环境变量。

## Play with the Service 使用服务 

The systemd service is now ready, you can test it with the `systemctl` and `journalctl` commands:

​	systemd 服务现已就绪，您可以使用 `systemctl` 和 `journalctl` 命令对其进行测试：

```bash
$ sudo systemctl start myapp.service
```

To start the service, and check if everything is running fine.

​	启动服务，并检查一切是否运行良好。

```bash
$ journalctl -u myapp.service -f
```

To read the logs from the standard output (`-u` to set the service name, `-f` to follow the logs).

​	从标准输出读取日志（ `-u` 设置服务名称， `-f` 跟踪日志）。

```bash
$ sudo systemctl stop myapp.service
```

To stop the service, for a maintenance (for instance).

​	停止服务，以便进行维护（例如）。

## Enable the Service on Startup 在启动时启用服务 

Once the service is working as you want, you can enable it on startup. This way, if the server need to restart, your app will restart as well.

​	一旦服务按您的需要工作，您就可以在启动时启用它。这样，如果服务器需要重新启动，您的应用程序也将重新启动。

```bash
$ sudo systemctl enable myapp.service
```
