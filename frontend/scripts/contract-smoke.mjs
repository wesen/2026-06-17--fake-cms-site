const endpoint = process.env.CMS_ENDPOINT || "http://localhost:8080/graphql";

async function gql(query, variables = {}) {
  const response = await fetch(endpoint, {
    method: "POST",
    headers: { "content-type": "application/json" },
    body: JSON.stringify({ query, variables }),
  });
  const payload = await response.json();
  return payload;
}

const good = await gql(`{
  articles(first: 1) {
    totalCount
    edges {
      node {
        slug
        postType
        blocks {
          __typename
          ... on Block { id order }
          ... on ParagraphBlock { text align }
        }
      }
    }
  }
}`);

if (good.errors?.length) {
  console.error(JSON.stringify(good.errors, null, 2));
  process.exit(1);
}
if ((good.data?.articles?.totalCount || 0) <= 0) {
  console.error("Expected seeded fake-cms to expose at least one article. Did you run with --path testdata/cms.db?");
  process.exit(1);
}

const bad = await gql(`{ articles(first: 1) { edges { node { blocks { __typename order } } } } }`);
if (!bad.errors?.some(error => /Cannot query field "order" on type "BlockUnion"/.test(error.message))) {
  console.error("Expected direct BlockUnion.order query to fail; executable schema may have changed.");
  console.error(JSON.stringify(bad, null, 2));
  process.exit(1);
}

console.log(`fake-cms contract smoke ok: ${good.data.articles.totalCount} articles at ${endpoint}`);
