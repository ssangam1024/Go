package main

import "fmt"

func main() {
	var x uint = 500                        //declaring x variable with data type as uint and intial value.
	var y uint = 4500                       // declaring y variable with data type as uint and intial value.
	fmt.Printf("Type: %T, value: %v", x, x) //this function printf() first formats its argument with formating verb given %T and then prints given x data type and formating verb given %v and then prints given x value.
	fmt.Printf("Type: %T, value: %v", y, y) //this function printf() first formats its argument with formating verb given %T and then prints given y data type and formating verb given %v and then prints given y value.
}
