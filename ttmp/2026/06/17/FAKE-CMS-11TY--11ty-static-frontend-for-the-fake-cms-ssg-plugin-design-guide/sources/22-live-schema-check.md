---
title: "Live GraphQL schema check — executable fake-cms server"
doc-type: reference
topics: [graphql, cms, research]
status: active
intent: long-term
summary: "Evidence captured from running ./fake-cms serve --path testdata/cms.db. Highlights differences between checked-in schema.graphql/API docs and executable GraphQL schema."
---

# Live GraphQL schema check

Command context:

```bash
./fake-cms serve --path testdata/cms.db --addr :18080
```

## Query fields and arguments

```graphql
{ __schema { queryType { fields { name args { name type { kind name ofType { kind name } } } type { kind name ofType { kind name } } } } } }
```

```json
{
  "data": {
    "__schema": {
      "queryType": {
        "fields": [
          {
            "args": [
              {
                "name": "id",
                "type": {
                  "kind": "SCALAR",
                  "name": "ID",
                  "ofType": null
                }
              },
              {
                "name": "slug",
                "type": {
                  "kind": "SCALAR",
                  "name": "String",
                  "ofType": null
                }
              }
            ],
            "name": "article",
            "type": {
              "kind": "OBJECT",
              "name": "Article",
              "ofType": null
            }
          },
          {
            "args": [
              {
                "name": "after",
                "type": {
                  "kind": "SCALAR",
                  "name": "String",
                  "ofType": null
                }
              },
              {
                "name": "filter",
                "type": {
                  "kind": "SCALAR",
                  "name": "String",
                  "ofType": null
                }
              },
              {
                "name": "postType",
                "type": {
                  "kind": "SCALAR",
                  "name": "String",
                  "ofType": null
                }
              },
              {
                "name": "first",
                "type": {
                  "kind": "SCALAR",
                  "name": "Int",
                  "ofType": null
                }
              }
            ],
            "name": "articles",
            "type": {
              "kind": "NON_NULL",
              "name": null,
              "ofType": {
                "kind": "OBJECT",
                "name": "ArticleConnection"
              }
            }
          },
          {
            "args": [
              {
                "name": "slug",
                "type": {
                  "kind": "NON_NULL",
                  "name": null,
                  "ofType": {
                    "kind": "SCALAR",
                    "name": "String"
                  }
                }
              }
            ],
            "name": "author",
            "type": {
              "kind": "OBJECT",
              "name": "Author",
              "ofType": null
            }
          },
          {
            "args": [],
            "name": "authors",
            "type": {
              "kind": "LIST",
              "name": null,
              "ofType": {
                "kind": "OBJECT",
                "name": "Author"
              }
            }
          },
          {
            "args": [],
            "name": "categories",
            "type": {
              "kind": "LIST",
              "name": null,
              "ofType": {
                "kind": "OBJECT",
                "name": "Category"
              }
            }
          },
          {
            "args": [
              {
                "name": "slug",
                "type": {
                  "kind": "NON_NULL",
                  "name": null,
                  "ofType": {
                    "kind": "SCALAR",
                    "name": "String"
                  }
                }
              }
            ],
            "name": "category",
            "type": {
              "kind": "OBJECT",
              "name": "Category",
              "ofType": null
            }
          },
          {
            "args": [
              {
                "name": "id",
                "type": {
                  "kind": "NON_NULL",
                  "name": null,
                  "ofType": {
                    "kind": "SCALAR",
                    "name": "ID"
                  }
                }
              }
            ],
            "name": "media",
            "type": {
              "kind": "OBJECT",
              "name": "Media",
              "ofType": null
            }
          },
          {
            "args": [
              {
                "name": "id",
                "type": {
                  "kind": "SCALAR",
                  "name": "ID",
                  "ofType": null
                }
              }
            ],
            "name": "node",
            "type": {
              "kind": "INTERFACE",
              "name": "Node",
              "ofType": null
            }
          },
          {
            "args": [
              {
                "name": "id",
                "type": {
                  "kind": "SCALAR",
                  "name": "ID",
                  "ofType": null
                }
              },
              {
                "name": "slug",
                "type": {
                  "kind": "SCALAR",
                  "name": "String",
                  "ofType": null
                }
              }
            ],
            "name": "page",
            "type": {
              "kind": "OBJECT",
              "name": "Page",
              "ofType": null
            }
          },
          {
            "args": [
              {
                "name": "slug",
                "type": {
                  "kind": "NON_NULL",
                  "name": null,
                  "ofType": {
                    "kind": "SCALAR",
                    "name": "String"
                  }
                }
              }
            ],
            "name": "tag",
            "type": {
              "kind": "OBJECT",
              "name": "Tag",
              "ofType": null
            }
          },
          {
            "args": [],
            "name": "tags",
            "type": {
              "kind": "LIST",
              "name": null,
              "ofType": {
                "kind": "OBJECT",
                "name": "Tag"
              }
            }
          }
        ]
      }
    }
  }
}

```

## SEO fields

```graphql
{ __type(name:"SEO") { fields { name type { kind name ofType { kind name } } } } }
```

```json
{
  "data": {
    "__type": {
      "fields": [
        {
          "name": "breadcrumbs",
          "type": {
            "kind": "LIST",
            "name": null,
            "ofType": {
              "kind": "SCALAR",
              "name": "String"
            }
          }
        },
        {
          "name": "canonical",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        },
        {
          "name": "jsonLd",
          "type": {
            "kind": "SCALAR",
            "name": "JSON",
            "ofType": null
          }
        },
        {
          "name": "metaDescription",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        },
        {
          "name": "og",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        },
        {
          "name": "robots",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        },
        {
          "name": "title",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        },
        {
          "name": "twitter",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        }
      ]
    }
  }
}

```

## Category fields

```graphql
{ __type(name:"Category") { fields { name type { kind name ofType { kind name } } } } }
```

```json
{
  "data": {
    "__type": {
      "fields": [
        {
          "name": "description",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        },
        {
          "name": "id",
          "type": {
            "kind": "SCALAR",
            "name": "ID",
            "ofType": null
          }
        },
        {
          "name": "name",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        },
        {
          "name": "slug",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        }
      ]
    }
  }
}

```

## Tag fields

```graphql
{ __type(name:"Tag") { fields { name type { kind name ofType { kind name } } } } }
```

```json
{
  "data": {
    "__type": {
      "fields": [
        {
          "name": "description",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        },
        {
          "name": "id",
          "type": {
            "kind": "SCALAR",
            "name": "ID",
            "ofType": null
          }
        },
        {
          "name": "name",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        },
        {
          "name": "slug",
          "type": {
            "kind": "SCALAR",
            "name": "String",
            "ofType": null
          }
        }
      ]
    }
  }
}

```

## Smoke query for renderable article

```graphql
{ articles(first:1){ totalCount edges{ node{ slug postType seo{ title jsonLd og twitter breadcrumbs } blocks { __typename ... on Block { id order } ... on ParagraphBlock { text align } ... on HeadingBlock { level text anchor } ... on ImageBlock { media { url alt width height } caption link size } ... on ListBlock { ordered items } ... on QuoteBlock { text citation } ... on EmbedBlock { provider url caption } ... on GalleryBlock { images { url alt width height } columns } } } } } }
```

```json
{
  "summary": {
    "totalCount": 140,
    "slug": "reportage-engagé--oenobiol-72",
    "postType": "ACTUALITES",
    "blockTypes": [
      "GalleryBlock",
      "QuoteBlock",
      "ParagraphBlock",
      "ListBlock",
      "ListBlock"
    ],
    "seoJsonLdRuntimeType": "dict",
    "seoOgRuntimeType": "str"
  }
}

```

## Expected invalid old-guide fields

```graphql
{ site { name } pages { slug } categories(first:20){ slug } articles(filter:{tagSlug:"audience"}){ totalCount } }
```

```json
{
  "data": null,
  "errors": [
    {
      "message": "Cannot query field \"site\" on type \"Query\".",
      "locations": [
        {
          "line": 1,
          "column": 3
        }
      ]
    },
    {
      "message": "Cannot query field \"pages\" on type \"Query\". Did you mean \"page\" or \"tags\"?",
      "locations": [
        {
          "line": 1,
          "column": 17
        }
      ]
    },
    {
      "message": "Unknown argument \"first\" on field \"categories\" of type \"Query\".",
      "locations": [
        {
          "line": 1,
          "column": 43
        }
      ]
    },
    {
      "message": "Argument \"filter\" has invalid value {tagSlug: \"audience\"}.\nExpected type \"String\", found {tagSlug: \"audience\"}.",
      "locations": [
        {
          "line": 1,
          "column": 77
        }
      ]
    }
  ]
}

```

