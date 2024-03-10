+++
title = "go help env"
date = 2024-02-24T10:56:10+08:00
type = "docs"
weight = 520
description = ""
isCJKLanguage = true
draft = false

+++



usage: go env [-json] [-u] [-w] [var ...]

Env prints Go environment information.

​	Env 打印 Go 环境信息。

By default env prints information as a shell script (on Windows, a batch file). If one or more variable names is given as arguments, env prints the value of each named variable on its own line.

​	默认情况下，env 将信息打印为 shell 脚本（在 Windows 上为批处理文件）。如果将一个或多个变量名作为参数给出，env 将在单独的行上打印每个已命名变量的值。

The -json flag prints the environment in JSON format instead of as a shell script.

​	-json 标志以 JSON 格式打印环境，而不是作为 shell 脚本。

The -u flag requires one or more arguments and unsets the default setting for the named environment variables, if one has been set with ‘go env -w’.

​	-u 标志需要一个或多个参数，并且取消已使用“go env -w”设置的已命名环境变量的默认设置（如果已设置）。

The -w flag requires one or more arguments of the form NAME=VALUE and changes the default settings of the named environment variables to the given values.

​	-w 标志需要一个或多个 NAME=VALUE 形式的参数，并将已命名环境变量的默认设置更改为给定值。

For more about environment variables, see ‘go help environment’.

​	有关环境变量的更多信息，请参阅“go help environment”。