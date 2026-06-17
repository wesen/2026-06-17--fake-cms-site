// Package server exposes the GraphQL Content Delivery API over net/http.
// It is deliberately stdlib-only (no third-party router): the org's convention
// (see glazed/pkg/web/static.go) is net/http with http.ServeMux.
package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	cmsgraphql "github.com/go-go-golems/fake-cms/internal/graphql"
	"github.com/go-go-golems/fake-cms/internal/repo"
	gql "github.com/graphql-go/graphql"
)

// graphqlRequest is the minimal POST body the handler accepts.
type graphqlRequest struct {
	Query         string         `json:"query"`
	Variables     map[string]any `json:"variables,omitempty"`
	OperationName string         `json:"operationName,omitempty"`
}

// New builds the HTTP handler serving /graphql, /playground, /healthz.
// The schema is built once at startup.
func New(ctx context.Context, r *repo.Repo) (http.Handler, error) {
	schema, err := cmsgraphql.Schema(r)
	if err != nil {
		return nil, fmt.Errorf("build graphql schema: %w", err)
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost && req.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read body: "+err.Error(), http.StatusBadRequest)
			return
		}
		var gq graphqlRequest
		if len(body) > 0 {
			if err := json.Unmarshal(body, &gq); err != nil {
				http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			// GET with ?query= (GraphiQL-style)
			gq.Query = req.URL.Query().Get("query")
		}
		if gq.Query == "" {
			http.Error(w, "missing query", http.StatusBadRequest)
			return
		}
		result := gql.Do(gql.Params{
			Schema:         schema,
			RequestString:  gq.Query,
			VariableValues: gq.Variables,
			OperationName:  gq.OperationName,
			Context:        ctx,
		})
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 200 even on field errors (GraphQL spec); client reads "errors".
		if err := json.NewEncoder(w).Encode(result); err != nil {
			log.Printf("encode graphql response: %v", err)
		}
	})

	mux.HandleFunc("/playground", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(graphiqlHTML))
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	// Root: helpful landing pointing at the playground + schema file.
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `<html><body>
<h1>fake-cms</h1>
<p>FAKE DATA — DO NOT EXPOSE PUBLICLY (unauthenticated read-only mock).</p>
<ul>
  <li><a href="/playground">GraphiQL playground</a></li>
  <li><code>POST /graphql</code> with a JSON <code>{"query":"..."}</code> body</li>
  <li><a href="/healthz">/healthz</a></li>
</ul>
</body></html>`)
	})

	return logging(mux), nil
}

// logging wraps the handler with a tiny request logger.
func logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Printf("%s %s", r.Method, r.URL.Path)
	})
}

// graphiqlHTML is a self-contained GraphiQL page (CDN-hosted assets) so the
// workshop has an instant query browser with no extra tooling.
const graphiqlHTML = `<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <title>fake-cms GraphiQL</title>
  <style>body { height: 100%; margin: 0; }</style>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/graphiql@3.0.10/graphiql.min.css" />
  <script src="https://cdn.jsdelivr.net/npm/react@18/umd/react.production.min.js"></script>
  <script src="https://cdn.jsdelivr.net/npm/react-dom@18/umd/react-dom.production.min.js"></script>
  <script src="https://cdn.jsdelivr.net/npm/graphiql@3.0.10/graphiql.min.js"></script>
</head>
<body>
  <div id="graphiql" style="height: 100vh;"></div>
  <script>
    const fetcher = async (params) => {
      const resp = await fetch('/graphql', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(params),
      });
      return resp.json();
    };
    ReactDOM.render(
      React.createElement(GraphiQL, { fetcher }),
      document.getElementById('graphiql'),
    );
  </script>
</body>
</html>`
