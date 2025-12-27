package storage

import (
	"encoding/json"
	"os"

	"github.com/wafash08/task-tracker-cli/task"
)

const filename = "tasks.json"

func LoadTasks() (task.Tasks, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return task.Tasks{}, nil // file belum ada
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks task.Tasks

	if len(data) == 0 {
		return task.Tasks{}, nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks task.Tasks) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
