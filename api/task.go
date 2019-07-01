package todo

import (
	"time"
)

// Task is a domain entity.
// It represent a simple task that is either todo or done.
type Task struct {
	ID        string
	Content   string
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
