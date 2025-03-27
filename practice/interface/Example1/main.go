package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	name string
}
type Robot struct {
	model string
}

func (p Person) Speak() string {
	return "Hello, I am " + p.name
}
func (r Robot) Speak() string {
	return "I am a robot model " + r.model
}
func Introduce(speaker Speaker) {
	fmt.Println(speaker.Speak())
}
func main() {
	var person Person
	fmt.Print("Enter person's name: ")
	fmt.Scan(&person.name)

	var robot Robot
	fmt.Print("Enter robot's model name: ")
	fmt.Scan(&robot.model)

	Introduce(person)
	Introduce(robot)
}
