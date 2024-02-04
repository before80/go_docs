+++
title = "使用 Supervisord 进行部署"
date = 2024-02-04T09:13:06+08:00
weight = 2
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/deploy/supervisor/]({{< ref "/beego/deployment/deploymentWithSupervisord" >}})

# Deployment with Supervisord 使用 Supervisord 进行部署



Supervisord is a very useful process manager implemented in Python. Supervisord can change your non-daemon application into a daemon application. The application needs to be a non-daemon app. So if you want to use Supervisord to manage nginx, you need to set daemon off to run nginx in non-daemon mode.

​	Supervisord 是一个用 Python 实现的非常有用的进程管理器。Supervisord 可以将非守护进程应用程序更改为守护进程应用程序。该应用程序需要是非守护进程应用程序。因此，如果您想使用 Supervisord 来管理 nginx，则需要将守护进程设置为关闭，以在非守护进程模式下运行 nginx。

## Install Supervisord 安装 Supervisord

1. install setuptools 
   ​	安装 setuptools

   ```
    wget https://pypi.python.org/packages/2.7/s/setuptools/setuptools-0.6c11-py2.7.egg
   
    sh setuptools-0.6c11-py2.7.egg 
   
    easy_install supervisor
   
    echo_supervisord_conf >/etc/supervisord.conf
   
    mkdir /etc/supervisord.conf.d
   ```

2. config `/etc/supervisord.conf` 
   ​	配置 `/etc/supervisord.conf`

   ```
    [include]
    files = /etc/supervisord.conf.d/*.conf
   ```

3. Create new application to be managed

   ​	创建要管理的新应用程序

   ```
    cd /etc/supervisord.conf.d
    vim beepkg.conf
   ```

   Configurations： 
   ​	配置：

   ```
    [program:beepkg]
    directory = /opt/app/beepkg
    command = /opt/app/beepkg/beepkg
    autostart = true
    startsecs = 5
    user = root
    redirect_stderr = true
    stdout_logfile = /var/log/supervisord/beepkg.log
   ```

## Supervisord Manage Supervisord 管理

Supervisord provides two commands, supervisord and supervisorctl:

​	Supervisord 提供两个命令，supervisord 和 supervisorctl：

- supervisord: Initialize Supervisord, run configed processes
  supervisord：初始化 Supervisord，运行已配置的进程
- supervisorctl stop programxxx: Stop process programxxx. programxxx is configed name in [program:beepkg]. Here is beepkg.
  supervisorctl stop programxxx：停止进程 programxxx。programxxx 是 [program:beepkg] 中配置的名称。此处为 beepkg。
- supervisorctl start programxxx: Run the process.
  supervisorctl start programxxx：运行进程。
- supervisorctl restart programxxx: Restart the process.
  supervisorctl restart programxxx：重新启动进程。
- supervisorctl stop groupworker: Restart all processes in group groupworker
  supervisorctl stop groupworker：重新启动 groupworker 组中的所有进程
- supervisorctl stop all: Stop all processes. Notes: start, restart and stop won’t reload the latest configs.
  supervisorctl stop all：停止所有进程。注意：start、restart 和 stop 不会重新加载最新的配置。
- supervisorctl reload: Reload the latest configs.
  supervisorctl reload：重新加载最新的配置。
- supervisorctl update: Reload all the processes who’s config has changed.
  supervisorctl update：重新加载所有已更改配置的进程。

> Notes: The processes stopped by `stop` manually won’t restart after reload or update.
>
> ​	注意：手动通过 `stop` 停止的进程在重新加载或更新后不会重新启动。
