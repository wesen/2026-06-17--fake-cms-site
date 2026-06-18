---
title: "Configuration — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/config/
summary: "Downloaded with defuddle. Config API: addGlobalData, addFilter, addShortcode, setTemplateFormats. Re-run: defuddle parse https://www.11ty.dev/docs/config/ --md | fold -w 100 -s."
---
- Introduction
	- [Why Eleventy?](https://www.11ty.dev/#why-should-you-use-eleventy)
		- [Performance](https://www.11ty.dev/docs/performance/)
	- Learn
	- [Starter Projects](https://www.11ty.dev/docs/starter/)
	- [Tutorials](https://www.11ty.dev/docs/tutorials/)
		- [Quick Tips](https://www.11ty.dev/docs/quicktips/)
- Community
	- [How can I contribute?](https://www.11ty.dev/docs/community/)
	- [Code of Conduct](https://www.11ty.dev/docs/code-of-conduct/)
	- [Blog](https://www.11ty.dev/blog/)
	- [Firehose](https://www.11ty.dev/firehose/)
	- [11ty Bundle](https://11tybundle.dev/)
	- [Leaderboards](https://www.11ty.dev/speedlify/)
	- [Eleventy Meetup](https://11tymeetup.dev/)
	- [11ty Conference](https://conf.11ty.dev/)
- [Guide](https://www.11ty.dev/docs/projects/) Guide
	- [Command Line Usage](https://www.11ty.dev/docs/usage/)
	- [Add a Configuration File](https://www.11ty.dev/docs/config/)
	- [Copy Files to Output](https://www.11ty.dev/docs/copy/)
	- [Add CSS, JS, Fonts](https://www.11ty.dev/docs/assets/)
	- [Importing Content](https://www.11ty.dev/docs/migrate/)
	- [Configure Templates with Data](https://www.11ty.dev/docs/data-configuration/)
	- [Using Data in Templates](https://www.11ty.dev/docs/data/)
		- [Eleventy Supplied Data](https://www.11ty.dev/docs/data-eleventy-supplied/)
			- [Data Cascade](https://www.11ty.dev/docs/data-cascade/)
			- [Front Matter Data](https://www.11ty.dev/docs/data-frontmatter/)
				- [Custom Front 
Matter](https://www.11ty.dev/docs/data-frontmatter-customize/)
					- [Template & Directory Data 
Files](https://www.11ty.dev/docs/data-template-dir/)
					- [Global Data 
Files](https://www.11ty.dev/docs/data-global/)
					- [Config Global 
Data](https://www.11ty.dev/docs/data-global-custom/)
					- [Computed Data](https://www.11ty.dev/docs/data-computed/)
			- [JavaScript Data Files](https://www.11ty.dev/docs/data-js/)
			- [Custom Data File Formats](https://www.11ty.dev/docs/data-custom/)
			- [Validate Data](https://www.11ty.dev/docs/data-validate/)
	- [Template Languages](https://www.11ty.dev/docs/languages/)
		- [HTML](https://www.11ty.dev/docs/languages/html/)
			- [Markdown](https://www.11ty.dev/docs/languages/markdown/)
			- [MDX](https://www.11ty.dev/docs/languages/mdx/)
			- [JavaScript](https://www.11ty.dev/docs/languages/javascript/)
			- [JSX](https://www.11ty.dev/docs/languages/jsx/)
					- 
[TypeScript](https://www.11ty.dev/docs/languages/typescript/)
			- [Custom](https://www.11ty.dev/docs/languages/custom/)
			- [WebC](https://www.11ty.dev/docs/languages/webc/)
			- [Nunjucks](https://www.11ty.dev/docs/languages/nunjucks/)
			- [Liquid](https://www.11ty.dev/docs/languages/liquid/)
			- [Handlebars](https://www.11ty.dev/docs/languages/handlebars/)
			- [Mustache](https://www.11ty.dev/docs/languages/mustache/)
			- [EJS](https://www.11ty.dev/docs/languages/ejs/)
			- [HAML](https://www.11ty.dev/docs/languages/haml/)
			- [Pug](https://www.11ty.dev/docs/languages/pug/)
			- [Sass](https://www.11ty.dev/docs/languages/sass/)
			- [Virtual Templates](https://www.11ty.dev/docs/virtual-templates/)
			- [Overriding Languages](https://www.11ty.dev/docs/template-overrides/)
	- Template Features
		- [Ignore Files](https://www.11ty.dev/docs/ignores/)
			- [Preprocess Content](https://www.11ty.dev/docs/config-preprocessors/)
			- [Postprocess Content](https://www.11ty.dev/docs/transforms/)
			- [Filters](https://www.11ty.dev/docs/filters/)
			- [`url`](https://www.11ty.dev/docs/filters/url/)
					- [`slugify`](https://www.11ty.dev/docs/filters/slugify/)
					- [`log`](https://www.11ty.dev/docs/filters/log/)
					- 
[`get*CollectionItem`](https://www.11ty.dev/docs/filters/collection-items/)
					- 
[`inputPathToUrl`](https://www.11ty.dev/docs/filters/inputpath-to-url/)
			- [Shortcodes](https://www.11ty.dev/docs/shortcodes/)
			- [`getBundle`](https://www.11ty.dev/docs/plugins/bundle/)
					- 
[`getBundleFileUrl`](https://www.11ty.dev/docs/plugins/bundle/)
	- [Environment Variables](https://www.11ty.dev/docs/environment-vars/)
	- [Internationalization (i18n)](https://www.11ty.dev/docs/i18n/)
	- [Watch Files and Dev Servers](https://www.11ty.dev/docs/watch-serve/)
		- [Eleventy Dev Server](https://www.11ty.dev/docs/dev-server/)
			- [Vite](https://www.11ty.dev/docs/server-vite/)
	- [Common Pitfalls](https://www.11ty.dev/docs/pitfalls/)
	- [Advanced](https://www.11ty.dev/docs/advanced/)
		- [Release History](https://www.11ty.dev/docs/versions/)
			- [Programmatic API](https://www.11ty.dev/docs/programmatic/)
			- [Configuration Events](https://www.11ty.dev/docs/events/)
			- [Order of Operations](https://www.11ty.dev/docs/advanced-order/)
- [Plugins](https://www.11ty.dev/docs/plugins/) Plugins
	- [Create or use Plugins](https://www.11ty.dev/docs/create-plugin/)
	- [Image](https://www.11ty.dev/docs/plugins/image/)
	- [Fetch](https://www.11ty.dev/docs/plugins/fetch/)
	- [`<is-land>`](https://www.11ty.dev/docs/plugins/is-land/)
	- [Render](https://www.11ty.dev/docs/plugins/render/)
	- [Internationalization (i18n)](https://www.11ty.dev/docs/plugins/i18n/)
	- [RSS](https://www.11ty.dev/docs/plugins/rss/)
	- [Upgrade Helper](https://www.11ty.dev/docs/plugins/upgrade-help/)
	- [Syntax Highlighting](https://www.11ty.dev/docs/plugins/syntaxhighlight/)
	- [InputPath to URL](https://www.11ty.dev/docs/plugins/inputpath-to-url/)
	- [Navigation](https://www.11ty.dev/docs/plugins/navigation/)
	- [HTML `<base>`](https://www.11ty.dev/docs/plugins/html-base/)
	- [Bundle](https://www.11ty.dev/docs/plugins/bundle/)
	- [Id Attribute](https://www.11ty.dev/docs/plugins/id-attribute/)
	- [Community Plugins](https://www.11ty.dev/docs/plugins/community/)
	- [Retired Plugins](https://www.11ty.dev/docs/plugins/retired/)
- [Services](https://www.11ty.dev/docs/services/) Services
	- [Deployment & Hosting](https://www.11ty.dev/docs/deployment/)
	- [Using a CMS](https://www.11ty.dev/docs/cms/)
	- [Runtime APIs](https://www.11ty.dev/docs/api-services/)
		- [Screenshots](https://www.11ty.dev/docs/services/screenshots/)
			- [OpenGraph Image](https://www.11ty.dev/docs/services/opengraph/)
			- [IndieWeb Avatar](https://www.11ty.dev/docs/services/indieweb-avatar/)
			- [Generator Image](https://www.11ty.dev/docs/services/generator/)
			- [Hosting Image](https://www.11ty.dev/docs/services/builtwith/)
			- [Sparklines](https://www.11ty.dev/docs/services/sparklines/)

Configuration files are optional. Add an `eleventy.config.js` file to the root directory of your 
project (read more about [default configuration filenames](#default-filenames)) to configure 
Eleventy to your own project’s needs. It might look like this:

eleventy.config.js

```js
export default async function(eleventyConfig) {
    // Configure Eleventy
};
```

```js
module.exports = async function(eleventyConfig) {
    // Configure Eleventy
};
```

There are a few different ways to [shape your configuration 
file](https://www.11ty.dev/docs/config-shapes/). Added in v3.0.0Eleventy v3 added support for both 
ESM and Asynchronous callbacks.

- Add [Filters](https://www.11ty.dev/docs/filters/).
- Add [Shortcodes](https://www.11ty.dev/docs/shortcodes/).
- Add [Custom Tags](https://www.11ty.dev/docs/custom-tags/).
- Add [JavaScript Template 
Functions](https://www.11ty.dev/docs/languages/javascript/#javascript-template-functions)
- Add custom [Collections](https://www.11ty.dev/docs/collections/) and use [Advanced Collection 
Filtering and 
Sorting](https://www.11ty.dev/docs/collections/#advanced-custom-filtering-and-sorting).
- Add [Plugins](https://www.11ty.dev/docs/plugins/).

Is your config file getting big and hard to understand? You can [create a project-specific 
plugin](https://www.11ty.dev/docs/quicktips/local-plugin/) to better organize your code.

## Default filenames

We look for the following configuration files:

1. `.eleventy.js`
2. `eleventy.config.js` Added in v2.0.0
3. `eleventy.config.mjs` Added in v3.0.0
4. `eleventy.config.cjs` Added in v2.0.0

The first configuration file found is used. The others are ignored.

## Configuration Options

### Input Directory

Controls the top level directory/file/glob that we’ll use to look for templates.

| Input Directory |  |
| --- | --- |
| *Object Key* | `dir.input` |
| *Default Value* | `.` *(current directory)* |
| *Valid Options* | Any valid directory. |
| *Configuration API* | `eleventyConfig.setInputDirectory()` Added in v3.0.0 |
| *Command Line Override* | `--input` |

#### Command Line

```bash
# The current directory
npx @11ty/eleventy --input=.

# A single file
npx @11ty/eleventy --input=README.md

# A glob of files
npx @11ty/eleventy --input=*.md

# A subdirectory
npx @11ty/eleventy --input=views
```

#### Configuration

Via named export (order doesn’t matter). Note that there are many [different shapes of 
configuration file](https://www.11ty.dev/docs/config-shapes/).

eleventy.config.js

```js
export const config = {
  dir: {
    input: "views"
  }
};
```

```js
module.exports.config = {
  dir: {
    input: "views"
  }
};
```

Or via method (not available in plugins) Added in v3.0.0:

eleventy.config.js

```js
export default function(eleventyConfig) {
    // Order matters, put this at the top of your configuration file.
  eleventyConfig.setInputDirectory("views");
};
```

```js
module.exports = function(eleventyConfig) {
    // Order matters, put this at the top of your configuration file.
  eleventyConfig.setInputDirectory("views");
};
```

### Directory for Includes

The includes directory is meant for [Eleventy layouts](https://www.11ty.dev/docs/layouts/), include 
files, extends files, partials, or macros. These files will not be processed as full template 
files, but can be consumed by other templates.

| Includes Directory |  |
| --- | --- |
| *Object Key* | `dir.includes` |
| *Default* | `_includes` |
| *Valid Options* | Any valid directory inside of `dir.input` (an empty string `""` is supported) |
| *Configuration API* | `eleventyConfig.setIncludesDirectory()` Added in v3.0.0 |
| *Command Line Override* | *None* |

Via named export (order doesn’t matter). Note that there are many [different shapes of 
configuration file](https://www.11ty.dev/docs/config-shapes/).

eleventy.config.js

```js
export const config = {
  dir: {
        // ⚠️ This value is relative to your input directory.
    includes: "my_includes"
  }
};
```

```js
module.exports.config = {
  dir: {
        // ⚠️ This value is relative to your input directory.
    includes: "my_includes"
  }
};
```

Or via method (not available in plugins) Added in v3.0.0:

eleventy.config.js

```js
export default function(eleventyConfig) {
    // Order matters, put this at the top of your configuration file.
    // This is relative to your input directory!
  eleventyConfig.setIncludesDirectory("my_includes");
};
```

```js
module.exports = function(eleventyConfig) {
    // Order matters, put this at the top of your configuration file.
    // This is relative to your input directory!
  eleventyConfig.setIncludesDirectory("my_includes");
};
```

### Directory for Layouts (Optional)

This configuration option is optional but useful if you want your [Eleventy 
layouts](https://www.11ty.dev/docs/layouts/) to live outside of the [Includes 
directory](#directory-for-includes). Just like the [Includes directory](#directory-for-includes), 
these files will not be processed as full template files, but can be consumed by other templates.

WARNING

This setting **only applies** to Eleventy's language-agnostic 
[layouts](https://www.11ty.dev/docs/layouts/) (when defined in front matter or data files).

When using `{% extends %}`, Eleventy will **still search the `_includes` directory**. See [this 
note about existing templating 
features](https://www.11ty.dev/docs/layout-chaining/#addendum-about-existing-templating-features).

| Includes Directory |  |
| --- | --- |
| *Object Key* | `dir.layouts` |
| *Default* | *The value in `dir.includes`* |
| *Valid Options* | Any valid directory inside of `dir.input` (an empty string `""` is supported) |
| *Configuration API* | `eleventyConfig.setLayoutsDirectory()` Added in v3.0.0 |
| *Command Line Override* | *None* |

Via named export (order doesn’t matter). Note that there are many [different shapes of 
configuration file](https://www.11ty.dev/docs/config-shapes/).

eleventy.config.js

```js
export const config = {
  dir: {
    // These are both relative to your input directory!
    includes: "_includes",
    layouts: "_layouts",
  }
};
```

```js
module.exports.config = {
  dir: {
    // These are both relative to your input directory!
    includes: "_includes",
    layouts: "_layouts",
  }
};
```

Or via method (not available in plugins) Added in v3.0.0:

eleventy.config.js

```js
export default function(eleventyConfig) {
    // Order matters, put this at the top of your configuration file.
    // This is relative to your input directory!
  eleventyConfig.setLayoutsDirectory("_layouts");
};
```

```js
module.exports = function(eleventyConfig) {
    // Order matters, put this at the top of your configuration file.
    // This is relative to your input directory!
  eleventyConfig.setLayoutsDirectory("_layouts");
};
```

### Directory for Global Data Files

Controls the directory inside which the global data template files, available to all templates, can 
be found. Read more about [Global Data Files](https://www.11ty.dev/docs/data-global/).

| Data Files Directory |  |
| --- | --- |
| *Object Key* | `dir.data` |
| *Default* | `_data` |
| *Valid Options* | Any valid directory inside of `dir.input` |
| *Configuration API* | `eleventyConfig.setDataDirectory()` Added in v3.0.0 |
| *Command Line Override* | *None* |

Via named export (order doesn’t matter). Note that there are many [different shapes of 
configuration file](https://www.11ty.dev/docs/config-shapes/).

eleventy.config.js

```js
export const config = {
  dir: {
    // ⚠️ This value is relative to your input directory.
    data: "lore",
  }
};
```

```js
module.exports.config = {
  dir: {
    // ⚠️ This value is relative to your input directory.
    data: "lore",
  }
};
```

Or via method (not available in plugins) Added in v3.0.0:

eleventy.config.js

```js
export default function(eleventyConfig) {
    // Order matters, put this at the top of your configuration file.
  eleventyConfig.setDataDirectory("lore");
};
```

```js
module.exports = function(eleventyConfig) {
    // Order matters, put this at the top of your configuration file.
  eleventyConfig.setDataDirectory("lore");
};
```

### Output Directory

Controls the directory inside which the finished templates will be written to.

| Output Directory |  |
| --- | --- |
| *Object Key* | `dir.output` |
| *Default* | `_site` |
| *Valid Options* | Any string that will work as a directory name. Eleventy creates this if it 
doesn’t exist. |
| *Configuration API* | `eleventyConfig.setOutputDirectory()` Added in v3.0.0 |
| *Command Line Override* | `--output` |

#### Command Line

```bash
npx @11ty/eleventy --output=_site
```

#### Configuration

Via named export (order doesn’t matter). Note that there are many [different shapes of 
configuration file](https://www.11ty.dev/docs/config-shapes/).

eleventy.config.js

```js
export const config = {
  dir: {
        output: "dist",
  }
};
```

```js
module.exports.config = {
  dir: {
        output: "dist",
  }
};
```

Or via method (not available in plugins) Added in v3.0.0:

eleventy.config.js

```js
export default function(eleventyConfig) {
    // Order matters, put this at the top of your configuration file.
  eleventyConfig.setOutputDirectory("dist");
};
```

```js
module.exports = function(eleventyConfig) {
    // Order matters, put this at the top of your configuration file.
  eleventyConfig.setOutputDirectory("dist");
};
```

### Default template engine for Markdown files

Markdown files run through this template engine before transforming to HTML.

| Markdown Template Engine |  |
| --- | --- |
| *Object Key* | `markdownTemplateEngine` |
| *Default* | `liquid` |
| *Valid Options* | A valid [template engine short name](https://www.11ty.dev/docs/languages/) or 
`false` |
| *Command Line Override* | *None* |
| *Configuration API* | `setMarkdownTemplateEngine` Pre-release only: v4.0.0-alpha.1 |

eleventy.config.js

```js
export const config = {
  markdownTemplateEngine: "njk",
};
```

```js
module.exports.config = {
  markdownTemplateEngine: "njk",
};
```

There are many [different shapes of configuration file](https://www.11ty.dev/docs/config-shapes/).

### Default template engine for HTML files

HTML templates run through this template engine before transforming to (better) HTML.

| HTML Template Engine |  |
| --- | --- |
| *Object Key* | `htmlTemplateEngine` |
| *Default* | `liquid` |
| *Valid Options* | A valid [template engine short name](https://www.11ty.dev/docs/languages/) or 
`false` |
| *Command Line Override* | *None* |
| *Configuration API* | `setHtmlTemplateEngine` Pre-release only: v4.0.0-alpha.1 |

eleventy.config.js

```js
export const config = {
  htmlTemplateEngine: "njk",
};
```

```js
module.exports.config = {
  htmlTemplateEngine: "njk",
};
```

There are many [different shapes of configuration file](https://www.11ty.dev/docs/config-shapes/).

### Template Formats

Specify which types of templates should be transformed.

| Template Formats |  |
| --- | --- |
| *Object Key* | `templateFormats` |
| *Default* | `html,liquid,ejs,md,hbs,mustache,haml,pug,njk,11ty.js` |
| *Valid Options* | Array of [template engine short names](https://www.11ty.dev/docs/languages/) |
| *Command Line Override* | `--formats` *(accepts a comma separated string)* |
| *Configuration API* | `setTemplateFormats` and `addTemplateFormats` |

**Case sensitivity**: File extensions should be considered case insensitive, cross-platform. While 
macOS already behaves this way (by default), other operating systems require additional Eleventy 
code to enable this behavior.

#### Command Line

```
npx @11ty/eleventy --formats=html,liquid,njk
```

#### Configuration File Static Export

eleventy.config.js

```js
export const config = {
  templateFormats: ["html", "liquid", "njk"],
};
```

```js
module.exports.config = {
  templateFormats: ["html", "liquid", "njk"],
};
```

There are many [different shapes of configuration file](https://www.11ty.dev/docs/config-shapes/).

#### Configuration API

eleventy.config.js

```js
export default function (eleventyConfig) {
    // Reset to this value
    eleventyConfig.setTemplateFormats("html,liquid,njk");

    // Additive to existing
    eleventyConfig.addTemplateFormats("pug,haml");

    // Or:
    // eleventyConfig.setTemplateFormats([ "html", "liquid", "njk" ]);
    // eleventyConfig.addTemplateFormats([ "pug", "haml" ]);
};
```

```js
module.exports = function (eleventyConfig) {
    // Reset to this value
    eleventyConfig.setTemplateFormats("html,liquid,njk");

    // Additive to existing
    eleventyConfig.addTemplateFormats("pug,haml");

    // Or:
    // eleventyConfig.setTemplateFormats([ "html", "liquid", "njk" ]);
    // eleventyConfig.addTemplateFormats([ "pug", "haml" ]);
};
```

### Enable Quiet Mode to Reduce Console Noise

In order to maximize user-friendliness to beginners, Eleventy will show each file it processes and 
the output file. To disable this noisy console output, use quiet mode!

| Quiet Mode |  |
| --- | --- |
| *Default* | `false` |
| *Valid Options* | `true` or `false` |
| *Command Line Override* | `--quiet` |

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.setQuietMode(true);
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.setQuietMode(true);
};
```

The command line will override any setting in configuration:

```bash
npx @11ty/eleventy --quiet
```

### Deploy to a subdirectory with a Path Prefix

If your site lives in a different subdirectory (particularly useful with GitHub pages), use 
pathPrefix to specify this. When paired with the [HTML `<base>` 
plugin](https://www.11ty.dev/docs/plugins/html-base/) it will transform any absolute URLs in your 
HTML to include this folder name and does **not** affect where things go in the output folder.

| Path Prefix |  |
| --- | --- |
| *Object Key* | `pathPrefix` |
| *Default* | `/` |
| *Valid Options* | A prefix directory added to urls in HTML files |
| *Command Line Override* | `--pathprefix` |

eleventy.config.js

```js
import { HtmlBasePlugin } from "@11ty/eleventy";

export default function (eleventyConfig) {
    eleventyConfig.addPlugin(HtmlBasePlugin);
};

export const config = {
    pathPrefix: "/eleventy-base-blog/",
}
```

```js
module.exports = async function (eleventyConfig) {
    const { HtmlBasePlugin } = await import("@11ty/eleventy");

    eleventyConfig.addPlugin(HtmlBasePlugin);
};

module.exports.config = {
    pathPrefix: "/eleventy-base-blog/",
}
```

Deploy to https://11ty.github.io/eleventy-base-blog/ on GitHub pages without modifying your config. 
This allows you to use the same code-base to deploy to either GitHub pages or Netlify, like the 
[`eleventy-base-blog`](https://github.com/11ty/eleventy-base-blog) project does.

```bash
npx @11ty/eleventy --pathprefix=eleventy-base-blog
```

### Change Base File Name for Data Files

Added in v2.0.0 When using [Directory Specific Data 
Files](https://www.11ty.dev/docs/data-template-dir/), looks for data files that match the current 
folder name. You can override this behavior to a static string with the `setDataFileBaseName` 
method.

| File Suffix |  |
| --- | --- |
| *Configuration API* | `setDataFileBaseName` |
| *Default* | *Current folder name* |
| *Valid Options* | String |
| *Command Line Override* | *None* |

eleventy.config.js

```js
export default function (eleventyConfig) {
    // Looks for index.json and index.11tydata.json instead of using folder names
    eleventyConfig.setDataFileBaseName("index");
};
```

```js
module.exports = function (eleventyConfig) {
    // Looks for index.json and index.11tydata.json instead of using folder names
    eleventyConfig.setDataFileBaseName("index");
};
```

### Change File Suffix for Data Files

Added in v2.0.0 When using [Template and Directory Specific Data 
Files](https://www.11ty.dev/docs/data-template-dir/), to prevent file name conflicts with 
non-Eleventy files in the project directory, we scope these files with a unique-to-Eleventy suffix. 
This suffix is customizable using the `setDataFileSuffixes` configuration API method.

| File Suffix |  |
| --- | --- |
| *Configuration API* | `setDataFileSuffixes` |
| *Default* | `[".11tydata", ""]` |
| *Valid Options* | Array |
| *Command Line Override* | *None* |

For example, using `".11tydata"` will search for `*.11tydata.js` and `*.11tydata.json` data files. 
The empty string (`""`) here represents a file without a suffix—and this entry only applies to 
`*.json` data files.

This feature can also be used to disable Template and Directory Data Files altogether with an empty 
array (`[]`).

Read more about [Template and Directory Specific Data 
Files](https://www.11ty.dev/docs/data-template-dir/).

eleventy.config.js

```js
export default function (eleventyConfig) {
    // e.g. file.json and file.11tydata.json
    eleventyConfig.setDataFileSuffixes([".11tydata", ""]);

    // e.g. file.11tydata.json
    eleventyConfig.setDataFileSuffixes([".11tydata"]);

    // No data files are used.
    eleventyConfig.setDataFileSuffixes([]);
};
```

```js
module.exports = function (eleventyConfig) {
    // e.g. file.json and file.11tydata.json
    eleventyConfig.setDataFileSuffixes([".11tydata", ""]);

    // e.g. file.11tydata.json
    eleventyConfig.setDataFileSuffixes([".11tydata"]);

    // No data files are used.
    eleventyConfig.setDataFileSuffixes([]);
};
```

***Backwards Compatibility Note*** (`v2.0.0`)

Prior to v2.0.0 this feature was exposed using a `jsDataFileSuffix` property in the configuration 
return object. When the `setDataFileSuffixes` method has not been used, Eleventy maintains 
backwards compatibility for old projects by using this property as a fallback.

eleventy.config.js

```js
export default function (eleventyConfig) {
    return {
        jsDataFileSuffix: ".11tydata",
    };
};
```

```js
module.exports = function (eleventyConfig) {
    return {
        jsDataFileSuffix: ".11tydata",
    };
};
```

### Transforms

- Documented moved to [Transforms](https://www.11ty.dev/docs/transforms/).

### Linters

Similar to Transforms, Linters are provided to analyze a template’s output without modifying it.

| Linters |  |
| --- | --- |
| *Configuration API* | `addLinter` |
| *Object Key* | *N/A* |
| *Valid Options* | Callback function |
| *Command Line Override* | *None* |

eleventy.config.js

```js
export default function (eleventyConfig) {
    // Sync or async
    eleventyConfig.addLinter("linter-name", async function (content) {
        console.log(this.inputPath);
        console.log(this.outputPath);

        // Eleventy 2.0+ has full access to Eleventy’s \`page\` variable
        console.log(this.page.inputPath);
        console.log(this.page.outputPath);
    });
};
```

```js
module.exports = function (eleventyConfig) {
    // Sync or async
    eleventyConfig.addLinter("linter-name", async function (content) {
        console.log(this.inputPath);
        console.log(this.outputPath);

        // Eleventy 2.0+ has full access to Eleventy’s \`page\` variable
        console.log(this.page.inputPath);
        console.log(this.page.outputPath);
    });
};
```

**Linters Example: Use Inclusive Language**

Inspired by the [CSS Tricks post *Words to Avoid in Educational 
Writing*](https://css-tricks.com/words-avoid-educational-writing/), this linter will log a warning 
to the console when it finds a trigger word in a markdown file.

This example has been packaged as a plugin in 
[`eleventy-plugin-inclusive-language`](https://www.11ty.dev/docs/plugins/inclusive-language/).

```js
export default function (eleventyConfig) {
    eleventyConfig.addLinter(
        "inclusive-language",
        function (content, inputPath, outputPath) {
            let words =
                "simply,obviously,basically,of course,clearly,just,everyone 
knows,however,easy".split(
                    ","
                );

            // Eleventy 1.0+: use this.inputPath and this.outputPath instead
            if (inputPath.endsWith(".md")) {
                for (let word of words) {
                    let regexp = new RegExp("\\b(" + word + ")\\b", "gi");
                    if (content.match(regexp)) {
                        console.warn(
                            \`Inclusive Language Linter (${inputPath}) Found: ${word}\`
                        );
                    }
                }
            }
        }
    );
};
```

### Data Filter Selectors

Added in v1.0.0

A `Set` of [`lodash` selectors](https://lodash.com/docs/4.17.15#get) that allow you to include data 
from the data cascade in the output from `--to=json`, `--to=ndjson`.

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.dataFilterSelectors.add("page");
    eleventyConfig.dataFilterSelectors.delete("page");
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.dataFilterSelectors.add("page");
    eleventyConfig.dataFilterSelectors.delete("page");
};
```

This will now include a `data` property in your JSON output that includes the `page` variable for 
each matching template.

### TypeScript Type Definitions

This may enable some extra autocomplete features in your IDE (where supported).

eleventy.config.js

```js
/** @param {import("@11ty/eleventy").UserConfig} eleventyConfig */
export default function (eleventyConfig) {
    // …
};
```

```js
/** @param {import("@11ty/eleventy").UserConfig} eleventyConfig */
module.exports = function (eleventyConfig) {
    // …
};
```

- Related: [GitHub #2091](https://github.com/11ty/eleventy/pull/2091) and [GitHub 
#3097](https://github.com/11ty/eleventy/issues/3097)

### Removed Features

#### Change exception case suffix for HTML files

Feature Removal

The `htmlOutputSuffix` feature was removed in Eleventy 3.0. You can read about the feature on the 
[v2 
documentation](https://v2-0-1.11ty.dev/docs/config/#change-exception-case-suffix-for-html-files). 
Related: [GitHub #3327](https://github.com/11ty/eleventy/issues/3327).

### Documentation Moved to Dedicated Pages

#### Copy Files to Output using Passthrough File Copy

Files found (that don’t have a valid template engine) from opt-in file extensions in 
`templateFormats` will passthrough to the output directory. Read more about [Passthrough 
Copy](https://www.11ty.dev/docs/copy/).

#### Customize Front Matter Parsing Options

- Documented at [Customize Front Matter 
Parsing](https://www.11ty.dev/docs/data-frontmatter-customize/).

#### Watch JavaScript Dependencies

- Documented at [Watch and Serve Configuration](https://www.11ty.dev/docs/watch-serve/).

#### Add Your Own Watch Targets

- Documented at [Watch and Serve Configuration](https://www.11ty.dev/docs/watch-serve/).

#### Override Browsersync Server Options

- Documented at [Watch and Serve Configuration](https://www.11ty.dev/docs/watch-serve/).

#### Transforms

- Documented at [Transforms](https://www.11ty.dev/docs/transforms/).

---

### Other pages in Eleventy Projects

- [Command Line Usage](https://www.11ty.dev/docs/usage/)
- [Add a Configuration File](https://www.11ty.dev/docs/config/)
- [Copy Files to Output](https://www.11ty.dev/docs/copy/)
- [Add CSS, JS, Fonts](https://www.11ty.dev/docs/assets/)
- [Importing Content](https://www.11ty.dev/docs/migrate/)
- [Configure Templates with Data](https://www.11ty.dev/docs/data-configuration/)
- [Using Data in Templates](https://www.11ty.dev/docs/data/)
	- [Eleventy Supplied Data](https://www.11ty.dev/docs/data-eleventy-supplied/)
		- [Data Cascade](https://www.11ty.dev/docs/data-cascade/)
		- [Front Matter Data](https://www.11ty.dev/docs/data-frontmatter/)
			- [Custom Front 
Matter](https://www.11ty.dev/docs/data-frontmatter-customize/)
				- [Template & Directory Data 
Files](https://www.11ty.dev/docs/data-template-dir/)
				- [Global Data Files](https://www.11ty.dev/docs/data-global/)
				- [Config Global 
Data](https://www.11ty.dev/docs/data-global-custom/)
				- [Computed Data](https://www.11ty.dev/docs/data-computed/)
		- [JavaScript Data Files](https://www.11ty.dev/docs/data-js/)
		- [Custom Data File Formats](https://www.11ty.dev/docs/data-custom/)
		- [Validate Data](https://www.11ty.dev/docs/data-validate/)
- [Template Languages](https://www.11ty.dev/docs/languages/)
	- [HTML](https://www.11ty.dev/docs/languages/html/)
		- [Markdown](https://www.11ty.dev/docs/languages/markdown/)
		- [MDX](https://www.11ty.dev/docs/languages/mdx/)
		- [JavaScript](https://www.11ty.dev/docs/languages/javascript/)
		- [JSX](https://www.11ty.dev/docs/languages/jsx/)
				- [TypeScript](https://www.11ty.dev/docs/languages/typescript/)
		- [Custom](https://www.11ty.dev/docs/languages/custom/)
		- [WebC](https://www.11ty.dev/docs/languages/webc/)
		- [Nunjucks](https://www.11ty.dev/docs/languages/nunjucks/)
		- [Liquid](https://www.11ty.dev/docs/languages/liquid/)
		- [Handlebars](https://www.11ty.dev/docs/languages/handlebars/)
		- [Mustache](https://www.11ty.dev/docs/languages/mustache/)
		- [EJS](https://www.11ty.dev/docs/languages/ejs/)
		- [HAML](https://www.11ty.dev/docs/languages/haml/)
		- [Pug](https://www.11ty.dev/docs/languages/pug/)
		- [Sass](https://www.11ty.dev/docs/languages/sass/)
		- [Virtual Templates](https://www.11ty.dev/docs/virtual-templates/)
		- [Overriding Languages](https://www.11ty.dev/docs/template-overrides/)
- Template Features
	- [Ignore Files](https://www.11ty.dev/docs/ignores/)
		- [Preprocess Content](https://www.11ty.dev/docs/config-preprocessors/)
		- [Postprocess Content](https://www.11ty.dev/docs/transforms/)
		- [Filters](https://www.11ty.dev/docs/filters/)
		- [`url`](https://www.11ty.dev/docs/filters/url/)
				- [`slugify`](https://www.11ty.dev/docs/filters/slugify/)
				- [`log`](https://www.11ty.dev/docs/filters/log/)
				- 
[`get*CollectionItem`](https://www.11ty.dev/docs/filters/collection-items/)
				- 
[`inputPathToUrl`](https://www.11ty.dev/docs/filters/inputpath-to-url/)
		- [Shortcodes](https://www.11ty.dev/docs/shortcodes/)
		- [`getBundle`](https://www.11ty.dev/docs/plugins/bundle/)
				- [`getBundleFileUrl`](https://www.11ty.dev/docs/plugins/bundle/)
- [Environment Variables](https://www.11ty.dev/docs/environment-vars/)
- [Internationalization (i18n)](https://www.11ty.dev/docs/i18n/)
- [Watch Files and Dev Servers](https://www.11ty.dev/docs/watch-serve/)
	- [Eleventy Dev Server](https://www.11ty.dev/docs/dev-server/)
		- [Vite](https://www.11ty.dev/docs/server-vite/)
- [Common Pitfalls](https://www.11ty.dev/docs/pitfalls/)
- [Advanced](https://www.11ty.dev/docs/advanced/)
	- [Release History](https://www.11ty.dev/docs/versions/)
		- [Programmatic API](https://www.11ty.dev/docs/programmatic/)
		- [Configuration Events](https://www.11ty.dev/docs/events/)
		- [Order of Operations](https://www.11ty.dev/docs/advanced-order/)
