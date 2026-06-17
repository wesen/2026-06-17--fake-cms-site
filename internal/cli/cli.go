// Package cli wires the fake-cms operator surface (serve / seed / db / version)
// as glazed commands and builds their cobra commands.
package cli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-go-golems/fake-cms/internal/build"
	"github.com/go-go-golems/fake-cms/internal/doc"
	"github.com/go-go-golems/glazed/pkg/cli"
	"github.com/go-go-golems/glazed/pkg/cmds"
	"github.com/go-go-golems/glazed/pkg/cmds/fields"
	"github.com/go-go-golems/glazed/pkg/cmds/schema"
	"github.com/go-go-golems/glazed/pkg/cmds/values"
	helpcmd "github.com/go-go-golems/glazed/pkg/help/cmd"
	"github.com/go-go-golems/glazed/pkg/help"
	"github.com/spf13/cobra"
)

// serveLog writes to stderr so it never collides with structured stdout.
var serveLog = log.New(os.Stderr, "fake-cms ", log.LstdFlags)

// dbConfig is decoded from the "db" section by decodeDB.
type dbConfig struct {
	Path     string `glazed:"path"`
	Recreate bool   `glazed:"recreate"`
}

// decodeDB extracts the db section values.
func decodeDB(v *values.Values) (dbConfig, error) {
	cfg := dbConfig{Path: "cms.db"}
	if err := v.DecodeSectionInto("db", &cfg); err != nil {
		return cfg, err
	}
	if cfg.Path == "" {
		cfg.Path = "cms.db"
	}
	return cfg, nil
}

// DBSection is the shared "database" field section (db path + migrations toggle).
// It is attached to every subcommand that needs to open SQLite.
func DBSection() *schema.SectionImpl {
	s, err := schema.NewSection("db", "Database",
		schema.WithFields(
			fields.New("path", fields.TypeString,
				fields.WithDefault("cms.db"),
				fields.WithHelp("Path to the SQLite file (created/migrated if missing)")),
			fields.New("recreate", fields.TypeBool,
				fields.WithDefault(false),
				fields.WithHelp("Drop and recreate the DB before use (seed)")),
		),
	)
	if err != nil {
		// WithFields/WithDefault never fail at init time; this is defensive.
		panic(err)
	}
	return s
}

// HTTPSection holds the HTTP server address (for the serve command).
func HTTPSection() *schema.SectionImpl {
	s, err := schema.NewSection("http", "HTTP",
		schema.WithFields(
			fields.New("addr", fields.TypeString,
				fields.WithDefault(":8080"),
				fields.WithHelp("Listen address for the GraphQL server")),
		),
	)
	if err != nil {
		panic(err)
	}
	return s
}

// versionCmd is a trivial BareCommand used to verify the glazed wiring works.
type versionCmd struct{ *cmds.CommandDescription }

func newVersionCmd() *versionCmd {
	return &versionCmd{
		CommandDescription: cmds.NewCommandDescription(
			"version",
			cmds.WithShort("Print fake-cms version"),
		),
	}
}

func (c *versionCmd) Run(_ context.Context, _ *values.Values) error {
	fmt.Printf("fake-cms %s (commit %s)\n", build.Version, build.Commit)
	return nil
}

var _ cmds.BareCommand = (*versionCmd)(nil)

// BuildRootCobra assembles the root command with all subcommands wired via
// glazed. Subcommands are added in later phases; version is always present so
// the CLI is testable from Phase 0.
func BuildRootCobra() (*cobra.Command, error) {
	root := &cobra.Command{
		Use:   "fake-cms",
		Short: "Fake internal CMS GraphQL API (SQLite-backed) for the SSG workshop",
	}

	// Wire the glazed help system: load embedded entries, attach the `help`
	// command and the rich help/usage templates to the root.
	if err := setupHelp(root); err != nil {
		return nil, fmt.Errorf("setup help: %w", err)
	}

	version, err := cli.BuildCobraCommand(newVersionCmd())
	if err != nil {
		return nil, fmt.Errorf("build version command: %w", err)
	}
	root.AddCommand(version)

	// serve / seed / db are registered here in their respective phases.
	if err := registerServe(root); err != nil {
		return nil, err
	}
	if err := registerSeed(root); err != nil {
		return nil, err
	}
	if err := registerDB(root); err != nil {
		return nil, err
	}

	return root, nil
}

// setupHelp builds the glazed HelpSystem from the embedded doc/*.md entries
// and attaches it to the root command (the `help` subcommand, rich help/usage
// templates). This is the canonical glazed initialization path.
func setupHelp(root *cobra.Command) error {
	hs := help.NewHelpSystem()
	if err := hs.LoadSectionsFromFS(doc.FS, "."); err != nil {
		return err
	}
	helpcmd.SetupCobraRootCommand(hs, root)
	return nil
}
