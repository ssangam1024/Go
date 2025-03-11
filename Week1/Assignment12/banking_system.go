package main

import "fmt"

type bankingAccount struct {
	AccountHolder string
	Balance       float64
}

func (deposit *bankingAccount) Deposit(depositammount float64) {
	deposit.Balance += depositammount
	fmt.Println("deposited:", depositammount)
}
func (withdraw *bankingAccount) withdraw(withdrawammount float64) {
	if withdraw.Balance >= withdrawammount {
		withdraw.Balance -= withdrawammount
		fmt.Println("withdraw:", withdrawammount)
	} else {
		fmt.Println("insuffient amount")
	}

}

func main() {
	var account bankingAccount
	fmt.Print("enter account holder name: ")
	fmt.Scan(&account.AccountHolder)

	fmt.Print("enter the balance: ")
	fmt.Scan(&account.Balance)

	var depositammount float64
	fmt.Println("enter deposit amount: ")
	fmt.Scan(&depositammount)
	account.Deposit(depositammount)

	var withdrawammount float64
	fmt.Println("enter withdraw amount: ")
	fmt.Scan(&withdrawammount)
	account.withdraw(withdrawammount)

	fmt.Println("final balance:", account.AccountHolder, account.Balance)

}
