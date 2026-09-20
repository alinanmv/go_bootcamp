package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	accountBalance, err := fileops.GetFloatFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)

	}
	fmt.Println("WELCOME TO GO BANK")
	fmt.Println(randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Account Balance:", accountBalance)
		case 2:
			var depositMoney float64
			fmt.Print("Enter Money for Deposit: ")
			fmt.Scan(&depositMoney)
			if depositMoney <= 0 {
				fmt.Println("Please enter a positive amount")
				continue
			}
			if depositMoney > accountBalance {
				fmt.Println("invalid input. deposit input can not be higher than account balance")
				continue
			}
			accountBalance += depositMoney
			fmt.Println("Updated account balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance)
		case 3:
			var withdrawMoney float64
			fmt.Print("Enter Money for Withdraw: ")
			fmt.Scan(&withdrawMoney)
			if withdrawMoney <= 0 {
				fmt.Println("Please enter a positive amount")
				continue
			}
			if withdrawMoney > accountBalance {
				fmt.Println("invalid input. withdraw input can not be higher than account balance")
				continue
			}
			accountBalance -= withdrawMoney
			fmt.Println("Updated account balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance)
		default:
			fmt.Println("Goodbye")
			return
		}
	}
}
