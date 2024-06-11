package cmd

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func newDeleteCommand(app *App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes the task with specific number",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				return
			}

			id := args[0]

			taskUUID, err := uuid.Parse(id)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			if err := app.State.Delete(taskUUID); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			if err := app.State.Save(".tmt.json"); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		},
	}

	return cmd
}
