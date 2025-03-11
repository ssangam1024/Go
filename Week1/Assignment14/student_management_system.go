package main

import "fmt"

type student struct {
	Name  string
	Age   int
	Marks []int
}

func (s student) averagemarks() float64 {
	if len(s.Marks) == 0 {
		return 0
	}
	sum := 0
	for _, mark := range s.Marks {
		sum += mark
	}
	return float64(sum) / float64(len(s.Marks))

}
func main() {
	var Student student
	fmt.Print("enter name: ")
	fmt.Scan(&Student.Name)

	fmt.Print("enter age: ")
	fmt.Scan(&Student.Age)

	fmt.Print("Enter number of subjects: ")
	var numSubjects int
	fmt.Scanln(&numSubjects)

	Student.Marks = make([]int, numSubjects)
	fmt.Println("Enter marks:")
	for i := 0; i < numSubjects; i++ {
		fmt.Scan(&Student.Marks[i])
	}
	fmt.Println("\nStudent Details:")
	fmt.Println("Name:", Student.Name)
	fmt.Println("Age:", Student.Age)
	fmt.Println("Marks:", Student.Marks)
	fmt.Println("Average Marks:", student.averagemarks(Student))
}
