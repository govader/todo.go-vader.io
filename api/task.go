package todo

import (
	"time"
)

// Task is a domain entity.
// It represent a simple task that is either todo or done.
type Task struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
