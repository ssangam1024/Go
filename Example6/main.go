package main

import "fmt"

func main() {

	var ( // this  condition tells us variables declareted in parathesis means written in block
		a int              // here the variable a has int data type and there is no assigned value to it.
		b int    = 5       // here the variable b has int data type and it is assigned value is 5.
		c string = "hello" // here the variable c has string data type and the value  assigned init is hello
	)

	fmt.Println(a) // prints varable a value as 0
	fmt.Println(b) // prints b value as 5
	fmt.Println(c) //prints c value as hello

}
