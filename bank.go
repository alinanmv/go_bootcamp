package main

import "fmt"

func main() {
	fmt.Println("WELCOME TO GO BANK")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2.Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scan(&choice)
}
