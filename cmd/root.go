/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"tmt/taskmanager"

	"github.com/spf13/cobra"
)

type App struct {
	rootCmd *cobra.Command
	State   *taskmanager.List
}

func NewApp(state *taskmanager.List) *App {
	app := &App{
		State: state,
	}

	app.rootCmd = &cobra.Command{
		Use:   "tmt",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
  examples and usage of using your application. For example:

  Cobra is a CLI library for Go that empowers applications.
  This application is a tool to generate the needed files
  to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	app.rootCmd.AddCommand(newAddCommand(app))
	app.rootCmd.AddCommand(newListCommand(app))
	app.rootCmd.AddCommand(newDeleteCommand(app))
	app.rootCmd.AddCommand(newCompleteCommand(app))

	return app
}

func (app *App) Execute() {
	if err := app.rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
