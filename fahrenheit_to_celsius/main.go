package main

import "fmt"

func fahrenheit_to_celsius(fahrenheit float64) {

	var celsius int

	celsius = (int(fahrenheit) - 32) * 5 / 9
	fmt.Println(celsius)
}
func main() {
	var fahrenheit float64
	fmt.Println("Enter fahrenheit value:")
	fmt.Scan(&fahrenheit)

	fahrenheit_to_celsius(fahrenheit)
}
