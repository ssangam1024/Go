package main // this tells us, this program belongs to main package, also its a main enty point of the application
import "fmt" // bring fmt files/funtions so that we can use them in this example



func main() {
	var student1 string = "John" // in this condition the both type and value of a variable is specified using var keyword
	var student2 = "Jane"        // this condition tells us we can also use var keyword without specifying type. but here the type of a variable is based on the value given.
	x := 2                       //in this statement we are creaing variable using ':='. here the type of a variable is based on value given.

	fmt.Println(student1) // this condition checks the 'student1' output value and prints the value using fmt package or functions in the fmt
	fmt.Println(student2) // this condition checks the 'student2' output value and prints the value using fmt package or functions in the fmt
	fmt.Println(x)        // this condition checks the 'x' output value and prints the value using fmt package or functions in the fmt
}
