package main

import (
	"fmt"
	"time"
)

func producer(ch chan int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println("producing: ", i)
		ch <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}
func consumer(ch chan int) {
	for num := range ch {
		fmt.Println("Consumed: ", num)
	}
}
func main() {
	var count int
	fmt.Print("Enter the number of values to produce: ")
	fmt.Scan(&count)

	ch := make(chan int)
	go producer(ch, count)
	go consumer(ch)
	time.Sleep(time.Second * 6)

}
