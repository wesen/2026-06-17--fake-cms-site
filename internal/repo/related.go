package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-go-golems/fake-cms/internal/domain"
)

// placeholders returns "?,?,?,?" for n args.
func placeholders(n int) string {
	if n <= 0 {
		// SQLite tolerates () but no IN () ; use a never-match sentinel.
		return "NULL"
	}
	return strings.Repeat("?,", n-1) + "?"
}

// ---- Authors ---------------------------------------------------------------

// AuthorByID returns a single author or sql.ErrNoRows.
func (r *Repo) AuthorByID(ctx context.Context, id int64) (*domain.Author, error) {
	a := &domain.Author{}
	err := r.Q.QueryRowContext(ctx,
		`SELECT id, slug, display_name, email, description, avatar_media_id, locale
		FROM author WHERE id = ?`, id).
		Scan(&a.ID, &a.Slug, &a.DisplayName, &a.Email, &a.Description, &a.AvatarMediaID, &a.Locale)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// AuthorBySlug returns a single author by slug or sql.ErrNoRows.
func (r *Repo) AuthorBySlug(ctx context.Context, slug string) (*domain.Author, error) {
	a := &domain.Author{}
	err := r.Q.QueryRowContext(ctx,
		`SELECT id, slug, display_name, email, description, avatar_media_id, locale
		FROM author WHERE slug = ?`, slug).
		Scan(&a.ID, &a.Slug, &a.DisplayName, &a.Email, &a.Description, &a.AvatarMediaID, &a.Locale)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// AuthorsByContentIDs batches author lookup for many content nodes in ONE
// query. Returns contentID -> *Author. This is the N+1 mitigation primitive.
func (r *Repo) AuthorsByContentIDs(ctx context.Context, ids []int64) (map[int64]*domain.Author, error) {
	out := map[int64]*domain.Author{}
	if len(ids) == 0 {
		return out, nil
	}
	q := fmt.Sprintf(`SELECT c.id, a.id, a.slug, a.display_name, a.email, a.description, a.avatar_media_id, a.locale
		FROM content_node c JOIN author a ON a.id = c.author_id
		WHERE c.id IN (%s)`, placeholders(len(ids)))
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	rows, err := r.Q.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var contentID int64
		a := &domain.Author{}
		if err := rows.Scan(&contentID, &a.ID, &a.Slug, &a.DisplayName, &a.Email,
			&a.Description, &a.AvatarMediaID, &a.Locale); err != nil {
			return nil, err
		}
		out[contentID] = a
	}
	return out, rows.Err()
}

// ---- Terms -----------------------------------------------------------------

// TermsByContentID returns all terms (categories + tags) attached to a node.
func (r *Repo) TermsByContentID(ctx context.Context, contentID int64) ([]*domain.Term, error) {
	rows, err := r.Q.QueryContext(ctx,
		`SELECT t.id, t.taxonomy_slug, t.slug, t.name, t.description, t.parent_term_id
		FROM content_term ct JOIN term t ON t.id = ct.term_id
		WHERE ct.content_id = ? ORDER BY t.taxonomy_slug, t.slug`, contentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var terms []*domain.Term
	for rows.Next() {
		t := &domain.Term{}
		if err := rows.Scan(&t.ID, &t.TaxonomySlug, &t.Slug, &t.Name, &t.Description, &t.ParentTermID); err != nil {
			return nil, err
		}
		terms = append(terms, t)
	}
	return terms, rows.Err()
}

// TermsByContentIDs batches term lookup for many nodes. Returns contentID -> []Term.
func (r *Repo) TermsByContentIDs(ctx context.Context, ids []int64) (map[int64][]*domain.Term, error) {
	out := map[int64][]*domain.Term{}
	if len(ids) == 0 {
		return out, nil
	}
	q := fmt.Sprintf(`SELECT ct.content_id, t.id, t.taxonomy_slug, t.slug, t.name, t.description, t.parent_term_id
		FROM content_term ct JOIN term t ON t.id = ct.term_id
		WHERE ct.content_id IN (%s)
		ORDER BY ct.content_id, t.taxonomy_slug, t.slug`, placeholders(len(ids)))
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	rows, err := r.Q.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var contentID int64
		t := &domain.Term{}
		if err := rows.Scan(&contentID, &t.ID, &t.TaxonomySlug, &t.Slug, &t.Name,
			&t.Description, &t.ParentTermID); err != nil {
			return nil, err
		}
		out[contentID] = append(out[contentID], t)
	}
	return out, rows.Err()
}

// TermBySlug returns a category or tag by (taxonomy, slug) or sql.ErrNoRows.
func (r *Repo) TermBySlug(ctx context.Context, taxonomy, slug string) (*domain.Term, error) {
	t := &domain.Term{}
	err := r.Q.QueryRowContext(ctx,
		`SELECT id, taxonomy_slug, slug, name, description, parent_term_id
		FROM term WHERE taxonomy_slug = ? AND slug = ?`, taxonomy, slug).
		Scan(&t.ID, &t.TaxonomySlug, &t.Slug, &t.Name, &t.Description, &t.ParentTermID)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// ---- Media -----------------------------------------------------------------

// MediaByID returns a single media asset or sql.ErrNoRows.
func (r *Repo) MediaByID(ctx context.Context, id int64) (*domain.Media, error) {
	m := &domain.Media{}
	err := r.Q.QueryRowContext(ctx,
		`SELECT id, slug, kind, source_url, alt, width, height, mime_type, file_size, caption, locale
		FROM media WHERE id = ?`, id).
		Scan(&m.ID, &m.Slug, &m.Kind, &m.SourceURL, &m.Alt, &m.Width, &m.Height,
			&m.MimeType, &m.FileSize, &m.Caption, &m.Locale)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// MediaByIDs batches media lookups.
func (r *Repo) MediaByIDs(ctx context.Context, ids []int64) (map[int64]*domain.Media, error) {
	out := map[int64]*domain.Media{}
	if len(ids) == 0 {
		return out, nil
	}
	q := fmt.Sprintf(`SELECT id, slug, kind, source_url, alt, width, height, mime_type, file_size, caption, locale
		FROM media WHERE id IN (%s)`, placeholders(len(ids)))
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	rows, err := r.Q.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		m := &domain.Media{}
		if err := rows.Scan(&m.ID, &m.Slug, &m.Kind, &m.SourceURL, &m.Alt, &m.Width,
			&m.Height, &m.MimeType, &m.FileSize, &m.Caption, &m.Locale); err != nil {
			return nil, err
		}
		out[m.ID] = m
	}
	return out, rows.Err()
}

// FeaturedMediaByContentIDs returns contentID -> featured media, batched.
func (r *Repo) FeaturedMediaByContentIDs(ctx context.Context, ids []int64) (map[int64]*domain.Media, error) {
	out := map[int64]*domain.Media{}
	if len(ids) == 0 {
		return out, nil
	}
	q := fmt.Sprintf(`SELECT cm.content_id, m.id, m.slug, m.kind, m.source_url, m.alt,
		m.width, m.height, m.mime_type, m.file_size, m.caption, m.locale
		FROM content_media cm JOIN media m ON m.id = cm.media_id
		WHERE cm.role = 'featured' AND cm.content_id IN (%s)`, placeholders(len(ids)))
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	rows, err := r.Q.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var contentID int64
		m := &domain.Media{}
		if err := rows.Scan(&contentID, &m.ID, &m.Slug, &m.Kind, &m.SourceURL, &m.Alt,
			&m.Width, &m.Height, &m.MimeType, &m.FileSize, &m.Caption, &m.Locale); err != nil {
			return nil, err
		}
		out[contentID] = m
	}
	return out, rows.Err()
}

// ---- Blocks ----------------------------------------------------------------

// BlocksByContentID returns the ordered block list for one node.
func (r *Repo) BlocksByContentID(ctx context.Context, contentID int64) ([]*domain.Block, error) {
	rows, err := r.Q.QueryContext(ctx,
		`SELECT id, content_id, order_index, type, data
		FROM block WHERE content_id = ? ORDER BY order_index`, contentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var blocks []*domain.Block
	for rows.Next() {
		b := &domain.Block{}
		if err := rows.Scan(&b.ID, &b.ContentID, &b.OrderIndex, &b.Type, &b.Data); err != nil {
			return nil, err
		}
		blocks = append(blocks, b)
	}
	return blocks, rows.Err()
}

// BlocksByContentIDs batches block lookups; contentID -> ordered []Block.
func (r *Repo) BlocksByContentIDs(ctx context.Context, ids []int64) (map[int64][]*domain.Block, error) {
	out := map[int64][]*domain.Block{}
	if len(ids) == 0 {
		return out, nil
	}
	q := fmt.Sprintf(`SELECT content_id, id, order_index, type, data
		FROM block WHERE content_id IN (%s) ORDER BY content_id, order_index`, placeholders(len(ids)))
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	rows, err := r.Q.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var contentID int64
		b := &domain.Block{}
		if err := rows.Scan(&contentID, &b.ID, &b.OrderIndex, &b.Type, &b.Data); err != nil {
			return nil, err
		}
		b.ContentID = contentID
		out[contentID] = append(out[contentID], b)
	}
	return out, rows.Err()
}

// ---- SEO -------------------------------------------------------------------

// SEOByContentID returns the Yoast-style SEO row, or sql.ErrNoRows.
func (r *Repo) SEOByContentID(ctx context.Context, contentID int64) (*domain.SEO, error) {
	s := &domain.SEO{}
	err := r.Q.QueryRowContext(ctx,
		`SELECT content_id, title, meta_description, canonical, robots,
		og_json, twitter_json, json_ld, breadcrumbs_json
		FROM seo WHERE content_id = ?`, contentID).
		Scan(&s.ContentID, &s.Title, &s.MetaDescription, &s.Canonical, &s.Robots,
			&s.OGJSON, &s.TwitterJSON, &s.JSONLD, &s.BreadcrumbsJSON)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// SEOByContentIDs batches SEO lookups.
func (r *Repo) SEOByContentIDs(ctx context.Context, ids []int64) (map[int64]*domain.SEO, error) {
	out := map[int64]*domain.SEO{}
	if len(ids) == 0 {
		return out, nil
	}
	q := fmt.Sprintf(`SELECT content_id, title, meta_description, canonical, robots,
		og_json, twitter_json, json_ld, breadcrumbs_json
		FROM seo WHERE content_id IN (%s)`, placeholders(len(ids)))
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	rows, err := r.Q.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		s := &domain.SEO{}
		if err := rows.Scan(&s.ContentID, &s.Title, &s.MetaDescription, &s.Canonical,
			&s.Robots, &s.OGJSON, &s.TwitterJSON, &s.JSONLD, &s.BreadcrumbsJSON); err != nil {
			return nil, err
		}
		out[s.ContentID] = s
	}
	return out, rows.Err()
}
