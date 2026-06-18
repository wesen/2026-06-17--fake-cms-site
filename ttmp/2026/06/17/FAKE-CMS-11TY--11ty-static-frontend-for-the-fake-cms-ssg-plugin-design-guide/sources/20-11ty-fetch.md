---
title: "Eleventy Fetch plugin — Eleventy docs"
doc-type: reference
topics: [11ty, research]
status: active
intent: long-term
external-url: https://www.11ty.dev/docs/plugins/fetch/
summary: "Downloaded with defuddle. Official fetch/caching plugin; useful but optional for CMS query caching. Re-run: defuddle parse https://www.11ty.dev/docs/plugins/fetch/ --md | fold -w 100 -s."
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

*This documentation is written for Eleventy Fetch v5.0+, requiring **Node 18** or newer.*

- [`11ty/eleventy-fetch`](https://github.com/11ty/eleventy-fetch) on GitHub

Fetch network resources and cache them locally to avoid bombarding your API (or other remote or 
rate-limited resources). Do this at configurable intervals (not with every build): once per minute, 
once per hour, once per day, or however often you like!

Successful API requests allow working offline too!

This plugin can save *any* kind of asset—JSON, HTML, images, videos, etc.

- Fetch a remote URL and saves it to a local cache.
- If the remote server goes down or linkrots away—we keep and continue to use the local asset. 
*Save remote images!*
- If a network connection fails (or if you’re offline), Fetch will continue to use the cached 
request (even if the cache is expired) and make a new request when the network connectivity is 
restored.
- Concurrency limits to avoid making too many network requests simultaneously.
- Uses built-in `fetch` Added in Fetch 5.0 (previously 
[`node-fetch`](https://www.npmjs.com/package/node-fetch))
- Supports and caches HTTP verbs separately (think `POST` versus `GET`) Added in Fetch 5.0

## Installation

- [`@11ty/eleventy-fetch` on npm](https://www.npmjs.com/package/@11ty/eleventy-fetch)

```bash
npm install @11ty/eleventy-fetch
```

This plugin was formerly known as 
[`@11ty/eleventy-cache-assets`](https://www.npmjs.com/package/@11ty/eleventy-cache-assets) and was 
renamed to `@11ty/eleventy-fetch` with v3: [Video 
Changelog](https://www.youtube.com/watch?v=JCQQgtOcjH4&t=246s).

WARNING

**Important Security and Privacy Notice**

This plugin caches complete network responses. Unless you’re willing to perform a full review of 
everything this plugin caches to disk for privacy and security exposure, it is *strongly* 
recommended that you add the `.cache` folder to your `.gitignore` file so that network responses 
aren’t checked in to your git repository.

Are you 100% sure that private e-mail addresses aren’t being returned from a cached API? I’m 
guessing no—add `.cache` to your `.gitignore` file. Right now. Do it.

## Usage

### Cache a JSON file from an API

Consider the following example, perhaps in an Eleventy [Global Data 
File](https://www.11ty.dev/docs/data-global/) like `_data/githubRepos.js`.

```js
import Fetch from "@11ty/eleventy-fetch";

export default async function () {
    let url = "https://api.github.com/repos/11ty/eleventy";

    let json = await Fetch(url, {
        duration: "1d", // save for 1 day
        type: "json", // we’ll parse JSON for you
    });

    return json;
};
```

```js
const Fetch = require("@11ty/eleventy-fetch");

module.exports = async function () {
    let url = "https://api.github.com/repos/11ty/eleventy";

    let json = await Fetch(url, {
        duration: "1d", // save for 1 day
        type: "json", // we’ll parse JSON for you
    });

    return json;
};
```

The first argument to Eleventy Fetch can be:

- a string (pointing to a URL)
- a [`URL` instance](https://developer.mozilla.org/en-US/docs/Web/API/URL) Added in Fetch 5.0
- a custom `function` (`async` -friendly) Added in Fetch 5.0

### Options

#### Verbose Output

Option to log requested remote URLs to the console.

- `verbose: true` (default: `false`) Added in Fetch 3.0

#### Change the Cache Duration

After this amount of time has passed, we’ll make a new network request to the URL to fetch fresh 
data. The default duration is `1d` (1 day).

The `duration` option supports the following shorthand values:

- `s` is seconds (e.g. `duration: "43s"`)
- `m` is minutes (e.g. `duration: "2m"`)
- `h` is hours (e.g. `duration: "99h"`)
- `d` is days (The default is `duration: "1d"`)
- `w` is weeks, or shorthand for 7 days (e.g. `duration: 2w` is 14 days)
- `y` is years, or shorthand for 365 days (not *exactly* one year) (e.g. `duration: 2y` is 730 days)

Special values:

- `duration: "*"` will *never* fetch new data (after the first success).
- `duration: "0s"` will *always* fetch new data (works with any unit, e.g. `"0m"`, `"0h"`).

#### Type

- `type: "buffer"` (default)
- `type: "json"`
- `type: "text"`
- `type: "xml"` (alias for `text`) Added in Fetch 5.0
- `type: "parsed-xml"` (uses [`parse-xml`](https://github.com/rgrove/parse-xml) to return a 
JavaScript representation of the XML) Added in Fetch 5.0

#### Return Type Added in Fetch 5.0

- `returnType: undefined` (default) returns the processed body of the request specific to the `type`
- `returnType: "response"` returns a cached object with `url`, `status`, `headers`, and `body` 
properties.

## What happens when a request fails?

1. If this is the first ever request to this URL (no entry exists in your cache folder), it will 
fail. Use a `try` / `catch` if you’d like to handle this gracefully.
2. If a failure happens and a cache entry already exists (*even if it’s expired*), it will use 
the cached entry.
3. If you prefer the build to *fail* when your API requests fail, leave out the `try` `catch` and 
let the error throw without handling it!

Consider the following [global data file](https://www.11ty.dev/docs/data-global/) in Eleventy (e.g. 
`_data/github.js`):

```js
import Fetch from "@11ty/eleventy-fetch";

export default async function () {
    try {
        let url = "https://api.github.com/repos/11ty/eleventy";

        /* This returns a promise */
        return Fetch(url, {
            duration: "1d",
            type: "json",
        });
    } catch (e) {
        return {
            // my failure fallback data
        };
    }
};
```

```js
const Fetch = require("@11ty/eleventy-fetch");

module.exports = async function () {
    try {
        let url = "https://api.github.com/repos/11ty/eleventy";

        /* This returns a promise */
        return Fetch(url, {
            duration: "1d",
            type: "json",
        });
    } catch (e) {
        return {
            // my failure fallback data
        };
    }
};
```

## Running this on your Build Server

This [documentation has moved to the Deployment 
page](https://www.11ty.dev/docs/deployment/#persisting-cache).

## More Examples

### Cache a Remote Image

This is what [`eleventy-img`](https://www.11ty.dev/docs/plugins/image/) uses internally.

```js
import Fetch from "@11ty/eleventy-fetch";

let url = "https://www.zachleat.com/img/avatar-2017-big.png";
let imageBuffer = await Fetch(url, {
    duration: "1d",
    type: "buffer",
});
// Use imageBuffer as an input to the \`sharp\` plugin, for example

// (Example truncated)
```

```js
const Fetch = require("@11ty/eleventy-fetch");

let url = "https://www.zachleat.com/img/avatar-2017-big.png";
let imageBuffer = await Fetch(url, {
    duration: "1d",
    type: "buffer",
});
// Use imageBuffer as an input to the \`sharp\` plugin, for example

// (Example truncated)
```

### Fetch Google Fonts CSS

Also a good example of using `fetchOptions` to pass in a custom user agent. Full option list is 
available on the [`fetch` documentation on 
MDN](https://developer.mozilla.org/en-US/docs/Web/API/Window/fetch#options).

```js
import Fetch from "@11ty/eleventy-fetch";

let url = "https://fonts.googleapis.com/css?family=Roboto+Mono:400&display=swap";
let fontCss = await Fetch(url, {
    duration: "1d",
    type: "text",
    fetchOptions: {
        headers: {
            // lol
            "user-agent":
                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like 
Gecko) Chrome/74.0.3729.169 Safari/537.36",
        },
    },
});
```

```js
const Fetch = require("@11ty/eleventy-fetch");

let url = "https://fonts.googleapis.com/css?family=Roboto+Mono:400&display=swap";
let fontCss = await Fetch(url, {
    duration: "1d",
    type: "text",
    fetchOptions: {
        headers: {
            // lol
            "user-agent":
                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like 
Gecko) Chrome/74.0.3729.169 Safari/537.36",
        },
    },
});
```

### Fetching GitHub Stars for a repo

- This specific example has been previously described in our quick tips section: head over to read 
[Quick Tip #009—Cache Data Requests](https://www.11ty.dev/docs/quicktips/cache-api-requests/).

## Advanced Usage

### Add a custom timeout

Use `AbortSignal` to supply a timeout for the request. Read more about [`AbortSignal` on 
MDN](https://developer.mozilla.org/en-US/docs/Web/API/AbortSignal#aborting_a_fetch_operation_with_a_
timeout).

```js
import Fetch from "@11ty/eleventy-fetch";

await Fetch("https://…", {
    fetchOptions: {
        signal: AbortSignal.timeout(5000)
    },
});
```

```js
const Fetch = require("@11ty/eleventy-fetch");

await Fetch("https://…", {
    fetchOptions: {
        signal: AbortSignal.timeout(5000)
    },
});
```

### Cache Directory

The `directory` option let’s you change where the cache is stored. It is **strongly recommended** 
that you add this folder to your `.gitignore` file.

```js
import Fetch from "@11ty/eleventy-fetch";

await Fetch("https://…", {
    directory: ".cache",
});
```

```js
const Fetch = require("@11ty/eleventy-fetch");

await Fetch("https://…", {
    directory: ".cache",
});
```

If you want to use this utility inside of a Netlify Function (or AWS Lambda), use a writeable 
location (`/tmp/`) like `directory: "/tmp/.cache/"`. You can also use `dryRun: true` to skip 
writing to the file system.

### Remove URL query params from Cache Identifier

Added in Fetch 2.0.3 If your fetched URL contains some query parameters that aren’t relevant to 
the identifier used in the cache, remove them using the `removeUrlQueryParams` option. This is 
useful if an API adds extra junk to your request URLs.

- `removeUrlQueryParams: true` (`false` is default)

```js
import Fetch from "@11ty/eleventy-fetch";

await Fetch(
    "https://www.zachleat.com/img/avatar-2017-big.png?Get=rid&of=these",
    {
        removeUrlQueryParams: true,
    }
);
```

```js
const Fetch = require("@11ty/eleventy-fetch");

await Fetch(
    "https://www.zachleat.com/img/avatar-2017-big.png?Get=rid&of=these",
    {
        removeUrlQueryParams: true,
    }
);
```

Note that query params are removed before—and **are relevant** to how—the cache key is 
calculated.

### Change the cache filename

This controls the name of the files inside your [cache directory](#cache-directory).

```js
import Fetch from "@11ty/eleventy-fetch";

await Fetch("https://…", {
    filenameFormat: function(cacheKey, hash) {
        // do not include the file extension
        return \`custom-name-${cacheKey}-${hash}\`
    }
});
```

```js
const Fetch = require("@11ty/eleventy-fetch");

await Fetch("https://…", {
    filenameFormat: function(cacheKey, hash) {
        // do not include the file extension
        return \`custom-name-${cacheKey}-${hash}\`
    }
});
```

- [`11ty/eleventy-fetch#49`](https://github.com/11ty/eleventy-fetch/pull/49)

### Manually store your own data in the cache

Added in Fetch 5.0 You can pass a function (async-friendly) in as your source to run your own logic 
and return any arbitrary data. You must supply a unique key for the request in the `requestId` 
property. Consider the following [Global Data File](https://www.11ty.dev/docs/data-global/):

```js
import Fetch from "@11ty/eleventy-fetch";

export default function() {
    return Fetch(async function() {
        // do some expensive operation here, this is simplified for brevity
        let fakeTwitterApiContents = { followerCount: 1000 };

        return fakeTwitterApiContents;
    }, {
        // must supply a unique id for the callback
        requestId: "zachleat_twitter_followers"
    });
}
```

```js
const Fetch = require("@11ty/eleventy-fetch");

module.exports = function() {
    return Fetch(async function() {
        // do some expensive operation here, this is simplified for brevity
        let fakeTwitterApiContents = { followerCount: 1000 };

        return fakeTwitterApiContents;
    }, {
        // must supply a unique id for the callback
        requestId: "zachleat_twitter_followers"
    });
}
```

#### Even lower-level access to the cache

**You probably won’t need to do this.** It’s more straightforward to pass in a [function 
source](#manually-store-your-own-data-in-the-cache). Consider the following [Global Data 
File](https://www.11ty.dev/docs/data-global/):

```js
import { AssetCache } from "@11ty/eleventy-fetch";

export default async function () {
    // Pass in your unique custom cache key
    // (normally this would be tied to your API URL)
    let asset = new AssetCache("zachleat_twitter_followers");

    // check if the cache is fresh within the last day
    if (asset.isCacheValid("1d")) {
        // return cached data.
        return asset.getCachedValue(); // a promise
    }

    // do some expensive operation here, this is simplified for brevity
    let fakeTwitterApiContents = { followerCount: 1000 };

    await asset.save(fakeTwitterApiContents, "json");

    return fakeTwitterApiContents;
};
```

```js
const { AssetCache } = require("@11ty/eleventy-fetch");

module.exports = async function () {
    // Pass in your unique custom cache key
    // (normally this would be tied to your API URL)
    let asset = new AssetCache("zachleat_twitter_followers");

    // check if the cache is fresh within the last day
    if (asset.isCacheValid("1d")) {
        // return cached data.
        return asset.getCachedValue(); // a promise
    }

    // do some expensive operation here, this is simplified for brevity
    let fakeTwitterApiContents = { followerCount: 1000 };

    await asset.save(fakeTwitterApiContents, "json");

    return fakeTwitterApiContents;
};
```

### Change Global Concurrency

```js
import Fetch from "@11ty/eleventy-fetch";
Fetch.concurrency = 4; // default is 10
```

### DEBUG mode

```js
DEBUG=Eleventy:Fetch npx @11ty/eleventy
```

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
