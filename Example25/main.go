package main

import "fmt"

func main() {
	var x int = 500                         // declaring x variable with data type and intial value
	var y int = -4500                       // declaring y varaiable with data type and negative value
	fmt.Printf("Type: %T, value: %v", x, x) //this function printf() first formats its argument with formating verb given %T and then prints given x data type and formating verb given %v and then prints given x value.
	fmt.Printf("Type: %T, value: %v", y, y) //this function printf() first formats its argument with formating verb given %T and then prints given y data type and formating verb given %v and then prints given y value.
}
