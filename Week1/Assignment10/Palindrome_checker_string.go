package main

import (
	"fmt"
	"strings"
)

func isPalindrome(word string) bool {

	word = strings.ToLower(word)
	reversed := ""
	for i := len(word) - 1; i >= 0; i-- {
		reversed += string(word[i])
	}
	return word == reversed
}

func main() {
	var word string
	fmt.Print("Enter a word: ")
	fmt.Scan(&word)

	if isPalindrome(word) {
		fmt.Println("It is a Palindrome word.")
	} else {
		fmt.Println("It is not a Palindrome word.")
	}
}
