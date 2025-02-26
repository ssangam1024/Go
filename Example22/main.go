package main

import "fmt"

func main() {
	var i = 3.141 // assigning value to variable i as 3.141

	fmt.Printf("%e\n", i)    // this function printf() first formats its argument with formating verb given %e and then prints given i value in  scientific notation with e as exponent and /n prints argument in new line.
	fmt.Printf("%f\n", i)    // this function printf() first formats its argument with formating verb given %f and then prints given i value in decimal pointand no exponent and /n prints argument in new line.
	fmt.Printf("%.2f\n", i)  // this function printf() first formats its argument with formating verb given %.2f and then prints given i value in default width with precision2 and /n prints argument in new line.
	fmt.Printf("%6.2f\n", i) // this function printf() first formats its argument with formating verb given %6.2f and then prints given i value in width6 and precision2 and /n prints argument in new line.
	fmt.Printf("%g\n", i)    // this function printf() first formats its argument with formating verb given %g and then prints given i value in exponent as needed,only necessary digits and /n prints argument in new line.
}
