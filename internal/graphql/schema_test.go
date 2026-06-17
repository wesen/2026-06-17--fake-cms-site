package graphql

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/go-go-golems/fake-cms/internal/repo"
	"github.com/graphql-go/graphql"
)

// seededSchema builds an in-memory repo, seeds it, and returns the schema +
// repo for tests.
func seededSchema(t *testing.T) (graphql.Schema, *repo.Repo) {
	t.Helper()
	ctx := context.Background()
	r, err := repo.Open(ctx, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	if err := r.Seed(ctx); err != nil {
		t.Fatal(err)
	}
	s, err := Schema(r)
	if err != nil {
		t.Fatal(err)
	}
	return s, r
}

func mustExec(t *testing.T, s graphql.Schema, q string, vars map[string]any) map[string]any {
	t.Helper()
	result := graphql.Do(graphql.Params{
		Schema:         s,
		RequestString:  q,
		VariableValues: vars,
		Context:        context.Background(),
	})
	if len(result.Errors) > 0 {
		t.Fatalf("graphql errors: %v", result.Errors)
	}
	out, _ := json.MarshalIndent(result.Data, "", "  ")
	t.Logf("RESULT:\n%s", out)
	return result.Data.(map[string]any)
}

// TestSchemaBuilds verifies the schema compiles with no type errors.
func TestSchemaBuilds(t *testing.T) {
	ctx := context.Background()
	r, err := repo.Open(ctx, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := Schema(r); err != nil {
		t.Fatalf("schema build: %v", err)
	}
}

// TestQueryArticles is the first golden query: list 3 articles with nested
// author + postType + featuredMedia url.
func TestQueryArticles(t *testing.T) {
	s, r := seededSchema(t)
	defer r.Close()
	q := `{
	  articles(first: 3) {
	    totalCount
	    pageInfo { hasNextPage endCursor }
	    edges {
	      cursor
	      node {
	        id slug title postType wordCount
	        publishedAt modifiedAt
	        author { slug displayName }
	        featuredMedia { url width height }
	      }
	    }
	  }
	}`
	data := mustExec(t, s, q, nil)
	conn := data["articles"].(map[string]any)
	if conn["totalCount"].(int) < 100 {
		t.Errorf("expected >=100 total, got %v", conn["totalCount"])
	}
	edges := conn["edges"].([]any)
	if len(edges) != 3 {
		t.Fatalf("expected 3 edges, got %d", len(edges))
	}
	first := edges[0].(map[string]any)["node"].(map[string]any)
	if first["slug"] == "" {
		t.Error("empty slug")
	}
	if first["postType"] == nil {
		t.Error("nil postType")
	}
	pi := conn["pageInfo"].(map[string]any)
	if pi["endCursor"] == nil {
		t.Error("nil endCursor")
	}
}

// TestQueryBlockUnion verifies the union dispatches: query blocks and read a
// per-type field via inline fragments.
func TestQueryBlockUnion(t *testing.T) {
	s, r := seededSchema(t)
	defer r.Close()
	q := `{
	  articles(first: 1) {
	    edges {
	      node {
	        slug
	        blocks {
	          ... on Block { order }
	          ... on HeadingBlock { level text }
	          ... on ParagraphBlock { text }
	          ... on ImageBlock { caption size }
	          ... on ListBlock { ordered items }
	          ... on QuoteBlock { text citation }
	          ... on EmbedBlock { provider url }
	          ... on GalleryBlock { columns }
	        }
	      }
	    }
	  }
	}`
	data := mustExec(t, s, q, nil)
	edges := data["articles"].(map[string]any)["edges"].([]any)
	blocks := edges[0].(map[string]any)["node"].(map[string]any)["blocks"].([]any)
	if len(blocks) == 0 {
		t.Fatal("expected blocks")
	}
	t.Logf("article has %d blocks", len(blocks))
}

// TestQuerySingletonSlug verifies article(slug:) + nested seo + categories.
func TestQuerySingletonSlug(t *testing.T) {
	s, r := seededSchema(t)
	defer r.Close()
	// find a slug first
	list := mustExec(t, s, `{ articles(first:1){ edges { node { slug } } } }`, nil)
	slug := list["articles"].(map[string]any)["edges"].([]any)[0].(map[string]any)["node"].(map[string]any)["slug"].(string)

	q := `query($slug:String!){
	  article(slug:$slug){
	    id slug title excerpt
	    seo { title canonical robots jsonLd }
	    categories { slug name }
	    tags { slug name }
	    author { email }
	  }
	}`
	data := mustExec(t, s, q, map[string]any{"slug": slug})
	art := data["article"].(map[string]any)
	if art["seo"] == nil {
		t.Error("nil seo")
	}
}
