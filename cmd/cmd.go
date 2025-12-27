package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wafash08/task-tracker-cli/storage"
	"github.com/wafash08/task-tracker-cli/task"
)

func Add(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: task description is required")
		fmt.Println(`Usage: task add "task description"`)
		os.Exit(1)
	}

	description := strings.Join(args, " ")

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	task, err := tasks.Add(description)
	if err != nil {
		fmt.Println("Error adding task:", err)
		os.Exit(1)
	}

	if err := storage.SaveTasks(tasks); err != nil {
		fmt.Println("Error saving tasks:", err)
		os.Exit(1)
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.Id)
}

func Delete(arg string) {
	if arg == "" {
		fmt.Println("Error: task id is required")
		fmt.Println(`Usage: task delete 1`)
		os.Exit(1)
	}

	id, err := strconv.Atoi(arg)
	if err != nil || id <= 0 {
		fmt.Printf("Invalid task id: %s\n", arg)
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	err = tasks.Delete(id)
	if err != nil {
		if err == task.ErrTaskNotFound {
			fmt.Printf("Task with ID %d not found\n", id)
			return
		}
		fmt.Println("Error deleting task:", err)
		return
	}

	err = storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task deleted successfully (ID: %d)\n", id)
}

func Update(args []string) {
	if len(args) <= 1 {
		fmt.Println("Error: task id and description are required")
		fmt.Println(`Usage: task update "new description"`)
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil || id <= 0 {
		fmt.Printf("Invalid task id: %s\n", args[0])
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	description := args[1]
	err = tasks.Update(id, description)
	if err != nil {
		if err == task.ErrTaskNotFound {
			fmt.Printf("Task with ID %d not found\n", id)
			return
		}
		fmt.Println("Error updating task:", err)
		return
	}

	err = storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task updated successfully (ID: %d)\n", id)
}

func MarkInProgress(arg string) {
	if arg == "" {
		fmt.Println("Error: task id is required")
		fmt.Println(`Usage: task delete 1`)
		os.Exit(1)
	}

	id, err := strconv.Atoi(arg)
	if err != nil || id <= 0 {
		fmt.Printf("Invalid task id: %s\n", arg)
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	err = tasks.MarkInProgress(id)
	if err != nil {
		if err == task.ErrTaskNotFound {
			fmt.Printf("Task with ID %d not found\n", id)
			return
		}
		fmt.Println("Error marking in progress a task:", err)
		return
	}

	err = storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Status task with (ID: %d) has been successfully updated to in progress\n", id)
}

func MarkDone(arg string) {
	if arg == "" {
		fmt.Println("Error: task id is required")
		fmt.Println(`Usage: task delete 1`)
		os.Exit(1)
	}

	id, err := strconv.Atoi(arg)
	if err != nil || id <= 0 {
		fmt.Printf("Invalid task id: %s\n", arg)
		return
	}

	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	err = tasks.MarkDone(id)
	if err != nil {
		if err == task.ErrTaskNotFound {
			fmt.Printf("Task with ID %d not found\n", id)
			return
		}
		fmt.Println("Error marking done a task:", err)
		return
	}

	err = storage.SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Status task with (ID: %d) has been successfully updated to done\n", id)
}

func List(args []string) {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	// case: task-cli list
	if len(args) == 0 {
		tasks.PrintTasks()
		return
	}

	if len(args) == 1 {
		status, err := ParseStatus(args[0])
		if err != nil {
			fmt.Println(err)
			fmt.Println("Valid statuses: todo, in-progress, done")
			return
		}

		filtered := tasks.ListByStatus(status)
		if len(filtered) == 0 {
			fmt.Printf("No tasks with status %s\n", status)
			return
		}

		filtered.PrintTasks()
		return
	}

	fmt.Println("Usage:")
	fmt.Println("  task-cli list")
	fmt.Println("  task-cli list <todo|in-progress|done>")
}
