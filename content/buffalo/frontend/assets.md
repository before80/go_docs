+++
title = "assets"
date = 2024-02-04T21:11:28+08:00
weight = 9
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

> 原文：[https://gobuffalo.io/documentation/frontend-layer/assets/](https://gobuffalo.io/documentation/frontend-layer/assets/)

# Assets 资产 

By default when a new Buffalo application is generated via the `buffalo new` command, a [Webpack](https://webpack.github.io/) configuration file is generated, and the application is set up to use Webpack as the asset pipeline for the application.

​	默认情况下，当通过 `buffalo new` 命令生成新的 Buffalo 应用程序时，将生成一个 Webpack 配置文件，并且应用程序被设置为使用 Webpack 作为应用程序的资产管道。

If [`npm`](https://www.npmjs.com/) is not found on the machine generating the new Buffalo application, then Webpack will not be configured and the asset pipeline would be skipped.

​	如果在生成新 Buffalo 应用程序的机器上找不到 `npm` ，那么不会配置 Webpack，并且会跳过资产管道。

The asset pipeline can also be skipped during application generation with the `--skip-webpack` flag.

​	还可以在应用程序生成期间使用 `--skip-webpack` 标志跳过资产管道。

## JavaScript

The asset pipeline is initially configured to support ES6 JavaScript files, with `/assets/js/application.js` being the main entry point.

​	资产管道最初被配置为支持 ES6 JavaScript 文件，其中 `/assets/js/application.js` 是主要入口点。

The following are automatically installed and configured during setup of the asset pipeline:

​	在设置资产管道期间，以下内容会自动安装和配置：

- [jQuery](https://jquery.com/)
- [Bootstrap 4](http://getbootstrap.com/)
- [jQuery UJS](https://github.com/rails/jquery-ujs)

None of the installed packages are required, and may be removed. They are included for convenience.

​	安装的包都不是必需的，可以删除。它们是为了方便而包含的。

## CSS

By default the asset pipeline is configured to use [.scss](http://sass-lang.com/) files, with `/assets/css/application.scss` as the main entry point. This, of course, can be changed.

​	默认情况下，资产管道配置为使用 .scss 文件，其中 `/assets/css/application.scss` 是主入口点。当然，这可以更改。

The following are automatically installed and configured during setup of the asset pipeline:

​	在设置资产管道期间，以下内容会自动安装和配置：

- [Bootstrap 4](http://getbootstrap.com/)
- [Font Awesome](http://fontawesome.io/)

None of the installed packages are required, and may be removed. They are included for convenience.

​	安装的包都不是必需的，可以删除。它们是为了方便而包含的。

## Other Assets 其他资产 

Any assets placed in the `/assets` folder will be copied to the “distribution” automatically, and can be found at `/assets/path/to/asset`.

​	放置在 `/assets` 文件夹中的任何资产都会自动复制到“distribution”，可以在 `/assets/path/to/asset` 中找到。

## Asset Fingerprinting 资产指纹 

In `v0.9.5` asset fingerprinting was introduced to the default Webpack configuration for new applications. Asset fingerprinting works by generating a hash of the file contents and appending it to the name of the file. So, for example, `application.js` might be come `application.a8adff90f4c6d47529c4.js`. The benefit of this is that it allows for assets to be cached but still allow for that cache to be broken when a change has been made to the contents of this file.

​	在 `v0.9.5` 中，为新应用程序的默认 Webpack 配置引入了资产指纹。资产指纹通过生成文件内容的哈希并将其附加到文件名来工作。因此，例如， `application.js` 可能会变成 `application.a8adff90f4c6d47529c4.js` 。这样做的好处是，它允许缓存资产，但仍允许在对该文件的内容进行更改时中断该缓存。

Note that in order for this to work, buffalo will expect a `/public/assets/manifest.json` file to be present, containing the mappings between the files you reference in the helpers (eg `javascriptTag("application.js")`) and their hashed counterparts. This is not something you need to worry about if you are using the default Webpack configuration. However, if you choose to use the `--skip-webpack` flag when generating the project, keep in mind you will have to handle that yourself. While having the manifest file present is not strictly required for your application to run, you may experience caching problems without it during development.

​	请注意，为了使其正常工作，buffalo 将期望存在一个 `/public/assets/manifest.json` 文件，其中包含您在帮助器中引用的文件（例如 `javascriptTag("application.js")` ）与其哈希对应项之间的映射。如果您使用的是默认 Webpack 配置，则无需担心此问题。但是，如果您选择在生成项目时使用 `--skip-webpack` 标志，请记住您必须自己处理它。虽然存在清单文件并不是应用程序运行的严格要求，但在开发过程中没有它可能会遇到缓存问题。

**NOTE:** Applications written before `v0.9.5` may need to set an `AssetsBox` in their `actions/render.go` file in the `render.Options`, if assets are not rendering properly. It is recommended to move the one from the `actions/app.go` file into there instead. This will not setup asset finger printing, but will make sure the assets render correctly. See changes https://github.com/gobuffalo/docs/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-c1ebdbddf205da1687721a8acd29043cR43 and https://github.com/gobuffalo/docs/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-25015af78e14806bd828e39a29a403fbR13 for examples.

​	 `v0.9.5` ： `AssetsBox` 之前写的代码可能需要在 `actions/render.go` 文件中的 `render.Options` 中设置一个 `actions/app.go` ，如果资产没有正确地处理。建议将一个从文件中移动到那里。这不会设置手指，但会确保资产呈献。查看更改 https://github.com/gobuffalo/rat/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-c1ebdbb205da1687721a8acd29043cR43 和 https://github.com/gobuffalo/rat/commit/00673ab3446a9a7209bbd243e4594bd679c81a69#diff-25015af78e14806bd828e39a29a403cR13 以获取示例。

By default new applications are setup to fingerprint only JavaScript and CSS files.

​	默认情况下，新的应用程序仅设置为指纹和文件。

## Asset Helpers 

With the introduction of asset fingerprinting in `v0.9.5` it became difficult to find asset files because the name of the file kept changing. To help with this, three new helpers were introduced.

​	由于在 `v0.9.5` 中引入了指纹识别，因此很难找到文件，因为文件的名字不断变化。为了解决这个难题，引入了三个新的助手。

1. `assetPath` - This helper will return the path of the requested asset. For example, `<%= assetPath("application.js") %>` would return something like `/assets/application.a8adff90f4c6d47529c4.js`.

   \- 该助手将返回的路径。例如，会返回类似的内容。

2. `javascriptTag` - This helper will generate a `<script src="xxx"></script>` style tag for the requested JavaScript file. Example: `<%= javascriptTag("application.js") %>` would return something like `<script src="/assets/application.bd76587ded82386f388f.js" type="text/javascript"></script>`.

   \- 该助手将为文件生成样式的标签。示例：会返回类似的内容。

3. `stylesheetTag` - This helper will generate a `<link href="xxx">` style tag for the requested CSS file. Example: `<%= stylesheetTag("application.css") %>` would return something like `<link href="/assets/application.bd76587ded82386f388f.css" media="screen" rel="stylesheet" />`.

   ​	 `stylesheetTag` - 此帮助程序将为请求的 CSS 文件生成 `<link href="xxx">` 样式标记。示例： `<%= stylesheetTag("application.css") %>` 将返回类似 `<link href="/assets/application.bd76587ded82386f388f.css" media="screen" rel="stylesheet" />` 的内容。

## Building Assets in Development 在开发中构建资产 

The `buffalo dev` command, in addition to watching and rebuilding the application’s Go binary, will watch, and rebuild the asset pipeline as well. Nothing special needs to be run.

​	 `buffalo dev` 命令除了监视和重新构建应用程序的 Go 二进制文件外，还将监视并重新构建资产管道。无需运行任何特殊命令。

## Building Assets for Deployment 构建资产以进行部署 

The `buffalo build` command will build the asset pipeline, and properly attach it to the generated Go binary. One binary to run them all! See [Packing](https://gobuffalo.io/documentation/deploy/packing) for more options on building assets for deployment.

​	 `buffalo build` 命令将构建资产管道，并将其正确附加到生成的 Go 二进制文件。一个二进制文件即可运行所有内容！有关构建资产以进行部署的更多选项，请参阅打包。