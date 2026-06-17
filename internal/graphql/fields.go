package graphql

import (
	"reflect"

	"github.com/graphql-go/graphql"
)

// structStringResolve reads a struct field by name and returns its string value.
// Works for any source struct that has the named exported field.
func structStringResolve(field string) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (any, error) {
		if p.Source == nil {
			return nil, nil
		}
		v := reflect.ValueOf(p.Source)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return nil, nil
			}
			v = v.Elem()
		}
		fv := v.FieldByName(field)
		if !fv.IsValid() {
			return nil, nil
		}
		if fv.Kind() == reflect.String {
			if fv.String() == "" {
				return nil, nil
			}
			return fv.String(), nil
		}
		return nil, nil
	}
}

// strFieldResolve / strPtrFieldResolve / intFieldResolve are explicit-typed
// helpers for the few non-string (DateTime/Int) fields.
func strFieldResolve(field string) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (any, error) {
		s, _ := structStringResolve(field)(p)
		return s, nil
	}
}

func strPtrFieldResolve(field string) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (any, error) {
		if p.Source == nil {
			return nil, nil
		}
		v := reflect.ValueOf(p.Source)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return nil, nil
			}
			v = v.Elem()
		}
		fv := v.FieldByName(field)
		if !fv.IsValid() {
			return nil, nil
		}
		// *string
		if fv.Kind() == reflect.Ptr {
			if fv.IsNil() {
				return nil, nil
			}
			return fv.Elem().String(), nil
		}
		if fv.Kind() == reflect.String {
			if fv.String() == "" {
				return nil, nil
			}
			return fv.String(), nil
		}
		return nil, nil
	}
}

func structPtrStringResolve(field string) graphql.FieldResolveFn {
	return strPtrFieldResolve(field)
}

func structPtrIntResolve(field string) graphql.FieldResolveFn {
	return intPtrFieldResolve(field)
}

func intPtrFieldResolve(field string) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (any, error) {
		if p.Source == nil {
			return nil, nil
		}
		v := reflect.ValueOf(p.Source)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return nil, nil
			}
			v = v.Elem()
		}
		fv := v.FieldByName(field)
		if !fv.IsValid() {
			return nil, nil
		}
		switch fv.Kind() {
		case reflect.Ptr:
			if fv.IsNil() {
				return nil, nil
			}
			return int(fv.Elem().Int()), nil
		case reflect.Int, reflect.Int64:
			return int(fv.Int()), nil
		}
		return nil, nil
	}
}

func intFieldResolve(field string) graphql.FieldResolveFn { return intPtrFieldResolve(field) }

func boolFieldResolve(field string) graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (any, error) {
		if p.Source == nil {
			return nil, nil
		}
		v := reflect.ValueOf(p.Source)
		if v.Kind() == reflect.Ptr {
			if v.IsNil() {
				return nil, nil
			}
			v = v.Elem()
		}
		fv := v.FieldByName(field)
		if !fv.IsValid() {
			return false, nil
		}
		return fv.Bool(), nil
	}
}

// argsIDSlug returns the common {id, slug} argument set for singletons.
func argsIDSlug() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"id":   &graphql.ArgumentConfig{Type: graphql.ID},
		"slug": &graphql.ArgumentConfig{Type: graphql.String},
	}
}

func argSlug() graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"slug": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	}
}
