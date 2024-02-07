package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"tmt"
)

var tasksFileName = ".tmt.json"

func main() {
	if fileNameFromEnv := os.Getenv("TMT_FILENAME_PATH"); fileNameFromEnv != "" {
		tasksFileName = fileNameFromEnv
	}
	flag.Usage = func() {
		output := flag.CommandLine.Output()
		fmt.Fprintln(output, "tmt (task management tool)")
		fmt.Fprintln(output, "Usage information: ")
		flag.PrintDefaults()
	}
	add := flag.Bool("add", false, "Add new task in the list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Task to be completed")
	flag.Parse()

	l := &tmt.List{}
	if err := l.Get(tasksFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		// list current items to do
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(tasksFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		newTask, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		l.Add(newTask)
		if err := l.Save(tasksFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()

	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return s.Text(), nil
}
