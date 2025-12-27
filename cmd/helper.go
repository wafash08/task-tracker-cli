package cmd

import (
	"fmt"

	"github.com/wafash08/task-tracker-cli/task"
)

func ParseStatus(input string) (string, error) {
	switch input {
	case task.Done.String():
		return task.Done.String(), nil
	case task.InProgress.String():
		return task.InProgress.String(), nil
	case task.Todo.String():
		return task.Todo.String(), nil
	default:
		return "", fmt.Errorf("invalid status: %s", input)
	}
}
