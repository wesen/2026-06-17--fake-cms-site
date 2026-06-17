// Package graphql implements the schema-first GraphQL Content Delivery API
// on top of internal/repo. It knows nothing about SQL; dependencies point UP.
package graphql

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Global IDs use the Relay convention base64("<TypeName>:<rowID>").
// node(id:) decodes the prefix to dispatch to the right loader.

// toGlobalID encodes a typed row id.
func toGlobalID(typeName string, id int64) string {
	return base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%d", typeName, id)))
}

// fromGlobalID decodes a global id into (typeName, rowID).
func fromGlobalID(s string) (typeName string, rowID int64, err error) {
	raw, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", 0, err
	}
	parts := strings.SplitN(string(raw), ":", 2)
	if len(parts) != 2 {
		return "", 0, errors.New("invalid global id")
	}
	id, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return "", 0, err
	}
	return parts[0], id, nil
}

// type names used in global IDs (must match the GraphQL type names).
const (
	typeArticle  = "Article"
	typePage     = "Page"
	typeAuthor   = "Author"
	typeCategory = "Category"
	typeTag      = "Tag"
	typeMedia    = "Media"
	typeBlock    = "Block"
)
