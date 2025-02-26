package main

import "fmt"

func main() { //boolen values are mostly used for conditional testing
	var b1 bool = true //variable declaration eith data type and initial value
	var b2 = true      //variable declaration with only initial value
	var b3 bool        // variable declaration with data type and without intial value
	b4 := true         // variable declaration with only intial value

	fmt.Println(b1) //prints b1 value true
	fmt.Println(b2) //prints b2 value true
	fmt.Println(b3) //prints b3 value false as it is default value of bool
	fmt.Println(b4) //prints b4 value true
}
