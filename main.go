package main

import (
	"fmt"
	"os"
	"tmt/cmd"
	"tmt/tm"
)

var tasksFileName = ".tmt.json"

func main() {
	d := tm.NewData()

	if err := d.Get(tasksFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	app := cmd.NewApp(&d)
	app.Execute()
}
