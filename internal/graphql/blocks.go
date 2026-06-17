package graphql

import (
	"context"
	"encoding/json"

	"github.com/go-go-golems/fake-cms/internal/domain"
	"github.com/graphql-go/graphql"
)

// blockValue is the resolver-level representation of a block: the typed
// payload plus the order index. The GraphQL id is global-encoded.
type blockValue struct {
	Type  string
	Order int64
	Data  map[string]any
	ID    int64
}

// resolveBlockType dispatches a blockValue to its concrete GraphQL object.
func resolveBlockType(p graphql.ResolveTypeParams) *graphql.Object {
	bv, ok := p.Value.(blockValue)
	if !ok {
		return nil
	}
	switch bv.Type {
	case domain.BlockParagraph:
		return objParagraph
	case domain.BlockHeading:
		return objHeading
	case domain.BlockImage:
		return objImage
	case domain.BlockList:
		return objList
	case domain.BlockQuote:
		return objQuote
	case domain.BlockEmbed:
		return objEmbed
	case domain.BlockGallery:
		return objGallery
	}
	return nil
}

// Concrete block object singletons, filled by blockTypes().
var (
	objParagraph, objHeading, objImage, objList, objQuote, objEmbed, objGallery *graphql.Object
)

// blockTypes builds the seven concrete block object types.
func blockTypes(iface *graphql.Interface) (par, hea, img, lst, quo, emb, gal *graphql.Object) {
	common := func(name string) *graphql.Object {
		return graphql.NewObject(graphql.ObjectConfig{
			Name: name,
			Interfaces: []*graphql.Interface{iface},
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.ID,
					Resolve: func(p graphql.ResolveParams) (any, error) {
						return toGlobalID(typeBlock, p.Source.(blockValue).ID), nil
					},
				},
				"order": &graphql.Field{
					Type: graphql.Int,
					Resolve: func(p graphql.ResolveParams) (any, error) {
						return int(p.Source.(blockValue).Order), nil
					},
				},
			},
		})
	}
	dataStr := func(key string) *graphql.Field {
		return &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (any, error) {
				if v, ok := p.Source.(blockValue).Data[key]; ok {
					if s, ok := v.(string); ok && s != "" {
						return s, nil
					}
				}
				return nil, nil
			},
		}
	}
	dataInt := func(key string) *graphql.Field {
		return &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (any, error) {
				if v, ok := p.Source.(blockValue).Data[key]; ok {
					if f, ok := v.(float64); ok {
						return int(f), nil
					}
				}
				return nil, nil
			},
		}
	}

	par = common("ParagraphBlock")
	par.AddFieldConfig("text", dataStr("text"))
	par.AddFieldConfig("align", dataStr("align"))

	hea = common("HeadingBlock")
	hea.AddFieldConfig("level", dataInt("level"))
	hea.AddFieldConfig("text", dataStr("text"))
	hea.AddFieldConfig("anchor", dataStr("anchor"))

	img = common("ImageBlock")
	img.AddFieldConfig("caption", dataStr("caption"))
	img.AddFieldConfig("link", dataStr("link"))
	img.AddFieldConfig("size", dataStr("size"))

	lst = common("ListBlock")
	lst.AddFieldConfig("ordered", &graphql.Field{
		Type: graphql.Boolean,
		Resolve: func(p graphql.ResolveParams) (any, error) {
			if v, ok := p.Source.(blockValue).Data["ordered"]; ok {
				if b, ok := v.(bool); ok {
					return b, nil
				}
			}
			return false, nil
		},
	})
	lst.AddFieldConfig("items", &graphql.Field{
		Type: graphql.NewList(graphql.String),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			if v, ok := p.Source.(blockValue).Data["items"]; ok {
				if arr, ok := v.([]any); ok {
					out := make([]string, 0, len(arr))
					for _, it := range arr {
						if s, ok := it.(string); ok {
							out = append(out, s)
						}
					}
					return out, nil
				}
			}
			return []string{}, nil
		},
	})

	quo = common("QuoteBlock")
	quo.AddFieldConfig("text", dataStr("text"))
	quo.AddFieldConfig("citation", dataStr("citation"))

	emb = common("EmbedBlock")
	emb.AddFieldConfig("provider", dataStr("provider"))
	emb.AddFieldConfig("url", dataStr("url"))
	emb.AddFieldConfig("caption", dataStr("caption"))

	gal = common("GalleryBlock")
	gal.AddFieldConfig("columns", dataInt("columns"))

	objParagraph, objHeading, objImage, objList, objQuote, objEmbed, objGallery = par, hea, img, lst, quo, emb, gal
	return
}

// toBlockValue converts a stored domain.Block into a resolver blockValue,
// decoding its JSON data payload.
func toBlockValue(b *domain.Block) blockValue {
	var data map[string]any
	if b.Data != "" {
		_ = json.Unmarshal([]byte(b.Data), &data)
	}
	if data == nil {
		data = map[string]any{}
	}
	return blockValue{Type: b.Type, Order: b.OrderIndex, Data: data, ID: b.ID}
}

// wireBlockMedia adds ImageBlock.media and GalleryBlock.images fields, both
// resolved through the repo (MediaByID), so blocks reference real Media nodes.
func wireBlockMedia(imageT, galleryT *graphql.Object, mediaType *graphql.Object, res *resolvers) {
	mediaByID := func(bv blockValue, idAny any) (*domain.Media, error) {
		if idAny == nil {
			return nil, nil
		}
		return res.r.MediaByID(context.Background(), toInt64(idAny))
	}
	_ = mediaByID

	imageT.AddFieldConfig("media", &graphql.Field{
		Type: graphql.NewNonNull(mediaType),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			bv := p.Source.(blockValue)
			m, err := res.r.MediaByID(p.Context, blockMediaID(bv, "media_id"))
			if err != nil || m == nil {
				// Fall back to a placeholder media so the non-null contract holds
				// even if a seed row was deleted.
				return &domain.Media{ID: 0, Kind: "IMAGE", SourceURL: "", Locale: "fr_FR"}, nil
			}
			return m, nil
		},
	})

	galleryT.AddFieldConfig("images", &graphql.Field{
		Type: graphql.NewList(graphql.NewNonNull(mediaType)),
		Resolve: func(p graphql.ResolveParams) (any, error) {
			bv := p.Source.(blockValue)
			raw, _ := bv.Data["images"].([]any)
			ids := make([]int64, 0, len(raw))
			for _, it := range raw {
				if id := toInt64(it); id > 0 {
					ids = append(ids, id)
				}
			}
			media, err := res.r.MediaByIDs(p.Context, ids)
			if err != nil {
				return nil, err
			}
			out := make([]*domain.Media, 0, len(ids))
			for _, id := range ids {
				if m, ok := media[id]; ok {
					out = append(out, m)
				}
			}
			return out, nil
		},
	})
}

func blockMediaID(bv blockValue, key string) int64 {
	return toInt64(bv.Data[key])
}

func toInt64(v any) int64 {
	switch n := v.(type) {
	case float64:
		return int64(n)
	case int:
		return int64(n)
	case int64:
		return n
	}
	return 0
}
