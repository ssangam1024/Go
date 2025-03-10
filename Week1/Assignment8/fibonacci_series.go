package main

import "fmt"

func fibonacci(numb int) {
	var a int
	a = 0
	var b int
	b = 1
	for c := 0; c < numb; c++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}
func main() {
	var numb int
	fmt.Print("enter the number of fibo term:")
	fmt.Scan(&numb)
	fibonacci(numb)

}
