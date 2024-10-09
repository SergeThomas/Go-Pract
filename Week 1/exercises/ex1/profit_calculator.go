package main

import "fmt"

func main() {
	taxRate := 5.5 
	var revenue float64
	var operatingExpenses float64

	fmt.Print("Revenue amount: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses amount: ")
	fmt.Scan(&operatingExpenses)

	// tax amount 
	var taxAmount = taxRate / revenue * 100
	
	// EBT
	var ebtAmount = revenue + taxAmount

	// Net Profit
	var netProfit = ebtAmount - taxAmount - operatingExpenses

	ratio := ebtAmount / netProfit
	
	// Output Section
	fmt.Print("Earnings before tax: ")
	fmt.Println(ebtAmount)

	fmt.Print("Earnings after tax: ")
	fmt.Println(netProfit)

	fmt.Print("Ratio: ")
	fmt.Println(ratio)
}