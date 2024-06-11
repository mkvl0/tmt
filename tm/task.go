package tm

import (
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
