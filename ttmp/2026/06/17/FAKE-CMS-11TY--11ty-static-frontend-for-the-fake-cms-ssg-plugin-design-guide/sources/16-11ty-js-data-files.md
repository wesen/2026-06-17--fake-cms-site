---
title: "JavaScript Data Files — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/languages/javascript/
summary: "Downloaded with defuddle. Async functions as data; JS template renderers. Re-run: defuddle parse https://www.11ty.dev/docs/languages/javascript/ --md | fold -w 100 -s."
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

| Eleventy Short Name | File Extension | Version Added |
| --- | --- | --- |
| `11ty.js` | `.11ty.js` | `*` |
| `11ty.js` | `.11ty.cjs` | `0.11.0` |
| `11ty.js` | `.11ty.mjs` | `3.0.0` |

- Related languages: [JSX](https://www.11ty.dev/docs/languages/jsx/), 
[TypeScript](https://www.11ty.dev/docs/languages/typescript/), 
[MDX](https://www.11ty.dev/docs/languages/mdx/)
- *[Front matter](https://www.11ty.dev/docs/data-frontmatter/) is not supported in JavaScript 
files. Use a `data` export instead.*

Eleventy supports many different types of JavaScript content that will be parsed as Eleventy 
templates. They are comprehensively described below.

## Raw Values

Raw values will not have access to Data or [JavaScript Template 
Functions](#javascript-template-functions). [Use a function](#function) that returns a value 
instead.

### String

```js
export default "<p>Zach</p>";
```

```js
module.exports = "<p>Zach</p>";
```

Or [template 
literals](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Template_literals):

```js
export default \`<p>These can
span
multiple
lines!</p>\`;
```

```js
module.exports = \`<p>These can
span
multiple
lines!</p>\`;
```

### Buffer

Some templating libraries return 
[Buffers](https://nodejs.org/api/buffer.html#buffer_class_method_buffer_from_string_encoding) (e.g. 
[viperHTML](https://github.com/WebReflection/viperHTML)).

```js
export default Buffer.from("<p>Zách</p>");
```

```js
module.exports = Buffer.from("<p>Zách</p>");
```

### Promise

```js
export default new Promise((resolve, reject) => {
    setTimeout(function () {
        resolve("<p>Zach</p>");
    }, 1000);
});
```

```js
module.exports = new Promise((resolve, reject) => {
    setTimeout(function () {
        resolve("<p>Zach</p>");
    }, 1000);
});
```

## Function

Can return any [raw value](#raw-values) (e.g. String, Buffer, Promise). Use [template 
literals](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Template_literals) to 
embed data values without having to concatenate strings!

```js
export default function (data) {
    return \`<p>${data.name}</p>\`;
};
```

```js
module.exports = function (data) {
    return \`<p>${data.name}</p>\`;
};
```

De-structuring syntax is a little bit easier to read:

```js
export default function ({ name }) {
    return \`<p>${name}</p>\`;
};
```

```js
module.exports = function ({ name }) {
    return \`<p>${name}</p>\`;
};
```

Maybe you like arrow functions:

```js
export default ({ name }) => \`<p>${name}</p>\`;
```

```js
module.exports = ({ name }) => \`<p>${name}</p>\`;
```

`async` functions work too:

```js
const getAnAsyncThing = require("./lib/asyncThing");

export default async function (data) {
    return \`<p>${await getAnAsyncThing()}</p>\`;
};
```

```js
const getAnAsyncThing = require("./lib/asyncThing");

module.exports = async function (data) {
    return \`<p>${await getAnAsyncThing()}</p>\`;
};
```

## Classes

Eleventy looks for classes that have a `render` method and uses `render` to return the content of 
the template. `render` methods can be `async`.

`render` can return any [raw value](#raw-values) (e.g. String, Buffer, Promise).

```js
class Test {
    // or \`async render({name}) {\`
    render({ name }) {
        return \`<p>${name}</p>\`;
    }
}

export default Test;
```

```js
class Test {
    // or \`async render({name}) {\`
    render({ name }) {
        return \`<p>${name}</p>\`;
    }
}

module.exports = Test;
```

### Optional data Method

[Front Matter](https://www.11ty.dev/docs/data-frontmatter/) is not supported in JavaScript template 
types. Use `data` methods instead! Additionally, there are more alternative options in the [Data 
Cascade](https://www.11ty.dev/docs/data-cascade/).

This data acts as Front Matter for the template and similarly to Front Matter will take precedence 
over all other data in the data cascade. The `data` method can be asynchronous `async data()` or it 
can be a getter `get data()`.

```js
class Test {
    // or \`async data() {\`
    // or \`get data() {\`
    data() {
        return {
            name: "Ted",
            layout: "teds-rad-layout",
            // … other front matter keys
        };
    }

    render({ name }) {
        // will always be "Ted"
        return \`<p>${name}</p>\`;
    }
}

export default Test;
```

```js
class Test {
    // or \`async data() {\`
    // or \`get data() {\`
    data() {
        return {
            name: "Ted",
            layout: "teds-rad-layout",
            // … other front matter keys
        };
    }

    render({ name }) {
        // will always be "Ted"
        return \`<p>${name}</p>\`;
    }
}

module.exports = Test;
```

### Permalinks

The `permalink` data key will work here. Permalinks can be a [raw value](#raw-values) (e.g. String, 
Buffer, Promise) or a Function that returns any raw value.

#### Permalink String

```js
class Test {
    data() {
        return {
            // Writes to "/my-permalink/index.html"
            permalink: "/my-permalink/",
        };
    }

    render(data) {
        /* … */
    }
}

export default Test;
```

```js
class Test {
    data() {
        return {
            // Writes to "/my-permalink/index.html"
            permalink: "/my-permalink/",
        };
    }

    render(data) {
        /* … */
    }
}

module.exports = Test;
```

#### Permalink Function

Permalink Functions can return any [raw value](#raw-values) (e.g. String, Buffer, Promise).

```js
class Test {
    data() {
        return {
            key: "hello",
            // Writes to "/my-permalink/hello/index.html"
            permalink: (data) => \`/my-permalink/${data.key}/\`,
        };
    }

    render(data) {
        /* … */
    }
}

export default Test;
```

```js
class Test {
    data() {
        return {
            key: "hello",
            // Writes to "/my-permalink/hello/index.html"
            permalink: (data) => \`/my-permalink/${data.key}/\`,
        };
    }

    render(data) {
        /* … */
    }
}

module.exports = Test;
```

#### Permalink Function using a Filter

Universal filters, shortcodes, and other JavaScript Template Functions work here and are exposed on 
`this`. Read more about [Eleventy provided Universal 
Filters](https://www.11ty.dev/docs/filters/#eleventy-provided-filters).

```js
class Test {
    data() {
        return {
            title: "This is my blog post title",
            // Writes to "/this-is-my-blog-post-title/index.html"
            permalink: function (data) {
                return \`/${this.slug(data.title)}/\`;
            },
        };
    }

    render(data) {
        /* … */
    }
}

export default Test;
```

```js
class Test {
    data() {
        return {
            title: "This is my blog post title",
            // Writes to "/this-is-my-blog-post-title/index.html"
            permalink: function (data) {
                return \`/${this.slug(data.title)}/\`;
            },
        };
    }

    render(data) {
        /* … */
    }
}

module.exports = Test;
```

### Markdown and JavaScript

Yes, you can use JavaScript as your preprocessor language for Markdown. Read more about 
[`templateEngineOverride`](https://www.11ty.dev/docs/template-overrides/).

```js
class Test {
    data() {
        return {
            myName: "Zach",
            templateEngineOverride: "11ty.js,md",
        };
    }

    render(data) {
        return \`# This is ${data.myName}\`;
    }
}

export default Test;
```

```js
class Test {
    data() {
        return {
            myName: "Zach",
            templateEngineOverride: "11ty.js,md",
        };
    }

    render(data) {
        return \`# This is ${data.myName}\`;
    }
}

module.exports = Test;
```

While `templateEngineOverride: 11ty.js,md` works to add markdown support, the special behavior of 
JavaScript templates does not allow other template engines to be supported here (e.g. 
`templateEngineOverride: njk,md`). One workaround is to use the [Render 
Plugin](https://www.11ty.dev/docs/plugins/render/).

## JavaScript Template Functions

A JavaScript Template Function allows you to extend your JavaScript templates with extra 
functionality. If you add any Universal Filters or Shortcodes, they will be exposed as JavaScript 
Template Functions.

eleventy.config.js

```js
export default function(eleventyConfig) {
  eleventyConfig.addJavaScriptFunction("myFunction", function(a, b) { /* … */ });
};
```

```js
module.exports = function(eleventyConfig) {
  eleventyConfig.addJavaScriptFunction("myFunction", function(a, b) { /* … */ });
};
```

js-fn-example.11ty.js

```js
export default function (data) {
    return \`<h1>${this.myFunction(data.a, data.b)}</h1>\`;
};
```

```js
module.exports = function (data) {
    return \`<h1>${this.myFunction(data.a, data.b)}</h1>\`;
};
```

### Asynchronous JavaScript Template Functions

This works the same as any `async` JavaScript function or function that returns a `Promise`.

This is the same as the example above but adds `async` before the `function`.

eleventy.config.js

```js
export default function(eleventyConfig) {
  eleventyConfig.addJavaScriptFunction("myAsyncFunction", async function(a, b) { /* … */ });
};
```

```js
module.exports = function(eleventyConfig) {
  eleventyConfig.addJavaScriptFunction("myAsyncFunction", async function(a, b) { /* … */ });
};
```

This is the same as the example above but adds `await` before the function is called.

**Filename** js-async-fn-example.11ty.js

```js
export default async function (data) {
    return \`<h1>${await this.myAsyncFunction(data.a, data.b)}</h1>\`;
};
```

```js
module.exports = async function (data) {
    return \`<h1>${await this.myAsyncFunction(data.a, data.b)}</h1>\`;
};
```

### Warning about Arrow Functions

WARNING

Note that by definition ([read on 
MDN](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/Arrow_functions#No_
separate_this)) arrow functions do not have access to `this`, so any use of JavaScript Functions 
inside of an arrow function template will throw an error.

**Filename** js-arrow-fn-example.11ty.js

```js
export default (data) => {
    // Using \`this\` in an arrow function will throw an error!
    return \`<h1>${this.myFunction(data.a, data.b)}</h1>\`;
};
```

```js
module.exports = (data) => {
    // Using \`this\` in an arrow function will throw an error!
    return \`<h1>${this.myFunction(data.a, data.b)}</h1>\`;
};
```

### Relationship to Filters and Shortcodes

Any universal filters or shortcodes will also be available as JavaScript Template Functions.

eleventy.config.js

```js
export default function(eleventyConfig) {
  // Universal filters (Adds to Liquid, Nunjucks, 11ty.js)
  eleventyConfig.addFilter("myFilter", function(myVariable) { /* … */ });

  // Universal Shortcodes (Adds to Liquid, Nunjucks, 11ty.js)
  eleventyConfig.addShortcode("user", function(firstName, lastName) { /* … */ });

  // Universal Paired Shortcodes (Adds to Liquid, Nunjucks, 11ty.js)
  eleventyConfig.addPairedShortcode("pairedUser", function(content, firstName, lastName) { /* … 
*/ });
};
```

```js
module.exports = function(eleventyConfig) {
  // Universal filters (Adds to Liquid, Nunjucks, 11ty.js)
  eleventyConfig.addFilter("myFilter", function(myVariable) { /* … */ });

  // Universal Shortcodes (Adds to Liquid, Nunjucks, 11ty.js)
  eleventyConfig.addShortcode("user", function(firstName, lastName) { /* … */ });

  // Universal Paired Shortcodes (Adds to Liquid, Nunjucks, 11ty.js)
  eleventyConfig.addPairedShortcode("pairedUser", function(content, firstName, lastName) { /* … 
*/ });
};
```

**Filename** universal-examples.11ty.js

```js
export default function (data) {
    return \`
<h1>${this.myFilter(data.myVar)}</h1>
<p>${this.user(data.firstName, data.lastName)}</p>
<p>${this.pairedUser(
        \`Here is some more content\`,
        data.firstName,
        data.lastName
    )}</p>
\`;
};
```

```js
module.exports = function (data) {
    return \`
<h1>${this.myFilter(data.myVar)}</h1>
<p>${this.user(data.firstName, data.lastName)}</p>
<p>${this.pairedUser(
        \`Here is some more content\`,
        data.firstName,
        data.lastName
    )}</p>
\`;
};
```

### Access to page data values

If you aren’t using an arrow function, JavaScript Functions (and Nunjucks, Liquid Shortcodes) 
will have access to Eleventy [`page` data 
values](https://www.11ty.dev/docs/data-eleventy-supplied/#page-variable-contents) without needing 
to pass them in as arguments.

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.addJavaScriptFunction("myFunction", function () {
        // Available in 0.11.0 and above
        console.log(this.page);

        // For example:
        console.log(this.page.url);
        console.log(this.page.inputPath);
        console.log(this.page.fileSlug);
    });
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.addJavaScriptFunction("myFunction", function () {
        // Available in 0.11.0 and above
        console.log(this.page);

        // For example:
        console.log(this.page.url);
        console.log(this.page.inputPath);
        console.log(this.page.fileSlug);
    });
};
```

## From the Community

[![https-willmartian-com-posts-conditional-rendering-eleventy](https://screenshot.11ty.app/https%3A%
2F%2Fwillmartian.com%2Fposts%2Fconditional-rendering-eleventy%2F/small/1:1/_wait:2/)](https://willma
rtian.com/posts/conditional-rendering-eleventy/ "Conditionally Rendering JavaScript Templates")

\+ [Add yours!](https://github.com/11ty/11ty-website/tree/main/src/_data/community/)

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

---

### Related Docs

- [Template Filters](https://www.11ty.dev/docs/filters/)
- [Template Shortcodes](https://www.11ty.dev/docs/shortcodes/)
- [Watch JavaScript Dependencies](https://www.11ty.dev/docs/config/#watch-javascript-dependencies)
