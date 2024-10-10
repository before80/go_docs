+++
title = "使用 Systemctl 部署"
date = 2024-02-04T09:13:21+08:00
weight = 3
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/deploy/systemctl/]({{< ref "/beego/deployment/deploymentWithSystemctl" >}})

# Deployment with Systemctl 使用 Systemctl 部署



Systemctl command is a command used to manage and control services. It allows you to enable, disable, view, start, stop, or restart system services.

​	Systemctl 命令用于管理和控制服务。它允许您启用、禁用、查看、启动、停止或重新启动系统服务。

## Install beego application as a service 将 beego 应用程序安装为服务

1. Pack your application using `bee pack` command. Copy the resultant `.tar.gz` file to target server.

   ​	使用 `bee pack` 命令打包您的应用程序。将生成的 `.tar.gz` 文件复制到目标服务器。

2. Unpack 
   ​	解压

   ```
    mkdir -p /usr/local/beepkg && cd "$_"
    tar -xvzf *.tar.gz
    rm -rf *.tar.gz
   ```

3. Configure service parameters

   ​	配置服务参数

   ```
    cat <<'EOF' > /etc/systemd/system/beepkg.service
    [Unit]
    Description=beepkg
    AssertPathExists=/usr/local/beepkg
   
    [Service]
    WorkingDirectory=/usr/local/beepkg
    ExecStart=/usr/local/beepkg/beepkg
   
    ExecReload=/bin/kill -HUP $MAINPID
    LimitNOFILE=65536
    Restart=always
    RestartSec=5
   
    [Install]
    WantedBy=multi-user.target
    EOF
   ```

4. Install service 
   ​	安装服务

   ```
    chmod +x beepkg
    chmod 644 /etc/systemd/system/beepkg.service
    systemctl daemon-reload
    systemctl enable beepkg
    systemctl start beepkg
    systemctl status beepkg
   ```
