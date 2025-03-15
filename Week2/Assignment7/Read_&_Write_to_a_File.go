package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	
	fmt.Print("Enter the input file name: ")
	var inputFile string
	fmt.Scanln(&inputFile)

	
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Print("Enter the word to replace: ")
	reader := bufio.NewReader(os.Stdin)
	oldWord, _ := reader.ReadString('\n')
	oldWord = strings.TrimSpace(oldWord) 

	fmt.Print("Enter the new word: ")
	newWord, _ := reader.ReadString('\n')
	newWord = strings.TrimSpace(newWord) 

	
	modifiedContent := strings.ReplaceAll(string(content), oldWord, newWord)

	
	outputFile := "output.txt"
	err = os.WriteFile(outputFile, []byte(modifiedContent), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	
	fmt.Println("File has been modified and saved as 'output.txt'.")
}
