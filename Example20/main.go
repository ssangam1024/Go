package main

import "fmt"

func main() {
	var txt = "Hello" // assigning value to varable txt as hello

	fmt.Printf("%s\n", txt)   //this function printf() first formats its argument with formating verb given %s and then prints given txt value as plain string and /n prints argument in new line.
	fmt.Printf("%q\n", txt)   //this function printf() first formats its argument with formating verb given %q and then prints given txt value in double-quoted string and /n prints argument in new line.
	fmt.Printf("%8s\n", txt)  //this function printf() first formats its argument with formating verb given %8s and then prints given txt value  asplain string and leaves space of width8 to rightside and /n prints argument in new line.
	fmt.Printf("%-8s\n", txt) //this function printf() first formats its argument with formating verb given %-8s and then prints given txt value  as plain string and leaves space of width8 on left side and /n prints argument in new line.
	fmt.Printf("%x\n", txt)   //this function printf() first formats its argument with formating verb given %x and then prints given txt value as hex dump of byte values  and /n prints argument in new line.
	fmt.Printf("% x\n", txt)  //this function printf() first formats its argument with formating verb given % x and then prints given txt value as hexdump with spaces and /n prints argument in new line.
}
