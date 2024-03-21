+++
title = "gres"
date = 2024-03-21T17:56:51+08:00
weight = 1
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：https://pkg.go.dev/github.com/gogf/gf/v2@v2.6.4/os/gres

Package gres provides resource management and packing/unpacking feature between files and bytes.

### Constants 

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gres/gres_instance.go#L11)

``` go
const (
	// DefaultName default group name for instance usage.
	DefaultName = "default"
)
```

[View Source](https://github.com/gogf/gf/blob/v2.6.4/os/gres/gres.go#L10)

``` go
const (
	// Separator for directories.
	Separator = "/"
)
```

### Variables 

This section is empty.

### Functions 

##### func Add 

``` go
func Add(content string, prefix ...string) error
```

Add unpacks and adds the `content` into the default resource object. The unnecessary parameter `prefix` indicates the prefix for each file storing into current resource object.

##### func Contains 

``` go
func Contains(path string) bool
```

Contains checks whether the `path` exists in the default resource object.

##### func Dump 

``` go
func Dump()
```

Dump prints the files of the default resource object.

##### func Export 

``` go
func Export(src, dst string, option ...ExportOption) error
```

Export exports and saves specified path `src` and all its sub files to specified system path `dst` recursively.

##### func GetContent 

``` go
func GetContent(path string) []byte
```

GetContent directly returns the content of `path` in default resource object.

##### func IsEmpty 

``` go
func IsEmpty() bool
```

IsEmpty checks and returns whether the resource manager is empty.

##### func Load 

``` go
func Load(path string, prefix ...string) error
```

Load loads, unpacks and adds the data from `path` into the default resource object. The unnecessary parameter `prefix` indicates the prefix for each file storing into current resource object.

<details class="Documentation-deprecatedDetails js-deprecatedDetails" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; display: block; color: var(--color-text-subtle);"><summary style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; list-style: none; opacity: 1;"><h4 tabindex="-1" id="Pack" data-kind="function" class="Documentation-functionHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 600; font-stretch: inherit; line-height: 1.25em; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1.125rem; margin: 1.5rem 0px 0.5rem; padding: 0px; vertical-align: baseline; word-break: break-word; align-items: baseline; display: flex; justify-content: space-between;"><span class="Documentation-deprecatedTitle" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; align-items: center; display: flex; gap: 0.5rem;">func<a class="Documentation-source" href="https://github.com/gogf/gf/blob/v2.6.4/os/gres/gres_func.go#L49" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 1;">Pack</a><span class="Documentation-deprecatedTag" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: 1.375; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.75rem; margin: 0px; padding: 0.125rem 0.25rem; vertical-align: middle; background-color: var(--color-border); border-radius: 0.125rem; color: var(--color-text-inverted); text-transform: uppercase;">DEPRECATED</span><span class="Documentation-deprecatedBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.87rem; margin: 0px 0.5rem 0px 0.25rem; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></span><span class="Documentation-sinceVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.9375rem; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></h4></summary><div class="go-Message go-Message--warning Documentation-deprecatedItemBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 1rem 1rem 0.5rem; vertical-align: baseline; color: var(--gray-1); width: 1164.69px; background-color: var(--color-background-warning);"><div class="Documentation-declaration" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre-wrap; scroll-padding-top: calc(var(--js-sticky-header-height, 3.5rem) + .75rem); word-break: break-all; overflow-wrap: break-word;"><a href="https://pkg.go.dev/builtin#string" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/builtin#string" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/builtin#byte" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/builtin#error" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></pre></div><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p></div></details>

<details class="Documentation-deprecatedDetails js-deprecatedDetails" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; display: block; color: var(--color-text-subtle);"><summary style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; list-style: none; opacity: 1;"><h4 tabindex="-1" id="PackToFile" data-kind="function" class="Documentation-functionHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 600; font-stretch: inherit; line-height: 1.25em; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1.125rem; margin: 1.5rem 0px 0.5rem; padding: 0px; vertical-align: baseline; word-break: break-word; align-items: baseline; display: flex; justify-content: space-between;"><span class="Documentation-deprecatedTitle" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; align-items: center; display: flex; gap: 0.5rem;">func<a class="Documentation-source" href="https://github.com/gogf/gf/blob/v2.6.4/os/gres/gres_func.go#L77" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 1;">PackToFile</a><span class="Documentation-deprecatedTag" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: 1.375; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.75rem; margin: 0px; padding: 0.125rem 0.25rem; vertical-align: middle; background-color: var(--color-border); border-radius: 0.125rem; color: var(--color-text-inverted); text-transform: uppercase;">DEPRECATED</span><span class="Documentation-deprecatedBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.87rem; margin: 0px 0.5rem 0px 0.25rem; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></span><span class="Documentation-sinceVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.9375rem; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></h4></summary><div class="go-Message go-Message--warning Documentation-deprecatedItemBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 1rem 1rem 0.5rem; vertical-align: baseline; color: var(--gray-1); width: 1164.69px; background-color: var(--color-background-warning);"><div class="Documentation-declaration" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre-wrap; scroll-padding-top: calc(var(--js-sticky-header-height, 3.5rem) + .75rem); word-break: break-all; overflow-wrap: break-word;"><a href="https://pkg.go.dev/builtin#string" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/builtin#string" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/builtin#error" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></pre></div><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p></div></details>

##### func PackToFileWithOption <-2.2.1

``` go
func PackToFileWithOption(srcPaths, dstPath string, option Option) error
```

PackToFileWithOption packs the path specified by `srcPaths` to target file `dstPath`.

Note that parameter `srcPaths` supports multiple paths join with ','.

<details class="Documentation-deprecatedDetails js-deprecatedDetails" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; display: block; color: var(--color-text-subtle);"><summary style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 16px; margin: 0px; padding: 0px; vertical-align: baseline; list-style: none; opacity: 1;"><h4 tabindex="-1" id="PackToGoFile" data-kind="function" class="Documentation-functionHeader" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 600; font-stretch: inherit; line-height: 1.25em; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1.125rem; margin: 1.5rem 0px 0.5rem; padding: 0px; vertical-align: baseline; word-break: break-word; align-items: baseline; display: flex; justify-content: space-between;"><span class="Documentation-deprecatedTitle" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; align-items: center; display: flex; gap: 0.5rem;">func<a class="Documentation-source" href="https://github.com/gogf/gf/blob/v2.6.4/os/gres/gres_func.go#L105" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 18px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none; opacity: 1;">PackToGoFile</a><span class="Documentation-deprecatedTag" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: 1.375; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.75rem; margin: 0px; padding: 0.125rem 0.25rem; vertical-align: middle; background-color: var(--color-border); border-radius: 0.125rem; color: var(--color-text-inverted); text-transform: uppercase;">DEPRECATED</span><span class="Documentation-deprecatedBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.87rem; margin: 0px 0.5rem 0px 0.25rem; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></span><span class="Documentation-sinceVersion" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: 400; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.9375rem; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle);"></span></h4></summary><div class="go-Message go-Message--warning Documentation-deprecatedItemBody" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 1rem 1rem 0.5rem; vertical-align: baseline; color: var(--gray-1); width: 1164.69px; background-color: var(--color-background-warning);"><div class="Documentation-declaration" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline;"><pre style="box-sizing: border-box; border: var(--border); font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5em; font-family: SFMono-Regular, Consolas, &quot;Liberation Mono&quot;, Menlo, monospace; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 0.875rem; margin: 0px; padding: 0.625rem; vertical-align: baseline; background-color: var(--color-background-accented); border-radius: var(--border-radius); color: var(--color-text); overflow-x: auto; tab-size: 4; white-space: pre-wrap; scroll-padding-top: calc(var(--js-sticky-header-height, 3.5rem) + .75rem); word-break: break-all; overflow-wrap: break-word;"><a href="https://pkg.go.dev/builtin#string" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/builtin#string" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a><a href="https://pkg.go.dev/builtin#error" style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: inherit; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 14px; margin: 0px; padding: 0px; vertical-align: baseline; color: var(--color-text-subtle); text-decoration: none;"></a></pre></div><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p><p style="box-sizing: border-box; border: 0px; font-style: inherit; font-variant: inherit; font-weight: inherit; font-stretch: inherit; line-height: 1.5rem; font-family: inherit; font-optical-sizing: inherit; font-kerning: inherit; font-feature-settings: inherit; font-variation-settings: inherit; font-size: 1rem; margin: 1rem 0px; padding: 0px; vertical-align: baseline; max-width: 60rem;"></p></div></details>

##### func PackToGoFileWithOption <-2.2.1

``` go
func PackToGoFileWithOption(srcPath, goFilePath, pkgName string, option Option) error
```

PackToGoFileWithOption packs the path specified by `srcPaths` to target go file `goFilePath` with given package name `pkgName`.

Note that parameter `srcPaths` supports multiple paths join with ','.

##### func PackWithOption <-2.2.1

``` go
func PackWithOption(srcPaths string, option Option) ([]byte, error)
```

PackWithOption packs the path specified by `srcPaths` into bytes.

Note that parameter `srcPaths` supports multiple paths join with ','.

### Types 

#### type ExportOption 

``` go
type ExportOption struct {
	RemovePrefix string // Remove the prefix of file name from resource.
}
```

ExportOption is the option for function Export.

#### type File 

``` go
type File struct {
	// contains filtered or unexported fields
}
```

##### func Get 

``` go
func Get(path string) *File
```

Get returns the file with given path.

##### func GetWithIndex 

``` go
func GetWithIndex(path string, indexFiles []string) *File
```

GetWithIndex searches file with `path`, if the file is directory it then does index files searching under this directory.

GetWithIndex is usually used for http static file service.

##### func ScanDir 

``` go
func ScanDir(path string, pattern string, recursive ...bool) []*File
```

ScanDir returns the files under the given path, the parameter `path` should be a folder type.

The pattern parameter `pattern` supports multiple file name patterns, using the ',' symbol to separate multiple patterns.

It scans directory recursively if given parameter `recursive` is true.

##### func ScanDirFile 

``` go
func ScanDirFile(path string, pattern string, recursive ...bool) []*File
```

ScanDirFile returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

Note that it returns only files, exclusive of directories.

##### func Unpack 

``` go
func Unpack(path string) ([]*File, error)
```

Unpack unpacks the content specified by `path` to []*File.

##### func UnpackContent 

``` go
func UnpackContent(content string) ([]*File, error)
```

UnpackContent unpacks the content to []*File.

##### (*File) Close 

``` go
func (f *File) Close() error
```

Close implements interface of http.File.

##### (*File) Content 

``` go
func (f *File) Content() []byte
```

Content returns the content of the file.

##### (*File) Export <-2.1.2

``` go
func (f *File) Export(dst string, option ...ExportOption) error
```

Export exports and saves all its sub files to specified system path `dst` recursively.

##### (*File) FileInfo 

``` go
func (f *File) FileInfo() os.FileInfo
```

FileInfo returns an os.FileInfo for the FileHeader.

##### (File) MarshalJSON 

``` go
func (f File) MarshalJSON() ([]byte, error)
```

MarshalJSON implements the interface MarshalJSON for json.Marshal.

##### (*File) Name 

``` go
func (f *File) Name() string
```

Name returns the name of the file.

##### (*File) Open 

``` go
func (f *File) Open() (io.ReadCloser, error)
```

Open returns a ReadCloser that provides access to the File's contents. Multiple files may be read concurrently.

##### (*File) Read 

``` go
func (f *File) Read(b []byte) (n int, err error)
```

Read implements the io.Reader interface.

##### (*File) Readdir 

``` go
func (f *File) Readdir(count int) ([]os.FileInfo, error)
```

Readdir implements Readdir interface of http.File.

##### (*File) Seek 

``` go
func (f *File) Seek(offset int64, whence int) (n int64, err error)
```

Seek implements the io.Seeker interface.

##### (*File) Stat 

``` go
func (f *File) Stat() (os.FileInfo, error)
```

Stat implements Stat interface of http.File.

#### type Option <-2.2.1

``` go
type Option struct {
	Prefix   string // The file path prefix for each file item in resource manager.
	KeepPath bool   // Keep the passed path when packing, usually for relative path.
}
```

Option contains the extra options for Pack functions.

#### type Resource 

``` go
type Resource struct {
	// contains filtered or unexported fields
}
```

##### func Instance 

``` go
func Instance(name ...string) *Resource
```

Instance returns an instance of Resource. The parameter `name` is the name for the instance.

##### func New 

``` go
func New() *Resource
```

New creates and returns a new resource object.

##### (*Resource) Add 

``` go
func (r *Resource) Add(content string, prefix ...string) error
```

Add unpacks and adds the `content` into current resource object. The unnecessary parameter `prefix` indicates the prefix for each file storing into current resource object.

##### (*Resource) Contains 

``` go
func (r *Resource) Contains(path string) bool
```

Contains checks whether the `path` exists in current resource object.

##### (*Resource) Dump 

``` go
func (r *Resource) Dump()
```

Dump prints the files of current resource object.

##### (*Resource) Export 

``` go
func (r *Resource) Export(src, dst string, option ...ExportOption) error
```

Export exports and saves specified path `srcPath` and all its sub files to specified system path `dstPath` recursively.

##### (*Resource) Get 

``` go
func (r *Resource) Get(path string) *File
```

Get returns the file with given path.

##### (*Resource) GetContent 

``` go
func (r *Resource) GetContent(path string) []byte
```

GetContent directly returns the content of `path`.

##### (*Resource) GetWithIndex 

``` go
func (r *Resource) GetWithIndex(path string, indexFiles []string) *File
```

GetWithIndex searches file with `path`, if the file is directory it then does index files searching under this directory.

GetWithIndex is usually used for http static file service.

##### (*Resource) IsEmpty 

``` go
func (r *Resource) IsEmpty() bool
```

IsEmpty checks and returns whether the resource manager is empty.

##### (*Resource) Load 

``` go
func (r *Resource) Load(path string, prefix ...string) error
```

Load loads, unpacks and adds the data from `path` into current resource object. The unnecessary parameter `prefix` indicates the prefix for each file storing into current resource object.

##### (*Resource) ScanDir 

``` go
func (r *Resource) ScanDir(path string, pattern string, recursive ...bool) []*File
```

ScanDir returns the files under the given path, the parameter `path` should be a folder type.

The pattern parameter `pattern` supports multiple file name patterns, using the ',' symbol to separate multiple patterns.

It scans directory recursively if given parameter `recursive` is true.

Note that the returned files does not contain given parameter `path`.

##### (*Resource) ScanDirFile 

``` go
func (r *Resource) ScanDirFile(path string, pattern string, recursive ...bool) []*File
```

ScanDirFile returns all sub-files with absolute paths of given `path`, It scans directory recursively if given parameter `recursive` is true.

Note that it returns only files, exclusive of directories.