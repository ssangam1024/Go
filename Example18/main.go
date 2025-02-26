package main

import (
	"fmt"
)

func main() {
	var i = 15.5             // assigning varable i value as 15.5
	var txt = "Hello World!" // assigning varable txt value as hello world

	fmt.Printf("%v\n", i)   //this function printf() first formats its argument with formating verb given %v and then prints given i value in the default format and /n prints argument in new line.
	fmt.Printf("%#v\n", i)  //this function printf() first formats its argument with formating verb given %#v and then prints given i value  in Go-syntax format and /n prints argument in new line.
	fmt.Printf("%v%%\n", i) //this function printf() first formats its argument with formating verb given %v%% and then prints given i value  in default format with %sign and /n prints argument in new line.
	fmt.Printf("%T\n", i)   //this function printf() first formats its argument with formating verb given %T and then prints given i data type value and /n prints argument in new line.

	fmt.Printf("%v\n", txt)  //this function printf() first formats its argument with formating verb given %v and then prints given txt value in the default format and /n prints argument in new line.
	fmt.Printf("%#v\n", txt) //this function printf() first formats its argument with formating verb given %#v and then prints given txt value in Go-syntax format and /n prints argument in new line.
	fmt.Printf("%T\n", txt)  //this function printf() first formats its argument with formating verb given %T and then prints given txt data type value  and /n prints argument in new line.
}
