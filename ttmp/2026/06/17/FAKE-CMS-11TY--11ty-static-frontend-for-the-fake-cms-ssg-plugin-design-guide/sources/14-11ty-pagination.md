---
title: "Pagination — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/pagination/
summary: "Downloaded with defuddle. Paginating a dataset to emit one page per item; permalink aliases. Re-run: defuddle parse https://www.11ty.dev/docs/pagination/ --md | fold -w 100 -s."
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

Pagination allows you to iterate over a data set and create multiple files from a single template. 
The input data can be in the form of an array or object defined in your frontmatter or in [global 
data](https://www.11ty.dev/docs/data-global/), or you can paginate a collection to make an easily 
digestible list of your posts.

## Paging an Array

To iterate over a data set and create pages for individual chunks of data, use pagination. Enable 
in your template’s front matter by adding the `pagination` key.

Consider the following template, which will result in two pages being created, each of which will 
display two items from `testdata`:

```
---
pagination:
  data: testdata
  size: 2
testdata:
 - item1
 - item2
 - item3
 - item4
---
<ol>
{%- for item in pagination.items %}
  <li>{{ item }}</li>
{% endfor -%}
</ol>
```

If the above file were named `paged.liquid`, it would create two pages in your output folder: 
`_site/paged/index.html` and `_site/paged/1/index.html`. These output paths are configurable with 
`permalink` (see below).

```
---
pagination:
  data: testdata
  size: 2
testdata:
 - item1
 - item2
 - item3
 - item4
---
<ol>
{%- for item in pagination.items %}
  <li>{{ item }}</li>
{% endfor -%}
</ol>
```

If the above file were named `paged.njk`, it would create two pages in your output folder: 
`_site/paged/index.html` and `_site/paged/1/index.html`. These output paths are configurable with 
`permalink` (see below).

```js
export const data = {
  pagination: {
    data: "testdata",
    size: 2
  },
  testdata: [
    "item1",
    "item2",
    "item3",
    "item4"
  ]
};

export function render(data) {
  return \`<ol>
    ${data.pagination.items.map(function(item) {
        return \`<li>${item}</li>\`;
      }).join("")
    }
  </ol>\`;
};
```

If the above file were named `paged.11ty.js`, it would create two pages in your output folder: 
`_site/paged/index.html` and `_site/paged/1/index.html`. These output paths are configurable with 
`permalink` (see below).

```js
exports.data = {
  pagination: {
    data: "testdata",
    size: 2
  },
  testdata: [
    "item1",
    "item2",
    "item3",
    "item4"
  ]
};

exports.render = function(data) {
  return \`<ol>
    ${data.pagination.items.map(function(item) {
        return \`<li>${item}</li>\`;
      }).join("")
    }
  </ol>\`;
};
```

If the above file were named `paged.11ty.js`, it would create two pages in your output folder: 
`_site/paged/index.html` and `_site/paged/1/index.html`. These output paths are configurable with 
`permalink` (see below).

We enable pagination and then give it a dataset with the `data` key. We control the number of items 
in each chunk with `size`. The pagination data variable will be populated with what you need to 
create each template. Here’s what’s in `pagination`:

**Syntax** JavaScript Object

```js
{
  items: [], // Array of current page’s chunk of data
  pageNumber: 0, // current page number, 0 indexed

  // Cool URLs
  hrefs: [], // Array of all page hrefs (in order)
  href: {
    next: "…", // put inside <a href="{{ pagination.href.next }}">Next Page</a>
    previous: "…", // put inside <a href="{{ pagination.href.previous }}">Previous Page</a>
    first: "…",
    last: "…",
  },

  pages: [], // Array of all chunks of paginated data (in order)
  page: {
    next: {}, // Next page’s chunk of data
    previous: {}, // Previous page’s chunk of data
    first: {},
    last: {},
  }
}
```

Expand to see all of the extra stuff in the `pagination` object that you probably don’t need any 
more but it’s still in there for backwards compatibility.

In addition to the `pagination` object entries documented above, it also has:

```js
{
  data: "…", // the original string key to the dataset
  size: 1, // page chunk sizes

  // Cool URLs
  // Use pagination.href.next, pagination.href.previous, et al instead.
  nextPageHref: "…", // put inside <a href="{{ pagination.nextPageHref }}">Next Page</a>
  previousPageHref: "…", // put inside <a href="{{ pagination.previousPageHref }}">Previous 
Page</a>
  firstPageHref: "…",
  lastPageHref: "…",

  // Uncool URLs
  // These include index.html file names, use \`hrefs\` instead
  links: [], // Array of all page links (in order)

  // Deprecated things:
  // nextPageLink
  // previousPageLink
  // firstPageLink
  // lastPageLink
  // pageLinks (alias to \`links\`)
}
```

## Creating Navigation Links to your Pages

Learn how to create a list of links to every paginated page on a pagination template with a full 
[Pagination Navigation](https://www.11ty.dev/docs/pagination/nav/) tutorial.

## Paging an Object

All of the examples thus far have paged Array data. Eleventy does allow paging objects too. Objects 
are resolved to pagination arrays using either the `Object.keys` or `Object.values` JavaScript 
functions. Consider the following templates:

```
---
pagination:
  data: testdata
  size: 1
testdata:
  itemkey1: itemvalue1
  itemkey2: itemvalue2
  itemkey3: itemvalue3
---
<ol>
{%- for item in pagination.items %}
  <li>{{ item }}={{testdata[item] }}</li>
{% endfor -%}
</ol>
```

```
---
pagination:
  data: testdata
  size: 1
testdata:
  itemkey1: itemvalue1
  itemkey2: itemvalue2
  itemkey3: itemvalue3
---
<ol>
{%- for item in pagination.items %}
  <li>{{ item }}={{testdata[item] }}</li>
{% endfor -%}
</ol>
```

```js
export const data = {
    pagination: {
        data: "testdata",
        size: 1,
    },
    testdata: {
        itemkey1: "itemvalue1",
        itemkey2: "itemvalue2",
        itemkey3: "itemvalue3",
    },
};

export function render(data) {
    return \`<ol>
        ${data.pagination.items
            .map(function (item) {
                return \`<li>${(item = data.testdata[item])}</li>\`;
            })
            .join("")}
    </ol>\`;
};
```

```js
exports.data = {
    pagination: {
        data: "testdata",
        size: 1,
    },
    testdata: {
        itemkey1: "itemvalue1",
        itemkey2: "itemvalue2",
        itemkey3: "itemvalue3",
    },
};

exports.render = function (data) {
    return \`<ol>
        ${data.pagination.items
            .map(function (item) {
                return \`<li>${(item = data.testdata[item])}</li>\`;
            })
            .join("")}
    </ol>\`;
};
```

In this example, we would get 3 pages that each print a key/value pair from `testdata`. The paged 
items hold the object keys:

**Syntax** JavaScript Object

```js
[
    ["itemkey1"], // pagination.items[0] holds the object key
    ["itemkey2"],
    ["itemkey3"],
];
```

You can use these keys to get access to the original value: `testdata[ pagination.items[0] ]`.

If you’d like the pagination to iterate over the values instead of the keys (using 
`Object.values` instead of `Object.keys`), add `resolve: values` to your `pagination` front matter:

**Syntax** YAML Front Matter

```markdown
---
pagination:
  data: testdata
  size: 1
  resolve: values
testdata:
  itemkey1: itemvalue1
  itemkey2: itemvalue2
  itemkey3: itemvalue3
---
```

This resolves to:

**Syntax** JavaScript Object

```js
[
    ["itemvalue1"], // pagination.items[0] holds the object value
    ["itemvalue2"],
    ["itemvalue3"],
];
```

## Paginate a global or local data file

[Read more about Template Data Files](https://www.11ty.dev/docs/data/). The only change here is 
that you point your `data` pagination key to the global or local data instead of data in the front 
matter. For example, consider the following `globalDataSet.json` file in your global data directory.

**Syntax** JavaScript Object

```json
{
    "myData": ["item1", "item2", "item3", "item4"]
}
```

Your front matter would look like this:

```
---
pagination:
  data: globalDataSet.myData
  size: 1
---
<ol>
{%- for item in pagination.items %}
  <li>{{ item }}</li>
{% endfor -%}
</ol>
```

```
---
pagination:
  data: globalDataSet.myData
  size: 1
---
<ol>
{%- for item in pagination.items %}
  <li>{{ item }}</li>
{% endfor -%}
</ol>
```

```js
export const data = {
    pagination: {
        data: "globalDataSet.myData",
        size: 1,
    },
};

export function render(data) {
    return \`<ol>
    ${data.pagination.items
            .map(function (item) {
                return \`<li>${item}</li>\`;
            })
            .join("")}
  </ol>\`;
};
```

```js
exports.data = {
    pagination: {
        data: "globalDataSet.myData",
        size: 1,
    },
};

exports.render = function (data) {
    return \`<ol>
    ${data.pagination.items
            .map(function (item) {
                return \`<li>${item}</li>\`;
            })
            .join("")}
  </ol>\`;
};
```

## Remapping with permalinks

Normally, front matter does not support template syntax, but `permalink` does, enabling parametric 
URLs via pagination variables. Here’s an example of a permalink using the pagination page number:

**Syntax** YAML Front Matter using Liquid, Nunjucks

```markdown
---
permalink: "different/page-{{ pagination.pageNumber }}/index.html"
---
```

Writes to `_site/different/page-0/index.html`, `_site/different/page-1/index.html`, et cetera.

That means Nunjucks will also let you start your page numbers with 1 instead of 0, by just adding 1 
here:

**Syntax** YAML Front Matter using Nunjucks

```markdown
---
permalink: "different/page-{{ pagination.pageNumber + 1 }}/index.html"
---
```

Writes to `_site/different/page-1/index.html`, `_site/different/page-2/index.html`, et cetera.

You can even use template logic here too:

```markdown
---
permalink: "different/{% if pagination.pageNumber > 0 %}page-{{ pagination.pageNumber + 1 }}/{% 
endif %}index.html"
---
```

Writes to `_site/different/index.html`, `_site/different/page-2/index.html`, et cetera.

Note that the above example works in Nunjucks but `{{ pagination.pageNumber + 1 }}` is not 
supported in Liquid. Use `{{ pagination.pageNumber | plus: 1 }}` instead.

### Use page item data in the permalink

You can do more advanced things like this:

**Syntax** YAML Front Matter using Liquid, Nunjucks

```markdown
---
pagination:
  data: testdata
  size: 1
testdata:
  - My Item
permalink: "different/{{ pagination.items[0] | slugify }}/index.html"
---
```

Using a universal `slug` filter (transforms `My Item` to `my-item`), this outputs: 
`_site/different/my-item/index.html`.

## Aliasing to a different variable

Ok, so `pagination.items[0]` is ugly. We provide an option to alias this to something different.

```
---
pagination:
  data: testdata
  size: 1
  alias: wonder
testdata:
  - Item1
  - Item2
permalink: "different/{{ wonder | slugify }}/index.html"
---
You can use the alias in your content too {{ wonder }}.
```

```
---
pagination:
  data: testdata
  size: 1
  alias: wonder
testdata:
  - Item1
  - Item2
permalink: "different/{{ wonder | slugify }}/index.html"
---
You can use the alias in your content too {{ wonder }}.
```

```js
export const data = {
    pagination: {
        data: "testdata",
        size: 1,
        alias: "wonder",
    },
    testdata: ["Item1", "Item2"],
    permalink: function (data) {
        return \`different/${this.slugify(data.wonder)}/index.html\`;
    },
};

export function render(data) {
    return \`You can use the alias in your content too ${data.wonder}.\`;
};
```

```js
exports.data = {
    pagination: {
        data: "testdata",
        size: 1,
        alias: "wonder",
    },
    testdata: ["Item1", "Item2"],
    permalink: function (data) {
        return \`different/${this.slugify(data.wonder)}/index.html\`;
    },
};

exports.render = function (data) {
    return \`You can use the alias in your content too ${data.wonder}.\`;
};
```

This writes to `_site/different/item1/index.html` and `_site/different/item2/index.html`.

Note that `page` is a reserved word so you cannot use `alias: page`. Read about Eleventy’s 
reserved data names in [Eleventy Supplied Data](https://www.11ty.dev/docs/data-eleventy-supplied/).

If your chunk `size` is greater than 1, the alias will be an array instead of a single value.

```
---
pagination:
  data: testdata
  size: 2
  alias: wonder
testdata:
  - Item1
  - Item2
  - Item3
  - Item4
permalink: "different/{{ wonder[0] | slugify }}/index.html"
---
You can use the alias in your content too {{ wonder[0] }}.
```

```
---
pagination:
  data: testdata
  size: 2
  alias: wonder
testdata:
  - Item1
  - Item2
  - Item3
  - Item4
permalink: "different/{{ wonder[0] | slugify }}/index.html"
---
You can use the alias in your content too {{ wonder[0] }}.
```

```js
export const data = {
  pagination: {
    data: "testdata",
    size: 2,
    alias: "wonder"
  },
  testdata: [
    "Item1",
    "Item2",
    "Item3",
    "Item4"
  ],
  permalink: {
    function(data) {
      return \`different/${this.slugify(data.wonder[0])}/index.html\`
    }
  }
};

export function render(data) {
  return \`You can use the alias in your content too ${data.wonder[0]}.\`;
}
```

```js
exports.data = {
  pagination: {
    data: "testdata",
    size: 2,
    alias: "wonder"
  },
  testdata: [
    "Item1",
    "Item2",
    "Item3",
    "Item4"
  ],
  permalink: {
    function(data) {
      return \`different/${this.slugify(data.wonder[0])}/index.html\`
    }
  }
};

exports.render = function (data) {
  return \`You can use the alias in your content too ${data.wonder[0]}.\`;
}
```

This writes to `_site/different/item1/index.html` and `_site/different/item3/index.html`.

## Paging a Collection

If you’d like to make a paginated list of all of your blog posts (any content with the tag `post` 
on it), use something like the following template to iterate over a specific collection:

```
---
title: My Posts
pagination:
  data: collections.post
  size: 6
  alias: posts
---

<ol>
{% for post in posts %}
  <li><a href="{{ post.url }}">{{ post.data.title }}</a></li>
{% endfor %}
</ol>
```

```
---
title: My Posts
pagination:
  data: collections.post
  size: 6
  alias: posts
---

<ol>
{% for post in posts %}
  <li><a href="{{ post.url }}">{{ post.data.title }}</a></li>
{% endfor %}
</ol>
```

```js
export const data = {
    title: "My Posts",
    pagination: {
        data: "collections.post",
        size: 6,
        alias: "posts",
    },
};

export function render(data) {
    return \`<ol>
        ${data.posts
            .map(function (post) {
                return \`<li><a href="${post.url}">${post.title}</a></li>\`;
            })
            .join("")}
    </ol>\`;
};
```

```js
exports.data = {
    title: "My Posts",
    pagination: {
        data: "collections.post",
        size: 6,
        alias: "posts",
    },
};

exports.render = function (data) {
    return \`<ol>
        ${data.posts
            .map(function (post) {
                return \`<li><a href="${post.url}">${post.title}</a></li>\`;
            })
            .join("")}
    </ol>\`;
};
```

The above generates a list of links but you could do a lot more. See what’s available in the 
[Collection documentation](https://www.11ty.dev/docs/collections/#collection-item-data-structure) 
(specifically `templateContent`). If you’d like to use this to automatically generate Tag pages 
for your content, please read [Quick Tip #004—Create Tag Pages for your 
Blog](https://www.11ty.dev/docs/quicktips/tag-pages/).

## Generating an Empty Results Page

Added in v2.0.0

By default, if the specified data set is empty, Eleventy will not render any pages. Use 
`generatePageOnEmptyData: true` to generate one pagination output with an empty chunk `[]` of items.

**Syntax** Liquid, Nunjucks

```markdown
---
title: Available Products
pagination:
  data: collections.available
  size: 6
  generatePageOnEmptyData: true
---
```

## Modifying the Data Set prior to Pagination

### Reverse the Data

Use `reverse: true`.

```markdown
---
pagination:
  data: testdata
  size: 2
  reverse: true
testdata:
  - item1
  - item2
  - item3
  - item4
---
```

Paginates to:

```js
[
    ["item4", "item3"],
    ["item2", "item1"],
];
```

*(More discussion at [Issue #194](https://github.com/11ty/eleventy/issues/194))*

As an aside, this could also be achieved in a more verbose way using the [Collection 
API](https://www.11ty.dev/docs/collections/#advanced-custom-filtering-and-sorting). This could also 
be done using the new `before` callback.

### Filtering Values

Use the `filter` pagination property to remove values from paginated data.

**Syntax** YAML Front Matter

```markdown
---
pagination:
  data: testdata
  size: 1
  filter:
    - item3
testdata:
  item1: itemvalue1
  item2: itemvalue2
  item3: itemvalue3
---
```

Paginates to:

**Syntax** JavaScript Object

```js
[["item1"], ["item2"]];
```

This will work the same with paginated arrays or with `resolve: values` for paginated objects.

**Syntax** YAML Front Matter

```markdown
---
pagination:
  data: testdata
  size: 1
  resolve: values
  filter:
    - itemvalue3
testdata:
  item1: itemvalue1
  item2: itemvalue2
  item3: itemvalue3
---
```

Paginates to:

**Syntax** JavaScript Object

```js
[["itemvalue1"], ["itemvalue2"]];
```

### The before Callback

The most powerful tool to change the data. Use this callback to modify, filter, or otherwise change 
the pagination data however you see fit *before* pagination occurs.

```js
---js
{
  pagination: {
    data: "testdata",
    size: 2,
    before: function(paginationData, fullData) {
      // \`fullData\` is new in v1.0.1 and contains the full Data Cascade thus far

      return paginationData.map(entry => \`${entry} with a suffix\`);
    }
  },
  testdata: [
    "item1",
    "item2",
    "item3",
    "item4"
  ]
}
---
<!-- the rest of the template -->
```

The above will iterate over a data set containing: `["item1 with a suffix", "item2 with a suffix", 
"item3 with a suffix", "item4 with a suffix"]`.

You can do anything in this `before` callback. Maybe a custom `.sort()`, `.filter()`, `.map()` to 
remap the entries, `.slice()` to paginate only a subset of the data, etc!

#### Use JavaScript Template Functions here

Added in v2.0.0 [JavaScript Template 
Functions](https://www.11ty.dev/docs/languages/javascript/#javascript-template-functions) (which 
are also populated by universal filters and shortcodes) are available in the `before` callback.

```js
// …
before: function() {
  let slug = this.slugify("My title.");
  // use Universal filters or shortcodes too…
},
// …
```

### Order of Operations

If you use more than one of these data set modification features, here’s the order in which they 
operate:

- The `before` callback
- `reverse: true`
- `filter` entries

## Add All Pagination Pages to Collections

By default, any tags listed in a paginated template will only add the very first page to the 
appropriate collection.

Consider the following pagination template:

**Filename** my-page.md

```yaml
---
tags:
  - myCollection
pagination:
  data: testdata
  size: 2
testdata:
  - item1
  - item2
  - item3
  - item4
---
```

This means that `collections.myCollection` will have only the first page added to the collection 
array (`_site/my-page/index.html`). However, if you’d like to add all the pagination pages to the 
collections, use `addAllPagesToCollections: true` to the pagination front matter options like so:

**Filename** my-page.md

```yaml
---
tags:
  - myCollection
pagination:
  data: testdata
  size: 2
  addAllPagesToCollections: true
testdata:
  - item1
  - item2
  - item3
  - item4
---
```

Now `collections.myCollection` will have both output pages in the collection array 
(`_site/my-page/index.html` and `_site/my-page/1/index.html`).

## Full Pagination Option List

- `data` (String) [Lodash.get path](https://lodash.com/docs/4.17.15#get) to point to the target 
data set.
- `size` (Number, required)
- `alias` (String) [Lodash.set path](https://lodash.com/docs/4.17.15#set) to point to the property 
to set.
- `generatePageOnEmptyData` (Boolean) if target data set is empty, render first page with empty 
chunk `[]`.
- `resolve: values`
- `filter` (Array)
- `reverse: true` (Boolean)
- `addAllPagesToCollections: true` (Boolean)

## Related

## From the Community

×72 resources via **[11tybundle.dev](https://11tybundle.dev/)** curated by [Bob 
Monsour](https://www.bobmonsour.com/).

- [Serving Markdown to LLMs with 
Eleventy](https://kittygiraudel.com/2026/03/11/serving-markdown-to-llms-with-11ty/)  —  
*Kitty Giraudel (2026)*
- [Tag Pages with Eleventy](https://kittygiraudel.com/2026/03/01/tag-pages-with-eleventy/) 
 —  *Kitty Giraudel (2026)*
- [Create pages from JSON files with 
Eleventy](https://saneef.com/blog/create-pages-from-json-files-with-eleventy/)  —  *Saneef H. 
Ansari (2026)*
- [eleventy lessons](https://leecat.art/eleventy-lessons/)  —  *Lee Cattarin (2026)*
- [Building a Lightweight Optimizely SaaS CMS Solution with 
11ty](https://world.optimizely.com/blogs/Minesh-Shah/Dates/2025/12/building-a-lightweight-optimizely
-saas-cms-solution-with-11ty/)  —  *Minesh Shah (2025)*
***Expand to see 67 more resources.***
- [11ty Hacks for Fun and 
Performance](https://infrequently.org/2025/10/11ty-hacks-for-fun-and-performance/)  —  *Alex 
Russell (2025)*
- [From Dotclear to Eleventy 
4](https://alix.guillard.fr/notes/dotclear-to-eleventy/page-navigation/)  —  *Alix Guillard 
(2025)*
- [exhibitionism](https://tlohde.com/blog/2025/09/exhibitionism/)  —  *tlohde (2025)*
- [Category and Tag Pages with 
Eleventy](https://joshtronic.com/2025/09/07/eleventy-category-tag-pages/)  —  *Josh Sherman 
(2025)*
- [Eleventy pagination with Vento](https://gfscott.com/blog/eleventy-pagination-vento/)  —  
*Graham F. Scott (2025)*
- [OG Images using 11ty Screenshot 
Service](https://my.stuffandthings.lol/blog/2025-08-23/og-images-using-11ty-screenshot-service.html)
  —  *Jason Moser (2025)*
- [Building a digital bookshelf with 
Eleventy](https://damianwalsh.co.uk/posts/building-a-digital-bookshelf-with-eleventy/)  —  
*Damian Walsh (2025)*
- [Building a Printful‑powered e‑commerce site with 11ty and 
Stripe](https://www.juanfernandes.uk/blog/building-a-printful%E2%80%91powered-ecommerce-site/) 
 —  *Juan Fernandes (2025)*
- [Double-Pagination in 
Elev­enty](https://chriskirknielsen.com/blog/double-pagination-in-eleventy/)  —  
*Christopher Kirk-Nielsen (2025)*
- [Creating a Journal With 
Eleventy](https://www.simplethread.com/creating-a-journal-with-eleventy/)  —  *Austin Carr 
(2025)*
- [Building a personal digital music library with Eleventy and 
APIs](https://damianwalsh.co.uk/posts/creating-connections-with-music-and-technology/)  —  
*Damian Walsh (2025)*
- [Building a seasonal veg app with Eleventy. Part 
1](https://yasmin.bearblog.dev/building-an-eleventy-app-part-1/)  —  *Yasmin (2025)*
- [Rewriting My Astro Blog with Eleventy](https://melkat.blog/p/11ty-rewrite/)  —  *Melanie 
Kat (2025)*
- [Adding Image Galleries to My Website](https://nathanupchurch.com/blog/galleries/)  —  
*Nathan Upchurch (2024)*
- [Notes on Upgrading to Eleventy 
3.0](https://chromamine.com/2024/11/notes-on-upgrading-to-eleventy-3.0/)  —  *Harris Lapiroff 
(2024)*
- [Oops, I built a headless frontend with 11ty](https://aaadaaam.com/notes/headless-frontend-11ty/) 
 —  *Adam Stoddard (2024)*
- [Eleventy (11ty) year, year-month, and year-monty-day 
indexes](https://blog.tomayac.com/2024/11/02/eleventy-11ty-year-year-month-and-year-month-day-indexe
s/)  —  *Thomas Steiner (2024)*
- [Building a Blog with 
Eleventy](https://blog.sebin-nyshkim.net/posts/building-a-blog-with-eleventy-blind-any/)  —  
*Sebin Nyshkim (2024)*
- [Eleventy Collections from an 
API](https://www.trovster.com/blog/2024/09/eleventy-collections-from-an-api)  —  *Trevor 
Morris (2024)*
- [This Is An Eleventy Blog 
Now!](https://arnaught.neocities.org/blog/2024/08/25/this-is-an-eleventy-blog-now) —  
*Arnaught (2024)*
- [Adding a Photo Stream to an Eleventy 
Site](https://sometimes.digital/posts/adding-a-photostream-to-eleventy/)  —  *nonnullish 
(2024)*
- [Nested pagination with 11ty](https://benwhite.com.au/blog/nested-pagination/)  —  *Ben 
White (2024)*
- [Category and genre pages 
return](https://arestelle.net/2024/06/20/category-and-genre-pages-return/)  —  *Nicki Hoffman 
(2024)*
- [Create a Blog with Eleventy and Storyblok](https://estelafranco.com/blog/eleventy-storyblok-3/) 
 —  *Estela Franco (2024)*
- [Surfing The Web And Sharing What I 
Find](https://flamedfury.com/posts/surfing-the-web-and-sharing-what-i-find/)  —  *fLaMEd 
(2024)*
- [Improving page load times with pagination in 
Eleventy](https://thomasrigby.com/posts/improving-page-load-times-with-pagination-in-eleventy/) 
 —  *Thomas Rigby (2024)*
- [Eleventy Nested Pagination](https://www.codeflood.net/blog/2024/04/17/11ty-nested-pagination/) 
 —  *Alistair Deneys (2024)*
- [Enhancing pagination with a page 
selector](https://www.coryd.dev/posts/2024/enhancing-pagination-with-a-page-selector/)  —  
*Cory Dransfeldt (2024)*
- [Assigning an active page in Eleventy 
navigation](https://anitacheng.com/notes/eleventy-active-nav/)  —  *Anita Cheng (2024)*
- [Low-tech Eleventy Categories](https://piccalil.li/blog/low-tech-eleventy-categories/)  —  
*Andy Bell (2024)*
- [Update: Tags Are Back](https://ttntm.me/blog/restoring-tags/)  —  *Tom Doe (2024)*
- [A Simple Guide to Redirects on Neocities with 
Eleventy](https://flamedfury.com/posts/a-simple-guide-to-redirects-on-neocities-with-eleventy/) 
 —  *fLaMEd (2024)*
- [Migrating from WordPress to Eleventy (part 
3)](https://publishing-project.rivendellweb.net/migrating-from-wordpress-to-eleventy-part-3/) 
 —  *Carlos Araya (2023)*
- [Building post types and category RSS feeds in 
Eleventy](https://localghost.dev/blog/building-post-types-and-category-rss-feeds-in-eleventy/) 
 —  *Sophie Koonin (2023)*
- [Migrating from WordPress to Eleventy (part 
2)](https://publishing-project.rivendellweb.net/migrating-from-wordpress-to-eleventy-part-2/) 
 —  *Carlos Araya (2023)*
- [Using Wordpress as a headless CMS for 
Eleventy](https://www.mikeaparicio.com/posts/2023-11-07-using-wordpress-as-a-headless-cms-for-eleven
ty/)  —  *Mike Aparicio (2023)*
- [Pagination in a Javacsript template with 
Eleventy](https://bobmonsour.com/blog/pagination-in-a-javascript-template-with-eleventy/)  —  
*Bob Monsour (2023)*
- [Eleventy Excellent demo 
branches](https://www.lenesaile.com/en/blog/eleventy-excellent-demo-branches/)  —  *Lene 
Saile (2023)*
- [What I learned making 
top-level.dev](https://ginger.wtf/posts/what-i-learned-making-top-level-dev/)  —  *Ginger 
(2023)*
- [Making Author Pages for an Academic Journal in Eleventy, or, How to Manipulate Collection Data 
in Eleventy](https://micah.torcellini.org/2023/09/23/author-page-manipulate-collections/)  —  
*Micah Torcellini (2023)*
- [11ty: Index ALL the things!](https://lea.verou.me/blog/2023/11ty-indices/) —  *Lea Verou 
(2023)*
- [Fixed Category Page 
Generation](https://johnwargo.com/posts/2023/fixed-category-page-generation/)  —  *John M. 
Wargo (2023)*
- [Generating Eleventy Category Pages Inside Eleventy 
Build](https://johnwargo.com/posts/2023/generating-eleventy-category-pages-inside-eleventy-build/) 
 —  *John M. Wargo (2023)*
- [Eleventy Paginated Category 
Pages](https://johnwargo.com/posts/2023/eleventy-paginated-category-pages/)  —  *John M. 
Wargo (2023)*
- [Lazy select-based pagination in 
Eleventy](https://www.coryd.dev/posts/2023/lazy-select-based-pagination-in-eleventy/)  —  
*Cory Dransfeldt (2023)*
- [From Notion to Eleventy, but faster](https://iamschulz.com/from-notion-to-eleventy-but-faster/) 
 —  *Daniel Schulz (2022)*
- [Creating 11ty Dynamic Categories plugin (with 2-level 
pagination)](https://www.youtube.com/watch?v=DC28C0sGG4w)  —  *Bryan Robinson (2022)*
- [Using Puppeteer with 11ty to automate generating social share 
images](https://photogabble.co.uk/tutorials/using-puppeteer-with-11ty-to-automate-generating-social-
share-images/)  —  *Simon Dann (2022)*
- [Taming Eleventy Tags: Or How I Learned To Tolerate Double 
Pagination](https://desmondrivet.com/2022/03/23/eleventy-pagination)  —  *Desmond Rivet 
(2022)*
- [11ty aliases the right way](https://boehs.org/node/11ty-aliases)  —  *Evan Boehs (2022)*
- [Building blocks for my first Eleventy 
site](https://samimaatta.fi/en/building-blocks-for-my-first-eleventy-site/)  —  *Sami 
Määttä (2022)*
- [Adding QR Codes to Your Jamstack 
Site](https://www.raymondcamden.com/2022/02/11/adding-qr-codes-to-your-jamstack-site)  —  
*Raymond Camden (2022)*
- [Add Pagination to Your Eleventy Static Generated Website in 
Minutes](https://www.thepolyglotdeveloper.com/2022/01/add-pagination-eleventy-static-generated-websi
te-minutes/)  —  *Nic Raboy (2022)*
- [When to Use Pagination in Eleventy](https://shivjm.blog/when-to-use-pagination-in-eleventy/) 
 —  *Shiv J.M. (2021)*
- [Sanity with 11ty: Paginating 
Data](https://www.crossingtheruby.com/2021/05/17/sanity-with-11ty-paginating-data.html)  —  
*crossingtheruby (2021)*
- [Building a Database Driven Eleventy 
Site](https://www.raymondcamden.com/2021/04/15/building-a-database-driven-eleventy-site)  —  
*Raymond Camden (2021)*
- [Using PDFs with the Jamstack - Now with 
Thumbnails](https://www.raymondcamden.com/2021/03/16/using-pdfs-with-the-jamstack-now-with-thumbnail
s)  —  *Raymond Camden (2021)*
- [Adding Simple Pagination to an 11ty 
Collection](https://www.brianperry.dev/til/2021/adding-simple-pagination-to-an-11ty-collection/) 
 —  *Brian Perry (2021)*
- [Using PDFs with the 
Jamstack](https://www.raymondcamden.com/2021/02/25/using-pdfs-with-the-jamstack)  —  *Raymond 
Camden (2021)*
- [Grouping blog posts by year in Eleventy](https://jamesdoc.com/blog/2021/11ty-posts-by-year/) 
 —  *James Doc (2021)*
- [Hooking Up FaunaDB to 
Eleventy](https://www.raymondcamden.com/2020/09/15/hooking-up-faunadb-to-eleventy)  —  
*Raymond Camden (2020)*
- [Eleventy Clock](https://multiline.co/mment/2020/09/eleventy-clock/)  —  *Ashur Cabrera 
(2020)*
- [Supporting Multiple Authors in an Eleventy 
Blog](https://www.raymondcamden.com/2020/08/24/supporting-multiple-authors-in-an-eleventy-blog) 
 —  *Raymond Camden (2020)*
- [Migrating from Node and Express to the Jamstack - Part 
1](https://www.raymondcamden.com/2020/08/06/migrating-from-node-and-express-to-the-jamstack-part-1) 
 —  *Raymond Camden (2020)*
- [Add pagination for dynamic data in 
Eleventy](https://dev.to/gabbersepp/add-pagination-for-dynamic-data-in-eleventy-5fk9)  —  
*Josef Biehler (2020)*
- [Why I'm Digging Eleventy](https://www.raymondcamden.com/2019/10/12/why-im-digging-eleventy) 
 —  *Raymond Camden (2019)*
- [Flexible tag-like functionality for custom keys in 
Eleventy](https://fuzzylogic.me/posts/flexible-tag-like-functionality-for-custom-keys-in-eleventy/) 
 —  *Laurence Hughes (2019)*

---

### Other pages in Create Pages From Data

- [Pagination](https://www.11ty.dev/docs/pagination/)
- [Pagination Navigation](https://www.11ty.dev/docs/pagination/nav/)

---

### Related Docs

- [Quick Tip: Zero Maintenance Tag Pages for your 
Blog](https://www.11ty.dev/docs/quicktips/tag-pages/)
- [Quick Tip: Transform Global Data using an \`eleventyComputed.js\` Global Data 
File](https://www.11ty.dev/docs/quicktips/create-multiple-computed-data-elements/)
- [Create a list of Navigation Links for your 
Pagination.](https://www.11ty.dev/docs/pagination/nav/)
