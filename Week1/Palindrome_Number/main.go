package main

import "fmt"

func palindrome(numb int) bool {

	orginal := numb
	reverse := 0
	for numb > 0 {
		reminder := numb % 10
		reverse = reverse*10 + reminder
		numb = numb / 10

	}
	return orginal == reverse

}
func main() {
	var numb int
	fmt.Println("Enter a number to check if the number is Palindrome:")
	fmt.Scan(&numb)

	if palindrome(numb) {
		fmt.Println("it is a Palindrome number")
	} else {
		fmt.Println("it is not a Palindrome number")
	}

}
