package main

import "fmt"

func multiplyByTen(numb *int) {
	*numb = *numb * 10
}

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	fmt.Println("Before function call: ", number)

	multiplyByTen(&number)

	fmt.Println("After function call: ", number)

}
