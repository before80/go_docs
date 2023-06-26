+++
title = "生成文档前的替换"
date = 2023-06-05T09:37:31+08:00
description = ""
isCJKLanguage = true
draft = false

+++

# 生成文档前的替换

## 替换掉 ¶ 符

```
// 查找匹配如下正则表达式
\[¶\]\(\S+\)

// 替换成空

```

## 替换掉 added in v

```
// 查找匹配如下字符
)added in v
或
) added in v

// 替换成
) <-  
```



## 替换掉(函数的)标题中的链接

```
// 查找匹配如下正则表达式
(\#+ func )\[(\S+)\]\(\S+\)

// 替换成
$1$2
```

## 替换掉(类型的)标题中的链接

```
// 查找匹配如下正则表达式
(\#+ type\s)\[(\S+)\]\(\S+\)

// 替换成
$1$2
```

## 替换掉(方法的)标题中的链接

```
// 查找匹配如下正则表达式
(\#+ )(func )(\(\*\S+\)\s)\[(\S+)\]\(\S+\)

// 替换成
$1$3$4

//以及 查找匹配如下正则表达式
(\#+ )(func )(\(\S+\)\s)\[(\S+)\]\(\S+\)
// 替换成
```



## 替换掉details嵌入的作为Exampe的html

````
// 查找匹配如下正则表达式
\<details[^\n]+text-decoration: none;"\>(Example[^\n]*)(\<span\>)[^\n]+\<\/div\>\<\/details\>(\n)

// 替换成
##### $1$3
``` go
```
````



## 替换掉details嵌入的作为DEPRECATED的html

```
// 查找匹配如下正则表达式
\<details[^\n]+\<span class="Documentation\-deprecatedTitle" style="box\-sizing: border-box; border: 0px; font\-style: inherit; font\-variant: inherit; font\-weight: inherit; font\-stretch: inherit; line\-height: inherit; font\-family: inherit; font\-size: 18px; margin: 0px; padding: 0px; vertical\-align: baseline; align\-items: center; display: flex; gap: 0.5rem;"\>([^\n]+)\<a class="Documentation\-source" href="https:\/\/[^\n]+" style="box\-sizing: border\-box; border: 0px; font\-style: inherit; font\-variant: inherit; font\-weight: inherit; font\-stretch: inherit; line\-height: inherit; font\-family: inherit; font\-size: 18px; margin: 0px; padding: 0px; vertical\-align: baseline; color: var\(\-\-color\-text\-subtle\); text\-decoration: none; opacity: 1;"\>([^\n]+)\<\/a\>\<span class="Documentation\-deprecatedTag" style="box\-sizing: border\-box; border: 0px; font\-style: inherit; font\-variant: inherit; font\-weight: 400; font\-stretch: inherit; line\-height: 1.375; font\-family: inherit; font\-size: 0.75rem; margin: 0px; padding: 0.125rem 0.25rem; vertical\-align: middle; background\-color: var\(\-\-color\-border\); border\-radius: 0.125rem; color: var\(\-\-color\-text\-inverted\); text\-transform: uppercase;"\>(DEPRECATED)\<\/span\>[^\n]+\<\/div\>\<\/details\>(\n)

// 替换成
#### $1 $2 <- $3$4
```



## 增加代码块的语言标识

````
// 查找匹配如下字符

```
func

或
```
type

或
```
package

或
```
const

或
```
var

或
```
const

// 分别在```后加上go


````

