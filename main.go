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

// command := os.Args[1]

// fmt.Println("command: ", command)
// switch command {
// case "add":
// 	file, err := os.Create("tasks.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	encoder := json.NewEncoder(file)
// 	encoder.SetIndent("", "")

// 	fmt.Println("add command")
// case "update":
// 	fmt.Println("update command")
// case "delete":
// 	fmt.Println("delete command")
// case "mark-in-progress":
// 	fmt.Println("mark-in-progress command")
// case "mark-done":
// 	fmt.Println("mark-done command")
// case "list":
// 	// - done
// 	// - todo
// 	// - in-progress
// 	fmt.Println("list command")
// }
