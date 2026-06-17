package graphql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-go-golems/fake-cms/internal/domain"
	"github.com/go-go-golems/fake-cms/internal/repo"
	"github.com/graphql-go/graphql"
)

// articleConn is the connection payload returned by resolveArticles.
type articleConn struct {
	Edges      []edgeData
	TotalCount int
	HasNext    bool
	EndCursor  string
}

// ---- Node -----------------------------------------------------------------

func (r *resolvers) resolveNode(articleT, pageT, authorT, mediaT *graphql.Object) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (any, error) {
		idArg, _ := p.Args["id"].(string)
		typeName, rowID, err := fromGlobalID(idArg)
		if err != nil {
			return nil, err
		}
		ctx := p.Context
		switch typeName {
		case typeArticle, typePage:
			n, err := r.r.GetContentNodeByID(ctx, rowID)
			if err != nil {
				return nil, err
			}
			return n, nil
		case typeAuthor:
			return r.r.AuthorByID(ctx, rowID)
		case typeMedia:
			return r.r.MediaByID(ctx, rowID)
		}
		return nil, fmt.Errorf("unknown node type %q", typeName)
	}
}

// ---- Singletons -----------------------------------------------------------

func (r *resolvers) resolveArticle(p graphql.ResolveParams) (any, error) {
	ctx := p.Context
	if idArg, ok := p.Args["id"].(string); ok && idArg != "" {
		_, rowID, err := fromGlobalID(idArg)
		if err == nil {
			return r.r.GetContentNodeByID(ctx, rowID)
		}
	}
	if slug, ok := p.Args["slug"].(string); ok && slug != "" {
		return r.r.GetArticleBySlug(ctx, slug)
	}
	return nil, nil
}

func (r *resolvers) resolvePage(p graphql.ResolveParams) (any, error) {
	ctx := p.Context
	if idArg, ok := p.Args["id"].(string); ok && idArg != "" {
		_, rowID, err := fromGlobalID(idArg)
		if err == nil {
			n, e := r.r.GetContentNodeByID(ctx, rowID)
			if e == nil && n.Kind == domain.KindPage {
				return n, nil
			}
		}
	}
	if slug, ok := p.Args["slug"].(string); ok && slug != "" {
		return r.r.GetPageBySlug(ctx, slug)
	}
	return nil, nil
}

func (r *resolvers) resolveArticles(p graphql.ResolveParams) (any, error) {
	ctx := p.Context
	first := 20
	if v, ok := p.Args["first"].(int); ok && v > 0 {
		first = v
	}
	after, _ := p.Args["after"].(string)

	f := domain.ArticleFilter{}
	if pt, ok := p.Args["postType"].(string); ok && pt != "" {
		f.PostType = pt
	}
	// The SDL defines `filter: ArticleFilter` as an input object; graphql-go's
	// schema-first object config receives it as a map. We also accept the flat
	// `postType` arg for convenience until a full input type is wired.
	if filt, ok := p.Args["filter"].(map[string]any); ok {
		if pt, ok := filt["postType"].(string); ok {
			f.PostType = pt
		}
		if s, ok := filt["categorySlug"].(string); ok {
			f.CategorySlug = s
		}
		if s, ok := filt["tagSlug"].(string); ok {
			f.TagSlug = s
		}
		if s, ok := filt["authorSlug"].(string); ok {
			f.AuthorSlug = s
		}
		if s, ok := filt["search"].(string); ok {
			f.Search = s
		}
	}

	res, err := r.r.ListArticles(ctx, f, first, after)
	if err != nil {
		return nil, err
	}
	edges := make([]edgeData, len(res.Nodes))
	for i, n := range res.Nodes {
		cursor := ""
		if n.PublishedAt != nil {
			cursor = repo.EncodeCursor(repo.Cursor{PublishedAt: *n.PublishedAt, RowID: n.ID})
		}
		edges[i] = edgeData{Node: n, Cursor: cursor}
	}
	startCur := ""
	if len(edges) > 0 {
		startCur = edges[0].Cursor
	}
	_ = startCur
	return &articleConn{Edges: edges, TotalCount: res.TotalCount, HasNext: res.HasNext, EndCursor: res.EndCursor}, nil
}

// ---- taxonomy singletons --------------------------------------------------

func (r *resolvers) resolveCategory(p graphql.ResolveParams) (any, error) {
	slug, _ := p.Args["slug"].(string)
	return r.r.TermBySlug(p.Context, "category", slug)
}
func (r *resolvers) resolveTag(p graphql.ResolveParams) (any, error) {
	slug, _ := p.Args["slug"].(string)
	return r.r.TermBySlug(p.Context, "post_tag", slug)
}
func (r *resolvers) resolveAuthorBySlug(p graphql.ResolveParams) (any, error) {
	slug, _ := p.Args["slug"].(string)
	return r.r.AuthorBySlug(p.Context, slug)
}
func (r *resolvers) resolveMedia(p graphql.ResolveParams) (any, error) {
	idArg, _ := p.Args["id"].(string)
	_, rowID, err := fromGlobalID(idArg)
	if err != nil {
		return nil, err
	}
	return r.r.MediaByID(p.Context, rowID)
}

// ---- list resolvers (taxonomy top-level) ----------------------------------

func (r *resolvers) resolveAllCategories(p graphql.ResolveParams) (any, error) {
	return r.listTerms(p.Context, "category")
}
func (r *resolvers) resolveAllTags(p graphql.ResolveParams) (any, error) {
	return r.listTerms(p.Context, "post_tag")
}

func (r *resolvers) listTerms(ctx context.Context, taxonomy string) ([]*domain.Term, error) {
	// Simple scan over the term table filtered by taxonomy.
	rows, err := r.r.DB.QueryContext(ctx,
		`SELECT id, taxonomy_slug, slug, name, description, parent_term_id
		FROM term WHERE taxonomy_slug = ? ORDER BY slug`, taxonomy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []*domain.Term
	for rows.Next() {
		t := &domain.Term{}
		if err := rows.Scan(&t.ID, &t.TaxonomySlug, &t.Slug, &t.Name, &t.Description, &t.ParentTermID); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, rows.Err()
}

func (r *resolvers) resolveAllAuthors(p graphql.ResolveParams) (any, error) {
	rows, err := r.r.DB.QueryContext(p.Context,
		`SELECT id, slug, display_name, email, description, avatar_media_id, locale
		FROM author ORDER BY display_name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []*domain.Author
	for rows.Next() {
		a := &domain.Author{}
		if err := rows.Scan(&a.ID, &a.Slug, &a.DisplayName, &a.Email, &a.Description, &a.AvatarMediaID, &a.Locale); err != nil {
			return nil, err
		}
		out = append(out, a)
	}
	return out, rows.Err()
}

// ---- relationship field resolvers ----------------------------------------

func (r *resolvers) resolveAuthor(p graphql.ResolveParams) (any, error) {
	n, ok := p.Source.(*domain.ContentNode)
	if !ok || n.AuthorID == nil {
		return nil, nil
	}
	return r.r.AuthorByID(p.Context, *n.AuthorID)
}

func (r *resolvers) resolveCategories(p graphql.ResolveParams) (any, error) {
	n, ok := p.Source.(*domain.ContentNode)
	if !ok {
		return nil, nil
	}
	terms, err := r.r.TermsByContentID(p.Context, n.ID)
	if err != nil {
		return nil, err
	}
	out := make([]*domain.Term, 0, len(terms))
	for _, t := range terms {
		if t.TaxonomySlug == "category" {
			out = append(out, t)
		}
	}
	return out, nil
}

func (r *resolvers) resolveTags(p graphql.ResolveParams) (any, error) {
	n, ok := p.Source.(*domain.ContentNode)
	if !ok {
		return nil, nil
	}
	terms, err := r.r.TermsByContentID(p.Context, n.ID)
	if err != nil {
		return nil, err
	}
	out := make([]*domain.Term, 0, len(terms))
	for _, t := range terms {
		if t.TaxonomySlug == "post_tag" {
			out = append(out, t)
		}
	}
	return out, nil
}

func (r *resolvers) resolveFeaturedMedia(p graphql.ResolveParams) (any, error) {
	n, ok := p.Source.(*domain.ContentNode)
	if !ok {
		return nil, nil
	}
	media, err := r.r.FeaturedMediaByContentIDs(p.Context, []int64{n.ID})
	if err != nil {
		return nil, err
	}
	return media[n.ID], nil
}

func (r *resolvers) resolveBlocks(p graphql.ResolveParams) (any, error) {
	n, ok := p.Source.(*domain.ContentNode)
	if !ok {
		return nil, nil
	}
	blocks, err := r.r.BlocksByContentID(p.Context, n.ID)
	if err != nil {
		return nil, err
	}
	out := make([]blockValue, len(blocks))
	for i, b := range blocks {
		out[i] = toBlockValue(b)
	}
	return out, nil
}

func (r *resolvers) resolvePostType(p graphql.ResolveParams) (any, error) {
	n, ok := p.Source.(*domain.ContentNode)
	if !ok {
		return nil, nil
	}
	wpSlug, err := r.r.GetArticlePostType(p.Context, n.ID)
	if err != nil || wpSlug == "" {
		return nil, nil
	}
	return wpToGraphqlEnum(wpSlug), nil
}

func (r *resolvers) resolveAuthorAvatar(p graphql.ResolveParams) (any, error) {
	a, ok := p.Source.(*domain.Author)
	if !ok || a.AvatarMediaID == nil {
		return nil, nil
	}
	return r.r.MediaByID(p.Context, *a.AvatarMediaID)
}

// ---- SEO resolvers --------------------------------------------------------

func (r *resolvers) resolveSEO(p graphql.ResolveParams) (any, error) {
	n, ok := p.Source.(*domain.ContentNode)
	if !ok {
		return nil, nil
	}
	s, err := r.r.SEOByContentID(p.Context, n.ID)
	if err == sql.ErrNoRows {
		return &domain.SEO{ContentID: n.ID, Title: n.Title}, nil
	}
	return s, err
}
func (r *resolvers) resolveSEOOG(p graphql.ResolveParams) (any, error) {
	s, ok := p.Source.(*domain.SEO)
	if !ok {
		return nil, nil
	}
	return decodeJSON(s.OGJSON), nil
}
func (r *resolvers) resolveSEOTwitter(p graphql.ResolveParams) (any, error) {
	s, ok := p.Source.(*domain.SEO)
	if !ok {
		return nil, nil
	}
	return decodeJSON(s.TwitterJSON), nil
}
func (r *resolvers) resolveSEOJSONLD(p graphql.ResolveParams) (any, error) {
	s, ok := p.Source.(*domain.SEO)
	if !ok {
		return nil, nil
	}
	return decodeJSON(s.JSONLD), nil
}
func (r *resolvers) resolveSEOBreadcrumbs(p graphql.ResolveParams) (any, error) {
	s, ok := p.Source.(*domain.SEO)
	if !ok {
		return nil, nil
	}
	return decodeJSON(s.BreadcrumbsJSON), nil
}

// ---- enum mapping ---------------------------------------------------------

func wpToGraphqlEnum(wpSlug string) string {
	switch wpSlug {
	case "actualites":
		return domain.PostActualites
	case "best-cases":
		return domain.PostBestCases
	case "etudes":
		return domain.PostEtudes
	case "blog":
		return domain.PostBlog
	case "slider-de-une":
		return domain.PostSliderDeUne
	case "cartouches-home":
		return domain.PostCartouchesHom
	case "non-classe":
		return domain.PostNonClasse
	}
	return ""
}
