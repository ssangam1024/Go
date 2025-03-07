package main

import "fmt"

func division(numb1, numb2 int) {

	var numb3 int
	if numb2 == 0 {
		fmt.Println("Error - you cannot assign zero as denominator")
	} else {
		numb3 = numb1 / numb2
		fmt.Println(numb3)
	}
}
func main() {
	var numb1 int
	fmt.Print("Enter numerator number: ")
	fmt.Scan(&numb1)

	var numb2 int
	fmt.Print("enter denominator number:")
	fmt.Scan(&numb2)

	division(numb1, numb2)

}
