package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newSearchCommand(app *App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var searchPattern string

			if len(args) == 1 {
				searchPattern = args[0]
			}

			r, err := app.State.Search(searchPattern)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Print(r)
		},
	}

	return cmd
}
