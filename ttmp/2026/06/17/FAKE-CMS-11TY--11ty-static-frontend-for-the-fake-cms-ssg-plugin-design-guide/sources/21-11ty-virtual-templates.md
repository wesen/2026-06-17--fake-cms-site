---
title: "Virtual Templates — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/virtual-templates/
summary: "Downloaded with defuddle. Documents eleventyConfig.addTemplate for virtual templates, the correct API if a plugin contributes generated pages."
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

## Virtual Templates Added in v3.0.0

In addition to template files in your input directory, Eleventy can also process virtual templates 
defined in your configuration file (or plugins). Related [GitHub 
#1612](https://github.com/11ty/eleventy/issues/1612).

The [RSS plugin offers a virtual template](https://www.11ty.dev/docs/plugins/rss/#virtual-template) 
to add feeds to your project.

## API

```js
eleventyConfig.addTemplate(virtualPath, content, data = {});
```
- `virtualPath`: used to determine the template language and data cascade for this template. This 
path is relative to your project’s input directory.
- `content`: usually a string, but maybe JavaScript (if using an `11ty.js` template). Can include 
front matter if the template language supports it.
- `data`: a data object tied to the template. A little more ergonomic than front matter but 
functionally the same.

## Example

### Markdown, HTML (via Liquid) Layout

eleventy.config.js

```js
export default function(eleventyConfig) {
    // Create content templates Files
    eleventyConfig.addTemplate("virtual.md", \`# Hello\`, {
            layout: "virtual.html"
    });

    // Works great with Layouts too
    eleventyConfig.addTemplate("_includes/virtual.html", \`<!-- Layout -->{{ content }}\`);
};
```

```js
module.exports = function(eleventyConfig) {
    // Create content templates Files
    eleventyConfig.addTemplate("virtual.md", \`# Hello\`, {
            layout: "virtual.html"
    });

    // Works great with Layouts too
    eleventyConfig.addTemplate("_includes/virtual.html", \`<!-- Layout -->{{ content }}\`);
};
```

### JavaScript

Any of the JavaScript shapes on [`11ty.js` 
templates](https://www.11ty.dev/docs/languages/javascript/) are also supported here.

eleventy.config.js

```js
export default function(eleventyConfig) {
    // Create content templates Files
    eleventyConfig.addTemplate("virtual.11ty.js", function(data) {
        return \`<h1>Hello</h1>\`;
    });
};
```

```js
module.exports = function(eleventyConfig) {
    // Create content templates Files
    eleventyConfig.addTemplate("virtual.11ty.js", function(data) {
        return \`<h1>Hello</h1>\`;
    });
};
```

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
