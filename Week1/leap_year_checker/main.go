package main

import "fmt"

func leapyear(year int) {
	if year == 0 {
		fmt.Println("it is a leap year")
	} else {
		fmt.Println("it is not a leap year")
	}

}

func main() {
	var year int
	fmt.Println("Enter year")
	fmt.Scan(&year)

	year = year % 4

	leapyear(year)

}
