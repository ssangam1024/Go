package main

import "fmt"

func palindrome(numb int)int {

	
	
	if palindrome == 0 {
		fmt.Println("it is a Palindrome number")
	} else {
		fmt.Println("it is not a Palindrome number")
	}

}

func main() {
	var numb int
	fmt.Println("Enter a number to check if the number is Palindrome:")
	fmt.Scan(&numb)

	palindrome(numb)
}
