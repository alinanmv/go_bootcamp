package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const file = "test.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return 1000, errors.New("no file was found")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("can't convert to a float")
	}
	return balance, nil
}
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(file, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)

	}
	fmt.Println("WELCOME TO GO BANK")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2.Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye")
			return
		}
	}
}
