package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var elements []string

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("To-Add Elements")
		fmt.Println("1. Add an Element")
		fmt.Println("2. Remove an Element")
		fmt.Println("3. Display all Elements")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addAnElement(reader)
		case "2":
			removeAnElement(reader)
		case "3":
			displayAllElements()
		case "4":
			fmt.Println("Exiting... bye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
func addAnElement(reader *bufio.Reader) {
	fmt.Print("Enter an Element: ")
	element, _ := reader.ReadString('\n')
	element = strings.TrimSpace(element)
	if element == "" {
		fmt.Println("Element cannot be empty.")
		return
	}
	elements = append(elements, element)
	fmt.Println("Element added successfully!")
}
func removeAnElement(reader *bufio.Reader) {
	displayAllElements()
	if len(elements) == 0 {
		return
	}
	fmt.Print("Enter Element number to remove: ")

	var index int
	_, err := fmt.Scanln(&index)
	if err != nil || index < 1 || index > len(elements) {
		fmt.Println("Invalid input. Please enter a valid element number.")
		return
	}

	elements = append(elements[:index-1], elements[index:]...)
	fmt.Println("Task removed successfully!")
}
func displayAllElements() {
	if len(elements) == 0 {
		fmt.Println("No elements available.")
		return
	}

	fmt.Println("\nYour elements:")
	for i, element := range elements {
		fmt.Printf("%d. %s\n", i+1, element)
	}
}
