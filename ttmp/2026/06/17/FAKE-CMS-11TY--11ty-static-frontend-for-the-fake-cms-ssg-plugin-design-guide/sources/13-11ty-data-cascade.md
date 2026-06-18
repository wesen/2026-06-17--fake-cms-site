---
title: "Data Cascade — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/data-cascade/
summary: "Downloaded with defuddle. Precedence order of data sources in the cascade. Re-run: defuddle parse https://www.11ty.dev/docs/data-cascade/ --md | fold -w 100 -s."
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

In Eleventy, data is merged from multiple different sources before the template is rendered. The 
data is merged in what Eleventy calls the Data Cascade.

## Sources of Data

When the data is merged in the Eleventy Data Cascade, the order of priority for sources of data is 
(from highest priority to lowest):

1. [Computed Data](https://www.11ty.dev/docs/data-computed/)
2. [Front Matter Data in a Template](https://www.11ty.dev/docs/data-frontmatter/)
3. [Template Data Files](https://www.11ty.dev/docs/data-template-dir/)
4. [Directory Data Files (and ascending Parent 
Directories)](https://www.11ty.dev/docs/data-template-dir/)
5. [Front Matter Data in Layouts](https://www.11ty.dev/docs/layouts/#front-matter-data-in-layouts) 
*(this [moved in 1.0](https://github.com/11ty/eleventy/issues/915))*
6. [Configuration API Global Data](https://www.11ty.dev/docs/data-global-custom/)
7. [Global Data Files](https://www.11ty.dev/docs/data-global/)

## Example

**Filename** my-template.md

```yaml
---
title: This is a Good Blog Post
tags:
  - CSS
  - HTML
layout: my-layout.njk
---
```

**Filename** \_includes/my-layout.njk

```yaml
---
title: This is a Very Good Blog Post
author: Zach
tags:
  - JavaScript
---
```

Note that when `my-template.md` and `my-layout.njk` share data with the same object key (`title` 
and `tags`), the “leaf template” `my-template.md` takes precedence.

The data cascade results in the following data when `my-template.md` is rendered:

**Syntax** JavaScript

```json
{
    "title": "This is a Good Blog Post",
    "author": "Zach",
    "tags": ["CSS", "HTML", "JavaScript"],
    "layout": "my-layout.njk"
}
```

Eleventy does a deep merge to combine Object literals and Arrays. (*Wait, where did the option to 
[opt-out of deep-merging](https://www.11ty.dev/docs/data-deep-merge/) go?)* You can override this 
on a per-property basis with the `override:` prefix.

## Using the override: prefix

Use the `override:` prefix on any data key to opt-out of deep-merge behavior for specific values or 
nested values.

**Filename** posts/posts.json

```json
{
    "tags": ["posts"]
}
```

**Filename** posts/firstpost.md

```markdown
---
# Instead of merging the array, this creates an empty set
override:tags: []
---
```

Even though normally the `posts/firstpost.md` file would inherit the `posts` tag from the directory 
data file (per normal [data cascade rules](https://www.11ty.dev/docs/data/)), we can override the 
`tags` value to be an empty array to opt-out of Array merge behavior.

## From the Community

[![https-benmyers-dev-blog-eleventy-data-cascade](https://screenshot.11ty.app/https%3A%2F%2Fbenmyers
.dev%2Fblog%2Feleventy-data-cascade%2F/small/1:1/_wait:2/)](https://benmyers.dev/blog/eleventy-data-
cascade/ "I Finally Understand Eleventy’s Data Cascade")

\+ [Add yours!](https://github.com/11ty/11ty-website/tree/main/src/_data/community/)

×52 resources via **[11tybundle.dev](https://11tybundle.dev/)** curated by [Bob 
Monsour](https://www.bobmonsour.com/).

- [Adding a transcript word count to your podcast About 
page](https://eleventy-plugin-podcaster.com/2026/06/adding-a-word-count-to-your-about-page/) 
 —  *Nathan Bottomley (2026)*
- [Making a Personal Reading 
List](https://claytonerrington.com/blog/making-a-personal-reading-list/?utm_source=rss)  —  
*Clayton Errington (2026)*
- [Create pages from JSON files with 
Eleventy](https://saneef.com/blog/create-pages-from-json-files-with-eleventy/)  —  *Saneef H. 
Ansari (2026)*
- [an 11ty tip-slash-hack](https://genehack.blog/2026/02/an-11ty-tip-slash-hack/)  —  *John 
Anderson (2026)*
- [exhibitionism](https://tlohde.com/blog/2025/09/exhibitionism/)  —  *tlohde (2025)*
***Expand to see 47 more resources.***
- [Last modified date of Eleventy 
posts](https://joshtronic.com/2025/07/20/last-modified-date-eleventy/)  —  *Josh Sherman 
(2025)*
- [Data vs. Collections in Eleventy](https://madbaker.com/blog/data-vs-collections-in-eleventy/) 
 —  *Mark Dyck (2025)*
- [Rendering FreeMarker and Python page templates with 
11ty](https://tfedder.de/blog/freemarker-and-python-page-templates-in-11ty/)  —  *Tobias 
Fedder (2025)*
- [Generating content driven CSS with 
Eleventy](https://httpster.io/article/generating-content-driven-css-with-eleventy/)  —  *Sami 
(2025)*
- [Building a personal digital music library with Eleventy and 
APIs](https://damianwalsh.co.uk/posts/creating-connections-with-music-and-technology/)  —  
*Damian Walsh (2025)*
- [Adding Random Content Using 
Eleventy](https://www.alwaystwisted.com/articles/random-content-using-eleventy)  —  *Stu 
Robson (2024)*
- [Here's how this is all put 
together](https://www.coryd.dev/posts/2024/heres-how-this-is-all-put-together)  —  *Cory 
Dransfeldt (2024)*
- [Adding Cooklang Support to Eleventy Three 
Ways](https://rknight.me/blog/adding-cooklang-support-to-eleventy-two-ways/)  —  *Robb Knight 
(2024)*
- [Building a Blog with 
Eleventy](https://blog.sebin-nyshkim.net/posts/building-a-blog-with-eleventy-blind-any/)  —  
*Sebin Nyshkim (2024)*
- [New Reading page, powered by the Airtable 
API](https://melanie-richards.com/blog/new-reading-page/)  —  *Melanie Richards (2024)*
- [Eleventy - Add CSV data file support](https://www.roboleary.net/blog/eleventy-csv)  —  
*Rob O'Leary (2024)*
- [Umami API Top 10 Pages](https://claytonerrington.com/blog/umami-api-top-10-pages/)  —  
*Clayton Errington (2024)*
- [Building pages from data in 
Eleventy](https://www.coryd.dev/posts/2024/building-pages-from-data-in-eleventy/)  —  *Cory 
Dransfeldt (2024)*
- [Eleventy Navigation Set URL to First Item in 
Collection](https://johnwargo.com/posts/2024/eleventy-navigation-set-url/)  —  *John M. Wargo 
(2024)*
- [Version 3](https://wonderfulfrog.com/posts/version-3/)  —  *Devin Haska (2024)*
- [Using B2 as a JSON data store](https://www.coryd.dev/posts/2024/using-b2-as-a-json-data-store/) 
 —  *Cory Dransfeldt (2024)*
- [Eleventy - Differentiate between dev and production builds with unique 
favicons](https://www.roboleary.net/2024/02/15/eleventy-favicon-modes.html)  —  *Rob O'Leary 
(2024)*
- [Community websites with Eleventy](https://hamatti.org/posts/community-websites-with-eleventy/) 
 —  *Juha-Matti Santala (2024)*
- [Creating an OPML File for my 
Blogroll](https://rknight.me/blog/creating-an-opml-file-for-my-blogroll/)  —  *Robb Knight 
(2024)*
- [Learning New Tricks](https://antonio.is/2024/02/21/learning-new-tricks/)  —  *Antonio 
Rodrigues (2024)*
- [Right here, right now](https://www.martingunnarsson.com/posts/right-here-right-now/)  —  
*Martin Gunnarsson (2024)*
- [Adding a git based changelog in 11ty](https://jamesdoc.com/blog/2023/git-changelog-in-11ty/) 
 —  *James Doc (2023)*
- [Programmatically importing your Last.fm listening data to 
ListenBrainz](https://www.coryd.dev/posts/2023/programmatically-importing-your-lastfm-listening-data
-to-listenbrainz/)  —  *Cory Dransfeldt (2023)*
- [Popular Pages with Eleventy and Fathom 
Analytics](https://rknight.me/popular-pages-with-eleventy-and-fathom-analytics/)  —  *Robb 
Knight (2023)*
- [Scratch that...use the Google Sheets 
API](https://bobmonsour.com/blog/scratch-that-use-google-sheets-api/)  —  *Bob Monsour (2023)*
- [Build a Blogroll with Eleventy](https://benmyers.dev/blog/eleventy-blogroll/)  —  *Ben 
Myers (2023)*
- [Using Eleventy filters in Directory Computed 
Data](https://illtron.net/2023/01/11ty-directory-data-filters/)  —  *Chris Coleman (2023)*
- [Relational data in Eleventy](https://danburzo.ro/eleventy-relational-data/)  —  *Dan 
Cătălin Burzo (2022)*
- [How to access post content in template listing 
files](https://www.alpower.com/tutorials/access-eleventy-post-content-in-listing/)  —  *Al 
Power (2022)*
- [Testing the Netlify Cache Plugin with 
Eleventy](https://www.raymondcamden.com/2022/06/26/testing-the-netlify-cache-plugin-with-eleventy) 
 —  *Raymond Camden (2022)*
- [Journey to the Center of the Eleventy Data Cascade with Ben 
Myers](https://www.youtube.com/watch?v=Nv5i_nuZOqw)  —  *Ben Myers (2022)*
- [11ty tips I wish I knew from the 
start](https://davidea.st/articles/11ty-tips-i-wish-i-knew-from-the-start/)  —  *David East 
(2022)*
- [Media Diary with Eleventy](https://anhvn.com/posts/2021/2021-09-21-media-diary-eleventy/) 
 —  *anh (2021)*
- [A Running Log - Garmin Connect and 
Eleventy](https://joshcrain.io/notes/2021/eleventy-garmin-connect/)  —  *Josh Crain (2021)*
- [Building a notification thingamajig with Eleventy 
data](https://multiline.co/mment/2021/07/building-notification-thingamajig-eleventy-data/) 
 —  *Ashur Cabrera (2021)*
- [Stop forgetting the Eleventy Data 
Cascade](https://blog.kaan.fyi/stop-forgetting-the-eleventy-data-cascade)  —  *Kaan Divringi 
(2021)*
- [Creating pages from data with Eleventy](https://florian.ec/blog/eleventy-data-pages/)  —  
*Florian Eckerstorfer (2021)*
- [I Finally Understand Eleventy's Data Cascade](https://benmyers.dev/blog/eleventy-data-cascade/) 
 —  *Ben Myers (2021)*
- [Site Metadata](https://11ty.rocks/tips/site-metadata/)  —  *Stephanie Eckles (2021)*
- [Using Template Content as Data for 
11ty](https://11ty.rocks/posts/using-template-content-as-data/)  —  *Stephanie Eckles (2021)*
- [How to use CSV data with Eleventy](https://www.maxkohler.com/posts/eleventy-csv/)  —  *Max 
Kohler (2021)*
- [Add structured data annotations to Eleventy 
blog](https://www.maxivanov.io/add-structured-data-to-eleventy-blog/)  —  *Max Ivanov (2020)*
- [Understanding Filters, Shortcodes and Data in 
11ty](https://www.madebymike.com.au/writing/11ty-filters-data-shortcodes/)  —  *Mike (2020)*
- [Composing data in Eleventy](https://www.falldowngoboone.com/blog/composing-data-in-eleventy/) 
 —  *Ryan Boone (2020)*
- [Using Eleventy data files for frequently used 
links](https://danabyerly.com/notes/using-eleventy-data-files-for-frequently-used-links/)  —  
*Dana Byerly (2020)*
- [Architecting data in Eleventy](https://sia.codes/posts/architecting-data-in-eleventy/)  —  
*Sia Karamalegos (2020)*
- [Five Critical Things To Do Before Working With 
11ty](https://khalidabuhakmeh.com/five-critical-things-before-working-with-11ty)  —  *Khalid 
Abuhakmeh (2020)*

---

### Other pages in Using Data

- [Eleventy Supplied Data](https://www.11ty.dev/docs/data-eleventy-supplied/)
- [Data Cascade](https://www.11ty.dev/docs/data-cascade/)
	- [Front Matter Data](https://www.11ty.dev/docs/data-frontmatter/)
		- [Custom Front Matter](https://www.11ty.dev/docs/data-frontmatter-customize/)
		- [Template & Directory Data Files](https://www.11ty.dev/docs/data-template-dir/)
		- [Global Data Files](https://www.11ty.dev/docs/data-global/)
		- [Config Global Data](https://www.11ty.dev/docs/data-global-custom/)
		- [Computed Data](https://www.11ty.dev/docs/data-computed/)
- [JavaScript Data Files](https://www.11ty.dev/docs/data-js/)
- [Custom Data File Formats](https://www.11ty.dev/docs/data-custom/)
- [Validate Data](https://www.11ty.dev/docs/data-validate/)
