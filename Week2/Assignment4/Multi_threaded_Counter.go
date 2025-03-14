package main

import (
	"fmt"
	"time"
)

func multiThreaded(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)

	}
}
func main() {
	var threaded1 int
	fmt.Print("Enter first mulithreaded number: ")
	fmt.Scan(&threaded1)

	var threaded2 int
	fmt.Print("Enter second mulithreaded number: ")
	fmt.Scan(&threaded2)

	var threaded3 int
	fmt.Print("Enter third multithreaded number: ")
	fmt.Scan(&threaded3)

	go multiThreaded(threaded1)
	go multiThreaded(threaded2)
	go multiThreaded(threaded3)

	time.Sleep(1 * time.Second)
}
