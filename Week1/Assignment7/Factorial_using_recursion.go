package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter a factorial number")
	fmt.Scan(&number)

	factorial(number) 
	fmt.Println("factorial number of", number, "=", factorial(number))

}

func factorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	} else {
		return number * factorial(number-1)

	}

}
