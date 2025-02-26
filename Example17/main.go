package main

import "fmt"

func main() {
	var i string = "Hello" // assigning variable i with data type string and value hello
	var j int = 15         // assiging variable j as data type int and value 15

	fmt.Printf("i has value: %v and type: %T\n", i, i) // this function printf() first formats its argument with formating verb %v(value) and %T(type) and then prints given i value
	fmt.Printf("j has value: %v and type: %T", j, j)   // this function printf() first formats its argument with formating verb %v(value) and %T(type) and then prints given j value
}
