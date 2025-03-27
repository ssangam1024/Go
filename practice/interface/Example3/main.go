package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func PrintTypeInfo(val interface{}) {
	typeInfo := reflect.TypeOf(val)
	fmt.Printf("Value: %v, Type: %s\n", val, typeInfo)
}
func main() {
	var input string
	fmt.Print("Enter a value: ")
	fmt.Scanln(&input)
	
	if intVal, err := strconv.Atoi(input); err == nil {
		PrintTypeInfo(intVal)
	} else if floatVal, err := strconv.ParseFloat(input, 64); err == nil {
		PrintTypeInfo(floatVal)
	} else if input == "true" || input == "false" {
		boolVal, _ := strconv.ParseBool(input)
		PrintTypeInfo(boolVal)
	} else {
		PrintTypeInfo(input)
	}
}
