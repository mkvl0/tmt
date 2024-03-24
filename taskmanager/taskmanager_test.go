package taskmanager_test

import (
	"io/ioutil"
	"os"
	"testing"
	"tmt/taskmanager"
)

func TestAdd(t *testing.T) {
	l := taskmanager.List{}

	newTask := "New task"
	l.Add(newTask)

	if len(l) == 0 {
		t.Errorf("Expected length of 'l' %d but got %d", 1, len(l))
	}

	if l[0].Task != newTask {
		t.Errorf("Expected %q but got %q", newTask, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := taskmanager.List{}

	l.Add("New task")

	if l[0].Done {
		t.Errorf("Task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("Task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := taskmanager.List{}
	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected length of list %d but got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q but got %q", tasks[2], l[1].Task)
	}
}

// tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	l1 := taskmanager.List{}
	l2 := taskmanager.List{}

	taskName := "New Task"
	l1.Add(taskName)

	// TempFile creates a new temporary file in the directory dir (first parameter)
	tf, err := ioutil.TempFile("", "")
	defer os.Remove(tf.Name()) // will be executed in the end of the function

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}
}
