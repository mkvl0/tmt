package tm

import "github.com/google/uuid"

type tag struct {
	Name string
	ID   uuid.UUID
}

func newTag(name string) tag {
	return tag{
		ID:   uuid.New(),
		Name: name,
	}
}
