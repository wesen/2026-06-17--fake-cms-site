package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-go-golems/fake-cms/internal/domain"
)

// Seed builds a deterministic dataset. The PRNG is seeded with a constant, so
// two Seed() runs on an empty DB produce byte-identical rows. No time.Now()
// is used: all timestamps come from a fixed distribution.
//
// The data is loosely derived from the scraped 20minutes-media.com sources in
// ttmp/.../sources/_raw (author slugs, post types, rubrique/tag slugs), but the
// text is synthetic placeholder content so the bundle is self-contained.
func (r *Repo) Seed(ctx context.Context) error {
	if err := r.Recreate(ctx); err != nil {
		return err
	}
	rng := rand.New(rand.NewSource(42)) //nolint:gosec // deterministic by design

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin seed tx: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	s := &seeder{tx: tx, rng: rng}

	authors, err := s.seedAuthors()
	if err != nil {
		return err
	}
	media, err := s.seedMedia(60)
	if err != nil {
		return err
	}
	cats, tags, err := s.seedTerms()
	if err != nil {
		return err
	}
	if err := s.seedArticles(authors, media, cats, tags); err != nil {
		return err
	}
	if err := s.seedPages(authors, media); err != nil {
		return err
	}
	if err := s.seedMenu(); err != nil {
		return err
	}

	return tx.Commit()
}

type seeder struct {
	tx  *sql.Tx
	rng *rand.Rand
}

// ---- helpers ---------------------------------------------------------------

func (s *seeder) exec(query string, args ...any) (int64, error) {
	res, err := s.tx.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return id, nil
}

func (s *seeder) pickString(pool []string) string {
	return pool[s.rng.Intn(len(pool))]
}

func (s *seeder) pickN(pool []string, n int) []string {
	if n > len(pool) {
		n = len(pool)
	}
	s.rng.Shuffle(len(pool), func(i, j int) { pool[i], pool[j] = pool[j], pool[i] })
	out := make([]string, n)
	copy(out, pool[:n])
	return out
}

func (s *seeder) mustJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		// only used for simple maps/slices; never errors in practice
		return "{}"
	}
	return string(b)
}

// fixedDate returns a deterministic timestamp in [2013-01-01, 2025-12-31].
var epochBase = time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC)

func (s *seeder) fixedDate() string {
	// ~13 years span in hours
	d := time.Duration(s.rng.Intn(13*365*24)) * time.Hour
	t := epochBase.Add(d)
	return t.UTC().Format("2006-01-02T15:04:05Z07:00")
}

// ---- authors ---------------------------------------------------------------

// real-ish WP user_nicename slugs (from ttmp sources/_raw/sitemap-urls.json).
var authorSpecs = []struct {
	slug, name, email, desc string
}{
	{"eking20minutes-fr", "Élise King", "eking@20minutes.fr", "Rédactrice en chef adjointe."},
	{"adminclic-clic-com", "admin@clic-clic.com", "admin@clic-clic.com", "Compte administration hérité de l'ancien CMS."}, // legacy email, preserved
	{"abaron", "Aurélie Baron", "abaron@20minutes.fr", "Journaliste culture."},
	{"jdoe", "Jean Doe", "", "Reporter."},
}

func (s *seeder) seedAuthors() ([]*domain.Author, error) {
	out := make([]*domain.Author, 0, len(authorSpecs))
	for _, a := range authorSpecs {
		id, err := s.exec(`INSERT INTO author (slug, display_name, email, description, locale)
			VALUES (?,?,?,?,?)`, a.slug, a.name, a.email, a.desc, domain.LocaleFR)
		if err != nil {
			return nil, fmt.Errorf("insert author %s: %w", a.slug, err)
		}
		out = append(out, &domain.Author{ID: id, Slug: a.slug, DisplayName: a.name})
	}
	return out, nil
}

// ---- media -----------------------------------------------------------------

func (s *seeder) seedMedia(n int) ([]int64, error) {
	ids := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		slug := fmt.Sprintf("media-%03d", i)
		w := 722
		h := 460
		seed := slug
		url := fmt.Sprintf("https://picsum.photos/seed/%s/%d/%d", seed, w, h)
		id, err := s.exec(`INSERT INTO media (slug, kind, source_url, alt, width, height, mime_type, locale)
			VALUES (?,?,?,?,?,?,?,?)`,
			slug, "IMAGE", url, "Illustration "+slug, w, h, "image/jpeg", domain.LocaleFR)
		if err != nil {
			return nil, fmt.Errorf("insert media %d: %w", i, err)
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// ---- terms -----------------------------------------------------------------

// categories mirror the real /archives/* tree; tags mirror /rubrique/* slugs.
var categorySpecs = []struct{ slug, name, parent string }{
	{"actualites", "Actualités", ""},
	{"best-cases", "Best cases", ""},
	{"etudes", "Études", ""},
	{"blog", "Blog", ""},
	{"actualites-en-une", "À la une", "actualites"},
	{"non-classe", "Non classé", ""},
}

var tagSpecs = []string{
	"moijeune", "audience", "offres", "offre", "communiques-de-presse",
	"conference", "contacts", "content", "culture-g", "dossier-special",
	"edito", "election", "chiffres-cles", "display", "video", "data",
	"print", "tv", "newsletter", "panel", "efficacite", "marketing",
	"rentree", "canne", "covid-19", "virale", "branding", "performance",
	"engagement", "innovation",
}

func (s *seeder) seedTerms() (cats, tags map[string]int64, err error) {
	cats = map[string]int64{}
	tags = map[string]int64{}

	// First pass: root categories.
	for _, c := range categorySpecs {
		if c.parent != "" {
			continue
		}
		id, e := s.exec(`INSERT INTO term (taxonomy_slug, slug, name) VALUES ('category', ?, ?)`,
			c.slug, c.name)
		if e != nil {
			return nil, nil, fmt.Errorf("insert category %s: %w", c.slug, e)
		}
		cats[c.slug] = id
	}
	// Second pass: children.
	for _, c := range categorySpecs {
		if c.parent == "" {
			continue
		}
		pid := cats[c.parent]
		id, e := s.exec(`INSERT INTO term (taxonomy_slug, slug, name, parent_term_id)
			VALUES ('category', ?, ?, ?)`, c.slug, c.name, pid)
		if e != nil {
			return nil, nil, fmt.Errorf("insert child category %s: %w", c.slug, e)
		}
		cats[c.slug] = id
	}

	for _, t := range tagSpecs {
		name := strings.ReplaceAll(t, "-", " ")
		id, e := s.exec(`INSERT INTO term (taxonomy_slug, slug, name) VALUES ('post_tag', ?, ?)`,
			t, name)
		if e != nil {
			return nil, nil, fmt.Errorf("insert tag %s: %w", t, e)
		}
		tags[t] = id
	}
	return cats, tags, nil
}

// ---- articles --------------------------------------------------------------

// postTypeCount mirrors the real volume distribution (best-cases dominates).
var postTypeSpecs = []struct {
	enum  string
	count int
}{
	{domain.PostBestCases, 40},
	{domain.PostActualites, 60},
	{domain.PostEtudes, 20},
	{domain.PostBlog, 10},
	{domain.PostSliderDeUne, 6},
	{domain.PostCartouchesHom, 4},
}

var titleNouns = []string{"Campagne", "Lancement", "Étude", "Cas client", "Reportage",
	"Enquête", "Interview", "Dossier", "Tribune", "Décryptage"}
var titleAdjs = []string{"historique", "inédit", "premium", "viral", "data-driven",
	"engagé", "créatif", "performant", "stratégique", "iconique"}
var brandPool = []string{"Suez", "RATP", "HBO", "Francine", "Auchan", "Cofidis",
	"Fleury Michon", "Oenobiol", "Mini", "Fuerteventura", "Warner", "Secours Catholique"}

func (s *seeder) seedArticles(authors []*domain.Author, mediaIDs []int64,
	cats, tags map[string]int64) error {
	// Map graphql enum -> WP post_type slug + a primary category.
	ptToCat := map[string]string{
		domain.PostBestCases:     "best-cases",
		domain.PostActualites:    "actualites",
		domain.PostEtudes:        "etudes",
		domain.PostBlog:          "blog",
		domain.PostSliderDeUne:   "actualites",
		domain.PostCartouchesHom: "actualites",
	}
	ptToWPSlug := map[string]string{
		domain.PostBestCases:     "best-cases",
		domain.PostActualites:    "actualites",
		domain.PostEtudes:        "etudes",
		domain.PostBlog:          "blog",
		domain.PostSliderDeUne:   "slider-de-une",
		domain.PostCartouchesHom: "cartouches-home",
	}

	a := 0
	for _, pt := range postTypeSpecs {
		for i := 0; i < pt.count; i++ {
			a++
			title := fmt.Sprintf("%s %s : %s",
				s.pickString(titleNouns), s.pickString(titleAdjs), s.pickString(brandPool))
			slug := fmt.Sprintf("%s-%d", strings.ToLower(strings.ReplaceAll(title, " ", "-")), a)
			slug = strings.ReplaceAll(slug, ":", "")
			author := authors[s.rng.Intn(len(authors))]
			pub := s.fixedDate()

			// word_count derived from block count later; set a placeholder.
			nodeID, err := s.exec(`INSERT INTO content_node
				(kind, slug, title, excerpt, status, locale, published_at, modified_at, author_id, word_count)
				VALUES ('ARTICLE', ?, ?, ?, 'PUBLISH', ?, ?, ?, ?, ?)`,
				slug, title, s.excerpt(), domain.LocaleFR, pub, pub, author.ID, s.rng.Intn(400)+40)
			if err != nil {
				return fmt.Errorf("insert article node #%d: %w", a, err)
			}

			if _, err := s.exec(`INSERT INTO article_post_type (content_id, post_type) VALUES (?, ?)`,
				nodeID, ptToWPSlug[pt.enum]); err != nil {
				return fmt.Errorf("insert post_type link #%d: %w", a, err)
			}

			// categories: primary + occasionally a child + "non-classe".
			primaryCat := ptToCat[pt.enum]
			chosenCats := []int64{cats[primaryCat]}
			if s.rng.Intn(3) == 0 {
				chosenCats = append(chosenCats, cats["actualites-en-une"])
			}
			if err := s.linkTerms(nodeID, chosenCats); err != nil {
				return err
			}
			// tags: 1-5 random.
			tagKeys := make([]string, 0, len(tags))
			for k := range tags {
				tagKeys = append(tagKeys, k)
			}
			picked := s.pickN(tagKeys, s.rng.Intn(5)+1)
			tagIDs := make([]int64, len(picked))
			for i, k := range picked {
				tagIDs[i] = tags[k]
			}
			if err := s.linkTerms(nodeID, tagIDs); err != nil {
				return err
			}

			// featured media + inline media.
			featured := mediaIDs[s.rng.Intn(len(mediaIDs))]
			if _, err := s.exec(`INSERT INTO content_media (content_id, media_id, role) VALUES (?, ?, 'featured')`,
				nodeID, featured); err != nil {
				return err
			}

			// blocks
			if err := s.seedBlocks(nodeID, mediaIDs); err != nil {
				return err
			}

			// seo
			if err := s.seedSEO(nodeID, title, slug, ptToWPSlug[pt.enum], featured); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *seeder) linkTerms(contentID int64, termIDs []int64) error {
	for _, tid := range termIDs {
		if _, err := s.exec(`INSERT OR IGNORE INTO content_term (content_id, term_id) VALUES (?, ?)`,
			contentID, tid); err != nil {
			return err
		}
	}
	return nil
}

func (s *seeder) excerpt() string {
	return "Lorem ipsum dolor sit amet, consectetur adipiscing elit — extrait éditorial synthétique pour le workshop."
}

// ---- blocks ----------------------------------------------------------------

func (s *seeder) seedBlocks(contentID int64, mediaIDs []int64) error {
	n := s.rng.Intn(10) + 5 // 5..14 blocks
	for i := 0; i < n; i++ {
		var btype, data string
		switch s.rng.Intn(7) {
		case 0: // heading
			level := s.rng.Intn(3) + 2 // h2-h4
			data = s.mustJSON(map[string]any{"level": level, "text": "Section " + s.pickString(titleNouns), "anchor": fmt.Sprintf("section-%d", i)})
			btype = domain.BlockHeading
		case 1: // image
			mid := mediaIDs[s.rng.Intn(len(mediaIDs))]
			data = s.mustJSON(map[string]any{"media_id": mid, "caption": "Légende illustrative.", "size": "large"})
			btype = domain.BlockImage
		case 2: // list
			items := []string{"Premier point", "Deuxième point", "Troisième point"}
			data = s.mustJSON(map[string]any{"ordered": s.rng.Intn(2) == 1, "items": items})
			btype = domain.BlockList
		case 3: // quote
			data = s.mustJSON(map[string]any{"text": "« Une citation marquante de l'article. »", "citation": "Source"})
			btype = domain.BlockQuote
		case 4: // embed
			data = s.mustJSON(map[string]any{"provider": "youtube", "url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ", "caption": "Vidéo associée."})
			btype = domain.BlockEmbed
		case 5: // gallery
			imgs := make([]int64, 3)
			for k := range imgs {
				imgs[k] = mediaIDs[s.rng.Intn(len(mediaIDs))]
			}
			data = s.mustJSON(map[string]any{"images": imgs, "columns": 3})
			btype = domain.BlockGallery
		default: // paragraph (most common)
			words := []string{"Lorem", "ipsum", "dolor", "sit", "amet", "editorial", "contenu", "marque", "audience", "stratégie"}
			var sb strings.Builder
			ln := s.rng.Intn(40) + 15
			for j := 0; j < ln; j++ {
				sb.WriteString(words[s.rng.Intn(len(words))])
				sb.WriteByte(' ')
			}
			data = s.mustJSON(map[string]any{"text": strings.TrimSpace(sb.String()), "align": "NONE"})
			btype = domain.BlockParagraph
		}
		if _, err := s.exec(`INSERT INTO block (content_id, order_index, type, data) VALUES (?, ?, ?, ?)`,
			contentID, i, btype, data); err != nil {
			return fmt.Errorf("insert block %d: %w", i, err)
		}
	}
	return nil
}

// ---- seo -------------------------------------------------------------------

func (s *seeder) seedSEO(contentID int64, title, slug, wpType string, featuredMedia int64) error {
	canonical := fmt.Sprintf("https://www.20minutes-media.com/%s/%s", wpType, slug)
	og := s.mustJSON(map[string]any{
		"title": title, "description": s.excerpt(),
		"image_id": featuredMedia, "type": "article", "locale": "fr_FR", "site_name": "20 Minutes Media",
	})
	tw := s.mustJSON(map[string]any{"card": "summary_large_image", "title": title, "description": s.excerpt()})
	jsonLd := s.mustJSON(map[string]any{
		"@context": "https://schema.org", "@type": "Article",
		"headline": title, "inLanguage": "fr-FR",
		"mainEntityOfPage": canonical,
	})
	crumbs := s.mustJSON([]map[string]any{
		{"label": "Accueil", "url": "https://www.20minutes-media.com/"},
		{"label": wpType, "url": fmt.Sprintf("https://www.20minutes-media.com/%s", wpType)},
		{"label": title, "url": canonical},
	})
	_, err := s.exec(`INSERT INTO seo
		(content_id, title, meta_description, canonical, robots, og_json, twitter_json, json_ld, breadcrumbs_json)
		VALUES (?,?,?,?,?,?,?,?,?)`,
		contentID, title, s.excerpt(), canonical, "index,follow", og, tw, jsonLd, crumbs)
	return err
}

// ---- pages -----------------------------------------------------------------

func (s *seeder) seedPages(authors []*domain.Author, mediaIDs []int64) error {
	pages := []struct{ slug, title, template string }{
		{"display", "Display", "page-display.php"},
		{"video", "Vidéo", "page-video.php"},
		{"presse", "Presse", "default"},
		{"a-propos", "À propos", "default"},
		{"panel", "Panel lecteurs", "default"},
		{"mentions-legales", "Mentions légales", "default"},
	}
	for i, p := range pages {
		pub := s.fixedDate()
		nodeID, err := s.exec(`INSERT INTO content_node
			(kind, slug, title, excerpt, status, locale, published_at, modified_at, menu_order, template)
			VALUES ('PAGE', ?, ?, ?, 'PUBLISH', ?, ?, ?, ?, ?)`,
			p.slug, p.title, p.title, domain.LocaleFR, pub, pub, i, p.template)
		if err != nil {
			return fmt.Errorf("insert page %s: %w", p.slug, err)
		}
		featured := mediaIDs[s.rng.Intn(len(mediaIDs))]
		if _, err := s.exec(`INSERT INTO content_media (content_id, media_id, role) VALUES (?, ?, 'featured')`,
			nodeID, featured); err != nil {
			return err
		}
		if err := s.seedBlocks(nodeID, mediaIDs); err != nil {
			return err
		}
		if err := s.seedSEO(nodeID, p.title, p.slug, "", featured); err != nil {
			return err
		}
	}
	return nil
}

// ---- menu ------------------------------------------------------------------

func (s *seeder) seedMenu() error {
	if _, err := s.exec(`INSERT INTO menu (slug, name) VALUES ('main', 'Menu principal')`); err != nil {
		return err
	}
	items := []struct {
		label, url string
	}{
		{"Accueil", "https://www.20minutes-media.com/"},
		{"Nos solutions", "https://www.20minutes-media.com/nos-solutions/display"},
		{"Presse", "https://www.20minutes-media.com/presse"},
		{"À propos", "https://www.20minutes-media.com/a-propos"},
		{"Contacts", "https://www.20minutes-media.com/contacts"},
	}
	for i, it := range items {
		if _, err := s.exec(`INSERT INTO menu_item (menu_slug, label, url, order_index) VALUES ('main', ?, ?, ?)`,
			it.label, it.url, i); err != nil {
			return err
		}
	}
	return nil
}
