package main

import (
	"fmt"
)

func main() {
	var x float32 = 123.78                    // declaring x variable with data type float32 and value 123.78
	var y float32 = 3.4e+38                   // declaring y variable with data type float32 and value 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x, x) //this function printf() first formats its argument with formating verb given %T and then prints given x data type and formating verb given %v and then prints given x value.
	fmt.Printf("Type: %T, value: %v", y, y)   //this function printf() first formats its argument with formating verb given %T and then prints given y data type and formating verb given %v and then prints given y value.
}
