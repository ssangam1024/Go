package main

import "fmt"

func greaterValue(numb1, numb2, numb3 int) {
	if (numb1 >= numb2) && (numb1 >= numb3) {
		fmt.Println(numb1)
	} else if (numb2 >= numb1) && (numb2 >= numb3) {
		fmt.Println(numb2)
	} else if (numb3 >= numb1) && (numb3 >= numb2) {
		fmt.Println(numb3)
	}
	return
}

func main() {
	var numb1 int
	fmt.Println("enter first number")
	fmt.Scan(&numb1)
	var numb2 int
	fmt.Println("enter second number")
	fmt.Scan(&numb2)
	var numb3 int
	
	fmt.Println("enter third number")
	fmt.Scan(&numb3)

	greaterValue(numb1, numb2, numb3)

}
