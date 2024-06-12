package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newAddCommand(app *App) *cobra.Command {
	var description string
	var tags []string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Creates and adds new task in the list",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				return
			}

			//newTask := args[0]

			fmt.Println("tags", tags)

			//app.State.Tasks.Add(newTask, description, []string{})

			if err := app.State.Save(".tmt.json"); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	cmd.Flags().StringVarP(&description, "description", "d", "", "Description for new task")
	cmd.Flags().StringSliceVarP(&tags, "tags", "", []string{}, "Tags for new task")

	return cmd
}
