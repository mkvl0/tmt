package tm_test

import (
	"testing"
	"tmt/tm"

	"github.com/google/uuid"
)

func TestAddTask(t *testing.T) {
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

func TestCompleteTask(t *testing.T) {
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

func TestDeleteTask(t *testing.T) {
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
