---
title: "Plugins — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/plugins/
summary: "Downloaded with defuddle. What a plugin is; addPlugin contract. Re-run: defuddle parse https://www.11ty.dev/docs/plugins/ --md | fold -w 100 -s."
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

Plugins are custom code that Eleventy can import into a project from an external repository.

All *official* plugins live under the `@11ty` npm organization and plugin names will include the 
`@11ty/` prefix.

- [Create or use Plugins](https://www.11ty.dev/docs/create-plugin/): Plugins are just 
configuration. Learn how to create a plugin of your own to reuse functionality or to organize your 
configuration file.
- [Image](https://www.11ty.dev/docs/plugins/image/): A utility to resize and generate images.
	- [Image Data Files](https://www.11ty.dev/docs/plugins/image-datafiles/): Use images to 
populate data in the Data Cascade.
		- [Image JavaScript API](https://www.11ty.dev/docs/plugins/image-js/): Low-level 
JavaScript API works independent of Eleventy.
		- [Image Shortcodes](https://www.11ty.dev/docs/plugins/image-shortcodes/): Use 
universal shortcodes in Nunjucks, Liquid, or 11ty.js templates.
		- [Image WebC](https://www.11ty.dev/docs/plugins/image-webc/): ~~Use a WebC 
component for WebC templates.~~
- [Fetch](https://www.11ty.dev/docs/plugins/fetch/): A utility to fetch and cache network requests.
- [`<is-land>`](https://www.11ty.dev/docs/plugins/is-land/): A plugin to smartly and efficiently 
load and initialize client-side components to your web site.
- [Render](https://www.11ty.dev/docs/plugins/render/): A plugin to add shortcodes to render an 
Eleventy template string (or file) inside of another template.
- [Internationalization (i18n)](https://www.11ty.dev/docs/plugins/i18n/): Utilities to manage pages 
and linking between localized content on Eleventy projects.
- [RSS](https://www.11ty.dev/docs/plugins/rss/): Generate an RSS or Atom feed to allow others to 
subscribe to your content using a feed reader.
- [Upgrade Helper](https://www.11ty.dev/docs/plugins/upgrade-help/): A plugin to help you upgrade 
your Eleventy project between major version releases.
- [Syntax Highlighting](https://www.11ty.dev/docs/plugins/syntaxhighlight/): Code syntax 
highlighting using PrismJS without client-side JavaScript.
- [InputPath to URL](https://www.11ty.dev/docs/plugins/inputpath-to-url/): Maps an Eleventy input 
file path to its output URL.
- [Navigation](https://www.11ty.dev/docs/plugins/navigation/): A plugin for creating hierarchical 
navigation in Eleventy projects. Supports breadcrumbs too!
- [HTML `<base>`](https://www.11ty.dev/docs/plugins/html-base/): Emulate the `<base>` element by 
adding a prefix to all URLs in `.html` output files.
- [Bundle](https://www.11ty.dev/docs/plugins/bundle/): A plugin create small plain-text bundles of 
code (CSS, JS, HTML, SVG, etc)
- [Id Attribute](https://www.11ty.dev/docs/plugins/id-attribute/): A plugin to add \`id\` 
attributes to headings.
- [Directory Output](https://www.11ty.dev/docs/plugins/directory-output/): A plugin to group and 
sort console output by directory, with file size and benchmarks.
- [Inclusive Language](https://www.11ty.dev/docs/plugins/inclusive-language/): A plugin to check 
for inclusive language in markdown files.
- [Community Plugins](https://www.11ty.dev/docs/plugins/community/)
- [Retired Plugins](https://www.11ty.dev/docs/plugins/retired/)
	- [Edge](https://www.11ty.dev/docs/plugins/edge/): A plugin to run Eleventy in an Edge 
Function to add dynamic content to your Eleventy sites.
		- [Serverless](https://www.11ty.dev/docs/plugins/serverless/): A plugin to run 
Eleventy in a serverless function for server side rendering or to speed up builds for very large 
sites.

## From the Community![Favicon for 
v1.indieweb-avatar.11ty.dev/https%3A%2F%2Fbryanlrobinson.com%2Fblog%2Fcreating-11ty-plugin-embed-svg
-contents%2F](https://v1.indieweb-avatar.11ty.dev/https%3A%2F%2Fbryanlrobinson.com%2Fblog%2Fcreating
-11ty-plugin-embed-svg-contents%2F/)

Favicon for 
v1.indieweb-avatar.11ty.dev/https%3A%2F%2Fbryanlrobinson.com%2Fblog%2Fcreating-11ty-plugin-embed-svg
-contents%2F

[![https-bryanlrobinson-com-blog-creating-11ty-plugin-embed-svg-contents](https://screenshot.11ty.ap
p/https%3A%2F%2Fbryanlrobinson.com%2Fblog%2Fcreating-11ty-plugin-embed-svg-contents%2F/small/1:1/_wa
it:2/)](https://bryanlrobinson.com/blog/creating-11ty-plugin-embed-svg-contents/ "Creating an 11ty 
Plugin—SVG Embed Tool")![Favicon for 
v1.indieweb-avatar.11ty.dev/https%3A%2F%2Fwww.lenesaile.com%2Fen%2Fblog%2Forganizing-the-eleventy-co
nfig-file%2F](https://v1.indieweb-avatar.11ty.dev/https%3A%2F%2Fwww.lenesaile.com%2Fen%2Fblog%2Forga
nizing-the-eleventy-config-file%2F/)

Favicon for 
v1.indieweb-avatar.11ty.dev/https%3A%2F%2Fwww.lenesaile.com%2Fen%2Fblog%2Forganizing-the-eleventy-co
nfig-file%2F

[![https-www-lenesaile-com-en-blog-organizing-the-eleventy-config-file](https://screenshot.11ty.app/
https%3A%2F%2Fwww.lenesaile.com%2Fen%2Fblog%2Forganizing-the-eleventy-config-file%2F/small/1:1/_wait
:2/)](https://www.lenesaile.com/en/blog/organizing-the-eleventy-config-file/ "Organizing the 
Eleventy config file")

\+ [Add yours!](https://github.com/11ty/11ty-website/tree/main/src/_data/community/)
