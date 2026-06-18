---
title: "Create or use Plugins — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/create-plugin/
summary: "Downloaded with defuddle. How to author a plugin: function signature, options, async addPlugin. Re-run: defuddle parse https://www.11ty.dev/docs/create-plugin/ --md | fold -w 100 -s."
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

## Plugins are Configuration

At their simplest, Eleventy plugins are a function passed to the `addPlugin` method. If you’re 
familiar with [Eleventy configuration files](https://www.11ty.dev/docs/config/), this will look and 
feel very similar!

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.addPlugin(function(eleventyConfig) {
        // I am a plugin!
    });
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.addPlugin(function(eleventyConfig) {
        // I am a plugin!
    });
};
```

The plugin can be defined elsewhere in the same file:

eleventy.config.js

```js
function myPlugin(eleventyConfig) {
    // I am a plugin!
}

export default function (eleventyConfig) {
    eleventyConfig.addPlugin(myPlugin);
};
```

```js
function myPlugin(eleventyConfig) {
    // I am a plugin!
}

module.exports = function (eleventyConfig) {
    eleventyConfig.addPlugin(myPlugin);
};
```

Or in a different file:

eleventy.config.js

```js
import myPlugin from "./_config/plugin.js";

export default function (eleventyConfig) {
    eleventyConfig.addPlugin(myPlugin);
};
```

```js
const myPlugin = require("./_config/plugin.js");

module.exports = function (eleventyConfig) {
    eleventyConfig.addPlugin(myPlugin);
};
```

Or in an `npm` package:

eleventy.config.js

```js
import pluginRss from "@11ty/eleventy-plugin-rss";

export default function (eleventyConfig) {
    eleventyConfig.addPlugin(pluginRss);
};
```

```js
const pluginRss = require("@11ty/eleventy-plugin-rss");

module.exports = function (eleventyConfig) {
    eleventyConfig.addPlugin(pluginRss);
};
```

Plugins are async-friendly Added in v3.0.0 but you must `await` the `addPlugin` method.

eleventy.config.js

```js
export default async function (eleventyConfig) {
    await eleventyConfig.addPlugin(async function(eleventyConfig) {
        // I am an asynchronous plugin!
    });
};
```

```js
module.exports = async function (eleventyConfig) {
    await eleventyConfig.addPlugin(async function(eleventyConfig) {
        // I am an asynchronous plugin!
    });
};
```

## Adding a Plugin

### Installation

We use the [`npm` command line tool](https://www.npmjs.com/) (included with Node.js) to install 
plugins.

Looking for a plugin? Check out the [official plugins](https://www.11ty.dev/docs/plugins/official/) 
or [community-contributed plugins](https://www.11ty.dev/docs/plugins/community/).

```bash
npm install @11ty/eleventy-plugin-rss --save
```

### Add the plugin to Eleventy in your config file

Your config file is probably named `eleventy.config.js` or `.eleventy.js`.

eleventy.config.js

```js
import pluginRss from "@11ty/eleventy-plugin-rss";

export default function (eleventyConfig) {
    eleventyConfig.addPlugin(pluginRss);
};
```

```js
const pluginRss = require("@11ty/eleventy-plugin-rss");

module.exports = function (eleventyConfig) {
    eleventyConfig.addPlugin(pluginRss);
};
```

### Plugin Configuration Options

Use an optional second argument to `addPlugin` to customize your plugin’s behavior. These options 
are specific to the plugin. Please consult the plugin’s documentation (e.g. the 
[`eleventy-plugin-syntaxhighlight` 
README](https://github.com/11ty/eleventy-plugin-syntaxhighlight/blob/master/README.md)) to learn 
what options are available to you.

eleventy.config.js

```js
import pluginSyntaxHighlight from "@11ty/eleventy-plugin-syntaxhighlight";

export default function (eleventyConfig) {
    eleventyConfig.addPlugin(pluginSyntaxHighlight, {
        // only install the markdown highlighter
        templateFormats: ["md"],

        init: function ({ Prism }) {
            // Add your own custom language to Prism!
        },
    });
};
```

```js
const pluginSyntaxHighlight = require("@11ty/eleventy-plugin-syntaxhighlight");

module.exports = function (eleventyConfig) {
    eleventyConfig.addPlugin(pluginSyntaxHighlight, {
        // only install the markdown highlighter
        templateFormats: ["md"],

        init: function ({ Prism }) {
            // Add your own custom language to Prism!
        },
    });
};
```

**Advanced Usage: Namespacing a plugin**

It’s unlikely you’ll need this feature *but* you can namespace parts of your configuration 
using `eleventyConfig.namespace`. This will add a string prefix to all filters, tags, helpers, 
shortcodes, collections, and transforms.

eleventy.config.js

```js
import pluginRss from "@11ty/eleventy-plugin-rss";

export default function (eleventyConfig) {
    eleventyConfig.namespace("myPrefix_", () => {
        // the rssLastUpdatedDate filter is now myPrefix_rssLastUpdatedDate
        eleventyConfig.addPlugin(pluginRss);
    });
};
```

```js
const pluginRss = require("@11ty/eleventy-plugin-rss");

module.exports = function (eleventyConfig) {
    eleventyConfig.namespace("myPrefix_", () => {
        // the rssLastUpdatedDate filter is now myPrefix_rssLastUpdatedDate
        eleventyConfig.addPlugin(pluginRss);
    });
};
```

WARNING

Plugin namespacing is an application feature and should not be used if you are creating your own 
plugin (in your plugin configuration code). Follow along at [Issue 
#256](https://github.com/11ty/eleventy/issues/256).

### Advanced: Execute a plugin immediately Added in v3.0.0

Plugins (by default) execute in a second stage of configuration after the user’s configuration 
file has finished, in order to have access to [the return object in the configuration 
callback](https://www.11ty.dev/docs/config/).

You are unlikely to need this, but you can execute a plugin’s code immediately using the 
`immediate` option.

eleventy.config.js

```js
export default function (eleventyConfig, pluginOptions) {
    console.log( "first" );

    eleventyConfig.addPlugin(eleventyConfig => {
        console.log("fourth");
    });

    eleventyConfig.addPlugin(eleventyConfig => {
        console.log("second");
    }, {
        immediate: true
    });

    console.log("third");
};
```

```js
module.exports = function (eleventyConfig, pluginOptions) {
    console.log( "first" );

    eleventyConfig.addPlugin(eleventyConfig => {
        console.log("fourth");
    });

    eleventyConfig.addPlugin(eleventyConfig => {
        console.log("second");
    }, {
        immediate: true
    });

    console.log("third");
};
```

## Creating a Plugin

A plugin primarily provides a “configuration function.” This function is called when Eleventy 
is first initialized, and operates similarly to a user’s configuration function (the same 
`eleventyConfig` argument passed to the user’s `eleventy.config.js` file is passed here), in 
addition to any config passed by the user:

plugin.js

```js
export default function (eleventyConfig, pluginOptions) {
    // Your plugin code goes here
};
```

```js
module.exports = function (eleventyConfig, pluginOptions) {
    // Your plugin code goes here
};
```

Note that plugins run as a second stage after the user’s primary configuration file has executed 
(to have access to the return object values).

**Advanced Usage: Custom Plugin Arguments**

If you want to allow developers to use custom arguments provided by your plugin, you can export an 
object. Prefer using the above syntax unless you need this behavior. For an example of how this is 
used, see the [syntax highlighting 
plugin](https://github.com/11ty/eleventy-plugin-syntaxhighlight/blob/23761d7fd54de031204052017595932
7b1a0ab9b/.eleventy.js#L10)

plugin-with-args.js

```js
export default {
    initArguments: {},
    configFunction: function (eleventyConfig, pluginOptions) {
        // Your plugin code goes here
    },
};
```

```js
module.exports = {
    initArguments: {},
    configFunction: function (eleventyConfig, pluginOptions) {
        // Your plugin code goes here
    },
};
```

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.addPlugin(require("./fancy-plugin.js"), {
        init: function (initArguments) {
            // \`this\` is the eleventyConfig object
            // initArguments will be the \`myInitArguments\` object from above
        },
    });
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.addPlugin(require("./fancy-plugin.js"), {
        init: function (initArguments) {
            // \`this\` is the eleventyConfig object
            // initArguments will be the \`myInitArguments\` object from above
        },
    });
};
```

### Feature Testing

If your plugin requires a specific feature in Eleventy, you should feature test it!

plugin.js

```js
export default function (eleventyConfig, pluginOptions) {
    if(!("addTemplate" in eleventyConfig)) {
        console.log( \`[my-test-plugin] WARN Eleventy plugin compatibility: Virtual Templates are 
required for this plugin, please use Eleventy v3.0 or newer.\` );
    }
};
```

```js
module.exports = function (eleventyConfig, pluginOptions) {
    if(!("addTemplate" in eleventyConfig)) {
        console.log( \`[my-test-plugin] WARN Eleventy plugin compatibility: Virtual Templates are 
required for this plugin, please use Eleventy v3.0 or newer.\` );
    }
};
```

### Version Checking

If feature testing is not available for your specific use case, you can add this code to your 
plugin configuration to show a warning if the plugin consumer does not have a compatible version of 
Eleventy:

plugin.js

```js
export default function (eleventyConfig, pluginOptions) {
    try {
        // Emit a warning message if the application is not using Eleventy 3.0 or newer (including 
prereleases).
        eleventyConfig.versionCheck(">=3.0");
    } catch(e) {
        console.log( \`[my-test-plugin] WARN Eleventy plugin compatibility: ${e.message}\` );
    }
};
```

```js
module.exports = function (eleventyConfig, pluginOptions) {
    try {
        // Emit a warning message if the application is not using Eleventy 3.0 or newer (including 
prereleases).
        eleventyConfig.versionCheck(">=3.0");
    } catch(e) {
        console.log( \`[my-test-plugin] WARN Eleventy plugin compatibility: ${e.message}\` );
    }
};
```

- This uses the [`semver` package](https://www.npmjs.com/package/semver) and is compatible with 
advanced range syntax.
- **Upper bounding your version number is *not recommended***. Eleventy works very hard to maintain 
backwards compatibility between major versions. Please ensure your plugin code does the same!
- The `versionCheck` method has been available in Eleventy core since v0.3.2 (~2018).

### Distributing a Plugin

If you’re distributing your plugin as a package, consider following these conventions. These are 
not hard requirements.

- Add `"eleventy-plugin"` to your package.json’s `keywords` field.
- Prefix your package name with `eleventy-plugin-`
- Open a PR to add your plugin to our [list of community 
plugins](https://github.com/11ty/11ty-website/tree/main/src/_data/plugins) for publication on [our 
community plugins directory](https://www.11ty.dev/docs/plugins/community/).

---

### Other pages in Plugins

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
