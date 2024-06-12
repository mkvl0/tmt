package tm

import (
	"errors"

	"github.com/google/uuid"
)

var ErrNoTag = errors.New("Tag with such ID does not exist")
var ErrNotUniqueTag = errors.New("Tag with such name already exists")

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

type TagsList map[uuid.UUID]tag

func NewTagsList() TagsList {
	return make(TagsList)
}

func (l TagsList) Add(
	name string,
) (tag, error) {
	for _, v := range l {
		if v.Name == name {
			return tag{}, ErrNotUniqueTag
		}
	}

	t := newTag(
		name,
	)

	l[t.ID] = t

	return t, nil
}

func (l TagsList) Delete(ID uuid.UUID) error {
	_, ok := l[ID]

	if !ok {
		return ErrNoTag
	}

	delete(l, ID)

	return nil
}

func (l TagsList) Get(ID uuid.UUID) (tag, error) {
	t, ok := l[ID]

	if !ok {
		return tag{}, ErrNoTag
	}

	return t, nil
}

func (l TagsList) Update(ID uuid.UUID, name string) (tag, error) {
	t, err := l.Get(ID)

	if err != nil {
		return tag{}, err
	}

	t.Name = name
	l[ID] = t

	return t, nil
}
