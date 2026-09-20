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
		if depositMoney <= 0 {
			fmt.Println("Please enter a positive amount")
			return
		}
		if depositMoney > accountBalance {
			fmt.Println("invalid input. deposit input can not be higher than account balance")
			return
		}
		accountBalance += depositMoney
		fmt.Println("Updated account balance:", accountBalance)

	} else if choice == 3 {
		var withdrawMoney float64
		fmt.Print("Enter Money for Withdraw: ")
		fmt.Scan(&withdrawMoney)
		if withdrawMoney <= 0 {
			fmt.Println("Please enter a positive amount")
			return
		}
		if withdrawMoney > accountBalance {
			fmt.Println("invalid input. withdraw input can not be higher than account balance")
			return
		}
		accountBalance -= withdrawMoney
		fmt.Println("Updated account balance:", accountBalance)
	} else {
		fmt.Println("Goodbye")
	}
}
