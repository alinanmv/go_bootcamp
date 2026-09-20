package main

import (
	"fmt"
)

func main() {

	fmt.Println("EBT & Ratio calculator")

	revenue := getUserInput("Revenue:")
	expenses := getUserInput("Expenses:")
	taxRate := getUserInput("Tax Rate:")

	EBT := calculateEBT(revenue, expenses)
	fmt.Println("EBT:", EBT)

	profit := calculateProfit(EBT, taxRate)
	fmt.Println("Profit:", profit)

	ratio := calculateRatio(EBT, profit)
	fmt.Println("Ratio:", ratio)

	formattedProfit := fmt.Sprintf("Profit: %.1f\n", profit)
	formattedRatio := fmt.Sprintln("Ratio:", ratio)

	fmt.Print(formattedProfit)
	fmt.Print(formattedRatio)
}

func getUserInput(info string) float64 {
	var userInput float64
	fmt.Print(info)
	fmt.Scan(&userInput)
	return userInput
}

func calculateEBT(revenue, expenses float64) float64 {
	EBT := revenue - expenses
	return EBT
}
func calculateProfit(EBT, taxRate float64) float64 {
	profit := EBT * (1 - taxRate/100)
	return profit
}
func calculateRatio(EBT, profit float64) float64 {
	ratio := EBT / profit
	return ratio
}
