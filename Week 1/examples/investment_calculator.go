package main

import (
	"math"
	"fmt"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount  float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Enter Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter current return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years you are investing for: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate / 100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	
}
