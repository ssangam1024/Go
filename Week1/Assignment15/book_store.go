package main

import "fmt"

type book struct {
	Title  string
	Author string
	Price  int
}

func (p *book) applydiscount() {
	if p.Price >= 50 {
		p.Price = p.Price - (p.Price * 10 / 100)
	}
}
func main() {
	var Book book

	fmt.Print("Enter Book Tittle: ")
	fmt.Scan(&Book.Title)

	fmt.Print("Enter Book Author: ")
	fmt.Scan(&Book.Author)

	fmt.Print("Enter Book Price: ")
	fmt.Scan(&Book.Price)

	Book.applydiscount()

	fmt.Println("Book Details: ")
	fmt.Println("Book tittle:", Book.Title)
	fmt.Println("Book author:", Book.Author)
	fmt.Println("Book price after discount:", Book.Price)

}
