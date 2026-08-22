package cli

import (
	"fmt"
	"io"
	"sort"

	"github.com/spf13/cobra"

	"github.com/green-threads/ndo/internal/config"
)

func newListCmd(deps Deps) *cobra.Command {
	var local, central bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List resolved recipes",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(deps, cmd.OutOrStdout(), local, central)
		},
	}

	cmd.Flags().BoolVar(&local, "local", false, "show only the local file's recipes, unmerged")
	cmd.Flags().BoolVar(&central, "central", false, "show only the central file's recipes, unmerged")

	return cmd
}

func runList(deps Deps, stdout io.Writer, local, central bool) error {
	res, err := loadResolved(deps)
	if err != nil {
		return err
	}

	rf := res.Merged
	switch {
	case local && central:
		return fmt.Errorf("--local and --central are mutually exclusive")
	case local:
		rf = res.Local
	case central:
		rf = res.Central
	}

	names := make([]string, 0, len(rf.Recipes))
	for name := range rf.Recipes {
		names = append(names, name)
	}
	sort.Strings(names)

	settings, err := config.LoadSettings(deps.NDOHome)
	if err != nil {
		return err
	}
	styler := newStyler(stdout, settings)

	for _, name := range names {
		r := rf.Recipes[name]
		if r.Description != "" {
			fmt.Fprintf(stdout, "%s\t%s\t%s\n", styler.BoldCyan(name), r.Command, styler.Dim(r.Description))
		} else {
			fmt.Fprintf(stdout, "%s\t%s\n", styler.BoldCyan(name), r.Command)
		}
	}
	return nil
}
