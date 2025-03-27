package main

import "fmt"

func main() {
	var numb int
	fmt.Print("Enter a number: ")
	fmt.Scan(&numb)

	var numb2 int
	fmt.Print("Enter a number to modify variable through pointer: ")
	fmt.Scan(&numb2)

	fmt.Println("Before modification: ", numb)

	var ptr *int = &numb
	*ptr = numb2
	fmt.Println("After modification: ", numb2)

}
