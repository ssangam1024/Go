package main

import "fmt"

type numb struct {
	orginalNumber int
}

func changnumber(orginal *numb) {
	var x int
	fmt.Print("Enter swaping number:")
	fmt.Scan(&x)
	orginal.orginalNumber = x
}

func main() {
	var orginalInput int
	fmt.Print("Enter orginal number: ")
	fmt.Scan(&orginalInput)

	orginal:= numb{orginalNumber: orginalInput}

	changnumber(&orginal)

	fmt.Println("Updated number:",orginal.orginalNumber)

}
