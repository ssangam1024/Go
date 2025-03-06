package main

import "fmt"

func reverse_number(numb int) int {

	reverse := 0
	for numb > 0 {
		reminder := numb % 10
		reverse = reverse*10 + reminder
		numb = numb / 10

	}
	fmt.Println(reverse)
	return reverse
}
func main() {
	var numb int
	fmt.Println("enter numbers")
	fmt.Scan(&numb)

	reverse_number(numb)
}
