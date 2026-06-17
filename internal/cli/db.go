package cli

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-go-golems/fake-cms/internal/repo"
	"github.com/go-go-golems/glazed/pkg/cli"
	"github.com/go-go-golems/glazed/pkg/cmds"
	"github.com/go-go-golems/glazed/pkg/cmds/fields"
	"github.com/go-go-golems/glazed/pkg/cmds/schema"
	"github.com/go-go-golems/glazed/pkg/cmds/values"
	"github.com/spf13/cobra"
)

// dbQueryCmd runs an ad-hoc SQL read against the DB and prints rows. It is the
// glazed-structured-output escape hatch for poking at the seed.
type dbQueryCmd struct{ *cmds.CommandDescription }

func newDBQueryCmd() (*dbQueryCmd, error) {
	sqlSection, err := schema.NewSection("sql", "SQL",
		schema.WithFields(
			fields.New("query", fields.TypeString,
				fields.WithHelp("SQL SELECT to run (read-only)")),
		),
	)
	if err != nil {
		return nil, err
	}
	return &dbQueryCmd{
		CommandDescription: cmds.NewCommandDescription(
			"query",
			cmds.WithShort("Run a read-only SQL query against the DB and print rows"),
			cmds.WithSections(DBSection(), sqlSection),
		),
	}, nil
}

func (c *dbQueryCmd) Run(ctx context.Context, v *values.Values) error {
	cfg, err := decodeDB(v)
	if err != nil {
		return err
	}
	var sqlCfg struct {
		Query string `glazed:"query"`
	}
	if err := v.DecodeSectionInto("sql", &sqlCfg); err != nil {
		return err
	}
	if sqlCfg.Query == "" {
		return fmt.Errorf("sql.query is required")
	}

	r, err := repo.Open(ctx, cfg.Path)
	if err != nil {
		return err
	}
	defer r.Close()

	rows, err := r.DB.QueryContext(ctx, sqlCfg.Query)
	if err != nil {
		return err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	// header
	for i, c := range cols {
		if i > 0 {
			fmt.Print("\t")
		}
		fmt.Print(c)
	}
	fmt.Println()
	n := 0
	for rows.Next() {
		vals := make([]sql.NullString, len(cols))
		ptrs := make([]any, len(cols))
		for i := range vals {
			ptrs[i] = &vals[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			return err
		}
		for i, vv := range vals {
			if i > 0 {
				fmt.Print("\t")
			}
			if vv.Valid {
				fmt.Print(vv.String)
			}
		}
		fmt.Println()
		n++
	}
	fmt.Printf("(%d rows)\n", n)
	return rows.Err()
}

var _ cmds.BareCommand = (*dbQueryCmd)(nil)

// registerDB wires the `db` command group.
func registerDB(root *cobra.Command) error {
	db := &cobra.Command{
		Use:   "db",
		Short: "Database utilities",
	}
	query, err := newDBQueryCmd()
	if err != nil {
		return err
	}
	queryCobra, err := cli.BuildCobraCommand(query)
	if err != nil {
		return err
	}
	db.AddCommand(queryCobra)
	root.AddCommand(db)
	return nil
}
