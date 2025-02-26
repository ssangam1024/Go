package main

import "fmt"

func main() {
	var txt1 string = "Hello!" // declaring variable  txt1 with data type string and value in double quotes
	var txt2 string            // just declaring variable txt2 with data type string
	txt3 := "World 1"          // assigning variable txt3 with value world1 in double quotes

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1) //this function printf() first formats its first argument with formating verb given %T and then prints given txt1 data type and then second formating verb given %v and then prints given txt1 value.
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2) //this function printf() first formats its first argument with formating verb given %T and then prints given txt2 data type and then second formating verb given %v and then prints given txt2 value.
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3) //this function printf() first formats its first argument with formating verb given %T and then prints given txt3 data type and then second formating verb given %v and then prints given txt3 value.
}
