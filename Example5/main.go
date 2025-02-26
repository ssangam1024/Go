package main

import "fmt"

func main() {

	var a, b = 6, "Hello" // this condition using var key, not specifing type keyword and declaring different type of variables (a,b) in sameline with the values  (6,Hello)

	c, d := 7, "World" // this condition without specifing type keyword and declaring dirrent type of variables in sameline

	fmt.Println(a) //prints given a value as 6
	fmt.Println(b) //prints given b value as Hello
	fmt.Println(c) //prints given c value as 7
	fmt.Println(d) //prints given d value as World

}
