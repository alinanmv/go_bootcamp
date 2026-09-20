package main

import (
	"fmt"
	"os"
)

func getUserInput(info string) float64 {
	var userInput float64
	fmt.Print(info)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		fmt.Println("Invalid input")
		fmt.Print(info)
		fmt.Scan(&userInput)
	}
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

const file = "results.txt"

func writeResultsToFile(revenue, expenses, taxRate, EBT, profit, ratio float64) {
	resultsText := fmt.Sprintf("Revenue: %.2f\nExpenses: %.2f\nTax Rate: %.2f\nEBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", revenue, expenses, taxRate, EBT, profit, ratio)
	os.WriteFile(file, []byte(resultsText), 0644)
}

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

	writeResultsToFile(revenue, expenses, taxRate, EBT, profit, ratio)

	formattedProfit := fmt.Sprintf("Profit: %.1f\n", profit)
	formattedRatio := fmt.Sprintf("Ratio: %.1f\n", ratio)

	fmt.Print(formattedProfit)
	fmt.Print(formattedRatio)

}
