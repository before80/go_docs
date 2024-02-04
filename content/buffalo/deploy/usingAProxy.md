+++
title = "使用代理"
date = 2024-02-04T21:20:13+08:00
weight = 4
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/deploy/proxy/](https://gobuffalo.io/documentation/deploy/proxy/)

# Using a Proxy 使用代理 

Buffalo can be used raw to handle queries. You just need to start your app and listen on a given address and port:

​	Buffalo 可以以原始方式用于处理查询。您只需启动应用程序并在给定的地址和端口上进行侦听：

```bash
# Env config
ADDR=0.0.0.0
PORT=80

# Start your app as a daemon, for example:
./myapp &
```

On common scenarios though, you’ll be using a proxy to distribute queries to a cluster, handle cases when your app is down, and so on.

​	然而，在常见场景中，您将使用代理将查询分发到集群，处理应用程序宕机的情况，等等。

## NGINX

NGINX allows two ways to be configured with your app:

​	NGINX 允许通过两种方式与您的应用程序进行配置：

### Using an IP address 使用 IP 地址 

#### Single backend app on same host 同一主机上的单个后端应用程序 

**app env config:
应用程序环境配置：**

```bash
ADDR=127.0.0.1
PORT=3000
```

**NGINX config:
NGINX 配置：**

```nginx
upstream buffalo_app {
    server 127.0.0.1:3000;
}

server {
    listen 80;
    server_name example.com;

    # Hide NGINX version (security best practice)
    server_tokens off;

    location / {
        proxy_redirect   off;
        proxy_set_header Host              $http_host;
        proxy_set_header X-Real-IP         $remote_addr;
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        proxy_pass       http://buffalo_app;
    }
}
```

#### Multiple backend apps 多个后端应用 

Using different ports just for the example:

​	仅为示例使用不同的端口：

**app1 env config:
app1 env 配置：**

```bash
ADDR=0.0.0.0
PORT=3000
```

**app2 env config:
app2 env 配置：**

```bash
ADDR=0.0.0.0
PORT=3001
```

**app3 env config:
app3 env 配置：**

```bash
ADDR=0.0.0.0
PORT=3002
```

**NGINX config:
NGINX 配置：**

```nginx
upstream buffalo_app_hosts {
    server host1.example.com:3000;
    server host2.example.com:3001;
    server host3.example.com:3002;
}

server {
    listen 80;
    server_name example.com;

    # Hide NGINX version (security best practice)
    server_tokens off;

    location / {
        proxy_redirect   off;
        proxy_set_header Host              $http_host;
        proxy_set_header X-Real-IP         $remote_addr;
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_pass       http://buffalo_app_hosts;
    }
}
```

### Using a UNIX domain socket 使用 UNIX 域套接字 

Since **0.10.3**
自 0.10.3 起



[UNIX sockets](https://en.wikipedia.org/wiki/Unix_domain_socket) are a common way to do inter-process communication (IPC) on UNIX systems. This means a program **A** can talk to a program **B**, using a file descriptor, just like they do using the TCP stack.

​	UNIX 套接字是 UNIX 系统上进行进程间通信 (IPC) 的常用方式。这意味着程序 A 可以使用文件描述符与程序 B 通信，就像它们使用 TCP 协议栈一样。

In our case, this allows you to have an instance of Buffalo running behind the proxy, without having to handle the full TCP stack between Buffalo and the proxy. This way, your app will answer faster!

​	在我们的案例中，这允许您在代理后面运行 Buffalo 实例，而无需在 Buffalo 和代理之间处理完整的 TCP 协议栈。这样，您的应用将更快地响应！

There are a couple of things to note about UNIX sockets. Since a UNIX socket is a file, UNIX file permissions apply. Therefore, whatever user owns the NGINX processes (typically `nginx`) needs to be able to both read from and write to the socket. **Executing `chmod 777` on the socket file will work, but this is almost always a bad idea!** Since, by default, groups have full read/write permissions on sockets created in Buffalo, a simpler and more secure solution would be to add the NGINX user to the user’s group that owns the app. The command to do this would be along the lines of `usermod -aG buffalo nginx`.

​	关于 UNIX 套接字，需要注意以下几点。由于 UNIX 套接字是一个文件，因此 UNIX 文件权限适用。因此，无论哪个用户拥有 NGINX 进程（通常是 `nginx` ），都需要能够对套接字进行读写。在套接字文件上执行 `chmod 777` 将起作用，但这几乎总是一个坏主意！由于默认情况下，组对在 Buffalo 中创建的套接字具有完全的读/写权限，因此一个更简单、更安全的方法是将 NGINX 用户添加到拥有该应用的用户组中。执行此操作的命令类似于 `usermod -aG buffalo nginx` 。

Socket files are typically created under the `/tmp` directory as in the example configuration below. However, in some more recent distributions of Linux, particularly newer [RedHat family](http://fedoraproject.org/wiki/Features/ServicesPrivateTmp) distros, `/tmp` and `/var/tmp` are namespaced so only the user that creates the file can see that it even exists. On these distributions, you will want to use something along the lines of `unix:/var/sock/buffalo.sock` instead of the example address given below.

​	套接字文件通常在 `/tmp` 目录下创建，如下面的示例配置所示。但是，在某些较新的 Linux 发行版中，尤其是较新的 RedHat 系列发行版中， `/tmp` 和 `/var/tmp` 是命名空间的，因此只有创建该文件的用户才能看到该文件的存在。在这些发行版上，您需要使用类似于 `unix:/var/sock/buffalo.sock` 的内容，而不是下面给出的示例地址。

**app env config:
应用环境配置：**

```bash
ADDR=unix:/tmp/buffalo.sock
```

**NGINX config:
NGINX 配置：**

```nginx
upstream buffalo_app {
    server unix:/tmp/buffalo.sock;
}

server {
    listen 80;
    server_name example.com;

    # Hide NGINX version (security best practice)
    server_tokens off;

    location / {
        proxy_redirect   off;
        proxy_set_header Host              $http_host;
        proxy_set_header X-Real-IP         $remote_addr;
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_pass       http://buffalo_app;
    }
}
```

## Apache 2

### Using an IP address 使用 IP 地址 

**app env config:
应用环境配置：**

```bash
ADDR=127.0.0.1
PORT=3000
```

**Apache 2 config:
Apache 2 配置：**

```apache
<VirtualHost *:80>
    ProxyPreserveHost On

    # Proxy requests to Buffalo
    ProxyPass / http://0.0.0.0:3000/
    ProxyPassReverse / http://0.0.0.0:3000/

    ServerName example.com
</VirtualHost>
```