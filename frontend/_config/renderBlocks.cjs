function escapeHtml(value) {
  return String(value ?? "").replace(/[&<>"']/g, char => ({
    "&": "&amp;",
    "<": "&lt;",
    ">": "&gt;",
    '"': "&quot;",
    "'": "&#39;",
  })[char]);
}

function attrs(values) {
  return Object.entries(values)
    .filter(([, value]) => value !== undefined && value !== null && value !== "")
    .map(([key, value]) => ` ${key}="${escapeHtml(value)}"`)
    .join("");
}

function clampHeadingLevel(level) {
  const n = Number(level || 2);
  if (!Number.isFinite(n)) return 2;
  return Math.min(6, Math.max(2, Math.trunc(n)));
}

function renderBlocks(blocks = []) {
  return [...blocks]
    .sort((a, b) => (a.order ?? 0) - (b.order ?? 0))
    .map(renderBlock)
    .join("\n");
}

function renderBlock(block) {
  switch (block?.__typename) {
    case "ParagraphBlock":
      return `<p${attrs({ class: block.align ? `align-${String(block.align).toLowerCase()}` : undefined })}>${escapeHtml(block.text)}</p>`;

    case "HeadingBlock": {
      const level = clampHeadingLevel(block.level);
      return `<h${level}${attrs({ id: block.anchor })}>${escapeHtml(block.text)}</h${level}>`;
    }

    case "ImageBlock": {
      const media = block.media || {};
      const img = `<img${attrs({ src: media.url, alt: media.alt || "", width: media.width, height: media.height })}>`;
      const linked = block.link ? `<a href="${escapeHtml(block.link)}">${img}</a>` : img;
      return `<figure class="block-image">${linked}${block.caption ? `<figcaption>${escapeHtml(block.caption)}</figcaption>` : ""}</figure>`;
    }

    case "ListBlock": {
      const tag = block.ordered ? "ol" : "ul";
      return `<${tag} class="block-list">${(block.items || []).map(item => `<li>${escapeHtml(item)}</li>`).join("")}</${tag}>`;
    }

    case "QuoteBlock":
      return `<blockquote class="block-quote"><p>${escapeHtml(block.text)}</p>${block.citation ? `<cite>${escapeHtml(block.citation)}</cite>` : ""}</blockquote>`;

    case "EmbedBlock":
      return `<figure class="block-embed block-embed-${escapeHtml(block.provider || "unknown")}"><a href="${escapeHtml(block.url)}">${escapeHtml(block.url)}</a>${block.caption ? `<figcaption>${escapeHtml(block.caption)}</figcaption>` : ""}</figure>`;

    case "GalleryBlock":
      return `<ul class="block-gallery columns-${escapeHtml(block.columns || 3)}">${(block.images || []).map(image => `<li><img${attrs({ src: image.url, alt: image.alt || "", width: image.width, height: image.height })}></li>`).join("")}</ul>`;

    default:
      throw new Error(`Unknown CMS block type: ${block?.__typename}`);
  }
}

module.exports = {
  escapeHtml,
  attrs,
  clampHeadingLevel,
  renderBlocks,
  renderBlock,
};
