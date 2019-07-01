package todo

// Store represents and kind of store you can use as long as it implements this interface.
type Store interface {
	// Create takes the content as a parameter and return a created Task.
	Create(content string) (*Task, error)

	// Get returns a task
	Get(ID string) (*Task, error)

	// MarkAsDone update a task to done
	MarkAsDone(ID string) (*Task, error)

	// All return a TaskQuerier for all tasks
	All() TaskQuerier

	// Done return a TaskQuerier for done tasks
	Done() TaskQuerier

	// NotDone return a TaskQuerier for tasks to be done
	NotDone() TaskQuerier
}

// TaskQuerier is a handy way to count and query at the same time
type TaskQuerier interface {
	// Count returns the number of row related to the query
	Count(int, error)

	// List returns a list of element from the index with a limited number of element
	List(from, limit int) ([]Task, error)
}
