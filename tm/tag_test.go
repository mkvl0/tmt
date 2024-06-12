package tm_test

import (
	"errors"
	"testing"
	"tmt/tm"
)

func TestAddTag(t *testing.T) {
	l := tm.NewTagsList()

	newTag := "something"
	got, err := l.Add(newTag)

	if err != nil {
		t.Fatal(err)
	}

	if len(l) == 0 {
		t.Errorf("Expected length of 'l' %d but got %d", 1, len(l))
	}

	if got.Name != newTag {
		t.Errorf("Expected %q but got %q", newTag, got.Name)
	}
}

func TestAddNonUniqueTag(t *testing.T) {
	l := tm.NewTagsList()
	l.Add("programming")

	_, err := l.Add("programming")

	if !errors.Is(err, tm.ErrNotUniqueTag) {
		t.Errorf("Expected error '%v' but got '%v'", tm.ErrNotUniqueTag, err)
	}
}

func TestUpdateTag(t *testing.T) {
	l := tm.NewTagsList()

	name := "programming"
	tag, _ := l.Add(name)

	newName := "coding"

	updTag, err := l.Update(tag.ID, newName)

	if err != nil {
		t.Fatal(err)
	}

	// we can test it like that but we're not sure if it updated in the list
	// if updTag.Name != newName {
	// 	t.Errorf("Expected %q but got %q", newName, updTag.Name)
	// }

	foundTag, _ := l.Get(updTag.ID)

	if foundTag.Name != newName {
		t.Errorf("Expected %q but got %q", newName, foundTag.Name)
	}
}

func TestGetTag(t *testing.T) {
	l := tm.NewTagsList()

	tag, _ := l.Add("tag 1")
	l.Add("tag 2")

	found, err := l.Get(tag.ID)

	if err != nil {
		t.Fatal(err)
	}

	if found.ID != tag.ID {
		t.Errorf("Expected tag with ID %q but got with ID %q", tag.ID, found.ID)
	}
}

func TestDeleteTag(t *testing.T) {
	l := tm.NewTagsList()

	tag, _ := l.Add("tag 1")
	l.Add("tag 2")

	err := l.Delete(tag.ID)

	if err != nil {
		t.Fatal(err)
	}

	found, _ := l.Get(tag.ID)

	if found.ID == tag.ID {
		t.Error("Expected no tag but got one")
	}
}
