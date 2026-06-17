package graphql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/go-go-golems/fake-cms/internal/domain"
	"github.com/go-go-golems/fake-cms/internal/repo"
	"github.com/graphql-go/graphql"
)

// resolvers holds the repo and provides field-resolution methods.
type resolvers struct{ r *repo.Repo }

// Schema builds the executable GraphQL schema over a repository.
func Schema(r *repo.Repo) (graphql.Schema, error) {
	res := &resolvers{r: r}
	dateTime, jsonScalar, urlScalar := defineScalars()

	// ---- Block interface + concrete types -------------------------------
	blockIface := graphql.NewInterface(graphql.InterfaceConfig{
		Name: "Block",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.ID},
			"order": &graphql.Field{Type: graphql.Int},
		},
		ResolveType: resolveBlockType,
	})

	blockParagraph, blockHeading, blockImage, blockList, blockQuote,
		blockEmbed, blockGallery := blockTypes(blockIface)

	blockUnion := graphql.NewUnion(graphql.UnionConfig{
		Name:  "BlockUnion",
		Types: []*graphql.Object{blockParagraph, blockHeading, blockImage, blockList, blockQuote, blockEmbed, blockGallery},
		ResolveType: resolveBlockType,
	})

	// ---- Media ----------------------------------------------------------
	mediaType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Media",
		Fields: graphql.Fields{
			"id":        idField(typeMedia),
			"slug":      strField("Slug"),
			"kind":      strField("Kind"),
			"url":       strField("SourceURL"),
			"sourceUrl": strField("SourceURL"),
			"alt":       strPtrField("Alt"),
			"width":     intPtrField("Width"),
			"height":    intPtrField("Height"),
			"mimeType":  strPtrField("MimeType"),
			"fileSize":  intPtrField("FileSize"),
			"caption":   strPtrField("Caption"),
			"locale":    strField("Locale"),
		},
	})

	// Wire ImageBlock.media + GalleryBlock.images to the Media type via repo.
	wireBlockMedia(blockImage, blockGallery, mediaType, res)

	// ---- Author / Category / Tag ---------------------------------------
	authorType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"id":          idField(typeAuthor),
			"slug":        strField("Slug"),
			"displayName": strField("DisplayName"),
			"email":       strPtrField("Email"),
			"description": strPtrField("Description"),
			"avatar": &graphql.Field{Type: mediaType, Resolve: res.resolveAuthorAvatar},
			"locale": strField("Locale"),
		},
	})

	categoryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id":          idField(typeCategory),
			"slug":        strField("Slug"),
			"name":        strField("Name"),
			"description": strPtrField("Description"),
		},
	})
	tagType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Tag",
		Fields: graphql.Fields{
			"id":          idField(typeTag),
			"slug":        strField("Slug"),
			"name":        strField("Name"),
			"description": strPtrField("Description"),
		},
	})

	// ---- SEO ------------------------------------------------------------
	seoType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SEO",
		Fields: graphql.Fields{
			"title":           strField("Title"),
			"metaDescription": nullableStr("MetaDescription"),
			"canonical":       nullableStr("Canonical"),
			"robots":          nullableStr("Robots"),
			"og":              &graphql.Field{Type: graphql.String, Resolve: res.resolveSEOOG},
			"twitter":         &graphql.Field{Type: graphql.String, Resolve: res.resolveSEOTwitter},
			"jsonLd":          &graphql.Field{Type: jsonScalar, Resolve: res.resolveSEOJSONLD},
			"breadcrumbs":     &graphql.Field{Type: graphql.NewList(graphql.String), Resolve: res.resolveSEOBreadcrumbs},
		},
	})

	// ---- Article / Page -------------------------------------------------
	// We use a single source *domain.ContentNode; Article vs Page resolvers
	// differ only in which fields are populated.
	articleType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"id":            idField(typeArticle),
			"slug":          strField("Slug"),
			"title":         strField("Title"),
			"excerpt":       nullableStr("Excerpt"),
			"status":        strField("Status"),
			"postType":      &graphql.Field{Type: graphql.String, Resolve: res.resolvePostType},
			"locale":        strField("Locale"),
			"publishedAt":   &graphql.Field{Type: dateTime, Resolve: strPtrFieldResolve("PublishedAt")},
			"modifiedAt":    &graphql.Field{Type: dateTime, Resolve: strFieldResolve("ModifiedAt")},
			"author":        &graphql.Field{Type: authorType, Resolve: res.resolveAuthor},
			"categories":    &graphql.Field{Type: graphql.NewList(categoryType), Resolve: res.resolveCategories},
			"tags":          &graphql.Field{Type: graphql.NewList(tagType), Resolve: res.resolveTags},
			"featuredMedia": &graphql.Field{Type: mediaType, Resolve: res.resolveFeaturedMedia},
			"blocks":        &graphql.Field{Type: graphql.NewList(blockUnion), Resolve: res.resolveBlocks},
			"seo":           &graphql.Field{Type: seoType, Resolve: res.resolveSEO},
			"wordCount":     intPtrField("WordCount"),
		},
	})
	pageType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Page",
		Fields: graphql.Fields{
			"id":            idField(typePage),
			"slug":          strField("Slug"),
			"title":         strField("Title"),
			"status":        strField("Status"),
			"locale":        strField("Locale"),
			"publishedAt":   &graphql.Field{Type: dateTime, Resolve: strPtrFieldResolve("PublishedAt")},
			"modifiedAt":    &graphql.Field{Type: dateTime, Resolve: strFieldResolve("ModifiedAt")},
			"menuOrder":     &graphql.Field{Type: graphql.Int, Resolve: intFieldResolve("MenuOrder")},
			"template":      nullableStr("Template"),
			"featuredMedia": &graphql.Field{Type: mediaType, Resolve: res.resolveFeaturedMedia},
			"blocks":        &graphql.Field{Type: graphql.NewList(blockUnion), Resolve: res.resolveBlocks},
			"seo":           &graphql.Field{Type: seoType, Resolve: res.resolveSEO},
		},
	})

	// ---- Connections ----------------------------------------------------
	pageInfoType := graphql.NewObject(graphql.ObjectConfig{
		Name: "PageInfo",
		Fields: graphql.Fields{
			"hasNextPage":     &graphql.Field{Type: graphql.Boolean, Resolve: boolFieldResolve("HasNext")},
			"hasPreviousPage": &graphql.Field{Type: graphql.Boolean, Resolve: boolFieldResolve("HasPrev")},
			"startCursor":     nullableStr("StartCursor"),
			"endCursor":       nullableStr("EndCursor"),
		},
	})
	articleEdgeType := graphql.NewObject(graphql.ObjectConfig{
		Name: "ArticleEdge",
		Fields: graphql.Fields{
			"node":   &graphql.Field{Type: graphql.NewNonNull(articleType), Resolve: edgeNode},
			"cursor": strField("Cursor"),
		},
	})
	articleConnType := graphql.NewObject(graphql.ObjectConfig{
		Name: "ArticleConnection",
		Fields: graphql.Fields{
			"edges":      &graphql.Field{Type: graphql.NewList(articleEdgeType), Resolve: connEdges},
			"pageInfo":   &graphql.Field{Type: graphql.NewNonNull(pageInfoType), Resolve: connPageInfo},
			"totalCount": &graphql.Field{Type: graphql.Int, Resolve: connTotalCount},
		},
	})

	// ---- Query ----------------------------------------------------------
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"node": &graphql.Field{
				Type: graphql.NewInterface(graphql.InterfaceConfig{
					Name: "Node",
					Fields: graphql.Fields{"id": &graphql.Field{Type: graphql.ID}},
					ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
						switch p.Value.(type) {
						case *domain.ContentNode:
							return articleType
						case *domain.Author:
							return authorType
						case *domain.Media:
							return mediaType
						}
						return nil
					},
				}),
				Args: graphql.FieldConfigArgument{"id": &graphql.ArgumentConfig{Type: graphql.ID}},
				Resolve: res.resolveNode(articleType, pageType, authorType, mediaType),
			},
			"article": &graphql.Field{
				Type: articleType,
				Args: argsIDSlug(),
				Resolve: res.resolveArticle,
			},
			"page": &graphql.Field{
				Type: pageType,
				Args: argsIDSlug(),
				Resolve: res.resolvePage,
			},
			"articles": &graphql.Field{
				Type: graphql.NewNonNull(articleConnType),
				Args: graphql.FieldConfigArgument{
					"filter": &graphql.ArgumentConfig{Type: graphql.String},
					"postType": &graphql.ArgumentConfig{Type: graphql.String},
					"first":    &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 20},
					"after":    &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: res.resolveArticles,
			},
			"categories": &graphql.Field{Type: graphql.NewList(categoryType), Resolve: res.resolveAllCategories},
			"category":   &graphql.Field{Type: categoryType, Args: argSlug(), Resolve: res.resolveCategory},
			"tags":       &graphql.Field{Type: graphql.NewList(tagType), Resolve: res.resolveAllTags},
			"tag":        &graphql.Field{Type: tagType, Args: argSlug(), Resolve: res.resolveTag},
			"authors":    &graphql.Field{Type: graphql.NewList(authorType), Resolve: res.resolveAllAuthors},
			"author":     &graphql.Field{Type: authorType, Args: argSlug(), Resolve: res.resolveAuthorBySlug},
			"media": &graphql.Field{
				Type: mediaType,
				Args: graphql.FieldConfigArgument{"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)}},
				Resolve: res.resolveMedia,
			},
		},
	})
	_ = urlScalar
	_ = sql.ErrNoRows

	return graphql.NewSchema(graphql.SchemaConfig{Query: queryType})
}

// ---- field helpers (struct-tag style resolution from domain structs) -------

func idField(typeName string) *graphql.Field {
	return &graphql.Field{Type: graphql.ID, Resolve: func(p graphql.ResolveParams) (any, error) {
		if n, ok := p.Source.(*domain.ContentNode); ok {
			return toGlobalID(typeName, n.ID), nil
		}
		if a, ok := p.Source.(*domain.Author); ok {
			return toGlobalID(typeName, a.ID), nil
		}
		if m, ok := p.Source.(*domain.Media); ok {
			return toGlobalID(typeName, m.ID), nil
		}
		if t, ok := p.Source.(*domain.Term); ok {
			return toGlobalID(typeName, t.ID), nil
		}
		return nil, nil
	}}
}

func strField(structField string) *graphql.Field {
	return &graphql.Field{Type: graphql.String, Resolve: structStringResolve(structField)}
}
func nullableStr(structField string) *graphql.Field {
	return &graphql.Field{Type: graphql.String, Resolve: structStringResolve(structField)}
}
func strPtrField(structField string) *graphql.Field {
	return &graphql.Field{Type: graphql.String, Resolve: structPtrStringResolve(structField)}
}
func intPtrField(structField string) *graphql.Field {
	return &graphql.Field{Type: graphql.Int, Resolve: structPtrIntResolve(structField)}
}

// connection/edge accessors -------------------------------------------------

type connResult interface {
	GetEdges() []edgeData
	GetTotalCount() int
	GetHasNext() bool
	GetEndCursor() string
}

type edgeData struct {
	Node   *domain.ContentNode
	Cursor string
}

func edgeNode(p graphql.ResolveParams) (any, error) { return p.Source.(edgeData).Node, nil }
func connEdges(p graphql.ResolveParams) (any, error) { return p.Source.(*articleConn).Edges, nil }
func connPageInfo(p graphql.ResolveParams) (any, error) {
	c := p.Source.(*articleConn)
	return struct {
		HasNext     bool
		HasPrev     bool
		StartCursor string
		EndCursor   string
	}{HasNext: c.HasNext, EndCursor: c.EndCursor}, nil
}
func connTotalCount(p graphql.ResolveParams) (any, error) {
	return p.Source.(*articleConn).TotalCount, nil
}

// json helpers -------------------------------------------------------------

func decodeJSON(s string) any {
	var out any
	if err := json.Unmarshal([]byte(s), &out); err != nil {
		return nil
	}
	return out
}

var _ = context.Background
var _ = fmt.Sprintf
