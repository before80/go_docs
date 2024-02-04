+++
title = "本地化"
date = 2024-02-04T21:18:22+08:00
weight = 11
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/guides/localization/]({{< ref "/buffalo/guides/localization" >}})

# Localization 本地化 

Translating your app is an effective way to **make it understandable to the many people** around the globe! Buffalo uses the [go-i18n](https://github.com/nicksnyder/go-i18n) project to provide the i18n (adapting the software to make it translatable without code change) and l10n (providing translation strings and specific formats) of your app.

​	翻译您的应用是让全球各地的人们理解它的有效方法！Buffalo 使用 go-i18n 项目来提供应用的 i18n（调整软件使其可翻译，无需更改代码）和 l10n（提供翻译字符串和特定格式）。

## Markup the Translatable Strings 标记可翻译字符串 

**This document only applies when using [github.com/gobuffalo/buffalo/render](https://github.com/gobuffalo/buffalo/tree/main/render). Please see [github.com/gobuffalo/plush](https://github.com/gobuffalo/plush) for more details on the underlying templating package.
此文档仅适用于使用 github.com/gobuffalo/buffalo/render 时。有关底层模板包的更多详细信息，请参阅 github.com/gobuffalo/plush。**

Translatable strings must use a **specific markup** to allow the l10n engine to get the translations.

​	可翻译字符串必须使用特定标记，以允许 l10n 引擎获取翻译。

In a plush template, you can use the `t` helper:

​	在 plush 模板中，您可以使用 `t` 帮助器：

```plain
<%= t("greetings") %>
```

### Translation Context 翻译上下文 

Since **0.10.2**
自 0.10.2 起



You can use a context with the `t` helper, to provide variables to the translation string:

​	您可以将上下文与 `t` 帮助器一起使用，以向翻译字符串提供变量：

```plain
<%= t("name-format", {name: "Mark"}) %>
```

### Plural Handling 复数处理 

Since **0.10.2**
自 0.10.2 起



You can use this helper with a numeric second arg to handle singular/plural cases:

​	您可以使用此助手和一个数字第二个参数来处理单数/复数情况：

```plain
<%= t("messages-notification", notificationsLen) %>
```

Provide a context using a third arg:

​	使用第三个参数提供上下文：

```plain
<%= t("messages-notification", notificationsLen, ctx) %>
```

The second arg is accessible as “Count” in the translations strings.
第二个参数在翻译字符串中可作为“计数”访问。

## Provide Translations 提供翻译 

Translations are stored in the `locales` folder. By default, they are stored in a `all.en-us.yaml` file for the American English strings.

​	翻译存储在 `locales` 文件夹中。默认情况下，它们存储在 `all.en-us.yaml` 文件中，用于美式英语字符串。

You can provide translations for another language by providing a new file `all.my-language-code.yaml`. If you want to split your strings into logical modules, you can even create multiples files, e.g. `users.en-us.yaml` for the user-related stuff, and `all.en-us.yaml` for the global stuff.

​	您可以通过提供新文件 `all.my-language-code.yaml` 来提供另一种语言的翻译。如果您想将字符串拆分为逻辑模块，您甚至可以创建多个文件，例如 `users.en-us.yaml` 用于与用户相关的内容， `all.en-us.yaml` 用于全局内容。

The localization format used by [go-i18n](https://github.com/nicksnyder/go-i18n) is the following:

​	go-i18n 使用的本地化格式如下：

```yaml
- id: greetings
  translation: "Welcome to Buffalo (EN)"

- id: messages-notification
  translation:
    one: "You have {{.Count}} notification"
    other: "You have {{.Count}} notifications"
```

## Define the Default Language 定义默认语言 

To define the default language of your app, you need to edit the `app.go` file in the `actions` folder:

​	要定义应用程序的默认语言，您需要编辑 `app.go` 文件夹中的 `actions` 文件：

```go
// Setup and use translations:
var err error
if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
  app.Stop(err)
}
app.Use(T.Middleware())
```

Changing `"en-US"` to another language code will change the default language.

​	将 `"en-US"` 更改为另一种语言代码将更改默认语言。

## Localized Views 本地化视图 

Since **0.10.2**
自 0.10.2 起



Sometimes, you have to **translate a whole page**, and marking every part of the page takes a lot of time. On some other cases, you’ll want to localize the page in a different way for a specific locale. Localized views is a complementary way to handle your translations.

​	有时，您必须翻译整个页面，而标记页面的每一部分都需要花费大量时间。在某些其他情况下，您会希望以不同的方式为特定区域设置本地化页面。本地化视图是处理翻译的补充方式。

Localized views are **included in the i18n middleware**, so you don’t need to setup anything else to use them.

​	本地化视图包含在 i18n 中间件中，因此您不需要设置其他任何内容即可使用它们。

### Create suffixed versions of the templates 创建模板的后缀版本 

First, create a version for the default locale, without a lang suffix:

​	首先，为默认区域设置创建一个版本，没有语言后缀：

**page.html**:

​	page.html：

```html
<p>This is my default language page.</p>
```

Then, create a new suffixed version for each language you want to support:

​	然后，为要支持的每种语言创建一个新的后缀版本：

**page.en-us.html**:

​	page.en-us.html：

```html
<p>This is my en-US version.</p>
```

**page.fr-fr.html**:

​	page.fr-fr.html：

```html
<p>This is my fr-FR version.</p>
```

The middleware will detect the user language and choose the right template for you! It also works with guest users, using the `Accept-Language` HTTP header.

​	中间件将检测用户语言并为您选择正确的模板！它还可与访客用户配合使用，使用 `Accept-Language` HTTP 标头。

## Use i18n in Actions 在操作中使用 i18n 

You’ll need to use the i18n features in actions, for instance, to translate flash messages. Here is the way to use it:

​	您需要在操作中使用 i18n 功能，例如，翻译闪存消息。以下是如何使用它：

```go
func Login(c buffalo.Context) error {
  // [...]
  // Set a translated flash message
  c.Flash().Add("success", T.Translate(c, "users.login-success"))
  return c.Redirect(303, "/users/signin")
}
```

`T.Translate` takes the `buffalo.Context` as first argument, then the following args are the same as the `t` helper ones (`t` calls `T.Translate` with the context, behind the scene).

​	 `T.Translate` 将 `buffalo.Context` 作为第一个参数，然后后面的参数与 `t` 帮助程序参数相同（ `t` 在幕后使用上下文调用 `T.Translate` ）。

## Refresh Translation Context 刷新翻译上下文 

Since **0.12.0**
自 0.12.0 起



If you provide translated versions of your app, you’ll probably have a language switch function. This way, the users can choose the correct language. Buffalo can’t detect when you change the language in an action, since it will extract the user languages once per request. You’ll then have to redirect to another page to see the changes. But even with that trick, if you use a flash message inside the action, the language used will be the old one.

​	如果您提供应用程序的翻译版本，您可能会有一个语言切换功能。这样，用户可以选择正确的语言。Buffalo 无法检测到您何时在操作中更改语言，因为它每次请求只提取一次用户语言。然后，您必须重定向到另一个页面才能看到更改。但即使使用此技巧，如果您在操作中使用闪存消息，使用的语言仍将是旧语言。

To solve that problem, you can use the `T.Refresh` method and refresh the language used for translations, within an action.

​	要解决此问题，您可以在操作中使用 `T.Refresh` 方法并刷新用于翻译的语言。

```go
func SwitchLanguage(c buffalo.Context) error {
  f := struct {
    Language string `form:"lang"`
    URL      string `form:"url"`
  }{}
  if err := c.Bind(&f); err != nil {
    return errors.WithStack(err)
  }

  // Set new current language using a cookie, for instance
  cookie := http.Cookie{
    Name:   "lang",
    Value:  f.Language,
    MaxAge: int((time.Hour * 24 * 265).Seconds()),
    Path:   "/",
  }
  http.SetCookie(c.Response(), &cookie)

  // Update language for the flash message
  T.Refresh(c, f.Language)

  c.Flash().Add("success", T.Translate(c, "users.language-changed", f))

  return c.Redirect(302, f.URL)
}
```

## Customize Generated Names 自定义生成名称 

Since **0.10.2**
自 0.10.2 起



Many Buffalo generators use [gobuffalo/flect](https://github.com/gobuffalo/flect) to generate a normalized version of a name. For example, when you want to generate a new model, the name you give to the command line is normalized in plural, capitalized, and so on forms.

​	许多 Buffalo 生成器使用 gobuffalo/flect 生成名称的规范化版本。例如，当您想要生成一个新模型时，您在命令行中给出的名称会以复数、大写等形式规范化。

Sometimes, the rules used by **flect** are not correct (in this case, feel free to open a PR on the repo!). Sometimes a rule is not correct for your use case, but it’s still correct in a general rule. In this case, you can provide custom rules using the `inflections.json` file at the root of your project.

​	有时，flect 使用的规则不正确（在这种情况下，请随时在仓库中打开一个 PR！）。有时，某个规则不适用于您的用例，但它仍然是一个通用的正确规则。在这种情况下，您可以使用项目根目录中的 `inflections.json` 文件提供自定义规则。

**inflections.json:
inflections.json：**

```json
{
  "singular form": "plural form"
}
```

## Related Resources 相关资源 # 翻译 Buffalo 应用 - 一篇关于使用 Buffalo i18n 工具的文章。

- [Translating a Buffalo app](https://blog.gobuffalo.io/translating-a-buffalo-app-1b4f32e6cb57) - An article about using Buffalo i18n tools.
