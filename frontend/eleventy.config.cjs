const fakeCmsPlugin = require("./_config/fakeCmsPlugin.cjs");

module.exports = function(eleventyConfig) {
  eleventyConfig.addPlugin(fakeCmsPlugin, {
    endpoint: process.env.CMS_ENDPOINT || "http://localhost:8080/graphql",
    site: {
      name: "Fake CMS 11ty Frontend",
      baseUrl: process.env.SITE_URL || "http://localhost:8081",
    },
    pageSlugs: (process.env.CMS_PAGE_SLUGS || "")
      .split(",")
      .map(slug => slug.trim())
      .filter(Boolean),
  });

  eleventyConfig.addPassthroughCopy({ "src/styles.css": "styles.css" });

  return {
    dir: {
      input: "src",
      includes: "_includes",
      output: "_site"
    },
    templateFormats: ["njk", "css"]
  };
};
