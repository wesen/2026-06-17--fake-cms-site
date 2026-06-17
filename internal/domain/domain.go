// Package domain defines plain Go structs for the CMS content model.
//
// These structs carry NO GraphQL and NO SQL tags: they are the boundary type
// exchanged between the repository (SQL) and the graphql resolvers. Keeping
// them framework-free is what makes the mock testable and teachable.
package domain

// Locale codes.
const (
	LocaleFR = "fr_FR"
	LocaleUS = "en_US"
)

// PostType graphql enum values (mirror schema.graphql PostType).
const (
	PostActualites    = "ACTUALITES"
	PostBestCases     = "BEST_CASES"
	PostEtudes        = "ETUDES"
	PostBlog          = "BLOG"
	PostSliderDeUne   = "SLIDER_DE_UNE"
	PostCartouchesHom = "CARTOUCHES_HOME"
	PostNonClasse     = "NON_CLASSE"
)

// Content kinds.
const (
	KindArticle = "ARTICLE"
	KindPage    = "PAGE"
)

// Status values.
const (
	StatusPublish = "PUBLISH"
	StatusDraft   = "DRAFT"
)

// Block kinds.
const (
	BlockParagraph = "paragraph"
	BlockHeading   = "heading"
	BlockImage     = "image"
	BlockList      = "list"
	BlockQuote     = "quote"
	BlockEmbed     = "embed"
	BlockGallery   = "gallery"
)

// Author is a person who wrote content.
type Author struct {
	ID            int64
	Slug          string
	DisplayName   string
	Email         *string // legacy: may be NULL or a raw email
	Description   *string
	AvatarMediaID *int64
	Locale        string
}

// Media is an asset (image, video, pdf...).
type Media struct {
	ID        int64
	Slug      string
	Kind      string // IMAGE|VIDEO|PDF|AUDIO|FILE
	SourceURL string
	Alt       *string
	Width     *int64
	Height    *int64
	MimeType  *string
	FileSize  *int64
	Caption   *string
	Locale    string
}

// Term is a category or tag.
type Term struct {
	ID           int64
	TaxonomySlug string // category | post_tag
	Slug         string
	Name         string
	Description  *string
	ParentTermID *int64
}

// ContentNode is the shared shape for Article and Page rows.
type ContentNode struct {
	ID           int64
	Kind         string // ARTICLE|PAGE
	Slug         string
	Title        string
	Excerpt      string
	Status       string
	Locale       string
	PublishedAt  *string // ISO-8601
	ModifiedAt   string
	AuthorID     *int64
	ParentPageID *int64
	MenuOrder    int64
	Template     string
	WordCount    *int64
}

// PostType links an ARTICLE content node to its WP custom post type.
type PostTypeRow struct {
	ContentID int64
	PostType  string // graphql enum value, e.g. BEST_CASES
}

// Block is one ordered, typed unit of body content. Data is the raw JSON
// payload (decoded lazily by resolvers).
type Block struct {
	ID         int64
	ContentID  int64
	OrderIndex int64
	Type       string
	Data       string // JSON string
}

// SEO is the Yoast-style per-content metadata.
type SEO struct {
	ContentID       int64
	Title           string
	MetaDescription string
	Canonical       string
	Robots          string
	OGJSON          string
	TwitterJSON     string
	JSONLD          string
	BreadcrumbsJSON string
}

// ArticleFilter captures the WHERE-clause inputs for ListArticles.
type ArticleFilter struct {
	PostType       string
	Status         string
	CategorySlug   string
	TagSlug        string
	AuthorSlug     string
	Search         string
	PublishedAfter *string
	PublishedBefore *string
}

// Page is a cursor pagination window.
type Page struct {
	First  int
	After  string // opaque base64 cursor (sort-key encoded)
}
