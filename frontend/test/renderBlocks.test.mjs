import test from "node:test";
import assert from "node:assert/strict";
import { createRequire } from "node:module";

const require = createRequire(import.meta.url);
const { escapeHtml, clampHeadingLevel, renderBlock, renderBlocks } = require("../_config/renderBlocks.cjs");

test("escapeHtml escapes text and attributes", () => {
  assert.equal(escapeHtml(`<script>alert("x")</script>`), "&lt;script&gt;alert(&quot;x&quot;)&lt;/script&gt;");
});

test("clampHeadingLevel keeps headings semantic and safe", () => {
  assert.equal(clampHeadingLevel(1), 2);
  assert.equal(clampHeadingLevel(4), 4);
  assert.equal(clampHeadingLevel(99), 6);
  assert.equal(clampHeadingLevel("nope"), 2);
});

test("renderBlock renders all seven CMS block variants distinctly", () => {
  const blocks = [
    { __typename: "ParagraphBlock", order: 1, text: "Hello", align: "CENTER" },
    { __typename: "HeadingBlock", order: 2, level: 3, text: "Head", anchor: "head" },
    { __typename: "ImageBlock", order: 3, media: { url: "https://example.test/a.jpg", alt: "A", width: 10, height: 20 }, caption: "Cap" },
    { __typename: "ListBlock", order: 4, ordered: true, items: ["One", "Two"] },
    { __typename: "QuoteBlock", order: 5, text: "Quote", citation: "Me" },
    { __typename: "EmbedBlock", order: 6, provider: "youtube", url: "https://youtu.be/example", caption: "Video" },
    { __typename: "GalleryBlock", order: 7, columns: 2, images: [{ url: "https://example.test/g.jpg", alt: "G" }] },
  ];

  const html = blocks.map(renderBlock);
  assert.equal(new Set(html.map(fragment => fragment.match(/^<[^\s>]+/)?.[0])).size, 6, "six distinct root tags/classes expected with figure reused for image/embed");
  assert.match(html[0], /<p class="align-center">Hello<\/p>/);
  assert.match(html[1], /<h3 id="head">Head<\/h3>/);
  assert.match(html[2], /block-image/);
  assert.match(html[3], /<ol class="block-list">/);
  assert.match(html[4], /block-quote/);
  assert.match(html[5], /block-embed-youtube/);
  assert.match(html[6], /block-gallery columns-2/);
});

test("renderBlocks sorts by order before rendering", () => {
  const html = renderBlocks([
    { __typename: "ParagraphBlock", order: 2, text: "Second" },
    { __typename: "ParagraphBlock", order: 1, text: "First" },
  ]);
  assert(html.indexOf("First") < html.indexOf("Second"));
});

test("renderBlock throws on unknown CMS block types", () => {
  assert.throws(() => renderBlock({ __typename: "NewFancyBlock" }), /Unknown CMS block type/);
});
