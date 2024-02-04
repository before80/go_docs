+++
title = "使用 nginx 部署"
date = 2024-02-04T09:13:32+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/deploy/nginx/](https://beego.wiki/docs/deploy/nginx/)

# Deployment with nginx 使用 nginx 部署



Go already has a standalone http server. But we still want to have nginx to do more for us such as logging, CC attack and act as a static file server because nginx performs well as a web server. So Go can just focus on functionality and logic. We can also use the nginx proxy to deploy multiple applications at the same time. Here is an example of two applications that share port 80 but have different domains, and requests are forwarding to different applications by nginx.

&zeroWidthSpace;Go 已经有一个独立的 http 服务器。但我们仍然希望 nginx 为我们做更多的事情，例如记录、CC 攻击和充当静态文件服务器，因为 nginx 作为 Web 服务器性能良好。因此，Go 只需专注于功能和逻辑即可。我们还可以使用 nginx 代理同时部署多个应用程序。以下是一个示例，其中两个应用程序共享端口 80 但具有不同的域，并且请求由 nginx 转发到不同的应用程序。

```
server {
    listen       80;
    server_name  .a.com;

    charset utf-8;
    access_log  /home/a.com.access.log  main;

    location /(css|js|fonts|img)/ {
        access_log off;
        expires 1d;

        root "/path/to/app_a/static"
        try_files $uri @backend
    }

    location / {
        try_files /_not_exists_ @backend;
    }

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass http://127.0.0.1:8080;
    }
}

server {
    listen       80;
    server_name  .b.com;

    charset utf-8;
    access_log  /home/b.com.access.log  main;

    location /(css|js|fonts|img)/ {
        access_log off;
        expires 1d;

        root "/path/to/app_b/static"
        try_files $uri @backend
    }

    location / {
        try_files /_not_exists_ @backend;
    }

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass http://127.0.0.1:8081;
    }
}
```