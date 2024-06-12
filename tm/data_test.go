package tm_test

import (
	"os"
	"testing"
	"tmt/tm"
)

// tests the Save and Get methods of the List type
func TestSaveGet(t *testing.T) {
	d := tm.NewData()

	taskName := "New Task"
	d.Tasks.Add(taskName, "", []string{})

	// TempFile creates a new temporary file in the directory dir (first parameter)
	tf, err := os.CreateTemp("", "")
	defer os.Remove(tf.Name()) // will be executed in the end of the function

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	if err := d.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := d.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}
}
