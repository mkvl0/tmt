package tm

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type task struct {
	ID          uuid.UUID
	Title       string
	Description string
	tagsIDs     []string // IDs or just strings?
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

func newTask(
	title, description string,
	tagsIDs []string,
) task {
	return task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		tagsIDs:     tagsIDs,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
}

// why not slice of tasks? because O(1)
// using map over slice is more effective in terms of time complexity
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
