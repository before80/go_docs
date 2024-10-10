+++
title = "生成文档前的替换"
date = 2023-06-05T09:37:31+08:00
description = ""
isCJKLanguage = true
draft = false

+++

# 生成文档前的替换

```
{{< ref "">}}

// 在浏览器控制台中执行以下代码
var elements = document.querySelectorAll('p > font > font > font');

elements.forEach(function(element) {        
        var newline = document.createElement('br');
        element.parentNode.insertBefore(newline, element);        
        element.firstChild.nodeValue = '&zeroWidthSpace;' +element.firstChild.nodeValue;
});


```

## 更换标题层级

```
//以下正则表达式请根据实际需要进行调整
^#{4} func

//替换成
##### func

^#{4} \(\*
//替换成
##### (*
```



## 替换掉标题中的链接

```
// 查找匹配如下正则表达式
(#{1,6} )\[([^\n]+)\]\([^\n]+\)

// 替换成
$1$2
```



## 替换掉`&zeroWidthSpace;`

```
// 查找匹配如下字符
&zeroWidthSpace;

// 替换成 tab
​	
```

## 替换掉youtube视频的HTML

```
// 查找匹配如下正则表达式
<iframe src="https:\/\/www\.youtube\.com\/embed\/([\w\-]+)([\w=\?]?)" [^\n]+><\/iframe>

或 

<iframe src="https:\/\/www\.youtube\.com\/embed\/([\w\-]+)\?[\w=]+" [^\n]+><\/iframe>

// 替换成
{{< youtube "$1">}}

```



## 替换掉标题中的重复章节数字

```
// 查找匹配如下字符
(\d+\.\d+ – )([a-zA-z\s]+ )(\d+\.\d+ –)
// 替换成
$1$2

// 查找匹配如下字符
(\d+\.\d+\.\d+ – )([a-zA-z\s]+ )(\d+\.\d+\.\d+ –)
// 替换成
$1$2
```



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
) added in

)added in
或
) added in v

// 替换成
) <-  

// 查找匹配如下字符
added in go1
// 替换成
<- go1
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
$1$3$4
```



## 替换掉details嵌入的作为Example的html

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


// 分别在```后加上go


````

## 修改原文版本

```
// 查找匹配如下正则表达式
(> 原文：)\[([^\]@]+)@(go1\.\d{2}\.\d)\]\(\2@\3\)

// 替换成 类似如下的版本
$1[$2@go1.23.0]($2@go1.23.0)
```

