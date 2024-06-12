package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newListCommand(app *App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Show list of the existing tasks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(app.State.Tasks)
		},
	}

	return cmd
}
