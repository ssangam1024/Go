package main

import "fmt"

func main() {
	var Principal float64
	fmt.Println("Enter Principal Amount")
	fmt.Scan(&Principal)

	var Rate float64
	fmt.Println("Enter Rate Amount")
	fmt.Scan(&Rate)

	var Time int
	fmt.Println("Enter Time in Years")
	fmt.Scan(&Time)

	var simple_interest float64
	simple_interest = (Principal * Rate * float64(Time)) / 100

	fmt.Println("Total interent:", simple_interest)

}
