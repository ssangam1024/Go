package main

import "fmt"

/*func employeeDetails (name string, age int, salary int) {
	fmt.Println("name: ", name)
	fmt.Println("age:", age)
	fmt.Println("salary:", salary)
}

func main() {
	var name string
	fmt.Print("Enter the Name of Employee:")
	fmt.Scan(&name)

	var age int
	fmt.Print("Enter the Age of Employee:")
	fmt.Scan(&age)

	var salary int
	fmt.Print("Enter the Salary of the Employee:")
	fmt.Scan(&salary)

	employeeDetails(name, age, salary)

}*/

type employeeDetails struct {
	name   string
	age    int
	salary int
}

func main() {
	var employee employeeDetails
	fmt.Print("Enter the Name of Employee:")
	fmt.Scan(&employee.name)

	fmt.Print("Enter the Age of Employee:")
	fmt.Scan(&employee.age)

	fmt.Print("Enter the Salary of the Employee:")
	fmt.Scan(&employee.salary)

	fmt.Println("Employee Details:", employee)

}
