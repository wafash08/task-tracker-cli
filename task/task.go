package task

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

var statusName = map[Status]string{
	Todo:       "todo",
	InProgress: "in-progress",
	Done:       "done",
}

func (s Status) String() string {
	return statusName[s]
}

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type Tasks []Task

func (t Tasks) NextID() int {
	max := 0
	for _, task := range t {
		if task.Id > max {
			max = task.Id
		}
	}
	return max + 1
}

// Adding a new task
// Output: Task added successfully (ID: 1)
func (t *Tasks) Add(description string) (Task, error) {
	description = strings.TrimSpace(description)
	if description == "" {
		return Task{}, errors.New("description cannot be empty")
	}
	task := Task{
		Id:          t.NextID(),
		Description: description,
		Status:      Todo.String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	*t = append(*t, task)
	return task, nil
}

// Deleting a task
func (t *Tasks) Delete(id int) error {
	for i, task := range *t {
		if task.Id == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
			return nil
		}
	}

	return ErrTaskNotFound
}

// Updating a task
func (t *Tasks) Update(id int, description string) error {
	tasks := *t
	for i, task := range tasks {
		if task.Id == id {
			now := time.Now()
			tasks[i].Description = description
			tasks[i].UpdatedAt = &now
			return nil
		}
	}

	return ErrTaskNotFound
}

func (t *Tasks) updateStatus(id int, status string) error {
	tasks := *t
	for i, task := range tasks {
		if task.Id == id {
			now := time.Now()
			tasks[i].Status = status
			tasks[i].UpdatedAt = &now
			return nil
		}
	}

	return ErrTaskNotFound
}

// Marking a task as in-progress
func (t *Tasks) MarkInProgress(id int) error {
	return t.updateStatus(id, InProgress.String())
}

// Marking a task as done
func (t *Tasks) MarkDone(id int) error {
	return t.updateStatus(id, Done.String())
}

// Listing by status: todo, in-progress, done
func (t Tasks) ListByStatus(status string) Tasks {
	var filtered Tasks
	for _, task := range t {
		if task.Status == status {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func (t Tasks) PrintTasks() {
	for _, task := range t {
		fmt.Printf(
			"[%d] %s (%s)\n",
			task.Id,
			task.Description,
			task.Status,
		)
	}
}
