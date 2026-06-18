# fake-cms 11ty frontend

Static Eleventy frontend for the fake-cms GraphQL workshop API.

## Development

From the repository root, start the seeded CMS API:

```bash
./fake-cms serve --path testdata/cms.db --addr :8080
```

Then, from this `frontend/` directory:

```bash
npm install
npm run build
npm run serve
```

The build reads the CMS at `CMS_ENDPOINT` (default `http://localhost:8080/graphql`) and writes static files to `_site/`.

Optional environment variables:

- `CMS_ENDPOINT` — GraphQL endpoint to query.
- `SITE_URL` — absolute base URL used by sitemap helpers.
- `CMS_PAGE_SLUGS` — comma-separated page slugs to fetch until the backend exposes page enumeration.

## Validation

Run fast unit tests:

```bash
npm test
```

Run the GraphQL contract smoke against an already-running backend:

```bash
CMS_ENDPOINT=http://localhost:8080/graphql npm run contract:smoke
```

Run the full integration acceptance test. This starts the seeded Go backend itself, builds Eleventy, and inspects `_site/`:

```bash
npm run test:integration
```

Current integration assertions:

- generated article page count equals `articles.totalCount`,
- no `/tag/` output exists,
- `/rubrique/` tag pages exist,
- `sitemap.xml` contains every generated HTML page URL,
- representative article JSON-LD parses as JSON.

## Known backend contract notes

The current executable fake-cms schema is narrower than the checked-in SDL in a few places:

- Start the backend with `--path testdata/cms.db`; the default `cms.db` may be empty.
- There is no `site` query yet, so site metadata lives in Eleventy config.
- There is no `pages` list yet, so page rendering uses optional `CMS_PAGE_SLUGS`.
- `categories`, `tags`, and `authors` currently take no `first` argument.
- Block fields shared by all variants must be queried with `... on Block { id order }`.
