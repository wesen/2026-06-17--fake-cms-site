const DEFAULT_POST_TYPES = [
  "ACTUALITES",
  "BEST_CASES",
  "ETUDES",
  "BLOG",
  "SLIDER_DE_UNE",
  "CARTOUCHES_HOME",
  "NON_CLASSE",
];

function postTypeSlug(postType) {
  return String(postType || "")
    .toLowerCase()
    .replaceAll("_", "-");
}

function pathSegment(slug) {
  return encodeURIComponent(String(slug || ""));
}

function indexByMany(items, getKeys) {
  const out = Object.create(null);
  for (const item of items) {
    for (const key of getKeys(item) || []) {
      if (!key) continue;
      (out[key] ||= []).push(item);
    }
  }
  return out;
}

function withUrl(entity, kind, urlPath) {
  return { ...entity, kind, urlPath };
}

function normalizeCms(raw, options = {}) {
  const site = options.site || { name: "Fake CMS", baseUrl: "http://localhost:8081" };
  const postTypes = options.postTypes || DEFAULT_POST_TYPES;

  const articles = (raw.articles || []).map(article => {
    const typeSlug = postTypeSlug(article.postType);
    return {
      ...article,
      kind: "article",
      postTypeSlug: typeSlug,
      urlPath: `${typeSlug}/${pathSegment(article.slug)}`,
    };
  });
  const pages = (raw.pages || []).map(page => withUrl(page, "page", pathSegment(page.slug)));
  const categories = (raw.categories || []).map(category => withUrl(category, "category", `archives/${pathSegment(category.slug)}`));
  const tags = (raw.tags || []).map(tag => withUrl(tag, "tag", `rubrique/${pathSegment(tag.slug)}`));
  const authors = (raw.authors || []).map(author => withUrl(author, "author", `author/${pathSegment(author.slug)}`));

  const articlesByPostType = indexByMany(articles, article => [article.postType]);
  const articlesByTagSlug = indexByMany(articles, article => (article.tags || []).map(tag => tag.slug));
  const articlesByCategorySlug = indexByMany(articles, article => (article.categories || []).map(category => category.slug));
  const articlesByAuthorSlug = indexByMany(articles, article => [article.author?.slug]);

  const sitemapUrls = [
    ...articles.map(article => `/${article.urlPath}/`),
    ...pages.map(page => `/${page.urlPath}/`),
    ...categories.map(category => `/${category.urlPath}/`),
    ...tags.map(tag => `/${tag.urlPath}/`),
    ...authors.map(author => `/${author.urlPath}/`),
  ];

  return {
    site,
    postTypes,
    totalArticleCount: raw.totalArticleCount ?? articles.length,
    articles,
    pages,
    categories,
    tags,
    authors,
    articlesByPostType,
    articlesByTagSlug,
    articlesByCategorySlug,
    articlesByAuthorSlug,
    sitemapUrls,
  };
}

module.exports = {
  DEFAULT_POST_TYPES,
  postTypeSlug,
  pathSegment,
  indexByMany,
  normalizeCms,
};
