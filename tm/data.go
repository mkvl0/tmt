package tm

import (
	"encoding/json"
	"errors"
	"os"
)

type Data struct {
	Tasks TasksList
	Tags  TagsList
}

func NewData() Data {
	return Data{
		Tasks: NewTasksList(),
		Tags:  NewTagsList(),
	}
}

func (d *Data) Save(filename string) error {
	js, err := json.Marshal(d)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (d *Data) Get(filename string) error {
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

	err = json.Unmarshal(file, d)

	if err != nil {
		return err
	}

	return nil
}
