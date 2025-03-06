package main

import (
	"fmt"
	"strconv"
)

func converting_int_to_string(number int) {

	strNum := strconv.Itoa(number)
	var arr []int

	for _, ch := range strNum {
		digit := int(ch - '0')
		arr = append(arr, digit)

	}
	fmt.Println(arr)

	var sumOfDigits int = 0

	for _, num := range arr {
		sumOfDigits = sumOfDigits + num
	}

	fmt.Println("sum of digits:", sumOfDigits)

}

func main() {
	var number int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)

	converting_int_to_string(number)

}
