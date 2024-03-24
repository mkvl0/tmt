/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"os"
	"tmt/cmd"
	"tmt/taskmanager"
)

var tasksFileName = ".tmt.json"

func main() {
	l := &taskmanager.List{}

	if err := l.Get(tasksFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	app := cmd.NewApp(l)
	app.Execute()
}
