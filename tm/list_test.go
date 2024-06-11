package tm_test

import (
	"os"
	"testing"
	"tmt/tm"

	"github.com/google/uuid"
)

func TestAdd(t *testing.T) {
	l := tm.NewTasksList()

	newTask := "New task"
	got := l.Add(newTask, "", []string{})

	if len(l) == 0 {
		t.Errorf("Expected length of 'l' %d but got %d", 1, len(l))
	}

	if got.Title != newTask {
		t.Errorf("Expected %q but got %q", newTask, got.Title)
	}
}

func TestComplete(t *testing.T) {
	l := tm.NewTasksList()

	newTask := l.Add("New task", "", []string{})

	if newTask.Done {
		t.Errorf("Task should not be completed")
	}

	l.Complete(newTask.ID)

	completedTask := l[newTask.ID]

	if !completedTask.Done {
		t.Errorf("Task should be completed")
	}
}

func TestDelete(t *testing.T) {
	l := tm.NewTasksList()
	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}

	newTaskIds := []uuid.UUID{}

	for _, task := range tasks {
		t := l.Add(task, "", []string{})
		newTaskIds = append(newTaskIds, t.ID)
	}

	l.Delete(newTaskIds[1])

	if len(l) != 2 {
		t.Errorf("Expected length of list %d but got %d", 2, len(l))
	}

	_, ok := l[newTaskIds[1]]

	if ok {
		t.Errorf("Expected %t but got %t", false, ok)
	}
}

// tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	l1 := tm.TasksList{}
	l2 := tm.TasksList{}

	taskName := "New Task"
	l1.Add(taskName, "", []string{})

	// TempFile creates a new temporary file in the directory dir (first parameter)
	tf, err := os.CreateTemp("", "")
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
