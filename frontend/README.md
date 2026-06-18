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

The initial scaffold uses hardcoded global data to prove the Eleventy pagination pattern. Later phases replace that data with the GraphQL client and normalizer.
