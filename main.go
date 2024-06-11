package main

import (
	"fmt"
	"os"
	"tmt/cmd"
	"tmt/tm"
)

var tasksFileName = ".tmt.json"

func main() {
	l := tm.NewTasksList()

	if err := l.Get(tasksFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	app := cmd.NewApp(&l)
	app.Execute()
}
