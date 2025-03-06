package main

import "fmt"

func main() {

	var x int
	fmt.Println("Enter Multiplication Table of:")
	fmt.Scan(&x)

	fmt.Println("Multiplication table of", x)

	var y int
	for y = 1; y < 11; y++ {

		fmt.Println(x, "*", y, "=", x*y)
	}

}
