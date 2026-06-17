package repo

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/go-go-golems/fake-cms/internal/domain"
)

// Cursor encodes the position of a row in the published_at-DESC ordering used
// by ListArticles. Format: base64("<iso8601>:<rowid>"). Rowid breaks ties so
// pagination is stable even when two articles share a published_at.
type Cursor struct {
	PublishedAt string
	RowID       int64
}

// EncodeCursor produces the opaque client-facing cursor.
func EncodeCursor(c Cursor) string {
	raw := fmt.Sprintf("%s:%d", c.PublishedAt, c.RowID)
	return base64.URLEncoding.EncodeToString([]byte(raw))
}

// DecodeCursor parses a client cursor. Returns ok=false on malformed input
// (the caller treats that as "start from the beginning").
func DecodeCursor(s string) (Cursor, bool) {
	if s == "" {
		return Cursor{}, false
	}
	raw, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return Cursor{}, false
	}
	parts := strings.SplitN(string(raw), ":", 2)
	if len(parts) != 2 {
		return Cursor{}, false
	}
	var id int64
	if _, err := fmt.Sscanf(parts[1], "%d", &id); err != nil {
		return Cursor{}, false
	}
	return Cursor{PublishedAt: parts[0], RowID: id}, true
}

// ListArticlesResult holds a page plus the total count.
type ListArticlesResult struct {
	Nodes      []*domain.ContentNode
	TotalCount int
	HasNext    bool
	EndCursor  string
}

// ListArticles returns a page of ARTICLE content nodes ordered by published_at
// DESC (then id DESC), filtered by the given filter. `after` is an opaque
// cursor; pagination is keyset, stable under inserts before the cursor.
//
// It issues exactly ONE query for rows + ONE for the count (2 queries total),
// so callers batching related lookups (AuthorsByContentIDs, TermsByContentID)
// keep the whole list operation at a fixed small query count.
func (r *Repo) ListArticles(ctx context.Context, f domain.ArticleFilter, first int, after string) (*ListArticlesResult, error) {
	if first <= 0 || first > 100 {
		first = 20
	}

	var (
		where []string
		args  []any
	)
	// status defaults to PUBLISH unless overridden
	status := f.Status
	if status == "" {
		status = domain.StatusPublish
	}
	where = append(where, "n.kind = 'ARTICLE'")
	where = append(where, "n.status = ?")
	args = append(args, status)

	// optional joins for taxonomy/author filters; built lazily.
	joins := ""
	if f.PostType != "" {
		wpSlug := postTypeGraphqlToWP(f.PostType)
		if wpSlug != "" {
			joins += " JOIN article_post_type apt ON apt.content_id = n.id"
			where = append(where, "apt.post_type = ?")
			args = append(args, wpSlug)
		}
	}
	if f.AuthorSlug != "" {
		joins += " JOIN author a ON a.id = n.author_id"
		where = append(where, "a.slug = ?")
		args = append(args, f.AuthorSlug)
	}
	if f.CategorySlug != "" {
		joins += " JOIN content_term ct_cat ON ct_cat.content_id = n.id"
		joins += " JOIN term t_cat ON t_cat.id = ct_cat.term_id"
		where = append(where, "t_cat.taxonomy_slug = 'category' AND t_cat.slug = ?")
		args = append(args, f.CategorySlug)
	}
	if f.TagSlug != "" {
		joins += " JOIN content_term ct_tag ON ct_tag.content_id = n.id"
		joins += " JOIN term t_tag ON t_tag.id = ct_tag.term_id"
		where = append(where, "t_tag.taxonomy_slug = 'post_tag' AND t_tag.slug = ?")
		args = append(args, f.TagSlug)
	}
	if f.Search != "" {
		where = append(where, "(n.title LIKE ? OR n.excerpt LIKE ?)")
		args = append(args, "%"+f.Search+"%", "%"+f.Search+"%")
	}
	if f.PublishedAfter != nil {
		where = append(where, "n.published_at >= ?")
		args = append(args, *f.PublishedAfter)
	}
	if f.PublishedBefore != nil {
		where = append(where, "n.published_at <= ?")
		args = append(args, *f.PublishedBefore)
	}

	// keyset cursor: published_at < cur.PublishedAt OR (== and id < cur.RowID)
	if cur, ok := DecodeCursor(after); ok {
		where = append(where, "(n.published_at < ? OR (n.published_at = ? AND n.id < ?))")
		args = append(args, cur.PublishedAt, cur.PublishedAt, cur.RowID)
	}

	whereSQL := strings.Join(where, " AND ")

	// fetch first+1 to learn HasNext without a separate query
	q := fmt.Sprintf(`SELECT n.id, n.kind, n.slug, n.title, n.excerpt, n.status,
		n.locale, n.published_at, n.modified_at, n.author_id, n.parent_page_id,
		n.menu_order, n.template, n.word_count
		FROM content_node n%s
		WHERE %s
		ORDER BY n.published_at DESC, n.id DESC
		LIMIT ?`, joins, whereSQL)
	args = append(args, first+1)

	rows, err := r.Q.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, fmt.Errorf("list articles query: %w", err)
	}
	defer rows.Close()

	nodes := make([]*domain.ContentNode, 0, first)
	for rows.Next() {
		n := &domain.ContentNode{}
		if err := scanContentNode(rows.Scan, n); err != nil {
			return nil, err
		}
		nodes = append(nodes, n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	hasNext := len(nodes) > first
	if hasNext {
		nodes = nodes[:first]
	}
	endCursor := ""
	if len(nodes) > 0 {
		last := nodes[len(nodes)-1]
		if last.PublishedAt != nil {
			endCursor = EncodeCursor(Cursor{PublishedAt: *last.PublishedAt, RowID: last.ID})
		}
	}

	// total count (separate, cheap COUNT)
	countQ := fmt.Sprintf(`SELECT count(*) FROM content_node n%s WHERE %s`, joins, whereSQL)
	// args for count exclude the LIMIT; rebuild without the trailing first+1.
	countArgs := args[:len(args)-1]
	var total int
	if err := r.Q.QueryRowContext(ctx, countQ, countArgs...).Scan(&total); err != nil {
		return nil, fmt.Errorf("count articles: %w", err)
	}

	return &ListArticlesResult{
		Nodes: nodes, TotalCount: total, HasNext: hasNext, EndCursor: endCursor,
	}, nil
}

// GetArticleBySlug returns a single article by its slug, or sql.ErrNoRows.
func (r *Repo) GetArticleBySlug(ctx context.Context, slug string) (*domain.ContentNode, error) {
	return r.getContentNode(ctx, "ARTICLE", slug)
}

// GetPageBySlug returns a single page by its slug, or sql.ErrNoRows.
func (r *Repo) GetPageBySlug(ctx context.Context, slug string) (*domain.ContentNode, error) {
	return r.getContentNode(ctx, "PAGE", slug)
}

func (r *Repo) getContentNode(ctx context.Context, kind, slug string) (*domain.ContentNode, error) {
	n := &domain.ContentNode{}
	q := `SELECT id, kind, slug, title, excerpt, status, locale, published_at,
		modified_at, author_id, parent_page_id, menu_order, template, word_count
		FROM content_node WHERE kind = ? AND slug = ?`
	row := r.Q.QueryRowContext(ctx, q, kind, slug)
	if err := scanContentNode(row.Scan, n); err != nil {
		return nil, err
	}
	return n, nil
}

// GetContentNodeByID returns a content node by id regardless of kind.
func (r *Repo) GetContentNodeByID(ctx context.Context, id int64) (*domain.ContentNode, error) {
	n := &domain.ContentNode{}
	q := `SELECT id, kind, slug, title, excerpt, status, locale, published_at,
		modified_at, author_id, parent_page_id, menu_order, template, word_count
		FROM content_node WHERE id = ?`
	row := r.Q.QueryRowContext(ctx, q, id)
	if err := scanContentNode(row.Scan, n); err != nil {
		return nil, err
	}
	return n, nil
}

// scanContentNode adapts the row.Scan func signature for reuse.
type scanFn func(dest ...any) error

func scanContentNode(scan scanFn, n *domain.ContentNode) error {
	var (
		excerpt   sql.NullString
		published sql.NullString
		authorID  sql.NullInt64
		parentID  sql.NullInt64
		template  sql.NullString
		wordCount sql.NullInt64
	)
	if err := scan(
		&n.ID, &n.Kind, &n.Slug, &n.Title, &excerpt, &n.Status, &n.Locale,
		&published, &n.ModifiedAt, &authorID, &parentID,
		&n.MenuOrder, &template, &wordCount,
	); err != nil {
		return err
	}
	n.Excerpt = excerpt.String
	if published.Valid {
		s := published.String
		n.PublishedAt = &s
	}
	if authorID.Valid {
		id := authorID.Int64
		n.AuthorID = &id
	}
	if parentID.Valid {
		id := parentID.Int64
		n.ParentPageID = &id
	}
	n.Template = template.String
	if wordCount.Valid {
		wc := wordCount.Int64
		n.WordCount = &wc
	}
	return nil
}

// postTypeGraphqlToWP maps the graphql PostType enum to the WP slug stored in
// article_post_type.post_type.
func postTypeGraphqlToWP(enum string) string {
	switch enum {
	case domain.PostActualites:
		return "actualites"
	case domain.PostBestCases:
		return "best-cases"
	case domain.PostEtudes:
		return "etudes"
	case domain.PostBlog:
		return "blog"
	case domain.PostSliderDeUne:
		return "slider-de-une"
	case domain.PostCartouchesHom:
		return "cartouches-home"
	case domain.PostNonClasse:
		return "non-classe"
	}
	return ""
}

// GetArticlePostType returns the WP post-type slug for an article, or "".
func (r *Repo) GetArticlePostType(ctx context.Context, contentID int64) (string, error) {
	var s string
	err := r.Q.QueryRowContext(ctx,
		"SELECT post_type FROM article_post_type WHERE content_id = ?", contentID).Scan(&s)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return s, err
}
