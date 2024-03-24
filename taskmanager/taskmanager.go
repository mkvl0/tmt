package taskmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/google/uuid"
)

type taskItem struct {
	ID          uuid.UUID
	Title       string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TasksList map[uuid.UUID]taskItem

func (l TasksList) Add(title string) taskItem {
	t := taskItem{
		ID:          uuid.New(),
		Title:       title,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	l[t.ID] = t

	return t
}

// Complete method doesn't require pointer receiver
// but its a good practice to keep methods for the same type
// with the same receiver type
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

func (l *TasksList) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, js, 0644)
}

func (l *TasksList) Sync(filename string) error {
	file, err := ioutil.ReadFile(filename)
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
