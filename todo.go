package todo

import (
	"fmt"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAT time.Time
}

type List []item

// Add creates  a new todo item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAT: time.Time{},
	}

	*l = append(*l, t)
}

//Complete method marks a ToDo item as completed by
// setting Done = true and CompletedAt to the current time

func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}
	// Adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompletedAT = time.Now()

	return nil
}
