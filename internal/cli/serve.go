package cli

import (
	"context"

	"github.com/go-go-golems/fake-cms/internal/repo"
	"github.com/go-go-golems/fake-cms/internal/server"
	"github.com/go-go-golems/glazed/pkg/cli"
	"github.com/go-go-golems/glazed/pkg/cmds"
	"github.com/go-go-golems/glazed/pkg/cmds/fields"
	"github.com/go-go-golems/glazed/pkg/cmds/schema"
	"github.com/go-go-golems/glazed/pkg/cmds/values"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"net/http"
)

// serveCmd opens the SQLite DB, builds the GraphQL schema, and serves the
// HTTP API until interrupted.
type serveCmd struct{ *cmds.CommandDescription }

func newServeCmd() (*serveCmd, error) {
	httpSection, err := schema.NewSection("http", "HTTP",
		schema.WithFields(
			fields.New("addr", fields.TypeString,
				fields.WithDefault(":8080"),
				fields.WithHelp("Listen address for the GraphQL server")),
		),
	)
	if err != nil {
		return nil, err
	}
	return &serveCmd{
		CommandDescription: cmds.NewCommandDescription(
			"serve",
			cmds.WithShort("Serve the fake CMS GraphQL API over HTTP"),
			cmds.WithSections(DBSection(), httpSection),
		),
	}, nil
}

func (c *serveCmd) Run(ctx context.Context, v *values.Values) error {
	cfg, err := decodeDB(v)
	if err != nil {
		return err
	}
	var httpCfg struct {
		Addr string `glazed:"addr"`
	}
	_ = v.DecodeSectionInto("http", &httpCfg) // optional; defaults below
	addr := httpCfg.Addr
	if addr == "" {
		addr = ":8080"
	}

	r, err := repo.Open(ctx, cfg.Path)
	if err != nil {
		return errors.Wrap(err, "open db")
	}
	defer r.Close()

	handler, err := server.New(ctx, r)
	if err != nil {
		return err
	}
	srv := &http.Server{Addr: addr, Handler: handler}
	go func() {
		<-ctx.Done()
		_ = srv.Shutdown(context.Background())
	}()
	serveLog.Printf("listening on %s (playground at http://localhost%s/playground)", addr, addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

var _ cmds.BareCommand = (*serveCmd)(nil)

// registerServe wires the `serve` command.
func registerServe(root *cobra.Command) error {
	cmd, err := newServeCmd()
	if err != nil {
		return err
	}
	cobraCmd, err := cli.BuildCobraCommand(cmd)
	if err != nil {
		return err
	}
	root.AddCommand(cobraCmd)
	return nil
}
