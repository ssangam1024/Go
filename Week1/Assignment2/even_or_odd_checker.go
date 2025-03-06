package main

import "fmt"

func main() {
	var given_number int
	fmt.Println("Enter a Number")
	fmt.Scan(&given_number)

	var x int
	x = given_number % 2

	if x == 0 {
		fmt.Println("is an even number")
	} else {
		fmt.Println("is an odd number")
	}
}
