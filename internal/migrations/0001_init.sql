-- 0001_init.sql — Fake CMS initial schema.
-- See design-doc §8. SQLite (modernc) with JSON1 enabled.
PRAGMA foreign_keys = ON;

-- ---------------------------------------------------------------------------
-- Lookup / enum tables
-- ---------------------------------------------------------------------------
CREATE TABLE locale (
    code  TEXT PRIMARY KEY,           -- 'fr_FR'
    label TEXT NOT NULL
);
INSERT INTO locale VALUES ('fr_FR','Français'), ('en_US','English');

CREATE TABLE post_type (
    slug          TEXT PRIMARY KEY,    -- WP slug, e.g. 'best-cases'
    graphql_enum  TEXT NOT NULL UNIQUE,-- 'BEST_CASES'
    label         TEXT NOT NULL,       -- 'Best cases'
    hierarchical  INTEGER NOT NULL DEFAULT 0
);
INSERT INTO post_type VALUES
 ('actualites','ACTUALITES','Actualités',0),
 ('best-cases','BEST_CASES','Best cases',0),
 ('etudes','ETUDES','Études',0),
 ('blog','BLOG','Blog',0),
 ('slider-de-une','SLIDER_DE_UNE','Slider de une',0),
 ('cartouches-home','CARTOUCHES_HOME','Cartouches home',0),
 ('non-classe','NON_CLASSE','Non classé',0);

CREATE TABLE taxonomy (
    slug         TEXT PRIMARY KEY,     -- 'category' | 'post_tag'
    label        TEXT NOT NULL,
    hierarchical INTEGER NOT NULL
);
INSERT INTO taxonomy VALUES ('category','Categories',1), ('post_tag','Tags',0);

-- ---------------------------------------------------------------------------
-- Media (images, video, pdf, ...)
-- ---------------------------------------------------------------------------
CREATE TABLE media (
    id         INTEGER PRIMARY KEY,
    slug       TEXT NOT NULL,
    kind       TEXT NOT NULL CHECK (kind IN ('IMAGE','VIDEO','PDF','AUDIO','FILE')),
    source_url TEXT NOT NULL,
    alt        TEXT,
    width      INTEGER,
    height     INTEGER,
    mime_type  TEXT,
    file_size  INTEGER,
    caption    TEXT,
    locale     TEXT NOT NULL DEFAULT 'fr_FR' REFERENCES locale(code)
);

-- ---------------------------------------------------------------------------
-- Authors
-- ---------------------------------------------------------------------------
CREATE TABLE author (
    id               INTEGER PRIMARY KEY,
    slug             TEXT NOT NULL UNIQUE,        -- WP user_nicename
    display_name     TEXT NOT NULL,
    email            TEXT,                        -- legacy: may be NULL / raw email
    description      TEXT,
    avatar_media_id  INTEGER REFERENCES media(id),
    locale           TEXT NOT NULL DEFAULT 'fr_FR' REFERENCES locale(code)
);

-- ---------------------------------------------------------------------------
-- Taxonomy terms (categories are hierarchical; tags are flat)
-- ---------------------------------------------------------------------------
CREATE TABLE term (
    id             INTEGER PRIMARY KEY,
    taxonomy_slug  TEXT NOT NULL REFERENCES taxonomy(slug),
    slug           TEXT NOT NULL,                 -- e.g. 'moijeune'
    name           TEXT NOT NULL,
    description    TEXT,
    parent_term_id INTEGER REFERENCES term(id),
    UNIQUE(taxonomy_slug, slug)
);

-- ---------------------------------------------------------------------------
-- Polymorphic content table (articles + pages in one table, like wp_posts)
-- ---------------------------------------------------------------------------
CREATE TABLE content_node (
    id             INTEGER PRIMARY KEY,
    kind           TEXT NOT NULL CHECK (kind IN ('ARTICLE','PAGE')),
    slug           TEXT NOT NULL,
    title          TEXT NOT NULL,
    excerpt        TEXT,
    status         TEXT NOT NULL DEFAULT 'PUBLISH',
    locale         TEXT NOT NULL DEFAULT 'fr_FR' REFERENCES locale(code),
    published_at   TEXT,                          -- ISO-8601; nullable for drafts
    modified_at    TEXT NOT NULL,
    author_id      INTEGER REFERENCES author(id),
    parent_page_id INTEGER REFERENCES content_node(id),
    menu_order     INTEGER NOT NULL DEFAULT 0,
    template       TEXT,                          -- page template PHP name
    word_count     INTEGER,
    UNIQUE(kind, slug)
);

-- Per-article post type (only for ARTICLE rows).
CREATE TABLE article_post_type (
    content_id INTEGER PRIMARY KEY REFERENCES content_node(id),
    post_type  TEXT NOT NULL REFERENCES post_type(slug)
);

-- ---------------------------------------------------------------------------
-- Content body: ordered, typed blocks. `data` is JSON (json_extract available).
-- ---------------------------------------------------------------------------
CREATE TABLE block (
    id          INTEGER PRIMARY KEY,
    content_id  INTEGER NOT NULL REFERENCES content_node(id),
    order_index INTEGER NOT NULL,
    type        TEXT NOT NULL CHECK (type IN
                  ('paragraph','heading','image','list','quote','embed','gallery')),
    data        TEXT NOT NULL,           -- JSON; validated by app layer
    UNIQUE(content_id, order_index)
);

-- M:N content <-> term
CREATE TABLE content_term (
    content_id INTEGER NOT NULL REFERENCES content_node(id),
    term_id    INTEGER NOT NULL REFERENCES term(id),
    PRIMARY KEY (content_id, term_id)
);

-- M:N content <-> media (featured / inline / gallery)
CREATE TABLE content_media (
    content_id INTEGER NOT NULL REFERENCES content_node(id),
    media_id   INTEGER NOT NULL REFERENCES media(id),
    role       TEXT NOT NULL DEFAULT 'inline',
    PRIMARY KEY (content_id, media_id, role)
);

-- ---------------------------------------------------------------------------
-- Yoast-style SEO (1:1 with content_node)
-- ---------------------------------------------------------------------------
CREATE TABLE seo (
    content_id       INTEGER PRIMARY KEY REFERENCES content_node(id),
    title            TEXT NOT NULL,
    meta_description TEXT,
    canonical        TEXT,
    robots           TEXT DEFAULT 'index,follow',
    og_json          TEXT,        -- JSON {title,description,image_id,type,...}
    twitter_json     TEXT,        -- JSON
    json_ld          TEXT NOT NULL,        -- raw Yoast graph blob (JSON)
    breadcrumbs_json TEXT NOT NULL         -- JSON array
);

-- ---------------------------------------------------------------------------
-- Menus (wp_nav_menu)
-- ---------------------------------------------------------------------------
CREATE TABLE menu (
    slug TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE menu_item (
    id          INTEGER PRIMARY KEY,
    menu_slug   TEXT NOT NULL REFERENCES menu(slug),
    parent_id   INTEGER REFERENCES menu_item(id),
    label       TEXT NOT NULL,
    url         TEXT NOT NULL,
    order_index INTEGER NOT NULL
);

-- schema_version is owned by the migration runner (repo.Migrate); do NOT
-- create or insert into it from migration files.
