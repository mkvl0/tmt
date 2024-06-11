package tm

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type TasksList map[uuid.UUID]task

func NewTasksList() TasksList {
	return make(TasksList)
}

func (l TasksList) Add(
	title, description string,
	tagsIDs []string,
) task {
	t := newTask(
		title,
		description,
		tagsIDs,
	)

	l[t.ID] = t

	return t
}

func (l TasksList) Complete(ID uuid.UUID) error {
	t, ok := l[ID]

	if !ok {
		return fmt.Errorf("Item with ID %q does not exist", ID)
	}

	t.Done = true
	t.CompletedAt = time.Now()

	l[ID] = t

	return nil
}

func (l TasksList) Delete(ID uuid.UUID) error {
	_, ok := l[ID]

	if !ok {
		return fmt.Errorf("Item with ID %q does not exist", ID)
	}

	delete(l, ID)

	return nil
}

func (l TasksList) Search(pattern string) (TasksList, error) {
	result := NewTasksList()

	for _, t := range l {
		if !strings.Contains(t.Title, pattern) {
			continue
		}

		result[t.ID] = t
	}

	return result, nil
}

func (l *TasksList) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (l *TasksList) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	err = json.Unmarshal(file, l)

	if err != nil {
		return err
	}

	return nil
}

func (l TasksList) String() string {
	formatted := ""

	for _, t := range l {
		prefix := "  "

		if t.Done {
			prefix = "X "
		}

		// 1: First task\n
		// X 1: Second task\n
		formatted += fmt.Sprintf("%s\t%s\t%s\n", prefix, t.ID, t.Title)
	}

	return formatted
}
