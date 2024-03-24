package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func newCompleteCommand(app *App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "complete",
		Short: "Completes the task with specific id",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				return
			}

			id, err := strconv.Atoi(args[0])

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			if err := app.State.Complete(id); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			if err := app.State.Save(".tmt.json"); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
