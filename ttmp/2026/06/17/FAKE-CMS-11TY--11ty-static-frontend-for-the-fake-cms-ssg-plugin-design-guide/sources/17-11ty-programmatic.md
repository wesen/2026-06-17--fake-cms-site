---
title: "Programmatic API — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/programmatic/
summary: "Downloaded with defuddle. Running Eleventy from a Node script (.write/.toJSON). Re-run: defuddle parse https://www.11ty.dev/docs/programmatic/ --md | fold -w 100 -s."
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

## Programmatic API Added in v1.0.0

You can run Eleventy in any arbitrary Node script.

## Write to the file system

Now create a file called `my-node-script.js` with the following contents:

my-node-script.js

```js
import Eleventy from "@11ty/eleventy";

let elev = new Eleventy();
await elev.write();
```

```js
(async function () {
    const { Eleventy } = await import("@11ty/eleventy");

    let elev = new Eleventy();
    await elev.write();
})();
```

Then run your new script from the command line.

```
node my-node-script.js
```

## Don’t write to the file system

Using `.write()` will write your output to the file system. If, instead, you want to retrieve the 
content programmatically without writing, use `.toJSON()` or `.toNDJSON()`.

### JSON Output

my-node-script.js

```js
import Eleventy from "@11ty/eleventy";

let elev = new Eleventy();
let json = await elev.toJSON();

// All results
console.log(json);
```

```js
(async function () {
    const { Eleventy } = await import("@11ty/eleventy");

    let elev = new Eleventy();
    let json = await elev.toJSON();

    // All results
    console.log(json);
})();
```

#### Adding data to JSON output

You can use the `eleventyConfig.dataFilterSelectors` configuration API `Set` to add or remove 
lodash-style selectors for Data Cascade entries to be included in individual entries from the 
`toJSON` method.

my-node-script.js

```js
import Eleventy from "@11ty/eleventy";

let elev = new Eleventy(".", "_site", {
    config: function(eleventyConfig) {
        eleventyConfig.dataFilterSelectors.add("globalData.key1");
        eleventyConfig.dataFilterSelectors.add("globalData.key2");
        eleventyConfig.dataFilterSelectors.add("someProperty.key");
    }
});

let json = await elev.toJSON();

// All results with
// json[…].data.globalData.key1
// json[…].data.globalData.key2
// json[…].data.someProperty.key
console.log(json);
```

```js
(async function () {
    const { Eleventy } = await import("@11ty/eleventy");

    let elev = new Eleventy(".", "_site", {
        config: function(eleventyConfig) {
            eleventyConfig.dataFilterSelectors.add("globalData.key1");
            eleventyConfig.dataFilterSelectors.add("globalData.key2");
            eleventyConfig.dataFilterSelectors.add("someProperty.key");
        }
    });

    let json = await elev.toJSON();

    // All results with
    // json[…].data.globalData.key1
    // json[…].data.globalData.key2
    // json[…].data.someProperty.key
    console.log(json);
})();
```

### ndjson Output

my-node-script.js

```js
import Eleventy from "@11ty/eleventy";

let elev = new Eleventy();
let stream = await elev.toNDJSON();
stream.on("data", (entry) => {
    // Stream one output result at a time
    let json = JSON.parse(entry.toString());
    console.log(json);
});
```

```js
(async function () {
    const { Eleventy } = await import("@11ty/eleventy");

    let elev = new Eleventy();
    let stream = await elev.toNDJSON();
    stream.on("data", (entry) => {
        // Stream one output result at a time
        let json = JSON.parse(entry.toString());
        console.log(json);
    });
})();
```

## Changing the Input and Output Directories

The first argument is the input directory. The second argument is the output directory.

my-node-script.js

```js
import Eleventy from "@11ty/eleventy";

let elev = new Eleventy(".", "_site");

// Use \`write\` or \`toJSON\` or \`toNDJSON\`
```

```js
(async function () {
    const { Eleventy } = await import("@11ty/eleventy");

    let elev = new Eleventy(".", "_site");

    // Use \`write\` or \`toJSON\` or \`toNDJSON\`
})();
```

## Full Options List

The third argument to Eleventy is an options object.

*(This documentation section is a work in progress but [you’re welcome to dig into the `Eleventy` 
class source code in `v3.1.6` to learn 
more](https://github.com/11ty/eleventy/blob/v3.1.6/src/Eleventy.js))*

my-node-script.js

```js
import Eleventy from "@11ty/eleventy";

let elev = new Eleventy(".", "_site", {
    // --quiet
    quietMode: true,

    // --config
    configPath: ".eleventy.js",

    config: function (eleventyConfig) {
        // Do some custom Configuration API stuff
        // Works great with eleventyConfig.addGlobalData
    },
});

// Use \`write\` or \`toJSON\` or \`toNDJSON\`
```

```js
(async function () {
    const { Eleventy } = await import("@11ty/eleventy");

    let elev = new Eleventy(".", "_site", {
        // --quiet
        quietMode: true,

        // --config
        configPath: ".eleventy.js",

        config: function (eleventyConfig) {
            // Do some custom Configuration API stuff
            // Works great with eleventyConfig.addGlobalData
        },
    });

    // Use \`write\` or \`toJSON\` or \`toNDJSON\`
})();
```

---

### Other pages in Advanced

- [Release History](https://www.11ty.dev/docs/versions/)
- [Programmatic API](https://www.11ty.dev/docs/programmatic/)
- [Configuration Events](https://www.11ty.dev/docs/events/)
- [Order of Operations](https://www.11ty.dev/docs/advanced-order/)
