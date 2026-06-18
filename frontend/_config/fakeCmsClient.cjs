const ARTICLE_FIELDS = `
  id slug title excerpt postType publishedAt modifiedAt wordCount
  author { id slug displayName description }
  categories { id slug name description }
  tags { id slug name description }
  featuredMedia { id url alt width height caption }
  seo { title metaDescription canonical robots jsonLd og twitter breadcrumbs }
  blocks {
    __typename
    ... on Block { id order }
    ... on ParagraphBlock { text align }
    ... on HeadingBlock { level text anchor }
    ... on ImageBlock { media { id url alt width height } caption link size }
    ... on ListBlock { ordered items }
    ... on QuoteBlock { text citation }
    ... on EmbedBlock { provider url caption }
    ... on GalleryBlock { images { id url alt width height } columns }
  }
`;

const PAGE_FIELDS = `
  id slug title status publishedAt modifiedAt template
  featuredMedia { id url alt width height caption }
  seo { title metaDescription canonical robots jsonLd og twitter breadcrumbs }
  blocks {
    __typename
    ... on Block { id order }
    ... on ParagraphBlock { text align }
    ... on HeadingBlock { level text anchor }
    ... on ImageBlock { media { id url alt width height } caption link size }
    ... on ListBlock { ordered items }
    ... on QuoteBlock { text citation }
    ... on EmbedBlock { provider url caption }
    ... on GalleryBlock { images { id url alt width height } columns }
  }
`;

async function gql(endpoint, query, variables = {}) {
  const response = await fetch(endpoint, {
    method: "POST",
    headers: { "content-type": "application/json" },
    body: JSON.stringify({ query, variables }),
  });
  if (!response.ok) {
    throw new Error(`GraphQL HTTP ${response.status}: ${await response.text()}`);
  }
  const payload = await response.json();
  if (payload.errors?.length) {
    throw new Error(`GraphQL errors:\n${JSON.stringify(payload.errors, null, 2)}`);
  }
  return payload.data;
}

async function fetchAllArticles(endpoint, { first = 50 } = {}) {
  const articles = [];
  let after = null;
  let totalCount = null;

  for (;;) {
    const data = await gql(endpoint, `
      query Articles($first: Int!, $after: String) {
        articles(first: $first, after: $after) {
          totalCount
          edges { cursor node { ${ARTICLE_FIELDS} } }
          pageInfo { hasNextPage endCursor }
        }
      }
    `, { first, after });

    totalCount = data.articles.totalCount;
    for (const edge of data.articles.edges || []) {
      articles.push(edge.node);
    }
    if (!data.articles.pageInfo.hasNextPage) break;
    after = data.articles.pageInfo.endCursor;
    if (!after) throw new Error("CMS pagination reported hasNextPage=true but endCursor was empty");
  }

  return { articles, totalCount };
}

async function fetchKnownPages(endpoint, pageSlugs = []) {
  const pages = [];
  for (const slug of pageSlugs) {
    const data = await gql(endpoint, `
      query Page($slug: String) {
        page(slug: $slug) { ${PAGE_FIELDS} }
      }
    `, { slug });
    if (data.page) pages.push(data.page);
  }
  return pages;
}

async function fetchCms({ endpoint, pageSlugs = [] }) {
  const [articleResult, categoriesData, tagsData, authorsData, pages] = await Promise.all([
    fetchAllArticles(endpoint),
    gql(endpoint, `{ categories { id slug name description } }`),
    gql(endpoint, `{ tags { id slug name description } }`),
    gql(endpoint, `{ authors { id slug displayName email description avatar { url alt } } }`),
    fetchKnownPages(endpoint, pageSlugs),
  ]);

  return {
    articles: articleResult.articles,
    totalArticleCount: articleResult.totalCount,
    categories: categoriesData.categories || [],
    tags: tagsData.tags || [],
    authors: authorsData.authors || [],
    pages,
  };
}

module.exports = {
  ARTICLE_FIELDS,
  PAGE_FIELDS,
  gql,
  fetchAllArticles,
  fetchKnownPages,
  fetchCms,
};
