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
	var addTask bool
	var list bool
	var completeTask int
	var deleteTask int
	flag.BoolVar(&addTask, "add", false, "Add new task in the list")
	flag.BoolVar(&list, "list", false, "List all tasks")
	flag.IntVar(&completeTask, "complete", 0, "Task to complete")
	flag.IntVar(&deleteTask, "delete", 0, "Task to delete")
	flag.Parse()

	l := &tmt.List{}
	if err := l.Get(tasksFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case list:
		handleShowList(l)
	case completeTask > 0:
		handleCompleteTask(completeTask, l)
	case addTask:
		handleAddTask(l)
	case deleteTask > 0:
		handleDeleteTask(deleteTask, l)
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

func handleCompleteTask(id int, l *tmt.List) {
	if err := l.Complete(id); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := l.Save(tasksFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func handleShowList(l *tmt.List) {
	// list current items to do
	fmt.Print(l)
}

func handleAddTask(l *tmt.List) {
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
}

func handleDeleteTask(id int, l *tmt.List) {
	if err := l.Delete(id); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	if err := l.Save(tasksFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
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
