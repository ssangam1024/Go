package main

import "fmt"

func main() {
	var x float64 = 1.7e+308                // declaring x variable with data type float64 and value 1.7e+308
	fmt.Printf("Type: %T, value: %v", x, x) //this function printf() first formats its argument with formating verb given %T and then prints given x data type and formating verb given %v and then prints given x value.
}
