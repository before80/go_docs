+++
title = "使用 Apache 部署"
date = 2024-02-04T09:13:44+08:00
weight = 5
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://beego.wiki/docs/deploy/apache/]({{< ref "/beego/deployment/deploymentWithApache" >}})

# Deployment with Apache 使用 Apache 部署



Apache is a web server and is set up to perform the same functions as nginx, serving as a reverse proxy and sending requests to the backend. Here is a configuration example:

​	Apache 是一个 Web 服务器，其设置与 nginx 的功能相同，用作反向代理并将请求发送到后端。这是一个配置示例：

```
NameVirtualHost *:80
<VirtualHost *:80>
	ServerAdmin webmaster@dummy-host.example.com
	ServerName www.a.com
	ProxyRequests Off
	<Proxy *>
		Order deny,allow
		Allow from all
	</Proxy>
	ProxyPass / http://127.0.0.1:8080/
	ProxyPassReverse / http://127.0.0.1:8080/
</VirtualHost>
 
<VirtualHost *:80>
	ServerAdmin webmaster@dummy-host.example.com
	ServerName www.b.com
	ProxyRequests Off
	<Proxy *>
		Order deny,allow
		Allow from all
	</Proxy>
	ProxyPass / http://127.0.0.1:8081/
	ProxyPassReverse / http://127.0.0.1:8081/
</VirtualHost>
```
