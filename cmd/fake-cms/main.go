// Command fake-cms is the operator CLI for the fake CMS GraphQL mock.
package main

import (
	"fmt"
	"os"

	"github.com/go-go-golems/fake-cms/internal/cli"
)

func main() {
	root, err := cli.BuildRootCobra()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	if err := root.Execute(); err != nil {
		// cobra prints the error itself; exit non-zero.
		os.Exit(1)
	}
}
