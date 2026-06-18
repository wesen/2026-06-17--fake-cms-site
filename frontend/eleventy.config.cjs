module.exports = function(eleventyConfig) {
  eleventyConfig.addGlobalData("cms", async () => ({
    site: {
      name: "Fake CMS 11ty Frontend",
      baseUrl: "http://localhost:8081"
    },
    articles: [
      {
        title: "Hardcoded Eleventy smoke article",
        slug: "hardcoded-smoke-article",
        postType: "BEST_CASES",
        postTypeSlug: "best-cases",
        urlPath: "best-cases/hardcoded-smoke-article",
        excerpt: "This article proves one template can emit one CMS-style page.",
        blocks: [{ __typename: "ParagraphBlock", order: 1, text: "Eleventy pagination is generating this page from global data." }],
        seo: { title: "Hardcoded Eleventy smoke article", jsonLd: { "@type": "Article", headline: "Hardcoded Eleventy smoke article" } }
      }
    ]
  }));

  eleventyConfig.addFilter("json", value => JSON.stringify(value, null, 2));
  eleventyConfig.addShortcode("renderBlocks", blocks => (blocks || []).map(block => {
    if (block.__typename === "ParagraphBlock") return `<p>${block.text}</p>`;
    return `<pre>${block.__typename}</pre>`;
  }).join("\n"));

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
