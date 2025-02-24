package main

import (
	"fmt"
)

func main() {
	//in the bellow condition we declared the variable types but didnot assign any values to them
	var a string
	var b int
	var c bool

	fmt.Println(a) // here the default value will be printed as we didnot assign any value ( )
	fmt.Println(b) // here the default value will be printed as we didnot assign any value (0)
	fmt.Println(c) // here the default value will be printed as we didnot assign any value (false)
}
