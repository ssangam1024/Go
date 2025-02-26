package main

import "fmt"

func main() {
	var i = 15 // assigning varable i value as 15

	fmt.Printf("%b\n", i)   //this function printf() first formats its argument with formating verb given %b and then prints given i value  as Base 2 and /n prints argument in new line.
	fmt.Printf("%d\n", i)   //this function printf() first formats its argument with formating verb given %d and then prints given i value  as base10 and /n prints argument in new line.
	fmt.Printf("%+d\n", i)  //this function printf() first formats its argument with formating verb given %+d and then prints given i value  as bace 10 and always show sign and /n prints argument in new line.
	fmt.Printf("%o\n", i)   //this function printf() first formats its argument with formating verb given %o and then prints given i value  as base8 and /n prints argument in new line.
	fmt.Printf("%O\n", i)   //this function printf() first formats its argument with formating verb given %0 and then prints given i value  as base8 with leading0o and /n prints argument in new line.
	fmt.Printf("%x\n", i)   //this function printf() first formats its argument with formating verb given %x and then prints given i value  as base16 in lowercase and /n prints argument in new line.
	fmt.Printf("%X\n", i)   //this function printf() first formats its argument with formating verb given %X and then prints given i value asbase16 with uppercase and /n prints argument in new line.
	fmt.Printf("%#x\n", i)  //this function printf() first formats its argument with formating verb given %#x and then prints given i value as base16 eith leadin0x  and /n prints argument in new line.
	fmt.Printf("%4d\n", i)  //this function printf() first formats its argument with formating verb given %4d and then prints given i value with pad spaces of width4 on right side and /n prints argument in new line.
	fmt.Printf("%-4d\n", i) //this function printf() first formats its argument with formating verb given %-4d and then prints given i value with pad spaces of width4 on left side and /n prints argument in new line.
	fmt.Printf("%04d\n", i) //this function printf() first formats its argument with formating verb given %04d and then prints given i value  with pad zeroes of width4 and /n prints argument in new line.
}
