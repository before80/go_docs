+++
title = "语句"
weight = 60
date = 2023-06-12T09:45:26+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false
+++

# 语句

## select

> 注意
>
> ​	select语句只能用于通道；
>
> ​	空select语句将会永久阻塞；
>
> ​	select语句中的case语句在读通道时，该通道不会发生阻塞，即使是nil通道也不会被阻塞。因case语句在编译后调用读通道时，会被传入不阻塞的参数。
>
> ​	