package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks []string

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("To-Do List Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. Remove Task")
		fmt.Println("3. View Tasks")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask(reader)
		case "2":
			removeTask(reader)
		case "3":
			viewTasks()
		case "4":
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)
	if task == "" {
		fmt.Println("Task cannot be empty.")
		return
	}
	tasks = append(tasks, task)
	fmt.Println("Task added successfully!")
}
func removeTask(reader *bufio.Reader) {
	viewTasks()
	if len(tasks) == 0 {
		return
	}
	fmt.Print("Enter task number to remove: ")

	var index int
	_, err := fmt.Scanln(&index)
	if err != nil || index < 1 || index > len(tasks) {
		fmt.Println("Invalid input. Please enter a valid task number.")
		return
	}

	tasks = append(tasks[:index-1], tasks[index:]...)
	fmt.Println("Task removed successfully!")
}
func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("\nYour Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}
