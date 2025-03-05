package main

import "fmt"

func addition(number1, number2 int) {
	fmt.Println(number1 + number2)
}
func subtraction(number1, number2 int) {
	fmt.Println(number1 - number2)
}
func multiplication(number1, number2 int) {
	fmt.Println(number1 * number2)
}
func division(number1, number2 int) {
	fmt.Println(number1 / number2)
}

func main() {
	var number1 int
	fmt.Println("Enter first number")
	fmt.Scan(&number1)

	var number2 int
	fmt.Println("Enter second number")
	fmt.Scan(&number2)

	var operator string
	fmt.Println("Enter operator")
	fmt.Scan(&operator)

	if operator == "+" {
		addition(number1, number2)
	} else if operator == "-" {
		subtraction(number1, number2)
	} else if operator == "*" {
		multiplication(number1, number2)
	} else if operator == "/" {
		division(number1, number2)
	} else {
		fmt.Println("operator not available, use only +,-,*,/")
	}
}
