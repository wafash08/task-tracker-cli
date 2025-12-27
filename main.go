package main

import (
	"fmt"
	"os"

	"github.com/wafash08/task-tracker-cli/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  task add \"task description\"")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		cmd.Add(args)
	case "delete":
		cmd.Delete(args[0])
	case "update":
		cmd.Update(args)
	case "mark-in-progress":
		cmd.MarkInProgress(args[0])
	case "mark-done":
		cmd.MarkDone(args[0])
	case "list":
		cmd.List(args)
	default:
		fmt.Printf("Unknown command: %s\n\n", command)
		fmt.Println("Usage:")
		fmt.Println("  task add \"task description\"")
	}
}
