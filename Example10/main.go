package main

import "fmt"

const ( // this condition is grouping together mlitiple constants into 1 block.
	// using 1 const keyword (menas the variable value is unchangeable and read only) and assigning constant A value as 1, B value as 3.14, C value as Hi.
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() { // this main functions calls constant A,B,C which is avilable in this file

	fmt.Println(A) // calls and prints constant A value
	fmt.Println(B) // calls and prints constant B value
	fmt.Println(C) // calls and prints constant C value

}
