const { fetchCms } = require("./fakeCmsClient.cjs");
const { normalizeCms, postTypeSlug, pathSegment } = require("./normalizeCms.cjs");
const { renderBlocks } = require("./renderBlocks.cjs");

function ensureLeadingSlash(path) {
  const value = String(path || "");
  return value.startsWith("/") ? value : `/${value}`;
}

module.exports = function fakeCmsPlugin(eleventyConfig, options = {}) {
  const endpoint = options.endpoint || process.env.CMS_ENDPOINT || "http://localhost:8080/graphql";
  const site = options.site || {
    name: "Fake CMS 11ty Frontend",
    baseUrl: process.env.SITE_URL || "http://localhost:8081",
  };
  const pageSlugs = options.pageSlugs || [];

  eleventyConfig.addGlobalData("cms", async () => {
    const raw = await fetchCms({ endpoint, pageSlugs });
    return normalizeCms(raw, { site });
  });

  eleventyConfig.addFilter("postTypeSlug", postTypeSlug);
  eleventyConfig.addFilter("pathSegment", pathSegment);
  eleventyConfig.addFilter("json", value => JSON.stringify(value, null, 2));
  eleventyConfig.addFilter("absoluteUrl", path => {
    const base = String(site.baseUrl || "").replace(/\/$/, "");
    return `${base}${ensureLeadingSlash(path)}`;
  });

  eleventyConfig.addShortcode("renderBlocks", blocks => renderBlocks(blocks));
};
