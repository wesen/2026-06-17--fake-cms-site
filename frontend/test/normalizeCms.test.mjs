import test from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";

const require = createRequire(import.meta.url);
const { postTypeSlug, pathSegment, normalizeCms } = require("../_config/normalizeCms.cjs");

test("postTypeSlug maps GraphQL enum values to legacy path slugs", () => {
  assert.equal(postTypeSlug("BEST_CASES"), "best-cases");
  assert.equal(postTypeSlug("SLIDER_DE_UNE"), "slider-de-une");
});

test("pathSegment URL-encodes accented slugs", () => {
  assert.equal(pathSegment("actualités"), "actualit%C3%A9s");
});

test("normalizeCms derives URLs, indexes, and sitemap entries", () => {
  const cms = normalizeCms({
    totalArticleCount: 1,
    articles: [{
      slug: "article-é",
      title: "Article",
      postType: "BEST_CASES",
      author: { slug: "admin" },
      tags: [{ slug: "audience", name: "Audience" }],
      categories: [{ slug: "actualites", name: "Actualités" }],
    }],
    categories: [{ slug: "actualites", name: "Actualités" }],
    tags: [{ slug: "audience", name: "Audience" }],
    authors: [{ slug: "admin", displayName: "Admin" }],
    pages: [{ slug: "mentions-legales", title: "Mentions" }],
  });

  assert.equal(cms.articles[0].urlPath, "best-cases/article-%C3%A9");
  assert.equal(cms.tags[0].urlPath, "rubrique/audience");
  assert.equal(cms.categories[0].urlPath, "archives/actualites");
  assert.equal(cms.authors[0].urlPath, "author/admin");
  assert.equal(cms.articlesByTagSlug.audience[0].title, "Article");
  assert.equal(cms.articlesByCategorySlug.actualites[0].title, "Article");
  assert.equal(cms.articlesByAuthorSlug.admin[0].title, "Article");
  assert(cms.sitemapUrls.includes("/rubrique/audience/"));
  assert(!cms.sitemapUrls.some(url => url.startsWith("/tag/")));
});
