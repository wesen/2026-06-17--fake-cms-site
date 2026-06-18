---
title: "Custom Template Languages — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/languages/custom/
summary: "Downloaded with defuddle. addExtension + compile/getData for arbitrary template types. Re-run: defuddle parse https://www.11ty.dev/docs/languages/custom/ --md | fold -w 100 -s."
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

| Eleventy Short Name | File Extension | npm Package |
| --- | --- | --- |
| *(Any)* | `.*` *(Any)* | *(Any)* |

Eleventy now allows the addition of custom template extensions, meaning that you can use Eleventy 
to process any arbitrary file extension and compile it to your site’s output folder. This feature 
is Added in v1.0.0.

## Introductory Example: \*.clowd

`clowd` is a pretend templating language that we’ve just created. It uses the `.clowd` file 
extension. The purpose of the language is to translate any occurrences of the word `cloud` to the 
word `butt` instead.

eleventy.config.js

```js
export default function (eleventyConfig) {
    // Add as a valid extension to process
    // Alternatively, add this to the list of formats you pass to the \`--formats\` CLI argument
    eleventyConfig.addTemplateFormats("clowd");

    // "clowd" here means that the extension will apply to any .clowd file
    eleventyConfig.addExtension("clowd", {
        compile: async (inputContent) => {
            // Replace any instances of cloud with butt
            let output = inputContent.replace(/cloud/gi, "butt");

            return async () => {
                return output;
            };
        },
    });
};
```

```js
module.exports = function (eleventyConfig) {
    // Add as a valid extension to process
    // Alternatively, add this to the list of formats you pass to the \`--formats\` CLI argument
    eleventyConfig.addTemplateFormats("clowd");

    // "clowd" here means that the extension will apply to any .clowd file
    eleventyConfig.addExtension("clowd", {
        compile: async (inputContent) => {
            // Replace any instances of cloud with butt
            let output = inputContent.replace(/cloud/gi, "butt");

            return async () => {
                return output;
            };
        },
    });
};
```

Situations where you might want to use `addExtension` but probably shouldn’t:

1. If you want to post-process the content of an existing template language (a file extension 
already processed by Eleventy), use a [Configuration API 
Transform](https://www.11ty.dev/docs/config/#transforms) instead.
2. If you want to pre-process `md` or `html` files using another template language, change the 
*Default Template Engine for [Markdown 
Files](https://www.11ty.dev/docs/config/#default-template-engine-for-markdown-files)* or *[HTML 
Files](https://www.11ty.dev/docs/config/#default-template-engine-for-html-files)*, respectively. 
This can also be done on [a per-template basis](https://www.11ty.dev/docs/template-overrides/). We 
will likely add additional hooks for preprocessing in the future.

## Example: Add Sass support to Eleventy

For a more realistic sample, here’s an example of Eleventy looking for all `.scss` files in a 
project’s input directory to process them to your output directory.

eleventy.config.js

```js
// Don’t forget to \`npm install sass\`!
import * as sass from "sass";

export default function (eleventyConfig) {
    eleventyConfig.addTemplateFormats("scss");

    // Creates the extension for use
    eleventyConfig.addExtension("scss", {
        outputFileExtension: "css", // optional, default: "html"

        // \`compile\` is called once per .scss file in the input directory
        compile: async function (inputContent) {
            let result = sass.compileString(inputContent);

            // This is the render function, \`data\` is the full data cascade
            return async (data) => {
                return result.css;
            };
        },
    });
};
```

```js
// Don’t forget to \`npm install sass\`!
const sass = require("sass");

module.exports = function (eleventyConfig) {
    eleventyConfig.addTemplateFormats("scss");

    // Creates the extension for use
    eleventyConfig.addExtension("scss", {
        outputFileExtension: "css", // optional, default: "html"

        // \`compile\` is called once per .scss file in the input directory
        compile: async function (inputContent) {
            let result = sass.compileString(inputContent);

            // This is the render function, \`data\` is the full data cascade
            return async (data) => {
                return result.css;
            };
        },
    });
};
```

We’re using `compileString` from the Sass library above for speed benefits over their 
asynchronous counterparts (reported by [the Sass 
documentation](https://sass-lang.com/documentation/js-api#usage)).

Note also that the `data` is not used in the above example. This is the full Eleventy data cascade 
and may be more useful in other templating languages.

The above extension would process a file located at `subdir/test.scss` to the output directory at 
`_site/subdir/test.css`.

## Using inputPath

You can pass in both the file’s `inputPath` and the Eleventy includes folder to provide a set of 
directories to look for when using Sass’ `@use`, `@forward`, and `@import` features. Read more 
about [`loadPaths` on the Sass 
documentation](https://sass-lang.com/documentation/js-api/interfaces/Options#loadPaths).

**Filename** eleventy.config.js

```
// some configuration omitted…
+compile: function (inputContent, inputPath) {
+    let { dir } = path.parse(inputPath);

let result = sass.compileString(inputContent, {
+        loadPaths: [
+            dir || ".",
+            this.config.dir.includes
+        ]
    });

    return (data) => {
        return result.css;
    };
}
// …
```

Make special note of the `this.config.dir.includes` folder above. Declaring your includes folder 
means that you don’t need to prefix any file paths with the includes folder name (e.g. 
`_includes/_code.scss` can be consumed with `@use "code"`).

## Registering Dependencies Added in v2.0.0

Eleventy includes two features to improve the performance of custom template compilation:

1. A compilation cache, which you can optionally disable with 
[`compileOptions.cache`](#compileoptions.cache-for-advanced-control-of-caching)
2. Hooks for incremental builds (via the `--incremental` command line flag)

To facilitate these features, if a template syntax allows use of other templates (think `@use` in 
Sass or `webc:import` in WebC), Eleventy needs to know about the dependencies a template file 
relies on. This is heavily dependent on each template compiler.

In our Sass example, this is exposed by Sass via the [`loadedUrls` property from the 
`compileString` function](https://sass-lang.com/documentation/js-api/interfaces/CompileResult), and 
you can see an example of how we register our dependencies in the `compile` method below:

```
// some configuration omitted…
compile: function (inputContent, inputPath) {
    let result = sass.compileString(inputContent);

+    this.addDependencies(inputPath, result.loadedUrls);

    return async (data) => {
        return result.css;
    };
}
// …
```

`addDependencies` ’s first parameter is the parent template file path. The second parameter is an 
Array of child file paths used by the template. The dependencies can be either relative or absolute 
paths and we will normalize them as needed.

## Skipping a template from inside of the compile function

To add support for Sass’ underscore convention (file names that start with an underscore aren’t 
written to the output directory), just return early in the `compile` function (don’t return a 
`render` function).

```
// some configuration omitted…
compile: async function (inputContent, inputPath) {
    let parsed = path.parse(inputPath);
+    if(parsed.name.startsWith("_")) {
+        return;
+    }

    let result = sass.compileString(inputContent);

    return async (data) => {
        return result.css;
    };
}
// …
```

Note that files inside of the `_includes` folder are left out of processing by default, so if you 
store your sass `@use`, `@forward`, and `@import` files in there you’ll get this for free (see 
the [Using `inputPath` example](#using-inputpath) above)!

This functionality is more-or-less identical to the [`compileOptions` `permalink: false` 
overrides](#compile-options-permalink-to-override-permalink-compilation), documented later on this 
page.

## Aliasing an Existing Template Language

Added in v2.0.0 If `key` is the *only* property in the options object, we treat the extension as an 
alias and use the existing upstream template syntax.

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.addExtension("11ty.jsx", {
        key: "11ty.js",
    });

    // Or, you can pass an array of extensions in v2.0.0 or newer.
    eleventyConfig.addExtension(["11ty.jsx", "11ty.ts", "11ty.tsx"], {
        key: "11ty.js",
    });
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.addExtension("11ty.jsx", {
        key: "11ty.js",
    });

    // Or, you can pass an array of extensions in v2.0.0 or newer.
    eleventyConfig.addExtension(["11ty.jsx", "11ty.ts", "11ty.tsx"], {
        key: "11ty.js",
    });
};
```

You can read about the above approach (and see more detailed examples of its usage) on the 
[TypeScript](https://www.11ty.dev/docs/languages/typescript/), 
[JSX](https://www.11ty.dev/docs/languages/jsx/), or [MDX](https://www.11ty.dev/docs/languages/mdx/) 
documentation pages.

Added in v3.0.0 `key` needn’t be the only property in the options object. If you want to add your 
own `compile` function, [keep reading](#overriding-or-extending-an-existing-template-language)!

Added in v3.0.0 **Breaking Change**: Starting in Eleventy 3.0 you must add the new alias to [your 
declared template formats](https://www.11ty.dev/docs/config/#template-formats) for the new template 
type to be processed.

## Overriding or Extending an Existing Template Language

You can override or extend existing template languages too! (Thank you to [Ben Holmes for this 
contribution](https://github.com/11ty/eleventy/pull/1871)).

In these example, we switch from the Eleventy default `markdown-it` to `marked` for markdown 
processing.

eleventy.config.js

```js
import { marked } from "marked";

export default function (eleventyConfig) {
    eleventyConfig.addExtension("md", {
        compile: function (inputContent, inputPath) {
            let html = marked.parse(inputContent);

            return function (data) {
                // Example: use \`marked\` only if useMarked is set in the Data Cascade
                if (data.useMarked) {
                    return html;
                }

                // You can also access the default \`markdown-it\` renderer here:
                return this.defaultRenderer(data);
            };
        },
    });
};
```

```js
const { marked } = require("marked");

module.exports = function (eleventyConfig) {
    eleventyConfig.addExtension("md", {
        compile: function (inputContent, inputPath) {
            let html = marked.parse(inputContent);

            return function (data) {
                // Example: use \`marked\` only if useMarked is set in the Data Cascade
                if (data.useMarked) {
                    return html;
                }

                // You can also access the default \`markdown-it\` renderer here:
                return this.defaultRenderer(data);
            };
        },
    });
};
```

Note that overriding `md` opts-out of the default pre-processing by another template language 
[Markdown Files](https://www.11ty.dev/docs/config/#default-template-engine-for-markdown-files). As 
mentioned elsewhere, improvements to add additional hooks for preprocessing will likely come later.

You can override a template language once. Any attempts to override an more than once via 
`addExtension` will throw an error.

Added in v3.0.0 Adding `key` in the options object unlocks use of the target `defaultRenderer`. You 
can read about this approach (and see examples of its usage) on the 
[TypeScript](https://www.11ty.dev/docs/languages/typescript/), 
[JSX](https://www.11ty.dev/docs/languages/jsx/), or [MDX](https://www.11ty.dev/docs/languages/mdx/) 
documentation pages (all of which use `key: "11ty.js"` to extend 
[JavaScript](https://www.11ty.dev/docs/languages/javascript/) templates).

## Access to Existing Filters and Shortcodes

If you want to add support for universal filters and shortcodes in your custom template language, 
you can do so with the following configuration API methods. Related [GitHub 
#3310](https://github.com/11ty/eleventy/issues/3310).

- `eleventyConfig.getFilter(name)`
- `eleventyConfig.getFilters()` Added in v3.0.0
- `eleventyConfig.getShortcode(name)` Added in v3.0.0
- `eleventyConfig.getShortcodes()` Added in v3.0.0
- `eleventyConfig.getPairedShortcode(name)` Added in v3.0.0
- `eleventyConfig.getPairedShortcodes()` Added in v3.0.0

## Full Options List

### compile

- *Required* for new file extensions. *Optional* for 
[aliases](#aliasing-an-existing-template-language).

`compile` is an async-friendly function that takes two parameters:

- `inputContent`: the full content of the file to parse (as a string).
- `inputPath`: the path to the file (as a string, useful for looking up relative imports)

`compile` can return:

- nothing (`undefined`) to indicate that the file should be ignored and not used as a page
- a render function (also async-friendly)
```js
// some configuration truncated …
    compile: async (inputContent, inputPath) => {
        return async () => {
            return inputContent;
        };
    },
```

The render function is passed the merged data object (i.e. the full Data Cascade available inside 
templates). The render function returned from `compile` is called once per output file generated 
(one for basic templates and more for [paginated templates](https://www.11ty.dev/docs/pagination/)).

`inputContent` will **not include front matter**. This will have been parsed, removed, and inserted 
into the Data Cascade. Also note that if `read: false` (as documented below), `inputContent` will 
be `undefined`.

**Advanced: Adding Eleventy’s Scoped Data to your Compile Function**

Shortcodes and Filters both provide access to `page` and `eleventy` (via `this.page` and 
`this.eleventy` specifically). If you’d like to add the same for your custom template, you can do 
so via the `augmentFunctionContext` method.

```js
compile: function(compileFn) {
        return function(data) {
            // Binds this.page and this.eleventy to your render context (and any future additions 
added later)
            let renderFn = eleventyConfig.augmentFunctionContext(compileFn, {
                source: data,

                // Overwrite existing values?
                // overwrite: true,

                // Lazily fetch the key using \`getter\`
                // lazy: false,
                // getter: (key, context) => context?.[key];
            });

            return renderFn(data);
        };
    }
```

### outputFileExtension

- *Optional*: Defaults to `html`

When the output file is written to the file system, what file extension should be used?

### init

- *Optional*

An async-friendly function that runs *once* (no matter how many files use the extension) for any 
additional setup at the beginning before any compilation or rendering.

Note that `init` will **not** re-run on watch/serve mode. If you’d like something that runs 
before *every* build, use the [`eleventy.before` 
event](https://www.11ty.dev/docs/events/#eleventy.before).

```js
// some configuration truncated …
  init: async function() {
    // has access to current configuration settings in \`this.config\`
  },
```

### read

- *Optional*: Defaults to `true`

Set to `false` to opt out of reading the contents of files from the file system. This is useful if 
you’re using an external bundler to read the files.

```js
// some configuration truncated …
  read: false,
```

Use with `compileOptions.setCacheKey` to get more fine-grained control over how the template is 
cached.

### useLayouts Added in v3.0.0

- *Optional*: Defaults to `true`

Whether or not [Layouts](https://www.11ty.dev/docs/layouts/) will be applied to this template 
language. This will also exclude data from layout files to play a part in the data cascade of this 
template type as well. Related [GitHub #2830](https://github.com/11ty/eleventy/issues/2830).

### useJavaScriptImport Added in v3.0.0

Use the JavaScript loader instead of reading from the file system. If enabled, this takes 
precedence over `read` option.

```js
useJavaScriptImport: true,
    getInstanceFromInputPath: async function(inputPath) {
        let mod = await import(inputPath);
        return mod.default;
    },
    compile: (compileFn) => compileFn,
```

### getData and getInstanceFromInputPath

- *Optional*

Controls if and how additional data should be retrieved from a JavaScript object to populate the 
Data Cascade. If your templates aren’t compiling JavaScript objects, you probably won’t need 
this.

Notably, this is separate from (in addition to) front matter parsing (which requires `read: true`).

```js
// some configuration truncated …
    // \`false\` is the default
    getData: false, // no additional data is used
```
```js
// some configuration truncated …
  getData: async function(inputPath) {
    // DIY, this object will be merged into data cascade
    return {};
  },
```
```js
// some configuration truncated …
  // get the \`data\` property from the instance.
  getData: ["data"],
  // * \`getData: true\` is aliased to ["data"]
  // * You can use more than one property name! ["data", "otherPropName"]

  getInstanceFromInputPath: function(inputPath) {
    // Return the JavaScript object from which the \`data\` property will be retrieved.
    let instance = doSomethingMyselfToFetchAJavaScriptObject(inputPath);
    return instance;
  }
```
***Advanced Use Case:*** overriding `getData` keys for one instance

If the JavaScript object returned from `getInstanceFromInputPath` has an `eleventyDataKey` 
property, this is used to override the keys returned from the `getData` Array for this specific 
instance only. Anything you can pass into a [`new Set()` 
constructor](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Set/Se
t) works here (Array, Map, another Set).

```js
// some configuration truncated …
  // if getData is \`false\`, \`eleventyDataKey\` will not be used.
  getData: true,

  getInstanceFromInputPath: function(inputPath) {
    return {
      // Overrides \`getData\` for this instance
      eleventyDataKey: ["myOverrideData"],

      // Will not be used
      data: {
        notAvailableOnGlobalData: 456
      },

      // Will be used.
      myOverrideData: {
        availableOnGlobalData: 123
      }
    }
  },
```

In the above example, the data cascade will include a top-level variable `availableOnGlobalData` 
with a value of `123`. Using `eleventyDataKey` overrides any keys set in `getData`, which means 
(for this instance) `data` will be ignored and `notAvailableOnGlobalData` will not be present.

### compileOptions

#### compileOptions.permalink to Override Permalink Compilation

*Optional*. This has the same signature as the `compile` function and expects a reusable `render` 
function to be returned.

```js
// some configuration truncated …
  compileOptions: {
    permalink: function(contents, inputPath) {
      return (data) => {
        // Return a string to override: you’ll want to use \`data.page\`
        // Or \`return;\` (return undefined) to fallback to default behavior
      }
    }
  },
```
- Don’t compile permalink strings in the parent template language
	- `permalink: "raw"` (new default in v3.0, related [GitHub 
#2780](https://github.com/11ty/eleventy/issues/2780))
- Don’t write *any* files to the file system:
	- `permalink: false`
		- `permalink: (contents, inputPath) => false`
		- `permalink: (contents, inputPath) => ((data) => false)`
- Override the default permalink function (return a string to override)
	- `permalink: (contents, inputPath) => "…"`
		- `permalink: (contents, inputPath) => ((data) => "…")` (use the data cascade)
		- If you return nothing (or `undefined`), this will revert to the default permalink 
behavior.

This provides another way to implement Sass’ underscore convention to skip writing the file to 
the output directory:

**Filename** eleventy.config.js

```js
// … some configuration truncated
  compileOptions: {
    permalink: function(contents, inputPath) {
      let parsed = path.parse(inputPath);
      if(parsed.name.startsWith("_")) {
        return false;
      }
    }
  },
```

#### compileOptions.spiderJavaScriptDependencies

- *Optional*: Defaults to `false`

Enable to use Eleventy to spider and watch files `require` ’d in these templates. This allows you 
to control the [Watch JavaScript 
Dependencies](https://www.11ty.dev/docs/watch-serve/#watch-javascript-dependencies) feature on a 
per-template language basis. Most template languages will want the default here and keep this 
feature disabled.

#### compileOptions.cache for advanced control of caching

- *Optional*: Defaults to the value of `read`

This controls caching for the compilation step and saves the compiled template function for reuse. 
For more efficient cleanup (and long term memory use), these caches are now segmented by 
`inputPath` (Added in v2.0.0).

By default, whether or not this `cache` is enabled is tied to boolean value of `read`. If `read: 
true`, then `cache` will also be `true`. It’s unlikely you will need this, but you can override 
this to mismatch `read`.

You can also granularly control the caching key using a `getCacheKey` callback. It might be useful 
to change this when using `read: false` and `contents` are unavailable.

If you’re using v2.0.0 or newer, you shouldn’t need a `getCacheKey` callback. It is preferred 
to use the [`addDependencies` method in the `compile` callback](#registering-dependencies) instead!

**Expand to see the default `getCacheKey` implementation** (you can override this!)
```js
// some configuration truncated …
  read: false,
  compileOptions: {
    cache: true,
    getCacheKey: function(contents, inputPath) {
      // return contents; // this is the default in 1.0

      // return inputPath + contents; // this is the new default in v2.0.0

      return inputPath; // override to cache by inputPath (this means the compile function will not 
get called when the file contents change)

      // Conditionally opt-out of cache with \`return false\`
      // if(someArbitraryCondition) {
      //   return false;
      // }
    }
  },
```

### isIncrementalMatch

If you’re using v2.0.0 or newer, you shouldn’t need an `isIncrementalMatch` callback. It is 
preferred to use the [`addDependencies` method in the `compile` 
callback](#registering-dependencies) instead!

- *Optional*

A callback used for advanced control of template dependency matching. This determines if a modified 
file (from a watch/serve rebuild) is relevant to each known full template file. If the callback 
returns true, the template will be rendered. If the callback returns false, the template will be 
skipped.

**Expand to see the default \`isIncrementalMatch\` implementation** (you can override this!)
```js
// some configuration truncated …
  // Called once for each template (matching this custom template’s file extension) in your 
project.
  isIncrementalMatch: function(modifiedFile) {
    // is modifiedFile relevant to this.inputPath?
    if (this.isFileRelevantToInputPath) {
      // True if they are the same file
      // Or if they are related by any \`addDependencies\` relationships
      return true;
    }

    // If \`modifiedFile\` is not a full template (maybe an include or layout)
    // and we have no record of any dependencies for this file, we re-render everything
    if (!this.doesFileHaveDependencies && !this.isFullTemplate) {
      return true;
    }

    // Skip it
    return false;
  },
```

You can see more advanced override implementations in 
[`@11ty/eleventy-plugin-webc`](https://github.com/11ty/eleventy-plugin-webc/blob/a33dc641dfc7845d179
f7bc60f9ab2d9a9177773/src/eleventyWebcTemplate.js) ~~and 
[`@11ty/eleventy-plugin-vue`](https://github.com/11ty/eleventy-plugin-vue/blob/f705297dea3442b918b06
59b5770d7eb069bb886/.eleventy.js)~~.

---

Template Languages:
- [HTML `*.html`](https://www.11ty.dev/docs/languages/html/)
- [Markdown `*.md`](https://www.11ty.dev/docs/languages/markdown/)
- [WebC `*.webc`](https://www.11ty.dev/docs/languages/webc/)
- [JavaScript `*.11ty.js`](https://www.11ty.dev/docs/languages/javascript/)
- [Liquid `*.liquid`](https://www.11ty.dev/docs/languages/liquid/)
- [Nunjucks `*.njk`](https://www.11ty.dev/docs/languages/nunjucks/)
- [Handlebars `*.hbs`](https://www.11ty.dev/docs/languages/handlebars/)
- [Mustache `*.mustache`](https://www.11ty.dev/docs/languages/mustache/)
- [EJS `*.ejs`](https://www.11ty.dev/docs/languages/ejs/)
- [Haml `*.haml`](https://www.11ty.dev/docs/languages/haml/)
- [Pug `*.pug`](https://www.11ty.dev/docs/languages/pug/)
- [TypeScript `*.ts`](https://www.11ty.dev/docs/languages/typescript/)
- [JSX `*.jsx`](https://www.11ty.dev/docs/languages/jsx/)
- [MDX `*.mdx`](https://www.11ty.dev/docs/languages/mdx/)
- [Sass `*.scss`](https://www.11ty.dev/docs/languages/sass/)
- [Custom `*.*`](https://www.11ty.dev/docs/languages/custom/)
