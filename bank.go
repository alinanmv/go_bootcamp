package main

import "fmt"

func main() {
	accountBalance := 1000.0
	fmt.Println("WELCOME TO GO BANK")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2.Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice:")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Account Balance:", accountBalance)
	} else if choice == 2 {
		var depositMoney float64
		fmt.Print("Enter Money for Deposit: ")
		fmt.Scan(&depositMoney)
		accountBalance += depositMoney
		fmt.Println("Updated account balance:", accountBalance)
	} else if choice == 3 {
		var withdrawMoney float64
		fmt.Print("Enter Money for Withdraw: ")
		fmt.Scan(&withdrawMoney)
		accountBalance -= withdrawMoney
		fmt.Println("Updated account balance:", accountBalance)
	} else {
		fmt.Println("Goodbye")
	}
}
