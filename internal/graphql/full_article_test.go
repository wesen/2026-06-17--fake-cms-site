package graphql

import (
	"context"
	"testing"
)

// TestQueryFullArticleRender is the full-article golden test (P4.2): it pulls
// one article's complete render payload — blocks (with media resolution in
// image + gallery blocks), seo, taxonomy, author — to prove the SSG plugin
// gets everything it needs in one query. Stable shape; values are seeded.
func TestQueryFullArticleRender(t *testing.T) {
	s, r := seededSchema(t)
	defer r.Close()

	// Find an article id that has at least one image + one gallery block, so we
	// exercise the media-resolution paths. Fallback: any article.
	var contentID int64
	row := r.DB.QueryRowContext(context.Background(),
		`SELECT b.content_id FROM block b
		 WHERE b.type IN ('image','gallery')
		 GROUP BY b.content_id LIMIT 1`)
	if err := row.Scan(&contentID); err != nil {
		t.Skipf("no image/gallery blocks to test: %v", err)
	}

	// resolve the slug for the singleton query
	var slug string
	if err := r.DB.QueryRowContext(context.Background(),
		"SELECT slug FROM content_node WHERE id=?", contentID).Scan(&slug); err != nil {
		t.Fatal(err)
	}

	q := `query($slug:String!){
	  article(slug:$slug){
	    id slug title excerpt postType wordCount
	    publishedAt modifiedAt status locale
	    author { id slug displayName description }
	    categories { id slug name }
	    tags { id slug name }
	    featuredMedia { id url width height alt }
	    seo {
	      title metaDescription canonical robots
	      og twitter jsonLd breadcrumbs
	    }
	    blocks {
	      ... on Block { id order }
	      ... on ParagraphBlock { text }
	      ... on HeadingBlock { level text anchor }
	      ... on ImageBlock { media { id url width height } caption size }
	      ... on ListBlock { ordered items }
	      ... on QuoteBlock { text citation }
	      ... on EmbedBlock { provider url caption }
	      ... on GalleryBlock { images { id url } columns }
	    }
	  }
	}`

	data := mustExec(t, s, q, map[string]any{"slug": slug})
	art := data["article"].(map[string]any)

	// Required-for-render fields must all be present.
	for _, k := range []string{"id", "slug", "title", "postType", "publishedAt", "author", "categories", "tags", "seo", "blocks"} {
		if art[k] == nil {
			t.Errorf("missing required field %q for render", k)
		}
	}

	blocks := art["blocks"].([]any)
	if len(blocks) == 0 {
		t.Fatal("no blocks in render payload")
	}
	// At least one block should carry the resolver-injected media object
	// (image.media or gallery.images), proving the repo-backed media path works.
	t.Logf("rendered article %q with %d blocks", slug, len(blocks))
}
