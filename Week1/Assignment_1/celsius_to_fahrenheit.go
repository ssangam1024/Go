package main

import "fmt"

func celsius_to_Fahrenhrit(celsius int) {
	var Fahrenheit float64
	Fahrenheit = (float64(celsius) * 9 / 5) + 32
	fmt.Println(Fahrenheit)

}
func main() {
	var celsius int
	fmt.Println("Enter celsius value:")
	fmt.Scan(&celsius)

	celsius_to_Fahrenhrit(celsius)

}
