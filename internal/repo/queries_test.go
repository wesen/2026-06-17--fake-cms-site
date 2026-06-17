package repo

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/go-go-golems/fake-cms/internal/dbcount"
	"github.com/go-go-golems/fake-cms/internal/domain"
)

// freshSeeded opens a fresh in-memory repo and seeds it.
func freshSeeded(t *testing.T) *Repo {
	t.Helper()
	ctx := context.Background()
	r, err := Open(ctx, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Seed(ctx); err != nil {
		t.Fatal(err)
	}
	return r
}

// TestListArticlesPagination verifies keyset pagination + counts.
func TestListArticlesPagination(t *testing.T) {
	ctx := context.Background()
	r := freshSeeded(t)
	defer r.Close()

	page1, err := r.ListArticles(ctx, domain.ArticleFilter{}, 10, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(page1.Nodes) != 10 {
		t.Fatalf("expected 10 nodes, got %d", len(page1.Nodes))
	}
	if !page1.HasNext {
		t.Fatal("expected HasNext on page 1")
	}
	if page1.TotalCount < 100 {
		t.Fatalf("expected >=100 total, got %d", page1.TotalCount)
	}
	// ordering: published_at DESC
	for i := 1; i < len(page1.Nodes); i++ {
		prev := page1.Nodes[i-1].PublishedAt
		cur := page1.Nodes[i].PublishedAt
		if prev != nil && cur != nil && *prev < *cur {
			t.Fatalf("page not DESC at %d", i)
		}
	}

	// page2 follows the cursor.
	page2, err := r.ListArticles(ctx, domain.ArticleFilter{}, 10, page1.EndCursor)
	if err != nil {
		t.Fatal(err)
	}
	if len(page2.Nodes) == 0 {
		t.Fatal("expected page2 nodes")
	}
	// no overlap: last of page1 != first of page2
	if page1.Nodes[len(page1.Nodes)-1].ID == page2.Nodes[0].ID {
		t.Fatal("pages overlap")
	}
}

// TestListArticlesFilters checks each filter narrows results.
func TestListArticlesFilters(t *testing.T) {
	ctx := context.Background()
	r := freshSeeded(t)
	defer r.Close()

	all, _ := r.ListArticles(ctx, domain.ArticleFilter{}, 100, "")

	cases := []domain.ArticleFilter{
		{PostType: domain.PostBestCases},
		{PostType: domain.PostActualites},
		{CategorySlug: "best-cases"},
		{TagSlug: "audience"},
		{Search: "Campagne"},
	}
	for i, f := range cases {
		got, err := r.ListArticles(ctx, f, 100, "")
		if err != nil {
			t.Fatalf("case %d: %v", i, err)
		}
		if got.TotalCount == 0 {
			t.Fatalf("case %d: filter returned 0", i)
		}
		if got.TotalCount > all.TotalCount {
			t.Fatalf("case %d: filter returned more than unfiltered", i)
		}
	}
}

// TestGetArticleBySlug round-trips a slug.
func TestGetArticleBySlug(t *testing.T) {
	ctx := context.Background()
	r := freshSeeded(t)
	defer r.Close()

	page, err := r.ListArticles(ctx, domain.ArticleFilter{}, 1, "")
	if err != nil || len(page.Nodes) == 0 {
		t.Fatal("need at least one article")
	}
	slug := page.Nodes[0].Slug
	got, err := r.GetArticleBySlug(ctx, slug)
	if err != nil {
		t.Fatalf("GetArticleBySlug(%q): %v", slug, err)
	}
	if got.Slug != slug {
		t.Fatalf("slug mismatch: %q vs %q", got.Slug, slug)
	}

	// missing
	if _, err := r.GetArticleBySlug(ctx, "does-not-exist"); !errors.Is(err, sql.ErrNoRows) {
		t.Fatalf("expected ErrNoRows, got %v", err)
	}
}

// TestBatchedLookups verifies the batched relationship loaders return data
// for every requested id in one call each.
func TestBatchedLookups(t *testing.T) {
	ctx := context.Background()
	r := freshSeeded(t)
	defer r.Close()

	page, err := r.ListArticles(ctx, domain.ArticleFilter{}, 20, "")
	if err != nil {
		t.Fatal(err)
	}
	ids := make([]int64, len(page.Nodes))
	for i, n := range page.Nodes {
		ids[i] = n.ID
	}

	authors, err := r.AuthorsByContentIDs(ctx, ids)
	if err != nil {
		t.Fatal(err)
	}
	if len(authors) != len(ids) {
		t.Errorf("authors: expected %d entries, got %d", len(ids), len(authors))
	}

	terms, err := r.TermsByContentIDs(ctx, ids)
	if err != nil {
		t.Fatal(err)
	}
	if len(terms) == 0 {
		t.Error("expected some terms")
	}

	blocks, err := r.BlocksByContentIDs(ctx, ids)
	if err != nil {
		t.Fatal(err)
	}
	if len(blocks) != len(ids) {
		t.Errorf("blocks: expected %d entries, got %d", len(ids), len(blocks))
	}

	media, err := r.FeaturedMediaByContentIDs(ctx, ids)
	if err != nil {
		t.Fatal(err)
	}
	if len(media) != len(ids) {
		t.Errorf("featured media: expected %d, got %d", len(ids), len(media))
	}

	seo, err := r.SEOByContentIDs(ctx, ids)
	if err != nil {
		t.Fatal(err)
	}
	if len(seo) != len(ids) {
		t.Errorf("seo: expected %d, got %d", len(ids), len(seo))
	}
}

// TestNoNPlus1 is the N+1 regression guard. Resolving 20 articles WITH their
// authors/blocks/terms/media/seo must use a fixed, small number of SQL
// statements: 1 (list) + 1 (count) + 5 (one batch per relationship) = 7.
// A naive per-row resolver would do 1 + 1 + 20*N.
func TestNoNPlus1(t *testing.T) {
	ctx := context.Background()
	r := freshSeeded(t)
	defer r.Close()

	counter := dbcount.New(r.DB)
	rc := r.WithQuerier(counter)
	counter.Reset()

	// 1+1: list + count
	page, err := rc.ListArticles(ctx, domain.ArticleFilter{}, 20, "")
	if err != nil {
		t.Fatal(err)
	}
	ids := make([]int64, len(page.Nodes))
	for i, n := range page.Nodes {
		ids[i] = n.ID
	}

	// +5: one batched query per relationship
	if _, err := rc.AuthorsByContentIDs(ctx, ids); err != nil {
		t.Fatal(err)
	}
	if _, err := rc.TermsByContentIDs(ctx, ids); err != nil {
		t.Fatal(err)
	}
	if _, err := rc.BlocksByContentIDs(ctx, ids); err != nil {
		t.Fatal(err)
	}
	if _, err := rc.FeaturedMediaByContentIDs(ctx, ids); err != nil {
		t.Fatal(err)
	}
	if _, err := rc.SEOByContentIDs(ctx, ids); err != nil {
		t.Fatal(err)
	}

	got := counter.Count()
	// list(1) + count(1) + 5 batches = 7. Allow a little slack for pragmas.
	const max = 10
	if got > max {
		t.Fatalf("N+1 regression: resolving 20 articles took %d SQL ops (max %d)", got, max)
	}
	t.Logf("resolved 20 articles + all relationships in %d SQL ops", got)
}
