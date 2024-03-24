package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newAddCommand(app *App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Creates and adds new task in the list",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				return
			}

			newTask := args[0]

			app.State.Add(newTask)

			if err := app.State.Save(".tmt.json"); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
