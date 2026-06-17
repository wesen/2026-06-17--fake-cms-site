package graphql

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
	ast "github.com/graphql-go/graphql/language/ast"
)

// defineScalars returns the custom scalar configs used by the schema.
// They are all "pass-through" JSON values; we do light type coercion in
// resolvers where needed.
func defineScalars() (dateTime, jsonScalar, url *graphql.Scalar) {
	dateTime = graphql.NewScalar(graphql.ScalarConfig{
		Name:        "DateTime",
		Description: "ISO-8601 date-time string.",
		Serialize:   serializeStringScalar,
		ParseValue:  parseStringScalar,
		ParseLiteral: func(vf ast.Value) any {
			return scalarLiteralString(vf)
		},
	})
	jsonScalar = graphql.NewScalar(graphql.ScalarConfig{
		Name:        "JSON",
		Description: "Opaque JSON value (Yoast jsonLd, meta).",
		Serialize: func(value any) any {
			// value is often a JSON string from the DB; decode it.
			if s, ok := value.(string); ok && s != "" {
				var out any
				if err := json.Unmarshal([]byte(s), &out); err == nil {
					return out
				}
			}
			return value
		},
		ParseValue:  parseStringScalar,
		ParseLiteral: func(vf ast.Value) any { return scalarLiteralString(vf) },
	})
	url = graphql.NewScalar(graphql.ScalarConfig{
		Name:        "URL",
		Description: "Absolute URL string.",
		Serialize:   serializeStringScalar,
		ParseValue:  parseStringScalar,
		ParseLiteral: func(vf ast.Value) any { return scalarLiteralString(vf) },
	})
	return
}

func serializeStringScalar(value any) any {
	switch v := value.(type) {
	case nil:
		return nil
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func parseStringScalar(value any) any {
	if s, ok := value.(string); ok {
		return s
	}
	return nil
}

func scalarLiteralString(vf ast.Value) any {
	if sv, ok := vf.(*ast.StringValue); ok {
		return sv.Value
	}
	return nil
}
