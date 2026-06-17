package cli

import (
	"context"
	"fmt"

	"github.com/go-go-golems/fake-cms/internal/repo"
	"github.com/go-go-golems/glazed/pkg/cli"
	"github.com/go-go-golems/glazed/pkg/cmds"
	"github.com/go-go-golems/glazed/pkg/cmds/values"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// seedCmd (re)creates and seeds the SQLite file from the deterministic
// generator. Useful to bootstrap testdata/cms.db or a fresh workshop DB.
type seedCmd struct{ *cmds.CommandDescription }

func newSeedCmd() (*seedCmd, error) {
	return &seedCmd{
		CommandDescription: cmds.NewCommandDescription(
			"seed",
			cmds.WithShort("Create/migrate + seed the SQLite DB (deterministic)"),
			cmds.WithSections(DBSection()),
		),
	}, nil
}

func (c *seedCmd) Run(ctx context.Context, v *values.Values) error {
	cfg, err := decodeDB(v)
	if err != nil {
		return err
	}
	r, err := repo.Open(ctx, cfg.Path)
	if err != nil {
		return errors.Wrap(err, "open db")
	}
	defer r.Close()
	// Seed always recreates so the artifact is byte-stable regardless of
	// whatever was in the file before.
	if err := r.Seed(ctx); err != nil {
		return errors.Wrap(err, "seed")
	}
	var n int
	_ = r.DB.QueryRowContext(ctx, "SELECT count(*) FROM content_node").Scan(&n)
	fmt.Printf("seeded %s with %d content nodes\n", cfg.Path, n)
	return nil
}

var _ cmds.BareCommand = (*seedCmd)(nil)

// registerSeed wires the `seed` command.
func registerSeed(root *cobra.Command) error {
	cmd, err := newSeedCmd()
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
