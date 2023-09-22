package todo_test

import (
	"os"
	"testing"
	"todo"
)

// TestAdd test the Add Method of the List type

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %s, got %s instead.", taskName, l[0].Task)
	}
}

// TestComplete tests the Complete method of the List Type
func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Done {
		t.Errorf("New task should not be completed.")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed")
	}
}

// TestDelete tests the Delete method of the List type
func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}

	for _, v := range tasks {
		l.Add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %s, got %s instead", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %s, got %s instead.", tasks[2], l[1].Task)
	}
}

// TestSaveGet tests the save and Get method of the List type
func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	taskName := "New Task"
	l1.Add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %s, got %s instead.", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Errorf("Error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %s should match %s task.", l1[0].Task, l2[0].Task)
	}

}
