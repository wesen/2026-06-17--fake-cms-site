---
title: "Collections — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/collections/
summary: "Downloaded with defuddle. addCollection API; building typed groups. Re-run: defuddle parse https://www.11ty.dev/docs/collections/ --md | fold -w 100 -s."
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

## Collections (Using Tags)

While [pagination](https://www.11ty.dev/docs/pagination/) allows you to iterate over a data set to 
create multiple templates, a collection allows you to group content in interesting ways. A piece of 
content can be a part of multiple collections, if you assign the same string value to the `tags` 
key in the front matter.

Take care to note that `tags` have a singular purpose in Eleventy: to construct collections of 
content. Some blogging platforms use Tags to refer to a hierarchy of labels for the content (e.g. a 
[tag cloud](https://en.wikipedia.org/wiki/Tag_cloud)).

## A Blog Example

For a blog site, your individual post files may use a tag called `post`, but it can be whatever you 
want. In this example, `mypost.md` has a single tag `post`:

**Syntax** Markdown

```markdown
---
tags: post
title: Hot Take—Social Media is Considered Harmful
---
```

This will place this `mypost.md` into the `post` collection with all other pieces of content 
sharing the `post` tag. To reference this collection and make a list of all posts, use the 
`collections` object in any template:

```
<ul>
{%- for post in collections.post -%}
  <li>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```
<ul>
{%- for post in collections.post -%}
  <li>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```js
export function render(data) {
  return \`<ul>
    ${data.collections.post
      .map((post) => \`<li>${post.data.title}</li>\`)
      .join("\n")}
  </ul>\`;
};
```

```js
exports.render = function (data) {
  return \`<ul>
    ${data.collections.post
      .map((post) => \`<li>${post.data.title}</li>\`)
      .join("\n")}
  </ul>\`;
};
```

### A note about using - in tags

If you use `-` in your collection names (e.g. `tags: "post-with-dash"`), remember that some 
template languages require square bracket notation to reference it in collections. Read more at 
[Issue #567](https://github.com/11ty/eleventy/issues/567).

```
<ul>
{%- for post in collections.post-with-dash -%}
  <li>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```
<ul>
{%- for post in collections['post-with-dash'] -%}
  <li>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```js
export function render(data) {
  return \`<ul>
    ${data.collections['post-with-dash']
      .map((post) => \`<li>${post.data.title}</li>\`)
      .join("\n")}
  </ul>\`;
};
```

```js
exports.render = function (data) {
  return \`<ul>
    ${data.collections['post-with-dash']
      .map((post) => \`<li>${post.data.title}</li>\`)
      .join("\n")}
  </ul>\`;
};
```

### Declare your collections for incremental builds

Added in v2.0.0Use the `eleventyImport` object to declare any collections you use (data cascade 
friendly) to inform the relationships for smarter incremental builds. This is an Array of 
collection names. Read more about [importing 
collections](https://github.com/11ty/eleventy/issues/975).

```
---
eleventyImport:
  collections: ["post"]
---
<ul>
{%- for post in collections.post -%}
  <li>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```
---
eleventyImport:
  collections: ["post"]
---
<ul>
{%- for post in collections.post -%}
  <li>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```js
export function data() {
  return {
    eleventyImport: {
      collections: ["post"],
    },
  };
};

export function render(data) {
  return \`<ul>
    ${data.collections.post
      .map((post) => \`<li>${post.data.title}</li>\`)
      .join("\n")}
  </ul>\`;
};
```

```js
exports.data = function () {
  return {
    eleventyImport: {
      collections: ["post"],
    },
  };
};

exports.render = function (data) {
  return \`<ul>
    ${data.collections.post
      .map((post) => \`<li>${post.data.title}</li>\`)
      .join("\n")}
  </ul>\`;
};
```

### Use an aria-current attribute on the current page

Compare the `post.url` and special Eleventy-provided `page.url` variable to find the current page. 
Building on the previous example:

```
<ul>
{%- for post in collections.post -%}
  <li{% if page.url == post.url %} aria-current="page"{% endif %}>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```
<ul>
{%- for post in collections.post -%}
  <li{% if page.url == post.url %} aria-current="page"{% endif %}>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```js
export function render(data) {
    return \`<ul>
    ${data.collections.post
            .map(
                (post) =>
                    \`<li${data.page.url === post.url ? \` aria-current="page"\` : ""}>${
                        post.data.title
                    }</li>\`
            )
            .join("\n")}
  </ul>\`;
};
```

```js
exports.render = function (data) {
    return \`<ul>
    ${data.collections.post
            .map(
                (post) =>
                    \`<li${data.page.url === post.url ? \` aria-current="page"\` : ""}>${
                        post.data.title
                    }</li>\`
            )
            .join("\n")}
    </ul>\`;
};
```

Background: `aria-current="page"` tells assistive technology, such as screen readers, which page of 
a set of pages is the current active one. It also provides a hook for your CSS styling, using its 
attribute selector: `[aria-current="page"] {}`.

## The Special all Collection

By default Eleventy puts all of your content (independent of whether or not it has any assigned 
tags) into the `collections.all` Collection. This allows you to iterate over all of your content 
inside of a template.

### Link to all Eleventy generated content

```
<ul>
{%- for post in collections.all -%}
  <li><a href="{{ post.url }}">{{ post.url }}</a></li>
{%- endfor -%}
</ul>
```

```
<ul>
{%- for post in collections.all -%}
  <li><a href="{{ post.url }}">{{ post.url }}</a></li>
{%- endfor -%}
</ul>
```

```js
export function render(data) {
  return \`<ul>
    ${data.collections.all
      .map((post) => \`<li><a href="${post.url}">${post.url}</a></li>\`)
      .join("\n")}
  </ul>\`;
};
```

```js
exports.render = function (data) {
  return \`<ul>
    ${data.collections.all
      .map((post) => \`<li><a href="${post.url}">${post.url}</a></li>\`)
      .join("\n")}
  </ul>\`;
};
```

## How to Exclude content from Collections

In front matter (or further upstream in the data cascade), set the `eleventyExcludeFromCollections` 
option to true to opt out of specific pieces of content added to all collections (including 
`collections.all`, collections set using tags, or collections added from the Configuration API in 
your config file). Useful for your RSS feed, `sitemap.xml`, custom templated `.htaccess` files, et 
cetera.

**Filename** excluded.md

```markdown
---
eleventyExcludeFromCollections: true
tags: post
---

This will not be available in \`collections.all\` or \`collections.post\`.
```

Added in v3.0.0 `eleventyExcludeFromCollections` can now also accept an array of tag names:

```markdown
---
eleventyExcludeFromCollections: ["post"]
---

This will be available in \`collections.all\` but not \`collections.post\`.
```

## Add to a Collection using Tags

You can use a single tag, as in the above example OR you can use any number of tags for the 
content, using YAML syntax for a list.

### A single tag: cat

```markdown
---
tags: cat
---
```

This content would show up in the template data inside of `collections.cat`.

### Using multiple words in a single tag

```markdown
---
tags: cat and dog
---
```

If you use multiple words for one tag you can access the content by the following syntax 
`collections['cat and dog']`.

### Multiple tags, single line

```markdown
---
tags: ["cat", "dog"]
---
```

This content would show up in the template data inside of `collections.cat` and `collections.dog`.

### Multiple tags, multiple lines

```markdown
---
tags:
  - cat
  - dog
---
```

This content would show up in the template data inside of `collections.cat` and `collections.dog`.

### Override tags

As of Eleventy 1.0, the [Data Cascade](https://www.11ty.dev/docs/data-cascade/) is combined using 
[deep data merge](https://www.11ty.dev/docs/data-deep-merge/) by default, which means tags are 
merged together with tags assigned higher in the data cascade (the Arrays are combined). To 
redefine `tags` in the front matter use [the `override:` 
prefix](https://www.11ty.dev/docs/data-deep-merge/#using-the-override-prefix):

```markdown
---
override:tags: []
---
```

This content would not show up in any of the collections it was added to with `tags` higher up in 
the data cascade.

## Collection Item Data Structure

```
<ul>
{%- for post in collections.post -%}
  <li>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```
<ul>
{%- for post in collections.post -%}
  <li>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```js
export function render(data) {
    return \`<ul>
    ${data.collections.post
            .map((post) => \`<li>${post.data.title}</li>\`)
            .join("\n")}
  </ul>\`;
};
```

```js
exports.render = function (data) {
    return \`<ul>
    ${data.collections.post
            .map((post) => \`<li>${post.data.title}</li>\`)
            .join("\n")}
  </ul>\`;
};
```

Note in the above example that we output the `post.data.title` value? Similarly, each collection 
item will have the following data:

- `page`: everything in [Eleventy’s supplied page 
variable](https://www.11ty.dev/docs/data-eleventy-supplied/#page-variable) for this template 
(including `inputPath`, `url`, `date`, and others). Added in v2.0.0
- `data`: all data for this piece of content (includes any data inherited from layouts)
- `rawInput`: the raw input of the template (before any processing). This does *not* include front 
matter. Added in v3.0.0 *(Related: [#1206](https://github.com/11ty/eleventy/issues/1206))*
- `content`: the rendered content of this template. This does *not* include layout wrappers. Added 
in v2.0.0

```json
{
  page: {
    inputPath: './test1.md',
    url: '/test1/',
    date: new Date(),
    // … and everything else in Eleventy’s \`page\`
  },
  data: { title: 'Test Title', tags: ['tag1', 'tag2'], date: 'Last Modified', /* … */ },
  content: '<h1>Test Title</h1>\n\n<p>This is text content…',

  // Available in v3.0.0 and newer:
  rawInput: '<h1>{{ title }}</h1>\n\n<p>This is text content…',
}
```

*Backwards compatibility notes:*

- Top level properties for `inputPath`, `fileSlug`, `outputPath`, `url`, `date` are still 
available, though use of `page.*` Added in v2.0.0 for these is encouraged moving forward.
- `content` Added in v2.0.0 is aliased to the previous property `templateContent`.

You can [view the previous Collection Item Data Structure docs for 
1.0](https://v1-0-2.11ty.dev/docs/collections/#collection-item-data-structure).

## Sorting

The default collection sorting algorithm sorts in ascending order using:

1. The input file’s Created Date (you can override using `date` in front matter, as shown below)
2. Files created at the exact same time are tie-broken using the input file’s full path including 
filename

For example, assume I only write blog posts on New Years Day:

```
posts/postA.md (created on 2008-01-01)
posts/postB.md (created on 2008-01-01)
posts/post3.md (created on 2007-01-01)
another-posts/post1.md (created on 2011-01-01)
```

This collection would be sorted like this:

1. `posts/post3.md`
2. `posts/postA.md`
3. `posts/postB.md`
4. `another-posts/post1.md`

### Sort descending

To sort descending in your template, you can use a filter to reverse the sort order. For example, 
it might look like this:

```
<ul>
{%- for post in collections.post reversed -%}
  <li>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```
<ul>
{%- for post in collections.post | reverse -%}
  <li>{{ post.data.title }}</li>
{%- endfor -%}
</ul>
```

```js
export function render(data) {
  // \`toReversed\` is Node 20+, see the note below
  let posts = data.collections.post.toReversed();
  return \`<ul>
    ${posts.map((post) => \`<li>${post.data.title}</li>\`).join("\n")}
  </ul>\`;
};
```

```js
exports.render = function (data) {
  // \`toReversed\` is Node 20+, see the note below
  let posts = data.collections.post.toReversed();
  return \`<ul>
    ${posts.map((post) => \`<li>${post.data.title}</li>\`).join("\n")}
  </ul>\`;
};
```

### Do not use Array reverse()

You should ***not*** use Array `reverse()` on collection arrays in your templates, like so:

`{%- for post in collections.post.reverse() -%}`

This will [mutate the array](https://doesitmutate.xyz/reverse/) and re-order it *in-place* and will 
have side effects for any use of that collection in other templates.

This is a [**Common Pitfall**](https://www.11ty.dev/docs/pitfalls/).

This applies any time you use `reverse`, for example in a custom shortcode:

eleventy.config.js

```js
export default function (eleventyConfig) {
    eleventyConfig.addShortcode("myShortcode", function (aCollection){
      // WARNING
      aCollection.reverse();
    })
};
```

```js
module.exports = function (eleventyConfig) {
    eleventyConfig.addShortcode("myShortcode", function (aCollection){
      // WARNING
      aCollection.reverse();
    })
};
```

Instead of `reverse` use:

- [JavaScript’s `.toReversed()` method](https://doesitmutate.xyz/toreversed/) (Node 20+)
- Create your own new array using [JavaScript `.filter(entry => 
entry).reverse()`](https://doesitmutate.xyz/toreversed/)
- [Liquid’s `reverse` filter](https://liquidjs.com/filters/reverse.html)
- [Nunjucks’ `reverse` filter](https://mozilla.github.io/nunjucks/templating.html#reverse)

### Overriding Content Dates

You can modify how a piece of content is sorted in a collection by changing its default `date`. 
[Read more at Content Dates](https://www.11ty.dev/docs/dates/).

```markdown
---
date: 2016-01-01
---
```

## Advanced: Custom Filtering and Sorting

This part of the docs has moved to its own page: [Collections 
API](https://www.11ty.dev/docs/collections-api/)

## From the Community

[![https-www-ovl-design-text-how-i-built-around-the-web](https://screenshot.11ty.app/https%3A%2F%2Fw
ww.ovl.design%2Ftext%2Fhow-i-built-around-the-web%2F/small/1:1/_wait:2/)](https://www.ovl.design/tex
t/how-i-built-around-the-web/ "How I built Around the Web")

[![https-darekkay-com-blog-eleventy-group-posts-by-year](https://screenshot.11ty.app/https%3A%2F%2Fd
arekkay.com%2Fblog%2Feleventy-group-posts-by-year%2F/small/1:1/_wait:2/)](https://darekkay.com/blog/
eleventy-group-posts-by-year/ "Group posts by year")

[![https-www-joshcanhelp-com-eleventy-custom-content-type-collections](https://screenshot.11ty.app/h
ttps%3A%2F%2Fwww.joshcanhelp.com%2Feleventy-custom-content-type-collections%2F/small/1:1/_wait:2/)](
https://www.joshcanhelp.com/eleventy-custom-content-type-collections/ "Custom Content Collections")

[![https-www-pborenstein-dev-posts-collections](https://screenshot.11ty.app/https%3A%2F%2Fwww.pboren
stein.dev%2Fposts%2Fcollections%2F/small/1:1/_wait:2/)](https://www.pborenstein.dev/posts/collection
s/ "Working with Collections")

\+ [Add yours!](https://github.com/11ty/11ty-website/tree/main/src/_data/community/)

×114 resources via **[11tybundle.dev](https://11tybundle.dev/)** curated by [Bob 
Monsour](https://www.bobmonsour.com/).

- [How to Provide a “Random Post” Feature With Eleventy and 
PHP](https://meiert.com/blog/random-posts-with-eleventy-and-php/)  —  *Jens Oliver Meiert 
(2026)*
- [Cleaning House: Refactoring My Eleventy Config Into 
Modules](https://brennan.day/cleaning-house-refactoring-my-eleventy-config-into-modules/)  —  
*Brennan Kenneth Brown (2026)*
- [How this site works now](https://cassie.ink/how-this-site-works-now/)  —  *Cassie (2026)*
- [Speeding Up Large 11ty Builds on 
Netlify](https://www.ryangittings.co.uk/blog/11ty-netlify-build-performance/)  —  *Ryan 
Gittings (2026)*
- [The Limits of AI and Where Humans 
Shine](https://blog.nicholas.clooney.io/posts/how-clever-is-ai/)  —  *Nicholas Clooney (2026)*
***Expand to see 109 more resources.***
- [Stats Page with Eleventy](https://kittygiraudel.com/2026/03/02/stats-page-with-11ty/)  —  
*Kitty Giraudel (2026)*
- [Backlinks in Eleventy](https://codemacabre.com/notes/backlinks-in-eleventy/)  —  *Myles 
Lewando (2026)*
- [Creating an Alphabetical Tag Page feat. Nunjucks 
Pitfalls](https://brennan.day/creatine-an-alphabetical-tag-page-feat-nunjucks-pitfalls/)  —  
*Brennan Kenneth Brown (2026)*
- [Eleventy: Sorted Collection](https://lea.codes/posts/2026-01-21-eleventy-sorted-collections/) 
 —  *Lea Rosema (2026)*
- [Syndicating posts to mastodon via the gitlab 
pipeline](https://grgml.xyz/blog/syndicating-posts-to-mastodon-via-the-gitlab-pipeline/)  —  
*Grigor Malo (2026)*
- [auld lang syne: The Commonplace 
(micro)Log](https://brennan.day/auld-lang-syne-the-commonplace-micro-log/)  —  *Brennan 
Kenneth Brown (2026)*
- [Migrating from WordPress to 11ty 
(Eleventy)](https://alfredbaudisch.com/blog/migrating-from-wordpress-to-11ty/)  —  *Alfred 
Reinold Baudisch (2025)*
- [11tyCMS: THE Eleventy Meetup appearance, metadata improvements, more post and collection 
goodness](https://11tycms.com/blog/posts/11tycms-11ty-meetup-metadata-improvements-more-post-managem
ent-and-collections/)  —  *Jessie Heald (2025)*
- [Building My New Archives 
Page](https://kpwags.com/posts/2025/09/25/building-my-new-archives-page/)  —  *Keith Wagner 
(2025)*
- [From Dotclear to Eleventy 
4](https://alix.guillard.fr/notes/dotclear-to-eleventy/page-navigation/)  —  *Alix Guillard 
(2025)*
- [Category and Tag Pages with 
Eleventy](https://joshtronic.com/2025/09/07/eleventy-category-tag-pages/)  —  *Josh Sherman 
(2025)*
- [Data vs. Collections in Eleventy](https://madbaker.com/blog/data-vs-collections-in-eleventy/) 
 —  *Mark Dyck (2025)*
- [Building a digital bookshelf with 
Eleventy](https://damianwalsh.co.uk/posts/building-a-digital-bookshelf-with-eleventy/)  —  
*Damian Walsh (2025)*
- [Double-Pagination in 
Elev­enty](https://chriskirknielsen.com/blog/double-pagination-in-eleventy/)  —  
*Christopher Kirk-Nielsen (2025)*
- [Creating a Journal With 
Eleventy](https://www.simplethread.com/creating-a-journal-with-eleventy/)  —  *Austin Carr 
(2025)*
- [Added a 'uses' Page Archive](https://ryan.himmelwright.net/post/added-uses-page-archive/) 
 —  *Ryan Himmelwright (2025)*
- [List blog posts grouped by year with 
Eleventy](https://hamatti.org/posts/list-blog-posts-grouped-by-year-with-eleventy/)  —  
*Juha-Matti Santala (2025)*
- [Building a personal digital music library with Eleventy and 
APIs](https://damianwalsh.co.uk/posts/creating-connections-with-music-and-technology/)  —  
*Damian Walsh (2025)*
- [Display Plausible Statistics in Your 11ty 
Blog](https://www.jjude.com/tech-notes/plausible-and-11ty/)  —  *Joseph Jude (2025)*
- [HTTP 301 redirects in Eleventy](https://danburzo.ro/notes/eleventy-301-redirects/)  —  
*Dan Cătălin Burzo (2024)*
- [Adding Image Galleries to My Website](https://nathanupchurch.com/blog/galleries/)  —  
*Nathan Upchurch (2024)*
- [Oops, I built a headless frontend with 11ty](https://aaadaaam.com/notes/headless-frontend-11ty/) 
 —  *Adam Stoddard (2024)*
- [Here's how this is all put 
together](https://www.coryd.dev/posts/2024/heres-how-this-is-all-put-together)  —  *Cory 
Dransfeldt (2024)*
- [Adding Cooklang Support to Eleventy Three 
Ways](https://rknight.me/blog/adding-cooklang-support-to-eleventy-two-ways/)  —  *Robb Knight 
(2024)*
- [Eleventy (11ty) year, year-month, and year-monty-day 
indexes](https://blog.tomayac.com/2024/11/02/eleventy-11ty-year-year-month-and-year-month-day-indexe
s/)  —  *Thomas Steiner (2024)*
- [Building an album releases calendar 
subscription](https://www.coryd.dev/posts/2024/building-an-album-releases-calendar-subscription/) 
 —  *Cory Dransfeldt (2024)*
- [Building a Blog with 
Eleventy](https://blog.sebin-nyshkim.net/posts/building-a-blog-with-eleventy-blind-any/)  —  
*Sebin Nyshkim (2024)*
- [Aggregating content using collections in 
Eleventy](https://www.coryd.dev/posts/2024/aggregating-content-using-collections-in-eleventy/) 
 —  *Cory Dransfeldt (2024)*
- [Eleventy Collections from an 
API](https://www.trovster.com/blog/2024/09/eleventy-collections-from-an-api)  —  *Trevor 
Morris (2024)*
- [Syndicating an RSS feed to Mastodon using a Cloudflare 
worker](https://www.coryd.dev/posts/2024/syndicating-an-rss-feed-to-mastodon-using-a-cloudflare-work
er/)  —  *Cory Dransfeldt (2024)*
- [Exclude specific tags in Eleventy using a custom 
filter](https://cri.dev/posts/2024-09-21-how-to-exclude-tags-collection-filter-eleventy/)  —  
*Christian Fei (2024)*
- [Dynamic Importing with 
Eleventy](https://www.trovster.com/blog/2024/08/dynamic-importing-with-eleventy)  —  *Trevor 
Morris (2024)*
- [Eleventy Filters in 
Collections](https://www.trovster.com/blog/2024/08/eleventy-filters-in-collections)  —  
*Trevor Morris (2024)*
- [Building a Custom Filter for Eleventy 
Collections](https://ttntm.me/blog/building-a-custom-filter-for-eleventy-collections/)  —  
*Tom Doe (2024)*
- [Twenty year celebration: Site update number three](https://ptsefton.com/2024/06/24/update/) 
 —  *Peter Sefton (2024)*
- [Category and genre pages 
return](https://arestelle.net/2024/06/20/category-and-genre-pages-return/)  —  *Nicki Hoffman 
(2024)*
- [Eleventy collections using the built in tags 
key](https://design2seo.com/blog/web-development/11ty/eleventy-collections-using-tags/)  —  
*Jeremy Faucher (2024)*
- [Eleventy - Convert a RSS Feed to a 
collection](https://www.roboleary.net/2024/06/05/eleventy-rss-collection.html)  —  *Rob 
O'Leary (2024)*
- [Integrating a (somewhat) custom CMS with 
Eleventy](https://www.coryd.dev/posts/2024/integrating-a-somewhat-custom-cms-with-eleventy/) 
 —  *Cory Dransfeldt (2024)*
- [Updating to Eleventy v3](https://mxb.dev/blog/eleventy-v3-update/)  —  *Max Böck (2024)*
- [Eleventy - Merge external data with an existing 
collection](https://www.roboleary.net/2024/05/26/eleventy-external-posts.html)  —  *Rob 
O'Leary (2024)*
- [A custom collection to sort events with 
Eleventy](https://samimaatta.fi/en/a-custom-collection-to-sort-events-with-eleventy/)  —  
*Sami Määttä (2024)*
- [Supporting a full-text RSS 
feed](https://dustinwhisman.com/writing/supporting-a-full-text-rss-feed/)  —  *Dustin Whisman 
(2024)*
- [Surfing The Web And Sharing What I 
Find](https://flamedfury.com/posts/surfing-the-web-and-sharing-what-i-find/)  —  *fLaMEd 
(2024)*
- [Notifications for New Eleventy Posts in GitLab - Part 
2](https://aarongoldenthal.com/posts/notifications-for-new-eleventy-posts-in-gitlab-part-2/) 
 —  *Aaron Goldenthal (2024)*
- [Eleventy Nested Pagination](https://www.codeflood.net/blog/2024/04/17/11ty-nested-pagination/) 
 —  *Alistair Deneys (2024)*
- [Notifications for New Eleventy Posts in GitLab - Part 
1](https://aarongoldenthal.com/posts/notifications-for-new-eleventy-posts-in-gitlab-part-1/) 
 —  *Aaron Goldenthal (2024)*
- [Eleventy Navigation Set URL to First Item in 
Collection](https://johnwargo.com/posts/2024/eleventy-navigation-set-url/)  —  *John M. Wargo 
(2024)*
- [11ty collections tag 
links](https://www.simoncox.com/shorts/2024-03-17-11ty-collection-tag-links/)  —  *Simon Cox 
(2024)*
- [Group posts by year in 
Eleventy.js](https://simpixelated.com/group-posts-by-year-in-eleventy-js/)  —  *Jordan Kohl 
(2024)*
- [Surfacing most used tags in 
Eleventy](https://www.coryd.dev/posts/2024/surfacing-most-used-tags-in-eleventy/)  —  *Cory 
Dransfeldt (2024)*
- [Draft Posts in Eleventy](https://olets.dev/posts/draft-posts-in-eleventy/)  —  *Henry 
Bley-Vroman (2024)*
- [My Eleventy site setup](https://anhvn.com/posts/2024/my-eleventy-site-setup/)  —  *anh 
(2024)*
- [Eleventy - Group posts by 
year](https://www.roboleary.net/2024/03/01/eleventy-posts-by-year.html)  —  *Rob O'Leary 
(2024)*
- [A roundup of recent updates to my 
website](https://grgml.xyz/blog/website-updates-bookmarks-status-and-more/)  —  *Grigør 
(2024)*
- [Intro: Bukmark.club](https://ttntm.me/blog/bukmark-club-intro/)  —  *Tom Doe (2024)*
- [Update: Bookmarks Are Back](https://ttntm.me/blog/bookmarks-are-back/)  —  *Tom Doe (2024)*
- [Eleventy 🤝 Immich](https://chris.bur.gs/eleventy-immich/)  —  *Chris Burgess (2024)*
- [Low-tech Eleventy Categories](https://piccalil.li/blog/low-tech-eleventy-categories/)  —  
*Andy Bell (2024)*
- [Getting up to Speed with Eleventy: Config and 
Collections](https://thenewstack.io/getting-up-to-speed-with-eleventy-config-and-collections/) 
 —  *David Eastman (2024)*
- [Update: Tags Are Back](https://ttntm.me/blog/restoring-tags/)  —  *Tom Doe (2024)*
- [Right here, right now](https://www.martingunnarsson.com/posts/right-here-right-now/)  —  
*Martin Gunnarsson (2024)*
- [Generating a static blog with 
Eleventy](https://tfedder.de/blog/generating-a-static-blog-with-eleventy/)  —  *Tobias Fedder 
(2023)*
- [Grouping posts by year with nunjucks in 
Eleventy](https://chriskirknielsen.com/blog/group-posts-by-year-with-nunjucks-in-eleventy/) 
 —  *Christopher Kirk-Nielsen (2023)*
- [Generating the Firehose page on the 11tybundle 
site](https://bobmonsour.com/blog/generating-the-firehose-page-on-the-11tybundle-site/)  —  
*Bob Monsour (2023)*
- [Building post types and category RSS feeds in 
Eleventy](https://localghost.dev/blog/building-post-types-and-category-rss-feeds-in-eleventy/) 
 —  *Sophie Koonin (2023)*
- [Eleventy Splitting Category Data Across Two Table 
Columns](https://johnwargo.com/posts/2023/eleventy-two-table-columns/)  —  *John M. Wargo 
(2023)*
- [Making a simple Eleventy blog 
template](https://dev.to/graphqleditor/making-a-simple-eleventy-blog-template-3558)  —  
*Tomek Poniatowicz (2023)*
- [Making Author Pages for an Academic Journal in Eleventy, or, How to Manipulate Collection Data 
in Eleventy](https://micah.torcellini.org/2023/09/23/author-page-manipulate-collections/)  —  
*Micah Torcellini (2023)*
- [How to create a drafts page for an 11ty blog](https://www.lkhrs.com/blog/2023/11ty-drafts-page/) 
 —  *Luke Harris (2023)*
- [Semi-automated hashtags for syndicated 
posts](https://www.coryd.dev/posts/2023/semi-automated-hashtags-for-syndicated-posts/)  —  
*Cory Dransfeldt (2023)*
- [CloudCannon + Eleventy](https://claytonerrington.com/blog/cloudcannon-eleventy)  —  
*Clayton Errington (2023)*
- [Rethinking Categorization](https://lea.verou.me/blog/2023/rethinking-categorization/)  —  
*Lea Verou (2023)*
- [11ty: Index ALL the things!](https://lea.verou.me/blog/2023/11ty-indices/) —  *Lea Verou 
(2023)*
- [Filtering tags within Eleventy.js 
collections](https://simpixelated.com/filtering-tags-within-eleventy-js-collections/)  —  
*Jordan Kohl (2023)*
- [Eleventy Collection Schemas](https://11ty.rocks/posts/eleventy-collection-schemas/)  —  
*Stephanie Eckles (2023)*
- [Adding tag list with post count to 
Eleventy.js](https://simpixelated.com/adding-tag-list-with-post-count-to-eleventy-js/)  —  
*Jordan Kohl (2023)*
- [How to programmatically add tags to posts in 
11ty](https://photogabble.co.uk/tutorials/how-to-programmatically-add-tags-to-posts-in-11ty/) 
 —  *Simon Dann (2023)*
- [Support Draft Blog Posts in 
Eleventy](https://www.raymondcamden.com/2022/08/14/support-draft-blog-posts-in-eleventy)  —  
*Raymond Camden (2022)*
- [Creating a now page archive with 
11ty](https://photogabble.co.uk/tutorials/creating-a-now-page-archive-with-11ty/)  —  *Simon 
Dann (2022)*
- [11ty aliases the right way](https://boehs.org/node/11ty-aliases)  —  *Evan Boehs (2022)*
- [Zero maintenance taxonomies in 11ty](https://boehs.org/node/11ty-taxonomies)  —  *Evan 
Boehs (2022)*
- [Eleventy Custom Content Type Collections and 
Layouts](https://www.joshcanhelp.com/eleventy-custom-content-type-collections/)  —  *Josh 
Cunningham (2022)*
- [Creating and Using Eleventy 
Collections](https://11ty.rocks/posts/creating-and-using-11ty-collections/)  —  *Stephanie 
Eckles (2021)*
- [Making an 11ty collection from a remote XML 
file](https://www.mikestreety.co.uk/blog/making-an-11ty-collection-from-a-remote-xml-file/) 
 —  *Mike Street (2021)*
- [Create an 11ty collection from any RSS 
feed](https://www.mikestreety.co.uk/blog/create-11ty-collection-from-rss/)  —  *Mike Street 
(2021)*
- [Creating an 11ty collection from a JSON 
API](https://www.mikestreety.co.uk/blog/creating-an-11ty-collection-from-json-api/)  —  *Mike 
Street (2021)*
- [Dynamic Short URLs with 
Eleventy](https://www.raymondcamden.com/2021/06/22/dynamic-short-urls-with-eleventy)  —  
*Raymond Camden (2021)*
- [Building Ale House Rock with 
11ty](https://www.mikestreety.co.uk/blog/building-ale-house-rock-with-11ty/)  —  *Mike Street 
(2021)*
- [Create posts in Eleventy using 
collections](https://events.lunch.dev/create-posts-in-eleventy-using-collections/)  —  
*Michael Chan (2021)*
- [Using 11ty JavaScript Data files to mix Markdown and CMS content into one 
collection](https://bryanlrobinson.com/blog/using-11ty-javascript-data-files-to-mix-markdown-and-cms
-content-into-one-collection/)  —  *Bryan Robinson (2021)*
- [Build an 11ty calendar to list all your 
posts](https://www.mikestreety.co.uk/blog/build-an-11ty-calendar-to-list-all-your-posts/)  —  
*Mike Street (2021)*
- [Collection archive in Eleventy](https://bnijenhuis.nl/notes/collection-archive-in-eleventy/) 
 —  *Bernard Nijenhuis (2021)*
- [Adding Simple Pagination to an 11ty 
Collection](https://www.brianperry.dev/til/2021/adding-simple-pagination-to-an-11ty-collection/) 
 —  *Brian Perry (2021)*
- [Using Template Content as Data for 
11ty](https://11ty.rocks/posts/using-template-content-as-data/)  —  *Stephanie Eckles (2021)*
- [Group posts by year in Eleventy](https://darekkay.com/blog/eleventy-group-posts-by-year/) 
 —  *Darek Kay (2021)*
- [Custom collection list markup in 
Eleventy](https://www.dandenney.com/posts/front-end-dev/custom-collection-list-markup-in-eleventy/) 
 —  *Dan Denney (2020)*
- [Create a Custom Collection with 
Eleventy](https://tannerdolby.com/writing/create-a-custom-collection-with-eleventy/)  —  
*Tanner Dolby (2020)*
- [Enhancing archives navigation, step 
1](https://nicolas-hoizey.com/articles/2020/10/26/enhancing-archives-navigation-step-1/)  —  
*Nicolas Hoizey (2020)*
- [Let's Learn Eleventy (11ty) - Collections](https://www.raresportan.com/eleventy-part-two/) 
 —  *Rares Portan (2020)*
- [Hiding Future Content with 
Eleventy](https://www.raymondcamden.com/2020/08/07/hiding-future-content-with-eleventy)  —  
*Raymond Camden (2020)*
- [Eleventy: Loop through a collection from within a 
layout?](https://dev.to/granttransition/eleventy-loop-through-a-collection-from-within-a-layout-1h4i
) —  *Grant Smith (2020)*
- [Basic custom taxonomies with 
Eleventy](https://www.webstoemp.com/blog/basic-custom-taxonomies-with-eleventy/)  —  
*Jérôme Coupé (2020)*
- [Adding a Last Edited Field to 
Eleventy](https://www.raymondcamden.com/2020/02/06/adding-a-last-edited-field-to-eleventy) 
 —  *Raymond Camden (2020)*
- [Adding Search to your Eleventy Static Site with 
Lunr](https://www.raymondcamden.com/2019/10/20/adding-search-to-your-eleventy-static-site-with-lunr)
  —  *Raymond Camden (2019)*
- [Flexible tag-like functionality for custom keys in 
Eleventy](https://fuzzylogic.me/posts/flexible-tag-like-functionality-for-custom-keys-in-eleventy/) 
 —  *Laurence Hughes (2019)*
- [Scheduled and draft 11ty posts](https://remysharp.com/2019/06/26/scheduled-and-draft-11ty-posts) 
 —  *Remy Sharp (2019)*
- [Implementing categories](https://pborenstein.dev/posts/categories/)  —  *Philip Borenstein 
(2019)*
- [Working with collections](https://pborenstein.dev/posts/collections/)  —  *Philip 
Borenstein (2019)*

---

### Other pages in Configure Templates with Data

- [Permalinks](https://www.11ty.dev/docs/permalinks/)
- [Layouts](https://www.11ty.dev/docs/layouts/)
- [Collections](https://www.11ty.dev/docs/collections/)
	- [Collections API](https://www.11ty.dev/docs/collections-api/)
- [Content Dates](https://www.11ty.dev/docs/dates/)
- [Create Pages From Data](https://www.11ty.dev/docs/pages-from-data/)
	- [Pagination](https://www.11ty.dev/docs/pagination/)
		- [Pagination Navigation](https://www.11ty.dev/docs/pagination/nav/)
