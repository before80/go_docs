+++
title = "数据类型"
date = 2024-08-19T09:33:07+08:00
weight = 35
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++



## Go的内置数据类型

| 序号 | 数据类型   | 名称       | 别名 | 默认值                      | 占用字节数 | 数据最小值                                                   | 数据最大值                                                   | 备注                                                         | 链接 |
| ---- | ---------- | ---------- | ---- | --------------------------- | ---------- | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ---- |
| 1    | bool       | 布尔型     |      | false                       | 1          | 无                                                           | 无                                                           | 只有true和false两种值                                        |      |
| 2    | byte       | 字节型     |      | '\x00'（使用%q）0（使用%d） | 1          | '\x00'（使用%q）0（使用%d）                                  | '\x7f'（使用%q）127（使用%d）                                | 是uint8的别名                                                |      |
| 3    | complex64  | 复数型     |      | (0+0i)                      | 8          | 无                                                           | 无                                                           |                                                              |      |
| 4    | complex128 | 复数型     |      | (0+0i)                      | 16         | 无                                                           | 无                                                           |                                                              |      |
| 5    | float32    | 浮点型     |      | 0                           | 4          | 1.401298464324817070923729583289916131280e-45 最小正非零值（使用%.39e保留39位小数） | 3.40282346638528859811704183484516925440e+38 （使用%.38e保留38位小数） |                                                              |      |
| 6    | float64    | 浮点型     |      | 0                           | 8          | 4.94065645841246544176568792868221372365059803e-324 最小正非零值（使用%.44e保留44位小数） | 1.797693134862315708145274237317043567980706e+308 （使用%.42e保留42位小数） |                                                              |      |
| 7    | int8       | 有符号整型 |      | 0                           | 1          | -128                                                         | 127                                                          |                                                              |      |
| 8    | int16      | 有符号整型 |      | 0                           | 2          | -32768                                                       | 32767                                                        |                                                              |      |
| 9    | int32      | 有符号整型 | rune | 0                           | 4          | -2147483648                                                  | 2147483647                                                   |                                                              |      |
| 10   | int64      | 有符号整型 |      | 0                           | 8          | -9223372036854775808                                         | 9223372036854775807（>922亿亿）                              |                                                              |      |
| 11   | int        | 有符号整型 |      | 0                           | 8          | -9223372036854775808                                         | 9223372036854775807（>922亿亿）                              | 请注意：这里给出的是64位系统的情况！                         |      |
| 12   | uint8      | 无符号整型 | byte | 0                           | 1          | 0                                                            | 255                                                          |                                                              |      |
| 13   | uint16     | 无符号整型 |      | 0                           | 2          | 0                                                            | 65535                                                        |                                                              |      |
| 14   | uint32     | 无符号整型 |      | 0                           | 4          | 0                                                            | 4294967295                                                   |                                                              |      |
| 15   | uint64     | 无符号整型 |      | 0                           | 8          | 0                                                            | 18446744073709551615                                         |                                                              |      |
| 16   | uint       | 无符号整型 |      | 0                           | 8          | 0                                                            | 18446744073709551615                                         |                                                              |      |
| 17   | rune       | 符文型     |      | '\x00'（使用%q）0（使用%d） | 4          | '\x00'（使用%q）0（使用%d）                                  |                                                              | 是int32的别名，而非uint32的别名                              |      |
| 18   | uintptr    | 指针整数型 |      | 无                          |            |                                                              |                                                              | uintptr 是一个整数类型，它足够大，可以容纳任何指针的比特模式 |      |
| 19   | string     | 字符串型   |      | ""                          |            |                                                              |                                                              |                                                              |      |
| 20   | [n]T       | 数组       |      | 空数组                      |            |                                                              |                                                              |                                                              |      |
| 21   | []T        | 切片       |      | nil                         |            |                                                              |                                                              |                                                              |      |
| 22   | map[K]V    | 映射       |      | 无                          |            |                                                              |                                                              |                                                              |      |
| 23   | struct{}   | 结构体     |      | 各自字段的零值              |            |                                                              |                                                              |                                                              |      |
| 24   | chan       | 通道       |      | nil                         |            |                                                              |                                                              |                                                              |      |
| 25   | *T         | 指针       |      | 无                          |            |                                                              |                                                              |                                                              |      |
| 26   | interface  | 接口       |      | 无                          |            |                                                              |                                                              |                                                              |      |
| 27   | error      | 错误类型   |      | 无                          |            |                                                              |                                                              |                                                              |      |
