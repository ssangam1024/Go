package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func modifyage(p *Person, newAge int) {
	p.Age = newAge
}

func main() {
	var person Person

	fmt.Print("Enter name: ")
	fmt.Scan(&person.Name)

	fmt.Print("Enter age: ")
	fmt.Scan(&person.Age)

	fmt.Println("Before modification: ", person.Age)

	var newAge int
	fmt.Print("Enter new age: ")
	fmt.Scan(&newAge)

	modifyage(&person, newAge)

	fmt.Println("After modification: ", person)

}
