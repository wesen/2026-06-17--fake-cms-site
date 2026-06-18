---
title: "Config Global Data — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/data-global-custom/
summary: "Downloaded with defuddle. Dedicated addGlobalData documentation used by the corrected design. Re-run: defuddle parse https://www.11ty.dev/docs/data-global-custom/ --md | fold -w 100 -s."
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

## Global Data from the Configuration API Added in v1.0.0

In addition to [Global Data Files](https://www.11ty.dev/docs/data-global/) global data can be added 
to the Eleventy config object using the `addGlobalData` method. This is especially useful for 
plugins.

The first value of `addGlobalData` is the key that will be available to your templates and the 
second value is the value of the value returned to the template.

## Literals

eleventy.config.js

```js
export default function (eleventyConfig) {
    // Values can be static:
    eleventyConfig.addGlobalData("myString", "myValue");
};
```

```js
module.exports = function (eleventyConfig) {
    // Values can be static:
    eleventyConfig.addGlobalData("myString", "myValue");
};
```

## More Complex Paths

The first argument can be any [lodash-set compatible path](https://lodash.com/docs/4.17.15#set):

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.addGlobalData("myNestedObject.myString", "myValue");
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.addGlobalData("myNestedObject.myString", "myValue");
};
```

## Functions

Importantly, passing a `function` to `addGlobalData` will evaluate that function before setting the 
value to the data cascade (and is async-friendly).

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.addGlobalData("myDate", () => new Date());

    // myDate’s value will be a Date instance
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.addGlobalData("myDate", () => new Date());

    // myDate’s value will be a Date instance
};
```

If you want a `function` returned, make sure you nest it:

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.addGlobalData("myFunction", () => {
        return () => new Date();
    });

    // myFunction’s value will be a function that returns a Date instance
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.addGlobalData("myFunction", () => {
        return () => new Date();
    });

    // myFunction’s value will be a function that returns a Date instance
};
```

The above is important to know when using this API with [Computed 
Data](https://www.11ty.dev/docs/data-computed/#using-javascript):

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.addGlobalData("eleventyComputed.myString", () => {
        return (data) => "This is a string!";
    });

    // myString’s value will be "This is a string!"
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.addGlobalData("eleventyComputed.myString", () => {
        return (data) => "This is a string!";
    });

    // myString’s value will be "This is a string!"
};
```

### Async/Promises

eleventy.config.js

```js
export default function (eleventyConfig) {
    // or a promise:
    eleventyConfig.addGlobalData("myFunctionPromise", () => {
        return new Promise((resolve) => {
            setTimeout(resolve, 100, "foo");
        });
    });

    // or async:
    eleventyConfig.addGlobalData("myAsyncFunction", async () => {
        return Promise.resolve("hi");
    });
};
```

```js
module.exports = function (eleventyConfig) {
    // or a promise:
    eleventyConfig.addGlobalData("myFunctionPromise", () => {
        return new Promise((resolve) => {
            setTimeout(resolve, 100, "foo");
        });
    });

    // or async:
    eleventyConfig.addGlobalData("myAsyncFunction", async () => {
        return Promise.resolve("hi");
    });
};
```

## Sources of Data

When the data is merged in the [Eleventy Data Cascade](https://www.11ty.dev/docs/data-cascade/), 
the order of priority for sources of data is (from highest priority to lowest):

1. [Computed Data](https://www.11ty.dev/docs/data-computed/)
2. [Front Matter Data in a Template](https://www.11ty.dev/docs/data-frontmatter/)
3. [Template Data Files](https://www.11ty.dev/docs/data-template-dir/)
4. [Directory Data Files (and ascending Parent 
Directories)](https://www.11ty.dev/docs/data-template-dir/)
5. [Front Matter Data in Layouts](https://www.11ty.dev/docs/layouts/#front-matter-data-in-layouts) 
*(this [moved in 1.0](https://github.com/11ty/eleventy/issues/915))*
6. [Configuration API Global Data](https://www.11ty.dev/docs/data-global-custom/) ⬅
7. [Global Data Files](https://www.11ty.dev/docs/data-global/)

---

### Other pages in Data Cascade

- [Front Matter Data](https://www.11ty.dev/docs/data-frontmatter/)
	- [Custom Front Matter](https://www.11ty.dev/docs/data-frontmatter-customize/)
- [Template & Directory Data Files](https://www.11ty.dev/docs/data-template-dir/)
- [Global Data Files](https://www.11ty.dev/docs/data-global/)
- [Config Global Data](https://www.11ty.dev/docs/data-global-custom/)
- [Computed Data](https://www.11ty.dev/docs/data-computed/)
